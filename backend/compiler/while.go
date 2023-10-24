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

	// Agregar la etiqueta de inicio
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

	// Comparar la condición con 0
	v.Generator.AddIf("(int)"+condition.GetValue(), "0", "==", LFinal)

	// Guardar el label para el continue
	v.Generator.ContinueStack = append(v.Generator.ContinueStack, LBegin)

	// Visitar el bloque
	v.Visit(ctx.Block())

	// Volver a evaluar la condición
	v.Generator.AddGoto(LBegin)
	v.Generator.AddLabel(LFinal)

	// Agregar break
	if len(v.Generator.BreakStack) > 0 {
		lastIndex := len(v.Generator.BreakStack) - 1
		lastElement := v.Generator.BreakStack[lastIndex]

		// Agregar la etiqueta de final
		v.Generator.AddLabel(lastElement)

		// Elimina el último elemento
		v.Generator.BreakStack = v.Generator.BreakStack[:lastIndex]
	}

	// Quitar el label para el continue
	if len(v.Generator.ContinueStack) > 0 {
		v.Generator.ContinueStack = v.Generator.ContinueStack[:len(v.Generator.ContinueStack)-1]
	}

	return nil
}
