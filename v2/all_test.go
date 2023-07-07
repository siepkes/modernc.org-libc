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

func TestXfmod(t *testing.T) {
	tls := rtl.NewTLS()

	defer tls.Close()

	x := 1.3518643030646695
	y := 6.283185307179586
	if g, e := Xfmod(tls, x, y), 1.3518643030646695; g != e {
		t.Fatal(g, e)
	}
}

func TestSwap(t *testing.T) {
	if g, e := ___bswap_16(nil, 0x1234), uint16(0x3412); g != e {
		t.Errorf("%#04x %#04x", g, e)
	}
	if g, e := ___bswap_32(nil, 0x12345678), uint32(0x78563412); g != e {
		t.Errorf("%#04x %#04x", g, e)
	}
}

var (
	valist       [256]byte
	formatString [256]byte
	srcString    [256]byte
	printBuf     [256]byte
	testPrintfS1 = [...]byte{'X', 'Y', 0}
)

// func TestPrintf(t *testing.T) {
// 	tls := rtl.NewTLS()
//
// 	defer tls.Close()
//
// 	i := uint64(0x123456789abcdef)
// 	j := uint64(0xf123456789abcde)
// 	k := uint64(0x23456789abcdef1)
// 	l := uint64(0xef123456789abcd)
// 	for itest, test := range []struct {
// 		fmt    string
// 		args   []interface{}
// 		result string
// 	}{
// 		{
// 			"%I64x %I32x %I64x %I32x",
// 			[]interface{}{int64(i), int32(j), int64(k), int32(l)},
// 			"123456789abcdef 789abcde 23456789abcdef1 6789abcd",
// 		},
// 		{
// 			"%llx %x %llx %x",
// 			[]interface{}{int64(i), int32(j), int64(k), int32(l)},
// 			"123456789abcdef 789abcde 23456789abcdef1 6789abcd",
// 		},
// 		{
// 			"%.1s\n",
// 			[]interface{}{uintptr(unsafe.Pointer(&testPrintfS1[0]))},
// 			"X\n",
// 		},
// 		{
// 			"%.2s\n",
// 			[]interface{}{uintptr(unsafe.Pointer(&testPrintfS1[0]))},
// 			"XY\n",
// 		},
// 	} {
// 		copy(formatString[:], test.fmt+"\x00")
// 		printBuf = [256]byte{}
// 		x_sprintf(tls, uintptr(unsafe.Pointer(&printBuf)), uintptr(unsafe.Pointer(&formatString[0])), rtl.VaList(uintptr(unsafe.Pointer(&valist[0])), test.args...))
// 		x := bytes.IndexByte(printBuf[:], 0)
// 		if x < 0 {
// 			t.Errorf("%v:", itest)
// 			continue
// 		}
//
// 		b := printBuf[:x]
// 		if g, e := string(b), test.result; g != e {
// 			t.Errorf("%v: %q %q", itest, g, e)
// 		}
// 	}
// }
