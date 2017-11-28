package pangram

import "unicode"
const testVersion = 1

func IsPangram(i string) bool{
	// Use a map as a set
	cases := map[rune]bool{}

	for _, c := range i {
		// Change uppercase to lowercase
		if unicode.IsUpper(c) {
			c = unicode.ToLower(c)
		}
		// Exclude anything that isn't a lowercase char
		if c < 'a' || c > 'z' {
			continue
		}
		cases[c] = true
	}

	// This should equal exactly 26 since we excluded every other character
	return len(cases) == 26
	

}
