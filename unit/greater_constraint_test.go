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
	NewDeclarative(t, "ExpectedInvalid").
		Call(
			func() {
				NewGreaterConstraint(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", nil))

	NewDeclarative(t, "ExpectedBool").
		Call(NewGreaterConstraint, true).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", true))

	NewDeclarative(t, "ExpectedInt").
		Call(NewGreaterConstraint, int(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int(5)}})

	NewDeclarative(t, "ExpectedInt8").
		Call(NewGreaterConstraint, int8(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int8(5)}})

	NewDeclarative(t, "ExpectedInt16").
		Call(NewGreaterConstraint, int16(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int16(5)}})

	NewDeclarative(t, "ExpectedInt32").
		Call(NewGreaterConstraint, int32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int32(5)}})

	NewDeclarative(t, "ExpectedInt64").
		Call(NewGreaterConstraint, int64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: int64(5)}})

	NewDeclarative(t, "ExpectedUint").
		Call(NewGreaterConstraint, uint(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint(5)}})

	NewDeclarative(t, "ExpectedUint8").
		Call(NewGreaterConstraint, uint8(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint8(5)}})

	NewDeclarative(t, "ExpectedUint16").
		Call(NewGreaterConstraint, uint16(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint16(5)}})

	NewDeclarative(t, "ExpectedUint32").
		Call(NewGreaterConstraint, uint32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint32(5)}})

	NewDeclarative(t, "ExpectedUint64").
		Call(NewGreaterConstraint, uint64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: uint64(5)}})

	NewDeclarative(t, "ExpectedUintPtr").
		Call(NewGreaterConstraint, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", uintptr(5)))

	NewDeclarative(t, "ExpectedFloat32").
		Call(NewGreaterConstraint, float32(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: float32(5)}})

	NewDeclarative(t, "ExpectedFloat64").
		Call(NewGreaterConstraint, float64(5)).
		ExpectResult(ConstraintAsValue{Value: &GreaterConstraint{expected: float64(5)}})

	NewDeclarative(t, "ExpectedComplex64").
		Call(NewGreaterConstraint, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex64(5)))

	NewDeclarative(t, "ExpectedComplex128").
		Call(NewGreaterConstraint, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex128(5)))

	NewDeclarative(t, "ExpectedArray").
		Call(NewGreaterConstraint, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", [1]int{5}))

	NewDeclarative(t, "ExpectedChan").
		Call(NewGreaterConstraint, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", make(chan int)))

	NewDeclarative(t, "ExpectedFunc").
		Call(NewGreaterConstraint, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", func() {}))

	NewDeclarative(t, "ExpectedInterface").
		Call(NewGreaterConstraint, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", (*interface{})(nil)))

	NewDeclarative(t, "ExpectedMap").
		Call(NewGreaterConstraint, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", map[int]int{1: 1}))

	NewDeclarative(t, "ExpectedPtr").
		Call(NewGreaterConstraint, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", new(int)))

	NewDeclarative(t, "ExpectedSlice").
		Call(NewGreaterConstraint, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", []int{5}))

	NewDeclarative(t, "ExpectedString").
		Call(NewGreaterConstraint, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", "data"))

	NewDeclarative(t, "ExpectedStruct").
		Call(NewGreaterConstraint, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", struct{}{}))

	NewDeclarative(t, "ExpectedUnsafePointer").
		Call(NewGreaterConstraint, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: int(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUintAndIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewDeclarative(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: int(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: int(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: int(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: int(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt8(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: int8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int8(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUintWihtIntComparison").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUintIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int8(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewDeclarative(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int8(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int8(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int8(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int8(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: int8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: int8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: int8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: int8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt16(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: int16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int16(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int16(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewDeclarative(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int16(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int16(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int16(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int16(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: int16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: int16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: int16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: int16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt32(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: int32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int32(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewDeclarative(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int32(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64WihtIntComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int32(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: int32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: int32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: int32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: int32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedInt64(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: int64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: int64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: int64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUintWithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "GreaterThanActualUintWithUintComparison").
			Call((&GreaterConstraint{expected: int64(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)

		NewDeclarative(t, "GreaterThanActualUintWithFloatComparison").
			Call((&GreaterConstraint{expected: int64(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(true)
	}

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64WithIntComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithUintComparison").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "GreaterThanActualUint64WithFloatComparison").
		Call((&GreaterConstraint{expected: int64(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: int64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: int64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: int64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: int64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: int64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: int64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: int64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: int64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: int64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: int64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: int64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: int64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: int64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: int64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: int64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "LessThanActualIntWithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(5)).
			ExpectResult(false)

		NewDeclarative(t, "LessThanActualIntWithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(-5)).
			ExpectResult(false)
	}

	NewDeclarative(t, "LessThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int8(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "LessThanActualInt8WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(5)).
			ExpectResult(false)

		NewDeclarative(t, "LessThanActualInt8WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(-5)).
			ExpectResult(false)
	}

	NewDeclarative(t, "LessThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int16(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "LessThanActualInt16WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(5)).
			ExpectResult(false)

		NewDeclarative(t, "LessThanActualInt16WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(-5)).
			ExpectResult(false)
	}

	NewDeclarative(t, "LessThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int32(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "LessThanActualInt32WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(5)).
			ExpectResult(false)

		NewDeclarative(t, "LessThanActualInt32WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(-5)).
			ExpectResult(false)
	}

	NewDeclarative(t, "LessThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint(5)}).Check, int64(10)).
		ExpectResult(true)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "LessThanActualInt64WithUintComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(5)).
			ExpectResult(false)

		NewDeclarative(t, "LessThanActualInt64WithFloatComparison").
			Call((&GreaterConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(-5)).
			ExpectResult(false)
	}

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: uint(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint8(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint8(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint16(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint16(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint32(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedUint64(t *testing.T) {
	NewDeclarative(t, "Uint64ndInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: uint64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualIntWithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualIntWithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualIntWithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(-5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt8WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(-5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt16WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(-5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt32WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(-5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64WithIntComparison").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64WithUintComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualInt64WithFloatComparison").
		Call((&GreaterConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(-5)).
		ExpectResult(false)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: uint64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: uint64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedFloat32(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: float32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: float32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: float32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: float32(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: float32(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: float32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: float32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: float32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: float32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: float32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: float32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: float32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: float32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: float32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: float32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: float32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: float32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestGreaterConstraint_Check_WithExpectedFloat64(t *testing.T) {
	NewDeclarative(t, "ActualInvalid").
		Call(
			func() {
				(&GreaterConstraint{expected: float64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "ActualBool").
		Call((&GreaterConstraint{expected: float64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "LessThanActualInt").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt8").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt16").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualInt64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualInt64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualInt64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, int64(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint8").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint8").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint8(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint16").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint16").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint16(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualUint64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualUint64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualUint64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uint64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualUintPtr").
		Call((&GreaterConstraint{expected: float64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "LessThanActualFloat32").
		Call((&GreaterConstraint{expected: float64(10)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat32").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float32(10)).
		ExpectResult(true)

	NewDeclarative(t, "LessThanActualFloat64").
		Call((&GreaterConstraint{expected: float64(10)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "EqualToActualFloat64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "GreaterThanActualFloat64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, float64(10)).
		ExpectResult(true)

	NewDeclarative(t, "ActualComplex64").
		Call((&GreaterConstraint{expected: float64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "ActualComplex128").
		Call((&GreaterConstraint{expected: float64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "ActualArray").
		Call((&GreaterConstraint{expected: float64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "ActualChan").
		Call((&GreaterConstraint{expected: float64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "ActualFunc").
		Call((&GreaterConstraint{expected: float64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "ActualInterface").
		Call((&GreaterConstraint{expected: float64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "ActualMap").
		Call((&GreaterConstraint{expected: float64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "ActualPtr").
		Call((&GreaterConstraint{expected: float64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "ActualSlice").
		Call((&GreaterConstraint{expected: float64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "ActualString").
		Call((&GreaterConstraint{expected: float64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "ActualStruct").
		Call((&GreaterConstraint{expected: float64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "ActualUnsafePointer").
		Call((&GreaterConstraint{expected: float64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

func TestGreaterConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&GreaterConstraint{expected: 55}).String).
		ExpectResult(fmt.Sprintf("be greater than %v", 55))
}

func TestGreaterConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&GreaterConstraint{expected: 55}).Details, 10).
		ExpectResult("")
}
