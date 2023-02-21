package screen

import (
	"bytes"

	"golang.org/x/term"
)

type Screen struct {
	Rows      int
	Cols      int
	Cx        int
	Cy        int
	Termios   *term.State
	QuitTimes int
	GridRows  []*bytes.Buffer
}

var S Screen
