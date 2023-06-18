#define _BSD_SOURCE
#include <glob.h>
#include <fnmatch.h>
#include <sys/stat.h>
#include <dirent.h>
#include <limits.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <stddef.h>
#include <unistd.h>
#include <pwd.h>
#include <assert.h>

struct match
{
	struct match *next;
	char name[];
};

int glob(const char *restrict pat, int flags, int (*errfunc)(const char *path, int err), glob_t *restrict g)
{
	assert(0);
}

void globfree(glob_t *g)
{
	assert(0);
}
