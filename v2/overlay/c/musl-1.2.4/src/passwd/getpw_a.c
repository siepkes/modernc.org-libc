#include <pthread.h>
#include <byteswap.h>
#include <string.h>
#include <unistd.h>
#include "pwf.h"
#include "nscd.h"
#include <assert.h>

int __getpw_a(const char *name, uid_t uid, struct passwd *pw, char **buf, size_t *size, struct passwd **res)
{
	assert(0);
}
