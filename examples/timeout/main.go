package main

import (
	"fmt"
	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(400, 300)

	fltk.NewButton(2, 2, 60, 30, "Test").SetCallback(func() {
		fltk.AddTimeout(2.0, timeoutCb)
	})

	win.End()
	win.Show()
	fltk.Run()
}

func timeoutCb() {
	fmt.Println("test")
	fltk.RepeatTimeout(2.0, timeoutCb)
}
