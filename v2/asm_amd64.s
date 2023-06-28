#include "textflag.h"

// static inline int a_cas(volatile int *p, int t, int s)
// {
// 	__asm__ __volatile__ (
// 		"lock ; cmpxchg %3, %1"
// 		: "=a"(t), "=m"(*p) : "a"(t), "r"(s) : "memory" );
// 	return t;
// }
TEXT ·a_cas(SB),NOSPLIT,$0-24
	MOVQ	p+0(FP), BX
	MOVL	t+8(FP), AX
	MOVL	s+12(FP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	MOVL	AX, ret+16(FP)
	RET
