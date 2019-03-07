package zerba
// type Solution struct {
// 	DrinksWater string
// 	OwnsZebra string
// }
// type House struct {
//     smoke string
//     drink string
//     pet string
//     nationality string
//     color string
//     location int
// }
// func SolvePuzzle()Solution{
//     color := [5]string{"red","ivory","yellow","green","blue"}
//     person := [5]string{"Japanese","Ukrainian","Englishman","",""}
//     pet := [5]string{"","","","",""}
//     smoke := [5]string{"","","","",""}
//     drink := [5]string{"water","orange juice","coffee","",""}
//     h1 := House{}
//     h2 := House{}
//     h3 := House{}
//     h4 := House{}
//     h5 := House{}
//     h1.nationality = person[2]
//     h1.color = color[0]
//     h1
// }
//

//h = make(map[string][House])
// var h [5]House
// h := Houses{}
// h.nationalities = "Englishman"
// h1.color = "Red"
// h2.nationality = "Spaniard"
// h2.pet = "Dog"

// type Houses struct {
//     smokes Smoke
//     drinks Drink
//     pets Pet
//     nationalities Nationality
//     colors Color
// }
// type Nationality struct {
//     marlboros string
//     menthols string
// }
// type Drink struct {
//     water string
//     oj string
// }
// type Pet struct {
//     water string
//     oj string
// }
// type Smoke struct {
//     water string
//     oj string
// }
// type Color struct {
//     water string
//     oj string
// }

//new version of zebra (not mine)

import "bytes"

//Solution structure to return solution in consistent way
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

//generate all permuations from byte slice, code from stack exchange
func permutations(arr []byte) [][]byte {
	var helper func([]byte, int)
	res := [][]byte{}

	helper = func(arr []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

//Absolute value function
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Map bytes to full nation names for the solution
var nationMap = map[byte]string{'E': "Englishman", 'S': "Spaniard", 'U': "Ukranian", 'J': "Japanese", 'N': "Norwegian"}

//SolvePuzzle solves the zebra puzzle
func SolvePuzzle() Solution {
	nations := []byte{'E', 'S', 'U', 'J', 'N'}
	drinks := []byte{'w', 'c', 't', 'o', 'm'}
	houses := []byte{'r', 'g', 'b', 'y', 'i'}
	smokes := []byte{'o', 'k', 'c', 'l', 'p'}
	pets := []byte{'z', 's', 'f', 'h', 'd'}
	for _, n := range permutations(nations) {
		if bytes.IndexByte(n, 'N') != 0 {
			//Norwegian lives in first house
			continue
		}
		for _, d := range permutations(drinks) {
			if bytes.IndexByte(d, 'm') != 2 {
				//Milk in middle house
				continue
			}
			if bytes.IndexByte(n, 'U') != bytes.IndexByte(d, 't') {
				//Ukrainian drinks tea
				continue
			}
			for _, h := range permutations(houses) {
				if bytes.IndexByte(h, 'g') != bytes.IndexByte(h, 'i')+1 {
					//green is right of ivory
					continue
				}
				if bytes.IndexByte(n, 'E') != bytes.IndexByte(h, 'r') {
					//English in red
					continue
				}
				if bytes.IndexByte(d, 'c') != bytes.IndexByte(h, 'g') {
					//coffee drinker in green
					continue
				}
				if abs(bytes.IndexByte(n, 'N')-bytes.IndexByte(h, 'b')) != 1 {
					//norweigain next to blue
					continue
				}
				for _, s := range permutations(smokes) {
					if bytes.IndexByte(h, 'y') != bytes.IndexByte(s, 'k') {
						//kools smoker in yellow
						continue
					}
					if bytes.IndexByte(s, 'l') != bytes.IndexByte(d, 'o') {
						//lucky strike drinks orange
						continue
					}
					if bytes.IndexByte(s, 'p') != bytes.IndexByte(n, 'J') {
						//Japanese smokes parliaments
						continue
					}
					for _, p := range permutations(pets) {
						if bytes.IndexByte(s, 'o') != bytes.IndexByte(p, 's') {
							//old gold has snails
							continue
						}
						if bytes.IndexByte(n, 'S') != bytes.IndexByte(p, 'd') {
							//spaniard has dog
							continue
						}
						if abs(bytes.IndexByte(s, 'c')-bytes.IndexByte(p, 'f')) != 1 {
							//chesterfieled next to fox
							continue
						}
						if abs(bytes.IndexByte(s, 'k')-bytes.IndexByte(p, 'h')) != 1 {
							//kools next to horse
							continue
						}
						return Solution{
							DrinksWater: nationMap[n[bytes.IndexByte(d, 'w')]],
							OwnsZebra:   nationMap[n[bytes.IndexByte(p, 'z')]],
						}
					}
				}
			}
		}
	}
	panic("No solution found, this can't happen!")
}
