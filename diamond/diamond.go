package diamond

import (
	"fmt"
	"math"
)

const (
	testVersion = 1
)

func Gen(letter byte) (string, error) {
	n := int(letter)
	if n < 65 || n > 90 {
		return "", fmt.Errorf("index out of allowed range")
	}
	n -= 65
	return Atum(n), nil
}

func Atum(n int) string {
	var (
		diamond              string
		dots                 string
		letters              string
		ascii_decimal_letter int = 65
	)

	for i := n; i >= -n; i-- {

		// generate dots
		for j := 1; j <= int(math.Abs(float64(i))); j++ {
			dots += " "
		}
		// add first dots
		diamond += dots

		// add chars
		for j := 1; j <= 2*int(float64(n)-math.Abs(float64(i)))+1; j++ {
			if j == 1 || j == 2*int(float64(n)-math.Abs(float64(i)))+1 {
				letters += string(rune(ascii_decimal_letter))
			} else {
				letters += " "
			}
		}
		diamond += letters
		// add last dots
		diamond += dots

		// add the line to the general shape
		diamond += "\n"

		// clear variables
		dots = ""
		letters = ""
		// change the char
		if i > 0 {
			ascii_decimal_letter++
		} else {
			ascii_decimal_letter--
		}
	}
	fmt.Println(diamond)
	// shape is done, return it
	return diamond
}
