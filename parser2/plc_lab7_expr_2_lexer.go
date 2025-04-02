// Code generated from PLC_Lab7_expr_2.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser2

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 7, 9, 42, 10, 9, 12, 9, 14, 9, 45, 11, 9, 3, 10, 3,
	10, 7, 10, 49, 10, 10, 12, 10, 14, 10, 52, 11, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 6, 11, 58, 10, 11, 13, 11, 14, 11, 59, 3, 12, 6, 12, 63, 10, 12,
	13, 12, 14, 12, 64, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 7, 3, 2, 51, 59,
	3, 2, 50, 59, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 71, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9,
	31, 3, 2, 2, 2, 11, 33, 3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 37, 3, 2, 2,
	2, 17, 39, 3, 2, 2, 2, 19, 46, 3, 2, 2, 2, 21, 53, 3, 2, 2, 2, 23, 62,
	3, 2, 2, 2, 25, 26, 7, 61, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 44, 2, 2,
	28, 6, 3, 2, 2, 2, 29, 30, 7, 49, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 45,
	2, 2, 32, 10, 3, 2, 2, 2, 33, 34, 7, 47, 2, 2, 34, 12, 3, 2, 2, 2, 35,
	36, 7, 42, 2, 2, 36, 14, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 16, 3, 2,
	2, 2, 39, 43, 9, 2, 2, 2, 40, 42, 9, 3, 2, 2, 41, 40, 3, 2, 2, 2, 42, 45,
	3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 18, 3, 2, 2, 2,
	45, 43, 3, 2, 2, 2, 46, 50, 7, 50, 2, 2, 47, 49, 9, 4, 2, 2, 48, 47, 3,
	2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	20, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 54, 7, 50, 2, 2, 54, 55, 7, 122,
	2, 2, 55, 57, 3, 2, 2, 2, 56, 58, 9, 5, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 22, 3, 2, 2, 2,
	61, 63, 9, 6, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3,
	2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 8, 12, 2, 2, 67,
	24, 3, 2, 2, 2, 7, 2, 43, 50, 59, 64, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "INT", "OCT", "HEX", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "INT", "OCT", "HEX",
	"WS",
}

type PLC_Lab7_expr_2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPLC_Lab7_expr_2Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PLC_Lab7_expr_2Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPLC_Lab7_expr_2Lexer(input antlr.CharStream) *PLC_Lab7_expr_2Lexer {
	l := new(PLC_Lab7_expr_2Lexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PLC_Lab7_expr_2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PLC_Lab7_expr_2Lexer tokens.
const (
	PLC_Lab7_expr_2LexerT__0 = 1
	PLC_Lab7_expr_2LexerT__1 = 2
	PLC_Lab7_expr_2LexerT__2 = 3
	PLC_Lab7_expr_2LexerT__3 = 4
	PLC_Lab7_expr_2LexerT__4 = 5
	PLC_Lab7_expr_2LexerT__5 = 6
	PLC_Lab7_expr_2LexerT__6 = 7
	PLC_Lab7_expr_2LexerINT  = 8
	PLC_Lab7_expr_2LexerOCT  = 9
	PLC_Lab7_expr_2LexerHEX  = 10
	PLC_Lab7_expr_2LexerWS   = 11
)
