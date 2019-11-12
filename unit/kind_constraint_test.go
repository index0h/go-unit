package unit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewKindConstraint(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewKindConstraint, reflect.Int, reflect.String, reflect.Ptr).
		ExpectResult(
			ConstraintAsValue{
				Value: &KindConstraint{expectedKinds: []reflect.Kind{reflect.Int, reflect.String, reflect.Ptr}},
			},
		)

	NewSubtest(t, "WithNegativeResultByEmptyKinds").
		Call(NewKindConstraint).
		ExpectPanic(&Error{message: "Variable 'expectedKinds' must be not empty"})
}

func TestKindConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&KindConstraint{expectedKinds: []reflect.Kind{reflect.Int, reflect.String, reflect.Ptr}}).Check, "data").
		ExpectResult(true)

	NewSubtest(t, "WithNegativeResult").
		Call((&KindConstraint{expectedKinds: []reflect.Kind{reflect.Int, reflect.String, reflect.Ptr}}).Check, []int{}).
		ExpectResult(false)
}

func TestKindConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&KindConstraint{expectedKinds: []reflect.Kind{reflect.Int, reflect.String, reflect.Ptr}}).String).
		ExpectResult("have one kind of int, string, ptr")
}

func TestKindConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(
			(&KindConstraint{expectedKinds: []reflect.Kind{reflect.Int, reflect.String, reflect.Ptr}}).Details,
			new(int),
		).
		ExpectResult(fmt.Sprintf("Actual kind is %s", reflect.ValueOf(new(int)).Kind().String()))
}
