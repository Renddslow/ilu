package lexer

func Lex(input []byte) []*Token {
	rawTokens := make([]string, 0)

	inString := false
	tokenValue := ""

	for i := 0; i < len(input); i++ {
		if (input[i] == ' ' && !inString) || input[i] == '\n' {
			if tokenValue != "" {
				rawTokens = append(rawTokens, tokenValue)
			}
			tokenValue = ""
			continue
		}

		if input[i] == '\'' {
			if inString {
				rawTokens = append(rawTokens, tokenValue)
				tokenValue = ""
			}
			inString = !inString
			continue
		}

		if input[i] == '\\' {
			tokenValue += string(input[i+1])
			i++
			continue
		}

		tokenValue += string(input[i])
	}

	if tokenValue != "" {
		rawTokens = append(rawTokens, tokenValue)
	}

	tokens := make([]*Token, len(rawTokens))
	for _, t := range rawTokens {
		tokens = append(tokens, NewToken(t))
	}

	return tokens
}
