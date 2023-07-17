// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// ~/tmp/musl/musl-1.2.4/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"modernc.org/cc/v4"
	util "modernc.org/ccgo/v3/lib"
	ccgo "modernc.org/ccgo/v4/lib"
	"modernc.org/ccorpus2"
)

const (
	archivePath          = "assets/musl.libc.org/releases/musl-1.2.4.tar.gz"
	extractedArchivePath = "musl-1.2.4"
	muslIncludes         = "include"
)

var (
	cCompiler = "gcc"
	goarch    = runtime.GOARCH
	goos      = runtime.GOOS
)

func fail(rc int, msg string, args ...any) {
	fmt.Fprintln(os.Stderr, strings.TrimSpace(fmt.Sprintf(msg, args...)))
	os.Exit(rc)
}

func main() {

	if os.Getenv(ccgo.CCEnvVar) != "" {
		if err := ccgo.NewTask(goos, goarch, os.Args, os.Stdout, os.Stderr, nil).Main(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	tempDir := os.Getenv("GO_GENERATE_DIR")
	switch {
	case tempDir != "":
		util.MustShell(true, "sh", "-c", fmt.Sprintf("rm -rf %s/*", tempDir))
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

	f, err := ccorpus2.FS.Open(archivePath)
	if err != nil {
		fail(1, "cannot open tar file: %v\n", err)
	}

	util.MustUntar(true, tempDir, f, nil)
	muslRoot := filepath.Join(tempDir, extractedArchivePath)
	util.MustCopyDir(true, tempDir, filepath.Join("overlay", "c"), nil)
	util.MustCopyFile(true, "COPYRIGHT-MUSL", filepath.Join(muslRoot, "COPYRIGHT"), nil)
	result := filepath.FromSlash("lib/libc.so.go")
	util.MustInDir(true, muslRoot, func() (err error) {
		var cflags string
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = fmt.Sprintf("CFLAGS=%s", s)
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CC=%s %s ./configure --disable-static --disable-optimize", cCompiler, cflags))
		if err := ccgo.NewTask(
			goos, goarch,
			[]string{
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
				"-exec-cc", cCompiler,
				"-extended-errors",
				"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6,__get_tp,__DOUBLE_BITS,__FLOAT_BITS",
				"-hide", "a_and,a_and_64,a_barrier,a_cas,a_cas_p,a_clz_64,a_crash,a_ctz_64,a_dec,a_fetch_add,a_inc,a_or,a_or_64,a_spin,a_store,a_swap,a_ctz_32",
				"-hide", "fabs,fabsf,fabsl",
				"-ignore-asm-errors",
				"-ignore-unsupported-alignment",
				"-ignore-unsupported-atomic-sizes",
				// "-positions", "-full-paths",
				"-exec", "make", // keep last
			},
			os.Stdout, os.Stderr, nil,
		).Main(); err != nil {
			return err
		}
		util.MustShell(true, "sed", "-i", `0,/^)$/{s/^)$/\n\t. "modernc.org\/libc\/v2\/rtl"\n)/}`, result)
		return nil
	})
	util.MustShell(true, "cp", filepath.Join(muslRoot, result), fmt.Sprintf("libc_so_%s_%s.go", goos, goarch))
}
