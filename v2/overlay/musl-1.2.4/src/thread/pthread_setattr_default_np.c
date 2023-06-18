#define _GNU_SOURCE
#include "pthread_impl.h"
#include <string.h>
#include <assert.h>

#define MIN(a,b) ((a)<(b) ? (a) : (b))
#define MAX(a,b) ((a)>(b) ? (a) : (b))

int pthread_setattr_default_np(const pthread_attr_t *attrp)
{
	assert(0);
}

int pthread_getattr_default_np(pthread_attr_t *attrp)
{
	assert(0);
}
