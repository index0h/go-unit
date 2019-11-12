package unit

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestNewTrueConstraint(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewTrueConstraint).
		ExpectResult(ConstraintAsValue{Value: &TrueConstraint{}})
}

func TestTrueConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithInvalidValue").
		Call(
			func() {
				(&TrueConstraint{}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidKindError("value", nil, reflect.Bool))

	NewSubtest(t, "WithBoolAndPositiveResult").
		Call((&TrueConstraint{}).Check, true).
		ExpectResult(true)

	NewSubtest(t, "WithBoolAndNegativeResult").
		Call((&TrueConstraint{}).Check, false).
		ExpectResult(false)

	NewSubtest(t, "WithInt").
		Call((&TrueConstraint{}).Check, 5).
		ExpectPanic(NewInvalidKindError("value", 5, reflect.Bool))

	NewSubtest(t, "WithInt8").
		Call((&TrueConstraint{}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("value", int8(5), reflect.Bool))

	NewSubtest(t, "WithInt16").
		Call((&TrueConstraint{}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("value", int16(5), reflect.Bool))

	NewSubtest(t, "WithInt32").
		Call((&TrueConstraint{}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("value", int32(5), reflect.Bool))

	NewSubtest(t, "WithInt64").
		Call((&TrueConstraint{}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("value", int64(5), reflect.Bool))

	NewSubtest(t, "WithUint").
		Call((&TrueConstraint{}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("value", uint(5), reflect.Bool))

	NewSubtest(t, "WithUint8").
		Call((&TrueConstraint{}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("value", uint8(5), reflect.Bool))

	NewSubtest(t, "WithUint16").
		Call((&TrueConstraint{}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("value", uint16(5), reflect.Bool))

	NewSubtest(t, "WithUint32").
		Call((&TrueConstraint{}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("value", uint32(5), reflect.Bool))

	NewSubtest(t, "WithUint64").
		Call((&TrueConstraint{}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("value", uint64(5), reflect.Bool))

	NewSubtest(t, "WithUintptr").
		Call((&TrueConstraint{}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("value", uintptr(5), reflect.Bool))

	NewSubtest(t, "WithFloat32").
		Call((&TrueConstraint{}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("value", float32(5), reflect.Bool))

	NewSubtest(t, "WithFloat64").
		Call((&TrueConstraint{}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("value", float64(5), reflect.Bool))

	NewSubtest(t, "WithComplex64").
		Call((&TrueConstraint{}).Check, complex64(5)).
		ExpectPanic(NewInvalidKindError("value", complex64(5), reflect.Bool))

	NewSubtest(t, "WithComplex128").
		Call((&TrueConstraint{}).Check, complex128(5)).
		ExpectPanic(NewInvalidKindError("value", complex128(5), reflect.Bool))

	NewSubtest(t, "WithArray").
		Call((&TrueConstraint{}).Check, [1]int{5}).
		ExpectPanic(NewInvalidKindError("value", [1]int{5}, reflect.Bool))

	NewSubtest(t, "WithChan").
		Call((&TrueConstraint{}).Check, make(chan int)).
		ExpectPanic(NewInvalidKindError("value", make(chan int), reflect.Bool))

	NewSubtest(t, "WithFunc").
		Call((&TrueConstraint{}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("value", func() {}, reflect.Bool))

	NewSubtest(t, "WithInterface").
		Call((&TrueConstraint{}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidKindError("value", (*interface{})(nil), reflect.Bool))

	NewSubtest(t, "WithMap").
		Call((&TrueConstraint{}).Check, map[int]int{}).
		ExpectPanic(NewInvalidKindError("value", map[int]int{}, reflect.Bool))

	NewSubtest(t, "WithPointer").
		Call((&TrueConstraint{}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("value", new(int), reflect.Bool))

	NewSubtest(t, "WithSlice").
		Call((&TrueConstraint{}).Check, []int{}).
		ExpectPanic(NewInvalidKindError("value", []int{}, reflect.Bool))

	NewSubtest(t, "WithString").
		Call((&TrueConstraint{}).Check, "data").
		ExpectPanic(NewInvalidKindError("value", "data", reflect.Bool))

	NewSubtest(t, "WithStruct").
		Call((&TrueConstraint{}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("value", struct{}{}, reflect.Bool))

	NewSubtest(t, "WithUnsafePointer").
		Call((&TrueConstraint{}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidKindError("value", unsafe.Pointer(new(int)), reflect.Bool))
}

func TestTrueConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&TrueConstraint{}).String).
		ExpectResult("be true")
}

func TestTrueConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&TrueConstraint{}).Details, 5).
		ExpectResult("")
}
