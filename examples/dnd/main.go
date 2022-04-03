package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

type Sender struct {
	*fltk.Box
}
func (s *Sender) OnEvent(event fltk.Event) bool {
	switch (event) {
		case fltk.PUSH:
		fltk.CopyToSelectionBuffer("It works!")
		fltk.DragAndDrop()
		return true
	}
	return false
}
func NewSender(x, y, w, h int) *Sender {
	s := &Sender {}
	s.Box = fltk.NewBox(fltk.FLAT_BOX, x, y, w, h, "Drag\nfrom\nhere..")
	s.Box.SetEventHandler(s.OnEvent)
	return s
}

type Receiver struct {
	*fltk.Box
}
func (r *Receiver) OnEvent(event fltk.Event) bool {
	switch (event) {
	case fltk.DND_ENTER,fltk.DND_DRAG,fltk.DND_RELEASE:
		return true
	case fltk.PASTE:
		r.SetLabel(fltk.EventText())
		fmt.Println("Pasted '", fltk.EventText(), "'\n")
		return true
	}
	return false
}
func NewReceiver(x, y, w, h int) *Receiver {
	r := &Receiver{}
	r.Box = fltk.NewBox(fltk.FLAT_BOX, x, y, w, h, "..to\nhere")
	r.SetEventHandler(r.OnEvent)
	return r
}

type SenderWindow struct {
	*fltk.Window
	sender *Sender
}
func NewSenderWindow() *SenderWindow {
	w := &SenderWindow{}
	w.Window = fltk.NewWindow(200, 100)
	w.SetPosition(0, 0)
	w.SetLabel("Sender")
	w.sender = NewSender(0, 0, 100, 100)
	w.End()
	return w
}
type ReceiverWindow struct {
	*fltk.Window
	receiver *Receiver
}
func NewReceiverWindow() *ReceiverWindow {
	w := &ReceiverWindow{}
	w.Window = fltk.NewWindow(200, 100)
	w.SetPosition(400, 0)
	w.SetLabel("Receiver")
	w.receiver = NewReceiver(100, 0, 100, 100)
	w.End()
	return w
}
func main() {
	win_a := NewSenderWindow() 
	win_a.Show()
	win_b := NewReceiverWindow()
	win_b.Show()
	fltk.Run()
}
