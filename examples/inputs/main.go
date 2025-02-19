package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	y := 0
	fltk.NewInput(70, y, 220, 20, "Normal:")
	y += 35
	fltk.NewIntInput(70, y, 220, 20, "Int:")
	y += 35
	fltk.NewFloatInput(70, y, 220, 20, "Float:")
	y += 35
	fltk.NewSecretInput(70, y, 220, 20, "Secret:")
	y += 35
	output := fltk.NewOutput(70, y, 220, 20, "Output:")
	output.SetValue("value")
	win.End()
	win.Show()
	fltk.Run()
}
