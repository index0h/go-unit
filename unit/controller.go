package unit

import (
	"reflect"
	"sync"
)

type Controller struct {
	mutex    sync.Mutex
	test     TestingT
	finishes []func()
}

func NewController(test TestingT) *Controller {
	if test == nil {
		err := NewNotNilError("test")

		panic(err)
	}

	return &Controller{
		test:     test,
		finishes: []func(){},
	}
}

func (c *Controller) TestingT() TestingT {
	return c.test
}

func (c *Controller) DeclarativeTest(name string) *Declarative {
	return NewDeclarative(c.test, name)
}

func (c *Controller) RegisterFinish(finish func()) {
	c.AssertNotNil(finish)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.finishes = append(c.finishes, finish)
}

func (c *Controller) Finish() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, finish := range c.finishes {
		finish()
	}
}

func (c *Controller) And(constraint ...Constraint) Constraint {
	return NewAndConstraint(constraint...)
}

func (c *Controller) Any() Constraint {
	return NewAnyConstraint()
}

func (c *Controller) Callback(callback func(value interface{}) bool) Constraint {
	return NewCallbackConstraint(callback)
}

func (c *Controller) NotCallback(callback func(value interface{}) bool) Constraint {
	return c.Not(NewCallbackConstraint(callback))
}

func (c *Controller) Contains(element interface{}, options ...interface{}) Constraint {
	return NewContainsConstraint(element, options...)
}

func (c *Controller) NotContains(element interface{}, options ...interface{}) Constraint {
	return c.Not(NewContainsConstraint(element, options...))
}

func (c *Controller) Elements(constraints ...Constraint) Constraint {
	return NewElementsConstraint(constraints...)
}

func (c *Controller) NotElements(constraints ...Constraint) Constraint {
	return c.Not(NewElementsConstraint(constraints...))
}

func (c *Controller) ValueElements(expectedValues []interface{}, options ...interface{}) Constraint {
	return NewValueElementsConstraint(NewEqualComparator(options...), expectedValues...)
}

func (c *Controller) NotValueElements(expectedValues []interface{}, options ...interface{}) Constraint {
	return c.Not(NewValueElementsConstraint(NewEqualComparator(options...), expectedValues...))
}

func (c *Controller) Empty() Constraint {
	return NewEmptyConstraint()
}

func (c *Controller) NotEmpty() Constraint {
	return c.Not(NewEmptyConstraint())
}

func (c *Controller) Equal(expected interface{}, options ...interface{}) Constraint {
	return NewEqualConstraint(expected, NewEqualComparator(options...))
}

func (c *Controller) NotEqual(expected interface{}, options ...interface{}) Constraint {
	return c.Not(NewEqualConstraint(expected, NewEqualComparator(options...)))
}

func (c *Controller) False() Constraint {
	return NewFalseConstraint()
}

func (c *Controller) Greater(expected interface{}) Constraint {
	return NewGreaterConstraint(expected)
}

func (c *Controller) GreaterOrEqual(expected interface{}, delta float64) Constraint {
	return c.Or(c.Greater(expected), c.Equal(expected, NumericDeltaOption{Value: delta}))
}

func (c *Controller) Less(expected interface{}) Constraint {
	return NewLessConstraint(expected)
}

func (c *Controller) LessOrEqual(expected interface{}, delta float64) Constraint {
	return c.Or(c.Less(expected), c.Equal(expected, NumericDeltaOption{Value: delta}))
}

func (c *Controller) Kind(allowedKinds ...reflect.Kind) Constraint {
	return NewKindConstraint(allowedKinds...)
}

func (c *Controller) NotKind(allowedKinds ...reflect.Kind) Constraint {
	return c.Not(NewKindConstraint(allowedKinds...))
}

func (c *Controller) Length(length int) Constraint {
	return NewLengthConstraint(length)
}

func (c *Controller) NotLength(length int) Constraint {
	return c.Not(c.Length(length))
}

func (c *Controller) LengthGreater(length int) Constraint {
	return NewLengthGreaterConstraint(length)
}

func (c *Controller) LengthGreaterOrEqual(length int) Constraint {
	return c.Or(c.LengthGreater(length), c.Length(length))
}

func (c *Controller) LengthLess(length int) Constraint {
	return NewLengthLessConstraint(length)
}

func (c *Controller) LengthLessOrEqual(length int) Constraint {
	return c.Or(c.LengthLess(length), c.Length(length))
}

func (c *Controller) Nil() Constraint {
	return NewNilConstraint()
}

func (c *Controller) NotNil() Constraint {
	return c.Not(NewNilConstraint())
}

func (c *Controller) Not(constraint Constraint) Constraint {
	return NewNotConstraint(constraint)
}

func (c *Controller) Or(constraint ...Constraint) Constraint {
	return NewOrConstraint(constraint...)
}

func (c *Controller) Regexp(pattern string) Constraint {
	return NewRegexpConstraint(pattern)
}

func (c *Controller) NotRegexp(pattern string) Constraint {
	return c.Not(NewRegexpConstraint(pattern))
}

func (c *Controller) Same(expected interface{}) Constraint {
	return NewSameConstraint(expected)
}

func (c *Controller) NotSame(expected interface{}) Constraint {
	return c.Not(NewSameConstraint(expected))
}

func (c *Controller) True() Constraint {
	return NewTrueConstraint()
}

func (c *Controller) Type(allowedTypes ...interface{}) Constraint {
	return NewTypeConstraint(allowedTypes...)
}

func (c *Controller) NotType(allowedTypes ...interface{}) Constraint {
	return c.Not(NewTypeConstraint(allowedTypes...))
}

func (c *Controller) Value(expected interface{}, options ...interface{}) Constraint {
	return NewValueConstraint(expected, NewEqualComparator(options...))
}

func (c *Controller) NotValue(expected interface{}, options ...interface{}) Constraint {
	return c.Not(c.Value(expected, options...))
}

func (c *Controller) AssertThat(value interface{}, constraint Constraint, context string) *Controller {
	c.test.Helper()

	if constraint == nil {
		c.test.Error(NewNotNilError("constraint"))
		c.test.Fail()

		return nil
	}

	if !constraint.Check(value) {
		c.test.Error(NewAssertError(value, context, constraint.String(), constraint.Details(value)))
		c.test.Fail()

		return nil
	}

	return c
}

func (c *Controller) AssertAnd(value interface{}, constraints ...Constraint) {
	c.test.Helper()

	c.AssertThat(value, c.And(constraints...), "'value' variable")
}

func (c *Controller) AssertCallback(value interface{}, callback func(value interface{}) bool) {
	c.test.Helper()

	c.AssertThat(value, c.Callback(callback), "'value' variable")
}

func (c *Controller) AssertNotCallback(value interface{}, callback func(value interface{}) bool) {
	c.test.Helper()

	c.AssertThat(value, c.NotCallback(callback), "'value' variable")
}

func (c *Controller) AssertContains(element interface{}, list interface{}) {
	c.test.Helper()

	c.AssertThat(list, c.Contains(element), "'list' variable")
}

func (c *Controller) AssertNotContains(element interface{}, list interface{}) {
	c.test.Helper()

	c.AssertThat(list, c.NotContains(element), "'list' variable")
}

func (c *Controller) AssertElements(constraints []Constraint, actual ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Elements(constraints...), "'actual' variable")
}

func (c *Controller) AssertNotElements(constraints []Constraint, actual ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.NotElements(constraints...), "'actual' variable")
}

func (c *Controller) AssertValueElements(expected []interface{}, actual []interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.ValueElements(expected, options...), "'actual' variable")
}

func (c *Controller) AssertNotValueElements(expected []interface{}, actual []interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.NotValueElements(expected, options...), "'actual' variable")
}

func (c *Controller) AssertEmpty(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Empty(), "'value' variable")
}

func (c *Controller) AssertNotEmpty(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.NotEmpty(), "'value' variable")
}

func (c *Controller) AssertEqual(expected interface{}, actual interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Equal(expected, options...), "'actual' variable")
}

func (c *Controller) AssertNotEqual(expected interface{}, actual interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.NotEqual(expected, options...), "'actual' variable")
}

func (c *Controller) AssertFalse(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.False(), "'value' variable")
}

func (c *Controller) AssertGreater(expected interface{}, actual interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Greater(expected), "'actual' variable")
}

func (c *Controller) AssertGreaterOrEqual(expected interface{}, actual interface{}, delta float64) {
	c.test.Helper()

	c.AssertThat(actual, c.GreaterOrEqual(expected, delta), "'actual' variable")
}

func (c *Controller) AssertLess(expected interface{}, actual interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Less(expected), "'actual' variable")
}

func (c *Controller) AssertLessOrEqual(expected interface{}, actual interface{}, delta float64) {
	c.test.Helper()

	c.AssertThat(actual, c.LessOrEqual(expected, delta), "'actual' variable")
}

func (c *Controller) AssertKind(value interface{}, expectedKinds ...reflect.Kind) {
	c.test.Helper()

	c.AssertThat(value, c.Kind(expectedKinds...), "'value' variable")
}

func (c *Controller) AssertNotKind(value interface{}, expectedKinds ...reflect.Kind) {
	c.test.Helper()

	c.AssertThat(value, c.NotKind(expectedKinds...), "'value' variable")
}

func (c *Controller) AssertLength(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Length(length), "'value' variable")
}

func (c *Controller) AssertNotLength(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.NotLength(length), "'value' variable")
}

func (c *Controller) AssertLengthGreater(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.LengthGreater(length), "'value' variable")
}

func (c *Controller) AssertLengthGreaterOrEqual(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.LengthGreaterOrEqual(length), "'value' variable")
}

func (c *Controller) AssertLengthLess(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.LengthLess(length), "'value' variable")
}

func (c *Controller) AssertLengthLessOrEqual(length int, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.LengthLessOrEqual(length), "'value' variable")
}

func (c *Controller) AssertNil(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Nil(), "'value' variable")
}

func (c *Controller) AssertNotNil(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.NotNil(), "'value' variable")
}

func (c *Controller) AssertNot(constraint Constraint, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Not(constraint), "'value' variable")
}

func (c *Controller) AssertOr(value interface{}, constraints ...Constraint) {
	c.test.Helper()

	c.AssertThat(value, c.Or(constraints...), "'value' variable")
}

func (c *Controller) AssertRegexp(pattern string, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Regexp(pattern), "'value' variable")
}

func (c *Controller) AssertNotRegexp(pattern string, value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.NotRegexp(pattern), "'value' variable")
}

func (c *Controller) AssertSame(expected interface{}, actual interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Same(expected), "'actual' variable")
}

func (c *Controller) AssertNotSame(expected interface{}, actual interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.NotSame(expected), "'actual' variable")
}

func (c *Controller) AssertTrue(value interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.True(), "'value' variable")
}

func (c *Controller) AssertType(value interface{}, expectedTypes ...interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.Type(expectedTypes...), "'value' variable")
}

func (c *Controller) AssertNotType(value interface{}, expectedTypes ...interface{}) {
	c.test.Helper()

	c.AssertThat(value, c.NotType(expectedTypes...), "'value' variable")
}

func (c *Controller) AssertValue(expected interface{}, actual interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.Value(expected, options...), "'actual' variable")
}

func (c *Controller) AssertNotValue(expected interface{}, actual interface{}, options ...interface{}) {
	c.test.Helper()

	c.AssertThat(actual, c.NotValue(expected, options...), "'actual' variable")
}
