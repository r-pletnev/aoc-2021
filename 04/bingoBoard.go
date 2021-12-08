package main

type Cell struct {
	v       int
	checked bool
}

type Board struct {
	data        []Cell
	columnCount int
}

func (b Board) rows() [][]Cell {
	var result = make([][]Cell, 0)
	var rowIndex int
	for i, elm := range b.data {
		rowIndex = i / b.columnCount
		if len(result) <= rowIndex {
			row := make([]Cell, 0)
			result = append(result, row)
		}
		result[rowIndex] = append(result[rowIndex], elm)
	}
	return result
}

func (b Board) columns() [][]Cell {
	var result = make([][]Cell, b.columnCount)
	var columnIndex int
	for i := 0; i < b.columnCount; i++ {
		result[i] = make([]Cell, 0)
	}
	for i, elm := range b.data {
		columnIndex = i % b.columnCount
		result[columnIndex] = append(result[columnIndex], elm)
	}
	return result
}

func (b Board) markCell(value int) {
	for i, cell := range b.data {
		if cell.v == value {
			b.data[i].checked = true
		}
	}
}

func (b Board) isBingo() bool {
	var isBingo bool

	for _, row := range b.rows() {
		isBingo = true
		for i := range row {
			if !row[i].checked {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}
	for _, col := range b.columns() {
		isBingo = true
		for i := range col {
			if !col[i].checked {
				isBingo = false
			}
		}
		if isBingo {
			return true
		}
	}
	return false
}

func (b Board) score() int {
	var result int
	if !b.isBingo() {
		return result
	}
	for _, cell := range b.data {
		if cell.checked {
			continue
		}
		result += cell.v
	}
	return result
}

func NewBoard(arr []int, width int) Board {
	if len(arr)%width != 0 {
		panic("bad table data")
	}
	var data = make([]Cell, len(arr))
	for i, val := range arr {
		data[i] = Cell{v: val}
	}
	return Board{data: data, columnCount: width}
}
