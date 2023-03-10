package terminal

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
	"golang.org/x/term"
	"github.com/jobutterfly/golitaire/screen"
	"github.com/jobutterfly/golitaire/consts"
)


func Die(err error) {
	term.Restore(int(os.Stdin.Fd()), screen.S.Termios)
	if _, err := os.Stdout.Write([]byte("\x1b[2J")); err != nil {
		log.Fatal("Could not clean screen")
	}
	fmt.Printf("%v\n", err)
	os.Exit(1)
}

func ReadKey() int {
	var b []byte = make([]byte, 1)

	nread, err := os.Stdin.Read(b)
	if err != nil {
		Die(err)
	}

	if nread != 1 {
		Die(errors.New(fmt.Sprintf("Wanted to read one character, got %d", nread)))
	}

	if b[0] == '\x1b' {
		var seq []byte = make([]byte, 3)

		_, err := os.Stdin.Read(seq)
		if err != nil {
			Die(err)
		}

		if seq[0] == '[' {
			if seq[1] >= '0' && seq[1] <= '9' {
				if seq[2] == '~' {
					switch seq[1] {
					case '1':
						return consts.HOME
					case '3':
						return consts.DELETE
					case '4':
						return consts.END
					case '5':
						return consts.PAGE_UP
					case '6':
						return consts.PAGE_DOWN
					case '7':
						return consts.HOME
					case '8':
						return consts.END
					}
				}
			} else {
				switch seq[1] {
				case 'A':
					return consts.UP
				case 'B':
					return consts.DOWN
				case 'C':
					return consts.RIGHT
				case 'D':
					return consts.LEFT
				case 'H':
					return consts.HOME
				case 'F':
					return consts.END
				}

			}
		} else if seq[0] == 'O' {
			switch seq[1] {
			case 'H':
				return consts.HOME
			case 'F':
				return consts.END
			}
		}

		return '\x1b'
	}

	return int(b[0])
}

func GetCursorPosition() (row int, col int, err error) {
	var buf []byte = make([]byte, 32)
	var i int = 0
	if _, err := os.Stdout.Write([]byte("\x1b[6n")); err != nil {
		return 0, 0, err
	}

	if _, err := os.Stdin.Read(buf); err != nil {
		return 0, 0, err
	}

	for i < len(buf) {
		if buf[i] == 'R' {
			break
		}
		i++
	}

	if buf[0] != '\x1b' || buf[1] != '[' {
		return 0, 0, errors.New("Could not find escape sequence when getting cursor position")
	}

	if _, err := fmt.Sscanf(string(buf[2:i]), "%d;%d", &row, &col); err != nil {
		return 0, 0, err
	}

	return row, col, nil
}

func GetWindowSize() (rows int, cols int, err error) {
	winSize, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}

	if winSize.Row < 1 || winSize.Col < 1 {
		_, err := os.Stdout.Write([]byte("\x1b[999C\x1b[999B"))
		if err != nil {
			return 0, 0, err
		}

		return GetCursorPosition()
	}
	return int(winSize.Row), int(winSize.Col), nil
}

func MakeGrid() []*bytes.Buffer {
	colLength := screen.S.Cols / 7
	rowLength := screen.S.Rows / 3
	var b []byte
	var line []byte
	var a []*bytes.Buffer

	for i := 0; i < screen.S.Cols; i++ {
		line = append(line, '-')
		if i % colLength == 0 {
			b = append(b, '|')
		} else {
			b = append(b, ' ')
		}
	}

	buf := bytes.NewBuffer(b)
	bufLine := bytes.NewBuffer(line)

	for j := 0; j < screen.S.Rows; j++ {
		if j % rowLength == 0 {
			a = append(a, bufLine)
		} else {
			a = append(a, buf)
		}
	}

	return a
}

func InitEditor() {
	screen.S.Cx = 0
	screen.S.Cy = 0
	rows, cols, err := GetWindowSize()
	if err != nil {
		Die(err)
	}
	screen.S.Rows = rows
	screen.S.Cols = cols
	screen.S.GridRows = MakeGrid()
}

