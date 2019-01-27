package hamming

func Distance(a, b string) (int, error) {
    distance := 0
    for i := 0; i < len(a)-1; i++ {
        if a[i] != b[i] {
            distance += 1
        }
    }
    return distance, nil
}
