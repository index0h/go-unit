package unit

import (
	"testing"
)

type MockEqualComparer struct {
	test *testing.T

	callCompare []struct {
		x      interface{}
		y      interface{}
		result bool
	}
	callDiff []struct {
		x      interface{}
		y      interface{}
		result string
	}
}

func NewMockEqualComparer(test *testing.T) *MockEqualComparer {
	return &MockEqualComparer{
		test: test,
	}
}

func (m *MockEqualComparer) Compare(x interface{}, y interface{}) bool {
	m.test.Helper()

	if len(m.callCompare) == 0 {
		m.test.Error("No call recorded for method 'Check'")
		m.test.Fail()
	}

	call := m.callCompare[0]

	m.callCompare = m.callCompare[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.x, x); diff != "" {
		m.test.Errorf("Expected argument 'x' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.y, y); diff != "" {
		m.test.Errorf("Expected argument 'y' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	return call.result
}

func (m *MockEqualComparer) Diff(x interface{}, y interface{}) string {
	m.test.Helper()

	if len(m.callDiff) == 0 {
		m.test.Error("No call recorded for method 'Details'")
		m.test.Fail()
	}

	call := m.callDiff[0]

	m.callDiff = m.callDiff[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.x, x); diff != "" {
		m.test.Errorf("Expected argument 'x' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.y, y); diff != "" {
		m.test.Errorf("Expected argument 'x' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	return call.result
}

func (m *MockEqualComparer) RecordCompare(x interface{}, y interface{}, result bool) *MockEqualComparer {
	call := struct {
		x      interface{}
		y      interface{}
		result bool
	}{
		x:      x,
		y:      y,
		result: result,
	}

	m.callCompare = append(m.callCompare, call)

	return m
}

func (m *MockEqualComparer) RecordDiff(x interface{}, y interface{}, result string) *MockEqualComparer {
	call := struct {
		x      interface{}
		y      interface{}
		result string
	}{
		x:      x,
		y:      y,
		result: result,
	}

	m.callDiff = append(m.callDiff, call)

	return m
}
