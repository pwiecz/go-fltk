package main

import (
	"fmt"

	"github.com/pwiecz/go-fltk"
)

func main() {
	win := fltk.NewWindow(300, 200)
	table := fltk.NewTableRow(5, 5, 295, 190)
	table.EnableColumnHeaders()
	table.EnableRowHeaders()
	table.SetColumnCount(3)
	table.SetRowCount(4)
	table.SetDrawCellCallback(func(tc fltk.TableContext, row, col, x, y, w, h int) {
		if tc == fltk.ContextCell {
			fltk.SetDrawFont(fltk.HELVETICA, 14)
			fltk.DrawBox(fltk.FLAT_BOX, x, y, w, h, fltk.BLACK)
			fltk.DrawBox(fltk.FLAT_BOX, x+1, y+1, w-2, h-2, fltk.WHITE)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(fmt.Sprintf("%d", row+col), x, y, w, h, fltk.ALIGN_CENTER)
		}
		if tc == fltk.ContextRowHeader {
			fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
			fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(fmt.Sprintf("row %d", row+1), x, y, w, h, fltk.ALIGN_CENTER)
		}
		if tc == fltk.ContextColHeader {
			fltk.SetDrawFont(fltk.HELVETICA_BOLD, 14)
			fltk.DrawBox(fltk.UP_BOX, x, y, w, h, fltk.BACKGROUND_COLOR)
			fltk.SetDrawColor(fltk.BLACK)
			fltk.Draw(fmt.Sprintf("col %d", col+1), x, y, w, h, fltk.ALIGN_CENTER)
		}
	})
	table.End()
	win.End()
	win.Show()
	fltk.Run()
}
