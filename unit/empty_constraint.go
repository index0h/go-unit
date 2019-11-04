package unit

import (
	"reflect"
)

type EmptyConstraint struct {
	comparator EqualComparer
}

func NewEmptyConstraint() Constraint {
	return &EmptyConstraint{
		comparator: NewEqualComparator(),
	}
}

func (c *EmptyConstraint) Check(value interface{}) bool {
	if value == nil {
		return true
	}

	valueValue := reflect.ValueOf(value)

	switch valueValue.Kind() {
	case reflect.Bool:
		return valueValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return valueValue.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return valueValue.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return valueValue.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return valueValue.Complex() == 0
	case reflect.String:
		return valueValue.Len() == 0
	case reflect.Chan, reflect.Map, reflect.Slice:
		return valueValue.IsNil() || valueValue.Len() == 0
	case reflect.Func, reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		return valueValue.IsNil()
	case reflect.Array:
		for i := 0; i < valueValue.Len(); i++ {
			if !c.Check(valueValue.Index(i).Interface()) {
				return false
			}
		}

		return true
	}

	return c.comparator.Compare(value, reflect.Zero(valueValue.Type()).Interface())
}

func (c *EmptyConstraint) String() string {
	return "be empty"
}

func (c *EmptyConstraint) Details(value interface{}) string {
	return ""
}
