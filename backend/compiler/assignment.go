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
					// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
					symbol.Value = expression.GetValue()
					v.currentEnv.Symbols[varId] = symbol

					// También se debe actualizar el valor en el ambiente de origen la variable
					if env, found := v.FindSymbolEnvironment(varId); found {
						env.Symbols[varId] = symbol
					}

					// C3D
					v.Generator.AddComment("-----Asignación de variable-----")

					v.Generator.AddSetStack("(int)"+symbol.Address, expression.GetValue())
					v.Generator.AddBr()

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

/*
func (v *Visitor) VisitIncrementVarAssignment(ctx *parser.IncrementVarAssignmentContext) interface{} {
	varId := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(varId); ok {
		if ctx.Expr(1) == nil {
			exprValue := v.Visit(ctx.Expr(0)).(PRIMITIVE)
			// Verificamos si los tipos son compatibles
			if exprValue.GetType() == symbol.DataType || (exprValue.GetType() == "Character" && symbol.DataType == "String" || (exprValue.GetType() == "String" && symbol.DataType == "Character")) {
				if exprValue.GetType() == "Int" {
					// Incrementamos el valor del primitivo con la expresión
					newValue := symbol.Value.(PRIMITIVE).GetValue().(int64) + exprValue.GetValue().(int64)
					symbol.Value = &IntPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura IntPrimitive
				} else if exprValue.GetType() == "Float" {
					// Incrementamos el valor del primitivo con la expresión
					newValue := symbol.Value.(PRIMITIVE).GetValue().(float64) + exprValue.GetValue().(float64)
					symbol.Value = &FloatPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
				} else if exprValue.GetType() == "String" || exprValue.GetType() == "Character" {
					// Incrementamos el valor del primitivo con la expresión
					newValue := symbol.Value.(PRIMITIVE).GetValue().(string) + exprValue.GetValue().(string)
					symbol.Value = &StringPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
				}

				v.currentEnv.Symbols[varId] = symbol
				if env, found := v.FindSymbolEnvironment(varId); found {
					//fmt.Println("Encontrado en el ambiente:", env.Name)
					env.Symbols[varId] = symbol
				}
			} else {
				// Los tipos no son compatibles o no son int, mostrar un mensaje de error o manejarlo según tu necesidad
				// fmt.Println("Tipos de variables incompatibles incremento, Linea ", ctx.GetStart().GetLine(), " Columna ", ctx.GetStart().GetColumn())
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Tipos de variables incompatibles",
				})
			}
		} else {
			vector := symbol.Value.(*VectorPrimitive).Values
			index := v.Visit(ctx.Expr(0)).(PRIMITIVE).GetValue().(int64)

			if index >= 0 && index < (int64(len(vector))) {
				exprValue := v.Visit(ctx.Expr(1)).(PRIMITIVE)
				vectorValue := vector[index].(PRIMITIVE)

				if exprValue.GetType() == vectorValue.GetType() {
					if exprValue.GetType() == "Int" {
						// Incrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(int64) + exprValue.GetValue().(int64)
						vector[index] = &IntPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura IntPrimitive
					} else if exprValue.GetType() == "Float" {
						// Incrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(float64) + exprValue.GetValue().(float64)
						vector[index] = &FloatPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
					} else if exprValue.GetType() == "String" {
						// Incrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(string) + exprValue.GetValue().(string)
						vector[index] = &StringPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
					} else {
						// Incrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(string) + exprValue.GetValue().(string)
						vector[index] = &CharPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
					}
					// Actualiza el valor del vector en el entorno actual
					symbol.Value = &VectorPrimitive{Values: vector, VectorType: symbol.DataType}

					// Guardamos el nuevo valor en la memoria
					v.currentEnv.Symbols[varId] = symbol
				} else {
					// Los tipos no son compatibles o no son int, mostrar un mensaje de error o manejarlo según tu necesidad
					//fmt.Println("Tipos de variables incompatibles, Linea ", ctx.GetStart().GetLine(), " Columna ", ctx.GetStart().GetColumn())
					v.SemanticErrors = append(v.SemanticErrors, SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "Tipos de variables incompatibles",
					})
				}
			} else {
				//fmt.Println("Índice fuera de rango. No se puede acceder al elemento incremento.")
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Índice fuera de rango",
				})
			}
		}
	}
	return nil
}

func (v *Visitor) VisitDecrementVarAssignment(ctx *parser.DecrementVarAssignmentContext) interface{} {
	varId := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(varId); ok {
		if ctx.Expr(1) == nil {
			exprValue := v.Visit(ctx.Expr(0)).(PRIMITIVE)

			// Verificamos si los tipos son compatibles
			if exprValue.GetType() == symbol.DataType {
				if exprValue.GetType() == "Int" {
					// Decrementamos el valor del primitivo con la expresión
					newValue := symbol.Value.(PRIMITIVE).GetValue().(int64) - exprValue.GetValue().(int64)
					symbol.Value = &IntPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura IntPrimitive
				} else if exprValue.GetType() == "Float" {
					// Decrementamos el valor del primitivo con la expresión
					newValue := symbol.Value.(PRIMITIVE).GetValue().(float64) - exprValue.GetValue().(float64)
					symbol.Value = &FloatPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
				}

				v.currentEnv.Symbols[varId] = symbol
				if env, found := v.FindSymbolEnvironment(varId); found {
					//fmt.Println("Encontrado en el ambiente:", env.Name)
					env.Symbols[varId] = symbol
				}
			} else {
				// Los tipos no son compatibles o no son int, mostrar un mensaje de error o manejarlo según tu necesidad
				//fmt.Println("Tipos de variables incompatibles, Linea ", ctx.GetStart().GetLine(), " Columna ", ctx.GetStart().GetColumn())
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Tipos de variables incompatibles",
				})

			}
		} else {
			vector := symbol.Value.(*VectorPrimitive).Values
			index := v.Visit(ctx.Expr(0)).(PRIMITIVE).GetValue().(int64)

			if index >= 0 && index < (int64(len(vector))) {
				exprValue := v.Visit(ctx.Expr(1)).(PRIMITIVE)
				vectorValue := vector[index].(PRIMITIVE)

				if exprValue.GetType() == vectorValue.GetType() {
					if exprValue.GetType() == "Int" {
						// Decrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(int64) - exprValue.GetValue().(int64)
						vector[index] = &IntPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura IntPrimitive
					} else if exprValue.GetType() == "Float" {
						// Decrementamos el valor del primitivo con la expresión
						newValue := vectorValue.GetValue().(float64) - exprValue.GetValue().(float64)
						vector[index] = &FloatPrimitive{Value: newValue} // Asignamos el nuevo valor a la estructura FloatPrimitive
					}
					// Actualiza el valor del vector en el entorno actual
					symbol.Value = &VectorPrimitive{Values: vector, VectorType: symbol.DataType}

					// Guardamos el nuevo valor en la memoria
					v.currentEnv.Symbols[varId] = symbol
				} else {
					// Los tipos no son compatibles o no son int, mostrar un mensaje de error o manejarlo según tu necesidad
					//fmt.Println("Tipos de variables incompatibles, Linea ", ctx.GetStart().GetLine(), " Columna ", ctx.GetStart().GetColumn())
					v.SemanticErrors = append(v.SemanticErrors, SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "Tipos de variables incompatibles",
					})
				}
			} else {
				//fmt.Println("Índice fuera de rango. No se puede acceder al elemento dec.")
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Índice fuera de rango",
				})
			}
		}
	}
	return nil
}
*/
