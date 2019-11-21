package main

import (
	"github.com/guillermo/reacty/events"
)

// Terminal represents anything that can be used as a terminal
type Terminal interface {
	Dimensions() (rows, cols int)
	Set(Row, Col int, Ch rune)
	NextEvent() events.Event
}

// Document holds the document
type Document struct {
	Terminal Terminal
	Body     []string
}

// Render renders the current document
func (d *Document) Render() {
	var line, col int

	rows, cols := d.Terminal.Dimensions()

	for _, content := range d.Body {
		for _, ch := range content {
			d.Terminal.Set(line+1, col+1, ch)
			col++
			if col >= cols {
				col = 0
				line++
				if line >= rows {
					return
				}
			}
		}

	}

}
