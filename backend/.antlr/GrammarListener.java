// Generated from /media/keviin/Almacenamiento/SEGUNDO SEMESTRE/COMPI 2/LAB/PROYECTOS/PROYECTO 2/OLC2_P2_202113553/backend/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GrammarParser}.
 */
public interface GrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GrammarParser#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(GrammarParser.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(GrammarParser.StartContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(GrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(GrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#stmts}.
	 * @param ctx the parse tree
	 */
	void enterStmts(GrammarParser.StmtsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#stmts}.
	 * @param ctx the parse tree
	 */
	void exitStmts(GrammarParser.StmtsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TypeValueVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterTypeValueVarDeclaration(GrammarParser.TypeValueVarDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeValueVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitTypeValueVarDeclaration(GrammarParser.TypeValueVarDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TypeVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterTypeVarDeclaration(GrammarParser.TypeVarDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitTypeVarDeclaration(GrammarParser.TypeVarDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ValueVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterValueVarDeclaration(GrammarParser.ValueVarDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ValueVarDeclaration}
	 * labeled alternative in {@link GrammarParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitValueVarDeclaration(GrammarParser.ValueVarDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ValueVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void enterValueVarAssignment(GrammarParser.ValueVarAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ValueVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void exitValueVarAssignment(GrammarParser.ValueVarAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IncrementVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void enterIncrementVarAssignment(GrammarParser.IncrementVarAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IncrementVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void exitIncrementVarAssignment(GrammarParser.IncrementVarAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DecrementVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void enterDecrementVarAssignment(GrammarParser.DecrementVarAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DecrementVarAssignment}
	 * labeled alternative in {@link GrammarParser#varAssignment}.
	 * @param ctx the parse tree
	 */
	void exitDecrementVarAssignment(GrammarParser.DecrementVarAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TypeValueLetDeclaration}
	 * labeled alternative in {@link GrammarParser#letDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterTypeValueLetDeclaration(GrammarParser.TypeValueLetDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeValueLetDeclaration}
	 * labeled alternative in {@link GrammarParser#letDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitTypeValueLetDeclaration(GrammarParser.TypeValueLetDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ValueLetDeclaration}
	 * labeled alternative in {@link GrammarParser#letDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterValueLetDeclaration(GrammarParser.ValueLetDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ValueLetDeclaration}
	 * labeled alternative in {@link GrammarParser#letDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitValueLetDeclaration(GrammarParser.ValueLetDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#vectorDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVectorDeclaration(GrammarParser.VectorDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#vectorDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVectorDeclaration(GrammarParser.VectorDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#valuesVectorDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterValuesVectorDeclaration(GrammarParser.ValuesVectorDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#valuesVectorDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitValuesVectorDeclaration(GrammarParser.ValuesVectorDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#vectorType}.
	 * @param ctx the parse tree
	 */
	void enterVectorType(GrammarParser.VectorTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#vectorType}.
	 * @param ctx the parse tree
	 */
	void exitVectorType(GrammarParser.VectorTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#vectorAppend}.
	 * @param ctx the parse tree
	 */
	void enterVectorAppend(GrammarParser.VectorAppendContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#vectorAppend}.
	 * @param ctx the parse tree
	 */
	void exitVectorAppend(GrammarParser.VectorAppendContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#vectorRemoveLast}.
	 * @param ctx the parse tree
	 */
	void enterVectorRemoveLast(GrammarParser.VectorRemoveLastContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#vectorRemoveLast}.
	 * @param ctx the parse tree
	 */
	void exitVectorRemoveLast(GrammarParser.VectorRemoveLastContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#vectorRemoveAt}.
	 * @param ctx the parse tree
	 */
	void enterVectorRemoveAt(GrammarParser.VectorRemoveAtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#vectorRemoveAt}.
	 * @param ctx the parse tree
	 */
	void exitVectorRemoveAt(GrammarParser.VectorRemoveAtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#matrixDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterMatrixDeclaration(GrammarParser.MatrixDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#matrixDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitMatrixDeclaration(GrammarParser.MatrixDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#valuesMatrixDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterValuesMatrixDeclaration(GrammarParser.ValuesMatrixDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#valuesMatrixDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitValuesMatrixDeclaration(GrammarParser.ValuesMatrixDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#structDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterStructDeclaration(GrammarParser.StructDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#structDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitStructDeclaration(GrammarParser.StructDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#structAttribute}.
	 * @param ctx the parse tree
	 */
	void enterStructAttribute(GrammarParser.StructAttributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#structAttribute}.
	 * @param ctx the parse tree
	 */
	void exitStructAttribute(GrammarParser.StructAttributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#structFunction}.
	 * @param ctx the parse tree
	 */
	void enterStructFunction(GrammarParser.StructFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#structFunction}.
	 * @param ctx the parse tree
	 */
	void exitStructFunction(GrammarParser.StructFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#structAccess}.
	 * @param ctx the parse tree
	 */
	void enterStructAccess(GrammarParser.StructAccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#structAccess}.
	 * @param ctx the parse tree
	 */
	void exitStructAccess(GrammarParser.StructAccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#structAssignment}.
	 * @param ctx the parse tree
	 */
	void enterStructAssignment(GrammarParser.StructAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#structAssignment}.
	 * @param ctx the parse tree
	 */
	void exitStructAssignment(GrammarParser.StructAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#funcDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFuncDeclaration(GrammarParser.FuncDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#funcDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFuncDeclaration(GrammarParser.FuncDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#paramsFuncDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterParamsFuncDeclaration(GrammarParser.ParamsFuncDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#paramsFuncDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitParamsFuncDeclaration(GrammarParser.ParamsFuncDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#funcCall}.
	 * @param ctx the parse tree
	 */
	void enterFuncCall(GrammarParser.FuncCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#funcCall}.
	 * @param ctx the parse tree
	 */
	void exitFuncCall(GrammarParser.FuncCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#paramsFuncCall}.
	 * @param ctx the parse tree
	 */
	void enterParamsFuncCall(GrammarParser.ParamsFuncCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#paramsFuncCall}.
	 * @param ctx the parse tree
	 */
	void exitParamsFuncCall(GrammarParser.ParamsFuncCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintStmt(GrammarParser.PrintStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintStmt(GrammarParser.PrintStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(GrammarParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(GrammarParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmt(GrammarParser.SwitchStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#switchStmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmt(GrammarParser.SwitchStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#cases}.
	 * @param ctx the parse tree
	 */
	void enterCases(GrammarParser.CasesContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#cases}.
	 * @param ctx the parse tree
	 */
	void exitCases(GrammarParser.CasesContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#caseStmt}.
	 * @param ctx the parse tree
	 */
	void enterCaseStmt(GrammarParser.CaseStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#caseStmt}.
	 * @param ctx the parse tree
	 */
	void exitCaseStmt(GrammarParser.CaseStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#defaultStmt}.
	 * @param ctx the parse tree
	 */
	void enterDefaultStmt(GrammarParser.DefaultStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#defaultStmt}.
	 * @param ctx the parse tree
	 */
	void exitDefaultStmt(GrammarParser.DefaultStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#whileStmt}.
	 * @param ctx the parse tree
	 */
	void enterWhileStmt(GrammarParser.WhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#whileStmt}.
	 * @param ctx the parse tree
	 */
	void exitWhileStmt(GrammarParser.WhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(GrammarParser.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(GrammarParser.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#range}.
	 * @param ctx the parse tree
	 */
	void enterRange(GrammarParser.RangeContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#range}.
	 * @param ctx the parse tree
	 */
	void exitRange(GrammarParser.RangeContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#guardStmt}.
	 * @param ctx the parse tree
	 */
	void enterGuardStmt(GrammarParser.GuardStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#guardStmt}.
	 * @param ctx the parse tree
	 */
	void exitGuardStmt(GrammarParser.GuardStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(GrammarParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(GrammarParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(GrammarParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(GrammarParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(GrammarParser.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(GrammarParser.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStringEmbededExpr(GrammarParser.StringEmbededExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStringEmbededExpr(GrammarParser.StringEmbededExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterStringExpr(GrammarParser.StringExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitStringExpr(GrammarParser.StringExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NilExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterNilExpr(GrammarParser.NilExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NilExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitNilExpr(GrammarParser.NilExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdExpr(GrammarParser.IdExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdExpr(GrammarParser.IdExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IsEmptyExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIsEmptyExpr(GrammarParser.IsEmptyExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IsEmptyExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIsEmptyExpr(GrammarParser.IsEmptyExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LogicalOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOperationExpr(GrammarParser.LogicalOperationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LogicalOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOperationExpr(GrammarParser.LogicalOperationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFloatEmbededExpr(GrammarParser.FloatEmbededExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFloatEmbededExpr(GrammarParser.FloatEmbededExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorAccessExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorAccessExpr(GrammarParser.VectorAccessExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorAccessExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorAccessExpr(GrammarParser.VectorAccessExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GroupExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterGroupExpr(GrammarParser.GroupExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GroupExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitGroupExpr(GrammarParser.GroupExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArithmeticOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterArithmeticOperationExpr(GrammarParser.ArithmeticOperationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArithmeticOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitArithmeticOperationExpr(GrammarParser.ArithmeticOperationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ComparisonOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterComparisonOperationExpr(GrammarParser.ComparisonOperationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ComparisonOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitComparisonOperationExpr(GrammarParser.ComparisonOperationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RelationalOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterRelationalOperationExpr(GrammarParser.RelationalOperationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RelationalOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitRelationalOperationExpr(GrammarParser.RelationalOperationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DigitExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterDigitExpr(GrammarParser.DigitExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DigitExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitDigitExpr(GrammarParser.DigitExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CountExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterCountExpr(GrammarParser.CountExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CountExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitCountExpr(GrammarParser.CountExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterIntEmbededExpr(GrammarParser.IntEmbededExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntEmbededExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitIntEmbededExpr(GrammarParser.IntEmbededExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnaryOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOperationExpr(GrammarParser.UnaryOperationExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnaryOperationExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOperationExpr(GrammarParser.UnaryOperationExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BooleanExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterBooleanExpr(GrammarParser.BooleanExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BooleanExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitBooleanExpr(GrammarParser.BooleanExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncCallExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterFuncCallExpr(GrammarParser.FuncCallExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncCallExpr}
	 * labeled alternative in {@link GrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitFuncCallExpr(GrammarParser.FuncCallExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link GrammarParser#type}.
	 * @param ctx the parse tree
	 */
	void enterType(GrammarParser.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#type}.
	 * @param ctx the parse tree
	 */
	void exitType(GrammarParser.TypeContext ctx);
}