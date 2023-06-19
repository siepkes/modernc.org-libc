#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <ctype.h>
#include <wchar.h>
#include <wctype.h>
#include <limits.h>
#include <string.h>

#include "stdio_impl.h"
#include "shgetc.h"
#include "intscan.h"
#include "floatscan.h"
#include <assert.h>

#define SIZE_hh -2
#define SIZE_h  -1
#define SIZE_def 0
#define SIZE_l   1
#define SIZE_L   2
#define SIZE_ll  3

#if 1
#undef getwc
#define getwc(f) \
	((f)->rpos != (f)->rend && *(f)->rpos < 128 ? *(f)->rpos++ : (getwc)(f))

#undef ungetwc
#define ungetwc(c,f) \
	((f)->rend && (c)<128U ? *--(f)->rpos : ungetwc((c),(f)))
#endif

int vfwscanf(FILE *restrict f, const wchar_t *restrict fmt, va_list ap)
{
	assert(0);
}

weak_alias(vfwscanf,__isoc99_vfwscanf);
