package fltk

type LayoutCalc struct {
	X, Y, W, H int
}

func NewLayoutCalc(w Widgety) *LayoutCalc {
	return  &LayoutCalc{0, 0, w.GetWidget().W(), w.GetWidget().H()}
}
// add a widget of h height
func (l *LayoutCalc) Add(w Widgety) {
	l.Y += w.GetWidget().H()
	l.H -= w.GetWidget().H()
}
