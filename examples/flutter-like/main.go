package main

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

// FLTK uses an RGBI color representation, the I is an index into FLTK's color map
// Passing 00 as I will use the RGB part of the value
const GRAY = 0x75757500
const LIGHT_GRAY = 0xeeeeee00
const BLUE = 0x42A5F500
const SEL_BLUE = 0x2196F300
const WIDTH = 600
const HEIGHT = 400

func main() {
	curr := 0
	fltk.InitStyles()
	win := fltk.NewWindow(WIDTH, HEIGHT)
	win.SetLabel("Flutter-like")
	win.SetColor(fltk.WHITE)
	bar := fltk.NewBox(fltk.FLAT_BOX, 0, 0, WIDTH, 60, "    FLTK App!")
	bar.SetDrawHandler(func(baseDraw func()) { // Shadow under the bar
		fltk.DrawBox(fltk.FLAT_BOX, 0, 0, WIDTH, 63, LIGHT_GRAY)
		baseDraw()
	})
	bar.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	bar.SetLabelColor(255) // this uses the index into the color map, here it's white
	bar.SetColor(BLUE)
	bar.SetLabelSize(22)
	text := fltk.NewBox(fltk.NO_BOX, 250, 180, 100, 40, "You have pushed the button this many times:")
	text.SetLabelSize(18)
	text.SetLabelFont(fltk.TIMES)
	count := fltk.NewBox(fltk.NO_BOX, 250, 180+40, 100, 40, "0")
	count.SetLabelSize(36)
	count.SetLabelColor(GRAY)
	btn := fltk.NewButton(WIDTH-100, HEIGHT-100, 60, 60, "@+6plus") // this translates into a plus sign
	btn.SetColor(BLUE)
	btn.SetSelectionColor(SEL_BLUE)
	btn.SetLabelColor(255)
	btn.SetBox(fltk.OFLAT_BOX)
	btn.ClearVisibleFocus()
	btn.SetCallback(func() {
		curr += 1
		count.SetLabel(strconv.Itoa(curr))
	})
	win.End()
	win.Show()
	fltk.Run()
}
