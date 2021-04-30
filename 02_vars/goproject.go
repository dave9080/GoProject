package main

import "fmt"

func main() {

	var name string = "Dave"
	var age int32 = 42
	const isCool = true

	// Shorthand
	// name := "Dave"

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)

}
