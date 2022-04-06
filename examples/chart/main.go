package main

import (
	"github.com/pwiecz/go-fltk"
	"math/rand"
	"strconv"
)

func main() {
	win := fltk.NewWindow(301, 440)
	colors := []fltk.Color{
		fltk.BLUE,
		fltk.RED,
		fltk.YELLOW,
		fltk.GREEN,
		fltk.CYAN,
		fltk.DARK_MAGENTA,
	}

	ch := fltk.NewChart(1, 1, 300, 300, "Example chart")

	ch.SetTextColor(fltk.DARK_RED)
	ch.SetType(fltk.SPECIALPIE_CHART)
	ch.SetAutosize(true)

	valueNameEditor := fltk.NewInput(85, ch.Y()+ch.H()+30, 200, 20, "Value name")
	valueEditor := fltk.NewInput(85, valueNameEditor.Y()+valueNameEditor.H()+20, 200, 20, "Value")

	addValueButton := fltk.NewButton(222, valueEditor.Y()+valueEditor.H()+10, 70, 20, "Add value")
	addValueButton.SetCallback(func() {
		val, err := strconv.ParseFloat(valueEditor.Value(), 64)
		if err != nil {
			panic(err)
		}
		ch.Add(val, colors[rand.Intn(len(colors))], valueNameEditor.Value())
	})

	win.End()
	win.Show()
	fltk.Run()
}
