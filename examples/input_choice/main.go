package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	inputChoice := fltk.NewInputChoice(5, 5, 100, 20, "Input choice")
	menu := inputChoice.MenuButton()
	menu.Add("An item", func() { fmt.Println("Item selected") })
	menu.Add("Add item at end", func() { menu.Add("Added item", func() { fmt.Println("Added item selected") }) })
	menu.Add("Insert item at start", func() { menu.Insert(0, "Inserted item", func() { fmt.Println("Inserted item selected") }) })
	menu.Add("Remove first item", func() { menu.Remove(0) })
	menu.Add("Replace first label", func() { menu.Replace(0, "New label") })
	
	win.End()
	win.Show()
	fltk.Run()
}
