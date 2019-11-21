package output

import (
	"errors"
	"io"
)

// Commander represents
type Commander interface {
	Sequence() []byte
}

// Output wraps an io.Writer and send commands through the Commander interface.
type Output struct {
	w io.Writer
}

// Open returns and output that represents the terminal output
func Open(w io.Writer) *Output {
	return &Output{w: w}
}

// Write a command to the io.Writer
func (o *Output) Write(c []byte) (n int, err error) {
	return io.Writer(o.w).Write(c)
}

// Send is a shortcut to send commands to the output
func (o *Output) Send(cmdName string, args ...interface{}) error {
	c, ok := Commands[cmdName]
	if !ok {
		return errors.New("Command not found")
	}

	_, err := o.Write(c.Sequence(args...))
	return err
}
