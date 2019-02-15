package beer
import (
	"fmt"
)
func Verses() {
	fmt.Printf("99 bottles of beer on the wall, 99 bottles of beer.\n")
	for beer := 98; beer > 1; beer-- {
		fmt.Printf("Take one down and pass it around, %d bottles of beer on the wall.\n\n", beer)
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", beer, beer)
	}
	fmt.Printf("Take one down and pass it around, 1 bottle of beer on the wall.\n\n")
	fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
	fmt.Printf("Take it down and pass it around, no more bottles of beer on the wall.\n\n")
	fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
	fmt.Printf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
}
