package unit

import (
	"fmt"
	"reflect"
)

type LengthGreaterConstraint struct {
	length int
}

func NewLengthGreaterConstraint(length int) Constraint {
	return &LengthGreaterConstraint{
		length: length,
	}
}

func (c *LengthGreaterConstraint) Check(value interface{}) bool {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return listValue.Len() > c.length
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}

func (c *LengthGreaterConstraint) String() string {
	return fmt.Sprintf("have length greater than %v", c.length)
}

func (c *LengthGreaterConstraint) Details(value interface{}) string {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return fmt.Sprintf("Actual length is %v", listValue.Len())
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}
