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
var setBoxTypeCb = make([]func(x, y, w, h int, c uint), 56, 56)

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

	setBoxTypeCb[b] = d

	C.go_fltk_set_boxtype(C.int(b), C.int(o[0]), C.int(o[1]), C.int(o[2]), C.int(o[3]))
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

// Couldn't figure out how to export a func array...
// For now, just gonna hide it at the bottom of the file and pretend it
// doesn't exist
// TODO: If it's possible, fix this
//export _go_drawBox0
func _go_drawBox0(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[0](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox1
func _go_drawBox1(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[1](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox2
func _go_drawBox2(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[2](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox3
func _go_drawBox3(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[3](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox4
func _go_drawBox4(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[4](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox5
func _go_drawBox5(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[5](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox6
func _go_drawBox6(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[6](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox7
func _go_drawBox7(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[7](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox8
func _go_drawBox8(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[8](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox9
func _go_drawBox9(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[9](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox10
func _go_drawBox10(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[10](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox11
func _go_drawBox11(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[11](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox12
func _go_drawBox12(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[12](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox13
func _go_drawBox13(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[13](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox14
func _go_drawBox14(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[14](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox15
func _go_drawBox15(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[15](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox16
func _go_drawBox16(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[16](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox17
func _go_drawBox17(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[17](int(x), int(y), int(w), int(h), uint(c))
}
//export _go_drawBox18
func _go_drawBox18(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[18](int(x), int(y), int(w), int(h), uint(c))
}
