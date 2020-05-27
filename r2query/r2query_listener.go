// Code generated from /Users/arturoeliasanton/github/arturoeanton/MyGraph/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query // R2Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// R2QueryListener is a complete listener for a parse tree produced by R2QueryParser.
type R2QueryListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterRel is called when entering the rel production.
	EnterRel(c *RelContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterRelations is called when entering the relations production.
	EnterRelations(c *RelationsContext)

	// EnterCreateGraph is called when entering the CreateGraph production.
	EnterCreateGraph(c *CreateGraphContext)

	// EnterDeleteGraph is called when entering the DeleteGraph production.
	EnterDeleteGraph(c *DeleteGraphContext)

	// EnterGetElements is called when entering the GetElements production.
	EnterGetElements(c *GetElementsContext)

	// EnterGetElementRelation is called when entering the GetElementRelation production.
	EnterGetElementRelation(c *GetElementRelationContext)

	// EnterAddElements is called when entering the AddElements production.
	EnterAddElements(c *AddElementsContext)

	// EnterDeleteElements is called when entering the DeleteElements production.
	EnterDeleteElements(c *DeleteElementsContext)

	// EnterAddRelations is called when entering the AddRelations production.
	EnterAddRelations(c *AddRelationsContext)

	// EnterDeleteRelations is called when entering the DeleteRelations production.
	EnterDeleteRelations(c *DeleteRelationsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitRel is called when exiting the rel production.
	ExitRel(c *RelContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitRelations is called when exiting the relations production.
	ExitRelations(c *RelationsContext)

	// ExitCreateGraph is called when exiting the CreateGraph production.
	ExitCreateGraph(c *CreateGraphContext)

	// ExitDeleteGraph is called when exiting the DeleteGraph production.
	ExitDeleteGraph(c *DeleteGraphContext)

	// ExitGetElements is called when exiting the GetElements production.
	ExitGetElements(c *GetElementsContext)

	// ExitGetElementRelation is called when exiting the GetElementRelation production.
	ExitGetElementRelation(c *GetElementRelationContext)

	// ExitAddElements is called when exiting the AddElements production.
	ExitAddElements(c *AddElementsContext)

	// ExitDeleteElements is called when exiting the DeleteElements production.
	ExitDeleteElements(c *DeleteElementsContext)

	// ExitAddRelations is called when exiting the AddRelations production.
	ExitAddRelations(c *AddRelationsContext)

	// ExitDeleteRelations is called when exiting the DeleteRelations production.
	ExitDeleteRelations(c *DeleteRelationsContext)
}
