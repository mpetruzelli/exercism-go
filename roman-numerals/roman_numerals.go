package romannumerals

import "fmt"

const (
	testVersion = 4
)

var roman2arabic = []struct {
	roman  string
	arabic int
}{
	{"I", 1},
	{"IV", 4},
	{"V", 5},
	{"IX", 9},
	{"X", 10},
	{"XL", 40},
	{"L", 50},
	{"XC", 90},
	{"C", 100},
	{"CD", 400},
	{"D", 500},
	{"CM", 900},
	{"M", 1000},
}

func Join(a []string, sep string) string {

	switch len(a) {

	case 0:

		return ""

	case 1:

		return a[0]

	case 2:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1]

	case 3:

		// Special case for common small values.

		// Remove if golang.org/issue/6714 is fixed

		return a[0] + sep + a[1] + sep + a[2]

	}

	n := len(sep) * (len(a) - 1)

	for i := 0; i < len(a); i++ {

		n += len(a[i])

	}

	b := make([]byte, n)

	bp := copy(b, a[0])

	for _, s := range a[1:] {

		bp += copy(b[bp:], sep)

		bp += copy(b[bp:], s)

	}

	return string(b)

}

func ToRomanNumeral(arg int) (string, error) {
	out := []string{}
	if arg <= 0 || arg > 3000 {
		return "", fmt.Errorf("Number needs to be between 1 and 3000")
	}
	i := len(roman2arabic) - 1
	for i >= 0 {
		if roman2arabic[i].arabic <= arg {
			out = append(out, roman2arabic[i].roman)
			arg -= roman2arabic[i].arabic
		} else {
			i -= 1
		}
	}
	return Join(out, ""), nil
}
