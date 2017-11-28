package anagram

import (
	"sort"
	"strings"
)

const (
	testVersion = 2
)

type ByAlphabeticOrder []rune

func (s ByAlphabeticOrder) Len() int {
	return len(s)
}

func (s ByAlphabeticOrder) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByAlphabeticOrder) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ByAlphabeticOrder) Crescent(i, j int) bool {
	return s[i] < s[j]
}

func Detect(s string, ls []string) []string {
	var out []string
	s = strings.ToLower(s)
	sorted_original := []rune(s)
	sort.Sort(ByAlphabeticOrder(sorted_original))
	for _, v := range ls {
		c := strings.ToLower(v)
		if c != s {
			sorted_temp := []rune(c)
			sort.Sort(ByAlphabeticOrder(sorted_temp))
			if string(sorted_temp) == string(sorted_original) {
				out = append(out, v)
			}
		}
	}

	return out

}
