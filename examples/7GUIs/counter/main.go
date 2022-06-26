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
		WIDGET_WIDTH*2+WIDGET_PADDING*3,
		WIDGET_HEIGHT*3+WIDGET_PADDING*2)
	win.SetLabel("Counter")

	hpack := fltk.NewPack(WIDGET_PADDING, WIDGET_PADDING, win.W(), WIDGET_HEIGHT)
	hpack.SetType(fltk.HORIZONTAL)
	hpack.SetSpacing(WIDGET_PADDING)

	text := fltk.NewOutput(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	text.SetValue("0")

	btn := fltk.NewButton(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	btn.SetLabel("Count")

	value := 0
	btn.SetCallback(func() {
		value++
		text.SetValue(strconv.Itoa(value))
	})

	win.End()
	win.Show()
	fltk.Run()
}
