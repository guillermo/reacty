// Package framebuffer converts a matrix of characters and its properties to a sequence of commands to write to the given output
//
package framebuffer

import (
	"io"
)

// Framebuffer represents the terminal state.
type Framebuffer struct {
	syncBuffer
	output               io.Writer
}

// Open returns a framebuffer and start throwing events.
func Open(o io.Writer) *Framebuffer {
	fb := &Framebuffer{
		output: o,
	}

	return fb
}

// Close closes the output device in case it implements the io.Closer
func (fb *Framebuffer) Close() {
	if closer, ok := fb.output.(io.Closer) ; ok{
		closer.Close()
	}
}



// Sync will write to the output all the changes
func (fb *Framebuffer) Sync() (err error) {
	_, err = fb.writeTo(fb.output)
	return
}

