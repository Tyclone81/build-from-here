/* Write a program named hiddenp that takes two strings as arguments. The program should check if the first string s1 is hidden in the second s2. s1 is considered hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1, but not necessarily consecutively.

If s1 is hidden in s2, the program should display 1 followed by a newline.
If s1 is not hidden in s2, the program should display 0 followed by a newline.
If s1 is an empty string, it is considered hidden in any string.
If the number of arguments is different from 2, the program should display nothing.*/


/* N/B: We only need a single index (i) for s1, and we can break early as soon as all characters of s1 are matched.*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check for exactly two arguments (os.Args[0] is the program name)
	if len(os.Args) != 3 {
		return
	}

	s1 := []rune(os.Args[1])
	i := 0

	// Scan s2 rune by rune; advance i when characters match in order
	for _, r := range os.Args[2] {
		if i == len(s1) {
			break // Early exit: all characters in s1 have already been matched
		}
		if r == s1[i] {
			i++
		}
	}

	// If i reached len(s1), all characters were found (also handles empty s1 naturally)
	if i == len(s1) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}