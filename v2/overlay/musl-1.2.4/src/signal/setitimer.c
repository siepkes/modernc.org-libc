#include <sys/time.h>
#include <errno.h>
#include "syscall.h"
#include <assert.h>

#define IS32BIT(x) !((x)+0x80000000ULL>>32)

int setitimer(int which, const struct itimerval *restrict new, struct itimerval *restrict old)
{
	assert(0);
}
