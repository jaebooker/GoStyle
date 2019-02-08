package reverse

func reverseString(stringy string) string{
    b := []byte(stringy)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
    stringy2 := ""
    for i:=len(stringy)-1;i>0;i--{
        stringy2 += stringy[i]
    }
    return stringy2
}
