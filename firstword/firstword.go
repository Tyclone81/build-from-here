/* Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
A word is a sequence of characters delimited by spaces or by the start/end of the argument. */
package main
import "fmt"
func FirstWord(s string) string {
    i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i >= len (s) {
		return "\n"
	}
	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[start:i] + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}