package queenattack

import (
	"errors"
	"math"
	"regexp"
)

const testVersion = 2

var validSquareExp = regexp.MustCompile(`^[a-h][1-8]$`)

func validSquare(s string) bool {
	return validSquareExp.MatchString(s)
}

func OffBoard(q1r, q2r, q1c, q2c float64) bool {
	return q1r < 1 || q2r < 1 || q1c < 1 || q2c < 1 || q1r > 8 || q2r > 8 || q1c > 8 || q2c > 8
}

func SameRow(q1r, q2r float64) bool {
	return q1r == q2r
}
func SameColum(q1c, q2c float64) bool {
	return q1c == q2c
}
func InDiagonal(q1r, q2r, q1c, q2c float64) bool {
	return math.Abs(q1c-q2c) == math.Abs(q1r-q2r)

}

//LookForInputErrors checks if the two queens arent in the same position and if their positions are inside the board
func LookForInputErrors(q1r, q2r, q1c, q2c float64) error {

	if q1r == q2r && q1c == q2c {
		return errors.New("Same position")
	}

	if OffBoard(q1r, q2r, q1c, q2c) {
		return errors.New("Off Board")
	}

	return nil
}

func CanQueenAttack(q1, q2 string) (bool, error) {

	if !(validSquare(q1) && validSquare(q2)) {
		return false, errors.New("Not a valid square")
	}

	q2r := float64(q2[0]-'a') + 1
	q1r := float64(q1[0]-'a') + 1
	//retrive colum number
	q1c := float64(q1[1] - '0')
	q2c := float64(q2[1] - '0')

	err := LookForInputErrors(q1r, q2r, q1c, q2c)

	if err != nil {
		return false, err
	}
	// transform row letter into a string number from 1-8

	if SameRow(q1r, q2r) || SameColum(q1c, q2c) || InDiagonal(q1r, q2r, q1c, q2c) {
		return true, nil
	}
	return false, nil
}
