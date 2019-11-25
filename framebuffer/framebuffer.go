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
	output               io.Writer
}

// Open returns a framebuffer and start throwing events.
func Open(o io.Writer) *Framebuffer {
	fb := &Framebuffer{
		output: o,
	}

	go fb.autoSync()
	return fb
}

// Close stop writing to the Output writer
func (fb *Framebuffer) Close() {
	fb.stopAutoSync()
}

// Origin returns the current origin of the document
func (fb *Framebuffer) Origin() (row, col int) {
	return fb.originRow + 1, fb.originCol + 1
}

// SetOrigin sets the document origin. A origin of 1,1 will be the start of the document
func (fb *Framebuffer) SetOrigin(row, col int) {
	fb.originRow = row - 1
	fb.originCol = col - 1
	fb.doc.CopyTo(&fb.terminal, fb.originRow+1, fb.originCol+1)
}

// SetTerminalSize sets the terminal Size
func (fb *Framebuffer) SetTerminalSize(rows, cols int) {
	fb.terminal.SetSize(rows, cols)
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
		fb.terminal.Fps = 25
	}

	fb.terminal.timer = time.NewTicker(time.Second / time.Duration(fb.terminal.Fps))
	for range fb.terminal.timer.C {
		fb.sync()
	}
}

// sync will write to the output all the changes
func (fb *Framebuffer) sync() (err error) {
	_, err = fb.terminal.WriteTo(fb.output)
	return
}

// stopAutoSync will stop the AutoSync
func (fb *Framebuffer) stopAutoSync() {
	fb.terminal.timer.Stop()
}

// Set sets the rune in position *row*/*col* to be *ch*.
func (fb *Framebuffer) Set(row, col int, ch rune) {
	fb.doc.Set(row, col, ch)
	fb.terminal.Set(row-(fb.originRow), col-(fb.originCol), ch)
}
