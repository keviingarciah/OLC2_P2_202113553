package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitGuardStmt(ctx *parser.GuardStmtContext) interface{} {
	// Entrar a un nuevo ámbito (entorno) y salir de él cuando se termine el bloque
	v.PushEnvironment(environmentName)
	defer v.PopEnvironment()

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

	// Agregar comentario
	v.Generator.AddComment("-----Guard-----")

	// C3D
	EV := v.Generator.NewLabel()
	EF := v.Generator.NewLabel()

	v.Generator.AddIf(condition.GetValue(), "1", "!=", EV)
	v.Generator.AddGoto(EF)

	v.Generator.AddLabel(EV)

	v.Visit(ctx.Block())

	v.Generator.AddLabel(EF)

	return nil
}
