package compiler

import (
	"backend/parser"
	"backend/structures"
	"strconv"
)

func (v *Visitor) VisitVectorValuesDeclaration(ctx *parser.VectorValuesDeclarationContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Declaración de Vector-----")

	vectorId := ctx.ID().GetText()
	vectorType := ctx.Type_().GetText()

	// C3D
	heapTemp := v.Generator.NewTemp()
	newTemp := v.Generator.NewTemp()

	// Recorrer los valores del vector
	if valuesCtx, ok := ctx.ValuesVectorDeclaration().(*parser.ValuesVectorDeclarationContext); ok {
		vectorSize := strconv.Itoa(len(valuesCtx.AllExpr()))

		v.Generator.AddAssign(heapTemp, "H")
		v.Generator.AddSetHeap("(int)H", vectorSize)
		v.Generator.AddExpression("H", "H", "1", "+")

		for _, value := range valuesCtx.AllExpr() {
			value := v.Visit(value).(structures.Primitive)

			v.Generator.AddSetHeap("(int)H", value.GetValue())
			v.Generator.AddExpression("H", "H", "1", "+")
		}
	} else {
		v.Generator.AddAssign(heapTemp, "H")
		v.Generator.AddSetHeap("(int)H", "0")
		v.Generator.AddExpression("H", "H", "1", "+")
	}
	// Save the value in stack
	v.Generator.AddAssign(newTemp, "P")
	v.Generator.AddExpression("P", "P", "1", "+")

	v.Generator.AddSetStack("(int)"+newTemp, heapTemp)

	// save the value in the memory
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", vectorId, vectorType, newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	// Aumentar el stack counter
	v.Generator.StackCounter++

	return nil
}

func (v *Visitor) VisitCountExpr(ctx *parser.CountExprContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Count-----")

	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		tempAdress := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(newTemp, "(int)"+tempAdress)

		return structures.Primitive{
			Value:      newTemp,
			DataType:   IntType,
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitIsEmptyExpr(ctx *parser.IsEmptyExprContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----IsEmpty-----")

	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		tempAdress := v.Generator.NewTemp()
		sizeTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(sizeTemp, "(int)"+tempAdress)

		// Validación para los booleanos
		lvl1 := v.Generator.NewLabel()
		lvl2 := v.Generator.NewLabel()

		v.Generator.AddIf(sizeTemp, "0", "!=", lvl1)
		v.Generator.AddAssign(newTemp, "1")
		v.Generator.AddGoto(lvl2)

		v.Generator.AddLabel(lvl1)
		v.Generator.AddAssign(newTemp, "0")

		v.Generator.AddLabel(lvl2)

		return structures.Primitive{
			Value:      newTemp,
			DataType:   BooleanType,
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitVectorAccessExpr(ctx *parser.VectorAccessExprContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Acceder Vector-----")

	vectorID := ctx.ID().GetText()
	index := v.Visit(ctx.Expr()).(structures.Primitive)

	if symbol, ok := v.FindSymbol(vectorID); ok {
		v.Generator.PrintBoundsError()

		tempAdress := v.Generator.NewTemp()
		sizeTemp := v.Generator.NewTemp()
		realSizeTemp := v.Generator.NewTemp()
		beginTemp := v.Generator.NewTemp()
		counterTemp := v.Generator.NewTemp()
		indexTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(sizeTemp, "(int)"+tempAdress)
		v.Generator.AddExpression(realSizeTemp, sizeTemp, "1", "-")

		lvl1 := v.Generator.NewLabel()
		lvl2 := v.Generator.NewLabel()
		lvl3 := v.Generator.NewLabel()
		lvl4 := v.Generator.NewLabel()
		lvl5 := v.Generator.NewLabel()

		v.Generator.AddIf(realSizeTemp, "0", "<", lvl1)
		v.Generator.AddIf(index.GetValue(), realSizeTemp, ">", lvl1)
		v.Generator.AddGoto(lvl2)

		v.Generator.AddLabel(lvl1)
		v.Generator.AddCall("_bounds_error_")
		v.Generator.AddAssign(newTemp, "9999999827968.00")

		v.Generator.AddGoto(lvl3)

		v.Generator.AddLabel(lvl2)

		v.Generator.AddExpression(beginTemp, tempAdress, "1", "+") // Comienzo de los valores
		v.Generator.AddAssign(counterTemp, "0")
		v.Generator.AddLabel(lvl4)
		v.Generator.AddIf(index.GetValue(), counterTemp, "==", lvl5)
		v.Generator.AddExpression(counterTemp, counterTemp, "1", "+")
		v.Generator.AddGoto(lvl4)

		v.Generator.AddLabel(lvl5)
		v.Generator.AddExpression(indexTemp, beginTemp, counterTemp, "+")
		v.Generator.AddGetHeap(newTemp, "(int)"+indexTemp)

		v.Generator.AddLabel(lvl3)

		return structures.Primitive{
			Value:      newTemp,
			DataType:   symbol.DataType,
			IsTemporal: true,
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitVectorRemoveLast(ctx *parser.VectorRemoveLastContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Remove Last-----")

	vectorID := ctx.ID().GetText()

	if symbol, ok := v.FindSymbol(vectorID); ok {
		tempAdress := v.Generator.NewTemp()
		sizeTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(sizeTemp, "(int)"+tempAdress)
		v.Generator.AddExpression(newTemp, sizeTemp, "1", "-")
		v.Generator.AddSetHeap("(int)"+tempAdress, newTemp)

	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitVectorAppend(ctx *parser.VectorAppendContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Append-----")

	vectorID := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(structures.Primitive)

	if symbol, ok := v.FindSymbol(vectorID); ok {
		tempAdress := v.Generator.NewTemp()
		sizeTemp := v.Generator.NewTemp()
		realSizeTemp := v.Generator.NewTemp() // Para iterar, osea se le resta 1
		counterTemp := v.Generator.NewTemp()
		beginTemp := v.Generator.NewTemp()
		newSizeTemp := v.Generator.NewTemp()
		heapTemp := v.Generator.NewTemp()
		valueAdressTemp := v.Generator.NewTemp()
		valueTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(sizeTemp, "(int)"+tempAdress)
		v.Generator.AddExpression(realSizeTemp, sizeTemp, "1", "-")

		v.Generator.AddExpression(beginTemp, tempAdress, "1", "+") // Comienzo de los valores
		v.Generator.AddAssign(counterTemp, "0")

		lvl1 := v.Generator.NewLabel()
		lvl2 := v.Generator.NewLabel()

		// Asignamos el nuevo tamaño
		v.Generator.AddExpression(newSizeTemp, sizeTemp, "1", "+")
		v.Generator.AddAssign(heapTemp, "H")
		v.Generator.AddSetHeap("(int)H", newSizeTemp)
		v.Generator.AddExpression("H", "H", "1", "+")

		// Iteramos para copiar los valores
		v.Generator.AddLabel(lvl1)
		v.Generator.AddIf(realSizeTemp, counterTemp, "<", lvl2)

		v.Generator.AddExpression(valueAdressTemp, beginTemp, counterTemp, "+")
		v.Generator.AddGetHeap(valueTemp, "(int)"+valueAdressTemp)

		v.Generator.AddSetHeap("(int)H", valueTemp)
		v.Generator.AddExpression("H", "H", "1", "+")

		v.Generator.AddExpression(counterTemp, counterTemp, "1", "+")
		v.Generator.AddGoto(lvl1)

		// Agregamos el nuevo valor
		v.Generator.AddLabel(lvl2)
		v.Generator.AddSetHeap("(int)H", value.GetValue())
		v.Generator.AddExpression("H", "H", "1", "+")

		// Save the value in stack
		v.Generator.AddAssign(newTemp, "P")
		v.Generator.AddExpression("P", "P", "1", "+")

		v.Generator.AddSetStack("(int)"+newTemp, heapTemp)

		// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
		symbol.Value = newTemp
		symbol.Address = strconv.Itoa(v.Generator.StackCounter)
		v.Generator.StackCounter++

		v.currentEnv.Symbols[vectorID] = symbol

		// También se debe actualizar el valor en el ambiente de origen la variable
		if env, found := v.FindSymbolEnvironment(vectorID); found {
			env.Symbols[vectorID] = symbol
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}

func (v *Visitor) VisitVectorRemoveAt(ctx *parser.VectorRemoveAtContext) interface{} {
	// Add comment
	v.Generator.AddComment("-----Remove at-----")

	vectorID := ctx.ID().GetText()
	value := v.Visit(ctx.Expr()).(structures.Primitive)

	if symbol, ok := v.FindSymbol(vectorID); ok {
		tempAdress := v.Generator.NewTemp()
		sizeTemp := v.Generator.NewTemp()
		realSizeTemp := v.Generator.NewTemp() // Para iterar, osea se le resta 1
		counterTemp := v.Generator.NewTemp()
		beginTemp := v.Generator.NewTemp()
		newSizeTemp := v.Generator.NewTemp()
		heapTemp := v.Generator.NewTemp()
		valueAdressTemp := v.Generator.NewTemp()
		valueTemp := v.Generator.NewTemp()
		newTemp := v.Generator.NewTemp()

		v.Generator.AddGetStack(tempAdress, "(int)"+symbol.Value)
		v.Generator.AddGetHeap(sizeTemp, "(int)"+tempAdress)
		v.Generator.AddExpression(realSizeTemp, sizeTemp, "2", "-")

		v.Generator.AddExpression(beginTemp, tempAdress, "1", "+") // Comienzo de los valores
		v.Generator.AddAssign(counterTemp, "0")

		lvl1 := v.Generator.NewLabel()
		lvl2 := v.Generator.NewLabel()
		lvl3 := v.Generator.NewLabel()
		lvl4 := v.Generator.NewLabel()

		// Asignamos el nuevo tamaño
		v.Generator.AddExpression(newSizeTemp, sizeTemp, "1", "-")
		v.Generator.AddAssign(heapTemp, "H")
		v.Generator.AddSetHeap("(int)H", newSizeTemp)
		v.Generator.AddExpression("H", "H", "1", "+")

		// Iteramos para copiar los valores
		v.Generator.AddLabel(lvl1)
		v.Generator.AddIf(realSizeTemp, counterTemp, "<", lvl2)
		v.Generator.AddIf(value.GetValue(), counterTemp, "==", lvl2)

		v.Generator.AddExpression(valueAdressTemp, beginTemp, counterTemp, "+")
		v.Generator.AddGetHeap(valueTemp, "(int)"+valueAdressTemp)

		v.Generator.AddSetHeap("(int)H", valueTemp)
		v.Generator.AddExpression("H", "H", "1", "+")

		v.Generator.AddExpression(counterTemp, counterTemp, "1", "+")
		v.Generator.AddGoto(lvl1)

		// Salimos del ciclo
		v.Generator.AddLabel(lvl2)

		// Asignamos nuevo valor a la contador
		v.Generator.AddExpression(counterTemp, counterTemp, "1", "+")

		// Iteramos los valores despues del valor a eliminar
		v.Generator.AddLabel(lvl3)
		v.Generator.AddIf(realSizeTemp, counterTemp, "<", lvl4)

		v.Generator.AddExpression(valueAdressTemp, beginTemp, counterTemp, "+")
		v.Generator.AddGetHeap(valueTemp, "(int)"+valueAdressTemp)

		v.Generator.AddSetHeap("(int)H", valueTemp)
		v.Generator.AddExpression("H", "H", "1", "+")

		v.Generator.AddExpression(counterTemp, counterTemp, "1", "+")
		v.Generator.AddGoto(lvl3)

		// Salimos del ciclo
		v.Generator.AddLabel(lvl4)

		// Save the value in stack
		v.Generator.AddAssign(newTemp, "P")
		v.Generator.AddExpression("P", "P", "1", "+")

		v.Generator.AddSetStack("(int)"+newTemp, heapTemp)

		// Asignamos el nuevo valor a la variable existente en el ambiente original donde se declaró
		symbol.Value = newTemp
		symbol.Address = strconv.Itoa(v.Generator.StackCounter)
		v.Generator.StackCounter++

		v.currentEnv.Symbols[vectorID] = symbol

		// También se debe actualizar el valor en el ambiente de origen la variable
		if env, found := v.FindSymbolEnvironment(vectorID); found {
			env.Symbols[vectorID] = symbol
		}
	} else {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Vector no declarada.",
		})
	}
	return nil
}
