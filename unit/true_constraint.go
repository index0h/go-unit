package unit

import "reflect"

type TrueConstraint struct {
}

func NewTrueConstraint() Constraint {
	return &TrueConstraint{}
}

func (c *TrueConstraint) Check(value interface{}) bool {
	valueValue := reflect.ValueOf(value)

	if valueValue.Kind() != reflect.Bool {
		panic(NewInvalidKindError("value", value, reflect.Bool))
	}

	return valueValue.Bool()
}

func (c *TrueConstraint) String() string {
	return "be true"
}

func (c *TrueConstraint) Details(value interface{}) string {
	return ""
}
