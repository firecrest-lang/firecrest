package main

import (
	"log"
	"regexp"
	"strings"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	Tokens []Token
}

func (l *Lexer) New(input string) Lexer {
	var lexer Lexer
	return lexer
}

func isLetter(input string) bool {
	matched, err := regexp.MatchString("/[a-zA-Z]/", input)
	if err != nil {
		log.Fatal(err)
	}
	return matched
}

func isNumber(input string) bool {
	matched, err := regexp.MatchString("/^[0-9]+$/", input)
	if err != nil {
		log.Fatal(err)
	}
	return matched
}

func isWhitespace(input string) bool {
	matched, err := regexp.MatchString("/\\s+/", input)
	if err != nil {
		log.Fatal(err)
	}
	return matched
}

func isOperator(input string) bool {
	operators := []string{"+", "-", "*", "/", "%"}
	var contained bool
	for _, substring := range operators {
		contained = strings.Contains(input, substring)
	}
	return contained
}

func isOpeningParenthesis(input string) bool {
	contained := strings.Contains(input, "(")
	return contained
}

func isClosingParenthesis(input string) bool {
	contained := strings.Contains(input, ")")
	return contained
}

func isParenthesis(input string) bool {
	if isOpeningParenthesis(input) == true || isClosingParenthesis(input) == true {
		return true
	} else {
		return false
	}
}

func isDoubleQuote(input string) bool {
	contained := strings.Contains(input, "\"")
	return contained
}

func isSingleQuote(input string) bool {
	contained := strings.Contains(input, "'")
	return contained
}

func isQuote(input string) bool {
	if isDoubleQuote(input) == true || isSingleQuote(input) == true {
		return true
	} else {
		return false
	}
}
