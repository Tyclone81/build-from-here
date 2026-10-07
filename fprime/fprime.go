/*Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

Factors must be displayed in ascending order and separated by *.

If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.*/

package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main () {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	// 1. sting-to-integer passer(atoi)
	num := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return
		}
		num = num*10 + int(char-'0')
	}
	if num <= 1 {
		return
	}
	//2. the core loop
	divisor := 2
	first := true  // init first at true to indicate we are printing 1st prime

	for num > 1 {
		if num%divisor == 0 {
			if ! first {
				z01.PrintRune('*')
			}
			first = false  //set first to false

			//print the current divisor right away
			temp := divisor
			digits := ""
			for temp > 0 {
				digits = string(rune('0' + temp%10)) + digits
				temp = temp/10
			}
			for _, r := range digits {
				z01.PrintRune(r)
			}
			//shrink the number
			num = num/divisor
		} else {
			divisor++  	//move to the next number
		}
	}
	z01.PrintRune('\n')
}