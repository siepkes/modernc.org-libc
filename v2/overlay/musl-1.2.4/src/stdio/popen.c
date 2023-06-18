#include <fcntl.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <spawn.h>
#include "stdio_impl.h"
#include "syscall.h"
#include <assert.h>

extern char **__environ;

FILE *popen(const char *cmd, const char *mode)
{
	assert(0);
}
