package unit

import (
	"fmt"
	"reflect"
	"regexp"
)

type RegexpConstraint struct {
	pattern *regexp.Regexp
}

func NewRegexpConstraint(pattern string) Constraint {
	return &RegexpConstraint{
		pattern: regexp.MustCompile(pattern),
	}
}

func (c *RegexpConstraint) Check(value interface{}) bool {
	valueValue := reflect.ValueOf(value)

	if valueValue.Kind() != reflect.String {
		panic(NewInvalidKindError("value", value, reflect.String))
	}

	return c.pattern.MatchString(valueValue.String())
}

func (c *RegexpConstraint) String() string {
	return fmt.Sprintf("match PCRE pattern '%s'", c.pattern.String())
}

func (c *RegexpConstraint) Details(value interface{}) string {
	return ""
}
