package framebuffer

//go:generate stringer -type=Attribute

import (
	"strings"
	"github.com/guillermo/reacty/commands"
)

// Attribute holds the properties of the character (Bold, Italic, Underline ...)
type Attribute uint

func (a Attribute) name() string {
	var values []string
	for i := uint(1) ; i <= uint(Double) ; i <<=1 {
		if uint(a) & i == 0{
			continue
		}
		if Attribute(i).String() == "" {
			panic(i)
		}
		values = append(values, Attribute(i).String())
	}

	return strings.Join(values, " ")
}

const (
	// Normal resets the character to the normal style
	Normal Attribute = 1 << iota
	// Bold makes the character bold
	Bold 
	// Faint lowers the intensity of the character
	Faint
	// Italic makes the character Italic
	Italic
	// Underline draws a line below the character
	Underline
	// Blink will show and hide the character
	Blink
	// Inverse will change the background and foreground colors
	Inverse
	// Invisible Will make the text invisible (but normally the text can be copy)
	Invisible
	// Crossed will cross the character
	Crossed
	// Double will make a double underline
	Double
)

var attributesCodes = map[Attribute]string{
	Normal: "0",
	Bold: "1",
	Faint: "2",
	Italic: "3",
	Underline: "4",
	Blink: "5",
	Inverse: "7",
	Invisible: "8",
	Crossed: "9",
	Double: "21",
}


func (a Attribute) sequence() []byte {
	var codes []string
	if a == 514 {
		a = 514
	}
	for i := uint(1) ; i <= uint(Double) ; i <<=1 {
		if uint(a) & i == 0{
			continue
		}
			if v, ok := attributesCodes[Attribute(i)] ; ok {
				codes = append(codes, v)
			}
	}
	if len(codes) == 0{
		return []byte("")
	}

	return commands.Commands["CHARSTYLE"].Sequence(strings.Join(codes,";"))

}