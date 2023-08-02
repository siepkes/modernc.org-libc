// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

var (
	testAtomicCASInt32  int32
	testAtomicCASUint64 uint64
	testAtomicCASp      uintptr
)

// func TestAtomicCASInt32(t *testing.T) {
// 	pi := uintptr(unsafe.Pointer(&testAtomicCASInt32))
// 	testAtomicCASInt32 = int32(0)
// 	j := a_cas(pi, 1, 2)
// 	if testAtomicCASInt32 != 0 || j != 0 {
// 		t.Fatal(testAtomicCASInt32, j)
// 	}
//
// 	if j = a_cas(pi, 0, 3); testAtomicCASInt32 != 3 || j != 0 {
// 		t.Fatal(testAtomicCASInt32, j)
// 	}
//
// 	if j = a_cas(pi, 4, 5); testAtomicCASInt32 != 3 || j != 3 {
// 		t.Fatal(testAtomicCASInt32, j)
// 	}
//
// 	if j = a_cas(pi, 3, 6); testAtomicCASInt32 != 6 || j != 3 {
// 		t.Fatal(testAtomicCASInt32, j)
// 	}
// }

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

// func TestAtomicOrInt32(t *testing.T) {
// 	pi := uintptr(unsafe.Pointer(&testAtomicCASInt32))
// 	testAtomicCASInt32 = int32(0)
// 	a_or(pi, 1)
// 	if j := testAtomicCASInt32; j != 1 {
// 		t.Fatalf("%032b", j)
// 	}
//
// 	a_or(pi, 2)
// 	if j := testAtomicCASInt32; j != 3 {
// 		t.Fatalf("%032b", j)
// 	}
//
// 	a_or(pi, Int32FromUint32(0x80000000))
// 	if j := testAtomicCASInt32; j != Int32FromUint32(0x80000003) {
// 		t.Fatalf("%032b", j)
// 	}
// }

func TestAtomicOrUint64(t *testing.T) {
	pi := uintptr(unsafe.Pointer(&testAtomicCASUint64))
	testAtomicCASUint64 = 0
	a_or_64(pi, 1)
	if j := testAtomicCASUint64; j != 1 {
		t.Fatalf("%064b", j)
	}

	a_or_64(pi, 2)
	if j := testAtomicCASUint64; j != 3 {
		t.Fatalf("%064b", j)
	}

	a_or_64(pi, Uint64FromUint64(0x8000000000000000))
	if j := testAtomicCASUint64; j != Uint64FromUint64(0x8000000000000003) {
		t.Fatalf("%064b", j)
	}
}

func TestXfmod(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	x := 1.3518643030646695
	y := 6.283185307179586
	if g, e := Xfmod(tls, x, y), 1.3518643030646695; g != e {
		t.Fatal(g, e)
	}
}

// func TestSwap(t *testing.T) {
// 	if g, e := ___bswap_16(nil, 0x1234), uint16(0x3412); g != e {
// 		t.Errorf("%#04x %#04x", g, e)
// 	}
// 	if g, e := ___bswap_32(nil, 0x12345678), uint32(0x78563412); g != e {
// 		t.Errorf("%#04x %#04x", g, e)
// 	}
// }

var (
	valist       [256]byte
	formatString [256]byte
	srcString    [256]byte
	printBuf     [256]byte
	testPrintfS1 = [...]byte{'X', 'Y', 0}
)

func TestSprintf(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	i := uint64(0x123456789abcdef)
	j := uint64(0xf123456789abcde)
	k := uint64(0x23456789abcdef1)
	l := uint64(0xef123456789abcd)
	for itest, test := range []struct {
		fmt    string
		args   []interface{}
		result string
	}{
		//TODO windows specific {
		//TODO windows specific 	"%I64x %I32x %I64x %I32x",
		//TODO windows specific 	[]interface{}{int64(i), int32(j), int64(k), int32(l)},
		//TODO windows specific 	"123456789abcdef 789abcde 23456789abcdef1 6789abcd",
		//TODO windows specific },

		{
			"%llx %x %llx %x",
			[]interface{}{int64(i), int32(j), int64(k), int32(l)},
			"123456789abcdef 789abcde 23456789abcdef1 6789abcd",
		},
		{
			"%.1s\n",
			[]interface{}{uintptr(unsafe.Pointer(&testPrintfS1[0]))},
			"X\n",
		},
		{
			"%.2s\n",
			[]interface{}{uintptr(unsafe.Pointer(&testPrintfS1[0]))},
			"XY\n",
		},
	} {
		copy(formatString[:], test.fmt+"\x00")
		printBuf = [256]byte{}
		rc := Xsprintf(tls, uintptr(unsafe.Pointer(&printBuf)), uintptr(unsafe.Pointer(&formatString[0])), VaList(uintptr(unsafe.Pointer(&valist[0])), test.args...))
		x := bytes.IndexByte(printBuf[:], 0)
		if x < 0 {
			t.Errorf("%v:", itest)
			continue
		}

		b := printBuf[:x]
		if g, e := string(b), test.result; g != e {
			t.Errorf("%v: %q %q, rc %v", itest, g, e, rc)
		}
	}
}

func TestStrtod(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	for itest, test := range []struct {
		s      string
		result float64
	}{
		{"+0", 0},
		{"+1", 1},
		{"+2", 2},
		{"-0", 0},
		{"-1", -1},
		{"-2", -2},
		{".5", .5},
		{"0", 0},
		{"1", 1},
		{"1.", 1},
		{"1.024e3", 1024},
		{"16", 16},
		{"2", 2},
		{"32", 32},
	} {
		copy(srcString[:], test.s+"\x00")
		if g, e := Xstrtod(tls, uintptr(unsafe.Pointer(&srcString[0])), 0), test.result; g != e {
			t.Errorf("%v: %q: %v %v", itest, test.s, g, e)
		}
	}
}

func TestRint(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	for itest, test := range []struct {
		x, y float64
	}{
		{-1.1, -1.0},
		{-1.0, -1.0},
		{-0.9, -1.0},
		{-0.51, -1.0},
		{-0.49, 0},
		{-0.1, 0},
		{-0, 0},
		{0.1, 0},
		{0.49, 0},
		{0.51, 1},
		{0.9, 1},
		{1, 1},
		{1.1, 1},
	} {
		if g, e := Xrint(tls, test.x), test.y; g != e {
			t.Errorf("#%d: x %v, got %v, expected %v", itest, test.x, g, e)
		}
	}
}

var testMemsetBuf [67]byte

func TestMemset(t *testing.T) {
	v := 0
	for start := 0; start < len(testMemsetBuf); start++ {
		for n := 0; n < len(testMemsetBuf)-start; n++ {
			for x := range testMemsetBuf {
				testMemsetBuf[x] = byte(v)
				v++
			}
			for x := start; x < start+n; x++ {
				testMemsetBuf[x] = byte(v)
			}
			e := testMemsetBuf
			Xmemset(nil, uintptr(unsafe.Pointer(&testMemsetBuf[start])), int32(v), uint64(n))
			if testMemsetBuf != e {
				t.Fatalf("start %v, v %#x n %v, exp\n%s\ngot\n%s", start, byte(v), n, hex.Dump(e[:]), hex.Dump(testMemsetBuf[:]))
			}
		}
	}
}

//TODO N/A musl 0.6.0
//TODO const testGetentropySize = 100
//TODO
//TODO var testGetentropyBuf [testGetentropySize]byte
//TODO
//TODO func TestGetentropy(t *testing.T) {
//TODO 	Xgetentropy(tls, uintptr(unsafe.Pointer(&testGetentropyBuf[0])), testGetentropySize)
//TODO 	t.Logf("\n%s", hex.Dump(testGetentropyBuf[:]))
//TODO }

//TODO N/A musl 0.6.0
// TODO func TestReallocArray(t *testing.T) {
// TODO 	const size = 16
// TODO 	p := Xmalloc(tls, size)
// TODO 	if p == 0 {
// TODO 		t.Fatal()
// TODO 	}
// TODO
// TODO 	//TODO defer Xfree(tls, p), crashes
// TODO
// TODO 	for i := 0; i < size; i++ {
// TODO 		unsafe.Slice((*byte)(unsafe.Pointer(p)), size)[i] = byte(i ^ 0x55)
// TODO 	}
// TODO
// TODO 	q := Xreallocarray(tls, p, 2, size)
// TODO 	if q == 0 {
// TODO 		t.Fatal()
// TODO 	}
// TODO
// TODO 	for i := 0; i < size; i++ {
// TODO 		if g, e := unsafe.Slice((*byte)(unsafe.Pointer(p)), size)[i], byte(i^0x55); g != e {
// TODO 			t.Fatal(i, g, e)
// TODO 		}
// TODO 	}
// TODO }

func mustTestCString(s string) uintptr {
	r, err := testCString(s)
	if err != nil {
		panic("testCString failed")
	}

	return r
}

func testCString(s string) (uintptr, error) {
	n := len(s)
	p, err := privateMalloc(n + 1)
	if err != nil {
		return 0, err
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), n), s)
	*(*byte)(unsafe.Pointer(p + uintptr(n))) = 0
	return p, nil
}

var testSnprintfBuf [3]byte

func TestSnprintf(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	testSnprintfBuf = [3]byte{0xff, 0xff, 0xff}
	p := uintptr(unsafe.Pointer(&testSnprintfBuf[0]))
	s := mustTestCString("12")
	if g, e := Xsnprintf(tls, p, 1, s, 0), int32(2); g != e {
		t.Fatal(g, e)
	}

	if g, e := testSnprintfBuf, [3]byte{0x00, 0xff, 0xff}; g != e {
		t.Fatal(g, e)
	}

	testSnprintfBuf = [3]byte{0xff, 0xff, 0xff}
	if g, e := Xsnprintf(tls, p, 2, s, 0), int32(2); g != e {
		t.Fatal(g, e)
	}

	if g, e := testSnprintfBuf, [3]byte{'1', 0x00, 0xff}; g != e {
		t.Fatal(g, e)
	}

	testSnprintfBuf = [3]byte{0xff, 0xff, 0xff}
	if g, e := Xsnprintf(tls, p, 3, s, 0), int32(2); g != e {
		t.Fatal(g, e)
	}

	if g, e := testSnprintfBuf, [3]byte{'1', '2', 0x00}; g != e {
		t.Fatal(g, e)
	}
}

var testFdopenBuf [100]byte

func TestFdopen(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	const s = "foobarbaz\n"
	tempdir := t.TempDir()
	f, err := os.Create(filepath.Join(tempdir, "test_fdopen"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := f.Write([]byte(s)); err != nil {
		t.Fatal(err)
	}

	if _, err := f.Seek(0, os.SEEK_SET); err != nil {
		t.Fatal(err)
	}

	p := Xfdopen(tls, int32(f.Fd()), mustTestCString("r"))

	bp := uintptr(unsafe.Pointer(&testFdopenBuf))
	if g, e := Xfread(tls, bp, 1, uint64(len(testFdopenBuf)), p), uint64(len(s)); g != e {
		t.Fatal(g, e)
	}

	if g, e := string(GoBytes(bp, len(s))), s; g != e {
		t.Fatalf("%q %q", g, e)
	}
}

func TestPow(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	for itest, test := range []struct{ x, y, z float64 }{
		{2, 12, 4096},
	} {
		if g, e := Xpow(tls, test.x, test.y), test.z; g != e {
			t.Errorf("%d: %v %v %v, %v", itest, test.x, test.y, test.z, g)
		}
	}
}
