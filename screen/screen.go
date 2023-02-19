package screen

import (
	"golang.org/x/term"
)

type Screen struct {
	Rows      int
	Cols      int
	Cx        int
	Cy        int
	Termios   *term.State
	QuitTimes int
}

var S Screen
