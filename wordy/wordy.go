package wordy

import (
	"strconv"
	"strings"
)

var operation = map[string]bool{
	"plus":       true,
	"minus":      true,
	"multiplied": true,
	"divided":    true,
	"raised":     true,
}

func Answer(s string) (int, bool) {
	r := strings.Replace(s, "?", "", -1)
	t := strings.Fields(r)
	result := 0
	numbersandoperations := []string{}

	for _, v := range t {
		if _, err := strconv.Atoi(v); err == nil {
			numbersandoperations = append(numbersandoperations, v)
		}
		if _, ok := operation[v]; ok == true {
			numbersandoperations = append(numbersandoperations, v)
		}
	}

	if len(numbersandoperations) < 3 {
		return result, false
	}
	var op string
	for k, v := range numbersandoperations {

		if k%2 == 0 {
			n, err := strconv.Atoi(v)
			if err != nil {

			} else {

				switch op {
				case "plus":
					result += n
					break
				case "minus":
					result -= n
					break
				case "multiplied":
					result *= n
					break
				case "divided":
					result /= n
					break
				case "raised":
					break
				default:
					result += n

				}
			}
		} else {
			op = v
		}

	}

	return result, true
}
