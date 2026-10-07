/*Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string. This rewrite must respect the order in which these characters appear in the second string.

If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.*/

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	//allocate pointers for s1 and s2 respectively
	i := 0
	j := 0

	// loop through both strings simulitaneously
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++   //always move forward in the second string
	}
	// if our pointer i matches every xter in s1
	if i == len(s1) {
		for k := 0; k < len(s1); k++ {
			z01.PrintRune(rune(s1[k]))
		}
		z01.PrintRune('\n')
	}
}