// Package terminal abstract the terminal by providing two main interfaces: Input Events and a framebuffer output
package terminal

import (
	"io"
	"os"

	"github.com/guillermo/reacty/commands"
	"github.com/guillermo/reacty/events"
	"github.com/guillermo/reacty/framebuffer"
	"github.com/guillermo/reacty/input"
)

// Terminal holds the state of the current terminal
type Terminal struct {
	term
	r io.Reader

	w          io.Writer
	input      *input.Input
	fb         *framebuffer.Framebuffer
	events     <-chan (events.Event)
	cols, rows int
	log        *os.File
	FixedSize  [2]int
}

// Open configures the controlling tty to be used with Terminal
func Open() (*Terminal, error) {
	eventsChan := make(chan (events.Event), 1024)
	log, err := os.Create("cmds.log")
	if err != nil {
		panic(err)
	}

	stdout := io.MultiWriter(log, os.Stdout)

	t := &Terminal{
		term: term{
			Fd: int(os.Stdin.Fd()),
		},
		r:      os.Stdin,
		w:      stdout,
		fb:     framebuffer.Open(stdout),
		log:    log,
		input:  input.Open(os.Stdin),
		events: eventsChan,
	}
	if t.FixedSize[0] == 0 {
		err = t.onResize(func(rows, cols int) {
			t.fb.SetSize(rows, cols)
			t.sendWinSize(eventsChan, rows, cols)
		})
		if err != nil {
			return nil, err
		}

		if err := t.saveTerminalState(); err != nil {
			return nil, err
		}

		if err := t.rawMode(); err != nil {
			return nil, err
		}
	} else {
		t.fb.SetSize(t.FixedSize[0],t.FixedSize[1])
	}

	go t.forwardInputEvents(eventsChan)

	t.saveScreen()
	t.hideCursor()

	if err != nil {
		panic(err)
	}
	return t, nil
}

// Close resets the terminal to the previous state
func (t *Terminal) Close() error {
	if t.FixedSize[0] == 0 {

	t.restore()
	}
	t.restoreScreen()
	t.showCursor()
	t.log.Close()
	//t.fb.Close()
	return nil
}

// Size returns the terminal size
func (t *Terminal) Size() (Rows, Cols int) {
	return t.fb.Size()
}

func (t *Terminal) Sync() {
	t.fb.Sync()
}

// NextEvent return the nextevent in the pipe
func (t *Terminal) NextEvent() events.Event {
	return <-t.events
}

// Set changes the character display in the given row/col.
func (t *Terminal) Set(row, col int, ch rune, attrs ...interface{}) {
	t.fb.Set(row, col, ch, attrs...)
}

// Send send a command to the output
func (t *Terminal) Send(cmd string, args ...interface{}) {
	t.w.Write(commands.Sequence(cmd, args...))
}

func (t *Terminal) forwardInputEvents(c chan (events.Event)) {
	for e := range t.input.Events {
		c <- e
	}
	panic("input was closed")
}

func (t *Terminal) saveScreen() {
	t.Send("SMCUP")
}

func (t *Terminal) restoreScreen() {
	t.Send("RMCUP")
}

func (t *Terminal) hideCursor() {
	t.Send("HIDECURSOR")
}
func (t *Terminal) showCursor() {
	t.Send("SHOWCURSOR")
}

func (t *Terminal) sendWinSize(c chan (events.Event), rows, cols int) {
	c <- &events.WindowSizeEvent{
		Cols: cols,
		Rows: rows,
	}
}
