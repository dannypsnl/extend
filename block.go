package extend

import (
	"github.com/llir/llvm/ir"
)

// Block receives a ir.Block and returns its extension
func Block(b *ir.Block) *ExtBlock {
	return &ExtBlock{
		Block: b,
	}
}

// ExtBlock is the extension of ir.Block
type ExtBlock struct {
	*ir.Block
}

// HasTerminator checks the end of Block is a terminator or not
func (e *ExtBlock) HasTerminator() bool {
	return e.Term != nil
}
