package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {
	vectorId := ctx.ID(0).GetText()
	vectorType := ctx.Type_().GetText()
	vectorAsigId := ""

	// Declarar otro vector
	if ctx.ID(1) != nil {
		//fmt.Println("Vector ID: ", ctx.ID(1).GetText())
		vectorAsigId = ctx.ID(1).GetText()

		// Verificar si el vector ya existe en el entorno actual
		if symbol, ok := v.FindSymbol(vectorAsigId); ok {
			fmt.Println(symbol)
			/*
				// Obtiene el vector
				vectorValues := symbol.Value.(*VectorPrimitive).Values

				// Actualiza el valor del vector en el entorno actual
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", vectorId, "vector", &VectorPrimitive{Values: vectorValues, VectorType: vectorType}, v.currentEnv.Name)
			*/
			return nil
		} else {
			v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Vector no declarado.",
			})
			return nil
		}
	}

	// C3D
	newTemp := v.Generator.NewTemp()

	// Declarar un vector con valores
	var vectorValues []interface{}
	if valuesCtx, ok := ctx.ValuesVectorDeclaration().(*parser.ValuesVectorDeclarationContext); ok {
		vectorValues = v.VisitValuesVectorDeclaration(valuesCtx, newTemp).([]interface{})
	}

	fmt.Println("Vector values: ", vectorValues)
	// save the value in the memory
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", vectorId, vectorType, newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	return nil
}

func (v *Visitor) VisitValuesVectorDeclaration(ctx *parser.ValuesVectorDeclarationContext, newTemp string) interface{} {
	var values []interface{}

	v.Generator.AddGetHeap(newTemp, "(int)H")
	v.Generator.AddSetHeap("(int)H", newTemp)
	v.Generator.AddExpression("H", "H", "+", "1")

	for _, exprCtx := range ctx.AllExpr() {
		value := v.Visit(exprCtx)

		v.Generator.AddSetHeap("(int)H", newTemp)
		v.Generator.AddExpression("H", "H", "+", "1")

		values = append(values, value)
	}
	//caracteres de escape
	v.Generator.AddSetHeap("(int)H", "-1")
	v.Generator.AddExpression("H", "H", "1", "+")

	return values
}

/*
func (v *Visitor) VisitVectorAppend(ctx *parser.VectorAppendContext) interface{} {
	vectorID := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())

	if symbol, ok := v.FindSymbol(vectorID); ok {
		if symbol.Line == 0 {
			// La variable no existe en ningún ambiente, mostrar un mensaje de error
			//fmt.Println("Variable no declarada:", vectorID)
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Variable no declarada: " + vectorID,
			})

			return nil
		}
		// Obtiene el vector
		vectorValues := symbol.Value.(*VectorPrimitive).Values

		// Agrega el valor al vector
		vectorValues = append(vectorValues, value)

		// Actualiza el valor del vector en el entorno actual
		symbol.Value = &VectorPrimitive{Values: vectorValues, VectorType: symbol.DataType}

		// Guardamos el nuevo valor en la memoria
		v.currentEnv.Symbols[vectorID] = symbol
	}
	return nil
}

func (v *Visitor) VisitVectorRemoveLast(ctx *parser.VectorRemoveLastContext) interface{} {
	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		if symbol.Line == 0 {
			// La variable no existe en ningún ambiente, mostrar un mensaje de error
			//fmt.Println("Variable no declarada:", vectorID)
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Variable no declarada: " + vectorID,
			})

			return nil
		}

		// Obtiene el vector
		vectorValues := symbol.Value.(*VectorPrimitive).Values

		// Verifica si el vector tiene elementos
		if len(vectorValues) > 0 {
			// Remueve el último elemento
			vectorValues = vectorValues[:len(vectorValues)-1]

			// Actualiza el valor del vector en el entorno actual
			symbol.Value = &VectorPrimitive{Values: vectorValues, VectorType: symbol.DataType}

			// Guardamos el nuevo valor en la memoria
			v.currentEnv.Symbols[vectorID] = symbol
			//fmt.Println("Vector después de eliminar el último elemento:", vectorValues)
		} else {
			//fmt.Println("El vector está vacío, no se puede eliminar el último elemento.")
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "El vector está vacío, no se puede eliminar el último elemento.",
			})

		}

	}
	return nil
}

func (v *Visitor) VisitVectorRemoveAt(ctx *parser.VectorRemoveAtContext) interface{} {
	vectorID := ctx.ID().GetText()
	index := v.Visit(ctx.Expr()).(PRIMITIVE).GetValue().(int64)

	if symbol, ok := v.FindSymbol(vectorID); ok {
		if symbol.Line == 0 {
			// La variable no existe en ningún ambiente, mostrar un mensaje de error
			//fmt.Println("Variable no declarada:", vectorID)
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Variable no declarada: " + vectorID,
			})

			return nil
		}
		// Obtiene el vector
		vectorValues := symbol.Value.(*VectorPrimitive).Values

		// Verifica si el índice está en el rango válido
		if index >= 0 && index < (int64(len(vectorValues))) {
			// Elimina el elemento en el índice especificado
			vectorValues = append(vectorValues[:index], vectorValues[index+1:]...)

			// Actualiza el valor del vector en el entorno actual
			symbol.Value = &VectorPrimitive{Values: vectorValues, VectorType: symbol.DataType}

			// Guardamos el nuevo valor en la memoria
			v.currentEnv.Symbols[vectorID] = symbol
			//fmt.Println("Vector después de eliminar el elemento en el índice", index, ":", vectorValues)
		} else {
			//fmt.Println("Índice fuera de rango. No se puede eliminar el elemento.")
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Índice fuera de rango. No se puede eliminar el elemento.",
			})

		}
	}
	return nil
}

func (v *Visitor) VisitCountExpr(ctx *parser.CountExprContext) interface{} {
	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		vector := symbol.Value.(*VectorPrimitive).Values
		return &IntPrimitive{Value: int64(len(vector))}
	}
	return nil
}

func (v *Visitor) VisitIsEmptyExpr(ctx *parser.IsEmptyExprContext) interface{} {
	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		vector := symbol.Value.(*VectorPrimitive).Values
		if len(vector) == 0 {
			return &BoolPrimitive{Value: true}
		} else {
			return &BoolPrimitive{Value: false}
		}
	}
	return nil
}

func (v *Visitor) VisitVectorAccessExpr(ctx *parser.VectorAccessExprContext) interface{} {
	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		vector := symbol.Value.(*VectorPrimitive).Values
		index := v.Visit(ctx.Expr()).(PRIMITIVE).GetValue().(int64)

		if index >= 0 && index < (int64(len(vector))) {
			return vector[index]
		} else {
			//fmt.Println("Índice fuera de rango. No se puede acceder al elemento.", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			v.SemanticErrors = append(v.SemanticErrors, SemanticError{
				Line:    ctx.GetStart().GetLine(),
				Column:  ctx.GetStart().GetColumn(),
				Message: "Índice fuera de rango. No se puede acceder al elemento.",
			})
		}
	}
	return nil
}
*/
