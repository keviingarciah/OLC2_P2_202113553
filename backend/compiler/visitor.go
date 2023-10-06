package compiler

import (
	"backend/generator"
	"backend/parser"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor Struct
type Visitor struct {
	parser.BaseGrammarVisitor
	Generator generator.Generator
}

// Create a new visitor
func NewVisitor() *Visitor {
	return &Visitor{
		Generator: generator.NewGenerator(),
	}
}

// Visit
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		node := tree.Accept(v)
		return node
	}
}

// Start the visitor
func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.Visit(ctx.Block())
}

// Visit Block
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// Visit stmts
	for _, stmt := range ctx.AllStmts() {
		v.Visit(stmt)
	}
	return nil
}

// Visit the stmts
func (v *Visitor) VisitStmts(ctx *parser.StmtsContext) interface{} {
	// Verify if the stmt which stmt is
	if ctx.VarDeclaration() != nil {
		return v.Visit(ctx.VarDeclaration())

	} else if ctx.VarAssignment() != nil {
		return v.Visit(ctx.VarAssignment())

	} else if ctx.LetDeclaration() != nil {
		return v.Visit(ctx.LetDeclaration())

	} else if ctx.VectorDeclaration() != nil {
		return v.Visit(ctx.VectorDeclaration())

	} else if ctx.VectorAppend() != nil {
		return v.Visit(ctx.VectorAppend())

	} else if ctx.VectorRemoveLast() != nil {
		return v.Visit(ctx.VectorRemoveLast())

	} else if ctx.VectorRemoveAt() != nil {
		return v.Visit(ctx.VectorRemoveAt())

	} else if ctx.FuncDeclaration() != nil {
		return v.Visit(ctx.FuncDeclaration())

	} else if ctx.FuncCall() != nil {
		return v.Visit(ctx.FuncCall())

	} else if ctx.PrintStmt() != nil {
		return v.Visit(ctx.PrintStmt())

	} else if ctx.IfStmt() != nil {
		return v.Visit(ctx.IfStmt())

	} else if ctx.SwitchStmt() != nil {
		return v.Visit(ctx.SwitchStmt())

	} else if ctx.WhileStmt() != nil {

		return v.Visit(ctx.WhileStmt())
	} else if ctx.ForStmt() != nil {
		return v.Visit(ctx.ForStmt())

	} else if ctx.GuardStmt() != nil {
		return v.Visit(ctx.GuardStmt())

	} else if ctx.BreakStmt() != nil {
		return v.Visit(ctx.BreakStmt())

	} else if ctx.ContinueStmt() != nil {
		return v.Visit(ctx.ContinueStmt())

	} else if ctx.ReturnStmt() != nil {
		return v.Visit(ctx.ReturnStmt())
	}
	return nil
}
