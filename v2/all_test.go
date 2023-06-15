// Copyright 2023 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/libc/v2"

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func Test(t *testing.T) {
	t.Log("TODO")
}
