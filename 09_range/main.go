package main

import "fmt"

func main() {
	ids := []int{22, 76, 54, 23, 11, 2}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("Id: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
}
