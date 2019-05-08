package main

import "fmt"

func main() {
	//using var
	name := "Brad"

	var age = 37
	const isCool = true
	size := 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
