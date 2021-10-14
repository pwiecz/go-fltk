package fltk

/*
#include "events.h"
*/
import "C"

type MouseButton int

var (
	LeftMouse   = MouseButton(C.go_FL_LEFT_MOUSE)
	MiddleMouse = MouseButton(C.go_FL_MIDDLE_MOUSE)
	RightMouse  = MouseButton(C.go_FL_RIGHT_MOUSE)
)

func EventType() Event {
	return Event(C.go_fltk_event())
}

func EventButton() MouseButton {
	return MouseButton(C.go_fltk_event_button())
}

func EventButton1() bool {
	return C.go_fltk_event_button1() != 0
}
func EventX() int {
	return int(C.go_fltk_event_x())
}
func EventY() int {
	return int(C.go_fltk_event_y())
}
func EventXRoot() int {
	return int(C.go_fltk_event_x_root())
}
func EventYRoot() int {
	return int(C.go_fltk_event_y_root())
}
func EventDX() int {
	return int(C.go_fltk_event_dx())
}
func EventDY() int {
	return int(C.go_fltk_event_dy())
}
func EventKey() int {
	return int(C.go_fltk_event_key())
}
func EventIsClick() bool {
	return C.go_fltk_event_is_click() != 0
}

var (
	SHIFT       = int(C.go_FL_SHIFT)
	CAPS_LOCK   = int(C.go_FL_CAPS_LOCK)
	CTRL        = int(C.go_FL_CTRL)
	ALT         = int(C.go_FL_ALT)
	NUM_LOCK    = int(C.go_FL_NUM_LOCK)
	META        = int(C.go_FL_META)
	SCROLL_LOCK = int(C.go_FL_SCROLL_LOCK)
	BUTTON1     = int(C.go_FL_BUTTON1)
	BUTTON2     = int(C.go_FL_BUTTON2)
	BUTTON3     = int(C.go_FL_BUTTON3)
)

func EventState() int {
	return int(C.go_fltk_event_state())
}
