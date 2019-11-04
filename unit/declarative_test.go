package unit

import (
	"reflect"
	"testing"
)

func TestNewDeclarative(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	expected := &Declarative{
		test:       testingT,
		name:       name,
		arguments:  []interface{}{},
		comparator: NewEqualComparator(),
	}

	actual := NewDeclarative(testingT, name)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestNewDeclarative_WithPanic(t *testing.T) {
	defer func() {
		expected := NewNotNilError("test")

		if diff := NewEqualComparator().Diff(expected, recover()); diff != "" {
			t.Errorf("Return must be %+v.\n%s", expected, diff)
		}
	}()

	name := "name"

	NewDeclarative(nil, name)
}

func TestDeclarative_SetCompareOptions(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	options := []interface{}{NumericDeltaOption{Value: 5}}
	expected := &Declarative{
		test:       testingT,
		name:       name,
		comparator: NewEqualComparator(options...),
	}

	declarative := &Declarative{
		test:       testingT,
		name:       name,
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	actual := declarative.SetCompareOptions(options...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	callable := func(a1 int, a2 bool) (r1 bool, r2 int) {
		return a2, a1
	}
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Declarative{
		test:      testingT,
		name:      name,
		callable:  callable,
		arguments: arguments,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()

	actual := declarative.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	callable := func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
		return a2, a1
	}
	name := "name"
	arguments := []interface{}{10, false, false, false, false, false, false}
	expected := &Declarative{
		test:      testingT,
		name:      name,
		callable:  callable,
		arguments: arguments,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()

	actual := declarative.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithNilCallable(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Declarative{
		test: testingT,
		name: name,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordError("Variable 'callable' must be not nil")
	testingT.RecordFail()

	actual := declarative.Call(nil, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithCallableNotFunc(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	arguments := []interface{}{10, false}
	expected := &Declarative{
		test: testingT,
		name: name,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordError(NewInvalidKindError("callable", new(int), reflect.String).Error())
	testingT.RecordFail()

	actual := declarative.Call(new(int), arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithCallableNotVariadicAndInvalidArgumentsCount(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	callable := func(a1 int, a2 bool) {}
	arguments := []interface{}{10}
	expected := &Declarative{
		test: testingT,
		name: name,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Invalid count of arguments, expected: %d, actual: %d", 2, 1)
	testingT.RecordFail()

	actual := declarative.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_Call_WithCallableVariadicAndInvalidArgumentsCount(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	callable := func(a1 int, a2 ...bool) {}
	arguments := []interface{}{}
	expected := &Declarative{
		test: testingT,
		name: name,
	}

	declarative := &Declarative{
		test: testingT,
		name: name,
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Count of arguments must be greater or equal than %d, actual: %d", 1, 0)
	testingT.RecordFail()

	actual := declarative.Call(callable, arguments...)

	if diff := NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}).Diff(expected, actual); diff != "" {
		t.Errorf("Return must be %+v.\n%s", expected, diff)
	}
}

func TestDeclarative_expectResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{true, 5}

	declarative := &Declarative{
		name: name,
		callable: func(a1 int, a2 bool) (r1 bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithVariadicArguments(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{[]bool{true, false, false}, 5}

	declarative := &Declarative{
		name: name,
		callable: func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true, false, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithNilAndNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{(*int)(nil)}

	declarative := &Declarative{
		name: name,
		callable: func(a1 *int) (r1 *int) {
			return a1
		},
		arguments:  []interface{}{nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithNilAndVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{(*int)(nil), (*int)(nil), (*int)(nil)}

	declarative := &Declarative{
		name: name,
		callable: func(a1 ...*int) (r1 *int, r2 *int, r3 *int) {
			return a1[0], a1[1], a1[2]
		},
		arguments:  []interface{}{nil, nil, nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithConstraints(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	constraint := NewMockConstraint(&testing.T{})
	results := []interface{}{
		5,
		ConstraintAsValue{Value: constraint},
	}

	declarative := &Declarative{
		name: name,
		callable: func() (interface{}, interface{}) {
			return 5, constraint
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(SamePointerOption{}),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithConstraintAsValue(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		ConstraintAsValue{Value: NewMockConstraint(&testing.T{})},
	}

	declarative := &Declarative{
		name: name,
		callable: func() interface{} {
			return NewMockConstraint(&testing.T{})
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(IgnoreUnexportedOption{Value: testing.T{}}),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithCallablePanic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{[]bool{true, false}, 5}

	declarative := &Declarative{
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

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithResultsCountMismatch(t *testing.T) {
	defer func() {
		expected := NewInvalidLengthError("list", 1, 2)

		if diff := NewEqualComparator().Diff(expected, recover()); diff != "" {
			t.Errorf("Panic must be %+v.\n%s", expected, diff)
		}
	}()

	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{true}

	declarative := &Declarative{
		name: name,
		callable: func(a1 int, a2 bool) (r1 bool, r2 int) {
			return a2, a1
		},
		arguments:  []interface{}{5, true},
		comparator: NewMockEqualComparer(t),
	}

	testingT.RecordHelper()

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithConstraintCheckError(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		NewMockConstraint(t).
			RecordCheck(10, false).
			RecordCheck(10, false).
			RecordString("stringResult").
			RecordDetails(10, "detailsResult"),
	}

	declarative := &Declarative{
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

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithConstraintCheckErrorAndWithoutDetails(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{
		NewMockConstraint(t).
			RecordCheck(10, false).
			RecordCheck(10, false).
			RecordString("stringResult").
			RecordDetails(10, ""),
	}

	declarative := &Declarative{
		name: name,
		callable: func(argument int) int {
			return argument * 2
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Failed assertion for expectedResults.\n%s", "")

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectResult_WithInvalidResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	results := []interface{}{5}

	declarative := &Declarative{
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

	declarative.expectResult(testingT, results...)
}

func TestDeclarative_expectPanic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := 5

	declarative := &Declarative{
		name: name,
		callable: func(a1 int) (r1 bool) {
			panic(a1)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithVariadicArguments(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := []bool{true, false, false}

	declarative := &Declarative{
		name: name,
		callable: func(a1 int, a2 ...bool) (r1 []bool, r2 int) {
			panic(a2)
		},
		arguments:  []interface{}{5, true, false, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithNilAndNotVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := (*int)(nil)

	declarative := &Declarative{
		name: name,
		callable: func(a1 *int) (r1 *int) {
			panic(a1)
		},
		arguments:  []interface{}{nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithNilAndVariadic(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := (*int)(nil)

	declarative := &Declarative{
		name: name,
		callable: func(a1 ...*int) (r1 *int, r2 *int, r3 *int) {
			panic(a1[2])
		},
		arguments:  []interface{}{nil, nil, nil},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithConstraints(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).RecordCheck(true, true)

	declarative := &Declarative{
		name: name,
		callable: func() (interface{}, interface{}, interface{}) {
			panic(true)
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithConstraintAsValue(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := ConstraintAsValue{Value: &MockConstraint{}}

	declarative := &Declarative{
		name: name,
		callable: func() (interface{}, interface{}, interface{}) {
			panic(&MockConstraint{})
		},
		arguments:  []interface{}{},
		comparator: NewEqualComparator(IgnoreUnexportedOption{Value: MockTestingT{}}),
	}

	testingT.RecordHelper()

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithCallableResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := "data"

	declarative := &Declarative{
		name: name,
		callable: func(a1 int, a2 ...bool) int {
			return a1
		},
		arguments:  []interface{}{5, true, false},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Callable must call panic and not return arguments: %+v", []interface{}{5})

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithConstraintCheckError(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).
		RecordCheck(10, false).
		RecordString("stringResult").
		RecordDetails(10, "detailsResult")

	declarative := &Declarative{
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

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithConstraintCheckErrorAndWithoutDetails(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := NewMockConstraint(t).
		RecordCheck(10, false).
		RecordDetails(10, "").
		RecordString("stringResult")

	declarative := &Declarative{
		name: name,
		callable: func(argument int) int {
			panic(argument * 2)
		},
		arguments:  []interface{}{5},
		comparator: NewEqualComparator(),
	}

	testingT.RecordHelper()
	testingT.RecordErrorf("Failed assertion for 'expectedPanic' variable, it must %s.", "stringResult")

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_expectPanic_WithInvalidResult(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"
	panics := 5

	declarative := &Declarative{
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

	declarative.expectPanic(testingT, panics)
}

func TestDeclarative_call_WithNotConfiguredCall(t *testing.T) {
	testingT := NewMockTestingT(t)
	name := "name"

	declarative := &Declarative{
		name: name,
	}

	testingT.RecordError("Call is not configured")
	testingT.RecordFail()

	results, panics, stackTrace := declarative.call(testingT)

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
