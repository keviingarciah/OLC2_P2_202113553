package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
	"strconv"
)

func (v *Visitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	exprCount := len(ctx.AllExpr())
	for i := 0; i < exprCount; i++ {
		expression := v.Visit(ctx.Expr(i)).(structures.Primitive)

		if expression.GetDataType() == IntType || expression.GetDataType() == FloatType {
			v.Generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", expression.Value))
			v.Generator.AddPrintf("c", "10")
			v.Generator.AddBr()
		} else if expression.GetDataType() == StringType {
			//llamar a generar printstring
			v.Generator.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := v.Generator.NewTemp()
			newTemp2 := v.Generator.NewTemp()
			size := strconv.Itoa(0)
			v.Generator.AddExpression(newTemp1, "P", size, "+")              //nuevo temporal en pos vacia
			v.Generator.AddExpression(newTemp1, newTemp1, "1", "+")          //se deja espacio de retorno
			v.Generator.AddSetStack("(int)"+newTemp1, expression.GetValue()) //se coloca string en parametro que se manda
			v.Generator.AddExpression("P", "P", size, "+")                   // cambio de entorno
			v.Generator.AddCall("_print_string_")                            //Llamada
			v.Generator.AddGetStack(newTemp2, "(int)P")                      //obtencion retorno
			v.Generator.AddExpression("P", "P", size, "-")                   //regreso del entorno
			v.Generator.AddPrintf("c", "10")                                 //salto de linea
			v.Generator.AddBr()
		} else if expression.GetDataType() == CharacterType {
			v.Generator.AddPrintf("c", "(int)"+fmt.Sprintf("%v", expression.Value))
			v.Generator.AddPrintf("c", "10")
			v.Generator.AddBr()
		}
	}
	return nil
}
