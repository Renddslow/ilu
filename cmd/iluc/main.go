package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Renddslow/ilu/internal/lexer"
)

func main() {
	programPath := os.Args[1]
	data, err := os.ReadFile(programPath)
	if err != nil {
		panic(err)
	}

	_, lexErr := lexer.Lex(data)

	if lexErr != nil {
		p, _ := filepath.Abs(programPath)
		fmt.Println(lexErr.Display(data, p))
		os.Exit(1)
	}
}
