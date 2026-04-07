package object

type Nowt struct{}

func (n *Nowt) Type() ObjectType { return TypeNowt }
func (n *Nowt) Inspect() string  { return "nowt" }
