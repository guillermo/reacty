package framebuffer

type growingframe frame

func (d *growingframe) Set(row, col int, ch rune) {
	for len(d.Rows) < row {
		d.Rows = append(d.Rows, make([]rune, col))
	}
	for len(d.Rows[row-1]) < col-1 {
		d.Rows[row-1] = append(d.Rows[row-1], ' ')
	}
	d.Rows[row-1][col-1] = ch
	if d.nCols < row {
		d.nCols = row
	}
	if d.nRows < col {
		d.nRows = col
	}
}

func (g *growingframe) get(row, col int) rune {
	if row <= 0 || len(g.Rows) < row {
		return ' '
	}
	r := g.Rows[row-1]

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

func (d *growingframe) CopyTo(f framer, originRow, originCol int) {
	rows, cols := f.Size()

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			f.Set(row+1, col+1, d.get(row+originRow, col+originCol))
		}

	}

}
