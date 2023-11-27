package main

import "github.com/pwiecz/go-fltk"

func main() {
	win := fltk.NewWindow(320, 180)
	grid := fltk.NewGrid(0, 0, win.W(), win.H())
	grid.SetLayout(3, 3, 10, 10)
	grid.SetColor(fltk.WHITE)
	b0 := fltk.NewButton(0, 0, 0, 0, "New")
	b1 := fltk.NewButton(0, 0, 0, 0, "Options")
	b2 := fltk.NewButton(0, 0, 0, 0, "About")
	b3 := fltk.NewButton(0, 0, 0, 0, "Help")
	b4 := fltk.NewButton(0, 0, 0, 0, "Quit")
	grid.SetWidget(b0, 0, 0, fltk.GridFill)
	grid.SetWidget(b1, 0, 2, fltk.GridFill)
	grid.SetWidget(b2, 1, 1, fltk.GridFill)
	grid.SetWidget(b3, 2, 0, fltk.GridFill)
	grid.SetWidget(b4, 2, 2, fltk.GridFill)
	grid.SetShowGrid(false)
	grid.End()
	win.End()
	win.Resizable(grid)
	win.Show()
	fltk.Run()
}
