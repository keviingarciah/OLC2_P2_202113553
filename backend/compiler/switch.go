package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	// Entrar a un nuevo ámbito (entorno) y salir de él cuando se termine el bloque
	v.PushEnvironment(environmentName)
	defer v.PopEnvironment()

	labelsList := []string{} // Lista de labels

	switchValue := v.Visit(ctx.Expr()).(structures.Primitive).GetValue()

	v.Generator.AddComment("-------Switch-------")

	// Agregar comparacion
	for _, caseCtx := range ctx.Cases().AllCaseStmt() {
		// C3D
		lCase := v.Generator.NewLabel()
		labelsList = append(labelsList, lCase)

		caseValue := v.Visit(caseCtx.Expr()).(structures.Primitive).GetValue()

		v.Generator.AddIf("(int)"+switchValue, "(int)"+caseValue, "==", lCase)
	}

	// Default
	lDefault := v.Generator.NewLabel()

	v.Generator.AddGoto(lDefault)

	// Final
	lFinal := v.Generator.NewLabel()

	// Visitar los casos
	for i, caseCtx := range ctx.Cases().AllCaseStmt() {
		lCase := labelsList[i]

		v.Generator.AddComment("-------Case-------")

		v.Generator.AddLabel(lCase)

		v.Visit(caseCtx.Block())

		v.Generator.AddGoto(lFinal)
	}

	for _, defaultCtx := range ctx.Cases().AllDefaultStmt() {
		// C3D
		v.Generator.AddComment("-------Default-------")

		v.Generator.AddLabel(lDefault)

		v.Visit(defaultCtx.Block())

		v.Generator.AddGoto(lFinal)
	}

	// Visitar el final
	v.Generator.AddLabel(lFinal)

	return nil
}
