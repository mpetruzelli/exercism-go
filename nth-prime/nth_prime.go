package prime

import "math"

// this solution belongs to richardmarshall. After testing performance of my first solution agains solutions of others, this one was the fastest by far!! impressive solution.

// Nth returns out the nth prime, and returns 0 / not ok if n < 1.

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}
	m := 12
	if n >= 6 {
		f := float64(n)
		m = int(f*math.Log(f) + f*math.Log(math.Log(f)))
	}
	return sieveOfEratosthenes(m)[n-1], true
}

func sieveOfEratosthenes(m int) []int {
	p := []int{2}
	np := make([]bool, m)
	for i := 3; i < m; i += 2 {
		if !np[i] {
			p = append(p, i)
			for y := i * i; y < m; y += i {
				np[y] = true
			}
		}
	}
	return p
}
