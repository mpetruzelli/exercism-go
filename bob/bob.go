package bob

import (
	"strings"
)
const testVersion = 3


func Hey(s string) string{
	trs:= strings.TrimSpace(s);
	if trs == strings.ToUpper(trs) && trs != strings.ToLower(trs){
		return "Whoa, chill out!"
	}
	if trs=="" {
		return "Fine. Be that way!"
	}

	if trs[len(trs)-1:]=="?" {
		return "Sure.";
	}
	return "Whatever.";
		
}

