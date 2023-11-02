grammar Grammar;
import Lexer;

// ? -> uno o nada * -> cero o mas + -> uno o mas
start: block EOF
    ;

block:
	(stmts)* // Return a slice 
    ; 

stmts:
	varDeclaration (SEMICOLON)?
	| varAssignment (SEMICOLON)?

	| letDeclaration (SEMICOLON)?

	| vectorDeclaration (SEMICOLON)?
	| vectorAppend (SEMICOLON)?
	| vectorRemoveLast (SEMICOLON)?
	| vectorRemoveAt (SEMICOLON)?

	| funcDeclaration (SEMICOLON)?
	| funcCall (SEMICOLON)?
	| printStmt (SEMICOLON)?
	| ifStmt (SEMICOLON)?
	| switchStmt (SEMICOLON)?
	| whileStmt (SEMICOLON)?
	| forStmt (SEMICOLON)?
	| guardStmt (SEMICOLON)?
	| breakStmt (SEMICOLON)?
	| continueStmt (SEMICOLON)?
	| returnStmt (SEMICOLON)?
    ;


// Variables
varDeclaration:
	VAR ID COLON type EQUAL expr		# TypeValueVarDeclaration // var value: String = "Hola"	
	| VAR ID EQUAL expr					# ValueVarDeclaration	// var value = "Hola"
	| VAR ID COLON type QUESTION_MARK	# TypeVarDeclaration	// var value: String?	
    ;

varAssignment:
	(ID|ID '[' expr ']') EQUAL expr	# ValueVarAssignment
	|(ID|ID '[' expr ']') INCREMENT expr	# IncrementVarAssignment
	|(ID|ID '[' expr ']') DECREMENT expr	# DecrementVarAssignment
	;	


// Constants
letDeclaration:
	LET ID COLON type EQUAL expr	# TypeValueLetDeclaration
	| LET ID EQUAL expr				# ValueLetDeclaration
    ;


// Vector
vectorDeclaration: 
	VAR ID COLON '[' type ']' '=' '[' (valuesVectorDeclaration)? ']' # VectorValuesDeclaration
	| VAR ID COLON '[' type ']' '=' ID	# VectorIdDeclaration
	;	

valuesVectorDeclaration: 
	expr (',' expr)*
	;	

vectorType: 
	'[' type ']'
	;

vectorAppend: 
	ID '.append' '(' expr ')'
	;

vectorRemoveLast: 
	ID '.removeLast' '('')'
	;

vectorRemoveAt: 
	ID '.remove' '(' 'at' ':' expr ')'
	;

// Matrix
matrixDeclaration:
	VAR ID COLON '[' '[' type ']' ']' '=' '[' (valuesMatrixDeclaration)? ']'
	| VAR ID COLON '[' '[' type ']' ']' '=' ID	
	;
	
valuesMatrixDeclaration:
	'[' (valuesVectorDeclaration)? ']' (',' '[' (valuesVectorDeclaration)? ']')*
	;


// Structs
structDeclaration: 'struct' ID '{' (structAttribute | funcDeclaration)* '}'
;

structAttribute
    : (VAR|LET) ID (':' type)? ('=' expr)? (';')?
;

structAccess: ID ('.' ID)+
;

structAssignment
    : ID '.' ID '=' expr
;


// Functions
funcDeclaration: 
	FUNC ID '(' (paramsFuncDeclaration)? ')' ('->' type)? '{' block '}'
	;	

paramsFuncDeclaration: 
    (ID | UNDERSCORE)? ID ':' INOUT? (type|vectorType) (',' paramsFuncDeclaration)?
    ;

funcCall: 
	ID '(' (paramsFuncCall)? ')'
	;	

paramsFuncCall: 
	(ID ':')? REFERENCE? expr (',' paramsFuncCall)?
	;


// Print 
printStmt: PRINT '(' expr ')'
		| PRINT '('expr (',' expr)* ')'
    ;


// If - Else 
ifStmt:
	IF expr '{' block '}'
	| IF expr '{' block '}' ELSE '{' block '}'
	| IF expr '{' block '}' ELSE ifStmt
    ;


// Switch - Case 
switchStmt: 
	SWITCH expr '{' cases '}'
    ;

cases: 
	(caseStmt | defaultStmt)*
    ;

caseStmt: 
	CASE expr ':' block (BREAK)?
    ;

defaultStmt: 
	DEFAULT ':' block (BREAK)?
    ;


// While  
whileStmt: 
	WHILE expr '{' block '}'
    ; 


// For 
forStmt: 
	FOR (ID|UNDERSCORE) IN (expr | range) '{' block '}'
	;	

range: 
	expr '...' expr
	;	


// Guard
guardStmt: 
	GUARD expr ELSE '{' block '}' 
	;	
	

// Transfers	
breakStmt: 
	BREAK
	;


continueStmt:
	CONTINUE
	;
	

returnStmt: 	
	RETURN
	| RETURN expr
	;	


// Expressions
expr:
	// Agroupation
	'(' expr ')' # GroupExpr
	// Unary
	| op = ('!' | '-') right = expr # UnaryOperationExpr
	// Arithmetic
	| left = expr op = ('*' | '/' | '%') right = expr	# ArithmeticOperationExpr
	| left = expr op = ('+' | '-') right = expr			# ArithmeticOperationExpr
	// Relational
	| left = expr op = ('>=' | '>') right = expr	# RelationalOperationExpr
	| left = expr op = ('<=' | '<') right = expr	# RelationalOperationExpr
	// Equality
	| left = expr op = ('==' | '!=') right = expr # ComparisonOperationExpr
	// Logical
	| left = expr op = ('&&' | '||') right = expr # LogicalOperationExpr
	// Primitives
	| ('true' | 'false')	# BooleanExpr
	| NIL					# NilExpr
	| DIGIT					# DigitExpr
	| STR					# StringExpr
	| CHAR 					# CharacterExpr	
	| funcCall              # FuncCallExpr
	// Vector
	| ID '.count'			# CountExpr
	| ID '.isEmpty'			# IsEmptyExpr
	| ID '[' expr ']'		# VectorAccessExpr
	// ID
	| ID					# IdExpr
	// Embeded
	| INT '(' expr ')' # IntEmbededExpr
	| FLOAT '(' expr ')' # FloatEmbededExpr 
	| STRING '(' expr ')' # StringEmbededExpr
    ;


// Types
type: (INT | FLOAT | STRING | BOOL | CHARACTER)
    ;
