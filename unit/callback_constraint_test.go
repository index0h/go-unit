package unit

import (
	"testing"
)

func TestNewCallbackConstraint(t *testing.T) {
	callback := func(actual interface{}) bool {
		return true
	}

	NewSubtest(t, "Success").
		Call(NewCallbackConstraint, callback).
		ExpectResult(ConstraintAsValue{Value: &CallbackConstraint{callback: callback}})

	NewSubtest(t, "Panic").
		Call(NewCallbackConstraint, nil).
		ExpectPanic(NewNotNilError("callback"))
}

func TestCallbackConstraint_Check(t *testing.T) {
	callback := func(expected interface{}) func(actual interface{}) bool {
		return func(actual interface{}) bool {
			return NewEqualComparator().Compare(expected, actual)
		}
	}

	NewSubtest(t, "Success: -> true").
		Call((&CallbackConstraint{callback: callback(5)}).Check, 5).
		ExpectResult(true)

	NewSubtest(t, "Success: -> true").
		Call((&CallbackConstraint{callback: callback(5)}).Check, "data").
		ExpectResult(false)
}

func TestCallbackConstraint_String(t *testing.T) {
	callback := func(actual interface{}) bool {
		return true
	}

	NewSubtest(t, "Success: -> true").
		Call((&CallbackConstraint{callback: callback}).String).
		ExpectResult("accept callback")
}

func TestCallbackConstraint_Details(t *testing.T) {
	callback := func(actual interface{}) bool {
		return true
	}

	NewSubtest(t, "Success: -> true").
		Call((&CallbackConstraint{callback: callback}).Details, 5).
		ExpectResult("")
}
