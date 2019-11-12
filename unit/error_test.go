package unit

import (
	"reflect"
	"testing"
)

func TestNewAssertError(t *testing.T) {
	NewSubtest(t, "WithDetails").
		Call(NewAssertError, 5, "context", "constraint", "details").
		ExpectResult(&Error{message: "Failed assertion for context, it must constraint.\ndetails"})

	NewSubtest(t, "WithoutDetails").
		Call(NewAssertError, 5, "context", "constraint", "").
		ExpectResult(&Error{message: "Failed assertion for context, it must constraint, actual value is 5."})
}

func TestNewNotNilError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewNotNilError, "variable").
		ExpectResult(&Error{message: "Variable 'variable' must be not nil"})
}

func TestNewNotEmptyError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewNotEmptyError, "variable").
		ExpectResult(&Error{message: "Variable 'variable' must be not empty"})
}

func TestNewErrorf(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewErrorf, "data %T", [1]int{5}).
		ExpectResult(&Error{message: "data [1]int"})
}

func TestNewInvalidKindError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewInvalidKindError, "variable", "data", reflect.Int, reflect.Bool).
		ExpectResult(&Error{message: "Variable 'variable' (string) must have one of kinds: int, bool"})
}

func TestNewInvalidTypeError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewInvalidTypeError, "variable", "data", "", new(int)).
		ExpectResult(&Error{message: "Variable 'variable' (string) must have one of types: string, *int"})
}

func TestNewInvalidNumericComparisonTypeError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewInvalidNumericComparisonTypeError, "variable", func() {}).
		ExpectResult(&Error{message: "Variable 'variable' (func()) is not acceptable for numeric comparison"})
}

func TestNewInvalidLengthComparisonTypeError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewInvalidLengthComparisonTypeError, "variable", func() {}).
		ExpectResult(&Error{message: "Variable 'variable' (func()) is not acceptable for length comparison"})
}

func TestNewInvalidNilComparisonTypeError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewInvalidNilComparisonTypeError, "variable", func() {}).
		ExpectResult(&Error{message: "Variable 'variable' (func()) is not acceptable for nil comparison"})
}

func TestNewLengthNotLessError(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewLengthNotLessError, "variable", 10, 5).
		ExpectResult(
			&Error{message: "Variable 'variable' must have length greater than or equal to 10, actual length: 5"},
		)
}

func TestError_Error(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&Error{message: "data"}).Error).
		ExpectResult("data")
}
