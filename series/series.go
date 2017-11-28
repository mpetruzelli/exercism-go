package series

const (
	testVersion = 2
)

func NumberOfDigitSeries(n int, s string) (o int) {

	for i := 0; i < len(s); i++ {
		if i+n <= len(s) {
			o++
		}
	}
	return
}

func All(n int, s string) []string {
	var slice []string
	for i := 0; i < len(s); i++ {
		if i+n <= len(s) {

			slice = append(slice, s[i:n+i])

		}
	}
	return slice

	//	size := go NDigitSeries(n, s)

}
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		first = s[:n]
		ok = true
		return
	} else {
		first = ""
		ok = false
		return
	}
}
