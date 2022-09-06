package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"runtime"

	"github.com/pwiecz/go-fltk"
)

//go:embed 8x8.png
var icon8data []byte

//go:embed 16x16.png
var icon16data []byte

//go:embed 32x32.png
var icon32data []byte

func main() {
	runtime.LockOSThread()
	icon8, _, err := image.Decode(bytes.NewReader(icon8data))
	if err != nil {
		panic(err)
	}
	icon16, _, err := image.Decode(bytes.NewReader(icon16data))
	if err != nil {
		panic(err)
	}
	icon32, _, err := image.Decode(bytes.NewReader(icon32data))
	if err != nil {
		panic(err)
	}
	win := fltk.NewWindow(300, 200)
	rgbIcon8, err := fltk.NewRgbImageFromImage(icon8)
	if err != nil {
		panic(err)
	}
	rgbIcon16, err := fltk.NewRgbImageFromImage(icon16)
	if err != nil {
		panic(err)
	}
	rgbIcon32, err := fltk.NewRgbImageFromImage(icon32)
	if err != nil {
		panic(err)
	}
	win.SetIcons([]*fltk.RgbImage{rgbIcon8, rgbIcon16, rgbIcon32})
	win.End()
	win.Show()
	fltk.Run()
}
