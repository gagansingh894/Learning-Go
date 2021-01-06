package main

import "fmt"

func main() {
	arrayExample()
	slicesExample()
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
func slicesExample() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[:3] //slice
	s2 := s[:1]         //slice of slice
	fmt.Println(x)
	fmt.Printf("%T", x)
	fmt.Println(s)
	fmt.Println(s2)

	// implicit
	var a []int = []int{5, 6, 7, 8, 9}
	a = append(a, 10) // append to slice
	fmt.Println(a)

	// make example
	b := make([]int, 5)
	fmt.Println(b)
	fmt.Printf("%T", b)
}

// Range Slice/Array
// Maps
