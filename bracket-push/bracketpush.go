package brackets
func Bracket(seqeunce string) bool {
    bracket := false
    brace := false
    paranthesis := false
	bracketCount := 0
	braceCount := 0
	paranthesisCount := 0
    m := make(map[string]int)
    for pos, char := range seqeunce {
        if char == '{' {
        bracket = true
	    bracketCount += 1
        m["{"] = pos
        }
        if char == '[' {
        brace = true
   	    braceCount += 1
        m["["] = pos
        }
        if char == '(' {
        paranthesis = true
	    paranthesisCount += 1
        m["("] = pos
        }
        if char == '{' {
	    if ((brace == true) || (paranthesis == true)) && (bracket == false) {
		return false
        }
        if (brace == true) {
            if m["["] < m["{"] {
                return false
            }
        if (paranthesis == true) {
            if m["("] < m["{"] {
                return false
            }
        }
        }
	bracket = false
	bracketCount -= 1
        }
	if char == ']' {
	if ((bracket == true) || (paranthesis == true)) && (brace == false) {
		return false
        }
        if (bracket == true) {
            if m["{"] < m["["] {
                return false
            }
        }
        if (paranthesis == true) {
            if m["("] < m["["] {
                return false
            }
        }
	brace = false
	braceCount -= 1
        }
	if char == ')' {
	if ((bracket == true) || (brace == true)) && (paranthesis == false){
		return false
        }
        if (brace == true) {
            if m["["] < m["("] {
                return false
            }
        }
        if (bracket == true) {
            if m["{"] < m["("] {
                return false
            }
        }
	    paranthesis = false
	    paranthesisCount -= 1
        }
     }
    if ((paranthesis == false) && (paranthesisCount == 0)) && ((brace == false) && (braceCount == 0)) && ((bracket == false) && (bracketCount == 0)) {
	       return true
    }
    return false
}
