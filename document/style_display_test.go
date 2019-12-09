package document

import (
	"testing"

	"golang.org/x/net/html/atom"
)

func Test_defaultDisplayStyleFor(t *testing.T) {
	type args struct {
		a atom.Atom
	}
	tests := []struct {
		name string
		el atom.Atom
		want StyleDisplay
	}{
		// TODO: Add test cases.
		{"P", atom.P, DisplayBlock},
		{"img", atom.Img, DisplayInline},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultDisplayStyleFor(tt.el); got != tt.want {
				t.Errorf("defaultDisplayStyleFor() = %v, want %v", got, tt.want)
			}
		})
	}
}
