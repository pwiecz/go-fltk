#pragma once

#ifdef __cplusplus
extern "C" {
#endif

	typedef struct GBrowser GBrowser;

	extern GBrowser   *go_fltk_new_Browser(int x, int y, int w, int h, const char *text);
	extern void        go_fltk_Browser_add(GBrowser *b, const char *str, void *d);
	extern void        go_fltk_Browser_bottomline(GBrowser *b, int i);
	extern void        go_fltk_Browser_clear(GBrowser *b);

#ifdef __cplusplus
}
#endif
