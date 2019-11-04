package unit

import "testing"

type TestingT interface {
	Name() string
	Helper()
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fail()
	Run(name string, test func(*testing.T)) bool
}

type Constraint interface {
	Check(value interface{}) bool
	String() string
	Details(value interface{}) string
}

type EqualComparer interface {
	Compare(x interface{}, y interface{}) bool
	Diff(x interface{}, y interface{}) string
}

type Equaler interface {
	Equal(value interface{}) bool
}
