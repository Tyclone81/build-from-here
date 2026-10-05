/*
  Write a function main.go:5 that takes a single word as a string and dissects it into its constituent anatomical components: its prefix, root, and suffix.
  The function must return three string values:
	 (prefix, root, suffix)
 The function must recognize the following 10 prefixes:
	 • "un", "re", "pre", "mis", "dis", "over", "under", "anti", "inter", "sub"
  The function must recognize the following 10 suffixes:
  • "ing", "ed", "er", "est", "ly", "ness", "ment", "tion", "able", "ful" */

  	package main

	import (
			"fmt"
	)
    
    // WordAnatomy dissects a word into its constituent prefix and suffix (returning 2 strings, excluding the root).
    func WordAnatomy(word string) (string, string) {
		//fixed-size array containing the 10 given prefixes
        prefixes := [...]string{
            "un", "re", "pre", "mis", "dis",
            "over", "under", "anti", "inter", "sub",
        }
		//fixed-size array containing the 10 given suffixes
        suffixes := [...]string{
            "ing", "ed", "er", "est", "ly",
            "ness", "ment", "tion", "able", "ful",
        }
    
		//empty variables to track the longest valid prefix and suffix found 
        bestPrefix := ""
        bestSuffix := ""
    
        // 1. Identify the longest matching prefix by looping through each prefix "p" in the given array
        for _, p := range prefixes {
            lp := len(p)	//get the length of each prefix "p" and then check it as follows

            if len(word) >= lp && word[:lp] == p {	//does 'word' have atleast 'lp' xters...and does the beginning slice of 'word' from start to lp match 'p'

                if lp > len(bestPrefix) {
                    bestPrefix = p
                }
            }
        }
    
        // 2. Identify the longest matching suffix by looping through each suffix 's' in the given arrray
        for _, s := range suffixes {
            ls := len(s)	//extract length of current suffix in the array and check it as we did in the prefix

            if len(word) >= ls && word[len(word)-ls:] == s {	//does the 'word' have atleast 'ls' xters and does the ending slice(from len(word)-ls) match 's'

                if ls > len(bestSuffix) {
                    bestSuffix = s
                }
            }
        }
    
        // 3. Prevent overlap: prefix and suffix cannot intersect
        if len(bestPrefix)+len(bestSuffix) > len(word) {	//means they overlap/intersect

            if len(bestPrefix) >= len(bestSuffix) {	//prioritize whichever affix is longer and if equal, keep prefix
                bestSuffix = "" 	//if the above condition is true, drop suffix so prefix remains otherwise drop prefix
            } else {
                bestPrefix = ""
            }
        }
    
        // 4. Return prefix and suffix only (root is intentionally omitted per requirements)
        return bestPrefix, bestSuffix
    }

	func main() {
	tests := []string{
		"unhappy",
		"runner",
		"redoing",
		"kindness",
		"planet",
		"careful",
		"misunderstanding",
		"greatest",
		"underpaid",
	}

	for _, word := range tests {
		// WordAnatomy now returns only 2 values: prefix and suffix
		prefix, suffix := WordAnatomy(word)
		fmt.Printf("%s -> Prefix:%q Suffix:%q\n", word, prefix, suffix)
	}
}
