package main

import (
	"os"

	"golang.org/x/term"
	"github.com/jobutterfly/golitaire/screen"
	"github.com/jobutterfly/golitaire/terminal"
	"github.com/jobutterfly/golitaire/io"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		terminal.Die(err)
	}
	screen.S.Termios = oldState
	defer term.Restore(int(os.Stdin.Fd()), screen.S.Termios)
	terminal.InitEditor()

	// io.SetStatusMsg("HELP: Ctrl-Q = quit | Ctrl-S = save | Ctrl-F = find")

	for {
		if err := io.RefreshScreen(); err != nil {
			terminal.Die(err)
		}
		if io.ProcessKeyPress(oldState) {
			break
		}
	}
}
