/* Given an array of non-negative integers representing the number of steps you can take forward from each position, implement the function CanJump() which takes a slice of unsigned integers []uint as input and returns a boolean value. This function should determine if it's possible to reach and stay at the last index of the array starting from the first index, based on the steps you need to advance. Be aware that:

Each value represents the exact number of steps you must take forward from that position.
The function should return true if it's possible to reach and stay at the last index without stepping out of the array, and false otherwise.
If the input has only one element, that is the last position in the array so the function will return true but if the array is empty it returns false. */

package main

import "fmt"

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}

	curr := uint(0)
	last := uint(len(arr) - 1)

	// Advance through array based on exact steps from each position
	for curr < last {
		step := arr[curr]
		// Cannot advance if step is 0 or if step overshoots beyond the last index
		if step == 0 || step > last-curr {
			return false
		}
		curr += step
	}

	return curr == last
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
