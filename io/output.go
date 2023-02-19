package io

import (
	"os"
	"bytes"

	"github.com/jobutterfly/golitaire/screen"
)

func DrawRows(buf *bytes.Buffer) error {
	for i := 0; i < screen.S.Rows; i++ {
		if _, err := buf.Write([]byte("\x1b[K")); err != nil {
			return err
		}
		if _, err := buf.Write([]byte("\r\n")); err != nil {
			return err
		}
	}
	return nil
}

func DrawStatusBar(buf *bytes.Buffer) error {
	if _, err := buf.Write([]byte("\x1b[7m")); err != nil {
		return err
	}

	status := "golitaire"
	length := len(status)
	if length > screen.S.Cols {
		length = screen.S.Cols
		status = status[:length]
	}

	if _, err := buf.Write([]byte(status)); err != nil {
		return err
	}

	for i := length; i < screen.S.Cols;i++ {
		if _, err := buf.Write([]byte(" ")); err != nil {
			return err
		}
	}

	if _, err := buf.Write([]byte("\x1b[m")); err != nil {
		return err
	}
	if _, err := buf.Write([]byte("\r\n")); err != nil {
		return err
	}

	return nil
}



func RefreshScreen() error {

	var mainBuf bytes.Buffer

	if _, err := mainBuf.Write([]byte("\x1b[?25l")); err != nil {
		return err
	}
	if _, err := mainBuf.Write([]byte("\x1b[H")); err != nil {
		return err
	}

	if err := DrawRows(&mainBuf); err != nil {
		return err
	}

	if err := DrawStatusBar(&mainBuf); err != nil {
		return err
	}
	/*

	if err := DrawMessageBar(&mainBuf); err != nil {
		return err
	}
	*/

	if _, err := mainBuf.Write([]byte("\x1b[H")); err != nil {
		return err
	}

	if _, err := mainBuf.Write([]byte("\x1b[?25h")); err != nil {
		return err
	}

	if _, err := os.Stdout.Write(mainBuf.Bytes()); err != nil {
		return err

	}

	return nil
}
