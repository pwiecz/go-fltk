package fltk

import (
	/*
		#include "grid.h"
	*/
	"C"
	"unsafe"
)

type Grid struct {
	Group
}

func NewGrid(x, y, w, h int, text ...string) *Grid {
	g := &Grid{}
	initWidget(g, unsafe.Pointer(C.go_fltk_new_Grid(C.int(x), C.int(y), C.int(w), C.int(h), cStringOpt(text))))
	return g
}

func (g *Grid) SetLayout(rows, columns, margin, gap int) {
	C.go_fltk_Grid_set_layout((*C.Fl_Grid)(g.ptr()), C.int(rows), C.int(columns), C.int(margin), C.int(gap))
}

// ClearLayout reset the layout without removing it's children.
// First removes all rows and columns from the grid setting them to 'zero',
// then hides all the children on the Grid.
// These can be displayed again by defining a new layout,
// showing each child and then adding them again
func (g *Grid) ClearLayout() {
	C.go_fltk_Grid_clear_layout((*C.Fl_Grid)(g.ptr()))
}

func (g *Grid) SetShowGrid(show bool) {
	if show {
		C.go_fltk_Grid_set_show_grid((*C.Fl_Grid)(g.ptr()), 1)
	} else {
		C.go_fltk_Grid_set_show_grid((*C.Fl_Grid)(g.ptr()), 0)
	}
}

func (g *Grid) SetShowGridAndColor(show bool, color Color) {
	if show {
		C.go_fltk_Grid_set_show_grid_and_color((*C.Fl_Grid)(g.ptr()), 1, C.uint(color))
	} else {
		C.go_fltk_Grid_set_show_grid_and_color((*C.Fl_Grid)(g.ptr()), 0, C.uint(color))
	}
}

func (g *Grid) SetColumnGap(column, gap int) {
	C.go_fltk_Grid_set_column_gap((*C.Fl_Grid)(g.ptr()), C.int(column), C.int(gap))
}
func (g *Grid) ColumnGap(column int) int {
	return int(C.go_fltk_Grid_column_gap((*C.Fl_Grid)(g.ptr()), C.int(column)))
}

func (g *Grid) SetColumnWeight(column, weight int) {
	C.go_fltk_Grid_set_column_weight((*C.Fl_Grid)(g.ptr()), C.int(column), C.int(weight))
}
func (g *Grid) ColumnWeight(column int) int {
	return int(C.go_fltk_Grid_column_weight((*C.Fl_Grid)(g.ptr()), C.int(column)))
}

func (g *Grid) SetRowGap(row, gap int) {
	C.go_fltk_Grid_set_row_gap((*C.Fl_Grid)(g.ptr()), C.int(row), C.int(gap))
}
func (g *Grid) RowGap(row int) int {
	return int(C.go_fltk_Grid_row_gap((*C.Fl_Grid)(g.ptr()), C.int(row)))
}

// Gets the current margins of the Grid.
// all_equal is 1 if all margins are equal, 0 otherwise
// margins is an array of the Grid margins, in order, {left, top, right, bottom}.
func (g *Grid) Margin() (all_equal int, margins [4]int) {
	var left, top, right, bottom C.int

	all_equal = int(C.go_fltk_Grid_margin(
		(*C.Fl_Grid)(g.ptr()),
		&left, &top, &right, &bottom,
	))

	margins = [4]int{int(left), int(top), int(right), int(bottom)}

	return all_equal, margins
}

func (g *Grid) SetMargin(left, top, right, bottom int) {
	C.go_fltk_Grid_set_margin((*C.Fl_Grid)(g.ptr()), C.int(left), C.int(top), C.int(right), C.int(bottom))
}

func (g *Grid) SetRowWeight(row, weight int) {
	C.go_fltk_Grid_set_row_weight((*C.Fl_Grid)(g.ptr()), C.int(row), C.int(weight))
}
func (g *Grid) RowWeight(row int) int {
	return int(C.go_fltk_Grid_row_weight((*C.Fl_Grid)(g.ptr()), C.int(row)))
}

type GridAlign int

var (
	GridBottom       = GridAlign(C.go_FL_GRID_BOTTOM)
	GridBottomLeft   = GridAlign(C.go_FL_GRID_BOTTOM_LEFT)
	GridBottomRight  = GridAlign(C.go_FL_GRID_BOTTOM_RIGHT)
	GridCenter       = GridAlign(C.go_FL_GRID_CENTER)
	GridFill         = GridAlign(C.go_FL_GRID_FILL)
	GridHorizontal   = GridAlign(C.go_FL_GRID_HORIZONTAL)
	GridLeft         = GridAlign(C.go_FL_GRID_LEFT)
	GridProportional = GridAlign(C.go_FL_GRID_PROPORTIONAL)
	GridRight        = GridAlign(C.go_FL_GRID_RIGHT)
	GridTop          = GridAlign(C.go_FL_GRID_TOP)
	GridTopLeft      = GridAlign(C.go_FL_GRID_TOP_LEFT)
	GridTopRight     = GridAlign(C.go_FL_GRID_TOP_RIGHT)
	GridVertical     = GridAlign(C.go_FL_GRID_VERTICAL)
)

func (g *Grid) SetWidget(w Widget, row, column int, align GridAlign) {
	C.go_fltk_Grid_set_widget((*C.Fl_Grid)(g.ptr()), w.getWidget().ptr(), C.int(row), C.int(column), C.int(align))
}
func (g *Grid) SetWidgetWithSpan(w Widget, row, column, rowSpan, columnSpan int, align GridAlign) {
	C.go_fltk_Grid_set_widget_with_span((*C.Fl_Grid)(g.ptr()), w.getWidget().ptr(), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(align))
}
