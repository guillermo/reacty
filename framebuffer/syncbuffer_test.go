package framebuffer

import (
	"bytes"
	"testing"
)

func (sb *syncBuffer) shouldOutput(t *testing.T, output string) {
	t.Helper()
	b := &bytes.Buffer{}

	sb.writeTo(b)
	if b.String() != output {
		t.Errorf("Expected syncBuffer to output %q. Got %q", output, b.String())
	}
}

func TestSyncBuffer(t *testing.T) {

	sb := &syncBuffer{}
	sb.SetSize(2, 2)
	sb.shouldOutput(t, "")
	sb.Set(1, 1, '😎')
	sb.shouldOutput(t, "\x1b[1;1H😎")
	sb.shouldOutput(t, "")
	sb.Set(1,2,'h')
	sb.shouldOutput(t, "h")
	sb.Set(2,1,'e')
	sb.shouldOutput(t, "e")
	sb.Set(1,1,':',Bold,Underline,Foreground(Red),Background(Blue))
	// Order is relevant
	sb.shouldOutput(t, "\x1b[1;1H\x1b[0;1;4m\x1b[38;2;255;0;0m\x1b[38;2;0;0;255m:")
	sb.Set(1,2,'O',Bold,Underline,Foreground(Red),Background(Blue))
	sb.shouldOutput(t, "O")
	sb.Set(2,1,'O',Underline,Foreground(Red),Background(Blue))
	sb.shouldOutput(t, "\x1b[0;4mO")

}
