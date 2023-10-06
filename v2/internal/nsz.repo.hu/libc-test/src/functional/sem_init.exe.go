// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/functional/sem_init.exe.go src/functional/sem_init.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/time.h:87:9:
const CLOCK_REALTIME = 0

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:11:9:
const EAGAIN = 11

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:111:9:
const ETIMEDOUT = 110

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:108:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:108:16:
type time_t = int64

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:237:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:237:13:
type clockid_t = int32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:252:1:
type timespec = struct {
	Ftv_sec   time_t
	Ftv_nsec  int32
	F__ccgo12 uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:296:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:296:26:
type pthread_t = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:302:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:302:13:
type pthread_once_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:307:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:307:18:
type pthread_key_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:312:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:312:13:
type pthread_spinlock_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:317:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:317:37:
type pthread_mutexattr_t = struct {
	F__attr uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:322:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:322:37:
type pthread_condattr_t = struct {
	F__attr uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:327:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:327:37:
type pthread_barrierattr_t = struct {
	F__attr uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:332:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:332:40:
type pthread_rwlockattr_t = struct {
	F__attr [2]uint32
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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:401:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:401:147:
type pthread_attr_t = struct {
	F__u struct {
		F__vi [0][9]int32
		F__s  [0][9]uint32
		F__i  [9]int32
	}
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:406:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:406:157:
type pthread_mutex_t = struct {
	F__u struct {
		F__vi [0][6]int32
		F__p  [0][6]uintptr
		F__i  [6]int32
	}
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:416:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:416:112:
type pthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][12]uintptr
		F__i  [12]int32
	}
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:426:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:426:139:
type pthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][8]uintptr
		F__i  [8]int32
	}
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:431:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:431:137:
type pthread_barrier_t = struct {
	F__u struct {
		F__vi [0][5]int32
		F__p  [0][5]uintptr
		F__i  [5]int32
	}
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:258:13:
type pid_t = int32

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sched.h:19:1:
type sched_param = struct {
	Fsched_priority int32
	F__reserved1    int32
	F__reserved2    [4]int32
	F__reserved3    int32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sched.h:87:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/sched.h:87:70:
type cpu_set_t = struct {
	F__bits [32]uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:232:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:232:14:
type timer_t = uintptr

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:242:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:242:14:
type clock_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:366:32:
type locale_t = uintptr

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/time.h:40:1:
type tm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	Ftm_gmtoff int32
	Ftm_zone   uintptr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/time.h:82:1:
type itimerspec = struct {
	Fit_interval timespec
	Fit_value    timespec
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/pthread.h:206:1:
type __ptcb = struct {
	F__f    uintptr
	F__x    uintptr
	F__next uintptr
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/pthread.h:219:1:
type cpu_set_t1 = struct {
	F__bits [32]uint32
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:88:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:88:15:
type ssize_t = int32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:175:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:175:18:
type mode_t = uint32

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:185:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:185:16:
type off_t = int64

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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/semaphore.h:17:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/semaphore.h:19:3:
type sem_t = struct {
	F__val [4]int32
}

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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:12:13:
func start(tls *libc.TLS, arg uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:13:1:
	var s uintptr
	var _ /* ts at bp+0 */ timespec
	_ = s
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:15:8:
	s = arg
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:16:2:
	if libc.Xclock_gettime(tls, CLOCK_REALTIME, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:16:2:
		t_printf(tls, __ccgo_ts, libc.VaList(bp+24, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:17:2:
	(*(*timespec)(unsafe.Pointer(bp))).Ftv_sec += int64(1)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:18:2:
	if libc.Xsem_post(tls, s) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:18:2:
		t_printf(tls, __ccgo_ts+77, libc.VaList(bp+24, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:19:2:
	if libc.Xsem_timedwait(tls, s+uintptr(1)*16, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:19:2:
		t_printf(tls, __ccgo_ts+131, libc.VaList(bp+24, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:20:2:
	return uintptr(0)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:23:13:
func many_waiters(tls *libc.TLS) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:24:1:
	var v1, v2, v3, v4, v5, v6 int32
	var _ /* p at bp+48 */ uintptr
	var _ /* r at bp+44 */ int32
	var _ /* s at bp+12 */ [2]sem_t
	var _ /* t at bp+0 */ [3]pthread_t
	_, _, _, _, _, _ = v1, v2, v3, v4, v5, v6
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:30:2:
	if libc.Xsem_init(tls, bp+12, 0, uint32(0)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:30:2:
		t_printf(tls, __ccgo_ts+197, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:31:2:
	if libc.Xsem_init(tls, bp+12+uintptr(1)*16, 0, uint32(0)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:31:2:
		t_printf(tls, __ccgo_ts+257, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:32:2:
	v1 = pthread_create(tls, bp, uintptr(0), __ccgo_fp(start), bp+12)
	*(*int32)(unsafe.Pointer(bp + 44)) = v1
	if v1 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:32:2:
		t_printf(tls, __ccgo_ts+319, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:33:2:
	v2 = pthread_create(tls, bp+uintptr(1)*4, uintptr(0), __ccgo_fp(start), bp+12)
	*(*int32)(unsafe.Pointer(bp + 44)) = v2
	if v2 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:33:2:
		t_printf(tls, __ccgo_ts+392, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:34:2:
	v3 = pthread_create(tls, bp+uintptr(2)*4, uintptr(0), __ccgo_fp(start), bp+12)
	*(*int32)(unsafe.Pointer(bp + 44)) = v3
	if v3 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:34:2:
		t_printf(tls, __ccgo_ts+467, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:35:2:
	if libc.Xsem_wait(tls, bp+12) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:35:2:
		t_printf(tls, __ccgo_ts+542, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:36:2:
	if libc.Xsem_wait(tls, bp+12) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:36:2:
		t_printf(tls, __ccgo_ts+596, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:37:2:
	if libc.Xsem_wait(tls, bp+12) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:37:2:
		t_printf(tls, __ccgo_ts+650, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:38:2:
	if libc.Xsem_getvalue(tls, bp+12, bp+44) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:38:2:
		t_printf(tls, __ccgo_ts+704, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:39:2:
	if *(*int32)(unsafe.Pointer(bp + 44)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:40:3:
		t_printf(tls, __ccgo_ts+766, libc.VaList(bp+64, *(*int32)(unsafe.Pointer(bp + 44))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:41:2:
	if libc.Xsem_post(tls, bp+12+uintptr(1)*16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:41:2:
		t_printf(tls, __ccgo_ts+827, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:42:2:
	if libc.Xsem_post(tls, bp+12+uintptr(1)*16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:42:2:
		t_printf(tls, __ccgo_ts+883, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:43:2:
	if libc.Xsem_post(tls, bp+12+uintptr(1)*16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:43:2:
		t_printf(tls, __ccgo_ts+939, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:44:2:
	v4 = pthread_join(tls, (*(*[3]pthread_t)(unsafe.Pointer(bp)))[0], bp+48)
	*(*int32)(unsafe.Pointer(bp + 44)) = v4
	if v4 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:44:2:
		t_printf(tls, __ccgo_ts+995, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:45:2:
	v5 = pthread_join(tls, (*(*[3]pthread_t)(unsafe.Pointer(bp)))[int32(1)], bp+48)
	*(*int32)(unsafe.Pointer(bp + 44)) = v5
	if v5 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:45:2:
		t_printf(tls, __ccgo_ts+1059, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:46:2:
	v6 = pthread_join(tls, (*(*[3]pthread_t)(unsafe.Pointer(bp)))[int32(2)], bp+48)
	*(*int32)(unsafe.Pointer(bp + 44)) = v6
	if v6 != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:46:2:
		t_printf(tls, __ccgo_ts+1123, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(bp + 44)))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:47:2:
	if libc.Xsem_getvalue(tls, bp+12+uintptr(1)*16, bp+44) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:47:2:
		t_printf(tls, __ccgo_ts+1187, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:48:2:
	if *(*int32)(unsafe.Pointer(bp + 44)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:49:3:
		t_printf(tls, __ccgo_ts+1251, libc.VaList(bp+64, *(*int32)(unsafe.Pointer(bp + 44))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:50:2:
	if libc.Xsem_destroy(tls, bp+12) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:50:2:
		t_printf(tls, __ccgo_ts+1312, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:51:2:
	if libc.Xsem_destroy(tls, bp+12+uintptr(1)*16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:51:2:
		t_printf(tls, __ccgo_ts+1369, libc.VaList(bp+64, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:54:13:
func single_thread(tls *libc.TLS) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:55:1:
	var _ /* r at bp+32 */ int32
	var _ /* s at bp+16 */ sem_t
	var _ /* ts at bp+0 */ timespec
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:60:2:
	if libc.Xsem_init(tls, bp+16, 0, uint32(1)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:60:2:
		t_printf(tls, __ccgo_ts+1428, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:61:2:
	if libc.Xsem_wait(tls, bp+16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:61:2:
		t_printf(tls, __ccgo_ts+1489, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:62:2:
	if libc.Xsem_getvalue(tls, bp+16, bp+32) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:62:2:
		t_printf(tls, __ccgo_ts+1544, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:63:2:
	if *(*int32)(unsafe.Pointer(bp + 32)) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:64:3:
		t_printf(tls, __ccgo_ts+1607, libc.VaList(bp+48, *(*int32)(unsafe.Pointer(bp + 32))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:65:2:
	if libc.Xsem_trywait(tls, bp+16) != -int32(1) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(EAGAIN) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:66:3:
		t_printf(tls, __ccgo_ts+1668, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:67:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:68:2:
	if libc.Xclock_gettime(tls, CLOCK_REALTIME, bp) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:68:2:
		t_printf(tls, __ccgo_ts+1743, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:69:2:
	if libc.Xsem_timedwait(tls, bp+16, bp) != -int32(1) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(ETIMEDOUT) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:70:3:
		t_printf(tls, __ccgo_ts+1820, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:71:2:
	if libc.Xsem_destroy(tls, bp+16) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:71:2:
		t_printf(tls, __ccgo_ts+1900, libc.VaList(bp+48, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
	}
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:74:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:75:1:
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:76:2:
	single_thread(tls)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:77:2:
	many_waiters(tls)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/sem_init.c:78:2:
	return libc.AtomicLoadPInt32(uintptr(unsafe.Pointer(&t_status)))
}

func main() {
	libc.Start(main1)
}

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:24:9:
const EMFILE = 24

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/fcntl.h:44:9:
const O_RDONLY = 0

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
		fd = libc.Xopen(tls, __ccgo_ts+1958, O_RDONLY, 0)
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
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:113:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:113:16:
type suseconds_t = int64

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:247:1:
type timeval = struct {
	Ftv_sec  time_t
	Ftv_usec suseconds_t
}

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
		t_printf(tls, __ccgo_ts+1968, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 2012,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 2020,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 2028,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 2038,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 2048,
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
				v2 = __ccgo_ts + 2057
			} else {
				v2 = __ccgo_ts + 2059
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+2060, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 2057
		} else {
			v3 = __ccgo_ts + 2059
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+2065, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts + 2059
	} else {
		v4 = __ccgo_ts + 2070
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+2072, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 2075
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 2078
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 2081
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 2084
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 2087
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+2090, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+2098, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+2103, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+2146, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2195) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2203) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2215) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2227) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2239) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2248) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+2059) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+2248) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+2254, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:6:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/print.c:6:14:
var t_status int32

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "src/functional/sem_init.c:16: clock_gettime(CLOCK_REALTIME, &ts) failed: %s\n\x00src/functional/sem_init.c:18: sem_post(s) failed: %s\n\x00src/functional/sem_init.c:19: sem_timedwait(s+1, &ts) failed: %s\n\x00src/functional/sem_init.c:30: sem_init(s, 0, 0) failed: %s\n\x00src/functional/sem_init.c:31: sem_init(s+1, 0, 0) failed: %s\n\x00src/functional/sem_init.c:32: pthread_create(t, 0, start, s) failed: %s\n\x00src/functional/sem_init.c:33: pthread_create(t+1, 0, start, s) failed: %s\n\x00src/functional/sem_init.c:34: pthread_create(t+2, 0, start, s) failed: %s\n\x00src/functional/sem_init.c:35: sem_wait(s) failed: %s\n\x00src/functional/sem_init.c:36: sem_wait(s) failed: %s\n\x00src/functional/sem_init.c:37: sem_wait(s) failed: %s\n\x00src/functional/sem_init.c:38: sem_getvalue(s, &r) failed: %s\n\x00src/functional/sem_init.c:40: sem value should be 0, got %d\n\x00src/functional/sem_init.c:41: sem_post(s+1) failed: %s\n\x00src/functional/sem_init.c:42: sem_post(s+1) failed: %s\n\x00src/functional/sem_init.c:43: sem_post(s+1) failed: %s\n\x00src/functional/sem_init.c:44: pthread_join(t[0],&p) failed: %s\n\x00src/functional/sem_init.c:45: pthread_join(t[1],&p) failed: %s\n\x00src/functional/sem_init.c:46: pthread_join(t[2],&p) failed: %s\n\x00src/functional/sem_init.c:47: sem_getvalue(s+1, &r) failed: %s\n\x00src/functional/sem_init.c:49: sem value should be 0, got %d\n\x00src/functional/sem_init.c:50: sem_destroy(s) failed: %s\n\x00src/functional/sem_init.c:51: sem_destroy(s+1) failed: %s\n\x00src/functional/sem_init.c:60: sem_init(&s, 0, 1) failed: %s\n\x00src/functional/sem_init.c:61: sem_wait(&s) failed: %s\n\x00src/functional/sem_init.c:62: sem_getvalue(&s, &r) failed: %s\n\x00src/functional/sem_init.c:64: sem value should be 0, got %d\n\x00src/functional/sem_init.c:66: sem_trywait should fail with EAGAIN, got %s\n\x00src/functional/sem_init.c:68: clock_gettime(CLOCK_REALTIME, &ts) failed: %s\n\x00src/functional/sem_init.c:70: sem_timedwait should fail with ETIMEDOUT, got %s\n\x00src/functional/sem_init.c:71: sem_destroy(&s) failed: %s\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
