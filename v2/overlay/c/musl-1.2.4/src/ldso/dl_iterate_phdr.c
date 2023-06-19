#include <elf.h>
#include <link.h>
#include "pthread_impl.h"
#include "libc.h"
#include <assert.h>

#define AUX_CNT 38

extern weak hidden const size_t _DYNAMIC[];

static int static_dl_iterate_phdr(int(*callback)(struct dl_phdr_info *info, size_t size, void *data), void *data)
{
	assert(0); //TODO
}

weak_alias(static_dl_iterate_phdr, dl_iterate_phdr);
