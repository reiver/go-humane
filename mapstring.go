package humane

import (
	"strings"
)

// mapString is used to change the ordering of strings to a more human-friendly ordering.
//
// It really only exists to be used inside of friendly.cmpStrings(str1,str2)
func mapString(str string) string {
	return strings.Map(mapRune, str)
}
