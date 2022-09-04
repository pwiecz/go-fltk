package main

import (
	"time"

	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT      = 25
	WIDGET_PADDING     = 10
	WIDGET_WIDTH       = 180
	WIDGET_LABEL_WIDTH = 110
)

const (
	DURATION_DEFAULT = 15.0
	DURATION_MAXIMUM = 30.0
)

func main() {
	fltk.SetScheme("gtk+")
	fltk.Lock() // enable the FLTK lock mechanism

	startTime := time.Now()

	win := fltk.NewWindow(
		WIDGET_LABEL_WIDTH+WIDGET_WIDTH+WIDGET_PADDING*2,
		WIDGET_HEIGHT*4+WIDGET_PADDING*5)
	win.SetLabel("Timer")

	vpack := fltk.NewPack(WIDGET_PADDING, WIDGET_PADDING, win.W()-WIDGET_PADDING*2, win.H()-WIDGET_PADDING*2)
	vpack.SetSpacing(WIDGET_PADDING)

	hpack := fltk.NewPack(0, 0, vpack.W(), WIDGET_HEIGHT)
	hpack.SetType(fltk.HORIZONTAL)
	fltk.NewBox(fltk.NO_BOX, 0, 0, WIDGET_LABEL_WIDTH, WIDGET_HEIGHT, "Elapsed Time:").SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	elapsedProgess := fltk.NewProgress(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	elapsedProgess.SetSelectionColor(fltk.BLUE)
	elapsedProgess.SetMaximum(DURATION_DEFAULT)
	hpack.End()

	elapsedLabel := fltk.NewBox(fltk.NO_BOX, 0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	elapsedLabel.SetLabel("0.0s")
	elapsedLabel.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)

	hpack = fltk.NewPack(0, 0, vpack.W(), WIDGET_HEIGHT)
	hpack.SetType(fltk.HORIZONTAL)
	fltk.NewBox(fltk.NO_BOX, 0, 0, WIDGET_LABEL_WIDTH, WIDGET_HEIGHT, "Duration:").SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	durationSlider := fltk.NewSlider(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	durationSlider.SetType(fltk.HOR_SLIDER)
	durationSlider.SetValue(DURATION_DEFAULT)
	durationSlider.SetMaximum(DURATION_MAXIMUM)
	// durationSlider.SetCallbackCondition(fltk.WhenChanged)
	durationSlider.SetCallback(func() {
		// log.Printf("change duration: %f", durationSlider.Value())
		elapsedProgess.SetMaximum(durationSlider.Value())
	})
	hpack.End()

	resetBtn := fltk.NewButton(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	resetBtn.SetLabel("Reset")
	resetBtn.SetCallback(func() {
		startTime = time.Now()
		elapsedProgess.SetValue(0.0)
	})

	stopCh := make(chan struct{})
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		for {
			select {
			case <-stopCh:
				return
			case t := <-ticker.C:
				fltk.Awake(func() {
					if durationSlider.Value()-elapsedProgess.Value() >= 0 {
						d := t.Sub(startTime)
						elapsedProgess.SetValue(d.Seconds())
						elapsedLabel.SetLabel(d.String())
					}
				})
			}
		}
	}()

	win.End()
	win.Show()
	fltk.Run()

	close(stopCh)
}
