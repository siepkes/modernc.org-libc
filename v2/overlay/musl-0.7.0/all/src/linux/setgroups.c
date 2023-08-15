#include <unistd.h>
#include "syscall.h"

int setgroups(size_t count, const gid_t list[])
{
	return syscall2(__NR_setgroups, count, (long)list);
}
