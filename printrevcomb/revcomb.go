//Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.
//These combinations are separated by a comma and a space.
package main
import "github.com/01-edu/z01"
func main () {
	//a flag that tracks if we are printing the first combination
	first := true

	for a := 9; a >= 2; a-- {
		for b := a-1; b >= 1; b-- {
			for c := b-1; c >= 0; c-- {
				if !first {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				//Set flag to false so the next loops trigger the comma
				first = false
				z01.PrintRune(rune('0' + a))
				z01.PrintRune(rune('0' + b))
				z01.PrintRune(rune('0' + c))
			}
		}
	}
	z01.PrintRune('\n')
}