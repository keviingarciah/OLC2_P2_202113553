package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	/*
		value, ok := v.Visit(ctx.Expr()).(structures.Primitive)

		if ok && !value {
			return v.Visit(ctx.Block(0))
		} else {
			if ctx.ELSE() != nil {
				// Si hay un "else", verifica si es seguido por otro "ifstmt" o un bloque
				if ctx.IfStmt() != nil {
					// Si es un "else if", visita el contexto "ifstmt"
					return v.Visit(ctx.IfStmt())
				} else if ctx.Block(1) != nil {
					// Si es un "else" tradicional, visita el bloque del "else"
					return v.Visit(ctx.Block(1))
				}
			}
		}
	*/
	return nil
}
