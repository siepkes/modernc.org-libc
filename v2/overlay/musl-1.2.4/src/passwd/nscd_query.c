#include <sys/socket.h>
#include <byteswap.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <limits.h>
#include "nscd.h"
#include <assert.h>

FILE *__nscd_query(int32_t req, const char *key, int32_t *buf, size_t len, int *swap)
{
	assert(0);
}
