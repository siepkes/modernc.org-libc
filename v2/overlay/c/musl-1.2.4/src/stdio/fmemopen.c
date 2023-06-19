#include "stdio_impl.h"
#include <errno.h>
#include <string.h>
#include <stdlib.h>
#include <stddef.h>
#include <inttypes.h>
#include "libc.h"
#include <assert.h>

struct cookie {
	size_t pos, len, size;
	unsigned char *buf;
	int mode;
};

struct mem_FILE {
	FILE f;
	struct cookie c;
	unsigned char buf[UNGET+BUFSIZ], buf2[];
};

FILE *fmemopen(void *restrict buf, size_t size, const char *restrict mode)
{
	assert(0);
}
