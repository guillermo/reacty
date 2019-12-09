package document

import (
	"testing"
)

var styleTest = []struct {
	in  string
	out Style
}{
	{"", Style{}},
	{"   diSplay : inlIne   ", Style{"display": DisplayInline}},
	{"diSplay :  Block ", Style{"display": DisplayBlock}},
}

func TestStyle(t *testing.T) {

	for _, tCase := range styleTest {
		t.Run(tCase.in, func(t *testing.T) {
			s, errs := ParseStyle(tCase.in)
			if len(errs) != 0 {
				t.Fatalf("Error parsing %q: %v. Expected: %v, Got: %v", tCase.in, errs, tCase.out, s)
			}
			if !tCase.out.Equal(s) {
				t.Errorf("Expected style %q to match %v. Got %v", tCase.in, tCase.out, s)
			}
		})
	}

}
