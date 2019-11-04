package unit

import (
	"reflect"
	"runtime/debug"
	"testing"
)

type Declarative struct {
	test       TestingT
	name       string
	callable   interface{}
	arguments  []interface{}
	comparator EqualComparer
}

func NewDeclarative(test TestingT, testName string) *Declarative {
	if test == nil {
		panic(NewNotNilError("test"))
	}

	return &Declarative{
		test:       test,
		name:       testName,
		arguments:  []interface{}{},
		comparator: NewEqualComparator(),
	}
}

func (d *Declarative) SetCompareOptions(options ...interface{}) *Declarative {
	d.test.Helper()

	d.comparator = NewEqualComparator(options...)

	return d
}

func (d *Declarative) Call(callable interface{}, arguments ...interface{}) *Declarative {
	d.test.Helper()

	if callable == nil {
		d.test.Error(NewNotNilError("callable").Error())
		d.test.Fail()

		return d
	}

	callableValue := reflect.ValueOf(callable)

	if callableValue.Kind() != reflect.Func {
		d.test.Error(NewInvalidKindError("callable", callable, reflect.String).Error())
		d.test.Fail()

		return d
	}

	callableType := callableValue.Type()
	argumentsCount := len(arguments)
	callableArgumentsCount := callableType.NumIn()

	if !callableType.IsVariadic() {
		if argumentsCount != callableArgumentsCount {
			d.test.Errorf("Invalid count of arguments, expected: %d, actual: %d", callableArgumentsCount, argumentsCount)
			d.test.Fail()

			return d
		}
	} else {
		if argumentsCount < callableArgumentsCount-1 {
			d.test.Errorf(
				"Count of arguments must be greater or equal than %d, actual: %d",
				callableArgumentsCount-1,
				argumentsCount,
			)
			d.test.Fail()

			return d
		}
	}

	d.callable = callable
	d.arguments = arguments

	return d
}

func (d *Declarative) ExpectResult(expectedResults ...interface{}) {
	d.test.Helper()

	d.test.Run(
		d.name,
		func(test *testing.T) {
			test.Helper()

			d.expectResult(test, expectedResults...)
		},
	)
}

func (d *Declarative) ExpectPanic(expectedPanic interface{}) {
	d.test.Helper()

	d.test.Run(
		d.name,
		func(test *testing.T) {
			test.Helper()

			d.expectPanic(test, expectedPanic)
		},
	)
}

func (d *Declarative) expectResult(test TestingT, expectedResults ...interface{}) {
	test.Helper()

	actualResults, actualPanics, stackTrace := d.call(test)

	if actualPanics != nil {
		test.Errorf("Callable must return result and not call panic:\n%+v\n%s", actualPanics, string(stackTrace))

		return
	}

	if constraint := NewValueElementsConstraint(d.comparator, expectedResults...); !constraint.Check(actualResults) {
		test.Errorf("Failed assertion for expectedResults.\n%s", constraint.Details(actualResults))
	}
}

func (d *Declarative) expectPanic(test TestingT, expectedPanic interface{}) {
	test.Helper()

	actualResults, actualPanic, _ := d.call(test)

	if actualResults != nil {
		test.Errorf("Callable must call panic and not return arguments: %+v", actualResults)

		return
	}

	if constraint := NewValueConstraint(expectedPanic, d.comparator); !constraint.Check(actualPanic) {
		if details := constraint.Details(actualPanic); details != "" {
			test.Errorf("Failed assertion for 'expectedPanic' variable, it must %s.\n%s", constraint.String(), details)
		} else {
			test.Errorf("Failed assertion for 'expectedPanic' variable, it must %s.", constraint.String())
		}
	}
}

func (d *Declarative) call(test TestingT) (results []interface{}, panics interface{}, stackTrace []byte) {
	if d.callable == nil {
		test.Error("Call is not configured")
		test.Fail()

		return nil, nil, nil
	}

	callableValue := reflect.ValueOf(d.callable)
	callableType := callableValue.Type()
	argumentsCount := len(d.arguments)
	callableArgumentsCount := callableType.NumIn()
	callableIsVariadic := callableType.IsVariadic()
	argumentValues := make([]reflect.Value, argumentsCount)

	for i := 0; i < argumentsCount; i++ {
		argumentValue := reflect.ValueOf(d.arguments[i])

		if argumentValue.IsValid() {
			argumentValues[i] = argumentValue

			continue
		}

		if !callableIsVariadic || (callableIsVariadic && i < callableArgumentsCount-1) {
			argumentValues[i] = reflect.New(callableType.In(i)).Elem()
		} else {
			argumentValues[i] = reflect.New(callableType.In(callableArgumentsCount - 1).Elem()).Elem()
		}
	}

	return func() (results []interface{}, panics interface{}, stackTrace []byte) {
		defer func() {
			if panics = recover(); panics != nil {
				stackTrace = debug.Stack()
			}
		}()

		resultsReflect := callableValue.Call(argumentValues)

		results = make([]interface{}, len(resultsReflect))

		for i, resultReflect := range resultsReflect {
			results[i] = resultReflect.Interface()
		}

		return results, panics, stackTrace
	}()
}
