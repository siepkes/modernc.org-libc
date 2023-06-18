#define _GNU_SOURCE
#include <net/if.h>
#include <errno.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include "netlink.h"
#include <assert.h>

#define IFADDRS_HASH_SIZE 64

struct ifnamemap {
	unsigned int hash_next;
	unsigned int index;
	unsigned char namelen;
	char name[IFNAMSIZ];
};

struct ifnameindexctx {
	unsigned int num, allocated, str_bytes;
	struct ifnamemap *list;
	unsigned int hash[IFADDRS_HASH_SIZE];
};


struct if_nameindex *if_nameindex()
{
	assert(0);
}
