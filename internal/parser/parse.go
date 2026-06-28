package parser

import (
	"fmt"

	"github.com/Renddslow/ilu/internal/lexer"
	"github.com/Renddslow/ilu/internal/opcodes"
	"github.com/Renddslow/ilu/internal/util"
)

type Instruction struct {
	Op opcodes.Opcode
}

type Symbol struct {
	Index  int
	Length int
}

type SymbolTable map[string]Symbol

func ParseFirstPass(tokens []*lexer.Token) ([]*Instruction, *util.Error) {
	instructions := make([]*Instruction, 0)
	//table := SymbolTable{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type == lexer.LabelDef {
			doneIdx := indexOfType(tokens[i:], lexer.Done)
			if doneIdx == -1 {
				return nil, util.NewParseError(
					fmt.Sprintf("label %s not terminated with an `end` instruction", token.Value),
					token.Line,
					token.Column,
				)
			}
		}
	}

	return instructions, nil
}

func indexOfType(tokens []*lexer.Token, tokenType lexer.TokenType) int {
	for i, token := range tokens {
		if token.Type == tokenType {
			return i
		}
	}
	return -1
}
