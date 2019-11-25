package framebuffer

import (
	"io"
	"time"
)

// Framebuffer represents the terminal state.
type Framebuffer struct {
	terminal             syncBuffer
	doc                  growingframe
	originRow, originCol int
	Output               io.Writer
}

func Open(o io.Writer) *Framebuffer {
	fb := &Framebuffer{
		Output: o,
	}

	go fb.autoSync()
	return fb
}

func (fb *Framebuffer) Close() {
	fb.stopAutoSync()
}

// Origin returns the current origin of the document
func (fb *Framebuffer) Origin() (row, col int) {
	return fb.originRow + 1, fb.originCol + 1
}

func (fb *Framebuffer) SetOrigin(row, col int) {
	fb.originRow = row - 1
	fb.originCol = col - 1
	fb.doc.CopyTo(&fb.terminal, fb.originRow+1, fb.originCol+1)
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

// DocumentSize return the size of the document
func (fb *Framebuffer) DocumentSize() (cols, rows int) {
	return fb.doc.nCols, fb.doc.nCols
}

// TerminalSize returns the size of the terminal
func (fb *Framebuffer) TerminalSize() (cols, rows int) {
	return fb.terminal.Size()
}

//autoSync will call Sync Framebuffer.Fps times per second. If not called, Sync
//have to be called manually to be able to see the changes in terminal
func (fb *Framebuffer) autoSync() {
	if fb.terminal.Fps == 0 {
		fb.terminal.Fps = 60
	}

	fb.terminal.timer = time.NewTimer(time.Second / time.Duration(fb.terminal.Fps))
	for _, ok := <-fb.terminal.timer.C; ok; {
		// todo
		fb.sync()
	}
}

// sync will write to the output all the changes
func (fb *Framebuffer) sync() (err error) {
	_, err = fb.terminal.WriteTo(fb.Output)
	return
}

// stopAutoSync will stop the AutoSync
func (fb *Framebuffer) stopAutoSync() {
	fb.terminal.timer.Stop()
}

func (fb *Framebuffer) Set(row, col int, ch rune) {
	fb.doc.Set(row, col, ch)
	fb.terminal.Set(row-(fb.originRow), col-(fb.originCol), ch)
}
