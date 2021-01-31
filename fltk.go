package fltk

/*
#cgo linux LDFLAGS: /usr/lib/x86_64-linux-gnu/libfltk_gl.a -lGLU -lGL /usr/lib/x86_64-linux-gnu/libfltk.a -lXrender -lXcursor -lXfixes -lXext -lXft -lfontconfig -lXinerama -lpthread -ldl -lm -lX11

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
	C.go_fltk_copy(C.CString(text), C.int(len(text)), 1 /* destination: clipboard */)
}
