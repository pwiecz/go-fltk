package fltk

import (
	"errors"
	"testing"
)

func testWidgetDestroyed(w *widget, t *testing.T) {
	if w.tracker != nil {
		t.Errorf("widget's tracker is not nil")
	}
	if w.callbackId != 0 {
		t.Errorf("callbackId is not 0")
	}
	if w.deletionHandlerId != 0 {
		t.Errorf("deletionHandlerId is not 0")
	}
	if w.resizeHandlerId != 0 {
		t.Errorf("resizeHandlerId is not 0")
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
		} else {
			err := r.(error)
			if err == nil {
				t.Errorf("Panicked with not an error")
			} else {
				if !errors.Is(err, ErrDestroyed) {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}
		testWidgetDestroyed(&b.widget, t)
		testGlobalMapsEmpty(t)
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		go Awake(func() {
			b.Destroy()
			go Awake(func() {
				b.SetLabel("bar")
			})
		})
		return true
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
		} else {
			err := r.(error)
			if err == nil {
				t.Errorf("Panicked with not an error")
			} else {
				if !errors.Is(err, ErrDestroyed) {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}
		testWidgetDestroyed(&g.widget, t)
		testWidgetDestroyed(&b.widget, t)
		testGlobalMapsEmpty(t)
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		go Awake(func() {
			g.Destroy()
			go Awake(func() {
				b.SetLabel("bar")
			})
		})
		return true
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
		} else {
			err := r.(error)
			if err == nil {
				t.Errorf("Panicked with not an error")
			} else {
				if !errors.Is(err, ErrDestroyed) {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}
		testWidgetDestroyed(&g.widget, t)
		testWidgetDestroyed(&b.widget, t)
		testWidgetDestroyed(&bParent.widget, t)
		testGlobalMapsEmpty(t)
	}()
	b.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		go Awake(func() {
			bParent = b.Parent()
			bParent.Destroy()
			go Awake(func() {
				b.SetLabel("bar")
			})
		})
		return true
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
		} else {
			err := r.(error)
			if err == nil {
				t.Errorf("Panicked with not an error")
			} else {
				if !errors.Is(err, ErrDestroyed) {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}
		testWidgetDestroyed(&tb.widget, t)
		testGlobalMapsEmpty(t)
	}()
	tb.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		go Awake(func() {
			tb.widget.Destroy()
			go Awake(func() {
				tb.IsRowSelected(0)
			})
		})
		return true
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
		} else {
			err := r.(error)
			if err == nil {
				t.Errorf("Panicked with not an error")
			} else {
				if !errors.Is(err, ErrDestroyed) {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}
		testWidgetDestroyed(&mb.widget, t)
		testGlobalMapsEmpty(t)
	}()
	mb.SetEventHandler(func(event Event) bool {
		if event != SHOW {
			return false
		}
		go Awake(func() {
			mb.widget.Destroy()
			go Awake(func() {
				mb.Redraw()
			})
		})
		return true
	})
	win.End()
	Lock()
	win.Show()
	Run()
}
