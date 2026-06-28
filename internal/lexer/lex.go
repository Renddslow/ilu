package lexer

import (
	"fmt"

	"github.com/Renddslow/ilu/internal/util"
)

func Lex(input []byte) ([]*Token, *util.Error) {
	tokens := make([]*Token, 0)
	line, col := 0, 0

	word := ""
	inString := false
	for i := 0; i < len(input); i++ {
		char := input[i]

		if char == '\n' {
			if inString {
				return nil, util.NewLexError("unterminated string", line, col)
			}

			if word != "" {
				token := NewToken(word, line, col)
				if token == nil {
					return nil, util.NewLexError(fmt.Sprintf("unrecognized keyword %s", word), line, col)
				}
				tokens = append(tokens, token)
			}

			line++
			col = 0
			word = ""
			continue
		}

		if char == ' ' && !inString {
			if word != "" {
				token := NewToken(word, line, col)
				if token == nil {
					return nil, util.NewLexError(fmt.Sprintf("unrecognized instruction %s", word), line, col)
				}
				tokens = append(tokens, NewToken(word, line, col))
			}
			word = ""
			col++
			continue
		}

		if char == '!' && !inString {
			tokens = append(tokens, NewToken("!", line, col))
			col++
			continue
		}

		if inString && char == '\'' {
			inString = false
			tokens = append(tokens, NewString(word, line, col))
			word = ""
			col++
			continue
		}

		if !inString && char == '\'' {
			inString = true
			col++
			continue
		}

		word += string(char)
		col++
	}

	if inString {
		return nil, util.NewLexError("unterminated string", line, col)
	}

	if word != "" {
		token := NewToken(word, line, col)
		if token == nil {
			return nil, util.NewLexError(fmt.Sprintf("unrecognized instruction `%s`", word), line, col)
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
