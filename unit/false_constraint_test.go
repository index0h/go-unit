package unit

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestNewFalseConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(NewFalseConstraint).
		ExpectResult(ConstraintAsValue{Value: &FalseConstraint{}})
}

func TestFalseConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithInvalidValue").
		Call(
			func() {
				(&FalseConstraint{}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidKindError("value", nil, reflect.Bool))

	NewDeclarative(t, "WithBoolAndPositiveResult").
		Call((&FalseConstraint{}).Check, false).
		ExpectResult(true)

	NewDeclarative(t, "WithBoolAndNegativeResult").
		Call((&FalseConstraint{}).Check, true).
		ExpectResult(false)

	NewDeclarative(t, "WithInt").
		Call((&FalseConstraint{}).Check, 5).
		ExpectPanic(NewInvalidKindError("value", 5, reflect.Bool))

	NewDeclarative(t, "WithInt8").
		Call((&FalseConstraint{}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("value", int8(5), reflect.Bool))

	NewDeclarative(t, "WithInt16").
		Call((&FalseConstraint{}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("value", int16(5), reflect.Bool))

	NewDeclarative(t, "WithInt32").
		Call((&FalseConstraint{}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("value", int32(5), reflect.Bool))

	NewDeclarative(t, "WithInt64").
		Call((&FalseConstraint{}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("value", int64(5), reflect.Bool))

	NewDeclarative(t, "WithUint").
		Call((&FalseConstraint{}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("value", uint(5), reflect.Bool))

	NewDeclarative(t, "WithUint8").
		Call((&FalseConstraint{}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("value", uint8(5), reflect.Bool))

	NewDeclarative(t, "WithUint16").
		Call((&FalseConstraint{}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("value", uint16(5), reflect.Bool))

	NewDeclarative(t, "WithUint32").
		Call((&FalseConstraint{}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("value", uint32(5), reflect.Bool))

	NewDeclarative(t, "WithUint64").
		Call((&FalseConstraint{}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("value", uint64(5), reflect.Bool))

	NewDeclarative(t, "WithUintptr").
		Call((&FalseConstraint{}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("value", uintptr(5), reflect.Bool))

	NewDeclarative(t, "WithFloat32").
		Call((&FalseConstraint{}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("value", float32(5), reflect.Bool))

	NewDeclarative(t, "WithFloat64").
		Call((&FalseConstraint{}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("value", float64(5), reflect.Bool))

	NewDeclarative(t, "WithComplex64").
		Call((&FalseConstraint{}).Check, complex64(5)).
		ExpectPanic(NewInvalidKindError("value", complex64(5), reflect.Bool))

	NewDeclarative(t, "WithComplex128").
		Call((&FalseConstraint{}).Check, complex128(5)).
		ExpectPanic(NewInvalidKindError("value", complex128(5), reflect.Bool))

	NewDeclarative(t, "WithArray").
		Call((&FalseConstraint{}).Check, [1]int{5}).
		ExpectPanic(NewInvalidKindError("value", [1]int{5}, reflect.Bool))

	NewDeclarative(t, "WithChan").
		Call((&FalseConstraint{}).Check, make(chan int)).
		ExpectPanic(NewInvalidKindError("value", make(chan int), reflect.Bool))

	NewDeclarative(t, "WithFunc").
		Call((&FalseConstraint{}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("value", func() {}, reflect.Bool))

	NewDeclarative(t, "WithInterface").
		Call((&FalseConstraint{}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidKindError("value", (*interface{})(nil), reflect.Bool))

	NewDeclarative(t, "WithMap").
		Call((&FalseConstraint{}).Check, map[int]int{}).
		ExpectPanic(NewInvalidKindError("value", map[int]int{}, reflect.Bool))

	NewDeclarative(t, "WithPointer").
		Call((&FalseConstraint{}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("value", new(int), reflect.Bool))

	NewDeclarative(t, "WithSlice").
		Call((&FalseConstraint{}).Check, []int{}).
		ExpectPanic(NewInvalidKindError("value", []int{}, reflect.Bool))

	NewDeclarative(t, "WithString").
		Call((&FalseConstraint{}).Check, "data").
		ExpectPanic(NewInvalidKindError("value", "data", reflect.Bool))

	NewDeclarative(t, "WithStruct").
		Call((&FalseConstraint{}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("value", struct{}{}, reflect.Bool))

	NewDeclarative(t, "WithUnsafePointer").
		Call((&FalseConstraint{}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidKindError("value", unsafe.Pointer(new(int)), reflect.Bool))
}

func TestFalseConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&FalseConstraint{}).String).
		ExpectResult("be false")
}

func TestFalseConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&FalseConstraint{}).Details, 5).
		ExpectResult("")
}
