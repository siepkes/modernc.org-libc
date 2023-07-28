#include <math.h>
#include <inttypes.h>

long double frexpl(long double x, int *e)
{
	return frexp(x, e);
}
