/* Write a program that takes a string, and displays this string with exactly:
one space between words.
without spaces nor tabs at the beginning nor at the end.
with the result followed by a newline ("\n").
A "word" is defined as a part of a string delimited either by spaces/tabs, or by the start/end of the string.
If the number of arguments is not 1, or if there are no words to display, the program displays a newline("\n"). */

package main

import (
	"os"

	"github.com/01-edu/z01"
)
func main() {
	// Check if exactly 1 argument is provided (os.Args[0] is the program name, os.Args[1] is the input)
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	s := os.Args[1]
	//tracking flags
	inWord := false
	wordEnded := false

	//single-pass loop execution
	for _, char := range s {
		if char != ' ' && char != '\t' {
			if wordEnded && !inWord {		//if previous word ended and a new one is starting, print one space delimeter
				z01.PrintRune(' ')
				wordEnded = false
			}

			z01.PrintRune(char)
			inWord = true
		} else {
		// If we leave a word and step into whitespace, mark it as ended
			if inWord {
			wordEnded = true
			inWord = false
			}
		}
	} 
	//3. always print a trailing newline if args were valid
	z01.PrintRune('\n')
}
