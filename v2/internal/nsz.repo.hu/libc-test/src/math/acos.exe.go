// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/acos.exe.go src/math/acos.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:5:19:
var t = [440]d_d{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(15),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(16),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(17),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948968),
		Fdy:   float32(0.7242340445518494),
		Fe:    int32(FE_INEXACT),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(18),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    float64(1.5707963267948968),
		Fdy:   float32(0.7242340445518494),
		Fe:    int32(FE_INEXACT),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.141592653589793),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.141592653589793),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.141592653589793),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.1415926535897936),
		Fdy:   float32(0.7242340445518494),
		Fe:    int32(FE_INEXACT),
	},
	16: {
		Ffile: __ccgo_ts,
		Fline: int32(27),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	17: {
		Ffile: __ccgo_ts,
		Fline: int32(28),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	18: {
		Ffile: __ccgo_ts,
		Fline: int32(29),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	19: {
		Ffile: __ccgo_ts,
		Fline: int32(30),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	20: {
		Ffile: __ccgo_ts,
		Fline: int32(31),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	21: {
		Ffile: __ccgo_ts,
		Fline: int32(32),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	22: {
		Ffile: __ccgo_ts,
		Fline: int32(33),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	23: {
		Ffile: __ccgo_ts,
		Fline: int32(34),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.18499994277954102),
		Fy:    float64(1.3847245226894906),
		Fdy:   float32(0.40587103366851807),
		Fe:    int32(FE_INEXACT),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.2997171878814697),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(-libc.Float64FromFloat64(0.25404220819473267)),
		Fe:    int32(FE_INEXACT),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.40296268463134766),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(0.13958610594272614),
		Fe:    int32(FE_INEXACT),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.4932067394256592),
		Fy:    float64(1.055024129940415),
		Fdy:   float32(0.29588502645492554),
		Fe:    int32(FE_INEXACT),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5696849822998047),
		Fy:    float64(0.9646738195527598),
		Fdy:   float32(0.0717492625117302),
		Fe:    int32(FE_INEXACT),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6396622657775879),
		Fy:    float64(0.8767375242854261),
		Fdy:   float32(0.18386588990688324),
		Fe:    int32(FE_INEXACT),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.696256160736084),
		Fy:    float64(0.8006278768874553),
		Fdy:   float32(-libc.Float64FromFloat64(0.42917510867118835)),
		Fe:    int32(FE_INEXACT),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7417607307434082),
		Fy:    float64(0.7351044138925751),
		Fdy:   float32(0.33165794610977173),
		Fe:    int32(FE_INEXACT),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7799997329711914),
		Fy:    float64(0.6761309362744347),
		Fdy:   float32(0.15576781332492828),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.18499994277954104),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.46693339943885803)),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(48),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.2997171878814698),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(0.008004593662917614),
		Fe:    int32(FE_INEXACT),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.4029626846313477),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(0.41274553537368774),
		Fe:    int32(FE_INEXACT),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.49320673942565924),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.4167296290397644)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5696849822998048),
		Fy:    float64(0.9646738195527597),
		Fdy:   float32(0.28849685192108154),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.639662265777588),
		Fy:    float64(0.876737524285426),
		Fdy:   float32(0.4848378300666809),
		Fe:    int32(FE_INEXACT),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6962561607360841),
		Fy:    float64(0.8006278768874552),
		Fdy:   float32(-libc.Float64FromFloat64(0.03601657226681709)),
		Fe:    int32(FE_INEXACT),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7417607307434083),
		Fy:    float64(0.7351044138925749),
		Fdy:   float32(-libc.Float64FromFloat64(0.17728380858898163)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7799997329711915),
		Fy:    float64(0.6761309362744344),
		Fdy:   float32(-libc.Float64FromFloat64(0.24622610211372375)),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.184999942779541),
		Fy:    float64(1.3847245226894906),
		Fdy:   float32(0.2786754369735718),
		Fe:    int32(FE_INEXACT),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.29971718788146967),
		Fy:    float64(1.2664001266200802),
		Fdy:   float32(0.48391100764274597),
		Fe:    int32(FE_INEXACT),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.4029626846313476),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(-libc.Float64FromFloat64(0.13357332348823547)),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.4932067394256591),
		Fy:    float64(1.055024129940415),
		Fdy:   float32(0.008499687537550926),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5696849822998046),
		Fy:    float64(0.9646738195527599),
		Fdy:   float32(-libc.Float64FromFloat64(0.14499834179878235)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6396622657775878),
		Fy:    float64(0.8767375242854262),
		Fdy:   float32(-libc.Float64FromFloat64(0.11710604280233383)),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6962561607360839),
		Fy:    float64(0.8006278768874555),
		Fdy:   float32(0.1776663362979889),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7417607307434081),
		Fy:    float64(0.7351044138925752),
		Fdy:   float32(-libc.Float64FromFloat64(0.1594003140926361)),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7799997329711913),
		Fy:    float64(0.6761309362744348),
		Fdy:   float32(-libc.Float64FromFloat64(0.4422382712364197)),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.18499994277954102),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.5941289663314819)),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.2997171878814697),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(-libc.Float64FromFloat64(0.25404220819473267)),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.40296268463134766),
		Fy:    float64(1.1560446376615316),
		Fdy:   float32(-libc.Float64FromFloat64(0.8604139089584351)),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.4932067394256592),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.7041149735450745)),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.5696849822998047),
		Fy:    float64(0.9646738195527597),
		Fdy:   float32(-libc.Float64FromFloat64(0.9282507300376892)),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.6396622657775879),
		Fy:    float64(0.876737524285426),
		Fdy:   float32(-libc.Float64FromFloat64(0.8161340951919556)),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.696256160736084),
		Fy:    float64(0.8006278768874553),
		Fdy:   float32(-libc.Float64FromFloat64(0.42917510867118835)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7417607307434082),
		Fy:    float64(0.735104413892575),
		Fdy:   float32(-libc.Float64FromFloat64(0.6683420538902283)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7799997329711914),
		Fy:    float64(0.6761309362744345),
		Fdy:   float32(-libc.Float64FromFloat64(0.8442322015762329)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.18499994277954104),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.46693339943885803)),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.2997171878814698),
		Fy:    float64(1.2664001266200797),
		Fdy:   float32(-libc.Float64FromFloat64(0.9919953942298889)),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.4029626846313477),
		Fy:    float64(1.1560446376615316),
		Fdy:   float32(-libc.Float64FromFloat64(0.5872544646263123)),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.49320673942565924),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.4167296290397644)),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.5696849822998048),
		Fy:    float64(0.9646738195527595),
		Fdy:   float32(-libc.Float64FromFloat64(0.7115031480789185)),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.639662265777588),
		Fy:    float64(0.8767375242854258),
		Fdy:   float32(-libc.Float64FromFloat64(0.5151621699333191)),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.6962561607360841),
		Fy:    float64(0.8006278768874552),
		Fdy:   float32(-libc.Float64FromFloat64(0.03601657226681709)),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7417607307434083),
		Fy:    float64(0.7351044138925749),
		Fdy:   float32(-libc.Float64FromFloat64(0.17728380858898163)),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7799997329711915),
		Fy:    float64(0.6761309362744344),
		Fdy:   float32(-libc.Float64FromFloat64(0.24622610211372375)),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.184999942779541),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.7213245630264282)),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.29971718788146967),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(-libc.Float64FromFloat64(0.5160889625549316)),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.4029626846313476),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(-libc.Float64FromFloat64(0.13357332348823547)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.4932067394256591),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.9915003180503845)),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.5696849822998046),
		Fy:    float64(0.9646738195527599),
		Fdy:   float32(-libc.Float64FromFloat64(0.14499834179878235)),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.6396622657775878),
		Fy:    float64(0.8767375242854262),
		Fdy:   float32(-libc.Float64FromFloat64(0.11710604280233383)),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.6962561607360839),
		Fy:    float64(0.8006278768874554),
		Fdy:   float32(-libc.Float64FromFloat64(0.8223336338996887)),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7417607307434081),
		Fy:    float64(0.7351044138925752),
		Fdy:   float32(-libc.Float64FromFloat64(0.1594003140926361)),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.7799997329711913),
		Fy:    float64(0.6761309362744348),
		Fdy:   float32(-libc.Float64FromFloat64(0.4422382712364197)),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.18499994277954102),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.5941289663314819)),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.2997171878814697),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(-libc.Float64FromFloat64(0.25404220819473267)),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.40296268463134766),
		Fy:    float64(1.1560446376615316),
		Fdy:   float32(-libc.Float64FromFloat64(0.8604139089584351)),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.4932067394256592),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.7041149735450745)),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5696849822998047),
		Fy:    float64(0.9646738195527597),
		Fdy:   float32(-libc.Float64FromFloat64(0.9282507300376892)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.6396622657775879),
		Fy:    float64(0.876737524285426),
		Fdy:   float32(-libc.Float64FromFloat64(0.8161340951919556)),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.696256160736084),
		Fy:    float64(0.8006278768874553),
		Fdy:   float32(-libc.Float64FromFloat64(0.42917510867118835)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7417607307434082),
		Fy:    float64(0.735104413892575),
		Fdy:   float32(-libc.Float64FromFloat64(0.6683420538902283)),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7799997329711914),
		Fy:    float64(0.6761309362744345),
		Fdy:   float32(-libc.Float64FromFloat64(0.8442322015762329)),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.18499994277954104),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.46693339943885803)),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.2997171878814698),
		Fy:    float64(1.2664001266200797),
		Fdy:   float32(-libc.Float64FromFloat64(0.9919953942298889)),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.4029626846313477),
		Fy:    float64(1.1560446376615316),
		Fdy:   float32(-libc.Float64FromFloat64(0.5872544646263123)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.49320673942565924),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.4167296290397644)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5696849822998048),
		Fy:    float64(0.9646738195527595),
		Fdy:   float32(-libc.Float64FromFloat64(0.7115031480789185)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.639662265777588),
		Fy:    float64(0.8767375242854258),
		Fdy:   float32(-libc.Float64FromFloat64(0.5151621699333191)),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.6962561607360841),
		Fy:    float64(0.8006278768874552),
		Fdy:   float32(-libc.Float64FromFloat64(0.03601657226681709)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7417607307434083),
		Fy:    float64(0.7351044138925749),
		Fdy:   float32(-libc.Float64FromFloat64(0.17728380858898163)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7799997329711915),
		Fy:    float64(0.6761309362744344),
		Fdy:   float32(-libc.Float64FromFloat64(0.24622610211372375)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.184999942779541),
		Fy:    float64(1.3847245226894904),
		Fdy:   float32(-libc.Float64FromFloat64(0.7213245630264282)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.29971718788146967),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(-libc.Float64FromFloat64(0.5160889625549316)),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.4029626846313476),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(-libc.Float64FromFloat64(0.13357332348823547)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.4932067394256591),
		Fy:    float64(1.0550241299404148),
		Fdy:   float32(-libc.Float64FromFloat64(0.9915003180503845)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5696849822998046),
		Fy:    float64(0.9646738195527599),
		Fdy:   float32(-libc.Float64FromFloat64(0.14499834179878235)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.6396622657775878),
		Fy:    float64(0.8767375242854262),
		Fdy:   float32(-libc.Float64FromFloat64(0.11710604280233383)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.6962561607360839),
		Fy:    float64(0.8006278768874554),
		Fdy:   float32(-libc.Float64FromFloat64(0.8223336338996887)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7417607307434081),
		Fy:    float64(0.7351044138925752),
		Fdy:   float32(-libc.Float64FromFloat64(0.1594003140926361)),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7799997329711913),
		Fy:    float64(0.6761309362744348),
		Fdy:   float32(-libc.Float64FromFloat64(0.4422382712364197)),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.18499994277954102),
		Fy:    float64(1.3847245226894906),
		Fdy:   float32(0.40587103366851807),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.2997171878814697),
		Fy:    float64(1.2664001266200802),
		Fdy:   float32(0.7459577918052673),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.40296268463134766),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(0.13958610594272614),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.4932067394256592),
		Fy:    float64(1.055024129940415),
		Fdy:   float32(0.29588502645492554),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.5696849822998047),
		Fy:    float64(0.9646738195527598),
		Fdy:   float32(0.0717492625117302),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.6396622657775879),
		Fy:    float64(0.8767375242854261),
		Fdy:   float32(0.18386588990688324),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.696256160736084),
		Fy:    float64(0.8006278768874554),
		Fdy:   float32(0.5708248615264893),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7417607307434082),
		Fy:    float64(0.7351044138925751),
		Fdy:   float32(0.33165794610977173),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7799997329711914),
		Fy:    float64(0.6761309362744347),
		Fdy:   float32(0.15576781332492828),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.18499994277954104),
		Fy:    float64(1.3847245226894906),
		Fdy:   float32(0.5330666303634644),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.2997171878814698),
		Fy:    float64(1.26640012662008),
		Fdy:   float32(0.008004593662917614),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.4029626846313477),
		Fy:    float64(1.1560446376615319),
		Fdy:   float32(0.41274553537368774),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.49320673942565924),
		Fy:    float64(1.055024129940415),
		Fdy:   float32(0.5832703709602356),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.5696849822998048),
		Fy:    float64(0.9646738195527597),
		Fdy:   float32(0.28849685192108154),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.639662265777588),
		Fy:    float64(0.876737524285426),
		Fdy:   float32(0.4848378300666809),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.6962561607360841),
		Fy:    float64(0.8006278768874553),
		Fdy:   float32(0.963983416557312),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7417607307434083),
		Fy:    float64(0.735104413892575),
		Fdy:   float32(0.8227161765098572),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7799997329711915),
		Fy:    float64(0.6761309362744345),
		Fdy:   float32(0.7537739276885986),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.184999942779541),
		Fy:    float64(1.3847245226894906),
		Fdy:   float32(0.2786754369735718),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.29971718788146967),
		Fy:    float64(1.2664001266200802),
		Fdy:   float32(0.48391100764274597),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.4029626846313476),
		Fy:    float64(1.156044637661532),
		Fdy:   float32(0.8664266467094421),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.4932067394256591),
		Fy:    float64(1.055024129940415),
		Fdy:   float32(0.008499687537550926),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.5696849822998046),
		Fy:    float64(0.96467381955276),
		Fdy:   float32(0.85500168800354),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.6396622657775878),
		Fy:    float64(0.8767375242854263),
		Fdy:   float32(0.882893979549408),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.6962561607360839),
		Fy:    float64(0.8006278768874555),
		Fdy:   float32(0.1776663362979889),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7417607307434081),
		Fy:    float64(0.7351044138925753),
		Fdy:   float32(0.8405997157096863),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.7799997329711913),
		Fy:    float64(0.6761309362744349),
		Fdy:   float32(0.5577617287635803),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.7041817951240164),
		Fy:    float64(0.7895262097099214),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5810268069553325),
		Fy:    float64(0.9508065906755769),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999994944723088),
		Fy:    float64(0.0010055125396986123),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999994450449933),
		Fy:    float64(0.0010535227174157129),
		Fdy:   float32(-libc.Float64FromFloat64(1.0450786069871043e-10)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999988467939773),
		Fy:    float64(0.0015186877521875936),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999985810440652),
		Fy:    float64(0.0016846105011686459),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999002352099),
		Fy:    float64(0.0004466873442456241),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999998810600929),
		Fy:    float64(0.0004877292476431849),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999907836514),
		Fy:    float64(0.00013576706955194443),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999711271722),
		Fy:    float64(0.00024030325813117852),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999952347282),
		Fy:    float64(9.762450355901458e-05),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.999999999372269),
		Fy:    float64(3.543249859590427e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999989854845),
		Fy:    float64(4.504476599567963e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.0427071828189582e-07)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999982041595),
		Fy:    float64(5.99306345179794e-05),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999987929652),
		Fy:    float64(4.9133182385882546e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999986625495),
		Fy:    float64(5.1719444776489514e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999996432989),
		Fy:    float64(2.67095905289605e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999997656777),
		Fy:    float64(2.1648202665297976e-05),
		Fdy:   float32(0.49999868869781494),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999997568902),
		Fy:    float64(2.205038569093586e-05),
		Fdy:   float32(-libc.Float64FromFloat64(4.1020783214662515e-08)),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999510638),
		Fy:    float64(9.893046895190767e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999503638),
		Fy:    float64(9.963551982326335e-06),
		Fdy:   float32(0.4999978244304657),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999998903168),
		Fy:    float64(1.4811019644320092e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999926505),
		Fy:    float64(3.833915590730347e-06),
		Fdy:   float32(-libc.Float64FromFloat64(5.386696557252435e-06)),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999913249),
		Fy:    float64(4.165347685326783e-06),
		Fdy:   float32(0.49998655915260315),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999880514),
		Fy:    float64(4.888466683520942e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1.5353114577010274e-05)),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.999999999974781),
		Fy:    float64(7.101978466789915e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1.0627240953908768e-05)),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999746938),
		Fy:    float64(7.114239464455442e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999925494194031)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999829755),
		Fy:    float64(5.8351508852055995e-06),
		Fdy:   float32(0.499986857175827),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999800517),
		Fy:    float64(6.316370043296419e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999893307685852)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999796226),
		Fy:    float64(6.383943367691589e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999917447566986)),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999976776),
		Fy:    float64(2.1551633455085493e-06),
		Fdy:   float32(-libc.Float64FromFloat64(8.940717088989913e-05)),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999972641),
		Fy:    float64(2.339197554540882e-06),
		Fdy:   float32(0.4999288022518158),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999930042),
		Fy:    float64(3.740547645754203e-06),
		Fdy:   float32(-libc.Float64FromFloat64(5.414176848717034e-05)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9592593671550103),
		Fy:    float64(0.2864273428429076),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9496636469901064),
		Fy:    float64(0.3186358618570418),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9314247618019921),
		Fy:    float64(0.3724879997767474),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999983016),
		Fy:    float64(1.843051373712651e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.49995189905166626)),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999988983),
		Fy:    float64(1.4843680859785048e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.49985095858573914)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999987157),
		Fy:    float64(1.602688987225317e-06),
		Fdy:   float32(0.49990877509117126),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999998481),
		Fy:    float64(5.511415603431214e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.0004926337860524654)),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999996072),
		Fy:    float64(8.86337301609727e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.0006143858772702515)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999997032),
		Fy:    float64(7.704058858579923e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.9990929961204529)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999769),
		Fy:    float64(2.149076029934881e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.002944539301097393)),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999764),
		Fy:    float64(2.1747528790423948e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.4920680820941925)),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999741),
		Fy:    float64(2.2745635393967895e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.98872971534729)),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.999999999999973),
		Fy:    float64(2.3228611451566115e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.4875737428665161)),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999963),
		Fy:    float64(8.560065398421929e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.023720979690551758)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999958),
		Fy:    float64(9.185692672385242e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.931171178817749)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999991),
		Fy:    float64(4.2146848510894035e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.036028336733579636)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999986),
		Fy:    float64(5.372690074837192e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.3553924858570099)),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.989604388279725),
		Fy:    float64(0.1443168269856668),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9750190558866111),
		Fy:    float64(0.22398951097064657),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9731132151474613),
		Fy:    float64(0.23241402466991679),
		Fdy:   float32(-libc.Float64FromFloat64(1.177841749507762e-14)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9708234702904849),
		Fy:    float64(0.24215500141800386),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9978992840741051),
		Fy:    float64(0.0648298054821695),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9976669282060323),
		Fy:    float64(0.06832246428239874),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9945050357993978),
		Fy:    float64(0.10488092320721203),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9981152108515527),
		Fy:    float64(0.06140653903124788),
		Fdy:   float32(-libc.Float64FromFloat64(8.740814197655614e-14)),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9996241129231636),
		Fy:    float64(0.02741935913594291),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999543854637087),
		Fy:    float64(0.009551427440593986),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999418310848611),
		Fy:    float64(0.0107860538738448),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999251622599064),
		Fy:    float64(0.012234269374193679),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999097531445593),
		Fy:    float64(0.01343489582312915),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9998953140274018),
		Fy:    float64(0.01446981680427433),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.999987594074304),
		Fy:    float64(0.004981154755116067),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999865990373268),
		Fy:    float64(0.005177063376913478),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999800240379935),
		Fy:    float64(0.006320763959111169),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999980093110205),
		Fy:    float64(0.0019953393896886165),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999969400023888),
		Fy:    float64(0.002473863040585165),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999926015825316),
		Fy:    float64(0.0038466678024550033),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.10089712900129905),
		Fy:    float64(1.4697272160826822),
		Fdy:   float32(-libc.Float64FromFloat64(4.839490980401923e-17)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0011644316078763548),
		Fy:    float64(1.569631894923878),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.0013321145347443867),
		Fy:    float64(1.5694642118661724),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0014353006688517845),
		Fy:    float64(1.5693610256332369),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0007061489590817238),
		Fy:    float64(1.5700901777771283),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0007119262749058442),
		Fy:    float64(1.570084400459852),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0007773030110915965),
		Fy:    float64(1.5700190237055305),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0008130847008781175),
		Fy:    float64(1.5699832420044293),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.0007740929382051827),
		Fy:    float64(1.5700222337793828),
		Fdy:   float32(-libc.Float64FromFloat64(2.4663092735916945e-16)),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.0007600242595518564),
		Fy:    float64(1.570036302462175),
		Fdy:   float32(-libc.Float64FromFloat64(5.067969647594298e-16)),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.00030893060976761445),
		Fy:    float64(1.570487396180215),
		Fdy:   float32(-libc.Float64FromFloat64(4.89326698464073e-16)),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0003049070403581824),
		Fy:    float64(1.570491419749814),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.00018597152086141465),
		Fy:    float64(1.5706103552729633),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.00018470235255988688),
		Fy:    float64(1.5706116244412867),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.13047359866068e-05),
		Fy:    float64(1.5707050220587833),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.8710818520541533e-05),
		Fy:    float64(1.570747615976357),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.0752934606245854e-05),
		Fy:    float64(1.5707455738602685),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.2363282679749265e-05),
		Fy:    float64(1.5707539635122043),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.70703223514168e-05),
		Fy:    float64(1.5707392564725142),
		Fdy:   float32(-libc.Float64FromFloat64(7.727158646489841e-17)),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7612756459656156e-05),
		Fy:    float64(1.5707787140384362),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.768307570607096e-05),
		Fy:    float64(1.5707786437191897),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.6603869356077167e-05),
		Fy:    float64(1.5707697229255375),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.6243030003361103e-05),
		Fy:    float64(1.5707700837648901),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.9535080732431775e-05),
		Fy:    float64(1.5707667917141597),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(7.968711453231403e-06),
		Fy:    float64(1.5707883580834434),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.2792692938013847e-05),
		Fy:    float64(1.5707835341019583),
		Fdy:   float32(-libc.Float64FromFloat64(3.999870894925078e-16)),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.359662869235591e-05),
		Fy:    float64(1.5707827301662036),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.749976313877041e-06),
		Fy:    float64(1.5707915768185825),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(6.130505625168775e-06),
		Fy:    float64(1.5707901962892712),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.8045065858644635e-06),
		Fy:    float64(1.5707895222883106),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.9205535580258834e-06),
		Fy:    float64(1.5707944062413386),
		Fdy:   float32(-libc.Float64FromFloat64(8.448404024703326e-16)),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.95617201032453e-06),
		Fy:    float64(1.5707943706228864),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.230798841266742e-06),
		Fy:    float64(1.5707940959960551),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0872421703555626e-06),
		Fy:    float64(1.5707952395527263),
		Fdy:   float32(-libc.Float64FromFloat64(2.797534890170588e-16)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0420373147190613e-06),
		Fy:    float64(1.570795284757582),
		Fdy:   float32(-libc.Float64FromFloat64(1.8203251550219187e-16)),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0996399226247995e-06),
		Fy:    float64(1.570795227154974),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.2186801609580316e-06),
		Fy:    float64(1.5707951081147355),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.142433787631935e-06),
		Fy:    float64(1.570795184361109),
		Fdy:   float32(-libc.Float64FromFloat64(2.7027725498009214e-16)),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.897309376224374e-06),
		Fy:    float64(1.5707944294855203),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.854653497589467e-07),
		Fy:    float64(1.5707958413295469),
		Fdy:   float32(-libc.Float64FromFloat64(7.315395813797029e-16)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(5.917033355008728e-07),
		Fy:    float64(1.570795735091561),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.732044728838702e-07),
		Fy:    float64(1.5707960535904235),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(3.825123296498587e-07),
		Fy:    float64(1.5707959442825667),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.39463105189764e-07),
		Fy:    float64(1.5707958873317915),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.47008093934653e-07),
		Fy:    float64(1.5707958797868027),
		Fdy:   float32(-libc.Float64FromFloat64(3.7284248932964524e-16)),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.708464931251013e-07),
		Fy:    float64(1.5707961559484034),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.8734007129500818e-07),
		Fy:    float64(1.570796139454825),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2.147961724653943e-07),
		Fy:    float64(1.5707961119987242),
		Fdy:   float32(-libc.Float64FromFloat64(8.326898736606473e-16)),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.1222300261865323e-07),
		Fy:    float64(1.570796214571894),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(3.9869993479106e-08),
		Fy:    float64(1.570796286924903),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.672153119990271e-08),
		Fy:    float64(1.5707962700733653),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.03506916689894219),
		Fy:    float64(1.5357199676329873),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.04164411458587163),
		Fy:    float64(1.5291401660444783),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.04221442248036426),
		Fy:    float64(1.5285693561614493),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.04823344744695509),
		Fy:    float64(1.5225441574997067),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.01931505551665312),
		Fy:    float64(1.5514800700942057),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.02042130345747097),
		Fy:    float64(1.5503736036895392),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.023821280197581677),
		Fy:    float64(1.546972793110796),
		Fdy:   float32(-libc.Float64FromFloat64(7.866507213016198e-16)),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.027047901364289167),
		Fy:    float64(1.5437451263533442),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.008200241427058575),
		Fy:    float64(1.562595993462273),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.011810850037121887),
		Fy:    float64(1.558985202145795),
		Fdy:   float32(-libc.Float64FromFloat64(2.0202659587323162e-16)),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.014353115913686952),
		Fy:    float64(1.5564427180166498),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.004122572793634592),
		Fy:    float64(1.5666737423235684),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.00437423835227992),
		Fy:    float64(1.5664220744930788),
		Fdy:   float32(-libc.Float64FromFloat64(3.448617617687243e-16)),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.005597999702486113),
		Fy:    float64(1.5651982978540182),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.005697304741148478),
		Fy:    float64(1.5650989912315618),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.005986059230534275),
		Fy:    float64(1.5648102318141373),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.0026939248540772324),
		Fy:    float64(1.5681023986824028),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.0033899170287759743),
		Fy:    float64(1.567406403273527),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.0035919827004359003),
		Fy:    float64(1.5672043363702521),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	273: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.141592653589793),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(41),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999999),
		Fy:    float64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.0416666679084301)),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(42),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999999),
		Fy:    float64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.0416666679084301)),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(43),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999999),
		Fy:    float64(1.490116119384766e-08),
		Fdy:   float32(0.9583333134651184),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(44),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999999),
		Fy:    float64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.0416666679084301)),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(45),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(3.141592638688632),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(46),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(3.141592638688632),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(47),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(3.1415926386886324),
		Fdy:   float32(0.7242340445518494),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(48),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9999999999999999),
		Fy:    float64(3.141592638688632),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(50),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(5.684341886080802e-14),
		Fy:    float64(1.5707963267948397),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.684341886080802e-14),
		Fy:    float64(1.5707963267949534),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(52),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(6.776263578034403e-21),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27573543787002563)),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.776263578034403e-21),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27579647302627563)),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(54),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2.2250738585072014e-308),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.2687083954988399),
		Fy:    float64(1.8428481854794552),
		Fdy:   float32(-libc.Float64FromFloat64(0.2862209379673004)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.712225905636645),
		Fy:    float64(0.7781321670027921),
		Fdy:   float32(0.08036109805107117),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.8804502671006565),
		Fy:    float64(0.4939853089730004),
		Fdy:   float32(0.11051065474748611),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6678552546757298),
		Fy:    float64(2.302119780711071),
		Fdy:   float32(-libc.Float64FromFloat64(0.08916228264570236)),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.2709030829738149),
		Fy:    float64(1.8451273986765893),
		Fdy:   float32(0.09290396422147751),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.039291951227476314),
		Fy:    float64(1.610098395247898),
		Fdy:   float32(0.17877931892871857),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5783313839080494),
		Fy:    float64(0.9541144916925824),
		Fdy:   float32(-libc.Float64FromFloat64(0.49562546610832214)),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.8550888492150711),
		Fy:    float64(0.5450741471324799),
		Fdy:   float32(-libc.Float64FromFloat64(0.43663010001182556)),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(66),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.1680688934754805),
		Fy:    float64(1.7396666950051183),
		Fdy:   float32(-libc.Float64FromFloat64(0.042779482901096344)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.9580241162600612),
		Fy:    float64(2.850825123202514),
		Fdy:   float32(0.026712287217378616),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(68),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.2687083954988399),
		Fy:    float64(1.8428481854794552),
		Fdy:   float32(-libc.Float64FromFloat64(0.2862209379673004)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(69),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.712225905636645),
		Fy:    float64(0.778132167002792),
		Fdy:   float32(-libc.Float64FromFloat64(0.9196389317512512)),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(70),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.8804502671006565),
		Fy:    float64(0.49398530897300036),
		Fdy:   float32(-libc.Float64FromFloat64(0.8894893527030945)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.6678552546757298),
		Fy:    float64(2.302119780711071),
		Fdy:   float32(-libc.Float64FromFloat64(0.08916228264570236)),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(72),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.2709030829738149),
		Fy:    float64(1.845127398676589),
		Fdy:   float32(-libc.Float64FromFloat64(0.9070960283279419)),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(73),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.039291951227476314),
		Fy:    float64(1.6100983952478978),
		Fdy:   float32(-libc.Float64FromFloat64(0.8212206959724426)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(74),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.5783313839080494),
		Fy:    float64(0.9541144916925824),
		Fdy:   float32(-libc.Float64FromFloat64(0.49562546610832214)),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(75),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.8550888492150711),
		Fy:    float64(0.5450741471324799),
		Fdy:   float32(-libc.Float64FromFloat64(0.43663010001182556)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(76),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.1680688934754805),
		Fy:    float64(1.7396666950051183),
		Fdy:   float32(-libc.Float64FromFloat64(0.042779482901096344)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(77),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.9580241162600612),
		Fy:    float64(2.8508251232025135),
		Fdy:   float32(-libc.Float64FromFloat64(0.9732877016067505)),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(78),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.2687083954988399),
		Fy:    float64(1.8428481854794554),
		Fdy:   float32(0.7137790322303772),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(79),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.712225905636645),
		Fy:    float64(0.7781321670027921),
		Fdy:   float32(0.08036109805107117),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(80),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.8804502671006565),
		Fy:    float64(0.4939853089730004),
		Fdy:   float32(0.11051065474748611),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(81),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.6678552546757298),
		Fy:    float64(2.3021197807110716),
		Fdy:   float32(0.910837709903717),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(82),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.2709030829738149),
		Fy:    float64(1.8451273986765893),
		Fdy:   float32(0.09290396422147751),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(83),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.039291951227476314),
		Fy:    float64(1.610098395247898),
		Fdy:   float32(0.17877931892871857),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(84),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.5783313839080494),
		Fy:    float64(0.9541144916925826),
		Fdy:   float32(0.5043745040893555),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(85),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.8550888492150711),
		Fy:    float64(0.54507414713248),
		Fdy:   float32(0.5633699297904968),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(86),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.1680688934754805),
		Fy:    float64(1.7396666950051185),
		Fdy:   float32(0.9572204947471619),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(87),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(0.9580241162600612),
		Fy:    float64(2.850825123202514),
		Fdy:   float32(0.026712287217378616),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(88),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.2687083954988399),
		Fy:    float64(1.8428481854794552),
		Fdy:   float32(-libc.Float64FromFloat64(0.2862209379673004)),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(89),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.712225905636645),
		Fy:    float64(0.778132167002792),
		Fdy:   float32(-libc.Float64FromFloat64(0.9196389317512512)),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(90),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.8804502671006565),
		Fy:    float64(0.49398530897300036),
		Fdy:   float32(-libc.Float64FromFloat64(0.8894893527030945)),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(91),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.6678552546757298),
		Fy:    float64(2.302119780711071),
		Fdy:   float32(-libc.Float64FromFloat64(0.08916228264570236)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(92),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.2709030829738149),
		Fy:    float64(1.845127398676589),
		Fdy:   float32(-libc.Float64FromFloat64(0.9070960283279419)),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(93),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.039291951227476314),
		Fy:    float64(1.6100983952478978),
		Fdy:   float32(-libc.Float64FromFloat64(0.8212206959724426)),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(94),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.5783313839080494),
		Fy:    float64(0.9541144916925824),
		Fdy:   float32(-libc.Float64FromFloat64(0.49562546610832214)),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(95),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.8550888492150711),
		Fy:    float64(0.5450741471324799),
		Fdy:   float32(-libc.Float64FromFloat64(0.43663010001182556)),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(96),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.1680688934754805),
		Fy:    float64(1.7396666950051183),
		Fdy:   float32(-libc.Float64FromFloat64(0.042779482901096344)),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(97),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(0.9580241162600612),
		Fy:    float64(2.8508251232025135),
		Fdy:   float32(-libc.Float64FromFloat64(0.9732877016067505)),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	330: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	331: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(101),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	332: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(102),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	333: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	334: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	335: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(65536.00000000001),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	336: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(131071.99999999999),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	337: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	338: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	339: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(109),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	340: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(110),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	341: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	342: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	343: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(113),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	344: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	345: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	346: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	347: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	348: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	349: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	350: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	351: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(121),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	352: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	353: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(123),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	354: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	355: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	356: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(126),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	357: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(127),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	358: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	359: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	360: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(130),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	361: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(131),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	362: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(132),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	363: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(133),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(0.9999999999999982),
		Fy:    float64(5.960464477539063e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.6666666865348816)),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.9999999999999982),
		Fy:    float64(5.960464477539064e-08),
		Fdy:   float32(0.3333333432674408),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	366: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	367: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	368: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	369: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(139),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	370: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	371: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	372: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	373: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	374: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(144),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	375: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	376: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	377: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	378: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	379: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	380: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(150),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	381: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(151),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	382: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(152),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	383: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(153),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	384: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(154),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	385: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(155),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	386: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(156),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	387: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(157),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	388: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(158),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	389: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(159),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	390: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(160),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	391: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(161),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	392: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(162),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	393: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(163),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	394: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(164),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	395: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(165),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	396: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(166),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	397: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(167),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	398: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(168),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999972),
		Fy:    float64(7.450580596923831e-08),
		Fdy:   float32(0.6979166865348816),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(169),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(0.9999999999999982),
		Fy:    float64(5.960464477539064e-08),
		Fdy:   float32(0.3333333432674408),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(170),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	401: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(171),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	402: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(172),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	403: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(173),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	404: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(174),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	405: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(175),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	406: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(176),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	407: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(177),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	408: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(178),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	409: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(179),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	410: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(180),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0000000000000004),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	411: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(181),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	412: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(182),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	413: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(183),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	414: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(184),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	415: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(185),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	416: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	417: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(187),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	418: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(188),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999972),
		Fy:    float64(7.45058059692383e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.3020833432674408)),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(189),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(0.9999999999999982),
		Fy:    float64(5.960464477539063e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.6666666865348816)),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	421: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(4.345239849338305),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	422: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	423: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	424: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(9.267056966972586),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	425: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.6619858980995045),
		Fy:    float64(0.8473310828433507),
		Fdy:   float32(-libc.Float64FromFloat64(0.41553276777267456)),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    float64(1.989530071088669),
		Fdy:   float32(0.4973946213722229),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.5617597462207241),
		Fy:    float64(0.9742849645674904),
		Fdy:   float32(-libc.Float64FromFloat64(0.4428897500038147)),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.7741522965913037),
		Fy:    float64(0.6854215158636222),
		Fdy:   float32(-libc.Float64FromFloat64(0.12589527666568756)),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    float64(2.316874138205964),
		Fdy:   float32(-libc.Float64FromFloat64(0.17284949123859406)),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0),
		Fy:    float64(1.5707963267948966),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(3.141592653589793),
		Fdy:   float32(-libc.Float64FromFloat64(0.27576595544815063)),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1),
		Fy:    float64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	433: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	434: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0000000000000002),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	435: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	436: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	437: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	438: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.5309227209592985),
		Fy:    float64(2.1304853799705463),
		Fdy:   float32(0.1391008496284485),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(0.4939556746399746),
		Fy:    float64(1.0541629875851946),
		Fdy:   float32(0.22054767608642578),
		Fe:    int32(FE_INEXACT),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:13:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:14:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:18:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:21:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(15840)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:22:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:24:3:
		if (*d_d)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:25:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:26:3:
		libc.Xfesetround(tls, (*d_d)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:27:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:28:3:
		y = libc.Xacos(tls, (*d_d)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:29:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:31:3:
		if !(checkexcept(tls, e, (*d_d)(unsafe.Pointer(p)).Fe, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:32:4:
			libc.Xprintf(tls, __ccgo_ts+91, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, estr(tls, (*d_d)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:34:4:
			libc.Xprintf(tls, __ccgo_ts+140, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:35:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:37:3:
		d = ulperr(tls, y, (*d_d)(unsafe.Pointer(p)).Fy, (*d_d)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:38:3:
		if !(checkulp(tls, d, (*d_d)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+149, libc.VaList(bp+8, (*d_d)(unsafe.Pointer(p)).Ffile, (*d_d)(unsafe.Pointer(p)).Fline, rstr(tls, (*d_d)(unsafe.Pointer(p)).Fr), (*d_d)(unsafe.Pointer(p)).Fx, (*d_d)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*d_d)(unsafe.Pointer(p)).Fdy), float64((*d_d)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:41:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/acos.c:44:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+206, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+216, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 260,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 268,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 276,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 286,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 296,
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
				v2 = __ccgo_ts + 305
			} else {
				v2 = __ccgo_ts + 23
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+307, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 305
		} else {
			v3 = __ccgo_ts + 23
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+312, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 23
	} else {
		v4 = __ccgo_ts + 317
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+319, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 322
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 325
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 328
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 331
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 334
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+337, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+345, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+350, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+393, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+442) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+450) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+462) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+474) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+486) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+495) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+23) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+495) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+501, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/acos.h\x00\x00src/math/ucb/acos.h\x00src/math/sanity/acos.h\x00src/math/special/acos.h\x00%s:%d: bad fp exception: %s acos(%a)=%a, want %s\x00 got %s\n\x00%s:%d: %s acos(%a) want %a got %a ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
