/*
Write a function CountRepeats() that takes a string as an argument and returns the number of times characters are repeated consecutively in the string.

If the string is empty or contains no consecutive repeats, the function returns 0.
*/

package main

import "fmt"

func CountRepeats(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	// Loop through string from the second character and compare with previous character
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountRepeats("aaabbcdddd")) // 6 ('a': 2 repeats, 'b': 1 repeat, 'd': 3 repeats)
	fmt.Println(CountRepeats("abc"))        // 0
	fmt.Println(CountRepeats("aabbcc"))     // 3
	fmt.Println(CountRepeats(""))           // 0
}
