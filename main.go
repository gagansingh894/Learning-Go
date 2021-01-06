package main

import "fmt"

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
	arrayExample()
	slicesExample()
}
