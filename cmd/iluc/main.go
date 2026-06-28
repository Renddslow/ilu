package main

import (
	"os"

	"github.com/Renddslow/ilu/internal/lexer"
)

func main() {
	programPath := os.Args[1]
	data, err := os.ReadFile(programPath)
	if err != nil {
		panic(err)
	}
	lexer.Lex(data)
}
