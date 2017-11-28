package strain

const (
	testVersion = 1
)

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(pass_test func(int) bool) (o Ints) {
	for _, number := range ints {
		if pass_test(number) {
			o = append(o, number)
		}
	}
	return
}
func (lists Lists) Keep(pass_test func([]int) bool) (o Lists) {
	for _, list := range lists {
		if pass_test(list) {
			o = append(o, list)
		}
	}

	return
}
func (strings Strings) Keep(pass_test func(string) bool) (o Strings) {
	for _, astring := range strings {
		if pass_test(astring) {
			o = append(o, astring)
		}
	}

	return
}
func (ints Ints) Discard(pass_test func(int) bool) (o Ints) {
	for _, number := range ints {
		if !pass_test(number) {
			o = append(o, number)
		}
	}

	return
}
