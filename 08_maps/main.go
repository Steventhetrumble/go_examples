package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")

	fmt.Println(len(emails))
}
