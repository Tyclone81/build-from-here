//Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.
//'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
package main
import "fmt"
func RepeatAlpha(s string) string {
	result := ""
	//loop through the characters of the input string one at a time
	for i := 0; i < len(s); i++ {
		//get the character at position i on the string and define it as char
		char := s[i]
		//check if the character is a letter and if it is lowercase
		if char >= 'a' && char <= 'z' {
			//define the position of the character in the string
			position := (int(char - 'a') + 1)
			//loop through repeating the character for the number of times as its position in the alphabet
			for j:= 0; j < position; j++ {
				result += string(char)
			}
			//repeat the operation but in this case for uppercase letters
		} else if char >= 'A' && char <= 'Z' {
			position := (int(char - 'A') + 1)
			for j:= 0 ; j < position; j++ {
				result += string(char)
			}
			//if character is anything other than a letter, print it as it is
		} else {
			result += string(char)
		}	
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}