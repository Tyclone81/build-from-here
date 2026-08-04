/* Write a function LastWord that takes a string and returns its last word followed by a \n.
A word is a section of string delimited by spaces or by the start/end of the string. */
package main
import "fmt"
func LastWord(s string) string{
	i := len(s)-1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	if i < 0 {
		return "\n"
	}
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return s[i+1: end+1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}