package pascal

func pascal(number int) [][]int triangle {
    num := 2
    intArray := [num][num]int{}
    for i:=1;i<num;i++ {
        for n:=0;n<i;n++ {
            temp := intArray[i-1][n]
            if intArray[i-1][n+1] != nil {
                temp2 := intArray[i-1][n+1]
            } else {
                temp2 := 0
            }
            intArray[i][n] = [temp1 + temp2]
        }
    }
}
