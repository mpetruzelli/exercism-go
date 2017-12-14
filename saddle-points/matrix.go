package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

//Matrix bla bla
type Matrix struct {
	ms [][]string
	mi [][]int
}

//New bla bla
func New(s string) (*Matrix, error) {

	ms, mi, err := CreateMatrixes(s)
	if err != nil {
		return nil, err
	}
	return &Matrix{ms: ms, mi: mi}, nil
}

//CreateMatrixes blabla
func CreateMatrixes(s string) ([][]string, [][]int, error) {

	rows := strings.Split(s, "\n")

	ncolums := len(strings.Split(rows[0], " "))

	ms := make([][]string, len(rows))

	mi := make([][]int, len(rows))

	for i := range ms {

		ms[i] = make([]string, 0)

		mi[i] = make([]int, 0)

	}

	for i := 0; i < len(rows); i++ {

		colums := strings.Split(rows[i], " ")

		if len(colums) != ncolums {

			return nil, nil, fmt.Errorf("uneven colums")

		}

		for j := 0; j < ncolums; j++ {

			if len(colums[j]) > 0 {

				ms[i] = append(ms[i], colums[j])

				n, err := strconv.Atoi(colums[j])

				if err != nil {

					return nil, nil, fmt.Errorf("overflows int")

				}

				//mi[i][j] = n
				mi[i] = append(mi[i], n)
			}

		}

	}
	return ms, mi, nil
}

//Rows blabla
func (m *Matrix) Rows() [][]int {

	rows := make([][]int, len(m.ms))

	for i := 0; i < len(m.ms); i++ {

		for j := 0; j < len(m.ms[0]); j++ {

			n, _ := strconv.Atoi(m.ms[i][j])

			rows[i] = append(rows[i], n)

		}

	}

	return rows
}

//Cols blabla
func (m *Matrix) Cols() [][]int {

	colums := make([][]int, len(m.ms[0]))

	for i := 0; i < len(m.ms[0]); i++ {

		for j := 0; j < len(m.ms); j++ {

			n, _ := strconv.Atoi(m.ms[j][i])

			colums[i] = append(colums[i], n)

		}

	}

	return colums

}

//Set blabla
func (m *Matrix) Set(x, y, n int) bool {
	if x < 0 || y < 0 {

		return false

	} else if x >= len(m.ms) || y >= len(m.ms[0]) {

		return false

	}

	m.ms[x][y] = strconv.Itoa(n)

	return true
}
