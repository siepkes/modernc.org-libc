// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/log10.exe.go src/math/log10.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:5:19:
var t = [398]d_d{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(15),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(16),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(17),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(18),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(27),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(28),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(29),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(30),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(31),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(32),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(33),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(100000),
		Fy:    float64(5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(34),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(100000),
		Fy:    float64(5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(100000),
		Fy:    float64(5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(100000),
		Fy:    float64(5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+06),
		Fy:    float64(6),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+06),
		Fy:    float64(6),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+06),
		Fy:    float64(6),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+06),
		Fy:    float64(6),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+07),
		Fy:    float64(7),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+07),
		Fy:    float64(7),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+07),
		Fy:    float64(7),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+07),
		Fy:    float64(7),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+08),
		Fy:    float64(8),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+08),
		Fy:    float64(8),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+08),
		Fy:    float64(8),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+08),
		Fy:    float64(8),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+09),
		Fy:    float64(9),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+09),
		Fy:    float64(9),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+09),
		Fy:    float64(9),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+09),
		Fy:    float64(9),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+10),
		Fy:    float64(10),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+10),
		Fy:    float64(10),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+10),
		Fy:    float64(10),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+10),
		Fy:    float64(10),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+11),
		Fy:    float64(11),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+11),
		Fy:    float64(11),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+11),
		Fy:    float64(11),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+11),
		Fy:    float64(11),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+12),
		Fy:    float64(12),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+12),
		Fy:    float64(12),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+12),
		Fy:    float64(12),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+12),
		Fy:    float64(12),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+13),
		Fy:    float64(13),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+13),
		Fy:    float64(13),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+13),
		Fy:    float64(13),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+13),
		Fy:    float64(13),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+14),
		Fy:    float64(14),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+14),
		Fy:    float64(14),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+14),
		Fy:    float64(14),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+14),
		Fy:    float64(14),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+15),
		Fy:    float64(15),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+15),
		Fy:    float64(15),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+15),
		Fy:    float64(15),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+15),
		Fy:    float64(15),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+16),
		Fy:    float64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+16),
		Fy:    float64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+16),
		Fy:    float64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+16),
		Fy:    float64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+17),
		Fy:    float64(17),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+17),
		Fy:    float64(17),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+17),
		Fy:    float64(17),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+17),
		Fy:    float64(17),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+18),
		Fy:    float64(18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+18),
		Fy:    float64(18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+18),
		Fy:    float64(18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+18),
		Fy:    float64(18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+19),
		Fy:    float64(19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+19),
		Fy:    float64(19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+19),
		Fy:    float64(19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+19),
		Fy:    float64(19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+20),
		Fy:    float64(20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+20),
		Fy:    float64(20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+20),
		Fy:    float64(20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+20),
		Fy:    float64(20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+21),
		Fy:    float64(21),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+21),
		Fy:    float64(21),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+21),
		Fy:    float64(21),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+21),
		Fy:    float64(21),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+22),
		Fy:    float64(22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+22),
		Fy:    float64(22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+22),
		Fy:    float64(22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+22),
		Fy:    float64(22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+23),
		Fy:    float64(23),
		Fdy:   float32(0.010254488326609135),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+23),
		Fy:    float64(23),
		Fdy:   float32(0.010254488326609135),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+23),
		Fy:    float64(22.999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.9897454977035522)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+23),
		Fy:    float64(22.999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.9897454977035522)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+24),
		Fy:    float64(24),
		Fdy:   float32(0.002050897805020213),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+24),
		Fy:    float64(24),
		Fdy:   float32(0.002050897805020213),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+24),
		Fy:    float64(23.999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.9979491233825684)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+24),
		Fy:    float64(23.999999999999996),
		Fdy:   float32(-libc.Float64FromFloat64(0.9979491233825684)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+25),
		Fy:    float64(25),
		Fdy:   float32(-libc.Float64FromFloat64(0.011074847541749477)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1e+25),
		Fy:    float64(25.000000000000004),
		Fdy:   float32(0.9889251589775085),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1e+25),
		Fy:    float64(25),
		Fdy:   float32(-libc.Float64FromFloat64(0.011074847541749477)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1e+25),
		Fy:    float64(25),
		Fdy:   float32(-libc.Float64FromFloat64(0.011074847541749477)),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.605753595336707e+129),
		Fy:    float64(129.41593334566474),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.605753595336707e+129),
		Fy:    float64(129.41593334566474),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.605753595336707e+129),
		Fy:    float64(129.4159333456647),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.605753595336707e+129),
		Fy:    float64(129.4159333456647),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(28.89107606708965),
		Fy:    float64(1.4607637177284503),
		Fdy:   float32(-libc.Float64FromFloat64(4.314574075076592e-21)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(28.89107606708965),
		Fy:    float64(1.4607637177284505),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(28.89107606708965),
		Fy:    float64(1.4607637177284503),
		Fdy:   float32(-libc.Float64FromFloat64(4.314574075076592e-21)),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(28.89107606708965),
		Fy:    float64(1.4607637177284503),
		Fdy:   float32(-libc.Float64FromFloat64(4.314574075076592e-21)),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.691902599264667),
		Fy:    -libc.Float64FromFloat64(0.15995503788301924),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5831178468116087),
		Fy:    -libc.Float64FromFloat64(0.2342436664280542),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.8453327248693686),
		Fy:    -libc.Float64FromFloat64(0.07297231811166777),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5371455821827291),
		Fy:    float64(0.18671500123071508),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.359722003851219),
		Fy:    float64(0.6394587975157083),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.006884920524832),
		Fy:    float64(0.7786493118441783),
		Fdy:   float32(-libc.Float64FromFloat64(5.355238754073226e-17)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.295308954980527),
		Fy:    float64(0.9188325657605579),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.0123778241895036),
		Fy:    float64(0.47890944185468787),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(12.711023516648847),
		Fy:    float64(1.104180522209078),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(19.726781754842012),
		Fy:    float64(1.2950562398332561),
		Fdy:   float32(-libc.Float64FromFloat64(2.1549108454234518e-16)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(62.75068251591416),
		Fy:    float64(1.7976184538388658),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(129.03919495369695),
		Fy:    float64(2.110721644924475),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(242.85239235763245),
		Fy:    float64(2.3853423860985377),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4410.928517551039),
		Fy:    float64(3.6445300197632893),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7310.826593252415),
		Fy:    float64(3.863966482921384),
		Fdy:   float32(-libc.Float64FromFloat64(1.5519496478270901e-18)),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(44834.05376371908),
		Fy:    float64(4.651608008275136),
		Fdy:   float32(-libc.Float64FromFloat64(3.3196068235843285e-17)),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.675667871174408e+09),
		Fy:    float64(9.885116174396305),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.4252757787229706e+10),
		Fy:    float64(10.534695543523627),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.584673386971999e+11),
		Fy:    float64(11.412405671144613),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.4252757787229706e+11),
		Fy:    float64(11.534695543523627),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.4513899321627742e+12),
		Fy:    float64(12.16178410619122),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.8845761724226977e+18),
		Fy:    float64(18.27521369582876),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.0742163559775997e+27),
		Fy:    float64(27.316854054515197),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.8265639051332115e+34),
		Fy:    float64(34.68363806075155),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.610533074958623e+46),
		Fy:    float64(46.20696964820229),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.665648800401708e+47),
		Fy:    float64(47.425802930488686),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.410031201858835e+49),
		Fy:    float64(49.73319976986739),
		Fdy:   float32(-libc.Float64FromFloat64(9.697389045538866e-18)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0822935911379228e+68),
		Fy:    float64(68.03434508675538),
		Fdy:   float32(-libc.Float64FromFloat64(4.181351674341252e-18)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.971558130181311e+79),
		Fy:    float64(79.77608766430353),
		Fdy:   float32(-libc.Float64FromFloat64(4.924067220185265e-18)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.7804727985331277e+94),
		Fy:    float64(94.44411865071494),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.515203759795911e+95),
		Fy:    float64(95.87594076011325),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.4107144565965284e+163),
		Fy:    float64(163.14943911689105),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.2177936106109515e+173),
		Fy:    float64(173.5075581849938),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.83567360032565e+174),
		Fy:    float64(174.99280410811778),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.461828085706754e+241),
		Fy:    float64(241.73733802634126),
		Fdy:   float32(-libc.Float64FromFloat64(2.626066834284348e-18)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.620343395914844e+286),
		Fy:    float64(286.82088051679364),
		Fdy:   float32(-libc.Float64FromFloat64(4.147285067993868e-19)),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(308.25471555991675),
		Fdy:   float32(0.05135570839047432),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(39),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(308.2547155599167),
		Fdy:   float32(-libc.Float64FromFloat64(0.9486442804336548)),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(40),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(308.25471555991675),
		Fdy:   float32(0.05135570839047432),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(41),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(308.2547155599167),
		Fdy:   float32(-libc.Float64FromFloat64(0.9486442804336548)),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	153: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	154: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	155: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	156: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(100000),
		Fy:    float64(5),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	157: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+06),
		Fy:    float64(6),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	158: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+07),
		Fy:    float64(7),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	159: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+08),
		Fy:    float64(8),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	160: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+09),
		Fy:    float64(9),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	161: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+10),
		Fy:    float64(10),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	162: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+11),
		Fy:    float64(11),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	163: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+12),
		Fy:    float64(12),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	164: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+13),
		Fy:    float64(13),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	165: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+14),
		Fy:    float64(14),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	166: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+15),
		Fy:    float64(15),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	167: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+16),
		Fy:    float64(16),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	168: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+17),
		Fy:    float64(17),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	169: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+18),
		Fy:    float64(18),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	170: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+19),
		Fy:    float64(19),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	171: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+20),
		Fy:    float64(20),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	172: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+21),
		Fy:    float64(21),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	173: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1e+22),
		Fy:    float64(22),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	174: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.000000000000007),
		Fy:    float64(3.0858478929705076e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.4030088484287262)),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	176: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(3.0858478929705297e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.007297479547560215)),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(69),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.000000000000007),
		Fy:    float64(3.0858478929705076e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.4030088484287262)),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(70),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	179: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(3.0858478929705293e-15),
		Fdy:   float32(0.9927025437355042),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(72),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.000000000000007),
		Fy:    float64(3.085847892970508e-15),
		Fdy:   float32(0.5969911813735962),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(73),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	182: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(74),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(3.0858478929705293e-15),
		Fdy:   float32(0.9927025437355042),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(75),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.000000000000007),
		Fy:    float64(3.0858478929705076e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.4030088484287262)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(76),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	185: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(77),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999929),
		Fy:    -libc.Float64FromFloat64(3.0858478929705297e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.007297479547560215)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(79),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(307.6526555685888),
		Fdy:   float32(-libc.Float64FromFloat64(0.048455704003572464)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5e-324),
		Fy:    -libc.Float64FromFloat64(323.3062153431158),
		Fdy:   float32(-libc.Float64FromFloat64(0.1018017828464508)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(81),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(307.6526555685887),
		Fdy:   float32(0.9515442848205566),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(82),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5e-324),
		Fy:    -libc.Float64FromFloat64(323.30621534311575),
		Fdy:   float32(0.8981981873512268),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(83),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(307.6526555685887),
		Fdy:   float32(0.9515442848205566),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(84),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5e-324),
		Fy:    -libc.Float64FromFloat64(323.30621534311575),
		Fdy:   float32(0.8981981873512268),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(85),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.2250738585072014e-308),
		Fy:    -libc.Float64FromFloat64(307.6526555685888),
		Fdy:   float32(-libc.Float64FromFloat64(0.048455704003572464)),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(86),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5e-324),
		Fy:    -libc.Float64FromFloat64(323.3062153431158),
		Fdy:   float32(-libc.Float64FromFloat64(0.1018017828464508)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(88),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(36.56459649885151),
		Fy:    float64(1.5630607852185152),
		Fdy:   float32(0.4208488464355469),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(85.61132705430306),
		Fy:    float64(1.9325312290665935),
		Fdy:   float32(-libc.Float64FromFloat64(0.012649922631680965)),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(90),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(94.02246329044654),
		Fy:    float64(1.9732316250636186),
		Fdy:   float32(0.20653603971004486),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(16.60725376974054),
		Fy:    float64(1.2202978220206646),
		Fdy:   float32(0.4733176529407501),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(36.45486097807018),
		Fy:    float64(1.561755446376819),
		Fdy:   float32(0.3410508930683136),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(48.03540361357637),
		Fy:    float64(1.6815614441771134),
		Fdy:   float32(0.11853763461112976),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(78.91655385898032),
		Fy:    float64(1.8971681121515425),
		Fdy:   float32(0.16663385927677155),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(92.75441225031035),
		Fy:    float64(1.967334577806064),
		Fdy:   float32(0.2901192009449005),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(41.596529063707386),
		Fy:    float64(1.6190570933342316),
		Fdy:   float32(-libc.Float64FromFloat64(0.05394209176301956)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.098780043531456),
		Fy:    float64(0.3219669260116395),
		Fdy:   float32(0.43094536662101746),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	205: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(101),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	206: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	207: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	208: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	209: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	210: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	211: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	212: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	213: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	214: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(112),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	215: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	216: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	217: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	218: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	219: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	220: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	221: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	222: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	223: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(121),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	224: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	225: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(123),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	226: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	227: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	228: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(126),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	229: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(127),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	230: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	231: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	232: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(130),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	233: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(131),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	234: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	235: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(133),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	236: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(134),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	237: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	238: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(136),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	239: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(137),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	240: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	241: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(139),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	242: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(140),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	243: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(141),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	244: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(142),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	245: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(143),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	246: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(144),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	247: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(145),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	248: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(146),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	249: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(147),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	250: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(148),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	251: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(149),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	252: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(150),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	253: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(151),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	254: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(152),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	255: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(153),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	256: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(154),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	257: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	258: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	259: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	260: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	261: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	262: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	263: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	264: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	265: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	266: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	267: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	268: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(166),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	269: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	270: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	271: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	272: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	273: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	274: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(172),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	275: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	276: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(174),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	277: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(175),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	278: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	279: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	280: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	281: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	282: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	283: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	284: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	285: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(183),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	286: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	287: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	288: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	289: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	290: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(188),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	291: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(189),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	292: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(190),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	293: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(191),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	294: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(192),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	295: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(193),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	296: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(194),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	297: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(195),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	298: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(196),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	299: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(197),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	300: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(198),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	301: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(199),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	302: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(200),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	303: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(201),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	304: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(202),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	305: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(203),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	306: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(204),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	307: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(205),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	308: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(206),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	309: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(207),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	310: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(208),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	311: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(209),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	312: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(210),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	313: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(211),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	314: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(212),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	315: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(213),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	316: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(214),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	317: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(215),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	318: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(216),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	319: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(217),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	320: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(218),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	321: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(219),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	322: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(220),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	323: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(221),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	324: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(222),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	325: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(223),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	326: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(224),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	327: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(225),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	328: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(226),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	329: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(227),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	330: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(228),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	331: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(229),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	332: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(230),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	333: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(231),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	334: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(232),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	335: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	336: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	337: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(235),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	338: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(236),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	339: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(237),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	340: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	341: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	342: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(240),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5e-323),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	343: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(241),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	344: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(242),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	345: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	346: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(244),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	347: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	348: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	349: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	350: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	351: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	352: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.220446049250313e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	353: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.881784197001252e-16),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	354: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.25),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	355: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.5),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	356: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999993),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	357: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999996),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	358: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(256),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	359: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	360: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	361: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(259),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	362: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	363: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000009),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	364: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	365: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.000000000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	366: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(264),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.999999999999998),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	367: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(265),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	368: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	369: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155792e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	370: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(268),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.494232837155794e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	371: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(269),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	372: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(270),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.988465674311584e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	373: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	374: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(272),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	375: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(273),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	376: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(10),
		Fy:    float64(1),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	377: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(100),
		Fy:    float64(2),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	378: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1000),
		Fy:    float64(3),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	379: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(10000),
		Fy:    float64(4),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	380: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	381: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.345239849338305),
		Fy:    float64(0.6380137537120029),
		Fdy:   float32(-libc.Float64FromFloat64(0.2088824063539505)),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	383: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	384: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.267056966972586),
		Fy:    float64(0.9669418327487274),
		Fdy:   float32(-libc.Float64FromFloat64(0.06120431795716286)),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6619858980995045),
		Fy:    -libc.Float64FromFloat64(0.17915126198447093),
		Fdy:   float32(0.39090874791145325),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	387: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5617597462207241),
		Fy:    -libc.Float64FromFloat64(0.25044938407454437),
		Fdy:   float32(-libc.Float64FromFloat64(0.3046841621398926)),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7741522965913037),
		Fy:    -libc.Float64FromFloat64(0.11117359349943837),
		Fdy:   float32(-libc.Float64FromFloat64(0.31503361463546753)),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts + 46,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	390: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	391: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	392: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.888609052210118e-31),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	393: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	394: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	395: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	396: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	397: {
		Ffile: __ccgo_ts + 70,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+45)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:13:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:14:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:18:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:21:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(14328)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:22:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:24:3:
		if (*d_d)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:25:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:26:3:
		libc.Xfesetround(tls, (*d_d)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:27:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:28:3:
		y = libc.Xlog10(tls, (*d_d)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:29:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:31:3:
		if !(checkexcept(tls, e, (*d_d)(unsafe.Pointer(p)).Fe, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:32:4:
			libc.Xprintf(tls, __ccgo_ts+95, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, estr(tls, (*d_d)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:34:4:
			libc.Xprintf(tls, __ccgo_ts+145, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:35:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:37:3:
		d = ulperr(tls, y, (*d_d)(unsafe.Pointer(p)).Fy, (*d_d)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:38:3:
		if !(checkulp(tls, d, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+154, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*d_d)(unsafe.Pointer(p)).Fdy), float64((*d_d)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:41:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log10.c:44:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+212, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+222, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 266,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 274,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 282,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 292,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 302,
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
				v2 = __ccgo_ts + 311
			} else {
				v2 = __ccgo_ts + 45
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+313, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 311
		} else {
			v3 = __ccgo_ts + 45
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+318, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 45
	} else {
		v4 = __ccgo_ts + 323
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+325, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 328
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 331
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 334
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 337
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 340
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+343, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+351, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+356, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+399, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+448) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+456) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+468) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+480) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+492) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+501) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+45) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+501) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+507, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/log10.h\x00src/math/ucb/log10.h\x00\x00src/math/sanity/log10.h\x00src/math/special/log10.h\x00%s:%d: bad fp exception: %s log10(%a)=%a, want %s\x00 got %s\n\x00%s:%d: %s log10(%a) want %a got %a ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
