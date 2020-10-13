package extend_test

import (
	"testing"

	"github.com/dannypsnl/extend"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
)

func TestExtBlock_HasTerminator(t *testing.T) {
	b := extend.Block(ir.NewBlock(""))
	assert.False(t, b.HasTerminator())
	b.NewRet(nil)
	assert.True(t, b.HasTerminator())
}

func TestExtBlock_BelongsToFunc_True(t *testing.T) {
	f := ir.NewFunc("", types.Void)
	b := extend.Block(f.NewBlock(""))
	assert.True(t, b.BelongsToFunc())
}
func TestExtBlock_BelongsToFunc_False(t *testing.T) {
	b := extend.Block(ir.NewBlock(""))
	assert.False(t, b.BelongsToFunc())
}
