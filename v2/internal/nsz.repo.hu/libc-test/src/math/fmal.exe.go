// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/fmal.exe.go src/math/fmal.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:5:21:
var t = [847]lll_l{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fx2:   libc.Float64FromFloat64(4.535662560676869),
		Fx3:   libc.Float64FromFloat64(0.6620717923376739),
		Fy:    -libc.Float64FromFloat64(35.92643043547104),
		Fdy:   float32(-libc.Float64FromFloat64(0.18241934478282928)),
		Fe:    int32(FE_INEXACT),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fx2:   -libc.Float64FromFloat64(8.88799136300345),
		Fx3:   libc.Float64FromFloat64(0.05215452675006225),
		Fy:    -libc.Float64FromFloat64(38.568299724347206),
		Fdy:   float32(0.21299950778484344),
		Fe:    int32(FE_INEXACT),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fx2:   -libc.Float64FromFloat64(2.763607337379588),
		Fx3:   libc.Float64FromFloat64(7.67640268511754),
		Fy:    libc.Float64FromFloat64(30.839393603267776),
		Fdy:   float32(-libc.Float64FromFloat64(0.48542988300323486)),
		Fe:    int32(FE_INEXACT),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fx2:   libc.Float64FromFloat64(4.567535276842744),
		Fx3:   -libc.Float64FromFloat64(0.792054511984896),
		Fy:    -libc.Float64FromFloat64(30.625704014196536),
		Fdy:   float32(0.3143090307712555),
		Fe:    int32(FE_INEXACT),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fx2:   libc.Float64FromFloat64(4.811392084359796),
		Fx3:   -libc.Float64FromFloat64(0.5587586823609152),
		Fy:    libc.Float64FromFloat64(44.02868585384228),
		Fdy:   float32(-libc.Float64FromFloat64(0.18550261855125427)),
		Fe:    int32(FE_INEXACT),
	},
	5: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	9: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	10: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	11: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	12: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	13: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	14: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	15: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	16: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	18: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(14),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	19: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(15),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	20: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(16),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	21: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(17),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	22: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(18),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	24: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(21),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	26: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(22),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	27: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(23),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(24),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	29: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(25),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	30: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(26),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	31: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(27),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	32: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	33: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(29),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	34: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(30),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	35: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(31),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	36: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(32),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	37: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(34),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	38: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	39: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(36),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	40: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(37),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	41: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	42: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	43: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(40),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	44: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	45: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	51: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	52: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740994e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740994e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.8014398509481982e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481982e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.801439850948198e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481982e+16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.2204460492503128e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.000000000000001),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	61: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	62: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.930380657631324e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	63: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	64: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(1.1102230246251573e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	65: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(7.395570986446986e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	66: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	67: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(2.220446049250314e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	68: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251575e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	69: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(9.860761315262648e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	70: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.1102230246251556e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	71: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251564e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	72: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.232595164407831e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	73: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(70),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	74: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(71),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	75: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(72),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(2.465190328815662e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	76: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.1102230246251568e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	77: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.4408920985006257e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999984),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(79),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999971),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(87),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(88),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(98),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3454935912653652),
		Fx2:   libc.Float64FromFloat64(1.7877956106220931),
		Fx3:   libc.Float64FromFloat64(1.571859389147983),
		Fy:    libc.Float64FromFloat64(3.9773269257323594),
		Fdy:   float32(-libc.Float64FromFloat64(0.49998700618743896)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.507081073074747),
		Fx2:   libc.Float64FromFloat64(1.1767029794364385),
		Fx3:   libc.Float64FromFloat64(1.6564241246154152),
		Fy:    libc.Float64FromFloat64(3.429810913554735),
		Fdy:   float32(0.49968236684799194),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(100),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4151145306345048),
		Fx2:   libc.Float64FromFloat64(1.852617872772964),
		Fx3:   libc.Float64FromFloat64(1.1245545266583719),
		Fy:    libc.Float64FromFloat64(3.74622099813258),
		Fdy:   float32(0.49986889958381653),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0367906778758385),
		Fx2:   libc.Float64FromFloat64(1.844830744728096),
		Fx3:   libc.Float64FromFloat64(1.4457733930486665),
		Fy:    libc.Float64FromFloat64(3.358476711441497),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999005198478699)),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(102),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6005432728600961),
		Fx2:   libc.Float64FromFloat64(1.0148247631930114),
		Fx3:   libc.Float64FromFloat64(1.064927127569435),
		Fy:    libc.Float64FromFloat64(2.6891980754298497),
		Fdy:   float32(0.4997788071632385),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7288061404769235),
		Fx2:   libc.Float64FromFloat64(1.5474381946254545),
		Fx3:   libc.Float64FromFloat64(1.6730020036854396),
		Fy:    libc.Float64FromFloat64(4.3482226565624496),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999816417694092)),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7427107597351077),
		Fx2:   libc.Float64FromFloat64(1.849296038642201),
		Fx3:   libc.Float64FromFloat64(1.3991238749230699),
		Fy:    libc.Float64FromFloat64(4.621911979400346),
		Fdy:   float32(0.499769002199173),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3511906677719727),
		Fx2:   libc.Float64FromFloat64(1.1793934286745371),
		Fx3:   libc.Float64FromFloat64(1.5349735130065458),
		Fy:    libc.Float64FromFloat64(3.12855890746317),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998961091041565)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.136443418649472),
		Fx2:   libc.Float64FromFloat64(1.8736700041485392),
		Fx3:   libc.Float64FromFloat64(1.1546802850525482),
		Fy:    libc.Float64FromFloat64(3.284000229988085),
		Fdy:   float32(0.4999006390571594),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6771789254072742),
		Fx2:   libc.Float64FromFloat64(1.9382297040483616),
		Fx3:   libc.Float64FromFloat64(1.5121591902639855),
		Fy:    libc.Float64FromFloat64(4.762917202492276),
		Fdy:   float32(0.4997826814651489),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3247007208186852),
		Fx2:   libc.Float64FromFloat64(1.7923295099356065),
		Fx3:   libc.Float64FromFloat64(1.8730695654036074),
		Fy:    libc.Float64FromFloat64(4.247369759159906),
		Fdy:   float32(0.49981001019477844),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3851677681273775),
		Fx2:   libc.Float64FromFloat64(1.0180411583854019),
		Fx3:   libc.Float64FromFloat64(1.2722177874065919),
		Fy:    libc.Float64FromFloat64(2.682375586629109),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999846816062927)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(110),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1848557739451757),
		Fx2:   libc.Float64FromFloat64(1.1146912830254707),
		Fx3:   libc.Float64FromFloat64(1.4282819970148823),
		Fy:    libc.Float64FromFloat64(2.7490303998739676),
		Fdy:   float32(0.4998341202735901),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(111),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0215820868427623),
		Fx2:   libc.Float64FromFloat64(1.1374585434694073),
		Fx3:   libc.Float64FromFloat64(1.355554299383184),
		Fy:    libc.Float64FromFloat64(2.5175615719177897),
		Fdy:   float32(-libc.Float64FromFloat64(0.49984169006347656)),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(112),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6527076557321632),
		Fx2:   libc.Float64FromFloat64(1.168399305179871),
		Fx3:   libc.Float64FromFloat64(1.7437347880407794),
		Fy:    libc.Float64FromFloat64(3.674757264663692),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996356666088104)),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(113),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8699885922722612),
		Fx2:   libc.Float64FromFloat64(1.3846347040049936),
		Fx3:   libc.Float64FromFloat64(1.4056845632903043),
		Fy:    libc.Float64FromFloat64(3.9949356642439215),
		Fdy:   float32(0.49997439980506897),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(114),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9841488139772177),
		Fx2:   libc.Float64FromFloat64(1.529783610792098),
		Fx3:   libc.Float64FromFloat64(1.6220561432786698),
		Fy:    libc.Float64FromFloat64(4.657374480273597),
		Fdy:   float32(0.4996892213821411),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(115),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.771627576178503),
		Fx2:   libc.Float64FromFloat64(1.881816970089244),
		Fx3:   libc.Float64FromFloat64(1.2055540563817846),
		Fy:    libc.Float64FromFloat64(4.539432893912566),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998099207878113)),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2773044213120623),
		Fx2:   libc.Float64FromFloat64(1.3979940522577083),
		Fx3:   libc.Float64FromFloat64(1.2166078778255152),
		Fy:    libc.Float64FromFloat64(3.0022718617422526),
		Fdy:   float32(0.49981433153152466),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(117),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9750724646447673),
		Fx2:   libc.Float64FromFloat64(1.7390724201323842),
		Fx3:   libc.Float64FromFloat64(1.5748621332789272),
		Fy:    libc.Float64FromFloat64(5.009656184305535),
		Fdy:   float32(-libc.Float64FromFloat64(0.49963924288749695)),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3512286149779318),
		Fx2:   libc.Float64FromFloat64(1.4346728222614198),
		Fx3:   libc.Float64FromFloat64(1.503703868528462),
		Fy:    libc.Float64FromFloat64(3.4422748390992406),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997270107269287)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(119),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3800210829977664),
		Fx2:   libc.Float64FromFloat64(1.0008685774549666),
		Fx3:   libc.Float64FromFloat64(1.8728670434743273),
		Fy:    libc.Float64FromFloat64(3.254086781672164),
		Fdy:   float32(-libc.Float64FromFloat64(0.49966543912887573)),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(120),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.828546075271343),
		Fx2:   libc.Float64FromFloat64(1.1034464547945304),
		Fx3:   libc.Float64FromFloat64(1.0224385913969178),
		Fy:    libc.Float64FromFloat64(3.0401412755835335),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992161989212036)),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(121),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1926850935939781),
		Fx2:   libc.Float64FromFloat64(1.6666455202633572),
		Fx3:   libc.Float64FromFloat64(1.5199276035567817),
		Fy:    libc.Float64FromFloat64(3.5077108718800685),
		Fdy:   float32(0.49998247623443604),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(122),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7591876656152672),
		Fx2:   libc.Float64FromFloat64(1.0403137081804417),
		Fx3:   libc.Float64FromFloat64(1.048704381024774),
		Fy:    libc.Float64FromFloat64(2.8788114248262873),
		Fdy:   float32(-libc.Float64FromFloat64(0.49970781803131104)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(123),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5101955454490188),
		Fx2:   libc.Float64FromFloat64(1.9631791369872196),
		Fx3:   libc.Float64FromFloat64(1.2709724200874426),
		Fy:    libc.Float64FromFloat64(4.23575680768399),
		Fdy:   float32(-libc.Float64FromFloat64(0.49991410970687866)),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(124),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5798734149109583),
		Fx2:   libc.Float64FromFloat64(1.7645793735126727),
		Fx3:   libc.Float64FromFloat64(1.6157975972511633),
		Fy:    libc.Float64FromFloat64(4.4036096379640695),
		Fdy:   float32(0.4999425709247589),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9009690218629107),
		Fx2:   libc.Float64FromFloat64(1.1070813638776524),
		Fx3:   libc.Float64FromFloat64(1.9059924296171635),
		Fy:    libc.Float64FromFloat64(4.010519807030321),
		Fdy:   float32(-libc.Float64FromFloat64(0.49982860684394836)),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(126),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8643964014101086e-308),
		Fx2:   libc.Float64FromFloat64(0.3985564425977608),
		Fx3:   libc.Float64FromFloat64(2.263143117703874e-308),
		Fy:    libc.Float64FromFloat64(3.803323200237475e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998026192188263)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2440091214668474e-308),
		Fx2:   libc.Float64FromFloat64(0.28230856262701254),
		Fx3:   libc.Float64FromFloat64(2.6382015318205453e-308),
		Fy:    libc.Float64FromFloat64(3.271704521423756e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49973583221435547)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(128),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.340676159708326e-308),
		Fx2:   libc.Float64FromFloat64(0.46968313261094646),
		Fx3:   libc.Float64FromFloat64(2.426421843515975e-308),
		Fy:    libc.Float64FromFloat64(3.995481087246488e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998250901699066)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(129),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.7351277891236763e-308),
		Fx2:   libc.Float64FromFloat64(0.4689706006971022),
		Fx3:   libc.Float64FromFloat64(3.5097788415977123e-308),
		Fy:    libc.Float64FromFloat64(5.261443964543483e-308),
		Fdy:   float32(0.4998818635940552),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(130),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.097320992463594e-308),
		Fx2:   libc.Float64FromFloat64(0.34826230693405064),
		Fx3:   libc.Float64FromFloat64(3.435524014779262e-308),
		Fy:    libc.Float64FromFloat64(4.514204168929897e-308),
		Fdy:   float32(0.4997435510158539),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3919970282279754e-308),
		Fx2:   libc.Float64FromFloat64(0.36158817703699403),
		Fx3:   libc.Float64FromFloat64(2.869146475882576e-308),
		Fy:    libc.Float64FromFloat64(3.7340643207974364e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999672472476959)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(132),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4670353809981666e-308),
		Fx2:   libc.Float64FromFloat64(0.33177786059632886),
		Fx3:   libc.Float64FromFloat64(2.96392593525395e-308),
		Fy:    libc.Float64FromFloat64(4.1142115165732993e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996887445449829)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.416907292701433e-308),
		Fx2:   libc.Float64FromFloat64(0.42857899707072544),
		Fx3:   libc.Float64FromFloat64(3.4967769500214586e-308),
		Fy:    libc.Float64FromFloat64(4.532612653540361e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997617304325104)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.35552039034126e-308),
		Fx2:   libc.Float64FromFloat64(0.2562922672045634),
		Fx3:   libc.Float64FromFloat64(2.344754868398631e-308),
		Fy:    libc.Float64FromFloat64(3.204748796890334e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499798059463501)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.016044492462986e-308),
		Fx2:   libc.Float64FromFloat64(0.45151647891729213),
		Fx3:   libc.Float64FromFloat64(4.336757390801355e-308),
		Fy:    libc.Float64FromFloat64(6.150067659213425e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998668432235718)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6821021859078436e-308),
		Fx2:   libc.Float64FromFloat64(0.2798750484876322),
		Fx3:   libc.Float64FromFloat64(3.7716877626876836e-308),
		Fy:    libc.Float64FromFloat64(4.802216290505057e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.258265962567917e-308),
		Fx2:   libc.Float64FromFloat64(0.49524943690783435),
		Fx3:   libc.Float64FromFloat64(3.141170011050849e-308),
		Fy:    libc.Float64FromFloat64(4.754824394308572e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999818801879883)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4332292654421257e-308),
		Fx2:   libc.Float64FromFloat64(0.3211879782928168),
		Fx3:   libc.Float64FromFloat64(2.352446916491974e-308),
		Fy:    libc.Float64FromFloat64(3.4551588832750626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998648464679718)),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(139),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.4084401076902946e-308),
		Fx2:   libc.Float64FromFloat64(0.32571038140117536),
		Fx3:   libc.Float64FromFloat64(4.0342635216289706e-308),
		Fy:    libc.Float64FromFloat64(5.470138230489015e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997462332248688)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6655261725631403e-308),
		Fx2:   libc.Float64FromFloat64(0.3684614987061577),
		Fx3:   libc.Float64FromFloat64(2.543428867796032e-308),
		Fy:    libc.Float64FromFloat64(3.8940341348852923e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998893439769745)),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.708060195971149e-308),
		Fx2:   libc.Float64FromFloat64(0.4948132989058098),
		Fx3:   libc.Float64FromFloat64(3.9971842970195486e-308),
		Fy:    libc.Float64FromFloat64(5.831981795129357e-308),
		Fdy:   float32(0.4999012053012848),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2467469201808124e-308),
		Fx2:   libc.Float64FromFloat64(0.41077268053646976),
		Fx3:   libc.Float64FromFloat64(2.846194566336063e-308),
		Fy:    libc.Float64FromFloat64(4.1798695017622626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997696876525879)),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.900149096752747e-308),
		Fx2:   libc.Float64FromFloat64(0.4818361440385291),
		Fx3:   libc.Float64FromFloat64(2.7739397646540755e-308),
		Fy:    libc.Float64FromFloat64(4.1713364225702425e-308),
		Fdy:   float32(0.49965453147888184),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(144),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2790590020478085e-308),
		Fx2:   libc.Float64FromFloat64(0.494545779910508),
		Fx3:   libc.Float64FromFloat64(2.3287382189364425e-308),
		Fy:    libc.Float64FromFloat64(3.9503830104767483e-308),
		Fdy:   float32(0.49974194169044495),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.118227690605252e-308),
		Fx2:   libc.Float64FromFloat64(0.31219030136512205),
		Fx3:   libc.Float64FromFloat64(2.418463142095881e-308),
		Fy:    libc.Float64FromFloat64(3.391943584551003e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999088048934937)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.355674709931492e-308),
		Fx2:   libc.Float64FromFloat64(0.35075065331385535),
		Fx3:   libc.Float64FromFloat64(3.300869363622636e-308),
		Fy:    libc.Float64FromFloat64(4.828625113753745e-308),
		Fdy:   float32(0.4996977150440216),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6386392473101836e-308),
		Fx2:   libc.Float64FromFloat64(0.4684098218690361),
		Fx3:   libc.Float64FromFloat64(3.0053806271059854e-308),
		Fy:    libc.Float64FromFloat64(4.241345166915196e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999159276485443)),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0712082054703013e-308),
		Fx2:   libc.Float64FromFloat64(0.48800863594675237),
		Fx3:   libc.Float64FromFloat64(2.310172278230616e-308),
		Fy:    libc.Float64FromFloat64(3.8089484052906506e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992144107818604)),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.37400617209041e-308),
		Fx2:   libc.Float64FromFloat64(0.37188700424329524),
		Fx3:   libc.Float64FromFloat64(4.240196210504512e-308),
		Fy:    libc.Float64FromFloat64(5.866832262384897e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997496008872986)),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.791294700965045e-308),
		Fx2:   libc.Float64FromFloat64(0.3902593880562056),
		Fx3:   libc.Float64FromFloat64(2.5077016231942145e-308),
		Fy:    libc.Float64FromFloat64(3.9872899731335676e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499637246131897)),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1197201035837955e-308),
		Fx2:   libc.Float64FromFloat64(0.4822890529471364),
		Fx3:   libc.Float64FromFloat64(3.825341952122656e-308),
		Fy:    libc.Float64FromFloat64(5.329948806340228e-308),
		Fdy:   float32(0.49997252225875854),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.4427874877577254e-308),
		Fx2:   libc.Float64FromFloat64(0.4954069183756685),
		Fx3:   libc.Float64FromFloat64(2.872822659697919e-308),
		Fy:    libc.Float64FromFloat64(5.073810318005952e-308),
		Fdy:   float32(0.49992209672927856),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0431807977782324e-308),
		Fx2:   libc.Float64FromFloat64(0.32859700760117405),
		Fx3:   libc.Float64FromFloat64(3.077972795118369e-308),
		Fy:    libc.Float64FromFloat64(4.07795289885765e-308),
		Fdy:   float32(0.499764621257782),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0198470689050903e-308),
		Fx2:   libc.Float64FromFloat64(0.26406302772077156),
		Fx3:   libc.Float64FromFloat64(2.4881330792781e-308),
		Fy:    libc.Float64FromFloat64(3.549626067267647e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999594986438751)),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2959596585511264e-308),
		Fx2:   libc.Float64FromFloat64(0.3776403512139897),
		Fx3:   libc.Float64FromFloat64(4.1126349789632075e-308),
		Fy:    libc.Float64FromFloat64(5.357322342005597e-308),
		Fdy:   float32(0.49985745549201965),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0238407539596875e-308),
		Fx2:   libc.Float64FromFloat64(0.414709494811271),
		Fx3:   libc.Float64FromFloat64(3.522493306241033e-308),
		Fy:    libc.Float64FromFloat64(4.776508777705388e-308),
		Fdy:   float32(0.49992141127586365),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.430991855783849e-308),
		Fx2:   libc.Float64FromFloat64(0.4548974710769909),
		Fx3:   libc.Float64FromFloat64(2.700659730809859e-308),
		Fy:    libc.Float64FromFloat64(3.806511778214693e-308),
		Fdy:   float32(0.4999784231185913),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8729480991804707e-308),
		Fx2:   libc.Float64FromFloat64(0.29311350591758767),
		Fx3:   libc.Float64FromFloat64(3.3401680603424384e-308),
		Fy:    libc.Float64FromFloat64(4.475381455930084e-308),
		Fdy:   float32(0.499789297580719),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.00207962237531e-308),
		Fx2:   libc.Float64FromFloat64(0.3021342035593944),
		Fx3:   libc.Float64FromFloat64(3.7891012620253474e-308),
		Fy:    libc.Float64FromFloat64(4.6961321977536e-308),
		Fdy:   float32(0.4997672736644745),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.750731580122295e-308),
		Fx2:   libc.Float64FromFloat64(0.40071775549369787),
		Fx3:   libc.Float64FromFloat64(3.252517654819901e-308),
		Fy:    libc.Float64FromFloat64(4.35478463957214e-308),
		Fdy:   float32(0.49982166290283203),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9912021758731145e-308),
		Fx2:   libc.Float64FromFloat64(0.46555874310141504),
		Fx3:   libc.Float64FromFloat64(2.9463566572156957e-308),
		Fy:    libc.Float64FromFloat64(4.338936982577401e-308),
		Fdy:   float32(0.49971434473991394),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.838669140247822e+307),
		Fx2:   libc.Float64FromFloat64(6.424191233707441),
		Fx3:   -libc.Float64FromFloat64(1.5592452099634468e+308),
		Fy:    libc.Float64FromFloat64(1.5492083773956314e+308),
		Fdy:   float32(0.4997844994068146),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.747665004798627e+307),
		Fx2:   libc.Float64FromFloat64(5.208765824948828),
		Fx3:   -libc.Float64FromFloat64(1.685884069462192e+308),
		Fy:    libc.Float64FromFloat64(7.870634530678682e+307),
		Fdy:   float32(0.4998374581336975),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.801614987372249e+307),
		Fx2:   libc.Float64FromFloat64(6.560281135384934),
		Fx3:   -libc.Float64FromFloat64(1.5788701442037314e+308),
		Fy:    libc.Float64FromFloat64(1.5711242779002418e+308),
		Fdy:   float32(0.49981388449668884),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.998909030732365e+307),
		Fx2:   libc.Float64FromFloat64(5.4718021401385455),
		Fx3:   -libc.Float64FromFloat64(1.4337839828307078e+308),
		Fy:    libc.Float64FromFloat64(1.301520130441218e+308),
		Fdy:   float32(0.4995970129966736),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(166),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.356524507205283e+307),
		Fx2:   libc.Float64FromFloat64(5.221732002025672),
		Fx3:   -libc.Float64FromFloat64(1.3149412007165367e+308),
		Fy:    libc.Float64FromFloat64(1.4820923431743255e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999943673610687)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.810357702135086e+307),
		Fx2:   libc.Float64FromFloat64(5.397646708150679),
		Fx3:   -libc.Float64FromFloat64(1.204117134943272e+308),
		Fy:    libc.Float64FromFloat64(1.3923440066523992e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998548924922943)),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.588316887017697e+307),
		Fx2:   libc.Float64FromFloat64(5.214681379814572),
		Fx3:   -libc.Float64FromFloat64(1.181755739589817e+308),
		Fy:    libc.Float64FromFloat64(1.7323734619336349e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997689127922058)),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.921493495074072e+307),
		Fx2:   libc.Float64FromFloat64(5.067366138955358),
		Fx3:   -libc.Float64FromFloat64(1.7736604119490367e+308),
		Fy:    libc.Float64FromFloat64(7.202405370537042e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994204342365265)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.613287579356673e+307),
		Fx2:   libc.Float64FromFloat64(4.705711329368165),
		Fx3:   -libc.Float64FromFloat64(9.004379644780656e+307),
		Fy:    libc.Float64FromFloat64(1.741013131239964e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49977052211761475)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.165326508807875e+307),
		Fx2:   libc.Float64FromFloat64(5.199316186584455),
		Fx3:   -libc.Float64FromFloat64(1.303655453861495e+308),
		Fy:    libc.Float64FromFloat64(1.3819611187623608e+308),
		Fdy:   float32(0.4996021091938019),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(172),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.108571665281885e+307),
		Fx2:   libc.Float64FromFloat64(5.333924115931674),
		Fx3:   -libc.Float64FromFloat64(1.5976450785694555e+308),
		Fy:    libc.Float64FromFloat64(1.6606206933649399e+308),
		Fdy:   float32(0.49984821677207947),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.026931649688656e+307),
		Fx2:   libc.Float64FromFloat64(4.823977588705634),
		Fx3:   -libc.Float64FromFloat64(1.156853129524485e+308),
		Fy:    libc.Float64FromFloat64(1.2681274322808266e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49989303946495056)),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(174),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.013778645030372e+307),
		Fx2:   libc.Float64FromFloat64(4.458342999119849),
		Fx3:   -libc.Float64FromFloat64(1.1987847093871429e+308),
		Fy:    libc.Float64FromFloat64(1.0365297827336333e+308),
		Fdy:   float32(0.49989405274391174),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.05548747672646e+307),
		Fx2:   libc.Float64FromFloat64(4.901075194729867),
		Fx3:   -libc.Float64FromFloat64(1.0880998722839145e+308),
		Fy:    libc.Float64FromFloat64(1.38963255466124e+308),
		Fdy:   float32(0.49962931871414185),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.569509420754374e+307),
		Fx2:   libc.Float64FromFloat64(5.065480295683743),
		Fx3:   -libc.Float64FromFloat64(1.4837555026035718e+308),
		Fy:    libc.Float64FromFloat64(8.309204905736796e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49980488419532776)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.557579228561147e+307),
		Fx2:   libc.Float64FromFloat64(4.428123891994576),
		Fx3:   -libc.Float64FromFloat64(1.74835254329516e+308),
		Fy:    libc.Float64FromFloat64(7.126123930692801e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49907225370407104)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.660579037345791e+307),
		Fx2:   libc.Float64FromFloat64(4.296516507172668),
		Fx3:   -libc.Float64FromFloat64(1.489576081148976e+308),
		Fy:    libc.Float64FromFloat64(5.128493955449336e+307),
		Fdy:   float32(0.4990600645542145),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.83791302701495e+307),
		Fx2:   libc.Float64FromFloat64(5.343009061408011),
		Fx3:   -libc.Float64FromFloat64(1.696970852225873e+308),
		Fy:    libc.Float64FromFloat64(1.4222313680794014e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997941851615906)),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.441003484470601e+307),
		Fx2:   libc.Float64FromFloat64(4.5373409725271845),
		Fx3:   -libc.Float64FromFloat64(1.1396064444589259e+308),
		Fy:    libc.Float64FromFloat64(1.3291623597162377e+308),
		Fdy:   float32(0.49984365701675415),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.801235155695051e+307),
		Fx2:   libc.Float64FromFloat64(5.063081192193199),
		Fx3:   -libc.Float64FromFloat64(1.6560392750859962e+308),
		Fy:    libc.Float64FromFloat64(7.748650665236435e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994792640209198)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.029173727473902e+307),
		Fx2:   libc.Float64FromFloat64(4.213927483311745),
		Fx3:   -libc.Float64FromFloat64(1.0933189427086426e+308),
		Fy:    libc.Float64FromFloat64(1.025938396146522e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995550811290741)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.713246459303928e+307),
		Fx2:   libc.Float64FromFloat64(4.167628920894302),
		Fx3:   -libc.Float64FromFloat64(1.6025945509751142e+308),
		Fy:    libc.Float64FromFloat64(3.6171167453465783e+307),
		Fdy:   float32(0.4981565773487091),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.993885217541583e+307),
		Fx2:   libc.Float64FromFloat64(4.81978436543766),
		Fx3:   -libc.Float64FromFloat64(1.5975361383511683e+308),
		Fy:    libc.Float64FromFloat64(1.2913872876223147e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997601807117462)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9535263638368559),
		Fx2:   libc.Float64FromFloat64(0.6453699054483613),
		Fx3:   libc.Float64FromFloat64(0.6159768642976889),
		Fy:    libc.Float64FromFloat64(0.0005996450257775657),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995484948158264)),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5812254895899245),
		Fx2:   -libc.Float64FromFloat64(0.8656947825024184),
		Fx3:   libc.Float64FromFloat64(0.5032226460046499),
		Fy:    libc.Float64FromFloat64(5.877220923862946e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.4792413115501404)),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.8867409402454873),
		Fx2:   libc.Float64FromFloat64(0.8159904749863807),
		Fx3:   libc.Float64FromFloat64(0.7235194532273546),
		Fy:    -libc.Float64FromFloat64(5.270779343044255e-05),
		Fdy:   float32(0.49455180764198303),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(188),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.8912024283674143),
		Fx2:   libc.Float64FromFloat64(0.732508220935042),
		Fx3:   libc.Float64FromFloat64(0.6513743765181268),
		Fy:    -libc.Float64FromFloat64(0.0014387287782770062),
		Fdy:   float32(0.4997765123844147),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(189),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.607226606106348),
		Fx2:   libc.Float64FromFloat64(0.5733792825297863),
		Fx3:   libc.Float64FromFloat64(0.5709429075459292),
		Fy:    libc.Float64FromFloat64(0.22277175180367423),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7669891973936926),
		Fx2:   -libc.Float64FromFloat64(0.18059667464261908),
		Fx3:   libc.Float64FromFloat64(0.14975452359555128),
		Fy:    libc.Float64FromFloat64(0.01123882505943902),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1939053833535347),
		Fx2:   -libc.Float64FromFloat64(0.843016615519224),
		Fx3:   libc.Float64FromFloat64(0.11779840359964053),
		Fy:    -libc.Float64FromFloat64(0.04566705640601396),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5692028022936819),
		Fx2:   -libc.Float64FromFloat64(0.9129964007875562),
		Fx3:   libc.Float64FromFloat64(0.11183249764296027),
		Fy:    -libc.Float64FromFloat64(0.4078476121693622),
		Fdy:   float32(0.49994081258773804),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.708688852152128),
		Fx2:   -libc.Float64FromFloat64(0.9042447353356446),
		Fx3:   libc.Float64FromFloat64(0.12170937600575682),
		Fy:    libc.Float64FromFloat64(0.7625375395553795),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997636675834656)),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(194),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6887843192514702),
		Fx2:   -libc.Float64FromFloat64(0.8945914094590282),
		Fx3:   libc.Float64FromFloat64(0.10712095229067814),
		Fy:    libc.Float64FromFloat64(0.723301487263128),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999738335609436)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9093487002930813),
		Fx2:   -libc.Float64FromFloat64(0.6353220424807806),
		Fx3:   libc.Float64FromFloat64(0.12209193072689345),
		Fy:    -libc.Float64FromFloat64(0.4556373428705502),
		Fdy:   float32(0.4999924302101135),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.8083244140324664),
		Fx2:   libc.Float64FromFloat64(0.7030885853349238),
		Fx3:   libc.Float64FromFloat64(0.1158722205072921),
		Fy:    -libc.Float64FromFloat64(0.4524514482464759),
		Fdy:   float32(0.4999793767929077),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7501997308042707),
		Fx2:   libc.Float64FromFloat64(1.7794178644284324),
		Fx3:   libc.Float64FromFloat64(0.034941880125231264),
		Fy:    libc.Float64FromFloat64(3.1492785474361837),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999973475933075)),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.100626123876368),
		Fx2:   -libc.Float64FromFloat64(1.7941957343733712),
		Fx3:   libc.Float64FromFloat64(0.028998076215368075),
		Fy:    -libc.Float64FromFloat64(1.9457406203835088),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2271453199908313),
		Fx2:   -libc.Float64FromFloat64(1.9066288212481188),
		Fx3:   libc.Float64FromFloat64(0.03643890739091297),
		Fy:    libc.Float64FromFloat64(2.376149542345177),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999966621398926)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.690189479774423),
		Fx2:   -libc.Float64FromFloat64(1.614737917080582),
		Fx3:   libc.Float64FromFloat64(0.015752263118042873),
		Fy:    -libc.Float64FromFloat64(2.7134607769244212),
		Fdy:   float32(0.49999865889549255),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.216633430136878),
		Fx2:   -libc.Float64FromFloat64(1.6298414477764052),
		Fx3:   libc.Float64FromFloat64(0.01518442879709737),
		Fy:    libc.Float64FromFloat64(1.9981040199845603),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(202),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7079042204871593),
		Fx2:   -libc.Float64FromFloat64(1.8249386066786561),
		Fx3:   libc.Float64FromFloat64(0.01542138096007668),
		Fy:    -libc.Float64FromFloat64(1.2764603608376999),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7906428884980627),
		Fx2:   libc.Float64FromFloat64(1.4457010439200353),
		Fx3:   libc.Float64FromFloat64(0.015292438404015948),
		Fy:    libc.Float64FromFloat64(1.1583256876736172),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999815225601196)),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(204),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8460153055895931),
		Fx2:   libc.Float64FromFloat64(1.4629262933341582),
		Fx3:   libc.Float64FromFloat64(0.006794043496898868),
		Fy:    libc.Float64FromFloat64(1.2444520786070474),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(205),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7682870382979763),
		Fx2:   -libc.Float64FromFloat64(1.0693842091290007),
		Fx3:   libc.Float64FromFloat64(0.006249663857703811),
		Fy:    -libc.Float64FromFloat64(1.8847285721056404),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(206),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5233173093240433),
		Fx2:   -libc.Float64FromFloat64(0.5456889357081449),
		Fx3:   libc.Float64FromFloat64(0.0011895152895775862),
		Fy:    libc.Float64FromFloat64(0.8324469165604097),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(207),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.191881851083905),
		Fx2:   libc.Float64FromFloat64(0.835413528942726),
		Fx3:   libc.Float64FromFloat64(0.001180710008124752),
		Fy:    -libc.Float64FromFloat64(0.994533513288669),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8351151527140477),
		Fx2:   libc.Float64FromFloat64(1.6137398506141194),
		Fx3:   libc.Float64FromFloat64(0.0005080982382209128),
		Fy:    libc.Float64FromFloat64(1.3481667000245756),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(209),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	213: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(210),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	214: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(211),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	215: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(212),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	216: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(213),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	217: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(214),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	218: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(215),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	219: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(216),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	220: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	221: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(218),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	222: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(219),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	223: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	224: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(221),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	225: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(222),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	226: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(223),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	227: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(224),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	228: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(225),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	229: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(226),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	230: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(227),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	231: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(228),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	232: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(229),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	233: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(230),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	234: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(231),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	235: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(232),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	236: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	237: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	238: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(235),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	239: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(236),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	240: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(237),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	241: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	242: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	243: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(240),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	244: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(242),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	245: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	246: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(244),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	247: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	248: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	249: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	250: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	251: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	252: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014402e-308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.450147717014402e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	258: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(256),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	259: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740992e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740992e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(259),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.8014398509481982e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481982e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.801439850948198e+16),
		Fy:    libc.Float64FromFloat64(1.801439850948198e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.2204460492503128e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(264),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.000000000000001),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(265),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	268: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	269: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.930380657631324e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	270: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(268),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	271: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(269),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(1.1102230246251573e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	272: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(270),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(7.395570986446986e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	273: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	274: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(272),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(2.220446049250314e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	275: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(273),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251575e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	276: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(9.860761315262648e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	277: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.1102230246251556e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	278: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251564e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	279: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.232595164407831e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	280: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(278),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	281: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(279),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	282: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(280),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(2.465190328815662e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	283: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(281),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.1102230246251568e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	284: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.4408920985006257e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(283),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(284),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999982),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(285),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(286),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(287),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(289),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(290),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(291),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(292),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(293),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(294),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(295),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(296),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(297),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(298),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(299),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(300),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(301),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(302),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(304),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(305),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3454935912653652),
		Fx2:   libc.Float64FromFloat64(1.7877956106220931),
		Fx3:   libc.Float64FromFloat64(1.571859389147983),
		Fy:    libc.Float64FromFloat64(3.9773269257323594),
		Fdy:   float32(-libc.Float64FromFloat64(0.49998700618743896)),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(307),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.507081073074747),
		Fx2:   libc.Float64FromFloat64(1.1767029794364385),
		Fx3:   libc.Float64FromFloat64(1.6564241246154152),
		Fy:    libc.Float64FromFloat64(3.4298109135547348),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003176331520081)),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(308),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4151145306345048),
		Fx2:   libc.Float64FromFloat64(1.852617872772964),
		Fx3:   libc.Float64FromFloat64(1.1245545266583719),
		Fy:    libc.Float64FromFloat64(3.7462209981325794),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001310706138611)),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(309),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0367906778758385),
		Fx2:   libc.Float64FromFloat64(1.844830744728096),
		Fx3:   libc.Float64FromFloat64(1.4457733930486665),
		Fy:    libc.Float64FromFloat64(3.358476711441497),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999005198478699)),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(310),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6005432728600961),
		Fx2:   libc.Float64FromFloat64(1.0148247631930114),
		Fx3:   libc.Float64FromFloat64(1.064927127569435),
		Fy:    libc.Float64FromFloat64(2.6891980754298492),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002211928367615)),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(311),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7288061404769235),
		Fx2:   libc.Float64FromFloat64(1.5474381946254545),
		Fx3:   libc.Float64FromFloat64(1.6730020036854396),
		Fy:    libc.Float64FromFloat64(4.3482226565624496),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999816417694092)),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(312),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7427107597351077),
		Fx2:   libc.Float64FromFloat64(1.849296038642201),
		Fx3:   libc.Float64FromFloat64(1.3991238749230699),
		Fy:    libc.Float64FromFloat64(4.621911979400345),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002310276031494)),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(313),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3511906677719727),
		Fx2:   libc.Float64FromFloat64(1.1793934286745371),
		Fx3:   libc.Float64FromFloat64(1.5349735130065458),
		Fy:    libc.Float64FromFloat64(3.12855890746317),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998961091041565)),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(314),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.136443418649472),
		Fx2:   libc.Float64FromFloat64(1.8736700041485392),
		Fx3:   libc.Float64FromFloat64(1.1546802850525482),
		Fy:    libc.Float64FromFloat64(3.2840002299880844),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000993609428406)),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(315),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6771789254072742),
		Fx2:   libc.Float64FromFloat64(1.9382297040483616),
		Fx3:   libc.Float64FromFloat64(1.5121591902639855),
		Fy:    libc.Float64FromFloat64(4.762917202492275),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002173185348511)),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(316),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3247007208186852),
		Fx2:   libc.Float64FromFloat64(1.7923295099356065),
		Fx3:   libc.Float64FromFloat64(1.8730695654036074),
		Fy:    libc.Float64FromFloat64(4.2473697591599056),
		Fdy:   float32(-libc.Float64FromFloat64(0.500190019607544)),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(317),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3851677681273775),
		Fx2:   libc.Float64FromFloat64(1.0180411583854019),
		Fx3:   libc.Float64FromFloat64(1.2722177874065919),
		Fy:    libc.Float64FromFloat64(2.682375586629109),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999846816062927)),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(318),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1848557739451757),
		Fx2:   libc.Float64FromFloat64(1.1146912830254707),
		Fx3:   libc.Float64FromFloat64(1.4282819970148823),
		Fy:    libc.Float64FromFloat64(2.749030399873967),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001658797264099)),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(319),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0215820868427623),
		Fx2:   libc.Float64FromFloat64(1.1374585434694073),
		Fx3:   libc.Float64FromFloat64(1.355554299383184),
		Fy:    libc.Float64FromFloat64(2.5175615719177897),
		Fdy:   float32(-libc.Float64FromFloat64(0.49984169006347656)),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6527076557321632),
		Fx2:   libc.Float64FromFloat64(1.168399305179871),
		Fx3:   libc.Float64FromFloat64(1.7437347880407794),
		Fy:    libc.Float64FromFloat64(3.674757264663692),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996356666088104)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(321),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8699885922722612),
		Fx2:   libc.Float64FromFloat64(1.3846347040049936),
		Fx3:   libc.Float64FromFloat64(1.4056845632903043),
		Fy:    libc.Float64FromFloat64(3.994935664243921),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000256299972534)),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(322),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9841488139772177),
		Fx2:   libc.Float64FromFloat64(1.529783610792098),
		Fx3:   libc.Float64FromFloat64(1.6220561432786698),
		Fy:    libc.Float64FromFloat64(4.657374480273596),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003107786178589)),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(323),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.771627576178503),
		Fx2:   libc.Float64FromFloat64(1.881816970089244),
		Fx3:   libc.Float64FromFloat64(1.2055540563817846),
		Fy:    libc.Float64FromFloat64(4.539432893912566),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998099207878113)),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(324),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2773044213120623),
		Fx2:   libc.Float64FromFloat64(1.3979940522577083),
		Fx3:   libc.Float64FromFloat64(1.2166078778255152),
		Fy:    libc.Float64FromFloat64(3.002271861742252),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001856684684753)),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(325),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9750724646447673),
		Fx2:   libc.Float64FromFloat64(1.7390724201323842),
		Fx3:   libc.Float64FromFloat64(1.5748621332789272),
		Fy:    libc.Float64FromFloat64(5.009656184305535),
		Fdy:   float32(-libc.Float64FromFloat64(0.49963924288749695)),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3512286149779318),
		Fx2:   libc.Float64FromFloat64(1.4346728222614198),
		Fx3:   libc.Float64FromFloat64(1.503703868528462),
		Fy:    libc.Float64FromFloat64(3.4422748390992406),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997270107269287)),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(327),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3800210829977664),
		Fx2:   libc.Float64FromFloat64(1.0008685774549666),
		Fx3:   libc.Float64FromFloat64(1.8728670434743273),
		Fy:    libc.Float64FromFloat64(3.254086781672164),
		Fdy:   float32(-libc.Float64FromFloat64(0.49966543912887573)),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(328),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.828546075271343),
		Fx2:   libc.Float64FromFloat64(1.1034464547945304),
		Fx3:   libc.Float64FromFloat64(1.0224385913969178),
		Fy:    libc.Float64FromFloat64(3.0401412755835335),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992161989212036)),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(329),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1926850935939781),
		Fx2:   libc.Float64FromFloat64(1.6666455202633572),
		Fx3:   libc.Float64FromFloat64(1.5199276035567817),
		Fy:    libc.Float64FromFloat64(3.507710871880068),
		Fdy:   float32(-libc.Float64FromFloat64(0.500017523765564)),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(330),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7591876656152672),
		Fx2:   libc.Float64FromFloat64(1.0403137081804417),
		Fx3:   libc.Float64FromFloat64(1.048704381024774),
		Fy:    libc.Float64FromFloat64(2.8788114248262873),
		Fdy:   float32(-libc.Float64FromFloat64(0.49970781803131104)),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(331),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5101955454490188),
		Fx2:   libc.Float64FromFloat64(1.9631791369872196),
		Fx3:   libc.Float64FromFloat64(1.2709724200874426),
		Fy:    libc.Float64FromFloat64(4.23575680768399),
		Fdy:   float32(-libc.Float64FromFloat64(0.49991410970687866)),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(332),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5798734149109583),
		Fx2:   libc.Float64FromFloat64(1.7645793735126727),
		Fx3:   libc.Float64FromFloat64(1.6157975972511633),
		Fy:    libc.Float64FromFloat64(4.403609637964069),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000573992729187)),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9009690218629107),
		Fx2:   libc.Float64FromFloat64(1.1070813638776524),
		Fx3:   libc.Float64FromFloat64(1.9059924296171635),
		Fy:    libc.Float64FromFloat64(4.010519807030321),
		Fdy:   float32(-libc.Float64FromFloat64(0.49982860684394836)),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(334),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.8643964014101086e-308),
		Fx2:   libc.Float64FromFloat64(0.3985564425977608),
		Fx3:   libc.Float64FromFloat64(2.263143117703874e-308),
		Fy:    libc.Float64FromFloat64(3.803323200237475e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998026192188263)),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(335),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2440091214668474e-308),
		Fx2:   libc.Float64FromFloat64(0.28230856262701254),
		Fx3:   libc.Float64FromFloat64(2.6382015318205453e-308),
		Fy:    libc.Float64FromFloat64(3.271704521423756e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49973583221435547)),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.340676159708326e-308),
		Fx2:   libc.Float64FromFloat64(0.46968313261094646),
		Fx3:   libc.Float64FromFloat64(2.426421843515975e-308),
		Fy:    libc.Float64FromFloat64(3.995481087246488e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998250901699066)),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(337),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.7351277891236763e-308),
		Fx2:   libc.Float64FromFloat64(0.4689706006971022),
		Fx3:   libc.Float64FromFloat64(3.5097788415977123e-308),
		Fy:    libc.Float64FromFloat64(5.261443964543482e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001181364059448)),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(338),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.097320992463594e-308),
		Fx2:   libc.Float64FromFloat64(0.34826230693405064),
		Fx3:   libc.Float64FromFloat64(3.435524014779262e-308),
		Fy:    libc.Float64FromFloat64(4.514204168929896e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002564787864685)),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.3919970282279754e-308),
		Fx2:   libc.Float64FromFloat64(0.36158817703699403),
		Fx3:   libc.Float64FromFloat64(2.869146475882576e-308),
		Fy:    libc.Float64FromFloat64(3.7340643207974364e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999672472476959)),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(340),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.4670353809981666e-308),
		Fx2:   libc.Float64FromFloat64(0.33177786059632886),
		Fx3:   libc.Float64FromFloat64(2.96392593525395e-308),
		Fy:    libc.Float64FromFloat64(4.1142115165732993e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996887445449829)),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(341),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.416907292701433e-308),
		Fx2:   libc.Float64FromFloat64(0.42857899707072544),
		Fx3:   libc.Float64FromFloat64(3.4967769500214586e-308),
		Fy:    libc.Float64FromFloat64(4.532612653540361e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997617304325104)),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(342),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.35552039034126e-308),
		Fx2:   libc.Float64FromFloat64(0.2562922672045634),
		Fx3:   libc.Float64FromFloat64(2.344754868398631e-308),
		Fy:    libc.Float64FromFloat64(3.204748796890334e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499798059463501)),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(343),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.016044492462986e-308),
		Fx2:   libc.Float64FromFloat64(0.45151647891729213),
		Fx3:   libc.Float64FromFloat64(4.336757390801355e-308),
		Fy:    libc.Float64FromFloat64(6.150067659213425e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998668432235718)),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(344),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6821021859078436e-308),
		Fx2:   libc.Float64FromFloat64(0.2798750484876322),
		Fx3:   libc.Float64FromFloat64(3.7716877626876836e-308),
		Fy:    libc.Float64FromFloat64(4.802216290505057e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(345),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.258265962567917e-308),
		Fx2:   libc.Float64FromFloat64(0.49524943690783435),
		Fx3:   libc.Float64FromFloat64(3.141170011050849e-308),
		Fy:    libc.Float64FromFloat64(4.754824394308572e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999818801879883)),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(346),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.4332292654421257e-308),
		Fx2:   libc.Float64FromFloat64(0.3211879782928168),
		Fx3:   libc.Float64FromFloat64(2.352446916491974e-308),
		Fy:    libc.Float64FromFloat64(3.4551588832750626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998648464679718)),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(347),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.4084401076902946e-308),
		Fx2:   libc.Float64FromFloat64(0.32571038140117536),
		Fx3:   libc.Float64FromFloat64(4.0342635216289706e-308),
		Fy:    libc.Float64FromFloat64(5.470138230489015e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997462332248688)),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(348),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6655261725631403e-308),
		Fx2:   libc.Float64FromFloat64(0.3684614987061577),
		Fx3:   libc.Float64FromFloat64(2.543428867796032e-308),
		Fy:    libc.Float64FromFloat64(3.8940341348852923e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998893439769745)),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(349),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.708060195971149e-308),
		Fx2:   libc.Float64FromFloat64(0.4948132989058098),
		Fx3:   libc.Float64FromFloat64(3.9971842970195486e-308),
		Fy:    libc.Float64FromFloat64(5.831981795129356e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000987648963928)),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(350),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.2467469201808124e-308),
		Fx2:   libc.Float64FromFloat64(0.41077268053646976),
		Fx3:   libc.Float64FromFloat64(2.846194566336063e-308),
		Fy:    libc.Float64FromFloat64(4.1798695017622626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997696876525879)),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(351),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.900149096752747e-308),
		Fx2:   libc.Float64FromFloat64(0.4818361440385291),
		Fx3:   libc.Float64FromFloat64(2.7739397646540755e-308),
		Fy:    libc.Float64FromFloat64(4.171336422570242e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003454685211182)),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(352),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.2790590020478085e-308),
		Fx2:   libc.Float64FromFloat64(0.494545779910508),
		Fx3:   libc.Float64FromFloat64(2.3287382189364425e-308),
		Fy:    libc.Float64FromFloat64(3.950383010476748e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002580285072327)),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(353),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.118227690605252e-308),
		Fx2:   libc.Float64FromFloat64(0.31219030136512205),
		Fx3:   libc.Float64FromFloat64(2.418463142095881e-308),
		Fy:    libc.Float64FromFloat64(3.391943584551003e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999088048934937)),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(354),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.355674709931492e-308),
		Fx2:   libc.Float64FromFloat64(0.35075065331385535),
		Fx3:   libc.Float64FromFloat64(3.300869363622636e-308),
		Fy:    libc.Float64FromFloat64(4.828625113753744e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500302255153656)),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(355),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6386392473101836e-308),
		Fx2:   libc.Float64FromFloat64(0.4684098218690361),
		Fx3:   libc.Float64FromFloat64(3.0053806271059854e-308),
		Fy:    libc.Float64FromFloat64(4.241345166915196e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999159276485443)),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(356),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.0712082054703013e-308),
		Fx2:   libc.Float64FromFloat64(0.48800863594675237),
		Fx3:   libc.Float64FromFloat64(2.310172278230616e-308),
		Fy:    libc.Float64FromFloat64(3.8089484052906506e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992144107818604)),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(357),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.37400617209041e-308),
		Fx2:   libc.Float64FromFloat64(0.37188700424329524),
		Fx3:   libc.Float64FromFloat64(4.240196210504512e-308),
		Fy:    libc.Float64FromFloat64(5.866832262384897e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997496008872986)),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(358),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.791294700965045e-308),
		Fx2:   libc.Float64FromFloat64(0.3902593880562056),
		Fx3:   libc.Float64FromFloat64(2.5077016231942145e-308),
		Fy:    libc.Float64FromFloat64(3.9872899731335676e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499637246131897)),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(359),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.1197201035837955e-308),
		Fx2:   libc.Float64FromFloat64(0.4822890529471364),
		Fx3:   libc.Float64FromFloat64(3.825341952122656e-308),
		Fy:    libc.Float64FromFloat64(5.329948806340227e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000274777412415)),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(360),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.4427874877577254e-308),
		Fx2:   libc.Float64FromFloat64(0.4954069183756685),
		Fx3:   libc.Float64FromFloat64(2.872822659697919e-308),
		Fy:    libc.Float64FromFloat64(5.073810318005951e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000779032707214)),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(361),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.0431807977782324e-308),
		Fx2:   libc.Float64FromFloat64(0.32859700760117405),
		Fx3:   libc.Float64FromFloat64(3.077972795118369e-308),
		Fy:    libc.Float64FromFloat64(4.0779528988576494e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500235378742218)),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(362),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.0198470689050903e-308),
		Fx2:   libc.Float64FromFloat64(0.26406302772077156),
		Fx3:   libc.Float64FromFloat64(2.4881330792781e-308),
		Fy:    libc.Float64FromFloat64(3.549626067267647e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999594986438751)),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(363),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.2959596585511264e-308),
		Fx2:   libc.Float64FromFloat64(0.3776403512139897),
		Fx3:   libc.Float64FromFloat64(4.1126349789632075e-308),
		Fy:    libc.Float64FromFloat64(5.357322342005596e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001425743103027)),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(364),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.0238407539596875e-308),
		Fx2:   libc.Float64FromFloat64(0.414709494811271),
		Fx3:   libc.Float64FromFloat64(3.522493306241033e-308),
		Fy:    libc.Float64FromFloat64(4.776508777705387e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000786185264587)),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(365),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.430991855783849e-308),
		Fx2:   libc.Float64FromFloat64(0.4548974710769909),
		Fx3:   libc.Float64FromFloat64(2.700659730809859e-308),
		Fy:    libc.Float64FromFloat64(3.8065117782146927e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000215768814087)),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(366),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.8729480991804707e-308),
		Fx2:   libc.Float64FromFloat64(0.29311350591758767),
		Fx3:   libc.Float64FromFloat64(3.3401680603424384e-308),
		Fy:    libc.Float64FromFloat64(4.475381455930083e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500210702419281)),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(367),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.00207962237531e-308),
		Fx2:   libc.Float64FromFloat64(0.3021342035593944),
		Fx3:   libc.Float64FromFloat64(3.7891012620253474e-308),
		Fy:    libc.Float64FromFloat64(4.696132197753599e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002327561378479)),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(368),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.750731580122295e-308),
		Fx2:   libc.Float64FromFloat64(0.40071775549369787),
		Fx3:   libc.Float64FromFloat64(3.252517654819901e-308),
		Fy:    libc.Float64FromFloat64(4.3547846395721397e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500178337097168)),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(369),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.9912021758731145e-308),
		Fx2:   libc.Float64FromFloat64(0.46555874310141504),
		Fx3:   libc.Float64FromFloat64(2.9463566572156957e-308),
		Fy:    libc.Float64FromFloat64(4.3389369825774004e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002856254577637)),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(370),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.838669140247822e+307),
		Fx2:   libc.Float64FromFloat64(6.424191233707441),
		Fx3:   -libc.Float64FromFloat64(1.5592452099634468e+308),
		Fy:    libc.Float64FromFloat64(1.5492083773956312e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500215470790863)),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(371),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.747665004798627e+307),
		Fx2:   libc.Float64FromFloat64(5.208765824948828),
		Fx3:   -libc.Float64FromFloat64(1.685884069462192e+308),
		Fy:    libc.Float64FromFloat64(7.870634530678681e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001625418663025)),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(372),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.801614987372249e+307),
		Fx2:   libc.Float64FromFloat64(6.560281135384934),
		Fx3:   -libc.Float64FromFloat64(1.5788701442037314e+308),
		Fy:    libc.Float64FromFloat64(1.5711242779002416e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001861453056335)),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(373),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.998909030732365e+307),
		Fx2:   libc.Float64FromFloat64(5.4718021401385455),
		Fx3:   -libc.Float64FromFloat64(1.4337839828307078e+308),
		Fy:    libc.Float64FromFloat64(1.3015201304412179e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5004029870033264)),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(374),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.356524507205283e+307),
		Fx2:   libc.Float64FromFloat64(5.221732002025672),
		Fx3:   -libc.Float64FromFloat64(1.3149412007165367e+308),
		Fy:    libc.Float64FromFloat64(1.4820923431743255e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999943673610687)),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(375),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.810357702135086e+307),
		Fx2:   libc.Float64FromFloat64(5.397646708150679),
		Fx3:   -libc.Float64FromFloat64(1.204117134943272e+308),
		Fy:    libc.Float64FromFloat64(1.3923440066523992e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998548924922943)),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(376),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.588316887017697e+307),
		Fx2:   libc.Float64FromFloat64(5.214681379814572),
		Fx3:   -libc.Float64FromFloat64(1.181755739589817e+308),
		Fy:    libc.Float64FromFloat64(1.7323734619336349e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997689127922058)),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(377),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.921493495074072e+307),
		Fx2:   libc.Float64FromFloat64(5.067366138955358),
		Fx3:   -libc.Float64FromFloat64(1.7736604119490367e+308),
		Fy:    libc.Float64FromFloat64(7.202405370537042e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994204342365265)),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(378),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.613287579356673e+307),
		Fx2:   libc.Float64FromFloat64(4.705711329368165),
		Fx3:   -libc.Float64FromFloat64(9.004379644780656e+307),
		Fy:    libc.Float64FromFloat64(1.741013131239964e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49977052211761475)),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(379),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.165326508807875e+307),
		Fx2:   libc.Float64FromFloat64(5.199316186584455),
		Fx3:   -libc.Float64FromFloat64(1.303655453861495e+308),
		Fy:    libc.Float64FromFloat64(1.3819611187623606e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003978610038757)),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.108571665281885e+307),
		Fx2:   libc.Float64FromFloat64(5.333924115931674),
		Fx3:   -libc.Float64FromFloat64(1.5976450785694555e+308),
		Fy:    libc.Float64FromFloat64(1.6606206933649397e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001518130302429)),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(381),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.026931649688656e+307),
		Fx2:   libc.Float64FromFloat64(4.823977588705634),
		Fx3:   -libc.Float64FromFloat64(1.156853129524485e+308),
		Fy:    libc.Float64FromFloat64(1.2681274322808266e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49989303946495056)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(382),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.013778645030372e+307),
		Fx2:   libc.Float64FromFloat64(4.458342999119849),
		Fx3:   -libc.Float64FromFloat64(1.1987847093871429e+308),
		Fy:    libc.Float64FromFloat64(1.0365297827336331e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001059770584106)),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.05548747672646e+307),
		Fx2:   libc.Float64FromFloat64(4.901075194729867),
		Fx3:   -libc.Float64FromFloat64(1.0880998722839145e+308),
		Fy:    libc.Float64FromFloat64(1.3896325546612397e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003706812858582)),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(384),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.569509420754374e+307),
		Fx2:   libc.Float64FromFloat64(5.065480295683743),
		Fx3:   -libc.Float64FromFloat64(1.4837555026035718e+308),
		Fy:    libc.Float64FromFloat64(8.309204905736796e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49980488419532776)),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(385),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.557579228561147e+307),
		Fx2:   libc.Float64FromFloat64(4.428123891994576),
		Fx3:   -libc.Float64FromFloat64(1.74835254329516e+308),
		Fy:    libc.Float64FromFloat64(7.126123930692801e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49907225370407104)),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(386),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.660579037345791e+307),
		Fx2:   libc.Float64FromFloat64(4.296516507172668),
		Fx3:   -libc.Float64FromFloat64(1.489576081148976e+308),
		Fy:    libc.Float64FromFloat64(5.128493955449335e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5009399056434631)),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(387),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.83791302701495e+307),
		Fx2:   libc.Float64FromFloat64(5.343009061408011),
		Fx3:   -libc.Float64FromFloat64(1.696970852225873e+308),
		Fy:    libc.Float64FromFloat64(1.4222313680794014e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997941851615906)),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(388),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.441003484470601e+307),
		Fx2:   libc.Float64FromFloat64(4.5373409725271845),
		Fx3:   -libc.Float64FromFloat64(1.1396064444589259e+308),
		Fy:    libc.Float64FromFloat64(1.3291623597162375e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001563429832458)),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(389),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.801235155695051e+307),
		Fx2:   libc.Float64FromFloat64(5.063081192193199),
		Fx3:   -libc.Float64FromFloat64(1.6560392750859962e+308),
		Fy:    libc.Float64FromFloat64(7.748650665236435e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994792640209198)),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(390),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.029173727473902e+307),
		Fx2:   libc.Float64FromFloat64(4.213927483311745),
		Fx3:   -libc.Float64FromFloat64(1.0933189427086426e+308),
		Fy:    libc.Float64FromFloat64(1.025938396146522e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995550811290741)),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(391),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.713246459303928e+307),
		Fx2:   libc.Float64FromFloat64(4.167628920894302),
		Fx3:   -libc.Float64FromFloat64(1.6025945509751142e+308),
		Fy:    libc.Float64FromFloat64(3.617116745346578e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5018433928489685)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(392),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.993885217541583e+307),
		Fx2:   libc.Float64FromFloat64(4.81978436543766),
		Fx3:   -libc.Float64FromFloat64(1.5975361383511683e+308),
		Fy:    libc.Float64FromFloat64(1.2913872876223147e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997601807117462)),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(393),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9535263638368559),
		Fx2:   libc.Float64FromFloat64(0.6453699054483613),
		Fx3:   libc.Float64FromFloat64(0.6159768642976889),
		Fy:    libc.Float64FromFloat64(0.0005996450257775657),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995484948158264)),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(394),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5812254895899245),
		Fx2:   -libc.Float64FromFloat64(0.8656947825024184),
		Fx3:   libc.Float64FromFloat64(0.5032226460046499),
		Fy:    libc.Float64FromFloat64(5.877220923862946e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.4792413115501404)),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(395),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.8867409402454873),
		Fx2:   libc.Float64FromFloat64(0.8159904749863807),
		Fx3:   libc.Float64FromFloat64(0.7235194532273546),
		Fy:    -libc.Float64FromFloat64(5.270779343044255e-05),
		Fdy:   float32(0.49455180764198303),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(396),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.8912024283674143),
		Fx2:   libc.Float64FromFloat64(0.732508220935042),
		Fx3:   libc.Float64FromFloat64(0.6513743765181268),
		Fy:    -libc.Float64FromFloat64(0.0014387287782770062),
		Fdy:   float32(0.4997765123844147),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(397),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.607226606106348),
		Fx2:   libc.Float64FromFloat64(0.5733792825297863),
		Fx3:   libc.Float64FromFloat64(0.5709429075459292),
		Fy:    libc.Float64FromFloat64(0.22277175180367423),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(398),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.7669891973936926),
		Fx2:   -libc.Float64FromFloat64(0.18059667464261908),
		Fx3:   libc.Float64FromFloat64(0.14975452359555128),
		Fy:    libc.Float64FromFloat64(0.01123882505943902),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(399),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.1939053833535347),
		Fx2:   -libc.Float64FromFloat64(0.843016615519224),
		Fx3:   libc.Float64FromFloat64(0.11779840359964053),
		Fy:    -libc.Float64FromFloat64(0.04566705640601396),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(400),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5692028022936819),
		Fx2:   -libc.Float64FromFloat64(0.9129964007875562),
		Fx3:   libc.Float64FromFloat64(0.11183249764296027),
		Fy:    -libc.Float64FromFloat64(0.4078476121693622),
		Fdy:   float32(0.49994081258773804),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(401),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.708688852152128),
		Fx2:   -libc.Float64FromFloat64(0.9042447353356446),
		Fx3:   libc.Float64FromFloat64(0.12170937600575682),
		Fy:    libc.Float64FromFloat64(0.7625375395553795),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997636675834656)),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(402),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.6887843192514702),
		Fx2:   -libc.Float64FromFloat64(0.8945914094590282),
		Fx3:   libc.Float64FromFloat64(0.10712095229067814),
		Fy:    libc.Float64FromFloat64(0.723301487263128),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999738335609436)),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(403),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9093487002930813),
		Fx2:   -libc.Float64FromFloat64(0.6353220424807806),
		Fx3:   libc.Float64FromFloat64(0.12209193072689345),
		Fy:    -libc.Float64FromFloat64(0.4556373428705502),
		Fdy:   float32(0.4999924302101135),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(404),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.8083244140324664),
		Fx2:   libc.Float64FromFloat64(0.7030885853349238),
		Fx3:   libc.Float64FromFloat64(0.1158722205072921),
		Fy:    -libc.Float64FromFloat64(0.4524514482464759),
		Fdy:   float32(0.4999793767929077),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(405),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7501997308042707),
		Fx2:   libc.Float64FromFloat64(1.7794178644284324),
		Fx3:   libc.Float64FromFloat64(0.034941880125231264),
		Fy:    libc.Float64FromFloat64(3.1492785474361837),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999973475933075)),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(406),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.100626123876368),
		Fx2:   -libc.Float64FromFloat64(1.7941957343733712),
		Fx3:   libc.Float64FromFloat64(0.028998076215368075),
		Fy:    -libc.Float64FromFloat64(1.9457406203835088),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(407),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2271453199908313),
		Fx2:   -libc.Float64FromFloat64(1.9066288212481188),
		Fx3:   libc.Float64FromFloat64(0.03643890739091297),
		Fy:    libc.Float64FromFloat64(2.376149542345177),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999966621398926)),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(408),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.690189479774423),
		Fx2:   -libc.Float64FromFloat64(1.614737917080582),
		Fx3:   libc.Float64FromFloat64(0.015752263118042873),
		Fy:    -libc.Float64FromFloat64(2.7134607769244212),
		Fdy:   float32(0.49999865889549255),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(409),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.216633430136878),
		Fx2:   -libc.Float64FromFloat64(1.6298414477764052),
		Fx3:   libc.Float64FromFloat64(0.01518442879709737),
		Fy:    libc.Float64FromFloat64(1.9981040199845603),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(410),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.7079042204871593),
		Fx2:   -libc.Float64FromFloat64(1.8249386066786561),
		Fx3:   libc.Float64FromFloat64(0.01542138096007668),
		Fy:    -libc.Float64FromFloat64(1.2764603608376999),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(411),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.7906428884980627),
		Fx2:   libc.Float64FromFloat64(1.4457010439200353),
		Fx3:   libc.Float64FromFloat64(0.015292438404015948),
		Fy:    libc.Float64FromFloat64(1.1583256876736172),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999815225601196)),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(412),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8460153055895931),
		Fx2:   libc.Float64FromFloat64(1.4629262933341582),
		Fx3:   libc.Float64FromFloat64(0.006794043496898868),
		Fy:    libc.Float64FromFloat64(1.2444520786070474),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(413),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7682870382979763),
		Fx2:   -libc.Float64FromFloat64(1.0693842091290007),
		Fx3:   libc.Float64FromFloat64(0.006249663857703811),
		Fy:    -libc.Float64FromFloat64(1.8847285721056404),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(414),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5233173093240433),
		Fx2:   -libc.Float64FromFloat64(0.5456889357081449),
		Fx3:   libc.Float64FromFloat64(0.0011895152895775862),
		Fy:    libc.Float64FromFloat64(0.8324469165604097),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(415),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.191881851083905),
		Fx2:   libc.Float64FromFloat64(0.835413528942726),
		Fx3:   libc.Float64FromFloat64(0.001180710008124752),
		Fy:    -libc.Float64FromFloat64(0.994533513288669),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(416),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8351151527140477),
		Fx2:   libc.Float64FromFloat64(1.6137398506141194),
		Fx3:   libc.Float64FromFloat64(0.0005080982382209128),
		Fy:    libc.Float64FromFloat64(1.3481667000245756),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(417),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	420: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(418),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	421: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(419),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	422: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(420),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	423: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(421),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	424: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(422),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	425: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(423),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	426: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(424),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	427: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(425),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	428: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(426),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	429: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(427),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	430: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(428),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	431: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(429),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	432: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(430),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	433: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(431),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	434: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(432),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	435: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(433),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	436: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(434),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	437: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(435),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	438: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(436),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	439: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(437),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	440: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(438),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	441: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(439),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	442: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(440),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	443: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(441),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	444: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(442),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	445: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(443),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	446: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(444),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	447: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(445),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	448: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(446),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	449: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(447),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	450: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(448),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	451: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(450),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	452: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(451),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	453: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(452),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	454: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(453),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	455: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(454),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	456: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(455),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	457: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(456),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	458: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(457),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	459: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(458),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0.9999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(459),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(460),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014402e-308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(461),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(462),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(463),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	465: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(464),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(1e-323),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	466: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(465),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740992e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(466),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740992e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(467),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.8014398509481982e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481982e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(468),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.801439850948198e+16),
		Fy:    libc.Float64FromFloat64(1.801439850948198e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(469),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(470),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.2204460492503128e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(471),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000007),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(472),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.000000000000001),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(473),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	475: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(474),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	476: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(475),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.930380657631324e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	477: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(476),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	478: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(477),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(1.1102230246251573e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	479: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(478),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(7.395570986446986e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	480: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(479),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	481: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(480),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(2.220446049250314e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	482: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(481),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251575e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	483: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(482),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(9.860761315262648e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	484: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(483),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.1102230246251556e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	485: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(484),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251564e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	486: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(485),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.232595164407831e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	487: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(486),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	488: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(487),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	489: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(488),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(2.465190328815662e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	490: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(489),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.1102230246251568e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	491: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(490),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.375)),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(491),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(492),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999982),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(493),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(494),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(495),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(496),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(497),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(498),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(499),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999978),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(500),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999973),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(501),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999997),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(502),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(503),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(504),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(505),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(506),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(507),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999998),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(508),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(509),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(510),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(511),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(512),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999993),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(513),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.999999999999999),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(514),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3454935912653652),
		Fx2:   libc.Float64FromFloat64(1.7877956106220931),
		Fx3:   libc.Float64FromFloat64(1.571859389147983),
		Fy:    libc.Float64FromFloat64(3.9773269257323594),
		Fdy:   float32(-libc.Float64FromFloat64(0.49998700618743896)),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(515),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.507081073074747),
		Fx2:   libc.Float64FromFloat64(1.1767029794364385),
		Fx3:   libc.Float64FromFloat64(1.6564241246154152),
		Fy:    libc.Float64FromFloat64(3.4298109135547348),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003176331520081)),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(516),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.4151145306345048),
		Fx2:   libc.Float64FromFloat64(1.852617872772964),
		Fx3:   libc.Float64FromFloat64(1.1245545266583719),
		Fy:    libc.Float64FromFloat64(3.7462209981325794),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001310706138611)),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(517),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0367906778758385),
		Fx2:   libc.Float64FromFloat64(1.844830744728096),
		Fx3:   libc.Float64FromFloat64(1.4457733930486665),
		Fy:    libc.Float64FromFloat64(3.358476711441497),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999005198478699)),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(518),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.6005432728600961),
		Fx2:   libc.Float64FromFloat64(1.0148247631930114),
		Fx3:   libc.Float64FromFloat64(1.064927127569435),
		Fy:    libc.Float64FromFloat64(2.6891980754298492),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002211928367615)),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(519),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7288061404769235),
		Fx2:   libc.Float64FromFloat64(1.5474381946254545),
		Fx3:   libc.Float64FromFloat64(1.6730020036854396),
		Fy:    libc.Float64FromFloat64(4.3482226565624496),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999816417694092)),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(520),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7427107597351077),
		Fx2:   libc.Float64FromFloat64(1.849296038642201),
		Fx3:   libc.Float64FromFloat64(1.3991238749230699),
		Fy:    libc.Float64FromFloat64(4.621911979400345),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002310276031494)),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(521),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3511906677719727),
		Fx2:   libc.Float64FromFloat64(1.1793934286745371),
		Fx3:   libc.Float64FromFloat64(1.5349735130065458),
		Fy:    libc.Float64FromFloat64(3.12855890746317),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998961091041565)),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(522),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.136443418649472),
		Fx2:   libc.Float64FromFloat64(1.8736700041485392),
		Fx3:   libc.Float64FromFloat64(1.1546802850525482),
		Fy:    libc.Float64FromFloat64(3.2840002299880844),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000993609428406)),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(523),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.6771789254072742),
		Fx2:   libc.Float64FromFloat64(1.9382297040483616),
		Fx3:   libc.Float64FromFloat64(1.5121591902639855),
		Fy:    libc.Float64FromFloat64(4.762917202492275),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002173185348511)),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(524),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3247007208186852),
		Fx2:   libc.Float64FromFloat64(1.7923295099356065),
		Fx3:   libc.Float64FromFloat64(1.8730695654036074),
		Fy:    libc.Float64FromFloat64(4.2473697591599056),
		Fdy:   float32(-libc.Float64FromFloat64(0.500190019607544)),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(525),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3851677681273775),
		Fx2:   libc.Float64FromFloat64(1.0180411583854019),
		Fx3:   libc.Float64FromFloat64(1.2722177874065919),
		Fy:    libc.Float64FromFloat64(2.682375586629109),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999846816062927)),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(526),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1848557739451757),
		Fx2:   libc.Float64FromFloat64(1.1146912830254707),
		Fx3:   libc.Float64FromFloat64(1.4282819970148823),
		Fy:    libc.Float64FromFloat64(2.749030399873967),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001658797264099)),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(527),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0215820868427623),
		Fx2:   libc.Float64FromFloat64(1.1374585434694073),
		Fx3:   libc.Float64FromFloat64(1.355554299383184),
		Fy:    libc.Float64FromFloat64(2.5175615719177897),
		Fdy:   float32(-libc.Float64FromFloat64(0.49984169006347656)),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(528),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.6527076557321632),
		Fx2:   libc.Float64FromFloat64(1.168399305179871),
		Fx3:   libc.Float64FromFloat64(1.7437347880407794),
		Fy:    libc.Float64FromFloat64(3.674757264663692),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996356666088104)),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(529),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.8699885922722612),
		Fx2:   libc.Float64FromFloat64(1.3846347040049936),
		Fx3:   libc.Float64FromFloat64(1.4056845632903043),
		Fy:    libc.Float64FromFloat64(3.994935664243921),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000256299972534)),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(530),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.9841488139772177),
		Fx2:   libc.Float64FromFloat64(1.529783610792098),
		Fx3:   libc.Float64FromFloat64(1.6220561432786698),
		Fy:    libc.Float64FromFloat64(4.657374480273596),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003107786178589)),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(531),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.771627576178503),
		Fx2:   libc.Float64FromFloat64(1.881816970089244),
		Fx3:   libc.Float64FromFloat64(1.2055540563817846),
		Fy:    libc.Float64FromFloat64(4.539432893912566),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998099207878113)),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(532),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.2773044213120623),
		Fx2:   libc.Float64FromFloat64(1.3979940522577083),
		Fx3:   libc.Float64FromFloat64(1.2166078778255152),
		Fy:    libc.Float64FromFloat64(3.002271861742252),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001856684684753)),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(533),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.9750724646447673),
		Fx2:   libc.Float64FromFloat64(1.7390724201323842),
		Fx3:   libc.Float64FromFloat64(1.5748621332789272),
		Fy:    libc.Float64FromFloat64(5.009656184305535),
		Fdy:   float32(-libc.Float64FromFloat64(0.49963924288749695)),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(534),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3512286149779318),
		Fx2:   libc.Float64FromFloat64(1.4346728222614198),
		Fx3:   libc.Float64FromFloat64(1.503703868528462),
		Fy:    libc.Float64FromFloat64(3.4422748390992406),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997270107269287)),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(535),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3800210829977664),
		Fx2:   libc.Float64FromFloat64(1.0008685774549666),
		Fx3:   libc.Float64FromFloat64(1.8728670434743273),
		Fy:    libc.Float64FromFloat64(3.254086781672164),
		Fdy:   float32(-libc.Float64FromFloat64(0.49966543912887573)),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(536),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.828546075271343),
		Fx2:   libc.Float64FromFloat64(1.1034464547945304),
		Fx3:   libc.Float64FromFloat64(1.0224385913969178),
		Fy:    libc.Float64FromFloat64(3.0401412755835335),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992161989212036)),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(537),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1926850935939781),
		Fx2:   libc.Float64FromFloat64(1.6666455202633572),
		Fx3:   libc.Float64FromFloat64(1.5199276035567817),
		Fy:    libc.Float64FromFloat64(3.507710871880068),
		Fdy:   float32(-libc.Float64FromFloat64(0.500017523765564)),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(538),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7591876656152672),
		Fx2:   libc.Float64FromFloat64(1.0403137081804417),
		Fx3:   libc.Float64FromFloat64(1.048704381024774),
		Fy:    libc.Float64FromFloat64(2.8788114248262873),
		Fdy:   float32(-libc.Float64FromFloat64(0.49970781803131104)),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(539),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.5101955454490188),
		Fx2:   libc.Float64FromFloat64(1.9631791369872196),
		Fx3:   libc.Float64FromFloat64(1.2709724200874426),
		Fy:    libc.Float64FromFloat64(4.23575680768399),
		Fdy:   float32(-libc.Float64FromFloat64(0.49991410970687866)),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(540),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.5798734149109583),
		Fx2:   libc.Float64FromFloat64(1.7645793735126727),
		Fx3:   libc.Float64FromFloat64(1.6157975972511633),
		Fy:    libc.Float64FromFloat64(4.403609637964069),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000573992729187)),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(541),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.9009690218629107),
		Fx2:   libc.Float64FromFloat64(1.1070813638776524),
		Fx3:   libc.Float64FromFloat64(1.9059924296171635),
		Fy:    libc.Float64FromFloat64(4.010519807030321),
		Fdy:   float32(-libc.Float64FromFloat64(0.49982860684394836)),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(542),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.8643964014101086e-308),
		Fx2:   libc.Float64FromFloat64(0.3985564425977608),
		Fx3:   libc.Float64FromFloat64(2.263143117703874e-308),
		Fy:    libc.Float64FromFloat64(3.803323200237475e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998026192188263)),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(543),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2440091214668474e-308),
		Fx2:   libc.Float64FromFloat64(0.28230856262701254),
		Fx3:   libc.Float64FromFloat64(2.6382015318205453e-308),
		Fy:    libc.Float64FromFloat64(3.271704521423756e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49973583221435547)),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(544),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.340676159708326e-308),
		Fx2:   libc.Float64FromFloat64(0.46968313261094646),
		Fx3:   libc.Float64FromFloat64(2.426421843515975e-308),
		Fy:    libc.Float64FromFloat64(3.995481087246488e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998250901699066)),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(545),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.7351277891236763e-308),
		Fx2:   libc.Float64FromFloat64(0.4689706006971022),
		Fx3:   libc.Float64FromFloat64(3.5097788415977123e-308),
		Fy:    libc.Float64FromFloat64(5.261443964543482e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001181364059448)),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(546),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.097320992463594e-308),
		Fx2:   libc.Float64FromFloat64(0.34826230693405064),
		Fx3:   libc.Float64FromFloat64(3.435524014779262e-308),
		Fy:    libc.Float64FromFloat64(4.514204168929896e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002564787864685)),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(547),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.3919970282279754e-308),
		Fx2:   libc.Float64FromFloat64(0.36158817703699403),
		Fx3:   libc.Float64FromFloat64(2.869146475882576e-308),
		Fy:    libc.Float64FromFloat64(3.7340643207974364e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999672472476959)),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(548),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.4670353809981666e-308),
		Fx2:   libc.Float64FromFloat64(0.33177786059632886),
		Fx3:   libc.Float64FromFloat64(2.96392593525395e-308),
		Fy:    libc.Float64FromFloat64(4.1142115165732993e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996887445449829)),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(549),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.416907292701433e-308),
		Fx2:   libc.Float64FromFloat64(0.42857899707072544),
		Fx3:   libc.Float64FromFloat64(3.4967769500214586e-308),
		Fy:    libc.Float64FromFloat64(4.532612653540361e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997617304325104)),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(550),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.35552039034126e-308),
		Fx2:   libc.Float64FromFloat64(0.2562922672045634),
		Fx3:   libc.Float64FromFloat64(2.344754868398631e-308),
		Fy:    libc.Float64FromFloat64(3.204748796890334e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499798059463501)),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(551),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.016044492462986e-308),
		Fx2:   libc.Float64FromFloat64(0.45151647891729213),
		Fx3:   libc.Float64FromFloat64(4.336757390801355e-308),
		Fy:    libc.Float64FromFloat64(6.150067659213425e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998668432235718)),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(552),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.6821021859078436e-308),
		Fx2:   libc.Float64FromFloat64(0.2798750484876322),
		Fx3:   libc.Float64FromFloat64(3.7716877626876836e-308),
		Fy:    libc.Float64FromFloat64(4.802216290505057e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(553),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.258265962567917e-308),
		Fx2:   libc.Float64FromFloat64(0.49524943690783435),
		Fx3:   libc.Float64FromFloat64(3.141170011050849e-308),
		Fy:    libc.Float64FromFloat64(4.754824394308572e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999818801879883)),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(554),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.4332292654421257e-308),
		Fx2:   libc.Float64FromFloat64(0.3211879782928168),
		Fx3:   libc.Float64FromFloat64(2.352446916491974e-308),
		Fy:    libc.Float64FromFloat64(3.4551588832750626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998648464679718)),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(555),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.4084401076902946e-308),
		Fx2:   libc.Float64FromFloat64(0.32571038140117536),
		Fx3:   libc.Float64FromFloat64(4.0342635216289706e-308),
		Fy:    libc.Float64FromFloat64(5.470138230489015e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997462332248688)),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(556),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.6655261725631403e-308),
		Fx2:   libc.Float64FromFloat64(0.3684614987061577),
		Fx3:   libc.Float64FromFloat64(2.543428867796032e-308),
		Fy:    libc.Float64FromFloat64(3.8940341348852923e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998893439769745)),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(557),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.708060195971149e-308),
		Fx2:   libc.Float64FromFloat64(0.4948132989058098),
		Fx3:   libc.Float64FromFloat64(3.9971842970195486e-308),
		Fy:    libc.Float64FromFloat64(5.831981795129356e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000987648963928)),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(558),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.2467469201808124e-308),
		Fx2:   libc.Float64FromFloat64(0.41077268053646976),
		Fx3:   libc.Float64FromFloat64(2.846194566336063e-308),
		Fy:    libc.Float64FromFloat64(4.1798695017622626e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997696876525879)),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(559),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.900149096752747e-308),
		Fx2:   libc.Float64FromFloat64(0.4818361440385291),
		Fx3:   libc.Float64FromFloat64(2.7739397646540755e-308),
		Fy:    libc.Float64FromFloat64(4.171336422570242e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003454685211182)),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(560),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.2790590020478085e-308),
		Fx2:   libc.Float64FromFloat64(0.494545779910508),
		Fx3:   libc.Float64FromFloat64(2.3287382189364425e-308),
		Fy:    libc.Float64FromFloat64(3.950383010476748e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002580285072327)),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(561),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.118227690605252e-308),
		Fx2:   libc.Float64FromFloat64(0.31219030136512205),
		Fx3:   libc.Float64FromFloat64(2.418463142095881e-308),
		Fy:    libc.Float64FromFloat64(3.391943584551003e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999088048934937)),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(562),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.355674709931492e-308),
		Fx2:   libc.Float64FromFloat64(0.35075065331385535),
		Fx3:   libc.Float64FromFloat64(3.300869363622636e-308),
		Fy:    libc.Float64FromFloat64(4.828625113753744e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500302255153656)),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(563),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.6386392473101836e-308),
		Fx2:   libc.Float64FromFloat64(0.4684098218690361),
		Fx3:   libc.Float64FromFloat64(3.0053806271059854e-308),
		Fy:    libc.Float64FromFloat64(4.241345166915196e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999159276485443)),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(564),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.0712082054703013e-308),
		Fx2:   libc.Float64FromFloat64(0.48800863594675237),
		Fx3:   libc.Float64FromFloat64(2.310172278230616e-308),
		Fy:    libc.Float64FromFloat64(3.8089484052906506e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49992144107818604)),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(565),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.37400617209041e-308),
		Fx2:   libc.Float64FromFloat64(0.37188700424329524),
		Fx3:   libc.Float64FromFloat64(4.240196210504512e-308),
		Fy:    libc.Float64FromFloat64(5.866832262384897e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997496008872986)),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(566),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.791294700965045e-308),
		Fx2:   libc.Float64FromFloat64(0.3902593880562056),
		Fx3:   libc.Float64FromFloat64(2.5077016231942145e-308),
		Fy:    libc.Float64FromFloat64(3.9872899731335676e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.499637246131897)),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(567),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.1197201035837955e-308),
		Fx2:   libc.Float64FromFloat64(0.4822890529471364),
		Fx3:   libc.Float64FromFloat64(3.825341952122656e-308),
		Fy:    libc.Float64FromFloat64(5.329948806340227e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000274777412415)),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(568),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.4427874877577254e-308),
		Fx2:   libc.Float64FromFloat64(0.4954069183756685),
		Fx3:   libc.Float64FromFloat64(2.872822659697919e-308),
		Fy:    libc.Float64FromFloat64(5.073810318005951e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000779032707214)),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(569),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.0431807977782324e-308),
		Fx2:   libc.Float64FromFloat64(0.32859700760117405),
		Fx3:   libc.Float64FromFloat64(3.077972795118369e-308),
		Fy:    libc.Float64FromFloat64(4.0779528988576494e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500235378742218)),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(570),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.0198470689050903e-308),
		Fx2:   libc.Float64FromFloat64(0.26406302772077156),
		Fx3:   libc.Float64FromFloat64(2.4881330792781e-308),
		Fy:    libc.Float64FromFloat64(3.549626067267647e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999594986438751)),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(571),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.2959596585511264e-308),
		Fx2:   libc.Float64FromFloat64(0.3776403512139897),
		Fx3:   libc.Float64FromFloat64(4.1126349789632075e-308),
		Fy:    libc.Float64FromFloat64(5.357322342005596e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001425743103027)),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(572),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.0238407539596875e-308),
		Fx2:   libc.Float64FromFloat64(0.414709494811271),
		Fx3:   libc.Float64FromFloat64(3.522493306241033e-308),
		Fy:    libc.Float64FromFloat64(4.776508777705387e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000786185264587)),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(573),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.430991855783849e-308),
		Fx2:   libc.Float64FromFloat64(0.4548974710769909),
		Fx3:   libc.Float64FromFloat64(2.700659730809859e-308),
		Fy:    libc.Float64FromFloat64(3.8065117782146927e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000215768814087)),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(574),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.8729480991804707e-308),
		Fx2:   libc.Float64FromFloat64(0.29311350591758767),
		Fx3:   libc.Float64FromFloat64(3.3401680603424384e-308),
		Fy:    libc.Float64FromFloat64(4.475381455930083e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500210702419281)),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(575),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.00207962237531e-308),
		Fx2:   libc.Float64FromFloat64(0.3021342035593944),
		Fx3:   libc.Float64FromFloat64(3.7891012620253474e-308),
		Fy:    libc.Float64FromFloat64(4.696132197753599e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002327561378479)),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(576),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.750731580122295e-308),
		Fx2:   libc.Float64FromFloat64(0.40071775549369787),
		Fx3:   libc.Float64FromFloat64(3.252517654819901e-308),
		Fy:    libc.Float64FromFloat64(4.3547846395721397e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500178337097168)),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(577),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.9912021758731145e-308),
		Fx2:   libc.Float64FromFloat64(0.46555874310141504),
		Fx3:   libc.Float64FromFloat64(2.9463566572156957e-308),
		Fy:    libc.Float64FromFloat64(4.3389369825774004e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002856254577637)),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(578),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.838669140247822e+307),
		Fx2:   libc.Float64FromFloat64(6.424191233707441),
		Fx3:   -libc.Float64FromFloat64(1.5592452099634468e+308),
		Fy:    libc.Float64FromFloat64(1.5492083773956312e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.500215470790863)),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(579),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.747665004798627e+307),
		Fx2:   libc.Float64FromFloat64(5.208765824948828),
		Fx3:   -libc.Float64FromFloat64(1.685884069462192e+308),
		Fy:    libc.Float64FromFloat64(7.870634530678681e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001625418663025)),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(580),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.801614987372249e+307),
		Fx2:   libc.Float64FromFloat64(6.560281135384934),
		Fx3:   -libc.Float64FromFloat64(1.5788701442037314e+308),
		Fy:    libc.Float64FromFloat64(1.5711242779002416e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001861453056335)),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(581),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.998909030732365e+307),
		Fx2:   libc.Float64FromFloat64(5.4718021401385455),
		Fx3:   -libc.Float64FromFloat64(1.4337839828307078e+308),
		Fy:    libc.Float64FromFloat64(1.3015201304412179e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5004029870033264)),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(582),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.356524507205283e+307),
		Fx2:   libc.Float64FromFloat64(5.221732002025672),
		Fx3:   -libc.Float64FromFloat64(1.3149412007165367e+308),
		Fy:    libc.Float64FromFloat64(1.4820923431743255e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999943673610687)),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(583),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.810357702135086e+307),
		Fx2:   libc.Float64FromFloat64(5.397646708150679),
		Fx3:   -libc.Float64FromFloat64(1.204117134943272e+308),
		Fy:    libc.Float64FromFloat64(1.3923440066523992e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998548924922943)),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(584),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.588316887017697e+307),
		Fx2:   libc.Float64FromFloat64(5.214681379814572),
		Fx3:   -libc.Float64FromFloat64(1.181755739589817e+308),
		Fy:    libc.Float64FromFloat64(1.7323734619336349e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997689127922058)),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(585),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.921493495074072e+307),
		Fx2:   libc.Float64FromFloat64(5.067366138955358),
		Fx3:   -libc.Float64FromFloat64(1.7736604119490367e+308),
		Fy:    libc.Float64FromFloat64(7.202405370537042e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994204342365265)),
		Fe:    int32(FE_INEXACT),
	},
	587: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(586),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.613287579356673e+307),
		Fx2:   libc.Float64FromFloat64(4.705711329368165),
		Fx3:   -libc.Float64FromFloat64(9.004379644780656e+307),
		Fy:    libc.Float64FromFloat64(1.741013131239964e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49977052211761475)),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(587),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.165326508807875e+307),
		Fx2:   libc.Float64FromFloat64(5.199316186584455),
		Fx3:   -libc.Float64FromFloat64(1.303655453861495e+308),
		Fy:    libc.Float64FromFloat64(1.3819611187623606e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003978610038757)),
		Fe:    int32(FE_INEXACT),
	},
	589: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(588),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.108571665281885e+307),
		Fx2:   libc.Float64FromFloat64(5.333924115931674),
		Fx3:   -libc.Float64FromFloat64(1.5976450785694555e+308),
		Fy:    libc.Float64FromFloat64(1.6606206933649397e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001518130302429)),
		Fe:    int32(FE_INEXACT),
	},
	590: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(589),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.026931649688656e+307),
		Fx2:   libc.Float64FromFloat64(4.823977588705634),
		Fx3:   -libc.Float64FromFloat64(1.156853129524485e+308),
		Fy:    libc.Float64FromFloat64(1.2681274322808266e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.49989303946495056)),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(590),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.013778645030372e+307),
		Fx2:   libc.Float64FromFloat64(4.458342999119849),
		Fx3:   -libc.Float64FromFloat64(1.1987847093871429e+308),
		Fy:    libc.Float64FromFloat64(1.0365297827336331e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001059770584106)),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(591),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.05548747672646e+307),
		Fx2:   libc.Float64FromFloat64(4.901075194729867),
		Fx3:   -libc.Float64FromFloat64(1.0880998722839145e+308),
		Fy:    libc.Float64FromFloat64(1.3896325546612397e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5003706812858582)),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(592),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.569509420754374e+307),
		Fx2:   libc.Float64FromFloat64(5.065480295683743),
		Fx3:   -libc.Float64FromFloat64(1.4837555026035718e+308),
		Fy:    libc.Float64FromFloat64(8.309204905736796e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49980488419532776)),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(593),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.557579228561147e+307),
		Fx2:   libc.Float64FromFloat64(4.428123891994576),
		Fx3:   -libc.Float64FromFloat64(1.74835254329516e+308),
		Fy:    libc.Float64FromFloat64(7.126123930692801e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.49907225370407104)),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(594),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.660579037345791e+307),
		Fx2:   libc.Float64FromFloat64(4.296516507172668),
		Fx3:   -libc.Float64FromFloat64(1.489576081148976e+308),
		Fy:    libc.Float64FromFloat64(5.128493955449335e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5009399056434631)),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(595),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.83791302701495e+307),
		Fx2:   libc.Float64FromFloat64(5.343009061408011),
		Fx3:   -libc.Float64FromFloat64(1.696970852225873e+308),
		Fy:    libc.Float64FromFloat64(1.4222313680794014e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997941851615906)),
		Fe:    int32(FE_INEXACT),
	},
	597: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(596),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.441003484470601e+307),
		Fx2:   libc.Float64FromFloat64(4.5373409725271845),
		Fx3:   -libc.Float64FromFloat64(1.1396064444589259e+308),
		Fy:    libc.Float64FromFloat64(1.3291623597162375e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.5001563429832458)),
		Fe:    int32(FE_INEXACT),
	},
	598: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(597),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.801235155695051e+307),
		Fx2:   libc.Float64FromFloat64(5.063081192193199),
		Fx3:   -libc.Float64FromFloat64(1.6560392750859962e+308),
		Fy:    libc.Float64FromFloat64(7.748650665236435e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4994792640209198)),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(598),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.029173727473902e+307),
		Fx2:   libc.Float64FromFloat64(4.213927483311745),
		Fx3:   -libc.Float64FromFloat64(1.0933189427086426e+308),
		Fy:    libc.Float64FromFloat64(1.025938396146522e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995550811290741)),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(599),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.713246459303928e+307),
		Fx2:   libc.Float64FromFloat64(4.167628920894302),
		Fx3:   -libc.Float64FromFloat64(1.6025945509751142e+308),
		Fy:    libc.Float64FromFloat64(3.617116745346578e+307),
		Fdy:   float32(-libc.Float64FromFloat64(0.5018433928489685)),
		Fe:    int32(FE_INEXACT),
	},
	601: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(600),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.993885217541583e+307),
		Fx2:   libc.Float64FromFloat64(4.81978436543766),
		Fx3:   -libc.Float64FromFloat64(1.5975361383511683e+308),
		Fy:    libc.Float64FromFloat64(1.2913872876223147e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997601807117462)),
		Fe:    int32(FE_INEXACT),
	},
	602: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(601),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9535263638368559),
		Fx2:   libc.Float64FromFloat64(0.6453699054483613),
		Fx3:   libc.Float64FromFloat64(0.6159768642976889),
		Fy:    libc.Float64FromFloat64(0.0005996450257775657),
		Fdy:   float32(-libc.Float64FromFloat64(0.4995484948158264)),
		Fe:    int32(FE_INEXACT),
	},
	603: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(602),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.5812254895899245),
		Fx2:   -libc.Float64FromFloat64(0.8656947825024184),
		Fx3:   libc.Float64FromFloat64(0.5032226460046499),
		Fy:    libc.Float64FromFloat64(5.877220923862946e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.4792413115501404)),
		Fe:    int32(FE_INEXACT),
	},
	604: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(603),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.8867409402454873),
		Fx2:   libc.Float64FromFloat64(0.8159904749863807),
		Fx3:   libc.Float64FromFloat64(0.7235194532273546),
		Fy:    -libc.Float64FromFloat64(5.270779343044256e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5054481625556946)),
		Fe:    int32(FE_INEXACT),
	},
	605: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(604),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.8912024283674143),
		Fx2:   libc.Float64FromFloat64(0.732508220935042),
		Fx3:   libc.Float64FromFloat64(0.6513743765181268),
		Fy:    -libc.Float64FromFloat64(0.0014387287782770064),
		Fdy:   float32(-libc.Float64FromFloat64(0.5002235174179077)),
		Fe:    int32(FE_INEXACT),
	},
	606: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(605),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.607226606106348),
		Fx2:   libc.Float64FromFloat64(0.5733792825297863),
		Fx3:   libc.Float64FromFloat64(0.5709429075459292),
		Fy:    libc.Float64FromFloat64(0.22277175180367423),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	607: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(606),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.7669891973936926),
		Fx2:   -libc.Float64FromFloat64(0.18059667464261908),
		Fx3:   libc.Float64FromFloat64(0.14975452359555128),
		Fy:    libc.Float64FromFloat64(0.01123882505943902),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	608: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(607),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.1939053833535347),
		Fx2:   -libc.Float64FromFloat64(0.843016615519224),
		Fx3:   libc.Float64FromFloat64(0.11779840359964053),
		Fy:    -libc.Float64FromFloat64(0.04566705640601397),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000004768371582)),
		Fe:    int32(FE_INEXACT),
	},
	609: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(608),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.5692028022936819),
		Fx2:   -libc.Float64FromFloat64(0.9129964007875562),
		Fx3:   libc.Float64FromFloat64(0.11183249764296027),
		Fy:    -libc.Float64FromFloat64(0.4078476121693623),
		Fdy:   float32(-libc.Float64FromFloat64(0.500059187412262)),
		Fe:    int32(FE_INEXACT),
	},
	610: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(609),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.708688852152128),
		Fx2:   -libc.Float64FromFloat64(0.9042447353356446),
		Fx3:   libc.Float64FromFloat64(0.12170937600575682),
		Fy:    libc.Float64FromFloat64(0.7625375395553795),
		Fdy:   float32(-libc.Float64FromFloat64(0.49997636675834656)),
		Fe:    int32(FE_INEXACT),
	},
	611: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(610),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.6887843192514702),
		Fx2:   -libc.Float64FromFloat64(0.8945914094590282),
		Fx3:   libc.Float64FromFloat64(0.10712095229067814),
		Fy:    libc.Float64FromFloat64(0.723301487263128),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999738335609436)),
		Fe:    int32(FE_INEXACT),
	},
	612: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(611),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9093487002930813),
		Fx2:   -libc.Float64FromFloat64(0.6353220424807806),
		Fx3:   libc.Float64FromFloat64(0.12209193072689345),
		Fy:    -libc.Float64FromFloat64(0.45563734287055024),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000075697898865)),
		Fe:    int32(FE_INEXACT),
	},
	613: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(612),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.8083244140324664),
		Fx2:   libc.Float64FromFloat64(0.7030885853349238),
		Fx3:   libc.Float64FromFloat64(0.1158722205072921),
		Fy:    -libc.Float64FromFloat64(0.45245144824647593),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000206232070923)),
		Fe:    int32(FE_INEXACT),
	},
	614: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(613),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7501997308042707),
		Fx2:   libc.Float64FromFloat64(1.7794178644284324),
		Fx3:   libc.Float64FromFloat64(0.034941880125231264),
		Fy:    libc.Float64FromFloat64(3.1492785474361837),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999973475933075)),
		Fe:    int32(FE_INEXACT),
	},
	615: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(614),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.100626123876368),
		Fx2:   -libc.Float64FromFloat64(1.7941957343733712),
		Fx3:   libc.Float64FromFloat64(0.028998076215368075),
		Fy:    -libc.Float64FromFloat64(1.945740620383509),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000000596046448)),
		Fe:    int32(FE_INEXACT),
	},
	616: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(615),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.2271453199908313),
		Fx2:   -libc.Float64FromFloat64(1.9066288212481188),
		Fx3:   libc.Float64FromFloat64(0.03643890739091297),
		Fy:    libc.Float64FromFloat64(2.376149542345177),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999966621398926)),
		Fe:    int32(FE_INEXACT),
	},
	617: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(616),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.690189479774423),
		Fx2:   -libc.Float64FromFloat64(1.614737917080582),
		Fx3:   libc.Float64FromFloat64(0.015752263118042873),
		Fy:    -libc.Float64FromFloat64(2.7134607769244217),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000013709068298)),
		Fe:    int32(FE_INEXACT),
	},
	618: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(617),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.216633430136878),
		Fx2:   -libc.Float64FromFloat64(1.6298414477764052),
		Fx3:   libc.Float64FromFloat64(0.01518442879709737),
		Fy:    libc.Float64FromFloat64(1.9981040199845603),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	619: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(618),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.7079042204871593),
		Fx2:   -libc.Float64FromFloat64(1.8249386066786561),
		Fx3:   libc.Float64FromFloat64(0.01542138096007668),
		Fy:    -libc.Float64FromFloat64(1.2764603608377),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000002384185791)),
		Fe:    int32(FE_INEXACT),
	},
	620: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(619),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.7906428884980627),
		Fx2:   libc.Float64FromFloat64(1.4457010439200353),
		Fx3:   libc.Float64FromFloat64(0.015292438404015948),
		Fy:    libc.Float64FromFloat64(1.1583256876736172),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999815225601196)),
		Fe:    int32(FE_INEXACT),
	},
	621: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(620),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.8460153055895931),
		Fx2:   libc.Float64FromFloat64(1.4629262933341582),
		Fx3:   libc.Float64FromFloat64(0.006794043496898868),
		Fy:    libc.Float64FromFloat64(1.2444520786070474),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	622: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(621),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7682870382979763),
		Fx2:   -libc.Float64FromFloat64(1.0693842091290007),
		Fx3:   libc.Float64FromFloat64(0.006249663857703811),
		Fy:    -libc.Float64FromFloat64(1.8847285721056406),
		Fdy:   float32(-libc.Float64FromFloat64(0.500000536441803)),
		Fe:    int32(FE_INEXACT),
	},
	623: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(622),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.5233173093240433),
		Fx2:   -libc.Float64FromFloat64(0.5456889357081449),
		Fx3:   libc.Float64FromFloat64(0.0011895152895775862),
		Fy:    libc.Float64FromFloat64(0.8324469165604097),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	624: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(623),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.191881851083905),
		Fx2:   libc.Float64FromFloat64(0.835413528942726),
		Fx3:   libc.Float64FromFloat64(0.001180710008124752),
		Fy:    -libc.Float64FromFloat64(0.9945335132886691),
		Fdy:   float32(-libc.Float64FromFloat64(0.5000002980232239)),
		Fe:    int32(FE_INEXACT),
	},
	625: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(624),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.8351151527140477),
		Fx2:   libc.Float64FromFloat64(1.6137398506141194),
		Fx3:   libc.Float64FromFloat64(0.0005080982382209128),
		Fy:    libc.Float64FromFloat64(1.3481667000245756),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	626: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(625),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	627: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(626),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	628: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(627),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	629: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(628),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	630: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(629),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	631: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(630),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	632: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(631),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	633: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(632),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	634: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(633),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	635: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(634),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	636: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(635),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	637: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(636),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	638: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(637),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	639: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(638),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	640: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(639),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	641: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(640),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	642: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(641),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	643: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(642),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	644: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(643),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	645: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(644),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	646: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(645),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	647: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(646),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	648: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(647),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	649: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(648),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	650: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(649),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	651: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(650),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	652: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(651),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	653: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(652),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	654: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(653),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	655: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(654),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	656: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(655),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	657: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(656),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	658: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(658),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	659: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(659),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	660: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(660),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	661: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(661),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fx3:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	662: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(662),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0),
		Fx3:   -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	663: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(663),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	664: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(664),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	665: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(665),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	666: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(666),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	667: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(667),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	668: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(668),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	669: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(669),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	670: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(670),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(4.4501477170144023e-308),
		Fy:    -libc.Float64FromFloat64(4.450147717014402e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	671: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(671),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	672: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(672),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fx3:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	673: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(673),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(1),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740994e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	674: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(674),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(9.007199254740992e+15),
		Fy:    libc.Float64FromFloat64(9.007199254740994e+15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	675: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(675),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.8014398509481982e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481984e+16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	676: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(676),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(1.801439850948198e+16),
		Fy:    libc.Float64FromFloat64(1.8014398509481982e+16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	677: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(677),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	678: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(678),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(2.2204460492503128e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000004),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	679: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(679),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000009),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	680: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(680),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000013),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    libc.Float64FromFloat64(1.0000000000000013),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	681: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(681),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	682: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(682),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	683: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(683),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.930380657631324e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	684: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(684),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	685: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(685),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(1.1102230246251573e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	686: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(686),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(7.395570986446986e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	687: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(687),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	688: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(688),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(2.220446049250314e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	689: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(689),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251575e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	690: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(690),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(9.860761315262648e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	691: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(691),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.1102230246251556e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	692: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(692),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251564e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	693: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(693),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.232595164407831e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	694: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(694),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	695: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(695),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(1.1102230246251563e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	696: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(696),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(2.465190328815662e-32),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	697: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(697),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.1102230246251568e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	698: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(698),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(4.4408920985006257e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	699: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(699),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	700: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(700),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(1.9999999999999984),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	701: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(701),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	702: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(702),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999993),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	703: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(703),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	704: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(704),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999992),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	705: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(705),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	706: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(706),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.9999999999999991),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	707: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(707),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(1.999999999999998),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	708: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(708),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999976),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	709: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(709),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fx2:   libc.Float64FromFloat64(0.999999999999999),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999971),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	710: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(710),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	711: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(711),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999999),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	712: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(712),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999999),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	713: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(713),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.9999999999999994),
		Fy:    libc.Float64FromFloat64(1.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	714: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(714),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999998),
		Fx3:   libc.Float64FromFloat64(0.999999999999999),
		Fy:    libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	715: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(715),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(2),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	716: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(716),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999996),
		Fy:    libc.Float64FromFloat64(1.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	717: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(717),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999997),
		Fx3:   libc.Float64FromFloat64(0.9999999999999991),
		Fy:    libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	718: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(718),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999997),
		Fy:    libc.Float64FromFloat64(1.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	719: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(719),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999996),
		Fx3:   libc.Float64FromFloat64(0.9999999999999992),
		Fy:    libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	720: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(720),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999998),
		Fy:    libc.Float64FromFloat64(1.9999999999999996),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	721: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(721),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   libc.Float64FromFloat64(0.9999999999999994),
		Fx3:   libc.Float64FromFloat64(0.9999999999999993),
		Fy:    libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	722: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(722),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3454935912653652),
		Fx2:   libc.Float64FromFloat64(1.7877956106220931),
		Fx3:   libc.Float64FromFloat64(1.571859389147983),
		Fy:    libc.Float64FromFloat64(3.97732692573236),
		Fdy:   float32(0.500012993812561),
		Fe:    int32(FE_INEXACT),
	},
	723: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(723),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.507081073074747),
		Fx2:   libc.Float64FromFloat64(1.1767029794364385),
		Fx3:   libc.Float64FromFloat64(1.6564241246154152),
		Fy:    libc.Float64FromFloat64(3.429810913554735),
		Fdy:   float32(0.49968236684799194),
		Fe:    int32(FE_INEXACT),
	},
	724: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(724),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.4151145306345048),
		Fx2:   libc.Float64FromFloat64(1.852617872772964),
		Fx3:   libc.Float64FromFloat64(1.1245545266583719),
		Fy:    libc.Float64FromFloat64(3.74622099813258),
		Fdy:   float32(0.49986889958381653),
		Fe:    int32(FE_INEXACT),
	},
	725: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(725),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0367906778758385),
		Fx2:   libc.Float64FromFloat64(1.844830744728096),
		Fx3:   libc.Float64FromFloat64(1.4457733930486665),
		Fy:    libc.Float64FromFloat64(3.3584767114414973),
		Fdy:   float32(0.5000994801521301),
		Fe:    int32(FE_INEXACT),
	},
	726: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(726),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.6005432728600961),
		Fx2:   libc.Float64FromFloat64(1.0148247631930114),
		Fx3:   libc.Float64FromFloat64(1.064927127569435),
		Fy:    libc.Float64FromFloat64(2.6891980754298497),
		Fdy:   float32(0.4997788071632385),
		Fe:    int32(FE_INEXACT),
	},
	727: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(727),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7288061404769235),
		Fx2:   libc.Float64FromFloat64(1.5474381946254545),
		Fx3:   libc.Float64FromFloat64(1.6730020036854396),
		Fy:    libc.Float64FromFloat64(4.34822265656245),
		Fdy:   float32(0.5000183582305908),
		Fe:    int32(FE_INEXACT),
	},
	728: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(728),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7427107597351077),
		Fx2:   libc.Float64FromFloat64(1.849296038642201),
		Fx3:   libc.Float64FromFloat64(1.3991238749230699),
		Fy:    libc.Float64FromFloat64(4.621911979400346),
		Fdy:   float32(0.499769002199173),
		Fe:    int32(FE_INEXACT),
	},
	729: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(729),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3511906677719727),
		Fx2:   libc.Float64FromFloat64(1.1793934286745371),
		Fx3:   libc.Float64FromFloat64(1.5349735130065458),
		Fy:    libc.Float64FromFloat64(3.1285589074631703),
		Fdy:   float32(0.5001038908958435),
		Fe:    int32(FE_INEXACT),
	},
	730: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(730),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.136443418649472),
		Fx2:   libc.Float64FromFloat64(1.8736700041485392),
		Fx3:   libc.Float64FromFloat64(1.1546802850525482),
		Fy:    libc.Float64FromFloat64(3.284000229988085),
		Fdy:   float32(0.4999006390571594),
		Fe:    int32(FE_INEXACT),
	},
	731: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(731),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.6771789254072742),
		Fx2:   libc.Float64FromFloat64(1.9382297040483616),
		Fx3:   libc.Float64FromFloat64(1.5121591902639855),
		Fy:    libc.Float64FromFloat64(4.762917202492276),
		Fdy:   float32(0.4997826814651489),
		Fe:    int32(FE_INEXACT),
	},
	732: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(732),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3247007208186852),
		Fx2:   libc.Float64FromFloat64(1.7923295099356065),
		Fx3:   libc.Float64FromFloat64(1.8730695654036074),
		Fy:    libc.Float64FromFloat64(4.247369759159906),
		Fdy:   float32(0.49981001019477844),
		Fe:    int32(FE_INEXACT),
	},
	733: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(733),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3851677681273775),
		Fx2:   libc.Float64FromFloat64(1.0180411583854019),
		Fx3:   libc.Float64FromFloat64(1.2722177874065919),
		Fy:    libc.Float64FromFloat64(2.6823755866291092),
		Fdy:   float32(0.5000153183937073),
		Fe:    int32(FE_INEXACT),
	},
	734: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(734),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1848557739451757),
		Fx2:   libc.Float64FromFloat64(1.1146912830254707),
		Fx3:   libc.Float64FromFloat64(1.4282819970148823),
		Fy:    libc.Float64FromFloat64(2.7490303998739676),
		Fdy:   float32(0.4998341202735901),
		Fe:    int32(FE_INEXACT),
	},
	735: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(735),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0215820868427623),
		Fx2:   libc.Float64FromFloat64(1.1374585434694073),
		Fx3:   libc.Float64FromFloat64(1.355554299383184),
		Fy:    libc.Float64FromFloat64(2.51756157191779),
		Fdy:   float32(0.5001583099365234),
		Fe:    int32(FE_INEXACT),
	},
	736: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(736),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.6527076557321632),
		Fx2:   libc.Float64FromFloat64(1.168399305179871),
		Fx3:   libc.Float64FromFloat64(1.7437347880407794),
		Fy:    libc.Float64FromFloat64(3.6747572646636923),
		Fdy:   float32(0.5003643035888672),
		Fe:    int32(FE_INEXACT),
	},
	737: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(737),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.8699885922722612),
		Fx2:   libc.Float64FromFloat64(1.3846347040049936),
		Fx3:   libc.Float64FromFloat64(1.4056845632903043),
		Fy:    libc.Float64FromFloat64(3.9949356642439215),
		Fdy:   float32(0.49997439980506897),
		Fe:    int32(FE_INEXACT),
	},
	738: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(738),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.9841488139772177),
		Fx2:   libc.Float64FromFloat64(1.529783610792098),
		Fx3:   libc.Float64FromFloat64(1.6220561432786698),
		Fy:    libc.Float64FromFloat64(4.657374480273597),
		Fdy:   float32(0.4996892213821411),
		Fe:    int32(FE_INEXACT),
	},
	739: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(739),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.771627576178503),
		Fx2:   libc.Float64FromFloat64(1.881816970089244),
		Fx3:   libc.Float64FromFloat64(1.2055540563817846),
		Fy:    libc.Float64FromFloat64(4.539432893912567),
		Fdy:   float32(0.5001900792121887),
		Fe:    int32(FE_INEXACT),
	},
	740: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(740),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.2773044213120623),
		Fx2:   libc.Float64FromFloat64(1.3979940522577083),
		Fx3:   libc.Float64FromFloat64(1.2166078778255152),
		Fy:    libc.Float64FromFloat64(3.0022718617422526),
		Fdy:   float32(0.49981433153152466),
		Fe:    int32(FE_INEXACT),
	},
	741: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(741),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.9750724646447673),
		Fx2:   libc.Float64FromFloat64(1.7390724201323842),
		Fx3:   libc.Float64FromFloat64(1.5748621332789272),
		Fy:    libc.Float64FromFloat64(5.009656184305536),
		Fdy:   float32(0.5003607869148254),
		Fe:    int32(FE_INEXACT),
	},
	742: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(742),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3512286149779318),
		Fx2:   libc.Float64FromFloat64(1.4346728222614198),
		Fx3:   libc.Float64FromFloat64(1.503703868528462),
		Fy:    libc.Float64FromFloat64(3.442274839099241),
		Fdy:   float32(0.5000272989273071),
		Fe:    int32(FE_INEXACT),
	},
	743: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(743),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3800210829977664),
		Fx2:   libc.Float64FromFloat64(1.0008685774549666),
		Fx3:   libc.Float64FromFloat64(1.8728670434743273),
		Fy:    libc.Float64FromFloat64(3.2540867816721644),
		Fdy:   float32(0.5003345608711243),
		Fe:    int32(FE_INEXACT),
	},
	744: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(744),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.828546075271343),
		Fx2:   libc.Float64FromFloat64(1.1034464547945304),
		Fx3:   libc.Float64FromFloat64(1.0224385913969178),
		Fy:    libc.Float64FromFloat64(3.040141275583534),
		Fdy:   float32(0.5000783801078796),
		Fe:    int32(FE_INEXACT),
	},
	745: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(745),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1926850935939781),
		Fx2:   libc.Float64FromFloat64(1.6666455202633572),
		Fx3:   libc.Float64FromFloat64(1.5199276035567817),
		Fy:    libc.Float64FromFloat64(3.5077108718800685),
		Fdy:   float32(0.49998247623443604),
		Fe:    int32(FE_INEXACT),
	},
	746: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(746),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7591876656152672),
		Fx2:   libc.Float64FromFloat64(1.0403137081804417),
		Fx3:   libc.Float64FromFloat64(1.048704381024774),
		Fy:    libc.Float64FromFloat64(2.8788114248262877),
		Fdy:   float32(0.500292181968689),
		Fe:    int32(FE_INEXACT),
	},
	747: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(747),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.5101955454490188),
		Fx2:   libc.Float64FromFloat64(1.9631791369872196),
		Fx3:   libc.Float64FromFloat64(1.2709724200874426),
		Fy:    libc.Float64FromFloat64(4.235756807683991),
		Fdy:   float32(0.5000858902931213),
		Fe:    int32(FE_INEXACT),
	},
	748: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(748),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.5798734149109583),
		Fx2:   libc.Float64FromFloat64(1.7645793735126727),
		Fx3:   libc.Float64FromFloat64(1.6157975972511633),
		Fy:    libc.Float64FromFloat64(4.4036096379640695),
		Fdy:   float32(0.4999425709247589),
		Fe:    int32(FE_INEXACT),
	},
	749: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(749),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.9009690218629107),
		Fx2:   libc.Float64FromFloat64(1.1070813638776524),
		Fx3:   libc.Float64FromFloat64(1.9059924296171635),
		Fy:    libc.Float64FromFloat64(4.010519807030322),
		Fdy:   float32(0.500171422958374),
		Fe:    int32(FE_INEXACT),
	},
	750: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(750),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.8643964014101086e-308),
		Fx2:   libc.Float64FromFloat64(0.3985564425977608),
		Fx3:   libc.Float64FromFloat64(2.263143117703874e-308),
		Fy:    libc.Float64FromFloat64(3.8033232002374756e-308),
		Fdy:   float32(0.5001973509788513),
		Fe:    int32(FE_INEXACT),
	},
	751: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(751),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2440091214668474e-308),
		Fx2:   libc.Float64FromFloat64(0.28230856262701254),
		Fx3:   libc.Float64FromFloat64(2.6382015318205453e-308),
		Fy:    libc.Float64FromFloat64(3.2717045214237564e-308),
		Fdy:   float32(0.5002641677856445),
		Fe:    int32(FE_INEXACT),
	},
	752: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(752),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.340676159708326e-308),
		Fx2:   libc.Float64FromFloat64(0.46968313261094646),
		Fx3:   libc.Float64FromFloat64(2.426421843515975e-308),
		Fy:    libc.Float64FromFloat64(3.9954810872464885e-308),
		Fdy:   float32(0.5001749396324158),
		Fe:    int32(FE_INEXACT),
	},
	753: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(753),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.7351277891236763e-308),
		Fx2:   libc.Float64FromFloat64(0.4689706006971022),
		Fx3:   libc.Float64FromFloat64(3.5097788415977123e-308),
		Fy:    libc.Float64FromFloat64(5.261443964543483e-308),
		Fdy:   float32(0.4998818635940552),
		Fe:    int32(FE_INEXACT),
	},
	754: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(754),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.097320992463594e-308),
		Fx2:   libc.Float64FromFloat64(0.34826230693405064),
		Fx3:   libc.Float64FromFloat64(3.435524014779262e-308),
		Fy:    libc.Float64FromFloat64(4.514204168929897e-308),
		Fdy:   float32(0.4997435510158539),
		Fe:    int32(FE_INEXACT),
	},
	755: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(755),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.3919970282279754e-308),
		Fx2:   libc.Float64FromFloat64(0.36158817703699403),
		Fx3:   libc.Float64FromFloat64(2.869146475882576e-308),
		Fy:    libc.Float64FromFloat64(3.734064320797437e-308),
		Fdy:   float32(0.5000327229499817),
		Fe:    int32(FE_INEXACT),
	},
	756: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(756),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.4670353809981666e-308),
		Fx2:   libc.Float64FromFloat64(0.33177786059632886),
		Fx3:   libc.Float64FromFloat64(2.96392593525395e-308),
		Fy:    libc.Float64FromFloat64(4.1142115165733e-308),
		Fdy:   float32(0.5003112554550171),
		Fe:    int32(FE_INEXACT),
	},
	757: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(757),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.416907292701433e-308),
		Fx2:   libc.Float64FromFloat64(0.42857899707072544),
		Fx3:   libc.Float64FromFloat64(3.4967769500214586e-308),
		Fy:    libc.Float64FromFloat64(4.532612653540362e-308),
		Fdy:   float32(0.5002382397651672),
		Fe:    int32(FE_INEXACT),
	},
	758: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(758),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.35552039034126e-308),
		Fx2:   libc.Float64FromFloat64(0.2562922672045634),
		Fx3:   libc.Float64FromFloat64(2.344754868398631e-308),
		Fy:    libc.Float64FromFloat64(3.2047487968903343e-308),
		Fdy:   float32(0.500201940536499),
		Fe:    int32(FE_INEXACT),
	},
	759: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(759),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.016044492462986e-308),
		Fx2:   libc.Float64FromFloat64(0.45151647891729213),
		Fx3:   libc.Float64FromFloat64(4.336757390801355e-308),
		Fy:    libc.Float64FromFloat64(6.150067659213426e-308),
		Fdy:   float32(0.5001331567764282),
		Fe:    int32(FE_INEXACT),
	},
	760: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(760),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.6821021859078436e-308),
		Fx2:   libc.Float64FromFloat64(0.2798750484876322),
		Fx3:   libc.Float64FromFloat64(3.7716877626876836e-308),
		Fy:    libc.Float64FromFloat64(4.802216290505058e-308),
		Fdy:   float32(0.5000009536743164),
		Fe:    int32(FE_INEXACT),
	},
	761: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(761),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.258265962567917e-308),
		Fx2:   libc.Float64FromFloat64(0.49524943690783435),
		Fx3:   libc.Float64FromFloat64(3.141170011050849e-308),
		Fy:    libc.Float64FromFloat64(4.754824394308573e-308),
		Fdy:   float32(0.5000181198120117),
		Fe:    int32(FE_INEXACT),
	},
	762: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(762),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.4332292654421257e-308),
		Fx2:   libc.Float64FromFloat64(0.3211879782928168),
		Fx3:   libc.Float64FromFloat64(2.352446916491974e-308),
		Fy:    libc.Float64FromFloat64(3.455158883275063e-308),
		Fdy:   float32(0.5001351833343506),
		Fe:    int32(FE_INEXACT),
	},
	763: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(763),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.4084401076902946e-308),
		Fx2:   libc.Float64FromFloat64(0.32571038140117536),
		Fx3:   libc.Float64FromFloat64(4.0342635216289706e-308),
		Fy:    libc.Float64FromFloat64(5.470138230489016e-308),
		Fdy:   float32(0.5002537369728088),
		Fe:    int32(FE_INEXACT),
	},
	764: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(764),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.6655261725631403e-308),
		Fx2:   libc.Float64FromFloat64(0.3684614987061577),
		Fx3:   libc.Float64FromFloat64(2.543428867796032e-308),
		Fy:    libc.Float64FromFloat64(3.894034134885293e-308),
		Fdy:   float32(0.5001106858253479),
		Fe:    int32(FE_INEXACT),
	},
	765: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(765),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.708060195971149e-308),
		Fx2:   libc.Float64FromFloat64(0.4948132989058098),
		Fx3:   libc.Float64FromFloat64(3.9971842970195486e-308),
		Fy:    libc.Float64FromFloat64(5.831981795129357e-308),
		Fdy:   float32(0.4999012053012848),
		Fe:    int32(FE_INEXACT),
	},
	766: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(766),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.2467469201808124e-308),
		Fx2:   libc.Float64FromFloat64(0.41077268053646976),
		Fx3:   libc.Float64FromFloat64(2.846194566336063e-308),
		Fy:    libc.Float64FromFloat64(4.179869501762263e-308),
		Fdy:   float32(0.5002303123474121),
		Fe:    int32(FE_INEXACT),
	},
	767: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(767),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.900149096752747e-308),
		Fx2:   libc.Float64FromFloat64(0.4818361440385291),
		Fx3:   libc.Float64FromFloat64(2.7739397646540755e-308),
		Fy:    libc.Float64FromFloat64(4.1713364225702425e-308),
		Fdy:   float32(0.49965453147888184),
		Fe:    int32(FE_INEXACT),
	},
	768: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(768),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.2790590020478085e-308),
		Fx2:   libc.Float64FromFloat64(0.494545779910508),
		Fx3:   libc.Float64FromFloat64(2.3287382189364425e-308),
		Fy:    libc.Float64FromFloat64(3.9503830104767483e-308),
		Fdy:   float32(0.49974194169044495),
		Fe:    int32(FE_INEXACT),
	},
	769: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(769),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.118227690605252e-308),
		Fx2:   libc.Float64FromFloat64(0.31219030136512205),
		Fx3:   libc.Float64FromFloat64(2.418463142095881e-308),
		Fy:    libc.Float64FromFloat64(3.3919435845510036e-308),
		Fdy:   float32(0.5000091195106506),
		Fe:    int32(FE_INEXACT),
	},
	770: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(770),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.355674709931492e-308),
		Fx2:   libc.Float64FromFloat64(0.35075065331385535),
		Fx3:   libc.Float64FromFloat64(3.300869363622636e-308),
		Fy:    libc.Float64FromFloat64(4.828625113753745e-308),
		Fdy:   float32(0.4996977150440216),
		Fe:    int32(FE_INEXACT),
	},
	771: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(771),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.6386392473101836e-308),
		Fx2:   libc.Float64FromFloat64(0.4684098218690361),
		Fx3:   libc.Float64FromFloat64(3.0053806271059854e-308),
		Fy:    libc.Float64FromFloat64(4.2413451669151963e-308),
		Fdy:   float32(0.5000841021537781),
		Fe:    int32(FE_INEXACT),
	},
	772: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(772),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.0712082054703013e-308),
		Fx2:   libc.Float64FromFloat64(0.48800863594675237),
		Fx3:   libc.Float64FromFloat64(2.310172278230616e-308),
		Fy:    libc.Float64FromFloat64(3.808948405290651e-308),
		Fdy:   float32(0.500078558921814),
		Fe:    int32(FE_INEXACT),
	},
	773: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(773),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.37400617209041e-308),
		Fx2:   libc.Float64FromFloat64(0.37188700424329524),
		Fx3:   libc.Float64FromFloat64(4.240196210504512e-308),
		Fy:    libc.Float64FromFloat64(5.866832262384898e-308),
		Fdy:   float32(0.5002503991127014),
		Fe:    int32(FE_INEXACT),
	},
	774: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(774),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.791294700965045e-308),
		Fx2:   libc.Float64FromFloat64(0.3902593880562056),
		Fx3:   libc.Float64FromFloat64(2.5077016231942145e-308),
		Fy:    libc.Float64FromFloat64(3.987289973133568e-308),
		Fdy:   float32(0.500362753868103),
		Fe:    int32(FE_INEXACT),
	},
	775: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(775),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.1197201035837955e-308),
		Fx2:   libc.Float64FromFloat64(0.4822890529471364),
		Fx3:   libc.Float64FromFloat64(3.825341952122656e-308),
		Fy:    libc.Float64FromFloat64(5.329948806340228e-308),
		Fdy:   float32(0.49997252225875854),
		Fe:    int32(FE_INEXACT),
	},
	776: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(776),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.4427874877577254e-308),
		Fx2:   libc.Float64FromFloat64(0.4954069183756685),
		Fx3:   libc.Float64FromFloat64(2.872822659697919e-308),
		Fy:    libc.Float64FromFloat64(5.073810318005952e-308),
		Fdy:   float32(0.49992209672927856),
		Fe:    int32(FE_INEXACT),
	},
	777: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(777),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.0431807977782324e-308),
		Fx2:   libc.Float64FromFloat64(0.32859700760117405),
		Fx3:   libc.Float64FromFloat64(3.077972795118369e-308),
		Fy:    libc.Float64FromFloat64(4.07795289885765e-308),
		Fdy:   float32(0.499764621257782),
		Fe:    int32(FE_INEXACT),
	},
	778: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(778),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.0198470689050903e-308),
		Fx2:   libc.Float64FromFloat64(0.26406302772077156),
		Fx3:   libc.Float64FromFloat64(2.4881330792781e-308),
		Fy:    libc.Float64FromFloat64(3.5496260672676476e-308),
		Fdy:   float32(0.5000404715538025),
		Fe:    int32(FE_INEXACT),
	},
	779: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(779),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.2959596585511264e-308),
		Fx2:   libc.Float64FromFloat64(0.3776403512139897),
		Fx3:   libc.Float64FromFloat64(4.1126349789632075e-308),
		Fy:    libc.Float64FromFloat64(5.357322342005597e-308),
		Fdy:   float32(0.49985745549201965),
		Fe:    int32(FE_INEXACT),
	},
	780: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(780),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.0238407539596875e-308),
		Fx2:   libc.Float64FromFloat64(0.414709494811271),
		Fx3:   libc.Float64FromFloat64(3.522493306241033e-308),
		Fy:    libc.Float64FromFloat64(4.776508777705388e-308),
		Fdy:   float32(0.49992141127586365),
		Fe:    int32(FE_INEXACT),
	},
	781: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(781),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.430991855783849e-308),
		Fx2:   libc.Float64FromFloat64(0.4548974710769909),
		Fx3:   libc.Float64FromFloat64(2.700659730809859e-308),
		Fy:    libc.Float64FromFloat64(3.806511778214693e-308),
		Fdy:   float32(0.4999784231185913),
		Fe:    int32(FE_INEXACT),
	},
	782: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(782),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.8729480991804707e-308),
		Fx2:   libc.Float64FromFloat64(0.29311350591758767),
		Fx3:   libc.Float64FromFloat64(3.3401680603424384e-308),
		Fy:    libc.Float64FromFloat64(4.475381455930084e-308),
		Fdy:   float32(0.499789297580719),
		Fe:    int32(FE_INEXACT),
	},
	783: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(783),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.00207962237531e-308),
		Fx2:   libc.Float64FromFloat64(0.3021342035593944),
		Fx3:   libc.Float64FromFloat64(3.7891012620253474e-308),
		Fy:    libc.Float64FromFloat64(4.6961321977536e-308),
		Fdy:   float32(0.4997672736644745),
		Fe:    int32(FE_INEXACT),
	},
	784: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(784),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.750731580122295e-308),
		Fx2:   libc.Float64FromFloat64(0.40071775549369787),
		Fx3:   libc.Float64FromFloat64(3.252517654819901e-308),
		Fy:    libc.Float64FromFloat64(4.35478463957214e-308),
		Fdy:   float32(0.49982166290283203),
		Fe:    int32(FE_INEXACT),
	},
	785: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(785),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.9912021758731145e-308),
		Fx2:   libc.Float64FromFloat64(0.46555874310141504),
		Fx3:   libc.Float64FromFloat64(2.9463566572156957e-308),
		Fy:    libc.Float64FromFloat64(4.338936982577401e-308),
		Fdy:   float32(0.49971434473991394),
		Fe:    int32(FE_INEXACT),
	},
	786: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(786),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.838669140247822e+307),
		Fx2:   libc.Float64FromFloat64(6.424191233707441),
		Fx3:   -libc.Float64FromFloat64(1.5592452099634468e+308),
		Fy:    libc.Float64FromFloat64(1.5492083773956314e+308),
		Fdy:   float32(0.4997844994068146),
		Fe:    int32(FE_INEXACT),
	},
	787: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(787),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.747665004798627e+307),
		Fx2:   libc.Float64FromFloat64(5.208765824948828),
		Fx3:   -libc.Float64FromFloat64(1.685884069462192e+308),
		Fy:    libc.Float64FromFloat64(7.870634530678682e+307),
		Fdy:   float32(0.4998374581336975),
		Fe:    int32(FE_INEXACT),
	},
	788: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(788),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.801614987372249e+307),
		Fx2:   libc.Float64FromFloat64(6.560281135384934),
		Fx3:   -libc.Float64FromFloat64(1.5788701442037314e+308),
		Fy:    libc.Float64FromFloat64(1.5711242779002418e+308),
		Fdy:   float32(0.49981388449668884),
		Fe:    int32(FE_INEXACT),
	},
	789: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(789),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.998909030732365e+307),
		Fx2:   libc.Float64FromFloat64(5.4718021401385455),
		Fx3:   -libc.Float64FromFloat64(1.4337839828307078e+308),
		Fy:    libc.Float64FromFloat64(1.301520130441218e+308),
		Fdy:   float32(0.4995970129966736),
		Fe:    int32(FE_INEXACT),
	},
	790: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(790),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.356524507205283e+307),
		Fx2:   libc.Float64FromFloat64(5.221732002025672),
		Fx3:   -libc.Float64FromFloat64(1.3149412007165367e+308),
		Fy:    libc.Float64FromFloat64(1.4820923431743257e+308),
		Fdy:   float32(0.5000056028366089),
		Fe:    int32(FE_INEXACT),
	},
	791: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(791),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.810357702135086e+307),
		Fx2:   libc.Float64FromFloat64(5.397646708150679),
		Fx3:   -libc.Float64FromFloat64(1.204117134943272e+308),
		Fy:    libc.Float64FromFloat64(1.3923440066523994e+308),
		Fdy:   float32(0.5001450777053833),
		Fe:    int32(FE_INEXACT),
	},
	792: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(792),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.588316887017697e+307),
		Fx2:   libc.Float64FromFloat64(5.214681379814572),
		Fx3:   -libc.Float64FromFloat64(1.181755739589817e+308),
		Fy:    libc.Float64FromFloat64(1.732373461933635e+308),
		Fdy:   float32(0.5002310872077942),
		Fe:    int32(FE_INEXACT),
	},
	793: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(793),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.921493495074072e+307),
		Fx2:   libc.Float64FromFloat64(5.067366138955358),
		Fx3:   -libc.Float64FromFloat64(1.7736604119490367e+308),
		Fy:    libc.Float64FromFloat64(7.202405370537043e+307),
		Fdy:   float32(0.5005795359611511),
		Fe:    int32(FE_INEXACT),
	},
	794: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(794),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.613287579356673e+307),
		Fx2:   libc.Float64FromFloat64(4.705711329368165),
		Fx3:   -libc.Float64FromFloat64(9.004379644780656e+307),
		Fy:    libc.Float64FromFloat64(1.7410131312399642e+308),
		Fdy:   float32(0.5002294778823853),
		Fe:    int32(FE_INEXACT),
	},
	795: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(795),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.165326508807875e+307),
		Fx2:   libc.Float64FromFloat64(5.199316186584455),
		Fx3:   -libc.Float64FromFloat64(1.303655453861495e+308),
		Fy:    libc.Float64FromFloat64(1.3819611187623608e+308),
		Fdy:   float32(0.4996021091938019),
		Fe:    int32(FE_INEXACT),
	},
	796: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(796),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.108571665281885e+307),
		Fx2:   libc.Float64FromFloat64(5.333924115931674),
		Fx3:   -libc.Float64FromFloat64(1.5976450785694555e+308),
		Fy:    libc.Float64FromFloat64(1.6606206933649399e+308),
		Fdy:   float32(0.49984821677207947),
		Fe:    int32(FE_INEXACT),
	},
	797: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(797),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.026931649688656e+307),
		Fx2:   libc.Float64FromFloat64(4.823977588705634),
		Fx3:   -libc.Float64FromFloat64(1.156853129524485e+308),
		Fy:    libc.Float64FromFloat64(1.2681274322808268e+308),
		Fdy:   float32(0.5001069903373718),
		Fe:    int32(FE_INEXACT),
	},
	798: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(798),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.013778645030372e+307),
		Fx2:   libc.Float64FromFloat64(4.458342999119849),
		Fx3:   -libc.Float64FromFloat64(1.1987847093871429e+308),
		Fy:    libc.Float64FromFloat64(1.0365297827336333e+308),
		Fdy:   float32(0.49989405274391174),
		Fe:    int32(FE_INEXACT),
	},
	799: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(799),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.05548747672646e+307),
		Fx2:   libc.Float64FromFloat64(4.901075194729867),
		Fx3:   -libc.Float64FromFloat64(1.0880998722839145e+308),
		Fy:    libc.Float64FromFloat64(1.38963255466124e+308),
		Fdy:   float32(0.49962931871414185),
		Fe:    int32(FE_INEXACT),
	},
	800: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(800),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.569509420754374e+307),
		Fx2:   libc.Float64FromFloat64(5.065480295683743),
		Fx3:   -libc.Float64FromFloat64(1.4837555026035718e+308),
		Fy:    libc.Float64FromFloat64(8.309204905736797e+307),
		Fdy:   float32(0.5001950860023499),
		Fe:    int32(FE_INEXACT),
	},
	801: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(801),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.557579228561147e+307),
		Fx2:   libc.Float64FromFloat64(4.428123891994576),
		Fx3:   -libc.Float64FromFloat64(1.74835254329516e+308),
		Fy:    libc.Float64FromFloat64(7.126123930692802e+307),
		Fdy:   float32(0.500927746295929),
		Fe:    int32(FE_INEXACT),
	},
	802: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(802),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.660579037345791e+307),
		Fx2:   libc.Float64FromFloat64(4.296516507172668),
		Fx3:   -libc.Float64FromFloat64(1.489576081148976e+308),
		Fy:    libc.Float64FromFloat64(5.128493955449336e+307),
		Fdy:   float32(0.4990600645542145),
		Fe:    int32(FE_INEXACT),
	},
	803: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(803),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.83791302701495e+307),
		Fx2:   libc.Float64FromFloat64(5.343009061408011),
		Fx3:   -libc.Float64FromFloat64(1.696970852225873e+308),
		Fy:    libc.Float64FromFloat64(1.4222313680794016e+308),
		Fdy:   float32(0.5002058148384094),
		Fe:    int32(FE_INEXACT),
	},
	804: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(804),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.441003484470601e+307),
		Fx2:   libc.Float64FromFloat64(4.5373409725271845),
		Fx3:   -libc.Float64FromFloat64(1.1396064444589259e+308),
		Fy:    libc.Float64FromFloat64(1.3291623597162377e+308),
		Fdy:   float32(0.49984365701675415),
		Fe:    int32(FE_INEXACT),
	},
	805: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(805),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.801235155695051e+307),
		Fx2:   libc.Float64FromFloat64(5.063081192193199),
		Fx3:   -libc.Float64FromFloat64(1.6560392750859962e+308),
		Fy:    libc.Float64FromFloat64(7.748650665236436e+307),
		Fdy:   float32(0.5005207061767578),
		Fe:    int32(FE_INEXACT),
	},
	806: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(806),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.029173727473902e+307),
		Fx2:   libc.Float64FromFloat64(4.213927483311745),
		Fx3:   -libc.Float64FromFloat64(1.0933189427086426e+308),
		Fy:    libc.Float64FromFloat64(1.0259383961465222e+308),
		Fdy:   float32(0.5004449486732483),
		Fe:    int32(FE_INEXACT),
	},
	807: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(807),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.713246459303928e+307),
		Fx2:   libc.Float64FromFloat64(4.167628920894302),
		Fx3:   -libc.Float64FromFloat64(1.6025945509751142e+308),
		Fy:    libc.Float64FromFloat64(3.6171167453465783e+307),
		Fdy:   float32(0.4981565773487091),
		Fe:    int32(FE_INEXACT),
	},
	808: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(808),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.993885217541583e+307),
		Fx2:   libc.Float64FromFloat64(4.81978436543766),
		Fx3:   -libc.Float64FromFloat64(1.5975361383511683e+308),
		Fy:    libc.Float64FromFloat64(1.2913872876223149e+308),
		Fdy:   float32(0.5002398490905762),
		Fe:    int32(FE_INEXACT),
	},
	809: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(809),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9535263638368559),
		Fx2:   libc.Float64FromFloat64(0.6453699054483613),
		Fx3:   libc.Float64FromFloat64(0.6159768642976889),
		Fy:    libc.Float64FromFloat64(0.0005996450257775658),
		Fdy:   float32(0.5004515051841736),
		Fe:    int32(FE_INEXACT),
	},
	810: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(810),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.5812254895899245),
		Fx2:   -libc.Float64FromFloat64(0.8656947825024184),
		Fx3:   libc.Float64FromFloat64(0.5032226460046499),
		Fy:    libc.Float64FromFloat64(5.877220923862947e-05),
		Fdy:   float32(0.5207586884498596),
		Fe:    int32(FE_INEXACT),
	},
	811: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(811),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.8867409402454873),
		Fx2:   libc.Float64FromFloat64(0.8159904749863807),
		Fx3:   libc.Float64FromFloat64(0.7235194532273546),
		Fy:    -libc.Float64FromFloat64(5.270779343044255e-05),
		Fdy:   float32(0.49455180764198303),
		Fe:    int32(FE_INEXACT),
	},
	812: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(812),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.8912024283674143),
		Fx2:   libc.Float64FromFloat64(0.732508220935042),
		Fx3:   libc.Float64FromFloat64(0.6513743765181268),
		Fy:    -libc.Float64FromFloat64(0.0014387287782770062),
		Fdy:   float32(0.4997765123844147),
		Fe:    int32(FE_INEXACT),
	},
	813: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(813),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.607226606106348),
		Fx2:   libc.Float64FromFloat64(0.5733792825297863),
		Fx3:   libc.Float64FromFloat64(0.5709429075459292),
		Fy:    libc.Float64FromFloat64(0.22277175180367426),
		Fdy:   float32(0.500000536441803),
		Fe:    int32(FE_INEXACT),
	},
	814: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(814),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.7669891973936926),
		Fx2:   -libc.Float64FromFloat64(0.18059667464261908),
		Fx3:   libc.Float64FromFloat64(0.14975452359555128),
		Fy:    libc.Float64FromFloat64(0.011238825059439022),
		Fdy:   float32(0.5000004172325134),
		Fe:    int32(FE_INEXACT),
	},
	815: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(815),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.1939053833535347),
		Fx2:   -libc.Float64FromFloat64(0.843016615519224),
		Fx3:   libc.Float64FromFloat64(0.11779840359964053),
		Fy:    -libc.Float64FromFloat64(0.04566705640601396),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	816: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(816),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.5692028022936819),
		Fx2:   -libc.Float64FromFloat64(0.9129964007875562),
		Fx3:   libc.Float64FromFloat64(0.11183249764296027),
		Fy:    -libc.Float64FromFloat64(0.4078476121693622),
		Fdy:   float32(0.49994081258773804),
		Fe:    int32(FE_INEXACT),
	},
	817: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(817),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.708688852152128),
		Fx2:   -libc.Float64FromFloat64(0.9042447353356446),
		Fx3:   libc.Float64FromFloat64(0.12170937600575682),
		Fy:    libc.Float64FromFloat64(0.7625375395553796),
		Fdy:   float32(0.5000236630439758),
		Fe:    int32(FE_INEXACT),
	},
	818: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(818),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.6887843192514702),
		Fx2:   -libc.Float64FromFloat64(0.8945914094590282),
		Fx3:   libc.Float64FromFloat64(0.10712095229067814),
		Fy:    libc.Float64FromFloat64(0.7233014872631282),
		Fdy:   float32(0.5000261664390564),
		Fe:    int32(FE_INEXACT),
	},
	819: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(819),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9093487002930813),
		Fx2:   -libc.Float64FromFloat64(0.6353220424807806),
		Fx3:   libc.Float64FromFloat64(0.12209193072689345),
		Fy:    -libc.Float64FromFloat64(0.4556373428705502),
		Fdy:   float32(0.4999924302101135),
		Fe:    int32(FE_INEXACT),
	},
	820: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(820),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.8083244140324664),
		Fx2:   libc.Float64FromFloat64(0.7030885853349238),
		Fx3:   libc.Float64FromFloat64(0.1158722205072921),
		Fy:    -libc.Float64FromFloat64(0.4524514482464759),
		Fdy:   float32(0.4999793767929077),
		Fe:    int32(FE_INEXACT),
	},
	821: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(821),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7501997308042707),
		Fx2:   libc.Float64FromFloat64(1.7794178644284324),
		Fx3:   libc.Float64FromFloat64(0.034941880125231264),
		Fy:    libc.Float64FromFloat64(3.149278547436184),
		Fdy:   float32(0.5000026226043701),
		Fe:    int32(FE_INEXACT),
	},
	822: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(822),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.100626123876368),
		Fx2:   -libc.Float64FromFloat64(1.7941957343733712),
		Fx3:   libc.Float64FromFloat64(0.028998076215368075),
		Fy:    -libc.Float64FromFloat64(1.9457406203835088),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	823: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(823),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.2271453199908313),
		Fx2:   -libc.Float64FromFloat64(1.9066288212481188),
		Fx3:   libc.Float64FromFloat64(0.03643890739091297),
		Fy:    libc.Float64FromFloat64(2.3761495423451775),
		Fdy:   float32(0.5000033378601074),
		Fe:    int32(FE_INEXACT),
	},
	824: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(824),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.690189479774423),
		Fx2:   -libc.Float64FromFloat64(1.614737917080582),
		Fx3:   libc.Float64FromFloat64(0.015752263118042873),
		Fy:    -libc.Float64FromFloat64(2.7134607769244212),
		Fdy:   float32(0.49999865889549255),
		Fe:    int32(FE_INEXACT),
	},
	825: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(825),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.216633430136878),
		Fx2:   -libc.Float64FromFloat64(1.6298414477764052),
		Fx3:   libc.Float64FromFloat64(0.01518442879709737),
		Fy:    libc.Float64FromFloat64(1.9981040199845606),
		Fdy:   float32(0.5000008940696716),
		Fe:    int32(FE_INEXACT),
	},
	826: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(826),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.7079042204871593),
		Fx2:   -libc.Float64FromFloat64(1.8249386066786561),
		Fx3:   libc.Float64FromFloat64(0.01542138096007668),
		Fy:    -libc.Float64FromFloat64(1.2764603608376999),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	827: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(827),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.7906428884980627),
		Fx2:   libc.Float64FromFloat64(1.4457010439200353),
		Fx3:   libc.Float64FromFloat64(0.015292438404015948),
		Fy:    libc.Float64FromFloat64(1.1583256876736174),
		Fdy:   float32(0.500001847743988),
		Fe:    int32(FE_INEXACT),
	},
	828: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(828),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.8460153055895931),
		Fx2:   libc.Float64FromFloat64(1.4629262933341582),
		Fx3:   libc.Float64FromFloat64(0.006794043496898868),
		Fy:    libc.Float64FromFloat64(1.2444520786070477),
		Fdy:   float32(0.5000006556510925),
		Fe:    int32(FE_INEXACT),
	},
	829: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(829),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7682870382979763),
		Fx2:   -libc.Float64FromFloat64(1.0693842091290007),
		Fx3:   libc.Float64FromFloat64(0.006249663857703811),
		Fy:    -libc.Float64FromFloat64(1.8847285721056404),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	830: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(830),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.5233173093240433),
		Fx2:   -libc.Float64FromFloat64(0.5456889357081449),
		Fx3:   libc.Float64FromFloat64(0.0011895152895775862),
		Fy:    libc.Float64FromFloat64(0.8324469165604098),
		Fdy:   float32(0.5000001788139343),
		Fe:    int32(FE_INEXACT),
	},
	831: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(831),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.191881851083905),
		Fx2:   libc.Float64FromFloat64(0.835413528942726),
		Fx3:   libc.Float64FromFloat64(0.001180710008124752),
		Fy:    -libc.Float64FromFloat64(0.994533513288669),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	832: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(832),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.8351151527140477),
		Fx2:   libc.Float64FromFloat64(1.6137398506141194),
		Fx3:   libc.Float64FromFloat64(0.0005080982382209128),
		Fy:    libc.Float64FromFloat64(1.3481667000245758),
		Fdy:   float32(0.5000000596046448),
		Fe:    int32(FE_INEXACT),
	},
	833: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(833),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    libc.Float64FromFloat64(8.69169475992e-311),
		Fdy:   float32(-libc.Float64FromFloat64(3.725290298461914e-09)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	834: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(834),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    -libc.Float64FromFloat64(8.691694760173e-311),
		Fdy:   float32(3.725290298461914e-09),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	835: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(835),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    libc.Float64FromFloat64(8.69169475992e-311),
		Fdy:   float32(-libc.Float64FromFloat64(3.725290298461914e-09)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	836: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(836),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    -libc.Float64FromFloat64(8.6916947601737e-311),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	837: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(837),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    libc.Float64FromFloat64(8.6916947599207e-311),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	838: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(838),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    -libc.Float64FromFloat64(8.691694760173e-311),
		Fdy:   float32(3.725290298461914e-09),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	839: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(839),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    libc.Float64FromFloat64(8.69169475992e-311),
		Fdy:   float32(-libc.Float64FromFloat64(3.725290298461914e-09)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	840: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(840),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.332636185167997e-302),
		Fx2:   libc.Float64FromFloat64(9.31322574629031e-10),
		Fx3:   -libc.Float64FromFloat64(1.265e-321),
		Fy:    -libc.Float64FromFloat64(8.691694760173e-311),
		Fdy:   float32(3.725290298461914e-09),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	841: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(841),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(5.293956235883739e-23),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(5.960464477539063e-08),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	842: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(842),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(5.293956235883739e-23),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(5.960464477539063e-08),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	843: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(843),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(5.293956235883739e-23),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	844: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(844),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(5.293956235883739e-23),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	845: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(848),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(6.223015277861142e-61),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	846: {
		Ffile: __ccgo_ts + 22,
		Fline: int32(849),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.332636185032189e-302),
		Fx2:   libc.Float64FromFloat64(6.223015277861142e-61),
		Fx3:   libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:15:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:16:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:20:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:23:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(44044)/libc.Uint32FromInt64(52)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:24:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*52
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:26:3:
		if (*lll_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:27:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:28:3:
		libc.Xfesetround(tls, (*lll_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:29:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:30:3:
		y = libc.Xfmal(tls, (*lll_l)(unsafe.Pointer(p)).Fx, (*lll_l)(unsafe.Pointer(p)).Fx2, (*lll_l)(unsafe.Pointer(p)).Fx3)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:31:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		/* do not check inexact by default */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:37:3:
		if !(checkexceptall(tls, e|int32(FE_INEXACT), (*lll_l)(unsafe.Pointer(p)).Fe|int32(FE_INEXACT), (*lll_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+46, libc.VaList(bp+8, (*lll_l)(unsafe.Pointer(p)).Ffile, (*lll_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*lll_l)(unsafe.Pointer(p)).Fr), (*lll_l)(unsafe.Pointer(p)).Fx, (*lll_l)(unsafe.Pointer(p)).Fx2, (*lll_l)(unsafe.Pointer(p)).Fx3, (*lll_l)(unsafe.Pointer(p)).Fy, estr(tls, (*lll_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:41:4:
			libc.Xprintf(tls, __ccgo_ts+105, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:42:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:44:3:
		d = ulperrl(tls, y, (*lll_l)(unsafe.Pointer(p)).Fy, (*lll_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:45:3:
		if !(checkcr(tls, y, (*lll_l)(unsafe.Pointer(p)).Fy, (*lll_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:46:4:
			libc.Xprintf(tls, __ccgo_ts+114, libc.VaList(bp+8, (*lll_l)(unsafe.Pointer(p)).Ffile, (*lll_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*lll_l)(unsafe.Pointer(p)).Fr), (*lll_l)(unsafe.Pointer(p)).Fx, (*lll_l)(unsafe.Pointer(p)).Fx2, (*lll_l)(unsafe.Pointer(p)).Fx3, (*lll_l)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*lll_l)(unsafe.Pointer(p)).Fdy), float64((*lll_l)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:48:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmal.c:51:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+182, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+192, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 236,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 244,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 252,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 262,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 272,
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
				v2 = __ccgo_ts + 281
			} else {
				v2 = __ccgo_ts + 45
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+283, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 281
		} else {
			v3 = __ccgo_ts + 45
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+288, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 45
	} else {
		v4 = __ccgo_ts + 293
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+295, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 298
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 301
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 304
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 307
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 310
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+313, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+321, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+326, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+369, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+418) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+426) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+438) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+450) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+462) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+471) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+45) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+471) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+477, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/sanity/fma.h\x00src/math/special/fma.h\x00\x00%s:%d: bad fp exception: %s fmal(%La,%La,%La)=%La, want %s\x00 got %s\n\x00%s:%d: %s fmal(%La,%La,%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
