// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !libc.membrk && !libc.memgrind
// +build !libc.membrk,!libc.memgrind

package libc // import "modernc.org/libc/v2"

import (
	"math"
	"math/bits"
	"sync"

	"modernc.org/memory"
)

const (
	isMemBrk = false
)

var (
	allocator   memory.Allocator
	allocatorMu sync.Mutex
)

func Xmalloc(tls *TLS, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v n=%v, (%v:)", tls, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	if n > math.MaxInt {
		tls.setErrno(ENOMEM)
		return 0
	}

	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	if n == 0 {
		// malloc(0) should return unique pointers
		// (often expected and gnulib replaces malloc if malloc(0) returns 0)
		n = 1
	}
	var err error
	if r, err = allocator.UintptrMalloc(int(n)); err != nil {
		r = 0
		tls.setErrno(ENOMEM)
	}
	return r
}

func Xcalloc(tls *TLS, m Tsize_t, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v m=%v n=%v, (%v:)", tls, m, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	hi, rq := bits.Mul(uint(m), uint(n))
	if hi != 0 || rq > math.MaxInt {
		tls.setErrno(ENOMEM)
		return 0
	}

	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	if rq == 0 {
		rq = 1
	}

	var err error
	if r, err = allocator.UintptrCalloc(int(rq)); err != nil {
		r = 0
		tls.setErrno(ENOMEM)
	}
	return r
}

func Xrealloc(tls *TLS, p uintptr, n Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v p=%v n=%v, (%v:)", tls, p, n, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	var err error
	if r, err = allocator.UintptrRealloc(p, int(n)); err != nil {
		r = 0
		tls.setErrno(ENOMEM)
	}
	return r
}

func Xfree(tls *TLS, p uintptr) {
	if __ccgo_strace {
		trc("tls=%v p=%v, (%v:)", tls, p, origin(2))
	}
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	allocator.UintptrFree(p)
}

func Xmalloc_usable_size(tls *TLS, p uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("tls=%v p=%v, (%v:)", tls, p, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	if p == 0 {
		return 0
	}

	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return Tsize_t(memory.UintptrUsableSize(p))
}

func MemAudit() (r []*MemAuditError) {
	panic(todo(""))
	return nil
}
