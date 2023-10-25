lexer grammar Lexer;

INT: 'Int';
FLOAT: 'Float';
STRING: 'String';
BOOL: 'Bool';
CHARACTER: 'Character';
NIL : 'nil';

VAR: 'var';
LET: 'let';

FUNC: 'func';
INOUT: 'inout';
REFERENCE: '&';

PRINT: 'print';

IF: 'if';
ELSE: 'else';

SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';

WHILE: 'while';

FOR: 'for';
IN: 'in';

GUARD: 'guard';

CONTINUE: 'continue';
BREAK: 'break';
RETURN: 'return';

SEMICOLON: ';';
COLON: ':';
EQUAL:'=';
QUESTION_MARK: '?';
INCREMENT: '+=';
DECREMENT: '-=';
UNDERSCORE: '_';

// ? -> uno o nada 
// * -> cero o mas
// + -> uno o mas
DIGIT: [0-9]+ ('.'[0-9]+)?;
STR: '"'~["]*'"'; 
CHAR: '\'' ~'\'' '\'';
ID:  ( '_' | [a-zA-Z] ) [a-zA-Z0-9_]*;

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
MULTI_COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment // fragment means that this rule is not a token
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;