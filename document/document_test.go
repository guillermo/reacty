package document

import (
	"strings"
	"testing"
)

var docTests = []struct {
	in  string
	out []string
}{
	{`<P>Paragraph</p><p>Paragraph 2</p>`, []string{"Paragraph ", "Paragraph ", "2         "}},
	{`<P style="display:inline;">P1</p><p style="display:inline;">P2</p>`, []string{"P1P2"}},
}

func TestDocument(t *testing.T) {

	for _, tCase := range docTests {
		t.Run(tCase.in, func(t *testing.T) {
			doc := &Document{Width: 10}
			doc.Parse(strings.NewReader(tCase.in))
			doc.Chars.shouldEqual(t, tCase.out...)
		})
	}

}
