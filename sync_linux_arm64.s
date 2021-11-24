// Copied from Go runtime/sys_linux_arm.s

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-GO file.

#include "go_asm.h"
// #include "go_tls.h"
#include "textflag.h"

// As for cas, memory barriers are complicated on ARM, but the kernel
// provides a user helper. ARMv5 does not support SMP and has no
// memory barrier instruction at all. ARMv6 added SMP support and has
// a memory barrier, but it requires writing to a coprocessor
// register. ARMv7 introduced the DMB instruction, but it's expensive
// even on single-core devices. The kernel helper takes care of all of
// this for us.

TEXT kernelPublicationBarrier<>(SB),NOSPLIT,$0
	// void __kuser_memory_barrier(void);
	MOVW	$0xffff0fa0, R11
	CALL	(R11)
	RET
