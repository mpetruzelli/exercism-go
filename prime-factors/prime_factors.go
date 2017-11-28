package prime

func Factors(n int64) []int64 {

	f := []int64{}

	if n == 1 {

		return f

	}
	i := 2
	for n != 1 {
		if n%int64(i) == 0 {
			f = append(f, int64(i))
			n = n / int64(i)

		} else {
			i += 1
		}

	}
	return f

}
