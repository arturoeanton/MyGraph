// Code generated from /Users/arturoeliasanton/github/arturoeanton/MyGraph/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query // R2Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseR2QueryListener is a complete listener for a parse tree produced by R2QueryParser.
type BaseR2QueryListener struct{}

var _ R2QueryListener = &BaseR2QueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseR2QueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseR2QueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseR2QueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseR2QueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseR2QueryListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseR2QueryListener) ExitStart(ctx *StartContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseR2QueryListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseR2QueryListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterRel is called when production rel is entered.
func (s *BaseR2QueryListener) EnterRel(ctx *RelContext) {}

// ExitRel is called when production rel is exited.
func (s *BaseR2QueryListener) ExitRel(ctx *RelContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseR2QueryListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseR2QueryListener) ExitRelation(ctx *RelationContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseR2QueryListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseR2QueryListener) ExitRelations(ctx *RelationsContext) {}

// EnterCreateGraph is called when production CreateGraph is entered.
func (s *BaseR2QueryListener) EnterCreateGraph(ctx *CreateGraphContext) {}

// ExitCreateGraph is called when production CreateGraph is exited.
func (s *BaseR2QueryListener) ExitCreateGraph(ctx *CreateGraphContext) {}

// EnterDeleteGraph is called when production DeleteGraph is entered.
func (s *BaseR2QueryListener) EnterDeleteGraph(ctx *DeleteGraphContext) {}

// ExitDeleteGraph is called when production DeleteGraph is exited.
func (s *BaseR2QueryListener) ExitDeleteGraph(ctx *DeleteGraphContext) {}

// EnterGetElements is called when production GetElements is entered.
func (s *BaseR2QueryListener) EnterGetElements(ctx *GetElementsContext) {}

// ExitGetElements is called when production GetElements is exited.
func (s *BaseR2QueryListener) ExitGetElements(ctx *GetElementsContext) {}

// EnterGetElementRelation is called when production GetElementRelation is entered.
func (s *BaseR2QueryListener) EnterGetElementRelation(ctx *GetElementRelationContext) {}

// ExitGetElementRelation is called when production GetElementRelation is exited.
func (s *BaseR2QueryListener) ExitGetElementRelation(ctx *GetElementRelationContext) {}

// EnterAddElements is called when production AddElements is entered.
func (s *BaseR2QueryListener) EnterAddElements(ctx *AddElementsContext) {}

// ExitAddElements is called when production AddElements is exited.
func (s *BaseR2QueryListener) ExitAddElements(ctx *AddElementsContext) {}

// EnterDeleteElements is called when production DeleteElements is entered.
func (s *BaseR2QueryListener) EnterDeleteElements(ctx *DeleteElementsContext) {}

// ExitDeleteElements is called when production DeleteElements is exited.
func (s *BaseR2QueryListener) ExitDeleteElements(ctx *DeleteElementsContext) {}

// EnterAddRelations is called when production AddRelations is entered.
func (s *BaseR2QueryListener) EnterAddRelations(ctx *AddRelationsContext) {}

// ExitAddRelations is called when production AddRelations is exited.
func (s *BaseR2QueryListener) ExitAddRelations(ctx *AddRelationsContext) {}

// EnterDeleteRelations is called when production DeleteRelations is entered.
func (s *BaseR2QueryListener) EnterDeleteRelations(ctx *DeleteRelationsContext) {}

// ExitDeleteRelations is called when production DeleteRelations is exited.
func (s *BaseR2QueryListener) ExitDeleteRelations(ctx *DeleteRelationsContext) {}
