package framebuffer

// frame represent a matrix of characters
type frame struct {
	Rows         [][]Char
	nRows, nCols int
}

// SetSize sets the size of the Frame and empty it
func (f *frame) SetSize(rows, cols int) {
	f.nRows = rows
	f.nCols = cols
	f.Rows = make([][]Char, rows, rows)
	for i := range f.Rows {
		f.Rows[i] = make([]Char, cols, cols)
		for k := 0; k < cols; k++ {
			f.Rows[i][k] = Char{Content:" ", Attr: Normal}
		}
	}
}

func (f *frame) Size() (rows, cols int) {
	return f.nRows, f.nCols
}

// Set changes a character in the given position
func (f *frame) Set(row, col int, ch rune, attrs ...interface{}) {
	if row <= 0 || row > f.nRows ||
		col <= 0 || col > f.nCols {
		return
	}
	char := NewChar(ch, attrs...)
	f.Rows[row-1][col-1] = char
}

const (
	firstASCIIChar = '!' //041
	lastASCIIChar  = '~' //026
)
