package fltk

/*
#include <stdlib.h>
#include "drawings.h"
*/
import "C"
import "unsafe"

func SetDrawColor(color Color) {
	C.go_fltk_color(C.uint(color))
}

func Draw(text string, x, y, w, h int, align Align) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_draw(textStr, C.int(x), C.int(y), C.int(w), C.int(h), C.uint(align))
}

func DrawBox(boxType BoxType, x, y, w, h int, color Color) {
	C.go_fltk_draw_box(
		C.int(boxType), C.int(x), C.int(y), C.int(w), C.int(h), C.uint(color))
}

func SetDrawFont(font Font, size int) {
	C.go_fltk_set_draw_font(C.int(font), C.int(size))
}

func PushClip(x, y, w, h int) {
	C.go_fltk_push_clip(C.int(x), C.int(y), C.int(w), C.int(h))
}

func PushNoClip() {
	C.go_fltk_push_no_clip()
}

func PopClip() {
	C.go_fltk_pop_clip()
}

func DrawPoint(x, y int) {
	C.go_fltk_point(C.int(x), C.int(y))
}

func SetLineStyle(style LineStyle, width int) {
	C.go_fltk_line_style(C.int(style), C.int(width), nil)
}

func DrawRect(x, y, w, h int) {
	C.go_fltk_rect(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawFocusRect(x, y, w, h int) {
	C.go_fltk_focus_rect(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawRectWithColor(x, y, w, h int, col Color) {
	C.go_fltk_rect_with_color(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

func DrawRectf(x, y, w, h int) {
	C.go_fltk_rectf(C.int(x), C.int(y), C.int(w), C.int(h))
}

func DrawRectfWithColor(x, y, w, h int, col Color) {
	C.go_fltk_rectf_with_color(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

func DrawLine(x, y, x1, y1 int) {
	C.go_fltk_line(C.int(x), C.int(y), C.int(x1), C.int(y1))
}

func DrawLine2(x, y, x1, y1, x2, y2 int) {
	C.go_fltk_line2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawLoop(x, y, x1, y1, x2, y2 int) {
	C.go_fltk_loop(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawLoop2(x, y, x1, y1, x2, y2, x3, y3 int) {
	C.go_fltk_loop2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

func DrawPolygon(x, y, x1, y1, x2, y2 int) {
	C.go_fltk_polygon(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func DrawPolygon2(x, y, x1, y1, x2, y2, x3, y3 int) {
	C.go_fltk_polygon2(C.int(x), C.int(y), C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

func DrawXyLine(x, y, x1 int) {
	C.go_fltk_xyline(C.int(x), C.int(y), C.int(x1))
}

func DrawXyLine2(x, y, x1, y2 int) {
	C.go_fltk_xyline2(C.int(x), C.int(y), C.int(x1), C.int(y2))
}

func DrawXyLine3(x, y, x1, y2, x3 int) {
	C.go_fltk_xyline3(C.int(x), C.int(y), C.int(x1), C.int(y2), C.int(x3))
}

func DrawYxLine(x, y, y1 int) {
	C.go_fltk_yxline(C.int(x), C.int(y), C.int(y1))
}

func DrawYxLine2(x, y, y1, x2 int) {
	C.go_fltk_yxline2(C.int(x), C.int(y), C.int(y1), C.int(x2))
}

func DrawYxLine3(x, y, y1, x2, y3 int) {
	C.go_fltk_yxline3(C.int(x), C.int(y), C.int(y1), C.int(x2), C.int(y3))
}

func DrawArc(x, y, w, h int, a1, a2 float64) {
	C.go_fltk_arc(C.int(x), C.int(y), C.int(w), C.int(h), C.double(a1), C.double(a2))
}

func DrawPie(x, y, w, h int, a1, a2 float64) {
	C.go_fltk_pie(C.int(x), C.int(y), C.int(w), C.int(h), C.double(a1), C.double(a2))
}

func DrawArc2(x, y, r, start, end float64) {
	C.go_fltk_arc2(C.double(x), C.double(y), C.double(r), C.double(start), C.double(end))
}

func DrawCirlce(x, y, r float64) {
	C.go_fltk_circle(C.double(x), C.double(y), C.double(r))
}

// returns the dx, dy, w, h of the string
func TextExtents(text string) (int, int, int, int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	dx := C.int(0)
	dy := C.int(0)
	w := C.int(0)
	h := C.int(0)
	C.go_fltk_text_extents(textStr, &dx, &dy, &w, &h)
	return int(dx), int(dy), int(w), int(h)
}

func DrawTextAngled(text string, x, y, angle int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_draw2(C.int(angle), textStr, C.int(x), C.int(y))
}

func DrawRtl(text string, x, y int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	C.go_fltk_rtl_draw(textStr, C.int(len(text)), C.int(x), C.int(y))
}

func MeasureText(text string, draw_symbols bool) (int, int) {
	textStr := C.CString(text)
	defer C.free(unsafe.Pointer(textStr))
	x := C.int(0)
	y := C.int(0)
	d := 0
	if draw_symbols {
		d = 1
	}
	C.go_fltk_measure(textStr, &x, &y, C.int(d))
	return int(x), int(y)
}

func DrawCheck(x, y, w, h int, col Color) {
	C.go_fltk_draw_check(C.int(x), C.int(y), C.int(w), C.int(h), C.uint(col))
}

type Offscreen struct {
	inner *C.GOffscreen
}

func NewOffscreen(w, h int) *Offscreen {
	o := &Offscreen{
		inner: C.go_fltk_create_offscreen(C.int(w), C.int(h)),
	}
	return o
}

func (offs *Offscreen) Begin() {
	C.go_fltk_begin_offscreen(offs.inner)
}

func (offs *Offscreen) End() {
	C.go_fltk_end_offscreen()
}

func (offs *Offscreen) Rescale() {
	C.go_fltk_rescale_offscreen(&offs.inner)
}

func (offs *Offscreen) Delete() {
	C.go_fltk_delete_offscreen(offs.inner)
}

func (offs *Offscreen) IsValid() bool {
	return offs.inner != nil
}

func (offs *Offscreen) Copy(x, y, w, h, srcx, srcy int) {
	C.go_fltk_copy_offscreen(C.int(x), C.int(y), C.int(w), C.int(h), offs.inner, C.int(srcx), C.int(srcy))
}
