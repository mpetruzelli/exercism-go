package triangle

import (
	"math"
	"sort"
)
const testVersion = 3



func KindFromSides(a, b, c float64) Kind{
	sides := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(sides))
	switch {
	case math.IsNaN(sides[0]) || sides[0] <= 0 || sides[2] == math.Inf(1) ||
		sides[1]+sides[0] < sides[2]:
		return NaT
	case sides[0] == sides[1] && sides[1] == sides[2]:
		return Equ
	case sides[0] == sides[1] || sides[1] == sides[2]:
		return Iso
	}
return Sca

}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	Equ = iota // equilateral
	Iso  
	Sca 
	NaT 
)
// Organize your code for readability.
