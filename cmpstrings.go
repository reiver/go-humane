package humane

// cmpStrings is used to change the ordering of strings to a more human-friendly ordering.
func cmpStrings(s1 string, s2 string) int {
	var m1 string = mapString(s1)
	var m2 string = mapString(s2)

	switch {
	case m1 < m2:
		return -1
	case m1 > m2:
		return 1
	default:
		return 0
	}
}
