package unit

func NewValueConstraint(expected interface{}, comparator EqualComparer) Constraint {
	if comparator == nil {
		panic(NewNotNilError("comparator"))
	}

	if expected == nil {
		return NewNilConstraint()
	}

	if constraint, ok := expected.(Constraint); ok {
		return constraint
	}

	if constraint, ok := expected.(ConstraintAsValue); ok {
		return NewEqualConstraint(constraint.Value, comparator)
	}

	return NewEqualConstraint(expected, comparator)
}
