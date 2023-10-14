package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitUnaryOperationExpr(ctx *parser.UnaryOperationExprContext) interface{} { // Pa mientras
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	switch sign {
	case "!":
		// Agregar comentario
		v.Generator.AddComment("-----Not-----")

		// Obtener los valores izquierdo y derecho
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if rightValue.GetDataType() == BooleanType {

			// Variable temporal para el resultado
			var resultValue string

			// Generación C3D
			EV := v.Generator.NewLabel()       // L0
			EF := v.Generator.NewLabel()       // L1
			newLabel := v.Generator.NewLabel() // L2

			v.Generator.AddIf(rightValue.GetValue(), "1", "==", EV) // L0
			v.Generator.AddGoto(EF)                                 // v.Generator
			v.Generator.AddLabel(EV)                                // L0
			v.Generator.AddExpression(newTemp, "0", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(EF) // L1
			v.Generator.AddExpression(newTemp, "1", "0", "+")
			v.Generator.AddGoto(newLabel)
			v.Generator.AddLabel(newLabel) // L32

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
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipo de dato incorrecto, no se puede realizar el cambio de signo.",
			})
		}
	case "-":
		// Agregar comentario
		v.Generator.AddComment("-----Negación-----")

		// Obtener los valores izquierdo y derecho
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType {

			// Agregar la expresión a la lista de expresiones
			v.Generator.AddExpression(newTemp, "0", rightValue.GetValue(), "-")

			// Realizar la operación de multiplicación
			var resultValue string
			dataType := rightValue.GetDataType()

			resultValue = newTemp

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipo de dato incorrecto, no se puede realizar la negación.",
			})
		}
	}
	return nil
}
