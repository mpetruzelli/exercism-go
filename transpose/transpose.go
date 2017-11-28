package transpose

func Transpose(s []string) []string {

	if len(s) == 0 {
		return []string{}
	}

	n := SizeOfLargestString(s)

	o := make([]string, n)

	for i, v := range s {

		for jj, vv := range v {
			o[jj] += string(vv)
		}
		rm := SizeOfLargestString(s[i:])
		for j := len(v); j < rm; j++ {
			o[j] += " "
		}
	}
	return o
}

func SizeOfLargestString(s []string) (o int) {
	o = 0
	for _, v := range s {
		if len(v) > o {
			o = len(v)
		}
	}
	return o
}
