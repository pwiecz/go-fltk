package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var reSumFormula = regexp.MustCompile(`(?i)SUM\(([A-Z]\d+):([A-Z]\d+)\)`)

type Context struct {
	Cells [][]*Cell
}

func NewContext(maxRow, maxCol int) *Context {
	ctx := &Context{}
	ctx.Cells = make([][]*Cell, maxRow)
	for row := range ctx.Cells {
		ctx.Cells[row] = make([]*Cell, maxCol)
		for col := range ctx.Cells[row] {
			cell := &Cell{}
			cell.Loc.Row = CellRow(row)
			cell.Loc.Col = CellCol(col)
			cell.Data = &CellDataText{}
			ctx.Cells[row][col] = cell
		}
	}

	return ctx
}

func (ctx *Context) FindCell(loc CellLoc) *Cell {
	if int(loc.Row) >= len(ctx.Cells) || int(loc.Col) >= len(ctx.Cells[loc.Row]) {
		return nil
	}
	return ctx.Cells[loc.Row][loc.Col]
}

func ParseCellData(value string) CellData {
	if strings.HasPrefix(value, "=") {
		formula := value[1:]
		matches := reSumFormula.FindStringSubmatch(formula)
		if matches == nil {
			return &CellDataInvalid{Reason: "SUPPORT FORMULA"}
		}

		strStartLoc := matches[1]
		startLoc, err := ParseCellLoc(strStartLoc)
		if err != nil {
			return &CellDataInvalid{Reason: "INVALID START CELL"}
		}

		strEndLoc := matches[2]
		endLoc, err := ParseCellLoc(strEndLoc)
		if err != nil {
			return &CellDataInvalid{Reason: "INVALID END CELL"}
		}

		formulaData := &CellDataFormula{}
		formulaData.Formula = formula
		formulaData.CalCells = make(map[CellLoc]*Cell)
		for row := startLoc.Row; row <= endLoc.Row; row++ {
			for col := startLoc.Col; col <= endLoc.Col; col++ {
				loc := CellLoc{Row: row, Col: col}
				formulaData.CalCells[loc] = ctx.FindCell(loc)
			}
		}
		return formulaData
	}

	num, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return &CellDataNumber{Number: num}
	}

	return &CellDataText{Text: value}
}

func (cell *Cell) Update(value string) {
	cellData := ParseCellData(value)
	cell.Data = cellData
	cell.RawValue = value
}

func (ctx *Context) FindAllChangedCells(cell *Cell) map[CellLoc]*Cell {
	changedCells := make(map[CellLoc]*Cell)
	changedCells[cell.Loc] = cell

	ctx.findAllChangedCells(cell, changedCells)

	return changedCells
}

func (ctx *Context) findAllChangedCells(cell *Cell, changedCells map[CellLoc]*Cell) {
	for row := range ctx.Cells {
		for col := range ctx.Cells[row] {
			changedCell := ctx.Cells[row][col]
			if changedCell.Data.IsDependOn(cell) {
				if _, exist := changedCells[changedCell.Loc]; !exist {
					changedCells[changedCell.Loc] = changedCell
					ctx.findAllChangedCells(changedCell, changedCells)
				} else {
					changedCell.Data = &CellDataInvalid{Reason: "RECURSIVE FORMULA"}
				}
			}
		}
	}
}

func (ctx *Context) UpdateCellAtLoc(locStr string, value string) {
	loc, err := ParseCellLoc(locStr)
	if err != nil {
		return
	}
	cell := ctx.FindCell(loc)
	if cell == nil {
		return
	}
	cell.Update(value)
}

type Cell struct {
	Loc      CellLoc
	Data     CellData
	RawValue string
}

type CellRow int

func (row CellRow) String() string {
	return strconv.Itoa(int(row))
}

type CellCol int

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (col CellCol) String() string {
	//  'A', ... 'Z', 'AA', 'AB', ... 'AZ', 'BA', ... 'BZ', ... 'ZZ', 'AAA', ...
	var stringBytes []byte
	for {
		stringBytes = append(stringBytes, 'A'+byte(col%26))
		col /= 26
		if col <= 0 {
			break
		}
		col -= 1
	}
	reverse(stringBytes)
	return string(stringBytes)
}

type CellLoc struct {
	Row CellRow
	Col CellCol
}

func (loc CellLoc) String() string {
	return loc.Col.String() + loc.Row.String()
}

var ZERO_CELL_LOC = CellLoc{0, 0}

func ParseCellLoc(str string) (CellLoc, error) {
	loc := ZERO_CELL_LOC

	if len(str) < 2 {
		return loc, errors.New("invalid cell location: len(str) < 2")
	}

	str = strings.ToUpper(str)

	runeCol := str[0]
	if runeCol > 'Z' || runeCol < 'A' {
		return ZERO_CELL_LOC, fmt.Errorf("invalid cell column: '%c'", runeCol)
	}
	loc.Col = CellCol(runeCol - 'A')

	strRow := str[1:]
	irow, err := strconv.Atoi(strRow)
	if err != nil {
		return ZERO_CELL_LOC, fmt.Errorf("invalid cell row: '%s'", strRow)
	}
	loc.Row = CellRow(irow)

	return loc, nil
}

type CellData interface {
	Eval() float64
	Display() string
	IsDependOn(cell *Cell) bool
}

type CellDataNumber struct {
	Number float64
}

func (data *CellDataNumber) Eval() float64 {
	return data.Number
}

func (data *CellDataNumber) Display() string {
	return fmt.Sprintf("%f", data.Number)
}

func (data *CellDataNumber) IsDependOn(cell *Cell) bool {
	return false
}

type CellDataFormula struct {
	Formula  string
	CalCells map[CellLoc]*Cell
}

func (data *CellDataFormula) Eval() float64 {
	sum := 0.0
	for _, calCell := range data.CalCells {
		if calCell != nil {
			sum += calCell.Data.Eval()
		}
	}
	return sum
}

func (data *CellDataFormula) Display() string {
	num := data.Eval()
	return fmt.Sprintf("%f", num)
}

func (data *CellDataFormula) IsDependOn(cell *Cell) bool {
	_, ok := data.CalCells[cell.Loc]
	return ok
}

type CellDataInvalid struct {
	Reason string
}

func (data *CellDataInvalid) Eval() float64 {
	return math.NaN()
}

func (data *CellDataInvalid) Display() string {
	return data.Reason
}

func (data *CellDataInvalid) IsDependOn(cell *Cell) bool {
	return false
}

type CellDataText struct {
	Text string
}

func (data *CellDataText) Eval() float64 {
	return 0.0
}

func (data *CellDataText) Display() string {
	return data.Text
}

func (data *CellDataText) IsDependOn(cell *Cell) bool {
	return false
}
