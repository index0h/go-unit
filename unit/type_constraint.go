package unit

import (
	"reflect"
	"strings"
)

type TypeConstraint struct {
	expectedTypes []interface{}
}

func NewTypeConstraint(expectedTypes ...interface{}) Constraint {
	if len(expectedTypes) == 0 {
		err := NewNotEmptyError("expectedTypes")

		panic(err)
	}

	for i, expectedType := range expectedTypes {
		if reflect.ValueOf(expectedType).Kind() == reflect.Invalid {
			err := NewErrorf("Invalid type of expectedTypes[%d]", i)

			panic(err)
		}
	}

	result := &TypeConstraint{
		expectedTypes: make([]interface{}, len(expectedTypes)),
	}

	copy(result.expectedTypes, expectedTypes)

	return result
}

func (c *TypeConstraint) Check(value interface{}) bool {
	valueType := reflect.TypeOf(value)

	for _, expectedType := range c.expectedTypes {
		if reflect.TypeOf(expectedType) == valueType {
			return true
		}
	}

	return false
}

func (c *TypeConstraint) String() string {
	parts := make([]string, len(c.expectedTypes))

	for i, expectedType := range c.expectedTypes {
		parts[i] = reflect.TypeOf(expectedType).String()
	}

	return "have one type of " + strings.Join(parts, ", ")
}

func (c *TypeConstraint) Details(value interface{}) string {
	return "Actual type is " + reflect.TypeOf(value).String()
}
