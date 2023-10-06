// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/isless.exe.go src/math/isless.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

//go:build linux && 386
// +build linux,386

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc/v2"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:8:9:
const FE_ALL_EXCEPT = 63

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:6:9:
const FE_INEXACT = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:1:9:
const FE_INVALID = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:10:9:
const FE_TONEAREST = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:5:9:
const FE_UNDERFLOW = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:33:9:
const FP_NAN = 0

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:7:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:7:14:
type __builtin_va_list = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:23:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:23:23:
type __predefined_size_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:27:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:27:24:
type __predefined_wchar_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:31:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/<builtin>:31:26:
type __predefined_ptrdiff_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:38:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:38:21:
type float_t = float64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:43:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:43:21:
type double_t = float64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:15:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:15:24:
type fexcept_t = uint16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:17:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:31:3:
type fenv_t = struct {
	F__control_word  uint16
	F__unused1       uint16
	F__status_word   uint16
	F__unused2       uint16
	F__tags          uint16
	F__unused3       uint16
	F__eip           uint32
	F__cs_selector   uint16
	F__ccgo_align8   [2]byte
	F__ccgo20        uint16
	F__data_offset   uint32
	F__data_selector uint16
	F__unused5       uint16
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:69:1:
type d_d = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:70:1:
type f_f = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fy    float32
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:71:1:
type l_l = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:72:1:
type ff_f = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fx2   float32
	Fy    float32
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:73:1:
type dd_d = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:74:1:
type ll_l = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:75:1:
type d_di = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:76:1:
type f_fi = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fy    float32
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:77:1:
type l_li = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:78:1:
type di_d = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fi    int64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:79:1:
type fi_f = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fi    int64
	Fy    float32
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:80:1:
type li_l = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fi    int64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:81:1:
type d_i = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:82:1:
type f_i = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:83:1:
type l_i = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:84:1:
type d_dd = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fy2   float64
	Fdy2  float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:85:1:
type f_ff = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fy    float32
	Fdy   float32
	Fy2   float32
	Fdy2  float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:86:1:
type l_ll = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fy    float64
	Fdy   float32
	Fy2   float64
	Fdy2  float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:87:1:
type ff_fi = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fx2   float32
	Fy    float32
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:88:1:
type dd_di = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fy    float64
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:89:1:
type ll_li = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fy    float64
	Fdy   float32
	Fi    int64
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:90:1:
type fff_f = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float32
	Fx2   float32
	Fx3   float32
	Fy    float32
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:91:1:
type ddd_d = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fx3   float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:92:1:
type lll_l = struct {
	Ffile uintptr
	Fline int32
	Fr    int32
	Fx    float64
	Fx2   float64
	Fx3   float64
	Fy    float64
	Fdy   float32
	Fe    int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:102:12:
func checkexcept(tls *libc.TLS, got int32, want int32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:103:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:104:2:
	if r == FE_TONEAREST {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:110:3:
		return libc.BoolInt32(got|int32(FE_INEXACT) == want|int32(FE_INEXACT))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:112:2:
	return libc.BoolInt32(got|int32(FE_INEXACT)|int32(FE_UNDERFLOW) == want|int32(FE_INEXACT)|int32(FE_UNDERFLOW))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:115:12:
func checkexceptall(tls *libc.TLS, got int32, want int32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:116:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:117:2:
	return libc.BoolInt32(got == want)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:120:12:
func checkulp(tls *libc.TLS, d float32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:123:2:
	if r == FE_TONEAREST {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:124:3:
		return libc.BoolInt32(float64(libc.Xfabsf(tls, d)) < float64(1.5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:126:2:
	return libc.BoolInt32(float64(libc.Xfabsf(tls, d)) < float64(3))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:129:12:
func checkcr(tls *libc.TLS, y float64, ywant float64, r int32) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:130:1:
	var v1, v3, v5, v7 uint64
	var v9 bool
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	_, _, _, _, _ = v1, v3, v5, v7, v9
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:131:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp)) = float64(ywant)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:132:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(y)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
		return libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:133:2:
	if v9 = y == ywant; v9 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(y)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v5 = *(*uint64)(unsafe.Pointer(bp))
		goto _6
	_6:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(ywant)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v7 = *(*uint64)(unsafe.Pointer(bp))
		goto _8
	_8:
	}
	return libc.BoolInt32(v9 && int32(v5>>libc.Int32FromInt32(63)) == int32(v7>>libc.Int32FromInt32(63)))
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:78:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:78:24:
type uintptr_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:93:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:93:15:
type intptr_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:119:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:119:25:
type int8_t = int8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:124:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:124:25:
type int16_t = int16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:129:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:129:25:
type int32_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:134:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:134:25:
type int64_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:139:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:139:25:
type intmax_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:144:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:144:25:
type uint8_t = uint8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:149:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:149:25:
type uint16_t = uint16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:154:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:154:25:
type uint32_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:159:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:159:25:
type uint64_t = uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:169:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:169:25:
type uintmax_t = uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:22:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:22:16:
type int_fast8_t = int8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:23:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:23:17:
type int_fast64_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:25:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:25:17:
type int_least8_t = int8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:26:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:26:17:
type int_least16_t = int16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:27:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:27:17:
type int_least32_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:28:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:28:17:
type int_least64_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:30:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:30:17:
type uint_fast8_t = uint8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:31:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:31:18:
type uint_fast64_t = uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:33:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:33:18:
type uint_least8_t = uint8

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:34:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:34:18:
type uint_least16_t = uint16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:35:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:35:18:
type uint_least32_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:36:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:36:18:
type uint_least64_t = uint64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:1:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:1:17:
type int_fast16_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:2:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:2:17:
type int_fast32_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:3:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:3:18:
type uint_fast16_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:4:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/stdint.h:4:18:
type uint_fast32_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:73:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:73:24:
type size_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:88:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:88:15:
type ssize_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:185:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:185:16:
type off_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:13:
type pid_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:268:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:268:18:
type uid_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:273:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:273:18:
type gid_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:18:
type useconds_t = uint32

const
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:5:7:
LESS = 0
const
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:5:12:
EQUAL = 1
const
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:5:18:
GREATER = 2
const
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:5:26:
UNORD = 3

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:28:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r264 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:29:1:
	var e, e1, e10, e100, e101, e102, e103, e104, e105, e106, e107, e108, e109, e11, e110, e111, e112, e113, e114, e115, e116, e117, e118, e119, e12, e120, e121, e122, e123, e124, e125, e126, e127, e128, e129, e13, e130, e131, e132, e133, e134, e135, e136, e137, e138, e139, e14, e140, e141, e142, e143, e144, e145, e146, e147, e148, e149, e15, e150, e151, e152, e153, e154, e155, e156, e157, e158, e159, e16, e160, e161, e162, e163, e164, e165, e166, e167, e168, e169, e17, e170, e171, e172, e173, e174, e175, e176, e177, e178, e179, e18, e180, e181, e182, e183, e184, e185, e186, e187, e188, e189, e19, e190, e191, e192, e193, e194, e195, e196, e197, e198, e199, e2, e20, e200, e201, e202, e203, e204, e205, e206, e207, e208, e209, e21, e210, e211, e212, e213, e214, e215, e216, e217, e218, e219, e22, e220, e221, e222, e223, e224, e225, e226, e227, e228, e229, e23, e230, e231, e232, e233, e234, e235, e236, e237, e238, e239, e24, e240, e241, e242, e243, e244, e245, e246, e247, e248, e249, e25, e250, e251, e252, e253, e254, e255, e256, e257, e258, e259, e26, e260, e261, e262, e263, e27, e28, e29, e3, e30, e31, e32, e33, e34, e35, e36, e37, e38, e39, e4, e40, e41, e42, e43, e44, e45, e46, e47, e48, e49, e5, e50, e51, e52, e53, e54, e55, e56, e57, e58, e59, e6, e60, e61, e62, e63, e64, e65, e66, e67, e68, e69, e7, e70, e71, e72, e73, e74, e75, e76, e77, e78, e79, e8, e80, e81, e82, e83, e84, e85, e86, e87, e88, e89, e9, e90, e91, e92, e93, e94, e95, e96, e97, e98, e99, r, r1, r10, r100, r101, r102, r103, r104, r105, r106, r107, r108, r109, r11, r110, r111, r112, r113, r114, r115, r116, r117, r118, r119, r12, r120, r121, r122, r123, r124, r125, r126, r127, r128, r129, r13, r130, r131, r132, r133, r134, r135, r136, r137, r138, r139, r14, r140, r141, r142, r143, r144, r145, r146, r147, r148, r149, r15, r150, r151, r152, r153, r154, r155, r156, r157, r158, r159, r16, r160, r161, r162, r163, r164, r165, r166, r167, r168, r169, r17, r170, r171, r172, r173, r174, r175, r176, r177, r178, r179, r18, r180, r181, r182, r183, r184, r185, r186, r187, r188, r189, r19, r190, r191, r192, r193, r194, r195, r196, r197, r198, r199, r2, r20, r200, r201, r202, r203, r204, r205, r206, r207, r208, r209, r21, r210, r211, r212, r213, r214, r215, r216, r217, r218, r219, r22, r220, r221, r222, r223, r224, r225, r226, r227, r228, r229, r23, r230, r231, r232, r233, r234, r235, r236, r237, r238, r239, r24, r240, r241, r242, r243, r244, r245, r246, r247, r248, r249, r25, r250, r251, r252, r253, r254, r255, r256, r257, r258, r259, r26, r260, r261, r262, r263, r27, r28, r29, r3, r30, r31, r32, r33, r34, r35, r36, r37, r38, r39, r4, r40, r41, r42, r43, r44, r45, r46, r47, r48, r49, r5, r50, r51, r52, r53, r54, r55, r56, r57, r58, r59, r6, r60, r61, r62, r63, r64, r65, r66, r67, r68, r69, r7, r70, r71, r72, r73, r74, r75, r76, r77, r78, r79, r8, r80, r81, r82, r83, r84, r85, r86, r87, r88, r89, r9, r90, r91, r92, r93, r94, r95, r96, r97, r98, r99, v1, v10, v100, v1001, v1008, v1010, v1017, v1019, v102, v1026, v1028, v1033, v1040, v1042, v1049, v1051, v1058, v1060, v1067, v1069, v107, v1076, v1078, v1083, v1090, v1092, v1099, v110, v1101, v1108, v1110, v1117, v1119, v112, v1126, v1128, v1133, v1140, v1142, v1149, v1151, v1158, v1160, v1167, v1169, v1176, v1178, v1183, v119, v1190, v1192, v1199, v1201, v1208, v121, v1210, v1217, v1219, v1226, v1228, v1233, v1240, v1242, v1249, v1251, v1258, v126, v1260, v1267, v1269, v1276, v1278, v1283, v129, v1290, v1292, v1299, v13, v1301, v1308, v131, v1310, v1317, v1319, v1326, v1328, v1333, v1340, v1342, v1349, v1351, v1358, v136, v1360, v1367, v1369, v1376, v1378, v1383, v139, v1390, v1392, v1399, v1401, v1408, v141, v1410, v1417, v1419, v1426, v1428, v1433, v1440, v1442, v1449, v1451, v1458, v1460, v1467, v1469, v1476, v1478, v148, v1483, v1490, v1492, v1499, v15, v150, v1501, v1508, v1510, v1517, v1519, v1526, v1528, v1533, v1540, v1542, v1549, v155, v1551, v1558, v1560, v1567, v1569, v1576, v1578, v158, v1583, v1590, v1592, v1599, v1601, v1608, v161, v1610, v1617, v1619, v1626, v1628, v1633, v164, v1640, v1642, v1649, v1651, v1658, v1660, v1667, v1669, v167, v1676, v1678, v1683, v169, v1690, v1692, v1699, v1701, v1708, v1710, v1717, v1719, v1726, v1728, v1733, v174, v1740, v1742, v1749, v1751, v1758, v1760, v1767, v1769, v177, v1776, v1778, v1783, v179, v1790, v1792, v1799, v1801, v1808, v1810, v1817, v1819, v1826, v1828, v1833, v1840, v1842, v1849, v1851, v1858, v186, v1860, v1867, v1869, v1876, v1878, v188, v1883, v1890, v1892, v1899, v1901, v1908, v1910, v1917, v1919, v1926, v1928, v193, v1933, v1940, v1942, v1949, v1951, v1958, v196, v1960, v1967, v1969, v1976, v1978, v198, v1983, v1990, v1992, v1999, v2, v20, v2001, v2008, v2010, v2017, v2019, v2026, v2028, v203, v2033, v2040, v2042, v2049, v2051, v2058, v206, v2060, v2067, v2069, v2076, v2078, v208, v2083, v2090, v2092, v2099, v2101, v2108, v2110, v2117, v2119, v2126, v2128, v2133, v2140, v2142, v2149, v215, v2151, v2158, v2160, v2167, v2169, v217, v2176, v2178, v2183, v2190, v2192, v2199, v2201, v2208, v2210, v2217, v2219, v222, v2226, v2228, v2233, v2240, v2242, v2249, v225, v2251, v2258, v2260, v2267, v2269, v227, v2276, v2278, v2283, v2290, v2292, v2299, v23, v2301, v2308, v2310, v2317, v2319, v232, v2326, v2328, v2333, v2340, v2342, v2349, v235, v2351, v2358, v2360, v2367, v2369, v237, v2376, v2378, v2383, v2390, v2392, v2399, v2401, v2408, v2410, v2417, v2419, v2426, v2428, v2433, v244, v2440, v2442, v2449, v2451, v2458, v246, v2460, v2467, v2469, v2476, v2478, v2483, v2490, v2492, v2499, v25, v2501, v2508, v251, v2510, v2517, v2519, v2526, v2528, v2533, v254, v2540, v2542, v2549, v2551, v2558, v256, v2560, v2567, v2569, v2576, v2578, v2583, v2590, v2592, v2599, v2601, v2608, v261, v2610, v2617, v2619, v2626, v2628, v2633, v264, v2640, v2642, v2649, v2651, v2658, v266, v2660, v2667, v2669, v2676, v2678, v2683, v2690, v2692, v2699, v2701, v2708, v2710, v2717, v2719, v2726, v2728, v273, v275, v280, v283, v285, v290, v293, v295, v302, v304, v309, v310, v313, v316, v319, v32, v322, v325, v327, v332, v335, v337, v34, v344, v346, v351, v354, v356, v361, v364, v366, v373, v375, v380, v383, v385, v39, v390, v393, v395, v402, v404, v409, v412, v414, v419, v42, v422, v424, v431, v433, v438, v44, v441, v443, v448, v451, v453, v460, v462, v467, v468, v471, v474, v477, v480, v483, v485, v49, v490, v493, v495, v5, v502, v504, v509, v512, v514, v519, v52, v522, v524, v531, v533, v538, v54, v541, v543, v548, v551, v553, v560, v562, v567, v570, v572, v577, v580, v582, v589, v591, v596, v599, v601, v606, v609, v61, v611, v618, v620, v625, v626, v629, v63, v632, v635, v638, v641, v643, v648, v651, v653, v660, v662, v667, v670, v672, v677, v68, v680, v682, v689, v691, v696, v699, v701, v706, v709, v71, v711, v718, v720, v725, v728, v73, v730, v735, v738, v740, v747, v749, v754, v757, v759, v764, v767, v769, v776, v778, v78, v783, v790, v792, v799, v801, v808, v81, v810, v817, v819, v826, v828, v83, v833, v840, v842, v849, v851, v858, v860, v867, v869, v876, v878, v883, v890, v892, v899, v90, v901, v908, v910, v917, v919, v92, v926, v928, v933, v940, v942, v949, v951, v958, v960, v967, v969, v97, v976, v978, v983, v990, v992, v999 int32
	var eps, epsl, huge, hugel, tiny, tinyl, v117, v118, v146, v147, v184, v185, v213, v214, v242, v243, v271, v272, v30, v300, v301, v31, v342, v343, v371, v372, v400, v401, v429, v430, v458, v459, v500, v501, v529, v530, v558, v559, v587, v588, v59, v60, v616, v617, v658, v659, v687, v688, v716, v717, v745, v746, v774, v775, v88, v89 float64
	var epsf, hugef, tinyf float32
	var v1002, v1004, v1011, v1013, v1020, v1022, v1029, v103, v1031, v1034, v1036, v1043, v1045, v105, v1052, v1054, v1061, v1063, v1070, v1072, v1079, v1081, v1084, v1086, v1093, v1095, v1102, v1104, v1111, v1113, v1120, v1122, v1129, v113, v1131, v1134, v1136, v1143, v1145, v115, v1152, v1154, v1161, v1163, v1170, v1172, v1179, v1181, v1184, v1186, v1193, v1195, v1202, v1204, v1211, v1213, v122, v1220, v1222, v1229, v1231, v124, v1243, v1245, v1252, v1254, v1261, v1263, v1270, v1272, v1279, v1281, v1293, v1295, v1302, v1304, v1311, v1313, v132, v1320, v1322, v1329, v1331, v134, v1343, v1345, v1352, v1354, v1361, v1363, v1370, v1372, v1379, v1381, v1393, v1395, v1402, v1404, v1411, v1413, v142, v1420, v1422, v1429, v1431, v144, v1443, v1445, v1452, v1454, v1461, v1463, v1470, v1472, v1479, v1481, v1493, v1495, v1502, v1504, v151, v1511, v1513, v1520, v1522, v1529, v153, v1531, v1543, v1545, v1552, v1554, v156, v1561, v1563, v1570, v1572, v1579, v1581, v1593, v1595, v16, v1602, v1604, v1611, v1613, v162, v1620, v1622, v1629, v1631, v1643, v1645, v1652, v1654, v1661, v1663, v1670, v1672, v1679, v1681, v1684, v1686, v1693, v1695, v170, v1702, v1704, v1711, v1713, v172, v1720, v1722, v1729, v1731, v1734, v1736, v1743, v1745, v1752, v1754, v1761, v1763, v1770, v1772, v1779, v1781, v1784, v1786, v1793, v1795, v18, v180, v1802, v1804, v1811, v1813, v182, v1820, v1822, v1829, v1831, v1834, v1836, v1843, v1845, v1852, v1854, v1861, v1863, v1870, v1872, v1879, v1881, v1884, v1886, v189, v1893, v1895, v1902, v1904, v191, v1911, v1913, v1920, v1922, v1929, v1931, v1934, v1936, v1943, v1945, v1952, v1954, v1961, v1963, v1970, v1972, v1979, v1981, v1984, v1986, v199, v1993, v1995, v2002, v2004, v201, v2011, v2013, v2020, v2022, v2029, v2031, v2034, v2036, v2043, v2045, v2052, v2054, v2061, v2063, v2070, v2072, v2079, v2081, v2084, v2086, v209, v2093, v2095, v2102, v2104, v211, v2111, v2113, v2120, v2122, v2129, v2131, v2134, v2136, v2143, v2145, v2152, v2154, v2161, v2163, v2170, v2172, v2179, v218, v2181, v2184, v2186, v2193, v2195, v220, v2202, v2204, v2211, v2213, v2220, v2222, v2229, v2231, v2234, v2236, v2243, v2245, v2252, v2254, v2261, v2263, v2270, v2272, v2279, v228, v2281, v2284, v2286, v2293, v2295, v230, v2302, v2304, v2311, v2313, v2320, v2322, v2329, v2331, v2343, v2345, v2352, v2354, v2361, v2363, v2370, v2372, v2379, v238, v2381, v2393, v2395, v240, v2402, v2404, v2411, v2413, v2420, v2422, v2429, v2431, v2443, v2445, v2452, v2454, v2461, v2463, v247, v2470, v2472, v2479, v2481, v249, v2493, v2495, v2502, v2504, v2511, v2513, v2520, v2522, v2529, v2531, v2534, v2536, v2543, v2545, v2552, v2554, v2561, v2563, v257, v2570, v2572, v2579, v2581, v2584, v2586, v259, v2593, v2595, v26, v2602, v2604, v2611, v2613, v2620, v2622, v2629, v2631, v2634, v2636, v2643, v2645, v2652, v2654, v2661, v2663, v267, v2670, v2672, v2679, v2681, v2684, v2686, v269, v2693, v2695, v2702, v2704, v2711, v2713, v2720, v2722, v2729, v2731, v276, v278, v28, v286, v288, v296, v298, v305, v307, v314, v320, v328, v330, v338, v340, v347, v349, v35, v357, v359, v367, v369, v37, v376, v378, v386, v388, v396, v398, v405, v407, v415, v417, v425, v427, v434, v436, v444, v446, v45, v454, v456, v463, v465, v47, v472, v478, v486, v488, v496, v498, v505, v507, v515, v517, v525, v527, v534, v536, v544, v546, v55, v554, v556, v563, v565, v57, v573, v575, v583, v585, v592, v594, v6, v602, v604, v612, v614, v621, v623, v630, v636, v64, v644, v646, v654, v656, v66, v663, v665, v673, v675, v683, v685, v692, v694, v702, v704, v712, v714, v721, v723, v731, v733, v74, v741, v743, v750, v752, v76, v760, v762, v770, v772, v779, v781, v784, v786, v793, v795, v8, v802, v804, v811, v813, v820, v822, v829, v831, v834, v836, v84, v843, v845, v852, v854, v86, v861, v863, v870, v872, v879, v881, v884, v886, v893, v895, v902, v904, v911, v913, v920, v922, v929, v93, v931, v934, v936, v943, v945, v95, v952, v954, v961, v963, v970, v972, v979, v981, v984, v986, v993, v995 uint64
	var v1006, v1007, v1015, v1016, v1024, v1025, v1038, v1039, v1047, v1048, v1056, v1057, v1065, v1066, v1074, v1075, v108, v1088, v1089, v109, v1097, v1098, v1106, v1107, v1115, v1116, v1124, v1125, v1138, v1139, v1147, v1148, v1156, v1157, v1165, v1166, v1174, v1175, v1188, v1189, v1197, v1198, v1206, v1207, v1215, v1216, v1224, v1225, v137, v138, v1688, v1689, v1697, v1698, v1706, v1707, v1715, v1716, v1724, v1725, v1738, v1739, v1747, v1748, v175, v1756, v1757, v176, v1765, v1766, v1774, v1775, v1788, v1789, v1797, v1798, v1806, v1807, v1815, v1816, v1824, v1825, v1838, v1839, v1847, v1848, v1856, v1857, v1865, v1866, v1874, v1875, v1888, v1889, v1897, v1898, v1906, v1907, v1915, v1916, v1924, v1925, v1938, v1939, v1947, v1948, v1956, v1957, v1965, v1966, v1974, v1975, v1988, v1989, v1997, v1998, v2006, v2007, v2015, v2016, v2024, v2025, v2038, v2039, v204, v2047, v2048, v205, v2056, v2057, v2065, v2066, v2074, v2075, v2088, v2089, v2097, v2098, v21, v2106, v2107, v2115, v2116, v2124, v2125, v2138, v2139, v2147, v2148, v2156, v2157, v2165, v2166, v2174, v2175, v2188, v2189, v2197, v2198, v22, v2206, v2207, v2215, v2216, v2224, v2225, v2238, v2239, v2247, v2248, v2256, v2257, v2265, v2266, v2274, v2275, v2288, v2289, v2297, v2298, v2306, v2307, v2315, v2316, v2324, v2325, v233, v234, v2538, v2539, v2547, v2548, v2556, v2557, v2565, v2566, v2574, v2575, v2588, v2589, v2597, v2598, v2606, v2607, v2615, v2616, v262, v2624, v2625, v263, v2638, v2639, v2647, v2648, v2656, v2657, v2665, v2666, v2674, v2675, v2688, v2689, v2697, v2698, v2706, v2707, v2715, v2716, v2724, v2725, v291, v292, v333, v334, v362, v363, v391, v392, v420, v421, v449, v450, v491, v492, v50, v51, v520, v521, v549, v550, v578, v579, v607, v608, v649, v650, v678, v679, v707, v708, v736, v737, v765, v766, v788, v789, v79, v797, v798, v80, v806, v807, v815, v816, v824, v825, v838, v839, v847, v848, v856, v857, v865, v866, v874, v875, v888, v889, v897, v898, v906, v907, v915, v916, v924, v925, v938, v939, v947, v948, v956, v957, v965, v966, v974, v975, v988, v989, v997, v998 double_t
	var v11, v12, v1238, v1239, v1247, v1248, v1256, v1257, v1265, v1266, v127, v1274, v1275, v128, v1288, v1289, v1297, v1298, v1306, v1307, v1315, v1316, v1324, v1325, v1338, v1339, v1347, v1348, v1356, v1357, v1365, v1366, v1374, v1375, v1388, v1389, v1397, v1398, v1406, v1407, v1415, v1416, v1424, v1425, v1438, v1439, v1447, v1448, v1456, v1457, v1465, v1466, v1474, v1475, v1488, v1489, v1497, v1498, v1506, v1507, v1515, v1516, v1524, v1525, v1538, v1539, v1547, v1548, v1556, v1557, v1565, v1566, v1574, v1575, v1588, v1589, v1597, v1598, v1606, v1607, v1615, v1616, v1624, v1625, v1638, v1639, v1647, v1648, v165, v1656, v1657, v166, v1665, v1666, v1674, v1675, v194, v195, v223, v224, v2338, v2339, v2347, v2348, v2356, v2357, v2365, v2366, v2374, v2375, v2388, v2389, v2397, v2398, v2406, v2407, v2415, v2416, v2424, v2425, v2438, v2439, v2447, v2448, v2456, v2457, v2465, v2466, v2474, v2475, v2488, v2489, v2497, v2498, v2506, v2507, v2515, v2516, v252, v2524, v2525, v253, v281, v282, v323, v324, v352, v353, v381, v382, v40, v41, v410, v411, v439, v440, v481, v482, v510, v511, v539, v540, v568, v569, v597, v598, v639, v640, v668, v669, v69, v697, v698, v70, v726, v727, v755, v756, v98, v99 float_t
	var v1234, v1236, v1284, v1286, v1334, v1336, v1384, v1386, v1434, v1436, v1484, v1486, v1534, v1536, v1584, v1586, v159, v1634, v1636, v2334, v2336, v2384, v2386, v2434, v2436, v2484, v2486, v3, v311, v317, v469, v475, v627, v633 uint32
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint32
		F__f float32
	}
	var _ /* __u at bp+8 */ struct {
		F__i [0]uint64
		F__f float64
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = e, e1, e10, e100, e101, e102, e103, e104, e105, e106, e107, e108, e109, e11, e110, e111, e112, e113, e114, e115, e116, e117, e118, e119, e12, e120, e121, e122, e123, e124, e125, e126, e127, e128, e129, e13, e130, e131, e132, e133, e134, e135, e136, e137, e138, e139, e14, e140, e141, e142, e143, e144, e145, e146, e147, e148, e149, e15, e150, e151, e152, e153, e154, e155, e156, e157, e158, e159, e16, e160, e161, e162, e163, e164, e165, e166, e167, e168, e169, e17, e170, e171, e172, e173, e174, e175, e176, e177, e178, e179, e18, e180, e181, e182, e183, e184, e185, e186, e187, e188, e189, e19, e190, e191, e192, e193, e194, e195, e196, e197, e198, e199, e2, e20, e200, e201, e202, e203, e204, e205, e206, e207, e208, e209, e21, e210, e211, e212, e213, e214, e215, e216, e217, e218, e219, e22, e220, e221, e222, e223, e224, e225, e226, e227, e228, e229, e23, e230, e231, e232, e233, e234, e235, e236, e237, e238, e239, e24, e240, e241, e242, e243, e244, e245, e246, e247, e248, e249, e25, e250, e251, e252, e253, e254, e255, e256, e257, e258, e259, e26, e260, e261, e262, e263, e27, e28, e29, e3, e30, e31, e32, e33, e34, e35, e36, e37, e38, e39, e4, e40, e41, e42, e43, e44, e45, e46, e47, e48, e49, e5, e50, e51, e52, e53, e54, e55, e56, e57, e58, e59, e6, e60, e61, e62, e63, e64, e65, e66, e67, e68, e69, e7, e70, e71, e72, e73, e74, e75, e76, e77, e78, e79, e8, e80, e81, e82, e83, e84, e85, e86, e87, e88, e89, e9, e90, e91, e92, e93, e94, e95, e96, e97, e98, e99, eps, epsf, epsl, huge, hugef, hugel, r, r1, r10, r100, r101, r102, r103, r104, r105, r106, r107, r108, r109, r11, r110, r111, r112, r113, r114, r115, r116, r117, r118, r119, r12, r120, r121, r122, r123, r124, r125, r126, r127, r128, r129, r13, r130, r131, r132, r133, r134, r135, r136, r137, r138, r139, r14, r140, r141, r142, r143, r144, r145, r146, r147, r148, r149, r15, r150, r151, r152, r153, r154, r155, r156, r157, r158, r159, r16, r160, r161, r162, r163, r164, r165, r166, r167, r168, r169, r17, r170, r171, r172, r173, r174, r175, r176, r177, r178, r179, r18, r180, r181, r182, r183, r184, r185, r186, r187, r188, r189, r19, r190, r191, r192, r193, r194, r195, r196, r197, r198, r199, r2, r20, r200, r201, r202, r203, r204, r205, r206, r207, r208, r209, r21, r210, r211, r212, r213, r214, r215, r216, r217, r218, r219, r22, r220, r221, r222, r223, r224, r225, r226, r227, r228, r229, r23, r230, r231, r232, r233, r234, r235, r236, r237, r238, r239, r24, r240, r241, r242, r243, r244, r245, r246, r247, r248, r249, r25, r250, r251, r252, r253, r254, r255, r256, r257, r258, r259, r26, r260, r261, r262, r263, r27, r28, r29, r3, r30, r31, r32, r33, r34, r35, r36, r37, r38, r39, r4, r40, r41, r42, r43, r44, r45, r46, r47, r48, r49, r5, r50, r51, r52, r53, r54, r55, r56, r57, r58, r59, r6, r60, r61, r62, r63, r64, r65, r66, r67, r68, r69, r7, r70, r71, r72, r73, r74, r75, r76, r77, r78, r79, r8, r80, r81, r82, r83, r84, r85, r86, r87, r88, r89, r9, r90, r91, r92, r93, r94, r95, r96, r97, r98, r99, tiny, tinyf, tinyl, v1, v10, v100, v1001, v1002, v1004, v1006, v1007, v1008, v1010, v1011, v1013, v1015, v1016, v1017, v1019, v102, v1020, v1022, v1024, v1025, v1026, v1028, v1029, v103, v1031, v1033, v1034, v1036, v1038, v1039, v1040, v1042, v1043, v1045, v1047, v1048, v1049, v105, v1051, v1052, v1054, v1056, v1057, v1058, v1060, v1061, v1063, v1065, v1066, v1067, v1069, v107, v1070, v1072, v1074, v1075, v1076, v1078, v1079, v108, v1081, v1083, v1084, v1086, v1088, v1089, v109, v1090, v1092, v1093, v1095, v1097, v1098, v1099, v11, v110, v1101, v1102, v1104, v1106, v1107, v1108, v1110, v1111, v1113, v1115, v1116, v1117, v1119, v112, v1120, v1122, v1124, v1125, v1126, v1128, v1129, v113, v1131, v1133, v1134, v1136, v1138, v1139, v1140, v1142, v1143, v1145, v1147, v1148, v1149, v115, v1151, v1152, v1154, v1156, v1157, v1158, v1160, v1161, v1163, v1165, v1166, v1167, v1169, v117, v1170, v1172, v1174, v1175, v1176, v1178, v1179, v118, v1181, v1183, v1184, v1186, v1188, v1189, v119, v1190, v1192, v1193, v1195, v1197, v1198, v1199, v12, v1201, v1202, v1204, v1206, v1207, v1208, v121, v1210, v1211, v1213, v1215, v1216, v1217, v1219, v122, v1220, v1222, v1224, v1225, v1226, v1228, v1229, v1231, v1233, v1234, v1236, v1238, v1239, v124, v1240, v1242, v1243, v1245, v1247, v1248, v1249, v1251, v1252, v1254, v1256, v1257, v1258, v126, v1260, v1261, v1263, v1265, v1266, v1267, v1269, v127, v1270, v1272, v1274, v1275, v1276, v1278, v1279, v128, v1281, v1283, v1284, v1286, v1288, v1289, v129, v1290, v1292, v1293, v1295, v1297, v1298, v1299, v13, v1301, v1302, v1304, v1306, v1307, v1308, v131, v1310, v1311, v1313, v1315, v1316, v1317, v1319, v132, v1320, v1322, v1324, v1325, v1326, v1328, v1329, v1331, v1333, v1334, v1336, v1338, v1339, v134, v1340, v1342, v1343, v1345, v1347, v1348, v1349, v1351, v1352, v1354, v1356, v1357, v1358, v136, v1360, v1361, v1363, v1365, v1366, v1367, v1369, v137, v1370, v1372, v1374, v1375, v1376, v1378, v1379, v138, v1381, v1383, v1384, v1386, v1388, v1389, v139, v1390, v1392, v1393, v1395, v1397, v1398, v1399, v1401, v1402, v1404, v1406, v1407, v1408, v141, v1410, v1411, v1413, v1415, v1416, v1417, v1419, v142, v1420, v1422, v1424, v1425, v1426, v1428, v1429, v1431, v1433, v1434, v1436, v1438, v1439, v144, v1440, v1442, v1443, v1445, v1447, v1448, v1449, v1451, v1452, v1454, v1456, v1457, v1458, v146, v1460, v1461, v1463, v1465, v1466, v1467, v1469, v147, v1470, v1472, v1474, v1475, v1476, v1478, v1479, v148, v1481, v1483, v1484, v1486, v1488, v1489, v1490, v1492, v1493, v1495, v1497, v1498, v1499, v15, v150, v1501, v1502, v1504, v1506, v1507, v1508, v151, v1510, v1511, v1513, v1515, v1516, v1517, v1519, v1520, v1522, v1524, v1525, v1526, v1528, v1529, v153, v1531, v1533, v1534, v1536, v1538, v1539, v1540, v1542, v1543, v1545, v1547, v1548, v1549, v155, v1551, v1552, v1554, v1556, v1557, v1558, v156, v1560, v1561, v1563, v1565, v1566, v1567, v1569, v1570, v1572, v1574, v1575, v1576, v1578, v1579, v158, v1581, v1583, v1584, v1586, v1588, v1589, v159, v1590, v1592, v1593, v1595, v1597, v1598, v1599, v16, v1601, v1602, v1604, v1606, v1607, v1608, v161, v1610, v1611, v1613, v1615, v1616, v1617, v1619, v162, v1620, v1622, v1624, v1625, v1626, v1628, v1629, v1631, v1633, v1634, v1636, v1638, v1639, v164, v1640, v1642, v1643, v1645, v1647, v1648, v1649, v165, v1651, v1652, v1654, v1656, v1657, v1658, v166, v1660, v1661, v1663, v1665, v1666, v1667, v1669, v167, v1670, v1672, v1674, v1675, v1676, v1678, v1679, v1681, v1683, v1684, v1686, v1688, v1689, v169, v1690, v1692, v1693, v1695, v1697, v1698, v1699, v170, v1701, v1702, v1704, v1706, v1707, v1708, v1710, v1711, v1713, v1715, v1716, v1717, v1719, v172, v1720, v1722, v1724, v1725, v1726, v1728, v1729, v1731, v1733, v1734, v1736, v1738, v1739, v174, v1740, v1742, v1743, v1745, v1747, v1748, v1749, v175, v1751, v1752, v1754, v1756, v1757, v1758, v176, v1760, v1761, v1763, v1765, v1766, v1767, v1769, v177, v1770, v1772, v1774, v1775, v1776, v1778, v1779, v1781, v1783, v1784, v1786, v1788, v1789, v179, v1790, v1792, v1793, v1795, v1797, v1798, v1799, v18, v180, v1801, v1802, v1804, v1806, v1807, v1808, v1810, v1811, v1813, v1815, v1816, v1817, v1819, v182, v1820, v1822, v1824, v1825, v1826, v1828, v1829, v1831, v1833, v1834, v1836, v1838, v1839, v184, v1840, v1842, v1843, v1845, v1847, v1848, v1849, v185, v1851, v1852, v1854, v1856, v1857, v1858, v186, v1860, v1861, v1863, v1865, v1866, v1867, v1869, v1870, v1872, v1874, v1875, v1876, v1878, v1879, v188, v1881, v1883, v1884, v1886, v1888, v1889, v189, v1890, v1892, v1893, v1895, v1897, v1898, v1899, v1901, v1902, v1904, v1906, v1907, v1908, v191, v1910, v1911, v1913, v1915, v1916, v1917, v1919, v1920, v1922, v1924, v1925, v1926, v1928, v1929, v193, v1931, v1933, v1934, v1936, v1938, v1939, v194, v1940, v1942, v1943, v1945, v1947, v1948, v1949, v195, v1951, v1952, v1954, v1956, v1957, v1958, v196, v1960, v1961, v1963, v1965, v1966, v1967, v1969, v1970, v1972, v1974, v1975, v1976, v1978, v1979, v198, v1981, v1983, v1984, v1986, v1988, v1989, v199, v1990, v1992, v1993, v1995, v1997, v1998, v1999, v2, v20, v2001, v2002, v2004, v2006, v2007, v2008, v201, v2010, v2011, v2013, v2015, v2016, v2017, v2019, v2020, v2022, v2024, v2025, v2026, v2028, v2029, v203, v2031, v2033, v2034, v2036, v2038, v2039, v204, v2040, v2042, v2043, v2045, v2047, v2048, v2049, v205, v2051, v2052, v2054, v2056, v2057, v2058, v206, v2060, v2061, v2063, v2065, v2066, v2067, v2069, v2070, v2072, v2074, v2075, v2076, v2078, v2079, v208, v2081, v2083, v2084, v2086, v2088, v2089, v209, v2090, v2092, v2093, v2095, v2097, v2098, v2099, v21, v2101, v2102, v2104, v2106, v2107, v2108, v211, v2110, v2111, v2113, v2115, v2116, v2117, v2119, v2120, v2122, v2124, v2125, v2126, v2128, v2129, v213, v2131, v2133, v2134, v2136, v2138, v2139, v214, v2140, v2142, v2143, v2145, v2147, v2148, v2149, v215, v2151, v2152, v2154, v2156, v2157, v2158, v2160, v2161, v2163, v2165, v2166, v2167, v2169, v217, v2170, v2172, v2174, v2175, v2176, v2178, v2179, v218, v2181, v2183, v2184, v2186, v2188, v2189, v2190, v2192, v2193, v2195, v2197, v2198, v2199, v22, v220, v2201, v2202, v2204, v2206, v2207, v2208, v2210, v2211, v2213, v2215, v2216, v2217, v2219, v222, v2220, v2222, v2224, v2225, v2226, v2228, v2229, v223, v2231, v2233, v2234, v2236, v2238, v2239, v224, v2240, v2242, v2243, v2245, v2247, v2248, v2249, v225, v2251, v2252, v2254, v2256, v2257, v2258, v2260, v2261, v2263, v2265, v2266, v2267, v2269, v227, v2270, v2272, v2274, v2275, v2276, v2278, v2279, v228, v2281, v2283, v2284, v2286, v2288, v2289, v2290, v2292, v2293, v2295, v2297, v2298, v2299, v23, v230, v2301, v2302, v2304, v2306, v2307, v2308, v2310, v2311, v2313, v2315, v2316, v2317, v2319, v232, v2320, v2322, v2324, v2325, v2326, v2328, v2329, v233, v2331, v2333, v2334, v2336, v2338, v2339, v234, v2340, v2342, v2343, v2345, v2347, v2348, v2349, v235, v2351, v2352, v2354, v2356, v2357, v2358, v2360, v2361, v2363, v2365, v2366, v2367, v2369, v237, v2370, v2372, v2374, v2375, v2376, v2378, v2379, v238, v2381, v2383, v2384, v2386, v2388, v2389, v2390, v2392, v2393, v2395, v2397, v2398, v2399, v240, v2401, v2402, v2404, v2406, v2407, v2408, v2410, v2411, v2413, v2415, v2416, v2417, v2419, v242, v2420, v2422, v2424, v2425, v2426, v2428, v2429, v243, v2431, v2433, v2434, v2436, v2438, v2439, v244, v2440, v2442, v2443, v2445, v2447, v2448, v2449, v2451, v2452, v2454, v2456, v2457, v2458, v246, v2460, v2461, v2463, v2465, v2466, v2467, v2469, v247, v2470, v2472, v2474, v2475, v2476, v2478, v2479, v2481, v2483, v2484, v2486, v2488, v2489, v249, v2490, v2492, v2493, v2495, v2497, v2498, v2499, v25, v2501, v2502, v2504, v2506, v2507, v2508, v251, v2510, v2511, v2513, v2515, v2516, v2517, v2519, v252, v2520, v2522, v2524, v2525, v2526, v2528, v2529, v253, v2531, v2533, v2534, v2536, v2538, v2539, v254, v2540, v2542, v2543, v2545, v2547, v2548, v2549, v2551, v2552, v2554, v2556, v2557, v2558, v256, v2560, v2561, v2563, v2565, v2566, v2567, v2569, v257, v2570, v2572, v2574, v2575, v2576, v2578, v2579, v2581, v2583, v2584, v2586, v2588, v2589, v259, v2590, v2592, v2593, v2595, v2597, v2598, v2599, v26, v2601, v2602, v2604, v2606, v2607, v2608, v261, v2610, v2611, v2613, v2615, v2616, v2617, v2619, v262, v2620, v2622, v2624, v2625, v2626, v2628, v2629, v263, v2631, v2633, v2634, v2636, v2638, v2639, v264, v2640, v2642, v2643, v2645, v2647, v2648, v2649, v2651, v2652, v2654, v2656, v2657, v2658, v266, v2660, v2661, v2663, v2665, v2666, v2667, v2669, v267, v2670, v2672, v2674, v2675, v2676, v2678, v2679, v2681, v2683, v2684, v2686, v2688, v2689, v269, v2690, v2692, v2693, v2695, v2697, v2698, v2699, v2701, v2702, v2704, v2706, v2707, v2708, v271, v2710, v2711, v2713, v2715, v2716, v2717, v2719, v272, v2720, v2722, v2724, v2725, v2726, v2728, v2729, v273, v2731, v275, v276, v278, v28, v280, v281, v282, v283, v285, v286, v288, v290, v291, v292, v293, v295, v296, v298, v3, v30, v300, v301, v302, v304, v305, v307, v309, v31, v310, v311, v313, v314, v316, v317, v319, v32, v320, v322, v323, v324, v325, v327, v328, v330, v332, v333, v334, v335, v337, v338, v34, v340, v342, v343, v344, v346, v347, v349, v35, v351, v352, v353, v354, v356, v357, v359, v361, v362, v363, v364, v366, v367, v369, v37, v371, v372, v373, v375, v376, v378, v380, v381, v382, v383, v385, v386, v388, v39, v390, v391, v392, v393, v395, v396, v398, v40, v400, v401, v402, v404, v405, v407, v409, v41, v410, v411, v412, v414, v415, v417, v419, v42, v420, v421, v422, v424, v425, v427, v429, v430, v431, v433, v434, v436, v438, v439, v44, v440, v441, v443, v444, v446, v448, v449, v45, v450, v451, v453, v454, v456, v458, v459, v460, v462, v463, v465, v467, v468, v469, v47, v471, v472, v474, v475, v477, v478, v480, v481, v482, v483, v485, v486, v488, v49, v490, v491, v492, v493, v495, v496, v498, v5, v50, v500, v501, v502, v504, v505, v507, v509, v51, v510, v511, v512, v514, v515, v517, v519, v52, v520, v521, v522, v524, v525, v527, v529, v530, v531, v533, v534, v536, v538, v539, v54, v540, v541, v543, v544, v546, v548, v549, v55, v550, v551, v553, v554, v556, v558, v559, v560, v562, v563, v565, v567, v568, v569, v57, v570, v572, v573, v575, v577, v578, v579, v580, v582, v583, v585, v587, v588, v589, v59, v591, v592, v594, v596, v597, v598, v599, v6, v60, v601, v602, v604, v606, v607, v608, v609, v61, v611, v612, v614, v616, v617, v618, v620, v621, v623, v625, v626, v627, v629, v63, v630, v632, v633, v635, v636, v638, v639, v64, v640, v641, v643, v644, v646, v648, v649, v650, v651, v653, v654, v656, v658, v659, v66, v660, v662, v663, v665, v667, v668, v669, v670, v672, v673, v675, v677, v678, v679, v68, v680, v682, v683, v685, v687, v688, v689, v69, v691, v692, v694, v696, v697, v698, v699, v70, v701, v702, v704, v706, v707, v708, v709, v71, v711, v712, v714, v716, v717, v718, v720, v721, v723, v725, v726, v727, v728, v73, v730, v731, v733, v735, v736, v737, v738, v74, v740, v741, v743, v745, v746, v747, v749, v750, v752, v754, v755, v756, v757, v759, v76, v760, v762, v764, v765, v766, v767, v769, v770, v772, v774, v775, v776, v778, v779, v78, v781, v783, v784, v786, v788, v789, v79, v790, v792, v793, v795, v797, v798, v799, v8, v80, v801, v802, v804, v806, v807, v808, v81, v810, v811, v813, v815, v816, v817, v819, v820, v822, v824, v825, v826, v828, v829, v83, v831, v833, v834, v836, v838, v839, v84, v840, v842, v843, v845, v847, v848, v849, v851, v852, v854, v856, v857, v858, v86, v860, v861, v863, v865, v866, v867, v869, v870, v872, v874, v875, v876, v878, v879, v88, v881, v883, v884, v886, v888, v889, v89, v890, v892, v893, v895, v897, v898, v899, v90, v901, v902, v904, v906, v907, v908, v910, v911, v913, v915, v916, v917, v919, v92, v920, v922, v924, v925, v926, v928, v929, v93, v931, v933, v934, v936, v938, v939, v940, v942, v943, v945, v947, v948, v949, v95, v951, v952, v954, v956, v957, v958, v960, v961, v963, v965, v966, v967, v969, v97, v970, v972, v974, v975, v976, v978, v979, v98, v981, v983, v984, v986, v988, v989, v99, v990, v992, v993, v995, v997, v998, v999
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:31:18:
	huge = float64(1.7976931348623157e+308)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:32:18:
	tiny = float64(2.2250738585072014e-308)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:33:18:
	eps = float64(2.220446049250313e-16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:34:17:
	hugef = libc.Float32FromFloat32(3.4028234663852886e+38)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:35:17:
	tinyf = libc.Float32FromFloat32(1.1754943508222875e-38)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:36:17:
	epsf = libc.Float32FromFloat32(1.1920928955078125e-07)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:37:23:
	hugel = libc.Float64FromFloat64(1.79769313486231570815e+308)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:38:23:
	tinyl = libc.Float64FromFloat64(2.22507385850720138309e-308)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:39:23:
	epsl = libc.Float64FromFloat64(2.22044604925031308085e-16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(4) == uint32(4) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _4
	_4:
		v2 = libc.BoolInt32(v3&uint32(0x7fffffff) > uint32(0x7f800000))
	} else {
		if uint32(4) == uint32(8) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v6 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _7
		_7:
			v5 = libc.BoolInt32(v6&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		} else {
			v5 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
		}
		v2 = v5
	}
	if v2 != 0 {
		_ = libc.Float64FromFloat64(1)
		v1 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v8 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _9
	_9:
		v1 = libc.BoolInt32(v8&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r = v1
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+1, libc.VaList(bp+24, r, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+84, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(8) == uint32(4) {
		v11 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v12 = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v11
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v16 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _17
	_17:
		if libc.BoolInt32(v16&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v12
			v15 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v12
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v18 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _19
		_19:
			v15 = libc.BoolInt32(v18&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v13 = libc.BoolInt32(!(v15 != 0) && v11 < v12)
		goto _14
	_14:
		v10 = v13
	} else {
		if uint32(8) == uint32(8) {
			v21 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v22 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v21
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v26 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _27
		_27:
			if libc.BoolInt32(v26&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v22
				v25 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v22
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v28 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _29
			_29:
				v25 = libc.BoolInt32(v28&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v23 = libc.BoolInt32(!(v25 != 0) && v21 < v22)
			goto _24
		_24:
			v20 = v23
		} else {
			v30 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v31 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:110:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v30
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v35 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _36
		_36:
			if libc.BoolInt32(v35&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v31
				v34 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v31
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v37 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _38
			_38:
				v34 = libc.BoolInt32(v37&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v32 = libc.BoolInt32(!(v34 != 0) && v30 < v31)
			goto _33
		_33:
			v20 = v32
		}
		v10 = v20
	}
	r1 = v10
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e1 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r1 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+173, libc.VaList(bp+24, r1, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e1&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+251, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(8) == uint32(4) {
		v40 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v41 = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v40
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v45 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _46
	_46:
		if libc.BoolInt32(v45&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v41
			v44 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v41
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v47 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _48
		_48:
			v44 = libc.BoolInt32(v47&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v42 = libc.BoolInt32(!(v44 != 0) && v40 <= v41)
		goto _43
	_43:
		v39 = v42
	} else {
		if uint32(8) == uint32(8) {
			v50 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v51 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v50
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v55 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _56
		_56:
			if libc.BoolInt32(v55&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v51
				v54 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v51
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v57 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _58
			_58:
				v54 = libc.BoolInt32(v57&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v52 = libc.BoolInt32(!(v54 != 0) && v50 <= v51)
			goto _53
		_53:
			v49 = v52
		} else {
			v59 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v60 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:113:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v59
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v64 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _65
		_65:
			if libc.BoolInt32(v64&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v60
				v63 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v60
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v66 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _67
			_67:
				v63 = libc.BoolInt32(v66&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v61 = libc.BoolInt32(!(v63 != 0) && v59 <= v60)
			goto _62
		_62:
			v49 = v61
		}
		v39 = v49
	}
	r2 = v39
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e2 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r2 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+335, libc.VaList(bp+24, r2, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e2&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+418, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(8) == uint32(4) {
		v69 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v70 = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v69
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v74 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _75
	_75:
		if libc.BoolInt32(v74&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v70
			v73 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v70
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v76 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _77
		_77:
			v73 = libc.BoolInt32(v76&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v71 = libc.BoolInt32(!(v73 != 0) && v69 != v70)
		goto _72
	_72:
		v68 = v71
	} else {
		if uint32(8) == uint32(8) {
			v79 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v80 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v79
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v84 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _85
		_85:
			if libc.BoolInt32(v84&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v80
				v83 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v80
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v86 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _87
			_87:
				v83 = libc.BoolInt32(v86&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v81 = libc.BoolInt32(!(v83 != 0) && v79 != v80)
			goto _82
		_82:
			v78 = v81
		} else {
			v88 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v89 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:116:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v88
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v93 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _94
		_94:
			if libc.BoolInt32(v93&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v89
				v92 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v89
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v95 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _96
			_96:
				v92 = libc.BoolInt32(v95&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v90 = libc.BoolInt32(!(v92 != 0) && v88 != v89)
			goto _91
		_91:
			v78 = v90
		}
		v68 = v78
	}
	r3 = v68
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e3 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r3 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+507, libc.VaList(bp+24, r3, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e3&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+592, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(8) == uint32(4) {
		v98 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v99 = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v98
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v103 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _104
	_104:
		if libc.BoolInt32(v103&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v99
			v102 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v99
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v105 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _106
		_106:
			v102 = libc.BoolInt32(v105&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v100 = libc.BoolInt32(!(v102 != 0) && v98 > v99)
		goto _101
	_101:
		v97 = v100
	} else {
		if uint32(8) == uint32(8) {
			v108 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v109 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v108
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v113 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _114
		_114:
			if libc.BoolInt32(v113&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v109
				v112 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v109
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v115 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _116
			_116:
				v112 = libc.BoolInt32(v115&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v110 = libc.BoolInt32(!(v112 != 0) && v108 > v109)
			goto _111
		_111:
			v107 = v110
		} else {
			v117 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v118 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:119:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v117
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v122 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _123
		_123:
			if libc.BoolInt32(v122&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v118
				v121 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v118
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v124 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _125
			_125:
				v121 = libc.BoolInt32(v124&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v119 = libc.BoolInt32(!(v121 != 0) && v117 > v118)
			goto _120
		_120:
			v107 = v119
		}
		v97 = v107
	}
	r4 = v97
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e4 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r4 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+683, libc.VaList(bp+24, r4, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e4&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+764, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if uint32(8) == uint32(4) {
		v127 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v128 = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v127
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v132 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _133
	_133:
		if libc.BoolInt32(v132&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v128
			v131 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v128
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v134 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _135
		_135:
			v131 = libc.BoolInt32(v134&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v129 = libc.BoolInt32(!(v131 != 0) && v127 >= v128)
		goto _130
	_130:
		v126 = v129
	} else {
		if uint32(8) == uint32(8) {
			v137 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v138 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v137
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v142 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _143
		_143:
			if libc.BoolInt32(v142&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v138
				v141 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v138
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v144 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _145
			_145:
				v141 = libc.BoolInt32(v144&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v139 = libc.BoolInt32(!(v141 != 0) && v137 >= v138)
			goto _140
		_140:
			v136 = v139
		} else {
			v146 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v147 = libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:122:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v146
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v151 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _152
		_152:
			if libc.BoolInt32(v151&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v147
				v150 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v147
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v153 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _154
			_154:
				v150 = libc.BoolInt32(v153&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v148 = libc.BoolInt32(!(v150 != 0) && v146 >= v147)
			goto _149
		_149:
			v136 = v148
		}
		v126 = v136
	}
	r5 = v126
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	e5 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if r5 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+851, libc.VaList(bp+24, r5, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
	if e5&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:41:2:
		t_printf(tls, __ccgo_ts+937, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v156 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _157
_157:
	if libc.BoolInt32(v156&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		libc.X__builtin_nanf(tls, __ccgo_ts)
		v155 = libc.Int32FromInt32(1)
	} else {
		if uint32(4) == uint32(4) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
			*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
			v159 = *(*uint32)(unsafe.Pointer(bp))
			goto _160
		_160:
			v158 = libc.BoolInt32(v159&uint32(0x7fffffff) > uint32(0x7f800000))
		} else {
			if uint32(4) == uint32(8) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v162 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _163
			_163:
				v161 = libc.BoolInt32(v162&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			} else {
				v161 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
			}
			v158 = v161
		}
		v155 = v158
	}
	r6 = v155
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e6 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r6 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1029, libc.VaList(bp+24, r6, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e6&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1112, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if uint32(8) == uint32(4) {
		v165 = libc.Float64FromFloat64(1)
		v166 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v165
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v170 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _171
	_171:
		if libc.BoolInt32(v170&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v166
			v169 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v166
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v172 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _173
		_173:
			v169 = libc.BoolInt32(v172&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v167 = libc.BoolInt32(!(v169 != 0) && v165 < v166)
		goto _168
	_168:
		v164 = v167
	} else {
		if uint32(8) == uint32(8) {
			v175 = libc.Float64FromFloat64(1)
			v176 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v175
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v180 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _181
		_181:
			if libc.BoolInt32(v180&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v176
				v179 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v176
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v182 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _183
			_183:
				v179 = libc.BoolInt32(v182&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v177 = libc.BoolInt32(!(v179 != 0) && v175 < v176)
			goto _178
		_178:
			v174 = v177
		} else {
			v184 = libc.Float64FromFloat64(1)
			v185 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:110:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v184
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v189 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _190
		_190:
			if libc.BoolInt32(v189&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v185
				v188 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v185
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v191 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _192
			_192:
				v188 = libc.BoolInt32(v191&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v186 = libc.BoolInt32(!(v188 != 0) && v184 < v185)
			goto _187
		_187:
			v174 = v186
		}
		v164 = v174
	}
	r7 = v164
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e7 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r7 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1201, libc.VaList(bp+24, r7, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e7&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1279, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if uint32(8) == uint32(4) {
		v194 = libc.Float64FromFloat64(1)
		v195 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v194
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v199 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _200
	_200:
		if libc.BoolInt32(v199&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v195
			v198 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v195
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v201 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _202
		_202:
			v198 = libc.BoolInt32(v201&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v196 = libc.BoolInt32(!(v198 != 0) && v194 <= v195)
		goto _197
	_197:
		v193 = v196
	} else {
		if uint32(8) == uint32(8) {
			v204 = libc.Float64FromFloat64(1)
			v205 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v204
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v209 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _210
		_210:
			if libc.BoolInt32(v209&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v205
				v208 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v205
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v211 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _212
			_212:
				v208 = libc.BoolInt32(v211&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v206 = libc.BoolInt32(!(v208 != 0) && v204 <= v205)
			goto _207
		_207:
			v203 = v206
		} else {
			v213 = libc.Float64FromFloat64(1)
			v214 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:113:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v213
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v218 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _219
		_219:
			if libc.BoolInt32(v218&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v214
				v217 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v214
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v220 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _221
			_221:
				v217 = libc.BoolInt32(v220&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v215 = libc.BoolInt32(!(v217 != 0) && v213 <= v214)
			goto _216
		_216:
			v203 = v215
		}
		v193 = v203
	}
	r8 = v193
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e8 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r8 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1363, libc.VaList(bp+24, r8, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e8&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1446, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if uint32(8) == uint32(4) {
		v223 = libc.Float64FromFloat64(1)
		v224 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v223
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v228 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _229
	_229:
		if libc.BoolInt32(v228&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v224
			v227 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v224
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v230 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _231
		_231:
			v227 = libc.BoolInt32(v230&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v225 = libc.BoolInt32(!(v227 != 0) && v223 != v224)
		goto _226
	_226:
		v222 = v225
	} else {
		if uint32(8) == uint32(8) {
			v233 = libc.Float64FromFloat64(1)
			v234 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v233
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v238 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _239
		_239:
			if libc.BoolInt32(v238&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v234
				v237 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v234
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v240 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _241
			_241:
				v237 = libc.BoolInt32(v240&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v235 = libc.BoolInt32(!(v237 != 0) && v233 != v234)
			goto _236
		_236:
			v232 = v235
		} else {
			v242 = libc.Float64FromFloat64(1)
			v243 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:116:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v242
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v247 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _248
		_248:
			if libc.BoolInt32(v247&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v243
				v246 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v243
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v249 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _250
			_250:
				v246 = libc.BoolInt32(v249&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v244 = libc.BoolInt32(!(v246 != 0) && v242 != v243)
			goto _245
		_245:
			v232 = v244
		}
		v222 = v232
	}
	r9 = v222
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e9 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r9 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1535, libc.VaList(bp+24, r9, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e9&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1620, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if uint32(8) == uint32(4) {
		v252 = libc.Float64FromFloat64(1)
		v253 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v252
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v257 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _258
	_258:
		if libc.BoolInt32(v257&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v253
			v256 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v253
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v259 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _260
		_260:
			v256 = libc.BoolInt32(v259&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v254 = libc.BoolInt32(!(v256 != 0) && v252 > v253)
		goto _255
	_255:
		v251 = v254
	} else {
		if uint32(8) == uint32(8) {
			v262 = libc.Float64FromFloat64(1)
			v263 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v262
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v267 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _268
		_268:
			if libc.BoolInt32(v267&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v263
				v266 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v263
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v269 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _270
			_270:
				v266 = libc.BoolInt32(v269&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v264 = libc.BoolInt32(!(v266 != 0) && v262 > v263)
			goto _265
		_265:
			v261 = v264
		} else {
			v271 = libc.Float64FromFloat64(1)
			v272 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:119:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v271
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v276 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _277
		_277:
			if libc.BoolInt32(v276&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v272
				v275 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v272
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v278 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _279
			_279:
				v275 = libc.BoolInt32(v278&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v273 = libc.BoolInt32(!(v275 != 0) && v271 > v272)
			goto _274
		_274:
			v261 = v273
		}
		v251 = v261
	}
	r10 = v251
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e10 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r10 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1711, libc.VaList(bp+24, r10, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e10&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1792, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if uint32(8) == uint32(4) {
		v281 = libc.Float64FromFloat64(1)
		v282 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v281
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v286 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _287
	_287:
		if libc.BoolInt32(v286&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v282
			v285 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v282
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v288 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _289
		_289:
			v285 = libc.BoolInt32(v288&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v283 = libc.BoolInt32(!(v285 != 0) && v281 >= v282)
		goto _284
	_284:
		v280 = v283
	} else {
		if uint32(8) == uint32(8) {
			v291 = libc.Float64FromFloat64(1)
			v292 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v291
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v296 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _297
		_297:
			if libc.BoolInt32(v296&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v292
				v295 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v292
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v298 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _299
			_299:
				v295 = libc.BoolInt32(v298&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v293 = libc.BoolInt32(!(v295 != 0) && v291 >= v292)
			goto _294
		_294:
			v290 = v293
		} else {
			v300 = libc.Float64FromFloat64(1)
			v301 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:122:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v300
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v305 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _306
		_306:
			if libc.BoolInt32(v305&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v301
				v304 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v301
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v307 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _308
			_308:
				v304 = libc.BoolInt32(v307&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v302 = libc.BoolInt32(!(v304 != 0) && v300 >= v301)
			goto _303
		_303:
			v290 = v302
		}
		v280 = v290
	}
	r11 = v280
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	e11 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if r11 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1879, libc.VaList(bp+24, r11, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
	if e11&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:42:2:
		t_printf(tls, __ccgo_ts+1965, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v311 = *(*uint32)(unsafe.Pointer(bp))
		goto _312
	_312:
		v310 = libc.BoolInt32(v311&uint32(0x7fffffff) > uint32(0x7f800000))
	} else {
		if uint32(4) == uint32(8) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v314 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _315
		_315:
			v313 = libc.BoolInt32(v314&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		} else {
			v313 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
		}
		v310 = v313
	}
	if v310 != 0 {
		libc.X__builtin_nanf(tls, __ccgo_ts)
		v309 = libc.Int32FromInt32(1)
	} else {
		if uint32(4) == uint32(4) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
			*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
			v317 = *(*uint32)(unsafe.Pointer(bp))
			goto _318
		_318:
			v316 = libc.BoolInt32(v317&uint32(0x7fffffff) > uint32(0x7f800000))
		} else {
			if uint32(4) == uint32(8) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v320 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _321
			_321:
				v319 = libc.BoolInt32(v320&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			} else {
				v319 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
			}
			v316 = v319
		}
		v309 = v316
	}
	r12 = v309
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e12 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r12 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2057, libc.VaList(bp+24, r12, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e12&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2155, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		v323 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v324 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v323
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v328 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _329
	_329:
		if libc.BoolInt32(v328&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v324
			v327 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v324
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v330 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _331
		_331:
			v327 = libc.BoolInt32(v330&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v325 = libc.BoolInt32(!(v327 != 0) && v323 < v324)
		goto _326
	_326:
		v322 = v325
	} else {
		if uint32(4) == uint32(8) {
			v333 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v334 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v333
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v338 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _339
		_339:
			if libc.BoolInt32(v338&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v334
				v337 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v334
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v340 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _341
			_341:
				v337 = libc.BoolInt32(v340&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v335 = libc.BoolInt32(!(v337 != 0) && v333 < v334)
			goto _336
		_336:
			v332 = v335
		} else {
			v342 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v343 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:110:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v342
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v347 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _348
		_348:
			if libc.BoolInt32(v347&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v343
				v346 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v343
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v349 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _350
			_350:
				v346 = libc.BoolInt32(v349&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v344 = libc.BoolInt32(!(v346 != 0) && v342 < v343)
			goto _345
		_345:
			v332 = v344
		}
		v322 = v332
	}
	r13 = v322
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e13 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r13 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2259, libc.VaList(bp+24, r13, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e13&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2352, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		v352 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v353 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v352
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v357 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _358
	_358:
		if libc.BoolInt32(v357&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v353
			v356 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v353
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v359 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _360
		_360:
			v356 = libc.BoolInt32(v359&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v354 = libc.BoolInt32(!(v356 != 0) && v352 <= v353)
		goto _355
	_355:
		v351 = v354
	} else {
		if uint32(4) == uint32(8) {
			v362 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v363 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v362
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v367 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _368
		_368:
			if libc.BoolInt32(v367&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v363
				v366 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v363
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v369 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _370
			_370:
				v366 = libc.BoolInt32(v369&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v364 = libc.BoolInt32(!(v366 != 0) && v362 <= v363)
			goto _365
		_365:
			v361 = v364
		} else {
			v371 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v372 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:113:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v371
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v376 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _377
		_377:
			if libc.BoolInt32(v376&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v372
				v375 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v372
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v378 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _379
			_379:
				v375 = libc.BoolInt32(v378&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v373 = libc.BoolInt32(!(v375 != 0) && v371 <= v372)
			goto _374
		_374:
			v361 = v373
		}
		v351 = v361
	}
	r14 = v351
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e14 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r14 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2451, libc.VaList(bp+24, r14, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e14&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2549, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		v381 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v382 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v381
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v386 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _387
	_387:
		if libc.BoolInt32(v386&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v382
			v385 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v382
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v388 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _389
		_389:
			v385 = libc.BoolInt32(v388&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v383 = libc.BoolInt32(!(v385 != 0) && v381 != v382)
		goto _384
	_384:
		v380 = v383
	} else {
		if uint32(4) == uint32(8) {
			v391 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v392 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v391
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v396 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _397
		_397:
			if libc.BoolInt32(v396&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v392
				v395 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v392
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v398 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _399
			_399:
				v395 = libc.BoolInt32(v398&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v393 = libc.BoolInt32(!(v395 != 0) && v391 != v392)
			goto _394
		_394:
			v390 = v393
		} else {
			v400 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v401 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:116:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v400
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v405 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _406
		_406:
			if libc.BoolInt32(v405&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v401
				v404 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v401
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v407 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _408
			_408:
				v404 = libc.BoolInt32(v407&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v402 = libc.BoolInt32(!(v404 != 0) && v400 != v401)
			goto _403
		_403:
			v390 = v402
		}
		v380 = v390
	}
	r15 = v380
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e15 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r15 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2653, libc.VaList(bp+24, r15, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e15&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2753, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		v410 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v411 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v410
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v415 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _416
	_416:
		if libc.BoolInt32(v415&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v411
			v414 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v411
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v417 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _418
		_418:
			v414 = libc.BoolInt32(v417&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v412 = libc.BoolInt32(!(v414 != 0) && v410 > v411)
		goto _413
	_413:
		v409 = v412
	} else {
		if uint32(4) == uint32(8) {
			v420 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v421 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v420
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v425 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _426
		_426:
			if libc.BoolInt32(v425&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v421
				v424 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v421
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v427 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _428
			_428:
				v424 = libc.BoolInt32(v427&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v422 = libc.BoolInt32(!(v424 != 0) && v420 > v421)
			goto _423
		_423:
			v419 = v422
		} else {
			v429 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v430 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:119:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v429
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v434 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _435
		_435:
			if libc.BoolInt32(v434&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v430
				v433 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v430
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v436 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _437
			_437:
				v433 = libc.BoolInt32(v436&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v431 = libc.BoolInt32(!(v433 != 0) && v429 > v430)
			goto _432
		_432:
			v419 = v431
		}
		v409 = v419
	}
	r16 = v409
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e16 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r16 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2859, libc.VaList(bp+24, r16, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e16&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+2955, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if uint32(4) == uint32(4) {
		v439 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v440 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v439
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v444 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _445
	_445:
		if libc.BoolInt32(v444&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v440
			v443 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v440
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v446 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _447
		_447:
			v443 = libc.BoolInt32(v446&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v441 = libc.BoolInt32(!(v443 != 0) && v439 >= v440)
		goto _442
	_442:
		v438 = v441
	} else {
		if uint32(4) == uint32(8) {
			v449 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v450 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v449
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v454 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _455
		_455:
			if libc.BoolInt32(v454&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v450
				v453 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v450
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v456 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _457
			_457:
				v453 = libc.BoolInt32(v456&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v451 = libc.BoolInt32(!(v453 != 0) && v449 >= v450)
			goto _452
		_452:
			v448 = v451
		} else {
			v458 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v459 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:122:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v458
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v463 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _464
		_464:
			if libc.BoolInt32(v463&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v459
				v462 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v459
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v465 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _466
			_466:
				v462 = libc.BoolInt32(v465&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v460 = libc.BoolInt32(!(v462 != 0) && v458 >= v459)
			goto _461
		_461:
			v448 = v460
		}
		v438 = v448
	}
	r17 = v438
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	e17 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if r17 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+3057, libc.VaList(bp+24, r17, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
	if e17&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:43:2:
		t_printf(tls, __ccgo_ts+3158, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(4) == uint32(4) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v469 = *(*uint32)(unsafe.Pointer(bp))
		goto _470
	_470:
		v468 = libc.BoolInt32(v469&uint32(0x7fffffff) > uint32(0x7f800000))
	} else {
		if uint32(4) == uint32(8) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v472 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _473
		_473:
			v471 = libc.BoolInt32(v472&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		} else {
			v471 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
		}
		v468 = v471
	}
	if v468 != 0 {
		_ = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		v467 = libc.Int32FromInt32(1)
	} else {
		if uint32(8) == uint32(4) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
			*(*float32)(unsafe.Pointer(bp)) = float32(float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
			v475 = *(*uint32)(unsafe.Pointer(bp))
			goto _476
		_476:
			v474 = libc.BoolInt32(v475&uint32(0x7fffffff) > uint32(0x7f800000))
		} else {
			if uint32(8) == uint32(8) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v478 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _479
			_479:
				v477 = libc.BoolInt32(v478&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			} else {
				v477 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))+libc.Float64FromFloat64(1)) == FP_NAN)
			}
			v474 = v477
		}
		v467 = v474
	}
	r18 = v467
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e18 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r18 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3265, libc.VaList(bp+24, r18, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e18&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3367, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(8) == uint32(4) {
		v481 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v482 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v481
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v486 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _487
	_487:
		if libc.BoolInt32(v486&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v482
			v485 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v482
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v488 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _489
		_489:
			v485 = libc.BoolInt32(v488&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v483 = libc.BoolInt32(!(v485 != 0) && v481 < v482)
		goto _484
	_484:
		v480 = v483
	} else {
		if uint32(8) == uint32(8) {
			v491 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v492 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v491
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v496 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _497
		_497:
			if libc.BoolInt32(v496&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v492
				v495 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v492
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v498 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _499
			_499:
				v495 = libc.BoolInt32(v498&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v493 = libc.BoolInt32(!(v495 != 0) && v491 < v492)
			goto _494
		_494:
			v490 = v493
		} else {
			v500 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v501 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:110:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v500
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v505 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _506
		_506:
			if libc.BoolInt32(v505&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v501
				v504 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v501
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v507 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _508
			_508:
				v504 = libc.BoolInt32(v507&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v502 = libc.BoolInt32(!(v504 != 0) && v500 < v501)
			goto _503
		_503:
			v490 = v502
		}
		v480 = v490
	}
	r19 = v480
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e19 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r19 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3475, libc.VaList(bp+24, r19, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e19&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3572, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(8) == uint32(4) {
		v510 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v511 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v510
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v515 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _516
	_516:
		if libc.BoolInt32(v515&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v511
			v514 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v511
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v517 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _518
		_518:
			v514 = libc.BoolInt32(v517&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v512 = libc.BoolInt32(!(v514 != 0) && v510 <= v511)
		goto _513
	_513:
		v509 = v512
	} else {
		if uint32(8) == uint32(8) {
			v520 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v521 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v520
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v525 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _526
		_526:
			if libc.BoolInt32(v525&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v521
				v524 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v521
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v527 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _528
			_528:
				v524 = libc.BoolInt32(v527&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v522 = libc.BoolInt32(!(v524 != 0) && v520 <= v521)
			goto _523
		_523:
			v519 = v522
		} else {
			v529 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v530 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:113:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v529
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v534 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _535
		_535:
			if libc.BoolInt32(v534&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v530
				v533 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v530
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v536 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _537
			_537:
				v533 = libc.BoolInt32(v536&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v531 = libc.BoolInt32(!(v533 != 0) && v529 <= v530)
			goto _532
		_532:
			v519 = v531
		}
		v509 = v519
	}
	r20 = v509
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e20 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r20 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3675, libc.VaList(bp+24, r20, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e20&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3777, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(8) == uint32(4) {
		v539 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v540 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v539
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v544 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _545
	_545:
		if libc.BoolInt32(v544&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v540
			v543 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v540
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v546 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _547
		_547:
			v543 = libc.BoolInt32(v546&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v541 = libc.BoolInt32(!(v543 != 0) && v539 != v540)
		goto _542
	_542:
		v538 = v541
	} else {
		if uint32(8) == uint32(8) {
			v549 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v550 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v549
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v554 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _555
		_555:
			if libc.BoolInt32(v554&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v550
				v553 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v550
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v556 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _557
			_557:
				v553 = libc.BoolInt32(v556&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v551 = libc.BoolInt32(!(v553 != 0) && v549 != v550)
			goto _552
		_552:
			v548 = v551
		} else {
			v558 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v559 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:116:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v558
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v563 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _564
		_564:
			if libc.BoolInt32(v563&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v559
				v562 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v559
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v565 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _566
			_566:
				v562 = libc.BoolInt32(v565&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v560 = libc.BoolInt32(!(v562 != 0) && v558 != v559)
			goto _561
		_561:
			v548 = v560
		}
		v538 = v548
	}
	r21 = v538
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e21 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r21 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3885, libc.VaList(bp+24, r21, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e21&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+3989, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(8) == uint32(4) {
		v568 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v569 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v568
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v573 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _574
	_574:
		if libc.BoolInt32(v573&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v569
			v572 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v569
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v575 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _576
		_576:
			v572 = libc.BoolInt32(v575&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v570 = libc.BoolInt32(!(v572 != 0) && v568 > v569)
		goto _571
	_571:
		v567 = v570
	} else {
		if uint32(8) == uint32(8) {
			v578 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v579 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v578
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v583 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _584
		_584:
			if libc.BoolInt32(v583&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v579
				v582 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v579
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v585 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _586
			_586:
				v582 = libc.BoolInt32(v585&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v580 = libc.BoolInt32(!(v582 != 0) && v578 > v579)
			goto _581
		_581:
			v577 = v580
		} else {
			v587 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v588 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:119:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v587
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v592 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _593
		_593:
			if libc.BoolInt32(v592&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v588
				v591 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v588
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v594 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _595
			_595:
				v591 = libc.BoolInt32(v594&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v589 = libc.BoolInt32(!(v591 != 0) && v587 > v588)
			goto _590
		_590:
			v577 = v589
		}
		v567 = v577
	}
	r22 = v567
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e22 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r22 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+4099, libc.VaList(bp+24, r22, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e22&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+4199, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if uint32(8) == uint32(4) {
		v597 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v598 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v597
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v602 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _603
	_603:
		if libc.BoolInt32(v602&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v598
			v601 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v598
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v604 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _605
		_605:
			v601 = libc.BoolInt32(v604&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v599 = libc.BoolInt32(!(v601 != 0) && v597 >= v598)
		goto _600
	_600:
		v596 = v599
	} else {
		if uint32(8) == uint32(8) {
			v607 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v608 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v607
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v612 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _613
		_613:
			if libc.BoolInt32(v612&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v608
				v611 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v608
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v614 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _615
			_615:
				v611 = libc.BoolInt32(v614&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v609 = libc.BoolInt32(!(v611 != 0) && v607 >= v608)
			goto _610
		_610:
			v606 = v609
		} else {
			v616 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v617 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + float64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:122:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v616
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v621 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _622
		_622:
			if libc.BoolInt32(v621&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v617
				v620 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v617
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v623 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _624
			_624:
				v620 = libc.BoolInt32(v623&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v618 = libc.BoolInt32(!(v620 != 0) && v616 >= v617)
			goto _619
		_619:
			v606 = v618
		}
		v596 = v606
	}
	r23 = v596
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	e23 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if r23 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+4305, libc.VaList(bp+24, r23, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
	if e23&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:44:2:
		t_printf(tls, __ccgo_ts+4410, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(4) == uint32(4) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.X__builtin_nanf(tls, __ccgo_ts)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v627 = *(*uint32)(unsafe.Pointer(bp))
		goto _628
	_628:
		v626 = libc.BoolInt32(v627&uint32(0x7fffffff) > uint32(0x7f800000))
	} else {
		if uint32(4) == uint32(8) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v630 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _631
		_631:
			v629 = libc.BoolInt32(v630&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		} else {
			v629 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))) == FP_NAN)
		}
		v626 = v629
	}
	if v626 != 0 {
		_ = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		v625 = libc.Int32FromInt32(1)
	} else {
		if uint32(8) == uint32(4) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
			*(*float32)(unsafe.Pointer(bp)) = float32(float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
			v633 = *(*uint32)(unsafe.Pointer(bp))
			goto _634
		_634:
			v632 = libc.BoolInt32(v633&uint32(0x7fffffff) > uint32(0x7f800000))
		} else {
			if uint32(8) == uint32(8) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v636 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _637
			_637:
				v635 = libc.BoolInt32(v636&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			} else {
				v635 = libc.BoolInt32(_fpclassifyl(tls, float64(libc.X__builtin_nanf(tls, __ccgo_ts))+libc.Float64FromFloat64(1)) == FP_NAN)
			}
			v632 = v635
		}
		v625 = v632
	}
	r24 = v625
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e24 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r24 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+4521, libc.VaList(bp+24, r24, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e24&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+4624, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(8) == uint32(4) {
		v639 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v640 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v639
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v644 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _645
	_645:
		if libc.BoolInt32(v644&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v640
			v643 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v640
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v646 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _647
		_647:
			v643 = libc.BoolInt32(v646&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v641 = libc.BoolInt32(!(v643 != 0) && v639 < v640)
		goto _642
	_642:
		v638 = v641
	} else {
		if uint32(8) == uint32(8) {
			v649 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v650 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v649
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v654 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _655
		_655:
			if libc.BoolInt32(v654&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v650
				v653 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v650
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v656 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _657
			_657:
				v653 = libc.BoolInt32(v656&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v651 = libc.BoolInt32(!(v653 != 0) && v649 < v650)
			goto _652
		_652:
			v648 = v651
		} else {
			v658 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v659 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:110:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v658
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v663 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _664
		_664:
			if libc.BoolInt32(v663&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v659
				v662 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v659
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v665 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _666
			_666:
				v662 = libc.BoolInt32(v665&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v660 = libc.BoolInt32(!(v662 != 0) && v658 < v659)
			goto _661
		_661:
			v648 = v660
		}
		v638 = v648
	}
	r25 = v638
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e25 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r25 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+4733, libc.VaList(bp+24, r25, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e25&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+4831, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(8) == uint32(4) {
		v668 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v669 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v668
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v673 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _674
	_674:
		if libc.BoolInt32(v673&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v669
			v672 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v669
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v675 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _676
		_676:
			v672 = libc.BoolInt32(v675&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v670 = libc.BoolInt32(!(v672 != 0) && v668 <= v669)
		goto _671
	_671:
		v667 = v670
	} else {
		if uint32(8) == uint32(8) {
			v678 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v679 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v678
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v683 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _684
		_684:
			if libc.BoolInt32(v683&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v679
				v682 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v679
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v685 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _686
			_686:
				v682 = libc.BoolInt32(v685&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v680 = libc.BoolInt32(!(v682 != 0) && v678 <= v679)
			goto _681
		_681:
			v677 = v680
		} else {
			v687 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v688 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:113:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v687
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v692 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _693
		_693:
			if libc.BoolInt32(v692&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v688
				v691 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v688
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v694 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _695
			_695:
				v691 = libc.BoolInt32(v694&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v689 = libc.BoolInt32(!(v691 != 0) && v687 <= v688)
			goto _690
		_690:
			v677 = v689
		}
		v667 = v677
	}
	r26 = v667
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e26 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r26 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+4935, libc.VaList(bp+24, r26, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e26&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5038, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(8) == uint32(4) {
		v697 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v698 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v697
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v702 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _703
	_703:
		if libc.BoolInt32(v702&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v698
			v701 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v698
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v704 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _705
		_705:
			v701 = libc.BoolInt32(v704&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v699 = libc.BoolInt32(!(v701 != 0) && v697 != v698)
		goto _700
	_700:
		v696 = v699
	} else {
		if uint32(8) == uint32(8) {
			v707 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v708 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v707
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v712 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _713
		_713:
			if libc.BoolInt32(v712&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v708
				v711 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v708
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v714 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _715
			_715:
				v711 = libc.BoolInt32(v714&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v709 = libc.BoolInt32(!(v711 != 0) && v707 != v708)
			goto _710
		_710:
			v706 = v709
		} else {
			v716 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v717 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:116:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v716
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v721 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _722
		_722:
			if libc.BoolInt32(v721&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v717
				v720 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v717
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v723 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _724
			_724:
				v720 = libc.BoolInt32(v723&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v718 = libc.BoolInt32(!(v720 != 0) && v716 != v717)
			goto _719
		_719:
			v706 = v718
		}
		v696 = v706
	}
	r27 = v696
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e27 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r27 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5147, libc.VaList(bp+24, r27, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e27&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5252, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(8) == uint32(4) {
		v726 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v727 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v726
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v731 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _732
	_732:
		if libc.BoolInt32(v731&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v727
			v730 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v727
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v733 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _734
		_734:
			v730 = libc.BoolInt32(v733&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v728 = libc.BoolInt32(!(v730 != 0) && v726 > v727)
		goto _729
	_729:
		v725 = v728
	} else {
		if uint32(8) == uint32(8) {
			v736 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v737 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v736
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v741 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _742
		_742:
			if libc.BoolInt32(v741&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v737
				v740 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v737
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v743 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _744
			_744:
				v740 = libc.BoolInt32(v743&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v738 = libc.BoolInt32(!(v740 != 0) && v736 > v737)
			goto _739
		_739:
			v735 = v738
		} else {
			v745 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v746 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:119:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v745
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v750 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _751
		_751:
			if libc.BoolInt32(v750&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v746
				v749 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v746
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v752 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _753
			_753:
				v749 = libc.BoolInt32(v752&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v747 = libc.BoolInt32(!(v749 != 0) && v745 > v746)
			goto _748
		_748:
			v735 = v747
		}
		v725 = v735
	}
	r28 = v725
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e28 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r28 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5363, libc.VaList(bp+24, r28, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e28&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5464, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if uint32(8) == uint32(4) {
		v755 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
		v756 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v755
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v760 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _761
	_761:
		if libc.BoolInt32(v760&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
			_ = v756
			v759 = libc.Int32FromInt32(1)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v756
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v762 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _763
		_763:
			v759 = libc.BoolInt32(v762&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
		}
		v757 = libc.BoolInt32(!(v759 != 0) && v755 >= v756)
		goto _758
	_758:
		v754 = v757
	} else {
		if uint32(8) == uint32(8) {
			v765 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v766 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v765
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v770 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _771
		_771:
			if libc.BoolInt32(v770&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v766
				v769 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v766
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v772 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _773
			_773:
				v769 = libc.BoolInt32(v772&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v767 = libc.BoolInt32(!(v769 != 0) && v765 >= v766)
			goto _768
		_768:
			v764 = v767
		} else {
			v774 = float64(libc.X__builtin_nanf(tls, __ccgo_ts))
			v775 = float64(libc.X__builtin_nanf(tls, __ccgo_ts)) + libc.Float64FromFloat64(1)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:122:1:
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
			*(*float64)(unsafe.Pointer(bp + 8)) = v774
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
			v779 = *(*uint64)(unsafe.Pointer(bp + 8))
			goto _780
		_780:
			if libc.BoolInt32(v779&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
				_ = v775
				v778 = libc.Int32FromInt32(1)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
				*(*float64)(unsafe.Pointer(bp + 8)) = v775
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
				v781 = *(*uint64)(unsafe.Pointer(bp + 8))
				goto _782
			_782:
				v778 = libc.BoolInt32(v781&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
			}
			v776 = libc.BoolInt32(!(v778 != 0) && v774 >= v775)
			goto _777
		_777:
			v764 = v776
		}
		v754 = v764
	}
	r29 = v754
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	e29 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if r29 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5571, libc.VaList(bp+24, r29, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
	if e29&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:45:2:
		t_printf(tls, __ccgo_ts+5677, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v784 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _785
_785:
	if libc.BoolInt32(v784&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1.1)
		v783 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(1.1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v786 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _787
	_787:
		v783 = libc.BoolInt32(v786&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r30 = v783
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e30 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r30 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+5789, libc.VaList(bp+24, r30, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e30&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+5857, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	v788 = libc.Float64FromFloat64(1)
	v789 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v788
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v793 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _794
_794:
	if libc.BoolInt32(v793&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v789
		v792 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v789
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v795 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _796
	_796:
		v792 = libc.BoolInt32(v795&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v790 = libc.BoolInt32(!(v792 != 0) && v788 < v789)
	goto _791
_791:
	r31 = v790
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e31 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r31 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+5931, libc.VaList(bp+24, r31, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e31&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+5994, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	v797 = libc.Float64FromFloat64(1)
	v798 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v797
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v802 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _803
_803:
	if libc.BoolInt32(v802&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v798
		v801 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v798
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v804 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _805
	_805:
		v801 = libc.BoolInt32(v804&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v799 = libc.BoolInt32(!(v801 != 0) && v797 <= v798)
	goto _800
_800:
	r32 = v799
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e32 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r32 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6063, libc.VaList(bp+24, r32, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e32&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6131, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	v806 = libc.Float64FromFloat64(1)
	v807 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v806
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v811 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _812
_812:
	if libc.BoolInt32(v811&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v807
		v810 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v807
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v813 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _814
	_814:
		v810 = libc.BoolInt32(v813&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v808 = libc.BoolInt32(!(v810 != 0) && v806 != v807)
	goto _809
_809:
	r33 = v808
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e33 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r33 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6205, libc.VaList(bp+24, r33, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e33&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6275, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	v815 = libc.Float64FromFloat64(1)
	v816 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v815
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v820 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _821
_821:
	if libc.BoolInt32(v820&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v816
		v819 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v816
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v822 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _823
	_823:
		v819 = libc.BoolInt32(v822&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v817 = libc.BoolInt32(!(v819 != 0) && v815 > v816)
	goto _818
_818:
	r34 = v817
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e34 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r34 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6351, libc.VaList(bp+24, r34, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e34&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6417, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	v824 = libc.Float64FromFloat64(1)
	v825 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v824
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v829 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _830
_830:
	if libc.BoolInt32(v829&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v825
		v828 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v825
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v831 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _832
	_832:
		v828 = libc.BoolInt32(v831&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v826 = libc.BoolInt32(!(v828 != 0) && v824 >= v825)
	goto _827
_827:
	r35 = v826
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	e35 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if r35 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6489, libc.VaList(bp+24, r35, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
	if e35&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:47:2:
		t_printf(tls, __ccgo_ts+6560, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v834 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _835
_835:
	if libc.BoolInt32(v834&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1) + eps
		v833 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(1) + eps
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v836 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _837
	_837:
		v833 = libc.BoolInt32(v836&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r36 = v833
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e36 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r36 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6637, libc.VaList(bp+24, r36, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e36&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6709, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	v838 = libc.Float64FromFloat64(1)
	v839 = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v838
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v843 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _844
_844:
	if libc.BoolInt32(v843&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v839
		v842 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v839
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v845 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _846
	_846:
		v842 = libc.BoolInt32(v845&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v840 = libc.BoolInt32(!(v842 != 0) && v838 < v839)
	goto _841
_841:
	r37 = v840
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e37 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r37 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6787, libc.VaList(bp+24, r37, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e37&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6854, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	v847 = libc.Float64FromFloat64(1)
	v848 = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v847
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v852 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _853
_853:
	if libc.BoolInt32(v852&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v848
		v851 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v848
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v854 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _855
	_855:
		v851 = libc.BoolInt32(v854&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v849 = libc.BoolInt32(!(v851 != 0) && v847 <= v848)
	goto _850
_850:
	r38 = v849
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e38 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r38 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6927, libc.VaList(bp+24, r38, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e38&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+6999, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	v856 = libc.Float64FromFloat64(1)
	v857 = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v856
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v861 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _862
_862:
	if libc.BoolInt32(v861&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v857
		v860 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v857
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v863 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _864
	_864:
		v860 = libc.BoolInt32(v863&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v858 = libc.BoolInt32(!(v860 != 0) && v856 != v857)
	goto _859
_859:
	r39 = v858
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e39 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r39 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7077, libc.VaList(bp+24, r39, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e39&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7151, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	v865 = libc.Float64FromFloat64(1)
	v866 = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v865
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v870 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _871
_871:
	if libc.BoolInt32(v870&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v866
		v869 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v866
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v872 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _873
	_873:
		v869 = libc.BoolInt32(v872&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v867 = libc.BoolInt32(!(v869 != 0) && v865 > v866)
	goto _868
_868:
	r40 = v867
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e40 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r40 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7231, libc.VaList(bp+24, r40, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e40&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7301, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	v874 = libc.Float64FromFloat64(1)
	v875 = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v874
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v879 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _880
_880:
	if libc.BoolInt32(v879&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v875
		v878 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v875
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v881 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _882
	_882:
		v878 = libc.BoolInt32(v881&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v876 = libc.BoolInt32(!(v878 != 0) && v874 >= v875)
	goto _877
_877:
	r41 = v876
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	e41 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if r41 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7377, libc.VaList(bp+24, r41, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
	if e41&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:48:2:
		t_printf(tls, __ccgo_ts+7452, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = float64(1) + eps
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v884 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _885
_885:
	if libc.BoolInt32(v884&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1)
		v883 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v886 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _887
	_887:
		v883 = libc.BoolInt32(v886&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r42 = v883
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e42 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r42 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7533, libc.VaList(bp+24, r42, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e42&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7605, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	v888 = float64(1) + eps
	v889 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v888
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v893 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _894
_894:
	if libc.BoolInt32(v893&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v889
		v892 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v889
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v895 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _896
	_896:
		v892 = libc.BoolInt32(v895&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v890 = libc.BoolInt32(!(v892 != 0) && v888 < v889)
	goto _891
_891:
	r43 = v890
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e43 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r43 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7683, libc.VaList(bp+24, r43, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e43&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7750, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	v897 = float64(1) + eps
	v898 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v897
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v902 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _903
_903:
	if libc.BoolInt32(v902&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v898
		v901 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v898
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v904 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _905
	_905:
		v901 = libc.BoolInt32(v904&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v899 = libc.BoolInt32(!(v901 != 0) && v897 <= v898)
	goto _900
_900:
	r44 = v899
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e44 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r44 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7823, libc.VaList(bp+24, r44, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e44&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7895, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	v906 = float64(1) + eps
	v907 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v906
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v911 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _912
_912:
	if libc.BoolInt32(v911&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v907
		v910 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v907
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v913 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _914
	_914:
		v910 = libc.BoolInt32(v913&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v908 = libc.BoolInt32(!(v910 != 0) && v906 != v907)
	goto _909
_909:
	r45 = v908
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e45 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r45 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+7973, libc.VaList(bp+24, r45, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e45&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+8047, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	v915 = float64(1) + eps
	v916 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v915
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v920 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _921
_921:
	if libc.BoolInt32(v920&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v916
		v919 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v916
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v922 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _923
	_923:
		v919 = libc.BoolInt32(v922&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v917 = libc.BoolInt32(!(v919 != 0) && v915 > v916)
	goto _918
_918:
	r46 = v917
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e46 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r46 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+8127, libc.VaList(bp+24, r46, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e46&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+8197, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	v924 = float64(1) + eps
	v925 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v924
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v929 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _930
_930:
	if libc.BoolInt32(v929&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v925
		v928 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v925
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v931 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _932
	_932:
		v928 = libc.BoolInt32(v931&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v926 = libc.BoolInt32(!(v928 != 0) && v924 >= v925)
	goto _927
_927:
	r47 = v926
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	e47 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if r47 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+8273, libc.VaList(bp+24, r47, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
	if e47&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:49:2:
		t_printf(tls, __ccgo_ts+8348, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = huge - libc.Float64FromInt32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v934 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _935
_935:
	if libc.BoolInt32(v934&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = huge
		v933 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = huge
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v936 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _937
	_937:
		v933 = libc.BoolInt32(v936&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r48 = v933
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e48 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r48 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8429, libc.VaList(bp+24, r48, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e48&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8501, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	v938 = huge - libc.Float64FromInt32(1)
	v939 = float64(huge)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v938
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v943 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _944
_944:
	if libc.BoolInt32(v943&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v939
		v942 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v939
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v945 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _946
	_946:
		v942 = libc.BoolInt32(v945&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v940 = libc.BoolInt32(!(v942 != 0) && v938 < v939)
	goto _941
_941:
	r49 = v940
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e49 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r49 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8579, libc.VaList(bp+24, r49, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e49&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8646, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	v947 = huge - libc.Float64FromInt32(1)
	v948 = float64(huge)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v947
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v952 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _953
_953:
	if libc.BoolInt32(v952&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v948
		v951 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v948
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v954 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _955
	_955:
		v951 = libc.BoolInt32(v954&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v949 = libc.BoolInt32(!(v951 != 0) && v947 <= v948)
	goto _950
_950:
	r50 = v949
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e50 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r50 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8719, libc.VaList(bp+24, r50, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e50&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8791, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	v956 = huge - libc.Float64FromInt32(1)
	v957 = float64(huge)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v956
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v961 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _962
_962:
	if libc.BoolInt32(v961&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v957
		v960 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v957
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v963 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _964
	_964:
		v960 = libc.BoolInt32(v963&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v958 = libc.BoolInt32(!(v960 != 0) && v956 != v957)
	goto _959
_959:
	r51 = v958
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e51 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r51 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8869, libc.VaList(bp+24, r51, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e51&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+8943, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	v965 = huge - libc.Float64FromInt32(1)
	v966 = float64(huge)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v965
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v970 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _971
_971:
	if libc.BoolInt32(v970&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v966
		v969 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v966
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v972 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _973
	_973:
		v969 = libc.BoolInt32(v972&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v967 = libc.BoolInt32(!(v969 != 0) && v965 > v966)
	goto _968
_968:
	r52 = v967
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e52 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r52 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+9023, libc.VaList(bp+24, r52, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e52&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+9093, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	v974 = huge - libc.Float64FromInt32(1)
	v975 = float64(huge)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v974
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v979 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _980
_980:
	if libc.BoolInt32(v979&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v975
		v978 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v975
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v981 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _982
	_982:
		v978 = libc.BoolInt32(v981&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v976 = libc.BoolInt32(!(v978 != 0) && v974 >= v975)
	goto _977
_977:
	r53 = v976
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	e53 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if r53 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+9169, libc.VaList(bp+24, r53, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
	if e53&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:50:2:
		t_printf(tls, __ccgo_ts+9244, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v984 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _985
_985:
	if libc.BoolInt32(v984&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = huge * huge
		v983 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = huge * huge
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v986 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _987
	_987:
		v983 = libc.BoolInt32(v986&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r54 = v983
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e54 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r54 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9325, libc.VaList(bp+24, r54, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e54&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9400, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	v988 = float64(huge)
	v989 = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v988
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v993 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _994
_994:
	if libc.BoolInt32(v993&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v989
		v992 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v989
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v995 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _996
	_996:
		v992 = libc.BoolInt32(v995&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v990 = libc.BoolInt32(!(v992 != 0) && v988 < v989)
	goto _991
_991:
	r55 = v990
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e55 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r55 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9481, libc.VaList(bp+24, r55, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e55&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9551, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	v997 = float64(huge)
	v998 = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v997
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1002 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1003
_1003:
	if libc.BoolInt32(v1002&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v998
		v1001 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v998
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1004 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1005
	_1005:
		v1001 = libc.BoolInt32(v1004&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v999 = libc.BoolInt32(!(v1001 != 0) && v997 <= v998)
	goto _1000
_1000:
	r56 = v999
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e56 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r56 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9627, libc.VaList(bp+24, r56, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e56&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9702, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	v1006 = float64(huge)
	v1007 = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1006
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1011 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1012
_1012:
	if libc.BoolInt32(v1011&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1007
		v1010 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1007
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1013 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1014
	_1014:
		v1010 = libc.BoolInt32(v1013&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1008 = libc.BoolInt32(!(v1010 != 0) && v1006 != v1007)
	goto _1009
_1009:
	r57 = v1008
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e57 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r57 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9783, libc.VaList(bp+24, r57, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e57&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9860, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	v1015 = float64(huge)
	v1016 = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1015
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1020 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1021
_1021:
	if libc.BoolInt32(v1020&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1016
		v1019 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1016
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1022 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1023
	_1023:
		v1019 = libc.BoolInt32(v1022&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1017 = libc.BoolInt32(!(v1019 != 0) && v1015 > v1016)
	goto _1018
_1018:
	r58 = v1017
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e58 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r58 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+9943, libc.VaList(bp+24, r58, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e58&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+10016, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	v1024 = float64(huge)
	v1025 = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1024
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1029 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1030
_1030:
	if libc.BoolInt32(v1029&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1025
		v1028 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1025
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1031 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1032
	_1032:
		v1028 = libc.BoolInt32(v1031&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1026 = libc.BoolInt32(!(v1028 != 0) && v1024 >= v1025)
	goto _1027
_1027:
	r59 = v1026
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	e59 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if r59 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+10095, libc.VaList(bp+24, r59, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
	if e59&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:51:2:
		t_printf(tls, __ccgo_ts+10173, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1034 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1035
_1035:
	if libc.BoolInt32(v1034&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v1033 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1036 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1037
	_1037:
		v1033 = libc.BoolInt32(v1036&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r60 = v1033
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e60 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r60 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10257, libc.VaList(bp+24, r60, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e60&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10326, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	v1038 = -libc.Float64FromFloat64(0)
	v1039 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1038
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1043 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1044
_1044:
	if libc.BoolInt32(v1043&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1039
		v1042 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1039
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1045 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1046
	_1046:
		v1042 = libc.BoolInt32(v1045&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1040 = libc.BoolInt32(!(v1042 != 0) && v1038 < v1039)
	goto _1041
_1041:
	r61 = v1040
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e61 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r61 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10401, libc.VaList(bp+24, r61, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e61&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10465, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	v1047 = -libc.Float64FromFloat64(0)
	v1048 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1047
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1052 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1053
_1053:
	if libc.BoolInt32(v1052&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1048
		v1051 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1048
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1054 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1055
	_1055:
		v1051 = libc.BoolInt32(v1054&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1049 = libc.BoolInt32(!(v1051 != 0) && v1047 <= v1048)
	goto _1050
_1050:
	r62 = v1049
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e62 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r62 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10535, libc.VaList(bp+24, r62, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e62&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10604, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	v1056 = -libc.Float64FromFloat64(0)
	v1057 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1056
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1061 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1062
_1062:
	if libc.BoolInt32(v1061&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1057
		v1060 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1057
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1063 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1064
	_1064:
		v1060 = libc.BoolInt32(v1063&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1058 = libc.BoolInt32(!(v1060 != 0) && v1056 != v1057)
	goto _1059
_1059:
	r63 = v1058
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e63 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r63 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10679, libc.VaList(bp+24, r63, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e63&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10750, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	v1065 = -libc.Float64FromFloat64(0)
	v1066 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1065
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1070 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1071
_1071:
	if libc.BoolInt32(v1070&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1066
		v1069 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1066
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1072 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1073
	_1073:
		v1069 = libc.BoolInt32(v1072&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1067 = libc.BoolInt32(!(v1069 != 0) && v1065 > v1066)
	goto _1068
_1068:
	r64 = v1067
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e64 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r64 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10827, libc.VaList(bp+24, r64, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e64&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10894, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	v1074 = -libc.Float64FromFloat64(0)
	v1075 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1074
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1079 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1080
_1080:
	if libc.BoolInt32(v1079&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1075
		v1078 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1075
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1081 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1082
	_1082:
		v1078 = libc.BoolInt32(v1081&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1076 = libc.BoolInt32(!(v1078 != 0) && v1074 >= v1075)
	goto _1077
_1077:
	r65 = v1076
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	e65 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if r65 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+10967, libc.VaList(bp+24, r65, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
	if e65&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:52:2:
		t_printf(tls, __ccgo_ts+11039, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1084 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1085
_1085:
	if libc.BoolInt32(v1084&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v1083 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1086 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1087
	_1087:
		v1083 = libc.BoolInt32(v1086&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r66 = v1083
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e66 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r66 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11117, libc.VaList(bp+24, r66, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e66&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11187, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	v1088 = -tiny
	v1089 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1088
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1093 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1094
_1094:
	if libc.BoolInt32(v1093&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1089
		v1092 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1089
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1095 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1096
	_1096:
		v1092 = libc.BoolInt32(v1095&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1090 = libc.BoolInt32(!(v1092 != 0) && v1088 < v1089)
	goto _1091
_1091:
	r67 = v1090
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e67 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r67 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11263, libc.VaList(bp+24, r67, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e67&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11328, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	v1097 = -tiny
	v1098 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1097
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1102 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1103
_1103:
	if libc.BoolInt32(v1102&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1098
		v1101 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1098
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1104 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1105
	_1105:
		v1101 = libc.BoolInt32(v1104&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1099 = libc.BoolInt32(!(v1101 != 0) && v1097 <= v1098)
	goto _1100
_1100:
	r68 = v1099
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e68 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r68 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11399, libc.VaList(bp+24, r68, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e68&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11469, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	v1106 = -tiny
	v1107 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1106
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1111 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1112
_1112:
	if libc.BoolInt32(v1111&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1107
		v1110 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1107
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1113 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1114
	_1114:
		v1110 = libc.BoolInt32(v1113&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1108 = libc.BoolInt32(!(v1110 != 0) && v1106 != v1107)
	goto _1109
_1109:
	r69 = v1108
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e69 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r69 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11545, libc.VaList(bp+24, r69, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e69&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11617, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	v1115 = -tiny
	v1116 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1115
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1120 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1121
_1121:
	if libc.BoolInt32(v1120&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1116
		v1119 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1116
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1122 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1123
	_1123:
		v1119 = libc.BoolInt32(v1122&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1117 = libc.BoolInt32(!(v1119 != 0) && v1115 > v1116)
	goto _1118
_1118:
	r70 = v1117
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e70 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r70 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11695, libc.VaList(bp+24, r70, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e70&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11763, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	v1124 = -tiny
	v1125 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1124
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1129 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1130
_1130:
	if libc.BoolInt32(v1129&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1125
		v1128 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1125
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1131 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1132
	_1132:
		v1128 = libc.BoolInt32(v1131&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1126 = libc.BoolInt32(!(v1128 != 0) && v1124 >= v1125)
	goto _1127
_1127:
	r71 = v1126
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	e71 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if r71 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11837, libc.VaList(bp+24, r71, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
	if e71&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:53:2:
		t_printf(tls, __ccgo_ts+11910, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1134 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1135
_1135:
	if libc.BoolInt32(v1134&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromInt32(2) * tiny
		v1133 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromInt32(2) * tiny
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1136 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1137
	_1137:
		v1133 = libc.BoolInt32(v1136&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r72 = v1133
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e72 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r72 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+11989, libc.VaList(bp+24, r72, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e72&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12061, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	v1138 = float64(tiny)
	v1139 = libc.Float64FromInt32(2) * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1138
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1143 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1144
_1144:
	if libc.BoolInt32(v1143&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1139
		v1142 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1139
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1145 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1146
	_1146:
		v1142 = libc.BoolInt32(v1145&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1140 = libc.BoolInt32(!(v1142 != 0) && v1138 < v1139)
	goto _1141
_1141:
	r73 = v1140
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e73 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r73 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12139, libc.VaList(bp+24, r73, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e73&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12206, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	v1147 = float64(tiny)
	v1148 = libc.Float64FromInt32(2) * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1147
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1152 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1153
_1153:
	if libc.BoolInt32(v1152&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1148
		v1151 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1148
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1154 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1155
	_1155:
		v1151 = libc.BoolInt32(v1154&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1149 = libc.BoolInt32(!(v1151 != 0) && v1147 <= v1148)
	goto _1150
_1150:
	r74 = v1149
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e74 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r74 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12279, libc.VaList(bp+24, r74, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e74&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12351, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	v1156 = float64(tiny)
	v1157 = libc.Float64FromInt32(2) * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1156
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1161 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1162
_1162:
	if libc.BoolInt32(v1161&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1157
		v1160 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1157
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1163 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1164
	_1164:
		v1160 = libc.BoolInt32(v1163&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1158 = libc.BoolInt32(!(v1160 != 0) && v1156 != v1157)
	goto _1159
_1159:
	r75 = v1158
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e75 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r75 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12429, libc.VaList(bp+24, r75, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e75&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12503, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	v1165 = float64(tiny)
	v1166 = libc.Float64FromInt32(2) * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1165
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1170 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1171
_1171:
	if libc.BoolInt32(v1170&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1166
		v1169 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1166
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1172 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1173
	_1173:
		v1169 = libc.BoolInt32(v1172&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1167 = libc.BoolInt32(!(v1169 != 0) && v1165 > v1166)
	goto _1168
_1168:
	r76 = v1167
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e76 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r76 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12583, libc.VaList(bp+24, r76, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e76&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12653, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	v1174 = float64(tiny)
	v1175 = libc.Float64FromInt32(2) * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1174
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1179 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1180
_1180:
	if libc.BoolInt32(v1179&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1175
		v1178 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1175
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1181 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1182
	_1182:
		v1178 = libc.BoolInt32(v1181&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1176 = libc.BoolInt32(!(v1178 != 0) && v1174 >= v1175)
	goto _1177
_1177:
	r77 = v1176
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	e77 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if r77 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12729, libc.VaList(bp+24, r77, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
	if e77&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:54:2:
		t_printf(tls, __ccgo_ts+12804, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tiny * libc.Float64FromFloat64(0.001953125)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1184 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1185
_1185:
	if libc.BoolInt32(v1184&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = tiny * libc.Float64FromFloat64(0.00390625)
		v1183 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = tiny * libc.Float64FromFloat64(0.00390625)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1186 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1187
	_1187:
		v1183 = libc.BoolInt32(v1186&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r78 = v1183
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e78 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r78 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+12885, libc.VaList(bp+24, r78, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e78&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+12969, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	v1188 = tiny * float64(0.001953125)
	v1189 = tiny * float64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1188
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1193 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1194
_1194:
	if libc.BoolInt32(v1193&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1189
		v1192 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1189
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1195 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1196
	_1196:
		v1192 = libc.BoolInt32(v1195&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1190 = libc.BoolInt32(!(v1192 != 0) && v1188 < v1189)
	goto _1191
_1191:
	r79 = v1190
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e79 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r79 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13059, libc.VaList(bp+24, r79, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e79&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13138, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	v1197 = tiny * float64(0.001953125)
	v1198 = tiny * float64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1197
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1202 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1203
_1203:
	if libc.BoolInt32(v1202&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1198
		v1201 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1198
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1204 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1205
	_1205:
		v1201 = libc.BoolInt32(v1204&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1199 = libc.BoolInt32(!(v1201 != 0) && v1197 <= v1198)
	goto _1200
_1200:
	r80 = v1199
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e80 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r80 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13223, libc.VaList(bp+24, r80, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e80&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13307, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	v1206 = tiny * float64(0.001953125)
	v1207 = tiny * float64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1206
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1211 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1212
_1212:
	if libc.BoolInt32(v1211&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1207
		v1210 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1207
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1213 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1214
	_1214:
		v1210 = libc.BoolInt32(v1213&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1208 = libc.BoolInt32(!(v1210 != 0) && v1206 != v1207)
	goto _1209
_1209:
	r81 = v1208
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e81 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r81 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13397, libc.VaList(bp+24, r81, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e81&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13483, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	v1215 = tiny * float64(0.001953125)
	v1216 = tiny * float64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1215
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1220 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1221
_1221:
	if libc.BoolInt32(v1220&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1216
		v1219 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1216
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1222 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1223
	_1223:
		v1219 = libc.BoolInt32(v1222&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1217 = libc.BoolInt32(!(v1219 != 0) && v1215 > v1216)
	goto _1218
_1218:
	r82 = v1217
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e82 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r82 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13575, libc.VaList(bp+24, r82, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e82&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13657, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	v1224 = tiny * float64(0.001953125)
	v1225 = tiny * float64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1224
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1229 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1230
_1230:
	if libc.BoolInt32(v1229&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1225
		v1228 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1225
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1231 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1232
	_1232:
		v1228 = libc.BoolInt32(v1231&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1226 = libc.BoolInt32(!(v1228 != 0) && v1224 >= v1225)
	goto _1227
_1227:
	r83 = v1226
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	e83 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if r83 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13745, libc.VaList(bp+24, r83, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
	if e83&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:55:2:
		t_printf(tls, __ccgo_ts+13832, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1234 = *(*uint32)(unsafe.Pointer(bp))
	goto _1235
_1235:
	if libc.BoolInt32(v1234&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(1.1)
		v1233 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1.1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1236 = *(*uint32)(unsafe.Pointer(bp))
		goto _1237
	_1237:
		v1233 = libc.BoolInt32(v1236&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r84 = v1233
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e84 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r84 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+13925, libc.VaList(bp+24, r84, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e84&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+13995, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	v1238 = libc.Float64FromFloat32(1)
	v1239 = libc.Float64FromFloat32(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1238
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1243 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1244
_1244:
	if libc.BoolInt32(v1243&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1239
		v1242 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1239
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1245 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1246
	_1246:
		v1242 = libc.BoolInt32(v1245&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1240 = libc.BoolInt32(!(v1242 != 0) && v1238 < v1239)
	goto _1241
_1241:
	r85 = v1240
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e85 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r85 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14071, libc.VaList(bp+24, r85, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e85&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14136, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	v1247 = libc.Float64FromFloat32(1)
	v1248 = libc.Float64FromFloat32(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1247
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1252 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1253
_1253:
	if libc.BoolInt32(v1252&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1248
		v1251 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1248
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1254 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1255
	_1255:
		v1251 = libc.BoolInt32(v1254&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1249 = libc.BoolInt32(!(v1251 != 0) && v1247 <= v1248)
	goto _1250
_1250:
	r86 = v1249
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e86 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r86 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14207, libc.VaList(bp+24, r86, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e86&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14277, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	v1256 = libc.Float64FromFloat32(1)
	v1257 = libc.Float64FromFloat32(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1256
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1261 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1262
_1262:
	if libc.BoolInt32(v1261&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1257
		v1260 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1257
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1263 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1264
	_1264:
		v1260 = libc.BoolInt32(v1263&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1258 = libc.BoolInt32(!(v1260 != 0) && v1256 != v1257)
	goto _1259
_1259:
	r87 = v1258
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e87 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r87 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14353, libc.VaList(bp+24, r87, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e87&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14425, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	v1265 = libc.Float64FromFloat32(1)
	v1266 = libc.Float64FromFloat32(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1265
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1270 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1271
_1271:
	if libc.BoolInt32(v1270&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1266
		v1269 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1266
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1272 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1273
	_1273:
		v1269 = libc.BoolInt32(v1272&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1267 = libc.BoolInt32(!(v1269 != 0) && v1265 > v1266)
	goto _1268
_1268:
	r88 = v1267
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e88 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r88 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14503, libc.VaList(bp+24, r88, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e88&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14571, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	v1274 = libc.Float64FromFloat32(1)
	v1275 = libc.Float64FromFloat32(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1274
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1279 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1280
_1280:
	if libc.BoolInt32(v1279&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1275
		v1278 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1275
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1281 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1282
	_1282:
		v1278 = libc.BoolInt32(v1281&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1276 = libc.BoolInt32(!(v1278 != 0) && v1274 >= v1275)
	goto _1277
_1277:
	r89 = v1276
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	e89 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if r89 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14645, libc.VaList(bp+24, r89, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
	if e89&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:57:2:
		t_printf(tls, __ccgo_ts+14718, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1284 = *(*uint32)(unsafe.Pointer(bp))
	goto _1285
_1285:
	if libc.BoolInt32(v1284&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(1) + epsf
		v1283 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1) + epsf
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1286 = *(*uint32)(unsafe.Pointer(bp))
		goto _1287
	_1287:
		v1283 = libc.BoolInt32(v1286&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r90 = v1283
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e90 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r90 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+14797, libc.VaList(bp+24, r90, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e90&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+14872, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	v1288 = libc.Float64FromFloat32(1)
	v1289 = float64(libc.Float32FromFloat32(1) + epsf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1288
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1293 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1294
_1294:
	if libc.BoolInt32(v1293&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1289
		v1292 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1289
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1295 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1296
	_1296:
		v1292 = libc.BoolInt32(v1295&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1290 = libc.BoolInt32(!(v1292 != 0) && v1288 < v1289)
	goto _1291
_1291:
	r91 = v1290
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e91 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r91 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+14953, libc.VaList(bp+24, r91, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e91&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15023, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	v1297 = libc.Float64FromFloat32(1)
	v1298 = float64(libc.Float32FromFloat32(1) + epsf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1297
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1302 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1303
_1303:
	if libc.BoolInt32(v1302&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1298
		v1301 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1298
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1304 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1305
	_1305:
		v1301 = libc.BoolInt32(v1304&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1299 = libc.BoolInt32(!(v1301 != 0) && v1297 <= v1298)
	goto _1300
_1300:
	r92 = v1299
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e92 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r92 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15099, libc.VaList(bp+24, r92, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e92&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15174, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	v1306 = libc.Float64FromFloat32(1)
	v1307 = float64(libc.Float32FromFloat32(1) + epsf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1306
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1311 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1312
_1312:
	if libc.BoolInt32(v1311&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1307
		v1310 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1307
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1313 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1314
	_1314:
		v1310 = libc.BoolInt32(v1313&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1308 = libc.BoolInt32(!(v1310 != 0) && v1306 != v1307)
	goto _1309
_1309:
	r93 = v1308
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e93 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r93 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15255, libc.VaList(bp+24, r93, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e93&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15332, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	v1315 = libc.Float64FromFloat32(1)
	v1316 = float64(libc.Float32FromFloat32(1) + epsf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1315
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1320 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1321
_1321:
	if libc.BoolInt32(v1320&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1316
		v1319 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1316
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1322 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1323
	_1323:
		v1319 = libc.BoolInt32(v1322&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1317 = libc.BoolInt32(!(v1319 != 0) && v1315 > v1316)
	goto _1318
_1318:
	r94 = v1317
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e94 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r94 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15415, libc.VaList(bp+24, r94, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e94&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15488, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	v1324 = libc.Float64FromFloat32(1)
	v1325 = float64(libc.Float32FromFloat32(1) + epsf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1324
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1329 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1330
_1330:
	if libc.BoolInt32(v1329&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1325
		v1328 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1325
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1331 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1332
	_1332:
		v1328 = libc.BoolInt32(v1331&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1326 = libc.BoolInt32(!(v1328 != 0) && v1324 >= v1325)
	goto _1327
_1327:
	r95 = v1326
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	e95 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if r95 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15567, libc.VaList(bp+24, r95, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
	if e95&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:58:2:
		t_printf(tls, __ccgo_ts+15645, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1) + epsf
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1334 = *(*uint32)(unsafe.Pointer(bp))
	goto _1335
_1335:
	if libc.BoolInt32(v1334&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(1)
		v1333 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1336 = *(*uint32)(unsafe.Pointer(bp))
		goto _1337
	_1337:
		v1333 = libc.BoolInt32(v1336&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r96 = v1333
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e96 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r96 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+15729, libc.VaList(bp+24, r96, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e96&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+15804, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	v1338 = float64(libc.Float32FromFloat32(1) + epsf)
	v1339 = libc.Float64FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1338
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1343 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1344
_1344:
	if libc.BoolInt32(v1343&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1339
		v1342 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1339
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1345 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1346
	_1346:
		v1342 = libc.BoolInt32(v1345&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1340 = libc.BoolInt32(!(v1342 != 0) && v1338 < v1339)
	goto _1341
_1341:
	r97 = v1340
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e97 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r97 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+15885, libc.VaList(bp+24, r97, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e97&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+15955, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	v1347 = float64(libc.Float32FromFloat32(1) + epsf)
	v1348 = libc.Float64FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1347
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1352 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1353
_1353:
	if libc.BoolInt32(v1352&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1348
		v1351 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1348
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1354 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1355
	_1355:
		v1351 = libc.BoolInt32(v1354&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1349 = libc.BoolInt32(!(v1351 != 0) && v1347 <= v1348)
	goto _1350
_1350:
	r98 = v1349
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e98 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r98 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16031, libc.VaList(bp+24, r98, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e98&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16106, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	v1356 = float64(libc.Float32FromFloat32(1) + epsf)
	v1357 = libc.Float64FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1356
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1361 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1362
_1362:
	if libc.BoolInt32(v1361&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1357
		v1360 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1357
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1363 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1364
	_1364:
		v1360 = libc.BoolInt32(v1363&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1358 = libc.BoolInt32(!(v1360 != 0) && v1356 != v1357)
	goto _1359
_1359:
	r99 = v1358
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e99 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r99 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16187, libc.VaList(bp+24, r99, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e99&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16264, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	v1365 = float64(libc.Float32FromFloat32(1) + epsf)
	v1366 = libc.Float64FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1365
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1370 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1371
_1371:
	if libc.BoolInt32(v1370&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1366
		v1369 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1366
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1372 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1373
	_1373:
		v1369 = libc.BoolInt32(v1372&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1367 = libc.BoolInt32(!(v1369 != 0) && v1365 > v1366)
	goto _1368
_1368:
	r100 = v1367
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e100 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r100 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16347, libc.VaList(bp+24, r100, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e100&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16420, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	v1374 = float64(libc.Float32FromFloat32(1) + epsf)
	v1375 = libc.Float64FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1374
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1379 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1380
_1380:
	if libc.BoolInt32(v1379&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1375
		v1378 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1375
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1381 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1382
	_1382:
		v1378 = libc.BoolInt32(v1381&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1376 = libc.BoolInt32(!(v1378 != 0) && v1374 >= v1375)
	goto _1377
_1377:
	r101 = v1376
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	e101 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if r101 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16499, libc.VaList(bp+24, r101, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
	if e101&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:59:2:
		t_printf(tls, __ccgo_ts+16577, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = hugef - libc.Float32FromInt32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1384 = *(*uint32)(unsafe.Pointer(bp))
	goto _1385
_1385:
	if libc.BoolInt32(v1384&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = hugef
		v1383 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = hugef
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1386 = *(*uint32)(unsafe.Pointer(bp))
		goto _1387
	_1387:
		v1383 = libc.BoolInt32(v1386&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r102 = v1383
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e102 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r102 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+16661, libc.VaList(bp+24, r102, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e102&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+16735, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	v1388 = float64(hugef - libc.Float32FromInt32(1))
	v1389 = float64(hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1388
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1393 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1394
_1394:
	if libc.BoolInt32(v1393&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1389
		v1392 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1389
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1395 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1396
	_1396:
		v1392 = libc.BoolInt32(v1395&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1390 = libc.BoolInt32(!(v1392 != 0) && v1388 < v1389)
	goto _1391
_1391:
	r103 = v1390
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e103 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r103 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+16815, libc.VaList(bp+24, r103, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e103&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+16884, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	v1397 = float64(hugef - libc.Float32FromInt32(1))
	v1398 = float64(hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1397
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1402 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1403
_1403:
	if libc.BoolInt32(v1402&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1398
		v1401 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1398
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1404 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1405
	_1405:
		v1401 = libc.BoolInt32(v1404&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1399 = libc.BoolInt32(!(v1401 != 0) && v1397 <= v1398)
	goto _1400
_1400:
	r104 = v1399
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e104 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r104 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+16959, libc.VaList(bp+24, r104, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e104&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17033, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	v1406 = float64(hugef - libc.Float32FromInt32(1))
	v1407 = float64(hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1406
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1411 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1412
_1412:
	if libc.BoolInt32(v1411&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1407
		v1410 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1407
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1413 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1414
	_1414:
		v1410 = libc.BoolInt32(v1413&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1408 = libc.BoolInt32(!(v1410 != 0) && v1406 != v1407)
	goto _1409
_1409:
	r105 = v1408
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e105 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r105 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17113, libc.VaList(bp+24, r105, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e105&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17189, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	v1415 = float64(hugef - libc.Float32FromInt32(1))
	v1416 = float64(hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1415
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1420 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1421
_1421:
	if libc.BoolInt32(v1420&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1416
		v1419 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1416
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1422 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1423
	_1423:
		v1419 = libc.BoolInt32(v1422&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1417 = libc.BoolInt32(!(v1419 != 0) && v1415 > v1416)
	goto _1418
_1418:
	r106 = v1417
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e106 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r106 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17271, libc.VaList(bp+24, r106, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e106&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17343, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	v1424 = float64(hugef - libc.Float32FromInt32(1))
	v1425 = float64(hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1424
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1429 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1430
_1430:
	if libc.BoolInt32(v1429&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1425
		v1428 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1425
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1431 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1432
	_1432:
		v1428 = libc.BoolInt32(v1431&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1426 = libc.BoolInt32(!(v1428 != 0) && v1424 >= v1425)
	goto _1427
_1427:
	r107 = v1426
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	e107 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if r107 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17421, libc.VaList(bp+24, r107, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
	if e107&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:60:2:
		t_printf(tls, __ccgo_ts+17498, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = hugef
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1434 = *(*uint32)(unsafe.Pointer(bp))
	goto _1435
_1435:
	if libc.BoolInt32(v1434&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = hugef * hugef
		v1433 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = hugef * hugef
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1436 = *(*uint32)(unsafe.Pointer(bp))
		goto _1437
	_1437:
		v1433 = libc.BoolInt32(v1436&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r108 = v1433
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e108 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r108 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17581, libc.VaList(bp+24, r108, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e108&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17659, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	v1438 = float64(hugef)
	v1439 = float64(hugef * hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1438
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1443 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1444
_1444:
	if libc.BoolInt32(v1443&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1439
		v1442 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1439
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1445 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1446
	_1446:
		v1442 = libc.BoolInt32(v1445&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1440 = libc.BoolInt32(!(v1442 != 0) && v1438 < v1439)
	goto _1441
_1441:
	r109 = v1440
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e109 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r109 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17743, libc.VaList(bp+24, r109, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e109&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17816, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	v1447 = float64(hugef)
	v1448 = float64(hugef * hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1447
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1452 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1453
_1453:
	if libc.BoolInt32(v1452&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1448
		v1451 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1448
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1454 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1455
	_1455:
		v1451 = libc.BoolInt32(v1454&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1449 = libc.BoolInt32(!(v1451 != 0) && v1447 <= v1448)
	goto _1450
_1450:
	r110 = v1449
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e110 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r110 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17895, libc.VaList(bp+24, r110, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e110&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+17973, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	v1456 = float64(hugef)
	v1457 = float64(hugef * hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1456
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1461 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1462
_1462:
	if libc.BoolInt32(v1461&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1457
		v1460 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1457
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1463 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1464
	_1464:
		v1460 = libc.BoolInt32(v1463&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1458 = libc.BoolInt32(!(v1460 != 0) && v1456 != v1457)
	goto _1459
_1459:
	r111 = v1458
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e111 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r111 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18057, libc.VaList(bp+24, r111, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e111&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18137, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	v1465 = float64(hugef)
	v1466 = float64(hugef * hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1465
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1470 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1471
_1471:
	if libc.BoolInt32(v1470&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1466
		v1469 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1466
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1472 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1473
	_1473:
		v1469 = libc.BoolInt32(v1472&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1467 = libc.BoolInt32(!(v1469 != 0) && v1465 > v1466)
	goto _1468
_1468:
	r112 = v1467
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e112 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r112 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18223, libc.VaList(bp+24, r112, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e112&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18299, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	v1474 = float64(hugef)
	v1475 = float64(hugef * hugef)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1474
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1479 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1480
_1480:
	if libc.BoolInt32(v1479&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1475
		v1478 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1475
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1481 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1482
	_1482:
		v1478 = libc.BoolInt32(v1481&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1476 = libc.BoolInt32(!(v1478 != 0) && v1474 >= v1475)
	goto _1477
_1477:
	r113 = v1476
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	e113 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if r113 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18381, libc.VaList(bp+24, r113, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
	if e113&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:61:2:
		t_printf(tls, __ccgo_ts+18462, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = -libc.Float32FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1484 = *(*uint32)(unsafe.Pointer(bp))
	goto _1485
_1485:
	if libc.BoolInt32(v1484&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(0)
		v1483 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1486 = *(*uint32)(unsafe.Pointer(bp))
		goto _1487
	_1487:
		v1483 = libc.BoolInt32(v1486&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r114 = v1483
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e114 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r114 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18549, libc.VaList(bp+24, r114, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e114&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18620, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	v1488 = float64(-libc.Float32FromFloat32(0))
	v1489 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1488
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1493 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1494
_1494:
	if libc.BoolInt32(v1493&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1489
		v1492 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1489
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1495 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1496
	_1496:
		v1492 = libc.BoolInt32(v1495&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1490 = libc.BoolInt32(!(v1492 != 0) && v1488 < v1489)
	goto _1491
_1491:
	r115 = v1490
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e115 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r115 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18697, libc.VaList(bp+24, r115, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e115&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18763, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	v1497 = float64(-libc.Float32FromFloat32(0))
	v1498 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1497
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1502 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1503
_1503:
	if libc.BoolInt32(v1502&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1498
		v1501 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1498
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1504 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1505
	_1505:
		v1501 = libc.BoolInt32(v1504&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1499 = libc.BoolInt32(!(v1501 != 0) && v1497 <= v1498)
	goto _1500
_1500:
	r116 = v1499
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e116 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r116 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18835, libc.VaList(bp+24, r116, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e116&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18906, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	v1506 = float64(-libc.Float32FromFloat32(0))
	v1507 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1506
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1511 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1512
_1512:
	if libc.BoolInt32(v1511&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1507
		v1510 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1507
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1513 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1514
	_1514:
		v1510 = libc.BoolInt32(v1513&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1508 = libc.BoolInt32(!(v1510 != 0) && v1506 != v1507)
	goto _1509
_1509:
	r117 = v1508
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e117 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r117 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+18983, libc.VaList(bp+24, r117, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e117&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+19056, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	v1515 = float64(-libc.Float32FromFloat32(0))
	v1516 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1515
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1520 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1521
_1521:
	if libc.BoolInt32(v1520&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1516
		v1519 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1516
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1522 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1523
	_1523:
		v1519 = libc.BoolInt32(v1522&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1517 = libc.BoolInt32(!(v1519 != 0) && v1515 > v1516)
	goto _1518
_1518:
	r118 = v1517
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e118 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r118 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+19135, libc.VaList(bp+24, r118, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e118&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+19204, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	v1524 = float64(-libc.Float32FromFloat32(0))
	v1525 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1524
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1529 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1530
_1530:
	if libc.BoolInt32(v1529&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1525
		v1528 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1525
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1531 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1532
	_1532:
		v1528 = libc.BoolInt32(v1531&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1526 = libc.BoolInt32(!(v1528 != 0) && v1524 >= v1525)
	goto _1527
_1527:
	r119 = v1526
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	e119 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if r119 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+19279, libc.VaList(bp+24, r119, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
	if e119&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:62:2:
		t_printf(tls, __ccgo_ts+19353, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = -tinyf
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1534 = *(*uint32)(unsafe.Pointer(bp))
	goto _1535
_1535:
	if libc.BoolInt32(v1534&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(0)
		v1533 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1536 = *(*uint32)(unsafe.Pointer(bp))
		goto _1537
	_1537:
		v1533 = libc.BoolInt32(v1536&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r120 = v1533
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e120 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r120 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19433, libc.VaList(bp+24, r120, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e120&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19505, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	v1538 = float64(-tinyf)
	v1539 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1538
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1543 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1544
_1544:
	if libc.BoolInt32(v1543&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1539
		v1542 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1539
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1545 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1546
	_1546:
		v1542 = libc.BoolInt32(v1545&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1540 = libc.BoolInt32(!(v1542 != 0) && v1538 < v1539)
	goto _1541
_1541:
	r121 = v1540
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e121 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r121 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19583, libc.VaList(bp+24, r121, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e121&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19650, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	v1547 = float64(-tinyf)
	v1548 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1547
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1552 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1553
_1553:
	if libc.BoolInt32(v1552&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1548
		v1551 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1548
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1554 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1555
	_1555:
		v1551 = libc.BoolInt32(v1554&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1549 = libc.BoolInt32(!(v1551 != 0) && v1547 <= v1548)
	goto _1550
_1550:
	r122 = v1549
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e122 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r122 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19723, libc.VaList(bp+24, r122, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e122&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19795, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	v1556 = float64(-tinyf)
	v1557 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1556
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1561 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1562
_1562:
	if libc.BoolInt32(v1561&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1557
		v1560 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1557
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1563 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1564
	_1564:
		v1560 = libc.BoolInt32(v1563&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1558 = libc.BoolInt32(!(v1560 != 0) && v1556 != v1557)
	goto _1559
_1559:
	r123 = v1558
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e123 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r123 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19873, libc.VaList(bp+24, r123, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e123&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+19947, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	v1565 = float64(-tinyf)
	v1566 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1565
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1570 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1571
_1571:
	if libc.BoolInt32(v1570&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1566
		v1569 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1566
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1572 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1573
	_1573:
		v1569 = libc.BoolInt32(v1572&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1567 = libc.BoolInt32(!(v1569 != 0) && v1565 > v1566)
	goto _1568
_1568:
	r124 = v1567
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e124 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r124 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+20027, libc.VaList(bp+24, r124, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e124&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+20097, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	v1574 = float64(-tinyf)
	v1575 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1574
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1579 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1580
_1580:
	if libc.BoolInt32(v1579&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1575
		v1578 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1575
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1581 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1582
	_1582:
		v1578 = libc.BoolInt32(v1581&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1576 = libc.BoolInt32(!(v1578 != 0) && v1574 >= v1575)
	goto _1577
_1577:
	r125 = v1576
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	e125 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if r125 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+20173, libc.VaList(bp+24, r125, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
	if e125&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:63:2:
		t_printf(tls, __ccgo_ts+20248, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = tinyf
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1584 = *(*uint32)(unsafe.Pointer(bp))
	goto _1585
_1585:
	if libc.BoolInt32(v1584&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromInt32(2) * tinyf
		v1583 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromInt32(2) * tinyf
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1586 = *(*uint32)(unsafe.Pointer(bp))
		goto _1587
	_1587:
		v1583 = libc.BoolInt32(v1586&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r126 = v1583
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e126 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r126 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20329, libc.VaList(bp+24, r126, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e126&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20403, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	v1588 = float64(tinyf)
	v1589 = float64(libc.Float32FromInt32(2) * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1588
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1593 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1594
_1594:
	if libc.BoolInt32(v1593&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1589
		v1592 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1589
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1595 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1596
	_1596:
		v1592 = libc.BoolInt32(v1595&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1590 = libc.BoolInt32(!(v1592 != 0) && v1588 < v1589)
	goto _1591
_1591:
	r127 = v1590
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e127 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r127 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20483, libc.VaList(bp+24, r127, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e127&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20552, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	v1597 = float64(tinyf)
	v1598 = float64(libc.Float32FromInt32(2) * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1597
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1602 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1603
_1603:
	if libc.BoolInt32(v1602&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1598
		v1601 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1598
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1604 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1605
	_1605:
		v1601 = libc.BoolInt32(v1604&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1599 = libc.BoolInt32(!(v1601 != 0) && v1597 <= v1598)
	goto _1600
_1600:
	r128 = v1599
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e128 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r128 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20627, libc.VaList(bp+24, r128, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e128&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20701, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	v1606 = float64(tinyf)
	v1607 = float64(libc.Float32FromInt32(2) * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1606
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1611 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1612
_1612:
	if libc.BoolInt32(v1611&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1607
		v1610 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1607
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1613 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1614
	_1614:
		v1610 = libc.BoolInt32(v1613&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1608 = libc.BoolInt32(!(v1610 != 0) && v1606 != v1607)
	goto _1609
_1609:
	r129 = v1608
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e129 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r129 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20781, libc.VaList(bp+24, r129, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e129&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20857, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	v1615 = float64(tinyf)
	v1616 = float64(libc.Float32FromInt32(2) * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1615
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1620 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1621
_1621:
	if libc.BoolInt32(v1620&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1616
		v1619 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1616
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1622 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1623
	_1623:
		v1619 = libc.BoolInt32(v1622&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1617 = libc.BoolInt32(!(v1619 != 0) && v1615 > v1616)
	goto _1618
_1618:
	r130 = v1617
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e130 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r130 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+20939, libc.VaList(bp+24, r130, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e130&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+21011, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	v1624 = float64(tinyf)
	v1625 = float64(libc.Float32FromInt32(2) * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1624
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1629 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1630
_1630:
	if libc.BoolInt32(v1629&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1625
		v1628 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1625
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1631 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1632
	_1632:
		v1628 = libc.BoolInt32(v1631&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1626 = libc.BoolInt32(!(v1628 != 0) && v1624 >= v1625)
	goto _1627
_1627:
	r131 = v1626
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	e131 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if r131 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+21089, libc.VaList(bp+24, r131, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
	if e131&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:64:2:
		t_printf(tls, __ccgo_ts+21166, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = tinyf * libc.Float32FromFloat32(0.001953125)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1634 = *(*uint32)(unsafe.Pointer(bp))
	goto _1635
_1635:
	if libc.BoolInt32(v1634&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = tinyf * libc.Float32FromFloat32(0.00390625)
		v1633 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = tinyf * libc.Float32FromFloat32(0.00390625)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v1636 = *(*uint32)(unsafe.Pointer(bp))
		goto _1637
	_1637:
		v1633 = libc.BoolInt32(v1636&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r132 = v1633
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e132 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r132 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21249, libc.VaList(bp+24, r132, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e132&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21337, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	v1638 = float64(tinyf * libc.Float32FromFloat32(0.001953125))
	v1639 = float64(tinyf * libc.Float32FromFloat32(0.00390625))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1638
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1643 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1644
_1644:
	if libc.BoolInt32(v1643&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1639
		v1642 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1639
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1645 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1646
	_1646:
		v1642 = libc.BoolInt32(v1645&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1640 = libc.BoolInt32(!(v1642 != 0) && v1638 < v1639)
	goto _1641
_1641:
	r133 = v1640
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e133 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r133 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21431, libc.VaList(bp+24, r133, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e133&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21514, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	v1647 = float64(tinyf * libc.Float32FromFloat32(0.001953125))
	v1648 = float64(tinyf * libc.Float32FromFloat32(0.00390625))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1647
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1652 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1653
_1653:
	if libc.BoolInt32(v1652&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1648
		v1651 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1648
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1654 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1655
	_1655:
		v1651 = libc.BoolInt32(v1654&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1649 = libc.BoolInt32(!(v1651 != 0) && v1647 <= v1648)
	goto _1650
_1650:
	r134 = v1649
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e134 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r134 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21603, libc.VaList(bp+24, r134, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e134&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21691, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	v1656 = float64(tinyf * libc.Float32FromFloat32(0.001953125))
	v1657 = float64(tinyf * libc.Float32FromFloat32(0.00390625))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1656
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1661 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1662
_1662:
	if libc.BoolInt32(v1661&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1657
		v1660 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1657
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1663 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1664
	_1664:
		v1660 = libc.BoolInt32(v1663&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1658 = libc.BoolInt32(!(v1660 != 0) && v1656 != v1657)
	goto _1659
_1659:
	r135 = v1658
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e135 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r135 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21785, libc.VaList(bp+24, r135, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e135&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21875, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	v1665 = float64(tinyf * libc.Float32FromFloat32(0.001953125))
	v1666 = float64(tinyf * libc.Float32FromFloat32(0.00390625))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1665
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1670 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1671
_1671:
	if libc.BoolInt32(v1670&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1666
		v1669 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1666
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1672 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1673
	_1673:
		v1669 = libc.BoolInt32(v1672&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1667 = libc.BoolInt32(!(v1669 != 0) && v1665 > v1666)
	goto _1668
_1668:
	r136 = v1667
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e136 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r136 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+21971, libc.VaList(bp+24, r136, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e136&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+22057, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	v1674 = float64(tinyf * libc.Float32FromFloat32(0.001953125))
	v1675 = float64(tinyf * libc.Float32FromFloat32(0.00390625))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1674
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1679 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1680
_1680:
	if libc.BoolInt32(v1679&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1675
		v1678 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1675
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1681 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1682
	_1682:
		v1678 = libc.BoolInt32(v1681&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1676 = libc.BoolInt32(!(v1678 != 0) && v1674 >= v1675)
	goto _1677
_1677:
	r137 = v1676
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	e137 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if r137 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+22149, libc.VaList(bp+24, r137, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
	if e137&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:65:2:
		t_printf(tls, __ccgo_ts+22240, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1684 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1685
_1685:
	if libc.BoolInt32(v1684&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1.1)
		v1683 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1.1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1686 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1687
	_1687:
		v1683 = libc.BoolInt32(v1686&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r138 = v1683
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e138 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r138 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22337, libc.VaList(bp+24, r138, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e138&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22407, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	v1688 = libc.Float64FromFloat64(1)
	v1689 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1688
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1693 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1694
_1694:
	if libc.BoolInt32(v1693&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1689
		v1692 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1689
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1695 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1696
	_1696:
		v1692 = libc.BoolInt32(v1695&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1690 = libc.BoolInt32(!(v1692 != 0) && v1688 < v1689)
	goto _1691
_1691:
	r139 = v1690
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e139 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r139 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22483, libc.VaList(bp+24, r139, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e139&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22548, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	v1697 = libc.Float64FromFloat64(1)
	v1698 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1697
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1702 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1703
_1703:
	if libc.BoolInt32(v1702&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1698
		v1701 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1698
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1704 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1705
	_1705:
		v1701 = libc.BoolInt32(v1704&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1699 = libc.BoolInt32(!(v1701 != 0) && v1697 <= v1698)
	goto _1700
_1700:
	r140 = v1699
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e140 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r140 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22619, libc.VaList(bp+24, r140, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e140&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22689, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	v1706 = libc.Float64FromFloat64(1)
	v1707 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1706
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1711 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1712
_1712:
	if libc.BoolInt32(v1711&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1707
		v1710 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1707
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1713 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1714
	_1714:
		v1710 = libc.BoolInt32(v1713&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1708 = libc.BoolInt32(!(v1710 != 0) && v1706 != v1707)
	goto _1709
_1709:
	r141 = v1708
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e141 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r141 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22765, libc.VaList(bp+24, r141, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e141&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22837, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	v1715 = libc.Float64FromFloat64(1)
	v1716 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1715
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1720 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1721
_1721:
	if libc.BoolInt32(v1720&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1716
		v1719 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1716
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1722 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1723
	_1723:
		v1719 = libc.BoolInt32(v1722&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1717 = libc.BoolInt32(!(v1719 != 0) && v1715 > v1716)
	goto _1718
_1718:
	r142 = v1717
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e142 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r142 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22915, libc.VaList(bp+24, r142, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e142&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+22983, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	v1724 = libc.Float64FromFloat64(1)
	v1725 = libc.Float64FromFloat64(1.1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1724
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1729 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1730
_1730:
	if libc.BoolInt32(v1729&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1725
		v1728 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1725
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1731 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1732
	_1732:
		v1728 = libc.BoolInt32(v1731&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1726 = libc.BoolInt32(!(v1728 != 0) && v1724 >= v1725)
	goto _1727
_1727:
	r143 = v1726
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	e143 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if r143 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+23057, libc.VaList(bp+24, r143, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
	if e143&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:67:2:
		t_printf(tls, __ccgo_ts+23130, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1734 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1735
_1735:
	if libc.BoolInt32(v1734&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1) + epsl
		v1733 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1) + epsl
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1736 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1737
	_1737:
		v1733 = libc.BoolInt32(v1736&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r144 = v1733
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e144 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r144 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23209, libc.VaList(bp+24, r144, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e144&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23284, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	v1738 = libc.Float64FromFloat64(1)
	v1739 = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1738
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1743 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1744
_1744:
	if libc.BoolInt32(v1743&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1739
		v1742 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1739
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1745 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1746
	_1746:
		v1742 = libc.BoolInt32(v1745&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1740 = libc.BoolInt32(!(v1742 != 0) && v1738 < v1739)
	goto _1741
_1741:
	r145 = v1740
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e145 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r145 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23365, libc.VaList(bp+24, r145, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e145&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23435, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	v1747 = libc.Float64FromFloat64(1)
	v1748 = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1747
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1752 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1753
_1753:
	if libc.BoolInt32(v1752&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1748
		v1751 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1748
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1754 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1755
	_1755:
		v1751 = libc.BoolInt32(v1754&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1749 = libc.BoolInt32(!(v1751 != 0) && v1747 <= v1748)
	goto _1750
_1750:
	r146 = v1749
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e146 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r146 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23511, libc.VaList(bp+24, r146, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e146&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23586, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	v1756 = libc.Float64FromFloat64(1)
	v1757 = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1756
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1761 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1762
_1762:
	if libc.BoolInt32(v1761&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1757
		v1760 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1757
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1763 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1764
	_1764:
		v1760 = libc.BoolInt32(v1763&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1758 = libc.BoolInt32(!(v1760 != 0) && v1756 != v1757)
	goto _1759
_1759:
	r147 = v1758
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e147 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r147 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23667, libc.VaList(bp+24, r147, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e147&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23744, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	v1765 = libc.Float64FromFloat64(1)
	v1766 = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1765
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1770 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1771
_1771:
	if libc.BoolInt32(v1770&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1766
		v1769 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1766
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1772 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1773
	_1773:
		v1769 = libc.BoolInt32(v1772&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1767 = libc.BoolInt32(!(v1769 != 0) && v1765 > v1766)
	goto _1768
_1768:
	r148 = v1767
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e148 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r148 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23827, libc.VaList(bp+24, r148, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e148&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23900, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	v1774 = libc.Float64FromFloat64(1)
	v1775 = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1774
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1779 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1780
_1780:
	if libc.BoolInt32(v1779&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1775
		v1778 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1775
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1781 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1782
	_1782:
		v1778 = libc.BoolInt32(v1781&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1776 = libc.BoolInt32(!(v1778 != 0) && v1774 >= v1775)
	goto _1777
_1777:
	r149 = v1776
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	e149 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if r149 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+23979, libc.VaList(bp+24, r149, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
	if e149&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:68:2:
		t_printf(tls, __ccgo_ts+24057, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1) + epsl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1784 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1785
_1785:
	if libc.BoolInt32(v1784&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1)
		v1783 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1786 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1787
	_1787:
		v1783 = libc.BoolInt32(v1786&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r150 = v1783
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e150 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r150 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24141, libc.VaList(bp+24, r150, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e150&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24216, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	v1788 = libc.Float64FromFloat64(1) + epsl
	v1789 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1788
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1793 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1794
_1794:
	if libc.BoolInt32(v1793&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1789
		v1792 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1789
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1795 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1796
	_1796:
		v1792 = libc.BoolInt32(v1795&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1790 = libc.BoolInt32(!(v1792 != 0) && v1788 < v1789)
	goto _1791
_1791:
	r151 = v1790
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e151 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r151 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24297, libc.VaList(bp+24, r151, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e151&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24367, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	v1797 = libc.Float64FromFloat64(1) + epsl
	v1798 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1797
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1802 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1803
_1803:
	if libc.BoolInt32(v1802&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1798
		v1801 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1798
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1804 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1805
	_1805:
		v1801 = libc.BoolInt32(v1804&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1799 = libc.BoolInt32(!(v1801 != 0) && v1797 <= v1798)
	goto _1800
_1800:
	r152 = v1799
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e152 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r152 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24443, libc.VaList(bp+24, r152, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e152&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24518, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	v1806 = libc.Float64FromFloat64(1) + epsl
	v1807 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1806
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1811 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1812
_1812:
	if libc.BoolInt32(v1811&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1807
		v1810 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1807
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1813 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1814
	_1814:
		v1810 = libc.BoolInt32(v1813&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1808 = libc.BoolInt32(!(v1810 != 0) && v1806 != v1807)
	goto _1809
_1809:
	r153 = v1808
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e153 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r153 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24599, libc.VaList(bp+24, r153, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e153&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24676, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	v1815 = libc.Float64FromFloat64(1) + epsl
	v1816 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1815
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1820 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1821
_1821:
	if libc.BoolInt32(v1820&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1816
		v1819 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1816
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1822 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1823
	_1823:
		v1819 = libc.BoolInt32(v1822&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1817 = libc.BoolInt32(!(v1819 != 0) && v1815 > v1816)
	goto _1818
_1818:
	r154 = v1817
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e154 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r154 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24759, libc.VaList(bp+24, r154, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e154&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24832, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	v1824 = libc.Float64FromFloat64(1) + epsl
	v1825 = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1824
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1829 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1830
_1830:
	if libc.BoolInt32(v1829&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1825
		v1828 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1825
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1831 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1832
	_1832:
		v1828 = libc.BoolInt32(v1831&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1826 = libc.BoolInt32(!(v1828 != 0) && v1824 >= v1825)
	goto _1827
_1827:
	r155 = v1826
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	e155 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if r155 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24911, libc.VaList(bp+24, r155, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
	if e155&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:69:2:
		t_printf(tls, __ccgo_ts+24989, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = hugel - libc.Float64FromInt32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1834 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1835
_1835:
	if libc.BoolInt32(v1834&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = hugel
		v1833 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = hugel
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1836 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1837
	_1837:
		v1833 = libc.BoolInt32(v1836&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r156 = v1833
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e156 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r156 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25073, libc.VaList(bp+24, r156, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e156&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25147, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	v1838 = hugel - libc.Float64FromInt32(1)
	v1839 = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1838
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1843 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1844
_1844:
	if libc.BoolInt32(v1843&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1839
		v1842 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1839
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1845 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1846
	_1846:
		v1842 = libc.BoolInt32(v1845&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1840 = libc.BoolInt32(!(v1842 != 0) && v1838 < v1839)
	goto _1841
_1841:
	r157 = v1840
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e157 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r157 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25227, libc.VaList(bp+24, r157, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e157&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25296, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	v1847 = hugel - libc.Float64FromInt32(1)
	v1848 = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1847
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1852 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1853
_1853:
	if libc.BoolInt32(v1852&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1848
		v1851 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1848
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1854 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1855
	_1855:
		v1851 = libc.BoolInt32(v1854&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1849 = libc.BoolInt32(!(v1851 != 0) && v1847 <= v1848)
	goto _1850
_1850:
	r158 = v1849
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e158 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r158 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25371, libc.VaList(bp+24, r158, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e158&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25445, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	v1856 = hugel - libc.Float64FromInt32(1)
	v1857 = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1856
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1861 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1862
_1862:
	if libc.BoolInt32(v1861&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1857
		v1860 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1857
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1863 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1864
	_1864:
		v1860 = libc.BoolInt32(v1863&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1858 = libc.BoolInt32(!(v1860 != 0) && v1856 != v1857)
	goto _1859
_1859:
	r159 = v1858
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e159 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r159 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25525, libc.VaList(bp+24, r159, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e159&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25601, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	v1865 = hugel - libc.Float64FromInt32(1)
	v1866 = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1865
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1870 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1871
_1871:
	if libc.BoolInt32(v1870&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1866
		v1869 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1866
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1872 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1873
	_1873:
		v1869 = libc.BoolInt32(v1872&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1867 = libc.BoolInt32(!(v1869 != 0) && v1865 > v1866)
	goto _1868
_1868:
	r160 = v1867
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e160 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r160 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25683, libc.VaList(bp+24, r160, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e160&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25755, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	v1874 = hugel - libc.Float64FromInt32(1)
	v1875 = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1874
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1879 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1880
_1880:
	if libc.BoolInt32(v1879&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1875
		v1878 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1875
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1881 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1882
	_1882:
		v1878 = libc.BoolInt32(v1881&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1876 = libc.BoolInt32(!(v1878 != 0) && v1874 >= v1875)
	goto _1877
_1877:
	r161 = v1876
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	e161 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if r161 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25833, libc.VaList(bp+24, r161, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
	if e161&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:70:2:
		t_printf(tls, __ccgo_ts+25910, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1884 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1885
_1885:
	if libc.BoolInt32(v1884&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = hugel * hugel
		v1883 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = hugel * hugel
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1886 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1887
	_1887:
		v1883 = libc.BoolInt32(v1886&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r162 = v1883
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e162 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r162 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+25993, libc.VaList(bp+24, r162, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e162&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26071, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	v1888 = hugel
	v1889 = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1888
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1893 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1894
_1894:
	if libc.BoolInt32(v1893&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1889
		v1892 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1889
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1895 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1896
	_1896:
		v1892 = libc.BoolInt32(v1895&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1890 = libc.BoolInt32(!(v1892 != 0) && v1888 < v1889)
	goto _1891
_1891:
	r163 = v1890
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e163 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r163 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26155, libc.VaList(bp+24, r163, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e163&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26228, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	v1897 = hugel
	v1898 = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1897
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1902 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1903
_1903:
	if libc.BoolInt32(v1902&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1898
		v1901 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1898
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1904 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1905
	_1905:
		v1901 = libc.BoolInt32(v1904&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1899 = libc.BoolInt32(!(v1901 != 0) && v1897 <= v1898)
	goto _1900
_1900:
	r164 = v1899
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e164 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r164 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26307, libc.VaList(bp+24, r164, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e164&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26385, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	v1906 = hugel
	v1907 = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1906
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1911 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1912
_1912:
	if libc.BoolInt32(v1911&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1907
		v1910 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1907
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1913 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1914
	_1914:
		v1910 = libc.BoolInt32(v1913&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1908 = libc.BoolInt32(!(v1910 != 0) && v1906 != v1907)
	goto _1909
_1909:
	r165 = v1908
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e165 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r165 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26469, libc.VaList(bp+24, r165, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e165&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26549, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	v1915 = hugel
	v1916 = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1915
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1920 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1921
_1921:
	if libc.BoolInt32(v1920&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1916
		v1919 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1916
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1922 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1923
	_1923:
		v1919 = libc.BoolInt32(v1922&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1917 = libc.BoolInt32(!(v1919 != 0) && v1915 > v1916)
	goto _1918
_1918:
	r166 = v1917
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e166 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r166 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26635, libc.VaList(bp+24, r166, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e166&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26711, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	v1924 = hugel
	v1925 = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1924
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1929 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1930
_1930:
	if libc.BoolInt32(v1929&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1925
		v1928 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1925
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1931 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1932
	_1932:
		v1928 = libc.BoolInt32(v1931&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1926 = libc.BoolInt32(!(v1928 != 0) && v1924 >= v1925)
	goto _1927
_1927:
	r167 = v1926
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	e167 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if r167 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26793, libc.VaList(bp+24, r167, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
	if e167&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:71:2:
		t_printf(tls, __ccgo_ts+26874, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1934 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1935
_1935:
	if libc.BoolInt32(v1934&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v1933 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1936 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1937
	_1937:
		v1933 = libc.BoolInt32(v1936&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r168 = v1933
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e168 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r168 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+26961, libc.VaList(bp+24, r168, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e168&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27032, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	v1938 = -libc.Float64FromFloat64(0)
	v1939 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1938
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1943 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1944
_1944:
	if libc.BoolInt32(v1943&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1939
		v1942 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1939
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1945 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1946
	_1946:
		v1942 = libc.BoolInt32(v1945&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1940 = libc.BoolInt32(!(v1942 != 0) && v1938 < v1939)
	goto _1941
_1941:
	r169 = v1940
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e169 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r169 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27109, libc.VaList(bp+24, r169, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e169&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27175, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	v1947 = -libc.Float64FromFloat64(0)
	v1948 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1947
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1952 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1953
_1953:
	if libc.BoolInt32(v1952&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1948
		v1951 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1948
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1954 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1955
	_1955:
		v1951 = libc.BoolInt32(v1954&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1949 = libc.BoolInt32(!(v1951 != 0) && v1947 <= v1948)
	goto _1950
_1950:
	r170 = v1949
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e170 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r170 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27247, libc.VaList(bp+24, r170, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e170&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27318, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	v1956 = -libc.Float64FromFloat64(0)
	v1957 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1956
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1961 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1962
_1962:
	if libc.BoolInt32(v1961&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1957
		v1960 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1957
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1963 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1964
	_1964:
		v1960 = libc.BoolInt32(v1963&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1958 = libc.BoolInt32(!(v1960 != 0) && v1956 != v1957)
	goto _1959
_1959:
	r171 = v1958
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e171 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r171 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27395, libc.VaList(bp+24, r171, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e171&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27468, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	v1965 = -libc.Float64FromFloat64(0)
	v1966 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1965
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1970 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1971
_1971:
	if libc.BoolInt32(v1970&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1966
		v1969 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1966
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1972 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1973
	_1973:
		v1969 = libc.BoolInt32(v1972&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1967 = libc.BoolInt32(!(v1969 != 0) && v1965 > v1966)
	goto _1968
_1968:
	r172 = v1967
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e172 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r172 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27547, libc.VaList(bp+24, r172, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e172&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27616, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	v1974 = -libc.Float64FromFloat64(0)
	v1975 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1974
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1979 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1980
_1980:
	if libc.BoolInt32(v1979&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1975
		v1978 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1975
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1981 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1982
	_1982:
		v1978 = libc.BoolInt32(v1981&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1976 = libc.BoolInt32(!(v1978 != 0) && v1974 >= v1975)
	goto _1977
_1977:
	r173 = v1976
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	e173 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if r173 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27691, libc.VaList(bp+24, r173, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
	if e173&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:72:2:
		t_printf(tls, __ccgo_ts+27765, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1984 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1985
_1985:
	if libc.BoolInt32(v1984&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v1983 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1986 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1987
	_1987:
		v1983 = libc.BoolInt32(v1986&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r174 = v1983
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e174 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r174 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+27845, libc.VaList(bp+24, r174, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e174&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+27917, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	v1988 = -tinyl
	v1989 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1988
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1993 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _1994
_1994:
	if libc.BoolInt32(v1993&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1989
		v1992 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1989
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v1995 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _1996
	_1996:
		v1992 = libc.BoolInt32(v1995&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1990 = libc.BoolInt32(!(v1992 != 0) && v1988 < v1989)
	goto _1991
_1991:
	r175 = v1990
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e175 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r175 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+27995, libc.VaList(bp+24, r175, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e175&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28062, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	v1997 = -tinyl
	v1998 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v1997
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2002 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2003
_2003:
	if libc.BoolInt32(v2002&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v1998
		v2001 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v1998
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2004 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2005
	_2005:
		v2001 = libc.BoolInt32(v2004&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v1999 = libc.BoolInt32(!(v2001 != 0) && v1997 <= v1998)
	goto _2000
_2000:
	r176 = v1999
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e176 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r176 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28135, libc.VaList(bp+24, r176, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e176&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28207, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	v2006 = -tinyl
	v2007 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2006
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2011 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2012
_2012:
	if libc.BoolInt32(v2011&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2007
		v2010 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2007
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2013 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2014
	_2014:
		v2010 = libc.BoolInt32(v2013&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2008 = libc.BoolInt32(!(v2010 != 0) && v2006 != v2007)
	goto _2009
_2009:
	r177 = v2008
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e177 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r177 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28285, libc.VaList(bp+24, r177, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e177&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28359, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	v2015 = -tinyl
	v2016 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2015
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2020 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2021
_2021:
	if libc.BoolInt32(v2020&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2016
		v2019 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2016
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2022 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2023
	_2023:
		v2019 = libc.BoolInt32(v2022&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2017 = libc.BoolInt32(!(v2019 != 0) && v2015 > v2016)
	goto _2018
_2018:
	r178 = v2017
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e178 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r178 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28439, libc.VaList(bp+24, r178, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e178&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28509, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	v2024 = -tinyl
	v2025 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2024
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2029 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2030
_2030:
	if libc.BoolInt32(v2029&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2025
		v2028 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2025
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2031 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2032
	_2032:
		v2028 = libc.BoolInt32(v2031&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2026 = libc.BoolInt32(!(v2028 != 0) && v2024 >= v2025)
	goto _2027
_2027:
	r179 = v2026
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	e179 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if r179 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28585, libc.VaList(bp+24, r179, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
	if e179&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:73:2:
		t_printf(tls, __ccgo_ts+28660, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2034 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2035
_2035:
	if libc.BoolInt32(v2034&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromInt32(2) * tinyl
		v2033 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromInt32(2) * tinyl
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2036 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2037
	_2037:
		v2033 = libc.BoolInt32(v2036&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r180 = v2033
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e180 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r180 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+28741, libc.VaList(bp+24, r180, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e180&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+28815, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	v2038 = tinyl
	v2039 = libc.Float64FromInt32(2) * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2038
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2043 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2044
_2044:
	if libc.BoolInt32(v2043&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2039
		v2042 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2039
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2045 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2046
	_2046:
		v2042 = libc.BoolInt32(v2045&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2040 = libc.BoolInt32(!(v2042 != 0) && v2038 < v2039)
	goto _2041
_2041:
	r181 = v2040
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e181 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r181 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+28895, libc.VaList(bp+24, r181, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e181&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+28964, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	v2047 = tinyl
	v2048 = libc.Float64FromInt32(2) * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2047
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2052 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2053
_2053:
	if libc.BoolInt32(v2052&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2048
		v2051 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2048
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2054 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2055
	_2055:
		v2051 = libc.BoolInt32(v2054&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2049 = libc.BoolInt32(!(v2051 != 0) && v2047 <= v2048)
	goto _2050
_2050:
	r182 = v2049
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e182 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r182 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29039, libc.VaList(bp+24, r182, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e182&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29113, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	v2056 = tinyl
	v2057 = libc.Float64FromInt32(2) * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2056
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2061 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2062
_2062:
	if libc.BoolInt32(v2061&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2057
		v2060 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2057
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2063 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2064
	_2064:
		v2060 = libc.BoolInt32(v2063&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2058 = libc.BoolInt32(!(v2060 != 0) && v2056 != v2057)
	goto _2059
_2059:
	r183 = v2058
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e183 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r183 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29193, libc.VaList(bp+24, r183, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e183&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29269, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	v2065 = tinyl
	v2066 = libc.Float64FromInt32(2) * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2065
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2070 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2071
_2071:
	if libc.BoolInt32(v2070&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2066
		v2069 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2066
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2072 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2073
	_2073:
		v2069 = libc.BoolInt32(v2072&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2067 = libc.BoolInt32(!(v2069 != 0) && v2065 > v2066)
	goto _2068
_2068:
	r184 = v2067
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e184 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r184 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29351, libc.VaList(bp+24, r184, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e184&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29423, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	v2074 = tinyl
	v2075 = libc.Float64FromInt32(2) * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2074
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2079 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2080
_2080:
	if libc.BoolInt32(v2079&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2075
		v2078 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2075
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2081 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2082
	_2082:
		v2078 = libc.BoolInt32(v2081&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2076 = libc.BoolInt32(!(v2078 != 0) && v2074 >= v2075)
	goto _2077
_2077:
	r185 = v2076
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	e185 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if r185 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29501, libc.VaList(bp+24, r185, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
	if e185&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:74:2:
		t_printf(tls, __ccgo_ts+29578, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tinyl * libc.Float64FromFloat64(0.001953125)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2084 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2085
_2085:
	if libc.BoolInt32(v2084&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = tinyl * libc.Float64FromFloat64(0.00390625)
		v2083 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = tinyl * libc.Float64FromFloat64(0.00390625)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2086 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2087
	_2087:
		v2083 = libc.BoolInt32(v2086&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r186 = v2083
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e186 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r186 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+29661, libc.VaList(bp+24, r186, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e186&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+29749, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	v2088 = tinyl * libc.Float64FromFloat64(0.001953125)
	v2089 = tinyl * libc.Float64FromFloat64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2088
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2093 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2094
_2094:
	if libc.BoolInt32(v2093&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2089
		v2092 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2089
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2095 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2096
	_2096:
		v2092 = libc.BoolInt32(v2095&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2090 = libc.BoolInt32(!(v2092 != 0) && v2088 < v2089)
	goto _2091
_2091:
	r187 = v2090
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e187 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r187 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+29843, libc.VaList(bp+24, r187, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e187&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+29926, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	v2097 = tinyl * libc.Float64FromFloat64(0.001953125)
	v2098 = tinyl * libc.Float64FromFloat64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2097
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2102 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2103
_2103:
	if libc.BoolInt32(v2102&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2098
		v2101 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2098
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2104 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2105
	_2105:
		v2101 = libc.BoolInt32(v2104&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2099 = libc.BoolInt32(!(v2101 != 0) && v2097 <= v2098)
	goto _2100
_2100:
	r188 = v2099
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e188 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r188 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30015, libc.VaList(bp+24, r188, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e188&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30103, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	v2106 = tinyl * libc.Float64FromFloat64(0.001953125)
	v2107 = tinyl * libc.Float64FromFloat64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2106
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2111 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2112
_2112:
	if libc.BoolInt32(v2111&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2107
		v2110 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2107
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2113 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2114
	_2114:
		v2110 = libc.BoolInt32(v2113&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2108 = libc.BoolInt32(!(v2110 != 0) && v2106 != v2107)
	goto _2109
_2109:
	r189 = v2108
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e189 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r189 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30197, libc.VaList(bp+24, r189, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e189&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30287, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	v2115 = tinyl * libc.Float64FromFloat64(0.001953125)
	v2116 = tinyl * libc.Float64FromFloat64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2115
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2120 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2121
_2121:
	if libc.BoolInt32(v2120&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2116
		v2119 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2116
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2122 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2123
	_2123:
		v2119 = libc.BoolInt32(v2122&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2117 = libc.BoolInt32(!(v2119 != 0) && v2115 > v2116)
	goto _2118
_2118:
	r190 = v2117
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e190 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r190 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30383, libc.VaList(bp+24, r190, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e190&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30469, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	v2124 = tinyl * libc.Float64FromFloat64(0.001953125)
	v2125 = tinyl * libc.Float64FromFloat64(0.00390625)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2124
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2129 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2130
_2130:
	if libc.BoolInt32(v2129&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2125
		v2128 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2125
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2131 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2132
	_2132:
		v2128 = libc.BoolInt32(v2131&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2126 = libc.BoolInt32(!(v2128 != 0) && v2124 >= v2125)
	goto _2127
_2127:
	r191 = v2126
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	e191 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if r191 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30561, libc.VaList(bp+24, r191, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
	if e191&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:75:2:
		t_printf(tls, __ccgo_ts+30652, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = huge * huge
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2134 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2135
_2135:
	if libc.BoolInt32(v2134&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = huge * huge * libc.Float64FromInt32(2)
		v2133 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = huge * huge * libc.Float64FromInt32(2)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2136 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2137
	_2137:
		v2133 = libc.BoolInt32(v2136&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r192 = v2133
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e192 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r192 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+30749, libc.VaList(bp+24, r192, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e192&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+30831, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	v2138 = huge * huge
	v2139 = huge * huge * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2138
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2143 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2144
_2144:
	if libc.BoolInt32(v2143&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2139
		v2142 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2139
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2145 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2146
	_2146:
		v2142 = libc.BoolInt32(v2145&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2140 = libc.BoolInt32(!(v2142 != 0) && v2138 < v2139)
	goto _2141
_2141:
	r193 = v2140
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e193 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r193 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+30919, libc.VaList(bp+24, r193, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e193&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+30996, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	v2147 = huge * huge
	v2148 = huge * huge * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2147
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2152 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2153
_2153:
	if libc.BoolInt32(v2152&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2148
		v2151 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2148
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2154 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2155
	_2155:
		v2151 = libc.BoolInt32(v2154&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2149 = libc.BoolInt32(!(v2151 != 0) && v2147 <= v2148)
	goto _2150
_2150:
	r194 = v2149
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e194 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r194 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31079, libc.VaList(bp+24, r194, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e194&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31161, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	v2156 = huge * huge
	v2157 = huge * huge * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2156
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2161 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2162
_2162:
	if libc.BoolInt32(v2161&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2157
		v2160 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2157
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2163 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2164
	_2164:
		v2160 = libc.BoolInt32(v2163&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2158 = libc.BoolInt32(!(v2160 != 0) && v2156 != v2157)
	goto _2159
_2159:
	r195 = v2158
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e195 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r195 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31249, libc.VaList(bp+24, r195, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e195&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31333, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	v2165 = huge * huge
	v2166 = huge * huge * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2165
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2170 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2171
_2171:
	if libc.BoolInt32(v2170&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2166
		v2169 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2166
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2172 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2173
	_2173:
		v2169 = libc.BoolInt32(v2172&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2167 = libc.BoolInt32(!(v2169 != 0) && v2165 > v2166)
	goto _2168
_2168:
	r196 = v2167
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e196 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r196 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31423, libc.VaList(bp+24, r196, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e196&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31503, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	v2174 = huge * huge
	v2175 = huge * huge * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2174
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2179 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2180
_2180:
	if libc.BoolInt32(v2179&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2175
		v2178 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2175
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2181 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2182
	_2182:
		v2178 = libc.BoolInt32(v2181&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2176 = libc.BoolInt32(!(v2178 != 0) && v2174 >= v2175)
	goto _2177
_2177:
	r197 = v2176
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	e197 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if r197 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31589, libc.VaList(bp+24, r197, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
	if e197&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:78:2:
		t_printf(tls, __ccgo_ts+31674, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tiny * tiny * libc.Float64FromFloat64(0.5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2184 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2185
_2185:
	if libc.BoolInt32(v2184&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = tiny * tiny
		v2183 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = tiny * tiny
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2186 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2187
	_2187:
		v2183 = libc.BoolInt32(v2186&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r198 = v2183
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e198 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r198 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+31765, libc.VaList(bp+24, r198, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e198&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+31849, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	v2188 = tiny * tiny * float64(0.5)
	v2189 = tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2188
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2193 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2194
_2194:
	if libc.BoolInt32(v2193&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2189
		v2192 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2189
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2195 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2196
	_2196:
		v2192 = libc.BoolInt32(v2195&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2190 = libc.BoolInt32(!(v2192 != 0) && v2188 < v2189)
	goto _2191
_2191:
	r199 = v2190
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e199 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r199 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+31939, libc.VaList(bp+24, r199, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e199&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32018, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	v2197 = tiny * tiny * float64(0.5)
	v2198 = tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2197
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2202 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2203
_2203:
	if libc.BoolInt32(v2202&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2198
		v2201 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2198
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2204 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2205
	_2205:
		v2201 = libc.BoolInt32(v2204&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2199 = libc.BoolInt32(!(v2201 != 0) && v2197 <= v2198)
	goto _2200
_2200:
	r200 = v2199
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e200 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r200 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32103, libc.VaList(bp+24, r200, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e200&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32187, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	v2206 = tiny * tiny * float64(0.5)
	v2207 = tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2206
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2211 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2212
_2212:
	if libc.BoolInt32(v2211&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2207
		v2210 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2207
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2213 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2214
	_2214:
		v2210 = libc.BoolInt32(v2213&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2208 = libc.BoolInt32(!(v2210 != 0) && v2206 != v2207)
	goto _2209
_2209:
	r201 = v2208
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e201 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r201 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32277, libc.VaList(bp+24, r201, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e201&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32363, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	v2215 = tiny * tiny * float64(0.5)
	v2216 = tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2215
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2220 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2221
_2221:
	if libc.BoolInt32(v2220&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2216
		v2219 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2216
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2222 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2223
	_2223:
		v2219 = libc.BoolInt32(v2222&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2217 = libc.BoolInt32(!(v2219 != 0) && v2215 > v2216)
	goto _2218
_2218:
	r202 = v2217
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e202 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r202 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32455, libc.VaList(bp+24, r202, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e202&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32537, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	v2224 = tiny * tiny * float64(0.5)
	v2225 = tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2224
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2229 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2230
_2230:
	if libc.BoolInt32(v2229&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2225
		v2228 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2225
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2231 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2232
	_2232:
		v2228 = libc.BoolInt32(v2231&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2226 = libc.BoolInt32(!(v2228 != 0) && v2224 >= v2225)
	goto _2227
_2227:
	r203 = v2226
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	e203 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if r203 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32625, libc.VaList(bp+24, r203, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
	if e203&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:79:2:
		t_printf(tls, __ccgo_ts+32712, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -tiny * tiny
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2234 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2235
_2235:
	if libc.BoolInt32(v2234&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v2233 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2236 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2237
	_2237:
		v2233 = libc.BoolInt32(v2236&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r204 = v2233
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e204 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r204 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+32805, libc.VaList(bp+24, r204, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e204&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+32880, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	v2238 = -tiny * tiny
	v2239 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2238
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2243 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2244
_2244:
	if libc.BoolInt32(v2243&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2239
		v2242 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2239
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2245 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2246
	_2246:
		v2242 = libc.BoolInt32(v2245&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2240 = libc.BoolInt32(!(v2242 != 0) && v2238 < v2239)
	goto _2241
_2241:
	r205 = v2240
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e205 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r205 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+32961, libc.VaList(bp+24, r205, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e205&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33031, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	v2247 = -tiny * tiny
	v2248 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2247
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2252 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2253
_2253:
	if libc.BoolInt32(v2252&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2248
		v2251 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2248
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2254 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2255
	_2255:
		v2251 = libc.BoolInt32(v2254&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2249 = libc.BoolInt32(!(v2251 != 0) && v2247 <= v2248)
	goto _2250
_2250:
	r206 = v2249
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e206 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r206 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33107, libc.VaList(bp+24, r206, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e206&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33182, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	v2256 = -tiny * tiny
	v2257 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2256
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2261 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2262
_2262:
	if libc.BoolInt32(v2261&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2257
		v2260 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2257
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2263 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2264
	_2264:
		v2260 = libc.BoolInt32(v2263&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2258 = libc.BoolInt32(!(v2260 != 0) && v2256 != v2257)
	goto _2259
_2259:
	r207 = v2258
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e207 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r207 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33263, libc.VaList(bp+24, r207, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e207&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33340, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	v2265 = -tiny * tiny
	v2266 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2265
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2270 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2271
_2271:
	if libc.BoolInt32(v2270&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2266
		v2269 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2266
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2272 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2273
	_2273:
		v2269 = libc.BoolInt32(v2272&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2267 = libc.BoolInt32(!(v2269 != 0) && v2265 > v2266)
	goto _2268
_2268:
	r208 = v2267
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e208 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r208 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33423, libc.VaList(bp+24, r208, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e208&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33496, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	v2274 = -tiny * tiny
	v2275 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2274
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2279 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2280
_2280:
	if libc.BoolInt32(v2279&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2275
		v2278 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2275
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2281 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2282
	_2282:
		v2278 = libc.BoolInt32(v2281&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2276 = libc.BoolInt32(!(v2278 != 0) && v2274 >= v2275)
	goto _2277
_2277:
	r209 = v2276
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	e209 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if r209 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33575, libc.VaList(bp+24, r209, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
	if e209&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:80:2:
		t_printf(tls, __ccgo_ts+33653, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = float64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2284 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2285
_2285:
	if libc.BoolInt32(v2284&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1) + eps*libc.Float64FromFloat64(0.25)
		v2283 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = float64(1) + eps*float64(0.25)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2286 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2287
	_2287:
		v2283 = libc.BoolInt32(v2286&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r210 = v2283
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e210 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r210 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+33737, libc.VaList(bp+24, r210, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e210&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+33814, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	v2288 = libc.Float64FromFloat64(1)
	v2289 = float64(1) + eps*float64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2288
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2293 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2294
_2294:
	if libc.BoolInt32(v2293&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2289
		v2292 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2289
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2295 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2296
	_2296:
		v2292 = libc.BoolInt32(v2295&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2290 = libc.BoolInt32(!(v2292 != 0) && v2288 < v2289)
	goto _2291
_2291:
	r211 = v2290
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e211 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r211 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+33897, libc.VaList(bp+24, r211, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e211&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+33969, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	v2297 = libc.Float64FromFloat64(1)
	v2298 = float64(1) + eps*float64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2297
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2302 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2303
_2303:
	if libc.BoolInt32(v2302&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2298
		v2301 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2298
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2304 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2305
	_2305:
		v2301 = libc.BoolInt32(v2304&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2299 = libc.BoolInt32(!(v2301 != 0) && v2297 <= v2298)
	goto _2300
_2300:
	r212 = v2299
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e212 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r212 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34047, libc.VaList(bp+24, r212, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e212&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34124, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	v2306 = libc.Float64FromFloat64(1)
	v2307 = float64(1) + eps*float64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2306
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2311 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2312
_2312:
	if libc.BoolInt32(v2311&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2307
		v2310 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2307
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2313 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2314
	_2314:
		v2310 = libc.BoolInt32(v2313&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2308 = libc.BoolInt32(!(v2310 != 0) && v2306 != v2307)
	goto _2309
_2309:
	r213 = v2308
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e213 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r213 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34207, libc.VaList(bp+24, r213, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e213&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34286, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	v2315 = libc.Float64FromFloat64(1)
	v2316 = float64(1) + eps*float64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2315
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2320 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2321
_2321:
	if libc.BoolInt32(v2320&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2316
		v2319 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2316
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2322 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2323
	_2323:
		v2319 = libc.BoolInt32(v2322&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2317 = libc.BoolInt32(!(v2319 != 0) && v2315 > v2316)
	goto _2318
_2318:
	r214 = v2317
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e214 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r214 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34371, libc.VaList(bp+24, r214, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e214&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34446, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	v2324 = libc.Float64FromFloat64(1)
	v2325 = float64(1) + eps*float64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2324
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2329 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2330
_2330:
	if libc.BoolInt32(v2329&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2325
		v2328 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2325
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2331 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2332
	_2332:
		v2328 = libc.BoolInt32(v2331&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2326 = libc.BoolInt32(!(v2328 != 0) && v2324 >= v2325)
	goto _2327
_2327:
	r215 = v2326
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	e215 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if r215 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34527, libc.VaList(bp+24, r215, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
	if e215&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:81:2:
		t_printf(tls, __ccgo_ts+34607, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = hugef * hugef
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v2334 = *(*uint32)(unsafe.Pointer(bp))
	goto _2335
_2335:
	if libc.BoolInt32(v2334&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = hugef * hugef * libc.Float32FromInt32(2)
		v2333 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = hugef * hugef * libc.Float32FromInt32(2)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v2336 = *(*uint32)(unsafe.Pointer(bp))
		goto _2337
	_2337:
		v2333 = libc.BoolInt32(v2336&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r216 = v2333
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e216 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r216 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+34693, libc.VaList(bp+24, r216, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e216&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+34779, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	v2338 = float64(hugef * hugef)
	v2339 = float64(hugef * hugef * libc.Float32FromInt32(2))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2338
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2343 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2344
_2344:
	if libc.BoolInt32(v2343&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2339
		v2342 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2339
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2345 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2346
	_2346:
		v2342 = libc.BoolInt32(v2345&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2340 = libc.BoolInt32(!(v2342 != 0) && v2338 < v2339)
	goto _2341
_2341:
	r217 = v2340
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e217 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r217 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+34871, libc.VaList(bp+24, r217, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e217&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+34952, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	v2347 = float64(hugef * hugef)
	v2348 = float64(hugef * hugef * libc.Float32FromInt32(2))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2347
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2352 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2353
_2353:
	if libc.BoolInt32(v2352&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2348
		v2351 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2348
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2354 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2355
	_2355:
		v2351 = libc.BoolInt32(v2354&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2349 = libc.BoolInt32(!(v2351 != 0) && v2347 <= v2348)
	goto _2350
_2350:
	r218 = v2349
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e218 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r218 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35039, libc.VaList(bp+24, r218, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e218&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35125, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	v2356 = float64(hugef * hugef)
	v2357 = float64(hugef * hugef * libc.Float32FromInt32(2))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2356
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2361 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2362
_2362:
	if libc.BoolInt32(v2361&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2357
		v2360 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2357
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2363 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2364
	_2364:
		v2360 = libc.BoolInt32(v2363&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2358 = libc.BoolInt32(!(v2360 != 0) && v2356 != v2357)
	goto _2359
_2359:
	r219 = v2358
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e219 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r219 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35217, libc.VaList(bp+24, r219, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e219&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35305, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	v2365 = float64(hugef * hugef)
	v2366 = float64(hugef * hugef * libc.Float32FromInt32(2))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2365
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2370 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2371
_2371:
	if libc.BoolInt32(v2370&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2366
		v2369 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2366
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2372 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2373
	_2373:
		v2369 = libc.BoolInt32(v2372&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2367 = libc.BoolInt32(!(v2369 != 0) && v2365 > v2366)
	goto _2368
_2368:
	r220 = v2367
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e220 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r220 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35399, libc.VaList(bp+24, r220, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e220&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35483, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	v2374 = float64(hugef * hugef)
	v2375 = float64(hugef * hugef * libc.Float32FromInt32(2))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2374
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2379 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2380
_2380:
	if libc.BoolInt32(v2379&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2375
		v2378 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2375
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2381 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2382
	_2382:
		v2378 = libc.BoolInt32(v2381&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2376 = libc.BoolInt32(!(v2378 != 0) && v2374 >= v2375)
	goto _2377
_2377:
	r221 = v2376
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	e221 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if r221 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35573, libc.VaList(bp+24, r221, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
	if e221&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:90:2:
		t_printf(tls, __ccgo_ts+35662, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = tinyf * tinyf * libc.Float32FromFloat32(0.5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v2384 = *(*uint32)(unsafe.Pointer(bp))
	goto _2385
_2385:
	if libc.BoolInt32(v2384&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = tinyf * tinyf
		v2383 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = tinyf * tinyf
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v2386 = *(*uint32)(unsafe.Pointer(bp))
		goto _2387
	_2387:
		v2383 = libc.BoolInt32(v2386&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r222 = v2383
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e222 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r222 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+35757, libc.VaList(bp+24, r222, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e222&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+35846, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	v2388 = float64(tinyf * tinyf * libc.Float32FromFloat32(0.5))
	v2389 = float64(tinyf * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2388
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2393 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2394
_2394:
	if libc.BoolInt32(v2393&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2389
		v2392 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2389
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2395 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2396
	_2396:
		v2392 = libc.BoolInt32(v2395&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2390 = libc.BoolInt32(!(v2392 != 0) && v2388 < v2389)
	goto _2391
_2391:
	r223 = v2390
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e223 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r223 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+35941, libc.VaList(bp+24, r223, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e223&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36025, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	v2397 = float64(tinyf * tinyf * libc.Float32FromFloat32(0.5))
	v2398 = float64(tinyf * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2397
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2402 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2403
_2403:
	if libc.BoolInt32(v2402&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2398
		v2401 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2398
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2404 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2405
	_2405:
		v2401 = libc.BoolInt32(v2404&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2399 = libc.BoolInt32(!(v2401 != 0) && v2397 <= v2398)
	goto _2400
_2400:
	r224 = v2399
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e224 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r224 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36115, libc.VaList(bp+24, r224, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e224&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36204, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	v2406 = float64(tinyf * tinyf * libc.Float32FromFloat32(0.5))
	v2407 = float64(tinyf * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2406
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2411 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2412
_2412:
	if libc.BoolInt32(v2411&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2407
		v2410 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2407
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2413 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2414
	_2414:
		v2410 = libc.BoolInt32(v2413&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2408 = libc.BoolInt32(!(v2410 != 0) && v2406 != v2407)
	goto _2409
_2409:
	r225 = v2408
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e225 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r225 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36299, libc.VaList(bp+24, r225, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e225&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36390, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	v2415 = float64(tinyf * tinyf * libc.Float32FromFloat32(0.5))
	v2416 = float64(tinyf * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2415
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2420 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2421
_2421:
	if libc.BoolInt32(v2420&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2416
		v2419 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2416
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2422 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2423
	_2423:
		v2419 = libc.BoolInt32(v2422&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2417 = libc.BoolInt32(!(v2419 != 0) && v2415 > v2416)
	goto _2418
_2418:
	r226 = v2417
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e226 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r226 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36487, libc.VaList(bp+24, r226, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e226&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36574, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	v2424 = float64(tinyf * tinyf * libc.Float32FromFloat32(0.5))
	v2425 = float64(tinyf * tinyf)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2424
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2429 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2430
_2430:
	if libc.BoolInt32(v2429&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2425
		v2428 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2425
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2431 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2432
	_2432:
		v2428 = libc.BoolInt32(v2431&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2426 = libc.BoolInt32(!(v2428 != 0) && v2424 >= v2425)
	goto _2427
_2427:
	r227 = v2426
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	e227 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if r227 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36667, libc.VaList(bp+24, r227, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
	if e227&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:91:2:
		t_printf(tls, __ccgo_ts+36759, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = -tinyf * tinyf
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v2434 = *(*uint32)(unsafe.Pointer(bp))
	goto _2435
_2435:
	if libc.BoolInt32(v2434&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(0)
		v2433 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v2436 = *(*uint32)(unsafe.Pointer(bp))
		goto _2437
	_2437:
		v2433 = libc.BoolInt32(v2436&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r228 = v2433
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e228 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r228 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+36857, libc.VaList(bp+24, r228, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e228&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+36935, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	v2438 = float64(-tinyf * tinyf)
	v2439 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2438
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2443 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2444
_2444:
	if libc.BoolInt32(v2443&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2439
		v2442 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2439
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2445 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2446
	_2446:
		v2442 = libc.BoolInt32(v2445&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2440 = libc.BoolInt32(!(v2442 != 0) && v2438 < v2439)
	goto _2441
_2441:
	r229 = v2440
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e229 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r229 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37019, libc.VaList(bp+24, r229, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e229&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37092, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	v2447 = float64(-tinyf * tinyf)
	v2448 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2447
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2452 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2453
_2453:
	if libc.BoolInt32(v2452&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2448
		v2451 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2448
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2454 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2455
	_2455:
		v2451 = libc.BoolInt32(v2454&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2449 = libc.BoolInt32(!(v2451 != 0) && v2447 <= v2448)
	goto _2450
_2450:
	r230 = v2449
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e230 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r230 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37171, libc.VaList(bp+24, r230, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e230&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37249, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	v2456 = float64(-tinyf * tinyf)
	v2457 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2456
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2461 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2462
_2462:
	if libc.BoolInt32(v2461&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2457
		v2460 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2457
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2463 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2464
	_2464:
		v2460 = libc.BoolInt32(v2463&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2458 = libc.BoolInt32(!(v2460 != 0) && v2456 != v2457)
	goto _2459
_2459:
	r231 = v2458
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e231 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r231 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37333, libc.VaList(bp+24, r231, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e231&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37413, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	v2465 = float64(-tinyf * tinyf)
	v2466 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2465
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2470 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2471
_2471:
	if libc.BoolInt32(v2470&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2466
		v2469 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2466
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2472 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2473
	_2473:
		v2469 = libc.BoolInt32(v2472&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2467 = libc.BoolInt32(!(v2469 != 0) && v2465 > v2466)
	goto _2468
_2468:
	r232 = v2467
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e232 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r232 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37499, libc.VaList(bp+24, r232, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e232&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37575, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	v2474 = float64(-tinyf * tinyf)
	v2475 = libc.Float64FromFloat32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2474
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2479 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2480
_2480:
	if libc.BoolInt32(v2479&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2475
		v2478 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2475
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2481 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2482
	_2482:
		v2478 = libc.BoolInt32(v2481&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2476 = libc.BoolInt32(!(v2478 != 0) && v2474 >= v2475)
	goto _2477
_2477:
	r233 = v2476
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	e233 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if r233 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37657, libc.VaList(bp+24, r233, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
	if e233&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:92:2:
		t_printf(tls, __ccgo_ts+37738, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v2484 = *(*uint32)(unsafe.Pointer(bp))
	goto _2485
_2485:
	if libc.BoolInt32(v2484&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		_ = libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25)
		v2483 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v2486 = *(*uint32)(unsafe.Pointer(bp))
		goto _2487
	_2487:
		v2483 = libc.BoolInt32(v2486&uint32(0x7fffffff) > uint32(0x7f800000))
	}
	r234 = v2483
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e234 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r234 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+37825, libc.VaList(bp+24, r234, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e234&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+37906, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	v2488 = libc.Float64FromFloat32(1)
	v2489 = float64(libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2488
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2493 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2494
_2494:
	if libc.BoolInt32(v2493&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2489
		v2492 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2489
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2495 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2496
	_2496:
		v2492 = libc.BoolInt32(v2495&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2490 = libc.BoolInt32(!(v2492 != 0) && v2488 < v2489)
	goto _2491
_2491:
	r235 = v2490
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e235 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r235 != libc.BoolInt32(true) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+37993, libc.VaList(bp+24, r235, libc.BoolInt32(true)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e235&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38069, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	v2497 = libc.Float64FromFloat32(1)
	v2498 = float64(libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2497
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2502 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2503
_2503:
	if libc.BoolInt32(v2502&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2498
		v2501 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2498
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2504 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2505
	_2505:
		v2501 = libc.BoolInt32(v2504&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2499 = libc.BoolInt32(!(v2501 != 0) && v2497 <= v2498)
	goto _2500
_2500:
	r236 = v2499
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e236 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r236 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38151, libc.VaList(bp+24, r236, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e236&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38232, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	v2506 = libc.Float64FromFloat32(1)
	v2507 = float64(libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:114:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2506
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2511 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2512
_2512:
	if libc.BoolInt32(v2511&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2507
		v2510 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2507
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2513 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2514
	_2514:
		v2510 = libc.BoolInt32(v2513&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2508 = libc.BoolInt32(!(v2510 != 0) && v2506 != v2507)
	goto _2509
_2509:
	r237 = v2508
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e237 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r237 != libc.BoolInt32(libc.Bool(true) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38319, libc.VaList(bp+24, r237, libc.BoolInt32(libc.Bool(true) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e237&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38402, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	v2515 = libc.Float64FromFloat32(1)
	v2516 = float64(libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2515
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2520 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2521
_2521:
	if libc.BoolInt32(v2520&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2516
		v2519 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2516
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2522 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2523
	_2523:
		v2519 = libc.BoolInt32(v2522&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2517 = libc.BoolInt32(!(v2519 != 0) && v2515 > v2516)
	goto _2518
_2518:
	r238 = v2517
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e238 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r238 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38491, libc.VaList(bp+24, r238, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e238&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38570, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	v2524 = libc.Float64FromFloat32(1)
	v2525 = float64(libc.Float32FromFloat32(1) + epsf*libc.Float32FromFloat32(0.25))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2524
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2529 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2530
_2530:
	if libc.BoolInt32(v2529&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2525
		v2528 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2525
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2531 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2532
	_2532:
		v2528 = libc.BoolInt32(v2531&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2526 = libc.BoolInt32(!(v2528 != 0) && v2524 >= v2525)
	goto _2527
_2527:
	r239 = v2526
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	e239 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if r239 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38655, libc.VaList(bp+24, r239, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
	if e239&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:93:2:
		t_printf(tls, __ccgo_ts+38739, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = hugel * hugel
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2534 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2535
_2535:
	if libc.BoolInt32(v2534&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = hugel * hugel * libc.Float64FromInt32(2)
		v2533 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = hugel * hugel * libc.Float64FromInt32(2)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2536 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2537
	_2537:
		v2533 = libc.BoolInt32(v2536&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r240 = v2533
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e240 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r240 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+38829, libc.VaList(bp+24, r240, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e240&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+38916, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	v2538 = hugel * hugel
	v2539 = hugel * hugel * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2538
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2543 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2544
_2544:
	if libc.BoolInt32(v2543&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2539
		v2542 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2539
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2545 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2546
	_2546:
		v2542 = libc.BoolInt32(v2545&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2540 = libc.BoolInt32(!(v2542 != 0) && v2538 < v2539)
	goto _2541
_2541:
	r241 = v2540
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e241 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r241 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39009, libc.VaList(bp+24, r241, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e241&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39091, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	v2547 = hugel * hugel
	v2548 = hugel * hugel * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2547
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2552 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2553
_2553:
	if libc.BoolInt32(v2552&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2548
		v2551 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2548
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2554 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2555
	_2555:
		v2551 = libc.BoolInt32(v2554&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2549 = libc.BoolInt32(!(v2551 != 0) && v2547 <= v2548)
	goto _2550
_2550:
	r242 = v2549
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e242 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r242 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39179, libc.VaList(bp+24, r242, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e242&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39266, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	v2556 = hugel * hugel
	v2557 = hugel * hugel * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2556
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2561 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2562
_2562:
	if libc.BoolInt32(v2561&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2557
		v2560 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2557
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2563 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2564
	_2564:
		v2560 = libc.BoolInt32(v2563&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2558 = libc.BoolInt32(!(v2560 != 0) && v2556 != v2557)
	goto _2559
_2559:
	r243 = v2558
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e243 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r243 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39359, libc.VaList(bp+24, r243, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e243&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39448, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	v2565 = hugel * hugel
	v2566 = hugel * hugel * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2565
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2570 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2571
_2571:
	if libc.BoolInt32(v2570&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2566
		v2569 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2566
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2572 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2573
	_2573:
		v2569 = libc.BoolInt32(v2572&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2567 = libc.BoolInt32(!(v2569 != 0) && v2565 > v2566)
	goto _2568
_2568:
	r244 = v2567
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e244 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r244 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39543, libc.VaList(bp+24, r244, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e244&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39628, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	v2574 = hugel * hugel
	v2575 = hugel * hugel * libc.Float64FromInt32(2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2574
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2579 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2580
_2580:
	if libc.BoolInt32(v2579&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2575
		v2578 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2575
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2581 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2582
	_2582:
		v2578 = libc.BoolInt32(v2581&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2576 = libc.BoolInt32(!(v2578 != 0) && v2574 >= v2575)
	goto _2577
_2577:
	r245 = v2576
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	e245 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if r245 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39719, libc.VaList(bp+24, r245, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
	if e245&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:101:2:
		t_printf(tls, __ccgo_ts+39809, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2584 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2585
_2585:
	if libc.BoolInt32(v2584&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = tinyl * tinyl
		v2583 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = tinyl * tinyl
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2586 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2587
	_2587:
		v2583 = libc.BoolInt32(v2586&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r246 = v2583
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e246 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r246 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+39905, libc.VaList(bp+24, r246, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e246&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+39995, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	v2588 = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	v2589 = tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2588
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2593 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2594
_2594:
	if libc.BoolInt32(v2593&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2589
		v2592 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2589
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2595 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2596
	_2596:
		v2592 = libc.BoolInt32(v2595&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2590 = libc.BoolInt32(!(v2592 != 0) && v2588 < v2589)
	goto _2591
_2591:
	r247 = v2590
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e247 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r247 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40091, libc.VaList(bp+24, r247, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e247&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40176, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	v2597 = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	v2598 = tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2597
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2602 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2603
_2603:
	if libc.BoolInt32(v2602&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2598
		v2601 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2598
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2604 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2605
	_2605:
		v2601 = libc.BoolInt32(v2604&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2599 = libc.BoolInt32(!(v2601 != 0) && v2597 <= v2598)
	goto _2600
_2600:
	r248 = v2599
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e248 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r248 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40267, libc.VaList(bp+24, r248, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e248&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40357, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	v2606 = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	v2607 = tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2606
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2611 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2612
_2612:
	if libc.BoolInt32(v2611&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2607
		v2610 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2607
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2613 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2614
	_2614:
		v2610 = libc.BoolInt32(v2613&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2608 = libc.BoolInt32(!(v2610 != 0) && v2606 != v2607)
	goto _2609
_2609:
	r249 = v2608
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e249 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r249 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40453, libc.VaList(bp+24, r249, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e249&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40545, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	v2615 = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	v2616 = tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2615
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2620 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2621
_2621:
	if libc.BoolInt32(v2620&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2616
		v2619 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2616
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2622 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2623
	_2623:
		v2619 = libc.BoolInt32(v2622&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2617 = libc.BoolInt32(!(v2619 != 0) && v2615 > v2616)
	goto _2618
_2618:
	r250 = v2617
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e250 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r250 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40643, libc.VaList(bp+24, r250, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e250&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40731, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	v2624 = tinyl * tinyl * libc.Float64FromFloat64(0.5)
	v2625 = tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2624
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2629 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2630
_2630:
	if libc.BoolInt32(v2629&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2625
		v2628 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2625
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2631 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2632
	_2632:
		v2628 = libc.BoolInt32(v2631&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2626 = libc.BoolInt32(!(v2628 != 0) && v2624 >= v2625)
	goto _2627
_2627:
	r251 = v2626
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	e251 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if r251 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40825, libc.VaList(bp+24, r251, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
	if e251&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:102:2:
		t_printf(tls, __ccgo_ts+40918, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = -tinyl * tinyl
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2634 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2635
_2635:
	if libc.BoolInt32(v2634&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(0)
		v2633 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2636 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2637
	_2637:
		v2633 = libc.BoolInt32(v2636&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r252 = v2633
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e252 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r252 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41017, libc.VaList(bp+24, r252, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e252&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41096, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	v2638 = -tinyl * tinyl
	v2639 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2638
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2643 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2644
_2644:
	if libc.BoolInt32(v2643&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2639
		v2642 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2639
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2645 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2646
	_2646:
		v2642 = libc.BoolInt32(v2645&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2640 = libc.BoolInt32(!(v2642 != 0) && v2638 < v2639)
	goto _2641
_2641:
	r253 = v2640
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e253 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r253 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41181, libc.VaList(bp+24, r253, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e253&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41255, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	v2647 = -tinyl * tinyl
	v2648 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2647
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2652 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2653
_2653:
	if libc.BoolInt32(v2652&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2648
		v2651 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2648
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2654 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2655
	_2655:
		v2651 = libc.BoolInt32(v2654&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2649 = libc.BoolInt32(!(v2651 != 0) && v2647 <= v2648)
	goto _2650
_2650:
	r254 = v2649
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e254 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r254 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41335, libc.VaList(bp+24, r254, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e254&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41414, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	v2656 = -tinyl * tinyl
	v2657 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2656
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2661 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2662
_2662:
	if libc.BoolInt32(v2661&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2657
		v2660 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2657
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2663 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2664
	_2664:
		v2660 = libc.BoolInt32(v2663&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2658 = libc.BoolInt32(!(v2660 != 0) && v2656 != v2657)
	goto _2659
_2659:
	r255 = v2658
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e255 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r255 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41499, libc.VaList(bp+24, r255, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e255&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41580, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	v2665 = -tinyl * tinyl
	v2666 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2665
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2670 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2671
_2671:
	if libc.BoolInt32(v2670&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2666
		v2669 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2666
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2672 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2673
	_2673:
		v2669 = libc.BoolInt32(v2672&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2667 = libc.BoolInt32(!(v2669 != 0) && v2665 > v2666)
	goto _2668
_2668:
	r256 = v2667
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e256 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r256 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41667, libc.VaList(bp+24, r256, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e256&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41744, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	v2674 = -tinyl * tinyl
	v2675 = libc.Float64FromFloat64(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2674
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2679 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2680
_2680:
	if libc.BoolInt32(v2679&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2675
		v2678 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2675
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2681 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2682
	_2682:
		v2678 = libc.BoolInt32(v2681&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2676 = libc.BoolInt32(!(v2678 != 0) && v2674 >= v2675)
	goto _2677
_2677:
	r257 = v2676
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	e257 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if r257 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41827, libc.VaList(bp+24, r257, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
	if e257&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:103:2:
		t_printf(tls, __ccgo_ts+41909, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2684 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2685
_2685:
	if libc.BoolInt32(v2684&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
		v2683 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2686 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2687
	_2687:
		v2683 = libc.BoolInt32(v2686&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	r258 = v2683
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e258 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r258 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+41997, libc.VaList(bp+24, r258, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e258&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42079, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	v2688 = libc.Float64FromFloat64(1)
	v2689 = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2688
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2693 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2694
_2694:
	if libc.BoolInt32(v2693&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2689
		v2692 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2689
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2695 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2696
	_2696:
		v2692 = libc.BoolInt32(v2695&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2690 = libc.BoolInt32(!(v2692 != 0) && v2688 < v2689)
	goto _2691
_2691:
	r259 = v2690
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e259 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r259 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42167, libc.VaList(bp+24, r259, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e259&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42244, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	v2697 = libc.Float64FromFloat64(1)
	v2698 = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2697
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2702 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2703
_2703:
	if libc.BoolInt32(v2702&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2698
		v2701 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2698
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2704 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2705
	_2705:
		v2701 = libc.BoolInt32(v2704&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2699 = libc.BoolInt32(!(v2701 != 0) && v2697 <= v2698)
	goto _2700
_2700:
	r260 = v2699
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e260 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r260 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42327, libc.VaList(bp+24, r260, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e260&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42409, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	v2706 = libc.Float64FromFloat64(1)
	v2707 = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:115:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2706
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2711 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2712
_2712:
	if libc.BoolInt32(v2711&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2707
		v2710 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2707
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2713 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2714
	_2714:
		v2710 = libc.BoolInt32(v2713&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2708 = libc.BoolInt32(!(v2710 != 0) && v2706 != v2707)
	goto _2709
_2709:
	r261 = v2708
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e261 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r261 != libc.BoolInt32(libc.Bool(false) || libc.Bool(false)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42497, libc.VaList(bp+24, r261, libc.BoolInt32(libc.Bool(false) || libc.Bool(false))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e261&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42581, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	v2715 = libc.Float64FromFloat64(1)
	v2716 = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2715
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2720 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2721
_2721:
	if libc.BoolInt32(v2720&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2716
		v2719 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2716
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2722 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2723
	_2723:
		v2719 = libc.BoolInt32(v2722&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2717 = libc.BoolInt32(!(v2719 != 0) && v2715 > v2716)
	goto _2718
_2718:
	r262 = v2717
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e262 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r262 != libc.BoolInt32(false) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42671, libc.VaList(bp+24, r262, libc.BoolInt32(false)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e262&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42751, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	feclearexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	v2724 = libc.Float64FromFloat64(1)
	v2725 = libc.Float64FromFloat64(1) + epsl*libc.Float64FromFloat64(0.25)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp + 8)) = v2724
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v2729 = *(*uint64)(unsafe.Pointer(bp + 8))
	goto _2730
_2730:
	if libc.BoolInt32(v2729&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		_ = v2725
		v2728 = libc.Int32FromInt32(1)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp + 8)) = v2725
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v2731 = *(*uint64)(unsafe.Pointer(bp + 8))
		goto _2732
	_2732:
		v2728 = libc.BoolInt32(v2731&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	v2726 = libc.BoolInt32(!(v2728 != 0) && v2724 >= v2725)
	goto _2727
_2727:
	r263 = v2726
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	e263 = fetestexcept(tls, int32(FE_ALL_EXCEPT))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if r263 != libc.BoolInt32(libc.Bool(false) || libc.Bool(true)) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42837, libc.VaList(bp+24, r263, libc.BoolInt32(libc.Bool(false) || libc.Bool(true))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
	if e263&int32(FE_INVALID) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:104:2:
		t_printf(tls, __ccgo_ts+42922, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/isless.c:106:2:
	return libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&t_status)))
}

func main() {
	libc.Start(main1)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:24:9:
const EMFILE = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:44:9:
const O_RDONLY = 0

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:175:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:175:18:
type mode_t = uint32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:378:1:
type iovec = struct {
	Fiov_base uintptr
	Fiov_len  size_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:24:1:
type flock1 = struct {
	Fl_type   int16
	Fl_whence int16
	Fl_start  off_t
	Fl_len    off_t
	Fl_pid    pid_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:167:1:
type file_handle = struct {
	Fhandle_bytes uint32
	Fhandle_type  int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:172:1:
type f_owner_ex = struct {
	Ftype1 int32
	Fpid   pid_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:6:6:
func t_fdfill(tls *libc.TLS) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:7:1:
	var fd int32
	_ = fd
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:8:6:
	fd = int32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:9:2:
	if libc.Xdup(tls, fd) == -int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:10:3:
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(EMFILE) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:11:4:
			return
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:12:3:
		fd = libc.Xopen(tls, __ccgo_ts+43013, O_RDONLY, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:14:2:
	for libc.Xdup(tls, fd) != -int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/fdfill.c:14:22:
	}
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/select.h:18:9:
const FD_SETSIZE = 1024

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/resource.h:78:9:
const RLIMIT_DATA = 2

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:12:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:12:24:
type wchar_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:64:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:64:35:
type div_t = struct {
	Fquot int32
	Frem  int32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:65:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:65:36:
type ldiv_t = struct {
	Fquot int32
	Frem  int32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:66:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdlib.h:66:41:
type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:32:
type locale_t = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:108:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:108:16:
type time_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:113:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:113:16:
type suseconds_t = int64

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:247:1:
type timeval = struct {
	Ftv_sec  time_t
	Ftv_usec suseconds_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:252:1:
type timespec = struct {
	Ftv_sec   time_t
	Ftv_nsec  int32
	F__ccgo12 uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:372:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:372:71:
type sigset_t = struct {
	F__bits [32]uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:372:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:372:71:
type __sigset_t = sigset_t

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/select.h:20:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/select.h:20:23:
type fd_mask = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/select.h:22:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/select.h:24:3:
type fd_set = struct {
	Ffds_bits [32]uint32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/time.h:17:1:
type itimerval = struct {
	Fit_interval timeval
	Fit_value    timeval
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/time.h:27:1:
type timezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:263:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:263:18:
type id_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/resource.h:20:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/resource.h:20:28:
type rlim_t = uint64

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/resource.h:22:1:
type rlimit = struct {
	Frlim_cur rlim_t
	Frlim_max rlim_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/resource.h:27:1:
type rusage = struct {
	Fru_utime    timeval
	Fru_stime    timeval
	Fru_maxrss   int32
	Fru_ixrss    int32
	Fru_idrss    int32
	Fru_isrss    int32
	Fru_minflt   int32
	Fru_majflt   int32
	Fru_nswap    int32
	Fru_inblock  int32
	Fru_oublock  int32
	Fru_msgsnd   int32
	Fru_msgrcv   int32
	Fru_nsignals int32
	Fru_nvcsw    int32
	Fru_nivcsw   int32
	F__reserved  [16]int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:7:5:
func t_memfill(tls *libc.TLS) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:8:1:
	var r int32
	_ = r
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:9:6:
	r = 0
	/* alloc mmap space with PROT_NONE */
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:11:2:
	if t_vmfill(tls, uintptr(0), uintptr(0), 0) < 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:12:3:
		t_printf(tls, __ccgo_ts+43023, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:13:3:
		r = -int32(1)
	}
	/* limit brk space */
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:16:2:
	if t_setrlim(tls, int32(RLIMIT_DATA), 0) < 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:17:3:
		r = -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:18:2:
	if !(r != 0) {
		/* use up libc reserves if any */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:20:3:
		for libc.Xmalloc(tls, uint32(1)) != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:20:20:
		}
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/memfill.c:21:2:
	return r
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:3:9:
const FE_DIVBYZERO = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:11:9:
const FE_DOWNWARD = 1024

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:4:9:
const FE_OVERFLOW = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:13:9:
const FE_TOWARDZERO = 3072

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:12:9:
const FE_UPWARD = 2048

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:34:9:
const FP_INFINITE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:338:1:
type _IO_FILE = struct {
	F__x int8
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:343:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:343:25:
type FILE = struct {
	F__x int8
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:349:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:349:27:
type va_list = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:354:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:354:27:
type __isoc_va_list = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:56:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:60:3:
type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:56:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:60:3:
type _G_fpos64_t = fpos_t

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:198:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:203:3:
type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:198:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdio.h:203:3:
type _IO_cookie_io_functions_t = cookie_io_functions_t

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:102:12:
func checkexcept1(tls *libc.TLS, got int32, want int32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:103:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:104:2:
	if r == FE_TONEAREST {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:110:3:
		return libc.BoolInt32(got|int32(FE_INEXACT) == want|int32(FE_INEXACT))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:112:2:
	return libc.BoolInt32(got|int32(FE_INEXACT)|int32(FE_UNDERFLOW) == want|int32(FE_INEXACT)|int32(FE_UNDERFLOW))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:115:12:
func checkexceptall1(tls *libc.TLS, got int32, want int32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:116:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:117:2:
	return libc.BoolInt32(got == want)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:120:12:
func checkulp1(tls *libc.TLS, d float32, r int32) (r1 int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:123:2:
	if r == FE_TONEAREST {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:124:3:
		return libc.BoolInt32(float64(libc.Xfabsf(tls, d)) < float64(1.5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:126:2:
	return libc.BoolInt32(float64(libc.Xfabsf(tls, d)) < float64(3))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:129:12:
func checkcr1(tls *libc.TLS, y float64, ywant float64, r int32) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:130:1:
	var v1, v3, v5, v7 uint64
	var v9 bool
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	_, _, _, _, _ = v1, v3, v5, v7, v9
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:131:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp)) = float64(ywant)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:132:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(y)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
		return libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:133:2:
	if v9 = y == ywant; v9 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(y)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v5 = *(*uint64)(unsafe.Pointer(bp))
		goto _6
	_6:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(ywant)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v7 = *(*uint64)(unsafe.Pointer(bp))
		goto _8
	_8:
	}
	return libc.BoolInt32(v9 && int32(v5>>libc.Int32FromInt32(63)) == int32(v7>>libc.Int32FromInt32(63)))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:5:5:
func eulpf(tls *libc.TLS, x float32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:6:1:
	var e int32
	var _ /* u at bp+0 */ struct {
		Fi [0]uint32_t
		Ff float32
	}
	_ = e
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:7:33:
	*(*struct {
		Fi [0]uint32_t
		Ff float32
	})(unsafe.Pointer(bp)) = struct {
		Fi [0]uint32_t
		Ff float32
	}{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:7:33:
	*(*float32)(unsafe.Pointer(bp)) = x
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:8:6:
	e = int32(*(*uint32_t)(unsafe.Pointer(bp)) >> int32(23) & uint32(0xff))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:10:2:
	if !(e != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:11:3:
		e++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:12:2:
	return e - int32(0x7f) - int32(23)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:15:5:
func eulp(tls *libc.TLS, x float64) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:16:1:
	var e int32
	var _ /* u at bp+0 */ struct {
		Fi [0]uint64_t
		Ff float64
	}
	_ = e
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:17:34:
	*(*struct {
		Fi [0]uint64_t
		Ff float64
	})(unsafe.Pointer(bp)) = struct {
		Fi [0]uint64_t
		Ff float64
	}{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:17:34:
	*(*float64)(unsafe.Pointer(bp)) = x
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:18:6:
	e = int32(*(*uint64_t)(unsafe.Pointer(bp)) >> int32(52) & uint64(0x7ff))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:20:2:
	if !(e != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:21:3:
		e++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:22:2:
	return e - int32(0x3ff) - int32(52)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:25:5:
func eulpl(tls *libc.TLS, x float64) (r int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:26:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:28:2:
	return eulp(tls, float64(x))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:42:7:
func ulperrf(tls *libc.TLS, got float32, want float32, dwant float32) (r float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:43:1:
	var v1, v10, v3, v6, v8 uint32
	var v5 bool
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint32
		F__f float32
	}
	_, _, _, _, _, _ = v1, v10, v3, v5, v6, v8
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:44:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = got
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v1 = *(*uint32)(unsafe.Pointer(bp))
	goto _2
_2:
	if v5 = libc.BoolInt32(v1&uint32(0x7fffffff) > uint32(0x7f800000)) != 0; v5 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = want
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v3 = *(*uint32)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 && libc.BoolInt32(v3&uint32(0x7fffffff) > uint32(0x7f800000)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:45:3:
		return libc.Float32FromInt32(0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:46:2:
	if got == want {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:47:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = got
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v6 = *(*uint32)(unsafe.Pointer(bp))
		goto _7
	_7:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
		*(*float32)(unsafe.Pointer(bp)) = want
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
		v8 = *(*uint32)(unsafe.Pointer(bp))
		goto _9
	_9:
		if int32(v6>>libc.Int32FromInt32(31)) == int32(v8>>libc.Int32FromInt32(31)) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:48:4:
			return dwant
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:49:3:
		return libc.X__builtin_inff(tls)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:51:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:58:2:
	*(*float32)(unsafe.Pointer(bp)) = got
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:59:2:
	v10 = *(*uint32)(unsafe.Pointer(bp))
	goto _11
_11:
	if libc.BoolInt32(v10&uint32(0x7fffffff) == uint32(0x7f800000)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:52:3:
		got = libc.Xcopysignf(tls, float32(1.7014118346046923e+38), got)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:53:3:
		want = float32(float64(want) * libc.Float64FromFloat64(0.5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:55:2:
	return float32(libc.Xscalbn(tls, float64(got-want), -eulpf(tls, want)) + float64(dwant))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:58:7:
func ulperr(tls *libc.TLS, got float64, want float64, dwant float32) (r float32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:59:1:
	var v1, v10, v3, v6, v8 uint64
	var v5 bool
	var _ /* __u at bp+0 */ struct {
		F__i [0]uint64
		F__f float64
	}
	_, _, _, _, _, _ = v1, v10, v3, v5, v6, v8
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:60:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp)) = got
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if v5 = libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0; v5 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = want
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
	}
	if v5 && libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:61:3:
		return libc.Float32FromInt32(0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:62:2:
	if got == want {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:63:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = got
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v6 = *(*uint64)(unsafe.Pointer(bp))
		goto _7
	_7:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = want
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v8 = *(*uint64)(unsafe.Pointer(bp))
		goto _9
	_9:
		if int32(v6>>libc.Int32FromInt32(63)) == int32(v8>>libc.Int32FromInt32(63)) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:64:4:
			return dwant
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:65:3:
		return libc.X__builtin_inff(tls) // treat 0 sign errors badly
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:67:2:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
	*(*float64)(unsafe.Pointer(bp)) = got
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v10 = *(*uint64)(unsafe.Pointer(bp))
	goto _11
_11:
	if libc.BoolInt32(v10&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) == libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:68:3:
		got = libc.Xcopysign(tls, float64(8.98846567431158e+307), got)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:69:3:
		want *= float64(0.5)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:71:2:
	return float32(libc.Xscalbn(tls, got-want, -eulp(tls, want)) + float64(dwant))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:74:7:
func ulperrl(tls *libc.TLS, got float64, want float64, dwant float32) (r float32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:75:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:77:2:
	return ulperr(tls, float64(got), float64(want), dwant)
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:99:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:102:3:
var eflags = [5]struct {
	Fflag int32
	Fs    uintptr
}{
	0: {
		Fflag: int32(FE_INEXACT),
		Fs:    __ccgo_ts + 43067,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 43075,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 43083,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 43093,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 43103,
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:110:6:
func estr(tls *libc.TLS, f int32) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:111:1:
	var all, i int32
	var p, v2, v3, v4 uintptr
	_, _, _, _, _, _ = all, i, p, v2, v3, v4
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:113:7:
	p = uintptr(unsafe.Pointer(&buf))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:114:9:
	all = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:116:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(8)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:117:3:
		if f&eflags[i].Fflag != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:118:4:
			if all != 0 {
				v2 = __ccgo_ts + 43112
			} else {
				v2 = __ccgo_ts
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+43114, libc.VaList(bp+8, v2, eflags[i].Fs)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:119:4:
			all |= eflags[i].Fflag
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:121:2:
	if all != f {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:122:3:
		if all != 0 {
			v3 = __ccgo_ts + 43112
		} else {
			v3 = __ccgo_ts
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+43119, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts
	} else {
		v4 = __ccgo_ts + 43124
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+43126, libc.VaList(bp+8, v4)))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:126:2:
	return uintptr(unsafe.Pointer(&buf))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:112:14:
var buf [256]int8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:129:6:
func rstr(tls *libc.TLS, r int32) (r1 uintptr) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:130:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:131:2:
	switch r {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:132:2:
	case FE_TONEAREST:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:132:11:
		return __ccgo_ts + 43129
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 43132
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 43135
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 43138
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 43141
}

// C documentation
//
//	/* relative path to p */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:6:6:
func t_pathrel(tls *libc.TLS, buf uintptr, n size_t, argv0 uintptr, p uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:7:1:
	var k int32
	var s uintptr
	_, _ = k, s
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:8:7:
	s = libc.Xstrrchr(tls, argv0, int32('/'))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:11:2:
	if s != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:12:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+43144, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+43152, libc.VaList(bp+8, p))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:15:2:
	if uint32(k) >= n {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:16:3:
		return uintptr(0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:17:2:
	return buf
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:8:5:
func t_printf(tls *libc.TLS, s uintptr, va uintptr) (r int32) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:9:1:
	var ap va_list
	var n int32
	var _ /* buf at bp+0 */ [512]int8
	_, _ = ap, n
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:14:2:
	libc.AtomicStorePInt32(uintptr(unsafe.Pointer(&t_status)), int32(1))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:15:2:
	ap = va
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:16:2:
	n = libc.Xvsnprintf(tls, bp, uint32(512), s, ap)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:17:2:
	_ = ap
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:18:2:
	if n < 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:19:3:
		n = 0
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:20:7:
		if uint32(n) >= uint32(512) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:21:3:
			n = int32(512)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:22:3:
			(*(*[512]int8)(unsafe.Pointer(bp)))[n-int32(1)] = int8('\n')
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:23:3:
			(*(*[512]int8)(unsafe.Pointer(bp)))[n-int32(2)] = int8('.')
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:24:3:
			(*(*[512]int8)(unsafe.Pointer(bp)))[n-int32(3)] = int8('.')
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:25:3:
			(*(*[512]int8)(unsafe.Pointer(bp)))[n-int32(4)] = int8('.')
		}
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:27:2:
	return libc.Xwrite(tls, int32(1), bp, uint32(n))
}

// C documentation
//
//	// TODO: use large period prng
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:6:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:6:17:
var seed = uint64(-libc.Int32FromInt32(1))

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:7:17:
func rand32(tls *libc.TLS) (r uint32_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:8:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:9:2:
	seed = uint64(6364136223846793005)*seed + uint64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:10:2:
	return uint32(seed >> int32(32))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:12:17:
func rand64(tls *libc.TLS) (r uint64_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:13:1:
	var u uint64_t
	_ = u
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:14:11:
	u = uint64(rand32(tls))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:15:2:
	return u<<int32(32) | uint64(rand32(tls))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:17:15:
func frand(tls *libc.TLS) (r float64) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:18:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:19:2:
	return float64(rand64(tls)) * float64(5.421010862427522e-20)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:21:14:
func frandf(tls *libc.TLS) (r float32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:22:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:23:2:
	return float32(rand32(tls)) * libc.Float32FromFloat32(2.3283064365386963e-10)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:25:20:
func frandl(tls *libc.TLS) (r float64) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:26:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:27:2:
	return float64(rand64(tls)) * libc.Float64FromFloat64(5.42101086242752217003726400434970855712890625e-20)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:34:6:
func t_randseed(tls *libc.TLS, s uint64_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:35:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:36:2:
	seed = s
}

// C documentation
//
//	/* uniform random in [0,n), n > 0 must hold */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:40:10:
func t_randn(tls *libc.TLS, n uint64_t) (r1 uint64_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:41:1:
	var m, r, v1 uint64_t
	_, _, _ = m, r, v1
	/* m is the largest multiple of n */
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:45:2:
	m = uint64(-libc.Int32FromInt32(1))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:46:2:
	m -= m % n
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:47:2:
	for {
		v1 = rand64(tls)
		r = v1
		if !(v1 >= m) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:47:29:
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:48:2:
	return r % n
}

// C documentation
//
//	/* uniform on [a,b], a <= b must hold */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:52:10:
func t_randint(tls *libc.TLS, a uint64_t, b uint64_t) (r uint64_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:53:1:
	var n uint64_t
	_ = n
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:54:11:
	n = b - a + uint64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:55:2:
	if n != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:56:3:
		return a + t_randn(tls, n)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:57:2:
	return rand64(tls)
}

// C documentation
//
//	/* shuffle the elements of p and q until the elements in p are well shuffled */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:61:13:
func shuffle2(tls *libc.TLS, p uintptr, q uintptr, np size_t, nq size_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:62:1:
	var r, v1 size_t
	var t uint64_t
	_, _, _ = r, t, v1
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:66:2:
	for np != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:67:3:
		v1 = np
		np--
		r = uint32(t_randn(tls, uint64(nq+v1)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:68:3:
		t = *(*uint64_t)(unsafe.Pointer(p + uintptr(np)*8))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:69:3:
		if r < nq {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:70:4:
			*(*uint64_t)(unsafe.Pointer(p + uintptr(np)*8)) = *(*uint64_t)(unsafe.Pointer(q + uintptr(r)*8))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:71:4:
			*(*uint64_t)(unsafe.Pointer(q + uintptr(r)*8)) = t
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:73:4:
			*(*uint64_t)(unsafe.Pointer(p + uintptr(np)*8)) = *(*uint64_t)(unsafe.Pointer(p + uintptr(r-nq)*8))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:74:4:
			*(*uint64_t)(unsafe.Pointer(p + uintptr(r-nq)*8)) = t
		}
	}
}

// C documentation
//
//	/* shuffle the elements of p */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:80:6:
func t_shuffle(tls *libc.TLS, p uintptr, n size_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:81:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:82:2:
	shuffle2(tls, p, uintptr(0), n, uint32(0))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:85:6:
func t_randrange(tls *libc.TLS, p uintptr, n size_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:86:1:
	var i size_t
	_ = i
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:88:2:
	i = uint32(0)
	for {
		if !(i < n) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:89:3:
		*(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8)) = uint64(i)
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:90:2:
	t_shuffle(tls, p, n)
}

// C documentation
//
//	/* hash table insert, 0 means empty, v > 0 must hold, len is power-of-2 */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:94:12:
func insert(tls *libc.TLS, tab uintptr, len1 size_t, v uint64_t) (r int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:95:1:
	var i, j, v1 size_t
	_, _, _ = i, j, v1
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:96:9:
	i = uint32(v & uint64(len1-libc.Uint32FromInt32(1)))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:97:9:
	j = uint32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:99:2:
	for *(*uint64_t)(unsafe.Pointer(tab + uintptr(i)*8)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:100:3:
		if *(*uint64_t)(unsafe.Pointer(tab + uintptr(i)*8)) == v {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:101:4:
			return -int32(1)
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:102:3:
		v1 = j
		j++
		i += v1
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:103:3:
		i &= len1 - uint32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:105:2:
	*(*uint64_t)(unsafe.Pointer(tab + uintptr(i)*8)) = v
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:106:2:
	return 0
}

// C documentation
//
//	/* choose k unique numbers from [0,n), k <= n */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:110:5:
func t_choose(tls *libc.TLS, n uint64_t, k size_t, p uintptr) (r int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:111:1:
	var i, j, len1, v2 size_t
	var tab, v10 uintptr
	var v1 uint64_t
	_, _, _, _, _, _, _ = i, j, len1, tab, v1, v10, v2
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:115:2:
	if n < uint64(k) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:116:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:118:2:
	if n < uint64(16) {
		/* no alloc */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:120:3:
		for k != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:121:4:
			v1 = n
			n--
			if t_randn(tls, v1) < uint64(k) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:122:5:
				k--
				v2 = k
				*(*uint64_t)(unsafe.Pointer(p + uintptr(v2)*8)) = n
			}
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:123:3:
		return 0
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:126:2:
	if k < uint32(8) {
		/* no alloc, n > 15 > 2*k */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:128:3:
		i = uint32(0)
		for {
			if !(i < k) {
				break
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:129:4:
			*(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8)) = t_randn(tls, n)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:130:4:
			j = uint32(0)
			for {
				if !(*(*uint64_t)(unsafe.Pointer(p + uintptr(j)*8)) != *(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8))) {
					break
				}
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:130:34:
				goto _4
			_4:
				j++
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:131:4:
			if j == i {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:132:5:
				i++
			}
			goto _3
		_3:
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:134:3:
		return 0
	}
	// TODO: if k < n/k use k*log(k) solution without alloc
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:139:2:
	if n < uint64(uint32(5)*k) && (n-uint64(k))*uint64(8) < uint64(uint32(-libc.Int32FromInt32(1))) {
		/* allocation is n-k < 4*k */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:141:3:
		tab = libc.Xmalloc(tls, uint32((n-uint64(k))*uint64(8)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:142:3:
		if !(tab != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:143:4:
			return -int32(1)
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:144:3:
		i = uint32(0)
		for {
			if !(i < k) {
				break
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:145:4:
			*(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8)) = uint64(i)
			goto _5
		_5:
			i++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:146:3:
		for {
			if !(uint64(i) < n) {
				break
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:147:4:
			*(*uint64_t)(unsafe.Pointer(tab + uintptr(i-k)*8)) = uint64(i)
			goto _6
		_6:
			i++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:148:3:
		if uint64(k) < n-uint64(k) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:149:4:
			shuffle2(tls, p, tab, k, uint32(n-uint64(k)))
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:151:4:
			shuffle2(tls, tab, p, uint32(n-uint64(k)), k)
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:152:3:
		libc.Xfree(tls, tab)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:153:3:
		return 0
	}
	/* allocation is 2*k <= len < 4*k */
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:157:2:
	len1 = uint32(16)
	for {
		if !(len1 < uint32(2)*k) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:157:37:
		goto _7
	_7:
		len1 *= uint32(2)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:158:2:
	tab = libc.Xcalloc(tls, len1, uint32(8))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:159:2:
	if !(tab != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:160:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:161:2:
	i = uint32(0)
	for {
		if !(i < k) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:162:3:
		for insert(tls, tab, len1, t_randn(tls, n)+uint64(1)) != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:162:41:
		}
		goto _8
	_8:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:163:2:
	i = uint32(0)
	for {
		if !(i < len1) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:164:3:
		if *(*uint64_t)(unsafe.Pointer(tab + uintptr(i)*8)) != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:165:4:
			v10 = p
			p += 8
			*(*uint64_t)(unsafe.Pointer(v10)) = *(*uint64_t)(unsafe.Pointer(tab + uintptr(i)*8)) - uint64(1)
		}
		goto _9
	_9:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:166:2:
	libc.Xfree(tls, tab)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:167:2:
	return 0
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:6:5:
func t_setrlim(tls *libc.TLS, r int32, lim int32) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:7:1:
	var _ /* rl at bp+0 */ rlimit
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:10:2:
	if libc.Xgetrlimit(tls, r, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:11:3:
		t_printf(tls, __ccgo_ts+43157, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:12:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:14:2:
	if uint64(lim) > (*(*rlimit)(unsafe.Pointer(bp))).Frlim_max {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:15:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:16:2:
	if uint64(lim) == (*(*rlimit)(unsafe.Pointer(bp))).Frlim_max && uint64(lim) == (*(*rlimit)(unsafe.Pointer(bp))).Frlim_cur {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:17:3:
		return 0
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:18:2:
	(*(*rlimit)(unsafe.Pointer(bp))).Frlim_max = uint64(lim)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:19:2:
	(*(*rlimit)(unsafe.Pointer(bp))).Frlim_cur = uint64(lim)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:20:2:
	if libc.Xsetrlimit(tls, r, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:21:3:
		t_printf(tls, __ccgo_ts+43200, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:22:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:24:2:
	return 0
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/langinfo.h:71:9:
const CODESET = 14

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/locale.h:18:9:
const LC_CTYPE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/locale.h:26:1:
type lconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    int8
	Ffrac_digits        int8
	Fp_cs_precedes      int8
	Fp_sep_by_space     int8
	Fn_cs_precedes      int8
	Fn_sep_by_space     int8
	Fp_sign_posn        int8
	Fn_sign_posn        int8
	Fint_p_cs_precedes  int8
	Fint_p_sep_by_space int8
	Fint_n_cs_precedes  int8
	Fint_n_sep_by_space int8
	Fint_p_sign_posn    int8
	Fint_n_sign_posn    int8
	F__ccgo_pad24       [2]byte
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/nl_types.h:11:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/nl_types.h:11:13:
type nl_item = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/nl_types.h:12:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/nl_types.h:12:14:
type nl_catd = uintptr

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:6:5:
func t_setutf8(tls *libc.TLS) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:7:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:8:2:
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43249) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43257) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43269) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43281) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43293) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+43302) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+43302) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+43308, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:20:2:
	return 0
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/mman.h:26:9:
const MAP_ANON = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/mman.h:22:9:
const MAP_PRIVATE = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:46:9:
const O_RDWR = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/limits.h:1:9:
const PAGESIZE = 4096

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/mman.h:57:9:
const PROT_NONE = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/stdint.h:50:9:
const UINT32_MAX = 4294967295

// C documentation
//
//	/* max mmap size, *start is the largest power-of-2 size considered */
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:15:15:
func mmax(tls *libc.TLS, fd int32, start uintptr) (r size_t) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:16:1:
	var i, n, v2 size_t
	var p, v3 uintptr
	_, _, _, _, _ = i, n, p, v2, v3
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:20:2:
	v2 = *(*size_t)(unsafe.Pointer(start))
	n = v2
	i = v2
	for {
		if !(i >= uint32(PAGESIZE)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:21:3:
		v3 = libc.Xmmap(tls, uintptr(0), n, PROT_NONE, libc.Int32FromInt32(MAP_PRIVATE)|libc.Int32FromInt32(MAP_ANON), fd, 0)
		p = v3
		if v3 == uintptr(-libc.Int32FromInt32(1)) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:22:4:
			n -= i / uint32(2)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:24:4:
			libc.Xmunmap(tls, p, n)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:25:4:
			if n == i {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:26:5:
				*(*size_t)(unsafe.Pointer(start)) = n
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:27:4:
			n += i / uint32(2)
		}
		goto _1
	_1:
		i /= uint32(2)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:30:2:
	return n & uint32(-libc.Int32FromInt32(PAGESIZE))
}

// C documentation
//
//	/*
//	fills the virtual memory with anonymous PROT_NONE mmaps,
//	returns the mappings in *p and *n in decreasing size order,
//	the return value is the number of mappings or -1 on failure.
//	*/
//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:38:5:
func t_vmfill(tls *libc.TLS, p uintptr, n uintptr, len1 int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:39:1:
	var fd, i int32
	var m size_t
	var q uintptr
	var _ /* start at bp+0 */ size_t
	_, _, _, _ = fd, i, m, q
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:40:6:
	fd = -int32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:41:9:
	*(*size_t)(unsafe.Pointer(bp)) = libc.Uint32FromUint32(0xffffffff)/libc.Uint32FromInt32(2) + libc.Uint32FromInt32(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:46:2:
	i = 0
	for {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:47:3:
		m = mmax(tls, fd, bp)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:48:3:
		if !(m != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:49:4:
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:50:3:
		q = libc.Xmmap(tls, uintptr(0), m, PROT_NONE, libc.Int32FromInt32(MAP_PRIVATE)|libc.Int32FromInt32(MAP_ANON), fd, 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:51:3:
		if q == uintptr(-libc.Int32FromInt32(1)) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:52:4:
			return -int32(1)
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:53:3:
		if i < len1 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:54:4:
			*(*uintptr)(unsafe.Pointer(p + uintptr(i)*4)) = q
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:55:4:
			*(*size_t)(unsafe.Pointer(n + uintptr(i)*4)) = m
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/vmfill.c:58:2:
	return i
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:6:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:6:14:
var t_status int32

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "\x00src/math/isless.c:41: isunordered(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: isunordered(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:41: isless(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: isless(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:41: islessequal(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: islessequal(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:41: islessgreater(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: islessgreater(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:41: isgreater(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: isgreater(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:41: isgreaterequal(__builtin_nanf(\"\"), 1.0) failed: got %d want %d\n\x00src/math/isless.c:41: isgreaterequal(__builtin_nanf(\"\"), 1.0) raised the invalid exception\n\x00src/math/isless.c:42: isunordered(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: isunordered(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:42: isless(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: isless(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:42: islessequal(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: islessequal(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:42: islessgreater(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: islessgreater(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:42: isgreater(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: isgreater(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:42: isgreaterequal(1.0, __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:42: isgreaterequal(1.0, __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:43: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")) failed: got %d want %d\n\x00src/math/isless.c:43: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")) raised the invalid exception\n\x00src/math/isless.c:44: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:44: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:44: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:44: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:44: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:44: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) failed: got %d want %d\n\x00src/math/isless.c:44: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0) raised the invalid exception\n\x00src/math/isless.c:45: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: isunordered(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:45: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: isless(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:45: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: islessequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:45: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: islessgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:45: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: isgreater(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:45: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) failed: got %d want %d\n\x00src/math/isless.c:45: isgreaterequal(__builtin_nanf(\"\"), __builtin_nanf(\"\")+1.0L) raised the invalid exception\n\x00src/math/isless.c:47: isunordered(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: isunordered(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:47: isless(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: isless(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:47: islessequal(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: islessequal(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:47: islessgreater(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: islessgreater(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:47: isgreater(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: isgreater(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:47: isgreaterequal(1.0, 1.1) failed: got %d want %d\n\x00src/math/isless.c:47: isgreaterequal(1.0, 1.1) raised the invalid exception\n\x00src/math/isless.c:48: isunordered(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: isunordered(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:48: isless(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: isless(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:48: islessequal(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: islessequal(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:48: islessgreater(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: islessgreater(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:48: isgreater(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: isgreater(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:48: isgreaterequal(1.0, 1.0+eps) failed: got %d want %d\n\x00src/math/isless.c:48: isgreaterequal(1.0, 1.0+eps) raised the invalid exception\n\x00src/math/isless.c:49: isunordered(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: isunordered(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:49: isless(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: isless(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:49: islessequal(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: islessequal(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:49: islessgreater(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: islessgreater(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:49: isgreater(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: isgreater(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:49: isgreaterequal(1.0+eps, 1.0) failed: got %d want %d\n\x00src/math/isless.c:49: isgreaterequal(1.0+eps, 1.0) raised the invalid exception\n\x00src/math/isless.c:50: isunordered(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: isunordered(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:50: isless(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: isless(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:50: islessequal(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: islessequal(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:50: islessgreater(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: islessgreater(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:50: isgreater(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: isgreater(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:50: isgreaterequal(huge-1, huge) failed: got %d want %d\n\x00src/math/isless.c:50: isgreaterequal(huge-1, huge) raised the invalid exception\n\x00src/math/isless.c:51: isunordered(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: isunordered(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:51: isless(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: isless(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:51: islessequal(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: islessequal(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:51: islessgreater(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: islessgreater(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:51: isgreater(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: isgreater(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:51: isgreaterequal(huge, huge*huge) failed: got %d want %d\n\x00src/math/isless.c:51: isgreaterequal(huge, huge*huge) raised the invalid exception\n\x00src/math/isless.c:52: isunordered(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: isunordered(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:52: isless(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: isless(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:52: islessequal(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: islessequal(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:52: islessgreater(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: islessgreater(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:52: isgreater(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: isgreater(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:52: isgreaterequal(-0.0, 0.0) failed: got %d want %d\n\x00src/math/isless.c:52: isgreaterequal(-0.0, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: isunordered(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: isunordered(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: isless(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: isless(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: islessequal(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: islessequal(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: islessgreater(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: islessgreater(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: isgreater(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: isgreater(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:53: isgreaterequal(-tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:53: isgreaterequal(-tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:54: isunordered(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: isunordered(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:54: isless(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: isless(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:54: islessequal(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: islessequal(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:54: islessgreater(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: islessgreater(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:54: isgreater(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: isgreater(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:54: isgreaterequal(tiny, 2*tiny) failed: got %d want %d\n\x00src/math/isless.c:54: isgreaterequal(tiny, 2*tiny) raised the invalid exception\n\x00src/math/isless.c:55: isunordered(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: isunordered(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:55: isless(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: isless(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:55: islessequal(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: islessequal(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:55: islessgreater(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: islessgreater(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:55: isgreater(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: isgreater(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:55: isgreaterequal(tiny*0x1p-9, tiny*0x1p-8) failed: got %d want %d\n\x00src/math/isless.c:55: isgreaterequal(tiny*0x1p-9, tiny*0x1p-8) raised the invalid exception\n\x00src/math/isless.c:57: isunordered(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: isunordered(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:57: isless(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: isless(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:57: islessequal(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: islessequal(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:57: islessgreater(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: islessgreater(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:57: isgreater(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: isgreater(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:57: isgreaterequal(1.0f, 1.1f) failed: got %d want %d\n\x00src/math/isless.c:57: isgreaterequal(1.0f, 1.1f) raised the invalid exception\n\x00src/math/isless.c:58: isunordered(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: isunordered(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:58: isless(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: isless(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:58: islessequal(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: islessequal(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:58: islessgreater(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: islessgreater(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:58: isgreater(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: isgreater(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:58: isgreaterequal(1.0f, 1.0f+epsf) failed: got %d want %d\n\x00src/math/isless.c:58: isgreaterequal(1.0f, 1.0f+epsf) raised the invalid exception\n\x00src/math/isless.c:59: isunordered(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: isunordered(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:59: isless(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: isless(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:59: islessequal(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: islessequal(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:59: islessgreater(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: islessgreater(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:59: isgreater(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: isgreater(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:59: isgreaterequal(1.0f+epsf, 1.0f) failed: got %d want %d\n\x00src/math/isless.c:59: isgreaterequal(1.0f+epsf, 1.0f) raised the invalid exception\n\x00src/math/isless.c:60: isunordered(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: isunordered(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:60: isless(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: isless(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:60: islessequal(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: islessequal(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:60: islessgreater(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: islessgreater(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:60: isgreater(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: isgreater(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:60: isgreaterequal(hugef-1, hugef) failed: got %d want %d\n\x00src/math/isless.c:60: isgreaterequal(hugef-1, hugef) raised the invalid exception\n\x00src/math/isless.c:61: isunordered(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: isunordered(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:61: isless(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: isless(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:61: islessequal(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: islessequal(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:61: islessgreater(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: islessgreater(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:61: isgreater(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: isgreater(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:61: isgreaterequal(hugef, hugef*hugef) failed: got %d want %d\n\x00src/math/isless.c:61: isgreaterequal(hugef, hugef*hugef) raised the invalid exception\n\x00src/math/isless.c:62: isunordered(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: isunordered(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:62: isless(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: isless(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:62: islessequal(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: islessequal(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:62: islessgreater(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: islessgreater(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:62: isgreater(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: isgreater(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:62: isgreaterequal(-0.0f, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:62: isgreaterequal(-0.0f, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: isunordered(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: isunordered(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: isless(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: isless(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: islessequal(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: islessequal(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: islessgreater(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: islessgreater(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: isgreater(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: isgreater(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:63: isgreaterequal(-tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:63: isgreaterequal(-tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:64: isunordered(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: isunordered(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:64: isless(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: isless(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:64: islessequal(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: islessequal(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:64: islessgreater(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: islessgreater(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:64: isgreater(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: isgreater(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:64: isgreaterequal(tinyf, 2*tinyf) failed: got %d want %d\n\x00src/math/isless.c:64: isgreaterequal(tinyf, 2*tinyf) raised the invalid exception\n\x00src/math/isless.c:65: isunordered(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: isunordered(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:65: isless(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: isless(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:65: islessequal(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: islessequal(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:65: islessgreater(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: islessgreater(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:65: isgreater(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: isgreater(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:65: isgreaterequal(tinyf*0x1p-9f, tinyf*0x1p-8f) failed: got %d want %d\n\x00src/math/isless.c:65: isgreaterequal(tinyf*0x1p-9f, tinyf*0x1p-8f) raised the invalid exception\n\x00src/math/isless.c:67: isunordered(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: isunordered(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:67: isless(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: isless(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:67: islessequal(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: islessequal(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:67: islessgreater(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: islessgreater(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:67: isgreater(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: isgreater(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:67: isgreaterequal(1.0L, 1.1L) failed: got %d want %d\n\x00src/math/isless.c:67: isgreaterequal(1.0L, 1.1L) raised the invalid exception\n\x00src/math/isless.c:68: isunordered(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: isunordered(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:68: isless(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: isless(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:68: islessequal(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: islessequal(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:68: islessgreater(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: islessgreater(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:68: isgreater(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: isgreater(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:68: isgreaterequal(1.0L, 1.0L+epsl) failed: got %d want %d\n\x00src/math/isless.c:68: isgreaterequal(1.0L, 1.0L+epsl) raised the invalid exception\n\x00src/math/isless.c:69: isunordered(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: isunordered(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:69: isless(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: isless(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:69: islessequal(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: islessequal(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:69: islessgreater(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: islessgreater(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:69: isgreater(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: isgreater(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:69: isgreaterequal(1.0L+epsl, 1.0L) failed: got %d want %d\n\x00src/math/isless.c:69: isgreaterequal(1.0L+epsl, 1.0L) raised the invalid exception\n\x00src/math/isless.c:70: isunordered(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: isunordered(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:70: isless(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: isless(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:70: islessequal(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: islessequal(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:70: islessgreater(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: islessgreater(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:70: isgreater(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: isgreater(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:70: isgreaterequal(hugel-1, hugel) failed: got %d want %d\n\x00src/math/isless.c:70: isgreaterequal(hugel-1, hugel) raised the invalid exception\n\x00src/math/isless.c:71: isunordered(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: isunordered(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:71: isless(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: isless(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:71: islessequal(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: islessequal(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:71: islessgreater(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: islessgreater(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:71: isgreater(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: isgreater(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:71: isgreaterequal(hugel, hugel*hugel) failed: got %d want %d\n\x00src/math/isless.c:71: isgreaterequal(hugel, hugel*hugel) raised the invalid exception\n\x00src/math/isless.c:72: isunordered(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: isunordered(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:72: isless(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: isless(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:72: islessequal(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: islessequal(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:72: islessgreater(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: islessgreater(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:72: isgreater(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: isgreater(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:72: isgreaterequal(-0.0L, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:72: isgreaterequal(-0.0L, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: isunordered(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: isunordered(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: isless(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: isless(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: islessequal(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: islessequal(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: islessgreater(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: islessgreater(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: isgreater(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: isgreater(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:73: isgreaterequal(-tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:73: isgreaterequal(-tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:74: isunordered(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: isunordered(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:74: isless(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: isless(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:74: islessequal(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: islessequal(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:74: islessgreater(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: islessgreater(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:74: isgreater(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: isgreater(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:74: isgreaterequal(tinyl, 2*tinyl) failed: got %d want %d\n\x00src/math/isless.c:74: isgreaterequal(tinyl, 2*tinyl) raised the invalid exception\n\x00src/math/isless.c:75: isunordered(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: isunordered(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:75: isless(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: isless(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:75: islessequal(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: islessequal(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:75: islessgreater(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: islessgreater(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:75: isgreater(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: isgreater(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:75: isgreaterequal(tinyl*0x1p-9L, tinyl*0x1p-8L) failed: got %d want %d\n\x00src/math/isless.c:75: isgreaterequal(tinyl*0x1p-9L, tinyl*0x1p-8L) raised the invalid exception\n\x00src/math/isless.c:78: isunordered(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: isunordered(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:78: isless(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: isless(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:78: islessequal(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: islessequal(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:78: islessgreater(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: islessgreater(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:78: isgreater(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: isgreater(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:78: isgreaterequal(huge*huge, huge*huge*2) failed: got %d want %d\n\x00src/math/isless.c:78: isgreaterequal(huge*huge, huge*huge*2) raised the invalid exception\n\x00src/math/isless.c:79: isunordered(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: isunordered(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:79: isless(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: isless(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:79: islessequal(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: islessequal(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:79: islessgreater(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: islessgreater(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:79: isgreater(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: isgreater(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:79: isgreaterequal(tiny*tiny*0.5, tiny*tiny) failed: got %d want %d\n\x00src/math/isless.c:79: isgreaterequal(tiny*tiny*0.5, tiny*tiny) raised the invalid exception\n\x00src/math/isless.c:80: isunordered(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: isunordered(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:80: isless(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: isless(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:80: islessequal(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: islessequal(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:80: islessgreater(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: islessgreater(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:80: isgreater(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: isgreater(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:80: isgreaterequal(-tiny*tiny, 0.0) failed: got %d want %d\n\x00src/math/isless.c:80: isgreaterequal(-tiny*tiny, 0.0) raised the invalid exception\n\x00src/math/isless.c:81: isunordered(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: isunordered(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:81: isless(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: isless(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:81: islessequal(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: islessequal(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:81: islessgreater(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: islessgreater(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:81: isgreater(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: isgreater(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:81: isgreaterequal(1.0, 1.0+eps*0.25) failed: got %d want %d\n\x00src/math/isless.c:81: isgreaterequal(1.0, 1.0+eps*0.25) raised the invalid exception\n\x00src/math/isless.c:90: isunordered(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: isunordered(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:90: isless(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: isless(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:90: islessequal(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: islessequal(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:90: islessgreater(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: islessgreater(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:90: isgreater(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: isgreater(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:90: isgreaterequal(hugef*hugef, hugef*hugef*2) failed: got %d want %d\n\x00src/math/isless.c:90: isgreaterequal(hugef*hugef, hugef*hugef*2) raised the invalid exception\n\x00src/math/isless.c:91: isunordered(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: isunordered(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:91: isless(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: isless(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:91: islessequal(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: islessequal(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:91: islessgreater(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: islessgreater(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:91: isgreater(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: isgreater(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:91: isgreaterequal(tinyf*tinyf*0.5f, tinyf*tinyf) failed: got %d want %d\n\x00src/math/isless.c:91: isgreaterequal(tinyf*tinyf*0.5f, tinyf*tinyf) raised the invalid exception\n\x00src/math/isless.c:92: isunordered(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: isunordered(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:92: isless(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: isless(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:92: islessequal(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: islessequal(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:92: islessgreater(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: islessgreater(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:92: isgreater(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: isgreater(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:92: isgreaterequal(-tinyf*tinyf, 0.0f) failed: got %d want %d\n\x00src/math/isless.c:92: isgreaterequal(-tinyf*tinyf, 0.0f) raised the invalid exception\n\x00src/math/isless.c:93: isunordered(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: isunordered(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:93: isless(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: isless(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:93: islessequal(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: islessequal(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:93: islessgreater(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: islessgreater(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:93: isgreater(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: isgreater(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:93: isgreaterequal(1.0f, 1.0f+epsf*0.25f) failed: got %d want %d\n\x00src/math/isless.c:93: isgreaterequal(1.0f, 1.0f+epsf*0.25f) raised the invalid exception\n\x00src/math/isless.c:101: isunordered(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: isunordered(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:101: isless(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: isless(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:101: islessequal(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: islessequal(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:101: islessgreater(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: islessgreater(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:101: isgreater(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: isgreater(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:101: isgreaterequal(hugel*hugel, hugel*hugel*2) failed: got %d want %d\n\x00src/math/isless.c:101: isgreaterequal(hugel*hugel, hugel*hugel*2) raised the invalid exception\n\x00src/math/isless.c:102: isunordered(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: isunordered(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:102: isless(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: isless(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:102: islessequal(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: islessequal(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:102: islessgreater(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: islessgreater(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:102: isgreater(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: isgreater(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:102: isgreaterequal(tinyl*tinyl*0.5L, tinyl*tinyl) failed: got %d want %d\n\x00src/math/isless.c:102: isgreaterequal(tinyl*tinyl*0.5L, tinyl*tinyl) raised the invalid exception\n\x00src/math/isless.c:103: isunordered(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: isunordered(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:103: isless(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: isless(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:103: islessequal(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: islessequal(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:103: islessgreater(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: islessgreater(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:103: isgreater(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: isgreater(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:103: isgreaterequal(-tinyl*tinyl, 0.0L) failed: got %d want %d\n\x00src/math/isless.c:103: isgreaterequal(-tinyl*tinyl, 0.0L) raised the invalid exception\n\x00src/math/isless.c:104: isunordered(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: isunordered(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00src/math/isless.c:104: isless(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: isless(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00src/math/isless.c:104: islessequal(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: islessequal(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00src/math/isless.c:104: islessgreater(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: islessgreater(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00src/math/isless.c:104: isgreater(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: isgreater(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00src/math/isless.c:104: isgreaterequal(1.0L, 1.0L+epsl*0.25L) failed: got %d want %d\n\x00src/math/isless.c:104: isgreaterequal(1.0L, 1.0L+epsl*0.25L) raised the invalid exception\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
