#include "file_chooser.h"

#include <Fl/Fl_File_Chooser.H>

#include "_cgo_export.h"


const int go_FL_SINGLE = Fl_File_Chooser::SINGLE;
const int go_FL_MULTI = Fl_File_Chooser::MULTI;
const int go_FL_CREATE = Fl_File_Chooser::CREATE;
const int go_FL_DIRECTORY = Fl_File_Chooser::DIRECTORY;


class GFileChooser : public Fl_File_Chooser {
public:
  GFileChooser(const char* message, const char* pattern, int type, const char* title, void* destroyCallbackId)
    : Fl_File_Chooser(message, pattern, type, title)
    , m_destroyCallbackId(destroyCallbackId) {}

  ~GFileChooser() {
    fprintf(stderr, "destroying filechooser");
    _go_callbackHandler(m_destroyCallbackId);
  }

private:
  void * const m_destroyCallbackId;
};

Fl_File_Chooser *go_fltk_new_FileChooser(const char* message, const char* pattern, int type, const char* title, void* destroyCallbackId) {
  return new GFileChooser(message, pattern, type, title, destroyCallbackId);
}

void go_fltk_FileChooser_preview(Fl_File_Chooser* fileChooser, int enable) {
  fileChooser->preview(enable);
}

char* go_fltk_file_chooser(const char* message, const char* pattern, const char* initialFilename, int relative) {
  return fl_file_chooser(message, pattern, initialFilename, relative);
}

