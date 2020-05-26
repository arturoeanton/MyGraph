// Code generated from /Users/arturoeliasanton/go/src/r2Server/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query // R2Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseR2QueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseR2QueryVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitRel(ctx *RelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitRelations(ctx *RelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitCreateGraph(ctx *CreateGraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitDeleteGraph(ctx *DeleteGraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitGetElements(ctx *GetElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitGetElementRelation(ctx *GetElementRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitAddElements(ctx *AddElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitDeleteElements(ctx *DeleteElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitAddRelations(ctx *AddRelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseR2QueryVisitor) VisitDeleteRelations(ctx *DeleteRelationsContext) interface{} {
	return v.VisitChildren(ctx)
}
