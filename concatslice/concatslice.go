/* Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.*/

package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	// Directly append the entire second slice to the first one using the ... unpack operator
	return append(slice1, slice2...)
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}