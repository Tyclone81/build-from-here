//Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.
//'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...
package main
import "fmt"
func RepeatAlpha(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'a' && char <= 'z' {
			position := (int(char - 'a') + 1)
			for j:= 0; j < position; j++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			position := (int(char - 'A') + 1)
			for j:= 0 ; j < position; j++ {
				result += string(char)
			}
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