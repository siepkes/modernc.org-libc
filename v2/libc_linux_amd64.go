// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"modernc.org/memory"
)

var (
	initLibcOnce sync.Once

	allocator   memory.Allocator
	allocatorMu sync.Mutex
	mainFP      uintptr      // Hook to start C main, if any.
	tid         atomic.Int32 // TLS id
)

// arch/x86_64/crt_arch.h
//
//	__asm__(
//	".text \n"
//	".global " START " \n"
//	START ": \n"
//	"	xor %rbp,%rbp \n"
//	"	mov %rsp,%rdi \n"
//	".weak _DYNAMIC \n"
//	".hidden _DYNAMIC \n"
//	"	lea _DYNAMIC(%rip),%rsi \n"
//	"	andq $-16,%rsp \n"
//	"	call " START "_c \n"
//	);
//
// crt/crt1.c
//
//	#include <features.h>
//	#include "libc.h"
//
//	#define START "_start"
//
//	#include "crt_arch.h"
//
//	int main();
//	weak void _init();
//	weak void _fini();
//	int __libc_start_main(int (*)(), int, char **,
//		void (*)(), void(*)(), void(*)());
//
//	void _start_c(long *p)
//	{
//		int argc = p[0];
//		char **argv = (void *)(p+1);
//		__libc_start_main(main, argc, argv, _init, _fini, 0);
//	}
func initLibc() {
	var argv []uintptr
	for _, v := range os.Args {
		p := mustCalloc(len(v) + 1)
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		argv = append(argv, p)
	}
	argv = append(argv, 0, 0, 0, 0) //TODO envp, auxv
	argvP := mustCalloc(len(argv) * int(unsafe.Sizeof(uintptr(0))))
	copy(unsafe.Slice((*uintptr)(unsafe.Pointer(argvP)), len(argv)), argv)

	t := newTLS()

	defer t.Close()

	x___libc_start_main(t, mainFP, int32(len(os.Args)), argvP, 0, 0, 0)
}

// CString returns a pointer to a zero-terminated version of s. The caller is
// responsible for freeing the allocated memory.
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

func mustCalloc(sz int) uintptr {
	if p, _ := calloc(sz); p != 0 {
		return p
	}

	panic(todo("OOM"))
}

func mustMalloc(sz int) uintptr {
	if p, _ := malloc(sz); p != 0 {
		return p
	}

	panic(todo("OOM"))
}

func calloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrCalloc(sz)
}

func malloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrMalloc(sz)
}

func free(p uintptr) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	allocator.UintptrFree(p)
}

// TLS emulates thread local storage.
type TLS struct {
	tp    uintptr   // *T__pthread, returned by ___get_tp()
	stack []uintptr //TODO reuse libc/v1 mechanism

	ID                 int32
	reentryGuard       int32 // memgrind
	stackHeaderBalance int32
}

// NewTLS returns a newly created TLS that must be eventually closed to prevent
// resource leaks.
func NewTLS() *TLS {
	initLibcOnce.Do(initLibc)
	return newTLS()
}

func newTLS() *TLS {
	t := &TLS{
		ID: tid.Add(1),
	}
	return t
}

func (t *TLS) Alloc(n int) (r uintptr) { //TODO reuse libc/v1 mechanism
	r = mustMalloc(n)
	t.stack = append(t.stack, r)
	return r
}

func (t *TLS) Free(n int) { //TODO reuse libc/v1 mechanism
	p := t.stack[len(t.stack)-1]
	t.stack = t.stack[:len(t.stack)-1]
	free(p)
}

func (t *TLS) Close() {
	*t = TLS{} //TODO more
}

// musl-1.2.4/src/env/__init_tls.c:81:extern weak hidden const Tsize_t _DYNAMIC[];
var __DYNAMIC [1]uint64

func ___clone(tls *TLS, func1 uintptr, stack uintptr, flags int32, arg uintptr, va uintptr) (r int32) {
	return -m_ENOSYS
}

// musl-1.2.4/include/setjmp.h:39:_Noreturn void longjmp (jmp_buf, int);
func _longjmp(tls *TLS, env uintptr, val int32) int64 {
	panic(todo(""))
}

// musl-1.2.4/src/thread/pthread_cancel.c:17:long __syscall_cp_asm(volatile void *, syscall_arg_t,
func ___syscall_cp_asm(tls *TLS, _ uintptr, _, _, _, _, _, _, _ int64) int64 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/pthread_impl.h:163:hidden void __unmapself(void *, size_t);
func ___unmapself(tls *TLS, _ uintptr, _ uint64) int64 {
	panic(todo(""))
}

// musl-1.2.4/src/thread/pthread_cancel.c:46:extern hidden const char __cp_begin[1], __cp_end[1], __cp_cancel[1];
var ___cp_begin [1]int8
var ___cp_cancel [1]int8
var ___cp_end [1]int8

func a_cas(p uintptr, t, s int32) int32

// musl-1.2.4/arch/x86_64/atomic_arch.h:2:static inline int a_cas(volatile int *p, int t, int s)
func _a_cas(tls *TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:38:static inline void a_and(volatile int *p, int v)
func _a_and(tls *TLS, p uintptr, v int32) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:100:static inline void a_spin()
func _a_spin(tls *TLS) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:29:static inline int a_fetch_add(volatile int *p, int v)
func _a_fetch_add(tls *TLS, p uintptr, v int32) int32 {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:119:static inline int a_clz_64(uint64_t x)
func _a_clz_64(tls *TLS, x uint64) int32 {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:112:static inline int a_ctz_64(uint64_t x)
func _a_ctz_64(tls *TLS, x uint64) int32 {
	panic(todo(""))
}

func a_cas_p(p uintptr, t, s uintptr) uintptr

// musl-1.2.4/arch/x86_64/atomic_arch.h:11:static inline void *a_cas_p(volatile void *p, void *t, void *s)
func _a_cas_p(tls *TLS, p, test, s uintptr) uintptr {
	return a_cas_p(p, test, s)
}

func a_or(p uintptr, v int32)

// musl-1.2.4/arch/x86_64/atomic_arch.h:46:static inline void a_or(volatile int *p, int v)
func _a_or(tls *TLS, p uintptr, v int32) {
	a_or(p, v)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:62:static inline void a_or_64(volatile uint64_t *p, uint64_t v)
func _a_or_64(tls *TLS, p uintptr, v uint64) {
	panic(todo(""))
}

// musl-1.2.4/src/internal/atomic.h:256:static inline int a_ctz_32(uint32_t x)
func _a_ctz_32(tls *TLS, x uint32) int32 {
	return int32(bits.TrailingZeros32(x))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:70:static inline void a_inc(volatile int *p)
func _a_inc(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), 1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:78:static inline void a_dec(volatile int *p)
func _a_dec(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), -1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:20:static inline int a_swap(volatile int *p, int v)
func _a_swap(tls *TLS, p uintptr, v int32) int32 {
	return atomic.SwapInt32((*int32)(unsafe.Pointer(p)), v)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:4:static __inline long __syscall0(long n)
func ___syscall0(tls *TLS, n int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), 0, 0, 0)
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:11:static __inline long __syscall1(long n, long a1)
func ___syscall1(tls *TLS, n, a1 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), 0, 0)
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:18:static __inline long __syscall2(long n, long a1, long a2)
func ___syscall2(tls *TLS, n, a1, a2 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), 0)
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:26:static __inline long __syscall3(long n, long a1, long a2, long a3)
func ___syscall3(tls *TLS, n, a1, a2, a3 int64) int64 {
	r1, _, err := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3))
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:34:static __inline long __syscall4(long n, long a1, long a2, long a3, long a4)
func ___syscall4(tls *TLS, n, a1, a2, a3, a4 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), 0, 0)
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:43:static __inline long __syscall5(long n, long a1, long a2, long a3, long a4, long a5)
func ___syscall5(tls *TLS, n, a1, a2, a3, a4, a5 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), 0)
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:53:static __inline long __syscall6(long n, long a1, long a2, long a3, long a4, long a5, long a6)
func ___syscall6(tls *TLS, n, a1, a2, a3, a4, a5, a6 int64) int64 {
	r1, _, err := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), uintptr(a6))
	if err != 0 {
		panic(todo(""))
	}

	return int64(r1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:86:static inline void a_store(volatile int *p, int x)
func _a_store(tls *TLS, p uintptr, x int32) {
	atomic.StoreInt32((*int32)(unsafe.Pointer(p)), x)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:94:static inline void a_barrier()
func _a_barrier(tls *TLS) {
	panic(todo(""))
}

// musl-1.2.4/include/math.h:61:static __inline unsigned long long __DOUBLE_BITS(double __f)
func ___DOUBLE_BITS(tls *TLS, f float64) uint64 {
	return math.Float64bits(f)
}

// musl-1.2.4/include/math.h:55:static __inline unsigned __FLOAT_BITS(float __f)
func ___FLOAT_BITS(tls *TLS, f float32) uint32 {
	return math.Float32bits(f)
}

// float __builtin_inff (void)
func ___builtin_inff(tls *TLS) float32 {
	return float32(math.Inf(1))
}

// float __builtin_nanf (char*)
func ___builtin_nanf(tls *TLS, _ uintptr) float32 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/pthread_impl.h:161:hidden int __set_thread_area(void *);
func ___set_thread_area(tls *TLS, p uintptr) int32 {
	tls.tp = p
	return 0
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:106:static inline void a_crash()
func _a_crash(tls *TLS) {
	panic(todo(""))
}

// int fegetround(void)
func _fegetround(tls *TLS) int32 {
	panic(todo(""))
}

// int fetestexcept(int mask)
func _fetestexcept(tls *TLS, mask int32) int32 {
	panic(todo(""))
}

// int fegetenv(fenv_t *envp)
func _fegetenv(tls *TLS, envp uintptr) int32 {
	panic(todo(""))
}

// int feclearexcept(int mask)
func _feclearexcept(tls *TLS, mask int32) int32 {
	panic(todo(""))
}

// int feraiseexcept(int mask)
func _feraiseexcept(tls *TLS, mask int32) int32 {
	panic(todo(""))
}

// int __fesetround(int r)
func ___fesetround(tls *TLS, r int32) int32 {
	panic(todo(""))
}

// int fesetenv(const fenv_t *envp)
func _fesetenv(tls *TLS, envp uintptr) int32 {
	panic(todo(""))
}

// musl-1.2.4/include/byteswap.h:7:static __inline uint16_t __bswap_16(uint16_t __x)
func ___bswap_16(tls *TLS, __x uint16) uint16 {
	return __x<<8 | __x>>8
}

// musl-1.2.4/include/endian.h:24:static __inline uint32_t __bswap32(uint32_t __x)
func ___bswap32(tls *TLS, __x uint32) uint32 {
	return __x>>int32(24) | __x>>int32(8)&uint32(0xff00) | __x<<int32(8)&uint32(0xff0000) | __x<<int32(24)
}

// musl-1.2.4/include/byteswap.h:12:static __inline uint32_t __bswap_32(uint32_t __x)
func ___bswap_32(tls *TLS, x uint32) uint32 {
	return ___bswap32(tls, x)
}

// musl-1.2.4/arch/x86_64/pthread_arch.h:1:static inline uintptr_t __get_tp()
//
//	static inline uintptr_t __get_tp()
//	{
//		uintptr_t tp;
//		__asm__ ("mov %%fs:0,%0" : "=r" (tp) );
//		return tp;
//	}
//
//	#define MC_PC gregs[REG_RIP]
//
// https://stackoverflow.com/questions/6611346/how-are-the-fs-gs-registers-used-in-linux-amd64
func ___get_tp(tls *TLS) uintptr {
	return tls.tp
}

// void *memcpy(void *restrict dest, const void *restrict src, size_t n);
func _memcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return _memmove(tls, dest, src, n)
}

// void *memmove(void *dest, const void *src, size_t n);
func _memmove(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	if n != 0 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(dest)), n), unsafe.Slice((*byte)(unsafe.Pointer(src)), n))
	}
	return dest
}

// void *memset(void *s, int c, size_t n);
func _memset(tls *TLS, s uintptr, c int32, n uint64) (r uintptr) {
	if n != 0 {
		c := byte(c & 0xff)
		b := unsafe.Slice((*byte)(unsafe.Pointer(s)), n)
		for i := range b {
			b[i] = c
		}
	}
	return s
}

// long __builtin_expect(long, long);
func ___builtin_expect(tls *TLS, exp, c int64) int64 {
	return exp
}

// double fabs(double x);
func _fabs(tls *TLS, x float64) float64 {
	panic(todo(""))
}

// float fabsf(float x);
func _fabsf(tls *TLS, x float32) float32 {
	panic(todo(""))
}

// long double fabsl(long double x);
func _fabsl(tls *TLS, x float64) float64 {
	panic(todo(""))
}

// musl-1.2.4/src/signal/x86_64/restore.s:2
func ___restore_rt(tls *TLS) {
	panic(todo(""))
}

// int setjmp(jmp_buf env);
func _setjmp(tls *TLS, env uintptr) int32 {
	panic(todo(""))
}

// ptrdiff_t __tlsdesc_static()
func ___tlsdesc_static(tls *TLS) int64 {
	panic(todo(""))
}

// ptrdiff_t __tlsdesc_dynamic()
func ___tlsdesc_dynamic(tls *TLS) int64 {
	panic(todo(""))
}

// ----------------------------------------------------------------------------

// Xfmod computes the floating-point remainder of dividing x by y.  The return
// value is x - n * y, where n is the quotient of x / y, rounded toward zero to
// an integer.
func Xfmod(tls *TLS, x float64, y float64) (r float64) {
	return x_fmod(tls, x, y)
}

// Xstrtod converts the initial portion of the string pointed to by nptr to double representation.
func Xstrtod(tls *TLS, nptr, endptr uintptr) float64 {
	return x_strtod(tls, nptr, endptr)
}

// Xrint rounds its argument to an integer value in floating-point format, using the current rounding direction.
func Xrint(tls *TLS, x float64) float64 {
	return x_rint(tls, x)
}

// Xmemset fills the first n bytes of the memory area pointed to by s with the constant byte c.
func Xmemset(tls *TLS, s uintptr, c int32, n uint64) uintptr {
	return _memset(tls, s, c, n)
}

// Xgetentropy writes length bytes of high-quality random data to the buffer
// starting at the location pointed to by buffer.  The maximum permitted value
// for the length argument is 256.
//
// A successful call to getentropy() always provides the requested number of
// bytes of entropy.
func Xgetentropy(tls *TLS, buffer uintptr, length uint64) int32 {
	return x_getentropy(tls, buffer, length)
}

// Xmalloc allocates  size  bytes and returns a pointer to the allocated
// memory.  The memory is not initialized.  If size is 0, then Xmalloc()
// returns either NULL, or a unique pointer value that can later be
// successfully passed to Xfree().
func Xmalloc(tls *TLS, size uint64) uintptr {
	return _default_malloc(tls, size)
}

// Xreallocarray changes the size of the memory block pointed to by ptr to be
// large enough for an array of nmemb elements, each of which is size bytes.
// It is equivalent to the call
//
//	Xrealloc(tls, ptr, nmemb * size)
//
// However, unlike that realloc() call, reallocarray() fails safely in the case
// where the multiplication would overflow.  If such an overflow occurs,
// reallocarray() returns NULL,  sets  errno to ENOMEM, and leaves the original
// block of memory unchanged.
func Xreallocarray(tls *TLS, ptr uintptr, nmemb, size uint64) uintptr {
	return x_reallocarray(tls, ptr, nmemb, size)
}

// Xfree frees the memory space pointed to by ptr, which must have been
// returned by a previous call to Xmalloc(), Xcalloc(), or Xrealloc().
// Otherwise, or if Xfree(ptr) has  already been called before, undefined
// behavior occurs.  If ptr is NULL, no operation is performed.
func Xfree(tls *TLS, ptr uintptr) {
	x_free(tls, ptr)
}

// Xsnprintf writes at most size bytes (including the terminating null byte
// ('\0')) to str, according to a format and args.
func Xsnprintf(tls *TLS, str uintptr, size uint64, format, args uintptr) int32 {
	return x_snprintf(tls, str, size, format, args)
}

// Xfdopen associates a stream with the existing file descriptor, fd.  The mode
// of the stream (one of the values "r", "r+", "w", "w+", "a", "a+") must  be
// compatible  with  the mode  of the file descriptor.  The file position
// indicator of the new stream is set to that belonging to fd, and the error
// and end-of-file indicators are cleared.  Modes "w" or "w+" do not cause
// truncation of the file.  The file descriptor is not dup'ed, and will be
// closed when the stream created by fdopen() is closed.  The result of
// applying fdopen() to a shared memory  object is undefined.
func Xfdopen(tls *TLS, fd int32, mode uintptr) uintptr {
	return x___fdopen(tls, fd, mode)
}

// Xfread reads nmemb items of data, each size bytes long, from the stream
// pointed to by stream, storing them at the location given by ptr.
func Xfread(tls *TLS, ptr uintptr, size, nmemb uint64, stream uintptr) uint64 {
	return x_fread(tls, ptr, size, nmemb, stream)
}
