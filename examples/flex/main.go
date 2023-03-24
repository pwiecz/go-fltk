package main

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

var i = 0

func main() {
	win := fltk.NewWindow(300, 200)
	column := fltk.NewFlex(0, 0, 300, 200)
	column.SetType(fltk.COLUMN)
	column.SetGap(5)
	inc := fltk.NewButton(0, 0, 0, 0, "Increment")
	column.Fixed(inc, 40)
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 0, 0, "0")
	dec := fltk.NewButton(0, 0, 0, 0, "Decrement")
	inc.SetCallback(func() {
		i++
		box.SetLabel(strconv.Itoa(i))
	})
	dec.SetCallback(func() {
		i++
		box.SetLabel(strconv.Itoa(i))
	})

	column.Fixed(dec, 40)
	column.End()
	win.End()
	win.Resizable(column)
	win.Show()
	fltk.Run()
}
