package isogram
import "strings"
func IsIsogram(word string) bool {
    word = strings.ToLower(word)
    for pos, char := range word {
        for pos2, char2 := range word {
            if (char == char2 && pos != pos2) && (char != '-' && char != ' ') {
                return false
            }
        }
    }
    return true
}
