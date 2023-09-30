// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

// https://posixtest.sourceforge.net/

// ~/tmp/musl/musl-1.2.4/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
)

const (
	archivePath = "master.zip"
)

var (
	extractedArchivePath string
	goarch               = or(os.Getenv("GO_GENERATE_GOARCH"), runtime.GOARCH)
	goos                 = runtime.GOOS
	j                    = fmt.Sprint(runtime.GOMAXPROCS(-1))
	muslArch             string
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func or(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

func main() {
	if goos != "linux" {
		return
	}

	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	util.Shell("find", filepath.Join("internal", "nsz.repo.hu"), "-name", "*.go", "-delete")
	f, err := os.Open(archivePath)
	if err != nil {
		fail(1, "cannot open zip file: %v\n", err)
	}

	f.Close()

	switch goarch {
	case "386":
		muslArch = "i386"
	case "amd64":
		muslArch = "x86_64"
	case "arm":
		muslArch = "arm"
	case "arm64":
		muslArch = "aarch64"
	case "loong64":
		muslArch = "mips"
	case "ppc64le":
		muslArch = "powerpc64"
	case "riscv64":
		muslArch = "riscv64"
	case "s390x":
		muslArch = "s390x"
	default:
		fail(1, "unsupported goarch: %s", goarch)
	}

	extractedArchivePath = "musl-master"
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev := os.Getenv("GO_GENERATE_DEV") != ""
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s", filepath.Join(tempDir, extractedArchivePath)))
	default:
		var err error
		if tempDir, err = os.MkdirTemp("", "libc-v2-generate"); err != nil {
			fail(1, "creating temp dir: %v\n", err)
		}

		defer func() {
			switch os.Getenv("GO_GENERATE_KEEP") {
			case "":
				os.RemoveAll(tempDir)
			default:
				fmt.Printf("%s: temporary directory kept\n", tempDir)
			}
		}()
	}
	libRoot := filepath.Join(tempDir, extractedArchivePath)
	makeRoot := libRoot
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)
	fmt.Fprintf(os.Stderr, "libRoot %s\n", libRoot)
	fmt.Fprintf(os.Stderr, "makeRoot %s\n", makeRoot)

	util.MustShell(true, "unzip", archivePath, "-d", tempDir)
	mustCopyDir(libRoot, filepath.Join("overlay", extractedArchivePath, "all"), nil, true)
	mustCopyDir(libRoot, filepath.Join("overlay", extractedArchivePath, goarch), nil, true)
	mustCopyFile("COPYRIGHT-MUSL", filepath.Join(makeRoot, "COPYRIGHT"), nil)
	result := "libc.a.go"
	util.MustInDir(true, makeRoot, func() (err error) {
		cflags := []string{
			"-DNDEBUG",
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure "+
			"--disable-shared "+
			"--disable-optimize "+
			fmt.Sprintf("--target %s ", muslArch)+
			"",
			strings.Join(cflags, " "),
		))
		args := []string{
			os.Args[0],

			"--package-name=libc",
			"--predef=float __builtin_inff(void);",
			"--predef=long __builtin_expect(long, long);",
			"--prefix-enumerator=_",
			"--prefix-external=x_",
			"--prefix-field=F",
			"--prefix-macro=m_",
			"--prefix-static-internal=_",
			"--prefix-static-none=_",
			"--prefix-tagged-enum=_",
			"--prefix-tagged-struct=T",
			"--prefix-tagged-union=T",
			"--prefix-typename=T",
			"--prefix-undefined=_",
			"-extended-errors",
			"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6,__get_tp,__DOUBLE_BITS,__FLOAT_BITS",
			"-hide", "a_and,a_and_64,a_barrier,a_cas,a_cas_p,a_clz_64,a_crash,a_ctz_64,a_dec,a_fetch_add,a_inc,a_or,a_or_64,a_spin,a_store,a_swap,a_ctz_32",
			"-hide", "fabs,fabsf,fabsl,sqrt,sqrtf,sqrtl",
			"-hide", "_Fork,fork,system,__synccall,__set_thread_area",
			"-hide", "calloc,free,malloc,malloc_usable_size,realloc",
			"-hide", "__libc_calloc,__libc_free,__libc_malloc,__libc_malloc_impl,__libc_realloc",
			"-hide", "__malloc_allzerop,__malloc_atfork,__malloc_donate,__simple_malloc,",
			"-ignore-asm-errors",
			"-isystem", "",
		}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-keep-object-files",
				"-positions",
			)
		}
		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "-j", j, "lib/libc.a"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			return err
		}

		return ccgo.NewTask(goos, goarch, append(args, "-o", result, "-nostdlib", "-ignore-link-errors", "lib/libc.a"), os.Stdout, os.Stderr, nil).Main()
	})

	os.RemoveAll(filepath.Join("include", goos, goarch))
	util.MustCopyDir(true, filepath.Join("include", goos, goarch), filepath.Join(tempDir, extractedArchivePath, "include"), nil)
	util.MustCopyDir(true, filepath.Join("include", goos, goarch, "bits"), filepath.Join(tempDir, extractedArchivePath, "obj", "include", "bits"), nil)
	util.MustCopyDir(true, filepath.Join("include", goos, goarch, "bits"), filepath.Join(tempDir, extractedArchivePath, "arch", "generic", "bits"), nil)
	util.MustCopyDir(true, filepath.Join("include", goos, goarch, "bits"), filepath.Join(tempDir, extractedArchivePath, "arch", muslArch, "bits"), nil)

	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	util.MustShell(true, "cp", filepath.Join(makeRoot, result), fn)
	util.MustShell(true, "sed", "-i", `s/\<T__\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/t__\1/g`, fn)
	util.MustShell(true, "sed", "-i", `s/\<x_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/X\1/g`, fn)
	util.MustShell(true, "sed", "-i", `s/\<Xpthread_\([a-zA-Z0-9][a-zA-Z0-9_]\+\)/x_pthread_\1/g`, fn)
	util.MustShell(true, "sed", "-i", `s/\<x___errno_location\>/X__errno_location/g`, fn)
	for _, v := range []struct{ old, new string }{
		{"___libc_calloc", "Xcalloc"},
		{"___libc_free", "Xfree"},
		{"___libc_malloc", "Xmalloc"},
		{"___libc_malloc_impl", "Xmalloc"},
		{"_calloc", "Xcalloc"},
		{"_free", "Xfree"},
		{"_malloc", "Xmalloc"},
		{"_realloc", "Xrealloc"},
		{"x__Exit", "X_Exit"},
		{"x___clock_gettime", "Xclock_gettime"},
		{"x___ctype_get_mb_cur_max", "X__ctype_get_mb_cur_max"},
		{"x___dn_expand", "Xdn_expand"},
		{"x___environ", "Xenviron"},
		{"x___fdopen", "Xfdopen"},
		{"x___fseeko", "Xfseeko"},
		{"x___fstat", "Xfstat"},
		{"x___ftello", "Xftello"},
		{"x___gmtime_r", "Xgmtime_r"},
		{"x___libc_current_sigrtmax", "X__libc_current_sigrtmax"},
		{"x___libc_current_sigrtmin", "X__libc_current_sigrtmin"},
		{"x___localtime_r", "Xlocaltime_r"},
		{"x___lseek", "Xlseek"},
		{"x___mmap", "Xmmap"},
		{"x___mremap", "Xmremap"},
		{"x___munmap", "Xmunmap"},
		{"x___newlocale", "Xnewlocale"},
		{"x___nl_langinfo", "Xnl_langinfo"},
		{"x___sigaction", "Xsigaction"},
		{"x___uselocale", "Xuselocale"},
	} {
		util.MustShell(true, "sed", "-i", fmt.Sprintf(`s/\<%s\>/%s/g`, v.old, v.new), fn)
	}

	m, err := filepath.Glob(fmt.Sprintf("*_%s_%s.go", runtime.GOOS, runtime.GOARCH))
	if err != nil {
		fail(1, "%s\n", err)
	}

	format := false
	for _, fn := range m {
		b, err := os.ReadFile(fn)
		if err != nil {
			fail(1, "%s\n", err)
		}

		a := strings.Split(string(b), "\n")
		w := false
		for i, v := range a {
			if strings.HasPrefix(v, "func X") {
				if i+1 < len(a) && !strings.Contains(a[i+1], "__ccgo_strace") {
					a[i] += "\n\t" + traceLine(v)
					w = true
					format = true
				}
			}
		}
		if w {
			if err := os.WriteFile(fn, []byte(strings.Join(a, "\n")), 0660); err != nil {
				fail(1, "%s\n", err)
			}
		}
	}
	if format {
		util.MustShell(true, "sh", "-c", "gofmt -w *.go")
	}
	util.Shell("sh", "-c", "./unconvert.sh")
	util.MustShell(true, "go", "test", "-run", "@")
	util.Shell("git", "status")
}

// func Xaio_fsync(tls *TLS, op int32, cb uintptr) (r int32) {
func traceLine(s string) string {
	var b strings.Builder
	parts := strings.Split(s, "(")
	for i, v := range parts {
		switch i {
		case 0:
			// "func Xaio_fsync"
		case 1:
			// "tls *TLS, op int32, cb uintptr) "
			a := strings.Split(v, ",")
			b.WriteString(`if __ccgo_strace { trc("`)
			var vals []string
			for j, w := range a {
				w = strings.TrimSpace(w)
				if x := strings.Index(w, " "); x > 0 {
					w := w[:x]
					if j != 0 {
						b.WriteString(" ")
					}
					fmt.Fprintf(&b, "%s=%%v", w)
					vals = append(vals, w)
				}
			}
			fmt.Fprintf(&b, `, (%%v:)", %s, origin(2))`, strings.Join(vals, ", "))
		case 2:
			// "r int32) {"
			r := v[:strings.Index(v, " ")]
			fmt.Fprintf(&b, `; defer func() { trc("-> %%v", %s)}()`, r)
		}
	}
	b.WriteString("}")
	return b.String()
}

func mustCopyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64) {
	file, bytes, err := copyDir(dst, src, canOverwrite, srcNotExistsOk)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return file, bytes
}

func copyDir(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool, srcNotExistsOk bool) (files int, bytes int64, rerr error) {
	dst = filepath.FromSlash(dst)
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) && srcNotExistsOk {
			err = nil
		}
		return 0, 0, err
	}

	if !si.IsDir() {
		return 0, 0, fmt.Errorf("cannot copy a file: %s", src)
	}

	return files, bytes, filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode()&os.ModeSymlink != 0 {
			target, err := filepath.EvalSymlinks(path)
			if err != nil {
				return fmt.Errorf("cannot evaluate symlink %s: %v", path, err)
			}

			if info, err = os.Stat(target); err != nil {
				return fmt.Errorf("cannot stat %s: %v", target, err)
			}

			if info.IsDir() {
				rel, err := filepath.Rel(src, path)
				if err != nil {
					return err
				}

				dst2 := filepath.Join(dst, rel)
				if err := os.MkdirAll(dst2, 0770); err != nil {
					return err
				}

				f, b, err := copyDir(dst2, target, canOverwrite, srcNotExistsOk)
				files += f
				bytes += b
				return err
			}

			path = target
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return os.MkdirAll(filepath.Join(dst, rel), 0770)
		}

		n, err := copyFile(filepath.Join(dst, rel), path, canOverwrite)
		if err != nil {
			return err
		}

		files++
		bytes += n
		return nil
	})
}

func mustCopyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) int64 {
	n, err := copyFile(dst, src, canOverwrite)
	if err != nil {
		fail(1, "%s\n", err)
	}

	return n
}

func copyFile(dst, src string, canOverwrite func(fn string, fi os.FileInfo) bool) (n int64, rerr error) {
	src = filepath.FromSlash(src)
	si, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if si.IsDir() {
		return 0, fmt.Errorf("cannot copy a directory: %s", src)
	}

	dst = filepath.FromSlash(dst)
	if si.Size() == 0 {
		return 0, os.Remove(dst)
	}

	dstDir := filepath.Dir(dst)
	di, err := os.Stat(dstDir)
	switch {
	case err != nil:
		if !os.IsNotExist(err) {
			return 0, err
		}

		if err := os.MkdirAll(dstDir, 0770); err != nil {
			return 0, err
		}
	case err == nil:
		if !di.IsDir() {
			return 0, fmt.Errorf("cannot create directory, file exists: %s", dst)
		}
	}

	di, err = os.Stat(dst)
	switch {
	case err != nil && !os.IsNotExist(err):
		return 0, err
	case err == nil:
		if di.IsDir() {
			return 0, fmt.Errorf("cannot overwite a directory: %s", dst)
		}

		if canOverwrite != nil && !canOverwrite(dst, di) {
			return 0, fmt.Errorf("cannot overwite: %s", dst)
		}
	}

	s, err := os.Open(src)
	if err != nil {
		return 0, err
	}

	defer s.Close()
	r := bufio.NewReader(s)

	d, err := os.Create(dst)

	defer func() {
		if err := d.Close(); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chmod(dst, si.Mode()); err != nil && rerr == nil {
			rerr = err
			return
		}

		if err := os.Chtimes(dst, si.ModTime(), si.ModTime()); err != nil && rerr == nil {
			rerr = err
			return
		}
	}()

	w := bufio.NewWriter(d)

	defer func() {
		if err := w.Flush(); err != nil && rerr == nil {
			rerr = err
		}
	}()

	return io.Copy(w, r)
}
