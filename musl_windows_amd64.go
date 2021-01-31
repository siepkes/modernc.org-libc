// Code generated by 'ccgo -D__environ=environ -export-externs X -hide __syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6 -nostdinc -nostdlib -o ../musl_windows_amd64.go -pkgname libc -Iarch/x86_64 -Iarch/generic -Iobj/src/internal -Isrc/include -Isrc/internal -Iobj/include -Iinclude copyright.c src/ctype/isalnum.c src/ctype/isalpha.c src/ctype/isdigit.c src/ctype/islower.c src/ctype/isprint.c src/ctype/isspace.c src/ctype/isxdigit.c src/env/putenv.c src/env/setenv.c src/env/unsetenv.c src/multibyte/wcrtomb.c src/multibyte/wcsrtombs.c src/multibyte/wcstombs.c src/string/strchrnul.c', DO NOT EDIT.

package libc

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

// musl as a whole is licensed under the following standard MIT license:
//
// ----------------------------------------------------------------------
// Copyright © 2005-2020 Rich Felker, et al.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
// ----------------------------------------------------------------------
//
// Authors/contributors include:
//
// A. Wilcox
// Ada Worcester
// Alex Dowad
// Alex Suykov
// Alexander Monakov
// Andre McCurdy
// Andrew Kelley
// Anthony G. Basile
// Aric Belsito
// Arvid Picciani
// Bartosz Brachaczek
// Benjamin Peterson
// Bobby Bingham
// Boris Brezillon
// Brent Cook
// Chris Spiegel
// Clément Vasseur
// Daniel Micay
// Daniel Sabogal
// Daurnimator
// David Carlier
// David Edelsohn
// Denys Vlasenko
// Dmitry Ivanov
// Dmitry V. Levin
// Drew DeVault
// Emil Renner Berthing
// Fangrui Song
// Felix Fietkau
// Felix Janda
// Gianluca Anzolin
// Hauke Mehrtens
// He X
// Hiltjo Posthuma
// Isaac Dunham
// Jaydeep Patil
// Jens Gustedt
// Jeremy Huntwork
// Jo-Philipp Wich
// Joakim Sindholt
// John Spencer
// Julien Ramseier
// Justin Cormack
// Kaarle Ritvanen
// Khem Raj
// Kylie McClain
// Leah Neukirchen
// Luca Barbato
// Luka Perkov
// M Farkas-Dyck (Strake)
// Mahesh Bodapati
// Markus Wichmann
// Masanori Ogino
// Michael Clark
// Michael Forney
// Mikhail Kremnyov
// Natanael Copa
// Nicholas J. Kain
// orc
// Pascal Cuoq
// Patrick Oppenlander
// Petr Hosek
// Petr Skocik
// Pierre Carrier
// Reini Urban
// Rich Felker
// Richard Pennington
// Ryan Fairfax
// Samuel Holland
// Segev Finer
// Shiz
// sin
// Solar Designer
// Stefan Kristiansson
// Stefan O'Rear
// Szabolcs Nagy
// Timo Teräs
// Trutz Behn
// Valentin Ochs
// Will Dietz
// William Haddon
// William Pitcock
//
// Portions of this software are derived from third-party works licensed
// under terms compatible with the above MIT license:
//
// The TRE regular expression implementation (src/regex/reg* and
// src/regex/tre*) is Copyright © 2001-2008 Ville Laurikari and licensed
// under a 2-clause BSD license (license text in the source files). The
// included version has been heavily modified by Rich Felker in 2012, in
// the interests of size, simplicity, and namespace cleanliness.
//
// Much of the math library code (src/math/* and src/complex/*) is
// Copyright © 1993,2004 Sun Microsystems or
// Copyright © 2003-2011 David Schultz or
// Copyright © 2003-2009 Steven G. Kargl or
// Copyright © 2003-2009 Bruce D. Evans or
// Copyright © 2008 Stephen L. Moshier or
// Copyright © 2017-2018 Arm Limited
// and labelled as such in comments in the individual source files. All
// have been licensed under extremely permissive terms.
//
// The ARM memcpy code (src/string/arm/memcpy.S) is Copyright © 2008
// The Android Open Source Project and is licensed under a two-clause BSD
// license. It was taken from Bionic libc, used on Android.
//
// The AArch64 memcpy and memset code (src/string/aarch64/*) are
// Copyright © 1999-2019, Arm Limited.
//
// The implementation of DES for crypt (src/crypt/crypt_des.c) is
// Copyright © 1994 David Burren. It is licensed under a BSD license.
//
// The implementation of blowfish crypt (src/crypt/crypt_blowfish.c) was
// originally written by Solar Designer and placed into the public
// domain. The code also comes with a fallback permissive license for use
// in jurisdictions that may not recognize the public domain.
//
// The smoothsort implementation (src/stdlib/qsort.c) is Copyright © 2011
// Valentin Ochs and is licensed under an MIT-style license.
//
// The x86_64 port was written by Nicholas J. Kain and is licensed under
// the standard MIT terms.
//
// The mips and microblaze ports were originally written by Richard
// Pennington for use in the ellcc project. The original code was adapted
// by Rich Felker for build system and code conventions during upstream
// integration. It is licensed under the standard MIT terms.
//
// The mips64 port was contributed by Imagination Technologies and is
// licensed under the standard MIT terms.
//
// The powerpc port was also originally written by Richard Pennington,
// and later supplemented and integrated by John Spencer. It is licensed
// under the standard MIT terms.
//
// All other files which have no copyright comments are original works
// produced specifically for use as part of this library, written either
// by Rich Felker, the main author of the library, or by one or more
// contibutors listed above. Details on authorship of individual files
// can be found in the git version control history of the project. The
// omission of copyright and license comments in each file is in the
// interest of source tree size.
//
// In addition, permission is hereby granted for all public header files
// (include/* and arch/*/bits/*) and crt files intended to be linked into
// applications (crt/*, ldso/dlstart.c, and arch/*/crt_arch.h) to omit
// the copyright notice and permission notice otherwise required by the
// license, and to use these files without any requirement of
// attribution. These files include substantial contributions from:
//
// Bobby Bingham
// John Spencer
// Nicholas J. Kain
// Rich Felker
// Richard Pennington
// Stefan Kristiansson
// Szabolcs Nagy
//
// all of whom have explicitly granted such permission.
//
// This file previously contained text expressing a belief that most of
// the files covered by the above exception were sufficiently trivial not
// to be subject to copyright, resulting in confusion over whether it
// negated the permissions granted in the license. In the spirit of
// permissive licensing, and of not having licensing issues being an
// obstacle to adoption, that text has been removed.
const ( /* copyright.c:194:1: */
	__musl__copyright__ = 0
)

const ( /* pthread_impl.h:58:1: */
	DT_EXITING  = 0
	DT_JOINABLE = 1
	DT_DETACHED = 2
)

type ptrdiff_t = int64 /* <builtin>:3:26 */

type size_t = uint64 /* <builtin>:9:23 */

type wchar_t = uint16 /* <builtin>:15:24 */

type va_list = uintptr /* <builtin>:51:27 */

type __locale_struct = struct{ cat [6]uintptr } /* alltypes.h:343:9 */

type locale_t = uintptr /* alltypes.h:343:32 */

func Xisalnum(tls *TLS, c int32) int32 { /* isalnum.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisalpha(tls, c)
		}
		return (Bool32((((uint32(c)) | uint32(32)) - uint32('a')) < uint32(26)))
	}() != 0) || (func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0)))
}

func X__isalnum_l(tls *TLS, c int32, l locale_t) int32 { /* isalnum.c:8:5: */
	return Xisalnum(tls, c)
}

func Xisalpha(tls *TLS, c int32) int32 { /* isalpha.c:4:5: */
	return (Bool32(((uint32(c) | uint32(32)) - uint32('a')) < uint32(26)))
}

func X__isalpha_l(tls *TLS, c int32, l locale_t) int32 { /* isalpha.c:9:5: */
	return Xisalpha(tls, c)
}

func Xisdigit(tls *TLS, c int32) int32 { /* isdigit.c:4:5: */
	return (Bool32((uint32(c) - uint32('0')) < uint32(10)))
}

func X__isdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isdigit.c:9:5: */
	return Xisdigit(tls, c)
}

func Xislower(tls *TLS, c int32) int32 { /* islower.c:4:5: */
	return (Bool32((uint32(c) - uint32('a')) < uint32(26)))
}

func X__islower_l(tls *TLS, c int32, l locale_t) int32 { /* islower.c:9:5: */
	return Xislower(tls, c)
}

func Xisprint(tls *TLS, c int32) int32 { /* isprint.c:4:5: */
	return (Bool32((uint32(c) - uint32(0x20)) < uint32(0x5f)))
}

func X__isprint_l(tls *TLS, c int32, l locale_t) int32 { /* isprint.c:9:5: */
	return Xisprint(tls, c)
}

func Xisspace(tls *TLS, c int32) int32 { /* isspace.c:4:5: */
	return (Bool32((c == ' ') || ((uint32(c) - uint32('\t')) < uint32(5))))
}

func X__isspace_l(tls *TLS, c int32, l locale_t) int32 { /* isspace.c:9:5: */
	return Xisspace(tls, c)
}

func Xisxdigit(tls *TLS, c int32) int32 { /* isxdigit.c:3:5: */
	return (Bool32((func() int32 {
		if 0 != 0 {
			return Xisdigit(tls, c)
		}
		return (Bool32(((uint32(c)) - uint32('0')) < uint32(10)))
	}() != 0) || (((uint32(c) | uint32(32)) - uint32('a')) < uint32(6))))
}

func X__isxdigit_l(tls *TLS, c int32, l locale_t) int32 { /* isxdigit.c:8:5: */
	return Xisxdigit(tls, c)
}

type div_t = struct {
	quot int32
	rem  int32
} /* stdlib.h:62:35 */
type ldiv_t = struct {
	quot int32
	rem  int32
} /* stdlib.h:63:36 */
type lldiv_t = struct {
	quot int64
	rem  int64
} /* stdlib.h:64:41 */

type ssize_t = int32 /* alltypes.h:65:15 */

type intptr_t = int32 /* alltypes.h:70:15 */

type off_t = int32 /* alltypes.h:162:16 */

type pid_t = int32 /* alltypes.h:235:13 */

type uid_t = uint32 /* alltypes.h:245:18 */

type gid_t = uint32 /* alltypes.h:250:18 */

type useconds_t = uint32 /* alltypes.h:260:18 */

func X__putenv(tls *TLS, s uintptr, l size_t, r uintptr) int32 { /* putenv.c:8:5: */
	var i size_t
	var newenv uintptr
	var tmp uintptr
	//TODO for (char **e = __environ; *e; e++, i++)
	var e uintptr
	i = uint64(0)
	if !(Environ() != 0) {
		goto __1
	}
	//TODO for (char **e = __environ; *e; e++, i++)
	e = Environ()
__2:
	if !(*(*uintptr)(unsafe.Pointer(e)) != 0) {
		goto __4
	}
	if !(!(Xstrncmp(tls, s, *(*uintptr)(unsafe.Pointer(e)), (l+uint64(1))) != 0)) {
		goto __5
	}
	tmp = *(*uintptr)(unsafe.Pointer(e))
	*(*uintptr)(unsafe.Pointer(e)) = s
	X__env_rm_add(tls, tmp, r)
	return 0
__5:
	;
	goto __3
__3:
	e += 8
	i++
	goto __2
	goto __4
__4:
	;
__1:
	;
	if !(Environ() == oldenv) {
		goto __6
	}
	newenv = Xrealloc(tls, oldenv, (uint64(unsafe.Sizeof(uintptr(0))) * (i + uint64(2))))
	if !(!(newenv != 0)) {
		goto __8
	}
	goto oom
__8:
	;
	goto __7
__6:
	newenv = Xmalloc(tls, (uint64(unsafe.Sizeof(uintptr(0))) * (i + uint64(2))))
	if !(!(newenv != 0)) {
		goto __9
	}
	goto oom
__9:
	;
	if !(i != 0) {
		goto __10
	}
	Xmemcpy(tls, newenv, Environ(), (uint64(unsafe.Sizeof(uintptr(0))) * i))
__10:
	;
	Xfree(tls, oldenv)
__7:
	;
	*(*uintptr)(unsafe.Pointer(newenv + uintptr(i)*8)) = s
	*(*uintptr)(unsafe.Pointer(newenv + uintptr((i+uint64(1)))*8)) = uintptr(0)
	*(*uintptr)(unsafe.Pointer(EnvironP())) = AssignPtrUintptr(uintptr(unsafe.Pointer(&oldenv)), newenv)
	if !(r != 0) {
		goto __11
	}
	X__env_rm_add(tls, uintptr(0), r)
__11:
	;
	return 0
oom:
	Xfree(tls, r)
	return -1
}

var oldenv uintptr /* putenv.c:22:14: */

func Xputenv(tls *TLS, s uintptr) int32 { /* putenv.c:43:5: */
	var l size_t = (size_t((int64(X__strchrnul(tls, s, '=')) - int64(s)) / 1))
	if !(l != 0) || !(int32(*(*int8)(unsafe.Pointer(s + uintptr(l)))) != 0) {
		return Xunsetenv(tls, s)
	}
	return X__putenv(tls, s, l, uintptr(0))
}

func X__env_rm_add(tls *TLS, old uintptr, new uintptr) { /* setenv.c:5:6: */
	//TODO for (size_t i=0; i < env_alloced_n; i++)
	var i size_t = uint64(0)
	for ; i < env_alloced_n; i++ {
		if *(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*8)) == old {
			*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*8)) = new
			Xfree(tls, old)
			return
		} else if !(int32(*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*8))) != 0) && (new != 0) {
			*(*uintptr)(unsafe.Pointer(env_alloced + uintptr(i)*8)) = new
			new = uintptr(0)
		}
	}
	if !(new != 0) {
		return
	}
	var t uintptr = Xrealloc(tls, env_alloced, (uint64(unsafe.Sizeof(uintptr(0))) * (env_alloced_n + uint64(1))))
	if !(t != 0) {
		return
	}
	*(*uintptr)(unsafe.Pointer((AssignPtrUintptr(uintptr(unsafe.Pointer(&env_alloced)), t)) + uintptr(PostIncUint64(&env_alloced_n, 1))*8)) = new
}

var env_alloced uintptr  /* setenv.c:7:14: */
var env_alloced_n size_t /* setenv.c:8:16: */

func Xsetenv(tls *TLS, var1 uintptr, value uintptr, overwrite int32) int32 { /* setenv.c:26:5: */
	var s uintptr
	var l1 size_t
	var l2 size_t

	if (!(var1 != 0) || !(int32(AssignUint64(&l1, (size_t((int64(X__strchrnul(tls, var1, '='))-int64(var1))/1)))) != 0)) || (*(*int8)(unsafe.Pointer(var1 + uintptr(l1))) != 0) {
		(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 22
		return -1
	}
	if !(overwrite != 0) && (Xgetenv(tls, var1) != 0) {
		return 0
	}

	l2 = Xstrlen(tls, value)
	s = Xmalloc(tls, ((l1 + l2) + uint64(2)))
	if !(s != 0) {
		return -1
	}
	Xmemcpy(tls, s, var1, l1)
	*(*int8)(unsafe.Pointer(s + uintptr(l1))) = int8('=')
	Xmemcpy(tls, ((s + uintptr(l1)) + uintptr(1)), value, (l2 + uint64(1)))
	return X__putenv(tls, s, l1, s)
}

func Xunsetenv(tls *TLS, name uintptr) int32 { /* unsetenv.c:9:5: */
	var l size_t = (size_t((int64(X__strchrnul(tls, name, '=')) - int64(name)) / 1))
	if !(l != 0) || (*(*int8)(unsafe.Pointer(name + uintptr(l))) != 0) {
		(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 22
		return -1
	}
	if Environ() != 0 {
		var e uintptr = Environ()
		var eo uintptr = e
		for ; *(*uintptr)(unsafe.Pointer(e)) != 0; e += 8 {
			//TODO if (!strncmp(name, *e, l) && l[*e] == '=')
			if !(Xstrncmp(tls, name, *(*uintptr)(unsafe.Pointer(e)), l) != 0) && (int32(*(*int8)(unsafe.Pointer((*(*uintptr)(unsafe.Pointer(e))) + uintptr(l)))) == '=') {
				X__env_rm_add(tls, *(*uintptr)(unsafe.Pointer(e)), uintptr(0))
			} else if eo != e {
				*(*uintptr)(unsafe.Pointer(PostIncUintptr(&eo, 8))) = *(*uintptr)(unsafe.Pointer(e))
			} else {
				eo += 8
			}
		}
		if eo != e {
			*(*uintptr)(unsafe.Pointer(eo)) = uintptr(0)
		}
	}
	return 0
}

type wint_t = uint32 /* alltypes.h:198:18 */

type wctype_t = uint32 /* alltypes.h:203:23 */

type __mbstate_t = struct {
	__opaque1 uint32
	__opaque2 uint32
} /* alltypes.h:337:9 */

type mbstate_t = __mbstate_t /* alltypes.h:337:63 */

type tm = struct {
	tm_sec    int32
	tm_min    int32
	tm_hour   int32
	tm_mday   int32
	tm_mon    int32
	tm_year   int32
	tm_wday   int32
	tm_yday   int32
	tm_isdst  int32
	tm_gmtoff int32
	tm_zone   uintptr
} /* wchar.h:138:1 */

type uintptr_t = uint32 /* alltypes.h:55:24 */

type int8_t = int8 /* alltypes.h:96:25 */

type int16_t = int16 /* alltypes.h:101:25 */

type int32_t = int32 /* alltypes.h:106:25 */

type int64_t = int32 /* alltypes.h:111:25 */

type intmax_t = int32 /* alltypes.h:116:25 */

type uint8_t = uint8 /* alltypes.h:121:25 */

type uint16_t = uint16 /* alltypes.h:126:25 */

type uint32_t = uint32 /* alltypes.h:131:25 */

type uint64_t = uint32 /* alltypes.h:136:25 */

type uintmax_t = uint32 /* alltypes.h:146:25 */

type int_fast8_t = int8_t   /* stdint.h:22:16 */
type int_fast64_t = int64_t /* stdint.h:23:17 */

type int_least8_t = int8_t   /* stdint.h:25:17 */
type int_least16_t = int16_t /* stdint.h:26:17 */
type int_least32_t = int32_t /* stdint.h:27:17 */
type int_least64_t = int64_t /* stdint.h:28:17 */

type uint_fast8_t = uint8_t   /* stdint.h:30:17 */
type uint_fast64_t = uint64_t /* stdint.h:31:18 */

type uint_least8_t = uint8_t   /* stdint.h:33:18 */
type uint_least16_t = uint16_t /* stdint.h:34:18 */
type uint_least32_t = uint32_t /* stdint.h:35:18 */
type uint_least64_t = uint64_t /* stdint.h:36:18 */

type int_fast16_t = int32_t   /* stdint.h:1:17 */
type int_fast32_t = int32_t   /* stdint.h:2:17 */
type uint_fast16_t = uint32_t /* stdint.h:3:18 */
type uint_fast32_t = uint32_t /* stdint.h:4:18 */

// Upper 6 state bits are a negative integer offset to bound-check next byte
//    equivalent to: ( (b-0x80) | (b+offset) ) & ~0x3f

// Interval [a,b). Either a must be 80 or b must be c0, lower 3 bits clear.

// Arbitrary encoding for representing code units instead of characters.

// Get inline definition of MB_CUR_MAX.

type lconv = struct {
	decimal_point      uintptr
	thousands_sep      uintptr
	grouping           uintptr
	int_curr_symbol    uintptr
	currency_symbol    uintptr
	mon_decimal_point  uintptr
	mon_thousands_sep  uintptr
	mon_grouping       uintptr
	positive_sign      uintptr
	negative_sign      uintptr
	int_frac_digits    int8
	frac_digits        int8
	p_cs_precedes      int8
	p_sep_by_space     int8
	n_cs_precedes      int8
	n_sep_by_space     int8
	p_sign_posn        int8
	n_sign_posn        int8
	int_p_cs_precedes  int8
	int_p_sep_by_space int8
	int_n_cs_precedes  int8
	int_n_sep_by_space int8
	int_p_sign_posn    int8
	int_n_sign_posn    int8
	_                  [2]byte
} /* locale.h:24:1 */

type _G_fpos64_t = struct {
	_        [0]uint64
	__opaque [16]int8
} /* stdio.h:54:9 */

type fpos_t = _G_fpos64_t /* stdio.h:58:3 */

// Support signed or unsigned plain-char

// Implementation choices...

// Arbitrary numbers...

// POSIX/SUS requirements follow. These numbers come directly
// from SUS and have nothing to do with the host system.

type __locale_map = struct {
	__map    uintptr
	map_size size_t
	name     [24]int8
	next     uintptr
} /* alltypes.h:343:9 */

type tls_module = struct {
	next   uintptr
	image  uintptr
	len    size_t
	size   size_t
	align  size_t
	offset size_t
} /* libc.h:14:1 */

type __libc = struct {
	can_do_threads  int8
	threaded        int8
	secure          int8
	need_locks      int8
	threads_minus_1 int32
	auxv            uintptr
	tls_head        uintptr
	tls_size        size_t
	tls_align       size_t
	tls_cnt         size_t
	page_size       size_t
	global_locale   struct{ cat [6]uintptr }
} /* libc.h:20:1 */

type time_t = int32 /* alltypes.h:85:16 */

type clockid_t = int32 /* alltypes.h:214:13 */

type timespec = struct {
	tv_sec  time_t
	tv_nsec int32
} /* alltypes.h:229:1 */

type __pthread = struct {
	self          uintptr
	dtv           uintptr
	prev          uintptr
	next          uintptr
	sysinfo       uintptr_t
	canary        uintptr_t
	canary2       uintptr_t
	tid           int32
	errno_val     int32
	detach_state  int32
	cancel        int32
	canceldisable uint8
	cancelasync   uint8
	tsd_used      uint8 /* unsigned char tsd_used: 1, unsigned char dlerror_flag: 1 */
	_             [1]byte
	map_base      uintptr
	map_size      size_t
	stack         uintptr
	stack_size    size_t
	guard_size    size_t
	result        uintptr
	cancelbuf     uintptr
	tsd           uintptr
	robust_list   struct {
		head    uintptr
		off     int32
		_       [4]byte
		pending uintptr
	}
	timer_id      int32
	_             [4]byte
	locale        locale_t
	killlock      [1]int32
	_             [4]byte
	dlerror_buf   uintptr
	stdio_locks   uintptr
	canary_at_end uintptr_t
	_             [4]byte
	dtv_copy      uintptr
} /* alltypes.h:273:9 */

type pthread_t = uintptr /* alltypes.h:273:26 */

type pthread_once_t = int32 /* alltypes.h:279:13 */

type pthread_key_t = uint32 /* alltypes.h:284:18 */

type pthread_spinlock_t = int32 /* alltypes.h:289:13 */

type pthread_mutexattr_t = struct{ __attr uint32 } /* alltypes.h:294:37 */

type pthread_condattr_t = struct{ __attr uint32 } /* alltypes.h:299:37 */

type pthread_barrierattr_t = struct{ __attr uint32 } /* alltypes.h:304:37 */

type pthread_rwlockattr_t = struct{ __attr [2]uint32 } /* alltypes.h:309:40 */

type __sigset_t = struct{ __bits [32]uint32 } /* alltypes.h:349:9 */

type sigset_t = __sigset_t /* alltypes.h:349:71 */

type pthread_attr_t = struct{ __u struct{ __i [9]int32 } } /* alltypes.h:372:147 */

type pthread_mutex_t = struct {
	__u struct {
		_   [0]uint64
		__i [6]int32
		_   [24]byte
	}
} /* alltypes.h:377:157 */

type pthread_cond_t = struct {
	__u struct {
		_   [0]uint64
		__i [12]int32
	}
} /* alltypes.h:387:112 */

type pthread_rwlock_t = struct {
	__u struct {
		_   [0]uint64
		__i [8]int32
		_   [32]byte
	}
} /* alltypes.h:397:139 */

type pthread_barrier_t = struct {
	__u struct {
		_   [0]uint64
		__i [5]int32
		_   [20]byte
	}
} /* alltypes.h:402:137 */

type sched_param = struct {
	sched_priority int32
	__reserved1    int32
	__reserved2    [2]struct {
		__reserved1 time_t
		__reserved2 int32
	}
	__reserved3 int32
} /* sched.h:19:1 */

type timer_t = uintptr /* alltypes.h:209:14 */

type clock_t = int32 /* alltypes.h:219:14 */

type itimerspec = struct {
	it_interval struct {
		tv_sec  time_t
		tv_nsec int32
	}
	it_value struct {
		tv_sec  time_t
		tv_nsec int32
	}
} /* time.h:80:1 */

type sigevent = struct {
	sigev_value struct {
		_         [0]uint64
		sival_int int32
		_         [4]byte
	}
	sigev_signo             int32
	sigev_notify            int32
	sigev_notify_function   uintptr
	sigev_notify_attributes uintptr
	__pad                   [44]int8
	_                       [4]byte
} /* time.h:107:1 */

type __ptcb = struct {
	__f    uintptr
	__x    uintptr
	__next uintptr
} /* alltypes.h:273:9 */

type sigaltstack = struct {
	ss_sp    uintptr
	ss_flags int32
	_        [4]byte
	ss_size  size_t
} /* signal.h:44:9 */

type stack_t = sigaltstack /* signal.h:44:28 */

type greg_t = int64        /* signal.h:59:19 */
type gregset_t = [23]int64 /* signal.h:59:27 */
type _fpstate = struct {
	cwd       uint16
	swd       uint16
	ftw       uint16
	fop       uint16
	rip       uint64
	rdp       uint64
	mxcsr     uint32
	mxcr_mask uint32
	_st       [8]struct {
		significand [4]uint16
		exponent    uint16
		padding     [3]uint16
	}
	_xmm    [16]struct{ element [4]uint32 }
	padding [24]uint32
} /* signal.h:60:9 */

type fpregset_t = uintptr /* signal.h:71:3 */
type sigcontext = struct {
	r8          uint32
	r9          uint32
	r10         uint32
	r11         uint32
	r12         uint32
	r13         uint32
	r14         uint32
	r15         uint32
	rdi         uint32
	rsi         uint32
	rbp         uint32
	rbx         uint32
	rdx         uint32
	rax         uint32
	rcx         uint32
	rsp         uint32
	rip         uint32
	eflags      uint32
	cs          uint16
	gs          uint16
	fs          uint16
	__pad0      uint16
	err         uint32
	trapno      uint32
	oldmask     uint32
	cr2         uint32
	fpstate     uintptr
	__reserved1 [8]uint32
} /* signal.h:72:1 */

type mcontext_t = struct {
	gregs       gregset_t
	fpregs      fpregset_t
	__reserved1 [8]uint64
} /* signal.h:84:3 */

type __ucontext = struct {
	uc_flags     uint32
	_            [4]byte
	uc_link      uintptr
	uc_stack     stack_t
	uc_mcontext  mcontext_t
	uc_sigmask   sigset_t
	__fpregs_mem [64]uint32
} /* signal.h:97:9 */

type ucontext_t = __ucontext /* signal.h:104:3 */

type sigval = struct {
	_         [0]uint64
	sival_int int32
	_         [4]byte
} /* time.h:107:1 */

type siginfo_t = struct {
	si_signo    int32
	si_errno    int32
	si_code     int32
	_           [4]byte
	__si_fields struct {
		_     [0]uint64
		__pad [116]int8
		_     [4]byte
	}
} /* signal.h:145:3 */

type sigaction = struct {
	__sa_handler struct{ sa_handler uintptr }
	sa_mask      sigset_t
	sa_flags     int32
	_            [4]byte
	sa_restorer  uintptr
} /* signal.h:167:1 */

type sig_t = uintptr /* signal.h:251:14 */

type sig_atomic_t = int32 /* signal.h:269:13 */

type mode_t = uint32 /* alltypes.h:152:18 */

type syscall_arg_t = int32 /* syscall.h:22:14 */

func a_cas(tls *TLS, p uintptr, t int32, s int32) int32 { /* atomic_arch.h:2:19: */
	panic(`arch/x86_64/atomic_arch.h:4:2: assembler statements not supported`)
	return t
}

func a_or(tls *TLS, p uintptr, v int32) { /* atomic_arch.h:46:20: */
	panic(`arch/x86_64/atomic_arch.h:48:2: assembler statements not supported`)
}

func a_or_64(tls *TLS, p uintptr, v uint64_t) { /* atomic_arch.h:62:20: */
	panic(`arch/x86_64/atomic_arch.h:64:2: assembler statements not supported`)
}

func a_ctz_64(tls *TLS, x uint64_t) int32 { /* atomic_arch.h:112:19: */
	panic(`arch/x86_64/atomic_arch.h:114:2: assembler statements not supported`)
	return int32(x)
}

func a_ctz_32(tls *TLS, x uint32_t) int32 { /* atomic.h:256:19: */
	return int32(debruijn32[(((x & -x) * uint32_t(0x076be629)) >> 27)])
}

var debruijn32 = [32]int8{
	int8(0), int8(1), int8(23), int8(2), int8(29), int8(24), int8(19), int8(3), int8(30), int8(27), int8(25), int8(11), int8(20), int8(8), int8(4), int8(13),
	int8(31), int8(22), int8(28), int8(18), int8(26), int8(10), int8(7), int8(12), int8(21), int8(17), int8(9), int8(6), int8(16), int8(5), int8(15), int8(14),
} /* atomic.h:261:20 */

type __timer = struct {
	timerid int32
	_       [4]byte
	thread  pthread_t
} /* pthread_impl.h:64:1 */

func __pthread_self(tls *TLS) uintptr { /* pthread_arch.h:1:30: */
	var self uintptr
	panic(`arch/x86_64/pthread_arch.h:4:2: assembler statements not supported`)
	return self
}

func Xwcrtomb(tls *TLS, s uintptr, wc wchar_t, st uintptr) size_t { /* wcrtomb.c:6:8: */
	if !(s != 0) {
		return uint64(1)
	}
	if uint32(wc) < uint32(0x80) {
		*(*int8)(unsafe.Pointer(s)) = int8(wc)
		return uint64(1)
	} else if (func() int32 {
		if !(!(int32(*(*uintptr)(unsafe.Pointer(((*__pthread)(unsafe.Pointer(__pthread_self(tls))).locale /* &.cat */) + uintptr(0)*8))) != 0)) {
			return 4
		}
		return 1
	}()) == 1 {
		if !(((uint32(wc)) - uint32(0xdf80)) < uint32(0x80)) {
			(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 84
			return Uint64FromInt32(-1)
		}
		*(*int8)(unsafe.Pointer(s)) = int8(wc)
		return uint64(1)
	} else if uint32(wc) < uint32(0x800) {
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0xc0 | (int32(wc) >> 6)))
		*(*int8)(unsafe.Pointer(s)) = (int8(0x80 | (int32(wc) & 0x3f)))
		return uint64(2)
	} else if (uint32(wc) < uint32(0xd800)) || ((uint32(wc) - uint32(0xe000)) < uint32(0x2000)) {
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0xe0 | (int32(wc) >> 12)))
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0x80 | ((int32(wc) >> 6) & 0x3f)))
		*(*int8)(unsafe.Pointer(s)) = (int8(0x80 | (int32(wc) & 0x3f)))
		return uint64(3)
	} else if (uint32(wc) - uint32(0x10000)) < uint32(0x100000) {
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0xf0 | (int32(wc) >> 18)))
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0x80 | ((int32(wc) >> 12) & 0x3f)))
		*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = (int8(0x80 | ((int32(wc) >> 6) & 0x3f)))
		*(*int8)(unsafe.Pointer(s)) = (int8(0x80 | (int32(wc) & 0x3f)))
		return uint64(4)
	}
	(*(*int32)(unsafe.Pointer(X___errno_location(tls)))) = 84
	return Uint64FromInt32(-1)
}

func Xwcsrtombs(tls *TLS, s uintptr, ws uintptr, n size_t, st uintptr) size_t { /* wcsrtombs.c:3:8: */
	bp := tls.Alloc(4)
	defer tls.Free(4)

	var ws2 uintptr
	// var buf [4]int8 at bp, 4

	var N size_t = n
	var l size_t
	if !(s != 0) {
		n = uint64(0)
		ws2 = *(*uintptr)(unsafe.Pointer(ws))
		for ; *(*wchar_t)(unsafe.Pointer(ws2)) != 0; ws2 += 2 {
			if uint32(*(*wchar_t)(unsafe.Pointer(ws2))) >= 0x80 {
				l = Xwcrtomb(tls, bp /* &buf[0] */, *(*wchar_t)(unsafe.Pointer(ws2)), uintptr(0))
				if !((l + uint64(1)) != 0) {
					return Uint64FromInt32(-1)
				}
				n = n + (l)
			} else {
				n++
			}
		}
		return n
	}
	for n >= uint64(4) {
		if (uint32(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws))))) - 1) >= 0x7f {
			if !(int32(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws))))) != 0) {
				*(*int8)(unsafe.Pointer(s)) = int8(0)
				*(*uintptr)(unsafe.Pointer(ws)) = uintptr(0)
				return (N - n)
			}
			l = Xwcrtomb(tls, s, *(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws)))), uintptr(0))
			if !((l + uint64(1)) != 0) {
				return Uint64FromInt32(-1)
			}
			s += uintptr(l)
			n = n - (l)
		} else {
			*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = int8(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws)))))
			n--
		}
		(*(*uintptr)(unsafe.Pointer(ws))) += 2
	}
	for n != 0 {
		if (uint32(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws))))) - 1) >= 0x7f {
			if !(int32(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws))))) != 0) {
				*(*int8)(unsafe.Pointer(s)) = int8(0)
				*(*uintptr)(unsafe.Pointer(ws)) = uintptr(0)
				return (N - n)
			}
			l = Xwcrtomb(tls, bp /* &buf[0] */, *(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws)))), uintptr(0))
			if !((l + uint64(1)) != 0) {
				return Uint64FromInt32(-1)
			}
			if l > n {
				return (N - n)
			}
			Xwcrtomb(tls, s, *(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws)))), uintptr(0))
			s += uintptr(l)
			n = n - (l)
		} else {
			*(*int8)(unsafe.Pointer(PostIncUintptr(&s, 1))) = int8(*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ws)))))
			n--
		}
		(*(*uintptr)(unsafe.Pointer(ws))) += 2
	}
	return N
}

func Xwcstombs(tls *TLS, s uintptr, ws uintptr, n size_t) size_t { /* wcstombs.c:4:8: */
	bp := tls.Alloc(8)
	defer tls.Free(8)
	*(*uintptr)(unsafe.Pointer(bp)) = ws

	//TODO return wcsrtombs(s, &(const wchar_t *){ws}, n, 0);
	return Xwcsrtombs(tls, s, bp /* &ws */, n, uintptr(0))
}

// Support signed or unsigned plain-char

// Implementation choices...

// Arbitrary numbers...

// POSIX/SUS requirements follow. These numbers come directly
// from SUS and have nothing to do with the host system.

func X__strchrnul(tls *TLS, s uintptr, c int32) uintptr { /* strchrnul.c:10:6: */
	c = int32(uint8(c))
	if !(c != 0) {
		return (s + uintptr(Xstrlen(tls, s)))
	}
	var w uintptr
	for ; (uint64(s) % (uint64(unsafe.Sizeof(size_t(0))))) != 0; s++ {
		if !(int32(*(*int8)(unsafe.Pointer(s))) != 0) || (int32(*(*uint8)(unsafe.Pointer(s))) == c) {
			return s
		}
	}
	var k size_t = ((Uint64(Uint64FromInt32(-1)) / uint64(255)) * size_t(c))
	for w = s; !(((((*(*uint64)(unsafe.Pointer(w))) - (Uint64(Uint64FromInt32(-1)) / uint64(255))) & ^(*(*uint64)(unsafe.Pointer(w)))) & ((Uint64(Uint64FromInt32(-1)) / uint64(255)) * (uint64((255 / 2) + 1)))) != 0) && !(((((*(*uint64)(unsafe.Pointer(w)) ^ k) - (Uint64(Uint64FromInt32(-1)) / uint64(255))) & ^(*(*uint64)(unsafe.Pointer(w)) ^ k)) & ((Uint64(Uint64FromInt32(-1)) / uint64(255)) * (uint64((255 / 2) + 1)))) != 0); w += 8 {
	}
	s = w
	for ; (*(*int8)(unsafe.Pointer(s)) != 0) && (int32(*(*uint8)(unsafe.Pointer(s))) != c); s++ {
	}
	return s
}
