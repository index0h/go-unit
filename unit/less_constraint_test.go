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
	NewDeclarative(t, "Invalid").
		Call(
			func() {
				NewLessConstraint(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", nil))

	NewDeclarative(t, "Bool").
		Call(NewLessConstraint, true).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", true))

	NewDeclarative(t, "Int").
		Call(NewLessConstraint, int(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int(5)}})

	NewDeclarative(t, "Int8").
		Call(NewLessConstraint, int8(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int8(5)}})

	NewDeclarative(t, "Int16").
		Call(NewLessConstraint, int16(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int16(5)}})

	NewDeclarative(t, "Int32").
		Call(NewLessConstraint, int32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int32(5)}})

	NewDeclarative(t, "Int64").
		Call(NewLessConstraint, int64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: int64(5)}})

	NewDeclarative(t, "Uint").
		Call(NewLessConstraint, uint(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint(5)}})

	NewDeclarative(t, "Uint8").
		Call(NewLessConstraint, uint8(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint8(5)}})

	NewDeclarative(t, "Uint16").
		Call(NewLessConstraint, uint16(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint16(5)}})

	NewDeclarative(t, "Uint32").
		Call(NewLessConstraint, uint32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint32(5)}})

	NewDeclarative(t, "Uint64").
		Call(NewLessConstraint, uint64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: uint64(5)}})

	NewDeclarative(t, "UintPtr").
		Call(NewLessConstraint, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", uintptr(5)))

	NewDeclarative(t, "Float32").
		Call(NewLessConstraint, float32(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: float32(5)}})

	NewDeclarative(t, "Float64").
		Call(NewLessConstraint, float64(5)).
		ExpectResult(ConstraintAsValue{Value: &LessConstraint{expected: float64(5)}})

	NewDeclarative(t, "Complex64").
		Call(NewLessConstraint, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex64(5)))

	NewDeclarative(t, "Complex128").
		Call(NewLessConstraint, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", complex128(5)))

	NewDeclarative(t, "Array").
		Call(NewLessConstraint, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", [1]int{5}))

	NewDeclarative(t, "Chan").
		Call(NewLessConstraint, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", make(chan int)))

	NewDeclarative(t, "Func").
		Call(NewLessConstraint, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", func() {}))

	NewDeclarative(t, "Interface").
		Call(NewLessConstraint, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", (*interface{})(nil)))

	NewDeclarative(t, "Map").
		Call(NewLessConstraint, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", map[int]int{1: 1}))

	NewDeclarative(t, "Ptr").
		Call(NewLessConstraint, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", new(int)))

	NewDeclarative(t, "Slice").
		Call(NewLessConstraint, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", []int{5}))

	NewDeclarative(t, "String").
		Call(NewLessConstraint, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", "data"))

	NewDeclarative(t, "Struct").
		Call(NewLessConstraint, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", struct{}{}))

	NewDeclarative(t, "UnsafePointer").
		Call(NewLessConstraint, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("expectedKinds", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt(t *testing.T) {
	NewDeclarative(t, "IntAndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "IntAndBool").
		Call((&LessConstraint{expected: int(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "IntLessInt").
		Call((&LessConstraint{expected: int(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualInt").
		Call((&LessConstraint{expected: int(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterInt").
		Call((&LessConstraint{expected: int(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessInt8").
		Call((&LessConstraint{expected: int(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualInt8").
		Call((&LessConstraint{expected: int(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterInt8").
		Call((&LessConstraint{expected: int(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessInt16").
		Call((&LessConstraint{expected: int(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualInt16").
		Call((&LessConstraint{expected: int(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterInt16").
		Call((&LessConstraint{expected: int(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessInt32").
		Call((&LessConstraint{expected: int(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualInt32").
		Call((&LessConstraint{expected: int(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterInt32").
		Call((&LessConstraint{expected: int(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessInt64").
		Call((&LessConstraint{expected: int(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualInt64").
		Call((&LessConstraint{expected: int(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterInt64").
		Call((&LessConstraint{expected: int(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessUintWithIntComparison").
		Call((&LessConstraint{expected: int(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualUintWithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUintAndIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "IntGreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewDeclarative(t, "IntGreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewDeclarative(t, "IntLessUint8").
		Call((&LessConstraint{expected: int(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualUint8").
		Call((&LessConstraint{expected: int(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint8").
		Call((&LessConstraint{expected: int(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessUint16").
		Call((&LessConstraint{expected: int(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualUint16").
		Call((&LessConstraint{expected: int(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint16").
		Call((&LessConstraint{expected: int(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessUint32").
		Call((&LessConstraint{expected: int(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualUint32").
		Call((&LessConstraint{expected: int(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint32").
		Call((&LessConstraint{expected: int(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessUint64WithIntComparison").
		Call((&LessConstraint{expected: int(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "IntAndUintPtr").
		Call((&LessConstraint{expected: int(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "IntLessFloat32").
		Call((&LessConstraint{expected: int(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualFloat32").
		Call((&LessConstraint{expected: int(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterFloat32").
		Call((&LessConstraint{expected: int(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntLessFloat64").
		Call((&LessConstraint{expected: int(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "IntEqualFloat64").
		Call((&LessConstraint{expected: int(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "IntGreaterFloat64").
		Call((&LessConstraint{expected: int(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "IntAndComplex64").
		Call((&LessConstraint{expected: int(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "IntAndComplex128").
		Call((&LessConstraint{expected: int(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "IntAndArray").
		Call((&LessConstraint{expected: int(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "IntAndChan").
		Call((&LessConstraint{expected: int(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "IntAndFunc").
		Call((&LessConstraint{expected: int(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "IntAndInterface").
		Call((&LessConstraint{expected: int(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "IntAndMap").
		Call((&LessConstraint{expected: int(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "IntAndPtr").
		Call((&LessConstraint{expected: int(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "IntAndSlice").
		Call((&LessConstraint{expected: int(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "IntAndString").
		Call((&LessConstraint{expected: int(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "IntAndStruct").
		Call((&LessConstraint{expected: int(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "IntAndUnsafePointer").
		Call((&LessConstraint{expected: int(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt8(t *testing.T) {
	NewDeclarative(t, "Int8AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Int8AndBool").
		Call((&LessConstraint{expected: int8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Int8LessInt").
		Call((&LessConstraint{expected: int8(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualInt").
		Call((&LessConstraint{expected: int8(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterInt").
		Call((&LessConstraint{expected: int8(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessInt8").
		Call((&LessConstraint{expected: int8(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualInt8").
		Call((&LessConstraint{expected: int8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterInt8").
		Call((&LessConstraint{expected: int8(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessInt16").
		Call((&LessConstraint{expected: int8(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualInt16").
		Call((&LessConstraint{expected: int8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterInt16").
		Call((&LessConstraint{expected: int8(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessInt32").
		Call((&LessConstraint{expected: int8(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualInt32").
		Call((&LessConstraint{expected: int8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterInt32").
		Call((&LessConstraint{expected: int8(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessInt64").
		Call((&LessConstraint{expected: int8(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualInt64").
		Call((&LessConstraint{expected: int8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterInt64").
		Call((&LessConstraint{expected: int8(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessUintWihtIntComparison").
		Call((&LessConstraint{expected: int8(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUintIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "Int8GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int8(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewDeclarative(t, "Int8GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int8(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewDeclarative(t, "Int8LessUint8").
		Call((&LessConstraint{expected: int8(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualUint8").
		Call((&LessConstraint{expected: int8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint8").
		Call((&LessConstraint{expected: int8(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessUint16").
		Call((&LessConstraint{expected: int8(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualUint16").
		Call((&LessConstraint{expected: int8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint16").
		Call((&LessConstraint{expected: int8(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessUint32").
		Call((&LessConstraint{expected: int8(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualUint32").
		Call((&LessConstraint{expected: int8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint32").
		Call((&LessConstraint{expected: int8(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int8(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int8(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int8AndUintPtr").
		Call((&LessConstraint{expected: int8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Int8LessFloat32").
		Call((&LessConstraint{expected: int8(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualFloat32").
		Call((&LessConstraint{expected: int8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterFloat32").
		Call((&LessConstraint{expected: int8(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8LessFloat64").
		Call((&LessConstraint{expected: int8(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int8EqualFloat64").
		Call((&LessConstraint{expected: int8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int8GreaterFloat64").
		Call((&LessConstraint{expected: int8(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int8AndComplex64").
		Call((&LessConstraint{expected: int8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Int8AndComplex128").
		Call((&LessConstraint{expected: int8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Int8AndArray").
		Call((&LessConstraint{expected: int8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Int8AndChan").
		Call((&LessConstraint{expected: int8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Int8AndFunc").
		Call((&LessConstraint{expected: int8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Int8AndInterface").
		Call((&LessConstraint{expected: int8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Int8AndMap").
		Call((&LessConstraint{expected: int8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Int8AndPtr").
		Call((&LessConstraint{expected: int8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Int8AndSlice").
		Call((&LessConstraint{expected: int8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Int8AndString").
		Call((&LessConstraint{expected: int8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Int8AndStruct").
		Call((&LessConstraint{expected: int8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Int8AndUnsafePointer").
		Call((&LessConstraint{expected: int8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt16(t *testing.T) {
	NewDeclarative(t, "Int16AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Int16AndBool").
		Call((&LessConstraint{expected: int16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Int16LessInt").
		Call((&LessConstraint{expected: int16(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualInt").
		Call((&LessConstraint{expected: int16(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterInt").
		Call((&LessConstraint{expected: int16(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessInt8").
		Call((&LessConstraint{expected: int16(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualInt8").
		Call((&LessConstraint{expected: int16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterInt8").
		Call((&LessConstraint{expected: int16(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessInt16").
		Call((&LessConstraint{expected: int16(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualInt16").
		Call((&LessConstraint{expected: int16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterInt16").
		Call((&LessConstraint{expected: int16(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessInt32").
		Call((&LessConstraint{expected: int16(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualInt32").
		Call((&LessConstraint{expected: int16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterInt32").
		Call((&LessConstraint{expected: int16(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessInt64").
		Call((&LessConstraint{expected: int16(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualInt64").
		Call((&LessConstraint{expected: int16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterInt64").
		Call((&LessConstraint{expected: int16(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessUintWithIntComparison").
		Call((&LessConstraint{expected: int16(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "Int16GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int16(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewDeclarative(t, "Int16GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int16(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewDeclarative(t, "Int16LessUint8").
		Call((&LessConstraint{expected: int16(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualUint8").
		Call((&LessConstraint{expected: int16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint8").
		Call((&LessConstraint{expected: int16(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessUint16").
		Call((&LessConstraint{expected: int16(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualUint16").
		Call((&LessConstraint{expected: int16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint16").
		Call((&LessConstraint{expected: int16(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessUint32").
		Call((&LessConstraint{expected: int16(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualUint32").
		Call((&LessConstraint{expected: int16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint32").
		Call((&LessConstraint{expected: int16(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int16(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int16(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int16AndUintPtr").
		Call((&LessConstraint{expected: int16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Int16LessFloat32").
		Call((&LessConstraint{expected: int16(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualFloat32").
		Call((&LessConstraint{expected: int16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterFloat32").
		Call((&LessConstraint{expected: int16(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16LessFloat64").
		Call((&LessConstraint{expected: int16(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int16EqualFloat64").
		Call((&LessConstraint{expected: int16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int16GreaterFloat64").
		Call((&LessConstraint{expected: int16(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int16AndComplex64").
		Call((&LessConstraint{expected: int16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Int16AndComplex128").
		Call((&LessConstraint{expected: int16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Int16AndArray").
		Call((&LessConstraint{expected: int16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Int16AndChan").
		Call((&LessConstraint{expected: int16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Int16AndFunc").
		Call((&LessConstraint{expected: int16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Int16AndInterface").
		Call((&LessConstraint{expected: int16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Int16AndMap").
		Call((&LessConstraint{expected: int16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Int16AndPtr").
		Call((&LessConstraint{expected: int16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Int16AndSlice").
		Call((&LessConstraint{expected: int16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Int16AndString").
		Call((&LessConstraint{expected: int16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Int16AndStruct").
		Call((&LessConstraint{expected: int16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Int16AndUnsafePointer").
		Call((&LessConstraint{expected: int16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt32(t *testing.T) {
	NewDeclarative(t, "Int32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Int32AndBool").
		Call((&LessConstraint{expected: int32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Int32LessInt").
		Call((&LessConstraint{expected: int32(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualInt").
		Call((&LessConstraint{expected: int32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterInt").
		Call((&LessConstraint{expected: int32(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessInt8").
		Call((&LessConstraint{expected: int32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualInt8").
		Call((&LessConstraint{expected: int32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterInt8").
		Call((&LessConstraint{expected: int32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessInt16").
		Call((&LessConstraint{expected: int32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualInt16").
		Call((&LessConstraint{expected: int32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterInt16").
		Call((&LessConstraint{expected: int32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessInt32").
		Call((&LessConstraint{expected: int32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualInt32").
		Call((&LessConstraint{expected: int32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterInt32").
		Call((&LessConstraint{expected: int32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessInt64").
		Call((&LessConstraint{expected: int32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualInt64").
		Call((&LessConstraint{expected: int32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterInt64").
		Call((&LessConstraint{expected: int32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessUintWithIntComparison").
		Call((&LessConstraint{expected: int32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "Int32GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int32(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewDeclarative(t, "Int32GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int32(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewDeclarative(t, "Int32LessUint8").
		Call((&LessConstraint{expected: int32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualUint8").
		Call((&LessConstraint{expected: int32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint8").
		Call((&LessConstraint{expected: int32(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessUint16").
		Call((&LessConstraint{expected: int32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualUint16").
		Call((&LessConstraint{expected: int32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint16").
		Call((&LessConstraint{expected: int32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessUint32").
		Call((&LessConstraint{expected: int32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualUint32").
		Call((&LessConstraint{expected: int32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint32").
		Call((&LessConstraint{expected: int32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint64WihtIntComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int32(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int32(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int32AndUintPtr").
		Call((&LessConstraint{expected: int32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Int32LessFloat32").
		Call((&LessConstraint{expected: int32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualFloat32").
		Call((&LessConstraint{expected: int32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterFloat32").
		Call((&LessConstraint{expected: int32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32LessFloat64").
		Call((&LessConstraint{expected: int32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int32EqualFloat64").
		Call((&LessConstraint{expected: int32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int32GreaterFloat64").
		Call((&LessConstraint{expected: int32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int32AndComplex64").
		Call((&LessConstraint{expected: int32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Int32AndComplex128").
		Call((&LessConstraint{expected: int32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Int32AndArray").
		Call((&LessConstraint{expected: int32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Int32AndChan").
		Call((&LessConstraint{expected: int32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Int32AndFunc").
		Call((&LessConstraint{expected: int32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Int32AndInterface").
		Call((&LessConstraint{expected: int32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Int32AndMap").
		Call((&LessConstraint{expected: int32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Int32AndPtr").
		Call((&LessConstraint{expected: int32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Int32AndSlice").
		Call((&LessConstraint{expected: int32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Int32AndString").
		Call((&LessConstraint{expected: int32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Int32AndStruct").
		Call((&LessConstraint{expected: int32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Int32AndUnsafePointer").
		Call((&LessConstraint{expected: int32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithInt64(t *testing.T) {
	NewDeclarative(t, "Int64AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: int64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Int64AndBool").
		Call((&LessConstraint{expected: int64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Int64LessInt").
		Call((&LessConstraint{expected: int64(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualInt").
		Call((&LessConstraint{expected: int64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterInt").
		Call((&LessConstraint{expected: int64(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessInt8").
		Call((&LessConstraint{expected: int64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualInt8").
		Call((&LessConstraint{expected: int64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterInt8").
		Call((&LessConstraint{expected: int64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessInt16").
		Call((&LessConstraint{expected: int64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualInt16").
		Call((&LessConstraint{expected: int64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterInt16").
		Call((&LessConstraint{expected: int64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessInt32").
		Call((&LessConstraint{expected: int64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualInt32").
		Call((&LessConstraint{expected: int64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterInt32").
		Call((&LessConstraint{expected: int64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessInt64").
		Call((&LessConstraint{expected: int64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualInt64").
		Call((&LessConstraint{expected: int64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterInt64").
		Call((&LessConstraint{expected: int64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessUintWithIntComparison").
		Call((&LessConstraint{expected: int64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualUintWithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUintWithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {

		NewDeclarative(t, "Int64GreaterUintWithUintComparison").
			Call((&LessConstraint{expected: int64(5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)

		NewDeclarative(t, "Int64GreaterUintWithFloatComparison").
			Call((&LessConstraint{expected: int64(-5)}).Check, uint(math.MaxInt64)+1).
			ExpectResult(false)
	}

	NewDeclarative(t, "Int64LessUint8").
		Call((&LessConstraint{expected: int64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualUint8").
		Call((&LessConstraint{expected: int64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint8").
		Call((&LessConstraint{expected: int64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessUint16").
		Call((&LessConstraint{expected: int64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualUint16").
		Call((&LessConstraint{expected: int64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint16").
		Call((&LessConstraint{expected: int64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessUint32").
		Call((&LessConstraint{expected: int64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualUint32").
		Call((&LessConstraint{expected: int64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint32").
		Call((&LessConstraint{expected: int64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint64WithIntComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint64WithUintComparison").
		Call((&LessConstraint{expected: int64(5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterUint64WithFloatComparison").
		Call((&LessConstraint{expected: int64(-5)}).Check, uint64(math.MaxInt64)+1).
		ExpectResult(false)

	NewDeclarative(t, "Int64AndUintPtr").
		Call((&LessConstraint{expected: int64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Int64LessFloat32").
		Call((&LessConstraint{expected: int64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualFloat32").
		Call((&LessConstraint{expected: int64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterFloat32").
		Call((&LessConstraint{expected: int64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64LessFloat64").
		Call((&LessConstraint{expected: int64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Int64EqualFloat64").
		Call((&LessConstraint{expected: int64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Int64GreaterFloat64").
		Call((&LessConstraint{expected: int64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Int64AndComplex64").
		Call((&LessConstraint{expected: int64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Int64AndComplex128").
		Call((&LessConstraint{expected: int64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Int64AndArray").
		Call((&LessConstraint{expected: int64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Int64AndChan").
		Call((&LessConstraint{expected: int64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Int64AndFunc").
		Call((&LessConstraint{expected: int64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Int64AndInterface").
		Call((&LessConstraint{expected: int64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Int64AndMap").
		Call((&LessConstraint{expected: int64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Int64AndPtr").
		Call((&LessConstraint{expected: int64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Int64AndSlice").
		Call((&LessConstraint{expected: int64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Int64AndString").
		Call((&LessConstraint{expected: int64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Int64AndStruct").
		Call((&LessConstraint{expected: int64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Int64AndUnsafePointer").
		Call((&LessConstraint{expected: int64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint(t *testing.T) {
	NewDeclarative(t, "UintAndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "UintAndBool").
		Call((&LessConstraint{expected: uint(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "UintLessIntWithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualIntWithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterIntWithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "UintLessIntWithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(5)).
			ExpectResult(true)

		NewDeclarative(t, "UintLessIntWithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int(-5)).
			ExpectResult(true)
	}

	NewDeclarative(t, "UintLessInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterInt8WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int8(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "UintLessInt8WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(5)).
			ExpectResult(true)

		NewDeclarative(t, "UintLessInt8WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int8(-5)).
			ExpectResult(true)
	}

	NewDeclarative(t, "UintLessInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterInt16WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int16(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "UintLessInt16WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(5)).
			ExpectResult(true)

		NewDeclarative(t, "UintLessInt16WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int16(-5)).
			ExpectResult(true)
	}

	NewDeclarative(t, "UintLessInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterInt32WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int32(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "UintLessInt32WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(5)).
			ExpectResult(true)

		NewDeclarative(t, "UintLessInt32WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int32(-5)).
			ExpectResult(true)
	}

	NewDeclarative(t, "UintLessInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterInt64WithIntComparison").
		Call((&LessConstraint{expected: uint(5)}).Check, int64(10)).
		ExpectResult(false)

	if !reflect.ValueOf(uint(0)).OverflowUint(math.MaxUint64) {
		NewDeclarative(t, "UintLessInt64WithUintComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(5)).
			ExpectResult(true)

		NewDeclarative(t, "UintLessInt64WithFloatComparison").
			Call((&LessConstraint{expected: uint(math.MaxInt64) + 1}).Check, int64(-5)).
			ExpectResult(true)
	}

	NewDeclarative(t, "UintLessUint").
		Call((&LessConstraint{expected: uint(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualUint").
		Call((&LessConstraint{expected: uint(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterUint").
		Call((&LessConstraint{expected: uint(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintLessUint8").
		Call((&LessConstraint{expected: uint(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualUint8").
		Call((&LessConstraint{expected: uint(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterUint8").
		Call((&LessConstraint{expected: uint(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintLessUint16").
		Call((&LessConstraint{expected: uint(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualUint16").
		Call((&LessConstraint{expected: uint(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterUint16").
		Call((&LessConstraint{expected: uint(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintLessUint32").
		Call((&LessConstraint{expected: uint(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualUint32").
		Call((&LessConstraint{expected: uint(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterUint32").
		Call((&LessConstraint{expected: uint(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintLessUint64").
		Call((&LessConstraint{expected: uint(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualUint64").
		Call((&LessConstraint{expected: uint(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterUint64").
		Call((&LessConstraint{expected: uint(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintAndUintPtr").
		Call((&LessConstraint{expected: uint(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "UintLessFloat32").
		Call((&LessConstraint{expected: uint(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualFloat32").
		Call((&LessConstraint{expected: uint(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterFloat32").
		Call((&LessConstraint{expected: uint(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintLessFloat64").
		Call((&LessConstraint{expected: uint(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "UintEqualFloat64").
		Call((&LessConstraint{expected: uint(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "UintGreaterFloat64").
		Call((&LessConstraint{expected: uint(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "UintAndComplex64").
		Call((&LessConstraint{expected: uint(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "UintAndComplex128").
		Call((&LessConstraint{expected: uint(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "UintAndArray").
		Call((&LessConstraint{expected: uint(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "UintAndChan").
		Call((&LessConstraint{expected: uint(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "UintAndFunc").
		Call((&LessConstraint{expected: uint(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "UintAndInterface").
		Call((&LessConstraint{expected: uint(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "UintAndMap").
		Call((&LessConstraint{expected: uint(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "UintAndPtr").
		Call((&LessConstraint{expected: uint(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "UintAndSlice").
		Call((&LessConstraint{expected: uint(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "UintAndString").
		Call((&LessConstraint{expected: uint(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "UintAndStruct").
		Call((&LessConstraint{expected: uint(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "UintAndUnsafePointer").
		Call((&LessConstraint{expected: uint(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint8(t *testing.T) {
	NewDeclarative(t, "Uint8AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint8(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Uint8AndBool").
		Call((&LessConstraint{expected: uint8(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Uint8LessInt").
		Call((&LessConstraint{expected: uint8(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualInt").
		Call((&LessConstraint{expected: uint8(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterInt").
		Call((&LessConstraint{expected: uint8(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessInt8").
		Call((&LessConstraint{expected: uint8(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualInt8").
		Call((&LessConstraint{expected: uint8(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterInt8").
		Call((&LessConstraint{expected: uint8(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessInt16").
		Call((&LessConstraint{expected: uint8(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualInt16").
		Call((&LessConstraint{expected: uint8(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterInt16").
		Call((&LessConstraint{expected: uint8(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessInt32").
		Call((&LessConstraint{expected: uint8(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualInt32").
		Call((&LessConstraint{expected: uint8(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterInt32").
		Call((&LessConstraint{expected: uint8(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessInt64").
		Call((&LessConstraint{expected: uint8(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualInt64").
		Call((&LessConstraint{expected: uint8(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterInt64").
		Call((&LessConstraint{expected: uint8(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessUint").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualUint").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterUint").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessUint8").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualUint8").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterUint8").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessUint16").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualUint16").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterUint16").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessUint32").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualUint32").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterUint32").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessUint64").
		Call((&LessConstraint{expected: uint8(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualUint64").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterUint64").
		Call((&LessConstraint{expected: uint8(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8AndUintPtr").
		Call((&LessConstraint{expected: uint8(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Uint8LessFloat32").
		Call((&LessConstraint{expected: uint8(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualFloat32").
		Call((&LessConstraint{expected: uint8(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterFloat32").
		Call((&LessConstraint{expected: uint8(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8LessFloat64").
		Call((&LessConstraint{expected: uint8(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint8EqualFloat64").
		Call((&LessConstraint{expected: uint8(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8GreaterFloat64").
		Call((&LessConstraint{expected: uint8(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint8AndComplex64").
		Call((&LessConstraint{expected: uint8(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Uint8AndComplex128").
		Call((&LessConstraint{expected: uint8(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Uint8AndArray").
		Call((&LessConstraint{expected: uint8(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Uint8AndChan").
		Call((&LessConstraint{expected: uint8(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Uint8AndFunc").
		Call((&LessConstraint{expected: uint8(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Uint8AndInterface").
		Call((&LessConstraint{expected: uint8(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Uint8AndMap").
		Call((&LessConstraint{expected: uint8(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Uint8AndPtr").
		Call((&LessConstraint{expected: uint8(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Uint8AndSlice").
		Call((&LessConstraint{expected: uint8(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Uint8AndString").
		Call((&LessConstraint{expected: uint8(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Uint8AndStruct").
		Call((&LessConstraint{expected: uint8(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Uint8AndUnsafePointer").
		Call((&LessConstraint{expected: uint8(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint16(t *testing.T) {
	NewDeclarative(t, "Uint16AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint16(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Uint16AndBool").
		Call((&LessConstraint{expected: uint16(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Uint16LessInt").
		Call((&LessConstraint{expected: uint16(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualInt").
		Call((&LessConstraint{expected: uint16(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterInt").
		Call((&LessConstraint{expected: uint16(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessInt8").
		Call((&LessConstraint{expected: uint16(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualInt8").
		Call((&LessConstraint{expected: uint16(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterInt8").
		Call((&LessConstraint{expected: uint16(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessInt16").
		Call((&LessConstraint{expected: uint16(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualInt16").
		Call((&LessConstraint{expected: uint16(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterInt16").
		Call((&LessConstraint{expected: uint16(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessInt32").
		Call((&LessConstraint{expected: uint16(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualInt32").
		Call((&LessConstraint{expected: uint16(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterInt32").
		Call((&LessConstraint{expected: uint16(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessInt64").
		Call((&LessConstraint{expected: uint16(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualInt64").
		Call((&LessConstraint{expected: uint16(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterInt64").
		Call((&LessConstraint{expected: uint16(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessUint").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualUint").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterUint").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessUint8").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualUint8").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterUint8").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessUint16").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualUint16").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterUint16").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessUint32").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualUint32").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterUint32").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessUint64").
		Call((&LessConstraint{expected: uint16(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualUint64").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterUint64").
		Call((&LessConstraint{expected: uint16(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16AndUintPtr").
		Call((&LessConstraint{expected: uint16(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Uint16LessFloat32").
		Call((&LessConstraint{expected: uint16(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualFloat32").
		Call((&LessConstraint{expected: uint16(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterFloat32").
		Call((&LessConstraint{expected: uint16(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16LessFloat64").
		Call((&LessConstraint{expected: uint16(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint16EqualFloat64").
		Call((&LessConstraint{expected: uint16(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16GreaterFloat64").
		Call((&LessConstraint{expected: uint16(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint16AndComplex64").
		Call((&LessConstraint{expected: uint16(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Uint16AndComplex128").
		Call((&LessConstraint{expected: uint16(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Uint16AndArray").
		Call((&LessConstraint{expected: uint16(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Uint16AndChan").
		Call((&LessConstraint{expected: uint16(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Uint16AndFunc").
		Call((&LessConstraint{expected: uint16(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Uint16AndInterface").
		Call((&LessConstraint{expected: uint16(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Uint16AndMap").
		Call((&LessConstraint{expected: uint16(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Uint16AndPtr").
		Call((&LessConstraint{expected: uint16(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Uint16AndSlice").
		Call((&LessConstraint{expected: uint16(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Uint16AndString").
		Call((&LessConstraint{expected: uint16(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Uint16AndStruct").
		Call((&LessConstraint{expected: uint16(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Uint16AndUnsafePointer").
		Call((&LessConstraint{expected: uint16(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint32(t *testing.T) {
	NewDeclarative(t, "Uint32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Uint32AndBool").
		Call((&LessConstraint{expected: uint32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Uint32LessInt").
		Call((&LessConstraint{expected: uint32(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualInt").
		Call((&LessConstraint{expected: uint32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterInt").
		Call((&LessConstraint{expected: uint32(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessInt8").
		Call((&LessConstraint{expected: uint32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualInt8").
		Call((&LessConstraint{expected: uint32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterInt8").
		Call((&LessConstraint{expected: uint32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessInt16").
		Call((&LessConstraint{expected: uint32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualInt16").
		Call((&LessConstraint{expected: uint32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterInt16").
		Call((&LessConstraint{expected: uint32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessInt32").
		Call((&LessConstraint{expected: uint32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualInt32").
		Call((&LessConstraint{expected: uint32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterInt32").
		Call((&LessConstraint{expected: uint32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessInt64").
		Call((&LessConstraint{expected: uint32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualInt64").
		Call((&LessConstraint{expected: uint32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterInt64").
		Call((&LessConstraint{expected: uint32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessUint").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualUint").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterUint").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessUint8").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualUint8").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessUint16").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualUint16").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterUint16").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessUint32").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualUint32").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterUint32").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessUint64").
		Call((&LessConstraint{expected: uint32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualUint64").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterUint64").
		Call((&LessConstraint{expected: uint32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32AndUintPtr").
		Call((&LessConstraint{expected: uint32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Uint32LessFloat32").
		Call((&LessConstraint{expected: uint32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualFloat32").
		Call((&LessConstraint{expected: uint32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterFloat32").
		Call((&LessConstraint{expected: uint32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32LessFloat64").
		Call((&LessConstraint{expected: uint32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint32EqualFloat64").
		Call((&LessConstraint{expected: uint32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32GreaterFloat64").
		Call((&LessConstraint{expected: uint32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint32AndComplex64").
		Call((&LessConstraint{expected: uint32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Uint32AndComplex128").
		Call((&LessConstraint{expected: uint32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Uint32AndArray").
		Call((&LessConstraint{expected: uint32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Uint32AndChan").
		Call((&LessConstraint{expected: uint32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Uint32AndFunc").
		Call((&LessConstraint{expected: uint32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Uint32AndInterface").
		Call((&LessConstraint{expected: uint32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Uint32AndMap").
		Call((&LessConstraint{expected: uint32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Uint32AndPtr").
		Call((&LessConstraint{expected: uint32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Uint32AndSlice").
		Call((&LessConstraint{expected: uint32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Uint32AndString").
		Call((&LessConstraint{expected: uint32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Uint32AndStruct").
		Call((&LessConstraint{expected: uint32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Uint32AndUnsafePointer").
		Call((&LessConstraint{expected: uint32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithUint64(t *testing.T) {
	NewDeclarative(t, "Uint64ndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: uint64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Uint64AndBool").
		Call((&LessConstraint{expected: uint64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Uint64LessIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterIntWithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessIntWithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessIntWithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int(-5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterInt8WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessInt8WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt8WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int8(-5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterInt16WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessInt16WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt16WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int16(-5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterInt32WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessInt32WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt32WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int32(-5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterInt64WithIntComparison").
		Call((&LessConstraint{expected: uint64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessInt64WithUintComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessInt64WithFloatComparison").
		Call((&LessConstraint{expected: uint64(math.MaxInt64) + 1}).Check, int64(-5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64LessUint").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualUint").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterUint").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessUint8").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualUint8").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterUint8").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessUint16").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualUint16").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterUint16").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessUint32").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualUint32").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterUint32").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessUint64").
		Call((&LessConstraint{expected: uint64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualUint64").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterUint64").
		Call((&LessConstraint{expected: uint64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64AndUintPtr").
		Call((&LessConstraint{expected: uint64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Uint64LessFloat32").
		Call((&LessConstraint{expected: uint64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualFloat32").
		Call((&LessConstraint{expected: uint64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterFloat32").
		Call((&LessConstraint{expected: uint64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64LessFloat64").
		Call((&LessConstraint{expected: uint64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Uint64EqualFloat64").
		Call((&LessConstraint{expected: uint64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64GreaterFloat64").
		Call((&LessConstraint{expected: uint64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Uint64AndComplex64").
		Call((&LessConstraint{expected: uint64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Uint64AndComplex128").
		Call((&LessConstraint{expected: uint64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Uint64AndArray").
		Call((&LessConstraint{expected: uint64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Uint64AndChan").
		Call((&LessConstraint{expected: uint64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Uint64AndFunc").
		Call((&LessConstraint{expected: uint64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Uint64AndInterface").
		Call((&LessConstraint{expected: uint64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Uint64AndMap").
		Call((&LessConstraint{expected: uint64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Uint64AndPtr").
		Call((&LessConstraint{expected: uint64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Uint64AndSlice").
		Call((&LessConstraint{expected: uint64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Uint64AndString").
		Call((&LessConstraint{expected: uint64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Uint64AndStruct").
		Call((&LessConstraint{expected: uint64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Uint64AndUnsafePointer").
		Call((&LessConstraint{expected: uint64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithFloat32(t *testing.T) {
	NewDeclarative(t, "Float32AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: float32(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Float32AndBool").
		Call((&LessConstraint{expected: float32(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Float32LessInt").
		Call((&LessConstraint{expected: float32(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualInt").
		Call((&LessConstraint{expected: float32(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterInt").
		Call((&LessConstraint{expected: float32(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessInt8").
		Call((&LessConstraint{expected: float32(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualInt8").
		Call((&LessConstraint{expected: float32(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterInt8").
		Call((&LessConstraint{expected: float32(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessInt16").
		Call((&LessConstraint{expected: float32(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualInt16").
		Call((&LessConstraint{expected: float32(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterInt16").
		Call((&LessConstraint{expected: float32(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessInt32").
		Call((&LessConstraint{expected: float32(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualInt32").
		Call((&LessConstraint{expected: float32(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterInt32").
		Call((&LessConstraint{expected: float32(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessInt64").
		Call((&LessConstraint{expected: float32(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualInt64").
		Call((&LessConstraint{expected: float32(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterInt64").
		Call((&LessConstraint{expected: float32(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessUint").
		Call((&LessConstraint{expected: float32(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualUint").
		Call((&LessConstraint{expected: float32(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterUint").
		Call((&LessConstraint{expected: float32(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessUint8").
		Call((&LessConstraint{expected: float32(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualUint8").
		Call((&LessConstraint{expected: float32(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterUint8").
		Call((&LessConstraint{expected: float32(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessUint16").
		Call((&LessConstraint{expected: float32(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualUint16").
		Call((&LessConstraint{expected: float32(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterUint16").
		Call((&LessConstraint{expected: float32(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessUint32").
		Call((&LessConstraint{expected: float32(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualUint32").
		Call((&LessConstraint{expected: float32(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterUint32").
		Call((&LessConstraint{expected: float32(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessUint64").
		Call((&LessConstraint{expected: float32(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualUint64").
		Call((&LessConstraint{expected: float32(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterUint64").
		Call((&LessConstraint{expected: float32(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32AndUintPtr").
		Call((&LessConstraint{expected: float32(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Float32LessFloat32").
		Call((&LessConstraint{expected: float32(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualFloat32").
		Call((&LessConstraint{expected: float32(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterFloat32").
		Call((&LessConstraint{expected: float32(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32LessFloat64").
		Call((&LessConstraint{expected: float32(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float32EqualFloat64").
		Call((&LessConstraint{expected: float32(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float32GreaterFloat64").
		Call((&LessConstraint{expected: float32(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float32AndComplex64").
		Call((&LessConstraint{expected: float32(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Float32AndComplex128").
		Call((&LessConstraint{expected: float32(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Float32AndArray").
		Call((&LessConstraint{expected: float32(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Float32AndChan").
		Call((&LessConstraint{expected: float32(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Float32AndFunc").
		Call((&LessConstraint{expected: float32(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Float32AndInterface").
		Call((&LessConstraint{expected: float32(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Float32AndMap").
		Call((&LessConstraint{expected: float32(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Float32AndPtr").
		Call((&LessConstraint{expected: float32(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Float32AndSlice").
		Call((&LessConstraint{expected: float32(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Float32AndString").
		Call((&LessConstraint{expected: float32(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Float32AndStruct").
		Call((&LessConstraint{expected: float32(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Float32AndUnsafePointer").
		Call((&LessConstraint{expected: float32(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestLessConstraint_Check_WithFloat64(t *testing.T) {
	NewDeclarative(t, "Float64AndInvalid").
		Call(
			func() {
				(&LessConstraint{expected: float64(5)}).Check(nil)
			},
		).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", nil))

	NewDeclarative(t, "Float64AndBool").
		Call((&LessConstraint{expected: float64(5)}).Check, false).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", false))

	NewDeclarative(t, "Float64LessInt").
		Call((&LessConstraint{expected: float64(10)}).Check, int(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualInt").
		Call((&LessConstraint{expected: float64(5)}).Check, int(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterInt").
		Call((&LessConstraint{expected: float64(5)}).Check, int(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessInt8").
		Call((&LessConstraint{expected: float64(10)}).Check, int8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualInt8").
		Call((&LessConstraint{expected: float64(5)}).Check, int8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterInt8").
		Call((&LessConstraint{expected: float64(5)}).Check, int8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessInt16").
		Call((&LessConstraint{expected: float64(10)}).Check, int16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualInt16").
		Call((&LessConstraint{expected: float64(5)}).Check, int16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterInt16").
		Call((&LessConstraint{expected: float64(5)}).Check, int16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessInt32").
		Call((&LessConstraint{expected: float64(10)}).Check, int32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualInt32").
		Call((&LessConstraint{expected: float64(5)}).Check, int32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterInt32").
		Call((&LessConstraint{expected: float64(5)}).Check, int32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessInt64").
		Call((&LessConstraint{expected: float64(10)}).Check, int64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualInt64").
		Call((&LessConstraint{expected: float64(5)}).Check, int64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterInt64").
		Call((&LessConstraint{expected: float64(5)}).Check, int64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessUint").
		Call((&LessConstraint{expected: float64(10)}).Check, uint(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualUint").
		Call((&LessConstraint{expected: float64(5)}).Check, uint(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterUint").
		Call((&LessConstraint{expected: float64(5)}).Check, uint(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessUint8").
		Call((&LessConstraint{expected: float64(10)}).Check, uint8(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualUint8").
		Call((&LessConstraint{expected: float64(5)}).Check, uint8(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterUint8").
		Call((&LessConstraint{expected: float64(5)}).Check, uint8(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessUint16").
		Call((&LessConstraint{expected: float64(10)}).Check, uint16(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualUint16").
		Call((&LessConstraint{expected: float64(5)}).Check, uint16(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterUint16").
		Call((&LessConstraint{expected: float64(5)}).Check, uint16(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessUint32").
		Call((&LessConstraint{expected: float64(10)}).Check, uint32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualUint32").
		Call((&LessConstraint{expected: float64(5)}).Check, uint32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterUint32").
		Call((&LessConstraint{expected: float64(5)}).Check, uint32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessUint64").
		Call((&LessConstraint{expected: float64(10)}).Check, uint64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualUint64").
		Call((&LessConstraint{expected: float64(5)}).Check, uint64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterUint64").
		Call((&LessConstraint{expected: float64(5)}).Check, uint64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64AndUintPtr").
		Call((&LessConstraint{expected: float64(5)}).Check, uintptr(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", uintptr(5)))

	NewDeclarative(t, "Float64LessFloat32").
		Call((&LessConstraint{expected: float64(10)}).Check, float32(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualFloat32").
		Call((&LessConstraint{expected: float64(5)}).Check, float32(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterFloat32").
		Call((&LessConstraint{expected: float64(5)}).Check, float32(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64LessFloat64").
		Call((&LessConstraint{expected: float64(10)}).Check, float64(5)).
		ExpectResult(true)

	NewDeclarative(t, "Float64EqualFloat64").
		Call((&LessConstraint{expected: float64(5)}).Check, float64(5)).
		ExpectResult(false)

	NewDeclarative(t, "Float64GreaterFloat64").
		Call((&LessConstraint{expected: float64(5)}).Check, float64(10)).
		ExpectResult(false)

	NewDeclarative(t, "Float64AndComplex64").
		Call((&LessConstraint{expected: float64(5)}).Check, complex64(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex64(5)))

	NewDeclarative(t, "Float64AndComplex128").
		Call((&LessConstraint{expected: float64(5)}).Check, complex128(5)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", complex128(5)))

	NewDeclarative(t, "Float64AndArray").
		Call((&LessConstraint{expected: float64(5)}).Check, [1]int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", [1]int{5}))

	NewDeclarative(t, "Float64AndChan").
		Call((&LessConstraint{expected: float64(5)}).Check, make(chan int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", make(chan int)))

	NewDeclarative(t, "Float64AndFunc").
		Call((&LessConstraint{expected: float64(5)}).Check, func() {}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", func() {}))

	NewDeclarative(t, "Float64AndInterface").
		Call((&LessConstraint{expected: float64(5)}).Check, (*interface{})(nil)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", (*interface{})(nil)))

	NewDeclarative(t, "Float64AndMap").
		Call((&LessConstraint{expected: float64(5)}).Check, map[int]int{1: 1}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", map[int]int{1: 1}))

	NewDeclarative(t, "Float64AndPtr").
		Call((&LessConstraint{expected: float64(5)}).Check, new(int)).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", new(int)))

	NewDeclarative(t, "Float64AndSlice").
		Call((&LessConstraint{expected: float64(5)}).Check, []int{5}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", []int{5}))

	NewDeclarative(t, "Float64AndString").
		Call((&LessConstraint{expected: float64(5)}).Check, "data").
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", "data"))

	NewDeclarative(t, "Float64AndStruct").
		Call((&LessConstraint{expected: float64(5)}).Check, struct{}{}).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", struct{}{}))

	NewDeclarative(t, "Float64AndUnsafePointer").
		Call((&LessConstraint{expected: float64(5)}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(NewInvalidNumericComparisonTypeError("actual", unsafe.Pointer(new(int))))
}

func TestLessConstraint_String(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&LessConstraint{expected: 55}).String).
		ExpectResult(fmt.Sprintf("be less than %v", 55))
}

func TestLessConstraint_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&LessConstraint{expected: 55}).Details, 10).
		ExpectResult("")
}
