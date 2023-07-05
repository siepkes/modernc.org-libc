#define _BSD_SOURCE
#include <stdlib.h>
#include <sys/mman.h>

#include "meta.h"
#include <assert.h>

void free(void *p)
{
	assert(0);
}
