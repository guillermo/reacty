package document


import (
	"testing"
)

func TestColor(t *testing.T) {
	r,g,b,a := White.RGBA()
	if r != 0xffff {
		t.Errorf("%b",r)
	}
	if g != 0xffff {
		t.Error(g)
	}
	if b != 0xffff {
		t.Error(b)
	}
	if a != 0xffff {
		t.Error(a)
	}
	if White != FromColor(White) {
		t.Fatal("Expected Black to equal Black", White, FromColor(White))
	}
}