package main

import "fmt"

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

// Range Slice/Array Examples
func loopArrayExample(ar []int) {
	lengthOfArray := len(ar)
	for i := 0; i < lengthOfArray; i++ {
		fmt.Println(ar[i])
	}
}

func rangeExample(ar []int) {
	for i, element := range ar {
		fmt.Println(i, element)
	}
	for _, element := range ar {
		fmt.Println(element)
	}
}

func printDuplicate(ar []int) {
	for i, e1 := range ar {
		for j, e2 := range ar {
			if j > i && e1 == e2 {
				fmt.Println(e1)
			}

		}
	}

}

// Maps
func mapExample() {
	// key value pair
	var mp map[string]int = map[string]int{
		"gagan": 1,
		"sahib": 2,
		"savy":  3,
	}
	fmt.Println(mp)         // print all
	fmt.Println(mp["savy"]) // print value based on key
	mp["gagan"] = 5         // update value
	fmt.Println(mp)
	mp["prashant"] = 10 // add new value
	fmt.Println(mp)
	delete(mp, "prashant") // delete if key is present else nothing
	fmt.Println(mp)
	// check if key exists
	val, ok := mp["gagan"] //key exists
	fmt.Println(val, ok)
	val, ok = mp["new"] // key does not exist
	fmt.Println(val, ok)
	fmt.Println(len(mp)) // print length
	var mp2 map[int]int = map[int]int{}
	fmt.Println(mp2) // empty map
}

func findFirstDuplicate(ar []int) {
	res := map[int]int{}

	for _, e := range ar {
		_, ok := res[e]
		if ok == false {
			res[e] = 1
		} else {
			fmt.Println(e)
			break
		}
	}

}

func findAllDuplicate(ar []int) {
	/*
		Function which prints all duplicate values
	*/
	res := map[int]int{}

	for _, e := range ar {
		_, ok := res[e]
		if ok == false {
			res[e] = 1
		} else {
			res[e]++
		}
	}
	for k, v := range res {
		if v != 1 {
			fmt.Println(k)
		}
	}
}
