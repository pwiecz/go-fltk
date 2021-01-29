package fltk

/*
#include <stdlib.h>
#include "file_chooser.h"
*/
import "C"
import "unsafe"

type FileChooser struct {
	Widget
	pathname, pattern, title *C.char
	destroyCallbackId        uintptr
}

type FileChooserType int

var (
	SINGLE    = FileChooserType(C.go_FL_SINGLE)
	MULTI     = FileChooserType(C.go_FL_MULTI)
	CREATE    = FileChooserType(C.go_FL_CREATE)
	DIRECTORY = FileChooserType(C.go_FL_DIRECTORY)
)

func NewFileChooser(pathname, pattern string, fctype FileChooserType, title string) *FileChooser {
	c := &FileChooser{}
	c.pathname = C.CString(pathname)
	c.pattern = C.CString(pattern)
	c.title = C.CString(title)
	c.destroyCallbackId = globalCallbackMap.register(func() { c.destroy() })
	initWidget(c, unsafe.Pointer(C.go_fltk_new_FileChooser(c.pathname, c.pattern, C.int(fctype), c.title, unsafe.Pointer(c.callbackId))))
	return c
}
func (c *FileChooser) destroy() {
	C.free(unsafe.Pointer(c.pathname))
	C.free(unsafe.Pointer(c.pattern))
	C.free(unsafe.Pointer(c.title))
	globalCallbackMap.unregister(c.destroyCallbackId)
}
func (c *FileChooser) SetPreview(enable bool) {
	if enable {
		C.go_fltk_FileChooser_preview((*C.Fl_File_Chooser)(c.ptr), C.int(1))
	} else {
		C.go_fltk_FileChooser_preview((*C.Fl_File_Chooser)(c.ptr), C.int(0))
	}
}

func ChooseFile(message, pattern, initialFilename string, relative bool) (string, bool) {
	var rel int
	if relative {
		rel = 1
	}
	messageStr := C.CString(message)
	defer C.free(unsafe.Pointer(messageStr))
	patternStr := C.CString(pattern)
	defer C.free(unsafe.Pointer(patternStr))
	initialFilenameStr := C.CString(initialFilename)
	defer C.free(unsafe.Pointer(initialFilenameStr))
	res := C.go_fltk_file_chooser(messageStr, patternStr, initialFilenameStr, C.int(rel))
	if res == nil {
		return "", false
	}
	return C.GoString(res), true
}
