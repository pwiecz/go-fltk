package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	packV := fltk.NewPack(0, 0, 300, 200)
	packV.SetType(fltk.VERTICAL)
	buttonPack := fltk.NewPack(0, 0, 300, 20)
	buttonPack.SetType(fltk.HORIZONTAL)
	prev := fltk.NewButton(0, 0, 100, 20, "Prev")
	next := fltk.NewButton(0, 0, 100, 20, "Next")
	buttonPack.End()
	wizard := fltk.NewWizard(0, 10, 300, 250)
	pack1 := fltk.NewPack(10, 10, 200, 100, "Pane1")
	pack1.SetType(fltk.VERTICAL)
	fltk.NewBox(fltk.NO_BOX, 0, 0, 100, 20, "Pane 1")
	pack1.End()
	wizard.Add(pack1)
	pack2 := fltk.NewPack(10, 10, 200, 100, "Pane2")
	pack2.SetType(fltk.VERTICAL)
	fltk.NewBox(fltk.NO_BOX, 0, 0, 100, 20, "Pane 2")
	pack2.End()
	wizard.Add(pack2)
	wizard.End()
	packV.End()
	prev.SetCallback(wizard.Prev)
	next.SetCallback(wizard.Next)

	win.End()
	win.Resizable(packV)

	win.Show()
	fltk.Run()

}
