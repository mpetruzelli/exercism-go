package dna

import (
	"errors"
	"fmt"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

const validNucleotides = "ACGT"

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	if !strings.Contains(validNucleotides, string(nucleotide)) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}
	return strings.Count(string(d), string(nucleotide)), nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	t := 0
	h := Histogram{}
	for _, v := range validNucleotides {
		nucleotide := v
		h[nucleotide], _ = d.Count(byte(nucleotide))
		t += h[nucleotide]
	}
	if t != len(d) {
		return nil, fmt.Errorf("invalid nucleotide present")
	}
	return h, nil
}
