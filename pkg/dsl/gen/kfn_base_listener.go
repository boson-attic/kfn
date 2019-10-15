// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Kfn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKfnListener is a complete listener for a parse tree produced by KfnParser.
type BaseKfnListener struct{}

var _ KfnListener = &BaseKfnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKfnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKfnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKfnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKfnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKfn is called when production kfn is entered.
func (s *BaseKfnListener) EnterKfn(ctx *KfnContext) {}

// ExitKfn is called when production kfn is exited.
func (s *BaseKfnListener) ExitKfn(ctx *KfnContext) {}

// EnterLine is called when production line is entered.
func (s *BaseKfnListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseKfnListener) ExitLine(ctx *LineContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseKfnListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseKfnListener) ExitStmt(ctx *StmtContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseKfnListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseKfnListener) ExitDecl(ctx *DeclContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseKfnListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseKfnListener) ExitIdent(ctx *IdentContext) {}

// EnterWire is called when production wire is entered.
func (s *BaseKfnListener) EnterWire(ctx *WireContext) {}

// ExitWire is called when production wire is exited.
func (s *BaseKfnListener) ExitWire(ctx *WireContext) {}

// EnterWireStmt is called when production wireStmt is entered.
func (s *BaseKfnListener) EnterWireStmt(ctx *WireStmtContext) {}

// ExitWireStmt is called when production wireStmt is exited.
func (s *BaseKfnListener) ExitWireStmt(ctx *WireStmtContext) {}

// EnterWireElement is called when production wireElement is entered.
func (s *BaseKfnListener) EnterWireElement(ctx *WireElementContext) {}

// ExitWireElement is called when production wireElement is exited.
func (s *BaseKfnListener) ExitWireElement(ctx *WireElementContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseKfnListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseKfnListener) ExitComponent(ctx *ComponentContext) {}

// EnterComponentIdent is called when production componentIdent is entered.
func (s *BaseKfnListener) EnterComponentIdent(ctx *ComponentIdentContext) {}

// ExitComponentIdent is called when production componentIdent is exited.
func (s *BaseKfnListener) ExitComponentIdent(ctx *ComponentIdentContext) {}

// EnterComponentValue is called when production componentValue is entered.
func (s *BaseKfnListener) EnterComponentValue(ctx *ComponentValueContext) {}

// ExitComponentValue is called when production componentValue is exited.
func (s *BaseKfnListener) ExitComponentValue(ctx *ComponentValueContext) {}

// EnterComponentOptionList is called when production componentOptionList is entered.
func (s *BaseKfnListener) EnterComponentOptionList(ctx *ComponentOptionListContext) {}

// ExitComponentOptionList is called when production componentOptionList is exited.
func (s *BaseKfnListener) ExitComponentOptionList(ctx *ComponentOptionListContext) {}

// EnterComponentOption is called when production componentOption is entered.
func (s *BaseKfnListener) EnterComponentOption(ctx *ComponentOptionContext) {}

// ExitComponentOption is called when production componentOption is exited.
func (s *BaseKfnListener) ExitComponentOption(ctx *ComponentOptionContext) {}

// EnterComponentOptionIdent is called when production componentOptionIdent is entered.
func (s *BaseKfnListener) EnterComponentOptionIdent(ctx *ComponentOptionIdentContext) {}

// ExitComponentOptionIdent is called when production componentOptionIdent is exited.
func (s *BaseKfnListener) ExitComponentOptionIdent(ctx *ComponentOptionIdentContext) {}

// EnterComponentOptionValue is called when production componentOptionValue is entered.
func (s *BaseKfnListener) EnterComponentOptionValue(ctx *ComponentOptionValueContext) {}

// ExitComponentOptionValue is called when production componentOptionValue is exited.
func (s *BaseKfnListener) ExitComponentOptionValue(ctx *ComponentOptionValueContext) {}
