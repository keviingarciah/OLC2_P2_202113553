package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitIntEmbededExpr(ctx *parser.IntEmbededExprContext) interface{} {
	exprValue := v.Visit(ctx.Expr()).(structures.Primitive)

	if exprValue.GetDataType() == StringType {
		// Agregar comentario
		v.Generator.AddComment("-----Int Embebida-----")

		// Llamar la función
		v.Generator.StringToNumber()

		newTemp := v.Generator.NewTemp()

		v.Generator.AddAssign(v.Generator.StringToNumberInput, exprValue.GetValue())

		v.Generator.AddCall("_string_to_number_")

		v.Generator.AddAssign(newTemp, "(int)"+v.Generator.StringToNumberOut)

		return structures.Primitive{
			Value:      newTemp,
			DataType:   IntType,
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

	// Agregar comentario
	v.Generator.AddComment("-----String Embebida-----")

	if (exprValue.GetDataType() == FloatType) || (exprValue.GetDataType() == IntType) {
		// Llamar la función
		v.Generator.NumberToString()

		v.Generator.AddAssign(v.Generator.NumberToStringInput, exprValue.GetValue())

		v.Generator.AddCall("_number_to_string_")

		return structures.Primitive{
			Value:      v.Generator.NumberToStringOut,
			DataType:   StringType,
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

		// Llamar la función
		v.Generator.StringToNumber()

		v.Generator.AddAssign(v.Generator.StringToNumberInput, exprValue.GetValue())

		v.Generator.AddCall("_string_to_number_")

		return structures.Primitive{
			Value:      v.Generator.StringToNumberOut,
			DataType:   FloatType,
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
