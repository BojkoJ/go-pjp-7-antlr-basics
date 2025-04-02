grammar PLC_Lab7_expr_2;

/** The start rule; begin parsing here. */
prog: expr_list EOF;

expr_list: (expr ';')+;

expr: expr ('*'|'/') expr    # MulDiv
    | expr ('+'|'-') expr    # AddSub
    | INT                    # Int
    | HEX                    # Hex
    | OCT                    # Oct
    | '(' expr ')'           # Parens
    ;

// Integer literals in different formats
INT : [1-9][0-9]* ;             // Decimal: starts with 1-9
OCT : '0'[0-7]* ;               // Octal: starts with 0
HEX : '0x'[0-9a-fA-F]+ ;        // Hexadecimal: starts with 0x

WS : [ \t\r\n]+ -> skip ;       // Skip whitespace
