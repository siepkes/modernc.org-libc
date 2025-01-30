// Copyright 2025 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc"

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	ccgo "modernc.org/ccgo/v4/lib"
)

// https://gitlab.com/cznic/libc/-/issues/42
func TestIssue42(t *testing.T) {
	dir := t.TempDir()

	defer os.Remove("test_pread_pwrite.txt")

	gof := filepath.Join(dir, "main.go")
	if err := ccgo.NewTask(
		goos, goarch,
		[]string{
			os.Args[0],
			"-o", gof,
			filepath.Join("testdata", "pread_test.c"),
		},
		os.Stdout, os.Stderr,
		nil,
	).Main(); err != nil {
		t.Fatal(err)
	}

	b, err := exec.Command("go", "run", gof).CombinedOutput()
	if err != nil {
		t.Fatalf("FAIL err=%v\n%s", err, b)
	}

	t.Logf("%s", b)
}
