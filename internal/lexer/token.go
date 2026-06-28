package lexer

import "fmt"

type Token struct{}

func NewToken(token string) *Token {
	fmt.Println(token)
	return &Token{}
}
