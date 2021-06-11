package fltk

/*
#include <stdlib.h>
#include "file_chooser.h"
*/
import "C"
import "unsafe"

type FileChooser struct {
	ptr                      *C.Fl_File_Chooser
	pathname, pattern, title *C.char
	callbackId               uintptr
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
	c.ptr = C.go_fltk_new_FileChooser(c.pathname, c.pattern, C.int(fctype), c.title)
	return c
}

func (c *FileChooser) Destroy() {
	C.free(unsafe.Pointer(c.pathname))
	C.free(unsafe.Pointer(c.pattern))
	C.free(unsafe.Pointer(c.title))
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	C.go_fltk_FileChooser_destroy(c.ptr)
}

func (c *FileChooser) SetCallback(callback func()) {
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	c.callbackId = globalCallbackMap.register(callback)
	C.go_fltk_FileChooser_set_callback(c.ptr, unsafe.Pointer(c.callbackId))
}
func (c *FileChooser) Show() {
	C.go_fltk_FileChooser_show((*C.Fl_File_Chooser)(c.ptr))
}
func (c *FileChooser) Popup() {
	C.go_fltk_FileChooser_popup((*C.Fl_File_Chooser)(c.ptr))
}
func (c *FileChooser) Shown() bool {
	return C.go_fltk_FileChooser_shown(c.ptr) != C.int(0)
}
func (c *FileChooser) SetPreview(enable bool) {
	if enable {
		C.go_fltk_FileChooser_preview(c.ptr, C.int(1))
	} else {
		C.go_fltk_FileChooser_preview(c.ptr, C.int(0))
	}
}
func (c *FileChooser) Selection() []string {
	count := int(C.go_fltk_FileChooser_count(c.ptr))
	var selection []string
	for i := 1; i <= count; i++ {
		value := C.GoString(C.go_fltk_FileChooser_value(c.ptr, C.int(i)))
		selection = append(selection, value)
	}
	return selection
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
