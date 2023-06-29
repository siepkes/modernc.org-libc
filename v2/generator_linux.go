// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// /tmp/libc-v2-generate2882489557/ : temporary directory kept

// /tmp/libc-v2-generate2882489557/musl-1.2.4/src/linux

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
	util.MustInDir(true, muslRoot, func() (err error) {
		var cflags string
		if s := cc.LongDouble64Flag(goos, goarch); s != "" {
			cflags = fmt.Sprintf("CFLAGS=%s", s)
		}
		util.MustShell(true, "sh", "-c", fmt.Sprintf("CC=%s %s ./configure --disable-static --disable-optimize", cCompiler, cflags))
		util.MustShell(true, "find", "src/string", "-name", "*.s", "-delete")
		return ccgo.NewTask(
			goos, goarch,
			[]string{
				os.Args[0],
				"-exec-cc", cCompiler,
				"-extended-errors",
				"-ignore-asm-errors",
				"-ignore-header-functions",
				"-ignore-unsupported-alignment",
				"--package-name=libc",
				// "-positions",
				"--prefix-enumerator=E",
				"--prefix-external=X",
				"--prefix-field=F",
				"--prefix-macro=M",
				"--prefix-static-internal=_",
				"--prefix-static-none=_",
				"--prefix-tagged-enum=TE",
				"--prefix-tagged-struct=TS",
				"--prefix-tagged-union=TU",
				"--prefix-typename=TN",
				"--prefix-undefined=_",
				"-exec", "make", // keep last
			},
			os.Stdout, os.Stderr, nil,
		).Main()
	})
}
