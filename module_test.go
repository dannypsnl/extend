package extend_test

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"

	"github.com/dannypsnl/extend"
)

func TestNoDupModule_AutoRenamingNewDefinition(t *testing.T) {
	m := extend.NoDup(ir.NewModule())
	foo1 := m.NewFunc("foo", types.Void)
	foo2 := m.NewGlobal("foo", types.Void)
	foo3 := m.NewGlobalDef("foo", constant.NewInt(types.I8, 1))

	assert.Equal(t, "foo", foo1.Name())
	assert.Equal(t, "foo.1", foo2.Name())
	assert.Equal(t, "foo.2", foo3.Name())
}

func TestNoDupModule_RemoveExistedDup(t *testing.T) {
	m := ir.NewModule()
	foo1 := m.NewGlobal("foo", types.Void)
	foo2 := m.NewGlobalDef("foo", constant.NewInt(types.I8, 1))
	foo3 := m.NewFunc("foo", types.Void)

	extend.NoDup(m)
	assert.Equal(t, "foo", foo1.Name())
	assert.Equal(t, "foo.1", foo2.Name())
	assert.Equal(t, "foo.2", foo3.Name())
}
