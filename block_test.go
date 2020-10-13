package extend_test

import (
	"testing"

	"github.com/dannypsnl/extend"
	"github.com/llir/llvm/ir"
	"github.com/stretchr/testify/assert"
)

func TestExtBlock_IsEndWithTerminator(t *testing.T) {
	b := extend.Block(ir.NewBlock(""))
	assert.False(t, b.HasTerminator())
	b.NewRet(nil)
	assert.True(t, b.HasTerminator())
}
