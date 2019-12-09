package document

import (
	"testing"
)

func (cm *CharArea) shouldSize(t *testing.T, rows, cols int) {
	if cm.Rows != rows {
		t.Errorf("Expecting %d rows. Got %d", rows, cm.Rows)
	}
	if cm.Cols != cols {
		t.Errorf("Expecting %d cols. Got %d", cols, cm.Cols)
	}
}

func (cm *CharArea) text() []string {

	lines := make([]string, cm.Rows)
	for row := 0; row < cm.Rows; row++ {
		for col := 0; col < cm.Cols; col++ {
			ch := cm.get(row+1, col+1)
			if len(ch.Content) == 0 {
				lines[row] += " "
			} else {
				lines[row] += ch.String()
			}

		}
	}
	return lines
}

func (cm *CharArea) shouldEqual(t *testing.T, Expected ...string) {
	t.Helper()
	lines := cm.text()

	if len(Expected) != len(lines) {
		t.Errorf("Expected document to be %d lines. Got %d lines", len(Expected), len(lines))
	}

	for n := range Expected {
		line := ""
		if len(lines) > n {
			line = lines[n]
		}
		if Expected[n] != line {
			t.Errorf("Expected document line %d to be %q. Got %q", n+1, Expected[n], line)
		}
	}
}

func TestDocumentShouldGrow(t *testing.T) {

	d := &CharArea{}

	d.Set(2, 2, Char{Content: []rune{'✌'}})

	d.shouldSize(t, 2, 2)
	d.shouldEqual(t, "  ", " ✌")

}
