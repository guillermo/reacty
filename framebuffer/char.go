package framebuffer



// Char holds a character and its attributes in the terminal
type Char struct {
	Content string
	Attr Attribute
	Bg Background
	Fg Foreground
}

func (c1 Char) equal(c2 Char) bool {
	if c1.Content != c2.Content {
		return false
	}
	if c1.Attr != c2.Attr {
		return false
	}
	if c1.Bg != c2.Bg {
		return false
	}
	if c1.Fg != c2.Fg {
		return false
	}

	return true
}


// NewChar generates a char with the given properties
func NewChar(ch rune, attrs ...interface{}) Char {

	char := Char{Content: string(ch), Attr: Normal}
	for _, attr := range attrs {
		switch v := attr.(type) {
		case Foreground:
			char.Fg = v
		case Background:
			char.Bg = v 
		case Attribute:
			char.Attr |= v
		}
	}
	return char

}