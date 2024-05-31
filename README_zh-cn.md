<!-- TOC -->

- [1. go-fltk](#1-go-fltk)
	- [1.1. 依赖](#11-依赖)
	- [1.2. 使用](#12-使用)
	- [1.3. 样式](#13-样式)
	- [1.4. 图像支持](#14-图像支持)
	- [1.5. 资源](#15-资源)

<!-- /TOC -->

# 1. go-fltk

---
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/pwiecz/go-fltk)

* [Document](./README.md)
* [中文文档](./README_zh-cn.md)

一个围绕 FLTK 1.4 库的简单封装，FLTK 是一个轻量级 GUI 库，允许创建小型、独立且快速的 GUI 应用程序。

## 1.1. 依赖
要构建 go-fltk，除了 Golang 编译器，你还需要一个 C++11 编译器，例如 Linux 上的 GCC 或 Clang，Windows 上的 MinGW 和 MacOS 上的 XCode。

go-fltk 带有一些架构的预构建 FLTK 库（linux/amd64, windows/amd64），但你也可以轻松地自己重建它们，或者为其他架构构建它们。
要为你的平台构建 FLTK 库，只需从 go-fltk 源代码树的根目录运行 go generate。

如果构建过程对你不起作用，你可以自行修改 fltk-build.sh 或 fltk-build.bat，或者在 https://github.com/pwiecz/go-fltk/discussions 提出问题。

要运行使用 go-fltk 构建的程序，你将需要一些系统库，这些库通常在带有图形用户界面的操作系统上是可用的：

- Windows: 除了 mingw64 没有外部依赖
- MacOS: 没有外部依赖
- Linux（和其他未测试的 Unix 系统）: 你需要：
    - X11
    - Xrender
    - Xcursor
    - Xfixes
    - Xext
    - Xft
    - Xinerama
    - OpenGL

## 1.2. 使用
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

可以使用 `fltk.New<WidgetType>` 函数创建小部件，并对你正在实例化的小部件进行修改。
函数和方法名与原始 C++ 名称相似，但遵循 Go 语言的 PascalCase 命名习惯。
设置器方法前缀为 `Set`。

## 1.3. 样式
FLTK 提供了 4 种内置样式：
- base (默认)
- gtk+
- gleam
- plastic
例如：可以使用 `fltk.SetScheme("gtk+")` 来设置这些样式.

FLTK 还允许自定义小部件的样式：
```go
package main

import (
	"strconv"

	"github.com/pwiecz/go-fltk"
)

// FLTK 使用 RGBI 颜色表示法，其中 I 是 FLTK 颜色表的索引
// 传递 00 作为 I 将使用 RGB 值
const GRAY = 0x75757500
const LIGHT_GRAY = 0xeeeeee00
const BLUE = 0x42A5F500
const SEL_BLUE = 0x2196F300
const WIDTH = 600
const HEIGHT = 400

func main() {
	curr := 0
	fltk.InitStyles()
	win := fltk.NewWindow(WIDTH, HEIGHT)
	win.SetLabel("Flutter-like")
	win.SetColor(fltk.WHITE)
	bar := fltk.NewBox(fltk.FLAT_BOX, 0, 0, WIDTH, 60, "    FLTK App!")
	bar.SetDrawHandler(func() { // Shadow under the bar
		fltk.DrawBox(fltk.FLAT_BOX, 0, 0, WIDTH, 63, LIGHT_GRAY)
	})
	bar.SetAlign(fltk.ALIGN_INSIDE | fltk.ALIGN_LEFT)
	bar.SetLabelColor(255) // this uses the index into the color map, here it's white
	bar.SetColor(BLUE)
	bar.SetLabelSize(22)
	text := fltk.NewBox(fltk.NO_BOX, 250, 180, 100, 40, "You have pushed the button this many times:")
	text.SetLabelSize(18)
	text.SetLabelFont(fltk.TIMES)
	count := fltk.NewBox(fltk.NO_BOX, 250, 180+40, 100, 40, "0")
	count.SetLabelSize(36)
	count.SetLabelColor(GRAY)
	btn := fltk.NewButton(WIDTH-100, HEIGHT-100, 60, 60, "@+6plus") // 这翻译成一个加号
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

标签属性可以在[这里](https://www.fltk.org/doc-1.3/common.html#common_labels)查看

## 1.4. 图像支持
FLTK 支持矢量和光栅图形，通过几种图像类型：
- SvgImage
- RgbImage
- JpegImage
- PngImage
- BmpImage
- SharedImage

其中一些可以从图像文件或数据中实例化：
```go
package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(400, 300)
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 400, 300, "")
	image, err := fltk.NewJpegImageLoad("image.jpg")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		box.SetImage(image)
	}
	win.End()
	win.Show()
	fltk.Run()
}
```

## 1.5. 资源
- [官方 FLTK 1.4 文档](https://www.fltk.org/doc-1.4/index.html)
- [go-fltk 文档](https://pkg.go.dev/github.com/pwiecz/go-fltk) 