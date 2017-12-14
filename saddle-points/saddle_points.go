package matrix

//find \n. from there. you know the lenght of array.
//one func for max other for min
//compare pairs

type Pair struct {
	row    int
	column int
}

// Saddle finds the saddle points in a matrix
func (m Matrix) Saddle() (points []Pair) {
	cols := m.Cols()
	rows := m.Rows()
	for r, row := range rows {
		for c, cell := range row {
			if is(cell, row, max) && is(cell, cols[c], min) {
				points = append(points, Pair{r, c})
			}
		}
	}

	return points
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func is(el int, a []int, f func(int, int) int) bool {
	m := a[0]
	for _, i := range a {
		m = f(m, i)
	}
	return el == m
}
