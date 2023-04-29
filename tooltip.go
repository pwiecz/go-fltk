package fltk

/*
#include "tooltip.h"
*/
import "C"

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
