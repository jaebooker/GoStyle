package acronym
import (
	"strings"
	"unicode"
)
// func abbreviate(s string) string {
// 	acronym := ""
// 	for i, l := range s {
// 		sl := strings.ToUpper(string(l))
// 		if i == 0 {
// 			acronym += sl
// 			continue
// 		}
// 		prev := rune(s[i-1])
// 		curr := rune(s[i])
// 		letterAfterNonLetter := unicode.IsLetter(curr) && !unicode.IsLetter(prev)
// 		upperAfterNonUpper := unicode.IsUpper(curr) && !unicode.IsUpper(prev)
// 		if letterAfterNonLetter || upperAfterNonUpper {
// 			acronym += sl
// 		}
// 	}
// 	return acronym
// }
func Abbreviate(s string) string {
	acronym := ""
	for _, v := range strings.FieldsFunc(s, split) {
		acronym += strings.ToUpper(string(v[0]))
	}
	return acronym
}
func split(c rune) bool {
	return unicode.In(c, unicode.Dash, unicode.Space)
}
