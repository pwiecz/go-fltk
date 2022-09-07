package main

import (
	"bytes"

	_ "embed"

	"github.com/pwiecz/go-fltk"
	"github.com/yuin/goldmark"
)

//go:embed example.md
var exampleText string

const (
	WIDGET_HEIGHT  = 600
	WIDGET_PADDING = 10
	WIDGET_WIDTH   = 400
)

func main() {
	fltk.SetScheme("gtk+")

	win := fltk.NewWindow(
		WIDGET_WIDTH*2+WIDGET_PADDING*3,
		WIDGET_HEIGHT*1+WIDGET_PADDING*2)
	win.SetLabel("MDViewer")
	win.Resizable(win)

	hpack := fltk.NewPack(WIDGET_PADDING, WIDGET_PADDING, win.W(), WIDGET_HEIGHT)
	hpack.SetType(fltk.HORIZONTAL)
	hpack.SetSpacing(WIDGET_PADDING)

	mdEditorBuf := fltk.NewTextBuffer()
	mdEditorBuf.SetText(exampleText)
	mdEditor := fltk.NewTextEditor(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	mdEditor.SetBuffer(mdEditorBuf)

	previewPanel := fltk.NewHelpView(0, 0, WIDGET_WIDTH, WIDGET_HEIGHT)
	updatePreview(mdEditorBuf, previewPanel)

	mdEditor.SetCallbackCondition(fltk.WhenChanged)
	mdEditor.SetCallback(func() {
		updatePreview(mdEditorBuf, previewPanel)
	})

	win.End()
	win.Show()
	fltk.Run()
}

func updatePreview(mdEditorBuf *fltk.TextBuffer, previewPanel *fltk.HelpView) {
	var buf bytes.Buffer
	source := []byte(mdEditorBuf.Text())
	if err := goldmark.Convert(source, &buf); err != nil {
		return
	}

	previewPanel.SetValue(buf.String())
}
