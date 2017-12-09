package wordy

import (
	"strconv"
	"strings"
)

var operation = map[string]string{
	"plus":       "+",
	"minus":      "-",
	"multiplied": "*",
	"divided":    "/",
	"raised":     "e",
}

func Answer(s string) (int, bool) {
	r := strings.Replace(s, "?", "", -1)
	t := strings.Fields(r)
	result := 0
	numbers := []int{}
	perform := []string{}

	for _, v := range t {
		if n, err := strconv.Atoi(v); err == nil {
			numbers = append(numbers, n)
		}
		if op, ok := operation[v]; ok == true {
			perform = append(perform, op)
		}
	}

	if len(perform) == 0 {
		return result, false
	}
	for i := 0; i < len(perform); i++ {

		switch perform[i] {

		case "+":
			result += numbers[i] + numbers[i+1]
		case "-":
			result += numbers[i] - numbers[i+1]
		case "*":
			result += numbers[i] * numbers[i+1]
		case "/":
			result += numbers[i] / numbers[i+1]
		case "e":
			result += numbers[i] * numbers[i+1]
		}

	}
	return result, true
}
