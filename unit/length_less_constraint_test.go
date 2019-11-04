package unit

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewLengthLessConstraint(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call(NewLengthLessConstraint, 5).
		ExpectResult(ConstraintAsValue{Value: &LengthLessConstraint{length: 5}})
}

func TestLengthLessConstraint_Check(t *testing.T) {
	NewDeclarative(t, "WithInvalidValue").
		Call(
			func() {
				(&LengthLessConstraint{length: 5}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", nil))

	NewDeclarative(t, "ValueBool").
		Call((&LengthLessConstraint{length: 5}).Check, true).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", true))

	NewDeclarative(t, "ValueInt").
		Call((&LengthLessConstraint{length: 5}).Check, int(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int(5)))

	NewDeclarative(t, "ValueInt8").
		Call((&LengthLessConstraint{length: 5}).Check, int8(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int8(5)))

	NewDeclarative(t, "ValueInt16").
		Call((&LengthLessConstraint{length: 5}).Check, int16(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int16(5)))

	NewDeclarative(t, "ValueInt32").
		Call((&LengthLessConstraint{length: 5}).Check, int32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int32(5)))

	NewDeclarative(t, "ValueInt64").
		Call((&LengthLessConstraint{length: 5}).Check, int64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int64(5)))

	NewDeclarative(t, "ValueUint").
		Call((&LengthLessConstraint{length: 5}).Check, uint(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint(5)))

	NewDeclarative(t, "ValueUint8").
		Call((&LengthLessConstraint{length: 5}).Check, uint8(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint8(5)))

	NewDeclarative(t, "ValueUint16").
		Call((&LengthLessConstraint{length: 5}).Check, uint16(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint16(5)))

	NewDeclarative(t, "ValueUint32").
		Call((&LengthLessConstraint{length: 5}).Check, uint32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint32(5)))

	NewDeclarative(t, "ValueUint64").
		Call((&LengthLessConstraint{length: 5}).Check, uint64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint64(5)))

	NewDeclarative(t, "ValueUintPtr").
		Call((&LengthLessConstraint{length: 5}).Check, uintptr(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uintptr(5)))

	NewDeclarative(t, "ValueFloat32").
		Call((&LengthLessConstraint{length: 5}).Check, float32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", float32(5)))

	NewDeclarative(t, "ValueFloat64").
		Call((&LengthLessConstraint{length: 5}).Check, float64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", float64(5)))

	NewDeclarative(t, "ValueComplex64").
		Call((&LengthLessConstraint{length: 5}).Check, complex64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", complex64(5)))

	NewDeclarative(t, "ValueComplex128").
		Call((&LengthLessConstraint{length: 5}).Check, complex128(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", complex128(5)))

	NewDeclarative(t, "ValueArrayWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Check, [2]int{5, 5}).
		ExpectResult(true)

	NewDeclarative(t, "ValueArrayWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Check, [2]int{5, 5}).
		ExpectResult(false)

	NewDeclarative(t, "ValueArrayWithNegativeResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Check, [2]int{5, 5}).
		ExpectResult(false)

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithPositiveResultByLessLength").
			Call((&LengthLessConstraint{length: 3}).Check, chanFixture).
			ExpectResult(true)
	}

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithNegativeResultBySameLength").
			Call((&LengthLessConstraint{length: 2}).Check, chanFixture).
			ExpectResult(false)
	}

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithNegativeResultByGreaterLength").
			Call((&LengthLessConstraint{length: 1}).Check, chanFixture).
			ExpectResult(false)
	}

	NewDeclarative(t, "ValueFunc").
		Call((&LengthLessConstraint{length: 5}).Check, func() {}).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", func() {}))

	NewDeclarative(t, "ValueInterface").
		Call((&LengthLessConstraint{length: 5}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", (*interface{})(nil)))

	NewDeclarative(t, "ValueMapWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Check, map[int]int{1: 1, 2: 2}).
		ExpectResult(true)

	NewDeclarative(t, "ValueMapWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Check, map[int]int{1: 1, 2: 2}).
		ExpectResult(false)

	NewDeclarative(t, "ValueMapWithNegativeResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Check, map[int]int{1: 1, 2: 2}).
		ExpectResult(false)

	NewDeclarative(t, "ValuePtr").
		Call((&LengthLessConstraint{length: 5}).Check, new(int)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", new(int)))

	NewDeclarative(t, "ValueSliceWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Check, []int{5, 5}).
		ExpectResult(true)

	NewDeclarative(t, "ValueSliceWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Check, []int{5, 5}).
		ExpectResult(false)

	NewDeclarative(t, "ValueSliceWithNegativeResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Check, []int{5, 5}).
		ExpectResult(false)

	NewDeclarative(t, "ValueStringWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Check, "ab").
		ExpectResult(true)

	NewDeclarative(t, "ValueStringWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Check, "ab").
		ExpectResult(false)

	NewDeclarative(t, "ValueStringWithNegativeResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Check, "ab").
		ExpectResult(false)

	NewDeclarative(t, "ValueStruct").
		Call((&LengthLessConstraint{length: 5}).Check, struct{}{}).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", struct{}{}))

	NewDeclarative(t, "ValueUnsafePointer").
		Call((&LengthLessConstraint{length: 5}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", unsafe.Pointer(new(int))))
}

func TestLengthLessConstraint_String(t *testing.T) {
	NewDeclarative(t, "NegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 55}).String).
		ExpectResult(fmt.Sprintf("have length less than %v", 55))
}

func TestLengthLessConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithInvalidValue").
		Call(
			func() {
				(&LengthLessConstraint{length: 5}).Details(nil)
			},
		).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", nil))

	NewDeclarative(t, "ValueBool").
		Call((&LengthLessConstraint{length: 5}).Details, true).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", true))

	NewDeclarative(t, "ValueInt").
		Call((&LengthLessConstraint{length: 5}).Details, int(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int(5)))

	NewDeclarative(t, "ValueInt8").
		Call((&LengthLessConstraint{length: 5}).Details, int8(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int8(5)))

	NewDeclarative(t, "ValueInt16").
		Call((&LengthLessConstraint{length: 5}).Details, int16(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int16(5)))

	NewDeclarative(t, "ValueInt32").
		Call((&LengthLessConstraint{length: 5}).Details, int32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int32(5)))

	NewDeclarative(t, "ValueInt64").
		Call((&LengthLessConstraint{length: 5}).Details, int64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", int64(5)))

	NewDeclarative(t, "ValueUint").
		Call((&LengthLessConstraint{length: 5}).Details, uint(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint(5)))

	NewDeclarative(t, "ValueUint8").
		Call((&LengthLessConstraint{length: 5}).Details, uint8(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint8(5)))

	NewDeclarative(t, "ValueUint16").
		Call((&LengthLessConstraint{length: 5}).Details, uint16(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint16(5)))

	NewDeclarative(t, "ValueUint32").
		Call((&LengthLessConstraint{length: 5}).Details, uint32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint32(5)))

	NewDeclarative(t, "ValueUint64").
		Call((&LengthLessConstraint{length: 5}).Details, uint64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uint64(5)))

	NewDeclarative(t, "ValueUintPtr").
		Call((&LengthLessConstraint{length: 5}).Details, uintptr(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", uintptr(5)))

	NewDeclarative(t, "ValueFloat32").
		Call((&LengthLessConstraint{length: 5}).Details, float32(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", float32(5)))

	NewDeclarative(t, "ValueFloat64").
		Call((&LengthLessConstraint{length: 5}).Details, float64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", float64(5)))

	NewDeclarative(t, "ValueComplex64").
		Call((&LengthLessConstraint{length: 5}).Details, complex64(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", complex64(5)))

	NewDeclarative(t, "ValueComplex128").
		Call((&LengthLessConstraint{length: 5}).Details, complex128(5)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", complex128(5)))

	NewDeclarative(t, "ValueArrayWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Details, [2]int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueArrayWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Details, [2]int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueArrayWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 1}).Details, [2]int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithResultByLessLength").
			Call((&LengthLessConstraint{length: 3}).Details, chanFixture).
			ExpectResult(fmt.Sprintf("Actual length is %v", 2))
	}

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithNegativeResultBySameLength").
			Call((&LengthLessConstraint{length: 2}).Details, chanFixture).
			ExpectResult(fmt.Sprintf("Actual length is %v", 2))
	}

	{
		chanFixture := make(chan int, 2)
		chanFixture <- 1
		chanFixture <- 1

		NewDeclarative(t, "ValueChanWithResultByGreaterLength").
			Call((&LengthLessConstraint{length: 1}).Details, chanFixture).
			ExpectResult(fmt.Sprintf("Actual length is %v", 2))
	}

	NewDeclarative(t, "ValueFunc").
		Call((&LengthLessConstraint{length: 5}).Details, func() {}).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", func() {}))

	NewDeclarative(t, "ValueInterface").
		Call((&LengthLessConstraint{length: 5}).Details, (*interface{})(nil)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", (*interface{})(nil)))

	NewDeclarative(t, "ValueMapWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Details, map[int]int{1: 1, 2: 2}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueMapWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Details, map[int]int{1: 1, 2: 2}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueMapWithPositiveResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Details, map[int]int{1: 1, 2: 2}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValuePtr").
		Call((&LengthLessConstraint{length: 5}).Details, new(int)).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", new(int)))

	NewDeclarative(t, "ValueSliceWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Details, []int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueSliceWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Details, []int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueSliceWithPositiveResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Details, []int{5, 5}).
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueStringWithPositiveResultByLessLength").
		Call((&LengthLessConstraint{length: 3}).Details, "ab").
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueStringWithNegativeResultBySameLength").
		Call((&LengthLessConstraint{length: 2}).Details, "ab").
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueStringWithPositiveResultByGreaterLength").
		Call((&LengthLessConstraint{length: 1}).Details, "ab").
		ExpectResult(fmt.Sprintf("Actual length is %v", 2))

	NewDeclarative(t, "ValueStruct").
		Call((&LengthLessConstraint{length: 5}).Details, struct{}{}).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", struct{}{}))

	NewDeclarative(t, "ValueUnsafePointer").
		Call((&LengthLessConstraint{length: 5}).Details, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidLengthComparisonTypeError("value", unsafe.Pointer(new(int))))
}
