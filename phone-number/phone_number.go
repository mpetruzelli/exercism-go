package phonenumber

import "fmt"

func Number(s string) (string, error) {
	r := make([]rune, 0)
	for _, v := range s {

		if v >= 48 && v <= 57 {
			r = append(r, v)
		}

	}

	if r[0] < 50 || r[0] > 57 {

		r = r[1:]

	}

	if len(r) != 10 {

		return "", fmt.Errorf("invalid number")

	}

	if r[3] < 50 || r[3] > 57 {

		return "", fmt.Errorf("invalid exchange code")

	}

	return string(r), nil
}

func AreaCode(s string) (string, error) {
	n, err := Number(s)
	if err != nil {
		return "", err
	}
	return n[:3], nil
}

func Format(s string) (string, error) {

	n, err := Number(s)

	if err != nil {
		return "", err
	}
	return "(" + n[:3] + ") " + n[3:6] + "-" + n[6:], nil

}
