package pythagorean

const (
	testVersion = 1
)

type Triplet [3]int

func Range(min, max int) (list []Triplet) {
	list = make([]Triplet, 0)
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					list = append(list, Triplet{i, j, k})
				}
			}
		}
	}
	return
}

func Sum(p int) (list []Triplet) {
	list = make([]Triplet, 0)
	for _, t := range Range(1, p) {
		a, b, c := t[0], t[1], t[2]
		if a+b+c == p {
			list = append(list, t)
		}
	}

	return

}
