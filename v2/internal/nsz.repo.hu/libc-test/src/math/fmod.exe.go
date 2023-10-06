// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/fmod.exe.go src/math/fmod.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:5:20:
var t = [1051]dd_d{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(8.988465674311579e+307),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(8.988465674311579e+307),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(8.988465674311578e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(8.988465674311578e+307),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.98846567431158e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623155e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(1.7976931348623155e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623155e+308),
		Fx2:   float64(8.988465674311579e+307),
		Fy:    float64(8.988465674311577e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fx2:   float64(8.988465674311579e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311577e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(26.87084500235544),
		Fx2:   float64(71.22259414414681),
		Fy:    -libc.Float64FromFloat64(26.87084500235544),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(88.04498526310279),
		Fx2:   -libc.Float64FromFloat64(66.78552142661842),
		Fy:    float64(21.25946383648437),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(27.090315010519465),
		Fx2:   -libc.Float64FromFloat64(3.929197512835828),
		Fy:    -libc.Float64FromFloat64(3.5151299335044977),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(57.833124341599536),
		Fx2:   float64(85.5088678167166),
		Fy:    float64(57.833124341599536),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(16.806885747674006),
		Fx2:   -libc.Float64FromFloat64(95.80244088519409),
		Fy:    -libc.Float64FromFloat64(16.806885747674006),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(41.68758991425401),
		Fx2:   float64(4.621773070236693),
		Fy:    -libc.Float64FromFloat64(0.09163228212377472),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(75.84228190126574),
		Fx2:   float64(71.21820707478629),
		Fy:    -libc.Float64FromFloat64(4.624074826479443),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(99.61681744275423),
		Fx2:   -libc.Float64FromFloat64(99.8875066991044),
		Fy:    -libc.Float64FromFloat64(99.61681744275423),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(45.536989697633864),
		Fx2:   float64(58.14706899562555),
		Fy:    -libc.Float64FromFloat64(45.536989697633864),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(28.70274214255075),
		Fx2:   float64(49.3734605429732),
		Fy:    float64(28.70274214255075),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623155e+308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623153e+308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623151e+308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.797693134862315e+308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623147e+308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507204e-308),
		Fx2:   float64(2.2250738585072043e-308),
		Fy:    float64(2.225073858507204e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507204e-308),
		Fx2:   float64(2.2250738585072034e-308),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(0.5),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(0.9999999999999999),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507202e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.225073858507203e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072034e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.2250738585072034e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072043e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.4501477170144023e-308),
		Fx2:   float64(4.450147717014403e-308),
		Fy:    float64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.139237815555687e-305),
		Fx2:   float64(5.696189077778436e-306),
		Fy:    float64(5.696189077778434e-306),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.125),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.25),
		Fx2:   float64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.25000000000000006),
		Fx2:   float64(0.5),
		Fy:    float64(0.25000000000000006),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.375),
		Fx2:   float64(0.5),
		Fy:    float64(0.375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.625),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999996),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(4.440892098500626e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(2),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000004),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000007),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000007),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000009),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000009),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000013),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.9999999999999998),
		Fx2:   float64(2),
		Fy:    float64(1.9999999999999998),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2),
		Fx2:   float64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.0000000000000036),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000036),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(2),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.999999999999993),
		Fx2:   float64(4),
		Fy:    float64(1.999999999999993),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(7),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.99584030953472e+292),
		Fx2:   float64(7.98336123813888e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.49423283715579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(4.49423283715579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311576e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(8.988465674311576e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.98846567431158e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311582e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(3.99168061906944e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311586e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311588e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311588e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.988465674311592e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(0.5),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(0.9999999999999999),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507202e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.225073858507203e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072034e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.2250738585072034e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072043e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.4501477170144023e-308),
		Fx2:   float64(4.450147717014403e-308),
		Fy:    float64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.139237815555687e-305),
		Fx2:   float64(5.696189077778436e-306),
		Fy:    float64(5.696189077778434e-306),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.125),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.25),
		Fx2:   float64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.25000000000000006),
		Fx2:   float64(0.5),
		Fy:    float64(0.25000000000000006),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.375),
		Fx2:   float64(0.5),
		Fy:    float64(0.375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.625),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999996),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(4.440892098500626e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(2),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000004),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000007),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000007),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000009),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000009),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000013),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.9999999999999998),
		Fx2:   float64(2),
		Fy:    float64(1.9999999999999998),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fx2:   float64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(448),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(449),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(450),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(451),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(452),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.0000000000000036),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000036),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(453),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(2),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(454),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(455),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(456),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(457),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(458),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(459),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(460),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(461),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.999999999999993),
		Fx2:   float64(4),
		Fy:    float64(1.999999999999993),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(462),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(463),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(464),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(465),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(466),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(467),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.99584030953472e+292),
		Fx2:   float64(7.98336123813888e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(468),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.49423283715579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(4.49423283715579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(469),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311576e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(8.988465674311576e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(470),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(471),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(472),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(473),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(474),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(475),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(476),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(477),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(478),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311582e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(3.99168061906944e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(479),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(480),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(481),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(482),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(483),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(484),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311586e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(485),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(486),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311588e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311588e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(487),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.988465674311592e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(488),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	431: {
		Ffile: __ccgo_ts,
		Fline: int32(489),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	432: {
		Ffile: __ccgo_ts,
		Fline: int32(490),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	433: {
		Ffile: __ccgo_ts,
		Fline: int32(491),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	434: {
		Ffile: __ccgo_ts,
		Fline: int32(492),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	435: {
		Ffile: __ccgo_ts,
		Fline: int32(493),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	436: {
		Ffile: __ccgo_ts,
		Fline: int32(494),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	437: {
		Ffile: __ccgo_ts,
		Fline: int32(495),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	438: {
		Ffile: __ccgo_ts,
		Fline: int32(496),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	439: {
		Ffile: __ccgo_ts,
		Fline: int32(497),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	440: {
		Ffile: __ccgo_ts,
		Fline: int32(498),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	441: {
		Ffile: __ccgo_ts,
		Fline: int32(499),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	442: {
		Ffile: __ccgo_ts,
		Fline: int32(500),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	443: {
		Ffile: __ccgo_ts,
		Fline: int32(501),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	444: {
		Ffile: __ccgo_ts,
		Fline: int32(502),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	445: {
		Ffile: __ccgo_ts,
		Fline: int32(503),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	446: {
		Ffile: __ccgo_ts,
		Fline: int32(504),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	447: {
		Ffile: __ccgo_ts,
		Fline: int32(505),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	448: {
		Ffile: __ccgo_ts,
		Fline: int32(506),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	449: {
		Ffile: __ccgo_ts,
		Fline: int32(507),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	450: {
		Ffile: __ccgo_ts,
		Fline: int32(508),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	451: {
		Ffile: __ccgo_ts,
		Fline: int32(509),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	452: {
		Ffile: __ccgo_ts,
		Fline: int32(510),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	453: {
		Ffile: __ccgo_ts,
		Fline: int32(511),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	454: {
		Ffile: __ccgo_ts,
		Fline: int32(512),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	455: {
		Ffile: __ccgo_ts,
		Fline: int32(513),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	456: {
		Ffile: __ccgo_ts,
		Fline: int32(514),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	457: {
		Ffile: __ccgo_ts,
		Fline: int32(515),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	458: {
		Ffile: __ccgo_ts,
		Fline: int32(516),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	459: {
		Ffile: __ccgo_ts,
		Fline: int32(517),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	460: {
		Ffile: __ccgo_ts,
		Fline: int32(518),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	461: {
		Ffile: __ccgo_ts,
		Fline: int32(519),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	462: {
		Ffile: __ccgo_ts,
		Fline: int32(520),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	463: {
		Ffile: __ccgo_ts,
		Fline: int32(521),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	464: {
		Ffile: __ccgo_ts,
		Fline: int32(522),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	465: {
		Ffile: __ccgo_ts,
		Fline: int32(523),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	466: {
		Ffile: __ccgo_ts,
		Fline: int32(524),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	467: {
		Ffile: __ccgo_ts,
		Fline: int32(525),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	468: {
		Ffile: __ccgo_ts,
		Fline: int32(526),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	469: {
		Ffile: __ccgo_ts,
		Fline: int32(527),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	470: {
		Ffile: __ccgo_ts,
		Fline: int32(528),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	471: {
		Ffile: __ccgo_ts,
		Fline: int32(529),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	472: {
		Ffile: __ccgo_ts,
		Fline: int32(530),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	473: {
		Ffile: __ccgo_ts,
		Fline: int32(531),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	474: {
		Ffile: __ccgo_ts,
		Fline: int32(532),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	475: {
		Ffile: __ccgo_ts,
		Fline: int32(533),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	476: {
		Ffile: __ccgo_ts,
		Fline: int32(534),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	477: {
		Ffile: __ccgo_ts,
		Fline: int32(535),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	478: {
		Ffile: __ccgo_ts,
		Fline: int32(536),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	479: {
		Ffile: __ccgo_ts,
		Fline: int32(537),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	480: {
		Ffile: __ccgo_ts,
		Fline: int32(538),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	481: {
		Ffile: __ccgo_ts,
		Fline: int32(539),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	482: {
		Ffile: __ccgo_ts,
		Fline: int32(540),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	483: {
		Ffile: __ccgo_ts,
		Fline: int32(541),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	484: {
		Ffile: __ccgo_ts,
		Fline: int32(542),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	485: {
		Ffile: __ccgo_ts,
		Fline: int32(543),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	486: {
		Ffile: __ccgo_ts,
		Fline: int32(544),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	487: {
		Ffile: __ccgo_ts,
		Fline: int32(545),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	488: {
		Ffile: __ccgo_ts,
		Fline: int32(546),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	489: {
		Ffile: __ccgo_ts,
		Fline: int32(547),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	490: {
		Ffile: __ccgo_ts,
		Fline: int32(548),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	491: {
		Ffile: __ccgo_ts,
		Fline: int32(549),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	492: {
		Ffile: __ccgo_ts,
		Fline: int32(550),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	493: {
		Ffile: __ccgo_ts,
		Fline: int32(551),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	494: {
		Ffile: __ccgo_ts,
		Fline: int32(552),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	495: {
		Ffile: __ccgo_ts,
		Fline: int32(553),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	496: {
		Ffile: __ccgo_ts,
		Fline: int32(554),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	497: {
		Ffile: __ccgo_ts,
		Fline: int32(555),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	498: {
		Ffile: __ccgo_ts,
		Fline: int32(556),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	499: {
		Ffile: __ccgo_ts,
		Fline: int32(557),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	500: {
		Ffile: __ccgo_ts,
		Fline: int32(558),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	501: {
		Ffile: __ccgo_ts,
		Fline: int32(559),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	502: {
		Ffile: __ccgo_ts,
		Fline: int32(560),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	503: {
		Ffile: __ccgo_ts,
		Fline: int32(561),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	504: {
		Ffile: __ccgo_ts,
		Fline: int32(562),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	505: {
		Ffile: __ccgo_ts,
		Fline: int32(563),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	506: {
		Ffile: __ccgo_ts,
		Fline: int32(564),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	507: {
		Ffile: __ccgo_ts,
		Fline: int32(565),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	508: {
		Ffile: __ccgo_ts,
		Fline: int32(566),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	509: {
		Ffile: __ccgo_ts,
		Fline: int32(567),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	510: {
		Ffile: __ccgo_ts,
		Fline: int32(568),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	511: {
		Ffile: __ccgo_ts,
		Fline: int32(569),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	512: {
		Ffile: __ccgo_ts,
		Fline: int32(570),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	513: {
		Ffile: __ccgo_ts,
		Fline: int32(571),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	514: {
		Ffile: __ccgo_ts,
		Fline: int32(572),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	515: {
		Ffile: __ccgo_ts,
		Fline: int32(573),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	516: {
		Ffile: __ccgo_ts,
		Fline: int32(574),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	517: {
		Ffile: __ccgo_ts,
		Fline: int32(575),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	518: {
		Ffile: __ccgo_ts,
		Fline: int32(576),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	519: {
		Ffile: __ccgo_ts,
		Fline: int32(577),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	520: {
		Ffile: __ccgo_ts,
		Fline: int32(578),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	521: {
		Ffile: __ccgo_ts,
		Fline: int32(579),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	522: {
		Ffile: __ccgo_ts,
		Fline: int32(580),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	523: {
		Ffile: __ccgo_ts,
		Fline: int32(581),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	524: {
		Ffile: __ccgo_ts,
		Fline: int32(582),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	525: {
		Ffile: __ccgo_ts,
		Fline: int32(583),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	526: {
		Ffile: __ccgo_ts,
		Fline: int32(584),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	527: {
		Ffile: __ccgo_ts,
		Fline: int32(585),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	528: {
		Ffile: __ccgo_ts,
		Fline: int32(586),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	529: {
		Ffile: __ccgo_ts,
		Fline: int32(587),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	530: {
		Ffile: __ccgo_ts,
		Fline: int32(588),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	531: {
		Ffile: __ccgo_ts,
		Fline: int32(589),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	532: {
		Ffile: __ccgo_ts,
		Fline: int32(590),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	533: {
		Ffile: __ccgo_ts,
		Fline: int32(591),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	534: {
		Ffile: __ccgo_ts,
		Fline: int32(592),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	535: {
		Ffile: __ccgo_ts,
		Fline: int32(593),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	536: {
		Ffile: __ccgo_ts,
		Fline: int32(594),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(0.5),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	537: {
		Ffile: __ccgo_ts,
		Fline: int32(595),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(0.9999999999999999),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	538: {
		Ffile: __ccgo_ts,
		Fline: int32(596),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	539: {
		Ffile: __ccgo_ts,
		Fline: int32(597),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	540: {
		Ffile: __ccgo_ts,
		Fline: int32(598),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	541: {
		Ffile: __ccgo_ts,
		Fline: int32(599),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	542: {
		Ffile: __ccgo_ts,
		Fline: int32(600),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	543: {
		Ffile: __ccgo_ts,
		Fline: int32(601),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	544: {
		Ffile: __ccgo_ts,
		Fline: int32(602),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	545: {
		Ffile: __ccgo_ts,
		Fline: int32(603),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	546: {
		Ffile: __ccgo_ts,
		Fline: int32(604),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	547: {
		Ffile: __ccgo_ts,
		Fline: int32(605),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	548: {
		Ffile: __ccgo_ts,
		Fline: int32(606),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	549: {
		Ffile: __ccgo_ts,
		Fline: int32(607),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	550: {
		Ffile: __ccgo_ts,
		Fline: int32(608),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	551: {
		Ffile: __ccgo_ts,
		Fline: int32(609),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	552: {
		Ffile: __ccgo_ts,
		Fline: int32(610),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	553: {
		Ffile: __ccgo_ts,
		Fline: int32(611),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	554: {
		Ffile: __ccgo_ts,
		Fline: int32(612),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	555: {
		Ffile: __ccgo_ts,
		Fline: int32(613),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	556: {
		Ffile: __ccgo_ts,
		Fline: int32(614),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	557: {
		Ffile: __ccgo_ts,
		Fline: int32(615),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	558: {
		Ffile: __ccgo_ts,
		Fline: int32(616),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	559: {
		Ffile: __ccgo_ts,
		Fline: int32(617),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	560: {
		Ffile: __ccgo_ts,
		Fline: int32(618),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	561: {
		Ffile: __ccgo_ts,
		Fline: int32(619),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	562: {
		Ffile: __ccgo_ts,
		Fline: int32(620),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507202e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	563: {
		Ffile: __ccgo_ts,
		Fline: int32(621),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	564: {
		Ffile: __ccgo_ts,
		Fline: int32(622),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	565: {
		Ffile: __ccgo_ts,
		Fline: int32(623),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	566: {
		Ffile: __ccgo_ts,
		Fline: int32(624),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.225073858507203e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	567: {
		Ffile: __ccgo_ts,
		Fline: int32(625),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	568: {
		Ffile: __ccgo_ts,
		Fline: int32(626),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072034e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.2250738585072034e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	569: {
		Ffile: __ccgo_ts,
		Fline: int32(627),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072043e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	570: {
		Ffile: __ccgo_ts,
		Fline: int32(628),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.4501477170144023e-308),
		Fx2:   float64(4.450147717014403e-308),
		Fy:    float64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	571: {
		Ffile: __ccgo_ts,
		Fline: int32(629),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.139237815555687e-305),
		Fx2:   float64(5.696189077778436e-306),
		Fy:    float64(5.696189077778434e-306),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	572: {
		Ffile: __ccgo_ts,
		Fline: int32(630),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.125),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	573: {
		Ffile: __ccgo_ts,
		Fline: int32(631),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.25),
		Fx2:   float64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	574: {
		Ffile: __ccgo_ts,
		Fline: int32(632),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.25000000000000006),
		Fx2:   float64(0.5),
		Fy:    float64(0.25000000000000006),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	575: {
		Ffile: __ccgo_ts,
		Fline: int32(633),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.375),
		Fx2:   float64(0.5),
		Fy:    float64(0.375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	576: {
		Ffile: __ccgo_ts,
		Fline: int32(634),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.625),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	577: {
		Ffile: __ccgo_ts,
		Fline: int32(635),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999996),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	578: {
		Ffile: __ccgo_ts,
		Fline: int32(636),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	579: {
		Ffile: __ccgo_ts,
		Fline: int32(637),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	580: {
		Ffile: __ccgo_ts,
		Fline: int32(638),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	581: {
		Ffile: __ccgo_ts,
		Fline: int32(639),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	582: {
		Ffile: __ccgo_ts,
		Fline: int32(640),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	583: {
		Ffile: __ccgo_ts,
		Fline: int32(641),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	584: {
		Ffile: __ccgo_ts,
		Fline: int32(642),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	585: {
		Ffile: __ccgo_ts,
		Fline: int32(643),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	586: {
		Ffile: __ccgo_ts,
		Fline: int32(644),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	587: {
		Ffile: __ccgo_ts,
		Fline: int32(645),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(4.440892098500626e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	588: {
		Ffile: __ccgo_ts,
		Fline: int32(646),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	589: {
		Ffile: __ccgo_ts,
		Fline: int32(647),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(2),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	590: {
		Ffile: __ccgo_ts,
		Fline: int32(648),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	591: {
		Ffile: __ccgo_ts,
		Fline: int32(649),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000004),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	592: {
		Ffile: __ccgo_ts,
		Fline: int32(650),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000007),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000007),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	593: {
		Ffile: __ccgo_ts,
		Fline: int32(651),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000009),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000009),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	594: {
		Ffile: __ccgo_ts,
		Fline: int32(652),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000013),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	595: {
		Ffile: __ccgo_ts,
		Fline: int32(653),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.9999999999999998),
		Fx2:   float64(2),
		Fy:    float64(1.9999999999999998),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	596: {
		Ffile: __ccgo_ts,
		Fline: int32(654),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2),
		Fx2:   float64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	597: {
		Ffile: __ccgo_ts,
		Fline: int32(655),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	598: {
		Ffile: __ccgo_ts,
		Fline: int32(656),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	599: {
		Ffile: __ccgo_ts,
		Fline: int32(657),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	600: {
		Ffile: __ccgo_ts,
		Fline: int32(658),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	601: {
		Ffile: __ccgo_ts,
		Fline: int32(659),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.0000000000000036),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000036),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	602: {
		Ffile: __ccgo_ts,
		Fline: int32(660),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(2),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	603: {
		Ffile: __ccgo_ts,
		Fline: int32(661),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	604: {
		Ffile: __ccgo_ts,
		Fline: int32(662),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	605: {
		Ffile: __ccgo_ts,
		Fline: int32(663),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	606: {
		Ffile: __ccgo_ts,
		Fline: int32(664),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	607: {
		Ffile: __ccgo_ts,
		Fline: int32(665),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	608: {
		Ffile: __ccgo_ts,
		Fline: int32(666),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	609: {
		Ffile: __ccgo_ts,
		Fline: int32(667),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	610: {
		Ffile: __ccgo_ts,
		Fline: int32(668),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.999999999999993),
		Fx2:   float64(4),
		Fy:    float64(1.999999999999993),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	611: {
		Ffile: __ccgo_ts,
		Fline: int32(669),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	612: {
		Ffile: __ccgo_ts,
		Fline: int32(670),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	613: {
		Ffile: __ccgo_ts,
		Fline: int32(671),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	614: {
		Ffile: __ccgo_ts,
		Fline: int32(672),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(7),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	615: {
		Ffile: __ccgo_ts,
		Fline: int32(673),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	616: {
		Ffile: __ccgo_ts,
		Fline: int32(674),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.99584030953472e+292),
		Fx2:   float64(7.98336123813888e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	617: {
		Ffile: __ccgo_ts,
		Fline: int32(675),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.49423283715579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(4.49423283715579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	618: {
		Ffile: __ccgo_ts,
		Fline: int32(676),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311576e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(8.988465674311576e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	619: {
		Ffile: __ccgo_ts,
		Fline: int32(677),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	620: {
		Ffile: __ccgo_ts,
		Fline: int32(678),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	621: {
		Ffile: __ccgo_ts,
		Fline: int32(679),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	622: {
		Ffile: __ccgo_ts,
		Fline: int32(680),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	623: {
		Ffile: __ccgo_ts,
		Fline: int32(681),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	624: {
		Ffile: __ccgo_ts,
		Fline: int32(682),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	625: {
		Ffile: __ccgo_ts,
		Fline: int32(683),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	626: {
		Ffile: __ccgo_ts,
		Fline: int32(684),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.98846567431158e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	627: {
		Ffile: __ccgo_ts,
		Fline: int32(685),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	628: {
		Ffile: __ccgo_ts,
		Fline: int32(686),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311582e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(3.99168061906944e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	629: {
		Ffile: __ccgo_ts,
		Fline: int32(687),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	630: {
		Ffile: __ccgo_ts,
		Fline: int32(688),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	631: {
		Ffile: __ccgo_ts,
		Fline: int32(689),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	632: {
		Ffile: __ccgo_ts,
		Fline: int32(690),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	633: {
		Ffile: __ccgo_ts,
		Fline: int32(691),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	634: {
		Ffile: __ccgo_ts,
		Fline: int32(692),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311586e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	635: {
		Ffile: __ccgo_ts,
		Fline: int32(693),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	636: {
		Ffile: __ccgo_ts,
		Fline: int32(694),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311588e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311588e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	637: {
		Ffile: __ccgo_ts,
		Fline: int32(695),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.988465674311592e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	638: {
		Ffile: __ccgo_ts,
		Fline: int32(696),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	639: {
		Ffile: __ccgo_ts,
		Fline: int32(697),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	640: {
		Ffile: __ccgo_ts,
		Fline: int32(698),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	641: {
		Ffile: __ccgo_ts,
		Fline: int32(699),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	642: {
		Ffile: __ccgo_ts,
		Fline: int32(700),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	643: {
		Ffile: __ccgo_ts,
		Fline: int32(701),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	644: {
		Ffile: __ccgo_ts,
		Fline: int32(702),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	645: {
		Ffile: __ccgo_ts,
		Fline: int32(703),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	646: {
		Ffile: __ccgo_ts,
		Fline: int32(704),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	647: {
		Ffile: __ccgo_ts,
		Fline: int32(705),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	648: {
		Ffile: __ccgo_ts,
		Fline: int32(706),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	649: {
		Ffile: __ccgo_ts,
		Fline: int32(707),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	650: {
		Ffile: __ccgo_ts,
		Fline: int32(708),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	651: {
		Ffile: __ccgo_ts,
		Fline: int32(709),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	652: {
		Ffile: __ccgo_ts,
		Fline: int32(710),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	653: {
		Ffile: __ccgo_ts,
		Fline: int32(711),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	654: {
		Ffile: __ccgo_ts,
		Fline: int32(712),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	655: {
		Ffile: __ccgo_ts,
		Fline: int32(713),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	656: {
		Ffile: __ccgo_ts,
		Fline: int32(714),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	657: {
		Ffile: __ccgo_ts,
		Fline: int32(715),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	658: {
		Ffile: __ccgo_ts,
		Fline: int32(716),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	659: {
		Ffile: __ccgo_ts,
		Fline: int32(717),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	660: {
		Ffile: __ccgo_ts,
		Fline: int32(718),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	661: {
		Ffile: __ccgo_ts,
		Fline: int32(719),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	662: {
		Ffile: __ccgo_ts,
		Fline: int32(720),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	663: {
		Ffile: __ccgo_ts,
		Fline: int32(721),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	664: {
		Ffile: __ccgo_ts,
		Fline: int32(722),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	665: {
		Ffile: __ccgo_ts,
		Fline: int32(723),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	666: {
		Ffile: __ccgo_ts,
		Fline: int32(724),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	667: {
		Ffile: __ccgo_ts,
		Fline: int32(725),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	668: {
		Ffile: __ccgo_ts,
		Fline: int32(726),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	669: {
		Ffile: __ccgo_ts,
		Fline: int32(727),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	670: {
		Ffile: __ccgo_ts,
		Fline: int32(728),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	671: {
		Ffile: __ccgo_ts,
		Fline: int32(729),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	672: {
		Ffile: __ccgo_ts,
		Fline: int32(730),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	673: {
		Ffile: __ccgo_ts,
		Fline: int32(731),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	674: {
		Ffile: __ccgo_ts,
		Fline: int32(732),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	675: {
		Ffile: __ccgo_ts,
		Fline: int32(733),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	676: {
		Ffile: __ccgo_ts,
		Fline: int32(734),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	677: {
		Ffile: __ccgo_ts,
		Fline: int32(735),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	678: {
		Ffile: __ccgo_ts,
		Fline: int32(736),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	679: {
		Ffile: __ccgo_ts,
		Fline: int32(737),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	680: {
		Ffile: __ccgo_ts,
		Fline: int32(738),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	681: {
		Ffile: __ccgo_ts,
		Fline: int32(739),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	682: {
		Ffile: __ccgo_ts,
		Fline: int32(740),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	683: {
		Ffile: __ccgo_ts,
		Fline: int32(741),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	684: {
		Ffile: __ccgo_ts,
		Fline: int32(742),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	685: {
		Ffile: __ccgo_ts,
		Fline: int32(743),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	686: {
		Ffile: __ccgo_ts,
		Fline: int32(744),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	687: {
		Ffile: __ccgo_ts,
		Fline: int32(745),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	688: {
		Ffile: __ccgo_ts,
		Fline: int32(746),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	689: {
		Ffile: __ccgo_ts,
		Fline: int32(747),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	690: {
		Ffile: __ccgo_ts,
		Fline: int32(748),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	691: {
		Ffile: __ccgo_ts,
		Fline: int32(749),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	692: {
		Ffile: __ccgo_ts,
		Fline: int32(750),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	693: {
		Ffile: __ccgo_ts,
		Fline: int32(751),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	694: {
		Ffile: __ccgo_ts,
		Fline: int32(752),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	695: {
		Ffile: __ccgo_ts,
		Fline: int32(753),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	696: {
		Ffile: __ccgo_ts,
		Fline: int32(754),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	697: {
		Ffile: __ccgo_ts,
		Fline: int32(755),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	698: {
		Ffile: __ccgo_ts,
		Fline: int32(756),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	699: {
		Ffile: __ccgo_ts,
		Fline: int32(757),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	700: {
		Ffile: __ccgo_ts,
		Fline: int32(758),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	701: {
		Ffile: __ccgo_ts,
		Fline: int32(759),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	702: {
		Ffile: __ccgo_ts,
		Fline: int32(760),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	703: {
		Ffile: __ccgo_ts,
		Fline: int32(761),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	704: {
		Ffile: __ccgo_ts,
		Fline: int32(762),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	705: {
		Ffile: __ccgo_ts,
		Fline: int32(763),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	706: {
		Ffile: __ccgo_ts,
		Fline: int32(764),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	707: {
		Ffile: __ccgo_ts,
		Fline: int32(765),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	708: {
		Ffile: __ccgo_ts,
		Fline: int32(766),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	709: {
		Ffile: __ccgo_ts,
		Fline: int32(767),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	710: {
		Ffile: __ccgo_ts,
		Fline: int32(768),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	711: {
		Ffile: __ccgo_ts,
		Fline: int32(769),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	712: {
		Ffile: __ccgo_ts,
		Fline: int32(770),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	713: {
		Ffile: __ccgo_ts,
		Fline: int32(771),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	714: {
		Ffile: __ccgo_ts,
		Fline: int32(772),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	715: {
		Ffile: __ccgo_ts,
		Fline: int32(773),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	716: {
		Ffile: __ccgo_ts,
		Fline: int32(774),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	717: {
		Ffile: __ccgo_ts,
		Fline: int32(775),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	718: {
		Ffile: __ccgo_ts,
		Fline: int32(776),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	719: {
		Ffile: __ccgo_ts,
		Fline: int32(777),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	720: {
		Ffile: __ccgo_ts,
		Fline: int32(778),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	721: {
		Ffile: __ccgo_ts,
		Fline: int32(779),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	722: {
		Ffile: __ccgo_ts,
		Fline: int32(780),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	723: {
		Ffile: __ccgo_ts,
		Fline: int32(781),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	724: {
		Ffile: __ccgo_ts,
		Fline: int32(782),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	725: {
		Ffile: __ccgo_ts,
		Fline: int32(783),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	726: {
		Ffile: __ccgo_ts,
		Fline: int32(784),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	727: {
		Ffile: __ccgo_ts,
		Fline: int32(785),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	728: {
		Ffile: __ccgo_ts,
		Fline: int32(786),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	729: {
		Ffile: __ccgo_ts,
		Fline: int32(787),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	730: {
		Ffile: __ccgo_ts,
		Fline: int32(788),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	731: {
		Ffile: __ccgo_ts,
		Fline: int32(789),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	732: {
		Ffile: __ccgo_ts,
		Fline: int32(790),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	733: {
		Ffile: __ccgo_ts,
		Fline: int32(791),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	734: {
		Ffile: __ccgo_ts,
		Fline: int32(792),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	735: {
		Ffile: __ccgo_ts,
		Fline: int32(793),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	736: {
		Ffile: __ccgo_ts,
		Fline: int32(794),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	737: {
		Ffile: __ccgo_ts,
		Fline: int32(795),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	738: {
		Ffile: __ccgo_ts,
		Fline: int32(796),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	739: {
		Ffile: __ccgo_ts,
		Fline: int32(797),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	740: {
		Ffile: __ccgo_ts,
		Fline: int32(798),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	741: {
		Ffile: __ccgo_ts,
		Fline: int32(799),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	742: {
		Ffile: __ccgo_ts,
		Fline: int32(800),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	743: {
		Ffile: __ccgo_ts,
		Fline: int32(801),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	744: {
		Ffile: __ccgo_ts,
		Fline: int32(802),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	745: {
		Ffile: __ccgo_ts,
		Fline: int32(803),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	746: {
		Ffile: __ccgo_ts,
		Fline: int32(804),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	747: {
		Ffile: __ccgo_ts,
		Fline: int32(805),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	748: {
		Ffile: __ccgo_ts,
		Fline: int32(806),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	749: {
		Ffile: __ccgo_ts,
		Fline: int32(807),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	750: {
		Ffile: __ccgo_ts,
		Fline: int32(808),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	751: {
		Ffile: __ccgo_ts,
		Fline: int32(809),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	752: {
		Ffile: __ccgo_ts,
		Fline: int32(810),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	753: {
		Ffile: __ccgo_ts,
		Fline: int32(811),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	754: {
		Ffile: __ccgo_ts,
		Fline: int32(812),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	755: {
		Ffile: __ccgo_ts,
		Fline: int32(813),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	756: {
		Ffile: __ccgo_ts,
		Fline: int32(814),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(0.5),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	757: {
		Ffile: __ccgo_ts,
		Fline: int32(815),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(0.9999999999999999),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	758: {
		Ffile: __ccgo_ts,
		Fline: int32(816),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	759: {
		Ffile: __ccgo_ts,
		Fline: int32(817),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	760: {
		Ffile: __ccgo_ts,
		Fline: int32(818),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	761: {
		Ffile: __ccgo_ts,
		Fline: int32(819),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	762: {
		Ffile: __ccgo_ts,
		Fline: int32(820),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	763: {
		Ffile: __ccgo_ts,
		Fline: int32(821),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	764: {
		Ffile: __ccgo_ts,
		Fline: int32(822),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	765: {
		Ffile: __ccgo_ts,
		Fline: int32(823),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	766: {
		Ffile: __ccgo_ts,
		Fline: int32(824),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	767: {
		Ffile: __ccgo_ts,
		Fline: int32(825),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	768: {
		Ffile: __ccgo_ts,
		Fline: int32(826),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	769: {
		Ffile: __ccgo_ts,
		Fline: int32(827),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	770: {
		Ffile: __ccgo_ts,
		Fline: int32(828),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	771: {
		Ffile: __ccgo_ts,
		Fline: int32(829),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	772: {
		Ffile: __ccgo_ts,
		Fline: int32(830),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	773: {
		Ffile: __ccgo_ts,
		Fline: int32(831),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	774: {
		Ffile: __ccgo_ts,
		Fline: int32(832),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	775: {
		Ffile: __ccgo_ts,
		Fline: int32(833),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	776: {
		Ffile: __ccgo_ts,
		Fline: int32(834),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	777: {
		Ffile: __ccgo_ts,
		Fline: int32(835),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	778: {
		Ffile: __ccgo_ts,
		Fline: int32(836),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	779: {
		Ffile: __ccgo_ts,
		Fline: int32(837),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	780: {
		Ffile: __ccgo_ts,
		Fline: int32(838),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	781: {
		Ffile: __ccgo_ts,
		Fline: int32(839),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	782: {
		Ffile: __ccgo_ts,
		Fline: int32(840),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	783: {
		Ffile: __ccgo_ts,
		Fline: int32(841),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507202e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    float64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	784: {
		Ffile: __ccgo_ts,
		Fline: int32(842),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	785: {
		Ffile: __ccgo_ts,
		Fline: int32(843),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	786: {
		Ffile: __ccgo_ts,
		Fline: int32(844),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	787: {
		Ffile: __ccgo_ts,
		Fline: int32(845),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.225073858507203e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	788: {
		Ffile: __ccgo_ts,
		Fline: int32(846),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	789: {
		Ffile: __ccgo_ts,
		Fline: int32(847),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072034e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(2.2250738585072034e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	790: {
		Ffile: __ccgo_ts,
		Fline: int32(848),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072043e-308),
		Fx2:   float64(2.225073858507204e-308),
		Fy:    float64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	791: {
		Ffile: __ccgo_ts,
		Fline: int32(849),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.4501477170144023e-308),
		Fx2:   float64(4.450147717014403e-308),
		Fy:    float64(4.4501477170144023e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	792: {
		Ffile: __ccgo_ts,
		Fline: int32(850),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.139237815555687e-305),
		Fx2:   float64(5.696189077778436e-306),
		Fy:    float64(5.696189077778434e-306),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	793: {
		Ffile: __ccgo_ts,
		Fline: int32(851),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.125),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	794: {
		Ffile: __ccgo_ts,
		Fline: int32(852),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.25),
		Fx2:   float64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	795: {
		Ffile: __ccgo_ts,
		Fline: int32(853),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.25000000000000006),
		Fx2:   float64(0.5),
		Fy:    float64(0.25000000000000006),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	796: {
		Ffile: __ccgo_ts,
		Fline: int32(854),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.375),
		Fx2:   float64(0.5),
		Fy:    float64(0.375),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	797: {
		Ffile: __ccgo_ts,
		Fline: int32(855),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.625),
		Fx2:   float64(0.5),
		Fy:    float64(0.125),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	798: {
		Ffile: __ccgo_ts,
		Fline: int32(856),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999996),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	799: {
		Ffile: __ccgo_ts,
		Fline: int32(857),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	800: {
		Ffile: __ccgo_ts,
		Fline: int32(858),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0.9999999999999999),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	801: {
		Ffile: __ccgo_ts,
		Fline: int32(859),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	802: {
		Ffile: __ccgo_ts,
		Fline: int32(860),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(1.9999999999999998),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	803: {
		Ffile: __ccgo_ts,
		Fline: int32(861),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	804: {
		Ffile: __ccgo_ts,
		Fline: int32(862),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	805: {
		Ffile: __ccgo_ts,
		Fline: int32(863),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	806: {
		Ffile: __ccgo_ts,
		Fline: int32(864),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	807: {
		Ffile: __ccgo_ts,
		Fline: int32(865),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	808: {
		Ffile: __ccgo_ts,
		Fline: int32(866),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(0.9999999999999998),
		Fy:    float64(4.440892098500626e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	809: {
		Ffile: __ccgo_ts,
		Fline: int32(867),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	810: {
		Ffile: __ccgo_ts,
		Fline: int32(868),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fx2:   float64(2),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	811: {
		Ffile: __ccgo_ts,
		Fline: int32(869),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    float64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	812: {
		Ffile: __ccgo_ts,
		Fline: int32(870),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000004),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	813: {
		Ffile: __ccgo_ts,
		Fline: int32(871),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000007),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000007),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	814: {
		Ffile: __ccgo_ts,
		Fline: int32(872),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000009),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(1.0000000000000009),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	815: {
		Ffile: __ccgo_ts,
		Fline: int32(873),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000013),
		Fx2:   float64(1.000000000000001),
		Fy:    float64(2.220446049250313e-16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	816: {
		Ffile: __ccgo_ts,
		Fline: int32(874),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.9999999999999998),
		Fx2:   float64(2),
		Fy:    float64(1.9999999999999998),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	817: {
		Ffile: __ccgo_ts,
		Fline: int32(875),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2),
		Fx2:   float64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	818: {
		Ffile: __ccgo_ts,
		Fline: int32(876),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	819: {
		Ffile: __ccgo_ts,
		Fline: int32(877),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	820: {
		Ffile: __ccgo_ts,
		Fline: int32(878),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	821: {
		Ffile: __ccgo_ts,
		Fline: int32(879),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	822: {
		Ffile: __ccgo_ts,
		Fline: int32(880),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.0000000000000036),
		Fx2:   float64(4),
		Fy:    float64(2.0000000000000036),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	823: {
		Ffile: __ccgo_ts,
		Fline: int32(881),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(2),
		Fy:    float64(0.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	824: {
		Ffile: __ccgo_ts,
		Fline: int32(882),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	825: {
		Ffile: __ccgo_ts,
		Fline: int32(883),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	826: {
		Ffile: __ccgo_ts,
		Fline: int32(884),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3),
		Fx2:   float64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	827: {
		Ffile: __ccgo_ts,
		Fline: int32(885),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	828: {
		Ffile: __ccgo_ts,
		Fline: int32(886),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	829: {
		Ffile: __ccgo_ts,
		Fline: int32(887),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	830: {
		Ffile: __ccgo_ts,
		Fline: int32(888),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5),
		Fx2:   float64(4),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	831: {
		Ffile: __ccgo_ts,
		Fline: int32(889),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.999999999999993),
		Fx2:   float64(4),
		Fy:    float64(1.999999999999993),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	832: {
		Ffile: __ccgo_ts,
		Fline: int32(890),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	833: {
		Ffile: __ccgo_ts,
		Fline: int32(891),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    float64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	834: {
		Ffile: __ccgo_ts,
		Fline: int32(892),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6),
		Fx2:   float64(4),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	835: {
		Ffile: __ccgo_ts,
		Fline: int32(893),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7),
		Fx2:   float64(4),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	836: {
		Ffile: __ccgo_ts,
		Fline: int32(894),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8),
		Fx2:   float64(4),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	837: {
		Ffile: __ccgo_ts,
		Fline: int32(895),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.99584030953472e+292),
		Fx2:   float64(7.98336123813888e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	838: {
		Ffile: __ccgo_ts,
		Fline: int32(896),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.49423283715579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(4.49423283715579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	839: {
		Ffile: __ccgo_ts,
		Fline: int32(897),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311576e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(8.988465674311576e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	840: {
		Ffile: __ccgo_ts,
		Fline: int32(898),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	841: {
		Ffile: __ccgo_ts,
		Fline: int32(899),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	842: {
		Ffile: __ccgo_ts,
		Fline: int32(900),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	843: {
		Ffile: __ccgo_ts,
		Fline: int32(901),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	844: {
		Ffile: __ccgo_ts,
		Fline: int32(902),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	845: {
		Ffile: __ccgo_ts,
		Fline: int32(903),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	846: {
		Ffile: __ccgo_ts,
		Fline: int32(904),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	847: {
		Ffile: __ccgo_ts,
		Fline: int32(905),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(8.98846567431158e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	848: {
		Ffile: __ccgo_ts,
		Fline: int32(906),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	849: {
		Ffile: __ccgo_ts,
		Fline: int32(907),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311582e+307),
		Fx2:   float64(8.988465674311578e+307),
		Fy:    float64(3.99168061906944e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	850: {
		Ffile: __ccgo_ts,
		Fline: int32(908),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	851: {
		Ffile: __ccgo_ts,
		Fline: int32(909),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	852: {
		Ffile: __ccgo_ts,
		Fline: int32(910),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	853: {
		Ffile: __ccgo_ts,
		Fline: int32(911),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	854: {
		Ffile: __ccgo_ts,
		Fline: int32(912),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	855: {
		Ffile: __ccgo_ts,
		Fline: int32(913),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311586e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	856: {
		Ffile: __ccgo_ts,
		Fline: int32(914),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	857: {
		Ffile: __ccgo_ts,
		Fline: int32(915),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311588e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(8.988465674311588e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	858: {
		Ffile: __ccgo_ts,
		Fline: int32(916),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.988465674311592e+307),
		Fx2:   float64(8.98846567431159e+307),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	859: {
		Ffile: __ccgo_ts,
		Fline: int32(917),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	860: {
		Ffile: __ccgo_ts,
		Fline: int32(918),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(2.2250738585072014e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	861: {
		Ffile: __ccgo_ts,
		Fline: int32(919),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(4.4501477170144023e-308),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	862: {
		Ffile: __ccgo_ts,
		Fline: int32(920),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    float64(8.988465674311578e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	863: {
		Ffile: __ccgo_ts,
		Fline: int32(921),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	864: {
		Ffile: __ccgo_ts,
		Fline: int32(922),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	865: {
		Ffile: __ccgo_ts,
		Fline: int32(923),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	866: {
		Ffile: __ccgo_ts,
		Fline: int32(924),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	867: {
		Ffile: __ccgo_ts,
		Fline: int32(925),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	868: {
		Ffile: __ccgo_ts,
		Fline: int32(926),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	869: {
		Ffile: __ccgo_ts,
		Fline: int32(927),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	870: {
		Ffile: __ccgo_ts,
		Fline: int32(928),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	871: {
		Ffile: __ccgo_ts,
		Fline: int32(929),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	872: {
		Ffile: __ccgo_ts,
		Fline: int32(930),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	873: {
		Ffile: __ccgo_ts,
		Fline: int32(931),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	874: {
		Ffile: __ccgo_ts,
		Fline: int32(932),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	875: {
		Ffile: __ccgo_ts,
		Fline: int32(933),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	876: {
		Ffile: __ccgo_ts,
		Fline: int32(934),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	877: {
		Ffile: __ccgo_ts,
		Fline: int32(935),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	878: {
		Ffile: __ccgo_ts,
		Fline: int32(936),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   float64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	879: {
		Ffile: __ccgo_ts,
		Fline: int32(937),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2e-323),
		Fx2:   -libc.Float64FromFloat64(2e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	880: {
		Ffile: __ccgo_ts,
		Fline: int32(938),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   float64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	881: {
		Ffile: __ccgo_ts,
		Fline: int32(939),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fx2:   -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    -libc.Float64FromFloat64(2.2250738585071994e-308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	882: {
		Ffile: __ccgo_ts,
		Fline: int32(940),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	883: {
		Ffile: __ccgo_ts,
		Fline: int32(941),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	884: {
		Ffile: __ccgo_ts,
		Fline: int32(942),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	885: {
		Ffile: __ccgo_ts,
		Fline: int32(943),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	886: {
		Ffile: __ccgo_ts,
		Fline: int32(944),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	887: {
		Ffile: __ccgo_ts,
		Fline: int32(945),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	888: {
		Ffile: __ccgo_ts,
		Fline: int32(946),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507203e-308),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	889: {
		Ffile: __ccgo_ts,
		Fline: int32(947),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	890: {
		Ffile: __ccgo_ts,
		Fline: int32(948),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	891: {
		Ffile: __ccgo_ts,
		Fline: int32(949),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	892: {
		Ffile: __ccgo_ts,
		Fline: int32(950),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	893: {
		Ffile: __ccgo_ts,
		Fline: int32(951),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   float64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	894: {
		Ffile: __ccgo_ts,
		Fline: int32(952),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fx2:   -libc.Float64FromFloat64(1.000000000000001),
		Fy:    -libc.Float64FromFloat64(1.0000000000000002),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	895: {
		Ffile: __ccgo_ts,
		Fline: int32(953),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	896: {
		Ffile: __ccgo_ts,
		Fline: int32(954),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	897: {
		Ffile: __ccgo_ts,
		Fline: int32(955),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	898: {
		Ffile: __ccgo_ts,
		Fline: int32(956),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.0000000000000004),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(2.0000000000000004),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	899: {
		Ffile: __ccgo_ts,
		Fline: int32(957),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	900: {
		Ffile: __ccgo_ts,
		Fline: int32(958),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.9999999999999996),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(2.9999999999999996),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	901: {
		Ffile: __ccgo_ts,
		Fline: int32(959),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   float64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	902: {
		Ffile: __ccgo_ts,
		Fline: int32(960),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3),
		Fx2:   -libc.Float64FromFloat64(2),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	903: {
		Ffile: __ccgo_ts,
		Fline: int32(961),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   float64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	904: {
		Ffile: __ccgo_ts,
		Fline: int32(962),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.999999999999999),
		Fx2:   -libc.Float64FromFloat64(4),
		Fy:    -libc.Float64FromFloat64(1.9999999999999991),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	905: {
		Ffile: __ccgo_ts,
		Fline: int32(963),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	906: {
		Ffile: __ccgo_ts,
		Fline: int32(964),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311578e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	907: {
		Ffile: __ccgo_ts,
		Fline: int32(965),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   float64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	908: {
		Ffile: __ccgo_ts,
		Fline: int32(966),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fx2:   -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    -libc.Float64FromFloat64(8.988465674311579e+307),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	909: {
		Ffile: __ccgo_ts,
		Fline: int32(967),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	910: {
		Ffile: __ccgo_ts,
		Fline: int32(968),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	911: {
		Ffile: __ccgo_ts,
		Fline: int32(969),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	912: {
		Ffile: __ccgo_ts,
		Fline: int32(970),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   float64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	913: {
		Ffile: __ccgo_ts,
		Fline: int32(971),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(1.5e-323),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	914: {
		Ffile: __ccgo_ts,
		Fline: int32(972),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fx2:   -libc.Float64FromFloat64(3),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	915: {
		Ffile: __ccgo_ts,
		Fline: int32(973),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   float64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	916: {
		Ffile: __ccgo_ts,
		Fline: int32(974),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311586e+307),
		Fx2:   -libc.Float64FromFloat64(5.987520928604159e+292),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	917: {
		Ffile: __ccgo_ts,
		Fline: int32(975),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	918: {
		Ffile: __ccgo_ts,
		Fline: int32(976),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	919: {
		Ffile: __ccgo_ts,
		Fline: int32(977),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	920: {
		Ffile: __ccgo_ts,
		Fline: int32(978),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	921: {
		Ffile: __ccgo_ts,
		Fline: int32(979),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    -libc.Float64FromFloat64(1.99584030953472e+292),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	922: {
		Ffile: __ccgo_ts,
		Fline: int32(980),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	923: {
		Ffile: __ccgo_ts,
		Fline: int32(981),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	924: {
		Ffile: __ccgo_ts,
		Fline: int32(982),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	925: {
		Ffile: __ccgo_ts,
		Fline: int32(983),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	926: {
		Ffile: __ccgo_ts,
		Fline: int32(984),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	927: {
		Ffile: __ccgo_ts,
		Fline: int32(985),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	928: {
		Ffile: __ccgo_ts,
		Fline: int32(986),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	929: {
		Ffile: __ccgo_ts,
		Fline: int32(987),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	930: {
		Ffile: __ccgo_ts,
		Fline: int32(988),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	931: {
		Ffile: __ccgo_ts,
		Fline: int32(989),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	932: {
		Ffile: __ccgo_ts,
		Fline: int32(990),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	933: {
		Ffile: __ccgo_ts,
		Fline: int32(991),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	934: {
		Ffile: __ccgo_ts,
		Fline: int32(992),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	935: {
		Ffile: __ccgo_ts,
		Fline: int32(993),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	936: {
		Ffile: __ccgo_ts,
		Fline: int32(994),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	937: {
		Ffile: __ccgo_ts,
		Fline: int32(995),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	938: {
		Ffile: __ccgo_ts,
		Fline: int32(996),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	939: {
		Ffile: __ccgo_ts,
		Fline: int32(997),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	940: {
		Ffile: __ccgo_ts,
		Fline: int32(998),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	941: {
		Ffile: __ccgo_ts,
		Fline: int32(999),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	942: {
		Ffile: __ccgo_ts,
		Fline: int32(1000),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	943: {
		Ffile: __ccgo_ts,
		Fline: int32(1001),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	944: {
		Ffile: __ccgo_ts,
		Fline: int32(1002),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	945: {
		Ffile: __ccgo_ts,
		Fline: int32(1003),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	946: {
		Ffile: __ccgo_ts,
		Fline: int32(1004),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	947: {
		Ffile: __ccgo_ts,
		Fline: int32(1005),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	948: {
		Ffile: __ccgo_ts,
		Fline: int32(1006),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	949: {
		Ffile: __ccgo_ts,
		Fline: int32(1007),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	950: {
		Ffile: __ccgo_ts,
		Fline: int32(1008),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	951: {
		Ffile: __ccgo_ts,
		Fline: int32(1009),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	952: {
		Ffile: __ccgo_ts,
		Fline: int32(1010),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	953: {
		Ffile: __ccgo_ts,
		Fline: int32(1011),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999999),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	954: {
		Ffile: __ccgo_ts,
		Fline: int32(1012),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	955: {
		Ffile: __ccgo_ts,
		Fline: int32(1013),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	956: {
		Ffile: __ccgo_ts,
		Fline: int32(1014),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	957: {
		Ffile: __ccgo_ts,
		Fline: int32(1015),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	958: {
		Ffile: __ccgo_ts,
		Fline: int32(1016),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	959: {
		Ffile: __ccgo_ts,
		Fline: int32(1017),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	960: {
		Ffile: __ccgo_ts,
		Fline: int32(1018),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	961: {
		Ffile: __ccgo_ts,
		Fline: int32(1019),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	962: {
		Ffile: __ccgo_ts,
		Fline: int32(1020),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	963: {
		Ffile: __ccgo_ts,
		Fline: int32(1021),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	964: {
		Ffile: __ccgo_ts,
		Fline: int32(1022),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	965: {
		Ffile: __ccgo_ts,
		Fline: int32(1023),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	966: {
		Ffile: __ccgo_ts,
		Fline: int32(1024),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	967: {
		Ffile: __ccgo_ts,
		Fline: int32(1025),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	968: {
		Ffile: __ccgo_ts,
		Fline: int32(1026),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	969: {
		Ffile: __ccgo_ts,
		Fline: int32(1027),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	970: {
		Ffile: __ccgo_ts,
		Fline: int32(1028),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	971: {
		Ffile: __ccgo_ts,
		Fline: int32(1029),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	972: {
		Ffile: __ccgo_ts,
		Fline: int32(1030),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	973: {
		Ffile: __ccgo_ts,
		Fline: int32(1031),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	974: {
		Ffile: __ccgo_ts,
		Fline: int32(1032),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	975: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fx2:   float64(4.535662560676869),
		Fy:    -libc.Float64FromFloat64(3.531185829902812),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	976: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.345239849338305),
		Fx2:   -libc.Float64FromFloat64(8.88799136300345),
		Fy:    float64(4.345239849338305),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	977: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fx2:   -libc.Float64FromFloat64(2.763607337379588),
		Fy:    -libc.Float64FromFloat64(0.09061141541648476),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	978: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fx2:   float64(4.567535276842744),
		Fy:    -libc.Float64FromFloat64(1.9641383050707404),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	979: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.267056966972586),
		Fx2:   float64(4.811392084359796),
		Fy:    float64(4.45566488261279),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	980: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.450045556060236),
		Fx2:   float64(0.6620717923376739),
		Fy:    -libc.Float64FromFloat64(0.4913994250211714),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	981: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.858890253041697),
		Fx2:   float64(0.05215452675006225),
		Fy:    float64(0.035711240532359426),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	982: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.792054511984896),
		Fx2:   float64(7.67640268511754),
		Fy:    -libc.Float64FromFloat64(0.792054511984896),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	983: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.615702673197924),
		Fx2:   float64(2.0119025790324803),
		Fy:    float64(0.615702673197924),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	984: {
		Ffile: __ccgo_ts + 21,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5587586823609152),
		Fx2:   float64(0.03223983060263804),
		Fy:    -libc.Float64FromFloat64(0.0106815621160685),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	985: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	986: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	987: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	988: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	989: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	990: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	991: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5),
		Fx2:   float64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	992: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	993: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fx2:   float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	994: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   float64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	995: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	996: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	997: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	998: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(14),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	999: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(15),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1000: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(16),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1001: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(17),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1002: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(18),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1003: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1004: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1005: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(21),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0.5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1006: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(22),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1007: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(23),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1008: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(24),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1009: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(25),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1010: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(26),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1011: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(27),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1012: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1013: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(29),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1014: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(30),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1015: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(31),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1016: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(32),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1017: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(33),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1018: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(34),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1019: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1020: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(36),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1021: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(37),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1022: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1023: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1024: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(40),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1025: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1026: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1027: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1028: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1029: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   -libc.Float64FromFloat64(0),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1030: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1031: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1032: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1033: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1034: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1035: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1036: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1037: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1038: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1039: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1040: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1041: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1042: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1043: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1044: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    -libc.Float64FromFloat64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1045: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1046: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fx2:   float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+20)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	1047: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.75),
		Fx2:   float64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1048: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.75),
		Fx2:   float64(0.5),
		Fy:    -libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1049: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.75),
		Fx2:   -libc.Float64FromFloat64(0.5),
		Fy:    float64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1050: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.75),
		Fx2:   -libc.Float64FromFloat64(0.5),
		Fy:    -libc.Float64FromFloat64(0.25),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:12:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:13:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:17:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:20:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(46244)/libc.Uint32FromInt64(44)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:21:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*44
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:23:3:
		if (*dd_d)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:24:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:25:3:
		libc.Xfesetround(tls, (*dd_d)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:26:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:27:3:
		y = libc.Xfmod(tls, (*dd_d)(unsafe.Pointer(p)).Fx, (*dd_d)(unsafe.Pointer(p)).Fx2)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:28:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:30:3:
		if !(checkexcept(tls, e, (*dd_d)(unsafe.Pointer(p)).Fe, (*dd_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:31:4:
			libc.Xprintf(tls, __ccgo_ts+68, libc.VaList(bp+8, (*dd_d)(unsafe.Pointer(p)).Ffile, (*dd_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*dd_d)(unsafe.Pointer(p)).Fr), (*dd_d)(unsafe.Pointer(p)).Fx, (*dd_d)(unsafe.Pointer(p)).Fx2, (*dd_d)(unsafe.Pointer(p)).Fy, estr(tls, (*dd_d)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:33:4:
			libc.Xprintf(tls, __ccgo_ts+120, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:34:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:36:3:
		d = ulperr(tls, y, (*dd_d)(unsafe.Pointer(p)).Fy, (*dd_d)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:37:3:
		if !(checkcr(tls, float64(y), (*dd_d)(unsafe.Pointer(p)).Fy, (*dd_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:38:4:
			libc.Xprintf(tls, __ccgo_ts+129, libc.VaList(bp+8, (*dd_d)(unsafe.Pointer(p)).Ffile, (*dd_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*dd_d)(unsafe.Pointer(p)).Fr), (*dd_d)(unsafe.Pointer(p)).Fx, (*dd_d)(unsafe.Pointer(p)).Fx2, (*dd_d)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*dd_d)(unsafe.Pointer(p)).Fdy), float64((*dd_d)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:40:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/fmod.c:43:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+189, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+199, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 243,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 251,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 259,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 269,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 279,
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
				v2 = __ccgo_ts + 288
			} else {
				v2 = __ccgo_ts + 20
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+290, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 288
		} else {
			v3 = __ccgo_ts + 20
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+295, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 20
	} else {
		v4 = __ccgo_ts + 300
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+302, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 305
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 308
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 311
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 314
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 317
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+320, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+328, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+333, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+376, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+425) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+433) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+445) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+457) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+469) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+478) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+20) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+478) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+484, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/ucb/fmod.h\x00\x00src/math/sanity/fmod.h\x00src/math/special/fmod.h\x00%s:%d: bad fp exception: %s fmod(%a,%a)=%a, want %s\x00 got %s\n\x00%s:%d: %s fmod(%a,%a) want %a got %a ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
