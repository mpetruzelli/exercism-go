package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const (
	testVersion = 2
)

func Encode(plain string) string {
	normalizedText := cleanText(plain)
	l := int(math.Ceil(math.Sqrt(float64(len(normalizedText)))))
	cipherText := make([]string, l)
	for i := 0; i < l; i++ {
		for j := i; j < len(normalizedText); j += l {
			cipherText[i] += string(normalizedText[j])
		}
	}
	return strings.Join(cipherText, " ")
}

func cleanText(text string) string {
	cleanText := ""
	for _, v := range strings.ToLower(text) {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			cleanText += string(v)
		}
	}
	return cleanText
}
