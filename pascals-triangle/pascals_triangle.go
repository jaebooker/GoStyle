package pascal

// func pascal(number int) [][]int triangle {
//     num := 2
//     intArray := [num][num]int{}
//     for i:=1;i<num;i++ {
//         for n:=0;n<i;n++ {
//             temp := intArray[i-1][n]
//             if intArray[i-1][n+1] != nil {
//                 temp2 := intArray[i-1][n+1]
//             } else {
//                 temp2 := 0
//             }
//             intArray[i][n] = [temp1 + temp2]
//         }
//     }
// }
func triangleHelper(n int) []int {
	result := make([]int, n+1)
	result[0] = 1
	result[n] = 1
	for i, j := 1, n-1; i <= j; i, j = i+1, j-1 {
		result[i] = result[i-1] * (n + 1 - i) / i
		result[j] = result[i]
	}
	return result
}
func Triangle(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = triangleHelper(i)
	}
	return result
}
