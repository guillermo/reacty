package document

//go:generate stringer -type StyleDisplay
import (
	"fmt"
	"golang.org/x/net/html/atom"
)

var inlineStytleItems = []atom.Atom{
	atom.A,
	atom.Abbr,
	atom.Acronym,
	atom.Audio,
	atom.B,
	atom.Bdi,
	atom.Bdo,
	atom.Big,
	atom.Br,
	atom.Button,
	atom.Canvas,
	atom.Cite,
	atom.Code,
	atom.Data,
	atom.Datalist,
	atom.Del,
	atom.Dfn,
	atom.Em,
	atom.Embed,
	atom.I,
	atom.Iframe,
	atom.Img,
	atom.Input,
	atom.Ins,
	atom.Kbd,
	atom.Label,
	atom.Map,
	atom.Mark,
	atom.Meter,
	atom.Noscript,
	atom.Object,
	atom.Output,
	atom.Picture,
	atom.Progress,
	atom.Q,
	atom.Ruby,
	atom.S,
	atom.Samp,
	atom.Script,
	atom.Select,
	atom.Slot,
	atom.Small,
	atom.Span,
	atom.Strong,
	atom.Sub,
	atom.Sup,
	atom.Svg,
	atom.Template,
	atom.Textarea,
	atom.Time,
	atom.U,
	atom.Tt,
	atom.Var,
	atom.Video,
	atom.Wbr,
}

func defaultDisplayStyleFor(a atom.Atom) StyleDisplay {
	for _, el := range inlineStytleItems {
		if el == a {
			return DisplayInline
		}
	}
	return DisplayBlock
}

type StyleDisplay uint8

const (
	DisplayInline StyleDisplay = iota
	DisplayBlock
	DisplayNone
)

func displayStyle(s string) (StyleDisplay, error) {
	if s == "inline" {
		return DisplayInline, nil
	}
	if s == "block" {
		return DisplayBlock, nil
	}
	if s == "none" {
		return DisplayNone, nil
	}
	return 0, fmt.Errorf("can't parse %q as display value", s)
}
