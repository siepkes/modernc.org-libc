#define _GNU_SOURCE
#define SYSCALL_NO_TLS 1
#include <stdlib.h>
#include <stdarg.h>
#include <stddef.h>
#include <string.h>
#include <unistd.h>
#include <stdint.h>
#include <elf.h>
#include <sys/mman.h>
#include <limits.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <errno.h>
#include <link.h>
#include <setjmp.h>
#include <pthread.h>
#include <ctype.h>
#include <dlfcn.h>
#include <semaphore.h>
#include <sys/membarrier.h>
#include "pthread_impl.h"
#include "fork_impl.h"
#include "dynlink.h"

#ifndef PAGE_SIZE
#define PAGE_SIZE ldso_page_size
#endif

#include "libc.h"
#include <assert.h>

#define malloc __libc_malloc
#define calloc __libc_calloc
#define realloc __libc_realloc
#define free __libc_free

#define MAXP2(a,b) (-(-(a)&-(b)))
#define ALIGN(x,y) ((x)+(y)-1 & -(y))

#define container_of(p,t,m) ((t*)((char *)(p)-offsetof(t,m)))
#define countof(a) ((sizeof (a))/(sizeof (a)[0]))

struct debug {
	int ver;
	void *head;
	void (*bp)(void);
	int state;
	void *base;
};

struct td_index {
	size_t args[2];
	struct td_index *next;
};

struct dso {
#if DL_FDPIC
	struct fdpic_loadmap *loadmap;
#else
	unsigned char *base;
#endif
	char *name;
	size_t *dynv;
	struct dso *next, *prev;

	Phdr *phdr;
	int phnum;
	size_t phentsize;
	Sym *syms;
	Elf_Symndx *hashtab;
	uint32_t *ghashtab;
	int16_t *versym;
	char *strings;
	struct dso *syms_next, *lazy_next;
	size_t *lazy, lazy_cnt;
	unsigned char *map;
	size_t map_len;
	dev_t dev;
	ino_t ino;
	char relocated;
	char constructed;
	char kernel_mapped;
	char mark;
	char bfs_built;
	char runtime_loaded;
	struct dso **deps, *needed_by;
	size_t ndeps_direct;
	size_t next_dep;
	pthread_t ctor_visitor;
	char *rpath_orig, *rpath;
	struct tls_module tls;
	size_t tls_id;
	size_t relro_start, relro_end;
	uintptr_t *new_dtv;
	unsigned char *new_tls;
	struct td_index *td_index;
	struct dso *fini_next;
	char *shortname;
#if DL_FDPIC
	unsigned char *base;
#else
	struct fdpic_loadmap *loadmap;
#endif
	struct funcdesc {
		void *addr;
		size_t *got;
	} *funcdescs;
	size_t *got;
	char buf[];
};

struct symdef {
	Sym *sym;
	struct dso *dso;
};

typedef void (*stage3_func)(size_t *, size_t *);

#define MIN_TLS_ALIGN offsetof(struct builtin_tls, pt)

#define ADDEND_LIMIT 4096

static struct debug debug;
struct debug *_dl_debug_addr = &debug;

extern hidden int __malloc_replaced;

hidden void (*const __init_array_start)(void)=0, (*const __fini_array_start)(void)=0;

extern hidden void (*const __init_array_end)(void), (*const __fini_array_end)(void);

weak_alias(__init_array_start, __init_array_end);
weak_alias(__fini_array_start, __fini_array_end);

#define strcmp(l,r) dl_strcmp(l,r)

#if DL_FDPIC
static void *laddr(const struct dso *p, size_t v)
{
	assert(0);
}
static void *laddr_pg(const struct dso *p, size_t v)
{
	assert(0);
}

static void (*fdbarrier(void *p))()
{
	assert(0);
}
#define fpaddr(p, v) fdbarrier((&(struct funcdesc){ \
	laddr(p, v), (p)->got }))
#else
#define laddr(p, v) (void *)((p)->base + (v))
#define laddr_pg(p, v) laddr(p, v)
#define fpaddr(p, v) ((void (*)())laddr(p, v))
#endif

#define OK_TYPES (1<<STT_NOTYPE | 1<<STT_OBJECT | 1<<STT_FUNC | 1<<STT_COMMON | 1<<STT_TLS)
#define OK_BINDS (1<<STB_GLOBAL | 1<<STB_WEAK | 1<<STB_GNU_UNIQUE)

#ifndef ARCH_SYM_REJECT_UND
#define ARCH_SYM_REJECT_UND(s) 0
#endif

#if defined(__GNUC__)
__attribute__((always_inline))
#endif
static inline struct symdef find_sym2(struct dso *dso, const char *s, int need_def, int use_deps)
{
	assert(0);
}

void __libc_exit_fini()
{
	assert(0);
}

void __ldso_atfork(int who)
{
	assert(0);
}

void __libc_start_init(void)
{
	assert(0);
}

static void dl_debug_state(void)
{
}

weak_alias(dl_debug_state, _dl_debug_state);

void __init_tls(size_t *auxv)
{
}

hidden void __dls2(unsigned char *base, size_t *sp)
{
	assert(0);
}

void __dls2b(size_t *sp, size_t *auxv)
{
	assert(0);
}

void __dls3(size_t *sp, size_t *auxv)
{
	assert(0);
}

void *dlopen(const char *file, int mode)
{
	assert(0);
}

hidden int __dl_invalid_handle(void *h)
{
	assert(0);
}

int dladdr(const void *addr_arg, Dl_info *info)
{
	assert(0);
}

hidden void *__dlsym(void *restrict p, const char *restrict s, void *restrict ra)
{
	assert(0);
}

hidden void *__dlsym_redir_time64(void *restrict p, const char *restrict s, void *restrict ra)
{
	assert(0);
}

int dl_iterate_phdr(int(*callback)(struct dl_phdr_info *info, size_t size, void *data), void *data)
{
	assert(0);
}
