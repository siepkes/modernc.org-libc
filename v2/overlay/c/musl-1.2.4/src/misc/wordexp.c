#include <wordexp.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <limits.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <signal.h>
#include <errno.h>
#include <fcntl.h>
#include "pthread_impl.h"
#include <assert.h>

int wordexp(const char *restrict s, wordexp_t *restrict we, int flags)
{
	assert(0);
}

void wordfree(wordexp_t *we)
{
	assert(0);
}
