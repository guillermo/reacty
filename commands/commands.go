// Package commands implement common output commands for the termianl
package commands

import (
	"fmt"
)

// Command implements the Commander Interface
type Command string

// Sequence render the Control Sequence
func (c Command) Sequence(args ...interface{}) []byte {
	if len(args) == 0 {
		return []byte(c)
	}

	return []byte(fmt.Sprintf(string(c), args...))
}

//https://invisible-island.net/xterm/ctlseqs/ctlseqs.html

// Commands is the list of fixed commands
var Commands = map[string]Command{
	"SMCUP":         "\x1b[?1049h\x1b[22;0;0t", // Activate alternate buffer
	"RMCUP":         "\x1b[?1049l\x1b[23;0;0t", // Undo ^^
	"HIDECURSOR":    "\x1b[?25l",
	"SHOWCURSOR":    "\x1b[?25h",
	"SOFTRESET":     "\x1b[>!p",
	"ENABLEMOUSE":   "\x1b[?1000;1006;1015h",
	"DISABLEMOUSE":  "\x1b[?1000;1006;1015l",
	"CLEAR":         "\x1b\x0c",
	"CURSORUP":      "\x1bA",
	"CURSORDOWN":    "\x1bB",
	"CURSORRIGHT":   "\x1bC",
	"CURSORLEFT":    "\x1bD",
	"CURSORHOME":    "\x1bH",
	"ERASEBELOW":    "\x1b[0J",
	"ERASEABOVE":    "\x1b[1J",
	"ERASEALL":      "\x1b[2J",
	"GETWINDOWSIZE": "\x1b[18t",
	"CURSORUPN":     "\x1b[%dA",
	"GOTO":          "\x1b[%d;%dH", // RowxColumn starting on 1
	"ENABLEPASTE":   "\033[?2004h", // Enable paste
	"DISABLEPASTE":  "\033[?2004l", // Disable paste

}
