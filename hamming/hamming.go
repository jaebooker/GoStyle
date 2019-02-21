package hamming

func Distance(a, b string) (int, error) {
    distance := 0
    if len(a) < 1 {
        return 0, nil
    }
    if len(a) != len(b) {
        return 0, nil
    }
    if len(a) < 2 {
        if a != b {
            distance += 1
        }
        return distance, nil
    }
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance += 1
        }
    }
    return distance, nil
}






// for char, pos := range a {
//     charmatch := false
//     for char2, pos2 := range b {
//         if (char != char2) && (pos == pos2) && (charmatch == false) {
//             distance += 1
//             charmatch = true
//         }
//     }
// }
