#pragma once

#include <stdint.h>

#ifdef __cplusplus

extern "C" {
#endif

        typedef struct GBrowser GBrowser;
        typedef struct Fl_Browser Fl_Browser;
	typedef struct GSelectBrowser GSelectBrowser;
	typedef struct GHoldBrowser GHoldBrowser;
        typedef struct GMultiBrowser GMultiBrowser;
        typedef struct GCheckBrowser GCheckBrowser;
        typedef struct Fl_Check_Browser Fl_Check_Browser;
	typedef struct Fl_Image Fl_Image;

	extern GBrowser   *go_fltk_new_Browser(int x, int y, int w, int h, const char *text);
	extern void        go_fltk_Browser_add(Fl_Browser* b, const char *str, uintptr_t id);
	extern int         go_fltk_Browser_topline(Fl_Browser* b);

	extern void        go_fltk_Browser_set_topline(Fl_Browser* b, int i);
	extern void        go_fltk_Browser_set_middleline(Fl_Browser* b, int i);
	extern void        go_fltk_Browser_set_bottomline(Fl_Browser* b, int i);

	extern void        go_fltk_Browser_clear(Fl_Browser* b);
	extern uintptr_t   go_fltk_Browser_data(Fl_Browser* b, int line);

	extern void        go_fltk_Browser_remove(Fl_Browser* b, int i);
	extern char        go_fltk_Browser_column_char(Fl_Browser* b);
	extern void        go_fltk_Browser_set_column_char(Fl_Browser* b, char c);
	extern void        go_fltk_Browser_hide_line(Fl_Browser* b, int line);
	extern int         go_fltk_Browser_size(Fl_Browser* b);
	extern Fl_Image*   go_fltk_Browser_icon(Fl_Browser* b, int line);
	extern void        go_fltk_Browser_set_icon(Fl_Browser* b, int line, Fl_Image *i);
	extern char        go_fltk_Browser_format_char(Fl_Browser* b);
	extern void        go_fltk_Browser_set_format_char(Fl_Browser* b, char c);
	extern int         go_fltk_Browser_displayed(Fl_Browser* b, int line);
	extern int         go_fltk_Browser_value(Fl_Browser* b);
	extern void        go_fltk_Browser_set_value(Fl_Browser* b, int line);
	extern const char* go_fltk_Browser_text(Fl_Browser* b, int line);
	extern void        go_fltk_Browser_set_column_widths(Fl_Browser* b, const int *arr);
	extern int         go_fltk_Browser_select(Fl_Browser* b, int line, int val);
        extern int         go_fltk_Browser_selected(Fl_Browser *b, int line);

	extern GSelectBrowser* go_fltk_new_Select_Browser(int x, int y, int w, int h, const char *text);
	extern GHoldBrowser*   go_fltk_new_Hold_Browser(int x, int y, int w, int h, const char *text);
        extern GMultiBrowser*  go_fltk_new_Multi_Browser(int x, int y, int w, int h, const char *text);
        extern GCheckBrowser*  go_fltk_new_Check_Browser(int x, int y, int w, int h, const char *text);
        extern void         go_fltk_Check_Browser_add(Fl_Check_Browser *b, const char *s, int checked);
        extern void         go_fltk_Check_Browser_set_checked(Fl_Check_Browser* b, int item, int checked);
        extern int          go_fltk_Check_Browser_is_checked(Fl_Check_Browser *b,int item);
        extern int          go_fltk_Check_Browser_nchecked(Fl_Check_Browser *b);
        extern void         go_fltk_Check_Browser_remove(Fl_Check_Browser* b, int item);
        extern void         go_fltk_Check_Browser_clear(Fl_Check_Browser *b);
        extern int          go_fltk_Check_Browser_nitems(Fl_Check_Browser* b);        
        extern const char*  go_fltk_Check_Browser_text(Fl_Check_Browser* b, int item);        
#ifdef __cplusplus
}
#endif
