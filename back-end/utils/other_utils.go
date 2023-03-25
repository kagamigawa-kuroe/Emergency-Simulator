package utils

import (
	"bytes"
	"errors"
	"io"
	"runtime"
	"strconv"
	"strings"
)

func SubStr(str string, start, length int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, length)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

func SubstrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos == true {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if hasPos == false {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
	}
}

func GetGoroutineId() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := string(buf[:n])

	str := stk[10:11]

	return strconv.Atoi(str)
}