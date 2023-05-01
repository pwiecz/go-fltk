package main

import (
	"strings"

	"github.com/pwiecz/go-fltk"
)

const (
	pad       = 6
	rowHeight = 32
	colWidth  = 60
)

func main() {
	window := makeWindow()
	window.Show()
	fltk.Run()
}

func makeWindow() *fltk.Window {
	width := 200
	height := 115
	window := fltk.NewWindow(width, height)
	window.SetLabel("UI Config")
	makeWidgets(width, height)
	window.End()
	return window
}

func makeWidgets(width, height int) {
	colFlex := fltk.NewFlex(0, 0, width, height)
	colFlex.SetSpacing(pad)
	rowFlex := makeScaleRow(width, rowHeight)
	colFlex.Fixed(rowFlex, rowHeight)
	rowFlex = makeThemeRow(rowHeight, rowHeight)
	colFlex.Fixed(rowFlex, rowHeight)
	rowFlex = makeTooltipRow(rowHeight, rowHeight)
	colFlex.Fixed(rowFlex, rowHeight)
	colFlex.End()
}

func makeScaleRow(width, height int) *fltk.Flex {
	rowFlex := fltk.NewFlex(0, 0, width, height)
	rowFlex.SetType(fltk.ROW)
	rowFlex.SetSpacing(pad)
	scaleLabel := makeAccelLabel(colWidth, rowHeight, "&Scale")
	scaleSpinner := makeScaleSpinner()
	scaleLabel.SetCallback(func() { scaleSpinner.TakeFocus() })
	rowFlex.Fixed(scaleLabel, colWidth)
	rowFlex.End()
	scaleSpinner.TakeFocus()
	return rowFlex
}

func makeScaleSpinner() *fltk.Spinner {
	spinner := fltk.NewSpinner(0, 0, colWidth, rowHeight)
	spinner.SetTooltip("Sets the application's scale.")
	spinner.SetType(fltk.SPINNER_FLOAT_INPUT)
	spinner.SetMinimum(0.5)
	spinner.SetMaximum(3.5)
	spinner.SetStep(0.1)
	spinner.SetValue(float64(fltk.ScreenScale(0)))
	spinner.SetCallback(func() {
		fltk.SetScreenScale(0, float32(spinner.Value()))
	})
	return spinner
}

func makeThemeRow(width, height int) *fltk.Flex {
	rowFlex := fltk.NewFlex(0, 0, width, height)
	rowFlex.SetType(fltk.ROW)
	rowFlex.SetSpacing(pad)
	themeLabel := makeAccelLabel(colWidth, rowHeight, "&Theme")
	themeChoice := makeThemeChoice()
	themeLabel.SetCallback(func() { themeChoice.TakeFocus() })
	rowFlex.Fixed(themeLabel, colWidth)
	rowFlex.End()
	return rowFlex
}

func makeThemeChoice() *fltk.Choice {
	choice := fltk.NewChoice(0, 0, colWidth, rowHeight)
	choice.SetTooltip("Sets the application's theme.")
	for i, name := range []string{"&Base", "&Gleam", "G&tk", "&Oxy", "&Plastic"} {
		theme := strings.ReplaceAll(name, "&", "")
		if theme == "Oxy" {
			choice.SetValue(i)
			fltk.SetScheme(theme)
		}
		choice.Add(name, func() { fltk.SetScheme(theme) })
	}
	return choice
}

func makeTooltipRow(width, height int) *fltk.Flex {
	rowFlex := fltk.NewFlex(0, 0, width, height)
	rowFlex.SetType(fltk.ROW)
	rowFlex.SetSpacing(pad)
	padBox := fltk.NewBox(fltk.NO_BOX, 0, 0, colWidth, rowHeight)
	checkButton := fltk.NewCheckButton(colWidth, 0, colWidth, rowHeight,
		"S&how Tooltips")
	checkButton.SetTooltip("If checked the application shows tooltips.")
	checkButton.SetValue(fltk.AreTooltipsEnabled())
	checkButton.SetCallback(func() {
		if checkButton.Value() {
			fltk.EnableTooltips()
		} else {
			fltk.DisableTooltips()
		}
	})
	rowFlex.Fixed(padBox, colWidth)
	rowFlex.End()
	return rowFlex
}

func makeAccelLabel(width, height int, label string) *fltk.Button {
	button := fltk.NewButton(0, 0, width, height, label)
	button.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	button.SetBox(fltk.NO_BOX)
	button.ClearVisibleFocus()
	return button
}
