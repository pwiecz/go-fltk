package fltk

/*
#include "enumerations.h"
*/
import "C"

type Align uint

var (
	ALIGN_CENTER             = Align(C.go_FL_ALIGN_CENTER)
	ALIGN_TOP                = Align(C.go_FL_ALIGN_TOP)
	ALIGN_BOTTOM             = Align(C.go_FL_ALIGN_BOTTOM)
	ALIGN_LEFT               = Align(C.go_FL_ALIGN_LEFT)
	ALIGN_RIGHT              = Align(C.go_FL_ALIGN_RIGHT)
	ALIGN_INSIDE             = Align(C.go_FL_ALIGN_INSIDE)
	ALIGN_TEXT_OVER_IMAGE    = Align(C.go_FL_ALIGN_TEXT_OVER_IMAGE)
	ALIGN_IMAGE_OVER_TEXT    = Align(C.go_FL_ALIGN_IMAGE_OVER_TEXT)
	ALIGN_CLIP               = Align(C.go_FL_ALIGN_CLIP)
	ALIGN_WRAP               = Align(C.go_FL_ALIGN_WRAP)
	ALIGN_IMAGE_NEXT_TO_TEXT = Align(C.go_FL_ALIGN_IMAGE_NEXT_TO_TEXT)
	ALIGN_TEXT_NEXT_TO_IMAGE = Align(C.go_FL_ALIGN_TEXT_NEXT_TO_IMAGE)
	ALIGN_IMAGE_BACKDROP     = Align(C.go_FL_ALIGN_IMAGE_BACKDROP)
	ALIGN_TOP_LEFT           = Align(C.go_FL_ALIGN_TOP_LEFT)
	ALIGN_TOP_RIGHT          = Align(C.go_FL_ALIGN_TOP_RIGHT)
	ALIGN_BOTTOM_LEFT        = Align(C.go_FL_ALIGN_BOTTOM_LEFT)
	ALIGN_BOTTOM_RIGHT       = Align(C.go_FL_ALIGN_BOTTOM_RIGHT)
	ALIGN_LEFT_TOP           = Align(C.go_FL_ALIGN_LEFT_TOP)
	ALIGN_RIGHT_TOP          = Align(C.go_FL_ALIGN_RIGHT_TOP)
	ALIGN_LEFT_BOTTOM        = Align(C.go_FL_ALIGN_LEFT_BOTTOM)
	ALIGN_RIGHT_BOTTOM       = Align(C.go_FL_ALIGN_RIGHT_BOTTOM)
	ALIGN_NOWRAP             = Align(C.go_FL_ALIGN_NOWRAP)
	ALIGN_POSITION_MASK      = Align(C.go_FL_ALIGN_POSITION_MASK)
	ALIGN_IMAGE_MASK         = Align(C.go_FL_ALIGN_IMAGE_MASK)
)

type BoxType int

const (
	NO_BOX                 = BoxType(0)
	FLAT_BOX               = BoxType(1)
	UP_BOX                 = BoxType(2)
	DOWN_BOX               = BoxType(3)
	UP_FRAME               = BoxType(4)
	DOWN_FRAME             = BoxType(5)
	THIN_UP_BOX            = BoxType(6)
	THIN_DOWN_BOX          = BoxType(7)
	THIN_UP_FRAME          = BoxType(8)
	THIN_DOWN_FRAME        = BoxType(9)
	ENGRAVED_BOX           = BoxType(10)
	EMBOSSED_BOX           = BoxType(11)
	ENGRAVED_FRAME         = BoxType(12)
	EMBOSSED_FRAME         = BoxType(13)
	BORDER_BOX             = BoxType(14)
	SHADOW_BOX             = BoxType(15)
	BORDER_FRAME           = BoxType(16)
	SHADOW_FRAME           = BoxType(17)
	ROUNDED_BOX            = BoxType(18)
	RSHADOW_BOX            = BoxType(19)
	ROUNDED_FRAME          = BoxType(20)
	RFLAT_BOX              = BoxType(21)
	ROUND_UP_BOX           = BoxType(22)
	ROUND_DOWN_BOX         = BoxType(23)
	DIAMOND_UP_BOX         = BoxType(24)
	DIAMOND_DOWN_BOX       = BoxType(25)
	OVAL_BOX               = BoxType(26)
	OSHADOW_BOX            = BoxType(27)
	OVAL_FRAME             = BoxType(28)
	OFLAT_BOX              = BoxType(29)
	PLASTIC_UP_BOX         = BoxType(30)
	PLASTIC_DOWN_BOX       = BoxType(31)
	PLASTIC_UP_FRAME       = BoxType(32)
	PLASTIC_DOWN_FRAME     = BoxType(33)
	PLASTIC_THIN_UP_BOX    = BoxType(34)
	PLASTIC_THIN_DOWN_BOX  = BoxType(35)
	PLASTIC_ROUND_UP_BOX   = BoxType(36)
	PLASTIC_ROUND_DOWN_BOX = BoxType(37)
	GTK_UP_BOX             = BoxType(38)
	GTK_DOWN_BOX           = BoxType(39)
	GTK_UP_FRAME           = BoxType(40)
	GTK_DOWN_FRAME         = BoxType(41)
	GTK_THIN_UP_BOX        = BoxType(42)
	GTK_THIN_DOWN_BOX      = BoxType(43)
	GTK_THIN_UP_FRAME      = BoxType(44)
	GTK_THIN_DOWN_FRAME    = BoxType(45)
	GTK_ROUND_UP_FRAME     = BoxType(46)
	GTK_ROUND_DOWN_FRAME   = BoxType(47)
	GLEAM_UP_BOX           = BoxType(48)
	GLEAM_DOWN_BOX         = BoxType(49)
	GLEAM_UP_FRAME         = BoxType(50)
	GLEAM_DOWN_FRAME       = BoxType(51)
	GLEAM_THIN_UP_BOX      = BoxType(52)
	GLEAM_THIN_DOWN_BOX    = BoxType(53)
	GLEAM_ROUND_UP_BOX     = BoxType(54)
	GLEAM_ROUND_DOWN_BOX   = BoxType(55)
	FREE_BOXTYPE           = BoxType(56)
)

type Font int

var (
	HELVETICA             = Font(C.go_FL_HELVETICA)
	HELVETICA_BOLD        = Font(C.go_FL_HELVETICA_BOLD)
	HELVETICA_ITALIC      = Font(C.go_FL_HELVETICA_ITALIC)
	HELVETICA_BOLD_ITALIC = Font(C.go_FL_HELVETICA_BOLD_ITALIC)
	COURIER               = Font(C.go_FL_COURIER)
	COURIER_BOLD          = Font(C.go_FL_COURIER_BOLD)
	COURIER_ITALIC        = Font(C.go_FL_COURIER_ITALIC)
	COURIER_BOLD_ITALIC   = Font(C.go_FL_COURIER_BOLD_ITALIC)
	TIMES                 = Font(C.go_FL_TIMES)
	TIMES_BOLD            = Font(C.go_FL_TIMES_BOLD)
	TIMES_ITALIC          = Font(C.go_FL_TIMES_ITALIC)
	TIMES_BOLD_ITALIC     = Font(C.go_FL_TIMES_BOLD_ITALIC)
	SYMBOL                = Font(C.go_FL_SYMBOL)
	SCREEN                = Font(C.go_FL_SCREEN)
	SCREEN_BOLD           = Font(C.go_FL_SCREEN_BOLD)
	ZAPF_DINGBATS         = Font(C.go_FL_ZAPF_DINGBATS)
	FREE_FONT             = Font(C.go_FL_FREE_FONT)
	BOLD                  = Font(C.go_FL_BOLD)
	ITALIC                = Font(C.go_FL_ITALIC)
	BOLD_ITALIC           = Font(C.go_FL_BOLD_ITALIC)
)

type LabelType int

var (
	NORMAL_LABEL = LabelType(C.go_FL_NORMAL_LABEL)
	NO_LABEL     = LabelType(C.go_FL_NO_LABEL)
)

type WrapMode int

const (
	WRAP_NONE      = WrapMode(0)
	WRAP_AT_COLUMN = WrapMode(1)
	WRAP_AT_PIXEL  = WrapMode(2)
	WRAP_AT_BOUNDS = WrapMode(3)
)

type Event int

var (
	NO_EVENT       = Event(C.go_FL_NO_EVENT)
	PUSH           = Event(C.go_FL_PUSH)
	DRAG           = Event(C.go_FL_DRAG)
	RELEASE        = Event(C.go_FL_RELEASE)
	MOVE           = Event(C.go_FL_MOVE)
	MOUSEWHEEL     = Event(C.go_FL_MOUSEWHEEL)
	ENTER          = Event(C.go_FL_ENTER)
	LEAVE          = Event(C.go_FL_LEAVE)
	FOCUS          = Event(C.go_FL_FOCUS)
	UNFOCUS        = Event(C.go_FL_UNFOCUS)
	KEY            = Event(C.go_FL_KEYDOWN)
	KEYDOWN        = Event(C.go_FL_KEYDOWN)
	KEYUP          = Event(C.go_FL_KEYUP)
	SHORTCUT       = Event(C.go_FL_SHORTCUT)
	DEACTIVATE     = Event(C.go_FL_DEACTIVATE)
	ACTIVATE       = Event(C.go_FL_ACTIVATE)
	HIDE           = Event(C.go_FL_HIDE)
	SHOW           = Event(C.go_FL_SHOW)
	PASTE          = Event(C.go_FL_PASTE)
	SELECTIONCLEAR = Event(C.go_FL_SELECTIONCLEAR)
	DND_ENTER      = Event(C.go_FL_DND_ENTER)
	DND_DRAG       = Event(C.go_FL_DND_DRAG)
	DND_LEAVE      = Event(C.go_FL_DND_LEAVE)
	DND_RELEASE    = Event(C.go_FL_DND_RELEASE)
)

type CallbackCondition int

var (
	WhenNever           = CallbackCondition(C.go_FL_WHEN_NEVER)
	WhenChanged         = CallbackCondition(C.go_FL_WHEN_CHANGED)
	WhenNotChanged      = CallbackCondition(C.go_FL_WHEN_NOT_CHANGED)
	WhenRelease         = CallbackCondition(C.go_FL_WHEN_RELEASE)
	WhenReleaseAlways   = CallbackCondition(C.go_FL_WHEN_RELEASE_ALWAYS)
	WhenEnterKey        = CallbackCondition(C.go_FL_WHEN_ENTER_KEY)
	WhenEnterKeyAlways  = CallbackCondition(C.go_FL_WHEN_ENTER_KEY_ALWAYS)
	WhenEnterKeyChanged = CallbackCondition(C.go_FL_WHEN_ENTER_KEY_CHANGED)
)

var (
	RGB         = int(C.go_FL_RGB)
	INDEX       = int(C.go_FL_INDEX)
	SINGLE      = int(C.go_FL_SINGLE)
	DOUBLE      = int(C.go_FL_DOUBLE)
	ACCUM       = int(C.go_FL_ACCUM)
	ALPHA       = int(C.go_FL_ALPHA)
	DEPTH       = int(C.go_FL_DEPTH)
	STENCIL     = int(C.go_FL_STENCIL)
	RGB8        = int(C.go_FL_RGB8)
	MULTISAMPLE = int(C.go_FL_MULTISAMPLE)
	STEREO      = int(C.go_FL_STEREO)
	FAKE_SINGLE = int(C.go_FL_FAKE_SINGLE)
	OPENGL3     = int(C.go_FL_OPENGL3)
)

type Color uint

var (
	FOREGROUND_COLOR  = Color(C.go_FL_FOREGROUND_COLOR)
	BACKGROUND2_COLOR = Color(C.go_FL_BACKGROUND2_COLOR)
	INACTIVE_COLOR    = Color(C.go_FL_INACTIVE_COLOR)
	SELECTION_COLOR   = Color(C.go_FL_SELECTION_COLOR)
	GRAY0             = Color(C.go_FL_GRAY0)
	DARK3             = Color(C.go_FL_DARK3)
	DARK2             = Color(C.go_FL_DARK2)
	DARK1             = Color(C.go_FL_DARK1)
	BACKGROUND_COLOR  = Color(C.go_FL_BACKGROUND_COLOR)
	LIGHT1            = Color(C.go_FL_LIGHT1)
	LIGHT2            = Color(C.go_FL_LIGHT2)
	LIGHT3            = Color(C.go_FL_LIGHT3)
	BLACK             = Color(C.go_FL_BLACK)
	RED               = Color(C.go_FL_RED)
	GREEN             = Color(C.go_FL_GREEN)
	YELLOW            = Color(C.go_FL_YELLOW)
	BLUE              = Color(C.go_FL_BLUE)
	MAGENTA           = Color(C.go_FL_MAGENTA)
	CYAN              = Color(C.go_FL_CYAN)
	DARK_RED          = Color(C.go_FL_DARK_RED)
	DARK_GREEN        = Color(C.go_FL_DARK_GREEN)
	DARK_YELLOW       = Color(C.go_FL_DARK_YELLOW)
	DARK_BLUE         = Color(C.go_FL_DARK_BLUE)
	DARK_MAGENTA      = Color(C.go_FL_DARK_MAGENTA)
	DARK_CYAN         = Color(C.go_FL_DARK_CYAN)
	WHITE             = Color(C.go_FL_WHITE)
)

func ColorFromRgb(r, g, b uint8) Color {
	r1 := uint(r)
	g1 := uint(g)
	b1 := uint(b)
	return Color(((r1 & 0xff) << 24) + ((g1 & 0xff) << 16) + ((b1 & 0xff) << 8) + 0x00)
}

type LineStyle int

var (
	SOLID        = LineStyle(0)
	DASH         = LineStyle(1)
	Dot          = LineStyle(2)
	DASH_DOT     = LineStyle(3)
	DASH_DOT_DOT = LineStyle(4)
	CAP_FLAT     = LineStyle(100)
	CAP_ROUND    = LineStyle(200)
	CAP_SQUARE   = LineStyle(300)
	JOIN_MITER   = LineStyle(1000)
	JOIN_ROUND   = LineStyle(2000)
	JOIN_BEVEL   = LineStyle(3000)
)

type callbackMap struct {
	callbackMap map[uintptr]func()
	id          uintptr
}

func newCallbackMap() *callbackMap {
	return &callbackMap{
		callbackMap: make(map[uintptr]func()),
	}
}
func (m *callbackMap) register(fn func()) uintptr {
	m.id++
	m.callbackMap[m.id] = fn
	return m.id
}
func (m *callbackMap) unregister(id uintptr) {
	delete(m.callbackMap, id)
}
func (m *callbackMap) invoke(id uintptr) {
	if callback, ok := m.callbackMap[id]; ok && callback != nil {
		callback()
	}
}
func (m *callbackMap) isEmpty() bool {
	return len(m.callbackMap) == 0
}
func (m *callbackMap) size() int {
	return len(m.callbackMap)
}
func (m *callbackMap) clear() {
	for id := range m.callbackMap {
		delete(m.callbackMap, id)
	}
}

var globalCallbackMap = newCallbackMap()

//export _go_callbackHandler
func _go_callbackHandler(id uintptr) {
	globalCallbackMap.invoke(id)
}

type eventHandlerMap struct {
	eventHandlerMap map[int]func(Event) bool
	id              int
}

func newEventHandlerMap() *eventHandlerMap {
	return &eventHandlerMap{
		eventHandlerMap: make(map[int]func(Event) bool),
	}
}
func (m *eventHandlerMap) register(fn func(Event) bool) int {
	m.id++
	m.eventHandlerMap[m.id] = fn
	return m.id
}
func (m *eventHandlerMap) unregister(id int) {
	delete(m.eventHandlerMap, id)
}
func (m *eventHandlerMap) invoke(id int, event Event) bool {
	if handler, ok := m.eventHandlerMap[id]; ok && handler != nil {
		return handler(event)
	}
	return false
}
func (m *eventHandlerMap) isEmpty() bool {
	return len(m.eventHandlerMap) == 0
}
func (m *eventHandlerMap) size() int {
	return len(m.eventHandlerMap)
}
func (m *eventHandlerMap) clear() {
	for id := range m.eventHandlerMap {
		delete(m.eventHandlerMap, id)
	}
}

var globalEventHandlerMap = newEventHandlerMap()

//export _go_eventHandler
func _go_eventHandler(handlerId C.int, event C.int) C.int {
	if globalEventHandlerMap.invoke(int(handlerId), Event(event)) {
		return 1
	}
	return 0
}

var (
	ESCAPE = int(C.go_FL_ESCAPE)
)
