// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Kfn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// KfnListener is a complete listener for a parse tree produced by KfnParser.
type KfnListener interface {
	antlr.ParseTreeListener

	// EnterKfn is called when entering the kfn production.
	EnterKfn(c *KfnContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterWire is called when entering the wire production.
	EnterWire(c *WireContext)

	// EnterWireStmt is called when entering the wireStmt production.
	EnterWireStmt(c *WireStmtContext)

	// EnterWireElement is called when entering the wireElement production.
	EnterWireElement(c *WireElementContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterComponentIdent is called when entering the componentIdent production.
	EnterComponentIdent(c *ComponentIdentContext)

	// EnterComponentValue is called when entering the componentValue production.
	EnterComponentValue(c *ComponentValueContext)

	// EnterComponentOptionList is called when entering the componentOptionList production.
	EnterComponentOptionList(c *ComponentOptionListContext)

	// EnterComponentOption is called when entering the componentOption production.
	EnterComponentOption(c *ComponentOptionContext)

	// EnterComponentOptionIdent is called when entering the componentOptionIdent production.
	EnterComponentOptionIdent(c *ComponentOptionIdentContext)

	// EnterComponentOptionValue is called when entering the componentOptionValue production.
	EnterComponentOptionValue(c *ComponentOptionValueContext)

	// ExitKfn is called when exiting the kfn production.
	ExitKfn(c *KfnContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitWire is called when exiting the wire production.
	ExitWire(c *WireContext)

	// ExitWireStmt is called when exiting the wireStmt production.
	ExitWireStmt(c *WireStmtContext)

	// ExitWireElement is called when exiting the wireElement production.
	ExitWireElement(c *WireElementContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitComponentIdent is called when exiting the componentIdent production.
	ExitComponentIdent(c *ComponentIdentContext)

	// ExitComponentValue is called when exiting the componentValue production.
	ExitComponentValue(c *ComponentValueContext)

	// ExitComponentOptionList is called when exiting the componentOptionList production.
	ExitComponentOptionList(c *ComponentOptionListContext)

	// ExitComponentOption is called when exiting the componentOption production.
	ExitComponentOption(c *ComponentOptionContext)

	// ExitComponentOptionIdent is called when exiting the componentOptionIdent production.
	ExitComponentOptionIdent(c *ComponentOptionIdentContext)

	// ExitComponentOptionValue is called when exiting the componentOptionValue production.
	ExitComponentOptionValue(c *ComponentOptionValueContext)
}
