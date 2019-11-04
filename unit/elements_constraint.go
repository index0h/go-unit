package unit

import (
	"fmt"
	"strings"
)

type ElementsConstraint struct {
	constraints []Constraint
}

func NewElementsConstraint(constraints ...Constraint) Constraint {
	result := &ElementsConstraint{
		constraints: make([]Constraint, len(constraints)),
	}

	for i, constraint := range constraints {
		if constraint == nil {
			panic(NewNotNilError(fmt.Sprintf("constraint[%d]", i)))
		}

		result.constraints[i] = constraint
	}

	return result
}

func NewValueElementsConstraint(comparator EqualComparer, values ...interface{}) Constraint {
	if comparator == nil {
		panic(NewNotNilError("comparator"))
	}

	result := &ElementsConstraint{
		constraints: make([]Constraint, len(values)),
	}

	for i, value := range values {
		result.constraints[i] = NewValueConstraint(value, comparator)
	}

	return result
}

func (c *ElementsConstraint) Check(list interface{}) bool {
	elements, ok := list.([]interface{})

	if !ok {
		panic(NewInvalidTypeError("list", list, []interface{}{}))
	}

	actualLength := len(elements)
	expectedLength := len(c.constraints)

	if actualLength != expectedLength {
		panic(NewInvalidLengthError("list", expectedLength, actualLength))
	}

	for i, value := range elements {
		if !c.constraints[i].Check(value) {
			return false
		}
	}

	return true
}

func (c *ElementsConstraint) String() string {
	return "apply all constraints for each elements"
}

func (c *ElementsConstraint) Details(list interface{}) string {
	elements, ok := list.([]interface{})

	if !ok {
		panic(NewInvalidTypeError("list", list, []interface{}{}))
	}

	actualLength := len(elements)
	expectedLength := len(c.constraints)

	if actualLength != expectedLength {
		panic(NewInvalidLengthError("list", expectedLength, actualLength))
	}

	result := ""

	for i, value := range elements {
		if !c.constraints[i].Check(value) {
			if details := c.constraints[i].Details(value); details != "" {
				result += fmt.Sprintf(
					"element[%d], it must %s.\n%s\n",
					i,
					c.constraints[i].String(),
					details,
				)
			} else {
				result += fmt.Sprintf(
					"element[%d], it must %s, actual value is %+v.\n",
					i,
					c.constraints[i].String(),
					value,
				)
			}
		}
	}

	return strings.TrimSpace(result)
}
