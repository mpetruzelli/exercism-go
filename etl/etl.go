package etl

import (
	"strings"
)

const (
	testVersion = 1
)

func Transform(m map[int][]string) map[string]int {
	o := make(map[string]int)
	for key, value := range m {
		for _, letter := range value {
			o[strings.ToLower(letter)] = key
		}
	}
	return o
}
