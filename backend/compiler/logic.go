package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

func (v *Visitor) VisitLogicalOperationExpr(ctx *parser.LogicalOperationExprContext) interface{} {
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	switch sign {
	case "&&":
		// Agregar comentario
		v.Generator.AddComment("-----And-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if leftValue.GetDataType() == BooleanType || rightValue.GetDataType() == BooleanType {

			// Variable temporal para el resultado
			var resultValue string

			// Generación C3D
			EV1 := v.Generator.NewLabel()      // L0
			EF1 := v.Generator.NewLabel()      // L1
			EV2 := v.Generator.NewLabel()      // L2
			EF2 := v.Generator.NewLabel()      // L3
			newLabel := v.Generator.NewLabel() // L4

			v.Generator.AddIf(leftValue.GetValue(), "1", "==", EV1)  // L0
			v.Generator.AddGoto(EF1)                                 //v.Generator.
			v.Generator.AddLabel(EV1)                                //v.Generator.
			v.Generator.AddIf(rightValue.GetValue(), "1", "==", EV2) // L2
			v.Generator.AddGoto(EF2)                                 //v.Generator.
			v.Generator.AddLabel(EV2)                                // L1
			v.Generator.AddExpression(newTemp, "1", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(EF1) // L1
			v.Generator.AddExpression(newTemp, "0", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(EF2) // L3
			v.Generator.AddExpression(newTemp, "0", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(newLabel) // L3

			reTemp := v.Generator.NewTemp()
			v.Generator.AddAssign(reTemp, newTemp)

			// Realizar la operación relacional
			resultValue = reTemp

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
	case "||":
		// Agregar comentario
		v.Generator.AddComment("-----Or-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if leftValue.GetDataType() == BooleanType || rightValue.GetDataType() == BooleanType {

			// Variable temporal para el resultado
			var resultValue string

			// Generación C3D
			EV1 := v.Generator.NewLabel()      // L0
			EF1 := v.Generator.NewLabel()      // L1
			EV2 := v.Generator.NewLabel()      // L2
			EF2 := v.Generator.NewLabel()      // L3
			newLabel := v.Generator.NewLabel() // L4

			v.Generator.AddIf(leftValue.GetValue(), "1", "==", EV1)  // L0
			v.Generator.AddGoto(EF1)                                 // v.Generator
			v.Generator.AddLabel(EF1)                                // v.Generator
			v.Generator.AddIf(rightValue.GetValue(), "1", "==", EV2) // L2
			v.Generator.AddGoto(EF2)                                 // v.Generator
			v.Generator.AddLabel(EV1)                                // L0
			v.Generator.AddExpression(newTemp, "1", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(EV2) // L2
			v.Generator.AddExpression(newTemp, "1", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(EF2) // L3
			v.Generator.AddExpression(newTemp, "0", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(newLabel) // L3

			reTemp := v.Generator.NewTemp()
			v.Generator.AddAssign(reTemp, newTemp)

			// Realizar la operación relacional
			resultValue = reTemp

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
