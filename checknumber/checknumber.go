/* Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false. */

package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, char := range arg {
		// Check if the character is an ASCII digit between '0' and '9'
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
	fmt.Println(CheckNumber("12345"))
	fmt.Println(CheckNumber("H e l l o"))
	fmt.Println(CheckNumber(""))
}
