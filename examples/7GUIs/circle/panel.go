package main

import (
	"fmt"
	"runtime"

	"github.com/pwiecz/go-fltk"
)

type Panel struct {
	win          *fltk.Window
	undoBtn      *fltk.Button
	redoBtn      *fltk.Button
	drawBox      *fltk.Box
	adjustDlg    *fltk.Window
	adjustTips   *fltk.Box
	adjustSlider *fltk.Slider
	popMenu      *fltk.MenuButton
}

func NewPanel(win *fltk.Window) *Panel {
	p := &Panel{}
	p.win = win

	col := fltk.NewFlex(WIDGET_PADDING, WIDGET_PADDING, win.W()-WIDGET_PADDING*2, win.H()-WIDGET_PADDING*2)
	col.SetGap(WIDGET_PADDING)

	row := fltk.NewFlex(0, 0, 0, 0)
	col.Fixed(row, WIDGET_HEIGHT)
	row.SetType(fltk.ROW)
	row.SetGap(WIDGET_PADDING)

	fltk.NewBox(fltk.NO_BOX, 0, 0, 0, 0) // invisible

	p.undoBtn = fltk.NewButton(0, 0, 0, 0)
	p.undoBtn.SetLabel("Undo")

	p.redoBtn = fltk.NewButton(0, 0, 0, 0)
	p.redoBtn.SetLabel("Redo")

	fltk.NewBox(fltk.NO_BOX, 0, 0, 0, 0) // invisible

	row.End()

	p.drawBox = fltk.NewBox(fltk.NO_BOX, 0, 0, 0, 0)

	col.End()
	win.Resizable(col)

	p.adjustDlg = fltk.NewWindow(WIDGET_WIDTH*2+WIDGET_PADDING*2, WIDGET_HEIGHT*2+WIDGET_PADDING*3)
	p.adjustDlg.SetModal()

	row = fltk.NewFlex(WIDGET_PADDING, WIDGET_PADDING, WIDGET_WIDTH*2, WIDGET_HEIGHT*2)
	p.adjustTips = fltk.NewBox(fltk.NO_BOX, 0, 0, 0, 0)
	p.adjustTips.SetLabel("Adjust diameter")
	p.adjustSlider = fltk.NewSlider(0, 0, 0, 0)
	p.adjustSlider.SetType(fltk.HOR_NICE_SLIDER)
	p.adjustSlider.SetMaximum(MAX_RADIUS)
	p.adjustSlider.SetMinimum(MIN_RADIUS)
	p.adjustSlider.SetValue(DEF_RADIUS)
	row.End()

	// SetModal makes the dialog's close button disappear on Windows
	// A workaround is to make the dialog resizable
	if runtime.GOOS == "windows" {
		p.adjustDlg.Resizable(row)
	}

	p.adjustDlg.End()

	p.popMenu = fltk.NewMenuButton(0, 0, 0, 0)
	p.popMenu.SetType(fltk.POPUP2)
	p.popMenu.Add("Adjust diameter..", func() {
		p.adjustDlg.Show()
	})

	return p
}

func (p *Panel) Bind(ctx *Context) {
	p.undoBtn.SetCallback(func() {
		if ctx.Undo() {
			p.Update(ctx)
		}
	})

	p.redoBtn.SetCallback(func() {
		if ctx.Redo() {
			p.Update(ctx)
		}
	})

	p.drawBox.SetDrawHandler(func() {
		fltk.DrawRectfWithColor(p.drawBox.X(), p.drawBox.Y(), p.drawBox.W(), p.drawBox.H(), fltk.WHITE)

		for _, c := range ctx.Circles() {
			x := c.X + p.drawBox.X() - c.R
			y := c.Y + p.drawBox.Y() - c.R
			w := c.R * 2
			h := c.R * 2

			if ctx.IsCirclePicked(c) {
				fltk.SetDrawColor(fltk.ColorFromRgb(128, 128, 128))
				fltk.DrawPie(x, y, w, h, 0, 360)
			} else {
				fltk.SetDrawColor(fltk.ColorFromRgb(0, 0, 0))
				fltk.DrawArc(x, y, w, h, 0, 360)
			}
		}
	})

	p.drawBox.SetEventHandler(func(e fltk.Event) bool {
		if fltk.EventIsClick() && e == fltk.RELEASE {
			x := (fltk.EventX() - p.drawBox.X())
			y := (fltk.EventY() - p.drawBox.Y())
			switch fltk.EventButton() {
			case fltk.LeftMouse:
				c := ctx.NewCircle(x, y)
				ctx.AddCircle(c)
				ctx.AddOp(OP_ADD, c)
				p.Update(ctx)
				return true
			case fltk.RightMouse:
				c := ctx.PickCircle(x, y)
				if c == nil {
					return true
				}
				p.Update(ctx)

				p.adjustTips.SetLabel(fmt.Sprintf("Adjust the circle at (%d, %d)", c.X, c.Y))
				p.adjustSlider.SetValue(float64(c.R))
				p.adjustDlg.SetPosition(fltk.EventXRoot()-c.R, fltk.EventYRoot()-c.R-p.adjustDlg.H()-WIDGET_PADDING)
				p.popMenu.Popup()
			}
		}
		return false
	})

	p.adjustDlg.SetEventHandler(func(e fltk.Event) bool {
		c := ctx.PickedCircle()
		if c == nil {
			return false
		}
		switch e {
		case fltk.SHOW:
			ctx.UpdateCircle(c)
			ctx.AddOp(OP_UPDATE, c)
			return true
		case fltk.HIDE:
			c.R = int(p.adjustSlider.Value())
			ctx.AddOp(OP_UPDATE, c)
			ctx.UnPickCircle()
			p.Update(ctx)
			return true
		}
		return false
	})

	p.adjustSlider.SetCallbackCondition(fltk.WhenChanged)
	p.adjustSlider.SetCallback(func() {
		c := ctx.PickedCircle()
		if c == nil {
			return
		}

		c.R = int(p.adjustSlider.Value())
		ctx.UpdateCircle(c)
		p.Update(ctx)
	})

	p.Update(ctx)
}

func (p *Panel) Update(ctx *Context) {
	p.drawBox.Redraw()

	if ctx.HasRedo() {
		p.redoBtn.Activate()
	} else {
		p.redoBtn.Deactivate()
	}

	if ctx.HasUndo() {
		p.undoBtn.Activate()
	} else {
		p.undoBtn.Deactivate()
	}
}
