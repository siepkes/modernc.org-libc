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
	char = int8
	long = int64
)

// musl-1.2.4/src/env/__init_tls.c:81:extern weak hidden const Tsize_t _DYNAMIC[];
// var __DYNAMIC [1]Tsize_t

// musl-1.2.4/src/thread/pthread_cancel.c:46:extern hidden const char __cp_begin[1], __cp_end[1], __cp_cancel[1];
var ___cp_begin [1]char
var ___cp_cancel [1]char
var ___cp_end [1]char

func a_cas(p uintptr, t, s int32) int32

// musl-1.2.4/arch/x86_64/atomic_arch.h:2:static inline int a_cas(volatile int *p, int t, int s)
func _a_cas(tls *rtl.TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:38:static inline void a_and(volatile int *p, int v)
func _a_and(tls *rtl.TLS, p uintptr, v int32) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:100:static inline void a_spin()
func _a_spin(tls *rtl.TLS) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:29:static inline int a_fetch_add(volatile int *p, int v)
func _a_fetch_add(tls *rtl.TLS, p uintptr, v int32) int32 {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:119:static inline int a_clz_64(uint64_t x)
func _a_clz_64(tls *rtl.TLS, x uint64) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/atomic.h:292:static inline int a_ctz_l(unsigned long x)
func _a_ctz_l(tls *rtl.TLS, x uint64) int32 {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:112:static inline int a_ctz_64(uint64_t x)
func _a_ctz_64(tls *rtl.TLS, x uint64) int32 {
	panic(todo(""))
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

// musl-1.2.4/src/internal/atomic.h:256:static inline int a_ctz_32(uint32_t x)
func _a_ctz_32(tls *rtl.TLS, x uint32) int32 {
	panic(todo(""))
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

// musl-1.2.4/include/byteswap.h:7:static __inline uint16_t __bswap_16(uint16_t __x)
func ___bswap_16(tls *rtl.TLS, x uint16) uint16 {
	panic(todo(""))
}

// musl-1.2.4/include/endian.h:24:static __inline uint32_t __bswap32(uint32_t __x)
func ___bswap32(tls *rtl.TLS, x uint32) uint32 {
	panic(todo(""))
}

// musl-1.2.4/include/byteswap.h:12:static __inline uint32_t __bswap_32(uint32_t __x)
func ___bswap_32(tls *rtl.TLS, x uint32) uint32 {
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

// musl-1.2.4/arch/x86_64/pthread_arch.h:1:static inline uintptr_t __get_tp()
func ___get_tp(tls *rtl.TLS) uintptr {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/glue.h:88:static inline void malloc_atfork(int who)
func _malloc_atfork(tls *rtl.TLS, who int32) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/glue.h:44:static inline uint64_t get_random_secret()
func _get_random_secret(tls *rtl.TLS) uint64 {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:102:static inline struct meta *dequeue_head(struct meta **phead)
func _dequeue_head(tls *rtl.TLS, phead uintptr) uintptr {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:277:static inline void decay_bounces(int sc)
func _decay_bounces(tls *rtl.TLS, sc int32) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/meta.h:264:static inline void account_bounce(int sc)
func _account_bounce(tls *rtl.TLS, sc int32) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/glue.h:68:static inline void rdlock()
func _rdlock(tls *rtl.TLS) {
	panic(todo(""))
}

// musl-1.2.4/src/malloc/mallocng/glue.h:80:static inline void upgradelock()
func _upgradelock(tls *rtl.TLS) {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:122:static inline float fp_barrierf(float x)
func _fp_barrierf(tls *rtl.TLS, x float32) float32 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:110:static inline double eval_as_double(double x)
func _eval_as_double(tls *rtl.TLS, x float64) float64 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:104:static inline float eval_as_float(float x)
func _eval_as_float(tls *rtl.TLS, x float32) float32 {
	panic(todo(""))
}

// int __builtin_expect(int, int);
func ___builtin_expect(tls *rtl.TLS, _, _ int32) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:155:static inline void fp_force_evalf(float x)
func _fp_force_evalf(tls *rtl.TLS, x float32) {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:164:static inline void fp_force_eval(double x)
func _fp_force_eval(tls *rtl.TLS, x float64) {
	panic(todo(""))
}

// musl-1.2.4/src/internal/libm.h:173:static inline void fp_force_evall(long double x)
func _fp_force_evall(tls *rtl.TLS, x float64) {
	panic(todo(""))
}

// musl-1.2.4/src/internal/pthread_impl.h:175:static inline void __futexwait(volatile void *addr, int val, int priv)
func ___futexwait(tls *rtl.TLS, addr uintptr, val, priv int32) {
	panic(todo(""))
}

// musl-1.2.4/src/stdio/getc.h:16:static inline int do_getc(FILE *f)
func _do_getc(tls *rtl.TLS, f uintptr) int32 {
	panic(todo(""))
}

// musl-1.2.4/src/stdio/putc.h:16:static inline int do_putc(int c, FILE *f)
func _do_putc(tls *rtl.TLS, c int32, f uintptr) int32 {
	panic(todo(""))
}

// ============================================================================

// func ___clone(tls *rtl.TLS, func1 uintptr, stack uintptr, flags int32, arg uintptr, va uintptr) (r int32) {
// 	return -m_ENOSYS
// }
//
// func default_malloc(tls *rtl.TLS, n Tsize_t) (r uintptr) {
// 	return _default_malloc(tls, n)
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:175:static inline Tsize_t get_stride(const struct meta *g)
// func _get_stride(tls *rtl.TLS, g uintptr) Tsize_t {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:184:static inline void set_size(unsigned char *p, unsigned char *end, Tsize_t n)
// func _set_size(tls *rtl.TLS, p, end uintptr, n Tsize_t) {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:159:static inline Tsize_t get_nominal_size(const unsigned char *p, const unsigned char *end)
// func _get_nominal_size(tls *rtl.TLS, p, end uintptr) Tsize_t {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/internal/libm.h:131:static inline double fp_barrier(double x)
// func _fp_barrier(tls *rtl.TLS, x float64) float64 {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:229:static inline int size_to_class(Tsize_t n)
// func _size_to_class(tls *rtl.TLS, n Tsize_t) int32 {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:196:static inline void *enframe(struct meta *g, int idx, Tsize_t n, int ctr)
// func _enframe(tls *rtl.TLS, g uintptr, idx int32, n Tsize_t, ctr int32) uintptr {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/malloc/mallocng/meta.h:240:static inline int size_overflows(Tsize_t n)
// func _size_overflows(tls *rtl.TLS, n Tsize_t) int32 {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/internal/syscall.h:61:static inline long __alt_socketcall(int sys, int sock, int cp, Tsyscall_arg_t a, Tsyscall_arg_t b, Tsyscall_arg_t c, Tsyscall_arg_t d, Tsyscall_arg_t e, Tsyscall_arg_t f)
// func ___alt_socketcall(tls *rtl.TLS, sys, sock, cp int32, a, b, c, d, e, f Tsyscall_arg_t) long {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/include/setjmp.h:39:_Noreturn void longjmp (jmp_buf, int);
// func _longjmp(tls *rtl.TLS, _ Tjmp_buf, _ int32) long {
// 	panic(todo(""))
// }
//
// // musl-1.2.4/src/thread/pthread_cancel.c:17:long __syscall_cp_asm(volatile void *, syscall_arg_t,
// // syscall_arg_t, syscall_arg_t, syscall_arg_t, syscall_arg_t, syscall_arg_t, syscall_arg_t);
// func ___syscall_cp_asm(tls *rtl.TLS, _ uintptr, _, _, _, _, _, _, _ Tsyscall_arg_t) long {
// 	panic(todo(""))
// }
