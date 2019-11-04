package unit

import (
	"fmt"
	"reflect"
	"strings"
)

type ContainsConstraint struct {
	element    interface{}
	comparator EqualComparer
}

func NewContainsConstraint(element interface{}, options ...interface{}) Constraint {
	return &ContainsConstraint{
		comparator: NewEqualComparator(options...),
		element:    element,
	}
}

func (c *ContainsConstraint) Check(list interface{}) bool {
	listValue := reflect.ValueOf(list)

	switch listValue.Kind() {
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		length := listValue.Len()

		for i := 0; i < length; i++ {
			if c.comparator.Compare(c.element, listValue.Index(i).Interface()) {
				return true
			}
		}

		return false
	case reflect.Map:
		keys := listValue.MapKeys()

		for i := 0; i < len(keys); i++ {
			if c.comparator.Compare(c.element, listValue.MapIndex(keys[i]).Interface()) {
				return true
			}
		}

		return false
	case reflect.String:
		if reflect.ValueOf(c.element).Kind() != reflect.String {
			err := NewInvalidKindError("element", c.element, reflect.String)

			panic(err)
		}

		return strings.Contains(list.(string), c.element.(string))
	default:
		err := NewInvalidKindError("list", list, reflect.Array, reflect.Slice, reflect.Map, reflect.String)

		panic(err)
	}
}

func (c *ContainsConstraint) String() string {
	return fmt.Sprintf("contain %+v", c.element)
}

func (c *ContainsConstraint) Details(list interface{}) string {
	return ""
}
