package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTERGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (i *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%v", i.Value)
}

type Null struct {
	Value bool
}

func (i *Null) Type() ObjectType {
	return NULL_OBJ
}

func (i *Null) Inspect() string {
	return "NULL"
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	store map[string]Object
}
