// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// ~/tmp/musl/musl-1.2.4/

package main

// https://posixtest.sourceforge.net/

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
	defaultCCompiler   = "gcc"
)

var (
	archivePath          = defaultArchivePath
	cCompiler            = defaultCCompiler
	extractedArchivePath string
	goarch               = runtime.GOARCH
	goos                 = runtime.GOOS

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
	if os.Getenv(ccgo.CCEnvVar) != "" {
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
	util.MustCopyDir(true, filepath.Join(tempDir, extractedArchivePath), filepath.Join("overlay", extractedArchivePath, goarch), nil)
	util.MustCopyFile(true, "COPYRIGHT-MUSL", filepath.Join(muslRoot, "COPYRIGHT"), nil)
	result := filepath.FromSlash("lib/libc.so.go")
	util.MustInDir(true, muslRoot, func() (err error) {
		var cflags string
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = fmt.Sprintf("CFLAGS=%s", s)
		}
		switch extractedArchivePath {
		case "musl-0.6.0":
			cCompiler = "cc"
		case "musl-1.2.4":
			util.MustShell(true, "sh", "-c", fmt.Sprintf("CC=%s %s ./configure --disable-static --disable-optimize", cCompiler, cflags))
		default:
			fail(1, "unsupported musl version: %s", extractedArchivePath)
		}
		args := []string{os.Args[0]}
		if dev {
			args = append(args, "-absolute-paths", "-positions")
		}
		args = append(args,
			"--package-name=libc",
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
			"-exec-cc", cCompiler,
			"-extended-errors",
			"-ignore-asm-errors",               //TODO- it is possible
			"-ignore-unsupported-alignment",    //TODO- only if possible
			"-ignore-unsupported-atomic-sizes", //TODO- it is possible
		)
		switch extractedArchivePath {
		case "musl-0.6.0":
			args = append(args,
				"-hide", "a_inc,a_dec,a_swap,a_store,a_ctz_64,a_and_64,a_or_64,a_spin,a_fetch_add,a_cas_p",
				"-hide", "syscall0,syscall1,syscall2,syscall3,syscall4,syscall5,syscall6",
				"-hide", "__pthread_self,sqrt,system",
			)
			return ccgo.NewTask(goos, goarch, append(args, "-exec", "make", "lib/libc.so"), os.Stdout, os.Stderr, nil).Main()
		// case "musl-1.2.4":
		// 	// Arguments when archiving/linking static and dynamic musl libc
		// 	//	$ diff -u static dynamic
		// 	//	--- static	2023-07-23 13:40:26.325791570 +0200
		// 	//	+++ dynamic	2023-07-23 13:42:18.021633710 +0200
		// 	//	@@ -1339,3 +1339,5 @@
		// 	//	 obj/src/unistd/usleep.lo
		// 	//	 obj/src/unistd/write.lo
		// 	//	 obj/src/unistd/writev.lo
		// 	//	+obj/ldso/dlstart.lo
		// 	//	+obj/ldso/dynlink.lo
		// 	//	$
		// 	args = append(args,
		// 		"--predef=float __builtin_inff(void);",
		// 		"--predef=long __builtin_expect(long, long);",
		// 		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6,__get_tp,__DOUBLE_BITS,__FLOAT_BITS",
		// 		"-hide", "a_and,a_and_64,a_barrier,a_cas,a_cas_p,a_clz_64,a_crash,a_ctz_64,a_dec,a_fetch_add,a_inc,a_or,a_or_64,a_spin,a_store,a_swap,a_ctz_32",
		// 		"-hide", "fabs,fabsf,fabsl",
		// 		"-ignore-file=obj/ldso/dlstart.lo.go,obj/ldso/dynlink.lo.go",
		// 	)
		// 	return ccgo.NewTask(goos, goarch, append(args, "-exec", "make"), os.Stdout, os.Stderr, nil).Main()
		default:
			fail(1, "unsupported musl version: %s", extractedArchivePath)
		}
		panic("unreachable")
	})
	fn := fmt.Sprintf("ccgo_%s_%s.go", goos, goarch)
	util.MustShell(true, "cp", filepath.Join(muslRoot, result), fn)
	util.MustShell(true, "sed", "-i", "s/\\<x_stdout\\>/Xstdout/g", fn)
	util.MustShell(true, "sed", "-i", "s/\\<x_stderr\\>/Xstderr/g", fn)
	util.MustShell(true, "sed", "-i", "s/\\<x_stdin\\>/Xstdin/g", fn)
}
