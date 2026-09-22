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
	var words []string
	word := ""

	// Iterate through the string to extract words delimited by spaces or tabs
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	// Append the trailing word if present
	if len(word) > 0 {
		words = append(words, word)
	}

	// If there are no words, print newline and exit
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Print words with a single space between them
	for i, w := range words {
		if i > 0 {
			z01.PrintRune(' ')
		}
		for _, r := range w {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
