package framebuffer

import (
	"bytes"
	"github.com/guillermo/reacty/commands"
	"io"
	"sync"
	"time"
)

type syncBuffer struct {
	sync.Mutex
	currentFrame frame
	lastFrame    frame
	timer        *time.Ticker
	Fps          int
	changed      bool
}

// WriteTo will write to w the changes synce the last WriteTo call
func (sb *syncBuffer) WriteTo(w io.Writer) (int64, error) {
	sb.Lock()
	defer sb.Unlock()
	if !sb.changed {
		return 0, nil
	}
	b := &bytes.Buffer{}
	for y := range sb.lastFrame.Rows {
		for x := range sb.lastFrame.Rows[y] {
			if sb.lastFrame.Rows[y][x] != sb.currentFrame.Rows[y][x] {
				b.Write(commands.Commands["GOTO"].Sequence(y+1, x+1))
				ch := sb.currentFrame.Rows[y][x]
				if isPrintable(ch) {
					b.Write([]byte(string(ch)))
				} else {
					b.Write([]byte(" "))
				}
				sb.lastFrame.Rows[y][x] = ch
			}
		}
	}

	n, err := w.Write(b.Bytes())
	sb.changed = false
	return int64(n), err

}

// SetSize sets the framebuffer size
func (sb *syncBuffer) SetSize(rows, cols int) {
	sb.Lock()
	defer sb.Unlock()
	sb.changed = true
	sb.currentFrame.SetSize(rows, cols)
	sb.lastFrame.SetSize(rows, cols)
}

// Set will set the character ch in the row and column given
func (sb *syncBuffer) Set(row, col int, ch rune) {
	sb.Lock()
	defer sb.Unlock()
	sb.changed = true

	sb.currentFrame.Set(row, col, ch)
	return
}

func (sb *syncBuffer) Size() (rows, cols int) {
	sb.Lock()
	defer sb.Unlock()
	return sb.currentFrame.Size()
}
