package serialize

import (
	"bytes"
	"encoding/binary"

	"github.com/Renddslow/ilu/internal/parser"
)

func Serialize(instructions []parser.Instruction) []byte {
	buf := bytes.NewBuffer(nil)
	for _, instruction := range instructions {
		switch inst := instruction.(type) {
		case parser.StringLiteral:
			if inst < 255 {
				buf.WriteByte(byte(inst))
			} else {
				binary.Write(buf, binary.BigEndian, inst)
			}
		case parser.IntegerLiteral:
			if inst < 255 {
				buf.WriteByte(byte(inst))
			} else {
				binary.Write(buf, binary.BigEndian, inst)
			}
		case parser.OpCode:
			buf.WriteByte(byte(inst))
		}
	}
	return buf.Bytes()
}
