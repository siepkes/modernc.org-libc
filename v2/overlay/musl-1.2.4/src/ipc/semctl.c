#include <sys/sem.h>
#include <stdarg.h>
#include <endian.h>
#include "syscall.h"
#include "ipc.h"
#include <assert.h>

#if __BYTE_ORDER != __BIG_ENDIAN
#undef SYSCALL_IPC_BROKEN_MODE
#endif

union semun {
	int val;
	struct semid_ds *buf;
	unsigned short *array;
};

int semctl(int id, int num, int cmd, ...)
{
	assert(0);
}
