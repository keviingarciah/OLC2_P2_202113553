package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

// ----------------------------------- DECLARATION -----------------------------------
func (v *Visitor) VisitFuncDeclaration(ctx *parser.FuncDeclarationContext) interface{} {
	funcName := ctx.ID().GetText()

	var returnType = ""
	if ctx.Type_() != nil {
		returnType = ctx.Type_().GetText()
	}

	// Evaluate the id
	if _, found := v.currentEnv.Symbols[funcName]; found {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Función ya declarada.",
		})
		return nil
	}

	v.Generator.MainCode = false

	v.Generator.FuncCode = append(v.Generator.FuncCode, fmt.Sprintf("void %s(){\n", funcName))

	v.Visit(ctx.Block())

	v.Generator.FuncCode = append(v.Generator.FuncCode, "}\n\n")

	v.Generator.MainCode = true

	// save the value in the memory
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "function", funcName, returnType, funcName, v.currentEnv.Name, v.Generator.StackCounter)
	v.Generator.StackCounter++

	return nil
}

// ----------------------------------- CALLS -----------------------------------
func (v *Visitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	funcName := ctx.ID().GetText()

	v.Generator.AddCall(funcName)

	return nil
}
