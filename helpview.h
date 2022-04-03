#pragma once

#ifdef __cplusplus
extern "C" {
#endif

	typedef struct GHelp_View GHelp_View;

	extern Fl_Help_View *go_fltk_new_HelpView(int x, int y, int w, int h, const char *text);
	extern void          go_fltk_HelpView_load(GHelp_View *h, const char *f);
	extern const char   *go_fltk_HelpView_directory(GHelp_View *h);
	extern const char   *go_fltk_HelpView_filename(GHelp_View *h);
	extern int           go_fltk_HelpView_find(GHelp_View *h, const char *s, int p);
	extern int           go_fltk_HelpView_leftline(GHelp_View *h);
	extern void          go_fltk_HelpView_set_leftline(GHelp_View *h, int i);
	extern int           go_fltk_HelpView_topline(GHelp_View *h);
	extern void          go_fltk_HelpView_set_topline(GHelp_View *h, int i);
	extern void          go_fltk_HelpView_set_toplinestring(GHelp_View *h, const char *s);

#ifdef __cplusplus
}
#endif
