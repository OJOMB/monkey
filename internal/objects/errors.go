package objects

// ErrorValue is a struct that represents an error value in the Donkey programming language.
// It contains a single field, Message, which is a string that describes the error.
type ErrorValue struct {
	Message string
}

// Type returns the type of the ErrorValue object, which is always TypeError.
func (ev *ErrorValue) Type() Type {
	return TypeErrorValue
}

// Inspect returns a string representation of the ErrorValue object, which includes the error message.
func (ev *ErrorValue) Inspect() string {
	return "ERROR: " + ev.Message
}
