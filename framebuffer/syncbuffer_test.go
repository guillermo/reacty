package framebuffer

import (
	"bytes"
	"testing"
)

func (sb *syncBuffer) shouldOutput(t *testing.T, output string) {
	t.Helper()
	b := &bytes.Buffer{}

	sb.WriteTo(b)
	if b.String() != output {
		t.Errorf("Expected syncBuffer to output %q. Got %q", output, b.String())
	}
}

func TestSyncBuffer(t *testing.T) {

	sb := &syncBuffer{}
	sb.SetSize(2, 2)
	sb.Set(2, 2, '😎')

	sb.shouldOutput(t, "\x1b[2;2H😎")
	sb.shouldOutput(t, "")

	sb.Set(2, 2, ' ')
	sb.shouldOutput(t, "\x1b[2;2H ")

}
