package fltk

/*
#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libfltk.a ${SRCDIR}/lib/linux/x64/libfltk_images.a ${SRCDIR}/lib/linux/x64/libfltk_z.a ${SRCDIR}/lib/linux/x64/libfltk_jpeg.a ${SRCDIR}/lib/linux/x64/libfltk_png.a ${SRCDIR}/lib/linux/x64/libfltk_gl.a -lGLU -lGL -lXrender -lXcursor -lXfixes -lXext -lXft -lfontconfig -lXinerama -lpthread -ldl -lm -lX11
#cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -lfltk -lfltk_images -lfltk_z -lfltk_png -lfltk_jpeg -lfltk_gl -lglu32 -lopengl32 -lgdiplus -lgdi32 -luser32 -lole32 -lcomctl32 -luuid -lws2_32
#cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/libfltk.a ${SRCDIR}/lib/macos/x64/libfltk_images.a ${SRCDIR}/lib/macos/x64/libfltk_z.a ${SRCDIR}/lib/macos/x64/libfltk_jpeg.a ${SRCDIR}/lib/macos/x64/libfltk_png.a ${SRCDIR}/lib/macos/x64/libfltk_gl.a -framework Cocoa -framework OpenGL -framework ApplicationServices -framework Carbon
#include <stdlib.h>
#include "fltk.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// Defined in SetBoxType() as the draw func in custom BoxTypes
var setBoxTypeCb func(x int, y int, w int, h int, c uint)

func Run() int {
	return int(C.go_fltk_run())
}
func Lock() bool {
	return C.go_fltk_lock() == 0
}
func Unlock() {
	C.go_fltk_unlock()
}

func InitStyles() {
	C.go_fltk_init_styles()
}

func SetScheme(scheme string) int {
	schemestr := C.CString(scheme)
	defer C.free(unsafe.Pointer(schemestr))
	return int(C.go_fltk_set_scheme(schemestr))
}

func SetBackgroundColor(r, g, b uint8) {
	C.go_fltk_set_background_color(C.uchar(r), C.uchar(g), C.uchar(b))
}

func SetBackground2Color(r, g, b uint8) {
	C.go_fltk_set_background2_color(C.uchar(r), C.uchar(g), C.uchar(b))
}

func SetBoxType(b BoxType, d func(int, int, int, int, uint), o ...int) {
	if len(o) < 5 {
		o = append(o, []int{0, 0, 0, 0, 0}...)
	}

	setBoxTypeCb = d

	C.go_fltk_set_boxtype(C.int(b), C.int(o[0]), C.int(o[1]), C.int(o[2]), C.int(o[3]))
}

// Is there a better way to do this? Feels a bit hacky.
//export _go_drawBox
func _go_drawBox(x C.int, y C.int, w C.int, h C.int, c C.uint) {
	setBoxTypeCb(int(x), int(y), int(w), int(h), uint(c))
}

func SetForegroundColor(r, g, b uint8) {
	C.go_fltk_set_foreground_color(C.uchar(r), C.uchar(g), C.uchar(b))
}

func SetColor(col Color, r, g, b uint8) {
	C.go_fltk_set_color(C.uint(col), C.uchar(r), C.uchar(g), C.uchar(b))
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

func Awake(fn func()) bool {
	awakeId := globalAwakeMap.register(fn)
	return C.go_fltk_awake(C.uintptr_t(awakeId)) == 0
}
func AwakeNullMessage() {
	C.go_fltk_awake_null_message()
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
func SetKeyboardScreenScaling(value bool) {
	if value {
		C.go_fltk_set_keyboard_screen_scaling(1)
	} else {
		C.go_fltk_set_keyboard_screen_scaling(0)
	}
}
