package fltk

/*
#cgo CPPFLAGS: -I${SRCDIR}/include
#cgo amd64,linux LDFLAGS: ${SRCDIR}/lib/linux/x64/libfltk.a ${SRCDIR}/lib/linux/x64/libfltk_images.a ${SRCDIR}/lib/linux/x64/libfltk_z.a ${SRCDIR}/lib/linux/x64/libfltk_jpeg.a ${SRCDIR}/lib/linux/x64/libfltk_png.a ${SRCDIR}/lib/linux/x64/libfltk_gl.a -lGLU -lGL -lXrender -lXcursor -lXfixes -lXext -lXft -lfontconfig -lXinerama -lpthread -ldl -lm -lX11
#cgo amd64,windows LDFLAGS: -L${SRCDIR}/lib/windows/x64 -lfltk -lfltk_images -lfltk_z -lfltk_png -lfltk_jpeg -lfltk_gl -lglu32 -lopengl32 -lgdiplus -lgdi32 -luser32 -lole32 -lcomctl32 -luuid -lws2_32 -lcomdlg32
#cgo amd64,darwin LDFLAGS: ${SRCDIR}/lib/macos/x64/libfltk.a ${SRCDIR}/lib/macos/x64/libfltk_images.a ${SRCDIR}/lib/macos/x64/libfltk_z.a ${SRCDIR}/lib/macos/x64/libfltk_jpeg.a ${SRCDIR}/lib/macos/x64/libfltk_png.a ${SRCDIR}/lib/macos/x64/libfltk_gl.a -framework Cocoa -framework OpenGL -framework ApplicationServices -framework Carbon
#cgo arm64,darwin LDFLAGS: ${SRCDIR}/lib/macos/arm64/libfltk.a ${SRCDIR}/lib/macos/arm64/libfltk_images.a ${SRCDIR}/lib/macos/arm64/libfltk_z.a ${SRCDIR}/lib/macos/arm64/libfltk_jpeg.a ${SRCDIR}/lib/macos/arm64/libfltk_png.a ${SRCDIR}/lib/macos/arm64/libfltk_gl.a -framework Cocoa -framework OpenGL -framework ApplicationServices -framework Carbon
#include <stdlib.h>
#include "fltk.h"
*/
import "C"
import (
	"image/color"
	"sync"
	"unsafe"
)

// Defined in SetBoxType() as the draw func in custom BoxTypes
var setBoxTypeCb = make([]func(x, y, w, h int, c Color), 57)

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

func SetBackgroundColor(c color.RGBA) {
	C.go_fltk_set_background_color(C.uchar(c.R), C.uchar(c.G), C.uchar(c.B))
}

func SetBackground2Color(c color.RGBA) {
	C.go_fltk_set_background2_color(C.uchar(c.R), C.uchar(c.G), C.uchar(c.B))
}

func SetBoxType(b BoxType, d func(int, int, int, int, Color), o ...int) {
	if len(o) < 5 {
		o = append(o, []int{0, 0, 0, 0, 0}...)
	}

	if b > 56 {
		panic("defining new box types not yet implemented!")
	}

	setBoxTypeCb[b] = d

	C.go_fltk_set_boxtype(C.int(b), C.int(o[0]), C.int(o[1]), C.int(o[2]), C.int(o[3]))
}

func SetForegroundColor(c color.RGBA) {
	C.go_fltk_set_foreground_color(C.uchar(c.R), C.uchar(c.G), C.uchar(c.B))
}

func SetColor(col Color, c color.RGBA) {
	C.go_fltk_set_color(C.uint(col), C.uchar(c.R), C.uchar(c.G), C.uchar(c.B))
}

//SetFont changes a face.
func SetFont(font Font, family string) {
	familystr := C.CString(family)
	defer C.free(unsafe.Pointer(familystr))

	C.go_fltk_set_font(C.int(font), familystr)
}

func SetFont2(font Font, font2 Font) {
	C.go_fltk_set_font2(C.int(font), C.int(font2))
}

func (col Color) Index() uint8 {
	return uint8(col)
}

// Maybe eventually the go-fltk Color type should just be replaced with Go's
// native RGBA?
func (col Color) RGBA() color.RGBA {
	var r, g, b C.uchar
	C.go_fltk_get_color(C.uint(col), &r, &g, &b)

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}

func (col Color) RGBI() uint32 {
	return uint32(C.go_fltk_get_colorindex(C.uint(col)))
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

func Wait(duration ...float64) {
	if len(duration) == 1 {
		C.go_fltk_wait_timed(C.double(duration[0]))
		return
	}

	C.go_fltk_wait()
}

func Check() {
	C.go_fltk_check()
}

type timeoutMap struct {
	mutex      sync.Mutex
	timeoutMap map[uintptr]func()
	id         uintptr
}

func newTimeoutMap() *timeoutMap {
	return &timeoutMap{timeoutMap: make(map[uintptr]func())}
}

var globalTimeoutMap = newTimeoutMap()

func (m *timeoutMap) register(fn func()) uintptr {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.id++
	m.timeoutMap[m.id] = fn
	return m.id
}
func (m *timeoutMap) fetchTimeout(id uintptr) func() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fn := m.timeoutMap[id]
	delete(m.timeoutMap, id)
	return fn
}
func (m *timeoutMap) invoke(id uintptr) {
	fn := m.fetchTimeout(id)
	fn()
}

//export _go_timeoutHandler
func _go_timeoutHandler(id C.uintptr_t) {
	globalTimeoutMap.invoke(uintptr(id))
}

// AddTimeout adds a one-shot timeout callback.  The function will be called by
//  Fl::wait() at t seconds after this function is called.
//  If you need more accurate, repeated timeouts, use RepeatTimeout() to
//  reschedule the subsequent timeouts.
func AddTimeout(t float64, fn func()) {
	timeoutId := globalTimeoutMap.register(fn)
	C.go_fltk_add_timeout(C.double(t), C.uintptr_t(timeoutId))
}

// RepeatTimeout repeats a timeout callback from the expiration of the
//  previous timeout, allowing for more accurate timing.
//  You may only call this method inside a timeout callback of the same timer
//  or at least a closely related timer, otherwise the timing accuracy can't
//  be improved and the behavior is undefined.
func RepeatTimeout(t float64, fn func()) {
	timeoutId := globalTimeoutMap.register(fn)
	C.go_fltk_repeat_timeout(C.double(t), C.uintptr_t(timeoutId))
}

//TODO: implement HasTimeout, RemoveTimeout

func CopyToClipboard(text string) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_copy(textStr, C.int(len(text)), 1 /* destination: clipboard */)
}
func CopyToSelectionBuffer(text string) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_copy(textStr, C.int(len(text)), 0 /* destination: selection buffer */)
}
func DragAndDrop() {
	C.go_fltk_dnd()
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
func ScrollbarSize() int {
	return int(C.go_fltk_scrollbar_size())
}
func SetScrollbarSize(size int) {
	C.go_fltk_set_scrollbar_size(C.int(size))
}

// Couldn't figure out how to export a func array...
// For now, just gonna hide it at the bottom of the file and pretend it
// doesn't exist
// TODO: If it's possible, fix this
//export _go_drawBox0
func _go_drawBox0(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[0](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox1
func _go_drawBox1(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[1](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox2
func _go_drawBox2(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[2](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox3
func _go_drawBox3(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[3](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox4
func _go_drawBox4(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[4](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox5
func _go_drawBox5(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[5](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox6
func _go_drawBox6(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[6](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox7
func _go_drawBox7(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[7](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox8
func _go_drawBox8(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[8](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox9
func _go_drawBox9(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[9](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox10
func _go_drawBox10(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[10](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox11
func _go_drawBox11(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[11](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox12
func _go_drawBox12(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[12](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox13
func _go_drawBox13(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[13](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox14
func _go_drawBox14(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[14](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox15
func _go_drawBox15(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[15](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox16
func _go_drawBox16(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[16](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox17
func _go_drawBox17(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[17](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox18
func _go_drawBox18(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[18](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox19
func _go_drawBox19(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[19](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox20
func _go_drawBox20(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[20](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox21
func _go_drawBox21(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[21](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox22
func _go_drawBox22(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[22](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox23
func _go_drawBox23(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[23](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox24
func _go_drawBox24(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[24](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox25
func _go_drawBox25(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[25](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox26
func _go_drawBox26(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[26](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox27
func _go_drawBox27(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[27](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox28
func _go_drawBox28(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[28](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox29
func _go_drawBox29(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[29](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox30
func _go_drawBox30(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[30](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox31
func _go_drawBox31(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[31](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox32
func _go_drawBox32(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[32](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox33
func _go_drawBox33(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[33](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox34
func _go_drawBox34(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[34](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox35
func _go_drawBox35(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[35](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox36
func _go_drawBox36(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[36](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox37
func _go_drawBox37(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[37](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox38
func _go_drawBox38(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[38](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox39
func _go_drawBox39(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[39](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox40
func _go_drawBox40(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[40](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox41
func _go_drawBox41(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[41](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox42
func _go_drawBox42(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[42](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox43
func _go_drawBox43(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[43](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox44
func _go_drawBox44(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[44](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox45
func _go_drawBox45(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[45](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox46
func _go_drawBox46(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[46](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox47
func _go_drawBox47(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[47](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox48
func _go_drawBox48(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[48](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox49
func _go_drawBox49(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[49](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox50
func _go_drawBox50(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[50](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox51
func _go_drawBox51(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[51](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox52
func _go_drawBox52(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[52](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox53
func _go_drawBox53(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[53](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox54
func _go_drawBox54(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[54](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox55
func _go_drawBox55(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[55](int(x), int(y), int(w), int(h), Color(c))
}

//export _go_drawBox56
func _go_drawBox56(x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[56](int(x), int(y), int(w), int(h), Color(c))
}
