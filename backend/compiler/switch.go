package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	/*
		// Entrar a un nuevo ámbito (entorno) para el bloque del if
		v.PushEnvironment(environmentName)
		defer v.PopEnvironment() // Salir del ámbito cuando se termine el bloque del if

		exprResult := v.Visit(ctx.Expr()).(PRIMITIVE) // Supongamos que PRIMITIVE es tu interfaz para los tipos de valores
		switchValue := exprResult.GetValue()          // Supongamos que GetValue() devuelve el valor subyacente

		var result interface{}
		var defaultCondition bool = true

		for _, caseCtx := range ctx.Cases().AllCaseStmt() {
			caseValue := v.Visit(caseCtx.Expr()).(PRIMITIVE).GetValue()
			//fmt.Printf("Comparando %v con %v\n", switchValue, caseValue)
			if switchValue == caseValue {
				defaultCondition = false

				result = v.Visit(caseCtx.Block())
				if _, ok := result.(*BreakTransference); ok {
					break
				}
			}
		}
		// Si no hay coincidencia en ningún caso, verifica el caso default
		//if result == nil {
		if defaultCondition {
			for _, defaultCtx := range ctx.Cases().AllDefaultStmt() {
				result = v.Visit(defaultCtx.Block())
				if _, ok := result.(*BreakTransference); ok {
					break
				}
			}
		}
	*/
	return nil
}
