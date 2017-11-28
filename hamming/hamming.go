package hamming
import "errors"
const testVersion = 6


func Distance(a, b string) (int,error) {
	
	if (len(a) != len(b)){
		return -1, errors.New("Two strings do not match.")
	}
	
	distance := 0

	for i:=0; i< len(a) && i< len(b); i++ {
		if a[i]!=b[i]{
			distance++;
		}
	}
	return distance,nil;

}



