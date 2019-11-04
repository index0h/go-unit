package unit

import (
	"fmt"
	"reflect"
	"strings"
)

type KindConstraint struct {
	expectedKinds []reflect.Kind
}

func NewKindConstraint(expectedKinds ...reflect.Kind) Constraint {
	if len(expectedKinds) == 0 {
		panic(NewNotEmptyError("expectedKinds"))
	}

	result := &KindConstraint{expectedKinds: make([]reflect.Kind, len(expectedKinds))}

	copy(result.expectedKinds, expectedKinds)

	return result
}

func (c *KindConstraint) Check(value interface{}) bool {
	valueKind := reflect.TypeOf(value).Kind()
	parts := make([]string, len(c.expectedKinds))

	for i, expectedKind := range c.expectedKinds {
		if expectedKind == valueKind {
			return true
		}

		parts[i] = expectedKind.String()
	}

	return false
}

func (c *KindConstraint) String() string {
	parts := make([]string, len(c.expectedKinds))

	for i, expectedKind := range c.expectedKinds {
		parts[i] = expectedKind.String()
	}

	return "have one kind of " + strings.Join(parts, ", ")
}

func (c *KindConstraint) Details(actual interface{}) string {
	return fmt.Sprintf("Actual kind is %s", reflect.ValueOf(actual).Kind().String())
}
