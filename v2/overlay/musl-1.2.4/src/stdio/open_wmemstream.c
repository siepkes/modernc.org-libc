#include "stdio_impl.h"
#include <wchar.h>
#include <errno.h>
#include <limits.h>
#include <string.h>
#include <stdlib.h>
#include "libc.h"
#include <assert.h>

struct cookie {
	wchar_t **bufp;
	size_t *sizep;
	size_t pos;
	wchar_t *buf;
	size_t len;
	size_t space;
	mbstate_t mbs;
};

struct wms_FILE {
	FILE f;
	struct cookie c;
	unsigned char buf[1];
};

FILE *open_wmemstream(wchar_t **bufp, size_t *sizep)
{
	assert(0);
}
