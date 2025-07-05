package lexer

import (
	"github.com/swagnikdutta/go-interpreter/token"
)

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpaces()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
			// when l.ch is not one of the recognized characters, it might be an identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // returning from here because we move the current pointer ahead by more than one step
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// if we end up here, we truly don't know how to handle the current character and
			// declare the token as illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// advancing the ch pointer
	l.readChar()
	return tok
}

// The purpose of readChar is to give us the next character and
// advance our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code for "NUL" character and signifies "we haven't read anything yet",
		// or "end of file" for us.
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

	// s := ""
	// for isLetter(l.ch) {
	// 	s += string(l.ch)
	// 	l.readChar()
	// }
	// return s
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// calling readChar for the first time sets readPosition to 1 — otherwise, these values are set to zero.
	l.readChar()
	return l
}
