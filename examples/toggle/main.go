package main

import (
	"github.com/pwiecz/go-fltk"
)

type MyToggleButton struct {
	btn *fltk.ToggleButton
}

func NewMyToggleButton(x, y, w, h int) *MyToggleButton {
	btn := fltk.NewToggleButton(x, y, w, h, "@+9circle")
	btn.SetColor(0x58585800)
	btn.SetSelectionColor(0x00008B00)
	btn.SetBox(fltk.RFLAT_BOX)
	btn.SetDownBox(fltk.RFLAT_BOX)
	btn.ClearVisibleFocus()
	btn.SetLabelColor(fltk.WHITE)
	btn.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	btn.SetCallback(func() {
		parent := btn.Parent()
		if btn.Value() {
			btn.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_RIGHT)
			parent.Redraw()
		} else {
			btn.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
			parent.Redraw()
		}
	})
	return &MyToggleButton{btn}
}

func main() {
	fltk.InitStyles()
	fltk.SetBackgroundColor(0, 0, 0)
	win := fltk.NewWindow(200, 200)
	NewMyToggleButton(70, 90, 60, 15)
	win.End()
	win.Show()
	fltk.Run()
}
