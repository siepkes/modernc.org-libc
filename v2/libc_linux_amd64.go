// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

type (
	long = int64
)

func a_cas(p uintptr, t, s int32) int32

// musl-1.2.4/arch/x86_64/atomic_arch.h:2:static inline int a_cas(volatile int *p, int t, int s)
func _a_cas(t *TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

func a_cas_p(p uintptr, t, s uintptr) uintptr

// musl-1.2.4/arch/x86_64/atomic_arch.h:11:static inline void *a_cas_p(volatile void *p, void *t, void *s)
func _a_cas_p(t *TLS, p, test, s uintptr) uintptr {
	return a_cas_p(p, test, s)
}

func a_or(p uintptr, v int32)

// musl-1.2.4/arch/x86_64/atomic_arch.h:46:static inline void a_or(volatile int *p, int v)
func _a_or(t *TLS, p uintptr, v int32) {
	a_or(p, v)
}
