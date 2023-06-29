#include "textflag.h"

// static inline int a_cas(volatile int *p, int t, int s)
// {
// 	__asm__ __volatile__ (
// 		"lock ; cmpxchg %3, %1"
// 		: "=a"(t), "=m"(*p) : "a"(t), "r"(s) : "memory" );
// 	return t;
// }
TEXT ·a_cas(SB),NOSPLIT,$0-20
	MOVQ	p+0(FP), BX
	MOVL	t+8(FP), AX
	MOVL	s+12(FP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	MOVL	AX, ret+16(FP)
	RET

// static inline void *a_cas_p(volatile void *p, void *t, void *s)
// {
// 	__asm__( "lock ; cmpxchg %3, %1"
// 		: "=a"(t), "=m"(*(void *volatile *)p)
// 		: "a"(t), "r"(s) : "memory" );
// 	return t;
// }
TEXT ·a_cas_p(SB),NOSPLIT,$0-32
	MOVQ	p+0(FP), BX
	MOVQ	t+8(FP), AX
	MOVQ	s+16(FP), CX
	LOCK
	CMPXCHGQ	CX, 0(BX)
	MOVQ	AX, ret+24(FP)
	RET
