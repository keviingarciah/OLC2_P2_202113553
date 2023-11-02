package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
	"strconv"
)

func (v *Visitor) VisitArithmeticOperationExpr(ctx *parser.ArithmeticOperationExprContext) interface{} {
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	// Operaciones aritméticas
	switch sign {
	case "%":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Módulo-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType) &&
			(rightValue.GetDataType() == IntType) {
			// Funcion
			v.Generator.PrintMathError()

			// Validación para división entre cero en C3D
			lvl1 := v.Generator.NewLabel()
			lvl2 := v.Generator.NewLabel()

			v.Generator.AddIf(rightValue.GetValue(), "0", "!=", lvl1)
			v.Generator.AddCall("_math_error_")
			//v.Generator.AddExpression(newTemp, "0", "", "")
			v.Generator.AddAssign(newTemp, "0")
			v.Generator.AddGoto(lvl2)
			v.Generator.AddLabel(lvl1)
			v.Generator.AddExpression(newTemp, "(int)"+leftValue.GetValue(), rightValue.GetValue(), "%")
			v.Generator.AddLabel(lvl2)

			// Verificar si el divisor es cero
			if rightValue.GetValue() == "0" {
				v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "El divisor no puede ser 0.",
				})

				return structures.Primitive{
					Value:    "Nil",
					DataType: NilType,
				}
			}

			// Retornar el valor
			dataType := leftValue.GetDataType()
			resultValue := newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden modular.",
			})
		}
	case "*":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Multiplicación-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {

			// Agregar la expresión a la lista de expresiones
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// Retornar el valor
			dataType := leftValue.GetDataType()
			resultValue := newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden multiplicar.",
			})
		}
	case "/":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----División-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {
			// Funcion
			v.Generator.PrintMathError()

			// Validación para división entre cero en C3D
			lvl1 := v.Generator.NewLabel()
			lvl2 := v.Generator.NewLabel()

			v.Generator.AddIf(rightValue.GetValue(), "0", "!=", lvl1)
			v.Generator.AddCall("_math_error_")
			//v.Generator.AddExpression(newTemp, "0", "", "")
			v.Generator.AddAssign(newTemp, "0")
			v.Generator.AddGoto(lvl2)
			v.Generator.AddLabel(lvl1)
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")
			v.Generator.AddLabel(lvl2)

			// Verificar si el divisor es cero
			if rightValue.GetValue() == "0" {
				v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "El divisor no puede ser 0.",
				})

				return structures.Primitive{
					Value:    "Nil",
					DataType: NilType,
				}
			}

			// Retornar el valor
			dataType := leftValue.GetDataType()
			resultValue := newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden dividir.",
			})
		}
	case "+":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Suma-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if ((leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType)) ||
			(leftValue.GetDataType() == StringType && rightValue.GetDataType() == StringType) {

			// GEneracion C3D
			if leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType {
				// Agregar la expresión a la lista de expresiones
				v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")
			} else {
				v.Generator.GenerateConcatString()

				size := strconv.Itoa(1)

				v.Generator.AddExpression(newTemp, "P", fmt.Sprintf("%v", size), "+")
				v.Generator.AddExpression(newTemp, newTemp, "1", "+")
				v.Generator.AddSetStack("(int)"+newTemp, leftValue.GetValue())
				v.Generator.AddExpression(newTemp, newTemp, "1", "+")
				v.Generator.AddSetStack("(int)"+newTemp, rightValue.GetValue())
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "+")

				v.Generator.AddCall("_concat_string_")

				newTemp = v.Generator.NewTemp()

				v.Generator.AddGetStack(newTemp, "(int)P")
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "-")
			}
			// Retornar el valor
			dataType := leftValue.GetDataType()
			resultValue := newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden sumar.",
			})
		}
	case "-":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Resta-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {

			// Agregar la expresión a la lista de expresiones
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// Retornar el valor
			dataType := leftValue.GetDataType()
			resultValue := newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden restar.",
			})
		}
	}
	return nil
}
