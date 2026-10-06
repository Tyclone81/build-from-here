/*Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').

If the number of arguments is different from 2, then the program displays a newline ('\n').*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len (args) != 3 {
		z01.PrintRune ('\n')
		return
	}
	s1 := args[1]
	s2 := args[2]

	seen := make(map[rune]bool)  //map to track already displayed xters to avoid doubles

	for _, char1 := range s1 {   // loop through 1st string and print unique xters
		if !seen[char1] {
			z01.PrintRune(char1)
			seen[char1] = true
		}
	}
	for _, char2 := range s2 {
		if !seen[char2] {
			z01.PrintRune(char2)
			seen[char2] = true
		}
	}
	z01.PrintRune('\n')
}