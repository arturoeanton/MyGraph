// Code generated from /Users/arturoeliasanton/go/src/r2Server/R2Query.g4 by ANTLR 4.8. DO NOT EDIT.

package r2query
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 131, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 81, 
	10, 14, 12, 14, 14, 14, 84, 11, 14, 3, 15, 6, 15, 87, 10, 15, 13, 15, 14, 
	15, 88, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 98, 10, 
	16, 12, 16, 14, 16, 101, 11, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 
	16, 7, 16, 109, 10, 16, 12, 16, 14, 16, 112, 11, 16, 3, 16, 3, 16, 3, 16, 
	3, 16, 7, 16, 118, 10, 16, 12, 16, 14, 16, 121, 11, 16, 3, 16, 5, 16, 124, 
	10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 130, 10, 17, 3, 119, 2, 18, 
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 2, 3, 2, 8, 4, 2, 67, 92, 99, 124, 
	5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 36, 
	36, 3, 2, 41, 41, 4, 2, 8223, 8223, 8245, 8245, 2, 142, 2, 3, 3, 2, 2, 
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 37, 
	3, 2, 2, 2, 7, 39, 3, 2, 2, 2, 9, 46, 3, 2, 2, 2, 11, 53, 3, 2, 2, 2, 13, 
	57, 3, 2, 2, 2, 15, 60, 3, 2, 2, 2, 17, 64, 3, 2, 2, 2, 19, 66, 3, 2, 2, 
	2, 21, 70, 3, 2, 2, 2, 23, 72, 3, 2, 2, 2, 25, 75, 3, 2, 2, 2, 27, 78, 
	3, 2, 2, 2, 29, 86, 3, 2, 2, 2, 31, 123, 3, 2, 2, 2, 33, 129, 3, 2, 2, 
	2, 35, 36, 7, 93, 2, 2, 36, 4, 3, 2, 2, 2, 37, 38, 7, 95, 2, 2, 38, 6, 
	3, 2, 2, 2, 39, 40, 7, 69, 2, 2, 40, 41, 7, 84, 2, 2, 41, 42, 7, 71, 2, 
	2, 42, 43, 7, 67, 2, 2, 43, 44, 7, 86, 2, 2, 44, 45, 7, 71, 2, 2, 45, 8, 
	3, 2, 2, 2, 46, 47, 7, 70, 2, 2, 47, 48, 7, 71, 2, 2, 48, 49, 7, 78, 2, 
	2, 49, 50, 7, 71, 2, 2, 50, 51, 7, 86, 2, 2, 51, 52, 7, 71, 2, 2, 52, 10, 
	3, 2, 2, 2, 53, 54, 7, 67, 2, 2, 54, 55, 7, 70, 2, 2, 55, 56, 7, 70, 2, 
	2, 56, 12, 3, 2, 2, 2, 57, 58, 7, 81, 2, 2, 58, 59, 7, 84, 2, 2, 59, 14, 
	3, 2, 2, 2, 60, 61, 7, 67, 2, 2, 61, 62, 7, 80, 2, 2, 62, 63, 7, 70, 2, 
	2, 63, 16, 3, 2, 2, 2, 64, 65, 7, 46, 2, 2, 65, 18, 3, 2, 2, 2, 66, 67, 
	7, 73, 2, 2, 67, 68, 7, 71, 2, 2, 68, 69, 7, 86, 2, 2, 69, 20, 3, 2, 2, 
	2, 70, 71, 7, 47, 2, 2, 71, 22, 3, 2, 2, 2, 72, 73, 7, 47, 2, 2, 73, 74, 
	7, 64, 2, 2, 74, 24, 3, 2, 2, 2, 75, 76, 7, 62, 2, 2, 76, 77, 7, 47, 2, 
	2, 77, 26, 3, 2, 2, 2, 78, 82, 9, 2, 2, 2, 79, 81, 9, 3, 2, 2, 80, 79, 
	3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 
	83, 28, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 87, 9, 4, 2, 2, 86, 85, 3, 
	2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 
	90, 3, 2, 2, 2, 90, 91, 8, 15, 2, 2, 91, 30, 3, 2, 2, 2, 92, 99, 7, 36, 
	2, 2, 93, 98, 5, 33, 17, 2, 94, 95, 7, 36, 2, 2, 95, 98, 7, 36, 2, 2, 96, 
	98, 10, 5, 2, 2, 97, 93, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2, 97, 96, 3, 2, 
	2, 2, 98, 101, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 
	102, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 124, 7, 36, 2, 2, 103, 110, 
	7, 41, 2, 2, 104, 109, 5, 33, 17, 2, 105, 106, 7, 41, 2, 2, 106, 109, 7, 
	41, 2, 2, 107, 109, 10, 6, 2, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 
	2, 2, 108, 107, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 
	110, 111, 3, 2, 2, 2, 111, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 
	124, 7, 41, 2, 2, 114, 119, 7, 8222, 2, 2, 115, 118, 5, 33, 17, 2, 116, 
	118, 11, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 116, 3, 2, 2, 2, 118, 121, 
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 122, 3, 2, 
	2, 2, 121, 119, 3, 2, 2, 2, 122, 124, 9, 7, 2, 2, 123, 92, 3, 2, 2, 2, 
	123, 103, 3, 2, 2, 2, 123, 114, 3, 2, 2, 2, 124, 32, 3, 2, 2, 2, 125, 126, 
	7, 98, 2, 2, 126, 130, 7, 41, 2, 2, 127, 128, 7, 98, 2, 2, 128, 130, 7, 
	36, 2, 2, 129, 125, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 34, 3, 2, 2, 
	2, 13, 2, 82, 88, 97, 99, 108, 110, 117, 119, 123, 129, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'CREATE'", "'DELETE'", "'ADD'", "'OR'", "'AND'", "','", 
	"'GET'", "'-'", "'->'", "'<-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "CREATE", "DELETE", "ADD", "OR", "AND", "COMMA", "GET", "MIDDLE", 
	"DIRECTION_REL_R", "DIRECTION_REL_L", "GRAPH_NAME", "WHITESPACE", "STRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "CREATE", "DELETE", "ADD", "OR", "AND", "COMMA", "GET", 
	"MIDDLE", "DIRECTION_REL_R", "DIRECTION_REL_L", "GRAPH_NAME", "WHITESPACE", 
	"STRING", "Escape",
}

type R2QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewR2QueryLexer(input antlr.CharStream) *R2QueryLexer {

	l := new(R2QueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "R2Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// R2QueryLexer tokens.
const (
	R2QueryLexerT__0 = 1
	R2QueryLexerT__1 = 2
	R2QueryLexerCREATE = 3
	R2QueryLexerDELETE = 4
	R2QueryLexerADD = 5
	R2QueryLexerOR = 6
	R2QueryLexerAND = 7
	R2QueryLexerCOMMA = 8
	R2QueryLexerGET = 9
	R2QueryLexerMIDDLE = 10
	R2QueryLexerDIRECTION_REL_R = 11
	R2QueryLexerDIRECTION_REL_L = 12
	R2QueryLexerGRAPH_NAME = 13
	R2QueryLexerWHITESPACE = 14
	R2QueryLexerSTRING = 15
)

