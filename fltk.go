package fltk

/*
#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libfltk_gl.a -lGLU -lGL ${SRCDIR}/lib/linux/x64/libfltk.a -lXrender -lXcursor -lXfixes -lXext -lXft -lfontconfig -lXinerama -lpthread -ldl -lm -lX11 
#cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -lfltk_gl -lglu32 -lopengl32 -lfltk -lgdiplus -lgdi32 -luser32 -lole32 -lcomctl32 -luuid -lws2_32
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
func _go_awakeHandler(id C.uintptr_t) {
	globalAwakeMap.invoke(uintptr(id))
}

func Awake(fn func()) {
	awakeId := globalAwakeMap.register(fn)
	C.go_fltk_awake(C.uintptr_t(awakeId))
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
func ScreenCount() int {
	return int(C.go_fltk_screen_count())
}
func ScreenScale(screenNum int) float32 {
	return float32(C.go_fltk_screen_scale(C.int(screenNum)))
}
func SetScreenScale(screenNum int, scale float32) {
	C.go_fltk_set_screen_scale(C.int(screenNum), C.float(scale))
}
