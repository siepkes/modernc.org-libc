// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
	"unsafe"

	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
)

func TestMain(m *testing.M) {
	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	rc := m.Run()
	os.Exit(rc)
}

var (
	goarch = runtime.GOARCH
	goos   = runtime.GOOS

	testAtomicCASInt32  int32
	testAtomicCASUint64 uint64
	testAtomicCASp      uintptr
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

	a_or(pi, Int32FromUint32(0x80000000))
	if j := testAtomicCASInt32; j != Int32FromUint32(0x80000003) {
		t.Fatalf("%032b", j)
	}
}

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

const testGetentropySize = 100

var testGetentropyBuf [testGetentropySize]byte

func TestGetentropy(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	Xgetentropy(tls, uintptr(unsafe.Pointer(&testGetentropyBuf[0])), testGetentropySize)
	t.Logf("\n%s", hex.Dump(testGetentropyBuf[:]))
}

func TestReallocArray(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	const size = 16
	p := Xmalloc(tls, size)
	if p == 0 {
		t.Fatal()
	}

	for i := 0; i < size; i++ {
		unsafe.Slice((*byte)(unsafe.Pointer(p)), size)[i] = byte(i ^ 0x55)
	}

	q := Xreallocarray(tls, p, 2, size)
	if q == 0 {
		t.Fatal()
	}

	defer Xfree(tls, q)

	for i := 0; i < size; i++ {
		if g, e := unsafe.Slice((*byte)(unsafe.Pointer(q)), size)[i], byte(i^0x55); g != e {
			t.Fatal(i, g, e)
		}
	}
}

func mustCString(s string) (r uintptr) {
	n := len(s)
	r = mustMalloc(Tsize_t(n + 1))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(r)), n), s)
	*(*byte)(unsafe.Pointer(r + uintptr(n))) = 0
	return r
}

var testSnprintfBuf [3]byte

func TestSnprintf(t *testing.T) {
	tls := NewTLS()

	defer tls.Close()

	testSnprintfBuf = [3]byte{0xff, 0xff, 0xff}
	p := uintptr(unsafe.Pointer(&testSnprintfBuf[0]))
	s := mustCString("12")
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

	p := Xfdopen(tls, int32(f.Fd()), mustCString("r"))

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

var testAtomicStore8 uint32

func TestAtomicStore8(t *testing.T) {
	testAtomicStore8 = 0x12345678
	AtomicStorePUint8(uintptr(unsafe.Pointer(&testAtomicStore8)), 0x9a)
	if g, e := *(*byte)(unsafe.Pointer(&testAtomicStore8)), byte(0x9a); g != e {
		t.Errorf("%#0x %#0x", g, e)
	}

	if g, e := testAtomicStore8, uint32(0x1234569a); g != e {
		t.Errorf("%#0x %#0x", g, e)
	}
}

func TestMemAuditBrk(t *testing.T) {
	if !isMemBrk {
		t.Skip("requires -tags=libc.membrk")
	}

	a := allocator

	defer func() { allocator = a }()

	mallocP := mustMalloc(1)
	t.Logf("mallocP %v %#0[1]x", mallocP)
	t.Logf("\n%s", hex.Dump(unsafe.Slice((*byte)(unsafe.Pointer(mallocP-heapGuard)), 4*heapGuard)))
	q := mallocP - heapGuard
	c := 0
	for ; q < mallocP; q++ {
		*(*byte)(unsafe.Pointer(q)) ^= 0x55
		c++
	}

	z := roundup(mallocP+1, heapAlign)
	for p := mallocP + 1; p < z; p++ {
		*(*byte)(unsafe.Pointer(p)) ^= 0x55
	}
	p := z
	z += heapGuard
	for ; p < z; p++ {
		*(*byte)(unsafe.Pointer(p)) ^= 0x55
		c++
	}
	p = mallocP + 2*heapGuard + 7
	*(*byte)(unsafe.Pointer(p)) ^= 0x55
	c++
	t.Logf("c %v, \n%s", c, hex.Dump(unsafe.Slice((*byte)(unsafe.Pointer(mallocP-heapGuard)), 4*heapGuard)))
	r := MemAudit()
	for i, v := range r {
		t.Log(i, v)
	}
	if g, e := len(r), c; g != e {
		t.Fatalf("got %v errors, expected %v", g, e)
	}
}

func mustShell(t *testing.T, cmd string, args ...string) []byte {
	r, err := shell(cmd, args...)
	if err != nil {
		t.Fatal(err)
	}

	return r
}

func shell(cmd string, args ...string) ([]byte, error) {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return nil, err
	}

	return exec.Command(cmd, args...).CombinedOutput()
}

var skipFiles = map[string]struct{}{
	// n/a
	"internal/nsz.repo.hu/libc-test/src/common/runtest.exe.go": {}, // src/common/runtest.c:37: usage: /tmp/TestLibc748082357/001/main [-t timeoutsec] [-w wrapcmd] cmd [args..]

	// fork
	"internal/nsz.repo.hu/libc-test/src/functional/ipc_msg.exe.go":        {}, // src/functional/ipc_msg.c:122: fork failed: Function not implemented
	"internal/nsz.repo.hu/libc-test/src/functional/ipc_sem.exe.go":        {}, // src/functional/ipc_sem.c:115: fork failed: Function not implemented
	"internal/nsz.repo.hu/libc-test/src/functional/ipc_shm.exe.go":        {}, // src/functional/ipc_shm.c:112: fork failed: Function not implemented
	"internal/nsz.repo.hu/libc-test/src/regression/daemon-failure.exe.go": {}, // src/regression/daemon-failure.c:35: fork failed: Function not implemented
	"internal/nsz.repo.hu/libc-test/src/regression/fflush-exit.exe.go":    {},

	// vfork
	"internal/nsz.repo.hu/libc-test/src/functional/vfork.exe.go": {},

	// clone
	"internal/nsz.repo.hu/libc-test/src/functional/popen.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/functional/spawn.exe.go": {},

	// dlopen
	"internal/nsz.repo.hu/libc-test/src/functional/tls_align_dlopen.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/functional/tls_init_dlopen.exe.go":  {},

	// vmfill
	"internal/nsz.repo.hu/libc-test/src/regression/malloc-brk-fail.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/regression/setenv-oom.exe.go":      {},

	// cbin fails too
	"internal/nsz.repo.hu/libc-test/src/functional/clocale_mbfuncs.exe.go": {},

	//TODO other
	"internal/nsz.repo.hu/libc-test/src/regression/sigaltstack.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/regression/sigreturn.exe.go":   {},

	//TODO miscompiles
	"internal/nsz.repo.hu/libc-test/src/functional/setjmp.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/functional/sscanf.exe.go": {},

	//TODO asm
	"internal/nsz.repo.hu/libc-test/src/functional/tgmath.exe.go": {},

	//TODO fpclassify
	"internal/nsz.repo.hu/libc-test/src/math/fpclassify.exe.go": {},

	//TODO pthread
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_cancel-points.exe.go":            {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_cancel.exe.go":                   {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_cond.exe.go":                     {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_mutex.exe.go":                    {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_mutex_pi.exe.go":                 {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_robust.exe.go":                   {},
	"internal/nsz.repo.hu/libc-test/src/functional/pthread_tsd.exe.go":                      {},
	"internal/nsz.repo.hu/libc-test/src/functional/sem_init.exe.go":                         {},
	"internal/nsz.repo.hu/libc-test/src/functional/tls_init.exe.go":                         {},
	"internal/nsz.repo.hu/libc-test/src/functional/tls_local_exec.exe.go":                   {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread-robust-detach.exe.go":            {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_atfork-errno-clobber.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_cancel-sem_wait.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_cond-smasher.exe.go":             {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_cond_wait-cancel_ignored.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_condattr_setclock.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_create-oom.exe.go":               {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_exit-cancel.exe.go":              {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_exit-dtor.exe.go":                {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_once-deadlock.exe.go":            {},
	"internal/nsz.repo.hu/libc-test/src/regression/pthread_rwlock-ebusy.exe.go":             {},
	"internal/nsz.repo.hu/libc-test/src/regression/raise-race.exe.go":                       {},
	"internal/nsz.repo.hu/libc-test/src/regression/tls_get_new-dtv.exe.go":                  {},

	//TODO fesetenv and friends
	"internal/nsz.repo.hu/libc-test/src/math/acos.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/acosf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/acosh.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/acoshf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/acoshl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/acosl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/asin.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/asinf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/asinh.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/asinhf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/asinhl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/asinl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/atan.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/atan2.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/atan2f.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/atan2l.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/atanf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/atanh.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/atanhf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/atanhl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/atanl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/cbrt.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/cbrtf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/cbrtl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/ceil.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/ceilf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/ceill.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/copysign.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/copysignf.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/copysignl.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/cos.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/cosf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/cosh.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/coshf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/coshl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/cosl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/drem.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/dremf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/erf.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/erfc.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/erfcf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/erfcl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/erff.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/erfl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/exp.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/exp10.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/exp10f.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/exp10l.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/exp2.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/exp2f.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/exp2l.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/expf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/expl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/expm1.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/expm1f.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/expm1l.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/fabs.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fabsf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fabsl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fdim.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fdimf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fdiml.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fenv.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/floor.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/floorf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/floorl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/fma.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/fmaf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fmal.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fmax.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fmaxf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fmaxl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fmin.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fminf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fminl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fmod.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/fmodf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/fmodl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/frexp.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/frexpf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/frexpl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/hypot.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/hypotf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/hypotl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/ilogb.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/ilogbf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/ilogbl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/isless.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/j0.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/j0f.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/j1.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/j1f.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/jn.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/jnf.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/ldexp.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/ldexpf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/ldexpl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/lgamma.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/lgamma_r.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/lgammaf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/lgammaf_r.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/lgammal.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/lgammal_r.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/llrint.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/llrintf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/llrintl.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/llround.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/llroundf.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/llroundl.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/log.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/log10.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/log10f.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/log10l.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/log1p.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/log1pf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/log1pl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/log2.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/log2f.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/log2l.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/logb.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/logbf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/logbl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/logf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/logl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/lrint.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/lrintf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/lrintl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/lround.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/lroundf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/lroundl.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/modf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/modff.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/modfl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/nearbyint.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/nearbyintf.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/nearbyintl.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/nextafter.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/nextafterf.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/nextafterl.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/nexttoward.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/nexttowardf.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/math/nexttowardl.exe.go": {},
	"internal/nsz.repo.hu/libc-test/src/math/pow.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/powf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/powl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/remainder.exe.go":   {},
	"internal/nsz.repo.hu/libc-test/src/math/remainderf.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/remainderl.exe.go":  {},
	"internal/nsz.repo.hu/libc-test/src/math/remquo.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/remquof.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/remquol.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/rint.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/rintf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/rintl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/round.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/roundf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/roundl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/scalb.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/scalbf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/scalbln.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/scalblnf.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/scalblnl.exe.go":    {},
	"internal/nsz.repo.hu/libc-test/src/math/scalbn.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/scalbnf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/scalbnl.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/sin.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/sincos.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/sincosf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/sincosl.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/sinf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/sinh.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/sinhf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/sinhl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/sinl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/sqrt.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/sqrtf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/sqrtl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/tan.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/tanf.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/tanh.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/tanhf.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/tanhl.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/tanl.exe.go":        {},
	"internal/nsz.repo.hu/libc-test/src/math/tgamma.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/tgammaf.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/tgammal.exe.go":     {},
	"internal/nsz.repo.hu/libc-test/src/math/trunc.exe.go":       {},
	"internal/nsz.repo.hu/libc-test/src/math/truncf.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/truncl.exe.go":      {},
	"internal/nsz.repo.hu/libc-test/src/math/y0.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/y0f.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/y1.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/y1f.exe.go":         {},
	"internal/nsz.repo.hu/libc-test/src/math/yn.exe.go":          {},
	"internal/nsz.repo.hu/libc-test/src/math/ynf.exe.go":         {},
}

func TestLibc(t *testing.T) {
	if testing.Short() {
		t.Skip("-short")
	}

	dir := filepath.Join("internal", "nsz.repo.hu", "libc-test")

	util.InDir(dir, func() error {
		mustShell(t, "make", "cleanall")
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"--prefix-field=F",
				"-absolute-paths",
				"-extended-errors",
				"-positions",
				"-exec", "make",
			},
			os.Stdout, os.Stderr,
			nil,
		).Exec(); err != nil {
			t.Fatal(err)
		}

		return nil
	})

	bin := filepath.Join(t.TempDir(), "main")
	var file, skip, buildfail, buildok, cfail, fail, pass int
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			t.Fatal(err)
		}

		if info.IsDir() || !strings.HasSuffix(path, ".exe.go") {
			return nil
		}

		file++
		if _, ok := skipFiles[filepath.ToSlash(path)]; ok {
			skip++
			return nil
		}

		cbin := path[:len(path)-len(".go")]
		_, cerr := run(t, cbin)
		if cerr != nil {
			t.Logf("%s: cbin fail=%s", path, cerr)
			cfail++
		}
		os.Remove(bin)
		if out, err := shell("go", "build", "-o", bin, path); err != nil {
			buildfail++
			t.Errorf("%s %s: BUILD FAIL=%s", out, path, err)
			return nil
		}

		buildok++
		switch out, err := run(t, bin); {
		case err != nil:
			fail++
			t.Errorf("%s %s: EXEC FAIL=%s", out, path, err)
		default:
			pass++
			t.Logf("%s: PASS", path)
		}
		return nil
	})
	t.Logf("files=%v skip=%v cfail=%v buildfail=%v buildok=%v fail=%v pass=%v", file, skip, cfail, buildfail, buildok, fail, pass)
	// all_linux_amd64_test.go:814: files=334 skip=232 cfail=16 buildfail=0 buildok=102 fail=6 pass=96
}

func run(t *testing.T, bin string) (out []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cmd := exec.CommandContext(ctx, bin)
	cmd.WaitDelay = 20 * time.Second
	return cmd.CombinedOutput()
}
