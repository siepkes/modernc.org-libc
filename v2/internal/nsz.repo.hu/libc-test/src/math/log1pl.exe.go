// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/math/log1pl.exe.go src/math/log1pl.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:5:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:5:19:
var t = [604]l_l{
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
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
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
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
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
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
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
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
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
		Fr:    int32(FE_UPWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	10: {
		Ffile: __ccgo_ts,
		Fline: int32(21),
		Fr:    int32(FE_DOWNWARD),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	11: {
		Ffile: __ccgo_ts,
		Fline: int32(22),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5e-324),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	12: {
		Ffile: __ccgo_ts,
		Fline: int32(23),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	13: {
		Ffile: __ccgo_ts,
		Fline: int32(24),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
		Fdy:   float32(0),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	14: {
		Ffile: __ccgo_ts,
		Fline: int32(25),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(1e-323),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    libc.Int32FromInt32(FE_INEXACT) | libc.Int32FromInt32(FE_UNDERFLOW),
	},
	15: {
		Ffile: __ccgo_ts,
		Fline: int32(26),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5e-324),
		Fy:    -libc.Float64FromFloat64(5e-324),
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
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
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
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
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
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
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
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	24: {
		Ffile: __ccgo_ts,
		Fline: int32(35),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	25: {
		Ffile: __ccgo_ts,
		Fline: int32(36),
		Fr:    int32(FE_UPWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	26: {
		Ffile: __ccgo_ts,
		Fline: int32(37),
		Fr:    int32(FE_DOWNWARD),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	27: {
		Ffile: __ccgo_ts,
		Fline: int32(38),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	28: {
		Ffile: __ccgo_ts,
		Fline: int32(39),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	29: {
		Ffile: __ccgo_ts,
		Fline: int32(40),
		Fr:    int32(FE_UPWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	30: {
		Ffile: __ccgo_ts,
		Fline: int32(41),
		Fr:    int32(FE_DOWNWARD),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	31: {
		Ffile: __ccgo_ts,
		Fline: int32(42),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	32: {
		Ffile: __ccgo_ts,
		Fline: int32(44),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.3322676295501873e-15),
		Fy:    libc.Float64FromFloat64(1.3322676295501865e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	33: {
		Ffile: __ccgo_ts,
		Fline: int32(49),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.726348535513887),
		Fy:    libc.Float64FromFloat64(0.5460085047865537),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	34: {
		Ffile: __ccgo_ts,
		Fline: int32(50),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2293000854655507),
		Fy:    libc.Float64FromFloat64(0.8016876731867886),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	35: {
		Ffile: __ccgo_ts,
		Fline: int32(51),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.706373313299537),
		Fy:    libc.Float64FromFloat64(0.9956094778458613),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	36: {
		Ffile: __ccgo_ts,
		Fline: int32(52),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0010741495427953741),
		Fy:    libc.Float64FromFloat64(0.00107357305695963),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	37: {
		Ffile: __ccgo_ts,
		Fline: int32(53),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0012563621481279006),
		Fy:    libc.Float64FromFloat64(0.001255573585615234),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	38: {
		Ffile: __ccgo_ts,
		Fline: int32(54),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00126529887455303),
		Fy:    libc.Float64FromFloat64(0.0012644990585318414),
		Fdy:   float32(-libc.Float64FromFloat64(9.15570618181382e-17)),
		Fe:    int32(FE_INEXACT),
	},
	39: {
		Ffile: __ccgo_ts,
		Fline: int32(55),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.001386038534474709),
		Fy:    libc.Float64FromFloat64(0.0013850788697177263),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	40: {
		Ffile: __ccgo_ts,
		Fline: int32(56),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0016168105960983048),
		Fy:    libc.Float64FromFloat64(0.0016155049649625313),
		Fdy:   float32(-libc.Float64FromFloat64(3.901381874285288e-16)),
		Fe:    int32(FE_INEXACT),
	},
	41: {
		Ffile: __ccgo_ts,
		Fline: int32(57),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.001797224520012697),
		Fy:    libc.Float64FromFloat64(0.001795611444441827),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	42: {
		Ffile: __ccgo_ts,
		Fline: int32(58),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0006762949071946357),
		Fy:    libc.Float64FromFloat64(0.000676066322848367),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	43: {
		Ffile: __ccgo_ts,
		Fline: int32(59),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0008207453525860209),
		Fy:    libc.Float64FromFloat64(0.0008204087252967217),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	44: {
		Ffile: __ccgo_ts,
		Fline: int32(60),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.000834984890151549),
		Fy:    libc.Float64FromFloat64(0.0008346364841971411),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	45: {
		Ffile: __ccgo_ts,
		Fline: int32(61),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.000358516119000968),
		Fy:    libc.Float64FromFloat64(0.0003584518674535278),
		Fdy:   float32(-libc.Float64FromFloat64(1.403694529860354e-16)),
		Fe:    int32(FE_INEXACT),
	},
	46: {
		Ffile: __ccgo_ts,
		Fline: int32(62),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00012821293410829644),
		Fy:    libc.Float64FromFloat64(0.00012820471553253774),
		Fdy:   float32(-libc.Float64FromFloat64(6.670699450004875e-17)),
		Fe:    int32(FE_INEXACT),
	},
	47: {
		Ffile: __ccgo_ts,
		Fline: int32(63),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0001629242333291465),
		Fy:    libc.Float64FromFloat64(0.00016291096261763764),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	48: {
		Ffile: __ccgo_ts,
		Fline: int32(64),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00019882570240539833),
		Fy:    libc.Float64FromFloat64(0.0001988059391950092),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	49: {
		Ffile: __ccgo_ts,
		Fline: int32(65),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.0002092775394366469),
		Fy:    libc.Float64FromFloat64(0.00020925564394716004),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	50: {
		Ffile: __ccgo_ts,
		Fline: int32(66),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.00021119991881665427),
		Fy:    libc.Float64FromFloat64(0.00021117761925352221),
		Fdy:   float32(-libc.Float64FromFloat64(6.283911383761856e-16)),
		Fe:    int32(FE_INEXACT),
	},
	51: {
		Ffile: __ccgo_ts,
		Fline: int32(67),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.00010101840866170902),
		Fy:    libc.Float64FromFloat64(0.00010101330664586021),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	52: {
		Ffile: __ccgo_ts,
		Fline: int32(68),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.48581699901167e-05),
		Fy:    libc.Float64FromFloat64(3.485756245822742e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	53: {
		Ffile: __ccgo_ts,
		Fline: int32(69),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.855323594827083e-05),
		Fy:    libc.Float64FromFloat64(3.855249279137047e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	54: {
		Ffile: __ccgo_ts,
		Fline: int32(70),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.193155921359673e-05),
		Fy:    libc.Float64FromFloat64(4.193068011034238e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	55: {
		Ffile: __ccgo_ts,
		Fline: int32(71),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.446535422039716e-05),
		Fy:    libc.Float64FromFloat64(4.4464365665838354e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	56: {
		Ffile: __ccgo_ts,
		Fline: int32(72),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.0206066447479943e-05),
		Fy:    libc.Float64FromFloat64(6.020425413500251e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	57: {
		Ffile: __ccgo_ts,
		Fline: int32(73),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.3599729651213705e-05),
		Fy:    libc.Float64FromFloat64(2.359945118197509e-05),
		Fdy:   float32(-libc.Float64FromFloat64(8.816943824888936e-16)),
		Fe:    int32(FE_INEXACT),
	},
	58: {
		Ffile: __ccgo_ts,
		Fline: int32(74),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.6816689109184868e-05),
		Fy:    libc.Float64FromFloat64(2.681632954820562e-05),
		Fdy:   float32(-libc.Float64FromFloat64(7.721578816949804e-17)),
		Fe:    int32(FE_INEXACT),
	},
	59: {
		Ffile: __ccgo_ts,
		Fline: int32(75),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.9408923495220864e-05),
		Fy:    libc.Float64FromFloat64(2.9408491061308546e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1.502109699815055e-17)),
		Fe:    int32(FE_INEXACT),
	},
	60: {
		Ffile: __ccgo_ts,
		Fline: int32(76),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.968078842982953e-05),
		Fy:    libc.Float64FromFloat64(2.968034796394418e-05),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	61: {
		Ffile: __ccgo_ts,
		Fline: int32(77),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.958693727991685e-06),
		Fy:    libc.Float64FromFloat64(7.958662057756791e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	62: {
		Ffile: __ccgo_ts,
		Fline: int32(78),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.921719477770297e-06),
		Fy:    libc.Float64FromFloat64(7.921688101116258e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	63: {
		Ffile: __ccgo_ts,
		Fline: int32(79),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.290531470131644e-05),
		Fy:    libc.Float64FromFloat64(1.290523142845911e-05),
		Fdy:   float32(-libc.Float64FromFloat64(9.119802572762079e-17)),
		Fe:    int32(FE_INEXACT),
	},
	64: {
		Ffile: __ccgo_ts,
		Fline: int32(80),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.261814511230699e-06),
		Fy:    libc.Float64FromFloat64(6.261794906152054e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	65: {
		Ffile: __ccgo_ts,
		Fline: int32(81),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.3083593475617995e-06),
		Fy:    libc.Float64FromFloat64(2.3083566833044604e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	66: {
		Ffile: __ccgo_ts,
		Fline: int32(82),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.917704733221056e-06),
		Fy:    libc.Float64FromFloat64(2.9177004767288804e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	67: {
		Ffile: __ccgo_ts,
		Fline: int32(83),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.469793658418255e-06),
		Fy:    libc.Float64FromFloat64(3.469787638698164e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	68: {
		Ffile: __ccgo_ts,
		Fline: int32(84),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6843664990071396e-06),
		Fy:    libc.Float64FromFloat64(3.6843597117455613e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1.5604146833437128e-16)),
		Fe:    int32(FE_INEXACT),
	},
	69: {
		Ffile: __ccgo_ts,
		Fline: int32(85),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5056064487773347),
		Fy:    libc.Float64FromFloat64(0.40919577303410304),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	70: {
		Ffile: __ccgo_ts,
		Fline: int32(86),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5074662373909726),
		Fy:    libc.Float64FromFloat64(0.4104302529445011),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	71: {
		Ffile: __ccgo_ts,
		Fline: int32(87),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.543799609375296),
		Fy:    libc.Float64FromFloat64(0.4342466565055341),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	72: {
		Ffile: __ccgo_ts,
		Fline: int32(88),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1600162034933231e-06),
		Fy:    libc.Float64FromFloat64(1.160015530675047e-06),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	73: {
		Ffile: __ccgo_ts,
		Fline: int32(89),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2368384476160648e-06),
		Fy:    libc.Float64FromFloat64(1.2368376827320229e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	74: {
		Ffile: __ccgo_ts,
		Fline: int32(90),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6108658190178803e-06),
		Fy:    libc.Float64FromFloat64(1.6108645215749302e-06),
		Fdy:   float32(-libc.Float64FromFloat64(4.669940924745622e-16)),
		Fe:    int32(FE_INEXACT),
	},
	75: {
		Ffile: __ccgo_ts,
		Fline: int32(91),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8150378441001636e-06),
		Fy:    libc.Float64FromFloat64(1.815036196920969e-06),
		Fdy:   float32(-libc.Float64FromFloat64(7.055955974800141e-16)),
		Fe:    int32(FE_INEXACT),
	},
	76: {
		Ffile: __ccgo_ts,
		Fline: int32(92),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.657174710290451e-07),
		Fy:    libc.Float64FromFloat64(5.657173110109769e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	77: {
		Ffile: __ccgo_ts,
		Fline: int32(93),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.596248053690229e-07),
		Fy:    libc.Float64FromFloat64(6.596245878166766e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	78: {
		Ffile: __ccgo_ts,
		Fline: int32(94),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.29054993115267e-07),
		Fy:    libc.Float64FromFloat64(7.290547273548048e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	79: {
		Ffile: __ccgo_ts,
		Fline: int32(95),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.755650353541528e-07),
		Fy:    libc.Float64FromFloat64(8.755646520473108e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	80: {
		Ffile: __ccgo_ts,
		Fline: int32(96),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.203265211132231e-07),
		Fy:    libc.Float64FromFloat64(4.203264327760557e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	81: {
		Ffile: __ccgo_ts,
		Fline: int32(97),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6649705275776103e-07),
		Fy:    libc.Float64FromFloat64(3.6649698559773254e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	82: {
		Ffile: __ccgo_ts,
		Fline: int32(98),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.873620930581013e-07),
		Fy:    libc.Float64FromFloat64(3.873620180334251e-07),
		Fdy:   float32(-libc.Float64FromFloat64(2.0258804636836322e-16)),
		Fe:    int32(FE_INEXACT),
	},
	83: {
		Ffile: __ccgo_ts,
		Fline: int32(99),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.738541335709014e-07),
		Fy:    libc.Float64FromFloat64(3.738540636874622e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	84: {
		Ffile: __ccgo_ts,
		Fline: int32(100),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.842603999897839e-07),
		Fy:    libc.Float64FromFloat64(3.8426032616177524e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	85: {
		Ffile: __ccgo_ts,
		Fline: int32(101),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.943527693522589e-07),
		Fy:    libc.Float64FromFloat64(3.94352691595226e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1.0279554790186747e-16)),
		Fe:    int32(FE_INEXACT),
	},
	86: {
		Ffile: __ccgo_ts,
		Fline: int32(102),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.633802238542776e-07),
		Fy:    libc.Float64FromFloat64(4.6338011649369486e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	87: {
		Ffile: __ccgo_ts,
		Fline: int32(103),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3202617372917325e-07),
		Fy:    libc.Float64FromFloat64(1.3202616501371874e-07),
		Fdy:   float32(-libc.Float64FromFloat64(6.749060056142147e-16)),
		Fe:    int32(FE_INEXACT),
	},
	88: {
		Ffile: __ccgo_ts,
		Fline: int32(104),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6961460528655013e-07),
		Fy:    libc.Float64FromFloat64(1.696145909019946e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	89: {
		Ffile: __ccgo_ts,
		Fline: int32(105),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.406603434028587e-08),
		Fy:    libc.Float64FromFloat64(6.406603228805757e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	90: {
		Ffile: __ccgo_ts,
		Fline: int32(106),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.95745443871373e-08),
		Fy:    libc.Float64FromFloat64(9.957453942959268e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	91: {
		Ffile: __ccgo_ts,
		Fline: int32(107),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.832604178600878e-08),
		Fy:    libc.Float64FromFloat64(7.832603871852451e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	92: {
		Ffile: __ccgo_ts,
		Fline: int32(108),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.96235750919584e-08),
		Fy:    libc.Float64FromFloat64(8.962357107576604e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	93: {
		Ffile: __ccgo_ts,
		Fline: int32(109),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1387033517479312e-07),
		Fy:    libc.Float64FromFloat64(1.13870328691567e-07),
		Fdy:   float32(-libc.Float64FromFloat64(1.417813245602103e-16)),
		Fe:    int32(FE_INEXACT),
	},
	94: {
		Ffile: __ccgo_ts,
		Fline: int32(110),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.153278926107579e-07),
		Fy:    libc.Float64FromFloat64(1.15327885960497e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	95: {
		Ffile: __ccgo_ts,
		Fline: int32(111),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.523079836230221e-08),
		Fy:    libc.Float64FromFloat64(4.523079733938968e-08),
		Fdy:   float32(-libc.Float64FromFloat64(3.101599071806793e-17)),
		Fe:    int32(FE_INEXACT),
	},
	96: {
		Ffile: __ccgo_ts,
		Fline: int32(112),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.1256006266036996e-08),
		Fy:    libc.Float64FromFloat64(5.125600495244795e-08),
		Fdy:   float32(-libc.Float64FromFloat64(8.547468763049453e-16)),
		Fe:    int32(FE_INEXACT),
	},
	97: {
		Ffile: __ccgo_ts,
		Fline: int32(113),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.6663579747679114e-08),
		Fy:    libc.Float64FromFloat64(4.6663578658934304e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	98: {
		Ffile: __ccgo_ts,
		Fline: int32(114),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.0969215396040908e-08),
		Fy:    libc.Float64FromFloat64(2.0969215176186913e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	99: {
		Ffile: __ccgo_ts,
		Fline: int32(115),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8429598420714387e-08),
		Fy:    libc.Float64FromFloat64(1.8429598250889337e-08),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	100: {
		Ffile: __ccgo_ts,
		Fline: int32(116),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.393349073739792e-08),
		Fy:    libc.Float64FromFloat64(2.3933490450991933e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	101: {
		Ffile: __ccgo_ts,
		Fline: int32(117),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.290926187970368e-09),
		Fy:    libc.Float64FromFloat64(8.290926153600638e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	102: {
		Ffile: __ccgo_ts,
		Fline: int32(118),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.498367243320052e-09),
		Fy:    libc.Float64FromFloat64(9.498367198210562e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	103: {
		Ffile: __ccgo_ts,
		Fline: int32(119),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.140500965186617e-09),
		Fy:    libc.Float64FromFloat64(9.140500923412239e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	104: {
		Ffile: __ccgo_ts,
		Fline: int32(120),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4431803935651913e-08),
		Fy:    libc.Float64FromFloat64(1.4431803831513432e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	105: {
		Ffile: __ccgo_ts,
		Fline: int32(121),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.750222268453456e-09),
		Fy:    libc.Float64FromFloat64(6.750222245670705e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	106: {
		Ffile: __ccgo_ts,
		Fline: int32(122),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.094957364184897e-09),
		Fy:    libc.Float64FromFloat64(6.094957345610644e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	107: {
		Ffile: __ccgo_ts,
		Fline: int32(123),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.369629129097454e-09),
		Fy:    libc.Float64FromFloat64(6.3696291088113655e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	108: {
		Ffile: __ccgo_ts,
		Fline: int32(124),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.809339046451396e-09),
		Fy:    libc.Float64FromFloat64(6.809339023267846e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	109: {
		Ffile: __ccgo_ts,
		Fline: int32(125),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2628331465533393e-09),
		Fy:    libc.Float64FromFloat64(2.2628331439931326e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	110: {
		Ffile: __ccgo_ts,
		Fline: int32(126),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.608843979110671e-09),
		Fy:    libc.Float64FromFloat64(2.608843975707637e-09),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	111: {
		Ffile: __ccgo_ts,
		Fline: int32(127),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.26494329502906727),
		Fy:    libc.Float64FromFloat64(0.2350272951107871),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	112: {
		Ffile: __ccgo_ts,
		Fline: int32(128),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.550544278289493e-10),
		Fy:    libc.Float64FromFloat64(9.550544273728847e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	113: {
		Ffile: __ccgo_ts,
		Fline: int32(129),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.132678983883741e-09),
		Fy:    libc.Float64FromFloat64(1.1326789832422602e-09),
		Fdy:   float32(-libc.Float64FromFloat64(2.836218884871692e-16)),
		Fe:    int32(FE_INEXACT),
	},
	114: {
		Ffile: __ccgo_ts,
		Fline: int32(130),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5949869037907344e-09),
		Fy:    libc.Float64FromFloat64(1.5949869025187428e-09),
		Fdy:   float32(-libc.Float64FromFloat64(8.454553483900392e-16)),
		Fe:    int32(FE_INEXACT),
	},
	115: {
		Ffile: __ccgo_ts,
		Fline: int32(131),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.63939270902679e-09),
		Fy:    libc.Float64FromFloat64(1.639392707682986e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	116: {
		Ffile: __ccgo_ts,
		Fline: int32(132),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.5374814258916699e-09),
		Fy:    libc.Float64FromFloat64(1.5374814247097454e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	117: {
		Ffile: __ccgo_ts,
		Fline: int32(133),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7217730657808563e-10),
		Fy:    libc.Float64FromFloat64(2.7217730654104536e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	118: {
		Ffile: __ccgo_ts,
		Fline: int32(134),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.043278370436311e-10),
		Fy:    libc.Float64FromFloat64(3.043278369973234e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	119: {
		Ffile: __ccgo_ts,
		Fline: int32(135),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.266127678996292e-10),
		Fy:    libc.Float64FromFloat64(3.2661276784629125e-10),
		Fdy:   float32(-libc.Float64FromFloat64(6.73797557123614e-16)),
		Fe:    int32(FE_INEXACT),
	},
	120: {
		Ffile: __ccgo_ts,
		Fline: int32(136),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.312664299541643e-10),
		Fy:    libc.Float64FromFloat64(4.312664298611689e-10),
		Fdy:   float32(-libc.Float64FromFloat64(6.96306928301305e-16)),
		Fe:    int32(FE_INEXACT),
	},
	121: {
		Ffile: __ccgo_ts,
		Fline: int32(137),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1800693755967932e-10),
		Fy:    libc.Float64FromFloat64(1.180069375527165e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.0839019558727557e-16)),
		Fe:    int32(FE_INEXACT),
	},
	122: {
		Ffile: __ccgo_ts,
		Fline: int32(138),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.214175426916e-10),
		Fy:    libc.Float64FromFloat64(1.2141754268422888e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.3354622708778104e-16)),
		Fe:    int32(FE_INEXACT),
	},
	123: {
		Ffile: __ccgo_ts,
		Fline: int32(139),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2482814782352842e-10),
		Fy:    libc.Float64FromFloat64(1.248281478157374e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.609138880833471e-16)),
		Fe:    int32(FE_INEXACT),
	},
	124: {
		Ffile: __ccgo_ts,
		Fline: int32(140),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.282387529554646e-10),
		Fy:    libc.Float64FromFloat64(1.2823875294724202e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.906191747248778e-16)),
		Fe:    int32(FE_INEXACT),
	},
	125: {
		Ffile: __ccgo_ts,
		Fline: int32(141),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.384705683513197e-10),
		Fy:    libc.Float64FromFloat64(1.3847056834173264e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.9507268374187714e-16)),
		Fe:    int32(FE_INEXACT),
	},
	126: {
		Ffile: __ccgo_ts,
		Fline: int32(142),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5211298887923503e-10),
		Fy:    libc.Float64FromFloat64(1.5211298886766585e-10),
		Fdy:   float32(-libc.Float64FromFloat64(5.753237830768748e-16)),
		Fe:    int32(FE_INEXACT),
	},
	127: {
		Ffile: __ccgo_ts,
		Fline: int32(143),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6575540940727446e-10),
		Fy:    libc.Float64FromFloat64(1.6575540939353703e-10),
		Fdy:   float32(-libc.Float64FromFloat64(8.111813777892556e-16)),
		Fe:    int32(FE_INEXACT),
	},
	128: {
		Ffile: __ccgo_ts,
		Fline: int32(144),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3505996321936024e-10),
		Fy:    libc.Float64FromFloat64(1.3505996321023964e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.575638678457498e-16)),
		Fe:    int32(FE_INEXACT),
	},
	129: {
		Ffile: __ccgo_ts,
		Fline: int32(145),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3164935808740855e-10),
		Fy:    libc.Float64FromFloat64(1.3164935807874277e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.2279149776484584e-16)),
		Fe:    int32(FE_INEXACT),
	},
	130: {
		Ffile: __ccgo_ts,
		Fline: int32(146),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6234480427525297e-10),
		Fy:    libc.Float64FromFloat64(1.6234480426207505e-10),
		Fdy:   float32(-libc.Float64FromFloat64(7.464499552897795e-16)),
		Fe:    int32(FE_INEXACT),
	},
	131: {
		Ffile: __ccgo_ts,
		Fline: int32(147),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4529177861526185e-10),
		Fy:    libc.Float64FromFloat64(1.45291778604707e-10),
		Fdy:   float32(-libc.Float64FromFloat64(4.788629828292854e-16)),
		Fe:    int32(FE_INEXACT),
	},
	132: {
		Ffile: __ccgo_ts,
		Fline: int32(148),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4870238374724457e-10),
		Fy:    libc.Float64FromFloat64(1.4870238373618837e-10),
		Fdy:   float32(-libc.Float64FromFloat64(5.254348130330174e-16)),
		Fe:    int32(FE_INEXACT),
	},
	133: {
		Ffile: __ccgo_ts,
		Fline: int32(149),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.418811734832869e-10),
		Fy:    libc.Float64FromFloat64(1.4188117347322176e-10),
		Fdy:   float32(-libc.Float64FromFloat64(4.354579176477617e-16)),
		Fe:    int32(FE_INEXACT),
	},
	134: {
		Ffile: __ccgo_ts,
		Fline: int32(150),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.691660145393037e-10),
		Fy:    libc.Float64FromFloat64(1.6916601452499513e-10),
		Fdy:   float32(-libc.Float64FromFloat64(8.800343037913935e-16)),
		Fe:    int32(FE_INEXACT),
	},
	135: {
		Ffile: __ccgo_ts,
		Fline: int32(151),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5552359401123326e-10),
		Fy:    libc.Float64FromFloat64(1.5552359399913946e-10),
		Fdy:   float32(-libc.Float64FromFloat64(6.286837882594619e-16)),
		Fe:    int32(FE_INEXACT),
	},
	136: {
		Ffile: __ccgo_ts,
		Fline: int32(152),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5893419914323924e-10),
		Fy:    libc.Float64FromFloat64(1.589341991306092e-10),
		Fdy:   float32(-libc.Float64FromFloat64(6.856722178902905e-16)),
		Fe:    int32(FE_INEXACT),
	},
	137: {
		Ffile: __ccgo_ts,
		Fline: int32(153),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8970106102410336e-10),
		Fy:    libc.Float64FromFloat64(1.8970106100611009e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	138: {
		Ffile: __ccgo_ts,
		Fline: int32(154),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.9046647784737756e-10),
		Fy:    libc.Float64FromFloat64(1.904664778292388e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	139: {
		Ffile: __ccgo_ts,
		Fline: int32(155),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.105290543524238e-10),
		Fy:    libc.Float64FromFloat64(2.1052905433026254e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	140: {
		Ffile: __ccgo_ts,
		Fline: int32(156),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.0542428274927651e-10),
		Fy:    libc.Float64FromFloat64(2.0542428272817695e-10),
		Fdy:   float32(-libc.Float64FromFloat64(8.784220824554134e-16)),
		Fe:    int32(FE_INEXACT),
	},
	141: {
		Ffile: __ccgo_ts,
		Fline: int32(157),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.832134775232282e-11),
		Fy:    libc.Float64FromFloat64(5.832134775062213e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	142: {
		Ffile: __ccgo_ts,
		Fline: int32(158),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.002665031821433e-11),
		Fy:    libc.Float64FromFloat64(6.002665031641273e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.7903144636908374e-17)),
		Fe:    int32(FE_INEXACT),
	},
	143: {
		Ffile: __ccgo_ts,
		Fline: int32(159),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.51425580159005e-11),
		Fy:    libc.Float64FromFloat64(6.514255801377873e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	144: {
		Ffile: __ccgo_ts,
		Fline: int32(160),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.173195288410778e-11),
		Fy:    libc.Float64FromFloat64(6.173195288220237e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	145: {
		Ffile: __ccgo_ts,
		Fline: int32(161),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.343725545000317e-11),
		Fy:    libc.Float64FromFloat64(6.343725544799103e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.4806039640557294e-17)),
		Fe:    int32(FE_INEXACT),
	},
	146: {
		Ffile: __ccgo_ts,
		Fline: int32(162),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.025846571360412e-11),
		Fy:    libc.Float64FromFloat64(7.0258465711136e-11),
		Fdy:   float32(-libc.Float64FromFloat64(5.2368685466723974e-17)),
		Fe:    int32(FE_INEXACT),
	},
	147: {
		Ffile: __ccgo_ts,
		Fline: int32(163),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.537437341132519e-11),
		Fy:    libc.Float64FromFloat64(7.537437340848455e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	148: {
		Ffile: __ccgo_ts,
		Fline: int32(164),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.684786058179977e-11),
		Fy:    libc.Float64FromFloat64(6.684786057956545e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.2916777157211245e-17)),
		Fe:    int32(FE_INEXACT),
	},
	149: {
		Ffile: __ccgo_ts,
		Fline: int32(165),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.196376827950921e-11),
		Fy:    libc.Float64FromFloat64(7.196376827691982e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	150: {
		Ffile: __ccgo_ts,
		Fline: int32(166),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.855316314770098e-11),
		Fy:    libc.Float64FromFloat64(6.855316314535122e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	151: {
		Ffile: __ccgo_ts,
		Fline: int32(167),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.231993183449124e-11),
		Fy:    libc.Float64FromFloat64(7.231993183187616e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	152: {
		Ffile: __ccgo_ts,
		Fline: int32(168),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.70796759772361e-11),
		Fy:    libc.Float64FromFloat64(7.707967597426546e-11),
		Fdy:   float32(-libc.Float64FromFloat64(7.586413534391699e-17)),
		Fe:    int32(FE_INEXACT),
	},
	153: {
		Ffile: __ccgo_ts,
		Fline: int32(169),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.366907084541623e-11),
		Fy:    libc.Float64FromFloat64(7.366907084270267e-11),
		Fdy:   float32(-libc.Float64FromFloat64(6.330205440098447e-17)),
		Fe:    int32(FE_INEXACT),
	},
	154: {
		Ffile: __ccgo_ts,
		Fline: int32(170),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.560618880681968e-11),
		Fy:    libc.Float64FromFloat64(8.560618880315547e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	155: {
		Ffile: __ccgo_ts,
		Fline: int32(171),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1289102986181285e-10),
		Fy:    libc.Float64FromFloat64(1.1289102985544066e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	156: {
		Ffile: __ccgo_ts,
		Fline: int32(172),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0095391190019227e-10),
		Fy:    libc.Float64FromFloat64(1.0095391189509642e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.2323881550154146e-16)),
		Fe:    int32(FE_INEXACT),
	},
	157: {
		Ffile: __ccgo_ts,
		Fline: int32(173),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.390088624089908e-11),
		Fy:    libc.Float64FromFloat64(8.39008862373794e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.0649843184012844e-16)),
		Fe:    int32(FE_INEXACT),
	},
	158: {
		Ffile: __ccgo_ts,
		Fline: int32(174),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.072209650459309e-11),
		Fy:    libc.Float64FromFloat64(9.072209650047784e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.4558926705371042e-16)),
		Fe:    int32(FE_INEXACT),
	},
	159: {
		Ffile: __ccgo_ts,
		Fline: int32(175),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.413270163645173e-11),
		Fy:    libc.Float64FromFloat64(9.413270163202125e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.6874820662282706e-16)),
		Fe:    int32(FE_INEXACT),
	},
	160: {
		Ffile: __ccgo_ts,
		Fline: int32(176),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.924860933425422e-11),
		Fy:    libc.Float64FromFloat64(9.924860932932909e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	161: {
		Ffile: __ccgo_ts,
		Fline: int32(177),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.901679393866668e-11),
		Fy:    libc.Float64FromFloat64(8.901679393470469e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	162: {
		Ffile: __ccgo_ts,
		Fline: int32(178),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.242739907052144e-11),
		Fy:    libc.Float64FromFloat64(9.242739906625003e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	163: {
		Ffile: __ccgo_ts,
		Fline: int32(179),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0436451703207417e-10),
		Fy:    libc.Float64FromFloat64(1.0436451702662819e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.54969702027092e-16)),
		Fe:    int32(FE_INEXACT),
	},
	164: {
		Ffile: __ccgo_ts,
		Fline: int32(180),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.878497854314893e-11),
		Fy:    libc.Float64FromFloat64(7.87849785400454e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	165: {
		Ffile: __ccgo_ts,
		Fline: int32(181),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0606981959801802e-10),
		Fy:    libc.Float64FromFloat64(1.0606981959239263e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	166: {
		Ffile: __ccgo_ts,
		Fline: int32(182),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0265921446613225e-10),
		Fy:    libc.Float64FromFloat64(1.026592144608628e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	167: {
		Ffile: __ccgo_ts,
		Fline: int32(183),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.731149137274221e-11),
		Fy:    libc.Float64FromFloat64(8.731149136893056e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.249000638005505e-16)),
		Fe:    int32(FE_INEXACT),
	},
	168: {
		Ffile: __ccgo_ts,
		Fline: int32(184),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.754330676831812e-11),
		Fy:    libc.Float64FromFloat64(9.754330676356077e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.9456600908315455e-16)),
		Fe:    int32(FE_INEXACT),
	},
	169: {
		Ffile: __ccgo_ts,
		Fline: int32(185),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.583800420238395e-11),
		Fy:    libc.Float64FromFloat64(9.58380041977915e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	170: {
		Ffile: __ccgo_ts,
		Fline: int32(186),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0777512216396382e-10),
		Fy:    libc.Float64FromFloat64(1.0777512215815608e-10),
		Fdy:   float32(-libc.Float64FromFloat64(2.899687593005049e-16)),
		Fe:    int32(FE_INEXACT),
	},
	171: {
		Ffile: __ccgo_ts,
		Fline: int32(187),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1221311298116892e-10),
		Fy:    libc.Float64FromFloat64(1.1221311297487302e-10),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	172: {
		Ffile: __ccgo_ts,
		Fline: int32(188),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1118572729586123e-10),
		Fy:    libc.Float64FromFloat64(1.111857272896801e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.28453039514514e-16)),
		Fe:    int32(FE_INEXACT),
	},
	173: {
		Ffile: __ccgo_ts,
		Fline: int32(189),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1630163499372189e-10),
		Fy:    libc.Float64FromFloat64(1.1630163498695886e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	174: {
		Ffile: __ccgo_ts,
		Fline: int32(190),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0948042472991156e-10),
		Fy:    libc.Float64FromFloat64(1.0948042472391858e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	175: {
		Ffile: __ccgo_ts,
		Fline: int32(191),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.049028110906371e-11),
		Fy:    libc.Float64FromFloat64(8.049028110582437e-11),
		Fdy:   float32(-libc.Float64FromFloat64(9.02091941710402e-17)),
		Fe:    int32(FE_INEXACT),
	},
	176: {
		Ffile: __ccgo_ts,
		Fline: int32(192),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.219558367498043e-11),
		Fy:    libc.Float64FromFloat64(8.219558367160238e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	177: {
		Ffile: __ccgo_ts,
		Fline: int32(193),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.145963324277664e-10),
		Fy:    libc.Float64FromFloat64(1.1459633242120024e-10),
		Fdy:   float32(-libc.Float64FromFloat64(3.7064655641388846e-16)),
		Fe:    int32(FE_INEXACT),
	},
	178: {
		Ffile: __ccgo_ts,
		Fline: int32(194),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.035438567197861e-11),
		Fy:    libc.Float64FromFloat64(3.0354385671517915e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.649150285669534e-18)),
		Fe:    int32(FE_INEXACT),
	},
	179: {
		Ffile: __ccgo_ts,
		Fline: int32(195),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.205968823783833e-11),
		Fy:    libc.Float64FromFloat64(3.2059688237324416e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.5409171604507024e-18)),
		Fe:    int32(FE_INEXACT),
	},
	180: {
		Ffile: __ccgo_ts,
		Fline: int32(196),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.171862772466623e-11),
		Fy:    libc.Float64FromFloat64(3.1718627724163194e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.35073179401251e-18)),
		Fe:    int32(FE_INEXACT),
	},
	181: {
		Ffile: __ccgo_ts,
		Fline: int32(197),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.00133251588069e-11),
		Fy:    libc.Float64FromFloat64(3.00133251583565e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.487869918556395e-18)),
		Fe:    int32(FE_INEXACT),
	},
	182: {
		Ffile: __ccgo_ts,
		Fline: int32(198),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.34239302905275e-11),
		Fy:    libc.Float64FromFloat64(3.342393028996892e-11),
		Fdy:   float32(-libc.Float64FromFloat64(5.3645772923167044e-18)),
		Fe:    int32(FE_INEXACT),
	},
	183: {
		Ffile: __ccgo_ts,
		Fline: int32(199),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.3764990803699986e-11),
		Fy:    libc.Float64FromFloat64(3.376499080312995e-11),
		Fdy:   float32(-libc.Float64FromFloat64(5.586923440970958e-18)),
		Fe:    int32(FE_INEXACT),
	},
	184: {
		Ffile: __ccgo_ts,
		Fline: int32(200),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.547029336956358e-11),
		Fy:    libc.Float64FromFloat64(3.547029336893451e-11),
		Fdy:   float32(-libc.Float64FromFloat64(6.804003907056981e-18)),
		Fe:    int32(FE_INEXACT),
	},
	185: {
		Ffile: __ccgo_ts,
		Fline: int32(201),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.512923285639071e-11),
		Fy:    libc.Float64FromFloat64(3.5129232855773676e-11),
		Fdy:   float32(-libc.Float64FromFloat64(6.5460823746180466e-18)),
		Fe:    int32(FE_INEXACT),
	},
	186: {
		Ffile: __ccgo_ts,
		Fline: int32(202),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.853983798812294e-11),
		Fy:    libc.Float64FromFloat64(3.853983798738028e-11),
		Fdy:   float32(-libc.Float64FromFloat64(9.483010300544723e-18)),
		Fe:    int32(FE_INEXACT),
	},
	187: {
		Ffile: __ccgo_ts,
		Fline: int32(203),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6834535422255855e-11),
		Fy:    libc.Float64FromFloat64(3.683453542157746e-11),
		Fdy:   float32(-libc.Float64FromFloat64(7.912743565233258e-18)),
		Fe:    int32(FE_INEXACT),
	},
	188: {
		Ffile: __ccgo_ts,
		Fline: int32(204),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.2291503633037353e-11),
		Fy:    libc.Float64FromFloat64(4.229150363214307e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.3750521107489498e-17)),
		Fe:    int32(FE_INEXACT),
	},
	189: {
		Ffile: __ccgo_ts,
		Fline: int32(205),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.7175595935429117e-11),
		Fy:    libc.Float64FromFloat64(3.7175595934738104e-11),
		Fdy:   float32(-libc.Float64FromFloat64(8.209893311041908e-18)),
		Fe:    int32(FE_INEXACT),
	},
	190: {
		Ffile: __ccgo_ts,
		Fline: int32(206),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.0586201067166e-11),
		Fy:    libc.Float64FromFloat64(4.058620106634238e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.1663273106777292e-17)),
		Fe:    int32(FE_INEXACT),
	},
	191: {
		Ffile: __ccgo_ts,
		Fline: int32(207),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.0245140553991965e-11),
		Fy:    libc.Float64FromFloat64(4.024514055318213e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.1276126110322873e-17)),
		Fe:    int32(FE_INEXACT),
	},
	192: {
		Ffile: __ccgo_ts,
		Fline: int32(208),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.888089850129659e-11),
		Fy:    libc.Float64FromFloat64(3.888089850054073e-11),
		Fdy:   float32(-libc.Float64FromFloat64(9.82317343820613e-18)),
		Fe:    int32(FE_INEXACT),
	},
	193: {
		Ffile: __ccgo_ts,
		Fline: int32(209),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.195044311986293e-11),
		Fy:    libc.Float64FromFloat64(4.195044311898301e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.3312287436403804e-17)),
		Fe:    int32(FE_INEXACT),
	},
	194: {
		Ffile: __ccgo_ts,
		Fline: int32(210),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.706635081748745e-11),
		Fy:    libc.Float64FromFloat64(4.706635081637983e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.109350266679623e-17)),
		Fe:    int32(FE_INEXACT),
	},
	195: {
		Ffile: __ccgo_ts,
		Fline: int32(211),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.399680619891064e-11),
		Fy:    libc.Float64FromFloat64(4.3996806197942784e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.610614620358331e-17)),
		Fe:    int32(FE_INEXACT),
	},
	196: {
		Ffile: __ccgo_ts,
		Fline: int32(212),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.570210876478587e-11),
		Fy:    libc.Float64FromFloat64(4.570210876374153e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.8752171251687337e-17)),
		Fe:    int32(FE_INEXACT),
	},
	197: {
		Ffile: __ccgo_ts,
		Fline: int32(213),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.536104825161067e-11),
		Fy:    libc.Float64FromFloat64(4.536104825058186e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.819863522065665e-17)),
		Fe:    int32(FE_INEXACT),
	},
	198: {
		Ffile: __ccgo_ts,
		Fline: int32(214),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.365574568573583e-11),
		Fy:    libc.Float64FromFloat64(4.365574568478292e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.5612484814011662e-17)),
		Fe:    int32(FE_INEXACT),
	},
	199: {
		Ffile: __ccgo_ts,
		Fline: int32(215),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.9112713896542144e-11),
		Fy:    libc.Float64FromFloat64(4.9112713895336115e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.500819778142999e-17)),
		Fe:    int32(FE_INEXACT),
	},
	200: {
		Ffile: __ccgo_ts,
		Fline: int32(216),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.877165338336617e-11),
		Fy:    libc.Float64FromFloat64(4.877165338217683e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.432075113539432e-17)),
		Fe:    int32(FE_INEXACT),
	},
	201: {
		Ffile: __ccgo_ts,
		Fline: int32(217),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.740741133066304e-11),
		Fy:    libc.Float64FromFloat64(4.740741132953931e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.1711598490275455e-17)),
		Fe:    int32(FE_INEXACT),
	},
	202: {
		Ffile: __ccgo_ts,
		Fline: int32(218),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.0476955949246825e-11),
		Fy:    libc.Float64FromFloat64(5.0476955947972863e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.790483870280288e-17)),
		Fe:    int32(FE_INEXACT),
	},
	203: {
		Ffile: __ccgo_ts,
		Fline: int32(219),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.081801646242319e-11),
		Fy:    libc.Float64FromFloat64(5.0818016461131953e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.866669189929892e-17)),
		Fe:    int32(FE_INEXACT),
	},
	204: {
		Ffile: __ccgo_ts,
		Fline: int32(220),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.4228621594191094e-11),
		Fy:    libc.Float64FromFloat64(5.422862159272072e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.71724644067286e-17)),
		Fe:    int32(FE_INEXACT),
	},
	205: {
		Ffile: __ccgo_ts,
		Fline: int32(221),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.218225851512942e-11),
		Fy:    libc.Float64FromFloat64(5.2182258513767926e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.187120282721915e-17)),
		Fe:    int32(FE_INEXACT),
	},
	206: {
		Ffile: __ccgo_ts,
		Fline: int32(222),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.252331902830617e-11),
		Fy:    libc.Float64FromFloat64(5.252331902692682e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.2712650650977494e-17)),
		Fe:    int32(FE_INEXACT),
	},
	207: {
		Ffile: __ccgo_ts,
		Fline: int32(223),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.559286364690043e-11),
		Fy:    libc.Float64FromFloat64(5.5592863645355144e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.1056613395701997e-17)),
		Fe:    int32(FE_INEXACT),
	},
	208: {
		Ffile: __ccgo_ts,
		Fline: int32(224),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.7639226725966754e-11),
		Fy:    libc.Float64FromFloat64(5.7639226724305614e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.744382415315066e-17)),
		Fe:    int32(FE_INEXACT),
	},
	209: {
		Ffile: __ccgo_ts,
		Fline: int32(225),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.3887561081013955e-11),
		Fy:    libc.Float64FromFloat64(5.388756107956202e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.624607506022841e-17)),
		Fe:    int32(FE_INEXACT),
	},
	210: {
		Ffile: __ccgo_ts,
		Fline: int32(226),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.5933924160077955e-11),
		Fy:    libc.Float64FromFloat64(5.593392415851365e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.207344997910118e-17)),
		Fe:    int32(FE_INEXACT),
	},
	211: {
		Ffile: __ccgo_ts,
		Fline: int32(227),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.729816621278884e-11),
		Fy:    libc.Float64FromFloat64(5.72981662111473e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.6330796390678906e-17)),
		Fe:    int32(FE_INEXACT),
	},
	212: {
		Ffile: __ccgo_ts,
		Fline: int32(228),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.7694194001486496e-11),
		Fy:    libc.Float64FromFloat64(5.769419399982219e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	213: {
		Ffile: __ccgo_ts,
		Fline: int32(229),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6551589358644582e-11),
		Fy:    libc.Float64FromFloat64(1.6551589358507606e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	214: {
		Ffile: __ccgo_ts,
		Fline: int32(230),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8092287805240843e-11),
		Fy:    libc.Float64FromFloat64(1.8092287805077178e-11),
		Fdy:   float32(-libc.Float64FromFloat64(6.761611496233679e-16)),
		Fe:    int32(FE_INEXACT),
	},
	215: {
		Ffile: __ccgo_ts,
		Fline: int32(231),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.8622185110624338e-11),
		Fy:    libc.Float64FromFloat64(1.8622185110450943e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	216: {
		Ffile: __ccgo_ts,
		Fline: int32(232),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.7394105069021052e-11),
		Fy:    libc.Float64FromFloat64(2.7394105068645832e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	217: {
		Ffile: __ccgo_ts,
		Fline: int32(233),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.721773065558615e-11),
		Fy:    libc.Float64FromFloat64(2.7217730655215743e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	218: {
		Ffile: __ccgo_ts,
		Fline: int32(234),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.503331289645425e-12),
		Fy:    libc.Float64FromFloat64(7.503331289617275e-12),
		Fdy:   float32(-libc.Float64FromFloat64(5.447480642029219e-20)),
		Fe:    int32(FE_INEXACT),
	},
	219: {
		Ffile: __ccgo_ts,
		Fline: int32(235),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.355982572562052e-12),
		Fy:    libc.Float64FromFloat64(8.355982572527141e-12),
		Fdy:   float32(-libc.Float64FromFloat64(8.380332221897234e-20)),
		Fe:    int32(FE_INEXACT),
	},
	220: {
		Ffile: __ccgo_ts,
		Fline: int32(236),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.673861546228711e-12),
		Fy:    libc.Float64FromFloat64(7.673861546199267e-12),
		Fdy:   float32(-libc.Float64FromFloat64(5.960994366302139e-20)),
		Fe:    int32(FE_INEXACT),
	},
	221: {
		Ffile: __ccgo_ts,
		Fline: int32(237),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.526512829145436e-12),
		Fy:    libc.Float64FromFloat64(8.526512829109085e-12),
		Fdy:   float32(-libc.Float64FromFloat64(9.087075337262541e-20)),
		Fe:    int32(FE_INEXACT),
	},
	222: {
		Ffile: __ccgo_ts,
		Fline: int32(238),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.208633855479165e-12),
		Fy:    libc.Float64FromFloat64(9.208633855436765e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1.2361387073992446e-19)),
		Fe:    int32(FE_INEXACT),
	},
	223: {
		Ffile: __ccgo_ts,
		Fline: int32(239),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.046143902593144e-12),
		Fy:    libc.Float64FromFloat64(9.046143902552228e-12),
		Fdy:   float32(-libc.Float64FromFloat64(3.3773514418788183e-16)),
		Fe:    int32(FE_INEXACT),
	},
	224: {
		Ffile: __ccgo_ts,
		Fline: int32(240),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1084466677898518e-11),
		Fy:    libc.Float64FromFloat64(1.1084466677837085e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.5953618899463796e-19)),
		Fe:    int32(FE_INEXACT),
	},
	225: {
		Ffile: __ccgo_ts,
		Fline: int32(241),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1937117960817181e-11),
		Fy:    libc.Float64FromFloat64(1.1937117960745934e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.4910992316678022e-19)),
		Fe:    int32(FE_INEXACT),
	},
	226: {
		Ffile: __ccgo_ts,
		Fline: int32(242),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.176658770423341e-11),
		Fy:    libc.Float64FromFloat64(1.1766587704164184e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.295752258207279e-19)),
		Fe:    int32(FE_INEXACT),
	},
	227: {
		Ffile: __ccgo_ts,
		Fline: int32(243),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0061285138396762e-11),
		Fy:    libc.Float64FromFloat64(1.0061285138346147e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.7618285302889447e-19)),
		Fe:    int32(FE_INEXACT),
	},
	228: {
		Ffile: __ccgo_ts,
		Fline: int32(244),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2327291909222797e-11),
		Fy:    libc.Float64FromFloat64(1.2327291909146816e-11),
		Fdy:   float32(-libc.Float64FromFloat64(3.4347556588076303e-16)),
		Fe:    int32(FE_INEXACT),
	},
	229: {
		Ffile: __ccgo_ts,
		Fline: int32(245),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.0913936421314844e-11),
		Fy:    libc.Float64FromFloat64(1.0913936421255287e-11),
		Fdy:   float32(-libc.Float64FromFloat64(2.439190190296368e-19)),
		Fe:    int32(FE_INEXACT),
	},
	230: {
		Ffile: __ccgo_ts,
		Fline: int32(246),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.023181539498034e-11),
		Fy:    libc.Float64FromFloat64(1.0231815394927994e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1.8843836098448013e-19)),
		Fe:    int32(FE_INEXACT),
	},
	231: {
		Ffile: __ccgo_ts,
		Fline: int32(247),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.278976924373633e-11),
		Fy:    libc.Float64FromFloat64(1.278976924365454e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.600447694774919e-19)),
		Fe:    int32(FE_INEXACT),
	},
	232: {
		Ffile: __ccgo_ts,
		Fline: int32(248),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.2619238987152461e-11),
		Fy:    libc.Float64FromFloat64(1.2619238987072838e-11),
		Fdy:   float32(-libc.Float64FromFloat64(4.360102095991511e-19)),
		Fe:    int32(FE_INEXACT),
	},
	233: {
		Ffile: __ccgo_ts,
		Fline: int32(249),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.449507180957608e-11),
		Fy:    libc.Float64FromFloat64(1.4495071809471026e-11),
		Fdy:   float32(-libc.Float64FromFloat64(7.589944602990565e-19)),
		Fe:    int32(FE_INEXACT),
	},
	234: {
		Ffile: __ccgo_ts,
		Fline: int32(250),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4324541552992017e-11),
		Fy:    libc.Float64FromFloat64(1.4324541552889421e-11),
		Fdy:   float32(-libc.Float64FromFloat64(7.238955325472064e-19)),
		Fe:    int32(FE_INEXACT),
	},
	235: {
		Ffile: __ccgo_ts,
		Fline: int32(251),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.379164112062645e-12),
		Fy:    libc.Float64FromFloat64(9.379164112018661e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1.3303711227812855e-19)),
		Fe:    int32(FE_INEXACT),
	},
	236: {
		Ffile: __ccgo_ts,
		Fline: int32(252),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3471890270071997e-11),
		Fy:    libc.Float64FromFloat64(1.3471890269981251e-11),
		Fdy:   float32(-libc.Float64FromFloat64(5.663209345783049e-19)),
		Fe:    int32(FE_INEXACT),
	},
	237: {
		Ffile: __ccgo_ts,
		Fline: int32(253),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.3642420526655962e-11),
		Fy:    libc.Float64FromFloat64(1.3642420526562904e-11),
		Fdy:   float32(-libc.Float64FromFloat64(5.955435712585782e-19)),
		Fe:    int32(FE_INEXACT),
	},
	238: {
		Ffile: __ccgo_ts,
		Fline: int32(254),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4551804205711455e-11),
		Fy:    libc.Float64FromFloat64(1.4551804205605575e-11),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	239: {
		Ffile: __ccgo_ts,
		Fline: int32(255),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.8996822551806555e-12),
		Fy:    libc.Float64FromFloat64(3.899682255173052e-12),
		Fdy:   float32(-libc.Float64FromFloat64(2.4754619939823693e-16)),
		Fe:    int32(FE_INEXACT),
	},
	240: {
		Ffile: __ccgo_ts,
		Fline: int32(256),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.6821199366207846e-12),
		Fy:    libc.Float64FromFloat64(3.682119936614005e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	241: {
		Ffile: __ccgo_ts,
		Fline: int32(257),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.1672825123013656e-12),
		Fy:    libc.Float64FromFloat64(5.167282512288014e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	242: {
		Ffile: __ccgo_ts,
		Fline: int32(258),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.840877128908382e-12),
		Fy:    libc.Float64FromFloat64(4.840877128896664e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	243: {
		Ffile: __ccgo_ts,
		Fline: int32(259),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.477094942915233e-12),
		Fy:    libc.Float64FromFloat64(6.477094942894257e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	244: {
		Ffile: __ccgo_ts,
		Fline: int32(260),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.126031644554177e-12),
		Fy:    libc.Float64FromFloat64(7.126031644528787e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	245: {
		Ffile: __ccgo_ts,
		Fline: int32(261),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.8869960368543016e-12),
		Fy:    libc.Float64FromFloat64(1.886996036852521e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	246: {
		Ffile: __ccgo_ts,
		Fline: int32(262),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.1551828251793376e-12),
		Fy:    libc.Float64FromFloat64(2.1551828251770148e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	247: {
		Ffile: __ccgo_ts,
		Fline: int32(263),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2615359756431715e-12),
		Fy:    libc.Float64FromFloat64(2.2615359756406144e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	248: {
		Ffile: __ccgo_ts,
		Fline: int32(264),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.5836412561484577e-12),
		Fy:    libc.Float64FromFloat64(2.58364125614512e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	249: {
		Ffile: __ccgo_ts,
		Fline: int32(265),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.8304940552823426e-12),
		Fy:    libc.Float64FromFloat64(2.8304940552783363e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	250: {
		Ffile: __ccgo_ts,
		Fline: int32(266),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.08289969541178377),
		Fy:    libc.Float64FromFloat64(0.07964234638129533),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	251: {
		Ffile: __ccgo_ts,
		Fline: int32(267),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.08604548379391067),
		Fy:    libc.Float64FromFloat64(0.0825431025812302),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	252: {
		Ffile: __ccgo_ts,
		Fline: int32(268),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.1048229242702921),
		Fy:    libc.Float64FromFloat64(0.09968507259924103),
		Fdy:   float32(-libc.Float64FromFloat64(4.999169396453567e-16)),
		Fe:    int32(FE_INEXACT),
	},
	253: {
		Ffile: __ccgo_ts,
		Fline: int32(269),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.1050460593243985),
		Fy:    libc.Float64FromFloat64(0.09988701674753325),
		Fdy:   float32(-libc.Float64FromFloat64(3.8417724600178587e-16)),
		Fe:    int32(FE_INEXACT),
	},
	254: {
		Ffile: __ccgo_ts,
		Fline: int32(270),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.13105003857866188),
		Fy:    libc.Float64FromFloat64(0.12314643893042172),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	255: {
		Ffile: __ccgo_ts,
		Fline: int32(271),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(9.849670420390535e-13),
		Fy:    libc.Float64FromFloat64(9.849670420385685e-13),
		Fdy:   float32(-libc.Float64FromFloat64(8.551935803055035e-17)),
		Fe:    int32(FE_INEXACT),
	},
	256: {
		Ffile: __ccgo_ts,
		Fline: int32(272),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.805489753492587e-13),
		Fy:    libc.Float64FromFloat64(9.805489753487781e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	257: {
		Ffile: __ccgo_ts,
		Fline: int32(273),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1937117960774433e-12),
		Fy:    libc.Float64FromFloat64(1.1937117960767308e-12),
		Fdy:   float32(-libc.Float64FromFloat64(2.6469779601696886e-22)),
		Fe:    int32(FE_INEXACT),
	},
	258: {
		Ffile: __ccgo_ts,
		Fline: int32(274),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.1647700612942416e-12),
		Fy:    libc.Float64FromFloat64(1.1647700612935633e-12),
		Fdy:   float32(-libc.Float64FromFloat64(5.864722719099583e-16)),
		Fe:    int32(FE_INEXACT),
	},
	259: {
		Ffile: __ccgo_ts,
		Fline: int32(275),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.778912954170593e-12),
		Fy:    libc.Float64FromFloat64(1.7789129541690107e-12),
		Fdy:   float32(-libc.Float64FromFloat64(4.761274899261277e-16)),
		Fe:    int32(FE_INEXACT),
	},
	260: {
		Ffile: __ccgo_ts,
		Fline: int32(276),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4068746168056581e-12),
		Fy:    libc.Float64FromFloat64(1.4068746168046686e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	261: {
		Ffile: __ccgo_ts,
		Fline: int32(277),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.4870557788858266e-12),
		Fy:    libc.Float64FromFloat64(1.4870557788847208e-12),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	262: {
		Ffile: __ccgo_ts,
		Fline: int32(278),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6200374375339033e-12),
		Fy:    libc.Float64FromFloat64(1.620037437532591e-12),
		Fdy:   float32(-libc.Float64FromFloat64(9.26442286059391e-22)),
		Fe:    int32(FE_INEXACT),
	},
	263: {
		Ffile: __ccgo_ts,
		Fline: int32(279),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.6083398116556757e-12),
		Fy:    libc.Float64FromFloat64(1.6083398116543822e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	264: {
		Ffile: __ccgo_ts,
		Fline: int32(280),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.849421637446287e-13),
		Fy:    libc.Float64FromFloat64(4.849421637445111e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	265: {
		Ffile: __ccgo_ts,
		Fline: int32(281),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.033331475609603e-13),
		Fy:    libc.Float64FromFloat64(5.033331475608335e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	266: {
		Ffile: __ccgo_ts,
		Fline: int32(282),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.516881676248313e-13),
		Fy:    libc.Float64FromFloat64(6.516881676246189e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1.7662601007898608e-16)),
		Fe:    int32(FE_INEXACT),
	},
	267: {
		Ffile: __ccgo_ts,
		Fline: int32(283),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.756966115559925e-13),
		Fy:    libc.Float64FromFloat64(7.756966115556917e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	268: {
		Ffile: __ccgo_ts,
		Fline: int32(284),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.690724855242186e-13),
		Fy:    libc.Float64FromFloat64(8.690724855238409e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	269: {
		Ffile: __ccgo_ts,
		Fline: int32(285),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.768799812616742e-13),
		Fy:    libc.Float64FromFloat64(8.768799812612897e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	270: {
		Ffile: __ccgo_ts,
		Fline: int32(286),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.557953848736579e-13),
		Fy:    libc.Float64FromFloat64(2.5579538487362516e-13),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	271: {
		Ffile: __ccgo_ts,
		Fline: int32(287),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.866051436015572e-13),
		Fy:    libc.Float64FromFloat64(2.866051436015161e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	272: {
		Ffile: __ccgo_ts,
		Fline: int32(288),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.65480674844853e-13),
		Fy:    libc.Float64FromFloat64(2.6548067484481774e-13),
		Fdy:   float32(-libc.Float64FromFloat64(1.897811729036742e-16)),
		Fe:    int32(FE_INEXACT),
	},
	273: {
		Ffile: __ccgo_ts,
		Fline: int32(289),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.8993278196914413e-13),
		Fy:    libc.Float64FromFloat64(2.8993278196910207e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	274: {
		Ffile: __ccgo_ts,
		Fline: int32(290),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.1974423109207916e-13),
		Fy:    libc.Float64FromFloat64(3.1974423109202807e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	275: {
		Ffile: __ccgo_ts,
		Fline: int32(291),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6279452078876165e-13),
		Fy:    libc.Float64FromFloat64(3.627945207886958e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	276: {
		Ffile: __ccgo_ts,
		Fline: int32(292),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.263256414561207e-13),
		Fy:    libc.Float64FromFloat64(4.263256414560298e-13),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	277: {
		Ffile: __ccgo_ts,
		Fline: int32(293),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(3.6237679523769487e-13),
		Fy:    libc.Float64FromFloat64(3.6237679523762923e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	278: {
		Ffile: __ccgo_ts,
		Fline: int32(294),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.139086659085962e-13),
		Fy:    libc.Float64FromFloat64(1.1390866590858972e-13),
		Fdy:   float32(-libc.Float64FromFloat64(7.076265707292931e-16)),
		Fe:    int32(FE_INEXACT),
	},
	279: {
		Ffile: __ccgo_ts,
		Fline: int32(295),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.586520391123188e-14),
		Fy:    libc.Float64FromFloat64(7.586520391122899e-14),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	280: {
		Ffile: __ccgo_ts,
		Fline: int32(296),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(4.47980149435476e-14),
		Fy:    libc.Float64FromFloat64(4.479801494354659e-14),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	281: {
		Ffile: __ccgo_ts,
		Fline: int32(297),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.50088302959277e-14),
		Fy:    libc.Float64FromFloat64(4.5008830295926683e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	282: {
		Ffile: __ccgo_ts,
		Fline: int32(298),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.9131948672901936e-14),
		Fy:    libc.Float64FromFloat64(1.913194867290175e-14),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	283: {
		Ffile: __ccgo_ts,
		Fline: int32(299),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.4030109916972704e-14),
		Fy:    libc.Float64FromFloat64(2.4030109916972414e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	284: {
		Ffile: __ccgo_ts,
		Fline: int32(300),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0126792072454743e-14),
		Fy:    libc.Float64FromFloat64(1.0126792072454692e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	285: {
		Ffile: __ccgo_ts,
		Fline: int32(301),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.164835077159059e-14),
		Fy:    libc.Float64FromFloat64(1.164835077159052e-14),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	286: {
		Ffile: __ccgo_ts,
		Fline: int32(302),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.615110447320561e-15),
		Fy:    libc.Float64FromFloat64(4.61511044732055e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	287: {
		Ffile: __ccgo_ts,
		Fline: int32(303),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(5.329070518200761e-15),
		Fy:    libc.Float64FromFloat64(5.329070518200747e-15),
		Fdy:   float32(0),
		Fe:    int32(FE_INEXACT),
	},
	288: {
		Ffile: __ccgo_ts,
		Fline: int32(304),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.884110950420531e-15),
		Fy:    libc.Float64FromFloat64(1.8841109504205295e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	289: {
		Ffile: __ccgo_ts,
		Fline: int32(305),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.038229654744293705),
		Fy:    libc.Float64FromFloat64(0.03751700761728945),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	290: {
		Ffile: __ccgo_ts,
		Fline: int32(306),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.0490877108652771),
		Fy:    libc.Float64FromFloat64(0.047920939708309386),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	291: {
		Ffile: __ccgo_ts,
		Fline: int32(307),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.058135531327070444),
		Fy:    libc.Float64FromFloat64(0.056508426675174195),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	292: {
		Ffile: __ccgo_ts,
		Fline: int32(308),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(0.060205904281236386),
		Fy:    libc.Float64FromFloat64(0.05846313858194019),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	293: {
		Ffile: __ccgo_ts,
		Fline: int32(309),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(8.881784197001256e-16),
		Fy:    libc.Float64FromFloat64(8.881784197001251e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	294: {
		Ffile: __ccgo_ts,
		Fline: int32(310),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.930136612989096e-16),
		Fy:    libc.Float64FromFloat64(9.93013661298909e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	295: {
		Ffile: __ccgo_ts,
		Fline: int32(311),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.087791964408415e-15),
		Fy:    libc.Float64FromFloat64(1.0877919644084144e-15),
		Fdy:   float32(-libc.Float64FromFloat64(5.312004074862998e-16)),
		Fe:    int32(FE_INEXACT),
	},
	296: {
		Ffile: __ccgo_ts,
		Fline: int32(312),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.1749496091904417e-15),
		Fy:    libc.Float64FromFloat64(1.1749496091904411e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	297: {
		Ffile: __ccgo_ts,
		Fline: int32(313),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.4728772825180599e-15),
		Fy:    libc.Float64FromFloat64(1.4728772825180587e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	298: {
		Ffile: __ccgo_ts,
		Fline: int32(314),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.6616296724220907e-15),
		Fy:    libc.Float64FromFloat64(1.6616296724220891e-15),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	299: {
		Ffile: __ccgo_ts,
		Fline: int32(315),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(5.438959822042076e-16),
		Fy:    libc.Float64FromFloat64(5.438959822042074e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	300: {
		Ffile: __ccgo_ts,
		Fline: int32(316),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.280369834735102e-16),
		Fy:    libc.Float64FromFloat64(6.280369834735099e-16),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	301: {
		Ffile: __ccgo_ts,
		Fline: int32(317),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.021666937153402e-16),
		Fy:    libc.Float64FromFloat64(7.0216669371534e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	302: {
		Ffile: __ccgo_ts,
		Fline: int32(318),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(7.691850745534257e-16),
		Fy:    libc.Float64FromFloat64(7.691850745534254e-16),
		Fdy:   float32(-libc.Float64FromFloat64(3.4762549792671706e-16)),
		Fe:    int32(FE_INEXACT),
	},
	303: {
		Ffile: __ccgo_ts,
		Fline: int32(319),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.01755015332170947),
		Fy:    libc.Float64FromFloat64(0.0173979278537315),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	304: {
		Ffile: __ccgo_ts,
		Fline: int32(320),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.012193423671798213),
		Fy:    libc.Float64FromFloat64(0.012119682712713818),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	305: {
		Ffile: __ccgo_ts,
		Fline: int32(321),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.005373041811775324),
		Fy:    libc.Float64FromFloat64(0.005358658520965404),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	306: {
		Ffile: __ccgo_ts,
		Fline: int32(322),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.005971983782683648),
		Fy:    libc.Float64FromFloat64(0.0059542221671725825),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	307: {
		Ffile: __ccgo_ts,
		Fline: int32(323),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.709361089888903),
		Fy:    libc.Float64FromFloat64(1.3108596488210371),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	308: {
		Ffile: __ccgo_ts,
		Fline: int32(324),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.428618981120854),
		Fy:    libc.Float64FromFloat64(1.6916847703128641),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	309: {
		Ffile: __ccgo_ts,
		Fline: int32(325),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.121327990593409),
		Fy:    libc.Float64FromFloat64(1.9630942235648072),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	310: {
		Ffile: __ccgo_ts,
		Fline: int32(326),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(8.345898577182753),
		Fy:    libc.Float64FromFloat64(2.23493759221665),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	311: {
		Ffile: __ccgo_ts,
		Fline: int32(327),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(11.464634503398177),
		Fy:    libc.Float64FromFloat64(2.52289539471636),
		Fdy:   float32(-libc.Float64FromFloat64(6.821492490439822e-17)),
		Fe:    int32(FE_INEXACT),
	},
	312: {
		Ffile: __ccgo_ts,
		Fline: int32(328),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(13.13476260116266),
		Fy:    libc.Float64FromFloat64(2.6486371958975217),
		Fdy:   float32(-libc.Float64FromFloat64(2.1028643765838193e-16)),
		Fe:    int32(FE_INEXACT),
	},
	313: {
		Ffile: __ccgo_ts,
		Fline: int32(329),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(18.85094962074969),
		Fy:    libc.Float64FromFloat64(2.9882518458251672),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	314: {
		Ffile: __ccgo_ts,
		Fline: int32(330),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(36.27248338224634),
		Fy:    libc.Float64FromFloat64(3.6182553434754277),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	315: {
		Ffile: __ccgo_ts,
		Fline: int32(331),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(66.36994795754265),
		Fy:    libc.Float64FromFloat64(4.210199042485149),
		Fdy:   float32(-libc.Float64FromFloat64(5.59162182685026e-17)),
		Fe:    int32(FE_INEXACT),
	},
	316: {
		Ffile: __ccgo_ts,
		Fline: int32(332),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(276.58366627903007),
		Fy:    libc.Float64FromFloat64(5.626122387849079),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	317: {
		Ffile: __ccgo_ts,
		Fline: int32(333),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(18507.534852627046),
		Fy:    libc.Float64FromFloat64(9.82598724806552),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	318: {
		Ffile: __ccgo_ts,
		Fline: int32(334),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(47480.619195666535),
		Fy:    libc.Float64FromFloat64(10.76809795083097),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	319: {
		Ffile: __ccgo_ts,
		Fline: int32(335),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(2.508292693232171e+06),
		Fy:    libc.Float64FromFloat64(14.735113276443638),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	320: {
		Ffile: __ccgo_ts,
		Fline: int32(336),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(1.5554081802898357e+07),
		Fy:    libc.Float64FromFloat64(16.559833721798675),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	321: {
		Ffile: __ccgo_ts,
		Fline: int32(337),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.2471337369034297e+09),
		Fy:    libc.Float64FromFloat64(20.94411374561095),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	322: {
		Ffile: __ccgo_ts,
		Fline: int32(338),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(6.987562283919435e+12),
		Fy:    libc.Float64FromFloat64(29.575152867990138),
		Fdy:   float32(-libc.Float64FromFloat64(7.709482127671828e-18)),
		Fe:    int32(FE_INEXACT),
	},
	323: {
		Ffile: __ccgo_ts,
		Fline: int32(339),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    libc.Float64FromFloat64(3.0799488994606656e+14),
		Fy:    libc.Float64FromFloat64(33.361104307680286),
		Fdy:   float32(-libc.Float64FromFloat64(1)),
		Fe:    int32(FE_INEXACT),
	},
	324: {
		Ffile: __ccgo_ts,
		Fline: int32(340),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(6.08494980157565e+18),
		Fy:    libc.Float64FromFloat64(43.25235015076265),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	325: {
		Ffile: __ccgo_ts,
		Fline: int32(341),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(2.2021387043525986e+21),
		Fy:    libc.Float64FromFloat64(49.143715979360856),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	326: {
		Ffile: __ccgo_ts,
		Fline: int32(342),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.860414956831622e+23),
		Fy:    libc.Float64FromFloat64(55.24798539146341),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	327: {
		Ffile: __ccgo_ts,
		Fline: int32(343),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(7.176922076666602e+24),
		Fy:    libc.Float64FromFloat64(57.23291284287763),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	328: {
		Ffile: __ccgo_ts,
		Fline: int32(344),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1.0293733571960176e+28),
		Fy:    libc.Float64FromFloat64(64.50133282985804),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	329: {
		Ffile: __ccgo_ts,
		Fline: int32(345),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.45208966738046846),
		Fy:    -libc.Float64FromFloat64(0.6016436320291095),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	330: {
		Ffile: __ccgo_ts,
		Fline: int32(346),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.4541539066674026),
		Fy:    -libc.Float64FromFloat64(0.6054182233161913),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	331: {
		Ffile: __ccgo_ts,
		Fline: int32(347),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.001350857658673194),
		Fy:    -libc.Float64FromFloat64(0.0013517708894026576),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	332: {
		Ffile: __ccgo_ts,
		Fline: int32(348),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0018758576649691486),
		Fy:    -libc.Float64FromFloat64(0.0018776192893412082),
		Fdy:   float32(1.454376481561519e-16),
		Fe:    int32(FE_INEXACT),
	},
	333: {
		Ffile: __ccgo_ts,
		Fline: int32(349),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0019323732509088777),
		Fy:    -libc.Float64FromFloat64(0.0019342426927938326),
		Fdy:   float32(7.797705373688692e-16),
		Fe:    int32(FE_INEXACT),
	},
	334: {
		Ffile: __ccgo_ts,
		Fline: int32(350),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0016489268452308842),
		Fy:    -libc.Float64FromFloat64(0.0016502878214072086),
		Fdy:   float32(5.348801303674093e-17),
		Fe:    int32(FE_INEXACT),
	},
	335: {
		Ffile: __ccgo_ts,
		Fline: int32(351),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.001663308515293271),
		Fy:    -libc.Float64FromFloat64(0.0016646933487183175),
		Fdy:   float32(3.932909499373277e-16),
		Fe:    int32(FE_INEXACT),
	},
	336: {
		Ffile: __ccgo_ts,
		Fline: int32(352),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0004907184920694575),
		Fy:    -libc.Float64FromFloat64(0.0004908389337922857),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	337: {
		Ffile: __ccgo_ts,
		Fline: int32(353),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0005225804863727815),
		Fy:    -libc.Float64FromFloat64(0.0005227170791443679),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	338: {
		Ffile: __ccgo_ts,
		Fline: int32(354),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0005375306628042406),
		Fy:    -libc.Float64FromFloat64(0.0005376751842030773),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	339: {
		Ffile: __ccgo_ts,
		Fline: int32(355),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0008175595548966595),
		Fy:    -libc.Float64FromFloat64(0.0008178939389745818),
		Fdy:   float32(3.7201144412927036e-16),
		Fe:    int32(FE_INEXACT),
	},
	340: {
		Ffile: __ccgo_ts,
		Fline: int32(356),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0009011491056111168),
		Fy:    -libc.Float64FromFloat64(0.0009015553845633363),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	341: {
		Ffile: __ccgo_ts,
		Fline: int32(357),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00024451620306825895),
		Fy:    -libc.Float64FromFloat64(0.0002445461020289933),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	342: {
		Ffile: __ccgo_ts,
		Fline: int32(358),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00026076482195816584),
		Fy:    -libc.Float64FromFloat64(0.00026079882701602827),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	343: {
		Ffile: __ccgo_ts,
		Fline: int32(359),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0002601027924107567),
		Fy:    -libc.Float64FromFloat64(0.00026013662500882926),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	344: {
		Ffile: __ccgo_ts,
		Fline: int32(360),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0003489019383173625),
		Fy:    -libc.Float64FromFloat64(0.00034896281875992474),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	345: {
		Ffile: __ccgo_ts,
		Fline: int32(361),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00029372200672236166),
		Fy:    -libc.Float64FromFloat64(0.00029376515147956146),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	346: {
		Ffile: __ccgo_ts,
		Fline: int32(362),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0003709229183728928),
		Fy:    -libc.Float64FromFloat64(0.0003709917272943099),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	347: {
		Ffile: __ccgo_ts,
		Fline: int32(363),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0004105702852934708),
		Fy:    -libc.Float64FromFloat64(0.00041065459234982494),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	348: {
		Ffile: __ccgo_ts,
		Fline: int32(364),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.00016605687378473682),
		Fy:    -libc.Float64FromFloat64(0.00016607066275392562),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	349: {
		Ffile: __ccgo_ts,
		Fline: int32(365),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00010691478129500695),
		Fy:    -libc.Float64FromFloat64(0.00010692049708764207),
		Fdy:   float32(5.547978454242982e-16),
		Fe:    int32(FE_INEXACT),
	},
	350: {
		Ffile: __ccgo_ts,
		Fline: int32(366),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00010471889098123223),
		Fy:    -libc.Float64FromFloat64(0.00010472437438711053),
		Fdy:   float32(7.852012361706309e-16),
		Fe:    int32(FE_INEXACT),
	},
	351: {
		Ffile: __ccgo_ts,
		Fline: int32(367),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.00010213455468851976),
		Fy:    -libc.Float64FromFloat64(0.00010213977077731543),
		Fdy:   float32(4.250924843046352e-16),
		Fe:    int32(FE_INEXACT),
	},
	352: {
		Ffile: __ccgo_ts,
		Fline: int32(368),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.7664368958092245e-05),
		Fy:    -libc.Float64FromFloat64(3.766507827824753e-05),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	353: {
		Ffile: __ccgo_ts,
		Fline: int32(369),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.678680048406363e-05),
		Fy:    -libc.Float64FromFloat64(1.6786941383975714e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	354: {
		Ffile: __ccgo_ts,
		Fline: int32(370),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.905985419099203e-05),
		Fy:    -libc.Float64FromFloat64(1.9060035832320964e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	355: {
		Ffile: __ccgo_ts,
		Fline: int32(371),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.588552128600496e-05),
		Fy:    -libc.Float64FromFloat64(2.5885856321892817e-05),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	356: {
		Ffile: __ccgo_ts,
		Fline: int32(372),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.186637734791678e-06),
		Fy:    -libc.Float64FromFloat64(8.186671245493272e-06),
		Fdy:   float32(3.641670190651891e-16),
		Fe:    int32(FE_INEXACT),
	},
	357: {
		Ffile: __ccgo_ts,
		Fline: int32(373),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0129041261286663e-05),
		Fy:    -libc.Float64FromFloat64(1.0129092560371506e-05),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	358: {
		Ffile: __ccgo_ts,
		Fline: int32(374),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.2786925829355926e-06),
		Fy:    -libc.Float64FromFloat64(4.278701736566812e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	359: {
		Ffile: __ccgo_ts,
		Fline: int32(375),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.4506177784950085e-06),
		Fy:    -libc.Float64FromFloat64(5.45063263316607e-06),
		Fdy:   float32(5.789977884855903e-16),
		Fe:    int32(FE_INEXACT),
	},
	360: {
		Ffile: __ccgo_ts,
		Fline: int32(376),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.405394354326238e-06),
		Fy:    -libc.Float64FromFloat64(2.4053972472918767e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	361: {
		Ffile: __ccgo_ts,
		Fline: int32(377),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5465408014785088e-06),
		Fy:    -libc.Float64FromFloat64(3.546547090469206e-06),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	362: {
		Ffile: __ccgo_ts,
		Fline: int32(378),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.9756487753565585e-06),
		Fy:    -libc.Float64FromFloat64(2.9756532026081582e-06),
		Fdy:   float32(2.459718827866464e-16),
		Fe:    int32(FE_INEXACT),
	},
	363: {
		Ffile: __ccgo_ts,
		Fline: int32(379),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.3652528224179302e-06),
		Fy:    -libc.Float64FromFloat64(3.365258484893913e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	364: {
		Ffile: __ccgo_ts,
		Fline: int32(380),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.23447363604033442),
		Fy:    -libc.Float64FromFloat64(0.2671916242949979),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	365: {
		Ffile: __ccgo_ts,
		Fline: int32(381),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.931060550754498e-07),
		Fy:    -libc.Float64FromFloat64(9.931065482055945e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	366: {
		Ffile: __ccgo_ts,
		Fline: int32(382),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0375324600242228e-06),
		Fy:    -libc.Float64FromFloat64(1.0375329982613978e-06),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	367: {
		Ffile: __ccgo_ts,
		Fline: int32(383),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.826923678010034e-06),
		Fy:    -libc.Float64FromFloat64(1.826925346837129e-06),
		Fdy:   float32(7.830684601490038e-17),
		Fe:    int32(FE_INEXACT),
	},
	368: {
		Ffile: __ccgo_ts,
		Fline: int32(384),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4855358716029911e-06),
		Fy:    -libc.Float64FromFloat64(1.485536975012497e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	369: {
		Ffile: __ccgo_ts,
		Fline: int32(385),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.869832931907907e-06),
		Fy:    -libc.Float64FromFloat64(1.869834680047683e-06),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	370: {
		Ffile: __ccgo_ts,
		Fline: int32(386),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.652667412889159e-07),
		Fy:    -libc.Float64FromFloat64(6.652669625789325e-07),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	371: {
		Ffile: __ccgo_ts,
		Fline: int32(387),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(8.496686359083757e-07),
		Fy:    -libc.Float64FromFloat64(8.496689968769756e-07),
		Fdy:   float32(5.690974027002861e-16),
		Fe:    int32(FE_INEXACT),
	},
	372: {
		Ffile: __ccgo_ts,
		Fline: int32(388),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.354206196269272e-07),
		Fy:    -libc.Float64FromFloat64(3.3542067588043583e-07),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	373: {
		Ffile: __ccgo_ts,
		Fline: int32(389),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.6324483144804217e-07),
		Fy:    -libc.Float64FromFloat64(3.6324489742146193e-07),
		Fdy:   float32(1.1235322945066858e-16),
		Fe:    int32(FE_INEXACT),
	},
	374: {
		Ffile: __ccgo_ts,
		Fline: int32(390),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.2820558408034993e-07),
		Fy:    -libc.Float64FromFloat64(4.282056757603872e-07),
		Fdy:   float32(3.0010152328093247e-16),
		Fe:    int32(FE_INEXACT),
	},
	375: {
		Ffile: __ccgo_ts,
		Fline: int32(391),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.8364139610194277e-07),
		Fy:    -libc.Float64FromFloat64(1.8364141296402602e-07),
		Fdy:   float32(5.615908908426e-16),
		Fe:    int32(FE_INEXACT),
	},
	376: {
		Ffile: __ccgo_ts,
		Fline: int32(392),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.984342368815845e-08),
		Fy:    -libc.Float64FromFloat64(5.98434254787762e-08),
		Fdy:   float32(3.5425935409049434e-17),
		Fe:    int32(FE_INEXACT),
	},
	377: {
		Ffile: __ccgo_ts,
		Fline: int32(393),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.055673588271345e-08),
		Fy:    -libc.Float64FromFloat64(6.055673771627266e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	378: {
		Ffile: __ccgo_ts,
		Fline: int32(394),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.167885169155561e-08),
		Fy:    -libc.Float64FromFloat64(6.167885359369605e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	379: {
		Ffile: __ccgo_ts,
		Fline: int32(395),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0460667493260404e-07),
		Fy:    -libc.Float64FromFloat64(1.0460668040388263e-07),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	380: {
		Ffile: __ccgo_ts,
		Fline: int32(396),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0150202705332658e-07),
		Fy:    -libc.Float64FromFloat64(1.0150203220465767e-07),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	381: {
		Ffile: __ccgo_ts,
		Fline: int32(397),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.452413653079915e-08),
		Fy:    -libc.Float64FromFloat64(8.452414010296419e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	382: {
		Ffile: __ccgo_ts,
		Fline: int32(398),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.452696356277372e-08),
		Fy:    -libc.Float64FromFloat64(7.4526966339908e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	383: {
		Ffile: __ccgo_ts,
		Fline: int32(399),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.213609793185516e-08),
		Fy:    -libc.Float64FromFloat64(9.21361021763857e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	384: {
		Ffile: __ccgo_ts,
		Fline: int32(400),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.124275859149637e-08),
		Fy:    -libc.Float64FromFloat64(4.124275944197896e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	385: {
		Ffile: __ccgo_ts,
		Fline: int32(401),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.58359350798313e-08),
		Fy:    -libc.Float64FromFloat64(5.583593663865718e-08),
		Fdy:   float32(1.6834912175577228e-17),
		Fe:    int32(FE_INEXACT),
	},
	386: {
		Ffile: __ccgo_ts,
		Fline: int32(402),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.5230796998418836e-08),
		Fy:    -libc.Float64FromFloat64(4.523079802133136e-08),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	387: {
		Ffile: __ccgo_ts,
		Fline: int32(403),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.82652777368832e-08),
		Fy:    -libc.Float64FromFloat64(4.8265278901651754e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	388: {
		Ffile: __ccgo_ts,
		Fline: int32(404),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.673697371809689e-08),
		Fy:    -libc.Float64FromFloat64(5.673697532763905e-08),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	389: {
		Ffile: __ccgo_ts,
		Fline: int32(405),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.519784540795373e-08),
		Fy:    -libc.Float64FromFloat64(2.5197845725419443e-08),
		Fdy:   float32(5.418403589136755e-17),
		Fe:    int32(FE_INEXACT),
	},
	390: {
		Ffile: __ccgo_ts,
		Fline: int32(406),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.5318060731846662e-08),
		Fy:    -libc.Float64FromFloat64(2.5318061052348765e-08),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	391: {
		Ffile: __ccgo_ts,
		Fline: int32(407),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.498367183174066e-09),
		Fy:    -libc.Float64FromFloat64(9.498367228283556e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	392: {
		Ffile: __ccgo_ts,
		Fline: int32(408),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.964171807342105e-09),
		Fy:    -libc.Float64FromFloat64(5.964171825127777e-09),
		Fdy:   float32(6.98459186080719e-16),
		Fe:    int32(FE_INEXACT),
	},
	393: {
		Ffile: __ccgo_ts,
		Fline: int32(409),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.0949573394192265e-09),
		Fy:    -libc.Float64FromFloat64(6.0949573579934795e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	394: {
		Ffile: __ccgo_ts,
		Fline: int32(410),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.809339015539997e-09),
		Fy:    -libc.Float64FromFloat64(6.809339038723546e-09),
		Fdy:   float32(1.3406308093549032e-17),
		Fe:    int32(FE_INEXACT),
	},
	395: {
		Ffile: __ccgo_ts,
		Fline: int32(411),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.807049201182368e-09),
		Fy:    -libc.Float64FromFloat64(2.807049205122131e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	396: {
		Ffile: __ccgo_ts,
		Fline: int32(412),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.8377256374095403e-09),
		Fy:    -libc.Float64FromFloat64(2.8377256414358837e-09),
		Fdy:   float32(6.175494342886857e-16),
		Fe:    int32(FE_INEXACT),
	},
	397: {
		Ffile: __ccgo_ts,
		Fline: int32(413),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.118160927754504e-09),
		Fy:    -libc.Float64FromFloat64(2.118160929997807e-09),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	398: {
		Ffile: __ccgo_ts,
		Fline: int32(414),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.9033703372263344e-09),
		Fy:    -libc.Float64FromFloat64(1.9033703390377434e-09),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	399: {
		Ffile: __ccgo_ts,
		Fline: int32(415),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.1350187346401479),
		Fy:    -libc.Float64FromFloat64(0.14504743082833094),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	400: {
		Ffile: __ccgo_ts,
		Fline: int32(416),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.19336376370148187),
		Fy:    -libc.Float64FromFloat64(0.21488247280794578),
		Fdy:   float32(1.4140555956898462e-16),
		Fe:    int32(FE_INEXACT),
	},
	401: {
		Ffile: __ccgo_ts,
		Fline: int32(417),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.19490627648363784),
		Fy:    -libc.Float64FromFloat64(0.216796581612239),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	402: {
		Ffile: __ccgo_ts,
		Fline: int32(418),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.20190197133932297),
		Fy:    -libc.Float64FromFloat64(0.22552384614372428),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	403: {
		Ffile: __ccgo_ts,
		Fline: int32(419),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.21282135934419988),
		Fy:    -libc.Float64FromFloat64(0.2393000669254704),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	404: {
		Ffile: __ccgo_ts,
		Fline: int32(420),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.6393927072350511e-09),
		Fy:    -libc.Float64FromFloat64(1.6393927085788553e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	405: {
		Ffile: __ccgo_ts,
		Fline: int32(421),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7500338959050033e-09),
		Fy:    -libc.Float64FromFloat64(1.7500338974363125e-09),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	406: {
		Ffile: __ccgo_ts,
		Fline: int32(422),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.998220583270616e-10),
		Fy:    -libc.Float64FromFloat64(4.998220584519726e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	407: {
		Ffile: __ccgo_ts,
		Fline: int32(423),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.293781370949918e-10),
		Fy:    -libc.Float64FromFloat64(6.293781372930501e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	408: {
		Ffile: __ccgo_ts,
		Fline: int32(424),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.060703692811513e-10),
		Fy:    -libc.Float64FromFloat64(8.06070369606026e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	409: {
		Ffile: __ccgo_ts,
		Fline: int32(425),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.066451676104171e-10),
		Fy:    -libc.Float64FromFloat64(9.066451680214198e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	410: {
		Ffile: __ccgo_ts,
		Fline: int32(426),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.467512511429239e-10),
		Fy:    -libc.Float64FromFloat64(2.4675125117336696e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	411: {
		Ffile: __ccgo_ts,
		Fline: int32(427),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.6922005514253674e-10),
		Fy:    -libc.Float64FromFloat64(3.692200552106985e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	412: {
		Ffile: __ccgo_ts,
		Fline: int32(428),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.108261347901489e-10),
		Fy:    -libc.Float64FromFloat64(3.1082613483845535e-10),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	413: {
		Ffile: __ccgo_ts,
		Fline: int32(429),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.2661276782851193e-10),
		Fy:    -libc.Float64FromFloat64(3.2661276788184983e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	414: {
		Ffile: __ccgo_ts,
		Fline: int32(430),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2005330062922226e-10),
		Fy:    -libc.Float64FromFloat64(1.2005330063642863e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	415: {
		Ffile: __ccgo_ts,
		Fline: int32(431),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.234639057605939e-10),
		Fy:    -libc.Float64FromFloat64(1.2346390576821554e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	416: {
		Ffile: __ccgo_ts,
		Fline: int32(432),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1664269549784287e-10),
		Fy:    -libc.Float64FromFloat64(1.166426955046456e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	417: {
		Ffile: __ccgo_ts,
		Fline: int32(433),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2687451089195778e-10),
		Fy:    -libc.Float64FromFloat64(1.2687451090000632e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	418: {
		Ffile: __ccgo_ts,
		Fline: int32(434),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.1800693755039556e-10),
		Fy:    -libc.Float64FromFloat64(1.1800693755735835e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	419: {
		Ffile: __ccgo_ts,
		Fline: int32(435),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2482814781314038e-10),
		Fy:    -libc.Float64FromFloat64(1.2482814782093139e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	420: {
		Ffile: __ccgo_ts,
		Fline: int32(436),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2141754268177185e-10),
		Fy:    -libc.Float64FromFloat64(1.2141754268914293e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	421: {
		Ffile: __ccgo_ts,
		Fline: int32(437),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3847056833853696e-10),
		Fy:    -libc.Float64FromFloat64(1.3847056834812398e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	422: {
		Ffile: __ccgo_ts,
		Fline: int32(438),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2823875294450116e-10),
		Fy:    -libc.Float64FromFloat64(1.2823875295272372e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	423: {
		Ffile: __ccgo_ts,
		Fline: int32(439),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3164935807585418e-10),
		Fy:    -libc.Float64FromFloat64(1.3164935808451993e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	424: {
		Ffile: __ccgo_ts,
		Fline: int32(440),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3369572115466227e-10),
		Fy:    -libc.Float64FromFloat64(1.3369572116359952e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	425: {
		Ffile: __ccgo_ts,
		Fline: int32(441),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3710632628600288e-10),
		Fy:    -libc.Float64FromFloat64(1.3710632629540193e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	426: {
		Ffile: __ccgo_ts,
		Fline: int32(442),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.3505996320719945e-10),
		Fy:    -libc.Float64FromFloat64(1.3505996321632002e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	427: {
		Ffile: __ccgo_ts,
		Fline: int32(443),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4051693141733574e-10),
		Fy:    -libc.Float64FromFloat64(1.4051693142720822e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	428: {
		Ffile: __ccgo_ts,
		Fline: int32(444),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5756995707388372e-10),
		Fy:    -libc.Float64FromFloat64(1.5756995708629784e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	429: {
		Ffile: __ccgo_ts,
		Fline: int32(445),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.302851160233139e-10),
		Fy:    -libc.Float64FromFloat64(1.3028511603180098e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	430: {
		Ffile: __ccgo_ts,
		Fline: int32(446),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4870238373250297e-10),
		Fy:    -libc.Float64FromFloat64(1.4870238374355914e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	431: {
		Ffile: __ccgo_ts,
		Fline: int32(447),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5893419912639918e-10),
		Fy:    -libc.Float64FromFloat64(1.589341991390292e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	432: {
		Ffile: __ccgo_ts,
		Fline: int32(448),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.507487468112878e-10),
		Fy:    -libc.Float64FromFloat64(1.5074874682265036e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	433: {
		Ffile: __ccgo_ts,
		Fline: int32(449),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.623448042576824e-10),
		Fy:    -libc.Float64FromFloat64(1.623448042708603e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	434: {
		Ffile: __ccgo_ts,
		Fline: int32(450),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4188117346986672e-10),
		Fy:    -libc.Float64FromFloat64(1.4188117347993182e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	435: {
		Ffile: __ccgo_ts,
		Fline: int32(451),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4529177860118872e-10),
		Fy:    -libc.Float64FromFloat64(1.4529177861174354e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	436: {
		Ffile: __ccgo_ts,
		Fline: int32(452),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4392753654866085e-10),
		Fy:    -libc.Float64FromFloat64(1.439275365590184e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	437: {
		Ffile: __ccgo_ts,
		Fline: int32(453),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.473381416799782e-10),
		Fy:    -libc.Float64FromFloat64(1.4733814169083244e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	438: {
		Ffile: __ccgo_ts,
		Fline: int32(454),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6780177246771945e-10),
		Fy:    -libc.Float64FromFloat64(1.6780177248179814e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	439: {
		Ffile: __ccgo_ts,
		Fline: int32(455),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5211298886380946e-10),
		Fy:    -libc.Float64FromFloat64(1.5211298887537862e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	440: {
		Ffile: __ccgo_ts,
		Fline: int32(456),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.5415935194258963e-10),
		Fy:    -libc.Float64FromFloat64(1.5415935195447216e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	441: {
		Ffile: __ccgo_ts,
		Fline: int32(457),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6098056220517005e-10),
		Fy:    -libc.Float64FromFloat64(1.609805622181274e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	442: {
		Ffile: __ccgo_ts,
		Fline: int32(458),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.555235939951082e-10),
		Fy:    -libc.Float64FromFloat64(1.5552359400720197e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	443: {
		Ffile: __ccgo_ts,
		Fline: int32(459),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.691660145202256e-10),
		Fy:    -libc.Float64FromFloat64(1.6916601453453415e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	444: {
		Ffile: __ccgo_ts,
		Fline: int32(460),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6439116733644863e-10),
		Fy:    -libc.Float64FromFloat64(1.6439116734996083e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	445: {
		Ffile: __ccgo_ts,
		Fline: int32(461),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.6575540938895789e-10),
		Fy:    -libc.Float64FromFloat64(1.657554094026953e-10),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	446: {
		Ffile: __ccgo_ts,
		Fline: int32(462),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7222839259124297e-10),
		Fy:    -libc.Float64FromFloat64(1.7222839260607426e-10),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	447: {
		Ffile: __ccgo_ts,
		Fline: int32(463),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.216523012723343e-11),
		Fy:    -libc.Float64FromFloat64(8.2165230130609e-11),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	448: {
		Ffile: __ccgo_ts,
		Fline: int32(464),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.884284375918393e-11),
		Fy:    -libc.Float64FromFloat64(9.884284376406888e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	449: {
		Ffile: __ccgo_ts,
		Fline: int32(465),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.069544618452226e-11),
		Fy:    -libc.Float64FromFloat64(3.0695446184993357e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	450: {
		Ffile: __ccgo_ts,
		Fline: int32(466),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.0013325158206365e-11),
		Fy:    -libc.Float64FromFloat64(3.001332515865676e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	451: {
		Ffile: __ccgo_ts,
		Fline: int32(467),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.1718627723995515e-11),
		Fy:    -libc.Float64FromFloat64(3.1718627724498544e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	452: {
		Ffile: __ccgo_ts,
		Fline: int32(468),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.240074875031063e-11),
		Fy:    -libc.Float64FromFloat64(3.240074875083553e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	453: {
		Ffile: __ccgo_ts,
		Fline: int32(469),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.410605131609707e-11),
		Fy:    -libc.Float64FromFloat64(3.410605131667867e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	454: {
		Ffile: __ccgo_ts,
		Fline: int32(470),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.3423930289782726e-11),
		Fy:    -libc.Float64FromFloat64(3.34239302903413e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	455: {
		Ffile: __ccgo_ts,
		Fline: int32(471),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.853983798713273e-11),
		Fy:    -libc.Float64FromFloat64(3.853983798787538e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	456: {
		Ffile: __ccgo_ts,
		Fline: int32(472),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.0927261579223424e-11),
		Fy:    -libc.Float64FromFloat64(4.092726158006094e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	457: {
		Ffile: __ccgo_ts,
		Fline: int32(473),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.047695594754821e-11),
		Fy:    -libc.Float64FromFloat64(5.0476955948822165e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	458: {
		Ffile: __ccgo_ts,
		Fline: int32(474),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.2632564145000166e-11),
		Fy:    -libc.Float64FromFloat64(4.263256414590893e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	459: {
		Ffile: __ccgo_ts,
		Fline: int32(475),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.706635081601062e-11),
		Fy:    -libc.Float64FromFloat64(4.7066350817118237e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	460: {
		Ffile: __ccgo_ts,
		Fline: int32(476),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.774847184231876e-11),
		Fy:    -libc.Float64FromFloat64(4.774847184345871e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	461: {
		Ffile: __ccgo_ts,
		Fline: int32(477),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.433786671077497e-11),
		Fy:    -libc.Float64FromFloat64(4.4337866711757886e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	462: {
		Ffile: __ccgo_ts,
		Fline: int32(478),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.536104825023892e-11),
		Fy:    -libc.Float64FromFloat64(4.536104825126773e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	463: {
		Ffile: __ccgo_ts,
		Fline: int32(479),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.945377440808775e-11),
		Fy:    -libc.Float64FromFloat64(4.945377440931058e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	464: {
		Ffile: __ccgo_ts,
		Fline: int32(480),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.388756107907804e-11),
		Fy:    -libc.Float64FromFloat64(5.388756108052997e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	465: {
		Ffile: __ccgo_ts,
		Fline: int32(481),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.0245140552912184e-11),
		Fy:    -libc.Float64FromFloat64(4.0245140553722013e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	466: {
		Ffile: __ccgo_ts,
		Fline: int32(482),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.922195901344474e-11),
		Fy:    -libc.Float64FromFloat64(3.922195901421392e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	467: {
		Ffile: __ccgo_ts,
		Fline: int32(483),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.365574568446528e-11),
		Fy:    -libc.Float64FromFloat64(4.3655745685418186e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	468: {
		Ffile: __ccgo_ts,
		Fline: int32(484),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.6043169276547834e-11),
		Fy:    -libc.Float64FromFloat64(4.6043169277607814e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	469: {
		Ffile: __ccgo_ts,
		Fline: int32(485),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.8771653381780385e-11),
		Fy:    -libc.Float64FromFloat64(4.8771653382969716e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	470: {
		Ffile: __ccgo_ts,
		Fline: int32(486),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.2864379539619906e-11),
		Fy:    -libc.Float64FromFloat64(5.286437954101722e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	471: {
		Ffile: __ccgo_ts,
		Fline: int32(487),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.1159076973854796e-11),
		Fy:    -libc.Float64FromFloat64(5.1159076975163415e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	472: {
		Ffile: __ccgo_ts,
		Fline: int32(488),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.2182258513314095e-11),
		Fy:    -libc.Float64FromFloat64(5.218225851467558e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	473: {
		Ffile: __ccgo_ts,
		Fline: int32(489),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5811353881881565e-11),
		Fy:    -libc.Float64FromFloat64(3.5811353882522785e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	474: {
		Ffile: __ccgo_ts,
		Fline: int32(490),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.19504431186897e-11),
		Fy:    -libc.Float64FromFloat64(4.1950443119569615e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	475: {
		Ffile: __ccgo_ts,
		Fline: int32(491),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.5129232855568e-11),
		Fy:    -libc.Float64FromFloat64(3.5129232856185024e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	476: {
		Ffile: __ccgo_ts,
		Fline: int32(492),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.683453542135133e-11),
		Fy:    -libc.Float64FromFloat64(3.683453542202972e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	477: {
		Ffile: __ccgo_ts,
		Fline: int32(493),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.7516656447664123e-11),
		Fy:    -libc.Float64FromFloat64(3.7516656448367867e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	478: {
		Ffile: __ccgo_ts,
		Fline: int32(494),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.627498467114431e-11),
		Fy:    -libc.Float64FromFloat64(5.627498467272774e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	479: {
		Ffile: __ccgo_ts,
		Fline: int32(495),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.559286364484005e-11),
		Fy:    -libc.Float64FromFloat64(5.5592863646385326e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	480: {
		Ffile: __ccgo_ts,
		Fline: int32(496),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.456968210538308e-11),
		Fy:    -libc.Float64FromFloat64(5.4569682106871996e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	481: {
		Ffile: __ccgo_ts,
		Fline: int32(497),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.729816621060012e-11),
		Fy:    -libc.Float64FromFloat64(5.729816621224165e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	482: {
		Ffile: __ccgo_ts,
		Fline: int32(498),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.7980287236903604e-11),
		Fy:    -libc.Float64FromFloat64(5.7980287238584454e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	483: {
		Ffile: __ccgo_ts,
		Fline: int32(499),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.725204785887985e-11),
		Fy:    -libc.Float64FromFloat64(1.7252047859028666e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	484: {
		Ffile: __ccgo_ts,
		Fline: int32(500),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.8321049240749534e-11),
		Fy:    -libc.Float64FromFloat64(1.832104924091736e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	485: {
		Ffile: __ccgo_ts,
		Fline: int32(501),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.4084107803116945e-11),
		Fy:    -libc.Float64FromFloat64(2.4084107803406965e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	486: {
		Ffile: __ccgo_ts,
		Fline: int32(502),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.232390942487842e-11),
		Fy:    -libc.Float64FromFloat64(2.23239094251276e-11),
		Fdy:   float32(3.3340384121209696e-16),
		Fe:    int32(FE_INEXACT),
	},
	487: {
		Ffile: __ccgo_ts,
		Fline: int32(503),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.3712632626326053e-11),
		Fy:    -libc.Float64FromFloat64(2.3712632626607197e-11),
		Fdy:   float32(2.1789274058516635e-16),
		Fe:    int32(FE_INEXACT),
	},
	488: {
		Ffile: __ccgo_ts,
		Fline: int32(504),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.469929948947168e-11),
		Fy:    -libc.Float64FromFloat64(2.4699299489776706e-11),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	489: {
		Ffile: __ccgo_ts,
		Fline: int32(505),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.4416878673150557e-11),
		Fy:    -libc.Float64FromFloat64(2.4416878673448645e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	490: {
		Ffile: __ccgo_ts,
		Fline: int32(506),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.5938667912447236e-11),
		Fy:    -libc.Float64FromFloat64(2.593866791278364e-11),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	491: {
		Ffile: __ccgo_ts,
		Fline: int32(507),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.364239873214453e-12),
		Fy:    -libc.Float64FromFloat64(7.364239873241569e-12),
		Fdy:   float32(6.957250166665413e-16),
		Fe:    int32(FE_INEXACT),
	},
	492: {
		Ffile: __ccgo_ts,
		Fline: int32(508),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.046143902538589e-12),
		Fy:    -libc.Float64FromFloat64(9.046143902579504e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	493: {
		Ffile: __ccgo_ts,
		Fline: int32(509),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.922195901390625e-12),
		Fy:    -libc.Float64FromFloat64(3.922195901398316e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	494: {
		Ffile: __ccgo_ts,
		Fline: int32(510),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.34852154284551e-12),
		Fy:    -libc.Float64FromFloat64(4.348521542854964e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	495: {
		Ffile: __ccgo_ts,
		Fline: int32(511),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.906475391572274e-12),
		Fy:    -libc.Float64FromFloat64(6.906475391596123e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	496: {
		Ffile: __ccgo_ts,
		Fline: int32(512),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.480149750118116e-12),
		Fy:    -libc.Float64FromFloat64(6.480149750139112e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	497: {
		Ffile: __ccgo_ts,
		Fline: int32(513),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.201172825754916e-12),
		Fy:    -libc.Float64FromFloat64(5.201172825768442e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	498: {
		Ffile: __ccgo_ts,
		Fline: int32(514),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.053824108663837e-12),
		Fy:    -libc.Float64FromFloat64(6.053824108682161e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	499: {
		Ffile: __ccgo_ts,
		Fline: int32(515),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.658024299528819e-12),
		Fy:    -libc.Float64FromFloat64(6.658024299550984e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	500: {
		Ffile: __ccgo_ts,
		Fline: int32(516),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.7748471843002735e-12),
		Fy:    -libc.Float64FromFloat64(4.774847184311672e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	501: {
		Ffile: __ccgo_ts,
		Fline: int32(517),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.627498467209437e-12),
		Fy:    -libc.Float64FromFloat64(5.627498467225271e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	502: {
		Ffile: __ccgo_ts,
		Fline: int32(518),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.8410599683070028e-12),
		Fy:    -libc.Float64FromFloat64(1.8410599683086975e-12),
		Fdy:   float32(1.7394308615811729e-16),
		Fe:    int32(FE_INEXACT),
	},
	503: {
		Ffile: __ccgo_ts,
		Fline: int32(519),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.9699340840761668e-12),
		Fy:    -libc.Float64FromFloat64(1.9699340840781067e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	504: {
		Ffile: __ccgo_ts,
		Fline: int32(520),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.606752670496494e-12),
		Fy:    -libc.Float64FromFloat64(2.606752670499891e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	505: {
		Ffile: __ccgo_ts,
		Fline: int32(521),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.5836412561440076e-12),
		Fy:    -libc.Float64FromFloat64(2.5836412561473454e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	506: {
		Ffile: __ccgo_ts,
		Fline: int32(522),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.922810439837678e-12),
		Fy:    -libc.Float64FromFloat64(2.922810439841949e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	507: {
		Ffile: __ccgo_ts,
		Fline: int32(523),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.08275236873532535),
		Fy:    -libc.Float64FromFloat64(0.08637779818161496),
		Fdy:   float32(3.0899396632878373e-16),
		Fe:    int32(FE_INEXACT),
	},
	508: {
		Ffile: __ccgo_ts,
		Fline: int32(524),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.09911403634610279),
		Fy:    -libc.Float64FromFloat64(0.10437659580587227),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	509: {
		Ffile: __ccgo_ts,
		Fline: int32(525),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0066662951214138e-12),
		Fy:    -libc.Float64FromFloat64(1.0066662951219205e-12),
		Fdy:   float32(1.4569842842478782e-16),
		Fe:    int32(FE_INEXACT),
	},
	510: {
		Ffile: __ccgo_ts,
		Fline: int32(526),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0231815394941953e-12),
		Fy:    -libc.Float64FromFloat64(1.0231815394947185e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	511: {
		Ffile: __ccgo_ts,
		Fline: int32(527),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.698843274887871e-13),
		Fy:    -libc.Float64FromFloat64(9.698843274892575e-13),
		Fdy:   float32(7.956526698276833e-16),
		Fe:    int32(FE_INEXACT),
	},
	512: {
		Ffile: __ccgo_ts,
		Fline: int32(528),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.380794976230464e-12),
		Fy:    -libc.Float64FromFloat64(1.3807949762314173e-12),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	513: {
		Ffile: __ccgo_ts,
		Fline: int32(529),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.278976924367635e-12),
		Fy:    -libc.Float64FromFloat64(1.2789769243684528e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	514: {
		Ffile: __ccgo_ts,
		Fline: int32(530),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0658141036397716e-12),
		Fy:    -libc.Float64FromFloat64(1.0658141036403395e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	515: {
		Ffile: __ccgo_ts,
		Fline: int32(531),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4921397450954682e-12),
		Fy:    -libc.Float64FromFloat64(1.4921397450965814e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	516: {
		Ffile: __ccgo_ts,
		Fline: int32(532),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2363443602220648e-12),
		Fy:    -libc.Float64FromFloat64(1.236344360222829e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	517: {
		Ffile: __ccgo_ts,
		Fline: int32(533),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.449507180949904e-12),
		Fy:    -libc.Float64FromFloat64(1.4495071809509544e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	518: {
		Ffile: __ccgo_ts,
		Fline: int32(534),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.705302565823271e-12),
		Fy:    -libc.Float64FromFloat64(1.705302565824725e-12),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	519: {
		Ffile: __ccgo_ts,
		Fline: int32(535),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.662670001677713e-12),
		Fy:    -libc.Float64FromFloat64(1.662670001679095e-12),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	520: {
		Ffile: __ccgo_ts,
		Fline: int32(536),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.924835210192842e-13),
		Fy:    -libc.Float64FromFloat64(4.924835210194054e-13),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	521: {
		Ffile: __ccgo_ts,
		Fline: int32(537),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.550880341145326e-13),
		Fy:    -libc.Float64FromFloat64(6.550880341147472e-13),
		Fdy:   float32(4.4404781111476556e-16),
		Fe:    int32(FE_INEXACT),
	},
	522: {
		Ffile: __ccgo_ts,
		Fline: int32(538),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(3.627945207886739e-13),
		Fy:    -libc.Float64FromFloat64(3.6279452078873974e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	523: {
		Ffile: __ccgo_ts,
		Fline: int32(539),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.2540626382475788e-13),
		Fy:    -libc.Float64FromFloat64(1.2540626382476573e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	524: {
		Ffile: __ccgo_ts,
		Fline: int32(540),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4785436552730705e-13),
		Fy:    -libc.Float64FromFloat64(1.4785436552731798e-13),
		Fdy:   float32(6.567786535100254e-16),
		Fe:    int32(FE_INEXACT),
	},
	525: {
		Ffile: __ccgo_ts,
		Fline: int32(541),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.695433295429486e-14),
		Fy:    -libc.Float64FromFloat64(5.695433295429647e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	526: {
		Ffile: __ccgo_ts,
		Fline: int32(542),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.843009815418733e-14),
		Fy:    -libc.Float64FromFloat64(6.843009815418968e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	527: {
		Ffile: __ccgo_ts,
		Fline: int32(543),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.460698725480866e-14),
		Fy:    -libc.Float64FromFloat64(7.460698725481144e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	528: {
		Ffile: __ccgo_ts,
		Fline: int32(544),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.243885181758519e-14),
		Fy:    -libc.Float64FromFloat64(9.243885181758947e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	529: {
		Ffile: __ccgo_ts,
		Fline: int32(545),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.07160237403795e-14),
		Fy:    -libc.Float64FromFloat64(9.07160237403836e-14),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	530: {
		Ffile: __ccgo_ts,
		Fline: int32(546),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(9.592326932761046e-14),
		Fy:    -libc.Float64FromFloat64(9.592326932761505e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	531: {
		Ffile: __ccgo_ts,
		Fline: int32(547),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0955960721836328e-13),
		Fy:    -libc.Float64FromFloat64(1.0955960721836927e-13),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	532: {
		Ffile: __ccgo_ts,
		Fline: int32(548),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.0863406069049192e-13),
		Fy:    -libc.Float64FromFloat64(1.0863406069049783e-13),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	533: {
		Ffile: __ccgo_ts,
		Fline: int32(549),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(5.329070518200657e-14),
		Fy:    -libc.Float64FromFloat64(5.329070518200798e-14),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	534: {
		Ffile: __ccgo_ts,
		Fline: int32(550),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.5587468005037544e-14),
		Fy:    -libc.Float64FromFloat64(1.5587468005037664e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	535: {
		Ffile: __ccgo_ts,
		Fline: int32(551),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.598721155460217e-14),
		Fy:    -libc.Float64FromFloat64(1.5987211554602295e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	536: {
		Ffile: __ccgo_ts,
		Fline: int32(552),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4755528111219778e-14),
		Fy:    -libc.Float64FromFloat64(1.475552811121989e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	537: {
		Ffile: __ccgo_ts,
		Fline: int32(553),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(2.1316282072802854e-14),
		Fy:    -libc.Float64FromFloat64(2.1316282072803078e-14),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	538: {
		Ffile: __ccgo_ts,
		Fline: int32(554),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.628768063130702e-14),
		Fy:    -libc.Float64FromFloat64(2.6287680631307367e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	539: {
		Ffile: __ccgo_ts,
		Fline: int32(555),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(9.3995968735235e-15),
		Fy:    -libc.Float64FromFloat64(9.399596873523543e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	540: {
		Ffile: __ccgo_ts,
		Fline: int32(556),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0204393146540996e-14),
		Fy:    -libc.Float64FromFloat64(1.0204393146541046e-14),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	541: {
		Ffile: __ccgo_ts,
		Fline: int32(557),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.184977810464563e-14),
		Fy:    -libc.Float64FromFloat64(1.18497781046457e-14),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	542: {
		Ffile: __ccgo_ts,
		Fline: int32(558),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.1443378446363074e-14),
		Fy:    -libc.Float64FromFloat64(1.144337844636314e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	543: {
		Ffile: __ccgo_ts,
		Fline: int32(559),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.3816813636417157e-14),
		Fy:    -libc.Float64FromFloat64(1.3816813636417253e-14),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	544: {
		Ffile: __ccgo_ts,
		Fline: int32(560),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.25955406352393e-15),
		Fy:    -libc.Float64FromFloat64(4.259554063523939e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	545: {
		Ffile: __ccgo_ts,
		Fline: int32(561),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.6151104473205465e-15),
		Fy:    -libc.Float64FromFloat64(4.6151104473205575e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	546: {
		Ffile: __ccgo_ts,
		Fline: int32(562),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(4.945168153041035e-15),
		Fy:    -libc.Float64FromFloat64(4.945168153041047e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	547: {
		Ffile: __ccgo_ts,
		Fline: int32(563),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.102196573270507e-15),
		Fy:    -libc.Float64FromFloat64(5.102196573270519e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	548: {
		Ffile: __ccgo_ts,
		Fline: int32(564),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.864753555590486e-15),
		Fy:    -libc.Float64FromFloat64(4.864753555590498e-15),
		Fdy:   float32(2.942776423729671e-16),
		Fe:    int32(FE_INEXACT),
	},
	549: {
		Ffile: __ccgo_ts,
		Fline: int32(565),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(2.2644195468014687e-15),
		Fy:    -libc.Float64FromFloat64(2.264419546801471e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	550: {
		Ffile: __ccgo_ts,
		Fline: int32(566),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(3.323259344844176e-15),
		Fy:    -libc.Float64FromFloat64(3.3232593448441814e-15),
		Fdy:   float32(6.880171227256456e-16),
		Fe:    int32(FE_INEXACT),
	},
	551: {
		Ffile: __ccgo_ts,
		Fline: int32(567),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.03986301344776611),
		Fy:    -libc.Float64FromFloat64(0.040679310374878326),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	552: {
		Ffile: __ccgo_ts,
		Fline: int32(568),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.05984590951848726),
		Fy:    -libc.Float64FromFloat64(0.06171149110833081),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	553: {
		Ffile: __ccgo_ts,
		Fline: int32(569),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.0877919644084144e-15),
		Fy:    -libc.Float64FromFloat64(1.087791964408415e-15),
		Fdy:   float32(5.565915965087868e-16),
		Fe:    int32(FE_INEXACT),
	},
	554: {
		Ffile: __ccgo_ts,
		Fline: int32(570),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.2560739669470195e-15),
		Fy:    -libc.Float64FromFloat64(1.2560739669470203e-15),
		Fdy:   float32(1.2817336968565555e-16),
		Fe:    int32(FE_INEXACT),
	},
	555: {
		Ffile: __ccgo_ts,
		Fline: int32(571),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.4728772825180583e-15),
		Fy:    -libc.Float64FromFloat64(1.4728772825180593e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	556: {
		Ffile: __ccgo_ts,
		Fline: int32(572),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(1.4043333874306797e-15),
		Fy:    -libc.Float64FromFloat64(1.4043333874306805e-15),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	557: {
		Ffile: __ccgo_ts,
		Fline: int32(573),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.7199501139797023e-15),
		Fy:    -libc.Float64FromFloat64(1.719950113979704e-15),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	558: {
		Ffile: __ccgo_ts,
		Fline: int32(574),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1.6011864169946876e-15),
		Fy:    -libc.Float64FromFloat64(1.6011864169946888e-15),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	559: {
		Ffile: __ccgo_ts,
		Fline: int32(575),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(4.440892098500626e-16),
		Fy:    -libc.Float64FromFloat64(4.440892098500627e-16),
		Fdy:   float32(2.9605945558685534e-16),
		Fe:    int32(FE_INEXACT),
	},
	560: {
		Ffile: __ccgo_ts,
		Fline: int32(576),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(5.438959822042073e-16),
		Fy:    -libc.Float64FromFloat64(5.438959822042075e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	561: {
		Ffile: __ccgo_ts,
		Fline: int32(577),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.021666937153399e-16),
		Fy:    -libc.Float64FromFloat64(7.021666937153401e-16),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	562: {
		Ffile: __ccgo_ts,
		Fline: int32(578),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(6.280369834735099e-16),
		Fy:    -libc.Float64FromFloat64(6.2803698347351e-16),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	563: {
		Ffile: __ccgo_ts,
		Fline: int32(579),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(7.691850745534254e-16),
		Fy:    -libc.Float64FromFloat64(7.691850745534257e-16),
		Fdy:   float32(4.215595363514171e-16),
		Fe:    int32(FE_INEXACT),
	},
	564: {
		Ffile: __ccgo_ts,
		Fline: int32(580),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.308148362110447e-16),
		Fy:    -libc.Float64FromFloat64(8.308148362110451e-16),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	565: {
		Ffile: __ccgo_ts,
		Fline: int32(581),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.019873956868282885),
		Fy:    -libc.Float64FromFloat64(0.020074100147248437),
		Fdy:   float32(1.9310924476277516e-16),
		Fe:    int32(FE_INEXACT),
	},
	566: {
		Ffile: __ccgo_ts,
		Fline: int32(582),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.025540398329229767),
		Fy:    -libc.Float64FromFloat64(0.02587221633684913),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	567: {
		Ffile: __ccgo_ts,
		Fline: int32(583),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.02783163552282319),
		Fy:    -libc.Float64FromFloat64(0.028226275038902182),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	568: {
		Ffile: __ccgo_ts,
		Fline: int32(584),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.02766279108851031),
		Fy:    -libc.Float64FromFloat64(0.0280526119370341),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	569: {
		Ffile: __ccgo_ts,
		Fline: int32(585),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.008409721886780203),
		Fy:    -libc.Float64FromFloat64(0.008445283111580014),
		Fdy:   float32(5.855150188004425e-16),
		Fe:    int32(FE_INEXACT),
	},
	570: {
		Ffile: __ccgo_ts,
		Fline: int32(586),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.011576448386940644),
		Fy:    -libc.Float64FromFloat64(0.011643977133495801),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	571: {
		Ffile: __ccgo_ts,
		Fline: int32(587),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.013433027226352352),
		Fy:    -libc.Float64FromFloat64(0.013524066544926843),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	572: {
		Ffile: __ccgo_ts,
		Fline: int32(588),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0042383330160938116),
		Fy:    -libc.Float64FromFloat64(0.004247340208802245),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	573: {
		Ffile: __ccgo_ts,
		Fline: int32(589),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.005069347943202042),
		Fy:    -libc.Float64FromFloat64(0.005082240677781677),
		Fdy:   float32(3.046708689846754e-17),
		Fe:    int32(FE_INEXACT),
	},
	574: {
		Ffile: __ccgo_ts,
		Fline: int32(590),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.004594183717371136),
		Fy:    -libc.Float64FromFloat64(0.004604769413584198),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	575: {
		Ffile: __ccgo_ts,
		Fline: int32(591),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0054335437464650745),
		Fy:    -libc.Float64FromFloat64(0.005448359136370934),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	576: {
		Ffile: __ccgo_ts,
		Fline: int32(592),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.005532228082852375),
		Fy:    -libc.Float64FromFloat64(0.005547587530805387),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	577: {
		Ffile: __ccgo_ts,
		Fline: int32(593),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.006760351388811436),
		Fy:    -libc.Float64FromFloat64(0.006783306077261629),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	578: {
		Ffile: __ccgo_ts,
		Fline: int32(594),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.0027918601488919927),
		Fy:    -libc.Float64FromFloat64(0.002795764659362213),
		Fdy:   float32(1),
		Fe:    int32(FE_INEXACT),
	},
	579: {
		Ffile: __ccgo_ts,
		Fline: int32(595),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.003239931116045147),
		Fy:    -libc.Float64FromFloat64(0.0032451910571675586),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	580: {
		Ffile: __ccgo_ts,
		Fline: int32(596),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.0027863242146242743),
		Fy:    -libc.Float64FromFloat64(0.002790213241678822),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	581: {
		Ffile: __ccgo_ts,
		Fline: int32(597),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.003093008333255691),
		Fy:    -libc.Float64FromFloat64(0.0030978015697625274),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	582: {
		Ffile: __ccgo_ts,
		Fline: int32(598),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.6627189711152338),
		Fy:    -libc.Float64FromFloat64(1.0868387825378942),
		Fdy:   float32(6.4771487157881235e-16),
		Fe:    int32(FE_INEXACT),
	},
	583: {
		Ffile: __ccgo_ts,
		Fline: int32(599),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.8175705997465633),
		Fy:    -libc.Float64FromFloat64(1.7013920287309274),
		Fdy:   float32(-libc.Float64FromFloat64(0.5)),
		Fe:    int32(FE_INEXACT),
	},
	584: {
		Ffile: __ccgo_ts,
		Fline: int32(600),
		Fr:    int32(FE_TOWARDZERO),
		Fx:    -libc.Float64FromFloat64(0.6627189711152338),
		Fy:    -libc.Float64FromFloat64(1.0868387825378942),
		Fdy:   float32(6.4771487157881235e-16),
		Fe:    int32(FE_INEXACT),
	},
	585: {
		Ffile: __ccgo_ts,
		Fline: int32(601),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.8718035039919387),
		Fy:    -libc.Float64FromFloat64(2.0541910670995107),
		Fdy:   float32(0.5),
		Fe:    int32(FE_INEXACT),
	},
	586: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.06684839057968),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	587: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(4.345239849338305),
		Fy:    libc.Float64FromFloat64(1.6762064170601734),
		Fdy:   float32(0.46188199520111084),
		Fe:    int32(FE_INEXACT),
	},
	588: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(8.38143342755525),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	589: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(6.531673581913484),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	590: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(9.267056966972586),
		Fy:    libc.Float64FromFloat64(2.3289404168523826),
		Fdy:   float32(-libc.Float64FromFloat64(0.411114901304245)),
		Fe:    int32(FE_INEXACT),
	},
	591: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.6619858980995045),
		Fy:    libc.Float64FromFloat64(0.5080132114992477),
		Fdy:   float32(-libc.Float64FromFloat64(0.29306045174598694)),
		Fe:    int32(FE_INEXACT),
	},
	592: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.4066039223853553),
		Fy:    -libc.Float64FromFloat64(0.5218931811663979),
		Fdy:   float32(-libc.Float64FromFloat64(0.25825726985931396)),
		Fe:    int32(FE_INEXACT),
	},
	593: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.5617597462207241),
		Fy:    libc.Float64FromFloat64(0.4458132279488102),
		Fdy:   float32(-libc.Float64FromFloat64(0.13274887204170227)),
		Fe:    int32(FE_INEXACT),
	},
	594: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(9),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0.7741522965913037),
		Fy:    libc.Float64FromFloat64(0.5733227294648414),
		Fdy:   float32(0.02716583013534546),
		Fe:    int32(FE_INEXACT),
	},
	595: {
		Ffile: __ccgo_ts + 25,
		Fline: int32(10),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0.6787637026394024),
		Fy:    -libc.Float64FromFloat64(1.1355782978128564),
		Fdy:   float32(0.2713092863559723),
		Fe:    int32(FE_INEXACT),
	},
	596: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(1),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(0),
		Fy:    libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	597: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(2),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(0),
		Fy:    -libc.Float64FromFloat64(0),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	598: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(3),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(7.888609052210118e-31),
		Fy:    -libc.Float64FromFloat64(7.888609052210118e-31),
		Fdy:   float32(1.7763568394002505e-15),
		Fe:    int32(FE_INEXACT),
	},
	599: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(4),
		Fr:    int32(FE_TONEAREST),
		Fx:    libc.Float64FromFloat64(1),
		Fy:    libc.Float64FromFloat64(0.6931471805599453),
		Fdy:   float32(-libc.Float64FromFloat64(0.2088811695575714)),
		Fe:    int32(FE_INEXACT),
	},
	600: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(5),
		Fr:    int32(FE_TONEAREST),
		Fx:    -libc.Float64FromFloat64(1),
		Fy:    float64(-libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(FE_DIVBYZERO),
	},
	601: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(6),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_inff(nil)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
	602: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(7),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(-libc.X__builtin_inff(nil)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(FE_INVALID),
	},
	603: {
		Ffile: __ccgo_ts + 49,
		Fline: int32(8),
		Fr:    int32(FE_TONEAREST),
		Fx:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fy:    float64(libc.X__builtin_nanf(nil, __ccgo_ts+24)),
		Fdy:   float32(0),
		Fe:    int32(0),
	},
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:18:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:19:1:
	var d float32
	var e, err, i int32
	var p uintptr
	var y float64
	_, _, _, _, _, _ = d, e, err, i, p, y
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:23:12:
	err = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:26:2:
	i = 0
	for {
		if !(uint32(i) < libc.Uint32FromInt64(21744)/libc.Uint32FromInt64(36)) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:27:3:
		p = uintptr(unsafe.Pointer(&t)) + uintptr(i)*36
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:29:3:
		if (*l_l)(unsafe.Pointer(p)).Fr < 0 {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:30:4:
			goto _1
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:31:3:
		libc.Xfesetround(tls, (*l_l)(unsafe.Pointer(p)).Fr)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:32:3:
		feclearexcept(tls, int32(FE_ALL_EXCEPT))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:33:3:
		y = libc.Xlog1pl(tls, (*l_l)(unsafe.Pointer(p)).Fx)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:34:3:
		e = fetestexcept(tls, libc.Int32FromInt32(FE_INEXACT)|libc.Int32FromInt32(FE_INVALID)|libc.Int32FromInt32(FE_DIVBYZERO)|libc.Int32FromInt32(FE_UNDERFLOW)|libc.Int32FromInt32(FE_OVERFLOW))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:36:3:
		if !(checkexcept(tls, e, (*l_l)(unsafe.Pointer(p)).Fe, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:37:4:
			libc.Xprintf(tls, __ccgo_ts+74, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, estr(tls, (*l_l)(unsafe.Pointer(p)).Fe)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:39:4:
			libc.Xprintf(tls, __ccgo_ts+127, libc.VaList(bp+8, estr(tls, e)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:40:4:
			err++
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:42:3:
		d = ulperrl(tls, y, (*l_l)(unsafe.Pointer(p)).Fy, (*l_l)(unsafe.Pointer(p)).Fdy)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:43:3:
		if !(checkulp(tls, d, (*l_l)(unsafe.Pointer(p)).Fr) != 0) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:44:4:
			libc.Xprintf(tls, __ccgo_ts+136, libc.VaList(bp+8, (*l_l)(unsafe.Pointer(p)).Ffile, (*l_l)(unsafe.Pointer(p)).Fline, rstr(tls, (*l_l)(unsafe.Pointer(p)).Fr), (*l_l)(unsafe.Pointer(p)).Fx, (*l_l)(unsafe.Pointer(p)).Fy, y, float64(d), float64(d-(*l_l)(unsafe.Pointer(p)).Fdy), float64((*l_l)(unsafe.Pointer(p)).Fdy)))
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:46:4:
			err++
		}
		goto _1
	_1:
		i++
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/math/log1pl.c:49:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+198, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+208, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 252,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 260,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 268,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 278,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 288,
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
				v2 = __ccgo_ts + 297
			} else {
				v2 = __ccgo_ts + 24
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+299, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 297
		} else {
			v3 = __ccgo_ts + 24
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+304, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 24
	} else {
		v4 = __ccgo_ts + 309
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+311, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 314
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 317
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 320
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 323
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 326
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+329, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+337, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+342, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+385, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+434) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+442) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+454) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+466) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+478) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+487) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+24) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+487) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+493, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "src/math/crlibm/log1p.h\x00\x00src/math/sanity/log1p.h\x00src/math/special/log1p.h\x00%s:%d: bad fp exception: %s log1pl(%La)=%La, want %s\x00 got %s\n\x00%s:%d: %s log1pl(%La) want %La got %La ulperr %.3f = %a + %a\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
