package wordy

import (
	"fmt"
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

	}
	fmt.Println(result)
	return result, true
}
