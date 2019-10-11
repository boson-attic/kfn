grammar Kfn;

kfn
    : line*
    ;

line
    : stmt (';' | '\n')+
    ;

stmt
    : decl
    | wire
    ;

decl
    : ident ':' component
    ;

ident
    : UPPER_IDENT
    ;

wire
    : wireStmt
    ;

wireStmt
    : wireElement ARROW (wireElement | wireStmt)
    ;

wireElement
    : (component | ident)
    ;

component
    : componentIdent (componentValue)? (COMPONENT_OPTION_SEP componentOptionList)?
    | componentValue
    ;

componentIdent
    : LOWER_IDENT
    ;

componentValue
    : STRING_LITERAL
    ;

componentOptionList
    : componentOption ((',' | 'and') componentOption)*
    ;

componentOption
    : componentOptionIdent (':' | '=')? componentOptionValue
    ;

componentOptionIdent
    : UPPER_IDENT
    ;

componentOptionValue
    : STRING_LITERAL
    ;

COMPONENT_OPTION_SEP
    : 'with'
    ;

UPPER_IDENT
    : 'a'..'z' IDENT_TAIL
    ;

LOWER_IDENT
    : 'A'..'Z' IDENT_TAIL
    ;

STRING_LITERAL
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

fragment IDENT_TAIL
    : ('a'..'z' | 'A'..'Z' | '0'..'9' | '_' | '-')*
    ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;
fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;
fragment HEX
   : [0-9a-fA-F]
   ;
fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

ARROW
    : '=>'
    | '->'
    ;

WS
    : [ \t]+ -> skip
    ;