#include <iconv.h>
#include <errno.h>
#include <wchar.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#include <stdint.h>
#include "locale_impl.h"
#include <assert.h>

#define UTF_32BE    0300
#define UTF_16LE    0301
#define UTF_16BE    0302
#define UTF_32LE    0303
#define UCS2BE      0304
#define UCS2LE      0305
#define WCHAR_T     0306
#define US_ASCII    0307
#define UTF_8       0310
#define UTF_16      0312
#define UTF_32      0313
#define UCS2        0314
#define EUC_JP      0320
#define SHIFT_JIS   0321
#define ISO2022_JP  0322
#define GB18030     0330
#define GBK         0331
#define GB2312      0332
#define BIG5        0340
#define EUC_KR      0350


struct stateful_cd {
	iconv_t base_cd;
	unsigned state;
};


iconv_t iconv_open(const char *to, const char *from)
{
	assert(0);
}


/* Adapt as needed */
#define mbrtowc_utf8 mbrtowc
#define wctomb_utf8 wctomb

size_t iconv(iconv_t cd, char **restrict in, size_t *restrict inb, char **restrict out, size_t *restrict outb)
{
	assert(0);
}
