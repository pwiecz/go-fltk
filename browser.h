#pragma once

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

	typedef struct GBrowser GBrowser;
	typedef struct GSelectBrowser GSelectBrowser;
	typedef struct GHoldBrowser GHoldBrowser;
	typedef struct GMultiBrowser GMultiBrowser;

	extern GBrowser   *go_fltk_new_Browser(int x, int y, int w, int h, const char *text);
	extern void        go_fltk_Browser_add(GBrowser *b, const char *str, uintptr_t id);
	extern void        go_fltk_Browser_bottomline(GBrowser *b, int i);
	extern void        go_fltk_Browser_clear(GBrowser *b);
	extern uintptr_t   go_fltk_Browser_data(GBrowser *b, int line);
	extern int         go_fltk_Browser_value(GBrowser *b);
	extern void        go_fltk_Browser_set_value(GBrowser *b, int line);

	extern GSelectBrowser  *go_fltk_new_Select_Browser(int x, int y, int w, int h, const char *text);
	extern GHoldBrowser    *go_fltk_new_Hold_Browser(int x, int y, int w, int h, const char *text);
	extern GMultiBrowser   *go_fltk_new_Multi_Browser(int x, int y, int w, int h, const char *text);

#ifdef __cplusplus
}
#endif
