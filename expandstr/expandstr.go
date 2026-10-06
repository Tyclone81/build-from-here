/*Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

The string will be followed by a newline ('\n').

A word, in this exercise, is a sequence of visible characters.

If the number of arguments is not 1, or if there are no word, the program displays nothing.*/

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	args := os.Args
	if len(args) != 2{
		return
	}
	s := args[1]

	// tracking flags
	inWord := false     
	hasWordBegun := false
	wordEnded := false


	for _, char := range s {
		if char != ' ' && char != '\t' {
			if wordEnded && !inWord {  // if a previous word ended and we are not inside a word, drop exactly 3 space delimeters before printing the xters
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				wordEnded = false		 //reset delimeter flag
			}
			z01.PrintRune(rune(char))
			inWord = true
			hasWordBegun = true
		} else {
			if inWord {				// If we were inside a word and hit whitespace, mark it as ended
				wordEnded = true
				inWord = false
			}
		}
	}
	if hasWordBegun {
		z01.PrintRune('\n')
	}
}
