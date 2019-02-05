package isogram

func isogram(word string) bool {
    for pos, char := range word {
        for pos2, char2 := range word {
            if (char == char2 && pos != pos2) && (char != '-' && char != ' ') {
                return false
            }
        }
    }
    return true
}
