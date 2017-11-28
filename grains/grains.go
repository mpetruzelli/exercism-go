package grains

import (
	"errors"
)

const (
	testVersion = 1
)

//sumatory of 2^i for 0<=i<=63
/**
func Compounding(i int) uint64 {
	return uint64(math.Exp2(float64(i - 1)))
}*/

// another solution that i need to start using
// Shifting left by one is the same as multiplication by 2, and shifting right by one is the same as dividing by two, discarding any remainder.

func Compounding(i int) uint64 {
	return 1 << (uint(i) - 1)
}

func Square(i int) (uint64, error) {
	if i > 64 || i < 1 {
		return 0, errors.New("square positions from 1 to 64")
	}
	return Compounding(i), nil
}

func Total() (o uint64) {

	return (1 << 64) - 1
}
