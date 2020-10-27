package extend

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Block receives a ir.Block and returns its extension
func Block(b *ir.Block) *ExtBlock {
	return &ExtBlock{
		Block: b,
	}
}

// ExtBlock is the extension of ir.Block, except several helper functions, it would
// * panic if users trying to add second terminator to block
type ExtBlock struct {
	*ir.Block
}

// HasTerminator checks the end of Block is a terminator or not
func (e *ExtBlock) HasTerminator() bool {
	return e.Term != nil
}

// BelongsToFunc checks the parent function exists or not
func (e *ExtBlock) BelongsToFunc() bool {
	return e.Parent != nil
}

func (e *ExtBlock) setTerminator(terminator ir.Terminator) {
	if e.HasTerminator() {
		panic("Block already has terminator")
	}
	e.Term = terminator
}
func (e *ExtBlock) NewRet(x value.Value) *ir.TermRet {
	term := ir.NewRet(x)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewBr(target *ir.Block) *ir.TermBr {
	term := ir.NewBr(target)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewCondBr(cond value.Value, targetTrue, targetFalse *ir.Block) *ir.TermCondBr {
	term := ir.NewCondBr(cond, targetTrue, targetFalse)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewSwitch(x value.Value, targetDefault *ir.Block, cases ...*ir.Case) *ir.TermSwitch {
	term := ir.NewSwitch(x, targetDefault, cases...)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewIndirectBr(addr constant.Constant, validTargets ...*ir.Block) *ir.TermIndirectBr {
	term := ir.NewIndirectBr(addr, validTargets...)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewInvoke(invokee value.Value, args []value.Value, normalRetTarget, exceptionRetTarget *ir.Block) *ir.TermInvoke {
	term := ir.NewInvoke(invokee, args, normalRetTarget, exceptionRetTarget)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewCallBr(callee value.Value, args []value.Value, normalRetTarget *ir.Block, otherRetTargets ...*ir.Block) *ir.TermCallBr {
	term := ir.NewCallBr(callee, args, normalRetTarget, otherRetTargets...)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewResume(x value.Value) *ir.TermResume {
	term := ir.NewResume(x)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewCatchSwitch(parentPad ir.ExceptionPad, handlers []*ir.Block, defaultUnwindTarget *ir.Block) *ir.TermCatchSwitch {
	term := ir.NewCatchSwitch(parentPad, handlers, defaultUnwindTarget)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewCatchRet(catchPad *ir.InstCatchPad, target *ir.Block) *ir.TermCatchRet {
	term := ir.NewCatchRet(catchPad, target)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewCleanupRet(cleanupPad *ir.InstCleanupPad, unwindTarget *ir.Block) *ir.TermCleanupRet {
	term := ir.NewCleanupRet(cleanupPad, unwindTarget)
	e.setTerminator(term)
	return term
}
func (e *ExtBlock) NewUnreachable() *ir.TermUnreachable {
	term := ir.NewUnreachable()
	e.setTerminator(term)
	return term
}
