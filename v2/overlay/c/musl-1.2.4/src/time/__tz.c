#include "time_impl.h"
#include <stdint.h>
#include <limits.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mman.h>
#include <ctype.h>
#include "libc.h"
#include "lock.h"
#include "fork_impl.h"
#include <assert.h>

#define malloc __libc_malloc
#define calloc undef
#define realloc undef
#define free undef

long  __timezone = 0;
int   __daylight = 0;
char *__tzname[2] = { 0, 0 };

weak_alias(__timezone, timezone);
weak_alias(__daylight, daylight);
weak_alias(__tzname, tzname);

const char __utc[] = "UTC";

static volatile int lock[1];
volatile int *const __timezone_lockptr = lock;


#define VEC(...) ((const unsigned char[]){__VA_ARGS__})

void __secs_to_zone(long long t, int local, int *isdst, long *offset, long *oppoff, const char **zonename)
{
	assert(0);
}

static void __tzset()
{
	assert(0);
}

weak_alias(__tzset, tzset);

const char *__tm_to_tzname(const struct tm *tm)
{
	assert(0);
}
