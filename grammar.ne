@{%
  const l = require("./lexer");
%}

@lexer l


#___________program___________#

program
    -> statements
        {%
            (data) => {
                return {
                    type: "program",
                    body: data[0]
                };
            }
        %}



#___________STATEMENTS___________#

statements
    -> null
        {%
            () =>[]
        %}
    |  _ statement _
        {%
            (data) => [data[1]]
        %}
    |  _ statement _ %newline:+ statements
        {%
            (data) => [data[1], ...data[4]]
        %}



#___________STATEMENT___________#

statement
    -> assignment           {% id %}
    | reassignment          {% id %}
    | function_call         {% id %}



#___________FUNCTION_CALL___________#

function_call -> %identifier _ "(" _ expression_list _ ")"
    {%
        (data) => {
            return {
                type: "function_call",
                fun_name: data[0],
                parameters: data[4]
            }
        }
    %}


#___________ASSIGNMENT___________#

assignment -> "var" _ %identifier _ "=" _ expression
    {%
        (data) => {
            return {
                type: "assignment",
                var_name: data[2],
                value: data[6]
            }
        }
    %}



#___________REASSIGNMENT___________#

reassignment -> %identifier _ "=" _ expression
    {%
        (data) => {
            return {
                type: "reassignment",
                var_name: data[0],
                value: data[4]
            }
        }
    %}


#___________EXPRESSION_LIST___________#

expression_list
    -> expression
        {%
            (data) => {
                return [data[0]]
            }
        %}
    |  expression __ expression_list
        {%
            (data) => {
                return [data[0], ...data[2]]
            }
        %}

#___________EXPRESSIONS___________#

expression
    -> %identifier    {% id %}
    |  literal        {% id %}



#___________LITERALS___________#

literal
    -> %number                   {% id %}
    |  %string                   {% id %}
    |  %boolean                  {% id %}



#___________WHITESPACE___________#

# optional whitespace
_ 
    -> null
    |  __

# mandatory whitespace
__ -> %whitespace