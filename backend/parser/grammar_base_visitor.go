// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar
import "github.com/antlr4-go/antlr/v4"

type BaseGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeValueVarDeclaration(ctx *TypeValueVarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValueVarDeclaration(ctx *ValueVarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeVarDeclaration(ctx *TypeVarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValueVarAssignment(ctx *ValueVarAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIncrementVarAssignment(ctx *IncrementVarAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDecrementVarAssignment(ctx *DecrementVarAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitTypeValueLetDeclaration(ctx *TypeValueLetDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValueLetDeclaration(ctx *ValueLetDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorValuesDeclaration(ctx *VectorValuesDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorIdDeclaration(ctx *VectorIdDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValuesVectorDeclaration(ctx *ValuesVectorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorType(ctx *VectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorAppend(ctx *VectorAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorRemoveLast(ctx *VectorRemoveLastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorRemoveAt(ctx *VectorRemoveAtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitMatrixDeclaration(ctx *MatrixDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitValuesMatrixDeclaration(ctx *ValuesMatrixDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructAttributes(ctx *StructAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructAccess(ctx *StructAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStructAssignment(ctx *StructAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFuncDeclaration(ctx *FuncDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitParamsFuncDeclaration(ctx *ParamsFuncDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitParamsFuncCall(ctx *ParamsFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCases(ctx *CasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCaseStmt(ctx *CaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDefaultStmt(ctx *DefaultStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitRange(ctx *RangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitGuardStmt(ctx *GuardStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStringEmbededExpr(ctx *StringEmbededExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIsEmptyExpr(ctx *IsEmptyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitLogicalOperationExpr(ctx *LogicalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFloatEmbededExpr(ctx *FloatEmbededExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitVectorAccessExpr(ctx *VectorAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitGroupExpr(ctx *GroupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitArithmeticOperationExpr(ctx *ArithmeticOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitComparisonOperationExpr(ctx *ComparisonOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitRelationalOperationExpr(ctx *RelationalOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCharacterExpr(ctx *CharacterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitDigitExpr(ctx *DigitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitCountExpr(ctx *CountExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitIntEmbededExpr(ctx *IntEmbededExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitUnaryOperationExpr(ctx *UnaryOperationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitFuncCallExpr(ctx *FuncCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}
