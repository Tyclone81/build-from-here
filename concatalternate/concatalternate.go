/*Write a function ConcatAlternate() that receives two slices of an int as arguments and returns a new slice with the result of the alternated values of each slice.

The input slices can be of different lengths.
The new slice should start with an element of the largest slice.
If the slices are of equal length, the new slice should return the elements of the first slice first and then the elements of the second slice.*/

package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	var first, second []int  // variables to store which slice comes first and second

	// 1. Identify the largest slice to start with
	if len(slice1) > len(slice2) {
		first = slice1
		second = slice2
	} else if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	} else {
		// If equal, slice1 must come first
		first = slice1
		second = slice2
	}

	result := []int{}
	
	// 2. Safely alternate up to the smaller slice's limit
	minLen := len(second)
	for i := 0; i < minLen; i++ {
		result = append(result, first[i])
		result = append(result, second[i])
	}
	// 3. Dump the remaining elements of the larger slice straight to the end
	for i := minLen; i < len(first); i++ {
		result = append(result, first[i])
	}
	return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}