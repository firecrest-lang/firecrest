// Generated automatically by nearley, version 2.19.2
// http://github.com/Hardmath123/nearley
(function () {
function id(x) { return x[0]; }

  const l = require("./lexer");
var grammar = {
    Lexer: l,
    ParserRules: [
    {"name": "program", "symbols": ["statements"], "postprocess": 
        (data) => {
            return {
                type: "program",
                body: data[0]
            };
        }
                },
    {"name": "statements", "symbols": [], "postprocess": 
        () =>[]
                },
    {"name": "statements", "symbols": ["_", "statement", "_"], "postprocess": 
        (data) => [data[1]]
                },
    {"name": "statements$ebnf$1", "symbols": [(l.has("newline") ? {type: "newline"} : newline)]},
    {"name": "statements$ebnf$1", "symbols": ["statements$ebnf$1", (l.has("newline") ? {type: "newline"} : newline)], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "statements", "symbols": ["_", "statement", "_", "statements$ebnf$1", "statements"], "postprocess": 
        (data) => [data[1], ...data[4]]
                },
    {"name": "statement", "symbols": ["assignment"], "postprocess": id},
    {"name": "statement", "symbols": ["reassignment"], "postprocess": id},
    {"name": "statement", "symbols": ["function_call"], "postprocess": id},
    {"name": "statement", "symbols": ["function_def"], "postprocess": id},
    {"name": "function_call", "symbols": [(l.has("identifier") ? {type: "identifier"} : identifier), "_", {"literal":"("}, "_", "expression_list", "_", {"literal":")"}], "postprocess": 
        (data) => {
            return {
                type: "function_call",
                fun_name: data[0],
                parameters: data[4]
            }
        }
            },
    {"name": "function_def$ebnf$1", "symbols": [(l.has("newline") ? {type: "newline"} : newline)]},
    {"name": "function_def$ebnf$1", "symbols": ["function_def$ebnf$1", (l.has("newline") ? {type: "newline"} : newline)], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "function_def$ebnf$2", "symbols": [(l.has("newline") ? {type: "newline"} : newline)]},
    {"name": "function_def$ebnf$2", "symbols": ["function_def$ebnf$2", (l.has("newline") ? {type: "newline"} : newline)], "postprocess": function arrpush(d) {return d[0].concat([d[1]]);}},
    {"name": "function_def", "symbols": [{"literal":"func"}, "_", (l.has("identifier") ? {type: "identifier"} : identifier), "_", {"literal":"("}, "_", "expression_list", "_", {"literal":")"}, "_", {"literal":"["}, "_", "function_def$ebnf$1", "statements", "function_def$ebnf$2", "_", {"literal":"]"}], "postprocess": 
        (data) => {
            return {
                type: "function_def",
                fun_name: data[2],
                parameters: data[6],
                body: data [13],
        
            }
        }
            },
    {"name": "assignment", "symbols": [{"literal":"var"}, "_", (l.has("identifier") ? {type: "identifier"} : identifier), "_", {"literal":"="}, "_", "expression"], "postprocess": 
        (data) => {
            return {
                type: "assignment",
                var_name: data[2],
                value: data[6]
            }
        }
            },
    {"name": "reassignment", "symbols": [(l.has("identifier") ? {type: "identifier"} : identifier), "_", {"literal":"="}, "_", "expression"], "postprocess": 
        (data) => {
            return {
                type: "reassignment",
                var_name: data[0],
                value: data[4]
            }
        }
            },
    {"name": "expression_list", "symbols": ["expression"], "postprocess": 
        (data) => {
            return [data[0]]
        }
                },
    {"name": "expression_list", "symbols": ["expression", "__", "expression_list"], "postprocess": 
        (data) => {
            return [data[0], ...data[2]]
        }
                },
    {"name": "expression", "symbols": [(l.has("identifier") ? {type: "identifier"} : identifier)], "postprocess": id},
    {"name": "expression", "symbols": ["literal"], "postprocess": id},
    {"name": "literal", "symbols": [(l.has("number") ? {type: "number"} : number)], "postprocess": id},
    {"name": "literal", "symbols": [(l.has("string") ? {type: "string"} : string)], "postprocess": id},
    {"name": "literal", "symbols": [(l.has("boolean") ? {type: "boolean"} : boolean)], "postprocess": id},
    {"name": "_", "symbols": []},
    {"name": "_", "symbols": ["__"]},
    {"name": "__", "symbols": [(l.has("whitespace") ? {type: "whitespace"} : whitespace)]}
]
  , ParserStart: "program"
}
if (typeof module !== 'undefined'&& typeof module.exports !== 'undefined') {
   module.exports = grammar;
} else {
   window.grammar = grammar;
}
})();
