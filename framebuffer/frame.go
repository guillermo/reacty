package framebuffer

import (
	"bytes"
	"github.com/guillermo/reacty/output"
	"io"
)

type Frame struct {
	Rows         [][]rune
	nRows, nCols int
}

func (f *Frame) SetSize(rows, columns int) {
	f.nRows = rows
	f.nCols = columns
	f.Rows = make([][]rune, rows, rows)
	for i, _ := range f.Rows {
		f.Rows[i] = make([]rune, columns, columns)
	}
}

func (f *Frame) Set(row, col int, ch rune) {
	if row < 1 || row > f.nRows ||
		col < 1 || col > f.nCols {
		panic("Set rune out of boundaries")
	}
	f.Rows[row-1][col-1] = ch
}

const (
	firstASCIIChar = '!' //041
	lastASCIIChar  = '~' //026
)

func (f *Frame) WriteTo(w io.Writer) (n int64, err error) {
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
