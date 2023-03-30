const moo = require("moo")

module.exports = moo.compile({
    whitespace:      /[ \t]+/,
    number:          { match: /0|[1-9][0-9]*/, value: Number },
    string:          { match: /'(?:\\["\\]|[^\n"\\])*'/, value: String },
    boolean:         { match: /^(?:true|false)$/, value: Boolean},
    left_bracket:    '[',
    right_bracket:   ']',
    left_paren:      '(',
    right_paren:     ')',
    param_separator: ',',
    assignment_op:   { match: /(?<![!=])\=(?!=)/},
    equal_to:        { match: /(?<!\=)\=\=(?!=)/},
    not_equal_to:    { match: /\!=/},
    identifier:      /[a-zA-Z_][a-zA-Z0-9_]*/,
    newline:         { match: /(?:\n|\r\n)/, lineBreaks: true },
});