package pascal

const (
	testVersion = 1
)

//O(n^2) time and O(1) extra space
func Triangle(n int) (o [][]int) {
	o = make([][]int, n)
	for r := 0; r < n; r++ {
		o[r] = make([]int, r+1)
		v := 1
		for c := 0; c <= r; c++ {
			o[r][c] = v
			v = v * ((r + 1) - (c + 1)) / (c + 1) //factorial
		}
	}
	return o
}
