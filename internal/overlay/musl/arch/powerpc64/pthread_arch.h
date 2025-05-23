static inline uintptr_t __get_tp();

#ifndef __CCGO__
static inline uintptr_t __get_tp()
{
	register uintptr_t tp __asm__("r13");
	__asm__ ("" : "=r" (tp) );
	return tp;
}
#endif

#define TLS_ABOVE_TP
#define GAP_ABOVE_TP 0

#define TP_OFFSET 0x7000
#define DTP_OFFSET 0x8000

// the kernel calls the ip "nip", it's the first saved value after the 32
// GPRs.
#define MC_PC gp_regs[32]
