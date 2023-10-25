package compiler

import (
	"backend/parser"
	"backend/structures"
)

// ----------------------------------- VAR -----------------------------------
func (v *Visitor) VisitValueVarAssignment(ctx *parser.ValueVarAssignmentContext) interface{} {
	varId := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(varId); ok {
		if ctx.Expr(1) == nil {
			expression := v.Visit(ctx.Expr(0)).(structures.Primitive)

			if symbol.Type == "var" {
				// Verificamos si los tipos son compatibles antes de asignar el nuevo valor
				if expression.GetDataType() == symbol.DataType {
					// C3D
					v.Generator.AddComment("-----Asignación de variable-----")

					v.Generator.AddSetStack("(int)"+symbol.Address, expression.GetValue())
					v.Generator.AddBr()

					// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
					symbol.Value = expression.GetValue()
					v.currentEnv.Symbols[varId] = symbol

					// También se debe actualizar el valor en el ambiente de origen la variable
					if env, found := v.FindSymbolEnvironment(varId); found {
						env.Symbols[varId] = symbol
					}
				} else {
					v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "Tipos de variables incompatibles.",
					})
				}
			} else {
				v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "La variable no es modificable, es una constante.",
				})
			}
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitIncrementVarAssignment(ctx *parser.IncrementVarAssignmentContext) interface{} {
	varId := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(varId); ok {
		if ctx.Expr(1) == nil {
			expression := v.Visit(ctx.Expr(0)).(structures.Primitive)

			if symbol.Type == "var" {
				// Verificamos si los tipos son compatibles antes de asignar el nuevo valor
				if expression.GetDataType() == symbol.DataType {
					// C3D
					v.Generator.AddComment("-----Incremento de variable-----")

					valueTemp := v.Generator.NewTemp()
					newTemp := v.Generator.NewTemp()

					v.Generator.AddAssign(valueTemp, symbol.Value)
					v.Generator.AddExpression(newTemp, valueTemp, expression.GetValue(), "+")
					v.Generator.AddSetStack("(int)"+symbol.Address, newTemp)
					v.Generator.AddBr()

					// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
					symbol.Value = expression.GetValue()
					v.currentEnv.Symbols[varId] = symbol

					// También se debe actualizar el valor en el ambiente de origen la variable
					if env, found := v.FindSymbolEnvironment(varId); found {
						env.Symbols[varId] = symbol
					}
				} else {
					v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "Tipos de variables incompatibles.",
					})
				}
			} else {
				v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "La variable no es modificable, es una constante.",
				})
			}
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitDecrementVarAssignment(ctx *parser.DecrementVarAssignmentContext) interface{} {
	varId := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(varId); ok {
		if ctx.Expr(1) == nil {
			expression := v.Visit(ctx.Expr(0)).(structures.Primitive)

			if symbol.Type == "var" {
				// Verificamos si los tipos son compatibles antes de asignar el nuevo valor
				if expression.GetDataType() == symbol.DataType {
					// C3D
					v.Generator.AddComment("-----Incremento de variable-----")

					valueTemp := v.Generator.NewTemp()
					newTemp := v.Generator.NewTemp()

					v.Generator.AddAssign(valueTemp, symbol.Value)
					v.Generator.AddExpression(newTemp, valueTemp, expression.GetValue(), "-")
					v.Generator.AddSetStack("(int)"+symbol.Address, newTemp)
					v.Generator.AddBr()

					// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
					symbol.Value = expression.GetValue()
					v.currentEnv.Symbols[varId] = symbol

					// También se debe actualizar el valor en el ambiente de origen la variable
					if env, found := v.FindSymbolEnvironment(varId); found {
						env.Symbols[varId] = symbol
					}
				} else {
					v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "Tipos de variables incompatibles.",
					})
				}
			} else {
				v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "La variable no es modificable, es una constante.",
				})
			}
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable no declarada.",
		})
	}
	return nil
}
