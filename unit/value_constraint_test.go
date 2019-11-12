package unit

import "testing"

func TestNewValueConstraint(t *testing.T) {
	NewSubtest(t, "WithNilAndPositiveResult").
		Call(NewValueConstraint, nil, NewMockEqualComparer(t)).
		ExpectResult(ConstraintAsValue{Value: &NilConstraint{}})

	{
		comparator := NewMockEqualComparer(t)
		constraint := NewMockConstraint(t)

		NewSubtest(t, "WithConstraintAndPositiveResult").
			SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
			Call(NewValueConstraint, constraint, comparator).
			ExpectResult(ConstraintAsValue{Value: constraint})
	}

	{
		comparator := NewMockEqualComparer(t)
		constraint := NewMockConstraint(t)

		NewSubtest(t, "WithConstraintAsValueAndPositiveResult").
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

		NewSubtest(t, "WithConstraintAsValueAndPositiveResult").
			SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
			Call(NewValueConstraint, "data", comparator).
			ExpectResult(
				ConstraintAsValue{
					NewEqualConstraint("data", comparator),
				},
			)
	}

	NewSubtest(t, "WithNilComparatorAndNegativeResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewValueConstraint, "data", nil).
		ExpectPanic(NewNotNilError("comparator"))
}
