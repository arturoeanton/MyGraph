// Code generated from /Users/arturoeliasanton/go/src/r2Server/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query // R2Query
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 80, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3, 
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24, 11, 
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 35, 10, 
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 45, 10, 6, 12, 
	6, 14, 6, 48, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 78, 10, 7, 3, 7, 2, 
	2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 84, 2, 14, 3, 2, 2, 2, 4, 17, 3, 2, 
	2, 2, 6, 34, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12, 77, 
	3, 2, 2, 2, 14, 15, 5, 12, 7, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2, 2, 
	17, 22, 7, 17, 2, 2, 18, 19, 7, 10, 2, 2, 19, 21, 7, 17, 2, 2, 20, 18, 
	3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 
	23, 5, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 12, 2, 2, 26, 27, 7, 
	17, 2, 2, 27, 35, 7, 13, 2, 2, 28, 29, 7, 14, 2, 2, 29, 30, 7, 17, 2, 2, 
	30, 35, 7, 12, 2, 2, 31, 32, 7, 14, 2, 2, 32, 33, 7, 17, 2, 2, 33, 35, 
	7, 13, 2, 2, 34, 25, 3, 2, 2, 2, 34, 28, 3, 2, 2, 2, 34, 31, 3, 2, 2, 2, 
	35, 7, 3, 2, 2, 2, 36, 37, 7, 3, 2, 2, 37, 38, 5, 6, 4, 2, 38, 39, 5, 4, 
	3, 2, 39, 40, 7, 4, 2, 2, 40, 9, 3, 2, 2, 2, 41, 46, 5, 8, 5, 2, 42, 43, 
	7, 10, 2, 2, 43, 45, 5, 8, 5, 2, 44, 42, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 
	46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 11, 3, 2, 2, 2, 48, 46, 3, 
	2, 2, 2, 49, 50, 7, 5, 2, 2, 50, 78, 7, 15, 2, 2, 51, 52, 7, 6, 2, 2, 52, 
	78, 7, 15, 2, 2, 53, 54, 7, 15, 2, 2, 54, 55, 7, 11, 2, 2, 55, 78, 5, 4, 
	3, 2, 56, 57, 7, 15, 2, 2, 57, 58, 7, 11, 2, 2, 58, 78, 5, 10, 6, 2, 59, 
	60, 7, 15, 2, 2, 60, 61, 7, 7, 2, 2, 61, 78, 5, 4, 3, 2, 62, 63, 7, 15, 
	2, 2, 63, 64, 7, 6, 2, 2, 64, 78, 5, 4, 3, 2, 65, 66, 7, 15, 2, 2, 66, 
	67, 7, 7, 2, 2, 67, 68, 5, 4, 3, 2, 68, 69, 5, 6, 4, 2, 69, 70, 5, 4, 3, 
	2, 70, 78, 3, 2, 2, 2, 71, 72, 7, 15, 2, 2, 72, 73, 7, 6, 2, 2, 73, 74, 
	5, 4, 3, 2, 74, 75, 5, 6, 4, 2, 75, 76, 5, 4, 3, 2, 76, 78, 3, 2, 2, 2, 
	77, 49, 3, 2, 2, 2, 77, 51, 3, 2, 2, 2, 77, 53, 3, 2, 2, 2, 77, 56, 3, 
	2, 2, 2, 77, 59, 3, 2, 2, 2, 77, 62, 3, 2, 2, 2, 77, 65, 3, 2, 2, 2, 77, 
	71, 3, 2, 2, 2, 78, 13, 3, 2, 2, 2, 6, 22, 34, 46, 77,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'CREATE'", "'DELETE'", "'ADD'", "'OR'", "'AND'", "','", 
	"'GET'", "'-'", "'->'", "'<-'",
}
var symbolicNames = []string{
	"", "", "", "CREATE", "DELETE", "ADD", "OR", "AND", "COMMA", "GET", "MIDDLE", 
	"DIRECTION_REL_R", "DIRECTION_REL_L", "GRAPH_NAME", "WHITESPACE", "STRING",
}

var ruleNames = []string{
	"start", "arguments", "rel", "relation", "relations", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type R2QueryParser struct {
	*antlr.BaseParser
}

func NewR2QueryParser(input antlr.TokenStream) *R2QueryParser {
	this := new(R2QueryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "R2Query.g4"

	return this
}

// R2QueryParser tokens.
const (
	R2QueryParserEOF = antlr.TokenEOF
	R2QueryParserT__0 = 1
	R2QueryParserT__1 = 2
	R2QueryParserCREATE = 3
	R2QueryParserDELETE = 4
	R2QueryParserADD = 5
	R2QueryParserOR = 6
	R2QueryParserAND = 7
	R2QueryParserCOMMA = 8
	R2QueryParserGET = 9
	R2QueryParserMIDDLE = 10
	R2QueryParserDIRECTION_REL_R = 11
	R2QueryParserDIRECTION_REL_L = 12
	R2QueryParserGRAPH_NAME = 13
	R2QueryParserWHITESPACE = 14
	R2QueryParserSTRING = 15
)

// R2QueryParser rules.
const (
	R2QueryParserRULE_start = 0
	R2QueryParserRULE_arguments = 1
	R2QueryParserRULE_rel = 2
	R2QueryParserRULE_relation = 3
	R2QueryParserRULE_relations = 4
	R2QueryParserRULE_expression = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(R2QueryParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *R2QueryParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, R2QueryParserRULE_start)

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
		p.SetState(12)
		p.Expression()
	}
	{
		p.SetState(13)
		p.Match(R2QueryParserEOF)
	}



	return localctx
}


// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(R2QueryParserSTRING)
}

func (s *ArgumentsContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(R2QueryParserSTRING, i)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2QueryParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2QueryParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *R2QueryParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, R2QueryParserRULE_arguments)
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
		p.SetState(15)
		p.Match(R2QueryParserSTRING)
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == R2QueryParserCOMMA {
		{
			p.SetState(16)
			p.Match(R2QueryParserCOMMA)
		}
		{
			p.SetState(17)
			p.Match(R2QueryParserSTRING)
		}


		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IRelContext is an interface to support dynamic dispatch.
type IRelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp2 returns the op2 token.
	GetOp2() antlr.Token 

	// GetTag returns the tag token.
	GetTag() antlr.Token 

	// GetOp1 returns the op1 token.
	GetOp1() antlr.Token 


	// SetOp2 sets the op2 token.
	SetOp2(antlr.Token) 

	// SetTag sets the tag token.
	SetTag(antlr.Token) 

	// SetOp1 sets the op1 token.
	SetOp1(antlr.Token) 


	// IsRelContext differentiates from other interfaces.
	IsRelContext()
}

type RelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op2 antlr.Token
	tag antlr.Token
	op1 antlr.Token
}

func NewEmptyRelContext() *RelContext {
	var p = new(RelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_rel
	return p
}

func (*RelContext) IsRelContext() {}

func NewRelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelContext {
	var p = new(RelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_rel

	return p
}

func (s *RelContext) GetParser() antlr.Parser { return s.parser }

func (s *RelContext) GetOp2() antlr.Token { return s.op2 }

func (s *RelContext) GetTag() antlr.Token { return s.tag }

func (s *RelContext) GetOp1() antlr.Token { return s.op1 }


func (s *RelContext) SetOp2(v antlr.Token) { s.op2 = v }

func (s *RelContext) SetTag(v antlr.Token) { s.tag = v }

func (s *RelContext) SetOp1(v antlr.Token) { s.op1 = v }


func (s *RelContext) MIDDLE() antlr.TerminalNode {
	return s.GetToken(R2QueryParserMIDDLE, 0)
}

func (s *RelContext) STRING() antlr.TerminalNode {
	return s.GetToken(R2QueryParserSTRING, 0)
}

func (s *RelContext) DIRECTION_REL_R() antlr.TerminalNode {
	return s.GetToken(R2QueryParserDIRECTION_REL_R, 0)
}

func (s *RelContext) DIRECTION_REL_L() antlr.TerminalNode {
	return s.GetToken(R2QueryParserDIRECTION_REL_L, 0)
}

func (s *RelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterRel(s)
	}
}

func (s *RelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitRel(s)
	}
}

func (s *RelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitRel(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *R2QueryParser) Rel() (localctx IRelContext) {
	localctx = NewRelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, R2QueryParserRULE_rel)

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

	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)

			var _m = p.Match(R2QueryParserMIDDLE)

			localctx.(*RelContext).op2 = _m
		}
		{
			p.SetState(24)

			var _m = p.Match(R2QueryParserSTRING)

			localctx.(*RelContext).tag = _m
		}
		{
			p.SetState(25)

			var _m = p.Match(R2QueryParserDIRECTION_REL_R)

			localctx.(*RelContext).op1 = _m
		}



	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)

			var _m = p.Match(R2QueryParserDIRECTION_REL_L)

			localctx.(*RelContext).op1 = _m
		}
		{
			p.SetState(27)

			var _m = p.Match(R2QueryParserSTRING)

			localctx.(*RelContext).tag = _m
		}
		{
			p.SetState(28)

			var _m = p.Match(R2QueryParserMIDDLE)

			localctx.(*RelContext).op2 = _m
		}



	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)

			var _m = p.Match(R2QueryParserDIRECTION_REL_L)

			localctx.(*RelContext).op1 = _m
		}
		{
			p.SetState(30)

			var _m = p.Match(R2QueryParserSTRING)

			localctx.(*RelContext).tag = _m
		}
		{
			p.SetState(31)

			var _m = p.Match(R2QueryParserDIRECTION_REL_R)

			localctx.(*RelContext).op2 = _m
		}


	}


	return localctx
}


// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r rule contexts.
	GetR() IRelContext

	// GetArg returns the arg rule contexts.
	GetArg() IArgumentsContext


	// SetR sets the r rule contexts.
	SetR(IRelContext)

	// SetArg sets the arg rule contexts.
	SetArg(IArgumentsContext)


	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	r IRelContext 
	arg IArgumentsContext 
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) GetR() IRelContext { return s.r }

func (s *RelationContext) GetArg() IArgumentsContext { return s.arg }


func (s *RelationContext) SetR(v IRelContext) { s.r = v }

func (s *RelationContext) SetArg(v IArgumentsContext) { s.arg = v }


func (s *RelationContext) Rel() IRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelContext)
}

func (s *RelationContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (s *RelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitRelation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *R2QueryParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, R2QueryParserRULE_relation)

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
		p.SetState(34)
		p.Match(R2QueryParserT__0)
	}
	{
		p.SetState(35)

		var _x = p.Rel()


		localctx.(*RelationContext).r = _x
	}
	{
		p.SetState(36)

		var _x = p.Arguments()


		localctx.(*RelationContext).arg = _x
	}
	{
		p.SetState(37)
		p.Match(R2QueryParserT__1)
	}



	return localctx
}


// IRelationsContext is an interface to support dynamic dispatch.
type IRelationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationsContext differentiates from other interfaces.
	IsRelationsContext()
}

type RelationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationsContext() *RelationsContext {
	var p = new(RelationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_relations
	return p
}

func (*RelationsContext) IsRelationsContext() {}

func NewRelationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationsContext {
	var p = new(RelationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_relations

	return p
}

func (s *RelationsContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationsContext) AllRelation() []IRelationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationContext)(nil)).Elem())
	var tst = make([]IRelationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationContext)
		}
	}

	return tst
}

func (s *RelationsContext) Relation(i int) IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *RelationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2QueryParserCOMMA)
}

func (s *RelationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2QueryParserCOMMA, i)
}

func (s *RelationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterRelations(s)
	}
}

func (s *RelationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitRelations(s)
	}
}

func (s *RelationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitRelations(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *R2QueryParser) Relations() (localctx IRelationsContext) {
	localctx = NewRelationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, R2QueryParserRULE_relations)
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
		p.SetState(39)
		p.Relation()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == R2QueryParserCOMMA {
		{
			p.SetState(40)
			p.Match(R2QueryParserCOMMA)
		}
		{
			p.SetState(41)
			p.Relation()
		}


		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = R2QueryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2QueryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type AddElementsContext struct {
	*ExpressionContext
}

func NewAddElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddElementsContext {
	var p = new(AddElementsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddElementsContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *AddElementsContext) ADD() antlr.TerminalNode {
	return s.GetToken(R2QueryParserADD, 0)
}

func (s *AddElementsContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}


func (s *AddElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterAddElements(s)
	}
}

func (s *AddElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitAddElements(s)
	}
}

func (s *AddElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitAddElements(s)

	default:
		return t.VisitChildren(s)
	}
}


type DeleteRelationsContext struct {
	*ExpressionContext
}

func NewDeleteRelationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteRelationsContext {
	var p = new(DeleteRelationsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DeleteRelationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteRelationsContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *DeleteRelationsContext) DELETE() antlr.TerminalNode {
	return s.GetToken(R2QueryParserDELETE, 0)
}

func (s *DeleteRelationsContext) AllArguments() []IArgumentsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentsContext)(nil)).Elem())
	var tst = make([]IArgumentsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentsContext)
		}
	}

	return tst
}

func (s *DeleteRelationsContext) Arguments(i int) IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DeleteRelationsContext) Rel() IRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelContext)
}


func (s *DeleteRelationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterDeleteRelations(s)
	}
}

func (s *DeleteRelationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitDeleteRelations(s)
	}
}

func (s *DeleteRelationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitDeleteRelations(s)

	default:
		return t.VisitChildren(s)
	}
}


type GetElementRelationContext struct {
	*ExpressionContext
}

func NewGetElementRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetElementRelationContext {
	var p = new(GetElementRelationContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GetElementRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetElementRelationContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *GetElementRelationContext) GET() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGET, 0)
}

func (s *GetElementRelationContext) Relations() IRelationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationsContext)
}


func (s *GetElementRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterGetElementRelation(s)
	}
}

func (s *GetElementRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitGetElementRelation(s)
	}
}

func (s *GetElementRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitGetElementRelation(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddRelationsContext struct {
	*ExpressionContext
}

func NewAddRelationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddRelationsContext {
	var p = new(AddRelationsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddRelationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddRelationsContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *AddRelationsContext) ADD() antlr.TerminalNode {
	return s.GetToken(R2QueryParserADD, 0)
}

func (s *AddRelationsContext) AllArguments() []IArgumentsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentsContext)(nil)).Elem())
	var tst = make([]IArgumentsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentsContext)
		}
	}

	return tst
}

func (s *AddRelationsContext) Arguments(i int) IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *AddRelationsContext) Rel() IRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelContext)
}


func (s *AddRelationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterAddRelations(s)
	}
}

func (s *AddRelationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitAddRelations(s)
	}
}

func (s *AddRelationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitAddRelations(s)

	default:
		return t.VisitChildren(s)
	}
}


type CreateGraphContext struct {
	*ExpressionContext
}

func NewCreateGraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreateGraphContext {
	var p = new(CreateGraphContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CreateGraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateGraphContext) CREATE() antlr.TerminalNode {
	return s.GetToken(R2QueryParserCREATE, 0)
}

func (s *CreateGraphContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}


func (s *CreateGraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterCreateGraph(s)
	}
}

func (s *CreateGraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitCreateGraph(s)
	}
}

func (s *CreateGraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitCreateGraph(s)

	default:
		return t.VisitChildren(s)
	}
}


type GetElementsContext struct {
	*ExpressionContext
}

func NewGetElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetElementsContext {
	var p = new(GetElementsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GetElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetElementsContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *GetElementsContext) GET() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGET, 0)
}

func (s *GetElementsContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}


func (s *GetElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterGetElements(s)
	}
}

func (s *GetElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitGetElements(s)
	}
}

func (s *GetElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitGetElements(s)

	default:
		return t.VisitChildren(s)
	}
}


type DeleteGraphContext struct {
	*ExpressionContext
}

func NewDeleteGraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteGraphContext {
	var p = new(DeleteGraphContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DeleteGraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteGraphContext) DELETE() antlr.TerminalNode {
	return s.GetToken(R2QueryParserDELETE, 0)
}

func (s *DeleteGraphContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}


func (s *DeleteGraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterDeleteGraph(s)
	}
}

func (s *DeleteGraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitDeleteGraph(s)
	}
}

func (s *DeleteGraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitDeleteGraph(s)

	default:
		return t.VisitChildren(s)
	}
}


type DeleteElementsContext struct {
	*ExpressionContext
}

func NewDeleteElementsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteElementsContext {
	var p = new(DeleteElementsContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DeleteElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteElementsContext) GRAPH_NAME() antlr.TerminalNode {
	return s.GetToken(R2QueryParserGRAPH_NAME, 0)
}

func (s *DeleteElementsContext) DELETE() antlr.TerminalNode {
	return s.GetToken(R2QueryParserDELETE, 0)
}

func (s *DeleteElementsContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}


func (s *DeleteElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.EnterDeleteElements(s)
	}
}

func (s *DeleteElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2QueryListener); ok {
		listenerT.ExitDeleteElements(s)
	}
}

func (s *DeleteElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2QueryVisitor:
		return t.VisitDeleteElements(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *R2QueryParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, R2QueryParserRULE_expression)

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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCreateGraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(R2QueryParserCREATE)
		}
		{
			p.SetState(48)
			p.Match(R2QueryParserGRAPH_NAME)
		}


	case 2:
		localctx = NewDeleteGraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(R2QueryParserDELETE)
		}
		{
			p.SetState(50)
			p.Match(R2QueryParserGRAPH_NAME)
		}


	case 3:
		localctx = NewGetElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(52)
			p.Match(R2QueryParserGET)
		}
		{
			p.SetState(53)
			p.Arguments()
		}


	case 4:
		localctx = NewGetElementRelationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(54)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(55)
			p.Match(R2QueryParserGET)
		}
		{
			p.SetState(56)
			p.Relations()
		}


	case 5:
		localctx = NewAddElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(58)
			p.Match(R2QueryParserADD)
		}
		{
			p.SetState(59)
			p.Arguments()
		}


	case 6:
		localctx = NewDeleteElementsContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(61)
			p.Match(R2QueryParserDELETE)
		}
		{
			p.SetState(62)
			p.Arguments()
		}


	case 7:
		localctx = NewAddRelationsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(63)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(64)
			p.Match(R2QueryParserADD)
		}
		{
			p.SetState(65)
			p.Arguments()
		}
		{
			p.SetState(66)
			p.Rel()
		}
		{
			p.SetState(67)
			p.Arguments()
		}


	case 8:
		localctx = NewDeleteRelationsContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(69)
			p.Match(R2QueryParserGRAPH_NAME)
		}
		{
			p.SetState(70)
			p.Match(R2QueryParserDELETE)
		}
		{
			p.SetState(71)
			p.Arguments()
		}
		{
			p.SetState(72)
			p.Rel()
		}
		{
			p.SetState(73)
			p.Arguments()
		}

	}


	return localctx
}


