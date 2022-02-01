package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	table := fltk.NewTableRow(5, 5, 295, 190)
	table.SetRowCount(2)
	table.SetColumnCount(3)
	table.SetBox(fltk.NO_BOX)
	table.Begin()
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			x, y, w, h, err := table.FindCell(fltk.ContextCell, i, j)
			if err == nil {
				if j == 0 {
					fltk.NewInput(x, y, w, h, "")
				} else {
					fltk.NewButton(x, y, w, h, "button")
				}
			}
		}
	}
	table.End()
	win.End()
	win.Show()
	fltk.Run()
}
