package unit

import (
	"fmt"
	"math"
	"reflect"
)

type LessConstraint struct {
	expected interface{}
}

func NewLessConstraint(expected interface{}) Constraint {
	switch reflect.ValueOf(expected).Kind() {
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		return &LessConstraint{expected: expected}
	}

	err := NewInvalidNumericComparisonTypeError("expectedKinds", expected)

	panic(err)
}

func (c *LessConstraint) Check(actual interface{}) bool {
	expectedValue := reflect.ValueOf(c.expected)
	actualValue := reflect.ValueOf(actual)

	switch expectedValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch actualValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return expectedValue.Int() > actualValue.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			expectedInt := expectedValue.Int()
			actualUint := actualValue.Uint()

			if actualUint <= math.MaxInt64 {
				return expectedInt > int64(actualUint)
			}

			if expectedInt >= 0 {
				return uint64(expectedInt) > actualUint
			}

			return float64(expectedInt) > float64(actualUint)
		case reflect.Float32, reflect.Float64:
			return float64(expectedValue.Int()) > actualValue.Float()
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch actualValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			xUint := expectedValue.Uint()
			yInt := actualValue.Int()

			if xUint <= math.MaxInt64 {
				return int64(xUint) > yInt
			}

			if yInt >= 0 {
				return xUint > uint64(yInt)
			}

			return float64(xUint) > float64(yInt)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return expectedValue.Uint() > actualValue.Uint()
		case reflect.Float32, reflect.Float64:
			return float64(expectedValue.Uint()) > actualValue.Float()
		}
	case reflect.Float32, reflect.Float64:
		switch actualValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return expectedValue.Float() > float64(actualValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return expectedValue.Float() > float64(actualValue.Uint())
		case reflect.Float32, reflect.Float64:
			return expectedValue.Float() > actualValue.Float()
		}
	}

	err := NewInvalidNumericComparisonTypeError("actual", actual)

	panic(err)
}

func (c *LessConstraint) String() string {
	return fmt.Sprintf("be less than %v", c.expected)
}

func (c *LessConstraint) Details(actual interface{}) string {
	return ""
}
