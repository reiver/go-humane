package humane

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// cmpStrings is used to change the ordering of strings to a more human-friendly ordering.
func cmpStrings(str1 string, str2 string) int {

	var len1 int = len(str1)
	var len2 int = len(str2)

	// Short-Cut
	{
		switch {
		case 0 < len1 && 0 < len2:
			r1, _ := utf8.DecodeRuneInString(str1)
			r2, _ := utf8.DecodeRuneInString(str2)

			var l1 rune = unicode.ToLower(r1)
			var l2 rune = unicode.ToLower(r2)

			m1 := mapRune(l1)
			m2 := mapRune(l2)

			switch {
			case m1 < m2:
				return -1
			case m1 > m2:
				return 1
			}
		}
	}

	var lower1 string = strings.ToLower(str1)
	var lower2 string = strings.ToLower(str2)

	var minlength int = len1
	if len2 < minlength {
		minlength = len2
	}

	var cutlow1 string = lower1[:minlength]
	var cutlow2 string = lower2[:minlength]

	var mcutlow1 string = mapString(cutlow1)
	var mcutlow2 string = mapString(cutlow2)

	switch {
	case  mcutlow1 != mcutlow2:
		return strings.Compare(mcutlow1,mcutlow2)
	default:
		switch {
		case len1 < len2:
			return -1
		case len1 > len2:
			return 1
		default:
			return strings.Compare(mapString(str1),mapString(str2))
		}
	}
}
