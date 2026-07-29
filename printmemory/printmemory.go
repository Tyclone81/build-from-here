//Write a function that takes (arr [10]byte), and displays the memory as in the example.
//After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.
//The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.
package main
import "github.com/01-edu/z01"
func PrintMemory(arr [10]byte) {
	//loop through the array to print the hexadecimal values of the bytes
	for i:=0; i<len(arr); i++ {
		b := arr[i]

		//print the first hexadecimal nibble of the byte, i.e the quotient of the byte
		first := b/16

		if first < 10 {
			z01.PrintRune(rune('0' + first))

		// if a,b,c,d,e,f
		} else {
			z01.PrintRune(rune('a' + first - 10))
		}
		//print second nibble i.e remainder
		second := b%16
		if second < 10 {
			z01.PrintRune(rune('0' + second))
		} else {
			z01.PrintRune(rune('a' + second - 10))
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	//loop through the array again to convert each of the hexadecimal character formulated to its ASCII representation
	for i:=0; i < len(arr); i++ {
		//printable ASCII characters i.e 33 through 126
		if arr[i] >= 33 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			//a dot for unprintable characters
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}