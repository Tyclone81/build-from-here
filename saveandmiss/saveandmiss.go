/*Write a function called SaveAndMiss() that takes a string and an int as an argument. The function should move through the string in sets determined by the int, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the string is reached. Return a string containing the saved characters.

If the int is 0 or a negative number return the original string.*/

package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	result := ""
	save := true 	//flag to track whether to save or skip the current chunk

	//loop through the string in sets determined by the int num
	for i := 0; i < len(arg); i += num {
		end := i + num 	// define the boundary of the cuurent chunk
		if end > len(arg) {
			end = len(arg)
		}

		// save the chunk if the toggle is true
		if save {
			result += arg[i:end]
		}
		save = !save		//toggle the flag for the next chunk
	}
	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}