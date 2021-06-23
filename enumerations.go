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

var (
	NO_BOX         = BoxType(C.go_FL_NO_BOX)
	FLAT_BOX       = BoxType(C.go_FL_FLAT_BOX)
	UP_BOX         = BoxType(C.go_FL_UP_BOX)
	DOWN_BOX       = BoxType(C.go_FL_DOWN_BOX)
	UP_FRAME       = BoxType(C.go_FL_UP_FRAME)
	DOWN_FRAME     = BoxType(C.go_FL_DOWN_FRAME)
	THIN_UP_BOX    = BoxType(C.go_FL_THIN_UP_BOX)
	THIN_DOWN_BOX  = BoxType(C.go_FL_THIN_DOWN_BOX)
	ENGRAVED_BOX   = BoxType(C.go_FL_ENGRAVED_BOX)
	EMBOSSED_BOX   = BoxType(C.go_FL_EMBOSSED_BOX)
	ENGRAVED_FRAME = BoxType(C.go_FL_ENGRAVED_FRAME)
	EMBOSSED_FRAME = BoxType(C.go_FL_EMBOSSED_FRAME)
	BORDER_BOX     = BoxType(C.go_FL_BORDER_BOX)
	BORDER_FRAME   = BoxType(C.go_FL_BORDER_FRAME)
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
