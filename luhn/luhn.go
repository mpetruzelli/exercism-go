package luhn

import (
	"strings"
	"unicode"
)

const (
	testVersion = 2
)

func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 1 || !validString(s) {
		return false
	}

	sum := 0
	index := 1
	for i := len(s) - 1; i >= 0; i-- {
		if (index)%2 == 0 {
			double := (int(s[i]) - 48) * 2
			if double > 9 {
				double -= 9
				sum += double
			} else {
				sum += double
			}
		} else {
			sum += (int(s[i]) - 48)
		}
		index++
	}
	return sum%10 == 0
}
func validString(text string) bool {
	for _, v := range strings.ToLower(text) {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}
