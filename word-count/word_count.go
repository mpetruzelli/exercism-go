package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Use this return type.
type Frequency map[string]int

// Just implement the function.
func WordCount(phrase string) (o Frequency) {
	o = make(map[string]int)
	splitted_str := strings.FieldsFunc(phrase, Split)
	for _, aword := range splitted_str {
		o[withoutPunctuation(aword)]++
	}
	return
}
func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '\n'
}

//without puctuation normalizes string too
func withoutPunctuation(s string) string {
	o := ""
	for _, v := range s {
		if !unicode.IsPunct(v) && !unicode.IsSymbol(v) {
			o += string(v)
		}
		if v == '\'' {
			o += string(v)
		}

	}
	if strings.Index(o, "'") == 0 || strings.Index(o, "'") == len(s)-1 {
		o = strings.Trim(o, "'")
	}
	return strings.ToLower(o)
}
