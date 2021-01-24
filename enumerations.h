#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Widget Fl_Widget;

  extern const int go_FL_NO_BOX;
  extern const int go_FL_FLAT_BOX;
  extern const int go_FL_UP_BOX;
  extern const int go_FL_DOWN_BOX;
  extern const int go_FL_UP_FRAME;
  extern const int go_FL_DOWN_FRAME;
  extern const int go_FL_THIN_UP_BOX;
  extern const int go_FL_THIN_DOWN_BOX;
  extern const int go_FL_ENGRAVED_BOX;
  extern const int go_FL_EMBOSSED_BOX;
  extern const int go_FL_ENGRAVED_FRAME;
  extern const int go_FL_EMBOSSED_FRAME;
  extern const int go_FL_BORDER_BOX;
  extern const int go_FL_BORDER_FRAME;

  extern const unsigned int go_FL_ALIGN_CENTER;
  extern const unsigned int go_FL_ALIGN_TOP;
  extern const unsigned int go_FL_ALIGN_BOTTOM;
  extern const unsigned int go_FL_ALIGN_LEFT;
  extern const unsigned int go_FL_ALIGN_INSIDE;
  extern const unsigned int go_FL_ALIGN_TEXT_OVER_IMAGE;
  extern const unsigned int go_FL_ALIGN_IMAGE_OVER_TEXT;
  extern const unsigned int go_FL_ALIGN_CLIP;
  extern const unsigned int go_FL_ALIGN_WRAP;
  extern const unsigned int go_FL_ALIGN_IMAGE_NEXT_TO_TEXT;
  extern const unsigned int go_FL_ALIGN_TEXT_NEXT_TO_IMAGE;
  extern const unsigned int go_FL_ALIGN_IMAGE_BACKDROP;
  extern const unsigned int go_FL_ALIGN_TOP_LEFT;
  extern const unsigned int go_FL_ALIGN_TOP_RIGHT;
  extern const unsigned int go_FL_ALIGN_BOTTOM_LEFT;
  extern const unsigned int go_FL_ALIGN_BOTTOM_RIGHT;
  extern const unsigned int go_FL_ALIGN_LEFT_TOP;
  extern const unsigned int go_FL_ALIGN_RIGHT_TOP;
  extern const unsigned int go_FL_ALIGN_LEFT_BOTTOM;
  extern const unsigned int go_FL_ALIGN_RIGHT_BOTTOM;
  extern const unsigned int go_FL_ALIGN_NOWRAP;
  extern const unsigned int go_FL_ALIGN_POSITION_MASK;
  extern const unsigned int go_FL_ALIGN_IMAGE_MASK;

  extern const int go_FL_HELVETICA;
  extern const int go_FL_HELVETICA_BOLD;
  extern const int go_FL_HELVETICA_ITALIC;
  extern const int go_FL_HELVETICA_BOLD_ITALIC;
  extern const int go_FL_COURIER;
  extern const int go_FL_COURIER_BOLD;
  extern const int go_FL_COURIER_ITALIC;
  extern const int go_FL_COURIER_BOLD_ITALIC;
  extern const int go_FL_TIMES;
  extern const int go_FL_TIMES_BOLD;
  extern const int go_FL_TIMES_ITALIC;
  extern const int go_FL_TIMES_BOLD_ITALIC;
  extern const int go_FL_SYMBOL;
  extern const int go_FL_SCREEN;
  extern const int go_FL_SCREEN_BOLD;
  extern const int go_FL_ZAPF_DINGBATS;
  extern const int go_FL_FREE_FONT;
  extern const int go_FL_BOLD;
  extern const int go_FL_ITALIC;
  extern const int go_FL_BOLD_ITALIC;

  extern const int go_FL_NORMAL_LABEL;
  extern const int go_FL_NO_LABEL;

  extern const int go_FL_NO_EVENT;
  extern const int go_FL_PUSH;
  extern const int go_FL_DRAG;
  extern const int go_FL_RELEASE;
  extern const int go_FL_MOVE;
  extern const int go_FL_MOUSEWHEEL;
  extern const int go_FL_ENTER;
  extern const int go_FL_LEAVE;
  extern const int go_FL_FOCUS;
  extern const int go_FL_UNFOCUS;
  extern const int go_FL_KEYBOARD;
  extern const int go_FL_KEYDOWN;
  extern const int go_FL_KEYUP;
  extern const int go_FL_SHORTCUT;
  extern const int go_FL_DEACTIVATE;
  extern const int go_FL_ACTIVATE;
  extern const int go_FL_HIDE;
  extern const int go_FL_SHOW;
  extern const int go_FL_PASTE;
  extern const int go_FL_SELECTIONCLEAR;
  extern const int go_FL_DND_ENTER;
  extern const int go_FL_DND_DRAG;
  extern const int go_FL_DND_LEAVE;
  extern const int go_FL_DND_RELEASE;

  extern const int go_FL_SHIFT;
  extern const int go_FL_CAPS_LOCK;
  extern const int go_FL_CTRL;
  extern const int go_FL_ALT;
  extern const int go_FL_NUM_LOCK;
  extern const int go_FL_META;
  extern const int go_FL_SCROLL_LOCK;
  extern const int go_FL_BUTTON1;
  extern const int go_FL_BUTTON2;
  extern const int go_FL_BUTTON3;

  extern void callback_handler(Fl_Widget *w, void* data);

#ifdef __cplusplus
}
#endif
