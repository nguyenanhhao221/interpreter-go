package lexer

import (
	"interpreter-go/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken Actually read the token and assign it then call readChar to update the position
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			/* readIdentifier() to set the Literal field of our current token.
			But what about its Type? Now that we have read identifiers like let, fn or foobar,
			we need to be able to tell user-defined identifiers apart from language keywords.
			We need a function that returns the correct TokenType for the token literal we have.
			What better place than the token package to add such a function? */
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func isDigit(char byte) bool {
	// Check if the ASCII value of char (byte) is in the range of ASCII value of the number 0 and 9
	// ASCII is the standard, it basically mean that each character in the alphabet is present as an integer
	// For example the number 0 is represent as the value 48 and 9 is 57
	// The single quote '' in this case ('0') and in go lang mean that Hey get the ASCII value of the number 0
	return '0' <= char && char <= '9'
}

func (l *Lexer) readNumber() string {
	currentPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[currentPosition:l.position]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// readIdentifier Read until the char is no longer a character, update position
func (l *Lexer) readIdentifier() string {
	currentPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[currentPosition:l.position]
}

// skipWhiteSpace if the character is white space skip to the next position
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
