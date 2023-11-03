// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by GrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by GrammarParser#TypeValueVarDeclaration.
	VisitTypeValueVarDeclaration(ctx *TypeValueVarDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValueVarDeclaration.
	VisitValueVarDeclaration(ctx *ValueVarDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#TypeVarDeclaration.
	VisitTypeVarDeclaration(ctx *TypeVarDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValueVarAssignment.
	VisitValueVarAssignment(ctx *ValueVarAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#IncrementVarAssignment.
	VisitIncrementVarAssignment(ctx *IncrementVarAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#DecrementVarAssignment.
	VisitDecrementVarAssignment(ctx *DecrementVarAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#TypeValueLetDeclaration.
	VisitTypeValueLetDeclaration(ctx *TypeValueLetDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValueLetDeclaration.
	VisitValueLetDeclaration(ctx *ValueLetDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorValuesDeclaration.
	VisitVectorValuesDeclaration(ctx *VectorValuesDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorIdDeclaration.
	VisitVectorIdDeclaration(ctx *VectorIdDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#valuesVectorDeclaration.
	VisitValuesVectorDeclaration(ctx *ValuesVectorDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#vectorType.
	VisitVectorType(ctx *VectorTypeContext) interface{}

	// Visit a parse tree produced by GrammarParser#vectorAppend.
	VisitVectorAppend(ctx *VectorAppendContext) interface{}

	// Visit a parse tree produced by GrammarParser#vectorRemoveLast.
	VisitVectorRemoveLast(ctx *VectorRemoveLastContext) interface{}

	// Visit a parse tree produced by GrammarParser#vectorRemoveAt.
	VisitVectorRemoveAt(ctx *VectorRemoveAtContext) interface{}

	// Visit a parse tree produced by GrammarParser#matrixDeclaration.
	VisitMatrixDeclaration(ctx *MatrixDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#valuesMatrixDeclaration.
	VisitValuesMatrixDeclaration(ctx *ValuesMatrixDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#structAttributes.
	VisitStructAttributes(ctx *StructAttributesContext) interface{}

	// Visit a parse tree produced by GrammarParser#structInstance.
	VisitStructInstance(ctx *StructInstanceContext) interface{}

	// Visit a parse tree produced by GrammarParser#attributesCall.
	VisitAttributesCall(ctx *AttributesCallContext) interface{}

	// Visit a parse tree produced by GrammarParser#structAccess.
	VisitStructAccess(ctx *StructAccessContext) interface{}

	// Visit a parse tree produced by GrammarParser#structAssignment.
	VisitStructAssignment(ctx *StructAssignmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#funcDeclaration.
	VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#paramsFuncDeclaration.
	VisitParamsFuncDeclaration(ctx *ParamsFuncDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#funcCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by GrammarParser#paramsFuncCall.
	VisitParamsFuncCall(ctx *ParamsFuncCallContext) interface{}

	// Visit a parse tree produced by GrammarParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#switchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#cases.
	VisitCases(ctx *CasesContext) interface{}

	// Visit a parse tree produced by GrammarParser#caseStmt.
	VisitCaseStmt(ctx *CaseStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#defaultStmt.
	VisitDefaultStmt(ctx *DefaultStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#range.
	VisitRange(ctx *RangeContext) interface{}

	// Visit a parse tree produced by GrammarParser#guardStmt.
	VisitGuardStmt(ctx *GuardStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#StringEmbededExpr.
	VisitStringEmbededExpr(ctx *StringEmbededExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#NilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructAccessExpr.
	VisitStructAccessExpr(ctx *StructAccessExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IsEmptyExpr.
	VisitIsEmptyExpr(ctx *IsEmptyExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#LogicalOperationExpr.
	VisitLogicalOperationExpr(ctx *LogicalOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#StructInstanceExpr.
	VisitStructInstanceExpr(ctx *StructInstanceExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#FloatEmbededExpr.
	VisitFloatEmbededExpr(ctx *FloatEmbededExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#VectorAccessExpr.
	VisitVectorAccessExpr(ctx *VectorAccessExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#GroupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ArithmeticOperationExpr.
	VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ComparisonOperationExpr.
	VisitComparisonOperationExpr(ctx *ComparisonOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#RelationalOperationExpr.
	VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#CharacterExpr.
	VisitCharacterExpr(ctx *CharacterExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#DigitExpr.
	VisitDigitExpr(ctx *DigitExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#CountExpr.
	VisitCountExpr(ctx *CountExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IntEmbededExpr.
	VisitIntEmbededExpr(ctx *IntEmbededExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#UnaryOperationExpr.
	VisitUnaryOperationExpr(ctx *UnaryOperationExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#BooleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#type.
	VisitType(ctx *TypeContext) interface{}
}
