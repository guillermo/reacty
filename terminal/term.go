package terminal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tredoe/term/sys"
)


type term struct {
	Fd                  int
	oldState, lastState sys.Termios
}

func (t *term) saveTerminalState() error {
	if err := sys.Getattr(t.Fd, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Getattr", err)
	}
	t.oldState = t.lastState
	return nil
}
func (t *term) restore() error {
	if err := sys.Setattr(t.Fd, sys.TCSANOW, &t.oldState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	t.lastState = t.oldState
	return nil
}
func (t *term) winSize() (rows, cols int, err error) {
	ws := sys.Winsize{}
	if err := sys.GetWinsize(t.Fd, &ws); err != nil {
		return 0,0,err
	}

	return int(ws.Row), int(ws.Col), nil
}

func (t *term) onResize(cbk func(rows, cols int)) error {
	sync := func() error {
		rows, cols, err := t.winSize()
		if err != nil {
			return err
		}
		cbk(rows, cols)
		return nil
	}

	sigChan := make(chan (os.Signal))
	signal.Notify(sigChan, syscall.SIGWINCH)
	go func() {
		for range sigChan {
			sync()
		}
	}()
	return sync()

}

func (t *term) rawMode() error {
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
	if err := sys.Setattr(t.Fd, sys.TCSAFLUSH, &t.lastState); err != nil {
		return os.NewSyscallError("sys.Setattr", err)
	}
	return nil
}
