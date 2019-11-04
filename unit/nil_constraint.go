package unit

import "reflect"

type NilConstraint struct {
}

func NewNilConstraint() Constraint {
	return &NilConstraint{}
}

func (c *NilConstraint) Check(value interface{}) bool {
	if value == nil {
		return true
	}

	valueValue := reflect.ValueOf(value)

	switch valueValue.Kind() {
	case
		reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return valueValue.IsNil()
	}

	panic(NewInvalidNilComparisonTypeError("value", value))
}

func (c *NilConstraint) String() string {
	return "be nil"
}

func (c *NilConstraint) Details(value interface{}) string {
	return ""
}
