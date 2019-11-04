package unit

import (
	"fmt"
	"testing"
)

func TestNewEqualConstraint(t *testing.T) {
	comparatorFixture := NewMockEqualComparer(t)

	NewDeclarative(t, "WithComparator").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewEqualConstraint, "data", comparatorFixture).
		ExpectResult(ConstraintAsValue{Value: &EqualConstraint{expected: "data", comparator: comparatorFixture}})

	NewDeclarative(t, "WithNilComparator").
		Call(NewEqualConstraint, "data", nil).
		ExpectPanic(NewNotNilError("comparator"))
}

func TestEqualConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&EqualConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", true),
			}).Check,
			"data",
		).
		ExpectResult(true)

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&EqualConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", false),
			}).Check,
			"data",
		).
		ExpectResult(false)
}

func TestEqualConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&EqualConstraint{
				expected:   "expected",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", true),
			}).String,
		).
		ExpectResult(fmt.Sprintf("be equal to %+v", "expected"))
}

func TestEqualConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&EqualConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordDiff("data", "data", "diff"),
			}).Details,
			"data",
		).
		ExpectResult("diff")
}
