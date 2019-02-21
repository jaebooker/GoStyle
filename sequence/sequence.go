package main

func sequences(seqeunce string) bool {
    bracket := false
    brace := false
    paranthesis := false
	bracketCount := 0
	braceCount := 0
	paranthesisCount := 0
    for _, char := range seqeunce {
        if char == '{' {
            bracket = true
	    bracketCount += 1
        }
        if char == '[' {
            brace = true
   	    braceCount += 1
        }
        if char == '(' {
            paranthesis = true
	    paranthesisCount += 1
        }
        if (char == '}') {
	    if (brace == true) || (paranthesis == true) {
		return false
        }
	bracket = false
	bracketCount -= 1
        }
	if char == ']' {
	if (bracket == true) || (paranthesis == true) {
		return false
        }
	brace = false
	braceCount -= 1
        }
	if char == ')' {
	if (bracket == true) || (brace == true) {
		return false
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

// func main(seqeunce string) bool {
//     for pos, char range := seqeunce {
//         bracket = false
//         brace = false
//         paranthesis = false
//         if char == '{' {
//             bracket = true
//         }
//         if char == '{' {
//             brace = true
//         }
//         if char == '{' {
//             paranthesis = true
//         }
//         if char == '}' {
//
//         }
//     }
// }
