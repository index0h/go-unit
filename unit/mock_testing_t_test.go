package unit

import (
	"testing"
)

type MockTestingT struct {
	test *testing.T

	callName []struct {
		result string
	}
	callFail   []struct{}
	callHelper []struct{}
	callErrorf []struct {
		format string
		args   []interface{}
	}
	callError []struct {
		args []interface{}
	}
	callRun []struct {
		name   string
		test   func(*testing.T)
		result bool
	}
}

func NewMockTestingT(test *testing.T) *MockTestingT {
	return &MockTestingT{
		test: test,
	}
}

func (m *MockTestingT) Name() string {
	m.test.Helper()

	if len(m.callFail) == 0 {
		m.test.Error("No call recorded for method 'Fail'")
		m.test.Fail()
	}

	call := m.callName[0]

	m.callName = m.callName[1:]

	return call.result
}

func (m *MockTestingT) Fail() {
	m.test.Helper()

	if len(m.callFail) == 0 {
		m.test.Error("No call recorded for method 'Fail'")
		m.test.Fail()
	}

	m.callFail = m.callFail[1:]
}

func (m *MockTestingT) Helper() {
	m.test.Helper()

	if len(m.callHelper) == 0 {
		m.test.Error("No call recorded for method 'Helper'")
		m.test.Fail()
	}

	m.callHelper = m.callHelper[1:]
}

func (m *MockTestingT) Errorf(format string, args ...interface{}) {
	m.test.Helper()

	if len(m.callErrorf) == 0 {
		m.test.Error("No call recorded for method 'Errorf'")
		m.test.Fail()
	}

	call := m.callErrorf[0]

	m.callErrorf = m.callErrorf[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.format, format); diff != "" {
		m.test.Errorf("Expected argument 'format' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	if len(call.args) == len(args) {
		for _, arg := range args {
			if constraint, ok := arg.(Constraint); constraint != nil && ok && !constraint.Check(arg) {
				if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.args, args); diff != "" {
					m.test.Errorf("Expected argument 'args' not equal with actual: \n%s", diff)
					m.test.Fail()
				}

				return
			}
		}
	} else {
		if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.args, args); diff != "" {
			m.test.Errorf("Expected argument 'args' not equal with actual: \n%s", diff)
			m.test.Fail()
		}
	}
}

func (m *MockTestingT) Run(name string, test func(*testing.T)) bool {
	m.test.Helper()

	if len(m.callErrorf) == 0 {
		m.test.Error("No call recorded for method 'Run'")
		m.test.Fail()
	}

	call := m.callRun[0]

	m.callRun = m.callRun[1:]

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.name, name); diff != "" {
		m.test.Errorf("Expected argument 'name' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	if diff := NewEqualComparator(SamePointerOption{Value: true}).Diff(call.test, test); diff != "" {
		m.test.Errorf("Expected argument 'test' not equal with actual: \n%s", diff)
		m.test.Fail()
	}

	return call.result
}

func (m *MockTestingT) Error(args ...interface{}) {
	m.test.Helper()

	if len(m.callError) == 0 {
		m.test.Error("No call recorded for method 'Errorf'")
		m.test.Fail()
	}

	call := m.callError[0]

	m.callError = m.callError[1:]

	if diff := NewEqualComparator().Diff(call.args, args); diff != "" {
		m.test.Errorf("Expected argument 'args' not equal with actual: \n%s", diff)
		m.test.Fail()
	}
}

func (m *MockTestingT) RecordName(result string) *MockTestingT {
	call := struct {
		result string
	}{
		result: result,
	}

	m.callName = append(m.callName, call)

	return m
}

func (m *MockTestingT) RecordFail() *MockTestingT {
	call := struct{}{}

	m.callFail = append(m.callFail, call)

	return m
}

func (m *MockTestingT) RecordHelper() *MockTestingT {
	call := struct{}{}

	m.callHelper = append(m.callHelper, call)

	return m
}

func (m *MockTestingT) RecordErrorf(format string, args ...interface{}) *MockTestingT {
	call := struct {
		format string
		args   []interface{}
	}{
		format: format,
		args:   args,
	}

	m.callErrorf = append(m.callErrorf, call)

	return m
}

func (m *MockTestingT) RecordRun(name string, test func(*testing.T), result bool) *MockTestingT {
	call := struct {
		name   string
		test   func(*testing.T)
		result bool
	}{
		name:   name,
		test:   test,
		result: result,
	}

	m.callRun = append(m.callRun, call)

	return m
}

func (m *MockTestingT) RecordError(args ...interface{}) *MockTestingT {
	call := struct {
		args []interface{}
	}{
		args: args,
	}

	m.callError = append(m.callError, call)

	return m
}
