/*Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.*/

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	arg := os.Args[1]

	//2. simple string-to-integer parser(atoi)
	if len(arg) == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	num := 0
	for i := 0; i < len(arg); i++ {
		char := arg[i]
		if char < '0' || char > '9' {		//if not a valid positive digit,fail and print 0
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
		}	
		num = num*10 + int(char-'0')
	}
	//double it is a strictly +ve number
	if num <= 0 {
	z01.PrintRune('0')
	z01.PrintRune('\n')
	return	
	}
	//3. loop thru numbers from 2 upto 'num' to find primes and sum them
	sum := 0
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			sum += i
		}
	}
	// Step 4: Turn the sum into a readable string of text
	sumText := ""
	
	if sum == 0 {
		sumText = "0"
	}
	// itoa
	for sum > 0 { 
		digit := string(rune('0' + sum % 10)) // Converts the raw math number to a text letter
		sumText = digit + sumText              // Pastes it onto the front of our text string
		sum = sum/10                                  // Chops off the last digit from our sum
	}

	// Now print our clean text string, character by character
	for i := 0; i < len(sumText); i++ {
		z01.PrintRune(rune(sumText[i]))
	}

	z01.PrintRune('\n')
	
}