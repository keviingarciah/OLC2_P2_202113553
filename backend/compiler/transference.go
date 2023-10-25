package compiler

import (
	"backend/parser"
	"backend/structures"
)

// ----------------------------------- BREAK -----------------------------------
func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	newLabel := v.Generator.NewLabel()
	v.Generator.AddGoto(newLabel)

	v.Generator.BreakStack = append(v.Generator.BreakStack, newLabel)
	return nil
}

// ----------------------------------- CONTINUE -----------------------------------
func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	lastIndex := len(v.Generator.ContinueStack) - 1
	lastElement := v.Generator.ContinueStack[lastIndex]

	// Ir al inicio del ciclo
	v.Generator.AddGoto(lastElement)

	// Elimina el último elemento
	v.Generator.ContinueStack = v.Generator.ContinueStack[:lastIndex]

	return nil
}

// ----------------------------------- RETURN -----------------------------------
func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if ctx.Expr() != nil {
		returnValue := v.Visit(ctx.Expr()).(structures.Primitive)
		newLabel := v.Generator.NewLabel()

		v.Generator.AddSetStack("(int)P", returnValue.GetValue())
		v.Generator.AddExpression("P", "P", "1", "+")

		v.Generator.AddGoto(newLabel)

		v.Generator.ReturnStack = append(v.Generator.ReturnStack, newLabel)
		return nil
	} else {
		newLabel := v.Generator.NewLabel()
		v.Generator.AddGoto(newLabel)

		v.Generator.ReturnStack = append(v.Generator.ReturnStack, newLabel)
		return nil
	}
}
