package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(400, 300)
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 400, 300, "")
	svgImage, err := fltk.NewSvgImageFromString(`<svg height="200" width="200">
<polygon points="10,10 20,100 10,190 100,180 190,190 180,100 190,10 100,20" style="stroke:blue;stroke-width:5;fill:red"/>
</svg>`)
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		svgImage.Scale(100, 100, true, true)
		box.SetImage(svgImage)
	}
	win.SetIcons([]*fltk.RgbImage{&svgImage.RgbImage})
	win.End()
	win.Show()
	fltk.Run()
}
