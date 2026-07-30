//Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.
//The second int must be between 2 and 36. If not, the function returns -1.
//If the first int is negative, reverse the sign and count the digits.
package main
import "fmt"
func DigitLen(n, base int) int {
	//check base validity
	if base < 2 || base > 36 {
		return -1
	}
	//convert negative int to positive
	if n < 0 {
		n = -n
	}
	count := 0
	//for int above 0 apply the digitlen logic
	for n > 0 {
		n = n/base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}