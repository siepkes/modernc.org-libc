#include <libintl.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <limits.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <ctype.h>
#include "locale_impl.h"
#include "atomic.h"
#include "pleval.h"
#include "lock.h"
#include "fork_impl.h"
#include <assert.h>

#define malloc __libc_malloc
#define calloc __libc_calloc
#define realloc undef
#define free undef

struct binding {
	struct binding *next;
	int dirlen;
	volatile int active;
	char *domainname;
	char *dirname;
	char buf[];
};

static volatile int lock[1];
volatile int *const __gettext_lockptr = lock;

char *bindtextdomain(const char *domainname, const char *dirname)
{
	assert(0);
}

struct msgcat {
	struct msgcat *next;
	const void *map;
	size_t map_size;
	const char *plural_rule;
	int nplurals;
	struct binding *binding;
	const struct __locale_map *lm;
	int cat;
};

static char *dummy_gettextdomain()
{
	assert(0);
}

weak_alias(dummy_gettextdomain, __gettextdomain);

char *dcngettext(const char *domainname, const char *msgid1, const char *msgid2, unsigned long int n, int category)
{
	assert(0);
}

char *dcgettext(const char *domainname, const char *msgid, int category)
{
	assert(0);
}

char *dngettext(const char *domainname, const char *msgid1, const char *msgid2, unsigned long int n)
{
	assert(0);
}

char *dgettext(const char *domainname, const char *msgid)
{
	assert(0);
}
