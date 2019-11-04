package unit

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewController(t *testing.T) {
	test := NewMockTestingT(t)

	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockTestingT{}}).
		Call(NewController, test).
		ExpectResult(&Controller{test: test, finishes: []func(){}})

	NewDeclarative(t, "NilTestingT").
		Call(NewController, nil).
		ExpectPanic(NewNotNilError("test"))
}

func TestController_TestingT(t *testing.T) {
	testingT := NewMockTestingT(t)

	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(SamePointerOption{Value: true}).
		Call((&Controller{test: testingT}).TestingT).
		ExpectResult(testingT)
}

func TestController_DeclarativeTest(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call((&Controller{test: NewMockTestingT(t)}).DeclarativeTest, "testName").
		ExpectResult(
			&Declarative{
				test:       NewMockTestingT(t),
				name:       "testName",
				arguments:  []interface{}{},
				comparator: NewEqualComparator(),
			},
		)
}

func TestController_RegisterFinish(t *testing.T) {
	finish := func() {}

	controller := NewController(NewMockTestingT(t).RecordHelper().RecordHelper().RecordHelper().RecordHelper())

	controller.RegisterFinish(finish)

	controller.AssertEqual([]func(){finish}, controller.finishes)
}

func TestController_Finish(t *testing.T) {
	testingT := NewMockTestingT(t).RecordFail()
	finish := func() {
		testingT.Fail()
	}

	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call((&Controller{test: testingT, finishes: []func(){finish}}).Finish).
		ExpectResult()
}

func TestController_And(t *testing.T) {
	constraint1 := NewMockConstraint(t)
	constraint2 := NewMockConstraint(t)

	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call((&Controller{test: NewMockTestingT(t)}).And, constraint1, constraint2).
		ExpectResult(ConstraintAsValue{Value: &AndConstraint{constraints: []Constraint{constraint1, constraint2}}})
}

func TestController_Any(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call((&Controller{test: NewMockTestingT(t)}).Any).
		ExpectResult(ConstraintAsValue{Value: &AnyConstraint{}})
}

func TestController_Callback(t *testing.T) {
	callback := func(actual interface{}) bool {
		return true
	}

	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Callback, callback).
		ExpectResult(ConstraintAsValue{Value: &CallbackConstraint{callback: callback}})

	NewDeclarative(t, "WithPanicResult").
		Call((&Controller{test: NewMockTestingT(t)}).Callback, nil).
		ExpectPanic(NewNotNilError("callback"))
}

func TestController_NotCallback(t *testing.T) {
	callback := func(actual interface{}) bool {
		return true
	}

	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotCallback, callback).
		ExpectResult(ConstraintAsValue{Value: &NotConstraint{constraint: &CallbackConstraint{callback: callback}}})

	NewDeclarative(t, "WithPanicResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotCallback, nil).
		ExpectPanic(NewNotNilError("callback"))
}

func TestController_Contains(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Contains, [1]int{5}, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &ContainsConstraint{
					comparator: NewEqualComparator(SamePointerOption{Value: true}),
					element:    [1]int{5},
				},
			},
		)
}

func TestController_NotContains(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotContains, [1]int{5}, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &ContainsConstraint{
						comparator: NewEqualComparator(SamePointerOption{Value: true}),
						element:    [1]int{5},
					},
				},
			},
		)
}

func TestController_Elements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).Elements,
			NewMockConstraint(t),
			NewMockConstraint(t),
			NewMockConstraint(t),
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &ElementsConstraint{
					constraints: []Constraint{
						NewMockConstraint(t),
						NewMockConstraint(t),
						NewMockConstraint(t),
					},
				},
			},
		)
}

func TestController_NotElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).NotElements,
			NewMockConstraint(t),
			NewMockConstraint(t),
			NewMockConstraint(t),
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &ElementsConstraint{
						constraints: []Constraint{
							NewMockConstraint(t),
							NewMockConstraint(t),
							NewMockConstraint(t),
						},
					},
				},
			},
		)
}

func TestController_ValueElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).ValueElements,
			[]interface{}{
				nil,
				NewMockConstraint(t),
				ConstraintAsValue{Value: NewMockConstraint(t)},
				"data",
			},
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &ElementsConstraint{
					constraints: []Constraint{
						&NilConstraint{},
						NewMockConstraint(t),
						&EqualConstraint{
							expected:   NewMockConstraint(t),
							comparator: NewEqualComparator(SamePointerOption{Value: true}),
						},
						&EqualConstraint{
							expected:   "data",
							comparator: NewEqualComparator(SamePointerOption{Value: true}),
						},
					},
				},
			},
		)
}

func TestController_NotValueElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).NotValueElements,
			[]interface{}{
				nil,
				NewMockConstraint(t),
				ConstraintAsValue{Value: NewMockConstraint(t)},
				"data",
			},
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &ElementsConstraint{
						constraints: []Constraint{
							&NilConstraint{},
							NewMockConstraint(t),
							&EqualConstraint{
								expected:   NewMockConstraint(t),
								comparator: NewEqualComparator(SamePointerOption{Value: true}),
							},
							&EqualConstraint{
								expected:   "data",
								comparator: NewEqualComparator(SamePointerOption{Value: true}),
							},
						},
					},
				},
			},
		)
}

func TestController_Empty(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Empty).
		ExpectResult(ConstraintAsValue{Value: &EmptyConstraint{comparator: NewEqualComparator()}})
}

func TestController_NotEmpty(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotEmpty).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &EmptyConstraint{
						comparator: NewEqualComparator(),
					},
				},
			},
		)
}

func TestController_Equal(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Equal, [1]int{5}, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &EqualConstraint{
					comparator: NewEqualComparator(SamePointerOption{Value: true}),
					expected:   [1]int{5},
				},
			},
		)
}

func TestController_NotEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotEqual, [1]int{5}, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &EqualConstraint{
						comparator: NewEqualComparator(SamePointerOption{Value: true}),
						expected:   [1]int{5},
					},
				},
			},
		)
}

func TestController_False(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).False).
		ExpectResult(ConstraintAsValue{Value: &FalseConstraint{}})
}

func TestController_Greater(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Greater, 55).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: 55}})
}

func TestController_GreaterOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).GreaterOrEqual, 55, float64(32)).
		ExpectResult(
			ConstraintAsValue{
				Value: &OrConstraint{
					constraints: []Constraint{
						&GreaterConstraint{expected: 55},
						&EqualConstraint{
							expected:   55,
							comparator: NewEqualComparator(NumericDeltaOption{Value: 32}),
						},
					},
				},
			},
		)
}

func TestController_Less(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Less, 55).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: 55}})
}

func TestController_LessOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).LessOrEqual, 55, float64(32)).
		ExpectResult(
			ConstraintAsValue{
				Value: &OrConstraint{
					constraints: []Constraint{
						&LessConstraint{expected: 55},
						&EqualConstraint{
							expected:   55,
							comparator: NewEqualComparator(NumericDeltaOption{Value: 32}),
						},
					},
				},
			},
		)
}

func TestController_Kind(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Kind, reflect.Bool, reflect.String).
		ExpectResult(
			ConstraintAsValue{
				Value: &KindConstraint{
					expectedKinds: []reflect.Kind{reflect.Bool, reflect.String},
				},
			},
		)
}

func TestController_NotKind(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotKind, reflect.Bool, reflect.String).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &KindConstraint{
						expectedKinds: []reflect.Kind{reflect.Bool, reflect.String},
					},
				},
			},
		)
}

func TestController_Length(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Length, 5).
		ExpectResult(ConstraintAsValue{Value: &LengthConstraint{length: 5}})
}

func TestController_NotLength(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotLength, 5).
		ExpectResult(ConstraintAsValue{Value: &NotConstraint{constraint: &LengthConstraint{length: 5}}})
}

func TestController_LengthGreater(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).LengthGreater, 5).
		ExpectResult(ConstraintAsValue{Value: &LengthGreaterConstraint{length: 5}})
}

func TestController_LengthGreaterOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).LengthGreaterOrEqual, 5).
		ExpectResult(
			ConstraintAsValue{
				Value: &OrConstraint{
					constraints: []Constraint{
						&LengthGreaterConstraint{length: 5},
						&LengthConstraint{length: 5},
					},
				},
			},
		)
}

func TestController_LengthLess(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).LengthLess, 5).
		ExpectResult(ConstraintAsValue{Value: &LengthLessConstraint{length: 5}})
}

func TestController_LengthLessOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).LengthLessOrEqual, 5).
		ExpectResult(
			ConstraintAsValue{
				Value: &OrConstraint{
					constraints: []Constraint{
						&LengthLessConstraint{length: 5},
						&LengthConstraint{length: 5},
					},
				},
			},
		)
}

func TestController_Nil(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Nil).
		ExpectResult(ConstraintAsValue{Value: &NilConstraint{}})
}

func TestController_NotNil(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotNil).
		ExpectResult(ConstraintAsValue{Value: &NotConstraint{constraint: &NilConstraint{}}})
}

func TestController_Not(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: testing.T{}}).
		Call((&Controller{test: NewMockTestingT(t)}).Not, NewMockConstraint(t)).
		ExpectResult(ConstraintAsValue{Value: &NotConstraint{constraint: NewMockConstraint(t)}})
}

func TestController_Regexp(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Regexp, "^\\d{4}$").
		ExpectResult(ConstraintAsValue{Value: &RegexpConstraint{pattern: regexp.MustCompile("^\\d{4}$")}})
}

func TestController_NotRegexp(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotRegexp, "^\\d{4}$").
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &RegexpConstraint{
						pattern: regexp.MustCompile("^\\d{4}$"),
					},
				},
			},
		)
}

func TestController_Same(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Same, [1]int{5}).
		ExpectResult(
			ConstraintAsValue{
				Value: &SameConstraint{
					comparator: NewEqualComparator(SameTypeOption{Value: true}, SamePointerOption{Value: true}),
					expected:   [1]int{5},
				},
			},
		)
}

func TestController_NotSame(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotSame, [1]int{5}).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &SameConstraint{
						comparator: NewEqualComparator(SameTypeOption{Value: true}, SamePointerOption{Value: true}),
						expected:   [1]int{5},
					},
				},
			},
		)
}

func TestController_True(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).True).
		ExpectResult(ConstraintAsValue{Value: &TrueConstraint{}})
}

func TestController_Type(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).Type, int8(0), new(string)).
		ExpectResult(
			ConstraintAsValue{
				Value: &TypeConstraint{
					expectedTypes: []interface{}{int8(0), new(string)},
				},
			},
		)
}

func TestController_NotType(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&Controller{test: NewMockTestingT(t)}).NotType, int8(0), new(string)).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &TypeConstraint{
						expectedTypes: []interface{}{int8(0), new(string)},
					},
				},
			},
		)
}

func TestController_Value(t *testing.T) {
	NewDeclarative(t, "WithNilExpected").
		Call((&Controller{test: NewMockTestingT(t)}).Value, nil, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &NilConstraint{},
			},
		)

	NewDeclarative(t, "WithConstraintExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).Value,
			NewMockConstraint(t),
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: NewMockConstraint(t),
			},
		)

	NewDeclarative(t, "WithConstraintExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).Value,
			ConstraintAsValue{Value: NewMockConstraint(t)},
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: NewEqualConstraint(
					NewMockConstraint(t),
					NewEqualComparator(SamePointerOption{Value: true}),
				),
			},
		)

	NewDeclarative(t, "WithValueExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).Value,
			"data",
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: NewEqualConstraint(
					"data",
					NewEqualComparator(SamePointerOption{Value: true}),
				),
			},
		)
}

func TestController_NotValue(t *testing.T) {
	NewDeclarative(t, "WithNilExpected").
		Call((&Controller{test: NewMockTestingT(t)}).NotValue, nil, SamePointerOption{Value: true}).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: &NilConstraint{},
				},
			},
		)

	NewDeclarative(t, "WithConstraintExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).NotValue,
			NewMockConstraint(t),
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: NewMockConstraint(t),
				},
			},
		)

	NewDeclarative(t, "WithConstraintExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).NotValue,
			ConstraintAsValue{Value: NewMockConstraint(t)},
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: NewEqualConstraint(
						NewMockConstraint(t),
						NewEqualComparator(SamePointerOption{Value: true}),
					),
				},
			},
		)

	NewDeclarative(t, "WithValueExpected").
		SetCompareOptions(IgnoreUnexportedOption{testing.T{}}).
		Call(
			(&Controller{test: NewMockTestingT(t)}).NotValue,
			"data",
			SamePointerOption{Value: true},
		).
		ExpectResult(
			ConstraintAsValue{
				Value: &NotConstraint{
					constraint: NewEqualConstraint(
						"data",
						NewEqualComparator(SamePointerOption{Value: true}),
					),
				},
			},
		)
}

func TestController_AssertThat(t *testing.T) {
	NewDeclarative(t, "ConstraintWithPositiveResult").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockTestingT{}}).
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper()}).AssertThat,
			5,
			NewMockConstraint(t).RecordCheck(5, true),
			"context",
		).
		ExpectResult(&Controller{test: NewMockTestingT(t)})

	NewDeclarative(t, "ConstraintFailWithoutDetails").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockTestingT{}}).
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordError(NewAssertError(5, "context", "stringResult", "")).
					RecordFail(),
			}).AssertThat,
			5,
			NewMockConstraint(t).
				RecordCheck(5, false).
				RecordDetails(5, "").
				RecordString("stringResult"),
			"context",
		).
		ExpectResult((*Controller)(nil))

	NewDeclarative(t, "ConstraintFailWithDetails").
		SetCompareOptions(IgnoreUnexportedOption{Value: MockTestingT{}}).
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordError(NewAssertError(5, "context", "stringResult", "detailsResult")).
					RecordFail(),
			}).AssertThat,
			5,
			NewMockConstraint(t).
				RecordCheck(5, false).
				RecordDetails(5, "detailsResult").
				RecordString("stringResult"),
			"context",
		).
		ExpectResult((*Controller)(nil))

	NewDeclarative(t, "NilConstraint").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordError(NewNotNilError("constraint")).
					RecordFail(),
			}).AssertThat,
			5,
			nil,
			"context",
		).
		ExpectResult((*Controller)(nil))
}

func TestController_AssertAnd(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertAnd,
			5,
			NewMockConstraint(t).RecordCheck(5, true),
			NewMockConstraint(t).RecordCheck(5, true),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError(5, "'value' variable", "(first) and (second)", "first details")).
					RecordFail(),
			}).AssertAnd,
			5,
			NewMockConstraint(t).RecordCheck(5, false).RecordString("first").RecordDetails(5, "first details"),
			NewMockConstraint(t).RecordString("second").RecordDetails(5, ""),
		).
		ExpectResult()
}

func TestController_AssertCallback(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertCallback,
			5,
			func(interface{}) bool {
				return true
			},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError(5, "'value' variable", "accept callback", "")).
					RecordFail(),
			}).AssertCallback,
			5,
			func(interface{}) bool {
				return false
			},
		).
		ExpectResult()
}

func TestController_AssertNotCallback(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotCallback,
			5,
			func(interface{}) bool {
				return false
			},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError(5, "'value' variable", "not accept callback", "")).
					RecordFail(),
			}).AssertNotCallback,
			5,
			func(interface{}) bool {
				return true
			},
		).
		ExpectResult()
}

func TestController_AssertContains(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertContains,
			5,
			[]int{1, 2, 5},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError([]int{1, 2, 3}, "'list' variable", "contain 5", "")).
					RecordFail(),
			}).AssertContains,
			5,
			[]int{1, 2, 3},
		).
		ExpectResult()
}

func TestController_AssertNotContains(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotContains,
			5,
			[]int{1, 2, 3},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError([]int{1, 2, 5}, "'list' variable", "not contain 5", "")).
					RecordFail(),
			}).AssertNotContains,
			5,
			[]int{1, 2, 5},
		).
		ExpectResult()
}

func TestController_AssertElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertElements,
			[]Constraint{&TrueConstraint{}, &FalseConstraint{}, &GreaterConstraint{expected: 5}},
			true,
			false,
			100,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						&Error{
							message: "Failed assertion for 'actual' variable, " +
								"it must apply all constraints for each elements.\n" +
								"element[1], it must be false, actual value is true.\n" +
								"element[2], it must be greater than 5, actual value is 0.",
						},
					).
					RecordFail(),
			}).AssertElements,
			[]Constraint{&TrueConstraint{}, &FalseConstraint{}, &GreaterConstraint{expected: 5}},
			true,
			true,
			0,
		).
		ExpectResult()
}

func TestController_AssertNotElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotElements,
			[]Constraint{&TrueConstraint{}, &FalseConstraint{}, &GreaterConstraint{expected: 5}},
			false,
			true,
			0,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						&Error{
							message: "Failed assertion for 'actual' variable, it must " +
								"not (apply all constraints for each elements), actual value is [true false 100].",
						},
					).
					RecordFail(),
			}).AssertNotElements,
			[]Constraint{&TrueConstraint{}, &FalseConstraint{}, &GreaterConstraint{expected: 5}},
			true,
			false,
			100,
		).
		ExpectResult()
}

func TestController_AssertValueElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertValueElements,
			[]interface{}{true, &FalseConstraint{}, "data"},
			[]interface{}{true, false, "data"},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						&Error{
							message: "Failed assertion for 'actual' variable, " +
								"it must apply all constraints for each elements.\n" +
								"element[1], it must be false, actual value is true.\n" +
								"element[2], it must be equal to data.\n" +
								"  interface{}(\n- \tstring(\"data\"),\n+ \tint(10),\n  )",
						},
					).
					RecordFail(),
			}).AssertValueElements,
			[]interface{}{true, &FalseConstraint{}, "data"},
			[]interface{}{true, true, 10},
		).
		ExpectResult()
}

func TestController_AssertNotValueElements(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotValueElements,
			[]interface{}{true, &FalseConstraint{}, "data"},
			[]interface{}{false, false, "data"},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						&Error{
							message: "Failed assertion for 'actual' variable, it must " +
								"not (apply all constraints for each elements), actual value is [true false data].",
						},
					).
					RecordFail(),
			}).AssertNotValueElements,
			[]interface{}{true, &FalseConstraint{}, "data"},
			[]interface{}{true, false, "data"},
		).
		ExpectResult()
}

func TestController_AssertEmpty(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertEmpty,
			[]int{},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError([]int{1, 2, 3}, "'value' variable", "be empty", "")).
					RecordFail(),
			}).AssertEmpty,
			[]int{1, 2, 3},
		).
		ExpectResult()
}

func TestController_AssertNotEmpty(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotEmpty,
			[]int{1, 2, 3},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError([]int{}, "'value' variable", "be not empty", "")).
					RecordFail(),
			}).AssertNotEmpty,
			[]int{},
		).
		ExpectResult()
}

func TestController_AssertEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertEqual,
			"data",
			"data",
			SameTypeOption{Value: true},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError("data", "'actual' variable", fmt.Sprintf("be equal to %+v", "data"), ""),
					).
					RecordFail(),
			}).AssertEqual,
			"data",
			"another data",
			SameTypeOption{Value: true},
		).
		ExpectResult()
}

func TestController_AssertNotEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotEqual,
			"data",
			"another data",
			SameTypeOption{Value: true},
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError("data", "'actual' variable", fmt.Sprintf("be not equal to %+v", "data"), ""),
					).
					RecordFail(),
			}).AssertNotEqual,
			"data",
			"data",
			SameTypeOption{Value: true},
		).
		ExpectResult()
}

func TestController_AssertFalse(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertFalse,
			false,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(true, "'value' variable", "be false", ""),
					).
					RecordFail(),
			}).AssertFalse,
			true,
		).
		ExpectResult()
}

func TestController_AssertGreater(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertGreater,
			10,
			100,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(10, "'actual' variable", fmt.Sprintf("be greater than %v", 100), ""),
					).
					RecordFail(),
			}).AssertGreater,
			100,
			10,
		).
		ExpectResult()
}

func TestController_AssertGreaterOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResultByGreater").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertGreaterOrEqual,
			10,
			100,
			float64(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResultByEqual").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertGreaterOrEqual,
			10,
			10,
			float64(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							10,
							"'actual' variable",
							fmt.Sprintf("(be greater than %v) or (be equal to %v)", 100, 100),
							"int(\n- \t100,\n+ \t10,\n  )",
						),
					).
					RecordFail(),
			}).AssertGreaterOrEqual,
			100,
			10,
			float64(0),
		).
		ExpectResult()
}

func TestController_AssertLess(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLess,
			100,
			10,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(100, "'actual' variable", fmt.Sprintf("be less than %v", 10), ""),
					).
					RecordFail(),
			}).AssertLess,
			10,
			100,
		).
		ExpectResult()
}

func TestController_AssertLessOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResultByLess").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLessOrEqual,
			100,
			10,
			float64(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResultByEqual").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLessOrEqual,
			10,
			10,
			float64(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							100,
							"'actual' variable",
							fmt.Sprintf("(be less than %v) or (be equal to %v)", 10, 10),
							"int(\n- \t10,\n+ \t100,\n  )",
						),
					).
					RecordFail(),
			}).AssertLessOrEqual,
			10,
			100,
			float64(0),
		).
		ExpectResult()
}

func TestController_AssertKind(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertKind,
			"data",
			reflect.String,
			reflect.Int,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have one kind of %s, %s", reflect.Array.String(), reflect.Int.String()),
							fmt.Sprintf("Actual kind is %s", reflect.String),
						),
					).
					RecordFail(),
			}).AssertKind,
			"data",
			reflect.Array,
			reflect.Int,
		).
		ExpectResult()
}

func TestController_AssertNotKind(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotKind,
			"data",
			reflect.Array,
			reflect.Int,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have no one kind of %s, %s", reflect.String.String(), reflect.Int.String()),
							fmt.Sprintf("Actual kind is %s", reflect.String),
						),
					).
					RecordFail(),
			}).AssertNotKind,
			"data",
			reflect.String,
			reflect.Int,
		).
		ExpectResult()
}

func TestController_AssertLength(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLength,
			4,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have length %d", 100),
							fmt.Sprintf("Actual length is %d", 4),
						),
					).
					RecordFail(),
			}).AssertLength,
			100,
			"data",
		).
		ExpectResult()
}

func TestController_AssertNotLength(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotLength,
			100,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("not have length %d", 4),
							fmt.Sprintf("Actual length is %d", 4),
						),
					).
					RecordFail(),
			}).AssertNotLength,
			4,
			"data",
		).
		ExpectResult()
}

func TestController_AssertLengthGreater(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthGreater,
			2,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have length greater than %d", 100),
							fmt.Sprintf("Actual length is %d", 4),
						),
					).
					RecordFail(),
			}).AssertLengthGreater,
			100,
			"data",
		).
		ExpectResult()
}

func TestController_AssertLengthGreaterOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResultByGreater").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthGreaterOrEqual,
			2,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResultByEqual").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthGreaterOrEqual,
			4,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("(have length greater than %d) or (have length %d)", 100, 100),
							fmt.Sprintf("(Actual length is %d) or (Actual length is %d)", 4, 4),
						),
					).
					RecordFail(),
			}).AssertLengthGreaterOrEqual,
			100,
			"data",
		).
		ExpectResult()
}

func TestController_AssertLengthLess(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthLess,
			100,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have length less than %d", 2),
							fmt.Sprintf("Actual length is %d", 4),
						),
					).
					RecordFail(),
			}).AssertLengthLess,
			2,
			"data",
		).
		ExpectResult()
}

func TestController_AssertLengthLessOrEqual(t *testing.T) {
	NewDeclarative(t, "WithPositiveResultByLess").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthLessOrEqual,
			100,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResultByEqual").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertLengthLessOrEqual,
			4,
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("(have length less than %d) or (have length %d)", 100, 100),
							fmt.Sprintf("(Actual length is %d) or (Actual length is %d)", 4, 4),
						),
					).
					RecordFail(),
			}).AssertLengthLessOrEqual,
			100,
			"data",
		).
		ExpectResult()
}

func TestController_AssertNil(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNil,
			nil,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError([]int{1, 2, 3}, "'value' variable", "be nil", "")).
					RecordFail(),
			}).AssertNil,
			[]int{1, 2, 3},
		).
		ExpectResult()
}

func TestController_AssertNotNil(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotNil,
			new(int),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError(nil, "'value' variable", "be not nil", "")).
					RecordFail(),
			}).AssertNotNil,
			nil,
		).
		ExpectResult()
}

func TestController_AssertNot(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNot,
			NewMockConstraint(t).RecordCheck("data", false),
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError("data", "'value' variable", "not (string data)", "details data")).
					RecordFail(),
			}).AssertNot,
			NewMockConstraint(t).
				RecordCheck("data", true).
				RecordString("string data").
				RecordDetails("data", "details data"),
			"data",
		).
		ExpectResult()
}

func TestController_AssertOr(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertOr,
			5,
			NewMockConstraint(t).RecordCheck(5, true),
			NewMockConstraint(t),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(NewAssertError(5, "'value' variable", "(first) or (second)", "first details")).
					RecordFail(),
			}).AssertOr,
			5,
			NewMockConstraint(t).RecordCheck(5, false).RecordString("first").RecordDetails(5, "first details"),
			NewMockConstraint(t).RecordCheck(5, false).RecordString("second").RecordDetails(5, ""),
		).
		ExpectResult()
}

func TestController_AssertRegexp(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertRegexp,
			"^\\d{4}$",
			"1234",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"-----",
							"'value' variable",
							fmt.Sprintf("match PCRE pattern '%s'", "^\\d{4}$"),
							"",
						),
					).
					RecordFail(),
			}).AssertRegexp,
			"^\\d{4}$",
			"-----",
		).
		ExpectResult()
}

func TestController_AssertNotRegexp(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotRegexp,
			"^\\d{4}$",
			"-----",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"1234",
							"'value' variable",
							fmt.Sprintf("not match PCRE pattern '%s'", "^\\d{4}$"),
							"",
						),
					).
					RecordFail(),
			}).AssertNotRegexp,
			"^\\d{4}$",
			"1234",
		).
		ExpectResult()
}

func TestController_AssertSame(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertSame,
			"data",
			"data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'actual' variable",
							fmt.Sprintf("be same as %+v", "data"),
							"  string(\n- \t\"data\",\n+ \t\"another data\",\n  )\n",
						),
					).
					RecordFail(),
			}).AssertSame,
			"data",
			"another data",
		).
		ExpectResult()
}

func TestController_AssertNotSame(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotSame,
			"data",
			"another data",
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError("data", "'actual' variable", fmt.Sprintf("not be same as %+v", "data"), ""),
					).
					RecordFail(),
			}).AssertNotSame,
			"data",
			"data",
		).
		ExpectResult()
}

func TestController_AssertTrue(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertTrue,
			true,
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(false, "'value' variable", "be true", ""),
					).
					RecordFail(),
			}).AssertTrue,
			false,
		).
		ExpectResult()
}

func TestController_AssertType(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertType,
			"data",
			"",
			int8(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have one type of %T, %T", []string{}, int8(0)),
							fmt.Sprintf("Actual type is %T", "data"),
						),
					).
					RecordFail(),
			}).AssertType,
			"data",
			[]string{},
			int8(0),
		).
		ExpectResult()
}

func TestController_AssertNotType(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{test: NewMockTestingT(t).RecordHelper().RecordHelper()}).AssertNotType,
			"data",
			[]string{},
			int8(0),
		).
		ExpectResult()

	NewDeclarative(t, "WithNegativeResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'value' variable",
							fmt.Sprintf("have no one type of %T, %T", "", int8(0)),
							fmt.Sprintf("Actual type is %T", "data"),
						),
					).
					RecordFail(),
			}).AssertNotType,
			"data",
			"",
			int8(0),
		).
		ExpectResult()
}

func TestController_AssertValue(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper(),
			}).AssertValue,
			"data",
			"data",
			SamePointerOption{Value: true},
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'actual' variable",
							fmt.Sprintf("be equal to %v", "data"),
							"  string(\n- \t\"data\",\n+ \t\"not data\",\n  )\n",
						),
					).
					RecordFail(),
			}).AssertValue,
			"data",
			"not data",
			SamePointerOption{Value: true},
		).
		ExpectResult()
}

func TestController_AssertNotValue(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper(),
			}).AssertNotValue,
			"data",
			"not data",
			SamePointerOption{Value: true},
		).
		ExpectResult()

	NewDeclarative(t, "WithPositiveResult").
		Call(
			(&Controller{
				test: NewMockTestingT(t).
					RecordHelper().
					RecordHelper().
					RecordError(
						NewAssertError(
							"data",
							"'actual' variable",
							fmt.Sprintf("be not equal to %v", "data"),
							"",
						),
					).
					RecordFail(),
			}).AssertNotValue,
			"data",
			"data",
			SamePointerOption{Value: true},
		).
		ExpectResult()
}
