package framebuffer

import (
	"strings"
	"testing"
)

func (d *growingframe) shouldSize(t *testing.T, rows, cols int) {
	if d.nRows != rows {
		t.Errorf("Expecting %d rows. Got %d", rows, d.nRows)
	}
	if d.nCols != cols {
		t.Errorf("Expecting %d cols. Got %d", cols, d.nCols)
	}
}

func (d *growingframe) shouldEqual(t *testing.T, data []string) {
	t.Helper()
	c := []string{}
	for _, r := range d.Rows {
		c = append(c, strings.Replace(string(r), "\x00", " ", -1))
	}

	if len(c) != len(data) {
		t.Fatalf("Expected document to be %d lines. Got %d lines", len(c), len(data))
	}

	for i, r := range data {
		if c[i] != r {
			t.Errorf("Expected document line %d to be %q. Got %q", i, r, c[i])
		}
	}
}

func (d *growingframe) expectedSize(t *testing.T, rows, cols int) {
	if d.nCols != rows {
		t.Errorf("Expected document to have %d rows. Got %d.", rows, d.nCols)
	}

	if d.nRows != cols {
		t.Errorf("Expected document to have %d cols. Got %d.", cols, d.nRows)
	}
}

func TestDocumentShouldGrow(t *testing.T) {

	d := &growingframe{}

	d.Set(2, 2, '✌')

	d.expectedSize(t, 2, 2)
	d.shouldEqual(t, []string{"  ", " ✌"})

}

func TestCopy(t *testing.T) {
	gf := &growingframe{}
	dest := &frame{}
	dest.SetSize(2, 2)
	dest.Set(1, 1, '1')
	dest.Set(1, 2, '2')
	dest.Set(2, 1, '3')
	dest.Set(2, 2, '4')
	dest.shouldEqual(t, "12", "34")

	gf.CopyTo(dest, 1, 1)
	dest.shouldEqual(t, "  ", "  ")

	gf.Set(1, 1, 'a')
	gf.CopyTo(dest, 1, 1)
	dest.shouldEqual(t, "a ", "  ")

	gf.Set(2, 2, 'b')
	gf.CopyTo(dest, 1, 1)
	dest.shouldEqual(t, "a ", " b")

	gf.CopyTo(dest, 2, 1)
	dest.shouldEqual(t, " b", "  ")

	gf.CopyTo(dest, 2, 2)
	dest.shouldEqual(t, "b ", "  ")

	gf.Set(3, 3, 'c')
	gf.CopyTo(dest, 3, 3)
	dest.shouldEqual(t, "c ", "  ")
}
