package main

import (
	"fmt"
	"github.com/guillermo/reacty/events"
	"runtime/debug"
)

func main() {
	doc, err := Open()
	if err != nil {
		panic(err)
	}
	doc.Body = []string{"hello", "world"}
	defer func() {
		if r := recover(); r != nil {
			doc.Close()
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	doc.Render()

	eventLoop(doc)

	doc.Close()

}

func eventLoop(doc *Document) {
EventLoop:
	for event := range doc.Terminal.Events {
		switch e := event.(type) {
		case events.KeyboardEvent:
			if e.Key == "q" {
				break EventLoop
			}
		}
	}

}
