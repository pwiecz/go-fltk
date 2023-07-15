#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct Fl_Widget Fl_Widget;

  extern const unsigned int go_FL_ALIGN_CENTER;
  extern const unsigned int go_FL_ALIGN_TOP;
  extern const unsigned int go_FL_ALIGN_BOTTOM;
  extern const unsigned int go_FL_ALIGN_LEFT;
  extern const unsigned int go_FL_ALIGN_RIGHT;
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
	
  extern const int go_FL_ARROW_SINGLE;
  extern const int go_FL_ARROW_DOUBLE;
  extern const int go_FL_ARROW_CHOICE;
  extern const int go_FL_ARROW_RETURN;
	
  extern const int go_FL_ORIENT_NONE;
  extern const int go_FL_ORIENT_RIGHT;
  extern const int go_FL_ORIENT_NE;
  extern const int go_FL_ORIENT_UP;
  extern const int go_FL_ORIENT_NW;
  extern const int go_FL_ORIENT_LEFT;
  extern const int go_FL_ORIENT_SW;
  extern const int go_FL_ORIENT_DOWN;
  extern const int go_FL_ORIENT_SE;
	
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
  extern const int go_FL_ENTER_KEY;
  extern const int go_FL_LEAVE;
  extern const int go_FL_FOCUS;
  extern const int go_FL_UNFOCUS;
  extern const int go_FL_KEYBOARD;
  extern const int go_FL_KEYDOWN;
  extern const int go_FL_KEYUP;
  extern const int go_FL_CLOSE;
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

  extern const int go_FL_WHEN_NEVER;
  extern const int go_FL_WHEN_CHANGED;
  extern const int go_FL_WHEN_NOT_CHANGED;
  extern const int go_FL_WHEN_RELEASE;
  extern const int go_FL_WHEN_RELEASE_ALWAYS;
  extern const int go_FL_WHEN_ENTER_KEY;
  extern const int go_FL_WHEN_ENTER_KEY_ALWAYS;
  extern const int go_FL_WHEN_ENTER_KEY_CHANGED;

  extern const int go_FL_ESCAPE;
  extern const int go_FL_TAB;
  extern const int go_FL_ENTER;
  extern const int go_FL_HOME;
  extern const int go_FL_LEFT;
  extern const int go_FL_UP;
  extern const int go_FL_RIGHT;
  extern const int go_FL_DOWN;
  extern const int go_FL_PAGE_UP;
  extern const int go_FL_PAGE_DOWN;
  extern const int go_FL_END;
  extern const int go_FL_MENU;
  extern const int go_FL_HELP;
  extern const int go_FL_F1;
  extern const int go_FL_F2;
  extern const int go_FL_F3;
  extern const int go_FL_F4;
  extern const int go_FL_F5;
  extern const int go_FL_F6;
  extern const int go_FL_F7;
  extern const int go_FL_F8;
  extern const int go_FL_F9;
  extern const int go_FL_F10;
  extern const int go_FL_F11;
  extern const int go_FL_F12;
  extern const int go_FL_DELETE;
  extern const int go_FL_BACKSPACE;
  extern const int go_FL_INSERT;  

  extern const int go_FL_RGB;
  extern const int go_FL_INDEX;
  extern const int go_FL_SINGLE;
  extern const int go_FL_DOUBLE;
  extern const int go_FL_ACCUM;
  extern const int go_FL_ALPHA;
  extern const int go_FL_DEPTH;
  extern const int go_FL_STENCIL;
  extern const int go_FL_RGB8;
  extern const int go_FL_MULTISAMPLE;
  extern const int go_FL_STEREO;
  extern const int go_FL_FAKE_SINGLE;
  extern const int go_FL_OPENGL3;

  extern const unsigned int go_FL_FOREGROUND_COLOR;
  extern const unsigned int go_FL_BACKGROUND2_COLOR;        
  extern const unsigned int go_FL_INACTIVE_COLOR;
  extern const unsigned int go_FL_SELECTION_COLOR;
  extern const unsigned int go_FL_GRAY0;
  extern const unsigned int go_FL_DARK3;
  extern const unsigned int go_FL_DARK2;
  extern const unsigned int go_FL_DARK1;
  extern const unsigned int go_FL_BACKGROUND_COLOR;
  extern const unsigned int go_FL_LIGHT1;
  extern const unsigned int go_FL_LIGHT2;
  extern const unsigned int go_FL_LIGHT3;
  extern const unsigned int go_FL_BLACK;
  extern const unsigned int go_FL_RED;
  extern const unsigned int go_FL_GREEN;
  extern const unsigned int go_FL_YELLOW;
  extern const unsigned int go_FL_BLUE;
  extern const unsigned int go_FL_MAGENTA;
  extern const unsigned int go_FL_CYAN;
  extern const unsigned int go_FL_DARK_RED;
  extern const unsigned int go_FL_DARK_GREEN;
  extern const unsigned int go_FL_DARK_YELLOW;
  extern const unsigned int go_FL_DARK_BLUE;
  extern const unsigned int go_FL_DARK_MAGENTA;
  extern const unsigned int go_FL_DARK_CYAN;
  extern const unsigned int go_FL_WHITE;

#ifdef __cplusplus
}
#endif
