package extend_test

import (
	"testing"

	"github.com/dannypsnl/extend"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
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

func TestExtBlock_PhiInStartAutomatically(t *testing.T) {
	b := extend.Block(ir.NewBlock(""))
	x := b.NewAdd(constant.NewInt(types.I1, 1), constant.NewInt(types.I1, 2))
	y := b.NewAdd(constant.NewInt(types.I1, 1), constant.NewInt(types.I1, 2))
	b.NewPhi(
		ir.NewIncoming(x, b.Block),
		ir.NewIncoming(y, b.Block),
	)
	assert.NotNil(t, b.Block.Insts[0].(*ir.InstPhi))
	assert.NotNil(t, b.Block.Insts[1].(*ir.InstAdd))
}
