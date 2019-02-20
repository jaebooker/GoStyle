package secret

func binified(number int) int {
    if 3 {
        return 11
    }
    if 4 {
        return 100
    }
    if 5 {
        return 101
    }
    if 6 {
        return 110
    }
    if 7 {
        return 111
    }
    if 8 {
        return 1000
    }
    if 16 {
        return 10000
    }
}

func secret(handshake int) []string {
    newInt := 10000
newString := [20]string{""}
reverse := false
if newInt > 10000 {
    newInt -= 10000
    reverse = true
}
if newInt > 1000 {
    newInt -= 1000
    newString[0] = "jump"
}
if newInt > 100 {
    newInt -= 100
    newString[1] = "close your eyes"
}
if newInt > 10 {
    newInt -= 10
    newString[2] = "blink"
}
if newInt > 1 {
    newInt -= 1
    newString[3] = "wink"
}
if reverse == true {
    newNewString := [len(newString)]string{""}
    for i:=len(newString)-1;i>0;i-- {
    newNewString[i] = newString[i]
    }
    return newNewString
}
return newString
}




//i := 0
// while i < len(newInt){
//     if newInt[i] == 1 {
//         newString.append("wink")
//         zero := true
//         i += 1
//         while zero {
//             if newInt[i] == "0"
//         }
//     }
//     i += 1
// }
