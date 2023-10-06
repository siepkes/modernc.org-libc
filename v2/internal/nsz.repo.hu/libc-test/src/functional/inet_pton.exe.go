// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/functional/inet_pton.exe.go src/functional/inet_pton.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:98:9:
const EAFNOSUPPORT = 97

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:28:9:
const ENOSPC = 28

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:103:9:
const PF_INET = 2

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:111:9:
const PF_INET6 = 10

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:73:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:73:24:
type size_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:32:
type locale_t = uintptr

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:12:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:12:24:
type wchar_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/inttypes.h:14:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/inttypes.h:14:40:
type imaxdiv_t = struct {
	Fquot intmax_t
	Frem  intmax_t
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:13:
type pid_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:268:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:268:18:
type uid_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:273:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:273:18:
type gid_t = uint32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:378:1:
type iovec = struct {
	Fiov_base uintptr
	Fiov_len  size_t
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:390:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:390:18:
type socklen_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:395:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:395:24:
type sa_family_t = uint16

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:22:1:
type msghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    socklen_t
	Fmsg_iov        uintptr
	Fmsg_iovlen     int32
	Fmsg_control    uintptr
	Fmsg_controllen socklen_t
	Fmsg_flags      int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:44:1:
type cmsghdr = struct {
	Fcmsg_len   socklen_t
	Fcmsg_level int32
	Fcmsg_type  int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:57:1:
type ucred = struct {
	Fpid pid_t
	Fuid uid_t
	Fgid gid_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:63:1:
type mmsghdr = struct {
	Fmsg_hdr msghdr
	Fmsg_len uint32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:74:1:
type linger = struct {
	Fl_onoff  int32
	Fl_linger int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:369:1:
type sockaddr = struct {
	Fsa_family sa_family_t
	Fsa_data   [14]int8
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sys/socket.h:374:1:
type sockaddr_storage = struct {
	Fss_family    sa_family_t
	F__ss_padding [122]int8
	F__ss_align   uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:12:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:12:18:
type in_port_t = uint16

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:13:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:13:18:
type in_addr_t = uint32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:14:1:
type in_addr = struct {
	Fs_addr in_addr_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:16:1:
type sockaddr_in = struct {
	Fsin_family sa_family_t
	Fsin_port   in_port_t
	Fsin_addr   in_addr
	Fsin_zero   [8]uint8_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:23:1:
type in6_addr = struct {
	F__in6_union struct {
		F__s6_addr16 [0][8]uint16_t
		F__s6_addr32 [0][4]uint32_t
		F__s6_addr   [16]uint8_t
	}
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:34:1:
type sockaddr_in6 = struct {
	Fsin6_family   sa_family_t
	Fsin6_port     in_port_t
	Fsin6_flowinfo uint32_t
	Fsin6_addr     in6_addr
	Fsin6_scope_id uint32_t
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:42:1:
type ipv6_mreq = struct {
	Fipv6mr_multiaddr in6_addr
	Fipv6mr_interface uint32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:233:1:
type ip_opts = struct {
	Fip_dst  in_addr
	Fip_opts [40]int8
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:251:1:
type ip_mreq = struct {
	Fimr_multiaddr in_addr
	Fimr_interface in_addr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:256:1:
type ip_mreqn = struct {
	Fimr_multiaddr in_addr
	Fimr_address   in_addr
	Fimr_ifindex   int32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:262:1:
type ip_mreq_source = struct {
	Fimr_multiaddr  in_addr
	Fimr_interface  in_addr
	Fimr_sourceaddr in_addr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:268:1:
type ip_msfilter = struct {
	Fimsf_multiaddr in_addr
	Fimsf_interface in_addr
	Fimsf_fmode     uint32_t
	Fimsf_numsrc    uint32_t
	Fimsf_slist     [1]in_addr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:279:1:
type group_req = struct {
	Fgr_interface uint32_t
	Fgr_group     sockaddr_storage
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:284:1:
type group_source_req = struct {
	Fgsr_interface uint32_t
	Fgsr_group     sockaddr_storage
	Fgsr_source    sockaddr_storage
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:290:1:
type group_filter = struct {
	Fgf_interface uint32_t
	Fgf_group     sockaddr_storage
	Fgf_fmode     uint32_t
	Fgf_numsrc    uint32_t
	Fgf_slist     [1]sockaddr_storage
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:301:1:
type in_pktinfo = struct {
	Fipi_ifindex  int32
	Fipi_spec_dst in_addr
	Fipi_addr     in_addr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:307:1:
type in6_pktinfo = struct {
	Fipi6_addr    in6_addr
	Fipi6_ifindex uint32
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/netinet/in.h:312:1:
type ip6_mtuinfo = struct {
	Fip6m_addr sockaddr_in6
	Fip6m_mtu  uint32_t
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:18:
type useconds_t = uint32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:8:12:
func digit(tls *libc.TLS, c int32) (r int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:9:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:10:2:
	c -= int32('0')
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:11:2:
	if c > int32(9) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:11:11:
		c -= libc.Int32FromUint8('a') - libc.Int32FromUint8('0') - libc.Int32FromInt32(10)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:12:2:
	return c
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:15:13:
func tobin(tls *libc.TLS, d uintptr, s uintptr) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:16:1:
	var i int32
	var p uintptr
	_, _ = i, p
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:18:16:
	p = d
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:19:2:
	i = 0
	for {
		if !(*(*int8)(unsafe.Pointer(s + uintptr(int32(2)*i))) != 0) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:19:25:
		*(*uint8)(unsafe.Pointer(p + uintptr(i))) = uint8(digit(tls, int32(*(*int8)(unsafe.Pointer(s + uintptr(int32(2)*i)))))*int32(16) + digit(tls, int32(*(*int8)(unsafe.Pointer(s + uintptr(int32(2)*i+int32(1)))))))
		goto _1
	_1:
		i++
	}
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:22:13:
func tohex(tls *libc.TLS, d uintptr, s uintptr, n int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:23:1:
	var i int32
	var p uintptr
	_, _ = i, p
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:25:16:
	p = s
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:26:2:
	i = 0
	for {
		if !(i < n) {
			break
		}
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:26:22:
		libc.Xsprintf(tls, d+uintptr(int32(2)*i), __ccgo_ts, libc.VaList(bp+8, int32(*(*uint8)(unsafe.Pointer(p + uintptr(i))))))
		goto _1
	_1:
		i++
	}
}

// ret and hex are the results of inet_pton and inet_addr respectively

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:90:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r50 int32) {
	bp := tls.Alloc(4096)
	defer tls.Free(4096)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:91:1:
	var in, in1, in10, in11, in12, in13, in14, in15, in16, in17, in18, in2, in3, in4, in5, in6, in7, in8, in9 in_addr
	var p, p1, p10, p11, p12, p13, p14, p15, p16, p17, p18, p2, p3, p4, p5, p6, p7, p8, p9 uintptr
	var r, r1, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r2, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r3, r30, r31, r32, r33, r34, r35, r36, r37, r38, r39, r4, r40, r41, r42, r43, r44, r45, r46, r47, r48, r49, r5, r6, r7, r8, r9 int32
	var _ /* a at bp+0 */ uint32_t
	var _ /* a at bp+120 */ uint32_t
	var _ /* a at bp+144 */ uint32_t
	var _ /* a at bp+168 */ uint32_t
	var _ /* a at bp+192 */ uint32_t
	var _ /* a at bp+216 */ uint32_t
	var _ /* a at bp+24 */ uint32_t
	var _ /* a at bp+240 */ uint32_t
	var _ /* a at bp+264 */ uint32_t
	var _ /* a at bp+288 */ uint32_t
	var _ /* a at bp+312 */ uint32_t
	var _ /* a at bp+336 */ uint32_t
	var _ /* a at bp+360 */ uint32_t
	var _ /* a at bp+384 */ uint32_t
	var _ /* a at bp+408 */ uint32_t
	var _ /* a at bp+432 */ uint32_t
	var _ /* a at bp+48 */ uint32_t
	var _ /* a at bp+72 */ uint32_t
	var _ /* a at bp+96 */ uint32_t
	var _ /* binaddr at bp+1036 */ [16]int8
	var _ /* binaddr at bp+1152 */ [16]int8
	var _ /* binaddr at bp+1268 */ [16]int8
	var _ /* binaddr at bp+1384 */ [16]int8
	var _ /* binaddr at bp+1500 */ [16]int8
	var _ /* binaddr at bp+1616 */ [16]int8
	var _ /* binaddr at bp+1732 */ [16]int8
	var _ /* binaddr at bp+1848 */ [16]int8
	var _ /* binaddr at bp+1964 */ [16]int8
	var _ /* binaddr at bp+2080 */ [16]int8
	var _ /* binaddr at bp+2196 */ [16]int8
	var _ /* binaddr at bp+2312 */ [16]int8
	var _ /* binaddr at bp+2428 */ [16]int8
	var _ /* binaddr at bp+2544 */ [16]int8
	var _ /* binaddr at bp+2660 */ [16]int8
	var _ /* binaddr at bp+2776 */ [16]int8
	var _ /* binaddr at bp+2892 */ [16]int8
	var _ /* binaddr at bp+3008 */ [16]int8
	var _ /* binaddr at bp+3124 */ [16]int8
	var _ /* binaddr at bp+3240 */ [16]int8
	var _ /* binaddr at bp+3356 */ [16]int8
	var _ /* binaddr at bp+3472 */ [16]int8
	var _ /* binaddr at bp+3588 */ [16]int8
	var _ /* binaddr at bp+3704 */ [16]int8
	var _ /* binaddr at bp+3820 */ [16]int8
	var _ /* binaddr at bp+3936 */ [16]int8
	var _ /* binaddr at bp+456 */ [16]int8
	var _ /* binaddr at bp+572 */ [16]int8
	var _ /* binaddr at bp+688 */ [16]int8
	var _ /* binaddr at bp+804 */ [16]int8
	var _ /* binaddr at bp+920 */ [16]int8
	var _ /* buf at bp+100 */ [20]int8
	var _ /* buf at bp+124 */ [20]int8
	var _ /* buf at bp+148 */ [20]int8
	var _ /* buf at bp+172 */ [20]int8
	var _ /* buf at bp+196 */ [20]int8
	var _ /* buf at bp+220 */ [20]int8
	var _ /* buf at bp+244 */ [20]int8
	var _ /* buf at bp+268 */ [20]int8
	var _ /* buf at bp+28 */ [20]int8
	var _ /* buf at bp+292 */ [20]int8
	var _ /* buf at bp+316 */ [20]int8
	var _ /* buf at bp+340 */ [20]int8
	var _ /* buf at bp+364 */ [20]int8
	var _ /* buf at bp+388 */ [20]int8
	var _ /* buf at bp+4 */ [20]int8
	var _ /* buf at bp+412 */ [20]int8
	var _ /* buf at bp+436 */ [20]int8
	var _ /* buf at bp+52 */ [20]int8
	var _ /* buf at bp+76 */ [20]int8
	var _ /* hexaddr at bp+1052 */ [40]int8
	var _ /* hexaddr at bp+1168 */ [40]int8
	var _ /* hexaddr at bp+1284 */ [40]int8
	var _ /* hexaddr at bp+1400 */ [40]int8
	var _ /* hexaddr at bp+1516 */ [40]int8
	var _ /* hexaddr at bp+1632 */ [40]int8
	var _ /* hexaddr at bp+1748 */ [40]int8
	var _ /* hexaddr at bp+1864 */ [40]int8
	var _ /* hexaddr at bp+1980 */ [40]int8
	var _ /* hexaddr at bp+2096 */ [40]int8
	var _ /* hexaddr at bp+2212 */ [40]int8
	var _ /* hexaddr at bp+2328 */ [40]int8
	var _ /* hexaddr at bp+2444 */ [40]int8
	var _ /* hexaddr at bp+2560 */ [40]int8
	var _ /* hexaddr at bp+2676 */ [40]int8
	var _ /* hexaddr at bp+2792 */ [40]int8
	var _ /* hexaddr at bp+2908 */ [40]int8
	var _ /* hexaddr at bp+3024 */ [40]int8
	var _ /* hexaddr at bp+3140 */ [40]int8
	var _ /* hexaddr at bp+3256 */ [40]int8
	var _ /* hexaddr at bp+3372 */ [40]int8
	var _ /* hexaddr at bp+3488 */ [40]int8
	var _ /* hexaddr at bp+3604 */ [40]int8
	var _ /* hexaddr at bp+3720 */ [40]int8
	var _ /* hexaddr at bp+3836 */ [40]int8
	var _ /* hexaddr at bp+3952 */ [40]int8
	var _ /* hexaddr at bp+472 */ [40]int8
	var _ /* hexaddr at bp+588 */ [40]int8
	var _ /* hexaddr at bp+704 */ [40]int8
	var _ /* hexaddr at bp+820 */ [40]int8
	var _ /* hexaddr at bp+936 */ [40]int8
	var _ /* txtaddr at bp+1092 */ [60]int8
	var _ /* txtaddr at bp+1208 */ [60]int8
	var _ /* txtaddr at bp+1324 */ [60]int8
	var _ /* txtaddr at bp+1440 */ [60]int8
	var _ /* txtaddr at bp+1556 */ [60]int8
	var _ /* txtaddr at bp+1672 */ [60]int8
	var _ /* txtaddr at bp+1788 */ [60]int8
	var _ /* txtaddr at bp+1904 */ [60]int8
	var _ /* txtaddr at bp+2020 */ [60]int8
	var _ /* txtaddr at bp+2136 */ [60]int8
	var _ /* txtaddr at bp+2252 */ [60]int8
	var _ /* txtaddr at bp+2368 */ [60]int8
	var _ /* txtaddr at bp+2484 */ [60]int8
	var _ /* txtaddr at bp+2600 */ [60]int8
	var _ /* txtaddr at bp+2716 */ [60]int8
	var _ /* txtaddr at bp+2832 */ [60]int8
	var _ /* txtaddr at bp+2948 */ [60]int8
	var _ /* txtaddr at bp+3064 */ [60]int8
	var _ /* txtaddr at bp+3180 */ [60]int8
	var _ /* txtaddr at bp+3296 */ [60]int8
	var _ /* txtaddr at bp+3412 */ [60]int8
	var _ /* txtaddr at bp+3528 */ [60]int8
	var _ /* txtaddr at bp+3644 */ [60]int8
	var _ /* txtaddr at bp+3760 */ [60]int8
	var _ /* txtaddr at bp+3876 */ [60]int8
	var _ /* txtaddr at bp+3992 */ [60]int8
	var _ /* txtaddr at bp+512 */ [60]int8
	var _ /* txtaddr at bp+628 */ [60]int8
	var _ /* txtaddr at bp+744 */ [60]int8
	var _ /* txtaddr at bp+860 */ [60]int8
	var _ /* txtaddr at bp+976 */ [60]int8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = in, in1, in10, in11, in12, in13, in14, in15, in16, in17, in18, in2, in3, in4, in5, in6, in7, in8, in9, p, p1, p10, p11, p12, p13, p14, p15, p16, p17, p18, p2, p3, p4, p5, p6, p7, p8, p9, r, r1, r10, r11, r12, r13, r14, r15, r16, r17, r18, r19, r2, r20, r21, r22, r23, r24, r25, r26, r27, r28, r29, r3, r30, r31, r32, r33, r34, r35, r36, r37, r38, r39, r4, r40, r41, r42, r43, r44, r45, r46, r47, r48, r49, r5, r6, r7, r8, r9
	// errors
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:94:1:
	if libc.Xinet_pton(tls, int32(12345), __ccgo_ts+5, uintptr(0)) != -int32(1) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(EAFNOSUPPORT) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:95:2:
		t_printf(tls, __ccgo_ts+6, libc.VaList(bp+4064, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:96:1:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:97:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), __ccgo_ts+95, __ccgo_ts+5, uint32(0)) != uintptr(0) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(ENOSPC) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:98:2:
		t_printf(tls, __ccgo_ts+100, libc.VaList(bp+4064, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:99:1:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	// dotted-decimal notation
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	*(*[20]int8)(unsafe.Pointer(bp + 4)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	*(*uint32_t)(unsafe.Pointer(bp)) = libc.Xinet_addr(tls, __ccgo_ts+181)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	tohex(tls, bp+4, bp, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if libc.Xstrcmp(tls, bp+4, __ccgo_ts+189) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+198, libc.VaList(bp+4064, bp+4, __ccgo_ts+189))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	r = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+181, bp)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if r != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+273, libc.VaList(bp+4064, r, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		goto _1
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	tohex(tls, bp+4, bp, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if libc.Xstrcmp(tls, bp+4, __ccgo_ts+189) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+363, libc.VaList(bp+4064, bp+4, __ccgo_ts+189))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	tobin(tls, bp, __ccgo_ts+189)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp, bp+4, uint32(20)) != bp+4 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+453, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if libc.Xstrcmp(tls, bp+4, __ccgo_ts+181) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+549, libc.VaList(bp+4064, bp+4, __ccgo_ts+181))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	in.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	p = libc.Xinet_ntoa(tls, in)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
	if libc.Xstrcmp(tls, p, __ccgo_ts+181) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:102:1:
		t_printf(tls, __ccgo_ts+642, libc.VaList(bp+4064, p, __ccgo_ts+181))
	}
	goto _2
_2:
	goto _1
_1: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	*(*[20]int8)(unsafe.Pointer(bp + 28)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	*(*uint32_t)(unsafe.Pointer(bp + 24)) = libc.Xinet_addr(tls, __ccgo_ts+720)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	tohex(tls, bp+28, bp+24, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if libc.Xstrcmp(tls, bp+28, __ccgo_ts+730) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+739, libc.VaList(bp+4064, bp+28, __ccgo_ts+730))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	r1 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+720, bp+24)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if r1 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+816, libc.VaList(bp+4064, r1, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		goto _3
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	tohex(tls, bp+28, bp+24, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if libc.Xstrcmp(tls, bp+28, __ccgo_ts+730) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+908, libc.VaList(bp+4064, bp+28, __ccgo_ts+730))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	tobin(tls, bp+24, __ccgo_ts+730)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+24, bp+28, uint32(20)) != bp+28 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+1000, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if libc.Xstrcmp(tls, bp+28, __ccgo_ts+720) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+1096, libc.VaList(bp+4064, bp+28, __ccgo_ts+720))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	in1.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 24))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	p1 = libc.Xinet_ntoa(tls, in1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
	if libc.Xstrcmp(tls, p1, __ccgo_ts+720) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:103:1:
		t_printf(tls, __ccgo_ts+1189, libc.VaList(bp+4064, p1, __ccgo_ts+720))
	}
	goto _4
_4:
	goto _3
_3: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	*(*[20]int8)(unsafe.Pointer(bp + 52)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	*(*uint32_t)(unsafe.Pointer(bp + 48)) = libc.Xinet_addr(tls, __ccgo_ts+1267)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	tohex(tls, bp+52, bp+48, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if libc.Xstrcmp(tls, bp+52, __ccgo_ts+1279) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1288, libc.VaList(bp+4064, bp+52, __ccgo_ts+1279))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	r2 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+1267, bp+48)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if r2 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1367, libc.VaList(bp+4064, r2, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		goto _5
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	tohex(tls, bp+52, bp+48, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if libc.Xstrcmp(tls, bp+52, __ccgo_ts+1279) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1461, libc.VaList(bp+4064, bp+52, __ccgo_ts+1279))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	tobin(tls, bp+48, __ccgo_ts+1279)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+48, bp+52, uint32(20)) != bp+52 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1555, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if libc.Xstrcmp(tls, bp+52, __ccgo_ts+1267) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1651, libc.VaList(bp+4064, bp+52, __ccgo_ts+1267))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	in2.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 48))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	p2 = libc.Xinet_ntoa(tls, in2)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
	if libc.Xstrcmp(tls, p2, __ccgo_ts+1267) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:104:1:
		t_printf(tls, __ccgo_ts+1744, libc.VaList(bp+4064, p2, __ccgo_ts+1267))
	}
	goto _6
_6:
	goto _5
_5: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	*(*[20]int8)(unsafe.Pointer(bp + 76)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	*(*uint32_t)(unsafe.Pointer(bp + 72)) = libc.Xinet_addr(tls, __ccgo_ts+1822)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	tohex(tls, bp+76, bp+72, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if libc.Xstrcmp(tls, bp+76, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+1847, libc.VaList(bp+4064, bp+76, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	r3 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+1822, bp+72)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if r3 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+1930, libc.VaList(bp+4064, r3, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		goto _7
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	tohex(tls, bp+76, bp+72, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if libc.Xstrcmp(tls, bp+76, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+2028, libc.VaList(bp+4064, bp+76, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	tobin(tls, bp+72, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+72, bp+76, uint32(20)) != bp+76 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+2126, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if libc.Xstrcmp(tls, bp+76, __ccgo_ts+1822) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+2222, libc.VaList(bp+4064, bp+76, __ccgo_ts+1822))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	in3.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 72))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	p3 = libc.Xinet_ntoa(tls, in3)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
	if libc.Xstrcmp(tls, p3, __ccgo_ts+1822) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:105:1:
		t_printf(tls, __ccgo_ts+2315, libc.VaList(bp+4064, p3, __ccgo_ts+1822))
	}
	goto _8
_8:
	goto _7
_7: /**/
	; //
	// numbers-and-dots notation, but not dotted-decimal
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	*(*[20]int8)(unsafe.Pointer(bp + 100)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	*(*uint32_t)(unsafe.Pointer(bp + 96)) = libc.Xinet_addr(tls, __ccgo_ts+2393)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	tohex(tls, bp+100, bp+96, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if libc.Xstrcmp(tls, bp+100, __ccgo_ts+2402) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2411, libc.VaList(bp+4064, bp+100, __ccgo_ts+2402))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	r4 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+2393, bp+96)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if r4 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2487, libc.VaList(bp+4064, r4, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		goto _9
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	tohex(tls, bp+100, bp+96, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if libc.Xstrcmp(tls, bp+100, __ccgo_ts+2402) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2578, libc.VaList(bp+4064, bp+100, __ccgo_ts+2402))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	tobin(tls, bp+96, __ccgo_ts+2402)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+96, bp+100, uint32(20)) != bp+100 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2669, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if libc.Xstrcmp(tls, bp+100, __ccgo_ts+2393) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2765, libc.VaList(bp+4064, bp+100, __ccgo_ts+2393))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	in4.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 96))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	p4 = libc.Xinet_ntoa(tls, in4)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
	if libc.Xstrcmp(tls, p4, __ccgo_ts+2393) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:108:1:
		t_printf(tls, __ccgo_ts+2858, libc.VaList(bp+4064, p4, __ccgo_ts+2393))
	}
	goto _10
_10:
	goto _9
_9: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	*(*[20]int8)(unsafe.Pointer(bp + 124)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	*(*uint32_t)(unsafe.Pointer(bp + 120)) = libc.Xinet_addr(tls, __ccgo_ts+2936)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	tohex(tls, bp+124, bp+120, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if libc.Xstrcmp(tls, bp+124, __ccgo_ts+2947) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+2956, libc.VaList(bp+4064, bp+124, __ccgo_ts+2947))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	r5 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+2936, bp+120)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if r5 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+3034, libc.VaList(bp+4064, r5, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		goto _11
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	tohex(tls, bp+124, bp+120, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if libc.Xstrcmp(tls, bp+124, __ccgo_ts+2947) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+3127, libc.VaList(bp+4064, bp+124, __ccgo_ts+2947))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	tobin(tls, bp+120, __ccgo_ts+2947)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+120, bp+124, uint32(20)) != bp+124 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+3220, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if libc.Xstrcmp(tls, bp+124, __ccgo_ts+2936) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+3316, libc.VaList(bp+4064, bp+124, __ccgo_ts+2936))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	in5.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 120))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	p5 = libc.Xinet_ntoa(tls, in5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
	if libc.Xstrcmp(tls, p5, __ccgo_ts+2936) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:109:1:
		t_printf(tls, __ccgo_ts+3409, libc.VaList(bp+4064, p5, __ccgo_ts+2936))
	}
	goto _12
_12:
	goto _11
_11: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	*(*[20]int8)(unsafe.Pointer(bp + 148)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	*(*uint32_t)(unsafe.Pointer(bp + 144)) = libc.Xinet_addr(tls, __ccgo_ts+3487)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	tohex(tls, bp+148, bp+144, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if libc.Xstrcmp(tls, bp+148, __ccgo_ts+3498) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3507, libc.VaList(bp+4064, bp+148, __ccgo_ts+3498))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	r6 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+3487, bp+144)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if r6 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3585, libc.VaList(bp+4064, r6, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		goto _13
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	tohex(tls, bp+148, bp+144, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if libc.Xstrcmp(tls, bp+148, __ccgo_ts+3498) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3678, libc.VaList(bp+4064, bp+148, __ccgo_ts+3498))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	tobin(tls, bp+144, __ccgo_ts+3498)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+144, bp+148, uint32(20)) != bp+148 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3771, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if libc.Xstrcmp(tls, bp+148, __ccgo_ts+3487) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3867, libc.VaList(bp+4064, bp+148, __ccgo_ts+3487))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	in6.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 144))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	p6 = libc.Xinet_ntoa(tls, in6)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
	if libc.Xstrcmp(tls, p6, __ccgo_ts+3487) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:110:1:
		t_printf(tls, __ccgo_ts+3960, libc.VaList(bp+4064, p6, __ccgo_ts+3487))
	}
	goto _14
_14:
	goto _13
_13: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	*(*[20]int8)(unsafe.Pointer(bp + 172)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	*(*uint32_t)(unsafe.Pointer(bp + 168)) = libc.Xinet_addr(tls, __ccgo_ts+4038)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	tohex(tls, bp+172, bp+168, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if libc.Xstrcmp(tls, bp+172, __ccgo_ts+4049) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4058, libc.VaList(bp+4064, bp+172, __ccgo_ts+4049))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	r7 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+4038, bp+168)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if r7 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4136, libc.VaList(bp+4064, r7, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		goto _15
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	tohex(tls, bp+172, bp+168, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if libc.Xstrcmp(tls, bp+172, __ccgo_ts+4049) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4229, libc.VaList(bp+4064, bp+172, __ccgo_ts+4049))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	tobin(tls, bp+168, __ccgo_ts+4049)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+168, bp+172, uint32(20)) != bp+172 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4322, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if libc.Xstrcmp(tls, bp+172, __ccgo_ts+4038) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4418, libc.VaList(bp+4064, bp+172, __ccgo_ts+4038))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	in7.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 168))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	p7 = libc.Xinet_ntoa(tls, in7)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
	if libc.Xstrcmp(tls, p7, __ccgo_ts+4038) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:111:1:
		t_printf(tls, __ccgo_ts+4511, libc.VaList(bp+4064, p7, __ccgo_ts+4038))
	}
	goto _16
_16:
	goto _15
_15: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	*(*[20]int8)(unsafe.Pointer(bp + 196)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	*(*uint32_t)(unsafe.Pointer(bp + 192)) = libc.Xinet_addr(tls, __ccgo_ts+4589)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	tohex(tls, bp+196, bp+192, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if libc.Xstrcmp(tls, bp+196, __ccgo_ts+4600) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+4609, libc.VaList(bp+4064, bp+196, __ccgo_ts+4600))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	r8 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+4589, bp+192)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if r8 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+4687, libc.VaList(bp+4064, r8, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		goto _17
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	tohex(tls, bp+196, bp+192, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if libc.Xstrcmp(tls, bp+196, __ccgo_ts+4600) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+4780, libc.VaList(bp+4064, bp+196, __ccgo_ts+4600))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	tobin(tls, bp+192, __ccgo_ts+4600)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+192, bp+196, uint32(20)) != bp+196 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+4873, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if libc.Xstrcmp(tls, bp+196, __ccgo_ts+4589) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+4969, libc.VaList(bp+4064, bp+196, __ccgo_ts+4589))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	in8.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 192))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	p8 = libc.Xinet_ntoa(tls, in8)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
	if libc.Xstrcmp(tls, p8, __ccgo_ts+4589) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:112:1:
		t_printf(tls, __ccgo_ts+5062, libc.VaList(bp+4064, p8, __ccgo_ts+4589))
	}
	goto _18
_18:
	goto _17
_17: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	*(*[20]int8)(unsafe.Pointer(bp + 220)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	*(*uint32_t)(unsafe.Pointer(bp + 216)) = libc.Xinet_addr(tls, __ccgo_ts+5140)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	tohex(tls, bp+220, bp+216, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if libc.Xstrcmp(tls, bp+220, __ccgo_ts+5158) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5167, libc.VaList(bp+4064, bp+220, __ccgo_ts+5158))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	r9 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+5140, bp+216)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if r9 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5252, libc.VaList(bp+4064, r9, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		goto _19
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	tohex(tls, bp+220, bp+216, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if libc.Xstrcmp(tls, bp+220, __ccgo_ts+5158) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5352, libc.VaList(bp+4064, bp+220, __ccgo_ts+5158))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	tobin(tls, bp+216, __ccgo_ts+5158)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+216, bp+220, uint32(20)) != bp+220 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5452, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if libc.Xstrcmp(tls, bp+220, __ccgo_ts+5140) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5548, libc.VaList(bp+4064, bp+220, __ccgo_ts+5140))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	in9.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 216))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	p9 = libc.Xinet_ntoa(tls, in9)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
	if libc.Xstrcmp(tls, p9, __ccgo_ts+5140) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:113:1:
		t_printf(tls, __ccgo_ts+5641, libc.VaList(bp+4064, p9, __ccgo_ts+5140))
	}
	goto _20
_20:
	goto _19
_19: /**/
	; //
	// invalid
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	*(*[20]int8)(unsafe.Pointer(bp + 244)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	*(*uint32_t)(unsafe.Pointer(bp + 240)) = libc.Xinet_addr(tls, __ccgo_ts+5719)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	tohex(tls, bp+244, bp+240, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if libc.Xstrcmp(tls, bp+244, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+5726, libc.VaList(bp+4064, bp+244, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	r10 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+5719, bp+240)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if r10 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+5800, libc.VaList(bp+4064, r10, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		goto _21
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	tohex(tls, bp+244, bp+240, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if libc.Xstrcmp(tls, bp+244, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+5889, libc.VaList(bp+4064, bp+244, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	tobin(tls, bp+240, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+240, bp+244, uint32(20)) != bp+244 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+5978, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if libc.Xstrcmp(tls, bp+244, __ccgo_ts+5719) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+6074, libc.VaList(bp+4064, bp+244, __ccgo_ts+5719))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	in10.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 240))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	p10 = libc.Xinet_ntoa(tls, in10)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
	if libc.Xstrcmp(tls, p10, __ccgo_ts+5719) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:116:1:
		t_printf(tls, __ccgo_ts+6167, libc.VaList(bp+4064, p10, __ccgo_ts+5719))
	}
	goto _22
_22:
	goto _21
_21: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	*(*[20]int8)(unsafe.Pointer(bp + 268)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	*(*uint32_t)(unsafe.Pointer(bp + 264)) = libc.Xinet_addr(tls, __ccgo_ts+6245)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	tohex(tls, bp+268, bp+264, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if libc.Xstrcmp(tls, bp+268, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6252, libc.VaList(bp+4064, bp+268, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	r11 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+6245, bp+264)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if r11 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6326, libc.VaList(bp+4064, r11, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		goto _23
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	tohex(tls, bp+268, bp+264, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if libc.Xstrcmp(tls, bp+268, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6415, libc.VaList(bp+4064, bp+268, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	tobin(tls, bp+264, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+264, bp+268, uint32(20)) != bp+268 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6504, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if libc.Xstrcmp(tls, bp+268, __ccgo_ts+6245) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6600, libc.VaList(bp+4064, bp+268, __ccgo_ts+6245))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	in11.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 264))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	p11 = libc.Xinet_ntoa(tls, in11)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
	if libc.Xstrcmp(tls, p11, __ccgo_ts+6245) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:117:1:
		t_printf(tls, __ccgo_ts+6693, libc.VaList(bp+4064, p11, __ccgo_ts+6245))
	}
	goto _24
_24:
	goto _23
_23: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	*(*[20]int8)(unsafe.Pointer(bp + 292)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	*(*uint32_t)(unsafe.Pointer(bp + 288)) = libc.Xinet_addr(tls, __ccgo_ts+6771)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	tohex(tls, bp+292, bp+288, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if libc.Xstrcmp(tls, bp+292, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+6778, libc.VaList(bp+4064, bp+292, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	r12 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+6771, bp+288)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if r12 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+6852, libc.VaList(bp+4064, r12, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		goto _25
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	tohex(tls, bp+292, bp+288, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if libc.Xstrcmp(tls, bp+292, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+6941, libc.VaList(bp+4064, bp+292, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	tobin(tls, bp+288, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+288, bp+292, uint32(20)) != bp+292 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+7030, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if libc.Xstrcmp(tls, bp+292, __ccgo_ts+6771) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+7126, libc.VaList(bp+4064, bp+292, __ccgo_ts+6771))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	in12.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 288))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	p12 = libc.Xinet_ntoa(tls, in12)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
	if libc.Xstrcmp(tls, p12, __ccgo_ts+6771) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:118:1:
		t_printf(tls, __ccgo_ts+7219, libc.VaList(bp+4064, p12, __ccgo_ts+6771))
	}
	goto _26
_26:
	goto _25
_25: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	*(*[20]int8)(unsafe.Pointer(bp + 316)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	*(*uint32_t)(unsafe.Pointer(bp + 312)) = libc.Xinet_addr(tls, __ccgo_ts+7297)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	tohex(tls, bp+316, bp+312, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if libc.Xstrcmp(tls, bp+316, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7307, libc.VaList(bp+4064, bp+316, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	r13 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+7297, bp+312)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if r13 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7384, libc.VaList(bp+4064, r13, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		goto _27
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	tohex(tls, bp+316, bp+312, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if libc.Xstrcmp(tls, bp+316, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7476, libc.VaList(bp+4064, bp+316, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	tobin(tls, bp+312, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+312, bp+316, uint32(20)) != bp+316 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7568, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if libc.Xstrcmp(tls, bp+316, __ccgo_ts+7297) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7664, libc.VaList(bp+4064, bp+316, __ccgo_ts+7297))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	in13.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 312))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	p13 = libc.Xinet_ntoa(tls, in13)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
	if libc.Xstrcmp(tls, p13, __ccgo_ts+7297) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:119:1:
		t_printf(tls, __ccgo_ts+7757, libc.VaList(bp+4064, p13, __ccgo_ts+7297))
	}
	goto _28
_28:
	goto _27
_27: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	*(*[20]int8)(unsafe.Pointer(bp + 340)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	*(*uint32_t)(unsafe.Pointer(bp + 336)) = libc.Xinet_addr(tls, __ccgo_ts+7835)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	tohex(tls, bp+340, bp+336, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if libc.Xstrcmp(tls, bp+340, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+7843, libc.VaList(bp+4064, bp+340, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	r14 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+7835, bp+336)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if r14 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+7918, libc.VaList(bp+4064, r14, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		goto _29
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	tohex(tls, bp+340, bp+336, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if libc.Xstrcmp(tls, bp+340, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+8008, libc.VaList(bp+4064, bp+340, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	tobin(tls, bp+336, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+336, bp+340, uint32(20)) != bp+340 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+8098, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if libc.Xstrcmp(tls, bp+340, __ccgo_ts+7835) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+8194, libc.VaList(bp+4064, bp+340, __ccgo_ts+7835))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	in14.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 336))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	p14 = libc.Xinet_ntoa(tls, in14)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
	if libc.Xstrcmp(tls, p14, __ccgo_ts+7835) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:120:1:
		t_printf(tls, __ccgo_ts+8287, libc.VaList(bp+4064, p14, __ccgo_ts+7835))
	}
	goto _30
_30:
	goto _29
_29: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	*(*[20]int8)(unsafe.Pointer(bp + 364)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	*(*uint32_t)(unsafe.Pointer(bp + 360)) = libc.Xinet_addr(tls, __ccgo_ts+8365)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	tohex(tls, bp+364, bp+360, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if libc.Xstrcmp(tls, bp+364, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8375, libc.VaList(bp+4064, bp+364, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	r15 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+8365, bp+360)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if r15 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8452, libc.VaList(bp+4064, r15, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		goto _31
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	tohex(tls, bp+364, bp+360, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if libc.Xstrcmp(tls, bp+364, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8544, libc.VaList(bp+4064, bp+364, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	tobin(tls, bp+360, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+360, bp+364, uint32(20)) != bp+364 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8636, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if libc.Xstrcmp(tls, bp+364, __ccgo_ts+8365) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8732, libc.VaList(bp+4064, bp+364, __ccgo_ts+8365))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	in15.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 360))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	p15 = libc.Xinet_ntoa(tls, in15)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
	if libc.Xstrcmp(tls, p15, __ccgo_ts+8365) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:121:1:
		t_printf(tls, __ccgo_ts+8825, libc.VaList(bp+4064, p15, __ccgo_ts+8365))
	}
	goto _32
_32:
	goto _31
_31: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	*(*[20]int8)(unsafe.Pointer(bp + 388)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	*(*uint32_t)(unsafe.Pointer(bp + 384)) = libc.Xinet_addr(tls, __ccgo_ts+8903)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	tohex(tls, bp+388, bp+384, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if libc.Xstrcmp(tls, bp+388, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+8920, libc.VaList(bp+4064, bp+388, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	r16 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+8903, bp+384)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if r16 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+9004, libc.VaList(bp+4064, r16, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		goto _33
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	tohex(tls, bp+388, bp+384, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if libc.Xstrcmp(tls, bp+388, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+9103, libc.VaList(bp+4064, bp+388, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	tobin(tls, bp+384, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+384, bp+388, uint32(20)) != bp+388 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+9202, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if libc.Xstrcmp(tls, bp+388, __ccgo_ts+8903) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+9298, libc.VaList(bp+4064, bp+388, __ccgo_ts+8903))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	in16.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 384))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	p16 = libc.Xinet_ntoa(tls, in16)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
	if libc.Xstrcmp(tls, p16, __ccgo_ts+8903) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:122:1:
		t_printf(tls, __ccgo_ts+9391, libc.VaList(bp+4064, p16, __ccgo_ts+8903))
	}
	goto _34
_34:
	goto _33
_33: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	*(*[20]int8)(unsafe.Pointer(bp + 412)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	*(*uint32_t)(unsafe.Pointer(bp + 408)) = libc.Xinet_addr(tls, __ccgo_ts+9469)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	tohex(tls, bp+412, bp+408, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if libc.Xstrcmp(tls, bp+412, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9487, libc.VaList(bp+4064, bp+412, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	r17 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+9469, bp+408)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if r17 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9572, libc.VaList(bp+4064, r17, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		goto _35
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	tohex(tls, bp+412, bp+408, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if libc.Xstrcmp(tls, bp+412, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9672, libc.VaList(bp+4064, bp+412, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	tobin(tls, bp+408, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+408, bp+412, uint32(20)) != bp+412 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9772, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if libc.Xstrcmp(tls, bp+412, __ccgo_ts+9469) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9868, libc.VaList(bp+4064, bp+412, __ccgo_ts+9469))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	in17.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 408))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	p17 = libc.Xinet_ntoa(tls, in17)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
	if libc.Xstrcmp(tls, p17, __ccgo_ts+9469) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:123:1:
		t_printf(tls, __ccgo_ts+9961, libc.VaList(bp+4064, p17, __ccgo_ts+9469))
	}
	goto _36
_36:
	goto _35
_35: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	*(*[20]int8)(unsafe.Pointer(bp + 436)) = [20]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	*(*uint32_t)(unsafe.Pointer(bp + 432)) = libc.Xinet_addr(tls, __ccgo_ts+10039)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	tohex(tls, bp+436, bp+432, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if libc.Xstrcmp(tls, bp+436, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10048, libc.VaList(bp+4064, bp+436, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	r18 = libc.Xinet_pton(tls, int32(PF_INET), __ccgo_ts+10039, bp+432)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if r18 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10124, libc.VaList(bp+4064, r18, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		goto _37
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	tohex(tls, bp+436, bp+432, int32(4))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if libc.Xstrcmp(tls, bp+436, __ccgo_ts+1838) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10215, libc.VaList(bp+4064, bp+436, __ccgo_ts+1838))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	tobin(tls, bp+432, __ccgo_ts+1838)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if libc.Xinet_ntop(tls, int32(PF_INET), bp+432, bp+436, uint32(20)) != bp+436 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10306, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if libc.Xstrcmp(tls, bp+436, __ccgo_ts+10039) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10402, libc.VaList(bp+4064, bp+436, __ccgo_ts+10039))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	in18.Fs_addr = *(*uint32_t)(unsafe.Pointer(bp + 432))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	p18 = libc.Xinet_ntoa(tls, in18)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
	if libc.Xstrcmp(tls, p18, __ccgo_ts+10039) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:124:1:
		t_printf(tls, __ccgo_ts+10495, libc.VaList(bp+4064, p18, __ccgo_ts+10039))
	}
	goto _38
_38:
	goto _37
_37: /**/
	; //
	// ipv6
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	*(*[16]int8)(unsafe.Pointer(bp + 456)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	*(*[40]int8)(unsafe.Pointer(bp + 472)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	*(*[60]int8)(unsafe.Pointer(bp + 512)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	r19 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+10573, bp+456)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if r19 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+10575, libc.VaList(bp+4064, r19, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		goto _39
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	tohex(tls, bp+472, bp+456, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if libc.Xstrcmp(tls, bp+472, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+10660, libc.VaList(bp+4064, bp+472, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	tobin(tls, bp+456, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+456, bp+512, uint32(60)) != bp+512 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+10745, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+512, bp+456) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+10834, libc.VaList(bp+4064, bp+512))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	tohex(tls, bp+472, bp+456, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if libc.Xstrcmp(tls, bp+472, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+10940, libc.VaList(bp+4064, bp+512, bp+472, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+512, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:127:1:
		t_printf(tls, __ccgo_ts+11062, libc.VaList(bp+4064, bp+512))
	}
	goto _40
_40:
	goto _39
_39: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	*(*[16]int8)(unsafe.Pointer(bp + 572)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	*(*[40]int8)(unsafe.Pointer(bp + 588)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	*(*[60]int8)(unsafe.Pointer(bp + 628)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	r20 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+11162, bp+572)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if r20 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11165, libc.VaList(bp+4064, r20, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		goto _41
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	tohex(tls, bp+588, bp+572, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if libc.Xstrcmp(tls, bp+588, __ccgo_ts+11251) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11284, libc.VaList(bp+4064, bp+588, __ccgo_ts+11251))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	tobin(tls, bp+572, __ccgo_ts+11251)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+572, bp+628, uint32(60)) != bp+628 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11370, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+628, bp+572) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11491, libc.VaList(bp+4064, bp+628))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	tohex(tls, bp+588, bp+572, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if libc.Xstrcmp(tls, bp+588, __ccgo_ts+11251) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11629, libc.VaList(bp+4064, bp+628, bp+588, __ccgo_ts+11251))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
	if libc.Xstrncmp(tls, __ccgo_ts+11251, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+628, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:128:1:
		t_printf(tls, __ccgo_ts+11758, libc.VaList(bp+4064, bp+628))
	}
	goto _42
_42:
	goto _41
_41: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	*(*[16]int8)(unsafe.Pointer(bp + 688)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	*(*[40]int8)(unsafe.Pointer(bp + 704)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	*(*[60]int8)(unsafe.Pointer(bp + 744)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	r21 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+11890, bp+688)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if r21 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+11894, libc.VaList(bp+4064, r21, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		goto _43
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	tohex(tls, bp+704, bp+688, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if libc.Xstrcmp(tls, bp+704, __ccgo_ts+11981) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+12014, libc.VaList(bp+4064, bp+704, __ccgo_ts+11981))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	tobin(tls, bp+688, __ccgo_ts+11981)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+688, bp+744, uint32(60)) != bp+744 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+12101, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+744, bp+688) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+12222, libc.VaList(bp+4064, bp+744))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	tohex(tls, bp+704, bp+688, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if libc.Xstrcmp(tls, bp+704, __ccgo_ts+11981) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+12360, libc.VaList(bp+4064, bp+744, bp+704, __ccgo_ts+11981))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
	if libc.Xstrncmp(tls, __ccgo_ts+11981, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+744, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:129:1:
		t_printf(tls, __ccgo_ts+12489, libc.VaList(bp+4064, bp+744))
	}
	goto _44
_44:
	goto _43
_43: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	*(*[16]int8)(unsafe.Pointer(bp + 804)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	*(*[40]int8)(unsafe.Pointer(bp + 820)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	*(*[60]int8)(unsafe.Pointer(bp + 860)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	r22 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+12621, bp+804)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if r22 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+12625, libc.VaList(bp+4064, r22, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		goto _45
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	tohex(tls, bp+820, bp+804, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if libc.Xstrcmp(tls, bp+820, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+12712, libc.VaList(bp+4064, bp+820, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	tobin(tls, bp+804, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+804, bp+860, uint32(60)) != bp+860 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+12799, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+860, bp+804) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+12888, libc.VaList(bp+4064, bp+860))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	tohex(tls, bp+820, bp+804, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if libc.Xstrcmp(tls, bp+820, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+12994, libc.VaList(bp+4064, bp+860, bp+820, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+860, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:130:1:
		t_printf(tls, __ccgo_ts+13091, libc.VaList(bp+4064, bp+860))
	}
	goto _46
_46:
	goto _45
_45: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	*(*[16]int8)(unsafe.Pointer(bp + 920)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	*(*[40]int8)(unsafe.Pointer(bp + 936)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	*(*[60]int8)(unsafe.Pointer(bp + 976)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	r23 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+13191, bp+920)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if r23 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13203, libc.VaList(bp+4064, r23, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		goto _47
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	tohex(tls, bp+936, bp+920, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if libc.Xstrcmp(tls, bp+936, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13298, libc.VaList(bp+4064, bp+936, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	tobin(tls, bp+920, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+920, bp+976, uint32(60)) != bp+976 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13393, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+976, bp+920) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13482, libc.VaList(bp+4064, bp+976))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	tohex(tls, bp+936, bp+920, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if libc.Xstrcmp(tls, bp+936, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13588, libc.VaList(bp+4064, bp+976, bp+936, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+976, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:131:1:
		t_printf(tls, __ccgo_ts+13685, libc.VaList(bp+4064, bp+976))
	}
	goto _48
_48:
	goto _47
_47: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1036)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1052)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1092)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	r24 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+13785, bp+1036)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if r24 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+13798, libc.VaList(bp+4064, r24, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		goto _49
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	tohex(tls, bp+1052, bp+1036, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if libc.Xstrcmp(tls, bp+1052, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+13894, libc.VaList(bp+4064, bp+1052, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	tobin(tls, bp+1036, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1036, bp+1092, uint32(60)) != bp+1092 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+13990, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1092, bp+1036) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+14079, libc.VaList(bp+4064, bp+1092))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	tohex(tls, bp+1052, bp+1036, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if libc.Xstrcmp(tls, bp+1052, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+14185, libc.VaList(bp+4064, bp+1092, bp+1052, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1092, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:132:1:
		t_printf(tls, __ccgo_ts+14282, libc.VaList(bp+4064, bp+1092))
	}
	goto _50
_50:
	goto _49
_49: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1152)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1168)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1208)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	r25 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+14382, bp+1152)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if r25 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+14396, libc.VaList(bp+4064, r25, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		goto _51
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	tohex(tls, bp+1168, bp+1152, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if libc.Xstrcmp(tls, bp+1168, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+14526, libc.VaList(bp+4064, bp+1168, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	tobin(tls, bp+1152, __ccgo_ts+14493)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1152, bp+1208, uint32(60)) != bp+1208 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+14623, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1208, bp+1152) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+14744, libc.VaList(bp+4064, bp+1208))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	tohex(tls, bp+1168, bp+1152, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if libc.Xstrcmp(tls, bp+1168, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+14882, libc.VaList(bp+4064, bp+1208, bp+1168, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
	if libc.Xstrncmp(tls, __ccgo_ts+14493, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1208, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:133:1:
		t_printf(tls, __ccgo_ts+15011, libc.VaList(bp+4064, bp+1208))
	}
	goto _52
_52:
	goto _51
_51: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1268)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1284)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1324)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	r26 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+15143, bp+1268)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if r26 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15167, libc.VaList(bp+4064, r26, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		goto _53
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	tohex(tls, bp+1284, bp+1268, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if libc.Xstrcmp(tls, bp+1284, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15274, libc.VaList(bp+4064, bp+1284, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	tobin(tls, bp+1268, __ccgo_ts+14493)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1268, bp+1324, uint32(60)) != bp+1324 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15381, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1324, bp+1268) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15502, libc.VaList(bp+4064, bp+1324))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	tohex(tls, bp+1284, bp+1268, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if libc.Xstrcmp(tls, bp+1284, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15640, libc.VaList(bp+4064, bp+1324, bp+1284, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
	if libc.Xstrncmp(tls, __ccgo_ts+14493, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1324, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:134:1:
		t_printf(tls, __ccgo_ts+15769, libc.VaList(bp+4064, bp+1324))
	}
	goto _54
_54:
	goto _53
_53: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1384)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1400)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1440)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	r27 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+15901, bp+1384)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if r27 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+15924, libc.VaList(bp+4064, r27, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		goto _55
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	tohex(tls, bp+1400, bp+1384, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if libc.Xstrcmp(tls, bp+1400, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+16030, libc.VaList(bp+4064, bp+1400, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	tobin(tls, bp+1384, __ccgo_ts+14493)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1384, bp+1440, uint32(60)) != bp+1440 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+16136, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1440, bp+1384) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+16257, libc.VaList(bp+4064, bp+1440))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	tohex(tls, bp+1400, bp+1384, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if libc.Xstrcmp(tls, bp+1400, __ccgo_ts+14493) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+16395, libc.VaList(bp+4064, bp+1440, bp+1400, __ccgo_ts+14493))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
	if libc.Xstrncmp(tls, __ccgo_ts+14493, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1440, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:135:1:
		t_printf(tls, __ccgo_ts+16524, libc.VaList(bp+4064, bp+1440))
	}
	goto _56
_56:
	goto _55
_55: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1500)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1516)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1556)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	r28 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+16656, bp+1500)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if r28 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+16671, libc.VaList(bp+4064, r28, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		goto _57
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	tohex(tls, bp+1516, bp+1500, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if libc.Xstrcmp(tls, bp+1516, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+16769, libc.VaList(bp+4064, bp+1516, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	tobin(tls, bp+1500, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1500, bp+1556, uint32(60)) != bp+1556 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+16867, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1556, bp+1500) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+16956, libc.VaList(bp+4064, bp+1556))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	tohex(tls, bp+1516, bp+1500, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if libc.Xstrcmp(tls, bp+1516, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+17062, libc.VaList(bp+4064, bp+1556, bp+1516, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1556, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:136:1:
		t_printf(tls, __ccgo_ts+17159, libc.VaList(bp+4064, bp+1556))
	}
	goto _58
_58:
	goto _57
_57: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1616)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1632)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1672)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	r29 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+17259, bp+1616)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if r29 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17277, libc.VaList(bp+4064, r29, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		goto _59
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	tohex(tls, bp+1632, bp+1616, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if libc.Xstrcmp(tls, bp+1632, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17378, libc.VaList(bp+4064, bp+1632, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	tobin(tls, bp+1616, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1616, bp+1672, uint32(60)) != bp+1672 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17479, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1672, bp+1616) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17568, libc.VaList(bp+4064, bp+1672))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	tohex(tls, bp+1632, bp+1616, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if libc.Xstrcmp(tls, bp+1632, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17674, libc.VaList(bp+4064, bp+1672, bp+1632, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1672, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:137:1:
		t_printf(tls, __ccgo_ts+17771, libc.VaList(bp+4064, bp+1672))
	}
	goto _60
_60:
	goto _59
_59: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1732)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1748)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1788)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	r30 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+17871, bp+1732)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if r30 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+17890, libc.VaList(bp+4064, r30, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		goto _61
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	tohex(tls, bp+1748, bp+1732, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if libc.Xstrcmp(tls, bp+1748, __ccgo_ts+17992) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+18025, libc.VaList(bp+4064, bp+1748, __ccgo_ts+17992))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	tobin(tls, bp+1732, __ccgo_ts+17992)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1732, bp+1788, uint32(60)) != bp+1788 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+18127, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1788, bp+1732) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+18248, libc.VaList(bp+4064, bp+1788))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	tohex(tls, bp+1748, bp+1732, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if libc.Xstrcmp(tls, bp+1748, __ccgo_ts+17992) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+18386, libc.VaList(bp+4064, bp+1788, bp+1748, __ccgo_ts+17992))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
	if libc.Xstrncmp(tls, __ccgo_ts+17992, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1788, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:138:1:
		t_printf(tls, __ccgo_ts+18515, libc.VaList(bp+4064, bp+1788))
	}
	goto _62
_62:
	goto _61
_61: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1848)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1864)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	*(*[60]int8)(unsafe.Pointer(bp + 1904)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	r31 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+18647, bp+1848)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if r31 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+18660, libc.VaList(bp+4064, r31, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		goto _63
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	tohex(tls, bp+1864, bp+1848, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if libc.Xstrcmp(tls, bp+1864, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+18756, libc.VaList(bp+4064, bp+1864, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	tobin(tls, bp+1848, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1848, bp+1904, uint32(60)) != bp+1904 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+18852, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+1904, bp+1848) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+18941, libc.VaList(bp+4064, bp+1904))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	tohex(tls, bp+1864, bp+1848, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if libc.Xstrcmp(tls, bp+1864, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+19047, libc.VaList(bp+4064, bp+1904, bp+1864, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+1904, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:139:1:
		t_printf(tls, __ccgo_ts+19144, libc.VaList(bp+4064, bp+1904))
	}
	goto _64
_64:
	goto _63
_63: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	*(*[16]int8)(unsafe.Pointer(bp + 1964)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	*(*[40]int8)(unsafe.Pointer(bp + 1980)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2020)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	r32 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+19244, bp+1964)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if r32 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19258, libc.VaList(bp+4064, r32, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		goto _65
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	tohex(tls, bp+1980, bp+1964, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if libc.Xstrcmp(tls, bp+1980, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19355, libc.VaList(bp+4064, bp+1980, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	tobin(tls, bp+1964, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+1964, bp+2020, uint32(60)) != bp+2020 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19452, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2020, bp+1964) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19541, libc.VaList(bp+4064, bp+2020))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	tohex(tls, bp+1980, bp+1964, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if libc.Xstrcmp(tls, bp+1980, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19647, libc.VaList(bp+4064, bp+2020, bp+1980, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2020, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:140:1:
		t_printf(tls, __ccgo_ts+19744, libc.VaList(bp+4064, bp+2020))
	}
	goto _66
_66:
	goto _65
_65: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2080)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2096)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2136)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	r33 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+19844, bp+2080)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if r33 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+19864, libc.VaList(bp+4064, r33, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		goto _67
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	tohex(tls, bp+2096, bp+2080, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if libc.Xstrcmp(tls, bp+2096, __ccgo_ts+19967) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+20000, libc.VaList(bp+4064, bp+2096, __ccgo_ts+19967))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	tobin(tls, bp+2080, __ccgo_ts+19967)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2080, bp+2136, uint32(60)) != bp+2136 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+20103, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2136, bp+2080) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+20224, libc.VaList(bp+4064, bp+2136))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	tohex(tls, bp+2096, bp+2080, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if libc.Xstrcmp(tls, bp+2096, __ccgo_ts+19967) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+20362, libc.VaList(bp+4064, bp+2136, bp+2096, __ccgo_ts+19967))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
	if libc.Xstrncmp(tls, __ccgo_ts+19967, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2136, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:141:1:
		t_printf(tls, __ccgo_ts+20491, libc.VaList(bp+4064, bp+2136))
	}
	goto _68
_68:
	goto _67
_67: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2196)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2212)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2252)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	r34 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+20623, bp+2196)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if r34 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+20647, libc.VaList(bp+4064, r34, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		goto _69
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	tohex(tls, bp+2212, bp+2196, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if libc.Xstrcmp(tls, bp+2212, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+20754, libc.VaList(bp+4064, bp+2212, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	tobin(tls, bp+2196, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2196, bp+2252, uint32(60)) != bp+2252 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+20861, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2252, bp+2196) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+20950, libc.VaList(bp+4064, bp+2252))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	tohex(tls, bp+2212, bp+2196, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if libc.Xstrcmp(tls, bp+2212, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+21056, libc.VaList(bp+4064, bp+2252, bp+2212, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2252, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:142:1:
		t_printf(tls, __ccgo_ts+21153, libc.VaList(bp+4064, bp+2252))
	}
	goto _70
_70:
	goto _69
_69: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2312)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2328)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2368)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	r35 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+21253, bp+2312)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if r35 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21267, libc.VaList(bp+4064, r35, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		goto _71
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	tohex(tls, bp+2328, bp+2312, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if libc.Xstrcmp(tls, bp+2328, __ccgo_ts+21364) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21397, libc.VaList(bp+4064, bp+2328, __ccgo_ts+21364))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	tobin(tls, bp+2312, __ccgo_ts+21364)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2312, bp+2368, uint32(60)) != bp+2368 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21494, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2368, bp+2312) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21615, libc.VaList(bp+4064, bp+2368))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	tohex(tls, bp+2328, bp+2312, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if libc.Xstrcmp(tls, bp+2328, __ccgo_ts+21364) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21753, libc.VaList(bp+4064, bp+2368, bp+2328, __ccgo_ts+21364))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
	if libc.Xstrncmp(tls, __ccgo_ts+21364, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2368, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:143:1:
		t_printf(tls, __ccgo_ts+21882, libc.VaList(bp+4064, bp+2368))
	}
	goto _72
_72:
	goto _71
_71: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2428)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2444)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2484)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	r36 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+22014, bp+2428)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if r36 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22030, libc.VaList(bp+4064, r36, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		goto _73
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	tohex(tls, bp+2444, bp+2428, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if libc.Xstrcmp(tls, bp+2444, __ccgo_ts+22129) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22162, libc.VaList(bp+4064, bp+2444, __ccgo_ts+22129))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	tobin(tls, bp+2428, __ccgo_ts+22129)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2428, bp+2484, uint32(60)) != bp+2484 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22261, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2484, bp+2428) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22382, libc.VaList(bp+4064, bp+2484))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	tohex(tls, bp+2444, bp+2428, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if libc.Xstrcmp(tls, bp+2444, __ccgo_ts+22129) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22520, libc.VaList(bp+4064, bp+2484, bp+2444, __ccgo_ts+22129))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
	if libc.Xstrncmp(tls, __ccgo_ts+22129, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2484, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:144:1:
		t_printf(tls, __ccgo_ts+22649, libc.VaList(bp+4064, bp+2484))
	}
	goto _74
_74:
	goto _73
_73: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2544)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2560)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2600)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	r37 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+22781, bp+2544)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if r37 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+22799, libc.VaList(bp+4064, r37, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		goto _75
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	tohex(tls, bp+2560, bp+2544, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if libc.Xstrcmp(tls, bp+2560, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+22900, libc.VaList(bp+4064, bp+2560, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	tobin(tls, bp+2544, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2544, bp+2600, uint32(60)) != bp+2600 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+23001, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2600, bp+2544) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+23090, libc.VaList(bp+4064, bp+2600))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	tohex(tls, bp+2560, bp+2544, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if libc.Xstrcmp(tls, bp+2560, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+23196, libc.VaList(bp+4064, bp+2600, bp+2560, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2600, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:145:1:
		t_printf(tls, __ccgo_ts+23293, libc.VaList(bp+4064, bp+2600))
	}
	goto _76
_76:
	goto _75
_75: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2660)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2676)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2716)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	r38 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+23393, bp+2660)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if r38 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23410, libc.VaList(bp+4064, r38, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		goto _77
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	tohex(tls, bp+2676, bp+2660, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if libc.Xstrcmp(tls, bp+2676, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23510, libc.VaList(bp+4064, bp+2676, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	tobin(tls, bp+2660, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2660, bp+2716, uint32(60)) != bp+2716 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23610, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2716, bp+2660) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23699, libc.VaList(bp+4064, bp+2716))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	tohex(tls, bp+2676, bp+2660, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if libc.Xstrcmp(tls, bp+2676, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23805, libc.VaList(bp+4064, bp+2716, bp+2676, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2716, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:146:1:
		t_printf(tls, __ccgo_ts+23902, libc.VaList(bp+4064, bp+2716))
	}
	goto _78
_78:
	goto _77
_77: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2776)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2792)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2832)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	r39 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+24002, bp+2776)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if r39 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24016, libc.VaList(bp+4064, r39, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		goto _79
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	tohex(tls, bp+2792, bp+2776, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if libc.Xstrcmp(tls, bp+2792, __ccgo_ts+24113) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24146, libc.VaList(bp+4064, bp+2792, __ccgo_ts+24113))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	tobin(tls, bp+2776, __ccgo_ts+24113)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2776, bp+2832, uint32(60)) != bp+2832 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24243, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2832, bp+2776) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24364, libc.VaList(bp+4064, bp+2832))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	tohex(tls, bp+2792, bp+2776, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if libc.Xstrcmp(tls, bp+2792, __ccgo_ts+24113) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24502, libc.VaList(bp+4064, bp+2832, bp+2792, __ccgo_ts+24113))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
	if libc.Xstrncmp(tls, __ccgo_ts+24113, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2832, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:147:1:
		t_printf(tls, __ccgo_ts+24631, libc.VaList(bp+4064, bp+2832))
	}
	goto _80
_80:
	goto _79
_79: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	*(*[16]int8)(unsafe.Pointer(bp + 2892)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	*(*[40]int8)(unsafe.Pointer(bp + 2908)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	*(*[60]int8)(unsafe.Pointer(bp + 2948)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	r40 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+24763, bp+2892)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if r40 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+24779, libc.VaList(bp+4064, r40, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		goto _81
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	tohex(tls, bp+2908, bp+2892, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if libc.Xstrcmp(tls, bp+2908, __ccgo_ts+24878) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+24911, libc.VaList(bp+4064, bp+2908, __ccgo_ts+24878))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	tobin(tls, bp+2892, __ccgo_ts+24878)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+2892, bp+2948, uint32(60)) != bp+2948 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+25010, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+2948, bp+2892) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+25131, libc.VaList(bp+4064, bp+2948))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	tohex(tls, bp+2908, bp+2892, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if libc.Xstrcmp(tls, bp+2908, __ccgo_ts+24878) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+25269, libc.VaList(bp+4064, bp+2948, bp+2908, __ccgo_ts+24878))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
	if libc.Xstrncmp(tls, __ccgo_ts+24878, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+2948, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:148:1:
		t_printf(tls, __ccgo_ts+25398, libc.VaList(bp+4064, bp+2948))
	}
	goto _82
_82:
	goto _81
_81: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3008)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3024)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3064)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	r41 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+25530, bp+3008)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if r41 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+25548, libc.VaList(bp+4064, r41, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		goto _83
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	tohex(tls, bp+3024, bp+3008, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if libc.Xstrcmp(tls, bp+3024, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+25649, libc.VaList(bp+4064, bp+3024, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	tobin(tls, bp+3008, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3008, bp+3064, uint32(60)) != bp+3064 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+25750, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3064, bp+3008) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+25839, libc.VaList(bp+4064, bp+3064))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	tohex(tls, bp+3024, bp+3008, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if libc.Xstrcmp(tls, bp+3024, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+25945, libc.VaList(bp+4064, bp+3064, bp+3024, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3064, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:149:1:
		t_printf(tls, __ccgo_ts+26042, libc.VaList(bp+4064, bp+3064))
	}
	goto _84
_84:
	goto _83
_83: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3124)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3140)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3180)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	r42 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+26142, bp+3124)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if r42 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26155, libc.VaList(bp+4064, r42, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		goto _85
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	tohex(tls, bp+3140, bp+3124, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if libc.Xstrcmp(tls, bp+3140, __ccgo_ts+26251) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26284, libc.VaList(bp+4064, bp+3140, __ccgo_ts+26251))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	tobin(tls, bp+3124, __ccgo_ts+26251)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3124, bp+3180, uint32(60)) != bp+3180 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26380, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3180, bp+3124) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26501, libc.VaList(bp+4064, bp+3180))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	tohex(tls, bp+3140, bp+3124, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if libc.Xstrcmp(tls, bp+3140, __ccgo_ts+26251) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26639, libc.VaList(bp+4064, bp+3180, bp+3140, __ccgo_ts+26251))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
	if libc.Xstrncmp(tls, __ccgo_ts+26251, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3180, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:150:1:
		t_printf(tls, __ccgo_ts+26768, libc.VaList(bp+4064, bp+3180))
	}
	goto _86
_86:
	goto _85
_85: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3240)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3256)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3296)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	r43 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+26900, bp+3240)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if r43 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+26914, libc.VaList(bp+4064, r43, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		goto _87
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	tohex(tls, bp+3256, bp+3240, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if libc.Xstrcmp(tls, bp+3256, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+27011, libc.VaList(bp+4064, bp+3256, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	tobin(tls, bp+3240, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3240, bp+3296, uint32(60)) != bp+3296 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+27108, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3296, bp+3240) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+27197, libc.VaList(bp+4064, bp+3296))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	tohex(tls, bp+3256, bp+3240, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if libc.Xstrcmp(tls, bp+3256, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+27303, libc.VaList(bp+4064, bp+3296, bp+3256, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3296, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:151:1:
		t_printf(tls, __ccgo_ts+27400, libc.VaList(bp+4064, bp+3296))
	}
	goto _88
_88:
	goto _87
_87: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3356)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3372)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3412)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	r44 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+27500, bp+3356)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if r44 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+27515, libc.VaList(bp+4064, r44, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		goto _89
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	tohex(tls, bp+3372, bp+3356, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if libc.Xstrcmp(tls, bp+3372, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+27613, libc.VaList(bp+4064, bp+3372, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	tobin(tls, bp+3356, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3356, bp+3412, uint32(60)) != bp+3412 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+27711, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3412, bp+3356) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+27800, libc.VaList(bp+4064, bp+3412))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	tohex(tls, bp+3372, bp+3356, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if libc.Xstrcmp(tls, bp+3372, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+27906, libc.VaList(bp+4064, bp+3412, bp+3372, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3412, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:152:1:
		t_printf(tls, __ccgo_ts+28003, libc.VaList(bp+4064, bp+3412))
	}
	goto _90
_90:
	goto _89
_89: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3472)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3488)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3528)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	r45 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+28103, bp+3472)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if r45 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28127, libc.VaList(bp+4064, r45, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		goto _91
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	tohex(tls, bp+3488, bp+3472, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if libc.Xstrcmp(tls, bp+3488, __ccgo_ts+28234) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28267, libc.VaList(bp+4064, bp+3488, __ccgo_ts+28234))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	tobin(tls, bp+3472, __ccgo_ts+28234)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3472, bp+3528, uint32(60)) != bp+3528 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28374, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3528, bp+3472) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28495, libc.VaList(bp+4064, bp+3528))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	tohex(tls, bp+3488, bp+3472, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if libc.Xstrcmp(tls, bp+3488, __ccgo_ts+28234) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28633, libc.VaList(bp+4064, bp+3528, bp+3488, __ccgo_ts+28234))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
	if libc.Xstrncmp(tls, __ccgo_ts+28234, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3528, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:153:1:
		t_printf(tls, __ccgo_ts+28762, libc.VaList(bp+4064, bp+3528))
	}
	goto _92
_92:
	goto _91
_91: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3588)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3604)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3644)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	r46 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+28894, bp+3588)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if r46 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+28916, libc.VaList(bp+4064, r46, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		goto _93
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	tohex(tls, bp+3604, bp+3588, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if libc.Xstrcmp(tls, bp+3604, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+29021, libc.VaList(bp+4064, bp+3604, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	tobin(tls, bp+3588, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3588, bp+3644, uint32(60)) != bp+3644 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+29126, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3644, bp+3588) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+29215, libc.VaList(bp+4064, bp+3644))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	tohex(tls, bp+3604, bp+3588, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if libc.Xstrcmp(tls, bp+3604, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+29321, libc.VaList(bp+4064, bp+3644, bp+3604, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3644, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:154:1:
		t_printf(tls, __ccgo_ts+29418, libc.VaList(bp+4064, bp+3644))
	}
	goto _94
_94:
	goto _93
_93: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3704)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3720)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3760)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	r47 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+29518, bp+3704)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if r47 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+29535, libc.VaList(bp+4064, r47, int32(1)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if int32(1) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		goto _95
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	tohex(tls, bp+3720, bp+3704, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if libc.Xstrcmp(tls, bp+3720, __ccgo_ts+28234) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+29635, libc.VaList(bp+4064, bp+3720, __ccgo_ts+28234))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	tobin(tls, bp+3704, __ccgo_ts+28234)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3704, bp+3760, uint32(60)) != bp+3760 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+29735, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3760, bp+3704) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+29856, libc.VaList(bp+4064, bp+3760))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	tohex(tls, bp+3720, bp+3704, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if libc.Xstrcmp(tls, bp+3720, __ccgo_ts+28234) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+29994, libc.VaList(bp+4064, bp+3760, bp+3720, __ccgo_ts+28234))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
	if libc.Xstrncmp(tls, __ccgo_ts+28234, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3760, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:155:1:
		t_printf(tls, __ccgo_ts+30123, libc.VaList(bp+4064, bp+3760))
	}
	goto _96
_96:
	goto _95
_95: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3820)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3836)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3876)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	r48 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+30255, bp+3820)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if r48 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30274, libc.VaList(bp+4064, r48, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		goto _97
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	tohex(tls, bp+3836, bp+3820, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if libc.Xstrcmp(tls, bp+3836, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30376, libc.VaList(bp+4064, bp+3836, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	tobin(tls, bp+3820, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3820, bp+3876, uint32(60)) != bp+3876 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30478, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3876, bp+3820) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30567, libc.VaList(bp+4064, bp+3876))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	tohex(tls, bp+3836, bp+3820, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if libc.Xstrcmp(tls, bp+3836, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30673, libc.VaList(bp+4064, bp+3876, bp+3836, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3876, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:156:1:
		t_printf(tls, __ccgo_ts+30770, libc.VaList(bp+4064, bp+3876))
	}
	goto _98
_98:
	goto _97
_97: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	*(*[16]int8)(unsafe.Pointer(bp + 3936)) = [16]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	*(*[40]int8)(unsafe.Pointer(bp + 3952)) = [40]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	*(*[60]int8)(unsafe.Pointer(bp + 3992)) = [60]int8{}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	r49 = libc.Xinet_pton(tls, int32(PF_INET6), __ccgo_ts+30870, bp+3936)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if r49 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+30875, libc.VaList(bp+4064, r49, 0))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if 0 != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		goto _99
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	tohex(tls, bp+3952, bp+3936, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if libc.Xstrcmp(tls, bp+3952, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+30963, libc.VaList(bp+4064, bp+3952, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	tobin(tls, bp+3936, __ccgo_ts+5)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if libc.Xinet_ntop(tls, int32(PF_INET6), bp+3936, bp+3992, uint32(60)) != bp+3992 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+31051, 0)
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if libc.Xinet_pton(tls, int32(PF_INET6), bp+3992, bp+3936) != int32(1) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+31140, libc.VaList(bp+4064, bp+3992))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	tohex(tls, bp+3952, bp+3936, int32(16))
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if libc.Xstrcmp(tls, bp+3952, __ccgo_ts+5) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+31246, libc.VaList(bp+4064, bp+3992, bp+3952, __ccgo_ts+5))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
	if libc.Xstrncmp(tls, __ccgo_ts+5, __ccgo_ts+11037, uint32(24)) == 0 && !(libc.Xstrchr(tls, bp+3992, int32('.')) != 0) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:157:1:
		t_printf(tls, __ccgo_ts+31343, libc.VaList(bp+4064, bp+3992))
	}
	goto _100
_100:
	goto _99
_99: /**/
	; //
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/inet_pton.c:159:1:
	return libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&t_status)))
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
		fd = libc.Xopen(tls, __ccgo_ts+31443, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+31453, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:34:9:
const FP_INFINITE = 1

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/math.h:33:9:
const FP_NAN = 0

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
		Fs:    __ccgo_ts + 31497,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 31505,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 31513,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 31523,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 31533,
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
				v2 = __ccgo_ts + 31542
			} else {
				v2 = __ccgo_ts + 5
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+31544, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 31542
		} else {
			v3 = __ccgo_ts + 5
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+31549, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 5
	} else {
		v4 = __ccgo_ts + 31554
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+31556, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 31559
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 31562
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 31565
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 31568
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 31571
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+31574, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+31582, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+31587, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+31630, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31679) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31687) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31699) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31711) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31723) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+31732) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+5) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+31732) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+31738, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "%02x\x00\x00src/functional/inet_pton.c:95: inet_pton(12345,,) should fail with EAFNOSUPPORT, got %s\n\x00xxxx\x00src/functional/inet_pton.c:98: inet_ntop(,,0,0) should fail with ENOSPC, got %s\n\x000.0.0.0\x0000000000\x00src/functional/inet_pton.c:102: inet_addr(\"0.0.0.0\") returned %s, want %s\n\x00src/functional/inet_pton.c:102: inet_pton(AF_INET, \"0.0.0.0\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:102: inet_pton(AF_INET, \"0.0.0.0\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:102: inet_ntop(AF_INET, <\"00000000\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:102: inet_ntop(AF_INET, <\"00000000\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:102: inet_ntoa(<\"00000000\">) returned %s, want %s\n\x00127.0.0.1\x007f000001\x00src/functional/inet_pton.c:103: inet_addr(\"127.0.0.1\") returned %s, want %s\n\x00src/functional/inet_pton.c:103: inet_pton(AF_INET, \"127.0.0.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:103: inet_pton(AF_INET, \"127.0.0.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:103: inet_ntop(AF_INET, <\"7f000001\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:103: inet_ntop(AF_INET, <\"7f000001\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:103: inet_ntoa(<\"7f000001\">) returned %s, want %s\n\x0010.0.128.31\x000a00801f\x00src/functional/inet_pton.c:104: inet_addr(\"10.0.128.31\") returned %s, want %s\n\x00src/functional/inet_pton.c:104: inet_pton(AF_INET, \"10.0.128.31\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:104: inet_pton(AF_INET, \"10.0.128.31\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:104: inet_ntop(AF_INET, <\"0a00801f\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:104: inet_ntop(AF_INET, <\"0a00801f\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:104: inet_ntoa(<\"0a00801f\">) returned %s, want %s\n\x00255.255.255.255\x00ffffffff\x00src/functional/inet_pton.c:105: inet_addr(\"255.255.255.255\") returned %s, want %s\n\x00src/functional/inet_pton.c:105: inet_pton(AF_INET, \"255.255.255.255\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:105: inet_pton(AF_INET, \"255.255.255.255\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:105: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:105: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:105: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.03.4\x0001020304\x00src/functional/inet_pton.c:108: inet_addr(\"1.2.03.4\") returned %s, want %s\n\x00src/functional/inet_pton.c:108: inet_pton(AF_INET, \"1.2.03.4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:108: inet_pton(AF_INET, \"1.2.03.4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:108: inet_ntop(AF_INET, <\"01020304\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:108: inet_ntop(AF_INET, <\"01020304\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:108: inet_ntoa(<\"01020304\">) returned %s, want %s\n\x001.2.0x33.4\x0001023304\x00src/functional/inet_pton.c:109: inet_addr(\"1.2.0x33.4\") returned %s, want %s\n\x00src/functional/inet_pton.c:109: inet_pton(AF_INET, \"1.2.0x33.4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:109: inet_pton(AF_INET, \"1.2.0x33.4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:109: inet_ntop(AF_INET, <\"01023304\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:109: inet_ntop(AF_INET, <\"01023304\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:109: inet_ntoa(<\"01023304\">) returned %s, want %s\n\x001.2.0XAB.4\x000102ab04\x00src/functional/inet_pton.c:110: inet_addr(\"1.2.0XAB.4\") returned %s, want %s\n\x00src/functional/inet_pton.c:110: inet_pton(AF_INET, \"1.2.0XAB.4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:110: inet_pton(AF_INET, \"1.2.0XAB.4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:110: inet_ntop(AF_INET, <\"0102ab04\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:110: inet_ntop(AF_INET, <\"0102ab04\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:110: inet_ntoa(<\"0102ab04\">) returned %s, want %s\n\x001.2.0xabcd\x000102abcd\x00src/functional/inet_pton.c:111: inet_addr(\"1.2.0xabcd\") returned %s, want %s\n\x00src/functional/inet_pton.c:111: inet_pton(AF_INET, \"1.2.0xabcd\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:111: inet_pton(AF_INET, \"1.2.0xabcd\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:111: inet_ntop(AF_INET, <\"0102abcd\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:111: inet_ntop(AF_INET, <\"0102abcd\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:111: inet_ntoa(<\"0102abcd\">) returned %s, want %s\n\x001.0xabcdef\x0001abcdef\x00src/functional/inet_pton.c:112: inet_addr(\"1.0xabcdef\") returned %s, want %s\n\x00src/functional/inet_pton.c:112: inet_pton(AF_INET, \"1.0xabcdef\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:112: inet_pton(AF_INET, \"1.0xabcdef\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:112: inet_ntop(AF_INET, <\"01abcdef\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:112: inet_ntop(AF_INET, <\"01abcdef\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:112: inet_ntoa(<\"01abcdef\">) returned %s, want %s\n\x0000377.0x0ff.65534\x00fffffffe\x00src/functional/inet_pton.c:113: inet_addr(\"00377.0x0ff.65534\") returned %s, want %s\n\x00src/functional/inet_pton.c:113: inet_pton(AF_INET, \"00377.0x0ff.65534\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:113: inet_pton(AF_INET, \"00377.0x0ff.65534\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:113: inet_ntop(AF_INET, <\"fffffffe\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:113: inet_ntop(AF_INET, <\"fffffffe\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:113: inet_ntoa(<\"fffffffe\">) returned %s, want %s\n\x00.1.2.3\x00src/functional/inet_pton.c:116: inet_addr(\".1.2.3\") returned %s, want %s\n\x00src/functional/inet_pton.c:116: inet_pton(AF_INET, \".1.2.3\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:116: inet_pton(AF_INET, \".1.2.3\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:116: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:116: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:116: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001..2.3\x00src/functional/inet_pton.c:117: inet_addr(\"1..2.3\") returned %s, want %s\n\x00src/functional/inet_pton.c:117: inet_pton(AF_INET, \"1..2.3\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:117: inet_pton(AF_INET, \"1..2.3\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:117: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:117: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:117: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.3.\x00src/functional/inet_pton.c:118: inet_addr(\"1.2.3.\") returned %s, want %s\n\x00src/functional/inet_pton.c:118: inet_pton(AF_INET, \"1.2.3.\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:118: inet_pton(AF_INET, \"1.2.3.\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:118: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:118: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:118: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.3.4.5\x00src/functional/inet_pton.c:119: inet_addr(\"1.2.3.4.5\") returned %s, want %s\n\x00src/functional/inet_pton.c:119: inet_pton(AF_INET, \"1.2.3.4.5\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:119: inet_pton(AF_INET, \"1.2.3.4.5\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:119: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:119: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:119: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.3.a\x00src/functional/inet_pton.c:120: inet_addr(\"1.2.3.a\") returned %s, want %s\n\x00src/functional/inet_pton.c:120: inet_pton(AF_INET, \"1.2.3.a\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:120: inet_pton(AF_INET, \"1.2.3.a\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:120: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:120: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:120: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.256.2.3\x00src/functional/inet_pton.c:121: inet_addr(\"1.256.2.3\") returned %s, want %s\n\x00src/functional/inet_pton.c:121: inet_pton(AF_INET, \"1.256.2.3\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:121: inet_pton(AF_INET, \"1.256.2.3\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:121: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:121: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:121: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.4294967296.3\x00src/functional/inet_pton.c:122: inet_addr(\"1.2.4294967296.3\") returned %s, want %s\n\x00src/functional/inet_pton.c:122: inet_pton(AF_INET, \"1.2.4294967296.3\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:122: inet_pton(AF_INET, \"1.2.4294967296.3\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:122: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:122: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:122: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2.-4294967295.3\x00src/functional/inet_pton.c:123: inet_addr(\"1.2.-4294967295.3\") returned %s, want %s\n\x00src/functional/inet_pton.c:123: inet_pton(AF_INET, \"1.2.-4294967295.3\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:123: inet_pton(AF_INET, \"1.2.-4294967295.3\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:123: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:123: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:123: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x001.2. 3.4\x00src/functional/inet_pton.c:124: inet_addr(\"1.2. 3.4\") returned %s, want %s\n\x00src/functional/inet_pton.c:124: inet_pton(AF_INET, \"1.2. 3.4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:124: inet_pton(AF_INET, \"1.2. 3.4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:124: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:124: inet_ntop(AF_INET, <\"ffffffff\">, buf, size) got %s, want %s\n\x00src/functional/inet_pton.c:124: inet_ntoa(<\"ffffffff\">) returned %s, want %s\n\x00:\x00src/functional/inet_pton.c:127: inet_pton(AF_INET6, \":\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:127: inet_pton(AF_INET6, \":\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:127: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:127: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:127: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x0000000000000000000000ffff\x00src/functional/inet_pton.c:127: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00::\x00src/functional/inet_pton.c:128: inet_pton(AF_INET6, \"::\", addr) returned %d, want %d\n\x0000000000000000000000000000000000\x00src/functional/inet_pton.c:128: inet_pton(AF_INET6, \"::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:128: inet_ntop(AF_INET6, <\"00000000000000000000000000000000\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:128: inet_ntop(AF_INET6, <\"00000000000000000000000000000000\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:128: inet_ntop(AF_INET6, <\"00000000000000000000000000000000\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:128: inet_ntop(AF_INET6, <\"00000000000000000000000000000000\">, buf, size) got %s, should be ipv4 mapped\n\x00::1\x00src/functional/inet_pton.c:129: inet_pton(AF_INET6, \"::1\", addr) returned %d, want %d\n\x0000000000000000000000000000000001\x00src/functional/inet_pton.c:129: inet_pton(AF_INET6, \"::1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:129: inet_ntop(AF_INET6, <\"00000000000000000000000000000001\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:129: inet_ntop(AF_INET6, <\"00000000000000000000000000000001\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:129: inet_ntop(AF_INET6, <\"00000000000000000000000000000001\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:129: inet_ntop(AF_INET6, <\"00000000000000000000000000000001\">, buf, size) got %s, should be ipv4 mapped\n\x00:::\x00src/functional/inet_pton.c:130: inet_pton(AF_INET6, \":::\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:130: inet_pton(AF_INET6, \":::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:130: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:130: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:130: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:130: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00192.168.1.1\x00src/functional/inet_pton.c:131: inet_pton(AF_INET6, \"192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:131: inet_pton(AF_INET6, \"192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:131: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:131: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:131: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:131: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00:192.168.1.1\x00src/functional/inet_pton.c:132: inet_pton(AF_INET6, \":192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:132: inet_pton(AF_INET6, \":192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:132: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:132: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:132: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:132: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00::192.168.1.1\x00src/functional/inet_pton.c:133: inet_pton(AF_INET6, \"::192.168.1.1\", addr) returned %d, want %d\n\x00000000000000000000000000c0a80101\x00src/functional/inet_pton.c:133: inet_pton(AF_INET6, \"::192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:133: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:133: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:133: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:133: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, should be ipv4 mapped\n\x000:0:0:0:0:0:192.168.1.1\x00src/functional/inet_pton.c:134: inet_pton(AF_INET6, \"0:0:0:0:0:0:192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:134: inet_pton(AF_INET6, \"0:0:0:0:0:0:192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:134: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:134: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:134: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:134: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, should be ipv4 mapped\n\x000:0::0:0:0:192.168.1.1\x00src/functional/inet_pton.c:135: inet_pton(AF_INET6, \"0:0::0:0:0:192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:135: inet_pton(AF_INET6, \"0:0::0:0:0:192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:135: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:135: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:135: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:135: inet_ntop(AF_INET6, <\"000000000000000000000000c0a80101\">, buf, size) got %s, should be ipv4 mapped\n\x00::012.34.56.78\x00src/functional/inet_pton.c:136: inet_pton(AF_INET6, \"::012.34.56.78\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:136: inet_pton(AF_INET6, \"::012.34.56.78\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:136: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:136: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:136: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:136: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00:ffff:192.168.1.1\x00src/functional/inet_pton.c:137: inet_pton(AF_INET6, \":ffff:192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:137: inet_pton(AF_INET6, \":ffff:192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:137: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:137: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:137: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:137: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00::ffff:192.168.1.1\x00src/functional/inet_pton.c:138: inet_pton(AF_INET6, \"::ffff:192.168.1.1\", addr) returned %d, want %d\n\x0000000000000000000000ffffc0a80101\x00src/functional/inet_pton.c:138: inet_pton(AF_INET6, \"::ffff:192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:138: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a80101\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:138: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a80101\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:138: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a80101\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:138: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a80101\">, buf, size) got %s, should be ipv4 mapped\n\x00.192.168.1.1\x00src/functional/inet_pton.c:139: inet_pton(AF_INET6, \".192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:139: inet_pton(AF_INET6, \".192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:139: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:139: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:139: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:139: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00:.192.168.1.1\x00src/functional/inet_pton.c:140: inet_pton(AF_INET6, \":.192.168.1.1\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:140: inet_pton(AF_INET6, \":.192.168.1.1\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:140: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:140: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:140: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:140: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00a:0b:00c:000d:E:F::\x00src/functional/inet_pton.c:141: inet_pton(AF_INET6, \"a:0b:00c:000d:E:F::\", addr) returned %d, want %d\n\x00000a000b000c000d000e000f00000000\x00src/functional/inet_pton.c:141: inet_pton(AF_INET6, \"a:0b:00c:000d:E:F::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:141: inet_ntop(AF_INET6, <\"000a000b000c000d000e000f00000000\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:141: inet_ntop(AF_INET6, <\"000a000b000c000d000e000f00000000\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:141: inet_ntop(AF_INET6, <\"000a000b000c000d000e000f00000000\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:141: inet_ntop(AF_INET6, <\"000a000b000c000d000e000f00000000\">, buf, size) got %s, should be ipv4 mapped\n\x00a:0b:00c:000d:0000e:f::\x00src/functional/inet_pton.c:142: inet_pton(AF_INET6, \"a:0b:00c:000d:0000e:f::\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:142: inet_pton(AF_INET6, \"a:0b:00c:000d:0000e:f::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:142: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:142: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:142: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:142: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x001:2:3:4:5:6::\x00src/functional/inet_pton.c:143: inet_pton(AF_INET6, \"1:2:3:4:5:6::\", addr) returned %d, want %d\n\x0000010002000300040005000600000000\x00src/functional/inet_pton.c:143: inet_pton(AF_INET6, \"1:2:3:4:5:6::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:143: inet_ntop(AF_INET6, <\"00010002000300040005000600000000\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:143: inet_ntop(AF_INET6, <\"00010002000300040005000600000000\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:143: inet_ntop(AF_INET6, <\"00010002000300040005000600000000\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:143: inet_ntop(AF_INET6, <\"00010002000300040005000600000000\">, buf, size) got %s, should be ipv4 mapped\n\x001:2:3:4:5:6:7::\x00src/functional/inet_pton.c:144: inet_pton(AF_INET6, \"1:2:3:4:5:6:7::\", addr) returned %d, want %d\n\x0000010002000300040005000600070000\x00src/functional/inet_pton.c:144: inet_pton(AF_INET6, \"1:2:3:4:5:6:7::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:144: inet_ntop(AF_INET6, <\"00010002000300040005000600070000\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:144: inet_ntop(AF_INET6, <\"00010002000300040005000600070000\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:144: inet_ntop(AF_INET6, <\"00010002000300040005000600070000\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:144: inet_ntop(AF_INET6, <\"00010002000300040005000600070000\">, buf, size) got %s, should be ipv4 mapped\n\x001:2:3:4:5:6:7:8::\x00src/functional/inet_pton.c:145: inet_pton(AF_INET6, \"1:2:3:4:5:6:7:8::\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:145: inet_pton(AF_INET6, \"1:2:3:4:5:6:7:8::\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:145: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:145: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:145: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:145: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x001:2:3:4:5:6:7::9\x00src/functional/inet_pton.c:146: inet_pton(AF_INET6, \"1:2:3:4:5:6:7::9\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:146: inet_pton(AF_INET6, \"1:2:3:4:5:6:7::9\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:146: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:146: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:146: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:146: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00::1:2:3:4:5:6\x00src/functional/inet_pton.c:147: inet_pton(AF_INET6, \"::1:2:3:4:5:6\", addr) returned %d, want %d\n\x0000000000000100020003000400050006\x00src/functional/inet_pton.c:147: inet_pton(AF_INET6, \"::1:2:3:4:5:6\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:147: inet_ntop(AF_INET6, <\"00000000000100020003000400050006\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:147: inet_ntop(AF_INET6, <\"00000000000100020003000400050006\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:147: inet_ntop(AF_INET6, <\"00000000000100020003000400050006\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:147: inet_ntop(AF_INET6, <\"00000000000100020003000400050006\">, buf, size) got %s, should be ipv4 mapped\n\x00::1:2:3:4:5:6:7\x00src/functional/inet_pton.c:148: inet_pton(AF_INET6, \"::1:2:3:4:5:6:7\", addr) returned %d, want %d\n\x0000000001000200030004000500060007\x00src/functional/inet_pton.c:148: inet_pton(AF_INET6, \"::1:2:3:4:5:6:7\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:148: inet_ntop(AF_INET6, <\"00000001000200030004000500060007\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:148: inet_ntop(AF_INET6, <\"00000001000200030004000500060007\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:148: inet_ntop(AF_INET6, <\"00000001000200030004000500060007\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:148: inet_ntop(AF_INET6, <\"00000001000200030004000500060007\">, buf, size) got %s, should be ipv4 mapped\n\x00::1:2:3:4:5:6:7:8\x00src/functional/inet_pton.c:149: inet_pton(AF_INET6, \"::1:2:3:4:5:6:7:8\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:149: inet_pton(AF_INET6, \"::1:2:3:4:5:6:7:8\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:149: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:149: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:149: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:149: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00a:b::c:d:e:f\x00src/functional/inet_pton.c:150: inet_pton(AF_INET6, \"a:b::c:d:e:f\", addr) returned %d, want %d\n\x00000a000b00000000000c000d000e000f\x00src/functional/inet_pton.c:150: inet_pton(AF_INET6, \"a:b::c:d:e:f\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:150: inet_ntop(AF_INET6, <\"000a000b00000000000c000d000e000f\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:150: inet_ntop(AF_INET6, <\"000a000b00000000000c000d000e000f\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:150: inet_ntop(AF_INET6, <\"000a000b00000000000c000d000e000f\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:150: inet_ntop(AF_INET6, <\"000a000b00000000000c000d000e000f\">, buf, size) got %s, should be ipv4 mapped\n\x00ffff:c0a8:5e4\x00src/functional/inet_pton.c:151: inet_pton(AF_INET6, \"ffff:c0a8:5e4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:151: inet_pton(AF_INET6, \"ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:151: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:151: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:151: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:151: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00:ffff:c0a8:5e4\x00src/functional/inet_pton.c:152: inet_pton(AF_INET6, \":ffff:c0a8:5e4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:152: inet_pton(AF_INET6, \":ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:152: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:152: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:152: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:152: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x000:0:0:0:0:ffff:c0a8:5e4\x00src/functional/inet_pton.c:153: inet_pton(AF_INET6, \"0:0:0:0:0:ffff:c0a8:5e4\", addr) returned %d, want %d\n\x0000000000000000000000ffffc0a805e4\x00src/functional/inet_pton.c:153: inet_pton(AF_INET6, \"0:0:0:0:0:ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:153: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:153: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:153: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:153: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s, should be ipv4 mapped\n\x000:0:0:0:ffff:c0a8:5e4\x00src/functional/inet_pton.c:154: inet_pton(AF_INET6, \"0:0:0:0:ffff:c0a8:5e4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:154: inet_pton(AF_INET6, \"0:0:0:0:ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:154: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:154: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:154: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:154: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x000::ffff:c0a8:5e4\x00src/functional/inet_pton.c:155: inet_pton(AF_INET6, \"0::ffff:c0a8:5e4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:155: inet_pton(AF_INET6, \"0::ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:155: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:155: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:155: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:155: inet_ntop(AF_INET6, <\"00000000000000000000ffffc0a805e4\">, buf, size) got %s, should be ipv4 mapped\n\x00::0::ffff:c0a8:5e4\x00src/functional/inet_pton.c:156: inet_pton(AF_INET6, \"::0::ffff:c0a8:5e4\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:156: inet_pton(AF_INET6, \"::0::ffff:c0a8:5e4\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:156: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:156: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:156: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:156: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00c0a8\x00src/functional/inet_pton.c:157: inet_pton(AF_INET6, \"c0a8\", addr) returned %d, want %d\n\x00src/functional/inet_pton.c:157: inet_pton(AF_INET6, \"c0a8\", addr) got addr %s, want %s\n\x00src/functional/inet_pton.c:157: inet_ntop(AF_INET6, <\"\">, buf, size) did not return buf\n\x00src/functional/inet_pton.c:157: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, it is rejected by inet_pton\n\x00src/functional/inet_pton.c:157: inet_ntop(AF_INET6, <\"\">, buf, size) got %s that is %s, want %s\n\x00src/functional/inet_pton.c:157: inet_ntop(AF_INET6, <\"\">, buf, size) got %s, should be ipv4 mapped\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
