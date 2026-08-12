/* The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.

The integers can be negative.

 */

package main
import "fmt"
func Slice(a []string, nbrs... int) []string{
	if len(nbrs) == 0 {
		return nil
	}
	//set start of slice as the first integer
	start := nbrs[0]

	//handle negative integer
	if start < 0 {
		start = len(a) + start
	}
	//default end of slice is end of string
	end := len(a)

	//in cases with 2 intergers or more, end of slice is the integer at index 1
	if len(nbrs) >= 2 {
		end = nbrs[1]
		if end < 0 {
		end = len(a) + end
		}
	}
	//clamp the indices to a valid range
	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}
	if start >= end {
		return nil
	}
	//return a slice of the input string from the defined start to end
	return a[start:end]
}

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}
