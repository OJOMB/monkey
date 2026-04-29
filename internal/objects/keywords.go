package objects

// Continue represents a continue statement in the Donkey programming language. It is used to skip the current iteration of a loop and continue with the next iteration.
type Continue struct{}

// Type returns the type of the Continue object, which is "CONTINUE".
func (c *Continue) Type() Type { return TypeContinue }

// Inspect returns a string representation of the Continue object, which is "continue".
func (c *Continue) Inspect() string { return "continue" }

// Break represents a break statement in the Donkey programming language. It is used to exit a loop immediately.
type Break struct{}

// Type returns the type of the Break object, which is "BREAK".
func (b *Break) Type() Type { return TypeBreak }

// Inspect returns a string representation of the Break object, which is "break".
func (b *Break) Inspect() string { return "break" }
