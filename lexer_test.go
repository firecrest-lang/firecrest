package main

import "testing"

func TestScanTokens(t *testing.T) {
	source := `var age = 42;
              print "Hello, world!";
              while (age > 18) {
                  print "You are over 18!";
                  age = age - 1;
              }
              return "Done!";`
	expectedTokens := []*Token{
		{VAR, "var", nil, 1},
		{IDENTIFIER, "age", nil, 1},
		{EQUAL, "=", nil, 1},
		{NUMBER, "42", 42.0, 1},
		{SEMICOLON, ";", nil, 1},
		{PRINT, "print", nil, 2},
		{STRING, "\"Hello, world!\"", "Hello, world!", 2},
		{SEMICOLON, ";", nil, 2},
		{WHILE, "while", nil, 3},
		{LEFT_PAREN, "(", nil, 3},
		{IDENTIFIER, "age", nil, 3},
		{GREATER, ">", nil, 3},
		{NUMBER, "18", 18.0, 3},
		{RIGHT_PAREN, ")", nil, 3},
		{LEFT_BRACE, "{", nil, 3},
		{PRINT, "print", nil, 4},
		{STRING, "\"You are over 18!\"", "You are over 18!", 4},
		{SEMICOLON, ";", nil, 4},
		{IDENTIFIER, "age", nil, 5},
		{EQUAL, "=", nil, 5},
		{IDENTIFIER, "age", nil, 5},
		{MINUS, "-", nil, 5},
		{NUMBER, "1", 1.0, 5},
		{SEMICOLON, ";", nil, 5},
		{RIGHT_BRACE, "}", nil, 6},
		{RETURN, "return", nil, 7},
		{STRING, "\"Done!\"", "Done!", 7},
		{SEMICOLON, ";", nil, 7},
		{EOF, "", nil, 7},
	}

	lexer := NewLexer(source)

	lexer.scanToken()
	if lexer.HadError {
		t.Errorf("ScanTokens() returned an error")
	}
	if len(lexer.Tokens) != len(expectedTokens) {
		t.Errorf("Expected %d tokens, but got %d", len(expectedTokens), len(lexer.Tokens))
	}
	for i, expected := range expectedTokens {
		actual := lexer.Tokens[i]
		if actual.Type != expected.Type || actual.Lexeme != expected.Lexeme || actual.Literal != expected.Literal || actual.Line != expected.Line {
			t.Errorf("Token %d: expected %v, but got %v", i, expected, actual)
		}
	}
}

func TestLexer_Simple(t *testing.T) {
	source := "var a = 123;"
	expectedTokens := []Token{
		{VAR, "var", nil, 1},
		{IDENTIFIER, "a", nil, 1},
		{EQUAL, "=", nil, 1},
		{NUMBER, "123", 123.0, 1},
		{SEMICOLON, ";", nil, 1},
		{EOF, "", nil, 1},
	}
	lexer := NewLexer(source)

	lexer.scanToken()
	if lexer.HadError {
		t.Errorf("ScanTokens() returned an error")
	}
	if len(lexer.Tokens) != len(expectedTokens) {
		t.Errorf("Expected %d tokens, but got %d", len(expectedTokens), len(lexer.Tokens))
	}
	for i, expected := range expectedTokens {
		actual := lexer.Tokens[i]
		if actual.Type != expected.Type || actual.Lexeme != expected.Lexeme || actual.Literal != expected.Literal || actual.Line != expected.Line {
			t.Errorf("Token %d: expected %v, but got %v", i, expected, actual)
		}
	}
}

func TestLexer_InDepth(t *testing.T) {
	source := `
	class Person {
        construct(name) {
            this.name = name;
        }
        sayHello() {
            print("Hello, my name is " + this.name);
        }
    }

    var p = Person("Nick");
    p.sayHello();
    `
	expectedTokens := []Token{
		{CLASS, "class", nil, 2},
		{IDENTIFIER, "Person", nil, 2},
		{LEFT_BRACE, "{", nil, 2},
		{IDENTIFIER, "construct", nil, 3},
		{LEFT_PAREN, "(", nil, 3},
		{IDENTIFIER, "name", nil, 3},
		{RIGHT_PAREN, ")", nil, 3},
		{LEFT_BRACE, "{", nil, 3},
		{THIS, "this", nil, 4},
		{DOT, ".", nil, 4},
		{IDENTIFIER, "name", nil, 4},
		{EQUAL, "=", nil, 4},
		{IDENTIFIER, "name", nil, 4},
		{SEMICOLON, ";", nil, 4},
		{RIGHT_BRACE, "}", nil, 5},
		{IDENTIFIER, "sayHello", nil, 6},
		{LEFT_PAREN, "(", nil, 6},
		{RIGHT_PAREN, ")", nil, 6},
		{LEFT_BRACE, "{", nil, 6},
		{PRINT, "print", nil, 7},
		{LEFT_PAREN, "(", nil, 7},
		{STRING, "\"Hello, my name is \"", "Hello, my name is ", 7},
		{PLUS, "+", nil, 7},
		{THIS, "this", nil, 7},
		{DOT, ".", nil, 7},
		{IDENTIFIER, "name", nil, 7},
		{RIGHT_PAREN, ")", nil, 7},
		{SEMICOLON, ";", nil, 7},
		{RIGHT_BRACE, "}", nil, 8},
		{RIGHT_BRACE, "}", nil, 9},
		{VAR, "var", nil, 11},
		{IDENTIFIER, "p", nil, 11},
		{EQUAL, "=", nil, 11},
		{IDENTIFIER, "Person", nil, 11},
		{LEFT_PAREN, "(", nil, 11},
		{STRING, "\"Nick\"", "Nick", 11},
		{RIGHT_PAREN, ")", nil, 11},
		{SEMICOLON, ";", nil, 11},
		{IDENTIFIER, "p", nil, 12},
		{DOT, ".", nil, 12},
		{IDENTIFIER, "sayHello", nil, 12},
		{LEFT_PAREN, "(", nil, 12},
		{RIGHT_PAREN, ")", nil, 12},
		{SEMICOLON, ";", nil, 12},
		{EOF, "", nil, 13},
	}
	lexer := NewLexer(source)

	lexer.scanToken()
	if lexer.HadError {
		t.Errorf("ScanTokens() returned an error")
	}
	if len(lexer.Tokens) != len(expectedTokens) {
		t.Errorf("Expected %d tokens, but got %d", len(expectedTokens), len(lexer.Tokens))
	}
	for i, expected := range expectedTokens {
		actual := lexer.Tokens[i]
		if actual.Type != expected.Type || actual.Lexeme != expected.Lexeme || actual.Literal != expected.Literal || actual.Line != expected.Line {
			t.Errorf("Token %d: expected %v, but got %v", i, expected, actual)
		}
	}
}
