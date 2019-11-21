// Package terminal abstract the terminal by providing two main interfaces: Input Events and a framebuffer output
package terminal

import (
	"github.com/guillermo/reacty/events"
	"github.com/guillermo/reacty/framebuffer"
	"github.com/guillermo/reacty/input"
	"github.com/guillermo/reacty/commands"
	"github.com/tredoe/term/sys"
	"io"
	"os"
	"os/signal"
	"syscall"
)

// Terminal holds the state of the current terminal
type Terminal struct {
	r  io.Reader
	fd int

	w                   io.Writer
	input               *input.Input
	fb                  *framebuffer.Framebuffer
	oldState, lastState sys.Termios
	events              <-chan (events.Event)
	rows, cols          int
}

// Open configures the controlling tty to be used with Terminal
func Open() (*Terminal, error) {
	eventsChan := make(chan (events.Event), 1024)
	t := &Terminal{
		r:  os.Stdin,
		w:  os.Stdout,
		fd: int(os.Stdin.Fd()),
		fb: &framebuffer.Framebuffer{
			Output: os.Stdout,
		},
		input:  input.Open(os.Stdin),
		events: eventsChan,
	}

	if err := t.saveTerminalState(); err != nil {
		return nil, err
	}

	if err := t.rawMode(); err != nil {
		return nil, err
	}

	go t.forwardInputEvents(eventsChan)
	go t.detectResize(eventsChan)

	t.saveScreen()
	t.hideCursor()
	t.getWinSize(eventsChan)

	go t.fb.AutoSync()
	return t, nil
}

// Close resets the terminal to the previous state
func (t *Terminal) Close() error {
	t.fb.StopAutoSync()
	t.Send("RMCUP")
	t.Send("SHOWCURSOR")
	return t.restore()
}

// Dimensions returns the size of the terminal
func (t *Terminal) Dimensions() (Rows, Cols int) {
	return t.rows, t.cols
}

// NextEvent return the nextevent in the pipe
func (t *Terminal) NextEvent() events.Event {
	return <-t.events
}

// Set changes the character display in the given row/col.
func (t *Terminal) Set(row, col int, ch rune) {
	t.fb.Set(row, col, ch)
}

// Send send a command to the output
func (t *Terminal) Send(cmd string, args ...interface{}) {
	t.w.Write(commands.Commands[cmd].Sequence(args...))
}

func (t *Terminal) detectResize(c chan (events.Event)) {
	sigChan := make(chan (os.Signal))
	signal.Notify(sigChan, syscall.SIGWINCH)
	for range sigChan {
		t.getWinSize(c)
	}
}

func (t *Terminal) forwardInputEvents(c chan (events.Event)) {
	for e := range t.input.Events {
		c <- e
	}
	panic("input was closed")
}

func (t *Terminal) saveTerminalState() error {
	if err := sys.Getattr(t.fd, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Getattr", err)
	}
	t.oldState = t.lastState
	return nil
}

func (t *Terminal) saveScreen() {
	t.Send("SMCUP")
}

func (t *Terminal) hideCursor() {
	t.Send("HIDECURSOR")
}

func (t *Terminal) getWinSize(c chan (events.Event)) error {

	ws := sys.Winsize{}
	if err := sys.GetWinsize(t.fd, &ws); err != nil {
		panic(err)
	}

	c <- &events.WindowSizeEvent{
		Cols: int(ws.Col),
		Rows: int(ws.Row),
	}
	t.cols = int(ws.Col)
	t.rows = int(ws.Row)
	t.fb.SetSize(t.rows, t.cols)

	return nil
}

func (t *Terminal) rawMode() error {
	// Input modes - no break, no CR to NL, no NL to CR, no carriage return,
	// no strip char, no start/stop output control, no parity check.
	t.lastState.Iflag &^= (sys.BRKINT | sys.IGNBRK | sys.ICRNL | sys.INLCR |
		sys.IGNCR | sys.ISTRIP | sys.IXON | sys.PARMRK)

	// Output modes - disable post processing.
	t.lastState.Oflag &^= sys.OPOST

	// Local modes - echoing off, canonical off, no extended functions,
	// no signal chars (^Z,^C).
	t.lastState.Lflag &^= (sys.ECHO | sys.ECHONL | sys.ICANON | sys.IEXTEN | sys.ISIG)

	// Control modes - set 8 bit chars.
	t.lastState.Cflag &^= (sys.CSIZE | sys.PARENB)
	t.lastState.Cflag |= sys.CS8

	// Control chars - set return condition: min number of bytes and timer.
	// We want read to return every single byte, without timeout.
	t.lastState.Cc[sys.VMIN] = 1 // Read returns when one char is available.
	t.lastState.Cc[sys.VTIME] = 0

	// Put the terminal in raw mode after flushing
	if err := sys.Setattr(t.fd, sys.TCSAFLUSH, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	return nil
}

func (t *Terminal) restore() error {
	if err := sys.Setattr(t.fd, sys.TCSANOW, &t.oldState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	t.lastState = t.oldState
	return nil
}
