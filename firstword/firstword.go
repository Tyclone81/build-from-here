/* Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
A word is a sequence of characters delimited by spaces or by the start/end of the argument. */
package main
import "fmt"
func FirstWord(s string) string {

	// set i to the index of the first character in the string
    i := 0

	//skip all leading spaces by moving forward
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i >= len (s) {
		return "\n"
	}
	// mark where the first word starts
	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}
	// slice from 1st character of word to one character past the last character of the word
	return s[start:i] + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}