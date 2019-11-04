package unit

import "testing"

func TestNewTypeConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(NewTypeConstraint, 5, "data", new(int)).
		ExpectResult(ConstraintAsValue{Value: &TypeConstraint{expectedTypes: []interface{}{5, "data", new(int)}}})

	NewDeclarative(t, "WithNegativeResultByEmptyTypes").
		Call(NewTypeConstraint).
		ExpectPanic(&Error{message: "Variable 'expectedTypes' must be not empty"})

	NewDeclarative(t, "WithNegativeResultByInvalidType").
		Call(NewTypeConstraint, nil).
		ExpectPanic(&Error{message: "Invalid type of expectedTypes[0]"})
}

func TestTypeConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&TypeConstraint{expectedTypes: []interface{}{5, "data", new(int)}}).Check, "data").
		ExpectResult(true)

	NewDeclarative(t, "WithNegativeResult").
		Call((&TypeConstraint{expectedTypes: []interface{}{5, "data", new(int)}}).Check, []int{}).
		ExpectResult(false)
}

func TestTypeConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&TypeConstraint{expectedTypes: []interface{}{5, "data", new(int)}}).String).
		ExpectResult("have one type of int, string, *int")
}

func TestTypeConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&TypeConstraint{expectedTypes: []interface{}{}}).Details, new(int)).
		ExpectResult("Actual type is *int")
}
