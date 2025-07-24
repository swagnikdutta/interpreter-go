package parser

import (
	"testing"

	"github.com/swagnikdutta/go-interpreter/ast"
	"github.com/swagnikdutta/go-interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// Remember there are three fields in a LetStatement — Token, Name(*Identifier) and Value(Expression) and we treat
	// Name as an Expression, even when it doesn't produce a value. See definition of LetStatement to recall.

	// First we are checking the TokenLiteral corresponding to the s (ast.Statement) Node
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// checking if the underlying type of s is a LetStatement. After all there can be various types of statements.
	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	// I think this is actually matching the variable name, i.e variable x of `let x = 5`
	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not %q. got=%q", name, letStatement.Name.Value)
		return false
	}

	// Here we are checking the TokenLiteral corresponding to the identifier (x of `let x = 5`) — which is treated as
	// an Expression Node.
	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral not %q. got=%q", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
