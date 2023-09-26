#ifndef _OBSTACK_H
#define _OBSTACK_H 1

#include <stddef.h>

struct obstack {
	int TODO;
};

extern int obstack_printf (struct obstack *obstack, char *format, ...);
extern int obstack_vprintf (struct obstack *obstack, char *format, __builtin_va_list va);

int obstack_begin (struct obstack *obstack_ptr, size_t chunk_size);
int obstack_init (struct obstack *obstack_ptr);
int obstack_specify_allocation (struct obstack *obstack_ptr, size_t chunk_size, size_t alignment, void *(*chunkfun) (size_t), void (*freefun) (void *));
int obstack_specify_allocation_with_arg (struct obstack *obstack_ptr, size_t chunk_size, size_t alignment, void *(*chunkfun) (void *, size_t), void (*freefun) (void *, void *), void *arg);
size_t obstack_alignment_mask (struct obstack *obstack_ptr);
size_t obstack_chunk_size (struct obstack *obstack_ptr);
size_t obstack_object_size (struct obstack *obstack_ptr);
size_t obstack_room (struct obstack *obstack_ptr);
void *obstack_alloc (struct obstack *obstack_ptr, size_t size);
void *obstack_base (struct obstack *obstack_ptr);
void *obstack_copy (struct obstack *obstack_ptr, void *address, size_t size);
void *obstack_copy0 (struct obstack *obstack_ptr, void *address, size_t size);
void *obstack_finish (struct obstack *obstack_ptr);
void *obstack_next_free (struct obstack *obstack_ptr);
void obstack_1grow (struct obstack *obstack_ptr, char data_char);
void obstack_1grow_fast (struct obstack *obstack_ptr, char data_char);
void obstack_blank (struct obstack *obstack_ptr, size_t size);
void obstack_blank_fast (struct obstack *obstack_ptr, size_t size);
void obstack_free (struct obstack *obstack_ptr, void *object);
void obstack_grow (struct obstack *obstack_ptr, void *address, size_t size);
void obstack_grow0 (struct obstack *obstack_ptr, void *address, size_t size);

#endif
