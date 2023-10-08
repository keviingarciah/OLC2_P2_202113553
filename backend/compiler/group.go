package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitGroupExpr(ctx *parser.GroupExprContext) interface{} {
	expr := v.Visit(ctx.Expr())

	return expr
}
