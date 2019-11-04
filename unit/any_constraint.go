package unit

type AnyConstraint struct {
}

func NewAnyConstraint() Constraint {
	return &AnyConstraint{}
}

func (c *AnyConstraint) Check(value interface{}) bool {
	return true
}

func (c *AnyConstraint) String() string {
	return "be anything"
}

func (c *AnyConstraint) Details(value interface{}) string {
	return ""
}
