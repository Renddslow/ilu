package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Renddslow/ilu/internal/lexer"
	"github.com/Renddslow/ilu/internal/parser"
	"github.com/Renddslow/ilu/internal/serialize"
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

	symbols := parser.SymbolTable{}

	rawOps, parseErr := parser.ParseFirstPass(tokens, &symbols)

	if parseErr != nil {
		fmt.Println(parseErr.Display(data, programAbsPath))
		os.Exit(1)
	}

	ops, parseErr := parser.ParseFinalPass(rawOps, symbols)

	if parseErr != nil {
		fmt.Println(parseErr.Display(data, programAbsPath))
		os.Exit(1)
	}

	out := bytes.NewBuffer(nil)
	out.WriteString("ILU")
	out.Write(symbols.Serialize())
	out.Write(serialize.Serialize(ops))

	cwd, _ := os.Getwd()
	outFile := path.Join(cwd, strings.Replace(path.Base(programAbsPath), ".ilu", "", 1))
	os.WriteFile(
		outFile,
		out.Bytes(),
		0644,
	)
}
