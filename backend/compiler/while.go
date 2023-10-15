package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	// Entrar a un nuevo ámbito (entorno) y salir de él cuando se termine el bloque
	v.PushEnvironment(environmentName)
	defer v.PopEnvironment()

	v.Generator.AddComment("-----While-----")
	LBegin := v.Generator.NewLabel()
	LFinal := v.Generator.NewLabel()

	v.Generator.AddLabel(LBegin)

	// Obtener la condición
	condition := v.Visit(ctx.Expr()).(structures.Primitive)

	// Verificar que la condición sea booleana
	if condition.GetDataType() != BooleanType {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "La condición debe ser de tipo Bool.",
		})
		return nil
	}

	v.Generator.AddIf("(int)"+condition.GetValue(), "0", "==", LFinal)

	v.Visit(ctx.Block())

	v.Generator.AddGoto(LBegin)
	v.Generator.AddLabel(LFinal)

	return nil
}
