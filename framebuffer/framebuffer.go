package framebuffer

import (
	"bytes"
	"io"
	"time"

	"github.com/guillermo/reacty/output"
	"sync"
)

// Framebuffer represents the terminal state.
type Framebuffer struct {
	sync.Mutex
	Output       io.Writer
	currentFrame Frame
	lastFrame    Frame
	timer        *time.Timer
	Fps          int
}

/*
AutoSync will call Sync Framebuffer.Fps times per second. If not called, Sync
have to be called manually to be able to see the changes in terminal
*/
func (fb *Framebuffer) AutoSync() {
	if fb.Fps == 0 {
		fb.Fps = 60
	}

	fb.timer = time.NewTimer(time.Second / time.Duration(fb.Fps))
	for _, ok := <-fb.timer.C; ok; {
		fb.Sync()
	}
}

// StopAutoSync will stop the AutoSync
func (fb *Framebuffer) StopAutoSync() {
	fb.timer.Stop()
}

// SetSize sets the framebuffer size
func (fb *Framebuffer) SetSize(rows, cols int) {
	fb.Lock()
	defer fb.Unlock()
	fb.currentFrame.SetSize(rows, cols)
	fb.lastFrame.SetSize(rows, cols)
}

// Sync will diff the current state and dump the changes into Output
func (fb *Framebuffer) Sync() error {
	fb.Lock()
	defer fb.Unlock()
	b := &bytes.Buffer{}
	for y := range fb.lastFrame.Rows {
		for x := range fb.lastFrame.Rows[y] {
			if fb.lastFrame.Rows[y][x] != fb.currentFrame.Rows[y][x] {
				b.Write(output.Commands["GOTO"].Sequence(y+1, x+1))
				ch := fb.currentFrame.Rows[y][x]
				if isPrintable(ch) {
					b.Write([]byte(string(ch)))
				}
				fb.lastFrame.Rows[y][x] = ch
			}
		}
	}
	fb.Output.Write(b.Bytes())
	return nil
}

// Set will set the character ch in the row and column given
func (fb *Framebuffer) Set(row, col int, ch rune) {
	fb.Lock()
	defer fb.Unlock()
	fb.currentFrame.Set(row, col, ch)
	return
}

func isPrintable(ch rune) bool {
	// Is utf8
	if ch > 127 {
		return true
	}
	// Is printable ASCII
	if ch >= firstASCIIChar && ch <= lastASCIIChar {
		return true
	}
	return false
}
