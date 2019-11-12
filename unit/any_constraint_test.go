package unit

import "testing"

func TestNewAnyConstraint(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewAnyConstraint).
		ExpectResult(ConstraintAsValue{Value: &AnyConstraint{}})
}

func TestAnyConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&AnyConstraint{}).Check, 5).
		ExpectResult(true)
}

func TestAnyConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&AnyConstraint{}).String).
		ExpectResult("be anything")
}

func TestAnyConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&AnyConstraint{}).Details, 5).
		ExpectResult("")
}
