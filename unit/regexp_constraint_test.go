package unit

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"unsafe"
)

func TestNewRegexpConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(NewRegexpConstraint, "^\\d+$").
		ExpectResult(ConstraintAsValue{Value: &RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}})
}

func TestRegexpConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithInvalidValue").
		Call(
			func() {
				(&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidKindError("value", nil, reflect.String))

	NewDeclarative(t, "WithBool").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, true).
		ExpectPanic(NewInvalidKindError("value", true, reflect.String))

	NewDeclarative(t, "WithInt").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, 5).
		ExpectPanic(NewInvalidKindError("value", 5, reflect.String))

	NewDeclarative(t, "WithInt8").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("value", int8(5), reflect.String))

	NewDeclarative(t, "WithInt16").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("value", int16(5), reflect.String))

	NewDeclarative(t, "WithInt32").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("value", int32(5), reflect.String))

	NewDeclarative(t, "WithInt64").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("value", int64(5), reflect.String))

	NewDeclarative(t, "WithUint").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("value", uint(5), reflect.String))

	NewDeclarative(t, "WithUint8").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("value", uint8(5), reflect.String))

	NewDeclarative(t, "WithUint16").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("value", uint16(5), reflect.String))

	NewDeclarative(t, "WithUint32").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("value", uint32(5), reflect.String))

	NewDeclarative(t, "WithUint64").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("value", uint64(5), reflect.String))

	NewDeclarative(t, "WithUintptr").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("value", uintptr(5), reflect.String))

	NewDeclarative(t, "WithFloat32").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("value", float32(5), reflect.String))

	NewDeclarative(t, "WithFloat64").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("value", float64(5), reflect.String))

	NewDeclarative(t, "WithComplex64").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, complex64(5)).
		ExpectPanic(NewInvalidKindError("value", complex64(5), reflect.String))

	NewDeclarative(t, "WithComplex128").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, complex128(5)).
		ExpectPanic(NewInvalidKindError("value", complex128(5), reflect.String))

	NewDeclarative(t, "WithArray").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, [1]int{5}).
		ExpectPanic(NewInvalidKindError("value", [1]int{5}, reflect.String))

	NewDeclarative(t, "WithChan").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, make(chan int)).
		ExpectPanic(NewInvalidKindError("value", make(chan int), reflect.String))

	NewDeclarative(t, "WithFunc").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("value", func() {}, reflect.String))

	NewDeclarative(t, "WithInterface").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidKindError("value", (*interface{})(nil), reflect.String))

	NewDeclarative(t, "WithMap").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, map[int]int{}).
		ExpectPanic(NewInvalidKindError("value", map[int]int{}, reflect.String))

	NewDeclarative(t, "WithPointer").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("value", new(int), reflect.String))

	NewDeclarative(t, "WithSlice").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, []int{}).
		ExpectPanic(NewInvalidKindError("value", []int{}, reflect.String))

	NewDeclarative(t, "WithStringAndPositiveResult").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^[dat]{4}$")}).Check, "data").
		ExpectResult(true)

	NewDeclarative(t, "WithStringAndNegativeResult").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d{10}$")}).Check, "data").
		ExpectResult(false)

	NewDeclarative(t, "WithStruct").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("value", struct{}{}, reflect.String))

	NewDeclarative(t, "WithUnsafePointer").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidKindError("value", unsafe.Pointer(new(int)), reflect.String))
}

func TestRegexpConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d{10}$")}).String).
		ExpectResult(fmt.Sprintf("match PCRE pattern '%s'", "^\\d{10}$"))
}

func TestRegexpConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&RegexpConstraint{pattern: regexp.MustCompile("^\\d+$")}).Details, 5).
		ExpectResult("")
}
