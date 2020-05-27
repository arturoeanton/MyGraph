// Code generated from /Users/arturoeliasanton/github/arturoeanton/MyGraph/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query // R2Query
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by R2QueryParser.
type R2QueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by R2QueryParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by R2QueryParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#rel.
	VisitRel(ctx *RelContext) interface{}

	// Visit a parse tree produced by R2QueryParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by R2QueryParser#relations.
	VisitRelations(ctx *RelationsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#CreateGraph.
	VisitCreateGraph(ctx *CreateGraphContext) interface{}

	// Visit a parse tree produced by R2QueryParser#DeleteGraph.
	VisitDeleteGraph(ctx *DeleteGraphContext) interface{}

	// Visit a parse tree produced by R2QueryParser#GetElements.
	VisitGetElements(ctx *GetElementsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#GetElementRelation.
	VisitGetElementRelation(ctx *GetElementRelationContext) interface{}

	// Visit a parse tree produced by R2QueryParser#AddElements.
	VisitAddElements(ctx *AddElementsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#DeleteElements.
	VisitDeleteElements(ctx *DeleteElementsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#AddRelations.
	VisitAddRelations(ctx *AddRelationsContext) interface{}

	// Visit a parse tree produced by R2QueryParser#DeleteRelations.
	VisitDeleteRelations(ctx *DeleteRelationsContext) interface{}

}