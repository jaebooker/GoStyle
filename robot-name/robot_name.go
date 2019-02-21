package robotname
import (
	"fmt"
	"math/rand"
	"time"
)
type Robot string
var namesInUse = make(map[Robot]bool)
func (r *Robot) Name() string {
    rand.Seed(time.Now().UnixNano())
    for inUse, ok := string(*r) == "", true; ok && inUse; inUse, ok = namesInUse[*r] {
        *r = Robot(letter() + letter() + number())
    }
    namesInUse[*r] = true
    return string(*r)
}
func (r *Robot) Reset() {
    *r = Robot("")
}
func letter() string {
    return string(rand.Intn('Z'-'A')+'A')
}
func number() string {
    return fmt.Sprintf("%03d", rand.Intn(1000))
}
    // func nameRobot(){
    //     word.random() += newStinrg
    //     word.random() += newStinrg
    //     string(num.random()) += newStinrg
    //     string(num.random()) += newStinrg
    //     string(num.random()) += newStinrg
    // }
    // func nameRobot()
    // for i in robot.name {
    //     if i == newString {
    //         func nameRobot()
    //     }
    // }
    // vector robot {
    //     name: string
    // }
    // robot(newStinrg)
