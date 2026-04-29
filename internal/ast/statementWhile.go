package ast

import "github.com/OJOMB/donkey/internal/tokens"

// StatementWhile represents a while loop statement in the AST. It contains a token, a condition expression, and a body statement block.
type StatementWhile struct {
	Token     tokens.Token
	Condition Expression
	Body      *StatementBlock
}

func (ew *StatementWhile) statementNode()      {}
func (ew *StatementWhile) TokenLexeme() string { return "while" }

func (ew *StatementWhile) String() string {
	return "while " + ew.Condition.String() + " " + ew.Body.String()
}
