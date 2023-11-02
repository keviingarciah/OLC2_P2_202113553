// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'['", "']'", "','", "'.append'", "'('", "')'", "'.removeLast'",
		"'.remove'", "'at'", "'struct'", "'{'", "'}'", "'mutating'", "'.'",
		"'->'", "'...'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='", "'>'",
		"'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'",
		"'.count'", "'.isEmpty'", "'Int'", "'Float'", "'String'", "'Bool'",
		"'Character'", "'nil'", "'var'", "'let'", "'func'", "'inout'", "'&'",
		"'print'", "'if'", "'else'", "'switch'", "'case'", "'default'", "'while'",
		"'for'", "'in'", "'guard'", "'continue'", "'break'", "'return'", "';'",
		"':'", "'='", "'?'", "'+='", "'-='", "'_'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "NIL", "VAR", "LET",
		"FUNC", "INOUT", "REFERENCE", "PRINT", "IF", "ELSE", "SWITCH", "CASE",
		"DEFAULT", "WHILE", "FOR", "IN", "GUARD", "CONTINUE", "BREAK", "RETURN",
		"SEMICOLON", "COLON", "EQUAL", "QUESTION_MARK", "INCREMENT", "DECREMENT",
		"UNDERSCORE", "DIGIT", "STR", "CHAR", "ID", "WHITESPACE", "MULTI_COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"start", "block", "stmts", "varDeclaration", "varAssignment", "letDeclaration",
		"vectorDeclaration", "valuesVectorDeclaration", "vectorType", "vectorAppend",
		"vectorRemoveLast", "vectorRemoveAt", "matrixDeclaration", "valuesMatrixDeclaration",
		"structDeclaration", "structAttribute", "structFunction", "structAccess",
		"structAssignment", "funcDeclaration", "paramsFuncDeclaration", "funcCall",
		"paramsFuncCall", "printStmt", "ifStmt", "switchStmt", "cases", "caseStmt",
		"defaultStmt", "whileStmt", "forStmt", "range", "guardStmt", "breakStmt",
		"continueStmt", "returnStmt", "expr", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 589, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 81, 8, 1, 10, 1, 12, 1, 84,
		9, 1, 1, 2, 1, 2, 3, 2, 88, 8, 2, 1, 2, 1, 2, 3, 2, 92, 8, 2, 1, 2, 1,
		2, 3, 2, 96, 8, 2, 1, 2, 1, 2, 3, 2, 100, 8, 2, 1, 2, 1, 2, 3, 2, 104,
		8, 2, 1, 2, 1, 2, 3, 2, 108, 8, 2, 1, 2, 1, 2, 3, 2, 112, 8, 2, 1, 2, 1,
		2, 3, 2, 116, 8, 2, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 3, 2, 124,
		8, 2, 1, 2, 1, 2, 3, 2, 128, 8, 2, 1, 2, 1, 2, 3, 2, 132, 8, 2, 1, 2, 1,
		2, 3, 2, 136, 8, 2, 1, 2, 1, 2, 3, 2, 140, 8, 2, 1, 2, 1, 2, 3, 2, 144,
		8, 2, 1, 2, 1, 2, 3, 2, 148, 8, 2, 1, 2, 1, 2, 3, 2, 152, 8, 2, 1, 2, 1,
		2, 3, 2, 156, 8, 2, 3, 2, 158, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		177, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 185, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 195, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 205, 8, 4, 1, 4, 1, 4, 3, 4, 209,
		8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 222, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 233, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 246, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 251, 8, 7, 10, 7,
		12, 7, 254, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 290, 8, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12,
		305, 8, 12, 1, 13, 1, 13, 3, 13, 309, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		3, 13, 315, 8, 13, 1, 13, 5, 13, 318, 8, 13, 10, 13, 12, 13, 321, 9, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 328, 8, 14, 10, 14, 12, 14, 331,
		9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 339, 8, 15, 1,
		15, 1, 15, 3, 15, 343, 8, 15, 1, 15, 3, 15, 346, 8, 15, 1, 16, 3, 16, 349,
		8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 4, 17, 356, 8, 17, 11, 17, 12,
		17, 357, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 3, 19, 370, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 375, 8, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 20, 3, 20, 382, 8, 20, 1, 20, 1, 20, 1, 20, 3,
		20, 387, 8, 20, 1, 20, 1, 20, 3, 20, 391, 8, 20, 1, 20, 1, 20, 3, 20, 395,
		8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 400, 8, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 3, 22, 406, 8, 22, 1, 22, 3, 22, 409, 8, 22, 1, 22, 1, 22, 1, 22, 3,
		22, 414, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 5, 23, 426, 8, 23, 10, 23, 12, 23, 429, 9, 23, 1, 23, 1,
		23, 3, 23, 433, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 459, 8, 24, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 5, 26, 469, 8, 26, 10,
		26, 12, 26, 472, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 479,
		8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 485, 8, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 498,
		8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 3, 35, 522, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 562, 8, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 5, 36, 582, 8, 36, 10, 36, 12, 36, 585, 9, 36, 1, 37, 1, 37, 1,
		37, 0, 1, 72, 38, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 0, 11, 1, 0, 41, 42, 2, 0, 65, 65, 69, 69, 1, 0, 17,
		18, 1, 0, 31, 32, 1, 0, 19, 21, 2, 0, 18, 18, 22, 22, 1, 0, 23, 24, 1,
		0, 25, 26, 1, 0, 27, 28, 1, 0, 29, 30, 1, 0, 35, 39, 649, 0, 76, 1, 0,
		0, 0, 2, 82, 1, 0, 0, 0, 4, 157, 1, 0, 0, 0, 6, 176, 1, 0, 0, 0, 8, 208,
		1, 0, 0, 0, 10, 221, 1, 0, 0, 0, 12, 245, 1, 0, 0, 0, 14, 247, 1, 0, 0,
		0, 16, 255, 1, 0, 0, 0, 18, 259, 1, 0, 0, 0, 20, 265, 1, 0, 0, 0, 22, 270,
		1, 0, 0, 0, 24, 304, 1, 0, 0, 0, 26, 306, 1, 0, 0, 0, 28, 322, 1, 0, 0,
		0, 30, 334, 1, 0, 0, 0, 32, 348, 1, 0, 0, 0, 34, 352, 1, 0, 0, 0, 36, 359,
		1, 0, 0, 0, 38, 365, 1, 0, 0, 0, 40, 381, 1, 0, 0, 0, 42, 396, 1, 0, 0,
		0, 44, 405, 1, 0, 0, 0, 46, 432, 1, 0, 0, 0, 48, 458, 1, 0, 0, 0, 50, 460,
		1, 0, 0, 0, 52, 470, 1, 0, 0, 0, 54, 473, 1, 0, 0, 0, 56, 480, 1, 0, 0,
		0, 58, 486, 1, 0, 0, 0, 60, 492, 1, 0, 0, 0, 62, 503, 1, 0, 0, 0, 64, 507,
		1, 0, 0, 0, 66, 514, 1, 0, 0, 0, 68, 516, 1, 0, 0, 0, 70, 521, 1, 0, 0,
		0, 72, 561, 1, 0, 0, 0, 74, 586, 1, 0, 0, 0, 76, 77, 3, 2, 1, 0, 77, 78,
		5, 0, 0, 1, 78, 1, 1, 0, 0, 0, 79, 81, 3, 4, 2, 0, 80, 79, 1, 0, 0, 0,
		81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 3, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 85, 87, 3, 6, 3, 0, 86, 88, 5, 59, 0, 0, 87,
		86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 158, 1, 0, 0, 0, 89, 91, 3, 8,
		4, 0, 90, 92, 5, 59, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		158, 1, 0, 0, 0, 93, 95, 3, 10, 5, 0, 94, 96, 5, 59, 0, 0, 95, 94, 1, 0,
		0, 0, 95, 96, 1, 0, 0, 0, 96, 158, 1, 0, 0, 0, 97, 99, 3, 12, 6, 0, 98,
		100, 5, 59, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 158, 1,
		0, 0, 0, 101, 103, 3, 18, 9, 0, 102, 104, 5, 59, 0, 0, 103, 102, 1, 0,
		0, 0, 103, 104, 1, 0, 0, 0, 104, 158, 1, 0, 0, 0, 105, 107, 3, 20, 10,
		0, 106, 108, 5, 59, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108,
		158, 1, 0, 0, 0, 109, 111, 3, 22, 11, 0, 110, 112, 5, 59, 0, 0, 111, 110,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 158, 1, 0, 0, 0, 113, 115, 3, 38,
		19, 0, 114, 116, 5, 59, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 158, 1, 0, 0, 0, 117, 119, 3, 42, 21, 0, 118, 120, 5, 59, 0, 0,
		119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 158, 1, 0, 0, 0, 121,
		123, 3, 46, 23, 0, 122, 124, 5, 59, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124,
		1, 0, 0, 0, 124, 158, 1, 0, 0, 0, 125, 127, 3, 48, 24, 0, 126, 128, 5,
		59, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 158, 1, 0, 0,
		0, 129, 131, 3, 50, 25, 0, 130, 132, 5, 59, 0, 0, 131, 130, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 158, 1, 0, 0, 0, 133, 135, 3, 58, 29, 0, 134,
		136, 5, 59, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 158,
		1, 0, 0, 0, 137, 139, 3, 60, 30, 0, 138, 140, 5, 59, 0, 0, 139, 138, 1,
		0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 158, 1, 0, 0, 0, 141, 143, 3, 64, 32,
		0, 142, 144, 5, 59, 0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		158, 1, 0, 0, 0, 145, 147, 3, 66, 33, 0, 146, 148, 5, 59, 0, 0, 147, 146,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 158, 1, 0, 0, 0, 149, 151, 3, 68,
		34, 0, 150, 152, 5, 59, 0, 0, 151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0,
		0, 152, 158, 1, 0, 0, 0, 153, 155, 3, 70, 35, 0, 154, 156, 5, 59, 0, 0,
		155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157,
		85, 1, 0, 0, 0, 157, 89, 1, 0, 0, 0, 157, 93, 1, 0, 0, 0, 157, 97, 1, 0,
		0, 0, 157, 101, 1, 0, 0, 0, 157, 105, 1, 0, 0, 0, 157, 109, 1, 0, 0, 0,
		157, 113, 1, 0, 0, 0, 157, 117, 1, 0, 0, 0, 157, 121, 1, 0, 0, 0, 157,
		125, 1, 0, 0, 0, 157, 129, 1, 0, 0, 0, 157, 133, 1, 0, 0, 0, 157, 137,
		1, 0, 0, 0, 157, 141, 1, 0, 0, 0, 157, 145, 1, 0, 0, 0, 157, 149, 1, 0,
		0, 0, 157, 153, 1, 0, 0, 0, 158, 5, 1, 0, 0, 0, 159, 160, 5, 41, 0, 0,
		160, 161, 5, 69, 0, 0, 161, 162, 5, 60, 0, 0, 162, 163, 3, 74, 37, 0, 163,
		164, 5, 61, 0, 0, 164, 165, 3, 72, 36, 0, 165, 177, 1, 0, 0, 0, 166, 167,
		5, 41, 0, 0, 167, 168, 5, 69, 0, 0, 168, 169, 5, 61, 0, 0, 169, 177, 3,
		72, 36, 0, 170, 171, 5, 41, 0, 0, 171, 172, 5, 69, 0, 0, 172, 173, 5, 60,
		0, 0, 173, 174, 3, 74, 37, 0, 174, 175, 5, 62, 0, 0, 175, 177, 1, 0, 0,
		0, 176, 159, 1, 0, 0, 0, 176, 166, 1, 0, 0, 0, 176, 170, 1, 0, 0, 0, 177,
		7, 1, 0, 0, 0, 178, 185, 5, 69, 0, 0, 179, 180, 5, 69, 0, 0, 180, 181,
		5, 1, 0, 0, 181, 182, 3, 72, 36, 0, 182, 183, 5, 2, 0, 0, 183, 185, 1,
		0, 0, 0, 184, 178, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 185, 186, 1, 0, 0,
		0, 186, 187, 5, 61, 0, 0, 187, 209, 3, 72, 36, 0, 188, 195, 5, 69, 0, 0,
		189, 190, 5, 69, 0, 0, 190, 191, 5, 1, 0, 0, 191, 192, 3, 72, 36, 0, 192,
		193, 5, 2, 0, 0, 193, 195, 1, 0, 0, 0, 194, 188, 1, 0, 0, 0, 194, 189,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 5, 63, 0, 0, 197, 209, 3, 72,
		36, 0, 198, 205, 5, 69, 0, 0, 199, 200, 5, 69, 0, 0, 200, 201, 5, 1, 0,
		0, 201, 202, 3, 72, 36, 0, 202, 203, 5, 2, 0, 0, 203, 205, 1, 0, 0, 0,
		204, 198, 1, 0, 0, 0, 204, 199, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206,
		207, 5, 64, 0, 0, 207, 209, 3, 72, 36, 0, 208, 184, 1, 0, 0, 0, 208, 194,
		1, 0, 0, 0, 208, 204, 1, 0, 0, 0, 209, 9, 1, 0, 0, 0, 210, 211, 5, 42,
		0, 0, 211, 212, 5, 69, 0, 0, 212, 213, 5, 60, 0, 0, 213, 214, 3, 74, 37,
		0, 214, 215, 5, 61, 0, 0, 215, 216, 3, 72, 36, 0, 216, 222, 1, 0, 0, 0,
		217, 218, 5, 42, 0, 0, 218, 219, 5, 69, 0, 0, 219, 220, 5, 61, 0, 0, 220,
		222, 3, 72, 36, 0, 221, 210, 1, 0, 0, 0, 221, 217, 1, 0, 0, 0, 222, 11,
		1, 0, 0, 0, 223, 224, 5, 41, 0, 0, 224, 225, 5, 69, 0, 0, 225, 226, 5,
		60, 0, 0, 226, 227, 5, 1, 0, 0, 227, 228, 3, 74, 37, 0, 228, 229, 5, 2,
		0, 0, 229, 230, 5, 61, 0, 0, 230, 232, 5, 1, 0, 0, 231, 233, 3, 14, 7,
		0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		235, 5, 2, 0, 0, 235, 246, 1, 0, 0, 0, 236, 237, 5, 41, 0, 0, 237, 238,
		5, 69, 0, 0, 238, 239, 5, 60, 0, 0, 239, 240, 5, 1, 0, 0, 240, 241, 3,
		74, 37, 0, 241, 242, 5, 2, 0, 0, 242, 243, 5, 61, 0, 0, 243, 244, 5, 69,
		0, 0, 244, 246, 1, 0, 0, 0, 245, 223, 1, 0, 0, 0, 245, 236, 1, 0, 0, 0,
		246, 13, 1, 0, 0, 0, 247, 252, 3, 72, 36, 0, 248, 249, 5, 3, 0, 0, 249,
		251, 3, 72, 36, 0, 250, 248, 1, 0, 0, 0, 251, 254, 1, 0, 0, 0, 252, 250,
		1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 15, 1, 0, 0, 0, 254, 252, 1, 0,
		0, 0, 255, 256, 5, 1, 0, 0, 256, 257, 3, 74, 37, 0, 257, 258, 5, 2, 0,
		0, 258, 17, 1, 0, 0, 0, 259, 260, 5, 69, 0, 0, 260, 261, 5, 4, 0, 0, 261,
		262, 5, 5, 0, 0, 262, 263, 3, 72, 36, 0, 263, 264, 5, 6, 0, 0, 264, 19,
		1, 0, 0, 0, 265, 266, 5, 69, 0, 0, 266, 267, 5, 7, 0, 0, 267, 268, 5, 5,
		0, 0, 268, 269, 5, 6, 0, 0, 269, 21, 1, 0, 0, 0, 270, 271, 5, 69, 0, 0,
		271, 272, 5, 8, 0, 0, 272, 273, 5, 5, 0, 0, 273, 274, 5, 9, 0, 0, 274,
		275, 5, 60, 0, 0, 275, 276, 3, 72, 36, 0, 276, 277, 5, 6, 0, 0, 277, 23,
		1, 0, 0, 0, 278, 279, 5, 41, 0, 0, 279, 280, 5, 69, 0, 0, 280, 281, 5,
		60, 0, 0, 281, 282, 5, 1, 0, 0, 282, 283, 5, 1, 0, 0, 283, 284, 3, 74,
		37, 0, 284, 285, 5, 2, 0, 0, 285, 286, 5, 2, 0, 0, 286, 287, 5, 61, 0,
		0, 287, 289, 5, 1, 0, 0, 288, 290, 3, 26, 13, 0, 289, 288, 1, 0, 0, 0,
		289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 5, 2, 0, 0, 292,
		305, 1, 0, 0, 0, 293, 294, 5, 41, 0, 0, 294, 295, 5, 69, 0, 0, 295, 296,
		5, 60, 0, 0, 296, 297, 5, 1, 0, 0, 297, 298, 5, 1, 0, 0, 298, 299, 3, 74,
		37, 0, 299, 300, 5, 2, 0, 0, 300, 301, 5, 2, 0, 0, 301, 302, 5, 61, 0,
		0, 302, 303, 5, 69, 0, 0, 303, 305, 1, 0, 0, 0, 304, 278, 1, 0, 0, 0, 304,
		293, 1, 0, 0, 0, 305, 25, 1, 0, 0, 0, 306, 308, 5, 1, 0, 0, 307, 309, 3,
		14, 7, 0, 308, 307, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0,
		0, 310, 319, 5, 2, 0, 0, 311, 312, 5, 3, 0, 0, 312, 314, 5, 1, 0, 0, 313,
		315, 3, 14, 7, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316,
		1, 0, 0, 0, 316, 318, 5, 2, 0, 0, 317, 311, 1, 0, 0, 0, 318, 321, 1, 0,
		0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 27, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 322, 323, 5, 10, 0, 0, 323, 324, 5, 69, 0, 0, 324,
		329, 5, 11, 0, 0, 325, 328, 3, 30, 15, 0, 326, 328, 3, 32, 16, 0, 327,
		325, 1, 0, 0, 0, 327, 326, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327,
		1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 332, 1, 0, 0, 0, 331, 329, 1, 0,
		0, 0, 332, 333, 5, 12, 0, 0, 333, 29, 1, 0, 0, 0, 334, 335, 7, 0, 0, 0,
		335, 338, 5, 69, 0, 0, 336, 337, 5, 60, 0, 0, 337, 339, 3, 74, 37, 0, 338,
		336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 341,
		5, 61, 0, 0, 341, 343, 3, 72, 36, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1,
		0, 0, 0, 343, 345, 1, 0, 0, 0, 344, 346, 5, 59, 0, 0, 345, 344, 1, 0, 0,
		0, 345, 346, 1, 0, 0, 0, 346, 31, 1, 0, 0, 0, 347, 349, 5, 13, 0, 0, 348,
		347, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 351,
		3, 38, 19, 0, 351, 33, 1, 0, 0, 0, 352, 355, 5, 69, 0, 0, 353, 354, 5,
		14, 0, 0, 354, 356, 5, 69, 0, 0, 355, 353, 1, 0, 0, 0, 356, 357, 1, 0,
		0, 0, 357, 355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 35, 1, 0, 0, 0,
		359, 360, 5, 69, 0, 0, 360, 361, 5, 14, 0, 0, 361, 362, 5, 69, 0, 0, 362,
		363, 5, 61, 0, 0, 363, 364, 3, 72, 36, 0, 364, 37, 1, 0, 0, 0, 365, 366,
		5, 43, 0, 0, 366, 367, 5, 69, 0, 0, 367, 369, 5, 5, 0, 0, 368, 370, 3,
		40, 20, 0, 369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0,
		0, 0, 371, 374, 5, 6, 0, 0, 372, 373, 5, 15, 0, 0, 373, 375, 3, 74, 37,
		0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376,
		377, 5, 11, 0, 0, 377, 378, 3, 2, 1, 0, 378, 379, 5, 12, 0, 0, 379, 39,
		1, 0, 0, 0, 380, 382, 7, 1, 0, 0, 381, 380, 1, 0, 0, 0, 381, 382, 1, 0,
		0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 5, 69, 0, 0, 384, 386, 5, 60, 0,
		0, 385, 387, 5, 44, 0, 0, 386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387,
		390, 1, 0, 0, 0, 388, 391, 3, 74, 37, 0, 389, 391, 3, 16, 8, 0, 390, 388,
		1, 0, 0, 0, 390, 389, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392, 393, 5, 3,
		0, 0, 393, 395, 3, 40, 20, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0,
		0, 395, 41, 1, 0, 0, 0, 396, 397, 5, 69, 0, 0, 397, 399, 5, 5, 0, 0, 398,
		400, 3, 44, 22, 0, 399, 398, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 401,
		1, 0, 0, 0, 401, 402, 5, 6, 0, 0, 402, 43, 1, 0, 0, 0, 403, 404, 5, 69,
		0, 0, 404, 406, 5, 60, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0,
		406, 408, 1, 0, 0, 0, 407, 409, 5, 45, 0, 0, 408, 407, 1, 0, 0, 0, 408,
		409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 413, 3, 72, 36, 0, 411, 412,
		5, 3, 0, 0, 412, 414, 3, 44, 22, 0, 413, 411, 1, 0, 0, 0, 413, 414, 1,
		0, 0, 0, 414, 45, 1, 0, 0, 0, 415, 416, 5, 46, 0, 0, 416, 417, 5, 5, 0,
		0, 417, 418, 3, 72, 36, 0, 418, 419, 5, 6, 0, 0, 419, 433, 1, 0, 0, 0,
		420, 421, 5, 46, 0, 0, 421, 422, 5, 5, 0, 0, 422, 427, 3, 72, 36, 0, 423,
		424, 5, 3, 0, 0, 424, 426, 3, 72, 36, 0, 425, 423, 1, 0, 0, 0, 426, 429,
		1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 430, 1, 0,
		0, 0, 429, 427, 1, 0, 0, 0, 430, 431, 5, 6, 0, 0, 431, 433, 1, 0, 0, 0,
		432, 415, 1, 0, 0, 0, 432, 420, 1, 0, 0, 0, 433, 47, 1, 0, 0, 0, 434, 435,
		5, 47, 0, 0, 435, 436, 3, 72, 36, 0, 436, 437, 5, 11, 0, 0, 437, 438, 3,
		2, 1, 0, 438, 439, 5, 12, 0, 0, 439, 459, 1, 0, 0, 0, 440, 441, 5, 47,
		0, 0, 441, 442, 3, 72, 36, 0, 442, 443, 5, 11, 0, 0, 443, 444, 3, 2, 1,
		0, 444, 445, 5, 12, 0, 0, 445, 446, 5, 48, 0, 0, 446, 447, 5, 11, 0, 0,
		447, 448, 3, 2, 1, 0, 448, 449, 5, 12, 0, 0, 449, 459, 1, 0, 0, 0, 450,
		451, 5, 47, 0, 0, 451, 452, 3, 72, 36, 0, 452, 453, 5, 11, 0, 0, 453, 454,
		3, 2, 1, 0, 454, 455, 5, 12, 0, 0, 455, 456, 5, 48, 0, 0, 456, 457, 3,
		48, 24, 0, 457, 459, 1, 0, 0, 0, 458, 434, 1, 0, 0, 0, 458, 440, 1, 0,
		0, 0, 458, 450, 1, 0, 0, 0, 459, 49, 1, 0, 0, 0, 460, 461, 5, 49, 0, 0,
		461, 462, 3, 72, 36, 0, 462, 463, 5, 11, 0, 0, 463, 464, 3, 52, 26, 0,
		464, 465, 5, 12, 0, 0, 465, 51, 1, 0, 0, 0, 466, 469, 3, 54, 27, 0, 467,
		469, 3, 56, 28, 0, 468, 466, 1, 0, 0, 0, 468, 467, 1, 0, 0, 0, 469, 472,
		1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 53, 1, 0,
		0, 0, 472, 470, 1, 0, 0, 0, 473, 474, 5, 50, 0, 0, 474, 475, 3, 72, 36,
		0, 475, 476, 5, 60, 0, 0, 476, 478, 3, 2, 1, 0, 477, 479, 5, 57, 0, 0,
		478, 477, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 55, 1, 0, 0, 0, 480, 481,
		5, 51, 0, 0, 481, 482, 5, 60, 0, 0, 482, 484, 3, 2, 1, 0, 483, 485, 5,
		57, 0, 0, 484, 483, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 57, 1, 0, 0,
		0, 486, 487, 5, 52, 0, 0, 487, 488, 3, 72, 36, 0, 488, 489, 5, 11, 0, 0,
		489, 490, 3, 2, 1, 0, 490, 491, 5, 12, 0, 0, 491, 59, 1, 0, 0, 0, 492,
		493, 5, 53, 0, 0, 493, 494, 7, 1, 0, 0, 494, 497, 5, 54, 0, 0, 495, 498,
		3, 72, 36, 0, 496, 498, 3, 62, 31, 0, 497, 495, 1, 0, 0, 0, 497, 496, 1,
		0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 500, 5, 11, 0, 0, 500, 501, 3, 2, 1,
		0, 501, 502, 5, 12, 0, 0, 502, 61, 1, 0, 0, 0, 503, 504, 3, 72, 36, 0,
		504, 505, 5, 16, 0, 0, 505, 506, 3, 72, 36, 0, 506, 63, 1, 0, 0, 0, 507,
		508, 5, 55, 0, 0, 508, 509, 3, 72, 36, 0, 509, 510, 5, 48, 0, 0, 510, 511,
		5, 11, 0, 0, 511, 512, 3, 2, 1, 0, 512, 513, 5, 12, 0, 0, 513, 65, 1, 0,
		0, 0, 514, 515, 5, 57, 0, 0, 515, 67, 1, 0, 0, 0, 516, 517, 5, 56, 0, 0,
		517, 69, 1, 0, 0, 0, 518, 522, 5, 58, 0, 0, 519, 520, 5, 58, 0, 0, 520,
		522, 3, 72, 36, 0, 521, 518, 1, 0, 0, 0, 521, 519, 1, 0, 0, 0, 522, 71,
		1, 0, 0, 0, 523, 524, 6, 36, -1, 0, 524, 525, 5, 5, 0, 0, 525, 526, 3,
		72, 36, 0, 526, 527, 5, 6, 0, 0, 527, 562, 1, 0, 0, 0, 528, 529, 7, 2,
		0, 0, 529, 562, 3, 72, 36, 20, 530, 562, 7, 3, 0, 0, 531, 562, 5, 40, 0,
		0, 532, 562, 5, 66, 0, 0, 533, 562, 5, 67, 0, 0, 534, 562, 5, 68, 0, 0,
		535, 562, 3, 42, 21, 0, 536, 537, 5, 69, 0, 0, 537, 562, 5, 33, 0, 0, 538,
		539, 5, 69, 0, 0, 539, 562, 5, 34, 0, 0, 540, 541, 5, 69, 0, 0, 541, 542,
		5, 1, 0, 0, 542, 543, 3, 72, 36, 0, 543, 544, 5, 2, 0, 0, 544, 562, 1,
		0, 0, 0, 545, 562, 5, 69, 0, 0, 546, 547, 5, 35, 0, 0, 547, 548, 5, 5,
		0, 0, 548, 549, 3, 72, 36, 0, 549, 550, 5, 6, 0, 0, 550, 562, 1, 0, 0,
		0, 551, 552, 5, 36, 0, 0, 552, 553, 5, 5, 0, 0, 553, 554, 3, 72, 36, 0,
		554, 555, 5, 6, 0, 0, 555, 562, 1, 0, 0, 0, 556, 557, 5, 37, 0, 0, 557,
		558, 5, 5, 0, 0, 558, 559, 3, 72, 36, 0, 559, 560, 5, 6, 0, 0, 560, 562,
		1, 0, 0, 0, 561, 523, 1, 0, 0, 0, 561, 528, 1, 0, 0, 0, 561, 530, 1, 0,
		0, 0, 561, 531, 1, 0, 0, 0, 561, 532, 1, 0, 0, 0, 561, 533, 1, 0, 0, 0,
		561, 534, 1, 0, 0, 0, 561, 535, 1, 0, 0, 0, 561, 536, 1, 0, 0, 0, 561,
		538, 1, 0, 0, 0, 561, 540, 1, 0, 0, 0, 561, 545, 1, 0, 0, 0, 561, 546,
		1, 0, 0, 0, 561, 551, 1, 0, 0, 0, 561, 556, 1, 0, 0, 0, 562, 583, 1, 0,
		0, 0, 563, 564, 10, 19, 0, 0, 564, 565, 7, 4, 0, 0, 565, 582, 3, 72, 36,
		20, 566, 567, 10, 18, 0, 0, 567, 568, 7, 5, 0, 0, 568, 582, 3, 72, 36,
		19, 569, 570, 10, 17, 0, 0, 570, 571, 7, 6, 0, 0, 571, 582, 3, 72, 36,
		18, 572, 573, 10, 16, 0, 0, 573, 574, 7, 7, 0, 0, 574, 582, 3, 72, 36,
		17, 575, 576, 10, 15, 0, 0, 576, 577, 7, 8, 0, 0, 577, 582, 3, 72, 36,
		16, 578, 579, 10, 14, 0, 0, 579, 580, 7, 9, 0, 0, 580, 582, 3, 72, 36,
		15, 581, 563, 1, 0, 0, 0, 581, 566, 1, 0, 0, 0, 581, 569, 1, 0, 0, 0, 581,
		572, 1, 0, 0, 0, 581, 575, 1, 0, 0, 0, 581, 578, 1, 0, 0, 0, 582, 585,
		1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584, 73, 1, 0,
		0, 0, 585, 583, 1, 0, 0, 0, 586, 587, 7, 10, 0, 0, 587, 75, 1, 0, 0, 0,
		63, 82, 87, 91, 95, 99, 103, 107, 111, 115, 119, 123, 127, 131, 135, 139,
		143, 147, 151, 155, 157, 176, 184, 194, 204, 208, 221, 232, 245, 252, 289,
		304, 308, 314, 319, 327, 329, 338, 342, 345, 348, 357, 369, 374, 381, 386,
		390, 394, 399, 405, 408, 413, 427, 432, 458, 468, 470, 478, 484, 497, 521,
		561, 581, 583,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF           = antlr.TokenEOF
	GrammarParserT__0          = 1
	GrammarParserT__1          = 2
	GrammarParserT__2          = 3
	GrammarParserT__3          = 4
	GrammarParserT__4          = 5
	GrammarParserT__5          = 6
	GrammarParserT__6          = 7
	GrammarParserT__7          = 8
	GrammarParserT__8          = 9
	GrammarParserT__9          = 10
	GrammarParserT__10         = 11
	GrammarParserT__11         = 12
	GrammarParserT__12         = 13
	GrammarParserT__13         = 14
	GrammarParserT__14         = 15
	GrammarParserT__15         = 16
	GrammarParserT__16         = 17
	GrammarParserT__17         = 18
	GrammarParserT__18         = 19
	GrammarParserT__19         = 20
	GrammarParserT__20         = 21
	GrammarParserT__21         = 22
	GrammarParserT__22         = 23
	GrammarParserT__23         = 24
	GrammarParserT__24         = 25
	GrammarParserT__25         = 26
	GrammarParserT__26         = 27
	GrammarParserT__27         = 28
	GrammarParserT__28         = 29
	GrammarParserT__29         = 30
	GrammarParserT__30         = 31
	GrammarParserT__31         = 32
	GrammarParserT__32         = 33
	GrammarParserT__33         = 34
	GrammarParserINT           = 35
	GrammarParserFLOAT         = 36
	GrammarParserSTRING        = 37
	GrammarParserBOOL          = 38
	GrammarParserCHARACTER     = 39
	GrammarParserNIL           = 40
	GrammarParserVAR           = 41
	GrammarParserLET           = 42
	GrammarParserFUNC          = 43
	GrammarParserINOUT         = 44
	GrammarParserREFERENCE     = 45
	GrammarParserPRINT         = 46
	GrammarParserIF            = 47
	GrammarParserELSE          = 48
	GrammarParserSWITCH        = 49
	GrammarParserCASE          = 50
	GrammarParserDEFAULT       = 51
	GrammarParserWHILE         = 52
	GrammarParserFOR           = 53
	GrammarParserIN            = 54
	GrammarParserGUARD         = 55
	GrammarParserCONTINUE      = 56
	GrammarParserBREAK         = 57
	GrammarParserRETURN        = 58
	GrammarParserSEMICOLON     = 59
	GrammarParserCOLON         = 60
	GrammarParserEQUAL         = 61
	GrammarParserQUESTION_MARK = 62
	GrammarParserINCREMENT     = 63
	GrammarParserDECREMENT     = 64
	GrammarParserUNDERSCORE    = 65
	GrammarParserDIGIT         = 66
	GrammarParserSTR           = 67
	GrammarParserCHAR          = 68
	GrammarParserID            = 69
	GrammarParserWHITESPACE    = 70
	GrammarParserMULTI_COMMENT = 71
	GrammarParserLINE_COMMENT  = 72
)

// GrammarParser rules.
const (
	GrammarParserRULE_start                   = 0
	GrammarParserRULE_block                   = 1
	GrammarParserRULE_stmts                   = 2
	GrammarParserRULE_varDeclaration          = 3
	GrammarParserRULE_varAssignment           = 4
	GrammarParserRULE_letDeclaration          = 5
	GrammarParserRULE_vectorDeclaration       = 6
	GrammarParserRULE_valuesVectorDeclaration = 7
	GrammarParserRULE_vectorType              = 8
	GrammarParserRULE_vectorAppend            = 9
	GrammarParserRULE_vectorRemoveLast        = 10
	GrammarParserRULE_vectorRemoveAt          = 11
	GrammarParserRULE_matrixDeclaration       = 12
	GrammarParserRULE_valuesMatrixDeclaration = 13
	GrammarParserRULE_structDeclaration       = 14
	GrammarParserRULE_structAttribute         = 15
	GrammarParserRULE_structFunction          = 16
	GrammarParserRULE_structAccess            = 17
	GrammarParserRULE_structAssignment        = 18
	GrammarParserRULE_funcDeclaration         = 19
	GrammarParserRULE_paramsFuncDeclaration   = 20
	GrammarParserRULE_funcCall                = 21
	GrammarParserRULE_paramsFuncCall          = 22
	GrammarParserRULE_printStmt               = 23
	GrammarParserRULE_ifStmt                  = 24
	GrammarParserRULE_switchStmt              = 25
	GrammarParserRULE_cases                   = 26
	GrammarParserRULE_caseStmt                = 27
	GrammarParserRULE_defaultStmt             = 28
	GrammarParserRULE_whileStmt               = 29
	GrammarParserRULE_forStmt                 = 30
	GrammarParserRULE_range                   = 31
	GrammarParserRULE_guardStmt               = 32
	GrammarParserRULE_breakStmt               = 33
	GrammarParserRULE_continueStmt            = 34
	GrammarParserRULE_returnStmt              = 35
	GrammarParserRULE_expr                    = 36
	GrammarParserRULE_type                    = 37
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Block()
	}
	{
		p.SetState(77)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmts() []IStmtsContext
	Stmts(i int) IStmtsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmts() []IStmtsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtsContext); ok {
			len++
		}
	}

	tst := make([]IStmtsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtsContext); ok {
			tst[i] = t.(IStmtsContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmts(i int) IStmtsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_block)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(79)
				p.Stmts()
			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtsContext is an interface to support dynamic dispatch.
type IStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDeclaration() IVarDeclarationContext
	SEMICOLON() antlr.TerminalNode
	VarAssignment() IVarAssignmentContext
	LetDeclaration() ILetDeclarationContext
	VectorDeclaration() IVectorDeclarationContext
	VectorAppend() IVectorAppendContext
	VectorRemoveLast() IVectorRemoveLastContext
	VectorRemoveAt() IVectorRemoveAtContext
	FuncDeclaration() IFuncDeclarationContext
	FuncCall() IFuncCallContext
	PrintStmt() IPrintStmtContext
	IfStmt() IIfStmtContext
	SwitchStmt() ISwitchStmtContext
	WhileStmt() IWhileStmtContext
	ForStmt() IForStmtContext
	GuardStmt() IGuardStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	ReturnStmt() IReturnStmtContext

	// IsStmtsContext differentiates from other interfaces.
	IsStmtsContext()
}

type StmtsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtsContext() *StmtsContext {
	var p = new(StmtsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmts
	return p
}

func InitEmptyStmtsContext(p *StmtsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmts
}

func (*StmtsContext) IsStmtsContext() {}

func NewStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtsContext {
	var p = new(StmtsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_stmts

	return p
}

func (s *StmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtsContext) VarDeclaration() IVarDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *StmtsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *StmtsContext) VarAssignment() IVarAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAssignmentContext)
}

func (s *StmtsContext) LetDeclaration() ILetDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetDeclarationContext)
}

func (s *StmtsContext) VectorDeclaration() IVectorDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorDeclarationContext)
}

func (s *StmtsContext) VectorAppend() IVectorAppendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorAppendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorAppendContext)
}

func (s *StmtsContext) VectorRemoveLast() IVectorRemoveLastContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorRemoveLastContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorRemoveLastContext)
}

func (s *StmtsContext) VectorRemoveAt() IVectorRemoveAtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorRemoveAtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorRemoveAtContext)
}

func (s *StmtsContext) FuncDeclaration() IFuncDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationContext)
}

func (s *StmtsContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *StmtsContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StmtsContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtsContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
}

func (s *StmtsContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtsContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StmtsContext) GuardStmt() IGuardStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardStmtContext)
}

func (s *StmtsContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StmtsContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StmtsContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Stmts() (localctx IStmtsContext) {
	localctx = NewStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_stmts)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.VarDeclaration()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(86)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.VarAssignment()
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(90)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.LetDeclaration()
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(94)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.VectorDeclaration()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(98)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(101)
			p.VectorAppend()
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(102)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.VectorRemoveLast()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(106)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.VectorRemoveAt()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(110)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(113)
			p.FuncDeclaration()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(114)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(117)
			p.FuncCall()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(118)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(121)
			p.PrintStmt()
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(122)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(125)
			p.IfStmt()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(126)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(129)
			p.SwitchStmt()
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(130)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(133)
			p.WhileStmt()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(134)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(137)
			p.ForStmt()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(138)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(141)
			p.GuardStmt()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(142)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(145)
			p.BreakStmt()
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(146)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(149)
			p.ContinueStmt()
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(150)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(153)
			p.ReturnStmt()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(154)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) CopyAll(ctx *VarDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeVarDeclarationContext struct {
	VarDeclarationContext
}

func NewTypeVarDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVarDeclarationContext {
	var p = new(TypeVarDeclarationContext)

	InitEmptyVarDeclarationContext(&p.VarDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclarationContext))

	return p
}

func (s *TypeVarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVarDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *TypeVarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TypeVarDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *TypeVarDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeVarDeclarationContext) QUESTION_MARK() antlr.TerminalNode {
	return s.GetToken(GrammarParserQUESTION_MARK, 0)
}

func (s *TypeVarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTypeVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeValueVarDeclarationContext struct {
	VarDeclarationContext
}

func NewTypeValueVarDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeValueVarDeclarationContext {
	var p = new(TypeValueVarDeclarationContext)

	InitEmptyVarDeclarationContext(&p.VarDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclarationContext))

	return p
}

func (s *TypeValueVarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeValueVarDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *TypeValueVarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TypeValueVarDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *TypeValueVarDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeValueVarDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *TypeValueVarDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeValueVarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTypeValueVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueVarDeclarationContext struct {
	VarDeclarationContext
}

func NewValueVarDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueVarDeclarationContext {
	var p = new(ValueVarDeclarationContext)

	InitEmptyVarDeclarationContext(&p.VarDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarDeclarationContext))

	return p
}

func (s *ValueVarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueVarDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *ValueVarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ValueVarDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *ValueVarDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueVarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValueVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_varDeclaration)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Type_()
		}
		{
			p.SetState(163)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.expr(0)
		}

	case 2:
		localctx = NewValueVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.expr(0)
		}

	case 3:
		localctx = NewTypeVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(170)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Type_()
		}
		{
			p.SetState(174)
			p.Match(GrammarParserQUESTION_MARK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarAssignmentContext is an interface to support dynamic dispatch.
type IVarAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarAssignmentContext differentiates from other interfaces.
	IsVarAssignmentContext()
}

type VarAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignmentContext() *VarAssignmentContext {
	var p = new(VarAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varAssignment
	return p
}

func InitEmptyVarAssignmentContext(p *VarAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_varAssignment
}

func (*VarAssignmentContext) IsVarAssignmentContext() {}

func NewVarAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignmentContext {
	var p = new(VarAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_varAssignment

	return p
}

func (s *VarAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignmentContext) CopyAll(ctx *VarAssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecrementVarAssignmentContext struct {
	VarAssignmentContext
}

func NewDecrementVarAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrementVarAssignmentContext {
	var p = new(DecrementVarAssignmentContext)

	InitEmptyVarAssignmentContext(&p.VarAssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAssignmentContext))

	return p
}

func (s *DecrementVarAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementVarAssignmentContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDECREMENT, 0)
}

func (s *DecrementVarAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *DecrementVarAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DecrementVarAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DecrementVarAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDecrementVarAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueVarAssignmentContext struct {
	VarAssignmentContext
}

func NewValueVarAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueVarAssignmentContext {
	var p = new(ValueVarAssignmentContext)

	InitEmptyVarAssignmentContext(&p.VarAssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAssignmentContext))

	return p
}

func (s *ValueVarAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueVarAssignmentContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *ValueVarAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ValueVarAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueVarAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ValueVarAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValueVarAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementVarAssignmentContext struct {
	VarAssignmentContext
}

func NewIncrementVarAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementVarAssignmentContext {
	var p = new(IncrementVarAssignmentContext)

	InitEmptyVarAssignmentContext(&p.VarAssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarAssignmentContext))

	return p
}

func (s *IncrementVarAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementVarAssignmentContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINCREMENT, 0)
}

func (s *IncrementVarAssignmentContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IncrementVarAssignmentContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncrementVarAssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IncrementVarAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIncrementVarAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VarAssignment() (localctx IVarAssignmentContext) {
	localctx = NewVarAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_varAssignment)
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(178)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(179)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(180)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(181)
				p.expr(0)
			}
			{
				p.SetState(182)
				p.Match(GrammarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(186)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.expr(0)
		}

	case 2:
		localctx = NewIncrementVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(188)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(189)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(190)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(191)
				p.expr(0)
			}
			{
				p.SetState(192)
				p.Match(GrammarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(196)
			p.Match(GrammarParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.expr(0)
		}

	case 3:
		localctx = NewDecrementVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(198)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(199)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(201)
				p.expr(0)
			}
			{
				p.SetState(202)
				p.Match(GrammarParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(206)
			p.Match(GrammarParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetDeclarationContext is an interface to support dynamic dispatch.
type ILetDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLetDeclarationContext differentiates from other interfaces.
	IsLetDeclarationContext()
}

type LetDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetDeclarationContext() *LetDeclarationContext {
	var p = new(LetDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_letDeclaration
	return p
}

func InitEmptyLetDeclarationContext(p *LetDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_letDeclaration
}

func (*LetDeclarationContext) IsLetDeclarationContext() {}

func NewLetDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetDeclarationContext {
	var p = new(LetDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_letDeclaration

	return p
}

func (s *LetDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LetDeclarationContext) CopyAll(ctx *LetDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LetDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueLetDeclarationContext struct {
	LetDeclarationContext
}

func NewValueLetDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueLetDeclarationContext {
	var p = new(ValueLetDeclarationContext)

	InitEmptyLetDeclarationContext(&p.LetDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*LetDeclarationContext))

	return p
}

func (s *ValueLetDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueLetDeclarationContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *ValueLetDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ValueLetDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *ValueLetDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueLetDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValueLetDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeValueLetDeclarationContext struct {
	LetDeclarationContext
}

func NewTypeValueLetDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeValueLetDeclarationContext {
	var p = new(TypeValueLetDeclarationContext)

	InitEmptyLetDeclarationContext(&p.LetDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*LetDeclarationContext))

	return p
}

func (s *TypeValueLetDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeValueLetDeclarationContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *TypeValueLetDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *TypeValueLetDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *TypeValueLetDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeValueLetDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *TypeValueLetDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TypeValueLetDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitTypeValueLetDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) LetDeclaration() (localctx ILetDeclarationContext) {
	localctx = NewLetDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_letDeclaration)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueLetDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Type_()
		}
		{
			p.SetState(214)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.expr(0)
		}

	case 2:
		localctx = NewValueLetDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorDeclarationContext is an interface to support dynamic dispatch.
type IVectorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVectorDeclarationContext differentiates from other interfaces.
	IsVectorDeclarationContext()
}

type VectorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorDeclarationContext() *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorDeclaration
	return p
}

func InitEmptyVectorDeclarationContext(p *VectorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorDeclaration
}

func (*VectorDeclarationContext) IsVectorDeclarationContext() {}

func NewVectorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorDeclarationContext {
	var p = new(VectorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vectorDeclaration

	return p
}

func (s *VectorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorDeclarationContext) CopyAll(ctx *VectorDeclarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VectorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorIdDeclarationContext struct {
	VectorDeclarationContext
}

func NewVectorIdDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorIdDeclarationContext {
	var p = new(VectorIdDeclarationContext)

	InitEmptyVectorDeclarationContext(&p.VectorDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDeclarationContext))

	return p
}

func (s *VectorIdDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorIdDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *VectorIdDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *VectorIdDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *VectorIdDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *VectorIdDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorIdDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *VectorIdDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorIdDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorValuesDeclarationContext struct {
	VectorDeclarationContext
}

func NewVectorValuesDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorValuesDeclarationContext {
	var p = new(VectorValuesDeclarationContext)

	InitEmptyVectorDeclarationContext(&p.VectorDeclarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorDeclarationContext))

	return p
}

func (s *VectorValuesDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorValuesDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *VectorValuesDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *VectorValuesDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *VectorValuesDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorValuesDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *VectorValuesDeclarationContext) ValuesVectorDeclaration() IValuesVectorDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesVectorDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesVectorDeclarationContext)
}

func (s *VectorValuesDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorValuesDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VectorDeclaration() (localctx IVectorDeclarationContext) {
	localctx = NewVectorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_vectorDeclaration)
	var _la int

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorValuesDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Type_()
		}
		{
			p.SetState(228)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1346472640544) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
			{
				p.SetState(231)
				p.ValuesVectorDeclaration()
			}

		}
		{
			p.SetState(234)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVectorIdDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Type_()
		}
		{
			p.SetState(241)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuesVectorDeclarationContext is an interface to support dynamic dispatch.
type IValuesVectorDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsValuesVectorDeclarationContext differentiates from other interfaces.
	IsValuesVectorDeclarationContext()
}

type ValuesVectorDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesVectorDeclarationContext() *ValuesVectorDeclarationContext {
	var p = new(ValuesVectorDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_valuesVectorDeclaration
	return p
}

func InitEmptyValuesVectorDeclarationContext(p *ValuesVectorDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_valuesVectorDeclaration
}

func (*ValuesVectorDeclarationContext) IsValuesVectorDeclarationContext() {}

func NewValuesVectorDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesVectorDeclarationContext {
	var p = new(ValuesVectorDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_valuesVectorDeclaration

	return p
}

func (s *ValuesVectorDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesVectorDeclarationContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ValuesVectorDeclarationContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValuesVectorDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesVectorDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesVectorDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValuesVectorDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ValuesVectorDeclaration() (localctx IValuesVectorDeclarationContext) {
	localctx = NewValuesVectorDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_valuesVectorDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.expr(0)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserT__2 {
		{
			p.SetState(248)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.expr(0)
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorTypeContext is an interface to support dynamic dispatch.
type IVectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext

	// IsVectorTypeContext differentiates from other interfaces.
	IsVectorTypeContext()
}

type VectorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorTypeContext() *VectorTypeContext {
	var p = new(VectorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorType
	return p
}

func InitEmptyVectorTypeContext(p *VectorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorType
}

func (*VectorTypeContext) IsVectorTypeContext() {}

func NewVectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorTypeContext {
	var p = new(VectorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vectorType

	return p
}

func (s *VectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VectorType() (localctx IVectorTypeContext) {
	localctx = NewVectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_vectorType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(GrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Type_()
	}
	{
		p.SetState(257)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorAppendContext is an interface to support dynamic dispatch.
type IVectorAppendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsVectorAppendContext differentiates from other interfaces.
	IsVectorAppendContext()
}

type VectorAppendContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorAppendContext() *VectorAppendContext {
	var p = new(VectorAppendContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorAppend
	return p
}

func InitEmptyVectorAppendContext(p *VectorAppendContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorAppend
}

func (*VectorAppendContext) IsVectorAppendContext() {}

func NewVectorAppendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorAppendContext {
	var p = new(VectorAppendContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vectorAppend

	return p
}

func (s *VectorAppendContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorAppendContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *VectorAppendContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAppendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VectorAppend() (localctx IVectorAppendContext) {
	localctx = NewVectorAppendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_vectorAppend)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(GrammarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.expr(0)
	}
	{
		p.SetState(263)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorRemoveLastContext is an interface to support dynamic dispatch.
type IVectorRemoveLastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVectorRemoveLastContext differentiates from other interfaces.
	IsVectorRemoveLastContext()
}

type VectorRemoveLastContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorRemoveLastContext() *VectorRemoveLastContext {
	var p = new(VectorRemoveLastContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorRemoveLast
	return p
}

func InitEmptyVectorRemoveLastContext(p *VectorRemoveLastContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorRemoveLast
}

func (*VectorRemoveLastContext) IsVectorRemoveLastContext() {}

func NewVectorRemoveLastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorRemoveLastContext {
	var p = new(VectorRemoveLastContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vectorRemoveLast

	return p
}

func (s *VectorRemoveLastContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorRemoveLastContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *VectorRemoveLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorRemoveLastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorRemoveLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorRemoveLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VectorRemoveLast() (localctx IVectorRemoveLastContext) {
	localctx = NewVectorRemoveLastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_vectorRemoveLast)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(GrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorRemoveAtContext is an interface to support dynamic dispatch.
type IVectorRemoveAtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsVectorRemoveAtContext differentiates from other interfaces.
	IsVectorRemoveAtContext()
}

type VectorRemoveAtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorRemoveAtContext() *VectorRemoveAtContext {
	var p = new(VectorRemoveAtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorRemoveAt
	return p
}

func InitEmptyVectorRemoveAtContext(p *VectorRemoveAtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vectorRemoveAt
}

func (*VectorRemoveAtContext) IsVectorRemoveAtContext() {}

func NewVectorRemoveAtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorRemoveAtContext {
	var p = new(VectorRemoveAtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vectorRemoveAt

	return p
}

func (s *VectorRemoveAtContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorRemoveAtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *VectorRemoveAtContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *VectorRemoveAtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorRemoveAtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorRemoveAtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorRemoveAtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorRemoveAt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) VectorRemoveAt() (localctx IVectorRemoveAtContext) {
	localctx = NewVectorRemoveAtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_vectorRemoveAt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(GrammarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(GrammarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.expr(0)
	}
	{
		p.SetState(276)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrixDeclarationContext is an interface to support dynamic dispatch.
type IMatrixDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EQUAL() antlr.TerminalNode
	ValuesMatrixDeclaration() IValuesMatrixDeclarationContext

	// IsMatrixDeclarationContext differentiates from other interfaces.
	IsMatrixDeclarationContext()
}

type MatrixDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixDeclarationContext() *MatrixDeclarationContext {
	var p = new(MatrixDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_matrixDeclaration
	return p
}

func InitEmptyMatrixDeclarationContext(p *MatrixDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_matrixDeclaration
}

func (*MatrixDeclarationContext) IsMatrixDeclarationContext() {}

func NewMatrixDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixDeclarationContext {
	var p = new(MatrixDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_matrixDeclaration

	return p
}

func (s *MatrixDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *MatrixDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *MatrixDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *MatrixDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *MatrixDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MatrixDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *MatrixDeclarationContext) ValuesMatrixDeclaration() IValuesMatrixDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesMatrixDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesMatrixDeclarationContext)
}

func (s *MatrixDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMatrixDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) MatrixDeclaration() (localctx IMatrixDeclarationContext) {
	localctx = NewMatrixDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_matrixDeclaration)
	var _la int

	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Type_()
		}
		{
			p.SetState(284)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserT__0 {
			{
				p.SetState(288)
				p.ValuesMatrixDeclaration()
			}

		}
		{
			p.SetState(291)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Type_()
		}
		{
			p.SetState(299)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuesMatrixDeclarationContext is an interface to support dynamic dispatch.
type IValuesMatrixDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValuesVectorDeclaration() []IValuesVectorDeclarationContext
	ValuesVectorDeclaration(i int) IValuesVectorDeclarationContext

	// IsValuesMatrixDeclarationContext differentiates from other interfaces.
	IsValuesMatrixDeclarationContext()
}

type ValuesMatrixDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesMatrixDeclarationContext() *ValuesMatrixDeclarationContext {
	var p = new(ValuesMatrixDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_valuesMatrixDeclaration
	return p
}

func InitEmptyValuesMatrixDeclarationContext(p *ValuesMatrixDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_valuesMatrixDeclaration
}

func (*ValuesMatrixDeclarationContext) IsValuesMatrixDeclarationContext() {}

func NewValuesMatrixDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesMatrixDeclarationContext {
	var p = new(ValuesMatrixDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_valuesMatrixDeclaration

	return p
}

func (s *ValuesMatrixDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesMatrixDeclarationContext) AllValuesVectorDeclaration() []IValuesVectorDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValuesVectorDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IValuesVectorDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValuesVectorDeclarationContext); ok {
			tst[i] = t.(IValuesVectorDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ValuesMatrixDeclarationContext) ValuesVectorDeclaration(i int) IValuesVectorDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesVectorDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesVectorDeclarationContext)
}

func (s *ValuesMatrixDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesMatrixDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesMatrixDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValuesMatrixDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ValuesMatrixDeclaration() (localctx IValuesMatrixDeclarationContext) {
	localctx = NewValuesMatrixDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_valuesMatrixDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(GrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1346472640544) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
		{
			p.SetState(307)
			p.ValuesVectorDeclaration()
		}

	}
	{
		p.SetState(310)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserT__2 {
		{
			p.SetState(311)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1346472640544) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
			{
				p.SetState(313)
				p.ValuesVectorDeclaration()
			}

		}
		{
			p.SetState(316)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclarationContext is an interface to support dynamic dispatch.
type IStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllStructAttribute() []IStructAttributeContext
	StructAttribute(i int) IStructAttributeContext
	AllStructFunction() []IStructFunctionContext
	StructFunction(i int) IStructFunctionContext

	// IsStructDeclarationContext differentiates from other interfaces.
	IsStructDeclarationContext()
}

type StructDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationContext() *StructDeclarationContext {
	var p = new(StructDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structDeclaration
	return p
}

func InitEmptyStructDeclarationContext(p *StructDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structDeclaration
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *StructDeclarationContext) AllStructAttribute() []IStructAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructAttributeContext); ok {
			len++
		}
	}

	tst := make([]IStructAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructAttributeContext); ok {
			tst[i] = t.(IStructAttributeContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationContext) StructAttribute(i int) IStructAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructAttributeContext)
}

func (s *StructDeclarationContext) AllStructFunction() []IStructFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructFunctionContext); ok {
			len++
		}
	}

	tst := make([]IStructFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructFunctionContext); ok {
			tst[i] = t.(IStructFunctionContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationContext) StructFunction(i int) IStructFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructFunctionContext)
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_structDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(GrammarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15393162797056) != 0 {
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserVAR, GrammarParserLET:
			{
				p.SetState(325)
				p.StructAttribute()
			}

		case GrammarParserT__12, GrammarParserFUNC:
			{
				p.SetState(326)
				p.StructFunction()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(332)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructAttributeContext is an interface to support dynamic dispatch.
type IStructAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	EQUAL() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsStructAttributeContext differentiates from other interfaces.
	IsStructAttributeContext()
}

type StructAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructAttributeContext() *StructAttributeContext {
	var p = new(StructAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAttribute
	return p
}

func InitEmptyStructAttributeContext(p *StructAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAttribute
}

func (*StructAttributeContext) IsStructAttributeContext() {}

func NewStructAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAttributeContext {
	var p = new(StructAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structAttribute

	return p
}

func (s *StructAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAttributeContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *StructAttributeContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *StructAttributeContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *StructAttributeContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *StructAttributeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructAttributeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *StructAttributeContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructAttributeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserSEMICOLON, 0)
}

func (s *StructAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructAttribute() (localctx IStructAttributeContext) {
	localctx = NewStructAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_structAttribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserVAR || _la == GrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(335)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserCOLON {
		{
			p.SetState(336)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Type_()
		}

	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserEQUAL {
		{
			p.SetState(340)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.expr(0)
		}

	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserSEMICOLON {
		{
			p.SetState(344)
			p.Match(GrammarParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructFunctionContext is an interface to support dynamic dispatch.
type IStructFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncDeclaration() IFuncDeclarationContext

	// IsStructFunctionContext differentiates from other interfaces.
	IsStructFunctionContext()
}

type StructFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructFunctionContext() *StructFunctionContext {
	var p = new(StructFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structFunction
	return p
}

func InitEmptyStructFunctionContext(p *StructFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structFunction
}

func (*StructFunctionContext) IsStructFunctionContext() {}

func NewStructFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructFunctionContext {
	var p = new(StructFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structFunction

	return p
}

func (s *StructFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructFunctionContext) FuncDeclaration() IFuncDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclarationContext)
}

func (s *StructFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructFunction() (localctx IStructFunctionContext) {
	localctx = NewStructFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_structFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__12 {
		{
			p.SetState(347)
			p.Match(GrammarParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(350)
		p.FuncDeclaration()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructAccessContext is an interface to support dynamic dispatch.
type IStructAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsStructAccessContext differentiates from other interfaces.
	IsStructAccessContext()
}

type StructAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructAccessContext() *StructAccessContext {
	var p = new(StructAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAccess
	return p
}

func InitEmptyStructAccessContext(p *StructAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAccess
}

func (*StructAccessContext) IsStructAccessContext() {}

func NewStructAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAccessContext {
	var p = new(StructAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structAccess

	return p
}

func (s *StructAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAccessContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *StructAccessContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *StructAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructAccess() (localctx IStructAccessContext) {
	localctx = NewStructAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_structAccess)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserT__13 {
		{
			p.SetState(353)
			p.Match(GrammarParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructAssignmentContext is an interface to support dynamic dispatch.
type IStructAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	Expr() IExprContext

	// IsStructAssignmentContext differentiates from other interfaces.
	IsStructAssignmentContext()
}

type StructAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructAssignmentContext() *StructAssignmentContext {
	var p = new(StructAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAssignment
	return p
}

func InitEmptyStructAssignmentContext(p *StructAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAssignment
}

func (*StructAssignmentContext) IsStructAssignmentContext() {}

func NewStructAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAssignmentContext {
	var p = new(StructAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structAssignment

	return p
}

func (s *StructAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAssignmentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *StructAssignmentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *StructAssignmentContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserEQUAL, 0)
}

func (s *StructAssignmentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructAssignment() (localctx IStructAssignmentContext) {
	localctx = NewStructAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_structAssignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(GrammarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.Match(GrammarParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclarationContext is an interface to support dynamic dispatch.
type IFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	Block() IBlockContext
	ParamsFuncDeclaration() IParamsFuncDeclarationContext
	Type_() ITypeContext

	// IsFuncDeclarationContext differentiates from other interfaces.
	IsFuncDeclarationContext()
}

type FuncDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclarationContext() *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcDeclaration
	return p
}

func InitEmptyFuncDeclarationContext(p *FuncDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcDeclaration
}

func (*FuncDeclarationContext) IsFuncDeclarationContext() {}

func NewFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclarationContext {
	var p = new(FuncDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_funcDeclaration

	return p
}

func (s *FuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GrammarParserFUNC, 0)
}

func (s *FuncDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *FuncDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclarationContext) ParamsFuncDeclaration() IParamsFuncDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsFuncDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsFuncDeclarationContext)
}

func (s *FuncDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFuncDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) FuncDeclaration() (localctx IFuncDeclarationContext) {
	localctx = NewFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_funcDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(366)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserUNDERSCORE || _la == GrammarParserID {
		{
			p.SetState(368)
			p.ParamsFuncDeclaration()
		}

	}
	{
		p.SetState(371)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__14 {
		{
			p.SetState(372)
			p.Match(GrammarParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Type_()
		}

	}
	{
		p.SetState(376)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Block()
	}
	{
		p.SetState(378)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsFuncDeclarationContext is an interface to support dynamic dispatch.
type IParamsFuncDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	VectorType() IVectorTypeContext
	INOUT() antlr.TerminalNode
	ParamsFuncDeclaration() IParamsFuncDeclarationContext
	UNDERSCORE() antlr.TerminalNode

	// IsParamsFuncDeclarationContext differentiates from other interfaces.
	IsParamsFuncDeclarationContext()
}

type ParamsFuncDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsFuncDeclarationContext() *ParamsFuncDeclarationContext {
	var p = new(ParamsFuncDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_paramsFuncDeclaration
	return p
}

func InitEmptyParamsFuncDeclarationContext(p *ParamsFuncDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_paramsFuncDeclaration
}

func (*ParamsFuncDeclarationContext) IsParamsFuncDeclarationContext() {}

func NewParamsFuncDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsFuncDeclarationContext {
	var p = new(ParamsFuncDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_paramsFuncDeclaration

	return p
}

func (s *ParamsFuncDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsFuncDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GrammarParserID)
}

func (s *ParamsFuncDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GrammarParserID, i)
}

func (s *ParamsFuncDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *ParamsFuncDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParamsFuncDeclarationContext) VectorType() IVectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorTypeContext)
}

func (s *ParamsFuncDeclarationContext) INOUT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINOUT, 0)
}

func (s *ParamsFuncDeclarationContext) ParamsFuncDeclaration() IParamsFuncDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsFuncDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsFuncDeclarationContext)
}

func (s *ParamsFuncDeclarationContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(GrammarParserUNDERSCORE, 0)
}

func (s *ParamsFuncDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsFuncDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsFuncDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitParamsFuncDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ParamsFuncDeclaration() (localctx IParamsFuncDeclarationContext) {
	localctx = NewParamsFuncDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_paramsFuncDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(380)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserUNDERSCORE || _la == GrammarParserID) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(383)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserINOUT {
		{
			p.SetState(385)
			p.Match(GrammarParserINOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserSTRING, GrammarParserBOOL, GrammarParserCHARACTER:
		{
			p.SetState(388)
			p.Type_()
		}

	case GrammarParserT__0:
		{
			p.SetState(389)
			p.VectorType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__2 {
		{
			p.SetState(392)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.ParamsFuncDeclaration()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ParamsFuncCall() IParamsFuncCallContext

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_funcCall
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *FuncCallContext) ParamsFuncCall() IParamsFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsFuncCallContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(397)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36530844729376) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&15) != 0) {
		{
			p.SetState(398)
			p.ParamsFuncCall()
		}

	}
	{
		p.SetState(401)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsFuncCallContext is an interface to support dynamic dispatch.
type IParamsFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	REFERENCE() antlr.TerminalNode
	ParamsFuncCall() IParamsFuncCallContext

	// IsParamsFuncCallContext differentiates from other interfaces.
	IsParamsFuncCallContext()
}

type ParamsFuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsFuncCallContext() *ParamsFuncCallContext {
	var p = new(ParamsFuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_paramsFuncCall
	return p
}

func InitEmptyParamsFuncCallContext(p *ParamsFuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_paramsFuncCall
}

func (*ParamsFuncCallContext) IsParamsFuncCallContext() {}

func NewParamsFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsFuncCallContext {
	var p = new(ParamsFuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_paramsFuncCall

	return p
}

func (s *ParamsFuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsFuncCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamsFuncCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ParamsFuncCallContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *ParamsFuncCallContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(GrammarParserREFERENCE, 0)
}

func (s *ParamsFuncCallContext) ParamsFuncCall() IParamsFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsFuncCallContext)
}

func (s *ParamsFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsFuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitParamsFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ParamsFuncCall() (localctx IParamsFuncCallContext) {
	localctx = NewParamsFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrammarParserRULE_paramsFuncCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(403)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserREFERENCE {
		{
			p.SetState(407)
			p.Match(GrammarParserREFERENCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(410)
		p.expr(0)
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__2 {
		{
			p.SetState(411)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.ParamsFuncCall()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GrammarParserPRINT, 0)
}

func (s *PrintStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrammarParserRULE_printStmt)
	var _la int

	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			p.Match(GrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expr(0)
		}
		{
			p.SetState(418)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(GrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.expr(0)
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserT__2 {
			{
				p.SetState(423)
				p.Match(GrammarParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(424)
				p.expr(0)
			}

			p.SetState(429)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(430)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode
	IfStmt() IIfStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GrammarParserRULE_ifStmt)
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(434)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.expr(0)
		}
		{
			p.SetState(436)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Block()
		}
		{
			p.SetState(438)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(441)
			p.expr(0)
		}
		{
			p.SetState(442)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.Block()
		}
		{
			p.SetState(444)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Block()
		}
		{
			p.SetState(448)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(450)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)
			p.expr(0)
		}
		{
			p.SetState(452)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Block()
		}
		{
			p.SetState(454)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.IfStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	Cases() ICasesContext

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
	return p
}

func InitEmptySwitchStmtContext(p *SwitchStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchStmt
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) SwitchStmt() (localctx ISwitchStmtContext) {
	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GrammarParserRULE_switchStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.expr(0)
	}
	{
		p.SetState(462)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.Cases()
	}
	{
		p.SetState(464)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCaseStmt() []ICaseStmtContext
	CaseStmt(i int) ICaseStmtContext
	AllDefaultStmt() []IDefaultStmtContext
	DefaultStmt(i int) IDefaultStmtContext

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) AllCaseStmt() []ICaseStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseStmtContext); ok {
			len++
		}
	}

	tst := make([]ICaseStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseStmtContext); ok {
			tst[i] = t.(ICaseStmtContext)
			i++
		}
	}

	return tst
}

func (s *CasesContext) CaseStmt(i int) ICaseStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseStmtContext)
}

func (s *CasesContext) AllDefaultStmt() []IDefaultStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefaultStmtContext); ok {
			len++
		}
	}

	tst := make([]IDefaultStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefaultStmtContext); ok {
			tst[i] = t.(IDefaultStmtContext)
			i++
		}
	}

	return tst
}

func (s *CasesContext) DefaultStmt(i int) IDefaultStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultStmtContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCases(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GrammarParserRULE_cases)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserCASE || _la == GrammarParserDEFAULT {
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserCASE:
			{
				p.SetState(466)
				p.CaseStmt()
			}

		case GrammarParserDEFAULT:
			{
				p.SetState(467)
				p.DefaultStmt()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseStmtContext is an interface to support dynamic dispatch.
type ICaseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Block() IBlockContext
	BREAK() antlr.TerminalNode

	// IsCaseStmtContext differentiates from other interfaces.
	IsCaseStmtContext()
}

type CaseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStmtContext() *CaseStmtContext {
	var p = new(CaseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseStmt
	return p
}

func InitEmptyCaseStmtContext(p *CaseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_caseStmt
}

func (*CaseStmtContext) IsCaseStmtContext() {}

func NewCaseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStmtContext {
	var p = new(CaseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_caseStmt

	return p
}

func (s *CaseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *CaseStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CaseStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *CaseStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *CaseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCaseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) CaseStmt() (localctx ICaseStmtContext) {
	localctx = NewCaseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GrammarParserRULE_caseStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.Match(GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(474)
		p.expr(0)
	}
	{
		p.SetState(475)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.Block()
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserBREAK {
		{
			p.SetState(477)
			p.Match(GrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultStmtContext is an interface to support dynamic dispatch.
type IDefaultStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Block() IBlockContext
	BREAK() antlr.TerminalNode

	// IsDefaultStmtContext differentiates from other interfaces.
	IsDefaultStmtContext()
}

type DefaultStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultStmtContext() *DefaultStmtContext {
	var p = new(DefaultStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultStmt
	return p
}

func InitEmptyDefaultStmtContext(p *DefaultStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_defaultStmt
}

func (*DefaultStmtContext) IsDefaultStmtContext() {}

func NewDefaultStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStmtContext {
	var p = new(DefaultStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_defaultStmt

	return p
}

func (s *DefaultStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultStmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *DefaultStmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *DefaultStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *DefaultStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDefaultStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) DefaultStmt() (localctx IDefaultStmtContext) {
	localctx = NewDefaultStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrammarParserRULE_defaultStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
		p.Block()
	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserBREAK {
		{
			p.SetState(483)
			p.Match(GrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(487)
		p.expr(0)
	}
	{
		p.SetState(488)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.Block()
	}
	{
		p.SetState(490)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	IN() antlr.TerminalNode
	Block() IBlockContext
	ID() antlr.TerminalNode
	UNDERSCORE() antlr.TerminalNode
	Expr() IExprContext
	Range_() IRangeContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForStmtContext) IN() antlr.TerminalNode {
	return s.GetToken(GrammarParserIN, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ForStmtContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(GrammarParserUNDERSCORE, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GrammarParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserUNDERSCORE || _la == GrammarParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(494)
		p.Match(GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(495)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(496)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(499)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(500)
		p.Block()
	}
	{
		p.SetState(501)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GrammarParserRULE_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.expr(0)
	}
	{
		p.SetState(504)
		p.Match(GrammarParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(505)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardStmtContext is an interface to support dynamic dispatch.
type IGuardStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsGuardStmtContext differentiates from other interfaces.
	IsGuardStmtContext()
}

type GuardStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardStmtContext() *GuardStmtContext {
	var p = new(GuardStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardStmt
	return p
}

func InitEmptyGuardStmtContext(p *GuardStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardStmt
}

func (*GuardStmtContext) IsGuardStmtContext() {}

func NewGuardStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardStmtContext {
	var p = new(GuardStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_guardStmt

	return p
}

func (s *GuardStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardStmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(GrammarParserGUARD, 0)
}

func (s *GuardStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *GuardStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitGuardStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) GuardStmt() (localctx IGuardStmtContext) {
	localctx = NewGuardStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GrammarParserRULE_guardStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(GrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		p.expr(0)
	}
	{
		p.SetState(509)
		p.Match(GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
		p.Block()
	}
	{
		p.SetState(512)
		p.Match(GrammarParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GrammarParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(514)
		p.Match(GrammarParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GrammarParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(516)
		p.Match(GrammarParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GrammarParserRULE_returnStmt)
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(518)
			p.Match(GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(519)
			p.Match(GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringEmbededExprContext struct {
	ExprContext
}

func NewStringEmbededExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringEmbededExprContext {
	var p = new(StringEmbededExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringEmbededExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringEmbededExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *StringEmbededExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StringEmbededExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStringEmbededExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExprContext struct {
	ExprContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STR() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTR, 0)
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(GrammarParserNIL, 0)
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNilExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsEmptyExprContext struct {
	ExprContext
}

func NewIsEmptyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEmptyExprContext {
	var p = new(IsEmptyExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IsEmptyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEmptyExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IsEmptyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIsEmptyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewLogicalOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOperationExprContext {
	var p = new(LogicalOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *LogicalOperationExprContext) GetRight() IExprContext { return s.right }

func (s *LogicalOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *LogicalOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *LogicalOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogicalOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitLogicalOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatEmbededExprContext struct {
	ExprContext
}

func NewFloatEmbededExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatEmbededExprContext {
	var p = new(FloatEmbededExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FloatEmbededExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatEmbededExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GrammarParserFLOAT, 0)
}

func (s *FloatEmbededExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloatEmbededExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFloatEmbededExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorAccessExprContext struct {
	ExprContext
}

func NewVectorAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorAccessExprContext {
	var p = new(VectorAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorAccessExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *VectorAccessExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVectorAccessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	ExprContext
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithmeticOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewArithmeticOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticOperationExprContext {
	var p = new(ArithmeticOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArithmeticOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *ArithmeticOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithmeticOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *ArithmeticOperationExprContext) GetRight() IExprContext { return s.right }

func (s *ArithmeticOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ArithmeticOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ArithmeticOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArithmeticOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArithmeticOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewComparisonOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonOperationExprContext {
	var p = new(ComparisonOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *ComparisonOperationExprContext) GetRight() IExprContext { return s.right }

func (s *ComparisonOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ComparisonOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ComparisonOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparisonOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitComparisonOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalOperationExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewRelationalOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalOperationExprContext {
	var p = new(RelationalOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RelationalOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalOperationExprContext) GetLeft() IExprContext { return s.left }

func (s *RelationalOperationExprContext) GetRight() IExprContext { return s.right }

func (s *RelationalOperationExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *RelationalOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *RelationalOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOperationExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RelationalOperationExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelationalOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitRelationalOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharacterExprContext struct {
	ExprContext
}

func NewCharacterExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharacterExprContext {
	var p = new(CharacterExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CharacterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserCHAR, 0)
}

func (s *CharacterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCharacterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DigitExprContext struct {
	ExprContext
}

func NewDigitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitExprContext {
	var p = new(DigitExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DigitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitExprContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIGIT, 0)
}

func (s *DigitExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDigitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CountExprContext struct {
	ExprContext
}

func NewCountExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountExprContext {
	var p = new(CountExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CountExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *CountExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCountExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntEmbededExprContext struct {
	ExprContext
}

func NewIntEmbededExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntEmbededExprContext {
	var p = new(IntEmbededExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntEmbededExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntEmbededExprContext) INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINT, 0)
}

func (s *IntEmbededExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IntEmbededExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIntEmbededExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryOperationExprContext struct {
	ExprContext
	op    antlr.Token
	right IExprContext
}

func NewUnaryOperationExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOperationExprContext {
	var p = new(UnaryOperationExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryOperationExprContext) GetOp() antlr.Token { return s.op }

func (s *UnaryOperationExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryOperationExprContext) GetRight() IExprContext { return s.right }

func (s *UnaryOperationExprContext) SetRight(v IExprContext) { s.right = v }

func (s *UnaryOperationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperationExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryOperationExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitUnaryOperationExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanExprContext struct {
	ExprContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallExprContext struct {
	ExprContext
}

func NewFuncCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallExprContext {
	var p = new(FuncCallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FuncCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExprContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *FuncCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitFuncCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 72
	p.EnterRecursionRule(localctx, 72, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(524)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.expr(0)
		}
		{
			p.SetState(526)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryOperationExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(528)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryOperationExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__16 || _la == GrammarParserT__17) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryOperationExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(529)

			var _x = p.expr(20)

			localctx.(*UnaryOperationExprContext).right = _x
		}

	case 3:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(530)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__30 || _la == GrammarParserT__31) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(531)
			p.Match(GrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDigitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(532)
			p.Match(GrammarParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewStringExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(533)
			p.Match(GrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewCharacterExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(534)
			p.Match(GrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewFuncCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(535)
			p.FuncCall()
		}

	case 9:
		localctx = NewCountExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(536)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(537)
			p.Match(GrammarParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewIsEmptyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(538)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(GrammarParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewVectorAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(540)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.expr(0)
		}
		{
			p.SetState(543)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(545)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewIntEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(546)
			p.Match(GrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)
			p.expr(0)
		}
		{
			p.SetState(549)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewFloatEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(551)
			p.Match(GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.expr(0)
		}
		{
			p.SetState(554)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewStringEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(556)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.expr(0)
		}
		{
			p.SetState(559)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(581)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(563)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(564)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3670016) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(565)

					var _x = p.expr(20)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(566)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(567)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__17 || _la == GrammarParserT__21) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(568)

					var _x = p.expr(19)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 3:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(569)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(570)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__22 || _la == GrammarParserT__23) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(571)

					var _x = p.expr(18)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 4:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(572)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(573)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__24 || _la == GrammarParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(574)

					var _x = p.expr(17)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 5:
				localctx = NewComparisonOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparisonOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(575)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(576)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__26 || _la == GrammarParserT__27) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(577)

					var _x = p.expr(16)

					localctx.(*ComparisonOperationExprContext).right = _x
				}

			case 6:
				localctx = NewLogicalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(578)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(579)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__28 || _la == GrammarParserT__29) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(580)

					var _x = p.expr(15)

					localctx.(*LogicalOperationExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(585)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINT, 0)
}

func (s *TypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GrammarParserFLOAT, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GrammarParserBOOL, 0)
}

func (s *TypeContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(GrammarParserCHARACTER, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(586)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1065151889408) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 36:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
