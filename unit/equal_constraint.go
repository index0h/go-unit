package unit

import (
	"fmt"
)

type EqualConstraint struct {
	expected   interface{}
	comparator EqualComparer
}

func NewEqualConstraint(expected interface{}, comparator EqualComparer) Constraint {
	if comparator == nil {
		panic(NewNotNilError("comparator"))
	}

	return &EqualConstraint{
		expected:   expected,
		comparator: comparator,
	}
}

func (c *EqualConstraint) Check(actual interface{}) bool {
	return c.comparator.Compare(c.expected, actual)
}

func (c *EqualConstraint) String() string {
	return fmt.Sprintf("be equal to %+v", c.expected)
}

func (c *EqualConstraint) Details(actual interface{}) string {
	return c.comparator.Diff(c.expected, actual)
}
