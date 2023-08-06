package main

import (
	goimage "image"
	"image/color"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(400, 300)

	{
		data := make([]uint8, 0, 10000*3)
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				data = append(data, uint8(x), uint8(y), uint8((x+y)%256))
			}
		}
		image, err := fltk.NewRgbImage(data, 100, 100, 3)
		if err != nil {
			panic(err)
		}
		image.Scale(100, 100, true, true)
		box1 := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 100, 100, "")
		box1.SetImage(image)
	}
	{
		grayImage := goimage.NewGray(goimage.Rect(0, 0, 100, 100))
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				grayImage.Set(x, y, color.Gray{uint8(x + y)})
			}
		}
		image, err := fltk.NewRgbImageFromImage(grayImage)
		if err != nil {
			panic(err)
		}
		box2 := fltk.NewBox(fltk.FLAT_BOX, 0, 100, 100, 100, "")
		box2.SetImage(image)
	}
	{
		rgbaImage := goimage.NewRGBA(goimage.Rect(0, 0, 100, 100))
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				rgbaImage.Set(x, y, color.RGBA{uint8(x), uint8(y), uint8(x + y), 255})
			}
		}
		image, err := fltk.NewRgbImageFromImage(rgbaImage)
		if err != nil {
			panic(err)
		}
		box3 := fltk.NewBox(fltk.FLAT_BOX, 0, 200, 100, 100, "")
		box3.SetImage(image)
	}
	{
		cmykImage := goimage.NewCMYK(goimage.Rect(0, 0, 100, 100))
		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				cmykImage.Set(x, y, color.CMYK{uint8(x), uint8(y), uint8(x + y), uint8(100 + x - y)})
			}
		}
		image, err := fltk.NewRgbImageFromImage(cmykImage)
		if err != nil {
			panic(err)
		}
		box3 := fltk.NewBox(fltk.FLAT_BOX, 100, 0, 100, 100, "")
		box3.SetImage(image)
	}

	win.End()
	win.Show()
	fltk.Run()
}
