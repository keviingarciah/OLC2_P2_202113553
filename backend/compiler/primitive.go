package compiler

import (
	"backend/parser"
	"backend/structures"
	"strings"
)

// Constants
const (
	IntType       = "Int"
	FloatType     = "Float"
	StringType    = "String"
	BooleanType   = "Bool"
	CharacterType = "Character"
	NilType       = "nil"
)

// Visit a digit primitive
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	digit := ctx.GetText() // get the digit

	// Evalue if the digit has a . to know if it is a float or an integer
	if strings.Contains(digit, ".") {
		return structures.Primitive{
			Value:      digit,
			DataType:   FloatType,
			IsTemporal: false,
		}
	} else {
		return structures.Primitive{
			Value:      digit,
			DataType:   FloatType,
			IsTemporal: false,
		}
	}
}

func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "\"") // get the string
	if len(str) == 1 {
		//fmt.Println("Character detected: ", str)
		return nil
	} else {
		//fmt.Println("Primitive String: ", str)
		return nil
	}
}

func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	boolean := ctx.GetText() // get the digit

	return structures.Primitive{
		Value:      boolean,
		DataType:   BooleanType,
		IsTemporal: false,
	}
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	value := ctx.GetText() // get the nil

	return structures.Primitive{
		Value:      value,
		DataType:   NilType,
		IsTemporal: false,
	}
}

/*
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()
	if symbol, ok := v.FindSymbol(id); ok {
		return symbol.Value
	} else {
		//fmt.Println("no such variable: " + id)
		v.SemanticErrors = append(v.SemanticErrors, SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable no declarada: " + id,
		})
	}
	return nil
}
*/
