/* Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

The display will be followed by a newline ('\n').

If the number of arguments is different from 2, the program displays nothing. */

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check that exactly two arguments are passed (os.Args[0] is the program name)
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	// Map to keep track of characters that have already been displayed
	seen := make(map[rune]bool)

	// Iterate through the characters in the first string
	for _, char1 := range str1 {
		// If the character hasn't been displayed yet
		if !seen[char1] {
			// Check if the character exists in the second string
			for _, char2 := range str2 {
				if char1 == char2 {
					z01.PrintRune(char1)
					seen[char1] = true
					break
				}
			}
		}
	}

	// Print a newline at the end of the display
	z01.PrintRune('\n')
}
