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
    {"name": "assignment", "symbols": [{"literal":"var"}, "_", (l.has("identifier") ? {type: "identifier"} : identifier), "_", {"literal":"="}, "_", "expression"], "postprocess": 
        (data) => {
            return {
                type: "assignment",
                var_name: data[2],
                value: data[6]
            }
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
