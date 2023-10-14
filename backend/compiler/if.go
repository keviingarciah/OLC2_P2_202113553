package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
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
	v.Generator.AddComment("-----If-----")

	// "ifstmt"
	EV := v.Generator.NewLabel()
	EF := v.Generator.NewLabel()
	newLabel := v.Generator.NewLabel()

	v.Generator.AddIf(condition.GetValue(), "1", "==", EV)
	v.Generator.AddGoto(EF)

	v.Generator.AddLabel(EV)

	v.Visit(ctx.Block(0))

	v.Generator.AddGoto(newLabel)

	v.Generator.AddLabel(EF)

	if ctx.ELSE() != nil {
		// Si hay un "else", verifica si es seguido por otro "ifstmt" o un bloque
		if ctx.IfStmt() != nil {
			// Si es un "else if", visita el contexto "ifstmt"

			// Agregar comentario
			v.Generator.AddComment("-----Else If-----")
			v.Visit(ctx.IfStmt())
		} else if ctx.Block(1) != nil {
			// Si es un "else" tradicional, visita el bloque del "else"

			// Agregar comentario
			v.Generator.AddComment("-----Else-----")
			v.Visit(ctx.Block(1))
		}
	}

	v.Generator.AddGoto(newLabel)

	v.Generator.AddLabel(newLabel)

	return nil
}
