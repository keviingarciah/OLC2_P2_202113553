package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitArithmeticOperationExpr(ctx *parser.ArithmeticOperationExprContext) interface{} {
	/*
		// get the left value
		leftValue := v.Visit(ctx.GetLeft())
		// get the right value
		rightValue := v.Visit(ctx.GetRight())
		// Get the operator
		sign := ctx.GetOp().GetText()
	*/
	//fmt.Println("Left: ", leftValue.GetValue(), "Right: ", rightValue.GetValue(), "Sign: ", sign)
	/*
		switch sign {
		case "%":
			// Check for division by zero
			if rightValue.GetValue() == 0 {
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Division by zero",
				})
				//fmt.Println("Error: Division by zero")

				return &NilPrimitive{Value: nil}
			}

			if leftValue.GetType() == "Int" && rightValue.GetType() == "Int" {
				return &IntPrimitive{Value: leftValue.GetValue().(int64) % rightValue.GetValue().(int64)}
			}
		case "*":
			if leftValue.GetType() == "Int" && rightValue.GetType() == "Int" {
				return &IntPrimitive{Value: leftValue.GetValue().(int64) * rightValue.GetValue().(int64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) * rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Int" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: float64(leftValue.GetValue().(int64)) * rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Int" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) * float64(rightValue.GetValue().(int64))}
			}
		case "/":
			// Check for division by zero
			if rightValue.GetValue() == 0 {
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "Division by zero",
				})
				//fmt.Println("Error: Division by zero")
				return &NilPrimitive{Value: nil}
			}

			if leftValue.GetType() == "Int" && rightValue.GetType() == "Int" {
				return &IntPrimitive{Value: leftValue.GetValue().(int64) / rightValue.GetValue().(int64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) / rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Int" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: float64(leftValue.GetValue().(int64)) / rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Int" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) / float64(rightValue.GetValue().(int64))}
			}
		case "+":
			if leftValue.GetType() == "Int" && rightValue.GetType() == "Int" {
				return &IntPrimitive{Value: leftValue.GetValue().(int64) + rightValue.GetValue().(int64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) + rightValue.GetValue().(float64)}
			} else if (leftValue.GetType() == "String" && rightValue.GetType() == "String") || (leftValue.GetType() == "Character" && rightValue.GetType() == "String") || (leftValue.GetType() == "String" && rightValue.GetType() == "Character") {
				return &StringPrimitive{Value: leftValue.GetValue().(string) + rightValue.GetValue().(string)}
			} else if leftValue.GetType() == "Int" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: float64(leftValue.GetValue().(int64)) + rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Int" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) + float64(rightValue.GetValue().(int64))}
			}
		case "-":
			if leftValue.GetType() == "Int" && rightValue.GetType() == "Int" {
				return &IntPrimitive{Value: leftValue.GetValue().(int64) - rightValue.GetValue().(int64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) - rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Int" && rightValue.GetType() == "Float" {
				return &FloatPrimitive{Value: float64(leftValue.GetValue().(int64)) - rightValue.GetValue().(float64)}
			} else if leftValue.GetType() == "Float" && rightValue.GetType() == "Int" {
				return &FloatPrimitive{Value: leftValue.GetValue().(float64) - float64(rightValue.GetValue().(int64))}
			}
		}
	*/
	return nil
}
