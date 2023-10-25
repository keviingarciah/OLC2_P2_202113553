package compiler

import (
	"backend/parser"
	"backend/structures"
	"fmt"
)

// ----------------------------------- VAR -----------------------------------
func (v *Visitor) VisitTypeValueVarDeclaration(ctx *parser.TypeValueVarDeclarationContext) interface{} {
	// get the id
	varId := ctx.ID().GetText()

	// Evaluate the id
	if _, found := v.currentEnv.Symbols[varId]; found {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable ya declarada.",
		})
	}

	// Get the values
	primitive := v.Visit(ctx.Expr()).(structures.Primitive)
	dataType := ctx.Type_().GetText()

	// Evaluate the type
	if primitive.GetDataType() != dataType {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: fmt.Sprintf("Tipo del dato incorrecto, se esperaba %s.", dataType),
		})
	}

	// Agregar comentario
	v.Generator.AddComment("-------Declaracion de Variable-------")

	newTemp := v.Generator.NewTemp()

	v.Generator.AddAssign(newTemp, primitive.GetValue())
	v.Generator.AddSetStack("(int)P", newTemp)
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, primitive.GetDataType(), newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	// Aumentar el stack counter
	v.Generator.StackCounter++

	return structures.Primitive{
		Value:    primitive.GetValue(),
		DataType: primitive.GetDataType(),
	}
}

func (v *Visitor) VisitValueVarDeclaration(ctx *parser.ValueVarDeclarationContext) interface{} {
	// get the id
	varId := ctx.ID().GetText()

	// Evaluate the id
	if _, found := v.currentEnv.Symbols[varId]; found {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable ya declarada.",
		})
	}

	// get the primitive value
	primitive := v.Visit(ctx.Expr()).(structures.Primitive)
	//size := strconv.Itoa(0)

	// Agregar comentario
	v.Generator.AddComment("-------Declaracion de Variable-------")

	newTemp := v.Generator.NewTemp()

	v.Generator.AddAssign(newTemp, primitive.GetValue())
	v.Generator.AddSetStack("(int)P", newTemp)
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, primitive.GetDataType(), newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	// Aumentar el stack counter
	v.Generator.StackCounter++

	return structures.Primitive{
		Value:    primitive.GetValue(),
		DataType: primitive.GetDataType(),
	}
}

func (v *Visitor) VisitTypeVarDeclaration(ctx *parser.TypeVarDeclarationContext) interface{} {
	/*
		// get the id
		varId := ctx.ID().GetText()

		// get Type
		varType := ctx.Type_().GetText()


			if varType == "Int" {
				// save the value in the memory
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, varType, &IntPrimitive{Value: 0}, v.currentEnv.Name)

			} else if varType == "Float" {
				// save the value in the memory
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, varType, &FloatPrimitive{Value: 0.0}, v.currentEnv.Name)

			} else if varType == "String" {
				// save the value in the memory
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, varType, &StringPrimitive{Value: ""}, v.currentEnv.Name)

			} else if varType == "Character" {
				// save the value in the memorya
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, varType, &CharPrimitive{Value: ""}, v.currentEnv.Name)

			} else if varType == "Bool" {
				// save the value in the memory
				v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, varType, &BoolPrimitive{Value: false}, v.currentEnv.Name)

			}
	*/
	return nil
}

// ----------------------------------- LET -----------------------------------
func (v *Visitor) VisitTypeValueLetDeclaration(ctx *parser.TypeValueLetDeclarationContext) interface{} {
	// get the id
	varId := ctx.ID().GetText()

	// Evaluate the id
	if _, found := v.currentEnv.Symbols[varId]; found {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable ya declarada.",
		})
	}

	// Get the values
	primitive := v.Visit(ctx.Expr()).(structures.Primitive)
	dataType := ctx.Type_().GetText()

	// Evaluate the type
	if primitive.GetDataType() != dataType {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: fmt.Sprintf("Tipo del dato incorrecto, se esperaba %s.", dataType),
		})
	}

	// Agregar comentario
	v.Generator.AddComment("-------Declaracion de Constante-------")

	newTemp := v.Generator.NewTemp()

	v.Generator.AddAssign(newTemp, primitive.GetValue())
	v.Generator.AddSetStack("(int)P", newTemp)
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "let", varId, primitive.GetDataType(), newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	// Aumentar el stack counter
	v.Generator.StackCounter++

	return structures.Primitive{
		Value:    primitive.GetValue(),
		DataType: primitive.GetDataType(),
	}
}

func (v *Visitor) VisitValueLetDeclaration(ctx *parser.ValueLetDeclarationContext) interface{} {
	// get the id
	varId := ctx.ID().GetText()

	// Evaluate the id
	if _, found := v.currentEnv.Symbols[varId]; found {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Variable ya declarada.",
		})
	}

	// get the primitive value
	primitive := v.Visit(ctx.Expr()).(structures.Primitive)
	//size := strconv.Itoa(0)

	// Agregar comentario
	v.Generator.AddComment("-------Declaracion de Constante-------")

	newTemp := v.Generator.NewTemp()

	v.Generator.AddAssign(newTemp, primitive.GetValue())
	v.Generator.AddSetStack("(int)P", newTemp)
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "let", varId, primitive.GetDataType(), newTemp, v.currentEnv.Name, v.Generator.StackCounter)

	// Aumentar el stack counter
	v.Generator.StackCounter++

	return structures.Primitive{
		Value:    primitive.GetValue(),
		DataType: primitive.GetDataType(),
	}
}
