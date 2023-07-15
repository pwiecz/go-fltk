package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(800, 600)
	box := fltk.NewBox(fltk.NO_BOX, 3, 3, 800-6, 600-6)
	box.SetColor(fltk.WHITE)
	win.End()
	win.Show()

	offs := fltk.NewOffscreen(box.W(), box.H())
	defer offs.Delete()
	offs.Begin()
	fltk.SetDrawColor(fltk.WHITE)
	fltk.DrawRectf(0, 0, box.W(), box.H())
	offs.End()

	box.SetDrawHandler(func(func()) {
		offs.Copy(3, 3, box.W(), box.H(), 0, 0)
	})

	x := 0
	y := 0

	box.SetEventHandler(func(e fltk.Event) bool {
		switch e {
		case fltk.PUSH:
			offs.Begin()
			fltk.SetDrawColor(fltk.RED)
			fltk.SetLineStyle(fltk.SOLID, 3)
			x = fltk.EventX()
			y = fltk.EventY()
			fltk.DrawPoint(x, y)
			offs.End()
			box.Redraw()
			fltk.SetLineStyle(fltk.SOLID, 0)
			return true
		case fltk.DRAG:
			offs.Begin()
			fltk.SetDrawColor(fltk.RED)
			fltk.SetLineStyle(fltk.SOLID, 3)
			nx := fltk.EventX()
			ny := fltk.EventY()
			fltk.DrawLine(x, y, nx, ny)
			x = nx
			y = ny
			offs.End()
			box.Redraw()
			fltk.SetLineStyle(fltk.SOLID, 0)
			return true
		default:
			return false
		}
	})

	fltk.Run()
}
