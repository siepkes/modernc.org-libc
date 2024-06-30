// Code generated for windows/amd64 by 'ccgo --cpp=/usr/bin/x86_64-w64-mingw32-gcc --goos=windows --goarch=amd64 --package-name libc --prefix-external=X --prefix-field=F --prefix-static-internal=_ --prefix-static-none=_ --prefix-tagged-struct=T --prefix-tagged-union=T --prefix-typename=T --winapi-test dump --winapi=ctype.h --winapi=float.h --winapi=io.h --winapi=libucrt.c --winapi=locale.h --winapi=malloc.h --winapi=math.h --winapi=process.h --winapi=types.h --winapi=stat.h --winapi=stdio.h --winapi=stdlib.h --winapi=string.h --winapi=time.h --winapi=timeb.h --winapi=wchar.h --winapi=winbase.h -build-lines=  -eval-all-macros -hide __acrt_iob_func -hide _errno -hide _wgetenv -hide _wputenv -hide exit -hide lldiv -hide qsort -hide __sep__ -hide __create_locale -hide __free_locale -hide __get_current_locale -hide __iob_func -hide __lock_fhandle -hide __updatetlocinfo -hide __updatetmbcinfo -hide _beginthread -hide _beginthreadex -hide _endthreadex -hide _filbuf -hide _flsbuf -hide _get_amblksiz -hide _get_osplatform -hide _get_osver -hide _get_output_format -hide _get_sbh_threshold -hide _get_winmajor -hide _get_winminor -hide _get_winver -hide _heapadd -hide _heapset -hide _heapused -hide _matherr -hide _onexit -hide _set_amblksiz -hide _set_malloc_crt_max_wait -hide _set_output_format -hide _set_sbh_threshold -hide _strcmpi -hide _strnset_l -hide _strset_l -hide _unlock_fhandle -hide _wcsncpy_l -hide _wctime -hide _wctime_s -hide _wgetdcwd_nolock -hide access -hide at_quick_exit -hide atexit -hide chdir -hide chmod -hide chsize -hide close -hide creat -hide cwait -hide dup -hide dup2 -hide eof -hide execv -hide execve -hide execvp -hide execvpe -hide fcloseall -hide fdopen -hide fgetchar -hide fgetpos64 -hide filelength -hide fileno -hide flushall -hide fopen64 -hide fpreset -hide fputchar -hide fsetpos64 -hide ftime -hide fwide -hide getcwd -hide getpid -hide getw -hide isatty -hide itoa -hide lltoa -hide lltow -hide locking -hide lseek -hide lseek64 -hide ltoa -hide memccpy -hide memicmp -hide mempcpy -hide mkdir -hide mkstemp -hide mktemp -hide onexit -hide putenv -hide putw -hide read -hide rmdir -hide rmtmp -hide setmode -hide spawnv -hide spawnve -hide spawnvp -hide spawnvpe -hide strcasecmp -hide strcmpi -hide strdup -hide stricmp -hide strlwr -hide strlwr_l -hide strncasecmp -hide strnicmp -hide strnset -hide strrev -hide strset -hide strtok_r -hide strupr -hide swab -hide tell -hide tempnam -hide tzset -hide ulltoa -hide ulltow -hide ultoa -hide umask -hide unlink -hide wcsdup -hide wcsicmp -hide wcsicoll -hide wcslwr -hide wcsnicmp -hide wcsnset -hide wcsrev -hide wcsset -hide wcsupr -hide wmemchr -hide wmemcmp -hide wmemcpy -hide wmemmove -hide wmempcpy -hide wmemset -hide write -hide wtoll -ignore-link-errors -import syscall -keep-strings -o libc_windows_arm64.go libucrt.c', DO NOT EDIT.

package libc

import (
	"reflect"
	"unsafe"

	"syscall"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const BUFSIZ = 512
const CLK_TCK = 1000
const CLOCKS_PER_SEC = 1000
const CLOCK_MONOTONIC = 1
const CLOCK_PROCESS_CPUTIME_ID = 2
const CLOCK_REALTIME = 0
const CLOCK_REALTIME_COARSE = 4
const CLOCK_THREAD_CPUTIME_ID = 3
const CW_DEFAULT = 524319
const DOMAIN = 1
const E2BIG = 7
const EACCES = 13
const EADDRINUSE = 100
const EADDRNOTAVAIL = 101
const EAFNOSUPPORT = 102
const EAGAIN = 11
const EALREADY = 103
const EBADF = 9
const EBADMSG = 104
const EBUSY = 16
const ECANCELED = 105
const ECHILD = 10
const ECONNABORTED = 106
const ECONNREFUSED = 107
const ECONNRESET = 108
const EDEADLK = 36
const EDEADLOCK = 36
const EDESTADDRREQ = 109
const EDOM = 33
const EEXIST = 17
const EFAULT = 14
const EFBIG = 27
const EHOSTUNREACH = 110
const EIDRM = 111
const EILSEQ = 42
const EINPROGRESS = 112
const EINTR = 4
const EINVAL = 22
const EIO = 5
const EISCONN = 113
const EISDIR = 21
const ELOOP = 114
const EMFILE = 24
const EMLINK = 31
const EMSGSIZE = 115
const ENAMETOOLONG = 38
const ENETDOWN = 116
const ENETRESET = 117
const ENETUNREACH = 118
const ENFILE = 23
const ENOBUFS = 119
const ENODATA = 120
const ENODEV = 19
const ENOENT = 2
const ENOEXEC = 8
const ENOFILE = 2
const ENOLCK = 39
const ENOLINK = 121
const ENOMEM = 12
const ENOMSG = 122
const ENOPROTOOPT = 123
const ENOSPC = 28
const ENOSR = 124
const ENOSTR = 125
const ENOSYS = 40
const ENOTCONN = 126
const ENOTDIR = 20
const ENOTEMPTY = 41
const ENOTRECOVERABLE = 127
const ENOTSOCK = 128
const ENOTSUP = 129
const ENOTTY = 25
const ENXIO = 6
const EOF = -1
const EOPNOTSUPP = 130
const EOVERFLOW = 132
const EOWNERDEAD = 133
const EPERM = 1
const EPIPE = 32
const EPROTO = 134
const EPROTONOSUPPORT = 135
const EPROTOTYPE = 136
const ERANGE = 34
const EROFS = 30
const ESPIPE = 29
const ESRCH = 3
const ETIME = 137
const ETIMEDOUT = 138
const ETXTBSY = 139
const EWOULDBLOCK = 140
const EXDEV = 18
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const FILENAME_MAX = 260
const FOPEN_MAX = 20
const FP_INFINITE = 1280
const FP_NAN = 256
const FP_NDENORM = 16
const FP_NINF = 4
const FP_NNORM = 8
const FP_NORMAL = 1024
const FP_NZERO = 32
const FP_PDENORM = 128
const FP_PINF = 512
const FP_PNORM = 256
const FP_PZERO = 64
const FP_QNAN = 2
const FP_SNAN = 1
const FP_SUBNORMAL = 17408
const FP_ZERO = 16384
const F_OK = 0
const HUGE = 0
const HUGE_VAL = 0
const HUGE_VALF = 0
const HUGE_VALL = 0
const INFINITY = 0
const LC_ALL = 0
const LC_COLLATE = 1
const LC_CTYPE = 2
const LC_MAX = 5
const LC_MIN = 0
const LC_MONETARY = 3
const LC_NUMERIC = 4
const LC_TIME = 5
const L_tmpnam = 12
const L_tmpnam_s = 12
const MB_CUR_MAX = 0
const MB_LEN_MAX = 5
const MCW_PC = 196608
const MINGW_HAS_DDK_H = 1
const MINGW_HAS_SECURE_API = 1
const M_1_PI = 0
const M_2_PI = 0
const M_2_SQRTPI = 0
const M_E = 0
const M_LN10 = 0
const M_LN2 = 0
const M_LOG10E = 0
const M_LOG2E = 0
const M_PI = 0
const M_PI_2 = 0
const M_PI_4 = 0
const M_SQRT1_2 = 0
const M_SQRT2 = 0
const NAN = 0
const OLD_P_OVERLAY = 2
const OVERFLOW = 3
const PATH_MAX = 260
const PC_24 = 131072
const PC_53 = 65536
const PC_64 = 0
const PLOSS = 6
const P_DETACH = 4
const P_NOWAIT = 1
const P_NOWAITO = 3
const P_OVERLAY = 2
const P_WAIT = 0
const P_tmpdir = "_P_tmpdir"
const RAND_MAX = 32767
const R_OK = 4
const SEEK_CUR = 1
const SEEK_END = 2
const SEEK_SET = 0
const SING = 2
const SIZE_MAX = 18446744073709551615
const SSIZE_MAX = 9223372036854775807
const STDERR_FILENO = 2
const STDIN_FILENO = 0
const STDOUT_FILENO = 1
const STRUNCATE = 80
const SYS_OPEN = 20
const TIMER_ABSTIME = 1
const TIME_UTC = 1
const TLOSS = 5
const TMP_MAX = 32767
const TMP_MAX_S = 32767
const UNALIGNED = 0
const UNDERFLOW = 4
const USE___UUIDOF = 0
const WAIT_CHILD = 0
const WAIT_GRANDCHILD = 1
const WCHAR_MAX = 65535
const WCHAR_MIN = 0
const WIN32 = 1
const WIN64 = 1
const WINNT = 1
const W_OK = 2
const X_OK = 1
const _ALLOCA_S_HEAP_MARKER = 56797
const _ALLOCA_S_MARKER_SIZE = 16
const _ALLOCA_S_STACK_MARKER = 52428
const _ALLOCA_S_THRESHOLD = 1024
const _ALPHA = 259
const _ANONYMOUS_STRUCT = 0
const _ANONYMOUS_UNION = 0
const _ARGMAX = 100
const _A_ARCH = 32
const _A_HIDDEN = 2
const _A_NORMAL = 0
const _A_RDONLY = 1
const _A_SUBDIR = 16
const _A_SYSTEM = 4
const _BLANK = 64
const _CALL_REPORTFAULT = 2
const _CONTROL = 32
const _CRTIMP2 = "_CRTIMP"
const _CRTIMP_ALTERNATIVE = "_CRTIMP"
const _CRTIMP_NOIA64 = "_CRTIMP"
const _CRTIMP_PURE = "_CRTIMP"
const _CRT_INTERNAL_LOCAL_PRINTF_OPTIONS = 4
const _CRT_INTERNAL_LOCAL_SCANF_OPTIONS = 2
const _CRT_INTERNAL_PRINTF_LEGACY_MSVCRT_COMPATIBILITY = 8
const _CRT_INTERNAL_PRINTF_LEGACY_THREE_DIGIT_EXPONENTS = 16
const _CRT_INTERNAL_PRINTF_LEGACY_VSPRINTF_NULL_TERMINATION = 1
const _CRT_INTERNAL_PRINTF_LEGACY_WIDE_SPECIFIERS = 4
const _CRT_INTERNAL_PRINTF_STANDARD_SNPRINTF_BEHAVIOR = 2
const _CRT_INTERNAL_SCANF_LEGACY_MSVCRT_COMPATIBILITY = 4
const _CRT_INTERNAL_SCANF_LEGACY_WIDE_SPECIFIERS = 2
const _CRT_INTERNAL_SCANF_SECURECRT = 1
const _CRT_SECURE_CPP_NOTHROW = 0
const _CVTBUFSIZE = 349
const _CW_DEFAULT = 524319
const _DIGIT = 4
const _DISABLE_PER_THREAD_LOCALE = 2
const _DISABLE_PER_THREAD_LOCALE_GLOBAL = 32
const _DISABLE_PER_THREAD_LOCALE_NEW = 512
const _DN_FLUSH = 16777216
const _DN_SAVE = 0
const _DOMAIN = 1
const _EM_DENORMAL = 524288
const _EM_INEXACT = 1
const _EM_INVALID = 16
const _EM_OVERFLOW = 4
const _EM_UNDERFLOW = 2
const _EM_ZERODIVIDE = 8
const _ENABLE_PER_THREAD_LOCALE = 1
const _ENABLE_PER_THREAD_LOCALE_GLOBAL = 16
const _ENABLE_PER_THREAD_LOCALE_NEW = 256
const _FPCLASS_ND = 16
const _FPCLASS_NINF = 4
const _FPCLASS_NN = 8
const _FPCLASS_NZ = 32
const _FPCLASS_PD = 128
const _FPCLASS_PINF = 512
const _FPCLASS_PN = 256
const _FPCLASS_PZ = 64
const _FPCLASS_QNAN = 2
const _FPCLASS_SNAN = 1
const _FPE_DENORMAL = 130
const _FPE_EXPLICITGEN = 140
const _FPE_INEXACT = 134
const _FPE_INVALID = 129
const _FPE_OVERFLOW = 132
const _FPE_SQRTNEG = 136
const _FPE_STACKOVERFLOW = 138
const _FPE_STACKUNDERFLOW = 139
const _FPE_UNDERFLOW = 133
const _FPE_UNEMULATED = 135
const _FPE_ZERODIVIDE = 131
const _FREEENTRY = 0
const _HEAPBADBEGIN = -3
const _HEAPBADNODE = -4
const _HEAPBADPTR = -6
const _HEAPEMPTY = -1
const _HEAPEND = -5
const _HEAPOK = -2
const _HEAP_MAXREQ = 18446744073709551584
const _HEX = 128
const _I16_MAX = 32767
const _I16_MIN = -32768
const _I32_MAX = 2147483647
const _I32_MIN = -2147483648
const _I64_MAX = 9223372036854775807
const _I64_MIN = -9223372036854775808
const _I8_MAX = 127
const _I8_MIN = -128
const _IC_AFFINE = 262144
const _IC_PROJECTIVE = 0
const _INTEGRAL_MAX_BITS = 64
const _IOB_ENTRIES = 20
const _IOFBF = 0
const _IOLBF = 64
const _IONBF = 4
const _LEADBYTE = 32768
const _LOWER = 2
const _MAX_DIR = 256
const _MAX_DRIVE = 3
const _MAX_ENV = 32767
const _MAX_EXT = 256
const _MAX_FNAME = 256
const _MAX_PATH = 260
const _MAX_WAIT_MALLOC_CRT = 60000
const _MCRTIMP = "_CRTIMP"
const _MCW_DN = 50331648
const _MCW_EM = 524319
const _MCW_IC = 262144
const _MCW_PC = 196608
const _MCW_RC = 768
const _MRTIMP2 = "_CRTIMP"
const _M_AMD64 = 100
const _M_X64 = 100
const _NFILE = 512
const _NLSCMPERROR = 2147483647
const _NSTREAM_ = 512
const _OLD_P_OVERLAY = 2
const _OUT_TO_DEFAULT = 0
const _OUT_TO_MSGBOX = 2
const _OUT_TO_STDERR = 1
const _OVERFLOW = 3
const _PC_24 = 131072
const _PC_53 = 65536
const _PC_64 = 0
const _PLOSS = 6
const _POSIX_CPUTIME = 200809
const _POSIX_MONOTONIC_CLOCK = 200809
const _POSIX_THREAD_CPUTIME = 200809
const _POSIX_TIMERS = 200809
const _PUNCT = 16
const _P_DETACH = 4
const _P_NOWAIT = 1
const _P_NOWAITO = 3
const _P_OVERLAY = 2
const _P_WAIT = 0
const _P_tmpdir = "\\\\"
const _RC_CHOP = 768
const _RC_DOWN = 256
const _RC_NEAR = 0
const _RC_UP = 512
const _REPORT_ERRMODE = 3
const _SECURECRT_FILL_BUFFER_PATTERN = 253
const _SING = 2
const _SPACE = 8
const _SW_SQRTNEG = 128
const _SW_STACKOVERFLOW = 512
const _SW_STACKUNDERFLOW = 1024
const _SW_UNEMULATED = 64
const _SYS_OPEN = 20
const _TLOSS = 5
const _TRUNCATE = -1
const _TWO_DIGIT_EXPONENT = 1
const _UI16_MAX = 65535
const _UI32_MAX = 4294967295
const _UI64_MAX = 18446744073709551615
const _UI8_MAX = 255
const _UNDERFLOW = 4
const _UPPER = 1
const _USEDENTRY = 1
const _WAIT_CHILD = 0
const _WAIT_GRANDCHILD = 1
const _WConst_return = 0
const _WIN32 = 1
const _WIN32_WINNT = 2560
const _WIN64 = 1
const _WRITE_ABORT_MSG = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BIGGEST_ALIGNMENT__ = 16
const __BYTE_ORDER__ = 1234
const __C89_NAMELESS = 0
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __CRTDECL = "__cdecl"
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 2
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DBL_MIN_10_EXP__ = -307
const __DBL_MIN_EXP__ = -1021
const __DEC128_EPSILON__ = 0
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = 0
const __DEC128_MIN_EXP__ = -6142
const __DEC128_MIN__ = 0
const __DEC128_SUBNORMAL_MIN__ = 0
const __DEC32_EPSILON__ = 0
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 0
const __DEC32_MIN_EXP__ = -94
const __DEC32_MIN__ = 0
const __DEC32_SUBNORMAL_MIN__ = 0
const __DEC64_EPSILON__ = 0
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = 0
const __DEC64_MIN_EXP__ = -382
const __DEC64_MIN__ = 0
const __DEC64_SUBNORMAL_MIN__ = 0
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = 1234
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 0
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 0
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 2
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = 0
const __FLT128_MIN_10_EXP__ = -4931
const __FLT128_MIN_EXP__ = -16381
const __FLT128_MIN__ = 0
const __FLT128_NORM_MAX__ = 0
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 0
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 0
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 2
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 0
const __FLT32X_MIN_10_EXP__ = -307
const __FLT32X_MIN_EXP__ = -1021
const __FLT32X_MIN__ = 0
const __FLT32X_NORM_MAX__ = 0
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 0
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 0
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 2
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 0
const __FLT32_MIN_10_EXP__ = -37
const __FLT32_MIN_EXP__ = -125
const __FLT32_MIN__ = 0
const __FLT32_NORM_MAX__ = 0
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 0
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 0
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 2
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = 0
const __FLT64X_MIN_10_EXP__ = -4931
const __FLT64X_MIN_EXP__ = -16381
const __FLT64X_MIN__ = 0
const __FLT64X_NORM_MAX__ = 0
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 0
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 0
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 2
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 0
const __FLT64_MIN_10_EXP__ = -307
const __FLT64_MIN_EXP__ = -1021
const __FLT64_MIN__ = 0
const __FLT64_NORM_MAX__ = 0
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 0
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 0
const __FLT_EVAL_METHOD_TS_18661_3__ = 2
const __FLT_EVAL_METHOD__ = 2
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 2
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 0
const __FLT_MIN_10_EXP__ = -37
const __FLT_MIN_EXP__ = -125
const __FLT_MIN__ = 0
const __FLT_NORM_MAX__ = 0
const __FLT_RADIX__ = 2
const __FUNCTION__ = 0
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 0
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-16LE"
const __GNUC__ = 12
const __GNU_EXTENSION = 0
const __GOT_SECURE_LIB__ = 200411
const __GXX_ABI_VERSION = 1017
const __GXX_MERGED_TYPEINFO_NAMES = 0
const __GXX_TYPEINFO_EQUALITY_INLINE = 0
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 32767
const __INT32_MAX__ = 2147483647
const __INT32_TYPE__ = 0
const __INT64_MAX__ = 9223372036854775807
const __INT8_MAX__ = 127
const __INTMAX_MAX__ = 9223372036854775807
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 9223372036854775807
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 32767
const __INT_FAST16_WIDTH__ = 16
const __INT_FAST32_MAX__ = 2147483647
const __INT_FAST32_TYPE__ = 0
const __INT_FAST32_WIDTH__ = 32
const __INT_FAST64_MAX__ = 9223372036854775807
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 127
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 32767
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 2147483647
const __INT_LEAST32_TYPE__ = 0
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 9223372036854775807
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 127
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 2147483647
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 0
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 0
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 2
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 0
const __LDBL_MIN_10_EXP__ = -307
const __LDBL_MIN_EXP__ = -1021
const __LDBL_MIN__ = 0
const __LDBL_NORM_MAX__ = 0
const __LONG32 = 0
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 9223372036854775807
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 2147483647
const __LONG_WIDTH__ = 32
const __MINGW32_MAJOR_VERSION = 3
const __MINGW32_MINOR_VERSION = 11
const __MINGW32__ = 1
const __MINGW64_VERSION_BUGFIX = 0
const __MINGW64_VERSION_MAJOR = 10
const __MINGW64_VERSION_MINOR = 0
const __MINGW64_VERSION_RC = 0
const __MINGW64_VERSION_STATE = "alpha"
const __MINGW64__ = 1
const __MINGW_ATTRIB_DEPRECATED_MSVC2005 = 0
const __MINGW_ATTRIB_DEPRECATED_SEC_WARN = 0
const __MINGW_DEBUGBREAK_IMPL = 1
const __MINGW_FORTIFY_LEVEL = 0
const __MINGW_FORTIFY_VA_ARG = 0
const __MINGW_FPCLASS_DEFINED = 1
const __MINGW_GCC_VERSION = 120000
const __MINGW_HAVE_ANSI_C99_PRINTF = 1
const __MINGW_HAVE_ANSI_C99_SCANF = 1
const __MINGW_HAVE_WIDE_C99_PRINTF = 1
const __MINGW_HAVE_WIDE_C99_SCANF = 1
const __MINGW_MSVC2005_DEPREC_STR = "This POSIX function is deprecated beginning in Visual C++ 2005, use _CRT_NONSTDC_NO_DEPRECATE to disable deprecation"
const __MINGW_SEC_WARN_STR = "This function or variable may be unsafe, use _CRT_SECURE_NO_WARNINGS to disable deprecation"
const __MINGW_USE_UNDERSCORE_PREFIX = 0
const __MSVCRT_VERSION__ = 3584
const __MSVCRT__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PCTYPE_FUNC = 0
const __PIC__ = 1
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = 0
const __PTRDIFF_MAX__ = 9223372036854775807
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 127
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SEH__ = 1
const __SHRT_MAX__ = 32767
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 2147483647
const __SIG_ATOMIC_MIN__ = -2147483648
const __SIG_ATOMIC_TYPE__ = 0
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 4
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 2
const __SIZEOF_WINT_T__ = 2
const __SIZE_MAX__ = 18446744073709551615
const __SIZE_WIDTH__ = 64
const __STDC_HOSTED__ = 1
const __STDC_SECURE_LIB__ = 200411
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 65535
const __UINT32_MAX__ = 4294967295
const __UINT64_MAX__ = 18446744073709551615
const __UINT8_MAX__ = 255
const __UINTMAX_MAX__ = 18446744073709551615
const __UINTPTR_MAX__ = 18446744073709551615
const __UINT_FAST16_MAX__ = 65535
const __UINT_FAST32_MAX__ = 4294967295
const __UINT_FAST64_MAX__ = 18446744073709551615
const __UINT_FAST8_MAX__ = 255
const __UINT_LEAST16_MAX__ = 65535
const __UINT_LEAST32_MAX__ = 4294967295
const __UINT_LEAST64_MAX__ = 18446744073709551615
const __UINT_LEAST8_MAX__ = 255
const __USE_MINGW_ANSI_STDIO = 0
const __VERSION__ = "12-win32"
const __WCHAR_MAX__ = 65535
const __WCHAR_MIN__ = 0
const __WCHAR_WIDTH__ = 16
const __WIN32 = 1
const __WIN32__ = 1
const __WIN64 = 1
const __WIN64__ = 1
const __WINNT = 1
const __WINNT__ = 1
const __WINT_MAX__ = 65535
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 16
const __amd64 = 1
const __amd64__ = 1
const __argc = 0
const __argv = 0
const __clockid_t_defined = 1
const __code_model_medium__ = 1
const __int16 = 0
const __int32 = 0
const __int8 = 0
const __k8 = 1
const __k8__ = 1
const __mb_cur_max = 0
const __mingw_bos_ovr = "__mingw_ovr"
const __mingw_choose_expr = 0
const __pic__ = 1
const __setusermatherr = 0
const __stat64 = 0
const __wargv = 0
const __x86_64 = 1
const __x86_64__ = 1
const _acmdln = 0
const _clear87 = 0
const _copysignl = 0
const _daylight = 0
const _doserrno = 0
const _dstbias = 0
const _environ = 0
const _finddata_t = 0
const _finddatai64_t = 0
const _findfirst = 0
const _findfirsti64 = 0
const _findnext = 0
const _findnexti64 = 0
const _fmode = 0
const _fpecode = 0
const _fstat = 0
const _fstati64 = 0
const _ftime = 0
const _ftime_s = 0
const _hypotl = 0
const _inline = 0
const _iob = 0
const _pctype = 0
const _pgmptr = 0
const _pwctype = 0
const _stat = 0
const _stati64 = 0
const _status87 = 0
const _sys_errlist = 0
const _sys_nerr = 0
const _timeb = 0
const _timezone = 0
const _tzname = 0
const _wP_tmpdir = "\\\\"
const _wcmdln = 0
const _wctype = 0
const _wenviron = 0
const _wfinddata_t = 0
const _wfinddatai64_t = 0
const _wfindfirst = 0
const _wfindfirsti64 = 0
const _wfindnext = 0
const _wfindnexti64 = 0
const _wpgmptr = 0
const _wstat = 0
const _wstati64 = 0
const environ = 0
const errno = 0
const fstat64 = 0
const isascii = 0
const iscsym = 0
const iscsymf = 0
const matherr = 0
const onexit_t = 0
const pclose = 0
const popen = 0
const stat64 = 0
const stderr = 0
const stdin = 0
const stdout = 0
const strcasecmp = 0
const strncasecmp = 0
const sys_errlist = 0
const sys_nerr = 0
const toascii = 0
const wcswcs = 0
const wpopen = 0

type T__builtin_va_list = uintptr

type T__predefined_size_t = uint64

type T__predefined_wchar_t = uint16

type T__predefined_ptrdiff_t = int64

type T__gnuc_va_list = uintptr

type Tva_list = uintptr

type Tsize_t = uint64

type Tssize_t = int64

type Trsize_t = uint64

type Tintptr_t = int64

type Tuintptr_t = uint64

type Tptrdiff_t = int64

type Twchar_t = uint16

type Twint_t = uint16

type Twctype_t = uint16

type Terrno_t = int32

type T__time32_t = int32

type T__time64_t = int64

type Ttime_t = int64

type Tthreadlocaleinfostruct = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

type Tpthreadlocinfo = uintptr

type Tpthreadmbcinfo = uintptr

type T_locale_tstruct = struct {
	Flocinfo Tpthreadlocinfo
	Fmbcinfo Tpthreadmbcinfo
}

type Tlocaleinfo_struct = T_locale_tstruct

type T_locale_t = uintptr

type TLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
}

type TtagLC_ID = TLC_ID

type TLPLC_ID = uintptr

type Tthreadlocinfo = struct {
	F_locale_pctype      uintptr
	F_locale_mb_cur_max  int32
	F_locale_lc_codepage uint32
}

var proc__pctype_func = dll.NewProc("__pctype_func")

// __attribute__ ((__dllimport__)) unsigned short* __pctype_func(void);
func X__pctype_func(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__pctype_func->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__pctype_func.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_isctype = dll.NewProc("_isctype")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isctype(int _C,int _Type);
func X_isctype(tls *TLS, __C int32, __Type int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Type=%+v", __C, __Type)
		defer func() { trc(`X_isctype->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isctype.Addr(), uintptr(__C), uintptr(__Type))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isctype_l = dll.NewProc("_isctype_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isctype_l(int _C,int _Type,_locale_t _Locale);
func X_isctype_l(tls *TLS, __C int32, __Type int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Type=%+v _Locale=%+v", __C, __Type, __Locale)
		defer func() { trc(`X_isctype_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isctype_l.Addr(), uintptr(__C), uintptr(__Type), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisalpha = dll.NewProc("isalpha")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isalpha(int _C);
func Xisalpha(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisalpha->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisalpha.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isalpha_l = dll.NewProc("_isalpha_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isalpha_l(int _C,_locale_t _Locale);
func X_isalpha_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isalpha_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isalpha_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisupper = dll.NewProc("isupper")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isupper(int _C);
func Xisupper(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisupper->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisupper.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isupper_l = dll.NewProc("_isupper_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isupper_l(int _C,_locale_t _Locale);
func X_isupper_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isupper_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isupper_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procislower = dll.NewProc("islower")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) islower(int _C);
func Xislower(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xislower->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procislower.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_islower_l = dll.NewProc("_islower_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _islower_l(int _C,_locale_t _Locale);
func X_islower_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_islower_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_islower_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisdigit = dll.NewProc("isdigit")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isdigit(int _C);
func Xisdigit(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisdigit->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisdigit.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isdigit_l = dll.NewProc("_isdigit_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isdigit_l(int _C,_locale_t _Locale);
func X_isdigit_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isdigit_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isdigit_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisxdigit = dll.NewProc("isxdigit")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isxdigit(int _C);
func Xisxdigit(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisxdigit->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisxdigit.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isxdigit_l = dll.NewProc("_isxdigit_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isxdigit_l(int _C,_locale_t _Locale);
func X_isxdigit_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isxdigit_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isxdigit_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisspace = dll.NewProc("isspace")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isspace(int _C);
func Xisspace(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisspace->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisspace.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isspace_l = dll.NewProc("_isspace_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isspace_l(int _C,_locale_t _Locale);
func X_isspace_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isspace_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isspace_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procispunct = dll.NewProc("ispunct")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) ispunct(int _C);
func Xispunct(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xispunct->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procispunct.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_ispunct_l = dll.NewProc("_ispunct_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _ispunct_l(int _C,_locale_t _Locale);
func X_ispunct_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_ispunct_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ispunct_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisalnum = dll.NewProc("isalnum")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isalnum(int _C);
func Xisalnum(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisalnum->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisalnum.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isalnum_l = dll.NewProc("_isalnum_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isalnum_l(int _C,_locale_t _Locale);
func X_isalnum_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isalnum_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isalnum_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisprint = dll.NewProc("isprint")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isprint(int _C);
func Xisprint(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisprint->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisprint.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isprint_l = dll.NewProc("_isprint_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isprint_l(int _C,_locale_t _Locale);
func X_isprint_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isprint_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isprint_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisgraph = dll.NewProc("isgraph")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) isgraph(int _C);
func Xisgraph(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisgraph->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisgraph.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isgraph_l = dll.NewProc("_isgraph_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isgraph_l(int _C,_locale_t _Locale);
func X_isgraph_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isgraph_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isgraph_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociscntrl = dll.NewProc("iscntrl")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) iscntrl(int _C);
func Xiscntrl(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiscntrl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociscntrl.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iscntrl_l = dll.NewProc("_iscntrl_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iscntrl_l(int _C,_locale_t _Locale);
func X_iscntrl_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iscntrl_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iscntrl_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proctoupper = dll.NewProc("toupper")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) toupper(int _C);
func Xtoupper(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xtoupper->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctoupper.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proctolower = dll.NewProc("tolower")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) tolower(int _C);
func Xtolower(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xtolower->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctolower.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_tolower = dll.NewProc("_tolower")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _tolower(int _C);
func X_tolower(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X_tolower->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_tolower.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_tolower_l = dll.NewProc("_tolower_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _tolower_l(int _C,_locale_t _Locale);
func X_tolower_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_tolower_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_tolower_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_toupper = dll.NewProc("_toupper")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _toupper(int _C);
func X_toupper(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X_toupper->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_toupper.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_toupper_l = dll.NewProc("_toupper_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _toupper_l(int _C,_locale_t _Locale);
func X_toupper_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_toupper_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_toupper_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__isascii = dll.NewProc("__isascii")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __isascii(int _C);
func X__isascii(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__isascii->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__isascii.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__toascii = dll.NewProc("__toascii")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __toascii(int _C);
func X__toascii(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__toascii->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__toascii.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__iscsymf = dll.NewProc("__iscsymf")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __iscsymf(int _C);
func X__iscsymf(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__iscsymf->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__iscsymf.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__iscsym = dll.NewProc("__iscsym")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __iscsym(int _C);
func X__iscsym(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__iscsym->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__iscsym.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisblank = dll.NewProc("isblank")

// int __attribute__((__cdecl__)) isblank(int _C);
func Xisblank(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisblank->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisblank.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswalpha = dll.NewProc("iswalpha")

// int __attribute__((__cdecl__)) iswalpha(wint_t _C);
func Xiswalpha(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswalpha->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswalpha.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswalpha_l = dll.NewProc("_iswalpha_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswalpha_l(wint_t _C,_locale_t _Locale);
func X_iswalpha_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswalpha_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswalpha_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswupper = dll.NewProc("iswupper")

// int __attribute__((__cdecl__)) iswupper(wint_t _C);
func Xiswupper(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswupper->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswupper.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswupper_l = dll.NewProc("_iswupper_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswupper_l(wint_t _C,_locale_t _Locale);
func X_iswupper_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswupper_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswupper_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswlower = dll.NewProc("iswlower")

// int __attribute__((__cdecl__)) iswlower(wint_t _C);
func Xiswlower(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswlower->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswlower.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswlower_l = dll.NewProc("_iswlower_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswlower_l(wint_t _C,_locale_t _Locale);
func X_iswlower_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswlower_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswlower_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswdigit = dll.NewProc("iswdigit")

// int __attribute__((__cdecl__)) iswdigit(wint_t _C);
func Xiswdigit(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswdigit->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswdigit.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswdigit_l = dll.NewProc("_iswdigit_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswdigit_l(wint_t _C,_locale_t _Locale);
func X_iswdigit_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswdigit_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswdigit_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswxdigit = dll.NewProc("iswxdigit")

// int __attribute__((__cdecl__)) iswxdigit(wint_t _C);
func Xiswxdigit(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswxdigit->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswxdigit.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswxdigit_l = dll.NewProc("_iswxdigit_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswxdigit_l(wint_t _C,_locale_t _Locale);
func X_iswxdigit_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswxdigit_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswxdigit_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswspace = dll.NewProc("iswspace")

// int __attribute__((__cdecl__)) iswspace(wint_t _C);
func Xiswspace(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswspace->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswspace.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswspace_l = dll.NewProc("_iswspace_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswspace_l(wint_t _C,_locale_t _Locale);
func X_iswspace_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswspace_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswspace_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswpunct = dll.NewProc("iswpunct")

// int __attribute__((__cdecl__)) iswpunct(wint_t _C);
func Xiswpunct(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswpunct->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswpunct.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswpunct_l = dll.NewProc("_iswpunct_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswpunct_l(wint_t _C,_locale_t _Locale);
func X_iswpunct_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswpunct_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswpunct_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswalnum = dll.NewProc("iswalnum")

// int __attribute__((__cdecl__)) iswalnum(wint_t _C);
func Xiswalnum(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswalnum->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswalnum.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswalnum_l = dll.NewProc("_iswalnum_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswalnum_l(wint_t _C,_locale_t _Locale);
func X_iswalnum_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswalnum_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswalnum_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswprint = dll.NewProc("iswprint")

// int __attribute__((__cdecl__)) iswprint(wint_t _C);
func Xiswprint(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswprint->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswprint.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswprint_l = dll.NewProc("_iswprint_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswprint_l(wint_t _C,_locale_t _Locale);
func X_iswprint_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswprint_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswprint_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswgraph = dll.NewProc("iswgraph")

// int __attribute__((__cdecl__)) iswgraph(wint_t _C);
func Xiswgraph(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswgraph->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswgraph.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswgraph_l = dll.NewProc("_iswgraph_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswgraph_l(wint_t _C,_locale_t _Locale);
func X_iswgraph_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswgraph_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswgraph_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswcntrl = dll.NewProc("iswcntrl")

// int __attribute__((__cdecl__)) iswcntrl(wint_t _C);
func Xiswcntrl(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswcntrl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswcntrl.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswcntrl_l = dll.NewProc("_iswcntrl_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswcntrl_l(wint_t _C,_locale_t _Locale);
func X_iswcntrl_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswcntrl_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswcntrl_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswascii = dll.NewProc("iswascii")

// int __attribute__((__cdecl__)) iswascii(wint_t _C);
func Xiswascii(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswascii->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswascii.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procisleadbyte = dll.NewProc("isleadbyte")

// int __attribute__((__cdecl__)) isleadbyte(int _C);
func Xisleadbyte(tls *TLS, __C int32) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xisleadbyte->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procisleadbyte.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isleadbyte_l = dll.NewProc("_isleadbyte_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isleadbyte_l(int _C,_locale_t _Locale);
func X_isleadbyte_l(tls *TLS, __C int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_isleadbyte_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isleadbyte_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proctowupper = dll.NewProc("towupper")

// wint_t __attribute__((__cdecl__)) towupper(wint_t _C);
func Xtowupper(tls *TLS, __C Twint_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xtowupper->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctowupper.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_towupper_l = dll.NewProc("_towupper_l")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _towupper_l(wint_t _C,_locale_t _Locale);
func X_towupper_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_towupper_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_towupper_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proctowlower = dll.NewProc("towlower")

// wint_t __attribute__((__cdecl__)) towlower(wint_t _C);
func Xtowlower(tls *TLS, __C Twint_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xtowlower->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctowlower.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_towlower_l = dll.NewProc("_towlower_l")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _towlower_l(wint_t _C,_locale_t _Locale);
func X_towlower_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_towlower_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_towlower_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var prociswctype = dll.NewProc("iswctype")

// int __attribute__((__cdecl__)) iswctype(wint_t _C,wctype_t _Type);
func Xiswctype(tls *TLS, __C Twint_t, __Type Twctype_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Type=%+v", __C, __Type)
		defer func() { trc(`Xiswctype->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswctype.Addr(), uintptr(__C), uintptr(__Type))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswctype_l = dll.NewProc("_iswctype_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswctype_l(wint_t _C,wctype_t _Type,_locale_t _Locale);
func X_iswctype_l(tls *TLS, __C Twint_t, __Type Twctype_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Type=%+v _Locale=%+v", __C, __Type, __Locale)
		defer func() { trc(`X_iswctype_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswctype_l.Addr(), uintptr(__C), uintptr(__Type), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__iswcsymf = dll.NewProc("__iswcsymf")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __iswcsymf(wint_t _C);
func X__iswcsymf(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__iswcsymf->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__iswcsymf.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswcsymf_l = dll.NewProc("_iswcsymf_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswcsymf_l(wint_t _C,_locale_t _Locale);
func X_iswcsymf_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswcsymf_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswcsymf_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc__iswcsym = dll.NewProc("__iswcsym")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) __iswcsym(wint_t _C);
func X__iswcsym(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`X__iswcsym->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__iswcsym.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_iswcsym_l = dll.NewProc("_iswcsym_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _iswcsym_l(wint_t _C,_locale_t _Locale);
func X_iswcsym_l(tls *TLS, __C Twint_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Locale=%+v", __C, __Locale)
		defer func() { trc(`X_iswcsym_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_iswcsym_l.Addr(), uintptr(__C), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procis_wctype = dll.NewProc("is_wctype")

// int __attribute__((__cdecl__)) is_wctype(wint_t _C,wctype_t _Type);
func Xis_wctype(tls *TLS, __C Twint_t, __Type Twctype_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v _Type=%+v", __C, __Type)
		defer func() { trc(`Xis_wctype->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procis_wctype.Addr(), uintptr(__C), uintptr(__Type))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var prociswblank = dll.NewProc("iswblank")

// int __attribute__((__cdecl__)) iswblank(wint_t _C);
func Xiswblank(tls *TLS, __C Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_C=%+v", __C)
		defer func() { trc(`Xiswblank->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(prociswblank.Addr(), uintptr(__C))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc___mb_cur_max_func = dll.NewProc("___mb_cur_max_func")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) ___mb_cur_max_func(void);
func X___mb_cur_max_func(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X___mb_cur_max_func->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc___mb_cur_max_func.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

type T_exception = struct {
	Ftype1  int32
	Fname   uintptr
	Farg1   float64
	Farg2   float64
	Fretval float64
}

type T__mingw_dbl_type_t = struct {
	Fval [0]uint64
	Flh  [0]struct {
		Flow  uint32
		Fhigh uint32
	}
	Fx float64
}

type T__mingw_flt_type_t = struct {
	Fval [0]uint32
	Fx   float32
}

type T__mingw_ldbl_type_t = struct {
	Flh [0]struct {
		Flow      uint32
		Fhigh     uint32
		F__ccgo8  uint32
		F__ccgo12 uint32
	}
	Fx           float64
	F__ccgo_pad2 [8]byte
}

var proc__setusermatherr = dll.NewProc("__setusermatherr")

// __attribute__ ((__dllimport__)) void __setusermatherr(int ( *)(struct _exception *));
func X__setusermatherr(tls *TLS, _0 uintptr) {
	X__ccgo_SyscallFP()
	panic(663)
}

var procabs = dll.NewProc("abs")

// int __attribute__((__cdecl__)) abs(int _X);
func Xabs(tls *TLS, __X int32) (r int32) {
	if __ccgo_strace {
		trc("_X=%+v", __X)
		defer func() { trc(`Xabs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procabs.Addr(), uintptr(__X))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proclabs = dll.NewProc("labs")

// long __attribute__((__cdecl__)) labs(long _X);
func Xlabs(tls *TLS, __X int32) (r int32) {
	if __ccgo_strace {
		trc("_X=%+v", __X)
		defer func() { trc(`Xlabs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proclabs.Addr(), uintptr(__X))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

type T_complex = struct {
	Fx float64
	Fy float64
}

type Tfloat_t = float32

type Tdouble_t = float64

type T_iobuf = struct {
	F_Placeholder uintptr
}

type TFILE = struct {
	F_Placeholder uintptr
}

type T_off_t = int32

type Toff32_t = int32

type T_off64_t = int64

type Toff64_t = int64

type Toff_t = int32

type Tfpos_t = int64

var proc_fsopen = dll.NewProc("_fsopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _fsopen(const char *_Filename,const char *_Mode,int _ShFlag);
func X_fsopen(tls *TLS, __Filename uintptr, __Mode uintptr, __ShFlag int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v _ShFlag=%+v", __Filename, __Mode, __ShFlag)
		defer func() { trc(`X_fsopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fsopen.Addr(), __Filename, __Mode, uintptr(__ShFlag))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procclearerr = dll.NewProc("clearerr")

// void __attribute__((__cdecl__)) clearerr(FILE *_File);
func Xclearerr(tls *TLS, __File uintptr) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
	}
	r0, r1, err := syscall.SyscallN(procclearerr.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var procfclose = dll.NewProc("fclose")

// int __attribute__((__cdecl__)) fclose(FILE *_File);
func Xfclose(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xfclose->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfclose.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fcloseall = dll.NewProc("_fcloseall")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fcloseall(void);
func X_fcloseall(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_fcloseall->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fcloseall.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fdopen = dll.NewProc("_fdopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _fdopen(int _FileHandle,const char *_Mode);
func X_fdopen(tls *TLS, __FileHandle int32, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Mode=%+v", __FileHandle, __Mode)
		defer func() { trc(`X_fdopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fdopen.Addr(), uintptr(__FileHandle), __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfeof = dll.NewProc("feof")

// int __attribute__((__cdecl__)) feof(FILE *_File);
func Xfeof(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xfeof->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfeof.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procferror = dll.NewProc("ferror")

// int __attribute__((__cdecl__)) ferror(FILE *_File);
func Xferror(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xferror->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procferror.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfflush = dll.NewProc("fflush")

// int __attribute__((__cdecl__)) fflush(FILE *_File);
func Xfflush(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xfflush->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfflush.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfgetc = dll.NewProc("fgetc")

// int __attribute__((__cdecl__)) fgetc(FILE *_File);
func Xfgetc(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xfgetc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfgetc.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fgetchar = dll.NewProc("_fgetchar")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fgetchar(void);
func X_fgetchar(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_fgetchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fgetchar.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfgetpos = dll.NewProc("fgetpos")

// int __attribute__((__cdecl__)) fgetpos(FILE * __restrict__ _File ,fpos_t * __restrict__ _Pos);
func Xfgetpos(tls *TLS, __File uintptr, __Pos uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Pos=%+v", __File, __Pos)
		defer func() { trc(`Xfgetpos->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfgetpos.Addr(), __File, __Pos)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfgets = dll.NewProc("fgets")

// char * __attribute__((__cdecl__)) fgets(char * __restrict__ _Buf,int _MaxCount,FILE * __restrict__ _File);
func Xfgets(tls *TLS, __Buf uintptr, __MaxCount int32, __File uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buf=%+v _MaxCount=%+v _File=%+v", __Buf, __MaxCount, __File)
		defer func() { trc(`Xfgets->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfgets.Addr(), __Buf, uintptr(__MaxCount), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_fileno = dll.NewProc("_fileno")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fileno(FILE *_File);
func X_fileno(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_fileno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fileno.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_tempnam = dll.NewProc("_tempnam")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _tempnam(const char *_DirName,const char *_FilePrefix);
func X_tempnam(tls *TLS, __DirName uintptr, __FilePrefix uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_DirName=%+v _FilePrefix=%+v", __DirName, __FilePrefix)
		defer func() { trc(`X_tempnam->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_tempnam.Addr(), __DirName, __FilePrefix)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_flushall = dll.NewProc("_flushall")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _flushall(void);
func X_flushall(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_flushall->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_flushall.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfopen = dll.NewProc("fopen")

// FILE * __attribute__((__cdecl__)) fopen(const char * __restrict__ _Filename,const char * __restrict__ _Mode);
func Xfopen(tls *TLS, __Filename uintptr, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v", __Filename, __Mode)
		defer func() { trc(`Xfopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfopen.Addr(), __Filename, __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfputc = dll.NewProc("fputc")

// int __attribute__((__cdecl__)) fputc(int _Ch,FILE *_File);
func Xfputc(tls *TLS, __Ch int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xfputc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfputc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fputchar = dll.NewProc("_fputchar")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fputchar(int _Ch);
func X_fputchar(tls *TLS, __Ch int32) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v", __Ch)
		defer func() { trc(`X_fputchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fputchar.Addr(), uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfputs = dll.NewProc("fputs")

// int __attribute__((__cdecl__)) fputs(const char * __restrict__ _Str,FILE * __restrict__ _File);
func Xfputs(tls *TLS, __Str uintptr, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _File=%+v", __Str, __File)
		defer func() { trc(`Xfputs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfputs.Addr(), __Str, __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfread = dll.NewProc("fread")

// size_t __attribute__((__cdecl__)) fread(void * __restrict__ _DstBuf,size_t _ElementSize,size_t _Count,FILE * __restrict__ _File);
func Xfread(tls *TLS, __DstBuf uintptr, __ElementSize Tsize_t, __Count Tsize_t, __File uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_DstBuf=%+v _ElementSize=%+v _Count=%+v _File=%+v", __DstBuf, __ElementSize, __Count, __File)
		defer func() { trc(`Xfread->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfread.Addr(), __DstBuf, uintptr(__ElementSize), uintptr(__Count), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procfreopen = dll.NewProc("freopen")

// FILE * __attribute__((__cdecl__)) freopen(const char * __restrict__ _Filename,const char * __restrict__ _Mode,FILE * __restrict__ _File);
func Xfreopen(tls *TLS, __Filename uintptr, __Mode uintptr, __File uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v _File=%+v", __Filename, __Mode, __File)
		defer func() { trc(`Xfreopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfreopen.Addr(), __Filename, __Mode, __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfsetpos = dll.NewProc("fsetpos")

// int __attribute__((__cdecl__)) fsetpos(FILE *_File,const fpos_t *_Pos);
func Xfsetpos(tls *TLS, __File uintptr, __Pos uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Pos=%+v", __File, __Pos)
		defer func() { trc(`Xfsetpos->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfsetpos.Addr(), __File, __Pos)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procfseek = dll.NewProc("fseek")

// int __attribute__((__cdecl__)) fseek(FILE *_File,long _Offset,int _Origin);
func Xfseek(tls *TLS, __File uintptr, __Offset int32, __Origin int32) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Offset=%+v _Origin=%+v", __File, __Offset, __Origin)
		defer func() { trc(`Xfseek->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfseek.Addr(), __File, uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procftell = dll.NewProc("ftell")

// long __attribute__((__cdecl__)) ftell(FILE *_File);
func Xftell(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xftell->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procftell.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fseeki64 = dll.NewProc("_fseeki64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fseeki64(FILE *_File, long long _Offset,int _Origin);
func X_fseeki64(tls *TLS, __File uintptr, __Offset int64, __Origin int32) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Offset=%+v _Origin=%+v", __File, __Offset, __Origin)
		defer func() { trc(`X_fseeki64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fseeki64.Addr(), __File, uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_ftelli64 = dll.NewProc("_ftelli64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _ftelli64(FILE *_File);
func X_ftelli64(tls *TLS, __File uintptr) (r int64) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_ftelli64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ftelli64.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var procfwrite = dll.NewProc("fwrite")

// size_t __attribute__((__cdecl__)) fwrite(const void * __restrict__ _Str,size_t _Size,size_t _Count,FILE * __restrict__ _File);
func Xfwrite(tls *TLS, __Str uintptr, __Size Tsize_t, __Count Tsize_t, __File uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Size=%+v _Count=%+v _File=%+v", __Str, __Size, __Count, __File)
		defer func() { trc(`Xfwrite->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfwrite.Addr(), __Str, uintptr(__Size), uintptr(__Count), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procgetc = dll.NewProc("getc")

// int __attribute__((__cdecl__)) getc(FILE *_File);
func Xgetc(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xgetc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgetc.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procgetchar = dll.NewProc("getchar")

// int __attribute__((__cdecl__)) getchar(void);
func Xgetchar(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xgetchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgetchar.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_getmaxstdio = dll.NewProc("_getmaxstdio")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _getmaxstdio(void);
func X_getmaxstdio(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getmaxstdio->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getmaxstdio.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procgets = dll.NewProc("gets")

// char * __attribute__((__cdecl__)) gets(char *_Buffer);
func Xgets(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`Xgets->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgets.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_getw = dll.NewProc("_getw")

// int __attribute__((__cdecl__)) _getw(FILE *_File);
func X_getw(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_getw->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getw.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procperror = dll.NewProc("perror")

// void __attribute__((__cdecl__)) perror(const char *_ErrMsg);
func Xperror(tls *TLS, __ErrMsg uintptr) {
	if __ccgo_strace {
		trc("_ErrMsg=%+v", __ErrMsg)
	}
	r0, r1, err := syscall.SyscallN(procperror.Addr(), __ErrMsg)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_pclose = dll.NewProc("_pclose")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _pclose(FILE *_File);
func X_pclose(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_pclose->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_pclose.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_popen = dll.NewProc("_popen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _popen(const char *_Command,const char *_Mode);
func X_popen(tls *TLS, __Command uintptr, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Command=%+v _Mode=%+v", __Command, __Mode)
		defer func() { trc(`X_popen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_popen.Addr(), __Command, __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procputc = dll.NewProc("putc")

// int __attribute__((__cdecl__)) putc(int _Ch,FILE *_File);
func Xputc(tls *TLS, __Ch int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xputc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procputc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procputchar = dll.NewProc("putchar")

// int __attribute__((__cdecl__)) putchar(int _Ch);
func Xputchar(tls *TLS, __Ch int32) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v", __Ch)
		defer func() { trc(`Xputchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procputchar.Addr(), uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procputs = dll.NewProc("puts")

// int __attribute__((__cdecl__)) puts(const char *_Str);
func Xputs(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`Xputs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procputs.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_putw = dll.NewProc("_putw")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _putw(int _Word,FILE *_File);
func X_putw(tls *TLS, __Word int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Word=%+v _File=%+v", __Word, __File)
		defer func() { trc(`X_putw->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putw.Addr(), uintptr(__Word), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procremove = dll.NewProc("remove")

// int __attribute__((__cdecl__)) remove(const char *_Filename);
func Xremove(tls *TLS, __Filename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v", __Filename)
		defer func() { trc(`Xremove->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procremove.Addr(), __Filename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procrename = dll.NewProc("rename")

// int __attribute__((__cdecl__)) rename(const char *_OldFilename,const char *_NewFilename);
func Xrename(tls *TLS, __OldFilename uintptr, __NewFilename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_OldFilename=%+v _NewFilename=%+v", __OldFilename, __NewFilename)
		defer func() { trc(`Xrename->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procrename.Addr(), __OldFilename, __NewFilename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_unlink = dll.NewProc("_unlink")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _unlink(const char *_Filename);
func X_unlink(tls *TLS, __Filename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v", __Filename)
		defer func() { trc(`X_unlink->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_unlink.Addr(), __Filename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procrewind = dll.NewProc("rewind")

// void __attribute__((__cdecl__)) rewind(FILE *_File);
func Xrewind(tls *TLS, __File uintptr) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
	}
	r0, r1, err := syscall.SyscallN(procrewind.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_rmtmp = dll.NewProc("_rmtmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _rmtmp(void);
func X_rmtmp(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_rmtmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_rmtmp.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procsetbuf = dll.NewProc("setbuf")

// void __attribute__((__cdecl__)) setbuf(FILE * __restrict__ _File,char * __restrict__ _Buffer);
func Xsetbuf(tls *TLS, __File uintptr, __Buffer uintptr) {
	if __ccgo_strace {
		trc("_File=%+v _Buffer=%+v", __File, __Buffer)
	}
	r0, r1, err := syscall.SyscallN(procsetbuf.Addr(), __File, __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_setmaxstdio = dll.NewProc("_setmaxstdio")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _setmaxstdio(int _Max);
func X_setmaxstdio(tls *TLS, __Max int32) (r int32) {
	if __ccgo_strace {
		trc("_Max=%+v", __Max)
		defer func() { trc(`X_setmaxstdio->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_setmaxstdio.Addr(), uintptr(__Max))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procsetvbuf = dll.NewProc("setvbuf")

// int __attribute__((__cdecl__)) setvbuf(FILE * __restrict__ _File,char * __restrict__ _Buf,int _Mode,size_t _Size);
func Xsetvbuf(tls *TLS, __File uintptr, __Buf uintptr, __Mode int32, __Size Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Buf=%+v _Mode=%+v _Size=%+v", __File, __Buf, __Mode, __Size)
		defer func() { trc(`Xsetvbuf->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procsetvbuf.Addr(), __File, __Buf, uintptr(__Mode), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proctmpfile = dll.NewProc("tmpfile")

// FILE * __attribute__((__cdecl__)) tmpfile(void);
func Xtmpfile(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xtmpfile->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctmpfile.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proctmpnam = dll.NewProc("tmpnam")

// char * __attribute__((__cdecl__)) tmpnam(char *_Buffer);
func Xtmpnam(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`Xtmpnam->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proctmpnam.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procungetc = dll.NewProc("ungetc")

// int __attribute__((__cdecl__)) ungetc(int _Ch,FILE *_File);
func Xungetc(tls *TLS, __Ch int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xungetc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procungetc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_set_printf_count_output = dll.NewProc("_set_printf_count_output")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _set_printf_count_output(int _Value);
func X_set_printf_count_output(tls *TLS, __Value int32) (r int32) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_set_printf_count_output->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_printf_count_output.Addr(), uintptr(__Value))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_get_printf_count_output = dll.NewProc("_get_printf_count_output")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _get_printf_count_output(void);
func X_get_printf_count_output(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_get_printf_count_output->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_printf_count_output.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wfsopen = dll.NewProc("_wfsopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _wfsopen(const wchar_t *_Filename,const wchar_t *_Mode,int _ShFlag);
func X_wfsopen(tls *TLS, __Filename uintptr, __Mode uintptr, __ShFlag int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v _ShFlag=%+v", __Filename, __Mode, __ShFlag)
		defer func() { trc(`X_wfsopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfsopen.Addr(), __Filename, __Mode, uintptr(__ShFlag))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfgetwc = dll.NewProc("fgetwc")

// wint_t __attribute__((__cdecl__)) fgetwc(FILE *_File);
func Xfgetwc(tls *TLS, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xfgetwc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfgetwc.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_fgetwchar = dll.NewProc("_fgetwchar")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _fgetwchar(void);
func X_fgetwchar(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_fgetwchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fgetwchar.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procfputwc = dll.NewProc("fputwc")

// wint_t __attribute__((__cdecl__)) fputwc(wchar_t _Ch,FILE *_File);
func Xfputwc(tls *TLS, __Ch Twchar_t, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xfputwc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfputwc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_fputwchar = dll.NewProc("_fputwchar")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _fputwchar(wchar_t _Ch);
func X_fputwchar(tls *TLS, __Ch Twchar_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v", __Ch)
		defer func() { trc(`X_fputwchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fputwchar.Addr(), uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procgetwc = dll.NewProc("getwc")

// wint_t __attribute__((__cdecl__)) getwc(FILE *_File);
func Xgetwc(tls *TLS, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`Xgetwc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgetwc.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procgetwchar = dll.NewProc("getwchar")

// wint_t __attribute__((__cdecl__)) getwchar(void);
func Xgetwchar(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xgetwchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgetwchar.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procputwc = dll.NewProc("putwc")

// wint_t __attribute__((__cdecl__)) putwc(wchar_t _Ch,FILE *_File);
func Xputwc(tls *TLS, __Ch Twchar_t, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xputwc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procputwc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procputwchar = dll.NewProc("putwchar")

// wint_t __attribute__((__cdecl__)) putwchar(wchar_t _Ch);
func Xputwchar(tls *TLS, __Ch Twchar_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v", __Ch)
		defer func() { trc(`Xputwchar->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procputwchar.Addr(), uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procungetwc = dll.NewProc("ungetwc")

// wint_t __attribute__((__cdecl__)) ungetwc(wint_t _Ch,FILE *_File);
func Xungetwc(tls *TLS, __Ch Twint_t, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`Xungetwc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procungetwc.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procfgetws = dll.NewProc("fgetws")

// wchar_t * __attribute__((__cdecl__)) fgetws(wchar_t * __restrict__ _Dst,int _SizeInWords,FILE * __restrict__ _File);
func Xfgetws(tls *TLS, __Dst uintptr, __SizeInWords int32, __File uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Dst=%+v _SizeInWords=%+v _File=%+v", __Dst, __SizeInWords, __File)
		defer func() { trc(`Xfgetws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfgetws.Addr(), __Dst, uintptr(__SizeInWords), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfputws = dll.NewProc("fputws")

// int __attribute__((__cdecl__)) fputws(const wchar_t * __restrict__ _Str,FILE * __restrict__ _File);
func Xfputws(tls *TLS, __Str uintptr, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _File=%+v", __Str, __File)
		defer func() { trc(`Xfputws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procfputws.Addr(), __Str, __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_getws = dll.NewProc("_getws")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _getws(wchar_t *_String);
func X_getws(tls *TLS, __String uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_getws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getws.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_putws = dll.NewProc("_putws")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _putws(const wchar_t *_Str);
func X_putws(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_putws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putws.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtempnam = dll.NewProc("_wtempnam")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wtempnam(const wchar_t *_Directory,const wchar_t *_FilePrefix);
func X_wtempnam(tls *TLS, __Directory uintptr, __FilePrefix uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Directory=%+v _FilePrefix=%+v", __Directory, __FilePrefix)
		defer func() { trc(`X_wtempnam->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtempnam.Addr(), __Directory, __FilePrefix)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wfdopen = dll.NewProc("_wfdopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _wfdopen(int _FileHandle ,const wchar_t *_Mode);
func X_wfdopen(tls *TLS, __FileHandle int32, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Mode=%+v", __FileHandle, __Mode)
		defer func() { trc(`X_wfdopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfdopen.Addr(), uintptr(__FileHandle), __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wfopen = dll.NewProc("_wfopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _wfopen(const wchar_t * __restrict__ _Filename,const wchar_t *__restrict__ _Mode);
func X_wfopen(tls *TLS, __Filename uintptr, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v", __Filename, __Mode)
		defer func() { trc(`X_wfopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfopen.Addr(), __Filename, __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wfreopen = dll.NewProc("_wfreopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _wfreopen(const wchar_t * __restrict__ _Filename,const wchar_t * __restrict__ _Mode,FILE * __restrict__ _OldFile);
func X_wfreopen(tls *TLS, __Filename uintptr, __Mode uintptr, __OldFile uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v _OldFile=%+v", __Filename, __Mode, __OldFile)
		defer func() { trc(`X_wfreopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfreopen.Addr(), __Filename, __Mode, __OldFile)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wperror = dll.NewProc("_wperror")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _wperror(const wchar_t *_ErrMsg);
func X_wperror(tls *TLS, __ErrMsg uintptr) {
	if __ccgo_strace {
		trc("_ErrMsg=%+v", __ErrMsg)
	}
	r0, r1, err := syscall.SyscallN(proc_wperror.Addr(), __ErrMsg)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_wpopen = dll.NewProc("_wpopen")

// __attribute__ ((__dllimport__)) FILE * __attribute__((__cdecl__)) _wpopen(const wchar_t *_Command,const wchar_t *_Mode);
func X_wpopen(tls *TLS, __Command uintptr, __Mode uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Command=%+v _Mode=%+v", __Command, __Mode)
		defer func() { trc(`X_wpopen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wpopen.Addr(), __Command, __Mode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wremove = dll.NewProc("_wremove")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wremove(const wchar_t *_Filename);
func X_wremove(tls *TLS, __Filename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v", __Filename)
		defer func() { trc(`X_wremove->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wremove.Addr(), __Filename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtmpnam = dll.NewProc("_wtmpnam")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wtmpnam(wchar_t *_Buffer);
func X_wtmpnam(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_wtmpnam->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtmpnam.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_fgetwc_nolock = dll.NewProc("_fgetwc_nolock")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _fgetwc_nolock(FILE *_File);
func X_fgetwc_nolock(tls *TLS, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_fgetwc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fgetwc_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_fputwc_nolock = dll.NewProc("_fputwc_nolock")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _fputwc_nolock(wchar_t _Ch,FILE *_File);
func X_fputwc_nolock(tls *TLS, __Ch Twchar_t, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`X_fputwc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fputwc_nolock.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_ungetwc_nolock = dll.NewProc("_ungetwc_nolock")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _ungetwc_nolock(wint_t _Ch,FILE *_File);
func X_ungetwc_nolock(tls *TLS, __Ch Twint_t, __File uintptr) (r Twint_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`X_ungetwc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ungetwc_nolock.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_fgetc_nolock = dll.NewProc("_fgetc_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fgetc_nolock(FILE *_File);
func X_fgetc_nolock(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_fgetc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fgetc_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fputc_nolock = dll.NewProc("_fputc_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fputc_nolock(int _Char, FILE *_File);
func X_fputc_nolock(tls *TLS, __Char int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Char=%+v _File=%+v", __Char, __File)
		defer func() { trc(`X_fputc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fputc_nolock.Addr(), uintptr(__Char), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_getc_nolock = dll.NewProc("_getc_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _getc_nolock(FILE *_File);
func X_getc_nolock(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_getc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getc_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_putc_nolock = dll.NewProc("_putc_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _putc_nolock(int _Char, FILE *_File);
func X_putc_nolock(tls *TLS, __Char int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Char=%+v _File=%+v", __Char, __File)
		defer func() { trc(`X_putc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putc_nolock.Addr(), uintptr(__Char), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_lock_file = dll.NewProc("_lock_file")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _lock_file(FILE *_File);
func X_lock_file(tls *TLS, __File uintptr) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
	}
	r0, r1, err := syscall.SyscallN(proc_lock_file.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_unlock_file = dll.NewProc("_unlock_file")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _unlock_file(FILE *_File);
func X_unlock_file(tls *TLS, __File uintptr) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
	}
	r0, r1, err := syscall.SyscallN(proc_unlock_file.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_fclose_nolock = dll.NewProc("_fclose_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fclose_nolock(FILE *_File);
func X_fclose_nolock(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_fclose_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fclose_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fflush_nolock = dll.NewProc("_fflush_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fflush_nolock(FILE *_File);
func X_fflush_nolock(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_fflush_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fflush_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fread_nolock = dll.NewProc("_fread_nolock")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _fread_nolock(void * __restrict__ _DstBuf,size_t _ElementSize,size_t _Count,FILE * __restrict__ _File);
func X_fread_nolock(tls *TLS, __DstBuf uintptr, __ElementSize Tsize_t, __Count Tsize_t, __File uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_DstBuf=%+v _ElementSize=%+v _Count=%+v _File=%+v", __DstBuf, __ElementSize, __Count, __File)
		defer func() { trc(`X_fread_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fread_nolock.Addr(), __DstBuf, uintptr(__ElementSize), uintptr(__Count), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_fseek_nolock = dll.NewProc("_fseek_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fseek_nolock(FILE *_File,long _Offset,int _Origin);
func X_fseek_nolock(tls *TLS, __File uintptr, __Offset int32, __Origin int32) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Offset=%+v _Origin=%+v", __File, __Offset, __Origin)
		defer func() { trc(`X_fseek_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fseek_nolock.Addr(), __File, uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_ftell_nolock = dll.NewProc("_ftell_nolock")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _ftell_nolock(FILE *_File);
func X_ftell_nolock(tls *TLS, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_ftell_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ftell_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fseeki64_nolock = dll.NewProc("_fseeki64_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _fseeki64_nolock(FILE *_File, long long _Offset,int _Origin);
func X_fseeki64_nolock(tls *TLS, __File uintptr, __Offset int64, __Origin int32) (r int32) {
	if __ccgo_strace {
		trc("_File=%+v _Offset=%+v _Origin=%+v", __File, __Offset, __Origin)
		defer func() { trc(`X_fseeki64_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fseeki64_nolock.Addr(), __File, uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_ftelli64_nolock = dll.NewProc("_ftelli64_nolock")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _ftelli64_nolock(FILE *_File);
func X_ftelli64_nolock(tls *TLS, __File uintptr) (r int64) {
	if __ccgo_strace {
		trc("_File=%+v", __File)
		defer func() { trc(`X_ftelli64_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ftelli64_nolock.Addr(), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_fwrite_nolock = dll.NewProc("_fwrite_nolock")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _fwrite_nolock(const void * __restrict__ _DstBuf,size_t _Size,size_t _Count,FILE * __restrict__ _File);
func X_fwrite_nolock(tls *TLS, __DstBuf uintptr, __Size Tsize_t, __Count Tsize_t, __File uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_DstBuf=%+v _Size=%+v _Count=%+v _File=%+v", __DstBuf, __Size, __Count, __File)
		defer func() { trc(`X_fwrite_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fwrite_nolock.Addr(), __DstBuf, uintptr(__Size), uintptr(__Count), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_ungetc_nolock = dll.NewProc("_ungetc_nolock")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _ungetc_nolock(int _Ch,FILE *_File);
func X_ungetc_nolock(tls *TLS, __Ch int32, __File uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _File=%+v", __Ch, __File)
		defer func() { trc(`X_ungetc_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ungetc_nolock.Addr(), uintptr(__Ch), __File)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wspawnv = dll.NewProc("_wspawnv")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wspawnv(int _Mode,const wchar_t *_Filename,const wchar_t *const *_ArgList);
func X_wspawnv(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v", __Mode, __Filename, __ArgList)
		defer func() { trc(`X_wspawnv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wspawnv.Addr(), uintptr(__Mode), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wspawnve = dll.NewProc("_wspawnve")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wspawnve(int _Mode,const wchar_t *_Filename,const wchar_t *const *_ArgList,const wchar_t *const *_Env);
func X_wspawnve(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v _Env=%+v", __Mode, __Filename, __ArgList, __Env)
		defer func() { trc(`X_wspawnve->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wspawnve.Addr(), uintptr(__Mode), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wspawnvp = dll.NewProc("_wspawnvp")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wspawnvp(int _Mode,const wchar_t *_Filename,const wchar_t *const *_ArgList);
func X_wspawnvp(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v", __Mode, __Filename, __ArgList)
		defer func() { trc(`X_wspawnvp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wspawnvp.Addr(), uintptr(__Mode), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wspawnvpe = dll.NewProc("_wspawnvpe")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wspawnvpe(int _Mode,const wchar_t *_Filename,const wchar_t *const *_ArgList,const wchar_t *const *_Env);
func X_wspawnvpe(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v _Env=%+v", __Mode, __Filename, __ArgList, __Env)
		defer func() { trc(`X_wspawnvpe->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wspawnvpe.Addr(), uintptr(__Mode), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_spawnv = dll.NewProc("_spawnv")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _spawnv(int _Mode,const char *_Filename,const char *const *_ArgList);
func X_spawnv(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v", __Mode, __Filename, __ArgList)
		defer func() { trc(`X_spawnv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_spawnv.Addr(), uintptr(__Mode), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_spawnve = dll.NewProc("_spawnve")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _spawnve(int _Mode,const char *_Filename,const char *const *_ArgList,const char *const *_Env);
func X_spawnve(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v _Env=%+v", __Mode, __Filename, __ArgList, __Env)
		defer func() { trc(`X_spawnve->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_spawnve.Addr(), uintptr(__Mode), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_spawnvp = dll.NewProc("_spawnvp")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _spawnvp(int _Mode,const char *_Filename,const char *const *_ArgList);
func X_spawnvp(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v", __Mode, __Filename, __ArgList)
		defer func() { trc(`X_spawnvp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_spawnvp.Addr(), uintptr(__Mode), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_spawnvpe = dll.NewProc("_spawnvpe")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _spawnvpe(int _Mode,const char *_Filename,const char *const *_ArgList,const char *const *_Env);
func X_spawnvpe(tls *TLS, __Mode int32, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Mode=%+v _Filename=%+v _ArgList=%+v _Env=%+v", __Mode, __Filename, __ArgList, __Env)
		defer func() { trc(`X_spawnvpe->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_spawnvpe.Addr(), uintptr(__Mode), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

type T_onexit_t = uintptr

type Tdiv_t = struct {
	Fquot int32
	Frem  int32
}

type T_div_t = Tdiv_t

type Tldiv_t = struct {
	Fquot int32
	Frem  int32
}

type T_ldiv_t = Tldiv_t

type T_LDOUBLE = struct {
	Fld [10]uint8
}

type T_CRT_DOUBLE = struct {
	Fx float64
}

type T_CRT_FLOAT = struct {
	Ff float32
}

type T_LONGDOUBLE = struct {
	Fx float64
}

type T_LDBL12 = struct {
	Fld12 [12]uint8
}

type T_purecall_handler = uintptr

var proc_set_purecall_handler = dll.NewProc("_set_purecall_handler")

// __attribute__ ((__dllimport__)) _purecall_handler __attribute__((__cdecl__)) _set_purecall_handler(_purecall_handler _Handler);
func X_set_purecall_handler(tls *TLS, __Handler T_purecall_handler) (r T_purecall_handler) {
	X__ccgo_SyscallFP()
	panic(663)
}

var proc_get_purecall_handler = dll.NewProc("_get_purecall_handler")

// __attribute__ ((__dllimport__)) _purecall_handler __attribute__((__cdecl__)) _get_purecall_handler(void);
func X_get_purecall_handler(tls *TLS) (r T_purecall_handler) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_get_purecall_handler->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_purecall_handler.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T_purecall_handler(r0)
}

type T_invalid_parameter_handler = uintptr

var proc_set_invalid_parameter_handler = dll.NewProc("_set_invalid_parameter_handler")

// __attribute__ ((__dllimport__)) _invalid_parameter_handler __attribute__((__cdecl__)) _set_invalid_parameter_handler(_invalid_parameter_handler _Handler);
func X_set_invalid_parameter_handler(tls *TLS, __Handler T_invalid_parameter_handler) (r T_invalid_parameter_handler) {
	X__ccgo_SyscallFP()
	panic(663)
}

var proc_get_invalid_parameter_handler = dll.NewProc("_get_invalid_parameter_handler")

// __attribute__ ((__dllimport__)) _invalid_parameter_handler __attribute__((__cdecl__)) _get_invalid_parameter_handler(void);
func X_get_invalid_parameter_handler(tls *TLS) (r T_invalid_parameter_handler) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_get_invalid_parameter_handler->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_invalid_parameter_handler.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T_invalid_parameter_handler(r0)
}

var proc_set_errno = dll.NewProc("_set_errno")

// errno_t __attribute__((__cdecl__)) _set_errno(int _Value);
func X_set_errno(tls *TLS, __Value int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_set_errno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_errno.Addr(), uintptr(__Value))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_errno = dll.NewProc("_get_errno")

// errno_t __attribute__((__cdecl__)) _get_errno(int *_Value);
func X_get_errno(tls *TLS, __Value uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_get_errno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_errno.Addr(), __Value)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc__doserrno = dll.NewProc("__doserrno")

// __attribute__ ((__dllimport__)) unsigned long * __attribute__((__cdecl__)) __doserrno(void);
func X__doserrno(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__doserrno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__doserrno.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_set_doserrno = dll.NewProc("_set_doserrno")

// errno_t __attribute__((__cdecl__)) _set_doserrno(unsigned long _Value);
func X_set_doserrno(tls *TLS, __Value uint32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_set_doserrno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_doserrno.Addr(), uintptr(__Value))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_doserrno = dll.NewProc("_get_doserrno")

// errno_t __attribute__((__cdecl__)) _get_doserrno(unsigned long *_Value);
func X_get_doserrno(tls *TLS, __Value uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_get_doserrno->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_doserrno.Addr(), __Value)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc__sys_errlist = dll.NewProc("__sys_errlist")

// __attribute__ ((__dllimport__)) char ** __attribute__((__cdecl__)) __sys_errlist(void);
func X__sys_errlist(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__sys_errlist->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__sys_errlist.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__sys_nerr = dll.NewProc("__sys_nerr")

// __attribute__ ((__dllimport__)) int * __attribute__((__cdecl__)) __sys_nerr(void);
func X__sys_nerr(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__sys_nerr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__sys_nerr.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p___argv = dll.NewProc("__p___argv")

// __attribute__ ((__dllimport__)) char *** __attribute__((__cdecl__)) __p___argv(void);
func X__p___argv(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p___argv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p___argv.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p__fmode = dll.NewProc("__p__fmode")

// __attribute__ ((__dllimport__)) int * __attribute__((__cdecl__)) __p__fmode(void);
func X__p__fmode(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p__fmode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p__fmode.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p___argc = dll.NewProc("__p___argc")

// __attribute__ ((__dllimport__)) int * __attribute__((__cdecl__)) __p___argc(void);
func X__p___argc(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p___argc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p___argc.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p___wargv = dll.NewProc("__p___wargv")

// __attribute__ ((__dllimport__)) wchar_t *** __attribute__((__cdecl__)) __p___wargv(void);
func X__p___wargv(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p___wargv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p___wargv.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p__environ = dll.NewProc("__p__environ")

// __attribute__ ((__dllimport__)) char *** __attribute__((__cdecl__)) __p__environ(void);
func X__p__environ(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p__environ->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p__environ.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p__wenviron = dll.NewProc("__p__wenviron")

// __attribute__ ((__dllimport__)) wchar_t *** __attribute__((__cdecl__)) __p__wenviron(void);
func X__p__wenviron(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p__wenviron->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p__wenviron.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p__pgmptr = dll.NewProc("__p__pgmptr")

// __attribute__ ((__dllimport__)) char ** __attribute__((__cdecl__)) __p__pgmptr(void);
func X__p__pgmptr(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p__pgmptr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p__pgmptr.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__p__wpgmptr = dll.NewProc("__p__wpgmptr")

// __attribute__ ((__dllimport__)) wchar_t ** __attribute__((__cdecl__)) __p__wpgmptr(void);
func X__p__wpgmptr(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__p__wpgmptr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__p__wpgmptr.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_get_pgmptr = dll.NewProc("_get_pgmptr")

// errno_t __attribute__((__cdecl__)) _get_pgmptr(char **_Value);
func X_get_pgmptr(tls *TLS, __Value uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_get_pgmptr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_pgmptr.Addr(), __Value)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_wpgmptr = dll.NewProc("_get_wpgmptr")

// errno_t __attribute__((__cdecl__)) _get_wpgmptr(wchar_t **_Value);
func X_get_wpgmptr(tls *TLS, __Value uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Value=%+v", __Value)
		defer func() { trc(`X_get_wpgmptr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_wpgmptr.Addr(), __Value)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_set_fmode = dll.NewProc("_set_fmode")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _set_fmode(int _Mode);
func X_set_fmode(tls *TLS, __Mode int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Mode=%+v", __Mode)
		defer func() { trc(`X_set_fmode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_fmode.Addr(), uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_fmode = dll.NewProc("_get_fmode")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _get_fmode(int *_PMode);
func X_get_fmode(tls *TLS, __PMode uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_PMode=%+v", __PMode)
		defer func() { trc(`X_get_fmode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_fmode.Addr(), __PMode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_exit = dll.NewProc("_exit")

// void __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _exit(int _Code) __attribute__ ((__noreturn__));
func X_exit(tls *TLS, __Code int32) {
	if __ccgo_strace {
		trc("_Code=%+v", __Code)
	}
	r0, r1, err := syscall.SyscallN(proc_exit.Addr(), uintptr(__Code))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var procquick_exit = dll.NewProc("quick_exit")

// void __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) quick_exit(int _Code) __attribute__ ((__noreturn__));
func Xquick_exit(tls *TLS, __Code int32) {
	if __ccgo_strace {
		trc("_Code=%+v", __Code)
	}
	r0, r1, err := syscall.SyscallN(procquick_exit.Addr(), uintptr(__Code))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_Exit = dll.NewProc("_Exit")

// void __attribute__((__cdecl__)) _Exit(int) __attribute__ ((__noreturn__));
func X_Exit(tls *TLS, _0 int32) {
	if __ccgo_strace {
		trc("0=%+v", _0)
	}
	r0, r1, err := syscall.SyscallN(proc_Exit.Addr(), uintptr(_0))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var procabort = dll.NewProc("abort")

// void __attribute__((__cdecl__)) __attribute__ ((__noreturn__)) abort(void);
func Xabort(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(procabort.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_set_abort_behavior = dll.NewProc("_set_abort_behavior")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) _set_abort_behavior(unsigned int _Flags,unsigned int _Mask);
func X_set_abort_behavior(tls *TLS, __Flags uint32, __Mask uint32) (r uint32) {
	if __ccgo_strace {
		trc("_Flags=%+v _Mask=%+v", __Flags, __Mask)
		defer func() { trc(`X_set_abort_behavior->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_abort_behavior.Addr(), uintptr(__Flags), uintptr(__Mask))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_abs64 = dll.NewProc("_abs64")

// long long __attribute__((__cdecl__)) _abs64( long long);
func X_abs64(tls *TLS, _x int64) (r int64) {
	if __ccgo_strace {
		trc("x=%+v", _x)
		defer func() { trc(`X_abs64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_abs64.Addr(), uintptr(_x))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var procatoi = dll.NewProc("atoi")

// int __attribute__((__cdecl__)) atoi(const char *_Str);
func Xatoi(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`Xatoi->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procatoi.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atoi_l = dll.NewProc("_atoi_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atoi_l(const char *_Str,_locale_t _Locale);
func X_atoi_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_atoi_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoi_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procatol = dll.NewProc("atol")

// long __attribute__((__cdecl__)) atol(const char *_Str);
func Xatol(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`Xatol->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procatol.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atol_l = dll.NewProc("_atol_l")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _atol_l(const char *_Str,_locale_t _Locale);
func X_atol_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_atol_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atol_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procbsearch = dll.NewProc("bsearch")

// void * __attribute__((__cdecl__)) bsearch(const void *_Key,const void *_Base,size_t _NumOfElements,size_t _SizeOfElements,int ( *_PtFuncCompare)(const void *,const void *));
func Xbsearch(tls *TLS, __Key uintptr, __Base uintptr, __NumOfElements Tsize_t, __SizeOfElements Tsize_t, __PtFuncCompare uintptr) (r uintptr) {
	X__ccgo_SyscallFP()
	panic(663)
}

var proc_byteswap_ushort = dll.NewProc("_byteswap_ushort")

// unsigned short __attribute__((__cdecl__)) _byteswap_ushort(unsigned short _Short);
func X_byteswap_ushort(tls *TLS, __Short uint16) (r uint16) {
	if __ccgo_strace {
		trc("_Short=%+v", __Short)
		defer func() { trc(`X_byteswap_ushort->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_byteswap_ushort.Addr(), uintptr(__Short))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint16(r0)
}

var proc_byteswap_ulong = dll.NewProc("_byteswap_ulong")

// unsigned long __attribute__((__cdecl__)) _byteswap_ulong (unsigned long _Long);
func X_byteswap_ulong(tls *TLS, __Long uint32) (r uint32) {
	if __ccgo_strace {
		trc("_Long=%+v", __Long)
		defer func() { trc(`X_byteswap_ulong->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_byteswap_ulong.Addr(), uintptr(__Long))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_byteswap_uint64 = dll.NewProc("_byteswap_uint64")

// unsigned long long __attribute__((__cdecl__)) _byteswap_uint64(unsigned long long _Int64);
func X_byteswap_uint64(tls *TLS, __Int64 uint64) (r uint64) {
	if __ccgo_strace {
		trc("_Int64=%+v", __Int64)
		defer func() { trc(`X_byteswap_uint64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_byteswap_uint64.Addr(), uintptr(__Int64))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var procdiv = dll.NewProc("div")

// div_t __attribute__((__cdecl__)) div(int _Numerator,int _Denominator);
func Xdiv(tls *TLS, __Numerator int32, __Denominator int32) (r Tdiv_t) {
	if __ccgo_strace {
		trc("_Numerator=%+v _Denominator=%+v", __Numerator, __Denominator)
		defer func() { trc(`Xdiv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procdiv.Addr(), uintptr(__Numerator), uintptr(__Denominator))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return *(*Tdiv_t)(unsafe.Pointer(&r0))
}

var procgetenv = dll.NewProc("getenv")

// char * __attribute__((__cdecl__)) getenv(const char *_VarName);
func Xgetenv(tls *TLS, __VarName uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_VarName=%+v", __VarName)
		defer func() { trc(`Xgetenv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procgetenv.Addr(), __VarName)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_itoa = dll.NewProc("_itoa")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _itoa(int _Value,char *_Dest,int _Radix);
func X_itoa(tls *TLS, __Value int32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_itoa->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_itoa.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_i64toa = dll.NewProc("_i64toa")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _i64toa( long long _Val,char *_DstBuf,int _Radix);
func X_i64toa(tls *TLS, __Val int64, __DstBuf uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Val=%+v _DstBuf=%+v _Radix=%+v", __Val, __DstBuf, __Radix)
		defer func() { trc(`X_i64toa->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_i64toa.Addr(), uintptr(__Val), __DstBuf, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ui64toa = dll.NewProc("_ui64toa")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _ui64toa(unsigned long long _Val,char *_DstBuf,int _Radix);
func X_ui64toa(tls *TLS, __Val uint64, __DstBuf uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Val=%+v _DstBuf=%+v _Radix=%+v", __Val, __DstBuf, __Radix)
		defer func() { trc(`X_ui64toa->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ui64toa.Addr(), uintptr(__Val), __DstBuf, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_atoi64 = dll.NewProc("_atoi64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _atoi64(const char *_String);
func X_atoi64(tls *TLS, __String uintptr) (r int64) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_atoi64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoi64.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_atoi64_l = dll.NewProc("_atoi64_l")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _atoi64_l(const char *_String,_locale_t _Locale);
func X_atoi64_l(tls *TLS, __String uintptr, __Locale T_locale_t) (r int64) {
	if __ccgo_strace {
		trc("_String=%+v _Locale=%+v", __String, __Locale)
		defer func() { trc(`X_atoi64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoi64_l.Addr(), __String, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_strtoi64 = dll.NewProc("_strtoi64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _strtoi64(const char *_String,char **_EndPtr,int _Radix);
func X_strtoi64(tls *TLS, __String uintptr, __EndPtr uintptr, __Radix int32) (r int64) {
	if __ccgo_strace {
		trc("_String=%+v _EndPtr=%+v _Radix=%+v", __String, __EndPtr, __Radix)
		defer func() { trc(`X_strtoi64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtoi64.Addr(), __String, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_strtoi64_l = dll.NewProc("_strtoi64_l")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _strtoi64_l(const char *_String,char **_EndPtr,int _Radix,_locale_t _Locale);
func X_strtoi64_l(tls *TLS, __String uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r int64) {
	if __ccgo_strace {
		trc("_String=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __String, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_strtoi64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtoi64_l.Addr(), __String, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_strtoui64 = dll.NewProc("_strtoui64")

// __attribute__ ((__dllimport__)) unsigned long long __attribute__((__cdecl__)) _strtoui64(const char *_String,char **_EndPtr,int _Radix);
func X_strtoui64(tls *TLS, __String uintptr, __EndPtr uintptr, __Radix int32) (r uint64) {
	if __ccgo_strace {
		trc("_String=%+v _EndPtr=%+v _Radix=%+v", __String, __EndPtr, __Radix)
		defer func() { trc(`X_strtoui64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtoui64.Addr(), __String, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_strtoui64_l = dll.NewProc("_strtoui64_l")

// __attribute__ ((__dllimport__)) unsigned long long __attribute__((__cdecl__)) _strtoui64_l(const char *_String,char **_EndPtr,int _Radix,_locale_t _Locale);
func X_strtoui64_l(tls *TLS, __String uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r uint64) {
	if __ccgo_strace {
		trc("_String=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __String, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_strtoui64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtoui64_l.Addr(), __String, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var procldiv = dll.NewProc("ldiv")

// ldiv_t __attribute__((__cdecl__)) ldiv(long _Numerator,long _Denominator);
func Xldiv(tls *TLS, __Numerator int32, __Denominator int32) (r Tldiv_t) {
	if __ccgo_strace {
		trc("_Numerator=%+v _Denominator=%+v", __Numerator, __Denominator)
		defer func() { trc(`Xldiv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procldiv.Addr(), uintptr(__Numerator), uintptr(__Denominator))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return *(*Tldiv_t)(unsafe.Pointer(&r0))
}

var proc_ltoa = dll.NewProc("_ltoa")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _ltoa(long _Value,char *_Dest,int _Radix);
func X_ltoa(tls *TLS, __Value int32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_ltoa->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ltoa.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procmblen = dll.NewProc("mblen")

// int __attribute__((__cdecl__)) mblen(const char *_Ch,size_t _MaxCount);
func Xmblen(tls *TLS, __Ch uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _MaxCount=%+v", __Ch, __MaxCount)
		defer func() { trc(`Xmblen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmblen.Addr(), __Ch, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_mblen_l = dll.NewProc("_mblen_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _mblen_l(const char *_Ch,size_t _MaxCount,_locale_t _Locale);
func X_mblen_l(tls *TLS, __Ch uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Ch=%+v _MaxCount=%+v _Locale=%+v", __Ch, __MaxCount, __Locale)
		defer func() { trc(`X_mblen_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mblen_l.Addr(), __Ch, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_mbstrlen = dll.NewProc("_mbstrlen")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _mbstrlen(const char *_Str);
func X_mbstrlen(tls *TLS, __Str uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_mbstrlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbstrlen.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_mbstrlen_l = dll.NewProc("_mbstrlen_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _mbstrlen_l(const char *_Str,_locale_t _Locale);
func X_mbstrlen_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_mbstrlen_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbstrlen_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_mbstrnlen = dll.NewProc("_mbstrnlen")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _mbstrnlen(const char *_Str,size_t _MaxCount);
func X_mbstrnlen(tls *TLS, __Str uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _MaxCount=%+v", __Str, __MaxCount)
		defer func() { trc(`X_mbstrnlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbstrnlen.Addr(), __Str, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_mbstrnlen_l = dll.NewProc("_mbstrnlen_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _mbstrnlen_l(const char *_Str,size_t _MaxCount,_locale_t _Locale);
func X_mbstrnlen_l(tls *TLS, __Str uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _MaxCount=%+v _Locale=%+v", __Str, __MaxCount, __Locale)
		defer func() { trc(`X_mbstrnlen_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbstrnlen_l.Addr(), __Str, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procmbtowc = dll.NewProc("mbtowc")

// int __attribute__((__cdecl__)) mbtowc(wchar_t * __restrict__ _DstCh,const char * __restrict__ _SrcCh,size_t _SrcSizeInBytes);
func Xmbtowc(tls *TLS, __DstCh uintptr, __SrcCh uintptr, __SrcSizeInBytes Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_DstCh=%+v _SrcCh=%+v _SrcSizeInBytes=%+v", __DstCh, __SrcCh, __SrcSizeInBytes)
		defer func() { trc(`Xmbtowc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmbtowc.Addr(), __DstCh, __SrcCh, uintptr(__SrcSizeInBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_mbtowc_l = dll.NewProc("_mbtowc_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _mbtowc_l(wchar_t * __restrict__ _DstCh,const char * __restrict__ _SrcCh,size_t _SrcSizeInBytes,_locale_t _Locale);
func X_mbtowc_l(tls *TLS, __DstCh uintptr, __SrcCh uintptr, __SrcSizeInBytes Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_DstCh=%+v _SrcCh=%+v _SrcSizeInBytes=%+v _Locale=%+v", __DstCh, __SrcCh, __SrcSizeInBytes, __Locale)
		defer func() { trc(`X_mbtowc_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbtowc_l.Addr(), __DstCh, __SrcCh, uintptr(__SrcSizeInBytes), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procmbstowcs = dll.NewProc("mbstowcs")

// size_t __attribute__((__cdecl__)) mbstowcs(wchar_t * __restrict__ _Dest,const char * __restrict__ _Source,size_t _MaxCount);
func Xmbstowcs(tls *TLS, __Dest uintptr, __Source uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _MaxCount=%+v", __Dest, __Source, __MaxCount)
		defer func() { trc(`Xmbstowcs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmbstowcs.Addr(), __Dest, __Source, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_mbstowcs_l = dll.NewProc("_mbstowcs_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _mbstowcs_l(wchar_t * __restrict__ _Dest,const char * __restrict__ _Source,size_t _MaxCount,_locale_t _Locale);
func X_mbstowcs_l(tls *TLS, __Dest uintptr, __Source uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _MaxCount=%+v _Locale=%+v", __Dest, __Source, __MaxCount, __Locale)
		defer func() { trc(`X_mbstowcs_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mbstowcs_l.Addr(), __Dest, __Source, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procrand = dll.NewProc("rand")

// int __attribute__((__cdecl__)) rand(void);
func Xrand(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xrand->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procrand.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_set_error_mode = dll.NewProc("_set_error_mode")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _set_error_mode(int _Mode);
func X_set_error_mode(tls *TLS, __Mode int32) (r int32) {
	if __ccgo_strace {
		trc("_Mode=%+v", __Mode)
		defer func() { trc(`X_set_error_mode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_set_error_mode.Addr(), uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procsrand = dll.NewProc("srand")

// void __attribute__((__cdecl__)) srand(unsigned int _Seed);
func Xsrand(tls *TLS, __Seed uint32) {
	if __ccgo_strace {
		trc("_Seed=%+v", __Seed)
	}
	r0, r1, err := syscall.SyscallN(procsrand.Addr(), uintptr(__Seed))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var procstrtol = dll.NewProc("strtol")

// long __attribute__((__cdecl__)) strtol(const char * __restrict__ _Str,char ** __restrict__ _EndPtr,int _Radix);
func Xstrtol(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`Xstrtol->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrtol.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strtol_l = dll.NewProc("_strtol_l")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _strtol_l(const char * __restrict__ _Str,char ** __restrict__ _EndPtr,int _Radix,_locale_t _Locale);
func X_strtol_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_strtol_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtol_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procstrtoul = dll.NewProc("strtoul")

// unsigned long __attribute__((__cdecl__)) strtoul(const char * __restrict__ _Str,char ** __restrict__ _EndPtr,int _Radix);
func Xstrtoul(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r uint32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`Xstrtoul->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrtoul.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_strtoul_l = dll.NewProc("_strtoul_l")

// __attribute__ ((__dllimport__)) unsigned long __attribute__((__cdecl__)) _strtoul_l(const char * __restrict__ _Str,char ** __restrict__ _EndPtr,int _Radix,_locale_t _Locale);
func X_strtoul_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r uint32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_strtoul_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtoul_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var procsystem = dll.NewProc("system")

// int __attribute__((__cdecl__)) system(const char *_Command);
func Xsystem(tls *TLS, __Command uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Command=%+v", __Command)
		defer func() { trc(`Xsystem->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procsystem.Addr(), __Command)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_ultoa = dll.NewProc("_ultoa")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _ultoa(unsigned long _Value,char *_Dest,int _Radix);
func X_ultoa(tls *TLS, __Value uint32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_ultoa->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ultoa.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwctomb = dll.NewProc("wctomb")

// int __attribute__((__cdecl__)) wctomb(char *_MbCh,wchar_t _WCh);
func Xwctomb(tls *TLS, __MbCh uintptr, __WCh Twchar_t) (r int32) {
	if __ccgo_strace {
		trc("_MbCh=%+v _WCh=%+v", __MbCh, __WCh)
		defer func() { trc(`Xwctomb->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwctomb.Addr(), __MbCh, uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wctomb_l = dll.NewProc("_wctomb_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wctomb_l(char *_MbCh,wchar_t _WCh,_locale_t _Locale);
func X_wctomb_l(tls *TLS, __MbCh uintptr, __WCh Twchar_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_MbCh=%+v _WCh=%+v _Locale=%+v", __MbCh, __WCh, __Locale)
		defer func() { trc(`X_wctomb_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wctomb_l.Addr(), __MbCh, uintptr(__WCh), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procwcstombs = dll.NewProc("wcstombs")

// size_t __attribute__((__cdecl__)) wcstombs(char * __restrict__ _Dest,const wchar_t * __restrict__ _Source,size_t _MaxCount);
func Xwcstombs(tls *TLS, __Dest uintptr, __Source uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _MaxCount=%+v", __Dest, __Source, __MaxCount)
		defer func() { trc(`Xwcstombs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstombs.Addr(), __Dest, __Source, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_wcstombs_l = dll.NewProc("_wcstombs_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _wcstombs_l(char * __restrict__ _Dest,const wchar_t * __restrict__ _Source,size_t _MaxCount,_locale_t _Locale);
func X_wcstombs_l(tls *TLS, __Dest uintptr, __Source uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _MaxCount=%+v _Locale=%+v", __Dest, __Source, __MaxCount, __Locale)
		defer func() { trc(`X_wcstombs_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstombs_l.Addr(), __Dest, __Source, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proccalloc = dll.NewProc("calloc")

// void * __attribute__((__cdecl__)) calloc(size_t _NumOfElements,size_t _SizeOfElements);
func Xcalloc(tls *TLS, __NumOfElements Tsize_t, __SizeOfElements Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_NumOfElements=%+v _SizeOfElements=%+v", __NumOfElements, __SizeOfElements)
		defer func() { trc(`Xcalloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proccalloc.Addr(), uintptr(__NumOfElements), uintptr(__SizeOfElements))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procfree = dll.NewProc("free")

// void __attribute__((__cdecl__)) free(void *_Memory);
func Xfree(tls *TLS, __Memory uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v", __Memory)
	}
	r0, r1, err := syscall.SyscallN(procfree.Addr(), __Memory)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var procmalloc = dll.NewProc("malloc")

// void * __attribute__((__cdecl__)) malloc(size_t _Size);
func Xmalloc(tls *TLS, __Size Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Size=%+v", __Size)
		defer func() { trc(`Xmalloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmalloc.Addr(), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procrealloc = dll.NewProc("realloc")

// void * __attribute__((__cdecl__)) realloc(void *_Memory,size_t _NewSize);
func Xrealloc(tls *TLS, __Memory uintptr, __NewSize Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _NewSize=%+v", __Memory, __NewSize)
		defer func() { trc(`Xrealloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procrealloc.Addr(), __Memory, uintptr(__NewSize))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_recalloc = dll.NewProc("_recalloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _recalloc(void *_Memory,size_t _Count,size_t _Size);
func X_recalloc(tls *TLS, __Memory uintptr, __Count Tsize_t, __Size Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _Count=%+v _Size=%+v", __Memory, __Count, __Size)
		defer func() { trc(`X_recalloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_recalloc.Addr(), __Memory, uintptr(__Count), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_free = dll.NewProc("_aligned_free")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _aligned_free(void *_Memory);
func X_aligned_free(tls *TLS, __Memory uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v", __Memory)
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_free.Addr(), __Memory)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_aligned_malloc = dll.NewProc("_aligned_malloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_malloc(size_t _Size,size_t _Alignment);
func X_aligned_malloc(tls *TLS, __Size Tsize_t, __Alignment Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Size=%+v _Alignment=%+v", __Size, __Alignment)
		defer func() { trc(`X_aligned_malloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_malloc.Addr(), uintptr(__Size), uintptr(__Alignment))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_offset_malloc = dll.NewProc("_aligned_offset_malloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_offset_malloc(size_t _Size,size_t _Alignment,size_t _Offset);
func X_aligned_offset_malloc(tls *TLS, __Size Tsize_t, __Alignment Tsize_t, __Offset Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Size=%+v _Alignment=%+v _Offset=%+v", __Size, __Alignment, __Offset)
		defer func() { trc(`X_aligned_offset_malloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_offset_malloc.Addr(), uintptr(__Size), uintptr(__Alignment), uintptr(__Offset))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_realloc = dll.NewProc("_aligned_realloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_realloc(void *_Memory,size_t _Size,size_t _Alignment);
func X_aligned_realloc(tls *TLS, __Memory uintptr, __Size Tsize_t, __Alignment Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _Size=%+v _Alignment=%+v", __Memory, __Size, __Alignment)
		defer func() { trc(`X_aligned_realloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_realloc.Addr(), __Memory, uintptr(__Size), uintptr(__Alignment))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_recalloc = dll.NewProc("_aligned_recalloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_recalloc(void *_Memory,size_t _Count,size_t _Size,size_t _Alignment);
func X_aligned_recalloc(tls *TLS, __Memory uintptr, __Count Tsize_t, __Size Tsize_t, __Alignment Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _Count=%+v _Size=%+v _Alignment=%+v", __Memory, __Count, __Size, __Alignment)
		defer func() { trc(`X_aligned_recalloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_recalloc.Addr(), __Memory, uintptr(__Count), uintptr(__Size), uintptr(__Alignment))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_offset_realloc = dll.NewProc("_aligned_offset_realloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_offset_realloc(void *_Memory,size_t _Size,size_t _Alignment,size_t _Offset);
func X_aligned_offset_realloc(tls *TLS, __Memory uintptr, __Size Tsize_t, __Alignment Tsize_t, __Offset Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _Size=%+v _Alignment=%+v _Offset=%+v", __Memory, __Size, __Alignment, __Offset)
		defer func() { trc(`X_aligned_offset_realloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_offset_realloc.Addr(), __Memory, uintptr(__Size), uintptr(__Alignment), uintptr(__Offset))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_aligned_offset_recalloc = dll.NewProc("_aligned_offset_recalloc")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _aligned_offset_recalloc(void *_Memory,size_t _Count,size_t _Size,size_t _Alignment,size_t _Offset);
func X_aligned_offset_recalloc(tls *TLS, __Memory uintptr, __Count Tsize_t, __Size Tsize_t, __Alignment Tsize_t, __Offset Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _Count=%+v _Size=%+v _Alignment=%+v _Offset=%+v", __Memory, __Count, __Size, __Alignment, __Offset)
		defer func() { trc(`X_aligned_offset_recalloc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_aligned_offset_recalloc.Addr(), __Memory, uintptr(__Count), uintptr(__Size), uintptr(__Alignment), uintptr(__Offset))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_itow = dll.NewProc("_itow")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _itow(int _Value,wchar_t *_Dest,int _Radix);
func X_itow(tls *TLS, __Value int32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_itow->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_itow.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ltow = dll.NewProc("_ltow")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _ltow(long _Value,wchar_t *_Dest,int _Radix);
func X_ltow(tls *TLS, __Value int32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_ltow->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ltow.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ultow = dll.NewProc("_ultow")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _ultow(unsigned long _Value,wchar_t *_Dest,int _Radix);
func X_ultow(tls *TLS, __Value uint32, __Dest uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Value=%+v _Dest=%+v _Radix=%+v", __Value, __Dest, __Radix)
		defer func() { trc(`X_ultow->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ultow.Addr(), uintptr(__Value), __Dest, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcstol = dll.NewProc("wcstol")

// long __attribute__((__cdecl__)) wcstol(const wchar_t * __restrict__ _Str,wchar_t ** __restrict__ _EndPtr,int _Radix);
func Xwcstol(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`Xwcstol->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstol.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcstol_l = dll.NewProc("_wcstol_l")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _wcstol_l(const wchar_t * __restrict__ _Str,wchar_t ** __restrict__ _EndPtr,int _Radix,_locale_t _Locale);
func X_wcstol_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_wcstol_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstol_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procwcstoul = dll.NewProc("wcstoul")

// unsigned long __attribute__((__cdecl__)) wcstoul(const wchar_t * __restrict__ _Str,wchar_t ** __restrict__ _EndPtr,int _Radix);
func Xwcstoul(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r uint32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`Xwcstoul->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstoul.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_wcstoul_l = dll.NewProc("_wcstoul_l")

// __attribute__ ((__dllimport__)) unsigned long __attribute__((__cdecl__)) _wcstoul_l(const wchar_t * __restrict__ _Str,wchar_t ** __restrict__ _EndPtr,int _Radix,_locale_t _Locale);
func X_wcstoul_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r uint32) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_wcstoul_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstoul_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_wsystem = dll.NewProc("_wsystem")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wsystem(const wchar_t *_Command);
func X_wsystem(tls *TLS, __Command uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Command=%+v", __Command)
		defer func() { trc(`X_wsystem->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wsystem.Addr(), __Command)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtoi = dll.NewProc("_wtoi")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wtoi(const wchar_t *_Str);
func X_wtoi(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_wtoi->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtoi.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtoi_l = dll.NewProc("_wtoi_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wtoi_l(const wchar_t *_Str,_locale_t _Locale);
func X_wtoi_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_wtoi_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtoi_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtol = dll.NewProc("_wtol")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _wtol(const wchar_t *_Str);
func X_wtol(tls *TLS, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_wtol->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtol.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wtol_l = dll.NewProc("_wtol_l")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _wtol_l(const wchar_t *_Str,_locale_t _Locale);
func X_wtol_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_wtol_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtol_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_i64tow = dll.NewProc("_i64tow")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _i64tow( long long _Val,wchar_t *_DstBuf,int _Radix);
func X_i64tow(tls *TLS, __Val int64, __DstBuf uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Val=%+v _DstBuf=%+v _Radix=%+v", __Val, __DstBuf, __Radix)
		defer func() { trc(`X_i64tow->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_i64tow.Addr(), uintptr(__Val), __DstBuf, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ui64tow = dll.NewProc("_ui64tow")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _ui64tow(unsigned long long _Val,wchar_t *_DstBuf,int _Radix);
func X_ui64tow(tls *TLS, __Val uint64, __DstBuf uintptr, __Radix int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Val=%+v _DstBuf=%+v _Radix=%+v", __Val, __DstBuf, __Radix)
		defer func() { trc(`X_ui64tow->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ui64tow.Addr(), uintptr(__Val), __DstBuf, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wtoi64 = dll.NewProc("_wtoi64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _wtoi64(const wchar_t *_Str);
func X_wtoi64(tls *TLS, __Str uintptr) (r int64) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_wtoi64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtoi64.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_wtoi64_l = dll.NewProc("_wtoi64_l")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _wtoi64_l(const wchar_t *_Str,_locale_t _Locale);
func X_wtoi64_l(tls *TLS, __Str uintptr, __Locale T_locale_t) (r int64) {
	if __ccgo_strace {
		trc("_Str=%+v _Locale=%+v", __Str, __Locale)
		defer func() { trc(`X_wtoi64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wtoi64_l.Addr(), __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_wcstoi64 = dll.NewProc("_wcstoi64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _wcstoi64(const wchar_t *_Str,wchar_t **_EndPtr,int _Radix);
func X_wcstoi64(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r int64) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`X_wcstoi64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstoi64.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_wcstoi64_l = dll.NewProc("_wcstoi64_l")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _wcstoi64_l(const wchar_t *_Str,wchar_t **_EndPtr,int _Radix,_locale_t _Locale);
func X_wcstoi64_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r int64) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_wcstoi64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstoi64_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_wcstoui64 = dll.NewProc("_wcstoui64")

// __attribute__ ((__dllimport__)) unsigned long long __attribute__((__cdecl__)) _wcstoui64(const wchar_t *_Str,wchar_t **_EndPtr,int _Radix);
func X_wcstoui64(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32) (r uint64) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v", __Str, __EndPtr, __Radix)
		defer func() { trc(`X_wcstoui64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstoui64.Addr(), __Str, __EndPtr, uintptr(__Radix))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_wcstoui64_l = dll.NewProc("_wcstoui64_l")

// __attribute__ ((__dllimport__)) unsigned long long __attribute__((__cdecl__)) _wcstoui64_l(const wchar_t *_Str ,wchar_t **_EndPtr,int _Radix,_locale_t _Locale);
func X_wcstoui64_l(tls *TLS, __Str uintptr, __EndPtr uintptr, __Radix int32, __Locale T_locale_t) (r uint64) {
	if __ccgo_strace {
		trc("_Str=%+v _EndPtr=%+v _Radix=%+v _Locale=%+v", __Str, __EndPtr, __Radix, __Locale)
		defer func() { trc(`X_wcstoui64_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcstoui64_l.Addr(), __Str, __EndPtr, uintptr(__Radix), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_putenv = dll.NewProc("_putenv")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _putenv(const char *_EnvString);
func X_putenv(tls *TLS, __EnvString uintptr) (r int32) {
	if __ccgo_strace {
		trc("_EnvString=%+v", __EnvString)
		defer func() { trc(`X_putenv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putenv.Addr(), __EnvString)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fullpath = dll.NewProc("_fullpath")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _fullpath(char *_FullPath,const char *_Path,size_t _SizeInBytes);
func X_fullpath(tls *TLS, __FullPath uintptr, __Path uintptr, __SizeInBytes Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_FullPath=%+v _Path=%+v _SizeInBytes=%+v", __FullPath, __Path, __SizeInBytes)
		defer func() { trc(`X_fullpath->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fullpath.Addr(), __FullPath, __Path, uintptr(__SizeInBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_atodbl = dll.NewProc("_atodbl")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atodbl(_CRT_DOUBLE *_Result,char *_Str);
func X_atodbl(tls *TLS, __Result uintptr, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v", __Result, __Str)
		defer func() { trc(`X_atodbl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atodbl.Addr(), __Result, __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atoldbl = dll.NewProc("_atoldbl")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atoldbl(_LDOUBLE *_Result,char *_Str);
func X_atoldbl(tls *TLS, __Result uintptr, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v", __Result, __Str)
		defer func() { trc(`X_atoldbl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoldbl.Addr(), __Result, __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atoflt = dll.NewProc("_atoflt")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atoflt(_CRT_FLOAT *_Result,char *_Str);
func X_atoflt(tls *TLS, __Result uintptr, __Str uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v", __Result, __Str)
		defer func() { trc(`X_atoflt->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoflt.Addr(), __Result, __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atodbl_l = dll.NewProc("_atodbl_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atodbl_l(_CRT_DOUBLE *_Result,char *_Str,_locale_t _Locale);
func X_atodbl_l(tls *TLS, __Result uintptr, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v _Locale=%+v", __Result, __Str, __Locale)
		defer func() { trc(`X_atodbl_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atodbl_l.Addr(), __Result, __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atoldbl_l = dll.NewProc("_atoldbl_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atoldbl_l(_LDOUBLE *_Result,char *_Str,_locale_t _Locale);
func X_atoldbl_l(tls *TLS, __Result uintptr, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v _Locale=%+v", __Result, __Str, __Locale)
		defer func() { trc(`X_atoldbl_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoldbl_l.Addr(), __Result, __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_atoflt_l = dll.NewProc("_atoflt_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _atoflt_l(_CRT_FLOAT *_Result,char *_Str,_locale_t _Locale);
func X_atoflt_l(tls *TLS, __Result uintptr, __Str uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Result=%+v _Str=%+v _Locale=%+v", __Result, __Str, __Locale)
		defer func() { trc(`X_atoflt_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_atoflt_l.Addr(), __Result, __Str, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_lrotl = dll.NewProc("_lrotl")

// unsigned long __attribute__((__cdecl__)) _lrotl(unsigned long,int);
func X_lrotl(tls *TLS, _0 uint32, _1 int32) (r uint32) {
	if __ccgo_strace {
		trc("0=%+v 1=%+v", _0, _1)
		defer func() { trc(`X_lrotl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_lrotl.Addr(), uintptr(_0), uintptr(_1))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_lrotr = dll.NewProc("_lrotr")

// unsigned long __attribute__((__cdecl__)) _lrotr(unsigned long,int);
func X_lrotr(tls *TLS, _0 uint32, _1 int32) (r uint32) {
	if __ccgo_strace {
		trc("0=%+v 1=%+v", _0, _1)
		defer func() { trc(`X_lrotr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_lrotr.Addr(), uintptr(_0), uintptr(_1))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_makepath = dll.NewProc("_makepath")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _makepath(char *_Path,const char *_Drive,const char *_Dir,const char *_Filename,const char *_Ext);
func X_makepath(tls *TLS, __Path uintptr, __Drive uintptr, __Dir uintptr, __Filename uintptr, __Ext uintptr) {
	if __ccgo_strace {
		trc("_Path=%+v _Drive=%+v _Dir=%+v _Filename=%+v _Ext=%+v", __Path, __Drive, __Dir, __Filename, __Ext)
	}
	r0, r1, err := syscall.SyscallN(proc_makepath.Addr(), __Path, __Drive, __Dir, __Filename, __Ext)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_rotl64 = dll.NewProc("_rotl64")

// unsigned long long __attribute__((__cdecl__)) _rotl64(unsigned long long _Val,int _Shift);
func X_rotl64(tls *TLS, __Val uint64, __Shift int32) (r uint64) {
	if __ccgo_strace {
		trc("_Val=%+v _Shift=%+v", __Val, __Shift)
		defer func() { trc(`X_rotl64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_rotl64.Addr(), uintptr(__Val), uintptr(__Shift))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_rotr64 = dll.NewProc("_rotr64")

// unsigned long long __attribute__((__cdecl__)) _rotr64(unsigned long long Value,int Shift);
func X_rotr64(tls *TLS, _Value uint64, _Shift int32) (r uint64) {
	if __ccgo_strace {
		trc("Value=%+v Shift=%+v", _Value, _Shift)
		defer func() { trc(`X_rotr64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_rotr64.Addr(), uintptr(_Value), uintptr(_Shift))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_rotr = dll.NewProc("_rotr")

// unsigned int __attribute__((__cdecl__)) _rotr(unsigned int _Val,int _Shift);
func X_rotr(tls *TLS, __Val uint32, __Shift int32) (r uint32) {
	if __ccgo_strace {
		trc("_Val=%+v _Shift=%+v", __Val, __Shift)
		defer func() { trc(`X_rotr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_rotr.Addr(), uintptr(__Val), uintptr(__Shift))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_rotl = dll.NewProc("_rotl")

// unsigned int __attribute__((__cdecl__)) _rotl(unsigned int _Val,int _Shift);
func X_rotl(tls *TLS, __Val uint32, __Shift int32) (r uint32) {
	if __ccgo_strace {
		trc("_Val=%+v _Shift=%+v", __Val, __Shift)
		defer func() { trc(`X_rotl->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_rotl.Addr(), uintptr(__Val), uintptr(__Shift))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_searchenv = dll.NewProc("_searchenv")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _searchenv(const char *_Filename,const char *_EnvVar,char *_ResultPath);
func X_searchenv(tls *TLS, __Filename uintptr, __EnvVar uintptr, __ResultPath uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _EnvVar=%+v _ResultPath=%+v", __Filename, __EnvVar, __ResultPath)
	}
	r0, r1, err := syscall.SyscallN(proc_searchenv.Addr(), __Filename, __EnvVar, __ResultPath)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_splitpath = dll.NewProc("_splitpath")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _splitpath(const char *_FullPath,char *_Drive,char *_Dir,char *_Filename,char *_Ext);
func X_splitpath(tls *TLS, __FullPath uintptr, __Drive uintptr, __Dir uintptr, __Filename uintptr, __Ext uintptr) {
	if __ccgo_strace {
		trc("_FullPath=%+v _Drive=%+v _Dir=%+v _Filename=%+v _Ext=%+v", __FullPath, __Drive, __Dir, __Filename, __Ext)
	}
	r0, r1, err := syscall.SyscallN(proc_splitpath.Addr(), __FullPath, __Drive, __Dir, __Filename, __Ext)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_swab = dll.NewProc("_swab")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _swab(char *_Buf1,char *_Buf2,int _SizeInBytes);
func X_swab(tls *TLS, __Buf1 uintptr, __Buf2 uintptr, __SizeInBytes int32) {
	if __ccgo_strace {
		trc("_Buf1=%+v _Buf2=%+v _SizeInBytes=%+v", __Buf1, __Buf2, __SizeInBytes)
	}
	r0, r1, err := syscall.SyscallN(proc_swab.Addr(), __Buf1, __Buf2, uintptr(__SizeInBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_wfullpath = dll.NewProc("_wfullpath")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wfullpath(wchar_t *_FullPath,const wchar_t *_Path,size_t _SizeInWords);
func X_wfullpath(tls *TLS, __FullPath uintptr, __Path uintptr, __SizeInWords Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_FullPath=%+v _Path=%+v _SizeInWords=%+v", __FullPath, __Path, __SizeInWords)
		defer func() { trc(`X_wfullpath->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfullpath.Addr(), __FullPath, __Path, uintptr(__SizeInWords))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wmakepath = dll.NewProc("_wmakepath")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _wmakepath(wchar_t *_ResultPath,const wchar_t *_Drive,const wchar_t *_Dir,const wchar_t *_Filename,const wchar_t *_Ext);
func X_wmakepath(tls *TLS, __ResultPath uintptr, __Drive uintptr, __Dir uintptr, __Filename uintptr, __Ext uintptr) {
	if __ccgo_strace {
		trc("_ResultPath=%+v _Drive=%+v _Dir=%+v _Filename=%+v _Ext=%+v", __ResultPath, __Drive, __Dir, __Filename, __Ext)
	}
	r0, r1, err := syscall.SyscallN(proc_wmakepath.Addr(), __ResultPath, __Drive, __Dir, __Filename, __Ext)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_wsearchenv = dll.NewProc("_wsearchenv")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _wsearchenv(const wchar_t *_Filename,const wchar_t *_EnvVar,wchar_t *_ResultPath);
func X_wsearchenv(tls *TLS, __Filename uintptr, __EnvVar uintptr, __ResultPath uintptr) {
	if __ccgo_strace {
		trc("_Filename=%+v _EnvVar=%+v _ResultPath=%+v", __Filename, __EnvVar, __ResultPath)
	}
	r0, r1, err := syscall.SyscallN(proc_wsearchenv.Addr(), __Filename, __EnvVar, __ResultPath)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_wsplitpath = dll.NewProc("_wsplitpath")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _wsplitpath(const wchar_t *_FullPath,wchar_t *_Drive,wchar_t *_Dir,wchar_t *_Filename,wchar_t *_Ext);
func X_wsplitpath(tls *TLS, __FullPath uintptr, __Drive uintptr, __Dir uintptr, __Filename uintptr, __Ext uintptr) {
	if __ccgo_strace {
		trc("_FullPath=%+v _Drive=%+v _Dir=%+v _Filename=%+v _Ext=%+v", __FullPath, __Drive, __Dir, __Filename, __Ext)
	}
	r0, r1, err := syscall.SyscallN(proc_wsplitpath.Addr(), __FullPath, __Drive, __Dir, __Filename, __Ext)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_beep = dll.NewProc("_beep")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _beep(unsigned _Frequency,unsigned _Duration) __attribute__ ((__deprecated__));
func X_beep(tls *TLS, __Frequency uint32, __Duration uint32) {
	if __ccgo_strace {
		trc("_Frequency=%+v _Duration=%+v", __Frequency, __Duration)
	}
	r0, r1, err := syscall.SyscallN(proc_beep.Addr(), uintptr(__Frequency), uintptr(__Duration))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_seterrormode = dll.NewProc("_seterrormode")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _seterrormode(int _Mode) __attribute__ ((__deprecated__));
func X_seterrormode(tls *TLS, __Mode int32) {
	if __ccgo_strace {
		trc("_Mode=%+v", __Mode)
	}
	r0, r1, err := syscall.SyscallN(proc_seterrormode.Addr(), uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_sleep = dll.NewProc("_sleep")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _sleep(unsigned long _Duration) __attribute__ ((__deprecated__));
func X_sleep(tls *TLS, __Duration uint32) {
	if __ccgo_strace {
		trc("_Duration=%+v", __Duration)
	}
	r0, r1, err := syscall.SyscallN(proc_sleep.Addr(), uintptr(__Duration))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

type Tlldiv_t = struct {
	Fquot int64
	Frem  int64
}

var procllabs = dll.NewProc("llabs")

// long long __attribute__((__cdecl__)) llabs(long long);
func Xllabs(tls *TLS, _0 int64) (r int64) {
	if __ccgo_strace {
		trc("0=%+v", _0)
		defer func() { trc(`Xllabs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procllabs.Addr(), uintptr(_0))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var procstrtoll = dll.NewProc("strtoll")

// long long __attribute__((__cdecl__)) strtoll(const char * __restrict__, char ** __restrict, int);
func Xstrtoll(tls *TLS, _0 uintptr, _1 uintptr, _2 int32) (r int64) {
	if __ccgo_strace {
		trc("0=%+v 1=%+v 2=%+v", _0, _1, _2)
		defer func() { trc(`Xstrtoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrtoll.Addr(), _0, _1, uintptr(_2))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var procstrtoull = dll.NewProc("strtoull")

// unsigned long long __attribute__((__cdecl__)) strtoull(const char * __restrict__, char ** __restrict__, int);
func Xstrtoull(tls *TLS, _0 uintptr, _1 uintptr, _2 int32) (r uint64) {
	if __ccgo_strace {
		trc("0=%+v 1=%+v 2=%+v", _0, _1, _2)
		defer func() { trc(`Xstrtoull->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrtoull.Addr(), _0, _1, uintptr(_2))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var procatoll = dll.NewProc("atoll")

// long long __attribute__((__cdecl__)) atoll (const char *);
func Xatoll(tls *TLS, _0 uintptr) (r int64) {
	if __ccgo_strace {
		trc("0=%+v", _0)
		defer func() { trc(`Xatoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procatoll.Addr(), _0)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

type T_HEAPINFO = struct {
	F_pentry  uintptr
	F_size    Tsize_t
	F_useflag int32
}

type T_heapinfo = T_HEAPINFO

var proc_resetstkoflw = dll.NewProc("_resetstkoflw")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _resetstkoflw (void);
func X_resetstkoflw(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_resetstkoflw->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_resetstkoflw.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_expand = dll.NewProc("_expand")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _expand(void *_Memory,size_t _NewSize);
func X_expand(tls *TLS, __Memory uintptr, __NewSize Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Memory=%+v _NewSize=%+v", __Memory, __NewSize)
		defer func() { trc(`X_expand->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_expand.Addr(), __Memory, uintptr(__NewSize))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_msize = dll.NewProc("_msize")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _msize(void *_Memory);
func X_msize(tls *TLS, __Memory uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Memory=%+v", __Memory)
		defer func() { trc(`X_msize->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_msize.Addr(), __Memory)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_heapchk = dll.NewProc("_heapchk")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _heapchk(void);
func X_heapchk(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_heapchk->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_heapchk.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_heapmin = dll.NewProc("_heapmin")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _heapmin(void);
func X_heapmin(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_heapmin->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_heapmin.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_heapwalk = dll.NewProc("_heapwalk")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _heapwalk(_HEAPINFO *_EntryInfo);
func X_heapwalk(tls *TLS, __EntryInfo uintptr) (r int32) {
	if __ccgo_strace {
		trc("_EntryInfo=%+v", __EntryInfo)
		defer func() { trc(`X_heapwalk->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_heapwalk.Addr(), __EntryInfo)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_get_heap_handle = dll.NewProc("_get_heap_handle")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _get_heap_handle(void);
func X_get_heap_handle(tls *TLS) (r Tintptr_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_get_heap_handle->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_heap_handle.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_memccpy = dll.NewProc("_memccpy")

// __attribute__ ((__dllimport__)) void * __attribute__((__cdecl__)) _memccpy(void *_Dst,const void *_Src,int _Val,size_t _MaxCount);
func X_memccpy(tls *TLS, __Dst uintptr, __Src uintptr, __Val int32, __MaxCount Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _Val=%+v _MaxCount=%+v", __Dst, __Src, __Val, __MaxCount)
		defer func() { trc(`X_memccpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_memccpy.Addr(), __Dst, __Src, uintptr(__Val), uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procmemchr = dll.NewProc("memchr")

// void * __attribute__((__cdecl__)) memchr(const void *_Buf ,int _Val,size_t _MaxCount);
func Xmemchr(tls *TLS, __Buf uintptr, __Val int32, __MaxCount Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Buf=%+v _Val=%+v _MaxCount=%+v", __Buf, __Val, __MaxCount)
		defer func() { trc(`Xmemchr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemchr.Addr(), __Buf, uintptr(__Val), uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_memicmp = dll.NewProc("_memicmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _memicmp(const void *_Buf1,const void *_Buf2,size_t _Size);
func X_memicmp(tls *TLS, __Buf1 uintptr, __Buf2 uintptr, __Size Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Buf1=%+v _Buf2=%+v _Size=%+v", __Buf1, __Buf2, __Size)
		defer func() { trc(`X_memicmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_memicmp.Addr(), __Buf1, __Buf2, uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_memicmp_l = dll.NewProc("_memicmp_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _memicmp_l(const void *_Buf1,const void *_Buf2,size_t _Size,_locale_t _Locale);
func X_memicmp_l(tls *TLS, __Buf1 uintptr, __Buf2 uintptr, __Size Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Buf1=%+v _Buf2=%+v _Size=%+v _Locale=%+v", __Buf1, __Buf2, __Size, __Locale)
		defer func() { trc(`X_memicmp_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_memicmp_l.Addr(), __Buf1, __Buf2, uintptr(__Size), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procmemcmp = dll.NewProc("memcmp")

// int __attribute__((__cdecl__)) memcmp(const void *_Buf1,const void *_Buf2,size_t _Size);
func Xmemcmp(tls *TLS, __Buf1 uintptr, __Buf2 uintptr, __Size Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Buf1=%+v _Buf2=%+v _Size=%+v", __Buf1, __Buf2, __Size)
		defer func() { trc(`Xmemcmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemcmp.Addr(), __Buf1, __Buf2, uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procmemcpy = dll.NewProc("memcpy")

// void * __attribute__((__cdecl__)) memcpy(void * __restrict__ _Dst,const void * __restrict__ _Src,size_t _Size);
func Xmemcpy(tls *TLS, __Dst uintptr, __Src uintptr, __Size Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _Size=%+v", __Dst, __Src, __Size)
		defer func() { trc(`Xmemcpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemcpy.Addr(), __Dst, __Src, uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procmemcpy_s = dll.NewProc("memcpy_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) memcpy_s (void *_dest,size_t _numberOfElements,const void *_src,size_t _count);
func Xmemcpy_s(tls *TLS, __dest uintptr, __numberOfElements Tsize_t, __src uintptr, __count Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_dest=%+v _numberOfElements=%+v _src=%+v _count=%+v", __dest, __numberOfElements, __src, __count)
		defer func() { trc(`Xmemcpy_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemcpy_s.Addr(), __dest, uintptr(__numberOfElements), __src, uintptr(__count))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var procmemset = dll.NewProc("memset")

// void * __attribute__((__cdecl__)) memset(void *_Dst,int _Val,size_t _Size);
func Xmemset(tls *TLS, __Dst uintptr, __Val int32, __Size Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dst=%+v _Val=%+v _Size=%+v", __Dst, __Val, __Size)
		defer func() { trc(`Xmemset->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemset.Addr(), __Dst, uintptr(__Val), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strset = dll.NewProc("_strset")

// char * __attribute__((__cdecl__)) _strset(char *_Str,int _Val);
func X_strset(tls *TLS, __Str uintptr, __Val int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Val=%+v", __Str, __Val)
		defer func() { trc(`X_strset->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strset.Addr(), __Str, uintptr(__Val))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrcpy = dll.NewProc("strcpy")

// char * __attribute__((__cdecl__)) strcpy(char * __restrict__ _Dest,const char * __restrict__ _Source);
func Xstrcpy(tls *TLS, __Dest uintptr, __Source uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v", __Dest, __Source)
		defer func() { trc(`Xstrcpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrcpy.Addr(), __Dest, __Source)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrcat = dll.NewProc("strcat")

// char * __attribute__((__cdecl__)) strcat(char * __restrict__ _Dest,const char * __restrict__ _Source);
func Xstrcat(tls *TLS, __Dest uintptr, __Source uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v", __Dest, __Source)
		defer func() { trc(`Xstrcat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrcat.Addr(), __Dest, __Source)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrcmp = dll.NewProc("strcmp")

// int __attribute__((__cdecl__)) strcmp(const char *_Str1,const char *_Str2);
func Xstrcmp(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`Xstrcmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrcmp.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procstrlen = dll.NewProc("strlen")

// size_t __attribute__((__cdecl__)) strlen(const char *_Str);
func Xstrlen(tls *TLS, __Str uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`Xstrlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrlen.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procstrnlen = dll.NewProc("strnlen")

// size_t __attribute__((__cdecl__)) strnlen(const char *_Str,size_t _MaxCount);
func Xstrnlen(tls *TLS, __Str uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _MaxCount=%+v", __Str, __MaxCount)
		defer func() { trc(`Xstrnlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrnlen.Addr(), __Str, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procmemmove = dll.NewProc("memmove")

// void * __attribute__((__cdecl__)) memmove(void *_Dst,const void *_Src,size_t _Size);
func Xmemmove(tls *TLS, __Dst uintptr, __Src uintptr, __Size Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _Size=%+v", __Dst, __Src, __Size)
		defer func() { trc(`Xmemmove->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmemmove.Addr(), __Dst, __Src, uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strdup = dll.NewProc("_strdup")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strdup(const char *_Src);
func X_strdup(tls *TLS, __Src uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Src=%+v", __Src)
		defer func() { trc(`X_strdup->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strdup.Addr(), __Src)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrchr = dll.NewProc("strchr")

// char * __attribute__((__cdecl__)) strchr(const char *_Str,int _Val);
func Xstrchr(tls *TLS, __Str uintptr, __Val int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Val=%+v", __Str, __Val)
		defer func() { trc(`Xstrchr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrchr.Addr(), __Str, uintptr(__Val))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_stricmp = dll.NewProc("_stricmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _stricmp(const char *_Str1,const char *_Str2);
func X_stricmp(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`X_stricmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stricmp.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_stricmp_l = dll.NewProc("_stricmp_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _stricmp_l(const char *_Str1,const char *_Str2,_locale_t _Locale);
func X_stricmp_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_stricmp_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stricmp_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procstrcoll = dll.NewProc("strcoll")

// int __attribute__((__cdecl__)) strcoll(const char *_Str1,const char *_Str2);
func Xstrcoll(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`Xstrcoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrcoll.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strcoll_l = dll.NewProc("_strcoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strcoll_l(const char *_Str1,const char *_Str2,_locale_t _Locale);
func X_strcoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_strcoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strcoll_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_stricoll = dll.NewProc("_stricoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _stricoll(const char *_Str1,const char *_Str2);
func X_stricoll(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`X_stricoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stricoll.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_stricoll_l = dll.NewProc("_stricoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _stricoll_l(const char *_Str1,const char *_Str2,_locale_t _Locale);
func X_stricoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_stricoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stricoll_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strncoll = dll.NewProc("_strncoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strncoll (const char *_Str1,const char *_Str2,size_t _MaxCount);
func X_strncoll(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_strncoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strncoll.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strncoll_l = dll.NewProc("_strncoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strncoll_l(const char *_Str1,const char *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_strncoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_strncoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strncoll_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strnicoll = dll.NewProc("_strnicoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strnicoll (const char *_Str1,const char *_Str2,size_t _MaxCount);
func X_strnicoll(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_strnicoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strnicoll.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strnicoll_l = dll.NewProc("_strnicoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strnicoll_l(const char *_Str1,const char *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_strnicoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_strnicoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strnicoll_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procstrcspn = dll.NewProc("strcspn")

// size_t __attribute__((__cdecl__)) strcspn(const char *_Str,const char *_Control);
func Xstrcspn(tls *TLS, __Str uintptr, __Control uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xstrcspn->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrcspn.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_strerror = dll.NewProc("_strerror")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strerror(const char *_ErrMsg);
func X_strerror(tls *TLS, __ErrMsg uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_ErrMsg=%+v", __ErrMsg)
		defer func() { trc(`X_strerror->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strerror.Addr(), __ErrMsg)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrerror = dll.NewProc("strerror")

// char * __attribute__((__cdecl__)) strerror(int);
func Xstrerror(tls *TLS, _0 int32) (r uintptr) {
	if __ccgo_strace {
		trc("0=%+v", _0)
		defer func() { trc(`Xstrerror->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrerror.Addr(), uintptr(_0))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strlwr = dll.NewProc("_strlwr")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strlwr(char *_String);
func X_strlwr(tls *TLS, __String uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_strlwr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strlwr.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrncat = dll.NewProc("strncat")

// char * __attribute__((__cdecl__)) strncat(char * __restrict__ _Dest,const char * __restrict__ _Source,size_t _Count);
func Xstrncat(tls *TLS, __Dest uintptr, __Source uintptr, __Count Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _Count=%+v", __Dest, __Source, __Count)
		defer func() { trc(`Xstrncat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrncat.Addr(), __Dest, __Source, uintptr(__Count))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrncmp = dll.NewProc("strncmp")

// int __attribute__((__cdecl__)) strncmp(const char *_Str1,const char *_Str2,size_t _MaxCount);
func Xstrncmp(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`Xstrncmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrncmp.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strnicmp = dll.NewProc("_strnicmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strnicmp(const char *_Str1,const char *_Str2,size_t _MaxCount);
func X_strnicmp(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_strnicmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strnicmp.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_strnicmp_l = dll.NewProc("_strnicmp_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _strnicmp_l(const char *_Str1,const char *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_strnicmp_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_strnicmp_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strnicmp_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procstrncpy = dll.NewProc("strncpy")

// char *strncpy(char * __restrict__ _Dest,const char * __restrict__ _Source,size_t _Count);
func Xstrncpy(tls *TLS, __Dest uintptr, __Source uintptr, __Count Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _Count=%+v", __Dest, __Source, __Count)
		defer func() { trc(`Xstrncpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrncpy.Addr(), __Dest, __Source, uintptr(__Count))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strnset = dll.NewProc("_strnset")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strnset(char *_Str,int _Val,size_t _MaxCount);
func X_strnset(tls *TLS, __Str uintptr, __Val int32, __MaxCount Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Val=%+v _MaxCount=%+v", __Str, __Val, __MaxCount)
		defer func() { trc(`X_strnset->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strnset.Addr(), __Str, uintptr(__Val), uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrpbrk = dll.NewProc("strpbrk")

// char * __attribute__((__cdecl__)) strpbrk(const char *_Str,const char *_Control);
func Xstrpbrk(tls *TLS, __Str uintptr, __Control uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xstrpbrk->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrpbrk.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrrchr = dll.NewProc("strrchr")

// char * __attribute__((__cdecl__)) strrchr(const char *_Str,int _Ch);
func Xstrrchr(tls *TLS, __Str uintptr, __Ch int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Ch=%+v", __Str, __Ch)
		defer func() { trc(`Xstrrchr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrrchr.Addr(), __Str, uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strrev = dll.NewProc("_strrev")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strrev(char *_Str);
func X_strrev(tls *TLS, __Str uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_strrev->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strrev.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrspn = dll.NewProc("strspn")

// size_t __attribute__((__cdecl__)) strspn(const char *_Str,const char *_Control);
func Xstrspn(tls *TLS, __Str uintptr, __Control uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xstrspn->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrspn.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procstrstr = dll.NewProc("strstr")

// char * __attribute__((__cdecl__)) strstr(const char *_Str,const char *_SubStr);
func Xstrstr(tls *TLS, __Str uintptr, __SubStr uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _SubStr=%+v", __Str, __SubStr)
		defer func() { trc(`Xstrstr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrstr.Addr(), __Str, __SubStr)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrtok = dll.NewProc("strtok")

// char * __attribute__((__cdecl__)) strtok(char * __restrict__ _Str,const char * __restrict__ _Delim);
func Xstrtok(tls *TLS, __Str uintptr, __Delim uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Delim=%+v", __Str, __Delim)
		defer func() { trc(`Xstrtok->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrtok.Addr(), __Str, __Delim)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strupr = dll.NewProc("_strupr")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strupr(char *_String);
func X_strupr(tls *TLS, __String uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_strupr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strupr.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strupr_l = dll.NewProc("_strupr_l")

// __attribute__ ((__dllimport__)) char *_strupr_l(char *_String,_locale_t _Locale);
func X_strupr_l(tls *TLS, __String uintptr, __Locale T_locale_t) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v _Locale=%+v", __String, __Locale)
		defer func() { trc(`X_strupr_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strupr_l.Addr(), __String, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procstrxfrm = dll.NewProc("strxfrm")

// size_t __attribute__((__cdecl__)) strxfrm(char * __restrict__ _Dst,const char * __restrict__ _Src,size_t _MaxCount);
func Xstrxfrm(tls *TLS, __Dst uintptr, __Src uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _MaxCount=%+v", __Dst, __Src, __MaxCount)
		defer func() { trc(`Xstrxfrm->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrxfrm.Addr(), __Dst, __Src, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_strxfrm_l = dll.NewProc("_strxfrm_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _strxfrm_l(char * __restrict__ _Dst,const char * __restrict__ _Src,size_t _MaxCount,_locale_t _Locale);
func X_strxfrm_l(tls *TLS, __Dst uintptr, __Src uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _MaxCount=%+v _Locale=%+v", __Dst, __Src, __MaxCount, __Locale)
		defer func() { trc(`X_strxfrm_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strxfrm_l.Addr(), __Dst, __Src, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_wcsdup = dll.NewProc("_wcsdup")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcsdup(const wchar_t *_Str);
func X_wcsdup(tls *TLS, __Str uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_wcsdup->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsdup.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcscat = dll.NewProc("wcscat")

// wchar_t * __attribute__((__cdecl__)) wcscat(wchar_t * __restrict__ _Dest,const wchar_t * __restrict__ _Source);
func Xwcscat(tls *TLS, __Dest uintptr, __Source uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v", __Dest, __Source)
		defer func() { trc(`Xwcscat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcscat.Addr(), __Dest, __Source)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcschr = dll.NewProc("wcschr")

// wchar_t * __attribute__((__cdecl__)) wcschr(const wchar_t *_Str,wchar_t _Ch);
func Xwcschr(tls *TLS, __Str uintptr, __Ch Twchar_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Ch=%+v", __Str, __Ch)
		defer func() { trc(`Xwcschr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcschr.Addr(), __Str, uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcscmp = dll.NewProc("wcscmp")

// int __attribute__((__cdecl__)) wcscmp(const wchar_t *_Str1,const wchar_t *_Str2);
func Xwcscmp(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`Xwcscmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcscmp.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procwcscpy = dll.NewProc("wcscpy")

// wchar_t * __attribute__((__cdecl__)) wcscpy(wchar_t * __restrict__ _Dest,const wchar_t * __restrict__ _Source);
func Xwcscpy(tls *TLS, __Dest uintptr, __Source uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v", __Dest, __Source)
		defer func() { trc(`Xwcscpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcscpy.Addr(), __Dest, __Source)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcscspn = dll.NewProc("wcscspn")

// size_t __attribute__((__cdecl__)) wcscspn(const wchar_t *_Str,const wchar_t *_Control);
func Xwcscspn(tls *TLS, __Str uintptr, __Control uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xwcscspn->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcscspn.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcslen = dll.NewProc("wcslen")

// size_t __attribute__((__cdecl__)) wcslen(const wchar_t *_Str);
func Xwcslen(tls *TLS, __Str uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`Xwcslen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcslen.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcsnlen = dll.NewProc("wcsnlen")

// size_t __attribute__((__cdecl__)) wcsnlen(const wchar_t *_Src,size_t _MaxCount);
func Xwcsnlen(tls *TLS, __Src uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Src=%+v _MaxCount=%+v", __Src, __MaxCount)
		defer func() { trc(`Xwcsnlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsnlen.Addr(), __Src, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcsncat = dll.NewProc("wcsncat")

// wchar_t *wcsncat(wchar_t * __restrict__ _Dest,const wchar_t * __restrict__ _Source,size_t _Count);
func Xwcsncat(tls *TLS, __Dest uintptr, __Source uintptr, __Count Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _Count=%+v", __Dest, __Source, __Count)
		defer func() { trc(`Xwcsncat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsncat.Addr(), __Dest, __Source, uintptr(__Count))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcsncmp = dll.NewProc("wcsncmp")

// int __attribute__((__cdecl__)) wcsncmp(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount);
func Xwcsncmp(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`Xwcsncmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsncmp.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procwcsncpy = dll.NewProc("wcsncpy")

// wchar_t *wcsncpy(wchar_t * __restrict__ _Dest,const wchar_t * __restrict__ _Source,size_t _Count);
func Xwcsncpy(tls *TLS, __Dest uintptr, __Source uintptr, __Count Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _Count=%+v", __Dest, __Source, __Count)
		defer func() { trc(`Xwcsncpy->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsncpy.Addr(), __Dest, __Source, uintptr(__Count))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcspbrk = dll.NewProc("wcspbrk")

// wchar_t * __attribute__((__cdecl__)) wcspbrk(const wchar_t *_Str,const wchar_t *_Control);
func Xwcspbrk(tls *TLS, __Str uintptr, __Control uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xwcspbrk->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcspbrk.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcsrchr = dll.NewProc("wcsrchr")

// wchar_t * __attribute__((__cdecl__)) wcsrchr(const wchar_t *_Str,wchar_t _Ch);
func Xwcsrchr(tls *TLS, __Str uintptr, __Ch Twchar_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Ch=%+v", __Str, __Ch)
		defer func() { trc(`Xwcsrchr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsrchr.Addr(), __Str, uintptr(__Ch))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcsspn = dll.NewProc("wcsspn")

// size_t __attribute__((__cdecl__)) wcsspn(const wchar_t *_Str,const wchar_t *_Control);
func Xwcsspn(tls *TLS, __Str uintptr, __Control uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Str=%+v _Control=%+v", __Str, __Control)
		defer func() { trc(`Xwcsspn->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsspn.Addr(), __Str, __Control)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcsstr = dll.NewProc("wcsstr")

// wchar_t * __attribute__((__cdecl__)) wcsstr(const wchar_t *_Str,const wchar_t *_SubStr);
func Xwcsstr(tls *TLS, __Str uintptr, __SubStr uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _SubStr=%+v", __Str, __SubStr)
		defer func() { trc(`Xwcsstr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsstr.Addr(), __Str, __SubStr)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcstok = dll.NewProc("wcstok")

// wchar_t * __attribute__((__cdecl__)) wcstok(wchar_t * __restrict__ _Str,const wchar_t * __restrict__ _Delim);
func Xwcstok(tls *TLS, __Str uintptr, __Delim uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Delim=%+v", __Str, __Delim)
		defer func() { trc(`Xwcstok->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstok.Addr(), __Str, __Delim)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcserror = dll.NewProc("_wcserror")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcserror(int _ErrNum);
func X_wcserror(tls *TLS, __ErrNum int32) (r uintptr) {
	if __ccgo_strace {
		trc("_ErrNum=%+v", __ErrNum)
		defer func() { trc(`X_wcserror->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcserror.Addr(), uintptr(__ErrNum))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__wcserror = dll.NewProc("__wcserror")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) __wcserror(const wchar_t *_Str);
func X__wcserror(tls *TLS, __Str uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X__wcserror->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__wcserror.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcsicmp = dll.NewProc("_wcsicmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsicmp(const wchar_t *_Str1,const wchar_t *_Str2);
func X_wcsicmp(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`X_wcsicmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsicmp.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsicmp_l = dll.NewProc("_wcsicmp_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsicmp_l(const wchar_t *_Str1,const wchar_t *_Str2,_locale_t _Locale);
func X_wcsicmp_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_wcsicmp_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsicmp_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsnicmp = dll.NewProc("_wcsnicmp")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsnicmp(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount);
func X_wcsnicmp(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_wcsnicmp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsnicmp.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsnicmp_l = dll.NewProc("_wcsnicmp_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsnicmp_l(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_wcsnicmp_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_wcsnicmp_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsnicmp_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsnset = dll.NewProc("_wcsnset")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcsnset(wchar_t *_Str,wchar_t _Val,size_t _MaxCount);
func X_wcsnset(tls *TLS, __Str uintptr, __Val Twchar_t, __MaxCount Tsize_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Val=%+v _MaxCount=%+v", __Str, __Val, __MaxCount)
		defer func() { trc(`X_wcsnset->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsnset.Addr(), __Str, uintptr(__Val), uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcsrev = dll.NewProc("_wcsrev")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcsrev(wchar_t *_Str);
func X_wcsrev(tls *TLS, __Str uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v", __Str)
		defer func() { trc(`X_wcsrev->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsrev.Addr(), __Str)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcsset = dll.NewProc("_wcsset")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcsset(wchar_t *_Str,wchar_t _Val);
func X_wcsset(tls *TLS, __Str uintptr, __Val Twchar_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Str=%+v _Val=%+v", __Str, __Val)
		defer func() { trc(`X_wcsset->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsset.Addr(), __Str, uintptr(__Val))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcslwr = dll.NewProc("_wcslwr")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcslwr(wchar_t *_String);
func X_wcslwr(tls *TLS, __String uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_wcslwr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcslwr.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcslwr_l = dll.NewProc("_wcslwr_l")

// __attribute__ ((__dllimport__)) wchar_t *_wcslwr_l(wchar_t *_String,_locale_t _Locale);
func X_wcslwr_l(tls *TLS, __String uintptr, __Locale T_locale_t) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v _Locale=%+v", __String, __Locale)
		defer func() { trc(`X_wcslwr_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcslwr_l.Addr(), __String, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcsupr = dll.NewProc("_wcsupr")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wcsupr(wchar_t *_String);
func X_wcsupr(tls *TLS, __String uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_wcsupr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsupr.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wcsupr_l = dll.NewProc("_wcsupr_l")

// __attribute__ ((__dllimport__)) wchar_t *_wcsupr_l(wchar_t *_String,_locale_t _Locale);
func X_wcsupr_l(tls *TLS, __String uintptr, __Locale T_locale_t) (r uintptr) {
	if __ccgo_strace {
		trc("_String=%+v _Locale=%+v", __String, __Locale)
		defer func() { trc(`X_wcsupr_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsupr_l.Addr(), __String, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procwcsxfrm = dll.NewProc("wcsxfrm")

// size_t __attribute__((__cdecl__)) wcsxfrm(wchar_t * __restrict__ _Dst,const wchar_t * __restrict__ _Src,size_t _MaxCount);
func Xwcsxfrm(tls *TLS, __Dst uintptr, __Src uintptr, __MaxCount Tsize_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _MaxCount=%+v", __Dst, __Src, __MaxCount)
		defer func() { trc(`Xwcsxfrm->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsxfrm.Addr(), __Dst, __Src, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_wcsxfrm_l = dll.NewProc("_wcsxfrm_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _wcsxfrm_l(wchar_t * __restrict__ _Dst,const wchar_t * __restrict__ _Src,size_t _MaxCount,_locale_t _Locale);
func X_wcsxfrm_l(tls *TLS, __Dst uintptr, __Src uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dst=%+v _Src=%+v _MaxCount=%+v _Locale=%+v", __Dst, __Src, __MaxCount, __Locale)
		defer func() { trc(`X_wcsxfrm_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsxfrm_l.Addr(), __Dst, __Src, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcscoll = dll.NewProc("wcscoll")

// int __attribute__((__cdecl__)) wcscoll(const wchar_t *_Str1,const wchar_t *_Str2);
func Xwcscoll(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`Xwcscoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcscoll.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcscoll_l = dll.NewProc("_wcscoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcscoll_l(const wchar_t *_Str1,const wchar_t *_Str2,_locale_t _Locale);
func X_wcscoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_wcscoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcscoll_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsicoll = dll.NewProc("_wcsicoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsicoll(const wchar_t *_Str1,const wchar_t *_Str2);
func X_wcsicoll(tls *TLS, __Str1 uintptr, __Str2 uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v", __Str1, __Str2)
		defer func() { trc(`X_wcsicoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsicoll.Addr(), __Str1, __Str2)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsicoll_l = dll.NewProc("_wcsicoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsicoll_l(const wchar_t *_Str1,const wchar_t *_Str2,_locale_t _Locale);
func X_wcsicoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _Locale=%+v", __Str1, __Str2, __Locale)
		defer func() { trc(`X_wcsicoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsicoll_l.Addr(), __Str1, __Str2, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsncoll = dll.NewProc("_wcsncoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsncoll(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount);
func X_wcsncoll(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_wcsncoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsncoll.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsncoll_l = dll.NewProc("_wcsncoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsncoll_l(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_wcsncoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_wcsncoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsncoll_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsnicoll = dll.NewProc("_wcsnicoll")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsnicoll(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount);
func X_wcsnicoll(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v", __Str1, __Str2, __MaxCount)
		defer func() { trc(`X_wcsnicoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsnicoll.Addr(), __Str1, __Str2, uintptr(__MaxCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcsnicoll_l = dll.NewProc("_wcsnicoll_l")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcsnicoll_l(const wchar_t *_Str1,const wchar_t *_Str2,size_t _MaxCount,_locale_t _Locale);
func X_wcsnicoll_l(tls *TLS, __Str1 uintptr, __Str2 uintptr, __MaxCount Tsize_t, __Locale T_locale_t) (r int32) {
	if __ccgo_strace {
		trc("_Str1=%+v _Str2=%+v _MaxCount=%+v _Locale=%+v", __Str1, __Str2, __MaxCount, __Locale)
		defer func() { trc(`X_wcsnicoll_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsnicoll_l.Addr(), __Str1, __Str2, uintptr(__MaxCount), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

type T__timeb32 = struct {
	Ftime     T__time32_t
	Fmillitm  uint16
	Ftimezone int16
	Fdstflag  int16
}

type Ttimeb = struct {
	Ftime     Ttime_t
	Fmillitm  uint16
	Ftimezone int16
	Fdstflag  int16
}

type T__timeb64 = struct {
	Ftime     T__time64_t
	Fmillitm  uint16
	Ftimezone int16
	Fdstflag  int16
}

var proc_ftime64 = dll.NewProc("_ftime64")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _ftime64(struct __timeb64 *_Time);
func X_ftime64(tls *TLS, __Time uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
	}
	r0, r1, err := syscall.SyscallN(proc_ftime64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_ftime32 = dll.NewProc("_ftime32")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _ftime32(struct __timeb32 *_Time);
func X_ftime32(tls *TLS, __Time uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
	}
	r0, r1, err := syscall.SyscallN(proc_ftime32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

type T_timespec32 = struct {
	Ftv_sec  T__time32_t
	Ftv_nsec int32
}

type T_timespec64 = struct {
	Ftv_sec  T__time64_t
	Ftv_nsec int32
}

type Ttimespec = struct {
	Ftv_sec  Ttime_t
	Ftv_nsec int32
}

type Titimerspec = struct {
	Fit_interval Ttimespec
	Fit_value    Ttimespec
}

type Tclock_t = int32

type Ttm = struct {
	Ftm_sec   int32
	Ftm_min   int32
	Ftm_hour  int32
	Ftm_mday  int32
	Ftm_mon   int32
	Ftm_year  int32
	Ftm_wday  int32
	Ftm_yday  int32
	Ftm_isdst int32
}

var proc__daylight = dll.NewProc("__daylight")

// __attribute__ ((__dllimport__)) int * __attribute__((__cdecl__)) __daylight(void);
func X__daylight(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__daylight->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__daylight.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__dstbias = dll.NewProc("__dstbias")

// __attribute__ ((__dllimport__)) long * __attribute__((__cdecl__)) __dstbias(void);
func X__dstbias(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__dstbias->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__dstbias.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__timezone = dll.NewProc("__timezone")

// __attribute__ ((__dllimport__)) long * __attribute__((__cdecl__)) __timezone(void);
func X__timezone(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__timezone->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__timezone.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc__tzname = dll.NewProc("__tzname")

// __attribute__ ((__dllimport__)) char ** __attribute__((__cdecl__)) __tzname(void);
func X__tzname(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__tzname->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__tzname.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_get_daylight = dll.NewProc("_get_daylight")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _get_daylight(int *_Daylight);
func X_get_daylight(tls *TLS, __Daylight uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Daylight=%+v", __Daylight)
		defer func() { trc(`X_get_daylight->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_daylight.Addr(), __Daylight)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_dstbias = dll.NewProc("_get_dstbias")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _get_dstbias(long *_Daylight_savings_bias);
func X_get_dstbias(tls *TLS, __Daylight_savings_bias uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Daylight_savings_bias=%+v", __Daylight_savings_bias)
		defer func() { trc(`X_get_dstbias->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_dstbias.Addr(), __Daylight_savings_bias)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_timezone = dll.NewProc("_get_timezone")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _get_timezone(long *_Timezone);
func X_get_timezone(tls *TLS, __Timezone uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Timezone=%+v", __Timezone)
		defer func() { trc(`X_get_timezone->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_timezone.Addr(), __Timezone)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_tzname = dll.NewProc("_get_tzname")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _get_tzname(size_t *_ReturnValue,char *_Buffer,size_t _SizeInBytes,int _Index);
func X_get_tzname(tls *TLS, __ReturnValue uintptr, __Buffer uintptr, __SizeInBytes Tsize_t, __Index int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_ReturnValue=%+v _Buffer=%+v _SizeInBytes=%+v _Index=%+v", __ReturnValue, __Buffer, __SizeInBytes, __Index)
		defer func() { trc(`X_get_tzname->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_tzname.Addr(), __ReturnValue, __Buffer, uintptr(__SizeInBytes), uintptr(__Index))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var procasctime = dll.NewProc("asctime")

// char * __attribute__((__cdecl__)) asctime(const struct tm *_Tm);
func Xasctime(tls *TLS, __Tm uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`Xasctime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procasctime.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var procasctime_s = dll.NewProc("asctime_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) asctime_s (char *_Buf,size_t _SizeInWords,const struct tm *_Tm);
func Xasctime_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Tm uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Tm=%+v", __Buf, __SizeInWords, __Tm)
		defer func() { trc(`Xasctime_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procasctime_s.Addr(), __Buf, uintptr(__SizeInWords), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_ctime32 = dll.NewProc("_ctime32")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _ctime32(const __time32_t *_Time);
func X_ctime32(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_ctime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ctime32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ctime32_s = dll.NewProc("_ctime32_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _ctime32_s (char *_Buf,size_t _SizeInBytes,const __time32_t *_Time);
func X_ctime32_s(tls *TLS, __Buf uintptr, __SizeInBytes Tsize_t, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInBytes=%+v _Time=%+v", __Buf, __SizeInBytes, __Time)
		defer func() { trc(`X_ctime32_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ctime32_s.Addr(), __Buf, uintptr(__SizeInBytes), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var procclock = dll.NewProc("clock")

// clock_t __attribute__((__cdecl__)) clock(void);
func Xclock(tls *TLS) (r Tclock_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xclock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procclock.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tclock_t(r0)
}

var proc_gmtime32 = dll.NewProc("_gmtime32")

// __attribute__ ((__dllimport__)) struct tm * __attribute__((__cdecl__)) _gmtime32(const __time32_t *_Time);
func X_gmtime32(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_gmtime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_gmtime32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_gmtime32_s = dll.NewProc("_gmtime32_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _gmtime32_s (struct tm *_Tm,const __time32_t *_Time);
func X_gmtime32_s(tls *TLS, __Tm uintptr, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Tm=%+v _Time=%+v", __Tm, __Time)
		defer func() { trc(`X_gmtime32_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_gmtime32_s.Addr(), __Tm, __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_localtime32 = dll.NewProc("_localtime32")

// __attribute__ ((__dllimport__)) struct tm * __attribute__((__cdecl__)) _localtime32(const __time32_t *_Time);
func X_localtime32(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_localtime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_localtime32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_localtime32_s = dll.NewProc("_localtime32_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _localtime32_s (struct tm *_Tm,const __time32_t *_Time);
func X_localtime32_s(tls *TLS, __Tm uintptr, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Tm=%+v _Time=%+v", __Tm, __Time)
		defer func() { trc(`X_localtime32_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_localtime32_s.Addr(), __Tm, __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var procstrftime = dll.NewProc("strftime")

// size_t __attribute__((__cdecl__)) strftime(char * __restrict__ _Buf,size_t _SizeInBytes,const char * __restrict__ _Format,const struct tm * __restrict__ _Tm);
func Xstrftime(tls *TLS, __Buf uintptr, __SizeInBytes Tsize_t, __Format uintptr, __Tm uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInBytes=%+v _Format=%+v _Tm=%+v", __Buf, __SizeInBytes, __Format, __Tm)
		defer func() { trc(`Xstrftime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procstrftime.Addr(), __Buf, uintptr(__SizeInBytes), __Format, __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_strftime_l = dll.NewProc("_strftime_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _strftime_l(char * __restrict__ _Buf,size_t _Max_size,const char * __restrict__ _Format,const struct tm * __restrict__ _Tm,_locale_t _Locale);
func X_strftime_l(tls *TLS, __Buf uintptr, __Max_size Tsize_t, __Format uintptr, __Tm uintptr, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _Max_size=%+v _Format=%+v _Tm=%+v _Locale=%+v", __Buf, __Max_size, __Format, __Tm, __Locale)
		defer func() { trc(`X_strftime_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strftime_l.Addr(), __Buf, uintptr(__Max_size), __Format, __Tm, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_strdate = dll.NewProc("_strdate")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strdate(char *_Buffer);
func X_strdate(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_strdate->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strdate.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strdate_s = dll.NewProc("_strdate_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _strdate_s (char *_Buf,size_t _SizeInBytes);
func X_strdate_s(tls *TLS, __Buf uintptr, __SizeInBytes Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInBytes=%+v", __Buf, __SizeInBytes)
		defer func() { trc(`X_strdate_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strdate_s.Addr(), __Buf, uintptr(__SizeInBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_strtime = dll.NewProc("_strtime")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _strtime(char *_Buffer);
func X_strtime(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_strtime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtime.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_strtime_s = dll.NewProc("_strtime_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _strtime_s (char *_Buf ,size_t _SizeInBytes);
func X_strtime_s(tls *TLS, __Buf uintptr, __SizeInBytes Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInBytes=%+v", __Buf, __SizeInBytes)
		defer func() { trc(`X_strtime_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_strtime_s.Addr(), __Buf, uintptr(__SizeInBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_time32 = dll.NewProc("_time32")

// __attribute__ ((__dllimport__)) __time32_t __attribute__((__cdecl__)) _time32(__time32_t *_Time);
func X_time32(tls *TLS, __Time uintptr) (r T__time32_t) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_time32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_time32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time32_t(r0)
}

var proc_timespec32_get = dll.NewProc("_timespec32_get")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _timespec32_get(struct _timespec32 *_Ts, int _Base);
func X_timespec32_get(tls *TLS, __Ts uintptr, __Base int32) (r int32) {
	if __ccgo_strace {
		trc("_Ts=%+v _Base=%+v", __Ts, __Base)
		defer func() { trc(`X_timespec32_get->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_timespec32_get.Addr(), __Ts, uintptr(__Base))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_mktime32 = dll.NewProc("_mktime32")

// __attribute__ ((__dllimport__)) __time32_t __attribute__((__cdecl__)) _mktime32(struct tm *_Tm);
func X_mktime32(tls *TLS, __Tm uintptr) (r T__time32_t) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_mktime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mktime32.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time32_t(r0)
}

var proc_mkgmtime32 = dll.NewProc("_mkgmtime32")

// __attribute__ ((__dllimport__)) __time32_t __attribute__((__cdecl__)) _mkgmtime32(struct tm *_Tm);
func X_mkgmtime32(tls *TLS, __Tm uintptr) (r T__time32_t) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_mkgmtime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mkgmtime32.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time32_t(r0)
}

var proc_tzset = dll.NewProc("_tzset")

// void __attribute__((__cdecl__)) _tzset(void);
func X_tzset(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(proc_tzset.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_ctime64 = dll.NewProc("_ctime64")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _ctime64(const __time64_t *_Time);
func X_ctime64(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_ctime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ctime64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_ctime64_s = dll.NewProc("_ctime64_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _ctime64_s (char *_Buf,size_t _SizeInBytes,const __time64_t *_Time);
func X_ctime64_s(tls *TLS, __Buf uintptr, __SizeInBytes Tsize_t, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInBytes=%+v _Time=%+v", __Buf, __SizeInBytes, __Time)
		defer func() { trc(`X_ctime64_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ctime64_s.Addr(), __Buf, uintptr(__SizeInBytes), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_gmtime64 = dll.NewProc("_gmtime64")

// __attribute__ ((__dllimport__)) struct tm * __attribute__((__cdecl__)) _gmtime64(const __time64_t *_Time);
func X_gmtime64(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_gmtime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_gmtime64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_gmtime64_s = dll.NewProc("_gmtime64_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _gmtime64_s (struct tm *_Tm,const __time64_t *_Time);
func X_gmtime64_s(tls *TLS, __Tm uintptr, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Tm=%+v _Time=%+v", __Tm, __Time)
		defer func() { trc(`X_gmtime64_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_gmtime64_s.Addr(), __Tm, __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_localtime64 = dll.NewProc("_localtime64")

// __attribute__ ((__dllimport__)) struct tm * __attribute__((__cdecl__)) _localtime64(const __time64_t *_Time);
func X_localtime64(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_localtime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_localtime64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_localtime64_s = dll.NewProc("_localtime64_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _localtime64_s (struct tm *_Tm,const __time64_t *_Time);
func X_localtime64_s(tls *TLS, __Tm uintptr, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Tm=%+v _Time=%+v", __Tm, __Time)
		defer func() { trc(`X_localtime64_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_localtime64_s.Addr(), __Tm, __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_mktime64 = dll.NewProc("_mktime64")

// __attribute__ ((__dllimport__)) __time64_t __attribute__((__cdecl__)) _mktime64(struct tm *_Tm);
func X_mktime64(tls *TLS, __Tm uintptr) (r T__time64_t) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_mktime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mktime64.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time64_t(r0)
}

var proc_mkgmtime64 = dll.NewProc("_mkgmtime64")

// __attribute__ ((__dllimport__)) __time64_t __attribute__((__cdecl__)) _mkgmtime64(struct tm *_Tm);
func X_mkgmtime64(tls *TLS, __Tm uintptr) (r T__time64_t) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_mkgmtime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mkgmtime64.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time64_t(r0)
}

var proc_time64 = dll.NewProc("_time64")

// __attribute__ ((__dllimport__)) __time64_t __attribute__((__cdecl__)) _time64(__time64_t *_Time);
func X_time64(tls *TLS, __Time uintptr) (r T__time64_t) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_time64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_time64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T__time64_t(r0)
}

var proc_timespec64_get = dll.NewProc("_timespec64_get")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _timespec64_get(struct _timespec64 *_Ts, int _Base);
func X_timespec64_get(tls *TLS, __Ts uintptr, __Base int32) (r int32) {
	if __ccgo_strace {
		trc("_Ts=%+v _Base=%+v", __Ts, __Base)
		defer func() { trc(`X_timespec64_get->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_timespec64_get.Addr(), __Ts, uintptr(__Base))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_getsystime = dll.NewProc("_getsystime")

// unsigned __attribute__((__cdecl__)) _getsystime(struct tm *_Tm);
func X_getsystime(tls *TLS, __Tm uintptr) (r uint32) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_getsystime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getsystime.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_setsystime = dll.NewProc("_setsystime")

// unsigned __attribute__((__cdecl__)) _setsystime(struct tm *_Tm,unsigned _MilliSec);
func X_setsystime(tls *TLS, __Tm uintptr, __MilliSec uint32) (r uint32) {
	if __ccgo_strace {
		trc("_Tm=%+v _MilliSec=%+v", __Tm, __MilliSec)
		defer func() { trc(`X_setsystime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_setsystime.Addr(), __Tm, uintptr(__MilliSec))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_wasctime = dll.NewProc("_wasctime")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wasctime(const struct tm *_Tm);
func X_wasctime(tls *TLS, __Tm uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Tm=%+v", __Tm)
		defer func() { trc(`X_wasctime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wasctime.Addr(), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wasctime_s = dll.NewProc("_wasctime_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _wasctime_s (wchar_t *_Buf,size_t _SizeInWords,const struct tm *_Tm);
func X_wasctime_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Tm uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Tm=%+v", __Buf, __SizeInWords, __Tm)
		defer func() { trc(`X_wasctime_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wasctime_s.Addr(), __Buf, uintptr(__SizeInWords), __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_wctime32 = dll.NewProc("_wctime32")

// wchar_t * __attribute__((__cdecl__)) _wctime32(const __time32_t *_Time);
func X_wctime32(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_wctime32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wctime32.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wctime32_s = dll.NewProc("_wctime32_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _wctime32_s (wchar_t *_Buf,size_t _SizeInWords,const __time32_t *_Time);
func X_wctime32_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Time=%+v", __Buf, __SizeInWords, __Time)
		defer func() { trc(`X_wctime32_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wctime32_s.Addr(), __Buf, uintptr(__SizeInWords), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var procwcsftime = dll.NewProc("wcsftime")

// size_t __attribute__((__cdecl__)) wcsftime(wchar_t * __restrict__ _Buf,size_t _SizeInWords,const wchar_t * __restrict__ _Format,const struct tm * __restrict__ _Tm);
func Xwcsftime(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Format uintptr, __Tm uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Format=%+v _Tm=%+v", __Buf, __SizeInWords, __Format, __Tm)
		defer func() { trc(`Xwcsftime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsftime.Addr(), __Buf, uintptr(__SizeInWords), __Format, __Tm)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_wcsftime_l = dll.NewProc("_wcsftime_l")

// __attribute__ ((__dllimport__)) size_t __attribute__((__cdecl__)) _wcsftime_l(wchar_t * __restrict__ _Buf,size_t _SizeInWords,const wchar_t * __restrict__ _Format,const struct tm * __restrict__ _Tm,_locale_t _Locale);
func X_wcsftime_l(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Format uintptr, __Tm uintptr, __Locale T_locale_t) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Format=%+v _Tm=%+v _Locale=%+v", __Buf, __SizeInWords, __Format, __Tm, __Locale)
		defer func() { trc(`X_wcsftime_l->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcsftime_l.Addr(), __Buf, uintptr(__SizeInWords), __Format, __Tm, __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var proc_wstrdate = dll.NewProc("_wstrdate")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wstrdate(wchar_t *_Buffer);
func X_wstrdate(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_wstrdate->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstrdate.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wstrdate_s = dll.NewProc("_wstrdate_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _wstrdate_s (wchar_t *_Buf,size_t _SizeInWords);
func X_wstrdate_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v", __Buf, __SizeInWords)
		defer func() { trc(`X_wstrdate_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstrdate_s.Addr(), __Buf, uintptr(__SizeInWords))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_wstrtime = dll.NewProc("_wstrtime")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wstrtime(wchar_t *_Buffer);
func X_wstrtime(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_wstrtime->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstrtime.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wstrtime_s = dll.NewProc("_wstrtime_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _wstrtime_s (wchar_t *_Buf,size_t _SizeInWords);
func X_wstrtime_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v", __Buf, __SizeInWords)
		defer func() { trc(`X_wstrtime_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstrtime_s.Addr(), __Buf, uintptr(__SizeInWords))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_wctime64 = dll.NewProc("_wctime64")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wctime64(const __time64_t *_Time);
func X_wctime64(tls *TLS, __Time uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Time=%+v", __Time)
		defer func() { trc(`X_wctime64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wctime64.Addr(), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wctime64_s = dll.NewProc("_wctime64_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _wctime64_s (wchar_t *_Buf,size_t _SizeInWords,const __time64_t *_Time);
func X_wctime64_s(tls *TLS, __Buf uintptr, __SizeInWords Tsize_t, __Time uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Buf=%+v _SizeInWords=%+v _Time=%+v", __Buf, __SizeInWords, __Time)
		defer func() { trc(`X_wctime64_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wctime64_s.Addr(), __Buf, uintptr(__SizeInWords), __Time)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

type Ttimeval = struct {
	Ftv_sec  int32
	Ftv_usec int32
}

type Ttimezone = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type Tclockid_t = int32

type T_fsize_t = uint32

type T_wfinddata32_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        T_fsize_t
	Fname        [260]Twchar_t
}

type T_wfinddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        int64
	Fname        [260]Twchar_t
}

type T_wfinddata64i32_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        T_fsize_t
	Fname        [260]Twchar_t
}

type T_wfinddata64_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        int64
	Fname        [260]Twchar_t
}

var proc_wgetcwd = dll.NewProc("_wgetcwd")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wgetcwd(wchar_t *_DstBuf,int _SizeInWords);
func X_wgetcwd(tls *TLS, __DstBuf uintptr, __SizeInWords int32) (r uintptr) {
	if __ccgo_strace {
		trc("_DstBuf=%+v _SizeInWords=%+v", __DstBuf, __SizeInWords)
		defer func() { trc(`X_wgetcwd->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wgetcwd.Addr(), __DstBuf, uintptr(__SizeInWords))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wgetdcwd = dll.NewProc("_wgetdcwd")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wgetdcwd(int _Drive,wchar_t *_DstBuf,int _SizeInWords);
func X_wgetdcwd(tls *TLS, __Drive int32, __DstBuf uintptr, __SizeInWords int32) (r uintptr) {
	if __ccgo_strace {
		trc("_Drive=%+v _DstBuf=%+v _SizeInWords=%+v", __Drive, __DstBuf, __SizeInWords)
		defer func() { trc(`X_wgetdcwd->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wgetdcwd.Addr(), uintptr(__Drive), __DstBuf, uintptr(__SizeInWords))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wchdir = dll.NewProc("_wchdir")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wchdir(const wchar_t *_Path);
func X_wchdir(tls *TLS, __Path uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Path=%+v", __Path)
		defer func() { trc(`X_wchdir->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wchdir.Addr(), __Path)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wmkdir = dll.NewProc("_wmkdir")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wmkdir(const wchar_t *_Path);
func X_wmkdir(tls *TLS, __Path uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Path=%+v", __Path)
		defer func() { trc(`X_wmkdir->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wmkdir.Addr(), __Path)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wrmdir = dll.NewProc("_wrmdir")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wrmdir(const wchar_t *_Path);
func X_wrmdir(tls *TLS, __Path uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Path=%+v", __Path)
		defer func() { trc(`X_wrmdir->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wrmdir.Addr(), __Path)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_waccess = dll.NewProc("_waccess")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _waccess(const wchar_t *_Filename,int _AccessMode);
func X_waccess(tls *TLS, __Filename uintptr, __AccessMode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _AccessMode=%+v", __Filename, __AccessMode)
		defer func() { trc(`X_waccess->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_waccess.Addr(), __Filename, uintptr(__AccessMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wchmod = dll.NewProc("_wchmod")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wchmod(const wchar_t *_Filename,int _Mode);
func X_wchmod(tls *TLS, __Filename uintptr, __Mode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v", __Filename, __Mode)
		defer func() { trc(`X_wchmod->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wchmod.Addr(), __Filename, uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wcreat = dll.NewProc("_wcreat")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wcreat(const wchar_t *_Filename,int _PermissionMode);
func X_wcreat(tls *TLS, __Filename uintptr, __PermissionMode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _PermissionMode=%+v", __Filename, __PermissionMode)
		defer func() { trc(`X_wcreat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wcreat.Addr(), __Filename, uintptr(__PermissionMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wfindfirst32 = dll.NewProc("_wfindfirst32")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wfindfirst32(const wchar_t *_Filename,struct _wfinddata32_t *_FindData);
func X_wfindfirst32(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_wfindfirst32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindfirst32.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wfindnext32 = dll.NewProc("_wfindnext32")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wfindnext32(intptr_t _FindHandle,struct _wfinddata32_t *_FindData);
func X_wfindnext32(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_wfindnext32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindnext32.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wunlink = dll.NewProc("_wunlink")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wunlink(const wchar_t *_Filename);
func X_wunlink(tls *TLS, __Filename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v", __Filename)
		defer func() { trc(`X_wunlink->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wunlink.Addr(), __Filename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wrename = dll.NewProc("_wrename")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wrename(const wchar_t *_OldFilename,const wchar_t *_NewFilename);
func X_wrename(tls *TLS, __OldFilename uintptr, __NewFilename uintptr) (r int32) {
	if __ccgo_strace {
		trc("_OldFilename=%+v _NewFilename=%+v", __OldFilename, __NewFilename)
		defer func() { trc(`X_wrename->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wrename.Addr(), __OldFilename, __NewFilename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wmktemp = dll.NewProc("_wmktemp")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wmktemp(wchar_t *_TemplateName);
func X_wmktemp(tls *TLS, __TemplateName uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_TemplateName=%+v", __TemplateName)
		defer func() { trc(`X_wmktemp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wmktemp.Addr(), __TemplateName)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wfindfirst32i64 = dll.NewProc("_wfindfirst32i64")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wfindfirst32i64(const wchar_t *_Filename,struct _wfinddata32i64_t *_FindData);
func X_wfindfirst32i64(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_wfindfirst32i64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindfirst32i64.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wfindfirst64i32 = dll.NewProc("_wfindfirst64i32")

// intptr_t __attribute__((__cdecl__)) _wfindfirst64i32(const wchar_t *_Filename,struct _wfinddata64i32_t *_FindData);
func X_wfindfirst64i32(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_wfindfirst64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindfirst64i32.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wfindfirst64 = dll.NewProc("_wfindfirst64")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wfindfirst64(const wchar_t *_Filename,struct _wfinddata64_t *_FindData);
func X_wfindfirst64(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_wfindfirst64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindfirst64.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wfindnext32i64 = dll.NewProc("_wfindnext32i64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wfindnext32i64(intptr_t _FindHandle,struct _wfinddata32i64_t *_FindData);
func X_wfindnext32i64(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_wfindnext32i64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindnext32i64.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wfindnext64i32 = dll.NewProc("_wfindnext64i32")

// int __attribute__((__cdecl__)) _wfindnext64i32(intptr_t _FindHandle,struct _wfinddata64i32_t *_FindData);
func X_wfindnext64i32(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_wfindnext64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindnext64i32.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wfindnext64 = dll.NewProc("_wfindnext64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wfindnext64(intptr_t _FindHandle,struct _wfinddata64_t *_FindData);
func X_wfindnext64(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_wfindnext64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wfindnext64.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wsopen_s = dll.NewProc("_wsopen_s")

// __attribute__ ((__dllimport__)) errno_t __attribute__((__cdecl__)) _wsopen_s(int *_FileHandle,const wchar_t *_Filename,int _OpenFlag,int _ShareFlag,int _PermissionFlag);
func X_wsopen_s(tls *TLS, __FileHandle uintptr, __Filename uintptr, __OpenFlag int32, __ShareFlag int32, __PermissionFlag int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Filename=%+v _OpenFlag=%+v _ShareFlag=%+v _PermissionFlag=%+v", __FileHandle, __Filename, __OpenFlag, __ShareFlag, __PermissionFlag)
		defer func() { trc(`X_wsopen_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wsopen_s.Addr(), __FileHandle, __Filename, uintptr(__OpenFlag), uintptr(__ShareFlag), uintptr(__PermissionFlag))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_wsetlocale = dll.NewProc("_wsetlocale")

// __attribute__ ((__dllimport__)) wchar_t * __attribute__((__cdecl__)) _wsetlocale(int _Category,const wchar_t *_Locale);
func X_wsetlocale(tls *TLS, __Category int32, __Locale uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Category=%+v _Locale=%+v", __Category, __Locale)
		defer func() { trc(`X_wsetlocale->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wsetlocale.Addr(), uintptr(__Category), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_wexecv = dll.NewProc("_wexecv")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wexecv(const wchar_t *_Filename,const wchar_t *const *_ArgList);
func X_wexecv(tls *TLS, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v", __Filename, __ArgList)
		defer func() { trc(`X_wexecv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wexecv.Addr(), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wexecve = dll.NewProc("_wexecve")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wexecve(const wchar_t *_Filename,const wchar_t *const *_ArgList,const wchar_t *const *_Env);
func X_wexecve(tls *TLS, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v _Env=%+v", __Filename, __ArgList, __Env)
		defer func() { trc(`X_wexecve->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wexecve.Addr(), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wexecvp = dll.NewProc("_wexecvp")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wexecvp(const wchar_t *_Filename,const wchar_t *const *_ArgList);
func X_wexecvp(tls *TLS, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v", __Filename, __ArgList)
		defer func() { trc(`X_wexecvp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wexecvp.Addr(), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_wexecvpe = dll.NewProc("_wexecvpe")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _wexecvpe(const wchar_t *_Filename,const wchar_t *const *_ArgList,const wchar_t *const *_Env);
func X_wexecvpe(tls *TLS, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v _Env=%+v", __Filename, __ArgList, __Env)
		defer func() { trc(`X_wexecvpe->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wexecvpe.Addr(), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

type T_ino_t = uint16

type Tino_t = uint16

type T_dev_t = uint32

type Tdev_t = uint32

type T_stat32 = struct {
	Fst_dev   T_dev_t
	Fst_ino   T_ino_t
	Fst_mode  uint16
	Fst_nlink int16
	Fst_uid   int16
	Fst_gid   int16
	Fst_rdev  T_dev_t
	Fst_size  T_off_t
	Fst_atime T__time32_t
	Fst_mtime T__time32_t
	Fst_ctime T__time32_t
}

type Tstat = struct {
	Fst_dev   T_dev_t
	Fst_ino   T_ino_t
	Fst_mode  uint16
	Fst_nlink int16
	Fst_uid   int16
	Fst_gid   int16
	Fst_rdev  T_dev_t
	Fst_size  T_off_t
	Fst_atime Ttime_t
	Fst_mtime Ttime_t
	Fst_ctime Ttime_t
}

type T_stat32i64 = struct {
	Fst_dev   T_dev_t
	Fst_ino   T_ino_t
	Fst_mode  uint16
	Fst_nlink int16
	Fst_uid   int16
	Fst_gid   int16
	Fst_rdev  T_dev_t
	Fst_size  int64
	Fst_atime T__time32_t
	Fst_mtime T__time32_t
	Fst_ctime T__time32_t
}

type T_stat64i32 = struct {
	Fst_dev   T_dev_t
	Fst_ino   T_ino_t
	Fst_mode  uint16
	Fst_nlink int16
	Fst_uid   int16
	Fst_gid   int16
	Fst_rdev  T_dev_t
	Fst_size  T_off_t
	Fst_atime T__time64_t
	Fst_mtime T__time64_t
	Fst_ctime T__time64_t
}

type T_stat64 = struct {
	Fst_dev   T_dev_t
	Fst_ino   T_ino_t
	Fst_mode  uint16
	Fst_nlink int16
	Fst_uid   int16
	Fst_gid   int16
	Fst_rdev  T_dev_t
	Fst_size  int64
	Fst_atime T__time64_t
	Fst_mtime T__time64_t
	Fst_ctime T__time64_t
}

var proc_wstat32 = dll.NewProc("_wstat32")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wstat32(const wchar_t *_Name,struct _stat32 *_Stat);
func X_wstat32(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_wstat32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstat32.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wstat32i64 = dll.NewProc("_wstat32i64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wstat32i64(const wchar_t *_Name,struct _stat32i64 *_Stat);
func X_wstat32i64(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_wstat32i64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstat32i64.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wstat64i32 = dll.NewProc("_wstat64i32")

// int __attribute__((__cdecl__)) _wstat64i32(const wchar_t *_Name,struct _stat64i32 *_Stat);
func X_wstat64i32(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_wstat64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstat64i32.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_wstat64 = dll.NewProc("_wstat64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _wstat64(const wchar_t *_Name,struct _stat64 *_Stat);
func X_wstat64(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_wstat64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_wstat64.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_cgetws = dll.NewProc("_cgetws")

// __attribute__ ((__dllimport__)) wchar_t *_cgetws(wchar_t *_Buffer);
func X_cgetws(tls *TLS, __Buffer uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Buffer=%+v", __Buffer)
		defer func() { trc(`X_cgetws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_cgetws.Addr(), __Buffer)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_getwch = dll.NewProc("_getwch")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _getwch(void);
func X_getwch(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getwch->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getwch.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_getwche = dll.NewProc("_getwche")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _getwche(void);
func X_getwche(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getwche->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getwche.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_putwch = dll.NewProc("_putwch")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _putwch(wchar_t _WCh);
func X_putwch(tls *TLS, __WCh Twchar_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_WCh=%+v", __WCh)
		defer func() { trc(`X_putwch->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putwch.Addr(), uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_ungetwch = dll.NewProc("_ungetwch")

// __attribute__ ((__dllimport__)) wint_t __attribute__((__cdecl__)) _ungetwch(wint_t _WCh);
func X_ungetwch(tls *TLS, __WCh Twint_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_WCh=%+v", __WCh)
		defer func() { trc(`X_ungetwch->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ungetwch.Addr(), uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_cputws = dll.NewProc("_cputws")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _cputws(const wchar_t *_String);
func X_cputws(tls *TLS, __String uintptr) (r int32) {
	if __ccgo_strace {
		trc("_String=%+v", __String)
		defer func() { trc(`X_cputws->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_cputws.Addr(), __String)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_putwch_nolock = dll.NewProc("_putwch_nolock")

// wint_t __attribute__((__cdecl__)) _putwch_nolock(wchar_t _WCh);
func X_putwch_nolock(tls *TLS, __WCh Twchar_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_WCh=%+v", __WCh)
		defer func() { trc(`X_putwch_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_putwch_nolock.Addr(), uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_getwch_nolock = dll.NewProc("_getwch_nolock")

// wint_t __attribute__((__cdecl__)) _getwch_nolock(void);
func X_getwch_nolock(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getwch_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getwch_nolock.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_getwche_nolock = dll.NewProc("_getwche_nolock")

// wint_t __attribute__((__cdecl__)) _getwche_nolock(void);
func X_getwche_nolock(tls *TLS) (r Twint_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getwche_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getwche_nolock.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var proc_ungetwch_nolock = dll.NewProc("_ungetwch_nolock")

// wint_t __attribute__((__cdecl__)) _ungetwch_nolock(wint_t _WCh);
func X_ungetwch_nolock(tls *TLS, __WCh Twint_t) (r Twint_t) {
	if __ccgo_strace {
		trc("_WCh=%+v", __WCh)
		defer func() { trc(`X_ungetwch_nolock->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_ungetwch_nolock.Addr(), uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

type T_Mbstatet = struct {
	F_Wchar uint32
	F_Byte  uint16
	F_State uint16
}

type Tmbstate_t = struct {
	F_Wchar uint32
	F_Byte  uint16
	F_State uint16
}

type T_Mbstatet1 = Tmbstate_t

type T_Wint_t = uint16

var procbtowc = dll.NewProc("btowc")

// wint_t __attribute__((__cdecl__)) btowc(int);
func Xbtowc(tls *TLS, _0 int32) (r Twint_t) {
	if __ccgo_strace {
		trc("0=%+v", _0)
		defer func() { trc(`Xbtowc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procbtowc.Addr(), uintptr(_0))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Twint_t(r0)
}

var procmbrlen = dll.NewProc("mbrlen")

// size_t __attribute__((__cdecl__)) mbrlen(const char * __restrict__ _Ch,size_t _SizeInBytes,mbstate_t * __restrict__ _State);
func Xmbrlen(tls *TLS, __Ch uintptr, __SizeInBytes Tsize_t, __State uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Ch=%+v _SizeInBytes=%+v _State=%+v", __Ch, __SizeInBytes, __State)
		defer func() { trc(`Xmbrlen->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmbrlen.Addr(), __Ch, uintptr(__SizeInBytes), __State)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procmbrtowc = dll.NewProc("mbrtowc")

// size_t __attribute__((__cdecl__)) mbrtowc(wchar_t * __restrict__ _DstCh,const char * __restrict__ _SrcCh,size_t _SizeInBytes,mbstate_t * __restrict__ _State);
func Xmbrtowc(tls *TLS, __DstCh uintptr, __SrcCh uintptr, __SizeInBytes Tsize_t, __State uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_DstCh=%+v _SrcCh=%+v _SizeInBytes=%+v _State=%+v", __DstCh, __SrcCh, __SizeInBytes, __State)
		defer func() { trc(`Xmbrtowc->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmbrtowc.Addr(), __DstCh, __SrcCh, uintptr(__SizeInBytes), __State)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procmbsrtowcs = dll.NewProc("mbsrtowcs")

// size_t __attribute__((__cdecl__)) mbsrtowcs(wchar_t * __restrict__ _Dest,const char ** __restrict__ _PSrc,size_t _Count,mbstate_t * __restrict__ _State);
func Xmbsrtowcs(tls *TLS, __Dest uintptr, __PSrc uintptr, __Count Tsize_t, __State uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _PSrc=%+v _Count=%+v _State=%+v", __Dest, __PSrc, __Count, __State)
		defer func() { trc(`Xmbsrtowcs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procmbsrtowcs.Addr(), __Dest, __PSrc, uintptr(__Count), __State)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcrtomb = dll.NewProc("wcrtomb")

// size_t __attribute__((__cdecl__)) wcrtomb(char * __restrict__ _Dest,wchar_t _Source,mbstate_t * __restrict__ _State);
func Xwcrtomb(tls *TLS, __Dest uintptr, __Source Twchar_t, __State uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _Source=%+v _State=%+v", __Dest, __Source, __State)
		defer func() { trc(`Xwcrtomb->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcrtomb.Addr(), __Dest, uintptr(__Source), __State)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwcsrtombs = dll.NewProc("wcsrtombs")

// size_t __attribute__((__cdecl__)) wcsrtombs(char * __restrict__ _Dest,const wchar_t ** __restrict__ _PSource,size_t _Count,mbstate_t * __restrict__ _State);
func Xwcsrtombs(tls *TLS, __Dest uintptr, __PSource uintptr, __Count Tsize_t, __State uintptr) (r Tsize_t) {
	if __ccgo_strace {
		trc("_Dest=%+v _PSource=%+v _Count=%+v _State=%+v", __Dest, __PSource, __Count, __State)
		defer func() { trc(`Xwcsrtombs->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcsrtombs.Addr(), __Dest, __PSource, uintptr(__Count), __State)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tsize_t(r0)
}

var procwctob = dll.NewProc("wctob")

// int __attribute__((__cdecl__)) wctob(wint_t _WCh);
func Xwctob(tls *TLS, __WCh Twint_t) (r int32) {
	if __ccgo_strace {
		trc("_WCh=%+v", __WCh)
		defer func() { trc(`Xwctob->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwctob.Addr(), uintptr(__WCh))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procwcstoll = dll.NewProc("wcstoll")

// long long __attribute__((__cdecl__)) wcstoll(const wchar_t * __restrict__ nptr,wchar_t ** __restrict__ endptr, int base);
func Xwcstoll(tls *TLS, _nptr uintptr, _endptr uintptr, _base int32) (r int64) {
	if __ccgo_strace {
		trc("nptr=%+v endptr=%+v base=%+v", _nptr, _endptr, _base)
		defer func() { trc(`Xwcstoll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstoll.Addr(), _nptr, _endptr, uintptr(_base))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var procwcstoull = dll.NewProc("wcstoull")

// unsigned long long __attribute__((__cdecl__)) wcstoull(const wchar_t * __restrict__ nptr,wchar_t ** __restrict__ endptr, int base);
func Xwcstoull(tls *TLS, _nptr uintptr, _endptr uintptr, _base int32) (r uint64) {
	if __ccgo_strace {
		trc("nptr=%+v endptr=%+v base=%+v", _nptr, _endptr, _base)
		defer func() { trc(`Xwcstoull->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procwcstoull.Addr(), _nptr, _endptr, uintptr(_base))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint64(r0)
}

var proc_getcwd = dll.NewProc("_getcwd")

// __attribute__ ((__dllimport__)) char* __attribute__((__cdecl__)) _getcwd (char*, int);
func X_getcwd(tls *TLS, _0 uintptr, _1 int32) (r uintptr) {
	if __ccgo_strace {
		trc("0=%+v 1=%+v", _0, _1)
		defer func() { trc(`X_getcwd->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getcwd.Addr(), _0, uintptr(_1))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

type T_finddata32_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        T_fsize_t
	Fname        [260]int8
}

type T_finddata32i64_t = struct {
	Fattrib      uint32
	Ftime_create T__time32_t
	Ftime_access T__time32_t
	Ftime_write  T__time32_t
	Fsize        int64
	Fname        [260]int8
}

type T_finddata64i32_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        T_fsize_t
	Fname        [260]int8
}

type T__finddata64_t = struct {
	Fattrib      uint32
	Ftime_create T__time64_t
	Ftime_access T__time64_t
	Ftime_write  T__time64_t
	Fsize        int64
	Fname        [260]int8
}

var proc_access = dll.NewProc("_access")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _access(const char *_Filename,int _AccessMode);
func X_access(tls *TLS, __Filename uintptr, __AccessMode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _AccessMode=%+v", __Filename, __AccessMode)
		defer func() { trc(`X_access->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_access.Addr(), __Filename, uintptr(__AccessMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_access_s = dll.NewProc("_access_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _access_s(const char *_Filename,int _AccessMode);
func X_access_s(tls *TLS, __Filename uintptr, __AccessMode int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _AccessMode=%+v", __Filename, __AccessMode)
		defer func() { trc(`X_access_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_access_s.Addr(), __Filename, uintptr(__AccessMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_chmod = dll.NewProc("_chmod")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _chmod(const char *_Filename,int _Mode);
func X_chmod(tls *TLS, __Filename uintptr, __Mode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _Mode=%+v", __Filename, __Mode)
		defer func() { trc(`X_chmod->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_chmod.Addr(), __Filename, uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_chsize = dll.NewProc("_chsize")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _chsize(int _FileHandle,long _Size);
func X_chsize(tls *TLS, __FileHandle int32, __Size int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Size=%+v", __FileHandle, __Size)
		defer func() { trc(`X_chsize->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_chsize.Addr(), uintptr(__FileHandle), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_chsize_s = dll.NewProc("_chsize_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _chsize_s (int _FileHandle, long long _Size);
func X_chsize_s(tls *TLS, __FileHandle int32, __Size int64) (r Terrno_t) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Size=%+v", __FileHandle, __Size)
		defer func() { trc(`X_chsize_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_chsize_s.Addr(), uintptr(__FileHandle), uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_close = dll.NewProc("_close")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _close(int _FileHandle);
func X_close(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_close->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_close.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_commit = dll.NewProc("_commit")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _commit(int _FileHandle);
func X_commit(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_commit->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_commit.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_creat = dll.NewProc("_creat")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _creat(const char *_Filename,int _PermissionMode);
func X_creat(tls *TLS, __Filename uintptr, __PermissionMode int32) (r int32) {
	if __ccgo_strace {
		trc("_Filename=%+v _PermissionMode=%+v", __Filename, __PermissionMode)
		defer func() { trc(`X_creat->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_creat.Addr(), __Filename, uintptr(__PermissionMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_dup = dll.NewProc("_dup")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _dup(int _FileHandle);
func X_dup(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_dup->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_dup.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_dup2 = dll.NewProc("_dup2")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _dup2(int _FileHandleSrc,int _FileHandleDst);
func X_dup2(tls *TLS, __FileHandleSrc int32, __FileHandleDst int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandleSrc=%+v _FileHandleDst=%+v", __FileHandleSrc, __FileHandleDst)
		defer func() { trc(`X_dup2->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_dup2.Addr(), uintptr(__FileHandleSrc), uintptr(__FileHandleDst))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_eof = dll.NewProc("_eof")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _eof(int _FileHandle);
func X_eof(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_eof->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_eof.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_filelength = dll.NewProc("_filelength")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _filelength(int _FileHandle);
func X_filelength(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_filelength->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_filelength.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_findfirst32 = dll.NewProc("_findfirst32")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _findfirst32(const char *_Filename,struct _finddata32_t *_FindData);
func X_findfirst32(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_findfirst32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findfirst32.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_findnext32 = dll.NewProc("_findnext32")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _findnext32(intptr_t _FindHandle,struct _finddata32_t *_FindData);
func X_findnext32(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_findnext32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findnext32.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_findclose = dll.NewProc("_findclose")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _findclose(intptr_t _FindHandle);
func X_findclose(tls *TLS, __FindHandle Tintptr_t) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v", __FindHandle)
		defer func() { trc(`X_findclose->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findclose.Addr(), uintptr(__FindHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_isatty = dll.NewProc("_isatty")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _isatty(int _FileHandle);
func X_isatty(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_isatty->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_isatty.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_locking = dll.NewProc("_locking")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _locking(int _FileHandle,int _LockMode,long _NumOfBytes);
func X_locking(tls *TLS, __FileHandle int32, __LockMode int32, __NumOfBytes int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _LockMode=%+v _NumOfBytes=%+v", __FileHandle, __LockMode, __NumOfBytes)
		defer func() { trc(`X_locking->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_locking.Addr(), uintptr(__FileHandle), uintptr(__LockMode), uintptr(__NumOfBytes))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_lseek = dll.NewProc("_lseek")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _lseek(int _FileHandle,long _Offset,int _Origin);
func X_lseek(tls *TLS, __FileHandle int32, __Offset int32, __Origin int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Offset=%+v _Origin=%+v", __FileHandle, __Offset, __Origin)
		defer func() { trc(`X_lseek->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_lseek.Addr(), uintptr(__FileHandle), uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_mktemp = dll.NewProc("_mktemp")

// __attribute__ ((__dllimport__)) char * __attribute__((__cdecl__)) _mktemp(char *_TemplateName);
func X_mktemp(tls *TLS, __TemplateName uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_TemplateName=%+v", __TemplateName)
		defer func() { trc(`X_mktemp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mktemp.Addr(), __TemplateName)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_mktemp_s = dll.NewProc("_mktemp_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _mktemp_s (char *_TemplateName,size_t _Size);
func X_mktemp_s(tls *TLS, __TemplateName uintptr, __Size Tsize_t) (r Terrno_t) {
	if __ccgo_strace {
		trc("_TemplateName=%+v _Size=%+v", __TemplateName, __Size)
		defer func() { trc(`X_mktemp_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mktemp_s.Addr(), __TemplateName, uintptr(__Size))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_pipe = dll.NewProc("_pipe")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _pipe(int *_PtHandles,unsigned int _PipeSize,int _TextMode);
func X_pipe(tls *TLS, __PtHandles uintptr, __PipeSize uint32, __TextMode int32) (r int32) {
	if __ccgo_strace {
		trc("_PtHandles=%+v _PipeSize=%+v _TextMode=%+v", __PtHandles, __PipeSize, __TextMode)
		defer func() { trc(`X_pipe->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_pipe.Addr(), __PtHandles, uintptr(__PipeSize), uintptr(__TextMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_read = dll.NewProc("_read")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _read(int _FileHandle,void *_DstBuf,unsigned int _MaxCharCount);
func X_read(tls *TLS, __FileHandle int32, __DstBuf uintptr, __MaxCharCount uint32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _DstBuf=%+v _MaxCharCount=%+v", __FileHandle, __DstBuf, __MaxCharCount)
		defer func() { trc(`X_read->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_read.Addr(), uintptr(__FileHandle), __DstBuf, uintptr(__MaxCharCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_setmode = dll.NewProc("_setmode")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _setmode(int _FileHandle,int _Mode);
func X_setmode(tls *TLS, __FileHandle int32, __Mode int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Mode=%+v", __FileHandle, __Mode)
		defer func() { trc(`X_setmode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_setmode.Addr(), uintptr(__FileHandle), uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_tell = dll.NewProc("_tell")

// __attribute__ ((__dllimport__)) long __attribute__((__cdecl__)) _tell(int _FileHandle);
func X_tell(tls *TLS, __FileHandle int32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_tell->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_tell.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_umask = dll.NewProc("_umask")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _umask(int _Mode);
func X_umask(tls *TLS, __Mode int32) (r int32) {
	if __ccgo_strace {
		trc("_Mode=%+v", __Mode)
		defer func() { trc(`X_umask->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_umask.Addr(), uintptr(__Mode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_umask_s = dll.NewProc("_umask_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _umask_s (int _NewMode,int *_OldMode);
func X_umask_s(tls *TLS, __NewMode int32, __OldMode uintptr) (r Terrno_t) {
	if __ccgo_strace {
		trc("_NewMode=%+v _OldMode=%+v", __NewMode, __OldMode)
		defer func() { trc(`X_umask_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_umask_s.Addr(), uintptr(__NewMode), __OldMode)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_write = dll.NewProc("_write")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _write(int _FileHandle,const void *_Buf,unsigned int _MaxCharCount);
func X_write(tls *TLS, __FileHandle int32, __Buf uintptr, __MaxCharCount uint32) (r int32) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Buf=%+v _MaxCharCount=%+v", __FileHandle, __Buf, __MaxCharCount)
		defer func() { trc(`X_write->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_write.Addr(), uintptr(__FileHandle), __Buf, uintptr(__MaxCharCount))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_filelengthi64 = dll.NewProc("_filelengthi64")

// __attribute__ ((__dllimport__)) long long __attribute__((__cdecl__)) _filelengthi64(int _FileHandle);
func X_filelengthi64(tls *TLS, __FileHandle int32) (r int64) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_filelengthi64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_filelengthi64.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_findfirst32i64 = dll.NewProc("_findfirst32i64")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _findfirst32i64(const char *_Filename,struct _finddata32i64_t *_FindData);
func X_findfirst32i64(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_findfirst32i64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findfirst32i64.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_findfirst64 = dll.NewProc("_findfirst64")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _findfirst64(const char *_Filename,struct __finddata64_t *_FindData);
func X_findfirst64(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_findfirst64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findfirst64.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_findfirst64i32 = dll.NewProc("_findfirst64i32")

// intptr_t __attribute__((__cdecl__)) _findfirst64i32(const char *_Filename,struct _finddata64i32_t *_FindData);
func X_findfirst64i32(tls *TLS, __Filename uintptr, __FindData uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _FindData=%+v", __Filename, __FindData)
		defer func() { trc(`X_findfirst64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findfirst64i32.Addr(), __Filename, __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_findnext32i64 = dll.NewProc("_findnext32i64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _findnext32i64(intptr_t _FindHandle,struct _finddata32i64_t *_FindData);
func X_findnext32i64(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_findnext32i64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findnext32i64.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_findnext64 = dll.NewProc("_findnext64")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _findnext64(intptr_t _FindHandle,struct __finddata64_t *_FindData);
func X_findnext64(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_findnext64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findnext64.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_findnext64i32 = dll.NewProc("_findnext64i32")

// int __attribute__((__cdecl__)) _findnext64i32(intptr_t _FindHandle,struct _finddata64i32_t *_FindData);
func X_findnext64i32(tls *TLS, __FindHandle Tintptr_t, __FindData uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FindHandle=%+v _FindData=%+v", __FindHandle, __FindData)
		defer func() { trc(`X_findnext64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_findnext64i32.Addr(), uintptr(__FindHandle), __FindData)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_lseeki64 = dll.NewProc("_lseeki64")

// long long __attribute__((__cdecl__)) _lseeki64(int _FileHandle, long long _Offset,int _Origin);
func X_lseeki64(tls *TLS, __FileHandle int32, __Offset int64, __Origin int32) (r int64) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Offset=%+v _Origin=%+v", __FileHandle, __Offset, __Origin)
		defer func() { trc(`X_lseeki64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_lseeki64.Addr(), uintptr(__FileHandle), uintptr(__Offset), uintptr(__Origin))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_telli64 = dll.NewProc("_telli64")

// long long __attribute__((__cdecl__)) _telli64(int _FileHandle);
func X_telli64(tls *TLS, __FileHandle int32) (r int64) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_telli64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_telli64.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int64(r0)
}

var proc_sopen_s = dll.NewProc("_sopen_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _sopen_s(int *_FileHandle,const char *_Filename,int _OpenFlag,int _ShareFlag,int _PermissionMode);
func X_sopen_s(tls *TLS, __FileHandle uintptr, __Filename uintptr, __OpenFlag int32, __ShareFlag int32, __PermissionMode int32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_FileHandle=%+v _Filename=%+v _OpenFlag=%+v _ShareFlag=%+v _PermissionMode=%+v", __FileHandle, __Filename, __OpenFlag, __ShareFlag, __PermissionMode)
		defer func() { trc(`X_sopen_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_sopen_s.Addr(), __FileHandle, __Filename, uintptr(__OpenFlag), uintptr(__ShareFlag), uintptr(__PermissionMode))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_get_osfhandle = dll.NewProc("_get_osfhandle")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _get_osfhandle(int _FileHandle);
func X_get_osfhandle(tls *TLS, __FileHandle int32) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_FileHandle=%+v", __FileHandle)
		defer func() { trc(`X_get_osfhandle->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_osfhandle.Addr(), uintptr(__FileHandle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_open_osfhandle = dll.NewProc("_open_osfhandle")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _open_osfhandle(intptr_t _OSFileHandle,int _Flags);
func X_open_osfhandle(tls *TLS, __OSFileHandle Tintptr_t, __Flags int32) (r int32) {
	if __ccgo_strace {
		trc("_OSFileHandle=%+v _Flags=%+v", __OSFileHandle, __Flags)
		defer func() { trc(`X_open_osfhandle->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_open_osfhandle.Addr(), uintptr(__OSFileHandle), uintptr(__Flags))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

type T_PVFV = uintptr

type T_PIFV = uintptr

type T_PVFI = uintptr

type T_onexit_table_t = struct {
	F_first uintptr
	F_last  uintptr
	F_end   uintptr
}

type T_pid_t = int64

type Tpid_t = int64

type T_mode_t = uint16

type Tmode_t = uint16

type Tuseconds_t = uint32

type T_sigset_t = uint64

type T_beginthread_proc_type = uintptr

type T_beginthreadex_proc_type = uintptr

var proc_endthread = dll.NewProc("_endthread")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _endthread(void) __attribute__ ((__noreturn__));
func X_endthread(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(proc_endthread.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

type T_tls_callback_type = uintptr

var proc_register_thread_local_exe_atexit_callback = dll.NewProc("_register_thread_local_exe_atexit_callback")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _register_thread_local_exe_atexit_callback(_tls_callback_type callback);
func X_register_thread_local_exe_atexit_callback(tls *TLS, _callback T_tls_callback_type) {
	X__ccgo_SyscallFP()
	panic(663)
}

var proc_cexit = dll.NewProc("_cexit")

// void __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _cexit(void);
func X_cexit(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(proc_cexit.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_c_exit = dll.NewProc("_c_exit")

// void __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _c_exit(void);
func X_c_exit(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(proc_c_exit.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc_getpid = dll.NewProc("_getpid")

// __attribute__ ((__dllimport__)) int __attribute__((__cdecl__)) _getpid(void);
func X_getpid(tls *TLS) (r int32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_getpid->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getpid.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_cwait = dll.NewProc("_cwait")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _cwait(int *_TermStat,intptr_t _ProcHandle,int _Action);
func X_cwait(tls *TLS, __TermStat uintptr, __ProcHandle Tintptr_t, __Action int32) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_TermStat=%+v _ProcHandle=%+v _Action=%+v", __TermStat, __ProcHandle, __Action)
		defer func() { trc(`X_cwait->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_cwait.Addr(), __TermStat, uintptr(__ProcHandle), uintptr(__Action))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_execv = dll.NewProc("_execv")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _execv(const char *_Filename,const char *const *_ArgList);
func X_execv(tls *TLS, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v", __Filename, __ArgList)
		defer func() { trc(`X_execv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_execv.Addr(), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_execve = dll.NewProc("_execve")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _execve(const char *_Filename,const char *const *_ArgList,const char *const *_Env);
func X_execve(tls *TLS, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v _Env=%+v", __Filename, __ArgList, __Env)
		defer func() { trc(`X_execve->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_execve.Addr(), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_execvp = dll.NewProc("_execvp")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _execvp(const char *_Filename,const char *const *_ArgList);
func X_execvp(tls *TLS, __Filename uintptr, __ArgList uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v", __Filename, __ArgList)
		defer func() { trc(`X_execvp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_execvp.Addr(), __Filename, __ArgList)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_execvpe = dll.NewProc("_execvpe")

// __attribute__ ((__dllimport__)) intptr_t __attribute__((__cdecl__)) _execvpe(const char *_Filename,const char *const *_ArgList,const char *const *_Env);
func X_execvpe(tls *TLS, __Filename uintptr, __ArgList uintptr, __Env uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v _ArgList=%+v _Env=%+v", __Filename, __ArgList, __Env)
		defer func() { trc(`X_execvpe->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_execvpe.Addr(), __Filename, __ArgList, __Env)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_loaddll = dll.NewProc("_loaddll")

// intptr_t __attribute__((__cdecl__)) _loaddll(char *_Filename);
func X_loaddll(tls *TLS, __Filename uintptr) (r Tintptr_t) {
	if __ccgo_strace {
		trc("_Filename=%+v", __Filename)
		defer func() { trc(`X_loaddll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_loaddll.Addr(), __Filename)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Tintptr_t(r0)
}

var proc_unloaddll = dll.NewProc("_unloaddll")

// int __attribute__((__cdecl__)) _unloaddll(intptr_t _Handle);
func X_unloaddll(tls *TLS, __Handle Tintptr_t) (r int32) {
	if __ccgo_strace {
		trc("_Handle=%+v", __Handle)
		defer func() { trc(`X_unloaddll->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_unloaddll.Addr(), uintptr(__Handle))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_getdllprocaddr = dll.NewProc("_getdllprocaddr")

// int ( * __attribute__((__cdecl__)) _getdllprocaddr(intptr_t _Handle,char *_ProcedureName,intptr_t _Ordinal))(void);
func X_getdllprocaddr(tls *TLS, __Handle Tintptr_t, __ProcedureName uintptr, __Ordinal Tintptr_t) (r uintptr) {
	if __ccgo_strace {
		trc("_Handle=%+v _ProcedureName=%+v _Ordinal=%+v", __Handle, __ProcedureName, __Ordinal)
		defer func() { trc(`X_getdllprocaddr->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_getdllprocaddr.Addr(), uintptr(__Handle), __ProcedureName, uintptr(__Ordinal))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_controlfp = dll.NewProc("_controlfp")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _controlfp (unsigned int _NewValue, unsigned int _Mask);
func X_controlfp(tls *TLS, __NewValue uint32, __Mask uint32) (r uint32) {
	if __ccgo_strace {
		trc("_NewValue=%+v _Mask=%+v", __NewValue, __Mask)
		defer func() { trc(`X_controlfp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_controlfp.Addr(), uintptr(__NewValue), uintptr(__Mask))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_controlfp_s = dll.NewProc("_controlfp_s")

// __attribute__((dllimport)) errno_t __attribute__((__cdecl__)) _controlfp_s(unsigned int *_CurrentState, unsigned int _NewValue, unsigned int _Mask);
func X_controlfp_s(tls *TLS, __CurrentState uintptr, __NewValue uint32, __Mask uint32) (r Terrno_t) {
	if __ccgo_strace {
		trc("_CurrentState=%+v _NewValue=%+v _Mask=%+v", __CurrentState, __NewValue, __Mask)
		defer func() { trc(`X_controlfp_s->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_controlfp_s.Addr(), __CurrentState, uintptr(__NewValue), uintptr(__Mask))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return Terrno_t(r0)
}

var proc_control87 = dll.NewProc("_control87")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _control87 (unsigned int _NewValue, unsigned int _Mask);
func X_control87(tls *TLS, __NewValue uint32, __Mask uint32) (r uint32) {
	if __ccgo_strace {
		trc("_NewValue=%+v _Mask=%+v", __NewValue, __Mask)
		defer func() { trc(`X_control87->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_control87.Addr(), uintptr(__NewValue), uintptr(__Mask))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_clearfp = dll.NewProc("_clearfp")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _clearfp (void);
func X_clearfp(tls *TLS) (r uint32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_clearfp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_clearfp.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_statusfp = dll.NewProc("_statusfp")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _statusfp (void);
func X_statusfp(tls *TLS) (r uint32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_statusfp->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_statusfp.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_fpreset = dll.NewProc("_fpreset")

// void __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) _fpreset (void);
func X_fpreset(tls *TLS) {
	if __ccgo_strace {
		trc("")
	}
	r0, r1, err := syscall.SyscallN(proc_fpreset.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc__fpecode = dll.NewProc("__fpecode")

// __attribute__ ((__dllimport__)) int * __attribute__((__cdecl__)) __attribute__ ((__nothrow__)) __fpecode(void);
func X__fpecode(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X__fpecode->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc__fpecode.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

type Tlconv = struct {
	Fdecimal_point        uintptr
	Fthousands_sep        uintptr
	Fgrouping             uintptr
	Fint_curr_symbol      uintptr
	Fcurrency_symbol      uintptr
	Fmon_decimal_point    uintptr
	Fmon_thousands_sep    uintptr
	Fmon_grouping         uintptr
	Fpositive_sign        uintptr
	Fnegative_sign        uintptr
	Fint_frac_digits      int8
	Ffrac_digits          int8
	Fp_cs_precedes        int8
	Fp_sep_by_space       int8
	Fn_cs_precedes        int8
	Fn_sep_by_space       int8
	Fp_sign_posn          int8
	Fn_sign_posn          int8
	F_W_decimal_point     uintptr
	F_W_thousands_sep     uintptr
	F_W_int_curr_symbol   uintptr
	F_W_currency_symbol   uintptr
	F_W_mon_decimal_point uintptr
	F_W_mon_thousands_sep uintptr
	F_W_positive_sign     uintptr
	F_W_negative_sign     uintptr
}

var proc_configthreadlocale = dll.NewProc("_configthreadlocale")

// int __attribute__((__cdecl__)) _configthreadlocale(int _Flag);
func X_configthreadlocale(tls *TLS, __Flag int32) (r int32) {
	if __ccgo_strace {
		trc("_Flag=%+v", __Flag)
		defer func() { trc(`X_configthreadlocale->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_configthreadlocale.Addr(), uintptr(__Flag))
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var procsetlocale = dll.NewProc("setlocale")

// char * __attribute__((__cdecl__)) setlocale(int _Category,const char *_Locale);
func Xsetlocale(tls *TLS, __Category int32, __Locale uintptr) (r uintptr) {
	if __ccgo_strace {
		trc("_Category=%+v _Locale=%+v", __Category, __Locale)
		defer func() { trc(`Xsetlocale->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(procsetlocale.Addr(), uintptr(__Category), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proclocaleconv = dll.NewProc("localeconv")

// __attribute__ ((__dllimport__)) struct lconv * __attribute__((__cdecl__)) localeconv(void);
func Xlocaleconv(tls *TLS) (r uintptr) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`Xlocaleconv->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proclocaleconv.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uintptr(r0)
}

var proc_get_current_locale = dll.NewProc("_get_current_locale")

// __attribute__ ((__dllimport__)) _locale_t __attribute__((__cdecl__)) _get_current_locale(void);
func X_get_current_locale(tls *TLS) (r T_locale_t) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X_get_current_locale->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_get_current_locale.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T_locale_t(r0)
}

var proc_create_locale = dll.NewProc("_create_locale")

// __attribute__ ((__dllimport__)) _locale_t __attribute__((__cdecl__)) _create_locale(int _Category,const char *_Locale);
func X_create_locale(tls *TLS, __Category int32, __Locale uintptr) (r T_locale_t) {
	if __ccgo_strace {
		trc("_Category=%+v _Locale=%+v", __Category, __Locale)
		defer func() { trc(`X_create_locale->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_create_locale.Addr(), uintptr(__Category), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return T_locale_t(r0)
}

var proc_free_locale = dll.NewProc("_free_locale")

// __attribute__ ((__dllimport__)) void __attribute__((__cdecl__)) _free_locale(_locale_t _Locale);
func X_free_locale(tls *TLS, __Locale T_locale_t) {
	if __ccgo_strace {
		trc("_Locale=%+v", __Locale)
	}
	r0, r1, err := syscall.SyscallN(proc_free_locale.Addr(), __Locale)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
}

var proc___lc_codepage_func = dll.NewProc("___lc_codepage_func")

// __attribute__ ((__dllimport__)) unsigned int __attribute__((__cdecl__)) ___lc_codepage_func(void);
func X___lc_codepage_func(tls *TLS) (r uint32) {
	if __ccgo_strace {
		trc("")
		defer func() { trc(`X___lc_codepage_func->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc___lc_codepage_func.Addr())
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return uint32(r0)
}

var proc_mkdir = dll.NewProc("_mkdir")

// int __attribute__((__cdecl__)) _mkdir(const char *_Path);
func X_mkdir(tls *TLS, __Path uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Path=%+v", __Path)
		defer func() { trc(`X_mkdir->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_mkdir.Addr(), __Path)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_fstat64 = dll.NewProc("_fstat64")

// int __attribute__((__cdecl__)) _fstat64(int _FileDes,struct _stat64 *_Stat);
func X_fstat64(tls *TLS, __FileDes int32, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_FileDes=%+v _Stat=%+v", __FileDes, __Stat)
		defer func() { trc(`X_fstat64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_fstat64.Addr(), uintptr(__FileDes), __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_stat64 = dll.NewProc("_stat64")

// int __attribute__((__cdecl__)) _stat64(const char *_Name,struct _stat64 *_Stat);
func X_stat64(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_stat64->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stat64.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

var proc_stat64i32 = dll.NewProc("_stat64i32")

// int __attribute__((__cdecl__)) _stat64i32(const char *_Name,struct _stat64i32 *_Stat);
func X_stat64i32(tls *TLS, __Name uintptr, __Stat uintptr) (r int32) {
	if __ccgo_strace {
		trc("_Name=%+v _Stat=%+v", __Name, __Stat)
		defer func() { trc(`X_stat64i32->%+v`, r) }()
	}
	r0, r1, err := syscall.SyscallN(proc_stat64i32.Addr(), __Name, __Stat)
	if err != 0 {
		if __ccgo_strace {
			trc(`r0=%v r1=%v err=%v`, r0, r1, err)
		}
		tls.setErrno(int32(err))
	}
	return int32(r0)
}

func init() {
	for _, nm := range []string{
		"_Exit",
		"___lc_codepage_func",
		"___mb_cur_max_func",
		"__daylight",
		"__doserrno",
		"__dstbias",
		"__fpecode",
		"__isascii",
		"__iscsym",
		"__iscsymf",
		"__iswcsym",
		"__iswcsymf",
		"__p___argc",
		"__p___argv",
		"__p___wargv",
		"__p__environ",
		"__p__fmode",
		"__p__pgmptr",
		"__p__wenviron",
		"__p__wpgmptr",
		"__pctype_func",
		"__setusermatherr",
		"__sys_errlist",
		"__sys_nerr",
		"__timezone",
		"__toascii",
		"__tzname",
		"__wcserror",
		"_abs64",
		"_access",
		"_access_s",
		"_aligned_free",
		"_aligned_malloc",
		"_aligned_offset_malloc",
		"_aligned_offset_realloc",
		"_aligned_offset_recalloc",
		"_aligned_realloc",
		"_aligned_recalloc",
		"_atodbl",
		"_atodbl_l",
		"_atoflt",
		"_atoflt_l",
		"_atoi64",
		"_atoi64_l",
		"_atoi_l",
		"_atol_l",
		"_atoldbl",
		"_atoldbl_l",
		"_beep",
		"_byteswap_uint64",
		"_byteswap_ulong",
		"_byteswap_ushort",
		"_c_exit",
		"_cexit",
		"_cgetws",
		"_chmod",
		"_chsize",
		"_chsize_s",
		"_clearfp",
		"_close",
		"_commit",
		"_configthreadlocale",
		"_control87",
		"_controlfp",
		"_controlfp_s",
		"_cputws",
		"_creat",
		"_create_locale",
		"_ctime32",
		"_ctime32_s",
		"_ctime64",
		"_ctime64_s",
		"_cwait",
		"_dup",
		"_dup2",
		"_endthread",
		"_eof",
		"_execv",
		"_execve",
		"_execvp",
		"_execvpe",
		"_exit",
		"_expand",
		"_fclose_nolock",
		"_fcloseall",
		"_fdopen",
		"_fflush_nolock",
		"_fgetc_nolock",
		"_fgetchar",
		"_fgetwc_nolock",
		"_fgetwchar",
		"_filelength",
		"_filelengthi64",
		"_fileno",
		"_findclose",
		"_findfirst32",
		"_findfirst32i64",
		"_findfirst64",
		"_findfirst64i32",
		"_findnext32",
		"_findnext32i64",
		"_findnext64",
		"_findnext64i32",
		"_flushall",
		"_fpreset",
		"_fputc_nolock",
		"_fputchar",
		"_fputwc_nolock",
		"_fputwchar",
		"_fread_nolock",
		"_free_locale",
		"_fseek_nolock",
		"_fseeki64",
		"_fseeki64_nolock",
		"_fsopen",
		"_fstat64",
		"_ftell_nolock",
		"_ftelli64",
		"_ftelli64_nolock",
		"_ftime32",
		"_ftime64",
		"_fullpath",
		"_fwrite_nolock",
		"_get_current_locale",
		"_get_daylight",
		"_get_doserrno",
		"_get_dstbias",
		"_get_errno",
		"_get_fmode",
		"_get_heap_handle",
		"_get_invalid_parameter_handler",
		"_get_osfhandle",
		"_get_pgmptr",
		"_get_printf_count_output",
		"_get_purecall_handler",
		"_get_timezone",
		"_get_tzname",
		"_get_wpgmptr",
		"_getc_nolock",
		"_getcwd",
		"_getdllprocaddr",
		"_getmaxstdio",
		"_getpid",
		"_getsystime",
		"_getw",
		"_getwch",
		"_getwch_nolock",
		"_getwche",
		"_getwche_nolock",
		"_getws",
		"_gmtime32",
		"_gmtime32_s",
		"_gmtime64",
		"_gmtime64_s",
		"_heapchk",
		"_heapmin",
		"_heapwalk",
		"_i64toa",
		"_i64tow",
		"_isalnum_l",
		"_isalpha_l",
		"_isatty",
		"_iscntrl_l",
		"_isctype",
		"_isctype_l",
		"_isdigit_l",
		"_isgraph_l",
		"_isleadbyte_l",
		"_islower_l",
		"_isprint_l",
		"_ispunct_l",
		"_isspace_l",
		"_isupper_l",
		"_iswalnum_l",
		"_iswalpha_l",
		"_iswcntrl_l",
		"_iswcsym_l",
		"_iswcsymf_l",
		"_iswctype_l",
		"_iswdigit_l",
		"_iswgraph_l",
		"_iswlower_l",
		"_iswprint_l",
		"_iswpunct_l",
		"_iswspace_l",
		"_iswupper_l",
		"_iswxdigit_l",
		"_isxdigit_l",
		"_itoa",
		"_itow",
		"_loaddll",
		"_localtime32",
		"_localtime32_s",
		"_localtime64",
		"_localtime64_s",
		"_lock_file",
		"_locking",
		"_lrotl",
		"_lrotr",
		"_lseek",
		"_lseeki64",
		"_ltoa",
		"_ltow",
		"_makepath",
		"_mblen_l",
		"_mbstowcs_l",
		"_mbstrlen",
		"_mbstrlen_l",
		"_mbstrnlen",
		"_mbstrnlen_l",
		"_mbtowc_l",
		"_memccpy",
		"_memicmp",
		"_memicmp_l",
		"_mkdir",
		"_mkgmtime32",
		"_mkgmtime64",
		"_mktemp",
		"_mktemp_s",
		"_mktime32",
		"_mktime64",
		"_msize",
		"_open_osfhandle",
		"_pclose",
		"_pipe",
		"_popen",
		"_putc_nolock",
		"_putenv",
		"_putw",
		"_putwch",
		"_putwch_nolock",
		"_putws",
		"_read",
		"_recalloc",
		"_register_thread_local_exe_atexit_callback",
		"_resetstkoflw",
		"_rmtmp",
		"_rotl",
		"_rotl64",
		"_rotr",
		"_rotr64",
		"_searchenv",
		"_set_abort_behavior",
		"_set_doserrno",
		"_set_errno",
		"_set_error_mode",
		"_set_fmode",
		"_set_invalid_parameter_handler",
		"_set_printf_count_output",
		"_set_purecall_handler",
		"_seterrormode",
		"_setmaxstdio",
		"_setmode",
		"_setsystime",
		"_sleep",
		"_sopen_s",
		"_spawnv",
		"_spawnve",
		"_spawnvp",
		"_spawnvpe",
		"_splitpath",
		"_stat64",
		"_stat64i32",
		"_statusfp",
		"_strcoll_l",
		"_strdate",
		"_strdate_s",
		"_strdup",
		"_strerror",
		"_strftime_l",
		"_stricmp",
		"_stricmp_l",
		"_stricoll",
		"_stricoll_l",
		"_strlwr",
		"_strncoll",
		"_strncoll_l",
		"_strnicmp",
		"_strnicmp_l",
		"_strnicoll",
		"_strnicoll_l",
		"_strnset",
		"_strrev",
		"_strset",
		"_strtime",
		"_strtime_s",
		"_strtoi64",
		"_strtoi64_l",
		"_strtol_l",
		"_strtoui64",
		"_strtoui64_l",
		"_strtoul_l",
		"_strupr",
		"_strupr_l",
		"_strxfrm_l",
		"_swab",
		"_tell",
		"_telli64",
		"_tempnam",
		"_time32",
		"_time64",
		"_timespec32_get",
		"_timespec64_get",
		"_tolower",
		"_tolower_l",
		"_toupper",
		"_toupper_l",
		"_towlower_l",
		"_towupper_l",
		"_tzset",
		"_ui64toa",
		"_ui64tow",
		"_ultoa",
		"_ultow",
		"_umask",
		"_umask_s",
		"_ungetc_nolock",
		"_ungetwc_nolock",
		"_ungetwch",
		"_ungetwch_nolock",
		"_unlink",
		"_unloaddll",
		"_unlock_file",
		"_waccess",
		"_wasctime",
		"_wasctime_s",
		"_wchdir",
		"_wchmod",
		"_wcreat",
		"_wcscoll_l",
		"_wcsdup",
		"_wcserror",
		"_wcsftime_l",
		"_wcsicmp",
		"_wcsicmp_l",
		"_wcsicoll",
		"_wcsicoll_l",
		"_wcslwr",
		"_wcslwr_l",
		"_wcsncoll",
		"_wcsncoll_l",
		"_wcsnicmp",
		"_wcsnicmp_l",
		"_wcsnicoll",
		"_wcsnicoll_l",
		"_wcsnset",
		"_wcsrev",
		"_wcsset",
		"_wcstoi64",
		"_wcstoi64_l",
		"_wcstol_l",
		"_wcstombs_l",
		"_wcstoui64",
		"_wcstoui64_l",
		"_wcstoul_l",
		"_wcsupr",
		"_wcsupr_l",
		"_wcsxfrm_l",
		"_wctime32",
		"_wctime32_s",
		"_wctime64",
		"_wctime64_s",
		"_wctomb_l",
		"_wexecv",
		"_wexecve",
		"_wexecvp",
		"_wexecvpe",
		"_wfdopen",
		"_wfindfirst32",
		"_wfindfirst32i64",
		"_wfindfirst64",
		"_wfindfirst64i32",
		"_wfindnext32",
		"_wfindnext32i64",
		"_wfindnext64",
		"_wfindnext64i32",
		"_wfopen",
		"_wfreopen",
		"_wfsopen",
		"_wfullpath",
		"_wgetcwd",
		"_wgetdcwd",
		"_wmakepath",
		"_wmkdir",
		"_wmktemp",
		"_wperror",
		"_wpopen",
		"_wremove",
		"_wrename",
		"_write",
		"_wrmdir",
		"_wsearchenv",
		"_wsetlocale",
		"_wsopen_s",
		"_wspawnv",
		"_wspawnve",
		"_wspawnvp",
		"_wspawnvpe",
		"_wsplitpath",
		"_wstat32",
		"_wstat32i64",
		"_wstat64",
		"_wstat64i32",
		"_wstrdate",
		"_wstrdate_s",
		"_wstrtime",
		"_wstrtime_s",
		"_wsystem",
		"_wtempnam",
		"_wtmpnam",
		"_wtoi",
		"_wtoi64",
		"_wtoi64_l",
		"_wtoi_l",
		"_wtol",
		"_wtol_l",
		"_wunlink",
		"abort",
		"abs",
		"asctime",
		"asctime_s",
		"atoi",
		"atol",
		"atoll",
		"bsearch",
		"btowc",
		"calloc",
		"clearerr",
		"clock",
		"div",
		"fclose",
		"feof",
		"ferror",
		"fflush",
		"fgetc",
		"fgetpos",
		"fgets",
		"fgetwc",
		"fgetws",
		"fopen",
		"fputc",
		"fputs",
		"fputwc",
		"fputws",
		"fread",
		"free",
		"freopen",
		"fseek",
		"fsetpos",
		"ftell",
		"fwrite",
		"getc",
		"getchar",
		"getenv",
		"gets",
		"getwc",
		"getwchar",
		"is_wctype",
		"isalnum",
		"isalpha",
		"isblank",
		"iscntrl",
		"isdigit",
		"isgraph",
		"isleadbyte",
		"islower",
		"isprint",
		"ispunct",
		"isspace",
		"isupper",
		"iswalnum",
		"iswalpha",
		"iswascii",
		"iswblank",
		"iswcntrl",
		"iswctype",
		"iswdigit",
		"iswgraph",
		"iswlower",
		"iswprint",
		"iswpunct",
		"iswspace",
		"iswupper",
		"iswxdigit",
		"isxdigit",
		"labs",
		"ldiv",
		"llabs",
		"localeconv",
		"malloc",
		"mblen",
		"mbrlen",
		"mbrtowc",
		"mbsrtowcs",
		"mbstowcs",
		"mbtowc",
		"memchr",
		"memcmp",
		"memcpy",
		"memcpy_s",
		"memmove",
		"memset",
		"perror",
		"putc",
		"putchar",
		"puts",
		"putwc",
		"putwchar",
		"quick_exit",
		"rand",
		"realloc",
		"remove",
		"rename",
		"rewind",
		"setbuf",
		"setlocale",
		"setvbuf",
		"srand",
		"strcat",
		"strchr",
		"strcmp",
		"strcoll",
		"strcpy",
		"strcspn",
		"strerror",
		"strftime",
		"strlen",
		"strncat",
		"strncmp",
		"strncpy",
		"strnlen",
		"strpbrk",
		"strrchr",
		"strspn",
		"strstr",
		"strtok",
		"strtol",
		"strtoll",
		"strtoul",
		"strtoull",
		"strxfrm",
		"system",
		"tmpfile",
		"tmpnam",
		"tolower",
		"toupper",
		"towlower",
		"towupper",
		"ungetc",
		"ungetwc",
		"wcrtomb",
		"wcscat",
		"wcschr",
		"wcscmp",
		"wcscoll",
		"wcscpy",
		"wcscspn",
		"wcsftime",
		"wcslen",
		"wcsncat",
		"wcsncmp",
		"wcsncpy",
		"wcsnlen",
		"wcspbrk",
		"wcsrchr",
		"wcsrtombs",
		"wcsspn",
		"wcsstr",
		"wcstok",
		"wcstol",
		"wcstoll",
		"wcstombs",
		"wcstoul",
		"wcstoull",
		"wcsxfrm",
		"wctob",
		"wctomb",
	} {
		func() {
			defer func() {
				if recover() != nil {
					println(nm)
				}
			}()
			dll.NewProc(nm).Addr()
		}()
	}
}
