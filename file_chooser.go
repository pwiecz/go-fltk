package fltk

/*
#include <stdlib.h>
#include "file_chooser.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type FileChooser struct {
	cPtr                     *C.Fl_File_Chooser
	pathname, pattern, title *C.char
	callbackId               uintptr
}

var ErrFileChooserDestroyed = errors.New("file chooser is destroyed")

type FileChooserType int

var (
	FileChooser_SINGLE    = FileChooserType(C.go_FL_FileChooser_SINGLE)
	FileChooser_MULTI     = FileChooserType(C.go_FL_FileChooser_MULTI)
	FileChooser_CREATE    = FileChooserType(C.go_FL_FileChooser_CREATE)
	FileChooser_DIRECTORY = FileChooserType(C.go_FL_FileChooser_DIRECTORY)
)

func NewFileChooser(pathname, pattern string, fctype FileChooserType, title string) *FileChooser {
	c := &FileChooser{}
	c.pathname = C.CString(pathname)
	c.pattern = C.CString(pattern)
	c.title = C.CString(title)
	c.cPtr = C.go_fltk_new_FileChooser(c.pathname, c.pattern, C.int(fctype), c.title)
	return c
}

func (c *FileChooser) ptr() *C.Fl_File_Chooser {
	if c.cPtr == nil {
		panic(ErrFileChooserDestroyed)
	}
	return c.cPtr
}
func (c *FileChooser) Destroy() {
	if c.pathname != nil {
		C.free(unsafe.Pointer(c.pathname))
	}
	c.pathname = nil
	if c.pattern != nil {
		C.free(unsafe.Pointer(c.pattern))
	}
	c.pattern = nil
	if c.title != nil {
		C.free(unsafe.Pointer(c.title))
	}
	c.title = nil
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	C.go_fltk_FileChooser_destroy(c.ptr())
	c.cPtr = nil
}

func (c *FileChooser) SetCallback(callback func()) {
	if c.callbackId > 0 {
		globalCallbackMap.unregister(c.callbackId)
	}
	c.callbackId = globalCallbackMap.register(callback)
	C.go_fltk_FileChooser_set_callback(c.ptr(), C.uintptr_t(c.callbackId))
}
func (c *FileChooser) Show() {
	C.go_fltk_FileChooser_show((*C.Fl_File_Chooser)(c.ptr()))
}
func (c *FileChooser) Popup() {
	C.go_fltk_FileChooser_popup((*C.Fl_File_Chooser)(c.ptr()))
}
func (c *FileChooser) Shown() bool {
	return C.go_fltk_FileChooser_shown(c.ptr()) != 0
}
func (c *FileChooser) SetPreview(enable bool) {
	if enable {
		C.go_fltk_FileChooser_preview(c.ptr(), 1)
	} else {
		C.go_fltk_FileChooser_preview(c.ptr(), 0)
	}
}
func (c *FileChooser) Selection() []string {
	count := int(C.go_fltk_FileChooser_count(c.ptr()))
	var selection []string
	for i := 1; i <= count; i++ {
		value := C.GoString(C.go_fltk_FileChooser_value(c.ptr(), C.int(i)))
		selection = append(selection, value)
	}
	return selection
}
func (c *FileChooser) SetValue(filename string) {
	filenameStr := C.CString(filename)
	defer C.free(unsafe.Pointer(filenameStr))
	C.go_fltk_FileChooser_set_value(c.ptr(), filenameStr)
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

type NativeFileChooser struct {
	ptr *C.Fl_Native_File_Chooser
}

func NewNativeFileChooser() *NativeFileChooser {
	return &NativeFileChooser{
		C.go_fltk_new_NativeFileChooser(),
	}
}

var ErrNativeFileChooserDestroyed = errors.New("native file chooser is destroyed")

func (c *NativeFileChooser) Destroy() {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	C.go_fltk_NativeFileChooser_destroy(c.ptr)
	c.ptr = nil
}
func (c *NativeFileChooser) Count() int {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return int(C.go_fltk_NativeFileChooser_count(c.ptr))
}

func (c *NativeFileChooser) Directory() string {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return C.GoString(C.go_fltk_NativeFileChooser_directory(c.ptr))
}

func (c *NativeFileChooser) SetDirectory(directory string) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	directoryStr := C.CString(directory)
	defer C.free(unsafe.Pointer(directoryStr))
	C.go_fltk_NativeFileChooser_set_directory(c.ptr, directoryStr)
}

func (c *NativeFileChooser) Filenames() []string {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	count := c.Count()
	var filenames []string
	for i := 0; i < count; i++ {
		filename := C.GoString(C.go_fltk_NativeFileChooser_nth_filename(c.ptr, C.int(i)))
		filenames = append(filenames, filename)
	}
	return filenames
}

func (c *NativeFileChooser) Filter() string {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return C.GoString(C.go_fltk_NativeFileChooser_filter(c.ptr))
}

func (c *NativeFileChooser) SetFilter(filter string) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	filterStr := C.CString(filter)
	defer C.free(unsafe.Pointer(filterStr))
	C.go_fltk_NativeFileChooser_set_filter(c.ptr, filterStr)
}

func (c *NativeFileChooser) FilterValue() int {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return int(C.go_fltk_NativeFileChooser_filter_value(c.ptr))
}

func (c *NativeFileChooser) SetFilterValue(v int) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	C.go_fltk_NativeFileChooser_set_filter_value(c.ptr, C.int(v))
}

func (c *NativeFileChooser) FilterCount() int {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return int(C.go_fltk_NativeFileChooser_filters(c.ptr))
}

func (c *NativeFileChooser) Options() int {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return int(C.go_fltk_NativeFileChooser_options(c.ptr))
}

type NativeFileChooserType int

var (
	// NativeFileChooser options
	NativeFileChooser_NO_OPTIONS     = int(C.go_FL_NativeFileChooser_NO_OPTIONS)
	NativeFileChooser_SAVEAS_CONFIRM = int(C.go_FL_NativeFileChooser_SAVEAS_CONFIRM)
	NativeFileChooser_NEW_FOLDER     = int(C.go_FL_NativeFileChooser_NEW_FOLDER)
	NativeFileChooser_PREVIEW        = int(C.go_FL_NativeFileChooser_PREVIEW)
	NativeFileChooser_USE_FILTER_EXT = int(C.go_FL_NativeFileChooser_USE_FILTER_EXT)

	NativeFileChooser_BROWSE_FILE            = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_FILE)
	NativeFileChooser_BROWSE_DIRECTORY       = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_DIRECTORY)
	NativeFileChooser_BROWSE_MULTI_FILE      = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_MULTI_FILE)
	NativeFileChooser_BROWSE_MULTI_DIRECTORY = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_MULTI_DIRECTORY)
	NativeFileChooser_BROWSE_SAVE_FILE       = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_SAVE_FILE)
	NativeFileChooser_BROWSE_SAVE_DIRECTORY  = NativeFileChooserType(C.go_FL_NativeFileChooser_BROWSE_SAVE_DIRECTORY)
)

func (c *NativeFileChooser) SetOptions(options int) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	C.go_fltk_NativeFileChooser_set_options(c.ptr, C.int(options))
}

func (c *NativeFileChooser) PresetFile() string {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return C.GoString(C.go_fltk_NativeFileChooser_preset_file(c.ptr))
}

func (c *NativeFileChooser) SetPresetFile(presetFile string) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	presetFileStr := C.CString(presetFile)
	defer C.free(unsafe.Pointer(presetFileStr))
	C.go_fltk_NativeFileChooser_set_preset_file(c.ptr, presetFileStr)
}

func (c *NativeFileChooser) Show() int {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return int(C.go_fltk_NativeFileChooser_show(c.ptr))
}

func (c *NativeFileChooser) Title() string {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return C.GoString(C.go_fltk_NativeFileChooser_title(c.ptr))
}

func (c *NativeFileChooser) SetTitle(title string) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	titleStr := C.CString(title)
	defer C.free(unsafe.Pointer(titleStr))
	C.go_fltk_NativeFileChooser_set_title(c.ptr, titleStr)
}

func (c *NativeFileChooser) Type() NativeFileChooserType {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	return NativeFileChooserType(C.go_fltk_NativeFileChooser_type(c.ptr))
}

func (c *NativeFileChooser) SetType(typ NativeFileChooserType) {
	if c.ptr == nil {
		panic(ErrNativeFileChooserDestroyed)
	}
	C.go_fltk_NativeFileChooser_set_type(c.ptr, C.int(typ))
}
