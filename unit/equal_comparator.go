package unit

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type NumericDeltaOption struct {
	Value float64
}

type SameTypeOption struct {
	Value bool
}

type SamePointerOption struct {
	Value bool
}

type UseEqualMethodOption struct {
	Value bool
}

type IgnoreUnexportedOption struct {
	Value interface{}
}

type IgnoreFieldsOption struct {
	Type   interface{}
	Fields []string
}

type EqualComparator struct {
	numericDelta     float64
	sameType         bool
	samePointer      bool
	useEqualMethod   bool
	ignoreUnexported []interface{}
	ignoresFields    []IgnoreFieldsOption
}

func NewEqualComparator(options ...interface{}) *EqualComparator {
	result := &EqualComparator{
		ignoreUnexported: []interface{}{},
		ignoresFields:    []IgnoreFieldsOption{},
	}

	for i, option := range options {
		switch option.(type) {
		case NumericDeltaOption:
			typedOption := option.(NumericDeltaOption)

			if typedOption.Value < 0 {
				err := NewErrorf(
					"Variable 'options[%d].Value' must be greater or equal to 0, actual: %f",
					i,
					typedOption.Value,
				)

				panic(err)
			}

			result.numericDelta = option.(NumericDeltaOption).Value
		case SameTypeOption:
			result.sameType = option.(SameTypeOption).Value
		case SamePointerOption:
			result.samePointer = option.(SamePointerOption).Value
		case UseEqualMethodOption:
			result.useEqualMethod = option.(UseEqualMethodOption).Value
		case IgnoreUnexportedOption:
			typedOption := option.(IgnoreUnexportedOption)

			if reflect.ValueOf(typedOption.Value).Kind() != reflect.Struct {
				err := NewInvalidKindError(
					fmt.Sprintf("options[%d].Value", i),
					typedOption.Value,
					reflect.Struct,
				)

				panic(err)
			}

			result.ignoreUnexported = append(result.ignoreUnexported, typedOption.Value)
		case IgnoreFieldsOption:
			typedOption := option.(IgnoreFieldsOption)

			typedOptionTypeValue := reflect.ValueOf(typedOption.Type)

			if typedOptionTypeValue.Kind() != reflect.Struct {
				err := NewInvalidKindError(fmt.Sprintf("options[%d].Type", i), typedOption.Type, reflect.Struct)

				panic(err)
			}

			typedOptionTypeValueType := typedOptionTypeValue.Type()

			for j, fieldName := range typedOption.Fields {
				if _, ok := typedOptionTypeValueType.FieldByName(fieldName); !ok {
					err := NewErrorf(
						"Variable 'options[%d].Fields[%d]' contains unknown field name: '%s'",
						i,
						j,
						fieldName,
					)

					panic(err)
				}
			}

			result.ignoresFields = append(result.ignoresFields, typedOption)

		default:
			err := NewErrorf("Variable 'options[%d]' has unknown type: %T", i, option)

			panic(err)
		}
	}

	return result
}

func (c *EqualComparator) Compare(x interface{}, y interface{}) bool {
	return cmp.Equal(x, y, c.buildCmpOptions(x, y)...)
}

func (c *EqualComparator) Diff(x interface{}, y interface{}) string {
	return strings.Replace(cmp.Diff(x, y, c.buildCmpOptions(x, y)...), "\u00a0", " ", -1)
}

func (c *EqualComparator) compareAny(x interface{}, y interface{}) bool {
	if c.sameType && reflect.TypeOf(x) != reflect.TypeOf(y) {
		return false
	}

	if c.useEqualMethod {
		if xEqualer, ok := x.(Equaler); ok && x != nil {
			return xEqualer.Equal(y)
		}

		if yEqualer, ok := y.(Equaler); ok && y != nil {
			return yEqualer.Equal(x)
		}
	}

	xValue := reflect.ValueOf(x)
	yValue := reflect.ValueOf(y)

	switch xValue.Kind() {
	case reflect.Bool:
		if yValue.Kind() != reflect.Bool {
			return false
		}

		return xValue.Bool() == yValue.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch yValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return math.Abs(float64(xValue.Int())-float64(yValue.Int())) <= c.numericDelta
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return math.Abs(float64(xValue.Int())-float64(yValue.Uint())) <= c.numericDelta
		case reflect.Float32, reflect.Float64:
			return math.Abs(float64(xValue.Int())-yValue.Float()) <= c.numericDelta
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch yValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return math.Abs(float64(xValue.Uint())-float64(yValue.Int())) <= c.numericDelta
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return math.Abs(float64(xValue.Uint())-float64(yValue.Uint())) <= c.numericDelta
		case reflect.Float32, reflect.Float64:
			return math.Abs(float64(xValue.Uint())-yValue.Float()) <= c.numericDelta
		default:
			return false
		}
	case reflect.Float32, reflect.Float64:
		switch yValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return math.Abs(xValue.Float()-float64(yValue.Int())) <= c.numericDelta
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return math.Abs(xValue.Float()-float64(yValue.Uint())) <= c.numericDelta
		case reflect.Float32, reflect.Float64:
			return math.Abs(xValue.Float()-yValue.Float()) <= c.numericDelta
		default:
			return false
		}
	case reflect.Uintptr:
		if yValue.Kind() != reflect.Uintptr {
			return false
		}

		return xValue.Uint() == yValue.Uint()
	case reflect.Complex64, reflect.Complex128:
		switch yValue.Kind() {
		case reflect.Complex64, reflect.Complex128:
			return xValue.Complex() == yValue.Complex()
		}

		return false
	case reflect.String:
		if yValue.Kind() != reflect.String {
			return false
		}

		return xValue.String() == yValue.String()

	case reflect.Func:
		if yValue.Kind() != reflect.Func {
			return false
		}

		return xValue.Pointer() == yValue.Pointer()
	case reflect.Ptr:
		if yValue.Kind() != reflect.Ptr {
			return false
		}

		if c.samePointer {
			return xValue.Pointer() == yValue.Pointer()
		}
	}

	return false
}

func (c *EqualComparator) buildCmpOptions(x, y interface{}) []cmp.Option {
	return []cmp.Option{
		c.buildEqualComparerOption(),
		c.buildAllowAllUnexported(x, y),
		c.buildIgnoreAllUnexported(),
		c.buildIgnoreFieldOption(),
	}
}

func (c *EqualComparator) buildIgnoreFieldOption() cmp.Option {
	result := cmp.Options{}

	for _, ignoreFields := range c.ignoresFields {
		result = append(result, cmpopts.IgnoreFields(ignoreFields.Type, ignoreFields.Fields...))
	}

	return result
}

func (c *EqualComparator) buildEqualComparerOption() cmp.Option {
	return cmp.FilterValues(
		func(x, y interface{}) bool {
			xKind := reflect.ValueOf(x).Kind()
			yKind := reflect.ValueOf(y).Kind()

			switch xKind {
			case
				reflect.Bool,
				reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64,
				reflect.Uintptr,
				reflect.Float32,
				reflect.Float64,
				reflect.Complex64,
				reflect.Complex128,
				reflect.String,
				reflect.Func:
				switch yKind {
				case
					reflect.Bool,
					reflect.Int,
					reflect.Int8,
					reflect.Int16,
					reflect.Int32,
					reflect.Int64,
					reflect.Uint,
					reflect.Uint8,
					reflect.Uint16,
					reflect.Uint32,
					reflect.Uint64,
					reflect.Uintptr,
					reflect.Float32,
					reflect.Float64,
					reflect.Complex64,
					reflect.Complex128,
					reflect.String,
					reflect.Func:
					return true
				}
			}

			if c.samePointer && xKind == reflect.Ptr && yKind == reflect.Ptr {
				return true
			}

			if c.useEqualMethod {
				if _, ok := x.(Equaler); ok && x != nil {
					return true
				}

				if _, ok := y.(Equaler); ok && y != nil {
					return true
				}
			}

			return false
		},
		cmp.Comparer(c.compareAny),
	)
}

func (c *EqualComparator) buildAllowAllUnexported(allowedTypes ...interface{}) cmp.Option {
	reflectTypes := make(map[reflect.Type]struct{})

	for _, allowedType := range allowedTypes {
		c.fetchUsedTypes(reflect.ValueOf(allowedType), reflectTypes)
	}

	for _, ignoredType := range c.ignoreUnexported {
		delete(reflectTypes, reflect.TypeOf(ignoredType))
	}

	var unexported []interface{}

	for reflectType := range reflectTypes {
		unexported = append(unexported, reflect.New(reflectType).Elem().Interface())
	}

	return cmp.AllowUnexported(unexported...)
}

func (c *EqualComparator) buildIgnoreAllUnexported() cmp.Option {
	return cmpopts.IgnoreUnexported(c.ignoreUnexported...)
}

func (c *EqualComparator) fetchUsedTypes(value reflect.Value, result map[reflect.Type]struct{}) {
	if !value.IsValid() {
		return
	}

	switch value.Kind() {
	case reflect.Ptr:
		if !value.IsNil() {
			c.fetchUsedTypes(value.Elem(), result)
		}
	case reflect.Interface:
		if !value.IsNil() {
			c.fetchUsedTypes(value.Elem(), result)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			c.fetchUsedTypes(value.Index(i), result)
		}
	case reflect.Map:
		for _, k := range value.MapKeys() {
			c.fetchUsedTypes(value.MapIndex(k), result)
		}
	case reflect.Struct:
		if _, ok := result[value.Type()]; ok {
			return
		}

		result[value.Type()] = struct{}{}

		for i := 0; i < value.NumField(); i++ {
			c.fetchUsedTypes(value.Field(i), result)
		}
	}
}
