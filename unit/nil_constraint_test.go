package unit

import (
	"testing"
	"unsafe"
)

func TestNewNilConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(NewNilConstraint).
		ExpectResult(ConstraintAsValue{Value: &NilConstraint{}})
}

func TestNilConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithInvalidValue").
		Call(
			func() bool {
				return (&NilConstraint{}).Check(nil)
			},
		).
		ExpectResult(true)

	NewDeclarative(t, "ValueBool").
		Call((&NilConstraint{}).Check, true).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", true))

	NewDeclarative(t, "ValueInt").
		Call((&NilConstraint{}).Check, int(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int(5)))

	NewDeclarative(t, "ValueInt8").
		Call((&NilConstraint{}).Check, int8(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int8(5)))

	NewDeclarative(t, "ValueInt16").
		Call((&NilConstraint{}).Check, int16(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int16(5)))

	NewDeclarative(t, "ValueInt32").
		Call((&NilConstraint{}).Check, int32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int32(5)))

	NewDeclarative(t, "ValueInt64").
		Call((&NilConstraint{}).Check, int64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int64(5)))

	NewDeclarative(t, "ValueUint").
		Call((&NilConstraint{}).Check, uint(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint(5)))

	NewDeclarative(t, "ValueUint8").
		Call((&NilConstraint{}).Check, uint8(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint8(5)))

	NewDeclarative(t, "ValueUint16").
		Call((&NilConstraint{}).Check, uint16(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint16(5)))

	NewDeclarative(t, "ValueUint32").
		Call((&NilConstraint{}).Check, uint32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint32(5)))

	NewDeclarative(t, "ValueUint64").
		Call((&NilConstraint{}).Check, uint64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint64(5)))

	NewDeclarative(t, "ValueUintptr").
		Call((&NilConstraint{}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uintptr(5)))

	NewDeclarative(t, "ValueFloat32").
		Call((&NilConstraint{}).Check, float32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", float32(5)))

	NewDeclarative(t, "ValueFloat64").
		Call((&NilConstraint{}).Check, float64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", float64(5)))

	NewDeclarative(t, "ValueComplex64").
		Call((&NilConstraint{}).Check, complex64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", complex64(5)))

	NewDeclarative(t, "ValueComplex128").
		Call((&NilConstraint{}).Check, complex128(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", complex128(5)))

	NewDeclarative(t, "ValueArray").
		Call((&NilConstraint{}).Check, [2]int{5, 5}).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", [2]int{5, 5}))

	NewDeclarative(t, "ValueChanWithPositiveResult").
		Call((&NilConstraint{}).Check, (chan int)(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueChanWithNegativeResult").
		Call((&NilConstraint{}).Check, make(chan int)).
		ExpectResult(false)

	NewDeclarative(t, "ValueFuncWithPositiveResult").
		Call((&NilConstraint{}).Check, (func())(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueFuncWithNegativeResult").
		Call((&NilConstraint{}).Check, func() {}).
		ExpectResult(false)

	NewDeclarative(t, "ValueInterfaceWithPositiveResult").
		Call((&NilConstraint{}).Check, (*interface{})(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueMapWithPositiveResult").
		Call((&NilConstraint{}).Check, (map[int]int)(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueMapWithNegativeResult").
		Call((&NilConstraint{}).Check, map[int]int{1: 1, 2: 2}).
		ExpectResult(false)

	NewDeclarative(t, "ValuePtrWithPositiveResult").
		Call((&NilConstraint{}).Check, (*int)(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValuePtrWithNegativeResult").
		Call((&NilConstraint{}).Check, new(int)).
		ExpectResult(false)

	NewDeclarative(t, "ValueSliceWithPositiveResult").
		Call((&NilConstraint{}).Check, ([]int)(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueSliceWithNegativeResult").
		Call((&NilConstraint{}).Check, []int{5, 5}).
		ExpectResult(false)

	NewDeclarative(t, "ValueString").
		Call((&NilConstraint{}).Check, "data").
		ExpectPanic(NewInvalidNilComparisonTypeError("value", "data"))

	NewDeclarative(t, "ValueStruct").
		Call((&NilConstraint{}).Check, struct{}{}).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", struct{}{}))

	NewDeclarative(t, "ValueUnsafePointerWithPositiveResult").
		Call((&NilConstraint{}).Check, (unsafe.Pointer)(nil)).
		ExpectResult(true)

	NewDeclarative(t, "ValueUnsafePointerWithNegativeResult").
		Call((&NilConstraint{}).Check, unsafe.Pointer(new(int))).
		ExpectResult(false)
}

func TestNilConstraint_String(t *testing.T) {
	NewDeclarative(t, "WitPositiveResult").
		Call((&NilConstraint{}).String).
		ExpectResult("be nil")
}

func TestNilConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WitPositiveResult").
		Call((&NilConstraint{}).Details, nil).
		ExpectResult("")
}
