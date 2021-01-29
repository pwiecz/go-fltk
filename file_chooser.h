#pragma once

#ifdef __cplusplus
extern "C" {
#endif

  extern const int go_FL_SINGLE;
  extern const int go_FL_MULTI;
  extern const int go_FL_CREATE;
  extern const int go_FL_DIRECTORY;

  typedef struct Fl_File_Chooser Fl_File_Chooser;

  extern Fl_File_Chooser* go_fltk_new_FileChooser(const char* pathname, const char* pattern, int type, const char* title, void* destroyCallbackId);

  extern void go_fltk_FileChooser_preview(Fl_File_Chooser* fileChooser, int enable);

  extern char* go_fltk_file_chooser(const char* message, const char* pattern, const char* initialFilename, int relative);

#ifdef __cplusplus
}
#endif
