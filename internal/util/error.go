package util

import (
	"fmt"
	"strings"
)

type LexError struct {
	Message string
	Line    int
	Column  int
}

func NewLexError(msg string, line int, column int) *LexError {
	return &LexError{
		Message: msg,
		Line:    line,
		Column:  column,
	}
}

func (err *LexError) Display(file []byte, path string) string {
	lines := strings.Split(string(file), "\n")
	line := lines[err.Line]

	errLine := ""
	for range err.Column - 1 {
		errLine += " "
	}
	errLine += "\x1b[31m^\x1b[39m"

	output := []string{
		line,
		errLine,
		"",
		fmt.Sprintf("lexing error: %s", err.Message),
		fmt.Sprintf("    \x1b[34m%s:%d:%d\x1b[39m", path, err.Line, err.Column),
	}
	return strings.Join(output, "\n")
}
