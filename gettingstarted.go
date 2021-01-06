package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	helloWorldExample("Gagan") // Hello World Example
	variableExample()          // varaibel example
	fmt.Println(add(2, 3))     // function with parameter example
	val, err := add(13, 3)
	fmt.Println(val, err)
	loops()
	infiniteLoopBreak()
	// consoleInputExample()
	conditionExample("DOG")
	switchExample()
}

func helloWorldExample(name string) {
	fmt.Print("Hello World" + name)
}

func variableExample() {

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

func infiniteLoopBreak() {
	condition := false
	i := 1
	for !condition {
		i++
		fmt.Println(i)
		if i == 3 {
			condition = true
		}
	}
}

func consoleInputExample() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter animal name")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf(input)
}

func conditionExample(nameofAnimal string) string {
	x := nameofAnimal
	if x == "dog" {
		fmt.Println("Dog")
	} else if x == "cat" {
		fmt.Println("Cat")
	} else {
		fmt.Println("New Animal")
	}
	return x
}

func switchExample() {

	alphabet := "c"

	switch alphabet {
	case "A", "a":
		fmt.Println("Apple")
	case "B", "b":
		fmt.Println("Banana")
	default:
		fmt.Println("Not a Fruit")
	}
}
