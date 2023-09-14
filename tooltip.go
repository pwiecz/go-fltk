package fltk

/*
#include <stdlib.h>
#include "tooltip.h"
*/
import "C"
import "unsafe"

func EnableTooltips() {
	C.go_fltk_enable_tooltips()
}

func DisableTooltips() {
	C.go_fltk_disable_tooltips()
}

func AreTooltipsEnabled() bool {
	return C.go_fltk_tooltips_enabled() != 0
}

func TooltipDelay() float32 {
	return float32(C.go_fltk_tooltip_delay())
}

// SetTooltipDelay sets the tooltip delay in seconds
func SetTooltipDelay(delay float32) {
	C.go_fltk_set_tooltip_delay(C.float(delay))
}

// As Fl_Tooltip::enter_area does not copy passed tooltip we have to store it here
// to avoid at least some of memory leaks, when TooltipEnterArea is called many times.
var tooltip *C.char

func TooltipEnterArea(wi Widget, x, y, w, h int, tip string) {
	if tooltip != nil {
		C.free(unsafe.Pointer(tooltip))
	}
	tooltip = C.CString(tip)
	C.go_fltk_tooltip_enter_area(wi.getWidget().ptr(), C.int(x), C.int(y), C.int(w), C.int(h), tooltip)
}
