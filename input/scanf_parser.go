package input

import (
	"fmt"
	"github.com/guillermo/reacty/events"
)

//https://play.golang.org/p/dudzyAar-Wa
func scanParser(data []rune) (e events.Event, n int, more bool) {

	var a, b, c int
	n, err := fmt.Sscanf(string(data), "\x1b[%d;%d;%dt", &a, &b, &c)
	if err == nil {
		return &events.WindowSizeEvent{}, len(data), false
	}
	errMsg := err.Error()
	if errMsg == "input does not match format" ||
		errMsg == "expected integer" {
		return nil, 0, false
	}
	if errMsg == "unexpected EOF" ||
		errMsg == "EOF" {
		return nil, 0, true
	}
	panic(err)

}
