package raindrops

import "strings" 
import "strconv"
const testVersion = 3

func Convert(i int) string{
  	s := make([]string, 0)
	if i%3==0{
		s=append(s,"Pling");
	}
	if i%5==0{
		s=append(s,"Plang");
	}
	if i%7==0{
		s=append(s,"Plong");
	}
	if(len(s)==0){
		return strconv.Itoa(i);
	}
	return strings.Join(s,"");
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
