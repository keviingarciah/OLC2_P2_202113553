package compiler

import (
	"backend/parser"
	"backend/structures"
)

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
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

	Linicio := v.Generator.NewLabel()
	EV := v.Generator.NewLabel()
	EF := v.Generator.NewLabel()
	tree.AddDisplay(Linicio, EF, "-1", false) // Display
	v.Generator.AddComment("While")
	v.Generator.AddLabel(Linicio)

	// Agregar comentario
	v.Generator.AddComment("-----While-----")

	for ok && value {
		//fmt.Println("aca")
		result := v.Visit(ctx.Block())
		//fmt.Println("Resultado de la llamada a función:", result)

		if result != nil {
			if _, ok := result.(*BreakTransference); ok {
				// Realizar acciones relacionadas con Break
				//fmt.Println("Realizando Break")
				break
			} else if _, ok := result.(*ContinueTransference); ok {
				// Realizar acciones relacionadas con Continue
				continue
			} else if rt, ok := result.(*ReturnTransference); ok {
				// Realizar acciones relacionadas con Return
				//fmt.Println("Realizando Return:", rt.GetReturn())
				if rt.GetReturn() != nil {
					return rt.GetReturn().(PRIMITIVE)
				} else {
					return nil
				}
			}
		}
		//fmt.Println("aca: ")
		value, ok = v.Visit(ctx.Expr()).(PRIMITIVE).GetValue().(bool)
		//fmt.Println("se quiebra")
	}
	return nil
}
