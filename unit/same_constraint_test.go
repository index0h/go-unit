package unit

import (
	"fmt"
	"testing"
)

func TestNewSameConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewSameConstraint, "data").
		ExpectResult(
			ConstraintAsValue{
				Value: &SameConstraint{
					expected:   "data",
					comparator: NewEqualComparator(SamePointerOption{Value: true}, SameTypeOption{Value: true}),
				},
			},
		)
}

func TestSameConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&SameConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", true),
			}).Check,
			"data",
		).
		ExpectResult(true)

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&SameConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", false),
			}).Check,
			"data",
		).
		ExpectResult(false)
}

func TestSameConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&SameConstraint{
				expected:   "expected",
				comparator: NewMockEqualComparer(t).RecordCompare("data", "data", true),
			}).String,
		).
		ExpectResult(fmt.Sprintf("be same as %+v", "expected"))
}

func TestSameConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&SameConstraint{
				expected:   "data",
				comparator: NewMockEqualComparer(t).RecordDiff("data", "data", "diff"),
			}).Details,
			"data",
		).
		ExpectResult("diff")
}
