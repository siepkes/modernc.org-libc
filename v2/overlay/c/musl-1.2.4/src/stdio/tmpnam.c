#include <stdio.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/stat.h>
#include <string.h>
#include <stdlib.h>
#include "syscall.h"
#include <assert.h>

#define MAXTRIES 100

char *tmpnam(char *buf)
{
	assert(0);
}
