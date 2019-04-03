package queenattack

func queenattack(array[]int w, array[]int b) bool wa bool ba {
    if (w[0] == b[0]) || (b[1] == b[1]) {
        return true, true
    }
    for(i:=w[0],x:=w[1], i<8,i++,x++){
        if (b[0] == i) && (b[1] == x) {
            return true, true
        }
    }
    for(i:=b[0],x:=b[1],i<8,i++,x++){
        if (w[0] == i) && (w[1] == x) {
            return true, true
        }
    }
    for(i:=w[0],x:=w[1], i>-1,i--,x++){
        if (b[0] == i) && (b[1] == x) {
            return true, true
        }
    }
    for(i:=b[0],x:=b[1],i>-1,i--,x++){
        if (w[0] == i) && (w[1] == x) {
            return true, true
        }
    }

}
