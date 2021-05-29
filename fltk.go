package fltk

/*
#cgo amd64,linux LDFLAGS: /usr/lib/x86_64-linux-gnu/libfltk_gl.a -lGLU -lGL /usr/lib/x86_64-linux-gnu/libfltk.a -lXrender -lXcursor -lXfixes -lXext -lXft -lfontconfig -lXinerama -lpthread -ldl -lm -lX11
#cgo amd64,windows LDFLAGS: -L. -lfltk_forms -lfltk_images -lfltk_gl -lfltk_png -lfltk_z -lfltk_jpeg -lfltk -lgdi32 -lglu32 -lopengl32 -lole32 -lcomctl32 -luuid
#include <stdlib.h>
#include "fltk.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

func Run() int {
	return int(C.go_fltk_run())
}
func Lock() int {
	return int(C.go_fltk_lock())
}

func cStringOpt(s []string) *C.char {
	if len(s) == 0 {
		return (*C.char)(nil)
	}
	return C.CString(s[0])
}

type awakeMap struct {
	mutex    sync.Mutex
	awakeMap map[uintptr]func()
	id       uintptr
}

func newAwakeMap() *awakeMap {
	return &awakeMap{awakeMap: make(map[uintptr]func())}
}

var globalAwakeMap = newAwakeMap()

func (m *awakeMap) register(fn func()) uintptr {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.id++
	m.awakeMap[m.id] = fn
	return m.id
}
func (m *awakeMap) fetchCallback(id uintptr) func() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fn := m.awakeMap[id]
	delete(m.awakeMap, id)
	return fn
}
func (m *awakeMap) invoke(id uintptr) {
	fn := m.fetchCallback(id)
	fn()
}

//export _go_awakeHandler
func _go_awakeHandler(id unsafe.Pointer) {
	globalAwakeMap.invoke(uintptr(id))
}

func Awake(fn func()) {
	awakeId := globalAwakeMap.register(fn)
	C.go_fltk_awake(unsafe.Pointer(awakeId))
}

func CopyToClipboard(text string) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_copy(textStr, C.int(len(text)), 1 /* destination: clipboard */)
}
func ScreenWorkArea(screenNum int) (int, int, int, int) {
	var x, y, w, h C.int
	C.go_fltk_screen_work_area(&x, &y, &w, &h, C.int(screenNum))
	return int(x), int(y), int(w), int(h)
}
func ScreenDPI(screenNum int) (float32, float32) {
	var w, h C.float
	C.go_fltk_screen_dpi(&w, &h, C.int(screenNum))
	return float32(w), float32(h)
}
