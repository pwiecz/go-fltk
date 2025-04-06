package fltk

/*
#include <stdlib.h>
#include "callbacks.h"
#include "widget.h"
*/
import "C"
import "fmt"

// Parameterless callbacks
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

// Event handlers
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

// Draw handlers
type drawHandlerMap struct {
	drawHandlerMap map[uintptr]func(func())
	id             uintptr
}

func newDrawHandlerMap() *drawHandlerMap {
	return &drawHandlerMap{
		drawHandlerMap: make(map[uintptr]func(func())),
	}
}
func (m *drawHandlerMap) register(fn func(func())) uintptr {
	m.id++
	m.drawHandlerMap[m.id] = fn
	return m.id
}
func (m *drawHandlerMap) unregister(id uintptr) {
	if _, ok := m.drawHandlerMap[id]; !ok {
		panic(fmt.Errorf("unknown draw handler id: %d", id))
	}
	delete(m.drawHandlerMap, id)
}
func (m *drawHandlerMap) invoke(id uintptr, baseDraw func()) {
	if handler, ok := m.drawHandlerMap[id]; ok && handler != nil {
		handler(baseDraw)
	}
}
func (m *drawHandlerMap) isEmpty() bool {
	return len(m.drawHandlerMap) == 0
}
func (m *drawHandlerMap) size() int {
	return len(m.drawHandlerMap)
}
func (m *drawHandlerMap) clear() {
	for id := range m.drawHandlerMap {
		delete(m.drawHandlerMap, id)
	}
}

var globalDrawHandlerMap = newDrawHandlerMap()

//export _go_drawHandler
func _go_drawHandler(handlerId uintptr, widget *C.Fl_Widget) {
	globalDrawHandlerMap.invoke(handlerId, func() {
		C.go_fltk_Widget_basedraw(widget)
	})
}

type helpViewHandlerMap struct {
	helpHandlerMap map[uintptr]func(string)
}

func newHelpViewHandlerMap() *helpViewHandlerMap {
	return &helpViewHandlerMap{
		helpHandlerMap: make(map[uintptr]func(string)),
	}
}
func (m *helpViewHandlerMap) register(id uintptr, fn func(string)) uintptr {
	m.helpHandlerMap[id] = fn
	return id
}
func (m *helpViewHandlerMap) unregister(id uintptr) {
	if _, ok := m.helpHandlerMap[id]; !ok {
		panic(fmt.Errorf("unknown draw handler id: %d", id))
	}
	delete(m.helpHandlerMap, id)
}
func (m *helpViewHandlerMap) invoke(id uintptr, str string) {
	if handler, ok := m.helpHandlerMap[id]; ok && handler != nil {
		handler(str)
	}
}
func (m *helpViewHandlerMap) isEmpty() bool {
	return len(m.helpHandlerMap) == 0
}
func (m *helpViewHandlerMap) size() int {
	return len(m.helpHandlerMap)
}
func (m *helpViewHandlerMap) clear() {
	for id := range m.helpHandlerMap {
		delete(m.helpHandlerMap, id)
	}
}

var globalHelpViewHandlerMap = newHelpViewHandlerMap()

//export _go_helpViewHandler
func _go_helpViewHandler(handlerId uintptr, str *C.char) {
	globalHelpViewHandlerMap.invoke(handlerId, C.GoString(str))
}
