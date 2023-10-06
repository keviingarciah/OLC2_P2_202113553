package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

func (v *Visitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	exprCount := len(ctx.AllExpr())
	for i := 0; i < exprCount; i++ {
		expression := v.Visit(ctx.Expr(i)).(structures.Primitive)

		if expression.GetDataType() == IntType || expression.GetDataType() == FloatType {
			v.Generator.AddPrintf("d", "(int)"+fmt.Sprintf("%v", expression.Value))
			v.Generator.AddPrintf("c", "10")
			v.Generator.AddBr()
		}
	}
	return nil
}
