package framebuffer

import (
	"bytes"
	"testing"
)

func (f *frame) shouldEqual(t *testing.T, rows ...string) {
	t.Helper()
	content := []string{}
	for _, r := range f.Rows {
		content = append(content, string(r))
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

func (f *frame) shouldOutput(t *testing.T, output string) {
	t.Helper()
	b := &bytes.Buffer{}
	if _, err := f.WriteTo(b); err != nil {
		t.Fatal(err)
	}

	if b.String() != output {
		t.Errorf("Expected frame output to be %q. Got %q", output, b.String())
	}

}

func TestFrame(t *testing.T) {

	f := &frame{}

	f.SetSize(2, 2)

	f.shouldOutput(t, "\x1b[1;1H  \n\r  ")

	f.Set(1, 2, '$')
	f.Set(2, 1, '❤')

	f.shouldOutput(t, "\x1b[1;1H $\n\r❤ ")
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
