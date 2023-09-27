#include "textflag.h"

// static inline int a_cas(volatile int *p, int t, int s)
// armv7
TEXT ·a_cas(SB),NOSPLIT,$0
	MOV R3,R0
	DMB ISH
L1:	LDREX R0,[R2]
	SUBS R0,R3,R0
	STREXEQ R0,R1,[R2]
	TEQEQ R0,$1
	BEQ L1
	DMB ISH
	BX LR
