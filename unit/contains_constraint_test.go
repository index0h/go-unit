package unit

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestNewContains(t *testing.T) {
	NewDeclarative(t, "Success").
		Call(NewContainsConstraint, "element").
		ExpectResult(
			ConstraintAsValue{Value: &ContainsConstraint{element: "element", comparator: NewEqualComparator()}},
		)
}

//noinspection GoRedundantConversion
func TestContains_Check_WithElementBool(t *testing.T) {
	NewDeclarative(t, "ElementBoolAndListBool").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInt").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInt8").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInt16").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInt32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInt64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUint").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUint8").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUint16").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUint32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUint64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUintptr").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListFloat32").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListFloat64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListComplex64").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementBoolAndListComplex128").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementBoolAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, [2]interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementBoolAndListArrayAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, [2]interface{}{false, true}).
		ExpectResult(true)

	NewDeclarative(t, "ElementBoolAndListChan").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementBoolAndListFunc").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListInterface").
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

	NewDeclarative(t, "ElementBoolAndListMapAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementBoolAndListMapAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": true}).
		ExpectResult(true)

	NewDeclarative(t, "ElementBoolAndListPtr").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementBoolAndListSliceAndContains").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, []interface{}{false, true}).
		ExpectResult(true)

	NewDeclarative(t, "ElementBoolAndListString").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", true, reflect.String))

	NewDeclarative(t, "ElementBoolAndListStruct").
		Call((&ContainsConstraint{element: true, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementBoolAndListUnsafePointer").
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
	NewDeclarative(t, "ElementIntAndListBool").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInt").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInt8").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInt16").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInt32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInt64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUint").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUint8").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUint16").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUint32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUint64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUintptr").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListFloat32").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListFloat64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListComplex64").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementIntAndListComplex128").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementIntAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementIntAndListArrayAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementIntAndListChan").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementIntAndListFunc").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListInterface").
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

	NewDeclarative(t, "ElementIntAndListMapAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementIntAndListMapAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": int(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementIntAndListPtr").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementIntAndListSliceAndContains").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementIntAndListString").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int(5), reflect.String))

	NewDeclarative(t, "ElementIntAndListStruct").
		Call((&ContainsConstraint{element: int(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementIntAndListUnsafePointer").
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
	NewDeclarative(t, "ElementInt8AndListBool").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInt").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInt8").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInt16").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInt32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInt64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUint").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUint8").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUint16").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUint32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUint64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUintptr").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListFloat32").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListFloat64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListComplex64").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt8AndListComplex128").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt8AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt8AndListArrayAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int8(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt8AndListChan").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt8AndListFunc").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListInterface").
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

	NewDeclarative(t, "ElementInt8AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt8AndListMapAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": int8(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt8AndListPtr").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt8AndListSliceAndContains").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int8(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt8AndListString").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int8(5), reflect.String))

	NewDeclarative(t, "ElementInt8AndListStruct").
		Call((&ContainsConstraint{element: int8(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt8AndListUnsafePointer").
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
	NewDeclarative(t, "ElementInt16AndListBool").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInt").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInt8").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInt16").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInt32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInt64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUint").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUint8").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUint16").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUint32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUint64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUintptr").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListFloat32").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListFloat64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListComplex64").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt16AndListComplex128").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt16AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt16AndListArrayAndContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int16(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt16AndListChan").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt16AndListFunc").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListInterface").
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

	NewDeclarative(t, "ElementInt16AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt16AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int16(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt16AndListPtr").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt16AndListSliceAndContains").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int16(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt16AndListString").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int16(5), reflect.String))

	NewDeclarative(t, "ElementInt16AndListStruct").
		Call((&ContainsConstraint{element: int16(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt16AndListUnsafePointer").
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
	NewDeclarative(t, "ElementInt32AndListBool").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInt").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInt8").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInt16").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInt32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInt64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUint").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUint8").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUint16").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUint32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUint64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUintptr").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListFloat32").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListFloat64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListComplex64").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt32AndListComplex128").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt32AndListArrayAndContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt32AndListChan").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt32AndListFunc").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListInterface").
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

	NewDeclarative(t, "ElementInt32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int32(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt32AndListPtr").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt32AndListSliceAndContains").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt32AndListString").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int32(5), reflect.String))

	NewDeclarative(t, "ElementInt32AndListStruct").
		Call((&ContainsConstraint{element: int32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt32AndListUnsafePointer").
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
	NewDeclarative(t, "ElementInt64AndListBool").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInt").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInt8").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInt16").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInt32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInt64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUint").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUint8").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUint16").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUint32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUint64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUintptr").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListFloat32").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListFloat64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListComplex64").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt64AndListComplex128").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt64AndListArrayAndContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, int64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt64AndListChan").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, make(chan int)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan int), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInt64AndListFunc").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListInterface").
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

	NewDeclarative(t, "ElementInt64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": int64(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt64AndListPtr").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInt64AndListSliceAndContains").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, int64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInt64AndListString").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", int64(5), reflect.String))

	NewDeclarative(t, "ElementInt64AndListStruct").
		Call((&ContainsConstraint{element: int64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInt64AndListUnsafePointer").
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
	NewDeclarative(t, "ElementUintAndListBool").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInt").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInt8").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInt16").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInt32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInt64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUint").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUint8").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUint16").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUint32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUint64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUintptr").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListFloat32").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListFloat64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListComplex64").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintAndListComplex128").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintAndListArrayAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintAndListChan").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintAndListFunc").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListInterface").
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

	NewDeclarative(t, "ElementUintAndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintAndListMapAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": uint(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintAndListPtr").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintAndListSliceAndContains").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintAndListString").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint(5), reflect.String))

	NewDeclarative(t, "ElementUintAndListStruct").
		Call((&ContainsConstraint{element: uint(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintAndListUnsafePointer").
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
	NewDeclarative(t, "ElementUint8AndListBool").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInt").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInt8").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInt16").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInt32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInt64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUint").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUint8").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUint16").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUint32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUint64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUintptr").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListFloat32").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListFloat64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListComplex64").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint8AndListComplex128").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint8AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint8AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint8(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint8AndListChan").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint8AndListFunc").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListInterface").
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

	NewDeclarative(t, "ElementUint8AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint8AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint8(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint8AndListPtr").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint8AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint8(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint8AndListString").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint8(5), reflect.String))

	NewDeclarative(t, "ElementUint8AndListStruct").
		Call((&ContainsConstraint{element: uint8(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint8AndListUnsafePointer").
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
	NewDeclarative(t, "ElementUint16AndListBool").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInt").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInt8").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInt16").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInt32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInt64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUint").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUint8").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUint16").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUint32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUint64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUintptr").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListFloat32").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListFloat64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListComplex64").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint16AndListComplex128").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint16AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint16AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint16(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint16AndListChan").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint16AndListFunc").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListInterface").
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

	NewDeclarative(t, "ElementUint16AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint16AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint16(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint16AndListPtr").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint16AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint16(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint16AndListString").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint16(5), reflect.String))

	NewDeclarative(t, "ElementUint16AndListStruct").
		Call((&ContainsConstraint{element: uint16(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint16AndListUnsafePointer").
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
	NewDeclarative(t, "ElementUint32AndListBool").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInt").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInt8").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInt16").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInt32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInt64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUint").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUint8").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUint16").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUint32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUint64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUintptr").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListFloat32").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListFloat64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListComplex64").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint32AndListComplex128").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint32AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint32AndListChan").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint32AndListFunc").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListInterface").
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

	NewDeclarative(t, "ElementUint32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint32(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint32AndListPtr").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint32AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint32AndListString").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint32(5), reflect.String))

	NewDeclarative(t, "ElementUint32AndListStruct").
		Call((&ContainsConstraint{element: uint32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint32AndListUnsafePointer").
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
	NewDeclarative(t, "ElementUint64AndListBool").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInt").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInt8").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInt16").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInt32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInt64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUint").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUint64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUint16").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUint32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUint64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUintptr").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListFloat32").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListFloat64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListComplex64").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint64AndListComplex128").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint64AndListArrayAndContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uint64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint64AndListChan").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUint64AndListFunc").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListInterface").
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

	NewDeclarative(t, "ElementUint64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uint64(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint64AndListPtr").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUint64AndListSliceAndContains").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uint64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUint64AndListString").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uint64(5), reflect.String))

	NewDeclarative(t, "ElementUint64AndListStruct").
		Call((&ContainsConstraint{element: uint64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUint64AndListUnsafePointer").
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
	NewDeclarative(t, "ElementUintptrAndListBool").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInt").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInt8").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInt16").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInt32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInt64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUint").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUint8").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUint16").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUint32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUint64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUintptr").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListFloat32").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListFloat64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListComplex64").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintptrAndListComplex128").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintptrAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintptrAndListArrayAndContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, uintptr(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintptrAndListChan").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUintptrAndListFunc").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListInterface").
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

	NewDeclarative(t, "ElementUintptrAndListMapAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintptrAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": uintptr(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintptrAndListPtr").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUintptrAndListSliceAndContains").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, []interface{}{false, uintptr(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUintptrAndListString").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", uintptr(5), reflect.String))

	NewDeclarative(t, "ElementUintptrAndListStruct").
		Call((&ContainsConstraint{element: uintptr(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUintptrAndListUnsafePointer").
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
	NewDeclarative(t, "ElementFloat32AndListBool").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInt").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInt8").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInt16").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInt32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInt64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUint").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUint8").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUint16").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUint32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUint64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUintptr").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListFloat32").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListFloat64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListComplex64").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat32AndListComplex128").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat32AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat32AndListArrayAndContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, float32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat32AndListChan").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat32AndListFunc").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListInterface").
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

	NewDeclarative(t, "ElementFloat32AndListMapAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat32AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": float32(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat32AndListPtr").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat32AndListSliceAndContains").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, []interface{}{false, float32(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat32AndListString").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", float32(5), reflect.String))

	NewDeclarative(t, "ElementFloat32AndListStruct").
		Call((&ContainsConstraint{element: float32(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat32AndListUnsafePointer").
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
	NewDeclarative(t, "ElementFloat64AndListBool").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInt").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInt8").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInt16").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInt32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInt64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUint").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUint8").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUint16").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUint32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUint64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUintptr").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListFloat32").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListFloat64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListComplex64").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat64AndListComplex128").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat64AndListArrayAndContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, float64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat64AndListChan").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFloat64AndListFunc").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListInterface").
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

	NewDeclarative(t, "ElementFloat64AndListMapAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": float64(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat64AndListPtr").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementFloat64AndListSliceAndContains").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, float64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementFloat64AndListString").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", float64(5), reflect.String))

	NewDeclarative(t, "ElementFloat64AndListStruct").
		Call((&ContainsConstraint{element: float64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFloat64AndListUnsafePointer").
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
	NewDeclarative(t, "ElementComplex64AndListBool").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInt").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInt8").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInt16").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInt32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInt64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUint").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUint8").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUint16").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUint32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUint64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUintptr").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListFloat32").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListFloat64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListComplex64").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex64AndListComplex128").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex64AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex64AndListArrayAndContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, complex64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex64AndListChan").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex64AndListFunc").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListInterface").
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

	NewDeclarative(t, "ElementComplex64AndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex64AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": complex64(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex64AndListPtr").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex64AndListSliceAndContains").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, []interface{}{false, complex64(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex64AndListString").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", complex64(5), reflect.String))

	NewDeclarative(t, "ElementComplex64AndListStruct").
		Call((&ContainsConstraint{element: complex64(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex64AndListUnsafePointer").
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
	NewDeclarative(t, "ElementComplex128AndListBool").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInt").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInt8").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInt16").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInt32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInt64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUint").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUint8").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUint16").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUint32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUint64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUintptr").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListFloat32").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListFloat64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListComplex64").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex128AndListComplex128").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex128AndListArrayAndNotContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex128AndListArrayAndContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, [2]interface{}{false, complex128(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex128AndListChan").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementComplex128AndListFunc").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListInterface").
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

	NewDeclarative(t, "ElementComplex128AndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex128AndListMapAndContains").
		Call(
			(&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": complex128(5)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex128AndListPtr").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListSliceAndNotContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementComplex128AndListSliceAndContains").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, []interface{}{false, complex128(5)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementComplex128AndListString").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", complex128(5), reflect.String))

	NewDeclarative(t, "ElementComplex128AndListStruct").
		Call((&ContainsConstraint{element: complex128(5), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementComplex128AndListUnsafePointer").
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
	NewDeclarative(t, "ElementArrayAndListBool").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInt").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInt8").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInt16").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInt32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInt64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUint").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUint8").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUint16").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUint32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUint64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUintptr").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListFloat32").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListFloat64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListComplex64").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementArrayAndListComplex128").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementArrayAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementArrayAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementArrayAndListChan").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementArrayAndListFunc").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListInterface").
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

	NewDeclarative(t, "ElementArrayAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementArrayAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementArrayAndListPtr").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementArrayAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, [2]interface{}{false, 10}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementArrayAndListString").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", [2]interface{}{false, 10}, reflect.String))

	NewDeclarative(t, "ElementArrayAndListStruct").
		Call((&ContainsConstraint{element: [2]interface{}{false, 10}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementArrayAndListUnsafePointer").
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

	NewDeclarative(t, "ElementChanAndListBool").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInt").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInt8").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInt16").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInt32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInt64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUint").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUint8").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUint16").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUint32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUint64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUintptr").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListFloat32").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListFloat64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListComplex64").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementChanAndListComplex128").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementChanAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementChanAndListArrayAndContains").
		Call((&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check, [2]interface{}{false, chanFixture}).
		ExpectResult(true)

	NewDeclarative(t, "ElementChanAndListChan").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementChanAndListFunc").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListInterface").
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

	NewDeclarative(t, "ElementChanAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementChanAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": chanFixture},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementChanAndListPtr").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementChanAndListSliceAndContains").
		Call((&ContainsConstraint{element: chanFixture, comparator: NewEqualComparator()}).Check, []interface{}{false, chanFixture}).
		ExpectResult(true)

	NewDeclarative(t, "ElementChanAndListString").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", make(chan int), reflect.String))

	NewDeclarative(t, "ElementChanAndListStruct").
		Call((&ContainsConstraint{element: make(chan int), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementChanAndListUnsafePointer").
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

	NewDeclarative(t, "ElementFuncAndListBool").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInt").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInt8").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInt16").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInt32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInt64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUint").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUint8").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUint16").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUint32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUint64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUintptr").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListFloat32").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListFloat64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListComplex64").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFuncAndListComplex128").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFuncAndListArrayAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, 10},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementFuncAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			[2]interface{}{funcFixture, false},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementFuncAndListChan").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementFuncAndListFunc").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListInterface").
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

	NewDeclarative(t, "ElementFuncAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementFuncAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": funcFixture},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementFuncAndListPtr").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListSliceAndNotContains").
		Call(
			(&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementFuncAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: funcFixture, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, funcFixture},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementFuncAndListString").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", func() {}, reflect.String))

	NewDeclarative(t, "ElementFuncAndListStruct").
		Call((&ContainsConstraint{element: func() {}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementFuncAndListUnsafePointer").
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
	NewDeclarative(t, "ElementInterfaceAndListBool").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInt").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInt8").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInt16").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInt32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInt64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUint").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUint8").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUint16").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUint32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUint64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUintptr").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListFloat32").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListFloat64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListComplex64").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInterfaceAndListComplex128").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInterfaceAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInterfaceAndListArrayAndContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, [2]interface{}{false, (*interface{})(nil)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInterfaceAndListChan").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementInterfaceAndListFunc").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListInterface").
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

	NewDeclarative(t, "ElementInterfaceAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementInterfaceAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": (*interface{})(nil)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementInterfaceAndListPtr").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementInterfaceAndListSliceAndContains").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, []interface{}{false, (*interface{})(nil)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementInterfaceAndListString").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", (*interface{})(nil), reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListStruct").
		Call((&ContainsConstraint{element: (*interface{})(nil), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementInterfaceAndListUnsafePointer").
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
	NewDeclarative(t, "ElementMapAndListBool").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInt").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInt8").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInt16").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInt32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInt64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUint").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUint8").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUint16").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUint32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUint64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUintptr").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListFloat32").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListFloat64").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListComplex64").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			complex64(5),
		).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementMapAndListComplex128").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			complex128(5),
		).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementMapAndListArrayAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, 10},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementMapAndListArrayAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[2]interface{}{false, map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementMapAndListChan").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			make(chan uint),
		).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementMapAndListFunc").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListInterface").
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

	NewDeclarative(t, "ElementMapAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementMapAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementMapAndListPtr").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListSliceAndNotContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementMapAndListSliceAndContains").
		Call(
			(&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check,
			[]interface{}{false, map[string]interface{}{"First": false, "Second": false}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementMapAndListString").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(
			NewInvalidKindError("element", map[string]interface{}{"First": false, "Second": false}, reflect.String),
		)

	NewDeclarative(t, "ElementMapAndListStruct").
		Call((&ContainsConstraint{element: map[string]interface{}{"First": false, "Second": false}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementMapAndListUnsafePointer").
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
	NewDeclarative(t, "ElementPtrAndListBool").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInt").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInt8").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInt16").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInt32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInt64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUint").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUint8").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUint16").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUint32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUint64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUintptr").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListFloat32").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListFloat64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListComplex64").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementPtrAndListComplex128").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementPtrAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementPtrAndListArrayAndContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, [2]interface{}{false, new(int)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementPtrAndListChan").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementPtrAndListFunc").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListInterface").
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

	NewDeclarative(t, "ElementPtrAndListMapAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementPtrAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": new(int)},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementPtrAndListPtr").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementPtrAndListSliceAndContains").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, []interface{}{false, new(int)}).
		ExpectResult(true)

	NewDeclarative(t, "ElementPtrAndListString").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", new(int), reflect.String))

	NewDeclarative(t, "ElementPtrAndListStruct").
		Call((&ContainsConstraint{element: new(int), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementPtrAndListUnsafePointer").
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
	NewDeclarative(t, "ElementSliceAndListBool").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInt").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInt8").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInt16").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInt32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInt64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUint").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUint8").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUint16").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUint32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUint64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUintptr").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListFloat32").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListFloat64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListComplex64").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementSliceAndListComplex128").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementSliceAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementSliceAndListArrayAndContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, []interface{}{false}}).
		ExpectResult(true)

	NewDeclarative(t, "ElementSliceAndListChan").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementSliceAndListFunc").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListInterface").
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

	NewDeclarative(t, "ElementSliceAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementSliceAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": []interface{}{false}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementSliceAndListPtr").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementSliceAndListSliceAndContains").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, []interface{}{false, []interface{}{false}}).
		ExpectResult(true)

	NewDeclarative(t, "ElementSliceAndListString").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", []interface{}{false}, reflect.String))

	NewDeclarative(t, "ElementSliceAndListStruct").
		Call((&ContainsConstraint{element: []interface{}{false}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementSliceAndListUnsafePointer").
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
	NewDeclarative(t, "ElementStringAndListBool").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInt").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInt8").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInt16").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInt32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInt64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUint").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUint8").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUint16").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUint32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUint64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUintptr").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListFloat32").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListFloat64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListComplex64").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStringAndListComplex128").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStringAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStringAndListArrayAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, [2]interface{}{false, "data"}).
		ExpectResult(true)

	NewDeclarative(t, "ElementStringAndListChan").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStringAndListFunc").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListInterface").
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

	NewDeclarative(t, "ElementStringAndListMapAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStringAndListMapAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": "data"}).
		ExpectResult(true)

	NewDeclarative(t, "ElementStringAndListPtr").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStringAndListSliceAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, []interface{}{false, "data"}).
		ExpectResult(true)

	NewDeclarative(t, "ElementStringAndListStringAndNotContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, "not contains").
		ExpectResult(false)

	NewDeclarative(t, "ElementStringAndListStringAndContains").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, "some data here").
		ExpectResult(true)

	NewDeclarative(t, "ElementStringAndListStruct").
		Call((&ContainsConstraint{element: "data", comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStringAndListUnsafePointer").
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
	NewDeclarative(t, "ElementStructAndListBool").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInt").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInt8").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInt16").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInt32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInt64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUint").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUint8").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUint16").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUint32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUint64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUintptr").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListFloat32").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListFloat64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListComplex64").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStructAndListComplex128").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStructAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStructAndListArrayAndContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, [2]interface{}{false, struct{}{}}).
		ExpectResult(true)

	NewDeclarative(t, "ElementStructAndListChan").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementStructAndListFunc").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListInterface").
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

	NewDeclarative(t, "ElementStructAndListMapAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, map[string]interface{}{"First": false, "Second": false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStructAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": struct{}{}},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementStructAndListPtr").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementStructAndListSliceAndContains").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, []interface{}{false, struct{}{}}).
		ExpectResult(true)

	NewDeclarative(t, "ElementStructAndListString").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", struct{}{}, reflect.String))

	NewDeclarative(t, "ElementStructAndListStruct").
		Call((&ContainsConstraint{element: struct{}{}, comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementStructAndListUnsafePointer").
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

	NewDeclarative(t, "ElementUnsafePointerAndListBool").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, true).
		ExpectPanic(NewInvalidKindError("list", true, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInt").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int(5)).
		ExpectPanic(NewInvalidKindError("list", int(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInt8").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int8(5)).
		ExpectPanic(NewInvalidKindError("list", int8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInt16").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int16(5)).
		ExpectPanic(NewInvalidKindError("list", int16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInt32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int32(5)).
		ExpectPanic(NewInvalidKindError("list", int32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInt64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, int64(5)).
		ExpectPanic(NewInvalidKindError("list", int64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUint").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint(5)).
		ExpectPanic(NewInvalidKindError("list", uint(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUint8").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint8(5)).
		ExpectPanic(NewInvalidKindError("list", uint8(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUint16").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint16(5)).
		ExpectPanic(NewInvalidKindError("list", uint16(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUint32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint32(5)).
		ExpectPanic(NewInvalidKindError("list", uint32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUint64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uint64(5)).
		ExpectPanic(NewInvalidKindError("list", uint64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUintptr").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, uintptr(5)).
		ExpectPanic(NewInvalidKindError("list", uintptr(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListFloat32").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, float32(5)).
		ExpectPanic(NewInvalidKindError("list", float32(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListFloat64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, float64(5)).
		ExpectPanic(NewInvalidKindError("list", float64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListComplex64").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, complex64(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex64(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUnsafePointerAndListComplex128").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, complex128(5)).
		ExpectPanic(
			NewInvalidKindError("list", complex128(5), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUnsafePointerAndListArrayAndNotContains").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, [2]interface{}{false, 10}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUnsafePointerAndListArrayAndContains").
		Call((&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check, [2]interface{}{false, unsafePointerFixture}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUnsafePointerAndListChan").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, make(chan uint)).
		ExpectPanic(
			NewInvalidKindError("list", make(chan uint), reflect.Array, reflect.Slice, reflect.Map, reflect.String),
		)

	NewDeclarative(t, "ElementUnsafePointerAndListFunc").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, func() {}).
		ExpectPanic(NewInvalidKindError("list", func() {}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListInterface").
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

	NewDeclarative(t, "ElementUnsafePointerAndListMapAndNotContains").
		Call(
			(&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": false},
		).
		ExpectResult(false)

	NewDeclarative(t, "ElementUnsafePointerAndListMapAndContains").
		Call(
			(&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check,
			map[string]interface{}{"First": false, "Second": unsafePointerFixture},
		).
		ExpectResult(true)

	NewDeclarative(t, "ElementUnsafePointerAndListPtr").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, new(int)).
		ExpectPanic(NewInvalidKindError("list", new(int), reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListSliceAndNotContains").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, []interface{}{false, false}).
		ExpectResult(false)

	NewDeclarative(t, "ElementUnsafePointerAndListSliceAndContains").
		Call((&ContainsConstraint{element: unsafePointerFixture, comparator: NewEqualComparator()}).Check, []interface{}{false, unsafePointerFixture}).
		ExpectResult(true)

	NewDeclarative(t, "ElementUnsafePointerAndListString").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, "data").
		ExpectPanic(NewInvalidKindError("element", unsafe.Pointer(new(int)), reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListStruct").
		Call((&ContainsConstraint{element: unsafe.Pointer(new(int)), comparator: NewEqualComparator()}).Check, struct{}{}).
		ExpectPanic(NewInvalidKindError("list", struct{}{}, reflect.Array, reflect.Slice, reflect.Map, reflect.String))

	NewDeclarative(t, "ElementUnsafePointerAndListUnsafePointer").
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
	NewDeclarative(t, "WithPositiveResult").
		Call((&ContainsConstraint{element: "element"}).String).
		ExpectResult(fmt.Sprintf("contain %+v", "element"))
}

func TestContains_Details(t *testing.T) {
	NewDeclarative(t, "WithPositiveResult").
		Call((&ContainsConstraint{element: "element"}).Details, [2]string{"First", "element"}).
		ExpectResult("")
}
