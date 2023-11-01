package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
	"strconv"
	"strings"
)

// Constants
const (
	IntType       = "Int"
	FloatType     = "Float"
	StringType    = "String"
	BooleanType   = "Bool"
	CharacterType = "Character"
	NilType       = "Nil"
	VectorType    = "vector"
)

// Visit a digit primitive
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	digit := ctx.GetText() // get the digit

	// Evalue if the digit has a . to know if it is a float or an integer
	if strings.Contains(digit, ".") {
		// Agrergar comentario
		//v.Generator.AddComment("-----Int Primitivo-----")

		return structures.Primitive{
			Value:      digit,
			DataType:   FloatType,
			IsTemporal: false,
		}
	} else {
		// Agrergar comentario
		//v.Generator.AddComment("-----Float Primitivo-----")

		return structures.Primitive{
			Value:      digit,
			DataType:   IntType,
			IsTemporal: false,
		}
	}
}

func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "\"") // get the string

	// Agregar comentario
	v.Generator.AddComment("-----String Primitivo-----")

	//nuevo temporal
	newTemp := v.Generator.NewTemp()
	//iguala a heap pointer
	v.Generator.AddAssign(newTemp, "H")
	//recorremos string en ascii
	byteArray := []byte(str)
	for _, asc := range byteArray {
		//se agrega ascii al heap
		v.Generator.AddSetHeap("(int)H", strconv.Itoa(int(asc)))
		//suma heap pointer
		v.Generator.AddExpression("H", "H", "1", "+")
	}
	//caracteres de escape
	v.Generator.AddSetHeap("(int)H", "-1")
	v.Generator.AddExpression("H", "H", "1", "+")

	return structures.Primitive{
		Value:      newTemp,
		DataType:   StringType,
		IsTemporal: true,
	}
}

func (v *Visitor) VisitCharacterExpr(ctx *parser.CharacterExprContext) interface{} {
	char := strings.Replace(ctx.GetText(), "'", "", -1) // get the char

	// Get the ascii value of the char
	charASCII := strconv.Itoa(int(char[0]))

	return structures.Primitive{
		Value:      charASCII,
		DataType:   CharacterType, // Puedes definir el tipo de dato según tus necesidades
		IsTemporal: false,
	}

}

func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	boolean := ctx.GetText()

	if boolean == "true" {
		//nuevo temporal
		newTemp := v.Generator.NewTemp()
		//iguala a heap pointer
		v.Generator.AddAssign(newTemp, "1")

		return structures.Primitive{
			Value:      newTemp,
			DataType:   BooleanType,
			IsTemporal: true,
		}
	} else {
		//nuevo temporal
		newTemp := v.Generator.NewTemp()
		//iguala a heap pointer
		v.Generator.AddAssign(newTemp, "0")

		return structures.Primitive{
			Value:      newTemp,
			DataType:   BooleanType,
			IsTemporal: true,
		}
	}
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	return structures.Primitive{
		Value:      "9999999827968.00",
		DataType:   NilType,
		IsTemporal: false,
	}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText()

	if symbol, ok := v.FindSymbol(id); ok {
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(newTemp, "(int)"+symbol.Address)

		return structures.Primitive{
			Value:      newTemp,
			DataType:   symbol.DataType,
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: fmt.Sprint("Variable ", id, " no declarada."),
		})
		return nil
	}
}
