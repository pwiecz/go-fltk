package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	column := fltk.NewFlex(0, 0, 300, 200)
	column.SetType(fltk.COLUMN)
	column.SetGap(5)
	i := fltk.NewButton(0, 0, 0, 0, "Increment")
	column.Fixed(i, 40)
	fltk.NewBox(fltk.FLAT_BOX, 0, 0, 0, 0, "0")
	d := fltk.NewButton(0, 0, 0, 0, "Decrement")
	column.Fixed(d, 40)
	column.End()
	win.End()
	win.Show()
	fltk.Run()
}
