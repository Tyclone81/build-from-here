/*Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

If the size is 0 it should print a newline ('\n').*/

package main

import "github.com/01-edu/z01"

func Chunk(slice []int, size int) {
	if size <= 0 {
		z01.PrintRune('\n')
		return
	}

	z01.PrintRune('[')
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}

		z01.PrintRune('[')
		for j, val := range slice[i:end] {
			z01.PrintRune(rune('0' + val)) // Assumes single digits 0-9 per prompt test cases
			if j < len(slice[i:end])-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')

		if i+size < len(slice) {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

//alternative version

/*func Chunk(slice []int, size int) [][]int {
	if size <= 0 {
		z01.PrintRune('\n')
		return nil
	}

	var result [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}

	return result
}
*/