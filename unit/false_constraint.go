package unit

import (
	"reflect"
)

type FalseConstraint struct {
}

func NewFalseConstraint() Constraint {
	return &FalseConstraint{}
}

func (c *FalseConstraint) Check(value interface{}) bool {
	valueValue := reflect.ValueOf(value)

	if valueValue.Kind() != reflect.Bool {
		panic(NewInvalidKindError("value", value, reflect.Bool))
	}

	return !valueValue.Bool()
}

func (c *FalseConstraint) String() string {
	return "be false"
}

func (c *FalseConstraint) Details(value interface{}) string {
	return ""
}
