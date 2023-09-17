// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

// https://posixtest.sourceforge.net/

// ~/tmp/musl/musl-1.2.4/

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"modernc.org/cc/v4"
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
	"modernc.org/ccorpus2"
)

const (
	defaultArchivePath = "assets/musl.libc.org/releases/musl-1.2.4.tar.gz"
)

var (
	archivePath          = defaultArchivePath
	extractedArchivePath string
	goarch               = runtime.GOARCH
	goos                 = runtime.GOOS
	muslArch             string

	dev bool
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	if dev {
		fmt.Fprintf(os.Stderr, "%s\n", debug.Stack())
	}
	os.Exit(rc)
}

func main() {
	if ccgo.IsExecEnv() {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	var f fs.File
	var err error

	switch s := os.Getenv("GO_GENERATE_ARCHIVE"); {
	case s != "":
		archivePath = s
		if f, err = os.Open(archivePath); err != nil {
			fail(1, "cannot open tar file: %v\n", err)
		}
	default:
		if f, err = ccorpus2.FS.Open(archivePath); err != nil {
			fail(1, "cannot open tar file: %v\n", err)
		}

	}
	switch goarch {
	case "amd64":
		muslArch = "x86_64"
	default:
		fail(1, "unsupported goarch: %s", goarch)
	}
	_, extractedArchivePath = filepath.Split(archivePath)
	extractedArchivePath = extractedArchivePath[:len(extractedArchivePath)-len(".tar.gz")]
	tempDir := os.Getenv("GO_GENERATE_DIR")
	dev = os.Getenv("GO_GENERATE_DEV") != ""
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
	fmt.Fprintf(os.Stderr, "archivePath %s\n", archivePath)
	fmt.Fprintf(os.Stderr, "extractedArchivePath %s\n", extractedArchivePath)
	fmt.Fprintf(os.Stderr, "tempDir %s\n", tempDir)

	util.MustUntar(true, tempDir, f, nil)
	muslRoot := filepath.Join(tempDir, extractedArchivePath)
	util.MustCopyDir(true, filepath.Join(tempDir, extractedArchivePath), filepath.Join("overlay", extractedArchivePath, "all"), nil)
	util.MustCopyDir(true, filepath.Join(tempDir, extractedArchivePath), filepath.Join("overlay", extractedArchivePath, goarch), nil)
	util.MustCopyFile(true, "COPYRIGHT-MUSL", filepath.Join(muslRoot, "COPYRIGHT"), nil)
	result := "libc.a.go"
	util.MustInDir(true, muslRoot, func() (err error) {
		cflags := []string{
			// "-UNDEBUG", //TODO-
		}
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = append(cflags, s)
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CFLAGS='%s' ./configure --disable-shared --disable-optimize", strings.Join(cflags, " ")))
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
			"-hide", "clone,_Fork,fork,system,__synccall",
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
		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "lib/libc.a"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
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
	util.MustShell(true, "cp", filepath.Join(muslRoot, result), fn)
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
		{"x___environ", "Xenviron"},
		{"x___fdopen", "Xfdopen"},
		{"x___fstat", "Xfstat"},
		{"x___gmtime_r", "Xgmtime_r"},
		{"x___localtime_r", "Xlocaltime_r"},
		{"x___lseek", "Xlseek"},
		{"x___mmap", "Xmmap"},
		{"x___mremap", "Xmremap"},
		{"x___munmap", "Xmunmap"},
		{"x___nl_langinfo", "Xnl_langinfo"},
		{"x___sigaction", "Xsigaction"},
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
		util.MustShell(true, "gofmt", "-w", ".")
	}
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
