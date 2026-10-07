/*Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

If there are no argument, the program displays nothing.*/

package main

import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) < 2 {
		return
	}
	//loop through each argmument passed
	for i := 0; i < len(os.Args); i++ {
		str := os.Args[i]

		// process the current string xter by xter
		for j := 0; j < len(str); j++ {
			char := str[j]

			//check if the xter is the last letter of a word
			isWordChar := (char != ' ' && char != '\t')
			isLastLetter := false

			if isWordChar {
				if j == len(str)-1 || str[j+1] == ' ' || str[j+1] == '\t' {
					isLastLetter = true
				}
			}
			// Step 3: Apply the formatting transformations
			if isLastLetter {
				// Turn into Uppercase if it's lowercase
				if char >= 'a' && char <= 'z' {
					char = char - 32
				}
			} else {
				// For all other characters, turn into Lowercase if it's uppercase
				if char >= 'A' && char <= 'Z' {
					char = char + 32
				}
			}

			z01.PrintRune(rune(char))
		}
		
		// Print a newline right after finishing each argument
		z01.PrintRune('\n')

	}
}