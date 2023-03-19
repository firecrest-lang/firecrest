package main

import (
	"fmt"
	"os"
	"strconv"
)

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	// End of file.
	EOF
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

type Lexer struct {
	Source   string
	Tokens   []*Token
	Start    int
	Current  int
	Line     int
	HadError bool
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		Source: source,
		Tokens: make([]*Token, 0),
		Line:   1,
	}
}

//This function simply checks whether the current field of the lexer is past the end of the source code. If it is, then the lexer has finished processing all of the input and is at the end of the file.

func (l *Lexer) isAtEnd() bool {
	return l.Current >= len(l.Source)
}

func (l *Lexer) scanToken() []*Token {
	for !l.isAtEnd() {
		l.Start = l.Current
		l.scanTokens()
	}

	l.Tokens = append(l.Tokens, &Token{
		Type:   EOF,
		Lexeme: "",
		Line:   l.Line,
	})

	return l.Tokens
}

// This function reads the next character from the source code and advances the current field of the lexer to point to the next character. It returns the character that it read.
func (l *Lexer) advance() byte {
	c := l.Source[l.Current]
	l.Current++
	return c
}

// This function creates a new Token object from the current lexeme, token type, and line number, and appends it to the list of tokens that the lexer has produced so far. The literal argument is an optional value that can be attached to the token, such as the numeric value for a NUMBER token or the string value for a STRING token. In the case of single-character tokens, the literal argument is nil.
func (l *Lexer) addToken(tokenType TokenType, literal interface{}) {
	lexeme := string(l.Source[l.Start:l.Current])
	token := Token{tokenType, lexeme, literal, l.Line}
	l.Tokens = append(l.Tokens, &token)
}

// This function checks whether the next character in the source code matches the expected character. If it does, it advances the lexer's current field to the next character and returns true. If it doesn't match, it returns false without advancing the lexer. If the lexer has already reached the end of the source code, it also returns false.
func (l *Lexer) match(expected byte) bool {
	if l.isAtEnd() {
		return false
	}
	if l.Source[l.Current] != expected {
		return false
	}
	l.Current++
	return true
}

// This function returns the next character in the source code without advancing the lexer's current field. If the lexer has already reached the end of the source code, it returns the null character '\x00'.
func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return '\x00'
	}
	return l.Source[l.Current]
}

// This function checks whether the given byte c represents a digit character between 0 and 9. It returns true if it is a digit, and false otherwise.
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// This function checks whether the given byte c represents an alphabetic character, including underscores. It returns true if it is an alphabetic character, and false otherwise.
func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// The peekNext() function returns the character immediately following the current character in the source code, without advancing the lexer's current field. If the lexer has already reached the end of the source code, it returns the null character '\x00'.
func (l *Lexer) peekNext() byte {
	if l.Current+1 >= len(l.Source) {
		return '\x00'
	}
	return l.Source[l.Current+1]
}

// The reportError() function prints an error message to standard error output along with the line number where the error occurred. It also sets the lexer's hadError field to true to indicate that an error has occurred.
func (l *Lexer) reportError(line int, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", line, message)
	l.HadError = true
}

/*
The string() function scans a string literal by iterating over the source code until it finds a closing double-quote or reaches the end of the input. It reports an error if it reaches the end of the input before finding a closing double-quote. Once it finds the closing double-quote, it trims the surrounding quotes and adds the resulting string value as a token to the lexer's token list using the addToken() function.
*/
func (l *Lexer) string() {
	for l.peek() != '"' && !l.isAtEnd() {
		if l.peek() == '\n' {
			l.Line++
		}
		l.advance()
	}

	if l.isAtEnd() {
		l.reportError(l.Line, "Unterminated string.")
		return
	}

	// Consume the closing double-quote.
	l.advance()

	// Trim the surrounding quotes and add the token.
	value := l.Source[l.Start+1 : l.Current-1]
	l.addToken(STRING, value)
}

/*
The number() function scans a numeric literal by iterating over the source code until it finds the end of the number. It first scans the integer part of the number, and then looks for a fractional part if there is a dot (.) followed by a digit. Once it has scanned the complete numeric literal, it converts the resulting string to a float64 value using Go's built-in strconv.ParseFloat() function, and adds the resulting value as a token to the lexer's token list using the addToken() function.
*/
func (l *Lexer) number() {
	for isDigit(l.peek()) {
		l.advance()
	}

	// Look for a fractional part.
	if l.peek() == '.' && isDigit(l.peekNext()) {
		// Consume the dot.
		l.advance()

		for isDigit(l.peek()) {
			l.advance()
		}
	}

	value, _ := strconv.ParseFloat(l.Source[l.Start:l.Current], 64)
	l.addToken(NUMBER, value)
}

// The identifier() function scans an identifier by iterating over the source code until it finds a non-alphanumeric character or reaches the end of the input. Once it has scanned the complete identifier, it checks whether the resulting string is a reserved keyword using a lookup table (keywords). If the string is a reserved keyword, it adds the corresponding token type to the lexer's token list using the addToken() function. Otherwise, it adds an IDENTIFIER token to the token list.
func (l *Lexer) identifier() {
	for isAlphaNumeric(l.peek()) {
		l.advance()
	}

	// Check if the identifier is a reserved word.
	text := l.Source[l.Start:l.Current]
	tokenType, ok := keywords[text]
	if ok {
		l.addToken(tokenType, nil)
	} else {
		l.addToken(IDENTIFIER, nil)
	}
}

// The isAlphaNumeric() function is a helper function that returns true if the given byte c is an alphabetic or numeric character. It is used to determine whether to continue scanning an identifier.
func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

/*
The scanToken() function is the core of the Wren lexer in Go. It reads the source code character by character, and identifies and categorizes each token in the input. Here's a brief overview of what it does:

It reads the current character from the source code using the advance() function.
It checks whether the current character is a single-character token (such as ( or +) by comparing it to a set of predefined characters. If it is, it adds the token to the list of tokens using the addToken() function and advances to the next character.
If the current character is not a single-character token, it checks whether it is the start of a multi-character token (such as == or <=) by looking at the next character. If it is, it adds the token to the list of tokens and advances to the next character. If it is not, it treats the character as the start of a new token.
If the current character is a digit, it scans a number using the number() function.
If the current character is a letter, it scans an identifier using the identifier() function.
If the current character is a double-quote ", it scans a string literal using the string() function.
If the current character is a whitespace character, it ignores it and advances to the next character.
If the current character is a newline character, it increments the line field of the lexer to keep track of the current line number.
If the current character is not any of the above, it reports an error using the reportError() function.
The scanToken() function repeats this process until it reaches the end of the input. At that point, it returns a list of tokens that the lexer has produced by calling the addToken() function.
*/
func (l *Lexer) scanTokens() {
	c := l.advance()

	switch c {
	case '(':
		l.addToken(LEFT_PAREN, nil)
	case ')':
		l.addToken(RIGHT_PAREN, nil)
	case '{':
		l.addToken(LEFT_BRACE, nil)
	case '}':
		l.addToken(RIGHT_BRACE, nil)
	case ',':
		l.addToken(COMMA, nil)
	case '.':
		l.addToken(DOT, nil)
	case '-':
		l.addToken(MINUS, nil)
	case '+':
		l.addToken(PLUS, nil)
	case ';':
		l.addToken(SEMICOLON, nil)
	case '*':
		l.addToken(STAR, nil)
	case '!':
		if l.match('=') {
			l.addToken(BANG_EQUAL, nil)
		} else {
			l.addToken(BANG, nil)
		}
	case '=':
		if l.match('=') {
			l.addToken(EQUAL_EQUAL, nil)
		} else {
			l.addToken(EQUAL, nil)
		}
	case '<':
		if l.match('=') {
			l.addToken(LESS_EQUAL, nil)
		} else {
			l.addToken(LESS, nil)
		}
	case '>':
		if l.match('=') {
			l.addToken(GREATER_EQUAL, nil)
		} else {
			l.addToken(GREATER, nil)
		}
	case '/':
		if l.match('/') {
			// A comment goes until the end of the line.
			for l.peek() != '\n' && !l.isAtEnd() {
				l.advance()
			}
		} else {
			l.addToken(SLASH, nil)
		}
	case ' ', '\r', '\t': // Skip whitespace.
		break
	case '\n':
		l.Line++
	case '"':
		l.string()
	default:
		if isDigit(c) {
			l.number()
		} else if isAlpha(c) {
			l.identifier()
		} else {
			l.reportError(l.Line, "Unexpected character.")
		}
	}
}
