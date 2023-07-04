// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"os"
	"testing"
	"unsafe"

	"modernc.org/libc/v2/rtl"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

var (
	testAtomicCASInt32 int32
	testAtomicCASp     uintptr
)

func TestAtomicCASInt32(t *testing.T) {
	pi := uintptr(unsafe.Pointer(&testAtomicCASInt32))
	testAtomicCASInt32 = int32(0)
	j := a_cas(pi, 1, 2)
	if testAtomicCASInt32 != 0 || j != 0 {
		t.Fatal(testAtomicCASInt32, j)
	}

	if j = a_cas(pi, 0, 3); testAtomicCASInt32 != 3 || j != 0 {
		t.Fatal(testAtomicCASInt32, j)
	}

	if j = a_cas(pi, 4, 5); testAtomicCASInt32 != 3 || j != 3 {
		t.Fatal(testAtomicCASInt32, j)
	}

	if j = a_cas(pi, 3, 6); testAtomicCASInt32 != 6 || j != 3 {
		t.Fatal(testAtomicCASInt32, j)
	}
}

func TestAtomicCASUintptr(t *testing.T) {
	pp := uintptr(unsafe.Pointer(&testAtomicCASp))
	testAtomicCASp = 0
	j := a_cas_p(pp, 1, 2)
	if testAtomicCASp != 0 || j != 0 {
		t.Fatal(testAtomicCASp, j)
	}

	if j = a_cas_p(pp, 0, 3); testAtomicCASp != 3 || j != 0 {
		t.Fatal(testAtomicCASp, j)
	}

	if j = a_cas_p(pp, 4, 5); testAtomicCASp != 3 || j != 3 {
		t.Fatal(testAtomicCASp, j)
	}

	if j = a_cas_p(pp, 3, 6); testAtomicCASp != 6 || j != 3 {
		t.Fatal(testAtomicCASp, j)
	}
}

func TestAtomicOrInt32(t *testing.T) {
	pi := uintptr(unsafe.Pointer(&testAtomicCASInt32))
	testAtomicCASInt32 = int32(0)
	a_or(pi, 1)
	if j := testAtomicCASInt32; j != 1 {
		t.Fatalf("%032b", j)
	}

	a_or(pi, 2)
	if j := testAtomicCASInt32; j != 3 {
		t.Fatalf("%032b", j)
	}

	a_or(pi, rtl.Int32FromUint32(0x80000000))
	if j := testAtomicCASInt32; j != rtl.Int32FromUint32(0x80000003) {
		t.Fatalf("%032b", j)
	}
}
