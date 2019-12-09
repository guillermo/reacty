package faketerm

import (
	"testing"
)

func TestFakeTerm(t *testing.T) {
	term := New(3, 5)
	r, c := term.Size()
	if r != 3 || c != 5 {
		t.Error("Incorrect size")
	}

}

func ExampleFakeTerm() {
	t := New(3, 5)
	t.Set(1, 1, 'H')
	t.Set(2, 3, 'e')
	t.Set(3, 5, 'l')
	t.Sync()
	// Output:
	// H
	//   e
	//     l
}
