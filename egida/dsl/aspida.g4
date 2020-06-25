grammar aspida;

/*

Aspida Full grammar

@author Antonio Paya Gonzalez
*/

/* Parser rules*/

program             : main hosts tasks? variables?;

// Blocks
main                : MAIN_KW ':' '{' main_content '}' ;
hosts               : HOSTS_KW ':' hosts_list  ;
tasks               : TASKS_KW  ':' '{' tasks_content '}' ;
variables           : VARS_KW  ':' '{' vars_content '}' ;

// MAIN Block
main_content        : main_prop (main_prop)*;
main_prop           : (name | connection | description) NS;
name                : 'name' ':' STRING
                    | 'NAME' ':' STRING
                    ;
connection          : 'connection' ':' (LOCAL_KW | SSH_KW | SSHPASS_KW)
                    | 'CONNECTION' ':' (LOCAL_KW | SSH_KW | SSHPASS_KW);
description         : 'description' ':' STRING
                    | 'DESCRIPTION' ':' STRING
                    ;

// HOSTS Block
hosts_list       : '[' host (',' host)* ']'
                 | '[' ']'                      /* localhost */
                 ;
host             : ip_v4 | ip_v6 ;


// TASKS Block
tasks_content        : tasks_prop (tasks_prop)*;
tasks_prop           : (sections | points | controls | exclusions) NS;
sections             : 'sections' ':' str_array
                     | 'SECTIONS' ':' str_array
                     ;
points               : 'points' ':' str_array
                     | 'POINTS' ':' str_array
                     ;
controls             : 'controls' ':' str_array
                     | 'CONTROLS' ':' str_array
                     ;
exclusions           : 'exclusions' ':' str_array
                     | 'EXCLUSIONS' ':' str_array
                     ;

// VARS Block
vars_content         : vars_prop (vars_prop)* ;
vars_prop            : STRING ':' value NS
                     | STRING ':' '{' vars_content '}' NS
                     ;

//Skippables
COMMENT             :  '#' ~('\r' | '\n')* -> skip ;
WHITESPACE          : (' ' | '\t') -> skip ;
NEWLINE             : ('\r'? '\n' | '\r')+ -> skip;

/*=====================================  LEXER RULES  =====================================*/

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
fragment INT
    : '0' | [1-9] [0-9]*
    ;
// no leading zeros
fragment EXP
   : [Ee] [+\-]? INT
   ;

value
   : STRING
   | NUMBER
   | 'true'
   | 'false'
   | 'null'
   | array
   ;

str_array
   : '[' cadena (',' cadena)* ']'
   | '[' ']'
   ;

cadena : STRING ;

array
   : '[' value (',' value)* ']'
   | '[' ']'
   ;

// IP V4
ip_v4                : single_ip | ip_range;
single_ip            : SINGLE_IP;
ip_range             : IP_RANGE;
//IP V6
ip_v6
 : h16 ':' h16 ':' h16 ':' h16 ':' h16 ':' h16 ':' ls32
 | '::' h16 ':' h16 ':' h16 ':' h16 ':' h16 ':' ls32
 | h16? '::' h16 ':' h16 ':' h16 ':' h16 ':' ls32
 | ((h16 ':')? h16)? '::' h16 ':' h16 ':' h16 ':' ls32
 | (((h16 ':')? h16 ':')? h16)? '::' h16 ':' h16 ':' ls32
 | ((((h16 ':')? h16 ':')? h16 ':')? h16)? '::' h16 ':' ls32
 | (((((h16 ':')? h16 ':')? h16 ':')? h16 ':')? h16)? '::' ls32
 | ((((((h16 ':')? h16 ':')? h16 ':')? h16 ':')? h16 ':')? h16)? '::' h16
 ;

h16
 : hexdigit hexdigit hexdigit hexdigit
 | hexdigit hexdigit hexdigit
 | hexdigit hexdigit
 | hexdigit
 ;
ls32
 : h16 ':' h16
 | ip_v4
 ;
//Hex numbers
hexdigit
 : DIGIT
 | HEX_CHAR
 ;


STRING
    : '"' (ESC | SAFECODEPOINT)* '"'
    ;
NUMBER
    : '-'? INT ('.' [0-9] +)? EXP?
    ;
DIGIT
    : [0-9] ;
HEX_CHAR
    : [a-fA-F] ;
NUMBER_RANGE
    : NUMBER'-'NUMBER ;
ARROW
    : '->'
    | '=>'
    ;
SINGLE_IP
    :  NUMBER '.' NUMBER  '.' NUMBER  '.' NUMBER
    ;

IP_RANGE
    :  (NUMBER | NUMBER_RANGE) '.' (NUMBER | NUMBER_RANGE)  '.' (NUMBER | NUMBER_RANGE)  '.' (NUMBER | NUMBER_RANGE)
    ;


// Keywords
NS              : ';';      /*New Sentence*/
MAIN_KW         : 'MAIN';
HOSTS_KW        : 'HOSTS';
TASKS_KW        : 'TASKS';
VARS_KW         : 'VARS';
LOCAL_KW        : 'LOCAL';
SSH_KW          : 'SSH';
SSHPASS_KW      : 'SSH-PASSWORD';
