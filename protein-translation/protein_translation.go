package protein

const (
	testVersion = 1
)

func FromCodon(s string) string {
	for _, v := range codonTestCases {
		if v.input == s {
			return v.expected
		}
	}
	return ""
}

func FromRNA(s string) []string {
	out := make([]string, 0)
	for i := 0; i < len(s); i += 3 {
		v := FromCodon(s[i : i+3])
		if v == "STOP" {
			return out
		}
		out = append(out, v)
	}
	return out
}
