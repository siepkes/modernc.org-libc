// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/sinhl.exe.go src/math/sinhl.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
	*(*float64)(unsafe.Pointer(bp)) = float64(float64(ywant))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:132:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(y))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
		return libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:133:2:
	if v9 = y == ywant; v9 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(y))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v5 = *(*uint64)(unsafe.Pointer(bp))
		goto _6
	_6:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(ywant))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v7 = *(*uint64)(unsafe.Pointer(bp))
		goto _8
	_8:
	}
	return libc.BoolInt32(v9 && int32(v5>>libc.Int32FromInt32(63)) == int32(v7>>libc.Int32FromInt32(63)))
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:5:19:
var t = [1770]l_l{
	0: {
		Ffile: __ccgo_ts,
		Fline: int32(11),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1: {
		Ffile: __ccgo_ts,
		Fline: int32(12),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	2: {
		Ffile: __ccgo_ts,
		Fline: int32(13),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	3: {
		Ffile: __ccgo_ts,
		Fline: int32(14),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	4: {
		Ffile: __ccgo_ts,
		Fline: int32(15),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	5: {
		Ffile: __ccgo_ts,
		Fline: int32(16),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	6: {
		Ffile: __ccgo_ts,
		Fline: int32(17),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	7: {
		Ffile: __ccgo_ts,
		Fline: int32(18),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	8: {
		Ffile: __ccgo_ts,
		Fline: int32(19),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	9: {
		Ffile: __ccgo_ts,
		Fline: int32(20),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(1e-323)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
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
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
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
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
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
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
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
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1000),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1000),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1000),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1000),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(43),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1000)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1000)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(45),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1000)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(46),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1000)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8773838938338423),
		Fy:    libc.Float64FromFloat64(0.9943656656548412),
		Fdy:   float32(-libc.Float64FromFloat64(3.2051589081248796e-16)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8324843825057703),
		Fy:    libc.Float64FromFloat64(0.9320279933884077),
		Fdy:   float32(-libc.Float64FromFloat64(1.7071925229108782e-16)),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.7016782729723596),
		Fy:    libc.Float64FromFloat64(0.760691287360039),
		Fdy:   float32(-libc.Float64FromFloat64(8.149637504756605e-16)),
		Fe:    int32(FE_INEXACT),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0010088821295156762),
		Fy:    libc.Float64FromFloat64(0.0010088823006629793),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0018137018051671279),
		Fy:    libc.Float64FromFloat64(0.00181370279953361),
		Fdy:   float32(-libc.Float64FromFloat64(7.271534330205833e-16)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.001853460121723812),
		Fy:    libc.Float64FromFloat64(0.0018534611829270423),
		Fdy:   float32(-libc.Float64FromFloat64(2.25456030120098e-16)),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0005413395436377842),
		Fy:    libc.Float64FromFloat64(0.0005413395700775749),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.000644010616011505),
		Fy:    libc.Float64FromFloat64(0.0006440106605287046),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0004849726431229318),
		Fy:    libc.Float64FromFloat64(0.00048497266213373555),
		Fdy:   float32(-libc.Float64FromFloat64(3.462221760913535e-16)),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.000130957003224479),
		Fy:    libc.Float64FromFloat64(0.00013095700359879201),
		Fdy:   float32(-libc.Float64FromFloat64(5.359160516619013e-16)),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.790008586109273e-05),
		Fy:    libc.Float64FromFloat64(8.790008597428496e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.2074613481787662e-17)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.49423109336556e-05),
		Fy:    libc.Float64FromFloat64(3.4942310940766155e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.1273182670831165e-16)),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.808964435412296e-05),
		Fy:    libc.Float64FromFloat64(3.808964436333317e-05),
		Fdy:   float32(-libc.Float64FromFloat64(5.835585844505219e-17)),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.634168703182075e-05),
		Fy:    libc.Float64FromFloat64(4.6341687048407616e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.073682023280606e-16)),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.1527907906210542e-05),
		Fy:    libc.Float64FromFloat64(2.1527907907873396e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.8402592813421606e-16)),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.8561568991070158e-05),
		Fy:    libc.Float64FromFloat64(2.85615689949534e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.866581416608688e-06),
		Fy:    libc.Float64FromFloat64(7.866581416689822e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4257829556795484e-05),
		Fy:    libc.Float64FromFloat64(1.4257829557278552e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4305114746044961e-05),
		Fy:    libc.Float64FromFloat64(1.430511474653285e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.577131872205953e-06),
		Fy:    libc.Float64FromFloat64(8.577131872311117e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0244019617673438e-05),
		Fy:    libc.Float64FromFloat64(1.0244019617852604e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0336331505177564e-05),
		Fy:    libc.Float64FromFloat64(1.033633150536162e-05),
		Fdy:   float32(-libc.Float64FromFloat64(5.732326175487816e-16)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.072883605954973e-05),
		Fy:    libc.Float64FromFloat64(1.0728836059755557e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.2504868092522025e-06),
		Fy:    libc.Float64FromFloat64(4.250486809265e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9650357407275533e-06),
		Fy:    libc.Float64FromFloat64(1.965035740728818e-06),
		Fdy:   float32(-libc.Float64FromFloat64(5.236359797508453e-16)),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6559185937208932e-06),
		Fy:    libc.Float64FromFloat64(3.655918593729037e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.31388071487577524),
		Fy:    libc.Float64FromFloat64(0.3190601423468649),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.3492801703451198),
		Fy:    libc.Float64FromFloat64(0.35642545076695176),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4984485542247685),
		Fy:    libc.Float64FromFloat64(0.5193464813863253),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4324351603899862),
		Fy:    libc.Float64FromFloat64(0.4460393129574175),
		Fdy:   float32(-libc.Float64FromFloat64(4.2620408269900846e-16)),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4753364096476337),
		Fy:    libc.Float64FromFloat64(0.4934396782290129),
		Fdy:   float32(-libc.Float64FromFloat64(3.0147673421035903e-16)),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.49482569032781143),
		Fy:    libc.Float64FromFloat64(0.5152675693026115),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4569123232416933),
		Fy:    libc.Float64FromFloat64(0.4729772810932722),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.29060888703511e-07),
		Fy:    libc.Float64FromFloat64(7.290608887035755e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.454940828631357e-07),
		Fy:    libc.Float64FromFloat64(9.454940828632766e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.1349775502573795e-07),
		Fy:    libc.Float64FromFloat64(3.134977550257431e-07),
		Fdy:   float32(-libc.Float64FromFloat64(2.631542372892755e-16)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.2188960857847e-07),
		Fy:    libc.Float64FromFloat64(3.218896085784755e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.034194528030488e-07),
		Fy:    libc.Float64FromFloat64(2.0341945280305018e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3924624706373388e-07),
		Fy:    libc.Float64FromFloat64(1.392462470637343e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.634599731097405e-07),
		Fy:    libc.Float64FromFloat64(1.634599731097412e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6444469549543377e-07),
		Fy:    libc.Float64FromFloat64(1.644446954954345e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0148200695957795e-07),
		Fy:    libc.Float64FromFloat64(2.0148200695957933e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.090209449380533e-07),
		Fy:    libc.Float64FromFloat64(2.0902094493805483e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2632966393753456e-07),
		Fy:    libc.Float64FromFloat64(2.263296639375365e-07),
		Fdy:   float32(-libc.Float64FromFloat64(2.1309429893897074e-16)),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1367922463794224e-07),
		Fy:    libc.Float64FromFloat64(1.136792246379425e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.096235686398444e-08),
		Fy:    libc.Float64FromFloat64(7.09623568639845e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.587139859288168e-08),
		Fy:    libc.Float64FromFloat64(7.587139859288175e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.2982386657816426e-08),
		Fy:    libc.Float64FromFloat64(4.298238665781644e-08),
		Fdy:   float32(-libc.Float64FromFloat64(7.733934910067876e-16)),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.920255018613676e-08),
		Fy:    libc.Float64FromFloat64(4.920255018613677e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.415441372490345e-08),
		Fy:    libc.Float64FromFloat64(5.415441372490347e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.833607377363538e-08),
		Fy:    libc.Float64FromFloat64(5.833607377363541e-08),
		Fdy:   float32(-libc.Float64FromFloat64(2.263467117339155e-16)),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1777127990000445),
		Fy:    libc.Float64FromFloat64(0.17864969338496478),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1801662095945897),
		Fy:    libc.Float64FromFloat64(0.18114248782368714),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.08690601661948748),
		Fy:    libc.Float64FromFloat64(0.08701545314212093),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.08962038872074476),
		Fy:    libc.Float64FromFloat64(0.08974040595832858),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11108243253708547),
		Fy:    libc.Float64FromFloat64(0.11131102022506617),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11517898837011448),
		Fy:    libc.Float64FromFloat64(0.11543382191536608),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.032088473695677924),
		Fy:    libc.Float64FromFloat64(0.032093980736414514),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.05883193938204074),
		Fy:    libc.Float64FromFloat64(0.05886588341210276),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.016374829370172404),
		Fy:    libc.Float64FromFloat64(0.016375561157562796),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.023857137199588468),
		Fy:    libc.Float64FromFloat64(0.023859400363937848),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.02881106049823509),
		Fy:    libc.Float64FromFloat64(0.028815046564440304),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.010420247635331758),
		Fy:    libc.Float64FromFloat64(0.01042043621081421),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0039101774105499395),
		Fy:    libc.Float64FromFloat64(0.003910187374658919),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0060800908664091335),
		Fy:    libc.Float64FromFloat64(0.006080128327443236),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.007365637349662377),
		Fy:    libc.Float64FromFloat64(0.007365703950689174),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0020809232504162846),
		Fy:    libc.Float64FromFloat64(0.0020809247522333386),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0021393409785517158),
		Fy:    libc.Float64FromFloat64(0.0021393426104341935),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0022360318188961083),
		Fy:    libc.Float64FromFloat64(0.00223603368219616),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0034741896753177847),
		Fy:    libc.Float64FromFloat64(0.003474196664230034),
		Fdy:   float32(-libc.Float64FromFloat64(5.794338516347487e-16)),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.003711892646510005),
		Fy:    libc.Float64FromFloat64(0.0037119011703496438),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0033766456879772244),
		Fy:    libc.Float64FromFloat64(0.0033766521045847228),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0035003122271196245),
		Fy:    libc.Float64FromFloat64(0.003500319374869898),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1582852936378536),
		Fy:    libc.Float64FromFloat64(1.4352219926746248),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2263946376216301),
		Fy:    libc.Float64FromFloat64(1.5577843363317636),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9656974021850588),
		Fy:    libc.Float64FromFloat64(3.499916026812682),
		Fdy:   float32(-libc.Float64FromFloat64(1.8182732177071952e-16)),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9075105287855971),
		Fy:    libc.Float64FromFloat64(3.2939242996646003),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9921958606526364),
		Fy:    libc.Float64FromFloat64(9.939609761586036),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.4705864249662364),
		Fy:    libc.Float64FromFloat64(16.062247485775032),
		Fdy:   float32(-libc.Float64FromFloat64(3.062605645382145e-16)),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.891291187158971),
		Fy:    libc.Float64FromFloat64(491.8348419887997),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(28.033688490833303),
		Fy:    libc.Float64FromFloat64(7.479046326036691e+11),
		Fdy:   float32(-libc.Float64FromFloat64(1.7127085602820758e-17)),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(22.585931539468277),
		Fy:    libc.Float64FromFloat64(3.2204416932372036e+09),
		Fdy:   float32(-libc.Float64FromFloat64(8.200796177188397e-16)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(22.94762359167755),
		Fy:    libc.Float64FromFloat64(4.623770853190357e+09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(28.510928379221593),
		Fy:    libc.Float64FromFloat64(1.2053380689421506e+12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(50.57145016271699),
		Fy:    libc.Float64FromFloat64(4.590624306365716e+21),
		Fdy:   float32(-libc.Float64FromFloat64(7.616565789731584e-16)),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(52.22836866388262),
		Fy:    libc.Float64FromFloat64(2.4069236479281332e+22),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(60.06174039698168),
		Fy:    libc.Float64FromFloat64(6.073687348575258e+25),
		Fdy:   float32(-libc.Float64FromFloat64(7.205126947141096e-18)),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(62.079094496527055),
		Fy:    libc.Float64FromFloat64(4.56644452180213e+26),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(70.02812768919067),
		Fy:    libc.Float64FromFloat64(1.293598304775298e+30),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(188.63649973560388),
		Fy:    libc.Float64FromFloat64(4.195279742863512e+81),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(250.1495613987814),
		Fy:    libc.Float64FromFloat64(2.1754252765260477e+108),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(288.2754631690901),
		Fy:    libc.Float64FromFloat64(7.859825909430946e+124),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(301.663664262519),
		Fy:    libc.Float64FromFloat64(5.126663166814058e+130),
		Fdy:   float32(-libc.Float64FromFloat64(1.1540876845899046e-16)),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(331.52131001037077),
		Fy:    libc.Float64FromFloat64(4.751662380502593e+143),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(339.1490441599899),
		Fy:    libc.Float64FromFloat64(9.761766738179631e+146),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(412.83872756953286),
		Fy:    libc.Float64FromFloat64(9.829949885546332e+178),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(455.12664958584196),
		Fy:    libc.Float64FromFloat64(2.2801451094916734e+197),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(510.87569028483415),
		Fy:    libc.Float64FromFloat64(3.710763136328248e+221),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(500.10095904114877),
		Fy:    libc.Float64FromFloat64(7.763488435777479e+216),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(535.5022871124662),
		Fy:    libc.Float64FromFloat64(1.839324449670326e+232),
		Fdy:   float32(-libc.Float64FromFloat64(1.4793597183407846e-16)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(314.1099726204526),
		Fy:    libc.Float64FromFloat64(1.3037605152680096e+136),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.2030902569538515)),
		Fy:    float64(-libc.Float64FromFloat64(12.283911423090622)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.9176400550478127)),
		Fy:    float64(-libc.Float64FromFloat64(1.0519574750489844)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.4357304533135248)),
		Fy:    float64(-libc.Float64FromFloat64(0.44964997575554405)),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.051317354887985),
		Fy:    libc.Float64FromFloat64(1568.9628465777837),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.051317354887985),
		Fy:    libc.Float64FromFloat64(1568.962846577784),
		Fdy:   float32(8.646399578537967e-07),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.051317354887985),
		Fy:    libc.Float64FromFloat64(1568.9628465777837),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.37088413683986104)),
		Fy:    float64(-libc.Float64FromFloat64(0.37944563966026995)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.002018281821645234),
		Fy:    libc.Float64FromFloat64(0.002018283191877733),
		Fdy:   float32(-libc.Float64FromFloat64(7.664326062695181e-07)),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.002018281821645234),
		Fy:    libc.Float64FromFloat64(0.0020182831918777334),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.002018281821645234),
		Fy:    libc.Float64FromFloat64(0.002018283191877733),
		Fdy:   float32(-libc.Float64FromFloat64(7.664326062695181e-07)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.03914805914962489)),
		Fy:    float64(-libc.Float64FromFloat64(0.03915805944289559)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(220.94359163179217)),
		Fy:    float64(-libc.Float64FromFloat64(4.503525870683281e+95)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(220.94359163179217)),
		Fy:    float64(-libc.Float64FromFloat64(4.5035258706832805e+95)),
		Fdy:   float32(8.2251489175178e-08),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(220.94359163179217)),
		Fy:    float64(-libc.Float64FromFloat64(4.5035258706832805e+95)),
		Fdy:   float32(8.2251489175178e-08),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(23.59024429461444)),
		Fy:    float64(-libc.Float64FromFloat64(8.791903816105587e+09)),
		Fdy:   float32(-libc.Float64FromFloat64(9.263247306989797e-07)),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(23.59024429461444)),
		Fy:    float64(-libc.Float64FromFloat64(8.791903816105585e+09)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(23.59024429461444)),
		Fy:    float64(-libc.Float64FromFloat64(8.791903816105585e+09)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.8614626113266568),
		Fy:    libc.Float64FromFloat64(0.9720379975734026),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.8614626113266568),
		Fy:    libc.Float64FromFloat64(0.9720379975734027),
		Fdy:   float32(4.2219281226607563e-07),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8614626113266568),
		Fy:    libc.Float64FromFloat64(0.9720379975734026),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0017344811551563634),
		Fy:    libc.Float64FromFloat64(0.0017344820248325371),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.8745093102756761),
		Fy:    libc.Float64FromFloat64(0.9903159297389821),
		Fdy:   float32(-libc.Float64FromFloat64(7.71222573803243e-07)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.8745093102756761),
		Fy:    libc.Float64FromFloat64(0.9903159297389822),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8745093102756761),
		Fy:    libc.Float64FromFloat64(0.9903159297389821),
		Fdy:   float32(-libc.Float64FromFloat64(7.71222573803243e-07)),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.011552864930981609),
		Fy:    libc.Float64FromFloat64(0.01155312192365101),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(104.25929145879475),
		Fy:    libc.Float64FromFloat64(9.510535522602076e+44),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.5602480352736698)),
		Fy:    float64(-libc.Float64FromFloat64(2.2749588952792497)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.18647907985343526),
		Fy:    libc.Float64FromFloat64(0.18756174508099566),
		Fdy:   float32(-libc.Float64FromFloat64(4.791321543962113e-07)),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.18647907985343526),
		Fy:    libc.Float64FromFloat64(0.18756174508099568),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.18647907985343526),
		Fy:    libc.Float64FromFloat64(0.18756174508099566),
		Fdy:   float32(-libc.Float64FromFloat64(4.791321543962113e-07)),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.1368735002461115),
		Fy:    libc.Float64FromFloat64(0.13730127354461383),
		Fdy:   float32(-libc.Float64FromFloat64(5.302902650328178e-07)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.1368735002461115),
		Fy:    libc.Float64FromFloat64(0.13730127354461386),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.1368735002461115),
		Fy:    libc.Float64FromFloat64(0.13730127354461383),
		Fdy:   float32(-libc.Float64FromFloat64(5.302902650328178e-07)),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.2079779672035478),
		Fy:    libc.Float64FromFloat64(0.20948055533857812),
		Fdy:   float32(-libc.Float64FromFloat64(2.806627037443832e-07)),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.2079779672035478),
		Fy:    libc.Float64FromFloat64(0.20948055533857815),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.2079779672035478),
		Fy:    libc.Float64FromFloat64(0.20948055533857812),
		Fdy:   float32(-libc.Float64FromFloat64(2.806627037443832e-07)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.014864080278925553)),
		Fy:    float64(-libc.Float64FromFloat64(0.014864627632141312)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995529651642)),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.733456292984288)),
		Fy:    float64(-libc.Float64FromFloat64(62336.917926662216)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.00304183182860822)),
		Fy:    float64(-libc.Float64FromFloat64(0.0030418365194906723)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.00304183182860822)),
		Fy:    float64(-libc.Float64FromFloat64(0.003041836519490672)),
		Fdy:   float32(2.463556469933792e-08),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.00304183182860822)),
		Fy:    float64(-libc.Float64FromFloat64(0.003041836519490672)),
		Fdy:   float32(2.463556469933792e-08),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4779809883234466),
		Fy:    libc.Float64FromFloat64(0.4963904180114629),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999940395355225)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.06109223935941238)),
		Fy:    float64(-libc.Float64FromFloat64(0.06113024848934278)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.19495767798746777),
		Fy:    libc.Float64FromFloat64(0.19619503517983408),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(17.60009175949559),
		Fy:    libc.Float64FromFloat64(2.2008616174284518e+07),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.039139750831924884)),
		Fy:    float64(-libc.Float64FromFloat64(0.0391497447591924)),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(675.3017485269734),
		Fy:    libc.Float64FromFloat64(9.523421563401886e+292),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.30537909998794),
		Fy:    libc.Float64FromFloat64(4.964129074715453),
		Fdy:   float32(-libc.Float64FromFloat64(4.928050429953146e-07)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.30537909998794),
		Fy:    libc.Float64FromFloat64(4.964129074715454),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.30537909998794),
		Fy:    libc.Float64FromFloat64(4.964129074715453),
		Fdy:   float32(-libc.Float64FromFloat64(4.928050429953146e-07)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(135.4691585640841)),
		Fy:    float64(-libc.Float64FromFloat64(3.4078309023107626e+58)),
		Fdy:   float32(-libc.Float64FromFloat64(2.646623045166052e-07)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(135.4691585640841)),
		Fy:    float64(-libc.Float64FromFloat64(3.407830902310762e+58)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(135.4691585640841)),
		Fy:    float64(-libc.Float64FromFloat64(3.407830902310762e+58)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(262.33731810340555)),
		Fy:    float64(-libc.Float64FromFloat64(4.27188599406152e+113)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(262.33731810340555)),
		Fy:    float64(-libc.Float64FromFloat64(4.2718859940615194e+113)),
		Fdy:   float32(4.981835388662148e-08),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(262.33731810340555)),
		Fy:    float64(-libc.Float64FromFloat64(4.2718859940615194e+113)),
		Fdy:   float32(4.981835388662148e-08),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.03444638413507548)),
		Fy:    float64(-libc.Float64FromFloat64(0.03445319661815227)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.03444638413507548)),
		Fy:    float64(-libc.Float64FromFloat64(0.034453196618152264)),
		Fdy:   float32(7.622465432177705e-07),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.03444638413507548)),
		Fy:    float64(-libc.Float64FromFloat64(0.034453196618152264)),
		Fdy:   float32(7.622465432177705e-07),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(486.95122179304525),
		Fy:    libc.Float64FromFloat64(1.5107708099422066e+211),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(59.43687140292632)),
		Fy:    float64(-libc.Float64FromFloat64(3.2514364961854605e+25)),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.73136571533244)),
		Fy:    float64(-libc.Float64FromFloat64(20.856054378729777)),
		Fdy:   float32(-libc.Float64FromFloat64(9.286142699238553e-07)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.73136571533244)),
		Fy:    float64(-libc.Float64FromFloat64(20.856054378729773)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(3.73136571533244)),
		Fy:    float64(-libc.Float64FromFloat64(20.856054378729773)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.4067726673770505),
		Fy:    libc.Float64FromFloat64(0.41808354580149343),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.4067726673770505),
		Fy:    libc.Float64FromFloat64(0.4180835458014935),
		Fdy:   float32(1.5870625702518737e-08),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4067726673770505),
		Fy:    libc.Float64FromFloat64(0.41808354580149343),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.19944543633673242)),
		Fy:    float64(-libc.Float64FromFloat64(0.20077034151484896)),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.05477687873473403)),
		Fy:    float64(-libc.Float64FromFloat64(0.05480427590761763)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.00379214158412666),
		Fy:    libc.Float64FromFloat64(0.0037921506728460187),
		Fdy:   float32(-libc.Float64FromFloat64(7.224122811066991e-08)),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.00379214158412666),
		Fy:    libc.Float64FromFloat64(0.003792150672846019),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00379214158412666),
		Fy:    libc.Float64FromFloat64(0.0037921506728460187),
		Fdy:   float32(-libc.Float64FromFloat64(7.224122811066991e-08)),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.2467513789191225)),
		Fy:    float64(-libc.Float64FromFloat64(4.6756108023305005)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0015491693925187935)),
		Fy:    float64(-libc.Float64FromFloat64(0.0015491700121674687)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0015491693925187935)),
		Fy:    float64(-libc.Float64FromFloat64(0.0015491700121674685)),
		Fdy:   float32(7.118841267583775e-07),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0015491693925187935)),
		Fy:    float64(-libc.Float64FromFloat64(0.0015491700121674685)),
		Fdy:   float32(7.118841267583775e-07),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.012889314459893554),
		Fy:    libc.Float64FromFloat64(0.012889671356004075),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(424.697147332777),
		Fy:    libc.Float64FromFloat64(1.3886652601451762e+184),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(424.697147332777),
		Fy:    libc.Float64FromFloat64(1.3886652601451764e+184),
		Fdy:   float32(2.2923701692434406e-07),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(424.697147332777),
		Fy:    libc.Float64FromFloat64(1.3886652601451762e+184),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.060236554711092336),
		Fy:    libc.Float64FromFloat64(0.06027298879984418),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(4.678228075774271)),
		Fy:    float64(-libc.Float64FromFloat64(53.78499289131485)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(22.098905811044816),
		Fy:    libc.Float64FromFloat64(1.978804339843703e+09),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.013679939328215169),
		Fy:    libc.Float64FromFloat64(0.013680366011202531),
		Fdy:   float32(-libc.Float64FromFloat64(4.602972580869391e-07)),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.013679939328215169),
		Fy:    libc.Float64FromFloat64(0.013680366011202533),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.013679939328215169),
		Fy:    libc.Float64FromFloat64(0.013680366011202531),
		Fdy:   float32(-libc.Float64FromFloat64(4.602972580869391e-07)),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(24.880002643289572)),
		Fy:    float64(-libc.Float64FromFloat64(3.1931392786850647e+10)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(24.880002643289572)),
		Fy:    float64(-libc.Float64FromFloat64(3.1931392786850643e+10)),
		Fdy:   float32(2.5453721264057094e-07),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(24.880002643289572)),
		Fy:    float64(-libc.Float64FromFloat64(3.1931392786850643e+10)),
		Fdy:   float32(2.5453721264057094e-07),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0013791945359087807),
		Fy:    libc.Float64FromFloat64(0.0013791949731543067),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0013791945359087807),
		Fy:    libc.Float64FromFloat64(0.001379194973154307),
		Fdy:   float32(6.076885483707883e-07),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0013791945359087807),
		Fy:    libc.Float64FromFloat64(0.0013791949731543067),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(115.14183632932038),
		Fy:    libc.Float64FromFloat64(5.063305809698415e+49),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(115.14183632932038),
		Fy:    libc.Float64FromFloat64(5.063305809698416e+49),
		Fdy:   float32(4.6504842998729146e-07),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(115.14183632932038),
		Fy:    libc.Float64FromFloat64(5.063305809698415e+49),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.29675638189035514)),
		Fy:    float64(-libc.Float64FromFloat64(0.3011312105057198)),
		Fdy:   float32(-libc.Float64FromFloat64(2.721487533108302e-07)),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.29675638189035514)),
		Fy:    float64(-libc.Float64FromFloat64(0.3011312105057197)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.29675638189035514)),
		Fy:    float64(-libc.Float64FromFloat64(0.3011312105057197)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(464.65293321018555)),
		Fy:    float64(-libc.Float64FromFloat64(3.127338545302208e+201)),
		Fdy:   float32(-libc.Float64FromFloat64(8.732300216252042e-07)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(464.65293321018555)),
		Fy:    float64(-libc.Float64FromFloat64(3.1273385453022076e+201)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(464.65293321018555)),
		Fy:    float64(-libc.Float64FromFloat64(3.1273385453022076e+201)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3410130052462377)),
		Fy:    float64(-libc.Float64FromFloat64(0.3476609350904727)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3410130052462377)),
		Fy:    float64(-libc.Float64FromFloat64(0.34766093509047263)),
		Fdy:   float32(3.5261146535958687e-07),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.3410130052462377)),
		Fy:    float64(-libc.Float64FromFloat64(0.34766093509047263)),
		Fdy:   float32(3.5261146535958687e-07),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1274356681803998),
		Fy:    libc.Float64FromFloat64(0.12778087104255825),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0027554690937075556),
		Fy:    libc.Float64FromFloat64(0.002755472580575878),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(621.470647173768),
		Fy:    libc.Float64FromFloat64(3.9832974360950064e+269),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.27730597762956)),
		Fy:    float64(-libc.Float64FromFloat64(0.28087374165898543)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.27730597762956)),
		Fy:    float64(-libc.Float64FromFloat64(0.2808737416589854)),
		Fdy:   float32(2.4272097221000877e-08),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.27730597762956)),
		Fy:    float64(-libc.Float64FromFloat64(0.2808737416589854)),
		Fdy:   float32(2.4272097221000877e-08),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(15.10351619030267)),
		Fy:    float64(-libc.Float64FromFloat64(1.8127743318161417e+06)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.88549694587863),
		Fy:    libc.Float64FromFloat64(3613.2023879414924),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.88549694587863),
		Fy:    libc.Float64FromFloat64(3613.202387941493),
		Fdy:   float32(6.43247688003612e-07),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.88549694587863),
		Fy:    libc.Float64FromFloat64(3613.2023879414924),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0035542384012813657),
		Fy:    libc.Float64FromFloat64(0.003554245884504384),
		Fdy:   float32(-libc.Float64FromFloat64(8.065375141086406e-07)),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0035542384012813657),
		Fy:    libc.Float64FromFloat64(0.0035542458845043844),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0035542384012813657),
		Fy:    libc.Float64FromFloat64(0.003554245884504384),
		Fdy:   float32(-libc.Float64FromFloat64(8.065375141086406e-07)),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0018095202025571468)),
		Fy:    float64(-libc.Float64FromFloat64(0.0018095211900617514)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.01517253010482324)),
		Fy:    float64(-libc.Float64FromFloat64(0.015173112245266311)),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(255.319215175061)),
		Fy:    float64(-libc.Float64FromFloat64(3.825571107013039e+110)),
		Fdy:   float32(-libc.Float64FromFloat64(4.0422060010314453e-07)),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(255.319215175061)),
		Fy:    float64(-libc.Float64FromFloat64(3.8255711070130385e+110)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(255.319215175061)),
		Fy:    float64(-libc.Float64FromFloat64(3.8255711070130385e+110)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7607030089685249)),
		Fy:    float64(-libc.Float64FromFloat64(0.8362211037581092)),
		Fdy:   float32(-libc.Float64FromFloat64(1.8761602404993027e-07)),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7607030089685249)),
		Fy:    float64(-libc.Float64FromFloat64(0.8362211037581091)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.7607030089685249)),
		Fy:    float64(-libc.Float64FromFloat64(0.8362211037581091)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(59.30557208882769),
		Fy:    libc.Float64FromFloat64(2.851364305482568e+25),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999916553497314)),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.542975583302663)),
		Fy:    float64(-libc.Float64FromFloat64(2565.2939841810594)),
		Fdy:   float32(-libc.Float64FromFloat64(9.413429324922618e-07)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.542975583302663)),
		Fy:    float64(-libc.Float64FromFloat64(2565.293984181059)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(8.542975583302663)),
		Fy:    float64(-libc.Float64FromFloat64(2565.293984181059)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.9493952180065532),
		Fy:    libc.Float64FromFloat64(25.942350203005244),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.9493952180065532),
		Fy:    libc.Float64FromFloat64(25.942350203005248),
		Fdy:   float32(1.013977666275423e-07),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.9493952180065532),
		Fy:    libc.Float64FromFloat64(25.942350203005244),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.22247355179804937),
		Fy:    libc.Float64FromFloat64(0.22431330094573432),
		Fdy:   float32(-libc.Float64FromFloat64(7.510515729336475e-07)),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.22247355179804937),
		Fy:    libc.Float64FromFloat64(0.22431330094573435),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.22247355179804937),
		Fy:    libc.Float64FromFloat64(0.22431330094573432),
		Fdy:   float32(-libc.Float64FromFloat64(7.510515729336475e-07)),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(47.586225270717385)),
		Fy:    float64(-libc.Float64FromFloat64(2.3195570257635813e+20)),
		Fdy:   float32(-libc.Float64FromFloat64(6.890074644161359e-08)),
		Fe:    int32(FE_INEXACT),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(47.586225270717385)),
		Fy:    float64(-libc.Float64FromFloat64(2.319557025763581e+20)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(47.586225270717385)),
		Fy:    float64(-libc.Float64FromFloat64(2.319557025763581e+20)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(7.7679935122764885),
		Fy:    libc.Float64FromFloat64(1181.8616606338474),
		Fdy:   float32(-libc.Float64FromFloat64(8.246815355050785e-07)),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(7.7679935122764885),
		Fy:    libc.Float64FromFloat64(1181.8616606338476),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.7679935122764885),
		Fy:    libc.Float64FromFloat64(1181.8616606338474),
		Fdy:   float32(-libc.Float64FromFloat64(8.246815355050785e-07)),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.03698864851274272)),
		Fy:    float64(-libc.Float64FromFloat64(0.03699708348869917)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.005533205275193556)),
		Fy:    float64(-libc.Float64FromFloat64(0.005533233509671457)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.002773075996949605)),
		Fy:    float64(-libc.Float64FromFloat64(0.0027730795510871563)),
		Fdy:   float32(-libc.Float64FromFloat64(9.411144930027149e-08)),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.002773075996949605)),
		Fy:    float64(-libc.Float64FromFloat64(0.002773079551087156)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.002773075996949605)),
		Fy:    float64(-libc.Float64FromFloat64(0.002773079551087156)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.9378812387323435),
		Fy:    libc.Float64FromFloat64(1.0815534054842582),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(157.49714733531803),
		Fy:    libc.Float64FromFloat64(1.2563539443368077e+68),
		Fdy:   float32(-libc.Float64FromFloat64(6.961474241506949e-07)),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(157.49714733531803),
		Fy:    libc.Float64FromFloat64(1.256353944336808e+68),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(157.49714733531803),
		Fy:    libc.Float64FromFloat64(1.2563539443368077e+68),
		Fdy:   float32(-libc.Float64FromFloat64(6.961474241506949e-07)),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(413.13308282886766),
		Fy:    libc.Float64FromFloat64(1.319435511557456e+179),
		Fdy:   float32(-libc.Float64FromFloat64(5.034079464394381e-08)),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(413.13308282886766),
		Fy:    libc.Float64FromFloat64(1.3194355115574563e+179),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(413.13308282886766),
		Fy:    libc.Float64FromFloat64(1.319435511557456e+179),
		Fdy:   float32(-libc.Float64FromFloat64(5.034079464394381e-08)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(103.94770340211716)),
		Fy:    float64(-libc.Float64FromFloat64(6.964404675265234e+44)),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.22860712515394274),
		Fy:    libc.Float64FromFloat64(0.23060354925242896),
		Fdy:   float32(-libc.Float64FromFloat64(8.432039635408728e-07)),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.22860712515394274),
		Fy:    libc.Float64FromFloat64(0.23060354925242899),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.22860712515394274),
		Fy:    libc.Float64FromFloat64(0.23060354925242896),
		Fdy:   float32(-libc.Float64FromFloat64(8.432039635408728e-07)),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.24577775802509466)),
		Fy:    float64(-libc.Float64FromFloat64(0.24825967988964937)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.24577775802509466)),
		Fy:    float64(-libc.Float64FromFloat64(0.24825967988964934)),
		Fdy:   float32(8.58448345297802e-07),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.24577775802509466)),
		Fy:    float64(-libc.Float64FromFloat64(0.24825967988964934)),
		Fdy:   float32(8.58448345297802e-07),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.18840562382704812),
		Fy:    libc.Float64FromFloat64(0.18952223277583394),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.18840562382704812),
		Fy:    libc.Float64FromFloat64(0.18952223277583397),
		Fdy:   float32(7.294478905350843e-07),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.18840562382704812),
		Fy:    libc.Float64FromFloat64(0.18952223277583394),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.004936457191744)),
		Fy:    float64(-libc.Float64FromFloat64(1.1828328951170173)),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0028567212225431783),
		Fy:    libc.Float64FromFloat64(0.002856725108093253),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0028567212225431783),
		Fy:    libc.Float64FromFloat64(0.0028567251080932536),
		Fdy:   float32(6.867040269753488e-07),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0028567212225431783),
		Fy:    libc.Float64FromFloat64(0.002856725108093253),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(11.55024858313248),
		Fy:    libc.Float64FromFloat64(51901.41861899404),
		Fdy:   float32(-libc.Float64FromFloat64(5.357226200430887e-07)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(11.55024858313248),
		Fy:    libc.Float64FromFloat64(51901.418618994045),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.55024858313248),
		Fy:    libc.Float64FromFloat64(51901.41861899404),
		Fdy:   float32(-libc.Float64FromFloat64(5.357226200430887e-07)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(250.28480828033375)),
		Fy:    float64(-libc.Float64FromFloat64(2.4904690255593755e+108)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(250.28480828033375)),
		Fy:    float64(-libc.Float64FromFloat64(2.490469025559375e+108)),
		Fdy:   float32(6.674870292044943e-07),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(250.28480828033375)),
		Fy:    float64(-libc.Float64FromFloat64(2.490469025559375e+108)),
		Fdy:   float32(6.674870292044943e-07),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.2929052658608045)),
		Fy:    float64(-libc.Float64FromFloat64(1.6844419519434584)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.17554341599404),
		Fy:    libc.Float64FromFloat64(1776.492433470843),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(66.62879654424248),
		Fy:    libc.Float64FromFloat64(4.320049065081397e+28),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01650768545364076),
		Fy:    libc.Float64FromFloat64(0.01650843519802584),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(166.03218173476108),
		Fy:    libc.Float64FromFloat64(6.394849825376102e+71),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(166.03218173476108),
		Fy:    libc.Float64FromFloat64(6.394849825376103e+71),
		Fdy:   float32(6.621335160161834e-07),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(166.03218173476108),
		Fy:    libc.Float64FromFloat64(6.394849825376102e+71),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.3724264642324049),
		Fy:    libc.Float64FromFloat64(0.3810957180164537),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.3724264642324049),
		Fy:    libc.Float64FromFloat64(0.38109571801645375),
		Fdy:   float32(2.237734264554092e-07),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.3724264642324049),
		Fy:    libc.Float64FromFloat64(0.3810957180164537),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.03442289147296586),
		Fy:    libc.Float64FromFloat64(0.0344296900265209),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.03442289147296586),
		Fy:    libc.Float64FromFloat64(0.03442969002652091),
		Fdy:   float32(5.516654937309795e-07),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.03442289147296586),
		Fy:    libc.Float64FromFloat64(0.0344296900265209),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.712644191268193),
		Fy:    libc.Float64FromFloat64(3039.647955706023),
		Fdy:   float32(0.49999967217445374),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.029913371544980628),
		Fy:    libc.Float64FromFloat64(0.029917832874232234),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.005944101950558901)),
		Fy:    float64(-libc.Float64FromFloat64(0.005944136953800515)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.005944101950558901)),
		Fy:    float64(-libc.Float64FromFloat64(0.005944136953800514)),
		Fdy:   float32(1.649702028316824e-07),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.005944101950558901)),
		Fy:    float64(-libc.Float64FromFloat64(0.005944136953800514)),
		Fdy:   float32(1.649702028316824e-07),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0029085292854841606),
		Fy:    libc.Float64FromFloat64(0.002908533386290463),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.06223843878973473)),
		Fy:    float64(-libc.Float64FromFloat64(0.06227862795021972)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0017765144077638057),
		Fy:    libc.Float64FromFloat64(0.0017765153422115505),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(511.09668849365715)),
		Fy:    float64(-libc.Float64FromFloat64(4.628513514721266e+221)),
		Fdy:   float32(-libc.Float64FromFloat64(4.4101048501943296e-07)),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(511.09668849365715)),
		Fy:    float64(-libc.Float64FromFloat64(4.628513514721265e+221)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(511.09668849365715)),
		Fy:    float64(-libc.Float64FromFloat64(4.628513514721265e+221)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6706181985271324),
		Fy:    libc.Float64FromFloat64(2.563661029614377),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.010976032000823628)),
		Fy:    float64(-libc.Float64FromFloat64(0.010976252388577813)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.010976032000823628)),
		Fy:    float64(-libc.Float64FromFloat64(0.010976252388577811)),
		Fdy:   float32(2.0649143550599547e-07),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.010976032000823628)),
		Fy:    float64(-libc.Float64FromFloat64(0.010976252388577811)),
		Fdy:   float32(2.0649143550599547e-07),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(30.627415587900483),
		Fy:    libc.Float64FromFloat64(1.0006624177197205e+13),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.23450645509616658)),
		Fy:    float64(-libc.Float64FromFloat64(0.23666175267178508)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.7852085541636632)),
		Fy:    float64(-libc.Float64FromFloat64(0.868419818985109)),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(160.9657804256401),
		Fy:    libc.Float64FromFloat64(4.031997168415856e+69),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(420.06783273943296)),
		Fy:    float64(-libc.Float64FromFloat64(1.355538285282819e+182)),
		Fdy:   float32(-libc.Float64FromFloat64(2.9577890359178127e-07)),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(420.06783273943296)),
		Fy:    float64(-libc.Float64FromFloat64(1.3555382852828186e+182)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(420.06783273943296)),
		Fy:    float64(-libc.Float64FromFloat64(1.3555382852828186e+182)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.003480231319644417)),
		Fy:    float64(-libc.Float64FromFloat64(0.0034802383450814518)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0033558923535781365),
		Fy:    libc.Float64FromFloat64(0.003355898652599175),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0033558923535781365),
		Fy:    libc.Float64FromFloat64(0.0033558986525991756),
		Fdy:   float32(2.901444133840414e-07),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0033558923535781365),
		Fy:    libc.Float64FromFloat64(0.003355898652599175),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.029230340220151925),
		Fy:    libc.Float64FromFloat64(0.029234502860732314),
		Fdy:   float32(0.49999961256980896),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.293114748260235)),
		Fy:    float64(-libc.Float64FromFloat64(0.29733003593464535)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23945489769554595)),
		Fy:    float64(-libc.Float64FromFloat64(0.2417498038487262)),
		Fdy:   float32(-libc.Float64FromFloat64(3.531154391112068e-07)),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23945489769554595)),
		Fy:    float64(-libc.Float64FromFloat64(0.24174980384872619)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.23945489769554595)),
		Fy:    float64(-libc.Float64FromFloat64(0.24174980384872619)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(248.36128540009028),
		Fy:    libc.Float64FromFloat64(3.638360867738952e+107),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999934434890747)),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(15.690499739263622)),
		Fy:    float64(-libc.Float64FromFloat64(3.26037428696075e+06)),
		Fdy:   float32(-libc.Float64FromFloat64(2.858526499949221e-07)),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(15.690499739263622)),
		Fy:    float64(-libc.Float64FromFloat64(3.2603742869607494e+06)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(15.690499739263622)),
		Fy:    float64(-libc.Float64FromFloat64(3.2603742869607494e+06)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.3373800077123154),
		Fy:    libc.Float64FromFloat64(0.34381692797926316),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.3373800077123154),
		Fy:    libc.Float64FromFloat64(0.3438169279792632),
		Fdy:   float32(1.976503938294627e-07),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.3373800077123154),
		Fy:    libc.Float64FromFloat64(0.34381692797926316),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.834420040756171),
		Fy:    libc.Float64FromFloat64(464.64408421602906),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.834420040756171),
		Fy:    libc.Float64FromFloat64(464.6440842160291),
		Fdy:   float32(8.846614605317882e-07),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.834420040756171),
		Fy:    libc.Float64FromFloat64(464.64408421602906),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.13747330226921378),
		Fy:    libc.Float64FromFloat64(0.13790672753096292),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.13747330226921378),
		Fy:    libc.Float64FromFloat64(0.13790672753096295),
		Fdy:   float32(6.953675324439246e-07),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.13747330226921378),
		Fy:    libc.Float64FromFloat64(0.13790672753096292),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(10.828905131474208)),
		Fy:    float64(-libc.Float64FromFloat64(25229.215583397996)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.03562314069779365),
		Fy:    libc.Float64FromFloat64(0.03563067551852974),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01670949296341074)),
		Fy:    float64(-libc.Float64FromFloat64(0.016710270542598177)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01670949296341074)),
		Fy:    float64(-libc.Float64FromFloat64(0.016710270542598174)),
		Fdy:   float32(4.542215208402922e-07),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.01670949296341074)),
		Fy:    float64(-libc.Float64FromFloat64(0.016710270542598174)),
		Fdy:   float32(4.542215208402922e-07),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(315.05195955891594),
		Fy:    libc.Float64FromFloat64(3.344240896870994e+136),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.5455044452576254),
		Fy:    libc.Float64FromFloat64(0.5729646068474747),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.5455044452576254),
		Fy:    libc.Float64FromFloat64(0.5729646068474749),
		Fdy:   float32(2.5811814907683583e-08),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.5455044452576254),
		Fy:    libc.Float64FromFloat64(0.5729646068474747),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.396448521825391),
		Fy:    libc.Float64FromFloat64(1.8966746259251075),
		Fdy:   float32(-libc.Float64FromFloat64(3.454772468103329e-08)),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.396448521825391),
		Fy:    libc.Float64FromFloat64(1.8966746259251077),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.396448521825391),
		Fy:    libc.Float64FromFloat64(1.8966746259251075),
		Fdy:   float32(-libc.Float64FromFloat64(3.454772468103329e-08)),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0033923505540281837),
		Fy:    libc.Float64FromFloat64(0.003392357060584196),
		Fdy:   float32(0.49999943375587463),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.23553709500026165),
		Fy:    libc.Float64FromFloat64(0.23772098769988698),
		Fdy:   float32(-libc.Float64FromFloat64(8.070480248534295e-07)),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.23553709500026165),
		Fy:    libc.Float64FromFloat64(0.237720987699887),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.23553709500026165),
		Fy:    libc.Float64FromFloat64(0.23772098769988698),
		Fdy:   float32(-libc.Float64FromFloat64(8.070480248534295e-07)),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.3620656817576133)),
		Fy:    float64(-libc.Float64FromFloat64(14.407029034646275)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.3620656817576133)),
		Fy:    float64(-libc.Float64FromFloat64(14.407029034646273)),
		Fdy:   float32(6.084349024604307e-07),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(3.3620656817576133)),
		Fy:    float64(-libc.Float64FromFloat64(14.407029034646273)),
		Fdy:   float32(6.084349024604307e-07),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(192.22679402125092)),
		Fy:    float64(-libc.Float64FromFloat64(1.520568262935169e+83)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.006340202773169752),
		Fy:    libc.Float64FromFloat64(0.006340245250681218),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.006340202773169752),
		Fy:    libc.Float64FromFloat64(0.006340245250681219),
		Fdy:   float32(4.652847849229147e-07),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.006340202773169752),
		Fy:    libc.Float64FromFloat64(0.006340245250681218),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(17.81298483795496),
		Fy:    libc.Float64FromFloat64(2.7230212197275124e+07),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(17.81298483795496),
		Fy:    libc.Float64FromFloat64(2.723021219727513e+07),
		Fdy:   float32(2.3618761701982294e-07),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(17.81298483795496),
		Fy:    libc.Float64FromFloat64(2.7230212197275124e+07),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(39.47528692852546),
		Fy:    libc.Float64FromFloat64(6.9641687965559864e+16),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(213.344905160842)),
		Fy:    float64(-libc.Float64FromFloat64(2.256758349163828e+92)),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004808612975988595),
		Fy:    libc.Float64FromFloat64(0.004808631507409649),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004808612975988595),
		Fy:    libc.Float64FromFloat64(0.00480863150740965),
		Fdy:   float32(5.680971639776544e-07),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004808612975988595),
		Fy:    libc.Float64FromFloat64(0.004808631507409649),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9152275039219852),
		Fy:    libc.Float64FromFloat64(1.048458912667124),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9152275039219852),
		Fy:    libc.Float64FromFloat64(1.0484589126671242),
		Fdy:   float32(6.305575084297743e-07),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9152275039219852),
		Fy:    libc.Float64FromFloat64(1.048458912667124),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(552.3011175756304),
		Fy:    libc.Float64FromFloat64(3.6332703643122223e+239),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.04230318738581301),
		Fy:    libc.Float64FromFloat64(0.04231580586112948),
		Fdy:   float32(0.49999961256980896),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(13.802624518826258),
		Fy:    libc.Float64FromFloat64(493598.31519403256),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.004883574753703039)),
		Fy:    float64(-libc.Float64FromFloat64(0.004883594165368015)),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.711829585337216),
		Fy:    libc.Float64FromFloat64(55.62326848596354),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(387.06297571199025)),
		Fy:    float64(-libc.Float64FromFloat64(6.284699328910981e+167)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(387.06297571199025)),
		Fy:    float64(-libc.Float64FromFloat64(6.28469932891098e+167)),
		Fdy:   float32(1.4118458580014703e-07),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(387.06297571199025)),
		Fy:    float64(-libc.Float64FromFloat64(6.28469932891098e+167)),
		Fdy:   float32(1.4118458580014703e-07),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4061912202200271)),
		Fy:    float64(-libc.Float64FromFloat64(0.4174533980317288)),
		Fdy:   float32(-libc.Float64FromFloat64(6.181154503792641e-07)),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4061912202200271)),
		Fy:    float64(-libc.Float64FromFloat64(0.41745339803172876)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.4061912202200271)),
		Fy:    float64(-libc.Float64FromFloat64(0.41745339803172876)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0015358939865735026),
		Fy:    libc.Float64FromFloat64(0.0015358945904282999),
		Fdy:   float32(-libc.Float64FromFloat64(4.409616281009221e-07)),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0015358939865735026),
		Fy:    libc.Float64FromFloat64(0.0015358945904283),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0015358939865735026),
		Fy:    libc.Float64FromFloat64(0.0015358945904282999),
		Fdy:   float32(-libc.Float64FromFloat64(4.409616281009221e-07)),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.001007414640438703)),
		Fy:    float64(-libc.Float64FromFloat64(0.001007414810840255)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995529651642)),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.06366795853300457),
		Fy:    libc.Float64FromFloat64(0.06371098141966683),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.06366795853300457),
		Fy:    libc.Float64FromFloat64(0.06371098141966684),
		Fdy:   float32(4.474165962164989e-07),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.06366795853300457),
		Fy:    libc.Float64FromFloat64(0.06371098141966683),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.37935676117930595)),
		Fy:    float64(-libc.Float64FromFloat64(0.38852142837208553)),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.8121988406458107)),
		Fy:    float64(-libc.Float64FromFloat64(8.293204104270405)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8448084576904168)),
		Fy:    float64(-libc.Float64FromFloat64(0.9489461532627864)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8448084576904168)),
		Fy:    float64(-libc.Float64FromFloat64(0.9489461532627863)),
		Fdy:   float32(7.386591391878028e-07),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.8448084576904168)),
		Fy:    float64(-libc.Float64FromFloat64(0.9489461532627863)),
		Fdy:   float32(7.386591391878028e-07),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(7.157519891545378),
		Fy:    libc.Float64FromFloat64(641.8612149107068),
		Fdy:   float32(-libc.Float64FromFloat64(5.005196612728469e-07)),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(7.157519891545378),
		Fy:    libc.Float64FromFloat64(641.861214910707),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.157519891545378),
		Fy:    libc.Float64FromFloat64(641.8612149107068),
		Fdy:   float32(-libc.Float64FromFloat64(5.005196612728469e-07)),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(348.35391058723087)),
		Fy:    float64(-libc.Float64FromFloat64(9.708477351180374e+150)),
		Fdy:   float32(0.4999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.895708049498697)),
		Fy:    float64(-libc.Float64FromFloat64(1.0203754256583368)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.895708049498697)),
		Fy:    float64(-libc.Float64FromFloat64(1.0203754256583366)),
		Fdy:   float32(8.62476611018792e-07),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.895708049498697)),
		Fy:    float64(-libc.Float64FromFloat64(1.0203754256583366)),
		Fdy:   float32(8.62476611018792e-07),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts,
		Fline: int32(448),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(38.90730405052602)),
		Fy:    float64(-libc.Float64FromFloat64(3.9463668942784344e+16)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts,
		Fline: int32(449),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.720327063112426),
		Fy:    libc.Float64FromFloat64(56.098016251215846),
		Fdy:   float32(-libc.Float64FromFloat64(2.335676185794e-07)),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts,
		Fline: int32(450),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.720327063112426),
		Fy:    libc.Float64FromFloat64(56.09801625121585),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	434: {
		Ffile: __ccgo_ts,
		Fline: int32(451),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.720327063112426),
		Fy:    libc.Float64FromFloat64(56.098016251215846),
		Fdy:   float32(-libc.Float64FromFloat64(2.335676185794e-07)),
		Fe:    int32(FE_INEXACT),
	},
	435: {
		Ffile: __ccgo_ts,
		Fline: int32(452),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(28.212110473529226)),
		Fy:    float64(-libc.Float64FromFloat64(8.939925627102595e+11)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts,
		Fline: int32(453),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(28.212110473529226)),
		Fy:    float64(-libc.Float64FromFloat64(8.939925627102594e+11)),
		Fdy:   float32(4.784643010680156e-07),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts,
		Fline: int32(454),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(28.212110473529226)),
		Fy:    float64(-libc.Float64FromFloat64(8.939925627102594e+11)),
		Fdy:   float32(4.784643010680156e-07),
		Fe:    int32(FE_INEXACT),
	},
	438: {
		Ffile: __ccgo_ts,
		Fline: int32(455),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0019028001962759313)),
		Fy:    float64(-libc.Float64FromFloat64(0.0019028013445046127)),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts,
		Fline: int32(456),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.13193245423579536),
		Fy:    libc.Float64FromFloat64(0.13231552731803387),
		Fdy:   float32(-libc.Float64FromFloat64(1.098900241913725e-07)),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts,
		Fline: int32(457),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.13193245423579536),
		Fy:    libc.Float64FromFloat64(0.1323155273180339),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	441: {
		Ffile: __ccgo_ts,
		Fline: int32(458),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.13193245423579536),
		Fy:    libc.Float64FromFloat64(0.13231552731803387),
		Fdy:   float32(-libc.Float64FromFloat64(1.098900241913725e-07)),
		Fe:    int32(FE_INEXACT),
	},
	442: {
		Ffile: __ccgo_ts,
		Fline: int32(459),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(218.66006581766686)),
		Fy:    float64(-libc.Float64FromFloat64(4.590183014265743e+94)),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	443: {
		Ffile: __ccgo_ts,
		Fline: int32(460),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(9.962441358722884)),
		Fy:    float64(-libc.Float64FromFloat64(10607.262382616791)),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	444: {
		Ffile: __ccgo_ts,
		Fline: int32(461),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(9.962441358722884)),
		Fy:    float64(-libc.Float64FromFloat64(10607.26238261679)),
		Fdy:   float32(5.328684551386687e-07),
		Fe:    int32(FE_INEXACT),
	},
	445: {
		Ffile: __ccgo_ts,
		Fline: int32(462),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(9.962441358722884)),
		Fy:    float64(-libc.Float64FromFloat64(10607.26238261679)),
		Fdy:   float32(5.328684551386687e-07),
		Fe:    int32(FE_INEXACT),
	},
	446: {
		Ffile: __ccgo_ts,
		Fline: int32(463),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(447.044987560902),
		Fy:    libc.Float64FromFloat64(7.049223243246581e+193),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	447: {
		Ffile: __ccgo_ts,
		Fline: int32(464),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(447.044987560902),
		Fy:    libc.Float64FromFloat64(7.049223243246582e+193),
		Fdy:   float32(4.581007431170292e-07),
		Fe:    int32(FE_INEXACT),
	},
	448: {
		Ffile: __ccgo_ts,
		Fline: int32(465),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(447.044987560902),
		Fy:    libc.Float64FromFloat64(7.049223243246581e+193),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	449: {
		Ffile: __ccgo_ts,
		Fline: int32(466),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(138.32947850775707)),
		Fy:    float64(-libc.Float64FromFloat64(5.9524972688123865e+59)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	450: {
		Ffile: __ccgo_ts,
		Fline: int32(467),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0019338548853052191)),
		Fy:    float64(-libc.Float64FromFloat64(0.001933856090675492)),
		Fdy:   float32(-libc.Float64FromFloat64(8.104841953127107e-08)),
		Fe:    int32(FE_INEXACT),
	},
	451: {
		Ffile: __ccgo_ts,
		Fline: int32(468),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0019338548853052191)),
		Fy:    float64(-libc.Float64FromFloat64(0.0019338560906754917)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	452: {
		Ffile: __ccgo_ts,
		Fline: int32(469),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0019338548853052191)),
		Fy:    float64(-libc.Float64FromFloat64(0.0019338560906754917)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	453: {
		Ffile: __ccgo_ts,
		Fline: int32(470),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.3362957031640077)),
		Fy:    float64(-libc.Float64FromFloat64(1.77105269149625)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	454: {
		Ffile: __ccgo_ts,
		Fline: int32(471),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.18783154047408937),
		Fy:    libc.Float64FromFloat64(0.1889379614196867),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	455: {
		Ffile: __ccgo_ts,
		Fline: int32(472),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.18783154047408937),
		Fy:    libc.Float64FromFloat64(0.18893796141968672),
		Fdy:   float32(8.833450806378096e-07),
		Fe:    int32(FE_INEXACT),
	},
	456: {
		Ffile: __ccgo_ts,
		Fline: int32(473),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.18783154047408937),
		Fy:    libc.Float64FromFloat64(0.1889379614196867),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	457: {
		Ffile: __ccgo_ts,
		Fline: int32(474),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0014842081109128645),
		Fy:    libc.Float64FromFloat64(0.0014842086558334306),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	458: {
		Ffile: __ccgo_ts,
		Fline: int32(475),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.03279343867537649)),
		Fy:    float64(-libc.Float64FromFloat64(0.032799316721337866)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	459: {
		Ffile: __ccgo_ts,
		Fline: int32(476),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0020467187984038442),
		Fy:    libc.Float64FromFloat64(0.0020467202273747147),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts,
		Fline: int32(477),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0034639227176956407),
		Fy:    libc.Float64FromFloat64(0.003463929644829697),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts,
		Fline: int32(478),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0034639227176956407),
		Fy:    libc.Float64FromFloat64(0.0034639296448296976),
		Fdy:   float32(2.2082613782004046e-07),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts,
		Fline: int32(479),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0034639227176956407),
		Fy:    libc.Float64FromFloat64(0.003463929644829697),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts,
		Fline: int32(480),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.004725058086661701)),
		Fy:    float64(-libc.Float64FromFloat64(0.004725075668759434)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts,
		Fline: int32(481),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(60.75478965133991),
		Fy:    libc.Float64FromFloat64(1.214618520912757e+26),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	465: {
		Ffile: __ccgo_ts,
		Fline: int32(482),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(60.75478965133991),
		Fy:    libc.Float64FromFloat64(1.2146185209127572e+26),
		Fdy:   float32(7.788086122673121e-07),
		Fe:    int32(FE_INEXACT),
	},
	466: {
		Ffile: __ccgo_ts,
		Fline: int32(483),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(60.75478965133991),
		Fy:    libc.Float64FromFloat64(1.214618520912757e+26),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts,
		Fline: int32(484),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.30526285604446757)),
		Fy:    float64(-libc.Float64FromFloat64(0.3100260022196166)),
		Fdy:   float32(-libc.Float64FromFloat64(3.797107126501942e-07)),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts,
		Fline: int32(485),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.30526285604446757)),
		Fy:    float64(-libc.Float64FromFloat64(0.31002600221961657)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts,
		Fline: int32(486),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.30526285604446757)),
		Fy:    float64(-libc.Float64FromFloat64(0.31002600221961657)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts,
		Fline: int32(487),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.269166117863374)),
		Fy:    float64(-libc.Float64FromFloat64(39183.81026951937)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts,
		Fline: int32(488),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.016202882136404338)),
		Fy:    float64(-libc.Float64FromFloat64(0.016203591111971987)),
		Fdy:   float32(0.4999992549419403),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts,
		Fline: int32(489),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.01017527276925483)),
		Fy:    float64(-libc.Float64FromFloat64(0.010175448354969061)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts,
		Fline: int32(490),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.00778127473010058)),
		Fy:    float64(-libc.Float64FromFloat64(0.007781353254081979)),
		Fdy:   float32(-libc.Float64FromFloat64(6.888102461743983e-07)),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts,
		Fline: int32(491),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.00778127473010058)),
		Fy:    float64(-libc.Float64FromFloat64(0.007781353254081978)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	475: {
		Ffile: __ccgo_ts,
		Fline: int32(492),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.00778127473010058)),
		Fy:    float64(-libc.Float64FromFloat64(0.007781353254081978)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	476: {
		Ffile: __ccgo_ts,
		Fline: int32(493),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(57.599806784238545),
		Fy:    libc.Float64FromFloat64(5.179027885777793e+24),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	477: {
		Ffile: __ccgo_ts,
		Fline: int32(494),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(57.599806784238545),
		Fy:    libc.Float64FromFloat64(5.179027885777795e+24),
		Fdy:   float32(4.758691716233443e-07),
		Fe:    int32(FE_INEXACT),
	},
	478: {
		Ffile: __ccgo_ts,
		Fline: int32(495),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(57.599806784238545),
		Fy:    libc.Float64FromFloat64(5.179027885777793e+24),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	479: {
		Ffile: __ccgo_ts,
		Fline: int32(496),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004611413364813623)),
		Fy:    float64(-libc.Float64FromFloat64(0.004611429708550924)),
		Fdy:   float32(-libc.Float64FromFloat64(6.62152046970732e-07)),
		Fe:    int32(FE_INEXACT),
	},
	480: {
		Ffile: __ccgo_ts,
		Fline: int32(497),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004611413364813623)),
		Fy:    float64(-libc.Float64FromFloat64(0.004611429708550923)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	481: {
		Ffile: __ccgo_ts,
		Fline: int32(498),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.004611413364813623)),
		Fy:    float64(-libc.Float64FromFloat64(0.004611429708550923)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	482: {
		Ffile: __ccgo_ts,
		Fline: int32(499),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(36.64472829643744),
		Fy:    libc.Float64FromFloat64(4.107459541478394e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	483: {
		Ffile: __ccgo_ts,
		Fline: int32(500),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(36.64472829643744),
		Fy:    libc.Float64FromFloat64(4.1074595414783945e+15),
		Fdy:   float32(3.5660193020703446e-07),
		Fe:    int32(FE_INEXACT),
	},
	484: {
		Ffile: __ccgo_ts,
		Fline: int32(501),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(36.64472829643744),
		Fy:    libc.Float64FromFloat64(4.107459541478394e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	485: {
		Ffile: __ccgo_ts,
		Fline: int32(502),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(18.261063066935396)),
		Fy:    float64(-libc.Float64FromFloat64(4.26234822704407e+07)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	486: {
		Ffile: __ccgo_ts,
		Fline: int32(503),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.9055288185608354),
		Fy:    libc.Float64FromFloat64(498.88750703715533),
		Fdy:   float32(-libc.Float64FromFloat64(9.511180110166606e-07)),
		Fe:    int32(FE_INEXACT),
	},
	487: {
		Ffile: __ccgo_ts,
		Fline: int32(504),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.9055288185608354),
		Fy:    libc.Float64FromFloat64(498.8875070371554),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	488: {
		Ffile: __ccgo_ts,
		Fline: int32(505),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.9055288185608354),
		Fy:    libc.Float64FromFloat64(498.88750703715533),
		Fdy:   float32(-libc.Float64FromFloat64(9.511180110166606e-07)),
		Fe:    int32(FE_INEXACT),
	},
	489: {
		Ffile: __ccgo_ts,
		Fline: int32(506),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(19.651797993248483),
		Fy:    libc.Float64FromFloat64(1.712527019322202e+08),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	490: {
		Ffile: __ccgo_ts,
		Fline: int32(507),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(554.3233909239034)),
		Fy:    float64(-libc.Float64FromFloat64(2.7451107624657635e+240)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	491: {
		Ffile: __ccgo_ts,
		Fline: int32(508),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(554.3233909239034)),
		Fy:    float64(-libc.Float64FromFloat64(2.745110762465763e+240)),
		Fdy:   float32(3.990522046137812e-08),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts,
		Fline: int32(509),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(554.3233909239034)),
		Fy:    float64(-libc.Float64FromFloat64(2.745110762465763e+240)),
		Fdy:   float32(3.990522046137812e-08),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts,
		Fline: int32(510),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0606374741882873)),
		Fy:    float64(-libc.Float64FromFloat64(0.06067464070842888)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts,
		Fline: int32(511),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.00468778171134664)),
		Fy:    float64(-libc.Float64FromFloat64(0.00468779888059836)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts,
		Fline: int32(512),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(20.68176012768859),
		Fy:    libc.Float64FromFloat64(4.7967192909220624e+08),
		Fdy:   float32(-libc.Float64FromFloat64(5.518028842743661e-07)),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts,
		Fline: int32(513),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(20.68176012768859),
		Fy:    libc.Float64FromFloat64(4.796719290922063e+08),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts,
		Fline: int32(514),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(20.68176012768859),
		Fy:    libc.Float64FromFloat64(4.7967192909220624e+08),
		Fdy:   float32(-libc.Float64FromFloat64(5.518028842743661e-07)),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts,
		Fline: int32(515),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.006880808066903397),
		Fy:    libc.Float64FromFloat64(0.006880862362937524),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts,
		Fline: int32(516),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(12.247971063935767),
		Fy:    libc.Float64FromFloat64(104278.85452233773),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts,
		Fline: int32(517),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.028902449167330622),
		Fy:    libc.Float64FromFloat64(0.028906473286442524),
		Fdy:   float32(-libc.Float64FromFloat64(6.629248900935636e-07)),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts,
		Fline: int32(518),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.028902449167330622),
		Fy:    libc.Float64FromFloat64(0.028906473286442528),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts,
		Fline: int32(519),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.028902449167330622),
		Fy:    libc.Float64FromFloat64(0.028906473286442524),
		Fdy:   float32(-libc.Float64FromFloat64(6.629248900935636e-07)),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts,
		Fline: int32(520),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(166.24433591717522)),
		Fy:    float64(-libc.Float64FromFloat64(7.906198989911125e+71)),
		Fdy:   float32(0.49999967217445374),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts,
		Fline: int32(521),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.8348633586187866)),
		Fy:    float64(-libc.Float64FromFloat64(8.484673426096462)),
		Fdy:   float32(-libc.Float64FromFloat64(4.619099627234391e-07)),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts,
		Fline: int32(522),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.8348633586187866)),
		Fy:    float64(-libc.Float64FromFloat64(8.48467342609646)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts,
		Fline: int32(523),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.8348633586187866)),
		Fy:    float64(-libc.Float64FromFloat64(8.48467342609646)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts,
		Fline: int32(524),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(105.91964108792094),
		Fy:    libc.Float64FromFloat64(5.003635371926832e+45),
		Fdy:   float32(-libc.Float64FromFloat64(5.71060603249407e-08)),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts,
		Fline: int32(525),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(105.91964108792094),
		Fy:    libc.Float64FromFloat64(5.003635371926833e+45),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts,
		Fline: int32(526),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(105.91964108792094),
		Fy:    libc.Float64FromFloat64(5.003635371926832e+45),
		Fdy:   float32(-libc.Float64FromFloat64(5.71060603249407e-08)),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts,
		Fline: int32(527),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(37.24889708827766)),
		Fy:    float64(-libc.Float64FromFloat64(7.515544779104878e+15)),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts,
		Fline: int32(528),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.49114296237883526),
		Fy:    libc.Float64FromFloat64(0.5111281886614168),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts,
		Fline: int32(529),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(12.107814184684846),
		Fy:    libc.Float64FromFloat64(90641.46007646355),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts,
		Fline: int32(530),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.710907465551164)),
		Fy:    float64(-libc.Float64FromFloat64(60947.02266610564)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts,
		Fline: int32(531),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(23.997527322008324)),
		Fy:    float64(-libc.Float64FromFloat64(1.3211851986441982e+10)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999994933605194)),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts,
		Fline: int32(532),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.9986125428187764),
		Fy:    libc.Float64FromFloat64(3.621644011665089),
		Fdy:   float32(-libc.Float64FromFloat64(3.9835313714320364e-07)),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts,
		Fline: int32(533),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.9986125428187764),
		Fy:    libc.Float64FromFloat64(3.6216440116650896),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts,
		Fline: int32(534),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9986125428187764),
		Fy:    libc.Float64FromFloat64(3.621644011665089),
		Fdy:   float32(-libc.Float64FromFloat64(3.9835313714320364e-07)),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts,
		Fline: int32(535),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.48297841386552676)),
		Fy:    float64(-libc.Float64FromFloat64(0.5019758875657032)),
		Fdy:   float32(-libc.Float64FromFloat64(2.5760132871255337e-07)),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts,
		Fline: int32(536),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.48297841386552676)),
		Fy:    float64(-libc.Float64FromFloat64(0.5019758875657031)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts,
		Fline: int32(537),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.48297841386552676)),
		Fy:    float64(-libc.Float64FromFloat64(0.5019758875657031)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts,
		Fline: int32(538),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.02211372765215255)),
		Fy:    float64(-libc.Float64FromFloat64(0.022115530027499004)),
		Fdy:   float32(-libc.Float64FromFloat64(4.974522767042799e-07)),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts,
		Fline: int32(539),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.02211372765215255)),
		Fy:    float64(-libc.Float64FromFloat64(0.022115530027499)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts,
		Fline: int32(540),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.02211372765215255)),
		Fy:    float64(-libc.Float64FromFloat64(0.022115530027499)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts,
		Fline: int32(541),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.00189614219625883)),
		Fy:    float64(-libc.Float64FromFloat64(0.001896143332476494)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts,
		Fline: int32(542),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(6.822978067454037)),
		Fy:    float64(-libc.Float64FromFloat64(459.35792633595787)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts,
		Fline: int32(543),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(6.822978067454037)),
		Fy:    float64(-libc.Float64FromFloat64(459.3579263359578)),
		Fdy:   float32(6.861841939098667e-07),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts,
		Fline: int32(544),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(6.822978067454037)),
		Fy:    float64(-libc.Float64FromFloat64(459.3579263359578)),
		Fdy:   float32(6.861841939098667e-07),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts,
		Fline: int32(545),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(73.70784042499079),
		Fy:    libc.Float64FromFloat64(5.12717775008987e+31),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts,
		Fline: int32(546),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.1324754243363918)),
		Fy:    float64(-libc.Float64FromFloat64(0.132863249323963)),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts,
		Fline: int32(547),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(17.838151931579034),
		Fy:    libc.Float64FromFloat64(2.792421386924552e+07),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts,
		Fline: int32(548),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(17.838151931579034),
		Fy:    libc.Float64FromFloat64(2.7924213869245525e+07),
		Fdy:   float32(2.8730491408168746e-07),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts,
		Fline: int32(549),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(17.838151931579034),
		Fy:    libc.Float64FromFloat64(2.792421386924552e+07),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts,
		Fline: int32(550),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(46.76738507888336)),
		Fy:    float64(-libc.Float64FromFloat64(1.0227918933271042e+20)),
		Fdy:   float32(-libc.Float64FromFloat64(1.6495391719217878e-07)),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts,
		Fline: int32(551),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(46.76738507888336)),
		Fy:    float64(-libc.Float64FromFloat64(1.0227918933271041e+20)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts,
		Fline: int32(552),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(46.76738507888336)),
		Fy:    float64(-libc.Float64FromFloat64(1.0227918933271041e+20)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts,
		Fline: int32(553),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00232004691963098),
		Fy:    libc.Float64FromFloat64(0.0023200490009524797),
		Fdy:   float32(0.49999964237213135),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts,
		Fline: int32(554),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(249.97527831177604)),
		Fy:    float64(-libc.Float64FromFloat64(1.8274857002656748e+108)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts,
		Fline: int32(555),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(249.97527831177604)),
		Fy:    float64(-libc.Float64FromFloat64(1.8274857002656746e+108)),
		Fdy:   float32(2.691785709885153e-07),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts,
		Fline: int32(556),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(249.97527831177604)),
		Fy:    float64(-libc.Float64FromFloat64(1.8274857002656746e+108)),
		Fdy:   float32(2.691785709885153e-07),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts,
		Fline: int32(557),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(343.3589647265288)),
		Fy:    float64(-libc.Float64FromFloat64(6.57466602995197e+148)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts,
		Fline: int32(558),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0025075753386883354),
		Fy:    libc.Float64FromFloat64(0.002507577966600566),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts,
		Fline: int32(559),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(37.25562745772068)),
		Fy:    float64(-libc.Float64FromFloat64(7.566297773652219e+15)),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999076128006)),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts,
		Fline: int32(560),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.290367311056003),
		Fy:    libc.Float64FromFloat64(13.407742489671243),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts,
		Fline: int32(561),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.4196860392358732),
		Fy:    libc.Float64FromFloat64(0.43211532764282995),
		Fdy:   float32(-libc.Float64FromFloat64(3.2030226293500164e-07)),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts,
		Fline: int32(562),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.4196860392358732),
		Fy:    libc.Float64FromFloat64(0.43211532764283),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts,
		Fline: int32(563),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.4196860392358732),
		Fy:    libc.Float64FromFloat64(0.43211532764282995),
		Fdy:   float32(-libc.Float64FromFloat64(3.2030226293500164e-07)),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts,
		Fline: int32(564),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23488176236988806)),
		Fy:    float64(-libc.Float64FromFloat64(0.23704744369021302)),
		Fdy:   float32(-libc.Float64FromFloat64(3.6041882367499056e-07)),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts,
		Fline: int32(565),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23488176236988806)),
		Fy:    float64(-libc.Float64FromFloat64(0.237047443690213)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts,
		Fline: int32(566),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.23488176236988806)),
		Fy:    float64(-libc.Float64FromFloat64(0.237047443690213)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts,
		Fline: int32(567),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.05527625622637421),
		Fy:    libc.Float64FromFloat64(0.0553044096335711),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts,
		Fline: int32(568),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.20277423863218966)),
		Fy:    float64(-libc.Float64FromFloat64(0.20416668955366502)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts,
		Fline: int32(569),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004026042230913948),
		Fy:    libc.Float64FromFloat64(0.004026053107286616),
		Fdy:   float32(-libc.Float64FromFloat64(5.158412363925891e-07)),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts,
		Fline: int32(570),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004026042230913948),
		Fy:    libc.Float64FromFloat64(0.004026053107286617),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts,
		Fline: int32(571),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004026042230913948),
		Fy:    libc.Float64FromFloat64(0.004026053107286616),
		Fdy:   float32(-libc.Float64FromFloat64(5.158412363925891e-07)),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts,
		Fline: int32(572),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.03503466370101309),
		Fy:    libc.Float64FromFloat64(0.03504183122676457),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts,
		Fline: int32(573),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.03503466370101309),
		Fy:    libc.Float64FromFloat64(0.035041831226764575),
		Fdy:   float32(4.33588752457581e-07),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts,
		Fline: int32(574),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.03503466370101309),
		Fy:    libc.Float64FromFloat64(0.03504183122676457),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts,
		Fline: int32(575),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.993876826830656)),
		Fy:    float64(-libc.Float64FromFloat64(80880.6302629405)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts,
		Fline: int32(576),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4947782241653227),
		Fy:    libc.Float64FromFloat64(2.1170246731257785),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999919533729553)),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts,
		Fline: int32(577),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.012728747770223658)),
		Fy:    float64(-libc.Float64FromFloat64(0.012729091493790583)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts,
		Fline: int32(578),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.11673856768590349),
		Fy:    libc.Float64FromFloat64(0.11700389853825394),
		Fdy:   float32(-libc.Float64FromFloat64(3.530482217684039e-07)),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts,
		Fline: int32(579),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.11673856768590349),
		Fy:    libc.Float64FromFloat64(0.11700389853825395),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts,
		Fline: int32(580),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.11673856768590349),
		Fy:    libc.Float64FromFloat64(0.11700389853825394),
		Fdy:   float32(-libc.Float64FromFloat64(3.530482217684039e-07)),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts,
		Fline: int32(581),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.11396626028873037)),
		Fy:    float64(-libc.Float64FromFloat64(0.11421312537643959)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts,
		Fline: int32(582),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.11396626028873037)),
		Fy:    float64(-libc.Float64FromFloat64(0.11421312537643957)),
		Fdy:   float32(6.057466634956654e-07),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts,
		Fline: int32(583),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.11396626028873037)),
		Fy:    float64(-libc.Float64FromFloat64(0.11421312537643957)),
		Fdy:   float32(6.057466634956654e-07),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts,
		Fline: int32(584),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(43.23970545174377)),
		Fy:    float64(-libc.Float64FromFloat64(3.0042459276869883e+18)),
		Fdy:   float32(-libc.Float64FromFloat64(8.006793450476835e-07)),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts,
		Fline: int32(585),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(43.23970545174377)),
		Fy:    float64(-libc.Float64FromFloat64(3.004245927686988e+18)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts,
		Fline: int32(586),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(43.23970545174377)),
		Fy:    float64(-libc.Float64FromFloat64(3.004245927686988e+18)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts,
		Fline: int32(587),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(507.5083473479991),
		Fy:    libc.Float64FromFloat64(1.2795132200304443e+220),
		Fdy:   float32(-libc.Float64FromFloat64(6.526752258650959e-07)),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts,
		Fline: int32(588),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(507.5083473479991),
		Fy:    libc.Float64FromFloat64(1.2795132200304445e+220),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts,
		Fline: int32(589),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(507.5083473479991),
		Fy:    libc.Float64FromFloat64(1.2795132200304443e+220),
		Fdy:   float32(-libc.Float64FromFloat64(6.526752258650959e-07)),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts,
		Fline: int32(590),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.02530200766357962)),
		Fy:    float64(-libc.Float64FromFloat64(0.02530470743875719)),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts,
		Fline: int32(591),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.02530200766357962)),
		Fy:    float64(-libc.Float64FromFloat64(0.025304707438757186)),
		Fdy:   float32(5.232120656728512e-07),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts,
		Fline: int32(592),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.02530200766357962)),
		Fy:    float64(-libc.Float64FromFloat64(0.025304707438757186)),
		Fdy:   float32(5.232120656728512e-07),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts,
		Fline: int32(593),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.14815976481990714),
		Fy:    libc.Float64FromFloat64(0.14870241036730367),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts,
		Fline: int32(594),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.14815976481990714),
		Fy:    libc.Float64FromFloat64(0.1487024103673037),
		Fdy:   float32(1.8995517336861667e-07),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts,
		Fline: int32(595),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.14815976481990714),
		Fy:    libc.Float64FromFloat64(0.14870241036730367),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts,
		Fline: int32(596),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.21277212341625848),
		Fy:    libc.Float64FromFloat64(0.21438119716036635),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts,
		Fline: int32(597),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.21277212341625848),
		Fy:    libc.Float64FromFloat64(0.21438119716036638),
		Fdy:   float32(7.829136166037642e-07),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts,
		Fline: int32(598),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.21277212341625848),
		Fy:    libc.Float64FromFloat64(0.21438119716036635),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts,
		Fline: int32(599),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(63.88103890072659)),
		Fy:    float64(-libc.Float64FromFloat64(2.767914668226584e+27)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts,
		Fline: int32(600),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(63.88103890072659)),
		Fy:    float64(-libc.Float64FromFloat64(2.7679146682265833e+27)),
		Fdy:   float32(4.043155570343515e-07),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts,
		Fline: int32(601),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(63.88103890072659)),
		Fy:    float64(-libc.Float64FromFloat64(2.7679146682265833e+27)),
		Fdy:   float32(4.043155570343515e-07),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts,
		Fline: int32(602),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(7.630541277251757)),
		Fy:    float64(-libc.Float64FromFloat64(1030.0821765305102)),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts,
		Fline: int32(603),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.296653635743203),
		Fy:    libc.Float64FromFloat64(5450.738847235007),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	587: {
		Ffile: __ccgo_ts,
		Fline: int32(604),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(329.58681474127246),
		Fy:    libc.Float64FromFloat64(6.866018267171989e+142),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts,
		Fline: int32(605),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.4428409200057175)),
		Fy:    float64(-libc.Float64FromFloat64(0.4574576228203432)),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	589: {
		Ffile: __ccgo_ts,
		Fline: int32(606),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.09395402513993366),
		Fy:    libc.Float64FromFloat64(0.09409231381107289),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	590: {
		Ffile: __ccgo_ts,
		Fline: int32(607),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.23908586978195723),
		Fy:    libc.Float64FromFloat64(0.24137016195738606),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts,
		Fline: int32(608),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.003693938589191232)),
		Fy:    float64(-libc.Float64FromFloat64(0.0036939469899412063)),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts,
		Fline: int32(609),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.17514276470102308),
		Fy:    libc.Float64FromFloat64(0.17603955608486418),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts,
		Fline: int32(610),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.17514276470102308),
		Fy:    libc.Float64FromFloat64(0.1760395560848642),
		Fdy:   float32(3.945540427707783e-08),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts,
		Fline: int32(611),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.17514276470102308),
		Fy:    libc.Float64FromFloat64(0.17603955608486418),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts,
		Fline: int32(612),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2234145693782823)),
		Fy:    float64(-libc.Float64FromFloat64(4.565292392290292)),
		Fdy:   float32(-libc.Float64FromFloat64(8.119993708533002e-07)),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts,
		Fline: int32(613),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2234145693782823)),
		Fy:    float64(-libc.Float64FromFloat64(4.565292392290291)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	597: {
		Ffile: __ccgo_ts,
		Fline: int32(614),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.2234145693782823)),
		Fy:    float64(-libc.Float64FromFloat64(4.565292392290291)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	598: {
		Ffile: __ccgo_ts,
		Fline: int32(615),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.042400675923642696),
		Fy:    libc.Float64FromFloat64(0.0424133818439857),
		Fdy:   float32(0.4999992549419403),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts,
		Fline: int32(616),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(27.054627970250813)),
		Fy:    float64(-libc.Float64FromFloat64(2.809607423809439e+11)),
		Fdy:   float32(-libc.Float64FromFloat64(6.184167773426452e-07)),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts,
		Fline: int32(617),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(27.054627970250813)),
		Fy:    float64(-libc.Float64FromFloat64(2.8096074238094385e+11)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	601: {
		Ffile: __ccgo_ts,
		Fline: int32(618),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(27.054627970250813)),
		Fy:    float64(-libc.Float64FromFloat64(2.8096074238094385e+11)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	602: {
		Ffile: __ccgo_ts,
		Fline: int32(619),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(539.4938216884754)),
		Fy:    float64(-libc.Float64FromFloat64(9.957717117978284e+233)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	603: {
		Ffile: __ccgo_ts,
		Fline: int32(620),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(539.4938216884754)),
		Fy:    float64(-libc.Float64FromFloat64(9.957717117978282e+233)),
		Fdy:   float32(8.970760632109887e-07),
		Fe:    int32(FE_INEXACT),
	},
	604: {
		Ffile: __ccgo_ts,
		Fline: int32(621),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(539.4938216884754)),
		Fy:    float64(-libc.Float64FromFloat64(9.957717117978282e+233)),
		Fdy:   float32(8.970760632109887e-07),
		Fe:    int32(FE_INEXACT),
	},
	605: {
		Ffile: __ccgo_ts,
		Fline: int32(622),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.2451174070820329)),
		Fy:    float64(-libc.Float64FromFloat64(0.24757933761045026)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	606: {
		Ffile: __ccgo_ts,
		Fline: int32(623),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.6099869955800075)),
		Fy:    float64(-libc.Float64FromFloat64(18.469259948604883)),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	607: {
		Ffile: __ccgo_ts,
		Fline: int32(624),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.001228968583421632)),
		Fy:    float64(-libc.Float64FromFloat64(0.0012289688927865942)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	608: {
		Ffile: __ccgo_ts,
		Fline: int32(625),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.107561179556031),
		Fy:    libc.Float64FromFloat64(11.160860130877317),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	609: {
		Ffile: __ccgo_ts,
		Fline: int32(626),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0829498784271337)),
		Fy:    float64(-libc.Float64FromFloat64(0.08304503645265181)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	610: {
		Ffile: __ccgo_ts,
		Fline: int32(627),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0829498784271337)),
		Fy:    float64(-libc.Float64FromFloat64(0.0830450364526518)),
		Fdy:   float32(4.716271178040188e-07),
		Fe:    int32(FE_INEXACT),
	},
	611: {
		Ffile: __ccgo_ts,
		Fline: int32(628),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0829498784271337)),
		Fy:    float64(-libc.Float64FromFloat64(0.0830450364526518)),
		Fdy:   float32(4.716271178040188e-07),
		Fe:    int32(FE_INEXACT),
	},
	612: {
		Ffile: __ccgo_ts,
		Fline: int32(629),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.2413114226922969)),
		Fy:    float64(-libc.Float64FromFloat64(0.2436602266860774)),
		Fdy:   float32(0.4999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	613: {
		Ffile: __ccgo_ts,
		Fline: int32(630),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.06961657891879666)),
		Fy:    float64(-libc.Float64FromFloat64(0.06967282496783433)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	614: {
		Ffile: __ccgo_ts,
		Fline: int32(631),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.06961657891879666)),
		Fy:    float64(-libc.Float64FromFloat64(0.06967282496783432)),
		Fdy:   float32(9.409273360461157e-08),
		Fe:    int32(FE_INEXACT),
	},
	615: {
		Ffile: __ccgo_ts,
		Fline: int32(632),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.06961657891879666)),
		Fy:    float64(-libc.Float64FromFloat64(0.06967282496783432)),
		Fdy:   float32(9.409273360461157e-08),
		Fe:    int32(FE_INEXACT),
	},
	616: {
		Ffile: __ccgo_ts,
		Fline: int32(633),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.003508627976279679),
		Fy:    libc.Float64FromFloat64(0.0035086351750941783),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999940395355225)),
		Fe:    int32(FE_INEXACT),
	},
	617: {
		Ffile: __ccgo_ts,
		Fline: int32(634),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.042564767775407894),
		Fy:    libc.Float64FromFloat64(0.042577621793188705),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	618: {
		Ffile: __ccgo_ts,
		Fline: int32(635),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.042564767775407894),
		Fy:    libc.Float64FromFloat64(0.04257762179318871),
		Fdy:   float32(3.186611081673618e-07),
		Fe:    int32(FE_INEXACT),
	},
	619: {
		Ffile: __ccgo_ts,
		Fline: int32(636),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.042564767775407894),
		Fy:    libc.Float64FromFloat64(0.042577621793188705),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	620: {
		Ffile: __ccgo_ts,
		Fline: int32(637),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(70.46416354974295),
		Fy:    libc.Float64FromFloat64(2.0006328977221363e+30),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	621: {
		Ffile: __ccgo_ts,
		Fline: int32(638),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(113.89612495224017)),
		Fy:    float64(-libc.Float64FromFloat64(1.4568961017693983e+49)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	622: {
		Ffile: __ccgo_ts,
		Fline: int32(639),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.30058156975192296),
		Fy:    libc.Float64FromFloat64(0.30512828223969213),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	623: {
		Ffile: __ccgo_ts,
		Fline: int32(640),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.30058156975192296),
		Fy:    libc.Float64FromFloat64(0.3051282822396922),
		Fdy:   float32(7.422396492984262e-07),
		Fe:    int32(FE_INEXACT),
	},
	624: {
		Ffile: __ccgo_ts,
		Fline: int32(641),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.30058156975192296),
		Fy:    libc.Float64FromFloat64(0.30512828223969213),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	625: {
		Ffile: __ccgo_ts,
		Fline: int32(642),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(151.46630769402964),
		Fy:    libc.Float64FromFloat64(3.0196153896522268e+65),
		Fdy:   float32(0.4999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	626: {
		Ffile: __ccgo_ts,
		Fline: int32(643),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.025966553520381897)),
		Fy:    float64(-libc.Float64FromFloat64(0.025969471661719392)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	627: {
		Ffile: __ccgo_ts,
		Fline: int32(644),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.909293944606481)),
		Fy:    float64(-libc.Float64FromFloat64(3.300068725075803)),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	628: {
		Ffile: __ccgo_ts,
		Fline: int32(645),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(117.81810723564494),
		Fy:    libc.Float64FromFloat64(7.357391141127855e+50),
		Fdy:   float32(-libc.Float64FromFloat64(3.2369419500355434e-07)),
		Fe:    int32(FE_INEXACT),
	},
	629: {
		Ffile: __ccgo_ts,
		Fline: int32(646),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(117.81810723564494),
		Fy:    libc.Float64FromFloat64(7.357391141127856e+50),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	630: {
		Ffile: __ccgo_ts,
		Fline: int32(647),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(117.81810723564494),
		Fy:    libc.Float64FromFloat64(7.357391141127855e+50),
		Fdy:   float32(-libc.Float64FromFloat64(3.2369419500355434e-07)),
		Fe:    int32(FE_INEXACT),
	},
	631: {
		Ffile: __ccgo_ts,
		Fline: int32(648),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(6.9449513876467615)),
		Fy:    float64(-libc.Float64FromFloat64(518.9477889495071)),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	632: {
		Ffile: __ccgo_ts,
		Fline: int32(649),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(11.02887830794932)),
		Fy:    float64(-libc.Float64FromFloat64(30814.206947037062)),
		Fdy:   float32(-libc.Float64FromFloat64(8.296756845993514e-07)),
		Fe:    int32(FE_INEXACT),
	},
	633: {
		Ffile: __ccgo_ts,
		Fline: int32(650),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(11.02887830794932)),
		Fy:    float64(-libc.Float64FromFloat64(30814.20694703706)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	634: {
		Ffile: __ccgo_ts,
		Fline: int32(651),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(11.02887830794932)),
		Fy:    float64(-libc.Float64FromFloat64(30814.20694703706)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	635: {
		Ffile: __ccgo_ts,
		Fline: int32(652),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.3889129468515815),
		Fy:    libc.Float64FromFloat64(109.48033117805711),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	636: {
		Ffile: __ccgo_ts,
		Fline: int32(653),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.011651258380550988),
		Fy:    libc.Float64FromFloat64(0.011651521995598959),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	637: {
		Ffile: __ccgo_ts,
		Fline: int32(654),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.011651258380550988),
		Fy:    libc.Float64FromFloat64(0.01165152199559896),
		Fdy:   float32(2.664144744812802e-07),
		Fe:    int32(FE_INEXACT),
	},
	638: {
		Ffile: __ccgo_ts,
		Fline: int32(655),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.011651258380550988),
		Fy:    libc.Float64FromFloat64(0.011651521995598959),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	639: {
		Ffile: __ccgo_ts,
		Fline: int32(656),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4451345950319313)),
		Fy:    float64(-libc.Float64FromFloat64(0.4599811070615632)),
		Fdy:   float32(-libc.Float64FromFloat64(9.457058354200854e-07)),
		Fe:    int32(FE_INEXACT),
	},
	640: {
		Ffile: __ccgo_ts,
		Fline: int32(657),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4451345950319313)),
		Fy:    float64(-libc.Float64FromFloat64(0.45998110706156314)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	641: {
		Ffile: __ccgo_ts,
		Fline: int32(658),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.4451345950319313)),
		Fy:    float64(-libc.Float64FromFloat64(0.45998110706156314)),
		Fdy:   float32(0.9999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	642: {
		Ffile: __ccgo_ts,
		Fline: int32(659),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.215164863488302),
		Fy:    libc.Float64FromFloat64(1.5370943234421397),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	643: {
		Ffile: __ccgo_ts,
		Fline: int32(660),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.215164863488302),
		Fy:    libc.Float64FromFloat64(1.53709432344214),
		Fdy:   float32(4.040669310256817e-08),
		Fe:    int32(FE_INEXACT),
	},
	644: {
		Ffile: __ccgo_ts,
		Fline: int32(661),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.215164863488302),
		Fy:    libc.Float64FromFloat64(1.5370943234421397),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	645: {
		Ffile: __ccgo_ts,
		Fline: int32(662),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(55.43837579068791),
		Fy:    libc.Float64FromFloat64(5.964179366194614e+23),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	646: {
		Ffile: __ccgo_ts,
		Fline: int32(663),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(55.43837579068791),
		Fy:    libc.Float64FromFloat64(5.9641793661946145e+23),
		Fdy:   float32(8.232533446062007e-07),
		Fe:    int32(FE_INEXACT),
	},
	647: {
		Ffile: __ccgo_ts,
		Fline: int32(664),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(55.43837579068791),
		Fy:    libc.Float64FromFloat64(5.964179366194614e+23),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	648: {
		Ffile: __ccgo_ts,
		Fline: int32(665),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(239.4822744973493)),
		Fy:    float64(-libc.Float64FromFloat64(5.067576717760302e+103)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	649: {
		Ffile: __ccgo_ts,
		Fline: int32(666),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.24338852623582716)),
		Fy:    float64(-libc.Float64FromFloat64(0.24579862754132714)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	650: {
		Ffile: __ccgo_ts,
		Fline: int32(667),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(21.255617313918563)),
		Fy:    float64(-libc.Float64FromFloat64(8.51466004789427e+08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999916553497314)),
		Fe:    int32(FE_INEXACT),
	},
	651: {
		Ffile: __ccgo_ts,
		Fline: int32(668),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0024174225308807775)),
		Fy:    float64(-libc.Float64FromFloat64(0.00241742488542349)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	652: {
		Ffile: __ccgo_ts,
		Fline: int32(669),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0024174225308807775)),
		Fy:    float64(-libc.Float64FromFloat64(0.0024174248854234894)),
		Fdy:   float32(3.9712762145427405e-07),
		Fe:    int32(FE_INEXACT),
	},
	653: {
		Ffile: __ccgo_ts,
		Fline: int32(670),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0024174225308807775)),
		Fy:    float64(-libc.Float64FromFloat64(0.0024174248854234894)),
		Fdy:   float32(3.9712762145427405e-07),
		Fe:    int32(FE_INEXACT),
	},
	654: {
		Ffile: __ccgo_ts,
		Fline: int32(671),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(333.6840455579976),
		Fy:    libc.Float64FromFloat64(4.131518381646137e+144),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	655: {
		Ffile: __ccgo_ts,
		Fline: int32(672),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(503.7738420434725),
		Fy:    libc.Float64FromFloat64(3.056115271151608e+218),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	656: {
		Ffile: __ccgo_ts,
		Fline: int32(673),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.24036513026713094)),
		Fy:    float64(-libc.Float64FromFloat64(0.24268635737966032)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	657: {
		Ffile: __ccgo_ts,
		Fline: int32(674),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.24036513026713094)),
		Fy:    float64(-libc.Float64FromFloat64(0.2426863573796603)),
		Fdy:   float32(4.5498521217268717e-07),
		Fe:    int32(FE_INEXACT),
	},
	658: {
		Ffile: __ccgo_ts,
		Fline: int32(675),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.24036513026713094)),
		Fy:    float64(-libc.Float64FromFloat64(0.2426863573796603)),
		Fdy:   float32(4.5498521217268717e-07),
		Fe:    int32(FE_INEXACT),
	},
	659: {
		Ffile: __ccgo_ts,
		Fline: int32(676),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.17096874084432426)),
		Fy:    float64(-libc.Float64FromFloat64(0.17180287056409177)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	660: {
		Ffile: __ccgo_ts,
		Fline: int32(677),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.05631102781234609),
		Fy:    libc.Float64FromFloat64(0.056340792269637695),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	661: {
		Ffile: __ccgo_ts,
		Fline: int32(678),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.05631102781234609),
		Fy:    libc.Float64FromFloat64(0.0563407922696377),
		Fdy:   float32(6.492148827419442e-07),
		Fe:    int32(FE_INEXACT),
	},
	662: {
		Ffile: __ccgo_ts,
		Fline: int32(679),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.05631102781234609),
		Fy:    libc.Float64FromFloat64(0.056340792269637695),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	663: {
		Ffile: __ccgo_ts,
		Fline: int32(680),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(114.60790362828183),
		Fy:    libc.Float64FromFloat64(2.9685894015384013e+49),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	664: {
		Ffile: __ccgo_ts,
		Fline: int32(681),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(17.715407989741408),
		Fy:    libc.Float64FromFloat64(2.469869112198705e+07),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	665: {
		Ffile: __ccgo_ts,
		Fline: int32(682),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0020249947703711584),
		Fy:    libc.Float64FromFloat64(0.002024996154321657),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	666: {
		Ffile: __ccgo_ts,
		Fline: int32(683),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0020249947703711584),
		Fy:    libc.Float64FromFloat64(0.0020249961543216573),
		Fdy:   float32(7.575753784294648e-07),
		Fe:    int32(FE_INEXACT),
	},
	667: {
		Ffile: __ccgo_ts,
		Fline: int32(684),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0020249947703711584),
		Fy:    libc.Float64FromFloat64(0.002024996154321657),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	668: {
		Ffile: __ccgo_ts,
		Fline: int32(685),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.34654267679495965),
		Fy:    libc.Float64FromFloat64(0.35352060205987984),
		Fdy:   float32(-libc.Float64FromFloat64(2.3711281471605616e-07)),
		Fe:    int32(FE_INEXACT),
	},
	669: {
		Ffile: __ccgo_ts,
		Fline: int32(686),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.34654267679495965),
		Fy:    libc.Float64FromFloat64(0.3535206020598799),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	670: {
		Ffile: __ccgo_ts,
		Fline: int32(687),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.34654267679495965),
		Fy:    libc.Float64FromFloat64(0.35352060205987984),
		Fdy:   float32(-libc.Float64FromFloat64(2.3711281471605616e-07)),
		Fe:    int32(FE_INEXACT),
	},
	671: {
		Ffile: __ccgo_ts,
		Fline: int32(688),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.3753761515336109)),
		Fy:    float64(-libc.Float64FromFloat64(0.3842540061627809)),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	672: {
		Ffile: __ccgo_ts,
		Fline: int32(689),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.316728455446214)),
		Fy:    float64(-libc.Float64FromFloat64(101.85570739954761)),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	673: {
		Ffile: __ccgo_ts,
		Fline: int32(690),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(74.30224044799778)),
		Fy:    float64(-libc.Float64FromFloat64(9.290156368707012e+31)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	674: {
		Ffile: __ccgo_ts,
		Fline: int32(691),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5314025221728155),
		Fy:    libc.Float64FromFloat64(0.5567683768955477),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	675: {
		Ffile: __ccgo_ts,
		Fline: int32(692),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.017438277458115065)),
		Fy:    float64(-libc.Float64FromFloat64(0.01743916128275111)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	676: {
		Ffile: __ccgo_ts,
		Fline: int32(693),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1086340244609192),
		Fy:    libc.Float64FromFloat64(0.1088478219573358),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	677: {
		Ffile: __ccgo_ts,
		Fline: int32(694),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(196.84401744338598)),
		Fy:    float64(-libc.Float64FromFloat64(1.5390069308642607e+85)),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	678: {
		Ffile: __ccgo_ts,
		Fline: int32(695),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.015890127419060724)),
		Fy:    float64(-libc.Float64FromFloat64(0.01589079612683409)),
		Fdy:   float32(-libc.Float64FromFloat64(6.996571642048366e-07)),
		Fe:    int32(FE_INEXACT),
	},
	679: {
		Ffile: __ccgo_ts,
		Fline: int32(696),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.015890127419060724)),
		Fy:    float64(-libc.Float64FromFloat64(0.015890796126834087)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	680: {
		Ffile: __ccgo_ts,
		Fline: int32(697),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.015890127419060724)),
		Fy:    float64(-libc.Float64FromFloat64(0.015890796126834087)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	681: {
		Ffile: __ccgo_ts,
		Fline: int32(698),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.670462713849081)),
		Fy:    float64(-libc.Float64FromFloat64(145.08266108541056)),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	682: {
		Ffile: __ccgo_ts,
		Fline: int32(699),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(165.64663198364144)),
		Fy:    float64(-libc.Float64FromFloat64(4.3489881137154225e+71)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	683: {
		Ffile: __ccgo_ts,
		Fline: int32(700),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(165.64663198364144)),
		Fy:    float64(-libc.Float64FromFloat64(4.348988113715422e+71)),
		Fdy:   float32(3.342244099258096e-07),
		Fe:    int32(FE_INEXACT),
	},
	684: {
		Ffile: __ccgo_ts,
		Fline: int32(701),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(165.64663198364144)),
		Fy:    float64(-libc.Float64FromFloat64(4.348988113715422e+71)),
		Fdy:   float32(3.342244099258096e-07),
		Fe:    int32(FE_INEXACT),
	},
	685: {
		Ffile: __ccgo_ts,
		Fline: int32(702),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.3661281823436593),
		Fy:    libc.Float64FromFloat64(0.37436308780558114),
		Fdy:   float32(0.49999961256980896),
		Fe:    int32(FE_INEXACT),
	},
	686: {
		Ffile: __ccgo_ts,
		Fline: int32(703),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(101.65268565511897)),
		Fy:    float64(-libc.Float64FromFloat64(7.017306646406745e+43)),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	687: {
		Ffile: __ccgo_ts,
		Fline: int32(704),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(11.010356254428826),
		Fy:    libc.Float64FromFloat64(30248.71773381185),
		Fdy:   float32(-libc.Float64FromFloat64(5.53638983546989e-07)),
		Fe:    int32(FE_INEXACT),
	},
	688: {
		Ffile: __ccgo_ts,
		Fline: int32(705),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(11.010356254428826),
		Fy:    libc.Float64FromFloat64(30248.717733811853),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	689: {
		Ffile: __ccgo_ts,
		Fline: int32(706),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.010356254428826),
		Fy:    libc.Float64FromFloat64(30248.71773381185),
		Fdy:   float32(-libc.Float64FromFloat64(5.53638983546989e-07)),
		Fe:    int32(FE_INEXACT),
	},
	690: {
		Ffile: __ccgo_ts,
		Fline: int32(707),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(138.86503316297095)),
		Fy:    float64(-libc.Float64FromFloat64(1.016921984551256e+60)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	691: {
		Ffile: __ccgo_ts,
		Fline: int32(708),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(428.8980014672162),
		Fy:    libc.Float64FromFloat64(9.268412214474193e+185),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	692: {
		Ffile: __ccgo_ts,
		Fline: int32(709),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.9298301808954798)),
		Fy:    float64(-libc.Float64FromFloat64(1.069729043631862)),
		Fdy:   float32(-libc.Float64FromFloat64(2.985236733366037e-07)),
		Fe:    int32(FE_INEXACT),
	},
	693: {
		Ffile: __ccgo_ts,
		Fline: int32(710),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.9298301808954798)),
		Fy:    float64(-libc.Float64FromFloat64(1.0697290436318618)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	694: {
		Ffile: __ccgo_ts,
		Fline: int32(711),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.9298301808954798)),
		Fy:    float64(-libc.Float64FromFloat64(1.0697290436318618)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	695: {
		Ffile: __ccgo_ts,
		Fline: int32(712),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(119.02006053315054)),
		Fy:    float64(-libc.Float64FromFloat64(2.4475159442215125e+51)),
		Fdy:   float32(-libc.Float64FromFloat64(7.997999773579068e-07)),
		Fe:    int32(FE_INEXACT),
	},
	696: {
		Ffile: __ccgo_ts,
		Fline: int32(713),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(119.02006053315054)),
		Fy:    float64(-libc.Float64FromFloat64(2.447515944221512e+51)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	697: {
		Ffile: __ccgo_ts,
		Fline: int32(714),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(119.02006053315054)),
		Fy:    float64(-libc.Float64FromFloat64(2.447515944221512e+51)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	698: {
		Ffile: __ccgo_ts,
		Fline: int32(715),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0030745412493466823)),
		Fy:    float64(-libc.Float64FromFloat64(0.003074546093188221)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	699: {
		Ffile: __ccgo_ts,
		Fline: int32(716),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0030745412493466823)),
		Fy:    float64(-libc.Float64FromFloat64(0.0030745460931882206)),
		Fdy:   float32(1.6535850022592058e-07),
		Fe:    int32(FE_INEXACT),
	},
	700: {
		Ffile: __ccgo_ts,
		Fline: int32(717),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0030745412493466823)),
		Fy:    float64(-libc.Float64FromFloat64(0.0030745460931882206)),
		Fdy:   float32(1.6535850022592058e-07),
		Fe:    int32(FE_INEXACT),
	},
	701: {
		Ffile: __ccgo_ts,
		Fline: int32(718),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(14.6562675600978)),
		Fy:    float64(-libc.Float64FromFloat64(1.1590605676956256e+06)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	702: {
		Ffile: __ccgo_ts,
		Fline: int32(719),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(632.63386871284),
		Fy:    libc.Float64FromFloat64(2.807813458816772e+274),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	703: {
		Ffile: __ccgo_ts,
		Fline: int32(720),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(632.63386871284),
		Fy:    libc.Float64FromFloat64(2.8078134588167726e+274),
		Fdy:   float32(6.613236678276735e-07),
		Fe:    int32(FE_INEXACT),
	},
	704: {
		Ffile: __ccgo_ts,
		Fline: int32(721),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(632.63386871284),
		Fy:    libc.Float64FromFloat64(2.807813458816772e+274),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	705: {
		Ffile: __ccgo_ts,
		Fline: int32(722),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.32634249679560534)),
		Fy:    float64(-libc.Float64FromFloat64(0.33216596833271494)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	706: {
		Ffile: __ccgo_ts,
		Fline: int32(723),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.025542510984535072),
		Fy:    libc.Float64FromFloat64(0.025545288482076496),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	707: {
		Ffile: __ccgo_ts,
		Fline: int32(724),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.08090991196034043),
		Fy:    libc.Float64FromFloat64(0.08099821915491343),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	708: {
		Ffile: __ccgo_ts,
		Fline: int32(725),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.40479997153620806),
		Fy:    libc.Float64FromFloat64(0.4159461941619743),
		Fdy:   float32(-libc.Float64FromFloat64(2.0309707338128646e-07)),
		Fe:    int32(FE_INEXACT),
	},
	709: {
		Ffile: __ccgo_ts,
		Fline: int32(726),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.40479997153620806),
		Fy:    libc.Float64FromFloat64(0.41594619416197437),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	710: {
		Ffile: __ccgo_ts,
		Fline: int32(727),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.40479997153620806),
		Fy:    libc.Float64FromFloat64(0.4159461941619743),
		Fdy:   float32(-libc.Float64FromFloat64(2.0309707338128646e-07)),
		Fe:    int32(FE_INEXACT),
	},
	711: {
		Ffile: __ccgo_ts,
		Fline: int32(728),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(105.19578751322277),
		Fy:    libc.Float64FromFloat64(2.4261633660542748e+45),
		Fdy:   float32(-libc.Float64FromFloat64(2.799187086566235e-07)),
		Fe:    int32(FE_INEXACT),
	},
	712: {
		Ffile: __ccgo_ts,
		Fline: int32(729),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(105.19578751322277),
		Fy:    libc.Float64FromFloat64(2.426163366054275e+45),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	713: {
		Ffile: __ccgo_ts,
		Fline: int32(730),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(105.19578751322277),
		Fy:    libc.Float64FromFloat64(2.4261633660542748e+45),
		Fdy:   float32(-libc.Float64FromFloat64(2.799187086566235e-07)),
		Fe:    int32(FE_INEXACT),
	},
	714: {
		Ffile: __ccgo_ts,
		Fline: int32(731),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.10040122306934736)),
		Fy:    float64(-libc.Float64FromFloat64(0.10056998895009632)),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	715: {
		Ffile: __ccgo_ts,
		Fline: int32(732),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.13557488530504533),
		Fy:    libc.Float64FromFloat64(0.13599059064811453),
		Fdy:   float32(-libc.Float64FromFloat64(4.7487404231105756e-08)),
		Fe:    int32(FE_INEXACT),
	},
	716: {
		Ffile: __ccgo_ts,
		Fline: int32(733),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.13557488530504533),
		Fy:    libc.Float64FromFloat64(0.13599059064811456),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	717: {
		Ffile: __ccgo_ts,
		Fline: int32(734),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.13557488530504533),
		Fy:    libc.Float64FromFloat64(0.13599059064811453),
		Fdy:   float32(-libc.Float64FromFloat64(4.7487404231105756e-08)),
		Fe:    int32(FE_INEXACT),
	},
	718: {
		Ffile: __ccgo_ts,
		Fline: int32(735),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(220.50145579377494),
		Fy:    libc.Float64FromFloat64(2.8942464362454596e+95),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	719: {
		Ffile: __ccgo_ts,
		Fline: int32(736),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(220.50145579377494),
		Fy:    libc.Float64FromFloat64(2.89424643624546e+95),
		Fdy:   float32(5.036933430346835e-07),
		Fe:    int32(FE_INEXACT),
	},
	720: {
		Ffile: __ccgo_ts,
		Fline: int32(737),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(220.50145579377494),
		Fy:    libc.Float64FromFloat64(2.8942464362454596e+95),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	721: {
		Ffile: __ccgo_ts,
		Fline: int32(738),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(271.746380013166)),
		Fy:    float64(-libc.Float64FromFloat64(5.211026974163252e+117)),
		Fdy:   float32(-libc.Float64FromFloat64(2.0441356696210278e-07)),
		Fe:    int32(FE_INEXACT),
	},
	722: {
		Ffile: __ccgo_ts,
		Fline: int32(739),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(271.746380013166)),
		Fy:    float64(-libc.Float64FromFloat64(5.211026974163251e+117)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	723: {
		Ffile: __ccgo_ts,
		Fline: int32(740),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(271.746380013166)),
		Fy:    float64(-libc.Float64FromFloat64(5.211026974163251e+117)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	724: {
		Ffile: __ccgo_ts,
		Fline: int32(741),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.014454496268462368)),
		Fy:    float64(-libc.Float64FromFloat64(0.01445499960930351)),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	725: {
		Ffile: __ccgo_ts,
		Fline: int32(742),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.09299196793754061),
		Fy:    libc.Float64FromFloat64(0.09312605066681913),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	726: {
		Ffile: __ccgo_ts,
		Fline: int32(743),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.09299196793754061),
		Fy:    libc.Float64FromFloat64(0.09312605066681914),
		Fdy:   float32(5.39329732873739e-07),
		Fe:    int32(FE_INEXACT),
	},
	727: {
		Ffile: __ccgo_ts,
		Fline: int32(744),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.09299196793754061),
		Fy:    libc.Float64FromFloat64(0.09312605066681913),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	728: {
		Ffile: __ccgo_ts,
		Fline: int32(745),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(42.8175282810755)),
		Fy:    float64(-libc.Float64FromFloat64(1.9696373246176963e+18)),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	729: {
		Ffile: __ccgo_ts,
		Fline: int32(746),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.003866892383516291),
		Fy:    libc.Float64FromFloat64(0.003866902020371447),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	730: {
		Ffile: __ccgo_ts,
		Fline: int32(747),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0020614314886876635)),
		Fy:    float64(-libc.Float64FromFloat64(0.002061432948696751)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	731: {
		Ffile: __ccgo_ts,
		Fline: int32(748),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9817066803548),
		Fy:    libc.Float64FromFloat64(1.1471681953318726),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	732: {
		Ffile: __ccgo_ts,
		Fline: int32(749),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9817066803548),
		Fy:    libc.Float64FromFloat64(1.1471681953318729),
		Fdy:   float32(4.407307585552189e-08),
		Fe:    int32(FE_INEXACT),
	},
	733: {
		Ffile: __ccgo_ts,
		Fline: int32(750),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9817066803548),
		Fy:    libc.Float64FromFloat64(1.1471681953318726),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	734: {
		Ffile: __ccgo_ts,
		Fline: int32(751),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(107.09876111679682),
		Fy:    libc.Float64FromFloat64(1.6269378990407906e+46),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	735: {
		Ffile: __ccgo_ts,
		Fline: int32(752),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.010040750094247182),
		Fy:    libc.Float64FromFloat64(0.010040918807583142),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	736: {
		Ffile: __ccgo_ts,
		Fline: int32(753),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.010040750094247182),
		Fy:    libc.Float64FromFloat64(0.010040918807583144),
		Fdy:   float32(9.521425567982078e-07),
		Fe:    int32(FE_INEXACT),
	},
	737: {
		Ffile: __ccgo_ts,
		Fline: int32(754),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.010040750094247182),
		Fy:    libc.Float64FromFloat64(0.010040918807583142),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	738: {
		Ffile: __ccgo_ts,
		Fline: int32(755),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.009313031265553771)),
		Fy:    float64(-libc.Float64FromFloat64(0.00931316588996467)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	739: {
		Ffile: __ccgo_ts,
		Fline: int32(756),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.009313031265553771)),
		Fy:    float64(-libc.Float64FromFloat64(0.009313165889964668)),
		Fdy:   float32(6.531234362228133e-07),
		Fe:    int32(FE_INEXACT),
	},
	740: {
		Ffile: __ccgo_ts,
		Fline: int32(757),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.009313031265553771)),
		Fy:    float64(-libc.Float64FromFloat64(0.009313165889964668)),
		Fdy:   float32(6.531234362228133e-07),
		Fe:    int32(FE_INEXACT),
	},
	741: {
		Ffile: __ccgo_ts,
		Fline: int32(758),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(30.283160948908858),
		Fy:    libc.Float64FromFloat64(7.092179176029353e+12),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	742: {
		Ffile: __ccgo_ts,
		Fline: int32(759),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(30.283160948908858),
		Fy:    libc.Float64FromFloat64(7.092179176029354e+12),
		Fdy:   float32(3.502055676563032e-07),
		Fe:    int32(FE_INEXACT),
	},
	743: {
		Ffile: __ccgo_ts,
		Fline: int32(760),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(30.283160948908858),
		Fy:    libc.Float64FromFloat64(7.092179176029353e+12),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	744: {
		Ffile: __ccgo_ts,
		Fline: int32(761),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(307.5166543927201),
		Fy:    libc.Float64FromFloat64(1.785484566025667e+133),
		Fdy:   float32(-libc.Float64FromFloat64(6.869610302828733e-08)),
		Fe:    int32(FE_INEXACT),
	},
	745: {
		Ffile: __ccgo_ts,
		Fline: int32(762),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(307.5166543927201),
		Fy:    libc.Float64FromFloat64(1.7854845660256674e+133),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	746: {
		Ffile: __ccgo_ts,
		Fline: int32(763),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(307.5166543927201),
		Fy:    libc.Float64FromFloat64(1.785484566025667e+133),
		Fdy:   float32(-libc.Float64FromFloat64(6.869610302828733e-08)),
		Fe:    int32(FE_INEXACT),
	},
	747: {
		Ffile: __ccgo_ts,
		Fline: int32(764),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(37.72296144884513)),
		Fy:    float64(-libc.Float64FromFloat64(1.2073800696130198e+16)),
		Fdy:   float32(-libc.Float64FromFloat64(5.396415190261905e-07)),
		Fe:    int32(FE_INEXACT),
	},
	748: {
		Ffile: __ccgo_ts,
		Fline: int32(765),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(37.72296144884513)),
		Fy:    float64(-libc.Float64FromFloat64(1.2073800696130196e+16)),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	749: {
		Ffile: __ccgo_ts,
		Fline: int32(766),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(37.72296144884513)),
		Fy:    float64(-libc.Float64FromFloat64(1.2073800696130196e+16)),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	750: {
		Ffile: __ccgo_ts,
		Fline: int32(767),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(379.4237324201999),
		Fy:    libc.Float64FromFloat64(3.0241497833340333e+164),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	751: {
		Ffile: __ccgo_ts,
		Fline: int32(768),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(14.220984932816732)),
		Fy:    float64(-libc.Float64FromFloat64(750006.9500900891)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	752: {
		Ffile: __ccgo_ts,
		Fline: int32(769),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.48253690277988304)),
		Fy:    float64(-libc.Float64FromFloat64(0.5014819213203519)),
		Fdy:   float32(0.4999990463256836),
		Fe:    int32(FE_INEXACT),
	},
	753: {
		Ffile: __ccgo_ts,
		Fline: int32(770),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(48.35457064429792),
		Fy:    libc.Float64FromFloat64(5.0014186583359016e+20),
		Fdy:   float32(-libc.Float64FromFloat64(9.020760671774042e-07)),
		Fe:    int32(FE_INEXACT),
	},
	754: {
		Ffile: __ccgo_ts,
		Fline: int32(771),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(48.35457064429792),
		Fy:    libc.Float64FromFloat64(5.001418658335902e+20),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	755: {
		Ffile: __ccgo_ts,
		Fline: int32(772),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(48.35457064429792),
		Fy:    libc.Float64FromFloat64(5.0014186583359016e+20),
		Fdy:   float32(-libc.Float64FromFloat64(9.020760671774042e-07)),
		Fe:    int32(FE_INEXACT),
	},
	756: {
		Ffile: __ccgo_ts,
		Fline: int32(773),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(110.74501155929501)),
		Fy:    float64(-libc.Float64FromFloat64(6.2361624924426995e+47)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	757: {
		Ffile: __ccgo_ts,
		Fline: int32(774),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(450.5262807874123),
		Fy:    libc.Float64FromFloat64(2.291119294426709e+195),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	758: {
		Ffile: __ccgo_ts,
		Fline: int32(775),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.005916599585397088),
		Fy:    libc.Float64FromFloat64(0.005916634105020248),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	759: {
		Ffile: __ccgo_ts,
		Fline: int32(776),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.005916599585397088),
		Fy:    libc.Float64FromFloat64(0.005916634105020249),
		Fdy:   float32(9.353823884339363e-07),
		Fe:    int32(FE_INEXACT),
	},
	760: {
		Ffile: __ccgo_ts,
		Fline: int32(777),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.005916599585397088),
		Fy:    libc.Float64FromFloat64(0.005916634105020248),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	761: {
		Ffile: __ccgo_ts,
		Fline: int32(778),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(180.83445478345396),
		Fy:    libc.Float64FromFloat64(1.7154412726504424e+78),
		Fdy:   float32(-libc.Float64FromFloat64(6.918447184034449e-07)),
		Fe:    int32(FE_INEXACT),
	},
	762: {
		Ffile: __ccgo_ts,
		Fline: int32(779),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(180.83445478345396),
		Fy:    libc.Float64FromFloat64(1.7154412726504426e+78),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	763: {
		Ffile: __ccgo_ts,
		Fline: int32(780),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(180.83445478345396),
		Fy:    libc.Float64FromFloat64(1.7154412726504424e+78),
		Fdy:   float32(-libc.Float64FromFloat64(6.918447184034449e-07)),
		Fe:    int32(FE_INEXACT),
	},
	764: {
		Ffile: __ccgo_ts,
		Fline: int32(781),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.4417384452372595)),
		Fy:    float64(-libc.Float64FromFloat64(115.41940552659561)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	765: {
		Ffile: __ccgo_ts,
		Fline: int32(782),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.4417384452372595)),
		Fy:    float64(-libc.Float64FromFloat64(115.4194055265956)),
		Fdy:   float32(6.521652835544955e-07),
		Fe:    int32(FE_INEXACT),
	},
	766: {
		Ffile: __ccgo_ts,
		Fline: int32(783),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(5.4417384452372595)),
		Fy:    float64(-libc.Float64FromFloat64(115.4194055265956)),
		Fdy:   float32(6.521652835544955e-07),
		Fe:    int32(FE_INEXACT),
	},
	767: {
		Ffile: __ccgo_ts,
		Fline: int32(784),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01679467380219048),
		Fy:    libc.Float64FromFloat64(0.016795463333930442),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	768: {
		Ffile: __ccgo_ts,
		Fline: int32(785),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.2381458477085943)),
		Fy:    float64(-libc.Float64FromFloat64(0.24040325136390694)),
		Fdy:   float32(-libc.Float64FromFloat64(9.422043234508237e-08)),
		Fe:    int32(FE_INEXACT),
	},
	769: {
		Ffile: __ccgo_ts,
		Fline: int32(786),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.2381458477085943)),
		Fy:    float64(-libc.Float64FromFloat64(0.24040325136390692)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	770: {
		Ffile: __ccgo_ts,
		Fline: int32(787),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.2381458477085943)),
		Fy:    float64(-libc.Float64FromFloat64(0.24040325136390692)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	771: {
		Ffile: __ccgo_ts,
		Fline: int32(788),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(330.945972363611)),
		Fy:    float64(-libc.Float64FromFloat64(2.6728809144772e+143)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	772: {
		Ffile: __ccgo_ts,
		Fline: int32(789),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(330.945972363611)),
		Fy:    float64(-libc.Float64FromFloat64(2.6728809144771996e+143)),
		Fdy:   float32(2.0904579400848888e-07),
		Fe:    int32(FE_INEXACT),
	},
	773: {
		Ffile: __ccgo_ts,
		Fline: int32(790),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(330.945972363611)),
		Fy:    float64(-libc.Float64FromFloat64(2.6728809144771996e+143)),
		Fdy:   float32(2.0904579400848888e-07),
		Fe:    int32(FE_INEXACT),
	},
	774: {
		Ffile: __ccgo_ts,
		Fline: int32(791),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0025290386274973695)),
		Fy:    float64(-libc.Float64FromFloat64(0.0025290413234687428)),
		Fdy:   float32(-libc.Float64FromFloat64(6.545114956679754e-07)),
		Fe:    int32(FE_INEXACT),
	},
	775: {
		Ffile: __ccgo_ts,
		Fline: int32(792),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0025290386274973695)),
		Fy:    float64(-libc.Float64FromFloat64(0.0025290413234687423)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	776: {
		Ffile: __ccgo_ts,
		Fline: int32(793),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0025290386274973695)),
		Fy:    float64(-libc.Float64FromFloat64(0.0025290413234687423)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	777: {
		Ffile: __ccgo_ts,
		Fline: int32(794),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.07827475010925083)),
		Fy:    float64(-libc.Float64FromFloat64(0.07835470533665716)),
		Fdy:   float32(-libc.Float64FromFloat64(3.789658649111516e-07)),
		Fe:    int32(FE_INEXACT),
	},
	778: {
		Ffile: __ccgo_ts,
		Fline: int32(795),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.07827475010925083)),
		Fy:    float64(-libc.Float64FromFloat64(0.07835470533665714)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	779: {
		Ffile: __ccgo_ts,
		Fline: int32(796),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.07827475010925083)),
		Fy:    float64(-libc.Float64FromFloat64(0.07835470533665714)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	780: {
		Ffile: __ccgo_ts,
		Fline: int32(797),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(34.6997421597002)),
		Fy:    float64(-libc.Float64FromFloat64(5.873223769575488e+14)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	781: {
		Ffile: __ccgo_ts,
		Fline: int32(798),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.6005204366903198)),
		Fy:    float64(-libc.Float64FromFloat64(2.3769096820560525)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999994933605194)),
		Fe:    int32(FE_INEXACT),
	},
	782: {
		Ffile: __ccgo_ts,
		Fline: int32(799),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(42.492031109923)),
		Fy:    float64(-libc.Float64FromFloat64(1.4224094870025725e+18)),
		Fdy:   float32(-libc.Float64FromFloat64(2.5092015221162e-07)),
		Fe:    int32(FE_INEXACT),
	},
	783: {
		Ffile: __ccgo_ts,
		Fline: int32(800),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(42.492031109923)),
		Fy:    float64(-libc.Float64FromFloat64(1.4224094870025723e+18)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	784: {
		Ffile: __ccgo_ts,
		Fline: int32(801),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(42.492031109923)),
		Fy:    float64(-libc.Float64FromFloat64(1.4224094870025723e+18)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	785: {
		Ffile: __ccgo_ts,
		Fline: int32(802),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.012719705068725863),
		Fy:    libc.Float64FromFloat64(0.012720048060249347),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	786: {
		Ffile: __ccgo_ts,
		Fline: int32(803),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6607962602819702),
		Fy:    libc.Float64FromFloat64(0.7099467980523012),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	787: {
		Ffile: __ccgo_ts,
		Fline: int32(804),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(144.05384859109014)),
		Fy:    float64(-libc.Float64FromFloat64(1.82289454104725e+62)),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	788: {
		Ffile: __ccgo_ts,
		Fline: int32(805),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.001856462339058039)),
		Fy:    float64(-libc.Float64FromFloat64(0.0018564634054264084)),
		Fdy:   float32(-libc.Float64FromFloat64(5.714080089092022e-07)),
		Fe:    int32(FE_INEXACT),
	},
	789: {
		Ffile: __ccgo_ts,
		Fline: int32(806),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.001856462339058039)),
		Fy:    float64(-libc.Float64FromFloat64(0.0018564634054264082)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	790: {
		Ffile: __ccgo_ts,
		Fline: int32(807),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.001856462339058039)),
		Fy:    float64(-libc.Float64FromFloat64(0.0018564634054264082)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	791: {
		Ffile: __ccgo_ts,
		Fline: int32(808),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(197.1059003891587)),
		Fy:    float64(-libc.Float64FromFloat64(1.9997462631868175e+85)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	792: {
		Ffile: __ccgo_ts,
		Fline: int32(809),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.4199497639219757)),
		Fy:    float64(-libc.Float64FromFloat64(0.43240263603007195)),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	793: {
		Ffile: __ccgo_ts,
		Fline: int32(810),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(13.846146880910164),
		Fy:    libc.Float64FromFloat64(515555.22224026266),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	794: {
		Ffile: __ccgo_ts,
		Fline: int32(811),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(207.28334277540935),
		Fy:    libc.Float64FromFloat64(5.259954194171787e+89),
		Fdy:   float32(-libc.Float64FromFloat64(6.81058907048282e-07)),
		Fe:    int32(FE_INEXACT),
	},
	795: {
		Ffile: __ccgo_ts,
		Fline: int32(812),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(207.28334277540935),
		Fy:    libc.Float64FromFloat64(5.259954194171788e+89),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	796: {
		Ffile: __ccgo_ts,
		Fline: int32(813),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(207.28334277540935),
		Fy:    libc.Float64FromFloat64(5.259954194171787e+89),
		Fdy:   float32(-libc.Float64FromFloat64(6.81058907048282e-07)),
		Fe:    int32(FE_INEXACT),
	},
	797: {
		Ffile: __ccgo_ts,
		Fline: int32(814),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(105.96987691400354)),
		Fy:    float64(-libc.Float64FromFloat64(5.261417876585533e+45)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999913573265076)),
		Fe:    int32(FE_INEXACT),
	},
	798: {
		Ffile: __ccgo_ts,
		Fline: int32(815),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0017595767687086708)),
		Fy:    float64(-libc.Float64FromFloat64(0.0017595776766828016)),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	799: {
		Ffile: __ccgo_ts,
		Fline: int32(816),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0026859608443908474)),
		Fy:    float64(-libc.Float64FromFloat64(0.0026859640739849114)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	800: {
		Ffile: __ccgo_ts,
		Fline: int32(817),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.060353617314336395)),
		Fy:    float64(-libc.Float64FromFloat64(0.060390264258024295)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999994933605194)),
		Fe:    int32(FE_INEXACT),
	},
	801: {
		Ffile: __ccgo_ts,
		Fline: int32(818),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(567.694333535856),
		Fy:    libc.Float64FromFloat64(1.7598938218938464e+246),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	802: {
		Ffile: __ccgo_ts,
		Fline: int32(819),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(567.694333535856),
		Fy:    libc.Float64FromFloat64(1.7598938218938468e+246),
		Fdy:   float32(6.865279829071369e-07),
		Fe:    int32(FE_INEXACT),
	},
	803: {
		Ffile: __ccgo_ts,
		Fline: int32(820),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(567.694333535856),
		Fy:    libc.Float64FromFloat64(1.7598938218938464e+246),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	804: {
		Ffile: __ccgo_ts,
		Fline: int32(821),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01423799121705459)),
		Fy:    float64(-libc.Float64FromFloat64(0.014238472277128569)),
		Fdy:   float32(-libc.Float64FromFloat64(1.0104984937697736e-07)),
		Fe:    int32(FE_INEXACT),
	},
	805: {
		Ffile: __ccgo_ts,
		Fline: int32(822),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01423799121705459)),
		Fy:    float64(-libc.Float64FromFloat64(0.014238472277128567)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	806: {
		Ffile: __ccgo_ts,
		Fline: int32(823),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.01423799121705459)),
		Fy:    float64(-libc.Float64FromFloat64(0.014238472277128567)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	807: {
		Ffile: __ccgo_ts,
		Fline: int32(824),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(238.94879514101186)),
		Fy:    float64(-libc.Float64FromFloat64(2.9724406471465195e+103)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	808: {
		Ffile: __ccgo_ts,
		Fline: int32(825),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.6198871666384125)),
		Fy:    float64(-libc.Float64FromFloat64(2.4272995808216575)),
		Fdy:   float32(-libc.Float64FromFloat64(8.690166168889846e-07)),
		Fe:    int32(FE_INEXACT),
	},
	809: {
		Ffile: __ccgo_ts,
		Fline: int32(826),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.6198871666384125)),
		Fy:    float64(-libc.Float64FromFloat64(2.427299580821657)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	810: {
		Ffile: __ccgo_ts,
		Fline: int32(827),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.6198871666384125)),
		Fy:    float64(-libc.Float64FromFloat64(2.427299580821657)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	811: {
		Ffile: __ccgo_ts,
		Fline: int32(828),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(206.26828047766193),
		Fy:    libc.Float64FromFloat64(1.9061014312510575e+89),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	812: {
		Ffile: __ccgo_ts,
		Fline: int32(829),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(206.26828047766193),
		Fy:    libc.Float64FromFloat64(1.9061014312510578e+89),
		Fdy:   float32(6.262856118155469e-07),
		Fe:    int32(FE_INEXACT),
	},
	813: {
		Ffile: __ccgo_ts,
		Fline: int32(830),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(206.26828047766193),
		Fy:    libc.Float64FromFloat64(1.9061014312510575e+89),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	814: {
		Ffile: __ccgo_ts,
		Fline: int32(831),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0011595271849870331)),
		Fy:    float64(-libc.Float64FromFloat64(0.0011595274448184037)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	815: {
		Ffile: __ccgo_ts,
		Fline: int32(832),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.9778345916655784),
		Fy:    libc.Float64FromFloat64(26.69127361515739),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999916553497314)),
		Fe:    int32(FE_INEXACT),
	},
	816: {
		Ffile: __ccgo_ts,
		Fline: int32(833),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(56.91823155450439)),
		Fy:    float64(-libc.Float64FromFloat64(2.6196537225888035e+24)),
		Fdy:   float32(-libc.Float64FromFloat64(3.749669872377126e-07)),
		Fe:    int32(FE_INEXACT),
	},
	817: {
		Ffile: __ccgo_ts,
		Fline: int32(834),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(56.91823155450439)),
		Fy:    float64(-libc.Float64FromFloat64(2.619653722588803e+24)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	818: {
		Ffile: __ccgo_ts,
		Fline: int32(835),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(56.91823155450439)),
		Fy:    float64(-libc.Float64FromFloat64(2.619653722588803e+24)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	819: {
		Ffile: __ccgo_ts,
		Fline: int32(836),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.0968259402528435),
		Fy:    libc.Float64FromFloat64(604.0627447398039),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999934434890747)),
		Fe:    int32(FE_INEXACT),
	},
	820: {
		Ffile: __ccgo_ts,
		Fline: int32(837),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(233.94810982439478),
		Fy:    libc.Float64FromFloat64(2.0014426618207477e+101),
		Fdy:   float32(-libc.Float64FromFloat64(4.6592833768954733e-07)),
		Fe:    int32(FE_INEXACT),
	},
	821: {
		Ffile: __ccgo_ts,
		Fline: int32(838),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(233.94810982439478),
		Fy:    libc.Float64FromFloat64(2.001442661820748e+101),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	822: {
		Ffile: __ccgo_ts,
		Fline: int32(839),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(233.94810982439478),
		Fy:    libc.Float64FromFloat64(2.0014426618207477e+101),
		Fdy:   float32(-libc.Float64FromFloat64(4.6592833768954733e-07)),
		Fe:    int32(FE_INEXACT),
	},
	823: {
		Ffile: __ccgo_ts,
		Fline: int32(840),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.014805962212000402),
		Fy:    libc.Float64FromFloat64(0.014806503169840928),
		Fdy:   float32(-libc.Float64FromFloat64(6.813569797259333e-08)),
		Fe:    int32(FE_INEXACT),
	},
	824: {
		Ffile: __ccgo_ts,
		Fline: int32(841),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.014805962212000402),
		Fy:    libc.Float64FromFloat64(0.01480650316984093),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	825: {
		Ffile: __ccgo_ts,
		Fline: int32(842),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.014805962212000402),
		Fy:    libc.Float64FromFloat64(0.014806503169840928),
		Fdy:   float32(-libc.Float64FromFloat64(6.813569797259333e-08)),
		Fe:    int32(FE_INEXACT),
	},
	826: {
		Ffile: __ccgo_ts,
		Fline: int32(843),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.6759972635211797),
		Fy:    libc.Float64FromFloat64(7.228995775613254),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	827: {
		Ffile: __ccgo_ts,
		Fline: int32(844),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.6759972635211797),
		Fy:    libc.Float64FromFloat64(7.228995775613255),
		Fdy:   float32(2.717883660352527e-07),
		Fe:    int32(FE_INEXACT),
	},
	828: {
		Ffile: __ccgo_ts,
		Fline: int32(845),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6759972635211797),
		Fy:    libc.Float64FromFloat64(7.228995775613254),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	829: {
		Ffile: __ccgo_ts,
		Fline: int32(846),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.09320357349888254)),
		Fy:    float64(-libc.Float64FromFloat64(0.09333857390439912)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	830: {
		Ffile: __ccgo_ts,
		Fline: int32(847),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0017926590450952772)),
		Fy:    float64(-libc.Float64FromFloat64(0.0017926600052515192)),
		Fdy:   float32(-libc.Float64FromFloat64(8.702125455783971e-07)),
		Fe:    int32(FE_INEXACT),
	},
	831: {
		Ffile: __ccgo_ts,
		Fline: int32(848),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0017926590450952772)),
		Fy:    float64(-libc.Float64FromFloat64(0.001792660005251519)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	832: {
		Ffile: __ccgo_ts,
		Fline: int32(849),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0017926590450952772)),
		Fy:    float64(-libc.Float64FromFloat64(0.001792660005251519)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	833: {
		Ffile: __ccgo_ts,
		Fline: int32(850),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.293478466169761),
		Fy:    libc.Float64FromFloat64(13.449636872441777),
		Fdy:   float32(-libc.Float64FromFloat64(3.266482053732034e-07)),
		Fe:    int32(FE_INEXACT),
	},
	834: {
		Ffile: __ccgo_ts,
		Fline: int32(851),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.293478466169761),
		Fy:    libc.Float64FromFloat64(13.44963687244178),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	835: {
		Ffile: __ccgo_ts,
		Fline: int32(852),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.293478466169761),
		Fy:    libc.Float64FromFloat64(13.449636872441777),
		Fdy:   float32(-libc.Float64FromFloat64(3.266482053732034e-07)),
		Fe:    int32(FE_INEXACT),
	},
	836: {
		Ffile: __ccgo_ts,
		Fline: int32(853),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(253.8724210348281)),
		Fy:    float64(-libc.Float64FromFloat64(9.00246760344834e+109)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	837: {
		Ffile: __ccgo_ts,
		Fline: int32(854),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(253.8724210348281)),
		Fy:    float64(-libc.Float64FromFloat64(9.002467603448338e+109)),
		Fdy:   float32(3.031455264590477e-07),
		Fe:    int32(FE_INEXACT),
	},
	838: {
		Ffile: __ccgo_ts,
		Fline: int32(855),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(253.8724210348281)),
		Fy:    float64(-libc.Float64FromFloat64(9.002467603448338e+109)),
		Fdy:   float32(3.031455264590477e-07),
		Fe:    int32(FE_INEXACT),
	},
	839: {
		Ffile: __ccgo_ts,
		Fline: int32(856),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.19724365404904126),
		Fy:    libc.Float64FromFloat64(0.19852510692918648),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	840: {
		Ffile: __ccgo_ts,
		Fline: int32(857),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.19724365404904126),
		Fy:    libc.Float64FromFloat64(0.1985251069291865),
		Fdy:   float32(5.566145659940958e-07),
		Fe:    int32(FE_INEXACT),
	},
	841: {
		Ffile: __ccgo_ts,
		Fline: int32(858),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.19724365404904126),
		Fy:    libc.Float64FromFloat64(0.19852510692918648),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	842: {
		Ffile: __ccgo_ts,
		Fline: int32(859),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0798641099483094),
		Fy:    libc.Float64FromFloat64(0.07994903625134969),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	843: {
		Ffile: __ccgo_ts,
		Fline: int32(860),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.002930899932110995)),
		Fy:    float64(-libc.Float64FromFloat64(0.002930904128269731)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	844: {
		Ffile: __ccgo_ts,
		Fline: int32(861),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.28568547078056383)),
		Fy:    float64(-libc.Float64FromFloat64(0.2895874532301797)),
		Fdy:   float32(-libc.Float64FromFloat64(1.4576170315194759e-07)),
		Fe:    int32(FE_INEXACT),
	},
	845: {
		Ffile: __ccgo_ts,
		Fline: int32(862),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.28568547078056383)),
		Fy:    float64(-libc.Float64FromFloat64(0.28958745323017965)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	846: {
		Ffile: __ccgo_ts,
		Fline: int32(863),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.28568547078056383)),
		Fy:    float64(-libc.Float64FromFloat64(0.28958745323017965)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	847: {
		Ffile: __ccgo_ts,
		Fline: int32(864),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(498.69545571174757)),
		Fy:    float64(-libc.Float64FromFloat64(1.9039457522039418e+216)),
		Fdy:   float32(0.4999992549419403),
		Fe:    int32(FE_INEXACT),
	},
	848: {
		Ffile: __ccgo_ts,
		Fline: int32(865),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0012301425255486155),
		Fy:    libc.Float64FromFloat64(0.0012301428358009647),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	849: {
		Ffile: __ccgo_ts,
		Fline: int32(866),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0012301425255486155),
		Fy:    libc.Float64FromFloat64(0.001230142835800965),
		Fdy:   float32(6.502892802018323e-07),
		Fe:    int32(FE_INEXACT),
	},
	850: {
		Ffile: __ccgo_ts,
		Fline: int32(867),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0012301425255486155),
		Fy:    libc.Float64FromFloat64(0.0012301428358009647),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	851: {
		Ffile: __ccgo_ts,
		Fline: int32(868),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.061294431680037),
		Fy:    libc.Float64FromFloat64(214.46398362622077),
		Fdy:   float32(-libc.Float64FromFloat64(6.119120143921464e-07)),
		Fe:    int32(FE_INEXACT),
	},
	852: {
		Ffile: __ccgo_ts,
		Fline: int32(869),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.061294431680037),
		Fy:    libc.Float64FromFloat64(214.4639836262208),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	853: {
		Ffile: __ccgo_ts,
		Fline: int32(870),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.061294431680037),
		Fy:    libc.Float64FromFloat64(214.46398362622077),
		Fdy:   float32(-libc.Float64FromFloat64(6.119120143921464e-07)),
		Fe:    int32(FE_INEXACT),
	},
	854: {
		Ffile: __ccgo_ts,
		Fline: int32(871),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.019102183945165174)),
		Fy:    float64(-libc.Float64FromFloat64(0.019103345676601782)),
		Fdy:   float32(-libc.Float64FromFloat64(3.141369404602301e-07)),
		Fe:    int32(FE_INEXACT),
	},
	855: {
		Ffile: __ccgo_ts,
		Fline: int32(872),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.019102183945165174)),
		Fy:    float64(-libc.Float64FromFloat64(0.01910334567660178)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	856: {
		Ffile: __ccgo_ts,
		Fline: int32(873),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.019102183945165174)),
		Fy:    float64(-libc.Float64FromFloat64(0.01910334567660178)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	857: {
		Ffile: __ccgo_ts,
		Fline: int32(874),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.026549101143415906),
		Fy:    libc.Float64FromFloat64(0.026552220130105605),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	858: {
		Ffile: __ccgo_ts,
		Fline: int32(875),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.026549101143415906),
		Fy:    libc.Float64FromFloat64(0.026552220130105608),
		Fdy:   float32(2.916452501722233e-07),
		Fe:    int32(FE_INEXACT),
	},
	859: {
		Ffile: __ccgo_ts,
		Fline: int32(876),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.026549101143415906),
		Fy:    libc.Float64FromFloat64(0.026552220130105605),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	860: {
		Ffile: __ccgo_ts,
		Fline: int32(877),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(18.256433505473886)),
		Fy:    float64(-libc.Float64FromFloat64(4.242661030662137e+07)),
		Fdy:   float32(-libc.Float64FromFloat64(2.0110341836243606e-07)),
		Fe:    int32(FE_INEXACT),
	},
	861: {
		Ffile: __ccgo_ts,
		Fline: int32(878),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(18.256433505473886)),
		Fy:    float64(-libc.Float64FromFloat64(4.2426610306621365e+07)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	862: {
		Ffile: __ccgo_ts,
		Fline: int32(879),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(18.256433505473886)),
		Fy:    float64(-libc.Float64FromFloat64(4.2426610306621365e+07)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	863: {
		Ffile: __ccgo_ts,
		Fline: int32(880),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0034012998847493606)),
		Fy:    float64(-libc.Float64FromFloat64(0.003401306442936028)),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	864: {
		Ffile: __ccgo_ts,
		Fline: int32(881),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0034012998847493606)),
		Fy:    float64(-libc.Float64FromFloat64(0.0034013064429360275)),
		Fdy:   float32(5.499677513398638e-07),
		Fe:    int32(FE_INEXACT),
	},
	865: {
		Ffile: __ccgo_ts,
		Fline: int32(882),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0034012998847493606)),
		Fy:    float64(-libc.Float64FromFloat64(0.0034013064429360275)),
		Fdy:   float32(5.499677513398638e-07),
		Fe:    int32(FE_INEXACT),
	},
	866: {
		Ffile: __ccgo_ts,
		Fline: int32(883),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1277229918755359),
		Fy:    libc.Float64FromFloat64(0.1280705362220061),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	867: {
		Ffile: __ccgo_ts,
		Fline: int32(884),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01066678182233932)),
		Fy:    float64(-libc.Float64FromFloat64(0.010666984101646175)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	868: {
		Ffile: __ccgo_ts,
		Fline: int32(885),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.01066678182233932)),
		Fy:    float64(-libc.Float64FromFloat64(0.010666984101646173)),
		Fdy:   float32(8.598089493716543e-07),
		Fe:    int32(FE_INEXACT),
	},
	869: {
		Ffile: __ccgo_ts,
		Fline: int32(886),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.01066678182233932)),
		Fy:    float64(-libc.Float64FromFloat64(0.010666984101646173)),
		Fdy:   float32(8.598089493716543e-07),
		Fe:    int32(FE_INEXACT),
	},
	870: {
		Ffile: __ccgo_ts,
		Fline: int32(887),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(256.1677608586269)),
		Fy:    float64(-libc.Float64FromFloat64(8.937478020049195e+110)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999994933605194)),
		Fe:    int32(FE_INEXACT),
	},
	871: {
		Ffile: __ccgo_ts,
		Fline: int32(888),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0022403293686252624)),
		Fy:    float64(-libc.Float64FromFloat64(0.0022403312426895077)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999940395355225)),
		Fe:    int32(FE_INEXACT),
	},
	872: {
		Ffile: __ccgo_ts,
		Fline: int32(889),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(606.8263056790327)),
		Fy:    float64(-libc.Float64FromFloat64(1.7389457539002224e+263)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	873: {
		Ffile: __ccgo_ts,
		Fline: int32(890),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.727700580911929),
		Fy:    libc.Float64FromFloat64(7.6161506810359745),
		Fdy:   float32(0.4999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	874: {
		Ffile: __ccgo_ts,
		Fline: int32(891),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.002311045285005879),
		Fy:    libc.Float64FromFloat64(0.002311047342195063),
		Fdy:   float32(-libc.Float64FromFloat64(1.0878680001269458e-07)),
		Fe:    int32(FE_INEXACT),
	},
	875: {
		Ffile: __ccgo_ts,
		Fline: int32(892),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.002311045285005879),
		Fy:    libc.Float64FromFloat64(0.0023110473421950636),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	876: {
		Ffile: __ccgo_ts,
		Fline: int32(893),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.002311045285005879),
		Fy:    libc.Float64FromFloat64(0.002311047342195063),
		Fdy:   float32(-libc.Float64FromFloat64(1.0878680001269458e-07)),
		Fe:    int32(FE_INEXACT),
	},
	877: {
		Ffile: __ccgo_ts,
		Fline: int32(894),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(7.259473878941525)),
		Fy:    float64(-libc.Float64FromFloat64(710.7538757082675)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	878: {
		Ffile: __ccgo_ts,
		Fline: int32(895),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(7.259473878941525)),
		Fy:    float64(-libc.Float64FromFloat64(710.7538757082674)),
		Fdy:   float32(9.291995866078651e-07),
		Fe:    int32(FE_INEXACT),
	},
	879: {
		Ffile: __ccgo_ts,
		Fline: int32(896),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(7.259473878941525)),
		Fy:    float64(-libc.Float64FromFloat64(710.7538757082674)),
		Fdy:   float32(9.291995866078651e-07),
		Fe:    int32(FE_INEXACT),
	},
	880: {
		Ffile: __ccgo_ts,
		Fline: int32(897),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.006530578744512849)),
		Fy:    float64(-libc.Float64FromFloat64(0.006530625164464857)),
		Fdy:   float32(-libc.Float64FromFloat64(7.069237994983268e-07)),
		Fe:    int32(FE_INEXACT),
	},
	881: {
		Ffile: __ccgo_ts,
		Fline: int32(898),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.006530578744512849)),
		Fy:    float64(-libc.Float64FromFloat64(0.006530625164464856)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	882: {
		Ffile: __ccgo_ts,
		Fline: int32(899),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.006530578744512849)),
		Fy:    float64(-libc.Float64FromFloat64(0.006530625164464856)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	883: {
		Ffile: __ccgo_ts,
		Fline: int32(900),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(107.28369194963955)),
		Fy:    float64(-libc.Float64FromFloat64(1.9574262867742917e+46)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	884: {
		Ffile: __ccgo_ts,
		Fline: int32(901),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(107.28369194963955)),
		Fy:    float64(-libc.Float64FromFloat64(1.9574262867742914e+46)),
		Fdy:   float32(6.472253630818159e-07),
		Fe:    int32(FE_INEXACT),
	},
	885: {
		Ffile: __ccgo_ts,
		Fline: int32(902),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(107.28369194963955)),
		Fy:    float64(-libc.Float64FromFloat64(1.9574262867742914e+46)),
		Fdy:   float32(6.472253630818159e-07),
		Fe:    int32(FE_INEXACT),
	},
	886: {
		Ffile: __ccgo_ts,
		Fline: int32(903),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(120.89186865308888),
		Fy:    libc.Float64FromFloat64(1.5908948184622513e+52),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	887: {
		Ffile: __ccgo_ts,
		Fline: int32(904),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(120.89186865308888),
		Fy:    libc.Float64FromFloat64(1.5908948184622515e+52),
		Fdy:   float32(5.467097707878565e-07),
		Fe:    int32(FE_INEXACT),
	},
	888: {
		Ffile: __ccgo_ts,
		Fline: int32(905),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(120.89186865308888),
		Fy:    libc.Float64FromFloat64(1.5908948184622513e+52),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	889: {
		Ffile: __ccgo_ts,
		Fline: int32(906),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.06038797537674587),
		Fy:    libc.Float64FromFloat64(0.06042468495071895),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	890: {
		Ffile: __ccgo_ts,
		Fline: int32(907),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(13.724339928077152)),
		Fy:    float64(-libc.Float64FromFloat64(456430.96803975827)),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	891: {
		Ffile: __ccgo_ts,
		Fline: int32(908),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(613.0651623378379),
		Fy:    libc.Float64FromFloat64(8.908140430923355e+265),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999943375587463)),
		Fe:    int32(FE_INEXACT),
	},
	892: {
		Ffile: __ccgo_ts,
		Fline: int32(909),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.3710172608390745),
		Fy:    libc.Float64FromFloat64(0.3795880283970969),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	893: {
		Ffile: __ccgo_ts,
		Fline: int32(910),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(9.385384233391369),
		Fy:    libc.Float64FromFloat64(5956.492295544711),
		Fdy:   float32(-libc.Float64FromFloat64(4.4369608076522127e-07)),
		Fe:    int32(FE_INEXACT),
	},
	894: {
		Ffile: __ccgo_ts,
		Fline: int32(911),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(9.385384233391369),
		Fy:    libc.Float64FromFloat64(5956.492295544712),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	895: {
		Ffile: __ccgo_ts,
		Fline: int32(912),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.385384233391369),
		Fy:    libc.Float64FromFloat64(5956.492295544711),
		Fdy:   float32(-libc.Float64FromFloat64(4.4369608076522127e-07)),
		Fe:    int32(FE_INEXACT),
	},
	896: {
		Ffile: __ccgo_ts,
		Fline: int32(913),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(237.16916500369416)),
		Fy:    float64(-libc.Float64FromFloat64(5.014523178052551e+102)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	897: {
		Ffile: __ccgo_ts,
		Fline: int32(914),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(237.16916500369416)),
		Fy:    float64(-libc.Float64FromFloat64(5.01452317805255e+102)),
		Fdy:   float32(1.4223735433915863e-07),
		Fe:    int32(FE_INEXACT),
	},
	898: {
		Ffile: __ccgo_ts,
		Fline: int32(915),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(237.16916500369416)),
		Fy:    float64(-libc.Float64FromFloat64(5.01452317805255e+102)),
		Fdy:   float32(1.4223735433915863e-07),
		Fe:    int32(FE_INEXACT),
	},
	899: {
		Ffile: __ccgo_ts,
		Fline: int32(916),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.029611638012924036),
		Fy:    libc.Float64FromFloat64(0.029615965692374627),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	900: {
		Ffile: __ccgo_ts,
		Fline: int32(917),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.029611638012924036),
		Fy:    libc.Float64FromFloat64(0.02961596569237463),
		Fdy:   float32(5.860066494278726e-07),
		Fe:    int32(FE_INEXACT),
	},
	901: {
		Ffile: __ccgo_ts,
		Fline: int32(918),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.029611638012924036),
		Fy:    libc.Float64FromFloat64(0.029615965692374627),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	902: {
		Ffile: __ccgo_ts,
		Fline: int32(919),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.002131275606595525)),
		Fy:    float64(-libc.Float64FromFloat64(0.0021312772200907746)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	903: {
		Ffile: __ccgo_ts,
		Fline: int32(920),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(138.64046594811924),
		Fy:    libc.Float64FromFloat64(8.123802078962541e+59),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	904: {
		Ffile: __ccgo_ts,
		Fline: int32(921),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(446.66327598259795),
		Fy:    libc.Float64FromFloat64(4.812447806666973e+193),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	905: {
		Ffile: __ccgo_ts,
		Fline: int32(922),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.1600721678917803)),
		Fy:    float64(-libc.Float64FromFloat64(0.16075663504444135)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	906: {
		Ffile: __ccgo_ts,
		Fline: int32(923),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01457445884947975),
		Fy:    libc.Float64FromFloat64(0.014574974826876676),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	907: {
		Ffile: __ccgo_ts,
		Fline: int32(924),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.633783995895476),
		Fy:    libc.Float64FromFloat64(6.927280704767608),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	908: {
		Ffile: __ccgo_ts,
		Fline: int32(925),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.633783995895476),
		Fy:    libc.Float64FromFloat64(6.927280704767609),
		Fdy:   float32(6.521615887322696e-07),
		Fe:    int32(FE_INEXACT),
	},
	909: {
		Ffile: __ccgo_ts,
		Fline: int32(926),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.633783995895476),
		Fy:    libc.Float64FromFloat64(6.927280704767608),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	910: {
		Ffile: __ccgo_ts,
		Fline: int32(927),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.3368397469312154),
		Fy:    libc.Float64FromFloat64(14.0472693828586),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	911: {
		Ffile: __ccgo_ts,
		Fline: int32(928),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.3368397469312154),
		Fy:    libc.Float64FromFloat64(14.047269382858602),
		Fdy:   float32(6.478624214878437e-08),
		Fe:    int32(FE_INEXACT),
	},
	912: {
		Ffile: __ccgo_ts,
		Fline: int32(929),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.3368397469312154),
		Fy:    libc.Float64FromFloat64(14.0472693828586),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	913: {
		Ffile: __ccgo_ts,
		Fline: int32(930),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.13046280031392768)),
		Fy:    float64(-libc.Float64FromFloat64(0.13083320666710666)),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	914: {
		Ffile: __ccgo_ts,
		Fline: int32(931),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4169775167373463)),
		Fy:    float64(-libc.Float64FromFloat64(0.4291663300541138)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	915: {
		Ffile: __ccgo_ts,
		Fline: int32(932),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.4169775167373463)),
		Fy:    float64(-libc.Float64FromFloat64(0.42916633005411375)),
		Fdy:   float32(2.807484236200253e-07),
		Fe:    int32(FE_INEXACT),
	},
	916: {
		Ffile: __ccgo_ts,
		Fline: int32(933),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.4169775167373463)),
		Fy:    float64(-libc.Float64FromFloat64(0.42916633005411375)),
		Fdy:   float32(2.807484236200253e-07),
		Fe:    int32(FE_INEXACT),
	},
	917: {
		Ffile: __ccgo_ts,
		Fline: int32(934),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.389751868726226)),
		Fy:    float64(-libc.Float64FromFloat64(14.812437545453937)),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	918: {
		Ffile: __ccgo_ts,
		Fline: int32(935),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.9848512468568043)),
		Fy:    float64(-libc.Float64FromFloat64(26.87934706592333)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	919: {
		Ffile: __ccgo_ts,
		Fline: int32(936),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.04492707130397345),
		Fy:    libc.Float64FromFloat64(0.0449421866086599),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	920: {
		Ffile: __ccgo_ts,
		Fline: int32(937),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.04492707130397345),
		Fy:    libc.Float64FromFloat64(0.04494218660865991),
		Fdy:   float32(7.656756224605488e-07),
		Fe:    int32(FE_INEXACT),
	},
	921: {
		Ffile: __ccgo_ts,
		Fline: int32(938),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.04492707130397345),
		Fy:    libc.Float64FromFloat64(0.0449421866086599),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	922: {
		Ffile: __ccgo_ts,
		Fline: int32(939),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.16790115008969986)),
		Fy:    float64(-libc.Float64FromFloat64(0.16869114064177698)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	923: {
		Ffile: __ccgo_ts,
		Fline: int32(940),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.002684762597519346),
		Fy:    libc.Float64FromFloat64(0.0026847658227930273),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	924: {
		Ffile: __ccgo_ts,
		Fline: int32(941),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.5740247310479),
		Fy:    libc.Float64FromFloat64(2646.1936151962673),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	925: {
		Ffile: __ccgo_ts,
		Fline: int32(942),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00346294903633166),
		Fy:    libc.Float64FromFloat64(0.0034629559576258677),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	926: {
		Ffile: __ccgo_ts,
		Fline: int32(943),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(91.54507710927231)),
		Fy:    float64(-libc.Float64FromFloat64(2.8608285965496126e+39)),
		Fdy:   float32(0.49999943375587463),
		Fe:    int32(FE_INEXACT),
	},
	927: {
		Ffile: __ccgo_ts,
		Fline: int32(944),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.3082895470332266)),
		Fy:    float64(-libc.Float64FromFloat64(0.3131962383265829)),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	928: {
		Ffile: __ccgo_ts,
		Fline: int32(945),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.744746433167145),
		Fy:    libc.Float64FromFloat64(424.79124803712824),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995529651642)),
		Fe:    int32(FE_INEXACT),
	},
	929: {
		Ffile: __ccgo_ts,
		Fline: int32(946),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.1039643949986134),
		Fy:    libc.Float64FromFloat64(0.10415178108637137),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	930: {
		Ffile: __ccgo_ts,
		Fline: int32(947),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23769741149063506)),
		Fy:    float64(-libc.Float64FromFloat64(0.23994206290916695)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	931: {
		Ffile: __ccgo_ts,
		Fline: int32(948),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.23769741149063506)),
		Fy:    float64(-libc.Float64FromFloat64(0.23994206290916692)),
		Fdy:   float32(9.107989740186895e-07),
		Fe:    int32(FE_INEXACT),
	},
	932: {
		Ffile: __ccgo_ts,
		Fline: int32(949),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.23769741149063506)),
		Fy:    float64(-libc.Float64FromFloat64(0.23994206290916692)),
		Fdy:   float32(9.107989740186895e-07),
		Fe:    int32(FE_INEXACT),
	},
	933: {
		Ffile: __ccgo_ts,
		Fline: int32(950),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(594.2089439850178),
		Fy:    libc.Float64FromFloat64(5.762825571230556e+257),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	934: {
		Ffile: __ccgo_ts,
		Fline: int32(951),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(27.245848986508715)),
		Fy:    float64(-libc.Float64FromFloat64(3.401667603555765e+11)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	935: {
		Ffile: __ccgo_ts,
		Fline: int32(952),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(81.12539124708067),
		Fy:    libc.Float64FromFloat64(8.536498437769605e+34),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	936: {
		Ffile: __ccgo_ts,
		Fline: int32(953),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(681.3473007604568)),
		Fy:    float64(-libc.Float64FromFloat64(4.0210825066931177e+295)),
		Fdy:   float32(-libc.Float64FromFloat64(3.616528090333304e-07)),
		Fe:    int32(FE_INEXACT),
	},
	937: {
		Ffile: __ccgo_ts,
		Fline: int32(954),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(681.3473007604568)),
		Fy:    float64(-libc.Float64FromFloat64(4.0210825066931172e+295)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	938: {
		Ffile: __ccgo_ts,
		Fline: int32(955),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(681.3473007604568)),
		Fy:    float64(-libc.Float64FromFloat64(4.0210825066931172e+295)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	939: {
		Ffile: __ccgo_ts,
		Fline: int32(956),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(614.6892325701142)),
		Fy:    float64(-libc.Float64FromFloat64(4.5197227485122005e+266)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	940: {
		Ffile: __ccgo_ts,
		Fline: int32(957),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(614.6892325701142)),
		Fy:    float64(-libc.Float64FromFloat64(4.5197227485122e+266)),
		Fdy:   float32(9.51432696183474e-07),
		Fe:    int32(FE_INEXACT),
	},
	941: {
		Ffile: __ccgo_ts,
		Fline: int32(958),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(614.6892325701142)),
		Fy:    float64(-libc.Float64FromFloat64(4.5197227485122e+266)),
		Fdy:   float32(9.51432696183474e-07),
		Fe:    int32(FE_INEXACT),
	},
	942: {
		Ffile: __ccgo_ts,
		Fline: int32(959),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.027973705735458543),
		Fy:    libc.Float64FromFloat64(0.02797735424720016),
		Fdy:   float32(-libc.Float64FromFloat64(5.492896093528543e-07)),
		Fe:    int32(FE_INEXACT),
	},
	943: {
		Ffile: __ccgo_ts,
		Fline: int32(960),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.027973705735458543),
		Fy:    libc.Float64FromFloat64(0.027977354247200163),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	944: {
		Ffile: __ccgo_ts,
		Fline: int32(961),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.027973705735458543),
		Fy:    libc.Float64FromFloat64(0.02797735424720016),
		Fdy:   float32(-libc.Float64FromFloat64(5.492896093528543e-07)),
		Fe:    int32(FE_INEXACT),
	},
	945: {
		Ffile: __ccgo_ts,
		Fline: int32(962),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(9.241786816621016)),
		Fy:    float64(-libc.Float64FromFloat64(5159.730490924843)),
		Fdy:   float32(0.49999961256980896),
		Fe:    int32(FE_INEXACT),
	},
	946: {
		Ffile: __ccgo_ts,
		Fline: int32(963),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(116.31819967014897)),
		Fy:    float64(-libc.Float64FromFloat64(1.6418076162535131e+50)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	947: {
		Ffile: __ccgo_ts,
		Fline: int32(964),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(116.31819967014897)),
		Fy:    float64(-libc.Float64FromFloat64(1.641807616253513e+50)),
		Fdy:   float32(1.6867879537585395e-07),
		Fe:    int32(FE_INEXACT),
	},
	948: {
		Ffile: __ccgo_ts,
		Fline: int32(965),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(116.31819967014897)),
		Fy:    float64(-libc.Float64FromFloat64(1.641807616253513e+50)),
		Fdy:   float32(1.6867879537585395e-07),
		Fe:    int32(FE_INEXACT),
	},
	949: {
		Ffile: __ccgo_ts,
		Fline: int32(966),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(45.967748305681525)),
		Fy:    float64(-libc.Float64FromFloat64(4.597369804038203e+19)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	950: {
		Ffile: __ccgo_ts,
		Fline: int32(967),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(126.94337499951817),
		Fy:    libc.Float64FromFloat64(6.75736348711597e+54),
		Fdy:   float32(-libc.Float64FromFloat64(3.1987124771148956e-07)),
		Fe:    int32(FE_INEXACT),
	},
	951: {
		Ffile: __ccgo_ts,
		Fline: int32(968),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(126.94337499951817),
		Fy:    libc.Float64FromFloat64(6.757363487115971e+54),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	952: {
		Ffile: __ccgo_ts,
		Fline: int32(969),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(126.94337499951817),
		Fy:    libc.Float64FromFloat64(6.75736348711597e+54),
		Fdy:   float32(-libc.Float64FromFloat64(3.1987124771148956e-07)),
		Fe:    int32(FE_INEXACT),
	},
	953: {
		Ffile: __ccgo_ts,
		Fline: int32(970),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(58.23587993961139)),
		Fy:    float64(-libc.Float64FromFloat64(9.783433804655248e+24)),
		Fdy:   float32(-libc.Float64FromFloat64(7.63364482736506e-07)),
		Fe:    int32(FE_INEXACT),
	},
	954: {
		Ffile: __ccgo_ts,
		Fline: int32(971),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(58.23587993961139)),
		Fy:    float64(-libc.Float64FromFloat64(9.783433804655246e+24)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	955: {
		Ffile: __ccgo_ts,
		Fline: int32(972),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(58.23587993961139)),
		Fy:    float64(-libc.Float64FromFloat64(9.783433804655246e+24)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	956: {
		Ffile: __ccgo_ts,
		Fline: int32(973),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.24379620977731725),
		Fy:    libc.Float64FromFloat64(0.24621846642159104),
		Fdy:   float32(-libc.Float64FromFloat64(2.665860279194021e-07)),
		Fe:    int32(FE_INEXACT),
	},
	957: {
		Ffile: __ccgo_ts,
		Fline: int32(974),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.24379620977731725),
		Fy:    libc.Float64FromFloat64(0.24621846642159106),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	958: {
		Ffile: __ccgo_ts,
		Fline: int32(975),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.24379620977731725),
		Fy:    libc.Float64FromFloat64(0.24621846642159104),
		Fdy:   float32(-libc.Float64FromFloat64(2.665860279194021e-07)),
		Fe:    int32(FE_INEXACT),
	},
	959: {
		Ffile: __ccgo_ts,
		Fline: int32(976),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.8978304458486255),
		Fy:    libc.Float64FromFloat64(66.99564189989096),
		Fdy:   float32(0.4999999701976776),
		Fe:    int32(FE_INEXACT),
	},
	960: {
		Ffile: __ccgo_ts,
		Fline: int32(977),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(18.62130945707799)),
		Fy:    float64(-libc.Float64FromFloat64(6.1108545577732936e+07)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	961: {
		Ffile: __ccgo_ts,
		Fline: int32(978),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.016545168006228825),
		Fy:    libc.Float64FromFloat64(0.016545922869402093),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	962: {
		Ffile: __ccgo_ts,
		Fline: int32(979),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.09933301031598128)),
		Fy:    float64(-libc.Float64FromFloat64(0.09949644483850721)),
		Fdy:   float32(-libc.Float64FromFloat64(2.7526917278919427e-07)),
		Fe:    int32(FE_INEXACT),
	},
	963: {
		Ffile: __ccgo_ts,
		Fline: int32(980),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.09933301031598128)),
		Fy:    float64(-libc.Float64FromFloat64(0.0994964448385072)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	964: {
		Ffile: __ccgo_ts,
		Fline: int32(981),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.09933301031598128)),
		Fy:    float64(-libc.Float64FromFloat64(0.0994964448385072)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	965: {
		Ffile: __ccgo_ts,
		Fline: int32(982),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(151.49225885853883),
		Fy:    libc.Float64FromFloat64(3.0990035779844554e+65),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	966: {
		Ffile: __ccgo_ts,
		Fline: int32(983),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.8358631714554887)),
		Fy:    float64(-libc.Float64FromFloat64(3.0555342888194605)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	967: {
		Ffile: __ccgo_ts,
		Fline: int32(984),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.8358631714554887)),
		Fy:    float64(-libc.Float64FromFloat64(3.05553428881946)),
		Fdy:   float32(8.17945476683235e-07),
		Fe:    int32(FE_INEXACT),
	},
	968: {
		Ffile: __ccgo_ts,
		Fline: int32(985),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.8358631714554887)),
		Fy:    float64(-libc.Float64FromFloat64(3.05553428881946)),
		Fdy:   float32(8.17945476683235e-07),
		Fe:    int32(FE_INEXACT),
	},
	969: {
		Ffile: __ccgo_ts,
		Fline: int32(986),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.014621991039588457)),
		Fy:    float64(-libc.Float64FromFloat64(0.014622512081828864)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	970: {
		Ffile: __ccgo_ts,
		Fline: int32(987),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.011702437297305943),
		Fy:    libc.Float64FromFloat64(0.011702704401490462),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	971: {
		Ffile: __ccgo_ts,
		Fline: int32(988),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.2273396893258596),
		Fy:    libc.Float64FromFloat64(0.2293030349387754),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	972: {
		Ffile: __ccgo_ts,
		Fline: int32(989),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.2273396893258596),
		Fy:    libc.Float64FromFloat64(0.22930303493877544),
		Fdy:   float32(8.520837013747951e-07),
		Fe:    int32(FE_INEXACT),
	},
	973: {
		Ffile: __ccgo_ts,
		Fline: int32(990),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.2273396893258596),
		Fy:    libc.Float64FromFloat64(0.2293030349387754),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	974: {
		Ffile: __ccgo_ts,
		Fline: int32(991),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.03815625137784059),
		Fy:    libc.Float64FromFloat64(0.03816551066318365),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	975: {
		Ffile: __ccgo_ts,
		Fline: int32(992),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.947257838275039),
		Fy:    libc.Float64FromFloat64(1.0954129763390839),
		Fdy:   float32(0.499999463558197),
		Fe:    int32(FE_INEXACT),
	},
	976: {
		Ffile: __ccgo_ts,
		Fline: int32(993),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(111.40716715586247),
		Fy:    libc.Float64FromFloat64(1.2091716175858173e+48),
		Fdy:   float32(-libc.Float64FromFloat64(7.188816653069807e-07)),
		Fe:    int32(FE_INEXACT),
	},
	977: {
		Ffile: __ccgo_ts,
		Fline: int32(994),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(111.40716715586247),
		Fy:    libc.Float64FromFloat64(1.2091716175858175e+48),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	978: {
		Ffile: __ccgo_ts,
		Fline: int32(995),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(111.40716715586247),
		Fy:    libc.Float64FromFloat64(1.2091716175858173e+48),
		Fdy:   float32(-libc.Float64FromFloat64(7.188816653069807e-07)),
		Fe:    int32(FE_INEXACT),
	},
	979: {
		Ffile: __ccgo_ts,
		Fline: int32(996),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004023236976895285),
		Fy:    libc.Float64FromFloat64(0.004023247830548555),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	980: {
		Ffile: __ccgo_ts,
		Fline: int32(997),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004023236976895285),
		Fy:    libc.Float64FromFloat64(0.004023247830548556),
		Fdy:   float32(7.394893941636838e-07),
		Fe:    int32(FE_INEXACT),
	},
	981: {
		Ffile: __ccgo_ts,
		Fline: int32(998),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004023236976895285),
		Fy:    libc.Float64FromFloat64(0.004023247830548555),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	982: {
		Ffile: __ccgo_ts,
		Fline: int32(999),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0012420330836808767)),
		Fy:    float64(-libc.Float64FromFloat64(0.0012420334030171668)),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	983: {
		Ffile: __ccgo_ts,
		Fline: int32(1000),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.23968812144064355),
		Fy:    libc.Float64FromFloat64(0.24198975254796568),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	984: {
		Ffile: __ccgo_ts,
		Fline: int32(1001),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.986527310372842)),
		Fy:    float64(-libc.Float64FromFloat64(3.576501146377928)),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	985: {
		Ffile: __ccgo_ts,
		Fline: int32(1002),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(628.3228024689478)),
		Fy:    float64(-libc.Float64FromFloat64(3.7678705235456825e+272)),
		Fdy:   float32(-libc.Float64FromFloat64(8.994917948257353e-07)),
		Fe:    int32(FE_INEXACT),
	},
	986: {
		Ffile: __ccgo_ts,
		Fline: int32(1003),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(628.3228024689478)),
		Fy:    float64(-libc.Float64FromFloat64(3.767870523545682e+272)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	987: {
		Ffile: __ccgo_ts,
		Fline: int32(1004),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(628.3228024689478)),
		Fy:    float64(-libc.Float64FromFloat64(3.767870523545682e+272)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	988: {
		Ffile: __ccgo_ts,
		Fline: int32(1005),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.08737937799006414)),
		Fy:    float64(-libc.Float64FromFloat64(0.08749061297235114)),
		Fdy:   float32(0.4999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	989: {
		Ffile: __ccgo_ts,
		Fline: int32(1006),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(11.519486686605731),
		Fy:    libc.Float64FromFloat64(50329.13967522563),
		Fdy:   float32(-libc.Float64FromFloat64(2.2806352717452683e-07)),
		Fe:    int32(FE_INEXACT),
	},
	990: {
		Ffile: __ccgo_ts,
		Fline: int32(1007),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(11.519486686605731),
		Fy:    libc.Float64FromFloat64(50329.13967522564),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	991: {
		Ffile: __ccgo_ts,
		Fline: int32(1008),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.519486686605731),
		Fy:    libc.Float64FromFloat64(50329.13967522563),
		Fdy:   float32(-libc.Float64FromFloat64(2.2806352717452683e-07)),
		Fe:    int32(FE_INEXACT),
	},
	992: {
		Ffile: __ccgo_ts,
		Fline: int32(1009),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(244.6004554856081),
		Fy:    libc.Float64FromFloat64(8.464437308574825e+105),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	993: {
		Ffile: __ccgo_ts,
		Fline: int32(1010),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.002201745269241301),
		Fy:    libc.Float64FromFloat64(0.0022017470481353017),
		Fdy:   float32(-libc.Float64FromFloat64(5.8208019737548966e-08)),
		Fe:    int32(FE_INEXACT),
	},
	994: {
		Ffile: __ccgo_ts,
		Fline: int32(1011),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.002201745269241301),
		Fy:    libc.Float64FromFloat64(0.002201747048135302),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	995: {
		Ffile: __ccgo_ts,
		Fline: int32(1012),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.002201745269241301),
		Fy:    libc.Float64FromFloat64(0.0022017470481353017),
		Fdy:   float32(-libc.Float64FromFloat64(5.8208019737548966e-08)),
		Fe:    int32(FE_INEXACT),
	},
	996: {
		Ffile: __ccgo_ts,
		Fline: int32(1013),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0017633366537313112),
		Fy:    libc.Float64FromFloat64(0.0017633375675383991),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	997: {
		Ffile: __ccgo_ts,
		Fline: int32(1014),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0017633366537313112),
		Fy:    libc.Float64FromFloat64(0.0017633375675383993),
		Fdy:   float32(2.0911353715291625e-07),
		Fe:    int32(FE_INEXACT),
	},
	998: {
		Ffile: __ccgo_ts,
		Fline: int32(1015),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0017633366537313112),
		Fy:    libc.Float64FromFloat64(0.0017633375675383991),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	999: {
		Ffile: __ccgo_ts,
		Fline: int32(1016),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(131.7142645922723)),
		Fy:    float64(-libc.Float64FromFloat64(7.975323435673107e+56)),
		Fdy:   float32(-libc.Float64FromFloat64(5.742569086919502e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1000: {
		Ffile: __ccgo_ts,
		Fline: int32(1017),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(131.7142645922723)),
		Fy:    float64(-libc.Float64FromFloat64(7.975323435673105e+56)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1001: {
		Ffile: __ccgo_ts,
		Fline: int32(1018),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(131.7142645922723)),
		Fy:    float64(-libc.Float64FromFloat64(7.975323435673105e+56)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1002: {
		Ffile: __ccgo_ts,
		Fline: int32(1019),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3025098046123219)),
		Fy:    float64(-libc.Float64FromFloat64(0.30714485073104775)),
		Fdy:   float32(-libc.Float64FromFloat64(3.918377444733778e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1003: {
		Ffile: __ccgo_ts,
		Fline: int32(1020),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3025098046123219)),
		Fy:    float64(-libc.Float64FromFloat64(0.3071448507310477)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1004: {
		Ffile: __ccgo_ts,
		Fline: int32(1021),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.3025098046123219)),
		Fy:    float64(-libc.Float64FromFloat64(0.3071448507310477)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1005: {
		Ffile: __ccgo_ts,
		Fline: int32(1022),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(69.44776590896484),
		Fy:    libc.Float64FromFloat64(7.240215937522972e+29),
		Fdy:   float32(0.4999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	1006: {
		Ffile: __ccgo_ts,
		Fline: int32(1023),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0026916945636421114),
		Fy:    libc.Float64FromFloat64(0.002691697813963001),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	1007: {
		Ffile: __ccgo_ts,
		Fline: int32(1024),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.2764973408254564),
		Fy:    libc.Float64FromFloat64(0.2800339053202855),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1008: {
		Ffile: __ccgo_ts,
		Fline: int32(1025),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.6527188958434986),
		Fy:    libc.Float64FromFloat64(2.514813936167272),
		Fdy:   float32(-libc.Float64FromFloat64(8.657349894747313e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1009: {
		Ffile: __ccgo_ts,
		Fline: int32(1026),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.6527188958434986),
		Fy:    libc.Float64FromFloat64(2.5148139361672723),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1010: {
		Ffile: __ccgo_ts,
		Fline: int32(1027),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6527188958434986),
		Fy:    libc.Float64FromFloat64(2.514813936167272),
		Fdy:   float32(-libc.Float64FromFloat64(8.657349894747313e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1011: {
		Ffile: __ccgo_ts,
		Fline: int32(1028),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(253.92531998445384)),
		Fy:    float64(-libc.Float64FromFloat64(9.491509551248593e+109)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1012: {
		Ffile: __ccgo_ts,
		Fline: int32(1029),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(253.92531998445384)),
		Fy:    float64(-libc.Float64FromFloat64(9.491509551248591e+109)),
		Fdy:   float32(8.811180691736809e-07),
		Fe:    int32(FE_INEXACT),
	},
	1013: {
		Ffile: __ccgo_ts,
		Fline: int32(1030),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(253.92531998445384)),
		Fy:    float64(-libc.Float64FromFloat64(9.491509551248591e+109)),
		Fdy:   float32(8.811180691736809e-07),
		Fe:    int32(FE_INEXACT),
	},
	1014: {
		Ffile: __ccgo_ts,
		Fline: int32(1031),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.050387930664732024)),
		Fy:    float64(-libc.Float64FromFloat64(0.050409255390314975)),
		Fdy:   float32(-libc.Float64FromFloat64(4.4961660705666873e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1015: {
		Ffile: __ccgo_ts,
		Fline: int32(1032),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.050387930664732024)),
		Fy:    float64(-libc.Float64FromFloat64(0.05040925539031497)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1016: {
		Ffile: __ccgo_ts,
		Fline: int32(1033),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.050387930664732024)),
		Fy:    float64(-libc.Float64FromFloat64(0.05040925539031497)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1017: {
		Ffile: __ccgo_ts,
		Fline: int32(1034),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.17994661906206963)),
		Fy:    float64(-libc.Float64FromFloat64(0.18091932806655744)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1018: {
		Ffile: __ccgo_ts,
		Fline: int32(1035),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.001682989471349168)),
		Fy:    float64(-libc.Float64FromFloat64(0.0016829902658475342)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1019: {
		Ffile: __ccgo_ts,
		Fline: int32(1036),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.001682989471349168)),
		Fy:    float64(-libc.Float64FromFloat64(0.001682990265847534)),
		Fdy:   float32(2.0550619694859051e-07),
		Fe:    int32(FE_INEXACT),
	},
	1020: {
		Ffile: __ccgo_ts,
		Fline: int32(1037),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.001682989471349168)),
		Fy:    float64(-libc.Float64FromFloat64(0.001682990265847534)),
		Fdy:   float32(2.0550619694859051e-07),
		Fe:    int32(FE_INEXACT),
	},
	1021: {
		Ffile: __ccgo_ts,
		Fline: int32(1038),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.737651053249887)),
		Fy:    float64(-libc.Float64FromFloat64(0.8063910516430637)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	1022: {
		Ffile: __ccgo_ts,
		Fline: int32(1039),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.627976655931801)),
		Fy:    float64(-libc.Float64FromFloat64(6.886751211301123)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1023: {
		Ffile: __ccgo_ts,
		Fline: int32(1040),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(233.1665973442606)),
		Fy:    float64(-libc.Float64FromFloat64(9.160867361250651e+100)),
		Fdy:   float32(-libc.Float64FromFloat64(6.77465550324996e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1024: {
		Ffile: __ccgo_ts,
		Fline: int32(1041),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(233.1665973442606)),
		Fy:    float64(-libc.Float64FromFloat64(9.16086736125065e+100)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1025: {
		Ffile: __ccgo_ts,
		Fline: int32(1042),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(233.1665973442606)),
		Fy:    float64(-libc.Float64FromFloat64(9.16086736125065e+100)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1026: {
		Ffile: __ccgo_ts,
		Fline: int32(1043),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.001355616945986008)),
		Fy:    float64(-libc.Float64FromFloat64(0.0013556173611883142)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999967217445374)),
		Fe:    int32(FE_INEXACT),
	},
	1027: {
		Ffile: __ccgo_ts,
		Fline: int32(1044),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.12802449991679965),
		Fy:    libc.Float64FromFloat64(0.12837451270902614),
		Fdy:   float32(-libc.Float64FromFloat64(8.796150297030181e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1028: {
		Ffile: __ccgo_ts,
		Fline: int32(1045),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.12802449991679965),
		Fy:    libc.Float64FromFloat64(0.12837451270902617),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1029: {
		Ffile: __ccgo_ts,
		Fline: int32(1046),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.12802449991679965),
		Fy:    libc.Float64FromFloat64(0.12837451270902614),
		Fdy:   float32(-libc.Float64FromFloat64(8.796150297030181e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1030: {
		Ffile: __ccgo_ts,
		Fline: int32(1047),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.129850899372178),
		Fy:    libc.Float64FromFloat64(1.3860567728653153),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1031: {
		Ffile: __ccgo_ts,
		Fline: int32(1048),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(124.0248767807216),
		Fy:    libc.Float64FromFloat64(3.6499726392748856e+53),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	1032: {
		Ffile: __ccgo_ts,
		Fline: int32(1049),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(472.6364846521285),
		Fy:    libc.Float64FromFloat64(9.170377944889652e+204),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	1033: {
		Ffile: __ccgo_ts,
		Fline: int32(1050),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(557.7186561145703),
		Fy:    libc.Float64FromFloat64(8.186623341586585e+241),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1034: {
		Ffile: __ccgo_ts,
		Fline: int32(1051),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(557.7186561145703),
		Fy:    libc.Float64FromFloat64(8.186623341586586e+241),
		Fdy:   float32(1.4373428314229386e-07),
		Fe:    int32(FE_INEXACT),
	},
	1035: {
		Ffile: __ccgo_ts,
		Fline: int32(1052),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(557.7186561145703),
		Fy:    libc.Float64FromFloat64(8.186623341586585e+241),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1036: {
		Ffile: __ccgo_ts,
		Fline: int32(1053),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.03045304070995063)),
		Fy:    float64(-libc.Float64FromFloat64(0.03045774789071961)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1037: {
		Ffile: __ccgo_ts,
		Fline: int32(1054),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.517860893810108),
		Fy:    libc.Float64FromFloat64(338.5634543474631),
		Fdy:   float32(-libc.Float64FromFloat64(2.9717864435951924e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1038: {
		Ffile: __ccgo_ts,
		Fline: int32(1055),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.517860893810108),
		Fy:    libc.Float64FromFloat64(338.56345434746316),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1039: {
		Ffile: __ccgo_ts,
		Fline: int32(1056),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.517860893810108),
		Fy:    libc.Float64FromFloat64(338.5634543474631),
		Fdy:   float32(-libc.Float64FromFloat64(2.9717864435951924e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1040: {
		Ffile: __ccgo_ts,
		Fline: int32(1057),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(40.00118745120798)),
		Fy:    float64(-libc.Float64FromFloat64(1.178324706867687e+17)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1041: {
		Ffile: __ccgo_ts,
		Fline: int32(1058),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(40.00118745120798)),
		Fy:    float64(-libc.Float64FromFloat64(1.1783247068676869e+17)),
		Fdy:   float32(8.083821398940927e-07),
		Fe:    int32(FE_INEXACT),
	},
	1042: {
		Ffile: __ccgo_ts,
		Fline: int32(1059),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(40.00118745120798)),
		Fy:    float64(-libc.Float64FromFloat64(1.1783247068676869e+17)),
		Fdy:   float32(8.083821398940927e-07),
		Fe:    int32(FE_INEXACT),
	},
	1043: {
		Ffile: __ccgo_ts,
		Fline: int32(1060),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0046881636630708295),
		Fy:    libc.Float64FromFloat64(0.0046881808365196505),
		Fdy:   float32(0.49999967217445374),
		Fe:    int32(FE_INEXACT),
	},
	1044: {
		Ffile: __ccgo_ts,
		Fline: int32(1061),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(25.673745696909656)),
		Fy:    float64(-libc.Float64FromFloat64(7.06213621954716e+10)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1045: {
		Ffile: __ccgo_ts,
		Fline: int32(1062),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(54.15840400595781),
		Fy:    libc.Float64FromFloat64(1.65831111943113e+23),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	1046: {
		Ffile: __ccgo_ts,
		Fline: int32(1063),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.013791665758930143)),
		Fy:    float64(-libc.Float64FromFloat64(0.013792102981981072)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1047: {
		Ffile: __ccgo_ts,
		Fline: int32(1064),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.010987678148172452)),
		Fy:    float64(-libc.Float64FromFloat64(0.01098789923820309)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1048: {
		Ffile: __ccgo_ts,
		Fline: int32(1065),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.010987678148172452)),
		Fy:    float64(-libc.Float64FromFloat64(0.010987899238203088)),
		Fdy:   float32(8.232130426222284e-07),
		Fe:    int32(FE_INEXACT),
	},
	1049: {
		Ffile: __ccgo_ts,
		Fline: int32(1066),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.010987678148172452)),
		Fy:    float64(-libc.Float64FromFloat64(0.010987899238203088)),
		Fdy:   float32(8.232130426222284e-07),
		Fe:    int32(FE_INEXACT),
	},
	1050: {
		Ffile: __ccgo_ts,
		Fline: int32(1067),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7112652440401097)),
		Fy:    float64(-libc.Float64FromFloat64(0.7727719208871113)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1051: {
		Ffile: __ccgo_ts,
		Fline: int32(1068),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7112652440401097)),
		Fy:    float64(-libc.Float64FromFloat64(0.7727719208871112)),
		Fdy:   float32(6.7936696268589e-07),
		Fe:    int32(FE_INEXACT),
	},
	1052: {
		Ffile: __ccgo_ts,
		Fline: int32(1069),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.7112652440401097)),
		Fy:    float64(-libc.Float64FromFloat64(0.7727719208871112)),
		Fdy:   float32(6.7936696268589e-07),
		Fe:    int32(FE_INEXACT),
	},
	1053: {
		Ffile: __ccgo_ts,
		Fline: int32(1070),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(30.76728506527727)),
		Fy:    float64(-libc.Float64FromFloat64(1.1508855342014209e+13)),
		Fdy:   float32(0.49999967217445374),
		Fe:    int32(FE_INEXACT),
	},
	1054: {
		Ffile: __ccgo_ts,
		Fline: int32(1071),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0021938355618367597),
		Fy:    libc.Float64FromFloat64(0.002193837321627671),
		Fdy:   float32(-libc.Float64FromFloat64(6.392904197127791e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1055: {
		Ffile: __ccgo_ts,
		Fline: int32(1072),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0021938355618367597),
		Fy:    libc.Float64FromFloat64(0.0021938373216276712),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1056: {
		Ffile: __ccgo_ts,
		Fline: int32(1073),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0021938355618367597),
		Fy:    libc.Float64FromFloat64(0.002193837321627671),
		Fdy:   float32(-libc.Float64FromFloat64(6.392904197127791e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1057: {
		Ffile: __ccgo_ts,
		Fline: int32(1074),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.91759748211262)),
		Fy:    float64(-libc.Float64FromFloat64(74940.54175524475)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	1058: {
		Ffile: __ccgo_ts,
		Fline: int32(1075),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(313.0743529787835)),
		Fy:    float64(-libc.Float64FromFloat64(4.6284324359748774e+135)),
		Fdy:   float32(-libc.Float64FromFloat64(2.6741082592707244e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1059: {
		Ffile: __ccgo_ts,
		Fline: int32(1076),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(313.0743529787835)),
		Fy:    float64(-libc.Float64FromFloat64(4.628432435974877e+135)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	1060: {
		Ffile: __ccgo_ts,
		Fline: int32(1077),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(313.0743529787835)),
		Fy:    float64(-libc.Float64FromFloat64(4.628432435974877e+135)),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	1061: {
		Ffile: __ccgo_ts,
		Fline: int32(1078),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0033682020118461814)),
		Fy:    float64(-libc.Float64FromFloat64(0.003368208380437621)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1062: {
		Ffile: __ccgo_ts,
		Fline: int32(1079),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.004181737512941932)),
		Fy:    float64(-libc.Float64FromFloat64(0.004181749700576825)),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1063: {
		Ffile: __ccgo_ts,
		Fline: int32(1080),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.8097755092085603)),
		Fy:    float64(-libc.Float64FromFloat64(0.9012227323497158)),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	1064: {
		Ffile: __ccgo_ts,
		Fline: int32(1081),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0724889347308385)),
		Fy:    float64(-libc.Float64FromFloat64(0.0725524353565681)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1065: {
		Ffile: __ccgo_ts,
		Fline: int32(1082),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0724889347308385)),
		Fy:    float64(-libc.Float64FromFloat64(0.07255243535656809)),
		Fdy:   float32(9.043952786669252e-07),
		Fe:    int32(FE_INEXACT),
	},
	1066: {
		Ffile: __ccgo_ts,
		Fline: int32(1083),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0724889347308385)),
		Fy:    float64(-libc.Float64FromFloat64(0.07255243535656809)),
		Fdy:   float32(9.043952786669252e-07),
		Fe:    int32(FE_INEXACT),
	},
	1067: {
		Ffile: __ccgo_ts,
		Fline: int32(1084),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.033841981836225256)),
		Fy:    float64(-libc.Float64FromFloat64(0.0338484419621427)),
		Fdy:   float32(-libc.Float64FromFloat64(8.926180612434109e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1068: {
		Ffile: __ccgo_ts,
		Fline: int32(1085),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.033841981836225256)),
		Fy:    float64(-libc.Float64FromFloat64(0.03384844196214269)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1069: {
		Ffile: __ccgo_ts,
		Fline: int32(1086),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.033841981836225256)),
		Fy:    float64(-libc.Float64FromFloat64(0.03384844196214269)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1070: {
		Ffile: __ccgo_ts,
		Fline: int32(1087),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0322218888134604)),
		Fy:    float64(-libc.Float64FromFloat64(0.0322274648328986)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1071: {
		Ffile: __ccgo_ts,
		Fline: int32(1088),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0322218888134604)),
		Fy:    float64(-libc.Float64FromFloat64(0.032227464832898595)),
		Fdy:   float32(6.744437541783554e-07),
		Fe:    int32(FE_INEXACT),
	},
	1072: {
		Ffile: __ccgo_ts,
		Fline: int32(1089),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0322218888134604)),
		Fy:    float64(-libc.Float64FromFloat64(0.032227464832898595)),
		Fdy:   float32(6.744437541783554e-07),
		Fe:    int32(FE_INEXACT),
	},
	1073: {
		Ffile: __ccgo_ts,
		Fline: int32(1090),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.26591642533694415)),
		Fy:    float64(-libc.Float64FromFloat64(0.2690614176904455)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1074: {
		Ffile: __ccgo_ts,
		Fline: int32(1091),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.062154599891279146),
		Fy:    libc.Float64FromFloat64(0.06219462683796566),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	1075: {
		Ffile: __ccgo_ts,
		Fline: int32(1092),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.006833146726447951),
		Fy:    libc.Float64FromFloat64(0.006833199901999411),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1076: {
		Ffile: __ccgo_ts,
		Fline: int32(1093),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.006833146726447951),
		Fy:    libc.Float64FromFloat64(0.006833199901999412),
		Fdy:   float32(7.160041377574089e-07),
		Fe:    int32(FE_INEXACT),
	},
	1077: {
		Ffile: __ccgo_ts,
		Fline: int32(1094),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.006833146726447951),
		Fy:    libc.Float64FromFloat64(0.006833199901999411),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1078: {
		Ffile: __ccgo_ts,
		Fline: int32(1095),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(248.846625269489)),
		Fy:    float64(-libc.Float64FromFloat64(5.911343537937787e+107)),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	1079: {
		Ffile: __ccgo_ts,
		Fline: int32(1096),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(7.798738149208579)),
		Fy:    float64(-libc.Float64FromFloat64(1218.7619170318235)),
		Fdy:   float32(-libc.Float64FromFloat64(2.7811406511091263e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1080: {
		Ffile: __ccgo_ts,
		Fline: int32(1097),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(7.798738149208579)),
		Fy:    float64(-libc.Float64FromFloat64(1218.7619170318233)),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1081: {
		Ffile: __ccgo_ts,
		Fline: int32(1098),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(7.798738149208579)),
		Fy:    float64(-libc.Float64FromFloat64(1218.7619170318233)),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1082: {
		Ffile: __ccgo_ts,
		Fline: int32(1099),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(58.51156052336972)),
		Fy:    float64(-libc.Float64FromFloat64(1.2888959741658462e+25)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995529651642)),
		Fe:    int32(FE_INEXACT),
	},
	1083: {
		Ffile: __ccgo_ts,
		Fline: int32(1100),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3892343846885517)),
		Fy:    float64(-libc.Float64FromFloat64(0.39913749472372523)),
		Fdy:   float32(-libc.Float64FromFloat64(8.145935908032698e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1084: {
		Ffile: __ccgo_ts,
		Fline: int32(1101),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3892343846885517)),
		Fy:    float64(-libc.Float64FromFloat64(0.3991374947237252)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	1085: {
		Ffile: __ccgo_ts,
		Fline: int32(1102),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.3892343846885517)),
		Fy:    float64(-libc.Float64FromFloat64(0.3991374947237252)),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	1086: {
		Ffile: __ccgo_ts,
		Fline: int32(1103),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0033486343107471126)),
		Fy:    float64(-libc.Float64FromFloat64(0.003348640568986355)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1087: {
		Ffile: __ccgo_ts,
		Fline: int32(1104),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0033486343107471126)),
		Fy:    float64(-libc.Float64FromFloat64(0.0033486405689863545)),
		Fdy:   float32(7.340882461903675e-07),
		Fe:    int32(FE_INEXACT),
	},
	1088: {
		Ffile: __ccgo_ts,
		Fline: int32(1105),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0033486343107471126)),
		Fy:    float64(-libc.Float64FromFloat64(0.0033486405689863545)),
		Fdy:   float32(7.340882461903675e-07),
		Fe:    int32(FE_INEXACT),
	},
	1089: {
		Ffile: __ccgo_ts,
		Fline: int32(1106),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(48.71147908435351)),
		Fy:    float64(-libc.Float64FromFloat64(7.14655229503112e+20)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1090: {
		Ffile: __ccgo_ts,
		Fline: int32(1107),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4136608569227536),
		Fy:    libc.Float64FromFloat64(1.9338632237548012),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1091: {
		Ffile: __ccgo_ts,
		Fline: int32(1108),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.1205185782772855),
		Fy:    libc.Float64FromFloat64(0.12081054015222809),
		Fdy:   float32(-libc.Float64FromFloat64(7.140037610042782e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1092: {
		Ffile: __ccgo_ts,
		Fline: int32(1109),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.1205185782772855),
		Fy:    libc.Float64FromFloat64(0.1208105401522281),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1093: {
		Ffile: __ccgo_ts,
		Fline: int32(1110),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.1205185782772855),
		Fy:    libc.Float64FromFloat64(0.12081054015222809),
		Fdy:   float32(-libc.Float64FromFloat64(7.140037610042782e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1094: {
		Ffile: __ccgo_ts,
		Fline: int32(1111),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(570.6185358621017),
		Fy:    libc.Float64FromFloat64(3.276811061150192e+247),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999934434890747)),
		Fe:    int32(FE_INEXACT),
	},
	1095: {
		Ffile: __ccgo_ts,
		Fline: int32(1112),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.0174955577806162)),
		Fy:    float64(-libc.Float64FromFloat64(1.2023794932702334)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	1096: {
		Ffile: __ccgo_ts,
		Fline: int32(1113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0031350820865199214),
		Fy:    libc.Float64FromFloat64(0.0031350872221734007),
		Fdy:   float32(-libc.Float64FromFloat64(2.9646128041349584e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1097: {
		Ffile: __ccgo_ts,
		Fline: int32(1114),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0031350820865199214),
		Fy:    libc.Float64FromFloat64(0.003135087222173401),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1098: {
		Ffile: __ccgo_ts,
		Fline: int32(1115),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0031350820865199214),
		Fy:    libc.Float64FromFloat64(0.0031350872221734007),
		Fdy:   float32(-libc.Float64FromFloat64(2.9646128041349584e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1099: {
		Ffile: __ccgo_ts,
		Fline: int32(1116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004704688201112392),
		Fy:    libc.Float64FromFloat64(0.004704705556797782),
		Fdy:   float32(-libc.Float64FromFloat64(4.4622655082093843e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1100: {
		Ffile: __ccgo_ts,
		Fline: int32(1117),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004704688201112392),
		Fy:    libc.Float64FromFloat64(0.004704705556797783),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1101: {
		Ffile: __ccgo_ts,
		Fline: int32(1118),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004704688201112392),
		Fy:    libc.Float64FromFloat64(0.004704705556797782),
		Fdy:   float32(-libc.Float64FromFloat64(4.4622655082093843e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1102: {
		Ffile: __ccgo_ts,
		Fline: int32(1119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(163.24623215553783),
		Fy:    libc.Float64FromFloat64(3.943735873857072e+70),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1103: {
		Ffile: __ccgo_ts,
		Fline: int32(1120),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(163.24623215553783),
		Fy:    libc.Float64FromFloat64(3.9437358738570724e+70),
		Fdy:   float32(3.3955785738726263e-07),
		Fe:    int32(FE_INEXACT),
	},
	1104: {
		Ffile: __ccgo_ts,
		Fline: int32(1121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(163.24623215553783),
		Fy:    libc.Float64FromFloat64(3.943735873857072e+70),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1105: {
		Ffile: __ccgo_ts,
		Fline: int32(1122),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.485465059394562)),
		Fy:    float64(-libc.Float64FromFloat64(2.0953110416520633)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	1106: {
		Ffile: __ccgo_ts,
		Fline: int32(1123),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.10684225942575322),
		Fy:    libc.Float64FromFloat64(0.10704564765556876),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	1107: {
		Ffile: __ccgo_ts,
		Fline: int32(1124),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(53.05229893378896)),
		Fy:    float64(-libc.Float64FromFloat64(5.486440624931597e+22)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999916553497314)),
		Fe:    int32(FE_INEXACT),
	},
	1108: {
		Ffile: __ccgo_ts,
		Fline: int32(1125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.08971628272820893),
		Fy:    libc.Float64FromFloat64(0.08983668573779514),
		Fdy:   float32(-libc.Float64FromFloat64(3.125140324300446e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1109: {
		Ffile: __ccgo_ts,
		Fline: int32(1126),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.08971628272820893),
		Fy:    libc.Float64FromFloat64(0.08983668573779516),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1110: {
		Ffile: __ccgo_ts,
		Fline: int32(1127),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.08971628272820893),
		Fy:    libc.Float64FromFloat64(0.08983668573779514),
		Fdy:   float32(-libc.Float64FromFloat64(3.125140324300446e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1111: {
		Ffile: __ccgo_ts,
		Fline: int32(1128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(12.419900951388778)),
		Fy:    float64(-libc.Float64FromFloat64(123841.00087755952)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	1112: {
		Ffile: __ccgo_ts,
		Fline: int32(1129),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(12.419900951388778)),
		Fy:    float64(-libc.Float64FromFloat64(123841.0008775595)),
		Fdy:   float32(2.852016507404187e-07),
		Fe:    int32(FE_INEXACT),
	},
	1113: {
		Ffile: __ccgo_ts,
		Fline: int32(1130),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(12.419900951388778)),
		Fy:    float64(-libc.Float64FromFloat64(123841.0008775595)),
		Fdy:   float32(2.852016507404187e-07),
		Fe:    int32(FE_INEXACT),
	},
	1114: {
		Ffile: __ccgo_ts,
		Fline: int32(1131),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.019683005202205804)),
		Fy:    float64(-libc.Float64FromFloat64(0.019684276160747225)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1115: {
		Ffile: __ccgo_ts,
		Fline: int32(1132),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.019683005202205804)),
		Fy:    float64(-libc.Float64FromFloat64(0.01968427616074722)),
		Fdy:   float32(7.220196494017728e-07),
		Fe:    int32(FE_INEXACT),
	},
	1116: {
		Ffile: __ccgo_ts,
		Fline: int32(1133),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.019683005202205804)),
		Fy:    float64(-libc.Float64FromFloat64(0.01968427616074722)),
		Fdy:   float32(7.220196494017728e-07),
		Fe:    int32(FE_INEXACT),
	},
	1117: {
		Ffile: __ccgo_ts,
		Fline: int32(1134),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0034705633803324247)),
		Fy:    float64(-libc.Float64FromFloat64(0.003470570347382808)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999940395355225)),
		Fe:    int32(FE_INEXACT),
	},
	1118: {
		Ffile: __ccgo_ts,
		Fline: int32(1135),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.010196301976728224),
		Fy:    libc.Float64FromFloat64(0.01019647865334519),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1119: {
		Ffile: __ccgo_ts,
		Fline: int32(1136),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.010196301976728224),
		Fy:    libc.Float64FromFloat64(0.010196478653345192),
		Fdy:   float32(7.071035952321836e-07),
		Fe:    int32(FE_INEXACT),
	},
	1120: {
		Ffile: __ccgo_ts,
		Fline: int32(1137),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.010196301976728224),
		Fy:    libc.Float64FromFloat64(0.01019647865334519),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1121: {
		Ffile: __ccgo_ts,
		Fline: int32(1138),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.835943360251591),
		Fy:    libc.Float64FromFloat64(62.97470557020158),
		Fdy:   float32(-libc.Float64FromFloat64(2.8501884230536234e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1122: {
		Ffile: __ccgo_ts,
		Fline: int32(1139),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.835943360251591),
		Fy:    libc.Float64FromFloat64(62.974705570201586),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1123: {
		Ffile: __ccgo_ts,
		Fline: int32(1140),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.835943360251591),
		Fy:    libc.Float64FromFloat64(62.97470557020158),
		Fdy:   float32(-libc.Float64FromFloat64(2.8501884230536234e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1124: {
		Ffile: __ccgo_ts,
		Fline: int32(1141),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.202901948388685),
		Fy:    libc.Float64FromFloat64(4.4703797906888205),
		Fdy:   float32(0.4999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1125: {
		Ffile: __ccgo_ts,
		Fline: int32(1142),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.02694509728196702)),
		Fy:    float64(-libc.Float64FromFloat64(0.02694835792895735)),
		Fdy:   float32(0.49999940395355225),
		Fe:    int32(FE_INEXACT),
	},
	1126: {
		Ffile: __ccgo_ts,
		Fline: int32(1143),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.008406090060550963),
		Fy:    libc.Float64FromFloat64(0.008406189059913884),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1127: {
		Ffile: __ccgo_ts,
		Fline: int32(1144),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.008406090060550963),
		Fy:    libc.Float64FromFloat64(0.008406189059913886),
		Fdy:   float32(1.9526416394910484e-07),
		Fe:    int32(FE_INEXACT),
	},
	1128: {
		Ffile: __ccgo_ts,
		Fline: int32(1145),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.008406090060550963),
		Fy:    libc.Float64FromFloat64(0.008406189059913884),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1129: {
		Ffile: __ccgo_ts,
		Fline: int32(1146),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(53.24801878439158)),
		Fy:    float64(-libc.Float64FromFloat64(6.672533066763922e+22)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	1130: {
		Ffile: __ccgo_ts,
		Fline: int32(1147),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(455.48742700464254)),
		Fy:    float64(-libc.Float64FromFloat64(3.270740802162682e+197)),
		Fdy:   float32(-libc.Float64FromFloat64(3.6132044556325127e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1131: {
		Ffile: __ccgo_ts,
		Fline: int32(1148),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(455.48742700464254)),
		Fy:    float64(-libc.Float64FromFloat64(3.270740802162681e+197)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1132: {
		Ffile: __ccgo_ts,
		Fline: int32(1149),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(455.48742700464254)),
		Fy:    float64(-libc.Float64FromFloat64(3.270740802162681e+197)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1133: {
		Ffile: __ccgo_ts,
		Fline: int32(1150),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(57.61641034307627),
		Fy:    libc.Float64FromFloat64(5.265736019851954e+24),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	1134: {
		Ffile: __ccgo_ts,
		Fline: int32(1151),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.030840753075434484)),
		Fy:    float64(-libc.Float64FromFloat64(0.030845642348870692)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	1135: {
		Ffile: __ccgo_ts,
		Fline: int32(1152),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(191.69758040086367)),
		Fy:    float64(-libc.Float64FromFloat64(8.957181338978069e+82)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1136: {
		Ffile: __ccgo_ts,
		Fline: int32(1153),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(191.69758040086367)),
		Fy:    float64(-libc.Float64FromFloat64(8.957181338978067e+82)),
		Fdy:   float32(2.0362674035823147e-07),
		Fe:    int32(FE_INEXACT),
	},
	1137: {
		Ffile: __ccgo_ts,
		Fline: int32(1154),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(191.69758040086367)),
		Fy:    float64(-libc.Float64FromFloat64(8.957181338978067e+82)),
		Fdy:   float32(2.0362674035823147e-07),
		Fe:    int32(FE_INEXACT),
	},
	1138: {
		Ffile: __ccgo_ts,
		Fline: int32(1155),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.4858565754905677),
		Fy:    libc.Float64FromFloat64(5.964075156927052),
		Fdy:   float32(-libc.Float64FromFloat64(2.089760755552561e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1139: {
		Ffile: __ccgo_ts,
		Fline: int32(1156),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.4858565754905677),
		Fy:    libc.Float64FromFloat64(5.964075156927053),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	1140: {
		Ffile: __ccgo_ts,
		Fline: int32(1157),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.4858565754905677),
		Fy:    libc.Float64FromFloat64(5.964075156927052),
		Fdy:   float32(-libc.Float64FromFloat64(2.089760755552561e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1141: {
		Ffile: __ccgo_ts,
		Fline: int32(1158),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(26.083025135540524)),
		Fy:    float64(-libc.Float64FromFloat64(1.063368747088575e+11)),
		Fdy:   float32(-libc.Float64FromFloat64(3.572953062302986e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1142: {
		Ffile: __ccgo_ts,
		Fline: int32(1159),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(26.083025135540524)),
		Fy:    float64(-libc.Float64FromFloat64(1.0633687470885748e+11)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1143: {
		Ffile: __ccgo_ts,
		Fline: int32(1160),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(26.083025135540524)),
		Fy:    float64(-libc.Float64FromFloat64(1.0633687470885748e+11)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1144: {
		Ffile: __ccgo_ts,
		Fline: int32(1161),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.975546843878052)),
		Fy:    float64(-libc.Float64FromFloat64(26.63023766728843)),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	1145: {
		Ffile: __ccgo_ts,
		Fline: int32(1162),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(64.07390027888627)),
		Fy:    float64(-libc.Float64FromFloat64(3.3566907321219466e+27)),
		Fdy:   float32(-libc.Float64FromFloat64(5.775034992439032e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1146: {
		Ffile: __ccgo_ts,
		Fline: int32(1163),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(64.07390027888627)),
		Fy:    float64(-libc.Float64FromFloat64(3.356690732121946e+27)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1147: {
		Ffile: __ccgo_ts,
		Fline: int32(1164),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(64.07390027888627)),
		Fy:    float64(-libc.Float64FromFloat64(3.356690732121946e+27)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1148: {
		Ffile: __ccgo_ts,
		Fline: int32(1165),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.041250813541761536)),
		Fy:    float64(-libc.Float64FromFloat64(0.041262513471511754)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1149: {
		Ffile: __ccgo_ts,
		Fline: int32(1166),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(166.26250423743244),
		Fy:    libc.Float64FromFloat64(8.051154153790877e+71),
		Fdy:   float32(0.49999934434890747),
		Fe:    int32(FE_INEXACT),
	},
	1150: {
		Ffile: __ccgo_ts,
		Fline: int32(1167),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(18.88326736432048),
		Fy:    libc.Float64FromFloat64(7.940883434036224e+07),
		Fdy:   float32(-libc.Float64FromFloat64(2.9353469699344714e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1151: {
		Ffile: __ccgo_ts,
		Fline: int32(1168),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(18.88326736432048),
		Fy:    libc.Float64FromFloat64(7.940883434036225e+07),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1152: {
		Ffile: __ccgo_ts,
		Fline: int32(1169),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(18.88326736432048),
		Fy:    libc.Float64FromFloat64(7.940883434036224e+07),
		Fdy:   float32(-libc.Float64FromFloat64(2.9353469699344714e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1153: {
		Ffile: __ccgo_ts,
		Fline: int32(1170),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.8731939990896205),
		Fy:    libc.Float64FromFloat64(65.36508917095715),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1154: {
		Ffile: __ccgo_ts,
		Fline: int32(1171),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.8731939990896205),
		Fy:    libc.Float64FromFloat64(65.36508917095716),
		Fdy:   float32(6.027840981914778e-07),
		Fe:    int32(FE_INEXACT),
	},
	1155: {
		Ffile: __ccgo_ts,
		Fline: int32(1172),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.8731939990896205),
		Fy:    libc.Float64FromFloat64(65.36508917095715),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1156: {
		Ffile: __ccgo_ts,
		Fline: int32(1173),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.16603573041936495)),
		Fy:    float64(-libc.Float64FromFloat64(0.16679965771798858)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1157: {
		Ffile: __ccgo_ts,
		Fline: int32(1174),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.16603573041936495)),
		Fy:    float64(-libc.Float64FromFloat64(0.16679965771798855)),
		Fdy:   float32(1.4539348569542199e-07),
		Fe:    int32(FE_INEXACT),
	},
	1158: {
		Ffile: __ccgo_ts,
		Fline: int32(1175),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.16603573041936495)),
		Fy:    float64(-libc.Float64FromFloat64(0.16679965771798855)),
		Fdy:   float32(1.4539348569542199e-07),
		Fe:    int32(FE_INEXACT),
	},
	1159: {
		Ffile: __ccgo_ts,
		Fline: int32(1176),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0035530258419969735),
		Fy:    libc.Float64FromFloat64(0.003553033317563701),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1160: {
		Ffile: __ccgo_ts,
		Fline: int32(1177),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0035530258419969735),
		Fy:    libc.Float64FromFloat64(0.0035530333175637016),
		Fdy:   float32(8.566917699681653e-07),
		Fe:    int32(FE_INEXACT),
	},
	1161: {
		Ffile: __ccgo_ts,
		Fline: int32(1178),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0035530258419969735),
		Fy:    libc.Float64FromFloat64(0.003553033317563701),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1162: {
		Ffile: __ccgo_ts,
		Fline: int32(1179),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.007031907391227668),
		Fy:    libc.Float64FromFloat64(0.0070319653433374),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	1163: {
		Ffile: __ccgo_ts,
		Fline: int32(1180),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(14.943862477632496),
		Fy:    libc.Float64FromFloat64(1.5452794056933569e+06),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1164: {
		Ffile: __ccgo_ts,
		Fline: int32(1181),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(14.943862477632496),
		Fy:    libc.Float64FromFloat64(1.545279405693357e+06),
		Fdy:   float32(4.117141259030177e-07),
		Fe:    int32(FE_INEXACT),
	},
	1165: {
		Ffile: __ccgo_ts,
		Fline: int32(1182),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(14.943862477632496),
		Fy:    libc.Float64FromFloat64(1.5452794056933569e+06),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1166: {
		Ffile: __ccgo_ts,
		Fline: int32(1183),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.002321292893131219)),
		Fy:    float64(-libc.Float64FromFloat64(0.0023212949778078206)),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	1167: {
		Ffile: __ccgo_ts,
		Fline: int32(1184),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(20.279025428844758)),
		Fy:    float64(-libc.Float64FromFloat64(3.206556129394113e+08)),
		Fdy:   float32(-libc.Float64FromFloat64(1.322384122204312e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1168: {
		Ffile: __ccgo_ts,
		Fline: int32(1185),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(20.279025428844758)),
		Fy:    float64(-libc.Float64FromFloat64(3.206556129394112e+08)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1169: {
		Ffile: __ccgo_ts,
		Fline: int32(1186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(20.279025428844758)),
		Fy:    float64(-libc.Float64FromFloat64(3.206556129394112e+08)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1170: {
		Ffile: __ccgo_ts,
		Fline: int32(1187),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.04293732805436054)),
		Fy:    float64(-libc.Float64FromFloat64(0.04295052258144166)),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	1171: {
		Ffile: __ccgo_ts,
		Fline: int32(1188),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(85.6474060122837)),
		Fy:    float64(-libc.Float64FromFloat64(7.85535515009925e+36)),
		Fdy:   float32(-libc.Float64FromFloat64(7.802340746820846e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1172: {
		Ffile: __ccgo_ts,
		Fline: int32(1189),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(85.6474060122837)),
		Fy:    float64(-libc.Float64FromFloat64(7.855355150099249e+36)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1173: {
		Ffile: __ccgo_ts,
		Fline: int32(1190),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(85.6474060122837)),
		Fy:    float64(-libc.Float64FromFloat64(7.855355150099249e+36)),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1174: {
		Ffile: __ccgo_ts,
		Fline: int32(1191),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.07363588227773775)),
		Fy:    float64(-libc.Float64FromFloat64(0.07370244559778795)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	1175: {
		Ffile: __ccgo_ts,
		Fline: int32(1192),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.131883584865253)),
		Fy:    float64(-libc.Float64FromFloat64(0.13226623217384192)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1176: {
		Ffile: __ccgo_ts,
		Fline: int32(1193),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.524753848006764),
		Fy:    libc.Float64FromFloat64(16.957978898935156),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1177: {
		Ffile: __ccgo_ts,
		Fline: int32(1194),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.007519402208472548),
		Fy:    libc.Float64FromFloat64(0.007519473068272875),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	1178: {
		Ffile: __ccgo_ts,
		Fline: int32(1195),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(104.31329652907378),
		Fy:    libc.Float64FromFloat64(1.0038274699080106e+45),
		Fdy:   float32(-libc.Float64FromFloat64(1.1904261043582665e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1179: {
		Ffile: __ccgo_ts,
		Fline: int32(1196),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(104.31329652907378),
		Fy:    libc.Float64FromFloat64(1.0038274699080108e+45),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1180: {
		Ffile: __ccgo_ts,
		Fline: int32(1197),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(104.31329652907378),
		Fy:    libc.Float64FromFloat64(1.0038274699080106e+45),
		Fdy:   float32(-libc.Float64FromFloat64(1.1904261043582665e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1181: {
		Ffile: __ccgo_ts,
		Fline: int32(1198),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(47.34866407991073)),
		Fy:    float64(-libc.Float64FromFloat64(1.829083532449287e+20)),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	1182: {
		Ffile: __ccgo_ts,
		Fline: int32(1199),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(452.8928564338384)),
		Fy:    float64(-libc.Float64FromFloat64(2.442521790887097e+196)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1183: {
		Ffile: __ccgo_ts,
		Fline: int32(1200),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(452.8928564338384)),
		Fy:    float64(-libc.Float64FromFloat64(2.4425217908870967e+196)),
		Fdy:   float32(4.686717716140265e-07),
		Fe:    int32(FE_INEXACT),
	},
	1184: {
		Ffile: __ccgo_ts,
		Fline: int32(1201),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(452.8928564338384)),
		Fy:    float64(-libc.Float64FromFloat64(2.4425217908870967e+196)),
		Fdy:   float32(4.686717716140265e-07),
		Fe:    int32(FE_INEXACT),
	},
	1185: {
		Ffile: __ccgo_ts,
		Fline: int32(1202),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.003317096021335891)),
		Fy:    float64(-libc.Float64FromFloat64(0.0033171021044101586)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	1186: {
		Ffile: __ccgo_ts,
		Fline: int32(1203),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.43428055887147)),
		Fy:    float64(-libc.Float64FromFloat64(2301.078790031394)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1187: {
		Ffile: __ccgo_ts,
		Fline: int32(1204),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.43428055887147)),
		Fy:    float64(-libc.Float64FromFloat64(2301.0787900313935)),
		Fdy:   float32(4.3234135205239e-07),
		Fe:    int32(FE_INEXACT),
	},
	1188: {
		Ffile: __ccgo_ts,
		Fline: int32(1205),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(8.43428055887147)),
		Fy:    float64(-libc.Float64FromFloat64(2301.0787900313935)),
		Fdy:   float32(4.3234135205239e-07),
		Fe:    int32(FE_INEXACT),
	},
	1189: {
		Ffile: __ccgo_ts,
		Fline: int32(1206),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.34953549938307515)),
		Fy:    float64(-libc.Float64FromFloat64(0.3566965250512138)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	1190: {
		Ffile: __ccgo_ts,
		Fline: int32(1207),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.05901657814409288)),
		Fy:    float64(-libc.Float64FromFloat64(0.0590508428063665)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1191: {
		Ffile: __ccgo_ts,
		Fline: int32(1208),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(98.5002015152501),
		Fy:    libc.Float64FromFloat64(2.999604446895706e+42),
		Fdy:   float32(-libc.Float64FromFloat64(0.499999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1192: {
		Ffile: __ccgo_ts,
		Fline: int32(1209),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.0909164872831094)),
		Fy:    float64(-libc.Float64FromFloat64(3.984377226336847)),
		Fdy:   float32(-libc.Float64FromFloat64(2.970055845707975e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1193: {
		Ffile: __ccgo_ts,
		Fline: int32(1210),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.0909164872831094)),
		Fy:    float64(-libc.Float64FromFloat64(3.9843772263368464)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1194: {
		Ffile: __ccgo_ts,
		Fline: int32(1211),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.0909164872831094)),
		Fy:    float64(-libc.Float64FromFloat64(3.9843772263368464)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1195: {
		Ffile: __ccgo_ts,
		Fline: int32(1212),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0012689815890789163),
		Fy:    libc.Float64FromFloat64(0.001268981929655471),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1196: {
		Ffile: __ccgo_ts,
		Fline: int32(1213),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0012689815890789163),
		Fy:    libc.Float64FromFloat64(0.0012689819296554713),
		Fdy:   float32(7.497792466892861e-07),
		Fe:    int32(FE_INEXACT),
	},
	1197: {
		Ffile: __ccgo_ts,
		Fline: int32(1214),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0012689815890789163),
		Fy:    libc.Float64FromFloat64(0.001268981929655471),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1198: {
		Ffile: __ccgo_ts,
		Fline: int32(1215),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.06405642846040122),
		Fy:    libc.Float64FromFloat64(0.06410024378272079),
		Fdy:   float32(-libc.Float64FromFloat64(5.236617894865958e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1199: {
		Ffile: __ccgo_ts,
		Fline: int32(1216),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.06405642846040122),
		Fy:    libc.Float64FromFloat64(0.0641002437827208),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1200: {
		Ffile: __ccgo_ts,
		Fline: int32(1217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.06405642846040122),
		Fy:    libc.Float64FromFloat64(0.06410024378272079),
		Fdy:   float32(-libc.Float64FromFloat64(5.236617894865958e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1201: {
		Ffile: __ccgo_ts,
		Fline: int32(1218),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.3389228676364389),
		Fy:    libc.Float64FromFloat64(1.7764021208452274),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1202: {
		Ffile: __ccgo_ts,
		Fline: int32(1219),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.3389228676364389),
		Fy:    libc.Float64FromFloat64(1.7764021208452276),
		Fdy:   float32(1.866389993665507e-07),
		Fe:    int32(FE_INEXACT),
	},
	1203: {
		Ffile: __ccgo_ts,
		Fline: int32(1220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3389228676364389),
		Fy:    libc.Float64FromFloat64(1.7764021208452274),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1204: {
		Ffile: __ccgo_ts,
		Fline: int32(1221),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.11371625047122623),
		Fy:    libc.Float64FromFloat64(0.11396149376535598),
		Fdy:   float32(-libc.Float64FromFloat64(4.894682774647663e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1205: {
		Ffile: __ccgo_ts,
		Fline: int32(1222),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.11371625047122623),
		Fy:    libc.Float64FromFloat64(0.113961493765356),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1206: {
		Ffile: __ccgo_ts,
		Fline: int32(1223),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.11371625047122623),
		Fy:    libc.Float64FromFloat64(0.11396149376535598),
		Fdy:   float32(-libc.Float64FromFloat64(4.894682774647663e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1207: {
		Ffile: __ccgo_ts,
		Fline: int32(1224),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(311.3991820756531)),
		Fy:    float64(-libc.Float64FromFloat64(8.667951027818371e+134)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1208: {
		Ffile: __ccgo_ts,
		Fline: int32(1225),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.004435114313271822),
		Fy:    libc.Float64FromFloat64(0.004435128853245858),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1209: {
		Ffile: __ccgo_ts,
		Fline: int32(1226),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(17.99543599704226)),
		Fy:    float64(-libc.Float64FromFloat64(3.2680489828708388e+07)),
		Fdy:   float32(-libc.Float64FromFloat64(4.1990304566752457e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1210: {
		Ffile: __ccgo_ts,
		Fline: int32(1227),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(17.99543599704226)),
		Fy:    float64(-libc.Float64FromFloat64(3.2680489828708384e+07)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1211: {
		Ffile: __ccgo_ts,
		Fline: int32(1228),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(17.99543599704226)),
		Fy:    float64(-libc.Float64FromFloat64(3.2680489828708384e+07)),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1212: {
		Ffile: __ccgo_ts,
		Fline: int32(1229),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(51.90434041593458),
		Fy:    libc.Float64FromFloat64(1.740758938191233e+22),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1213: {
		Ffile: __ccgo_ts,
		Fline: int32(1230),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(170.40718212848634),
		Fy:    libc.Float64FromFloat64(5.080060441034557e+73),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1214: {
		Ffile: __ccgo_ts,
		Fline: int32(1231),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(170.40718212848634),
		Fy:    libc.Float64FromFloat64(5.080060441034558e+73),
		Fdy:   float32(5.809894219055423e-07),
		Fe:    int32(FE_INEXACT),
	},
	1215: {
		Ffile: __ccgo_ts,
		Fline: int32(1232),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(170.40718212848634),
		Fy:    libc.Float64FromFloat64(5.080060441034557e+73),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1216: {
		Ffile: __ccgo_ts,
		Fline: int32(1233),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0015639804223160476)),
		Fy:    float64(-libc.Float64FromFloat64(0.0015639810599078722)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	1217: {
		Ffile: __ccgo_ts,
		Fline: int32(1234),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.009642395749675728),
		Fy:    libc.Float64FromFloat64(0.009642545168606707),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	1218: {
		Ffile: __ccgo_ts,
		Fline: int32(1235),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(49.09644854750649)),
		Fy:    float64(-libc.Float64FromFloat64(1.0502354888004449e+21)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1219: {
		Ffile: __ccgo_ts,
		Fline: int32(1236),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(20.757271252061834)),
		Fy:    float64(-libc.Float64FromFloat64(5.172951045691734e+08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1220: {
		Ffile: __ccgo_ts,
		Fline: int32(1237),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.013422779068729478),
		Fy:    libc.Float64FromFloat64(0.013423182138277163),
		Fdy:   float32(-libc.Float64FromFloat64(3.879552252783469e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1221: {
		Ffile: __ccgo_ts,
		Fline: int32(1238),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.013422779068729478),
		Fy:    libc.Float64FromFloat64(0.013423182138277165),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1222: {
		Ffile: __ccgo_ts,
		Fline: int32(1239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.013422779068729478),
		Fy:    libc.Float64FromFloat64(0.013423182138277163),
		Fdy:   float32(-libc.Float64FromFloat64(3.879552252783469e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1223: {
		Ffile: __ccgo_ts,
		Fline: int32(1240),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.26431369241603325)),
		Fy:    float64(-libc.Float64FromFloat64(0.2674020290755088)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999964237213135)),
		Fe:    int32(FE_INEXACT),
	},
	1224: {
		Ffile: __ccgo_ts,
		Fline: int32(1241),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(247.56259246692213)),
		Fy:    float64(-libc.Float64FromFloat64(1.6369591370780732e+107)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1225: {
		Ffile: __ccgo_ts,
		Fline: int32(1242),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11586295790588907),
		Fy:    libc.Float64FromFloat64(0.11612236036149998),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1226: {
		Ffile: __ccgo_ts,
		Fline: int32(1243),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(247.0002100978724),
		Fy:    libc.Float64FromFloat64(9.328209232515992e+106),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1227: {
		Ffile: __ccgo_ts,
		Fline: int32(1244),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(247.0002100978724),
		Fy:    libc.Float64FromFloat64(9.328209232515993e+106),
		Fdy:   float32(6.295862817751186e-07),
		Fe:    int32(FE_INEXACT),
	},
	1228: {
		Ffile: __ccgo_ts,
		Fline: int32(1245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(247.0002100978724),
		Fy:    libc.Float64FromFloat64(9.328209232515992e+106),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1229: {
		Ffile: __ccgo_ts,
		Fline: int32(1246),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(6.178901613471149)),
		Fy:    float64(-libc.Float64FromFloat64(241.22983151902304)),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1230: {
		Ffile: __ccgo_ts,
		Fline: int32(1247),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(6.178901613471149)),
		Fy:    float64(-libc.Float64FromFloat64(241.229831519023)),
		Fdy:   float32(5.655664949699712e-07),
		Fe:    int32(FE_INEXACT),
	},
	1231: {
		Ffile: __ccgo_ts,
		Fline: int32(1248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(6.178901613471149)),
		Fy:    float64(-libc.Float64FromFloat64(241.229831519023)),
		Fdy:   float32(5.655664949699712e-07),
		Fe:    int32(FE_INEXACT),
	},
	1232: {
		Ffile: __ccgo_ts,
		Fline: int32(1249),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.037538865335789554)),
		Fy:    float64(-libc.Float64FromFloat64(0.03754768237502156)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1233: {
		Ffile: __ccgo_ts,
		Fline: int32(1250),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.003094096521153616),
		Fy:    libc.Float64FromFloat64(0.0030941014580104145),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1234: {
		Ffile: __ccgo_ts,
		Fline: int32(1251),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.003094096521153616),
		Fy:    libc.Float64FromFloat64(0.003094101458010415),
		Fdy:   float32(3.4075929988830467e-07),
		Fe:    int32(FE_INEXACT),
	},
	1235: {
		Ffile: __ccgo_ts,
		Fline: int32(1252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.003094096521153616),
		Fy:    libc.Float64FromFloat64(0.0030941014580104145),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1236: {
		Ffile: __ccgo_ts,
		Fline: int32(1253),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0050982439829042)),
		Fy:    float64(-libc.Float64FromFloat64(0.005098266068603762)),
		Fdy:   float32(-libc.Float64FromFloat64(3.2239060487881943e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1237: {
		Ffile: __ccgo_ts,
		Fline: int32(1254),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0050982439829042)),
		Fy:    float64(-libc.Float64FromFloat64(0.005098266068603761)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1238: {
		Ffile: __ccgo_ts,
		Fline: int32(1255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0050982439829042)),
		Fy:    float64(-libc.Float64FromFloat64(0.005098266068603761)),
		Fdy:   float32(0.9999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1239: {
		Ffile: __ccgo_ts,
		Fline: int32(1256),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(11.106046755377752),
		Fy:    libc.Float64FromFloat64(33286.24656514406),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1240: {
		Ffile: __ccgo_ts,
		Fline: int32(1257),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(11.106046755377752),
		Fy:    libc.Float64FromFloat64(33286.246565144065),
		Fdy:   float32(8.491747394145932e-08),
		Fe:    int32(FE_INEXACT),
	},
	1241: {
		Ffile: __ccgo_ts,
		Fline: int32(1258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.106046755377752),
		Fy:    libc.Float64FromFloat64(33286.24656514406),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1242: {
		Ffile: __ccgo_ts,
		Fline: int32(1259),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.800305357186696)),
		Fy:    float64(-libc.Float64FromFloat64(60.76965001456885)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1243: {
		Ffile: __ccgo_ts,
		Fline: int32(1260),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.800305357186696)),
		Fy:    float64(-libc.Float64FromFloat64(60.769650014568846)),
		Fdy:   float32(4.128646935441793e-07),
		Fe:    int32(FE_INEXACT),
	},
	1244: {
		Ffile: __ccgo_ts,
		Fline: int32(1261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(4.800305357186696)),
		Fy:    float64(-libc.Float64FromFloat64(60.769650014568846)),
		Fdy:   float32(4.128646935441793e-07),
		Fe:    int32(FE_INEXACT),
	},
	1245: {
		Ffile: __ccgo_ts,
		Fline: int32(1262),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.428981275547701),
		Fy:    libc.Float64FromFloat64(0.4422600786115471),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	1246: {
		Ffile: __ccgo_ts,
		Fline: int32(1263),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.2768460281557712)),
		Fy:    float64(-libc.Float64FromFloat64(1.6531991442929364)),
		Fdy:   float32(0.4999997913837433),
		Fe:    int32(FE_INEXACT),
	},
	1247: {
		Ffile: __ccgo_ts,
		Fline: int32(1264),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.04055516635616158)),
		Fy:    float64(-libc.Float64FromFloat64(0.04056628426287842)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1248: {
		Ffile: __ccgo_ts,
		Fline: int32(1265),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6168131067575104),
		Fy:    libc.Float64FromFloat64(2.419240951563887),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1249: {
		Ffile: __ccgo_ts,
		Fline: int32(1266),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0026923068572599464),
		Fy:    libc.Float64FromFloat64(0.002692310109799443),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997913837433)),
		Fe:    int32(FE_INEXACT),
	},
	1250: {
		Ffile: __ccgo_ts,
		Fline: int32(1267),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0015627950905462045)),
		Fy:    float64(-libc.Float64FromFloat64(0.0015627957266894457)),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1251: {
		Ffile: __ccgo_ts,
		Fline: int32(1268),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(176.62941192732478),
		Fy:    libc.Float64FromFloat64(2.5594635332764036e+76),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	1252: {
		Ffile: __ccgo_ts,
		Fline: int32(1269),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(480.29569470589394),
		Fy:    libc.Float64FromFloat64(1.9441953361957527e+208),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1253: {
		Ffile: __ccgo_ts,
		Fline: int32(1270),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(480.29569470589394),
		Fy:    libc.Float64FromFloat64(1.944195336195753e+208),
		Fdy:   float32(5.163862510926265e-07),
		Fe:    int32(FE_INEXACT),
	},
	1254: {
		Ffile: __ccgo_ts,
		Fline: int32(1271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(480.29569470589394),
		Fy:    libc.Float64FromFloat64(1.9441953361957527e+208),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1255: {
		Ffile: __ccgo_ts,
		Fline: int32(1272),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.41790909454925457),
		Fy:    libc.Float64FromFloat64(0.43018026132460013),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	1256: {
		Ffile: __ccgo_ts,
		Fline: int32(1273),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.8771669588955652),
		Fy:    libc.Float64FromFloat64(0.9940597597828145),
		Fdy:   float32(-libc.Float64FromFloat64(8.28992597234901e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1257: {
		Ffile: __ccgo_ts,
		Fline: int32(1274),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.8771669588955652),
		Fy:    libc.Float64FromFloat64(0.9940597597828146),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	1258: {
		Ffile: __ccgo_ts,
		Fline: int32(1275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.8771669588955652),
		Fy:    libc.Float64FromFloat64(0.9940597597828145),
		Fdy:   float32(-libc.Float64FromFloat64(8.28992597234901e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1259: {
		Ffile: __ccgo_ts,
		Fline: int32(1276),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(65.58312064027763)),
		Fy:    float64(-libc.Float64FromFloat64(1.5182993440617435e+28)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	1260: {
		Ffile: __ccgo_ts,
		Fline: int32(1277),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.46619572198332937),
		Fy:    libc.Float64FromFloat64(0.4832672273504518),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	1261: {
		Ffile: __ccgo_ts,
		Fline: int32(1278),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(213.24795269033928)),
		Fy:    float64(-libc.Float64FromFloat64(2.0482319432961357e+92)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1262: {
		Ffile: __ccgo_ts,
		Fline: int32(1279),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.599471564217223)),
		Fy:    float64(-libc.Float64FromFloat64(135.13992118421058)),
		Fdy:   float32(-libc.Float64FromFloat64(6.56900681406114e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1263: {
		Ffile: __ccgo_ts,
		Fline: int32(1280),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.599471564217223)),
		Fy:    float64(-libc.Float64FromFloat64(135.13992118421055)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1264: {
		Ffile: __ccgo_ts,
		Fline: int32(1281),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(5.599471564217223)),
		Fy:    float64(-libc.Float64FromFloat64(135.13992118421055)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1265: {
		Ffile: __ccgo_ts,
		Fline: int32(1282),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004614803834693622),
		Fy:    libc.Float64FromFloat64(0.004614820214506895),
		Fdy:   float32(-libc.Float64FromFloat64(2.124103559708601e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1266: {
		Ffile: __ccgo_ts,
		Fline: int32(1283),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004614803834693622),
		Fy:    libc.Float64FromFloat64(0.004614820214506896),
		Fdy:   float32(0.9999997615814209),
		Fe:    int32(FE_INEXACT),
	},
	1267: {
		Ffile: __ccgo_ts,
		Fline: int32(1284),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004614803834693622),
		Fy:    libc.Float64FromFloat64(0.004614820214506895),
		Fdy:   float32(-libc.Float64FromFloat64(2.124103559708601e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1268: {
		Ffile: __ccgo_ts,
		Fline: int32(1285),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(19.19321839775755),
		Fy:    libc.Float64FromFloat64(1.0826269764162035e+08),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1269: {
		Ffile: __ccgo_ts,
		Fline: int32(1286),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004235330973886889)),
		Fy:    float64(-libc.Float64FromFloat64(0.00423534363614617)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1270: {
		Ffile: __ccgo_ts,
		Fline: int32(1287),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004235330973886889)),
		Fy:    float64(-libc.Float64FromFloat64(0.0042353436361461695)),
		Fdy:   float32(2.5625743660384614e-07),
		Fe:    int32(FE_INEXACT),
	},
	1271: {
		Ffile: __ccgo_ts,
		Fline: int32(1288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.004235330973886889)),
		Fy:    float64(-libc.Float64FromFloat64(0.0042353436361461695)),
		Fdy:   float32(2.5625743660384614e-07),
		Fe:    int32(FE_INEXACT),
	},
	1272: {
		Ffile: __ccgo_ts,
		Fline: int32(1289),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.29691675115125343),
		Fy:    libc.Float64FromFloat64(0.3012986970208847),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1273: {
		Ffile: __ccgo_ts,
		Fline: int32(1290),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.013858821472876934)),
		Fy:    float64(-libc.Float64FromFloat64(0.013859265114025684)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995529651642)),
		Fe:    int32(FE_INEXACT),
	},
	1274: {
		Ffile: __ccgo_ts,
		Fline: int32(1291),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.004427635069365701)),
		Fy:    float64(-libc.Float64FromFloat64(0.004427649535904337)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	1275: {
		Ffile: __ccgo_ts,
		Fline: int32(1292),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.2933303310514345),
		Fy:    libc.Float64FromFloat64(0.2975549531663258),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1276: {
		Ffile: __ccgo_ts,
		Fline: int32(1293),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(27.808749188469875)),
		Fy:    float64(-libc.Float64FromFloat64(5.972502297965205e+11)),
		Fdy:   float32(0.4999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1277: {
		Ffile: __ccgo_ts,
		Fline: int32(1294),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(497.30907363000256)),
		Fy:    float64(-libc.Float64FromFloat64(4.759446860543728e+215)),
		Fdy:   float32(0.4999997019767761),
		Fe:    int32(FE_INEXACT),
	},
	1278: {
		Ffile: __ccgo_ts,
		Fline: int32(1295),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.4825209533154668),
		Fy:    libc.Float64FromFloat64(0.5014640787580181),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	1279: {
		Ffile: __ccgo_ts,
		Fline: int32(1296),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.473564891730571)),
		Fy:    float64(-libc.Float64FromFloat64(5.890191973272671)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1280: {
		Ffile: __ccgo_ts,
		Fline: int32(1297),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.001312086904249768),
		Fy:    libc.Float64FromFloat64(0.0013120872807251562),
		Fdy:   float32(0.49999988079071045),
		Fe:    int32(FE_INEXACT),
	},
	1281: {
		Ffile: __ccgo_ts,
		Fline: int32(1298),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004538338290454858)),
		Fy:    float64(-libc.Float64FromFloat64(0.004538353869462591)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1282: {
		Ffile: __ccgo_ts,
		Fline: int32(1299),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.004538338290454858)),
		Fy:    float64(-libc.Float64FromFloat64(0.00453835386946259)),
		Fdy:   float32(1.7509218253053405e-07),
		Fe:    int32(FE_INEXACT),
	},
	1283: {
		Ffile: __ccgo_ts,
		Fline: int32(1300),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.004538338290454858)),
		Fy:    float64(-libc.Float64FromFloat64(0.00453835386946259)),
		Fdy:   float32(1.7509218253053405e-07),
		Fe:    int32(FE_INEXACT),
	},
	1284: {
		Ffile: __ccgo_ts,
		Fline: int32(1301),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0016453392000921223),
		Fy:    libc.Float64FromFloat64(0.0016453399424531134),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1285: {
		Ffile: __ccgo_ts,
		Fline: int32(1302),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0016453392000921223),
		Fy:    libc.Float64FromFloat64(0.0016453399424531136),
		Fdy:   float32(6.698763854728895e-07),
		Fe:    int32(FE_INEXACT),
	},
	1286: {
		Ffile: __ccgo_ts,
		Fline: int32(1303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0016453392000921223),
		Fy:    libc.Float64FromFloat64(0.0016453399424531134),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1287: {
		Ffile: __ccgo_ts,
		Fline: int32(1304),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.710507744445426),
		Fy:    libc.Float64FromFloat64(7.485202726037672),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1288: {
		Ffile: __ccgo_ts,
		Fline: int32(1305),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.710507744445426),
		Fy:    libc.Float64FromFloat64(7.485202726037673),
		Fdy:   float32(2.1719702658629103e-07),
		Fe:    int32(FE_INEXACT),
	},
	1289: {
		Ffile: __ccgo_ts,
		Fline: int32(1306),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.710507744445426),
		Fy:    libc.Float64FromFloat64(7.485202726037672),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1290: {
		Ffile: __ccgo_ts,
		Fline: int32(1307),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6914436654887405),
		Fy:    libc.Float64FromFloat64(0.7478716933677023),
		Fdy:   float32(0.49999943375587463),
		Fe:    int32(FE_INEXACT),
	},
	1291: {
		Ffile: __ccgo_ts,
		Fline: int32(1308),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.7254532048208885)),
		Fy:    float64(-libc.Float64FromFloat64(0.7907809688360167)),
		Fdy:   float32(0.499999076128006),
		Fe:    int32(FE_INEXACT),
	},
	1292: {
		Ffile: __ccgo_ts,
		Fline: int32(1309),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(50.662948915507336),
		Fy:    libc.Float64FromFloat64(5.030476856344886e+21),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1293: {
		Ffile: __ccgo_ts,
		Fline: int32(1310),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(50.662948915507336),
		Fy:    libc.Float64FromFloat64(5.030476856344887e+21),
		Fdy:   float32(1.6315132711497426e-07),
		Fe:    int32(FE_INEXACT),
	},
	1294: {
		Ffile: __ccgo_ts,
		Fline: int32(1311),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(50.662948915507336),
		Fy:    libc.Float64FromFloat64(5.030476856344886e+21),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998211860657)),
		Fe:    int32(FE_INEXACT),
	},
	1295: {
		Ffile: __ccgo_ts,
		Fline: int32(1312),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.006901060201725372),
		Fy:    libc.Float64FromFloat64(0.006901114978597788),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1296: {
		Ffile: __ccgo_ts,
		Fline: int32(1313),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.006901060201725372),
		Fy:    libc.Float64FromFloat64(0.006901114978597789),
		Fdy:   float32(8.873419119481696e-07),
		Fe:    int32(FE_INEXACT),
	},
	1297: {
		Ffile: __ccgo_ts,
		Fline: int32(1314),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.006901060201725372),
		Fy:    libc.Float64FromFloat64(0.006901114978597788),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1298: {
		Ffile: __ccgo_ts,
		Fline: int32(1315),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.08775912975183811)),
		Fy:    float64(-libc.Float64FromFloat64(0.08787182170641879)),
		Fdy:   float32(-libc.Float64FromFloat64(3.8599219465140777e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1299: {
		Ffile: __ccgo_ts,
		Fline: int32(1316),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.08775912975183811)),
		Fy:    float64(-libc.Float64FromFloat64(0.08787182170641877)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1300: {
		Ffile: __ccgo_ts,
		Fline: int32(1317),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.08775912975183811)),
		Fy:    float64(-libc.Float64FromFloat64(0.08787182170641877)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1301: {
		Ffile: __ccgo_ts,
		Fline: int32(1318),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.007158910063167654)),
		Fy:    float64(-libc.Float64FromFloat64(0.007158971212339803)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1302: {
		Ffile: __ccgo_ts,
		Fline: int32(1319),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.007158910063167654)),
		Fy:    float64(-libc.Float64FromFloat64(0.007158971212339802)),
		Fdy:   float32(9.754290886121453e-08),
		Fe:    int32(FE_INEXACT),
	},
	1303: {
		Ffile: __ccgo_ts,
		Fline: int32(1320),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.007158910063167654)),
		Fy:    float64(-libc.Float64FromFloat64(0.007158971212339802)),
		Fdy:   float32(9.754290886121453e-08),
		Fe:    int32(FE_INEXACT),
	},
	1304: {
		Ffile: __ccgo_ts,
		Fline: int32(1321),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.014012969497574227),
		Fy:    libc.Float64FromFloat64(0.014013428107598827),
		Fdy:   float32(-libc.Float64FromFloat64(8.864235496730544e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1305: {
		Ffile: __ccgo_ts,
		Fline: int32(1322),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.014012969497574227),
		Fy:    libc.Float64FromFloat64(0.014013428107598829),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1306: {
		Ffile: __ccgo_ts,
		Fline: int32(1323),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.014012969497574227),
		Fy:    libc.Float64FromFloat64(0.014013428107598827),
		Fdy:   float32(-libc.Float64FromFloat64(8.864235496730544e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1307: {
		Ffile: __ccgo_ts,
		Fline: int32(1324),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8821988556028172)),
		Fy:    float64(-libc.Float64FromFloat64(1.0011674466896832)),
		Fdy:   float32(-libc.Float64FromFloat64(5.262346576273558e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1308: {
		Ffile: __ccgo_ts,
		Fline: int32(1325),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8821988556028172)),
		Fy:    float64(-libc.Float64FromFloat64(1.001167446689683)),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1309: {
		Ffile: __ccgo_ts,
		Fline: int32(1326),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.8821988556028172)),
		Fy:    float64(-libc.Float64FromFloat64(1.001167446689683)),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1310: {
		Ffile: __ccgo_ts,
		Fline: int32(1327),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(31.177089505567004),
		Fy:    libc.Float64FromFloat64(1.7338356900024002e+13),
		Fdy:   float32(0.49999985098838806),
		Fe:    int32(FE_INEXACT),
	},
	1311: {
		Ffile: __ccgo_ts,
		Fline: int32(1328),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7560524763553639)),
		Fy:    float64(-libc.Float64FromFloat64(0.8301678815388579)),
		Fdy:   float32(-libc.Float64FromFloat64(1.1603687255501427e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1312: {
		Ffile: __ccgo_ts,
		Fline: int32(1329),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.7560524763553639)),
		Fy:    float64(-libc.Float64FromFloat64(0.8301678815388578)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1313: {
		Ffile: __ccgo_ts,
		Fline: int32(1330),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.7560524763553639)),
		Fy:    float64(-libc.Float64FromFloat64(0.8301678815388578)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1314: {
		Ffile: __ccgo_ts,
		Fline: int32(1331),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.28692413914401343),
		Fy:    libc.Float64FromFloat64(0.29087723650466124),
		Fdy:   float32(-libc.Float64FromFloat64(7.370698540398735e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1315: {
		Ffile: __ccgo_ts,
		Fline: int32(1332),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.28692413914401343),
		Fy:    libc.Float64FromFloat64(0.2908772365046613),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1316: {
		Ffile: __ccgo_ts,
		Fline: int32(1333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.28692413914401343),
		Fy:    libc.Float64FromFloat64(0.29087723650466124),
		Fdy:   float32(-libc.Float64FromFloat64(7.370698540398735e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1317: {
		Ffile: __ccgo_ts,
		Fline: int32(1334),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.04839747009951355),
		Fy:    libc.Float64FromFloat64(0.04841636599998516),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1318: {
		Ffile: __ccgo_ts,
		Fline: int32(1335),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.04839747009951355),
		Fy:    libc.Float64FromFloat64(0.04841636599998517),
		Fdy:   float32(3.768928706904262e-07),
		Fe:    int32(FE_INEXACT),
	},
	1319: {
		Ffile: __ccgo_ts,
		Fline: int32(1336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.04839747009951355),
		Fy:    libc.Float64FromFloat64(0.04841636599998516),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1320: {
		Ffile: __ccgo_ts,
		Fline: int32(1337),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.21429097411529302)),
		Fy:    float64(-libc.Float64FromFloat64(0.2159348063247585)),
		Fdy:   float32(-libc.Float64FromFloat64(5.945448720012791e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1321: {
		Ffile: __ccgo_ts,
		Fline: int32(1338),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.21429097411529302)),
		Fy:    float64(-libc.Float64FromFloat64(0.21593480632475848)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1322: {
		Ffile: __ccgo_ts,
		Fline: int32(1339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.21429097411529302)),
		Fy:    float64(-libc.Float64FromFloat64(0.21593480632475848)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1323: {
		Ffile: __ccgo_ts,
		Fline: int32(1340),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.003838968213670434)),
		Fy:    float64(-libc.Float64FromFloat64(0.003838977643256272)),
		Fdy:   float32(0.4999997317790985),
		Fe:    int32(FE_INEXACT),
	},
	1324: {
		Ffile: __ccgo_ts,
		Fline: int32(1341),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(174.39193144081784)),
		Fy:    float64(-libc.Float64FromFloat64(2.7316403394731262e+75)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	1325: {
		Ffile: __ccgo_ts,
		Fline: int32(1342),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8322958436390403)),
		Fy:    float64(-libc.Float64FromFloat64(0.9317702782054633)),
		Fdy:   float32(-libc.Float64FromFloat64(1.1102081742819792e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1326: {
		Ffile: __ccgo_ts,
		Fline: int32(1343),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.8322958436390403)),
		Fy:    float64(-libc.Float64FromFloat64(0.9317702782054632)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1327: {
		Ffile: __ccgo_ts,
		Fline: int32(1344),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.8322958436390403)),
		Fy:    float64(-libc.Float64FromFloat64(0.9317702782054632)),
		Fdy:   float32(0.9999998807907104),
		Fe:    int32(FE_INEXACT),
	},
	1328: {
		Ffile: __ccgo_ts,
		Fline: int32(1345),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0054450100023686904),
		Fy:    libc.Float64FromFloat64(0.005445036908139539),
		Fdy:   float32(-libc.Float64FromFloat64(2.0802829681088042e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1329: {
		Ffile: __ccgo_ts,
		Fline: int32(1346),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0054450100023686904),
		Fy:    libc.Float64FromFloat64(0.00544503690813954),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1330: {
		Ffile: __ccgo_ts,
		Fline: int32(1347),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0054450100023686904),
		Fy:    libc.Float64FromFloat64(0.005445036908139539),
		Fdy:   float32(-libc.Float64FromFloat64(2.0802829681088042e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1331: {
		Ffile: __ccgo_ts,
		Fline: int32(1348),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(11.509022476665512)),
		Fy:    float64(-libc.Float64FromFloat64(49805.230917708206)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1332: {
		Ffile: __ccgo_ts,
		Fline: int32(1349),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(30.7109525853602),
		Fy:    libc.Float64FromFloat64(1.0878455665803428e+13),
		Fdy:   float32(-libc.Float64FromFloat64(5.413661483544274e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1333: {
		Ffile: __ccgo_ts,
		Fline: int32(1350),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(30.7109525853602),
		Fy:    libc.Float64FromFloat64(1.087845566580343e+13),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1334: {
		Ffile: __ccgo_ts,
		Fline: int32(1351),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(30.7109525853602),
		Fy:    libc.Float64FromFloat64(1.0878455665803428e+13),
		Fdy:   float32(-libc.Float64FromFloat64(5.413661483544274e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1335: {
		Ffile: __ccgo_ts,
		Fline: int32(1352),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(209.11625925336006),
		Fy:    libc.Float64FromFloat64(3.2885729054612166e+90),
		Fdy:   float32(-libc.Float64FromFloat64(7.513970672334835e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1336: {
		Ffile: __ccgo_ts,
		Fline: int32(1353),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(209.11625925336006),
		Fy:    libc.Float64FromFloat64(3.288572905461217e+90),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1337: {
		Ffile: __ccgo_ts,
		Fline: int32(1354),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(209.11625925336006),
		Fy:    libc.Float64FromFloat64(3.2885729054612166e+90),
		Fdy:   float32(-libc.Float64FromFloat64(7.513970672334835e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1338: {
		Ffile: __ccgo_ts,
		Fline: int32(1355),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(6.600351002849948),
		Fy:    libc.Float64FromFloat64(367.6759475730082),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1339: {
		Ffile: __ccgo_ts,
		Fline: int32(1356),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(6.600351002849948),
		Fy:    libc.Float64FromFloat64(367.6759475730083),
		Fdy:   float32(5.360393515729811e-07),
		Fe:    int32(FE_INEXACT),
	},
	1340: {
		Ffile: __ccgo_ts,
		Fline: int32(1357),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.600351002849948),
		Fy:    libc.Float64FromFloat64(367.6759475730082),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1341: {
		Ffile: __ccgo_ts,
		Fline: int32(1358),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(508.30974732036645)),
		Fy:    float64(-libc.Float64FromFloat64(2.851598405542914e+220)),
		Fdy:   float32(-libc.Float64FromFloat64(5.003794285585172e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1342: {
		Ffile: __ccgo_ts,
		Fline: int32(1359),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(508.30974732036645)),
		Fy:    float64(-libc.Float64FromFloat64(2.8515984055429134e+220)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1343: {
		Ffile: __ccgo_ts,
		Fline: int32(1360),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(508.30974732036645)),
		Fy:    float64(-libc.Float64FromFloat64(2.8515984055429134e+220)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1344: {
		Ffile: __ccgo_ts,
		Fline: int32(1361),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(30.15709726547301)),
		Fy:    float64(-libc.Float64FromFloat64(6.252172285144656e+12)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1345: {
		Ffile: __ccgo_ts,
		Fline: int32(1362),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(30.15709726547301)),
		Fy:    float64(-libc.Float64FromFloat64(6.252172285144655e+12)),
		Fdy:   float32(8.126878014991235e-07),
		Fe:    int32(FE_INEXACT),
	},
	1346: {
		Ffile: __ccgo_ts,
		Fline: int32(1363),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(30.15709726547301)),
		Fy:    float64(-libc.Float64FromFloat64(6.252172285144655e+12)),
		Fdy:   float32(8.126878014991235e-07),
		Fe:    int32(FE_INEXACT),
	},
	1347: {
		Ffile: __ccgo_ts,
		Fline: int32(1364),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.05277333988511369),
		Fy:    libc.Float64FromFloat64(0.05279783914511154),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999916553497314)),
		Fe:    int32(FE_INEXACT),
	},
	1348: {
		Ffile: __ccgo_ts,
		Fline: int32(1365),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5062345148846947),
		Fy:    libc.Float64FromFloat64(2.143987064669041),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999701976776)),
		Fe:    int32(FE_INEXACT),
	},
	1349: {
		Ffile: __ccgo_ts,
		Fline: int32(1366),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.8512185322480907)),
		Fy:    float64(-libc.Float64FromFloat64(8.625543260725578)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	1350: {
		Ffile: __ccgo_ts,
		Fline: int32(1367),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(431.18147132805916),
		Fy:    libc.Float64FromFloat64(9.09292692769934e+186),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1351: {
		Ffile: __ccgo_ts,
		Fline: int32(1368),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(431.18147132805916),
		Fy:    libc.Float64FromFloat64(9.092926927699343e+186),
		Fdy:   float32(5.454712663777173e-07),
		Fe:    int32(FE_INEXACT),
	},
	1352: {
		Ffile: __ccgo_ts,
		Fline: int32(1369),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(431.18147132805916),
		Fy:    libc.Float64FromFloat64(9.09292692769934e+186),
		Fdy:   float32(-libc.Float64FromFloat64(0.999999463558197)),
		Fe:    int32(FE_INEXACT),
	},
	1353: {
		Ffile: __ccgo_ts,
		Fline: int32(1370),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.012170900804288736)),
		Fy:    float64(-libc.Float64FromFloat64(0.012171201287446606)),
		Fdy:   float32(-libc.Float64FromFloat64(5.777062597189797e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1354: {
		Ffile: __ccgo_ts,
		Fline: int32(1371),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.012170900804288736)),
		Fy:    float64(-libc.Float64FromFloat64(0.012171201287446604)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1355: {
		Ffile: __ccgo_ts,
		Fline: int32(1372),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.012170900804288736)),
		Fy:    float64(-libc.Float64FromFloat64(0.012171201287446604)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1356: {
		Ffile: __ccgo_ts,
		Fline: int32(1373),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0023393478261702918),
		Fy:    libc.Float64FromFloat64(0.0023393499598698517),
		Fdy:   float32(-libc.Float64FromFloat64(8.974171805675724e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1357: {
		Ffile: __ccgo_ts,
		Fline: int32(1374),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0023393478261702918),
		Fy:    libc.Float64FromFloat64(0.002339349959869852),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1358: {
		Ffile: __ccgo_ts,
		Fline: int32(1375),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0023393478261702918),
		Fy:    libc.Float64FromFloat64(0.0023393499598698517),
		Fdy:   float32(-libc.Float64FromFloat64(8.974171805675724e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1359: {
		Ffile: __ccgo_ts,
		Fline: int32(1376),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.004009501448589755),
		Fy:    libc.Float64FromFloat64(0.004009512191457343),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1360: {
		Ffile: __ccgo_ts,
		Fline: int32(1377),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.004009501448589755),
		Fy:    libc.Float64FromFloat64(0.004009512191457344),
		Fdy:   float32(3.9734561596560525e-07),
		Fe:    int32(FE_INEXACT),
	},
	1361: {
		Ffile: __ccgo_ts,
		Fline: int32(1378),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.004009501448589755),
		Fy:    libc.Float64FromFloat64(0.004009512191457343),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1362: {
		Ffile: __ccgo_ts,
		Fline: int32(1379),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(343.1107617981876),
		Fy:    libc.Float64FromFloat64(5.12956497046992e+148),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1363: {
		Ffile: __ccgo_ts,
		Fline: int32(1380),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(343.1107617981876),
		Fy:    libc.Float64FromFloat64(5.129564970469921e+148),
		Fdy:   float32(5.859269549546298e-07),
		Fe:    int32(FE_INEXACT),
	},
	1364: {
		Ffile: __ccgo_ts,
		Fline: int32(1381),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(343.1107617981876),
		Fy:    libc.Float64FromFloat64(5.12956497046992e+148),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1365: {
		Ffile: __ccgo_ts,
		Fline: int32(1382),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(57.6548048153395)),
		Fy:    float64(-libc.Float64FromFloat64(5.47184253234507e+24)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	1366: {
		Ffile: __ccgo_ts,
		Fline: int32(1383),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(26.00744465390955)),
		Fy:    float64(-libc.Float64FromFloat64(9.859609303215547e+10)),
		Fdy:   float32(0.49999940395355225),
		Fe:    int32(FE_INEXACT),
	},
	1367: {
		Ffile: __ccgo_ts,
		Fline: int32(1384),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(27.64908830624743),
		Fy:    libc.Float64FromFloat64(5.0911569469158624e+11),
		Fdy:   float32(-libc.Float64FromFloat64(2.0573472170326568e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1368: {
		Ffile: __ccgo_ts,
		Fline: int32(1385),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(27.64908830624743),
		Fy:    libc.Float64FromFloat64(5.091156946915863e+11),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1369: {
		Ffile: __ccgo_ts,
		Fline: int32(1386),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(27.64908830624743),
		Fy:    libc.Float64FromFloat64(5.0911569469158624e+11),
		Fdy:   float32(-libc.Float64FromFloat64(2.0573472170326568e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1370: {
		Ffile: __ccgo_ts,
		Fline: int32(1387),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.014093652344241236),
		Fy:    libc.Float64FromFloat64(0.01409411892167035),
		Fdy:   float32(-libc.Float64FromFloat64(6.816337645432213e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1371: {
		Ffile: __ccgo_ts,
		Fline: int32(1388),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.014093652344241236),
		Fy:    libc.Float64FromFloat64(0.014094118921670352),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1372: {
		Ffile: __ccgo_ts,
		Fline: int32(1389),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.014093652344241236),
		Fy:    libc.Float64FromFloat64(0.01409411892167035),
		Fdy:   float32(-libc.Float64FromFloat64(6.816337645432213e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1373: {
		Ffile: __ccgo_ts,
		Fline: int32(1390),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(207.1772835858378)),
		Fy:    float64(-libc.Float64FromFloat64(4.730652436410769e+89)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999937415122986)),
		Fe:    int32(FE_INEXACT),
	},
	1374: {
		Ffile: __ccgo_ts,
		Fline: int32(1391),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.013177935358016918),
		Fy:    libc.Float64FromFloat64(0.013178316770268468),
		Fdy:   float32(0.49999910593032837),
		Fe:    int32(FE_INEXACT),
	},
	1375: {
		Ffile: __ccgo_ts,
		Fline: int32(1392),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.061236181789170466),
		Fy:    libc.Float64FromFloat64(0.061274460251855754),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1376: {
		Ffile: __ccgo_ts,
		Fline: int32(1393),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.061236181789170466),
		Fy:    libc.Float64FromFloat64(0.06127446025185576),
		Fdy:   float32(4.638629320652399e-07),
		Fe:    int32(FE_INEXACT),
	},
	1377: {
		Ffile: __ccgo_ts,
		Fline: int32(1394),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.061236181789170466),
		Fy:    libc.Float64FromFloat64(0.061274460251855754),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1378: {
		Ffile: __ccgo_ts,
		Fline: int32(1395),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.0346682457391594)),
		Fy:    float64(-libc.Float64FromFloat64(1.2294141093471405)),
		Fdy:   float32(-libc.Float64FromFloat64(9.174262345368334e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1379: {
		Ffile: __ccgo_ts,
		Fline: int32(1396),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.0346682457391594)),
		Fy:    float64(-libc.Float64FromFloat64(1.2294141093471402)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1380: {
		Ffile: __ccgo_ts,
		Fline: int32(1397),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.0346682457391594)),
		Fy:    float64(-libc.Float64FromFloat64(1.2294141093471402)),
		Fdy:   float32(0.9999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1381: {
		Ffile: __ccgo_ts,
		Fline: int32(1398),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.010804785602257724),
		Fy:    libc.Float64FromFloat64(0.010804995834704893),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1382: {
		Ffile: __ccgo_ts,
		Fline: int32(1399),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.010804785602257724),
		Fy:    libc.Float64FromFloat64(0.010804995834704895),
		Fdy:   float32(8.33589353987918e-07),
		Fe:    int32(FE_INEXACT),
	},
	1383: {
		Ffile: __ccgo_ts,
		Fline: int32(1400),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.010804785602257724),
		Fy:    libc.Float64FromFloat64(0.010804995834704893),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1384: {
		Ffile: __ccgo_ts,
		Fline: int32(1401),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(266.2547128391924),
		Fy:    libc.Float64FromFloat64(2.147447578964194e+115),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	1385: {
		Ffile: __ccgo_ts,
		Fline: int32(1402),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0015534325093140707),
		Fy:    libc.Float64FromFloat64(0.001553433134092419),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1386: {
		Ffile: __ccgo_ts,
		Fline: int32(1403),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0015534325093140707),
		Fy:    libc.Float64FromFloat64(0.0015534331340924191),
		Fdy:   float32(4.063824690092588e-07),
		Fe:    int32(FE_INEXACT),
	},
	1387: {
		Ffile: __ccgo_ts,
		Fline: int32(1404),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0015534325093140707),
		Fy:    libc.Float64FromFloat64(0.001553433134092419),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995827674866)),
		Fe:    int32(FE_INEXACT),
	},
	1388: {
		Ffile: __ccgo_ts,
		Fline: int32(1405),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.16975910060033444),
		Fy:    libc.Float64FromFloat64(0.17057563352871488),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1389: {
		Ffile: __ccgo_ts,
		Fline: int32(1406),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.16975910060033444),
		Fy:    libc.Float64FromFloat64(0.1705756335287149),
		Fdy:   float32(5.996039362798911e-07),
		Fe:    int32(FE_INEXACT),
	},
	1390: {
		Ffile: __ccgo_ts,
		Fline: int32(1407),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.16975910060033444),
		Fy:    libc.Float64FromFloat64(0.17057563352871488),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1391: {
		Ffile: __ccgo_ts,
		Fline: int32(1408),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0992123849369932)),
		Fy:    float64(-libc.Float64FromFloat64(0.0993752245854566)),
		Fdy:   float32(-libc.Float64FromFloat64(3.6889755961055926e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1392: {
		Ffile: __ccgo_ts,
		Fline: int32(1409),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.0992123849369932)),
		Fy:    float64(-libc.Float64FromFloat64(0.09937522458545658)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1393: {
		Ffile: __ccgo_ts,
		Fline: int32(1410),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.0992123849369932)),
		Fy:    float64(-libc.Float64FromFloat64(0.09937522458545658)),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1394: {
		Ffile: __ccgo_ts,
		Fline: int32(1411),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(186.82768879671644),
		Fy:    libc.Float64FromFloat64(6.873917500157631e+80),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1395: {
		Ffile: __ccgo_ts,
		Fline: int32(1412),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(186.82768879671644),
		Fy:    libc.Float64FromFloat64(6.873917500157632e+80),
		Fdy:   float32(5.055869678471936e-07),
		Fe:    int32(FE_INEXACT),
	},
	1396: {
		Ffile: __ccgo_ts,
		Fline: int32(1413),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(186.82768879671644),
		Fy:    libc.Float64FromFloat64(6.873917500157631e+80),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1397: {
		Ffile: __ccgo_ts,
		Fline: int32(1414),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.11133873973003598),
		Fy:    libc.Float64FromFloat64(0.11156891403029393),
		Fdy:   float32(0.49999919533729553),
		Fe:    int32(FE_INEXACT),
	},
	1398: {
		Ffile: __ccgo_ts,
		Fline: int32(1415),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(109.45088607491294)),
		Fy:    float64(-libc.Float64FromFloat64(1.709565976922297e+47)),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	1399: {
		Ffile: __ccgo_ts,
		Fline: int32(1416),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.10940359878294229)),
		Fy:    float64(-libc.Float64FromFloat64(0.10962197406408625)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1400: {
		Ffile: __ccgo_ts,
		Fline: int32(1417),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.612153591739522),
		Fy:    libc.Float64FromFloat64(50.34542712487623),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	1401: {
		Ffile: __ccgo_ts,
		Fline: int32(1418),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.612153591739522),
		Fy:    libc.Float64FromFloat64(50.34542712487624),
		Fdy:   float32(3.090949860506953e-07),
		Fe:    int32(FE_INEXACT),
	},
	1402: {
		Ffile: __ccgo_ts,
		Fline: int32(1419),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.612153591739522),
		Fy:    libc.Float64FromFloat64(50.34542712487623),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997019767761)),
		Fe:    int32(FE_INEXACT),
	},
	1403: {
		Ffile: __ccgo_ts,
		Fline: int32(1420),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.26274713660018056)),
		Fy:    float64(-libc.Float64FromFloat64(0.26578076024462216)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999919533729553)),
		Fe:    int32(FE_INEXACT),
	},
	1404: {
		Ffile: __ccgo_ts,
		Fline: int32(1421),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(209.81223446482798)),
		Fy:    float64(-libc.Float64FromFloat64(6.595772508586045e+90)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1405: {
		Ffile: __ccgo_ts,
		Fline: int32(1422),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.006916852313748974),
		Fy:    libc.Float64FromFloat64(0.006916907467530838),
		Fdy:   float32(-libc.Float64FromFloat64(3.419154950279335e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1406: {
		Ffile: __ccgo_ts,
		Fline: int32(1423),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.006916852313748974),
		Fy:    libc.Float64FromFloat64(0.006916907467530839),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1407: {
		Ffile: __ccgo_ts,
		Fline: int32(1424),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.006916852313748974),
		Fy:    libc.Float64FromFloat64(0.006916907467530838),
		Fdy:   float32(-libc.Float64FromFloat64(3.419154950279335e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1408: {
		Ffile: __ccgo_ts,
		Fline: int32(1425),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(50.12729281882736)),
		Fy:    float64(-libc.Float64FromFloat64(2.9442634542313835e+21)),
		Fdy:   float32(-libc.Float64FromFloat64(6.285833364927385e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1409: {
		Ffile: __ccgo_ts,
		Fline: int32(1426),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(50.12729281882736)),
		Fy:    float64(-libc.Float64FromFloat64(2.944263454231383e+21)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1410: {
		Ffile: __ccgo_ts,
		Fline: int32(1427),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(50.12729281882736)),
		Fy:    float64(-libc.Float64FromFloat64(2.944263454231383e+21)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1411: {
		Ffile: __ccgo_ts,
		Fline: int32(1428),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(37.483040927760534),
		Fy:    libc.Float64FromFloat64(9.49834290453368e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1412: {
		Ffile: __ccgo_ts,
		Fline: int32(1429),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(37.483040927760534),
		Fy:    libc.Float64FromFloat64(9.498342904533682e+15),
		Fdy:   float32(1.1692226564719022e-07),
		Fe:    int32(FE_INEXACT),
	},
	1413: {
		Ffile: __ccgo_ts,
		Fline: int32(1430),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(37.483040927760534),
		Fy:    libc.Float64FromFloat64(9.49834290453368e+15),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1414: {
		Ffile: __ccgo_ts,
		Fline: int32(1431),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.710344455086933),
		Fy:    libc.Float64FromFloat64(8243.63998280067),
		Fdy:   float32(0.4999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1415: {
		Ffile: __ccgo_ts,
		Fline: int32(1432),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0031612135069855134),
		Fy:    libc.Float64FromFloat64(0.0031612187721319356),
		Fdy:   float32(-libc.Float64FromFloat64(6.68895665967284e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1416: {
		Ffile: __ccgo_ts,
		Fline: int32(1433),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0031612135069855134),
		Fy:    libc.Float64FromFloat64(0.003161218772131936),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1417: {
		Ffile: __ccgo_ts,
		Fline: int32(1434),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0031612135069855134),
		Fy:    libc.Float64FromFloat64(0.0031612187721319356),
		Fdy:   float32(-libc.Float64FromFloat64(6.68895665967284e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1418: {
		Ffile: __ccgo_ts,
		Fline: int32(1435),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(31.60775734280074)),
		Fy:    float64(-libc.Float64FromFloat64(2.66713257484656e+13)),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	1419: {
		Ffile: __ccgo_ts,
		Fline: int32(1436),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(259.75662256875586),
		Fy:    libc.Float64FromFloat64(3.2347284166841928e+112),
		Fdy:   float32(-libc.Float64FromFloat64(7.985730121617962e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1420: {
		Ffile: __ccgo_ts,
		Fline: int32(1437),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(259.75662256875586),
		Fy:    libc.Float64FromFloat64(3.234728416684193e+112),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1421: {
		Ffile: __ccgo_ts,
		Fline: int32(1438),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(259.75662256875586),
		Fy:    libc.Float64FromFloat64(3.2347284166841928e+112),
		Fdy:   float32(-libc.Float64FromFloat64(7.985730121617962e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1422: {
		Ffile: __ccgo_ts,
		Fline: int32(1439),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.030740518519594053),
		Fy:    libc.Float64FromFloat64(0.030745360274884746),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1423: {
		Ffile: __ccgo_ts,
		Fline: int32(1440),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.030740518519594053),
		Fy:    libc.Float64FromFloat64(0.03074536027488475),
		Fdy:   float32(1.3655676411872264e-07),
		Fe:    int32(FE_INEXACT),
	},
	1424: {
		Ffile: __ccgo_ts,
		Fline: int32(1441),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.030740518519594053),
		Fy:    libc.Float64FromFloat64(0.030745360274884746),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999998807907104)),
		Fe:    int32(FE_INEXACT),
	},
	1425: {
		Ffile: __ccgo_ts,
		Fline: int32(1442),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0018476080154770976),
		Fy:    libc.Float64FromFloat64(0.0018476090666601172),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	1426: {
		Ffile: __ccgo_ts,
		Fline: int32(1443),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.226843641641054),
		Fy:    libc.Float64FromFloat64(5083.200955085254),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999910593032837)),
		Fe:    int32(FE_INEXACT),
	},
	1427: {
		Ffile: __ccgo_ts,
		Fline: int32(1444),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(27.26074946976305),
		Fy:    libc.Float64FromFloat64(3.452733603939799e+11),
		Fdy:   float32(-libc.Float64FromFloat64(6.73071781420731e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1428: {
		Ffile: __ccgo_ts,
		Fline: int32(1445),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(27.26074946976305),
		Fy:    libc.Float64FromFloat64(3.4527336039398e+11),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1429: {
		Ffile: __ccgo_ts,
		Fline: int32(1446),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(27.26074946976305),
		Fy:    libc.Float64FromFloat64(3.452733603939799e+11),
		Fdy:   float32(-libc.Float64FromFloat64(6.73071781420731e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1430: {
		Ffile: __ccgo_ts,
		Fline: int32(1447),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.0051912214774361195)),
		Fy:    float64(-libc.Float64FromFloat64(0.0051912447936488275)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1431: {
		Ffile: __ccgo_ts,
		Fline: int32(1448),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(51.144331512108806)),
		Fy:    float64(-libc.Float64FromFloat64(8.140872643860235e+21)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1432: {
		Ffile: __ccgo_ts,
		Fline: int32(1449),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(290.1550505154486),
		Fy:    libc.Float64FromFloat64(5.148815615391829e+125),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	1433: {
		Ffile: __ccgo_ts,
		Fline: int32(1450),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.20039860696732337),
		Fy:    libc.Float64FromFloat64(0.20174262426267484),
		Fdy:   float32(-libc.Float64FromFloat64(5.257960538074258e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1434: {
		Ffile: __ccgo_ts,
		Fline: int32(1451),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.20039860696732337),
		Fy:    libc.Float64FromFloat64(0.20174262426267486),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1435: {
		Ffile: __ccgo_ts,
		Fline: int32(1452),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.20039860696732337),
		Fy:    libc.Float64FromFloat64(0.20174262426267484),
		Fdy:   float32(-libc.Float64FromFloat64(5.257960538074258e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1436: {
		Ffile: __ccgo_ts,
		Fline: int32(1453),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.001347796106549111),
		Fy:    libc.Float64FromFloat64(0.001347796514606627),
		Fdy:   float32(-libc.Float64FromFloat64(4.363542700502876e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1437: {
		Ffile: __ccgo_ts,
		Fline: int32(1454),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.001347796106549111),
		Fy:    libc.Float64FromFloat64(0.0013477965146066271),
		Fdy:   float32(0.9999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1438: {
		Ffile: __ccgo_ts,
		Fline: int32(1455),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.001347796106549111),
		Fy:    libc.Float64FromFloat64(0.001347796514606627),
		Fdy:   float32(-libc.Float64FromFloat64(4.363542700502876e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1439: {
		Ffile: __ccgo_ts,
		Fline: int32(1456),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(27.319637845736487)),
		Fy:    float64(-libc.Float64FromFloat64(3.6621655114545557e+11)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999994039535522)),
		Fe:    int32(FE_INEXACT),
	},
	1440: {
		Ffile: __ccgo_ts,
		Fline: int32(1457),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(27.319637845736487)),
		Fy:    float64(-libc.Float64FromFloat64(3.662165511454555e+11)),
		Fdy:   float32(5.797102744509175e-07),
		Fe:    int32(FE_INEXACT),
	},
	1441: {
		Ffile: __ccgo_ts,
		Fline: int32(1458),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(27.319637845736487)),
		Fy:    float64(-libc.Float64FromFloat64(3.662165511454555e+11)),
		Fdy:   float32(5.797102744509175e-07),
		Fe:    int32(FE_INEXACT),
	},
	1442: {
		Ffile: __ccgo_ts,
		Fline: int32(1459),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(335.42003456531734)),
		Fy:    float64(-libc.Float64FromFloat64(2.344445470080423e+145)),
		Fdy:   float32(0.4999994933605194),
		Fe:    int32(FE_INEXACT),
	},
	1443: {
		Ffile: __ccgo_ts,
		Fline: int32(1460),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.08667382125314456),
		Fy:    libc.Float64FromFloat64(0.08678238272154257),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1444: {
		Ffile: __ccgo_ts,
		Fline: int32(1461),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.08667382125314456),
		Fy:    libc.Float64FromFloat64(0.08678238272154258),
		Fdy:   float32(7.334890739230104e-08),
		Fe:    int32(FE_INEXACT),
	},
	1445: {
		Ffile: __ccgo_ts,
		Fline: int32(1462),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.08667382125314456),
		Fy:    libc.Float64FromFloat64(0.08678238272154257),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1446: {
		Ffile: __ccgo_ts,
		Fline: int32(1463),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(11.57276825696518),
		Fy:    libc.Float64FromFloat64(53083.481518330715),
		Fdy:   float32(-libc.Float64FromFloat64(1.9964700470609387e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1447: {
		Ffile: __ccgo_ts,
		Fline: int32(1464),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(11.57276825696518),
		Fy:    libc.Float64FromFloat64(53083.48151833072),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1448: {
		Ffile: __ccgo_ts,
		Fline: int32(1465),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.57276825696518),
		Fy:    libc.Float64FromFloat64(53083.481518330715),
		Fdy:   float32(-libc.Float64FromFloat64(1.9964700470609387e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1449: {
		Ffile: __ccgo_ts,
		Fline: int32(1466),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.6210361088879343)),
		Fy:    float64(-libc.Float64FromFloat64(0.661733867313747)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1450: {
		Ffile: __ccgo_ts,
		Fline: int32(1467),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.6210361088879343)),
		Fy:    float64(-libc.Float64FromFloat64(0.6617338673137468)),
		Fdy:   float32(7.209569048427511e-07),
		Fe:    int32(FE_INEXACT),
	},
	1451: {
		Ffile: __ccgo_ts,
		Fline: int32(1468),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.6210361088879343)),
		Fy:    float64(-libc.Float64FromFloat64(0.6617338673137468)),
		Fdy:   float32(7.209569048427511e-07),
		Fe:    int32(FE_INEXACT),
	},
	1452: {
		Ffile: __ccgo_ts,
		Fline: int32(1469),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(45.4855407082272)),
		Fy:    float64(-libc.Float64FromFloat64(2.8385028871866847e+19)),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	1453: {
		Ffile: __ccgo_ts,
		Fline: int32(1470),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.008415857700526323)),
		Fy:    float64(-libc.Float64FromFloat64(0.008415957045394636)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	1454: {
		Ffile: __ccgo_ts,
		Fline: int32(1471),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.9646084090084114),
		Fy:    libc.Float64FromFloat64(1.1213137979757526),
		Fdy:   float32(-libc.Float64FromFloat64(5.110093752591638e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1455: {
		Ffile: __ccgo_ts,
		Fline: int32(1472),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.9646084090084114),
		Fy:    libc.Float64FromFloat64(1.1213137979757528),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1456: {
		Ffile: __ccgo_ts,
		Fline: int32(1473),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.9646084090084114),
		Fy:    libc.Float64FromFloat64(1.1213137979757526),
		Fdy:   float32(-libc.Float64FromFloat64(5.110093752591638e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1457: {
		Ffile: __ccgo_ts,
		Fline: int32(1474),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(14.073086670185601),
		Fy:    libc.Float64FromFloat64(646895.140242154),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1458: {
		Ffile: __ccgo_ts,
		Fline: int32(1475),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(14.073086670185601),
		Fy:    libc.Float64FromFloat64(646895.1402421541),
		Fdy:   float32(2.1929699300926586e-07),
		Fe:    int32(FE_INEXACT),
	},
	1459: {
		Ffile: __ccgo_ts,
		Fline: int32(1476),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(14.073086670185601),
		Fy:    libc.Float64FromFloat64(646895.140242154),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1460: {
		Ffile: __ccgo_ts,
		Fline: int32(1477),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.999864005519551)),
		Fy:    float64(-libc.Float64FromFloat64(201.68572702235355)),
		Fdy:   float32(0.4999993145465851),
		Fe:    int32(FE_INEXACT),
	},
	1461: {
		Ffile: __ccgo_ts,
		Fline: int32(1478),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.05789935929915894),
		Fy:    libc.Float64FromFloat64(0.057931714404507645),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1462: {
		Ffile: __ccgo_ts,
		Fline: int32(1479),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.10940960685926365)),
		Fy:    float64(-libc.Float64FromFloat64(0.10962801813403941)),
		Fdy:   float32(0.49999991059303284),
		Fe:    int32(FE_INEXACT),
	},
	1463: {
		Ffile: __ccgo_ts,
		Fline: int32(1480),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(12.250489918343725)),
		Fy:    float64(-libc.Float64FromFloat64(104541.84885789081)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999993145465851)),
		Fe:    int32(FE_INEXACT),
	},
	1464: {
		Ffile: __ccgo_ts,
		Fline: int32(1481),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01324516094255626),
		Fy:    libc.Float64FromFloat64(0.013245548221850784),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1465: {
		Ffile: __ccgo_ts,
		Fline: int32(1482),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.10880405754468958),
		Fy:    libc.Float64FromFloat64(0.10901886091255766),
		Fdy:   float32(-libc.Float64FromFloat64(8.313147077387839e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1466: {
		Ffile: __ccgo_ts,
		Fline: int32(1483),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.10880405754468958),
		Fy:    libc.Float64FromFloat64(0.10901886091255768),
		Fdy:   float32(0.9999991655349731),
		Fe:    int32(FE_INEXACT),
	},
	1467: {
		Ffile: __ccgo_ts,
		Fline: int32(1484),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.10880405754468958),
		Fy:    libc.Float64FromFloat64(0.10901886091255766),
		Fdy:   float32(-libc.Float64FromFloat64(8.313147077387839e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1468: {
		Ffile: __ccgo_ts,
		Fline: int32(1485),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.011583338078854298),
		Fy:    libc.Float64FromFloat64(0.011583597110520755),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992847442627)),
		Fe:    int32(FE_INEXACT),
	},
	1469: {
		Ffile: __ccgo_ts,
		Fline: int32(1486),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.729149439057955)),
		Fy:    float64(-libc.Float64FromFloat64(56.59520227858825)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999993443489075)),
		Fe:    int32(FE_INEXACT),
	},
	1470: {
		Ffile: __ccgo_ts,
		Fline: int32(1487),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.729149439057955)),
		Fy:    float64(-libc.Float64FromFloat64(56.595202278588246)),
		Fdy:   float32(6.576256055268459e-07),
		Fe:    int32(FE_INEXACT),
	},
	1471: {
		Ffile: __ccgo_ts,
		Fline: int32(1488),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(4.729149439057955)),
		Fy:    float64(-libc.Float64FromFloat64(56.595202278588246)),
		Fdy:   float32(6.576256055268459e-07),
		Fe:    int32(FE_INEXACT),
	},
	1472: {
		Ffile: __ccgo_ts,
		Fline: int32(1489),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.08483368053531662)),
		Fy:    float64(-libc.Float64FromFloat64(0.08493547166917344)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1473: {
		Ffile: __ccgo_ts,
		Fline: int32(1490),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.08483368053531662)),
		Fy:    float64(-libc.Float64FromFloat64(0.08493547166917342)),
		Fdy:   float32(4.612694226580061e-08),
		Fe:    int32(FE_INEXACT),
	},
	1474: {
		Ffile: __ccgo_ts,
		Fline: int32(1491),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.08483368053531662)),
		Fy:    float64(-libc.Float64FromFloat64(0.08493547166917342)),
		Fdy:   float32(4.612694226580061e-08),
		Fe:    int32(FE_INEXACT),
	},
	1475: {
		Ffile: __ccgo_ts,
		Fline: int32(1492),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.055411044053062),
		Fy:    libc.Float64FromFloat64(78.43131120819022),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999991059303284)),
		Fe:    int32(FE_INEXACT),
	},
	1476: {
		Ffile: __ccgo_ts,
		Fline: int32(1493),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.184234717269313)),
		Fy:    float64(-libc.Float64FromFloat64(32.81400609366026)),
		Fdy:   float32(-libc.Float64FromFloat64(7.197834293037886e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1477: {
		Ffile: __ccgo_ts,
		Fline: int32(1494),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.184234717269313)),
		Fy:    float64(-libc.Float64FromFloat64(32.814006093660254)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1478: {
		Ffile: __ccgo_ts,
		Fline: int32(1495),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(4.184234717269313)),
		Fy:    float64(-libc.Float64FromFloat64(32.814006093660254)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1479: {
		Ffile: __ccgo_ts,
		Fline: int32(1496),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(10.896414479513387),
		Fy:    libc.Float64FromFloat64(26991.230572278593),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992251396179)),
		Fe:    int32(FE_INEXACT),
	},
	1480: {
		Ffile: __ccgo_ts,
		Fline: int32(1497),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.007933874448787976),
		Fy:    libc.Float64FromFloat64(0.0079339576838078),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	1481: {
		Ffile: __ccgo_ts,
		Fline: int32(1498),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(45.27703862940457),
		Fy:    libc.Float64FromFloat64(2.304294790599182e+19),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999961256980896)),
		Fe:    int32(FE_INEXACT),
	},
	1482: {
		Ffile: __ccgo_ts,
		Fline: int32(1499),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3511492270210299)),
		Fy:    float64(-libc.Float64FromFloat64(0.35841030426302595)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999991655349731)),
		Fe:    int32(FE_INEXACT),
	},
	1483: {
		Ffile: __ccgo_ts,
		Fline: int32(1500),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.3511492270210299)),
		Fy:    float64(-libc.Float64FromFloat64(0.3584103042630259)),
		Fdy:   float32(8.314930255437503e-07),
		Fe:    int32(FE_INEXACT),
	},
	1484: {
		Ffile: __ccgo_ts,
		Fline: int32(1501),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.3511492270210299)),
		Fy:    float64(-libc.Float64FromFloat64(0.3584103042630259)),
		Fdy:   float32(8.314930255437503e-07),
		Fe:    int32(FE_INEXACT),
	},
	1485: {
		Ffile: __ccgo_ts,
		Fline: int32(1502),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.03461413807766779),
		Fy:    libc.Float64FromFloat64(0.034621050580657294),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999985098838806)),
		Fe:    int32(FE_INEXACT),
	},
	1486: {
		Ffile: __ccgo_ts,
		Fline: int32(1503),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.0032883107085060173),
		Fy:    libc.Float64FromFloat64(0.003288316634586218),
		Fdy:   float32(-libc.Float64FromFloat64(7.880376529101341e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1487: {
		Ffile: __ccgo_ts,
		Fline: int32(1504),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.0032883107085060173),
		Fy:    libc.Float64FromFloat64(0.0032883166345862185),
		Fdy:   float32(0.9999992251396179),
		Fe:    int32(FE_INEXACT),
	},
	1488: {
		Ffile: __ccgo_ts,
		Fline: int32(1505),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0032883107085060173),
		Fy:    libc.Float64FromFloat64(0.003288316634586218),
		Fdy:   float32(-libc.Float64FromFloat64(7.880376529101341e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1489: {
		Ffile: __ccgo_ts,
		Fline: int32(1506),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(45.93037514362743)),
		Fy:    float64(-libc.Float64FromFloat64(4.428722625905879e+19)),
		Fdy:   float32(-libc.Float64FromFloat64(4.933423269903869e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1490: {
		Ffile: __ccgo_ts,
		Fline: int32(1507),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(45.93037514362743)),
		Fy:    float64(-libc.Float64FromFloat64(4.428722625905878e+19)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1491: {
		Ffile: __ccgo_ts,
		Fline: int32(1508),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(45.93037514362743)),
		Fy:    float64(-libc.Float64FromFloat64(4.428722625905878e+19)),
		Fdy:   float32(0.9999995231628418),
		Fe:    int32(FE_INEXACT),
	},
	1492: {
		Ffile: __ccgo_ts,
		Fline: int32(1509),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(10.565391348018991),
		Fy:    libc.Float64FromFloat64(19384.792905116996),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999995231628418)),
		Fe:    int32(FE_INEXACT),
	},
	1493: {
		Ffile: __ccgo_ts,
		Fline: int32(1510),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(58.205987758249414),
		Fy:    libc.Float64FromFloat64(9.495313355059902e+24),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1494: {
		Ffile: __ccgo_ts,
		Fline: int32(1511),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(58.205987758249414),
		Fy:    libc.Float64FromFloat64(9.495313355059903e+24),
		Fdy:   float32(2.1612920875213604e-07),
		Fe:    int32(FE_INEXACT),
	},
	1495: {
		Ffile: __ccgo_ts,
		Fline: int32(1512),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(58.205987758249414),
		Fy:    libc.Float64FromFloat64(9.495313355059902e+24),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999997615814209)),
		Fe:    int32(FE_INEXACT),
	},
	1496: {
		Ffile: __ccgo_ts,
		Fline: int32(1513),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.29953978320113245),
		Fy:    libc.Float64FromFloat64(0.30403924333407895),
		Fdy:   float32(0.49999913573265076),
		Fe:    int32(FE_INEXACT),
	},
	1497: {
		Ffile: __ccgo_ts,
		Fline: int32(1514),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.022060479076481548)),
		Fy:    float64(-libc.Float64FromFloat64(0.022062268462897508)),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999988079071045)),
		Fe:    int32(FE_INEXACT),
	},
	1498: {
		Ffile: __ccgo_ts,
		Fline: int32(1515),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.773924597093861),
		Fy:    libc.Float64FromFloat64(160.89753916567253),
		Fdy:   float32(0.49999937415122986),
		Fe:    int32(FE_INEXACT),
	},
	1499: {
		Ffile: __ccgo_ts,
		Fline: int32(1516),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.5067206006975318)),
		Fy:    float64(-libc.Float64FromFloat64(16.654381050593173)),
		Fdy:   float32(-libc.Float64FromFloat64(7.030915298855689e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1500: {
		Ffile: __ccgo_ts,
		Fline: int32(1517),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.5067206006975318)),
		Fy:    float64(-libc.Float64FromFloat64(16.65438105059317)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1501: {
		Ffile: __ccgo_ts,
		Fline: int32(1518),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(3.5067206006975318)),
		Fy:    float64(-libc.Float64FromFloat64(16.65438105059317)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1502: {
		Ffile: __ccgo_ts,
		Fline: int32(1519),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.5197331375018729)),
		Fy:    float64(-libc.Float64FromFloat64(0.5434498075530381)),
		Fdy:   float32(-libc.Float64FromFloat64(1.9808783235930605e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1503: {
		Ffile: __ccgo_ts,
		Fline: int32(1520),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.5197331375018729)),
		Fy:    float64(-libc.Float64FromFloat64(0.543449807553038)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1504: {
		Ffile: __ccgo_ts,
		Fline: int32(1521),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.5197331375018729)),
		Fy:    float64(-libc.Float64FromFloat64(0.543449807553038)),
		Fdy:   float32(0.9999998211860657),
		Fe:    int32(FE_INEXACT),
	},
	1505: {
		Ffile: __ccgo_ts,
		Fline: int32(1522),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.425995113908073),
		Fy:    libc.Float64FromFloat64(5.612545952910487),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1506: {
		Ffile: __ccgo_ts,
		Fline: int32(1523),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.425995113908073),
		Fy:    libc.Float64FromFloat64(5.612545952910488),
		Fdy:   float32(3.316776258088794e-07),
		Fe:    int32(FE_INEXACT),
	},
	1507: {
		Ffile: __ccgo_ts,
		Fline: int32(1524),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.425995113908073),
		Fy:    libc.Float64FromFloat64(5.612545952910487),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1508: {
		Ffile: __ccgo_ts,
		Fline: int32(1525),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.8210233391691695)),
		Fy:    float64(-libc.Float64FromFloat64(168.65705113214443)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999999403953552)),
		Fe:    int32(FE_INEXACT),
	},
	1509: {
		Ffile: __ccgo_ts,
		Fline: int32(1526),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.985798193342399)),
		Fy:    float64(-libc.Float64FromFloat64(26.904830033097028)),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	1510: {
		Ffile: __ccgo_ts,
		Fline: int32(1527),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(360.5775132653414)),
		Fy:    float64(-libc.Float64FromFloat64(1.9760335495728697e+156)),
		Fdy:   float32(0.49999916553497314),
		Fe:    int32(FE_INEXACT),
	},
	1511: {
		Ffile: __ccgo_ts,
		Fline: int32(1528),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.0466577638100714)),
		Fy:    float64(-libc.Float64FromFloat64(3.806408198893866)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999992549419403)),
		Fe:    int32(FE_INEXACT),
	},
	1512: {
		Ffile: __ccgo_ts,
		Fline: int32(1529),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.5655063792806123)),
		Fy:    float64(-libc.Float64FromFloat64(6.46518152988839)),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999990463256836)),
		Fe:    int32(FE_INEXACT),
	},
	1513: {
		Ffile: __ccgo_ts,
		Fline: int32(1530),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.5655063792806123)),
		Fy:    float64(-libc.Float64FromFloat64(6.4651815298883895)),
		Fdy:   float32(9.476165132582537e-07),
		Fe:    int32(FE_INEXACT),
	},
	1514: {
		Ffile: __ccgo_ts,
		Fline: int32(1531),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.5655063792806123)),
		Fy:    float64(-libc.Float64FromFloat64(6.4651815298883895)),
		Fdy:   float32(9.476165132582537e-07),
		Fe:    int32(FE_INEXACT),
	},
	1515: {
		Ffile: __ccgo_ts,
		Fline: int32(1532),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(11.27897568112739),
		Fy:    libc.Float64FromFloat64(39570.07779581674),
		Fdy:   float32(-libc.Float64FromFloat64(0.49999934434890747)),
		Fe:    int32(FE_INEXACT),
	},
	1516: {
		Ffile: __ccgo_ts,
		Fline: int32(1533),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(48.824495420354246)),
		Fy:    float64(-libc.Float64FromFloat64(8.001638873714679e+20)),
		Fdy:   float32(-libc.Float64FromFloat64(6.635634690610459e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1517: {
		Ffile: __ccgo_ts,
		Fline: int32(1534),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(48.824495420354246)),
		Fy:    float64(-libc.Float64FromFloat64(8.001638873714678e+20)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1518: {
		Ffile: __ccgo_ts,
		Fline: int32(1535),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(48.824495420354246)),
		Fy:    float64(-libc.Float64FromFloat64(8.001638873714678e+20)),
		Fdy:   float32(0.9999993443489075),
		Fe:    int32(FE_INEXACT),
	},
	1519: {
		Ffile: __ccgo_ts,
		Fline: int32(1536),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(22.191925596998523),
		Fy:    libc.Float64FromFloat64(2.1717050133906517e+09),
		Fdy:   float32(-libc.Float64FromFloat64(5.48562411495368e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1520: {
		Ffile: __ccgo_ts,
		Fline: int32(1537),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(22.191925596998523),
		Fy:    libc.Float64FromFloat64(2.171705013390652e+09),
		Fdy:   float32(0.999999463558197),
		Fe:    int32(FE_INEXACT),
	},
	1521: {
		Ffile: __ccgo_ts,
		Fline: int32(1538),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(22.191925596998523),
		Fy:    libc.Float64FromFloat64(2.1717050133906517e+09),
		Fdy:   float32(-libc.Float64FromFloat64(5.48562411495368e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1522: {
		Ffile: __ccgo_ts,
		Fline: int32(1539),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(182.41059344582766)),
		Fy:    float64(-libc.Float64FromFloat64(8.296294851757244e+78)),
		Fdy:   float32(0.4999995827674866),
		Fe:    int32(FE_INEXACT),
	},
	1523: {
		Ffile: __ccgo_ts,
		Fline: int32(1540),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0.003920879904025643),
		Fy:    libc.Float64FromFloat64(0.003920889950176695),
		Fdy:   float32(-libc.Float64FromFloat64(2.5241128298603144e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1524: {
		Ffile: __ccgo_ts,
		Fline: int32(1541),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0.003920879904025643),
		Fy:    libc.Float64FromFloat64(0.0039208899501766955),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1525: {
		Ffile: __ccgo_ts,
		Fline: int32(1542),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.003920879904025643),
		Fy:    libc.Float64FromFloat64(0.003920889950176695),
		Fdy:   float32(-libc.Float64FromFloat64(2.5241128298603144e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1526: {
		Ffile: __ccgo_ts,
		Fline: int32(1543),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(47.2876134512479)),
		Fy:    float64(-libc.Float64FromFloat64(1.7207571729797448e+20)),
		Fdy:   float32(-libc.Float64FromFloat64(7.132177159974162e-08)),
		Fe:    int32(FE_INEXACT),
	},
	1527: {
		Ffile: __ccgo_ts,
		Fline: int32(1544),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(47.2876134512479)),
		Fy:    float64(-libc.Float64FromFloat64(1.7207571729797444e+20)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1528: {
		Ffile: __ccgo_ts,
		Fline: int32(1545),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(47.2876134512479)),
		Fy:    float64(-libc.Float64FromFloat64(1.7207571729797444e+20)),
		Fdy:   float32(0.9999999403953552),
		Fe:    int32(FE_INEXACT),
	},
	1529: {
		Ffile: __ccgo_ts,
		Fline: int32(1546),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(162.56712752196054),
		Fy:    libc.Float64FromFloat64(1.9997533177291967e+70),
		Fdy:   float32(0.4999995529651642),
		Fe:    int32(FE_INEXACT),
	},
	1530: {
		Ffile: __ccgo_ts,
		Fline: int32(1547),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.8119960348157726),
		Fy:    libc.Float64FromFloat64(8.291510181562067),
		Fdy:   float32(-libc.Float64FromFloat64(3.6491147170636395e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1531: {
		Ffile: __ccgo_ts,
		Fline: int32(1548),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.8119960348157726),
		Fy:    libc.Float64FromFloat64(8.291510181562069),
		Fdy:   float32(0.9999996423721313),
		Fe:    int32(FE_INEXACT),
	},
	1532: {
		Ffile: __ccgo_ts,
		Fline: int32(1549),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.8119960348157726),
		Fy:    libc.Float64FromFloat64(8.291510181562067),
		Fdy:   float32(-libc.Float64FromFloat64(3.6491147170636395e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1533: {
		Ffile: __ccgo_ts,
		Fline: int32(1550),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.035616792023675424)),
		Fy:    float64(-libc.Float64FromFloat64(0.03562432281644391)),
		Fdy:   float32(-libc.Float64FromFloat64(6.959020879548916e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1534: {
		Ffile: __ccgo_ts,
		Fline: int32(1551),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.035616792023675424)),
		Fy:    float64(-libc.Float64FromFloat64(0.035624322816443905)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1535: {
		Ffile: __ccgo_ts,
		Fline: int32(1552),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.035616792023675424)),
		Fy:    float64(-libc.Float64FromFloat64(0.035624322816443905)),
		Fdy:   float32(0.9999992847442627),
		Fe:    int32(FE_INEXACT),
	},
	1536: {
		Ffile: __ccgo_ts,
		Fline: int32(1553),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.1008543484713104)),
		Fy:    float64(-libc.Float64FromFloat64(0.1010254104553604)),
		Fdy:   float32(-libc.Float64FromFloat64(6.062239776838396e-07)),
		Fe:    int32(FE_INEXACT),
	},
	1537: {
		Ffile: __ccgo_ts,
		Fline: int32(1554),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0.1008543484713104)),
		Fy:    float64(-libc.Float64FromFloat64(0.10102541045536038)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1538: {
		Ffile: __ccgo_ts,
		Fline: int32(1555),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0.1008543484713104)),
		Fy:    float64(-libc.Float64FromFloat64(0.10102541045536038)),
		Fdy:   float32(0.9999994039535522),
		Fe:    int32(FE_INEXACT),
	},
	1539: {
		Ffile: __ccgo_ts,
		Fline: int32(1556),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(53.74555411975277),
		Fy:    libc.Float64FromFloat64(1.0974066444336659e+23),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1540: {
		Ffile: __ccgo_ts,
		Fline: int32(1557),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(53.74555411975277),
		Fy:    libc.Float64FromFloat64(1.097406644433666e+23),
		Fdy:   float32(3.316703498512652e-07),
		Fe:    int32(FE_INEXACT),
	},
	1541: {
		Ffile: __ccgo_ts,
		Fline: int32(1558),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(53.74555411975277),
		Fy:    libc.Float64FromFloat64(1.0974066444336659e+23),
		Fdy:   float32(-libc.Float64FromFloat64(0.9999996423721313)),
		Fe:    int32(FE_INEXACT),
	},
	1542: {
		Ffile: __ccgo_ts,
		Fline: int32(1559),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.011301030201687476),
		Fy:    libc.Float64FromFloat64(0.011301270751836099),
		Fdy:   float32(-libc.Float64FromFloat64(0.4999997317790985)),
		Fe:    int32(FE_INEXACT),
	},
	1543: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(38),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(710.4758600739439),
		Fy:    libc.Float64FromFloat64(1.7976931348621744e+308),
		Fdy:   float32(0.10319651663303375),
		Fe:    int32(FE_INEXACT),
	},
	1544: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(710.4758600739439)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348621744e+308)),
		Fdy:   float32(-libc.Float64FromFloat64(0.10319651663303375)),
		Fe:    int32(FE_INEXACT),
	},
	1545: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(40),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(710.4758600739439),
		Fy:    libc.Float64FromFloat64(1.7976931348621742e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.8968034982681274)),
		Fe:    int32(FE_INEXACT),
	},
	1546: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(41),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(710.4758600739439)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348621742e+308)),
		Fdy:   float32(0.8968034982681274),
		Fe:    int32(FE_INEXACT),
	},
	1547: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(42),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(710.4758600739439),
		Fy:    libc.Float64FromFloat64(1.7976931348621744e+308),
		Fdy:   float32(0.10319651663303375),
		Fe:    int32(FE_INEXACT),
	},
	1548: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(43),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.4758600739439)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348621742e+308)),
		Fdy:   float32(0.8968034982681274),
		Fe:    int32(FE_INEXACT),
	},
	1549: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(44),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(710.4758600739439),
		Fy:    libc.Float64FromFloat64(1.7976931348621742e+308),
		Fdy:   float32(-libc.Float64FromFloat64(0.8968034982681274)),
		Fe:    int32(FE_INEXACT),
	},
	1550: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(45),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.4758600739439)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348621744e+308)),
		Fdy:   float32(-libc.Float64FromFloat64(0.10319651663303375)),
		Fe:    int32(FE_INEXACT),
	},
	1551: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(46),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(710.475860073944),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1552: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(47),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(710.475860073944)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1553: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(48),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(710.475860073944),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1554: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(49),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(710.475860073944)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1555: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(50),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(710.475860073944),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1556: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(51),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.475860073944)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1557: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(52),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(710.475860073944),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1558: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(53),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.475860073944)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1559: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.776263578034403e-21),
		Fy:    libc.Float64FromFloat64(6.776263578034403e-21),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1560: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(56),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(6.776263578034403e-21)),
		Fy:    float64(-libc.Float64FromFloat64(6.776263578034403e-21)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1561: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1562: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1563: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.121995791e-314),
		Fy:    libc.Float64FromFloat64(2.121995791e-314),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1564: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.121995791e-314)),
		Fy:    float64(-libc.Float64FromFloat64(2.121995791e-314)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1565: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(61),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1566: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(62),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1567: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1568: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(65),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1569: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(66),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1570: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(67),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1571: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(68),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1572: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(69),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1573: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(70),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1574: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(71),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1575: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(73),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.061254105181627)),
		Fy:    float64(-libc.Float64FromFloat64(1584.6309578891396)),
		Fdy:   float32(-libc.Float64FromFloat64(0.23715293407440186)),
		Fe:    int32(FE_INEXACT),
	},
	1576: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(74),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(21.366775245021078),
		Fy:    libc.Float64FromFloat64(9.515740483834388e+08),
		Fdy:   float32(0.4591154158115387),
		Fe:    int32(FE_INEXACT),
	},
	1577: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(75),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(26.413498513041322),
		Fy:    libc.Float64FromFloat64(1.4798123801141684e+11),
		Fdy:   float32(-libc.Float64FromFloat64(0.18828228116035461)),
		Fe:    int32(FE_INEXACT),
	},
	1578: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(76),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(20.03565953145929)),
		Fy:    float64(-libc.Float64FromFloat64(2.5138906352810937e+08)),
		Fdy:   float32(0.06728991121053696),
		Fe:    int32(FE_INEXACT),
	},
	1579: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(77),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.127097202525357)),
		Fy:    float64(-libc.Float64FromFloat64(1692.4795714846805)),
		Fdy:   float32(-libc.Float64FromFloat64(0.021291883662343025)),
		Fe:    int32(FE_INEXACT),
	},
	1580: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.1787587219330506)),
		Fy:    float64(-libc.Float64FromFloat64(1.4713383659375732)),
		Fdy:   float32(0.414265513420105),
		Fe:    int32(FE_INEXACT),
	},
	1581: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(79),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(17.349942297752893),
		Fy:    libc.Float64FromFloat64(1.7137765879068326e+07),
		Fdy:   float32(0.10599929094314575),
		Fe:    int32(FE_INEXACT),
	},
	1582: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(80),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(25.652653355546338),
		Fy:    libc.Float64FromFloat64(6.914739170629799e+10),
		Fdy:   float32(-libc.Float64FromFloat64(0.4773549437522888)),
		Fe:    int32(FE_INEXACT),
	},
	1583: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(81),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.042064912260298)),
		Fy:    float64(-libc.Float64FromFloat64(77.3914256028152)),
		Fdy:   float32(-libc.Float64FromFloat64(0.2838144302368164)),
		Fe:    int32(FE_INEXACT),
	},
	1584: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(28.740725577743834)),
		Fy:    float64(-libc.Float64FromFloat64(1.5167308790828335e+12)),
		Fdy:   float32(0.015985963866114616),
		Fe:    int32(FE_INEXACT),
	},
	1585: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(84),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1586: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1587: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1588: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(87),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1589: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(88),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1590: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(89),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1591: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(90),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1592: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(91),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1593: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(92),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1594: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(93),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1595: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(94),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fy:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.1666666716337204)),
		Fe:    int32(FE_INEXACT),
	},
	1596: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(95),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2351741790771484e-08),
		Fy:    libc.Float64FromFloat64(2.2351741790771484e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5625)),
		Fe:    int32(FE_INEXACT),
	},
	1597: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(96),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1598: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(97),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1599: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(98),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507203e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1600: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(99),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fy:    float64(-libc.Float64FromFloat64(4.450147717014404e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1601: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(100),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fy:    float64(-libc.Float64FromFloat64(8.900295434028808e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	1602: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(101),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1603: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(102),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1e-323),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1604: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(103),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1605: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(104),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1606: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(105),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1607: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(106),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1608: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(107),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(1e-323)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1609: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(108),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1e-323)),
		Fy:    float64(-libc.Float64FromFloat64(1.5e-323)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1610: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(109),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fy:    float64(-libc.Float64FromFloat64(5.56268464626801e-309)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1611: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(110),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fy:    float64(-libc.Float64FromFloat64(1.112536929253601e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1612: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(111),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1613: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(112),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1614: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(113),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(710.5),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1615: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(114),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1616: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(115),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1617: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(116),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1618: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(117),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1619: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(118),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.5)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1620: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(119),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.49423283715579e+307)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1621: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(120),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.98846567431158e+307)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1622: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(121),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623155e+308)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1623: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(122),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1624: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(123),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1625: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(124),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1626: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(125),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fy:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.6666666865348816)),
		Fe:    int32(FE_INEXACT),
	},
	1627: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(126),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(3.725290298461914e-08),
		Fy:    libc.Float64FromFloat64(3.725290298461915e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.3020833432674408)),
		Fe:    int32(FE_INEXACT),
	},
	1628: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(127),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fy:    float64(-libc.Float64FromFloat64(1.490116119384766e-08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.8333333134651184)),
		Fe:    int32(FE_INEXACT),
	},
	1629: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(128),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2351741790771484e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.2351741790771488e-08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.4375)),
		Fe:    int32(FE_INEXACT),
	},
	1630: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(129),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.980232238769532e-08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.3333333432674408)),
		Fe:    int32(FE_INEXACT),
	},
	1631: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(130),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.725290298461914e-08)),
		Fy:    float64(-libc.Float64FromFloat64(3.7252902984619154e-08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.6979166865348816)),
		Fe:    int32(FE_INEXACT),
	},
	1632: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1633: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(132),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1634: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1635: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1636: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(135),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fy:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.1666666716337204)),
		Fe:    int32(FE_INEXACT),
	},
	1637: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(136),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fy:    libc.Float64FromFloat64(2.980232238769532e-08),
		Fdy:   float32(0.3333333432674408),
		Fe:    int32(FE_INEXACT),
	},
	1638: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(137),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.725290298461914e-08),
		Fy:    libc.Float64FromFloat64(3.725290298461915e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.3020833432674408)),
		Fe:    int32(FE_INEXACT),
	},
	1639: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(138),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1640: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(139),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1641: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(140),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fy:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1642: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(141),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fy:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1643: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(142),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fy:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fdy:   float32(0.1666666716337204),
		Fe:    int32(FE_INEXACT),
	},
	1644: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(143),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.980232238769532e-08)),
		Fdy:   float32(-libc.Float64FromFloat64(0.3333333432674408)),
		Fe:    int32(FE_INEXACT),
	},
	1645: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(144),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(3.725290298461914e-08)),
		Fy:    float64(-libc.Float64FromFloat64(3.725290298461915e-08)),
		Fdy:   float32(0.3020833432674408),
		Fe:    int32(FE_INEXACT),
	},
	1646: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(145),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1e-323),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1647: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(146),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1648: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(147),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1649: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(148),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1650: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(149),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1651: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(150),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1e-323)),
		Fy:    float64(-libc.Float64FromFloat64(1e-323)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1652: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(151),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fy:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1653: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(152),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fy:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1654: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(153),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1655: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1656: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(710.5),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1657: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(156),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1658: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1659: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(158),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1660: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1661: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(710.5)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1662: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(161),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(4.49423283715579e+307)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1663: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(162),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.98846567431158e+307)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1664: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623155e+308)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1665: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(164),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1666: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1667: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(166),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1668: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(167),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1669: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(168),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2351741790771484e-08),
		Fy:    libc.Float64FromFloat64(2.2351741790771488e-08),
		Fdy:   float32(0.4375),
		Fe:    int32(FE_INEXACT),
	},
	1670: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(169),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1671: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(170),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1672: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(171),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1673: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(172),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fy:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1674: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(173),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fy:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1675: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(174),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1676: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(175),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1e-323)),
		Fy:    float64(-libc.Float64FromFloat64(1e-323)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1677: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(176),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fy:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1678: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(177),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fy:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1679: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(178),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1680: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(179),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1681: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(180),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(710.5),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1682: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(181),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1683: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(182),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1684: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(183),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1685: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(184),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1686: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(185),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(710.5)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1687: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(186),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(4.49423283715579e+307)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1688: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(187),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(8.98846567431158e+307)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1689: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(188),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623155e+308)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1690: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(189),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1691: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(190),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1692: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(191),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1693: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(192),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1694: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(193),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1695: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(194),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507203e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1696: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(195),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014404e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1697: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(196),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(8.900295434028808e-308),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	1698: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(197),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fy:    libc.Float64FromFloat64(1.490116119384766e-08),
		Fdy:   float32(0.8333333134651184),
		Fe:    int32(FE_INEXACT),
	},
	1699: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(198),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fy:    libc.Float64FromFloat64(2.980232238769532e-08),
		Fdy:   float32(0.3333333432674408),
		Fe:    int32(FE_INEXACT),
	},
	1700: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(199),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(3.725290298461914e-08),
		Fy:    libc.Float64FromFloat64(3.7252902984619154e-08),
		Fdy:   float32(0.6979166865348816),
		Fe:    int32(FE_INEXACT),
	},
	1701: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(200),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fy:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fdy:   float32(0.1666666716337204),
		Fe:    int32(FE_INEXACT),
	},
	1702: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(201),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.2351741790771484e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.2351741790771484e-08)),
		Fdy:   float32(0.5625),
		Fe:    int32(FE_INEXACT),
	},
	1703: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(202),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fdy:   float32(0.6666666865348816),
		Fe:    int32(FE_INEXACT),
	},
	1704: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(203),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(-libc.Float64FromFloat64(3.725290298461914e-08)),
		Fy:    float64(-libc.Float64FromFloat64(3.725290298461915e-08)),
		Fdy:   float32(0.3020833432674408),
		Fe:    int32(FE_INEXACT),
	},
	1705: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(204),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1706: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(205),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1e-323),
		Fy:    libc.Float64FromFloat64(1.5e-323),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1707: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(206),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(5.56268464626801e-309),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1708: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(207),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    libc.Float64FromFloat64(1.112536929253601e-308),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1709: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(208),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1710: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(209),
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1711: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(210),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1712: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(211),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1713: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(212),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072014e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1714: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(213),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507202e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1715: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(214),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072024e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1716: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(215),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fy:    libc.Float64FromFloat64(4.450147717014403e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1717: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(216),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fy:    libc.Float64FromFloat64(8.900295434028806e-308),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1718: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fy:    libc.Float64FromFloat64(1.4901161193847656e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.1666666716337204)),
		Fe:    int32(FE_INEXACT),
	},
	1719: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(218),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2351741790771484e-08),
		Fy:    libc.Float64FromFloat64(2.2351741790771484e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5625)),
		Fe:    int32(FE_INEXACT),
	},
	1720: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(219),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072014e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1721: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507202e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1722: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(221),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072024e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1723: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(222),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fy:    float64(-libc.Float64FromFloat64(4.450147717014403e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1724: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(223),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fy:    float64(-libc.Float64FromFloat64(8.900295434028806e-308)),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	1725: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(224),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1726: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(225),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1e-323),
		Fy:    libc.Float64FromFloat64(1e-323),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1727: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(226),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fy:    libc.Float64FromFloat64(5.562684646268003e-309),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1728: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(227),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fy:    libc.Float64FromFloat64(1.1125369292536007e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1729: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(228),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fy:    libc.Float64FromFloat64(2.2250738585072004e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1730: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(229),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fy:    libc.Float64FromFloat64(2.225073858507201e-308),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1731: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(230),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(5e-324)),
		Fy:    float64(-libc.Float64FromFloat64(5e-324)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1732: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(231),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1e-323)),
		Fy:    float64(-libc.Float64FromFloat64(1e-323)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1733: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(232),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fy:    float64(-libc.Float64FromFloat64(5.562684646268003e-309)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1734: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fy:    float64(-libc.Float64FromFloat64(1.1125369292536007e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1735: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.2250738585072004e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1736: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(235),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fy:    float64(-libc.Float64FromFloat64(2.225073858507201e-308)),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	1737: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(236),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(710.5),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1738: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(237),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.49423283715579e+307),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1739: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.98846567431158e+307),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1740: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7976931348623155e+308),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1741: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(240),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fy:    libc.Float64FromFloat64(1.7976931348623157e+308),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1742: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(241),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(710.5)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1743: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(242),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(4.49423283715579e+307)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1744: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(8.98846567431158e+307)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1745: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(244),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623155e+308)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1746: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fy:    float64(-libc.Float64FromFloat64(1.7976931348623157e+308)),
		Fdy:   float32(1),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_OVERFLOW),
	},
	1747: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1748: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1749: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fy:    libc.Float64FromFloat64(2.9802322387695312e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.6666666865348816)),
		Fe:    int32(FE_INEXACT),
	},
	1750: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.725290298461914e-08),
		Fy:    libc.Float64FromFloat64(3.725290298461915e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.3020833432674408)),
		Fe:    int32(FE_INEXACT),
	},
	1751: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fy:    float64(-libc.Float64FromFloat64(1.4901161193847656e-08)),
		Fdy:   float32(0.1666666716337204),
		Fe:    int32(FE_INEXACT),
	},
	1752: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.2351741790771484e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.2351741790771484e-08)),
		Fdy:   float32(0.5625),
		Fe:    int32(FE_INEXACT),
	},
	1753: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fy:    float64(-libc.Float64FromFloat64(2.9802322387695312e-08)),
		Fdy:   float32(0.6666666865348816),
		Fe:    int32(FE_INEXACT),
	},
	1754: {
		Ffile: __ccgo_ts + 24,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(-libc.Float64FromFloat64(3.725290298461914e-08)),
		Fy:    float64(-libc.Float64FromFloat64(3.725290298461915e-08)),
		Fdy:   float32(0.3020833432674408),
		Fe:    int32(FE_INEXACT),
	},
	1755: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.06684839057968)),
		Fy:    float64(-libc.Float64FromFloat64(1593.5206801156262)),
		Fdy:   float32(-libc.Float64FromFloat64(0.2138727605342865)),
		Fe:    int32(FE_INEXACT),
	},
	1756: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fy:    libc.Float64FromFloat64(38.54878088685412),
		Fdy:   float32(0.21537430584430695),
		Fe:    int32(FE_INEXACT),
	},
	1757: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(8.38143342755525)),
		Fy:    float64(-libc.Float64FromFloat64(2182.6307505145546)),
		Fdy:   float32(0.16213826835155487),
		Fe:    int32(FE_INEXACT),
	},
	1758: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(6.531673581913484)),
		Fy:    float64(-libc.Float64FromFloat64(343.2723926847529)),
		Fdy:   float32(0.20479513704776764),
		Fe:    int32(FE_INEXACT),
	},
	1759: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fy:    libc.Float64FromFloat64(5291.7790755194055),
		Fdy:   float32(-libc.Float64FromFloat64(0.48676517605781555)),
		Fe:    int32(FE_INEXACT),
	},
	1760: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619858980995045),
		Fy:    libc.Float64FromFloat64(0.7114062568229157),
		Fdy:   float32(-libc.Float64FromFloat64(0.4584641456604004)),
		Fe:    int32(FE_INEXACT),
	},
	1761: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.4066039223853553)),
		Fy:    float64(-libc.Float64FromFloat64(0.41790065258739445)),
		Fdy:   float32(0.37220045924186707),
		Fe:    int32(FE_INEXACT),
	},
	1762: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5617597462207241),
		Fy:    libc.Float64FromFloat64(0.5917755935451237),
		Fdy:   float32(0.46178996562957764),
		Fe:    int32(FE_INEXACT),
	},
	1763: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7741522965913037),
		Fy:    libc.Float64FromFloat64(0.8538292008852542),
		Fdy:   float32(-libc.Float64FromFloat64(0.07019051909446716)),
		Fe:    int32(FE_INEXACT),
	},
	1764: {
		Ffile: __ccgo_ts + 44,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0.6787637026394024)),
		Fy:    float64(-libc.Float64FromFloat64(0.732097615653169)),
		Fdy:   float32(0.26858529448509216),
		Fe:    int32(FE_INEXACT),
	},
	1765: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1766: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.Float64FromFloat64(0)),
		Fy:    float64(-libc.Float64FromFloat64(0)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1767: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1768: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	1769: {
		Ffile: __ccgo_ts + 67,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+23)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:19:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:20:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:24:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:27:2:
	i = 0
	for {
		if !(uint32(uint32(i)) < libc.Uint32FromInt64(63720)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:28:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:30:3:
		if (*l_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:31:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:32:3:
		libc.Xfesetround(tls, (*l_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:33:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:34:3:
		y = libc.Xsinhl(tls, (*l_l)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:35:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:37:3:
		if !(checkexcept(tls, e, (*l_l)(unsafe.Pointer(p)).Fe, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:38:4:
			libc.Xprintf(tls, __ccgo_ts+91, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, estr(tls, (*l_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:40:4:
			libc.Xprintf(tls, __ccgo_ts+143, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:41:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:43:3:
		d = ulperrl(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:44:3:
		if !(checkulp(tls, d, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:45:4:
			if libc.Xfabsf(tls, d) < libc.Float32FromFloat32(5) || (*l_l)(unsafe.Pointer(p)).Fr != FE_TONEAREST {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:46:5:
				libc.Xprintf(tls, __ccgo_ts+152, 0)
			} else {
				// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:48:5:
				err++
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:49:4:
			libc.Xprintf(tls, __ccgo_ts+155, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, y, float64(float64(d)), float64(d-(*l_l)(unsafe.Pointer(p)).Fdy), float64((*l_l)(unsafe.Pointer(p)).Fdy)))
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/sinhl.c:53:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+216, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+226, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	*(*float64)(unsafe.Pointer(bp)) = float64(float64(ywant))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
	v1 = *(*uint64)(unsafe.Pointer(bp))
	goto _2
_2:
	if libc.BoolInt32(v1&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:132:3:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(y))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v3 = *(*uint64)(unsafe.Pointer(bp))
		goto _4
	_4:
		return libc.BoolInt32(v3&(-libc.Uint64FromUint64(1)>>libc.Int32FromInt32(1)) > libc.Uint64FromUint64(0x7ff)<<libc.Int32FromInt32(52))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.h:133:2:
	if v9 = y == ywant; v9 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(y))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:65:2:
		v5 = *(*uint64)(unsafe.Pointer(bp))
		goto _6
	_6:
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:64:2:
		*(*float64)(unsafe.Pointer(bp)) = float64(float64(ywant))
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
	return eulp(tls, float64(float64(x)))
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
	return float32(libc.Xscalbn(tls, float64(got-want), -eulpf(tls, want)) + float64(float64(dwant)))
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
	return float32(libc.Xscalbn(tls, got-want, -eulp(tls, want)) + float64(float64(dwant)))
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:74:7:
func ulperrl(tls *libc.TLS, got float64, want float64, dwant float32) (r float32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:75:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:77:2:
	return ulperr(tls, float64(float64(got)), float64(float64(want)), dwant)
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
		Fs:    __ccgo_ts + 270,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 278,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 286,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 296,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 306,
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
		if !(uint32(uint32(i)) < libc.Uint32FromInt64(40)/libc.Uint32FromInt64(8)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:117:3:
		if f&eflags[i].Fflag != 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:118:4:
			if all != 0 {
				v2 = __ccgo_ts + 315
			} else {
				v2 = __ccgo_ts + 23
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+317, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 315
		} else {
			v3 = __ccgo_ts + 23
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+322, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 23
	} else {
		v4 = __ccgo_ts + 327
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+329, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 332
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 335
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 338
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 341
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 344
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+347, libc.VaList(bp+8, int32(int32(s))-int32(int32(argv0)), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+355, libc.VaList(bp+8, p))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:15:2:
	if uint32(uint32(k)) >= n {
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
		if uint32(uint32(n)) >= uint32(512) {
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
	return libc.Xwrite(tls, int32(1), bp, uint32(uint32(n)))
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
		*(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8)) = uint64(uint64(i))
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
	if n < uint64(uint64(k)) {
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
			if t_randn(tls, v1) < uint64(uint64(k)) {
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
	if n < uint64(uint32(5)*k) && (n-uint64(uint64(k)))*uint64(8) < uint64(uint32(-libc.Int32FromInt32(1))) {
		/* allocation is n-k < 4*k */
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:141:3:
		tab = libc.Xmalloc(tls, uint32((n-uint64(uint64(k)))*uint64(8)))
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
			*(*uint64_t)(unsafe.Pointer(p + uintptr(i)*8)) = uint64(uint64(i))
			goto _5
		_5:
			i++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:146:3:
		for {
			if !(uint64(uint64(i)) < n) {
				break
			}
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:147:4:
			*(*uint64_t)(unsafe.Pointer(tab + uintptr(i-k)*8)) = uint64(uint64(i))
			goto _6
		_6:
			i++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:148:3:
		if uint64(uint64(k)) < n-uint64(uint64(k)) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:149:4:
			shuffle2(tls, p, tab, k, uint32(n-uint64(uint64(k))))
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/rand.c:151:4:
			shuffle2(tls, tab, p, uint32(n-uint64(uint64(k))), k)
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
		t_printf(tls, __ccgo_ts+360, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:12:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:14:2:
	if uint64(uint64(lim)) > (*(*rlimit)(unsafe.Pointer(bp))).Frlim_max {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:15:3:
		return -int32(1)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:16:2:
	if uint64(uint64(lim)) == (*(*rlimit)(unsafe.Pointer(bp))).Frlim_max && uint64(uint64(lim)) == (*(*rlimit)(unsafe.Pointer(bp))).Frlim_cur {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:17:3:
		return 0
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:18:2:
	(*(*rlimit)(unsafe.Pointer(bp))).Frlim_max = uint64(uint64(lim))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:19:2:
	(*(*rlimit)(unsafe.Pointer(bp))).Frlim_cur = uint64(uint64(lim))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:20:2:
	if libc.Xsetrlimit(tls, r, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/setrlim.c:21:3:
		t_printf(tls, __ccgo_ts+403, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+452) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+460) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+472) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+484) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+496) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+505) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+23) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+505) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+511, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/sinh.h\x00\x00src/math/ucb/sinh.h\x00src/math/sanity/sinh.h\x00src/math/special/sinh.h\x00%s:%d: bad fp exception: %s sinhl(%La)=%La, want %s\x00 got %s\n\x00X \x00%s:%d: %s sinhl(%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
