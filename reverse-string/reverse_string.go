package reverse

func String(s string) string{
    //b := []byte(stringy)
	// for i := range b {
	// 	b[i] := charset[source.Int63()%int64(len(charset))]
	// }
	// return string(b)
    // stringy2 := []rune(stringy)
    // i:=len(stringy)
    // for n:=0;n<i;n++{
    //     stringy2[n],stringy2[i] = stringy2[i],stringy2[n]
    //     i--
    // }
    // return string(stringy2)
    // stringRunes := []rune(stringy)
    // for i, j := 0, len(stringRunes)-1; i < j; i, j = i+1, j-1 {
    // runes[i], runes[j] = runes[j], runes[i]
    // }
    // return string(runes)
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
    }
