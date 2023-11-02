package compiler

import (
	"backend/parser"
)

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	/*
		// Add comment
		v.Generator.AddComment("-----Declaración de Struct-----")

		structName := ctx.ID().GetText()

		// C3D
		heapTemp := v.Generator.NewTemp()
		pointerTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		// Obtén la lista de atributos y procesa cada uno
		structAttributes := ctx.AllStructAttributes()
		for _, attributeCtx := range structAttributes {

			v.Generator.AddAssign(heapTemp, "H")
			v.Generator.AddExpression(pointerTemp, "H", "1", "+")
			v.Generator.AddSetHeap("(int)H", pointerTemp)
			v.Generator.AddExpression("H", "H", "1", "+")

			attributeName := attributeCtx.ID().GetText()
			dataType := attributeCtx.Type_().GetText()

			// Procesa el atributo (varType, attributeName, dataType)
			// ...
		}
	*/
	return nil
}
