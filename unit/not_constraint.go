package unit

import (
	"fmt"
	"strings"
)

type NotConstraint struct {
	constraint Constraint
}

func NewNotConstraint(constraint Constraint) Constraint {
	if constraint == nil {
		panic(NewNotNilError("constraint"))
	}

	return &NotConstraint{
		constraint: constraint,
	}
}

func (c *NotConstraint) Check(actual interface{}) bool {
	return !c.constraint.Check(actual)
}

func (c *NotConstraint) String() string {
	switch c.constraint.(type) {
	case *AndConstraint:
		return "be not " + c.formatMultilineString(c.constraint.String())
	case *AnyConstraint:
		return "be nothing"
	case *CallbackConstraint:
		return "not accept callback"
	case *ContainsConstraint:
		return fmt.Sprintf("not contain %+v", c.constraint.(*ContainsConstraint).element)
	case *EmptyConstraint:
		return "be not empty"
	case *EqualConstraint:
		return fmt.Sprintf("be not equal to %+v", c.constraint.(*EqualConstraint).expected)
	case *FalseConstraint:
		return "be true"
	case *GreaterConstraint:
		return fmt.Sprintf("be less than or equal to %v", c.constraint.(*GreaterConstraint).expected)
	case *KindConstraint:
		constraint := c.constraint.(*KindConstraint)
		parts := make([]string, len(constraint.expectedKinds))

		for i, expectedKind := range constraint.expectedKinds {
			parts[i] = expectedKind.String()
		}

		return fmt.Sprintf("have no one kind of %s", strings.Join(parts, ", "))
	case *LengthConstraint:
		return fmt.Sprintf("not have length %v", c.constraint.(*LengthConstraint).length)
	case *LengthGreaterConstraint:
		return fmt.Sprintf("have length less than or equal to %v", c.constraint.(*LengthGreaterConstraint).length)
	case *LengthLessConstraint:
		return fmt.Sprintf("have length greater than or equal to %v", c.constraint.(*LengthLessConstraint).length)
	case *LessConstraint:
		return fmt.Sprintf("be greater than or equal to %v", c.constraint.(*LessConstraint).expected)
	case *NilConstraint:
		return "be not nil"
	case *NotConstraint:
		return c.constraint.(*NotConstraint).constraint.String()
	case *OrConstraint:
		return "be not " + c.formatMultilineString(c.constraint.String())
	case *RegexpConstraint:
		return fmt.Sprintf("not match PCRE pattern '%s'", c.constraint.(*RegexpConstraint).pattern.String())
	case *SameConstraint:
		return fmt.Sprintf("not be same as %+v", c.constraint.(*SameConstraint).expected)
	case *TrueConstraint:
		return "be false"
	case *TypeConstraint:
		constraint := c.constraint.(*TypeConstraint)
		parts := make([]string, len(constraint.expectedTypes))

		for i, expectedType := range constraint.expectedTypes {
			parts[i] = fmt.Sprintf("%T", expectedType)
		}

		return fmt.Sprintf("have no one type of %v", strings.Join(parts, ", "))
	default:
		return "not " + c.formatMultilineString(c.constraint.String())
	}
}

func (c *NotConstraint) Details(value interface{}) string {
	return c.constraint.Details(value)
}

func (c *NotConstraint) formatMultilineString(input string) string {
	input = strings.TrimSpace(input)

	if strings.Index(input, "\n") != -1 {
		return "(\n\t" + strings.Replace(input, "\n", "\n\t", -1) + "\n)"
	} else {
		return "(" + input + ")"
	}
}
