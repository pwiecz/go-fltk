package fltk

import (
	"errors"
	"testing"
)

func TestPanicWhenTestBufferIsMissing(t *testing.T) {
	win := NewWindow(400, 400)
	textEditor := NewTextEditor(2, 2, 300, 300, "")
	win.End()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrNoTextBufferAssociated) {
			t.Errorf("Unexpected error: %v", err)
		}
		textEditor.Destroy()
		Unlock()
		// clean up after outselves as other tests check if this map is empty
		globalCallbackMap.clear()
	}()
	textEditor.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		textEditor.SelectAll()
		panic("Should have panicked")
	})
	Lock()
	win.Show()
	Run()
}
