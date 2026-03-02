package lexer

import (
	"github.com/umid-stha/interpreter/token"
)

/*
	Goal:
	Achive somthing like this:
	[
		LET,
		IDENTIFIER("x"),
		EQUAL_SIGN,
		INTEGER(5),
		PLUS_SIGN,
		INTEGER(5),
		SEMICOLON
	]

*/

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination. uses byte for simplicity, but could be extended to support Unicode characters using rune.
}

// initialize the lexer with the input and read the first character
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// readChar advances our position in the input string or sets the current character to 0 if we've reached the end of the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

/*
readIdentifer reads rest of the identifiers/keywords unitil it encounters a non-letter character from the input,
and returns slice of the origin input as token literal
*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readInt is same as readIdentifier for numbers
func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhiteSpace ignores all whitespaces
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}

// return the char at readPostion
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

/*
nextToken checks the current character and returns the appropriate token based on it.
It uses a switch statement to determine the type of token to return.
After determining the token, it calls readChar to advance to the next character in the input.
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readInt()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// newToken is a helper function that creates a new token based on the given token type and character.
func newToken[C string | byte](tokenType token.TokenType, ch C) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// checks if the character is valid letter including '_'
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// checks if the character is valid number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
