/*Write a function that takes two strings's and returns the number of characters that are not included in both, without repeating characters.

If there is no unique characters return 0.
If both strings are empty return -1*/

package main

import (
	"fmt"
	"strings"
)

func WeAreUnique(str1, str2 string) int {
	// 1. Edge Case: If both strings are empty, return -1 as specified
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	// seen tracks unique characters already counted to prevent duplicates
	seen := ""
	count := 0

	// 2. Iterate through str1: find characters not in str2
	for _, ch := range str1 {
		charStr := string(ch)
		// Check if the character is absent from str2 and hasn't been counted yet
		if !strings.Contains(str2, charStr) && !strings.Contains(seen, charStr) {
			seen += charStr
			count++
		}
	}

	// 3. Iterate through str2: find characters not in str1
	for _, ch := range str2 {
		charStr := string(ch)
		// Check if the character is absent from str1 and hasn't been counted yet
		if !strings.Contains(str1, charStr) && !strings.Contains(seen, charStr) {
			seen += charStr
			count++
		}
	}

	// 4. Returns the total count of unique characters (returns 0 if none found)
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
