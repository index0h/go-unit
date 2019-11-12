package unit

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

//noinspection GoRedundantConversion
func TestNewLessConstraint(t *testing.T) {
	NewSubtest(t, "Invalid").
		Call(
			func() {
				NewLessConstraint(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", nil))

	NewSubtest(t, "Bool").
		Call(NewLessConstraint, true).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", true))

	NewSubtest(t, "Int").
		Call(NewLessConstraint, int(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int(5)}})

	NewSubtest(t, "Int8").
		Call(NewLessConstraint, int8(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int8(5)}})

	NewSubtest(t, "Int16").
		Call(NewLessConstraint, int16(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int16(5)}})

	NewSubtest(t, "Int32").
		Call(NewLessConstraint, int32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int32(5)}})

	NewSubtest(t, "Int64").
		Call(NewLessConstraint, int64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int64(5)}})

	NewSubtest(t, "Uint").
		Call(NewLessConstraint, uint(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint(5)}})

	NewSubtest(t, "Uint8").
		Call(NewLessConstraint, uint8(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint8(5)}})

	NewSubtest(t, "Uint16").
		Call(NewLessConstraint, uint16(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint16(5)}})

	NewSubtest(t, "Uint32").
		Call(NewLessConstraint, uint32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint32(5)}})

	NewSubtest(t, "Uint64").
		Call(NewLessConstraint, uint64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint64(5)}})

	NewSubtest(t, "UintPtr").
		Call(NewLessConstraint, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", uintptr(5)))

	NewSubtest(t, "Float32").
		Call(NewLessConstraint, float32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: float32(5)}})

	NewSubtest(t, "Float64").
		Call(NewLessConstraint, float64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: float64(5)}})

	NewSubtest(t, "Complex64").
		Call(NewLessConstraint, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex64(5)))

	NewSubtest(t, "Complex128").
		Call(NewLessConstraint, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex128(5)))

	NewSubtest(t, "Array").
		Call(NewLessConstraint, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", [1]int{5}))

	NewSubtest(t, "Chan").
		Call(NewLessConstraint, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", make(chan int)))

	NewSubtest(t, "Func").
		Call(NewLessConstraint, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", func() {}))

	NewSubtest(t, "Interface").
		Call(NewLessConstraint, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", (*interface{})(nil)))

	NewSubtest(t, "Map").
		Call(NewLessConstraint, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", map[int]int{1: 1}))

	NewSubtest(t, "Ptr").
		Call(NewLessConstraint, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", new(int)))

	NewSubtest(t, "Slice").
		Call(NewLessConstraint, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", []int{5}))

	NewSubtest(t, "String").
		Call(NewLessConstraint, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", "data"))

	NewSubtest(t, "Struct").
		Call(NewLessConstraint, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", struct{}{}))

	NewSubtest(t, "UnsafePointer").
		Call(NewLessConstraint, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt(t *testing.T) {
	NewSubtest(t, "IntAndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "IntAndBool").
		Call((&LessConstraint{expected: int(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "IntLessInt").
		Call((&LessConstraint{expected: int(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualInt").
		Call((&LessConstraint{expected: int(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterInt").
		Call((&LessConstraint{expected: int(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessInt8").
		Call((&LessConstraint{expected: int(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualInt8").
		Call((&LessConstraint{expected: int(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterInt8").
		Call((&LessConstraint{expected: int(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessInt16").
		Call((&LessConstraint{expected: int(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualInt16").
		Call((&LessConstraint{expected: int(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterInt16").
		Call((&LessConstraint{expected: int(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessInt32").
		Call((&LessConstraint{expected: int(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualInt32").
		Call((&LessConstraint{expected: int(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterInt32").
		Call((&LessConstraint{expected: int(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessInt64").
		Call((&LessConstraint{expected: int(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualInt64").
		Call((&LessConstraint{expected: int(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterInt64").
		Call((&LessConstraint{expected: int(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessUintWithIntComparison").
		Call((&LessConstraint{expected: int(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualUintWithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUintAndIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "IntGreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewSubtest(t, "IntGreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewSubtest(t, "IntLessUint8").
		Call((&LessConstraint{expected: int(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualUint8").
		Call((&LessConstraint{expected: int(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint8").
		Call((&LessConstraint{expected: int(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessUint16").
		Call((&LessConstraint{expected: int(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualUint16").
		Call((&LessConstraint{expected: int(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint16").
		Call((&LessConstraint{expected: int(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessUint32").
		Call((&LessConstraint{expected: int(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualUint32").
		Call((&LessConstraint{expected: int(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint32").
		Call((&LessConstraint{expected: int(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessUint64WithIntComparison").
		Call((&LessConstraint{expected: int(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "IntAndUintPtr").
		Call((&LessConstraint{expected: int(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "IntLessFloat32").
		Call((&LessConstraint{expected: int(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualFloat32").
		Call((&LessConstraint{expected: int(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterFloat32").
		Call((&LessConstraint{expected: int(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "IntLessFloat64").
		Call((&LessConstraint{expected: int(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "IntEqualFloat64").
		Call((&LessConstraint{expected: int(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "IntGreaterFloat64").
		Call((&LessConstraint{expected: int(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "IntAndComplex64").
		Call((&LessConstraint{expected: int(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "IntAndComplex128").
		Call((&LessConstraint{expected: int(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "IntAndArray").
		Call((&LessConstraint{expected: int(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "IntAndChan").
		Call((&LessConstraint{expected: int(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "IntAndFunc").
		Call((&LessConstraint{expected: int(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "IntAndInterface").
		Call((&LessConstraint{expected: int(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "IntAndMap").
		Call((&LessConstraint{expected: int(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "IntAndPtr").
		Call((&LessConstraint{expected: int(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "IntAndSlice").
		Call((&LessConstraint{expected: int(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "IntAndString").
		Call((&LessConstraint{expected: int(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "IntAndStruct").
		Call((&LessConstraint{expected: int(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "IntAndUnsafePointer").
		Call((&LessConstraint{expected: int(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt8(t *testing.T) {
	NewSubtest(t, "Int8AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Int8AndBool").
		Call((&LessConstraint{expected: int8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Int8LessInt").
		Call((&LessConstraint{expected: int8(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualInt").
		Call((&LessConstraint{expected: int8(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterInt").
		Call((&LessConstraint{expected: int8(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessInt8").
		Call((&LessConstraint{expected: int8(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualInt8").
		Call((&LessConstraint{expected: int8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterInt8").
		Call((&LessConstraint{expected: int8(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessInt16").
		Call((&LessConstraint{expected: int8(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualInt16").
		Call((&LessConstraint{expected: int8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterInt16").
		Call((&LessConstraint{expected: int8(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessInt32").
		Call((&LessConstraint{expected: int8(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualInt32").
		Call((&LessConstraint{expected: int8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterInt32").
		Call((&LessConstraint{expected: int8(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessInt64").
		Call((&LessConstraint{expected: int8(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualInt64").
		Call((&LessConstraint{expected: int8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterInt64").
		Call((&LessConstraint{expected: int8(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessUintWihtIntComparison").
		Call((&LessConstraint{expected: int8(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUintIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "Int8GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int8(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewSubtest(t, "Int8GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int8(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewSubtest(t, "Int8LessUint8").
		Call((&LessConstraint{expected: int8(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualUint8").
		Call((&LessConstraint{expected: int8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint8").
		Call((&LessConstraint{expected: int8(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessUint16").
		Call((&LessConstraint{expected: int8(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualUint16").
		Call((&LessConstraint{expected: int8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint16").
		Call((&LessConstraint{expected: int8(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessUint32").
		Call((&LessConstraint{expected: int8(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualUint32").
		Call((&LessConstraint{expected: int8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint32").
		Call((&LessConstraint{expected: int8(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int8(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int8AndUintPtr").
		Call((&LessConstraint{expected: int8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Int8LessFloat32").
		Call((&LessConstraint{expected: int8(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualFloat32").
		Call((&LessConstraint{expected: int8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterFloat32").
		Call((&LessConstraint{expected: int8(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8LessFloat64").
		Call((&LessConstraint{expected: int8(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int8EqualFloat64").
		Call((&LessConstraint{expected: int8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int8GreaterFloat64").
		Call((&LessConstraint{expected: int8(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int8AndComplex64").
		Call((&LessConstraint{expected: int8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Int8AndComplex128").
		Call((&LessConstraint{expected: int8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Int8AndArray").
		Call((&LessConstraint{expected: int8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Int8AndChan").
		Call((&LessConstraint{expected: int8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Int8AndFunc").
		Call((&LessConstraint{expected: int8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Int8AndInterface").
		Call((&LessConstraint{expected: int8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Int8AndMap").
		Call((&LessConstraint{expected: int8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Int8AndPtr").
		Call((&LessConstraint{expected: int8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Int8AndSlice").
		Call((&LessConstraint{expected: int8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Int8AndString").
		Call((&LessConstraint{expected: int8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Int8AndStruct").
		Call((&LessConstraint{expected: int8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Int8AndUnsafePointer").
		Call((&LessConstraint{expected: int8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt16(t *testing.T) {
	NewSubtest(t, "Int16AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Int16AndBool").
		Call((&LessConstraint{expected: int16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Int16LessInt").
		Call((&LessConstraint{expected: int16(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualInt").
		Call((&LessConstraint{expected: int16(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterInt").
		Call((&LessConstraint{expected: int16(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessInt8").
		Call((&LessConstraint{expected: int16(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualInt8").
		Call((&LessConstraint{expected: int16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterInt8").
		Call((&LessConstraint{expected: int16(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessInt16").
		Call((&LessConstraint{expected: int16(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualInt16").
		Call((&LessConstraint{expected: int16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterInt16").
		Call((&LessConstraint{expected: int16(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessInt32").
		Call((&LessConstraint{expected: int16(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualInt32").
		Call((&LessConstraint{expected: int16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterInt32").
		Call((&LessConstraint{expected: int16(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessInt64").
		Call((&LessConstraint{expected: int16(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualInt64").
		Call((&LessConstraint{expected: int16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterInt64").
		Call((&LessConstraint{expected: int16(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessUintWithIntComparison").
		Call((&LessConstraint{expected: int16(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "Int16GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int16(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewSubtest(t, "Int16GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int16(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewSubtest(t, "Int16LessUint8").
		Call((&LessConstraint{expected: int16(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualUint8").
		Call((&LessConstraint{expected: int16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint8").
		Call((&LessConstraint{expected: int16(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessUint16").
		Call((&LessConstraint{expected: int16(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualUint16").
		Call((&LessConstraint{expected: int16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint16").
		Call((&LessConstraint{expected: int16(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessUint32").
		Call((&LessConstraint{expected: int16(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualUint32").
		Call((&LessConstraint{expected: int16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint32").
		Call((&LessConstraint{expected: int16(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int16(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int16AndUintPtr").
		Call((&LessConstraint{expected: int16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Int16LessFloat32").
		Call((&LessConstraint{expected: int16(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualFloat32").
		Call((&LessConstraint{expected: int16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterFloat32").
		Call((&LessConstraint{expected: int16(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16LessFloat64").
		Call((&LessConstraint{expected: int16(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int16EqualFloat64").
		Call((&LessConstraint{expected: int16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int16GreaterFloat64").
		Call((&LessConstraint{expected: int16(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int16AndComplex64").
		Call((&LessConstraint{expected: int16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Int16AndComplex128").
		Call((&LessConstraint{expected: int16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Int16AndArray").
		Call((&LessConstraint{expected: int16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Int16AndChan").
		Call((&LessConstraint{expected: int16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Int16AndFunc").
		Call((&LessConstraint{expected: int16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Int16AndInterface").
		Call((&LessConstraint{expected: int16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Int16AndMap").
		Call((&LessConstraint{expected: int16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Int16AndPtr").
		Call((&LessConstraint{expected: int16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Int16AndSlice").
		Call((&LessConstraint{expected: int16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Int16AndString").
		Call((&LessConstraint{expected: int16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Int16AndStruct").
		Call((&LessConstraint{expected: int16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Int16AndUnsafePointer").
		Call((&LessConstraint{expected: int16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt32(t *testing.T) {
	NewSubtest(t, "Int32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Int32AndBool").
		Call((&LessConstraint{expected: int32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Int32LessInt").
		Call((&LessConstraint{expected: int32(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualInt").
		Call((&LessConstraint{expected: int32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterInt").
		Call((&LessConstraint{expected: int32(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessInt8").
		Call((&LessConstraint{expected: int32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualInt8").
		Call((&LessConstraint{expected: int32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterInt8").
		Call((&LessConstraint{expected: int32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessInt16").
		Call((&LessConstraint{expected: int32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualInt16").
		Call((&LessConstraint{expected: int32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterInt16").
		Call((&LessConstraint{expected: int32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessInt32").
		Call((&LessConstraint{expected: int32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualInt32").
		Call((&LessConstraint{expected: int32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterInt32").
		Call((&LessConstraint{expected: int32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessInt64").
		Call((&LessConstraint{expected: int32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualInt64").
		Call((&LessConstraint{expected: int32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterInt64").
		Call((&LessConstraint{expected: int32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessUintWithIntComparison").
		Call((&LessConstraint{expected: int32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "Int32GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int32(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewSubtest(t, "Int32GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int32(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewSubtest(t, "Int32LessUint8").
		Call((&LessConstraint{expected: int32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualUint8").
		Call((&LessConstraint{expected: int32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint8").
		Call((&LessConstraint{expected: int32(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessUint16").
		Call((&LessConstraint{expected: int32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualUint16").
		Call((&LessConstraint{expected: int32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint16").
		Call((&LessConstraint{expected: int32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessUint32").
		Call((&LessConstraint{expected: int32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualUint32").
		Call((&LessConstraint{expected: int32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint32").
		Call((&LessConstraint{expected: int32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint64WihtIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int32(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int32AndUintPtr").
		Call((&LessConstraint{expected: int32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Int32LessFloat32").
		Call((&LessConstraint{expected: int32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualFloat32").
		Call((&LessConstraint{expected: int32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterFloat32").
		Call((&LessConstraint{expected: int32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32LessFloat64").
		Call((&LessConstraint{expected: int32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int32EqualFloat64").
		Call((&LessConstraint{expected: int32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int32GreaterFloat64").
		Call((&LessConstraint{expected: int32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int32AndComplex64").
		Call((&LessConstraint{expected: int32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Int32AndComplex128").
		Call((&LessConstraint{expected: int32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Int32AndArray").
		Call((&LessConstraint{expected: int32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Int32AndChan").
		Call((&LessConstraint{expected: int32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Int32AndFunc").
		Call((&LessConstraint{expected: int32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Int32AndInterface").
		Call((&LessConstraint{expected: int32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Int32AndMap").
		Call((&LessConstraint{expected: int32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Int32AndPtr").
		Call((&LessConstraint{expected: int32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Int32AndSlice").
		Call((&LessConstraint{expected: int32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Int32AndString").
		Call((&LessConstraint{expected: int32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Int32AndStruct").
		Call((&LessConstraint{expected: int32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Int32AndUnsafePointer").
		Call((&LessConstraint{expected: int32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt64(t *testing.T) {
	NewSubtest(t, "Int64AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Int64AndBool").
		Call((&LessConstraint{expected: int64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Int64LessInt").
		Call((&LessConstraint{expected: int64(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualInt").
		Call((&LessConstraint{expected: int64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterInt").
		Call((&LessConstraint{expected: int64(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessInt8").
		Call((&LessConstraint{expected: int64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualInt8").
		Call((&LessConstraint{expected: int64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterInt8").
		Call((&LessConstraint{expected: int64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessInt16").
		Call((&LessConstraint{expected: int64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualInt16").
		Call((&LessConstraint{expected: int64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterInt16").
		Call((&LessConstraint{expected: int64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessInt32").
		Call((&LessConstraint{expected: int64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualInt32").
		Call((&LessConstraint{expected: int64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterInt32").
		Call((&LessConstraint{expected: int64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessInt64").
		Call((&LessConstraint{expected: int64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualInt64").
		Call((&LessConstraint{expected: int64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterInt64").
		Call((&LessConstraint{expected: int64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessUintWithIntComparison").
		Call((&LessConstraint{expected: int64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewSubtest(t, "Int64GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int64(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewSubtest(t, "Int64GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int64(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewSubtest(t, "Int64LessUint8").
		Call((&LessConstraint{expected: int64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualUint8").
		Call((&LessConstraint{expected: int64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint8").
		Call((&LessConstraint{expected: int64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessUint16").
		Call((&LessConstraint{expected: int64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualUint16").
		Call((&LessConstraint{expected: int64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint16").
		Call((&LessConstraint{expected: int64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessUint32").
		Call((&LessConstraint{expected: int64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualUint32").
		Call((&LessConstraint{expected: int64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint32").
		Call((&LessConstraint{expected: int64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int64(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewSubtest(t, "Int64AndUintPtr").
		Call((&LessConstraint{expected: int64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Int64LessFloat32").
		Call((&LessConstraint{expected: int64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualFloat32").
		Call((&LessConstraint{expected: int64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterFloat32").
		Call((&LessConstraint{expected: int64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64LessFloat64").
		Call((&LessConstraint{expected: int64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Int64EqualFloat64").
		Call((&LessConstraint{expected: int64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Int64GreaterFloat64").
		Call((&LessConstraint{expected: int64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Int64AndComplex64").
		Call((&LessConstraint{expected: int64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Int64AndComplex128").
		Call((&LessConstraint{expected: int64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Int64AndArray").
		Call((&LessConstraint{expected: int64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Int64AndChan").
		Call((&LessConstraint{expected: int64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Int64AndFunc").
		Call((&LessConstraint{expected: int64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Int64AndInterface").
		Call((&LessConstraint{expected: int64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Int64AndMap").
		Call((&LessConstraint{expected: int64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Int64AndPtr").
		Call((&LessConstraint{expected: int64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Int64AndSlice").
		Call((&LessConstraint{expected: int64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Int64AndString").
		Call((&LessConstraint{expected: int64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Int64AndStruct").
		Call((&LessConstraint{expected: int64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Int64AndUnsafePointer").
		Call((&LessConstraint{expected: int64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint(t *testing.T) {
	NewSubtest(t, "UintAndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "UintAndBool").
		Call((&LessConstraint{expected: uint(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "UintLessIntWithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualIntWithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterIntWithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "UintLessIntWithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(5)).
			ExpectResult(true)

		NewSubtest(t, "UintLessIntWithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(-5)).
			ExpectResult(true)
	}

	NewSubtest(t, "UintLessInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int8(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "UintLessInt8WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(5)).
			ExpectResult(true)

		NewSubtest(t, "UintLessInt8WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(-5)).
			ExpectResult(true)
	}

	NewSubtest(t, "UintLessInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int16(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "UintLessInt16WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(5)).
			ExpectResult(true)

		NewSubtest(t, "UintLessInt16WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(-5)).
			ExpectResult(true)
	}

	NewSubtest(t, "UintLessInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int32(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "UintLessInt32WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(5)).
			ExpectResult(true)

		NewSubtest(t, "UintLessInt32WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(-5)).
			ExpectResult(true)
	}

	NewSubtest(t, "UintLessInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int64(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewSubtest(t, "UintLessInt64WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(5)).
			ExpectResult(true)

		NewSubtest(t, "UintLessInt64WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(-5)).
			ExpectResult(true)
	}

	NewSubtest(t, "UintLessUint").
		Call((&LessConstraint{expected: uint(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualUint").
		Call((&LessConstraint{expected: uint(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterUint").
		Call((&LessConstraint{expected: uint(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "UintLessUint8").
		Call((&LessConstraint{expected: uint(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualUint8").
		Call((&LessConstraint{expected: uint(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterUint8").
		Call((&LessConstraint{expected: uint(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "UintLessUint16").
		Call((&LessConstraint{expected: uint(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualUint16").
		Call((&LessConstraint{expected: uint(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterUint16").
		Call((&LessConstraint{expected: uint(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "UintLessUint32").
		Call((&LessConstraint{expected: uint(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualUint32").
		Call((&LessConstraint{expected: uint(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterUint32").
		Call((&LessConstraint{expected: uint(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "UintLessUint64").
		Call((&LessConstraint{expected: uint(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualUint64").
		Call((&LessConstraint{expected: uint(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterUint64").
		Call((&LessConstraint{expected: uint(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "UintAndUintPtr").
		Call((&LessConstraint{expected: uint(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "UintLessFloat32").
		Call((&LessConstraint{expected: uint(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualFloat32").
		Call((&LessConstraint{expected: uint(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterFloat32").
		Call((&LessConstraint{expected: uint(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "UintLessFloat64").
		Call((&LessConstraint{expected: uint(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "UintEqualFloat64").
		Call((&LessConstraint{expected: uint(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "UintGreaterFloat64").
		Call((&LessConstraint{expected: uint(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "UintAndComplex64").
		Call((&LessConstraint{expected: uint(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "UintAndComplex128").
		Call((&LessConstraint{expected: uint(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "UintAndArray").
		Call((&LessConstraint{expected: uint(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "UintAndChan").
		Call((&LessConstraint{expected: uint(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "UintAndFunc").
		Call((&LessConstraint{expected: uint(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "UintAndInterface").
		Call((&LessConstraint{expected: uint(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "UintAndMap").
		Call((&LessConstraint{expected: uint(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "UintAndPtr").
		Call((&LessConstraint{expected: uint(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "UintAndSlice").
		Call((&LessConstraint{expected: uint(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "UintAndString").
		Call((&LessConstraint{expected: uint(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "UintAndStruct").
		Call((&LessConstraint{expected: uint(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "UintAndUnsafePointer").
		Call((&LessConstraint{expected: uint(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint8(t *testing.T) {
	NewSubtest(t, "Uint8AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Uint8AndBool").
		Call((&LessConstraint{expected: uint8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Uint8LessInt").
		Call((&LessConstraint{expected: uint8(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualInt").
		Call((&LessConstraint{expected: uint8(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterInt").
		Call((&LessConstraint{expected: uint8(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessInt8").
		Call((&LessConstraint{expected: uint8(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualInt8").
		Call((&LessConstraint{expected: uint8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterInt8").
		Call((&LessConstraint{expected: uint8(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessInt16").
		Call((&LessConstraint{expected: uint8(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualInt16").
		Call((&LessConstraint{expected: uint8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterInt16").
		Call((&LessConstraint{expected: uint8(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessInt32").
		Call((&LessConstraint{expected: uint8(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualInt32").
		Call((&LessConstraint{expected: uint8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterInt32").
		Call((&LessConstraint{expected: uint8(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessInt64").
		Call((&LessConstraint{expected: uint8(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualInt64").
		Call((&LessConstraint{expected: uint8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterInt64").
		Call((&LessConstraint{expected: uint8(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessUint").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualUint").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterUint").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessUint8").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualUint8").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterUint8").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessUint16").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualUint16").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterUint16").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessUint32").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualUint32").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterUint32").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessUint64").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualUint64").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterUint64").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8AndUintPtr").
		Call((&LessConstraint{expected: uint8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Uint8LessFloat32").
		Call((&LessConstraint{expected: uint8(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualFloat32").
		Call((&LessConstraint{expected: uint8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterFloat32").
		Call((&LessConstraint{expected: uint8(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8LessFloat64").
		Call((&LessConstraint{expected: uint8(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint8EqualFloat64").
		Call((&LessConstraint{expected: uint8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint8GreaterFloat64").
		Call((&LessConstraint{expected: uint8(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint8AndComplex64").
		Call((&LessConstraint{expected: uint8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Uint8AndComplex128").
		Call((&LessConstraint{expected: uint8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Uint8AndArray").
		Call((&LessConstraint{expected: uint8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Uint8AndChan").
		Call((&LessConstraint{expected: uint8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Uint8AndFunc").
		Call((&LessConstraint{expected: uint8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Uint8AndInterface").
		Call((&LessConstraint{expected: uint8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Uint8AndMap").
		Call((&LessConstraint{expected: uint8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Uint8AndPtr").
		Call((&LessConstraint{expected: uint8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Uint8AndSlice").
		Call((&LessConstraint{expected: uint8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Uint8AndString").
		Call((&LessConstraint{expected: uint8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Uint8AndStruct").
		Call((&LessConstraint{expected: uint8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Uint8AndUnsafePointer").
		Call((&LessConstraint{expected: uint8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint16(t *testing.T) {
	NewSubtest(t, "Uint16AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Uint16AndBool").
		Call((&LessConstraint{expected: uint16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Uint16LessInt").
		Call((&LessConstraint{expected: uint16(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualInt").
		Call((&LessConstraint{expected: uint16(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterInt").
		Call((&LessConstraint{expected: uint16(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessInt8").
		Call((&LessConstraint{expected: uint16(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualInt8").
		Call((&LessConstraint{expected: uint16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterInt8").
		Call((&LessConstraint{expected: uint16(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessInt16").
		Call((&LessConstraint{expected: uint16(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualInt16").
		Call((&LessConstraint{expected: uint16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterInt16").
		Call((&LessConstraint{expected: uint16(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessInt32").
		Call((&LessConstraint{expected: uint16(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualInt32").
		Call((&LessConstraint{expected: uint16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterInt32").
		Call((&LessConstraint{expected: uint16(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessInt64").
		Call((&LessConstraint{expected: uint16(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualInt64").
		Call((&LessConstraint{expected: uint16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterInt64").
		Call((&LessConstraint{expected: uint16(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessUint").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualUint").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterUint").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessUint8").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualUint8").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterUint8").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessUint16").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualUint16").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterUint16").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessUint32").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualUint32").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterUint32").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessUint64").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualUint64").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterUint64").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16AndUintPtr").
		Call((&LessConstraint{expected: uint16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Uint16LessFloat32").
		Call((&LessConstraint{expected: uint16(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualFloat32").
		Call((&LessConstraint{expected: uint16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterFloat32").
		Call((&LessConstraint{expected: uint16(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16LessFloat64").
		Call((&LessConstraint{expected: uint16(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint16EqualFloat64").
		Call((&LessConstraint{expected: uint16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint16GreaterFloat64").
		Call((&LessConstraint{expected: uint16(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint16AndComplex64").
		Call((&LessConstraint{expected: uint16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Uint16AndComplex128").
		Call((&LessConstraint{expected: uint16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Uint16AndArray").
		Call((&LessConstraint{expected: uint16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Uint16AndChan").
		Call((&LessConstraint{expected: uint16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Uint16AndFunc").
		Call((&LessConstraint{expected: uint16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Uint16AndInterface").
		Call((&LessConstraint{expected: uint16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Uint16AndMap").
		Call((&LessConstraint{expected: uint16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Uint16AndPtr").
		Call((&LessConstraint{expected: uint16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Uint16AndSlice").
		Call((&LessConstraint{expected: uint16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Uint16AndString").
		Call((&LessConstraint{expected: uint16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Uint16AndStruct").
		Call((&LessConstraint{expected: uint16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Uint16AndUnsafePointer").
		Call((&LessConstraint{expected: uint16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint32(t *testing.T) {
	NewSubtest(t, "Uint32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Uint32AndBool").
		Call((&LessConstraint{expected: uint32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Uint32LessInt").
		Call((&LessConstraint{expected: uint32(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualInt").
		Call((&LessConstraint{expected: uint32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterInt").
		Call((&LessConstraint{expected: uint32(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessInt8").
		Call((&LessConstraint{expected: uint32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualInt8").
		Call((&LessConstraint{expected: uint32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterInt8").
		Call((&LessConstraint{expected: uint32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessInt16").
		Call((&LessConstraint{expected: uint32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualInt16").
		Call((&LessConstraint{expected: uint32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterInt16").
		Call((&LessConstraint{expected: uint32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessInt32").
		Call((&LessConstraint{expected: uint32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualInt32").
		Call((&LessConstraint{expected: uint32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterInt32").
		Call((&LessConstraint{expected: uint32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessInt64").
		Call((&LessConstraint{expected: uint32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualInt64").
		Call((&LessConstraint{expected: uint32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterInt64").
		Call((&LessConstraint{expected: uint32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessUint").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualUint").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterUint").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessUint8").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualUint8").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessUint16").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualUint16").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterUint16").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessUint32").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualUint32").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterUint32").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessUint64").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualUint64").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterUint64").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32AndUintPtr").
		Call((&LessConstraint{expected: uint32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Uint32LessFloat32").
		Call((&LessConstraint{expected: uint32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualFloat32").
		Call((&LessConstraint{expected: uint32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterFloat32").
		Call((&LessConstraint{expected: uint32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32LessFloat64").
		Call((&LessConstraint{expected: uint32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint32EqualFloat64").
		Call((&LessConstraint{expected: uint32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint32GreaterFloat64").
		Call((&LessConstraint{expected: uint32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint32AndComplex64").
		Call((&LessConstraint{expected: uint32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Uint32AndComplex128").
		Call((&LessConstraint{expected: uint32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Uint32AndArray").
		Call((&LessConstraint{expected: uint32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Uint32AndChan").
		Call((&LessConstraint{expected: uint32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Uint32AndFunc").
		Call((&LessConstraint{expected: uint32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Uint32AndInterface").
		Call((&LessConstraint{expected: uint32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Uint32AndMap").
		Call((&LessConstraint{expected: uint32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Uint32AndPtr").
		Call((&LessConstraint{expected: uint32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Uint32AndSlice").
		Call((&LessConstraint{expected: uint32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Uint32AndString").
		Call((&LessConstraint{expected: uint32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Uint32AndStruct").
		Call((&LessConstraint{expected: uint32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Uint32AndUnsafePointer").
		Call((&LessConstraint{expected: uint32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint64(t *testing.T) {
	NewSubtest(t, "Uint64ndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Uint64AndBool").
		Call((&LessConstraint{expected: uint64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Uint64LessIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessIntWithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessIntWithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(-5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessInt8WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt8WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(-5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessInt16WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt16WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(-5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessInt32WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt32WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(-5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessInt64WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessInt64WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(-5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64LessUint").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualUint").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterUint").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessUint8").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualUint8").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterUint8").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessUint16").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualUint16").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterUint16").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessUint32").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualUint32").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterUint32").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessUint64").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualUint64").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterUint64").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64AndUintPtr").
		Call((&LessConstraint{expected: uint64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Uint64LessFloat32").
		Call((&LessConstraint{expected: uint64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualFloat32").
		Call((&LessConstraint{expected: uint64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterFloat32").
		Call((&LessConstraint{expected: uint64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64LessFloat64").
		Call((&LessConstraint{expected: uint64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Uint64EqualFloat64").
		Call((&LessConstraint{expected: uint64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Uint64GreaterFloat64").
		Call((&LessConstraint{expected: uint64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Uint64AndComplex64").
		Call((&LessConstraint{expected: uint64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Uint64AndComplex128").
		Call((&LessConstraint{expected: uint64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Uint64AndArray").
		Call((&LessConstraint{expected: uint64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Uint64AndChan").
		Call((&LessConstraint{expected: uint64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Uint64AndFunc").
		Call((&LessConstraint{expected: uint64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Uint64AndInterface").
		Call((&LessConstraint{expected: uint64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Uint64AndMap").
		Call((&LessConstraint{expected: uint64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Uint64AndPtr").
		Call((&LessConstraint{expected: uint64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Uint64AndSlice").
		Call((&LessConstraint{expected: uint64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Uint64AndString").
		Call((&LessConstraint{expected: uint64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Uint64AndStruct").
		Call((&LessConstraint{expected: uint64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Uint64AndUnsafePointer").
		Call((&LessConstraint{expected: uint64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithFloat32(t *testing.T) {
	NewSubtest(t, "Float32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: float32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Float32AndBool").
		Call((&LessConstraint{expected: float32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Float32LessInt").
		Call((&LessConstraint{expected: float32(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualInt").
		Call((&LessConstraint{expected: float32(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterInt").
		Call((&LessConstraint{expected: float32(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessInt8").
		Call((&LessConstraint{expected: float32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualInt8").
		Call((&LessConstraint{expected: float32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterInt8").
		Call((&LessConstraint{expected: float32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessInt16").
		Call((&LessConstraint{expected: float32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualInt16").
		Call((&LessConstraint{expected: float32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterInt16").
		Call((&LessConstraint{expected: float32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessInt32").
		Call((&LessConstraint{expected: float32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualInt32").
		Call((&LessConstraint{expected: float32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterInt32").
		Call((&LessConstraint{expected: float32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessInt64").
		Call((&LessConstraint{expected: float32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualInt64").
		Call((&LessConstraint{expected: float32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterInt64").
		Call((&LessConstraint{expected: float32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessUint").
		Call((&LessConstraint{expected: float32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualUint").
		Call((&LessConstraint{expected: float32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterUint").
		Call((&LessConstraint{expected: float32(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessUint8").
		Call((&LessConstraint{expected: float32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualUint8").
		Call((&LessConstraint{expected: float32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterUint8").
		Call((&LessConstraint{expected: float32(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessUint16").
		Call((&LessConstraint{expected: float32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualUint16").
		Call((&LessConstraint{expected: float32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterUint16").
		Call((&LessConstraint{expected: float32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessUint32").
		Call((&LessConstraint{expected: float32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualUint32").
		Call((&LessConstraint{expected: float32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterUint32").
		Call((&LessConstraint{expected: float32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessUint64").
		Call((&LessConstraint{expected: float32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualUint64").
		Call((&LessConstraint{expected: float32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterUint64").
		Call((&LessConstraint{expected: float32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32AndUintPtr").
		Call((&LessConstraint{expected: float32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Float32LessFloat32").
		Call((&LessConstraint{expected: float32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualFloat32").
		Call((&LessConstraint{expected: float32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterFloat32").
		Call((&LessConstraint{expected: float32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32LessFloat64").
		Call((&LessConstraint{expected: float32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float32EqualFloat64").
		Call((&LessConstraint{expected: float32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float32GreaterFloat64").
		Call((&LessConstraint{expected: float32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float32AndComplex64").
		Call((&LessConstraint{expected: float32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Float32AndComplex128").
		Call((&LessConstraint{expected: float32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Float32AndArray").
		Call((&LessConstraint{expected: float32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Float32AndChan").
		Call((&LessConstraint{expected: float32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Float32AndFunc").
		Call((&LessConstraint{expected: float32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Float32AndInterface").
		Call((&LessConstraint{expected: float32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Float32AndMap").
		Call((&LessConstraint{expected: float32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Float32AndPtr").
		Call((&LessConstraint{expected: float32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Float32AndSlice").
		Call((&LessConstraint{expected: float32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Float32AndString").
		Call((&LessConstraint{expected: float32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Float32AndStruct").
		Call((&LessConstraint{expected: float32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Float32AndUnsafePointer").
		Call((&LessConstraint{expected: float32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithFloat64(t *testing.T) {
	NewSubtest(t, "Float64AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: float64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewSubtest(t, "Float64AndBool").
		Call((&LessConstraint{expected: float64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewSubtest(t, "Float64LessInt").
		Call((&LessConstraint{expected: float64(10)}).Check, int(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualInt").
		Call((&LessConstraint{expected: float64(5)}).Check, int(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterInt").
		Call((&LessConstraint{expected: float64(5)}).Check, int(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessInt8").
		Call((&LessConstraint{expected: float64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualInt8").
		Call((&LessConstraint{expected: float64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterInt8").
		Call((&LessConstraint{expected: float64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessInt16").
		Call((&LessConstraint{expected: float64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualInt16").
		Call((&LessConstraint{expected: float64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterInt16").
		Call((&LessConstraint{expected: float64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessInt32").
		Call((&LessConstraint{expected: float64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualInt32").
		Call((&LessConstraint{expected: float64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterInt32").
		Call((&LessConstraint{expected: float64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessInt64").
		Call((&LessConstraint{expected: float64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualInt64").
		Call((&LessConstraint{expected: float64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterInt64").
		Call((&LessConstraint{expected: float64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessUint").
		Call((&LessConstraint{expected: float64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualUint").
		Call((&LessConstraint{expected: float64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterUint").
		Call((&LessConstraint{expected: float64(5)}).Check, uint(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessUint8").
		Call((&LessConstraint{expected: float64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualUint8").
		Call((&LessConstraint{expected: float64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterUint8").
		Call((&LessConstraint{expected: float64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessUint16").
		Call((&LessConstraint{expected: float64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualUint16").
		Call((&LessConstraint{expected: float64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterUint16").
		Call((&LessConstraint{expected: float64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessUint32").
		Call((&LessConstraint{expected: float64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualUint32").
		Call((&LessConstraint{expected: float64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterUint32").
		Call((&LessConstraint{expected: float64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessUint64").
		Call((&LessConstraint{expected: float64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualUint64").
		Call((&LessConstraint{expected: float64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterUint64").
		Call((&LessConstraint{expected: float64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64AndUintPtr").
		Call((&LessConstraint{expected: float64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewSubtest(t, "Float64LessFloat32").
		Call((&LessConstraint{expected: float64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualFloat32").
		Call((&LessConstraint{expected: float64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterFloat32").
		Call((&LessConstraint{expected: float64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64LessFloat64").
		Call((&LessConstraint{expected: float64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewSubtest(t, "Float64EqualFloat64").
		Call((&LessConstraint{expected: float64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewSubtest(t, "Float64GreaterFloat64").
		Call((&LessConstraint{expected: float64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewSubtest(t, "Float64AndComplex64").
		Call((&LessConstraint{expected: float64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewSubtest(t, "Float64AndComplex128").
		Call((&LessConstraint{expected: float64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewSubtest(t, "Float64AndArray").
		Call((&LessConstraint{expected: float64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewSubtest(t, "Float64AndChan").
		Call((&LessConstraint{expected: float64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewSubtest(t, "Float64AndFunc").
		Call((&LessConstraint{expected: float64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewSubtest(t, "Float64AndInterface").
		Call((&LessConstraint{expected: float64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewSubtest(t, "Float64AndMap").
		Call((&LessConstraint{expected: float64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewSubtest(t, "Float64AndPtr").
		Call((&LessConstraint{expected: float64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewSubtest(t, "Float64AndSlice").
		Call((&LessConstraint{expected: float64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewSubtest(t, "Float64AndString").
		Call((&LessConstraint{expected: float64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewSubtest(t, "Float64AndStruct").
		Call((&LessConstraint{expected: float64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewSubtest(t, "Float64AndUnsafePointer").
		Call((&LessConstraint{expected: float64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

func TestLessConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&LessConstraint{expected: 55}).String).
		ExpectResult(fmt.Sprintf("be less than %v", 55))
}

func TestLessConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&LessConstraint{expected: 55}).Details, 10).
		ExpectResult("")
}
