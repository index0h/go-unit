package unit

import (
	"fmt"
	"reflect"
)

type LengthConstraint struct {
	length int
}

func NewLengthConstraint(length int) Constraint {
	return &LengthConstraint{
		length: length,
	}
}

func (c *LengthConstraint) Check(value interface{}) bool {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return listValue.Len() == c.length
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}

func (c *LengthConstraint) String() string {
	return fmt.Sprintf("have length %v", c.length)
}

func (c *LengthConstraint) Details(value interface{}) string {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return fmt.Sprintf("Actual length is %v", listValue.Len())
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}
