package fltk

/*
#include "fltk.h"
*/
import "C"

func EventButton1() bool {
	return C.go_fltk_event_button1() != 0
}
func EventX() int {
	return int(C.go_fltk_event_x())
}
func EventY() int {
	return int(C.go_fltk_event_y())
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
func EventState() int {
	return int(C.go_fltk_event_state())
}
