// package main
//
// import (
//     "testing"
// )
//
// func BenchmarkCalculate(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         Calculate(2)
//     }
// }
// package main
// import (
// 	"testing"
// )
//
// func TestCalculate(t *testing.T) {
// 	if Calculate(2) != 4 {
// 		t.Error("Expected 2 + 2 to equal 4")
// 	}
// }
// package main
// func TestTableCalculate(t *testing.T) {
// 	var tests = []struct {
// 		input    int
// 		expected int
// 	}{
// 		{2, 4},
// 		{-1, 1},
// 		{0, 2},
// 		{-5, -3},
// 		{99999, 100001},
// 	}
//
// 	for _, test := range tests {
// 		if output := Calculate(test.input); output != test.expected {
// 			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
// 		}
// 	}
// }
//package main
//
// import (
// 	"fmt"
// 	"testing"
// )
//
// func TestCalculate(t *testing.T) {
// 	fmt.Println("Test Calculate")
// 	expected := 4
// 	result := Calculate(2)
// 	if expected != result {
// 		t.Error("Failed")
// 	}
// }
//
// func BenchmarkCalculate(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Calculate(2)
// 	}
// }
//
// func TestOther(t *testing.T) {
// 	fmt.Println("Testing something else")
// 	fmt.Println("This shouldn't run with -run=calc")
// }
package main

import (
	"testing"
)

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
