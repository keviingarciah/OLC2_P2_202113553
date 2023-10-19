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
			// Agregar comentario
			v.Generator.AddComment("-------Print Entero o Float-------")

			v.Generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", expression.GetValue()))
			v.Generator.AddPrintf("c", "32")
			//v.Generator.AddBr()
		} else if expression.GetDataType() == StringType {
			// Agregar comentario
			v.Generator.AddComment("-------Print String-------")

			//llamar a generar printstring
			v.Generator.PrintString()
			//agregar codigo en el main
			newTemp1 := v.Generator.NewTemp()
			newTemp2 := v.Generator.NewTemp()
			size := strconv.Itoa(1)

			v.Generator.AddExpression(newTemp1, "P", size, "+")              //nuevo temporal en pos vacia
			v.Generator.AddExpression(newTemp1, newTemp1, "1", "+")          //se deja espacio de retorno
			v.Generator.AddSetStack("(int)"+newTemp1, expression.GetValue()) //se coloca string en parametro que se manda
			v.Generator.AddExpression("P", "P", size, "+")                   // cambio de entorno
			v.Generator.AddCall("_print_string_")                            //Llamada
			v.Generator.AddGetStack(newTemp2, "(int)P")                      //obtencion retorno
			v.Generator.AddExpression("P", "P", size, "-")                   //regreso del entorno
			v.Generator.AddPrintf("c", "32")
			//v.Generator.AddBr()
		} else if expression.GetDataType() == CharacterType {
			// Agregar comentario
			v.Generator.AddComment("-------Print Character-------")

			v.Generator.AddPrintf("c", "(int)"+fmt.Sprintf("%v", expression.GetValue()))
			v.Generator.AddPrintf("c", "32")
			//v.Generator.AddBr()
		} else if expression.GetDataType() == BooleanType {
			// Agregar comentario
			v.Generator.AddComment("-------Print Bool-------")

			// Validación para los booleanos
			lvl1 := v.Generator.NewLabel()
			lvl2 := v.Generator.NewLabel()

			v.Generator.AddIf(expression.GetValue(), "1", "!=", lvl1)
			v.Generator.AddPrintf("c", "116")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "117")
			v.Generator.AddPrintf("c", "101")
			v.Generator.AddGoto(lvl2)
			v.Generator.AddLabel(lvl1)

			v.Generator.AddPrintf("c", "102")
			v.Generator.AddPrintf("c", "97")
			v.Generator.AddPrintf("c", "108")
			v.Generator.AddPrintf("c", "115")
			v.Generator.AddPrintf("c", "101")
			v.Generator.AddLabel(lvl2)

			v.Generator.AddPrintf("c", "32")
		}
	}
	// Salto de linea
	v.Generator.AddPrintf("c", "10")
	return nil
}
