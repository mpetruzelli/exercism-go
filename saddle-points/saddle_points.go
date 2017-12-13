package matrix

import (
	"strconv"
	"strings"
)

type Pair [2]int
type Matrix [][]int

func New(s string) (Matrix, error) {

	var m Matrix = [][]int{{}}
	ms := strings.Split(s, "\n")
	for i, v := range ms {
		numbers := strings.Split(v, " ")
		for j := range numbers {
			n, err := strconv.Atoi(numbers[j])
			if err != nil {
				return nil, err
			}
			m[i][j] = n
		}
	}
	return m, nil
}

func (m *Matrix) Sanddle() []Pair {
	tmpRow := -128
	tmpColum:=127
	for i := 0; i < len((*m)); i++ {
		for j := 0; j < (*m); j++ {
			if (*m)[i][j] >= tmpRow {
				
				tmpRow = (*m)[i][j]

			}
			if(*m)[j][i] <= tmpColum{
				tmpColum=
			}
		}
	}

}
