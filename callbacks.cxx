#include "callbacks.h"

#include <cstdint>

#include "_cgo_export.h"


void callback_handler(Fl_Widget *w, void* data) {
  _go_callbackHandler((uintptr_t)data);
}
