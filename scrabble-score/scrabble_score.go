package scrabble

import (
	"strings"
)

const (
	testVersion = 5
)

type scrabble_points struct {
	letters string
	points  int
}

var scrabbleValues = []scrabble_points{
	{"aeioulnrst", 1},
	{"dg", 2},
	{"bcmp", 3},
	{"fhvwy", 4},
	{"k", 5},
	{"jx", 8},
	{"qz", 10},
}

func Score(s string) int {
	total_points := 0
	for _, r := range s {
		c := strings.ToLower(string(r))
		for _, rule := range scrabbleValues {
			if strings.ContainsAny(rule.letters, c) {
				total_points += rule.points
			}
		}
	}
	return total_points
}
