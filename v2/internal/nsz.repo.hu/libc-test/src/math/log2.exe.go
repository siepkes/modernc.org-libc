// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/log2.exe.go src/math/log2.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:5:19:
var t = [449]d_d{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7067244082187474),
		Fy:    float64(0.7712301192941611),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.293341100188992e-244),
		Fy:    -libc.Float64FromFloat64(808.1462765290187),
		Fdy:   float32(0.37797051668167114),
		Fe:    int32(FE_INEXACT),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.0257277093424366e-44),
		Fy:    -libc.Float64FromFloat64(144.5675540123593),
		Fdy:   float32(0.4138624966144562),
		Fe:    int32(FE_INEXACT),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5710275344015058),
		Fy:    -libc.Float64FromFloat64(0.808367782247107),
		Fdy:   float32(5.285554412293799e-16),
		Fe:    int32(FE_INEXACT),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5729535380737182),
		Fy:    -libc.Float64FromFloat64(0.8035099421327152),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.471765769146129),
		Fy:    float64(0.5575480854030254),
		Fdy:   float32(-libc.Float64FromFloat64(5.172307695432675e-16)),
		Fe:    int32(FE_INEXACT),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.9119621062425292),
		Fy:    float64(0.9350539303909772),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.471765769146129),
		Fy:    float64(0.5575480854030254),
		Fdy:   float32(-libc.Float64FromFloat64(5.172307695432675e-16)),
		Fe:    int32(FE_INEXACT),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.9119621062425292),
		Fy:    float64(0.9350539303909772),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.0094070305020213),
		Fy:    float64(1.5894792483041358),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.076943476364032),
		Fy:    float64(1.6214979415205624),
		Fdy:   float32(-libc.Float64FromFloat64(1.7176054695083897e-16)),
		Fe:    int32(FE_INEXACT),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(27),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.7892221739993412),
		Fy:    float64(1.9219017320688592),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.8922310293086957),
		Fy:    float64(1.960597346024947),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(29),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.431861864920373),
		Fy:    float64(2.1479129152544147),
		Fdy:   float32(-libc.Float64FromFloat64(3.240793054820276e-17)),
		Fe:    int32(FE_INEXACT),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(30),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.1950938638155035),
		Fy:    float64(2.8470135078103054),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(31),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(28.32197270313262),
		Fy:    float64(4.823849851477344),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(33),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.64380155103743e-265),
		Fy:    -libc.Float64FromFloat64(878.9083312557786),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(34),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.887684129119793e+110),
		Fy:    float64(368.56389798144636),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.887684129119793e+110),
		Fy:    float64(368.5638979814464),
		Fdy:   float32(5.6959166272463335e-08),
		Fe:    int32(FE_INEXACT),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.887684129119793e+110),
		Fy:    float64(368.56389798144636),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.218173622828685e+84),
		Fy:    float64(281.1185784491785),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.218173622828685e+84),
		Fy:    float64(281.11857844917853),
		Fdy:   float32(1.8403159174340544e-07),
		Fe:    int32(FE_INEXACT),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.218173622828685e+84),
		Fy:    float64(281.1185784491785),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.286266690475151e-264),
		Fy:    -libc.Float64FromFloat64(875.796023348118),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.286266690475151e-264),
		Fy:    -libc.Float64FromFloat64(875.7960233481178),
		Fdy:   float32(1.243382996563014e-07),
		Fe:    int32(FE_INEXACT),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.286266690475151e-264),
		Fy:    -libc.Float64FromFloat64(875.7960233481178),
		Fdy:   float32(1.243382996563014e-07),
		Fe:    int32(FE_INEXACT),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1114408152853843e-203),
		Fy:    -libc.Float64FromFloat64(674.1989721358727),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5670140455468004e-39),
		Fy:    -libc.Float64FromFloat64(128.90717758949486),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.7014515127389058e+20),
		Fy:    float64(67.87229668559492),
		Fdy:   float32(-libc.Float64FromFloat64(1.7564364895861218e-07)),
		Fe:    int32(FE_INEXACT),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.7014515127389058e+20),
		Fy:    float64(67.87229668559493),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.7014515127389058e+20),
		Fy:    float64(67.87229668559492),
		Fdy:   float32(-libc.Float64FromFloat64(1.7564364895861218e-07)),
		Fe:    int32(FE_INEXACT),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.4260691763131828e+132),
		Fy:    float64(439.00655249141806),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5319969820851656e+254),
		Fy:    float64(844.3851495566588),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.429077056757102e-296),
		Fy:    -libc.Float64FromFloat64(980.0535995236916),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.429077056757102e-296),
		Fy:    -libc.Float64FromFloat64(980.0535995236914),
		Fdy:   float32(1.80357773160722e-07),
		Fe:    int32(FE_INEXACT),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.429077056757102e-296),
		Fy:    -libc.Float64FromFloat64(980.0535995236914),
		Fdy:   float32(1.80357773160722e-07),
		Fe:    int32(FE_INEXACT),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.570657878503617e-167),
		Fy:    -libc.Float64FromFloat64(552.0459540204611),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.8846766620733993e+81),
		Fy:    float64(269.99049271969614),
		Fdy:   float32(-libc.Float64FromFloat64(2.783874997192015e-08)),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.8846766620733993e+81),
		Fy:    float64(269.9904927196962),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.8846766620733993e+81),
		Fy:    float64(269.99049271969614),
		Fdy:   float32(-libc.Float64FromFloat64(2.783874997192015e-08)),
		Fe:    int32(FE_INEXACT),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.635263288664831e-97),
		Fy:    -libc.Float64FromFloat64(321.51750226580003),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.635263288664831e-97),
		Fy:    -libc.Float64FromFloat64(321.5175022658),
		Fdy:   float32(1.158567641823538e-07),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.635263288664831e-97),
		Fy:    -libc.Float64FromFloat64(321.5175022658),
		Fdy:   float32(1.158567641823538e-07),
		Fe:    int32(FE_INEXACT),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.8871859615316325e+260),
		Fy:    float64(865.6600307980367),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.8871859615316325e+260),
		Fy:    float64(865.6600307980368),
		Fdy:   float32(1.6223033583173674e-07),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.8871859615316325e+260),
		Fy:    float64(865.6600307980367),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.5376744471831127e-71),
		Fy:    -libc.Float64FromFloat64(234.51338773620822),
		Fdy:   float32(-libc.Float64FromFloat64(1.9425760910962708e-07)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.5376744471831127e-71),
		Fy:    -libc.Float64FromFloat64(234.5133877362082),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.5376744471831127e-71),
		Fy:    -libc.Float64FromFloat64(234.5133877362082),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.324838749857688e+130),
		Fy:    float64(433.96329867871515),
		Fdy:   float32(-libc.Float64FromFloat64(1.286587689719454e-07)),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.324838749857688e+130),
		Fy:    float64(433.9632986787152),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.324838749857688e+130),
		Fy:    float64(433.96329867871515),
		Fdy:   float32(-libc.Float64FromFloat64(1.286587689719454e-07)),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.6825235505336408e-229),
		Fy:    -libc.Float64FromFloat64(759.2979428923943),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.6825235505336408e-229),
		Fy:    -libc.Float64FromFloat64(759.2979428923942),
		Fdy:   float32(1.1306534020150139e-07),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.6825235505336408e-229),
		Fy:    -libc.Float64FromFloat64(759.2979428923942),
		Fdy:   float32(1.1306534020150139e-07),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.431983280192438e-128),
		Fy:    -libc.Float64FromFloat64(423.42775362119545),
		Fdy:   float32(-libc.Float64FromFloat64(2.2590610626593843e-07)),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.431983280192438e-128),
		Fy:    -libc.Float64FromFloat64(423.4277536211954),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.431983280192438e-128),
		Fy:    -libc.Float64FromFloat64(423.4277536211954),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5244387760920075e+94),
		Fy:    float64(312.8695191299748),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.720328623925579e-234),
		Fy:    -libc.Float64FromFloat64(774.8150761736865),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.762569887530403e+197),
		Fy:    float64(656.9465470352343),
		Fdy:   float32(-libc.Float64FromFloat64(1.2235837232310587e-07)),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.762569887530403e+197),
		Fy:    float64(656.9465470352344),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.762569887530403e+197),
		Fy:    float64(656.9465470352343),
		Fdy:   float32(-libc.Float64FromFloat64(1.2235837232310587e-07)),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.323777238909358e-231),
		Fy:    -libc.Float64FromFloat64(766.1489181427186),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0698811339103488e-246),
		Fy:    -libc.Float64FromFloat64(817.096860823276),
		Fdy:   float32(-libc.Float64FromFloat64(1.4652773927537055e-07)),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0698811339103488e-246),
		Fy:    -libc.Float64FromFloat64(817.0968608232758),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0698811339103488e-246),
		Fy:    -libc.Float64FromFloat64(817.0968608232758),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7918130062514963e-298),
		Fy:    -libc.Float64FromFloat64(989.0931521909527),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7918130062514963e-298),
		Fy:    -libc.Float64FromFloat64(989.0931521909525),
		Fdy:   float32(1.499738431220976e-07),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7918130062514963e-298),
		Fy:    -libc.Float64FromFloat64(989.0931521909525),
		Fdy:   float32(1.499738431220976e-07),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.523585676661139e+215),
		Fy:    float64(715.5500154682775),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.18131110803778e-291),
		Fy:    -libc.Float64FromFloat64(965.5558800641538),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.18131110803778e-291),
		Fy:    -libc.Float64FromFloat64(965.5558800641537),
		Fdy:   float32(2.058094423773582e-07),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.18131110803778e-291),
		Fy:    -libc.Float64FromFloat64(965.5558800641537),
		Fdy:   float32(2.058094423773582e-07),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.869256282502346e+120),
		Fy:    float64(401.78018452116424),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.4353631263944388e+38),
		Fy:    float64(127.5174045080437),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.4353631263944388e+38),
		Fy:    float64(127.51740450804371),
		Fdy:   float32(1.9556499353257095e-07),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.4353631263944388e+38),
		Fy:    float64(127.5174045080437),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0580397562205371e+72),
		Fy:    float64(239.26021667014032),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.8897005521272254e-159),
		Fy:    -libc.Float64FromFloat64(526.6556470872174),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.8897005521272254e-159),
		Fy:    -libc.Float64FromFloat64(526.6556470872173),
		Fdy:   float32(4.22074535322281e-08),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.8897005521272254e-159),
		Fy:    -libc.Float64FromFloat64(526.6556470872173),
		Fdy:   float32(4.22074535322281e-08),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.1932460679690458e+143),
		Fy:    float64(476.71074129855003),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.5227350302983266e+212),
		Fy:    float64(705.5837448002197),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.413511436360485e+232),
		Fy:    float64(771.1866015695618),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.413511436360485e+232),
		Fy:    float64(771.186601569562),
		Fdy:   float32(2.061880195469712e-07),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.413511436360485e+232),
		Fy:    float64(771.1866015695618),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.694129841877631e+298),
		Fy:    float64(990.695116726697),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1296389003573722e+99),
		Fy:    float64(329.0467430692158),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1605609653189456e+19),
		Fy:    float64(63.33145611361045),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.2197577663346064e-243),
		Fy:    -libc.Float64FromFloat64(806.9419323885347),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.2197577663346064e-243),
		Fy:    -libc.Float64FromFloat64(806.9419323885346),
		Fdy:   float32(1.4304511353202543e-07),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.2197577663346064e-243),
		Fy:    -libc.Float64FromFloat64(806.9419323885346),
		Fdy:   float32(1.4304511353202543e-07),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0053173097190037e+53),
		Fy:    float64(176.06983996218102),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0053173097190037e+53),
		Fy:    float64(176.06983996218105),
		Fdy:   float32(1.7596211421277985e-07),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0053173097190037e+53),
		Fy:    float64(176.06983996218102),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.2452978506569114e+17),
		Fy:    float64(56.789268460818484),
		Fdy:   float32(-libc.Float64FromFloat64(1.1002602917642434e-07)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.2452978506569114e+17),
		Fy:    float64(56.78926846081849),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.2452978506569114e+17),
		Fy:    float64(56.789268460818484),
		Fdy:   float32(-libc.Float64FromFloat64(1.1002602917642434e-07)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(7.38700350823881e+241),
		Fy:    float64(803.469660130671),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(7.38700350823881e+241),
		Fy:    float64(803.4696601306711),
		Fdy:   float32(1.7090644632844487e-07),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7.38700350823881e+241),
		Fy:    float64(803.469660130671),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.3809448904637181e-211),
		Fy:    -libc.Float64FromFloat64(700.4611722742959),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.3809448904637181e-211),
		Fy:    -libc.Float64FromFloat64(700.4611722742958),
		Fdy:   float32(9.484347884836097e-08),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.3809448904637181e-211),
		Fy:    -libc.Float64FromFloat64(700.4611722742958),
		Fdy:   float32(9.484347884836097e-08),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.762943267812643e+34),
		Fy:    float64(114.41176116541604),
		Fdy:   float32(-libc.Float64FromFloat64(1.2273949323571287e-07)),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.762943267812643e+34),
		Fy:    float64(114.41176116541605),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.762943267812643e+34),
		Fy:    float64(114.41176116541604),
		Fdy:   float32(-libc.Float64FromFloat64(1.2273949323571287e-07)),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.6329902802770947e-259),
		Fy:    -libc.Float64FromFloat64(858.5182190698629),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1187149122437587e-101),
		Fy:    -libc.Float64FromFloat64(335.3528951497167),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.567913549931045e-150),
		Fy:    -libc.Float64FromFloat64(495.0310098788578),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.747146667134389e+180),
		Fy:    float64(600.1941177034716),
		Fdy:   float32(-libc.Float64FromFloat64(1.211654989674571e-07)),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.747146667134389e+180),
		Fy:    float64(600.1941177034718),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.747146667134389e+180),
		Fy:    float64(600.1941177034716),
		Fdy:   float32(-libc.Float64FromFloat64(1.211654989674571e-07)),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1662759151658542e-191),
		Fy:    -libc.Float64FromFloat64(634.2663569847322),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.542157602849816e+110),
		Fy:    float64(367.23671884162275),
		Fdy:   float32(-libc.Float64FromFloat64(8.08097553317566e-08)),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.542157602849816e+110),
		Fy:    float64(367.2367188416228),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.542157602849816e+110),
		Fy:    float64(367.23671884162275),
		Fdy:   float32(-libc.Float64FromFloat64(8.08097553317566e-08)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.392156372850442e+55),
		Fy:    float64(185.38234792049641),
		Fdy:   float32(-libc.Float64FromFloat64(1.967923424217588e-07)),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.392156372850442e+55),
		Fy:    float64(185.38234792049644),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.392156372850442e+55),
		Fy:    float64(185.38234792049641),
		Fdy:   float32(-libc.Float64FromFloat64(1.967923424217588e-07)),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.181523575370562e+78),
		Fy:    float64(262.3091249744883),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.181523575370562e+78),
		Fy:    float64(262.3091249744884),
		Fdy:   float32(1.2715779007521633e-07),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.181523575370562e+78),
		Fy:    float64(262.3091249744883),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.288077595425543e+18),
		Fy:    float64(61.89503672046661),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.3393392361720464e-08),
		Fy:    -libc.Float64FromFloat64(26.15390333737849),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.3393392361720464e-08),
		Fy:    -libc.Float64FromFloat64(26.153903337378487),
		Fdy:   float32(1.6499511445999815e-07),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.3393392361720464e-08),
		Fy:    -libc.Float64FromFloat64(26.153903337378487),
		Fdy:   float32(1.6499511445999815e-07),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.457820423589657e+38),
		Fy:    float64(128.3896061067466),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.112952917855509e+264),
		Fy:    float64(877.143409612743),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.112952917855509e+264),
		Fy:    float64(877.1434096127431),
		Fdy:   float32(1.149399579958299e-07),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.112952917855509e+264),
		Fy:    float64(877.143409612743),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.43370106579888e+207),
		Fy:    float64(690.0810508397884),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.43370106579888e+207),
		Fy:    float64(690.0810508397885),
		Fdy:   float32(1.380311971388437e-07),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.43370106579888e+207),
		Fy:    float64(690.0810508397884),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.82720299806966e-194),
		Fy:    -libc.Float64FromFloat64(642.1828629105246),
		Fdy:   float32(-libc.Float64FromFloat64(1.3830597822561685e-07)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.82720299806966e-194),
		Fy:    -libc.Float64FromFloat64(642.1828629105245),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.82720299806966e-194),
		Fy:    -libc.Float64FromFloat64(642.1828629105245),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.343872148547027e-07),
		Fy:    -libc.Float64FromFloat64(20.835611172279744),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.6726849446749785e+294),
		Fy:    float64(978.0651496800158),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.6726849446749785e+294),
		Fy:    float64(978.0651496800159),
		Fdy:   float32(1.1872887029085177e-07),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.6726849446749785e+294),
		Fy:    float64(978.0651496800158),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.699051487358758e+202),
		Fy:    float64(674.307318833991),
		Fdy:   float32(-libc.Float64FromFloat64(5.207492037584416e-08)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.699051487358758e+202),
		Fy:    float64(674.3073188339911),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.699051487358758e+202),
		Fy:    float64(674.307318833991),
		Fdy:   float32(-libc.Float64FromFloat64(5.207492037584416e-08)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.007720190187349e+42),
		Fy:    float64(141.52376177166641),
		Fdy:   float32(-libc.Float64FromFloat64(4.703699119090743e-08)),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.007720190187349e+42),
		Fy:    float64(141.52376177166644),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.007720190187349e+42),
		Fy:    float64(141.52376177166641),
		Fdy:   float32(-libc.Float64FromFloat64(4.703699119090743e-08)),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.3685832417863484e+13),
		Fy:    float64(45.60960664768264),
		Fdy:   float32(-libc.Float64FromFloat64(2.1661051619048521e-07)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.3685832417863484e+13),
		Fy:    float64(45.60960664768265),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.3685832417863484e+13),
		Fy:    float64(45.60960664768264),
		Fdy:   float32(-libc.Float64FromFloat64(2.1661051619048521e-07)),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.15920990893684e+10),
		Fy:    float64(36.05908132919082),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.1678071194026492e+11),
		Fy:    float64(37.6574454423991),
		Fdy:   float32(-libc.Float64FromFloat64(1.4884659549352364e-07)),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.1678071194026492e+11),
		Fy:    float64(37.65744544239911),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.1678071194026492e+11),
		Fy:    float64(37.6574454423991),
		Fdy:   float32(-libc.Float64FromFloat64(1.4884659549352364e-07)),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.60336490091812e+58),
		Fy:    float64(195.93536950047633),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.60336490091812e+58),
		Fy:    float64(195.93536950047636),
		Fdy:   float32(1.005295757749991e-08),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.60336490091812e+58),
		Fy:    float64(195.93536950047633),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.601285918402335e+98),
		Fy:    float64(328.81218094061427),
		Fdy:   float32(-libc.Float64FromFloat64(1.3238378926416772e-07)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.601285918402335e+98),
		Fy:    float64(328.8121809406143),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.601285918402335e+98),
		Fy:    float64(328.81218094061427),
		Fdy:   float32(-libc.Float64FromFloat64(1.3238378926416772e-07)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.3411321257967105e-202),
		Fy:    -libc.Float64FromFloat64(669.289138131757),
		Fdy:   float32(-libc.Float64FromFloat64(4.416074972368733e-08)),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.3411321257967105e-202),
		Fy:    -libc.Float64FromFloat64(669.2891381317569),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.3411321257967105e-202),
		Fy:    -libc.Float64FromFloat64(669.2891381317569),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.809233358023687e+171),
		Fy:    float64(569.9792048975075),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.809233358023687e+171),
		Fy:    float64(569.9792048975077),
		Fdy:   float32(2.09860573363585e-07),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.809233358023687e+171),
		Fy:    float64(569.9792048975075),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.1798058579385668e-198),
		Fy:    -libc.Float64FromFloat64(656.072824102798),
		Fdy:   float32(-libc.Float64FromFloat64(1.571039547343389e-07)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.1798058579385668e-198),
		Fy:    -libc.Float64FromFloat64(656.0728241027979),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.1798058579385668e-198),
		Fy:    -libc.Float64FromFloat64(656.0728241027979),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.519455532745495e+209),
		Fy:    float64(696.7474977906815),
		Fdy:   float32(-libc.Float64FromFloat64(1.6250162104824994e-07)),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.519455532745495e+209),
		Fy:    float64(696.7474977906816),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.519455532745495e+209),
		Fy:    float64(696.7474977906815),
		Fdy:   float32(-libc.Float64FromFloat64(1.6250162104824994e-07)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.8367833097292497e-147),
		Fy:    -libc.Float64FromFloat64(486.38353266038627),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0279716568765754e-97),
		Fy:    -libc.Float64FromFloat64(322.1872247168257),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.4425031543568829e+186),
		Fy:    float64(618.4072001227001),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(7.265441319221895e+80),
		Fy:    float64(268.6152980240454),
		Fdy:   float32(-libc.Float64FromFloat64(4.285656274305438e-08)),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(7.265441319221895e+80),
		Fy:    float64(268.61529802404544),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7.265441319221895e+80),
		Fy:    float64(268.6152980240454),
		Fdy:   float32(-libc.Float64FromFloat64(4.285656274305438e-08)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.994171975950923e+42),
		Fy:    float64(142.3271332553971),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.620262975806255e-132),
		Fy:    -libc.Float64FromFloat64(435.3867766432764),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.620262975806255e-132),
		Fy:    -libc.Float64FromFloat64(435.38677664327633),
		Fdy:   float32(1.4174828777413495e-07),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.620262975806255e-132),
		Fy:    -libc.Float64FromFloat64(435.38677664327633),
		Fdy:   float32(1.4174828777413495e-07),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.7087848311413374e-16),
		Fy:    -libc.Float64FromFloat64(51.259902947323695),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.127141657900579e+130),
		Fy:    float64(432.02332117855815),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.11490847662199e+63),
		Fy:    float64(212.30204480144255),
		Fdy:   float32(-libc.Float64FromFloat64(8.016701968927009e-08)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.11490847662199e+63),
		Fy:    float64(212.30204480144258),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.11490847662199e+63),
		Fy:    float64(212.30204480144255),
		Fdy:   float32(-libc.Float64FromFloat64(8.016701968927009e-08)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.7120365812937125e+41),
		Fy:    float64(138.09126281835566),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.7120365812937125e+41),
		Fy:    float64(138.0912628183557),
		Fdy:   float32(2.3507051594151562e-07),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.7120365812937125e+41),
		Fy:    float64(138.09126281835566),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.158695266422826e+118),
		Fy:    float64(393.0976747954884),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.142389163740648e+265),
		Fy:    float64(882.6733839391719),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.142389163740648e+265),
		Fy:    float64(882.673383939172),
		Fdy:   float32(5.499062183389469e-08),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.142389163740648e+265),
		Fy:    float64(882.6733839391719),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.5461877435720934e-216),
		Fy:    -libc.Float64FromFloat64(715.7101995777217),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.5069800346385503e-151),
		Fy:    -libc.Float64FromFloat64(501.0194820246052),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.569601076884624e-144),
		Fy:    -libc.Float64FromFloat64(475.6418398947354),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.8701590396760317e+26),
		Fy:    float64(87.89126114801239),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.186248832577276e-192),
		Fy:    -libc.Float64FromFloat64(635.7445361497672),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.186248832577276e-192),
		Fy:    -libc.Float64FromFloat64(635.7445361497671),
		Fdy:   float32(3.548946381215501e-08),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.186248832577276e-192),
		Fy:    -libc.Float64FromFloat64(635.7445361497671),
		Fdy:   float32(3.548946381215501e-08),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.751224505659147e+176),
		Fy:    float64(587.9449280961987),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.751224505659147e+176),
		Fy:    float64(587.9449280961989),
		Fdy:   float32(7.83603866238991e-08),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.751224505659147e+176),
		Fy:    float64(587.9449280961987),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.7156373981389626e+26),
		Fy:    float64(88.8850448613743),
		Fdy:   float32(-libc.Float64FromFloat64(4.7134509628676824e-08)),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.7156373981389626e+26),
		Fy:    float64(88.88504486137431),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.7156373981389626e+26),
		Fy:    float64(88.8850448613743),
		Fdy:   float32(-libc.Float64FromFloat64(4.7134509628676824e-08)),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.6907957300593165e+184),
		Fy:    float64(611.9924718333741),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.0183614312708062e+165),
		Fy:    float64(549.7119012269849),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1411143477659629e-76),
		Fy:    -libc.Float64FromFloat64(252.27609184427473),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7183195160460042e-199),
		Fy:    -libc.Float64FromFloat64(660.2826925566122),
		Fdy:   float32(-libc.Float64FromFloat64(2.082099115341407e-07)),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7183195160460042e-199),
		Fy:    -libc.Float64FromFloat64(660.2826925566121),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7183195160460042e-199),
		Fy:    -libc.Float64FromFloat64(660.2826925566121),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.941334646044936e+46),
		Fy:    float64(155.60390544867508),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.229444497453363e-69),
		Fy:    -libc.Float64FromFloat64(227.5217525215324),
		Fdy:   float32(-libc.Float64FromFloat64(4.426535582524593e-09)),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.229444497453363e-69),
		Fy:    -libc.Float64FromFloat64(227.52175252153236),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.229444497453363e-69),
		Fy:    -libc.Float64FromFloat64(227.52175252153236),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.84839785413112e+244),
		Fy:    float64(813.6958714080346),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.278238436914253e-220),
		Fy:    -libc.Float64FromFloat64(729.1112600850767),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.278238436914253e-220),
		Fy:    -libc.Float64FromFloat64(729.1112600850765),
		Fdy:   float32(3.328691633441849e-08),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.278238436914253e-220),
		Fy:    -libc.Float64FromFloat64(729.1112600850765),
		Fdy:   float32(3.328691633441849e-08),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.697440464250781e-138),
		Fy:    -libc.Float64FromFloat64(455.30548619611807),
		Fdy:   float32(-libc.Float64FromFloat64(1.1783620124106164e-07)),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.697440464250781e-138),
		Fy:    -libc.Float64FromFloat64(455.305486196118),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.697440464250781e-138),
		Fy:    -libc.Float64FromFloat64(455.305486196118),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.025218582222422e-200),
		Fy:    -libc.Float64FromFloat64(664.349687444455),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.000764052102743e+58),
		Fy:    float64(194.25715938879983),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.000764052102743e+58),
		Fy:    float64(194.25715938879986),
		Fdy:   float32(1.8474729301942716e-07),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.000764052102743e+58),
		Fy:    float64(194.25715938879983),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.3169606002030296e-197),
		Fy:    -libc.Float64FromFloat64(653.2076011813547),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.5005808316840458e+291),
		Fy:    float64(967.2665966468059),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.5005808316840458e+291),
		Fy:    float64(967.266596646806),
		Fdy:   float32(3.907239332079371e-08),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.5005808316840458e+291),
		Fy:    float64(967.2665966468059),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.3849224714228564e+16),
		Fy:    float64(53.620654734021414),
		Fdy:   float32(-libc.Float64FromFloat64(2.302469113146799e-07)),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.3849224714228564e+16),
		Fy:    float64(53.62065473402142),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.3849224714228564e+16),
		Fy:    float64(53.620654734021414),
		Fdy:   float32(-libc.Float64FromFloat64(2.302469113146799e-07)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.468946395287264e+38),
		Fy:    float64(129.31545010363118),
		Fdy:   float32(-libc.Float64FromFloat64(1.4141831172764796e-07)),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.468946395287264e+38),
		Fy:    float64(129.3154501036312),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.468946395287264e+38),
		Fy:    float64(129.31545010363118),
		Fdy:   float32(-libc.Float64FromFloat64(1.4141831172764796e-07)),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.010943448042358e-231),
		Fy:    -libc.Float64FromFloat64(766.3575174083846),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7450676962935243e+74),
		Fy:    float64(246.62596202555682),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.265079685323101e+209),
		Fy:    float64(694.6222000922675),
		Fdy:   float32(-libc.Float64FromFloat64(5.42662021985052e-08)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.265079685323101e+209),
		Fy:    float64(694.6222000922676),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.265079685323101e+209),
		Fy:    float64(694.6222000922675),
		Fdy:   float32(-libc.Float64FromFloat64(5.42662021985052e-08)),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.5431888023440114e+224),
		Fy:    float64(744.7378078348515),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.5431888023440114e+224),
		Fy:    float64(744.7378078348517),
		Fdy:   float32(1.0962097718447694e-07),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.5431888023440114e+224),
		Fy:    float64(744.7378078348515),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0661033432027428e+57),
		Fy:    float64(189.4422487017421),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.698750655674708e+64),
		Fy:    float64(214.49043611995586),
		Fdy:   float32(-libc.Float64FromFloat64(4.079477378127194e-08)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.698750655674708e+64),
		Fy:    float64(214.4904361199559),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.698750655674708e+64),
		Fy:    float64(214.49043611995586),
		Fdy:   float32(-libc.Float64FromFloat64(4.079477378127194e-08)),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0334563366568723e-129),
		Fy:    -libc.Float64FromFloat64(428.4812468040231),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.9083937123260755e-174),
		Fy:    -libc.Float64FromFloat64(577.0831296722732),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.9083937123260755e-174),
		Fy:    -libc.Float64FromFloat64(577.0831296722731),
		Fdy:   float32(9.392132405139364e-09),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.9083937123260755e-174),
		Fy:    -libc.Float64FromFloat64(577.0831296722731),
		Fdy:   float32(9.392132405139364e-09),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.7803767264302897e-194),
		Fy:    -libc.Float64FromFloat64(642.9787700343613),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.788096879913045e-11),
		Fy:    -libc.Float64FromFloat64(33.579938213208074),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.3125463616098785e-100),
		Fy:    -libc.Float64FromFloat64(330.9833271998666),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.3125463616098785e-100),
		Fy:    -libc.Float64FromFloat64(330.98332719986655),
		Fdy:   float32(3.475582133205535e-08),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.3125463616098785e-100),
		Fy:    -libc.Float64FromFloat64(330.98332719986655),
		Fdy:   float32(3.475582133205535e-08),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.566661747317308e-181),
		Fy:    -libc.Float64FromFloat64(599.0778452413755),
		Fdy:   float32(-libc.Float64FromFloat64(2.2449472680818872e-07)),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.566661747317308e-181),
		Fy:    -libc.Float64FromFloat64(599.0778452413754),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.566661747317308e-181),
		Fy:    -libc.Float64FromFloat64(599.0778452413754),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.407507073192908e+305),
		Fy:    float64(1016.2597470281298),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.669607707142408),
		Fy:    float64(3.273457361218098),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.866622804657321e+44),
		Fy:    float64(148.71736641175343),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.0162982978238514e-205),
		Fy:    -libc.Float64FromFloat64(679.9835503602164),
		Fdy:   float32(-libc.Float64FromFloat64(9.800487532629631e-08)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.0162982978238514e-205),
		Fy:    -libc.Float64FromFloat64(679.9835503602163),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.0162982978238514e-205),
		Fy:    -libc.Float64FromFloat64(679.9835503602163),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.912034711098946e-282),
		Fy:    -libc.Float64FromFloat64(935.2416952058479),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(7.092063453310183e+166),
		Fy:    float64(554.2662691955148),
		Fdy:   float32(-libc.Float64FromFloat64(1.2211324929012335e-07)),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(7.092063453310183e+166),
		Fy:    float64(554.266269195515),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7.092063453310183e+166),
		Fy:    float64(554.2662691955148),
		Fdy:   float32(-libc.Float64FromFloat64(1.2211324929012335e-07)),
		Fe:    int32(FE_INEXACT),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.410529030874095e+249),
		Fy:    float64(828.930091170021),
		Fdy:   float32(-libc.Float64FromFloat64(8.608466117721036e-08)),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.410529030874095e+249),
		Fy:    float64(828.9300911700211),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.410529030874095e+249),
		Fy:    float64(828.930091170021),
		Fdy:   float32(-libc.Float64FromFloat64(8.608466117721036e-08)),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.0600626124445356e-278),
		Fy:    -libc.Float64FromFloat64(921.156855141968),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.0600626124445356e-278),
		Fy:    -libc.Float64FromFloat64(921.1568551419679),
		Fdy:   float32(5.944794523315977e-08),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.0600626124445356e-278),
		Fy:    -libc.Float64FromFloat64(921.1568551419679),
		Fdy:   float32(5.944794523315977e-08),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.400477624232371e-214),
		Fy:    -libc.Float64FromFloat64(707.8221409486564),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.400477624232371e-214),
		Fy:    -libc.Float64FromFloat64(707.8221409486563),
		Fdy:   float32(2.2100377350398048e-07),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.400477624232371e-214),
		Fy:    -libc.Float64FromFloat64(707.8221409486563),
		Fdy:   float32(2.2100377350398048e-07),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.684561295137765e+255),
		Fy:    float64(849.8324970762524),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.0769136291102504e+287),
		Fy:    float64(955.0148471795609),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.4425459992716953e-103),
		Fy:    -libc.Float64FromFloat64(340.007207059029),
		Fdy:   float32(-libc.Float64FromFloat64(1.3908113771776698e-07)),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.4425459992716953e-103),
		Fy:    -libc.Float64FromFloat64(340.00720705902893),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.4425459992716953e-103),
		Fy:    -libc.Float64FromFloat64(340.00720705902893),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.257715657640541e+114),
		Fy:    float64(380.40366350413694),
		Fdy:   float32(-libc.Float64FromFloat64(1.6324347029694763e-07)),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.257715657640541e+114),
		Fy:    float64(380.403663504137),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.257715657640541e+114),
		Fy:    float64(380.40366350413694),
		Fdy:   float32(-libc.Float64FromFloat64(1.6324347029694763e-07)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.5750931375923394e-105),
		Fy:    -libc.Float64FromFloat64(346.9644691359975),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.977966622418147e+262),
		Fy:    float64(871.3291789417536),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.977966622418147e+262),
		Fy:    float64(871.3291789417538),
		Fdy:   float32(1.2155385320511414e-07),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.977966622418147e+262),
		Fy:    float64(871.3291789417536),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.3963385000147476e-42),
		Fy:    -libc.Float64FromFloat64(139.03933126356787),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.3963385000147476e-42),
		Fy:    -libc.Float64FromFloat64(139.03933126356785),
		Fdy:   float32(1.900298940427092e-07),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.3963385000147476e-42),
		Fy:    -libc.Float64FromFloat64(139.03933126356785),
		Fdy:   float32(1.900298940427092e-07),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.6556869780851722e+38),
		Fy:    float64(126.9606975503864),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.6556869780851722e+38),
		Fy:    float64(126.96069755038641),
		Fdy:   float32(1.919219094759228e-08),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.6556869780851722e+38),
		Fy:    float64(126.9606975503864),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.652024254307643e-66),
		Fy:    -libc.Float64FromFloat64(216.51346083364223),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.092736517981007e-21),
		Fy:    -libc.Float64FromFloat64(67.41204891392954),
		Fdy:   float32(-libc.Float64FromFloat64(1.8662053946627566e-07)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.092736517981007e-21),
		Fy:    -libc.Float64FromFloat64(67.41204891392952),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.092736517981007e-21),
		Fy:    -libc.Float64FromFloat64(67.41204891392952),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0296896350282387e-266),
		Fy:    -libc.Float64FromFloat64(883.590663688533),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0971833019151535e+127),
		Fy:    float64(422.0186726217165),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.5743251638375265e+92),
		Fy:    float64(307.81094365235725),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.114930344205035e+261),
		Fy:    float64(870.0438114768256),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.114930344205035e+261),
		Fy:    float64(870.0438114768257),
		Fdy:   float32(1.9889445468379563e-07),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.114930344205035e+261),
		Fy:    float64(870.0438114768256),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.868744617365426e+278),
		Fy:    float64(925.0164299202825),
		Fdy:   float32(-libc.Float64FromFloat64(1.477753812650917e-07)),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.868744617365426e+278),
		Fy:    float64(925.0164299202826),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.868744617365426e+278),
		Fy:    float64(925.0164299202825),
		Fdy:   float32(-libc.Float64FromFloat64(1.477753812650917e-07)),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.847812982784006e-83),
		Fy:    -libc.Float64FromFloat64(272.9443885741011),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.847812982784006e-83),
		Fy:    -libc.Float64FromFloat64(272.944388574101),
		Fdy:   float32(1.4439259565790508e-08),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.847812982784006e-83),
		Fy:    -libc.Float64FromFloat64(272.944388574101),
		Fdy:   float32(1.4439259565790508e-08),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.26449161340012e-77),
		Fy:    -libc.Float64FromFloat64(252.57675149542246),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.082589565643832e-232),
		Fy:    -libc.Float64FromFloat64(767.6725004244837),
		Fdy:   float32(-libc.Float64FromFloat64(2.0252443277968268e-07)),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.082589565643832e-232),
		Fy:    -libc.Float64FromFloat64(767.6725004244836),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.082589565643832e-232),
		Fy:    -libc.Float64FromFloat64(767.6725004244836),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.7045491973621964e-246),
		Fy:    -libc.Float64FromFloat64(815.3050133490987),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.2334257935615034e+40),
		Fy:    float64(134.57018729671395),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0408740281300768e+235),
		Fy:    float64(780.7108977754392),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.375955995171274e-220),
		Fy:    -libc.Float64FromFloat64(729.0688847762581),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0599925623698468e+208),
		Fy:    float64(691.0450978784638),
		Fdy:   float32(-libc.Float64FromFloat64(2.8301666787911017e-08)),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0599925623698468e+208),
		Fy:    float64(691.0450978784639),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0599925623698468e+208),
		Fy:    float64(691.0450978784638),
		Fdy:   float32(-libc.Float64FromFloat64(2.8301666787911017e-08)),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.0383758083862713e-284),
		Fy:    -libc.Float64FromFloat64(941.4138037749177),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7079945898055961e-164),
		Fy:    -libc.Float64FromFloat64(544.023904156401),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.917002549612509e-238),
		Fy:    -libc.Float64FromFloat64(787.4623277534897),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0092960108706409e+08),
		Fy:    float64(26.588774115688786),
		Fdy:   float32(-libc.Float64FromFloat64(1.9297164044473902e-07)),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0092960108706409e+08),
		Fy:    float64(26.58877411568879),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0092960108706409e+08),
		Fy:    float64(26.588774115688786),
		Fdy:   float32(-libc.Float64FromFloat64(1.9297164044473902e-07)),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.9242386718246215e+51),
		Fy:    float64(171.9849626034004),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0624372970169214e+71),
		Fy:    float64(235.9442724357099),
		Fdy:   float32(-libc.Float64FromFloat64(8.496215997411127e-08)),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0624372970169214e+71),
		Fy:    float64(235.94427243570993),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0624372970169214e+71),
		Fy:    float64(235.9442724357099),
		Fdy:   float32(-libc.Float64FromFloat64(8.496215997411127e-08)),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.6177582149971135e+187),
		Fy:    float64(621.8945497471989),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.826035639706036e+302),
		Fy:    float64(1006.5188941278358),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.30718214634746e+153),
		Fy:    float64(509.4611304235034),
		Fdy:   float32(-libc.Float64FromFloat64(1.0074347756017232e-07)),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.30718214634746e+153),
		Fy:    float64(509.46113042350345),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.30718214634746e+153),
		Fy:    float64(509.4611304235034),
		Fdy:   float32(-libc.Float64FromFloat64(1.0074347756017232e-07)),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.037544907890825e+79),
		Fy:    float64(263.45915135307837),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.037544907890825e+79),
		Fy:    float64(263.4591513530784),
		Fdy:   float32(1.8193054529547226e-07),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.037544907890825e+79),
		Fy:    float64(263.45915135307837),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.806974810332374e-54),
		Fy:    -libc.Float64FromFloat64(176.6171033529662),
		Fdy:   float32(-libc.Float64FromFloat64(2.3219338629587583e-07)),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.806974810332374e-54),
		Fy:    -libc.Float64FromFloat64(176.61710335296618),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.806974810332374e-54),
		Fy:    -libc.Float64FromFloat64(176.61710335296618),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.6390272285993895e+254),
		Fy:    float64(845.6332889474085),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(5.50941698321517e-77),
		Fy:    -libc.Float64FromFloat64(253.32656364813428),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(5.50941698321517e-77),
		Fy:    -libc.Float64FromFloat64(253.32656364813425),
		Fdy:   float32(1.3605047399778414e-07),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.50941698321517e-77),
		Fy:    -libc.Float64FromFloat64(253.32656364813425),
		Fdy:   float32(1.3605047399778414e-07),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(7.779199869124979e-262),
		Fy:    -libc.Float64FromFloat64(867.3855390862811),
		Fdy:   float32(-libc.Float64FromFloat64(1.2960435924469493e-07)),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(7.779199869124979e-262),
		Fy:    -libc.Float64FromFloat64(867.385539086281),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(7.779199869124979e-262),
		Fy:    -libc.Float64FromFloat64(867.385539086281),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.795206660935492e-126),
		Fy:    -libc.Float64FromFloat64(417.07998500451566),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.652444721325411e-220),
		Fy:    -libc.Float64FromFloat64(728.9553184367901),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.4557596307925752e-144),
		Fy:    -libc.Float64FromFloat64(477.061476306832),
		Fdy:   float32(-libc.Float64FromFloat64(7.403230917191195e-09)),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.4557596307925752e-144),
		Fy:    -libc.Float64FromFloat64(477.06147630683193),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.4557596307925752e-144),
		Fy:    -libc.Float64FromFloat64(477.06147630683193),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.406811894293943e+26),
		Fy:    float64(88.1385527588963),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.820865622644905e-28),
		Fy:    -libc.Float64FromFloat64(91.08008713690337),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.820865622644905e-28),
		Fy:    -libc.Float64FromFloat64(91.08008713690336),
		Fdy:   float32(6.594058987730023e-08),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.820865622644905e-28),
		Fy:    -libc.Float64FromFloat64(91.08008713690336),
		Fdy:   float32(6.594058987730023e-08),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.1916394926990522e+204),
		Fy:    float64(677.9262791995662),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.1916394926990522e+204),
		Fy:    float64(677.9262791995664),
		Fdy:   float32(2.174545414845852e-07),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.1916394926990522e+204),
		Fy:    float64(677.9262791995662),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.130202739343946e+128),
		Fy:    float64(428.0407392442554),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.6757657653906656e-281),
		Fy:    -libc.Float64FromFloat64(931.5837498281355),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.6757657653906656e-281),
		Fy:    -libc.Float64FromFloat64(931.5837498281354),
		Fdy:   float32(6.588714285271635e-08),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.6757657653906656e-281),
		Fy:    -libc.Float64FromFloat64(931.5837498281354),
		Fdy:   float32(6.588714285271635e-08),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.2846327382876345e+265),
		Fy:    float64(880.6723011136111),
		Fdy:   float32(-libc.Float64FromFloat64(8.488806102491253e-09)),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.2846327382876345e+265),
		Fy:    float64(880.6723011136112),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.2846327382876345e+265),
		Fy:    float64(880.6723011136111),
		Fdy:   float32(-libc.Float64FromFloat64(8.488806102491253e-09)),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.0905011568001617e+149),
		Fy:    float64(496.99956374718784),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.9601049289002103e-269),
		Fy:    -libc.Float64FromFloat64(892.0330092077243),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.135444806421597e-202),
		Fy:    -libc.Float64FromFloat64(670.8462175884038),
		Fdy:   float32(-libc.Float64FromFloat64(5.519743240256503e-08)),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.135444806421597e-202),
		Fy:    -libc.Float64FromFloat64(670.8462175884036),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.135444806421597e-202),
		Fy:    -libc.Float64FromFloat64(670.8462175884036),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.3976759541683983e+241),
		Fy:    float64(801.8463075595906),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.174746211998537e+193),
		Fy:    float64(643.193810810292),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1555153969616633e-107),
		Fy:    -libc.Float64FromFloat64(355.2377696694536),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.547880857730658e-232),
		Fy:    -libc.Float64FromFloat64(769.3380201970195),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.9085291355151747e-59),
		Fy:    -libc.Float64FromFloat64(195.0612963875188),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(9.26504053432982e-82),
		Fy:    -libc.Float64FromFloat64(269.1863064927412),
		Fdy:   float32(-libc.Float64FromFloat64(1.4716714247242635e-07)),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(9.26504053432982e-82),
		Fy:    -libc.Float64FromFloat64(269.18630649274115),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(9.26504053432982e-82),
		Fy:    -libc.Float64FromFloat64(269.18630649274115),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(3.379198992686702e-152),
		Fy:    -libc.Float64FromFloat64(503.17638911312633),
		Fdy:   float32(-libc.Float64FromFloat64(3.716994712021915e-08)),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(3.379198992686702e-152),
		Fy:    -libc.Float64FromFloat64(503.1763891131263),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.379198992686702e-152),
		Fy:    -libc.Float64FromFloat64(503.1763891131263),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.904014340653078e-270),
		Fy:    -libc.Float64FromFloat64(895.9915412748121),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.3622128378041964e-107),
		Fy:    -libc.Float64FromFloat64(354.2061671940506),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.1193559985285155e-149),
		Fy:    -libc.Float64FromFloat64(493.883660193653),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.2254202729281516e-25),
		Fy:    -libc.Float64FromFloat64(82.75492574811217),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.8018563536790782e+68),
		Fy:    float64(226.74059445456018),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.8018563536790782e+68),
		Fy:    float64(226.7405944545602),
		Fdy:   float32(8.160436237858448e-08),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.8018563536790782e+68),
		Fy:    float64(226.74059445456018),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.942567509238741e+133),
		Fy:    float64(443.7955720787535),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.0029375735398722e+67),
		Fy:    float64(223.57129981417657),
		Fdy:   float32(-libc.Float64FromFloat64(8.277142882207045e-08)),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.0029375735398722e+67),
		Fy:    float64(223.5712998141766),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.0029375735398722e+67),
		Fy:    float64(223.57129981417657),
		Fdy:   float32(-libc.Float64FromFloat64(8.277142882207045e-08)),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.6556598612884334e+212),
		Fy:    float64(706.4677417760943),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.4940415843720632e-186),
		Fy:    -libc.Float64FromFloat64(617.2994053451442),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.1775738400147423e-155),
		Fy:    -libc.Float64FromFloat64(513.7761330671019),
		Fdy:   float32(-libc.Float64FromFloat64(2.3563076467780775e-07)),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.1775738400147423e-155),
		Fy:    -libc.Float64FromFloat64(513.7761330671018),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.1775738400147423e-155),
		Fy:    -libc.Float64FromFloat64(513.7761330671018),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.881320041462119e+118),
		Fy:    float64(394.77020053992067),
		Fdy:   float32(-libc.Float64FromFloat64(1.5205839076770644e-07)),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.881320041462119e+118),
		Fy:    float64(394.7702005399207),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.881320041462119e+118),
		Fy:    float64(394.77020053992067),
		Fdy:   float32(-libc.Float64FromFloat64(1.5205839076770644e-07)),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.630604209494402e-165),
		Fy:    -libc.Float64FromFloat64(547.4127290118838),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.630604209494402e-165),
		Fy:    -libc.Float64FromFloat64(547.4127290118837),
		Fdy:   float32(7.962990267174064e-09),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.630604209494402e-165),
		Fy:    -libc.Float64FromFloat64(547.4127290118837),
		Fdy:   float32(7.962990267174064e-09),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.8042411565838117e-226),
		Fy:    -libc.Float64FromFloat64(749.9043572610753),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.8042411565838117e-226),
		Fy:    -libc.Float64FromFloat64(749.9043572610752),
		Fdy:   float32(1.7788889294934052e-07),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.8042411565838117e-226),
		Fy:    -libc.Float64FromFloat64(749.9043572610752),
		Fdy:   float32(1.7788889294934052e-07),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.148411421638898e+29),
		Fy:    float64(98.9561284580925),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.148411421638898e+29),
		Fy:    float64(98.95612845809251),
		Fdy:   float32(3.908802526098043e-08),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.148411421638898e+29),
		Fy:    float64(98.9561284580925),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(6.039776333415696e-291),
		Fy:    -libc.Float64FromFloat64(964.0865804879447),
		Fdy:   float32(-libc.Float64FromFloat64(7.731304663138872e-08)),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(6.039776333415696e-291),
		Fy:    -libc.Float64FromFloat64(964.0865804879446),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.039776333415696e-291),
		Fy:    -libc.Float64FromFloat64(964.0865804879446),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.803219485789802e-244),
		Fy:    -libc.Float64FromFloat64(809.6998801418869),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.803219485789802e-244),
		Fy:    -libc.Float64FromFloat64(809.6998801418868),
		Fdy:   float32(1.7722263123687299e-07),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.803219485789802e-244),
		Fy:    -libc.Float64FromFloat64(809.6998801418868),
		Fdy:   float32(1.7722263123687299e-07),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.3636963773910445e-245),
		Fy:    -libc.Float64FromFloat64(811.4491556744474),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.146980807032071e-90),
		Fy:    -libc.Float64FromFloat64(296.3536505593706),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.375083840701094e+300),
		Fy:    float64(997.0379480505742),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2.6020997997220898e+11),
		Fy:    float64(37.92088533933053),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2.6020997997220898e+11),
		Fy:    float64(37.92088533933054),
		Fdy:   float32(4.411322507280602e-08),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.6020997997220898e+11),
		Fy:    float64(37.92088533933053),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.5619270909007666e+55),
		Fy:    float64(184.06327463785195),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	432: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.345239849338305),
		Fy:    float64(2.1194358133804485),
		Fdy:   float32(-libc.Float64FromFloat64(0.10164877772331238)),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	434: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	435: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.267056966972586),
		Fy:    float64(3.2121112403298744),
		Fdy:   float32(-libc.Float64FromFloat64(0.15739446878433228)),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6619858980995045),
		Fy:    -libc.Float64FromFloat64(0.5951276104207402),
		Fdy:   float32(0.3321485221385956),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	438: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5617597462207241),
		Fy:    -libc.Float64FromFloat64(0.8319748453044644),
		Fdy:   float32(0.057555437088012695),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7741522965913037),
		Fy:    -libc.Float64FromFloat64(0.36931068365537134),
		Fdy:   float32(-libc.Float64FromFloat64(0.19838279485702515)),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts + 23,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	441: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	442: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	443: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.888609052210118e-31),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	444: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	445: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	446: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	447: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	448: {
		Ffile: __ccgo_ts + 47,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+46)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:12:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:13:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:17:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:20:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(16164)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:21:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:23:3:
		if (*d_d)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:24:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:25:3:
		libc.Xfesetround(tls, (*d_d)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:26:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:27:3:
		y = libc.Xlog2(tls, (*d_d)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:28:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:30:3:
		if !(checkexcept(tls, e, (*d_d)(unsafe.Pointer(p)).Fe, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:31:4:
			libc.Xprintf(tls, __ccgo_ts+71, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, estr(tls, (*d_d)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:33:4:
			libc.Xprintf(tls, __ccgo_ts+120, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:34:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:36:3:
		d = ulperr(tls, y, (*d_d)(unsafe.Pointer(p)).Fy, (*d_d)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:37:3:
		if !(checkulp(tls, d, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:38:4:
			libc.Xprintf(tls, __ccgo_ts+129, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*d_d)(unsafe.Pointer(p)).Fdy), float64((*d_d)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:40:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log2.c:43:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+186, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+196, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 240,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 248,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 256,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 266,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 276,
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
				v2 = __ccgo_ts + 285
			} else {
				v2 = __ccgo_ts + 46
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+287, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 285
		} else {
			v3 = __ccgo_ts + 46
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+292, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 46
	} else {
		v4 = __ccgo_ts + 297
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+299, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 302
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 305
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 308
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 311
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 314
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+317, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+325, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+330, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+373, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+422) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+430) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+442) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+454) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+466) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+475) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+46) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+475) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+481, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/log2.h\x00src/math/sanity/log2.h\x00\x00src/math/special/log2.h\x00%s:%d: bad fp exception: %s log2(%a)=%a, want %s\x00 got %s\n\x00%s:%d: %s log2(%a) want %a got %a ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
