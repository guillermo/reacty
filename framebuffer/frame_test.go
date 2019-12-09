package framebuffer

import (
	"testing"
)

func (f *frame) shouldEqual(t *testing.T, rows ...string) {
	t.Helper()
	content := []string{}
	for _, r := range f.Rows {
		line := ""
		for _, char := range r {
			line += char.Content
		}
		content = append(content, line)
	}

	if len(content) != len(rows) {
		t.Errorf("Expecting (%d rows got %d) %v. Got: %v", len(rows), len(content), rows, content)
	}
	for i, r := range content {
		if r != rows[i] {
			t.Errorf("Expecting line %d to be %q. Got: %q", i+1, rows[i], r)
		}
	}
}

func TestSetOutOfRange(t *testing.T) {

	f := &frame{}
	f.SetSize(1, 1)
	f.Set(0, 1, '❌')
	f.Set(2, 1, '❌')
	f.Set(1, 0, '❌')
	f.Set(1, 2, '❌')
	f.Set(1, 1, '✔')

}
