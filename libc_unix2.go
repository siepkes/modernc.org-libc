// Copyright 2024 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build illumos

package libc // import "modernc.org/libc"

import (
	"unsafe"

	"modernc.org/libc/errno"
	"modernc.org/libc/sys/types"
	ctime "modernc.org/libc/time"
)

// ssize_t recvmsg(int sockfd, struct msghdr *msg, int flags);
func Xrecvmsg(t *TLS, sockfd int32, msg uintptr, flags int32) types.Ssize_t {
	if __ccgo_strace {
		trc("t=%v sockfd=%v msg=%v flags=%v, (%v:)", t, sockfd, msg, flags, origin(2))
	}
	panic(todo(""))
	//	n, _, err := unix.Syscall(unix.SYS_RECVMSG, uintptr(sockfd), msg, uintptr(flags))
	//	if err != 0 {
	//		t.setErrno(err)
	//		return -1
	//	}
	//
	// return types.Ssize_t(n)
}

func Xgmtime_r(tls *TLS, t uintptr, tm uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("tls=%v t=%v tm=%v, (%v:)", tls, t, tm, origin(2))
		defer func() { trc("-> %v", r) }()
	}
	if x___secs_to_tm(tls, int64(*(*time_t)(unsafe.Pointer(t))), tm) < 0 {
		*(*int32)(unsafe.Pointer(X__errno_location(tls))) = int32(errno.EOVERFLOW)
		return uintptr(0)
	}
	(*ctime.Tm)(unsafe.Pointer(tm)).Ftm_isdst = 0
	return tm
}
