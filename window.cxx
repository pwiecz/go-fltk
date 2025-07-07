#include "window.h"

#include <FL/Fl_Window.H>
#include <FL/Fl_Double_Window.H>
#include <FL/platform.H>

#include "event_handler.h"


class GWindow : public EventHandler<Fl_Double_Window> {
public:
  GWindow(int w, int h, const char* title)
    : EventHandler<Fl_Double_Window>(w, h, title) {}
  GWindow(int x, int y, int w, int h, const char* title)
    : EventHandler<Fl_Double_Window>(x, y, w, h, title) {}
};

GWindow *go_fltk_new_Window(int w, int h, const char* title) {
  return new GWindow(w, h, title);
}

GWindow *go_fltk_new_Window_with_position(int x, int y, int w, int h, const char* title) {
  return new GWindow(x, y, w, h, title);
}

int go_fltk_Window_shown(Fl_Window *w) {
  return w->shown();
}

void go_fltk_Window_show(Fl_Window *w) {
  w->show();
}

int go_fltk_Window_x_root(Fl_Window* w) {
  return w->x_root();
}

int go_fltk_Window_y_root(Fl_Window* w) {
  return w->y_root();
}

void go_fltk_Window_set_xclass(Fl_Window *w, const char *wmclass) {
  w->xclass(wmclass);
}

void go_fltk_Window_set_label(Fl_Window *w, const char *label) {
  w->copy_label(label);
}

void go_fltk_Window_set_cursor(Fl_Window* w, int cursor) {
  w->cursor((Fl_Cursor)cursor);
}

void go_fltk_Window_set_fullscreen(Fl_Window* w, int flag) {
  if (flag) w->fullscreen(); else w->fullscreen_off();
}

int go_fltk_Window_fullscreen_active(Fl_Window* w) {
  return w->fullscreen_active();
}

void go_fltk_Window_set_modal(Fl_Window* w) {
  w->set_modal();
}

void go_fltk_Window_set_non_modal(Fl_Window* w) {
  w->set_non_modal();
}

void go_fltk_Window_border(Fl_Window *w, int n) {
  w->border(n);
}

void go_fltk_Window_clear_border(Fl_Window *w) {
  w->clear_border();
}

void go_fltk_Window_set_icons(Fl_Window* w, const Fl_RGB_Image *images[], int length) {
  w->icons(images, length);
}

void go_fltk_Window_size_range(Fl_Window* w, int minW, int minH, int maxW, int maxH, int deltaX, int deltaY, int aspectRatio) {
  w->size_range(minW, minH, maxW, maxH, deltaX, deltaY, aspectRatio);
}

void *go_fltk_Window_xid(Fl_Window *w) {
#if defined(_WIN32)
  return fl_win32_xid(w);
#elif defined(__APPLE__)
  return fl_xid(w);
#elif defined(__unix__)
// TODO: Recognize in runtime if we're using Wayland or X11
#if FLTK_USE_X11
  return (void*)(uintptr_t)fl_xid(w);
#elif FLTK_USE_WAYLAND
  return fl_wl_xid(w);
#endif
#endif
}

const int go_FL_CURSOR_DEFAULT = (int)FL_CURSOR_DEFAULT;
const int go_FL_CURSOR_ARROW = (int)FL_CURSOR_ARROW;
const int go_FL_CURSOR_CROSS = (int)FL_CURSOR_CROSS;
const int go_FL_CURSOR_WAIT = (int)FL_CURSOR_WAIT;
const int go_FL_CURSOR_INSERT = (int)FL_CURSOR_INSERT;
const int go_FL_CURSOR_HAND = (int)FL_CURSOR_HAND;
const int go_FL_CURSOR_HELP = (int)FL_CURSOR_HELP;
const int go_FL_CURSOR_MOVE = (int)FL_CURSOR_MOVE;
const int go_FL_CURSOR_NS = (int)FL_CURSOR_NS;
const int go_FL_CURSOR_WE = (int)FL_CURSOR_WE;
const int go_FL_CURSOR_NWSE = (int)FL_CURSOR_NWSE;
const int go_FL_CURSOR_NESW = (int)FL_CURSOR_NESW;
const int go_FL_CURSOR_N = (int)FL_CURSOR_N;
const int go_FL_CURSOR_NE = (int)FL_CURSOR_NE;
const int go_FL_CURSOR_E = (int)FL_CURSOR_E;
const int go_FL_CURSOR_SE = (int)FL_CURSOR_SE;
const int go_FL_CURSOR_S = (int)FL_CURSOR_S;
const int go_FL_CURSOR_SW = (int)FL_CURSOR_SW;
const int go_FL_CURSOR_W = (int)FL_CURSOR_W;
const int go_FL_CURSOR_NW = (int)FL_CURSOR_NW;
const int go_FL_CURSOR_NONE = (int)FL_CURSOR_NONE;
