#include "textflag.h"

// static inline int a_cas(volatile int *p, int t, int s)
TEXT ·a_cas(SB),NOSPLIT,$0-20
	MOVQ	p+0(FP), BX
	MOVL	t+8(FP), AX
	MOVL	s+12(FP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	MOVL	AX, ret+16(FP)
	RET

// static inline void *a_cas_p(volatile void *p, void *t, void *s)
TEXT ·a_cas_p(SB),NOSPLIT,$0-32
	MOVQ	p+0(FP), BX
	MOVQ	t+8(FP), AX
	MOVQ	s+16(FP), CX
	LOCK
	CMPXCHGQ	CX, 0(BX)
	MOVQ	AX, ret+24(FP)
	RET

// static inline void a_or(volatile int *p, int v)
TEXT ·a_or(SB),NOSPLIT,$0-12
	MOVQ	p+0(FP), BX
	MOVL	v+8(FP), AX
	LOCK
	ORL	AX, 0(BX)
	RET

// static inline void a_or_64(volatile uint64_t *p, uint64_t v)
TEXT ·a_or_64(SB),NOSPLIT,$0-16
	MOVQ	p+0(FP), BX
	MOVQ	v+8(FP), AX
	LOCK
	ORQ	AX, 0(BX)
	RET

// static inline void a_and_64(volatile uint64_t *p, uint64_t v)
TEXT ·a_and_64(SB),NOSPLIT,$0-16
	MOVQ	p+0(FP), BX
	MOVQ	v+8(FP), AX
	LOCK
	ANDQ	AX, 0(BX)
	RET
