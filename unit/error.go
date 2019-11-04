package unit

import (
	"fmt"
	"reflect"
	"strings"
)

type Error struct {
	message string
}

func NewAssertError(value interface{}, context string, assertion string, details string) *Error {
	if details == "" {
		return &Error{
			message: fmt.Sprintf(
				"Failed assertion for %s, it must %s, actual value is %+v.",
				context,
				assertion,
				value,
			),
		}
	}

	return &Error{
		message: fmt.Sprintf("Failed assertion for %s, it must %s.\n%s", context, assertion, details),
	}
}

func NewNotNilError(name string) *Error {
	return &Error{message: fmt.Sprintf("Variable '%s' must be not nil", name)}
}

func NewNotEmptyError(name string) *Error {
	return &Error{message: fmt.Sprintf("Variable '%s' must be not empty", name)}
}

func NewErrorf(format string, arguments ...interface{}) error {
	return &Error{message: strings.TrimSpace(fmt.Sprintf(format, arguments...))}
}

func NewInvalidKindError(name string, value interface{}, expectedKinds ...reflect.Kind) *Error {
	parts := make([]string, len(expectedKinds))

	for i, expectedKind := range expectedKinds {
		parts[i] = expectedKind.String()
	}

	return &Error{
		message: fmt.Sprintf(
			"Variable '%s' (%T) must have one of kinds: %s",
			name,
			value,
			strings.Join(parts, ", "),
		),
	}
}

func NewInvalidTypeError(name string, value interface{}, expectedTypes ...interface{}) *Error {
	parts := make([]string, len(expectedTypes))

	for i, expectedType := range expectedTypes {
		parts[i] = fmt.Sprintf("%T", expectedType)
	}

	return &Error{
		message: fmt.Sprintf(
			"Variable '%s' (%T) must have one of types: %s",
			name,
			value,
			strings.Join(parts, ", "),
		),
	}
}

func NewInvalidNumericComparisonTypeError(name string, value interface{}) *Error {
	return &Error{
		message: fmt.Sprintf("Variable '%s' (%T) is not acceptable for numeric comparison", name, value),
	}
}

func NewInvalidLengthComparisonTypeError(name string, value interface{}) *Error {
	return &Error{
		message: fmt.Sprintf("Variable '%s' (%T) is not acceptable for length comparison", name, value),
	}
}

func NewInvalidNilComparisonTypeError(name string, value interface{}) *Error {
	return &Error{
		message: fmt.Sprintf("Variable '%s' (%T) is not acceptable for nil comparison", name, value),
	}
}

func NewLengthNotLessError(variableName string, expectedLength int, actualLength int) *Error {
	return &Error{
		message: fmt.Sprintf(
			"Variable '%s' must have length greater than or equal to %d, actual length: %d",
			variableName,
			expectedLength,
			actualLength,
		),
	}
}

func NewInvalidLengthError(variableName string, expectedLength int, actualLength int) *Error {
	return &Error{
		message: fmt.Sprintf(
			"Variable '%s' must have length %d, actual length: %d",
			variableName,
			expectedLength,
			actualLength,
		),
	}
}

func (e *Error) Error() string {
	return e.message
}
