package fltk

/*
#include "file_chooser.h"
*/
import "C"
import "unsafe"

type FileChooser struct {
	Widget
}

type FileChooserType int
var (
	SINGLE = FileChooserType(C.go_FL_SINGLE)
	MULTI = FileChooserType(C.go_FL_MULTI)
	CREATE = FileChooserType(C.go_FL_CREATE)
	DIRECTORY = FileChooserType(C.go_FL_DIRECTORY)
)

func NewFileChooser(pathname, pattern string, fctype FileChooserType, title string) *FileChooser {
	c := &FileChooser{}
	initWidget(c, unsafe.Pointer(C.go_fltk_new_FileChooser(C.CString(pathname), C.CString(pattern), C.int(fctype), C.CString(title))))
	return c
}
func (c* FileChooser) SetPreview(enable bool) {
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
	res := C.go_fltk_file_chooser(C.CString(message), C.CString(pattern), C.CString(initialFilename), C.int(rel))
	if res == nil {
		return "", false
	}
	return C.GoString(res), true
}

