package unit

import "testing"

func TestNewValueConstraint(t *testing.T) {
	NewDeclarative(t, "WithNilAndPositiveResult").
		Call(NewValueConstraint, nil, NewMockEqualComparer(t)).
		ExpectResult(ConstraintAsValue{Value: &NilConstraint{}})

	{
		comparator := NewMockEqualComparer(t)
		constraint := NewMockConstraint(t)

		NewDeclarative(t, "WithConstraintAndPositiveResult").
			SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
			Call(NewValueConstraint, constraint, comparator).
			ExpectResult(ConstraintAsValue{Value: constraint})
	}

	{
		comparator := NewMockEqualComparer(t)
		constraint := NewMockConstraint(t)

		NewDeclarative(t, "WithConstraintAsValueAndPositiveResult").
			SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
			Call(NewValueConstraint, ConstraintAsValue{Value: constraint}, comparator).
			ExpectResult(
				ConstraintAsValue{
					Value: NewEqualConstraint(constraint, comparator),
				},
			)
	}

	{
		comparator := NewMockEqualComparer(t)

		NewDeclarative(t, "WithConstraintAsValueAndPositiveResult").
			SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
			Call(NewValueConstraint, "data", comparator).
			ExpectResult(
				ConstraintAsValue{
					NewEqualConstraint("data", comparator),
				},
			)
	}

	NewDeclarative(t, "WithNilComparatorAndNegativeResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewValueConstraint, "data", nil).
		ExpectPanic(NewNotNilError("comparator"))
}
