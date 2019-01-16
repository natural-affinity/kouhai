package spec

// IsInvalidError compares errors
func IsInvalidError(actual error, expected error) bool {
	a := (actual != nil && expected != nil && actual.Error() != expected.Error())
	b := (actual == nil && expected != nil)

	return a || b
}
