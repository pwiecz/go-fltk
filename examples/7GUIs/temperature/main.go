package main

import (
	"log"
	"strconv"

	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT  = 25
	WIDGET_PADDING = 5
	WIDGET_WIDTH   = 70
)

func main() {
	fltk.SetScheme("gtk+")
	fltk.InitStyles()

	var cInput, fInput *fltk.FloatInput

	win := fltk.NewWindow(
		WIDGET_WIDTH*4+WIDGET_PADDING*5,
		WIDGET_HEIGHT+WIDGET_PADDING*2)
	win.SetLabel("TempConv")

	pack := fltk.NewPack(WIDGET_PADDING, WIDGET_PADDING, win.W()-WIDGET_PADDING*2, win.H()-WIDGET_PADDING*2)
	pack.SetType(fltk.HORIZONTAL)
	pack.SetSpacing(WIDGET_PADDING)

	cInput = fltk.NewFloatInput(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	cInput.SetCallbackCondition(fltk.WhenChanged)
	cInput.SetCallback(func() {
		cVal, err := strconv.ParseFloat(cInput.Value(), 64)
		if err != nil {
			return
		}
		fVal := cVal*(9.0/5.0) + 32.0

		// log to make sure no circular calling
		log.Printf("conv %fC to %fF", cVal, fVal)
		fInput.SetValue(strconv.FormatFloat(fVal, 'f', 2, 64))
	})

	cText := fltk.NewBox(fltk.NO_BOX, 0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	cText.SetLabel("Celsius = ")

	fInput = fltk.NewFloatInput(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	fInput.SetCallbackCondition(fltk.WhenChanged)
	fInput.SetCallback(func() {
		fVal, err := strconv.ParseFloat(fInput.Value(), 64)
		if err != nil {
			return
		}
		cVal := (fVal - 32.0) * (5.0 / 9.0)

		// log to make sure no circular calling
		log.Printf("conv %fC to %fF", fVal, cVal)
		cInput.SetValue(strconv.FormatFloat(cVal, 'f', 2, 64))
	})

	fText := fltk.NewBox(fltk.NO_BOX, 0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	fText.SetLabel("Fahrenheit")

	win.End()
	win.Show()
	fltk.Run()
}
