#include <sys/socket.h>
#include <sys/time.h>
#include <errno.h>
#include "syscall.h"
#include <assert.h>

int getsockopt(int fd, int level, int optname, void *restrict optval, socklen_t *restrict optlen)
{
	assert(0);
}
