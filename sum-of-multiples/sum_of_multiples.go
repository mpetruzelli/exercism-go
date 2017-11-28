// Package summultiples
package summultiples

const (
	testVersion = 2
)

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func SumMultiples(limit int, divisors ...int) int {
	o := make([]int, 0)

	for _, v := range divisors {
		for i := 1; i < limit; i++ {
			if i%v == 0 {
				o = AppendIfMissing(o, i)
			}
		}

	}

	total := 0
	for _, v := range o {
		total += v
	}
	return total
}
