package lexer

import (
	"strings"
)

type TokenType int

const (
	Op TokenType = iota
	String
	LabelDef
	LabelRef
	Negation
	Done
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

var ops = []string{"loads", "printc", "jmp", "eos", "end"}

func containsOp(token string) bool {
	for _, op := range ops {
		// Ops can have a condition suffix
		if strings.HasPrefix(token, op) {
			return true
		}
	}
	return false
}

func NewToken(token string, line, col int) *Token {
	if containsOp(token) {
		return &Token{
			Type:   Op,
			Value:  token,
			Line:   line,
			Column: col,
		}
	}

	if strings.HasSuffix(token, ":") {
		return &Token{
			Type:   LabelDef,
			Value:  strings.TrimSuffix(token, ":"),
			Line:   line,
			Column: col,
		}
	}

	if strings.HasPrefix(token, "@") {
		return &Token{
			Type:   LabelRef,
			Value:  strings.TrimPrefix(token, "@"),
			Line:   line,
			Column: col,
		}
	}

	if token == "!" {
		return &Token{
			Type:   Negation,
			Value:  "",
			Line:   line,
			Column: col,
		}
	}

	return nil
}

func NewString(token string, line, col int) *Token {
	return &Token{
		Type:   String,
		Value:  token,
		Line:   line,
		Column: col,
	}
}
