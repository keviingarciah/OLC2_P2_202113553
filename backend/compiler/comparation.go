package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

func (v *Visitor) VisitComparisonOperationExpr(ctx *parser.ComparisonOperationExprContext) interface{} {
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	switch sign {
	case "==":
		// Agregar comentario
		v.Generator.AddComment("-----Igual-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

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
				// Se obtiene el tamaño de la cadena 💀
			}

			// Retornar el valor
			return structures.Primitive{
				Value:    resultValue,
				DataType: BooleanType,
			}
		} else {
			fmt.Print("ERROR: No se puede comparar.")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	case "!=":
		// Agregar comentario
		v.Generator.AddComment("-----Diferente-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

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
				// Se obtiene el tamaño de la cadena 💀
			}

			// Retornar el valor
			return structures.Primitive{
				Value:    resultValue,
				DataType: BooleanType,
			}
		} else {
			fmt.Print("ERROR: No se puede comparar.")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	}
	return nil
}
