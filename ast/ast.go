package ast

import "github.com/swagnikdutta/go-interpreter/token"

// Node — Every node in our AST has to implement the node interface, meaning
// it has to provide a TokenLiteral() method that returns the literal value
// of the token it is associated with.
//
// The AST we are building consists solely of Nodes — that are connected to each
// other (it's a tree after all). Some of these nodes implement the Statement
// and some the Expression interface.
//
// Statement and Expression interfaces contain dummy methods — statementNode(),
// and expressionNode() — which act as guiding constructs/safety checks — to
// throw errors when we use a Statement where an Expression was required and
// vice versa.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program — A program in Monkey is a series of statements.
// Also note that, it implements TokenLiteral() below, which means Program is a Node.
type Program struct {
	// Statements are just a slice of AST Nodes that implement the Statement interface.
	Statements []Statement
}

// TokenLiteral — The program node is going to be the root node in our AST.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// This enforces that the node representing a let statement — is a statement and not an expression.
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier — To hold an identifier in a binding, i.e, the x in `let x = 5`, we are using this Identifier struct type.
// It implements the Expression interface, as you can see, it has implementations for expressionNode() and TokenLiteral().
//
// But an identifier in a let statement does not produce any value — while Expressions do.
// Then why is it being treated as an Expression?
// - To keep things simple.
// - Because, identifiers in other parts of Monkey program `let x = valueProducingIdentifier` do produce values.
// - We want to reduce the number of different node types.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// We are treating identifier as an expression. This might feel weird because,
// the identifier in a let statement doesn't produce a value.
// So why will it be an expression?
// To keep things simple.
// Besides, identifiers in other parts of a Monkey program do produce values: let x = valueProducingIdentifier
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
