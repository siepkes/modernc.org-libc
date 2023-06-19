#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <limits.h>
#include <assert.h>

extern char **__environ;

int __execvpe(const char *file, char *const argv[], char *const envp[])
{
	assert(0);
}

int execvp(const char *file, char *const argv[])
{
	assert(0);
}

weak_alias(__execvpe, execvpe);
