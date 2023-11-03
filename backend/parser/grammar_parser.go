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
		"'.remove'", "'at'", "'struct'", "'{'", "'}'", "'.'", "'->'", "'...'",
		"'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='", "'>'", "'<='", "'<'",
		"'=='", "'!='", "'&&'", "'||'", "'true'", "'false'", "'.count'", "'.isEmpty'",
		"'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'nil'", "'var'",
		"'let'", "'func'", "'inout'", "'&'", "'print'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'while'", "'for'", "'in'", "'guard'", "'continue'",
		"'break'", "'return'", "';'", "':'", "'='", "'?'", "'+='", "'-='", "'_'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"INT", "FLOAT", "STRING", "BOOL", "CHARACTER", "NIL", "VAR", "LET",
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
		"structDeclaration", "structAttributes", "structInstance", "attributesCall",
		"structAccess", "structAssignment", "funcDeclaration", "paramsFuncDeclaration",
		"funcCall", "paramsFuncCall", "printStmt", "ifStmt", "switchStmt", "cases",
		"caseStmt", "defaultStmt", "whileStmt", "forStmt", "range", "guardStmt",
		"breakStmt", "continueStmt", "returnStmt", "expr", "type",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 605, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 83, 8, 1, 10,
		1, 12, 1, 86, 9, 1, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 2, 1, 2, 3, 2, 94, 8,
		2, 1, 2, 1, 2, 3, 2, 98, 8, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 1, 2,
		3, 2, 106, 8, 2, 1, 2, 1, 2, 3, 2, 110, 8, 2, 1, 2, 1, 2, 3, 2, 114, 8,
		2, 1, 2, 1, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 2, 1, 2,
		3, 2, 126, 8, 2, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 2, 1, 2, 3, 2, 134, 8,
		2, 1, 2, 1, 2, 3, 2, 138, 8, 2, 1, 2, 1, 2, 3, 2, 142, 8, 2, 1, 2, 1, 2,
		3, 2, 146, 8, 2, 1, 2, 1, 2, 3, 2, 150, 8, 2, 1, 2, 1, 2, 3, 2, 154, 8,
		2, 1, 2, 1, 2, 3, 2, 158, 8, 2, 1, 2, 1, 2, 3, 2, 162, 8, 2, 1, 2, 1, 2,
		3, 2, 166, 8, 2, 1, 2, 1, 2, 3, 2, 170, 8, 2, 1, 2, 1, 2, 3, 2, 174, 8,
		2, 3, 2, 176, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 195, 8, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 203, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 213, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 223, 8, 4, 1, 4, 1, 4, 3, 4, 227, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 240,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 251,
		8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 264, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 269, 8, 7, 10, 7, 12, 7, 272,
		9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 308, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 323, 8, 12,
		1, 13, 1, 13, 3, 13, 327, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 333,
		8, 13, 1, 13, 5, 13, 336, 8, 13, 10, 13, 12, 13, 339, 9, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 345, 8, 14, 10, 14, 12, 14, 348, 9, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 360,
		8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 369, 8,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 385, 8, 20, 1, 20, 1, 20, 1, 20, 3,
		20, 390, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 3, 21, 397, 8, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 402, 8, 21, 1, 21, 1, 21, 3, 21, 406, 8, 21, 1,
		21, 1, 21, 3, 21, 410, 8, 21, 1, 22, 1, 22, 1, 22, 3, 22, 415, 8, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 3, 23, 421, 8, 23, 1, 23, 3, 23, 424, 8, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 429, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 441, 8, 24, 10, 24, 12, 24, 444,
		9, 24, 1, 24, 1, 24, 3, 24, 448, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 474,
		8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 484,
		8, 27, 10, 27, 12, 27, 487, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3,
		28, 494, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 500, 8, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		513, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 3, 36, 537, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 3, 37, 578, 8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 5, 37, 598, 8, 37, 10, 37, 12, 37, 601, 9, 37, 1,
		38, 1, 38, 1, 38, 0, 1, 74, 39, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 0, 11, 1, 0, 40, 41, 2, 0, 64,
		64, 68, 68, 1, 0, 16, 17, 1, 0, 30, 31, 1, 0, 18, 20, 2, 0, 17, 17, 21,
		21, 1, 0, 22, 23, 1, 0, 24, 25, 1, 0, 26, 27, 1, 0, 28, 29, 2, 0, 34, 38,
		68, 68, 669, 0, 78, 1, 0, 0, 0, 2, 84, 1, 0, 0, 0, 4, 175, 1, 0, 0, 0,
		6, 194, 1, 0, 0, 0, 8, 226, 1, 0, 0, 0, 10, 239, 1, 0, 0, 0, 12, 263, 1,
		0, 0, 0, 14, 265, 1, 0, 0, 0, 16, 273, 1, 0, 0, 0, 18, 277, 1, 0, 0, 0,
		20, 283, 1, 0, 0, 0, 22, 288, 1, 0, 0, 0, 24, 322, 1, 0, 0, 0, 26, 324,
		1, 0, 0, 0, 28, 340, 1, 0, 0, 0, 30, 351, 1, 0, 0, 0, 32, 356, 1, 0, 0,
		0, 34, 363, 1, 0, 0, 0, 36, 370, 1, 0, 0, 0, 38, 374, 1, 0, 0, 0, 40, 380,
		1, 0, 0, 0, 42, 396, 1, 0, 0, 0, 44, 411, 1, 0, 0, 0, 46, 420, 1, 0, 0,
		0, 48, 447, 1, 0, 0, 0, 50, 473, 1, 0, 0, 0, 52, 475, 1, 0, 0, 0, 54, 485,
		1, 0, 0, 0, 56, 488, 1, 0, 0, 0, 58, 495, 1, 0, 0, 0, 60, 501, 1, 0, 0,
		0, 62, 507, 1, 0, 0, 0, 64, 518, 1, 0, 0, 0, 66, 522, 1, 0, 0, 0, 68, 529,
		1, 0, 0, 0, 70, 531, 1, 0, 0, 0, 72, 536, 1, 0, 0, 0, 74, 577, 1, 0, 0,
		0, 76, 602, 1, 0, 0, 0, 78, 79, 3, 2, 1, 0, 79, 80, 5, 0, 0, 1, 80, 1,
		1, 0, 0, 0, 81, 83, 3, 4, 2, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 3, 1, 0, 0, 0, 86, 84, 1, 0,
		0, 0, 87, 89, 3, 6, 3, 0, 88, 90, 5, 58, 0, 0, 89, 88, 1, 0, 0, 0, 89,
		90, 1, 0, 0, 0, 90, 176, 1, 0, 0, 0, 91, 93, 3, 8, 4, 0, 92, 94, 5, 58,
		0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 176, 1, 0, 0, 0, 95,
		97, 3, 10, 5, 0, 96, 98, 5, 58, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0,
		0, 0, 98, 176, 1, 0, 0, 0, 99, 101, 3, 12, 6, 0, 100, 102, 5, 58, 0, 0,
		101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 176, 1, 0, 0, 0, 103,
		105, 3, 18, 9, 0, 104, 106, 5, 58, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106,
		1, 0, 0, 0, 106, 176, 1, 0, 0, 0, 107, 109, 3, 20, 10, 0, 108, 110, 5,
		58, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 176, 1, 0, 0,
		0, 111, 113, 3, 22, 11, 0, 112, 114, 5, 58, 0, 0, 113, 112, 1, 0, 0, 0,
		113, 114, 1, 0, 0, 0, 114, 176, 1, 0, 0, 0, 115, 117, 3, 40, 20, 0, 116,
		118, 5, 58, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 176,
		1, 0, 0, 0, 119, 121, 3, 44, 22, 0, 120, 122, 5, 58, 0, 0, 121, 120, 1,
		0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 176, 1, 0, 0, 0, 123, 125, 3, 48, 24,
		0, 124, 126, 5, 58, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		176, 1, 0, 0, 0, 127, 129, 3, 50, 25, 0, 128, 130, 5, 58, 0, 0, 129, 128,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 176, 1, 0, 0, 0, 131, 133, 3, 52,
		26, 0, 132, 134, 5, 58, 0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0,
		0, 134, 176, 1, 0, 0, 0, 135, 137, 3, 60, 30, 0, 136, 138, 5, 58, 0, 0,
		137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 176, 1, 0, 0, 0, 139,
		141, 3, 62, 31, 0, 140, 142, 5, 58, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142,
		1, 0, 0, 0, 142, 176, 1, 0, 0, 0, 143, 145, 3, 66, 33, 0, 144, 146, 5,
		58, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 176, 1, 0, 0,
		0, 147, 149, 3, 68, 34, 0, 148, 150, 5, 58, 0, 0, 149, 148, 1, 0, 0, 0,
		149, 150, 1, 0, 0, 0, 150, 176, 1, 0, 0, 0, 151, 153, 3, 70, 35, 0, 152,
		154, 5, 58, 0, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 176,
		1, 0, 0, 0, 155, 157, 3, 72, 36, 0, 156, 158, 5, 58, 0, 0, 157, 156, 1,
		0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 176, 1, 0, 0, 0, 159, 161, 3, 28, 14,
		0, 160, 162, 5, 58, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		176, 1, 0, 0, 0, 163, 165, 3, 32, 16, 0, 164, 166, 5, 58, 0, 0, 165, 164,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 176, 1, 0, 0, 0, 167, 169, 3, 36,
		18, 0, 168, 170, 5, 58, 0, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0,
		0, 170, 176, 1, 0, 0, 0, 171, 173, 3, 38, 19, 0, 172, 174, 5, 58, 0, 0,
		173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175,
		87, 1, 0, 0, 0, 175, 91, 1, 0, 0, 0, 175, 95, 1, 0, 0, 0, 175, 99, 1, 0,
		0, 0, 175, 103, 1, 0, 0, 0, 175, 107, 1, 0, 0, 0, 175, 111, 1, 0, 0, 0,
		175, 115, 1, 0, 0, 0, 175, 119, 1, 0, 0, 0, 175, 123, 1, 0, 0, 0, 175,
		127, 1, 0, 0, 0, 175, 131, 1, 0, 0, 0, 175, 135, 1, 0, 0, 0, 175, 139,
		1, 0, 0, 0, 175, 143, 1, 0, 0, 0, 175, 147, 1, 0, 0, 0, 175, 151, 1, 0,
		0, 0, 175, 155, 1, 0, 0, 0, 175, 159, 1, 0, 0, 0, 175, 163, 1, 0, 0, 0,
		175, 167, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 176, 5, 1, 0, 0, 0, 177, 178,
		5, 40, 0, 0, 178, 179, 5, 68, 0, 0, 179, 180, 5, 59, 0, 0, 180, 181, 3,
		76, 38, 0, 181, 182, 5, 60, 0, 0, 182, 183, 3, 74, 37, 0, 183, 195, 1,
		0, 0, 0, 184, 185, 5, 40, 0, 0, 185, 186, 5, 68, 0, 0, 186, 187, 5, 60,
		0, 0, 187, 195, 3, 74, 37, 0, 188, 189, 5, 40, 0, 0, 189, 190, 5, 68, 0,
		0, 190, 191, 5, 59, 0, 0, 191, 192, 3, 76, 38, 0, 192, 193, 5, 61, 0, 0,
		193, 195, 1, 0, 0, 0, 194, 177, 1, 0, 0, 0, 194, 184, 1, 0, 0, 0, 194,
		188, 1, 0, 0, 0, 195, 7, 1, 0, 0, 0, 196, 203, 5, 68, 0, 0, 197, 198, 5,
		68, 0, 0, 198, 199, 5, 1, 0, 0, 199, 200, 3, 74, 37, 0, 200, 201, 5, 2,
		0, 0, 201, 203, 1, 0, 0, 0, 202, 196, 1, 0, 0, 0, 202, 197, 1, 0, 0, 0,
		203, 204, 1, 0, 0, 0, 204, 205, 5, 60, 0, 0, 205, 227, 3, 74, 37, 0, 206,
		213, 5, 68, 0, 0, 207, 208, 5, 68, 0, 0, 208, 209, 5, 1, 0, 0, 209, 210,
		3, 74, 37, 0, 210, 211, 5, 2, 0, 0, 211, 213, 1, 0, 0, 0, 212, 206, 1,
		0, 0, 0, 212, 207, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 5, 62, 0,
		0, 215, 227, 3, 74, 37, 0, 216, 223, 5, 68, 0, 0, 217, 218, 5, 68, 0, 0,
		218, 219, 5, 1, 0, 0, 219, 220, 3, 74, 37, 0, 220, 221, 5, 2, 0, 0, 221,
		223, 1, 0, 0, 0, 222, 216, 1, 0, 0, 0, 222, 217, 1, 0, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 225, 5, 63, 0, 0, 225, 227, 3, 74, 37, 0, 226, 202, 1,
		0, 0, 0, 226, 212, 1, 0, 0, 0, 226, 222, 1, 0, 0, 0, 227, 9, 1, 0, 0, 0,
		228, 229, 5, 41, 0, 0, 229, 230, 5, 68, 0, 0, 230, 231, 5, 59, 0, 0, 231,
		232, 3, 76, 38, 0, 232, 233, 5, 60, 0, 0, 233, 234, 3, 74, 37, 0, 234,
		240, 1, 0, 0, 0, 235, 236, 5, 41, 0, 0, 236, 237, 5, 68, 0, 0, 237, 238,
		5, 60, 0, 0, 238, 240, 3, 74, 37, 0, 239, 228, 1, 0, 0, 0, 239, 235, 1,
		0, 0, 0, 240, 11, 1, 0, 0, 0, 241, 242, 5, 40, 0, 0, 242, 243, 5, 68, 0,
		0, 243, 244, 5, 59, 0, 0, 244, 245, 5, 1, 0, 0, 245, 246, 3, 76, 38, 0,
		246, 247, 5, 2, 0, 0, 247, 248, 5, 60, 0, 0, 248, 250, 5, 1, 0, 0, 249,
		251, 3, 14, 7, 0, 250, 249, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 253, 5, 2, 0, 0, 253, 264, 1, 0, 0, 0, 254, 255, 5, 40,
		0, 0, 255, 256, 5, 68, 0, 0, 256, 257, 5, 59, 0, 0, 257, 258, 5, 1, 0,
		0, 258, 259, 3, 76, 38, 0, 259, 260, 5, 2, 0, 0, 260, 261, 5, 60, 0, 0,
		261, 262, 5, 68, 0, 0, 262, 264, 1, 0, 0, 0, 263, 241, 1, 0, 0, 0, 263,
		254, 1, 0, 0, 0, 264, 13, 1, 0, 0, 0, 265, 270, 3, 74, 37, 0, 266, 267,
		5, 3, 0, 0, 267, 269, 3, 74, 37, 0, 268, 266, 1, 0, 0, 0, 269, 272, 1,
		0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 15, 1, 0, 0,
		0, 272, 270, 1, 0, 0, 0, 273, 274, 5, 1, 0, 0, 274, 275, 3, 76, 38, 0,
		275, 276, 5, 2, 0, 0, 276, 17, 1, 0, 0, 0, 277, 278, 5, 68, 0, 0, 278,
		279, 5, 4, 0, 0, 279, 280, 5, 5, 0, 0, 280, 281, 3, 74, 37, 0, 281, 282,
		5, 6, 0, 0, 282, 19, 1, 0, 0, 0, 283, 284, 5, 68, 0, 0, 284, 285, 5, 7,
		0, 0, 285, 286, 5, 5, 0, 0, 286, 287, 5, 6, 0, 0, 287, 21, 1, 0, 0, 0,
		288, 289, 5, 68, 0, 0, 289, 290, 5, 8, 0, 0, 290, 291, 5, 5, 0, 0, 291,
		292, 5, 9, 0, 0, 292, 293, 5, 59, 0, 0, 293, 294, 3, 74, 37, 0, 294, 295,
		5, 6, 0, 0, 295, 23, 1, 0, 0, 0, 296, 297, 5, 40, 0, 0, 297, 298, 5, 68,
		0, 0, 298, 299, 5, 59, 0, 0, 299, 300, 5, 1, 0, 0, 300, 301, 5, 1, 0, 0,
		301, 302, 3, 76, 38, 0, 302, 303, 5, 2, 0, 0, 303, 304, 5, 2, 0, 0, 304,
		305, 5, 60, 0, 0, 305, 307, 5, 1, 0, 0, 306, 308, 3, 26, 13, 0, 307, 306,
		1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 5, 2,
		0, 0, 310, 323, 1, 0, 0, 0, 311, 312, 5, 40, 0, 0, 312, 313, 5, 68, 0,
		0, 313, 314, 5, 59, 0, 0, 314, 315, 5, 1, 0, 0, 315, 316, 5, 1, 0, 0, 316,
		317, 3, 76, 38, 0, 317, 318, 5, 2, 0, 0, 318, 319, 5, 2, 0, 0, 319, 320,
		5, 60, 0, 0, 320, 321, 5, 68, 0, 0, 321, 323, 1, 0, 0, 0, 322, 296, 1,
		0, 0, 0, 322, 311, 1, 0, 0, 0, 323, 25, 1, 0, 0, 0, 324, 326, 5, 1, 0,
		0, 325, 327, 3, 14, 7, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327,
		328, 1, 0, 0, 0, 328, 337, 5, 2, 0, 0, 329, 330, 5, 3, 0, 0, 330, 332,
		5, 1, 0, 0, 331, 333, 3, 14, 7, 0, 332, 331, 1, 0, 0, 0, 332, 333, 1, 0,
		0, 0, 333, 334, 1, 0, 0, 0, 334, 336, 5, 2, 0, 0, 335, 329, 1, 0, 0, 0,
		336, 339, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338,
		27, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 340, 341, 5, 10, 0, 0, 341, 342,
		5, 68, 0, 0, 342, 346, 5, 11, 0, 0, 343, 345, 3, 30, 15, 0, 344, 343, 1,
		0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0,
		0, 347, 349, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 350, 5, 12, 0, 0, 350,
		29, 1, 0, 0, 0, 351, 352, 7, 0, 0, 0, 352, 353, 5, 68, 0, 0, 353, 354,
		5, 59, 0, 0, 354, 355, 3, 76, 38, 0, 355, 31, 1, 0, 0, 0, 356, 357, 5,
		68, 0, 0, 357, 359, 5, 5, 0, 0, 358, 360, 3, 34, 17, 0, 359, 358, 1, 0,
		0, 0, 359, 360, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362, 5, 6, 0, 0,
		362, 33, 1, 0, 0, 0, 363, 364, 5, 68, 0, 0, 364, 365, 5, 59, 0, 0, 365,
		368, 3, 74, 37, 0, 366, 367, 5, 3, 0, 0, 367, 369, 3, 34, 17, 0, 368, 366,
		1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 35, 1, 0, 0, 0, 370, 371, 5, 68,
		0, 0, 371, 372, 5, 13, 0, 0, 372, 373, 5, 68, 0, 0, 373, 37, 1, 0, 0, 0,
		374, 375, 5, 68, 0, 0, 375, 376, 5, 13, 0, 0, 376, 377, 5, 68, 0, 0, 377,
		378, 5, 60, 0, 0, 378, 379, 3, 74, 37, 0, 379, 39, 1, 0, 0, 0, 380, 381,
		5, 42, 0, 0, 381, 382, 5, 68, 0, 0, 382, 384, 5, 5, 0, 0, 383, 385, 3,
		42, 21, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 386, 1, 0,
		0, 0, 386, 389, 5, 6, 0, 0, 387, 388, 5, 14, 0, 0, 388, 390, 3, 76, 38,
		0, 389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391,
		392, 5, 11, 0, 0, 392, 393, 3, 2, 1, 0, 393, 394, 5, 12, 0, 0, 394, 41,
		1, 0, 0, 0, 395, 397, 7, 1, 0, 0, 396, 395, 1, 0, 0, 0, 396, 397, 1, 0,
		0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 5, 68, 0, 0, 399, 401, 5, 59, 0,
		0, 400, 402, 5, 43, 0, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402,
		405, 1, 0, 0, 0, 403, 406, 3, 76, 38, 0, 404, 406, 3, 16, 8, 0, 405, 403,
		1, 0, 0, 0, 405, 404, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 408, 5, 3,
		0, 0, 408, 410, 3, 42, 21, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0,
		0, 410, 43, 1, 0, 0, 0, 411, 412, 5, 68, 0, 0, 412, 414, 5, 5, 0, 0, 413,
		415, 3, 46, 23, 0, 414, 413, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416,
		1, 0, 0, 0, 416, 417, 5, 6, 0, 0, 417, 45, 1, 0, 0, 0, 418, 419, 5, 68,
		0, 0, 419, 421, 5, 59, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0,
		421, 423, 1, 0, 0, 0, 422, 424, 5, 44, 0, 0, 423, 422, 1, 0, 0, 0, 423,
		424, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 428, 3, 74, 37, 0, 426, 427,
		5, 3, 0, 0, 427, 429, 3, 46, 23, 0, 428, 426, 1, 0, 0, 0, 428, 429, 1,
		0, 0, 0, 429, 47, 1, 0, 0, 0, 430, 431, 5, 45, 0, 0, 431, 432, 5, 5, 0,
		0, 432, 433, 3, 74, 37, 0, 433, 434, 5, 6, 0, 0, 434, 448, 1, 0, 0, 0,
		435, 436, 5, 45, 0, 0, 436, 437, 5, 5, 0, 0, 437, 442, 3, 74, 37, 0, 438,
		439, 5, 3, 0, 0, 439, 441, 3, 74, 37, 0, 440, 438, 1, 0, 0, 0, 441, 444,
		1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 445, 1, 0,
		0, 0, 444, 442, 1, 0, 0, 0, 445, 446, 5, 6, 0, 0, 446, 448, 1, 0, 0, 0,
		447, 430, 1, 0, 0, 0, 447, 435, 1, 0, 0, 0, 448, 49, 1, 0, 0, 0, 449, 450,
		5, 46, 0, 0, 450, 451, 3, 74, 37, 0, 451, 452, 5, 11, 0, 0, 452, 453, 3,
		2, 1, 0, 453, 454, 5, 12, 0, 0, 454, 474, 1, 0, 0, 0, 455, 456, 5, 46,
		0, 0, 456, 457, 3, 74, 37, 0, 457, 458, 5, 11, 0, 0, 458, 459, 3, 2, 1,
		0, 459, 460, 5, 12, 0, 0, 460, 461, 5, 47, 0, 0, 461, 462, 5, 11, 0, 0,
		462, 463, 3, 2, 1, 0, 463, 464, 5, 12, 0, 0, 464, 474, 1, 0, 0, 0, 465,
		466, 5, 46, 0, 0, 466, 467, 3, 74, 37, 0, 467, 468, 5, 11, 0, 0, 468, 469,
		3, 2, 1, 0, 469, 470, 5, 12, 0, 0, 470, 471, 5, 47, 0, 0, 471, 472, 3,
		50, 25, 0, 472, 474, 1, 0, 0, 0, 473, 449, 1, 0, 0, 0, 473, 455, 1, 0,
		0, 0, 473, 465, 1, 0, 0, 0, 474, 51, 1, 0, 0, 0, 475, 476, 5, 48, 0, 0,
		476, 477, 3, 74, 37, 0, 477, 478, 5, 11, 0, 0, 478, 479, 3, 54, 27, 0,
		479, 480, 5, 12, 0, 0, 480, 53, 1, 0, 0, 0, 481, 484, 3, 56, 28, 0, 482,
		484, 3, 58, 29, 0, 483, 481, 1, 0, 0, 0, 483, 482, 1, 0, 0, 0, 484, 487,
		1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 55, 1, 0,
		0, 0, 487, 485, 1, 0, 0, 0, 488, 489, 5, 49, 0, 0, 489, 490, 3, 74, 37,
		0, 490, 491, 5, 59, 0, 0, 491, 493, 3, 2, 1, 0, 492, 494, 5, 56, 0, 0,
		493, 492, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 57, 1, 0, 0, 0, 495, 496,
		5, 50, 0, 0, 496, 497, 5, 59, 0, 0, 497, 499, 3, 2, 1, 0, 498, 500, 5,
		56, 0, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 59, 1, 0, 0,
		0, 501, 502, 5, 51, 0, 0, 502, 503, 3, 74, 37, 0, 503, 504, 5, 11, 0, 0,
		504, 505, 3, 2, 1, 0, 505, 506, 5, 12, 0, 0, 506, 61, 1, 0, 0, 0, 507,
		508, 5, 52, 0, 0, 508, 509, 7, 1, 0, 0, 509, 512, 5, 53, 0, 0, 510, 513,
		3, 74, 37, 0, 511, 513, 3, 64, 32, 0, 512, 510, 1, 0, 0, 0, 512, 511, 1,
		0, 0, 0, 513, 514, 1, 0, 0, 0, 514, 515, 5, 11, 0, 0, 515, 516, 3, 2, 1,
		0, 516, 517, 5, 12, 0, 0, 517, 63, 1, 0, 0, 0, 518, 519, 3, 74, 37, 0,
		519, 520, 5, 15, 0, 0, 520, 521, 3, 74, 37, 0, 521, 65, 1, 0, 0, 0, 522,
		523, 5, 54, 0, 0, 523, 524, 3, 74, 37, 0, 524, 525, 5, 47, 0, 0, 525, 526,
		5, 11, 0, 0, 526, 527, 3, 2, 1, 0, 527, 528, 5, 12, 0, 0, 528, 67, 1, 0,
		0, 0, 529, 530, 5, 56, 0, 0, 530, 69, 1, 0, 0, 0, 531, 532, 5, 55, 0, 0,
		532, 71, 1, 0, 0, 0, 533, 537, 5, 57, 0, 0, 534, 535, 5, 57, 0, 0, 535,
		537, 3, 74, 37, 0, 536, 533, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 537, 73,
		1, 0, 0, 0, 538, 539, 6, 37, -1, 0, 539, 540, 5, 5, 0, 0, 540, 541, 3,
		74, 37, 0, 541, 542, 5, 6, 0, 0, 542, 578, 1, 0, 0, 0, 543, 544, 7, 2,
		0, 0, 544, 578, 3, 74, 37, 21, 545, 578, 7, 3, 0, 0, 546, 578, 5, 39, 0,
		0, 547, 578, 5, 65, 0, 0, 548, 578, 5, 66, 0, 0, 549, 578, 5, 67, 0, 0,
		550, 551, 5, 68, 0, 0, 551, 578, 5, 32, 0, 0, 552, 553, 5, 68, 0, 0, 553,
		578, 5, 33, 0, 0, 554, 555, 5, 68, 0, 0, 555, 556, 5, 1, 0, 0, 556, 557,
		3, 74, 37, 0, 557, 558, 5, 2, 0, 0, 558, 578, 1, 0, 0, 0, 559, 578, 3,
		32, 16, 0, 560, 578, 3, 36, 18, 0, 561, 578, 5, 68, 0, 0, 562, 563, 5,
		34, 0, 0, 563, 564, 5, 5, 0, 0, 564, 565, 3, 74, 37, 0, 565, 566, 5, 6,
		0, 0, 566, 578, 1, 0, 0, 0, 567, 568, 5, 35, 0, 0, 568, 569, 5, 5, 0, 0,
		569, 570, 3, 74, 37, 0, 570, 571, 5, 6, 0, 0, 571, 578, 1, 0, 0, 0, 572,
		573, 5, 36, 0, 0, 573, 574, 5, 5, 0, 0, 574, 575, 3, 74, 37, 0, 575, 576,
		5, 6, 0, 0, 576, 578, 1, 0, 0, 0, 577, 538, 1, 0, 0, 0, 577, 543, 1, 0,
		0, 0, 577, 545, 1, 0, 0, 0, 577, 546, 1, 0, 0, 0, 577, 547, 1, 0, 0, 0,
		577, 548, 1, 0, 0, 0, 577, 549, 1, 0, 0, 0, 577, 550, 1, 0, 0, 0, 577,
		552, 1, 0, 0, 0, 577, 554, 1, 0, 0, 0, 577, 559, 1, 0, 0, 0, 577, 560,
		1, 0, 0, 0, 577, 561, 1, 0, 0, 0, 577, 562, 1, 0, 0, 0, 577, 567, 1, 0,
		0, 0, 577, 572, 1, 0, 0, 0, 578, 599, 1, 0, 0, 0, 579, 580, 10, 20, 0,
		0, 580, 581, 7, 4, 0, 0, 581, 598, 3, 74, 37, 21, 582, 583, 10, 19, 0,
		0, 583, 584, 7, 5, 0, 0, 584, 598, 3, 74, 37, 20, 585, 586, 10, 18, 0,
		0, 586, 587, 7, 6, 0, 0, 587, 598, 3, 74, 37, 19, 588, 589, 10, 17, 0,
		0, 589, 590, 7, 7, 0, 0, 590, 598, 3, 74, 37, 18, 591, 592, 10, 16, 0,
		0, 592, 593, 7, 8, 0, 0, 593, 598, 3, 74, 37, 17, 594, 595, 10, 15, 0,
		0, 595, 596, 7, 9, 0, 0, 596, 598, 3, 74, 37, 16, 597, 579, 1, 0, 0, 0,
		597, 582, 1, 0, 0, 0, 597, 585, 1, 0, 0, 0, 597, 588, 1, 0, 0, 0, 597,
		591, 1, 0, 0, 0, 597, 594, 1, 0, 0, 0, 598, 601, 1, 0, 0, 0, 599, 597,
		1, 0, 0, 0, 599, 600, 1, 0, 0, 0, 600, 75, 1, 0, 0, 0, 601, 599, 1, 0,
		0, 0, 602, 603, 7, 10, 0, 0, 603, 77, 1, 0, 0, 0, 63, 84, 89, 93, 97, 101,
		105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161,
		165, 169, 173, 175, 194, 202, 212, 222, 226, 239, 250, 263, 270, 307, 322,
		326, 332, 337, 346, 359, 368, 384, 389, 396, 401, 405, 409, 414, 420, 423,
		428, 442, 447, 473, 483, 485, 493, 499, 512, 536, 577, 597, 599,
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
	GrammarParserINT           = 34
	GrammarParserFLOAT         = 35
	GrammarParserSTRING        = 36
	GrammarParserBOOL          = 37
	GrammarParserCHARACTER     = 38
	GrammarParserNIL           = 39
	GrammarParserVAR           = 40
	GrammarParserLET           = 41
	GrammarParserFUNC          = 42
	GrammarParserINOUT         = 43
	GrammarParserREFERENCE     = 44
	GrammarParserPRINT         = 45
	GrammarParserIF            = 46
	GrammarParserELSE          = 47
	GrammarParserSWITCH        = 48
	GrammarParserCASE          = 49
	GrammarParserDEFAULT       = 50
	GrammarParserWHILE         = 51
	GrammarParserFOR           = 52
	GrammarParserIN            = 53
	GrammarParserGUARD         = 54
	GrammarParserCONTINUE      = 55
	GrammarParserBREAK         = 56
	GrammarParserRETURN        = 57
	GrammarParserSEMICOLON     = 58
	GrammarParserCOLON         = 59
	GrammarParserEQUAL         = 60
	GrammarParserQUESTION_MARK = 61
	GrammarParserINCREMENT     = 62
	GrammarParserDECREMENT     = 63
	GrammarParserUNDERSCORE    = 64
	GrammarParserDIGIT         = 65
	GrammarParserSTR           = 66
	GrammarParserCHAR          = 67
	GrammarParserID            = 68
	GrammarParserWHITESPACE    = 69
	GrammarParserMULTI_COMMENT = 70
	GrammarParserLINE_COMMENT  = 71
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
	GrammarParserRULE_structAttributes        = 15
	GrammarParserRULE_structInstance          = 16
	GrammarParserRULE_attributesCall          = 17
	GrammarParserRULE_structAccess            = 18
	GrammarParserRULE_structAssignment        = 19
	GrammarParserRULE_funcDeclaration         = 20
	GrammarParserRULE_paramsFuncDeclaration   = 21
	GrammarParserRULE_funcCall                = 22
	GrammarParserRULE_paramsFuncCall          = 23
	GrammarParserRULE_printStmt               = 24
	GrammarParserRULE_ifStmt                  = 25
	GrammarParserRULE_switchStmt              = 26
	GrammarParserRULE_cases                   = 27
	GrammarParserRULE_caseStmt                = 28
	GrammarParserRULE_defaultStmt             = 29
	GrammarParserRULE_whileStmt               = 30
	GrammarParserRULE_forStmt                 = 31
	GrammarParserRULE_range                   = 32
	GrammarParserRULE_guardStmt               = 33
	GrammarParserRULE_breakStmt               = 34
	GrammarParserRULE_continueStmt            = 35
	GrammarParserRULE_returnStmt              = 36
	GrammarParserRULE_expr                    = 37
	GrammarParserRULE_type                    = 38
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
		p.SetState(78)
		p.Block()
	}
	{
		p.SetState(79)
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
	p.SetState(84)
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
				p.SetState(81)
				p.Stmts()
			}

		}
		p.SetState(86)
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
	StructDeclaration() IStructDeclarationContext
	StructInstance() IStructInstanceContext
	StructAccess() IStructAccessContext
	StructAssignment() IStructAssignmentContext

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

func (s *StmtsContext) StructDeclaration() IStructDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *StmtsContext) StructInstance() IStructInstanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInstanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInstanceContext)
}

func (s *StmtsContext) StructAccess() IStructAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructAccessContext)
}

func (s *StmtsContext) StructAssignment() IStructAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructAssignmentContext)
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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.VarDeclaration()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(88)
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
			p.SetState(91)
			p.VarAssignment()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(92)
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
			p.SetState(95)
			p.LetDeclaration()
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(96)
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
			p.SetState(99)
			p.VectorDeclaration()
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(100)
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
			p.SetState(103)
			p.VectorAppend()
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(104)
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
			p.SetState(107)
			p.VectorRemoveLast()
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(108)
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
			p.SetState(111)
			p.VectorRemoveAt()
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(112)
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
			p.SetState(115)
			p.FuncDeclaration()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(116)
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
			p.SetState(119)
			p.FuncCall()
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(120)
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
			p.SetState(123)
			p.PrintStmt()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(124)
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
			p.SetState(127)
			p.IfStmt()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(128)
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
			p.SetState(131)
			p.SwitchStmt()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(132)
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
			p.SetState(135)
			p.WhileStmt()
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(136)
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
			p.SetState(139)
			p.ForStmt()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(140)
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
			p.SetState(143)
			p.GuardStmt()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(144)
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
			p.SetState(147)
			p.BreakStmt()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(148)
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
			p.SetState(151)
			p.ContinueStmt()
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(152)
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
			p.SetState(155)
			p.ReturnStmt()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(156)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(159)
			p.StructDeclaration()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(160)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(163)
			p.StructInstance()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(164)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(167)
			p.StructAccess()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(168)
				p.Match(GrammarParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(171)
			p.StructAssignment()
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserSEMICOLON {
			{
				p.SetState(172)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Type_()
		}
		{
			p.SetState(181)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.expr(0)
		}

	case 2:
		localctx = NewValueVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	case 3:
		localctx = NewTypeVarDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Type_()
		}
		{
			p.SetState(192)
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
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(196)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(197)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(198)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(199)
				p.expr(0)
			}
			{
				p.SetState(200)
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
			p.SetState(204)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.expr(0)
		}

	case 2:
		localctx = NewIncrementVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(206)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(207)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(208)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(209)
				p.expr(0)
			}
			{
				p.SetState(210)
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
			p.SetState(214)
			p.Match(GrammarParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.expr(0)
		}

	case 3:
		localctx = NewDecrementVarAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(216)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(217)
				p.Match(GrammarParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(218)
				p.Match(GrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(219)
				p.expr(0)
			}
			{
				p.SetState(220)
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
			p.SetState(224)
			p.Match(GrammarParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
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
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeValueLetDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Type_()
		}
		{
			p.SetState(232)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.expr(0)
		}

	case 2:
		localctx = NewValueLetDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(GrammarParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
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

	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorValuesDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Type_()
		}
		{
			p.SetState(246)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&-1152921483568211967) != 0 {
			{
				p.SetState(249)
				p.ValuesVectorDeclaration()
			}

		}
		{
			p.SetState(252)
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
			p.SetState(254)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Type_()
		}
		{
			p.SetState(259)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
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
		p.SetState(265)
		p.expr(0)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserT__2 {
		{
			p.SetState(266)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.expr(0)
		}

		p.SetState(272)
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
		p.SetState(273)
		p.Match(GrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Type_()
	}
	{
		p.SetState(275)
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
		p.SetState(277)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Match(GrammarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.expr(0)
	}
	{
		p.SetState(281)
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
		p.SetState(283)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(GrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
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
		p.SetState(288)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(GrammarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(GrammarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.expr(0)
	}
	{
		p.SetState(294)
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

	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Type_()
		}
		{
			p.SetState(302)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GrammarParserT__0 {
			{
				p.SetState(306)
				p.ValuesMatrixDeclaration()
			}

		}
		{
			p.SetState(309)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
			p.Match(GrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Type_()
		}
		{
			p.SetState(317)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(GrammarParserEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
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
		p.SetState(324)
		p.Match(GrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&-1152921483568211967) != 0 {
		{
			p.SetState(325)
			p.ValuesVectorDeclaration()
		}

	}
	{
		p.SetState(328)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserT__2 {
		{
			p.SetState(329)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&-1152921483568211967) != 0 {
			{
				p.SetState(331)
				p.ValuesVectorDeclaration()
			}

		}
		{
			p.SetState(334)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(339)
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
	AllStructAttributes() []IStructAttributesContext
	StructAttributes(i int) IStructAttributesContext

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

func (s *StructDeclarationContext) AllStructAttributes() []IStructAttributesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructAttributesContext); ok {
			len++
		}
	}

	tst := make([]IStructAttributesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructAttributesContext); ok {
			tst[i] = t.(IStructAttributesContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclarationContext) StructAttributes(i int) IStructAttributesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAttributesContext); ok {
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

	return t.(IStructAttributesContext)
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
		p.SetState(340)
		p.Match(GrammarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(341)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserVAR || _la == GrammarParserLET {
		{
			p.SetState(343)
			p.StructAttributes()
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(349)
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

// IStructAttributesContext is an interface to support dynamic dispatch.
type IStructAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode

	// IsStructAttributesContext differentiates from other interfaces.
	IsStructAttributesContext()
}

type StructAttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructAttributesContext() *StructAttributesContext {
	var p = new(StructAttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAttributes
	return p
}

func InitEmptyStructAttributesContext(p *StructAttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structAttributes
}

func (*StructAttributesContext) IsStructAttributesContext() {}

func NewStructAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructAttributesContext {
	var p = new(StructAttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structAttributes

	return p
}

func (s *StructAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *StructAttributesContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *StructAttributesContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *StructAttributesContext) Type_() ITypeContext {
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

func (s *StructAttributesContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *StructAttributesContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *StructAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructAttributes() (localctx IStructAttributesContext) {
	localctx = NewStructAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_structAttributes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserVAR || _la == GrammarParserLET) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(352)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)
		p.Type_()
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

// IStructInstanceContext is an interface to support dynamic dispatch.
type IStructInstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AttributesCall() IAttributesCallContext

	// IsStructInstanceContext differentiates from other interfaces.
	IsStructInstanceContext()
}

type StructInstanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructInstanceContext() *StructInstanceContext {
	var p = new(StructInstanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structInstance
	return p
}

func InitEmptyStructInstanceContext(p *StructInstanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_structInstance
}

func (*StructInstanceContext) IsStructInstanceContext() {}

func NewStructInstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructInstanceContext {
	var p = new(StructInstanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_structInstance

	return p
}

func (s *StructInstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *StructInstanceContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *StructInstanceContext) AttributesCall() IAttributesCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributesCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributesCallContext)
}

func (s *StructInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructInstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructInstance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) StructInstance() (localctx IStructInstanceContext) {
	localctx = NewStructInstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_structInstance)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserID {
		{
			p.SetState(358)
			p.AttributesCall()
		}

	}
	{
		p.SetState(361)
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

// IAttributesCallContext is an interface to support dynamic dispatch.
type IAttributesCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Expr() IExprContext
	AttributesCall() IAttributesCallContext

	// IsAttributesCallContext differentiates from other interfaces.
	IsAttributesCallContext()
}

type AttributesCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesCallContext() *AttributesCallContext {
	var p = new(AttributesCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_attributesCall
	return p
}

func InitEmptyAttributesCallContext(p *AttributesCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_attributesCall
}

func (*AttributesCallContext) IsAttributesCallContext() {}

func NewAttributesCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesCallContext {
	var p = new(AttributesCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_attributesCall

	return p
}

func (s *AttributesCallContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AttributesCallContext) COLON() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOLON, 0)
}

func (s *AttributesCallContext) Expr() IExprContext {
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

func (s *AttributesCallContext) AttributesCall() IAttributesCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributesCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributesCallContext)
}

func (s *AttributesCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAttributesCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) AttributesCall() (localctx IAttributesCallContext) {
	localctx = NewAttributesCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_attributesCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(365)
		p.expr(0)
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__2 {
		{
			p.SetState(366)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.AttributesCall()
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
	p.EnterRule(localctx, 36, GrammarParserRULE_structAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(GrammarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(372)
		p.Match(GrammarParserID)
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
	p.EnterRule(localctx, 38, GrammarParserRULE_structAssignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)
		p.Match(GrammarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(GrammarParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
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
	p.EnterRule(localctx, 40, GrammarParserRULE_funcDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(GrammarParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserUNDERSCORE || _la == GrammarParserID {
		{
			p.SetState(383)
			p.ParamsFuncDeclaration()
		}

	}
	{
		p.SetState(386)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__13 {
		{
			p.SetState(387)
			p.Match(GrammarParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.Type_()
		}

	}
	{
		p.SetState(391)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Block()
	}
	{
		p.SetState(393)
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
	p.EnterRule(localctx, 42, GrammarParserRULE_paramsFuncDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(395)
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
		p.SetState(398)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserINOUT {
		{
			p.SetState(400)
			p.Match(GrammarParserINOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserINT, GrammarParserFLOAT, GrammarParserSTRING, GrammarParserBOOL, GrammarParserCHARACTER, GrammarParserID:
		{
			p.SetState(403)
			p.Type_()
		}

	case GrammarParserT__0:
		{
			p.SetState(404)
			p.VectorType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__2 {
		{
			p.SetState(407)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
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
	p.EnterRule(localctx, 44, GrammarParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&-1152920933812398079) != 0 {
		{
			p.SetState(413)
			p.ParamsFuncCall()
		}

	}
	{
		p.SetState(416)
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
	p.EnterRule(localctx, 46, GrammarParserRULE_paramsFuncCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(418)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(GrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserREFERENCE {
		{
			p.SetState(422)
			p.Match(GrammarParserREFERENCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(425)
		p.expr(0)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserT__2 {
		{
			p.SetState(426)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
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
	p.EnterRule(localctx, 48, GrammarParserRULE_printStmt)
	var _la int

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(430)
			p.Match(GrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.expr(0)
		}
		{
			p.SetState(433)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)
			p.Match(GrammarParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.expr(0)
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrammarParserT__2 {
			{
				p.SetState(438)
				p.Match(GrammarParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(439)
				p.expr(0)
			}

			p.SetState(444)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(445)
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
	p.EnterRule(localctx, 50, GrammarParserRULE_ifStmt)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(449)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.expr(0)
		}
		{
			p.SetState(451)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.Block()
		}
		{
			p.SetState(453)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.expr(0)
		}
		{
			p.SetState(457)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Block()
		}
		{
			p.SetState(459)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Block()
		}
		{
			p.SetState(463)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.expr(0)
		}
		{
			p.SetState(467)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.Block()
		}
		{
			p.SetState(469)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
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
	p.EnterRule(localctx, 52, GrammarParserRULE_switchStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.expr(0)
	}
	{
		p.SetState(477)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(478)
		p.Cases()
	}
	{
		p.SetState(479)
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
	p.EnterRule(localctx, 54, GrammarParserRULE_cases)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserCASE || _la == GrammarParserDEFAULT {
		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrammarParserCASE:
			{
				p.SetState(481)
				p.CaseStmt()
			}

		case GrammarParserDEFAULT:
			{
				p.SetState(482)
				p.DefaultStmt()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(487)
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
	p.EnterRule(localctx, 56, GrammarParserRULE_caseStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Match(GrammarParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(489)
		p.expr(0)
	}
	{
		p.SetState(490)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)
		p.Block()
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserBREAK {
		{
			p.SetState(492)
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
	p.EnterRule(localctx, 58, GrammarParserRULE_defaultStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.Match(GrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
		p.Match(GrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)
		p.Block()
	}
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrammarParserBREAK {
		{
			p.SetState(498)
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
	p.EnterRule(localctx, 60, GrammarParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)
		p.expr(0)
	}
	{
		p.SetState(503)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.Block()
	}
	{
		p.SetState(505)
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
	p.EnterRule(localctx, 62, GrammarParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(GrammarParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GrammarParserUNDERSCORE || _la == GrammarParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(509)
		p.Match(GrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(510)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(511)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(514)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(515)
		p.Block()
	}
	{
		p.SetState(516)
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
	p.EnterRule(localctx, 64, GrammarParserRULE_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.expr(0)
	}
	{
		p.SetState(519)
		p.Match(GrammarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(520)
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
	p.EnterRule(localctx, 66, GrammarParserRULE_guardStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Match(GrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(523)
		p.expr(0)
	}
	{
		p.SetState(524)
		p.Match(GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.Match(GrammarParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(526)
		p.Block()
	}
	{
		p.SetState(527)
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
	p.EnterRule(localctx, 68, GrammarParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
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
	p.EnterRule(localctx, 70, GrammarParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 72, GrammarParserRULE_returnStmt)
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(533)
			p.Match(GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(534)
			p.Match(GrammarParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
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

type StructAccessExprContext struct {
	ExprContext
}

func NewStructAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAccessExprContext {
	var p = new(StructAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAccessExprContext) StructAccess() IStructAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructAccessContext)
}

func (s *StructAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructAccessExpr(s)

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

type StructInstanceExprContext struct {
	ExprContext
}

func NewStructInstanceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInstanceExprContext {
	var p = new(StructInstanceExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructInstanceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInstanceExprContext) StructInstance() IStructInstanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructInstanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructInstanceContext)
}

func (s *StructInstanceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStructInstanceExpr(s)

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

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(577)
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
			p.SetState(539)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.expr(0)
		}
		{
			p.SetState(541)
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
			p.SetState(543)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryOperationExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__15 || _la == GrammarParserT__16) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryOperationExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(544)

			var _x = p.expr(21)

			localctx.(*UnaryOperationExprContext).right = _x
		}

	case 3:
		localctx = NewBooleanExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(545)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__29 || _la == GrammarParserT__30) {
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
			p.SetState(546)
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
			p.SetState(547)
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
			p.SetState(548)
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
			p.SetState(549)
			p.Match(GrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewCountExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(550)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(GrammarParserT__31)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewIsEmptyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(552)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(553)
			p.Match(GrammarParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewVectorAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(554)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.expr(0)
		}
		{
			p.SetState(557)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewStructInstanceExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(559)
			p.StructInstance()
		}

	case 12:
		localctx = NewStructAccessExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(560)
			p.StructAccess()
		}

	case 13:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(561)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewIntEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(562)
			p.Match(GrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.expr(0)
		}
		{
			p.SetState(565)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewFloatEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(567)
			p.Match(GrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)
			p.expr(0)
		}
		{
			p.SetState(570)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewStringEmbededExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(572)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.expr(0)
		}
		{
			p.SetState(575)
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
	p.SetState(599)
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
			p.SetState(597)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(579)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(580)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(581)

					var _x = p.expr(21)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 2:
				localctx = NewArithmeticOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ArithmeticOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(582)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(583)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ArithmeticOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__16 || _la == GrammarParserT__20) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ArithmeticOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(584)

					var _x = p.expr(20)

					localctx.(*ArithmeticOperationExprContext).right = _x
				}

			case 3:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(585)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(586)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__21 || _la == GrammarParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(587)

					var _x = p.expr(19)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 4:
				localctx = NewRelationalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RelationalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(588)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(589)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__23 || _la == GrammarParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(590)

					var _x = p.expr(18)

					localctx.(*RelationalOperationExprContext).right = _x
				}

			case 5:
				localctx = NewComparisonOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ComparisonOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(591)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(592)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__25 || _la == GrammarParserT__26) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(593)

					var _x = p.expr(17)

					localctx.(*ComparisonOperationExprContext).right = _x
				}

			case 6:
				localctx = NewLogicalOperationExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LogicalOperationExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(594)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(595)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__27 || _la == GrammarParserT__28) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(596)

					var _x = p.expr(16)

					localctx.(*LogicalOperationExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(601)
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
	ID() antlr.TerminalNode

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

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
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
	p.EnterRule(localctx, 76, GrammarParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(602)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-34)) & ^0x3f) == 0 && ((int64(1)<<(_la-34))&17179869215) != 0) {
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
	case 37:
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
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
