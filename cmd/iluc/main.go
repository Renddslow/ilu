package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Renddslow/ilu/internal/lexer"
	"github.com/Renddslow/ilu/internal/parser"
)

func main() {
	programPath := os.Args[1]
	programAbsPath, _ := filepath.Abs(programPath)

	data, err := os.ReadFile(programPath)
	if err != nil {
		panic(err)
	}

	tokens, lexErr := lexer.Lex(data)

	if lexErr != nil {
		fmt.Println(lexErr.Display(data, programAbsPath))
		os.Exit(1)
	}

	_, parseErr := parser.ParseFirstPass(tokens)

	if parseErr != nil {
		fmt.Println(parseErr.Display(data, programAbsPath))
		os.Exit(1)
	}
}
