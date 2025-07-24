package parser

import (
	"github.com/swagnikdutta/go-interpreter/ast"
	"github.com/swagnikdutta/go-interpreter/lexer"
	"github.com/swagnikdutta/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so that curToken and peekToken both are set.
	p.nextToken()
	p.nextToken()
	return p
}

// This function has to be called twice in order to set both curToken and peekToken.
// In the first call, curToken will be nil, peekToken will point to the first token.
// In the second call, curToken will point to the first token and peekToken will point to the second token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
