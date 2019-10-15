// Code generated from Kfn.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gen // Kfn
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 100,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 7, 2, 34, 10, 2,
	12, 2, 14, 2, 37, 11, 2, 3, 3, 3, 3, 6, 3, 41, 10, 3, 13, 3, 14, 3, 42,
	3, 4, 3, 4, 5, 4, 47, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 61, 10, 8, 3, 9, 3, 9, 5, 9, 65, 10,
	9, 3, 10, 3, 10, 5, 10, 69, 10, 10, 3, 10, 3, 10, 5, 10, 73, 10, 10, 3,
	10, 5, 10, 76, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	7, 13, 85, 10, 13, 12, 13, 14, 13, 88, 11, 13, 3, 14, 3, 14, 5, 14, 92,
	10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 2, 2, 17, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 5, 3, 2, 3, 4, 3,
	2, 6, 7, 4, 2, 5, 5, 8, 8, 2, 94, 2, 35, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2,
	6, 46, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 54, 3, 2,
	2, 2, 14, 56, 3, 2, 2, 2, 16, 64, 3, 2, 2, 2, 18, 75, 3, 2, 2, 2, 20, 77,
	3, 2, 2, 2, 22, 79, 3, 2, 2, 2, 24, 81, 3, 2, 2, 2, 26, 89, 3, 2, 2, 2,
	28, 95, 3, 2, 2, 2, 30, 97, 3, 2, 2, 2, 32, 34, 5, 4, 3, 2, 33, 32, 3,
	2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36,
	3, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 40, 5, 6, 4, 2, 39, 41, 9, 2, 2,
	2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 47, 5, 8, 5, 2, 45, 47, 5, 12, 7, 2,
	46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 7, 3, 2, 2, 2, 48, 49, 5, 10,
	6, 2, 49, 50, 7, 5, 2, 2, 50, 51, 5, 18, 10, 2, 51, 9, 3, 2, 2, 2, 52,
	53, 7, 10, 2, 2, 53, 11, 3, 2, 2, 2, 54, 55, 5, 14, 8, 2, 55, 13, 3, 2,
	2, 2, 56, 57, 5, 16, 9, 2, 57, 60, 7, 13, 2, 2, 58, 61, 5, 16, 9, 2, 59,
	61, 5, 14, 8, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 15, 3, 2,
	2, 2, 62, 65, 5, 18, 10, 2, 63, 65, 5, 10, 6, 2, 64, 62, 3, 2, 2, 2, 64,
	63, 3, 2, 2, 2, 65, 17, 3, 2, 2, 2, 66, 68, 5, 20, 11, 2, 67, 69, 5, 22,
	12, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70,
	71, 7, 9, 2, 2, 71, 73, 5, 24, 13, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2,
	2, 2, 73, 76, 3, 2, 2, 2, 74, 76, 5, 22, 12, 2, 75, 66, 3, 2, 2, 2, 75,
	74, 3, 2, 2, 2, 76, 19, 3, 2, 2, 2, 77, 78, 7, 11, 2, 2, 78, 21, 3, 2,
	2, 2, 79, 80, 7, 12, 2, 2, 80, 23, 3, 2, 2, 2, 81, 86, 5, 26, 14, 2, 82,
	83, 9, 3, 2, 2, 83, 85, 5, 26, 14, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2,
	2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 25, 3, 2, 2, 2, 88, 86,
	3, 2, 2, 2, 89, 91, 5, 28, 15, 2, 90, 92, 9, 4, 2, 2, 91, 90, 3, 2, 2,
	2, 91, 92, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 5, 30, 16, 2, 94, 27,
	3, 2, 2, 2, 95, 96, 7, 10, 2, 2, 96, 29, 3, 2, 2, 2, 97, 98, 7, 12, 2,
	2, 98, 31, 3, 2, 2, 2, 12, 35, 42, 46, 60, 64, 68, 72, 75, 86, 91,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'\n'", "':'", "','", "'and'", "'='", "'with'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "COMPONENT_OPTION_SEP", "UPPER_IDENT", "LOWER_IDENT",
	"STRING_LITERAL", "ARROW", "WS",
}

var ruleNames = []string{
	"kfn", "line", "stmt", "decl", "ident", "wire", "wireStmt", "wireElement",
	"component", "componentIdent", "componentValue", "componentOptionList",
	"componentOption", "componentOptionIdent", "componentOptionValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type KfnParser struct {
	*antlr.BaseParser
}

func NewKfnParser(input antlr.TokenStream) *KfnParser {
	this := new(KfnParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Kfn.g4"

	return this
}

// KfnParser tokens.
const (
	KfnParserEOF                  = antlr.TokenEOF
	KfnParserT__0                 = 1
	KfnParserT__1                 = 2
	KfnParserT__2                 = 3
	KfnParserT__3                 = 4
	KfnParserT__4                 = 5
	KfnParserT__5                 = 6
	KfnParserCOMPONENT_OPTION_SEP = 7
	KfnParserUPPER_IDENT          = 8
	KfnParserLOWER_IDENT          = 9
	KfnParserSTRING_LITERAL       = 10
	KfnParserARROW                = 11
	KfnParserWS                   = 12
)

// KfnParser rules.
const (
	KfnParserRULE_kfn                  = 0
	KfnParserRULE_line                 = 1
	KfnParserRULE_stmt                 = 2
	KfnParserRULE_decl                 = 3
	KfnParserRULE_ident                = 4
	KfnParserRULE_wire                 = 5
	KfnParserRULE_wireStmt             = 6
	KfnParserRULE_wireElement          = 7
	KfnParserRULE_component            = 8
	KfnParserRULE_componentIdent       = 9
	KfnParserRULE_componentValue       = 10
	KfnParserRULE_componentOptionList  = 11
	KfnParserRULE_componentOption      = 12
	KfnParserRULE_componentOptionIdent = 13
	KfnParserRULE_componentOptionValue = 14
)

// IKfnContext is an interface to support dynamic dispatch.
type IKfnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKfnContext differentiates from other interfaces.
	IsKfnContext()
}

type KfnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKfnContext() *KfnContext {
	var p = new(KfnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_kfn
	return p
}

func (*KfnContext) IsKfnContext() {}

func NewKfnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KfnContext {
	var p = new(KfnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_kfn

	return p
}

func (s *KfnContext) GetParser() antlr.Parser { return s.parser }

func (s *KfnContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *KfnContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *KfnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KfnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KfnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterKfn(s)
	}
}

func (s *KfnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitKfn(s)
	}
}

func (s *KfnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitKfn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Kfn() (localctx IKfnContext) {
	localctx = NewKfnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KfnParserRULE_kfn)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KfnParserUPPER_IDENT)|(1<<KfnParserLOWER_IDENT)|(1<<KfnParserSTRING_LITERAL))) != 0 {
		{
			p.SetState(30)
			p.Line()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Stmt() IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitLine(s)
	}
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KfnParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Stmt()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KfnParserT__0 || _la == KfnParserT__1 {
		{
			p.SetState(37)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KfnParserT__0 || _la == KfnParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl() IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *StmtContext) Wire() IWireContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWireContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWireContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KfnParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Decl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Wire()
		}

	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *DeclContext) Component() IComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KfnParserRULE_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Ident()
	}
	{
		p.SetState(47)
		p.Match(KfnParserT__2)
	}
	{
		p.SetState(48)
		p.Component()
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) UPPER_IDENT() antlr.TerminalNode {
	return s.GetToken(KfnParserUPPER_IDENT, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KfnParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(KfnParserUPPER_IDENT)
	}

	return localctx
}

// IWireContext is an interface to support dynamic dispatch.
type IWireContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWireContext differentiates from other interfaces.
	IsWireContext()
}

type WireContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWireContext() *WireContext {
	var p = new(WireContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_wire
	return p
}

func (*WireContext) IsWireContext() {}

func NewWireContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WireContext {
	var p = new(WireContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_wire

	return p
}

func (s *WireContext) GetParser() antlr.Parser { return s.parser }

func (s *WireContext) WireStmt() IWireStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWireStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWireStmtContext)
}

func (s *WireContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WireContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WireContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterWire(s)
	}
}

func (s *WireContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitWire(s)
	}
}

func (s *WireContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitWire(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Wire() (localctx IWireContext) {
	localctx = NewWireContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KfnParserRULE_wire)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.WireStmt()
	}

	return localctx
}

// IWireStmtContext is an interface to support dynamic dispatch.
type IWireStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWireStmtContext differentiates from other interfaces.
	IsWireStmtContext()
}

type WireStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWireStmtContext() *WireStmtContext {
	var p = new(WireStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_wireStmt
	return p
}

func (*WireStmtContext) IsWireStmtContext() {}

func NewWireStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WireStmtContext {
	var p = new(WireStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_wireStmt

	return p
}

func (s *WireStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WireStmtContext) AllWireElement() []IWireElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWireElementContext)(nil)).Elem())
	var tst = make([]IWireElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWireElementContext)
		}
	}

	return tst
}

func (s *WireStmtContext) WireElement(i int) IWireElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWireElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWireElementContext)
}

func (s *WireStmtContext) ARROW() antlr.TerminalNode {
	return s.GetToken(KfnParserARROW, 0)
}

func (s *WireStmtContext) WireStmt() IWireStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWireStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWireStmtContext)
}

func (s *WireStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WireStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WireStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterWireStmt(s)
	}
}

func (s *WireStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitWireStmt(s)
	}
}

func (s *WireStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitWireStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) WireStmt() (localctx IWireStmtContext) {
	localctx = NewWireStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KfnParserRULE_wireStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.WireElement()
	}
	{
		p.SetState(55)
		p.Match(KfnParserARROW)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(56)
			p.WireElement()
		}

	case 2:
		{
			p.SetState(57)
			p.WireStmt()
		}

	}

	return localctx
}

// IWireElementContext is an interface to support dynamic dispatch.
type IWireElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWireElementContext differentiates from other interfaces.
	IsWireElementContext()
}

type WireElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWireElementContext() *WireElementContext {
	var p = new(WireElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_wireElement
	return p
}

func (*WireElementContext) IsWireElementContext() {}

func NewWireElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WireElementContext {
	var p = new(WireElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_wireElement

	return p
}

func (s *WireElementContext) GetParser() antlr.Parser { return s.parser }

func (s *WireElementContext) Component() IComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *WireElementContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *WireElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WireElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WireElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterWireElement(s)
	}
}

func (s *WireElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitWireElement(s)
	}
}

func (s *WireElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitWireElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) WireElement() (localctx IWireElementContext) {
	localctx = NewWireElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KfnParserRULE_wireElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KfnParserLOWER_IDENT, KfnParserSTRING_LITERAL:
		{
			p.SetState(60)
			p.Component()
		}

	case KfnParserUPPER_IDENT:
		{
			p.SetState(61)
			p.Ident()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_component
	return p
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) ComponentIdent() IComponentIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentIdentContext)
}

func (s *ComponentContext) ComponentValue() IComponentValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentValueContext)
}

func (s *ComponentContext) COMPONENT_OPTION_SEP() antlr.TerminalNode {
	return s.GetToken(KfnParserCOMPONENT_OPTION_SEP, 0)
}

func (s *ComponentContext) ComponentOptionList() IComponentOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentOptionListContext)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponent(s)
	}
}

func (s *ComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponent(s)
	}
}

func (s *ComponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) Component() (localctx IComponentContext) {
	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KfnParserRULE_component)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KfnParserLOWER_IDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.ComponentIdent()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == KfnParserSTRING_LITERAL {
			{
				p.SetState(65)
				p.ComponentValue()
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == KfnParserCOMPONENT_OPTION_SEP {
			{
				p.SetState(68)
				p.Match(KfnParserCOMPONENT_OPTION_SEP)
			}
			{
				p.SetState(69)
				p.ComponentOptionList()
			}

		}

	case KfnParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.ComponentValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComponentIdentContext is an interface to support dynamic dispatch.
type IComponentIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentIdentContext differentiates from other interfaces.
	IsComponentIdentContext()
}

type ComponentIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentIdentContext() *ComponentIdentContext {
	var p = new(ComponentIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentIdent
	return p
}

func (*ComponentIdentContext) IsComponentIdentContext() {}

func NewComponentIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentIdentContext {
	var p = new(ComponentIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentIdent

	return p
}

func (s *ComponentIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentIdentContext) LOWER_IDENT() antlr.TerminalNode {
	return s.GetToken(KfnParserLOWER_IDENT, 0)
}

func (s *ComponentIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentIdent(s)
	}
}

func (s *ComponentIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentIdent(s)
	}
}

func (s *ComponentIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentIdent() (localctx IComponentIdentContext) {
	localctx = NewComponentIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KfnParserRULE_componentIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(KfnParserLOWER_IDENT)
	}

	return localctx
}

// IComponentValueContext is an interface to support dynamic dispatch.
type IComponentValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentValueContext differentiates from other interfaces.
	IsComponentValueContext()
}

type ComponentValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentValueContext() *ComponentValueContext {
	var p = new(ComponentValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentValue
	return p
}

func (*ComponentValueContext) IsComponentValueContext() {}

func NewComponentValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentValueContext {
	var p = new(ComponentValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentValue

	return p
}

func (s *ComponentValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(KfnParserSTRING_LITERAL, 0)
}

func (s *ComponentValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentValue(s)
	}
}

func (s *ComponentValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentValue(s)
	}
}

func (s *ComponentValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentValue() (localctx IComponentValueContext) {
	localctx = NewComponentValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KfnParserRULE_componentValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(KfnParserSTRING_LITERAL)
	}

	return localctx
}

// IComponentOptionListContext is an interface to support dynamic dispatch.
type IComponentOptionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentOptionListContext differentiates from other interfaces.
	IsComponentOptionListContext()
}

type ComponentOptionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentOptionListContext() *ComponentOptionListContext {
	var p = new(ComponentOptionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentOptionList
	return p
}

func (*ComponentOptionListContext) IsComponentOptionListContext() {}

func NewComponentOptionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentOptionListContext {
	var p = new(ComponentOptionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentOptionList

	return p
}

func (s *ComponentOptionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentOptionListContext) AllComponentOption() []IComponentOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentOptionContext)(nil)).Elem())
	var tst = make([]IComponentOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentOptionContext)
		}
	}

	return tst
}

func (s *ComponentOptionListContext) ComponentOption(i int) IComponentOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentOptionContext)
}

func (s *ComponentOptionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentOptionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentOptionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentOptionList(s)
	}
}

func (s *ComponentOptionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentOptionList(s)
	}
}

func (s *ComponentOptionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentOptionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentOptionList() (localctx IComponentOptionListContext) {
	localctx = NewComponentOptionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KfnParserRULE_componentOptionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.ComponentOption()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KfnParserT__3 || _la == KfnParserT__4 {
		{
			p.SetState(80)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KfnParserT__3 || _la == KfnParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(81)
			p.ComponentOption()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComponentOptionContext is an interface to support dynamic dispatch.
type IComponentOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentOptionContext differentiates from other interfaces.
	IsComponentOptionContext()
}

type ComponentOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentOptionContext() *ComponentOptionContext {
	var p = new(ComponentOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentOption
	return p
}

func (*ComponentOptionContext) IsComponentOptionContext() {}

func NewComponentOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentOptionContext {
	var p = new(ComponentOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentOption

	return p
}

func (s *ComponentOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentOptionContext) ComponentOptionIdent() IComponentOptionIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentOptionIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentOptionIdentContext)
}

func (s *ComponentOptionContext) ComponentOptionValue() IComponentOptionValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentOptionValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentOptionValueContext)
}

func (s *ComponentOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentOption(s)
	}
}

func (s *ComponentOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentOption(s)
	}
}

func (s *ComponentOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentOption() (localctx IComponentOptionContext) {
	localctx = NewComponentOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KfnParserRULE_componentOption)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.ComponentOptionIdent()
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KfnParserT__2 || _la == KfnParserT__5 {
		{
			p.SetState(88)
			_la = p.GetTokenStream().LA(1)

			if !(_la == KfnParserT__2 || _la == KfnParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(91)
		p.ComponentOptionValue()
	}

	return localctx
}

// IComponentOptionIdentContext is an interface to support dynamic dispatch.
type IComponentOptionIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentOptionIdentContext differentiates from other interfaces.
	IsComponentOptionIdentContext()
}

type ComponentOptionIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentOptionIdentContext() *ComponentOptionIdentContext {
	var p = new(ComponentOptionIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentOptionIdent
	return p
}

func (*ComponentOptionIdentContext) IsComponentOptionIdentContext() {}

func NewComponentOptionIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentOptionIdentContext {
	var p = new(ComponentOptionIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentOptionIdent

	return p
}

func (s *ComponentOptionIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentOptionIdentContext) UPPER_IDENT() antlr.TerminalNode {
	return s.GetToken(KfnParserUPPER_IDENT, 0)
}

func (s *ComponentOptionIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentOptionIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentOptionIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentOptionIdent(s)
	}
}

func (s *ComponentOptionIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentOptionIdent(s)
	}
}

func (s *ComponentOptionIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentOptionIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentOptionIdent() (localctx IComponentOptionIdentContext) {
	localctx = NewComponentOptionIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KfnParserRULE_componentOptionIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(KfnParserUPPER_IDENT)
	}

	return localctx
}

// IComponentOptionValueContext is an interface to support dynamic dispatch.
type IComponentOptionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentOptionValueContext differentiates from other interfaces.
	IsComponentOptionValueContext()
}

type ComponentOptionValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentOptionValueContext() *ComponentOptionValueContext {
	var p = new(ComponentOptionValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KfnParserRULE_componentOptionValue
	return p
}

func (*ComponentOptionValueContext) IsComponentOptionValueContext() {}

func NewComponentOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentOptionValueContext {
	var p = new(ComponentOptionValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KfnParserRULE_componentOptionValue

	return p
}

func (s *ComponentOptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentOptionValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(KfnParserSTRING_LITERAL, 0)
}

func (s *ComponentOptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentOptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentOptionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.EnterComponentOptionValue(s)
	}
}

func (s *ComponentOptionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KfnListener); ok {
		listenerT.ExitComponentOptionValue(s)
	}
}

func (s *ComponentOptionValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KfnVisitor:
		return t.VisitComponentOptionValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KfnParser) ComponentOptionValue() (localctx IComponentOptionValueContext) {
	localctx = NewComponentOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KfnParserRULE_componentOptionValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(KfnParserSTRING_LITERAL)
	}

	return localctx
}
