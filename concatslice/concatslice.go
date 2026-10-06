/* Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.*/

package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	// easiest option
	// Directly append the entire second slice to the first one using the ... unpack operator
	// return append(slice1, slice2...)

	result := make([]int, len(slice1) + len(slice2))

	idx := 0

	for i := 0; i < len(slice1); i++ {
		result[idx] = slice1[i]
		idx++
	}
	for i := 0; i < len(slice2); i++ {
		result[idx] = slice2[i]
		idx++
	}
	return result
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}