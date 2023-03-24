package main

import (
	"fmt"
	"runtime"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	group := fltk.NewGroup(5, 20, 285, 170, "Group")
	fltk.NewButton(10, 10, 50, 20, "Button")
	group.End()
	{
		button := group.Child(0)
		button.SetCallback(func() {
			fmt.Println("Button pressed")
			// Even though the go button object gets GC'd the callback is still being called.
			// We may say that it's a memory leak, but it's hard to prevent it...
			runtime.GC()
		})
	}
	win.End()
	win.Show()
	fltk.Run()
}
