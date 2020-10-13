package extend

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// Module receives a ir.Module and returns its extension
func Module(b *ir.Module) *ExtModule {
	return &ExtModule{
		Module:      b,
		existedFunc: make(map[string]uint64),
	}
}

// ExtModule is the extension of ir.Module
type ExtModule struct {
	*ir.Module
	existedFunc map[string]uint64
}

// NewFunc
func (e *ExtModule) NewFunc(name string, ret types.Type, params ...*ir.Param) *ir.Func {
	return e.Module.NewFunc(e.autoNewName(name), ret, params...)
}

// NewGlobal
func (e *ExtModule) NewGlobal(name string, ty types.Type) *ir.Global {
	return e.Module.NewGlobal(e.autoNewName(name), ty)
}

// NewGlobalDef
func (e *ExtModule) NewGlobalDef(name string, c constant.Constant) *ir.Global {
	return e.Module.NewGlobalDef(e.autoNewName(name), c)
}

// autoNewName would automatically renaming duplicated global name via given name
func (e ExtModule) autoNewName(name string) string {
	suffix := e.existedFunc[name]
	e.existedFunc[name]++
	if suffix != 0 {
		return fmt.Sprintf("%s.%d", name, suffix)
	} else {
		return name
	}
}
