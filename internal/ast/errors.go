package ast

import "fmt"

var (
	ErrInvalidForLoopInitializer = fmt.Errorf("invalid for loop initializer: expected a statement")
	ErrInvalidForLoopStep        = fmt.Errorf("invalid for loop step: expected a statement")
	ErrInvalidLoopCondition      = fmt.Errorf("invalid for loop condition: expected an expression")
	ErrInvalidLoopConditionType  = fmt.Errorf("invalid for loop condition: expected a boolean expression")
)
