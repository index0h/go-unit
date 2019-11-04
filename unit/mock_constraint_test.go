package unit

import (
	"testing"
)

type MockConstraint struct {
	test *testing.T

	callCheck []struct {
		value  interface{}
		result bool
	}
	callString []struct {
		result string
	}
	callDetails []struct {
		value  interface{}
		result string
	}
}

func NewMockConstraint(test *testing.T) *MockConstraint {
	return &MockConstraint{
		test: test,
	}
}

func (m *MockConstraint) Check(value interface{}) bool {
	m.test.Helper()

	if len(m.callCheck) == 0 {
		m.test.Error("No call recorded for method 'Check'")
		m.test.Fail()
	}

	call := m.callCheck[0]

	m.callCheck = m.callCheck[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.value, value); diff != "" {
		m.test.Errorf("Expected argument not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	return call.result
}

func (m *MockConstraint) String() string {
	m.test.Helper()

	if len(m.callString) == 0 {
		m.test.Error("No call recorded for method 'String'")
		m.test.Fail()
	}

	call := m.callString[0]

	m.callString = m.callString[1:]

	return call.result
}

func (m *MockConstraint) Details(value interface{}) string {
	m.test.Helper()

	if len(m.callDetails) == 0 {
		m.test.Error("No call recorded for method 'Details'")
		m.test.Fail()
	}

	call := m.callDetails[0]

	m.callDetails = m.callDetails[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.value, value); diff != "" {
		m.test.Errorf("Expected args not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	return call.result
}

func (m *MockConstraint) RecordCheck(value interface{}, result bool) *MockConstraint {
	call := struct {
		value  interface{}
		result bool
	}{
		value:  value,
		result: result,
	}

	m.callCheck = append(m.callCheck, call)

	return m
}

func (m *MockConstraint) RecordString(result string) *MockConstraint {
	call := struct {
		result string
	}{
		result: result,
	}

	m.callString = append(m.callString, call)

	return m
}

func (m *MockConstraint) RecordDetails(value interface{}, result string) *MockConstraint {
	call := struct {
		value  interface{}
		result string
	}{
		value:  value,
		result: result,
	}

	m.callDetails = append(m.callDetails, call)

	return m
}
