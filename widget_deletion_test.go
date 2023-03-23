package fltk

import (
	"errors"
	"testing"
)

func testWidgetDestroyed(name string, w *widget, t *testing.T) {
	if w.tracker != nil {
		t.Errorf("%s's widget's tracker is not nil", name)
	}
	if w.callbackId != 0 {
		t.Errorf("%s's callbackId is not 0", name)
	}
	if w.deletionHandlerId != 0 {
		t.Errorf("%s's deletionHandlerId is not 0", name)
	}
	if w.resizeHandlerId != 0 {
		t.Errorf("%s's resizeHandlerId is not 0", name)
	}
}

func testGlobalMapsEmpty(t *testing.T) {
	// actually in our tests we do not destroy the main windows, so the callback map should
	// contain their deletion handlers.
	if globalCallbackMap.size() != 1 {
		t.Errorf("global callback map is not empty: %d", globalCallbackMap.size())
	}
	globalCallbackMap.clear()
	if !globalEventHandlerMap.isEmpty() {
		t.Errorf("global event handler map is not empty: %d", globalEventHandlerMap.size())
	}
	globalEventHandlerMap.clear()
	if !globalTableCallbackMap.isEmpty() {
		t.Errorf("global table callback map is not empty: %d", globalTableCallbackMap.size())
	}
	globalTableCallbackMap.clear()
}

func TestPanicWhenAccessingDeletedWidget(t *testing.T) {
	win := NewWindow(400, 400)
	b := NewButton(2, 2, 50, 50, "foo")
	b.SetResizeHandler(func() {})
	win.End()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrDestroyed) {
			t.Errorf("Unexpected error: %v", err)
		}
		testWidgetDestroyed("button", &b.widget, t)
		testGlobalMapsEmpty(t)
		Unlock()
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		b.Destroy()
		Wait()
		b.SetLabel("bar")
		panic("Should have panicked")
	})
	Lock()
	win.Show()
	Run()
}

func TestPanicWhenAccessingChildOfDeletedWidget(t *testing.T) {
	win := NewWindow(400, 400)
	g := NewGroup(1, 1, 398, 398)
	g.SetResizeHandler(func() {})
	b := NewButton(2, 2, 50, 50, "foo")
	b.SetResizeHandler(func() {})
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrDestroyed) {
			t.Errorf("Unexpected error: %v", err)
		}
		testWidgetDestroyed("group", &g.widget, t)
		testWidgetDestroyed("button", &b.widget, t)
		testGlobalMapsEmpty(t)
		Unlock()
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		g.Destroy()
		Wait()
		b.SetLabel("bar")
		panic("Should have panicked")
	})
	g.End()
	win.End()
	Lock()
	win.Show()
	Run()
}

func TestPanicWhenAccessingChildOfWidgetDeletedViaParent(t *testing.T) {
	win := NewWindow(400, 400)
	g := NewGroup(1, 1, 398, 398)
	g.SetResizeHandler(func() {})
	b := NewButton(2, 2, 50, 50, "foo")
	b.SetResizeHandler(func() {})
	var bParent *Group
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrDestroyed) {
			t.Errorf("Unexpected error: %v", err)
		}
		testWidgetDestroyed("group", &g.widget, t)
		testWidgetDestroyed("button", &b.widget, t)
		testGlobalMapsEmpty(t)
		Unlock()
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		bParent = b.Parent()
		bParent.Destroy()
		Wait()
		b.SetLabel("bar")
		panic("Should have panicked")
	})
	g.End()
	win.End()
	Lock()
	win.Show()
	Run()
}

// TableRow uses custom cleanup procedure
func TestDestroyingTableRow(t *testing.T) {
	win := NewWindow(400, 400)
	tb := NewTableRow(20, 20, 50, 50)
	tb.SetResizeHandler(func() {})
	tb.SetDrawCellCallback(nil)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrDestroyed) {
			t.Errorf("Unexpected error: %v", err)
		}
		testWidgetDestroyed("table row", &tb.widget, t)
		testGlobalMapsEmpty(t)
		Unlock()
	}()
	tb.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		tb.Destroy()
		Wait()
		tb.IsRowSelected(0)
		panic("Should have panicked")
	})
	win.End()
	Lock()
	win.Show()
	Run()
}

// MenuBar uses custom cleanup procedure and may have to clear many callbacks
func TestDestroyingMenu(t *testing.T) {
	win := NewWindow(400, 400)
	mb := NewMenuBar(20, 20, 50, 50)
	mb.SetResizeHandler(func() {})
	mb.Add("&File/&Load", func() {})
	mb.AddEx("&File/&Save", CTRL+int('s'), func() {}, 0)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Did not panic")
		} else if err, ok := r.(error); !ok {
			t.Errorf("Panicked with not an error: %v", r)
		} else if !errors.Is(err, ErrDestroyed) {
			t.Errorf("Unexpected error: %v", err)
		}
		testWidgetDestroyed("menu bar", &mb.widget, t)
		testGlobalMapsEmpty(t)
		Unlock()
	}()
	mb.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		mb.Destroy()
		Wait()
		mb.Redraw()
		panic("Should have panicked")
	})
	win.End()
	Lock()
	win.Show()
	Run()
}
