#include <stdio.h>
#include "getc.h"

int getc(FILE *f)
{
	return do_getc(f);
}

#ifndef __CCGO__
weak_alias(getc, _IO_getc);
#endif
