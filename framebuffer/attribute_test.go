package framebuffer

import "testing"

func TestAttribute_Sequence(t *testing.T) {
	tests := []struct {
		a    Attribute
		want string
	}{
	{Normal, "\x1b[0m"},
	{Bold, "\x1b[1m"},
	{Faint, "\x1b[2m"},
	{Italic, "\x1b[3m"},
	{Underline, "\x1b[4m"},
	{Blink, "\x1b[5m"},
	{Inverse, "\x1b[7m"},
	{Invisible, "\x1b[8m"},
	{Crossed, "\x1b[9m"},
	{Double, "\x1b[21m"},
	{Double | Bold, "\x1b[1;21m"},
	{Italic | Bold, "\x1b[1;3m"},
	{Underline | Blink | Bold, "\x1b[1;4;5m"},
	}
	for _, tt := range tests {
		t.Run(tt.a.name(), func(t *testing.T) {
			if got := tt.a.sequence(); string(got) != tt.want {
				t.Errorf("Attribute(%.10b %d).Sequence() = %q, want %q", tt.a,tt.a,string(got), tt.want)
			}
		})
	}
}
