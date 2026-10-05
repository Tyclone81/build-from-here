/*Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

The hash equation is computed as follows:
(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

If the resulting character is unprintable add 33 to it.*/

package main

import (
	"fmt"
)

func HashCode(dec string) string {
	result:= "" //empty variable to store our hashed function

	for i:= 0; i<len(dec); i++ {
		char:= dec[i]
		hash := (int(char) + len(dec)) % 127
		if hash < 33 {
			hash += 33
		}
		result += string(rune(hash))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}