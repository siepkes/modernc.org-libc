#include "stdio_impl.h"
#include <errno.h>
#include <limits.h>
#include <string.h>
#include <stdlib.h>
#include "libc.h"
#include <assert.h>

struct cookie {
	char **bufp;
	size_t *sizep;
	size_t pos;
	char *buf;
	size_t len;
	size_t space;
};

struct ms_FILE {
	FILE f;
	struct cookie c;
	unsigned char buf[BUFSIZ];
};

FILE *open_memstream(char **bufp, size_t *sizep)
{
	assert(0);
}
