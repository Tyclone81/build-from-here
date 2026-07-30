//Write a function that takes a string and returns a new string that replaces every character with the number of duplicates and the character itself, deleting the extra duplications.
//The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.
package main
import "fmt"
func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	count := 1

	for i := 0; i < len(s); i++ {
		//condition to make sure we stop at the last character on the string
		if i < len(s)-1 && s[i] == s[i+1] {
			count++
		} else {
			//otherwise convert count to string using itoa and then add current character as a string
			result += itoa(count) + string(s[i])
			count = 1 // reset count back to one for the next group of characters
		}
	}
	return result
}
// the itoa helper function we used to convert count to string
//basically the itoa function
func itoa(n int) string {
	if n==0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string(byte('0' + n%10)) + result
		n = n/10
	}
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}