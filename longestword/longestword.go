/*Write a function `LongestWord` that takes a `string` as an argument and returns the longest word in that string, followed by a newline (`'\n'`).
    
    - A "word" is defined as a sequence of characters delimited by spaces (`' '`), tabs (`'\t'`), or the start/end of the string.
    - If there are multiple words with the same maximum length, return the **first** one that appears in the string.
    - If the string is empty or contains only spaces/tabs, return a newline (`"\n"`)*/
	
	package main
    
    import "fmt"
    
    // LongestWord returns the first word with the maximum length in string s
    func LongestWord(s string) string {
        longest := ""
        currentWord := ""
    
        for i := 0; i < len(s); i++ {
            // Check if the current character is a delimiter (space or tab)
            if s[i] == ' ' || s[i] == '\t' {
                // When encountering a delimiter, check if the word just ended is strictly longer
                if len(currentWord) > len(longest) {
                    longest = currentWord
                }
                // Reset current word buffer
                currentWord = ""
            } else {
                // Append character to the current word being built
                currentWord += string(s[i])
            }
        }
    
        // Check the trailing word at the end of the string
        if len(currentWord) > len(longest) {
            longest = currentWord
        }
    
        // Return the longest word followed by a newline
        return longest + "\n"
    }
    
    func main() {
        fmt.Print(LongestWord("The quick brown fox jumps over the lazy dog"))
        fmt.Print(LongestWord("  one   three  seven  eleven  "))
        fmt.Print(LongestWord("tie cat dog"))
        fmt.Print(LongestWord("       "))
        fmt.Print(LongestWord(""))
    }