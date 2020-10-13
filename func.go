package extend

import (
	"github.com/llir/llvm/ir"
)

// ExtFunc is extension of ir.Func
type ExtFunc struct {
	*ir.Func
}

// Func extends ir.Func
func Func(f *ir.Func) *ExtFunc {
	return &ExtFunc{
		Func: f,
	}
}

// IsDefinition returns true if this is a definition
func (e *ExtFunc) IsDefinition() bool {
	return e.Blocks != nil
}

// IsDeclaration returns true if this is a declaration
func (e *ExtFunc) IsDeclaration() bool {
	return e.Blocks == nil
}
