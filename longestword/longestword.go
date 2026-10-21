/*Write a function `LongestWord` that takes a `string` as an argument and returns the longest word in that string, followed by a newline (`'\n'`).
    
    - A "word" is defined as a sequence of characters delimited by spaces (`' '`), tabs (`'\t'`), or the start/end of the string.
    - If there are multiple words with the same maximum length, return the **first** one that appears in the string.
    - If the string is empty or contains only spaces/tabs, return a newline (`"\n"`)*/
	
	package main
    
    import "fmt"
    
    func LongestWord(arr string) string {
	longest := ""
	current := ""
	hasDigit := false // Tracks if the current word has any numbers

	for i := 0; i < len(arr); i++ {
		if arr[i] != ' ' && arr[i] != '\t' && arr[i] != '\n' {
			// If the character is a digit, flag it
			if arr[i] >= '0' && arr[i] <= '9' {
				hasDigit = true
			}
			current += string(arr[i])
		} else {
			// Only check the length if the word has NO digits
			if !hasDigit && len(current) > len(longest) {
				longest = current
			}
			current = "" 
			hasDigit = false // Reset the flag for the next word
		}
	}

	// Check the final word remaining after the loop ends
	if !hasDigit && len(current) > len(longest) {
		longest = current
	}

	return longest + "\n"
    }

    func main() {
	fmt.Println(LongestWord("Hello i am a girl"))               // Expected: "Hello"
	fmt.Println(LongestWord("Hel2lo i am a girl"))             // Expected: "girl" ("Hel2lo" is ignored)
	fmt.Println(LongestWord("Hello i am a girrrrrrrrrl"))       // Expected: "girrrrrrrrrl"
	fmt.Println(LongestWord("Hello i am a umemen2yanawekoce"))   // Expected: "umemenyanawekoce"
	fmt.Println(LongestWord("123 123uijjj42"))                  // Expected: "" (Both have digits)
    fmt.Print(LongestWord("The quick brown fox jumps over the lazy dog"))
    fmt.Print(LongestWord("  one   three  seven  eleven  "))
    fmt.Print(LongestWord("tie cat dog"))
    fmt.Print(LongestWord("       "))
    fmt.Print(LongestWord(""))
}
