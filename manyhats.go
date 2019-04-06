package manyhats

func manyhats([100]bool h) [100]bool nh{
    nh := [100]bool
    for (i:=0;i<len(h);i++){
        c := 0
        for (x:=i+1;x<len(h)-1;x++){
            if h[x]{
                c++
            }
        }
        if nh[i-1] != nil {
            c2 := 0
            for (y:=i-1;i>0;i--){
                if nh[i]{
                    c2 += 1
                }
            }
            c =- c2
            if c % 2 == 0 && c2 % 2 != 0 {
                nh[i] = true
            }
            if c % 2 != 0 && c2 % 2 == 0 {
                nh[i] = false
            }
            if c % 2 == 0 && c2 % 2 == 0 {
                nh[i] = true
            }
        } else {
            if c % 2 == 0 {
                nh[i] = true
            } else {
                nh[i] = false
            }
        }
    }
    return nh
}

func main(){

}
