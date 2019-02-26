package zerba
type Solution struct {
	DrinksWater string
	OwnsZebra string
}
type House struct {
    smoke string
    drink string
    pet string
    nationality string
    color string
    location int
}
func SolvePuzzle()Solution{
    color := [5]string{"red","ivory","yellow","green","blue"}
    person := [5]string{"Japanese","Ukrainian","Englishman","",""}
    pet := [5]string{"","","","",""}
    smoke := [5]string{"","","","",""}
    drink := [5]string{"water","orange juice","coffee","",""}
    h1 := House{}
    h2 := House{}
    h3 := House{}
    h4 := House{}
    h5 := House{}
    h1.nationality = person[2]
    h1.color = color[0]
}


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
