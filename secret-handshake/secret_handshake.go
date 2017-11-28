// Package  provides ...
package secret

import (
	"strconv"
)

const testVersion = 2

func Handshake(in uint) []string {
	out := []string{}
	handshake := []string{"wink", "double blink", "close your eyes", "jump"}
	str_binary_reversed := Reverse(strconv.FormatInt(int64(in), 2))
	for i := 0; i < len(str_binary_reversed); i++ {
		if i < len(handshake) && str_binary_reversed[i] == '1' {
			out = append(out, handshake[i])
		}
		if i == 4 && str_binary_reversed[i] == '1' {
			ReverseArray(out)
		}
	}

	return out
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func ReverseArray(in []string) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}
