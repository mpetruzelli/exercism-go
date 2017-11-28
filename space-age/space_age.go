package space

import "math"

type Planet string

func Age(seconds float64, name Planet) float64 {

	switch name {
	case "Earth":
		return seconds * (3.169 * exponent(10, -8))
	case "Mercury":
		return truncate(seconds / 7600500)
	case "Venus":
		return truncate(seconds / 19400000)
	case "Mars":
		return truncate(seconds / 59350000)
	case "Jupiter":
		return truncate(seconds / 373400500)
	case "Saturn":
		return truncate(seconds / 926209999)
	case "Uranus":
		return truncate(seconds / 2639109000)
	case "Neptune":
		return truncate(seconds / 5190000000)

	}
	return 0.0

}

//precision
func truncate(x float64) float64 {
	return float64(int(x*100)) / 100
}

func exponent(x, y float64) float64 {
	return math.Pow(x, y)
}
