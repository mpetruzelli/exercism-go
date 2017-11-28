package strand

const (
	testVersion = 3
)

var translation = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(s string) string {
	out := []rune{}
	for _, v := range s {
		out = append(out, translation[v])
	}
	return string(out)
}
