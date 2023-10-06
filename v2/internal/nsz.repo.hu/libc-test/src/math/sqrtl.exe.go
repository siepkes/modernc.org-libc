// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/sqrtl.exe.go src/math/sqrtl.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:3:9:
const FE_DIVBYZERO = 4

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:11:9:
const FE_DOWNWARD = 1024

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:6:9:
const FE_INEXACT = 32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:1:9:
const FE_INVALID = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:4:9:
const FE_OVERFLOW = 8

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:10:9:
const FE_TONEAREST = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:13:9:
const FE_TOWARDZERO = 3072

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:5:9:
const FE_UNDERFLOW = 16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/fenv.h:12:9:
const FE_UPWARD = 2048

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

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:38:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:38:21:
type float_t = float64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:43:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:43:21:
type double_t = float64

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:5:19:
var t = [678]l_l{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942596e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942596e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942597e+154),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942596e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.457404589857209e+15),
		Fy:    libc.Float64FromFloat64(6.6763797e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.457404589857209e+15),
		Fy:    libc.Float64FromFloat64(6.6763797e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.457404589857209e+15),
		Fy:    libc.Float64FromFloat64(6.6763797e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.457404589857209e+15),
		Fy:    libc.Float64FromFloat64(6.6763797e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.911242719325041e+15),
		Fy:    libc.Float64FromFloat64(6.2539929e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.911242719325041e+15),
		Fy:    libc.Float64FromFloat64(6.2539929e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.911242719325041e+15),
		Fy:    libc.Float64FromFloat64(6.2539929e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.911242719325041e+15),
		Fy:    libc.Float64FromFloat64(6.2539929e+07),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(2.2227587494850775e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(2.2227587494850775e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(2.2227587494850775e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(2.2227587494850775e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9251665078627895),
		Fy:    libc.Float64FromFloat64(1.7103118159747333),
		Fdy:   float32(0.37662771344184875),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.848903589087301),
		Fy:    libc.Float64FromFloat64(2.617040998740238),
		Fdy:   float32(0.4624446928501129),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.5218011419322055),
		Fy:    libc.Float64FromFloat64(2.742590224939228),
		Fdy:   float32(-libc.Float64FromFloat64(0.12186521291732788)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.32857988648769),
		Fy:    libc.Float64FromFloat64(1.1526403977336948),
		Fdy:   float32(-libc.Float64FromFloat64(0.1205209344625473)),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.916389245089347),
		Fy:    libc.Float64FromFloat64(1.7077439050072312),
		Fdy:   float32(-libc.Float64FromFloat64(0.22862079739570618)),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.842830859393445),
		Fy:    libc.Float64FromFloat64(1.9603139695960556),
		Fdy:   float32(0.35526472330093384),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.31332709665502),
		Fy:    libc.Float64FromFloat64(2.5126334982752696),
		Fdy:   float32(-libc.Float64FromFloat64(0.38196054100990295)),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.420354043232288),
		Fy:    libc.Float64FromFloat64(2.7240326802797883),
		Fdy:   float32(-libc.Float64FromFloat64(0.1511644721031189)),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.3277226721946023),
		Fy:    libc.Float64FromFloat64(1.8242046683951345),
		Fdy:   float32(0.47830772399902344),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1679023732498908),
		Fy:    libc.Float64FromFloat64(0.4097589208911635),
		Fdy:   float32(-libc.Float64FromFloat64(0.14166373014450073)),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2e-323),
		Fy:    libc.Float64FromFloat64(4.445517498970155e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(7.458340731200207e-155),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(2.983336292480083e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.802596928649634e-45),
		Fy:    libc.Float64FromFloat64(5.293955920339377e-23),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1210387714598537e-44),
		Fy:    libc.Float64FromFloat64(1.0587911840678754e-22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.938735877055719e-39),
		Fy:    libc.Float64FromFloat64(5.421010862427522e-20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1754943508222875e-38),
		Fy:    libc.Float64FromFloat64(1.0842021724855044e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.70197740328915e-38),
		Fy:    libc.Float64FromFloat64(2.168404344971009e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.4332275390625e-05),
		Fy:    libc.Float64FromFloat64(0.005859375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.00390625),
		Fy:    libc.Float64FromFloat64(0.0625),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.03515625),
		Fy:    libc.Float64FromFloat64(0.1875),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0625),
		Fy:    libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(9),
		Fy:    libc.Float64FromFloat64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(256),
		Fy:    libc.Float64FromFloat64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2304),
		Fy:    libc.Float64FromFloat64(48),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(26569),
		Fy:    libc.Float64FromFloat64(163),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(27889),
		Fy:    libc.Float64FromFloat64(167),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(29929),
		Fy:    libc.Float64FromFloat64(173),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(32041),
		Fy:    libc.Float64FromFloat64(179),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(32761),
		Fy:    libc.Float64FromFloat64(181),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.359296e+06),
		Fy:    libc.Float64FromFloat64(1536),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.1267647932558654e+37),
		Fy:    libc.Float64FromFloat64(4.611686018427388e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.507059173023462e+37),
		Fy:    libc.Float64FromFloat64(9.223372036854776e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1235582092889474e+307),
		Fy:    libc.Float64FromFloat64(3.3519519824856493e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240041e-154),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.494232837155789e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971297e+153),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.4e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507197e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507198e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.4501477170144013e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014404e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(6.661338147750939e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.0000000000000018),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155788e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2e-323),
		Fy:    libc.Float64FromFloat64(4.445517498970155e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(7.458340731200207e-155),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(2.983336292480083e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.802596928649634e-45),
		Fy:    libc.Float64FromFloat64(5.293955920339377e-23),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1210387714598537e-44),
		Fy:    libc.Float64FromFloat64(1.0587911840678754e-22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.938735877055719e-39),
		Fy:    libc.Float64FromFloat64(5.421010862427522e-20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1754943508222875e-38),
		Fy:    libc.Float64FromFloat64(1.0842021724855044e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.70197740328915e-38),
		Fy:    libc.Float64FromFloat64(2.168404344971009e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4332275390625e-05),
		Fy:    libc.Float64FromFloat64(0.005859375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00390625),
		Fy:    libc.Float64FromFloat64(0.0625),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.03515625),
		Fy:    libc.Float64FromFloat64(0.1875),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0625),
		Fy:    libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9),
		Fy:    libc.Float64FromFloat64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(256),
		Fy:    libc.Float64FromFloat64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2304),
		Fy:    libc.Float64FromFloat64(48),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(26569),
		Fy:    libc.Float64FromFloat64(163),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(27889),
		Fy:    libc.Float64FromFloat64(167),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(29929),
		Fy:    libc.Float64FromFloat64(173),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(32041),
		Fy:    libc.Float64FromFloat64(179),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(32761),
		Fy:    libc.Float64FromFloat64(181),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.359296e+06),
		Fy:    libc.Float64FromFloat64(1536),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1267647932558654e+37),
		Fy:    libc.Float64FromFloat64(4.611686018427388e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.507059173023462e+37),
		Fy:    libc.Float64FromFloat64(9.223372036854776e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1235582092889474e+307),
		Fy:    libc.Float64FromFloat64(3.3519519824856493e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400412e-154),
		Fdy:   float32(5.551117770103743e-17),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400417e-154),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(8.881784197001252e-16),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(4.996003610813204e-16),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(2.220446313948109e-16),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(5.551117770103743e-17),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(4.44089183380283e-16),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(9.992007221626409e-16),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.494232837155789e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971298e+153),
		Fdy:   float32(5.551117770103743e-17),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    libc.Float64FromFloat64(6.7039039649713e+153),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.4e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507197e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507198e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.4501477170144013e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.450147717014404e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.661338147750939e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.0000000000000018),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155788e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2e-323),
		Fy:    libc.Float64FromFloat64(4.445517498970155e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(7.458340731200207e-155),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(2.983336292480083e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.802596928649634e-45),
		Fy:    libc.Float64FromFloat64(5.293955920339377e-23),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1210387714598537e-44),
		Fy:    libc.Float64FromFloat64(1.0587911840678754e-22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.938735877055719e-39),
		Fy:    libc.Float64FromFloat64(5.421010862427522e-20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1754943508222875e-38),
		Fy:    libc.Float64FromFloat64(1.0842021724855044e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.70197740328915e-38),
		Fy:    libc.Float64FromFloat64(2.168404344971009e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.4332275390625e-05),
		Fy:    libc.Float64FromFloat64(0.005859375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.00390625),
		Fy:    libc.Float64FromFloat64(0.0625),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.03515625),
		Fy:    libc.Float64FromFloat64(0.1875),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0625),
		Fy:    libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(9),
		Fy:    libc.Float64FromFloat64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(256),
		Fy:    libc.Float64FromFloat64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2304),
		Fy:    libc.Float64FromFloat64(48),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(26569),
		Fy:    libc.Float64FromFloat64(163),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(27889),
		Fy:    libc.Float64FromFloat64(167),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(29929),
		Fy:    libc.Float64FromFloat64(173),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(32041),
		Fy:    libc.Float64FromFloat64(179),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(32761),
		Fy:    libc.Float64FromFloat64(181),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.359296e+06),
		Fy:    libc.Float64FromFloat64(1536),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.1267647932558654e+37),
		Fy:    libc.Float64FromFloat64(4.611686018427388e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.507059173023462e+37),
		Fy:    libc.Float64FromFloat64(9.223372036854776e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1235582092889474e+307),
		Fy:    libc.Float64FromFloat64(3.3519519824856493e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400412e-154),
		Fdy:   float32(5.551115123125783e-17),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400417e-154),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(8.881784197001252e-16),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(4.996003610813204e-16),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(2.220446049250313e-16),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(5.551115123125783e-17),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(4.44089183380283e-16),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(9.992007221626409e-16),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.0000000000000009),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.494232837155789e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971298e+153),
		Fdy:   float32(5.551115123125783e-17),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    libc.Float64FromFloat64(6.7039039649713e+153),
		Fdy:   float32(1.1102227599273605e-16),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.4e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507197e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507198e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.4501477170144013e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014404e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(6.661338147750939e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.0000000000000018),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155788e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2e-323),
		Fy:    libc.Float64FromFloat64(4.445517498970155e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(7.458340731200207e-155),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(2.983336292480083e-154),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.802596928649634e-45),
		Fy:    libc.Float64FromFloat64(5.293955920339377e-23),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1210387714598537e-44),
		Fy:    libc.Float64FromFloat64(1.0587911840678754e-22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.938735877055719e-39),
		Fy:    libc.Float64FromFloat64(5.421010862427522e-20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1754943508222875e-38),
		Fy:    libc.Float64FromFloat64(1.0842021724855044e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.70197740328915e-38),
		Fy:    libc.Float64FromFloat64(2.168404344971009e-19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.4332275390625e-05),
		Fy:    libc.Float64FromFloat64(0.005859375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00390625),
		Fy:    libc.Float64FromFloat64(0.0625),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.03515625),
		Fy:    libc.Float64FromFloat64(0.1875),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0625),
		Fy:    libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9),
		Fy:    libc.Float64FromFloat64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(256),
		Fy:    libc.Float64FromFloat64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2304),
		Fy:    libc.Float64FromFloat64(48),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(26569),
		Fy:    libc.Float64FromFloat64(163),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(27889),
		Fy:    libc.Float64FromFloat64(167),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(29929),
		Fy:    libc.Float64FromFloat64(173),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(32041),
		Fy:    libc.Float64FromFloat64(179),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(32761),
		Fy:    libc.Float64FromFloat64(181),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.359296e+06),
		Fy:    libc.Float64FromFloat64(1536),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.1267647932558654e+37),
		Fy:    libc.Float64FromFloat64(4.611686018427388e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.507059173023462e+37),
		Fy:    libc.Float64FromFloat64(9.223372036854776e+18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1235582092889474e+307),
		Fy:    libc.Float64FromFloat64(3.3519519824856493e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240041e-154),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400413e-154),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.494232837155789e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971297e+153),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    libc.Float64FromFloat64(6.703903964971299e+153),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.4e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507197e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507198e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.4501477170144013e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.450147717014404e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.661338147750939e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.0000000000000018),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155788e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	342: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	343: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fy:    libc.Float64FromFloat64(2.0845238903256313),
		Fdy:   float32(-libc.Float64FromFloat64(0.07180261611938477)),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	345: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	346: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fy:    libc.Float64FromFloat64(3.0441841217266385),
		Fdy:   float32(-libc.Float64FromFloat64(0.01546262577176094)),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619858980995045),
		Fy:    libc.Float64FromFloat64(0.8136251582267503),
		Fdy:   float32(-libc.Float64FromFloat64(0.08618157356977463)),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	349: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5617597462207241),
		Fy:    libc.Float64FromFloat64(0.7495063350104014),
		Fdy:   float32(-libc.Float64FromFloat64(0.0981396734714508)),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7741522965913037),
		Fy:    libc.Float64FromFloat64(0.879859248170583),
		Fdy:   float32(-libc.Float64FromFloat64(0.37124353647232056)),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	352: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	353: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	354: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	355: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	356: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	357: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	358: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	359: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	360: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1e-323),
		Fy:    libc.Float64FromFloat64(3.1434555694052576e-162),
		Fdy:   float32(0.43537619709968567),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5e-323),
		Fy:    libc.Float64FromFloat64(3.849931087076416e-162),
		Fdy:   float32(-libc.Float64FromFloat64(0.45194002985954285)),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(2.2227587494850775e-162),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	363: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	364: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(14),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.414213562373095),
		Fdy:   float32(-libc.Float64FromFloat64(0.21107041835784912)),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(15),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(16),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0000000000000004),
		Fy:    libc.Float64FromFloat64(1.4142135623730951),
		Fdy:   float32(-libc.Float64FromFloat64(0.27173060178756714)),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(17),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(18),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	371: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942596e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(21),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.797693134862315e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942593e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(22),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623141e+308),
		Fy:    libc.Float64FromFloat64(1.340780792994259e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(23),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623133e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942587e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(24),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623125e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942584e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(25),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623117e+308),
		Fy:    libc.Float64FromFloat64(1.340780792994258e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(26),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.797693134862311e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942578e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(27),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623101e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942575e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623093e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942572e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(29),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623085e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942569e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(30),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623077e+308),
		Fy:    libc.Float64FromFloat64(1.3407807929942566e+154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(31),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507203e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400417e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(32),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507205e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400423e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(33),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507207e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240043e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(34),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507209e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400437e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507211e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400443e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(36),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072127e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240045e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(37),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072147e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400457e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072167e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400463e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072187e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240047e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(40),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072207e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400476e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072226e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400483e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072246e-308),
		Fy:    libc.Float64FromFloat64(1.491668146240049e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072266e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400496e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072286e-308),
		Fy:    libc.Float64FromFloat64(1.4916681462400503e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(92.35130391890645),
		Fy:    libc.Float64FromFloat64(9.609958580499006),
		Fdy:   float32(0.4998137056827545),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(93.3599596388916),
		Fy:    libc.Float64FromFloat64(9.662295774757238),
		Fdy:   float32(-libc.Float64FromFloat64(0.49979978799819946)),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(95.42049628886124),
		Fy:    libc.Float64FromFloat64(9.76834153215689),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997270107269287)),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(95.87916941885449),
		Fy:    libc.Float64FromFloat64(9.791790919890728),
		Fdy:   float32(0.4998766779899597),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(96.84804174884022),
		Fy:    libc.Float64FromFloat64(9.841140266698785),
		Fdy:   float32(0.499801903963089),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(97.43639050883155),
		Fy:    libc.Float64FromFloat64(9.87098731175517),
		Fdy:   float32(0.4997696280479431),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(97.50957979883047),
		Fy:    libc.Float64FromFloat64(9.874693909120955),
		Fdy:   float32(0.49999818205833435),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(97.80496893882612),
		Fy:    libc.Float64FromFloat64(9.88963947466368),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999580681324005)),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(98.2751822888192),
		Fy:    libc.Float64FromFloat64(9.913383997849534),
		Fdy:   float32(0.49979931116104126),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(99.47293564880155),
		Fy:    libc.Float64FromFloat64(9.973611966023219),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999540448188782)),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(100.57047130878539),
		Fy:    libc.Float64FromFloat64(10.028483001370914),
		Fdy:   float32(-libc.Float64FromFloat64(0.49996453523635864)),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(100.60954608878481),
		Fy:    libc.Float64FromFloat64(10.030431002144665),
		Fdy:   float32(0.49975672364234924),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(100.67909109878379),
		Fy:    libc.Float64FromFloat64(10.033897104255344),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997771382331848)),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(101.12268095877725),
		Fy:    libc.Float64FromFloat64(10.055977374615422),
		Fdy:   float32(0.49988678097724915),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(101.3027691287746),
		Fy:    libc.Float64FromFloat64(10.064927676281366),
		Fdy:   float32(0.4999105632305145),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.45932313565507e-307),
		Fy:    libc.Float64FromFloat64(4.9591563149945874e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998999834060669)),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.610957305180409e-307),
		Fy:    libc.Float64FromFloat64(7.490632353266584e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999343752861023)),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.8073887977408524e-307),
		Fy:    libc.Float64FromFloat64(7.62062254526548e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.49989569187164307)),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.026137080471427e-307),
		Fy:    libc.Float64FromFloat64(8.382205605013174e-154),
		Fdy:   float32(0.49980640411376953),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.438697769194972e-307),
		Fy:    libc.Float64FromFloat64(9.186238495268328e-154),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999065697193146)),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1607792515836795e-306),
		Fy:    libc.Float64FromFloat64(1.0773946591586944e-153),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997684359550476)),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2827413827423193e-306),
		Fy:    libc.Float64FromFloat64(1.1325817333606962e-153),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999513030052185)),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7116604596087457e-306),
		Fy:    libc.Float64FromFloat64(1.3083044216117078e-153),
		Fdy:   float32(-libc.Float64FromFloat64(0.49986395239830017)),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.038173251686994e-306),
		Fy:    libc.Float64FromFloat64(1.4276460526639628e-153),
		Fdy:   float32(0.4998403787612915),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.171572060856931e-306),
		Fy:    libc.Float64FromFloat64(1.4736254818836879e-153),
		Fdy:   float32(0.4999290406703949),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(70),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4681399631804094e-306),
		Fy:    libc.Float64FromFloat64(1.5710314965589996e-153),
		Fdy:   float32(0.49989044666290283),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(71),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5175533964200588e-306),
		Fy:    libc.Float64FromFloat64(1.5866799918131124e-153),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997701048851013)),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(72),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6461505468829625e-306),
		Fy:    libc.Float64FromFloat64(1.6266992797941982e-153),
		Fdy:   float32(0.4998672902584076),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8167076367720413e-306),
		Fy:    libc.Float64FromFloat64(1.9536395872248397e-153),
		Fdy:   float32(0.49983471632003784),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.5743220778562766e-306),
		Fy:    libc.Float64FromFloat64(2.1387664851161936e-153),
		Fdy:   float32(0.49985939264297485),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.507669773981944e+31),
		Fy:    libc.Float64FromFloat64(7.421367646183515e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.7081172777095287e+31),
		Fy:    libc.Float64FromFloat64(6.089431235927974e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.303704665253099e+31),
		Fy:    libc.Float64FromFloat64(7.282653819352599e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.878979092276793e+31),
		Fy:    libc.Float64FromFloat64(6.22814506275889e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(79),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.152316475392746e+31),
		Fy:    libc.Float64FromFloat64(7.843670362395877e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.211634565897406e+31),
		Fy:    libc.Float64FromFloat64(5.667128519715612e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.384153063706236e+31),
		Fy:    libc.Float64FromFloat64(7.990089526223243e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0478231792192007e+31),
		Fy:    libc.Float64FromFloat64(5.520709355888246e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	434: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.704599973929435e+31),
		Fy:    libc.Float64FromFloat64(8.777585074454953e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	435: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.240331294899048e+31),
		Fy:    libc.Float64FromFloat64(4.733213807656536e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.061213246847164e+31),
		Fy:    libc.Float64FromFloat64(6.372764899827362e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.386174466068633e+31),
		Fy:    libc.Float64FromFloat64(8.594285581750605e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	438: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(87),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4172103032625467e+31),
		Fy:    libc.Float64FromFloat64(4.916513300360884e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(88),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.298149066126225e+31),
		Fy:    libc.Float64FromFloat64(6.556027048545655e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.398490666089125e+31),
		Fy:    libc.Float64FromFloat64(7.347442184930158e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	441: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.798696577668996e+31),
		Fy:    libc.Float64FromFloat64(6.163356697181331e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	442: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.287154726749991e+31),
		Fy:    libc.Float64FromFloat64(7.929158042787388e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	443: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.115471445921065e+31),
		Fy:    libc.Float64FromFloat64(5.581640839324101e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	444: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.787540622643213e+31),
		Fy:    libc.Float64FromFloat64(6.91920560660197e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	445: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.175401572250135e+31),
		Fy:    libc.Float64FromFloat64(8.470774210336465e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	446: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5401848692100933e+31),
		Fy:    libc.Float64FromFloat64(5.040024671775024e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	447: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.749856479441335e+31),
		Fy:    libc.Float64FromFloat64(8.803326916252364e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	448: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.216029230934957e+31),
		Fy:    libc.Float64FromFloat64(4.707471965859125e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	449: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(98),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.079398540600521e+31),
		Fy:    libc.Float64FromFloat64(7.12698992604909e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	450: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.622062995668232e+31),
		Fy:    libc.Float64FromFloat64(8.730442712525083e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	451: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(100),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.285180510810281e+31),
		Fy:    libc.Float64FromFloat64(4.780356169586406e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	452: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.344520591848479e+31),
		Fy:    libc.Float64FromFloat64(6.59129774160482e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	453: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(102),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.658370290376833e+31),
		Fy:    libc.Float64FromFloat64(8.751211510629162e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	454: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.265367194677404e+31),
		Fy:    libc.Float64FromFloat64(4.759587371482327e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	455: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.0005785058422105e+31),
		Fy:    libc.Float64FromFloat64(7.071476865437807e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	456: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.673943956380633e+31),
		Fy:    libc.Float64FromFloat64(6.836624866394699e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	457: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.48346189049597e+31),
		Fy:    libc.Float64FromFloat64(6.695865806970724e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	458: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.77975503859052e+31),
		Fy:    libc.Float64FromFloat64(6.913577249579642e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	459: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.454161568144607e+31),
		Fy:    libc.Float64FromFloat64(7.385229561865093e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.752259949714389e+31),
		Fy:    libc.Float64FromFloat64(6.125569320246396e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(110),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.908261047786581e+31),
		Fy:    libc.Float64FromFloat64(8.311594941878834e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(111),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.703172161213076e+31),
		Fy:    libc.Float64FromFloat64(5.199203940232655e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(112),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.283689040604468e+31),
		Fy:    libc.Float64FromFloat64(6.544989717795184e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(113),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.848237713216801e+31),
		Fy:    libc.Float64FromFloat64(7.647377140704388e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	465: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(114),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4379714517605475e+31),
		Fy:    libc.Float64FromFloat64(5.863421741407101e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	466: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(115),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.767142266040733e+31),
		Fy:    libc.Float64FromFloat64(7.594170307571943e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.5006493689057854e+31),
		Fy:    libc.Float64FromFloat64(5.916628574539546e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(117),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.133544220083508e+31),
		Fy:    libc.Float64FromFloat64(7.164875588650168e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0270742446494973e+31),
		Fy:    libc.Float64FromFloat64(6.345923293461321e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(119),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.661532935280845e+31),
		Fy:    libc.Float64FromFloat64(8.161821448231299e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(120),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8611559588159497e+31),
		Fy:    libc.Float64FromFloat64(5.34897743388019e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(121),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.272657376678041e+31),
		Fy:    libc.Float64FromFloat64(7.261306615670517e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(122),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.905615358830551e+31),
		Fy:    libc.Float64FromFloat64(6.249492266440972e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(123),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.6932334597492285e+31),
		Fy:    libc.Float64FromFloat64(8.771107945835138e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	475: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(124),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.246467017142019e+31),
		Fy:    libc.Float64FromFloat64(4.739690936276351e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	476: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.3167516277557395e+31),
		Fy:    libc.Float64FromFloat64(6.570199104864129e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	477: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(126),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.886984611824975e+31),
		Fy:    libc.Float64FromFloat64(8.298785821928998e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	478: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.716508013951285e+31),
		Fy:    libc.Float64FromFloat64(5.212013060182491e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	479: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(128),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.735523649844082e+31),
		Fy:    libc.Float64FromFloat64(8.795182573343252e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	480: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(129),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.223703717152097e+31),
		Fy:    libc.Float64FromFloat64(4.715616308768237e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	481: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(130),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.186795308651041e+31),
		Fy:    libc.Float64FromFloat64(7.865618417296279e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	482: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1868062480331265e+31),
		Fy:    libc.Float64FromFloat64(5.64518046481521e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	483: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(132),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.127537651073882e+31),
		Fy:    libc.Float64FromFloat64(7.160682684684389e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	484: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0323975720826006e+31),
		Fy:    libc.Float64FromFloat64(6.3501161974271e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	485: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.435142396797359e+31),
		Fy:    libc.Float64FromFloat64(8.622727176941968e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	486: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3893244994878864e+31),
		Fy:    libc.Float64FromFloat64(4.888071705169521e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	487: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.261204807270303e+31),
		Fy:    libc.Float64FromFloat64(6.527790443381515e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	488: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.317954488117151e+31),
		Fy:    libc.Float64FromFloat64(6.571114432207943e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	489: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.255211542363269e+31),
		Fy:    libc.Float64FromFloat64(8.51775295624572e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	490: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(139),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.493050761780475e+31),
		Fy:    libc.Float64FromFloat64(4.993045925865769e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	491: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.907490741116252e+31),
		Fy:    libc.Float64FromFloat64(8.311131536148524e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.703654050867354e+31),
		Fy:    libc.Float64FromFloat64(5.199667345962965e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.3234971100141965e+31),
		Fy:    libc.Float64FromFloat64(7.296229923744315e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8620867338300856e+31),
		Fy:    libc.Float64FromFloat64(6.214568958367174e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(144),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.307442544293903e+31),
		Fy:    libc.Float64FromFloat64(8.548358055377596e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.462581895883536e+31),
		Fy:    libc.Float64FromFloat64(4.962440826733893e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.316369271993604e+31),
		Fy:    libc.Float64FromFloat64(6.569908121118289e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.768910627203775e+31),
		Fy:    libc.Float64FromFloat64(8.227338954488125e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7914948806801877e+31),
		Fy:    libc.Float64FromFloat64(5.283459927623364e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.969194414804752e+31),
		Fy:    libc.Float64FromFloat64(7.049251318264056e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.390722143384783e+31),
		Fy:    libc.Float64FromFloat64(6.626252442659262e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.4842430899827755e+31),
		Fy:    libc.Float64FromFloat64(6.696449126203212e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.136349434961683e+31),
		Fy:    libc.Float64FromFloat64(6.431445743347046e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.614251909634059e+31),
		Fy:    libc.Float64FromFloat64(8.725968089349204e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.289460571536615e+31),
		Fy:    libc.Float64FromFloat64(4.784830792762285e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.198445335931199e+31),
		Fy:    libc.Float64FromFloat64(7.210024504765015e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.969975775422584e+31),
		Fy:    libc.Float64FromFloat64(6.300774377346474e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.379373912472161e+31),
		Fy:    libc.Float64FromFloat64(8.590328231489273e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.421103142363461e+31),
		Fy:    libc.Float64FromFloat64(4.920470650622216e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.85766864234487e+31),
		Fy:    libc.Float64FromFloat64(8.864349182170607e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1589494814080708e+31),
		Fy:    libc.Float64FromFloat64(4.646449699940882e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.703311235145105e+31),
		Fy:    libc.Float64FromFloat64(8.7768509359252e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.241026315720138e+31),
		Fy:    libc.Float64FromFloat64(4.733947946186289e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.713995310168961e+31),
		Fy:    libc.Float64FromFloat64(8.782935335164982e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.23526937185456e+31),
		Fy:    libc.Float64FromFloat64(4.727863546946507e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.118209896894663e+31),
		Fy:    libc.Float64FromFloat64(8.436948439391261e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(166),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.574395831509225e+31),
		Fy:    libc.Float64FromFloat64(5.073850442720228e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.919505841940028e+31),
		Fy:    libc.Float64FromFloat64(8.318356713882874e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6961455670398675e+31),
		Fy:    libc.Float64FromFloat64(5.192442168228615e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.5310355331118335e+31),
		Fy:    libc.Float64FromFloat64(8.6781539126198e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3354457401153323e+31),
		Fy:    libc.Float64FromFloat64(4.832644969491689e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.5229906655387195e+31),
		Fy:    libc.Float64FromFloat64(6.725318331156318e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(172),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.313721081533377e+31),
		Fy:    libc.Float64FromFloat64(8.552029631340959e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.458939248238732e+31),
		Fy:    libc.Float64FromFloat64(4.95876925077053e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(174),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.976300383777964e+31),
		Fy:    libc.Float64FromFloat64(7.054289747223291e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.737091015954304e+31),
		Fy:    libc.Float64FromFloat64(7.574358729261709e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.5241321688367113e+31),
		Fy:    libc.Float64FromFloat64(5.93644015284978e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.175690201054387e+31),
		Fy:    libc.Float64FromFloat64(7.194226991869514e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.989908044459507e+31),
		Fy:    libc.Float64FromFloat64(6.316571890241975e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.826421288996464e+31),
		Fy:    libc.Float64FromFloat64(8.262215979382568e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.754762248681834e+31),
		Fy:    libc.Float64FromFloat64(5.248582902728921e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.331005065546803e+31),
		Fy:    libc.Float64FromFloat64(7.301373203409619e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8556967259322173e+31),
		Fy:    libc.Float64FromFloat64(6.20942567870187e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.858382793875266e+31),
		Fy:    libc.Float64FromFloat64(7.654007312431356e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.430200749067627e+31),
		Fy:    libc.Float64FromFloat64(5.856791569680133e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.220823443233025e+31),
		Fy:    libc.Float64FromFloat64(8.497542846748715e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5132736076101274e+31),
		Fy:    libc.Float64FromFloat64(5.013256035362774e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.7488837976548005e+31),
		Fy:    libc.Float64FromFloat64(7.582139406298727e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(188),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.514900318014445e+31),
		Fy:    libc.Float64FromFloat64(5.928659475812762e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(189),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.52016789985723e+31),
		Fy:    libc.Float64FromFloat64(6.723219392417022e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460664e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740989e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651693e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370498e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.946583574312755e+31),
		Fy:    libc.Float64FromFloat64(7.711409452436535e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.363291775702557e+31),
		Fy:    libc.Float64FromFloat64(5.799389429674952e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(194),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460662e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740988e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.02824096036517e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370499e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.480945779714492e+31),
		Fy:    libc.Float64FromFloat64(6.693986689346261e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.732811254779502e+31),
		Fy:    libc.Float64FromFloat64(7.571533038149871e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.5274878765249093e+31),
		Fy:    libc.Float64FromFloat64(5.939265843961616e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.669612804798942e+31),
		Fy:    libc.Float64FromFloat64(8.757632559544241e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.259259008998744e+31),
		Fy:    libc.Float64FromFloat64(4.753166322567246e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.11296384146066e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740987e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(202),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365171e+31),
		Fy:    libc.Float64FromFloat64(4.5035996273705e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.267237211790956e+31),
		Fy:    libc.Float64FromFloat64(6.532409365456941e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(204),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.009267739787761e+31),
		Fy:    libc.Float64FromFloat64(8.949451234454412e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(205),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0805892362786736e+31),
		Fy:    libc.Float64FromFloat64(4.561347647657075e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(206),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460658e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740986e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(207),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365172e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370501e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.987545969068988e+31),
		Fy:    libc.Float64FromFloat64(7.062255991585824e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(209),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.566164171128131e+31),
		Fy:    libc.Float64FromFloat64(6.757339839854238e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(210),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.734587595271395e+31),
		Fy:    libc.Float64FromFloat64(8.206453311432043e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(211),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.813608193318664e+31),
		Fy:    libc.Float64FromFloat64(5.304345570679444e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(212),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.1129638414606565e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740985e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365173e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370502e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.3443367859042575e+31),
		Fy:    libc.Float64FromFloat64(7.965134516067043e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.075439326081513e+31),
		Fy:    libc.Float64FromFloat64(5.545664366044444e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(216),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.180864436119667e+31),
		Fy:    libc.Float64FromFloat64(7.861847388572018e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(217),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1910652976361803e+31),
		Fy:    libc.Float64FromFloat64(5.648951493539469e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(218),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460655e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740984e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(219),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365174e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370503e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(220),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.054817174561428e+31),
		Fy:    libc.Float64FromFloat64(8.399295907730259e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(221),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6127462657108146e+31),
		Fy:    libc.Float64FromFloat64(5.111502974381228e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(222),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.101368825578754e+31),
		Fy:    libc.Float64FromFloat64(7.142386733843774e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(223),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0556673290203794e+31),
		Fy:    libc.Float64FromFloat64(6.368412148267713e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(224),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.871777744127253e+31),
		Fy:    libc.Float64FromFloat64(6.979812135098804e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.2409655399062555e+31),
		Fy:    libc.Float64FromFloat64(6.512269604297917e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460653e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740983e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(227),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651747e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370504e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(228),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.782986823264654e+31),
		Fy:    libc.Float64FromFloat64(6.915914128489923e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(229),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.760897398579499e+31),
		Fy:    libc.Float64FromFloat64(8.222467633611882e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(230),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7966447393857396e+31),
		Fy:    libc.Float64FromFloat64(5.288331248499605e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.526867208575129e+31),
		Fy:    libc.Float64FromFloat64(7.434290287966383e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(232),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6923956694719314e+31),
		Fy:    libc.Float64FromFloat64(6.076508594145104e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(233),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.6914391141316585e+31),
		Fy:    libc.Float64FromFloat64(8.180121706999021e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(234),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8416119145265026e+31),
		Fy:    libc.Float64FromFloat64(5.330677175112466e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.689375752172573e+31),
		Fy:    libc.Float64FromFloat64(6.847901687504408e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	587: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(236),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460651e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740982e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(237),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651756e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370505e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	589: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(238),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.543471967518235e+31),
		Fy:    libc.Float64FromFloat64(8.089172990805818e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	590: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(239),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9394027305275995e+31),
		Fy:    libc.Float64FromFloat64(5.421625891305669e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(240),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.783885350610401e+31),
		Fy:    libc.Float64FromFloat64(6.91656370650224e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(241),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.368141886728405e+31),
		Fy:    libc.Float64FromFloat64(7.98006383854691e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.058903012211327e+31),
		Fy:    libc.Float64FromFloat64(5.530735043564577e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(243),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.688490129351167e+31),
		Fy:    libc.Float64FromFloat64(7.542207985299243e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(244),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.5624077293509993e+31),
		Fy:    libc.Float64FromFloat64(5.968590896812244e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(245),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.183322690149958e+31),
		Fy:    libc.Float64FromFloat64(6.46786107623684e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	597: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(246),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460649e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740981e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	598: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(247),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651765e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370506e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(248),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.787592217770977e+31),
		Fy:    libc.Float64FromFloat64(6.919242890498191e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(249),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.152622205525523e+31),
		Fy:    libc.Float64FromFloat64(7.843865249687505e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	601: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(250),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2114136794298073e+31),
		Fy:    libc.Float64FromFloat64(5.666933632423982e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	602: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(251),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.221576081237102e+31),
		Fy:    libc.Float64FromFloat64(7.2260473851457e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	603: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(252),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.9498101378613707e+31),
		Fy:    libc.Float64FromFloat64(6.284751496965787e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	604: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(253),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.456488031009233e+31),
		Fy:    libc.Float64FromFloat64(7.386804472171463e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	605: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(254),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.750330753297667e+31),
		Fy:    libc.Float64FromFloat64(6.123994409940024e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	606: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(255),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.169685502135344e+31),
		Fy:    libc.Float64FromFloat64(7.190052504770284e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	607: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(256),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.9951834766671948e+31),
		Fy:    libc.Float64FromFloat64(6.320746377341203e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	608: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(257),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.319207980675688e+31),
		Fy:    libc.Float64FromFloat64(6.57206815293001e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	609: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(258),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.815857135430063e+31),
		Fy:    libc.Float64FromFloat64(6.939637696184191e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	610: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(259),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.1129638414606475e+31),
		Fy:    libc.Float64FromFloat64(9.00719925474098e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	611: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(260),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651774e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370507e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	612: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(261),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.350308431816392e+31),
		Fy:    libc.Float64FromFloat64(6.595686796548477e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	613: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(262),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.040348902906065e+31),
		Fy:    libc.Float64FromFloat64(8.966799263341443e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	614: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(263),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.064793253538231e+31),
		Fy:    libc.Float64FromFloat64(4.543999618770044e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	615: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(264),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.754470669887585e+31),
		Fy:    libc.Float64FromFloat64(7.585822743702613e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	616: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(265),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.510534224071454e+31),
		Fy:    libc.Float64FromFloat64(5.924976138408874e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	617: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(266),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.421297166841715e+31),
		Fy:    libc.Float64FromFloat64(6.649283545496999e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	618: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(267),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.600045152055993e+31),
		Fy:    libc.Float64FromFloat64(8.124066193757897e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	619: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(268),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.90168890557771e+31),
		Fy:    libc.Float64FromFloat64(5.38673268835359e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	620: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(269),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460646e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740979e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	621: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(270),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0282409603651783e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370508e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	622: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(271),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.26059965555696e+31),
		Fy:    libc.Float64FromFloat64(7.912395121299845e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	623: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1342124669069942e+31),
		Fy:    libc.Float64FromFloat64(5.598403760811642e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	624: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(273),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.6043458947675985e+31),
		Fy:    libc.Float64FromFloat64(8.126712677809889e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	625: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(274),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8988384255350794e+31),
		Fy:    libc.Float64FromFloat64(5.384086204301598e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	626: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(275),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.820280852686347e+31),
		Fy:    libc.Float64FromFloat64(8.25849916915074e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	627: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(276),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.758665227476755e+31),
		Fy:    libc.Float64FromFloat64(5.252299712960747e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	628: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(277),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.253952072616471e+31),
		Fy:    libc.Float64FromFloat64(8.517013603732514e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	629: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(278),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.493789140655456e+31),
		Fy:    libc.Float64FromFloat64(4.993785278378973e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	630: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(279),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.311590395646856e+31),
		Fy:    libc.Float64FromFloat64(7.944551841134184e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	631: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(280),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0983106121188587e+31),
		Fy:    libc.Float64FromFloat64(5.566247040977303e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	632: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(281),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.088966193091328e+31),
		Fy:    libc.Float64FromFloat64(7.803182807733859e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	633: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(282),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.257688125249389e+31),
		Fy:    libc.Float64FromFloat64(5.707616074377628e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	634: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(283),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.436050488258726e+31),
		Fy:    libc.Float64FromFloat64(8.022499914776394e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	635: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(284),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0121425554851454e+31),
		Fy:    libc.Float64FromFloat64(5.488298967335093e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	636: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(285),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460644e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740978e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	637: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(286),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365179e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370509e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	638: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(287),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.756864938196381e+31),
		Fy:    libc.Float64FromFloat64(8.807306590664583e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	639: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(288),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2122839735700452e+31),
		Fy:    libc.Float64FromFloat64(4.703492291446904e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	640: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(289),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.988506363748954e+31),
		Fy:    libc.Float64FromFloat64(8.359728682050006e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	641: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(290),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.653352420596143e+31),
		Fy:    libc.Float64FromFloat64(5.151070200061481e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	642: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(291),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.81285807372299e+31),
		Fy:    libc.Float64FromFloat64(8.254003921566181e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	643: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(292),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.763389325721453e+31),
		Fy:    libc.Float64FromFloat64(5.256794960545306e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	644: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(293),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.594389264201628e+31),
		Fy:    libc.Float64FromFloat64(6.778192431763521e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	645: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(294),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.555912141310338e+31),
		Fy:    libc.Float64FromFloat64(8.69247498777554e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	646: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(295),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.321624515072873e+31),
		Fy:    libc.Float64FromFloat64(4.818323894335947e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	647: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(296),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.512126496492307e+31),
		Fy:    libc.Float64FromFloat64(7.424369667852151e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	648: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(297),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.7044620580189524e+31),
		Fy:    libc.Float64FromFloat64(6.086429214259336e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	649: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(298),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.740887432351866e+31),
		Fy:    libc.Float64FromFloat64(7.576864412375257e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	650: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(299),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.52115782911238e+31),
		Fy:    libc.Float64FromFloat64(5.93393446973623e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	651: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(300),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.482789790355091e+31),
		Fy:    libc.Float64FromFloat64(8.051577355000131e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	652: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(301),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.980309968207605e+31),
		Fy:    libc.Float64FromFloat64(5.459221527111356e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	653: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(302),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.197678029390556e+31),
		Fy:    libc.Float64FromFloat64(6.478949011522282e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	654: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(303),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.526220660977826e+31),
		Fy:    libc.Float64FromFloat64(6.727719272515631e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	655: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(304),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.112963841460642e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740977e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	656: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(305),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.02824096036518e+31),
		Fy:    libc.Float64FromFloat64(4.50359962737051e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	657: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(306),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.509055950654503e+31),
		Fy:    libc.Float64FromFloat64(6.71495044706549e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	658: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(307),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.344705169301233e+31),
		Fy:    libc.Float64FromFloat64(6.591437756135783e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	659: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(308),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.896987142885946e+31),
		Fy:    libc.Float64FromFloat64(6.997847628296822e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	660: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(309),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.4714068413986925e+31),
		Fy:    libc.Float64FromFloat64(6.686857887976005e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	661: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(310),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.655187424265802e+31),
		Fy:    libc.Float64FromFloat64(8.157933209009376e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	662: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(311),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.865317091427492e+31),
		Fy:    libc.Float64FromFloat64(5.352865673102111e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	663: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(312),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.436582898922373e+31),
		Fy:    libc.Float64FromFloat64(6.660767897864609e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	664: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(313),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.473129231297679e+31),
		Fy:    libc.Float64FromFloat64(6.688145655783581e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	665: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(314),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.192786206097684e+31),
		Fy:    libc.Float64FromFloat64(8.481029540154711e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	666: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(315),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5298579633288304e+31),
		Fy:    libc.Float64FromFloat64(5.029769341956776e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	667: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(316),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.346289085711224e+31),
		Fy:    libc.Float64FromFloat64(8.571049577333702e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	668: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4401123194052615e+31),
		Fy:    libc.Float64FromFloat64(4.939749304777785e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	669: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(318),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.051568048423239e+31),
		Fy:    libc.Float64FromFloat64(8.973053019136373e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	670: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(319),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0591137516947767e+31),
		Fy:    libc.Float64FromFloat64(4.537745862975114e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	671: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(320),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.11296384146064e+31),
		Fy:    libc.Float64FromFloat64(9.007199254740976e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	672: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(321),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.028240960365181e+31),
		Fy:    libc.Float64FromFloat64(4.503599627370511e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	673: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(322),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.579015546492041e+31),
		Fy:    libc.Float64FromFloat64(7.469280786322094e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	674: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(323),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.64999409017507e+31),
		Fy:    libc.Float64FromFloat64(6.041518095789393e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	675: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(324),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.572825431706964e+31),
		Fy:    libc.Float64FromFloat64(6.762266951035698e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	676: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(325),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.538042043330116e+31),
		Fy:    libc.Float64FromFloat64(8.682189840892743e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	677: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(326),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3315465272939403e+31),
		Fy:    libc.Float64FromFloat64(4.828609041218744e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:18:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:19:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:23:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:26:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(24408)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:27:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:29:3:
		if (*l_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:30:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:31:3:
		libc.Xfesetround(tls, (*l_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:32:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:33:3:
		y = sqrtl(tls, (*l_l)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:34:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:36:3:
		if !(checkexceptall(tls, e, (*l_l)(unsafe.Pointer(p)).Fe, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:37:4:
			libc.Xprintf(tls, __ccgo_ts+68, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, estr(tls, (*l_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+120, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:40:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:42:3:
		d = ulperrl(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:43:3:
		if !(checkcr(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:44:4:
			libc.Xprintf(tls, __ccgo_ts+129, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*l_l)(unsafe.Pointer(p)).Fdy), float64((*l_l)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:46:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sqrtl.c:49:2:
	return libc.BoolInt32(!!(err != 0))
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

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:13:
type pid_t = int32

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
		fd = libc.Xopen(tls, __ccgo_ts+190, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+200, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:34:9:
const FP_INFINITE = 1

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
		Fs:    __ccgo_ts + 244,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 252,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 260,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 270,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 280,
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
				v2 = __ccgo_ts + 289
			} else {
				v2 = __ccgo_ts + 20
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+291, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 289
		} else {
			v3 = __ccgo_ts + 20
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+296, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 20
	} else {
		v4 = __ccgo_ts + 301
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+303, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 306
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 309
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 312
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 315
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 318
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+321, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+329, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+334, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+377, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+426) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+434) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+446) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+458) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+470) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+479) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+20) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+479) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+485, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/ucb/sqrt.h\x00\x00src/math/sanity/sqrt.h\x00src/math/special/sqrt.h\x00%s:%d: bad fp exception: %s sqrtl(%La)=%La, want %s\x00 got %s\n\x00%s:%d: %s sqrtl(%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
