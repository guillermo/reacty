package framebuffer

import (
	"bytes"
	"testing"
)

func TestOrigin(t *testing.T) {
	b := &bytes.Buffer{}
	fb := &Framebuffer{Output: b}
	row, col := fb.Origin()
	if row != 1 || col != 1 {
		t.Errorf("Expecting origin to be 1x1. Got %dx%d", row, col)
	}

	fb.SetOrigin(1, 1)
	row, col = fb.Origin()
	if row != 1 || col != 1 {
		t.Errorf("Expecting origin to be 1x1. Got %dx%d", row, col)
	}

}

func TestSizes(t *testing.T) {
	b := &bytes.Buffer{}
	fb := &Framebuffer{Output: b}
	cols, rows := fb.DocumentSize()
	if cols != 0 || rows != 0 {
		t.Error("Expecting a terminal of 0x0")
	}

}

func TestBigFrame(t *testing.T) {
	b := &bytes.Buffer{}
	fb := &Framebuffer{Output: b}

	fb.terminal.shouldOutput(t, "")
	fb.terminal.lastFrame.shouldEqual(t)
	fb.terminal.SetSize(2, 2)

	fb.Set(1, 1, 'a')
	fb.Set(2, 2, 'b')
	fb.sync()
	fb.terminal.lastFrame.shouldEqual(t, "a ", " b")

	// After changing the origin it should update
	fb.SetOrigin(2, 1)
	fb.sync()
	fb.terminal.lastFrame.shouldEqual(t, " b", "  ")

	fb.SetOrigin(2, 2)
	fb.sync()
	fb.terminal.lastFrame.shouldEqual(t, "b ", "  ")

	fb.SetOrigin(1, 1)
	fb.sync()
	fb.terminal.lastFrame.shouldEqual(t, "a ", " b")

	// If we add a char, it should show in the correct position
	fb.SetOrigin(2, 2)
	fb.Set(3, 3, 'c')
	fb.sync()
	fb.doc.shouldSize(t, 3, 3)

	fb.terminal.lastFrame.shouldEqual(t, "b ", " c")

}
