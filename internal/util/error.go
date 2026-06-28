package util

import (
	"fmt"
	"strings"
)

type CompilerStage int

const (
	Lexing CompilerStage = iota
	Parsing
)

func (s CompilerStage) String() string {
	switch s {
	case Lexing:
		return "lexing"
	case Parsing:
		return "parsing"
	default:
		panic("unhandled default case")
	}
}

type Error struct {
	Stage   CompilerStage
	Message string
	Line    int
	Column  int
}

func NewLexError(msg string, line int, column int) *Error {
	return &Error{
		Stage:   Lexing,
		Message: msg,
		Line:    line,
		Column:  column,
	}
}

func NewParseError(msg string, line int, column int) *Error {
	return &Error{
		Stage:   Parsing,
		Message: msg,
		Line:    line,
		Column:  column,
	}
}

func (err *Error) Display(file []byte, path string) string {
	lines := strings.Split(string(file), "\n")
	line := lines[err.Line]

	errLine := ""
	for range err.Column - 1 {
		errLine += " "
	}
	errLine += "\x1b[31m^\x1b[39m"

	start := max(0, err.Line-3)
	end := min(len(lines), err.Line+1+4)

	prefix := lines[start:err.Line]
	suffix := lines[err.Line+1 : end]

	output := make([]string, 0, len(prefix)+len(suffix)+2+3)
	output = append(output, prefix...)
	output = append(output, line, errLine)
	output = append(output, suffix...)
	output = append(
		output,
		"",
		fmt.Sprintf("%s error: %s", err.Stage, err.Message),
		fmt.Sprintf("    \x1b[34m%s:%d:%d\x1b[39m", path, err.Line, err.Column),
	)

	return strings.Join(output, "\n")
}
