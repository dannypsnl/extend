package extend

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// NoDup receives a ir.Module and returns a NoDupModule
// it also removes existed duplicates
func NoDup(b *ir.Module) *NoDupModule {
	m := &NoDupModule{
		Module:    b,
		existedId: make(map[string]uint64),
	}
	for _, g := range b.Globals {
		g.SetName(m.autoNewName(g.Name()))
	}
	for _, f := range b.Funcs {
		f.SetName(m.autoNewName(f.Name()))
	}
	return m
}

// NoDupModule is the extension of ir.Module
// it wraps NewFunc, NewGlobal, and NewGlobalDef, automatically renaming existed global identifier
type NoDupModule struct {
	*ir.Module
	existedId map[string]uint64
}

func (e *NoDupModule) NewFunc(name string, ret types.Type, params ...*ir.Param) *ir.Func {
	return e.Module.NewFunc(e.autoNewName(name), ret, params...)
}

func (e *NoDupModule) NewGlobal(name string, ty types.Type) *ir.Global {
	return e.Module.NewGlobal(e.autoNewName(name), ty)
}

func (e *NoDupModule) NewGlobalDef(name string, c constant.Constant) *ir.Global {
	return e.Module.NewGlobalDef(e.autoNewName(name), c)
}

// autoNewName would automatically renaming duplicated global name via given name
func (e NoDupModule) autoNewName(name string) string {
	suffix := e.existedId[name]
	e.existedId[name]++
	if suffix != 0 {
		return fmt.Sprintf("%s.%d", name, suffix)
	} else {
		return name
	}
}
