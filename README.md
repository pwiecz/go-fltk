# go-fltk

A simple wrapper around FLTK 1.4 library

## Usage
```go
package main

import "github.com/pwiecz/go-fltk"

func main() {
    win := fltk.NewWindow(400, 300)
    win.SetLabel("Main Window")
    btn := fltk.NewButton(160, 200, 80, 30, "Click")
    btn.SetCallback(func() {
        btn.SetLabel("Clicked")
    })
    win.End()
    win.Show()
    fltk.Run()
}
```

Widgets are created using the `fltk.New<WidgetType>` functions, modified for whatever widget you're instantiating.
Function and method names resemble the original C++ names, while however, following the Golang convention of PascalCase. 