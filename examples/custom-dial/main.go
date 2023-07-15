package main

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

type MyDial struct {
	main_wid  *fltk.Group
	value     *int
	value_box *fltk.Box
}

func NewMyDial(x, y, w, h int, label string) *MyDial {
	value := 0
	main_wid := fltk.NewGroup(x, y, w, h, label)
	main_wid.SetAlign(fltk.ALIGN_TOP)
	main_wid.SetLabelSize(22)
	main_wid.SetLabelColor(fltk.ColorFromRgb(0x79, 0x79, 0x79))
	value_box :=
		fltk.NewBox(fltk.NO_BOX, main_wid.X(), main_wid.Y()+80, main_wid.W(), 40, "0")
	value_box.SetLabelSize(26)
	main_wid.End()

	main_wid.SetDrawHandler(func(func()) {
		fltk.SetDrawColor(fltk.ColorFromRgb(230, 230, 230))
		fltk.DrawPie(main_wid.X(), main_wid.Y(), main_wid.W(), main_wid.H(), 0., 180.)
		fltk.SetDrawColor(0xb0bf1a00)
		fltk.DrawPie(
			main_wid.X(),
			main_wid.Y(),
			main_wid.W(),
			main_wid.H(),
			float64(100-value)*1.8,
			180.,
		)
		fltk.SetDrawColor(fltk.WHITE)
		fltk.DrawPie(
			main_wid.X()-50+main_wid.W()/2,
			main_wid.Y()-50+main_wid.H()/2,
			100,
			100,
			0.,
			360.,
		)
		main_wid.DrawChildren()
	})
	return &MyDial{
		main_wid,
		&value,
		value_box,
	}
}

func (d *MyDial) SetValue(val int) {
	*d.value = val
	d.value_box.SetLabel(strconv.Itoa(val))
	d.main_wid.Redraw()
}

func (d *MyDial) Value() int {
	return *d.value
}

func main() {
	win := fltk.NewWindow(400, 300)
	win.SetColor(fltk.WHITE)
	dial := NewMyDial(100, 100, 200, 200, "CPU Load %")
	win.End()
	win.Show()
	dial.SetValue(26)
	fltk.Run()
}
