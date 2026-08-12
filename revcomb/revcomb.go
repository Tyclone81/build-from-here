/* Write a program that prints in descending order on a single line all unique combinations of three different digits so that the first digit is greater than the second and the second is greater than the third.
These combinations are separated by a comma and a space.*/

package main
import "github.com/01-edu/z01"
func main() {
	// flag first to track we are printing the first combination
	first := true

	// init 3 characters and set them in a way that the next character is smaller than the previous
	for a := 9; a >= 2; a-- {
		for b := a-1; b >= 1; b-- {
			for c := b-1; c >= 0; c-- {

				// condition to print comma and space incase we are not printing the first combination
				if ! first{
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				//set the flag used to track the first combination to false to show we are no longer printing the first combination
				first = false
				
				z01.PrintRune(rune('0' + a))
				z01.PrintRune(rune('0' + b))
				z01.PrintRune(rune('0' + c))
			}
		}
	}
	z01.PrintRune('\n')
}