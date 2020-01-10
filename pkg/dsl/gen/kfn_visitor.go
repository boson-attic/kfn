// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Kfn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by KfnParser.
type KfnVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KfnParser#kfn.
	VisitKfn(ctx *KfnContext) interface{}

	// Visit a parse tree produced by KfnParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by KfnParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by KfnParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by KfnParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by KfnParser#wire.
	VisitWire(ctx *WireContext) interface{}

	// Visit a parse tree produced by KfnParser#wireStmt.
	VisitWireStmt(ctx *WireStmtContext) interface{}

	// Visit a parse tree produced by KfnParser#wireElement.
	VisitWireElement(ctx *WireElementContext) interface{}

	// Visit a parse tree produced by KfnParser#component.
	VisitComponent(ctx *ComponentContext) interface{}

	// Visit a parse tree produced by KfnParser#componentIdent.
	VisitComponentIdent(ctx *ComponentIdentContext) interface{}

	// Visit a parse tree produced by KfnParser#componentValue.
	VisitComponentValue(ctx *ComponentValueContext) interface{}

	// Visit a parse tree produced by KfnParser#componentOptionList.
	VisitComponentOptionList(ctx *ComponentOptionListContext) interface{}

	// Visit a parse tree produced by KfnParser#componentOption.
	VisitComponentOption(ctx *ComponentOptionContext) interface{}

	// Visit a parse tree produced by KfnParser#componentOptionIdent.
	VisitComponentOptionIdent(ctx *ComponentOptionIdentContext) interface{}

	// Visit a parse tree produced by KfnParser#componentOptionValue.
	VisitComponentOptionValue(ctx *ComponentOptionValueContext) interface{}
}
