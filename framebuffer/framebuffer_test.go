package framebuffer

import (
	"bytes"
	"testing"
)

func TestFramebuffer(t *testing.T) {

	b := &bytes.Buffer{}

	fb := &Framebuffer{Output: b}
	fb.SetSize(2, 2)

	err := fb.Sync()
	if err != nil {
		t.Fatal(err)
	}

	if len(b.Bytes()) != 0 {
		t.Fatal("Expected 0 Got:", len(b.Bytes()))
	}

	fb.Set(2, 2, '😎')

	err = fb.Sync()
	if err != nil {
		t.Fatal(err)
	}

	if exp := "\x1b[2;2H😎"; b.String() != exp {
		t.Errorf("Expected %q Got %q", exp, b.String())
	}

	b.Reset()

	err = fb.Sync()
	if err != nil {
		t.Fatal(err)
	}

	if exp := ""; b.String() != exp {
		t.Fatalf("Expected %q Got %q", exp, b.String())
	}
}
