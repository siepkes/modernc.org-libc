// Code generated by 'ccgo stdio/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o stdio/stdio_linux_s390x.go -pkgname stdio', DO NOT EDIT.

package stdio

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

const (
	BUFSIZ                  = 8192
	EOF                     = -1
	FILENAME_MAX            = 4096
	FOPEN_MAX               = 16
	L_ctermid               = 9
	L_tmpnam                = 20
	P_tmpdir                = "/tmp"
	SEEK_CUR                = 1
	SEEK_END                = 2
	SEEK_SET                = 0
	TMP_MAX                 = 238328
	X_ATFILE_SOURCE         = 1
	X_BITS_STDIO_LIM_H      = 1
	X_BITS_TYPESIZES_H      = 1
	X_BITS_TYPES_H          = 1
	X_BSD_SIZE_T_           = 0
	X_BSD_SIZE_T_DEFINED_   = 0
	X_DEFAULT_SOURCE        = 1
	X_FEATURES_H            = 1
	X_FILE_OFFSET_BITS      = 64
	X_GCC_SIZE_T            = 0
	X_G_BUFSIZ              = 8192
	X_G_HAVE_MMAP           = 1
	X_G_HAVE_MREMAP         = 1
	X_G_HAVE_ST_BLKSIZE     = 0
	X_G_IO_IO_FILE_VERSION  = 0x20001
	X_G_config_h            = 1
	X_IOFBF                 = 0
	X_IOLBF                 = 1
	X_IONBF                 = 2
	X_IOS_APPEND            = 8
	X_IOS_ATEND             = 4
	X_IOS_BIN               = 128
	X_IOS_INPUT             = 1
	X_IOS_NOCREATE          = 32
	X_IOS_NOREPLACE         = 64
	X_IOS_OUTPUT            = 2
	X_IOS_TRUNC             = 16
	X_IO_BAD_SEEN           = 0x4000
	X_IO_BOOLALPHA          = 0200000
	X_IO_BUFSIZ             = 8192
	X_IO_CURRENTLY_PUTTING  = 0x800
	X_IO_DEC                = 020
	X_IO_DELETE_DONT_CLOSE  = 0x40
	X_IO_DONT_CLOSE         = 0100000
	X_IO_EOF_SEEN           = 0x10
	X_IO_ERR_SEEN           = 0x20
	X_IO_FIXED              = 010000
	X_IO_FLAGS2_MMAP        = 1
	X_IO_FLAGS2_NOTCANCEL   = 2
	X_IO_FLAGS2_USER_WBUF   = 8
	X_IO_HAVE_ST_BLKSIZE    = 0
	X_IO_HEX                = 0100
	X_IO_INTERNAL           = 010
	X_IO_IN_BACKUP          = 0x100
	X_IO_IS_APPENDING       = 0x1000
	X_IO_IS_FILEBUF         = 0x2000
	X_IO_LEFT               = 02
	X_IO_LINE_BUF           = 0x200
	X_IO_LINKED             = 0x80
	X_IO_MAGIC              = 0xFBAD0000
	X_IO_MAGIC_MASK         = 0xFFFF0000
	X_IO_NO_READS           = 4
	X_IO_NO_WRITES          = 8
	X_IO_OCT                = 040
	X_IO_RIGHT              = 04
	X_IO_SCIENTIFIC         = 04000
	X_IO_SHOWBASE           = 0200
	X_IO_SHOWPOINT          = 0400
	X_IO_SHOWPOS            = 02000
	X_IO_SKIPWS             = 01
	X_IO_STDIO              = 040000
	X_IO_STDIO_H            = 0
	X_IO_TIED_PUT_GET       = 0x400
	X_IO_UNBUFFERED         = 2
	X_IO_UNIFIED_JUMPTABLES = 1
	X_IO_UNITBUF            = 020000
	X_IO_UPPERCASE          = 01000
	X_IO_USER_BUF           = 1
	X_IO_USER_LOCK          = 0x8000
	X_LP64                  = 1
	X_OLD_STDIO_MAGIC       = 0xFABC0000
	X_POSIX_C_SOURCE        = 200809
	X_POSIX_SOURCE          = 1
	X_SIZET_                = 0
	X_SIZE_T                = 0
	X_SIZE_T_               = 0
	X_SIZE_T_DECLARED       = 0
	X_SIZE_T_DEFINED        = 0
	X_SIZE_T_DEFINED_       = 0
	X_STDC_PREDEF_H         = 1
	X_STDIO_H               = 1
	X_STDIO_USES_IOSTREAM   = 0
	X_SYS_CDEFS_H           = 1
	X_SYS_SIZE_T_H          = 0
	X_T_SIZE                = 0
	X_T_SIZE_               = 0
	X_VA_LIST_DEFINED       = 0
	Linux                   = 1
	Unix                    = 1
)

// This is the structure from the libstdc++ codecvt class.
const ( /* libio.h:176:1: */
	X__codecvt_ok      = 0
	X__codecvt_partial = 1
	X__codecvt_error   = 2
	X__codecvt_noconv  = 3
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

//  In 4.3bsd-net2, leave these undefined to indicate that size_t, etc.
//     are already defined.
//  BSD/OS 3.1 and FreeBSD [23].x require the MACHINE_ANSI_H check here.
//  NetBSD 5 requires the I386_ANSI_H and X86_64_ANSI_H checks here.

// A null pointer constant.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2017 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Copyright (C) 1991-2017 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Determine the wordsize from the preprocessor defines.

// Convenience types.
type X__u_char = uint8   /* types.h:30:23 */
type X__u_short = uint16 /* types.h:31:28 */
type X__u_int = uint32   /* types.h:32:22 */
type X__u_long = uint64  /* types.h:33:27 */

// Fixed-size types, underlying types depend on word size and compiler.
type X__int8_t = int8     /* types.h:36:21 */
type X__uint8_t = uint8   /* types.h:37:23 */
type X__int16_t = int16   /* types.h:38:26 */
type X__uint16_t = uint16 /* types.h:39:28 */
type X__int32_t = int32   /* types.h:40:20 */
type X__uint32_t = uint32 /* types.h:41:22 */
type X__int64_t = int64   /* types.h:43:25 */
type X__uint64_t = uint64 /* types.h:44:27 */

// quad_t is also 64 bits.
type X__quad_t = int64    /* types.h:52:18 */
type X__u_quad_t = uint64 /* types.h:53:27 */

// Largest integral types.
type X__intmax_t = int64   /* types.h:61:18 */
type X__uintmax_t = uint64 /* types.h:62:27 */

// The machine-dependent file <bits/typesizes.h> defines __*_T_TYPE
//    macros for each of the OS types we define below.  The definitions
//    of those macros must use the following macros for underlying types.
//    We define __S<SIZE>_TYPE and __U<SIZE>_TYPE for the signed and unsigned
//    variants of each of the following integer types on this machine.
//
// 	16		-- "natural" 16-bit type (always short)
// 	32		-- "natural" 32-bit type (always int)
// 	64		-- "natural" 64-bit type (long or long long)
// 	LONG32		-- 32-bit type, traditionally long
// 	QUAD		-- 64-bit type, always long long
// 	WORD		-- natural type of __WORDSIZE bits (int or long)
// 	LONGWORD	-- type of __WORDSIZE bits, traditionally long
//
//    We distinguish WORD/LONGWORD, 32/LONG32, and 64/QUAD so that the
//    conventional uses of `long' or `long long' type modifiers match the
//    types we define, even when a less-adorned type would be the same size.
//    This matters for (somewhat) portably writing printf/scanf formats for
//    these types, where using the appropriate l or ll format modifiers can
//    make the typedefs and the formats match up across all GNU platforms.  If
//    we used `long' when it's 64 bits where `long long' is expected, then the
//    compiler would warn about the formats not matching the argument types,
//    and the programmer changing them to shut up the compiler would break the
//    program's portability.
//
//    Here we assume what is presently the case in all the GCC configurations
//    we support: long long is always 64 bits, long is always word/address size,
//    and int is always 32 bits.

// No need to mark the typedef with __extension__.
// bits/typesizes.h -- underlying types for *_t.  Linux/s390 version.
//    Copyright (C) 2003-2017 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// See <bits/types.h> for the meaning of these macros.  This file exists so
//    that <bits/types.h> need not vary across different GNU platforms.

// size_t is unsigned long int on s390 -m31.

// Tell the libc code that off_t and off64_t are actually the same type
//    for all ABI purposes, even if possibly expressed as different base types
//    for C type-checking purposes.

// Same for ino_t and ino64_t.

// And for __rlim_t and __rlim64_t.

// Number of descriptors that can fit in an `fd_set'.

type X__dev_t = uint64                     /* types.h:133:25 */ // Type of device numbers.
type X__uid_t = uint32                     /* types.h:134:25 */ // Type of user identifications.
type X__gid_t = uint32                     /* types.h:135:25 */ // Type of group identifications.
type X__ino_t = uint64                     /* types.h:136:25 */ // Type of file serial numbers.
type X__ino64_t = uint64                   /* types.h:137:27 */ // Type of file serial numbers (LFS).
type X__mode_t = uint32                    /* types.h:138:26 */ // Type of file attribute bitmasks.
type X__nlink_t = uint64                   /* types.h:139:27 */ // Type of file link counts.
type X__off_t = int64                      /* types.h:140:25 */ // Type of file sizes and offsets.
type X__off64_t = int64                    /* types.h:141:27 */ // Type of file sizes and offsets (LFS).
type X__pid_t = int32                      /* types.h:142:25 */ // Type of process identifications.
type X__fsid_t = struct{ F__val [2]int32 } /* types.h:143:26 */ // Type of file system IDs.
type X__clock_t = int64                    /* types.h:144:27 */ // Type of CPU usage counts.
type X__rlim_t = uint64                    /* types.h:145:26 */ // Type for resource measurement.
type X__rlim64_t = uint64                  /* types.h:146:28 */ // Type for resource measurement (LFS).
type X__id_t = uint32                      /* types.h:147:24 */ // General type for IDs.
type X__time_t = int64                     /* types.h:148:26 */ // Seconds since the Epoch.
type X__useconds_t = uint32                /* types.h:149:30 */ // Count of microseconds.
type X__suseconds_t = int64                /* types.h:150:31 */ // Signed count of microseconds.

type X__daddr_t = int32 /* types.h:152:27 */ // The type of a disk address.
type X__key_t = int32   /* types.h:153:25 */ // Type of an IPC key.

// Clock ID used in clock and timer functions.
type X__clockid_t = int32 /* types.h:156:29 */

// Timer ID returned by `timer_create'.
type X__timer_t = uintptr /* types.h:159:12 */

// Type to represent block size.
type X__blksize_t = int64 /* types.h:162:29 */

// Types from the Large File Support interface.

// Type to count number of disk blocks.
type X__blkcnt_t = int64   /* types.h:167:28 */
type X__blkcnt64_t = int64 /* types.h:168:30 */

// Type to count file system blocks.
type X__fsblkcnt_t = uint64   /* types.h:171:30 */
type X__fsblkcnt64_t = uint64 /* types.h:172:32 */

// Type to count file system nodes.
type X__fsfilcnt_t = uint64   /* types.h:175:30 */
type X__fsfilcnt64_t = uint64 /* types.h:176:32 */

// Type of miscellaneous file system fields.
type X__fsword_t = int64 /* types.h:179:28 */

type X__ssize_t = int64 /* types.h:181:27 */ // Type of a byte count, or error.

// Signed long type used in system calls.
type X__syscall_slong_t = int64 /* types.h:184:33 */
// Unsigned long type used in system calls.
type X__syscall_ulong_t = uint64 /* types.h:186:33 */

// These few don't really vary by system, they always correspond
//    to one of the other defined types.
type X__loff_t = X__off64_t /* types.h:190:19 */ // Type of file sizes and offsets (LFS).
type X__qaddr_t = uintptr   /* types.h:191:18 */
type X__caddr_t = uintptr   /* types.h:192:14 */

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int64 /* types.h:195:25 */

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* types.h:198:23 */

// C99: An integer type that can be accessed as an atomic entity,
//    even in the presence of asynchronous interrupts.
//    It is not currently necessary for this to be machine-specific.
type X__sig_atomic_t = int32 /* types.h:203:13 */

type X_IO_FILE1 = struct {
	F_flags          int32
	_                [4]byte
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     X__off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]int8
	_                [4]byte
	F_lock           uintptr
	F_offset         X__off64_t
	F__pad1          uintptr
	F__pad2          uintptr
	F__pad3          uintptr
	F__pad4          uintptr
	F__pad5          Size_t
	F_mode           int32
	F_unused2        [20]int8
} /* __FILE.h:4:1 */

type X__FILE = X_IO_FILE1 /* __FILE.h:5:25 */

// The opaque type of streams.  This is the definition used elsewhere.
type FILE = X_IO_FILE1 /* FILE.h:7:25 */

// Copyright (C) 1991-2017 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//    Written by Per Bothner <bothner@cygnus.com>.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.
//
//    As a special exception, if you link the code in this file with
//    files compiled with a GNU compiler to produce an executable,
//    that does not cause the resulting executable to be covered by
//    the GNU Lesser General Public License.  This exception does not
//    however invalidate any other reasons why the executable file
//    might be covered by the GNU Lesser General Public License.
//    This exception applies to code released by its copyright holders
//    in files containing the exception.

// This file is needed by libio to define various configuration parameters.
//    These are always the same in the GNU C library.

// Define types for libio in terms of the standard internal type names.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2017 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Copyright (C) 1989-2017 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3, or (at your option)
// any later version.
//
// GCC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// ISO C Standard:  7.17  Common definitions  <stddef.h>

// Any one of these symbols __need_* means that GNU libc
//    wants us just to define one data type.  So don't define
//    the symbols that indicate this file's entire job has been done.

// This avoids lossage on SunOS but only if stdtypes.h comes first.
//    There's no way to win with the other order!  Sun lossage.

// On 4.3bsd-net2, make sure ansi.h is included, so we have
//    one less case to deal with in the following.
// On FreeBSD 5, machine/ansi.h does not exist anymore...

// In 4.3bsd-net2, machine/ansi.h defines these symbols, which are
//    defined if the corresponding type is *not* defined.
//    FreeBSD-2.1 defines _MACHINE_ANSI_H_ instead of _ANSI_H_.
//    NetBSD defines _I386_ANSI_H_ and _X86_64_ANSI_H_ instead of _ANSI_H_

// Sequent's header files use _PTRDIFF_T_ in some conflicting way.
//    Just ignore it.

// On VxWorks, <type/vxTypesBase.h> may have defined macros like
//    _TYPE_size_t which will typedef size_t.  fixincludes patched the
//    vxTypesBase.h so that this macro is only defined if _GCC_SIZE_T is
//    not defined, and so that defining this macro defines _GCC_SIZE_T.
//    If we find that the macros are still defined at this point, we must
//    invoke them so that the type is defined as expected.

// In case nobody has defined these types, but we aren't running under
//    GCC 2.00, make sure that __PTRDIFF_TYPE__, __SIZE_TYPE__, and
//    __WCHAR_TYPE__ have reasonable values.  This can happen if the
//    parts of GCC is compiled by an older compiler, that actually
//    include gstddef.h, such as collect2.

// Signed type of difference of two pointers.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// Unsigned type of `sizeof' something.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

//  In 4.3bsd-net2, leave these undefined to indicate that size_t, etc.
//     are already defined.
//  BSD/OS 3.1 and FreeBSD [23].x require the MACHINE_ANSI_H check here.
//  NetBSD 5 requires the I386_ANSI_H and X86_64_ANSI_H checks here.

// A null pointer constant.

// Integral type unchanged by default argument promotions that can
//    hold any value corresponding to members of the extended character
//    set, as well as at least one value that does not correspond to any
//    member of the extended character set.

// Conversion state information.
type X__mbstate_t = struct {
	F__count int32
	F__value struct{ F__wch uint32 }
} /* __mbstate_t.h:21:3 */

type X_G_fpos_t = struct {
	F__pos   X__off_t
	F__state X__mbstate_t
} /* _G_config.h:26:3 */
type X_G_fpos64_t = struct {
	F__pos   X__off64_t
	F__state X__mbstate_t
} /* _G_config.h:31:3 */

// These library features are always available in the GNU C library.

// This is defined by <bits/stat.h> if `st_blksize' exists.

// ALL of these should be defined in _G_config.h

// This define avoids name pollution if we're using GNU stdarg.h
// Copyright (C) 1989-2017 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3, or (at your option)
// any later version.
//
// GCC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// ISO C Standard:  7.15  Variable arguments  <stdarg.h>

// Define __gnuc_va_list.

type X__gnuc_va_list = X__builtin_va_list /* stdarg.h:40:27 */

// A streammarker remembers a position in a buffer.

type X_IO_marker = struct {
	F_next uintptr
	F_sbuf uintptr
	F_pos  int32
	_      [4]byte
} /* __FILE.h:4:1 */

type X_IO_FILE = X_IO_FILE1 /* libio.h:310:25 */

type Va_list = X__gnuc_va_list /* stdio.h:46:20 */

type Off_t = X__off64_t /* stdio.h:59:19 */

type Ssize_t = X__ssize_t /* stdio.h:71:19 */

// The type of the second argument to `fgetpos' and `fsetpos'.
type Fpos_t = X_G_fpos64_t /* stdio.h:80:21 */

// If we are compiling with optimizing read this file.  It contains
//    several optimizing inline functions and macros.

var _ int8 /* gen.c:2:13: */
