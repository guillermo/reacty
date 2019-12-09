package framebuffer

import (
	"bytes"
	"github.com/guillermo/reacty/commands"
	"io"
	"sync"
)

type cursor struct {
	row  int
	col  int
	attr Attribute
	fg   Foreground
	bg   Background
}

type syncBuffer struct {
	sync.Mutex
	c            cursor
	currentFrame frame
	lastFrame    frame
}

// writeTo will write to w the changes synce the last writeTo call
func (sb *syncBuffer) writeTo(w io.Writer) (int64, error) {
	b := &bytes.Buffer{}

	sb.eachChange(func(row, col int, current, previous Char) {
		b.Write(sb.draw(row,col, current))
		sb.updateChar(row, col)
	})

	n, err := w.Write(b.Bytes())
	return int64(n), err
}

func (sb *syncBuffer) draw(row, col int, char Char) []byte {
	if sb.c.attr == 0 {
		sb.c.attr = Normal
	}
	b := &bytes.Buffer{}

	if sb.c.row != row || sb.c.col != col {
		b.Write(commands.Sequence("GOTO",row,col))
		sb.c.row = row 
		sb.c.col = col
	}
	if sb.c.attr != char.Attr {
		b.Write(char.Attr.sequence())
		sb.c.attr = char.Attr
	}
	if sb.c.fg != char.Fg {
		b.Write(char.Fg.sequence())
		sb.c.fg = char.Fg
	}
	if sb.c.bg != char.Bg {
		b.Write(char.Bg.sequence())
		sb.c.bg = char.Bg
	}
	runes := []rune(char.Content)
	if len(runes) == 0 || !isPrintable(runes[0]) {
		b.Write([]byte(" "))
	} else {
		b.Write([]byte(string(runes[0])))
	}
	sb.c.col++
	if sb.c.col > sb.currentFrame.nCols {
		sb.c.col = 1
		sb.c.row++
	}

	return b.Bytes()
}


func isPrintable(ch rune) bool {
	// Is utf8
	if ch > 127 {
		return true
	}
	// Is printable ASCII
	if ch >= firstASCIIChar && ch <= lastASCIIChar {
		return true
	}
	return false
}


func (sb *syncBuffer) updateChar(row, col int) {
	sb.lastFrame.Rows[row-1][col-1] = sb.currentFrame.Rows[row-1][col-1]
}

func (sb *syncBuffer) eachChange(fn func(row, col int, current, previous Char)) {
	for y := range sb.lastFrame.Rows {
		for x, lastChar := range sb.lastFrame.Rows[y] {
			currentChar := sb.currentFrame.Rows[y][x]
			if !lastChar.equal(currentChar) {
				fn(y+1, x+1, currentChar, lastChar)
			}
		}
	}
}

// SetSize sets the framebuffer size
func (sb *syncBuffer) SetSize(rows, cols int) {
	sb.Lock()
	defer sb.Unlock()
	sb.currentFrame.SetSize(rows, cols)
	sb.lastFrame.SetSize(rows, cols)
}

// Set will set the character ch in the row and column given
func (sb *syncBuffer) Set(row, col int, ch rune, attributes ...interface{}) {
	sb.Lock()
	defer sb.Unlock()

	sb.currentFrame.Set(row, col, ch, attributes...)
	return
}

func (sb *syncBuffer) Size() (rows, cols int) {
	rows, cols = sb.currentFrame.Size()
	sb.Lock()
	defer sb.Unlock()
	return sb.currentFrame.Size()
}
