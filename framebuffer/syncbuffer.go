package framebuffer

import (
	"bytes"
	"github.com/guillermo/reacty/commands"
	"io"
	"time"
)

type syncBuffer struct {
	currentFrame frame
	lastFrame    frame
	timer        *time.Timer
	Fps          int
}

// WriteTo will write to w the changes synce the last WriteTo call
func (sb *syncBuffer) WriteTo(w io.Writer) (int64, error) {
	b := &bytes.Buffer{}
	for y := range sb.lastFrame.Rows {
		for x := range sb.lastFrame.Rows[y] {
			if sb.lastFrame.Rows[y][x] != sb.currentFrame.Rows[y][x] {
				b.Write(commands.Commands["GOTO"].Sequence(y+1, x+1))
				ch := sb.currentFrame.Rows[y][x]
				if isPrintable(ch) {
					b.Write([]byte(string(ch)))
				}
				sb.lastFrame.Rows[y][x] = ch
			}
		}
	}

	n, err := w.Write(b.Bytes())
	return int64(n), err

}

// SetSize sets the framebuffer size
func (sb *syncBuffer) SetSize(rows, cols int) {
	sb.currentFrame.SetSize(rows, cols)
	sb.lastFrame.SetSize(rows, cols)
}

// Set will set the character ch in the row and column given
func (sb *syncBuffer) Set(row, col int, ch rune) {

	sb.currentFrame.Set(row, col, ch)
	return
}

func (sb *syncBuffer) Size() (rows, cols int) {
	return sb.currentFrame.Size()
}
