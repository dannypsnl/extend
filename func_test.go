package extend_test

import (
	"testing"

	"github.com/dannypsnl/extend"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
)

func TestExtFunction_IsDefinition(t *testing.T) {
	f := extend.Func(ir.NewFunc("", types.Void))
	f.NewBlock("")
	assert.True(t, f.IsDefinition())
	assert.False(t, f.IsDeclaration())
}
func TestExtFunction_IsDeclaration(t *testing.T) {
	f := extend.Func(ir.NewFunc("", types.Void))
	assert.False(t, f.IsDefinition())
	assert.True(t, f.IsDeclaration())
}
