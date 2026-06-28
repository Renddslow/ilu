package parser

import (
	"slices"

	"github.com/Renddslow/ilu/internal/lexer"
	"github.com/Renddslow/ilu/internal/opcodes"
)

type Instruction interface {
	isInstruction()
}

type StringLiteral int32

type IntegerLiteral int32

type OpCode opcodes.Opcode

type TempLabelRef string

func (l StringLiteral) isInstruction()  {}
func (l IntegerLiteral) isInstruction() {}
func (l OpCode) isInstruction()         {}
func (l TempLabelRef) isInstruction()   {}

type Symbol struct {
	ID     int32
	Index  int32
	Length int32
}

func getNullaryOpFromToken(tok *lexer.Token) Instruction {
	switch tok.Value {
	case "printc":
		return OpCode(opcodes.Printc)
	case "eos":
		return OpCode(opcodes.EndOfStack)
	default:
		return nil
	}
}

func getUnaryFromTokens(tok *lexer.Token, arg *lexer.Token) []Instruction {
	instructions := make([]Instruction, 0)
	if tok.Value == "loads" {
		for _, char := range reverseString(arg.Value) {
			instructions = append(instructions, OpCode(opcodes.Push), StringLiteral(char))
		}
	}

	if tok.Value == "jmp" {
		instructions = append(instructions, OpCode(opcodes.Jump), TempLabelRef(arg.Value))
	}

	return instructions
}

func reverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
