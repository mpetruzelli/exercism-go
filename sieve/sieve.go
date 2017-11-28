package sieve

const (
	testVersion = 1
)

func Sieve(max int) (o []int) {
	var (
		sieve  []bool = make([]bool, max+1)
		primes []int  = make([]int, 0)
	)
	for i := 2; i <= max; i++ {

		if sieve[i] == false { // i has not been marked -- it is prime
			primes = append(primes, i)
			for j := i << 1; j <= max; j += i {
				sieve[j] = true
			}
		}
	}
	return primes

}
