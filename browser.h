#pragma once

#ifdef __cplusplus

extern "C" {
#endif

	typedef struct GBrowser GBrowser;
	typedef struct Fl_Image Fl_Image;

	extern GBrowser*   go_fltk_new_Browser(int x, int y, int w, int h, const char *text);
	extern void        go_fltk_Browser_add(GBrowser *b, const char *str, void *d);
	extern int         go_fltk_Browser_topline(GBrowser *b);

	extern void        go_fltk_Browser_set_topline(GBrowser *b, int i);
	extern void        go_fltk_Browser_set_middleline(GBrowser *b, int i);
	extern void        go_fltk_Browser_set_bottomline(GBrowser *b, int i);

	extern void        go_fltk_Browser_clear(GBrowser *b);
	extern void        go_fltk_Browser_remove(GBrowser *b, int i);
	extern char        go_fltk_Browser_column_char(GBrowser *b);
	extern void        go_fltk_Browser_set_column_char(GBrowser *b, char c);
	extern void        go_fltk_Browser_hide_line(GBrowser *b, int line);
	extern int         go_fltk_Browser_size(GBrowser *b);
	extern Fl_Image*   go_fltk_Browser_icon(GBrowser *b, int line);
	extern void        go_fltk_Browser_set_icon(GBrowser *b, int line, Fl_Image *i);
	extern char        go_fltk_Browser_format_char(GBrowser *b);
	extern void        go_fltk_Browser_set_format_char(GBrowser *b, char c);
	extern int         go_fltk_Browser_displayed(GBrowser *b, int line);
	extern int         go_fltk_Browser_value(GBrowser *b);
	extern const char* go_fltk_Browser_text(GBrowser *b, int line);
	extern void        go_fltk_Browser_set_column_widths(GBrowser *b, const int *arr);

#ifdef __cplusplus
}
#endif
