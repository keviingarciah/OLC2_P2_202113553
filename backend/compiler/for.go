package compiler

import (
	"backend/parser"
	"backend/structures"
	"strconv"
)

func (v *Visitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	// Entrar a un nuevo ámbito (entorno) y salir de él cuando se termine el bloque
	v.PushEnvironment(environmentName)
	defer v.PopEnvironment()

	// Get variable name
	var variableName string

	// If the variable name is not specified, use _
	if ctx.UNDERSCORE() != nil {
		variableName = "_"
	} else {
		variableName = ctx.ID().GetText()
	}

	// If in range
	inRange := ctx.Range_() != nil

	// Start and end values
	var start, end structures.Primitive

	// Type of the for loop
	if inRange {
		// Get start and end values
		start = v.Visit(ctx.Range_().Expr(0)).(structures.Primitive)
		end = v.Visit(ctx.Range_().Expr(1)).(structures.Primitive)

		// Save symbol address
		address := strconv.Itoa(v.Generator.StackCounter)

		// C3D
		l1 := v.Generator.NewLabel()
		l2 := v.Generator.NewLabel()
		l3 := v.Generator.NewLabel()
		t1 := v.Generator.NewTemp()
		t2 := v.Generator.NewTemp()

		// Agregar comentario
		v.Generator.AddComment("-------Declaracion de Variable-------")

		v.Generator.AddSetStack("(int)P", "0")
		v.Generator.AddExpression("P", "P", "1", "+")

		// Save the value in the symbol table (env and global)
		v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", variableName, IntType, "0", v.currentEnv.Name, v.Generator.StackCounter)

		// Aumentar el stack counter
		v.Generator.StackCounter++

		// Agregar comentario
		v.Generator.AddComment("-------For-------")

		v.Generator.AddSetStack("(int)"+address, start.GetValue())

		v.Generator.AddLabel(l1)

		v.Generator.AddGetStack(t1, "(int)"+address)

		v.Generator.AddIf("(int)"+t1, "(int)"+end.GetValue(), ">", l3)

		v.Visit(ctx.Block())

		v.Generator.AddComment("-------Actualización-------")

		v.Generator.AddLabel(l2)

		v.Generator.AddGetStack(t1, "(int)"+address)

		v.Generator.AddExpression(t2, t1, "1", "+")

		v.Generator.AddSetStack("(int)"+address, t2)

		v.Generator.AddGoto(l1)

		v.Generator.AddLabel(l3)

		/*
			// Evaluar y ejecutar el bucle for similar a Python
			for i := startInt; i <= endInt; i++ {

				if symbol, ok := v.FindSymbol(variableName); ok {
					// Verificamos si los tipos son compatibles antes de asignar el nuevo valor
					if IntType == symbol.DataType {
						// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
						symbol.Value = strconv.Itoa(i)
						v.currentEnv.Symbols[variableName] = symbol

						// También se debe actualizar el valor en el ambiente de origen la variable
						if env, found := v.FindSymbolEnvironment(variableName); found {
							env.Symbols[variableName] = symbol
						}
					} else {
						v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
							Line:    ctx.GetStart().GetLine(),
							Column:  ctx.GetStart().GetColumn(),
							Message: "Tipos de variables incompatibles.",
						})
					}
				}
			}
		*/
	}
	return nil
}
