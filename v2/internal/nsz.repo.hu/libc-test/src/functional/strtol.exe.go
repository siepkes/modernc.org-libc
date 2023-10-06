// Code generated for linux/386 by 'cc --prefix-field=F -absolute-paths -extended-errors -positions -o src/functional/strtol.exe.go src/functional/strtol.o.go src/common/libtest.a -lpthread -lm -lrt -lcrypt -ldl -lresolv -lutil -lpthread', DO NOT EDIT.

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

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:22:9:
const EINVAL = 22

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/errno.h:34:9:
const ERANGE = 34

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

//
// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:1:

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/include/linux/386/bits/alltypes.h:283:18:
type useconds_t = uint32

/* r = place to store result
 * f = function call to test (or any expression)
 * x = expected result
 * m = message to print on failure (with formats for r & x)
**/

// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:21:5:
func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:22:1:
	var i, l, v1, v10, v101, v103, v108, v110, v115, v117, v12, v122, v124, v129, v131, v136, v138, v14, v143, v145, v150, v152, v157, v159, v164, v166, v17, v171, v173, v175, v177, v179, v182, v185, v187, v19, v190, v192, v195, v197, v199, v202, v204, v207, v24, v26, v3, v31, v33, v38, v40, v45, v47, v52, v54, v59, v61, v63, v66, v68, v7, v70, v73, v75, v80, v82, v87, v89, v94, v96 int32
	var ll, v119, v126 int64
	var msg, s, v106, v113, v120, v127, v134, v141, v148, v15, v155, v162, v169, v180, v183, v188, v193, v200, v205, v22, v29, v36, v43, v50, v57, v64, v71, v78, v8, v85, v92, v99 uintptr
	var ul, v105, v112, v21, v28, v35, v42, v49, v5, v56, v77, v84, v91, v98 uint32
	var ull, v133, v140, v147, v154, v161, v168 uint64
	var v100, v102, v104, v107, v109, v11, v111, v114, v116, v118, v121, v123, v125, v128, v13, v130, v132, v135, v137, v139, v142, v144, v146, v149, v151, v153, v156, v158, v16, v160, v163, v165, v167, v170, v172, v174, v176, v178, v18, v181, v184, v186, v189, v191, v194, v196, v198, v2, v20, v201, v203, v206, v208, v23, v25, v27, v30, v32, v34, v37, v39, v4, v41, v44, v46, v48, v51, v53, v55, v58, v6, v60, v62, v65, v67, v69, v72, v74, v76, v79, v81, v83, v86, v88, v9, v90, v93, v95, v97 bool
	var _ /* c at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, l, ll, msg, s, ul, ull, v1, v10, v100, v101, v102, v103, v104, v105, v106, v107, v108, v109, v11, v110, v111, v112, v113, v114, v115, v116, v117, v118, v119, v12, v120, v121, v122, v123, v124, v125, v126, v127, v128, v129, v13, v130, v131, v132, v133, v134, v135, v136, v137, v138, v139, v14, v140, v141, v142, v143, v144, v145, v146, v147, v148, v149, v15, v150, v151, v152, v153, v154, v155, v156, v157, v158, v159, v16, v160, v161, v162, v163, v164, v165, v166, v167, v168, v169, v17, v170, v171, v172, v173, v174, v175, v176, v177, v178, v179, v18, v180, v181, v182, v183, v184, v185, v186, v187, v188, v189, v19, v190, v191, v192, v193, v194, v195, v196, v197, v198, v199, v2, v20, v200, v201, v202, v203, v204, v205, v206, v207, v208, v21, v22, v23, v24, v25, v26, v27, v28, v29, v3, v30, v31, v32, v33, v34, v35, v36, v37, v38, v39, v4, v40, v41, v42, v43, v44, v45, v46, v47, v48, v49, v5, v50, v51, v52, v53, v54, v55, v56, v57, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v67, v68, v69, v7, v70, v71, v72, v73, v74, v75, v76, v77, v78, v79, v8, v80, v81, v82, v83, v84, v85, v86, v87, v88, v89, v9, v90, v91, v92, v93, v94, v95, v96, v97, v98, v99
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:28:7:
	msg = __ccgo_ts
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:31:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 1
	v1 = libc.Xatol(tls, __ccgo_ts+20)
	l = v1
	if v2 = v1 == int32(2147483647); !v2 {
		t_printf(tls, __ccgo_ts+31, libc.VaList(bp+16, __ccgo_ts+1, l, int32(2147483647)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:32:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 100
	v3 = libc.Xstrtol(tls, __ccgo_ts+20, uintptr(0), 0)
	l = v3
	if v4 = v3 == int32(2147483647); !v4 {
		t_printf(tls, __ccgo_ts+127, libc.VaList(bp+16, __ccgo_ts+100, l, int32(2147483647)))
	}
	_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:33:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 196
	v5 = libc.Xstrtoul(tls, __ccgo_ts+224, uintptr(0), 0)
	ul = v5
	if v6 = v5 == uint32(4294967295); !v6 {
		t_printf(tls, __ccgo_ts+235, libc.VaList(bp+16, __ccgo_ts+196, ul, uint32(4294967295)))
	}
	_ = v6 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:35:2:
	if uint32(4) == uint32(4) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:36:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 306
		v8 = __ccgo_ts + 336
		s = v8
		v7 = libc.Xstrtol(tls, v8, bp, 0)
		l = v7
		if v9 = v7 == int32(2147483647); !v9 {
			t_printf(tls, __ccgo_ts+347, libc.VaList(bp+16, __ccgo_ts+306, l, int32(2147483647)))
		}
		_ = v9 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:37:3:
		v10 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v10
		if v11 = v10 == int32(10); !v11 {
			t_printf(tls, __ccgo_ts+417, libc.VaList(bp+16, msg, i, int32(10)))
		}
		_ = v11 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:38:3:
		v12 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v12
		if v13 = v12 == int32(ERANGE); !v13 {
			t_printf(tls, __ccgo_ts+488, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v13 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:39:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 552
		v15 = __ccgo_ts + 583
		s = v15
		v14 = libc.Xstrtol(tls, v15, bp, 0)
		l = v14
		if v16 = v14 == -libc.Int32FromInt32(2147483647)-libc.Int32FromInt32(1); !v16 {
			t_printf(tls, __ccgo_ts+595, libc.VaList(bp+16, __ccgo_ts+552, l, -libc.Int32FromInt32(2147483647)-libc.Int32FromInt32(1)))
		}
		_ = v16 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:40:3:
		v17 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v17
		if v18 = v17 == int32(11); !v18 {
			t_printf(tls, __ccgo_ts+665, libc.VaList(bp+16, msg, i, int32(11)))
		}
		_ = v18 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:41:3:
		v19 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v19
		if v20 = v19 == int32(ERANGE); !v20 {
			t_printf(tls, __ccgo_ts+736, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v20 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:42:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 800
		v22 = __ccgo_ts + 831
		s = v22
		v21 = libc.Xstrtoul(tls, v22, bp, 0)
		ul = v21
		if v23 = v21 == uint32(4294967295); !v23 {
			t_printf(tls, __ccgo_ts+842, libc.VaList(bp+16, __ccgo_ts+800, ul, uint32(4294967295)))
		}
		_ = v23 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:43:3:
		v24 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v24
		if v25 = v24 == int32(10); !v25 {
			t_printf(tls, __ccgo_ts+912, libc.VaList(bp+16, msg, i, int32(10)))
		}
		_ = v25 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:44:3:
		v26 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v26
		if v27 = v26 == int32(ERANGE); !v27 {
			t_printf(tls, __ccgo_ts+983, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v27 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:45:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 1047
		v29 = __ccgo_ts + 1070
		s = v29
		v28 = libc.Xstrtoul(tls, v29, bp, 0)
		ul = v28
		if v30 = v28 == -libc.Uint32FromUint32(1); !v30 {
			t_printf(tls, __ccgo_ts+1073, libc.VaList(bp+16, __ccgo_ts+1047, ul, -libc.Uint32FromUint32(1)))
		}
		_ = v30 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:46:3:
		v31 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v31
		if v32 = v31 == int32(2); !v32 {
			t_printf(tls, __ccgo_ts+1143, libc.VaList(bp+16, msg, i, int32(2)))
		}
		_ = v32 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:47:3:
		v33 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v33
		if v34 = v33 == 0; !v34 {
			t_printf(tls, __ccgo_ts+1214, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v34 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:48:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 1279
		v36 = __ccgo_ts + 1302
		s = v36
		v35 = libc.Xstrtoul(tls, v36, bp, 0)
		ul = v35
		if v37 = v35 == -libc.Uint32FromUint32(2); !v37 {
			t_printf(tls, __ccgo_ts+1305, libc.VaList(bp+16, __ccgo_ts+1279, ul, -libc.Uint32FromUint32(2)))
		}
		_ = v37 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:49:3:
		v38 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v38
		if v39 = v38 == int32(2); !v39 {
			t_printf(tls, __ccgo_ts+1375, libc.VaList(bp+16, msg, i, int32(2)))
		}
		_ = v39 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:50:3:
		v40 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v40
		if v41 = v40 == 0; !v41 {
			t_printf(tls, __ccgo_ts+1446, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v41 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:51:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 1511
		v43 = __ccgo_ts + 1543
		s = v43
		v42 = libc.Xstrtoul(tls, v43, bp, 0)
		ul = v42
		if v44 = v42 == -libc.Uint32FromUint32(2147483648); !v44 {
			t_printf(tls, __ccgo_ts+1555, libc.VaList(bp+16, __ccgo_ts+1511, ul, -libc.Uint32FromUint32(2147483648)))
		}
		_ = v44 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:52:3:
		v45 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v45
		if v46 = v45 == int32(11); !v46 {
			t_printf(tls, __ccgo_ts+1625, libc.VaList(bp+16, msg, i, int32(11)))
		}
		_ = v46 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:53:3:
		v47 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v47
		if v48 = v47 == 0; !v48 {
			t_printf(tls, __ccgo_ts+1696, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v48 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:54:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 1761
		v50 = __ccgo_ts + 583
		s = v50
		v49 = libc.Xstrtoul(tls, v50, bp, 0)
		ul = v49
		if v51 = v49 == -libc.Uint32FromUint32(2147483649); !v51 {
			t_printf(tls, __ccgo_ts+1793, libc.VaList(bp+16, __ccgo_ts+1761, ul, -libc.Uint32FromUint32(2147483649)))
		}
		_ = v51 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:55:3:
		v52 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v52
		if v53 = v52 == int32(11); !v53 {
			t_printf(tls, __ccgo_ts+1863, libc.VaList(bp+16, msg, i, int32(11)))
		}
		_ = v53 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:56:3:
		v54 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v54
		if v55 = v54 == 0; !v55 {
			t_printf(tls, __ccgo_ts+1934, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v55 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:57:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 1999
		v57 = __ccgo_ts + 2031
		s = v57
		v56 = libc.Xstrtoul(tls, v57, bp, 0)
		ul = v56
		if v58 = v56 == uint32(4294967295); !v58 {
			t_printf(tls, __ccgo_ts+2043, libc.VaList(bp+16, __ccgo_ts+1999, ul, uint32(4294967295)))
		}
		_ = v58 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:58:3:
		v59 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v59
		if v60 = v59 == int32(11); !v60 {
			t_printf(tls, __ccgo_ts+2122, libc.VaList(bp+16, msg, i, int32(11)))
		}
		_ = v60 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:59:3:
		v61 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v61
		if v62 = v61 == int32(ERANGE); !v62 {
			t_printf(tls, __ccgo_ts+2193, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v62 || libc.Bool(libc.Int32FromInt32(0) != 0)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:60:9:
		if uint32(4) == uint32(8) {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:61:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 2258
			v64 = __ccgo_ts + 2297
			s = v64
			v63 = libc.Xstrtol(tls, v64, bp, 0)
			l = v63
			if v65 = int64(v63) == int64(9223372036854775807); !v65 {
				t_printf(tls, __ccgo_ts+2317, libc.VaList(bp+16, __ccgo_ts+2258, l, int64(9223372036854775807)))
			}
			_ = v65 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:62:3:
			v66 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v66
			if v67 = v66 == int32(19); !v67 {
				t_printf(tls, __ccgo_ts+2387, libc.VaList(bp+16, msg, i, int32(19)))
			}
			_ = v67 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:63:3:
			v68 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v68
			if v69 = v68 == int32(ERANGE); !v69 {
				t_printf(tls, __ccgo_ts+2458, libc.VaList(bp+16, msg, i, int32(ERANGE)))
			}
			_ = v69 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:64:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 2522
			v71 = __ccgo_ts + 2562
			s = v71
			v70 = libc.Xstrtol(tls, v71, bp, 0)
			l = v70
			if v72 = int64(v70) == -libc.Int64FromInt64(9223372036854775807)-libc.Int64FromInt32(1); !v72 {
				t_printf(tls, __ccgo_ts+2583, libc.VaList(bp+16, __ccgo_ts+2522, l, -libc.Int64FromInt64(9223372036854775807)-libc.Int64FromInt32(1)))
			}
			_ = v72 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:65:3:
			v73 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v73
			if v74 = v73 == int32(20); !v74 {
				t_printf(tls, __ccgo_ts+2653, libc.VaList(bp+16, msg, i, int32(20)))
			}
			_ = v74 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:66:3:
			v75 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v75
			if v76 = v75 == int32(ERANGE); !v76 {
				t_printf(tls, __ccgo_ts+2724, libc.VaList(bp+16, msg, i, int32(ERANGE)))
			}
			_ = v76 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:67:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 2788
			v78 = __ccgo_ts + 2829
			s = v78
			v77 = libc.Xstrtoul(tls, v78, bp, 0)
			ul = v77
			if v79 = uint64(v77) == uint64(18446744073709551615); !v79 {
				t_printf(tls, __ccgo_ts+2850, libc.VaList(bp+16, __ccgo_ts+2788, ul, uint64(18446744073709551615)))
			}
			_ = v79 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:68:3:
			v80 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v80
			if v81 = v80 == int32(20); !v81 {
				t_printf(tls, __ccgo_ts+2920, libc.VaList(bp+16, msg, i, int32(20)))
			}
			_ = v81 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:69:3:
			v82 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v82
			if v83 = v82 == int32(ERANGE); !v83 {
				t_printf(tls, __ccgo_ts+2991, libc.VaList(bp+16, msg, i, int32(ERANGE)))
			}
			_ = v83 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:70:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 1047
			v85 = __ccgo_ts + 1070
			s = v85
			v84 = libc.Xstrtoul(tls, v85, bp, 0)
			ul = v84
			if v86 = v84 == -libc.Uint32FromUint32(1); !v86 {
				t_printf(tls, __ccgo_ts+3055, libc.VaList(bp+16, __ccgo_ts+1047, ul, -libc.Uint32FromUint32(1)))
			}
			_ = v86 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:71:3:
			v87 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v87
			if v88 = v87 == int32(2); !v88 {
				t_printf(tls, __ccgo_ts+3125, libc.VaList(bp+16, msg, i, int32(2)))
			}
			_ = v88 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:72:3:
			v89 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v89
			if v90 = v89 == 0; !v90 {
				t_printf(tls, __ccgo_ts+3196, libc.VaList(bp+16, msg, i, 0))
			}
			_ = v90 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:73:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 1279
			v92 = __ccgo_ts + 1302
			s = v92
			v91 = libc.Xstrtoul(tls, v92, bp, 0)
			ul = v91
			if v93 = v91 == -libc.Uint32FromUint32(2); !v93 {
				t_printf(tls, __ccgo_ts+3261, libc.VaList(bp+16, __ccgo_ts+1279, ul, -libc.Uint32FromUint32(2)))
			}
			_ = v93 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:74:3:
			v94 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v94
			if v95 = v94 == int32(2); !v95 {
				t_printf(tls, __ccgo_ts+3331, libc.VaList(bp+16, msg, i, int32(2)))
			}
			_ = v95 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:75:3:
			v96 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v96
			if v97 = v96 == 0; !v97 {
				t_printf(tls, __ccgo_ts+3402, libc.VaList(bp+16, msg, i, 0))
			}
			_ = v97 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:76:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 3467
			v99 = __ccgo_ts + 3508
			s = v99
			v98 = libc.Xstrtoul(tls, v99, bp, 0)
			ul = v98
			if v100 = uint64(v98) == -libc.Uint64FromUint64(9223372036854775808); !v100 {
				t_printf(tls, __ccgo_ts+3529, libc.VaList(bp+16, __ccgo_ts+3467, ul, -libc.Uint64FromUint64(9223372036854775808)))
			}
			_ = v100 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:77:3:
			v101 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v101
			if v102 = v101 == int32(20); !v102 {
				t_printf(tls, __ccgo_ts+3599, libc.VaList(bp+16, msg, i, int32(20)))
			}
			_ = v102 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:78:3:
			v103 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v103
			if v104 = v103 == 0; !v104 {
				t_printf(tls, __ccgo_ts+3670, libc.VaList(bp+16, msg, i, 0))
			}
			_ = v104 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:79:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 3735
			v106 = __ccgo_ts + 2562
			s = v106
			v105 = libc.Xstrtoul(tls, v106, bp, 0)
			ul = v105
			if v107 = uint64(v105) == -libc.Uint64FromUint64(9223372036854775809); !v107 {
				t_printf(tls, __ccgo_ts+3776, libc.VaList(bp+16, __ccgo_ts+3735, ul, -libc.Uint64FromUint64(9223372036854775809)))
			}
			_ = v107 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:80:3:
			v108 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v108
			if v109 = v108 == int32(20); !v109 {
				t_printf(tls, __ccgo_ts+3846, libc.VaList(bp+16, msg, i, int32(20)))
			}
			_ = v109 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:81:3:
			v110 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v110
			if v111 = v110 == 0; !v111 {
				t_printf(tls, __ccgo_ts+3917, libc.VaList(bp+16, msg, i, 0))
			}
			_ = v111 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:82:3:
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			msg = __ccgo_ts + 3982
			v113 = __ccgo_ts + 4024
			s = v113
			v112 = libc.Xstrtoul(tls, v113, bp, 0)
			ul = v112
			if v114 = uint64(v112) == uint64(18446744073709551615); !v114 {
				t_printf(tls, __ccgo_ts+4046, libc.VaList(bp+16, __ccgo_ts+3982, ul, uint64(18446744073709551615)))
			}
			_ = v114 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:83:3:
			v115 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
			i = v115
			if v116 = v115 == int32(21); !v116 {
				t_printf(tls, __ccgo_ts+4125, libc.VaList(bp+16, msg, i, int32(21)))
			}
			_ = v116 || libc.Bool(libc.Int32FromInt32(0) != 0)
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:84:3:
			v117 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
			i = v117
			if v118 = v117 == int32(ERANGE); !v118 {
				t_printf(tls, __ccgo_ts+4196, libc.VaList(bp+16, msg, i, int32(ERANGE)))
			}
			_ = v118 || libc.Bool(libc.Int32FromInt32(0) != 0)
		} else {
			// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:86:3:
			t_printf(tls, __ccgo_ts+4261, libc.VaList(bp+16, libc.Int32FromInt64(4)))
		}
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:89:2:
	if uint32(8) == uint32(8) {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:90:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 4326
		v120 = __ccgo_ts + 2297
		s = v120
		v119 = libc.Xstrtoll(tls, v120, bp, 0)
		ll = v119
		if v121 = v119 == int64(9223372036854775807); !v121 {
			t_printf(tls, __ccgo_ts+4366, libc.VaList(bp+16, __ccgo_ts+4326, ll, int64(9223372036854775807)))
		}
		_ = v121 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:91:3:
		v122 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v122
		if v123 = v122 == int32(19); !v123 {
			t_printf(tls, __ccgo_ts+4438, libc.VaList(bp+16, msg, i, int32(19)))
		}
		_ = v123 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:92:3:
		v124 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v124
		if v125 = v124 == int32(ERANGE); !v125 {
			t_printf(tls, __ccgo_ts+4509, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v125 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:93:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 4573
		v127 = __ccgo_ts + 2562
		s = v127
		v126 = libc.Xstrtoll(tls, v127, bp, 0)
		ll = v126
		if v128 = v126 == -libc.Int64FromInt64(9223372036854775807)-libc.Int64FromInt32(1); !v128 {
			t_printf(tls, __ccgo_ts+4614, libc.VaList(bp+16, __ccgo_ts+4573, ll, -libc.Int64FromInt64(9223372036854775807)-libc.Int64FromInt32(1)))
		}
		_ = v128 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:94:3:
		v129 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v129
		if v130 = v129 == int32(20); !v130 {
			t_printf(tls, __ccgo_ts+4686, libc.VaList(bp+16, msg, i, int32(20)))
		}
		_ = v130 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:95:3:
		v131 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v131
		if v132 = v131 == int32(ERANGE); !v132 {
			t_printf(tls, __ccgo_ts+4757, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v132 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:96:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 4821
		v134 = __ccgo_ts + 2829
		s = v134
		v133 = libc.Xstrtoull(tls, v134, bp, 0)
		ull = v133
		if v135 = v133 == uint64(18446744073709551615); !v135 {
			t_printf(tls, __ccgo_ts+4863, libc.VaList(bp+16, __ccgo_ts+4821, ull, uint64(18446744073709551615)))
		}
		_ = v135 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:97:3:
		v136 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v136
		if v137 = v136 == int32(20); !v137 {
			t_printf(tls, __ccgo_ts+4935, libc.VaList(bp+16, msg, i, int32(20)))
		}
		_ = v137 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:98:3:
		v138 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v138
		if v139 = v138 == int32(ERANGE); !v139 {
			t_printf(tls, __ccgo_ts+5006, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v139 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:99:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 5070
		v141 = __ccgo_ts + 1070
		s = v141
		v140 = libc.Xstrtoull(tls, v141, bp, 0)
		ull = v140
		if v142 = v140 == -libc.Uint64FromUint64(1); !v142 {
			t_printf(tls, __ccgo_ts+5094, libc.VaList(bp+16, __ccgo_ts+5070, ull, -libc.Uint64FromUint64(1)))
		}
		_ = v142 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:100:3:
		v143 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v143
		if v144 = v143 == int32(2); !v144 {
			t_printf(tls, __ccgo_ts+5166, libc.VaList(bp+16, msg, i, int32(2)))
		}
		_ = v144 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:101:3:
		v145 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v145
		if v146 = v145 == 0; !v146 {
			t_printf(tls, __ccgo_ts+5238, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v146 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:102:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 5304
		v148 = __ccgo_ts + 1302
		s = v148
		v147 = libc.Xstrtoull(tls, v148, bp, 0)
		ull = v147
		if v149 = v147 == -libc.Uint64FromUint64(2); !v149 {
			t_printf(tls, __ccgo_ts+5328, libc.VaList(bp+16, __ccgo_ts+5304, ull, -libc.Uint64FromUint64(2)))
		}
		_ = v149 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:103:3:
		v150 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v150
		if v151 = v150 == int32(2); !v151 {
			t_printf(tls, __ccgo_ts+5401, libc.VaList(bp+16, msg, i, int32(2)))
		}
		_ = v151 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:104:3:
		v152 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v152
		if v153 = v152 == 0; !v153 {
			t_printf(tls, __ccgo_ts+5473, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v153 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:105:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 5539
		v155 = __ccgo_ts + 3508
		s = v155
		v154 = libc.Xstrtoull(tls, v155, bp, 0)
		ull = v154
		if v156 = v154 == -libc.Uint64FromUint64(9223372036854775808); !v156 {
			t_printf(tls, __ccgo_ts+5581, libc.VaList(bp+16, __ccgo_ts+5539, ull, -libc.Uint64FromUint64(9223372036854775808)))
		}
		_ = v156 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:106:3:
		v157 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v157
		if v158 = v157 == int32(20); !v158 {
			t_printf(tls, __ccgo_ts+5654, libc.VaList(bp+16, msg, i, int32(20)))
		}
		_ = v158 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:107:3:
		v159 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v159
		if v160 = v159 == 0; !v160 {
			t_printf(tls, __ccgo_ts+5726, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v160 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:108:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 5792
		v162 = __ccgo_ts + 2562
		s = v162
		v161 = libc.Xstrtoull(tls, v162, bp, 0)
		ull = v161
		if v163 = v161 == -libc.Uint64FromUint64(9223372036854775809); !v163 {
			t_printf(tls, __ccgo_ts+5834, libc.VaList(bp+16, __ccgo_ts+5792, ull, -libc.Uint64FromUint64(9223372036854775809)))
		}
		_ = v163 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:109:3:
		v164 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v164
		if v165 = v164 == int32(20); !v165 {
			t_printf(tls, __ccgo_ts+5907, libc.VaList(bp+16, msg, i, int32(20)))
		}
		_ = v165 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:110:3:
		v166 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v166
		if v167 = v166 == 0; !v167 {
			t_printf(tls, __ccgo_ts+5979, libc.VaList(bp+16, msg, i, 0))
		}
		_ = v167 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:111:3:
		*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
		msg = __ccgo_ts + 6045
		v169 = __ccgo_ts + 4024
		s = v169
		v168 = libc.Xstrtoull(tls, v169, bp, 0)
		ull = v168
		if v170 = v168 == uint64(18446744073709551615); !v170 {
			t_printf(tls, __ccgo_ts+6088, libc.VaList(bp+16, __ccgo_ts+6045, ull, uint64(18446744073709551615)))
		}
		_ = v170 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:112:3:
		v171 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
		i = v171
		if v172 = v171 == int32(21); !v172 {
			t_printf(tls, __ccgo_ts+6170, libc.VaList(bp+16, msg, i, int32(21)))
		}
		_ = v172 || libc.Bool(libc.Int32FromInt32(0) != 0)
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:113:3:
		v173 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
		i = v173
		if v174 = v173 == int32(ERANGE); !v174 {
			t_printf(tls, __ccgo_ts+6242, libc.VaList(bp+16, msg, i, int32(ERANGE)))
		}
		_ = v174 || libc.Bool(libc.Int32FromInt32(0) != 0)
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:115:3:
		t_printf(tls, __ccgo_ts+6308, libc.VaList(bp+16, libc.Int32FromInt64(8)))
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:118:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6379
	v175 = libc.Xstrtol(tls, __ccgo_ts+6398, uintptr(0), int32(36))
	l = v175
	if v176 = v175 == libc.Int32FromInt32(35); !v176 {
		t_printf(tls, __ccgo_ts+6400, libc.VaList(bp+16, __ccgo_ts+6379, l, int32(35)))
	}
	_ = v176 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:119:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6453
	v177 = libc.Xstrtol(tls, __ccgo_ts+6502, uintptr(0), int32(2))
	l = v177
	if v178 = v177 == libc.Int32FromInt32(0x12345678); !v178 {
		t_printf(tls, __ccgo_ts+6535, libc.VaList(bp+16, __ccgo_ts+6453, l, int32(0x12345678)))
	}
	_ = v178 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:120:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6588
	v180 = __ccgo_ts + 6613
	s = v180
	v179 = libc.Xstrtol(tls, v180, bp, int32(16))
	l = v179
	if v181 = v179 == libc.Int32FromInt32(0x0f5f); !v181 {
		t_printf(tls, __ccgo_ts+6618, libc.VaList(bp+16, __ccgo_ts+6588, l, int32(0x0f5f)))
	}
	_ = v181 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:122:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6671
	v183 = __ccgo_ts + 6695
	s = v183
	v182 = libc.Xstrtol(tls, v183, bp, int32(16))
	l = v182
	if v184 = v182 == libc.Int32FromInt32(0); !v184 {
		t_printf(tls, __ccgo_ts+6699, libc.VaList(bp+16, __ccgo_ts+6671, l, 0))
	}
	_ = v184 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:123:2:
	v185 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
	i = v185
	if v186 = v185 == int32(1); !v186 {
		t_printf(tls, __ccgo_ts+6752, libc.VaList(bp+16, msg, i, int32(1)))
	}
	_ = v186 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:125:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6826
	v188 = __ccgo_ts + 6853
	s = v188
	v187 = libc.Xstrtol(tls, v188, bp, int32(16))
	l = v187
	if v189 = v187 == libc.Int32FromInt32(0x1234); !v189 {
		t_printf(tls, __ccgo_ts+6860, libc.VaList(bp+16, __ccgo_ts+6826, l, int32(0x1234)))
	}
	_ = v189 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:126:2:
	v190 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
	i = v190
	if v191 = v190 == int32(6); !v191 {
		t_printf(tls, __ccgo_ts+6913, libc.VaList(bp+16, msg, i, int32(6)))
	}
	_ = v191 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:128:2:
	*(*uintptr)(unsafe.Pointer(bp)) = libc.UintptrFromInt32(0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:129:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 6987
	v193 = __ccgo_ts + 7011
	s = v193
	v192 = libc.Xstrtol(tls, v193, bp, int32(37))
	l = v192
	if v194 = v192 == libc.Int32FromInt32(0); !v194 {
		t_printf(tls, __ccgo_ts+7015, libc.VaList(bp+16, __ccgo_ts+6987, l, 0))
	}
	_ = v194 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:130:2:
	v195 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
	i = v195
	if v196 = v195 == 0; !v196 {
		t_printf(tls, __ccgo_ts+7068, libc.VaList(bp+16, msg, i, 0))
	}
	_ = v196 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:131:2:
	v197 = *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))
	i = v197
	if v198 = v197 == int32(EINVAL); !v198 {
		t_printf(tls, __ccgo_ts+7140, libc.VaList(bp+16, msg, i, int32(EINVAL)))
	}
	_ = v198 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:133:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 7191
	v200 = __ccgo_ts + 7218
	s = v200
	v199 = libc.Xstrtol(tls, v200, bp, int32(8))
	l = v199
	if v201 = v199 == libc.Int32FromInt32(015437); !v201 {
		t_printf(tls, __ccgo_ts+7226, libc.VaList(bp+16, __ccgo_ts+7191, l, int32(015437)))
	}
	_ = v201 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:134:2:
	v202 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
	i = v202
	if v203 = v202 == int32(7); !v203 {
		t_printf(tls, __ccgo_ts+7279, libc.VaList(bp+16, msg, i, int32(7)))
	}
	_ = v203 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:136:2:
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	msg = __ccgo_ts + 7351
	v205 = __ccgo_ts + 7374
	s = v205
	v204 = libc.Xstrtol(tls, v205, bp, 0)
	l = v204
	if v206 = v204 == libc.Int32FromInt32(1); !v206 {
		t_printf(tls, __ccgo_ts+7378, libc.VaList(bp+16, __ccgo_ts+7351, l, int32(1)))
	}
	_ = v206 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:137:2:
	v207 = int32(*(*uintptr)(unsafe.Pointer(bp))) - int32(s)
	i = v207
	if v208 = v207 == int32(3); !v208 {
		t_printf(tls, __ccgo_ts+7431, libc.VaList(bp+16, msg, i, int32(3)))
	}
	_ = v208 || libc.Bool(libc.Int32FromInt32(0) != 0)
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/functional/strtol.c:138:2:
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
		fd = libc.Xopen(tls, __ccgo_ts+7503, O_RDONLY, 0)
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
		t_printf(tls, __ccgo_ts+7513, libc.VaList(bp+8, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		Fs:    __ccgo_ts + 7557,
	},
	1: {
		Fflag: int32(FE_INVALID),
		Fs:    __ccgo_ts + 7565,
	},
	2: {
		Fflag: int32(FE_DIVBYZERO),
		Fs:    __ccgo_ts + 7573,
	},
	3: {
		Fflag: int32(FE_UNDERFLOW),
		Fs:    __ccgo_ts + 7583,
	},
	4: {
		Fflag: int32(FE_OVERFLOW),
		Fs:    __ccgo_ts + 7593,
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
				v2 = __ccgo_ts + 7602
			} else {
				v2 = __ccgo_ts
			}
			p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+7604, libc.VaList(bp+8, v2, eflags[i].Fs)))
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
			v3 = __ccgo_ts + 7602
		} else {
			v3 = __ccgo_ts
		}
		p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+7609, libc.VaList(bp+8, v3, f & ^all)))
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:123:3:
		all = f
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:125:2:
	if all != 0 {
		v4 = __ccgo_ts
	} else {
		v4 = __ccgo_ts + 7614
	}
	p += uintptr(libc.Xsprintf(tls, p, __ccgo_ts+7616, libc.VaList(bp+8, v4)))
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
		return __ccgo_ts + 7619
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:2:
		fallthrough
	case int32(FE_TOWARDZERO):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:134:11:
		return __ccgo_ts + 7622
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:2:
		fallthrough
	case int32(FE_UPWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:137:11:
		return __ccgo_ts + 7625
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:2:
		fallthrough
	case int32(FE_DOWNWARD):
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:140:11:
		return __ccgo_ts + 7628
	}
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/mtest.c:143:2:
	return __ccgo_ts + 7631
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
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+7634, libc.VaList(bp+8, int32(s)-int32(argv0), argv0, p))
	} else {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/path.c:14:3:
		k = libc.Xsnprintf(tls, buf, n, __ccgo_ts+7642, libc.VaList(bp+8, p))
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
		t_printf(tls, __ccgo_ts+7647, libc.VaList(bp+24, r, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
		t_printf(tls, __ccgo_ts+7690, libc.VaList(bp+24, r, lim, libc.Xstrerror(tls, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))))
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
	_ = libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7739) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7747) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7759) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7771) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7783) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts+7792) != 0 || libc.Xsetlocale(tls, LC_CTYPE, __ccgo_ts) != 0
	// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:17:2:
	if libc.Xstrcmp(tls, libc.Xnl_langinfo(tls, int32(CODESET)), __ccgo_ts+7792) != 0 {
		// /home/jnml/src/modernc.org/builder/.exclude/modernc.org/libc/v2/internal/nsz.repo.hu/libc-test/src/common/utf8.c:18:3:
		return t_printf(tls, __ccgo_ts+7798, libc.VaList(bp+8, libc.Xnl_langinfo(tls, int32(CODESET))))
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

var __ccgo_ts1 = "\x00atol(\"2147483647\")\x002147483647\x00src/functional/strtol.c:31: %s failed (max 32bit signed %ld != %ld)\n\x00strtol(\"2147483647\", 0, 0)\x00src/functional/strtol.c:32: %s failed (max 32bit signed %ld != %ld)\n\x00strtoul(\"4294967295\", 0, 0)\x004294967295\x00src/functional/strtol.c:33: %s failed (max 32bit unsigned %lu != %lu)\n\x00strtol(s=\"2147483648\", &c, 0)\x002147483648\x00src/functional/strtol.c:36: %s failed (uncaught overflow %ld != %ld)\n\x00src/functional/strtol.c:37: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:38: %s failed (missing errno %d != %d)\n\x00strtol(s=\"-2147483649\", &c, 0)\x00-2147483649\x00src/functional/strtol.c:39: %s failed (uncaught overflow %ld != %ld)\n\x00src/functional/strtol.c:40: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:41: %s failed (missing errno %d != %d)\n\x00strtoul(s=\"4294967296\", &c, 0)\x004294967296\x00src/functional/strtol.c:42: %s failed (uncaught overflow %lu != %lu)\n\x00src/functional/strtol.c:43: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:44: %s failed (missing errno %d != %d)\n\x00strtoul(s=\"-1\", &c, 0)\x00-1\x00src/functional/strtol.c:45: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:46: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:47: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-2\", &c, 0)\x00-2\x00src/functional/strtol.c:48: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:49: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:50: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-2147483648\", &c, 0)\x00-2147483648\x00src/functional/strtol.c:51: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:52: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:53: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-2147483649\", &c, 0)\x00src/functional/strtol.c:54: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:55: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:56: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-4294967296\", &c, 0)\x00-4294967296\x00src/functional/strtol.c:57: %s failed (uncaught negative overflow %lu != %lu)\n\x00src/functional/strtol.c:58: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:59: %s failed (spurious errno %d != %d)\n\x00strtol(s=\"9223372036854775808\", &c, 0)\x009223372036854775808\x00src/functional/strtol.c:61: %s failed (uncaught overflow %ld != %ld)\n\x00src/functional/strtol.c:62: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:63: %s failed (missing errno %d != %d)\n\x00strtol(s=\"-9223372036854775809\", &c, 0)\x00-9223372036854775809\x00src/functional/strtol.c:64: %s failed (uncaught overflow %ld != %ld)\n\x00src/functional/strtol.c:65: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:66: %s failed (missing errno %d != %d)\n\x00strtoul(s=\"18446744073709551616\", &c, 0)\x0018446744073709551616\x00src/functional/strtol.c:67: %s failed (uncaught overflow %lu != %lu)\n\x00src/functional/strtol.c:68: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:69: %s failed (missing errno %d != %d)\n\x00src/functional/strtol.c:70: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:71: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:72: %s failed (spurious errno %d != %d)\n\x00src/functional/strtol.c:73: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:74: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:75: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-9223372036854775808\", &c, 0)\x00-9223372036854775808\x00src/functional/strtol.c:76: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:77: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:78: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-9223372036854775809\", &c, 0)\x00src/functional/strtol.c:79: %s failed (rejected negative %lu != %lu)\n\x00src/functional/strtol.c:80: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:81: %s failed (spurious errno %d != %d)\n\x00strtoul(s=\"-18446744073709551616\", &c, 0)\x00-18446744073709551616\x00src/functional/strtol.c:82: %s failed (uncaught negative overflow %lu != %lu)\n\x00src/functional/strtol.c:83: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:84: %s failed (spurious errno %d != %d)\n\x00src/functional/strtol.c:86: sizeof(long) == %d, not implemented\n\x00strtoll(s=\"9223372036854775808\", &c, 0)\x00src/functional/strtol.c:90: %s failed (uncaught overflow %lld != %lld)\n\x00src/functional/strtol.c:91: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:92: %s failed (missing errno %d != %d)\n\x00strtoll(s=\"-9223372036854775809\", &c, 0)\x00src/functional/strtol.c:93: %s failed (uncaught overflow %lld != %lld)\n\x00src/functional/strtol.c:94: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:95: %s failed (missing errno %d != %d)\n\x00strtoull(s=\"18446744073709551616\", &c, 0)\x00src/functional/strtol.c:96: %s failed (uncaught overflow %llu != %llu)\n\x00src/functional/strtol.c:97: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:98: %s failed (missing errno %d != %d)\n\x00strtoull(s=\"-1\", &c, 0)\x00src/functional/strtol.c:99: %s failed (rejected negative %llu != %llu)\n\x00src/functional/strtol.c:100: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:101: %s failed (spurious errno %d != %d)\n\x00strtoull(s=\"-2\", &c, 0)\x00src/functional/strtol.c:102: %s failed (rejected negative %llu != %llu)\n\x00src/functional/strtol.c:103: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:104: %s failed (spurious errno %d != %d)\n\x00strtoull(s=\"-9223372036854775808\", &c, 0)\x00src/functional/strtol.c:105: %s failed (rejected negative %llu != %llu)\n\x00src/functional/strtol.c:106: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:107: %s failed (spurious errno %d != %d)\n\x00strtoull(s=\"-9223372036854775809\", &c, 0)\x00src/functional/strtol.c:108: %s failed (rejected negative %llu != %llu)\n\x00src/functional/strtol.c:109: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:110: %s failed (spurious errno %d != %d)\n\x00strtoull(s=\"-18446744073709551616\", &c, 0)\x00src/functional/strtol.c:111: %s failed (uncaught negative overflow %llu != %llu)\n\x00src/functional/strtol.c:112: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:113: %s failed (spurious errno %d != %d)\n\x00src/functional/strtol.c:115: sizeof(long long) == %d, not implemented\n\x00strtol(\"z\", 0, 36)\x00z\x00src/functional/strtol.c:118: %s failed (%ld != %ld)\n\x00strtol(\"00010010001101000101011001111000\", 0, 2)\x0000010010001101000101011001111000\x00src/functional/strtol.c:119: %s failed (%ld != %ld)\n\x00strtol(s=\"0F5F\", &c, 16)\x000F5F\x00src/functional/strtol.c:120: %s failed (%ld != %ld)\n\x00strtol(s=\"0xz\", &c, 16)\x000xz\x00src/functional/strtol.c:122: %s failed (%ld != %ld)\n\x00src/functional/strtol.c:123: %s failed (wrong final position %ld != %ld)\n\x00strtol(s=\"0x1234\", &c, 16)\x000x1234\x00src/functional/strtol.c:125: %s failed (%ld != %ld)\n\x00src/functional/strtol.c:126: %s failed (wrong final position %ld != %ld)\n\x00strtol(s=\"123\", &c, 37)\x00123\x00src/functional/strtol.c:129: %s failed (%ld != %ld)\n\x00src/functional/strtol.c:130: %s failed (wrong final position %d != %d)\n\x00src/functional/strtol.c:131: %s failed (%d != %d)\n\x00strtol(s=\"  15437\", &c, 8)\x00  15437\x00src/functional/strtol.c:133: %s failed (%ld != %ld)\n\x00src/functional/strtol.c:134: %s failed (wrong final position %d != %d)\n\x00strtol(s=\"  1\", &c, 0)\x00  1\x00src/functional/strtol.c:136: %s failed (%ld != %ld)\n\x00src/functional/strtol.c:137: %s failed (wrong final position %d != %d)\n\x00/dev/null\x00src/common/memfill.c:12: vmfill failed: %s\n\x00INEXACT\x00INVALID\x00DIVBYZERO\x00UNDERFLOW\x00OVERFLOW\x00|\x00%s%s\x00%s%d\x000\x00%s\x00RN\x00RZ\x00RU\x00RD\x00R?\x00%.*s/%s\x00./%s\x00src/common/setrlim.c:11: getrlimit %d: %s\n\x00src/common/setrlim.c:21: setrlimit(%d, %ld): %s\n\x00C.UTF-8\x00POSIX.UTF-8\x00en_US.UTF-8\x00en_GB.UTF-8\x00en.UTF-8\x00UTF-8\x00src/common/utf8.c:18: cannot set UTF-8 locale for test (codeset=%s)\n\x00"
