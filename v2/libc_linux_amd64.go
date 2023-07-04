// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"math"
	"sync/atomic"
	"unsafe"

	"modernc.org/libc/v2/rtl"
)

type (
	long = int64
)

// // musl-1.2.4/src/env/__init_tls.c:81:extern weak hidden const size_t _DYNAMIC[];
// var __DYNAMIC [1]Tsize_t

func a_cas(p uintptr, t, s int32) int32

// musl-1.2.4/arch/x86_64/atomic_arch.h:2:static inline int a_cas(volatile int *p, int t, int s)
func _a_cas(tls *rtl.TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

func a_cas_p(p uintptr, t, s uintptr) uintptr

// musl-1.2.4/arch/x86_64/atomic_arch.h:11:static inline void *a_cas_p(volatile void *p, void *t, void *s)
func _a_cas_p(tls *rtl.TLS, p, test, s uintptr) uintptr {
	return a_cas_p(p, test, s)
}

func a_or(p uintptr, v int32)

// musl-1.2.4/arch/x86_64/atomic_arch.h:46:static inline void a_or(volatile int *p, int v)
func _a_or(tls *rtl.TLS, p uintptr, v int32) {
	a_or(p, v)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:70:static inline void a_inc(volatile int *p)
func _a_inc(tls *rtl.TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), 1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:78:static inline void a_dec(volatile int *p)
func _a_dec(tls *rtl.TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), -1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:20:static inline int a_swap(volatile int *p, int v)
func _a_swap(tls *rtl.TLS, p uintptr, v int32) int32 {
	return atomic.SwapInt32((*int32)(unsafe.Pointer(p)), v)
}

// musl-1.2.4/src/internal/pthread_impl.h:168:static inline void __wake(volatile void *addr, int cnt, int priv)
func ___wake(tls *rtl.TLS, addr uintptr, cnt, prin int32) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:4:static __inline long __syscall0(long n)
func ___syscall0(tls *rtl.TLS, n long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:11:static __inline long __syscall1(long n, long a1)
func ___syscall1(tls *rtl.TLS, n, a1 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:18:static __inline long __syscall2(long n, long a1, long a2)
func ___syscall2(tls *rtl.TLS, n, a1, a2 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:26:static __inline long __syscall3(long n, long a1, long a2, long a3)
func ___syscall3(tls *rtl.TLS, n, a1, s2, a3 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:34:static __inline long __syscall4(long n, long a1, long a2, long a3, long a4)
func ___syscall4(tls *rtl.TLS, n, a1, s2, a3, a4 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:43:static __inline long __syscall5(long n, long a1, long a2, long a3, long a4, long a5)
func ___syscall5(tls *rtl.TLS, n, a1, s2, a3, a4, a5 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:53:static __inline long __syscall6(long n, long a1, long a2, long a3, long a4, long a5, long a6)
func ___syscall6(tls *rtl.TLS, n, a1, s2, a3, a4, a5, a6 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:86:static inline void a_store(volatile int *p, int x)
func _a_store(tls *rtl.TLS, p uintptr, x int32) {
	atomic.StoreInt32((*int32)(unsafe.Pointer(p)), x)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:94:static inline void a_barrier()
func _a_barrier(tls *rtl.TLS) {
	panic(todo(""))
}

// musl-1.2.4/include/math.h:61:static __inline unsigned long long __DOUBLE_BITS(double __f)
func ___DOUBLE_BITS(tls *rtl.TLS, f float64) uint64 {
	return math.Float64bits(f)
}

// musl-1.2.4/include/math.h:55:static __inline unsigned __FLOAT_BITS(float __f)
func ___FLOAT_BITS(tls *rtl.TLS, f float32) uint32 {
	return math.Float32bits(f)
}

// float __builtin_inff (void)
func ___builtin_inff(tls *rtl.TLS) float32 {
	return float32(math.Inf(1))
}

// float __builtin_nanf (char*)
func ___builtin_nanf(tls *rtl.TLS, _ uintptr) float32 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/pthread_impl.h:161:hidden int __set_thread_area(void *);
func ___set_thread_area(tls *rtl.TLS, p uintptr) int32 {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:106:static inline void a_crash()
func _a_crash(tls *rtl.TLS) {
	panic(todo(""))
}

// int fegetround(void)
func _fegetround(tls *rtl.TLS) int32 {
	panic(todo(""))
}

// int fetestexcept(int mask)
func _fetestexcept(tls *rtl.TLS, mask int32) int32 {
	panic(todo(""))
}

// int fegetenv(fenv_t *envp)
func _fegetenv(tls *rtl.TLS, envp uintptr) int32 {
	panic(todo(""))
}

// int feclearexcept(int mask)
func _feclearexcept(tls *rtl.TLS, mask int32) int32 {
	panic(todo(""))
}

// int feraiseexcept(int mask)
func _feraiseexcept(tls *rtl.TLS, mask int32) int32 {
	panic(todo(""))
}

// int __fesetround(int r)
func ___fesetround(tls *rtl.TLS, r int32) int32 {
	panic(todo(""))
}

// int fesetenv(const fenv_t *envp)
func _fesetenv(tls *rtl.TLS, envp uintptr) int32 {
	panic(todo(""))
}

// static __inline int __isspace(int _c)
//
//	{
//		return _c == ' ' || (unsigned)_c-'\t' < 5;
//	}
func ___isspace(tls *rtl.TLS, c int32) int32 {
	return rtl.Bool32(c == ' ' || uint32(c)-'\t' < 5)
}

// musl-1.2.4/include/endian.h:24:static __inline uint32_t __bswap32(uint32_t __x)
func ___bswap32(tls *rtl.TLS, x uint32) uint32 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:129:static inline struct meta *get_meta(const unsigned char *p)
func _get_meta(tls *rtl.TLS, p uintptr) uintptr {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:124:static inline int get_slot_index(const unsigned char *p)
func _get_slot_index(tls *rtl.TLS, p uintptr) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:75:static inline void queue(struct meta **phead, struct meta *m)
func _queue(tls *rtl.TLS, phead, m uintptr) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:249:static inline void step_seq(void)
func _step_seq(tls *rtl.TLS) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:259:static inline void record_seq(int sc)
func _record_seq(tls *rtl.TLS, sc int32) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:109:static inline void free_meta(struct meta *m)
func _free_meta(tls *rtl.TLS, m uintptr) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:283:static inline int is_bouncing(int sc)
func _is_bouncing(tls *rtl.TLS, sc int32) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:90:static inline void dequeue(struct meta **phead, struct meta *m)
func _dequeue(tls *rtl.TLS, phead, m uintptr) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:115:static inline uint32_t activate_group(struct meta *m)
func _activate_group(tls *rtl.TLS, m uintptr) uint32 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/glue.h:72:static inline void wrlock()
func _wrlock(tls *rtl.TLS) {
	panic(todo(""))
}

// // musl-1.2.4/arch/x86_64/pthread_arch.h:1:static inline uintptr_t __get_tp()
// func ___get_tp(tls *rtl.TLS) Tuintptr_t {
// 	panic(todo(""))
// }

// func default_malloc(tls *rtl.TLS, n Tsize_t) (r uintptr) {
// 	return _default_malloc(tls, n)
// }

// func ___clone(tls *rtl.TLS, func1 uintptr, stack uintptr, flags int32, arg uintptr, va uintptr) (r int32) {
// 	return -m_ENOSYS
// }

// // musl-1.2.4/src/malloc/mallocng/meta.h:175:static inline size_t get_stride(const struct meta *g)
// func _get_stride(tls *rtl.TLS, g uintptr) Tsize_t {
// 	panic(todo(""))
// }

// // musl-1.2.4/src/malloc/mallocng/meta.h:184:static inline void set_size(unsigned char *p, unsigned char *end, size_t n)
// func _set_size(tls *rtl.TLS, p, end uintptr, n Tsize_t) {
// 	panic(todo(""))
// }

// // musl-1.2.4/src/malloc/mallocng/meta.h:159:static inline size_t get_nominal_size(const unsigned char *p, const unsigned char *end)
// func _get_nominal_size(tls *rtl.TLS, p, end uintptr) Tsize_t {
// 	panic(todo(""))
// }
