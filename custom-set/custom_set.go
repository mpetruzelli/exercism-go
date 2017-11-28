package stringset

import (
	"strings"
)

type Set map[string]bool

func New() Set {

	return Set{}

}

func NewFromSlice(slice []string) Set {

	if len(slice) == 0 {
		return New()
	}

	s := make(map[string]bool)

	for _, v := range slice {

		s[v] = true

	}
	return Set(s)

}

func Equal(s1, s2 Set) bool {

	if len(s1) != len(s2) {

		return false

	}

	return Subset(s1, s2)
}

func Disjoint(s1, s2 Set) bool {
	for k, _ := range s1 {
		if s2[k] == true {
			return false
		}
	}
	return true

}

func Intersection(s1, s2 Set) Set {

	m := make(map[string]bool)

	for k, _ := range s1 {

		if s2[k] == true {

			m[k] = true

		}

	}

	return Set(m)

}

func Difference(s1, s2 Set) Set {
	m := make(map[string]bool)

	for k, _ := range s1 {

		if s2[k] == false {

			m[k] = true

		}

	}

	return Set(m)

}

func Union(s1, s2 Set) Set {
	for k, _ := range s1 {
		s2[k] = true
	}
	return s2
}

func Subset(s1, s2 Set) bool {
	flag := true
	for k, _ := range s1 {

		flag = flag && s2[k]

	}
	return flag

}

func (s Set) Add(st string) {
	s[st] = true
}

func (s Set) String() string {

	pieces := make([]string, 0)

	for el := range s {

		pieces = append(pieces, "\""+el+"\"")

	}

	return "{" + strings.Join(pieces, ", ") + "}"

}

func (s Set) IsEmpty() bool {

	return len(s) == 0

}

func (s Set) Has(st string) bool {
	return s[st]
}
