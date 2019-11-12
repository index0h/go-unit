package unit

import (
	"testing"
)

func TestNewOrConstraint(t *testing.T) {
	constraint1 := NewMockConstraint(t)
	constraint2 := NewMockConstraint(t)
	constraint3 := NewMockConstraint(t)

	NewSubtest(t, "Success: (<Value>, <Value>)").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockConstraint{}}).
		Call(NewOrConstraint, constraint1, constraint2).
		ExpectResult(ConstraintAsValue{Value: &OrConstraint{constraints: []Constraint{constraint1, constraint2}}})

	NewSubtest(t, "Success: (<Value>, <Value>, <Value>)").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockConstraint{}}).
		Call(NewOrConstraint, constraint1, constraint2, constraint3).
		ExpectResult(ConstraintAsValue{Value: &OrConstraint{constraints: []Constraint{constraint1, constraint2, constraint3}}})

	NewSubtest(t, "Panic: ()").
		Call(NewOrConstraint).
		ExpectPanic(NewLengthNotLessError("argumentsConstraint", 2, 0))

	NewSubtest(t, "Panic: (<Value>)").
		Call(NewOrConstraint, constraint1).
		ExpectPanic(NewLengthNotLessError("argumentsConstraint", 2, 1))

	NewSubtest(t, "Panic: (<Value>, nil, <Value>)").
		Call(NewOrConstraint, constraint1, nil, constraint3).
		ExpectPanic(NewNotNilError("argumentsConstraint[1]"))
}

func TestOrConstraint_Check(t *testing.T) {
	NewSubtest(t, "Success: (false && false && false) -> false").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck(5, false),
					NewMockConstraint(t).RecordCheck(5, false),
					NewMockConstraint(t).RecordCheck(5, false),
				},
			}).Check,
			5,
		).
		ExpectResult(false)

	NewSubtest(t, "Success: (true && ? && ?) -> true").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck(5, true),
					NewMockConstraint(t),
					NewMockConstraint(t),
				},
			}).Check,
			5,
		).
		ExpectResult(true)

	NewSubtest(t, "Success: (false && true && ?) -> true").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck(5, false),
					NewMockConstraint(t).RecordCheck(5, true),
					NewMockConstraint(t),
				},
			}).Check,
			5,
		).
		ExpectResult(true)

	NewSubtest(t, "Success: (false && false && true) -> true").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordCheck(5, false),
					NewMockConstraint(t).RecordCheck(5, false),
					NewMockConstraint(t).RecordCheck(5, true),
				},
			}).Check,
			5,
		).
		ExpectResult(true)
}

func TestOrConstraint_String(t *testing.T) {
	NewSubtest(t, "Success: ('First' && 'Second' && 'third')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordString("First"),
					NewMockConstraint(t).RecordString("Second"),
					NewMockConstraint(t).RecordString("third"),
				},
			}).String,
		).
		ExpectResult("(First) or (Second) or (third)")

	NewSubtest(t, "Success: ('First\\nline' && 'Second\\nline' && 'third\\nline')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordString("First\nline"),
					NewMockConstraint(t).RecordString("Second\nline"),
					NewMockConstraint(t).RecordString("third\nline"),
				},
			}).String,
		).
		ExpectResult("(\n\tFirst\n\tline\n) or (\n\tSecond\n\tline\n) or (\n\tthird\n\tline\n)")

	NewSubtest(t, "Success: ('First\\nline' && '' && 'third\\nline')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordString("First\nline"),
					NewMockConstraint(t).RecordString(""),
					NewMockConstraint(t).RecordString("third\nline"),
				},
			}).String,
		).
		ExpectResult("(\n\tFirst\n\tline\n) or (\n\tthird\n\tline\n)")

	NewSubtest(t, "Success: ('First\\nline' && '' && '')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordString("First\nline"),
					NewMockConstraint(t).RecordString(""),
					NewMockConstraint(t).RecordString(""),
				},
			}).String,
		).
		ExpectResult("First\nline")

	NewSubtest(t, "Success: ('f' && '' && '')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordString(""),
					NewMockConstraint(t).RecordString(""),
					NewMockConstraint(t).RecordString(""),
				},
			}).String,
		).
		ExpectResult("")
}

func TestOrConstraint_Details(t *testing.T) {
	NewSubtest(t, "Success: ('First' && 'Second' && 'third')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordDetails(5, "First"),
					NewMockConstraint(t).RecordDetails(5, "Second"),
					NewMockConstraint(t).RecordDetails(5, "third"),
				},
			}).Details,
			5,
		).
		ExpectResult("(First) or (Second) or (third)")

	NewSubtest(t, "Success: ('First\\nline' && 'Second\\nline' && 'third\\nline')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordDetails(5, "First\nline"),
					NewMockConstraint(t).RecordDetails(5, "Second\nline"),
					NewMockConstraint(t).RecordDetails(5, "third\nline"),
				},
			}).Details,
			5,
		).
		ExpectResult("(\n\tFirst\n\tline\n) or (\n\tSecond\n\tline\n) or (\n\tthird\n\tline\n)")

	NewSubtest(t, "Success: ('First\\nline' && '' && 'third\\nline')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordDetails(5, "First\nline"),
					NewMockConstraint(t).RecordDetails(5, ""),
					NewMockConstraint(t).RecordDetails(5, "third\nline"),
				},
			}).Details,
			5,
		).
		ExpectResult("(\n\tFirst\n\tline\n) or (\n\tthird\n\tline\n)")

	NewSubtest(t, "Success: ('First\\nline' && '' && '')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordDetails(5, "First\nline"),
					NewMockConstraint(t).RecordDetails(5, ""),
					NewMockConstraint(t).RecordDetails(5, ""),
				},
			}).Details,
			5,
		).
		ExpectResult("First\nline")

	NewSubtest(t, "Success: ('f' && '' && '')").
		Call(
			(&OrConstraint{
				constraints: []Constraint{
					NewMockConstraint(t).RecordDetails(5, ""),
					NewMockConstraint(t).RecordDetails(5, ""),
					NewMockConstraint(t).RecordDetails(5, ""),
				},
			}).Details,
			5,
		).
		ExpectResult("")
}
