package main

import "C"
import (
	"MyGraph/r2query"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"strings"
)

type r2queryListener struct {
	manager *ClientManager
	*r2query.BaseR2QueryListener
	rm       *RelationManager
	distinct map[string]bool
	result   []string
	messages   []string
}

// ExitStart is called when exiting the start production.
func (l *r2queryListener) ExitStart(c *r2query.StartContext) {}

// ExitArguments is called when exiting the arguments production.
func (l *r2queryListener) ExitArguments(c *r2query.ArgumentsContext) {}

// ExitRel is called when exiting the rel production.
func (l *r2queryListener) ExitRel(c *r2query.RelContext) {}

func valueFormatter(element1, element2, tag, direction string) string {
	return fmt.Sprintf(
		"{\"element1\":%s, \"element2\":%s, \"tag\":\"%s\", \"direction\":\"%s\"}",
		element1, element2, tag, direction)
}

// EnterRelation is called when entering the relation production.
func (l *r2queryListener) EnterRelation(c *r2query.RelationContext) {
	//for j:= 0; j< c.Relations().GetChildCount();j=j+3 {
	tag := c.GetR().GetTag().GetText()
	//c.GetParent().(*r2query.RelationsContext).GetParent().(*r2query.GetElementRelationContext).GRAPH_NAME()

	if c.GetR().GetOp1().GetTokenType() == r2query.R2QueryParserDIRECTION_REL_L && c.GetR().GetOp2().GetTokenType() == r2query.R2QueryParserDIRECTION_REL_R {
		for _, term := range c.GetArg().(*r2query.ArgumentsContext).AllSTRING() {
			element := term.GetText()
			for _, item := range l.rm.GetElementAll(l.rm.objects[element].relations[tag]) {
				mset, ok := l.rm.objects[item].relations[tag]
				if !ok {
					continue
				}
				if mset.Contains(l.rm.ids[element]) {
					value := valueFormatter(item, element, tag, "<->")
					if _, ok := l.distinct[value]; !ok {
						value2 := valueFormatter(element, item, tag, "<->")
						if _, ok := l.distinct[value2]; !ok {
							l.distinct[value] = true
							l.result = append(l.result, value)
						}
					}
					continue
				}
			}
		}
		return
	}

	if c.GetR().GetOp1().GetTokenType() == r2query.R2QueryParserDIRECTION_REL_L {
		for _, term := range c.GetArg().(*r2query.ArgumentsContext).AllSTRING() {
			element := term.GetText()
			for _, item := range l.rm.GetElementAll(l.rm.objects[element].relations[tag]) {
				mset, ok := l.rm.objects[item].relations[tag]
				if ok {
					if mset.Contains(l.rm.ids[element]) {
						value := valueFormatter(item, element, tag, "<->")
						if _, ok := l.distinct[value]; !ok {
							value2 := valueFormatter(element, item, tag, "<->")
							if _, ok := l.distinct[value2]; !ok {
								l.distinct[value] = true
								l.result = append(l.result, value)
							}
						}
						continue
					}
				}
				value := valueFormatter(item, element, tag, "<-")
				if _, ok := l.distinct[value]; !ok {
					value2 := valueFormatter(element, item, tag, "->")
					if _, ok := l.distinct[value2]; !ok {
						l.distinct[value] = true
						l.result = append(l.result, value)
					}
				}
				continue
			}
		}
		return
	}

	if c.GetR().GetOp1().GetTokenType() == r2query.R2QueryParserDIRECTION_REL_R {
		tagReverse := "_INTERNAL@[REVERSE-HIDE]-" + tag
		for _, term := range c.GetArg().(*r2query.ArgumentsContext).AllSTRING() {
			element := term.GetText()
			for _, item := range l.rm.GetElementAll(l.rm.objects[element].relations[tagReverse]) {
				mset, ok := l.rm.objects[element].relations[tag]
				if ok {
					if mset.Contains(l.rm.ids[item]) {
						value := valueFormatter(item, element, tag, "<->")
						if _, ok := l.distinct[value]; !ok {
							value2 := valueFormatter(element, item, tag, "<->")
							if _, ok := l.distinct[value2]; !ok {
								l.distinct[value] = true
								l.result = append(l.result, value)
							}
						}
						continue
					}
				}
				value := valueFormatter(item, element, tag, "->")
				if _, ok := l.distinct[value]; !ok {
					value2 := valueFormatter(element, item, tag, "<-")
					if _, ok := l.distinct[value2]; !ok {
						l.distinct[value] = true
						l.result = append(l.result, value)
					}
				}
				continue
			}
		}
		return
	}
	//}
	//l.result +=fmt.Sprintln("{", i, " Count}")
}

// ExitRelations is called when exiting the relations production.
func (l *r2queryListener) ExitRelations(c *r2query.RelationsContext) {}

// ExitCreateGraph is called when exiting the CreateGraph production.
func (l *r2queryListener) ExitCreateGraph(c *r2query.CreateGraphContext) {
	if graph == nil {
		graph = make(map[string]RelationManager)
	}
	if _, ok := graph[c.GRAPH_NAME().GetText()]; ok {
		l.messages = append(l.result, "\"EXIST GRAPH OK\"")
		return
	}
	graph[c.GRAPH_NAME().GetText()] = NewRelationManager()
	l.messages = append(l.messages, "\"CREATE GRAPH OK\"")
}

// ExitDeleteGraph is called when exiting the DeleteGraph production.
func (l *r2queryListener) ExitDeleteGraph(c *r2query.DeleteGraphContext) {
	if graph == nil {
		graph = make(map[string]RelationManager)
	}
	if _, ok := graph[c.GRAPH_NAME().GetText()]; ok {
		delete(graph, c.GRAPH_NAME().GetText())
		l.messages = append(l.messages,"\"DELETE GRAPH OK\"")
		return
	}
	l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
}

// ExitGetElements is called when exiting the GetElements production.
func (l *r2queryListener) ExitGetElements(c *r2query.GetElementsContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
		return
	}
	for _, v := range c.Arguments().GetChildren() {
		vv := v.(*antlr.TerminalNodeImpl)
		if vv.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
			e1 := vv.GetText()
			for tag, r := range mr.objects[e1].relations {
				if strings.HasPrefix(tag, "_INTERNAL@[REVERSE-HIDE]-") {
					continue
				}
				for _, e2 := range mr.GetElementAll(r) {
					rel1 := ""
					rel2 := "->"
					mset, ok := mr.objects[e2].relations[tag]
					if ok {
						if mset.Contains(mr.ids[e1]) {
							rel1 = "<"
						}
					}
					value := valueFormatter(e1, e2, tag, rel1+rel2)
					if _, ok := l.distinct[value]; !ok {
						l.distinct[value] = true
						l.result = append(l.result, value)
					}
				}
			}
		}
	}
}

// ExitGetElementRelation is called when exiting the GetElementRelation production.
func (l *r2queryListener) EnterGetElementRelation(c *r2query.GetElementRelationContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages,"\"GRAPH NOT FOUND \"")
		panic("\"GRAPH NOT FOUND \"")
	}
	l.rm = &mr
	//i := 0
}

// ExitAddElements is called when exiting the AddElements production.
func (l *r2queryListener) ExitAddElements(c *r2query.AddElementsContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
		return
	}
	for _, v := range c.Arguments().GetChildren() {
		vv := v.(*antlr.TerminalNodeImpl)
		if vv.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
			if mr.AddObject(vv.GetText()) {
			}
		}
	}
}

// ExitDeleteElements is called when exiting the DeleteElements production.
func (l *r2queryListener) ExitDeleteElements(c *r2query.DeleteElementsContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
		return
	}
	for _, v := range c.Arguments().GetChildren() {
		vv := v.(*antlr.TerminalNodeImpl)
		if vv.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
			if mr.RemoveObject(vv.GetText()) {
			}
		}
	}
}

// ExitAddRelations is called when exiting the AddRelations production.
func (l *r2queryListener) ExitAddRelations(c *r2query.AddRelationsContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
		return
	}
	for _, v0 := range c.Arguments(0).GetChildren() {
		vv0 := v0.(*antlr.TerminalNodeImpl)
		if vv0.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
			op1 := c.Rel().GetOp1().GetTokenType()
			op2 := c.Rel().GetOp2().GetTokenType()
			for _, v1 := range c.Arguments(1).GetChildren() {
				vv1 := v1.(*antlr.TerminalNodeImpl)
				if vv1.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
					if op1 == r2query.R2QueryParserDIRECTION_REL_L && op2 == r2query.R2QueryParserDIRECTION_REL_R {
						mr.RelationBi(vv0.GetText(), c.Rel().GetTag().GetText(), vv1.GetText())
						continue
					}
					if op1 == r2query.R2QueryParserDIRECTION_REL_L && op2 == r2query.R2QueryLexerMIDDLE {
						mr.RelationL(vv0.GetText(), c.Rel().GetTag().GetText(), vv1.GetText())
						continue
					}
					if op1 == r2query.R2QueryParserDIRECTION_REL_R && op2 == r2query.R2QueryLexerMIDDLE {
						mr.RelationR(vv0.GetText(), c.Rel().GetTag().GetText(), vv1.GetText())
						continue
					}
				}
			}
		}
	}
}

// ExitDeleteRelations is called when exiting the DeleteRelations production.
func (l *r2queryListener) ExitDeleteRelations(c *r2query.DeleteRelationsContext) {
	mr, ok := graph[c.GRAPH_NAME().GetText()]
	if !ok {
		l.messages = append(l.messages, "\"GRAPH NOT FOUND \"")
		return
	}
	i := 0
	for _, v0 := range c.Arguments(0).GetChildren() {
		vv0 := v0.(*antlr.TerminalNodeImpl)
		if vv0.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
			for _, v1 := range c.Arguments(1).GetChildren() {
				vv1 := v1.(*antlr.TerminalNodeImpl)
				if vv1.GetSymbol().GetTokenType() == r2query.R2QueryParserSTRING {
					mr.RemoveRelation(vv0.GetText(), c.Rel().GetTag().GetText(), vv1.GetText())
					i++
				}
			}
		}
	}
	l.messages = append(l.messages, fmt.Sprint("\"", i, " Count\""))
}

func (l *r2queryListener) execute(cmd string) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprint("Recovered in R2QueryParser", r)
			l.messages = append(l.messages, "\""+str+"\"")
			log.Print(str)
		}
	}()
	is := antlr.NewInputStream(cmd)
	lexer := r2query.NewR2QueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := r2query.NewR2QueryParser(stream)
	l.distinct = make(map[string]bool)
	antlr.ParseTreeWalkerDefault.Walk(l, parser.Start())
}

/*
CREATE A
A ADD "A1","A2"-"GHJ"->"AAA"
A ADD "A3","A4"<-"GHJ"->"AAA3"
A ADD "A5","A6"<-"GHJ"-"AAA5"
A GET "A1","A2","AAA"
A GET "A3","A4", "AAA3"
A GET "A5","A6","AAA5"
*/
