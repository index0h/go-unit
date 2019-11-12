package unit

import (
	"testing"
	"unsafe"
)

func TestNewEmptyConstraint(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call(NewEmptyConstraint).
		ExpectResult(ConstraintAsValue{Value: &EmptyConstraint{comparator: NewEqualComparator()}})
}

//noinspection GoRedundantConversion
func TestEmptyConstraint_Check(t *testing.T) {
	NewSubtest(t, "WithInvalidValueWithPositiveResult").
		Call(
			func() bool {
				return (&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check(nil)
			},
		).
		ExpectResult(true)

	NewSubtest(t, "WithBoolValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t).RecordCompare(false, false, true)}).Check, true).
		ExpectResult(true)

	NewSubtest(t, "WithBoolValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t).RecordCompare(true, false, false)}).Check, false).
		ExpectResult(false)

	NewSubtest(t, "WithIntValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t).RecordCompare(int(0), int(0), true)}).Check, int(0)).
		ExpectResult(true)

	NewSubtest(t, "WithIntValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int(5), int(0), false),
			}).Check,
			int(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithInt8ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int8(0), int8(0), true),
			}).Check,
			int8(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithInt8ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int8(5), int8(0), false),
			}).Check,
			int8(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithInt16ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int16(0), int16(0), true),
			}).Check,
			int16(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithInt16ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int16(5), int16(0), false),
			}).Check,
			int16(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithInt32ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int32(0), int32(0), true),
			}).Check,
			int32(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithInt32ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int32(5), int32(0), false),
			}).Check,
			int32(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithInt64ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int64(0), int64(0), true),
			}).Check,
			int64(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithInt64ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(int64(5), int64(0), false),
			}).Check,
			int64(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUintValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint(0), uint(0), true),
			}).Check,
			uint(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUintValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint(5), uint(0), false),
			}).Check,
			uint(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUint8ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint8(0), uint8(0), true),
			}).Check,
			uint8(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUint8ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint8(5), uint8(0), false),
			}).Check,
			uint8(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUint16ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint16(0), uint16(0), true),
			}).Check,
			uint16(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUint16ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint16(5), uint16(0), false),
			}).Check,
			uint16(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUint32ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint32(0), uint32(0), true),
			}).Check,
			uint32(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUint32ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint32(5), uint32(0), false),
			}).Check,
			uint32(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUint64ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint64(0), uint64(0), true),
			}).Check,
			uint64(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUint64ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uint64(5), uint64(0), false),
			}).Check,
			uint64(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithUintptrValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uintptr(0), uintptr(0), true),
			}).Check,
			uintptr(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithUintptrValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(uintptr(5), uintptr(0), false),
			}).Check,
			uintptr(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithFloat32ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(float32(0), float32(0), true),
			}).Check,
			float32(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithFloat32ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(float32(5), float32(0), false),
			}).Check,
			float32(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithFloat64ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(float64(0), float64(0), true),
			}).Check,
			float64(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithFloat64ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(float64(5), float64(0), false),
			}).Check,
			float64(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithComplex64ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(complex64(0), complex64(0), true),
			}).Check,
			complex64(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithComplex64ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(complex64(5), complex64(0), false),
			}).Check,
			complex64(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithComplex128ValueAndPositiveResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(complex128(0), complex128(0), true),
			}).Check,
			complex128(0),
		).
		ExpectResult(true)

	NewSubtest(t, "WithComplex128ValueAndNegativeResult").
		Call(
			(&EmptyConstraint{
				comparator: NewMockEqualComparer(t).RecordCompare(complex128(5), complex128(0), false),
			}).Check,
			complex128(5),
		).
		ExpectResult(false)

	NewSubtest(t, "WithArrayValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, [2]int{0, 0}).
		ExpectResult(true)

	NewSubtest(t, "WithArrayValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, [2]int{1, 5}).
		ExpectResult(false)

	NewSubtest(t, "WithChanValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, make(chan int)).
		ExpectResult(true)

	{
		chanFixture := make(chan int, 1)
		chanFixture <- 1

		NewSubtest(t, "WithChanValueAndPositiveResult").
			Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, chanFixture).
			ExpectResult(false)
	}

	NewSubtest(t, "WithFuncValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, (func())(nil)).
		ExpectResult(true)

	NewSubtest(t, "WithFuncValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, func() {}).
		ExpectResult(false)

	NewSubtest(t, "WithInterfaceValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, (*interface{})(nil)).
		ExpectResult(true)

	NewSubtest(t, "WithMapValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, map[string]int{}).
		ExpectResult(true)

	NewSubtest(t, "WithMapValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, map[string]int{"data": 1}).
		ExpectResult(false)

	NewSubtest(t, "WithPtrValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, (*int)(nil)).
		ExpectResult(true)

	NewSubtest(t, "WithPtrValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, new(int)).
		ExpectResult(false)

	NewSubtest(t, "WithSliceValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, []int{}).
		ExpectResult(true)

	NewSubtest(t, "WithSliceValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, []int{1}).
		ExpectResult(false)

	NewSubtest(t, "WithStringValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, "").
		ExpectResult(true)

	NewSubtest(t, "WithStringValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, "data").
		ExpectResult(false)

	{
		structFixtureEmpty := struct {
			data  int
			value string
		}{
			data:  0,
			value: "",
		}

		structFixtureNotEmpty := struct {
			data  int
			value string
		}{
			data:  1000,
			value: "data",
		}

		NewSubtest(t, "WithStructValueAndPositiveResult").
			Call(
				(&EmptyConstraint{
					comparator: NewMockEqualComparer(t).RecordCompare(
						structFixtureEmpty,
						structFixtureEmpty,
						true,
					),
				}).Check,
				structFixtureEmpty,
			).
			ExpectResult(true)

		NewSubtest(t, "WithStructValueAndNegativeResult").
			Call(
				(&EmptyConstraint{
					comparator: NewMockEqualComparer(t).RecordCompare(
						structFixtureNotEmpty,
						structFixtureEmpty,
						false,
					),
				}).Check,
				structFixtureNotEmpty,
			).
			ExpectResult(false)
	}

	NewSubtest(t, "WithUnsafePointerValueAndPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, (unsafe.Pointer)(nil)).
		ExpectResult(true)

	NewSubtest(t, "WithUnsafePointerValueAndNegativeResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Check, unsafe.Pointer(new(int))).
		ExpectResult(false)
}

func TestEmptyConstraint_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).String).
		ExpectResult("be empty")
}

func TestEmptyConstraint_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&EmptyConstraint{comparator: NewMockEqualComparer(t)}).Details, "data").
		ExpectResult("")
}
