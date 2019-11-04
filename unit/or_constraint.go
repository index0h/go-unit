package unit

import (
	"fmt"
	"strings"
)

type OrConstraint struct {
	constraints []Constraint
}

func NewOrConstraint(constraints ...Constraint) Constraint {
	if len(constraints) < 2 {
		err := NewLengthNotLessError("argumentsConstraint", 2, len(constraints))

		panic(err)
	}

	for i, constraint := range constraints {
		if constraint == nil {
			err := NewNotNilError(fmt.Sprintf("argumentsConstraint[%d]", i))

			panic(err)
		}
	}

	result := &OrConstraint{constraints: make([]Constraint, len(constraints))}

	copy(result.constraints, constraints)

	return result
}

func (c *OrConstraint) Check(value interface{}) bool {
	for _, constraint := range c.constraints {
		if constraint.Check(value) {
			return true
		}
	}

	return false
}

func (c *OrConstraint) String() string {
	parts := make([]string, len(c.constraints))

	for i, constraint := range c.constraints {
		parts[i] = constraint.String()
	}

	return c.joinParts(parts)
}

func (c *OrConstraint) Details(value interface{}) string {
	parts := make([]string, len(c.constraints))

	for i, constraint := range c.constraints {
		parts[i] = constraint.Details(value)
	}

	return c.joinParts(parts)
}

func (c *OrConstraint) joinParts(parts []string) string {
	actualParts := make([]string, 0, len(parts))

	for _, part := range parts {
		if part = strings.TrimSpace(part); part != "" {
			actualParts = append(actualParts, part)
		}
	}

	if len(actualParts) == 0 {
		return ""
	}

	if len(actualParts) == 1 {
		return actualParts[0]
	}

	for i, part := range actualParts {
		if strings.Contains(part, "\n") {
			actualParts[i] = "(\n\t" + strings.Replace(part, "\n", "\n\t", -1) + "\n)"
		} else {
			actualParts[i] = "(" + part + ")"
		}
	}

	return strings.Join(actualParts, " or ")
}
