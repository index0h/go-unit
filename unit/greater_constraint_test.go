package unit

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

//noinspection GoRedundantConversion
func TestNewGreaterConstraint(t *testing.T) {
	NewSubtest(t, "ExpectedInvalid").
		Call(
			func() {
				NewGreaterConstraint(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", nil))

	NewSubtest(t, "ExpectedBool").
		Call(NewGreaterConstraint, true).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", true))

	NewSubtest(t, "ExpectedInt").
		Call(NewGreaterConstraint, int(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int(5)}})

	NewSubtest(t, "ExpectedInt8").
		Call(NewGreaterConstraint, int8(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int8(5)}})

	NewSubtest(t, "ExpectedInt16").
		Call(NewGreaterConstraint, int16(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int16(5)}})

	NewSubtest(t, "ExpectedInt32").
		Call(NewGreaterConstraint, int32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int32(5)}})

	NewSubtest(t, "ExpectedInt64").
		Call(NewGreaterConstraint, int64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int64(5)}})

	NewSubtest(t, "ExpectedUint").
		Call(NewGreaterConstraint, uint(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint(5)}})

	NewSubtest(t, "ExpectedUint8").
		Call(NewGreaterConstraint, uint8(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint8(5)}})

	NewSubtest(t, "ExpectedUint16").
		Call(NewGreaterConstraint, uint16(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint16(5)}})

	NewSubtest(t, "ExpectedUint32").
		Call(NewGreaterConstraint, uint32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint32(5)}})

	NewSubtest(t, "ExpectedUint64").
		Call(NewGreaterConstraint, uint64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint64(5)}})

	NewSubtest(t, "ExpectedUintPtr").
		Call(NewGreaterConstraint, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", uintptr(5)))

	NewSubtest(t, "ExpectedFloat32").
		Call(NewGreaterConstraint, float32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: float32(5)}})

	NewSubtest(t, "ExpectedFloat64").
		Call(NewGreaterConstraint, float64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: float64(5)}})

	NewSubtest(t, "ExpectedComplex64").
		Call(NewGreaterConstraint, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex64(5)))

	NewSubtest(t, "ExpectedComplex128").
		Call(NewGreaterConstraint, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex128(5)))

	NewSubtest(t, "ExpectedArray").
		Call(NewGreaterConstraint, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", [1]int{5}))

	NewSubtest(t, "ExpectedChan").
		Call(NewGreaterConstraint, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", make(chan int)))

	NewSubtest(t, "ExpectedFunc").
		Call(NewGreaterConstraint, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", func() {}))

	NewSubtest(t, "ExpectedInterface").
		Call(NewGreaterConstraint, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", (*interface{})(nil)))

	NewSubtest(t, "ExpectedMap").
		Call(NewGreaterConstraint, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", map[int]int{1: 1}))

	NewSubtest(t, "ExpectedPtr").
		Call(NewGreaterConstraint, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", new(int)))

	NewSubtest(t, "ExpectedSlice").
		Call(NewGreaterConstraint, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", []int{5}))

	NewSubtest(t, "ExpectedString").
		Call(NewGreaterConstraint, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", "data"))

	NewSubtest(t, "ExpectedStruct").
		Call(NewGreaterConstraint, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", struct{}{}))

	NewSubtest(t, "ExpectedUnsafePointer").
		Call(NewGreaterConstraint, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: int(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUintAndIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewSubtest(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: int(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: int(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: int(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: int(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt8(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: int8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUintWihtIntComparison").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUintIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int8(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewSubtest(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int8(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int8(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int8(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: int8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: int8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: int8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: int8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt16(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: int16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int16(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewSubtest(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int16(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int16(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int16(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: int16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: int16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: int16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: int16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt32(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: int32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int32(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewSubtest(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int32(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64WihtIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int32(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: int32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: int32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: int32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: int32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt64(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: int64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int64(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewSubtest(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int64(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int64(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: int64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: int64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: int64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: int64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "LessThanActualIntWithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(5)).
			ExpectResult(false)

		NewSubtest(t, "LessThanActualIntWithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(-5)).
			ExpectResult(false)
	}

	NewSubtest(t, "LessThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int8(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "LessThanActualInt8WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(5)).
			ExpectResult(false)

		NewSubtest(t, "LessThanActualInt8WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(-5)).
			ExpectResult(false)
	}

	NewSubtest(t, "LessThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int16(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "LessThanActualInt16WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(5)).
			ExpectResult(false)

		NewSubtest(t, "LessThanActualInt16WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(-5)).
			ExpectResult(false)
	}

	NewSubtest(t, "LessThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int32(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "LessThanActualInt32WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(5)).
			ExpectResult(false)

		NewSubtest(t, "LessThanActualInt32WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(-5)).
			ExpectResult(false)
	}

	NewSubtest(t, "LessThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int64(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "LessThanActualInt64WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(5)).
			ExpectResult(false)

		NewSubtest(t, "LessThanActualInt64WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(-5)).
			ExpectResult(false)
	}

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: uint(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint8(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint16(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint32(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint64(t *testing.T) {
	NewSubtest(t, "Uint64ndInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualIntWithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualIntWithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(-5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt8WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(-5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt16WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(-5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt32WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(-5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualInt64WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(-5)).
		ExpectResult(false)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedFloat32(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: float32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: float32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: float32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: float32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: float32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: float32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: float32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: float32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: float32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: float32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: float32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: float32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: float32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedFloat64(t *testing.T) {
	NewSubtest(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: float64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "ActualBool").
		Call((&GreaterConstraint{expected: float64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewSubtest(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewSubtest(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: float64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "ActualArray").
		Call((&GreaterConstraint{expected: float64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "ActualChan").
		Call((&GreaterConstraint{expected: float64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "ActualFunc").
		Call((&GreaterConstraint{expected: float64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "ActualInterface").
		Call((&GreaterConstraint{expected: float64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "ActualMap").
		Call((&GreaterConstraint{expected: float64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "ActualPtr").
		Call((&GreaterConstraint{expected: float64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "ActualSlice").
		Call((&GreaterConstraint{expected: float64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "ActualString").
		Call((&GreaterConstraint{expected: float64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "ActualStruct").
		Call((&GreaterConstraint{expected: float64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: float64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

func TestGreaterConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&GreaterConstraint{expected: 55}).String).
		ExpectResult(fmt.Sprintf("be greater than %v", 55))
}

func TestGreaterConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&GreaterConstraint{expected: 55}).Details, 10).
		ExpectResult("")
}
