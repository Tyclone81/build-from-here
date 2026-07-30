//Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).
//If the number of arguments is different from 3, the program displays nothing.
//If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').
package main
import (
	"os"
	"github.com/01-edu/z01"
) 
func main() {
	// get command line arguments
	args:= os.Args

	//confirm we only have 4 arguments i.e, program name and the 3 input arguments
	if len(args) != 4 {
		return
	}
	// define the arguments
	str := args[1]
	oldchar :=args[2]
	newchar :=args[3]

	//loop through the input string character by character checking if the character equals the oldchar
	for i := 0; i < len(str); i++ {
		char:= str[i]
		if string(char) == oldchar {
			//loop through replacing the oldchar with newchar, for all the instances it occurs in the string
			for _, r := range newchar {
				z01.PrintRune(rune(r))
			}
		} else {
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}

