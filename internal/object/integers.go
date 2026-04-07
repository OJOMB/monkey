package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return TypeInteger }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
