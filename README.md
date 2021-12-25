# go-fltk

A simple wrapper around FLTK 1.4 library, which is a lightweight GUI library which allows creating small, self-contained and fast gui applications.

## Requirements
In addition to Go, you will also need a C++11 compiler. The FLTK libraries are bundled with the repo for x86_64 Linux, MacOS and Windows (mingw64).
You also need system libs which are normally available on operating systems with graphical user interfaces:
- Windows: No external dependencies (for mingw64)
- MacOS: No external dependencies
- Linux: You need the X11 libs:
    - x11
    - Xrender
    - Xcursor
    - Xfixes
    - Xext
    - Xft
    - Xinerama
    - OpenGL

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
Setter methods are also preceded by a `Set` prefix.

## Styling
FLTk offers 4 builtin schemes:
- base (default)
- gtk+
- gleam
- plastic
These can be set using `fltk.SetScheme("gtk+")` for example.

FLTK also allows custom styling your widgets:
```go
package main

import (
    "github.com/pwiecz/go-fltk"
    "strconv"
)

const WIDTH = 600
const HEIGHT = 400
// FLTK uses an RGBI color representation, the I is an index into FLTK's color map
// Passing 00 as I will use the RGB part of the value
const GRAY = 0x75757500
const BLUE = 0x42A5F500
const SEL_BLUE = 0x2196F300

func main() {
    curr := 0
    fltk.InitStyles()
    win := fltk.NewWindow(WIDTH, HEIGHT)
    win.SetLabel("Flutter-like")
    win.SetColor(0xffffff00) 
    bar := fltk.NewBox(fltk.FLAT_BOX, 0, 0, WIDTH, 60, "    FLTK App!")
    bar.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
    bar.SetLabelColor(255) // this uses the index into the color map
    bar.SetColor(BLUE)
    bar.SetLabelSize(22)
    text := fltk.NewBox(fltk.NO_BOX, 250, 180, 100, 40, "You have pushed the button this many times:")
    text.SetLabelSize(18)
    text.SetLabelFont(fltk.TIMES)
    count := fltk.NewBox(fltk.NO_BOX, 250, 180+40, 100, 40, "0")
    count.SetLabelSize(36)
    count.SetLabelColor(GRAY)
    btn := fltk.NewButton(WIDTH-100, HEIGHT-100, 60, 60, "@+6plus") // this translates into a plus sign
    btn.SetColor(BLUE)
    btn.SetSelectionColor(SEL_BLUE)
    btn.SetLabelColor(255)
    btn.SetBox(fltk.OFLAT_BOX)
    btn.ClearVisibleFocus()
    btn.SetCallback(func() {
        curr += 1
        count.SetLabel(strconv.Itoa(curr))
    })
    win.End()
    win.Show()
    fltk.Run()
}
```

![image](https://user-images.githubusercontent.com/37966791/147374840-2d993522-fc86-46fc-9e95-2b3391d31013.png)

The label attributes can be seen [here](https://www.fltk.org/doc-1.3/common.html#common_labels)