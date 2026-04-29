package ast

import (
	"github.com/OJOMB/donkey/internal/objects"
	"github.com/OJOMB/donkey/internal/tokens"
)

// StatementFor represents a for loop statement in the AST.
// for (initializer; condition; step) { body }
type StatementFor struct {
	Token       tokens.Token    // The 'for' token
	Initializer Statement       // The initializer statement (e.g., var i = 0)
	Step        Statement       // The step statement (e.g., i = i + 1)
	Condition   Expression      // The condition expression
	Body        *StatementBlock // The body of the loop
}

func (s *StatementFor) statementNode() {}

func (s *StatementFor) TokenLexeme() string {
	return s.Token.Lexeme
}

func (s *StatementFor) String() string {
	result := s.TokenLexeme() + " ("
	if s.Initializer != nil {
		result += s.Initializer.String()
	}

	result += "; "
	if s.Condition != nil {
		result += s.Condition.String()
	}

	result += "; "
	if s.Step != nil {
		result += s.Step.String()
	}

	result += ") "
	if s.Body == nil {
		result += "{}"
		return result
	}

	return result + s.Body.String()
}

func (s *StatementFor) EvalInitializer(env *objects.Environment, evaluator Evaluator) error {
	if s.Initializer == nil {
		return ErrInvalidForLoopInitializer
	}

	result := evaluator.Eval(s.Initializer, env)
	if result == nil {
		return ErrInvalidForLoopInitializer
	}

	return nil
}

func (s *StatementFor) EvalCondition(env *objects.Environment, evaluator Evaluator) (bool, error) {
	if s.Condition == nil {
		return true, ErrInvalidLoopCondition
	}

	conditionValue := evaluator.Eval(s.Condition, env)
	conditionBool, ok := conditionValue.(*objects.Boolean)
	if !ok {
		return false, ErrInvalidLoopConditionType
	}

	return conditionBool.Value, nil
}
