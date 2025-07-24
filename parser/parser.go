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

// First thing parse program does is construct the root node of the AST — an *ast.Program (Recall that a Program is
// a Node, see definition). It then iterates over every token in the input until it encounters a token.EOF.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// It's like we are reading the first token, and generating a Statement
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// Till here, we have read til '=' of 'let x = 5'
	// For now, we will skip expressions until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}
