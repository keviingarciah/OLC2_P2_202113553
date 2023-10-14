package compiler

import (
	"backend/parser"
	"backend/structures"
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

	// Get the primitive value
	primitive := v.Visit(ctx.Expr()).(structures.Primitive)

	// Evaluate the type
	if primitive.GetDataType() != ctx.Type_().GetText() {
		v.SemanticErrors = append(v.SemanticErrors, structures.SemanticError{
			Line:    ctx.GetStart().GetLine(),
			Column:  ctx.GetStart().GetColumn(),
			Message: "Tipo del dato incorrecto.",
		})
	}

	// Agregar comentario
	v.Generator.AddComment("-------Declaracion de variable-------")

	v.Generator.AddSetStack("(int)P", primitive.GetValue())
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, primitive.GetDataType(), primitive.GetValue(), v.currentEnv.Name, v.Generator.StackCounter)

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
	v.Generator.AddComment("-------Declaracion de variable-------")

	v.Generator.AddSetStack("(int)P", primitive.GetValue())
	v.Generator.AddExpression("P", "P", "1", "+")
	v.Generator.AddBr()

	// Save the value in the symbol table (env and global)
	v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, primitive.GetDataType(), primitive.GetValue(), v.currentEnv.Name, v.Generator.StackCounter)

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
	/*
		// get the id
		varId := ctx.ID().GetText()

		// get the primitive value
		primitive := v.Visit(ctx.Expr()).(PRIMITIVE)

		// save the value in the memory
		v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "let", varId, primitive.GetType(), primitive, v.currentEnv.Name)
	*/
	return nil
}

func (v *Visitor) VisitValueLetDeclaration(ctx *parser.ValueLetDeclarationContext) interface{} {
	/*
		// get the id
		varId := ctx.ID().GetText()

		// get the primitive value
		primitive := v.Visit(ctx.Expr()).(PRIMITIVE)

		// save the value in the memory
		v.SaveSymbol(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "var", varId, primitive.GetType(), primitive, v.currentEnv.Name)
	*/
	return nil
}
