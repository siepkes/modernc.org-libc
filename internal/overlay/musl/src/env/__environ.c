#include <unistd.h>

char **__environ = 0;
#ifndef __CCGO__
weak_alias(__environ, ___environ);
weak_alias(__environ, _environ);
#endif
weak_alias(__environ, environ);
