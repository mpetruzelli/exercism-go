package lsproduct

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	testVersion = 5
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return -1, fmt.Errorf("span can't be negative")
	}
	if len(digits) < span {
		return -1, fmt.Errorf("len of digits cant be smaller than span")
	}
	if hasLetters(digits) {
		return -1, fmt.Errorf("digits can't have letters")
	}
	if span == 0 {
		return 1, nil
	}
	largestserie := 0
	for i := 0; i <= len(digits)-span; i++ {
		serie := 1
		for _, v := range digits[i : i+span] {
			serie *= int(v) - 48
		}
		if serie > largestserie {
			largestserie = serie
		}
	}
	return largestserie, nil

}

func hasLetters(text string) bool {
	for _, v := range strings.ToLower(text) {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}
