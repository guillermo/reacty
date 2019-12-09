// Package commands implement common output commands for the termianl
package commands

import (
	"fmt"
)

// Command represent a CSI sequence
type Command string

// Sequence render the Control Sequence
func (c Command) Sequence(args ...interface{}) []byte {
	if len(args) == 0 {
		return []byte(c)
	}

	return []byte(fmt.Sprintf(string(c), args...))
}

// Sequence generate the given command with the provided arguments
func Sequence(name string, args ...interface{}) []byte {
	return Commands[name].Sequence(args...)
}

//https://invisible-island.net/xterm/ctlseqs/ctlseqs.html

// Commands is the list of fixed commands
var Commands = map[string]Command{
	// SMCUP will enable the alternative buffer
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
	"ENABLEPASTE":   "\x1b[?2004h", // Enable paste
	"DISABLEPASTE":  "\x1b[?2004l", // Disable paste
	"CHARSTYLE": "\x1b[%sm",
	"BGCOLOR": "\x1b[48;2;%d;%d;%dm",
	"FGCOLOR": "\x1b[38;2;%d;%d;%dm",

}
