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
	NewDeclarative(t, "WithoutOptions").
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

	NewDeclarative(t, "WithNumericDeltaOption").
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

	NewDeclarative(t, "WithInvalidNumericDeltaOption").
		Call(NewEqualComparator, NumericDeltaOption{Value: -5}).
		ExpectPanic(NewErrorf("Variable 'options[0].Value' must be greater or equal to 0, actual: %f", -5.0))

	NewDeclarative(t, "WithSameTypeOption").
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

	NewDeclarative(t, "WithSamePointerOption").
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

	NewDeclarative(t, "WithUseEqualMethodOption").
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

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInvalid").
		Call(
			func() {
				NewEqualComparator(IgnoreUnexportedOption{Value: nil})
			},
		).
		ExpectPanic(NewInvalidKindError("options[0].Value", nil, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndBool").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: false}).
		ExpectPanic(NewInvalidKindError("options[0].Value", false, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInt").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInt8").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int8(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int8(10), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInt16").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int16(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInt32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInt64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: int64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", int64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUint").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint(10), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUint8").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint8(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint8(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUint16").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint16(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUint32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUint64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uint64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uint64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUintptr").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: uintptr(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", uintptr(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndFloat32").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: float32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", float32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndFloat64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: float64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", float64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndComplex64").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: complex64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", complex64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndComplex128").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: complex128(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", complex128(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndArray").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: [1]int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", [1]int{5}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndChan").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: make(chan int)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", make(chan int), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndFunc").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: func() {}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", func() {}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndInterface").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: (*interface{})(nil)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", (*interface{})(nil), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndMap").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: map[int]int{1: 1}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", map[int]int{1: 1}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndPtr").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: new(int)}).
		ExpectPanic(NewInvalidKindError("options[0].Value", new(int), reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndSlice").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: []int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Value", []int{5}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndString").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: "data"}).
		ExpectPanic(NewInvalidKindError("options[0].Value", "data", reflect.Struct))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewDeclarative(t, "WithIgnoreUnexportedOptionAndStruct").
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

	NewDeclarative(t, "WithIgnoreUnexportedOptionAndUnsafePointer").
		Call(NewEqualComparator, IgnoreUnexportedOption{Value: unsafe.Pointer(new(int))}).
		ExpectPanic(NewInvalidKindError("options[0].Value", unsafe.Pointer(new(int)), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInvalid").
		Call(
			func() {
				NewEqualComparator(IgnoreFieldsOption{Type: nil})
			},
		).
		ExpectPanic(NewInvalidKindError("options[0].Type", nil, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndBool").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: false}).
		ExpectPanic(NewInvalidKindError("options[0].Type", false, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInt").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInt8").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int8(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int8(10), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInt16").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int16(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInt32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInt64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: int64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", int64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUint").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint(10)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint(10), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUint8").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint8(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint8(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUint16").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint16(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint16(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUint32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUint64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uint64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uint64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUintptr").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: uintptr(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", uintptr(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndFloat32").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: float32(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", float32(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndFloat64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: float64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", float64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndComplex64").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: complex64(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", complex64(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndComplex128").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: complex128(5)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", complex128(5), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndArray").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: [1]int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", [1]int{5}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndChan").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: make(chan int)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", make(chan int), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndFunc").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: func() {}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", func() {}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndInterface").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: (*interface{})(nil)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", (*interface{})(nil), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndMap").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: map[int]int{1: 1}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", map[int]int{1: 1}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndPtr").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: new(int)}).
		ExpectPanic(NewInvalidKindError("options[0].Type", new(int), reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndSlice").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: []int{5}}).
		ExpectPanic(NewInvalidKindError("options[0].Type", []int{5}, reflect.Struct))

	NewDeclarative(t, "WithIgnoreFieldsOptionAndString").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: "data"}).
		ExpectPanic(NewInvalidKindError("options[0].Type", "data", reflect.Struct))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewDeclarative(t, "WithIgnoreFieldsOptionAndStruct").
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

		NewDeclarative(t, "WithIgnoreFieldsOptionAndStructAndUndefinedField").
			Call(NewEqualComparator, IgnoreFieldsOption{Type: testStruct{}, Fields: []string{"undefined"}}).
			ExpectPanic(NewErrorf("Variable 'options[0].Fields[0]' contains unknown field name: 'undefined'"))
	}

	NewDeclarative(t, "WithIgnoreFieldsOptionAndUnsafePointer").
		Call(NewEqualComparator, IgnoreFieldsOption{Type: unsafe.Pointer(new(int))}).
		ExpectPanic(NewInvalidKindError("options[0].Type", unsafe.Pointer(new(int)), reflect.Struct))

	NewDeclarative(t, "WithOptionInvalid").
		Call(
			func() {
				NewEqualComparator(nil)
			},
		).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", nil))

	NewDeclarative(t, "WithOptionBool").
		Call(NewEqualComparator, false).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", false))

	NewDeclarative(t, "WithOptionInt").
		Call(NewEqualComparator, int(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int(5)))

	NewDeclarative(t, "WithOptionInt8").
		Call(NewEqualComparator, int8(10)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int8(10)))

	NewDeclarative(t, "WithOptionInt16").
		Call(NewEqualComparator, int16(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int16(5)))

	NewDeclarative(t, "WithOptionInt32").
		Call(NewEqualComparator, int32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int32(5)))

	NewDeclarative(t, "WithOptionInt64").
		Call(NewEqualComparator, int64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", int64(5)))

	NewDeclarative(t, "WithOptionUint").
		Call(NewEqualComparator, uint(10)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint(10)))

	NewDeclarative(t, "WithOptionUint8").
		Call(NewEqualComparator, uint8(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint8(5)))

	NewDeclarative(t, "WithOptionUint16").
		Call(NewEqualComparator, uint16(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint16(5)))

	NewDeclarative(t, "WithOptionUint32").
		Call(NewEqualComparator, uint32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint32(5)))

	NewDeclarative(t, "WithOptionUint64").
		Call(NewEqualComparator, uint64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uint64(5)))

	NewDeclarative(t, "WithOptionUintptr").
		Call(NewEqualComparator, uintptr(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", uintptr(5)))

	NewDeclarative(t, "WithOptionFloat32").
		Call(NewEqualComparator, float32(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", float32(5)))

	NewDeclarative(t, "WithOptionFloat64").
		Call(NewEqualComparator, float64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", float64(5)))

	NewDeclarative(t, "WithOptionComplex64").
		Call(NewEqualComparator, complex64(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", complex64(5)))

	NewDeclarative(t, "WithOptionComplex128").
		Call(NewEqualComparator, complex128(5)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", complex128(5)))

	NewDeclarative(t, "WithOptionArray").
		Call(NewEqualComparator, [1]int{5}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", [1]int{5}))

	NewDeclarative(t, "WithOptionChan").
		Call(NewEqualComparator, make(chan int)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", make(chan int)))

	NewDeclarative(t, "WithOptionFunc").
		Call(NewEqualComparator, func() {}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", func() {}))

	NewDeclarative(t, "WithOptionInterface").
		Call(NewEqualComparator, (*interface{})(nil)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", (*interface{})(nil)))

	NewDeclarative(t, "WithOptionMap").
		Call(NewEqualComparator, map[int]int{1: 1}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", map[int]int{1: 1}))

	NewDeclarative(t, "WithOptionPtr").
		Call(NewEqualComparator, new(int)).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", new(int)))

	NewDeclarative(t, "WithOptionSlice").
		Call(NewEqualComparator, []int{5}).
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", []int{5}))

	NewDeclarative(t, "WithOptionString").
		Call(NewEqualComparator, "data").
		ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", "data"))

	{
		type testStruct struct {
			First  string
			second string
		}

		NewDeclarative(t, "WithOptionStruct").
			Call(NewEqualComparator, testStruct{}).
			ExpectPanic(NewErrorf("Variable 'options[0]' has unknown type: %T", testStruct{}))
	}

	NewDeclarative(t, "WithOptionUnsafePointer").
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

	NewDeclarative(t, "InvalidAndInvalid").
		Call(
			func() bool {
				return (&EqualComparator{}).Compare(nil, nil)
			},
		).
		ExpectResult(true)

	NewDeclarative(t, "InvalidAndBool").
		Call(doubleRunWithInvalid, false).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInt").
		Call(doubleRunWithInvalid, int(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInt8").
		Call(doubleRunWithInvalid, int8(10)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInt16").
		Call(doubleRunWithInvalid, int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInt32").
		Call(doubleRunWithInvalid, int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInt64").
		Call(doubleRunWithInvalid, int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUint").
		Call(doubleRunWithInvalid, uint(10)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUint8").
		Call(doubleRunWithInvalid, uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUint16").
		Call(doubleRunWithInvalid, uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUint32").
		Call(doubleRunWithInvalid, uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUint64").
		Call(doubleRunWithInvalid, uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUintptr").
		Call(doubleRunWithInvalid, uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndFloat32").
		Call(doubleRunWithInvalid, float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndFloat64").
		Call(doubleRunWithInvalid, float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndComplex64").
		Call(doubleRunWithInvalid, complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndComplex128").
		Call(doubleRunWithInvalid, complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndArray").
		Call(doubleRunWithInvalid, [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndChan").
		Call(doubleRunWithInvalid, make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndFunc").
		Call(doubleRunWithInvalid, func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndInterface").
		Call(doubleRunWithInvalid, (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndMap").
		Call(doubleRunWithInvalid, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndPtr").
		Call(doubleRunWithInvalid, new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndSlice").
		Call(doubleRunWithInvalid, []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndString").
		Call(doubleRunWithInvalid, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndStruct").
		Call(doubleRunWithInvalid, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "InvalidAndUnsafePointer").
		Call(doubleRunWithInvalid, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndBoolAndPositiveResult").
		Call(doubleRun, true, true).
		ExpectResult(true, true)

	NewDeclarative(t, "BoolAndBoolAndNegativeResult").
		Call(doubleRun, true, false).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInt").
		Call(doubleRun, true, int(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInt8").
		Call(doubleRun, true, int8(10)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInt16").
		Call(doubleRun, true, int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInt32").
		Call(doubleRun, true, int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInt64").
		Call(doubleRun, true, int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUint").
		Call(doubleRun, true, uint(10)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUint8").
		Call(doubleRun, true, uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUint16").
		Call(doubleRun, true, uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUint32").
		Call(doubleRun, true, uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUint64").
		Call(doubleRun, true, uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUintptr").
		Call(doubleRun, true, uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndFloat32").
		Call(doubleRun, true, float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndFloat64").
		Call(doubleRun, true, float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndComplex64").
		Call(doubleRun, true, complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndComplex128").
		Call(doubleRun, true, complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndArray").
		Call(doubleRun, true, [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndChan").
		Call(doubleRun, true, make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndFunc").
		Call(doubleRun, true, func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndInterface").
		Call(doubleRun, true, (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndMap").
		Call(doubleRun, true, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndPtr").
		Call(doubleRun, true, new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndSlice").
		Call(doubleRun, true, []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndString").
		Call(doubleRun, true, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndStruct").
		Call(doubleRun, true, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "BoolAndUnsafePointer").
		Call(doubleRun, true, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndIntWithPositiveResult").
		Call(doubleRun, int(5), int(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndIntAndNegativeResult").
		Call(doubleRun, int(10), int(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt8AndNegativeResult").
		Call(doubleRun, int(10), int8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt8AndPositiveResult").
		Call(doubleRun, int(5), int8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt16AndNegativeResult").
		Call(doubleRun, int(10), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt16AndPositiveResult").
		Call(doubleRun, int(5), int16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt32AndNegativeResult").
		Call(doubleRun, int(10), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt32AndPositiveResult").
		Call(doubleRun, int(5), int32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt64AndNegativeResult").
		Call(doubleRun, int(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt64AndPositiveResult").
		Call(doubleRun, int(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint8AndNegativeResult").
		Call(doubleRun, int(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint8AndPositiveResult").
		Call(doubleRun, int(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint16AndNegativeResult").
		Call(doubleRun, int(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint16AndPositiveResult").
		Call(doubleRun, int(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint32AndNegativeResult").
		Call(doubleRun, int(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint32AndPositiveResult").
		Call(doubleRun, int(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUintptr").
		Call(doubleRun, int(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFloat32AndNegativeResult").
		Call(doubleRun, int(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFloat32AndPositiveResult").
		Call(doubleRun, int(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndFloat64AndNegativeResult").
		Call(doubleRun, int(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFloat64AndPositiveResult").
		Call(doubleRun, int(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndComplex64").
		Call(doubleRun, int(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndComplex128").
		Call(doubleRun, int(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndArray").
		Call(doubleRun, int(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndChan").
		Call(doubleRun, int(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFunc").
		Call(doubleRun, int(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInterface").
		Call(doubleRun, int(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndMap").
		Call(doubleRun, int(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndPtr").
		Call(doubleRun, int(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndSlice").
		Call(doubleRun, int(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndString").
		Call(doubleRun, int(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndStruct").
		Call(doubleRun, int(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUnsafePointer").
		Call(doubleRun, int(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt8AndPositiveResult").
		Call(doubleRun, int8(5), int8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt8AndNegativeResult").
		Call(doubleRun, int8(10), int8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt16AndNegativeResult").
		Call(doubleRun, int8(10), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt16AndPositiveResult").
		Call(doubleRun, int8(5), int16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt32AndNegativeResult").
		Call(doubleRun, int8(10), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt32AndPositiveResult").
		Call(doubleRun, int8(5), int32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt64AndNegativeResult").
		Call(doubleRun, int8(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt64AndPositiveResult").
		Call(doubleRun, int8(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint8AndNegativeResult").
		Call(doubleRun, int8(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint8AndPositiveResult").
		Call(doubleRun, int8(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint16AndNegativeResult").
		Call(doubleRun, int8(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint16AndPositiveResult").
		Call(doubleRun, int8(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint32AndNegativeResult").
		Call(doubleRun, int8(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint32AndPositiveResult").
		Call(doubleRun, int8(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUintptr").
		Call(doubleRun, int8(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFloat32AndNegativeResult").
		Call(doubleRun, int8(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFloat32AndPositiveResult").
		Call(doubleRun, int8(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndFloat64AndNegativeResult").
		Call(doubleRun, int8(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFloat64AndPositiveResult").
		Call(doubleRun, int8(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndComplex64").
		Call(doubleRun, int8(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndComplex128").
		Call(doubleRun, int8(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndArray").
		Call(doubleRun, int8(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndChan").
		Call(doubleRun, int8(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFunc").
		Call(doubleRun, int8(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInterface").
		Call(doubleRun, int8(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndMap").
		Call(doubleRun, int8(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndPtr").
		Call(doubleRun, int8(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndSlice").
		Call(doubleRun, int8(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndString").
		Call(doubleRun, int8(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndStruct").
		Call(doubleRun, int8(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUnsafePointer").
		Call(doubleRun, int8(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(5), int16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt16AndNegativeResult").
		Call(doubleRun, int16(10), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt32AndNegativeResult").
		Call(doubleRun, int16(10), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt32AndPositiveResult").
		Call(doubleRun, int16(5), int32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt64AndNegativeResult").
		Call(doubleRun, int16(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt64AndPositiveResult").
		Call(doubleRun, int16(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint8AndNegativeResult").
		Call(doubleRun, int16(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint8AndPositiveResult").
		Call(doubleRun, int16(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint16AndNegativeResult").
		Call(doubleRun, int16(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint16AndPositiveResult").
		Call(doubleRun, int16(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint32AndNegativeResult").
		Call(doubleRun, int16(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint32AndPositiveResult").
		Call(doubleRun, int16(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUintptr").
		Call(doubleRun, int16(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFloat32AndNegativeResult").
		Call(doubleRun, int16(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFloat32AndPositiveResult").
		Call(doubleRun, int16(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndFloat64AndNegativeResult").
		Call(doubleRun, int16(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFloat64AndPositiveResult").
		Call(doubleRun, int16(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndComplex64").
		Call(doubleRun, int16(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndComplex128").
		Call(doubleRun, int16(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndArray").
		Call(doubleRun, int16(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndChan").
		Call(doubleRun, int16(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFunc").
		Call(doubleRun, int16(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInterface").
		Call(doubleRun, int16(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndMap").
		Call(doubleRun, int16(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndPtr").
		Call(doubleRun, int16(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndSlice").
		Call(doubleRun, int16(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndString").
		Call(doubleRun, int16(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndStruct").
		Call(doubleRun, int16(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUnsafePointer").
		Call(doubleRun, int16(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInt32AndNegativeResult").
		Call(doubleRun, int32(10), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(5), int32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndInt64AndNegativeResult").
		Call(doubleRun, int32(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInt64AndPositiveResult").
		Call(doubleRun, int32(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint8AndNegativeResult").
		Call(doubleRun, int32(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint8AndPositiveResult").
		Call(doubleRun, int32(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint16AndNegativeResult").
		Call(doubleRun, int32(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint16AndPositiveResult").
		Call(doubleRun, int32(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint32AndNegativeResult").
		Call(doubleRun, int32(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint32AndPositiveResult").
		Call(doubleRun, int32(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUintptr").
		Call(doubleRun, int32(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFloat32AndNegativeResult").
		Call(doubleRun, int32(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFloat32AndPositiveResult").
		Call(doubleRun, int32(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndFloat64AndNegativeResult").
		Call(doubleRun, int32(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFloat64AndPositiveResult").
		Call(doubleRun, int32(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndComplex64").
		Call(doubleRun, int32(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndComplex128").
		Call(doubleRun, int32(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndArray").
		Call(doubleRun, int32(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndChan").
		Call(doubleRun, int32(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFunc").
		Call(doubleRun, int32(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInterface").
		Call(doubleRun, int32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndMap").
		Call(doubleRun, int32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndPtr").
		Call(doubleRun, int32(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndSlice").
		Call(doubleRun, int32(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndString").
		Call(doubleRun, int32(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndStruct").
		Call(doubleRun, int32(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUnsafePointer").
		Call(doubleRun, int32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndInt64AndNegativeResult").
		Call(doubleRun, int64(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint8AndNegativeResult").
		Call(doubleRun, int64(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint8AndPositiveResult").
		Call(doubleRun, int64(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint16AndNegativeResult").
		Call(doubleRun, int64(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint16AndPositiveResult").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint32AndNegativeResult").
		Call(doubleRun, int64(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint32AndPositiveResult").
		Call(doubleRun, int64(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUintptr").
		Call(doubleRun, int64(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFloat32AndNegativeResult").
		Call(doubleRun, int64(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFloat32AndPositiveResult").
		Call(doubleRun, int64(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndFloat64AndNegativeResult").
		Call(doubleRun, int64(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFloat64AndPositiveResult").
		Call(doubleRun, int64(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndComplex64").
		Call(doubleRun, int64(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndComplex128").
		Call(doubleRun, int64(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndArray").
		Call(doubleRun, int64(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndChan").
		Call(doubleRun, int64(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFunc").
		Call(doubleRun, int64(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndInterface").
		Call(doubleRun, int64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndMap").
		Call(doubleRun, int64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndPtr").
		Call(doubleRun, int64(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndSlice").
		Call(doubleRun, int64(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndString").
		Call(doubleRun, int64(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndStruct").
		Call(doubleRun, int64(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUnsafePointer").
		Call(doubleRun, int64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUintAndNegativeResult").
		Call(doubleRun, uint(10), uint(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(5), uint(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint8AndNegativeResult").
		Call(doubleRun, uint(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint8AndPositiveResult").
		Call(doubleRun, uint(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint16AndNegativeResult").
		Call(doubleRun, uint(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint16AndPositiveResult").
		Call(doubleRun, uint(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint32AndNegativeResult").
		Call(doubleRun, uint(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint32AndPositiveResult").
		Call(doubleRun, uint(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint64AndNegativeResult").
		Call(doubleRun, uint(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint64AndPositiveResult").
		Call(doubleRun, uint(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUintptr").
		Call(doubleRun, uint(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFloat32AndNegativeResult").
		Call(doubleRun, uint(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFloat32AndPositiveResult").
		Call(doubleRun, uint(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndFloat64AndNegativeResult").
		Call(doubleRun, uint(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFloat64AndPositiveResult").
		Call(doubleRun, uint(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndComplex64").
		Call(doubleRun, uint(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndComplex128").
		Call(doubleRun, uint(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndArray").
		Call(doubleRun, uint(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndChan").
		Call(doubleRun, uint(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFunc").
		Call(doubleRun, uint(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndInterface").
		Call(doubleRun, uint(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndMap").
		Call(doubleRun, uint(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndPtr").
		Call(doubleRun, uint(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndSlice").
		Call(doubleRun, uint(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndString").
		Call(doubleRun, uint(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndStruct").
		Call(doubleRun, uint(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUnsafePointer").
		Call(doubleRun, uint(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint8AndNegativeResult").
		Call(doubleRun, uint8(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint16AndNegativeResult").
		Call(doubleRun, uint8(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint16AndPositiveResult").
		Call(doubleRun, uint8(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint32AndNegativeResult").
		Call(doubleRun, uint8(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint32AndPositiveResult").
		Call(doubleRun, uint8(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint64AndNegativeResult").
		Call(doubleRun, uint8(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint64AndPositiveResult").
		Call(doubleRun, uint8(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUintptr").
		Call(doubleRun, uint8(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFloat32AndNegativeResult").
		Call(doubleRun, uint8(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFloat32AndPositiveResult").
		Call(doubleRun, uint8(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndFloat64AndNegativeResult").
		Call(doubleRun, uint8(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFloat64AndPositiveResult").
		Call(doubleRun, uint8(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndComplex64").
		Call(doubleRun, uint8(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndComplex128").
		Call(doubleRun, uint8(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndArray").
		Call(doubleRun, uint8(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndChan").
		Call(doubleRun, uint8(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFunc").
		Call(doubleRun, uint8(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndInterface").
		Call(doubleRun, uint8(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndMap").
		Call(doubleRun, uint8(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndPtr").
		Call(doubleRun, uint8(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndSlice").
		Call(doubleRun, uint8(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndString").
		Call(doubleRun, uint8(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndStruct").
		Call(doubleRun, uint8(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUnsafePointer").
		Call(doubleRun, uint8(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint16AndNegativeResult").
		Call(doubleRun, uint16(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint32AndNegativeResult").
		Call(doubleRun, uint16(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint32AndPositiveResult").
		Call(doubleRun, uint16(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint64AndNegativeResult").
		Call(doubleRun, uint16(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint64AndPositiveResult").
		Call(doubleRun, uint16(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUintptr").
		Call(doubleRun, uint16(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFloat32AndNegativeResult").
		Call(doubleRun, uint16(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFloat32AndPositiveResult").
		Call(doubleRun, uint16(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndFloat64AndNegativeResult").
		Call(doubleRun, uint16(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFloat64AndPositiveResult").
		Call(doubleRun, uint16(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndComplex64").
		Call(doubleRun, uint16(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndComplex128").
		Call(doubleRun, uint16(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndArray").
		Call(doubleRun, uint16(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndChan").
		Call(doubleRun, uint16(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFunc").
		Call(doubleRun, uint16(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndInterface").
		Call(doubleRun, uint16(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndMap").
		Call(doubleRun, uint16(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndPtr").
		Call(doubleRun, uint16(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndSlice").
		Call(doubleRun, uint16(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndString").
		Call(doubleRun, uint16(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndStruct").
		Call(doubleRun, uint16(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUnsafePointer").
		Call(doubleRun, uint16(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUint32AndNegativeResult").
		Call(doubleRun, uint32(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndUint64AndNegativeResult").
		Call(doubleRun, uint32(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUint64AndPositiveResult").
		Call(doubleRun, uint32(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndUintptr").
		Call(doubleRun, uint32(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFloat32AndNegativeResult").
		Call(doubleRun, uint32(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFloat32AndPositiveResult").
		Call(doubleRun, uint32(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndFloat64AndNegativeResult").
		Call(doubleRun, uint32(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFloat64AndPositiveResult").
		Call(doubleRun, uint32(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndComplex64").
		Call(doubleRun, uint32(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndComplex128").
		Call(doubleRun, uint32(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndArray").
		Call(doubleRun, uint32(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndChan").
		Call(doubleRun, uint32(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFunc").
		Call(doubleRun, uint32(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndInterface").
		Call(doubleRun, uint32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndMap").
		Call(doubleRun, uint32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndPtr").
		Call(doubleRun, uint32(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndSlice").
		Call(doubleRun, uint32(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndString").
		Call(doubleRun, uint32(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndStruct").
		Call(doubleRun, uint32(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUnsafePointer").
		Call(doubleRun, uint32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndUint64AndNegativeResult").
		Call(doubleRun, uint64(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndUintptr").
		Call(doubleRun, uint64(5), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndFloat32AndNegativeResult").
		Call(doubleRun, uint64(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndFloat32AndPositiveResult").
		Call(doubleRun, uint64(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndFloat64AndNegativeResult").
		Call(doubleRun, uint64(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndFloat64AndPositiveResult").
		Call(doubleRun, uint64(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndComplex64").
		Call(doubleRun, uint64(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndComplex128").
		Call(doubleRun, uint64(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndArray").
		Call(doubleRun, uint64(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndChan").
		Call(doubleRun, uint64(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndFunc").
		Call(doubleRun, uint64(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndInterface").
		Call(doubleRun, uint64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndMap").
		Call(doubleRun, uint64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndPtr").
		Call(doubleRun, uint64(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndSlice").
		Call(doubleRun, uint64(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndString").
		Call(doubleRun, uint64(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndStruct").
		Call(doubleRun, uint64(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndUnsafePointer").
		Call(doubleRun, uint64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndUintptrAndPositiveResult").
		Call(doubleRun, uintptr(5), uintptr(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintptrAndUintptrAndNegativeResult").
		Call(doubleRun, uintptr(5), uintptr(10)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndFloat32").
		Call(doubleRun, uintptr(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndFloat64").
		Call(doubleRun, uintptr(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndComplex64").
		Call(doubleRun, uintptr(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndComplex128").
		Call(doubleRun, uintptr(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndArray").
		Call(doubleRun, uintptr(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndChan").
		Call(doubleRun, uintptr(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndFunc").
		Call(doubleRun, uintptr(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndInterface").
		Call(doubleRun, uintptr(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndMap").
		Call(doubleRun, uintptr(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndPtr").
		Call(doubleRun, uintptr(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndSlice").
		Call(doubleRun, uintptr(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndString").
		Call(doubleRun, uintptr(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndStruct").
		Call(doubleRun, uintptr(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndUnsafePointer").
		Call(doubleRun, uintptr(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFloat32AndNegativeResult").
		Call(doubleRun, float32(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float32AndFloat64AndNegativeResult").
		Call(doubleRun, float32(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFloat64AndPositiveResult").
		Call(doubleRun, float32(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float32AndComplex64").
		Call(doubleRun, float32(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndComplex128").
		Call(doubleRun, float32(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndArray").
		Call(doubleRun, float32(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndChan").
		Call(doubleRun, float32(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFunc").
		Call(doubleRun, float32(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndInterface").
		Call(doubleRun, float32(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndMap").
		Call(doubleRun, float32(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndPtr").
		Call(doubleRun, float32(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndSlice").
		Call(doubleRun, float32(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndString").
		Call(doubleRun, float32(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndStruct").
		Call(doubleRun, float32(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndUnsafePointer").
		Call(doubleRun, float32(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndFloat64AndNegativeResult").
		Call(doubleRun, float64(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float64AndComplex64").
		Call(doubleRun, float64(5), complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndComplex128").
		Call(doubleRun, float64(5), complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndArray").
		Call(doubleRun, float64(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndChan").
		Call(doubleRun, float64(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndFunc").
		Call(doubleRun, float64(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndInterface").
		Call(doubleRun, float64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndMap").
		Call(doubleRun, float64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndPtr").
		Call(doubleRun, float64(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndSlice").
		Call(doubleRun, float64(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndString").
		Call(doubleRun, float64(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndStruct").
		Call(doubleRun, float64(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndUnsafePointer").
		Call(doubleRun, float64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndComplex64AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)+5i).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex64AndComplex64AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndComplex128AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex64AndComplex128AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndArray").
		Call(doubleRun, complex64(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndChan").
		Call(doubleRun, complex64(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndFunc").
		Call(doubleRun, complex64(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndInterface").
		Call(doubleRun, complex64(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndMap").
		Call(doubleRun, complex64(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndPtr").
		Call(doubleRun, complex64(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndSlice").
		Call(doubleRun, complex64(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndString").
		Call(doubleRun, complex64(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndStruct").
		Call(doubleRun, complex64(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndUnsafePointer").
		Call(doubleRun, complex64(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndComplex128AndPositiveResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex128AndComplex128AndNegativeResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndArray").
		Call(doubleRun, complex128(5), [1]int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndChan").
		Call(doubleRun, complex128(5), make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndFunc").
		Call(doubleRun, complex128(5), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndInterface").
		Call(doubleRun, complex128(5), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndMap").
		Call(doubleRun, complex128(5), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndPtr").
		Call(doubleRun, complex128(5), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndSlice").
		Call(doubleRun, complex128(5), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndString").
		Call(doubleRun, complex128(5), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndStruct").
		Call(doubleRun, complex128(5), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndUnsafePointer").
		Call(doubleRun, complex128(5), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndPositiveResult").
		Call(doubleRun, [2]string{"data", "string"}, [2]string{"data", "string"}).
		ExpectResult(true, true)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByType").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByElementValue").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByElementsCount").
		Call(doubleRun, [2]string{"data", "string"}, [3]interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndChan").
		Call(doubleRun, [1]string{"data"}, make(chan int)).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndFunc").
		Call(doubleRun, [1]string{"data"}, func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndInterface").
		Call(doubleRun, [1]string{"data"}, (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndMap").
		Call(doubleRun, [1]string{"data"}, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndPtr").
		Call(doubleRun, [1]string{"data"}, new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndSlice").
		Call(doubleRun, [1]string{"data"}, []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndString").
		Call(doubleRun, [1]string{"data"}, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndStruct").
		Call(doubleRun, [1]string{"data"}, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndUnsafePointer").
		Call(doubleRun, [1]string{"data"}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndChanAndNegativeResult").
		Call(doubleRun, make(chan int), make(chan int)).
		ExpectResult(false, false)

	{
		chanFixture := make(chan int)

		NewDeclarative(t, "ChanAndChanAndNegativeResult").
			Call(doubleRun, chanFixture, chanFixture).
			ExpectResult(true, true)
	}

	NewDeclarative(t, "ChanAndFunc").
		Call(doubleRun, make(chan int), func() {}).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndInterface").
		Call(doubleRun, make(chan int), (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndMap").
		Call(doubleRun, make(chan int), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndPtr").
		Call(doubleRun, make(chan int), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndSlice").
		Call(doubleRun, make(chan int), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndString").
		Call(doubleRun, make(chan int), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndStruct").
		Call(doubleRun, make(chan int), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndUnsafePointer").
		Call(doubleRun, make(chan int), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndFuncWithNegativeResult").
		Call(doubleRun, func() {}, func() {}).
		ExpectResult(false, false)

	{
		funcFixture := func() {}

		NewDeclarative(t, "FuncAndFuncWithPositiveResult").
			Call(doubleRun, funcFixture, funcFixture).
			ExpectResult(true, true)
	}

	NewDeclarative(t, "FuncAndInterface").
		Call(doubleRun, func() {}, (*interface{})(nil)).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndMap").
		Call(doubleRun, func() {}, map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndPtr").
		Call(doubleRun, func() {}, new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndSlice").
		Call(doubleRun, func() {}, []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndString").
		Call(doubleRun, func() {}, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndStruct").
		Call(doubleRun, func() {}, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "FuncAndUnsafePointer").
		Call(doubleRun, func() {}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndInterface").
		Call(doubleRun, (*interface{})(nil), (*interface{})(nil)).
		ExpectResult(true, true)

	NewDeclarative(t, "InterfaceAndMap").
		Call(doubleRun, (*interface{})(nil), map[int]int{1: 1}).
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndPtr").
		Call(doubleRun, (*interface{})(nil), new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndSlice").
		Call(doubleRun, (*interface{})(nil), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndString").
		Call(doubleRun, (*interface{})(nil), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndStruct").
		Call(doubleRun, (*interface{})(nil), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "InterfaceAndUnsafePointer").
		Call(doubleRun, (*interface{})(nil), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndMapWithPositiveResult").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "MapAndMapWithNegativeResultByType").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]interface{}{"First": 1, "Second": 2},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndMapWithNegativeResultByElementValue").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2000},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndMapWithNegativeResultByElementCount").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2, "third": 3},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndPtr").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, new(int)).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndSlice").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndString").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndStruct").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndUnsafePointer").
		Call(doubleRun, map[string]int{"First": 1, "Second": 2}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	{
		stringFixture1 := "First"
		stringFixture2 := "First"
		stringFixture3 := "data"

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(true, true)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	NewDeclarative(t, "PtrAndSlice").
		Call(doubleRun, new(string), []int{5}).
		ExpectResult(false, false)

	NewDeclarative(t, "PtrAndString").
		Call(doubleRun, new(string), "data").
		ExpectResult(false, false)

	NewDeclarative(t, "PtrAndStruct").
		Call(doubleRun, new(string), struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "PtrAndUnsafePointer").
		Call(doubleRun, new(string), unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndSliceAndPositiveResult").
		Call(doubleRun, []string{"data", "string"}, []string{"data", "string"}).
		ExpectResult(true, true)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByType").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByElementValue").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByElementsCount").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndString").
		Call(doubleRun, []string{"First", "Second"}, "data").
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndStruct").
		Call(doubleRun, []string{"First", "Second"}, struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndUnsafePointer").
		Call(doubleRun, []string{"First", "Second"}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	NewDeclarative(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", "data").
		ExpectResult(true, true)

	NewDeclarative(t, "StringAndStringWithNegativeResult").
		Call(doubleRun, "data", "First").
		ExpectResult(false, false)

	NewDeclarative(t, "StringAndStruct").
		Call(doubleRun, "data", struct{}{}).
		ExpectResult(false, false)

	NewDeclarative(t, "StringAndUnsafePointer").
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

		NewDeclarative(t, "StructAndStructWithPositiveResult").
			Call(doubleRun, structFixture1, structFixture1).
			ExpectResult(true, true)

		NewDeclarative(t, "StructAndStructWithNegativeResultByFieldValue").
			Call(doubleRun, structFixture1, structFixture2).
			ExpectResult(false, false)

		NewDeclarative(t, "StructAndStructWithNegativeResultByFieldName").
			Call(doubleRun, structFixture1, structFixture3).
			ExpectResult(false, false)
	}

	NewDeclarative(t, "StructAndUnsafePointer").
		Call(doubleRun, struct{}{}, unsafe.Pointer(new(int))).
		ExpectResult(false, false)

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
		Call(doubleRun, unsafe.Pointer(new(int)), unsafe.Pointer(new(int))).
		ExpectResult(false, false)
}

//noinspection GoRedundantConversion
func TestEqualComparator_Compare_WithDelta(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}, delta float64) (bool, bool) {
		return (&EqualComparator{numericDelta: delta}).Compare(x, y),
			(&EqualComparator{numericDelta: delta}).Compare(y, x)
	}

	NewDeclarative(t, "IntAndIntAndPositiveResult").
		Call(doubleRun, int(10), int(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt8AndPositiveResult").
		Call(doubleRun, int(10), int8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt16AndPositiveResult").
		Call(doubleRun, int(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt32AndPositiveResult").
		Call(doubleRun, int(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndInt64AndPositiveResult").
		Call(doubleRun, int(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint8AndPositiveResult").
		Call(doubleRun, int(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint16AndPositiveResult").
		Call(doubleRun, int(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndUint32AndPositiveResult").
		Call(doubleRun, int(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndFloat32AndPositiveResult").
		Call(doubleRun, int(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndFloat64AndPositiveResult").
		Call(doubleRun, int(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt8AndPositiveResult").
		Call(doubleRun, int8(10), int8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt16AndPositiveResult").
		Call(doubleRun, int8(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt32AndPositiveResult").
		Call(doubleRun, int8(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt64AndPositiveResult").
		Call(doubleRun, int8(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint8AndPositiveResult").
		Call(doubleRun, int8(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint16AndPositiveResult").
		Call(doubleRun, int8(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndUint32AndPositiveResult").
		Call(doubleRun, int8(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndFloat32AndPositiveResult").
		Call(doubleRun, int8(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndFloat64AndPositiveResult").
		Call(doubleRun, int8(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(10), int16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt32AndPositiveResult").
		Call(doubleRun, int16(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt64AndPositiveResult").
		Call(doubleRun, int16(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint8AndPositiveResult").
		Call(doubleRun, int16(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint16AndPositiveResult").
		Call(doubleRun, int16(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndUint32AndPositiveResult").
		Call(doubleRun, int16(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndFloat32AndPositiveResult").
		Call(doubleRun, int16(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndFloat64AndPositiveResult").
		Call(doubleRun, int16(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(10), int32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndInt64AndPositiveResult").
		Call(doubleRun, int32(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint8AndPositiveResult").
		Call(doubleRun, int32(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint16AndPositiveResult").
		Call(doubleRun, int32(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndUint32AndPositiveResult").
		Call(doubleRun, int32(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndFloat32AndPositiveResult").
		Call(doubleRun, int32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndFloat64AndPositiveResult").
		Call(doubleRun, int32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(10), int64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint8AndPositiveResult").
		Call(doubleRun, int64(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint16AndPositiveResult").
		Call(doubleRun, int64(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint32AndPositiveResult").
		Call(doubleRun, int64(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndFloat32AndPositiveResult").
		Call(doubleRun, int64(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndFloat64AndPositiveResult").
		Call(doubleRun, int64(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(10), uint(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint8AndPositiveResult").
		Call(doubleRun, uint(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint16AndPositiveResult").
		Call(doubleRun, uint(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint32AndPositiveResult").
		Call(doubleRun, uint(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint64AndPositiveResult").
		Call(doubleRun, uint(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndFloat32AndPositiveResult").
		Call(doubleRun, uint(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndFloat64AndPositiveResult").
		Call(doubleRun, uint(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(10), uint8(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint16AndPositiveResult").
		Call(doubleRun, uint8(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint32AndPositiveResult").
		Call(doubleRun, uint8(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint64AndPositiveResult").
		Call(doubleRun, uint8(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndFloat32AndPositiveResult").
		Call(doubleRun, uint8(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndFloat64AndPositiveResult").
		Call(doubleRun, uint8(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(10), uint16(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint32AndPositiveResult").
		Call(doubleRun, uint16(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint64AndPositiveResult").
		Call(doubleRun, uint16(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndFloat32AndPositiveResult").
		Call(doubleRun, uint16(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndFloat64AndPositiveResult").
		Call(doubleRun, uint16(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(10), uint32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndUint64AndPositiveResult").
		Call(doubleRun, uint32(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndFloat32AndPositiveResult").
		Call(doubleRun, uint32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndFloat64AndPositiveResult").
		Call(doubleRun, uint32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(10), uint64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndFloat32AndPositiveResult").
		Call(doubleRun, uint64(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndFloat64AndPositiveResult").
		Call(doubleRun, uint64(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(10), float32(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float32AndFloat64AndPositiveResult").
		Call(doubleRun, float32(10), float64(5), float64(20)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(10), float64(5), float64(20)).
		ExpectResult(true, true)
}

//noinspection GoRedundantConversion
func TestEqualComparator_Compare_WithSameType(t *testing.T) {
	doubleRun := func(x interface{}, y interface{}) (bool, bool) {
		return (&EqualComparator{sameType: true}).Compare(x, y), (&EqualComparator{sameType: true}).Compare(y, x)
	}

	NewDeclarative(t, "InvalidAndInvalid").
		Call(
			func() bool {
				return (&EqualComparator{sameType: true}).Compare(nil, nil)
			},
		).
		ExpectResult(true)

	NewDeclarative(t, "BoolAndBoolAndPositiveResult").
		Call(doubleRun, true, true).
		ExpectResult(true, true)

	NewDeclarative(t, "BoolAndBoolAndNegativeResult").
		Call(doubleRun, true, false).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndIntWithPositiveResult").
		Call(doubleRun, int(5), int(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "IntAndIntAndNegativeResult").
		Call(doubleRun, int(10), int(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt8").
		Call(doubleRun, int(5), int8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt16").
		Call(doubleRun, int(5), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt32").
		Call(doubleRun, int(5), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndInt64").
		Call(doubleRun, int(5), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint8").
		Call(doubleRun, int(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint").
		Call(doubleRun, int(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndUint32").
		Call(doubleRun, int(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFloat32").
		Call(doubleRun, int(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "IntAndFloat64").
		Call(doubleRun, int(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt8WithPositiveResult").
		Call(doubleRun, int8(5), int8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int8AndInt8WithNegativeResult").
		Call(doubleRun, int8(10), int8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt16").
		Call(doubleRun, int8(5), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt32").
		Call(doubleRun, int8(5), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndInt64").
		Call(doubleRun, int8(5), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint8").
		Call(doubleRun, int8(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint16").
		Call(doubleRun, int8(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndUint32").
		Call(doubleRun, int8(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFloat32").
		Call(doubleRun, int8(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int8AndFloat64").
		Call(doubleRun, int8(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt16AndPositiveResult").
		Call(doubleRun, int16(5), int16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int16AndInt16AndNegativeResult").
		Call(doubleRun, int16(10), int16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt32").
		Call(doubleRun, int16(5), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndInt64").
		Call(doubleRun, int16(5), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint8").
		Call(doubleRun, int16(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint16").
		Call(doubleRun, int16(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndUint32").
		Call(doubleRun, int16(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFloat32").
		Call(doubleRun, int16(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int16AndFloat64").
		Call(doubleRun, int16(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInt32AndNegativeResult").
		Call(doubleRun, int32(10), int32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndInt32AndPositiveResult").
		Call(doubleRun, int32(5), int32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int32AndInt64").
		Call(doubleRun, int32(5), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint8").
		Call(doubleRun, int32(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint16").
		Call(doubleRun, int32(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndUint32").
		Call(doubleRun, int32(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFloat32").
		Call(doubleRun, int32(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int32AndFloat64").
		Call(doubleRun, int32(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndInt64AndNegativeResult").
		Call(doubleRun, int64(10), int64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndInt64AndPositiveResult").
		Call(doubleRun, int64(5), int64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Int64AndUint8").
		Call(doubleRun, int64(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint16AndNegativeResult").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint16").
		Call(doubleRun, int64(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndUint32").
		Call(doubleRun, int64(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFloat32").
		Call(doubleRun, int64(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Int64AndFloat64").
		Call(doubleRun, int64(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUintAndNegativeResult").
		Call(doubleRun, uint(10), uint(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUintAndPositiveResult").
		Call(doubleRun, uint(5), uint(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintAndUint8").
		Call(doubleRun, uint(5), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint16").
		Call(doubleRun, uint(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint32").
		Call(doubleRun, uint(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndUint64").
		Call(doubleRun, uint(5), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFloat32").
		Call(doubleRun, uint(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintAndFloat64").
		Call(doubleRun, uint(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint8AndNegativeResult").
		Call(doubleRun, uint8(10), uint8(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint8AndPositiveResult").
		Call(doubleRun, uint8(5), uint8(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint8AndUint16").
		Call(doubleRun, uint8(5), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint32").
		Call(doubleRun, uint8(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndUint64").
		Call(doubleRun, uint8(5), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFloat32").
		Call(doubleRun, uint8(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint8AndFloat64").
		Call(doubleRun, uint8(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint16AndNegativeResult").
		Call(doubleRun, uint16(10), uint16(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint16AndPositiveResult").
		Call(doubleRun, uint16(5), uint16(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint16AndUint32").
		Call(doubleRun, uint16(5), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndUint64").
		Call(doubleRun, uint16(5), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFloat32").
		Call(doubleRun, uint16(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint16AndFloat64").
		Call(doubleRun, uint16(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUint32AndNegativeResult").
		Call(doubleRun, uint32(10), uint32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndUint32AndPositiveResult").
		Call(doubleRun, uint32(5), uint32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint32AndUint64").
		Call(doubleRun, uint32(5), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFloat32").
		Call(doubleRun, uint32(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint32AndFloat64").
		Call(doubleRun, uint32(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndUint64AndNegativeResult").
		Call(doubleRun, uint64(10), uint64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndUint64AndPositiveResult").
		Call(doubleRun, uint64(5), uint64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Uint64AndFloat32").
		Call(doubleRun, uint64(5), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Uint64AndFloat64").
		Call(doubleRun, uint64(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "UintptrAndUintptrAndPositiveResult").
		Call(doubleRun, uintptr(5), uintptr(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "UintptrAndUintptrAndNegativeResult").
		Call(doubleRun, uintptr(10), uintptr(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFloat32AndNegativeResult").
		Call(doubleRun, float32(10), float32(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float32AndFloat32AndPositiveResult").
		Call(doubleRun, float32(5), float32(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Float32AndFloat64").
		Call(doubleRun, float32(5), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndFloat64AndNegativeResult").
		Call(doubleRun, float64(10), float64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Float64AndFloat64AndPositiveResult").
		Call(doubleRun, float64(5), float64(5)).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex64AndComplex64AndPositiveResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)+5i).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex64AndComplex64AndNegativeResult").
		Call(doubleRun, complex64(5)+5i, complex64(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex64AndComplex128").
		Call(doubleRun, complex64(5)+5i, complex128(5)+5i).
		ExpectResult(false, false)

	NewDeclarative(t, "Complex128AndComplex128AndPositiveResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)+5i).
		ExpectResult(true, true)

	NewDeclarative(t, "Complex128AndComplex128AndNegativeResult").
		Call(doubleRun, complex128(5)+5i, complex128(5)).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndPositiveResult").
		Call(doubleRun, [2]string{"data", "string"}, [2]string{"data", "string"}).
		ExpectResult(true, true)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByType").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByElementValue").
		Call(doubleRun, [2]string{"data", "string"}, [2]interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ArrayAndArrayAndNegativeResultByElementsCount").
		Call(doubleRun, [2]string{"data", "string"}, [3]interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewDeclarative(t, "ChanAndChanAndNegativeResult").
		Call(doubleRun, make(chan int), make(chan int)).
		ExpectResult(false, false)

	{
		chanFixture := make(chan int)

		NewDeclarative(t, "ChanAndChanAndNegativeResult").
			Call(doubleRun, chanFixture, chanFixture).
			ExpectResult(true, true)
	}

	NewDeclarative(t, "FuncAndFuncWithNegativeResult").
		Call(doubleRun, func() {}, func() {}).
		ExpectResult(false, false)

	{
		funcFixture := func() {}

		NewDeclarative(t, "FuncAndFuncWithPositiveResult").
			Call(doubleRun, funcFixture, funcFixture).
			ExpectResult(true, true)
	}

	NewDeclarative(t, "InterfaceAndInterface").
		Call(doubleRun, (*interface{})(nil), (*interface{})(nil)).
		ExpectResult(true, true)

	NewDeclarative(t, "MapAndMapWithPositiveResult").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "MapAndMapWithNegativeResultByType").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]interface{}{"First": 1, "Second": 2},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndMapWithNegativeResultByElementValue").
		Call(
			doubleRun,
			map[string]int{"First": 1, "Second": 2},
			map[string]int{"First": 1, "Second": 2000},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "MapAndMapWithNegativeResultByElementCount").
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

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(true, true)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	NewDeclarative(t, "SliceAndSliceAndPositiveResult").
		Call(doubleRun, []string{"data", "string"}, []string{"data", "string"}).
		ExpectResult(true, true)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByType").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByElementValue").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "invalid"}).
		ExpectResult(false, false)

	NewDeclarative(t, "SliceAndSliceAndNegativeResultByElementsCount").
		Call(doubleRun, []string{"data", "string"}, []interface{}{"data", "string", "data"}).
		ExpectResult(false, false)

	NewDeclarative(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", "data").
		ExpectResult(true, true)

	NewDeclarative(t, "StringAndStringWithNegativeResult").
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

		NewDeclarative(t, "StructAndStructWithPositiveResult").
			Call(doubleRun, structFixture1, structFixture1).
			ExpectResult(true, true)

		NewDeclarative(t, "StructAndStructWithNegativeResultByFieldValue").
			Call(doubleRun, structFixture1, structFixture2).
			ExpectResult(false, false)

		NewDeclarative(t, "StructAndStructWithNegativeResultByFieldName").
			Call(doubleRun, structFixture1, structFixture3).
			ExpectResult(false, false)
	}

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
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

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySamePointer").
			Call(doubleRun, &stringFixture1, &stringFixture1).
			ExpectResult(true, true)

		NewDeclarative(t, "PtrAndPtrWithPositiveResultBySameValue").
			Call(doubleRun, &stringFixture1, &stringFixture2).
			ExpectResult(false, false)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByType").
			Call(doubleRun, &stringFixture1, new(int)).
			ExpectResult(false, false)

		NewDeclarative(t, "PtrAndPtrWithNegativeResultByValue").
			Call(doubleRun, &stringFixture1, &stringFixture3).
			ExpectResult(false, false)
	}

	unsafePointerFixture := unsafe.Pointer(new(string))

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithPositiveResult").
		Call(doubleRun, unsafePointerFixture, unsafePointerFixture).
		ExpectResult(true, true)

	NewDeclarative(t, "UnsafePointerAndUnsafePointerWithNegativeResult").
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

	NewDeclarative(t, "StringAndStringWithPositiveResult").
		Call(doubleRun, "data", equalFixture("data")).
		ExpectResult(true, true)

	NewDeclarative(t, "StringAndStringWithNegativeResult").
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

	NewDeclarative(t, "IgnoreUnexportedDisabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "Second"},
			&testType{Exported: "First", unexported: "Second"},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "IgnoreUnexportedDisabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "First"},
			&testType{Exported: "First", unexported: "Second"},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "IgnoreUnexportedEnabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{Exported: "First", unexported: "First"},
			&testType{Exported: "First", unexported: "Second"},
			testType{},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "IgnoreUnexportedEnabledAndNegativeResult").
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

	NewDeclarative(t, "IgnoreFieldsDisabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "second"},
			&testType{First: "first", Second: "second"},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "IgnoreFieldsDisabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "first", Second: "second"},
		).
		ExpectResult(false, false)

	NewDeclarative(t, "IgnoreFieldsEnabledAndPositiveResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "first", Second: "second"},
			IgnoreFieldsOption{Type: testType{}, Fields: []string{"Second"}},
		).
		ExpectResult(true, true)

	NewDeclarative(t, "IgnoreFieldsEnabledAndNegativeResult").
		Call(
			doubleRun,
			&testType{First: "first", Second: "data"},
			&testType{First: "second", Second: "second"},
			IgnoreFieldsOption{Type: testType{}, Fields: []string{"Second"}},
		).
		ExpectResult(false, false)
}
