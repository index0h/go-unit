package unit

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestNewContains(t *testing.T) {
	NewSubtest(t, "Success").
		Call(NewContainsConstraint, "element").
		ExpectResult(
			ConstraintAsValue{Value: &ContainsConstraint{element: "element", comparator: NewEqualComparator()}},
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementBool(t *testing.T) {
	NewSubtest(t, "ElementBoolAndListBool").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInt").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInt8").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInt16").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInt32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInt64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUint").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUint8").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUint16").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUint32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUint64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUintptr").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListFloat32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListFloat64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListComplex64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementBoolAndListComplex128").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementBoolAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, [2]interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementBoolAndListArrayAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, [2]interface{}{false, true}).
		ExpectResult(true)

	NewSubtest(t, "ElementBoolAndListChan").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementBoolAndListFunc").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListInterface").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementBoolAndListMapAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementBoolAndListMapAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": true}).
		ExpectResult(true)

	NewSubtest(t, "ElementBoolAndListPtr").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementBoolAndListSliceAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, []interface{}{false, true}).
		ExpectResult(true)

	NewSubtest(t, "ElementBoolAndListString").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", true, reflect.String))

	NewSubtest(t, "ElementBoolAndListStruct").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementBoolAndListUnsafePointer").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInt(t *testing.T) {
	NewSubtest(t, "ElementIntAndListBool").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInt").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInt8").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInt16").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInt32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInt64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUint").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUint8").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUint16").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUint32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUint64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUintptr").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListFloat32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListFloat64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListComplex64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementIntAndListComplex128").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementIntAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementIntAndListArrayAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementIntAndListChan").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementIntAndListFunc").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListInterface").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementIntAndListMapAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementIntAndListMapAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": int(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementIntAndListPtr").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementIntAndListSliceAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementIntAndListString").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int(5), reflect.String))

	NewSubtest(t, "ElementIntAndListStruct").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementIntAndListUnsafePointer").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInt8(t *testing.T) {
	NewSubtest(t, "ElementInt8AndListBool").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInt").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInt8").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInt16").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInt32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInt64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUint").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUint8").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUint16").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUint32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUint64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUintptr").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListFloat32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListFloat64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListComplex64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt8AndListComplex128").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt8AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt8AndListArrayAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int8(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt8AndListChan").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt8AndListFunc").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListInterface").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementInt8AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt8AndListMapAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": int8(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt8AndListPtr").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt8AndListSliceAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int8(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt8AndListString").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int8(5), reflect.String))

	NewSubtest(t, "ElementInt8AndListStruct").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt8AndListUnsafePointer").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInt16(t *testing.T) {
	NewSubtest(t, "ElementInt16AndListBool").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInt").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInt8").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInt16").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInt32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInt64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUint").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUint8").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUint16").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUint32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUint64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUintptr").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListFloat32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListFloat64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListComplex64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt16AndListComplex128").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt16AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt16AndListArrayAndContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int16(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt16AndListChan").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt16AndListFunc").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListInterface").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementInt16AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt16AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int16(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementInt16AndListPtr").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt16AndListSliceAndContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int16(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt16AndListString").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int16(5), reflect.String))

	NewSubtest(t, "ElementInt16AndListStruct").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt16AndListUnsafePointer").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInt32(t *testing.T) {
	NewSubtest(t, "ElementInt32AndListBool").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInt").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInt8").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInt16").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInt32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInt64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUint").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUint8").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUint16").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUint32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUint64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUintptr").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListFloat32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListFloat64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListComplex64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt32AndListComplex128").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt32AndListArrayAndContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt32AndListChan").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt32AndListFunc").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListInterface").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementInt32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int32(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementInt32AndListPtr").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt32AndListSliceAndContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt32AndListString").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int32(5), reflect.String))

	NewSubtest(t, "ElementInt32AndListStruct").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt32AndListUnsafePointer").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInt64(t *testing.T) {
	NewSubtest(t, "ElementInt64AndListBool").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInt").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInt8").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInt16").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInt32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInt64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUint").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUint8").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUint16").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUint32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUint64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUintptr").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListFloat32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListFloat64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListComplex64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt64AndListComplex128").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt64AndListArrayAndContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt64AndListChan").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInt64AndListFunc").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListInterface").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementInt64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int64(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementInt64AndListPtr").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInt64AndListSliceAndContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInt64AndListString").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int64(5), reflect.String))

	NewSubtest(t, "ElementInt64AndListStruct").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInt64AndListUnsafePointer").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUint(t *testing.T) {
	NewSubtest(t, "ElementUintAndListBool").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInt").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInt8").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInt16").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInt32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInt64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUint").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUint8").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUint16").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUint32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUint64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUintptr").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListFloat32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListFloat64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListComplex64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintAndListComplex128").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintAndListArrayAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUintAndListChan").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintAndListFunc").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListInterface").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUintAndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintAndListMapAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": uint(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUintAndListPtr").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintAndListSliceAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUintAndListString").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint(5), reflect.String))

	NewSubtest(t, "ElementUintAndListStruct").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintAndListUnsafePointer").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUint8(t *testing.T) {
	NewSubtest(t, "ElementUint8AndListBool").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInt").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInt8").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInt16").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInt32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInt64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUint").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUint8").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUint16").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUint32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUint64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUintptr").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListFloat32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListFloat64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListComplex64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint8AndListComplex128").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint8AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint8AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint8(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint8AndListChan").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint8AndListFunc").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListInterface").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUint8AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint8AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint8(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUint8AndListPtr").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint8AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint8(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint8AndListString").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint8(5), reflect.String))

	NewSubtest(t, "ElementUint8AndListStruct").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint8AndListUnsafePointer").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUint16(t *testing.T) {
	NewSubtest(t, "ElementUint16AndListBool").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInt").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInt8").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInt16").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInt32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInt64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUint").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUint8").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUint16").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUint32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUint64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUintptr").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListFloat32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListFloat64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListComplex64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint16AndListComplex128").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint16AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint16AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint16(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint16AndListChan").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint16AndListFunc").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListInterface").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUint16AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint16AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint16(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUint16AndListPtr").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint16AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint16(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint16AndListString").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint16(5), reflect.String))

	NewSubtest(t, "ElementUint16AndListStruct").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint16AndListUnsafePointer").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUint32(t *testing.T) {
	NewSubtest(t, "ElementUint32AndListBool").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInt").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInt8").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInt16").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInt32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInt64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUint").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUint8").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUint16").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUint32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUint64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUintptr").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListFloat32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListFloat64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListComplex64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint32AndListComplex128").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint32AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint32AndListChan").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint32AndListFunc").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListInterface").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUint32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint32(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUint32AndListPtr").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint32AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint32AndListString").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint32(5), reflect.String))

	NewSubtest(t, "ElementUint32AndListStruct").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint32AndListUnsafePointer").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUint64(t *testing.T) {
	NewSubtest(t, "ElementUint64AndListBool").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInt").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInt8").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInt16").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInt32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInt64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUint").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUint64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUint16").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUint32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUint64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUintptr").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListFloat32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListFloat64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListComplex64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint64AndListComplex128").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint64AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint64AndListChan").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUint64AndListFunc").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListInterface").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUint64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint64(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUint64AndListPtr").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUint64AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUint64AndListString").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint64(5), reflect.String))

	NewSubtest(t, "ElementUint64AndListStruct").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUint64AndListUnsafePointer").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUintptr(t *testing.T) {
	NewSubtest(t, "ElementUintptrAndListBool").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInt").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInt8").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInt16").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInt32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInt64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUint").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUint8").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUint16").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUint32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUint64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUintptr").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListFloat32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListFloat64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListComplex64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintptrAndListComplex128").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintptrAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintptrAndListArrayAndContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uintptr(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUintptrAndListChan").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUintptrAndListFunc").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListInterface").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUintptrAndListMapAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintptrAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uintptr(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUintptrAndListPtr").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUintptrAndListSliceAndContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uintptr(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementUintptrAndListString").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uintptr(5), reflect.String))

	NewSubtest(t, "ElementUintptrAndListStruct").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUintptrAndListUnsafePointer").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementFloat32(t *testing.T) {
	NewSubtest(t, "ElementFloat32AndListBool").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInt").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInt8").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInt16").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInt32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInt64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUint").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUint8").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUint16").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUint32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUint64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUintptr").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListFloat32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListFloat64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListComplex64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat32AndListComplex128").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat32AndListArrayAndContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, float32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat32AndListChan").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat32AndListFunc").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListInterface").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementFloat32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": float32(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat32AndListPtr").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat32AndListSliceAndContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, float32(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat32AndListString").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", float32(5), reflect.String))

	NewSubtest(t, "ElementFloat32AndListStruct").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat32AndListUnsafePointer").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementFloat64(t *testing.T) {
	NewSubtest(t, "ElementFloat64AndListBool").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInt").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInt8").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInt16").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInt32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInt64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUint").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUint8").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUint16").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUint32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUint64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUintptr").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListFloat32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListFloat64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListComplex64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat64AndListComplex128").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat64AndListArrayAndContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, float64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat64AndListChan").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFloat64AndListFunc").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListInterface").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementFloat64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": float64(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat64AndListPtr").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementFloat64AndListSliceAndContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, float64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementFloat64AndListString").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", float64(5), reflect.String))

	NewSubtest(t, "ElementFloat64AndListStruct").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFloat64AndListUnsafePointer").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementComplex64(t *testing.T) {
	NewSubtest(t, "ElementComplex64AndListBool").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInt").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInt8").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInt16").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInt32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInt64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUint").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUint8").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUint16").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUint32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUint64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUintptr").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListFloat32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListFloat64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListComplex64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex64AndListComplex128").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex64AndListArrayAndContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, complex64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex64AndListChan").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex64AndListFunc").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListInterface").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementComplex64AndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": complex64(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex64AndListPtr").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex64AndListSliceAndContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, complex64(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex64AndListString").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", complex64(5), reflect.String))

	NewSubtest(t, "ElementComplex64AndListStruct").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex64AndListUnsafePointer").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementComplex128(t *testing.T) {
	NewSubtest(t, "ElementComplex128AndListBool").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInt").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInt8").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInt16").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInt32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInt64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUint").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUint8").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUint16").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUint32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUint64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUintptr").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListFloat32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListFloat64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListComplex64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex128AndListComplex128").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex128AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex128AndListArrayAndContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, complex128(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex128AndListChan").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementComplex128AndListFunc").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListInterface").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementComplex128AndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex128AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": complex128(5)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex128AndListPtr").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementComplex128AndListSliceAndContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, []interface{}{false, complex128(5)}).
		ExpectResult(true)

	NewSubtest(t, "ElementComplex128AndListString").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", complex128(5), reflect.String))

	NewSubtest(t, "ElementComplex128AndListStruct").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementComplex128AndListUnsafePointer").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementArray(t *testing.T) {
	NewSubtest(t, "ElementArrayAndListBool").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInt").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInt8").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInt16").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInt32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInt64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUint").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUint8").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUint16").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUint32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUint64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUintptr").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListFloat32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListFloat64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListComplex64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementArrayAndListComplex128").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementArrayAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementArrayAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementArrayAndListChan").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementArrayAndListFunc").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListInterface").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementArrayAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementArrayAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementArrayAndListPtr").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementArrayAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementArrayAndListString").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", [2]interface{}{false, 10}, reflect.String))

	NewSubtest(t, "ElementArrayAndListStruct").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementArrayAndListUnsafePointer").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementChan(t *testing.T) {
	chanFixture := make(chan int)

	NewSubtest(t, "ElementChanAndListBool").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInt").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInt8").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInt16").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInt32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInt64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUint").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUint8").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUint16").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUint32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUint64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUintptr").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListFloat32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListFloat64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListComplex64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementChanAndListComplex128").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementChanAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementChanAndListArrayAndContains").
		Call((&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check, [2]interface{}{false, chanFixture}).
		ExpectResult(true)

	NewSubtest(t, "ElementChanAndListChan").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementChanAndListFunc").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListInterface").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementChanAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementChanAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": chanFixture},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementChanAndListPtr").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementChanAndListSliceAndContains").
		Call((&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check, []interface{}{false, chanFixture}).
		ExpectResult(true)

	NewSubtest(t, "ElementChanAndListString").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", make(chan int), reflect.String))

	NewSubtest(t, "ElementChanAndListStruct").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementChanAndListUnsafePointer").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementFunc(t *testing.T) {
	funcFixture := func() {}

	NewSubtest(t, "ElementFuncAndListBool").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInt").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInt8").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInt16").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInt32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInt64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUint").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUint8").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUint16").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUint32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUint64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUintptr").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListFloat32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListFloat64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListComplex64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFuncAndListComplex128").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFuncAndListArrayAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, 10},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementFuncAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			[2]interface{}{funcFixture, false},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementFuncAndListChan").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementFuncAndListFunc").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListInterface").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementFuncAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementFuncAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": funcFixture},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementFuncAndListPtr").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListSliceAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementFuncAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, funcFixture},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementFuncAndListString").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", func() {}, reflect.String))

	NewSubtest(t, "ElementFuncAndListStruct").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementFuncAndListUnsafePointer").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementInterface(t *testing.T) {
	NewSubtest(t, "ElementInterfaceAndListBool").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInt").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInt8").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInt16").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInt32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInt64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUint").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUint8").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUint16").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUint32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUint64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUintptr").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListFloat32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListFloat64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListComplex64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInterfaceAndListComplex128").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInterfaceAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementInterfaceAndListArrayAndContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, [2]interface{}{false, (*interface{})(nil)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInterfaceAndListChan").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementInterfaceAndListFunc").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListInterface").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementInterfaceAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementInterfaceAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": (*interface{})(nil)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementInterfaceAndListPtr").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementInterfaceAndListSliceAndContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, []interface{}{false, (*interface{})(nil)}).
		ExpectResult(true)

	NewSubtest(t, "ElementInterfaceAndListString").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", (*interface{})(nil), reflect.String))

	NewSubtest(t, "ElementInterfaceAndListStruct").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementInterfaceAndListUnsafePointer").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementMap(t *testing.T) {
	NewSubtest(t, "ElementMapAndListBool").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInt").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInt8").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInt16").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInt32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInt64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUint").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUint8").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUint16").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUint32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUint64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUintptr").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListFloat32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListFloat64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListComplex64").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			complex64(5),
		).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementMapAndListComplex128").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			complex128(5),
		).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementMapAndListArrayAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, 10},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementMapAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementMapAndListChan").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			make(chan uint),
		).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementMapAndListFunc").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListInterface").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			(*interface{})(nil),
		).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementMapAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementMapAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementMapAndListPtr").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListSliceAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementMapAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementMapAndListString").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(
			NewInvalidKindError("element", map[string]interface{}{"First": false, "Second": false}, reflect.String),
		)

	NewSubtest(t, "ElementMapAndListStruct").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementMapAndListUnsafePointer").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			unsafe.Pointer(new(int)),
		).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementPtr(t *testing.T) {
	NewSubtest(t, "ElementPtrAndListBool").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInt").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInt8").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInt16").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInt32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInt64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUint").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUint8").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUint16").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUint32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUint64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUintptr").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListFloat32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListFloat64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListComplex64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementPtrAndListComplex128").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementPtrAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementPtrAndListArrayAndContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, new(int)}).
		ExpectResult(true)

	NewSubtest(t, "ElementPtrAndListChan").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementPtrAndListFunc").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListInterface").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementPtrAndListMapAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementPtrAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": new(int)},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementPtrAndListPtr").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementPtrAndListSliceAndContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, []interface{}{false, new(int)}).
		ExpectResult(true)

	NewSubtest(t, "ElementPtrAndListString").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", new(int), reflect.String))

	NewSubtest(t, "ElementPtrAndListStruct").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementPtrAndListUnsafePointer").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementSlice(t *testing.T) {
	NewSubtest(t, "ElementSliceAndListBool").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInt").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInt8").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInt16").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInt32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInt64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUint").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUint8").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUint16").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUint32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUint64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUintptr").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListFloat32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListFloat64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListComplex64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementSliceAndListComplex128").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementSliceAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementSliceAndListArrayAndContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, []interface{}{false}}).
		ExpectResult(true)

	NewSubtest(t, "ElementSliceAndListChan").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementSliceAndListFunc").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListInterface").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementSliceAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementSliceAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": []interface{}{false}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementSliceAndListPtr").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementSliceAndListSliceAndContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, []interface{}{false, []interface{}{false}}).
		ExpectResult(true)

	NewSubtest(t, "ElementSliceAndListString").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", []interface{}{false}, reflect.String))

	NewSubtest(t, "ElementSliceAndListStruct").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementSliceAndListUnsafePointer").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementString(t *testing.T) {
	NewSubtest(t, "ElementStringAndListBool").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInt").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInt8").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInt16").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInt32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInt64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUint").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUint8").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUint16").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUint32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUint64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUintptr").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListFloat32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListFloat64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListComplex64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStringAndListComplex128").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStringAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementStringAndListArrayAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, [2]interface{}{false, "data"}).
		ExpectResult(true)

	NewSubtest(t, "ElementStringAndListChan").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStringAndListFunc").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListInterface").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementStringAndListMapAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementStringAndListMapAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": "data"}).
		ExpectResult(true)

	NewSubtest(t, "ElementStringAndListPtr").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementStringAndListSliceAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, []interface{}{false, "data"}).
		ExpectResult(true)

	NewSubtest(t, "ElementStringAndListStringAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, "not contains").
		ExpectResult(false)

	NewSubtest(t, "ElementStringAndListStringAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, "some data here").
		ExpectResult(true)

	NewSubtest(t, "ElementStringAndListStruct").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStringAndListUnsafePointer").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementStruct(t *testing.T) {
	NewSubtest(t, "ElementStructAndListBool").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInt").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInt8").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInt16").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInt32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInt64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUint").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUint8").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUint16").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUint32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUint64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUintptr").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListFloat32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListFloat64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListComplex64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStructAndListComplex128").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStructAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementStructAndListArrayAndContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, struct{}{}}).
		ExpectResult(true)

	NewSubtest(t, "ElementStructAndListChan").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementStructAndListFunc").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListInterface").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementStructAndListMapAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewSubtest(t, "ElementStructAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": struct{}{}},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementStructAndListPtr").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementStructAndListSliceAndContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, []interface{}{false, struct{}{}}).
		ExpectResult(true)

	NewSubtest(t, "ElementStructAndListString").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", struct{}{}, reflect.String))

	NewSubtest(t, "ElementStructAndListStruct").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementStructAndListUnsafePointer").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementUnsafePointer(t *testing.T) {
	unsafePointerFixture := unsafe.Pointer(new(int))

	NewSubtest(t, "ElementUnsafePointerAndListBool").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInt").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInt8").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInt16").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInt32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInt64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUint").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUint8").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUint16").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUint32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUint64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUintptr").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListFloat32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListFloat64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListComplex64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUnsafePointerAndListComplex128").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUnsafePointerAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewSubtest(t, "ElementUnsafePointerAndListArrayAndContains").
		Call((&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check, [2]interface{}{false, unsafePointerFixture}).
		ExpectResult(true)

	NewSubtest(t, "ElementUnsafePointerAndListChan").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewSubtest(t, "ElementUnsafePointerAndListFunc").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListInterface").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, (*interface{})(nil)).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				(*interface{})(nil),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)

	NewSubtest(t, "ElementUnsafePointerAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewSubtest(t, "ElementUnsafePointerAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": unsafePointerFixture},
		).
		ExpectResult(true)

	NewSubtest(t, "ElementUnsafePointerAndListPtr").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewSubtest(t, "ElementUnsafePointerAndListSliceAndContains").
		Call((&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check, []interface{}{false, unsafePointerFixture}).
		ExpectResult(true)

	NewSubtest(t, "ElementUnsafePointerAndListString").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", unsafe.Pointer(new(int)), reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListStruct").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewSubtest(t, "ElementUnsafePointerAndListUnsafePointer").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, unsafe.Pointer(new(int))).
		ExpectPanic(
			NewInvalidKindError(
				"list",
				unsafe.Pointer(new(int)),
				reflect.Array,
				reflect.Slice,
				reflect.Map,
				reflect.String,
			),
		)
}

func TestContains_String(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&ContainsConstraint{element: "element"}).String).
		ExpectResult(fmt.Sprintf("contain %+v", "element"))
}

func TestContains_Details(t *testing.T) {
	NewSubtest(t, "WithPositiveResult").
		Call((&ContainsConstraint{element: "element"}).Details, [2]string{"First", "element"}).
		ExpectResult("")
}
