/*
Write a function CountRepeats() that takes a string as an argument and returns the number of times characters are repeated consecutively in the string.

If the string is empty or contains no consecutive repeats, the function returns 0.
*/
package main

import (
	"fmt"
)

func CountRepeats(s string) string {
	if s == "" {
		return ""
	}

	result := ""
	counter := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			counter++
		} else {
			result += string(s[i])
			// Only append the counter if it's greater than 1
			if counter > 1 {
				result += string(rune(counter + '0'))
			}
			counter = 1
		}
	}

	// Handle the last character group
	result += string(s[len(s)-1])
	if counter > 1 {
		result += string(rune(counter + '0'))
	}

	return result
}

func main() {
	fmt.Println(CountRepeats("aaabbbccc")) // Output: a3b3c3
	fmt.Println(CountRepeats("abc"))       // Output: abc
	fmt.Println(CountRepeats("abbccc"))    // Output: ab2c3
	fmt.Println(CountRepeats("aaabbcdddd")) 
	fmt.Println(CountRepeats(""))
}

