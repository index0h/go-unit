package unit

import "testing"

func TestNewElementsConstraint(t *testing.T) {
	constraint1 := NewMockConstraint(t)
	constraint2 := NewMockConstraint(t)

	NewSubtest(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewElementsConstraint, constraint1, constraint2).
		ExpectResult(
			ConstraintAsValue{
				Value: &ElementsConstraint{
					constraints: []Constraint{constraint1, constraint2},
				},
			},
		)

	NewSubtest(t, "WithNilConstraint").
		Call(NewElementsConstraint, nil).
		ExpectPanic(NewNotNilError("constraint[0]"))
}

func TestNewValueElementsConstraint(t *testing.T) {
	constraint1 := NewMockConstraint(t)
	constraint2 := NewMockConstraint(t)
	comparator := NewMockEqualComparer(t)

	NewSubtest(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewValueElementsConstraint, comparator, nil, constraint1, ConstraintAsValue{Value: constraint2}, "data").
		ExpectResult(
			ConstraintAsValue{
				Value: &ElementsConstraint{
					constraints: []Constraint{
						&NilConstraint{},
						constraint1,
						&EqualConstraint{expected: constraint2, comparator: comparator},
						&EqualConstraint{expected: "data", comparator: comparator},
					},
				},
			},
		)

	NewSubtest(t, "WithNilComparator").
		Call(NewValueElementsConstraint, nil).
		ExpectPanic(NewNotNilError("comparator"))
}

func TestElementsConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck("data1", true),
					NewMockConstraint(t).RecordCheck("data2", true),
				},
			}).Check,
			[]interface{}{"data1", "data2"},
		).
		ExpectResult(true)

	NewSubtest(t, "WithNegativeResult").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck("data1", false),
					NewMockConstraint(t),
				},
			}).Check,
			[]interface{}{"data1", "data2"},
		).
		ExpectResult(false)

	NewSubtest(t, "WithInvalidList").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).Check,
			"data",
		).
		ExpectPanic(NewInvalidTypeError("list", "data", []interface{}{}))

	NewSubtest(t, "WithInvalidElementsCount").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t),
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).Check,
			[]interface{}{"data1"},
		).
		ExpectPanic(NewInvalidLengthError("list", 3, 1))
}

func TestElementsConstraint_String(t *testing.T) {
	NewSubtest(t, "WithInvalidElementsCount").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t),
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).String,
		).
		ExpectResult("apply all constraints for each elements")
}

func TestElementsConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithInvalidElementsCount").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck("data1", true),
					NewMockConstraint(t).
						RecordCheck("data2", false).
						RecordDetails("data2", "details2").
						RecordString("string2"),
					NewMockConstraint(t).
						RecordCheck("data3", false).
						RecordDetails("data3", "").
						RecordString("string3"),
				},
			}).Details,
			[]interface{}{"data1", "data2", "data3"},
		).
		ExpectResult("element[1], it must string2.\ndetails2\nelement[2], it must string3, actual value is data3.")

	NewSubtest(t, "WithInvalidList").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).Details,
			"data",
		).
		ExpectPanic(NewInvalidTypeError("list", "data", []interface{}{}))

	NewSubtest(t, "WithInvalidElementsCount").
		Call(
			(&ElementsConstraint{
				constraints: []Constraint{
					NewMockConstraint(t),
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).Details,
			[]interface{}{"data1"},
		).
		ExpectPanic(NewInvalidLengthError("list", 3, 1))
}
