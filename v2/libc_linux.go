// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"sync/atomic"
	"unsafe"
)

// musl-1.2.4/arch/x86_64/atomic_arch.h:70:static inline void a_inc(volatile int *p)
func __ccgo_undefined_sia_inc(t *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), 1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:78:static inline void a_dec(volatile int *p)
func __ccgo_undefined_sia_dec(t *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), -1)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:20:static inline int a_swap(volatile int *p, int v)
func __ccgo_undefined_sia_swap(t *TLS, p uintptr, v int32) int32 {
	return atomic.SwapInt32((*int32)(unsafe.Pointer(p)), v)
}

// musl-1.2.4/src/internal/pthread_impl.h:168:static inline void __wake(volatile void *addr, int cnt, int priv)
func __ccgo_undefined_si__wake(t *TLS, addr uintptr, cnt, prin int32) {
	panic(todo(""))
}

// musl-1.2.4/arch/x86_64/syscall_arch.h:26:static __inline long __syscall3(long n, long a1, long a2, long a3)
func __ccgo_undefined_si__syscall3(t *TLS, n, a1, s2, a3 long) long {
	panic(todo("", n))
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:86:static inline void a_store(volatile int *p, int x)
func __ccgo_undefined_sia_store(t *TLS, p uintptr, x int32) {
	atomic.StoreInt32((*int32)(unsafe.Pointer(p)), x)
}

// musl-1.2.4/arch/x86_64/atomic_arch.h:94:static inline void a_barrier()
func __ccgo_undefined_sia_barrier(t *TLS) {
	panic(todo(""))
}

func a_cas(p *int32, t, s int32) int32

// musl-1.2.4/arch/x86_64/atomic_arch.h:2:static inline int a_cas(volatile int *p, int t, int s)
func __ccgo_undefined_sia_cas(t *TLS, p uintptr, test, s int32) int32 {
	return a_cas((*int32)(unsafe.Pointer(p)), test, s)
}

// // musl-1.2.4/arch/x86_64/pthread_arch.h:1:static inline uintptr_t __get_tp()
// func __ccgo_undefined_si__get_tp(t *TLS) TNuintptr_t {
// 	panic(todo(""))
// }
//
// func __ccgo_undefined_sidefault_malloc(tls *TLS, n TNsize_t) (r uintptr) {
// 	return _default_malloc(tls, n)
// }
