package main

import (
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	tabs := fltk.NewTabs(5, 5, 290, 190)
	pack := fltk.NewPack(10, 10, 200, 100, "Tab1")
	pack.SetType(fltk.VERTICAL)
	fltk.NewButton(0, 0, 100, 20, "Button")
	pack.End()
	tabs.Add(pack)
	pack2 := fltk.NewPack(10, 10, 200, 100, "Tab2")
	pack2.SetType(fltk.VERTICAL)
	pack2.End()
	tabs.Add(pack2)
	tabs.End()
	tabs.Resizable(pack)

	win.End()
	win.Resizable(tabs)

	win.Show()
	fltk.Run()

}
