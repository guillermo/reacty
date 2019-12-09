package document

import (
	"strings"
	"testing"
)

func TestDocument(t *testing.T) {

	doc := &Document{Width: 10}
	doc.Parse(strings.NewReader(`<P>Paragraph</p><p>Paragraph 2</p>`))
	doc.Chars.shouldEqual(t, "Paragraph ", "Paragraph ", "2         ")

}
