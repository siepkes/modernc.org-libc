#include "textflag.h"

// static inline void a_or_64(volatile uint64_t *p, uint64_t v)
TEXT ·a_or_64(SB),NOSPLIT,$0-12
	MOVL	p+0(FP), BX
	MOVL	v+4(FP), AX
	LOCK
	ORL	AX, 0(BX)
	MOVL	v+8(FP), AX
	LOCK
	ORL	AX, 4(BX)
	RET

// static inline void a_and_64(volatile uint64_t *p, uint64_t v)
TEXT ·a_and_64(SB),NOSPLIT,$0-12
	MOVL	p+0(FP), BX
	MOVL	v+4(FP), AX
	LOCK
	ANDL	AX, 0(BX)
	MOVL	v+8(FP), AX
	LOCK
	ANDL	AX, 4(BX)
	RET
