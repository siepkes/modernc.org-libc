#define _GNU_SOURCE
#include "pthread_impl.h"
#include "stdio_impl.h"
#include "libc.h"
#include "lock.h"
#include <sys/mman.h>
#include <string.h>
#include <stddef.h>
#include <assert.h>

static void dummy_0()
{
}
weak_alias(dummy_0, __acquire_ptc);
weak_alias(dummy_0, __release_ptc);
weak_alias(dummy_0, __pthread_tsd_run_dtors);
weak_alias(dummy_0, __do_orphaned_stdio_locks);
weak_alias(dummy_0, __dl_thread_cleanup);
weak_alias(dummy_0, __membarrier_init);

void __tl_lock(void)
{
	assert(0);
}

void __tl_unlock(void)
{
	assert(0);
}

void __tl_sync(pthread_t td)
{
	assert(0);
}

_Noreturn void __pthread_exit(void *result)
{
	assert(0);
}

void __do_cleanup_push(struct __ptcb *cb)
{
	assert(0);
}

void __do_cleanup_pop(struct __ptcb *cb)
{
	assert(0);
}

struct start_args {
	void *(*start_func)(void *);
	void *start_arg;
	volatile int control;
	unsigned long sig_mask[_NSIG/8/sizeof(long)];
};

#define ROUND(x) (((x)+PAGE_SIZE-1)&-PAGE_SIZE)

/* pthread_key_create.c overrides this */
static volatile size_t dummy = 0;
weak_alias(dummy, __pthread_tsd_size);
static void *dummy_tsd[1] = { 0 };
weak_alias(dummy_tsd, __pthread_tsd_main);

static FILE *volatile dummy_file = 0;
weak_alias(dummy_file, __stdin_used);
weak_alias(dummy_file, __stdout_used);
weak_alias(dummy_file, __stderr_used);

int __pthread_create(pthread_t *restrict res, const pthread_attr_t *restrict attrp, void *(*entry)(void *), void *restrict arg)
{
	assert(0);
}

weak_alias(__pthread_exit, pthread_exit);
weak_alias(__pthread_create, pthread_create);
