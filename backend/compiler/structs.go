package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

// Constants
const (
	StructType    = "Struct"
	PrimitiveType = "Primitive"
)

func (v *Visitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	structName := ctx.ID().GetText()

	// Make attributes map
	attributes := make(map[string]structures.StructAttribute)

	// Obtén la lista de atributos y procesa cada uno
	structAttributes := ctx.AllStructAttributes()
	for i, attributeCtx := range structAttributes {
		attributeName := attributeCtx.ID().GetText()
		attributeType := attributeCtx.Type_().GetText()

		var exprType string
		var primitiveType string
		var structType string

		// Evaluate the id
		if _, found := v.currentEnv.Structs[attributeType]; found {
			exprType = StructType
			structType = attributeType
		} else {
			exprType = PrimitiveType
			primitiveType = attributeType
		}

		// Save the symbol
		attributes[attributeName] = structures.StructAttribute{
			Index:         i,
			PrimitiveType: primitiveType,
			ExprType:      exprType,
			StructType:    structType,
		}
	}
	// Make the struct
	newStruct := structures.Struct{Attributes: attributes}

	// Save the struct
	v.SaveStruct(structName, newStruct)

	return nil
}

func (v *Visitor) VisitStructInstanceExpr(ctx *parser.StructInstanceExprContext) interface{} {
	return v.Visit(ctx.StructInstance())
}

func (v *Visitor) VisitStructInstance(ctx *parser.StructInstanceContext) interface{} {
	structName := ctx.ID().GetText()

	// Get the struct
	structure := v.currentEnv.Structs[structName]

	// Make attributes map
	attributes := make(map[string]structures.StructAttribute)

	if attributeCtx, ok := ctx.AttributesCall().(*parser.AttributesCallContext); ok {
		attributes = v.VisitAttributesCall(attributeCtx).(map[string]structures.StructAttribute)
	}

	// Agregar comentario
	v.Generator.AddComment("-----Instancia de Struct-----")

	// Creamos el struct
	temp1 := v.Generator.NewTemp()

	v.Generator.AddAssign(temp1, "H")
	for i, attribute := range attributes {
		// Obtenemos la posicion del heap del atributo
		temp2 := v.Generator.NewTemp()
		index := structure.Attributes[i].Index
		v.Generator.AddAssign(temp2, temp1+"+"+fmt.Sprintf("%v", index))

		// Guardamos el valor en el heap
		v.Generator.AddSetHeap("(int) "+temp2, fmt.Sprintf("%v", attribute.Value))
		v.Generator.AddExpression("H", "H", "1", "+")
	}

	return structures.Primitive{
		Value:      temp1,
		DataType:   structName,
		IsTemporal: true,
	}
}

func (v *Visitor) VisitAttributesCall(ctx *parser.AttributesCallContext) interface{} {
	attributes := make(map[string]structures.StructAttribute)

	index := 0
	v.visitAttributesCall(ctx, attributes, index)
	return attributes
}

func (v *Visitor) visitAttributesCall(ctx *parser.AttributesCallContext, params map[string]structures.StructAttribute, index int) {
	// Id of the attribute
	attributeId := ctx.ID().GetText()

	// Value of the attribute
	attributeValue := v.Visit(ctx.Expr()).(structures.Primitive)

	// Add the attribute to the map
	params[attributeId] = structures.StructAttribute{
		Index:         index,
		PrimitiveType: attributeValue.GetDataType(),
		ExprType:      PrimitiveType,
		Value:         attributeValue.GetValue(),
	}

	// Increase the index
	index++

	if ctx.AttributesCall() != nil {
		if nextParamsCtx, ok := ctx.AttributesCall().(*parser.AttributesCallContext); ok {
			v.visitAttributesCall(nextParamsCtx, params, index)
		}
	}
}

func (v *Visitor) VisitStructAccessExpr(ctx *parser.StructAccessExprContext) interface{} {
	return v.Visit(ctx.StructAccess())
}

func (v *Visitor) VisitStructAccess(ctx *parser.StructAccessContext) interface{} {
	// Get the struct
	structID := ctx.ID(0).GetText()
	attrID := ctx.ID(1).GetText()

	// Get the attribute
	if symbol, ok := v.FindSymbol(structID); ok {
		// obtenemos el struct
		structure := v.currentEnv.Structs[symbol.DataType]
		structAttribute := structure.Attributes[attrID]

		//* Obtenemos la posicion del primer struct en el heap
		PS := v.Generator.NewTemp()
		PH := v.Generator.NewTemp()
		v.Generator.AddAssign(PS, fmt.Sprintf("%v", symbol.Address))
		v.Generator.AddGetStack(PH, "(int) "+PS)

		//* Obtenemos la posicion del atributo
		attribute_index := structAttribute.Index

		//* Obtenemos la posicion en un temporal
		tPos := v.Generator.NewTemp()
		v.Generator.AddExpression(tPos, PH, fmt.Sprintf("%v", attribute_index), "+")

		//* Obtenemos el valor del atributo
		valueAtributo := v.Generator.NewTemp()
		v.Generator.AddGetHeap(valueAtributo, "(int) "+tPos)

		//* Creamos el nuevo valor
		return structures.Primitive{
			Value:      valueAtributo,
			DataType:   structAttribute.PrimitiveType,
			IsTemporal: true,
		}
	}
	return nil
}

func (v *Visitor) VisitStructAssignment(ctx *parser.StructAssignmentContext) interface{} {
	// Get the struct
	structID := ctx.ID(0).GetText()
	attrID := ctx.ID(1).GetText()

	expr := v.Visit(ctx.Expr()).(structures.Primitive)

	// Get the attribute
	if symbol, ok := v.FindSymbol(structID); ok {
		// obtenemos el struct
		structure := v.currentEnv.Structs[symbol.DataType]
		structAttribute := structure.Attributes[attrID]

		//* Obtenemos la posicion del primer struct en el heap
		PS := v.Generator.NewTemp()
		PH := v.Generator.NewTemp()
		v.Generator.AddAssign(PS, fmt.Sprintf("%v", symbol.Address))
		v.Generator.AddGetStack(PH, "(int) "+PS)

		//* Obtenemos la posicion del atributo
		attribute_index := structAttribute.Index

		//* Obtenemos la posicion en un temporal
		tPos := v.Generator.NewTemp()
		v.Generator.AddExpression(tPos, PH, fmt.Sprintf("%v", attribute_index), "+")

		// Asignamos el valor
		v.Generator.AddSetHeap("(int) "+tPos, expr.GetValue())
	}
	return nil
}
