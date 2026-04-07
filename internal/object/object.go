package object

type ObjectType string

const (
	// Object types

	// TypeInteger represents the type of an integer object
	TypeInteger = "INTEGER"
	// TypeBoolean represents the type of a boolean object
	TypeBoolean = "BOOLEAN"
	// TypeNowt represents the type of a nowt object (the equivalent of null or nil)
	TypeNowt = "NOWT"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
