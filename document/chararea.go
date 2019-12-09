package document

type CharArea struct {
	Content    [][]Char
	Cols, Rows int
}

func (cm *CharArea) Set(row, col int, ch Char) {
	if row <= 0 {
		panic("Got a no row")
	}
	if col <= 0 {
		panic("Got a 0 row")
	}

	for len(cm.Content) < row {
		cm.Content = append(cm.Content, make([]Char, col))
	}
	for len(cm.Content[row-1]) < col {
		cm.Content[row-1] = append(cm.Content[row-1], Char{})
	}
	cm.Content[row-1][col-1] = ch
	if cm.Cols < col {
		cm.Cols = col
	}
	if cm.Rows < row {
		cm.Rows = row
	}
}

func (cm *CharArea) get(row, col int) Char {
	if row <= 0 || len(cm.Content) < row {
		return Char{Content: []rune{' '}}
	}
	r := cm.Content[row-1]

	if col <= 0 || len(r) < col {
		return Char{Content: []rune{' '}}
	}

	return r[col-1]
}

type framer interface {
	Set(row, col int, ch rune)
	Size() (rows, cols int)
}
