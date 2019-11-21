package main

import (
	"github.com/guillermo/reacty/terminal"
)

type Document struct {
	Terminal *terminal.Terminal
	Body     []string
}

func Open() (*Document, error) {

	term, err := terminal.Open()

	return &Document{Terminal: term}, err
}

func (d *Document) Close() error {
	return d.Terminal.Close()
}

func (d *Document) Render() {

	for row, content := range d.Body {
		for i, ch := range content {
			d.Terminal.Set(row+1, 1+i, ch)

		}

	}

}
