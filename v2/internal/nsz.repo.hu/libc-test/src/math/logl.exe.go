// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/logl.exe.go src/math/logl.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:5:19:
var t = [1509]l_l{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(15),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(16),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(17),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(18),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213813),
		Fdy:   float32(-libc.Float64FromFloat64(0.6109976768493652)),
		Fe:    int32(FE_INEXACT),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(27),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(29),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(30),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(31),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(32),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(33),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(34),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7421227853654382),
		Fy:    libc.Float64FromFloat64(0.5551043612307478),
		Fdy:   float32(-libc.Float64FromFloat64(9.596701180307563e-16)),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.726348535513887),
		Fy:    libc.Float64FromFloat64(0.5460085047865537),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.030451738637547),
		Fy:    libc.Float64FromFloat64(1.39387846363126),
		Fdy:   float32(-libc.Float64FromFloat64(3.07812487716262e-16)),
		Fe:    int32(FE_INEXACT),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.09249642977486),
		Fy:    libc.Float64FromFloat64(1.4091551578803312),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.3525143247996634),
		Fy:    libc.Float64FromFloat64(1.2097106090354592),
		Fdy:   float32(-libc.Float64FromFloat64(2.754493440351453e-16)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.716861910767595),
		Fy:    libc.Float64FromFloat64(1.743420037847665),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.477122252887574),
		Fy:    libc.Float64FromFloat64(1.868276314947396),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.201173315249939),
		Fy:    libc.Float64FromFloat64(1.974243973196778),
		Fdy:   float32(-libc.Float64FromFloat64(3.584595787176916e-16)),
		Fe:    int32(FE_INEXACT),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.11023548637455),
		Fy:    libc.Float64FromFloat64(2.0931269042501035),
		Fdy:   float32(-libc.Float64FromFloat64(2.709724572715511e-16)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(17.095327585176175),
		Fy:    libc.Float64FromFloat64(2.838805185538619),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(22.661177212233824),
		Fy:    libc.Float64FromFloat64(3.1206532053640323),
		Fdy:   float32(-libc.Float64FromFloat64(2.684326554489887e-16)),
		Fe:    int32(FE_INEXACT),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(49.39281360008446),
		Fy:    libc.Float64FromFloat64(3.8998049399290537),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(71.03510946091578),
		Fy:    libc.Float64FromFloat64(4.263174254266271),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(222.92105098013496),
		Fy:    libc.Float64FromFloat64(5.406817677296964),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(428.31524716519823),
		Fy:    libc.Float64FromFloat64(6.059859483252683),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(138236.95959852284),
		Fy:    libc.Float64FromFloat64(11.836724590151075),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.754458852196048e+06),
		Fy:    libc.Float64FromFloat64(15.374593441427612),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9979003645601828e+07),
		Fy:    libc.Float64FromFloat64(16.810192462353758),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1227898974582405e+09),
		Fy:    libc.Float64FromFloat64(20.839082404779035),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0330780385933456e+10),
		Fy:    libc.Float64FromFloat64(23.058393662813447),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.5672925182012714e+07),
		Fy:    libc.Float64FromFloat64(17.83500450381271),
		Fdy:   float32(-libc.Float64FromFloat64(9.762054717105811e-19)),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.874672998883447e+07),
		Fy:    libc.Float64FromFloat64(17.702148676006242),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.385338848000864e+11),
		Fy:    libc.Float64FromFloat64(26.197777211816646),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4451911223729425e+14),
		Fy:    libc.Float64FromFloat64(33.13031459044816),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.1271191358301664e+17),
		Fy:    libc.Float64FromFloat64(40.77849051037793),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9643960027070173e+19),
		Fy:    libc.Float64FromFloat64(44.42430158720062),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.057563107060638e+27),
		Fy:    libc.Float64FromFloat64(62.89131983565637),
		Fdy:   float32(-libc.Float64FromFloat64(6.6741166985514544e-18)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8866265548661536e+24),
		Fy:    libc.Float64FromFloat64(56.61958380369774),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.771633820386861e+41),
		Fy:    libc.Float64FromFloat64(94.97789099594698),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.020425758407352e+44),
		Fy:    libc.Float64FromFloat64(102.70513189881244),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2580908456763673e+45),
		Fy:    libc.Float64FromFloat64(103.84592455477211),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.152017462376579e+62),
		Fy:    libc.Float64FromFloat64(143.52668152238468),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3777209648246645e+74),
		Fy:    libc.Float64FromFloat64(170.71172754077602),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.849257202283994e+102),
		Fy:    libc.Float64FromFloat64(236.21155968060643),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.61872491219804e+103),
		Fy:    libc.Float64FromFloat64(238.69638325235653),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5719651818861894e+125),
		Fy:    libc.Float64FromFloat64(288.2754631690901),
		Fdy:   float32(-libc.Float64FromFloat64(6.162429389071052e-19)),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.272373929833471e+137),
		Fy:    libc.Float64FromFloat64(317.1166384623794),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.371102901169479e+156),
		Fy:    libc.Float64FromFloat64(361.3280601512514),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.360205199608655e+139),
		Fy:    libc.Float64FromFloat64(321.53184704655627),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.20500185550749e+160),
		Fy:    libc.Float64FromFloat64(370.5183588297328),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0854514513414446e+163),
		Fy:    libc.Float64FromFloat64(376.7288023946394),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7012604239472096e+204),
		Fy:    libc.Float64FromFloat64(470.25872837297845),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.1498325550794055e+213),
		Fy:    libc.Float64FromFloat64(491.59797410203873),
		Fdy:   float32(-libc.Float64FromFloat64(3.7335624128193457e-19)),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7331122411293055e+233),
		Fy:    libc.Float64FromFloat64(537.0522454432017),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.645454875450314e+289),
		Fy:    libc.Float64FromFloat64(666.9829811707267),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5632533809874178),
		Fy:    -libc.Float64FromFloat64(0.5740256970470217),
		Fdy:   float32(6.592782477373728e-16),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.574197928566325),
		Fy:    -libc.Float64FromFloat64(0.5547811188141762),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5837592676331093),
		Fy:    -libc.Float64FromFloat64(0.5382665940940305),
		Fdy:   float32(8.069118552790611e-16),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.6378065772902207),
		Fy:    -libc.Float64FromFloat64(0.4497202119942843),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.6572138124139603),
		Fy:    -libc.Float64FromFloat64(0.41974587597664104),
		Fdy:   float32(8.274927441940909e-16),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.964721556182863),
		Fy:    -libc.Float64FromFloat64(0.03591576209532481),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0828996954117838),
		Fy:    libc.Float64FromFloat64(0.07964234638129533),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2783287003041894),
		Fy:    libc.Float64FromFloat64(0.2455535218714178),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3170653074678824),
		Fy:    libc.Float64FromFloat64(0.2754060095862774),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4711698109944988),
		Fy:    libc.Float64FromFloat64(0.3860578741100104),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.726348535513887),
		Fy:    libc.Float64FromFloat64(0.5460085047865537),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(4.440892098500625e-16),
		Fdy:   float32(-libc.Float64FromFloat64(5.921189111737107e-16)),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.240257022715406e-46),
		Fy:    -libc.Float64FromFloat64(104.08789290665531),
		Fdy:   float32(-libc.Float64FromFloat64(0.49965426325798035)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.7396630234059945e+237),
		Fy:    libc.Float64FromFloat64(547.2686330806602),
		Fdy:   float32(-libc.Float64FromFloat64(0.49917951226234436)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4451440459536037e-116),
		Fy:    -libc.Float64FromFloat64(266.20576675161743),
		Fdy:   float32(-libc.Float64FromFloat64(0.4991205930709839)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.438831660406453e+63),
		Fy:    libc.Float64FromFloat64(146.75642512811908),
		Fdy:   float32(-libc.Float64FromFloat64(0.000828784191980958)),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.438831660406453e+63),
		Fy:    libc.Float64FromFloat64(146.75642512811908),
		Fdy:   float32(-libc.Float64FromFloat64(0.000828784191980958)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.438831660406453e+63),
		Fy:    libc.Float64FromFloat64(146.7564251281191),
		Fdy:   float32(0.9991711974143982),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.613307411535328e-54),
		Fy:    -libc.Float64FromFloat64(122.81064997905968),
		Fdy:   float32(0.9993534684181213),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.613307411535328e-54),
		Fy:    -libc.Float64FromFloat64(122.81064997905968),
		Fdy:   float32(0.9993534684181213),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.613307411535328e-54),
		Fy:    -libc.Float64FromFloat64(122.81064997905969),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006465032929554582)),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0651951942848229e+39),
		Fy:    libc.Float64FromFloat64(89.86397669015386),
		Fdy:   float32(-libc.Float64FromFloat64(0.4997045397758484)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.376992814880975e+217),
		Fy:    libc.Float64FromFloat64(500.87795079344573),
		Fdy:   float32(-libc.Float64FromFloat64(0.49964672327041626)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7605918393031102e+70),
		Fy:    libc.Float64FromFloat64(162.19640160085456),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.849382817174757e-174),
		Fy:    -libc.Float64FromFloat64(398.5893712742764),
		Fdy:   float32(0.9999268054962158),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(7.849382817174757e-174),
		Fy:    -libc.Float64FromFloat64(398.5893712742764),
		Fdy:   float32(0.9999268054962158),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(7.849382817174757e-174),
		Fy:    -libc.Float64FromFloat64(398.58937127427646),
		Fdy:   float32(-libc.Float64FromFloat64(7.316956180147827e-05)),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.646086612907075e+125),
		Fy:    libc.Float64FromFloat64(289.55409929454794),
		Fdy:   float32(-libc.Float64FromFloat64(0.9998994469642639)),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.646086612907075e+125),
		Fy:    libc.Float64FromFloat64(289.55409929454794),
		Fdy:   float32(-libc.Float64FromFloat64(0.9998994469642639)),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.646086612907075e+125),
		Fy:    libc.Float64FromFloat64(289.554099294548),
		Fdy:   float32(0.00010054100857814774),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6736883116634555e+18),
		Fy:    libc.Float64FromFloat64(41.96156143487446),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999533295631409)),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.6736883116634555e+18),
		Fy:    libc.Float64FromFloat64(41.96156143487446),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999533295631409)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.6736883116634555e+18),
		Fy:    libc.Float64FromFloat64(41.96156143487447),
		Fdy:   float32(4.664263178710826e-05),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.971896830449522e-160),
		Fy:    -libc.Float64FromFloat64(367.03437110740873),
		Fdy:   float32(0.999961793422699),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.971896830449522e-160),
		Fy:    -libc.Float64FromFloat64(367.03437110740873),
		Fdy:   float32(0.999961793422699),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.971896830449522e-160),
		Fy:    -libc.Float64FromFloat64(367.0343711074088),
		Fdy:   float32(-libc.Float64FromFloat64(3.821008431259543e-05)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.327353370923574e-56),
		Fy:    -libc.Float64FromFloat64(127.47980908326339),
		Fdy:   float32(0.9999094605445862),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.327353370923574e-56),
		Fy:    -libc.Float64FromFloat64(127.47980908326339),
		Fdy:   float32(0.9999094605445862),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.327353370923574e-56),
		Fy:    -libc.Float64FromFloat64(127.4798090832634),
		Fdy:   float32(-libc.Float64FromFloat64(9.051470260601491e-05)),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.707831708937764e+87),
		Fy:    libc.Float64FromFloat64(202.36714000749757),
		Fdy:   float32(-libc.Float64FromFloat64(0.49994680285453796)),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6383808675082076e+96),
		Fy:    libc.Float64FromFloat64(222.01833434867441),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999009072780609)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2187014418702487e-32),
		Fy:    -libc.Float64FromFloat64(72.88580088716162),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999526143074036)),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.879079408094281e-185),
		Fy:    -libc.Float64FromFloat64(424.62264434437037),
		Fdy:   float32(1.9103258637187537e-06),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.879079408094281e-185),
		Fy:    -libc.Float64FromFloat64(424.62264434437037),
		Fdy:   float32(1.9103258637187537e-06),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.879079408094281e-185),
		Fy:    -libc.Float64FromFloat64(424.6226443443704),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999980926513672)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.014729326042983e+300),
		Fy:    libc.Float64FromFloat64(692.3879073452982),
		Fdy:   float32(-libc.Float64FromFloat64(1.3890728951082565e-05)),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.014729326042983e+300),
		Fy:    libc.Float64FromFloat64(692.3879073452982),
		Fdy:   float32(-libc.Float64FromFloat64(1.3890728951082565e-05)),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.014729326042983e+300),
		Fy:    libc.Float64FromFloat64(692.3879073452983),
		Fdy:   float32(0.9999861121177673),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.066860594449529e-156),
		Fy:    -libc.Float64FromFloat64(357.40043323604937),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999924600124359)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.809288065449752e-297),
		Fy:    -libc.Float64FromFloat64(682.2972235574097),
		Fdy:   float32(6.198263236001367e-06),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.809288065449752e-297),
		Fy:    -libc.Float64FromFloat64(682.2972235574097),
		Fdy:   float32(6.198263236001367e-06),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.809288065449752e-297),
		Fy:    -libc.Float64FromFloat64(682.2972235574098),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999938011169434)),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.466980052484247e+161),
		Fy:    libc.Float64FromFloat64(372.21291254887285),
		Fdy:   float32(-libc.Float64FromFloat64(0.49998629093170166)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2035764523939172e+208),
		Fy:    libc.Float64FromFloat64(480.1019671698441),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.839053941045579e-43),
		Fy:    -libc.Float64FromFloat64(97.96768812141758),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999815225601196)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.176145894361368e-42),
		Fy:    -libc.Float64FromFloat64(96.54633100413716),
		Fdy:   float32(0.9999988675117493),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.176145894361368e-42),
		Fy:    -libc.Float64FromFloat64(96.54633100413716),
		Fdy:   float32(0.9999988675117493),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.176145894361368e-42),
		Fy:    -libc.Float64FromFloat64(96.54633100413717),
		Fdy:   float32(-libc.Float64FromFloat64(1.1115938605144038e-06)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.414032486625784e-15),
		Fy:    -libc.Float64FromFloat64(34.192330852717355),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999983310699463)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.379836090751614e+134),
		Fy:    libc.Float64FromFloat64(310.399544867124),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1610104271440014e+242),
		Fy:    libc.Float64FromFloat64(558.3764842364723),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999821186065674)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.852254521726481e-306),
		Fy:    -libc.Float64FromFloat64(702.3033381417298),
		Fdy:   float32(1.7284099840253475e-07),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(9.852254521726481e-306),
		Fy:    -libc.Float64FromFloat64(702.3033381417298),
		Fdy:   float32(1.7284099840253475e-07),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(9.852254521726481e-306),
		Fy:    -libc.Float64FromFloat64(702.3033381417299),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.635428434830941e-107),
		Fy:    -libc.Float64FromFloat64(245.08587798254538),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999982416629791)),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.724506177502807e-187),
		Fy:    -libc.Float64FromFloat64(428.8386561013714),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7855179997835594e+80),
		Fy:    libc.Float64FromFloat64(184.78651600861372),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.602850080515328e-26),
		Fy:    -libc.Float64FromFloat64(57.60515248004554),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.434306131930632e+297),
		Fy:    libc.Float64FromFloat64(686.0000800712567),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.434306131930632e+297),
		Fy:    libc.Float64FromFloat64(686.0000800712567),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.434306131930632e+297),
		Fy:    libc.Float64FromFloat64(686.0000800712568),
		Fdy:   float32(5.699967164929376e-10),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2975323128560564e+77),
		Fy:    libc.Float64FromFloat64(178.13088780004549),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2975323128560564e+77),
		Fy:    libc.Float64FromFloat64(178.13088780004549),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2975323128560564e+77),
		Fy:    libc.Float64FromFloat64(178.1308878000455),
		Fdy:   float32(8.713315224895268e-08),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.643992397019563e-147),
		Fy:    -libc.Float64FromFloat64(336.21367349816984),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(9.643992397019563e-147),
		Fy:    -libc.Float64FromFloat64(336.21367349816984),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(9.643992397019563e-147),
		Fy:    -libc.Float64FromFloat64(336.2136734981699),
		Fdy:   float32(-libc.Float64FromFloat64(1.4301173223429942e-07)),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2941524555964039e+84),
		Fy:    libc.Float64FromFloat64(193.6750038179465),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5175711230458892e+184),
		Fy:    libc.Float64FromFloat64(424.0927682223216),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.5175711230458892e+184),
		Fy:    libc.Float64FromFloat64(424.0927682223216),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.5175711230458892e+184),
		Fy:    libc.Float64FromFloat64(424.09276822232164),
		Fdy:   float32(2.6526727481268608e-08),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3735308192926206e+166),
		Fy:    libc.Float64FromFloat64(382.5465101018224),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.0028393234327676e+202),
		Fy:    libc.Float64FromFloat64(466.2217470670165),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.0028393234327676e+202),
		Fy:    libc.Float64FromFloat64(466.2217470670165),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.0028393234327676e+202),
		Fy:    libc.Float64FromFloat64(466.22174706701657),
		Fdy:   float32(1.8826657566251015e-08),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1506872841537574e-147),
		Fy:    -libc.Float64FromFloat64(337.71422121201954),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.312997553441186e-63),
		Fy:    -libc.Float64FromFloat64(143.22025024678246),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.3086423164618517e-21),
		Fy:    -libc.Float64FromFloat64(47.15774902392681),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9263883495896065e-128),
		Fy:    -libc.Float64FromFloat64(293.6571228857321),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.017989105292117e+108),
		Fy:    libc.Float64FromFloat64(249.3812915665087),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.306500692862194e+142),
		Fy:    libc.Float64FromFloat64(329.0841216294109),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.1810458932903514e+48),
		Fy:    libc.Float64FromFloat64(111.30388899302791),
		Fdy:   float32(-libc.Float64FromFloat64(1.5123329255573026e-09)),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.1810458932903514e+48),
		Fy:    libc.Float64FromFloat64(111.30388899302791),
		Fdy:   float32(-libc.Float64FromFloat64(1.5123329255573026e-09)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.1810458932903514e+48),
		Fy:    libc.Float64FromFloat64(111.30388899302793),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.502454551965461e-191),
		Fy:    -libc.Float64FromFloat64(438.2891300578011),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.502454551965461e-191),
		Fy:    -libc.Float64FromFloat64(438.2891300578011),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.502454551965461e-191),
		Fy:    -libc.Float64FromFloat64(438.2891300578012),
		Fdy:   float32(-libc.Float64FromFloat64(3.203395104733886e-09)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6705525679302026e+144),
		Fy:    libc.Float64FromFloat64(332.5545387964236),
		Fdy:   float32(-libc.Float64FromFloat64(2.839232182694218e-09)),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.6705525679302026e+144),
		Fy:    libc.Float64FromFloat64(332.5545387964236),
		Fdy:   float32(-libc.Float64FromFloat64(2.839232182694218e-09)),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.6705525679302026e+144),
		Fy:    libc.Float64FromFloat64(332.5545387964237),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.270885881037552e-165),
		Fy:    -libc.Float64FromFloat64(377.6996614047083),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3209861418882308e+29),
		Fy:    libc.Float64FromFloat64(67.0533462316904),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.935488738265274e+167),
		Fy:    libc.Float64FromFloat64(385.1920704025098),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0008322397534104),
		Fy:    libc.Float64FromFloat64(0.0008318936339297721),
		Fdy:   float32(-libc.Float64FromFloat64(0.12009137123823166)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9835850093081454),
		Fy:    -libc.Float64FromFloat64(0.01655120939456235),
		Fdy:   float32(0.4384218454360962),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9866338084047975),
		Fy:    -libc.Float64FromFloat64(0.013456323179109275),
		Fdy:   float32(-libc.Float64FromFloat64(0.13074573874473572)),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0290075380718406),
		Fy:    libc.Float64FromFloat64(0.028594782453691195),
		Fdy:   float32(-libc.Float64FromFloat64(0.17229929566383362)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0160286263618885),
		Fy:    libc.Float64FromFloat64(0.015901524312406583),
		Fdy:   float32(0.4639400839805603),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0193777275917844),
		Fy:    libc.Float64FromFloat64(0.01919237013870613),
		Fdy:   float32(-libc.Float64FromFloat64(0.30323463678359985)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9888805810724204),
		Fy:    -libc.Float64FromFloat64(0.011181701796136154),
		Fdy:   float32(-libc.Float64FromFloat64(0.0510830320417881)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9908142338328195),
		Fy:    -libc.Float64FromFloat64(0.009228215470102843),
		Fdy:   float32(0.3999338746070862),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0171254615924559),
		Fy:    libc.Float64FromFloat64(0.016980473855356407),
		Fdy:   float32(0.4382856488227844),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0219925534836878),
		Fy:    libc.Float64FromFloat64(0.021754205535484485),
		Fdy:   float32(-libc.Float64FromFloat64(0.337487131357193)),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9949543748283334),
		Fy:    -libc.Float64FromFloat64(0.005058397318777487),
		Fdy:   float32(0.05873437225818634),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.741341202236685e+172),
		Fy:    libc.Float64FromFloat64(397.7923288369396),
		Fdy:   float32(-libc.Float64FromFloat64(0.0009575664298608899)),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8428943217149917e-112),
		Fy:    -libc.Float64FromFloat64(256.8447077547533),
		Fdy:   float32(-libc.Float64FromFloat64(0.0001291344378842041)),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.996552912589139e-159),
		Fy:    -libc.Float64FromFloat64(364.3198449964904),
		Fdy:   float32(-libc.Float64FromFloat64(0.0013127467827871442)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.152683588140919e-15),
		Fy:    -libc.Float64FromFloat64(32.324729270302456),
		Fdy:   float32(0.011537257581949234),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.427354002934097e+13),
		Fy:    libc.Float64FromFloat64(32.06508905301064),
		Fdy:   float32(-libc.Float64FromFloat64(0.004408243112266064)),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9394980218120699e+121),
		Fy:    libc.Float64FromFloat64(279.2752254402295),
		Fdy:   float32(2.4246377506642602e-05),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.288050045747404e-286),
		Fy:    -libc.Float64FromFloat64(657.0835045014802),
		Fdy:   float32(-libc.Float64FromFloat64(0.0008952662465162575)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.009320336293199e+143),
		Fy:    libc.Float64FromFloat64(330.3713825495163),
		Fdy:   float32(-libc.Float64FromFloat64(0.00012530524691101164)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.4250470408041695e+208),
		Fy:    libc.Float64FromFloat64(480.6287259132664),
		Fdy:   float32(0.0006021694862283766),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1403438989409624e+288),
		Fy:    libc.Float64FromFloat64(663.2758366649373),
		Fdy:   float32(0.00022887045633979142),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3195079107728942),
		Fy:    libc.Float64FromFloat64(0.2772588722239781),
		Fdy:   float32(0.4548284709453583),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2311444133449163),
		Fy:    libc.Float64FromFloat64(0.20794415416798362),
		Fdy:   float32(0.18117491900920868),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.699999999999992),
		Fy:    libc.Float64FromFloat64(0.9932517730102804),
		Fdy:   float32(-libc.Float64FromFloat64(0.42995986342430115)),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.899999999999997),
		Fy:    libc.Float64FromFloat64(1.0647107369924271),
		Fdy:   float32(-libc.Float64FromFloat64(0.49179887771606445)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7873808054318492),
		Fy:    libc.Float64FromFloat64(0.5807513110867752),
		Fdy:   float32(-libc.Float64FromFloat64(0.49695637822151184)),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.101631084009031),
		Fy:    libc.Float64FromFloat64(0.09679188518111734),
		Fdy:   float32(0.49142447113990784),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5828174211072428),
		Fy:    libc.Float64FromFloat64(0.9488808267943696),
		Fdy:   float32(0.49506866931915283),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3243713696080635),
		Fy:    libc.Float64FromFloat64(0.8434496238172203),
		Fdy:   float32(-libc.Float64FromFloat64(0.4957887530326843)),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3720330376470498),
		Fy:    libc.Float64FromFloat64(0.31629360893145164),
		Fdy:   float32(-libc.Float64FromFloat64(0.49346116185188293)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9560474639218035),
		Fy:    libc.Float64FromFloat64(0.6709258371269086),
		Fdy:   float32(-libc.Float64FromFloat64(0.49915921688079834)),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9861652490003219),
		Fy:    libc.Float64FromFloat64(0.6862057691102773),
		Fdy:   float32(0.49405014514923096),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9841198793913846),
		Fy:    libc.Float64FromFloat64(0.6851754301176153),
		Fdy:   float32(0.4978952407836914),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9827574699742507),
		Fy:    libc.Float64FromFloat64(0.6844885374558497),
		Fdy:   float32(0.4961881637573242),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9800354570206753),
		Fy:    libc.Float64FromFloat64(0.6831147521323047),
		Fdy:   float32(0.4962450861930847),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9658060783284743),
		Fy:    libc.Float64FromFloat64(0.6759023791837),
		Fdy:   float32(-libc.Float64FromFloat64(0.4980953633785248)),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9617593559473196),
		Fy:    libc.Float64FromFloat64(0.6738417011983828),
		Fdy:   float32(0.49582648277282715),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9597391197314769),
		Fy:    libc.Float64FromFloat64(0.6728113622057269),
		Fdy:   float32(-libc.Float64FromFloat64(0.49907368421554565)),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9543619884569003),
		Fy:    libc.Float64FromFloat64(0.6700637915586349),
		Fdy:   float32(0.49321919679641724),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.948330349121657),
		Fy:    libc.Float64FromFloat64(0.6669727745806637),
		Fdy:   float32(0.49128198623657227),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9436519467371645),
		Fy:    libc.Float64FromFloat64(0.664568650264462),
		Fdy:   float32(-libc.Float64FromFloat64(0.4963008761405945)),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9429845212265147),
		Fy:    libc.Float64FromFloat64(0.6642252039335733),
		Fdy:   float32(0.49253129959106445),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9409836194926928),
		Fy:    libc.Float64FromFloat64(0.6631948649409037),
		Fdy:   float32(0.4926552474498749),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9403171102519186),
		Fy:    libc.Float64FromFloat64(0.662851418610029),
		Fdy:   float32(-libc.Float64FromFloat64(0.49822714924812317)),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.934328816847641),
		Fy:    libc.Float64FromFloat64(0.6597604016320563),
		Fdy:   float32(0.49932998418807983),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9330005968024173),
		Fy:    libc.Float64FromFloat64(0.6590735089702834),
		Fdy:   float32(-libc.Float64FromFloat64(0.49977031350135803)),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9310099765976125),
		Fy:    libc.Float64FromFloat64(0.6580431699776256),
		Fdy:   float32(-libc.Float64FromFloat64(0.4977264404296875)),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9191092306621644),
		Fy:    libc.Float64FromFloat64(0.6518611360216776),
		Fdy:   float32(0.4918416142463684),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9131863899639374),
		Fy:    libc.Float64FromFloat64(0.6487701190437036),
		Fdy:   float32(-libc.Float64FromFloat64(0.4946287274360657)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9099038250031526),
		Fy:    libc.Float64FromFloat64(0.6470528873892726),
		Fdy:   float32(-libc.Float64FromFloat64(0.49593690037727356)),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9085923765447872),
		Fy:    libc.Float64FromFloat64(0.6463659947274977),
		Fdy:   float32(0.4947270452976227),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9072818286014444),
		Fy:    libc.Float64FromFloat64(0.6456791020657301),
		Fdy:   float32(0.4936542809009552),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9020486296175882),
		Fy:    libc.Float64FromFloat64(0.6429315314186387),
		Fdy:   float32(0.49363088607788086),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8955273183997507),
		Fy:    libc.Float64FromFloat64(0.6394970681097802),
		Fdy:   float32(-libc.Float64FromFloat64(0.4934515357017517)),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.894876418278117),
		Fy:    libc.Float64FromFloat64(0.6391536217788613),
		Fdy:   float32(-libc.Float64FromFloat64(0.490817129611969)),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8857872503323154),
		Fy:    libc.Float64FromFloat64(0.6343453731464892),
		Fdy:   float32(-libc.Float64FromFloat64(0.490698903799057)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8786763571143712),
		Fy:    libc.Float64FromFloat64(0.6305674635067444),
		Fdy:   float32(0.4915044903755188),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8683810779995265),
		Fy:    libc.Float64FromFloat64(0.6250723222125631),
		Fdy:   float32(-libc.Float64FromFloat64(0.49692854285240173)),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8606966604753445),
		Fy:    libc.Float64FromFloat64(0.6209509662419352),
		Fdy:   float32(0.4964044690132141),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.850499911367864),
		Fy:    libc.Float64FromFloat64(0.6154558249477547),
		Fdy:   float32(0.49955394864082336),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8479594671604433),
		Fy:    libc.Float64FromFloat64(0.6140820396242144),
		Fdy:   float32(-libc.Float64FromFloat64(0.4987090528011322)),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8473249012377164),
		Fy:    libc.Float64FromFloat64(0.6137385932933224),
		Fdy:   float32(-libc.Float64FromFloat64(0.49816015362739563)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8422562120429549),
		Fy:    libc.Float64FromFloat64(0.6109910226462344),
		Fdy:   float32(-libc.Float64FromFloat64(0.49055591225624084)),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8397270851331367),
		Fy:    libc.Float64FromFloat64(0.6096172373226979),
		Fdy:   float32(0.49639299511909485),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8384638240107538),
		Fy:    libc.Float64FromFloat64(0.6089303446609211),
		Fdy:   float32(-libc.Float64FromFloat64(0.4993882477283478)),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.837201430315228),
		Fy:    libc.Float64FromFloat64(0.6082434519991503),
		Fdy:   float32(-libc.Float64FromFloat64(0.49634653329849243)),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8365705585663812),
		Fy:    libc.Float64FromFloat64(0.6079000056682635),
		Fdy:   float32(-libc.Float64FromFloat64(0.4986101984977722)),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8340492371610513),
		Fy:    libc.Float64FromFloat64(0.6065262203447045),
		Fdy:   float32(-libc.Float64FromFloat64(0.49065500497817993)),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.833419447835258),
		Fy:    libc.Float64FromFloat64(0.6061827740136913),
		Fdy:   float32(-libc.Float64FromFloat64(0.49018043279647827)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.829645250779472),
		Fy:    libc.Float64FromFloat64(0.6041220960285203),
		Fdy:   float32(0.49995192885398865),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8290169737271673),
		Fy:    libc.Float64FromFloat64(0.6037786496976335),
		Fdy:   float32(0.49813365936279297),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8239985178152924),
		Fy:    libc.Float64FromFloat64(0.6010310790505444),
		Fdy:   float32(0.4917064607143402),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8233721797797005),
		Fy:    libc.Float64FromFloat64(0.600687632719651),
		Fdy:   float32(0.49271926283836365),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8221201488643952),
		Fy:    libc.Float64FromFloat64(0.600000740057886),
		Fdy:   float32(0.49449554085731506),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8208689776646199),
		Fy:    libc.Float64FromFloat64(0.5993138473961108),
		Fdy:   float32(-libc.Float64FromFloat64(0.496764600276947)),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.820243714273578),
		Fy:    libc.Float64FromFloat64(0.5989704010652284),
		Fdy:   float32(0.4980434477329254),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8177448070476216),
		Fy:    libc.Float64FromFloat64(0.5975966157416499),
		Fdy:   float32(0.4903089702129364),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8164966402061462),
		Fy:    libc.Float64FromFloat64(0.5969097230799119),
		Fdy:   float32(0.49379226565361023),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8158728782205076),
		Fy:    libc.Float64FromFloat64(0.596566276749022),
		Fdy:   float32(0.49634501338005066),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8140028771213308),
		Fy:    libc.Float64FromFloat64(0.5955359377563678),
		Fdy:   float32(0.4908742308616638),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8115125375794967),
		Fy:    libc.Float64FromFloat64(0.5941621524328238),
		Fdy:   float32(0.4953368008136749),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.810890487071702),
		Fy:    libc.Float64FromFloat64(0.5938187061019344),
		Fdy:   float32(-libc.Float64FromFloat64(0.4974379539489746)),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.810268650168187),
		Fy:    libc.Float64FromFloat64(0.5934752597710442),
		Fdy:   float32(-libc.Float64FromFloat64(0.49774107336997986)),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8077834371303876),
		Fy:    libc.Float64FromFloat64(0.592101474447509),
		Fdy:   float32(0.4933202564716339),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8046817181350279),
		Fy:    libc.Float64FromFloat64(0.590384242793077),
		Fdy:   float32(0.4926154911518097),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8028232417862318),
		Fy:    libc.Float64FromFloat64(0.5893539038004179),
		Fdy:   float32(0.4982609450817108),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8003482501185593),
		Fy:    libc.Float64FromFloat64(0.5879801184768579),
		Fdy:   float32(-libc.Float64FromFloat64(0.49070197343826294)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7954084554370604),
		Fy:    libc.Float64FromFloat64(0.5852325478297848),
		Fdy:   float32(-libc.Float64FromFloat64(0.49442774057388306)),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7947919348675285),
		Fy:    libc.Float64FromFloat64(0.5848891014989022),
		Fdy:   float32(0.4973377287387848),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3233283780039616e+300),
		Fy:    libc.Float64FromFloat64(691.0556779596564),
		Fdy:   float32(0.49036815762519836),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.810926104414746e+299),
		Fy:    libc.Float64FromFloat64(690.648935359327),
		Fdy:   float32(0.49113383889198303),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.866451600965818e+299),
		Fy:    libc.Float64FromFloat64(690.2421927589975),
		Fdy:   float32(0.4902323782444),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.676180134646991e+298),
		Fy:    libc.Float64FromFloat64(688.2084797573505),
		Fdy:   float32(0.4909690320491791),
		Fe:    int32(FE_INEXACT),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.964671801661398e+297),
		Fy:    libc.Float64FromFloat64(684.9545389547153),
		Fdy:   float32(0.49072694778442383),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9739245830789938e+297),
		Fy:    libc.Float64FromFloat64(684.5477963543859),
		Fdy:   float32(0.49079570174217224),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.750611883692923e+296),
		Fy:    libc.Float64FromFloat64(683.734311153727),
		Fdy:   float32(0.491062730550766),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9443774685465656e+294),
		Fy:    libc.Float64FromFloat64(678.0399147491154),
		Fdy:   float32(0.49021828174591064),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8526818885266785e+293),
		Fy:    libc.Float64FromFloat64(676.0062017474684),
		Fdy:   float32(0.4909018576145172),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.041187107544512e+292),
		Fy:    libc.Float64FromFloat64(673.9724887458214),
		Fdy:   float32(0.49108728766441345),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4879715112641626e+292),
		Fy:    libc.Float64FromFloat64(672.7522609448332),
		Fdy:   float32(0.49043917655944824),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.296337890061707e+291),
		Fy:    libc.Float64FromFloat64(670.3118053428567),
		Fdy:   float32(0.4907771050930023),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8263087805198706e+290),
		Fy:    libc.Float64FromFloat64(669.0915775418686),
		Fdy:   float32(0.4906221330165863),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5476158811530418e+290),
		Fy:    libc.Float64FromFloat64(668.6848349415392),
		Fdy:   float32(0.4905865788459778),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.129384475770039e+290),
		Fy:    libc.Float64FromFloat64(667.8713497408803),
		Fdy:   float32(0.49118268489837646),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.57213518172034e+287),
		Fy:    libc.Float64FromFloat64(662.9904385369275),
		Fdy:   float32(0.49074938893318176),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.707460891688494e+287),
		Fy:    libc.Float64FromFloat64(662.5836959365981),
		Fdy:   float32(0.4904578626155853),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1216533991791074e+287),
		Fy:    libc.Float64FromFloat64(660.9567255352805),
		Fdy:   float32(0.49064159393310547),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.468142737057775e+286),
		Fy:    libc.Float64FromFloat64(660.5499829349511),
		Fdy:   float32(0.49016574025154114),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.771973387025498e+285),
		Fy:    libc.Float64FromFloat64(658.5162699333041),
		Fdy:   float32(0.4911714196205139),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1139752448182877e+284),
		Fy:    libc.Float64FromFloat64(654.0421013296806),
		Fdy:   float32(0.491011381149292),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.938367503505401e+283),
		Fy:    libc.Float64FromFloat64(653.2286161290218),
		Fdy:   float32(0.49048542976379395),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1892294028182166e+283),
		Fy:    libc.Float64FromFloat64(652.415130928363),
		Fdy:   float32(0.4909510016441345),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.705080423361042e+282),
		Fy:    libc.Float64FromFloat64(651.6016457277042),
		Fdy:   float32(0.4902869462966919),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.904562454628899e+280),
		Fy:    libc.Float64FromFloat64(646.313991923422),
		Fdy:   float32(0.4911923110485077),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.397298674693025e+278),
		Fy:    libc.Float64FromFloat64(642.246565920128),
		Fdy:   float32(0.4900752007961273),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.572665086637977e+276),
		Fy:    libc.Float64FromFloat64(637.7723973165045),
		Fdy:   float32(0.49091583490371704),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.552779098149713e+275),
		Fy:    libc.Float64FromFloat64(634.9251991141987),
		Fdy:   float32(0.49115467071533203),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.697126663431185e+275),
		Fy:    libc.Float64FromFloat64(634.5184565138693),
		Fdy:   float32(0.4904699921607971),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0912547047672801e+275),
		Fy:    libc.Float64FromFloat64(633.298228712881),
		Fdy:   float32(0.49042338132858276),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.837644998968694e+274),
		Fy:    libc.Float64FromFloat64(632.4847435122223),
		Fdy:   float32(0.49086105823516846),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.329999284992692e+273),
		Fy:    libc.Float64FromFloat64(630.4510305105753),
		Fdy:   float32(0.49038389325141907),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.806154168292663e+273),
		Fy:    libc.Float64FromFloat64(629.6375453099164),
		Fdy:   float32(0.49100491404533386),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2439971730954586e+273),
		Fy:    libc.Float64FromFloat64(628.8240601092576),
		Fdy:   float32(0.4903148114681244),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.514768162616216e+272),
		Fy:    libc.Float64FromFloat64(628.0105749085988),
		Fdy:   float32(0.49008092284202576),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.671818391522856e+272),
		Fy:    libc.Float64FromFloat64(627.6038323082694),
		Fdy:   float32(0.49023473262786865),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0837846413090249e+272),
		Fy:    libc.Float64FromFloat64(626.3836045072812),
		Fdy:   float32(0.49055635929107666),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.216006658963816e+271),
		Fy:    libc.Float64FromFloat64(625.9768619069517),
		Fdy:   float32(0.49140506982803345),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.786944950115442e+270),
		Fy:    libc.Float64FromFloat64(622.7229211043166),
		Fdy:   float32(0.49033427238464355),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.477017426740952e+269),
		Fy:    libc.Float64FromFloat64(621.095950702999),
		Fdy:   float32(0.4904177784919739),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.408411364891969e+268),
		Fy:    libc.Float64FromFloat64(617.4352673000343),
		Fdy:   float32(0.4904983937740326),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.377421860625968e+267),
		Fy:    libc.Float64FromFloat64(617.0285246997049),
		Fdy:   float32(0.490109384059906),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.243633283865956e+267),
		Fy:    libc.Float64FromFloat64(616.6217820993755),
		Fdy:   float32(0.49068447947502136),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4113977540785127e+266),
		Fy:    libc.Float64FromFloat64(613.3678412967403),
		Fdy:   float32(0.4907104969024658),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.200893182368507e+264),
		Fy:    libc.Float64FromFloat64(609.7071578937756),
		Fdy:   float32(0.4900393784046173),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.128651098197378e+264),
		Fy:    libc.Float64FromFloat64(609.3004152934462),
		Fdy:   float32(0.4903954565525055),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.113792851496236e+263),
		Fy:    libc.Float64FromFloat64(607.6734448921286),
		Fdy:   float32(0.49025610089302063),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.596928310066665e+263),
		Fy:    libc.Float64FromFloat64(606.8599596914698),
		Fdy:   float32(0.49104443192481995),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.133685303029817e+262),
		Fy:    libc.Float64FromFloat64(604.4195040894934),
		Fdy:   float32(0.49046790599823),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7301026686970786e+261),
		Fy:    libc.Float64FromFloat64(601.979048487517),
		Fdy:   float32(0.49031904339790344),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.058250779399749e+260),
		Fy:    libc.Float64FromFloat64(600.7588206865288),
		Fdy:   float32(0.4903261661529541),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3784968370690105e+260),
		Fy:    libc.Float64FromFloat64(599.5385928855405),
		Fdy:   float32(0.49025940895080566),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.020440736853479e+259),
		Fy:    libc.Float64FromFloat64(598.3183650845524),
		Fdy:   float32(0.49151530861854553),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.186160798092838e+258),
		Fy:    libc.Float64FromFloat64(596.2846520829054),
		Fdy:   float32(0.49154505133628845),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.116288694522055e+258),
		Fy:    libc.Float64FromFloat64(595.8779094825759),
		Fdy:   float32(0.49114421010017395),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8053016358975792e+258),
		Fy:    libc.Float64FromFloat64(594.6576816815877),
		Fdy:   float32(0.4909677505493164),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2019978997250075e+258),
		Fy:    libc.Float64FromFloat64(594.2509390812584),
		Fdy:   float32(0.49030864238739014),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.003088914417475e+257),
		Fy:    libc.Float64FromFloat64(593.8441964809289),
		Fdy:   float32(0.49068495631217957),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.090929552795538e+256),
		Fy:    libc.Float64FromFloat64(590.5902556782937),
		Fdy:   float32(0.4909314513206482),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0579889569595307e+256),
		Fy:    libc.Float64FromFloat64(590.1835130779643),
		Fdy:   float32(0.49078601598739624),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3702410471103538e+256),
		Fy:    libc.Float64FromFloat64(589.7767704776348),
		Fdy:   float32(0.4903731942176819),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.123277949752446e+255),
		Fy:    libc.Float64FromFloat64(589.3700278773055),
		Fdy:   float32(0.4902334213256836),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.074420316335468e+255),
		Fy:    libc.Float64FromFloat64(588.9632852769761),
		Fdy:   float32(0.49090662598609924),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.292104826346508e+254),
		Fy:    libc.Float64FromFloat64(586.5228296749997),
		Fdy:   float32(0.4905882179737091),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5620331180410878e+254),
		Fy:    libc.Float64FromFloat64(585.3026018740114),
		Fdy:   float32(0.4903557300567627),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.924654408089309e+253),
		Fy:    libc.Float64FromFloat64(584.4891166733527),
		Fdy:   float32(0.4910874664783478),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.499445500672079e+251),
		Fy:    libc.Float64FromFloat64(579.2014628690704),
		Fdy:   float32(0.4901328980922699),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3515455884713254e+250),
		Fy:    libc.Float64FromFloat64(575.9475220664352),
		Fdy:   float32(0.4901290237903595),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.991541414101042e+249),
		Fy:    libc.Float64FromFloat64(575.1340368657764),
		Fdy:   float32(0.4908254146575928),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6561122927044817e+249),
		Fy:    libc.Float64FromFloat64(574.3205516651176),
		Fdy:   float32(0.4903213381767273),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7684808642391826e+249),
		Fy:    libc.Float64FromFloat64(573.9138090647882),
		Fdy:   float32(0.49079379439353943),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.475490433312895e+248),
		Fy:    libc.Float64FromFloat64(572.2868386634706),
		Fdy:   float32(0.49096059799194336),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.9505269796334776e+246),
		Fy:    libc.Float64FromFloat64(568.2194126601765),
		Fdy:   float32(0.49084845185279846),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.786191738839783e+245),
		Fy:    libc.Float64FromFloat64(566.1856996585295),
		Fdy:   float32(0.4903239905834198),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3331052091661506e+244),
		Fy:    libc.Float64FromFloat64(562.1182736552355),
		Fdy:   float32(0.49097755551338196),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.909793304942806e+243),
		Fy:    libc.Float64FromFloat64(561.3047884545766),
		Fdy:   float32(0.4908464848995209),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.732892223947301e+242),
		Fy:    libc.Float64FromFloat64(559.2710754529296),
		Fdy:   float32(0.490185409784317),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.148679881056635e+242),
		Fy:    libc.Float64FromFloat64(558.8643328526002),
		Fdy:   float32(0.4901961386203766),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.679957564993364e+239),
		Fy:    libc.Float64FromFloat64(552.3564512473298),
		Fdy:   float32(0.49081486463546753),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5092964580573183e+239),
		Fy:    libc.Float64FromFloat64(550.7294808460122),
		Fdy:   float32(0.49013394117355347),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.584127180397495e+237),
		Fy:    libc.Float64FromFloat64(546.6620548427181),
		Fdy:   float32(0.49003514647483826),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7205520571722976e+237),
		Fy:    libc.Float64FromFloat64(546.2553122423888),
		Fdy:   float32(0.4911401867866516),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.3812987947097143e+236),
		Fy:    libc.Float64FromFloat64(544.6283418410711),
		Fdy:   float32(0.4906133711338043),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.566437823423874e+234),
		Fy:    libc.Float64FromFloat64(539.7474306371183),
		Fdy:   float32(0.49100151658058167),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.92566106090242e+232),
		Fy:    libc.Float64FromFloat64(535.2732620334948),
		Fdy:   float32(0.4912262260913849),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.523317833314286e+230),
		Fy:    libc.Float64FromFloat64(531.6125786305303),
		Fdy:   float32(0.49079424142837524),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.3351646259103464e+230),
		Fy:    libc.Float64FromFloat64(530.7990934298714),
		Fdy:   float32(0.49064502120018005),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.220604466238246e+230),
		Fy:    libc.Float64FromFloat64(530.392350829542),
		Fdy:   float32(0.4909760057926178),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.844169325323489e+229),
		Fy:    libc.Float64FromFloat64(529.5788656288832),
		Fdy:   float32(0.49068209528923035),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.576353509455381e+228),
		Fy:    libc.Float64FromFloat64(527.1384100269067),
		Fdy:   float32(0.4909501373767853),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2054035638219563e+227),
		Fy:    libc.Float64FromFloat64(523.4777266239422),
		Fdy:   float32(0.49085453152656555),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4683919845016583e+227),
		Fy:    libc.Float64FromFloat64(523.0709840236127),
		Fdy:   float32(0.49123919010162354),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.3341486711279006e+226),
		Fy:    libc.Float64FromFloat64(521.8507562226246),
		Fdy:   float32(0.4904187321662903),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9213729436463445e+226),
		Fy:    libc.Float64FromFloat64(521.0372710219657),
		Fdy:   float32(0.49047577381134033),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.517645029504567e+225),
		Fy:    libc.Float64FromFloat64(520.2237858213069),
		Fdy:   float32(0.4909400939941406),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.671180501078594e+225),
		Fy:    libc.Float64FromFloat64(519.8170432209776),
		Fdy:   float32(0.49088403582572937),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.420670289892081e+224),
		Fy:    libc.Float64FromFloat64(517.7833302193305),
		Fdy:   float32(0.49095475673675537),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4483573623670477e+221),
		Fy:    libc.Float64FromFloat64(509.24173561241304),
		Fdy:   float32(0.49185994267463684),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.420718023175229e+220),
		Fy:    libc.Float64FromFloat64(508.4282504117542),
		Fdy:   float32(0.49141135811805725),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.2750138952518947e+220),
		Fy:    libc.Float64FromFloat64(508.02150781142484),
		Fdy:   float32(0.4907272756099701),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.57737577280762e+217),
		Fy:    libc.Float64FromFloat64(501.9203688064838),
		Fdy:   float32(0.49356162548065186),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.376765707669469e+217),
		Fy:    libc.Float64FromFloat64(501.51362620615436),
		Fdy:   float32(0.4908583462238312),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.882184791914928e+217),
		Fy:    libc.Float64FromFloat64(500.29339840516616),
		Fdy:   float32(0.49057143926620483),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.269318958422813e+215),
		Fy:    libc.Float64FromFloat64(497.03945760253094),
		Fdy:   float32(0.4935648739337921),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.8400255927927514e+215),
		Fy:    libc.Float64FromFloat64(496.6327150022015),
		Fdy:   float32(0.49156755208969116),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.222564296996089e+215),
		Fy:    libc.Float64FromFloat64(496.2259724018721),
		Fdy:   float32(0.49049490690231323),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.145633416430279e+215),
		Fy:    libc.Float64FromFloat64(495.81922980154275),
		Fdy:   float32(0.490043580532074),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.517481967170883e+213),
		Fy:    libc.Float64FromFloat64(492.1585463985781),
		Fdy:   float32(0.4933522343635559),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0843179692039527e+213),
		Fy:    libc.Float64FromFloat64(490.53157599726046),
		Fdy:   float32(0.4918052852153778),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.230076126774052e+210),
		Fy:    libc.Float64FromFloat64(485.6506647933076),
		Fdy:   float32(0.49202895164489746),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6174079925002306e+210),
		Fy:    libc.Float64FromFloat64(484.02369439199003),
		Fdy:   float32(0.4914403259754181),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0768953904021742e+210),
		Fy:    libc.Float64FromFloat64(483.6169517916606),
		Fdy:   float32(0.49090665578842163),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.170136955220665e+209),
		Fy:    libc.Float64FromFloat64(483.2102091913312),
		Fdy:   float32(0.49008500576019287),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.382036468058782e+208),
		Fy:    libc.Float64FromFloat64(481.1764961896842),
		Fdy:   float32(0.49315834045410156),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.246705761247024e+208),
		Fy:    libc.Float64FromFloat64(480.7697535893548),
		Fdy:   float32(0.4904451072216034),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.606336204869732e+207),
		Fy:    libc.Float64FromFloat64(477.10907018639017),
		Fdy:   float32(0.49243834614753723),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0695236220431802e+207),
		Fy:    libc.Float64FromFloat64(476.7023275860608),
		Fdy:   float32(0.49069085717201233),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3994585721480998e+206),
		Fy:    libc.Float64FromFloat64(474.6686145844138),
		Fdy:   float32(0.4913855493068695),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5953402079376247e+204),
		Fy:    libc.Float64FromFloat64(470.1944459807903),
		Fdy:   float32(0.4930509924888611),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0622023163095796e+204),
		Fy:    libc.Float64FromFloat64(469.7877033804609),
		Fdy:   float32(0.4900451600551605),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0874831406021094e+203),
		Fy:    libc.Float64FromFloat64(468.1607329791433),
		Fdy:   float32(0.4901706874370575),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.16147621025903e+202),
		Fy:    libc.Float64FromFloat64(466.9405051781551),
		Fdy:   float32(0.49135932326316833),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8186393150288044e+202),
		Fy:    libc.Float64FromFloat64(465.7202773771669),
		Fdy:   float32(0.4911947548389435),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.190681286402212e+199),
		Fy:    libc.Float64FromFloat64(460.4326235728846),
		Fdy:   float32(0.49168118834495544),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.077409524409759e+196),
		Fy:    libc.Float64FromFloat64(453.11125676695536),
		Fdy:   float32(0.4902292788028717),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.361530890113813e+194),
		Fy:    libc.Float64FromFloat64(447.0101177620143),
		Fdy:   float32(0.49123239517211914),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5414790489565964e+189),
		Fy:    libc.Float64FromFloat64(435.62132495279104),
		Fdy:   float32(0.49081215262413025),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0263407191754986e+189),
		Fy:    libc.Float64FromFloat64(435.2145823524616),
		Fdy:   float32(0.49097388982772827),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1699972038694318e+187),
		Fy:    libc.Float64FromFloat64(430.7404137488382),
		Fdy:   float32(0.4923788011074066),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5309270245162277e+186),
		Fy:    libc.Float64FromFloat64(428.7067007471912),
		Fdy:   float32(0.49079233407974243),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.736697562180802e+183),
		Fy:    libc.Float64FromFloat64(423.41904694290895),
		Fdy:   float32(0.4917660355567932),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.740299876630533e+182),
		Fy:    libc.Float64FromFloat64(420.9785913409325),
		Fdy:   float32(0.4908123314380646),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.98948659905785e+182),
		Fy:    libc.Float64FromFloat64(419.7583635399443),
		Fdy:   float32(0.49072110652923584),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.683736854204179e+180),
		Fy:    libc.Float64FromFloat64(416.5044227373091),
		Fdy:   float32(0.49187278747558594),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.8320290767315684e+178),
		Fy:    libc.Float64FromFloat64(411.62351153335624),
		Fdy:   float32(0.4905877411365509),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9472760241315083e+176),
		Fy:    libc.Float64FromFloat64(406.335857729074),
		Fdy:   float32(0.4917433559894562),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.359805026575628e+174),
		Fy:    libc.Float64FromFloat64(401.86168912545054),
		Fdy:   float32(0.49024292826652527),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.916898651638154e+173),
		Fy:    libc.Float64FromFloat64(400.64146132446234),
		Fdy:   float32(0.4935462176799774),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.92710077188906e+173),
		Fy:    libc.Float64FromFloat64(399.42123352347414),
		Fdy:   float32(0.4918329119682312),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.639716134820248e+172),
		Fy:    libc.Float64FromFloat64(398.20100572248595),
		Fdy:   float32(0.49057161808013916),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2216972205184255e+171),
		Fy:    libc.Float64FromFloat64(394.5403223195213),
		Fdy:   float32(0.4926331639289856),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.3661696271457796e+170),
		Fy:    libc.Float64FromFloat64(392.91335191820366),
		Fdy:   float32(0.49110323190689087),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6862892412391378e+169),
		Fy:    libc.Float64FromFloat64(389.65941111556845),
		Fdy:   float32(0.4921484589576721),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2064888377733743e+168),
		Fy:    libc.Float64FromFloat64(387.62569811392143),
		Fdy:   float32(0.49253004789352417),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.777818299708299e+166),
		Fy:    libc.Float64FromFloat64(383.5582721106274),
		Fdy:   float32(0.4912410080432892),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.909159428350362e+164),
		Fy:    libc.Float64FromFloat64(378.27061830634517),
		Fdy:   float32(0.4908115267753601),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.847771415830941e+161),
		Fy:    libc.Float64FromFloat64(371.7627367010747),
		Fdy:   float32(0.49104949831962585),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.726274050674535e+160),
		Fy:    libc.Float64FromFloat64(369.72902369942767),
		Fdy:   float32(0.4911133050918579),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6518958358008406e+160),
		Fy:    libc.Float64FromFloat64(368.91553849876885),
		Fdy:   float32(0.49098697304725647),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.323025132417057e+159),
		Fy:    libc.Float64FromFloat64(368.1020532981101),
		Fdy:   float32(0.4926440417766571),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.875783998512716e+159),
		Fy:    libc.Float64FromFloat64(367.69531069778066),
		Fdy:   float32(0.49215391278266907),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.246372799530095e+159),
		Fy:    libc.Float64FromFloat64(367.2885680974513),
		Fdy:   float32(0.4916616976261139),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.161485487614506e+159),
		Fy:    libc.Float64FromFloat64(366.88182549712184),
		Fdy:   float32(0.4925493001937866),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8282773113857515e+158),
		Fy:    libc.Float64FromFloat64(364.8481124954748),
		Fdy:   float32(0.49106383323669434),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.842407358101619e+156),
		Fy:    libc.Float64FromFloat64(360.7806864921808),
		Fdy:   float32(0.4911682605743408),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1466892776314716e+156),
		Fy:    libc.Float64FromFloat64(359.967201291522),
		Fdy:   float32(0.49153849482536316),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.336230773643207e+155),
		Fy:    libc.Float64FromFloat64(358.7469734905338),
		Fdy:   float32(0.4903601109981537),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0848515595483257e+154),
		Fy:    libc.Float64FromFloat64(354.67954748723974),
		Fdy:   float32(0.49053096771240234),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.809259193793128e+153),
		Fy:    libc.Float64FromFloat64(353.8660622865809),
		Fdy:   float32(0.492278516292572),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2020795501569484e+153),
		Fy:    libc.Float64FromFloat64(353.45931968625155),
		Fdy:   float32(0.4900107979774475),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7705920188017334e+149),
		Fy:    libc.Float64FromFloat64(344.1042398786753),
		Fdy:   float32(0.4901861250400543),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2282321043276842e+149),
		Fy:    libc.Float64FromFloat64(343.29075467801647),
		Fdy:   float32(0.4902476370334625),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts,
		Fline: int32(448),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.625285754072531e+148),
		Fy:    libc.Float64FromFloat64(342.07052687702827),
		Fdy:   float32(0.4903561770915985),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts,
		Fline: int32(449),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4001472424482978e+147),
		Fy:    libc.Float64FromFloat64(338.816586074393),
		Fdy:   float32(0.49177002906799316),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts,
		Fline: int32(450),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8320755337507494e+146),
		Fy:    libc.Float64FromFloat64(336.782873072746),
		Fdy:   float32(0.4915969967842102),
		Fe:    int32(FE_INEXACT),
	},
	434: {
		Ffile: __ccgo_ts,
		Fline: int32(451),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.407607657778202e+145),
		Fy:    libc.Float64FromFloat64(335.5626452717578),
		Fdy:   float32(0.4911704659461975),
		Fe:    int32(FE_INEXACT),
	},
	435: {
		Ffile: __ccgo_ts,
		Fline: int32(452),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0627250235927295e+145),
		Fy:    libc.Float64FromFloat64(333.9356748704402),
		Fdy:   float32(0.49139976501464844),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts,
		Fline: int32(453),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.136770220345118e+144),
		Fy:    libc.Float64FromFloat64(332.71544706945195),
		Fdy:   float32(0.4906404912471771),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts,
		Fline: int32(454),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3905626892671775e+144),
		Fy:    libc.Float64FromFloat64(331.9019618687932),
		Fdy:   float32(0.49039191007614136),
		Fe:    int32(FE_INEXACT),
	},
	438: {
		Ffile: __ccgo_ts,
		Fline: int32(455),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.381043746085339e+141),
		Fy:    libc.Float64FromFloat64(324.98733766319333),
		Fdy:   float32(0.49160224199295044),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts,
		Fline: int32(456),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.195203999233609e+140),
		Fy:    libc.Float64FromFloat64(324.5805950628639),
		Fdy:   float32(0.4916040301322937),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts,
		Fline: int32(457),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2031811915285054e+140),
		Fy:    libc.Float64FromFloat64(322.5468820612169),
		Fdy:   float32(0.49072182178497314),
		Fe:    int32(FE_INEXACT),
	},
	441: {
		Ffile: __ccgo_ts,
		Fline: int32(458),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.010967455235522e+139),
		Fy:    libc.Float64FromFloat64(322.1401394608875),
		Fdy:   float32(0.4900331497192383),
		Fe:    int32(FE_INEXACT),
	},
	442: {
		Ffile: __ccgo_ts,
		Fline: int32(459),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.364540553000867e+139),
		Fy:    libc.Float64FromFloat64(320.9199116598993),
		Fdy:   float32(0.49115145206451416),
		Fe:    int32(FE_INEXACT),
	},
	443: {
		Ffile: __ccgo_ts,
		Fline: int32(460),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5743478663101007e+139),
		Fy:    libc.Float64FromFloat64(320.51316905956986),
		Fdy:   float32(0.49234721064567566),
		Fe:    int32(FE_INEXACT),
	},
	444: {
		Ffile: __ccgo_ts,
		Fline: int32(461),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.6468911466963096e+138),
		Fy:    libc.Float64FromFloat64(319.29294125858166),
		Fdy:   float32(0.4915231168270111),
		Fe:    int32(FE_INEXACT),
	},
	445: {
		Ffile: __ccgo_ts,
		Fline: int32(462),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.048424999304951e+137),
		Fy:    libc.Float64FromFloat64(316.85248565660527),
		Fdy:   float32(0.49163374304771423),
		Fe:    int32(FE_INEXACT),
	},
	446: {
		Ffile: __ccgo_ts,
		Fline: int32(463),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.045913314038611e+135),
		Fy:    libc.Float64FromFloat64(311.564831852323),
		Fdy:   float32(0.4922950267791748),
		Fe:    int32(FE_INEXACT),
	},
	447: {
		Ffile: __ccgo_ts,
		Fline: int32(464),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3622008962901964e+135),
		Fy:    libc.Float64FromFloat64(311.1580892519936),
		Fdy:   float32(0.4900952875614166),
		Fe:    int32(FE_INEXACT),
	},
	448: {
		Ffile: __ccgo_ts,
		Fline: int32(465),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.901666459593163e+133),
		Fy:    libc.Float64FromFloat64(308.31089104968777),
		Fdy:   float32(0.4918755292892456),
		Fe:    int32(FE_INEXACT),
	},
	449: {
		Ffile: __ccgo_ts,
		Fline: int32(466),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0339233873121812e+133),
		Fy:    libc.Float64FromFloat64(306.27717804804075),
		Fdy:   float32(0.49183204770088196),
		Fe:    int32(FE_INEXACT),
	},
	450: {
		Ffile: __ccgo_ts,
		Fline: int32(467),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.051758469444187e+132),
		Fy:    libc.Float64FromFloat64(305.05695024705255),
		Fdy:   float32(0.4910182058811188),
		Fe:    int32(FE_INEXACT),
	},
	451: {
		Ffile: __ccgo_ts,
		Fline: int32(468),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0268457790503162e+130),
		Fy:    libc.Float64FromFloat64(299.3625538424409),
		Fdy:   float32(0.4918397068977356),
		Fe:    int32(FE_INEXACT),
	},
	452: {
		Ffile: __ccgo_ts,
		Fline: int32(469),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.965853682959888e+128),
		Fy:    libc.Float64FromFloat64(296.1086130398057),
		Fdy:   float32(0.49297165870666504),
		Fe:    int32(FE_INEXACT),
	},
	453: {
		Ffile: __ccgo_ts,
		Fline: int32(470),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.170572956765408e+128),
		Fy:    libc.Float64FromFloat64(294.8883852388175),
		Fdy:   float32(0.49033424258232117),
		Fe:    int32(FE_INEXACT),
	},
	454: {
		Ffile: __ccgo_ts,
		Fline: int32(471),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1625599289805774e+125),
		Fy:    libc.Float64FromFloat64(287.9737610332176),
		Fdy:   float32(0.49182572960853577),
		Fe:    int32(FE_INEXACT),
	},
	455: {
		Ffile: __ccgo_ts,
		Fline: int32(472),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5211954413308598e+124),
		Fy:    libc.Float64FromFloat64(285.9400480315706),
		Fdy:   float32(0.4913940131664276),
		Fe:    int32(FE_INEXACT),
	},
	456: {
		Ffile: __ccgo_ts,
		Fline: int32(473),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.490004896619718e+123),
		Fy:    libc.Float64FromFloat64(284.7198202305824),
		Fdy:   float32(0.4917212426662445),
		Fe:    int32(FE_INEXACT),
	},
	457: {
		Ffile: __ccgo_ts,
		Fline: int32(474),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.989515075029701e+123),
		Fy:    libc.Float64FromFloat64(284.31307763025296),
		Fdy:   float32(0.49231135845184326),
		Fe:    int32(FE_INEXACT),
	},
	458: {
		Ffile: __ccgo_ts,
		Fline: int32(475),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.911743894280602e+122),
		Fy:    libc.Float64FromFloat64(282.27936462860595),
		Fdy:   float32(0.4921264946460724),
		Fe:    int32(FE_INEXACT),
	},
	459: {
		Ffile: __ccgo_ts,
		Fline: int32(476),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.4592690644080663e+120),
		Fy:    libc.Float64FromFloat64(277.80519602498254),
		Fdy:   float32(0.49162939190864563),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts,
		Fline: int32(477),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.763536708996752e+119),
		Fy:    libc.Float64FromFloat64(276.1782256236649),
		Fdy:   float32(0.4916137754917145),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts,
		Fline: int32(478),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.8348990046418566e+119),
		Fy:    libc.Float64FromFloat64(275.7714830233355),
		Fdy:   float32(0.49157142639160156),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts,
		Fline: int32(479),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.884966483841679e+119),
		Fy:    libc.Float64FromFloat64(275.3647404230061),
		Fdy:   float32(0.4925258755683899),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts,
		Fline: int32(480),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.634893948452487e+118),
		Fy:    libc.Float64FromFloat64(273.7377700216885),
		Fdy:   float32(0.4919740557670593),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts,
		Fline: int32(481),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.99016530667343e+117),
		Fy:    libc.Float64FromFloat64(271.7040570200415),
		Fdy:   float32(0.4912319779396057),
		Fe:    int32(FE_INEXACT),
	},
	465: {
		Ffile: __ccgo_ts,
		Fline: int32(482),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.5826301258606555e+115),
		Fy:    libc.Float64FromFloat64(266.82314581608864),
		Fdy:   float32(0.4904383718967438),
		Fe:    int32(FE_INEXACT),
	},
	466: {
		Ffile: __ccgo_ts,
		Fline: int32(483),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.3614583787952473e+115),
		Fy:    libc.Float64FromFloat64(266.0096606154298),
		Fdy:   float32(0.4903961420059204),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts,
		Fline: int32(484),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.238111255749255e+115),
		Fy:    libc.Float64FromFloat64(265.60291801510044),
		Fdy:   float32(0.4927603304386139),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts,
		Fline: int32(485),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.921778734341683e+114),
		Fy:    libc.Float64FromFloat64(264.7894328144416),
		Fdy:   float32(0.49369966983795166),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts,
		Fline: int32(486),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5513787597541663e+113),
		Fy:    libc.Float64FromFloat64(261.128749411477),
		Fdy:   float32(0.49192312359809875),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts,
		Fline: int32(487),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.338447887026099e+112),
		Fy:    libc.Float64FromFloat64(259.09503640982996),
		Fdy:   float32(0.4920327365398407),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts,
		Fline: int32(488),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.4791733302895e+109),
		Fy:    libc.Float64FromFloat64(252.9938974048889),
		Fdy:   float32(0.4935656487941742),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts,
		Fline: int32(489),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.9797498965969464e+109),
		Fy:    libc.Float64FromFloat64(252.58715480455953),
		Fdy:   float32(0.49489355087280273),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts,
		Fline: int32(490),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.78640678329264e+108),
		Fy:    libc.Float64FromFloat64(250.9601844032419),
		Fdy:   float32(0.49332720041275024),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts,
		Fline: int32(491),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.3384153695964095e+108),
		Fy:    libc.Float64FromFloat64(250.1466992025831),
		Fdy:   float32(0.4912772476673126),
		Fe:    int32(FE_INEXACT),
	},
	475: {
		Ffile: __ccgo_ts,
		Fline: int32(492),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4597759143266845e+106),
		Fy:    libc.Float64FromFloat64(244.45230279797144),
		Fdy:   float32(0.4900543987751007),
		Fe:    int32(FE_INEXACT),
	},
	476: {
		Ffile: __ccgo_ts,
		Fline: int32(493),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8688109255882514e+105),
		Fy:    libc.Float64FromFloat64(242.82533239665383),
		Fdy:   float32(0.49288392066955566),
		Fe:    int32(FE_INEXACT),
	},
	477: {
		Ffile: __ccgo_ts,
		Fline: int32(494),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.377128090916262e+103),
		Fy:    libc.Float64FromFloat64(239.1646489936892),
		Fdy:   float32(0.49112117290496826),
		Fe:    int32(FE_INEXACT),
	},
	478: {
		Ffile: __ccgo_ts,
		Fline: int32(495),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.2792223792295477e+102),
		Fy:    libc.Float64FromFloat64(236.31745079138338),
		Fdy:   float32(0.49308812618255615),
		Fe:    int32(FE_INEXACT),
	},
	479: {
		Ffile: __ccgo_ts,
		Fline: int32(496),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.849172797504646e+102),
		Fy:    libc.Float64FromFloat64(235.91070819105397),
		Fdy:   float32(0.49005332589149475),
		Fe:    int32(FE_INEXACT),
	},
	480: {
		Ffile: __ccgo_ts,
		Fline: int32(497),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8970235502232178e+102),
		Fy:    libc.Float64FromFloat64(235.5039655907246),
		Fdy:   float32(0.4927107095718384),
		Fe:    int32(FE_INEXACT),
	},
	481: {
		Ffile: __ccgo_ts,
		Fline: int32(498),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2630677764624174e+102),
		Fy:    libc.Float64FromFloat64(235.09722299039518),
		Fdy:   float32(0.4960947036743164),
		Fe:    int32(FE_INEXACT),
	},
	482: {
		Ffile: __ccgo_ts,
		Fline: int32(499),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1003991702457248e+101),
		Fy:    libc.Float64FromFloat64(232.65676738841876),
		Fdy:   float32(0.49069294333457947),
		Fe:    int32(FE_INEXACT),
	},
	483: {
		Ffile: __ccgo_ts,
		Fline: int32(500),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.32662877600973e+100),
		Fy:    libc.Float64FromFloat64(232.25002478808935),
		Fdy:   float32(0.4924061894416809),
		Fe:    int32(FE_INEXACT),
	},
	484: {
		Ffile: __ccgo_ts,
		Fline: int32(501),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1477456584197126e+97),
		Fy:    libc.Float64FromFloat64(224.1151727815013),
		Fdy:   float32(0.49296891689300537),
		Fe:    int32(FE_INEXACT),
	},
	485: {
		Ffile: __ccgo_ts,
		Fline: int32(502),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4300024545650976e+97),
		Fy:    libc.Float64FromFloat64(223.70843018117188),
		Fdy:   float32(0.49351072311401367),
		Fe:    int32(FE_INEXACT),
	},
	486: {
		Ffile: __ccgo_ts,
		Fline: int32(503),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2458345813080797e+96),
		Fy:    libc.Float64FromFloat64(221.26797457919545),
		Fdy:   float32(0.4956727623939514),
		Fe:    int32(FE_INEXACT),
	},
	487: {
		Ffile: __ccgo_ts,
		Fline: int32(504),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.677241739271557e+95),
		Fy:    libc.Float64FromFloat64(220.04774677820726),
		Fdy:   float32(0.4936850965023041),
		Fe:    int32(FE_INEXACT),
	},
	488: {
		Ffile: __ccgo_ts,
		Fline: int32(505),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.456002516550464e+93),
		Fy:    libc.Float64FromFloat64(216.38706337524263),
		Fdy:   float32(0.4924456775188446),
		Fe:    int32(FE_INEXACT),
	},
	489: {
		Ffile: __ccgo_ts,
		Fline: int32(506),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.238178122314902e+92),
		Fy:    libc.Float64FromFloat64(213.9466077732662),
		Fdy:   float32(0.49100789427757263),
		Fe:    int32(FE_INEXACT),
	},
	490: {
		Ffile: __ccgo_ts,
		Fline: int32(507),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.778688366343745e+91),
		Fy:    libc.Float64FromFloat64(211.09940957096038),
		Fdy:   float32(0.4931267201900482),
		Fe:    int32(FE_INEXACT),
	},
	491: {
		Ffile: __ccgo_ts,
		Fline: int32(508),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.228836514482408e+90),
		Fy:    libc.Float64FromFloat64(207.43872616799575),
		Fdy:   float32(0.49367088079452515),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts,
		Fline: int32(509),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.447559394580902e+89),
		Fy:    libc.Float64FromFloat64(206.62524096733694),
		Fdy:   float32(0.49231991171836853),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts,
		Fline: int32(510),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.128064808394073e+88),
		Fy:    libc.Float64FromFloat64(204.59152796568992),
		Fdy:   float32(0.4932084381580353),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts,
		Fline: int32(511),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.7459763833500766e+88),
		Fy:    libc.Float64FromFloat64(204.1847853653605),
		Fdy:   float32(0.49290353059768677),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts,
		Fline: int32(512),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.103940428924214e+88),
		Fy:    libc.Float64FromFloat64(203.37130016470172),
		Fdy:   float32(0.4904153645038605),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts,
		Fline: int32(513),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.2100520231647115e+87),
		Fy:    libc.Float64FromFloat64(202.1510723637135),
		Fdy:   float32(0.4902549982070923),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts,
		Fline: int32(514),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.832977093849189e+87),
		Fy:    libc.Float64FromFloat64(200.9308445627253),
		Fdy:   float32(0.49360084533691406),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts,
		Fline: int32(515),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5969107206584518e+86),
		Fy:    libc.Float64FromFloat64(198.49038896074887),
		Fdy:   float32(0.49570971727371216),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts,
		Fline: int32(516),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0895381372796304e+85),
		Fy:    libc.Float64FromFloat64(196.45667595910186),
		Fdy:   float32(0.4928642213344574),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts,
		Fline: int32(517),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.263138729945093e+84),
		Fy:    libc.Float64FromFloat64(195.64319075844304),
		Fdy:   float32(0.4932849109172821),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts,
		Fline: int32(518),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.106445228222621e+84),
		Fy:    libc.Float64FromFloat64(194.82970555778425),
		Fdy:   float32(0.49273058772087097),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts,
		Fline: int32(519),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.734135083861001e+84),
		Fy:    libc.Float64FromFloat64(194.42296295745484),
		Fdy:   float32(0.4908994138240814),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts,
		Fline: int32(520),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.070152973779008e+83),
		Fy:    libc.Float64FromFloat64(193.20273515646662),
		Fdy:   float32(0.4900767505168915),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts,
		Fline: int32(521),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0559696415237794e+83),
		Fy:    libc.Float64FromFloat64(191.1690221548196),
		Fdy:   float32(0.4917001724243164),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts,
		Fline: int32(522),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.030810065454137e+82),
		Fy:    libc.Float64FromFloat64(190.76227955449022),
		Fdy:   float32(0.49180081486701965),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts,
		Fline: int32(523),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.078335011323747e+81),
		Fy:    libc.Float64FromFloat64(187.9150813521844),
		Fdy:   float32(0.4915797710418701),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts,
		Fline: int32(524),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.715418864365381e+81),
		Fy:    libc.Float64FromFloat64(187.50833875185498),
		Fdy:   float32(0.49458935856819153),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts,
		Fline: int32(525),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.050417219807457e+78),
		Fy:    libc.Float64FromFloat64(181.00045714658452),
		Fdy:   float32(0.49113717675209045),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts,
		Fline: int32(526),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.960044455682121e+77),
		Fy:    libc.Float64FromFloat64(179.37348674526692),
		Fdy:   float32(0.491471529006958),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts,
		Fline: int32(527),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3628712321750758e+76),
		Fy:    libc.Float64FromFloat64(175.30606074197289),
		Fdy:   float32(0.4936870038509369),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts,
		Fline: int32(528),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.074208575984287e+75),
		Fy:    libc.Float64FromFloat64(174.89931814164348),
		Fdy:   float32(0.4954296052455902),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts,
		Fline: int32(529),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5536317531665822e+74),
		Fy:    libc.Float64FromFloat64(170.83189213834945),
		Fdy:   float32(0.4953838586807251),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts,
		Fline: int32(530),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.000391038231264e+72),
		Fy:    libc.Float64FromFloat64(167.57795133571423),
		Fdy:   float32(0.4940897226333618),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts,
		Fline: int32(531),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.77109294514582e+72),
		Fy:    libc.Float64FromFloat64(166.357723534726),
		Fdy:   float32(0.49557459354400635),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts,
		Fline: int32(532),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.22760966803592e+71),
		Fy:    libc.Float64FromFloat64(165.1374957337378),
		Fdy:   float32(0.4934104382991791),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts,
		Fline: int32(533),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.317453446629227e+71),
		Fy:    libc.Float64FromFloat64(164.324010533079),
		Fdy:   float32(0.49094176292419434),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts,
		Fline: int32(534),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.027351087463935e+71),
		Fy:    libc.Float64FromFloat64(163.51052533242017),
		Fdy:   float32(0.4924500584602356),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts,
		Fline: int32(535),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.005171302774353e+67),
		Fy:    libc.Float64FromFloat64(154.9689307255027),
		Fdy:   float32(0.49152591824531555),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts,
		Fline: int32(536),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.8891318246461e+66),
		Fy:    libc.Float64FromFloat64(154.1554455248439),
		Fdy:   float32(0.494808554649353),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts,
		Fline: int32(537),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.9185221855131655e+66),
		Fy:    libc.Float64FromFloat64(153.7487029245145),
		Fdy:   float32(0.4904435873031616),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts,
		Fline: int32(538),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.746928295450148e+66),
		Fy:    libc.Float64FromFloat64(152.5284751235263),
		Fdy:   float32(0.4969553053379059),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts,
		Fline: int32(539),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1631320220109382e+66),
		Fy:    libc.Float64FromFloat64(152.1217325231969),
		Fdy:   float32(0.49033257365226746),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts,
		Fline: int32(540),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.521944018060633e+65),
		Fy:    libc.Float64FromFloat64(150.08801952154988),
		Fdy:   float32(0.49383580684661865),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts,
		Fline: int32(541),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7349698909024902e+63),
		Fy:    libc.Float64FromFloat64(145.61385091792644),
		Fdy:   float32(0.4970724880695343),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts,
		Fline: int32(542),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5879439869035713e+60),
		Fy:    libc.Float64FromFloat64(139.105969312656),
		Fdy:   float32(0.49259650707244873),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts,
		Fline: int32(543),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7230933462913268e+60),
		Fy:    libc.Float64FromFloat64(138.69922671232658),
		Fdy:   float32(0.4962015450000763),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts,
		Fline: int32(544),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.638651061962569e+59),
		Fy:    libc.Float64FromFloat64(137.8857415116678),
		Fdy:   float32(0.49081578850746155),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts,
		Fline: int32(545),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2546465588783374e+59),
		Fy:    libc.Float64FromFloat64(136.66551371067956),
		Fdy:   float32(0.491875559091568),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts,
		Fline: int32(546),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5011787362871763e+59),
		Fy:    libc.Float64FromFloat64(136.25877111035015),
		Fdy:   float32(0.49128302931785583),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts,
		Fline: int32(547),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9501774331629635e+58),
		Fy:    libc.Float64FromFloat64(134.63180070903255),
		Fdy:   float32(0.4900747537612915),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts,
		Fline: int32(548),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8602710712555906e+57),
		Fy:    libc.Float64FromFloat64(132.59808770738553),
		Fdy:   float32(0.49361515045166016),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts,
		Fline: int32(549),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.400591611697199e+55),
		Fy:    libc.Float64FromFloat64(128.1239191037621),
		Fdy:   float32(0.49245685338974),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts,
		Fline: int32(550),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9508277865679163e+55),
		Fy:    libc.Float64FromFloat64(127.31043390310329),
		Fdy:   float32(0.49606770277023315),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts,
		Fline: int32(551),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.534429921244915e+53),
		Fy:    libc.Float64FromFloat64(124.05649310046806),
		Fdy:   float32(0.4960584044456482),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts,
		Fline: int32(552),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2238843425687973e+53),
		Fy:    libc.Float64FromFloat64(122.83626529947985),
		Fdy:   float32(0.49141696095466614),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts,
		Fline: int32(553),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.909925449545792e+52),
		Fy:    libc.Float64FromFloat64(120.80255229783283),
		Fdy:   float32(0.4900602102279663),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts,
		Fline: int32(554),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.589020951724632e+51),
		Fy:    libc.Float64FromFloat64(119.58232449684462),
		Fdy:   float32(0.49853235483169556),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts,
		Fline: int32(555),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.8076018432385386e+51),
		Fy:    libc.Float64FromFloat64(118.76883929618582),
		Fdy:   float32(0.49507227540016174),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts,
		Fline: int32(556),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.687949287598655e+51),
		Fy:    libc.Float64FromFloat64(117.95535409552701),
		Fdy:   float32(0.49879318475723267),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts,
		Fline: int32(557),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.482853814046525e+50),
		Fy:    libc.Float64FromFloat64(117.14186889486821),
		Fdy:   float32(0.49454930424690247),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts,
		Fline: int32(558),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9242108491518145e+49),
		Fy:    libc.Float64FromFloat64(113.48118549190357),
		Fdy:   float32(0.4943091571331024),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts,
		Fline: int32(559),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1935418434643832e+47),
		Fy:    libc.Float64FromFloat64(109.00701688828012),
		Fdy:   float32(-libc.Float64FromFloat64(0.49778175354003906)),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts,
		Fline: int32(560),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.724197918877798e+46),
		Fy:    libc.Float64FromFloat64(108.19353168762132),
		Fdy:   float32(-libc.Float64FromFloat64(0.49828293919563293)),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts,
		Fline: int32(561),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.911038876478935e+46),
		Fy:    libc.Float64FromFloat64(106.56656128630371),
		Fdy:   float32(0.4979231357574463),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts,
		Fline: int32(562),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.471833040668761e+45),
		Fy:    libc.Float64FromFloat64(105.7530760856449),
		Fdy:   float32(0.4962787330150604),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts,
		Fline: int32(563),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5005709853211116e+45),
		Fy:    libc.Float64FromFloat64(104.5328482846567),
		Fdy:   float32(0.4964560866355896),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts,
		Fline: int32(564),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.413840044033365e+42),
		Fy:    libc.Float64FromFloat64(98.83845188004506),
		Fdy:   float32(0.49011197686195374),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts,
		Fline: int32(565),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.602065527693206e+42),
		Fy:    libc.Float64FromFloat64(98.43170927971563),
		Fdy:   float32(-libc.Float64FromFloat64(0.49458232522010803)),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts,
		Fline: int32(566),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4405674003428797e+41),
		Fy:    libc.Float64FromFloat64(94.77102587675103),
		Fdy:   float32(0.49697065353393555),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts,
		Fline: int32(567),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.591521744583649e+40),
		Fy:    libc.Float64FromFloat64(94.36428327642162),
		Fdy:   float32(0.49144428968429565),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts,
		Fline: int32(568),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8310615736224536e+40),
		Fy:    libc.Float64FromFloat64(93.1440554754334),
		Fdy:   float32(0.49847331643104553),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts,
		Fline: int32(569),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8849648157519174e+40),
		Fy:    libc.Float64FromFloat64(92.73731287510401),
		Fdy:   float32(0.49753308296203613),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts,
		Fline: int32(570),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1488025590504006e+38),
		Fy:    libc.Float64FromFloat64(88.26314427148057),
		Fdy:   float32(0.49001264572143555),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts,
		Fline: int32(571),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.813993608875372e+35),
		Fy:    libc.Float64FromFloat64(82.16200526653951),
		Fdy:   float32(0.4925099313259125),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts,
		Fline: int32(572),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.420912414596516e+35),
		Fy:    libc.Float64FromFloat64(80.94177746555131),
		Fdy:   float32(0.4934936761856079),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts,
		Fline: int32(573),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.180727388070626e+32),
		Fy:    libc.Float64FromFloat64(75.65412366126907),
		Fdy:   float32(0.4915544390678406),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts,
		Fline: int32(574),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.781039947746529e+32),
		Fy:    libc.Float64FromFloat64(75.24738106093966),
		Fdy:   float32(0.4904050827026367),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts,
		Fline: int32(575),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1832907373593933e+32),
		Fy:    libc.Float64FromFloat64(74.84063846061025),
		Fdy:   float32(0.4950158894062042),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts,
		Fline: int32(576),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.165296977290502e+31),
		Fy:    libc.Float64FromFloat64(72.80692545896322),
		Fdy:   float32(-libc.Float64FromFloat64(0.49583566188812256)),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts,
		Fline: int32(577),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.229441222065934e+31),
		Fy:    libc.Float64FromFloat64(71.58669765797502),
		Fdy:   float32(0.49428167939186096),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts,
		Fline: int32(578),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1614998724857624e+29),
		Fy:    libc.Float64FromFloat64(67.9260142550104),
		Fdy:   float32(0.49463748931884766),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts,
		Fline: int32(579),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.597696556700635e+27),
		Fy:    libc.Float64FromFloat64(62.63836045072816),
		Fdy:   float32(0.4903721213340759),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts,
		Fline: int32(580),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.108466000344566e+25),
		Fy:    libc.Float64FromFloat64(58.97767704776353),
		Fdy:   float32(0.49079629778862),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts,
		Fline: int32(581),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.6835262820491106e+23),
		Fy:    libc.Float64FromFloat64(54.50350844414008),
		Fdy:   float32(-libc.Float64FromFloat64(0.49129071831703186)),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts,
		Fline: int32(582),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1183646225027014e+23),
		Fy:    libc.Float64FromFloat64(54.09676584381059),
		Fdy:   float32(-libc.Float64FromFloat64(0.49031969904899597)),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts,
		Fline: int32(583),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.128336829787126e+22),
		Fy:    libc.Float64FromFloat64(52.469795442493066),
		Fdy:   float32(-libc.Float64FromFloat64(0.49422401189804077)),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts,
		Fline: int32(584),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.080341950473528e+22),
		Fy:    libc.Float64FromFloat64(52.063052842163664),
		Fdy:   float32(-libc.Float64FromFloat64(0.49586549401283264)),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts,
		Fline: int32(585),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.366868295383857e+21),
		Fy:    libc.Float64FromFloat64(49.215854639857824),
		Fdy:   float32(-libc.Float64FromFloat64(0.49034780263900757)),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts,
		Fline: int32(586),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.986117593561772e+20),
		Fy:    libc.Float64FromFloat64(47.995626838869626),
		Fdy:   float32(-libc.Float64FromFloat64(0.49459126591682434)),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts,
		Fline: int32(587),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.651465720584657e+20),
		Fy:    libc.Float64FromFloat64(47.588884238540224),
		Fdy:   float32(-libc.Float64FromFloat64(0.4965655207633972)),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts,
		Fline: int32(588),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.141249385092951e+19),
		Fy:    libc.Float64FromFloat64(45.96191383722261),
		Fdy:   float32(0.49586251378059387),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts,
		Fline: int32(589),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.086385977422035e+19),
		Fy:    libc.Float64FromFloat64(45.555171236893216),
		Fdy:   float32(0.4994893968105316),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts,
		Fline: int32(590),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6981578688127054e+19),
		Fy:    libc.Float64FromFloat64(44.7416860362344),
		Fdy:   float32(0.4941146969795227),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts,
		Fline: int32(591),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7964754625765622e+19),
		Fy:    libc.Float64FromFloat64(44.334943435905004),
		Fdy:   float32(0.49390074610710144),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts,
		Fline: int32(592),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.075817906401565e+17),
		Fy:    libc.Float64FromFloat64(40.267517432610944),
		Fdy:   float32(-libc.Float64FromFloat64(0.494322806596756)),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts,
		Fline: int32(593),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6796879273659516e+16),
		Fy:    libc.Float64FromFloat64(37.82706183063455),
		Fdy:   float32(0.49544525146484375),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts,
		Fline: int32(594),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.909445156343236e+15),
		Fy:    libc.Float64FromFloat64(36.60683402964634),
		Fdy:   float32(0.4981684684753418),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts,
		Fline: int32(595),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.506338862782641e+15),
		Fy:    libc.Float64FromFloat64(35.79334882898753),
		Fdy:   float32(0.494545042514801),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts,
		Fline: int32(596),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.034941220263468e+15),
		Fy:    libc.Float64FromFloat64(34.57312102799932),
		Fdy:   float32(0.49557289481163025),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts,
		Fline: int32(597),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.5880014964074125e+14),
		Fy:    libc.Float64FromFloat64(33.75963582734052),
		Fdy:   float32(0.4995199143886566),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts,
		Fline: int32(598),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.997119658830151e+13),
		Fy:    libc.Float64FromFloat64(31.31918022536408),
		Fdy:   float32(-libc.Float64FromFloat64(0.4978776276111603)),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts,
		Fline: int32(599),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.661344419861769e+13),
		Fy:    libc.Float64FromFloat64(30.912437625034695),
		Fdy:   float32(0.49075886607170105),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts,
		Fline: int32(600),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1798015135268096e+13),
		Fy:    libc.Float64FromFloat64(30.098952424375884),
		Fdy:   float32(0.49194616079330444),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts,
		Fline: int32(601),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.027856644538637e+12),
		Fy:    libc.Float64FromFloat64(27.65849682239945),
		Fdy:   float32(0.49228301644325256),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts,
		Fline: int32(602),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.033851694272748e+11),
		Fy:    libc.Float64FromFloat64(26.438269021411255),
		Fdy:   float32(0.49145349860191345),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts,
		Fline: int32(603),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.96975781803216e+10),
		Fy:    libc.Float64FromFloat64(24.40455601976423),
		Fdy:   float32(0.49014368653297424),
		Fe:    int32(FE_INEXACT),
	},
	587: {
		Ffile: __ccgo_ts,
		Fline: int32(604),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1717253126767712e+10),
		Fy:    libc.Float64FromFloat64(23.184328218776027),
		Fdy:   float32(-libc.Float64FromFloat64(0.4972979426383972)),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts,
		Fline: int32(605),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.302722182086328e+09),
		Fy:    libc.Float64FromFloat64(21.557357817458417),
		Fdy:   float32(0.4949679374694824),
		Fe:    int32(FE_INEXACT),
	},
	589: {
		Ffile: __ccgo_ts,
		Fline: int32(606),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0130837978500843e+08),
		Fy:    libc.Float64FromFloat64(19.523644815811387),
		Fdy:   float32(0.49163344502449036),
		Fe:    int32(FE_INEXACT),
	},
	590: {
		Ffile: __ccgo_ts,
		Fline: int32(607),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.006158042897095e+08),
		Fy:    libc.Float64FromFloat64(19.116902215481996),
		Fdy:   float32(0.49820560216903687),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts,
		Fline: int32(608),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7477879547910742e+07),
		Fy:    libc.Float64FromFloat64(16.676446613505576),
		Fdy:   float32(0.492167204618454),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts,
		Fline: int32(609),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2869591519230707e+06),
		Fy:    libc.Float64FromFloat64(14.642733611858558),
		Fdy:   float32(0.49395832419395447),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts,
		Fline: int32(610),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.522692963163255e+06),
		Fy:    libc.Float64FromFloat64(14.235991011529155),
		Fdy:   float32(-libc.Float64FromFloat64(0.49964046478271484)),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts,
		Fline: int32(611),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(449442.50258006505),
		Fy:    libc.Float64FromFloat64(13.015763210540939),
		Fdy:   float32(-libc.Float64FromFloat64(0.4941127896308899)),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts,
		Fline: int32(612),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(88326.26632424773),
		Fy:    libc.Float64FromFloat64(11.388792809223329),
		Fdy:   float32(0.4972361922264099),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts,
		Fline: int32(613),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(11557.383867517256),
		Fy:    libc.Float64FromFloat64(9.355079807576313),
		Fdy:   float32(-libc.Float64FromFloat64(0.4935089945793152)),
		Fe:    int32(FE_INEXACT),
	},
	597: {
		Ffile: __ccgo_ts,
		Fline: int32(614),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5019174628430672),
		Fy:    libc.Float64FromFloat64(0.40674260032943543),
		Fdy:   float32(-libc.Float64FromFloat64(0.4933646023273468)),
		Fe:    int32(FE_INEXACT),
	},
	598: {
		Ffile: __ccgo_ts,
		Fline: int32(615),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6868921304654562),
		Fy:    -libc.Float64FromFloat64(0.37557801442211075),
		Fdy:   float32(-libc.Float64FromFloat64(0.44673043489456177)),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts,
		Fline: int32(616),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5242453321059702),
		Fy:    -libc.Float64FromFloat64(0.6457955131863953),
		Fdy:   float32(-libc.Float64FromFloat64(0.472727507352829)),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts,
		Fline: int32(617),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7550665897480207),
		Fy:    -libc.Float64FromFloat64(0.2809493352806082),
		Fdy:   float32(-libc.Float64FromFloat64(0.4986448585987091)),
		Fe:    int32(FE_INEXACT),
	},
	601: {
		Ffile: __ccgo_ts,
		Fline: int32(618),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7056230658427625),
		Fy:    -libc.Float64FromFloat64(0.3486740851373969),
		Fdy:   float32(0.49320825934410095),
		Fe:    int32(FE_INEXACT),
	},
	602: {
		Ffile: __ccgo_ts,
		Fline: int32(619),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5029332434841948),
		Fy:    -libc.Float64FromFloat64(0.6872978344212657),
		Fdy:   float32(-libc.Float64FromFloat64(0.4925854206085205)),
		Fe:    int32(FE_INEXACT),
	},
	603: {
		Ffile: __ccgo_ts,
		Fline: int32(620),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9699999999999963),
		Fy:    -libc.Float64FromFloat64(0.030459207484712352),
		Fdy:   float32(-libc.Float64FromFloat64(0.4935954809188843)),
		Fe:    int32(FE_INEXACT),
	},
	604: {
		Ffile: __ccgo_ts,
		Fline: int32(621),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9696688545572445),
		Fy:    -libc.Float64FromFloat64(0.03080065282492184),
		Fdy:   float32(0.49030396342277527),
		Fe:    int32(FE_INEXACT),
	},
	605: {
		Ffile: __ccgo_ts,
		Fline: int32(622),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9683454028881043),
		Fy:    -libc.Float64FromFloat64(0.03216643418569543),
		Fdy:   float32(0.4992527365684509),
		Fe:    int32(FE_INEXACT),
	},
	606: {
		Ffile: __ccgo_ts,
		Fline: int32(623),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.966363611448499),
		Fy:    -libc.Float64FromFloat64(0.03421510622687656),
		Fdy:   float32(0.49811500310897827),
		Fe:    int32(FE_INEXACT),
	},
	607: {
		Ffile: __ccgo_ts,
		Fline: int32(624),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9660337074217422),
		Fy:    -libc.Float64FromFloat64(0.03455655156706658),
		Fdy:   float32(0.4934997856616974),
		Fe:    int32(FE_INEXACT),
	},
	608: {
		Ffile: __ccgo_ts,
		Fline: int32(625),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9650446709374537),
		Fy:    -libc.Float64FromFloat64(0.03558088758765723),
		Fdy:   float32(0.49389392137527466),
		Fe:    int32(FE_INEXACT),
	},
	609: {
		Ffile: __ccgo_ts,
		Fline: int32(626),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9640566470401606),
		Fy:    -libc.Float64FromFloat64(0.03660522360824577),
		Fdy:   float32(-libc.Float64FromFloat64(0.4935038387775421)),
		Fe:    int32(FE_INEXACT),
	},
	610: {
		Ffile: __ccgo_ts,
		Fline: int32(627),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9633985264782978),
		Fy:    -libc.Float64FromFloat64(0.037288114288641774),
		Fdy:   float32(-libc.Float64FromFloat64(0.49069562554359436)),
		Fe:    int32(FE_INEXACT),
	},
	611: {
		Ffile: __ccgo_ts,
		Fline: int32(628),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9610986405085552),
		Fy:    -libc.Float64FromFloat64(0.03967823167002075),
		Fdy:   float32(0.4933150112628937),
		Fe:    int32(FE_INEXACT),
	},
	612: {
		Ffile: __ccgo_ts,
		Fline: int32(629),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9607705338747173),
		Fy:    -libc.Float64FromFloat64(0.04001967701021239),
		Fdy:   float32(-libc.Float64FromFloat64(0.492381751537323)),
		Fe:    int32(FE_INEXACT),
	},
	613: {
		Ffile: __ccgo_ts,
		Fline: int32(630),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9604425392522347),
		Fy:    -libc.Float64FromFloat64(0.04036112235040616),
		Fdy:   float32(-libc.Float64FromFloat64(0.4905974268913269)),
		Fe:    int32(FE_INEXACT),
	},
	614: {
		Ffile: __ccgo_ts,
		Fline: int32(631),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9594592270705958),
		Fy:    -libc.Float64FromFloat64(0.041385458370995054),
		Fdy:   float32(-libc.Float64FromFloat64(0.49960315227508545)),
		Fe:    int32(FE_INEXACT),
	},
	615: {
		Ffile: __ccgo_ts,
		Fline: int32(632),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9591316801112558),
		Fy:    -libc.Float64FromFloat64(0.04172690371121444),
		Fdy:   float32(-libc.Float64FromFloat64(0.49060484766960144)),
		Fe:    int32(FE_INEXACT),
	},
	616: {
		Ffile: __ccgo_ts,
		Fline: int32(633),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9568419802273874),
		Fy:    -libc.Float64FromFloat64(0.04411702109256391),
		Fdy:   float32(0.4914343059062958),
		Fe:    int32(FE_INEXACT),
	},
	617: {
		Ffile: __ccgo_ts,
		Fline: int32(634),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9565153267622388),
		Fy:    -libc.Float64FromFloat64(0.044458466432782257),
		Fdy:   float32(0.49053871631622314),
		Fe:    int32(FE_INEXACT),
	},
	618: {
		Ffile: __ccgo_ts,
		Fline: int32(635),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9548837314043811),
		Fy:    -libc.Float64FromFloat64(0.0461656931337446),
		Fdy:   float32(-libc.Float64FromFloat64(0.4945697784423828)),
		Fe:    int32(FE_INEXACT),
	},
	619: {
		Ffile: __ccgo_ts,
		Fline: int32(636),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9539061103945167),
		Fy:    -libc.Float64FromFloat64(0.04719002915433551),
		Fdy:   float32(0.4909532070159912),
		Fe:    int32(FE_INEXACT),
	},
	620: {
		Ffile: __ccgo_ts,
		Fline: int32(637),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9516288853215642),
		Fy:    -libc.Float64FromFloat64(0.049580146535707664),
		Fdy:   float32(0.4939286410808563),
		Fe:    int32(FE_INEXACT),
	},
	621: {
		Ffile: __ccgo_ts,
		Fline: int32(638),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9493570965845065),
		Fy:    -libc.Float64FromFloat64(0.051970263917085),
		Fdy:   float32(0.4952352046966553),
		Fe:    int32(FE_INEXACT),
	},
	622: {
		Ffile: __ccgo_ts,
		Fline: int32(639),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9467674066904423),
		Fy:    -libc.Float64FromFloat64(0.054701826638651696),
		Fdy:   float32(0.4933057725429535),
		Fe:    int32(FE_INEXACT),
	},
	623: {
		Ffile: __ccgo_ts,
		Fline: int32(640),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9461210887592751),
		Fy:    -libc.Float64FromFloat64(0.05538471731904609),
		Fdy:   float32(-libc.Float64FromFloat64(0.49211832880973816)),
		Fe:    int32(FE_INEXACT),
	},
	624: {
		Ffile: __ccgo_ts,
		Fline: int32(641),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9457980952676897),
		Fy:    -libc.Float64FromFloat64(0.05572616265925983),
		Fdy:   float32(0.49081093072891235),
		Fe:    int32(FE_INEXACT),
	},
	625: {
		Ffile: __ccgo_ts,
		Fline: int32(642),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9445072235830247),
		Fy:    -libc.Float64FromFloat64(0.05709194402003101),
		Fdy:   float32(0.4908629357814789),
		Fe:    int32(FE_INEXACT),
	},
	626: {
		Ffile: __ccgo_ts,
		Fline: int32(643),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9438624485829527),
		Fy:    -libc.Float64FromFloat64(0.05777483470042363),
		Fdy:   float32(0.4993323087692261),
		Fe:    int32(FE_INEXACT),
	},
	627: {
		Ffile: __ccgo_ts,
		Fline: int32(644),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9428961112899324),
		Fy:    -libc.Float64FromFloat64(0.058799170721011845),
		Fdy:   float32(0.49546197056770325),
		Fe:    int32(FE_INEXACT),
	},
	628: {
		Ffile: __ccgo_ts,
		Fline: int32(645),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9409664037328768),
		Fy:    -libc.Float64FromFloat64(0.06084784276219493),
		Fdy:   float32(0.49053680896759033),
		Fe:    int32(FE_INEXACT),
	},
	629: {
		Ffile: __ccgo_ts,
		Fline: int32(646),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9380792447917545),
		Fy:    -libc.Float64FromFloat64(0.06392085082395581),
		Fdy:   float32(0.4948404133319855),
		Fe:    int32(FE_INEXACT),
	},
	630: {
		Ffile: __ccgo_ts,
		Fline: int32(647),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.937438857900023),
		Fy:    -libc.Float64FromFloat64(0.0646037415043485),
		Fdy:   float32(0.49738574028015137),
		Fe:    int32(FE_INEXACT),
	},
	631: {
		Ffile: __ccgo_ts,
		Fline: int32(648),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9367989081732485),
		Fy:    -libc.Float64FromFloat64(0.06528663218473936),
		Fdy:   float32(-libc.Float64FromFloat64(0.49470648169517517)),
		Fe:    int32(FE_INEXACT),
	},
	632: {
		Ffile: __ccgo_ts,
		Fline: int32(649),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9364790971534491),
		Fy:    -libc.Float64FromFloat64(0.06562807752493668),
		Fdy:   float32(0.4993959367275238),
		Fe:    int32(FE_INEXACT),
	},
	633: {
		Ffile: __ccgo_ts,
		Fline: int32(650),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9361593953129137),
		Fy:    -libc.Float64FromFloat64(0.06596952286521739),
		Fdy:   float32(0.4902285933494568),
		Fe:    int32(FE_INEXACT),
	},
	634: {
		Ffile: __ccgo_ts,
		Fline: int32(651),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9352009444950148),
		Fy:    -libc.Float64FromFloat64(0.06699385888572211),
		Fdy:   float32(-libc.Float64FromFloat64(0.4980614483356476)),
		Fe:    int32(FE_INEXACT),
	},
	635: {
		Ffile: __ccgo_ts,
		Fline: int32(652),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.93074113239686),
		Fy:    -libc.Float64FromFloat64(0.07177409364847037),
		Fdy:   float32(-libc.Float64FromFloat64(0.49917978048324585)),
		Fe:    int32(FE_INEXACT),
	},
	636: {
		Ffile: __ccgo_ts,
		Fline: int32(653),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9272519196151529),
		Fy:    -libc.Float64FromFloat64(0.07552999239063493),
		Fdy:   float32(0.49672776460647583),
		Fe:    int32(FE_INEXACT),
	},
	637: {
		Ffile: __ccgo_ts,
		Fline: int32(654),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9253542290688181),
		Fy:    -libc.Float64FromFloat64(0.07757866443181219),
		Fdy:   float32(0.49406591057777405),
		Fe:    int32(FE_INEXACT),
	},
	638: {
		Ffile: __ccgo_ts,
		Fline: int32(655),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9250383251143067),
		Fy:    -libc.Float64FromFloat64(0.07792010977200924),
		Fdy:   float32(-libc.Float64FromFloat64(0.4963134527206421)),
		Fe:    int32(FE_INEXACT),
	},
	639: {
		Ffile: __ccgo_ts,
		Fline: int32(656),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9228300150464307),
		Fy:    -libc.Float64FromFloat64(0.08031022715338185),
		Fdy:   float32(0.4932170510292053),
		Fe:    int32(FE_INEXACT),
	},
	640: {
		Ffile: __ccgo_ts,
		Fline: int32(657),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9225149728259101),
		Fy:    -libc.Float64FromFloat64(0.08065167249357766),
		Fdy:   float32(0.49184224009513855),
		Fe:    int32(FE_INEXACT),
	},
	641: {
		Ffile: __ccgo_ts,
		Fline: int32(658),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9203126866643934),
		Fy:    -libc.Float64FromFloat64(0.08304178987496676),
		Fdy:   float32(0.492173969745636),
		Fe:    int32(FE_INEXACT),
	},
	642: {
		Ffile: __ccgo_ts,
		Fline: int32(659),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9193704598900464),
		Fy:    -libc.Float64FromFloat64(0.08406612589554589),
		Fdy:   float32(0.4998122453689575),
		Fe:    int32(FE_INEXACT),
	},
	643: {
		Ffile: __ccgo_ts,
		Fline: int32(660),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9178022251398503),
		Fy:    -libc.Float64FromFloat64(0.08577335259652598),
		Fdy:   float32(-libc.Float64FromFloat64(0.4974932372570038)),
		Fe:    int32(FE_INEXACT),
	},
	644: {
		Ffile: __ccgo_ts,
		Fline: int32(661),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.914049367249251),
		Fy:    -libc.Float64FromFloat64(0.08987069667888475),
		Fdy:   float32(0.49881115555763245),
		Fe:    int32(FE_INEXACT),
	},
	645: {
		Ffile: __ccgo_ts,
		Fline: int32(662),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9137373226282179),
		Fy:    -libc.Float64FromFloat64(0.09021214201907993),
		Fdy:   float32(0.49479401111602783),
		Fe:    int32(FE_INEXACT),
	},
	646: {
		Ffile: __ccgo_ts,
		Fline: int32(663),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9121786967157957),
		Fy:    -libc.Float64FromFloat64(0.0919193687200642),
		Fdy:   float32(0.4996547996997833),
		Fe:    int32(FE_INEXACT),
	},
	647: {
		Ffile: __ccgo_ts,
		Fline: int32(664),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9118672907174673),
		Fy:    -libc.Float64FromFloat64(0.09226081406025761),
		Fdy:   float32(0.49428629875183105),
		Fe:    int32(FE_INEXACT),
	},
	648: {
		Ffile: __ccgo_ts,
		Fline: int32(665),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9109337104371722),
		Fy:    -libc.Float64FromFloat64(0.09328515008084877),
		Fdy:   float32(-libc.Float64FromFloat64(0.4940306842327118)),
		Fe:    int32(FE_INEXACT),
	},
	649: {
		Ffile: __ccgo_ts,
		Fline: int32(666),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.908759071798632),
		Fy:    -libc.Float64FromFloat64(0.09567526746221898),
		Fdy:   float32(0.49990609288215637),
		Fe:    int32(FE_INEXACT),
	},
	650: {
		Ffile: __ccgo_ts,
		Fline: int32(667),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9068992282512988),
		Fy:    -libc.Float64FromFloat64(0.09772393950339679),
		Fdy:   float32(-libc.Float64FromFloat64(0.4970739781856537)),
		Fe:    int32(FE_INEXACT),
	},
	651: {
		Ffile: __ccgo_ts,
		Fline: int32(668),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9056614476498946),
		Fy:    -libc.Float64FromFloat64(0.09908972086483721),
		Fdy:   float32(-libc.Float64FromFloat64(0.4900471866130829)),
		Fe:    int32(FE_INEXACT),
	},
	652: {
		Ffile: __ccgo_ts,
		Fline: int32(669),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9025743821593739),
		Fy:    -libc.Float64FromFloat64(0.10250417426614657),
		Fdy:   float32(-libc.Float64FromFloat64(0.49502959847450256)),
		Fe:    int32(FE_INEXACT),
	},
	653: {
		Ffile: __ccgo_ts,
		Fline: int32(670),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9013425043208705),
		Fy:    -libc.Float64FromFloat64(0.10386995562694772),
		Fdy:   float32(0.49165818095207214),
		Fe:    int32(FE_INEXACT),
	},
	654: {
		Ffile: __ccgo_ts,
		Fline: int32(671),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8994978393347296),
		Fy:    -libc.Float64FromFloat64(0.10591862766810894),
		Fdy:   float32(0.4920984208583832),
		Fe:    int32(FE_INEXACT),
	},
	655: {
		Ffile: __ccgo_ts,
		Fline: int32(672),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8991907624169408),
		Fy:    -libc.Float64FromFloat64(0.10626007300831189),
		Fdy:   float32(0.4976048171520233),
		Fe:    int32(FE_INEXACT),
	},
	656: {
		Ffile: __ccgo_ts,
		Fline: int32(673),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8988837903311513),
		Fy:    -libc.Float64FromFloat64(0.10660151834860968),
		Fdy:   float32(-libc.Float64FromFloat64(0.4904788136482239)),
		Fe:    int32(FE_INEXACT),
	},
	657: {
		Ffile: __ccgo_ts,
		Fline: int32(674),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8985769230418454),
		Fy:    -libc.Float64FromFloat64(0.10694296368869874),
		Fdy:   float32(-libc.Float64FromFloat64(0.49140575528144836)),
		Fe:    int32(FE_INEXACT),
	},
	658: {
		Ffile: __ccgo_ts,
		Fline: int32(675),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8973505011318538),
		Fy:    -libc.Float64FromFloat64(0.10830874504948952),
		Fdy:   float32(0.4985845983028412),
		Fe:    int32(FE_INEXACT),
	},
	659: {
		Ffile: __ccgo_ts,
		Fline: int32(676),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8961257531027199),
		Fy:    -libc.Float64FromFloat64(0.10967452641026859),
		Fdy:   float32(0.49015653133392334),
		Fe:    int32(FE_INEXACT),
	},
	660: {
		Ffile: __ccgo_ts,
		Fline: int32(677),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8955140060794091),
		Fy:    -libc.Float64FromFloat64(0.11035741709066246),
		Fdy:   float32(-libc.Float64FromFloat64(0.49283912777900696)),
		Fe:    int32(FE_INEXACT),
	},
	661: {
		Ffile: __ccgo_ts,
		Fline: int32(678),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8952082891907264),
		Fy:    -libc.Float64FromFloat64(0.11069886243085889),
		Fdy:   float32(-libc.Float64FromFloat64(0.49308550357818604)),
		Fe:    int32(FE_INEXACT),
	},
	662: {
		Ffile: __ccgo_ts,
		Fline: int32(679),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8945971684810904),
		Fy:    -libc.Float64FromFloat64(0.11138175311125728),
		Fdy:   float32(-libc.Float64FromFloat64(0.4986288547515869)),
		Fe:    int32(FE_INEXACT),
	},
	663: {
		Ffile: __ccgo_ts,
		Fline: int32(680),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8942917645888986),
		Fy:    -libc.Float64FromFloat64(0.11172319845144932),
		Fdy:   float32(-libc.Float64FromFloat64(0.4901041090488434)),
		Fe:    int32(FE_INEXACT),
	},
	664: {
		Ffile: __ccgo_ts,
		Fline: int32(681),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8933761783355497),
		Fy:    -libc.Float64FromFloat64(0.11274753447205134),
		Fdy:   float32(0.4928297996520996),
		Fe:    int32(FE_INEXACT),
	},
	665: {
		Ffile: __ccgo_ts,
		Fline: int32(682),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8912434541494845),
		Fy:    -libc.Float64FromFloat64(0.11513765185341131),
		Fdy:   float32(0.4933846592903137),
		Fe:    int32(FE_INEXACT),
	},
	666: {
		Ffile: __ccgo_ts,
		Fline: int32(683),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8909391951719419),
		Fy:    -libc.Float64FromFloat64(0.11547909719361193),
		Fdy:   float32(0.4902154803276062),
		Fe:    int32(FE_INEXACT),
	},
	667: {
		Ffile: __ccgo_ts,
		Fline: int32(684),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8906350400644807),
		Fy:    -libc.Float64FromFloat64(0.11582054253380661),
		Fdy:   float32(-libc.Float64FromFloat64(0.4969140887260437)),
		Fe:    int32(FE_INEXACT),
	},
	668: {
		Ffile: __ccgo_ts,
		Fline: int32(685),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8903309887916321),
		Fy:    -libc.Float64FromFloat64(0.11616198787400535),
		Fdy:   float32(-libc.Float64FromFloat64(0.4924226999282837)),
		Fe:    int32(FE_INEXACT),
	},
	669: {
		Ffile: __ccgo_ts,
		Fline: int32(686),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8891158213376159),
		Fy:    -libc.Float64FromFloat64(0.11752776923481097),
		Fdy:   float32(-libc.Float64FromFloat64(0.49083951115608215)),
		Fe:    int32(FE_INEXACT),
	},
	670: {
		Ffile: __ccgo_ts,
		Fline: int32(687),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8857826557771602),
		Fy:    -libc.Float64FromFloat64(0.12128366797694587),
		Fdy:   float32(-libc.Float64FromFloat64(0.49141690135002136)),
		Fe:    int32(FE_INEXACT),
	},
	671: {
		Ffile: __ccgo_ts,
		Fline: int32(688),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8827633497479924),
		Fy:    -libc.Float64FromFloat64(0.12469812137890898),
		Fdy:   float32(-libc.Float64FromFloat64(0.49866199493408203)),
		Fe:    int32(FE_INEXACT),
	},
	672: {
		Ffile: __ccgo_ts,
		Fline: int32(689),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8824619857683135),
		Fy:    -libc.Float64FromFloat64(0.1250395667191082),
		Fdy:   float32(-libc.Float64FromFloat64(0.49587559700012207)),
		Fe:    int32(FE_INEXACT),
	},
	673: {
		Ffile: __ccgo_ts,
		Fline: int32(690),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8800547747255121),
		Fy:    -libc.Float64FromFloat64(0.12777112944069896),
		Fdy:   float32(-libc.Float64FromFloat64(0.49249520897865295)),
		Fe:    int32(FE_INEXACT),
	},
	674: {
		Ffile: __ccgo_ts,
		Fline: int32(691),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8797543354183023),
		Fy:    -libc.Float64FromFloat64(0.12811257478087468),
		Fdy:   float32(-libc.Float64FromFloat64(0.49718695878982544)),
		Fe:    int32(FE_INEXACT),
	},
	675: {
		Ffile: __ccgo_ts,
		Fline: int32(692),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8776541301582385),
		Fy:    -libc.Float64FromFloat64(0.13050269216225208),
		Fdy:   float32(-libc.Float64FromFloat64(0.4943873882293701)),
		Fe:    int32(FE_INEXACT),
	},
	676: {
		Ffile: __ccgo_ts,
		Fline: int32(693),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.876157053880922),
		Fy:    -libc.Float64FromFloat64(0.13220991886323166),
		Fdy:   float32(0.49791496992111206),
		Fe:    int32(FE_INEXACT),
	},
	677: {
		Ffile: __ccgo_ts,
		Fline: int32(694),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8758579452051334),
		Fy:    -libc.Float64FromFloat64(0.1325513642034264),
		Fdy:   float32(0.49067845940589905),
		Fe:    int32(FE_INEXACT),
	},
	678: {
		Ffile: __ccgo_ts,
		Fline: int32(695),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8746625312717331),
		Fy:    -libc.Float64FromFloat64(0.1339171455642125),
		Fdy:   float32(0.49099066853523254),
		Fe:    int32(FE_INEXACT),
	},
	679: {
		Ffile: __ccgo_ts,
		Fline: int32(696),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8743639328066132),
		Fy:    -libc.Float64FromFloat64(0.13425859090441014),
		Fdy:   float32(-libc.Float64FromFloat64(0.49670711159706116)),
		Fe:    int32(FE_INEXACT),
	},
	680: {
		Ffile: __ccgo_ts,
		Fline: int32(697),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8719788119193593),
		Fy:    -libc.Float64FromFloat64(0.13699015362598033),
		Fdy:   float32(-libc.Float64FromFloat64(0.49120867252349854)),
		Fe:    int32(FE_INEXACT),
	},
	681: {
		Ffile: __ccgo_ts,
		Fline: int32(698),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.871681129641284),
		Fy:    -libc.Float64FromFloat64(0.1373315989661778),
		Fdy:   float32(0.4927307665348053),
		Fe:    int32(FE_INEXACT),
	},
	682: {
		Ffile: __ccgo_ts,
		Fline: int32(699),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8713835489880849),
		Fy:    -libc.Float64FromFloat64(0.13767304430637506),
		Fdy:   float32(-libc.Float64FromFloat64(0.4984346330165863)),
		Fe:    int32(FE_INEXACT),
	},
	683: {
		Ffile: __ccgo_ts,
		Fline: int32(700),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8684133241975921),
		Fy:    -libc.Float64FromFloat64(0.14108749770833884),
		Fdy:   float32(0.49445483088493347),
		Fe:    int32(FE_INEXACT),
	},
	684: {
		Ffile: __ccgo_ts,
		Fline: int32(701),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8669320106118782),
		Fy:    -libc.Float64FromFloat64(0.14279472440931792),
		Fdy:   float32(0.4939049482345581),
		Fe:    int32(FE_INEXACT),
	},
	685: {
		Ffile: __ccgo_ts,
		Fline: int32(702),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.863976959472129),
		Fy:    -libc.Float64FromFloat64(0.146209177811288),
		Fdy:   float32(0.4909650981426239),
		Fe:    int32(FE_INEXACT),
	},
	686: {
		Ffile: __ccgo_ts,
		Fline: int32(703),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8633871590659558),
		Fy:    -libc.Float64FromFloat64(0.14689206849167688),
		Fdy:   float32(0.4996567368507385),
		Fe:    int32(FE_INEXACT),
	},
	687: {
		Ffile: __ccgo_ts,
		Fline: int32(704),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8625032133052423),
		Fy:    -libc.Float64FromFloat64(0.14791640451226676),
		Fdy:   float32(0.4957877993583679),
		Fe:    int32(FE_INEXACT),
	},
	688: {
		Ffile: __ccgo_ts,
		Fline: int32(705),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8610319810106725),
		Fy:    -libc.Float64FromFloat64(0.14962363121325023),
		Fdy:   float32(-libc.Float64FromFloat64(0.4974651038646698)),
		Fe:    int32(FE_INEXACT),
	},
	689: {
		Ffile: __ccgo_ts,
		Fline: int32(706),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8607380358389554),
		Fy:    -libc.Float64FromFloat64(0.14996507655344812),
		Fdy:   float32(0.4903302788734436),
		Fe:    int32(FE_INEXACT),
	},
	690: {
		Ffile: __ccgo_ts,
		Fline: int32(707),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8589764709406532),
		Fy:    -libc.Float64FromFloat64(0.1520137485946192),
		Fdy:   float32(-libc.Float64FromFloat64(0.4927206039428711)),
		Fe:    int32(FE_INEXACT),
	},
	691: {
		Ffile: __ccgo_ts,
		Fline: int32(708),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8575112544581619),
		Fy:    -libc.Float64FromFloat64(0.15372097529562032),
		Fdy:   float32(-libc.Float64FromFloat64(0.4933449327945709)),
		Fe:    int32(FE_INEXACT),
	},
	692: {
		Ffile: __ccgo_ts,
		Fline: int32(709),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8557562934097933),
		Fy:    -libc.Float64FromFloat64(0.15576964733678333),
		Fdy:   float32(0.4929549992084503),
		Fe:    int32(FE_INEXACT),
	},
	693: {
		Ffile: __ccgo_ts,
		Fline: int32(710),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8548801602176047),
		Fy:    -libc.Float64FromFloat64(0.15679398335737812),
		Fdy:   float32(0.4998406767845154),
		Fe:    int32(FE_INEXACT),
	},
	694: {
		Ffile: __ccgo_ts,
		Fline: int32(711),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8513845882230047),
		Fy:    -libc.Float64FromFloat64(0.16089132743972542),
		Fdy:   float32(0.4952109158039093),
		Fe:    int32(FE_INEXACT),
	},
	695: {
		Ffile: __ccgo_ts,
		Fline: int32(712),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8490621512053441),
		Fy:    -libc.Float64FromFloat64(0.16362289016129647),
		Fdy:   float32(0.49513253569602966),
		Fe:    int32(FE_INEXACT),
	},
	696: {
		Ffile: __ccgo_ts,
		Fline: int32(713),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8487722923785611),
		Fy:    -libc.Float64FromFloat64(0.1639643355014997),
		Fdy:   float32(-libc.Float64FromFloat64(0.4924023747444153)),
		Fe:    int32(FE_INEXACT),
	},
	697: {
		Ffile: __ccgo_ts,
		Fline: int32(714),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8450131201641286),
		Fy:    -libc.Float64FromFloat64(0.16840312492404913),
		Fdy:   float32(0.49467232823371887),
		Fe:    int32(FE_INEXACT),
	},
	698: {
		Ffile: __ccgo_ts,
		Fline: int32(715),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8404092945763008),
		Fy:    -libc.Float64FromFloat64(0.17386625036718947),
		Fdy:   float32(0.4980604350566864),
		Fe:    int32(FE_INEXACT),
	},
	699: {
		Ffile: __ccgo_ts,
		Fline: int32(716),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8372587226820964),
		Fy:    -libc.Float64FromFloat64(0.1776221491093648),
		Fdy:   float32(-libc.Float64FromFloat64(0.4924791753292084)),
		Fe:    int32(FE_INEXACT),
	},
	700: {
		Ffile: __ccgo_ts,
		Fline: int32(717),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8318446198639039),
		Fy:    -libc.Float64FromFloat64(0.18410961057307895),
		Fdy:   float32(-libc.Float64FromFloat64(0.499199241399765)),
		Fe:    int32(FE_INEXACT),
	},
	701: {
		Ffile: __ccgo_ts,
		Fline: int32(718),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.82816041950793),
		Fy:    -libc.Float64FromFloat64(0.18854839999563455),
		Fdy:   float32(-libc.Float64FromFloat64(0.49829092621803284)),
		Fe:    int32(FE_INEXACT),
	},
	702: {
		Ffile: __ccgo_ts,
		Fline: int32(719),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8270301054998307),
		Fy:    -libc.Float64FromFloat64(0.18991418135642063),
		Fdy:   float32(0.49566569924354553),
		Fe:    int32(FE_INEXACT),
	},
	703: {
		Ffile: __ccgo_ts,
		Fline: int32(720),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8267477681282397),
		Fy:    -libc.Float64FromFloat64(0.19025562669661575),
		Fdy:   float32(-libc.Float64FromFloat64(0.49103525280952454)),
		Fe:    int32(FE_INEXACT),
	},
	704: {
		Ffile: __ccgo_ts,
		Fline: int32(721),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8259013341997944),
		Fy:    -libc.Float64FromFloat64(0.1912799627172072),
		Fdy:   float32(-libc.Float64FromFloat64(0.4960572421550751)),
		Fe:    int32(FE_INEXACT),
	},
	705: {
		Ffile: __ccgo_ts,
		Fline: int32(722),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8256193821761142),
		Fy:    -libc.Float64FromFloat64(0.19162140805740033),
		Fdy:   float32(-libc.Float64FromFloat64(0.4977794885635376)),
		Fe:    int32(FE_INEXACT),
	},
	706: {
		Ffile: __ccgo_ts,
		Fline: int32(723),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8247741035022609),
		Fy:    -libc.Float64FromFloat64(0.1926457440779899),
		Fdy:   float32(-libc.Float64FromFloat64(0.4930289089679718)),
		Fe:    int32(FE_INEXACT),
	},
	707: {
		Ffile: __ccgo_ts,
		Fline: int32(724),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8244925363005422),
		Fy:    -libc.Float64FromFloat64(0.19298718941819018),
		Fdy:   float32(-libc.Float64FromFloat64(0.49325960874557495)),
		Fe:    int32(FE_INEXACT),
	},
	708: {
		Ffile: __ccgo_ts,
		Fline: int32(725),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.822805150533457),
		Fy:    -libc.Float64FromFloat64(0.19503586145936735),
		Fdy:   float32(-libc.Float64FromFloat64(0.4941540062427521)),
		Fe:    int32(FE_INEXACT),
	},
	709: {
		Ffile: __ccgo_ts,
		Fline: int32(726),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8222434563740615),
		Fy:    -libc.Float64FromFloat64(0.19571875213975864),
		Fdy:   float32(-libc.Float64FromFloat64(0.4941878914833069)),
		Fe:    int32(FE_INEXACT),
	},
	710: {
		Ffile: __ccgo_ts,
		Fline: int32(727),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8214016340120976),
		Fy:    -libc.Float64FromFloat64(0.19674308816034555),
		Fdy:   float32(-libc.Float64FromFloat64(0.49610137939453125)),
		Fe:    int32(FE_INEXACT),
	},
	711: {
		Ffile: __ccgo_ts,
		Fline: int32(728),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8197205740081908),
		Fy:    -libc.Float64FromFloat64(0.1987917602015258),
		Fdy:   float32(-libc.Float64FromFloat64(0.49868014454841614)),
		Fe:    int32(FE_INEXACT),
	},
	712: {
		Ffile: __ccgo_ts,
		Fline: int32(729),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8191609855582914),
		Fy:    -libc.Float64FromFloat64(0.19947465088191624),
		Fdy:   float32(0.49146661162376404),
		Fe:    int32(FE_INEXACT),
	},
	713: {
		Ffile: __ccgo_ts,
		Fline: int32(730),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8186017791156784),
		Fy:    -libc.Float64FromFloat64(0.20015754156230853),
		Fdy:   float32(-libc.Float64FromFloat64(0.4922550916671753)),
		Fe:    int32(FE_INEXACT),
	},
	714: {
		Ffile: __ccgo_ts,
		Fline: int32(731),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8177636851450142),
		Fy:    -libc.Float64FromFloat64(0.2011818775828988),
		Fdy:   float32(0.4970836341381073),
		Fe:    int32(FE_INEXACT),
	},
	715: {
		Ffile: __ccgo_ts,
		Fline: int32(732),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8135860684900397),
		Fy:    -libc.Float64FromFloat64(0.2063035576858871),
		Fdy:   float32(-libc.Float64FromFloat64(0.4903571903705597)),
		Fe:    int32(FE_INEXACT),
	},
	716: {
		Ffile: __ccgo_ts,
		Fline: int32(733),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8110897495113072),
		Fy:    -libc.Float64FromFloat64(0.2093765657476146),
		Fdy:   float32(-libc.Float64FromFloat64(0.4966326951980591)),
		Fe:    int32(FE_INEXACT),
	},
	717: {
		Ffile: __ccgo_ts,
		Fline: int32(734),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8097062167721057),
		Fy:    -libc.Float64FromFloat64(0.21108379244859615),
		Fdy:   float32(-libc.Float64FromFloat64(0.4926416277885437)),
		Fe:    int32(FE_INEXACT),
	},
	718: {
		Ffile: __ccgo_ts,
		Fline: int32(735),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8091534646989178),
		Fy:    -libc.Float64FromFloat64(0.2117666831289873),
		Fdy:   float32(0.4964464604854584),
		Fe:    int32(FE_INEXACT),
	},
	719: {
		Ffile: __ccgo_ts,
		Fline: int32(736),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8066707462393806),
		Fy:    -libc.Float64FromFloat64(0.21483969119075663),
		Fdy:   float32(0.49042031168937683),
		Fe:    int32(FE_INEXACT),
	},
	720: {
		Ffile: __ccgo_ts,
		Fline: int32(737),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8044702812226966),
		Fy:    -libc.Float64FromFloat64(0.21757125391232546),
		Fdy:   float32(-libc.Float64FromFloat64(0.4927595257759094)),
		Fe:    int32(FE_INEXACT),
	},
	721: {
		Ffile: __ccgo_ts,
		Fline: int32(738),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8041956454830617),
		Fy:    -libc.Float64FromFloat64(0.21791269925252174),
		Fdy:   float32(0.49411964416503906),
		Fe:    int32(FE_INEXACT),
	},
	722: {
		Ffile: __ccgo_ts,
		Fline: int32(739),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8025497988241774),
		Fy:    -libc.Float64FromFloat64(0.21996137129370685),
		Fdy:   float32(0.4990193247795105),
		Fe:    int32(FE_INEXACT),
	},
	723: {
		Ffile: __ccgo_ts,
		Fline: int32(740),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8017281390566221),
		Fy:    -libc.Float64FromFloat64(0.2209857073142914),
		Fdy:   float32(-libc.Float64FromFloat64(0.49329790472984314)),
		Fe:    int32(FE_INEXACT),
	},
	724: {
		Ffile: __ccgo_ts,
		Fline: int32(741),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8011808332785263),
		Fy:    -libc.Float64FromFloat64(0.2216685979946869),
		Fdy:   float32(-libc.Float64FromFloat64(0.4959532916545868)),
		Fe:    int32(FE_INEXACT),
	},
	725: {
		Ffile: __ccgo_ts,
		Fline: int32(742),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.8000873423343984),
		Fy:    -libc.Float64FromFloat64(0.22303437935568693),
		Fdy:   float32(-libc.Float64FromFloat64(0.49052006006240845)),
		Fe:    int32(FE_INEXACT),
	},
	726: {
		Ffile: __ccgo_ts,
		Fline: int32(743),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7998142028734677),
		Fy:    -libc.Float64FromFloat64(0.22337582469574854),
		Fdy:   float32(0.49039843678474426),
		Fe:    int32(FE_INEXACT),
	},
	727: {
		Ffile: __ccgo_ts,
		Fline: int32(744),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.794642233669141),
		Fy:    -libc.Float64FromFloat64(0.22986328615939966),
		Fdy:   float32(0.49015167355537415),
		Fe:    int32(FE_INEXACT),
	},
	728: {
		Ffile: __ccgo_ts,
		Fline: int32(745),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7940997651379443),
		Fy:    -libc.Float64FromFloat64(0.23054617683980616),
		Fdy:   float32(0.4912400543689728),
		Fe:    int32(FE_INEXACT),
	},
	729: {
		Ffile: __ccgo_ts,
		Fline: int32(746),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7919335916885177),
		Fy:    -libc.Float64FromFloat64(0.23327773956135905),
		Fdy:   float32(0.4908215403556824),
		Fe:    int32(FE_INEXACT),
	},
	730: {
		Ffile: __ccgo_ts,
		Fline: int32(747),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7908527218339777),
		Fy:    -libc.Float64FromFloat64(0.23464352092214436),
		Fdy:   float32(-libc.Float64FromFloat64(0.49861642718315125)),
		Fe:    int32(FE_INEXACT),
	},
	731: {
		Ffile: __ccgo_ts,
		Fline: int32(748),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7862754586069147),
		Fy:    -libc.Float64FromFloat64(0.24044809170548168),
		Fdy:   float32(0.4962637722492218),
		Fe:    int32(FE_INEXACT),
	},
	732: {
		Ffile: __ccgo_ts,
		Fline: int32(749),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7849342533461222),
		Fy:    -libc.Float64FromFloat64(0.24215531840647292),
		Fdy:   float32(-libc.Float64FromFloat64(0.4968024492263794)),
		Fe:    int32(FE_INEXACT),
	},
	733: {
		Ffile: __ccgo_ts,
		Fline: int32(750),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7838629365322646),
		Fy:    -libc.Float64FromFloat64(0.24352109976725314),
		Fdy:   float32(-libc.Float64FromFloat64(0.4939420819282532)),
		Fe:    int32(FE_INEXACT),
	},
	734: {
		Ffile: __ccgo_ts,
		Fline: int32(751),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7830604085904834),
		Fy:    -libc.Float64FromFloat64(0.24454543578783908),
		Fdy:   float32(0.4949719309806824),
		Fe:    int32(FE_INEXACT),
	},
	735: {
		Ffile: __ccgo_ts,
		Fline: int32(752),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7817246874662636),
		Fy:    -libc.Float64FromFloat64(0.2462526624888255),
		Fdy:   float32(-libc.Float64FromFloat64(0.4981144666671753)),
		Fe:    int32(FE_INEXACT),
	},
	736: {
		Ffile: __ccgo_ts,
		Fline: int32(753),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.781191037195693),
		Fy:    -libc.Float64FromFloat64(0.24693555316922278),
		Fdy:   float32(0.49488934874534607),
		Fe:    int32(FE_INEXACT),
	},
	737: {
		Ffile: __ccgo_ts,
		Fline: int32(754),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7798585047895361),
		Fy:    -libc.Float64FromFloat64(0.24864277987019573),
		Fdy:   float32(0.4946981966495514),
		Fe:    int32(FE_INEXACT),
	},
	738: {
		Ffile: __ccgo_ts,
		Fline: int32(755),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7793261284826896),
		Fy:    -libc.Float64FromFloat64(0.24932567055058716),
		Fdy:   float32(0.49108976125717163),
		Fe:    int32(FE_INEXACT),
	},
	739: {
		Ffile: __ccgo_ts,
		Fline: int32(756),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7790600766312865),
		Fy:    -libc.Float64FromFloat64(0.24966711589079202),
		Fdy:   float32(0.49895092844963074),
		Fe:    int32(FE_INEXACT),
	},
	740: {
		Ffile: __ccgo_ts,
		Fline: int32(757),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7787941156065508),
		Fy:    -libc.Float64FromFloat64(0.25000856123098475),
		Fdy:   float32(0.49177148938179016),
		Fe:    int32(FE_INEXACT),
	},
	741: {
		Ffile: __ccgo_ts,
		Fline: int32(758),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7764045484708353),
		Fy:    -libc.Float64FromFloat64(0.25308156929275333),
		Fdy:   float32(0.49430570006370544),
		Fe:    int32(FE_INEXACT),
	},
	742: {
		Ffile: __ccgo_ts,
		Fline: int32(759),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.774815578374795),
		Fy:    -libc.Float64FromFloat64(0.2551302413339307),
		Fdy:   float32(0.4980194568634033),
		Fe:    int32(FE_INEXACT),
	},
	743: {
		Ffile: __ccgo_ts,
		Fline: int32(760),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7742866446598106),
		Fy:    -libc.Float64FromFloat64(0.25581313201431916),
		Fdy:   float32(-libc.Float64FromFloat64(0.4929640591144562)),
		Fe:    int32(FE_INEXACT),
	},
	744: {
		Ffile: __ccgo_ts,
		Fline: int32(761),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7737580720254205),
		Fy:    -libc.Float64FromFloat64(0.25649602269471644),
		Fdy:   float32(-libc.Float64FromFloat64(0.4944615960121155)),
		Fe:    int32(FE_INEXACT),
	},
	745: {
		Ffile: __ccgo_ts,
		Fline: int32(762),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7727020090126306),
		Fy:    -libc.Float64FromFloat64(0.25786180405551595),
		Fdy:   float32(0.4905416667461395),
		Fe:    int32(FE_INEXACT),
	},
	746: {
		Ffile: __ccgo_ts,
		Fline: int32(763),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7716473873665202),
		Fy:    -libc.Float64FromFloat64(0.2592275854162889),
		Fdy:   float32(0.4922045171260834),
		Fe:    int32(FE_INEXACT),
	},
	747: {
		Ffile: __ccgo_ts,
		Fline: int32(764),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7700681531572044),
		Fy:    -libc.Float64FromFloat64(0.261276257457467),
		Fdy:   float32(-libc.Float64FromFloat64(0.4962473213672638)),
		Fe:    int32(FE_INEXACT),
	},
	748: {
		Ffile: __ccgo_ts,
		Fline: int32(765),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7671812799448212),
		Fy:    -libc.Float64FromFloat64(0.26503215619962883),
		Fdy:   float32(-libc.Float64FromFloat64(0.4953794479370117)),
		Fe:    int32(FE_INEXACT),
	},
	749: {
		Ffile: __ccgo_ts,
		Fline: int32(766),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7640443052883066),
		Fy:    -libc.Float64FromFloat64(0.2691295002819811),
		Fdy:   float32(-libc.Float64FromFloat64(0.49803632497787476)),
		Fe:    int32(FE_INEXACT),
	},
	750: {
		Ffile: __ccgo_ts,
		Fline: int32(767),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.763262067890441),
		Fy:    -libc.Float64FromFloat64(0.2701538363025731),
		Fdy:   float32(-libc.Float64FromFloat64(0.493619829416275)),
		Fe:    int32(FE_INEXACT),
	},
	751: {
		Ffile: __ccgo_ts,
		Fline: int32(768),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7630015001013756),
		Fy:    -libc.Float64FromFloat64(0.27049528164276404),
		Fdy:   float32(-libc.Float64FromFloat64(0.4909423291683197)),
		Fe:    int32(FE_INEXACT),
	},
	752: {
		Ffile: __ccgo_ts,
		Fline: int32(769),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7627410212667733),
		Fy:    -libc.Float64FromFloat64(0.27083672698296374),
		Fdy:   float32(0.49688634276390076),
		Fe:    int32(FE_INEXACT),
	},
	753: {
		Ffile: __ccgo_ts,
		Fline: int32(770),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7598816164033821),
		Fy:    -libc.Float64FromFloat64(0.2745926257251247),
		Fdy:   float32(0.49013441801071167),
		Fe:    int32(FE_INEXACT),
	},
	754: {
		Ffile: __ccgo_ts,
		Fline: int32(771),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7588444926586534),
		Fy:    -libc.Float64FromFloat64(0.27595840708590824),
		Fdy:   float32(-libc.Float64FromFloat64(0.49473750591278076)),
		Fe:    int32(FE_INEXACT),
	},
	755: {
		Ffile: __ccgo_ts,
		Fline: int32(772),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7578087844313413),
		Fy:    -libc.Float64FromFloat64(0.2773241884466905),
		Fdy:   float32(-libc.Float64FromFloat64(0.4944440722465515)),
		Fe:    int32(FE_INEXACT),
	},
	756: {
		Ffile: __ccgo_ts,
		Fline: int32(773),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7575500783226484),
		Fy:    -libc.Float64FromFloat64(0.2776656337868868),
		Fdy:   float32(-libc.Float64FromFloat64(0.4913228750228882)),
		Fe:    int32(FE_INEXACT),
	},
	757: {
		Ffile: __ccgo_ts,
		Fline: int32(774),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7565161367755828),
		Fy:    -libc.Float64FromFloat64(0.27903141514767243),
		Fdy:   float32(0.4986989200115204),
		Fe:    int32(FE_INEXACT),
	},
	758: {
		Ffile: __ccgo_ts,
		Fline: int32(775),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7526514085293461),
		Fy:    -libc.Float64FromFloat64(0.28415309525062166),
		Fdy:   float32(-libc.Float64FromFloat64(0.4953634738922119)),
		Fe:    int32(FE_INEXACT),
	},
	759: {
		Ffile: __ccgo_ts,
		Fline: int32(776),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7511110510143703),
		Fy:    -libc.Float64FromFloat64(0.2862017672918037),
		Fdy:   float32(-libc.Float64FromFloat64(0.4909941256046295)),
		Fe:    int32(FE_INEXACT),
	},
	760: {
		Ffile: __ccgo_ts,
		Fline: int32(777),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7500858977679541),
		Fy:    -libc.Float64FromFloat64(0.2875675486526095),
		Fdy:   float32(-libc.Float64FromFloat64(0.4908120334148407)),
		Fe:    int32(FE_INEXACT),
	},
	761: {
		Ffile: __ccgo_ts,
		Fline: int32(778),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7439642943116302),
		Fy:    -libc.Float64FromFloat64(0.2957622368173009),
		Fdy:   float32(-libc.Float64FromFloat64(0.4992794394493103)),
		Fe:    int32(FE_INEXACT),
	},
	762: {
		Ffile: __ccgo_ts,
		Fline: int32(779),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7432026150612813),
		Fy:    -libc.Float64FromFloat64(0.29678657283788973),
		Fdy:   float32(-libc.Float64FromFloat64(0.4926382303237915)),
		Fe:    int32(FE_INEXACT),
	},
	763: {
		Ffile: __ccgo_ts,
		Fline: int32(780),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7424417156269629),
		Fy:    -libc.Float64FromFloat64(0.297810908858474),
		Fdy:   float32(-libc.Float64FromFloat64(0.49433445930480957)),
		Fe:    int32(FE_INEXACT),
	},
	764: {
		Ffile: __ccgo_ts,
		Fline: int32(781),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7411752806595144),
		Fy:    -libc.Float64FromFloat64(0.29951813555945783),
		Fdy:   float32(-libc.Float64FromFloat64(0.4978193938732147)),
		Fe:    int32(FE_INEXACT),
	},
	765: {
		Ffile: __ccgo_ts,
		Fline: int32(782),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7406693117482015),
		Fy:    -libc.Float64FromFloat64(0.30020102623985445),
		Fdy:   float32(0.49148818850517273),
		Fe:    int32(FE_INEXACT),
	},
	766: {
		Ffile: __ccgo_ts,
		Fline: int32(783),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7389011390642257),
		Fy:    -libc.Float64FromFloat64(0.3025911436212234),
		Fdy:   float32(-libc.Float64FromFloat64(0.4912014901638031)),
		Fe:    int32(FE_INEXACT),
	},
	767: {
		Ffile: __ccgo_ts,
		Fline: int32(784),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.737892650505224),
		Fy:    -libc.Float64FromFloat64(0.303956924982021),
		Fdy:   float32(-libc.Float64FromFloat64(0.4923177659511566)),
		Fe:    int32(FE_INEXACT),
	},
	768: {
		Ffile: __ccgo_ts,
		Fline: int32(785),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7373889225062212),
		Fy:    -libc.Float64FromFloat64(0.3046398156624085),
		Fdy:   float32(-libc.Float64FromFloat64(0.4992019236087799)),
		Fe:    int32(FE_INEXACT),
	},
	769: {
		Ffile: __ccgo_ts,
		Fline: int32(786),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7366339751975747),
		Fy:    -libc.Float64FromFloat64(0.30566415168301464),
		Fdy:   float32(0.49381038546562195),
		Fe:    int32(FE_INEXACT),
	},
	770: {
		Ffile: __ccgo_ts,
		Fline: int32(787),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7363824978946492),
		Fy:    -libc.Float64FromFloat64(0.306005597023189),
		Fdy:   float32(0.49077844619750977),
		Fe:    int32(FE_INEXACT),
	},
	771: {
		Ffile: __ccgo_ts,
		Fline: int32(788),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7353774469007994),
		Fy:    -libc.Float64FromFloat64(0.30737137838398176),
		Fdy:   float32(-libc.Float64FromFloat64(0.4905019700527191)),
		Fe:    int32(FE_INEXACT),
	},
	772: {
		Ffile: __ccgo_ts,
		Fline: int32(789),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7348754359245617),
		Fy:    -libc.Float64FromFloat64(0.3080542690643686),
		Fdy:   float32(-libc.Float64FromFloat64(0.49550676345825195)),
		Fe:    int32(FE_INEXACT),
	},
	773: {
		Ffile: __ccgo_ts,
		Fline: int32(790),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7333714582697353),
		Fy:    -libc.Float64FromFloat64(0.31010294110555836),
		Fdy:   float32(-libc.Float64FromFloat64(0.49081799387931824)),
		Fe:    int32(FE_INEXACT),
	},
	774: {
		Ffile: __ccgo_ts,
		Fline: int32(791),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7331210947478422),
		Fy:    -libc.Float64FromFloat64(0.3104443864457482),
		Fdy:   float32(0.4950374662876129),
		Fe:    int32(FE_INEXACT),
	},
	775: {
		Ffile: __ccgo_ts,
		Fline: int32(792),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7311212610740591),
		Fy:    -libc.Float64FromFloat64(0.31317594916731617),
		Fdy:   float32(0.497704416513443),
		Fe:    int32(FE_INEXACT),
	},
	776: {
		Ffile: __ccgo_ts,
		Fline: int32(793),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7308716657403485),
		Fy:    -libc.Float64FromFloat64(0.3135173945075196),
		Fdy:   float32(-libc.Float64FromFloat64(0.4912841022014618)),
		Fe:    int32(FE_INEXACT),
	},
	777: {
		Ffile: __ccgo_ts,
		Fline: int32(794),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7303727306697032),
		Fy:    -libc.Float64FromFloat64(0.31420028518790194),
		Fdy:   float32(-libc.Float64FromFloat64(0.49085164070129395)),
		Fe:    int32(FE_INEXACT),
	},
	778: {
		Ffile: __ccgo_ts,
		Fline: int32(795),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7301233908745768),
		Fy:    -libc.Float64FromFloat64(0.314541730528113),
		Fdy:   float32(0.49043720960617065),
		Fe:    int32(FE_INEXACT),
	},
	779: {
		Ffile: __ccgo_ts,
		Fline: int32(796),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7283803940807347),
		Fy:    -libc.Float64FromFloat64(0.3169318479094784),
		Fdy:   float32(0.4946947991847992),
		Fe:    int32(FE_INEXACT),
	},
	780: {
		Ffile: __ccgo_ts,
		Fline: int32(797),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7268897090158154),
		Fy:    -libc.Float64FromFloat64(0.3189805199506639),
		Fdy:   float32(0.490811288356781),
		Fe:    int32(FE_INEXACT),
	},
	781: {
		Ffile: __ccgo_ts,
		Fline: int32(798),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7263934922580472),
		Fy:    -libc.Float64FromFloat64(0.3196634106310448),
		Fdy:   float32(0.4997774362564087),
		Fe:    int32(FE_INEXACT),
	},
	782: {
		Ffile: __ccgo_ts,
		Fline: int32(799),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7234232972917924),
		Fy:    -libc.Float64FromFloat64(0.32376075471344457),
		Fdy:   float32(-libc.Float64FromFloat64(0.49057692289352417)),
		Fe:    int32(FE_INEXACT),
	},
	783: {
		Ffile: __ccgo_ts,
		Fline: int32(800),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7226826481518305),
		Fy:    -libc.Float64FromFloat64(0.3247850907339927),
		Fdy:   float32(-libc.Float64FromFloat64(0.4950527548789978)),
		Fe:    int32(FE_INEXACT),
	},
	784: {
		Ffile: __ccgo_ts,
		Fline: int32(801),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7221893033759262),
		Fy:    -libc.Float64FromFloat64(0.3254679814144502),
		Fdy:   float32(0.49099284410476685),
		Fe:    int32(FE_INEXACT),
	},
	785: {
		Ffile: __ccgo_ts,
		Fline: int32(802),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7204652473382305),
		Fy:    -libc.Float64FromFloat64(0.3278580987957566),
		Fdy:   float32(0.49922969937324524),
		Fe:    int32(FE_INEXACT),
	},
	786: {
		Ffile: __ccgo_ts,
		Fline: int32(803),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7189907612143241),
		Fy:    -libc.Float64FromFloat64(0.32990677083693626),
		Fdy:   float32(0.4908960163593292),
		Fe:    int32(FE_INEXACT),
	},
	787: {
		Ffile: __ccgo_ts,
		Fline: int32(804),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7172743409389752),
		Fy:    -libc.Float64FromFloat64(0.33229688821831127),
		Fdy:   float32(-libc.Float64FromFloat64(0.49215611815452576)),
		Fe:    int32(FE_INEXACT),
	},
	788: {
		Ffile: __ccgo_ts,
		Fline: int32(805),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7148294175392654),
		Fy:    -libc.Float64FromFloat64(0.335711341620278),
		Fdy:   float32(-libc.Float64FromFloat64(0.4905996322631836)),
		Fe:    int32(FE_INEXACT),
	},
	789: {
		Ffile: __ccgo_ts,
		Fline: int32(806),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.709964543856365),
		Fy:    -libc.Float64FromFloat64(0.3425402484242027),
		Fdy:   float32(0.49830830097198486),
		Fe:    int32(FE_INEXACT),
	},
	790: {
		Ffile: __ccgo_ts,
		Fline: int32(807),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7097221711520301),
		Fy:    -libc.Float64FromFloat64(0.34288169376439986),
		Fdy:   float32(-libc.Float64FromFloat64(0.4919234812259674)),
		Fe:    int32(FE_INEXACT),
	},
	791: {
		Ffile: __ccgo_ts,
		Fline: int32(808),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.707302990310754),
		Fy:    -libc.Float64FromFloat64(0.3462961471663587),
		Fdy:   float32(-libc.Float64FromFloat64(0.4982340633869171)),
		Fe:    int32(FE_INEXACT),
	},
	792: {
		Ffile: __ccgo_ts,
		Fline: int32(809),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7068201445746446),
		Fy:    -libc.Float64FromFloat64(0.3469790378467559),
		Fdy:   float32(-libc.Float64FromFloat64(0.4963652789592743)),
		Fe:    int32(FE_INEXACT),
	},
	793: {
		Ffile: __ccgo_ts,
		Fline: int32(810),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7046514145412596),
		Fy:    -libc.Float64FromFloat64(0.35005204590852845),
		Fdy:   float32(0.4981462061405182),
		Fe:    int32(FE_INEXACT),
	},
	794: {
		Ffile: __ccgo_ts,
		Fline: int32(811),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7036896716881657),
		Fy:    -libc.Float64FromFloat64(0.3514178272693044),
		Fdy:   float32(0.49618637561798096),
		Fe:    int32(FE_INEXACT),
	},
	795: {
		Ffile: __ccgo_ts,
		Fline: int32(812),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7027292414689272),
		Fy:    -libc.Float64FromFloat64(0.35278360863010544),
		Fdy:   float32(-libc.Float64FromFloat64(0.49039512872695923)),
		Fe:    int32(FE_INEXACT),
	},
	796: {
		Ffile: __ccgo_ts,
		Fline: int32(813),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7020097791426726),
		Fy:    -libc.Float64FromFloat64(0.3538079446506856),
		Fdy:   float32(-libc.Float64FromFloat64(0.49549609422683716)),
		Fe:    int32(FE_INEXACT),
	},
	797: {
		Ffile: __ccgo_ts,
		Fline: int32(814),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.701530546857208),
		Fy:    -libc.Float64FromFloat64(0.3544908353310755),
		Fdy:   float32(0.49012354016304016),
		Fe:    int32(FE_INEXACT),
	},
	798: {
		Ffile: __ccgo_ts,
		Fline: int32(815),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7010516417232864),
		Fy:    -libc.Float64FromFloat64(0.35517372601146735),
		Fdy:   float32(0.49028611183166504),
		Fe:    int32(FE_INEXACT),
	},
	799: {
		Ffile: __ccgo_ts,
		Fline: int32(816),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6986620155155403),
		Fy:    -libc.Float64FromFloat64(0.35858817941342874),
		Fdy:   float32(0.4946027398109436),
		Fe:    int32(FE_INEXACT),
	},
	800: {
		Ffile: __ccgo_ts,
		Fline: int32(817),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6984235013480071),
		Fy:    -libc.Float64FromFloat64(0.3589296247537053),
		Fdy:   float32(-libc.Float64FromFloat64(0.49029362201690674)),
		Fe:    int32(FE_INEXACT),
	},
	801: {
		Ffile: __ccgo_ts,
		Fline: int32(818),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6977084472883437),
		Fy:    -libc.Float64FromFloat64(0.3599539607742157),
		Fdy:   float32(0.4905036389827728),
		Fe:    int32(FE_INEXACT),
	},
	802: {
		Ffile: __ccgo_ts,
		Fline: int32(819),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6972321513395939),
		Fy:    -libc.Float64FromFloat64(0.3606368514546187),
		Fdy:   float32(0.49018704891204834),
		Fe:    int32(FE_INEXACT),
	},
	803: {
		Ffile: __ccgo_ts,
		Fline: int32(820),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6927235252085766),
		Fy:    -libc.Float64FromFloat64(0.3671243129183407),
		Fdy:   float32(0.49816617369651794),
		Fe:    int32(FE_INEXACT),
	},
	804: {
		Ffile: __ccgo_ts,
		Fline: int32(821),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6924870383648073),
		Fy:    -libc.Float64FromFloat64(0.36746575825854494),
		Fdy:   float32(-libc.Float64FromFloat64(0.49114713072776794)),
		Fe:    int32(FE_INEXACT),
	},
	805: {
		Ffile: __ccgo_ts,
		Fline: int32(822),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6917780621245666),
		Fy:    -libc.Float64FromFloat64(0.36849009427912655),
		Fdy:   float32(-libc.Float64FromFloat64(0.49624186754226685)),
		Fe:    int32(FE_INEXACT),
	},
	806: {
		Ffile: __ccgo_ts,
		Fline: int32(823),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6913058145981126),
		Fy:    -libc.Float64FromFloat64(0.36917298495951734),
		Fdy:   float32(-libc.Float64FromFloat64(0.49652406573295593)),
		Fe:    int32(FE_INEXACT),
	},
	807: {
		Ffile: __ccgo_ts,
		Fline: int32(824),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.690362286475159),
		Fy:    -libc.Float64FromFloat64(0.370538766320307),
		Fdy:   float32(0.49702391028404236),
		Fe:    int32(FE_INEXACT),
	},
	808: {
		Ffile: __ccgo_ts,
		Fline: int32(825),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6901266057277304),
		Fy:    -libc.Float64FromFloat64(0.37088021166050417),
		Fdy:   float32(-libc.Float64FromFloat64(0.496623694896698)),
		Fe:    int32(FE_INEXACT),
	},
	809: {
		Ffile: __ccgo_ts,
		Fline: int32(826),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6896554855738334),
		Fy:    -libc.Float64FromFloat64(0.37156310235052775),
		Fdy:   float32(-libc.Float64FromFloat64(0.4900000989437103)),
		Fe:    int32(FE_INEXACT),
	},
	810: {
		Ffile: __ccgo_ts,
		Fline: int32(827),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6863666374497533),
		Fy:    -libc.Float64FromFloat64(0.3763433371036471),
		Fdy:   float32(-libc.Float64FromFloat64(0.4940560460090637)),
		Fe:    int32(FE_INEXACT),
	},
	811: {
		Ffile: __ccgo_ts,
		Fline: int32(828),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.684494342239632),
		Fy:    -libc.Float64FromFloat64(0.3790748998252191),
		Fdy:   float32(-libc.Float64FromFloat64(0.4972333610057831)),
		Fe:    int32(FE_INEXACT),
	},
	812: {
		Ffile: __ccgo_ts,
		Fline: int32(829),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6842606647323933),
		Fy:    -libc.Float64FromFloat64(0.37941634516542777),
		Fdy:   float32(-libc.Float64FromFloat64(0.4910392761230469)),
		Fe:    int32(FE_INEXACT),
	},
	813: {
		Ffile: __ccgo_ts,
		Fline: int32(830),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6837935490141405),
		Fy:    -libc.Float64FromFloat64(0.3800992358458),
		Fdy:   float32(-libc.Float64FromFloat64(0.49754759669303894)),
		Fe:    int32(FE_INEXACT),
	},
	814: {
		Ffile: __ccgo_ts,
		Fline: int32(831),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6830934732688436),
		Fy:    -libc.Float64FromFloat64(0.3811235718664044),
		Fdy:   float32(-libc.Float64FromFloat64(0.49967849254608154)),
		Fe:    int32(FE_INEXACT),
	},
	815: {
		Ffile: __ccgo_ts,
		Fline: int32(832),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6819282727657349),
		Fy:    -libc.Float64FromFloat64(0.3828307985673763),
		Fdy:   float32(-libc.Float64FromFloat64(0.4931097626686096)),
		Fe:    int32(FE_INEXACT),
	},
	816: {
		Ffile: __ccgo_ts,
		Fline: int32(833),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6798359182437135),
		Fy:    -libc.Float64FromFloat64(0.3859038066291388),
		Fdy:   float32(-libc.Float64FromFloat64(0.4985475540161133)),
		Fe:    int32(FE_INEXACT),
	},
	817: {
		Ffile: __ccgo_ts,
		Fline: int32(834),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.679603831062101),
		Fy:    -libc.Float64FromFloat64(0.38624525196937665),
		Fdy:   float32(0.4902213215827942),
		Fe:    int32(FE_INEXACT),
	},
	818: {
		Ffile: __ccgo_ts,
		Fline: int32(835),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6747482611516169),
		Fy:    -libc.Float64FromFloat64(0.39341560411347637),
		Fdy:   float32(-libc.Float64FromFloat64(0.49569371342658997)),
		Fe:    int32(FE_INEXACT),
	},
	819: {
		Ffile: __ccgo_ts,
		Fline: int32(836),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6733673382662091),
		Fy:    -libc.Float64FromFloat64(0.39546427615464025),
		Fdy:   float32(0.4981290400028229),
		Fe:    int32(FE_INEXACT),
	},
	820: {
		Ffile: __ccgo_ts,
		Fline: int32(837),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.672907658959646),
		Fy:    -libc.Float64FromFloat64(0.39614716683503504),
		Fdy:   float32(-libc.Float64FromFloat64(0.49209046363830566)),
		Fe:    int32(FE_INEXACT),
	},
	821: {
		Ffile: __ccgo_ts,
		Fline: int32(838),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.671759833115219),
		Fy:    -libc.Float64FromFloat64(0.39785439353603275),
		Fdy:   float32(-libc.Float64FromFloat64(0.4901471436023712)),
		Fe:    int32(FE_INEXACT),
	},
	822: {
		Ffile: __ccgo_ts,
		Fline: int32(839),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6715305030045514),
		Fy:    -libc.Float64FromFloat64(0.3981958388762234),
		Fdy:   float32(0.4904076159000397),
		Fe:    int32(FE_INEXACT),
	},
	823: {
		Ffile: __ccgo_ts,
		Fline: int32(840),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6710720776274797),
		Fy:    -libc.Float64FromFloat64(0.39887872955660797),
		Fdy:   float32(-libc.Float64FromFloat64(0.49911725521087646)),
		Fe:    int32(FE_INEXACT),
	},
	824: {
		Ffile: __ccgo_ts,
		Fline: int32(841),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6696986783271535),
		Fy:    -libc.Float64FromFloat64(0.40092740159778956),
		Fdy:   float32(0.49907174706459045),
		Fe:    int32(FE_INEXACT),
	},
	825: {
		Ffile: __ccgo_ts,
		Fline: int32(842),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6690130330735221),
		Fy:    -libc.Float64FromFloat64(0.4019517376183745),
		Fdy:   float32(-libc.Float64FromFloat64(0.49455586075782776)),
		Fe:    int32(FE_INEXACT),
	},
	826: {
		Ffile: __ccgo_ts,
		Fline: int32(843),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6680999312335066),
		Fy:    -libc.Float64FromFloat64(0.40331751897915463),
		Fdy:   float32(-libc.Float64FromFloat64(0.4900614619255066)),
		Fe:    int32(FE_INEXACT),
	},
	827: {
		Ffile: __ccgo_ts,
		Fline: int32(844),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6678718505659452),
		Fy:    -libc.Float64FromFloat64(0.40365896431937),
		Fdy:   float32(-libc.Float64FromFloat64(0.49198153614997864)),
		Fe:    int32(FE_INEXACT),
	},
	828: {
		Ffile: __ccgo_ts,
		Fline: int32(845),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6676438477621764),
		Fy:    -libc.Float64FromFloat64(0.4040004096595773),
		Fdy:   float32(-libc.Float64FromFloat64(0.49108970165252686)),
		Fe:    int32(FE_INEXACT),
	},
	829: {
		Ffile: __ccgo_ts,
		Fline: int32(846),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6662774645912211),
		Fy:    -libc.Float64FromFloat64(0.40604908170072446),
		Fdy:   float32(-libc.Float64FromFloat64(0.4920632839202881)),
		Fe:    int32(FE_INEXACT),
	},
	830: {
		Ffile: __ccgo_ts,
		Fline: int32(847),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6653680963893964),
		Fy:    -libc.Float64FromFloat64(0.40741486306151403),
		Fdy:   float32(0.4905575215816498),
		Fe:    int32(FE_INEXACT),
	},
	831: {
		Ffile: __ccgo_ts,
		Fline: int32(848),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6631001022127874),
		Fy:    -libc.Float64FromFloat64(0.4108293164634763),
		Fdy:   float32(0.49386417865753174),
		Fe:    int32(FE_INEXACT),
	},
	832: {
		Ffile: __ccgo_ts,
		Fline: int32(849),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6624212126577099),
		Fy:    -libc.Float64FromFloat64(0.411853652484063),
		Fdy:   float32(0.49683329463005066),
		Fe:    int32(FE_INEXACT),
	},
	833: {
		Ffile: __ccgo_ts,
		Fline: int32(850),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619690058065835),
		Fy:    -libc.Float64FromFloat64(0.41253654316446137),
		Fdy:   float32(-libc.Float64FromFloat64(0.49726584553718567)),
		Fe:    int32(FE_INEXACT),
	},
	834: {
		Ffile: __ccgo_ts,
		Fline: int32(851),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6608398387909463),
		Fy:    -libc.Float64FromFloat64(0.414243769865442),
		Fdy:   float32(0.4967559278011322),
		Fe:    int32(FE_INEXACT),
	},
	835: {
		Ffile: __ccgo_ts,
		Fline: int32(852),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6603887114766576),
		Fy:    -libc.Float64FromFloat64(0.4149266605458303),
		Fdy:   float32(-libc.Float64FromFloat64(0.495334655046463)),
		Fe:    int32(FE_INEXACT),
	},
	836: {
		Ffile: __ccgo_ts,
		Fline: int32(853),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.659937892127833),
		Fy:    -libc.Float64FromFloat64(0.4156095512262321),
		Fdy:   float32(0.4940595030784607),
		Fe:    int32(FE_INEXACT),
	},
	837: {
		Ffile: __ccgo_ts,
		Fline: int32(854),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6581376901847479),
		Fy:    -libc.Float64FromFloat64(0.4183411139477983),
		Fdy:   float32(-libc.Float64FromFloat64(0.49567198753356934)),
		Fe:    int32(FE_INEXACT),
	},
	838: {
		Ffile: __ccgo_ts,
		Fline: int32(855),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6561183320980267),
		Fy:    -libc.Float64FromFloat64(0.4214141220095664),
		Fdy:   float32(0.49234333634376526),
		Fe:    int32(FE_INEXACT),
	},
	839: {
		Ffile: __ccgo_ts,
		Fline: int32(856),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6558943417933121),
		Fy:    -libc.Float64FromFloat64(0.421755567349765),
		Fdy:   float32(-libc.Float64FromFloat64(0.49299314618110657)),
		Fe:    int32(FE_INEXACT),
	},
	840: {
		Ffile: __ccgo_ts,
		Fline: int32(857),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6554465905599487),
		Fy:    -libc.Float64FromFloat64(0.42243845803016217),
		Fdy:   float32(0.49023500084877014),
		Fe:    int32(FE_INEXACT),
	},
	841: {
		Ffile: __ccgo_ts,
		Fline: int32(858),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6538818669527107),
		Fy:    -libc.Float64FromFloat64(0.42482857541156416),
		Fdy:   float32(-libc.Float64FromFloat64(0.49208420515060425)),
		Fe:    int32(FE_INEXACT),
	},
	842: {
		Ffile: __ccgo_ts,
		Fline: int32(859),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.650319368337433),
		Fy:    -libc.Float64FromFloat64(0.430291700854673),
		Fdy:   float32(-libc.Float64FromFloat64(0.49146831035614014)),
		Fe:    int32(FE_INEXACT),
	},
	843: {
		Ffile: __ccgo_ts,
		Fline: int32(860),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6467762790329339),
		Fy:    -libc.Float64FromFloat64(0.4357548262978189),
		Fdy:   float32(-libc.Float64FromFloat64(0.4970642626285553)),
		Fe:    int32(FE_INEXACT),
	},
	844: {
		Ffile: __ccgo_ts,
		Fline: int32(861),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.646555477984203),
		Fy:    -libc.Float64FromFloat64(0.436096271638008),
		Fdy:   float32(0.4934425950050354),
		Fe:    int32(FE_INEXACT),
	},
	845: {
		Ffile: __ccgo_ts,
		Fline: int32(862),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6463347523140877),
		Fy:    -libc.Float64FromFloat64(0.4364377169782035),
		Fdy:   float32(-libc.Float64FromFloat64(0.4996820390224457)),
		Fe:    int32(FE_INEXACT),
	},
	846: {
		Ffile: __ccgo_ts,
		Fline: int32(863),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.645011979803946),
		Fy:    -libc.Float64FromFloat64(0.43848638901939574),
		Fdy:   float32(0.4959268271923065),
		Fe:    int32(FE_INEXACT),
	},
	847: {
		Ffile: __ccgo_ts,
		Fline: int32(864),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6443516090778327),
		Fy:    -libc.Float64FromFloat64(0.4395107250399715),
		Fdy:   float32(-libc.Float64FromFloat64(0.49780476093292236)),
		Fe:    int32(FE_INEXACT),
	},
	848: {
		Ffile: __ccgo_ts,
		Fline: int32(865),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6441316357800307),
		Fy:    -libc.Float64FromFloat64(0.4398521703801682),
		Fdy:   float32(-libc.Float64FromFloat64(0.4980933666229248)),
		Fe:    int32(FE_INEXACT),
	},
	849: {
		Ffile: __ccgo_ts,
		Fline: int32(866),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6436919144468954),
		Fy:    -libc.Float64FromFloat64(0.44053506106056634),
		Fdy:   float32(-libc.Float64FromFloat64(0.4903792440891266)),
		Fe:    int32(FE_INEXACT),
	},
	850: {
		Ffile: __ccgo_ts,
		Fline: int32(867),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6434721663603026),
		Fy:    -libc.Float64FromFloat64(0.4408765064007595),
		Fdy:   float32(-libc.Float64FromFloat64(0.49915575981140137)),
		Fe:    int32(FE_INEXACT),
	},
	851: {
		Ffile: __ccgo_ts,
		Fline: int32(868),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6428133721129602),
		Fy:    -libc.Float64FromFloat64(0.4419008424214045),
		Fdy:   float32(0.49039793014526367),
		Fe:    int32(FE_INEXACT),
	},
	852: {
		Ffile: __ccgo_ts,
		Fline: int32(869),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6425939239493832),
		Fy:    -libc.Float64FromFloat64(0.44224228776154106),
		Fdy:   float32(0.4922305643558502),
		Fe:    int32(FE_INEXACT),
	},
	853: {
		Ffile: __ccgo_ts,
		Fline: int32(870),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6423745507025285),
		Fy:    -libc.Float64FromFloat64(0.44258373310173976),
		Fdy:   float32(-libc.Float64FromFloat64(0.49324601888656616)),
		Fe:    int32(FE_INEXACT),
	},
	854: {
		Ffile: __ccgo_ts,
		Fline: int32(871),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6401849330335967),
		Fy:    -libc.Float64FromFloat64(0.445998186503708),
		Fdy:   float32(-libc.Float64FromFloat64(0.49471238255500793)),
		Fe:    int32(FE_INEXACT),
	},
	855: {
		Ffile: __ccgo_ts,
		Fline: int32(872),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6399663821850442),
		Fy:    -libc.Float64FromFloat64(0.446339631843923),
		Fdy:   float32(0.49022406339645386),
		Fe:    int32(FE_INEXACT),
	},
	856: {
		Ffile: __ccgo_ts,
		Fline: int32(873),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6377849730734735),
		Fy:    -libc.Float64FromFloat64(0.44975408524586424),
		Fdy:   float32(-libc.Float64FromFloat64(0.49421465396881104)),
		Fe:    int32(FE_INEXACT),
	},
	857: {
		Ffile: __ccgo_ts,
		Fline: int32(874),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6375672415401911),
		Fy:    -libc.Float64FromFloat64(0.4500955305860658),
		Fdy:   float32(-libc.Float64FromFloat64(0.4992421567440033)),
		Fe:    int32(FE_INEXACT),
	},
	858: {
		Ffile: __ccgo_ts,
		Fline: int32(875),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6373495843376313),
		Fy:    -libc.Float64FromFloat64(0.4504369759262737),
		Fdy:   float32(0.49009349942207336),
		Fe:    int32(FE_INEXACT),
	},
	859: {
		Ffile: __ccgo_ts,
		Fline: int32(876),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6366970584606351),
		Fy:    -libc.Float64FromFloat64(0.4514613119468437),
		Fdy:   float32(0.492910236120224),
		Fe:    int32(FE_INEXACT),
	},
	860: {
		Ffile: __ccgo_ts,
		Fline: int32(877),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6364796983273023),
		Fy:    -libc.Float64FromFloat64(0.4518027572870695),
		Fdy:   float32(-libc.Float64FromFloat64(0.49270951747894287)),
		Fe:    int32(FE_INEXACT),
	},
	861: {
		Ffile: __ccgo_ts,
		Fline: int32(878),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6356109995801411),
		Fy:    -libc.Float64FromFloat64(0.4531685386478269),
		Fdy:   float32(-libc.Float64FromFloat64(0.4985910952091217)),
		Fe:    int32(FE_INEXACT),
	},
	862: {
		Ffile: __ccgo_ts,
		Fline: int32(879),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6351770949237577),
		Fy:    -libc.Float64FromFloat64(0.45385142932822004),
		Fdy:   float32(-libc.Float64FromFloat64(0.49607616662979126)),
		Fe:    int32(FE_INEXACT),
	},
	863: {
		Ffile: __ccgo_ts,
		Fline: int32(880),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6343101740336697),
		Fy:    -libc.Float64FromFloat64(0.45521721068900484),
		Fdy:   float32(0.4904749393463135),
		Fe:    int32(FE_INEXACT),
	},
	864: {
		Ffile: __ccgo_ts,
		Fline: int32(881),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6332281866295435),
		Fy:    -libc.Float64FromFloat64(0.4569244373899922),
		Fdy:   float32(0.49066975712776184),
		Fe:    int32(FE_INEXACT),
	},
	865: {
		Ffile: __ccgo_ts,
		Fline: int32(882),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6323639257052004),
		Fy:    -libc.Float64FromFloat64(0.4582902187507738),
		Fdy:   float32(0.4926498532295227),
		Fe:    int32(FE_INEXACT),
	},
	866: {
		Ffile: __ccgo_ts,
		Fline: int32(883),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6319322376883539),
		Fy:    -libc.Float64FromFloat64(0.4589731094311724),
		Fdy:   float32(-libc.Float64FromFloat64(0.49676454067230225)),
		Fe:    int32(FE_INEXACT),
	},
	867: {
		Ffile: __ccgo_ts,
		Fline: int32(884),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6315008443665936),
		Fy:    -libc.Float64FromFloat64(0.4596560001115768),
		Fdy:   float32(0.4911830425262451),
		Fe:    int32(FE_INEXACT),
	},
	868: {
		Ffile: __ccgo_ts,
		Fline: int32(885),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.628703957087348),
		Fy:    -libc.Float64FromFloat64(0.464094789534119),
		Fdy:   float32(-libc.Float64FromFloat64(0.49373435974121094)),
		Fe:    int32(FE_INEXACT),
	},
	869: {
		Ffile: __ccgo_ts,
		Fline: int32(886),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6269889561779067),
		Fy:    -libc.Float64FromFloat64(0.4668263522556852),
		Fdy:   float32(0.49830731749534607),
		Fe:    int32(FE_INEXACT),
	},
	870: {
		Ffile: __ccgo_ts,
		Fline: int32(887),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6267749102650293),
		Fy:    -libc.Float64FromFloat64(0.4671677975958899),
		Fdy:   float32(0.4994850158691406),
		Fe:    int32(FE_INEXACT),
	},
	871: {
		Ffile: __ccgo_ts,
		Fline: int32(888),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6259194570891667),
		Fy:    -libc.Float64FromFloat64(0.4685335789566644),
		Fdy:   float32(0.4981178641319275),
		Fe:    int32(FE_INEXACT),
	},
	872: {
		Ffile: __ccgo_ts,
		Fline: int32(889),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6257057762894396),
		Fy:    -libc.Float64FromFloat64(0.4688750242968584),
		Fdy:   float32(0.49767228960990906),
		Fe:    int32(FE_INEXACT),
	},
	873: {
		Ffile: __ccgo_ts,
		Fline: int32(890),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6244252225247747),
		Fy:    -libc.Float64FromFloat64(0.47092369633804326),
		Fdy:   float32(0.49830612540245056),
		Fe:    int32(FE_INEXACT),
	},
	874: {
		Ffile: __ccgo_ts,
		Fline: int32(891),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6210232104748497),
		Fy:    -libc.Float64FromFloat64(0.4763868217811808),
		Fdy:   float32(0.4952779710292816),
		Fe:    int32(FE_INEXACT),
	},
	875: {
		Ffile: __ccgo_ts,
		Fline: int32(892),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6205992642830304),
		Fy:    -libc.Float64FromFloat64(0.4770697124615822),
		Fdy:   float32(0.4934687912464142),
		Fe:    int32(FE_INEXACT),
	},
	876: {
		Ffile: __ccgo_ts,
		Fline: int32(893),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6174288791588094),
		Fy:    -libc.Float64FromFloat64(0.4821913925645272),
		Fdy:   float32(0.492437481880188),
		Fe:    int32(FE_INEXACT),
	},
	877: {
		Ffile: __ccgo_ts,
		Fline: int32(894),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6167967483302021),
		Fy:    -libc.Float64FromFloat64(0.48321572858510714),
		Fdy:   float32(0.4969865083694458),
		Fe:    int32(FE_INEXACT),
	},
	878: {
		Ffile: __ccgo_ts,
		Fline: int32(895),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6151142283263032),
		Fy:    -libc.Float64FromFloat64(0.48594729130668146),
		Fdy:   float32(0.4954165816307068),
		Fe:    int32(FE_INEXACT),
	},
	879: {
		Ffile: __ccgo_ts,
		Fline: int32(896),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6128082547822366),
		Fy:    -libc.Float64FromFloat64(0.48970319004883867),
		Fdy:   float32(0.49839410185813904),
		Fe:    int32(FE_INEXACT),
	},
	880: {
		Ffile: __ccgo_ts,
		Fline: int32(897),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6117629447116112),
		Fy:    -libc.Float64FromFloat64(0.4914104167498224),
		Fdy:   float32(-libc.Float64FromFloat64(0.49069786071777344)),
		Fe:    int32(FE_INEXACT),
	},
	881: {
		Ffile: __ccgo_ts,
		Fline: int32(898),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6113453201102602),
		Fy:    -libc.Float64FromFloat64(0.49209330743021745),
		Fdy:   float32(0.496625155210495),
		Fe:    int32(FE_INEXACT),
	},
	882: {
		Ffile: __ccgo_ts,
		Fline: int32(899),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6098858778485787),
		Fy:    -libc.Float64FromFloat64(0.49448342481158747),
		Fdy:   float32(-libc.Float64FromFloat64(0.49243125319480896)),
		Fe:    int32(FE_INEXACT),
	},
	883: {
		Ffile: __ccgo_ts,
		Fline: int32(900),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6090534756534998),
		Fy:    -libc.Float64FromFloat64(0.49584920617237704),
		Fdy:   float32(-libc.Float64FromFloat64(0.4913181960582733)),
		Fe:    int32(FE_INEXACT),
	},
	884: {
		Ffile: __ccgo_ts,
		Fline: int32(901),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6086377006916709),
		Fy:    -libc.Float64FromFloat64(0.49653209685281124),
		Fdy:   float32(0.49005696177482605),
		Fe:    int32(FE_INEXACT),
	},
	885: {
		Ffile: __ccgo_ts,
		Fline: int32(902),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6042891629542656),
		Fy:    -libc.Float64FromFloat64(0.5037024489968965),
		Fdy:   float32(-libc.Float64FromFloat64(0.49144256114959717)),
		Fe:    int32(FE_INEXACT),
	},
	886: {
		Ffile: __ccgo_ts,
		Fline: int32(903),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6036704847191009),
		Fy:    -libc.Float64FromFloat64(0.5047267850174793),
		Fdy:   float32(-libc.Float64FromFloat64(0.49353837966918945)),
		Fe:    int32(FE_INEXACT),
	},
	887: {
		Ffile: __ccgo_ts,
		Fline: int32(904),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6026407615845096),
		Fy:    -libc.Float64FromFloat64(0.5064340117184671),
		Fdy:   float32(0.4927366375923157),
		Fe:    int32(FE_INEXACT),
	},
	888: {
		Ffile: __ccgo_ts,
		Fline: int32(905),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6020237710017339),
		Fy:    -libc.Float64FromFloat64(0.5074583477390521),
		Fdy:   float32(0.49998369812965393),
		Fe:    int32(FE_INEXACT),
	},
	889: {
		Ffile: __ccgo_ts,
		Fline: int32(906),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6016127949209797),
		Fy:    -libc.Float64FromFloat64(0.5081412384194467),
		Fdy:   float32(-libc.Float64FromFloat64(0.490567147731781)),
		Fe:    int32(FE_INEXACT),
	},
	890: {
		Ffile: __ccgo_ts,
		Fline: int32(907),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5997668716732534),
		Fy:    -libc.Float64FromFloat64(0.5112142464812612),
		Fdy:   float32(0.4913415014743805),
		Fe:    int32(FE_INEXACT),
	},
	891: {
		Ffile: __ccgo_ts,
		Fline: int32(908),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5985394038209056),
		Fy:    -libc.Float64FromFloat64(0.5132629185223958),
		Fdy:   float32(-libc.Float64FromFloat64(0.4982002377510071)),
		Fe:    int32(FE_INEXACT),
	},
	892: {
		Ffile: __ccgo_ts,
		Fline: int32(909),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5983350702169109),
		Fy:    -libc.Float64FromFloat64(0.5136043638625919),
		Fdy:   float32(-libc.Float64FromFloat64(0.4985455274581909)),
		Fe:    int32(FE_INEXACT),
	},
	893: {
		Ffile: __ccgo_ts,
		Fline: int32(910),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.597314448073487),
		Fy:    -libc.Float64FromFloat64(0.5153115905635691),
		Fdy:   float32(-libc.Float64FromFloat64(0.4930005371570587)),
		Fe:    int32(FE_INEXACT),
	},
	894: {
		Ffile: __ccgo_ts,
		Fline: int32(911),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5967029106320941),
		Fy:    -libc.Float64FromFloat64(0.5163359265841573),
		Fdy:   float32(0.4967649281024933),
		Fe:    int32(FE_INEXACT),
	},
	895: {
		Ffile: __ccgo_ts,
		Fline: int32(912),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5962955668769181),
		Fy:    -libc.Float64FromFloat64(0.5170188172645488),
		Fdy:   float32(0.49142730236053467),
		Fe:    int32(FE_INEXACT),
	},
	896: {
		Ffile: __ccgo_ts,
		Fline: int32(913),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5946689707214952),
		Fy:    -libc.Float64FromFloat64(0.5197503799861358),
		Fdy:   float32(-libc.Float64FromFloat64(0.4902055263519287)),
		Fe:    int32(FE_INEXACT),
	},
	897: {
		Ffile: __ccgo_ts,
		Fline: int32(914),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.594465958433265),
		Fy:    -libc.Float64FromFloat64(0.5200918253263157),
		Fdy:   float32(-libc.Float64FromFloat64(0.49516722559928894)),
		Fe:    int32(FE_INEXACT),
	},
	898: {
		Ffile: __ccgo_ts,
		Fline: int32(915),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5926419637611084),
		Fy:    -libc.Float64FromFloat64(0.5231648333880832),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998883605003357)),
		Fe:    int32(FE_INEXACT),
	},
	899: {
		Ffile: __ccgo_ts,
		Fline: int32(916),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5924396434668038),
		Fy:    -libc.Float64FromFloat64(0.5235062787282796),
		Fdy:   float32(0.49842125177383423),
		Fe:    int32(FE_INEXACT),
	},
	900: {
		Ffile: __ccgo_ts,
		Fline: int32(917),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5904202354599157),
		Fy:    -libc.Float64FromFloat64(0.5269207321302458),
		Fdy:   float32(0.4905199706554413),
		Fe:    int32(FE_INEXACT),
	},
	901: {
		Ffile: __ccgo_ts,
		Fline: int32(918),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.58961440092945),
		Fy:    -libc.Float64FromFloat64(0.5282865134910296),
		Fdy:   float32(-libc.Float64FromFloat64(0.49761608242988586)),
		Fe:    int32(FE_INEXACT),
	},
	902: {
		Ffile: __ccgo_ts,
		Fline: int32(919),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5894131142059021),
		Fy:    -libc.Float64FromFloat64(0.5286279588312266),
		Fdy:   float32(-libc.Float64FromFloat64(0.491918683052063)),
		Fe:    int32(FE_INEXACT),
	},
	903: {
		Ffile: __ccgo_ts,
		Fline: int32(920),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5874040225198262),
		Fy:    -libc.Float64FromFloat64(0.5320424122331892),
		Fdy:   float32(0.4935992360115051),
		Fe:    int32(FE_INEXACT),
	},
	904: {
		Ffile: __ccgo_ts,
		Fline: int32(921),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5858016810338561),
		Fy:    -libc.Float64FromFloat64(0.5347739749547652),
		Fdy:   float32(0.4976060092449188),
		Fe:    int32(FE_INEXACT),
	},
	905: {
		Ffile: __ccgo_ts,
		Fline: int32(922),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5856016959235215),
		Fy:    -libc.Float64FromFloat64(0.535115420294964),
		Fdy:   float32(-libc.Float64FromFloat64(0.49679431319236755)),
		Fe:    int32(FE_INEXACT),
	},
	906: {
		Ffile: __ccgo_ts,
		Fline: int32(923),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5854017790854964),
		Fy:    -libc.Float64FromFloat64(0.5354568656351938),
		Fdy:   float32(-libc.Float64FromFloat64(0.490285187959671)),
		Fe:    int32(FE_INEXACT),
	},
	907: {
		Ffile: __ccgo_ts,
		Fline: int32(924),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5838048993900695),
		Fy:    -libc.Float64FromFloat64(0.538188428356725),
		Fdy:   float32(-libc.Float64FromFloat64(0.49117255210876465)),
		Fe:    int32(FE_INEXACT),
	},
	908: {
		Ffile: __ccgo_ts,
		Fline: int32(925),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5818149240376601),
		Fy:    -libc.Float64FromFloat64(0.5416028817586923),
		Fdy:   float32(-libc.Float64FromFloat64(0.49548885226249695)),
		Fe:    int32(FE_INEXACT),
	},
	909: {
		Ffile: __ccgo_ts,
		Fline: int32(926),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5816162999545564),
		Fy:    -libc.Float64FromFloat64(0.541944327098885),
		Fdy:   float32(-libc.Float64FromFloat64(0.4922885596752167)),
		Fe:    int32(FE_INEXACT),
	},
	910: {
		Ffile: __ccgo_ts,
		Fline: int32(927),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5810208344587946),
		Fy:    -libc.Float64FromFloat64(0.5429686631194767),
		Fdy:   float32(-libc.Float64FromFloat64(0.4990759789943695)),
		Fe:    int32(FE_INEXACT),
	},
	911: {
		Ffile: __ccgo_ts,
		Fline: int32(928),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.580822481467591),
		Fy:    -libc.Float64FromFloat64(0.5433101084596731),
		Fdy:   float32(-libc.Float64FromFloat64(0.4915362298488617)),
		Fe:    int32(FE_INEXACT),
	},
	912: {
		Ffile: __ccgo_ts,
		Fline: int32(929),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5806241961914859),
		Fy:    -libc.Float64FromFloat64(0.5436515537999467),
		Fdy:   float32(0.49100857973098755),
		Fe:    int32(FE_INEXACT),
	},
	913: {
		Ffile: __ccgo_ts,
		Fline: int32(930),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5804259786074982),
		Fy:    -libc.Float64FromFloat64(0.5439929991400634),
		Fdy:   float32(0.49495047330856323),
		Fe:    int32(FE_INEXACT),
	},
	914: {
		Ffile: __ccgo_ts,
		Fline: int32(931),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5802278286923805),
		Fy:    -libc.Float64FromFloat64(0.5443344444802615),
		Fdy:   float32(-libc.Float64FromFloat64(0.4998927712440491)),
		Fe:    int32(FE_INEXACT),
	},
	915: {
		Ffile: __ccgo_ts,
		Fline: int32(932),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5790403489558092),
		Fy:    -libc.Float64FromFloat64(0.5463831165214573),
		Fdy:   float32(-libc.Float64FromFloat64(0.4930100739002228)),
		Fe:    int32(FE_INEXACT),
	},
	916: {
		Ffile: __ccgo_ts,
		Fline: int32(933),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5786450626819583),
		Fy:    -libc.Float64FromFloat64(0.5470660072018421),
		Fdy:   float32(0.4942411780357361),
		Fe:    int32(FE_INEXACT),
	},
	917: {
		Ffile: __ccgo_ts,
		Fline: int32(934),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5774608221946624),
		Fy:    -libc.Float64FromFloat64(0.5491146792430123),
		Fdy:   float32(0.4911934733390808),
		Fe:    int32(FE_INEXACT),
	},
	918: {
		Ffile: __ccgo_ts,
		Fline: int32(935),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5760822711521428),
		Fy:    -libc.Float64FromFloat64(0.5515047966243858),
		Fdy:   float32(0.49824458360671997),
		Fe:    int32(FE_INEXACT),
	},
	919: {
		Ffile: __ccgo_ts,
		Fline: int32(936),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5756890042324652),
		Fy:    -libc.Float64FromFloat64(0.5521876873047934),
		Fdy:   float32(0.49369579553604126),
		Fe:    int32(FE_INEXACT),
	},
	920: {
		Ffile: __ccgo_ts,
		Fline: int32(937),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5750996071706281),
		Fy:    -libc.Float64FromFloat64(0.5532120233253847),
		Fdy:   float32(-libc.Float64FromFloat64(0.4951796531677246)),
		Fe:    int32(FE_INEXACT),
	},
	921: {
		Ffile: __ccgo_ts,
		Fline: int32(938),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5747070110739704),
		Fy:    -libc.Float64FromFloat64(0.5538949140057566),
		Fdy:   float32(0.49030524492263794),
		Fe:    int32(FE_INEXACT),
	},
	922: {
		Ffile: __ccgo_ts,
		Fline: int32(939),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5741186193880897),
		Fy:    -libc.Float64FromFloat64(0.5549192500263452),
		Fdy:   float32(0.4954734146595001),
		Fe:    int32(FE_INEXACT),
	},
	923: {
		Ffile: __ccgo_ts,
		Fline: int32(940),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5731393049452781),
		Fy:    -libc.Float64FromFloat64(0.5566264767273362),
		Fdy:   float32(-libc.Float64FromFloat64(0.49742624163627625)),
		Fe:    int32(FE_INEXACT),
	},
	924: {
		Ffile: __ccgo_ts,
		Fline: int32(941),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.572552518295069),
		Fy:    -libc.Float64FromFloat64(0.5576508127479204),
		Fdy:   float32(-libc.Float64FromFloat64(0.4930580258369446)),
		Fe:    int32(FE_INEXACT),
	},
	925: {
		Ffile: __ccgo_ts,
		Fline: int32(942),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5669111283779115),
		Fy:    -libc.Float64FromFloat64(0.5675527276136244),
		Fdy:   float32(-libc.Float64FromFloat64(0.49354881048202515)),
		Fe:    int32(FE_INEXACT),
	},
	926: {
		Ffile: __ccgo_ts,
		Fline: int32(943),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5665241222082267),
		Fy:    -libc.Float64FromFloat64(0.5682356182940127),
		Fdy:   float32(0.49413731694221497),
		Fe:    int32(FE_INEXACT),
	},
	927: {
		Ffile: __ccgo_ts,
		Fline: int32(944),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5661373802312307),
		Fy:    -libc.Float64FromFloat64(0.5689185089744013),
		Fdy:   float32(-libc.Float64FromFloat64(0.4979185461997986)),
		Fe:    int32(FE_INEXACT),
	},
	928: {
		Ffile: __ccgo_ts,
		Fline: int32(945),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5659441082586241),
		Fy:    -libc.Float64FromFloat64(0.5692599543145987),
		Fdy:   float32(0.49293625354766846),
		Fe:    int32(FE_INEXACT),
	},
	929: {
		Ffile: __ccgo_ts,
		Fline: int32(946),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5657509022665668),
		Fy:    -libc.Float64FromFloat64(0.5696013996547964),
		Fdy:   float32(0.4909191429615021),
		Fe:    int32(FE_INEXACT),
	},
	930: {
		Ffile: __ccgo_ts,
		Fline: int32(947),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5655577622325352),
		Fy:    -libc.Float64FromFloat64(0.5699428449949919),
		Fdy:   float32(0.4969032108783722),
		Fe:    int32(FE_INEXACT),
	},
	931: {
		Ffile: __ccgo_ts,
		Fline: int32(948),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5647858612264249),
		Fy:    -libc.Float64FromFloat64(0.5713086263557742),
		Fdy:   float32(-libc.Float64FromFloat64(0.494773805141449)),
		Fe:    int32(FE_INEXACT),
	},
	932: {
		Ffile: __ccgo_ts,
		Fline: int32(949),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5644003058864343),
		Fy:    -libc.Float64FromFloat64(0.5719915170361652),
		Fdy:   float32(-libc.Float64FromFloat64(0.4972938597202301)),
		Fe:    int32(FE_INEXACT),
	},
	933: {
		Ffile: __ccgo_ts,
		Fline: int32(950),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5628607147529939),
		Fy:    -libc.Float64FromFloat64(0.5747230797577495),
		Fdy:   float32(-libc.Float64FromFloat64(0.4937868118286133)),
		Fe:    int32(FE_INEXACT),
	},
	934: {
		Ffile: __ccgo_ts,
		Fline: int32(951),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5620924948097227),
		Fy:    -libc.Float64FromFloat64(0.5760888611185351),
		Fdy:   float32(0.4977673888206482),
		Fe:    int32(FE_INEXACT),
	},
	935: {
		Ffile: __ccgo_ts,
		Fline: int32(952),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5613253233707604),
		Fy:    -libc.Float64FromFloat64(0.5774546424793098),
		Fdy:   float32(-libc.Float64FromFloat64(0.49158912897109985)),
		Fe:    int32(FE_INEXACT),
	},
	936: {
		Ffile: __ccgo_ts,
		Fline: int32(953),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5611336941720728),
		Fy:    -libc.Float64FromFloat64(0.5777960878195035),
		Fdy:   float32(0.4980987012386322),
		Fe:    int32(FE_INEXACT),
	},
	937: {
		Ffile: __ccgo_ts,
		Fline: int32(954),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5605591990050461),
		Fy:    -libc.Float64FromFloat64(0.5788204238400922),
		Fdy:   float32(-libc.Float64FromFloat64(0.49866196513175964)),
		Fe:    int32(FE_INEXACT),
	},
	938: {
		Ffile: __ccgo_ts,
		Fline: int32(955),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5603678313512954),
		Fy:    -libc.Float64FromFloat64(0.5791618691802927),
		Fdy:   float32(-libc.Float64FromFloat64(0.4990600347518921)),
		Fe:    int32(FE_INEXACT),
	},
	939: {
		Ffile: __ccgo_ts,
		Fline: int32(956),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5594119725932027),
		Fy:    -libc.Float64FromFloat64(0.5808690958812736),
		Fdy:   float32(0.49269190430641174),
		Fe:    int32(FE_INEXACT),
	},
	940: {
		Ffile: __ccgo_ts,
		Fline: int32(957),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5592209965876975),
		Fy:    -libc.Float64FromFloat64(0.5812105412214739),
		Fdy:   float32(0.4975578784942627),
		Fe:    int32(FE_INEXACT),
	},
	941: {
		Ffile: __ccgo_ts,
		Fline: int32(958),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5576955337574534),
		Fy:    -libc.Float64FromFloat64(0.5839421039431474),
		Fdy:   float32(-libc.Float64FromFloat64(0.4900452196598053)),
		Fe:    int32(FE_INEXACT),
	},
	942: {
		Ffile: __ccgo_ts,
		Fline: int32(959),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5571245586188904),
		Fy:    -libc.Float64FromFloat64(0.5849664399636278),
		Fdy:   float32(0.49661991000175476),
		Fe:    int32(FE_INEXACT),
	},
	943: {
		Ffile: __ccgo_ts,
		Fline: int32(960),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5567442333250698),
		Fy:    -libc.Float64FromFloat64(0.5856493306440228),
		Fdy:   float32(0.49516603350639343),
		Fe:    int32(FE_INEXACT),
	},
	944: {
		Ffile: __ccgo_ts,
		Fline: int32(961),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5565541680512164),
		Fy:    -libc.Float64FromFloat64(0.5859907759842211),
		Fdy:   float32(-libc.Float64FromFloat64(0.49516746401786804)),
		Fe:    int32(FE_INEXACT),
	},
	945: {
		Ffile: __ccgo_ts,
		Fline: int32(962),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5561742321388329),
		Fy:    -libc.Float64FromFloat64(0.5866736666646127),
		Fdy:   float32(-libc.Float64FromFloat64(0.4975254237651825)),
		Fe:    int32(FE_INEXACT),
	},
	946: {
		Ffile: __ccgo_ts,
		Fline: int32(963),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5552255266973865),
		Fy:    -libc.Float64FromFloat64(0.5883808933655909),
		Fdy:   float32(0.4947531223297119),
		Fe:    int32(FE_INEXACT),
	},
	947: {
		Ffile: __ccgo_ts,
		Fline: int32(964),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5533329678745501),
		Fy:    -libc.Float64FromFloat64(0.5917953467675565),
		Fdy:   float32(-libc.Float64FromFloat64(0.4947388172149658)),
		Fe:    int32(FE_INEXACT),
	},
	948: {
		Ffile: __ccgo_ts,
		Fline: int32(965),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5531440671625558),
		Fy:    -libc.Float64FromFloat64(0.592136792107763),
		Fdy:   float32(0.49022653698921204),
		Fe:    int32(FE_INEXACT),
	},
	949: {
		Ffile: __ccgo_ts,
		Fline: int32(966),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5520120163737606),
		Fy:    -libc.Float64FromFloat64(0.5941854641489339),
		Fdy:   float32(-libc.Float64FromFloat64(0.4908985495567322)),
		Fe:    int32(FE_INEXACT),
	},
	950: {
		Ffile: __ccgo_ts,
		Fline: int32(967),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5514468600859794),
		Fy:    -libc.Float64FromFloat64(0.595209800169522),
		Fdy:   float32(0.49641019105911255),
		Fe:    int32(FE_INEXACT),
	},
	951: {
		Ffile: __ccgo_ts,
		Fline: int32(968),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5510704107160018),
		Fy:    -libc.Float64FromFloat64(0.5958926908499251),
		Fdy:   float32(-libc.Float64FromFloat64(0.4928133487701416)),
		Fe:    int32(FE_INEXACT),
	},
	952: {
		Ffile: __ccgo_ts,
		Fline: int32(969),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5506942183320461),
		Fy:    -libc.Float64FromFloat64(0.596575581530306),
		Fdy:   float32(0.4946102499961853),
		Fe:    int32(FE_INEXACT),
	},
	953: {
		Ffile: __ccgo_ts,
		Fline: int32(970),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5503182827586555),
		Fy:    -libc.Float64FromFloat64(0.5972584722107075),
		Fdy:   float32(-libc.Float64FromFloat64(0.49783188104629517)),
		Fe:    int32(FE_INEXACT),
	},
	954: {
		Ffile: __ccgo_ts,
		Fline: int32(971),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5501304112211418),
		Fy:    -libc.Float64FromFloat64(0.5975999175508977),
		Fdy:   float32(-libc.Float64FromFloat64(0.49071335792541504)),
		Fe:    int32(FE_INEXACT),
	},
	955: {
		Ffile: __ccgo_ts,
		Fline: int32(972),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5495671813424795),
		Fy:    -libc.Float64FromFloat64(0.5986242535714893),
		Fdy:   float32(-libc.Float64FromFloat64(0.4979199469089508)),
		Fe:    int32(FE_INEXACT),
	},
	956: {
		Ffile: __ccgo_ts,
		Fline: int32(973),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5476939097298121),
		Fy:    -libc.Float64FromFloat64(0.6020387069734532),
		Fdy:   float32(-libc.Float64FromFloat64(0.49402520060539246)),
		Fe:    int32(FE_INEXACT),
	},
	957: {
		Ffile: __ccgo_ts,
		Fline: int32(974),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5475069341192733),
		Fy:    -libc.Float64FromFloat64(0.6023801523136488),
		Fdy:   float32(0.49635785818099976),
		Fe:    int32(FE_INEXACT),
	},
	958: {
		Ffile: __ccgo_ts,
		Fline: int32(975),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5473200223397838),
		Fy:    -libc.Float64FromFloat64(0.6027215976538511),
		Fdy:   float32(0.4971492886543274),
		Fe:    int32(FE_INEXACT),
	},
	959: {
		Ffile: __ccgo_ts,
		Fline: int32(976),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.546946390186817),
		Fy:    -libc.Float64FromFloat64(0.603404488334237),
		Fdy:   float32(-libc.Float64FromFloat64(0.4963313937187195)),
		Fe:    int32(FE_INEXACT),
	},
	960: {
		Ffile: __ccgo_ts,
		Fline: int32(977),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5465730130966605),
		Fy:    -libc.Float64FromFloat64(0.6040873790146316),
		Fdy:   float32(0.4992106258869171),
		Fe:    int32(FE_INEXACT),
	},
	961: {
		Ffile: __ccgo_ts,
		Fline: int32(978),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5435951577672431),
		Fy:    -libc.Float64FromFloat64(0.6095505044577673),
		Fdy:   float32(0.494458943605423),
		Fe:    int32(FE_INEXACT),
	},
	962: {
		Ffile: __ccgo_ts,
		Fline: int32(979),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5432240684212368),
		Fy:    -libc.Float64FromFloat64(0.6102333951381776),
		Fdy:   float32(-libc.Float64FromFloat64(0.49110326170921326)),
		Fe:    int32(FE_INEXACT),
	},
	963: {
		Ffile: __ccgo_ts,
		Fline: int32(980),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5408181547945178),
		Fy:    -libc.Float64FromFloat64(0.6146721845607163),
		Fdy:   float32(-libc.Float64FromFloat64(0.4990432858467102)),
		Fe:    int32(FE_INEXACT),
	},
	964: {
		Ffile: __ccgo_ts,
		Fline: int32(981),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5406335264777039),
		Fy:    -libc.Float64FromFloat64(0.6150136299009101),
		Fdy:   float32(-libc.Float64FromFloat64(0.4966411888599396)),
		Fe:    int32(FE_INEXACT),
	},
	965: {
		Ffile: __ccgo_ts,
		Fline: int32(982),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5402644589117088),
		Fy:    -libc.Float64FromFloat64(0.6156965205813006),
		Fdy:   float32(-libc.Float64FromFloat64(0.4951445162296295)),
		Fe:    int32(FE_INEXACT),
	},
	966: {
		Ffile: __ccgo_ts,
		Fline: int32(983),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5397113299091432),
		Fy:    -libc.Float64FromFloat64(0.6167208566018969),
		Fdy:   float32(-libc.Float64FromFloat64(0.49299871921539307)),
		Fe:    int32(FE_INEXACT),
	},
	967: {
		Ffile: __ccgo_ts,
		Fline: int32(984),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5384228968450299),
		Fy:    -libc.Float64FromFloat64(0.6191109739833982),
		Fdy:   float32(0.4904833137989044),
		Fe:    int32(FE_INEXACT),
	},
	968: {
		Ffile: __ccgo_ts,
		Fline: int32(985),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5373209740346317),
		Fy:    -libc.Float64FromFloat64(0.6211596460244478),
		Fdy:   float32(0.49381476640701294),
		Fe:    int32(FE_INEXACT),
	},
	969: {
		Ffile: __ccgo_ts,
		Fline: int32(986),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5365876119836465),
		Fy:    -libc.Float64FromFloat64(0.6225254273852301),
		Fdy:   float32(0.4901638627052307),
		Fe:    int32(FE_INEXACT),
	},
	970: {
		Ffile: __ccgo_ts,
		Fline: int32(987),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5360382473794971),
		Fy:    -libc.Float64FromFloat64(0.6235497634058277),
		Fdy:   float32(0.49919766187667847),
		Fe:    int32(FE_INEXACT),
	},
	971: {
		Ffile: __ccgo_ts,
		Fline: int32(988),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5358552508611946),
		Fy:    -libc.Float64FromFloat64(0.623891208746024),
		Fdy:   float32(-libc.Float64FromFloat64(0.49309808015823364)),
		Fe:    int32(FE_INEXACT),
	},
	972: {
		Ffile: __ccgo_ts,
		Fline: int32(989),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5323901625375514),
		Fy:    -libc.Float64FromFloat64(0.6303786702097453),
		Fdy:   float32(0.49746939539909363),
		Fe:    int32(FE_INEXACT),
	},
	973: {
		Ffile: __ccgo_ts,
		Fline: int32(990),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5322084114281849),
		Fy:    -libc.Float64FromFloat64(0.6307201155499459),
		Fdy:   float32(-libc.Float64FromFloat64(0.4923783242702484)),
		Fe:    int32(FE_INEXACT),
	},
	974: {
		Ffile: __ccgo_ts,
		Fline: int32(991),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5320267223662954),
		Fy:    -libc.Float64FromFloat64(0.631061560890144),
		Fdy:   float32(-libc.Float64FromFloat64(0.4960213601589203)),
		Fe:    int32(FE_INEXACT),
	},
	975: {
		Ffile: __ccgo_ts,
		Fline: int32(992),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.528586391081725),
		Fy:    -libc.Float64FromFloat64(0.6375490223538707),
		Fdy:   float32(0.4904926121234894),
		Fe:    int32(FE_INEXACT),
	},
	976: {
		Ffile: __ccgo_ts,
		Fline: int32(993),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5278649504184717),
		Fy:    -libc.Float64FromFloat64(0.638914803714656),
		Fdy:   float32(-libc.Float64FromFloat64(0.49026548862457275)),
		Fe:    int32(FE_INEXACT),
	},
	977: {
		Ffile: __ccgo_ts,
		Fline: int32(994),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5276847441579418),
		Fy:    -libc.Float64FromFloat64(0.6392562490548548),
		Fdy:   float32(0.49356281757354736),
		Fe:    int32(FE_INEXACT),
	},
	978: {
		Ffile: __ccgo_ts,
		Fline: int32(995),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5275045994174986),
		Fy:    -libc.Float64FromFloat64(0.6395976943950493),
		Fdy:   float32(-libc.Float64FromFloat64(0.4914540648460388)),
		Fe:    int32(FE_INEXACT),
	},
	979: {
		Ffile: __ccgo_ts,
		Fline: int32(996),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5251683065732053),
		Fy:    -libc.Float64FromFloat64(0.644036483817605),
		Fdy:   float32(-libc.Float64FromFloat64(0.4963913559913635)),
		Fe:    int32(FE_INEXACT),
	},
	980: {
		Ffile: __ccgo_ts,
		Fline: int32(997),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.523735733876236),
		Fy:    -libc.Float64FromFloat64(0.6467680465392233),
		Fdy:   float32(0.4905742406845093),
		Fe:    int32(FE_INEXACT),
	},
	981: {
		Ffile: __ccgo_ts,
		Fline: int32(998),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5228423610590809),
		Fy:    -libc.Float64FromFloat64(0.6484752732401555),
		Fdy:   float32(-libc.Float64FromFloat64(0.49768775701522827)),
		Fe:    int32(FE_INEXACT),
	},
	982: {
		Ffile: __ccgo_ts,
		Fline: int32(999),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5224854387667215),
		Fy:    -libc.Float64FromFloat64(0.6491581639205516),
		Fdy:   float32(-libc.Float64FromFloat64(0.49436497688293457)),
		Fe:    int32(FE_INEXACT),
	},
	983: {
		Ffile: __ccgo_ts,
		Fline: int32(1000),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5203490159646328),
		Fy:    -libc.Float64FromFloat64(0.6532555080029052),
		Fdy:   float32(0.49539875984191895),
		Fe:    int32(FE_INEXACT),
	},
	984: {
		Ffile: __ccgo_ts,
		Fline: int32(1001),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5191068058425568),
		Fy:    -libc.Float64FromFloat64(0.6556456253842803),
		Fdy:   float32(0.4927680492401123),
		Fe:    int32(FE_INEXACT),
	},
	985: {
		Ffile: __ccgo_ts,
		Fline: int32(1002),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5176907679218755),
		Fy:    -libc.Float64FromFloat64(0.6583771881058493),
		Fdy:   float32(-libc.Float64FromFloat64(0.4951101541519165)),
		Fe:    int32(FE_INEXACT),
	},
	986: {
		Ffile: __ccgo_ts,
		Fline: int32(1003),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5169841981401506),
		Fy:    -libc.Float64FromFloat64(0.6597429694666334),
		Fdy:   float32(0.4994032680988312),
		Fe:    int32(FE_INEXACT),
	},
	987: {
		Ffile: __ccgo_ts,
		Fline: int32(1004),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5168077064275904),
		Fy:    -libc.Float64FromFloat64(0.6600844148068321),
		Fdy:   float32(-libc.Float64FromFloat64(0.4987243711948395)),
		Fe:    int32(FE_INEXACT),
	},
	988: {
		Ffile: __ccgo_ts,
		Fline: int32(1005),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.515221989887079),
		Fy:    -libc.Float64FromFloat64(0.6631574228686167),
		Fdy:   float32(-libc.Float64FromFloat64(0.4909955561161041)),
		Fe:    int32(FE_INEXACT),
	},
	989: {
		Ffile: __ccgo_ts,
		Fline: int32(1006),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5136411387867006),
		Fy:    -libc.Float64FromFloat64(0.6662304309303672),
		Fdy:   float32(-libc.Float64FromFloat64(0.49850842356681824)),
		Fe:    int32(FE_INEXACT),
	},
	990: {
		Ffile: __ccgo_ts,
		Fline: int32(1007),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5134657883513232),
		Fy:    -libc.Float64FromFloat64(0.6665718762705706),
		Fdy:   float32(0.49170950055122375),
		Fe:    int32(FE_INEXACT),
	},
	991: {
		Ffile: __ccgo_ts,
		Fline: int32(1008),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5103196972039252),
		Fy:    -libc.Float64FromFloat64(0.6727178923941117),
		Fdy:   float32(-libc.Float64FromFloat64(0.4919610321521759)),
		Fe:    int32(FE_INEXACT),
	},
	992: {
		Ffile: __ccgo_ts,
		Fline: int32(1009),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5092752896977203),
		Fy:    -libc.Float64FromFloat64(0.6747665644352819),
		Fdy:   float32(0.49584776163101196),
		Fe:    int32(FE_INEXACT),
	},
	993: {
		Ffile: __ccgo_ts,
		Fline: int32(1010),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5075393582672915),
		Fy:    -libc.Float64FromFloat64(0.6781810178372492),
		Fdy:   float32(-libc.Float64FromFloat64(0.4917784631252289)),
		Fe:    int32(FE_INEXACT),
	},
	994: {
		Ffile: __ccgo_ts,
		Fline: int32(1011),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5042573736244309),
		Fy:    -libc.Float64FromFloat64(0.6846684793009753),
		Fdy:   float32(0.4964941143989563),
		Fe:    int32(FE_INEXACT),
	},
	995: {
		Ffile: __ccgo_ts,
		Fline: int32(1012),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5021954846774855),
		Fy:    -libc.Float64FromFloat64(0.6887658233833307),
		Fdy:   float32(-libc.Float64FromFloat64(0.4970729649066925)),
		Fe:    int32(FE_INEXACT),
	},
	996: {
		Ffile: __ccgo_ts,
		Fline: int32(1013),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5013388545758872),
		Fy:    -libc.Float64FromFloat64(0.6904730500843093),
		Fdy:   float32(-libc.Float64FromFloat64(0.49277254939079285)),
		Fe:    int32(FE_INEXACT),
	},
	997: {
		Ffile: __ccgo_ts,
		Fline: int32(1014),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5009966118148667),
		Fy:    -libc.Float64FromFloat64(0.691155940764718),
		Fdy:   float32(-libc.Float64FromFloat64(0.4903624951839447)),
		Fe:    int32(FE_INEXACT),
	},
	998: {
		Ffile: __ccgo_ts,
		Fline: int32(1015),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5008255780573116),
		Fy:    -libc.Float64FromFloat64(0.6914973861049105),
		Fdy:   float32(-libc.Float64FromFloat64(0.49111950397491455)),
		Fe:    int32(FE_INEXACT),
	},
	999: {
		Ffile: __ccgo_ts,
		Fline: int32(1016),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4989480559971554),
		Fy:    -libc.Float64FromFloat64(0.6952532848470889),
		Fdy:   float32(0.490249365568161),
		Fe:    int32(FE_INEXACT),
	},
	1000: {
		Ffile: __ccgo_ts,
		Fline: int32(1017),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.49724732629345153),
		Fy:    -libc.Float64FromFloat64(0.6986677382490343),
		Fdy:   float32(0.4920831322669983),
		Fe:    int32(FE_INEXACT),
	},
	1001: {
		Ffile: __ccgo_ts,
		Fline: int32(1018),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4962296723740231),
		Fy:    -libc.Float64FromFloat64(0.7007164102902042),
		Fdy:   float32(0.49359583854675293),
		Fe:    int32(FE_INEXACT),
	},
	1002: {
		Ffile: __ccgo_ts,
		Fline: int32(1019),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4960602659878799),
		Fy:    -libc.Float64FromFloat64(0.7010578556304012),
		Fdy:   float32(0.4954407215118408),
		Fe:    int32(FE_INEXACT),
	},
	1003: {
		Ffile: __ccgo_ts,
		Fline: int32(1020),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4952141011596031),
		Fy:    -libc.Float64FromFloat64(0.7027650823313855),
		Fdy:   float32(0.4954765737056732),
		Fe:    int32(FE_INEXACT),
	},
	1004: {
		Ffile: __ccgo_ts,
		Fline: int32(1021),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.49436937969409483),
		Fy:    -libc.Float64FromFloat64(0.7044723090323606),
		Fdy:   float32(0.4999547600746155),
		Fe:    int32(FE_INEXACT),
	},
	1005: {
		Ffile: __ccgo_ts,
		Fline: int32(1022),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4928525104741696),
		Fy:    -libc.Float64FromFloat64(0.7075453170941329),
		Fdy:   float32(-libc.Float64FromFloat64(0.49709880352020264)),
		Fe:    int32(FE_INEXACT),
	},
	1006: {
		Ffile: __ccgo_ts,
		Fline: int32(1023),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4908372555726581),
		Fy:    -libc.Float64FromFloat64(0.7116426611764841),
		Fdy:   float32(-libc.Float64FromFloat64(0.49192336201667786)),
		Fe:    int32(FE_INEXACT),
	},
	1007: {
		Ffile: __ccgo_ts,
		Fline: int32(1024),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.49066969008780353),
		Fy:    -libc.Float64FromFloat64(0.7119841065166794),
		Fdy:   float32(-libc.Float64FromFloat64(0.4935777187347412)),
		Fe:    int32(FE_INEXACT),
	},
	1008: {
		Ffile: __ccgo_ts,
		Fline: int32(1025),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.699999999999973e+300),
		Fy:    libc.Float64FromFloat64(691.768779671224),
		Fdy:   float32(0.40125057101249695),
		Fe:    int32(FE_INEXACT),
	},
	1009: {
		Ffile: __ccgo_ts,
		Fline: int32(1026),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.855025408745231e+284),
		Fy:    libc.Float64FromFloat64(654.5520648037113),
		Fdy:   float32(0.00044815544970333576),
		Fe:    int32(FE_INEXACT),
	},
	1010: {
		Ffile: __ccgo_ts,
		Fline: int32(1027),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.674333751102181e+283),
		Fy:    libc.Float64FromFloat64(653.1736679567664),
		Fdy:   float32(-libc.Float64FromFloat64(0.00031353349913842976)),
		Fe:    int32(FE_INEXACT),
	},
	1011: {
		Ffile: __ccgo_ts,
		Fline: int32(1028),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1778488808665296e+283),
		Fy:    libc.Float64FromFloat64(651.7952711098216),
		Fdy:   float32(-libc.Float64FromFloat64(0.00030700548086315393)),
		Fe:    int32(FE_INEXACT),
	},
	1012: {
		Ffile: __ccgo_ts,
		Fline: int32(1029),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.597637387241978e+278),
		Fy:    libc.Float64FromFloat64(642.1464931812071),
		Fdy:   float32(0.0007162937545217574),
		Fe:    int32(FE_INEXACT),
	},
	1013: {
		Ffile: __ccgo_ts,
		Fline: int32(1030),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0630774924920024e+276),
		Fy:    libc.Float64FromFloat64(636.6329057934275),
		Fdy:   float32(-libc.Float64FromFloat64(0.000215521824429743)),
		Fe:    int32(FE_INEXACT),
	},
	1014: {
		Ffile: __ccgo_ts,
		Fline: int32(1031),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.111768366939115e+273),
		Fy:    libc.Float64FromFloat64(629.7409215587029),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004894029698334634)),
		Fe:    int32(FE_INEXACT),
	},
	1015: {
		Ffile: __ccgo_ts,
		Fline: int32(1032),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.841102248316935e+272),
		Fy:    libc.Float64FromFloat64(628.362524711758),
		Fdy:   float32(0.0004123070102650672),
		Fe:    int32(FE_INEXACT),
	},
	1016: {
		Ffile: __ccgo_ts,
		Fline: int32(1033),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2545462702162287e+271),
		Fy:    libc.Float64FromFloat64(624.2273341709233),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004549950244836509)),
		Fe:    int32(FE_INEXACT),
	},
	1017: {
		Ffile: __ccgo_ts,
		Fline: int32(1034),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0072258902774538e+269),
		Fy:    libc.Float64FromFloat64(620.0921436300886),
		Fdy:   float32(0.00015255789912771434),
		Fe:    int32(FE_INEXACT),
	},
	1018: {
		Ffile: __ccgo_ts,
		Fline: int32(1035),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.211484398981715e+267),
		Fy:    libc.Float64FromFloat64(615.9569530892538),
		Fdy:   float32(0.0002531104546505958),
		Fe:    int32(FE_INEXACT),
	},
	1019: {
		Ffile: __ccgo_ts,
		Fline: int32(1036),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.138251800586988e+265),
		Fy:    libc.Float64FromFloat64(611.8217625484191),
		Fdy:   float32(-libc.Float64FromFloat64(6.556255539180711e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1020: {
		Ffile: __ccgo_ts,
		Fline: int32(1037),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.221005705214308e+263),
		Fy:    libc.Float64FromFloat64(607.6865720075843),
		Fdy:   float32(-libc.Float64FromFloat64(0.00010529689461691305)),
		Fe:    int32(FE_INEXACT),
	},
	1021: {
		Ffile: __ccgo_ts,
		Fline: int32(1038),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.071547066402716e+263),
		Fy:    libc.Float64FromFloat64(606.3081751606394),
		Fdy:   float32(-libc.Float64FromFloat64(0.000536039297003299)),
		Fe:    int32(FE_INEXACT),
	},
	1022: {
		Ffile: __ccgo_ts,
		Fline: int32(1039),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3153293654749535e+262),
		Fy:    libc.Float64FromFloat64(603.5513814667496),
		Fdy:   float32(0.000208690034924075),
		Fe:    int32(FE_INEXACT),
	},
	1023: {
		Ffile: __ccgo_ts,
		Fline: int32(1040),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.351687334263114e+260),
		Fy:    libc.Float64FromFloat64(600.7945877728597),
		Fdy:   float32(-libc.Float64FromFloat64(0.00030435333610512316)),
		Fe:    int32(FE_INEXACT),
	},
	1024: {
		Ffile: __ccgo_ts,
		Fline: int32(1041),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1044765101955897e+260),
		Fy:    libc.Float64FromFloat64(599.4161909259149),
		Fdy:   float32(-libc.Float64FromFloat64(0.00030044038430787623)),
		Fe:    int32(FE_INEXACT),
	},
	1025: {
		Ffile: __ccgo_ts,
		Fline: int32(1042),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.387201445628428e+256),
		Fy:    libc.Float64FromFloat64(591.1458098442454),
		Fdy:   float32(0.000286802212940529),
		Fe:    int32(FE_INEXACT),
	},
	1026: {
		Ffile: __ccgo_ts,
		Fline: int32(1043),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.619315584065118e+254),
		Fy:    libc.Float64FromFloat64(587.0106193034106),
		Fdy:   float32(8.280808106064796e-05),
		Fe:    int32(FE_INEXACT),
	},
	1027: {
		Ffile: __ccgo_ts,
		Fline: int32(1044),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.171914064144765e+254),
		Fy:    libc.Float64FromFloat64(585.6322224564657),
		Fdy:   float32(-libc.Float64FromFloat64(0.0001266393082914874)),
		Fe:    int32(FE_INEXACT),
	},
	1028: {
		Ffile: __ccgo_ts,
		Fline: int32(1045),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2771438692056195e+245),
		Fy:    libc.Float64FromFloat64(564.956269752292),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004877792962361127)),
		Fe:    int32(FE_INEXACT),
	},
	1029: {
		Ffile: __ccgo_ts,
		Fline: int32(1046),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.73799711516787e+244),
		Fy:    libc.Float64FromFloat64(563.577872905347),
		Fdy:   float32(0.00039563412428833544),
		Fe:    int32(FE_INEXACT),
	},
	1030: {
		Ffile: __ccgo_ts,
		Fline: int32(1047),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5159260413299677e+235),
		Fy:    libc.Float64FromFloat64(541.5235233542285),
		Fdy:   float32(-libc.Float64FromFloat64(0.0003305026621092111)),
		Fe:    int32(FE_INEXACT),
	},
	1031: {
		Ffile: __ccgo_ts,
		Fline: int32(1048),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.54002329800433e+232),
		Fy:    libc.Float64FromFloat64(534.6315391195039),
		Fdy:   float32(-libc.Float64FromFloat64(0.0002671186812222004)),
		Fe:    int32(FE_INEXACT),
	},
	1032: {
		Ffile: __ccgo_ts,
		Fline: int32(1049),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.880584516745012e+231),
		Fy:    libc.Float64FromFloat64(533.253142272559),
		Fdy:   float32(0.0007054543239064515),
		Fe:    int32(FE_INEXACT),
	},
	1033: {
		Ffile: __ccgo_ts,
		Fline: int32(1050),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.778382061567214e+230),
		Fy:    libc.Float64FromFloat64(531.874745425614),
		Fdy:   float32(-libc.Float64FromFloat64(8.363384404219687e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1034: {
		Ffile: __ccgo_ts,
		Fline: int32(1051),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.463978179817935e+230),
		Fy:    libc.Float64FromFloat64(530.4963485786692),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006807362078689039)),
		Fe:    int32(FE_INEXACT),
	},
	1035: {
		Ffile: __ccgo_ts,
		Fline: int32(1052),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.564503606201741e+229),
		Fy:    libc.Float64FromFloat64(527.7395548847793),
		Fdy:   float32(-libc.Float64FromFloat64(8.178473945008591e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1036: {
		Ffile: __ccgo_ts,
		Fline: int32(1053),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5031457335243e+227),
		Fy:    libc.Float64FromFloat64(523.6043643439446),
		Fdy:   float32(0.0007151241879910231),
		Fe:    int32(FE_INEXACT),
	},
	1037: {
		Ffile: __ccgo_ts,
		Fline: int32(1054),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.004937117705146e+225),
		Fy:    libc.Float64FromFloat64(519.4691738031098),
		Fdy:   float32(-libc.Float64FromFloat64(0.0001418407482560724)),
		Fe:    int32(FE_INEXACT),
	},
	1038: {
		Ffile: __ccgo_ts,
		Fline: int32(1055),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.58335856775157e+221),
		Fy:    libc.Float64FromFloat64(509.82039587449543),
		Fdy:   float32(0.0011012349277734756),
		Fe:    int32(FE_INEXACT),
	},
	1039: {
		Ffile: __ccgo_ts,
		Fline: int32(1056),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.509603635353377e+220),
		Fy:    libc.Float64FromFloat64(508.44199902755054),
		Fdy:   float32(0.0004427394596859813),
		Fe:    int32(FE_INEXACT),
	},
	1040: {
		Ffile: __ccgo_ts,
		Fline: int32(1057),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6663785955769362e+217),
		Fy:    libc.Float64FromFloat64(500.17161794588105),
		Fdy:   float32(0.00022261348203755915),
		Fe:    int32(FE_INEXACT),
	},
	1041: {
		Ffile: __ccgo_ts,
		Fline: int32(1058),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0580675362434805e+216),
		Fy:    libc.Float64FromFloat64(497.4148242519912),
		Fdy:   float32(-libc.Float64FromFloat64(0.0010811034590005875)),
		Fe:    int32(FE_INEXACT),
	},
	1042: {
		Ffile: __ccgo_ts,
		Fline: int32(1059),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.718202659491164e+214),
		Fy:    libc.Float64FromFloat64(494.6580305581014),
		Fdy:   float32(0.000537406129296869),
		Fe:    int32(FE_INEXACT),
	},
	1043: {
		Ffile: __ccgo_ts,
		Fline: int32(1060),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.692867455613212e+214),
		Fy:    libc.Float64FromFloat64(493.27963371115646),
		Fdy:   float32(0.00025838054716587067),
		Fe:    int32(FE_INEXACT),
	},
	1044: {
		Ffile: __ccgo_ts,
		Fline: int32(1061),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.2657245807044e+213),
		Fy:    libc.Float64FromFloat64(491.90123686421157),
		Fdy:   float32(0.00018303742399439216),
		Fe:    int32(FE_INEXACT),
	},
	1045: {
		Ffile: __ccgo_ts,
		Fline: int32(1062),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.333532809681597e+210),
		Fy:    libc.Float64FromFloat64(485.009252629487),
		Fdy:   float32(0.00016897985187824816),
		Fe:    int32(FE_INEXACT),
	},
	1046: {
		Ffile: __ccgo_ts,
		Fline: int32(1063),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0919731114360445e+210),
		Fy:    libc.Float64FromFloat64(483.6308557825421),
		Fdy:   float32(-libc.Float64FromFloat64(0.0002381162194069475)),
		Fe:    int32(FE_INEXACT),
	},
	1047: {
		Ffile: __ccgo_ts,
		Fline: int32(1064),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.933486200122759e+208),
		Fy:    libc.Float64FromFloat64(480.87406208865224),
		Fdy:   float32(0.0001982752582989633),
		Fe:    int32(FE_INEXACT),
	},
	1048: {
		Ffile: __ccgo_ts,
		Fline: int32(1065),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7471150748258009e+208),
		Fy:    libc.Float64FromFloat64(479.49566524170734),
		Fdy:   float32(0.00026375919696874917),
		Fe:    int32(FE_INEXACT),
	},
	1049: {
		Ffile: __ccgo_ts,
		Fline: int32(1066),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.402418922575371e+207),
		Fy:    libc.Float64FromFloat64(478.11726839476245),
		Fdy:   float32(0.00043132659629918635),
		Fe:    int32(FE_INEXACT),
	},
	1050: {
		Ffile: __ccgo_ts,
		Fline: int32(1067),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.774887327066178e+205),
		Fy:    libc.Float64FromFloat64(472.60368100698275),
		Fdy:   float32(0.0009138844325207174),
		Fe:    int32(FE_INEXACT),
	},
	1051: {
		Ffile: __ccgo_ts,
		Fline: int32(1068),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.144879494586778e+201),
		Fy:    libc.Float64FromFloat64(462.9549030783684),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004165813443250954)),
		Fe:    int32(FE_INEXACT),
	},
	1052: {
		Ffile: __ccgo_ts,
		Fline: int32(1069),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9773383228113497e+194),
		Fy:    libc.Float64FromFloat64(447.79253776197436),
		Fdy:   float32(0.0017525587463751435),
		Fe:    int32(FE_INEXACT),
	},
	1053: {
		Ffile: __ccgo_ts,
		Fline: int32(1070),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0246662819695654e+191),
		Fy:    libc.Float64FromFloat64(440.90055352724977),
		Fdy:   float32(0.0005647700163535774),
		Fe:    int32(FE_INEXACT),
	},
	1054: {
		Ffile: __ccgo_ts,
		Fline: int32(1071),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9205126670594646e+190),
		Fy:    libc.Float64FromFloat64(438.1437598333599),
		Fdy:   float32(-libc.Float64FromFloat64(0.0014597040135413408)),
		Fe:    int32(FE_INEXACT),
	},
	1055: {
		Ffile: __ccgo_ts,
		Fline: int32(1072),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.742774266107443e+187),
		Fy:    libc.Float64FromFloat64(432.6301724455803),
		Fdy:   float32(-libc.Float64FromFloat64(0.001506447559222579)),
		Fe:    int32(FE_INEXACT),
	},
	1056: {
		Ffile: __ccgo_ts,
		Fline: int32(1073),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.951041258443708e+187),
		Fy:    libc.Float64FromFloat64(431.2517755986354),
		Fdy:   float32(-libc.Float64FromFloat64(0.0013384955236688256)),
		Fe:    int32(FE_INEXACT),
	},
	1057: {
		Ffile: __ccgo_ts,
		Fline: int32(1074),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2388141703873702e+186),
		Fy:    libc.Float64FromFloat64(428.49498190474554),
		Fdy:   float32(0.00038424646481871605),
		Fe:    int32(FE_INEXACT),
	},
	1058: {
		Ffile: __ccgo_ts,
		Fline: int32(1075),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.982055134256258e+184),
		Fy:    libc.Float64FromFloat64(424.3597913639108),
		Fdy:   float32(-libc.Float64FromFloat64(0.00033982854802161455)),
		Fe:    int32(FE_INEXACT),
	},
	1059: {
		Ffile: __ccgo_ts,
		Fline: int32(1076),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1712121552525986e+182),
		Fy:    libc.Float64FromFloat64(420.22460082307606),
		Fdy:   float32(0.0003078729205299169),
		Fe:    int32(FE_INEXACT),
	},
	1060: {
		Ffile: __ccgo_ts,
		Fline: int32(1077),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.2216219451701444e+179),
		Fy:    libc.Float64FromFloat64(413.3326165883515),
		Fdy:   float32(-libc.Float64FromFloat64(0.0015292088501155376)),
		Fe:    int32(FE_INEXACT),
	},
	1061: {
		Ffile: __ccgo_ts,
		Fline: int32(1078),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2988350412903906e+177),
		Fy:    libc.Float64FromFloat64(407.81902920057183),
		Fdy:   float32(-libc.Float64FromFloat64(0.0011345547391101718)),
		Fe:    int32(FE_INEXACT),
	},
	1062: {
		Ffile: __ccgo_ts,
		Fline: int32(1079),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3194814056342417e+174),
		Fy:    libc.Float64FromFloat64(400.9270449658473),
		Fdy:   float32(0.00028139923233538866),
		Fe:    int32(FE_INEXACT),
	},
	1063: {
		Ffile: __ccgo_ts,
		Fline: int32(1080),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.111119615119698e+172),
		Fy:    libc.Float64FromFloat64(396.79185442501256),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004297379346098751)),
		Fe:    int32(FE_INEXACT),
	},
	1064: {
		Ffile: __ccgo_ts,
		Fline: int32(1081),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.14467810666128e+169),
		Fy:    libc.Float64FromFloat64(389.89987019028797),
		Fdy:   float32(0.000526917923707515),
		Fe:    int32(FE_INEXACT),
	},
	1065: {
		Ffile: __ccgo_ts,
		Fline: int32(1082),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3617639391033954e+168),
		Fy:    libc.Float64FromFloat64(387.1430764963981),
		Fdy:   float32(0.0010329196229577065),
		Fe:    int32(FE_INEXACT),
	},
	1066: {
		Ffile: __ccgo_ts,
		Fline: int32(1083),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.577383774666566e+162),
		Fy:    libc.Float64FromFloat64(374.7375048738939),
		Fdy:   float32(0.0005877484218217432),
		Fe:    int32(FE_INEXACT),
	},
	1067: {
		Ffile: __ccgo_ts,
		Fline: int32(1084),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2485883224511097e+160),
		Fy:    libc.Float64FromFloat64(369.22391748611426),
		Fdy:   float32(-libc.Float64FromFloat64(0.0010295957326889038)),
		Fe:    int32(FE_INEXACT),
	},
	1068: {
		Ffile: __ccgo_ts,
		Fline: int32(1085),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2843320252992648e+157),
		Fy:    libc.Float64FromFloat64(362.3319332513897),
		Fdy:   float32(-libc.Float64FromFloat64(0.0011092235799878836)),
		Fe:    int32(FE_INEXACT),
	},
	1069: {
		Ffile: __ccgo_ts,
		Fline: int32(1086),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3575330154018488e+151),
		Fy:    libc.Float64FromFloat64(348.54796478194055),
		Fdy:   float32(6.027381095918827e-05),
		Fe:    int32(FE_INEXACT),
	},
	1070: {
		Ffile: __ccgo_ts,
		Fline: int32(1087),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3950085100114543e+148),
		Fy:    libc.Float64FromFloat64(341.655980547216),
		Fdy:   float32(0.0007253738003782928),
		Fe:    int32(FE_INEXACT),
	},
	1071: {
		Ffile: __ccgo_ts,
		Fline: int32(1088),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.228385305206043e+141),
		Fy:    libc.Float64FromFloat64(326.49361523082194),
		Fdy:   float32(0.0004822110931854695),
		Fe:    int32(FE_INEXACT),
	},
	1072: {
		Ffile: __ccgo_ts,
		Fline: int32(1089),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.530152203966285e+132),
		Fy:    libc.Float64FromFloat64(305.81766252664823),
		Fdy:   float32(0.0010940146166831255),
		Fe:    int32(FE_INEXACT),
	},
	1073: {
		Ffile: __ccgo_ts,
		Fline: int32(1090),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.146321893527306e+131),
		Fy:    libc.Float64FromFloat64(303.06086883275844),
		Fdy:   float32(-libc.Float64FromFloat64(0.0008326743263751268)),
		Fe:    int32(FE_INEXACT),
	},
	1074: {
		Ffile: __ccgo_ts,
		Fline: int32(1091),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.044799293771313e+131),
		Fy:    libc.Float64FromFloat64(301.6824719858135),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006173266447149217)),
		Fe:    int32(FE_INEXACT),
	},
	1075: {
		Ffile: __ccgo_ts,
		Fline: int32(1092),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.632708198485911e+130),
		Fy:    libc.Float64FromFloat64(300.3040751388686),
		Fdy:   float32(0.0008095950470305979),
		Fe:    int32(FE_INEXACT),
	},
	1076: {
		Ffile: __ccgo_ts,
		Fline: int32(1093),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.633955918323995e+129),
		Fy:    libc.Float64FromFloat64(298.9256782919237),
		Fdy:   float32(3.091900725848973e-05),
		Fe:    int32(FE_INEXACT),
	},
	1077: {
		Ffile: __ccgo_ts,
		Fline: int32(1094),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.674557895290199e+127),
		Fy:    libc.Float64FromFloat64(293.412090904144),
		Fdy:   float32(-libc.Float64FromFloat64(0.0005964654847048223)),
		Fe:    int32(FE_INEXACT),
	},
	1078: {
		Ffile: __ccgo_ts,
		Fline: int32(1095),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0782796906294104e+125),
		Fy:    libc.Float64FromFloat64(287.89850351636437),
		Fdy:   float32(0.0015082211466506124),
		Fe:    int32(FE_INEXACT),
	},
	1079: {
		Ffile: __ccgo_ts,
		Fline: int32(1096),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7526301132214923e+120),
		Fy:    libc.Float64FromFloat64(276.87132874080504),
		Fdy:   float32(-libc.Float64FromFloat64(0.0013665660517290235)),
		Fe:    int32(FE_INEXACT),
	},
	1080: {
		Ffile: __ccgo_ts,
		Fline: int32(1097),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.112832960507674e+119),
		Fy:    libc.Float64FromFloat64(274.11453504691525),
		Fdy:   float32(-libc.Float64FromFloat64(0.0010270075872540474)),
		Fe:    int32(FE_INEXACT),
	},
	1081: {
		Ffile: __ccgo_ts,
		Fline: int32(1098),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.804140925572836e+118),
		Fy:    libc.Float64FromFloat64(272.7361381999703),
		Fdy:   float32(-libc.Float64FromFloat64(0.0011487321462482214)),
		Fe:    int32(FE_INEXACT),
	},
	1082: {
		Ffile: __ccgo_ts,
		Fline: int32(1099),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1667500088074774e+110),
		Fy:    libc.Float64FromFloat64(253.43858234274154),
		Fdy:   float32(-libc.Float64FromFloat64(0.0014848082792013884)),
		Fe:    int32(FE_INEXACT),
	},
	1083: {
		Ffile: __ccgo_ts,
		Fline: int32(1100),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.40828231056086e+108),
		Fy:    libc.Float64FromFloat64(250.6817886488517),
		Fdy:   float32(-libc.Float64FromFloat64(0.0003088162047788501)),
		Fe:    int32(FE_INEXACT),
	},
	1084: {
		Ffile: __ccgo_ts,
		Fline: int32(1101),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9867367507705986e+106),
		Fy:    libc.Float64FromFloat64(245.16820126107206),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006259044166654348)),
		Fe:    int32(FE_INEXACT),
	},
	1085: {
		Ffile: __ccgo_ts,
		Fline: int32(1102),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.778664187935149e+104),
		Fy:    libc.Float64FromFloat64(241.03301072023731),
		Fdy:   float32(0.0015685339458286762),
		Fe:    int32(FE_INEXACT),
	},
	1086: {
		Ffile: __ccgo_ts,
		Fline: int32(1103),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.204138293391806e+104),
		Fy:    libc.Float64FromFloat64(239.6546138732924),
		Fdy:   float32(-libc.Float64FromFloat64(0.0017646930646151304)),
		Fe:    int32(FE_INEXACT),
	},
	1087: {
		Ffile: __ccgo_ts,
		Fline: int32(1104),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0824461681893475e+100),
		Fy:    libc.Float64FromFloat64(231.3842327916229),
		Fdy:   float32(-libc.Float64FromFloat64(0.0018820523982867599)),
		Fe:    int32(FE_INEXACT),
	},
	1088: {
		Ffile: __ccgo_ts,
		Fline: int32(1105),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.931795583043576e+98),
		Fy:    libc.Float64FromFloat64(227.24904225078816),
		Fdy:   float32(-libc.Float64FromFloat64(0.001083637005649507)),
		Fe:    int32(FE_INEXACT),
	},
	1089: {
		Ffile: __ccgo_ts,
		Fline: int32(1106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.231791492087817e+91),
		Fy:    libc.Float64FromFloat64(210.70828008744922),
		Fdy:   float32(5.803331077913754e-05),
		Fe:    int32(FE_INEXACT),
	},
	1090: {
		Ffile: __ccgo_ts,
		Fline: int32(1107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.143539154113069e+90),
		Fy:    libc.Float64FromFloat64(209.3298832405043),
		Fdy:   float32(0.0015903724124655128),
		Fe:    int32(FE_INEXACT),
	},
	1091: {
		Ffile: __ccgo_ts,
		Fline: int32(1108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0520268747823837e+90),
		Fy:    libc.Float64FromFloat64(207.95148639355938),
		Fdy:   float32(-libc.Float64FromFloat64(0.0011568719055503607)),
		Fe:    int32(FE_INEXACT),
	},
	1092: {
		Ffile: __ccgo_ts,
		Fline: int32(1109),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4969523998511605e+76),
		Fy:    libc.Float64FromFloat64(176.24835891382637),
		Fdy:   float32(0.0008024497656151652),
		Fe:    int32(FE_INEXACT),
	},
	1093: {
		Ffile: __ccgo_ts,
		Fline: int32(1110),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.594989647404178e+74),
		Fy:    libc.Float64FromFloat64(172.11316837299165),
		Fdy:   float32(-libc.Float64FromFloat64(0.001523628132417798)),
		Fe:    int32(FE_INEXACT),
	},
	1094: {
		Ffile: __ccgo_ts,
		Fline: int32(1111),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.951768733223654e+72),
		Fy:    libc.Float64FromFloat64(167.9779778321569),
		Fdy:   float32(-libc.Float64FromFloat64(0.00018457486294209957)),
		Fe:    int32(FE_INEXACT),
	},
	1095: {
		Ffile: __ccgo_ts,
		Fline: int32(1112),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.609011581781005e+70),
		Fy:    libc.Float64FromFloat64(162.46439044437724),
		Fdy:   float32(0.0028411620296537876),
		Fe:    int32(FE_INEXACT),
	},
	1096: {
		Ffile: __ccgo_ts,
		Fline: int32(1113),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.094066617798543e+69),
		Fy:    libc.Float64FromFloat64(161.08599359743232),
		Fdy:   float32(-libc.Float64FromFloat64(0.0013529036659747362)),
		Fe:    int32(FE_INEXACT),
	},
	1097: {
		Ffile: __ccgo_ts,
		Fline: int32(1114),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.72466167911917e+64),
		Fy:    libc.Float64FromFloat64(148.68042197492812),
		Fdy:   float32(0.0007249926566146314),
		Fe:    int32(FE_INEXACT),
	},
	1098: {
		Ffile: __ccgo_ts,
		Fline: int32(1115),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.959315756096997e+62),
		Fy:    libc.Float64FromFloat64(144.54523143409338),
		Fdy:   float32(0.0016438864404335618),
		Fe:    int32(FE_INEXACT),
	},
	1099: {
		Ffile: __ccgo_ts,
		Fline: int32(1116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.054045418008295e+59),
		Fy:    libc.Float64FromFloat64(137.6532471993688),
		Fdy:   float32(0.0010570957092568278),
		Fe:    int32(FE_INEXACT),
	},
	1100: {
		Ffile: __ccgo_ts,
		Fline: int32(1117),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5255116557807657e+59),
		Fy:    libc.Float64FromFloat64(136.2748503524239),
		Fdy:   float32(0.00033413487835787237),
		Fe:    int32(FE_INEXACT),
	},
	1101: {
		Ffile: __ccgo_ts,
		Fline: int32(1118),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5497612859914164e+56),
		Fy:    libc.Float64FromFloat64(129.3828661176993),
		Fdy:   float32(-libc.Float64FromFloat64(0.001488643465563655)),
		Fe:    int32(FE_INEXACT),
	},
	1102: {
		Ffile: __ccgo_ts,
		Fline: int32(1119),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4795585869109036e+54),
		Fy:    libc.Float64FromFloat64(125.24767557686458),
		Fdy:   float32(0.0010112300515174866),
		Fe:    int32(FE_INEXACT),
	},
	1103: {
		Ffile: __ccgo_ts,
		Fline: int32(1120),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.448263628942141e+47),
		Fy:    libc.Float64FromFloat64(110.08531026047054),
		Fdy:   float32(0.00041983870323747396),
		Fe:    int32(FE_INEXACT),
	},
	1104: {
		Ffile: __ccgo_ts,
		Fline: int32(1121),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.094326720850124e+46),
		Fy:    libc.Float64FromFloat64(107.3285165665807),
		Fdy:   float32(0.004319945350289345),
		Fe:    int32(FE_INEXACT),
	},
	1105: {
		Ffile: __ccgo_ts,
		Fline: int32(1122),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0480973640037883e+43),
		Fy:    libc.Float64FromFloat64(99.05813548491122),
		Fdy:   float32(0.0026578253600746393),
		Fe:    int32(FE_INEXACT),
	},
	1106: {
		Ffile: __ccgo_ts,
		Fline: int32(1123),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.641018748265132e+42),
		Fy:    libc.Float64FromFloat64(97.67973863796631),
		Fdy:   float32(-libc.Float64FromFloat64(0.0031704579014331102)),
		Fe:    int32(FE_INEXACT),
	},
	1107: {
		Ffile: __ccgo_ts,
		Fline: int32(1124),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.6548970241115044e+41),
		Fy:    libc.Float64FromFloat64(96.3013417910214),
		Fdy:   float32(0.0022679148241877556),
		Fe:    int32(FE_INEXACT),
	},
	1108: {
		Ffile: __ccgo_ts,
		Fline: int32(1125),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.292697922456709e+37),
		Fy:    libc.Float64FromFloat64(86.652563862407),
		Fdy:   float32(-libc.Float64FromFloat64(0.007486560847610235)),
		Fe:    int32(FE_INEXACT),
	},
	1109: {
		Ffile: __ccgo_ts,
		Fline: int32(1126),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.977328527322507e+32),
		Fy:    libc.Float64FromFloat64(75.62538908684769),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006277129868976772)),
		Fe:    int32(FE_INEXACT),
	},
	1110: {
		Ffile: __ccgo_ts,
		Fline: int32(1127),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.430256619990795e+31),
		Fy:    libc.Float64FromFloat64(72.86859539295787),
		Fdy:   float32(0.005739106796681881),
		Fe:    int32(FE_INEXACT),
	},
	1111: {
		Ffile: __ccgo_ts,
		Fline: int32(1128),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1521188803671966e+25),
		Fy:    libc.Float64FromFloat64(57.70623007656382),
		Fdy:   float32(0.011879983358085155),
		Fe:    int32(FE_INEXACT),
	},
	1112: {
		Ffile: __ccgo_ts,
		Fline: int32(1129),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.315381921283572e+23),
		Fy:    libc.Float64FromFloat64(54.949436382673994),
		Fdy:   float32(0.003798835910856724),
		Fe:    int32(FE_INEXACT),
	},
	1113: {
		Ffile: __ccgo_ts,
		Fline: int32(1130),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1890383131925658e+19),
		Fy:    libc.Float64FromFloat64(43.922261607114685),
		Fdy:   float32(-libc.Float64FromFloat64(0.009014979004859924)),
		Fe:    int32(FE_INEXACT),
	},
	1114: {
		Ffile: __ccgo_ts,
		Fline: int32(1131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9961648463183626e+18),
		Fy:    libc.Float64FromFloat64(42.54386476016977),
		Fdy:   float32(-libc.Float64FromFloat64(0.01178494282066822)),
		Fe:    int32(FE_INEXACT),
	},
	1115: {
		Ffile: __ccgo_ts,
		Fline: int32(1132),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.963378225268694e+11),
		Fy:    libc.Float64FromFloat64(26.003102596830807),
		Fdy:   float32(-libc.Float64FromFloat64(0.00440884567797184)),
		Fe:    int32(FE_INEXACT),
	},
	1116: {
		Ffile: __ccgo_ts,
		Fline: int32(1133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1912645959609547e+06),
		Fy:    libc.Float64FromFloat64(14.975927821271501),
		Fdy:   float32(-libc.Float64FromFloat64(0.011949931271374226)),
		Fe:    int32(FE_INEXACT),
	},
	1117: {
		Ffile: __ccgo_ts,
		Fline: int32(1134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(804141.8591505148),
		Fy:    libc.Float64FromFloat64(13.597530974326588),
		Fdy:   float32(-libc.Float64FromFloat64(0.03107423335313797)),
		Fe:    int32(FE_INEXACT),
	},
	1118: {
		Ffile: __ccgo_ts,
		Fline: int32(1135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(202629.43112159194),
		Fy:    libc.Float64FromFloat64(12.219134127381674),
		Fdy:   float32(0.0014447752619162202),
		Fe:    int32(FE_INEXACT),
	},
	1119: {
		Ffile: __ccgo_ts,
		Fline: int32(1136),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(12865.961164306518),
		Fy:    libc.Float64FromFloat64(9.462340433491846),
		Fdy:   float32(0.005155033897608519),
		Fe:    int32(FE_INEXACT),
	},
	1120: {
		Ffile: __ccgo_ts,
		Fline: int32(1137),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(816.9245492384185),
		Fy:    libc.Float64FromFloat64(6.70554673960202),
		Fdy:   float32(0.07645606994628906),
		Fe:    int32(FE_INEXACT),
	},
	1121: {
		Ffile: __ccgo_ts,
		Fline: int32(1138),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.565047575371527e-07),
		Fy:    -libc.Float64FromFloat64(13.970405964571682),
		Fdy:   float32(-libc.Float64FromFloat64(0.006799963302910328)),
		Fe:    int32(FE_INEXACT),
	},
	1122: {
		Ffile: __ccgo_ts,
		Fline: int32(1139),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.158239492171566e-07),
		Fy:    -libc.Float64FromFloat64(15.348802811516595),
		Fdy:   float32(-libc.Float64FromFloat64(0.010654755868017673)),
		Fe:    int32(FE_INEXACT),
	},
	1123: {
		Ffile: __ccgo_ts,
		Fline: int32(1140),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4367697596335356e-17),
		Fy:    -libc.Float64FromFloat64(38.78154920958013),
		Fdy:   float32(0.004450969863682985),
		Fe:    int32(FE_INEXACT),
	},
	1124: {
		Ffile: __ccgo_ts,
		Fline: int32(1141),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6204039838791536e-18),
		Fy:    -libc.Float64FromFloat64(40.15994605652504),
		Fdy:   float32(-libc.Float64FromFloat64(0.004730952437967062)),
		Fe:    int32(FE_INEXACT),
	},
	1125: {
		Ffile: __ccgo_ts,
		Fline: int32(1142),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.122773442719956e-19),
		Fy:    -libc.Float64FromFloat64(41.53834290346995),
		Fdy:   float32(-libc.Float64FromFloat64(0.006756927818059921)),
		Fe:    int32(FE_INEXACT),
	},
	1126: {
		Ffile: __ccgo_ts,
		Fline: int32(1143),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.677954155645899e-21),
		Fy:    -libc.Float64FromFloat64(47.051930291249604),
		Fdy:   float32(-libc.Float64FromFloat64(0.006255533546209335)),
		Fe:    int32(FE_INEXACT),
	},
	1127: {
		Ffile: __ccgo_ts,
		Fline: int32(1144),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.978127256083782e-26),
		Fy:    -libc.Float64FromFloat64(58.079105066808914),
		Fdy:   float32(0.004611059091985226),
		Fe:    int32(FE_INEXACT),
	},
	1128: {
		Ffile: __ccgo_ts,
		Fline: int32(1145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.56477420455485e-28),
		Fy:    -libc.Float64FromFloat64(62.214295607643656),
		Fdy:   float32(-libc.Float64FromFloat64(0.004013489466160536)),
		Fe:    int32(FE_INEXACT),
	},
	1129: {
		Ffile: __ccgo_ts,
		Fline: int32(1146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.41015280304252e-28),
		Fy:    -libc.Float64FromFloat64(63.592692454588565),
		Fdy:   float32(0.0035098388325423002),
		Fe:    int32(FE_INEXACT),
	},
	1130: {
		Ffile: __ccgo_ts,
		Fline: int32(1147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.1072510950916696e-42),
		Fy:    -libc.Float64FromFloat64(95.29581993432159),
		Fdy:   float32(0.0050625004805624485),
		Fe:    int32(FE_INEXACT),
	},
	1131: {
		Ffile: __ccgo_ts,
		Fline: int32(1148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6079001845391032e-43),
		Fy:    -libc.Float64FromFloat64(98.0526136282114),
		Fdy:   float32(-libc.Float64FromFloat64(0.00253826892003417)),
		Fe:    int32(FE_INEXACT),
	},
	1132: {
		Ffile: __ccgo_ts,
		Fline: int32(1149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.649355531577439e-46),
		Fy:    -libc.Float64FromFloat64(104.94459786293598),
		Fdy:   float32(0.001327008823864162),
		Fe:    int32(FE_INEXACT),
	},
	1133: {
		Ffile: __ccgo_ts,
		Fline: int32(1150),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.238867183922625e-48),
		Fy:    -libc.Float64FromFloat64(109.0797884037707),
		Fdy:   float32(0.004307890776544809),
		Fe:    int32(FE_INEXACT),
	},
	1134: {
		Ffile: __ccgo_ts,
		Fline: int32(1151),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.068119059223667e-48),
		Fy:    -libc.Float64FromFloat64(110.45818525071563),
		Fdy:   float32(-libc.Float64FromFloat64(0.004236247856169939)),
		Fe:    int32(FE_INEXACT),
	},
	1135: {
		Ffile: __ccgo_ts,
		Fline: int32(1152),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0850979476569464e-51),
		Fy:    -libc.Float64FromFloat64(117.35016948544019),
		Fdy:   float32(-libc.Float64FromFloat64(0.0027454225346446037)),
		Fe:    int32(FE_INEXACT),
	},
	1136: {
		Ffile: __ccgo_ts,
		Fline: int32(1153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7637124721771683e-56),
		Fy:    -libc.Float64FromFloat64(128.3773442609995),
		Fdy:   float32(-libc.Float64FromFloat64(0.00027726709959097207)),
		Fe:    int32(FE_INEXACT),
	},
	1137: {
		Ffile: __ccgo_ts,
		Fline: int32(1154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.5148873526284527e-60),
		Fy:    -libc.Float64FromFloat64(136.647725342669),
		Fdy:   float32(-libc.Float64FromFloat64(0.0015743798576295376)),
		Fe:    int32(FE_INEXACT),
	},
	1138: {
		Ffile: __ccgo_ts,
		Fline: int32(1155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1376712273225353e-60),
		Fy:    -libc.Float64FromFloat64(138.0261221896139),
		Fdy:   float32(0.0006056312122382224),
		Fe:    int32(FE_INEXACT),
	},
	1139: {
		Ffile: __ccgo_ts,
		Fline: int32(1156),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1741276878234548e-66),
		Fy:    -libc.Float64FromFloat64(151.81009065906304),
		Fdy:   float32(-libc.Float64FromFloat64(0.0018615764565765858)),
		Fe:    int32(FE_INEXACT),
	},
	1140: {
		Ffile: __ccgo_ts,
		Fline: int32(1157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.053400272569548e-73),
		Fy:    -libc.Float64FromFloat64(166.97245597545708),
		Fdy:   float32(-libc.Float64FromFloat64(0.0003235733893234283)),
		Fe:    int32(FE_INEXACT),
	},
	1141: {
		Ffile: __ccgo_ts,
		Fline: int32(1158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.962980673688199e-78),
		Fy:    -libc.Float64FromFloat64(177.99963075101638),
		Fdy:   float32(-libc.Float64FromFloat64(0.0019678790122270584)),
		Fe:    int32(FE_INEXACT),
	},
	1142: {
		Ffile: __ccgo_ts,
		Fline: int32(1159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2505827661293958e-78),
		Fy:    -libc.Float64FromFloat64(179.3780275979613),
		Fdy:   float32(0.001930116442963481),
		Fe:    int32(FE_INEXACT),
	},
	1143: {
		Ffile: __ccgo_ts,
		Fline: int32(1160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.252226992490821e-85),
		Fy:    -libc.Float64FromFloat64(194.54039291435535),
		Fdy:   float32(0.002585108857601881),
		Fe:    int32(FE_INEXACT),
	},
	1144: {
		Ffile: __ccgo_ts,
		Fline: int32(1161),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3111738285773845e-87),
		Fy:    -libc.Float64FromFloat64(200.05398030213502),
		Fdy:   float32(0.00037346474709920585),
		Fe:    int32(FE_INEXACT),
	},
	1145: {
		Ffile: __ccgo_ts,
		Fline: int32(1162),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.286152573961755e-90),
		Fy:    -libc.Float64FromFloat64(205.56756768991465),
		Fdy:   float32(-libc.Float64FromFloat64(0.0007958468049764633)),
		Fe:    int32(FE_INEXACT),
	},
	1146: {
		Ffile: __ccgo_ts,
		Fline: int32(1163),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.332016331067914e-90),
		Fy:    -libc.Float64FromFloat64(206.94596453685958),
		Fdy:   float32(0.0008837159257382154),
		Fe:    int32(FE_INEXACT),
	},
	1147: {
		Ffile: __ccgo_ts,
		Fline: int32(1164),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.457641267056697e-92),
		Fy:    -libc.Float64FromFloat64(209.7027582307494),
		Fdy:   float32(0.0014609688660129905),
		Fe:    int32(FE_INEXACT),
	},
	1148: {
		Ffile: __ccgo_ts,
		Fline: int32(1165),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.1650523082464757e-95),
		Fy:    -libc.Float64FromFloat64(217.97313931241888),
		Fdy:   float32(0.0018417028477415442),
		Fe:    int32(FE_INEXACT),
	},
	1149: {
		Ffile: __ccgo_ts,
		Fline: int32(1166),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.464000611287356e-97),
		Fy:    -libc.Float64FromFloat64(222.10832985325362),
		Fdy:   float32(0.0014844238758087158),
		Fe:    int32(FE_INEXACT),
	},
	1150: {
		Ffile: __ccgo_ts,
		Fline: int32(1167),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.199468117304924e-98),
		Fy:    -libc.Float64FromFloat64(224.86512354714344),
		Fdy:   float32(0.0014113979414105415),
		Fe:    int32(FE_INEXACT),
	},
	1151: {
		Ffile: __ccgo_ts,
		Fline: int32(1168),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.71986898744926e-105),
		Fy:    -libc.Float64FromFloat64(240.02748886353749),
		Fdy:   float32(0.0013273241929709911),
		Fe:    int32(FE_INEXACT),
	},
	1152: {
		Ffile: __ccgo_ts,
		Fline: int32(1169),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.151570885165047e-107),
		Fy:    -libc.Float64FromFloat64(244.16267940437223),
		Fdy:   float32(-libc.Float64FromFloat64(0.0010737342527136207)),
		Fe:    int32(FE_INEXACT),
	},
	1153: {
		Ffile: __ccgo_ts,
		Fline: int32(1170),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.487491497429517e-111),
		Fy:    -libc.Float64FromFloat64(255.18985417993153),
		Fdy:   float32(0.001013478497043252),
		Fe:    int32(FE_INEXACT),
	},
	1154: {
		Ffile: __ccgo_ts,
		Fline: int32(1171),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.807795575996853e-115),
		Fy:    -libc.Float64FromFloat64(263.46023526160104),
		Fdy:   float32(0.0005114726955071092),
		Fe:    int32(FE_INEXACT),
	},
	1155: {
		Ffile: __ccgo_ts,
		Fline: int32(1172),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.594966890626116e-116),
		Fy:    -libc.Float64FromFloat64(264.83863210854594),
		Fdy:   float32(0.0008591783116571605),
		Fe:    int32(FE_INEXACT),
	},
	1156: {
		Ffile: __ccgo_ts,
		Fline: int32(1173),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5351578826901497e-117),
		Fy:    -libc.Float64FromFloat64(268.9738226493807),
		Fdy:   float32(-libc.Float64FromFloat64(0.0005603072349913418)),
		Fe:    int32(FE_INEXACT),
	},
	1157: {
		Ffile: __ccgo_ts,
		Fline: int32(1174),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.189170814845369e-120),
		Fy:    -libc.Float64FromFloat64(274.4874100371603),
		Fdy:   float32(-libc.Float64FromFloat64(0.0007727963966317475)),
		Fe:    int32(FE_INEXACT),
	},
	1158: {
		Ffile: __ccgo_ts,
		Fline: int32(1175),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.992284268615321e-124),
		Fy:    -libc.Float64FromFloat64(284.13618796577475),
		Fdy:   float32(0.0015199100598692894),
		Fe:    int32(FE_INEXACT),
	},
	1159: {
		Ffile: __ccgo_ts,
		Fline: int32(1176),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.120216267115348e-130),
		Fy:    -libc.Float64FromFloat64(297.9201564352239),
		Fdy:   float32(-libc.Float64FromFloat64(0.00164736807346344)),
		Fe:    int32(FE_INEXACT),
	},
	1160: {
		Ffile: __ccgo_ts,
		Fline: int32(1177),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.592187917844459e-132),
		Fy:    -libc.Float64FromFloat64(302.0553469760586),
		Fdy:   float32(-libc.Float64FromFloat64(0.0013288217596709728)),
		Fe:    int32(FE_INEXACT),
	},
	1161: {
		Ffile: __ccgo_ts,
		Fline: int32(1178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0547247699354287e-133),
		Fy:    -libc.Float64FromFloat64(306.19053751689336),
		Fdy:   float32(-libc.Float64FromFloat64(0.001281566685065627)),
		Fe:    int32(FE_INEXACT),
	},
	1162: {
		Ffile: __ccgo_ts,
		Fline: int32(1179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.6999658954654556e-137),
		Fy:    -libc.Float64FromFloat64(314.46091859856284),
		Fdy:   float32(0.0007282128790393472),
		Fe:    int32(FE_INEXACT),
	},
	1163: {
		Ffile: __ccgo_ts,
		Fline: int32(1180),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7428847248946112e-140),
		Fy:    -libc.Float64FromFloat64(321.3529028332874),
		Fdy:   float32(-libc.Float64FromFloat64(6.367784953908995e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1164: {
		Ffile: __ccgo_ts,
		Fline: int32(1181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1058264611880535e-142),
		Fy:    -libc.Float64FromFloat64(326.86649022106707),
		Fdy:   float32(-libc.Float64FromFloat64(0.001082897069863975)),
		Fe:    int32(FE_INEXACT),
	},
	1165: {
		Ffile: __ccgo_ts,
		Fline: int32(1182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7974044885739664e-147),
		Fy:    -libc.Float64FromFloat64(337.89366499662634),
		Fdy:   float32(0.001266073901206255),
		Fe:    int32(FE_INEXACT),
	},
	1166: {
		Ffile: __ccgo_ts,
		Fline: int32(1183),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.875778207960991e-149),
		Fy:    -libc.Float64FromFloat64(342.0288555374611),
		Fdy:   float32(0.0004161993565503508),
		Fe:    int32(FE_INEXACT),
	},
	1167: {
		Ffile: __ccgo_ts,
		Fline: int32(1184),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.6011347773726225e-151),
		Fy:    -libc.Float64FromFloat64(346.1640460782958),
		Fdy:   float32(0.000728690589312464),
		Fe:    int32(FE_INEXACT),
	},
	1168: {
		Ffile: __ccgo_ts,
		Fline: int32(1185),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1594040427614375e-151),
		Fy:    -libc.Float64FromFloat64(347.5424429252407),
		Fdy:   float32(0.0008618015563115478),
		Fe:    int32(FE_INEXACT),
	},
	1169: {
		Ffile: __ccgo_ts,
		Fline: int32(1186),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.361639079447041e-153),
		Fy:    -libc.Float64FromFloat64(350.29923661913057),
		Fdy:   float32(0.0003124908544123173),
		Fe:    int32(FE_INEXACT),
	},
	1170: {
		Ffile: __ccgo_ts,
		Fline: int32(1187),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.855001977345962e-153),
		Fy:    -libc.Float64FromFloat64(351.67763346607546),
		Fdy:   float32(-libc.Float64FromFloat64(0.0007261309074237943)),
		Fe:    int32(FE_INEXACT),
	},
	1171: {
		Ffile: __ccgo_ts,
		Fline: int32(1188),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.015110415076262e-158),
		Fy:    -libc.Float64FromFloat64(362.7048082416348),
		Fdy:   float32(0.0009498826693743467),
		Fe:    int32(FE_INEXACT),
	},
	1172: {
		Ffile: __ccgo_ts,
		Fline: int32(1189),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.597541419135944e-159),
		Fy:    -libc.Float64FromFloat64(364.0832050885797),
		Fdy:   float32(-libc.Float64FromFloat64(0.00013283509179018438)),
		Fe:    int32(FE_INEXACT),
	},
	1173: {
		Ffile: __ccgo_ts,
		Fline: int32(1190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.0630388018131433e-161),
		Fy:    -libc.Float64FromFloat64(369.5967924763594),
		Fdy:   float32(9.751915058586746e-05),
		Fe:    int32(FE_INEXACT),
	},
	1174: {
		Ffile: __ccgo_ts,
		Fline: int32(1191),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.9786470882061086e-166),
		Fy:    -libc.Float64FromFloat64(380.62396725191866),
		Fdy:   float32(0.00010975053010042757),
		Fe:    int32(FE_INEXACT),
	},
	1175: {
		Ffile: __ccgo_ts,
		Fline: int32(1192),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2744725189812159e-169),
		Fy:    -libc.Float64FromFloat64(388.89434833358814),
		Fdy:   float32(-libc.Float64FromFloat64(0.0009085416095331311)),
		Fe:    int32(FE_INEXACT),
	},
	1176: {
		Ffile: __ccgo_ts,
		Fline: int32(1193),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.211443833729335e-170),
		Fy:    -libc.Float64FromFloat64(390.2727451805331),
		Fdy:   float32(-libc.Float64FromFloat64(0.00037306291051208973)),
		Fe:    int32(FE_INEXACT),
	},
	1177: {
		Ffile: __ccgo_ts,
		Fline: int32(1194),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.092266677859032e-171),
		Fy:    -libc.Float64FromFloat64(391.651142027478),
		Fdy:   float32(0.0006968426168896258),
		Fe:    int32(FE_INEXACT),
	},
	1178: {
		Ffile: __ccgo_ts,
		Fline: int32(1195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.314353947357611e-176),
		Fy:    -libc.Float64FromFloat64(404.0567136499822),
		Fdy:   float32(-libc.Float64FromFloat64(0.000841476721689105)),
		Fe:    int32(FE_INEXACT),
	},
	1179: {
		Ffile: __ccgo_ts,
		Fline: int32(1196),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.302839130463357e-178),
		Fy:    -libc.Float64FromFloat64(408.19190419081696),
		Fdy:   float32(-libc.Float64FromFloat64(0.0012502162717282772)),
		Fe:    int32(FE_INEXACT),
	},
	1180: {
		Ffile: __ccgo_ts,
		Fline: int32(1197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.379039998312719e-184),
		Fy:    -libc.Float64FromFloat64(423.354269507211),
		Fdy:   float32(8.395931945415214e-05),
		Fe:    int32(FE_INEXACT),
	},
	1181: {
		Ffile: __ccgo_ts,
		Fline: int32(1198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4749352638750525e-185),
		Fy:    -libc.Float64FromFloat64(424.7326663541559),
		Fdy:   float32(0.001146144000813365),
		Fe:    int32(FE_INEXACT),
	},
	1182: {
		Ffile: __ccgo_ts,
		Fline: int32(1199),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.559763074729752e-187),
		Fy:    -libc.Float64FromFloat64(428.86785689499067),
		Fdy:   float32(-libc.Float64FromFloat64(0.0011598450364544988)),
		Fe:    int32(FE_INEXACT),
	},
	1183: {
		Ffile: __ccgo_ts,
		Fline: int32(1200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.530173055625715e-188),
		Fy:    -libc.Float64FromFloat64(431.62465058888046),
		Fdy:   float32(-libc.Float64FromFloat64(0.0010448852553963661)),
		Fe:    int32(FE_INEXACT),
	},
	1184: {
		Ffile: __ccgo_ts,
		Fline: int32(1201),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.895407568732158e-189),
		Fy:    -libc.Float64FromFloat64(433.0030474358254),
		Fdy:   float32(0.0014824189711362123),
		Fe:    int32(FE_INEXACT),
	},
	1185: {
		Ffile: __ccgo_ts,
		Fline: int32(1202),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2414843285873012e-189),
		Fy:    -libc.Float64FromFloat64(434.3814442827703),
		Fdy:   float32(0.0012188251130282879),
		Fe:    int32(FE_INEXACT),
	},
	1186: {
		Ffile: __ccgo_ts,
		Fline: int32(1203),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.64814142183095e-190),
		Fy:    -libc.Float64FromFloat64(435.7598411297152),
		Fdy:   float32(0.0009733540937304497),
		Fe:    int32(FE_INEXACT),
	},
	1187: {
		Ffile: __ccgo_ts,
		Fline: int32(1204),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.44585484729281e-193),
		Fy:    -libc.Float64FromFloat64(444.0302222113847),
		Fdy:   float32(0.00030666353995911777),
		Fe:    int32(FE_INEXACT),
	},
	1188: {
		Ffile: __ccgo_ts,
		Fline: int32(1205),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3133122723078825e-195),
		Fy:    -libc.Float64FromFloat64(448.1654127522194),
		Fdy:   float32(0.0011125061428174376),
		Fe:    int32(FE_INEXACT),
	},
	1189: {
		Ffile: __ccgo_ts,
		Fline: int32(1206),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.7600455475193263e-200),
		Fy:    -libc.Float64FromFloat64(459.19258752777876),
		Fdy:   float32(0.0005225290660746396),
		Fe:    int32(FE_INEXACT),
	},
	1190: {
		Ffile: __ccgo_ts,
		Fline: int32(1207),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.3874419289752497e-201),
		Fy:    -libc.Float64FromFloat64(461.94938122166855),
		Fdy:   float32(0.0004960378282703459),
		Fe:    int32(FE_INEXACT),
	},
	1191: {
		Ffile: __ccgo_ts,
		Fline: int32(1208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.540003845493904e-205),
		Fy:    -libc.Float64FromFloat64(471.598159150283),
		Fdy:   float32(-libc.Float64FromFloat64(0.0009690800798125565)),
		Fe:    int32(FE_INEXACT),
	},
	1192: {
		Ffile: __ccgo_ts,
		Fline: int32(1209),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.942220738756091e-209),
		Fy:    -libc.Float64FromFloat64(479.86854023195247),
		Fdy:   float32(0.0010400231694802642),
		Fe:    int32(FE_INEXACT),
	},
	1193: {
		Ffile: __ccgo_ts,
		Fline: int32(1210),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.407664764303911e-214),
		Fy:    -libc.Float64FromFloat64(490.89571500751174),
		Fdy:   float32(-libc.Float64FromFloat64(0.0008846340933814645)),
		Fe:    int32(FE_INEXACT),
	},
	1194: {
		Ffile: __ccgo_ts,
		Fline: int32(1211),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.0685484608579646e-215),
		Fy:    -libc.Float64FromFloat64(493.6525087014016),
		Fdy:   float32(0.0003553261631168425),
		Fe:    int32(FE_INEXACT),
	},
	1195: {
		Ffile: __ccgo_ts,
		Fline: int32(1212),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0252017734597075e-215),
		Fy:    -libc.Float64FromFloat64(495.0309055483465),
		Fdy:   float32(-libc.Float64FromFloat64(0.0009924067417159677)),
		Fe:    int32(FE_INEXACT),
	},
	1196: {
		Ffile: __ccgo_ts,
		Fline: int32(1213),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.5095214105375704e-217),
		Fy:    -libc.Float64FromFloat64(497.78769924223633),
		Fdy:   float32(0.00014937073865439743),
		Fe:    int32(FE_INEXACT),
	},
	1197: {
		Ffile: __ccgo_ts,
		Fline: int32(1214),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.1332223657058305e-218),
		Fy:    -libc.Float64FromFloat64(500.5444929361261),
		Fdy:   float32(0.0009875799296423793),
		Fe:    int32(FE_INEXACT),
	},
	1198: {
		Ffile: __ccgo_ts,
		Fline: int32(1215),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6663575470337744e-220),
		Fy:    -libc.Float64FromFloat64(506.0580803239058),
		Fdy:   float32(0.00010991079761879519),
		Fe:    int32(FE_INEXACT),
	},
	1199: {
		Ffile: __ccgo_ts,
		Fline: int32(1216),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.198924331054276e-221),
		Fy:    -libc.Float64FromFloat64(507.4364771708507),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004232325591146946)),
		Fe:    int32(FE_INEXACT),
	},
	1200: {
		Ffile: __ccgo_ts,
		Fline: int32(1217),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.7181177997966e-223),
		Fy:    -libc.Float64FromFloat64(511.57166771168545),
		Fdy:   float32(-libc.Float64FromFloat64(0.000287248200038448)),
		Fe:    int32(FE_INEXACT),
	},
	1201: {
		Ffile: __ccgo_ts,
		Fline: int32(1218),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.933398621115243e-229),
		Fy:    -libc.Float64FromFloat64(525.3556361811346),
		Fdy:   float32(0.00014190976799000055),
		Fe:    int32(FE_INEXACT),
	},
	1202: {
		Ffile: __ccgo_ts,
		Fline: int32(1219),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.04361236614208e-232),
		Fy:    -libc.Float64FromFloat64(532.2476204158592),
		Fdy:   float32(0.0007234865333884954),
		Fe:    int32(FE_INEXACT),
	},
	1203: {
		Ffile: __ccgo_ts,
		Fline: int32(1220),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8397157438435856e-234),
		Fy:    -libc.Float64FromFloat64(537.7612078036389),
		Fdy:   float32(-libc.Float64FromFloat64(0.00014867981371935457)),
		Fe:    int32(FE_INEXACT),
	},
	1204: {
		Ffile: __ccgo_ts,
		Fline: int32(1221),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.155578075862806e-235),
		Fy:    -libc.Float64FromFloat64(539.1396046505837),
		Fdy:   float32(-libc.Float64FromFloat64(1.408677599101793e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1205: {
		Ffile: __ccgo_ts,
		Fline: int32(1222),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.269323599618685e-238),
		Fy:    -libc.Float64FromFloat64(546.0315888853083),
		Fdy:   float32(5.159335341886617e-05),
		Fe:    int32(FE_INEXACT),
	},
	1206: {
		Ffile: __ccgo_ts,
		Fline: int32(1223),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.384877229448362e-241),
		Fy:    -libc.Float64FromFloat64(552.9235731200329),
		Fdy:   float32(0.0003806101158261299),
		Fe:    int32(FE_INEXACT),
	},
	1207: {
		Ffile: __ccgo_ts,
		Fline: int32(1224),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.7635668930101736e-245),
		Fy:    -libc.Float64FromFloat64(562.5723510486473),
		Fdy:   float32(-libc.Float64FromFloat64(5.595584661932662e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1208: {
		Ffile: __ccgo_ts,
		Fline: int32(1225),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.839288855016746e-248),
		Fy:    -libc.Float64FromFloat64(569.4643352833718),
		Fdy:   float32(-libc.Float64FromFloat64(4.875911690760404e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1209: {
		Ffile: __ccgo_ts,
		Fline: int32(1226),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9510166142385655e-250),
		Fy:    -libc.Float64FromFloat64(574.9779226711514),
		Fdy:   float32(-libc.Float64FromFloat64(1.778921068762429e-05)),
		Fe:    int32(FE_INEXACT),
	},
	1210: {
		Ffile: __ccgo_ts,
		Fline: int32(1227),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.91621449814272e-251),
		Fy:    -libc.Float64FromFloat64(576.3563195180964),
		Fdy:   float32(0.000118538475362584),
		Fe:    int32(FE_INEXACT),
	},
	1211: {
		Ffile: __ccgo_ts,
		Fline: int32(1228),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.865754541782381e-253),
		Fy:    -libc.Float64FromFloat64(580.4915100589311),
		Fdy:   float32(-libc.Float64FromFloat64(0.0002639983722474426)),
		Fe:    int32(FE_INEXACT),
	},
	1212: {
		Ffile: __ccgo_ts,
		Fline: int32(1229),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.9943629561796515e-254),
		Fy:    -libc.Float64FromFloat64(583.2483037528209),
		Fdy:   float32(0.0003083782212343067),
		Fe:    int32(FE_INEXACT),
	},
	1213: {
		Ffile: __ccgo_ts,
		Fline: int32(1230),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.258490542569773e-254),
		Fy:    -libc.Float64FromFloat64(584.6267005997659),
		Fdy:   float32(-libc.Float64FromFloat64(0.00031685474095866084)),
		Fe:    int32(FE_INEXACT),
	},
	1214: {
		Ffile: __ccgo_ts,
		Fline: int32(1231),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2784955881984868e-257),
		Fy:    -libc.Float64FromFloat64(591.5186848344904),
		Fdy:   float32(-libc.Float64FromFloat64(0.0003560609184205532)),
		Fe:    int32(FE_INEXACT),
	},
	1215: {
		Ffile: __ccgo_ts,
		Fline: int32(1232),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.246852408752646e-262),
		Fy:    -libc.Float64FromFloat64(601.1674627631048),
		Fdy:   float32(0.000124257494462654),
		Fe:    int32(FE_INEXACT),
	},
	1216: {
		Ffile: __ccgo_ts,
		Fline: int32(1233),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.361746738242124e-269),
		Fy:    -libc.Float64FromFloat64(619.0866217733887),
		Fdy:   float32(-libc.Float64FromFloat64(0.00012405644520185888)),
		Fe:    int32(FE_INEXACT),
	},
	1217: {
		Ffile: __ccgo_ts,
		Fline: int32(1234),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.4313593274847667e-270),
		Fy:    -libc.Float64FromFloat64(620.4650186203336),
		Fdy:   float32(-libc.Float64FromFloat64(0.000513899081852287)),
		Fe:    int32(FE_INEXACT),
	},
	1218: {
		Ffile: __ccgo_ts,
		Fline: int32(1235),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.383393150918106e-272),
		Fy:    -libc.Float64FromFloat64(625.9786060081133),
		Fdy:   float32(2.2048656319384463e-05),
		Fe:    int32(FE_INEXACT),
	},
	1219: {
		Ffile: __ccgo_ts,
		Fline: int32(1236),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.577313325007993e-275),
		Fy:    -libc.Float64FromFloat64(631.492193395893),
		Fdy:   float32(3.920057861250825e-05),
		Fe:    int32(FE_INEXACT),
	},
	1220: {
		Ffile: __ccgo_ts,
		Fline: int32(1237),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.541316591747485e-276),
		Fy:    -libc.Float64FromFloat64(634.2489870897828),
		Fdy:   float32(0.0003507465007714927),
		Fe:    int32(FE_INEXACT),
	},
	1221: {
		Ffile: __ccgo_ts,
		Fline: int32(1238),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.923487295702502e-277),
		Fy:    -libc.Float64FromFloat64(635.6273839367277),
		Fdy:   float32(0.0003265703562647104),
		Fe:    int32(FE_INEXACT),
	},
	1222: {
		Ffile: __ccgo_ts,
		Fline: int32(1239),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2843031711951766e-280),
		Fy:    -libc.Float64FromFloat64(643.8977650183972),
		Fdy:   float32(1.274151054531103e-05),
		Fe:    int32(FE_INEXACT),
	},
	1223: {
		Ffile: __ccgo_ts,
		Fline: int32(1240),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.034918794841685e-290),
		Fy:    -libc.Float64FromFloat64(665.9521145695157),
		Fdy:   float32(0.00040247890865430236),
		Fe:    int32(FE_INEXACT),
	},
	1224: {
		Ffile: __ccgo_ts,
		Fline: int32(1241),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.655638487244009e-292),
		Fy:    -libc.Float64FromFloat64(670.0873051103505),
		Fdy:   float32(0.0005686437361873686),
		Fe:    int32(FE_INEXACT),
	},
	1225: {
		Ffile: __ccgo_ts,
		Fline: int32(1242),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.32731219052115e-299),
		Fy:    -libc.Float64FromFloat64(686.6280672736895),
		Fdy:   float32(0.0004820417962037027),
		Fe:    int32(FE_INEXACT),
	},
	1226: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(709.782712893384),
		Fdy:   float32(-libc.Float64FromFloat64(0.2079046070575714)),
		Fe:    int32(FE_INEXACT),
	},
	1227: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(39),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(709.782712893384),
		Fdy:   float32(-libc.Float64FromFloat64(0.2079046070575714)),
		Fe:    int32(FE_INEXACT),
	},
	1228: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(40),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(709.7827128933841),
		Fdy:   float32(0.792095422744751),
		Fe:    int32(FE_INEXACT),
	},
	1229: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(41),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(709.782712893384),
		Fdy:   float32(-libc.Float64FromFloat64(0.2079046070575714)),
		Fe:    int32(FE_INEXACT),
	},
	1230: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.000000000000007),
		Fy:    libc.Float64FromFloat64(7.105427357600977e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1.515824548129971e-13)),
		Fe:    int32(FE_INEXACT),
	},
	1231: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1232: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(7.105427357601027e-15),
		Fdy:   float32(7.579122740649855e-14),
		Fe:    int32(FE_INEXACT),
	},
	1233: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(46),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000000000000007),
		Fy:    libc.Float64FromFloat64(7.105427357600977e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1.515824548129971e-13)),
		Fe:    int32(FE_INEXACT),
	},
	1234: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(47),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1235: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(48),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(7.105427357601027e-15),
		Fdy:   float32(7.579122740649855e-14),
		Fe:    int32(FE_INEXACT),
	},
	1236: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(49),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000007),
		Fy:    libc.Float64FromFloat64(7.1054273576009774e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1237: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(50),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1238: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(51),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(7.105427357601027e-15),
		Fdy:   float32(7.579122740649855e-14),
		Fe:    int32(FE_INEXACT),
	},
	1239: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(52),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000007),
		Fy:    libc.Float64FromFloat64(7.105427357600977e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1.515824548129971e-13)),
		Fe:    int32(FE_INEXACT),
	},
	1240: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(53),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1241: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(54),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(7.105427357601029e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1242: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(708.3964185322641),
		Fdy:   float32(0.24167631566524506),
		Fe:    int32(FE_INEXACT),
	},
	1243: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	1244: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(58),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(708.3964185322641),
		Fdy:   float32(0.24167631566524506),
		Fe:    int32(FE_INEXACT),
	},
	1245: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(59),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	1246: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(60),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(708.3964185322641),
		Fdy:   float32(0.24167631566524506),
		Fe:    int32(FE_INEXACT),
	},
	1247: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(61),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213812),
		Fdy:   float32(0.38900232315063477),
		Fe:    int32(FE_INEXACT),
	},
	1248: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(62),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(708.3964185322642),
		Fdy:   float32(-libc.Float64FromFloat64(0.7583236694335938)),
		Fe:    int32(FE_INEXACT),
	},
	1249: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(63),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(744.4400719213813),
		Fdy:   float32(-libc.Float64FromFloat64(0.6109976768493652)),
		Fe:    int32(FE_INEXACT),
	},
	1250: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(36.56459649885151),
		Fy:    libc.Float64FromFloat64(3.599080463487721),
		Fdy:   float32(0.4961016774177551),
		Fe:    int32(FE_INEXACT),
	},
	1251: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(85.61132705430306),
		Fy:    libc.Float64FromFloat64(4.4498175997941996),
		Fdy:   float32(0.020959023386240005),
		Fe:    int32(FE_INEXACT),
	},
	1252: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(94.02246329044654),
		Fy:    libc.Float64FromFloat64(4.543533724895904),
		Fdy:   float32(0.24992160499095917),
		Fe:    int32(FE_INEXACT),
	},
	1253: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(16.60725376974054),
		Fy:    libc.Float64FromFloat64(2.809839573997883),
		Fdy:   float32(0.00979764387011528),
		Fe:    int32(FE_INEXACT),
	},
	1254: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(36.45486097807018),
		Fy:    libc.Float64FromFloat64(3.596074809729525),
		Fdy:   float32(-libc.Float64FromFloat64(0.24034123122692108)),
		Fe:    int32(FE_INEXACT),
	},
	1255: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(70),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(48.03540361357637),
		Fy:    libc.Float64FromFloat64(3.87193831431576),
		Fdy:   float32(-libc.Float64FromFloat64(0.360834538936615)),
		Fe:    int32(FE_INEXACT),
	},
	1256: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(71),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(78.91655385898032),
		Fy:    libc.Float64FromFloat64(4.368391013943797),
		Fdy:   float32(-libc.Float64FromFloat64(0.3010043501853943)),
		Fe:    int32(FE_INEXACT),
	},
	1257: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(72),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(92.75441225031035),
		Fy:    libc.Float64FromFloat64(4.529955271787977),
		Fdy:   float32(-libc.Float64FromFloat64(0.21752695739269257)),
		Fe:    int32(FE_INEXACT),
	},
	1258: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(41.596529063707386),
		Fy:    libc.Float64FromFloat64(3.7280167278176712),
		Fdy:   float32(0.3062756061553955),
		Fe:    int32(FE_INEXACT),
	},
	1259: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.098780043531456),
		Fy:    libc.Float64FromFloat64(0.7413562442715179),
		Fdy:   float32(-libc.Float64FromFloat64(0.1252504587173462)),
		Fe:    int32(FE_INEXACT),
	},
	1260: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1261: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1262: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1263: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1264: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1265: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1266: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1267: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1268: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(87),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1269: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(88),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1270: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(89),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1271: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(90),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    -libc.Float64FromFloat64(4.440892098500628e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1272: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(91),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    -libc.Float64FromFloat64(3.3306690738754706e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.875)),
		Fe:    int32(FE_INEXACT),
	},
	1273: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(92),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(2.2204460492503136e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	1274: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(93),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251568e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.75)),
		Fe:    int32(FE_INEXACT),
	},
	1275: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(94),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1.4802972779342767e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1276: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(95),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(4.440892098500625e-16),
		Fdy:   float32(-libc.Float64FromFloat64(5.921189111737107e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1277: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(96),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(8.881784197001248e-16),
		Fdy:   float32(-libc.Float64FromFloat64(2.3684758564530796e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1278: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(97),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1279: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(98),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1280: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(99),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1281: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(100),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1282: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(101),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1283: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(102),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1284: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(103),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1285: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(104),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1286: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(105),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1287: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(106),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1288: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(107),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1289: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(108),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1290: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(109),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1291: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1292: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1293: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(112),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1294: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1295: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1296: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1297: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1298: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1299: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1300: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1301: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1302: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(121),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1303: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1304: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(123),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1305: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1306: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1307: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(126),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1308: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(127),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1309: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1310: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1311: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(130),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1312: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(131),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1313: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1314: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(133),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1315: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(134),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1316: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    -libc.Float64FromFloat64(9.992007221626415e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.46875)),
		Fe:    int32(FE_INEXACT),
	},
	1317: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(136),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    -libc.Float64FromFloat64(8.881784197001258e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1318: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(137),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    -libc.Float64FromFloat64(7.7715611723761e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.9375)),
		Fe:    int32(FE_INEXACT),
	},
	1319: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    -libc.Float64FromFloat64(6.661338147750942e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.75)),
		Fe:    int32(FE_INEXACT),
	},
	1320: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(139),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    -libc.Float64FromFloat64(5.551115123125785e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.4375)),
		Fe:    int32(FE_INEXACT),
	},
	1321: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(140),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(6.661338147750936e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.75)),
		Fe:    int32(FE_INEXACT),
	},
	1322: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(141),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.875)),
		Fe:    int32(FE_INEXACT),
	},
	1323: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(142),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.5543122344752178e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.875)),
		Fe:    int32(FE_INEXACT),
	},
	1324: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(143),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000018),
		Fy:    libc.Float64FromFloat64(1.7763568394002489e-15),
		Fdy:   float32(-libc.Float64FromFloat64(9.473903425812318e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1325: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(144),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000002),
		Fy:    libc.Float64FromFloat64(1.9984014443252794e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9375)),
		Fe:    int32(FE_INEXACT),
	},
	1326: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    -libc.Float64FromFloat64(9.992007221626415e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.46875)),
		Fe:    int32(FE_INEXACT),
	},
	1327: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    -libc.Float64FromFloat64(8.881784197001256e-16),
		Fdy:   float32(1.1842379282265398e-15),
		Fe:    int32(FE_INEXACT),
	},
	1328: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    -libc.Float64FromFloat64(7.771561172376099e-16),
		Fdy:   float32(0.0625),
		Fe:    int32(FE_INEXACT),
	},
	1329: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    -libc.Float64FromFloat64(6.661338147750941e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1330: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    -libc.Float64FromFloat64(5.551115123125785e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.4375)),
		Fe:    int32(FE_INEXACT),
	},
	1331: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    -libc.Float64FromFloat64(4.440892098500627e-16),
		Fdy:   float32(2.9605948205663494e-16),
		Fe:    int32(FE_INEXACT),
	},
	1332: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    -libc.Float64FromFloat64(3.33066907387547e-16),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1333: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251565e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1334: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1.4802972779342767e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1335: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(4.440892098500625e-16),
		Fdy:   float32(-libc.Float64FromFloat64(5.921189111737107e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1336: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(6.661338147750937e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1337: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(8.881784197001248e-16),
		Fdy:   float32(-libc.Float64FromFloat64(2.3684758564530796e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1338: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-15),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1339: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.554312234475218e-15),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1340: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0000000000000018),
		Fy:    libc.Float64FromFloat64(1.7763568394002489e-15),
		Fdy:   float32(-libc.Float64FromFloat64(9.473903425812318e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1341: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.000000000000002),
		Fy:    libc.Float64FromFloat64(1.9984014443252798e-15),
		Fdy:   float32(0.0625),
		Fe:    int32(FE_INEXACT),
	},
	1342: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1343: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1344: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1345: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1346: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1347: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(166),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1348: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1349: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1350: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1351: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1352: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1353: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(172),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1354: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1355: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(174),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1356: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1357: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1358: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1359: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1360: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1361: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1362: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1363: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1364: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1365: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1366: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1367: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1368: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1369: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(188),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1370: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(189),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1371: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1372: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1373: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1374: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1375: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(194),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1376: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(195),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1377: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(196),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1378: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(197),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    -libc.Float64FromFloat64(4.440892098500627e-16),
		Fdy:   float32(2.9605945558685534e-16),
		Fe:    int32(FE_INEXACT),
	},
	1379: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(198),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    -libc.Float64FromFloat64(3.33066907387547e-16),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1380: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(199),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	1381: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(200),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251565e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1382: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(201),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(2.220446049250313e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	1383: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(202),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(4.4408920985006257e-16),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1384: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(203),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(8.881784197001249e-16),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1385: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(204),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1386: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(205),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1387: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(206),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1388: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(207),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1389: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(208),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1390: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(209),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1391: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(210),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1392: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(211),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1393: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(212),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1394: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(213),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1395: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(214),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1396: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(215),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1397: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(216),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1398: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(217),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1399: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(218),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1400: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(219),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1401: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(220),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1402: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(221),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1403: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(222),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1404: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(223),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1405: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(224),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1406: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(225),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1407: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(226),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1408: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(227),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1409: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(228),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1410: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(229),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1411: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(230),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1412: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(231),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1413: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(232),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1414: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(233),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1415: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(234),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1416: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(235),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1417: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(236),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1418: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(237),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1419: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(238),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1420: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(239),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1421: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(240),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1422: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(241),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1423: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(242),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    -libc.Float64FromFloat64(9.992007221626413e-16),
		Fdy:   float32(0.53125),
		Fe:    int32(FE_INEXACT),
	},
	1424: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(243),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    -libc.Float64FromFloat64(8.881784197001256e-16),
		Fdy:   float32(1.1842378223474214e-15),
		Fe:    int32(FE_INEXACT),
	},
	1425: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(244),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    -libc.Float64FromFloat64(7.771561172376099e-16),
		Fdy:   float32(0.0625),
		Fe:    int32(FE_INEXACT),
	},
	1426: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(245),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    -libc.Float64FromFloat64(6.661338147750941e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1427: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(246),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-16),
		Fdy:   float32(0.5625),
		Fe:    int32(FE_INEXACT),
	},
	1428: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(247),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(6.661338147750937e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1429: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(248),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.110223024625156e-15),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1430: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(249),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.554312234475218e-15),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1431: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(250),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.0000000000000018),
		Fy:    libc.Float64FromFloat64(1.776356839400249e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1432: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(251),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.000000000000002),
		Fy:    libc.Float64FromFloat64(1.9984014443252798e-15),
		Fdy:   float32(0.0625),
		Fe:    int32(FE_INEXACT),
	},
	1433: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1434: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1435: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1436: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999996),
		Fy:    -libc.Float64FromFloat64(4.440892098500627e-16),
		Fdy:   float32(2.9605945558685534e-16),
		Fe:    int32(FE_INEXACT),
	},
	1437: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(256),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999997),
		Fy:    -libc.Float64FromFloat64(3.33066907387547e-16),
		Fdy:   float32(0.125),
		Fe:    int32(FE_INEXACT),
	},
	1438: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999998),
		Fy:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	1439: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999999),
		Fy:    -libc.Float64FromFloat64(1.1102230246251565e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1440: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(259),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000002),
		Fy:    libc.Float64FromFloat64(2.2204460492503128e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1.4802972779342767e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1441: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000004),
		Fy:    libc.Float64FromFloat64(4.440892098500625e-16),
		Fdy:   float32(-libc.Float64FromFloat64(5.921189111737107e-16)),
		Fe:    int32(FE_INEXACT),
	},
	1442: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000009),
		Fy:    libc.Float64FromFloat64(8.881784197001248e-16),
		Fdy:   float32(-libc.Float64FromFloat64(2.3684758564530796e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1443: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1444: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1445: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(264),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1446: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(265),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1447: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1448: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1449: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(268),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1450: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(269),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1451: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(270),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1452: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1453: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(272),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1454: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(273),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1455: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1456: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1457: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1458: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1459: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(278),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1460: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(279),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1461: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(280),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1462: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(281),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1463: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1464: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(283),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1465: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(284),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1466: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(285),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1467: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(286),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1468: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(287),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1469: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1470: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(289),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1471: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(290),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1472: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(291),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1473: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(292),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1474: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(293),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1475: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(294),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1476: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(295),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1477: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(296),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1478: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(297),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1479: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(298),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1480: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(299),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1481: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(300),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.999999999999999),
		Fy:    -libc.Float64FromFloat64(9.992007221626413e-16),
		Fdy:   float32(0.53125),
		Fe:    int32(FE_INEXACT),
	},
	1482: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(301),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999991),
		Fy:    -libc.Float64FromFloat64(8.881784197001256e-16),
		Fdy:   float32(1.1842378223474214e-15),
		Fe:    int32(FE_INEXACT),
	},
	1483: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(302),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999992),
		Fy:    -libc.Float64FromFloat64(7.771561172376099e-16),
		Fdy:   float32(0.0625),
		Fe:    int32(FE_INEXACT),
	},
	1484: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999993),
		Fy:    -libc.Float64FromFloat64(6.661338147750941e-16),
		Fdy:   float32(0.25),
		Fe:    int32(FE_INEXACT),
	},
	1485: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(304),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9999999999999994),
		Fy:    -libc.Float64FromFloat64(5.551115123125784e-16),
		Fdy:   float32(0.5625),
		Fe:    int32(FE_INEXACT),
	},
	1486: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(305),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000007),
		Fy:    libc.Float64FromFloat64(6.661338147750936e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.75)),
		Fe:    int32(FE_INEXACT),
	},
	1487: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000000000000001),
		Fy:    libc.Float64FromFloat64(1.1102230246251558e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.875)),
		Fe:    int32(FE_INEXACT),
	},
	1488: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(307),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000016),
		Fy:    libc.Float64FromFloat64(1.5543122344752178e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.875)),
		Fe:    int32(FE_INEXACT),
	},
	1489: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(308),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0000000000000018),
		Fy:    libc.Float64FromFloat64(1.7763568394002489e-15),
		Fdy:   float32(-libc.Float64FromFloat64(9.473903425812318e-15)),
		Fe:    int32(FE_INEXACT),
	},
	1490: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(309),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.000000000000002),
		Fy:    libc.Float64FromFloat64(1.9984014443252794e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9375)),
		Fe:    int32(FE_INEXACT),
	},
	1491: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1492: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fy:    libc.Float64FromFloat64(1.4690809584224322),
		Fdy:   float32(-libc.Float64FromFloat64(0.3412533402442932)),
		Fe:    int32(FE_INEXACT),
	},
	1493: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1494: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1495: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fy:    libc.Float64FromFloat64(2.2264658498795615),
		Fdy:   float32(0.3638114035129547),
		Fe:    int32(FE_INEXACT),
	},
	1496: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619858980995045),
		Fy:    -libc.Float64FromFloat64(0.4125110252365137),
		Fdy:   float32(-libc.Float64FromFloat64(0.29108747839927673)),
		Fe:    int32(FE_INEXACT),
	},
	1497: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1498: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5617597462207241),
		Fy:    -libc.Float64FromFloat64(0.5766810183195862),
		Fdy:   float32(-libc.Float64FromFloat64(0.10983199626207352)),
		Fe:    int32(FE_INEXACT),
	},
	1499: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7741522965913037),
		Fy:    -libc.Float64FromFloat64(0.2559866591263865),
		Fdy:   float32(-libc.Float64FromFloat64(0.057990044355392456)),
		Fe:    int32(FE_INEXACT),
	},
	1500: {
		Ffile: __ccgo_ts + 42,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1501: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1502: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	1503: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.888609052210118e-31),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1504: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1505: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1506: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1507: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1508: {
		Ffile: __ccgo_ts + 64,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+22)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:19:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:20:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:24:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:27:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(54324)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:28:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:30:3:
		if (*l_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:31:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:32:3:
		libc.Xfesetround(tls, (*l_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:33:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:34:3:
		y = libc.Xlogl(tls, (*l_l)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:35:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:37:3:
		if !(checkexcept(tls, e, (*l_l)(unsafe.Pointer(p)).Fe, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:38:4:
			libc.Xprintf(tls, __ccgo_ts+87, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, estr(tls, (*l_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:40:4:
			libc.Xprintf(tls, __ccgo_ts+138, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:41:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:43:3:
		d = ulperrl(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:44:3:
		if !(checkulp(tls, d, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:45:4:
			libc.Xprintf(tls, __ccgo_ts+147, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*l_l)(unsafe.Pointer(p)).Fdy), float64((*l_l)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:47:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/logl.c:50:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+207, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+217, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 261,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 269,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 277,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 287,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 297,
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
				v2 = __ccgo_ts + 306
			} else {
				v2 = __ccgo_ts + 22
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+308, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 306
		} else {
			v3 = __ccgo_ts + 22
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+313, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 22
	} else {
		v4 = __ccgo_ts + 318
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+320, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 323
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 326
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 329
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 332
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 335
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+338, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+346, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+351, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+394, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+443) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+451) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+463) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+475) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+487) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+496) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+22) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+496) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+502, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/log.h\x00\x00src/math/ucb/log.h\x00src/math/sanity/log.h\x00src/math/special/log.h\x00%s:%d: bad fp exception: %s logl(%La)=%La, want %s\x00 got %s\n\x00%s:%d: %s logl(%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
