package unit

import "fmt"

type SameConstraint struct {
	expected   interface{}
	comparator EqualComparer
}

func NewSameConstraint(value interface{}) Constraint {
	return &SameConstraint{
		expected:   value,
		comparator: NewEqualComparator(SameTypeOption{Value: true}, SamePointerOption{Value: true}),
	}
}

func (c *SameConstraint) Check(actual interface{}) bool {
	return c.comparator.Compare(c.expected, actual)
}

func (c *SameConstraint) String() string {
	return fmt.Sprintf("be same as %+v", c.expected)
}

func (c *SameConstraint) Details(actual interface{}) string {
	return c.comparator.Diff(c.expected, actual)
}
