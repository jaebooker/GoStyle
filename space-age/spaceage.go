package space

func main(Age int, Planet string) float64 {
    seconds := float64(31557600)
    new_age := float64(Age)
    new_age = new_age/seconds
    if Planet == "Earth" {
        return new_age
    }
    if Planet == "Mercury" {
        return new_age * 0.2408467
    }
    if Planet == "Venus" {
        return new_age * 0.61519726
    }
    if Planet == "Mars" {
        return new_age * 1.8808158
    }
    if Planet == "Jupiter" {
        return new_age * 11.862615
    }
    if Planet == "Saturn" {
        return new_age * 29.447498
    }
    if Planet == "Uranus" {
        return new_age * 84.016846
    }
    if Planet == "Neptune" {
        return new_age * 164.79132
    }
    return 0
}
