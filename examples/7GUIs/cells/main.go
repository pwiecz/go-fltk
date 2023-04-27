package main

import (
	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT = 200
	WIDGET_WIDTH  = 450
)

const (
	MAX_ROW_COUNT = 100 // 0~99
	MAX_COL_COUNT = 26  // A~Z
)

var ctx = NewContext(MAX_ROW_COUNT, MAX_COL_COUNT)

func init() {
	ctx.UpdateCellAtLoc("B1", "5")
	ctx.UpdateCellAtLoc("B2", "1")
	ctx.UpdateCellAtLoc("B3", "10.3")
	ctx.UpdateCellAtLoc("B4", "22.87")
	ctx.UpdateCellAtLoc("B5", "=SUM(B1:B4)")
	ctx.UpdateCellAtLoc("C1", "6")
	ctx.UpdateCellAtLoc("C2", "7")
	ctx.UpdateCellAtLoc("C3", "2")
	ctx.UpdateCellAtLoc("C4", "5")
	ctx.UpdateCellAtLoc("C5", "=SUM(C1:C4)")
	ctx.UpdateCellAtLoc("A5", "Sum")
	ctx.UpdateCellAtLoc("D5", "=SUM(B5:C5)")
}

func main() {
	fltk.SetScheme("gtk+")

	win := fltk.NewWindow(
		WIDGET_WIDTH,
		WIDGET_HEIGHT)
	win.SetLabel("Cells")

	p := NewPanel(win, MAX_ROW_COUNT, MAX_COL_COUNT)
	p.Bind(ctx)

	win.End()
	win.Show()
	fltk.Run()
}
