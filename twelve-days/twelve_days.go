package twelve
import "fmt"
func twelve_days(){
    twelve_strings := [12]string{"twelve Drummers Drumming", "eleven Pipers Piping", "ten Lords-a-Leaping", "nine Ladies Dancing", "eight Maids-a-Milking", "seven Swans-a-Swimming", "six Geese-a-Laying", "five Gold Rings", "four Calling Birds", "three French Hens", "two Turtle Doves", "and a Partridge in a Pear Tree."}
    fmt.Println("On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.")
    for i := len(twelve_strings)-2; i >= 0; i-- {
        temp_string := twelve_strings[i]
        for n := i+1;n<len(twelve_strings);n++ {
            temp_string += ", "
            temp_string += twelve_strings[n]
        }
        fmt.Println("On the first day of Christmas my true love gave to me:", temp_string)
    }
}
