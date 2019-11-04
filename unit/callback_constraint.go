package unit

type CallbackConstraint struct {
	callback func(value interface{}) bool
}

func NewCallbackConstraint(callback func(value interface{}) bool) Constraint {
	if callback == nil {
		panic(NewNotNilError("callback"))
	}

	return &CallbackConstraint{callback: callback}
}

func (c *CallbackConstraint) Check(value interface{}) bool {
	return c.callback(value)
}

func (c *CallbackConstraint) String() string {
	return "accept callback"
}

func (c *CallbackConstraint) Details(value interface{}) string {
	return ""
}
