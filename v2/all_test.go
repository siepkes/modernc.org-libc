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

// 	__builtin_printf("\t%d\n", i);
// 	__builtin_printf("%d\n", a_cas(&i, 1, 2));
// 	__builtin_printf("\t%d\n", i);
// 	__builtin_printf("%d\n", a_cas(&i, 0, 3));
// 	__builtin_printf("\t%d\n", i);
// 	__builtin_printf("%d\n", a_cas(&i, 4, 5));
// 	__builtin_printf("\t%d\n", i);
// 	__builtin_printf("%d\n", a_cas(&i, 3, 6));
// 	__builtin_printf("\t%d\n", i);

func TestAtomicCAS(t *testing.T) {
	i := int32(0)
	j := a_cas(&i, 1, 2)
	if i != 0 || j != 0 {
		t.Fatal(i, j)
	}

	if j = a_cas(&i, 0, 3); i != 3 || j != 0 {
		t.Fatal(i, j)
	}

	if j = a_cas(&i, 4, 5); i != 3 || j != 3 {
		t.Fatal(i, j)
	}

	if j = a_cas(&i, 3, 6); i != 6 || j != 3 {
		t.Fatal(i, j)
	}
}
