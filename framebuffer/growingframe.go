package framebuffer

type growingframe frame

func (f *growingframe) Set(row, col int, ch rune) {
	if row <= 0 {
		panic("Got a no row")
	}
	if col <= 0 {
		panic("Got a 0 row")
	}

	for len(f.Rows) < row {
		f.Rows = append(f.Rows, make([]rune, col))
	}
	for len(f.Rows[row-1]) < col {
		f.Rows[row-1] = append(f.Rows[row-1], ' ')
	}
	f.Rows[row-1][col-1] = ch
	if f.nCols < row {
		f.nCols = row
	}
	if f.nRows < col {
		f.nRows = col
	}
}

func (f *growingframe) get(row, col int) rune {
	if row <= 0 || len(f.Rows) < row {
		return ' '
	}
	r := f.Rows[row-1]

	if col <= 0 || len(r) < col {
		return ' '
	}
	ch := r[col-1]
	if ch == 0 {
		return ' '
	}
	return ch
}

type framer interface {
	Set(row, col int, ch rune)
	Size() (rows, cols int)
}

func (f *growingframe) CopyTo(dest framer, originRow, originCol int) {
	rows, cols := dest.Size()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			dest.Set(row+1, col+1, f.get(row+originRow, col+originCol))
		}

	}

}
