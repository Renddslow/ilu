package opcodes

type Opcode uint8

const (
	Push Opcode = iota
	Printc
	NotZero
	Jump
)

// loads -> push ! push d ... push H
// nz
