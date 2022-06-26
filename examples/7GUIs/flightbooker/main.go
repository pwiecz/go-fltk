package main

import (
	"fmt"
	"time"

	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT  = 25
	WIDGET_PADDING = 5
	WIDGET_WIDTH   = 200
)

const (
	DATE_FORMAT = "02.01.2006"
)

type BookOption int

const (
	BookOptionOneWay BookOption = iota
	BookOptionReturn
)

func main() {
	var (
		option       BookOption
		optionChoice *fltk.Choice
		startInput   *fltk.Input
		returnInput  *fltk.Input
		bookBtn      *fltk.Button
	)

	update := func() {
		option = BookOption(optionChoice.Value())
		switch option {
		case BookOptionOneWay:
			returnInput.Deactivate()
			if _, ok := validateInput(startInput); ok {
				bookBtn.Activate()
			} else {
				bookBtn.Deactivate()
			}
		case BookOptionReturn:
			returnInput.Activate()
			t1, ok1 := validateInput(startInput)
			t2, ok2 := validateInput(returnInput)
			if ok1 && ok2 && !t1.After(t2) {
				bookBtn.Activate()
			} else {
				bookBtn.Deactivate()
			}
		}
	}

	fltk.SetScheme("gtk+")

	win := fltk.NewWindow(
		WIDGET_WIDTH+WIDGET_PADDING*2,
		WIDGET_HEIGHT*4+WIDGET_PADDING*5)
	win.SetLabel("Book Flight")

	vpack := fltk.NewPack(WIDGET_PADDING, WIDGET_PADDING, win.W()-WIDGET_PADDING*2, win.H()-WIDGET_PADDING*2)
	vpack.SetSpacing(WIDGET_PADDING)

	option = BookOptionOneWay

	optionChoice = fltk.NewChoice(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	optionChoice.Add("one-way flight", update)
	optionChoice.Add("return flight", update)
	optionChoice.SetValue(int(option))

	now := time.Now()
	startInput = fltk.NewInput(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	startInput.SetValue(now.Format(DATE_FORMAT))
	startInput.SetCallbackCondition(fltk.WhenChanged)
	startInput.SetCallback(update)

	returnInput = fltk.NewInput(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	returnInput.SetValue(now.Format(DATE_FORMAT))
	returnInput.SetCallbackCondition(fltk.WhenChanged)
	returnInput.SetCallback(update)
	// defaultInputColor = startInput.Color()

	bookBtn = fltk.NewButton(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	bookBtn.SetLabel("Book")
	bookBtn.SetCallback(func() {
		switch option {
		case BookOptionOneWay:
			fltk.MessageBox("Book successful", fmt.Sprintf("You have booked a one-way flight on %s.", startInput.Value()))
		case BookOptionReturn:
			fltk.MessageBox("Book successful", fmt.Sprintf("You have booked a return flight on %s and %s.", startInput.Value(), returnInput.Value()))
		}
	})

	update()

	win.End()
	win.Show()
	fltk.Run()
}

func validateInput(input *fltk.Input) (time.Time, bool) {
	defer input.Redraw()

	t, err := time.Parse(DATE_FORMAT, input.Value())
	if err != nil {
		input.SetColor(fltk.RED)
		return t, false
	}

	input.SetColor(fltk.BACKGROUND2_COLOR)
	return t, true
}
