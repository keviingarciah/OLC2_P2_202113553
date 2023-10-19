package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
	"strconv"
)

func (v *Visitor) VisitComparisonOperationExpr(ctx *parser.ComparisonOperationExprContext) interface{} {
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	switch sign {
	case "==":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Igual-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || rightValue.GetDataType() == IntType) ||
			(leftValue.GetDataType() == FloatType || rightValue.GetDataType() == FloatType) ||
			(leftValue.GetDataType() == StringType || rightValue.GetDataType() == StringType) ||
			(leftValue.GetDataType() == CharacterType || rightValue.GetDataType() == CharacterType) ||
			(leftValue.GetDataType() == BooleanType || rightValue.GetDataType() == BooleanType) {

			// Variable temporal para el resultado
			var resultValue string

			if leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType || leftValue.GetDataType() == CharacterType || leftValue.GetDataType() == BooleanType {
				// Generación C3D
				lvl1 := v.Generator.NewLabel()
				lvl2 := v.Generator.NewLabel()

				v.Generator.AddIf(leftValue.GetValue(), rightValue.GetValue(), "==", lvl1)
				v.Generator.AddAssign(newTemp, "0")
				v.Generator.AddGoto(lvl2)

				v.Generator.AddLabel(lvl1)
				v.Generator.AddAssign(newTemp, "1")

				v.Generator.AddLabel(lvl2)
				reTemp := v.Generator.NewTemp()
				v.Generator.AddAssign(reTemp, newTemp)

				// Realizar la operación relacional
				resultValue = reTemp

			} else if leftValue.GetDataType() == StringType {
				v.Generator.GenerateCompareString()

				// Generación C3D
				lvl1 := v.Generator.NewLabel()
				lvl2 := v.Generator.NewLabel()
				size := strconv.Itoa(1)
				temp1 := v.Generator.NewTemp()

				v.Generator.AddExpression(temp1, "P", fmt.Sprintf("%v", size), "+")
				v.Generator.AddExpression(temp1, temp1, "1", "+")
				v.Generator.AddSetStack("(int)"+temp1, leftValue.GetValue())
				v.Generator.AddExpression(temp1, temp1, "1", "+")
				v.Generator.AddSetStack("(int)"+temp1, rightValue.GetValue())
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "+")

				v.Generator.AddCall("_compare_string_")

				temp2 := v.Generator.NewTemp()

				v.Generator.AddGetStack(temp2, "(int)P")
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "-")

				v.Generator.AddIf(temp2, "1", "==", lvl1)
				v.Generator.AddAssign(newTemp, "0")
				v.Generator.AddGoto(lvl2)

				v.Generator.AddLabel(lvl1)
				v.Generator.AddAssign(newTemp, "1")

				v.Generator.AddLabel(lvl2)
				reTemp := v.Generator.NewTemp()
				v.Generator.AddAssign(reTemp, newTemp)

				// Realizar la operación relacional
				resultValue = reTemp
			}

			// Retornar el valor
			return structures.Primitive{
				Value:    resultValue,
				DataType: BooleanType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden comparar.",
			})
		}
	case "!=":
		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Agregar comentario
		v.Generator.AddComment("-----Diferente-----")

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || rightValue.GetDataType() == IntType) ||
			(leftValue.GetDataType() == FloatType || rightValue.GetDataType() == FloatType) ||
			(leftValue.GetDataType() == StringType || rightValue.GetDataType() == StringType) ||
			(leftValue.GetDataType() == CharacterType || rightValue.GetDataType() == CharacterType) ||
			(leftValue.GetDataType() == BooleanType || rightValue.GetDataType() == BooleanType) {

			// Variable temporal para el resultado
			var resultValue string

			if leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType || leftValue.GetDataType() == CharacterType || leftValue.GetDataType() == BooleanType {
				// Generación C3D
				lvl1 := v.Generator.NewLabel()
				lvl2 := v.Generator.NewLabel()

				v.Generator.AddIf(leftValue.GetValue(), rightValue.GetValue(), "!=", lvl1)
				v.Generator.AddAssign(newTemp, "0")
				v.Generator.AddGoto(lvl2)

				v.Generator.AddLabel(lvl1)
				v.Generator.AddAssign(newTemp, "1")

				v.Generator.AddLabel(lvl2)
				reTemp := v.Generator.NewTemp()
				v.Generator.AddAssign(reTemp, newTemp)

				// Realizar la operación relacional
				resultValue = reTemp

			} else if leftValue.GetDataType() == StringType {
				v.Generator.GenerateCompareString()

				// Generación C3D
				lvl1 := v.Generator.NewLabel()
				lvl2 := v.Generator.NewLabel()
				size := strconv.Itoa(1)
				temp1 := v.Generator.NewTemp()

				v.Generator.AddExpression(temp1, "P", fmt.Sprintf("%v", size), "+")
				v.Generator.AddExpression(temp1, temp1, "1", "+")
				v.Generator.AddSetStack("(int)"+temp1, leftValue.GetValue())
				v.Generator.AddExpression(temp1, temp1, "1", "+")
				v.Generator.AddSetStack("(int)"+temp1, rightValue.GetValue())
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "+")

				v.Generator.AddCall("_compare_string_")

				temp2 := v.Generator.NewTemp()

				v.Generator.AddGetStack(temp2, "(int)P")
				v.Generator.AddExpression("P", "P", fmt.Sprintf("%v", size), "-")

				v.Generator.AddIf(temp2, "1", "!=", lvl1)
				v.Generator.AddAssign(newTemp, "0")
				v.Generator.AddGoto(lvl2)

				v.Generator.AddLabel(lvl1)
				v.Generator.AddAssign(newTemp, "1")

				v.Generator.AddLabel(lvl2)
				reTemp := v.Generator.NewTemp()
				v.Generator.AddAssign(reTemp, newTemp)

				// Realizar la operación relacional
				resultValue = reTemp
			}

			// Retornar el valor
			return structures.Primitive{
				Value:    resultValue,
				DataType: BooleanType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipos de datos distintos, no se pueden comparar.",
			})
		}
	}
	return nil
}
