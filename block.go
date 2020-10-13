package extend

import (
	"github.com/llir/llvm/ir"
)

type extBlock struct {
	*ir.Block
}

func (e *extBlock) HasTerminator() bool {
	return e.Term != nil
}

func Block(b *ir.Block) *extBlock {
	return &extBlock{
		Block: b,
	}
}
