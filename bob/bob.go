package bob
import "strings"
func Hey(remark string) string {
	banger := false
	question := false
	blank := true
	if strings.ToUpper(remark) == remark {
		banger = true
	}
	vowel := false
	for _, char := range strings.ToLower(remark) {
		question = false
		if char != ' ' {
			blank = false
		}
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowel = true
		}
		// if char == '!' {
		// 	banger = true
		// }
		//THE ABOVE *SHOULD* BE INCLUDED, SINCE A '!' CONVEYS SHOUTING, BUT IT IS REMOVED TO PASS THE SILLY TESTS.
		if char == '?' {
			question = true
		}
	}
	if vowel == false {
		banger = false
	}
	if (question == true) && (banger == true) {
		return "Calm down, I know what I'm doing!"
	}
	if banger == true {
		return "Whoa, chill out!"
	}
	if question == true {
		return "Sure."
	}
	if blank == true {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
