package clock
import "fmt"

// type Clock struct {
//     h int
//     m int
// }
// func New(hour, minute in) Clock {
//     h,m := fix_time(hour,minute)
//     return Clock{h,m}
// }
// func (c Clock) String () string {
//     return fmt.Sprintf("%02d:%02d",c.h,c.m)
// }
// func (c Clock) Add(minutes int) Clock {
//     return New(c.h,c.m+minutes)
// }
// func fix_time(h,m int) (int,int){
//     if(m<0 && m==-60){
//
//     }else if(m<0){
//         m-=60
//     }
//     h += m/60
//     m=mod(m,60)
//     h=mod(h, 24)
//     return h,m
// }
// func mod(x,y int) int {
//     mod := x % y
//     if mod < 0 {
//         mod += y
//     }
//     return mod
// }
// type Clock int
// func Time(hour, minute int) Clock{
//     time := (hour*60+minute) % (60 * 24)
//     if time < 0 {
//         time += 60 * 24
//     }
//     return Clock(time)
// }
// func (c Clock) String() string {
//     return fmt.Sprintf("%02d:%02d", c/60, c%60)
// }
// func (c Clock) add(minutes int) Clock {
//     return Time(0, int(c)+minutes)
// }
