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
		var cflags string
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = fmt.Sprintf("CFLAGS=%s", s)
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("%s ./configure --disable-shared --disable-optimize", cflags))
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
			"-hide", "fork,system",
			"-ignore-asm-errors",
			"-isystem", "",
		}
		if dev {
			args = append(
				args,
				"-absolute-paths",
				"-positions",
				// "-verify-types",
			)
		}
		if err := ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "lib/libc.a"), os.Stdout, os.Stderr, nil).Exec(); err != nil {
			return err
		}

		return ccgo.NewTask(goos, goarch, append(args, "-o", result, "-nostdlib", "-ignore-link-errors", "lib/libc.a"), os.Stdout, os.Stderr, nil).Main()
	})

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
	for _, v := range []string{
		"environ",
		"fdopen",
		"fstat",
		"gmtime_r",
		"localtime_r",
		"lseek",
		"mmap",
		"mremap",
		"munmap",
		"nl_langinfo",
		"sigaction",
	} {
		util.MustShell(true, "sed", "-i", fmt.Sprintf(`s/\<x___%s\>/X%[1]s/g`, v), fn)
	}
}
