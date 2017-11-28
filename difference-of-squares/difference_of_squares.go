package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	out := MakeRange(0, n)
	sum := 0
	for _, value := range out {
		sum += value
	}
	return sum * sum
}

func MakeRange(min, max int) []int {
	o := make([]int, max-min+1)
	for i := range o {
		o[i] = min + i
	}
	return o
}

func SumOfSquares(n int) int {
	out := MakeRange(0, n)
	sum := 0
	for _, value := range out {
		sum += value * value
	}
	return sum

}
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
