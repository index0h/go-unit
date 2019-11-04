package unit

import (
	"fmt"
	"reflect"
)

type LengthLessConstraint struct {
	length int
}

func NewLengthLessConstraint(length int) Constraint {
	return &LengthLessConstraint{
		length: length,
	}
}

func (c *LengthLessConstraint) Check(value interface{}) bool {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return listValue.Len() < c.length
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}

func (c *LengthLessConstraint) String() string {
	return fmt.Sprintf("have length less than %v", c.length)
}

func (c *LengthLessConstraint) Details(value interface{}) string {
	listValue := reflect.ValueOf(value)

	switch listValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return fmt.Sprintf("Actual length is %v", listValue.Len())
	default:
		panic(NewInvalidLengthComparisonTypeError("value", value))
	}
}
