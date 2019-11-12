package unit

import (
	"testing"
	"unsafe"
)

func TestNewNilConstraint(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewNilConstraint).
		ExpectResult(ConstraintAsValue{Value: &NilConstraint{}})
}

func TestNilConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithInvalidValue").
		Call(
			func() bool {
				return (&NilConstraint{}).Check(nil)
			},
		).
		ExpectResult(true)

	NewSubtest(t, "ValueBool").
		Call((&NilConstraint{}).Check, true).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", true))

	NewSubtest(t, "ValueInt").
		Call((&NilConstraint{}).Check, int(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int(5)))

	NewSubtest(t, "ValueInt8").
		Call((&NilConstraint{}).Check, int8(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int8(5)))

	NewSubtest(t, "ValueInt16").
		Call((&NilConstraint{}).Check, int16(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int16(5)))

	NewSubtest(t, "ValueInt32").
		Call((&NilConstraint{}).Check, int32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int32(5)))

	NewSubtest(t, "ValueInt64").
		Call((&NilConstraint{}).Check, int64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", int64(5)))

	NewSubtest(t, "ValueUint").
		Call((&NilConstraint{}).Check, uint(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint(5)))

	NewSubtest(t, "ValueUint8").
		Call((&NilConstraint{}).Check, uint8(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint8(5)))

	NewSubtest(t, "ValueUint16").
		Call((&NilConstraint{}).Check, uint16(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint16(5)))

	NewSubtest(t, "ValueUint32").
		Call((&NilConstraint{}).Check, uint32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint32(5)))

	NewSubtest(t, "ValueUint64").
		Call((&NilConstraint{}).Check, uint64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uint64(5)))

	NewSubtest(t, "ValueUintptr").
		Call((&NilConstraint{}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", uintptr(5)))

	NewSubtest(t, "ValueFloat32").
		Call((&NilConstraint{}).Check, float32(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", float32(5)))

	NewSubtest(t, "ValueFloat64").
		Call((&NilConstraint{}).Check, float64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", float64(5)))

	NewSubtest(t, "ValueComplex64").
		Call((&NilConstraint{}).Check, complex64(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", complex64(5)))

	NewSubtest(t, "ValueComplex128").
		Call((&NilConstraint{}).Check, complex128(5)).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", complex128(5)))

	NewSubtest(t, "ValueArray").
		Call((&NilConstraint{}).Check, [2]int{5, 5}).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", [2]int{5, 5}))

	NewSubtest(t, "ValueChanWithPositiveResult").
		Call((&NilConstraint{}).Check, (chan int)(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueChanWithNegativeResult").
		Call((&NilConstraint{}).Check, make(chan int)).
		ExpectResult(false)

	NewSubtest(t, "ValueFuncWithPositiveResult").
		Call((&NilConstraint{}).Check, (func())(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueFuncWithNegativeResult").
		Call((&NilConstraint{}).Check, func() {}).
		ExpectResult(false)

	NewSubtest(t, "ValueInterfaceWithPositiveResult").
		Call((&NilConstraint{}).Check, (*interface{})(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueMapWithPositiveResult").
		Call((&NilConstraint{}).Check, (map[int]int)(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueMapWithNegativeResult").
		Call((&NilConstraint{}).Check, map[int]int{1: 1, 2: 2}).
		ExpectResult(false)

	NewSubtest(t, "ValuePtrWithPositiveResult").
		Call((&NilConstraint{}).Check, (*int)(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValuePtrWithNegativeResult").
		Call((&NilConstraint{}).Check, new(int)).
		ExpectResult(false)

	NewSubtest(t, "ValueSliceWithPositiveResult").
		Call((&NilConstraint{}).Check, ([]int)(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueSliceWithNegativeResult").
		Call((&NilConstraint{}).Check, []int{5, 5}).
		ExpectResult(false)

	NewSubtest(t, "ValueString").
		Call((&NilConstraint{}).Check, "data").
		ExpectPanic(NewInvalidNilComparisonTypeError("value", "data"))

	NewSubtest(t, "ValueStruct").
		Call((&NilConstraint{}).Check, struct{}{}).
		ExpectPanic(NewInvalidNilComparisonTypeError("value", struct{}{}))

	NewSubtest(t, "ValueUnsafePointerWithPositiveResult").
		Call((&NilConstraint{}).Check, (unsafe.Pointer)(nil)).
		ExpectResult(true)

	NewSubtest(t, "ValueUnsafePointerWithNegativeResult").
		Call((&NilConstraint{}).Check, unsafe.Pointer(new(int))).
		ExpectResult(false)
}

func TestNilConstraint_String(t *testing.T) {
	NewSubtest(t, "WitPositiveResult").
		Call((&NilConstraint{}).String).
		ExpectResult("be nil")
}

func TestNilConstraint_Details(t *testing.T) {
	NewSubtest(t, "WitPositiveResult").
		Call((&NilConstraint{}).Details, nil).
		ExpectResult("")
}
