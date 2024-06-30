// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc"

import (
	"encoding/hex"
	"testing"
	"unsafe"
)

func TestXfmod(t *testing.T) {
	x := 1.3518643030646695
	y := 6.283185307179586
	if g, e := Xfmod(nil, x, y), 1.3518643030646695; g != e {
		t.Fatal(g, e)
	}
}

func TestSwap(t *testing.T) {
	if g, e := X__builtin_bswap16(nil, 0x1234), uint16(0x3412); g != e {
		t.Errorf("%#04x %#04x", g, e)
	}
	if g, e := X__builtin_bswap32(nil, 0x12345678), uint32(0x78563412); g != e {
		t.Errorf("%#04x %#04x", g, e)
	}
	if g, e := X__builtin_bswap64(nil, 0x123456789abcdef0), uint64(0xf0debc9a78563412); g != e {
		t.Errorf("%#04x %#04x", g, e)
	}
}

var (
	valist       [256]byte
	formatString [256]byte
	srcString    [256]byte
	testPrintfS1 = [...]byte{'X', 'Y', 0}
)

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
			Xmemset(nil, uintptr(unsafe.Pointer(&testMemsetBuf[start])), int32(v), Tsize_t(n))
			if testMemsetBuf != e {
				t.Fatalf("start %v, v %#x n %v, exp\n%s\ngot\n%s", start, byte(v), n, hex.Dump(e[:]), hex.Dump(testMemsetBuf[:]))
			}
		}
	}
}

func TestABI(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	t.Run("div", func(t *testing.T) { testABIdiv(t, tls) })
	t.Run("ldiv", func(t *testing.T) { testABIldiv(t, tls) })
}

func testABIdiv(t *testing.T, tls *TLS) {
	if g, e := Xdiv(tls, 31, 5), (Tdiv_t{6, 1}); g != e {
		t.Error(g, e)
	}
}

func testABIldiv(t *testing.T, tls *TLS) {
	if g, e := Xldiv(tls, 31, 5), (Tldiv_t{6, 1}); g != e {
		t.Error(g, e)
	}
}
