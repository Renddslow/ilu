package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"slices"

	"github.com/Renddslow/ilu/internal/lexer"
	"github.com/Renddslow/ilu/internal/util"
)

type SymbolTable map[string]Symbol

func (s SymbolTable) Serialize() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteByte(byte(len(s) * 12))

	for _, symbol := range s {
		binary.Write(buf, binary.BigEndian, symbol.ID)
		binary.Write(buf, binary.BigEndian, symbol.Index)
		binary.Write(buf, binary.BigEndian, symbol.Length)
	}

	return buf.Bytes()
}

var nullaryOps = []string{"printc", "eos"}
var unaryOps = []string{"loads", "jmp"}

func ParseFirstPass(tokens []*lexer.Token, table *SymbolTable) ([]Instruction, *util.Error) {
	instructions := make([]Instruction, 0)

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
			childInstructions, err := ParseFirstPass(tokens[i+1:i+doneIdx+1], table)

			if err != nil {
				return nil, err
			}

			hash := fnv.New32()
			hash.Write([]byte(token.Value))
			(*table)[token.Value] = Symbol{
				ID:     int32(hash.Sum32()),
				Index:  int32(len(instructions)),
				Length: int32(len(childInstructions)),
			}
			instructions = append(instructions, childInstructions...)
			i += doneIdx + 1
			continue
		}

		if token.Type == lexer.Op {
			if slices.Contains(nullaryOps, token.Value) {
				instructions = append(instructions, getNullaryOpFromToken(token))
				continue
			}

			if slices.Contains(unaryOps, token.Value) {
				instructions = append(instructions, getUnaryFromTokens(token, tokens[i+1])...)
				i += 1
				continue
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

func ParseFinalPass(instructions []Instruction, table SymbolTable) ([]Instruction, *util.Error) {
	finalInstructions := make([]Instruction, 0, len(instructions))
	for _, instruction := range instructions {
		if label, ok := instruction.(TempLabelRef); ok {
			finalInstructions = append(finalInstructions, IntegerLiteral(table[string(label)].ID))
			continue
		}
		finalInstructions = append(finalInstructions, instruction)
	}
	return finalInstructions, nil
}
