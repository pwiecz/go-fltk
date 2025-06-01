package fltk

/*
#include <stdlib.h>
#include "fltk.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// Defined in SetBoxType() as the draw func in custom BoxTypes
var setBoxTypeCb = make([]func(x, y, w, h int, c Color), 256)

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

//export _go_drawBox
func _go_drawBox(i, x, y, w, h C.int, c C.uint) {
	setBoxTypeCb[i](int(x), int(y), int(w), int(h), Color(c))
}

func SetBoxType(b BoxType, d func(int, int, int, int, Color), o ...int) {
	if len(o) < 5 {
		o = append(o, []int{0, 0, 0, 0, 0}...)
	}

	if b > 256 {
		panic("reached the max box type limit!")
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

func GetFont(font Font) string {
	return C.GoString(C.go_fltk_get_font(C.int(font)))
}

// GetFontName Returs human readable font name and a font attribute (BOLD, ITALIC or BOLD_ITALIC).
func GetFontName(font Font) (string, Font) {
	var attribute C.int
	fontName := C.go_fltk_get_font_name(C.int(font), &attribute)
	return C.GoString(fontName), Font(attribute)
}

// SetFont assigns font specified by the name to the font number
func SetFont(fontNumber Font, fontName string) {
	fontNameStr := C.CString(fontName)
	// fontNameStr is not freed as the fltk just depends on it living on
	// throughout the lifetime of the app.

	C.go_fltk_set_font(C.int(fontNumber), fontNameStr)
}

func SetFont2(font Font, font2 Font) {
	C.go_fltk_set_font2(C.int(font), C.int(font2))
}

func SetFonts(xstarname ...string) int {
	return int(C.go_fltk_set_fonts(cStringOpt(xstarname)))
}

func (col Color) Index() uint {
	return uint(col)
}

func (col Color) RGB() (int, int, int) {
	var r, g, b C.uchar
	C.go_fltk_get_color(C.uint(col), &r, &g, &b)
	return int(r), int(g), int(b)
}

func (col Color) RGBI() uint {
	return uint(C.go_fltk_get_colorindex(C.uint(col)))
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
//
//	Fl::wait() at t seconds after this function is called.
//	If you need more accurate, repeated timeouts, use RepeatTimeout() to
//	reschedule the subsequent timeouts.
func AddTimeout(t float64, fn func()) {
	timeoutId := globalTimeoutMap.register(fn)
	C.go_fltk_add_timeout(C.double(t), C.uintptr_t(timeoutId))
}

// RepeatTimeout repeats a timeout callback from the expiration of the
//
//	previous timeout, allowing for more accurate timing.
//	You may only call this method inside a timeout callback of the same timer
//	or at least a closely related timer, otherwise the timing accuracy can't
//	be improved and the behavior is undefined.
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

// ScreenNum gets the screen number of a screen that contains the specified screen position x, y.
func ScreenNum(x, y int) int {
	return int(C.go_fltk_screen_num(C.int(x), C.int(y)))
}

// ScreenWorkArea gets the bounding box of the work area of the given screen.
func ScreenWorkArea(screenNum int) (int, int, int, int) {
	var x, y, w, h C.int
	C.go_fltk_screen_work_area(&x, &y, &w, &h, C.int(screenNum))
	return int(x), int(y), int(w), int(h)
}

// ScreenDPI gets the screen resolution in dots-per-inch for the given screen.
func ScreenDPI(screenNum int) (float32, float32) {
	var w, h C.float
	C.go_fltk_screen_dpi(&w, &h, C.int(screenNum))
	return float32(w), float32(h)
}

// ScreenCount gets the total count of available screens.
func ScreenCount() int {
	return int(C.go_fltk_screen_count())
}

// ScreenScale gets current value of the GUI scaling factor for the given screen.
func ScreenScale(screenNum int) float32 {
	return float32(C.go_fltk_screen_scale(C.int(screenNum)))
}

// SetScreenScale sets the value of the GUI scaling factor for the given screen.
func SetScreenScale(screenNum int, scale float32) {
	C.go_fltk_set_screen_scale(C.int(screenNum), C.float(scale))
}

// SetKeyboardScreenScaling controls the possibility to scale all windows by ctrl/+/-/0/ or cmd/+/-/0/
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
func MenuLinespacing() int {
	return int(C.go_fltk_menu_linespacing())
}
func SetMenuLinespacing(size int) {
	C.go_fltk_set_menu_linespacing(C.int(size))
}
func TestShortcut(shortcut int) bool {
	return C.go_fltk_test_shortcut(C.int(shortcut)) != 0
}

func Version() string {
	return "v1.4.0"
}

func GoVersion() string {
	return "v0.1.0"
}
