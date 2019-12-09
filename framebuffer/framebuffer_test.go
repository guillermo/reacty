package framebuffer

import (
	"bytes"
	"testing"
)
func TestBigFrame(t *testing.T) {
	b := &bytes.Buffer{}
	fb := &Framebuffer{output: b}

	fb.shouldOutput(t, "")
	fb.lastFrame.shouldEqual(t)
	fb.SetSize(2, 2)

	fb.Set(1, 1, 'a')
	fb.Set(2, 2, 'b')
	fb.Sync()
	fb.lastFrame.shouldEqual(t, "a ", " b")


}
