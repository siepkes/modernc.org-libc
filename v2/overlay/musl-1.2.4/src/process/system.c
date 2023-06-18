#include <unistd.h>
#include <stdlib.h>
#include <signal.h>
#include <sys/wait.h>
#include <spawn.h>
#include <errno.h>
#include "pthread_impl.h"
#include <assert.h>

extern char **__environ;

int system(const char *cmd)
{
	assert(0);
}
