// Package faketerm implements a fake terminal. Mostly intended for testing.
package faketerm

import (
	"fmt"
	"github.com/guillermo/reacty/events"
	"strings"
	"testing"
)

type FakeTerm struct {
	Rows       []string
	Events     chan (events.Event)
	rows, cols int
}

func New(rows int, cols int) *FakeTerm {
	return &FakeTerm{
		Events: make(chan (events.Event)),
		Rows:   make([]string, rows, rows),
		rows:   rows,
		cols:   cols,
	}
}

func (t *FakeTerm) Size() (int, int) {
	return t.rows, t.cols
}

func (t *FakeTerm) NextEvent() events.Event {
	return <-t.Events
}

func (t *FakeTerm) get(Row, Col int) rune {
	if Row < 1 || Col < 1 || Row > t.rows || Col > t.cols {
		panic("get outside terminal size")
	}
	if len(t.Rows) < Row {
		return ' '
	}

	row := t.Rows[Row-1]
	if len([]rune(row)) < Col {
		return ' '
	}
	return []rune(row)[Col-1]
}

func (t *FakeTerm) Set(Row, Col int, ch rune, attrs ...interface{}) {
	if Row <= 0 {
		panic("Got a no row")
	}
	if Col <= 0 {
		panic("Got a no col")
	}
	for len(t.Rows) < Row {
		t.Rows = append(t.Rows, "")
	}
	for len(t.Rows[Row-1]) < Col {
		t.Rows[Row-1] += " "
	}

	row := t.Rows[Row-1]
	s := row[:Col-1] + string(ch) + row[Col:]
	t.Rows[Row-1] = s
	if t.rows < Row {
		t.rows = Row
	}
	if t.cols < Col {
		t.cols = Col
	}
}
func (t *FakeTerm) String() string {
	s := "\n"
	for _, row := range t.Rows {
		s += row + "\n"
	}
	return s
}

func (t *FakeTerm) Sync() {
	for _, l := range t.Lines() {
		fmt.Println(strings.TrimSpace(l))
	}
	return
}

func (t *FakeTerm) Lines() []string {
	rows := []string{}
	for row := 0; row < t.rows; row++ {
		s := []rune{}
		for col := 0; col < t.cols; col++ {
			s = append(s, t.get(row+1, col+1))
		}
		rows = append(rows, string(s))
	}
	return rows
}

func (t *FakeTerm) Equal(test *testing.T, rows ...string) {
	test.Helper()
	lines := t.Lines()
	if len(lines) != len(rows) {
		test.Errorf("Expected %d rows. Got %d", len(rows), len(lines))
	}
	for i, row := range rows {
		if row != lines[i] {
			test.Errorf("Expected row %d to be %q. Got: %q", i+1, rows[i], lines[i])
		}
	}
}
