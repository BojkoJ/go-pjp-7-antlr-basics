// Code generated from PLC_Lab7_expr_2.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser2 // PLC_Lab7_expr_2
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 40, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 6,
	3, 15, 10, 3, 13, 3, 14, 3, 16, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 27, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 35,
	10, 4, 12, 4, 14, 4, 38, 11, 4, 3, 4, 2, 3, 6, 5, 2, 4, 6, 2, 4, 3, 2,
	4, 5, 3, 2, 6, 7, 2, 42, 2, 8, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 26, 3,
	2, 2, 2, 8, 9, 5, 4, 3, 2, 9, 10, 7, 2, 2, 3, 10, 3, 3, 2, 2, 2, 11, 12,
	5, 6, 4, 2, 12, 13, 7, 3, 2, 2, 13, 15, 3, 2, 2, 2, 14, 11, 3, 2, 2, 2,
	15, 16, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 5, 3, 2,
	2, 2, 18, 19, 8, 4, 1, 2, 19, 27, 7, 10, 2, 2, 20, 27, 7, 12, 2, 2, 21,
	27, 7, 11, 2, 2, 22, 23, 7, 8, 2, 2, 23, 24, 5, 6, 4, 2, 24, 25, 7, 9,
	2, 2, 25, 27, 3, 2, 2, 2, 26, 18, 3, 2, 2, 2, 26, 20, 3, 2, 2, 2, 26, 21,
	3, 2, 2, 2, 26, 22, 3, 2, 2, 2, 27, 36, 3, 2, 2, 2, 28, 29, 12, 8, 2, 2,
	29, 30, 9, 2, 2, 2, 30, 35, 5, 6, 4, 9, 31, 32, 12, 7, 2, 2, 32, 33, 9,
	3, 2, 2, 33, 35, 5, 6, 4, 8, 34, 28, 3, 2, 2, 2, 34, 31, 3, 2, 2, 2, 35,
	38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 7, 3, 2, 2,
	2, 38, 36, 3, 2, 2, 2, 6, 16, 26, 34, 36,
}
var literalNames = []string{
	"", "';'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "INT", "OCT", "HEX", "WS",
}

var ruleNames = []string{
	"prog", "expr_list", "expr",
}

type PLC_Lab7_expr_2Parser struct {
	*antlr.BaseParser
}

// NewPLC_Lab7_expr_2Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PLC_Lab7_expr_2Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPLC_Lab7_expr_2Parser(input antlr.TokenStream) *PLC_Lab7_expr_2Parser {
	this := new(PLC_Lab7_expr_2Parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PLC_Lab7_expr_2.g4"

	return this
}

// PLC_Lab7_expr_2Parser tokens.
const (
	PLC_Lab7_expr_2ParserEOF  = antlr.TokenEOF
	PLC_Lab7_expr_2ParserT__0 = 1
	PLC_Lab7_expr_2ParserT__1 = 2
	PLC_Lab7_expr_2ParserT__2 = 3
	PLC_Lab7_expr_2ParserT__3 = 4
	PLC_Lab7_expr_2ParserT__4 = 5
	PLC_Lab7_expr_2ParserT__5 = 6
	PLC_Lab7_expr_2ParserT__6 = 7
	PLC_Lab7_expr_2ParserINT  = 8
	PLC_Lab7_expr_2ParserOCT  = 9
	PLC_Lab7_expr_2ParserHEX  = 10
	PLC_Lab7_expr_2ParserWS   = 11
)

// PLC_Lab7_expr_2Parser rules.
const (
	PLC_Lab7_expr_2ParserRULE_prog      = 0
	PLC_Lab7_expr_2ParserRULE_expr_list = 1
	PLC_Lab7_expr_2ParserRULE_expr      = 2
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Expr_list() IExpr_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_listContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(PLC_Lab7_expr_2ParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *PLC_Lab7_expr_2Parser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PLC_Lab7_expr_2ParserRULE_prog)

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
		p.SetState(6)
		p.Expr_list()
	}
	{
		p.SetState(7)
		p.Match(PLC_Lab7_expr_2ParserEOF)
	}

	return localctx
}

// IExpr_listContext is an interface to support dynamic dispatch.
type IExpr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpr_listContext differentiates from other interfaces.
	IsExpr_listContext()
}

type Expr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_listContext() *Expr_listContext {
	var p = new(Expr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_expr_list
	return p
}

func (*Expr_listContext) IsExpr_listContext() {}

func NewExpr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_listContext {
	var p = new(Expr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_expr_list

	return p
}

func (s *Expr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_listContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Expr_listContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterExpr_list(s)
	}
}

func (s *Expr_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitExpr_list(s)
	}
}

func (p *PLC_Lab7_expr_2Parser) Expr_list() (localctx IExpr_listContext) {
	localctx = NewExpr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PLC_Lab7_expr_2ParserRULE_expr_list)
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
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PLC_Lab7_expr_2ParserT__5)|(1<<PLC_Lab7_expr_2ParserINT)|(1<<PLC_Lab7_expr_2ParserOCT)|(1<<PLC_Lab7_expr_2ParserHEX))) != 0) {
		{
			p.SetState(9)
			p.expr(0)
		}
		{
			p.SetState(10)
			p.Match(PLC_Lab7_expr_2ParserT__0)
		}

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PLC_Lab7_expr_2ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OctContext struct {
	*ExprContext
}

func NewOctContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctContext {
	var p = new(OctContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OctContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctContext) OCT() antlr.TerminalNode {
	return s.GetToken(PLC_Lab7_expr_2ParserOCT, 0)
}

func (s *OctContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterOct(s)
	}
}

func (s *OctContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitOct(s)
	}
}

type MulDivContext struct {
	*ExprContext
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*ExprContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitAddSub(s)
	}
}

type ParensContext struct {
	*ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitParens(s)
	}
}

type HexContext struct {
	*ExprContext
}

func NewHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexContext {
	var p = new(HexContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *HexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexContext) HEX() antlr.TerminalNode {
	return s.GetToken(PLC_Lab7_expr_2ParserHEX, 0)
}

func (s *HexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterHex(s)
	}
}

func (s *HexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitHex(s)
	}
}

type IntContext struct {
	*ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(PLC_Lab7_expr_2ParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PLC_Lab7_expr_2Listener); ok {
		listenerT.ExitInt(s)
	}
}

func (p *PLC_Lab7_expr_2Parser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PLC_Lab7_expr_2Parser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, PLC_Lab7_expr_2ParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PLC_Lab7_expr_2ParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(17)
			p.Match(PLC_Lab7_expr_2ParserINT)
		}

	case PLC_Lab7_expr_2ParserHEX:
		localctx = NewHexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(PLC_Lab7_expr_2ParserHEX)
		}

	case PLC_Lab7_expr_2ParserOCT:
		localctx = NewOctContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Match(PLC_Lab7_expr_2ParserOCT)
		}

	case PLC_Lab7_expr_2ParserT__5:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Match(PLC_Lab7_expr_2ParserT__5)
		}
		{
			p.SetState(21)
			p.expr(0)
		}
		{
			p.SetState(22)
			p.Match(PLC_Lab7_expr_2ParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLC_Lab7_expr_2ParserRULE_expr)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(27)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PLC_Lab7_expr_2ParserT__1 || _la == PLC_Lab7_expr_2ParserT__2) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expr(7)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PLC_Lab7_expr_2ParserRULE_expr)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(30)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PLC_Lab7_expr_2ParserT__3 || _la == PLC_Lab7_expr_2ParserT__4) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.expr(6)
				}

			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

func (p *PLC_Lab7_expr_2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PLC_Lab7_expr_2Parser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
