package unit

import (
	"reflect"
	"testing"
)

func TestNewSubtest(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	expected := &Subtest{
		test:       testingT,
		name:       name,
		arguments:  []interface{}{},
		comparator: NewEqualComparator(),
	}

	actual := NewSubtest(testingT, name)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestNewSubtest_WithPanic(t *testing.T) {
	defer func() {
		expected := NewNotNilError("test")

		if diff := NewEqualComparator().Diff(expected, recover()); diff != "" {
			t.Errorf("Return must be %+v.\n%s", expected, diff)
		}
	}()

	name := "name"

	NewSubtest(nil, name)
}

func TestSubtest_SetCompareOptions(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	options := []interface{}{NumericDeltaOption{Value: 5}}
	expected := &Subtest{
		test:       testingT,
		name:       name,
		comparator: NewEqualComparator(options...),
	}

	subtest := &Subtest{
		test:       testingT,
		name:       name,
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	actual := subtest.SetCompareOptions(options...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	callable := func(a1 int, a2 bool) (r1 bool, r2 int) {
		return a2, a1
	}
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Subtest{
		test:      testingT,
		name:      name,
		callable:  callable,
		arguments: arguments,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()

	actual := subtest.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	callable := func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
		return a2, a1
	}
	name := "name"
	arguments := []interface{}{10, false, false, false, false, false, false}
	expected := &Subtest{
		test:      testingT,
		name:      name,
		callable:  callable,
		arguments: arguments,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()

	actual := subtest.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithNilCallable(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Subtest{
		test: testingT,
		name: name,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordError("Variable 'callable' must be not nil")
	testingT.RecordFail()

	actual := subtest.Call(nil, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithCallableNotFunc(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Subtest{
		test: testingT,
		name: name,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordError(NewInvalidKindError("callable", new(int), reflect.String).Error())
	testingT.RecordFail()

	actual := subtest.Call(new(int), arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithCallableNotVariadicAndInvalidArgumentsCount(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	callable := func(a1 int, a2 bool) {}
	arguments := []interface{}{10}
	expected := &Subtest{
		test: testingT,
		name: name,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Invalid count of arguments, expected: %d, actual: %d", 2, 1)
	testingT.RecordFail()

	actual := subtest.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_Call_WithCallableVariadicAndInvalidArgumentsCount(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	callable := func(a1 int, a2 ...bool) {}
	arguments := []interface{}{}
	expected := &Subtest{
		test: testingT,
		name: name,
	}

	subtest := &Subtest{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Count of arguments must be greater or equal than %d, actual: %d", 1, 0)
	testingT.RecordFail()

	actual := subtest.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestSubtest_expectResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{true, 5}

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 bool) (r1 bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithVariadicArguments(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{[]bool{true, false, false}, 5}

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true, false, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithNilAndNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{(*int)(nil)}

	subtest := &Subtest{
		name: name,
		callable: func(a1 *int) (r1 *int) {
			return a1
		},
		arguments:  []interface{}{nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithNilAndVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{(*int)(nil), (*int)(nil), (*int)(nil)}

	subtest := &Subtest{
		name: name,
		callable: func(a1 ...*int) (r1 *int, r2 *int, r3 *int) {
			return a1[0], a1[1], a1[2]
		},
		arguments:  []interface{}{nil, nil, nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithConstraints(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	constraint := NewMockConstraint(&testing.T{})
	results := []interface{}{
		5,
		ConstraintAsValue{Value: constraint},
	}

	subtest := &Subtest{
		name: name,
		callable: func() (interface{}, interface{}) {
			return 5, constraint
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(SamePointerOption{}),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithConstraintAsValue(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		ConstraintAsValue{Value: NewMockConstraint(&testing.T{})},
	}

	subtest := &Subtest{
		name: name,
		callable: func() interface{} {
			return NewMockConstraint(&testing.T{})
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(IgnoreUnexportedOption{Value: testing.T{}}),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithCallablePanic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{[]bool{true, false}, 5}

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
			panic("data")
		},
		arguments:  []interface{}{5, true, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf(
		"Callable must return result and not call panic:\n%+v\n%s",
		"data",
		NewNotConstraint(NewNilConstraint()),
	)

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithResultsCountMismatch(t *testing.T) {
	defer func() {
		expected := NewInvalidLengthError("list", 1, 2)

		if diff := NewEqualComparator().Diff(expected, recover()); diff != "" {
			t.Errorf("Panic must be %+v.\n%s", expected, diff)
		}
	}()

	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{true}

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 bool) (r1 bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true},
		comparator: NewMockEqualComparer(t),
	}

	testingT.RecordHelper()

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithConstraintCheckError(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		NewMockConstraint(t).
			RecordCheck(10, false).
			RecordCheck(10, false).
			RecordString("stringResult").
			RecordDetails(10, "detailsResult"),
	}

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			return argument * 2
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf(
		"Failed assertion for expectedResults.\n%s",
		"element[0], it must stringResult.\ndetailsResult",
	)

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithConstraintCheckErrorAndWithoutDetails(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		NewMockConstraint(t).
			RecordCheck(10, false).
			RecordCheck(10, false).
			RecordString("stringResult").
			RecordDetails(10, ""),
	}

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			return argument * 2
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Failed assertion for expectedResults.\n%s", "")

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectResult_WithInvalidResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{5}

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			return argument * 2
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf(
		"Failed assertion for expectedResults.\n%s",
		"element[0], it must be equal to 5.\n  int(\n- \t5,\n+ \t10,\n  )",
	)

	subtest.expectResult(testingT, results...)
}

func TestSubtest_expectPanic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := 5

	subtest := &Subtest{
		name: name,
		callable: func(a1 int) (r1 bool) {
			panic(a1)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithVariadicArguments(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := []bool{true, false, false}

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
			panic(a2)
		},
		arguments:  []interface{}{5, true, false, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithNilAndNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := (*int)(nil)

	subtest := &Subtest{
		name: name,
		callable: func(a1 *int) (r1 *int) {
			panic(a1)
		},
		arguments:  []interface{}{nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithNilAndVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := (*int)(nil)

	subtest := &Subtest{
		name: name,
		callable: func(a1 ...*int) (r1 *int, r2 *int, r3 *int) {
			panic(a1[2])
		},
		arguments:  []interface{}{nil, nil, nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithConstraints(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).RecordCheck(true, true)

	subtest := &Subtest{
		name: name,
		callable: func() (interface{}, interface{}, interface{}) {
			panic(true)
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithConstraintAsValue(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := ConstraintAsValue{Value: &MockConstraint{}}

	subtest := &Subtest{
		name: name,
		callable: func() (interface{}, interface{}, interface{}) {
			panic(&MockConstraint{})
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}),
	}

	testingT.RecordHelper()

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithCallableResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := "data"

	subtest := &Subtest{
		name: name,
		callable: func(a1 int, a2 ...bool) int {
			return a1
		},
		arguments:  []interface{}{5, true, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Callable must call panic and not return arguments: %+v", []interface{}{5})

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithConstraintCheckError(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).
		RecordCheck(10, false).
		RecordString("stringResult").
		RecordDetails(10, "detailsResult")

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			panic(argument * 2)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf(
		"Failed assertion for 'expectedPanic' variable, it must %s.\n%s",
		"stringResult",
		"detailsResult",
	)

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithConstraintCheckErrorAndWithoutDetails(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).
		RecordCheck(10, false).
		RecordDetails(10, "").
		RecordString("stringResult")

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			panic(argument * 2)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Failed assertion for 'expectedPanic' variable, it must %s.", "stringResult")

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_expectPanic_WithInvalidResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := 5

	subtest := &Subtest{
		name: name,
		callable: func(argument int) int {
			panic(argument * 2)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf(
		"Failed assertion for 'expectedPanic' variable, it must %s.\n%s",
		"be equal to 5",
		"  int(\n- \t5,\n+ \t10,\n  )\n",
	)

	subtest.expectPanic(testingT, panics)
}

func TestSubtest_call_WithNotConfiguredCall(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"

	subtest := &Subtest{
		name: name,
	}

	testingT.RecordError("Call is not configured")
	testingT.RecordFail()

	results, panics, stackTrace := subtest.call(testingT)

	if diff := NewEqualComparator().Diff([]interface{}(nil), results); diff != "" {
		t.Errorf("Results must be nil.\n%s", diff)
	}

	if diff := NewEqualComparator().Diff(interface{}(nil), panics); diff != "" {
		t.Errorf("Panics must be nil.\n%s", diff)
	}

	if diff := NewEqualComparator().Diff([]byte(nil), stackTrace); diff != "" {
		t.Errorf("StackTrace must be nil.\n%s", diff)
	}
}
