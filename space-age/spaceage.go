package space
type Planet string
func Age(age float64, planet Planet) float64 {
    var year float64
    switch planet {
    case "Mercury":
    year = 0.2408467
case "Venus":
    year = 0.61519726
case "Earth":
    year = 1.0
case "Mars":
    year = 1.8808158
case "Jupiter":
    year = 11.862615
case "Saturn":
    year = 29.447498
case "Uranus":
    year = 84.016846
case "Neptune":
    year = 164.79132
}
return age / (31557600 * year)
//     var year float64
// switch p {
// case "Mercury":
//     year = 0.2408467
// case "Venus":
//     year = 0.61519726
// case "Earth":
//     year = 1.0
// case "Mars":
//     year = 1.8808158
// case "Jupiter":
//     year = 11.862615
// case "Saturn":
//     year = 29.447498
// case "Uranus":
//     year = 84.016846
// case "Neptune":
//     year = 164.79132
// }
// return seconds / (EarthSeconds * year)
    // seconds := float64(31557600)
    // new_age := float64(Age)
    // new_age = new_age/seconds
    // if Planet == "Earth" {
    //     return new_age
    // }
    // if Planet == "Mercury" {
    //     return new_age * 0.2408467
    // }
    // if Planet == "Venus" {
    //     return new_age * 0.61519726
    // }
    // if Planet == "Mars" {
    //     return new_age * 1.8808158
    // }
    // if Planet == "Jupiter" {
    //     return new_age * 11.862615
    // }
    // if Planet == "Saturn" {
    //     return new_age * 29.447498
    // }
    // if Planet == "Uranus" {
    //     return new_age * 84.016846
    // }
    // if Planet == "Neptune" {
    //     return new_age * 164.79132
    // }
    // return 0
}
