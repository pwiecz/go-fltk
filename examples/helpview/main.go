package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 135)
	helpView := fltk.NewHelpView(5, 5, 290, 100)
	helpContent := "List:<ul>"
	for i := 0; i < 100; i++ {
		helpContent += fmt.Sprintf("<li>%d</li>", i)
	}
	helpContent += "</ul>The end"
	helpView.SetValue(helpContent)
	scrollToTop := fltk.NewButton(5, 110, 65, 20, "Top")
	scrollToBottom := fltk.NewButton(155, 110, 65, 20, "Bottom")
	scrollToBottom.SetCallback(func() {
		fmt.Println("Scrolling to bottom")
		helpView.SetTopLine(1000000)
		helpView.SetTopLine(helpView.TopLine() - helpView.H())
	})
	scrollToTop.SetCallback(func() {
		fmt.Println("Scrolling to top")
		helpView.SetTopLine(0)
	})
	win.End()
	win.Show()
	fltk.Run()
}
