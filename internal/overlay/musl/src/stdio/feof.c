#include "stdio_impl.h"

#undef feof

int feof(FILE *f)
{
	FLOCK(f);
	int ret = !!(f->flags & F_EOF);
	FUNLOCK(f);
	return ret;
}

weak_alias(feof, feof_unlocked);
#ifndef __CCGO__
weak_alias(feof, _IO_feof_unlocked);
#endif
