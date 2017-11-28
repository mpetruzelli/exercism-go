package allergies

const (
	testVersion = 1
)

var allergies = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(in uint) []string {
	o := make([]string, 0)

	for k := range allergies {
		if AllergicTo(in, k) {
			o = append(o, k)
		}
	}

	return o
}
func AllergicTo(in uint, s string) bool {
	return in&allergies[s] != 0
}
