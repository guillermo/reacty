package input

import (
	"fmt"
)

// BytesEvent holds all the unrecognized bytes from the input
type BytesEvent string

func (b BytesEvent) String() string {
	return fmt.Sprintf("BytesEvent: %s", []byte(b))
}

// Event holds the information of the event.
type Event interface {
	String() string
}

//https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/code/code_values

// KeyboardEvent represent a key press.
type KeyboardEvent struct {
	// Key Is the printibal character
	Key string
	// Ctrl is true if the Control key is pressed
	Ctrl bool
	// Alt is true if the Alt key is pressed
	Alt bool
	// Shift is true if the Shift key is pressed
	Shift bool
	// Code holds the name of the key. For example: Right, Escape or Enter
	Code string
}

func (ke KeyboardEvent) String() string {
	s := ke.Code
	if len(s) == 0 {
		s = ke.Key
	}

	if ke.Alt {
		s = "Alt+" + s
	}
	if ke.Ctrl {
		s = "Ctrl+" + s
	}

	return "KeyboardEvent: " + s
}

// WindowSizeEvent represents the window terminal size in characters
type WindowSizeEvent struct {
	Width  int
	Height int
}

func (wse *WindowSizeEvent) String() string {
	return fmt.Sprintf("WindowSizeEvent: %dx%d", wse.Width, wse.Height)
}

// ErrorEvent holds an error produce while processing the input
type ErrorEvent string

func (e ErrorEvent) String() string {
	return "ErrorEvent: " + string(e)
}
