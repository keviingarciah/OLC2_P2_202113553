package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitIntEmbededExpr(ctx *parser.IntEmbededExprContext) interface{} {
	exprValue := v.Visit(ctx.Expr()).(structures.Primitive)

	if exprValue.GetDataType() == StringType {
		// Agregar comentario
		v.Generator.AddComment("-----Float Embebida-----")

		// Setear temporales
		inputTemp := v.Generator.NewTemp()
		outputTemp := v.Generator.NewTemp()

		v.Generator.StringToFloatInput = inputTemp
		v.Generator.StringToFloatOut = outputTemp

		// Llamar la función
		v.Generator.StringToFloat()

		v.Generator.AddAssign(inputTemp, exprValue.GetValue())

		v.Generator.AddCall("_string_to_float_")

		return structures.Primitive{
			Value:      outputTemp,
			DataType:   exprValue.GetDataType(),
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Tipo de expresión no compatible con Float",
		})
	}
	return nil
}

func (v *Visitor) VisitStringEmbededExpr(ctx *parser.StringEmbededExprContext) interface{} {
	exprValue := v.Visit(ctx.Expr()).(structures.Primitive)

	if (exprValue.GetDataType() == FloatType) || (exprValue.GetDataType() == IntType) {
		// Setear temporales
		inputTemp := v.Generator.NewTemp()
		outputTemp := v.Generator.NewTemp()

		v.Generator.NumberToStringInput = inputTemp
		v.Generator.NumberToStringOut = outputTemp

		// Llamar la función
		v.Generator.NumberToString()

		v.Generator.AddAssign(inputTemp, exprValue.GetValue())

		v.Generator.AddCall("_number_to_string_")

		return structures.Primitive{
			Value:      outputTemp,
			DataType:   exprValue.GetDataType(),
			IsTemporal: true,
		}
	} else if exprValue.GetDataType() == BooleanType {
		newTemp := v.Generator.NewTemp()
		lvl1 := v.Generator.NewLabel()
		lvl2 := v.Generator.NewLabel()

		v.Generator.AddAssign(newTemp, "H")

		v.Generator.AddIf(exprValue.GetValue(), "1", "!=", lvl1)

		v.Generator.AddSetHeap("(int)H", "116")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "114")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "117")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "101")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddGoto(lvl2)

		v.Generator.AddLabel(lvl1)
		v.Generator.AddSetHeap("(int)H", "102")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "97")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "108")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "115")
		v.Generator.AddExpression("H", "H", "1", "+")
		v.Generator.AddSetHeap("(int)H", "101")
		v.Generator.AddExpression("H", "H", "1", "+")

		v.Generator.AddLabel(lvl2)

		v.Generator.AddSetHeap("(int)H", "-1")
		v.Generator.AddExpression("H", "H", "1", "+")

		return structures.Primitive{
			Value:      newTemp,
			DataType:   StringType,
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Tipo de expresión no compatible con String",
		})
	}
	return nil
}

func (v *Visitor) VisitFloatEmbededExpr(ctx *parser.FloatEmbededExprContext) interface{} {
	exprValue := v.Visit(ctx.Expr()).(structures.Primitive)

	if exprValue.GetDataType() == StringType {
		// Agregar comentario
		v.Generator.AddComment("-----Float Embebida-----")

		// Setear temporales
		inputTemp := v.Generator.NewTemp()
		outputTemp := v.Generator.NewTemp()

		v.Generator.StringToFloatInput = inputTemp
		v.Generator.StringToFloatOut = outputTemp

		// Llamar la función
		v.Generator.StringToFloat()

		v.Generator.AddAssign(inputTemp, exprValue.GetValue())

		v.Generator.AddCall("_string_to_float_")

		return structures.Primitive{
			Value:      outputTemp,
			DataType:   exprValue.GetDataType(),
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Tipo de expresión no compatible con Float",
		})
	}
	return nil
}
