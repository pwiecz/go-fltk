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
	grid := fltk.NewGrid(0, 0, width, height)
	grid.SetLayout(3, 2, pad, pad)
	makeScaleRow(grid, 0)
	makeThemeRow(grid, 1)
	makeTooltipRow(grid, 2)
	// grid.SetShowGrid(true) // DEBUG
	grid.End()
}

func makeScaleRow(grid *fltk.Grid, row int) {
	scaleLabel := makeAccelLabel(colWidth, rowHeight, "&Scale")
	scaleSpinner := makeScaleSpinner()
	scaleLabel.SetCallback(func() { scaleSpinner.TakeFocus() })
	scaleSpinner.TakeFocus()
	grid.SetWidget(scaleLabel, row, 0, fltk.GridLeft)
	grid.SetWidget(scaleSpinner, row, 1, fltk.GridFill)
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

func makeThemeRow(grid *fltk.Grid, row int) {
	themeLabel := makeAccelLabel(colWidth, rowHeight, "&Theme")
	themeChoice := makeThemeChoice()
	themeLabel.SetCallback(func() { themeChoice.TakeFocus() })
	grid.SetWidget(themeLabel, row, 0, fltk.GridLeft)
	grid.SetWidget(themeChoice, row, 1, fltk.GridFill)
}

func makeThemeChoice() *fltk.Choice {
	choice := fltk.NewChoice(0, 0, colWidth, rowHeight)
	choice.SetTooltip("Sets the application's theme.")
	for i, name := range []string{"&Base", "&Gleam", "G&tk", "&Oxy",
		"&Plastic"} {
		theme := strings.ReplaceAll(name, "&", "")
		if theme == "Oxy" {
			choice.SetValue(i)
			fltk.SetScheme(theme)
		}
		choice.Add(name, func() { fltk.SetScheme(theme) })
	}
	return choice
}

func makeTooltipRow(grid *fltk.Grid, row int) {
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
	grid.SetWidgetWithSpan(checkButton, row, 0, 1, 2, fltk.GridCenter)
}

func makeAccelLabel(width, height int, label string) *fltk.Button {
	button := fltk.NewButton(0, 0, width, height, label)
	button.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	button.SetBox(fltk.NO_BOX)
	button.ClearVisibleFocus()
	return button
}
