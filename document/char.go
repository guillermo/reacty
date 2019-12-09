package document

type Char struct {
	Content []rune // A char can be compose of several runes
}

func (c Char) String() string {
	return string(c.Content)
}
