package sublist

type Relation string

//Sublist returns the relation between two lists
func Sublist(x, y []int) Relation {

	switch {

	case len(x) == len(y):

		if testEq(x, y) == true {

			return "equal"

		} else {

			return "unequal"

		}

	case len(x) == 0 && len(y) > 0:

		return "sublist"

	case len(x) > 0 && len(y) == 0:

		return "superlist"

	case len(x) < len(y):

		if SubList(x, y) {

			return "sublist"

		} else {

			return "unequal"

		}
	case len(x) > len(y):

		if SubList(y, x) {
			return "superlist"
		} else {
			return "unequal"
		}
	}
	return "consecutive"
}

//Sublist verifies if s is a sublist of f
func SubList(f, s []int) bool {

	for i := 0; i < len(s); i++ {

		if (i + len(f)) > len(s) {

			break

		}

		if testEq(f, s[i:i+(len(f))]) {

			return true

		}

	}

	return false
}

//TestEq tests Equility between two integer slices
func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
