package extend

import (
	"fmt"
	"strings"

	"github.com/dannypsnl/extend/enc"

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
// * automatically lift Phi instructions to the front of the block
type ExtBlock struct {
	*ir.Block
	Phis []*ir.InstPhi
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
		panic(fmt.Sprintf("Block `%s` already has terminator", e.Block.Name()))
	}
	e.Term = terminator
}

func (block *ExtBlock) LLString() string {
	buf := &strings.Builder{}
	if block.IsUnnamed() {
		fmt.Fprintf(buf, "%s\n", enc.LabelID(block.LocalID))
	} else {
		fmt.Fprintf(buf, "%s\n", enc.LabelName(block.LocalName))
	}
	for _, inst := range block.Phis {
		fmt.Fprintf(buf, "\t%s\n", inst.LLString())
	}
	for _, inst := range block.Insts {
		fmt.Fprintf(buf, "\t%s\n", inst.LLString())
	}
	if !block.HasTerminator() {
		panic(fmt.Sprintf("missing terminator in basic block %q.\ncurrent instructions:\n%s", block.Name(), buf.String()))
	}
	fmt.Fprintf(buf, "\t%s", block.Term.LLString())
	return buf.String()
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

func (e *ExtBlock) NewPhi(incs ...*ir.Incoming) *ir.InstPhi {
	instPhi := ir.NewPhi(incs...)
	e.Phis = append(e.Phis, instPhi)
	return instPhi
}
