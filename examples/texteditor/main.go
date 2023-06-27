package main

import (
	"fmt"
	"os"

	"github.com/pwiecz/go-fltk"
)

const (
	WIDGET_HEIGHT  = 400
	WIDGET_PADDING = 0
	WIDGET_WIDTH   = 800
)

type EditorApp struct {
	Win        *fltk.Window
	TextBuffer *fltk.TextBuffer
	TextEditor *fltk.TextEditor
	FileName   string
	IsChanged  bool
}

func (app *EditorApp) BuildGUI() {
	fltk.SetScheme("gtk+")

	app.Win = fltk.NewWindow(WIDGET_WIDTH, WIDGET_HEIGHT)

	app.Win.SetLabel("TextEditor")
	app.Win.Resizable(app.Win)

	hpack := fltk.NewFlex(WIDGET_PADDING, WIDGET_PADDING, app.Win.W(), WIDGET_HEIGHT)
	hpack.SetType(fltk.COLUMN)
	hpack.SetSpacing(WIDGET_PADDING)

	menuBar := fltk.NewMenuBar(0, 0, 0, 0)
	hpack.Fixed(menuBar, 20)
	menuBar.SetType(uint8(fltk.FLAT_BOX))
	menuBar.Activate()
	menuBar.AddEx("File", fltk.ALT+'f', nil, fltk.SUBMENU)
	menuBar.AddEx("File/&Open", fltk.CTRL+'o', app.callbackMenuFileOpen, fltk.MENU_VALUE)
	menuBar.AddEx("File/&Save", fltk.CTRL+'s', app.callbackMenuFileSave, fltk.MENU_VALUE)
	menuBar.Add("File/Save &As", app.callbackMenuFileSaveAs)
	menuBar.AddEx("Help", 0, nil, fltk.SUBMENU)
	menuBar.Add("Help/&About", app.callbackMenuHelpAbout)

	app.TextBuffer = fltk.NewTextBuffer()
	app.TextEditor = fltk.NewTextEditor(0, 0, 0, 0)
	app.TextEditor.SetBuffer(app.TextBuffer)

	app.TextEditor.SetCallbackCondition(fltk.WhenChanged)
	app.TextEditor.SetCallback(func() {
		app.IsChanged = true
	})
	app.TextEditor.Parent().Resizable(app.TextEditor)
	hpack.End()
	app.Win.End()
	app.IsChanged = false
}

func main() {
	myapp := EditorApp{}
	myapp.BuildGUI()
	myapp.Win.Show()
	fltk.Run()
}

func (app *EditorApp) callbackMenuFileOpen() {
	fChooser := fltk.NewFileChooser("./", "*.*", fltk.FileChooser_SINGLE, "Open text file")
	defer fChooser.Destroy()
	fChooser.Popup()
	fnames := fChooser.Selection()
	if len(fnames) == 1 {
		fmt.Printf("select file name %s\n", fnames[0])
		textByte, err := os.ReadFile(fnames[0])
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		app.TextBuffer.SetText(string(textByte))
		app.FileName = fnames[0]
	}
}

func (app *EditorApp) callbackMenuFileSave() {
	if app.IsChanged {
		info, _ := os.Stat(app.FileName)
		os.WriteFile(app.FileName, []byte(app.TextBuffer.Text()), info.Mode())
		app.IsChanged = false
	}
}

func (app *EditorApp) callbackMenuFileSaveAs() {
	fChooser := fltk.NewFileChooser("./", "*.*", fltk.FileChooser_CREATE, "Enter/Select file name")
	defer fChooser.Destroy()
	fChooser.Popup()
	fnames := fChooser.Selection()
	if len(fnames) == 1 {
		os.WriteFile(fnames[0], []byte(app.TextBuffer.Text()), 0640)
		app.IsChanged = false
		app.FileName = fnames[0]
	}
}

func (app *EditorApp) callbackMenuHelpAbout() {
	fltk.MessageBox("About", "Sample Text Editor")
}
