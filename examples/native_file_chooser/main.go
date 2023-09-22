package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {

	win := fltk.NewWindow(300, 200)
	button := fltk.NewButton(5, 5, 180, 20, "Show file chooser")
	button.SetCallback(func() {
		nfc := fltk.NewNativeFileChooser()
		defer nfc.Destroy()
		nfc.SetOptions(fltk.NativeFileChooser_PREVIEW | fltk.NativeFileChooser_NEW_FOLDER)
		nfc.SetType(fltk.NativeFileChooser_BROWSE_MULTI_FILE)
		nfc.SetDirectory("/home")
		nfc.SetFilter("C++ Files\t*.{cxx,H}\nTxt Files\t*.txt")
		nfc.SetTitle("Native file chooser example")
		nfc.Show()
		fmt.Println("Selected files:")
		for _, filename := range nfc.Filenames() {
			fmt.Println(filename)
		}
	})
	win.End()
	win.Show()
	fltk.Run()
}
