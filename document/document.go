package document

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
)

type Document struct {
	Chars CharArea
	node  *html.Node
	Width int
}

func (d *Document) Parse(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	d.node = doc
	d.render()
	return nil
}

type cursor struct {
	row, col int
}

func (d *Document) render() {
	cursor := &cursor{}
	d.renderNode(cursor, d.node)

}

func (d *Document) renderNode(cursor *cursor, node *html.Node) {
	if node.Type == html.TextNode {
		d.renderTextNode(cursor, node)
	}
	if node.Type == html.ElementNode {
		d.renderElementNode(cursor, node)
	}
	if node.Type == html.DocumentNode {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			d.renderNode(cursor, c)
		}

	}
}

func (d *Document) renderTextNode(cursor *cursor, node *html.Node) {
	runes := []rune(node.Data)
	for _, ch := range runes {
		d.Chars.Set(cursor.row+1, cursor.col+1, Char{Content: []rune{ch}})
		cursor.col++
		if cursor.col >= d.Width {
			cursor.col = 0
			cursor.row++
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		d.renderNode(cursor, c)
	}

}

func (d *Document) renderElementNode(cursor *cursor, node *html.Node) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		d.renderNode(cursor, c)
	}

	if node.DataAtom == atom.P {
		cursor.row++
		cursor.col = 0
	}

}
