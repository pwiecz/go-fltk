#pragma once

#ifdef __cplusplus
extern "C" {
#endif

	typedef struct GHelp_View GHelp_View;

	extern GHelp_View  *go_fltk_new_HelpView(int x, int y, int w, int h, const char *text);
	extern void         go_fltk_HelpView_load(GHelp_View *h, const char *f);
	extern const char  *go_fltk_HelpView_directory(GHelp_View *h);
	extern const char  *go_fltk_HelpView_filename(GHelp_View *h);
	extern int          go_fltk_HelpView_find(GHelp_View *h, const char *s, int p);

#ifdef __cplusplus
}
#endif
