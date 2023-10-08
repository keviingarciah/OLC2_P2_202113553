package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
	"strconv"
)

func (v *Visitor) VisitArithmeticOperationExpr(ctx *parser.ArithmeticOperationExprContext) interface{} {
	// Get the operator
	sign := ctx.GetOp().GetText()

	// Crear un nuevo temporal
	newTemp := v.Generator.NewTemp()

	// Operaciones aritméticas
	switch sign {
	case "%":
		// Agregar comentario
		v.Generator.AddComment("-----Módulo-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType) &&
			(rightValue.GetDataType() == IntType) {

			// Validación para división entre cero en C3D
			lvl1 := v.Generator.NewLabel()
			lvl2 := v.Generator.NewLabel()

			v.Generator.AddIf(rightValue.GetValue(), "0", "!=", lvl1)
			v.Generator.AddPrintf("c", "77")
			v.Generator.AddPrintf("c", "97")
			v.Generator.AddPrintf("c", "116")
			v.Generator.AddPrintf("c", "104")
			v.Generator.AddPrintf("c", "69")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "111")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "10")
			//v.Generator.AddExpression(newTemp, "0", "", "")
			v.Generator.AddAssign(newTemp, "0")
			v.Generator.AddGoto(lvl2)
			v.Generator.AddLabel(lvl1)
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "%")
			v.Generator.AddLabel(lvl2)

			// Verificar si el divisor es cero
			if rightValue.GetValue() == "0" {
				fmt.Print("ERROR: No se puede modular entre cero")
				/*
					v.SemanticErrors = append(v.SemanticErrors, SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "No se puede multiplicar",
					})
				*/
				return structures.Primitive{
					Value:    "Nil",
					DataType: NilType,
				}
			}

			// Realizar la operación de modulo
			var resultValue string
			dataType := leftValue.GetDataType()

			leftInt, _ := strconv.Atoi(leftValue.GetValue())
			rightInt, _ := strconv.Atoi(rightValue.GetValue())
			resultValue = strconv.Itoa(leftInt % rightInt)

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			fmt.Print("ERROR: No se realizar el módulo")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	case "*":
		// Agregar comentario
		v.Generator.AddComment("-----Multiplicación-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {

			// Agregar la expresión a la lista de expresiones
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// Realizar la operación de multiplicación
			var resultValue string
			dataType := leftValue.GetDataType()

			if dataType == IntType {
				leftInt, _ := strconv.Atoi(leftValue.GetValue())
				rightInt, _ := strconv.Atoi(rightValue.GetValue())
				resultValue = strconv.Itoa(leftInt * rightInt)
			} else {
				leftFloat, _ := strconv.ParseFloat(leftValue.GetValue(), 64)
				rightFloat, _ := strconv.ParseFloat(rightValue.GetValue(), 64)
				resultValue = strconv.FormatFloat(leftFloat*rightFloat, 'f', 4, 64)
			}

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			fmt.Print("ERROR: No se puede multiplicar")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	case "/":
		// Agregar comentario
		v.Generator.AddComment("-----División-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {

			// Validación para división entre cero en C3D
			lvl1 := v.Generator.NewLabel()
			lvl2 := v.Generator.NewLabel()

			v.Generator.AddIf(rightValue.GetValue(), "0", "!=", lvl1)
			v.Generator.AddPrintf("c", "77")
			v.Generator.AddPrintf("c", "97")
			v.Generator.AddPrintf("c", "116")
			v.Generator.AddPrintf("c", "104")
			v.Generator.AddPrintf("c", "69")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "111")
			v.Generator.AddPrintf("c", "114")
			v.Generator.AddPrintf("c", "10")
			//v.Generator.AddExpression(newTemp, "0", "", "")
			v.Generator.AddAssign(newTemp, "0")
			v.Generator.AddGoto(lvl2)
			v.Generator.AddLabel(lvl1)
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")
			v.Generator.AddLabel(lvl2)

			// Verificar si el divisor es cero
			if rightValue.GetValue() == "0" {
				fmt.Print("ERROR: No se puede dividir entre cero")
				/*
					v.SemanticErrors = append(v.SemanticErrors, SemanticError{
						Line:    ctx.GetStart().GetLine(),
						Column:  ctx.GetStart().GetColumn(),
						Message: "No se puede multiplicar",
					})
				*/
				return structures.Primitive{
					Value:    "Nil",
					DataType: NilType,
				}
			}

			// Realizar la operación de división
			var resultValue string
			dataType := leftValue.GetDataType()

			if dataType == IntType {
				leftInt, _ := strconv.Atoi(leftValue.GetValue())
				rightInt, _ := strconv.Atoi(rightValue.GetValue())
				resultValue = strconv.Itoa(leftInt / rightInt)
			} else {
				leftFloat, _ := strconv.ParseFloat(leftValue.GetValue(), 64)
				rightFloat, _ := strconv.ParseFloat(rightValue.GetValue(), 64)
				resultValue = strconv.FormatFloat(leftFloat/rightFloat, 'f', 4, 64)
			}

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			fmt.Print("ERROR: No se realizar la división")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	case "+":
		// Agregar comentario
		v.Generator.AddComment("-----Suma-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if ((leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType)) ||
			(leftValue.GetDataType() == StringType && rightValue.GetDataType() == StringType) {

			// GEneracion C3D
			if leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType {
				// Agregar la expresión a la lista de expresiones
				v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")
			} else {
				fmt.Println("Suma de strings")
			}

			// Realizar la operación de multiplicación
			var resultValue string
			dataType := leftValue.GetDataType()

			if dataType == IntType {
				fmt.Print("Suma de enteros")
				leftInt, _ := strconv.Atoi(leftValue.GetValue())
				rightInt, _ := strconv.Atoi(rightValue.GetValue())
				resultValue = strconv.Itoa(leftInt + rightInt)
			} else if dataType == FloatType {
				fmt.Print("Suma de floats")
				leftFloat, _ := strconv.ParseFloat(leftValue.GetValue(), 64)
				rightFloat, _ := strconv.ParseFloat(rightValue.GetValue(), 64)
				resultValue = strconv.FormatFloat(leftFloat+rightFloat, 'f', 4, 64)
			} else {
				fmt.Print("Suma de Strings")
				resultValue = leftValue.GetValue() + rightValue.GetValue()
			}

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			fmt.Print("ERROR: No se puede sumar")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	case "-":
		// Agregar comentario
		v.Generator.AddComment("-----Resta-----")

		// Obtener los valores izquierdo y derecho
		leftValue := v.Visit(ctx.GetLeft()).(structures.Primitive)
		rightValue := v.Visit(ctx.GetRight()).(structures.Primitive)

		// Verificar si ambos operandos son de tipo IntType o FloatType
		if (leftValue.GetDataType() == IntType || leftValue.GetDataType() == FloatType) &&
			(rightValue.GetDataType() == IntType || rightValue.GetDataType() == FloatType) {

			// Agregar la expresión a la lista de expresiones
			v.Generator.AddExpression(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// Realizar la operación de multiplicación
			var resultValue string
			dataType := leftValue.GetDataType()

			if dataType == IntType {
				leftInt, _ := strconv.Atoi(leftValue.GetValue())
				rightInt, _ := strconv.Atoi(rightValue.GetValue())
				resultValue = strconv.Itoa(leftInt - rightInt)
			} else {
				leftFloat, _ := strconv.ParseFloat(leftValue.GetValue(), 64)
				rightFloat, _ := strconv.ParseFloat(rightValue.GetValue(), 64)
				resultValue = strconv.FormatFloat(leftFloat-rightFloat, 'f', 4, 64)
			}

			return structures.Primitive{
				Value:    resultValue,
				DataType: dataType,
			}
		} else {
			fmt.Print("ERROR: No se puede restar")
			/*
				v.SemanticErrors = append(v.SemanticErrors, SemanticError{
					Line:    ctx.GetStart().GetLine(),
					Column:  ctx.GetStart().GetColumn(),
					Message: "No se puede multiplicar",
				})
			*/
		}
	}
	return nil
}
