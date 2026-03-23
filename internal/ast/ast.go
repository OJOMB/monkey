package ast

import "github.com/OJOMB/monkey/internal/tokens"

// Node is the base interface for all nodes in the AST.
// It has a method TokenLexeme() that returns the lexeme of the token associated with the node.
type Node interface {
	// TokenLexeme returns the lexeme of the token associated with the node.
	TokenLexeme() string
}

// Statement represents a statement in the AST.
// It embeds the Node interface and has an additional method statementNode() that is used to differentiate it from expressions.
type Statement interface {
	Node
	// statementNode is a marker method to indicate that a struct is a Statement.
	statementNode()
}

// Expression represents an expression in the AST.
// It embeds the Node interface and has an additional method expressionNode() that is used to differentiate it from statements.
type Expression interface {
	Node
	// expressionNode is a marker method to indicate that a struct is an Expression.
	expressionNode()
}

// Program is the root node of the AST. It contains a slice of statements.
type Program struct {
	Statements []Statement
}

// TokenLexeme returns the lexeme of the first statement's token in the program, or an empty string if there are no statements.
func (p *Program) TokenLexeme() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLexeme()
	}

	return ""
}

// LetStatement represents a let statement in the AST. It contains a token, an identifier for the variable name, and an expression for the value.
type LetStatement struct {
	// Token is the token associated with the let statement, which is typically a token.TokenTypeLet token.
	Token tokens.Token
	// Name is the identifier for the variable name being declared in the let statement.
	// LHS of the let statement, which is the variable name being declared.
	Name *Identifier
	// Value is the expression that represents the value being assigned to the variable in the let statement.
	// RHS of the let statement, which is the expression representing the value being assigned to the variable.
	Value Expression
}

func (ls *LetStatement) statementNode()      {}
func (ls *LetStatement) TokenLexeme() string { return ls.Token.Lexeme }

// Identifier represents an identifier in the AST. It contains a token and a value for the identifier name.
type Identifier struct {
	// Token is the token associated with the identifier, which is typically a token.TokenTypeIdent token.
	Token tokens.Token
	// Value is the name of the identifier, which is the lexeme of the token.
	Value string
}

func (i *Identifier) expressionNode()     {}
func (i *Identifier) TokenLexeme() string { return i.Token.Lexeme }
