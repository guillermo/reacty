package framebuffer

import (
	"bytes"
	"testing"
)

func TestFrame(t *testing.T) {

	f := &Frame{}

	f.SetSize(2, 2)

	b := &bytes.Buffer{}
	n, err := f.WriteTo(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 12 {
		t.Fatal(n)
	}

	if exp := "\x1b[1;1H  \n\r  "; b.String() != exp {
		t.Errorf("Expecting: %q Got: %q", exp, b.String())
	}
	b.Reset()

	f.Set(1, 2, '$')
	f.Set(2, 1, '❤')

	n, err = f.WriteTo(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 14 {
		t.Fatal(n)
	}

	if exp := "\x1b[1;1H $\n\r❤ "; b.String() != exp {
		t.Errorf("Expecting: %q Got: %q", exp, b.String())
	}
	b.Reset()
}

func TestSetOutOfRange(t *testing.T) {

	f := &Frame{}
	f.SetSize(1, 1)
	f.Set(0, 1, '❌')
	f.Set(2, 1, '❌')
	f.Set(1, 0, '❌')
	f.Set(1, 2, '❌')
	f.Set(1, 1, '✔')

}
