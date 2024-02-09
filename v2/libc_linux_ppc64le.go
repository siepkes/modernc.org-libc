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
)

const (
	heapAlign = 16
	heapGuard = 16
)

type long = int64

var (
	_ error = (*MemAuditError)(nil)

	argc         int32   // Value established in initLibc
	argv         uintptr // Value established in initLibc
	initLibcOnce sync.Once
	tid          atomic.Int32 // TLS Go ID
)

// Start executes C's main.
func Start(main func(*TLS, int32, uintptr) int32) {
	runtime.LockOSThread()

	if isMemBrk {
		defer func() {
			trc("==== PANIC")
			for _, v := range MemAudit() {
				trc("", v.Error())
			}
		}()
	}

	tls := NewTLS()
	rc := main(tls, argc, argv)
	if isMemBrk {
		trc("==== RC %v", rc)
		for _, v := range MemAudit() {
			trc("", v.Error())
		}
	}
	Xexit(tls, rc)
}

func initLibc(tls *TLS) {
	var args []uintptr
	for _, v := range os.Args {
		p := mustCalloc(Tsize_t(len(v) + 1))
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		args = append(args, p)
	}
	args = append(args, 0)
	for _, v := range os.Environ() {
		p := mustCalloc(Tsize_t(len(v) + 1))
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		args = append(args, p)
	}
	args = append(args, 0)
	argv = mustCalloc(Tsize_t(len(args) * int(unsafe.Sizeof(uintptr(0)))))
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

type MemAuditError struct {
	Caller  string
	Message string
}

func (e *MemAuditError) Error() string {
	return fmt.Sprintf("%s: %s", e.Caller, e.Message)
}

// CString returns a pointer to a zero-terminated version of s. The caller is
// responsible for freeing the allocated memory using Xfree.
func CString(tls *TLS, s string) (uintptr, error) {
	n := len(s)
	p := Xmalloc(tls, Tsize_t(n)+1)
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

func mustMalloc(sz Tsize_t) (r uintptr) {
	if r = Xmalloc(nil, sz); r != 0 || sz == 0 {
		return r
	}

	panic(todo("OOM"))
}

func mustCalloc(sz Tsize_t) (r uintptr) {
	if r := Xcalloc(nil, 1, sz); r != 0 || sz == 0 {
		return r
	}

	panic(todo("OOM"))
}

type tlsStackSlot struct {
	p  uintptr
	sz Tsize_t
}

// TLS emulates thread local storage. TLS is not safe for concurrent use by
// multiple goroutines.
type TLS struct {
	allocas        []uintptr
	freeFS         uintptr
	fs             uintptr // *T__pthread
	sp             int
	stack          []tlsStackSlot
	threadFuncWait chan struct{}

	ID int32
}

// NewTLS returns a newly created TLS that must be eventually closed to prevent
// resource leaks.
func NewTLS() (r *TLS) {
	r = &TLS{
		ID: tid.Add(1),
		fs: mustMalloc(Tsize_t(unsafe.Sizeof(t__pthread{}))),
	}
	r.freeFS = r.fs
	*(*t__pthread)(unsafe.Pointer(r.fs)) = t__pthread{
		Ftid: r.ID,
	}
	initLibcOnce.Do(func() { initLibc(r) })
	return r
}

func (t *TLS) String() string {
	return fmt.Sprintf("{%p: ID=%v fs=%#0x}", t, t.ID, t.fs)
}

// Alloc allocates n bytes in t's local storage. Calls to Alloc() must be
// strictly paired with calls to TLS.Free on function exit. That also means any
// memory from Alloc must not be used after a function returns.
//
// The order matters. This is ok:
//
//	p := t.Alloc(11)
//		q := t.Alloc(22)
//		t.Free(22)
//		// q is no more usable here.
//	t.Free(11)
//	// p is no more usable here.
//
// This is not correct:
//
//	t.Alloc(11)
//		t.Alloc(22)
//		t.Free(11)
//	t.Free(22)
func (t *TLS) Alloc(n Tsize_t) (r uintptr) {
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
	if t.sp < len(t.stack) {
		p := t.stack[t.sp].p
		sz := t.stack[t.sp].sz
		if sz >= n && sz <= shrinkSegment*n {
			t.sp++
			return p
		}

		Xfree(t, p)
		r = mustMalloc(n)
		t.stack[t.sp] = tlsStackSlot{p: r, sz: Xmalloc_usable_size(t, r)}
		t.sp++
		return r

	}

	r = mustMalloc(n)
	t.stack = append(t.stack, tlsStackSlot{p: r, sz: Xmalloc_usable_size(t, r)})
	t.sp++
	return r
}

// Free manages memory of the preceding TLS.Alloc()
func (t *TLS) Free(n int) {
	//TODO shrink stacks
	t.sp--
}

func (t *TLS) alloca(n Tsize_t) (r uintptr) {
	r = mustMalloc(n)
	t.allocas = append(t.allocas, r)
	return r
}

func (t *TLS) FreeAlloca() {
	for _, v := range t.allocas {
		Xfree(t, v)
	}
	t.allocas = t.allocas[:0]
}

func (t *TLS) Close() {
	defer func() { *t = TLS{} }()

	for _, v := range t.allocas {
		Xfree(t, v)
	}
	for _, v := range t.stack[:t.sp] {
		Xfree(t, v.p)
	}
	Xfree(t, t.freeFS)
}

func (tls *TLS) setErrno(err int32) {
	if tls != nil {
		*(*int32)(unsafe.Pointer(X__errno_location(tls))) = err
	}
}

// musl-1.2.4/src/env/__init_tls.c:81:extern weak hidden const Tsize_t _DYNAMIC[];
var __DYNAMIC [1]uint64

var ___cp_begin [1]int8

var ___cp_cancel [1]int8

var ___cp_end [1]int8

func ___clone(tls *TLS, func1 uintptr, stack uintptr, flags int32, arg uintptr, va uintptr) (r int32) {
	panic(todo(""))
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
	if __ccgo_strace {
		trc("tls=%v env=%v val=%v, (%v:)", tls, env, val, origin(2))
	}
	panic(todo(""))
}

func X__builtin_longjmp(tls *TLS, env uintptr, val int32) int64 {
	if __ccgo_strace {
		trc("tls=%v env=%v val=%v, (%v:)", tls, env, val, origin(2))
	}
	panic(todo(""))
}

func X__builtin_iswblank(tls *TLS, wc uint32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v wc=%v, (%v:)", tls, wc, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xiswblank(tls, wc)
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

// int16 stores are atomic on this CPU when properly aligned.
func a_store_16(addr uintptr, val uint16) {
	if addr&1 != 0 {
		panic(fmt.Errorf("unaligned atomic 16 bit access at %#0x", addr))
	}

	*(*uint16)(unsafe.Pointer(addr)) = val
}

func a_crash()

func _a_crash(tls *TLS) {
	a_crash()
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

// // func a_and_64(p uintptr, x uint64)
// //
// // func _a_and_64(tls *TLS, p uintptr, x uint64) {
// // 	a_and_64(p, x)
// // }

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
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
	}
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

func ___syscall0(tls *TLS, n long) long {
	switch n {
	case SYS_sched_yield:
		runtime.Gosched()
		return 0
	default:
		r1, _, err := syscall.Syscall(uintptr(n), 0, 0, 0)
		if err != 0 {
			return long(-err)
		}

		return long(r1)
	}
}

func ___syscall1(tls *TLS, n, a1 long) long {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), 0, 0)
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___syscall2(tls *TLS, n, a1, a2 long) long {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), 0)
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___syscall3(tls *TLS, n, a1, a2, a3 long) long {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3))
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___syscall4(tls *TLS, n, a1, a2, a3, a4 long) long {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), 0, 0)
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___syscall5(tls *TLS, n, a1, a2, a3, a4, a5 long) long {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), 0)
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___syscall6(tls *TLS, n, a1, a2, a3, a4, a5, a6 long) long {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), uintptr(a6))
	if err != 0 {
		return long(-err)
	}

	return long(r1)
}

func ___unmapself(tls *TLS, base uintptr, size Tsize_t) {
	panic(todo(""))
}

func X__builtin_vsnprintf(tls *TLS, str uintptr, n Tsize_t, format uintptr, ap uintptr) (r int32) {
	if __ccgo_strace {
		trc("tls=%v str=%v n=%v format=%v ap=%v, (%v:)", tls, str, n, format, ap, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xvsnprintf(tls, str, n, format, ap)
}

func X__builtin_sprintf(tls *TLS, str uintptr, fmt uintptr, va uintptr) (r int32) {
	if __ccgo_strace {
		trc("tls=%v str=%v fmt=%v va=%v, (%v:)", tls, str, fmt, va, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xsprintf(tls, str, fmt, va)
}

func X__builtin_rintf(tls *TLS, x float32) (r float32) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xrintf(tls, x)
}

func X__builtin_memset(tls *TLS, s uintptr, c int32, n Tsize_t) uintptr {
	if __ccgo_strace {
		trc("tls=%v s=%v c=%v n=%v, (%v:)", tls, s, c, n, origin(2))
	}
	return Xmemset(tls, s, c, n)
}

func X__builtin_snprintf(tls *TLS, str uintptr, size Tsize_t, format, args uintptr) int32 {
	if __ccgo_strace {
		trc("tls=%v str=%v size=%v args=%v, (%v:)", tls, str, size, args, origin(2))
	}
	return Xsnprintf(tls, str, size, format, args)
}

func X__builtin_malloc(tls *TLS, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v n=%v, (%v:)", tls, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xmalloc(tls, n)
}

func X__builtin_free(tls *TLS, p uintptr) {
	if __ccgo_strace {
		trc("tls=%v p=%v, (%v:)", tls, p, origin(2))
	}
	Xfree(tls, p)
}

func X__builtin_printf(tls *TLS, format, va uintptr) (r int32) {
	if __ccgo_strace {
		trc("tls=%v va=%v, (%v:)", tls, va, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xprintf(tls, format, va)
}

func X__builtin_round(tls *TLS, x float64) (r float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xround(tls, x)
}

func X__builtin_strcpy(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v dest=%v src=%v, (%v:)", tls, dest, src, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xstrcpy(tls, dest, src)
}

func X__builtin_strncpy(tls *TLS, dest uintptr, src uintptr, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v dest=%v src=%v n=%v, (%v:)", tls, dest, src, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xstrncpy(tls, dest, src, n)
}

func X__builtin_strcmp(tls *TLS, s1 uintptr, s2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("tls=%v s1=%v s2=%v, (%v:)", tls, s1, s2, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xstrcmp(tls, s1, s2)
}

func X__builtin_strlen(tls *TLS, s uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("tls=%v s=%v, (%v:)", tls, s, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xstrlen(tls, s)
}

func X__builtin_memcpy(tls *TLS, dest, src uintptr, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v src=%v n=%v, (%v:)", tls, src, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xmemcpy(tls, dest, src, n)
}

func X__builtin_memcmp(tls *TLS, s1 uintptr, s2 uintptr, n Tsize_t) (r1 int32) {
	if __ccgo_strace {
		trc("tls=%v s1=%v s2=%v n=%v, (%v:)", tls, s1, s2, n, origin(2))
		defer func() { trc("-> %v", r1) }()
	}
	return Xmemcmp(tls, s1, s2, n)
}

func X__builtin_fprintf(tls *TLS, stream uintptr, fmt uintptr, va uintptr) (r int32) {
	if __ccgo_strace {
		trc("tls=%v stream=%v fmt=%v va=%v, (%v:)", tls, stream, fmt, va, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xfprintf(tls, stream, fmt, va)
}

func X__builtin_exit(tls *TLS, code int32) {
	if __ccgo_strace {
		trc("tls=%v code=%v, (%v:)", tls, code, origin(2))
	}
	Xexit(tls, code)
}

func X__builtin_ffs(tls *TLS, i int32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v i=%v, (%v:)", tls, i, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xffs(tls, i)
}

func Xfabs(tls *TLS, x float64) (r float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return _fabs(tls, x)
}

func X__builtin_fabs(tls *TLS, x float64) (r float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xfabs(tls, x)
}

func Xfabsf(tls *TLS, x float32) (r float32) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return _fabsf(tls, x)
}

func X__assert_fail(tls *TLS, expr uintptr, file uintptr, line int32, func1 uintptr) {
	if __ccgo_strace {
		trc("tls=%v expr=%v file=%v line=%v func1=%v, (%v:)", tls, expr, file, line, func1, origin(2))
	}
	x___assert_fail(tls, expr, file, line, func1)
}

func Xsqrt(tls *TLS, x float64) float64 {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
	}
	return _sqrt(tls, x)
}

func Xsqrtf(tls *TLS, x float32) float32 {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
	}
	return _sqrtf(tls, x)
}

func X__builtin_memmove(tls *TLS, dest uintptr, src uintptr, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v dest=%v src=%v n=%v, (%v:)", tls, dest, src, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xmemmove(tls, dest, src, n)
}

func Xpread64(tls *TLS, fd int32, buf uintptr, size Tsize_t, ofs int64) (r Tssize_t) {
	if __ccgo_strace {
		trc("tls=%v fd=%v buf=%v size=%v ofs=%v, (%v:)", tls, fd, buf, size, ofs, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xpread(tls, fd, buf, size, ofs)
}

func Xpwrite64(tls *TLS, fd int32, buf uintptr, size Tsize_t, ofs int64) (r Tssize_t) {
	if __ccgo_strace {
		trc("tls=%v fd=%v buf=%v size=%v ofs=%v, (%v:)", tls, fd, buf, size, ofs, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xpwrite(tls, fd, buf, size, ofs)
}

func Xstrchrnul(tls *TLS, s uintptr, c int32) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v s=%v c=%v, (%v:)", tls, s, c, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return x___strchrnul(tls, s, c)
}

func X__builtin_log2(tls *TLS, x float64) (r1 float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
		defer func() { trc("-> %v", r1) }()
	}
	return Xlog2(tls, x)
}

func X__builtin_abort(tls *TLS) {
	if __ccgo_strace {
		trc("tls=%v, (%v:)", tls, origin(2))
	}
	Xabort(tls)
}

func X__builtin_llabs(tls *TLS, j int64) (r int64) {
	if __ccgo_strace {
		trc("tls=%v j=%v, (%v:)", tls, j, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xllabs(tls, j)
}

func X__builtin_abs(tls *TLS, j int32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v j=%v, (%v:)", tls, j, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xabs(tls, j)
}

func X__builtin_bswap16(t *TLS, x uint16) uint16 {
	if __ccgo_strace {
		trc("t=%v x=%v, (%v:)", t, x, origin(2))
	}
	return x<<8 | x>>8
}

func X__builtin_bswap32(tls *TLS, x uint32) uint32 {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
	}
	return x>>int32(24) | x>>int32(8)&uint32(0xff00) | x<<int32(8)&uint32(0xff0000) | x<<int32(24)
}

func X__builtin_bswap64(tls *TLS, x uint64) uint64 {
	if __ccgo_strace {
		trc("tls=%v x=%v, (%v:)", tls, x, origin(2))
	}
	return x<<56 |
		x&0xff00<<40 |
		x&0xff0000<<24 |
		x&0xff000000<<8 |
		x&0xff00000000>>8 |
		x&0xff0000000000>>24 |
		x&0xff000000000000>>40 |
		x>>56
}

func Xalloca(tls *TLS, size Tsize_t) uintptr {
	if __ccgo_strace {
		trc("tls=%v size=%v, (%v:)", tls, size, origin(2))
	}
	return tls.alloca(size)
}

func X__builtin_alloca(tls *TLS, size Tsize_t) uintptr {
	if __ccgo_strace {
		trc("tls=%v size=%v, (%v:)", tls, size, origin(2))
	}
	return Xalloca(tls, size)
}

func X__builtin_copysign(tls *TLS, x float64, y float64) (r float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v y=%v, (%v:)", tls, x, y, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xcopysign(tls, x, y)
}

func X__builtin_copysignf(tls *TLS, x float32, y float32) (r float32) {
	if __ccgo_strace {
		trc("tls=%v x=%v y=%v, (%v:)", tls, x, y, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xcopysignf(tls, x, y)
}

func X__builtin_copysignl(t *TLS, x, y float64) float64 {
	if __ccgo_strace {
		trc("t=%v y=%v, (%v:)", t, y, origin(2))
	}
	return Xcopysign(t, x, y)
}

func X__builtin_expect(t *TLS, exp, c int64) int64 {
	if __ccgo_strace {
		trc("t=%v c=%v, (%v:)", t, c, origin(2))
	}
	return exp
}

func ___builtin_expect(t *TLS, exp, c long) long {
	return exp
}

func X__builtin_huge_val(t *TLS) float64 {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	return math.Inf(1)
}

func X__builtin_huge_valf(t *TLS) float32 {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	return float32(math.Inf(1))
}

func X__builtin_inf(t *TLS) float64 {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	return math.Inf(1)
}

func X__builtin_inff(t *TLS) float32 {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	return ___builtin_inff(t)
}

func X__builtin_isinfl(t *TLS, x float64) int32 {
	if __ccgo_strace {
		trc("t=%v x=%v, (%v:)", t, x, origin(2))
	}
	return Bool32(math.IsInf(x, 0))
}

func X__builtin_isinff(t *TLS, x float32) int32 {
	if __ccgo_strace {
		trc("t=%v x=%v, (%v:)", t, x, origin(2))
	}
	return Bool32(math.IsInf(float64(x), 0))
}

func X__builtin_isinf(t *TLS, x float64) int32 {
	if __ccgo_strace {
		trc("t=%v x=%v, (%v:)", t, x, origin(2))
	}
	return Bool32(math.IsInf(x, 0))
}

func ___builtin_inff(t *TLS) float32 {
	return float32(math.Inf(1))
}

func X__builtin_infl(t *TLS) float64 {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	return math.Inf(1)
}

func X__builtin_isunordered(t *TLS, a, b float64) int32 {
	if __ccgo_strace {
		trc("t=%v b=%v, (%v:)", t, b, origin(2))
	}
	return Bool32(math.IsNaN(a) || math.IsNaN(b))
}

func X__builtin_nan(t *TLS, s uintptr) float64 {
	if __ccgo_strace {
		trc("t=%v s=%v, (%v:)", t, s, origin(2))
	}
	return math.NaN()
}

func X__builtin_nanf(t *TLS, s uintptr) float32 {
	if __ccgo_strace {
		trc("t=%v s=%v, (%v:)", t, s, origin(2))
	}
	return ___builtin_nanf(t, s)
}

func ___builtin_nanf(t *TLS, s uintptr) float32 {
	return float32(math.NaN())
}

func X__builtin_nanl(t *TLS, s uintptr) float64 {
	if __ccgo_strace {
		trc("t=%v s=%v, (%v:)", t, s, origin(2))
	}
	return math.NaN()
}

func X__builtin_prefetch(t *TLS, addr, args uintptr) {
	if __ccgo_strace {
		trc("t=%v args=%v, (%v:)", t, args, origin(2))
	}
	// nop
}

func X__builtin_unreachable(t *TLS) {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	fmt.Fprintf(os.Stderr, "unrechable\n")
	os.Stderr.Sync()
	Xexit(t, 1)
}

func X__builtin_trap(t *TLS) {
	if __ccgo_strace {
		trc("t=%v, (%v:)", t, origin(2))
	}
	Xabort(t)
}

func Xsystem(t *TLS, command uintptr) int32 {
	if __ccgo_strace {
		trc("t=%v command=%v, (%v:)", t, command, origin(2))
	}
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

func X__builtin_isprint(tls *TLS, c int32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v c=%v, (%v:)", tls, c, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xisprint(tls, c)
}

func X__builtin_isalnum(tls *TLS, c int32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v c=%v, (%v:)", tls, c, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xisalnum(tls, c)
}

func X__builtin_isblank(tls *TLS, c int32) (r int32) {
	if __ccgo_strace {
		trc("tls=%v c=%v, (%v:)", tls, c, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xisblank(tls, c)
}

func X__builtin_clz(t *TLS, n uint32) int32 {
	if __ccgo_strace {
		trc("t=%v n=%v, (%v:)", t, n, origin(2))
	}
	return int32(bits.LeadingZeros32(n))
}

func X__builtin_clzll(t *TLS, n uint64) int32 {
	if __ccgo_strace {
		trc("t=%v n=%v, (%v:)", t, n, origin(2))
	}
	return int32(bits.LeadingZeros64(n))
}

func Xdlopen(tls *TLS, filename uintptr, flags int32) uintptr {
	if __ccgo_strace {
		trc("tls=%v filename=%v flags=%v, (%v:)", tls, filename, flags, origin(2))
	}
	panic(todo(""))
}

func Xdlsym(tls *TLS, handle, symbol uintptr) uintptr {
	if __ccgo_strace {
		trc("tls=%v symbol=%v, (%v:)", tls, symbol, origin(2))
	}
	panic(todo(""))
}

func Xtzset(tls *TLS) {
	if __ccgo_strace {
		trc("tls=%v, (%v:)", tls, origin(2))
	}
	___tzset(tls)
}

func Xfork(tls *TLS) int32 {
	if __ccgo_strace {
		trc("tls=%v, (%v:)", tls, origin(2))
	}
	return _fork(tls)
}

func X_exit(tls *TLS, status int32) {
	if __ccgo_strace {
		trc("tls=%v status=%v, (%v:)", tls, status, origin(2))
	}
	x__exit(tls, status)
}

func X__builtin_ldexp(tls *TLS, x float64, n int32) (r float64) {
	if __ccgo_strace {
		trc("tls=%v x=%v n=%v, (%v:)", tls, x, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	return Xldexp(tls, x, n)
}

func X__builtin_add_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	if __ccgo_strace {
		trc("t=%v b=%v res=%v, (%v:)", t, b, res, origin(2))
	}
	r, ovf := mathutil.AddOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func X__builtin_sub_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	if __ccgo_strace {
		trc("t=%v b=%v res=%v, (%v:)", t, b, res, origin(2))
	}
	r, ovf := mathutil.SubOverflowInt64(a, b)
	*(*int64)(unsafe.Pointer(res)) = r
	return Bool32(ovf)
}

func X__builtin_mul_overflowInt64(t *TLS, a, b int64, res uintptr) int32 {
	if __ccgo_strace {
		trc("t=%v b=%v res=%v, (%v:)", t, b, res, origin(2))
	}
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

func _setjmp(tls *TLS, jmp_buf uintptr) int32 {
	panic(todo(""))
}

func X__builtin_setjmp(tls *TLS, jmp_buf uintptr) int32 {
	if __ccgo_strace {
		trc("tls=%v jmp_buf=%v, (%v:)", tls, jmp_buf, origin(2))
	}
	return _setjmp(tls, jmp_buf)
}

func ___get_tp(tls *TLS) uintptr {
	return tls.fs
}

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

func _fork(tls *TLS) int32 {
	tls.setErrno(ENOSYS)
	return -1
}

func _abort(tls *TLS) {
	Xabort(tls)
}
