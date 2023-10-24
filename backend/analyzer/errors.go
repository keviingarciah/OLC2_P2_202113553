package analyzer

import "github.com/antlr4-go/antlr/v4"

// ------------------ Error Handling ------------------
type LexicalError struct {
	Line    int
	Column  int
	Message string
}

type SyntaxError struct {
	Line    int
	Column  int
	Message string
}

type MyErrorListener struct {
	*antlr.DefaultErrorListener
	LexerErrors  []LexicalError
	ParserErrors []SyntaxError
}

func NewMyErrorListener() *MyErrorListener {
	return &MyErrorListener{}
}

func (e *MyErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	e.ParserErrors = append(e.ParserErrors, SyntaxError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}

func (e *MyErrorListener) SyntaxErrorUnimpl(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	e.LexerErrors = append(e.LexerErrors, LexicalError{
		Line:    line,
		Column:  column,
		Message: msg,
	})
}
