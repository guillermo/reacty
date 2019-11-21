package framebuffer

import (
	"bytes"
	"github.com/guillermo/reacty/output"
	"io"
	"sync"
)

// Frame represent a matrix of characters
type Frame struct {
	sync.Mutex
	Rows         [][]rune
	nRows, nCols int
}

// SetSize sets the size of the Frame
func (f *Frame) SetSize(rows, columns int) {
	f.Lock()
	defer f.Unlock()
	f.nRows = rows
	f.nCols = columns
	f.Rows = make([][]rune, rows, rows)
	for i := range f.Rows {
		f.Rows[i] = make([]rune, columns, columns)
	}
}

// Set changes a character in the given position
func (f *Frame) Set(row, col int, ch rune) {
	f.Lock()
	defer f.Unlock()
	if row <= 0 || row > len(f.Rows) ||
		col <= 0 || col > len(f.Rows[row-1]) {
		return
	}
	f.Rows[row-1][col-1] = ch
}

const (
	firstASCIIChar = '!' //041
	lastASCIIChar  = '~' //026
)

// WriteTo implements the WriteTo interface and writes the sequences to render the full frame
func (f *Frame) WriteTo(w io.Writer) (n int64, err error) {
	f.Lock()
	defer f.Unlock()
	b := &bytes.Buffer{}

	// Go to 1,1
	b.Write(output.Commands["GOTO"].Sequence(1, 1))
	for i, row := range f.Rows {
		for _, ch := range row {
			if isPrintable(ch) {
				b.Write([]byte(string(ch)))
			} else {
				b.Write([]byte(string(" ")))
			}
		}

		if i+1 < len(f.Rows) {
			b.Write([]byte("\n\r"))
		}
	}

	nn, err := w.Write(b.Bytes())
	return int64(nn), err
}
