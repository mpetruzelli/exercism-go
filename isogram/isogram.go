package isogram

import "strings"

const (
	testVersion = 1
)

func IsIsogram(s string) bool {
	asciiTable := make([]int, 255)
	s = strings.Replace(strings.Replace(strings.ToLower(s), "-", "", -1), " ", "", -1)
	for _, r := range s {

		if asciiTable[r] == 1 {
			return false
		}
		asciiTable[r] = 1
	}
	return true
}
