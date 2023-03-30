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
    | function_def          {% id %}
    | conditional           {% id %}



#___________CONDITIONAL_CALL___________#

conditional 
    -> "if" _ "(" _ comparison _ ")" _ "[" _  %newline:+ statements  %newline:+ _ "]" _ "else" _ "[" _  %newline:+ statements  %newline:+ _ "]"
        {%
            (data) => {
                return {
                    type: "conditional_ifelse",
                    comparison: data[4],
                    if_body: data [11],
                    else_body: data[21]
                }
            }
        %}
    | "if" _ "(" _ comparison _ ")" _ "[" _  %newline:+ statements  %newline:+ _ "]"
        {%
            (data) => {
                return {
                    type: "conditional_if",
                    comparison: data[4],
                    if_body: data [11],
                }
            }
        %}



#___________COMPARISON___________#

comparison 
    -> expression _ %equal_to _ expression
        {%
            (data) => {
                return {
                    type: "comparison",
                    comparison: data[2],
                    expr_one: data[0],
                    expr_two: data[4]
                }
            }
        %}

    |  expression _ %not_equal_to _ expression
        {%
            (data) => {
                return {
                    type: "comparison",
                    comparison: data[2],
                    expr_one: data[0],
                    expr_two: data[4]
                }
            }
        %}



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



#___________FUNCTION_CALL___________#

function_def -> "func" _ %identifier _ "(" _ expression_list _ ")" _ "[" _  %newline:+ statements  %newline:+ _ "]"
    {%
        (data) => {
            return {
                type: "function_def",
                fun_name: data[2],
                parameters: data[6],
                body: data [13],

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