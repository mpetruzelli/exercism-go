package flatten

func Flatten(f interface{}) []interface{} {

	output := []interface{}{}

	switch e := f.(type) {
	case []interface{}:
		for _, v := range e {
			output = append(output, Flatten(v)...)
		}
	case interface{}:
		output = append(output, e)
	}

	return output

}
