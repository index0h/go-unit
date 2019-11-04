package unit

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewNotConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(NewNotConstraint, NewMockConstraint(t)).
		ExpectResult(ConstraintAsValue{Value: &NotConstraint{constraint: NewMockConstraint(t)}})

	NewDeclarative(t, "WithNegativeResultByNiConstraint").
		Call(NewNotConstraint, nil).
		ExpectPanic(NewNotNilError("constraint"))
}

func TestNotConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordCheck("data", false)}).Check, "data").
		ExpectResult(true)

	NewDeclarative(t, "WithNegativeResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordCheck("data", true)}).Check, "data").
		ExpectResult(false)
}

func TestNotConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithAndConstraintAndOneLineResult").
		Call(
			(&NotConstraint{
				constraint: &AndConstraint{
					constraints: []Constraint{
						NewMockConstraint(t).RecordString("first data"),
						NewMockConstraint(t).RecordString("second data"),
					},
				},
			}).String,
		).
		ExpectResult("be not ((first data) and (second data))")

	NewDeclarative(t, "WithAndConstraintAndMultilineResult").
		Call(
			(&NotConstraint{
				constraint: &AndConstraint{
					constraints: []Constraint{
						NewMockConstraint(t).RecordString("first\ndata"),
						NewMockConstraint(t).RecordString("second data"),
					},
				},
			}).String,
		).
		ExpectResult("be not (\n\t(\n\t\tfirst\n\t\tdata\n\t) and (second data)\n)")

	NewDeclarative(t, "WithAnythingConstraint").
		Call((&NotConstraint{constraint: &AnyConstraint{}}).String).
		ExpectResult("be nothing")

	NewDeclarative(t, "WithCallbackConstraint").
		Call(
			(&NotConstraint{
				constraint: &CallbackConstraint{
					callback: func(interface{}) bool {
						return true
					},
				},
			}).String,
		).
		ExpectResult("not accept callback")

	NewDeclarative(t, "WithContainsConstraint").
		Call(
			(&NotConstraint{
				constraint: &ContainsConstraint{
					element: [1]int{1},
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("not contain %+v", [1]int{1}))

	NewDeclarative(t, "WithEmptyConstraint").
		Call((&NotConstraint{constraint: &EmptyConstraint{}}).String).
		ExpectResult("be not empty")

	NewDeclarative(t, "WithEqualConstraint").
		Call(
			(&NotConstraint{
				constraint: &EqualConstraint{
					expected: [1]int{1},
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("be not equal to %+v", [1]int{1}))

	NewDeclarative(t, "WithFalseConstraint").
		Call((&NotConstraint{constraint: &FalseConstraint{}}).String).
		ExpectResult("be true")

	NewDeclarative(t, "WithGreaterConstraint").
		Call(
			(&NotConstraint{
				constraint: &GreaterConstraint{
					expected: 10,
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("be less than or equal to %v", 10))

	NewDeclarative(t, "WithGreaterConstraint").
		Call(
			(&NotConstraint{
				constraint: &KindConstraint{
					expectedKinds: []reflect.Kind{reflect.String, reflect.Bool},
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("have no one kind of %s, %s", reflect.String, reflect.Bool))

	NewDeclarative(t, "WithLengthConstraint").
		Call(
			(&NotConstraint{
				constraint: &LengthConstraint{
					length: 10,
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("not have length %v", 10))

	NewDeclarative(t, "WithLengthGreaterConstraint").
		Call(
			(&NotConstraint{
				constraint: &LengthGreaterConstraint{
					length: 10,
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("have length less than or equal to %v", 10))

	NewDeclarative(t, "WithLengthLessConstraint").
		Call(
			(&NotConstraint{
				constraint: &LengthLessConstraint{
					length: 10,
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("have length greater than or equal to %v", 10))

	NewDeclarative(t, "WithLessConstraint").
		Call(
			(&NotConstraint{
				constraint: &LessConstraint{
					expected: 10,
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("be greater than or equal to %v", 10))

	NewDeclarative(t, "WithNilConstraint").
		Call((&NotConstraint{constraint: &NilConstraint{}}).String).
		ExpectResult("be not nil")

	NewDeclarative(t, "WithNotConstraint").
		Call(
			(&NotConstraint{
				constraint: &NotConstraint{
					constraint: NewMockConstraint(t).RecordString("data"),
				},
			}).String,
		).
		ExpectResult("data")

	NewDeclarative(t, "WithAndConstraintAndOneLineResult").
		Call(
			(&NotConstraint{
				constraint: &AndConstraint{
					constraints: []Constraint{
						NewMockConstraint(t).RecordString("first data"),
						NewMockConstraint(t).RecordString("second data"),
					},
				},
			}).String,
		).
		ExpectResult("be not ((first data) and (second data))")

	NewDeclarative(t, "WithOrConstraintAndMultilineResult").
		Call(
			(&NotConstraint{
				constraint: &OrConstraint{
					constraints: []Constraint{
						NewMockConstraint(t).RecordString("first\ndata"),
						NewMockConstraint(t).RecordString("second data"),
					},
				},
			}).String,
		).
		ExpectResult("be not (\n\t(\n\t\tfirst\n\t\tdata\n\t) or (second data)\n)")

	NewDeclarative(t, "WithRegexpConstraint").
		Call(
			(&NotConstraint{
				constraint: &RegexpConstraint{
					pattern: regexp.MustCompile("^\\d{3}$"),
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("not match PCRE pattern '%s'", regexp.MustCompile("^\\d{3}$")))

	NewDeclarative(t, "WithSameConstraint").
		Call(
			(&NotConstraint{
				constraint: &SameConstraint{
					expected: [1]int{5},
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("not be same as %+v", [1]int{5}))

	NewDeclarative(t, "WithSameConstraint").
		Call((&NotConstraint{constraint: &TrueConstraint{}}).String).
		ExpectResult("be false")

	NewDeclarative(t, "WithTypeConstraint").
		Call(
			(&NotConstraint{
				constraint: &TypeConstraint{
					expectedTypes: []interface{}{int32(5), [1]string{"data"}},
				},
			}).String,
		).
		ExpectResult(fmt.Sprintf("have no one type of %T, %T", int32(5), [1]string{"data"}))

	NewDeclarative(t, "WithUnknownConstraintAndOneLineResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordString("data")}).String).
		ExpectResult("not (data)")

	NewDeclarative(t, "WithUnknownConstraintAndMultiLineResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordString("first\ndata")}).String).
		ExpectResult("not (\n\tfirst\n\tdata\n)")
}

func TestNotConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithEmptyResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordDetails(10, "")}).Details, 10).
		ExpectResult("")

	NewDeclarative(t, "WithOneLineResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordDetails(10, "data")}).Details, 10).
		ExpectResult("data")

	NewDeclarative(t, "WithMultiLineResult").
		Call((&NotConstraint{constraint: NewMockConstraint(t).RecordDetails(10, "first\ndata")}).Details, 10).
		ExpectResult("first\ndata")
}
