package main

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT  = 25
	WIDGET_PADDING = 5
	WIDGET_WIDTH   = 70
)

func main() {
	fltk.SetScheme("gtk+")

	win := fltk.NewWindow(
		WIDGET_WIDTH*2+WIDGET_PADDING*2,
		WIDGET_HEIGHT+WIDGET_PADDING*2)
	win.SetLabel("Counter")

	row := fltk.NewFlex(WIDGET_PADDING, WIDGET_PADDING, WIDGET_WIDTH*2, WIDGET_HEIGHT)
	row.SetType(fltk.ROW)
	row.SetGap(WIDGET_PADDING)

	text := fltk.NewOutput(0, 0, 0, 0)
	text.SetValue("0")

	btn := fltk.NewButton(0, 0, 0, 0)
	btn.SetLabel("Count")

	value := 0
	btn.SetCallback(func() {
		value++
		text.SetValue(strconv.Itoa(value))
	})

	row.End()
	win.End()
	win.Show()
	fltk.Run()
}
