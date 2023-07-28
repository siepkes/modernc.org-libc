// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"modernc.org/memory"
)

var (
	initLibcOnce sync.Once

	allocator   memory.Allocator
	allocatorMu sync.Mutex

	tid  atomic.Int32 // TLS Go ID
	tls1 *TLS
)

// Start executes a transpilled main program.
func Start(main func(*TLS, int32, uintptr) int32) {
	runtime.LockOSThread()
	tls := NewTLS()
	argv := Xcalloc(tls, 1, uint64((len(os.Args)+1)*int(unsafe.Sizeof(uintptr(0)))))
	if argv == 0 {
		panic("OOM")
	}

	p := argv
	for _, v := range os.Args {
		s := Xcalloc(tls, 1, uint64(len(v)+1))
		if s == 0 {
			panic("OOM")
		}

		copy(unsafe.Slice((*byte)(unsafe.Pointer(s)), len(v)), v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += unsafe.Sizeof(uintptr(0))
	}
	//TODO SetEnviron(tls, os.Environ())
	Xexit(tls, main(tls, int32(len(os.Args)), argv))
}

// musl-0.6.0/crt/x86_64/crt1.s
//
//	/* Written 2011 Nicholas J. Kain, released as Public Domain */
//	.text
//	.global _start
//	_start:
//		xor %rbp,%rbp   /* rbp:undefined -> mark as zero 0 (ABI) */
//		mov %rdx,%r9    /* 6th arg: ptr to register with atexit() */
//		pop %rsi        /* 2nd arg: argc */
//		mov %rsp,%rdx   /* 3rd arg: argv */
//		andq $-16,%rsp  /* align stack pointer */
//		push %rax       /* 8th arg: glibc ABI compatible */
//		push %rsp       /* 7th arg: glibc ABI compatible */
//		xor %r8,%r8     /* 5th arg: always 0 */
//		xor %rcx,%rcx   /* 4th arg: always 0 */
//		mov $main,%rdi  /* 1st arg: application entry ip */
//		call __libc_start_main /* musl init will run the program */
//	.L0:	jmp .L0
func initLibc() {
	var argv []uintptr
	for _, v := range os.Args {
		p := mustPrivateCalloc(len(v) + 1)
		copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), len(v)), v)
		argv = append(argv, p)
	}
	argv = append(argv, 0, 0, 0, 0) //TODO envp, auxv, ?
	argvP := mustPrivateCalloc(len(argv) * int(unsafe.Sizeof(uintptr(0))))
	copy(unsafe.Slice((*uintptr)(unsafe.Pointer(argvP)), len(argv)), argv)

	tls1 = newTLS()

	// mov $main,%rdi  /* 1st arg: application entry ip */
	// pop %rsi        /* 2nd arg: argc */
	// mov %rsp,%rdx   /* 3rd arg: argv */
	// xor %rcx,%rcx   /* 4th arg: always 0 */
	// xor %r8,%r8     /* 5th arg: always 0 */
	// mov %rdx,%r9    /* 6th arg: ptr to register with atexit() */

	x___libc_start_main(tls1, 0, int32(len(os.Args)), argvP, 0, 0, 0)
}

// CString returns a pointer to a zero-terminated version of s. The caller is
// responsible for freeing the allocated memory using Xfree.
func CString(s string) (uintptr, error) {
	n := len(s)
	p := Xmalloc(nil, uint64(n)+1)
	if p == 0 {
		return 0, fmt.Errorf("CString: cannot allocate %d bytes", n+1)
	}

	copy(unsafe.Slice((*byte)(unsafe.Pointer(p)), n), s)
	*(*byte)(unsafe.Pointer(p + uintptr(n))) = 0
	return p, nil
}

// GoBytes returns a byte slice from a C char* having length len bytes.
func GoBytes(s uintptr, len int) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(s)), len)
}

// GoString returns the value of a C string at s.
func GoString(s uintptr) string {
	if s == 0 {
		return ""
	}

	var buf []byte
	for {
		b := *(*byte)(unsafe.Pointer(s))
		if b == 0 {
			return string(buf)
		}

		buf = append(buf, b)
		s++
	}
}

// Go libc private heap, outside of the C heap.
func mustPrivateCalloc(sz int) uintptr {
	if p, _ := privateCalloc(sz); p != 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func mustPrivateMalloc(sz int) uintptr {
	if p, _ := privateMalloc(sz); p != 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func mustPrivateRealloc(p uintptr, sz int) uintptr {
	if p, _ := privateRealloc(p, sz); p != 0 || sz == 0 {
		return p
	}

	panic(todo("OOM"))
}

// Go libc private heap, outside of the C heap.
func privateCalloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrCalloc(sz)
}

// Go libc private heap, outside of the C heap.
func privateMalloc(sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrMalloc(sz)
}

// Go libc private heap, outside of the C heap.
func privateRealloc(p uintptr, sz int) (uintptr, error) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	return allocator.UintptrRealloc(p, sz)
}

// Go libc private heap, outside of the C heap.
func privateFree(p uintptr) {
	allocatorMu.Lock()

	defer allocatorMu.Unlock()

	allocator.UintptrFree(p)
}

type tlsStack struct {
	p  uintptr
	sz int
}

// TLS emulates thread local storage.
type TLS struct {
	fs    uintptr
	sp    int
	stack []tlsStack

	ID int32
}

// NewTLS returns a newly created TLS that must be eventually closed to prevent
// resource leaks.
func NewTLS() *TLS {
	initLibcOnce.Do(initLibc)
	return tls1 //TODO
}

func newTLS() *TLS {
	t := &TLS{
		ID: tid.Add(1),
	}
	return t
}

func (t *TLS) Alloc(n int) (r uintptr) {
	return mustPrivateMalloc(n) //TODO-

	sp := t.sp
	if sp < len(t.stack) {
		s := t.stack[sp]
		if s.sz >= n {
			t.sp++
			return s.p
		}

		r = mustPrivateRealloc(s.p, n)
		t.stack[sp].p = r
		t.stack[sp].sz = n
		t.sp++
		return r

	}

	r = mustPrivateMalloc(n)
	t.stack = append(t.stack, tlsStack{r, n})
	t.sp++
	return r
}

func (t *TLS) Free(n int) {
	t.sp--
}

func (t *TLS) Close() {
	if t.ID == 1 { //TODO-
		t.sp = 0
		return
	}

	for _, v := range t.stack[:t.sp] {
		privateFree(v.p)
	}
	*t = TLS{}
}

func _sqrt(tls *TLS, x float64) float64 {
	return math.Sqrt(x)
}

func _sqrtf(tls *TLS, x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

func ___restore_rt(tls *TLS) {
	panic(todo(""))
}

func ___restore(tls *TLS) {
	panic(todo(""))
}

func _longjmp(tls *TLS, env uintptr, val int32) int64 {
	panic(todo(""))
}

func a_cas(p uintptr, t, s int32) int32

func _a_cas(tls *TLS, p uintptr, test, s int32) int32 {
	return a_cas(p, test, s)
}

func _a_spin(tls *TLS) {
	panic(todo(""))
}

func _a_fetch_add(tls *TLS, p uintptr, v int32) int32 {
	panic(todo(""))
}

func a_and_64(p uintptr, x uint64)

func _a_and_64(tls *TLS, p uintptr, x uint64) {
	a_and_64(p, x)
}

func a_cas_p(p uintptr, t, s uintptr) uintptr

func _a_cas_p(tls *TLS, p, test, s uintptr) uintptr {
	return a_cas_p(p, test, s)
}

func a_or(p uintptr, v int32)

func _a_or(tls *TLS, p uintptr, v int32) {
	a_or(p, v)
}

func a_or_64(p uintptr, v uint64)

func _a_or_64(tls *TLS, p uintptr, v uint64) {
	a_or_64(p, v)
}

func _a_ctz_64(tls *TLS, x uint64) int32 {
	return int32(bits.TrailingZeros64(x))
}

func _a_inc(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), 1)
}

func _a_dec(tls *TLS, p uintptr) {
	atomic.AddInt32((*int32)(unsafe.Pointer(p)), -1)
}

func _a_swap(tls *TLS, p uintptr, v int32) int32 {
	return atomic.SwapInt32((*int32)(unsafe.Pointer(p)), v)
}

func _a_store(tls *TLS, p uintptr, v int32) {
	atomic.StoreInt32((*int32)(unsafe.Pointer(p)), v)
}

func ___pthread_self(tls *TLS) uintptr {
	return tls.fs
}

func _syscall0(tls *TLS, n int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), 0, 0, 0)
	return int64(r1)
}

func _syscall1(tls *TLS, n, a1 int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), uintptr(a1), 0, 0)
	return int64(r1)
}

func _syscall2(tls *TLS, n, a1, a2 int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), 0)
	return int64(r1)
}

func _syscall3(tls *TLS, n, a1, a2, a3 int64) int64 {
	r1, _, _ := syscall.Syscall(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3))
	return int64(r1)
}

func _syscall4(tls *TLS, n, a1, a2, a3, a4 int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), 0, 0)
	return int64(r1)
}

func _syscall5(tls *TLS, n, a1, a2, a3, a4, a5 int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), 0)
	return int64(r1)
}

func _syscall6(tls *TLS, n, a1, a2, a3, a4, a5, a6 int64) int64 {
	r1, _, _ := syscall.Syscall6(uintptr(n), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5), uintptr(a6))
	return int64(r1)
}

// This is not the set_thread_area syscall, but arch_prctl syscall with
// subcommand ARCH_SET_FS.
func ___set_thread_area(tls *TLS, p uintptr) int32 {
	tls.fs = p
	return 0
}

func ___bswap_16(tls *TLS, __x uint16) uint16 {
	return __x<<8 | __x>>8
}

func ___bswap32(tls *TLS, __x uint32) uint32 {
	return __x>>int32(24) | __x>>int32(8)&uint32(0xff00) | __x<<int32(8)&uint32(0xff0000) | __x<<int32(24)
}

func ___bswap_32(tls *TLS, x uint32) uint32 {
	return ___bswap32(tls, x)
}

func _memcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return _memmove(tls, dest, src, n)
}

func _memmove(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	if n != 0 {
		copy(unsafe.Slice((*byte)(unsafe.Pointer(dest)), n), unsafe.Slice((*byte)(unsafe.Pointer(src)), n))
	}
	return dest
}

func _memset(tls *TLS, s uintptr, c int32, n uint64) (r uintptr) {
	if n != 0 {
		c := byte(c & 0xff)
		b := unsafe.Slice((*byte)(unsafe.Pointer(s)), n)
		for i := range b {
			b[i] = c
		}
	}
	return s
}

func ___setjmp(tls *TLS, env uintptr) int32 {
	panic(todo(""))
}

func ___unmapself(tls *TLS, base uintptr, size uint64) {
	panic(todo(""))
}

func ___uniclone(tls *TLS, start, stack, new uintptr) int32 {
	panic(todo(""))
}

// ----------------------------------------------------------------------------

// Xfmod computes the floating-point remainder of dividing x by y.  The return
// value is x - n * y, where n is the quotient of x / y, rounded toward zero to
// an integer.
func Xfmod(tls *TLS, x float64, y float64) (r float64) {
	return x_fmod(tls, x, y)
}

// Xsprintf provides formatted output conversion. Xsprintf writes to the
// character string str.  Upon successful return, this functions returns the
// number of characters printed (excluding the null byte used to end output to
// strings).
func Xsprintf(tls *TLS, str uintptr, fmt uintptr, va uintptr) (r int32) {
	return x_sprintf(tls, str, fmt, va)
}

// Xstrtod converts the initial portion of the string pointed to by nptr to
// double representation.
func Xstrtod(tls *TLS, nptr, endptr uintptr) float64 {
	return x_strtod(tls, nptr, endptr)
}

// Xrint rounds its argument to an integer value in floating-point format,
// using the current rounding direction.
func Xrint(tls *TLS, x float64) float64 {
	return x_rint(tls, x)
}

// Xmemset fills the first n bytes of the memory area pointed to by s with the constant byte c.
func Xmemset(tls *TLS, s uintptr, c int32, n uint64) uintptr {
	return _memset(tls, s, c, n)
}

// Xsnprintf writes at most size bytes (including the terminating null byte
// ('\0')) to str, according to a format and args.
func Xsnprintf(tls *TLS, str uintptr, size uint64, format, args uintptr) int32 {
	return x_snprintf(tls, str, size, format, args)
}

// Xfdopen associates a stream with the existing file descriptor, fd.  The mode
// of the stream (one of the values "r", "r+", "w", "w+", "a", "a+") must  be
// compatible  with  the mode  of the file descriptor.  The file position
// indicator of the new stream is set to that belonging to fd, and the error
// and end-of-file indicators are cleared.  Modes "w" or "w+" do not cause
// truncation of the file.  The file descriptor is not dup'ed, and will be
// closed when the stream created by fdopen() is closed.  The result of
// applying fdopen() to a shared memory  object is undefined.
func Xfdopen(tls *TLS, fd int32, mode uintptr) uintptr {
	return x___fdopen(tls, fd, mode)
}

// Xfread reads nmemb items of data, each size bytes long, from the stream
// pointed to by stream, storing them at the location given by ptr.
func Xfread(tls *TLS, ptr uintptr, size, nmemb uint64, stream uintptr) uint64 {
	return x_fread(tls, ptr, size, nmemb, stream)
}

// Xmalloc allocates size bytes and returns a pointer to the allocated
// memory. The memory is not initialized. If size is 0, then Xmalloc returns
// either NULL, or a unique pointer value that can later be successfully passed
// to free().
func Xmalloc(tls *TLS, n uint64) (r uintptr) {
	return x_malloc(tls, n)
}

// The Xree function frees the memory space pointed to by ptr, which must
// have been returned by a previous call to Xmalloc(), Xcalloc(), or Xrealloc().
// Otherwise, or if Xfree(ptr) has  already been called before, undefined
// behavior occurs.  If ptr is NULL, no operation is performed.
func Xfree(tls *TLS, p uintptr) {
	x_free(tls, p)
}

// Xcalloc allocates memory for an array of nmemb elements of size bytes each
// and returns a pointer to the allocated memory.  The memory is set to zero.
// If nmemb or size is 0, then Xcalloc returns either NULL, or a unique pointer
// value that can later be successfully passed to free().  If the
// multiplication of nmemb and size would  result  in  integer  overflow, then
// Xcalloc  returns an error.  By contrast, an integer overflow would not be
// detected in the following call to malloc(), with the result that an
// incorrectly sized block of memory would be allocated:
//
//	Xmalloc(nmemb * size)
func Xcalloc(tls *TLS, nmemb, size uint64) (r uintptr) {
	return x_calloc(tls, nmemb, size)
}

// Xprintf produces output according to a format.  It writes output to stdout,
// the standard output stream;
func Xprintf(tls *TLS, format, va uintptr) (r int32) {
	return x_printf(tls, format, va)
}

// Xsin returns the sine of x, where x is given in radians.
func Xsin(tls *TLS, x float64) (r float64) {
	return x_sin(tls, x)
}

// Xstrcpy  copies the string pointed to by src, including the terminating null
// byte ('\0'), to the buffer pointed to by dest.  The strings may not overlap,
// and the destination string dest must be large enough to receive the copy.
// Beware of buffer overruns!
func Xstrcpy(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	return x_strcpy(tls, dest, src)
}

// Xstrncpy is line Xstrcpy, except that at most n bytes of src are copied.
// Warning: If there is no null byte among the first n bytes of src, the string
// placed in dest will not  be null-terminated.
func Xstrncpy(tls *TLS, dest uintptr, src uintptr, n uint64) (r uintptr) {
	return x_strncpy(tls, dest, src, n)
}

// Xstrcmp compares the two strings s1 and s2. The locale is not taken into
// account (for a locale-aware comparison, see strcoll(3)). The comparison is
// done using unsigned characters.
func Xstrcmp(tls *TLS, s1 uintptr, s2 uintptr) (r int32) {
	return x_strcmp(tls, s1, s2)
}

// Xstrlen function calculates the length of the string pointed to by s,
// excluding the terminating null byte ('\0').
func Xstrlen(tls *TLS, s uintptr) (r uint64) {
	return x_strlen(tls, s)
}

// Xstrcat appends the src string to the dest string, overwriting the
// terminating null byte ('\0') at the end of dest, and then adds a terminating
// null byte. The strings may not overlap, and the dest string must have enough
// space for the result. If dest is not large enough, program behavior is
// unpredictable; buffer overruns are a favorite avenue for attacking secure
// programs.
func Xstrcat(tls *TLS, dest uintptr, src uintptr) (r uintptr) {
	return x_strcat(tls, dest, src)
}

// Xstrncmp is line Xstrcmp, except it compares only the first (at most) n
// bytes of s1 and s2.
func Xstrncmp(tls *TLS, s1 uintptr, s2 uintptr, n uint64) (r1 int32) {
	return x_strncmp(tls, s1, s2, n)
}

// Xstrchr returns a pointer to the first occurrence of the character c in the
// string s.
func Xstrchr(tls *TLS, s uintptr, c int32) (r uintptr) {
	return x_strchr(tls, s, c)
}

// Xstrrchr returns a pointer to the last occurrence of the character c in the
// string s.
func Xstrrchr(tls *TLS, s uintptr, c int32) (r uintptr) {
	return x_strrchr(tls, s, c)
}

// Xmemcpy copies n bytes from memory area src to memory area dest.  The memory
// areas must not overlap.  Use memmove(3) if the memory areas do overlap.
func Xmemcpy(tls *TLS, dest, src uintptr, n uint64) (r uintptr) {
	return _memcpy(tls, dest, src, n)
}

// Xmemcmp compares the first n bytes (each interpreted as unsigned char) of
// the memory areas s1 and s2.
func Xmemcmp(tls *TLS, s1 uintptr, s2 uintptr, n uint64) (r1 int32) {
	return x_memcmp(tls, s1, s2, n)
}

// Xfopen opens the file whose name is the string pointed to by pathname and
// associates a stream with it.
//
// The argument mode points to a string beginning with one of the following
// sequences (possibly followed by additional characters, as described below):
//
//	r
//
// Open text file for reading.  The stream is positioned at the beginning of
// the file.
//
//	r+
//
// Open for reading and writing.  The stream is positioned at the beginning of
// the file.
//
//	w
//
// Truncate file to zero length or create text file for writing.  The stream is
// positioned at the beginning of the file.
//
//	w+
//
// Open for reading and writing.  The file is created if it does not exist,
// otherwise it is truncated.  The stream is positioned at the beginning of the
// file.
//
//	a
//
// Open for appending (writing at end of file).  The file is created if it does
// not exist.  The stream is positioned at the end of the file.
//
//	a+
//
// Open  for  reading and appending (writing at end of file).  The file is
// created if it does not exist.  Output is always appended to the end of the
// file.  POSIX is silent on what the initial read position is when using this
// mode.  For glibc, the initial file position for reading is at the beginning
// of the file, but for Android/BSD/MacOS, the initial  file  position for
// reading is at the end of the file.
//
// The  mode  string can also include the letter 'b' either as a last character
// or as a character between the characters in any of the two-character strings
// described above.  This is strictly for compatibility with C89 and has no
// effect; the 'b' is ignored on all POSIX conforming systems, including Linux.
// (Other systems may treat text files and binary  files  differently,  and
// adding the 'b' may be a good idea if you do I/O to a binary file and expect
// that your program may be ported to non-UNIX environments.)
func Xfopen(tls *TLS, pathname uintptr, mode uintptr) (r uintptr) {
	return x_fopen(tls, pathname, mode)
}

// Xfwrite writes nmemb items of data, each size bytes long, to the stream
// pointed to by stream, obtaining them from the location given by ptr.
func Xfwrite(tls *TLS, ptr uintptr, size uint64, nmemb uint64, stream uintptr) (r uint64) {
	return x_fwrite(tls, ptr, size, nmemb, stream)
}

// Xfclose flushes the stream pointed to by stream (writing any buffered output
// data using fflush(3)) and closes the underlying file descriptor.
func Xfclose(tls *TLS, stream uintptr) (r int32) {
	return x_fclose(tls, stream)
}

// Xfgetc reads the next character from stream and returns it as an unsigned
// char cast to an int, or EOF on end of file or error.
func Xfgetc(tls *TLS, stream uintptr) (r int32) {
	return x_fgetc(tls, stream)
}

// Xfprintf is like printf(), except it writes output to the given output
// stream.
func Xfprintf(tls *TLS, stream uintptr, fmt uintptr, va uintptr) (r int32) {
	return x_fprintf(tls, stream, fmt, va)
}

// Xgetc is equivalent to Xfgetc.
func Xgetc(tls *TLS, stream uintptr) (r int32) {
	return Xfgetc(tls, stream)
}

// Xfgets reads in at most one less than size characters from stream and stores
// them into the buffer pointed to by s. Reading stops after an EOF or a
// newline. If a newline is read, it is stored into the buffer. A terminating
// null byte ('\0') is stored after the last character in the buffer.
func Xfgets(tls *TLS, s uintptr, size int32, stream uintptr) (r uintptr) {
	return x_fgets(tls, s, size, stream)
}

// Xexit causes normal process termination and the least significant byte of
// status (i.e., status & 0xFF) is returned to the parent (see wait(2)).
func Xexit(tls *TLS, code int32) {
	x_exit(tls, code)
}

// Xtolower returns a lowercase equivalent of c, if c is an uppercase letter
// and a lowercase representation exists in the current locale.  Otherwise, it
// returns  c.
func Xtolower(tls *TLS, c int32) (r int32) {
	return x_tolower(tls, c)
}

// Xrealloc function changes the size of the memory block pointed to by ptr to
// size bytes.  The contents will be unchanged in the range from the start of
// the region up to the minimum  of the  old and new sizes.  If the new size is
// larger than the old size, the added memory will not be initialized.  If ptr
// is NULL, then the call is equivalent to Xmalloc(size), for all values of
// size; if size is equal to zero, and ptr is not NULL, then the call is
// equivalent to Xfree(ptr).  Unless ptr is NULL, it must have been returned by
// an earlier call to Xmalloc(), Xcalloc(), or Xrealloc().  If the area pointed
// to was moved, a Xfree(ptr) is done.
func Xrealloc(tls *TLS, ptr uintptr, size uint64) (r uintptr) {
	return x_realloc(tls, ptr, size)
}

// Xfabs returns the absolute value of the floating-point number x.
func Xfabs(tls *TLS, x float64) (r float64) {
	return x_fabs(tls, x)
}

// Xfabsf returns the absolute value of the floating-point number x.
func Xfabsf(tls *TLS, x float32) (r float32) {
	return x_fabsf(tls, x)
}

// X__assert_fail aborts a program.
func X__assert_fail(tls *TLS, expr uintptr, file uintptr, line uint32, func1 uintptr) {
	x___assert_fail(tls, expr, file, int32(line), func1)
}
