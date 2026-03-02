package lexer

import (
	"testing"

	"github.com/umid-stha/interpreter/token"
)

/*
TestNextToken tests the nextToken method of the Lexer struct.
It initializes a Lexer with a sample input string and defines a series of expected tokens.
The test iterates through the expected tokens, calling nextToken on the Lexer and comparing the returned token's type and literal value to the expected values.
If there is a mismatch, the test fails with an appropriate error message.
*/
func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, ")"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.nextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
