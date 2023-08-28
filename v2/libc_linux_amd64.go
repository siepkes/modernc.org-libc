// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

//TODO https://gcc.gnu.org/onlinedocs/gcc-6.2.0/gcc/Integer-Overflow-Builtins.html

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"modernc.org/mathutil"
	"modernc.org/memory"
)

var (
	initLibcOnce sync.Once

	allocator   memory.Allocator
	allocatorMu sync.Mutex

	argc int32   // Value established in initLibc
	argv uintptr // Value established in initLibc

	tid atomic.Int32 // TLS Go ID
)

// Start executes C's main.
func Start(main func(*TLS, int32, uintptr) int32) {
	runtime.LockOSThread()
	tls := NewTLS()
	Xexit(tls, main(tls, argc, argv))
}

func initLibc(tls *TLS) {
	var args []uintptr
	for _, v := range os.Args {
		p := mustPrivateCalloc(len(v) + 1)
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		args = append(args, p)
	}
	args = append(args, 0)
	for _, v := range os.Environ() {
		p := mustPrivateCalloc(len(v) + 1)
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		args = append(args, p)
	}
	args = append(args, 0)
	argv = mustPrivateCalloc(len(args) * int(unsafe.Sizeof(uintptr(0))))
	copy(unsafe.Slice((*uintptr)(unsafe.Pointer(argv)), len(args)), args)

	// mov $main,%rdi  /* 1st arg: application entry ip */
	// pop %rsi        /* 2nd arg: argc */
	// mov %rsp,%rdx   /* 3rd arg: argv */
	// xor %rcx,%rcx   /* 4th arg: always 0 */
	// xor %r8,%r8     /* 5th arg: always 0 */
	// mov %rdx,%r9    /* 6th arg: ptr to register with atexit() */

	argc = int32(len(os.Args))
	x___libc_start_main(tls, 0, argc, argv, 0, 0, 0)
}

// CString returns a pointer to a zero-terminated version of s. The caller is
// responsible for freeing the allocated memory using Xfree.
func CString(s string) (uintptr, error) {
	n := len(s)
	p := Xmalloc(nil, uint64(n)+1)
	if p == 0 {
		return 0, fmt.Errorf("CString: cannot allocate %d bytes", n+1)
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), n), s)
	*(*byte)(unsafe.Pointer(p + uintptr(n))) = 0
	return p, nil
}

// GoBytes returns a byte slice from a C char* having length len bytes.
func GoBytes(s uintptr, len int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(s)), len)
}

// GoString returns the value of a C string at s.
func GoString(s uintptr) string {
	if s == 0 {
		return ""
	}

	var buf []byte
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			return string(buf)
		}

		buf = append(buf, b)
		s++
	}
}

// Go libc private heap, outside of the C heap.
func mustPrivateCalloc(sz int) uintptr {
	if p, _ := privateCalloc(sz); p != 0 || sz == 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func mustPrivateMalloc(sz int) uintptr {
	if p, _ := privateMalloc(sz); p != 0 || sz == 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func mustPrivateRealloc(p uintptr, sz int) uintptr {
	if p, _ := privateRealloc(p, sz); p != 0 || sz == 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func privateCalloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrCalloc(sz)
}

// Go libc private heap, outside of the C heap.
func privateMalloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrMalloc(sz)
}

// Go libc private heap, outside of the C heap.
func privateRealloc(p uintptr, sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrRealloc(p, sz)
}

// Go libc private heap, outside of the C heap.
func privateFree(p uintptr) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	allocator.UintptrFree(p)
}

type tlsStack struct {
	p  uintptr
	sz int
}

// TLS emulates thread local storage. TLS is not safe for concurrent use by
// multiple goroutines.
type TLS struct {
	allocas        []uintptr
	freeFS         uintptr
	fs             uintptr // *T__pthread
	sp             int
	stack          []tlsStack
	threadFuncWait chan struct{}

	ID int32
}

// NewTLS returns a newly created TLS that must be eventually closed to prevent
// resource leaks.
func NewTLS() (r *TLS) {
	//TODO keep count of TLSs, notify libc when more than one and when back to one.
	r = &TLS{
		ID: tid.Add(1),
		fs: mustPrivateMalloc(int(unsafe.Sizeof(T__pthread{}))),
	}
	r.freeFS = r.fs
	*(*T__pthread)(unsafe.Pointer(r.fs)) = T__pthread{
		Ftid: r.ID,
	}
	initLibcOnce.Do(func() { initLibc(r) })
	if r.ID > 1 {
		x___libc.Fneed_locks = -1
	}
	return r
}

// var nallocs, nmallocs, nreallocs int //TODO-

func (t *TLS) Alloc(n int) (r uintptr) {
	n = int(roundup(uintptr(n), 16))
	// shrink	stats									speedtest1
	// -----------------------------------------------------------------------------------------------
	//    0		total  2,544, nallocs 107,553,070, nmallocs 25, nreallocs 107,553,045	10.984s
	//    1		total  2,544, nallocs 107,553,070, nmallocs 25, nreallocs  38,905,980	 9.597s
	//    2		total  2,616, nallocs 107,553,070, nmallocs 25, nreallocs  18,201,284	 9.206s
	//    3		total  2,624, nallocs 107,553,070, nmallocs 25, nreallocs  16,716,302	 9.155s
	//    4		total  2,624, nallocs 107,553,070, nmallocs 25, nreallocs  16,156,102	 9.398s
	//    8		total  3,408, nallocs 107,553,070, nmallocs 25, nreallocs  14,364,274	 9.198s
	//   16		total  3,976, nallocs 107,553,070, nmallocs 25, nreallocs   6,219,602	 8.910s
	// ---------------------------------------------------------------------------------------------
	//   32		total  5,120, nallocs 107,553,070, nmallocs 25, nreallocs   1,089,037	 8.836s
	// ---------------------------------------------------------------------------------------------
	//   64		total  6,520, nallocs 107,553,070, nmallocs 25, nreallocs       1,788	 8.420s
	//  128		total  8,848, nallocs 107,553,070, nmallocs 25, nreallocs       1,098	 8.833s
	//  256		total  8,848, nallocs 107,553,070, nmallocs 25, nreallocs       1,049	 9.508s
	//  512		total 33,336, nallocs 107,553,070, nmallocs 25, nreallocs          88	 8.667s
	// none		total 33,336, nallocs 107,553,070, nmallocs 25, nreallocs          88	 8.408s

	const shrinkSegment = 32
	// nallocs++
	sp := t.sp
	if sp < len(t.stack) {
		s := t.stack[sp]
		if s.sz >= n && s.sz <= shrinkSegment*n {
			t.sp++
			return s.p
		}

		r = mustPrivateRealloc(s.p, n)
		// nreallocs++
		t.stack[sp].p = r
		t.stack[sp].sz = n
		t.sp++
		return r

	}

	r = mustPrivateMalloc(n)
	// nmallocs++
	t.stack = append(t.stack, tlsStack{p: r, sz: n})
	t.sp++
	return r
}

func (t *TLS) alloca(n uint64) (r uintptr) {
	r = mustPrivateMalloc(int(n))
	t.allocas = append(t.allocas, r)
	return r
}

func (t *TLS) Free(n int) {
	//TODO shrink stacks
	t.sp--
}

func (t *TLS) Close() {
	// var sum int
	// for i, v := range t.stack {
	// 	sum += v.avail
	// 	trc("%4d: %#0x, avail %v", i, v.p, v.avail)
	// }
	// trc("total %v, nallocs %v, nmallocs %v, nreallocs %v", sum, nallocs, nmallocs, nreallocs)

	defer func() { *t = TLS{} }()

	for _, v := range t.allocas {
		privateFree(v)
	}
	for _, v := range t.stack[:t.sp] {
		privateFree(v.p)
	}
	privateFree(t.freeFS)
}

// musl-1.2.4/src/env/__init_tls.c:81:extern weak hidden const Tsize_t _DYNAMIC[];
var __DYNAMIC [1]uint64

var ___cp_begin [1]int8
var ___cp_cancel [1]int8
var ___cp_end [1]int8

func ___clone(tls *TLS, func1 uintptr, stack uintptr, flags int32, arg uintptr, va uintptr) (r int32) {
	return -m_ENOSYS
}

func ___tlsdesc_dynamic(tls *TLS) int64 {
	panic(todo(""))
}

func ___tlsdesc_static(tls *TLS) int64 {
	panic(todo(""))
}

func _sqrt(tls *TLS, x float64) float64 {
	return math.Sqrt(x)
}

func _sqrtf(tls *TLS, x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func ___restore_rt(tls *TLS) {
	panic(todo(""))
}

func _longjmp(tls *TLS, env uintptr, val int32) int64 {
	panic(todo(""))
}

func Xlongjmp(tls *TLS, env uintptr, val int32) int64 {
	panic(todo(""))
}

func X__builtin_longjmp(tls *TLS, env uintptr, val int32) int64 {
	panic(todo(""))
}

// Byte stores are atomic on this CPU.
func a_store_8(addr uintptr, val int32) {
	*(*byte)(unsafe.Pointer(addr)) = byte(val)
}

// Byte loads are atomic on this CPU.
func a_load_8(addr uintptr) uint32 {
	return uint32(*(*byte)(unsafe.Pointer(addr)))
}

// int16 loads are atomic on this CPU when properly aligned.
func a_load_16(addr uintptr) uint32 {
	if addr&1 != 0 {
		panic(fmt.Errorf("unaligned atomic 16 bit access at %#0x", addr))
	}

	return uint32(*(*uint16)(unsafe.Pointer(addr)))
}

func _a_crash(tls *TLS) {
	panic(todo(""))
}

func a_cas(p uintptr, t, s int32) int32

func _a_cas(tls *TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

func a_spin()

func _a_spin(tls *TLS) {
	a_spin()
}

func a_fetch_add(p uintptr, v int32) int32

func _a_fetch_add(tls *TLS, p uintptr, v int32) int32 {
	return a_fetch_add(p, v)
}

// func a_and_64(p uintptr, x uint64)
//
// func _a_and_64(tls *TLS, p uintptr, x uint64) {
// 	a_and_64(p, x)
// }

func _a_and(tls *TLS, p uintptr, x int32) {
	panic(todo(""))
}

func a_cas_p(p uintptr, t, s uintptr) uintptr

func _a_cas_p(tls *TLS, p, test, s uintptr) uintptr {
	return a_cas_p(p, test, s)
}

func a_or(p uintptr, v int32)

func _a_or(tls *TLS, p uintptr, v int32) {
	a_or(p, v)
}

func a_or_64(p uintptr, v uint64)

func _a_or_64(tls *TLS, p uintptr, v uint64) {
	a_or_64(p, v)
}

func _a_ctz_64(tls *TLS, x uint64) int32 {
	return int32(bits.TrailingZeros64(x))
}

func _a_ctz_32(tls *TLS, x uint32) int32 {
	return X__builtin_ctz(tls, x)
}

func X__builtin_ctz(tls *TLS, x uint32) int32 {
	return int32(bits.TrailingZeros32(x))
}

func _a_clz_64(tls *TLS, x uint64) int32 {
	return int32(bits.LeadingZeros64(x))
}

func _a_inc(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), 1)
}

func _a_dec(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), -1)
}

func _a_swap(tls *TLS, p uintptr, v int32) int32 {
	return atomic.SwapInt32((*int32)(unsafe.Pointer(p)), v)
}

func _a_store(tls *TLS, p uintptr, v int32) {
	atomic.StoreInt32((*int32)(unsafe.Pointer(p)), v)
}

func a_barrier()

func _a_barrier(tls *TLS) {
	a_barrier()
}

func ___syscall0(tls *TLS, n int64) int64 {
	// #define __NR_sched_yield		24
	const m___NR_sched_yield = 24
	switch n {
	case m___NR_sched_yield:
		runtime.Gosched()
		return 0
	default:
		r1, _, err := syscall.Syscall(uintptr(n), 0, 0, 0)
		if err != 0 {
			return int64(-err)
		}

		return int64(r1)
	}
}

func ___syscall1(tls *TLS, n, a1 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), 0, 0)
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func ___syscall2(tls *TLS, n, a1, a2 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), 0)
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func ___syscall3(tls *TLS, n, a1, a2, a3 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3))
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func ___syscall4(tls *TLS, n, a1, a2, a3, a4 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), 0, 0)
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func ___syscall5(tls *TLS, n, a1, a2, a3, a4, a5 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), 0)
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func ___syscall6(tls *TLS, n, a1, a2, a3, a4, a5, a6 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), uintptr(a6))
	if err != 0 {
		return int64(-err)
	}

	return int64(r1)
}

func _fork(tls *TLS) int32 {
	panic(todo(""))
}

func ___unmapself(tls *TLS, base uintptr, size uint64) {
	panic(todo(""))
}

func Xsignal(tls *TLS, sig int32, function uintptr) (r uintptr) {
	return x_signal(tls, sig, function)
}

// func ___uniclone(tls *TLS, start, stack, new uintptr) int32 {
// 	panic(todo(""))
// }

// ----------------------------------------------------------------------------

// Xfmod computes the floating-point remainder of dividing x by y.
func Xfmod(tls *TLS, x float64, y float64) (r float64) {
	return x_fmod(tls, x, y)
}

// Xvsnprintf() produces output according to a format. Xvsnprintf writes to the
// character string str.
func Xvsnprintf(tls *TLS, str uintptr, n uint64, format uintptr, ap uintptr) (r int32) {
	return x_vsnprintf(tls, str, n, format, ap)
}

// X__builtin_vsnprintf() is equivalent to Xvsnprintf.
func X__builtin_vsnprintf(tls *TLS, str uintptr, n uint64, format uintptr, ap uintptr) (r int32) {
	return Xvsnprintf(tls, str, n, format, ap)
}

// Xsprintf provides formatted output conversion. Xsprintf writes to the
// character string str.
func Xsprintf(tls *TLS, str uintptr, fmt uintptr, va uintptr) (r int32) {
	return x_sprintf(tls, str, fmt, va)
}

// X__builtin_sprintf is equivalent to Xsprintf.
func X__builtin_sprintf(tls *TLS, str uintptr, fmt uintptr, va uintptr) (r int32) {
	return Xsprintf(tls, str, fmt, va)
}

// Xstrtod converts the initial portion of the string pointed to by nptr to
// double representation.
func Xstrtod(tls *TLS, nptr, endptr uintptr) float64 {
	return x_strtod(tls, nptr, endptr)
}

// Xrint rounds its argument to an integer value in floating-point format,
// using the current rounding direction.
func Xrint(tls *TLS, x float64) float64 {
	return x_rint(tls, x)
}

func Xrintf(tls *TLS, x float32) (r float32) {
	return x_rintf(tls, x)
}

func X__builtin_rintf(tls *TLS, x float32) (r float32) {
	return Xrintf(tls, x)
}

// Xmemset fills the first n bytes of the memory area pointed to by s with the constant byte c.
func Xmemset(tls *TLS, s uintptr, c int32, n uint64) uintptr {
	return _memset(tls, s, c, n)
}

// X__builtin_memset is equivalent to Xmemset
func X__builtin_memset(tls *TLS, s uintptr, c int32, n uint64) uintptr {
	return Xmemset(tls, s, c, n)
}

func _memset(tls *TLS, s uintptr, c int32, n uint64) uintptr {
	if n != 0 {
		c := byte(c & 0xff)
		b := unsafe.Slice((*byte)(unsafe.Pointer(s)), n)
		for i := range b {
			b[i] = c
		}
	}
	return s
}

// Xsnprintf writes at most size bytes (including the terminating null byte
// ('\0')) to str, according to a format and args.
func Xsnprintf(tls *TLS, str uintptr, size uint64, format, args uintptr) int32 {
	return x_snprintf(tls, str, size, format, args)
}

// X__builtin_snprintf is equivalent to Xsnprintf.
func X__builtin_snprintf(tls *TLS, str uintptr, size uint64, format, args uintptr) int32 {
	return Xsnprintf(tls, str, size, format, args)
}

// Xfdopen associates a stream with the existing file descriptor, fd.
func Xfdopen(tls *TLS, fd int32, mode uintptr) uintptr {
	return x___fdopen(tls, fd, mode)
}

// Xfread reads nmemb items of data, each size bytes long, from the stream
// pointed to by stream, storing them at the location given by ptr.
func Xfread(tls *TLS, ptr uintptr, size, nmemb uint64, stream uintptr) uint64 {
	return x_fread(tls, ptr, size, nmemb, stream)
}

// Xmalloc allocates size bytes and returns a pointer to the allocated
// memory.
func Xmalloc(tls *TLS, n uint64) (r uintptr) {
	return _default_malloc(tls, n)
}

// X__builtin_malloc is equivalent to Xmalloc.
func X__builtin_malloc(tls *TLS, n uint64) (r uintptr) {
	return Xmalloc(tls, n)
}

// The Xfree function frees the memory space pointed to by ptr, which must have
// been returned by a previous call to Xmalloc(), Xcalloc(), or Xrealloc().
func Xfree(tls *TLS, p uintptr) {
	x_free(tls, p)
}

// The X__builtin_free is equivalent to Xfree.
func X__builtin_free(tls *TLS, p uintptr) {
	Xfree(tls, p)
}

// Xcalloc allocates memory for an array of nmemb elements of size bytes each
// and returns a pointer to the allocated memory.
func Xcalloc(tls *TLS, nmemb, size uint64) (r uintptr) {
	return x_calloc(tls, nmemb, size)
}

// Xprintf produces output according to a format. It writes output to stdout,
// the standard output stream;
func Xprintf(tls *TLS, format, va uintptr) (r int32) {
	return x_printf(tls, format, va)
}

// X__builtin_printf is equivalent to Xprintf.
func X__builtin_printf(tls *TLS, format, va uintptr) (r int32) {
	return Xprintf(tls, format, va)
}

// Xexp returns the value of e (the base of natural logarithms) raised to the
// power of x.
func Xexp(tls *TLS, x float64) (r float64) {
	return x_exp(tls, x)
}

// Xlog returns the natural logarithm of x.
func Xlog(tls *TLS, x float64) (r float64) {
	return x_log(tls, x)
}

// Xlog10 returns the base 10 logarithm of x.
func Xlog10(tls *TLS, x float64) (r float64) {
	return x_log10(tls, x)
}

// Xround rounds x to the nearest integer,
func Xround(tls *TLS, x float64) (r float64) {
	return x_round(tls, x)
}

func X__builtin_round(tls *TLS, x float64) (r float64) {
	return Xround(tls, x)
}

// Xceil returns the smallest integral value that is not less than x.
func Xceil(tls *TLS, x float64) float64 {
	return math.Ceil(x)
}

// Xsinh returns the hyperbolic sine of x,
func Xsinh(tls *TLS, x float64) (r float64) {
	return x_sinh(tls, x)
}

// Xcosh returns the hyperbolic cosine of x.
func Xcosh(tls *TLS, x float64) (r float64) {
	return x_cosh(tls, x)
}

// Xtanh returns the hyperbolic tangent of x.
func Xtanh(tls *TLS, x float64) (r float64) {
	return x_tanh(tls, x)
}

// Xatan returns the arc tangent of x, where x is given in radians.
func Xatan(tls *TLS, x float64) (r float64) {
	return x_atan(tls, x)
}

// Xatanf returns the arc tangent of x, where x is given in radians.
func Xatanf(tls *TLS, x float32) (r float32) {
	return x_atanf(tls, x)
}

// Xtan returns the tangent of x, where x is given in radians.
func Xtan(tls *TLS, x float64) (r float64) {
	return x_tan(tls, x)
}

// Xtanf returns the tangent of x, where x is given in radians.
func Xtanf(tls *TLS, x float32) (r float32) {
	return x_tanf(tls, x)
}

// Xacos calculatse the arc cosine of x; that is the value whose cosine is x.
func Xacos(tls *TLS, x float64) (r float64) {
	return x_acos(tls, x)
}

// Xsin returns the sine of x, where x is given in radians.
func Xsin(tls *TLS, x float64) (r float64) {
	return x_sin(tls, x)
}

// Xsinf returns the sine of x, where x is given in radians.
func Xsinf(tls *TLS, x float32) (r float32) {
	return x_sinf(tls, x)
}

// Xstrcpy copies the string pointed to by src, including the terminating null
// byte ('\0'), to the buffer pointed to by dest.
func Xstrcpy(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	return x_strcpy(tls, dest, src)
}

// X__builtin_strcpy is equivalent to Xstrcpy.
func X__builtin_strcpy(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	return Xstrcpy(tls, dest, src)
}

// Xstrncpy is line Xstrcpy, except that at most n bytes of src are copied.
func Xstrncpy(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return x_strncpy(tls, dest, src, n)
}

func X__builtin_strncpy(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return Xstrncpy(tls, dest, src, n)
}

// Xstrcmp compares the two strings s1 and s2.
func Xstrcmp(tls *TLS, s1 uintptr, s2 uintptr) (r int32) {
	return x_strcmp(tls, s1, s2)
}

// X__builtin_strcmp is equivalent to Xstrcmp.
func X__builtin_strcmp(tls *TLS, s1 uintptr, s2 uintptr) (r int32) {
	return Xstrcmp(tls, s1, s2)
}

// Xstrlen function calculates the length of the string pointed to by s,
// excluding the terminating null byte ('\0').
func Xstrlen(tls *TLS, s uintptr) (r uint64) {
	return x_strlen(tls, s)
}

// X__builtin_strlen is equivalent to Xstrlen.
func X__builtin_strlen(tls *TLS, s uintptr) (r uint64) {
	return Xstrlen(tls, s)
}

// Xstrcat appends the src string to the dest string, overwriting the
// terminating null byte ('\0') at the end of dest, and then adds a terminating
// null byte.
func Xstrcat(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	return x_strcat(tls, dest, src)
}

// Xstrncmp is line Xstrcmp, except it compares only the first (at most) n
// bytes of s1 and s2.
func Xstrncmp(tls *TLS, s1 uintptr, s2 uintptr, n uint64) (r1 int32) {
	return x_strncmp(tls, s1, s2, n)
}

// Xstrchr returns a pointer to the first occurrence of the character c in the
// string s.
func Xstrchr(tls *TLS, s uintptr, c int32) (r uintptr) {
	return x_strchr(tls, s, c)
}

// Xstrrchr returns a pointer to the last occurrence of the character c in the
// string s.
func Xstrrchr(tls *TLS, s uintptr, c int32) (r uintptr) {
	return x_strrchr(tls, s, c)
}

// Xmemcpy copies n bytes from memory area src to memory area dest.
func Xmemcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return _memcpy(tls, dest, src, n)
}

// X__builtin_memcpy is equivalent to Xmemcpy.
func X__builtin_memcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return Xmemcpy(tls, dest, src, n)
}

func _memcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return _memmove(tls, dest, src, n)
}

// Xmemcmp compares the first n bytes (each interpreted as unsigned char) of
// the memory areas s1 and s2.
func Xmemcmp(tls *TLS, s1 uintptr, s2 uintptr, n uint64) (r1 int32) {
	return x_memcmp(tls, s1, s2, n)
}

// X__builtin_memcmp is equivalent to Xmemcmp.
func X__builtin_memcmp(tls *TLS, s1 uintptr, s2 uintptr, n uint64) (r1 int32) {
	return Xmemcmp(tls, s1, s2, n)
}

// Xfopen opens the file whose name is the string pointed to by pathname and
// associates a stream with it.
func Xfopen(tls *TLS, pathname uintptr, mode uintptr) (r uintptr) {
	return x_fopen(tls, pathname, mode)
}

// Xfwrite writes nmemb items of data, each size bytes long, to the stream
// pointed to by stream, obtaining them from the location given by ptr.
func Xfwrite(tls *TLS, ptr uintptr, size uint64, nmemb uint64, stream uintptr) (r uint64) {
	return x_fwrite(tls, ptr, size, nmemb, stream)
}

// Xfclose flushes the stream pointed to by stream (writing any buffered output
// data using fflush(3)) and closes the underlying file descriptor.
func Xfclose(tls *TLS, stream uintptr) (r int32) {
	return x_fclose(tls, stream)
}

// Xfgetc reads the next character from stream and returns it as an unsigned
// char cast to an int, or EOF on end of file or error.
func Xfgetc(tls *TLS, stream uintptr) (r int32) {
	return x_fgetc(tls, stream)
}

// Xfprintf is like printf(), except it writes output to the given output
// stream.
func Xfprintf(tls *TLS, stream uintptr, fmt uintptr, va uintptr) (r int32) {
	return x_fprintf(tls, stream, fmt, va)
}

func X__builtin_fprintf(tls *TLS, stream uintptr, fmt uintptr, va uintptr) (r int32) {
	return Xfprintf(tls, stream, fmt, va)
}

// Xgetc is equivalent to Xfgetc.
func Xgetc(tls *TLS, stream uintptr) (r int32) {
	return Xfgetc(tls, stream)
}

// Xfgets reads in at most one less than size characters from stream and stores
// them into the buffer pointed to by s.
func Xfgets(tls *TLS, s uintptr, size int32, stream uintptr) (r uintptr) {
	return x_fgets(tls, s, size, stream)
}

// Xexit causes normal process termination and the least significant byte of
// status (i.e., status & 0xFF) is returned to the parent (see wait(2)).
func Xexit(tls *TLS, code int32) {
	tls.Close() //TODO-DBG
	x_exit(tls, code)
}

// X__builtin_exit is equivalent to Xexit.
func X__builtin_exit(tls *TLS, code int32) {
	Xexit(tls, code)
}

// Xtolower returns a lowercase equivalent of c, if c is an uppercase letter
// and a lowercase representation exists in the current locale.
func Xtolower(tls *TLS, c int32) (r int32) {
	return x_tolower(tls, c)
}

// Xrealloc function changes the size of the memory block pointed to by ptr to
// size bytes.
func Xrealloc(tls *TLS, ptr uintptr, size uint64) (r uintptr) {
	return x_realloc(tls, ptr, size)
}

// Xffs returns the position of the first (least significant) bit set in the
// word i.
func Xffs(tls *TLS, i int32) (r int32) {
	return x_ffs(tls, i)
}

func X__builtin_ffs(tls *TLS, i int32) (r int32) {
	return Xffs(tls, i)
}

// Xfabs returns the absolute value of the floating-point number x.
func Xfabs(tls *TLS, x float64) (r float64) {
	return _fabs(tls, x)
}

func X__builtin_fabs(tls *TLS, x float64) (r float64) {
	return Xfabs(tls, x)
}

// Xfabsf returns the absolute value of the floating-point number x.
func Xfabsf(tls *TLS, x float32) (r float32) {
	return _fabsf(tls, x)
}

// X__assert_fail aborts a program.
func X__assert_fail(tls *TLS, expr uintptr, file uintptr, line int32, func1 uintptr) {
	x___assert_fail(tls, expr, file, line, func1)
}

// Xatoi converts the initial portion of the string pointed to by nptr to int.
func Xatoi(tls *TLS, nptr uintptr) (r int32) {
	return x_atoi(tls, nptr)
}

// Xatol converts the initial portion of the string pointed to by nptr to long.
func Xatol(tls *TLS, nptr uintptr) (r int64) {
	return x_atol(tls, nptr)
}

// Xfputs writes the string s to stream, without its terminating null byte ('\0').
func Xfputs(tls *TLS, s uintptr, stream uintptr) (r int32) {
	return x_fputs(tls, s, stream)
}

// Xfputc writes the character c, cast to an unsigned char, to stream.
func Xfputc(tls *TLS, c int32, stream uintptr) (r int32) {
	return x_fputc(tls, c, stream)
}

// Xputc is equivalent to Xfputc.
func Xputc(tls *TLS, c int32, stream uintptr) (r int32) {
	return x_fputc(tls, c, stream)
}

// Xutchar(c) is equivalent to Xputc(c, Xstdout).
func Xputchar(tls *TLS, c int32) (r int32) {
	return x_putchar(tls, c)
}

// Xfloor returns the largest integral value that is not greater than x.
func Xfloor(tls *TLS, x float64) (r float64) {
	return x_floor(tls, x)
}

// Xfloorf returns the largest integral value that is not greater than x.
func Xfloorf(tls *TLS, x float32) (r float32) {
	return x_floorf(tls, x)
}

// Xpow returns the value of x raised to the power of y.
func Xpow(tls *TLS, x float64, y float64) (r float64) {
	return x_pow(tls, x, y)
}

// Xpowf returns the value of x raised to the power of y.
func Xpowf(tls *TLS, x float32, y float32) (r1 float32) {
	return x_powf(tls, x, y)
}

// Xcos returns the cosine of x, where x is given in radians.
func Xcos(tls *TLS, x float64) (r float64) {
	return x_cos(tls, x)
}

// Xcosf returns the cosine of x, where x is given in radians.
func Xcosf(tls *TLS, x float32) (r float32) {
	return x_cosf(tls, x)
}

// Xatan2 calculates the principal value of the arc tangent of y/x, using the
// signs of the two arguments to determine the quadrant of the result.
func Xatan2(tls *TLS, y float64, x float64) (r float64) {
	return x_atan2(tls, y, x)
}

// Xatan2f calculates the principal value of the arc tangent of y/x, using the
// signs of the two arguments to determine the quadrant of the result.
func Xatan2f(tls *TLS, y float32, x float32) (r float32) {
	return x_atan2f(tls, y, x)
}

// Xsqrt returns the nonnegative square root of x.
func Xsqrt(tls *TLS, x float64) float64 {
	return _sqrt(tls, x)
}

// Xsqrtf returns the nonnegative square root of x.
func Xsqrtf(tls *TLS, x float32) float32 {
	return _sqrtf(tls, x)
}

// Xasin calculates the principal value of the arc sine of x; that is the value
// whose sine is x.
func Xasin(tls *TLS, x float64) (r float64) {
	return x_asin(tls, x)
}

// Xasinf calculates the principal value of the arc sine of x; that is the value
// whose sine is x.
func Xasinf(tls *TLS, x float32) (r float32) {
	return x_asinf(tls, x)
}

// For output streams, Xfflush forces a write of all user-space buffered data
// for the given output or update stream via the stream's underlying write
// function.
//
// For input streams associated with seekable files (e.g., disk files, but not
// pipes or terminals), Xfflush discards any buffered data that has been
// fetched from the underlying file, but has not been consumed by the
// application.
func Xfflush(tls *TLS, stream uintptr) (r int32) {
	return x_fflush(tls, stream)
}

// Xrand returns a pseudo-random integer in the range 0 to RAND_MAX inclusive
// (i.e., the mathematical range [0, RAND_MAX]).
func Xrand(tls *TLS) (r int32) {
	return x_rand(tls)
}

// Xperror produces a message on standard error describing the last error
// encountered during a call to a system or library function.
func Xperror(tls *TLS, s uintptr) {
	x_perror(tls, s)
}

// The qsort() function sorts an array with nmemb elements of size size. The
// base argument points to the start of the array.
func Xqsort(tls *TLS, array uintptr, nmemb uint64, size uint64, compar uintptr) {
	x_qsort(tls, array, nmemb, size, compar)
}

// Xtoupper returns the uppercase equivalent of c if c is a lowercase letter
// and if an uppercase representation exists in the current locale.
func Xtoupper(tls *TLS, c int32) (r int32) {
	return x_toupper(tls, c)
}

// Xputs writes the string s and a trailing newline to stdout.
func Xputs(tls *TLS, s uintptr) (r int32) {
	return x_puts(tls, s)
}

// Xsscanf scans input according to format.
func Xsscanf(tls *TLS, str uintptr, format uintptr, va uintptr) (r int32) {
	return x_sscanf(tls, str, format, va)
}

// Xwrite() writes up to count bytes from the buffer starting at buf to the
// file referred to by the file descriptor fd.
func Xwrite(tls *TLS, fd int32, buf uintptr, count uint64) (r int64) {
	return x_write(tls, fd, buf, count)
}

// Xmemmove copies n bytes from memory area src to memory area dest.
func Xmemmove(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return _memmove(tls, dest, src, n)
}

func X__builtin_memmove(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return Xmemmove(tls, dest, src, n)
}

func _memmove(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	if n != 0 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(dest)), n), unsafe.Slice((*byte)(unsafe.Pointer(src)), n))
	}
	return dest
}

// Xmemchr scans the initial n bytes of the memory area pointed to by s for the
// first instance of c.
func Xmemchr(tls *TLS, s uintptr, c int32, n uint64) (r uintptr) {
	return x_memchr(tls, s, c, n)
}

// Xfileno examines the argument stream and returns the integer file descriptor
// used to implement this stream.
func Xfileno(tls *TLS, stream uintptr) (r int32) {
	return x_fileno(tls, stream)
}

// Xread attempts to read up to count bytes from file descriptor fd into the
// buffer starting at buf.
func Xread(tls *TLS, fd int32, buf uintptr, count uint64) (r int64) {
	return x_read(tls, fd, buf, count)
}

// Xabort aborts the program.
func Xabort(tls *TLS) {
	x_abort(tls)
}

// X__builtin_abort is equivalent to Xabort.
func X__builtin_abort(tls *TLS) {
	Xabort(tls)
}

// Xabs computes the absolute value of the integer argument j.
func Xabs(tls *TLS, j int32) (r int32) {
	return x_abs(tls, j)
}

// Xlabs computes the absolute value of the integer argument j.
func Xlabs(tls *TLS, j int64) (r int64) {
	return x_labs(tls, j)
}

// Xllabs computes the absolute value of the integer argument j.
func Xllabs(tls *TLS, j int64) (r int64) {
	return x_llabs(tls, j)
}

// X__builtin_llabs is equivalent to Xllabs.
func X__builtin_llabs(tls *TLS, j int64) (r int64) {
	return Xllabs(tls, j)
}

// X__builtin_abs is equivalent to Xabs.
func X__builtin_abs(tls *TLS, j int32) (r int32) {
	return Xabs(tls, j)
}

func X__builtin_bswap16(t *TLS, x uint16) uint16 {
	return x<<8 | x>>8
}

func X__builtin_bswap32(tls *TLS, x uint32) uint32 {
	return x>>int32(24) | x>>int32(8)&uint32(0xff00) | x<<int32(8)&uint32(0xff0000) | x<<int32(24)
}

func X__builtin_bswap64(tls *TLS, x uint64) uint64 {
	return x<<56 |
		x&0xff00<<40 |
		x&0xff0000<<24 |
		x&0xff000000<<8 |
		x&0xff00000000>>8 |
		x&0xff0000000000>>24 |
		x&0xff000000000000>>40 |
		x>>56
}

// Xalloca allocates  size bytes of space in the stack frame of the caller.
func Xalloca(tls *TLS, size uint64) uintptr {
	return tls.alloca(size)
}

func X__builtin_alloca(tls *TLS, size uint64) uintptr {
	return Xalloca(tls, size)
}

// Xcopysign returns a value whose absolute value matches that of x, but whose
// sign bit matches that of y.
func Xcopysign(tls *TLS, x float64, y float64) (r float64) {
	return x_copysign(tls, x, y)
}

// X__builtin_copysign is equivalent to Xcopysign.
func X__builtin_copysign(tls *TLS, x float64, y float64) (r float64) {
	return Xcopysign(tls, x, y)
}

// Xcopysignf returns a value whose absolute value matches that of x, but whose
// sign bit matches that of y.
func Xcopysignf(tls *TLS, x float32, y float32) (r float32) {
	return x_copysignf(tls, x, y)
}

// X__builtin_copysignf is equivalent to Xcopysignf.
func X__builtin_copysignf(tls *TLS, x float32, y float32) (r float32) {
	return Xcopysignf(tls, x, y)
}

func X__builtin_copysignl(t *TLS, x, y float64) float64 {
	return Xcopysign(t, x, y)
}

func X__builtin_expect(t *TLS, exp, c int64) int64 {
	return exp
}

func ___builtin_expect(t *TLS, exp, c int64) int64 {
	return exp
}

func X__builtin_huge_val(t *TLS) float64 {
	return math.Inf(1)
}

func X__builtin_huge_valf(t *TLS) float32 {
	return float32(math.Inf(1))
}

func X__builtin_inf(t *TLS) float64 {
	return math.Inf(1)
}

func X__builtin_inff(t *TLS) float32 {
	return ___builtin_inff(t)
}

func X__builtin_isinfl(t *TLS, x float64) int32 {
	return Bool32(math.IsInf(x, 0))
}

func X__builtin_isinff(t *TLS, x float32) int32 {
	return Bool32(math.IsInf(float64(x), 0))
}

func X__builtin_isinf(t *TLS, x float64) int32 {
	return Bool32(math.IsInf(x, 0))
}

func ___builtin_inff(t *TLS) float32 {
	return float32(math.Inf(1))
}

func X__builtin_infl(t *TLS) float64 {
	return math.Inf(1)
}

func X__builtin_isunordered(t *TLS, a, b float64) int32 {
	return Bool32(math.IsNaN(a) || math.IsNaN(b))
}

func X__builtin_nan(t *TLS, s uintptr) float64 {
	return math.NaN()
}

func X__builtin_nanf(t *TLS, s uintptr) float32 {
	return ___builtin_nanf(t, s)
}

func ___builtin_nanf(t *TLS, s uintptr) float32 {
	return float32(math.NaN())
}

func X__builtin_nanl(t *TLS, s uintptr) float64 {
	return math.NaN()
}

func X__builtin_prefetch(t *TLS, addr, args uintptr) {}

func X__builtin_unreachable(t *TLS) {
	fmt.Fprintf(os.Stderr, "unrechable\n")
	os.Stderr.Sync()
	Xexit(t, 1)
}

// Xopen opens the file specified by pathname.
func Xopen(tls *TLS, pathname uintptr, flags int32, va uintptr) (r int32) {
	return x_open(tls, pathname, flags, va)
}

// Xvfprintf provides formatted output conversion. Xvfprintf writes output to
// the given output stream.
func Xvfprintf(tls *TLS, stream uintptr, fmt uintptr, ap uintptr) (r int32) {
	return x_vfprintf(tls, stream, fmt, ap)
}

// Xvprintf provides formatted output conversion. Xvprintf writes output to
// stdout.
func Xvprintf(tls *TLS, fmt uintptr, ap uintptr) (r int32) {
	return x_vprintf(tls, fmt, ap)
}

func X__builtin_trap(t *TLS) {
	Xabort(t)
}

// Xgetrusage returns resource usage measures for who.
func Xgetrusage(tls *TLS, who int32, ru uintptr) (r int32) {
	return x_getrusage(tls, who, ru)
}

// Xstat retrieves information about the file pointed to by pathname.
func Xstat(tls *TLS, pathname uintptr, buf uintptr) (r int32) {
	return x_stat(tls, pathname, buf)
}

// Xfseek sets the file position indicator for the stream pointed to by stream.
func Xfseek(tls *TLS, stream uintptr, offset int64, whence int32) (r int32) {
	return x_fseek(tls, stream, offset, whence)
}

// Xftell function obtains the current value of the file position indicator for
// the stream pointed to by stream.
func Xftell(tls *TLS, stream uintptr) (r int64) {
	return x_ftell(tls, stream)
}

// Xrewind sets the file position indicator for the stream pointed to by stream
// to the beginning of the file.
func Xrewind(tls *TLS, stream uintptr) {
	x_rewind(tls, stream)
}

// Xlstat is identical to Xstat(), except that if pathname is a symbolic link,
// then it returns information about the link itself, not the file that the
// link refers to.
func Xlstat(tls *TLS, pathname uintptr, buf uintptr) (r int32) {
	return x_lstat(tls, pathname, buf)
}

// Xmkdir attempts to create a directory named pathname.
func Xmkdir(tls *TLS, pathname uintptr, mode uint32) (r int32) {
	return x_mkdir(tls, pathname, mode)
}

// Xsymlink creates a symbolic link named linkpath which contains the string
// target.
func Xsymlink(tls *TLS, target uintptr, linkpath uintptr) (r int32) {
	return x_symlink(tls, target, linkpath)
}

// Xchmod changes the mode of the file specified whose pathname is given in
// pathname, which is dereferenced if it is a symbolic link.
func Xchmod(tls *TLS, pathname uintptr, mode uint32) (r int32) {
	return x_chmod(tls, pathname, mode)
}

// Xime() returns the time as the number of seconds since the Epoch, 1970-01-01
// 00:00:00 +0000 (UTC).
func Xtime(tls *TLS, tloc uintptr) (r int64) {
	return x_time(tls, tloc)
}

// Xutime changes the access and modification times of the inode specified by
// filename to the actime and modtime fields of times respectively.
func Xutime(tls *TLS, filename uintptr, times uintptr) (r int32) {
	return x_utime(tls, filename, times)
}

// Xutimes is similar to Xutime, but the times argument refers to an array
// rather than a structure.
func Xutimes(tls *TLS, filename uintptr, times uintptr) (r int32) {
	return x_utimes(tls, filename, times)
}

// Xclosedir function closes the directory stream associated with dirp.
func Xclosedir(tls *TLS, dirp uintptr) (r int32) {
	return x_closedir(tls, dirp)
}

// Xopendir opens a directory stream corresponding to the directory name, and
// returns a pointer to the directory stream.
func Xopendir(tls *TLS, name uintptr) (r uintptr) {
	return x_opendir(tls, name)
}

// Xreaddir returns a pointer to a dirent structure representing the next
// directory entry in the directory stream pointed to by dirp.
func Xreaddir(tls *TLS, dirp uintptr) (r uintptr) {
	return x_readdir(tls, dirp)
}

// Xreadlink places the contents of the symbolic link pathname in the buffer
// buf, which has size bufsiz. readlink() does not append a null byte to buf.
func Xreadlink(tls *TLS, pathname uintptr, buf uintptr, bufsiz uint64) (r int64) {
	return x_readlink(tls, pathname, buf, bufsiz)
}

// Xstrstr function finds the first occurrence of the substring needle in the
// string haystack.
func Xstrstr(tls *TLS, haystack uintptr, needle uintptr) (r uintptr) {
	return x_strstr(tls, haystack, needle)
}

// Xgetenv function searches the environment list to find the environment
// variable name, and returns a pointer to the corresponding value string.
func Xgetenv(tls *TLS, name uintptr) (r uintptr) {
	r = x_getenv(tls, name)
	return r
}

// Xsystem executes the shell command specified in command in a child process.
func Xsystem(t *TLS, command uintptr) int32 {
	s := GoString(command)
	if command == 0 {
		_, err := exec.LookPath("sh")
		return Bool32(err == nil)
	}

	cmd := exec.Command("sh", "-c", s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		ps := err.(*exec.ExitError)
		return int32(ps.ExitCode())
	}

	return 0
}

// Xunlink deletes a name from the filesystem.
func Xunlink(tls *TLS, name uintptr) (r int32) {
	return x_unlink(tls, name)
}

// Xgetpid returns the process ID (PID) of the calling process.
func Xgetpid(tls *TLS) (r int32) {
	return x_getpid(tls)
}

// Xaccess checks whether the calling process can access the file pathname.
func Xaccess(tls *TLS, pathname uintptr, mode int32) (r int32) {
	return x_access(tls, pathname, mode)
}

// Xpclose waits for the associated process to terminate and returns the exit
// status of the command as returned by wait4(2).
func Xpclose(tls *TLS, stream uintptr) (r int32) {
	return x_pclose(tls, stream)
}

// Xchdir changes the current working directory of the calling process to the
// directory specified in path.
func Xchdir(tls *TLS, path uintptr) (r int32) {
	return x_chdir(tls, path)
}

// Xstrdup returns a pointer to a new string which is a duplicate of the string
// s.
func Xstrdup(tls *TLS, s uintptr) (r uintptr) {
	return x_strdup(tls, s)
}

// Xpopen opens a process by creating a pipe, forking, and invoking the shell.
func Xpopen(tls *TLS, cmd uintptr, mode uintptr) (r uintptr) {
	return x_popen(tls, cmd, mode)
}

// Xstrtol converts the initial part of the string in nptr to a long integer
// value according to the given base, which must be between 2 and 36 inclusive,
// or be the special value 0.
func Xstrtol(tls *TLS, nptr uintptr, endptr uintptr, base int32) (r int64) {
	return x_strtol(tls, nptr, endptr, base)
}

// Xgetuid returns the real user ID of the calling process.
func Xgetuid(tls *TLS) (r uint32) {
	return x_getuid(tls)
}

// Xgetpwuid returns a pointer to a structure containing the broken-out fields
// of the record in the password database that matches the user ID uid.
func Xgetpwuid(tls *TLS, uid uint32) (r uintptr) {
	return x_getpwuid(tls, uid)
}

// Xsetvbuf adjusts file buffering.
func Xsetvbuf(tls *TLS, stream uintptr, buf uintptr, mode int32, size uint64) (r int32) {
	return x_setvbuf(tls, stream, buf, mode, size)
}

// Xisatty function tests whether fd is an open file descriptor referring to a
// terminal.
func Xisatty(tls *TLS, fd int32) (r int32) {
	return x_isatty(tls, fd)
}

// Xisprint checks for any printable character including space.
func Xisprint(tls *TLS, c int32) (r int32) {
	return x_isprint(tls, c)
}

func X__builtin_isprint(tls *TLS, c int32) (r int32) {
	return Xisprint(tls, c)
}

// Xisalnum checks for an alphanumeric character; it is equivalent to
// (Xisalpha(c) || Xisdigit(c)).
func Xisalnum(tls *TLS, c int32) (r int32) {
	return x_isalnum(tls, c)
}

// X__builtin_isalnum is equivalent to Xisalnum.
func X__builtin_isalnum(tls *TLS, c int32) (r int32) {
	return Xisalnum(tls, c)
}

// Xisblank checks for a blank character; that is, a space or a tab.
func Xisblank(tls *TLS, c int32) (r int32) {
	return x_isblank(tls, c)
}

// X__builtin_isblank is equivalent to Xisblank.
func X__builtin_isblank(tls *TLS, c int32) (r int32) {
	return Xisblank(tls, c)
}

// The atexit() function registers the given function to be called at normal
// process termination, either via exit(3) or via return from the program's
// main().
func Xatexit(tls *TLS, function uintptr) (r int32) {
	return x_atexit(tls, function)
}

// Xraise function sends a signal to the calling process or thread.
func Xraise(tls *TLS, sig int32) (r int32) {
	return x_raise(tls, sig)
}

func X__ctype_b_loc(tls *TLS) uintptr {
	panic(todo("")) // Not provided in musl v0.6.0
}

// Xbacktrace returns a backtrace for the calling program, in the array pointed
// to by buffer.
func Xbacktrace(tls *TLS, buffer uintptr, size int32) int32 {
	panic(todo("")) // Not provided in musl v0.6.0
}

// Xbacktrace_symbols_fd takes the same buffer and size arguments as
// backtrace_symbols(), but instead of returning an array of strings to the
// caller, it writes the strings, one per line, to the file descriptor fd.
func Xbacktrace_symbols_fd(tls *TLS, buffer uintptr, size, fd int32) {
	panic(todo("")) // Not provided in musl v0.6.0
}

func X__builtin_clz(t *TLS, n uint32) int32 {
	return int32(bits.LeadingZeros32(n))
}

func X__builtin_clzll(t *TLS, n uint64) int32 {
	return int32(bits.LeadingZeros64(n))
}

// Xclose closes a file descriptor, so that it no longer refers to any file and
// may be reused.
func Xclose(tls *TLS, fd int32) (r int32) {
	return x_close(tls, fd)
}

// Xgetcwd copies an absolute pathname of the current working directory to the
// array pointed to by buf, which is of length size.
func Xgetcwd(tls *TLS, buf uintptr, size uint64) (r uintptr) {
	return x_getcwd(tls, buf, size)
}

// Xfstat is identical to Xstat, except that the file about which information
// is to be retrieved is specified by the file descriptor fd.
func Xfstat(tls *TLS, fd int32, buf uintptr) (r int32) {
	return x___fstat(tls, fd, buf)
}

// Xftruncate causes the regular file referenced by fd to be truncated to a
// size of precisely length bytes.
func Xftruncate(tls *TLS, fd int32, length int64) (r int32) {
	return x_ftruncate(tls, fd, length)
}

// Xfcntl performs an operation on the open file descriptor fd.
func Xfcntl(tls *TLS, fd int32, cmd int32, va uintptr) (r1 int32) {
	return x_fcntl(tls, fd, cmd, va)
}

// Xpread reads up to count bytes from file descriptor fd at offset offset
// (from the start of the file) into the buffer starting at buf.
func Xpread(tls *TLS, fd int32, buf uintptr, count uint64, offset int64) (r int64) {
	return x_pread(tls, fd, buf, count, offset)
}

// Xpwrite writes up to count bytes from the buffer starting at buf to the file
// descriptor fd at offset offset.
func Xpwrite(tls *TLS, fd int32, buf uintptr, count uint64, offset int64) (r int64) {
	return x_pwrite(tls, fd, buf, count, offset)
}

// Xfchmod changes file mode bits.
func Xfchmod(tls *TLS, fd int32, mode uint32) (r int32) {
	return x_fchmod(tls, fd, mode)
}

// Xrmdir deletes a directory, which must be empty.
func Xrmdir(tls *TLS, pathname uintptr) (r int32) {
	return x_rmdir(tls, pathname)
}

// Xfchown changes the ownership of the file referred to by the open file
// descriptor fd.
func Xfchown(tls *TLS, fd int32, uid uint32, gid uint32) (r int32) {
	return x_fchown(tls, fd, uid, gid)
}

// Xgeteuid returns the effective user ID of the calling process.
func Xgeteuid(tls *TLS) (r uint32) {
	return x_geteuid(tls)
}

// Xmmap creates a new mapping in the virtual address space of the calling
// process.
func Xmmap(tls *TLS, start uintptr, len uint64, prot int32, flags int32, fd int32, off int64) (r uintptr) {
	return x___mmap(tls, start, len, prot, flags, fd, off)
}

// Xmunmap deletes the mappings for the specified address range, and causes
// further references to addresses within the range to generate invalid memory
// references.
func Xmunmap(tls *TLS, start uintptr, len uint64) (r int32) {
	return x___munmap(tls, start, len)
}

// Xmremap expands (or shrinks) an existing memory mapping, potentially moving
// it at the same time
func Xmremap(tls *TLS, old_addr uintptr, old_len uint64, new_len uint64, flags int32, va uintptr) (r uintptr) {
	return x___mremap(tls, old_addr, old_len, new_len, flags, va)
}

// Xstrerror function returns a pointer to a string that describes the error
// code passed in the argument errnum,
func Xstrerror(tls *TLS, errnum int32) (r uintptr) {
	return x_strerror(tls, errnum)
}

// Xfsync transfers ("flushes") all modified in-core data of (i.e., modified
// buffer cache pages for) the file referred to by the file descriptor fd to
// the disk device (or other permanent storage device) so that all changed
// information can be retrieved even if the system crashes or is rebooted.
func Xfsync(tls *TLS, fd int32) (r int32) {
	return x_fsync(tls, fd) //TODO nop in musl v0.6.0.
}

// Xsysconf gets configuration information at run time
func Xsysconf(tls *TLS, name int32) (r int64) {
	return x_sysconf(tls, name)
}

// Xdlopen() loads the dynamic shared object (shared library) file named by the
// null-terminated string filename and returns an opaque "handle" for the
// loaded object.
func Xdlopen(tls *TLS, filename uintptr, flags int32) uintptr {
	panic(todo(""))
}

// Xdlerror function returns a human-readable, null-terminated string
// describing the most recent error that occurred from a call to one of the
// functions in the dlopen API since the last call to dlerror().
func Xdlerror(tls *TLS) uintptr {
	panic(todo(""))
}

// Xdlsym takes a "handle" of a dynamic loaded shared object returned by
// dlopen(3) along with a null-terminated symbol name, and returns the address
// where that symbol is loaded into memory.
func Xdlsym(tls *TLS, handle, symbol uintptr) uintptr {
	panic(todo(""))
}

// Xdlclose decrements the reference count on the dynamically loaded shared
// object referred to by handle.
func Xdlclose(tls *TLS, handle uintptr) int32 {
	panic(todo(""))
}

// Xnanosleep suspends the execution of the calling thread until either at
// least the time specified in *req has elapsed, or the delivery of a signal
// that triggers the invocation of a handler in the calling thread or that
// terminates the process.
func Xnanosleep(tls *TLS, req uintptr, rem uintptr) (r int32) {
	return x_nanosleep(tls, req, rem)
}

// Xgettimeofday gets the time as well as a timezone.
func Xgettimeofday(tls *TLS, tv uintptr, tz uintptr) (r int32) {
	return x_gettimeofday(tls, tv, tz)
}

// Xstrcspn calculates the length of the initial segment of s which consists
// entirely of bytes not in reject.
func Xstrcspn(tls *TLS, s uintptr, reject uintptr) (r uint64) {
	return x_strcspn(tls, s, reject)
}

// Xlocaltime function converts the calendar time timep to broken-down time
// representation, expressed relative to the user's specified timezone.
func Xlocaltime(tls *TLS, timep uintptr) (r uintptr) {
	return x_localtime(tls, timep)
}

// Xgetopt parses the command-line arguments.
func Xgetopt(tls *TLS, argc int32, argv uintptr, optstring uintptr) (r int32) {
	return x_getopt(tls, argc, argv, optstring)
}

// Xungetc pushes c back to stream, cast to unsigned char, where it is
// available for subsequent read operations
func Xungetc(tls *TLS, c int32, stream uintptr) (r int32) {
	return x_ungetc(tls, c, stream)
}

// Xfscanf scans input from stream according to format
func Xfscanf(tls *TLS, stream uintptr, format uintptr, va uintptr) (r int32) {
	return x_fscanf(tls, stream, format, va)
}

// Xisspace checks for white-space characters.
func Xisspace(tls *TLS, c int32) (r int32) {
	return x_isspace(tls, c)
}

// Xsleep causes the calling thread to sleep either the number of real-time
// seconds specified in 'seconds' have elapsed or until a signal arrives which
// is not ignored.
func Xsleep(tls *TLS, seconds uint32) (r uint32) {
	return x_sleep(tls, seconds)
}

// Xusleep suspends execution of the calling thread for (at least) usec
// microseconds.
func Xusleep(tls *TLS, usec uint32) (r int32) {
	return x_usleep(tls, usec)
}

// Xgetentropy function writes length bytes of high-quality random data to the
// buffer starting at the location pointed to by buffer.
func Xgetentropy(tls *TLS, buffer uintptr, length uint64) (r int32) {
	return x_getentropy(tls, buffer, length)
}

// Xreallocarray() function changes the size of the memory block pointed to by
// ptr to be large enough for an array of nmemb elements, each of which is size
// bytes.
func Xreallocarray(tls *TLS, ptr uintptr, nmemb uint64, size uint64) (r uintptr) {
	return x_reallocarray(tls, ptr, nmemb, size)
}

// Xtmpnam returns a pointer to a string that is a valid filename, and such
// that a file with this name did not exist at some point in time, so that
// naive programmers may think it a suitable name for a temporary file.
func Xtmpnam(tls *TLS, buf uintptr) (r uintptr) {
	return x_tmpnam(tls, buf)
}

// Xfreopen opens the file whose name is the string pointed to by pathname and
// associates the stream pointed to by stream with it.
func Xfreopen(tls *TLS, pathname uintptr, mode uintptr, stream uintptr) (r uintptr) {
	return x_freopen(tls, pathname, mode, stream)
}

// Xmempcpy copies n bytes from the object beginning at src into the object
// pointed to by dest.
func Xmempcpy(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return x_mempcpy(tls, dest, src, n)
}

// Xremove() deletes a name from the filesystem.
func Xremove(tls *TLS, name uintptr) (r int32) {
	return x_remove(tls, name)
}

// Xstrspn calculates the length (in bytes) of the initial segment of s which
// consists entirely of bytes in accept.
func Xstrspn(tls *TLS, s uintptr, accept uintptr) (r uint64) {
	return x_strspn(tls, s, accept)
}

// Xlseek repositions the file offset of the open file description associated
// with the file descriptor fd to the argument offset according to the
// directive whence.
func Xlseek(tls *TLS, fd int32, offset int64, whence int32) (r int64) {
	return x___lseek(tls, fd, offset, whence)
}

// Xferror tests the error indicator for the stream pointed to by stream,
// returning nonzero if it is set.
func Xferror(tls *TLS, stream uintptr) (r int32) {
	return x_ferror(tls, stream)
}

func X__builtin_add_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.AddOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func X__builtin_sub_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.SubOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func X__builtin_mul_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	r, ovf := mathutil.MulOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func _fabs(tls *TLS, x float64) float64 {
	return math.Abs(x)
}

func _fabsf(tls *TLS, x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func _fabsl(tls *TLS, x float64) float64 {
	return math.Abs(x)
}

func ___DOUBLE_BITS(tls *TLS, f float64) uint64 {
	return math.Float64bits(f)
}

func ___FLOAT_BITS(tls *TLS, f float32) uint32 {
	return math.Float32bits(f)
}

func ___fesetround(tls *TLS, r int32) (r1 int32) {
	panic(todo(""))
}

func _fegetround(tls *TLS) int32 {
	panic(todo(""))
}

func _fegetenv(tls *TLS, _ uintptr) int32 {
	panic(todo(""))
}

func _feclearexcept(tls *TLS, _ int32) int32 {
	panic(todo(""))
}

func _feraiseexcept(tls *TLS, _ int32) int32 {
	panic(todo(""))
}

func _fetestexcept(tls *TLS, _ int32) int32 {
	panic(todo(""))
}

func _fesetenv(tls *TLS, _ uintptr) int32 {
	panic(todo(""))
}

// Xsetjmp saves various information about the calling environment (typically,
// the stack pointer, the instruction pointer, possibly the values of other
// registers and the  signal mask) in the buffer env for later use by
// longjmp().
func setjmp(tls *TLS, env uintptr) int32 {
	panic(todo(""))
}

func _setjmp(tls *TLS, jmp_buf uintptr) int32 {
	panic(todo(""))
}

func X__builtin_setjmp(tls *TLS, jmp_buf uintptr) int32 {
	return _setjmp(tls, jmp_buf)
}

func ___syscall_cp_asm(tls *TLS, _ uintptr, a, b, c, d, e, f, g int64) int64 {
	panic(todo(""))
}

func ___get_tp(tls *TLS) uintptr {
	return tls.fs
}

// // func ___pthread_self(tls *TLS) uintptr {
// // 	return uintptr(unsafe.Pointer(tls.fs))
// // }

// This is not the set_thread_area syscall, but arch_prctl syscall with
// subcommand ARCH_SET_FS.
//
//	/* Copyright 2011 Nicholas J. Kain, licensed GNU LGPL 2.1 or later */
//	.text
//	.global __set_thread_area
//	.type __set_thread_area,@function
//	__set_thread_area:
//	        mov %rdi,%rsi           /* shift for syscall */
//	        movl $0x1002,%edi       /* SET_FS register */
//	        movl $158,%eax          /* set fs segment to */
//	        syscall                 /* arch_prctl(SET_FS, arg)*/
//		ret
func ___set_thread_area(tls *TLS, p uintptr) int32 {
	tls.fs = p
	return 0
}

func X__errno_location(tls *TLS) uintptr {
	return x___errno_location(tls)
}

// Xpthread_create function  starts a new simulated thread in the calling
// process. C threads are simulated by Go goroutines.
func Xpthread_create(tls *TLS, res uintptr, attr uintptr, entry uintptr, arg uintptr) (r int32) {
	panic(todo(""))
}

// // 	if tls.fs == 0 {
// // 		return m_ENOSYS
// // 	}
// //
// // 	var v2 int32
// // 	var v3 bool
// // 	if v3 = !(_init1 != 0); v3 {
// // 		_init1++
// // 		v2 = _init1
// // 	}
// // 	if v3 && v2 != 0 {
// // 		_init_threads(tls)
// // 	}
// // 	if attr == 0 {
// // 		attr = uintptr(unsafe.Pointer(&_default_attr))
// // 	}
// //
// // 	self := tls.fs
// // 	newTLS := NewTLS()
// // 	new1 := uintptr(unsafe.Pointer(newTLS.fs))
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fpid = (*T__pthread)(unsafe.Pointer(self)).Fpid
// // 	(*T__pthread)(unsafe.Pointer(new1)).Ferrno_ptr = new1 + 52
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fstart = entry
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fstart_arg = arg
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fself = uintptr(unsafe.Pointer(newTLS))
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fdetached = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&(*Tpthread_attr_t)(unsafe.Pointer(attr)).F__u.F__i)) + uintptr(Uint64FromInt32(2)*(Uint64FromInt64(8)/Uint64FromInt64(4))+Uint64FromInt32(0))*4))
// // 	(*T__pthread)(unsafe.Pointer(new1)).Fattr = *(*Tpthread_attr_t)(unsafe.Pointer(attr))
// // 	_a_inc(tls, uintptr(unsafe.Pointer(&x___libc.Fthreads_minus_1)))
// // 	newTLS.threadFuncWait = make(chan struct{})
// //
// // 	go func() {
// //
// // 		defer close(newTLS.threadFuncWait)
// //
// // 		(*T__pthread)(unsafe.Pointer(new1)).Fresult = (*(*func(*TLS, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{entry})))(newTLS, arg)
// // 	}()
// //
// // 	*(*uintptr)(unsafe.Pointer(res)) = new1
// // 	return 0
// // }

// Xpthread_join function waits for the simulated thread specified by 'thread'
// to terminate.
func Xpthread_join(tls *TLS, thread uintptr, res uintptr) (r int32) {
	panic(todo(""))
}

// // 	defer (*TLS)(unsafe.Pointer((*T__pthread)(unsafe.Pointer(thread)).Fself)).Close()
// //
// // 	<-(*TLS)(unsafe.Pointer((*T__pthread)(unsafe.Pointer(thread)).Fself)).threadFuncWait
// // 	if res != 0 {
// // 		*(*uintptr)(unsafe.Pointer(res)) = (*T__pthread)(unsafe.Pointer(thread)).Fresult
// // 	}
// // 	return 0
// // }

// Xpthread_mutex_lock locks the given mutex.
func Xpthread_mutex_lock(tls *TLS, mutex uintptr) (r int32) {
	panic(todo(""))
	// return x_pthread_mutex_lock(tls, mutex)
}

// Xpthread_mutex_unlock unlocks the given mutex.
func Xpthread_mutex_unlock(tls *TLS, mutex uintptr) (r int32) {
	panic(todo(""))
	// return x_pthread_mutex_unlock(tls, mutex)
}

// Xpthread_self function returns the ID of the calling thread.
func Xpthread_self(tls *TLS) (r uintptr) {
	panic(todo(""))
	// return x_pthread_self(tls)
}

// Xpthread_cond_wait atomically unlocks the mutex (as per
// pthread_unlock_mutex) and waits for the condition variable cond to be
// signaled.
func Xpthread_cond_wait(tls *TLS, cond uintptr, mutex uintptr) (r int32) {
	panic(todo(""))
	// return x_pthread_cond_wait(tls, cond, mutex)
}

// Xpthread_cond_signal restarts one of the threads that are waiting on the
// condition variable cond.
func Xpthread_cond_signal(tls *TLS, cond uintptr) (r int32) {
	panic(todo(""))
	// return x_pthread_cond_signal(tls, cond)
}
