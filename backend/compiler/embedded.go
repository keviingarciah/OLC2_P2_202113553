package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitIntEmbededExpr(ctx *parser.IntEmbededExprContext) interface{} {
	/*
		exprValue := v.Visit(ctx.Expr()).(PRIMITIVE) // Suponiendo que el resultado de Visit(ctx.Expr()) es de tipo PRIMITIVE
		fmt.Println(exprValue.GetType())

		switch exprValue.GetType() {
		case "String":
			intValue, err := strconv.Atoi(exprValue.GetValue().(string))
			if err != nil {
				// Manejo de error en caso de que la conversión falle
				//fmt.Println("Error: No se pudo convertir el valor de String a Int")
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se pudo convertir el valor de String a Int",
				})
				return nil
			}
			return &IntPrimitive{Value: int64(intValue)}

		case "Character":
			intValue, err := strconv.Atoi(exprValue.GetValue().(string))
			if err != nil {
				// Manejo de error en caso de que la conversión falle
				//fmt.Println("Error: No se pudo convertir el valor de String a Int")
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se pudo convertir el valor de String a Int",
				})

				return nil
			}
			return &IntPrimitive{Value: int64(intValue)}
		case "Float":
			floatValue := int(exprValue.GetValue().(float64))
			return &IntPrimitive{Value: int64(floatValue)}
		default:
			// Manejo de error para tipos de expresiones no compatibles con Int
			//fmt.Println("Error: Tipo de expresión no compatible con Int")
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipo de expresión no compatible con Int",
			})

			return nil
		}
	*/
	return nil
}

func (v *Visitor) VisitFloatEmbededExpr(ctx *parser.FloatEmbededExprContext) interface{} {
	/*
		exprValue := v.Visit(ctx.Expr()).(PRIMITIVE) // Suponiendo que el resultado de Visit(ctx.Expr()) es de tipo PRIMITIVE
		switch exprValue.GetType() {
		case "String":
			floatValue, err := strconv.ParseFloat(exprValue.GetValue().(string), 64)
			if err != nil {
				// Manejo de error en caso de que la conversión falle
				//fmt.Println("Error: No se pudo convertir el valor de String a Float")
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se pudo convertir el valor de String a Float",
				})

				return nil
			}
			return &FloatPrimitive{Value: floatValue}
		default:
			// Manejo de error para tipos de expresiones no compatibles con Float
			//fmt.Println("Error: Tipo de expresión no compatible con Float")
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Tipo de expresión no compatible con Float",
			})

			return nil
		}
	*/
	return nil
}

func (v *Visitor) VisitStringEmbededExpr(ctx *parser.StringEmbededExprContext) interface{} {
	/*
		exprValue := v.Visit(ctx.Expr()).(PRIMITIVE) // Suponiendo que el resultado de Visit(ctx.Expr()) es de tipo PRIMITIVE
		stringValue := fmt.Sprintf("%v", exprValue.GetValue())
		return &StringPrimitive{Value: stringValue}
	*/
	return nil
}
