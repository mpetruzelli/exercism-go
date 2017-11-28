package acronym

import "strings"
import "bytes"
const testVersion = 3

func Abbreviate(s string) string{
	a := strings.FieldsFunc(s, Split)	
	o := make([]byte,len(a));
	
	for index, a:= range a{
		o[index]=a[0];
	}
	return string(bytes.ToUpper(o));
}
func Split(r rune) bool {
    return r == '-' || r == ' '
}
