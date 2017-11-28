package brackets

const (
	testVersion = 5
)

func Bracket(s string) (bool, error) {
	var counter [3]int
	var r []rune //opposite brackets

	for _, v := range s {
		switch v {

		case '{':
			counter[0]++
			r = append([]rune{'}'}, r...) //Prepend

		case '[':
			counter[1]++
			r = append([]rune{']'}, r...) //Prepend

		case '(':
			counter[2]++
			r = append([]rune{')'}, r...) //Prepend

		case '}':
			if len(r) > 0 && r[0] == '}' {
				counter[0]--
				r = r[1:] //remove first of array
			} else {
				return false, nil
			}
		case ']':
			if len(r) > 0 && r[0] == ']' {
				counter[1]--
				r = r[1:] //remove first of array
			} else {
				return false, nil
			}
		case ')':
			if len(r) > 0 && r[0] == ')' {
				counter[2]--
				r = r[1:] //remove first of array
			} else {
				return false, nil
			}

		}
	}

	for _, v := range counter {
		if v != 0 {
			return false, nil
		}
	}

	return true, nil
}
