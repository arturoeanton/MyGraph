// R2Query.g4
grammar R2Query;

CREATE: 'CREATE';
DELETE: 'DELETE';
ADD: 'ADD';
COMMA: ',';
GET: 'GET';
MIDDLE: '-';
DIRECTION_REL_R: '->';
DIRECTION_REL_L: '<-';

GRAPH_NAME:[a-zA-Z][a-zA-Z0-9]*;
TAG_NAME:':'('_'|'-'|'$'|'&'|'@'|[a-zA-Z0-9])('_'|'-'|'$'|'&'|'@'|[a-zA-Z0-9])*;
//NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

STRING
    : '"'      (Escape | '""'   | ~["])* '"'
    | '\''     (Escape | '\'\'' | ~['])* '\''
    | '\u201C' (Escape | .)*? ('\u201D' | '\u2033')   // smart quotes
    ;
fragment Escape
    : '\u0060\''    // backtick single-quote
    | '\u0060"'     // backtick double-quote
    ;


arguments : STRING(COMMA STRING)* ;
rel: (op2=MIDDLE '[' tag=TAG_NAME ']' op1=DIRECTION_REL_R) |(op1=DIRECTION_REL_L '[' tag=TAG_NAME ']'op2=MIDDLE)|(op1=DIRECTION_REL_L '[' tag=TAG_NAME']' op2=DIRECTION_REL_R);
relation: '(' r=rel arg=arguments ')';
relations: relation (','relation)*;

expression
    : CREATE GRAPH_NAME                              # CreateGraph
    | DELETE GRAPH_NAME                              # DeleteGraph
    | GRAPH_NAME GET  arguments                      # GetElements
    | GRAPH_NAME GET  relations                      # GetElementRelation
    | GRAPH_NAME ADD arguments                       # AddElements
    | GRAPH_NAME DELETE arguments                    # DeleteElements
    | GRAPH_NAME ADD arguments rel arguments         # AddRelations
    | GRAPH_NAME DELETE arguments rel arguments      # DeleteRelations
    ;

/*
CREATE G1
G1 ADD "NODE1"
G1 ADD "NODE2"-[:TAG]-> "NODE3"
G1 ADD "NODE1"<-[:TAG]-> "NODE3"
G1 ADD "NODE3"<-[:TAG]- "NODE3"
G1 GET   (-[:TAG]->"NODE1")
G1 GET (-[:TAG]->"NODE3")
G1 GET (<-[:TAG]-"NODE2")
G1 GET (<-[:TAG]-"NODE2"),  (<-[:TAG]-"NODE3"), (<-[:TAG]-"NODE1"),  (-[:TAG]->"NODE3")
*/