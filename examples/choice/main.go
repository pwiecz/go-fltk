package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	choice := fltk.NewChoice(5, 5, 100, 20, "Choice")
	choice.Add("An item", func() { fmt.Println("Item selected") })
	choice.Add("Add item at end", func() { choice.Add("Added item", func() { fmt.Println("Added item selected") }) })
	choice.Add("Insert item at start", func() { choice.Insert(0, "Inserted item", func() { fmt.Println("Inserted item selected") }) })
	choice.Add("Remove first item", func() { choice.Remove(0) })
	choice.Add("Replace first label", func() { choice.Replace(0, "New label") })
	win.End()
	win.Show()
	fltk.Run()
}
