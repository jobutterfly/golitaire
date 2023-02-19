package io

import (
	"os"

	"golang.org/x/term"
	"github.com/jobutterfly/golitaire/consts"
	"github.com/jobutterfly/golitaire/terminal"
	"github.com/jobutterfly/golitaire/screen"
	/*
	"github.com/jobutterfly/gowrite/editor"
	"github.com/jobutterfly/gowrite/operations"
	*/
)


func ProcessKeyPress(oldState *term.State) bool {
	var char int = terminal.ReadKey()

	switch char {
	// see references in readme for ascii control codes
	// ctrl q
	case 17:
		if screen.S.QuitTimes > 0 {
			// SetStatusMsg(fmt.Sprintf("Warning! File has unsaved changes. Press Ctrl-Q %d times to quit anyway.", editor.E.QuitTimes))
			screen.S.QuitTimes--
			return false
		}
		if _, err := os.Stdout.Write([]byte("\x1b[2J")); err != nil {
			terminal.Die(err)
		}
		if _, err := os.Stdout.Write([]byte("\x1b[H")); err != nil {
			terminal.Die(err)
		}
		return true
	// ctrl s
	/*
	case 19: 
		EditorSave()
	case consts.UP:
		MoveCursor(char)
	case consts.DOWN:
		MoveCursor(char)
	case consts.RIGHT:
		MoveCursor(char)
	case consts.LEFT:
		MoveCursor(char)
	*/
	}

	screen.S.QuitTimes = consts.QUIT_TIMES 
	return false
}





