/* Write a function that converts a string from camelCase to snake_case. If the string is empty, return an empty string. If the string is not camelCase, return the string unchanged. If the string is camelCase, return the snake_case version of the string. Rules for writing in camelCase:

The word does not end on a capitalized letter (CamelCasE).
No two capitalized letters shall follow directly each other (CamelCAse).
Numbers or punctuation are not allowed in the word anywhere (camelCase1). */
package main
import "fmt"
func CamelToSnakeCase(s string) string{
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		//check if the character is not a letter
		if s[i] < 'A' || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			return s
		}
		//check that the text does not have capital letters following each other
		if i < len(s) - 1 && s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
		//check that text does not end with a capital letter case
		lastChar := s[len(s) - 1]
		if lastChar >= 'A' && lastChar <= 'Z' {
			return s
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		//if the character is uppercase and its index position is not zero, add underscore before it
		if s[i] >= 'A' && s[i] <= 'Z' {
			if i != 0 {
				result += "_"
			}
			//convert the upper to lower case (optional in this exercise)
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}