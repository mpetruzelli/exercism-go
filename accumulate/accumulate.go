package accumulate

const testVersion = 1

func Accumulate(i []string, f func(string) string) []string {
	o := make([]string, len(i))
	for index, i := range i {
		o[index] = f(i)
	}
	return o
}
