#include <signal.h>
#include <errno.h>
#include <string.h>
#include "syscall.h"
#include "pthread_impl.h"
#include "libc.h"
#include "lock.h"
#include "ksigaction.h"
#include <assert.h>

void __get_handler_set(sigset_t *set)
{
	assert(0);
}

volatile int __eintr_valid_flag;

int __libc_sigaction(int sig, const struct sigaction *restrict sa, struct sigaction *restrict old)
{
	assert(0);
}

int __sigaction(int sig, const struct sigaction *restrict sa, struct sigaction *restrict old)
{
	assert(0);
}

weak_alias(__sigaction, sigaction);
