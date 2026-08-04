/* Write a function LastWord that takes a string and returns its last word followed by a \n.
A word is a section of string delimited by spaces or by the start/end of the string. */
package main
import "fmt"
func LastWord(s string) string{

	// set i to the index of the last character in the string
	i := len(s)-1

	// move i leftwards as long these conditions are true, to remove trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}
	//mark where the last word ends i.e the last non-space character we meet
	end := i

	//move the i leftwards again but in this case until we meet an empty space
	for i >= 0 && s[i] != ' ' {
		i--
	}
	//slice one character after the space we found to one character after the last character
	return s[i+1: end+1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}