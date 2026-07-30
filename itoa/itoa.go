//Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as anint into a number represented as a string.
//For this exercise the handling of the signs + or - does have to be taken into account.
//the func mannually converts each digit to its character representation
package main
import "fmt"
func Itoa(n int) string {

	if n == 0 {
		return "0"
	}
	//a flag isNeg set to false to track negative numbers
	isNeg := false
	//convert negative numbers to positive
	if n < 0 {
		isNeg = true
		n = -n
	}

	result := ""
	//loop through the digits one after another converting them to their respecitve character representations
	for n > 0 {
		result = string(byte('0' + n%10)) + result
		n = n/10
	}
	//append neg sign before outputting neg numbers
	if isNeg {
		result = "-" + result
	}
	return result

}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}