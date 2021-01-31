package fltk

/*
#include <stdlib.h>
#include "dialogs.h"
*/
import "C"
import "unsafe"

func MessageBox(title, message string) {
	titleStr := C.CString(title)
	defer C.free(unsafe.Pointer(titleStr))
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	C.go_fltk_message_box(titleStr, messageStr)
}
