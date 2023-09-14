#include <stdlib.h>
#include <stdint.h>
#include <limits.h>
#include <errno.h>
#include <sys/mman.h>
#include "libc.h"
#include "lock.h"
#include "syscall.h"
#include "fork_impl.h"

#define ALIGN 16


static volatile int lock[1];
volatile int *const __bump_lockptr = lock;


void *__libc_malloc(size_t n)
{
	return __libc_malloc_impl(n);
}
