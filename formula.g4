grammar formula;

// starting point for parsing a apexcode file
compilationUnit
    :   expression
    ;

expression
    :   primary                            # primaryExpression
    |   Identifier '(' arguments? ')'      # functionCall
    |   Identifier ('.' Identifier)*       # fieldReference
    |   expression op=('&'|'^') expression # binaryExpression
    |   expression op=('*'|'/') expression # binaryExpression
    |   expression op=('+'|'-') expression # binaryExpression
    |   expression op=('<'|'>'|'<='|'>='|'='|'=='|'!='|'<>') expression # binaryExpression
    |   expression op=('&&'|'||') expression # binaryExpression
    ;

primary
    :   literal
    |   '(' expression ')'
    ;

arguments
    :   expression (',' expression)*
    ;

literal
    :   StringLiteral
    |   IntegerLiteral
    |   FloatingPointLiteral
    |   BooleanLiteral
    |   NullLiteral
    ;

StringLiteral
    :   QUOTE StringCharacters? QUOTE
    ;

fragment
StringCharacters
    :   StringCharacter+
    ;

fragment
StringCharacter
    :   ~["\\]
    ;

IntegerLiteral
    :   Digits
    ;

fragment
Digits
    :   Digit+
    ;

fragment
Digit
    :   '0'
    |   NonZeroDigit
    ;

fragment
NonZeroDigit
    :   [1-9]
    ;

FloatingPointLiteral
    :   Digits '.' Digits?
    ;

BooleanLiteral
    :   T R U E
    |   F A L S E
    ;

NullLiteral
    :   N U L L
    ;

Identifier
    :   Letter LetterOrDigit*
    ;

fragment
Letter
    :   [a-zA-Z$_] // these are the "java letters" below 0xFF
    ;

fragment
LetterOrDigit
    :   [a-zA-Z0-9$_] // these are the "java letters or digits" below 0xFF
    ;

WS  :  [ \t\r\n\u000C]+ -> skip
    ;

QUOTE : '"' -> skip;

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
fragment SPACE : ' ';
