// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/expm1l.exe.go src/math/expm1l.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:5:19:
var t = [751]l_l{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(15),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(16),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(17),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(18),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(27),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(29),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(30),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(31),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(32),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(33),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(34),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(709.7827128933839),
		Fy:    libc.Float64FromFloat64(1.7976931348620688e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(709.782712893384),
		Fy:    libc.Float64FromFloat64(1.7976931348622732e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(709.7827128933841),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(709.7827128933842),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933839),
		Fy:    libc.Float64FromFloat64(1.797693134862069e+308),
		Fdy:   float32(0.8943153619766235),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(709.782712893384),
		Fy:    libc.Float64FromFloat64(1.7976931348622734e+308),
		Fdy:   float32(0.8943153619766235),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933841),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933842),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933839),
		Fy:    libc.Float64FromFloat64(1.7976931348620688e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(709.782712893384),
		Fy:    libc.Float64FromFloat64(1.7976931348622732e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933841),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(709.7827128933842),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(709.7827128933839),
		Fy:    libc.Float64FromFloat64(1.7976931348620688e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(709.782712893384),
		Fy:    libc.Float64FromFloat64(1.7976931348622732e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.10568465292453766)),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(709.7827128933841),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(709.7827128933842),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(36.04365338911715),
		Fy:    -libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(5.0763438388273276e-15)),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(37.42994775023704),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(37.42994775023705),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(37.429947750237055),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.4178516392292583e+24),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(36.04365338911715),
		Fy:    -libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(36.625),
		Fy:    -libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(0.8817101716995239),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(37.42994775023704),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(37.42994775023705),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(37.429947750237055),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.4178516392292583e+24),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(36.04365338911715),
		Fy:    -libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(5.0763438388273276e-15)),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(36.625),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.11828982084989548)),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(37.42994775023704),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(37.42994775023705),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(37.429947750237055),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.4178516392292583e+24),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(36.04365338911715),
		Fy:    -libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(36.625),
		Fy:    -libc.Float64FromFloat64(0.9999999999999998),
		Fdy:   float32(0.8817101716995239),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(37.42994775023704),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(37.42994775023705),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(37.429947750237055),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.4178516392292583e+24),
		Fy:    -libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125781e-17),
		Fy:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(0.375),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125786e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125781e-17),
		Fy:    -libc.Float64FromFloat64(5.55111512312578e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125781e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.25)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fdy:   float32(-libc.Float64FromFloat64(0.125)),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125781e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fy:    -libc.Float64FromFloat64(5.5511151231257815e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125782e-17),
		Fdy:   float32(0.75),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125783e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.551115123125785e-17),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-17),
		Fdy:   float32(0.875),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.1588039009762205),
		Fy:    -libc.Float64FromFloat64(0.9843737627262741),
		Fdy:   float32(-libc.Float64FromFloat64(0.0312496405094862)),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6300653316969685),
		Fy:    libc.Float64FromFloat64(0.8777332507568573),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.837522455340574),
		Fy:    libc.Float64FromFloat64(1.3106351774748008),
		Fdy:   float32(3.634668669249446e-17),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.3864676908333609),
		Fy:    libc.Float64FromFloat64(0.47177284454410945),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.16334680970792906),
		Fy:    libc.Float64FromFloat64(0.17744496823242875),
		Fdy:   float32(3.2604389252406554e-17),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11941886519949067),
		Fy:    libc.Float64FromFloat64(0.12684181427238037),
		Fdy:   float32(-libc.Float64FromFloat64(6.5893075247076176e-18)),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.6300653316969685),
		Fy:    libc.Float64FromFloat64(0.8777332507568573),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.837522455340574),
		Fy:    libc.Float64FromFloat64(1.3106351774748006),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.3864676908333609),
		Fy:    libc.Float64FromFloat64(0.47177284454410945),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.16334680970792906),
		Fy:    libc.Float64FromFloat64(0.17744496823242872),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.11941886519949067),
		Fy:    libc.Float64FromFloat64(0.12684181427238037),
		Fdy:   float32(-libc.Float64FromFloat64(6.589281054928016e-18)),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.6300653316969685),
		Fy:    libc.Float64FromFloat64(0.8777332507568574),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.837522455340574),
		Fy:    libc.Float64FromFloat64(1.3106351774748008),
		Fdy:   float32(3.634666022271486e-17),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.3864676908333609),
		Fy:    libc.Float64FromFloat64(0.4717728445441095),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.16334680970792906),
		Fy:    libc.Float64FromFloat64(0.17744496823242875),
		Fdy:   float32(3.2604389252406554e-17),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.11941886519949067),
		Fy:    libc.Float64FromFloat64(0.1268418142723804),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.6300653316969685),
		Fy:    libc.Float64FromFloat64(0.8777332507568573),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.837522455340574),
		Fy:    libc.Float64FromFloat64(1.3106351774748006),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.3864676908333609),
		Fy:    libc.Float64FromFloat64(0.47177284454410945),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.16334680970792906),
		Fy:    libc.Float64FromFloat64(0.17744496823242872),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.11941886519949067),
		Fy:    libc.Float64FromFloat64(0.12684181427238037),
		Fdy:   float32(-libc.Float64FromFloat64(6.589281054928016e-18)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6645352591003745e-15),
		Fy:    libc.Float64FromFloat64(2.6645352591003777e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5006933289508785),
		Fy:    libc.Float64FromFloat64(0.6498647732549397),
		Fdy:   float32(-libc.Float64FromFloat64(1.1838172175495504e-16)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5139746479610767),
		Fy:    libc.Float64FromFloat64(0.6719233126339155),
		Fdy:   float32(-libc.Float64FromFloat64(4.538887457750844e-16)),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5370849970421203),
		Fy:    libc.Float64FromFloat64(0.7110119806899063),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.628493326460252),
		Fy:    libc.Float64FromFloat64(0.874783763165878),
		Fdy:   float32(-libc.Float64FromFloat64(8.321396728218464e-16)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.837522455340574),
		Fy:    libc.Float64FromFloat64(1.3106351774748006),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0012644990585318414),
		Fy:    libc.Float64FromFloat64(0.0012652988745530298),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0014220669368195187),
		Fy:    libc.Float64FromFloat64(0.0014230785534779516),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0015863115853247444),
		Fy:    libc.Float64FromFloat64(0.0015875704431065037),
		Fdy:   float32(-libc.Float64FromFloat64(4.486116246345717e-16)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0016155049649625313),
		Fy:    libc.Float64FromFloat64(0.0016168105960983046),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0016363372996054091),
		Fy:    libc.Float64FromFloat64(0.001637676830026004),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00179502442962638),
		Fy:    libc.Float64FromFloat64(0.0017966364503724328),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0008159650125383392),
		Fy:    libc.Float64FromFloat64(0.0008162980025524241),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0003072959030527372),
		Fy:    libc.Float64FromFloat64(0.0003073431232754902),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.000328182351347372),
		Fy:    libc.Float64FromFloat64(0.0003282362090667964),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00039213426389762384),
		Fy:    libc.Float64FromFloat64(0.000392211158588771),
		Fdy:   float32(-libc.Float64FromFloat64(7.296081345017262e-16)),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0004233823455238492),
		Fy:    libc.Float64FromFloat64(0.0004234719844791705),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0001406545188323459),
		Fy:    libc.Float64FromFloat64(0.00014066441114297384),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00021117761925352221),
		Fy:    libc.Float64FromFloat64(0.00021119991881665424),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00021679979544436472),
		Fy:    libc.Float64FromFloat64(0.00021682329821845193),
		Fdy:   float32(-libc.Float64FromFloat64(7.406167570193944e-16)),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0002297732788421618),
		Fy:    libc.Float64FromFloat64(0.00022979967874395537),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.64073763060803e-05),
		Fy:    libc.Float64FromFloat64(8.641110953094564e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.519273388116257e-05),
		Fy:    libc.Float64FromFloat64(8.519636288516983e-05),
		Fdy:   float32(-libc.Float64FromFloat64(5.619933903112234e-16)),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.973645553229528e-05),
		Fy:    libc.Float64FromFloat64(9.97414293779329e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.5250279590812134e-05),
		Fy:    libc.Float64FromFloat64(3.5250900889218046e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.707378358985801e-05),
		Fy:    libc.Float64FromFloat64(5.70754123292306e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.550665228249693e-05),
		Fy:    libc.Float64FromFloat64(1.550677251125088e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8095676534180272e-05),
		Fy:    libc.Float64FromFloat64(1.8095840261922472e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9374836215883414e-05),
		Fy:    libc.Float64FromFloat64(1.937502390923478e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8221945188971074e-05),
		Fy:    libc.Float64FromFloat64(1.8222111209622714e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.1599388293708596e-16)),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.014564325605376e-05),
		Fy:    libc.Float64FromFloat64(2.014584618088754e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.269404180270066e-05),
		Fy:    libc.Float64FromFloat64(2.2694299314415324e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.5616663475283292e-05),
		Fy:    libc.Float64FromFloat64(2.5616991584808776e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.958662057756793e-06),
		Fy:    libc.Float64FromFloat64(7.958693727991685e-06),
		Fdy:   float32(-libc.Float64FromFloat64(5.361931373147719e-16)),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.361520362501089e-05),
		Fy:    libc.Float64FromFloat64(1.3615296312316417e-05),
		Fdy:   float32(-libc.Float64FromFloat64(6.98297561606471e-16)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4182583203136826e-05),
		Fy:    libc.Float64FromFloat64(1.4182683776445444e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.979662307128111e-06),
		Fy:    libc.Float64FromFloat64(5.979680185344399e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.044123712576142e-06),
		Fy:    libc.Float64FromFloat64(7.044148522473836e-06),
		Fdy:   float32(-libc.Float64FromFloat64(4.512274741339298e-16)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.568089560214602e-06),
		Fy:    libc.Float64FromFloat64(6.568111130162061e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6801863664118377e-06),
		Fy:    libc.Float64FromFloat64(2.6801899581145256e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.894103088933567e-06),
		Fy:    libc.Float64FromFloat64(2.8941072768539516e-06),
		Fdy:   float32(-libc.Float64FromFloat64(6.47506289715551e-16)),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0482536195206217e-06),
		Fy:    libc.Float64FromFloat64(3.0482582654504066e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.2780576512590657e-06),
		Fy:    libc.Float64FromFloat64(3.278063024095919e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1.8570453567743727e-16)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.453778296756362e-06),
		Fy:    libc.Float64FromFloat64(3.45378426105549e-06),
		Fdy:   float32(-libc.Float64FromFloat64(8.174405806925505e-16)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6957044736884278e-06),
		Fy:    libc.Float64FromFloat64(3.6957113028126186e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.647498315271071e-06),
		Fy:    libc.Float64FromFloat64(3.6475049674011384e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.40981933237688895),
		Fy:    libc.Float64FromFloat64(0.5065455765155363),
		Fdy:   float32(-libc.Float64FromFloat64(4.772160860820271e-16)),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4342466565055342),
		Fy:    libc.Float64FromFloat64(0.543799609375296),
		Fdy:   float32(-libc.Float64FromFloat64(2.081256301401566e-16)),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.056484317642688e-06),
		Fy:    libc.Float64FromFloat64(1.0564848757224412e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0203408080216896e-06),
		Fy:    libc.Float64FromFloat64(1.0203413285695488e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1600155306750473e-06),
		Fy:    libc.Float64FromFloat64(1.1600162034933231e-06),
		Fdy:   float32(-libc.Float64FromFloat64(7.2160282611722585e-16)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.565559426830293e-06),
		Fy:    libc.Float64FromFloat64(1.5655606523190918e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.815036196920969e-06),
		Fy:    libc.Float64FromFloat64(1.8150378441001634e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.746156293526846e-07),
		Fy:    libc.Float64FromFloat64(6.746158569058595e-07),
		Fdy:   float32(-libc.Float64FromFloat64(5.049115752979601e-16)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.881378892084897e-07),
		Fy:    libc.Float64FromFloat64(6.881381259754213e-07),
		Fdy:   float32(-libc.Float64FromFloat64(3.231889679753449e-16)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.041226320679659e-07),
		Fy:    libc.Float64FromFloat64(8.041229553746563e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.755646520473109e-07),
		Fy:    libc.Float64FromFloat64(8.755650353541528e-07),
		Fdy:   float32(-libc.Float64FromFloat64(2.349682895271026e-16)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.233873626761689e-07),
		Fy:    libc.Float64FromFloat64(9.233877889984108e-07),
		Fdy:   float32(-libc.Float64FromFloat64(7.88446431149799e-16)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.5274289378632815e-07),
		Fy:    libc.Float64FromFloat64(2.52742925725816e-07),
		Fdy:   float32(-libc.Float64FromFloat64(4.489475790772765e-16)),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.897363454728493e-07),
		Fy:    libc.Float64FromFloat64(2.897363874464283e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.615221474550486e-07),
		Fy:    libc.Float64FromFloat64(4.615222539564113e-07),
		Fdy:   float32(-libc.Float64FromFloat64(5.305791088194888e-16)),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.873620180334251e-07),
		Fy:    libc.Float64FromFloat64(3.8736209305810124e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4152820019176016e-07),
		Fy:    libc.Float64FromFloat64(3.415282585125226e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5890828523183288e-07),
		Fy:    libc.Float64FromFloat64(1.5890829785775512e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9515836569307815e-07),
		Fy:    libc.Float64FromFloat64(1.9515838473647323e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.167885232560243e-08),
		Fy:    libc.Float64FromFloat64(6.167885422774288e-08),
		Fdy:   float32(-libc.Float64FromFloat64(3.813165775108917e-16)),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.57286690024966e-08),
		Fy:    libc.Float64FromFloat64(6.572867116262562e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.452696448848515e-08),
		Fy:    libc.Float64FromFloat64(7.452696726561944e-08),
		Fdy:   float32(-libc.Float64FromFloat64(6.351485025315844e-16)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000008013398725e-07),
		Fy:    libc.Float64FromFloat64(1.000008063399528e-07),
		Fdy:   float32(-libc.Float64FromFloat64(3.3314599908899683e-17)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.980073398518289e-08),
		Fy:    libc.Float64FromFloat64(8.980073801726892e-08),
		Fdy:   float32(-libc.Float64FromFloat64(4.510055515017492e-16)),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.623200713013559e-08),
		Fy:    libc.Float64FromFloat64(8.623201084811521e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.376315367589579e-08),
		Fy:    libc.Float64FromFloat64(5.3763155121134165e-08),
		Fdy:   float32(-libc.Float64FromFloat64(8.058698459352607e-16)),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.490849474166314e-08),
		Fy:    libc.Float64FromFloat64(5.4908496249134555e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.7986998003419136e-08),
		Fy:    libc.Float64FromFloat64(5.798699968466514e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.900270111151338e-08),
		Fy:    libc.Float64FromFloat64(5.900270285217278e-08),
		Fdy:   float32(-libc.Float64FromFloat64(8.785555960237243e-16)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.5197845513775635e-08),
		Fy:    libc.Float64FromFloat64(2.5197845831241347e-08),
		Fdy:   float32(-libc.Float64FromFloat64(7.692573232462264e-17)),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1454881239003747e-08),
		Fy:    libc.Float64FromFloat64(2.1454881469159715e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1872492534950583e-08),
		Fy:    libc.Float64FromFloat64(1.1872492605428621e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2976658584245984e-08),
		Fy:    libc.Float64FromFloat64(1.297665866844282e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4186172727824836e-08),
		Fy:    libc.Float64FromFloat64(1.4186172828448585e-08),
		Fdy:   float32(-libc.Float64FromFloat64(4.537687847339295e-16)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.9701676228849325e-09),
		Fy:    libc.Float64FromFloat64(3.970167630766048e-09),
		Fdy:   float32(-libc.Float64FromFloat64(7.027128267231525e-16)),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.49448036243847e-09),
		Fy:    libc.Float64FromFloat64(6.494480383527607e-09),
		Fdy:   float32(-libc.Float64FromFloat64(6.721584424915585e-16)),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.750222245670706e-09),
		Fy:    libc.Float64FromFloat64(6.750222268453456e-09),
		Fdy:   float32(-libc.Float64FromFloat64(8.322260701824663e-16)),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.01626289935116e-09),
		Fy:    libc.Float64FromFloat64(6.01626291744887e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.369629108811366e-09),
		Fy:    libc.Float64FromFloat64(6.369629129097454e-09),
		Fdy:   float32(-libc.Float64FromFloat64(6.253197439698823e-16)),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.366615357536089e-09),
		Fy:    libc.Float64FromFloat64(7.366615384669599e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6088439757076376e-09),
		Fy:    libc.Float64FromFloat64(2.608843979110671e-09),
		Fdy:   float32(-libc.Float64FromFloat64(8.657949917337791e-17)),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.21987045952630488),
		Fy:    libc.Float64FromFloat64(0.24591532367202992),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.24896832141868444),
		Fy:    libc.Float64FromFloat64(0.2827013982656397),
		Fdy:   float32(-libc.Float64FromFloat64(2.275991028859902e-16)),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8311983573747601e-09),
		Fy:    libc.Float64FromFloat64(1.831198359051404e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.009957332081318e-10),
		Fy:    libc.Float64FromFloat64(7.009957334538292e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.258646011715036e-10),
		Fy:    libc.Float64FromFloat64(9.258646016001162e-10),
		Fdy:   float32(-libc.Float64FromFloat64(4.4095397033536e-17)),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.017672645180204e-10),
		Fy:    libc.Float64FromFloat64(9.017672649246125e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.397582903371179e-10),
		Fy:    libc.Float64FromFloat64(2.3975829036585996e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.873299824013314e-10),
		Fy:    libc.Float64FromFloat64(2.873299824426107e-10),
		Fdy:   float32(-libc.Float64FromFloat64(4.575526926675513e-16)),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.989122517419991e-10),
		Fy:    libc.Float64FromFloat64(2.989122517866733e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.87588692829632e-10),
		Fy:    libc.Float64FromFloat64(3.875886929047445e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.8892452649521795e-17)),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1732481652641356e-10),
		Fy:    libc.Float64FromFloat64(1.173248165332961e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1664269550011046e-10),
		Fy:    libc.Float64FromFloat64(1.166426955069132e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.200533006316244e-10),
		Fy:    libc.Float64FromFloat64(1.2005330063883077e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2073542165792672e-10),
		Fy:    libc.Float64FromFloat64(1.207354216652152e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.275566319209414e-10),
		Fy:    libc.Float64FromFloat64(1.2755663192907672e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.24146026789436e-10),
		Fy:    libc.Float64FromFloat64(1.241460267971421e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3096723705244293e-10),
		Fy:    libc.Float64FromFloat64(1.309672370610191e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2687451089464063e-10),
		Fy:    libc.Float64FromFloat64(1.2687451090268918e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3028511602614293e-10),
		Fy:    libc.Float64FromFloat64(1.3028511603463001e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2346390576313445e-10),
		Fy:    libc.Float64FromFloat64(1.234639057707561e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3221490851987375e-10),
		Fy:    libc.Float64FromFloat64(1.3221490852861413e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3437784218394058e-10),
		Fy:    libc.Float64FromFloat64(1.3437784219296926e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.371063262891359e-10),
		Fy:    libc.Float64FromFloat64(1.3710632629853495e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3778844731543435e-10),
		Fy:    libc.Float64FromFloat64(1.3778844732492715e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5415935194655049e-10),
		Fy:    libc.Float64FromFloat64(1.54159351958433e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5484147297284505e-10),
		Fy:    libc.Float64FromFloat64(1.5484147298483297e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6507328836724494e-10),
		Fy:    libc.Float64FromFloat64(1.650732883808695e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4119905244692425e-10),
		Fy:    libc.Float64FromFloat64(1.411990524568928e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4392753655211337e-10),
		Fy:    libc.Float64FromFloat64(1.4392753656247091e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5143086784137067e-10),
		Fy:    libc.Float64FromFloat64(1.514308678528363e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6780177247241235e-10),
		Fy:    libc.Float64FromFloat64(1.6780177248649105e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6098056220948917e-10),
		Fy:    libc.Float64FromFloat64(1.6098056222244652e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4733814168359629e-10),
		Fy:    libc.Float64FromFloat64(1.4733814169445052e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3369572115764136e-10),
		Fy:    libc.Float64FromFloat64(1.336957211665786e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4051693142062658e-10),
		Fy:    libc.Float64FromFloat64(1.4051693143049906e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5074874681507532e-10),
		Fy:    libc.Float64FromFloat64(1.507487468264379e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4460965757841026e-10),
		Fy:    libc.Float64FromFloat64(1.4460965758886622e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.480202627098924e-10),
		Fy:    libc.Float64FromFloat64(1.4802026272084738e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5756995707802177e-10),
		Fy:    libc.Float64FromFloat64(1.575699570904359e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6848389349870382e-10),
		Fy:    libc.Float64FromFloat64(1.684838935128972e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5825207810431556e-10),
		Fy:    libc.Float64FromFloat64(1.582520781168374e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.616626832357822e-10),
		Fy:    libc.Float64FromFloat64(1.6166268324884958e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.643911673409527e-10),
		Fy:    libc.Float64FromFloat64(1.643911673544649e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.897010610061101e-10),
		Fy:    libc.Float64FromFloat64(1.8970106102410336e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.7324864162785367e-16)),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.0542428272817695e-10),
		Fy:    libc.Float64FromFloat64(2.054242827492765e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2554385572465513e-10),
		Fy:    libc.Float64FromFloat64(2.2554385575009016e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.866240826378032e-11),
		Fy:    libc.Float64FromFloat64(5.866240826550095e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.93445292900966e-11),
		Fy:    libc.Float64FromFloat64(5.934452929185748e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.033130293470063e-11),
		Fy:    libc.Float64FromFloat64(6.033130293652055e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.207301339536018e-11),
		Fy:    libc.Float64FromFloat64(6.207301339728669e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.275513442167568e-11),
		Fy:    libc.Float64FromFloat64(6.275513442364477e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.889422365850825e-11),
		Fy:    libc.Float64FromFloat64(6.889422366088144e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.616573955325088e-11),
		Fy:    libc.Float64FromFloat64(6.616573955543982e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.57154339216408e-11),
		Fy:    libc.Float64FromFloat64(7.57154339245072e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.98081600795129e-11),
		Fy:    libc.Float64FromFloat64(7.980816008269755e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.95763446848222e-11),
		Fy:    libc.Float64FromFloat64(6.957634468724262e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.230482879007646e-11),
		Fy:    libc.Float64FromFloat64(7.230482879269044e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.548361852693615e-11),
		Fy:    libc.Float64FromFloat64(6.548361852908019e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.298694981638964e-11),
		Fy:    libc.Float64FromFloat64(7.298694981905318e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.639755494795321e-11),
		Fy:    libc.Float64FromFloat64(7.639755495087149e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.113789787869701e-11),
		Fy:    libc.Float64FromFloat64(7.113789788122732e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.276845957940435e-11),
		Fy:    libc.Float64FromFloat64(9.276845958370733e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.253664418475785e-11),
		Fy:    libc.Float64FromFloat64(8.253664418816399e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1323209036859265e-10),
		Fy:    libc.Float64FromFloat64(1.1323209037500339e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.617906471094543e-11),
		Fy:    libc.Float64FromFloat64(9.617906471557062e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.002717908687896e-10),
		Fy:    libc.Float64FromFloat64(1.002717908738168e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.93578544478594e-11),
		Fy:    libc.Float64FromFloat64(8.93578544518518e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.345058060571287e-11),
		Fy:    libc.Float64FromFloat64(9.345058061007937e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.321876521106871e-11),
		Fy:    libc.Float64FromFloat64(8.321876521453138e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.662937034262064e-11),
		Fy:    libc.Float64FromFloat64(8.662937034637295e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.00399754741687e-11),
		Fy:    libc.Float64FromFloat64(9.003997547822228e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.686118573725318e-11),
		Fy:    libc.Float64FromFloat64(9.686118574194421e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0368239600032215e-10),
		Fy:    libc.Float64FromFloat64(1.0368239600569715e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0641088010554539e-10),
		Fy:    libc.Float64FromFloat64(1.0641088011120702e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0709300113185081e-10),
		Fy:    libc.Float64FromFloat64(1.0709300113758526e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.958966984248262e-11),
		Fy:    libc.Float64FromFloat64(9.958966984744166e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0300027497401595e-10),
		Fy:    libc.Float64FromFloat64(1.0300027497932046e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.105036062633756e-10),
		Fy:    libc.Float64FromFloat64(1.1050360626948112e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0937707793617611e-10),
		Fy:    libc.Float64FromFloat64(1.0937707794215777e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0982148523707096e-10),
		Fy:    libc.Float64FromFloat64(1.0982148524310132e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1033300279268596e-10),
		Fy:    libc.Float64FromFloat64(1.1033300279877265e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.912603905320127e-11),
		Fy:    libc.Float64FromFloat64(7.912603905633172e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.594724931631056e-11),
		Fy:    libc.Float64FromFloat64(8.594724932000402e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1391421139489652e-10),
		Fy:    libc.Float64FromFloat64(1.1391421140138473e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.583335304535564e-11),
		Fy:    libc.Float64FromFloat64(4.5833353046405984e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.466560206605262e-11),
		Fy:    libc.Float64FromFloat64(1.4665602066160158e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.637090463186804e-11),
		Fy:    libc.Float64FromFloat64(1.637090463200204e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.807620719768249e-11),
		Fy:    libc.Float64FromFloat64(1.8076207197845862e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0518899538323396e-11),
		Fy:    libc.Float64FromFloat64(2.051889953853391e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.978150976349597e-11),
		Fy:    libc.Float64FromFloat64(1.9781509763691622e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.1486812329308482e-11),
		Fy:    libc.Float64FromFloat64(2.148681232953932e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.306042995726106e-11),
		Fy:    libc.Float64FromFloat64(2.3060429957526953e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.4897417460930597e-11),
		Fy:    libc.Float64FromFloat64(2.4897417461240534e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.3192114895120024e-11),
		Fy:    libc.Float64FromFloat64(2.3192114895388958e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.66027200267402e-11),
		Fy:    libc.Float64FromFloat64(2.660272002709405e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.8308022592548834e-11),
		Fy:    libc.Float64FromFloat64(2.8308022592949503e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2165799458467402e-11),
		Fy:    libc.Float64FromFloat64(2.2165799458713065e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.673861546199267e-12),
		Fy:    libc.Float64FromFloat64(7.67386154622871e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.014922059363224e-12),
		Fy:    libc.Float64FromFloat64(8.014922059395341e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.526512829109085e-12),
		Fy:    libc.Float64FromFloat64(8.526512829145434e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.720224625182423e-12),
		Fy:    libc.Float64FromFloat64(9.720224625229663e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.867573342272945e-12),
		Fy:    libc.Float64FromFloat64(8.86757334231226e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.131466648584592e-12),
		Fy:    libc.Float64FromFloat64(9.131466648626283e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.057287590809166e-11),
		Fy:    libc.Float64FromFloat64(1.0572875908147551e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1084466677837085e-11),
		Fy:    libc.Float64FromFloat64(1.1084466677898516e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1937117960745934e-11),
		Fy:    libc.Float64FromFloat64(1.193711796081718e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2274822674619144e-11),
		Fy:    libc.Float64FromFloat64(1.227482267469448e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2327291909146816e-11),
		Fy:    libc.Float64FromFloat64(1.2327291909222796e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0937521680111614e-11),
		Fy:    libc.Float64FromFloat64(1.0937521680171428e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0231815394927994e-11),
		Fy:    libc.Float64FromFloat64(1.0231815394980338e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1425527191000654e-11),
		Fy:    libc.Float64FromFloat64(1.1425527191065924e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.278976924365454e-11),
		Fy:    libc.Float64FromFloat64(1.2789769243736328e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.17110549218567e-11),
		Fy:    libc.Float64FromFloat64(1.1711054921925274e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2278178473909406e-11),
		Fy:    libc.Float64FromFloat64(1.2278178473984781e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4495071809471026e-11),
		Fy:    libc.Float64FromFloat64(1.4495071809576078e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3130829756817915e-11),
		Fy:    libc.Float64FromFloat64(1.3130829756904123e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.379164112018661e-12),
		Fy:    libc.Float64FromFloat64(9.379164112062644e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3642420526562904e-11),
		Fy:    libc.Float64FromFloat64(1.364242052665596e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3983481039726182e-11),
		Fy:    libc.Float64FromFloat64(1.398348103982395e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4551804205605577e-11),
		Fy:    libc.Float64FromFloat64(1.4551804205711455e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.7678561791877776e-16)),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.990493025328017e-12),
		Fy:    libc.Float64FromFloat64(3.990493025335978e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.433786671139749e-12),
		Fy:    libc.Float64FromFloat64(4.433786671149577e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.66453853298107e-12),
		Fy:    libc.Float64FromFloat64(4.6645385329919495e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.5657333242957705e-12),
		Fy:    libc.Float64FromFloat64(4.565733324306193e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.840877128896665e-12),
		Fy:    libc.Float64FromFloat64(4.840877128908382e-12),
		Fdy:   float32(-libc.Float64FromFloat64(7.285173678238995e-16)),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.50714167885154e-12),
		Fy:    libc.Float64FromFloat64(5.507141678866704e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.9917405198712384e-12),
		Fy:    libc.Float64FromFloat64(6.99174051989568e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.286437954050488e-12),
		Fy:    libc.Float64FromFloat64(5.28643795406446e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.139089236960984e-12),
		Fy:    libc.Float64FromFloat64(6.139089236979828e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.364242079004326e-12),
		Fy:    libc.Float64FromFloat64(6.3642420790245785e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.875832822406078e-12),
		Fy:    libc.Float64FromFloat64(1.875832822407837e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.728484105317544e-12),
		Fy:    libc.Float64FromFloat64(2.728484105321266e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.155182825177015e-12),
		Fy:    libc.Float64FromFloat64(2.1551828251793376e-12),
		Fdy:   float32(-libc.Float64FromFloat64(4.123998279389275e-16)),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.3021584638618413e-12),
		Fy:    libc.Float64FromFloat64(2.302158463864491e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.154809746773186e-12),
		Fy:    libc.Float64FromFloat64(3.154809746778162e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.369621102881549e-12),
		Fy:    libc.Float64FromFloat64(3.3696211028872256e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.5811353882287675e-12),
		Fy:    libc.Float64FromFloat64(3.5811353882351794e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11339152528017672),
		Fy:    libc.Float64FromFloat64(0.12007038298136355),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.379164112031856e-13),
		Fy:    libc.Float64FromFloat64(9.379164112036253e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.303608728285667e-12),
		Fy:    libc.Float64FromFloat64(1.3036087282865168e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3063942201973668e-12),
		Fy:    libc.Float64FromFloat64(1.30639422019822e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1084466677855515e-12),
		Fy:    libc.Float64FromFloat64(1.1084466677861656e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5347723092414238e-12),
		Fy:    libc.Float64FromFloat64(1.5347723092426014e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7789129541690107e-12),
		Fy:    libc.Float64FromFloat64(1.7789129541705928e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3642420526590822e-12),
		Fy:    libc.Float64FromFloat64(1.3642420526600125e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.487055778884721e-12),
		Fy:    libc.Float64FromFloat64(1.4870557788858266e-12),
		Fdy:   float32(-libc.Float64FromFloat64(2.9086034088682885e-16)),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7905676941149181e-12),
		Fy:    libc.Float64FromFloat64(1.790567694116521e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.791686098835062e-13),
		Fy:    libc.Float64FromFloat64(5.791686098836738e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.175256585893963e-13),
		Fy:    libc.Float64FromFloat64(8.175256585897305e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.189484742661675e-13),
		Fy:    libc.Float64FromFloat64(7.18948474266426e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.69072485523841e-13),
		Fy:    libc.Float64FromFloat64(8.690724855242186e-13),
		Fdy:   float32(-libc.Float64FromFloat64(4.667562879746205e-16)),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4909565102799116e-13),
		Fy:    libc.Float64FromFloat64(2.490956510280222e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.7205529873804976e-13),
		Fy:    libc.Float64FromFloat64(2.7205529873808677e-13),
		Fdy:   float32(-libc.Float64FromFloat64(6.590203262049339e-16)),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.410605131648287e-13),
		Fy:    libc.Float64FromFloat64(3.410605131648868e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.526325597726108e-13),
		Fy:    libc.Float64FromFloat64(3.5263255977267293e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.263256414560298e-13),
		Fy:    libc.Float64FromFloat64(4.2632564145612065e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2337690232431684e-13),
		Fy:    libc.Float64FromFloat64(1.2337690232432446e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.776586443654602e-14),
		Fy:    libc.Float64FromFloat64(3.7765864436546726e-14),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.275510661177438e-14),
		Fy:    libc.Float64FromFloat64(5.275510661177577e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5383701491068472e-14),
		Fy:    libc.Float64FromFloat64(1.538370149106859e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8028052943983022e-14),
		Fy:    libc.Float64FromFloat64(1.8028052943983183e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7222418539888605e-14),
		Fy:    libc.Float64FromFloat64(1.7222418539888753e-14),
		Fdy:   float32(-libc.Float64FromFloat64(5.972666480315303e-16)),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.723103470137968e-14),
		Fy:    libc.Float64FromFloat64(2.723103470138005e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7461810307976428e-14),
		Fy:    libc.Float64FromFloat64(2.7461810307976806e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.440892098500623e-15),
		Fy:    libc.Float64FromFloat64(4.440892098500633e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.699798436761761e-15),
		Fy:    libc.Float64FromFloat64(4.6997984367617716e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.958081967793449e-15),
		Fy:    libc.Float64FromFloat64(5.9580819677934665e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.165926057296533e-15),
		Fy:    libc.Float64FromFloat64(4.165926057296541e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.4750994875343035e-15),
		Fy:    libc.Float64FromFloat64(5.475099487534318e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.764165321960913e-15),
		Fy:    libc.Float64FromFloat64(6.764165321960936e-15),
		Fdy:   float32(-libc.Float64FromFloat64(5.512279721285357e-16)),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.552713678800499e-15),
		Fy:    libc.Float64FromFloat64(3.552713678800506e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.044456616251767944),
		Fy:    libc.Float64FromFloat64(0.04545961976965475),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.04668417805888616),
		Fy:    libc.Float64FromFloat64(0.04779104141839956),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.05447064424404796),
		Fy:    libc.Float64FromFloat64(0.05598147682419537),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3322676295501877e-15),
		Fy:    libc.Float64FromFloat64(1.3322676295501886e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.021666937153402e-16),
		Fy:    libc.Float64FromFloat64(7.021666937153405e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.881784197001249e-16),
		Fy:    libc.Float64FromFloat64(8.881784197001252e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01744543912144125),
		Fy:    libc.Float64FromFloat64(0.01759849956793867),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.005240007166606304),
		Fy:    libc.Float64FromFloat64(0.005253760015341516),
		Fdy:   float32(-libc.Float64FromFloat64(7.651595896056469e-16)),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.006140304489335447),
		Fy:    libc.Float64FromFloat64(0.0061591948032472255),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0020367416504511963),
		Fy:    libc.Float64FromFloat64(0.0020388172176187013),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.003045637107548267),
		Fy:    libc.Float64FromFloat64(0.003050279772337934),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0033009818573028283),
		Fy:    libc.Float64FromFloat64(0.0033064360977122418),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0037731210197621975),
		Fy:    libc.Float64FromFloat64(0.0037802482019648758),
		Fdy:   float32(-libc.Float64FromFloat64(5.774820230264788e-16)),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5007991627201707),
		Fy:    libc.Float64FromFloat64(3.4852721006879204),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6270060846924657),
		Fy:    libc.Float64FromFloat64(4.088617001442459),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6744336219614115),
		Fy:    libc.Float64FromFloat64(4.335772228886831),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6486371958975217),
		Fy:    libc.Float64FromFloat64(13.134762601162658),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.7460971663977936),
		Fy:    libc.Float64FromFloat64(41.35545271444171),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.800446273003557),
		Fy:    libc.Float64FromFloat64(120.56465643950725),
		Fdy:   float32(-libc.Float64FromFloat64(1.3789127281038573e-16)),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.657914718791208),
		Fy:    libc.Float64FromFloat64(777.924964819056),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(11.022872793631722),
		Fy:    libc.Float64FromFloat64(61258.41271820104),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(11.411195701885317),
		Fy:    libc.Float64FromFloat64(90326.36165653409),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(11.794490387560606),
		Fy:    libc.Float64FromFloat64(132519.20290772576),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(16.55983372179868),
		Fy:    libc.Float64FromFloat64(1.5554081802898357e+07),
		Fdy:   float32(-libc.Float64FromFloat64(4.9277915181752236e-17)),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(33.36110430768029),
		Fy:    libc.Float64FromFloat64(3.0799488994606656e+14),
		Fdy:   float32(-libc.Float64FromFloat64(7.398695680807977e-16)),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5521286037396312),
		Fy:    -libc.Float64FromFloat64(0.42427698100389327),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5678971675654444),
		Fy:    -libc.Float64FromFloat64(0.43328410483884316),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5875302365511949),
		Fy:    -libc.Float64FromFloat64(0.4443019659825088),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.7504361301451186),
		Fy:    -libc.Float64FromFloat64(0.527839415634408),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0010342508382093867),
		Fy:    -libc.Float64FromFloat64(0.0010337161851488954),
		Fdy:   float32(5.962046804739102e-16),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0009993047459634164),
		Fy:    -libc.Float64FromFloat64(0.0009988056072535024),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0015016910926750034),
		Fy:    -libc.Float64FromFloat64(0.001500564118798892),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.001592280836930195),
		Fy:    -libc.Float64FromFloat64(0.0015910138303643474),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0007661371015089736),
		Fy:    -libc.Float64FromFloat64(0.0007658436934148785),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0002592159863363192),
		Fy:    -libc.Float64FromFloat64(0.00025918239277525844),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0002816961508676599),
		Fy:    -libc.Float64FromFloat64(0.00028165647823225005),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.000285703525987575),
		Fy:    -libc.Float64FromFloat64(0.0002856627166217466),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.000348414684871073),
		Fy:    -libc.Float64FromFloat64(0.0003483539955233141),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(448),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0003062755798036508),
		Fy:    -libc.Float64FromFloat64(0.00030622868222624183),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(449),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00031943569383397147),
		Fy:    -libc.Float64FromFloat64(0.0003193846796847818),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(450),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0003403726872763659),
		Fy:    -libc.Float64FromFloat64(0.0003403147670649164),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(451),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00037099172729430994),
		Fy:    -libc.Float64FromFloat64(0.0003709229183728928),
		Fdy:   float32(5.728346179426904e-16),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(452),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00047748723033089457),
		Fy:    -libc.Float64FromFloat64(0.00047737325144520603),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts,
		Fline: int32(453),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00022062165813775716),
		Fy:    -libc.Float64FromFloat64(0.00022059732296939206),
		Fdy:   float32(9.459052503027227e-17),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts,
		Fline: int32(454),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.629310138688924e-05),
		Fy:    -libc.Float64FromFloat64(7.629019114224062e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts,
		Fline: int32(455),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00012152133611286769),
		Fy:    -libc.Float64FromFloat64(0.00012151395269438635),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	434: {
		Ffile: __ccgo_ts,
		Fline: int32(456),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.981783153260638e-05),
		Fy:    -libc.Float64FromFloat64(5.981604248179431e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	435: {
		Ffile: __ccgo_ts,
		Fline: int32(457),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.792297759687184e-05),
		Fy:    -libc.Float64FromFloat64(4.792182930932415e-05),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts,
		Fline: int32(458),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.3662775599579165e-05),
		Fy:    -libc.Float64FromFloat64(5.366133577859168e-05),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts,
		Fline: int32(459),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.795512057379581e-05),
		Fy:    -libc.Float64FromFloat64(2.795472983305374e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	438: {
		Ffile: __ccgo_ts,
		Fline: int32(460),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.266651918724206e-05),
		Fy:    -libc.Float64FromFloat64(2.2666262303636913e-05),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts,
		Fline: int32(461),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.706013270055256e-06),
		Fy:    -libc.Float64FromFloat64(8.705975372831704e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts,
		Fline: int32(462),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.762046061947656e-06),
		Fy:    -libc.Float64FromFloat64(8.762007675334174e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	441: {
		Ffile: __ccgo_ts,
		Fline: int32(463),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.329032253041802e-06),
		Fy:    -libc.Float64FromFloat64(9.328988737755731e-06),
		Fdy:   float32(7.159378168264299e-16),
		Fe:    int32(FE_INEXACT),
	},
	442: {
		Ffile: __ccgo_ts,
		Fline: int32(464),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.085329616300062e-05),
		Fy:    -libc.Float64FromFloat64(1.0853237266194895e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	443: {
		Ffile: __ccgo_ts,
		Fline: int32(465),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1401820789458584e-05),
		Fy:    -libc.Float64FromFloat64(1.1401755788946967e-05),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	444: {
		Ffile: __ccgo_ts,
		Fline: int32(466),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1773710732438632e-05),
		Fy:    -libc.Float64FromFloat64(1.1773641422578438e-05),
		Fdy:   float32(8.5190158675599965e-16),
		Fe:    int32(FE_INEXACT),
	},
	445: {
		Ffile: __ccgo_ts,
		Fline: int32(467),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4450292537476033e-05),
		Fy:    -libc.Float64FromFloat64(1.4450188132501718e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	446: {
		Ffile: __ccgo_ts,
		Fline: int32(468),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4804043617729241e-05),
		Fy:    -libc.Float64FromFloat64(1.4803934038416263e-05),
		Fdy:   float32(7.728447724756444e-17),
		Fe:    int32(FE_INEXACT),
	},
	447: {
		Ffile: __ccgo_ts,
		Fline: int32(469),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.516178746541967e-06),
		Fy:    -libc.Float64FromFloat64(9.516133467856625e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	448: {
		Ffile: __ccgo_ts,
		Fline: int32(470),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.017322351301034e-05),
		Fy:    -libc.Float64FromFloat64(1.0173171765947496e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	449: {
		Ffile: __ccgo_ts,
		Fline: int32(471),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.126566341989102e-06),
		Fy:    -libc.Float64FromFloat64(4.126557827725925e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	450: {
		Ffile: __ccgo_ts,
		Fline: int32(472),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.102614274306003e-06),
		Fy:    -libc.Float64FromFloat64(6.102595653393392e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	451: {
		Ffile: __ccgo_ts,
		Fline: int32(473),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.369854005410582e-06),
		Fy:    -libc.Float64FromFloat64(6.369833717933633e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	452: {
		Ffile: __ccgo_ts,
		Fline: int32(474),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.481634881982701e-06),
		Fy:    -libc.Float64FromFloat64(5.4816198578496635e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	453: {
		Ffile: __ccgo_ts,
		Fline: int32(475),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.0815542618514994e-06),
		Fy:    -libc.Float64FromFloat64(7.081529187705305e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	454: {
		Ffile: __ccgo_ts,
		Fline: int32(476),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.935536130353019e-06),
		Fy:    -libc.Float64FromFloat64(5.9355185150932944e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	455: {
		Ffile: __ccgo_ts,
		Fline: int32(477),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.234128940404327e-06),
		Fy:    -libc.Float64FromFloat64(2.2341264447401244e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	456: {
		Ffile: __ccgo_ts,
		Fline: int32(478),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.0421209556059794e-06),
		Fy:    -libc.Float64FromFloat64(2.0421188704784e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	457: {
		Ffile: __ccgo_ts,
		Fline: int32(479),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.405397247291877e-06),
		Fy:    -libc.Float64FromFloat64(2.405394354326238e-06),
		Fdy:   float32(8.509028819716276e-16),
		Fe:    int32(FE_INEXACT),
	},
	458: {
		Ffile: __ccgo_ts,
		Fline: int32(480),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.2709384257796003),
		Fy:    -libc.Float64FromFloat64(0.23733654463298864),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	459: {
		Ffile: __ccgo_ts,
		Fline: int32(481),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.41974587597664104),
		Fy:    -libc.Float64FromFloat64(0.3427861875860397),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts,
		Fline: int32(482),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.48741146936167723),
		Fy:    -libc.Float64FromFloat64(0.3857857494035244),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts,
		Fline: int32(483),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.708002338751466e-07),
		Fy:    -libc.Float64FromFloat64(9.707997626487521e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts,
		Fline: int32(484),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1888428565786063e-06),
		Fy:    -libc.Float64FromFloat64(1.1888421499052176e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts,
		Fline: int32(485),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.3293391510174467e-06),
		Fy:    -libc.Float64FromFloat64(1.329338267446549e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts,
		Fline: int32(486),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7286630181992272e-06),
		Fy:    -libc.Float64FromFloat64(1.7286615240621727e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	465: {
		Ffile: __ccgo_ts,
		Fline: int32(487),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.826925346837129e-06),
		Fy:    -libc.Float64FromFloat64(1.8269236780100337e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	466: {
		Ffile: __ccgo_ts,
		Fline: int32(488),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.053067495121433e-07),
		Fy:    -libc.Float64FromFloat64(5.053066218447093e-07),
		Fdy:   float32(4.908216586764177e-16),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts,
		Fline: int32(489),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.750171183516263e-07),
		Fy:    -libc.Float64FromFloat64(5.750169530293147e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts,
		Fline: int32(490),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.496689968769756e-07),
		Fy:    -libc.Float64FromFloat64(8.496686359083756e-07),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts,
		Fline: int32(491),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.5295270738157185e-07),
		Fy:    -libc.Float64FromFloat64(2.529526753890385e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts,
		Fline: int32(492),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.6324489742146193e-07),
		Fy:    -libc.Float64FromFloat64(3.632448314480421e-07),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts,
		Fline: int32(493),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.111424594927801e-07),
		Fy:    -libc.Float64FromFloat64(4.111423749737307e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts,
		Fline: int32(494),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.8371355081004816e-07),
		Fy:    -libc.Float64FromFloat64(1.837135339347148e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts,
		Fline: int32(495),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4751049351043472e-07),
		Fy:    -libc.Float64FromFloat64(1.4751048263076238e-07),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts,
		Fline: int32(496),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.9666795642587178e-07),
		Fy:    -libc.Float64FromFloat64(1.966679370867305e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	475: {
		Ffile: __ccgo_ts,
		Fline: int32(497),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.9714282315011486e-07),
		Fy:    -libc.Float64FromFloat64(1.9714280371746976e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	476: {
		Ffile: __ccgo_ts,
		Fline: int32(498),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0103032385985334e-07),
		Fy:    -libc.Float64FromFloat64(1.0103031875629033e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	477: {
		Ffile: __ccgo_ts,
		Fline: int32(499),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.590318781107688e-08),
		Fy:    -libc.Float64FromFloat64(3.590318716655744e-08),
		Fdy:   float32(7.151141831643435e-17),
		Fe:    int32(FE_INEXACT),
	},
	478: {
		Ffile: __ccgo_ts,
		Fline: int32(500),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.349324366295251e-08),
		Fy:    -libc.Float64FromFloat64(4.3493242717121395e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	479: {
		Ffile: __ccgo_ts,
		Fline: int32(501),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.062084194041831e-08),
		Fy:    -libc.Float64FromFloat64(5.062084065918351e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	480: {
		Ffile: __ccgo_ts,
		Fline: int32(502),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.376315463938804e-08),
		Fy:    -libc.Float64FromFloat64(5.376315319414966e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	481: {
		Ffile: __ccgo_ts,
		Fline: int32(503),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.4687161574958236e-08),
		Fy:    -libc.Float64FromFloat64(5.4687160079615436e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	482: {
		Ffile: __ccgo_ts,
		Fline: int32(504),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.550672766331306e-08),
		Fy:    -libc.Float64FromFloat64(1.550672754308376e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	483: {
		Ffile: __ccgo_ts,
		Fline: int32(505),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.9348617623328615e-08),
		Fy:    -libc.Float64FromFloat64(1.934861743614411e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	484: {
		Ffile: __ccgo_ts,
		Fline: int32(506),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0920076380191543e-08),
		Fy:    -libc.Float64FromFloat64(1.0920076320567509e-08),
		Fdy:   float32(8.846300398653994e-16),
		Fe:    int32(FE_INEXACT),
	},
	485: {
		Ffile: __ccgo_ts,
		Fline: int32(507),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2892140904151637e-08),
		Fy:    -libc.Float64FromFloat64(1.2892140821047989e-08),
		Fdy:   float32(4.759398191887516e-16),
		Fe:    int32(FE_INEXACT),
	},
	486: {
		Ffile: __ccgo_ts,
		Fline: int32(508),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4397018257994494e-08),
		Fy:    -libc.Float64FromFloat64(1.4397018154357427e-08),
		Fdy:   float32(1.9489320202881113e-16),
		Fe:    int32(FE_INEXACT),
	},
	487: {
		Ffile: __ccgo_ts,
		Fline: int32(509),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4822964329134183e-08),
		Fy:    -libc.Float64FromFloat64(1.4822964219274046e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	488: {
		Ffile: __ccgo_ts,
		Fline: int32(510),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.818763781256577e-09),
		Fy:    -libc.Float64FromFloat64(3.818763773965099e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	489: {
		Ffile: __ccgo_ts,
		Fline: int32(511),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.494480376497895e-09),
		Fy:    -libc.Float64FromFloat64(6.494480355408756e-09),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	490: {
		Ffile: __ccgo_ts,
		Fline: int32(512),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.425039586285558e-09),
		Fy:    -libc.Float64FromFloat64(2.4250395833451498e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	491: {
		Ffile: __ccgo_ts,
		Fline: int32(513),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.14504743082833096),
		Fy:    -libc.Float64FromFloat64(0.1350187346401479),
		Fdy:   float32(1.4406071669105123e-17),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts,
		Fline: int32(514),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.15050763135276474),
		Fy:    -libc.Float64FromFloat64(0.13972883505037),
		Fdy:   float32(7.021032276989254e-17),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts,
		Fline: int32(515),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.1540712945707682),
		Fy:    -libc.Float64FromFloat64(0.14278909565064318),
		Fdy:   float32(2.8833957083580005e-16),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts,
		Fline: int32(516),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.21488247280794578),
		Fy:    -libc.Float64FromFloat64(0.19336376370148184),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts,
		Fline: int32(517),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.2255238461437243),
		Fy:    -libc.Float64FromFloat64(0.20190197133932297),
		Fdy:   float32(5.764810947806202e-16),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts,
		Fline: int32(518),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.2429577998684119),
		Fy:    -libc.Float64FromFloat64(0.21569538916901634),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts,
		Fline: int32(519),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2969470187939886e-09),
		Fy:    -libc.Float64FromFloat64(1.2969470179529529e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts,
		Fline: int32(520),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2973573721252216e-09),
		Fy:    -libc.Float64FromFloat64(1.2973573712836536e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts,
		Fline: int32(521),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.748458030205084e-09),
		Fy:    -libc.Float64FromFloat64(1.7484580286765312e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts,
		Fline: int32(522),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.188218884163258e-10),
		Fy:    -libc.Float64FromFloat64(5.188218882817376e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts,
		Fline: int32(523),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.156972093412989e-10),
		Fy:    -libc.Float64FromFloat64(7.156972090851877e-10),
		Fdy:   float32(5.205532033393133e-16),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts,
		Fline: int32(524),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.46751251173367e-10),
		Fy:    -libc.Float64FromFloat64(2.467512511429239e-10),
		Fdy:   float32(6.435712922799627e-16),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts,
		Fline: int32(525),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.0876494186569483e-10),
		Fy:    -libc.Float64FromFloat64(4.087649417821504e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts,
		Fline: int32(526),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.4406812274761725e-10),
		Fy:    -libc.Float64FromFloat64(4.4406812264901903e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts,
		Fline: int32(527),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.453639936167446e-10),
		Fy:    -libc.Float64FromFloat64(1.4536399360617925e-10),
		Fdy:   float32(5.337417180862995e-16),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts,
		Fline: int32(528),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.8717411458937195e-10),
		Fy:    -libc.Float64FromFloat64(1.8717411457185487e-10),
		Fdy:   float32(4.1386355381134215e-16),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts,
		Fline: int32(529),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.866240826492742e-11),
		Fy:    -libc.Float64FromFloat64(5.866240826320678e-11),
		Fdy:   float32(2.5451672468876818e-17),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts,
		Fline: int32(530),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.548361852836552e-11),
		Fy:    -libc.Float64FromFloat64(6.548361852622147e-11),
		Fdy:   float32(3.9519142717317035e-17),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts,
		Fline: int32(531),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.207301339664453e-11),
		Fy:    -libc.Float64FromFloat64(6.2073013394718e-11),
		Fdy:   float32(3.190717525769786e-17),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts,
		Fline: int32(532),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.571543392355175e-11),
		Fy:    -libc.Float64FromFloat64(7.571543392068533e-11),
		Fdy:   float32(7.063415688087491e-17),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts,
		Fline: int32(533),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.176692162977554e-11),
		Fy:    -libc.Float64FromFloat64(8.176692162643262e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts,
		Fline: int32(534),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.906860308958056e-11),
		Fy:    -libc.Float64FromFloat64(7.906860308645464e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts,
		Fline: int32(535),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.253664418702862e-11),
		Fy:    -libc.Float64FromFloat64(8.253664418362247e-11),
		Fdy:   float32(9.973884422324311e-17),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts,
		Fline: int32(536),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.230482879181913e-11),
		Fy:    -libc.Float64FromFloat64(7.230482878920513e-11),
		Fdy:   float32(5.874165548274692e-17),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts,
		Fline: int32(537),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.889422366009038e-11),
		Fy:    -libc.Float64FromFloat64(6.889422365771718e-11),
		Fdy:   float32(4.841815027050952e-17),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts,
		Fline: int32(538),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.900819143670994e-11),
		Fy:    -libc.Float64FromFloat64(6.900819143432887e-11),
		Fdy:   float32(1.2099201260059678e-16),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts,
		Fline: int32(539),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.912603905528825e-11),
		Fy:    -libc.Float64FromFloat64(7.912603905215778e-11),
		Fdy:   float32(8.424706160421519e-17),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts,
		Fline: int32(540),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0982148524109121e-10),
		Fy:    -libc.Float64FromFloat64(1.0982148523506083e-10),
		Fdy:   float32(3.1262691710933703e-16),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts,
		Fline: int32(541),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.884284376406889e-11),
		Fy:    -libc.Float64FromFloat64(9.884284375918393e-11),
		Fdy:   float32(3.828270754838625e-16),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts,
		Fline: int32(542),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.276845958227301e-11),
		Fy:    -libc.Float64FromFloat64(9.276845957797002e-11),
		Fdy:   float32(1.5917631080237984e-16),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts,
		Fline: int32(543),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.594724931877287e-11),
		Fy:    -libc.Float64FromFloat64(8.594724931507941e-11),
		Fdy:   float32(1.1727488792091009e-16),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts,
		Fline: int32(544),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.9357854450521e-11),
		Fy:    -libc.Float64FromFloat64(8.935785444652859e-11),
		Fdy:   float32(1.3702759037176196e-16),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts,
		Fline: int32(545),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.030002749775523e-10),
		Fy:    -libc.Float64FromFloat64(1.0300027497224777e-10),
		Fdy:   float32(2.418970984449816e-16),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts,
		Fline: int32(546),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.958966984578866e-11),
		Fy:    -libc.Float64FromFloat64(9.958966984082961e-11),
		Fdy:   float32(2.1141426202765103e-16),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts,
		Fline: int32(547),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.61790647140289e-11),
		Fy:    -libc.Float64FromFloat64(9.617906470940369e-11),
		Fdy:   float32(1.8390744940071887e-16),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts,
		Fline: int32(548),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0641088010931982e-10),
		Fy:    -libc.Float64FromFloat64(1.0641088010365818e-10),
		Fdy:   float32(2.755632170269918e-16),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts,
		Fline: int32(549),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1323209037286648e-10),
		Fy:    -libc.Float64FromFloat64(1.1323209036645573e-10),
		Fdy:   float32(3.5330940664014864e-16),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts,
		Fline: int32(550),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.0354385671825045e-11),
		Fy:    -libc.Float64FromFloat64(3.035438567136435e-11),
		Fdy:   float32(3.649150285669534e-18),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts,
		Fline: int32(551),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.069544618499336e-11),
		Fy:    -libc.Float64FromFloat64(3.069544618452226e-11),
		Fdy:   float32(3.8159363669398264e-18),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts,
		Fline: int32(552),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.2400748750835537e-11),
		Fy:    -libc.Float64FromFloat64(3.240074875031063e-11),
		Fdy:   float32(4.737243515756488e-18),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts,
		Fline: int32(553),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.2059688237667024e-11),
		Fy:    -libc.Float64FromFloat64(3.205968823715311e-11),
		Fdy:   float32(4.5409171604507024e-18),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts,
		Fline: int32(554),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.3764990803509973e-11),
		Fy:    -libc.Float64FromFloat64(3.3764990802939936e-11),
		Fdy:   float32(5.586923440970958e-18),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts,
		Fline: int32(555),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.3260638279668915e-11),
		Fy:    -libc.Float64FromFloat64(3.3260638279115783e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts,
		Fline: int32(556),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.410605131667868e-11),
		Fy:    -libc.Float64FromFloat64(3.410605131609707e-11),
		Fdy:   float32(5.81609879276245e-18),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts,
		Fline: int32(557),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.547029336935389e-11),
		Fy:    -libc.Float64FromFloat64(3.547029336872482e-11),
		Fdy:   float32(6.804003907056981e-18),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts,
		Fline: int32(558),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.698187572774437e-11),
		Fy:    -libc.Float64FromFloat64(3.6981875727060546e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts,
		Fline: int32(559),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.717559593519878e-11),
		Fy:    -libc.Float64FromFloat64(3.7175595934507767e-11),
		Fdy:   float32(8.209893311041908e-18),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts,
		Fline: int32(560),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.8880898501044636e-11),
		Fy:    -libc.Float64FromFloat64(3.8880898500288774e-11),
		Fdy:   float32(9.82317343820613e-18),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts,
		Fline: int32(561),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.3996806198588023e-11),
		Fy:    -libc.Float64FromFloat64(4.3996806197620164e-11),
		Fdy:   float32(1.610614620358331e-17),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts,
		Fline: int32(562),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.0927261580060944e-11),
		Fy:    -libc.Float64FromFloat64(4.0927261579223424e-11),
		Fdy:   float32(1.2060266861243542e-17),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts,
		Fline: int32(563),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.2632564145908934e-11),
		Fy:    -libc.Float64FromFloat64(4.2632564145000166e-11),
		Fdy:   float32(1.419947503931388e-17),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts,
		Fline: int32(564),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.774847184345872e-11),
		Fy:    -libc.Float64FromFloat64(4.774847184231876e-11),
		Fdy:   float32(2.2343167431571943e-17),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts,
		Fline: int32(565),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.433786671175789e-11),
		Fy:    -libc.Float64FromFloat64(4.433786671077497e-11),
		Fdy:   float32(1.66114013566205e-17),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts,
		Fline: int32(566),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.252331902784639e-11),
		Fy:    -libc.Float64FromFloat64(5.252331902646704e-11),
		Fdy:   float32(3.2712650650977494e-17),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts,
		Fline: int32(567),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.9453774409310586e-11),
		Fy:    -libc.Float64FromFloat64(4.945377440808775e-11),
		Fdy:   float32(2.5710149866687388e-17),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts,
		Fline: int32(568),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.9221959014213924e-11),
		Fy:    -libc.Float64FromFloat64(3.922195901344474e-11),
		Fdy:   float32(1.0172415710270918e-17),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts,
		Fline: int32(569),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.058620106689146e-11),
		Fy:    -libc.Float64FromFloat64(4.058620106606784e-11),
		Fdy:   float32(1.1663273106777292e-17),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts,
		Fline: int32(570),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.229150363273926e-11),
		Fy:    -libc.Float64FromFloat64(4.229150363184497e-11),
		Fdy:   float32(1.3750521107489498e-17),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts,
		Fline: int32(571),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.604316927760782e-11),
		Fy:    -libc.Float64FromFloat64(4.6043169276547834e-11),
		Fdy:   float32(1.9318253958249226e-17),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts,
		Fline: int32(572),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.740741133028846e-11),
		Fy:    -libc.Float64FromFloat64(4.740741132916473e-11),
		Fdy:   float32(2.1711598490275455e-17),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts,
		Fline: int32(573),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.081801646199278e-11),
		Fy:    -libc.Float64FromFloat64(5.081801646070154e-11),
		Fdy:   float32(2.866669189929892e-17),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts,
		Fline: int32(574),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.286437954101723e-11),
		Fy:    -libc.Float64FromFloat64(5.2864379539619906e-11),
		Fdy:   float32(3.35706420869869e-17),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts,
		Fline: int32(575),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.115907697516342e-11),
		Fy:    -libc.Float64FromFloat64(5.1159076973854796e-11),
		Fdy:   float32(2.9444056386641556e-17),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts,
		Fline: int32(576),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.581135388252279e-11),
		Fy:    -libc.Float64FromFloat64(3.5811353881881565e-11),
		Fdy:   float32(7.069522266241603e-18),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts,
		Fline: int32(577),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.751665644836787e-11),
		Fy:    -libc.Float64FromFloat64(3.7516656447664123e-11),
		Fdy:   float32(8.51535456764549e-18),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts,
		Fline: int32(578),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.570210876443776e-11),
		Fy:    -libc.Float64FromFloat64(4.5702108763393416e-11),
		Fdy:   float32(1.8752171251687337e-17),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts,
		Fline: int32(579),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.9112713896140135e-11),
		Fy:    -libc.Float64FromFloat64(4.9112713894934105e-11),
		Fdy:   float32(2.500819778142999e-17),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts,
		Fline: int32(580),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.422862159370097e-11),
		Fy:    -libc.Float64FromFloat64(5.42286215922306e-11),
		Fdy:   float32(3.71724644067286e-17),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts,
		Fline: int32(581),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.627498467272775e-11),
		Fy:    -libc.Float64FromFloat64(5.627498467114431e-11),
		Fdy:   float32(4.310905363623797e-17),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts,
		Fline: int32(582),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.4569682106872e-11),
		Fy:    -libc.Float64FromFloat64(5.456968210538308e-11),
		Fdy:   float32(3.8116456156663914e-17),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts,
		Fline: int32(583),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.798028723858446e-11),
		Fy:    -libc.Float64FromFloat64(5.7980287236903604e-11),
		Fdy:   float32(4.857675718988289e-17),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts,
		Fline: int32(584),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.763922672541304e-11),
		Fy:    -libc.Float64FromFloat64(5.76392267237519e-11),
		Fdy:   float32(4.744382415315066e-17),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts,
		Fline: int32(585),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.7759011600109234e-11),
		Fy:    -libc.Float64FromFloat64(5.775901159844118e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts,
		Fline: int32(586),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.593392415955652e-11),
		Fy:    -libc.Float64FromFloat64(5.593392415799222e-11),
		Fdy:   float32(4.207344997910118e-17),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts,
		Fline: int32(587),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6029844118790686e-11),
		Fy:    -libc.Float64FromFloat64(1.6029844118662208e-11),
		Fdy:   float32(5.675914839991863e-19),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts,
		Fline: int32(588),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7735146684624523e-11),
		Fy:    -libc.Float64FromFloat64(1.7735146684467255e-11),
		Fdy:   float32(8.505004883821226e-19),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts,
		Fline: int32(589),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.4556356947969565e-11),
		Fy:    -libc.Float64FromFloat64(2.4556356947668058e-11),
		Fdy:   float32(3.126001561621597e-18),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts,
		Fline: int32(590),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.626165951380825e-11),
		Fy:    -libc.Float64FromFloat64(2.626165951346341e-11),
		Fdy:   float32(4.0890780226497366e-18),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts,
		Fline: int32(591),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.8262933297252568e-11),
		Fy:    -libc.Float64FromFloat64(1.82629332970858e-11),
		Fdy:   float32(6.232137024256529e-16),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts,
		Fline: int32(592),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.944044925045933e-11),
		Fy:    -libc.Float64FromFloat64(1.9440449250270364e-11),
		Fdy:   float32(1.2278801361635151e-18),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts,
		Fline: int32(593),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.1145751816295105e-11),
		Fy:    -libc.Float64FromFloat64(2.1145751816071534e-11),
		Fdy:   float32(1.7188151384361873e-18),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts,
		Fline: int32(594),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.285105438213185e-11),
		Fy:    -libc.Float64FromFloat64(2.2851054381870765e-11),
		Fdy:   float32(2.344004862848666e-18),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts,
		Fline: int32(595),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.3712632626607197e-11),
		Fy:    -libc.Float64FromFloat64(2.371263262632605e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts,
		Fline: int32(596),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.441687867344865e-11),
		Fy:    -libc.Float64FromFloat64(2.4416878673150557e-11),
		Fdy:   float32(8.5711634509577075e-16),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts,
		Fline: int32(597),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.7966962079647902e-11),
		Fy:    -libc.Float64FromFloat64(2.7966962079256826e-11),
		Fdy:   float32(5.2591746299427474e-18),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts,
		Fline: int32(598),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.503331289636041e-12),
		Fy:    -libc.Float64FromFloat64(7.503331289607891e-12),
		Fdy:   float32(5.447480642029219e-20),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts,
		Fline: int32(599),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.03354586952622e-12),
		Fy:    -libc.Float64FromFloat64(8.033545869493953e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts,
		Fline: int32(600),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.87973633632019e-12),
		Fy:    -libc.Float64FromFloat64(7.879736336289145e-12),
		Fdy:   float32(6.84220244600019e-16),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts,
		Fline: int32(601),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.355982572550415e-12),
		Fy:    -libc.Float64FromFloat64(8.355982572515504e-12),
		Fdy:   float32(8.380332221897234e-20),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts,
		Fline: int32(602),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.208633855465032e-12),
		Fy:    -libc.Float64FromFloat64(9.208633855422632e-12),
		Fdy:   float32(1.2361387073992446e-19),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts,
		Fline: int32(603),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.402475395023153e-12),
		Fy:    -libc.Float64FromFloat64(9.40247539497895e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts,
		Fline: int32(604),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.006128513837989e-11),
		Fy:    -libc.Float64FromFloat64(1.0061285138329275e-11),
		Fdy:   float32(1.7618285302889447e-19),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts,
		Fline: int32(605),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1711054921902417e-11),
		Fy:    -libc.Float64FromFloat64(1.1711054921833843e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts,
		Fline: int32(606),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1766587704210335e-11),
		Fy:    -libc.Float64FromFloat64(1.1766587704141108e-11),
		Fdy:   float32(3.295752258207279e-19),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts,
		Fline: int32(607),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.193028125860296e-11),
		Fy:    -libc.Float64FromFloat64(1.1930281258531795e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts,
		Fline: int32(608),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1158480399339078e-11),
		Fy:    -libc.Float64FromFloat64(1.1158480399276823e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	587: {
		Ffile: __ccgo_ts,
		Fline: int32(609),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.261923898712592e-11),
		Fy:    -libc.Float64FromFloat64(1.2619238987046298e-11),
		Fdy:   float32(4.360102095991511e-19),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts,
		Fline: int32(610),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0913936421294991e-11),
		Fy:    -libc.Float64FromFloat64(1.0913936421235434e-11),
		Fdy:   float32(2.439190190296368e-19),
		Fe:    int32(FE_INEXACT),
	},
	589: {
		Ffile: __ccgo_ts,
		Fline: int32(611),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2274822674669368e-11),
		Fy:    -libc.Float64FromFloat64(1.2274822674594032e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	590: {
		Ffile: __ccgo_ts,
		Fline: int32(612),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3608865327700474e-11),
		Fy:    -libc.Float64FromFloat64(1.3608865327607872e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts,
		Fline: int32(613),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3471890270041748e-11),
		Fy:    -libc.Float64FromFloat64(1.3471890269951002e-11),
		Fdy:   float32(5.663209345783049e-19),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts,
		Fline: int32(614),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.3507280978945485e-11),
		Fy:    -libc.Float64FromFloat64(1.3507280978854263e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts,
		Fline: int32(615),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4324541552957818e-11),
		Fy:    -libc.Float64FromFloat64(1.4324541552855222e-11),
		Fdy:   float32(7.238955325472064e-19),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts,
		Fline: int32(616),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.026665180489709e-12),
		Fy:    -libc.Float64FromFloat64(4.026665180481601e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts,
		Fline: int32(617),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.694056946832428e-12),
		Fy:    -libc.Float64FromFloat64(3.694056946825605e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts,
		Fline: int32(618),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.6043169277289825e-12),
		Fy:    -libc.Float64FromFloat64(4.604316927718383e-12),
		Fdy:   float32(1.5431881507789284e-20),
		Fe:    int32(FE_INEXACT),
	},
	597: {
		Ffile: __ccgo_ts,
		Fline: int32(619),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.565733324302719e-12),
		Fy:    -libc.Float64FromFloat64(4.565733324292296e-12),
		Fdy:   float32(1.5557927951274605e-16),
		Fe:    int32(FE_INEXACT),
	},
	598: {
		Ffile: __ccgo_ts,
		Fline: int32(620),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.751665644815675e-12),
		Fy:    -libc.Float64FromFloat64(3.751665644808637e-12),
		Fdy:   float32(6.8027333576361e-21),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts,
		Fline: int32(621),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.4337866711463016e-12),
		Fy:    -libc.Float64FromFloat64(4.433786671136472e-12),
		Fdy:   float32(1.3287829360051837e-20),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts,
		Fline: int32(622),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.991740519887533e-12),
		Fy:    -libc.Float64FromFloat64(6.991740519863091e-12),
		Fdy:   float32(8.216219588366713e-20),
		Fe:    int32(FE_INEXACT),
	},
	601: {
		Ffile: __ccgo_ts,
		Fline: int32(623),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.286437954059803e-12),
		Fy:    -libc.Float64FromFloat64(5.28643795404583e-12),
		Fdy:   float32(2.6840356516120642e-20),
		Fe:    int32(FE_INEXACT),
	},
	602: {
		Ffile: __ccgo_ts,
		Fline: int32(624),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.309619493556325e-12),
		Fy:    -libc.Float64FromFloat64(6.309619493536419e-12),
		Fdy:   float32(5.450127619989389e-20),
		Fe:    int32(FE_INEXACT),
	},
	603: {
		Ffile: __ccgo_ts,
		Fline: int32(625),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.1622707764703596e-12),
		Fy:    -libc.Float64FromFloat64(7.1622707764447105e-12),
		Fdy:   float32(9.047370667859995e-20),
		Fe:    int32(FE_INEXACT),
	},
	604: {
		Ffile: __ccgo_ts,
		Fline: int32(626),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.4569682106425325e-12),
		Fy:    -libc.Float64FromFloat64(5.456968210627643e-12),
		Fdy:   float32(3.0466716321553115e-20),
		Fe:    int32(FE_INEXACT),
	},
	605: {
		Ffile: __ccgo_ts,
		Fline: int32(627),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.139089236973547e-12),
		Fy:    -libc.Float64FromFloat64(6.139089236954703e-12),
		Fdy:   float32(4.8836743365130754e-20),
		Fe:    int32(FE_INEXACT),
	},
	606: {
		Ffile: __ccgo_ts,
		Fline: int32(628),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.76158995246347e-12),
		Fy:    -libc.Float64FromFloat64(2.7615899524596564e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	607: {
		Ffile: __ccgo_ts,
		Fline: int32(629),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.06460165726209977),
		Fy:    -libc.Float64FromFloat64(0.06255918824826757),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	608: {
		Ffile: __ccgo_ts,
		Fline: int32(630),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.07221637344469926),
		Fy:    -libc.Float64FromFloat64(0.06967042458314097),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	609: {
		Ffile: __ccgo_ts,
		Fline: int32(631),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.07433682283844194),
		Fy:    -libc.Float64FromFloat64(0.0716410513217996),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	610: {
		Ffile: __ccgo_ts,
		Fline: int32(632),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.09377161882257719),
		Fy:    -libc.Float64FromFloat64(0.08950932256915058),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	611: {
		Ffile: __ccgo_ts,
		Fline: int32(633),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.10318926109567506),
		Fy:    -libc.Float64FromFloat64(0.09804374790195503),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	612: {
		Ffile: __ccgo_ts,
		Fline: int32(634),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.10903804320464712),
		Fy:    -libc.Float64FromFloat64(0.10330369635164652),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	613: {
		Ffile: __ccgo_ts,
		Fline: int32(635),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.1118679760732042),
		Fy:    -libc.Float64FromFloat64(0.10583769947395787),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	614: {
		Ffile: __ccgo_ts,
		Fline: int32(636),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.365158115655264e-13),
		Fy:    -libc.Float64FromFloat64(9.36515811565088e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	615: {
		Ffile: __ccgo_ts,
		Fline: int32(637),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0077690519377925e-12),
		Fy:    -libc.Float64FromFloat64(1.0077690519372848e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	616: {
		Ffile: __ccgo_ts,
		Fline: int32(638),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.379164112034789e-13),
		Fy:    -libc.Float64FromFloat64(9.37916411203039e-13),
		Fdy:   float32(1.0587911840678754e-22),
		Fe:    int32(FE_INEXACT),
	},
	617: {
		Ffile: __ccgo_ts,
		Fline: int32(639),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3642420526597025e-12),
		Fy:    -libc.Float64FromFloat64(1.364242052658772e-12),
		Fdy:   float32(4.499862532288471e-22),
		Fe:    int32(FE_INEXACT),
	},
	618: {
		Ffile: __ccgo_ts,
		Fline: int32(640),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4147475076854451e-12),
		Fy:    -libc.Float64FromFloat64(1.4147475076844443e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	619: {
		Ffile: __ccgo_ts,
		Fline: int32(641),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.278976924368453e-12),
		Fy:    -libc.Float64FromFloat64(1.278976924367635e-12),
		Fdy:   float32(3.441071348220595e-22),
		Fe:    int32(FE_INEXACT),
	},
	620: {
		Ffile: __ccgo_ts,
		Fline: int32(642),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.3063942201979357e-12),
		Fy:    -libc.Float64FromFloat64(1.3063942201970825e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	621: {
		Ffile: __ccgo_ts,
		Fline: int32(643),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6848105514421937e-12),
		Fy:    -libc.Float64FromFloat64(1.6848105514407744e-12),
		Fdy:   float32(4.018829345169201e-16),
		Fe:    int32(FE_INEXACT),
	},
	622: {
		Ffile: __ccgo_ts,
		Fline: int32(644),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7905676941159868e-12),
		Fy:    -libc.Float64FromFloat64(1.7905676941143838e-12),
		Fdy:   float32(1.402898318889935e-21),
		Fe:    int32(FE_INEXACT),
	},
	623: {
		Ffile: __ccgo_ts,
		Fline: int32(645),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7053025658247251e-12),
		Fy:    -libc.Float64FromFloat64(1.705302565823271e-12),
		Fdy:   float32(1.138200522872966e-21),
		Fe:    int32(FE_INEXACT),
	},
	624: {
		Ffile: __ccgo_ts,
		Fline: int32(646),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.602649920770685e-13),
		Fy:    -libc.Float64FromFloat64(4.602649920769624e-13),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	625: {
		Ffile: __ccgo_ts,
		Fline: int32(647),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.550880341147472e-13),
		Fy:    -libc.Float64FromFloat64(6.550880341145325e-13),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	626: {
		Ffile: __ccgo_ts,
		Fline: int32(648),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.358745046066486e-13),
		Fy:    -libc.Float64FromFloat64(2.3587450460662074e-13),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	627: {
		Ffile: __ccgo_ts,
		Fline: int32(649),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.57608471749767e-13),
		Fy:    -libc.Float64FromFloat64(3.5760847174970304e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	628: {
		Ffile: __ccgo_ts,
		Fline: int32(650),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.137978058668471e-13),
		Fy:    -libc.Float64FromFloat64(1.1379780586684062e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	629: {
		Ffile: __ccgo_ts,
		Fline: int32(651),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4921397450962475e-13),
		Fy:    -libc.Float64FromFloat64(1.4921397450961362e-13),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	630: {
		Ffile: __ccgo_ts,
		Fline: int32(652),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.1796412355255547e-13),
		Fy:    -libc.Float64FromFloat64(2.1796412355253172e-13),
		Fdy:   float32(2.885929925057067e-16),
		Fe:    int32(FE_INEXACT),
	},
	631: {
		Ffile: __ccgo_ts,
		Fline: int32(653),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.140289677804273e-14),
		Fy:    -libc.Float64FromFloat64(8.140289677803941e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	632: {
		Ffile: __ccgo_ts,
		Fline: int32(654),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.999524775355733e-14),
		Fy:    -libc.Float64FromFloat64(7.999524775355412e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	633: {
		Ffile: __ccgo_ts,
		Fline: int32(655),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0272197169173017e-13),
		Fy:    -libc.Float64FromFloat64(1.027219716917249e-13),
		Fdy:   float32(8.550601196767517e-16),
		Fe:    int32(FE_INEXACT),
	},
	634: {
		Ffile: __ccgo_ts,
		Fline: int32(656),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.197442310920468e-14),
		Fy:    -libc.Float64FromFloat64(3.197442310920417e-14),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	635: {
		Ffile: __ccgo_ts,
		Fline: int32(657),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.066467347002827e-14),
		Fy:    -libc.Float64FromFloat64(3.06646734700278e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	636: {
		Ffile: __ccgo_ts,
		Fline: int32(658),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.240993826809858e-14),
		Fy:    -libc.Float64FromFloat64(4.2409938268097683e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	637: {
		Ffile: __ccgo_ts,
		Fline: int32(659),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.4798014943547264e-14),
		Fy:    -libc.Float64FromFloat64(4.4798014943546255e-14),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	638: {
		Ffile: __ccgo_ts,
		Fline: int32(660),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.2755106611775306e-14),
		Fy:    -libc.Float64FromFloat64(5.275510661177391e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	639: {
		Ffile: __ccgo_ts,
		Fline: int32(661),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.329070518200799e-14),
		Fy:    -libc.Float64FromFloat64(5.329070518200657e-14),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	640: {
		Ffile: __ccgo_ts,
		Fline: int32(662),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7852165294699987e-14),
		Fy:    -libc.Float64FromFloat64(1.785216529469983e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	641: {
		Ffile: __ccgo_ts,
		Fline: int32(663),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.057678187205895e-15),
		Fy:    -libc.Float64FromFloat64(9.057678187205854e-15),
		Fdy:   float32(4.3200845537940736e-16),
		Fe:    int32(FE_INEXACT),
	},
	642: {
		Ffile: __ccgo_ts,
		Fline: int32(664),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0357851278622317e-14),
		Fy:    -libc.Float64FromFloat64(1.0357851278622264e-14),
		Fdy:   float32(2.845759916929124e-16),
		Fe:    int32(FE_INEXACT),
	},
	643: {
		Ffile: __ccgo_ts,
		Fline: int32(665),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1982181969500577e-14),
		Fy:    -libc.Float64FromFloat64(1.1982181969500506e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	644: {
		Ffile: __ccgo_ts,
		Fline: int32(666),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2992930143631854e-14),
		Fy:    -libc.Float64FromFloat64(1.2992930143631768e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	645: {
		Ffile: __ccgo_ts,
		Fline: int32(667),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.0242958677880845e-15),
		Fy:    -libc.Float64FromFloat64(5.024295867788071e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	646: {
		Ffile: __ccgo_ts,
		Fline: int32(668),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.864753555590498e-15),
		Fy:    -libc.Float64FromFloat64(4.8647535555904854e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	647: {
		Ffile: __ccgo_ts,
		Fline: int32(669),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.404745667978761e-15),
		Fy:    -libc.Float64FromFloat64(6.40474566797874e-15),
		Fdy:   float32(2.917725159616829e-16),
		Fe:    int32(FE_INEXACT),
	},
	648: {
		Ffile: __ccgo_ts,
		Fline: int32(670),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.2633758932252457e-15),
		Fy:    -libc.Float64FromFloat64(3.2633758932252406e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	649: {
		Ffile: __ccgo_ts,
		Fline: int32(671),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.0119595631484904e-15),
		Fy:    -libc.Float64FromFloat64(3.0119595631484857e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	650: {
		Ffile: __ccgo_ts,
		Fline: int32(672),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.03791540701945831),
		Fy:    -libc.Float64FromFloat64(0.03720561690705369),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	651: {
		Ffile: __ccgo_ts,
		Fline: int32(673),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0487060321253327),
		Fy:    -libc.Float64FromFloat64(0.047538918492096204),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	652: {
		Ffile: __ccgo_ts,
		Fline: int32(674),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.06037856877519905),
		Fy:    -libc.Float64FromFloat64(0.058591921586100465),
		Fdy:   float32(7.262695012005642e-16),
		Fe:    int32(FE_INEXACT),
	},
	653: {
		Ffile: __ccgo_ts,
		Fline: int32(675),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.332267629550188e-15),
		Fy:    -libc.Float64FromFloat64(1.3322676295501873e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	654: {
		Ffile: __ccgo_ts,
		Fline: int32(676),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5383701491068516e-15),
		Fy:    -libc.Float64FromFloat64(1.5383701491068503e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	655: {
		Ffile: __ccgo_ts,
		Fline: int32(677),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4043333874306807e-15),
		Fy:    -libc.Float64FromFloat64(1.4043333874306797e-15),
		Fdy:   float32(3.3289707728162247e-16),
		Fe:    int32(FE_INEXACT),
	},
	656: {
		Ffile: __ccgo_ts,
		Fline: int32(678),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6616296724220901e-15),
		Fy:    -libc.Float64FromFloat64(1.6616296724220887e-15),
		Fdy:   float32(7.637445151881402e-16),
		Fe:    int32(FE_INEXACT),
	},
	657: {
		Ffile: __ccgo_ts,
		Fline: int32(679),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    -libc.Float64FromFloat64(4.440892098500625e-16),
		Fdy:   float32(2.9605945558685534e-16),
		Fe:    int32(FE_INEXACT),
	},
	658: {
		Ffile: __ccgo_ts,
		Fline: int32(680),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.438959822042073e-16),
		Fy:    -libc.Float64FromFloat64(5.438959822042072e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	659: {
		Ffile: __ccgo_ts,
		Fline: int32(681),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.01718250623968118),
		Fy:    -libc.Float64FromFloat64(0.017035728849501188),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	660: {
		Ffile: __ccgo_ts,
		Fline: int32(682),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.026200864479135753),
		Fy:    -libc.Float64FromFloat64(0.025860600047320253),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	661: {
		Ffile: __ccgo_ts,
		Fline: int32(683),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.02539445466106131),
		Fy:    -libc.Float64FromFloat64(0.0250747276460027),
		Fdy:   float32(6.774380517913538e-16),
		Fe:    int32(FE_INEXACT),
	},
	662: {
		Ffile: __ccgo_ts,
		Fline: int32(684),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.027882488208207105),
		Fy:    -libc.Float64FromFloat64(0.02749735938532294),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	663: {
		Ffile: __ccgo_ts,
		Fline: int32(685),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.008445283111580014),
		Fy:    -libc.Float64FromFloat64(0.008409721886780202),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	664: {
		Ffile: __ccgo_ts,
		Fline: int32(686),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.011990336563785989),
		Fy:    -libc.Float64FromFloat64(0.011918738923963144),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	665: {
		Ffile: __ccgo_ts,
		Fline: int32(687),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00442437387193802),
		Fy:    -libc.Float64FromFloat64(0.004414600748488612),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	666: {
		Ffile: __ccgo_ts,
		Fline: int32(688),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.005082240677781677),
		Fy:    -libc.Float64FromFloat64(0.005069347943202041),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	667: {
		Ffile: __ccgo_ts,
		Fline: int32(689),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0028657859325618654),
		Fy:    -libc.Float64FromFloat64(0.0028616834879013468),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	668: {
		Ffile: __ccgo_ts,
		Fline: int32(690),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0027957646593622133),
		Fy:    -libc.Float64FromFloat64(0.0027918601488919927),
		Fdy:   float32(3.531920131676151e-16),
		Fe:    int32(FE_INEXACT),
	},
	669: {
		Ffile: __ccgo_ts,
		Fline: int32(691),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0034071705243246472),
		Fy:    -libc.Float64FromFloat64(0.003401372705422137),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	670: {
		Ffile: __ccgo_ts,
		Fline: int32(692),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0403571407112755),
		Fy:    -libc.Float64FromFloat64(0.6466715285589396),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	671: {
		Ffile: __ccgo_ts,
		Fline: int32(693),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.3358234132908153),
		Fy:    -libc.Float64FromFloat64(0.737058423295845),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	672: {
		Ffile: __ccgo_ts,
		Fline: int32(694),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5654316209026886),
		Fy:    -libc.Float64FromFloat64(0.7910022141103128),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	673: {
		Ffile: __ccgo_ts,
		Fline: int32(695),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0403571407112755),
		Fy:    -libc.Float64FromFloat64(0.6466715285589396),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	674: {
		Ffile: __ccgo_ts,
		Fline: int32(696),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5654316209026886),
		Fy:    -libc.Float64FromFloat64(0.7910022141103128),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	675: {
		Ffile: __ccgo_ts,
		Fline: int32(697),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(35.63818828100899),
		Fy:    -libc.Float64FromFloat64(0.9999999999999997),
		Fdy:   float32(5.033383174775537e-15),
		Fe:    int32(FE_INEXACT),
	},
	676: {
		Ffile: __ccgo_ts,
		Fline: int32(698),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(35.48403760118173),
		Fy:    -libc.Float64FromFloat64(0.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	677: {
		Ffile: __ccgo_ts,
		Fline: int32(699),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(35.232723172900826),
		Fy:    -libc.Float64FromFloat64(0.9999999999999994),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	678: {
		Ffile: __ccgo_ts,
		Fline: int32(700),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(35.032052477438675),
		Fy:    -libc.Float64FromFloat64(0.9999999999999993),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	679: {
		Ffile: __ccgo_ts,
		Fline: int32(701),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(34.59673440618083),
		Fy:    -libc.Float64FromFloat64(0.999999999999999),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	680: {
		Ffile: __ccgo_ts,
		Fline: int32(702),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(34.2944535343079),
		Fy:    -libc.Float64FromFloat64(0.9999999999999988),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	681: {
		Ffile: __ccgo_ts,
		Fline: int32(703),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(33.81902983759282),
		Fy:    -libc.Float64FromFloat64(0.9999999999999979),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	682: {
		Ffile: __ccgo_ts,
		Fline: int32(704),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(33.25556048034141),
		Fy:    -libc.Float64FromFloat64(0.9999999999999963),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	683: {
		Ffile: __ccgo_ts,
		Fline: int32(705),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.89734825708379),
		Fy:    -libc.Float64FromFloat64(0.9999999999999949),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	684: {
		Ffile: __ccgo_ts,
		Fline: int32(706),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.86559955876921),
		Fy:    -libc.Float64FromFloat64(0.9999999999999947),
		Fdy:   float32(6.367041632425116e-14),
		Fe:    int32(FE_INEXACT),
	},
	685: {
		Ffile: __ccgo_ts,
		Fline: int32(707),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.7385998680079),
		Fy:    -libc.Float64FromFloat64(0.9999999999999939),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	686: {
		Ffile: __ccgo_ts,
		Fline: int32(708),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.71144887894195),
		Fy:    -libc.Float64FromFloat64(0.9999999999999937),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	687: {
		Ffile: __ccgo_ts,
		Fline: int32(709),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.50996682440892),
		Fy:    -libc.Float64FromFloat64(0.9999999999999923),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	688: {
		Ffile: __ccgo_ts,
		Fline: int32(710),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.48830532762774),
		Fy:    -libc.Float64FromFloat64(0.9999999999999921),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	689: {
		Ffile: __ccgo_ts,
		Fline: int32(711),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.474120692635786),
		Fy:    -libc.Float64FromFloat64(0.9999999999999921),
		Fdy:   float32(1.2666713652357414e-14),
		Fe:    int32(FE_INEXACT),
	},
	690: {
		Ffile: __ccgo_ts,
		Fline: int32(712),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.48118785985888),
		Fy:    -libc.Float64FromFloat64(0.9999999999999921),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	691: {
		Ffile: __ccgo_ts,
		Fline: int32(713),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.460134450661045),
		Fy:    -libc.Float64FromFloat64(0.9999999999999919),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	692: {
		Ffile: __ccgo_ts,
		Fline: int32(714),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.43951516345831),
		Fy:    -libc.Float64FromFloat64(0.9999999999999918),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	693: {
		Ffile: __ccgo_ts,
		Fline: int32(715),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.40606722939077),
		Fy:    -libc.Float64FromFloat64(0.9999999999999916),
		Fdy:   float32(8.545984422512684e-14),
		Fe:    int32(FE_INEXACT),
	},
	694: {
		Ffile: __ccgo_ts,
		Fline: int32(716),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.342351415004664),
		Fy:    -libc.Float64FromFloat64(0.999999999999991),
		Fdy:   float32(1.3184370045368765e-13),
		Fe:    int32(FE_INEXACT),
	},
	695: {
		Ffile: __ccgo_ts,
		Fline: int32(717),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.259463755198894),
		Fy:    -libc.Float64FromFloat64(0.9999999999999901),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	696: {
		Ffile: __ccgo_ts,
		Fline: int32(718),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.26516177631353),
		Fy:    -libc.Float64FromFloat64(0.9999999999999903),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	697: {
		Ffile: __ccgo_ts,
		Fline: int32(719),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.236990899346836),
		Fy:    -libc.Float64FromFloat64(0.9999999999999899),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	698: {
		Ffile: __ccgo_ts,
		Fline: int32(720),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.24256194439629),
		Fy:    -libc.Float64FromFloat64(0.9999999999999901),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	699: {
		Ffile: __ccgo_ts,
		Fline: int32(721),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.22046159739563),
		Fy:    -libc.Float64FromFloat64(0.9999999999999899),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	700: {
		Ffile: __ccgo_ts,
		Fline: int32(722),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.1882007351774),
		Fy:    -libc.Float64FromFloat64(0.9999999999999895),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	701: {
		Ffile: __ccgo_ts,
		Fline: int32(723),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.036320203884685),
		Fy:    -libc.Float64FromFloat64(0.9999999999999877),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	702: {
		Ffile: __ccgo_ts,
		Fline: int32(724),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.02277597877693),
		Fy:    -libc.Float64FromFloat64(0.9999999999999876),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	703: {
		Ffile: __ccgo_ts,
		Fline: int32(725),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(32.01384734803263),
		Fy:    -libc.Float64FromFloat64(0.9999999999999876),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	704: {
		Ffile: __ccgo_ts,
		Fline: int32(726),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(32.000602121282604),
		Fy:    -libc.Float64FromFloat64(0.9999999999999872),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	705: {
		Ffile: __ccgo_ts,
		Fline: int32(727),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.896558261509526),
		Fy:    -libc.Float64FromFloat64(0.9999999999999859),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	706: {
		Ffile: __ccgo_ts,
		Fline: int32(728),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.85399864709073),
		Fy:    -libc.Float64FromFloat64(0.9999999999999852),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	707: {
		Ffile: __ccgo_ts,
		Fline: int32(729),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.846451441455347),
		Fy:    -libc.Float64FromFloat64(0.9999999999999851),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	708: {
		Ffile: __ccgo_ts,
		Fline: int32(730),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.831525791238672),
		Fy:    -libc.Float64FromFloat64(0.999999999999985),
		Fdy:   float32(4.70569789121837e-15),
		Fe:    int32(FE_INEXACT),
	},
	709: {
		Ffile: __ccgo_ts,
		Fline: int32(731),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.80593024404971),
		Fy:    -libc.Float64FromFloat64(0.9999999999999847),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	710: {
		Ffile: __ccgo_ts,
		Fline: int32(732),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.753193947968764),
		Fy:    -libc.Float64FromFloat64(0.9999999999999837),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	711: {
		Ffile: __ccgo_ts,
		Fline: int32(733),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.709635973629634),
		Fy:    -libc.Float64FromFloat64(0.999999999999983),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	712: {
		Ffile: __ccgo_ts,
		Fline: int32(734),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.63085509577652),
		Fy:    -libc.Float64FromFloat64(0.9999999999999817),
		Fdy:   float32(8.827727231933841e-14),
		Fe:    int32(FE_INEXACT),
	},
	713: {
		Ffile: __ccgo_ts,
		Fline: int32(735),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.609864819884685),
		Fy:    -libc.Float64FromFloat64(0.9999999999999812),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	714: {
		Ffile: __ccgo_ts,
		Fline: int32(736),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.555017019385016),
		Fy:    -libc.Float64FromFloat64(0.9999999999999801),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	715: {
		Ffile: __ccgo_ts,
		Fline: int32(737),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.530050396654556),
		Fy:    -libc.Float64FromFloat64(0.9999999999999798),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	716: {
		Ffile: __ccgo_ts,
		Fline: int32(738),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.513745687629612),
		Fy:    -libc.Float64FromFloat64(0.9999999999999795),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	717: {
		Ffile: __ccgo_ts,
		Fline: int32(739),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.49241154515462),
		Fy:    -libc.Float64FromFloat64(0.999999999999979),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	718: {
		Ffile: __ccgo_ts,
		Fline: int32(740),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.463801011113354),
		Fy:    -libc.Float64FromFloat64(0.9999999999999782),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	719: {
		Ffile: __ccgo_ts,
		Fline: int32(741),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.423594590635314),
		Fy:    -libc.Float64FromFloat64(0.9999999999999775),
		Fdy:   float32(7.997074546626723e-14),
		Fe:    int32(FE_INEXACT),
	},
	720: {
		Ffile: __ccgo_ts,
		Fline: int32(742),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.39446631771229),
		Fy:    -libc.Float64FromFloat64(0.9999999999999767),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	721: {
		Ffile: __ccgo_ts,
		Fline: int32(743),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.331873468070807),
		Fy:    -libc.Float64FromFloat64(0.9999999999999754),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	722: {
		Ffile: __ccgo_ts,
		Fline: int32(744),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.32738915562348),
		Fy:    -libc.Float64FromFloat64(0.9999999999999752),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	723: {
		Ffile: __ccgo_ts,
		Fline: int32(745),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.305264359342843),
		Fy:    -libc.Float64FromFloat64(0.9999999999999747),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	724: {
		Ffile: __ccgo_ts,
		Fline: int32(746),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.279344981790768),
		Fy:    -libc.Float64FromFloat64(0.999999999999974),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	725: {
		Ffile: __ccgo_ts,
		Fline: int32(747),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.27296876465149),
		Fy:    -libc.Float64FromFloat64(0.9999999999999737),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	726: {
		Ffile: __ccgo_ts,
		Fline: int32(748),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.262431259348705),
		Fy:    -libc.Float64FromFloat64(0.9999999999999735),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	727: {
		Ffile: __ccgo_ts,
		Fline: int32(749),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.219347673212393),
		Fy:    -libc.Float64FromFloat64(0.9999999999999722),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	728: {
		Ffile: __ccgo_ts,
		Fline: int32(750),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.199466302658564),
		Fy:    -libc.Float64FromFloat64(0.9999999999999717),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	729: {
		Ffile: __ccgo_ts,
		Fline: int32(751),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.201436746645864),
		Fy:    -libc.Float64FromFloat64(0.9999999999999719),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	730: {
		Ffile: __ccgo_ts,
		Fline: int32(752),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(31.195537024518675),
		Fy:    -libc.Float64FromFloat64(0.9999999999999716),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	731: {
		Ffile: __ccgo_ts,
		Fline: int32(753),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(31.197499733686524),
		Fy:    -libc.Float64FromFloat64(0.9999999999999717),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	732: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    -libc.Float64FromFloat64(0.9996862293931839),
		Fdy:   float32(-libc.Float64FromFloat64(0.2760058343410492)),
		Fe:    int32(FE_INEXACT),
	},
	733: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fy:    libc.Float64FromFloat64(76.11053017112141),
		Fdy:   float32(-libc.Float64FromFloat64(0.02792675793170929)),
		Fe:    int32(FE_INEXACT),
	},
	734: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    -libc.Float64FromFloat64(0.9997709186615084),
		Fdy:   float32(0.10052496194839478),
		Fe:    int32(FE_INEXACT),
	},
	735: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    -libc.Float64FromFloat64(0.9985434338739069),
		Fdy:   float32(-libc.Float64FromFloat64(0.27437829971313477)),
		Fe:    int32(FE_INEXACT),
	},
	736: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fy:    libc.Float64FromFloat64(10582.558245524993),
		Fdy:   float32(0.17696762084960938),
		Fe:    int32(FE_INEXACT),
	},
	737: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619858980995045),
		Fy:    libc.Float64FromFloat64(0.9386384525571999),
		Fdy:   float32(0.007150684483349323),
		Fe:    int32(FE_INEXACT),
	},
	738: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    -libc.Float64FromFloat64(0.3340921107161975),
		Fdy:   float32(-libc.Float64FromFloat64(0.21216636896133423)),
		Fe:    int32(FE_INEXACT),
	},
	739: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5617597462207241),
		Fy:    libc.Float64FromFloat64(0.7537559518626312),
		Fdy:   float32(0.21675777435302734),
		Fe:    int32(FE_INEXACT),
	},
	740: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7741522965913037),
		Fy:    libc.Float64FromFloat64(1.1687528885129248),
		Fdy:   float32(0.4007748067378998),
		Fe:    int32(FE_INEXACT),
	},
	741: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    -libc.Float64FromFloat64(0.4927562910597158),
		Fdy:   float32(-libc.Float64FromFloat64(0.05476519837975502)),
		Fe:    int32(FE_INEXACT),
	},
	742: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	743: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	744: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.7182818284590453),
		Fdy:   float32(0.348938524723053),
		Fe:    int32(FE_INEXACT),
	},
	745: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0.6321205588285577),
		Fdy:   float32(0.11194825917482376),
		Fe:    int32(FE_INEXACT),
	},
	746: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	747: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	748: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	749: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	750: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:18:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:19:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:23:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:26:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(27036)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:27:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:29:3:
		if (*l_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:30:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:31:3:
		libc.Xfesetround(tls, (*l_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:32:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:33:3:
		y = libc.Xexpm1l(tls, (*l_l)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:34:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:36:3:
		if !(checkexcept(tls, e, (*l_l)(unsafe.Pointer(p)).Fe, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:37:4:
			libc.Xprintf(tls, __ccgo_ts+74, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, estr(tls, (*l_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+127, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:40:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:42:3:
		d = ulperrl(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:43:3:
		if !(checkulp(tls, d, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:44:4:
			if libc.Xfabsf(tls, d) < libc.Float32FromFloat32(2.5) {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:45:5:
				libc.Xprintf(tls, __ccgo_ts+136, 0)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:47:5:
				err++
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:48:4:
			libc.Xprintf(tls, __ccgo_ts+139, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*l_l)(unsafe.Pointer(p)).Fdy), float64((*l_l)(unsafe.Pointer(p)).Fdy)))
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/expm1l.c:52:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+201, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+211, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 255,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 263,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 271,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 281,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 291,
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
				v2 = __ccgo_ts + 300
			} else {
				v2 = __ccgo_ts + 24
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+302, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 300
		} else {
			v3 = __ccgo_ts + 24
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+307, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 24
	} else {
		v4 = __ccgo_ts + 312
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+314, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 317
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 320
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 323
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 326
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 329
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+332, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+340, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+345, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+388, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+437) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+445) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+457) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+469) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+481) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+490) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+24) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+490) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+496, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/expm1.h\x00\x00src/math/sanity/expm1.h\x00src/math/special/expm1.h\x00%s:%d: bad fp exception: %s expm1l(%La)=%La, want %s\x00 got %s\n\x00X \x00%s:%d: %s expm1l(%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
