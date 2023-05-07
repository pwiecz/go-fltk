package main

import (
	"log"

	"github.com/pwiecz/go-fltk"
)

type Panel struct {
	tb         *fltk.TableRow
	cellValues map[CellLoc]string

	editInput *fltk.Input // the input box to show on editing cell
	editCell  *Cell       // current editing cell meta, nil means not editing
}

func NewPanel(win *fltk.Window, rowCount, colCount int) *Panel {
	p := &Panel{}

	p.cellValues = make(map[CellLoc]string)

	p.tb = fltk.NewTableRow(0, 0, win.W(), win.H())
	p.tb.SetRowCount(rowCount)
	p.tb.SetColumnCount(colCount)
	p.tb.EnableColumnHeaders()
	p.tb.EnableRowHeaders()
	p.tb.AllowColumnResizing()
	p.tb.AllowRowResizing()

	p.tb.Begin()

	p.editInput = fltk.NewInput(0, 0, 0, 0)
	p.editInput.Hide()
	p.editInput.SetColor(fltk.YELLOW)

	p.tb.End()
	win.Resizable(p.tb)

	return p
}

func (p *Panel) Bind(ctx *Context) {
	for row := range ctx.Cells {
		for col := range ctx.Cells[row] {
			cell := ctx.Cells[row][col]
			p.cellValues[cell.Loc] = cell.Data.Display()
		}
	}

	p.tb.SetDrawCellCallback(func(tc fltk.TableContext, i, j, x, y, w, h int) {
		row := CellRow(i)
		col := CellCol(j)

		switch tc {
		case fltk.ContextRowHeader:
			fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
			fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(row.String(), x, y, w, h, fltk.ALIGN_CENTER)
		case fltk.ContextColHeader:
			fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
			fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(col.String(), x, y, w, h, fltk.ALIGN_CENTER)
		case fltk.ContextCell:
			loc := CellLoc{Row: row, Col: col}
			if p.IsEditingAt(col, row) {
				p.editInput.Resize(x, y, w, h)
				return
			}
			fltk.SetDrawFont(fltk.HELVETICA, 14)
			fltk.DrawBox(fltk.FLAT_BOX, x, y, w, h, fltk.BLACK)
			fltk.DrawBox(fltk.FLAT_BOX, x+1, y+1, w-2, h-2, fltk.WHITE)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(p.cellValues[loc], x, y, w, h, fltk.ALIGN_CENTER)
		}
	})

	// p.tb.SetCallbackCondition(fltk.WhenNotChanged)
	p.tb.SetCallback(func() {
		tc := p.tb.CallbackContext()
		if tc != fltk.ContextCell {
			p.DoneEditing(ctx)
			return
		}

		if fltk.EventClicks() == 0 {
			p.DoneEditing(ctx)
			return
		}

		p.StartEditing(ctx)
	})

	p.editInput.SetCallbackCondition(fltk.WhenEnterKeyAlways)
	p.editInput.SetCallback(func() {
		p.DoneEditing(ctx)
	})
}

func (p *Panel) IsEditing() bool {
	return p.editCell != nil
}

func (p *Panel) IsEditingAt(col CellCol, row CellRow) bool {
	return p.editCell != nil && p.editCell.Loc.Col == col && p.editCell.Loc.Row == row
}

func (p *Panel) StartEditing(ctx *Context) {
	if p.IsEditing() {
		p.DoneEditing(ctx)
	}

	row := CellRow(p.tb.CallbackRow())
	col := CellCol(p.tb.CallbackColumn())
	loc := CellLoc{Row: row, Col: col}

	x, y, w, h, err := p.tb.FindCell(fltk.ContextCell, int(row), int(col))
	if err != nil {
		log.Panic("should not go here")
		return
	}

	// log.Print("show input:", x, y, w, h)
	p.editCell = ctx.FindCell(loc)
	p.editInput.Resize(x, y, w, h)
	p.editInput.SetValue(p.editCell.RawValue)
	p.editInput.Show()
	p.editInput.TakeFocus()
}

func (p *Panel) DoneEditing(ctx *Context) {
	if p.IsEditing() {
		// log.Print("done editing")
		p.editCell.Update(p.editInput.Value())
		p.ApplyChangedCells(ctx, p.editCell)
		p.editInput.Hide()
		p.editCell = nil
	}
}

func (p *Panel) ApplyChangedCells(ctx *Context, changedCell *Cell) {
	scells := ctx.FindAllChangedCells(changedCell)
	for loc, scell := range scells {
		p.cellValues[loc] = scell.Data.Display()
	}
}
