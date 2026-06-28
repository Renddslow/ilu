package opcodes

type Opcode uint8

const (
	Push Opcode = iota
	Printc
	EndOfStack
	Jump
)
