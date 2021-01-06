package main

import (
	"fmt"
	"errors"
)

func main() {
	hello_world_example() // Hello World Example
	variable_example() // varaibel example
	fmt.Println(add(2,3)) // function with parameter example
	val, err := add(2,3)
	fmt.Println(val, err)
	loops()
}

func hello_world_example() {
	fmt.Print("Hello World\n")
}

func variable_example() {

	// default values
	var a int 
	var b float32
	var c string
	var d bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func add(a int, b int) (int, error) {
	sum := a + b
	if sum > 10 {
		return -1, errors.New("xyz")
	}
	return sum, nil
}

func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}