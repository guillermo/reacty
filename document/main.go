package main

import (
	"fmt"
	"github.com/guillermo/reacty/events"
	"github.com/guillermo/reacty/terminal"
	"runtime/debug"
)

func main() {
	term, err := terminal.Open()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	doc := &Document{
		Terminal: term,
	}
	doc.Body = []string{"hello", "world"}
	defer func() {
		if r := recover(); r != nil {
			term.Close()
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	doc.Render()

	eventLoop(doc)

}

func eventLoop(doc *Document) {
EventLoop:
	for {
		event := doc.Terminal.NextEvent()
		switch e := event.(type) {
		case events.KeyboardEvent:
			if e.Key == "q" {
				break EventLoop
			}
		}
	}

}
