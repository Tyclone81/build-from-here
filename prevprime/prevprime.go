/*Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.
If there are no primes inferior to the int passed as parameter the function should return 0 */
package main
import "fmt"
func FindPrevPrime(nb int) int {
	// least prime is 2
	for i := nb; i >= 2; i-- {

		//flag isPrime to track we are printing a prime number
		isPrime := true
		
		// loop to confirm the previous numbers are either prime or not 
		for j := 2; j * j <= i; j++ {
			if i % j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return i
		}
	}
	//if no primes inferior to nb return 0
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}