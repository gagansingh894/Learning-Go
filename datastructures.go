package main

import "fmt"

func main() {
	arrayExample()
}

// Arrays
func arrayExample() {
	var arr [5]int      // fixed size empty array
	fmt.Println(arr)    // print array
	fmt.Println(arr[0]) // print specific array
	arr[0] = 2          // Change array element
	fmt.Println(arr[0])
	length := len(arr) // get length of array
	fmt.Println(length)

	// alternate approach
	arr1 := []int{4, 5, 6}
	fmt.Println(len(arr1))

	// calculate array sum
	s2 := arraySum(arr1)
	fmt.Println(s2)
}

// Sum of Array
func arraySum(ar []int) int {
	lengthOfArray := len(ar)
	sum := 0

	for i := 0; i < lengthOfArray; i++ {
		sum += ar[i]
	}

	return sum
}

// Slices
// Range Slice/Array
// Maps
