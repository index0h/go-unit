package unit

import (
	"reflect"
	"testing"
	"unsafe"
)

type equalComparatorUseEqualMethodFixture struct {
	callback func(x interface{}) bool
}

func (e equalComparatorUseEqualMethodFixture) Equal(x interface{}) bool {
	return e.callback(x)
}

//noinspection GoRedundantConversion
func TestNewEqualComparator(t *testing.T) {
	NewSubtest(t, "WithoutOptions").
		Call(NewEqualComparator).
		ExpectResult(
			&EqualComparator{
				numericDelta:     0,
				sameType:         false,
				samePointer:      false,
				useEqualMethod:   false,
				ignoreUnexported: []interface{}{},
				ignoresFields:    []IgnoreFieldsOption{},
			},
		)

	NewSubtest(t, "WithNumericDeltaOption").
		Call(NewEqualComparator, NumericDeltaOption{Value: 5}).
		ExpectResult(
			&EqualComparator{
				numericDelta:     5,
				sameType:         false,
				samePointer:      false,
				useEqualMethod:   false,
				ignoreUnexported: []interface{}{},
				ignoresFields:    []IgnoreFieldsOption{},
			},
		)

	NewSubtest(t, "WithInvalidNumericDeltaOption").
		Call(NewEqualComparator, NumericDeltaOption{Value: -5}).
		ExpectPanic(NewErrorf("Variable 'options[0].Value' must be greater or equal to 0, actual: %f", -5.0))

	NewSubtest(t, "WithSameTypeOption").
		Call(NewEqualComparator, SameTypeOption{Value: true}).
		ExpectResult(
			&EqualComparator{
				numericDelta:     0,
				sameType:         true,
				samePointer:      false,
				useEqualMethod:   false,
				ignoreUnexported: []interface{}{},
				ignoresFields:    []IgnoreFieldsOption{},
			},
		)

	NewSubtest(t, "WithSamePointerOption").
		Call(NewEqualComparator, SamePointerOption{Value: true}).
		ExpectResult(
			&EqualComparator{
				numericDelta:     0,
				sameType:         false,
				samePointer:      true,
				useEqualMethod:   false,
				ignoreUnexported: []interface{}{},
				ignoresFields:    []IgnoreFieldsOption{},
			},
		)

	NewSubtest(t, "WithUseEqualMethodOption").
		Call(NewEqualComparator, UseEqualMethodOption{Value: true}).
		ExpectResult(
			&EqualComparator{
				numericDelta:     0,
				sameType:         false,
				samePointer:      false,
				useEqualMethod:   true,
				ignoreUnexported: []interface{}{},
				ignoresFields:    []IgnoreFieldsOption{},
			},
		)

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInvalid").
		Call(
			func() {
				NewEqualComparator(IgnoreUnexportedOption{Value: nil})
			},
		).
		ExpectPanic(NewInvalidKindError("options[0].Value", nil, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndBool").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: false}).
		ExpectPanic(NewInvalidKindError("options[0].Value", false, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInt").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInt8").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int8(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int8(10), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInt16").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int16(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInt32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInt64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUint").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint(10), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUint8").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint8(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint8(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUint16").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint16(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUint32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUint64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUintptr").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uintptr(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uintptr(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndFloat32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: float32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", float32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndFloat64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: float64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", float64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndComplex64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: complex64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", complex64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndComplex128").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: complex128(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", complex128(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndArray").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: [1]int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", [1]int{5}, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndChan").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: make(chan int)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", make(chan int), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndFunc").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: func() {}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", func() {}, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndInterface").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: (*interface{})(nil)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", (*interface{})(nil), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndMap").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: map[int]int{1: 1}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", map[int]int{1: 1}, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndPtr").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: new(int)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", new(int), reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndSlice").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: []int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", []int{5}, reflect.Struct))

	NewSubtest(t, "WithIgnoreUnexportedOptionAndString").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: "data"}).
		ExpectPanic(NewInvalidKindError("options[0].Value", "data", reflect.Struct))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewSubtest(t, "WithIgnoreUnexportedOptionAndStruct").
			Call(NewEqualComparator, IgnoreUnexportedOption{Value: testStruct{}}).
			ExpectResult(
				&EqualComparator{
					numericDelta:     0,
					sameType:         false,
					samePointer:      false,
					useEqualMethod:   false,
					ignoreUnexported: []interface{}{testStruct{}},
					ignoresFields:    []IgnoreFieldsOption{},
				},
			)
	}

	NewSubtest(t, "WithIgnoreUnexportedOptionAndUnsafePointer").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: unsafe.Pointer(new(int))}).
		ExpectPanic(NewInvalidKindError("options[0].Value", unsafe.Pointer(new(int)), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInvalid").
		Call(
			func() {
				NewEqualComparator(IgnoreFieldsOption{Type: nil})
			},
		).
		ExpectPanic(NewInvalidKindError("options[0].Type", nil, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndBool").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: false}).
		ExpectPanic(NewInvalidKindError("options[0].Type", false, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInt").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInt8").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int8(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int8(10), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInt16").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int16(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInt32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInt64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUint").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint(10), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUint8").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint8(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint8(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUint16").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint16(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUint32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUint64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndUintptr").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uintptr(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uintptr(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndFloat32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: float32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", float32(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndFloat64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: float64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", float64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndComplex64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: complex64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", complex64(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndComplex128").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: complex128(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", complex128(5), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndArray").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: [1]int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", [1]int{5}, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndChan").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: make(chan int)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", make(chan int), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndFunc").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: func() {}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", func() {}, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndInterface").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: (*interface{})(nil)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", (*interface{})(nil), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndMap").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: map[int]int{1: 1}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", map[int]int{1: 1}, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndPtr").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: new(int)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", new(int), reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndSlice").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: []int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", []int{5}, reflect.Struct))

	NewSubtest(t, "WithIgnoreFieldsOptionAndString").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: "data"}).
		ExpectPanic(NewInvalidKindError("options[0].Type", "data", reflect.Struct))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewSubtest(t, "WithIgnoreFieldsOptionAndStruct").
			Call(NewEqualComparator, IgnoreFieldsOption{Type: testStruct{}}).
			ExpectResult(
				&EqualComparator{
					numericDelta:     0,
					sameType:         false,
					samePointer:      false,
					useEqualMethod:   false,
					ignoreUnexported: []interface{}{},
					ignoresFields:    []IgnoreFieldsOption{{Type: testStruct{}}},
				},
			)

		NewSubtest(t, "WithIgnoreFieldsOptionAndStructAndUndefinedField").
			Call(NewEqualComparator, IgnoreFieldsOption{Type: testStruct{}, Fields: []string{"undefined"}}).
			ExpectPanic(NewErrorf("Variable 'options[0].Fields[0]' contains unknown field name: 'undefined'"))
	}

	NewSubtest(t, "WithIgnoreFieldsOptionAndUnsafePointer").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: unsafe.Pointer(new(int))}).
		ExpectPanic(NewInvalidKindError("options[0].Type", unsafe.Pointer(new(int)), reflect.Struct))

	NewSubtest(t, "WithOptionInvalid").
		Call(
			func() {
				NewEqualComparator(nil)
			},
		).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", nil))

	NewSubtest(t, "WithOptionBool").
		Call(NewEqualComparator, false).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", false))

	NewSubtest(t, "WithOptionInt").
		Call(NewEqualComparator, int(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int(5)))

	NewSubtest(t, "WithOptionInt8").
		Call(NewEqualComparator, int8(10)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int8(10)))

	NewSubtest(t, "WithOptionInt16").
		Call(NewEqualComparator, int16(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int16(5)))

	NewSubtest(t, "WithOptionInt32").
		Call(NewEqualComparator, int32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int32(5)))

	NewSubtest(t, "WithOptionInt64").
		Call(NewEqualComparator, int64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int64(5)))

	NewSubtest(t, "WithOptionUint").
		Call(NewEqualComparator, uint(10)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint(10)))

	NewSubtest(t, "WithOptionUint8").
		Call(NewEqualComparator, uint8(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint8(5)))

	NewSubtest(t, "WithOptionUint16").
		Call(NewEqualComparator, uint16(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint16(5)))

	NewSubtest(t, "WithOptionUint32").
		Call(NewEqualComparator, uint32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint32(5)))

	NewSubtest(t, "WithOptionUint64").
		Call(NewEqualComparator, uint64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint64(5)))

	NewSubtest(t, "WithOptionUintptr").
		Call(NewEqualComparator, uintptr(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uintptr(5)))

	NewSubtest(t, "WithOptionFloat32").
		Call(NewEqualComparator, float32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", float32(5)))

	NewSubtest(t, "WithOptionFloat64").
		Call(NewEqualComparator, float64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", float64(5)))

	NewSubtest(t, "WithOptionComplex64").
		Call(NewEqualComparator, complex64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", complex64(5)))

	NewSubtest(t, "WithOptionComplex128").
		Call(NewEqualComparator, complex128(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", complex128(5)))

	NewSubtest(t, "WithOptionArray").
		Call(NewEqualComparator, [1]int{5}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", [1]int{5}))

	NewSubtest(t, "WithOptionChan").
		Call(NewEqualComparator, make(chan int)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", make(chan int)))

	NewSubtest(t, "WithOptionFunc").
		Call(NewEqualComparator, func() {}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", func() {}))

	NewSubtest(t, "WithOptionInterface").
		Call(NewEqualComparator, (*interface{})(nil)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", (*interface{})(nil)))

	NewSubtest(t, "WithOptionMap").
		Call(NewEqualComparator, map[int]int{1: 1}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", map[int]int{1: 1}))

	NewSubtest(t, "WithOptionPtr").
		Call(NewEqualComparator, new(int)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", new(int)))

	NewSubtest(t, "WithOptionSlice").
		Call(NewEqualComparator, []int{5}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", []int{5}))

	NewSubtest(t, "WithOptionString").
		Call(NewEqualComparator, "data").
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", "data"))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewSubtest(t, "WithOptionStruct").
			Call(NewEqualComparator, testStruct{}).
			ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", testStruct{}))
	}

	NewSubtest(t, "WithOptionUnsafePointer").
		Call(NewEqualComparator, unsafe.Pointer(new(int))).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", unsafe.Pointer(new(int))))
}

//noinspection GoRedundantConversion
func TestEqualComparator_Compare(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}) (bool, bool) {
		return (&EqualComparator{}).Compare(x, y), (&EqualComparator{}).Compare(y, x)
	}

	doubleRunWithInvalid := func(x interface{}) (bool, bool) {
		return (&EqualComparator{}).Compare(x, nil), (&EqualComparator{}).Compare(nil, x)
	}

	NewSubtest(t, "InvalidAndInvalid").
		Call(
			func() bool {
				return (&EqualComparator{}).Compare(nil, nil)
			},
		).
		ExpectResult(true)

	NewSubtest(t, "InvalidAndBool").
		Call(doubleRunWithInvalid, false).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInt").
		Call(doubleRunWithInvalid, int(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInt8").
		Call(doubleRunWithInvalid, int8(10)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInt16").
		Call(doubleRunWithInvalid, int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInt32").
		Call(doubleRunWithInvalid, int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInt64").
		Call(doubleRunWithInvalid, int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUint").
		Call(doubleRunWithInvalid, uint(10)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUint8").
		Call(doubleRunWithInvalid, uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUint16").
		Call(doubleRunWithInvalid, uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUint32").
		Call(doubleRunWithInvalid, uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUint64").
		Call(doubleRunWithInvalid, uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUintptr").
		Call(doubleRunWithInvalid, uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndFloat32").
		Call(doubleRunWithInvalid, float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndFloat64").
		Call(doubleRunWithInvalid, float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndComplex64").
		Call(doubleRunWithInvalid, complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndComplex128").
		Call(doubleRunWithInvalid, complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndArray").
		Call(doubleRunWithInvalid, [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndChan").
		Call(doubleRunWithInvalid, make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndFunc").
		Call(doubleRunWithInvalid, func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndInterface").
		Call(doubleRunWithInvalid, (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndMap").
		Call(doubleRunWithInvalid, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndPtr").
		Call(doubleRunWithInvalid, new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndSlice").
		Call(doubleRunWithInvalid, []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndString").
		Call(doubleRunWithInvalid, "data").
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndStruct").
		Call(doubleRunWithInvalid, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "InvalidAndUnsafePointer").
		Call(doubleRunWithInvalid, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndBoolAndPositiveResult").
		Call(doubleRun, true, true).
		ExpectResult(true, true)

	NewSubtest(t, "BoolAndBoolAndNegativeResult").
		Call(doubleRun, true, false).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInt").
		Call(doubleRun, true, int(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInt8").
		Call(doubleRun, true, int8(10)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInt16").
		Call(doubleRun, true, int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInt32").
		Call(doubleRun, true, int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInt64").
		Call(doubleRun, true, int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUint").
		Call(doubleRun, true, uint(10)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUint8").
		Call(doubleRun, true, uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUint16").
		Call(doubleRun, true, uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUint32").
		Call(doubleRun, true, uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUint64").
		Call(doubleRun, true, uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUintptr").
		Call(doubleRun, true, uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndFloat32").
		Call(doubleRun, true, float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndFloat64").
		Call(doubleRun, true, float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndComplex64").
		Call(doubleRun, true, complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndComplex128").
		Call(doubleRun, true, complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndArray").
		Call(doubleRun, true, [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndChan").
		Call(doubleRun, true, make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndFunc").
		Call(doubleRun, true, func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndInterface").
		Call(doubleRun, true, (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndMap").
		Call(doubleRun, true, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndPtr").
		Call(doubleRun, true, new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndSlice").
		Call(doubleRun, true, []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndString").
		Call(doubleRun, true, "data").
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndStruct").
		Call(doubleRun, true, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "BoolAndUnsafePointer").
		Call(doubleRun, true, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndIntWithPositiveResult").
		Call(doubleRun, int(5), int(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndIntAndNegativeResult").
		Call(doubleRun, int(10), int(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt8AndNegativeResult").
		Call(doubleRun, int(10), int8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt8AndPositiveResult").
		Call(doubleRun, int(5), int8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt16AndNegativeResult").
		Call(doubleRun, int(10), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt16AndPositiveResult").
		Call(doubleRun, int(5), int16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt32AndNegativeResult").
		Call(doubleRun, int(10), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt32AndPositiveResult").
		Call(doubleRun, int(5), int32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt64AndNegativeResult").
		Call(doubleRun, int(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt64AndPositiveResult").
		Call(doubleRun, int(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint8AndNegativeResult").
		Call(doubleRun, int(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint8AndPositiveResult").
		Call(doubleRun, int(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint16AndNegativeResult").
		Call(doubleRun, int(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint16AndPositiveResult").
		Call(doubleRun, int(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint32AndNegativeResult").
		Call(doubleRun, int(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint32AndPositiveResult").
		Call(doubleRun, int(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUintptr").
		Call(doubleRun, int(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFloat32AndNegativeResult").
		Call(doubleRun, int(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFloat32AndPositiveResult").
		Call(doubleRun, int(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndFloat64AndNegativeResult").
		Call(doubleRun, int(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFloat64AndPositiveResult").
		Call(doubleRun, int(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndComplex64").
		Call(doubleRun, int(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndComplex128").
		Call(doubleRun, int(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndArray").
		Call(doubleRun, int(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndChan").
		Call(doubleRun, int(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFunc").
		Call(doubleRun, int(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInterface").
		Call(doubleRun, int(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndMap").
		Call(doubleRun, int(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndPtr").
		Call(doubleRun, int(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndSlice").
		Call(doubleRun, int(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndString").
		Call(doubleRun, int(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "IntAndStruct").
		Call(doubleRun, int(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUnsafePointer").
		Call(doubleRun, int(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt8AndPositiveResult").
		Call(doubleRun, int8(5), int8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt8AndNegativeResult").
		Call(doubleRun, int8(10), int8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt16AndNegativeResult").
		Call(doubleRun, int8(10), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt16AndPositiveResult").
		Call(doubleRun, int8(5), int16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt32AndNegativeResult").
		Call(doubleRun, int8(10), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt32AndPositiveResult").
		Call(doubleRun, int8(5), int32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt64AndNegativeResult").
		Call(doubleRun, int8(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt64AndPositiveResult").
		Call(doubleRun, int8(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint8AndNegativeResult").
		Call(doubleRun, int8(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint8AndPositiveResult").
		Call(doubleRun, int8(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint16AndNegativeResult").
		Call(doubleRun, int8(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint16AndPositiveResult").
		Call(doubleRun, int8(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint32AndNegativeResult").
		Call(doubleRun, int8(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint32AndPositiveResult").
		Call(doubleRun, int8(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUintptr").
		Call(doubleRun, int8(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFloat32AndNegativeResult").
		Call(doubleRun, int8(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFloat32AndPositiveResult").
		Call(doubleRun, int8(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndFloat64AndNegativeResult").
		Call(doubleRun, int8(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFloat64AndPositiveResult").
		Call(doubleRun, int8(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndComplex64").
		Call(doubleRun, int8(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndComplex128").
		Call(doubleRun, int8(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndArray").
		Call(doubleRun, int8(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndChan").
		Call(doubleRun, int8(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFunc").
		Call(doubleRun, int8(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInterface").
		Call(doubleRun, int8(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndMap").
		Call(doubleRun, int8(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndPtr").
		Call(doubleRun, int8(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndSlice").
		Call(doubleRun, int8(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndString").
		Call(doubleRun, int8(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndStruct").
		Call(doubleRun, int8(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUnsafePointer").
		Call(doubleRun, int8(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(5), int16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt16AndNegativeResult").
		Call(doubleRun, int16(10), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt32AndNegativeResult").
		Call(doubleRun, int16(10), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt32AndPositiveResult").
		Call(doubleRun, int16(5), int32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt64AndNegativeResult").
		Call(doubleRun, int16(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt64AndPositiveResult").
		Call(doubleRun, int16(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint8AndNegativeResult").
		Call(doubleRun, int16(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint8AndPositiveResult").
		Call(doubleRun, int16(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint16AndNegativeResult").
		Call(doubleRun, int16(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint16AndPositiveResult").
		Call(doubleRun, int16(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint32AndNegativeResult").
		Call(doubleRun, int16(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint32AndPositiveResult").
		Call(doubleRun, int16(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUintptr").
		Call(doubleRun, int16(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFloat32AndNegativeResult").
		Call(doubleRun, int16(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFloat32AndPositiveResult").
		Call(doubleRun, int16(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndFloat64AndNegativeResult").
		Call(doubleRun, int16(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFloat64AndPositiveResult").
		Call(doubleRun, int16(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndComplex64").
		Call(doubleRun, int16(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndComplex128").
		Call(doubleRun, int16(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndArray").
		Call(doubleRun, int16(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndChan").
		Call(doubleRun, int16(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFunc").
		Call(doubleRun, int16(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInterface").
		Call(doubleRun, int16(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndMap").
		Call(doubleRun, int16(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndPtr").
		Call(doubleRun, int16(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndSlice").
		Call(doubleRun, int16(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndString").
		Call(doubleRun, int16(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndStruct").
		Call(doubleRun, int16(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUnsafePointer").
		Call(doubleRun, int16(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInt32AndNegativeResult").
		Call(doubleRun, int32(10), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(5), int32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndInt64AndNegativeResult").
		Call(doubleRun, int32(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInt64AndPositiveResult").
		Call(doubleRun, int32(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint8AndNegativeResult").
		Call(doubleRun, int32(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint8AndPositiveResult").
		Call(doubleRun, int32(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint16AndNegativeResult").
		Call(doubleRun, int32(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint16AndPositiveResult").
		Call(doubleRun, int32(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint32AndNegativeResult").
		Call(doubleRun, int32(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint32AndPositiveResult").
		Call(doubleRun, int32(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUintptr").
		Call(doubleRun, int32(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFloat32AndNegativeResult").
		Call(doubleRun, int32(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFloat32AndPositiveResult").
		Call(doubleRun, int32(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndFloat64AndNegativeResult").
		Call(doubleRun, int32(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFloat64AndPositiveResult").
		Call(doubleRun, int32(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndComplex64").
		Call(doubleRun, int32(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndComplex128").
		Call(doubleRun, int32(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndArray").
		Call(doubleRun, int32(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndChan").
		Call(doubleRun, int32(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFunc").
		Call(doubleRun, int32(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInterface").
		Call(doubleRun, int32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndMap").
		Call(doubleRun, int32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndPtr").
		Call(doubleRun, int32(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndSlice").
		Call(doubleRun, int32(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndString").
		Call(doubleRun, int32(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndStruct").
		Call(doubleRun, int32(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUnsafePointer").
		Call(doubleRun, int32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndInt64AndNegativeResult").
		Call(doubleRun, int64(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint8AndNegativeResult").
		Call(doubleRun, int64(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint8AndPositiveResult").
		Call(doubleRun, int64(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint16AndNegativeResult").
		Call(doubleRun, int64(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint16AndPositiveResult").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint32AndNegativeResult").
		Call(doubleRun, int64(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint32AndPositiveResult").
		Call(doubleRun, int64(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUintptr").
		Call(doubleRun, int64(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFloat32AndNegativeResult").
		Call(doubleRun, int64(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFloat32AndPositiveResult").
		Call(doubleRun, int64(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndFloat64AndNegativeResult").
		Call(doubleRun, int64(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFloat64AndPositiveResult").
		Call(doubleRun, int64(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndComplex64").
		Call(doubleRun, int64(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndComplex128").
		Call(doubleRun, int64(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndArray").
		Call(doubleRun, int64(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndChan").
		Call(doubleRun, int64(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFunc").
		Call(doubleRun, int64(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndInterface").
		Call(doubleRun, int64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndMap").
		Call(doubleRun, int64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndPtr").
		Call(doubleRun, int64(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndSlice").
		Call(doubleRun, int64(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndString").
		Call(doubleRun, int64(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndStruct").
		Call(doubleRun, int64(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUnsafePointer").
		Call(doubleRun, int64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUintAndNegativeResult").
		Call(doubleRun, uint(10), uint(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(5), uint(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint8AndNegativeResult").
		Call(doubleRun, uint(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint8AndPositiveResult").
		Call(doubleRun, uint(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint16AndNegativeResult").
		Call(doubleRun, uint(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint16AndPositiveResult").
		Call(doubleRun, uint(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint32AndNegativeResult").
		Call(doubleRun, uint(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint32AndPositiveResult").
		Call(doubleRun, uint(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint64AndNegativeResult").
		Call(doubleRun, uint(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint64AndPositiveResult").
		Call(doubleRun, uint(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUintptr").
		Call(doubleRun, uint(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFloat32AndNegativeResult").
		Call(doubleRun, uint(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFloat32AndPositiveResult").
		Call(doubleRun, uint(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndFloat64AndNegativeResult").
		Call(doubleRun, uint(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFloat64AndPositiveResult").
		Call(doubleRun, uint(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndComplex64").
		Call(doubleRun, uint(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndComplex128").
		Call(doubleRun, uint(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndArray").
		Call(doubleRun, uint(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndChan").
		Call(doubleRun, uint(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFunc").
		Call(doubleRun, uint(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndInterface").
		Call(doubleRun, uint(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndMap").
		Call(doubleRun, uint(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndPtr").
		Call(doubleRun, uint(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndSlice").
		Call(doubleRun, uint(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndString").
		Call(doubleRun, uint(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "UintAndStruct").
		Call(doubleRun, uint(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUnsafePointer").
		Call(doubleRun, uint(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint8AndNegativeResult").
		Call(doubleRun, uint8(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint16AndNegativeResult").
		Call(doubleRun, uint8(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint16AndPositiveResult").
		Call(doubleRun, uint8(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint32AndNegativeResult").
		Call(doubleRun, uint8(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint32AndPositiveResult").
		Call(doubleRun, uint8(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint64AndNegativeResult").
		Call(doubleRun, uint8(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint64AndPositiveResult").
		Call(doubleRun, uint8(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUintptr").
		Call(doubleRun, uint8(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFloat32AndNegativeResult").
		Call(doubleRun, uint8(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFloat32AndPositiveResult").
		Call(doubleRun, uint8(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndFloat64AndNegativeResult").
		Call(doubleRun, uint8(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFloat64AndPositiveResult").
		Call(doubleRun, uint8(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndComplex64").
		Call(doubleRun, uint8(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndComplex128").
		Call(doubleRun, uint8(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndArray").
		Call(doubleRun, uint8(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndChan").
		Call(doubleRun, uint8(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFunc").
		Call(doubleRun, uint8(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndInterface").
		Call(doubleRun, uint8(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndMap").
		Call(doubleRun, uint8(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndPtr").
		Call(doubleRun, uint8(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndSlice").
		Call(doubleRun, uint8(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndString").
		Call(doubleRun, uint8(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndStruct").
		Call(doubleRun, uint8(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUnsafePointer").
		Call(doubleRun, uint8(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint16AndNegativeResult").
		Call(doubleRun, uint16(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint32AndNegativeResult").
		Call(doubleRun, uint16(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint32AndPositiveResult").
		Call(doubleRun, uint16(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint64AndNegativeResult").
		Call(doubleRun, uint16(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint64AndPositiveResult").
		Call(doubleRun, uint16(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUintptr").
		Call(doubleRun, uint16(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFloat32AndNegativeResult").
		Call(doubleRun, uint16(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFloat32AndPositiveResult").
		Call(doubleRun, uint16(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndFloat64AndNegativeResult").
		Call(doubleRun, uint16(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFloat64AndPositiveResult").
		Call(doubleRun, uint16(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndComplex64").
		Call(doubleRun, uint16(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndComplex128").
		Call(doubleRun, uint16(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndArray").
		Call(doubleRun, uint16(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndChan").
		Call(doubleRun, uint16(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFunc").
		Call(doubleRun, uint16(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndInterface").
		Call(doubleRun, uint16(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndMap").
		Call(doubleRun, uint16(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndPtr").
		Call(doubleRun, uint16(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndSlice").
		Call(doubleRun, uint16(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndString").
		Call(doubleRun, uint16(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndStruct").
		Call(doubleRun, uint16(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUnsafePointer").
		Call(doubleRun, uint16(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUint32AndNegativeResult").
		Call(doubleRun, uint32(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndUint64AndNegativeResult").
		Call(doubleRun, uint32(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUint64AndPositiveResult").
		Call(doubleRun, uint32(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndUintptr").
		Call(doubleRun, uint32(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFloat32AndNegativeResult").
		Call(doubleRun, uint32(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFloat32AndPositiveResult").
		Call(doubleRun, uint32(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndFloat64AndNegativeResult").
		Call(doubleRun, uint32(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFloat64AndPositiveResult").
		Call(doubleRun, uint32(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndComplex64").
		Call(doubleRun, uint32(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndComplex128").
		Call(doubleRun, uint32(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndArray").
		Call(doubleRun, uint32(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndChan").
		Call(doubleRun, uint32(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFunc").
		Call(doubleRun, uint32(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndInterface").
		Call(doubleRun, uint32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndMap").
		Call(doubleRun, uint32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndPtr").
		Call(doubleRun, uint32(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndSlice").
		Call(doubleRun, uint32(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndString").
		Call(doubleRun, uint32(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndStruct").
		Call(doubleRun, uint32(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUnsafePointer").
		Call(doubleRun, uint32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndUint64AndNegativeResult").
		Call(doubleRun, uint64(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndUintptr").
		Call(doubleRun, uint64(5), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndFloat32AndNegativeResult").
		Call(doubleRun, uint64(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndFloat32AndPositiveResult").
		Call(doubleRun, uint64(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndFloat64AndNegativeResult").
		Call(doubleRun, uint64(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndFloat64AndPositiveResult").
		Call(doubleRun, uint64(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndComplex64").
		Call(doubleRun, uint64(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndComplex128").
		Call(doubleRun, uint64(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndArray").
		Call(doubleRun, uint64(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndChan").
		Call(doubleRun, uint64(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndFunc").
		Call(doubleRun, uint64(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndInterface").
		Call(doubleRun, uint64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndMap").
		Call(doubleRun, uint64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndPtr").
		Call(doubleRun, uint64(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndSlice").
		Call(doubleRun, uint64(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndString").
		Call(doubleRun, uint64(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndStruct").
		Call(doubleRun, uint64(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndUnsafePointer").
		Call(doubleRun, uint64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndUintptrAndPositiveResult").
		Call(doubleRun, uintptr(5), uintptr(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintptrAndUintptrAndNegativeResult").
		Call(doubleRun, uintptr(5), uintptr(10)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndFloat32").
		Call(doubleRun, uintptr(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndFloat64").
		Call(doubleRun, uintptr(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndComplex64").
		Call(doubleRun, uintptr(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndComplex128").
		Call(doubleRun, uintptr(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndArray").
		Call(doubleRun, uintptr(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndChan").
		Call(doubleRun, uintptr(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndFunc").
		Call(doubleRun, uintptr(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndInterface").
		Call(doubleRun, uintptr(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndMap").
		Call(doubleRun, uintptr(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndPtr").
		Call(doubleRun, uintptr(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndSlice").
		Call(doubleRun, uintptr(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndString").
		Call(doubleRun, uintptr(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndStruct").
		Call(doubleRun, uintptr(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndUnsafePointer").
		Call(doubleRun, uintptr(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFloat32AndNegativeResult").
		Call(doubleRun, float32(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Float32AndFloat64AndNegativeResult").
		Call(doubleRun, float32(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFloat64AndPositiveResult").
		Call(doubleRun, float32(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Float32AndComplex64").
		Call(doubleRun, float32(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndComplex128").
		Call(doubleRun, float32(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndArray").
		Call(doubleRun, float32(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndChan").
		Call(doubleRun, float32(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFunc").
		Call(doubleRun, float32(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndInterface").
		Call(doubleRun, float32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndMap").
		Call(doubleRun, float32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndPtr").
		Call(doubleRun, float32(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndSlice").
		Call(doubleRun, float32(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndString").
		Call(doubleRun, float32(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndStruct").
		Call(doubleRun, float32(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndUnsafePointer").
		Call(doubleRun, float32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndFloat64AndNegativeResult").
		Call(doubleRun, float64(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Float64AndComplex64").
		Call(doubleRun, float64(5), complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndComplex128").
		Call(doubleRun, float64(5), complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndArray").
		Call(doubleRun, float64(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndChan").
		Call(doubleRun, float64(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndFunc").
		Call(doubleRun, float64(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndInterface").
		Call(doubleRun, float64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndMap").
		Call(doubleRun, float64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndPtr").
		Call(doubleRun, float64(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndSlice").
		Call(doubleRun, float64(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndString").
		Call(doubleRun, float64(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndStruct").
		Call(doubleRun, float64(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndUnsafePointer").
		Call(doubleRun, float64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndComplex64AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)+5i).
		ExpectResult(true, true)

	NewSubtest(t, "Complex64AndComplex64AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndComplex128AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewSubtest(t, "Complex64AndComplex128AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndArray").
		Call(doubleRun, complex64(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndChan").
		Call(doubleRun, complex64(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndFunc").
		Call(doubleRun, complex64(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndInterface").
		Call(doubleRun, complex64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndMap").
		Call(doubleRun, complex64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndPtr").
		Call(doubleRun, complex64(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndSlice").
		Call(doubleRun, complex64(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndString").
		Call(doubleRun, complex64(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndStruct").
		Call(doubleRun, complex64(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndUnsafePointer").
		Call(doubleRun, complex64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndComplex128AndPositiveResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewSubtest(t, "Complex128AndComplex128AndNegativeResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndArray").
		Call(doubleRun, complex128(5), [1]int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndChan").
		Call(doubleRun, complex128(5), make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndFunc").
		Call(doubleRun, complex128(5), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndInterface").
		Call(doubleRun, complex128(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndMap").
		Call(doubleRun, complex128(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndPtr").
		Call(doubleRun, complex128(5), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndSlice").
		Call(doubleRun, complex128(5), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndString").
		Call(doubleRun, complex128(5), "data").
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndStruct").
		Call(doubleRun, complex128(5), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndUnsafePointer").
		Call(doubleRun, complex128(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndPositiveResult").
		Call(doubleRun, [2]string{"data", "string"}, [2]string{"data", "string"}).
		ExpectResult(true, true)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByType").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByElementValue").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByElementsCount").
		Call(doubleRun, [2]string{"data", "string"}, [3]interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndChan").
		Call(doubleRun, [1]string{"data"}, make(chan int)).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndFunc").
		Call(doubleRun, [1]string{"data"}, func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndInterface").
		Call(doubleRun, [1]string{"data"}, (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndMap").
		Call(doubleRun, [1]string{"data"}, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndPtr").
		Call(doubleRun, [1]string{"data"}, new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndSlice").
		Call(doubleRun, [1]string{"data"}, []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndString").
		Call(doubleRun, [1]string{"data"}, "data").
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndStruct").
		Call(doubleRun, [1]string{"data"}, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndUnsafePointer").
		Call(doubleRun, [1]string{"data"}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndChanAndNegativeResult").
		Call(doubleRun, make(chan int), make(chan int)).
		ExpectResult(false, false)

	{
		chanFixture := make(chan int)

		NewSubtest(t, "ChanAndChanAndNegativeResult").
			Call(doubleRun, chanFixture, chanFixture).
			ExpectResult(true, true)
	}

	NewSubtest(t, "ChanAndFunc").
		Call(doubleRun, make(chan int), func() {}).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndInterface").
		Call(doubleRun, make(chan int), (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndMap").
		Call(doubleRun, make(chan int), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndPtr").
		Call(doubleRun, make(chan int), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndSlice").
		Call(doubleRun, make(chan int), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndString").
		Call(doubleRun, make(chan int), "data").
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndStruct").
		Call(doubleRun, make(chan int), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndUnsafePointer").
		Call(doubleRun, make(chan int), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndFuncWithNegativeResult").
		Call(doubleRun, func() {}, func() {}).
		ExpectResult(false, false)

	{
		funcFixture := func() {}

		NewSubtest(t, "FuncAndFuncWithPositiveResult").
			Call(doubleRun, funcFixture, funcFixture).
			ExpectResult(true, true)
	}

	NewSubtest(t, "FuncAndInterface").
		Call(doubleRun, func() {}, (*interface{})(nil)).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndMap").
		Call(doubleRun, func() {}, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndPtr").
		Call(doubleRun, func() {}, new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndSlice").
		Call(doubleRun, func() {}, []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndString").
		Call(doubleRun, func() {}, "data").
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndStruct").
		Call(doubleRun, func() {}, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "FuncAndUnsafePointer").
		Call(doubleRun, func() {}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndInterface").
		Call(doubleRun, (*interface{})(nil), (*interface{})(nil)).
		ExpectResult(true, true)

	NewSubtest(t, "InterfaceAndMap").
		Call(doubleRun, (*interface{})(nil), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndPtr").
		Call(doubleRun, (*interface{})(nil), new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndSlice").
		Call(doubleRun, (*interface{})(nil), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndString").
		Call(doubleRun, (*interface{})(nil), "data").
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndStruct").
		Call(doubleRun, (*interface{})(nil), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "InterfaceAndUnsafePointer").
		Call(doubleRun, (*interface{})(nil), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndMapWithPositiveResult").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2},
		).
		ExpectResult(true, true)

	NewSubtest(t, "MapAndMapWithNegativeResultByType").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]interface{}{"First": 1, "Second": 2},
		).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndMapWithNegativeResultByElementValue").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2000},
		).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndMapWithNegativeResultByElementCount").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2, "third": 3},
		).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndPtr").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, new(int)).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndSlice").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndString").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, "data").
		ExpectResult(false, false)

	NewSubtest(t, "MapAndStruct").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndUnsafePointer").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	{
		stringFixture1 := "First"
		stringFixture2 := "First"
		stringFixture3 := "data"

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(true, true)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	NewSubtest(t, "PtrAndSlice").
		Call(doubleRun, new(string), []int{5}).
		ExpectResult(false, false)

	NewSubtest(t, "PtrAndString").
		Call(doubleRun, new(string), "data").
		ExpectResult(false, false)

	NewSubtest(t, "PtrAndStruct").
		Call(doubleRun, new(string), struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "PtrAndUnsafePointer").
		Call(doubleRun, new(string), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndSliceAndPositiveResult").
		Call(doubleRun, []string{"data", "string"}, []string{"data", "string"}).
		ExpectResult(true, true)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByType").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByElementValue").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByElementsCount").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndString").
		Call(doubleRun, []string{"First", "Second"}, "data").
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndStruct").
		Call(doubleRun, []string{"First", "Second"}, struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndUnsafePointer").
		Call(doubleRun, []string{"First", "Second"}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewSubtest(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", "data").
		ExpectResult(true, true)

	NewSubtest(t, "StringAndStringWithNegativeResult").
		Call(doubleRun, "data", "First").
		ExpectResult(false, false)

	NewSubtest(t, "StringAndStruct").
		Call(doubleRun, "data", struct{}{}).
		ExpectResult(false, false)

	NewSubtest(t, "StringAndUnsafePointer").
		Call(doubleRun, "data", unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	{
		structFixture1 := struct {
			field string
		}{field: "data"}

		structFixture2 := struct {
			field string
		}{field: "secondData"}

		structFixture3 := struct {
			secondField string
		}{secondField: "data"}

		NewSubtest(t, "StructAndStructWithPositiveResult").
			Call(doubleRun, structFixture1, structFixture1).
			ExpectResult(true, true)

		NewSubtest(t, "StructAndStructWithNegativeResultByFieldValue").
			Call(doubleRun, structFixture1, structFixture2).
			ExpectResult(false, false)

		NewSubtest(t, "StructAndStructWithNegativeResultByFieldName").
			Call(doubleRun, structFixture1, structFixture3).
			ExpectResult(false, false)
	}

	NewSubtest(t, "StructAndUnsafePointer").
		Call(doubleRun, struct{}{}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
		Call(doubleRun, unsafe.Pointer(new(int)), unsafe.Pointer(new(int))).
		ExpectResult(false, false)
}

//noinspection GoRedundantConversion
func TestEqualComparator_Compare_WithDelta(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}, delta float64) (bool, bool) {
		return (&EqualComparator{numericDelta: delta}).Compare(x, y),
			(&EqualComparator{numericDelta: delta}).Compare(y, x)
	}

	NewSubtest(t, "IntAndIntAndPositiveResult").
		Call(doubleRun, int(10), int(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt8AndPositiveResult").
		Call(doubleRun, int(10), int8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt16AndPositiveResult").
		Call(doubleRun, int(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt32AndPositiveResult").
		Call(doubleRun, int(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndInt64AndPositiveResult").
		Call(doubleRun, int(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint8AndPositiveResult").
		Call(doubleRun, int(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint16AndPositiveResult").
		Call(doubleRun, int(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndUint32AndPositiveResult").
		Call(doubleRun, int(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndFloat32AndPositiveResult").
		Call(doubleRun, int(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndFloat64AndPositiveResult").
		Call(doubleRun, int(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt8AndPositiveResult").
		Call(doubleRun, int8(10), int8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt16AndPositiveResult").
		Call(doubleRun, int8(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt32AndPositiveResult").
		Call(doubleRun, int8(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt64AndPositiveResult").
		Call(doubleRun, int8(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint8AndPositiveResult").
		Call(doubleRun, int8(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint16AndPositiveResult").
		Call(doubleRun, int8(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndUint32AndPositiveResult").
		Call(doubleRun, int8(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndFloat32AndPositiveResult").
		Call(doubleRun, int8(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndFloat64AndPositiveResult").
		Call(doubleRun, int8(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt32AndPositiveResult").
		Call(doubleRun, int16(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt64AndPositiveResult").
		Call(doubleRun, int16(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint8AndPositiveResult").
		Call(doubleRun, int16(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint16AndPositiveResult").
		Call(doubleRun, int16(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndUint32AndPositiveResult").
		Call(doubleRun, int16(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndFloat32AndPositiveResult").
		Call(doubleRun, int16(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndFloat64AndPositiveResult").
		Call(doubleRun, int16(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndInt64AndPositiveResult").
		Call(doubleRun, int32(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint8AndPositiveResult").
		Call(doubleRun, int32(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint16AndPositiveResult").
		Call(doubleRun, int32(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndUint32AndPositiveResult").
		Call(doubleRun, int32(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndFloat32AndPositiveResult").
		Call(doubleRun, int32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndFloat64AndPositiveResult").
		Call(doubleRun, int32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint8AndPositiveResult").
		Call(doubleRun, int64(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint16AndPositiveResult").
		Call(doubleRun, int64(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint32AndPositiveResult").
		Call(doubleRun, int64(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndFloat32AndPositiveResult").
		Call(doubleRun, int64(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndFloat64AndPositiveResult").
		Call(doubleRun, int64(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(10), uint(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint8AndPositiveResult").
		Call(doubleRun, uint(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint16AndPositiveResult").
		Call(doubleRun, uint(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint32AndPositiveResult").
		Call(doubleRun, uint(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint64AndPositiveResult").
		Call(doubleRun, uint(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndFloat32AndPositiveResult").
		Call(doubleRun, uint(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndFloat64AndPositiveResult").
		Call(doubleRun, uint(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint16AndPositiveResult").
		Call(doubleRun, uint8(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint32AndPositiveResult").
		Call(doubleRun, uint8(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint64AndPositiveResult").
		Call(doubleRun, uint8(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndFloat32AndPositiveResult").
		Call(doubleRun, uint8(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndFloat64AndPositiveResult").
		Call(doubleRun, uint8(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint32AndPositiveResult").
		Call(doubleRun, uint16(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint64AndPositiveResult").
		Call(doubleRun, uint16(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndFloat32AndPositiveResult").
		Call(doubleRun, uint16(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndFloat64AndPositiveResult").
		Call(doubleRun, uint16(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndUint64AndPositiveResult").
		Call(doubleRun, uint32(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndFloat32AndPositiveResult").
		Call(doubleRun, uint32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndFloat64AndPositiveResult").
		Call(doubleRun, uint32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndFloat32AndPositiveResult").
		Call(doubleRun, uint64(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndFloat64AndPositiveResult").
		Call(doubleRun, uint64(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Float32AndFloat64AndPositiveResult").
		Call(doubleRun, float32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewSubtest(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(10), float64(5), float64(20)).
		ExpectResult(true, true)
}

//noinspection GoRedundantConversion
func TestEqualComparator_Compare_WithSameType(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}) (bool, bool) {
		return (&EqualComparator{sameType: true}).Compare(x, y), (&EqualComparator{sameType: true}).Compare(y, x)
	}

	NewSubtest(t, "InvalidAndInvalid").
		Call(
			func() bool {
				return (&EqualComparator{sameType: true}).Compare(nil, nil)
			},
		).
		ExpectResult(true)

	NewSubtest(t, "BoolAndBoolAndPositiveResult").
		Call(doubleRun, true, true).
		ExpectResult(true, true)

	NewSubtest(t, "BoolAndBoolAndNegativeResult").
		Call(doubleRun, true, false).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndIntWithPositiveResult").
		Call(doubleRun, int(5), int(5)).
		ExpectResult(true, true)

	NewSubtest(t, "IntAndIntAndNegativeResult").
		Call(doubleRun, int(10), int(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt8").
		Call(doubleRun, int(5), int8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt16").
		Call(doubleRun, int(5), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt32").
		Call(doubleRun, int(5), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndInt64").
		Call(doubleRun, int(5), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint8").
		Call(doubleRun, int(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint").
		Call(doubleRun, int(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndUint32").
		Call(doubleRun, int(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFloat32").
		Call(doubleRun, int(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "IntAndFloat64").
		Call(doubleRun, int(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt8WithPositiveResult").
		Call(doubleRun, int8(5), int8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int8AndInt8WithNegativeResult").
		Call(doubleRun, int8(10), int8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt16").
		Call(doubleRun, int8(5), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt32").
		Call(doubleRun, int8(5), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndInt64").
		Call(doubleRun, int8(5), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint8").
		Call(doubleRun, int8(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint16").
		Call(doubleRun, int8(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndUint32").
		Call(doubleRun, int8(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFloat32").
		Call(doubleRun, int8(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int8AndFloat64").
		Call(doubleRun, int8(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(5), int16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int16AndInt16AndNegativeResult").
		Call(doubleRun, int16(10), int16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt32").
		Call(doubleRun, int16(5), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndInt64").
		Call(doubleRun, int16(5), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint8").
		Call(doubleRun, int16(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint16").
		Call(doubleRun, int16(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndUint32").
		Call(doubleRun, int16(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFloat32").
		Call(doubleRun, int16(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int16AndFloat64").
		Call(doubleRun, int16(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInt32AndNegativeResult").
		Call(doubleRun, int32(10), int32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(5), int32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int32AndInt64").
		Call(doubleRun, int32(5), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint8").
		Call(doubleRun, int32(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint16").
		Call(doubleRun, int32(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndUint32").
		Call(doubleRun, int32(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFloat32").
		Call(doubleRun, int32(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int32AndFloat64").
		Call(doubleRun, int32(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndInt64AndNegativeResult").
		Call(doubleRun, int64(10), int64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(5), int64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Int64AndUint8").
		Call(doubleRun, int64(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint16AndNegativeResult").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint16").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndUint32").
		Call(doubleRun, int64(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFloat32").
		Call(doubleRun, int64(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Int64AndFloat64").
		Call(doubleRun, int64(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUintAndNegativeResult").
		Call(doubleRun, uint(10), uint(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(5), uint(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintAndUint8").
		Call(doubleRun, uint(5), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint16").
		Call(doubleRun, uint(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint32").
		Call(doubleRun, uint(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndUint64").
		Call(doubleRun, uint(5), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFloat32").
		Call(doubleRun, uint(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintAndFloat64").
		Call(doubleRun, uint(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint8AndNegativeResult").
		Call(doubleRun, uint8(10), uint8(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(5), uint8(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint8AndUint16").
		Call(doubleRun, uint8(5), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint32").
		Call(doubleRun, uint8(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndUint64").
		Call(doubleRun, uint8(5), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFloat32").
		Call(doubleRun, uint8(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint8AndFloat64").
		Call(doubleRun, uint8(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint16AndNegativeResult").
		Call(doubleRun, uint16(10), uint16(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(5), uint16(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint16AndUint32").
		Call(doubleRun, uint16(5), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndUint64").
		Call(doubleRun, uint16(5), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFloat32").
		Call(doubleRun, uint16(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint16AndFloat64").
		Call(doubleRun, uint16(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUint32AndNegativeResult").
		Call(doubleRun, uint32(10), uint32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(5), uint32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint32AndUint64").
		Call(doubleRun, uint32(5), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFloat32").
		Call(doubleRun, uint32(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint32AndFloat64").
		Call(doubleRun, uint32(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndUint64AndNegativeResult").
		Call(doubleRun, uint64(10), uint64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(5), uint64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Uint64AndFloat32").
		Call(doubleRun, uint64(5), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Uint64AndFloat64").
		Call(doubleRun, uint64(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "UintptrAndUintptrAndPositiveResult").
		Call(doubleRun, uintptr(5), uintptr(5)).
		ExpectResult(true, true)

	NewSubtest(t, "UintptrAndUintptrAndNegativeResult").
		Call(doubleRun, uintptr(10), uintptr(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFloat32AndNegativeResult").
		Call(doubleRun, float32(10), float32(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(5), float32(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Float32AndFloat64").
		Call(doubleRun, float32(5), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndFloat64AndNegativeResult").
		Call(doubleRun, float64(10), float64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(5), float64(5)).
		ExpectResult(true, true)

	NewSubtest(t, "Complex64AndComplex64AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)+5i).
		ExpectResult(true, true)

	NewSubtest(t, "Complex64AndComplex64AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)).
		ExpectResult(false, false)

	NewSubtest(t, "Complex64AndComplex128").
		Call(doubleRun, complex64(5)+5i, complex128(5)+5i).
		ExpectResult(false, false)

	NewSubtest(t, "Complex128AndComplex128AndPositiveResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewSubtest(t, "Complex128AndComplex128AndNegativeResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndPositiveResult").
		Call(doubleRun, [2]string{"data", "string"}, [2]string{"data", "string"}).
		ExpectResult(true, true)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByType").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByElementValue").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "ArrayAndArrayAndNegativeResultByElementsCount").
		Call(doubleRun, [2]string{"data", "string"}, [3]interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewSubtest(t, "ChanAndChanAndNegativeResult").
		Call(doubleRun, make(chan int), make(chan int)).
		ExpectResult(false, false)

	{
		chanFixture := make(chan int)

		NewSubtest(t, "ChanAndChanAndNegativeResult").
			Call(doubleRun, chanFixture, chanFixture).
			ExpectResult(true, true)
	}

	NewSubtest(t, "FuncAndFuncWithNegativeResult").
		Call(doubleRun, func() {}, func() {}).
		ExpectResult(false, false)

	{
		funcFixture := func() {}

		NewSubtest(t, "FuncAndFuncWithPositiveResult").
			Call(doubleRun, funcFixture, funcFixture).
			ExpectResult(true, true)
	}

	NewSubtest(t, "InterfaceAndInterface").
		Call(doubleRun, (*interface{})(nil), (*interface{})(nil)).
		ExpectResult(true, true)

	NewSubtest(t, "MapAndMapWithPositiveResult").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2},
		).
		ExpectResult(true, true)

	NewSubtest(t, "MapAndMapWithNegativeResultByType").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]interface{}{"First": 1, "Second": 2},
		).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndMapWithNegativeResultByElementValue").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2000},
		).
		ExpectResult(false, false)

	NewSubtest(t, "MapAndMapWithNegativeResultByElementCount").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2, "third": 3},
		).
		ExpectResult(false, false)

	{
		stringFixture1 := "First"
		stringFixture2 := "First"
		stringFixture3 := "data"

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(true, true)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	NewSubtest(t, "SliceAndSliceAndPositiveResult").
		Call(doubleRun, []string{"data", "string"}, []string{"data", "string"}).
		ExpectResult(true, true)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByType").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByElementValue").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewSubtest(t, "SliceAndSliceAndNegativeResultByElementsCount").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewSubtest(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", "data").
		ExpectResult(true, true)

	NewSubtest(t, "StringAndStringWithNegativeResult").
		Call(doubleRun, "data", "First").
		ExpectResult(false, false)

	{
		structFixture1 := struct {
			field string
		}{field: "data"}

		structFixture2 := struct {
			field string
		}{field: "secondData"}

		structFixture3 := struct {
			secondField string
		}{secondField: "data"}

		NewSubtest(t, "StructAndStructWithPositiveResult").
			Call(doubleRun, structFixture1, structFixture1).
			ExpectResult(true, true)

		NewSubtest(t, "StructAndStructWithNegativeResultByFieldValue").
			Call(doubleRun, structFixture1, structFixture2).
			ExpectResult(false, false)

		NewSubtest(t, "StructAndStructWithNegativeResultByFieldName").
			Call(doubleRun, structFixture1, structFixture3).
			ExpectResult(false, false)
	}

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
		Call(doubleRun, unsafe.Pointer(new(string)), unsafe.Pointer(new(int))).
		ExpectResult(false, false)
}

func TestEqualComparator_Compare_WithSamePointer(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}) (bool, bool) {
		return (&EqualComparator{samePointer: true}).Compare(x, y), (&EqualComparator{samePointer: true}).Compare(y, x)
	}

	{
		stringFixture1 := "First"
		stringFixture2 := "First"
		stringFixture3 := "data"

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewSubtest(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(false, false)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewSubtest(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewSubtest(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
		Call(doubleRun, unsafe.Pointer(new(string)), unsafe.Pointer(new(int))).
		ExpectResult(false, false)
}

func TestEqualComparator_Compare_WithUseEqualMethod(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}) (bool, bool) {
		return (&EqualComparator{useEqualMethod: true}).Compare(x, y),
			(&EqualComparator{useEqualMethod: true}).Compare(y, x)
	}

	equalFixture := func(y interface{}) *equalComparatorUseEqualMethodFixture {
		return &equalComparatorUseEqualMethodFixture{
			callback: func(x interface{}) bool {
				return x == y
			},
		}
	}

	NewSubtest(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", equalFixture("data")).
		ExpectResult(true, true)

	NewSubtest(t, "StringAndStringWithNegativeResult").
		Call(doubleRun, "data", equalFixture("another data")).
		ExpectResult(false, false)
}

func TestEqualComparator_Compare_WithIgnoreUnexported(t *testing.T) {
	type testType struct {
		Exported   string
		unexported string
	}

	doubleRun := func(x interface{}, y interface{}, ignoreUnexported ...interface{}) (bool, bool) {
		return (&EqualComparator{ignoreUnexported: ignoreUnexported}).Compare(x, y),
			(&EqualComparator{ignoreUnexported: ignoreUnexported}).Compare(y, x)
	}

	NewSubtest(t, "IgnoreUnexportedDisabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "Second"},
			&testType{Exported: "First", unexported: "Second"},
		).
		ExpectResult(true, true)

	NewSubtest(t, "IgnoreUnexportedDisabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "First"},
			&testType{Exported: "First", unexported: "Second"},
		).
		ExpectResult(false, false)

	NewSubtest(t, "IgnoreUnexportedEnabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "First"},
			&testType{Exported: "First", unexported: "Second"},
			testType{},
		).
		ExpectResult(true, true)

	NewSubtest(t, "IgnoreUnexportedEnabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "First"},
			&testType{Exported: "Second", unexported: "Second"},
			testType{},
		).
		ExpectResult(false, false)
}

func TestEqualComparator_Compare_WithIgnoreFields(t *testing.T) {
	type testType struct {
		First  string
		Second string
	}

	doubleRun := func(x interface{}, y interface{}, options ...IgnoreFieldsOption) (bool, bool) {
		return (&EqualComparator{ignoresFields: options}).Compare(x, y),
			(&EqualComparator{ignoresFields: options}).Compare(y, x)
	}

	NewSubtest(t, "IgnoreFieldsDisabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "second"},
			&testType{First: "first", Second: "second"},
		).
		ExpectResult(true, true)

	NewSubtest(t, "IgnoreFieldsDisabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "first", Second: "second"},
		).
		ExpectResult(false, false)

	NewSubtest(t, "IgnoreFieldsEnabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "first", Second: "second"},
			IgnoreFieldsOption{Type: testType{}, Fields: []string{"Second"}},
		).
		ExpectResult(true, true)

	NewSubtest(t, "IgnoreFieldsEnabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "second", Second: "second"},
			IgnoreFieldsOption{Type: testType{}, Fields: []string{"Second"}},
		).
		ExpectResult(false, false)
}
