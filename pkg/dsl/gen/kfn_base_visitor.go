// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Kfn
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseKfnVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKfnVisitor) VisitKfn(ctx *KfnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitWire(ctx *WireContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitWireStmt(ctx *WireStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitWireElement(ctx *WireElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponent(ctx *ComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentIdent(ctx *ComponentIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentValue(ctx *ComponentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentOptionList(ctx *ComponentOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentOption(ctx *ComponentOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentOptionIdent(ctx *ComponentOptionIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKfnVisitor) VisitComponentOptionValue(ctx *ComponentOptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}
