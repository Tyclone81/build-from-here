//Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.

package main
import "fmt"
func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}